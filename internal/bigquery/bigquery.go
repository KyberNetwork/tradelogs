package bigquery

import (
	"context"
	"errors"
	"fmt"
	"time"

	"cloud.google.com/go/bigquery"
	"github.com/KyberNetwork/tradelogs/internal/storage"
	"github.com/KyberNetwork/tradelogs/pkg/parser"
	"go.uber.org/zap"
	"google.golang.org/api/iterator"
)

type WorkerState string

const (
	limit            = 5_000
	minBlockDateTime = "2018-08-08T00:00:00+00:00"

	stateStopped WorkerState = "stopped"
	stateRunning WorkerState = "running"
)

var (
	ErrNoTopic       = errors.New("no topic")
	ErrNoParser      = errors.New("no parser")
	ErrWorkerRunning = errors.New("worker is running")

	minBlockTime, _ = time.Parse(time.RFC3339, minBlockDateTime)
)

type Worker struct {
	l              *zap.SugaredLogger
	client         *bigquery.Client
	storage        *storage.Storage
	parserTopicMap map[string]parser.Parser
	parserNameMap  map[string]parser.Parser
	state          WorkerState
}

func NewWorker(
	projectID string, storage *storage.Storage, parserNameMap map[string]parser.Parser,
) (*Worker, error) {
	client, err := bigquery.NewClient(
		context.Background(),
		projectID,
	)
	if err != nil {
		return nil, err
	}

	parserTopicMap := make(map[string]parser.Parser)
	for _, ps := range parserNameMap {
		for _, topic := range ps.Topics() {
			parserTopicMap[topic] = ps
		}
	}

	return &Worker{
		l:              zap.S(),
		client:         client,
		storage:        storage,
		parserTopicMap: parserTopicMap,
		parserNameMap:  parserNameMap,
		state:          stateStopped,
	}, nil
}

func (w *Worker) queryDateData(
	ctx context.Context, dateString string,
	minTime, maxTime, offset int64,
	topics []string,
) (*bigquery.RowIterator, error) {
	q := w.client.Query(
		`SELECT *
        FROM ` + "`bigquery-public-data.crypto_ethereum.logs`" + `
		WHERE 
			DATE(block_timestamp) = @date
			AND block_timestamp >= TIMESTAMP_SECONDS(@minTime)
			AND block_timestamp <= TIMESTAMP_SECONDS(@maxTime) 
			AND ARRAY_LENGTH(topics) > 0
			AND topics[OFFSET(0)] IN UNNEST(@topics)
		ORDER BY 
			block_timestamp DESC
		LIMIT @limit
		OFFSET @offset 
		`,
	)

	q.Parameters = []bigquery.QueryParameter{
		{
			Name:  "date",
			Value: dateString,
		},
		{
			Name:  "minTime",
			Value: minTime,
		},
		{
			Name:  "maxTime",
			Value: maxTime,
		},
		{
			Name:  "topics",
			Value: topics,
		},
		{
			Name:  "limit",
			Value: limit,
		},
		{
			Name:  "offset",
			Value: offset,
		},
	}

	return q.Read(ctx)
}

func (w *Worker) logsFromRowIterator(iter *bigquery.RowIterator) (int, []storage.TradeLog, error) {
	count := 0
	res := make([]storage.TradeLog, 0, limit)
	for {
		var row BQLog
		if err := iter.Next(&row); err != nil {
			if errors.Is(err, iterator.Done) {
				return count, res, nil
			}
			return 0, nil, err
		}

		count++
		log, err := w.parseLog(row)
		if err != nil {
			return 0, nil, err
		}

		res = append(res, log)
	}
}

func (w *Worker) parseLog(row BQLog) (storage.TradeLog, error) {
	if len(row.Topics) == 0 {
		return storage.TradeLog{}, ErrNoTopic
	}

	ps := w.parserTopicMap[row.Topics[0]]
	if ps == nil {
		return storage.TradeLog{}, fmt.Errorf("%w: %s", ErrNoTopic, row.Topics[0])
	}

	ehtLog := row.ToEthLog()
	return ps.Parse(ehtLog, uint64(row.BlockTimestamp.Unix()))
}

func endOfDay(t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, 23, 59, 59, 1e9-1, t.Location())
}

func (w *Worker) run(minTime, maxTime time.Time, topics []string) {
	l := w.l.With("minTime", minTime, "maxTime", maxTime)
	l.Info("Start running worker")
	w.state = stateRunning
	defer func() {
		w.l.Info("Stop running worker")
		w.state = stateStopped
	}()

	temp := endOfDay(maxTime)
	offset := int64(0)
	minTs, maxTs := minTime.Unix(), maxTime.Unix()
	for temp.After(minTime) {
		iter, err := w.queryDateData(
			context.Background(),
			temp.Format("2006-01-02"),
			minTs, maxTs, offset, topics,
		)
		if err != nil {
			l.Errorw("Failed to queryDateData", "err", err, "temp", temp, "offset", offset)
			return
		}
		l.Infow("Successfully queryDateData", "temp", temp)

		count, tradelogs, err := w.logsFromRowIterator(iter)
		if err != nil {
			l.Errorw("Failed to get logsFromRowIterator",
				"err", err, "temp", temp, "offset", offset)
			return
		}
		l.Infow("Successfully get logsFromRowIterator", "temp", temp, "count", count)

		err = w.storage.Insert(tradelogs)
		if err != nil {
			l.Errorw("Failed to insert tradelogs", "err", err, "temp", temp, "tradelogs", tradelogs)
			return
		}

		if count < limit {
			offset = 0
			temp = temp.Add(-24 * time.Hour)
		} else {
			offset += limit
		}
	}

	l.Info("Backfill successfully!")
}

func (w *Worker) topicsFromExchanges(exchanges []string) ([]string, error) {
	topics := make([]string, 0)
	if len(exchanges) == 0 {
		for topic := range w.parserTopicMap {
			topics = append(topics, topic)
		}
		return topics, nil
	}

	for _, ex := range exchanges {
		parser, ok := w.parserNameMap[ex]
		if !ok {
			return nil, fmt.Errorf("%w: %s", ErrNoParser, ex)
		}
		topics = append(topics, parser.Topics()...)
	}
	return topics, nil
}

// BackFillAllData fills data from minBlockTime to now.
func (w *Worker) BackFillAllData(exchanges []string) error {
	topics, err := w.topicsFromExchanges(exchanges)
	if err != nil {
		w.l.Errorw("Failed to get topicsFromExchanges", "err", err)
		return err
	}

	if w.state == stateRunning {
		return ErrWorkerRunning
	}

	go w.run(minBlockTime, time.Now().UTC(), topics)
	return nil
}

// BackFillPartialData fills data fromTime, toTime (unix seconds).
func (w *Worker) BackFillPartialData(fromTime, toTime int64, exchanges []string) error {
	topics, err := w.topicsFromExchanges(exchanges)
	if err != nil {
		w.l.Errorw("Failed to get topicsFromExchanges", "err", err)
		return err
	}

	if w.state == stateRunning {
		return ErrWorkerRunning
	}

	// minTime = max(fromTime, minBlockTime)
	minTime := time.Unix(fromTime, 0).UTC()
	if minTime.Before(minBlockTime) {
		minTime = minBlockTime
	}

	// maxTime = min(toTime, now)
	now := time.Now().UTC()
	maxTime := time.Unix(toTime, 0).UTC()
	if maxTime.After(now) {
		maxTime = now
	}
	go w.run(minTime, maxTime, topics)
	return nil
}
