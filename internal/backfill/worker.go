package backfill

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/KyberNetwork/tradelogs/internal/bigquery"
	"github.com/KyberNetwork/tradelogs/pkg/parser/oneinch"

	"github.com/KyberNetwork/tradelogs/pkg/parser"
	"github.com/KyberNetwork/tradelogs/pkg/storage"
	"go.uber.org/zap"
)

type WorkerState string

var (
	defaultOneinchRouterContractAddresses = "0x1111111254eeb25477b68fb85ed929f73a960582"
	defaultOneinchRFQMethodIds            = []string{"0x3eca9c0a", "0x9570eeee", "0x5a099843"}
)

const (
	limit            = 5_000
	minBlockDateTime = "2018-08-08T00:00:00+00:00"

	stateStopped WorkerState = "stopped"
	stateRunning WorkerState = "running"
)

var (
	ErrNoParser      = errors.New("no parser")
	ErrWorkerRunning = errors.New("worker is running")
	ErrNoTopic       = errors.New("no topic")

	minBlockTime, _ = time.Parse(time.RFC3339, minBlockDateTime)
)

type Worker struct {
	l                  *zap.SugaredLogger
	bigqueryDatasource *bigquery.Datasource
	storage            *storage.Storage
	parserTopicMap     map[string]parser.Parser
	parserNameMap      map[string]parser.Parser
	oneinchParser      *oneinch.Parser
	state              WorkerState
}

func NewWorker(
	projectID string, storage *storage.Storage, parserNameMap map[string]parser.Parser,
	oneinchParser *oneinch.Parser,
) (*Worker, error) {
	bigqueryDatasource, err := bigquery.NewDatasource(projectID, parserNameMap)
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
		l:                  zap.S(),
		bigqueryDatasource: bigqueryDatasource,
		storage:            storage,
		parserTopicMap:     parserTopicMap,
		parserNameMap:      parserNameMap,
		state:              stateStopped,
		oneinchParser:      oneinchParser,
	}, nil
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
		count, bqLogs, err := w.bigqueryDatasource.QueryEventLogs(
			context.Background(),
			temp.Format("2006-01-02"),
			minTs, maxTs, offset, topics,
		)
		if err != nil {
			l.Errorw("Failed to queryDateData", "err", err, "temp", temp, "offset", offset)
			return
		}

		l.Infow("Successfully get QueryEventLogs", "temp", temp, "count", count)

		tradelogs := make([]storage.TradeLog, 0, len(bqLogs))
		for _, bqLog := range bqLogs {
			log, err := w.parseLog(bqLog)
			if err != nil {
				l.Errorw("Failed to parse the bigquery log to ethereum log", "err", err)
				return
			}
			tradelogs = append(tradelogs, log)
		}

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

func (w *Worker) backfillOneInch(contractAddress string, fourBytesMethodIds []string, minTime, maxTime time.Time) {
	l := w.l.With("minTime", minTime, "maxTime", maxTime)
	l.Info("Start running backfill oneinch")
	w.state = stateRunning
	defer func() {
		w.l.Info("Stop running worker")
		w.state = stateStopped
	}()

	dateOfBlockTime := endOfDay(maxTime)
	offset := int64(0)
	minTs, maxTs := minTime.Unix(), maxTime.Unix()
	for dateOfBlockTime.After(minTime) {
		count, traceCalls, err := w.bigqueryDatasource.QueryTraceCall(
			context.Background(),
			dateOfBlockTime.Format("2006-01-02"),
			minTs, maxTs,
			contractAddress,
			fourBytesMethodIds,
			limit,
			offset,
		)
		if err != nil {
			l.Errorw("Failed to query trace calls", "err", err, "temp", dateOfBlockTime, "offset", offset)
			return
		}

		l.Infow("Fetch trace call successfully", "temp", dateOfBlockTime, "count", count)
		tradeLogs := make([]storage.TradeLog, 0, len(traceCalls))
		for _, traceCall := range traceCalls {
			var (
				err error
			)
			order := storage.TradeLog{
				TxHash:          traceCall.TransactionHash,
				BlockNumber:     uint64(traceCall.BlockNumber),
				Timestamp:       uint64(traceCall.BlockTimestamp.Unix()),
				ContractAddress: traceCall.ToAddress,
			}

			order, err = w.oneinchParser.ParseFromInternalCall(order, traceCall.ToTraceCall())
			if err != nil {
				l.Errorw("Failed to parse the internal call from trace call bigquery", "err", err)
				continue
			}
			tradeLogs = append(tradeLogs, order)
		}

		err = w.storage.Insert(tradeLogs)
		if err != nil {
			l.Errorw("Failed to insert tradelogs", "err", err, "temp", dateOfBlockTime, "tradelogs", tradeLogs)
			return
		}

		if count < limit {
			offset = 0
			dateOfBlockTime = dateOfBlockTime.Add(-24 * time.Hour)
		} else {
			offset += limit
		}
	}

	l.Info("Backfill oneinch successfully!")
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

func (w *Worker) BackFillOneInch(from, to int64, contractAddress string, fourByteMethodIds []string) error {
	if w.state == stateRunning {
		return ErrWorkerRunning
	}

	minTime := time.Unix(from, 0).UTC()
	if minTime.Before(minBlockTime) {
		minTime = minBlockTime
	}

	now := time.Now().UTC()
	maxTime := time.Unix(to, 0).UTC()
	if maxTime.After(now) {
		maxTime = now
	}

	if contractAddress == "" {
		contractAddress = defaultOneinchRouterContractAddresses
	}
	if len(fourByteMethodIds) == 0 {
		fourByteMethodIds = defaultOneinchRFQMethodIds
	}

	go w.backfillOneInch(contractAddress, fourByteMethodIds, minTime, maxTime)
	return nil
}

func (w *Worker) parseLog(row bigquery.BQLog) (storage.TradeLog, error) {
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
