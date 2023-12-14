package bigquery

import (
	"cloud.google.com/go/bigquery"
	"context"
	"errors"
	"github.com/KyberNetwork/tradelogs/pkg/parser"
	"go.uber.org/zap"
	"google.golang.org/api/iterator"
)

type WorkerState string

const (
	limit = 5_000
)

type Datasource struct {
	l              *zap.SugaredLogger
	client         *bigquery.Client
	parserTopicMap map[string]parser.Parser
}

func NewDatasource(
	projectID string, parserNameMap map[string]parser.Parser,
) (*Datasource, error) {
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

	return &Datasource{
		l:              zap.S(),
		client:         client,
		parserTopicMap: parserTopicMap,
	}, nil
}

func (w *Datasource) QueryEventLogs(
	ctx context.Context, dateString string,
	minTime, maxTime, offset int64,
	topics []string,
) (int, []BQLog, error) {
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

	iter, err := q.Read(ctx)
	if err != nil {
		return 0, nil, err
	}

	return w.logsFromRowIterator(iter)
}

func (w *Datasource) QueryTraceCall(
	ctx context.Context,
	dateString string,
	minTime, maxTime int64,
	toAddress string,
	fourBytesMethods []string,
	limit int64,
	offset int64,
) (int, []BQTraceCall, error) {
	q := w.client.Query(
		`SELECT transaction_hash, from_address, to_address, input, output,block_timestamp, block_number
        FROM ` + "`bigquery-public-data.crypto_ethereum.traces`" + `
		WHERE
			DATE(block_timestamp) = @date
			AND block_timestamp >= TIMESTAMP_SECONDS(@minTime)
			AND block_timestamp <= TIMESTAMP_SECONDS(@maxTime) 
			AND to_address = @toAddress
			AND status = @status
			AND error is null
			AND substr(input, 0, 10) in UNNEST(@fourBytesMethods)
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
			Name:  "toAddress",
			Value: toAddress,
		},
		{
			Name:  "fourBytesMethods",
			Value: fourBytesMethods,
		},
		{
			Name:  "limit",
			Value: limit,
		},
		{
			Name:  "offset",
			Value: offset,
		},
		{
			Name:  "status",
			Value: 1,
		},
	}

	iter, err := q.Read(ctx)
	if err != nil {
		return 0, nil, err
	}

	count := 0
	bqTraceCalls := make([]BQTraceCall, 0, limit)
	for {
		var row BQTraceCall
		if err := iter.Next(&row); err != nil {
			if errors.Is(err, iterator.Done) {
				return count, bqTraceCalls, nil
			}
			return 0, nil, err
		}

		count++
		bqTraceCalls = append(bqTraceCalls, row)
	}
}

func (w *Datasource) logsFromRowIterator(iter *bigquery.RowIterator) (int, []BQLog, error) {
	count := 0
	res := make([]BQLog, 0, limit)
	for {
		var row BQLog
		if err := iter.Next(&row); err != nil {
			if errors.Is(err, iterator.Done) {
				return count, res, nil
			}
			return 0, nil, err
		}

		count++
		res = append(res, row)
	}
}
