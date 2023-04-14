package bigquery

import (
	"context"

	"cloud.google.com/go/bigquery"
	"github.com/KyberNetwork/tradelogs/internal/parser"
	"github.com/KyberNetwork/tradelogs/internal/storage"
	"go.uber.org/zap"
)

const (
	limit             = 1_000
	minBlockTimestamp = "2015-08-08 06:22:55 UTC"
)

type Worker struct {
	l           *zap.SugaredLogger
	client      *bigquery.Client
	storage     *storage.Storage
	parserMap   map[string]parser.Parser
	queryTopics []string
}

func NewWorker(
	projectID string, storage *storage.Storage, parsers ...parser.Parser,
) (*Worker, error) {
	client, err := bigquery.NewClient(
		context.Background(),
		projectID,
	)
	if err != nil {
		return nil, err
	}

	p := make(map[string]parser.Parser)
	topics := make([]string, 0)
	for _, ps := range parsers {
		psTopics := ps.Topics()
		topics = append(topics, psTopics...)
		for _, topic := range psTopics {
			p[topic] = ps
		}
	}

	return &Worker{
		l:           zap.S(),
		storage:     storage,
		parserMap:   p,
		client:      client,
		queryTopics: topics,
	}, nil
}

func (w *Worker) queryDateData(
	ctx context.Context, dateString string, offset int64,
) (*bigquery.RowIterator, error) {
	q := w.client.Query(
		`SELECT *
        FROM ` + "`bigquery-public-data.crypto_ethereum.logs`" + `
		WHERE 
			DATE(block_timestamp) = @date
			AND ARRAY_LENGTH(topics) > 0
			AND topics[OFFSET(0)] IN @topics
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
			Name:  "topics",
			Value: w.queryTopics,
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

func (w *Worker) queryRangeOfTimeData(
	ctx context.Context, fromTime, toTime, offset int64,
) (*bigquery.RowIterator, error) {
	q := w.client.Query(
		`SELECT *
        FROM ` + "`bigquery-public-data.crypto_ethereum.logs`" + `
		WHERE 
			block_timestamp	 >= TIMESTAMP_MILLIS(@fromTime)
			AND block_timestamp <= TIMESTAMP_MILLIS(@toTime) 
			AND ARRAY_LENGTH(topics) > 0
			AND topics[OFFSET(0)] IN @topics
		ORDER BY 
			block_timestamp DESC
		LIMIT @limit
		OFFSET @offset 
		`,
	)

	q.Parameters = []bigquery.QueryParameter{
		{
			Name:  "fromTime",
			Value: fromTime,
		},
		{
			Name:  "toTime",
			Value: toTime,
		},
		{
			Name:  "topics",
			Value: w.queryTopics,
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
