package server

import (
	"encoding/json"
	"time"

	"github.com/KyberNetwork/tradelogs/pkg/dune"
	"github.com/KyberNetwork/tradelogs/pkg/storage"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"go.uber.org/zap"
)

const (
	limit = 100
	retry = 5
)

type DuneWorker struct {
	client *dune.Client
	s      *storage.Storage
}

func NewDuneWoker(c *dune.Client, s *storage.Storage) *DuneWorker {
	return &DuneWorker{
		client: c,
		s:      s,
	}
}

// func (d *DuneWorker) backfill(l *zap.SugaredLogger,
// 	queryID int64, from, to uint64, ps parser.Parser, topic string) error {
// 	l.Infow("start backfill data", "query", queryID, "topic", topic, "from", from, "to", to)
// 	queryRes, err := d.client.ExecuteQuery(queryID, topic, from, to)
// 	if err != nil {
// 		return err
// 	}

// 	l.Infow("executing query", "data", queryRes)
// 	errCount := 0
// 	for {
// 		time.Sleep(5 * time.Second)
// 		state, err := d.client.ExecuteState(queryRes.ExecutionID)
// 		if err != nil {
// 			l.Errorw("error when get state", "error", err)
// 			errCount++
// 			if errCount > retry {
// 				return err
// 			}
// 			continue
// 		}
// 		l.Infow("execute query", "state", state)
// 		if state.IsExecutionFinished {
// 			break
// 		}
// 	}
// 	l.Infow("executed query")

// 	var progress uint64 = 0
// 	errCount = 0

// 	for {
// 		data := []storage.TradeLog{}

// 		time.Sleep(5 * time.Second)
// 		logs, rowCount, err := d.client.GetLastestExecuteResult(queryID, limit, progress)
// 		if err != nil {
// 			l.Errorw("error when collect data", "error", err)
// 			errCount++
// 			if errCount > retry {
// 				return err
// 			}
// 			continue
// 		}

// 		l.Infow("collect data", "progress", progress, "len", len(logs), "total", rowCount)
// 		for _, l := range logs {
// 			ts, err := time.Parse("2006-01-02 15:04:05.999 UTC", l.BlockTime)
// 			if err != nil {
// 				return err
// 			}
// 			parsedLog, err := ps.Parse(DuneLogToETHLog(l), uint64(ts.Unix()))
// 			if err != nil {
// 				return err
// 			}
// 			data = append(data, parsedLog)
// 		}
// 		l.Infow("parsed data", "len", len(data))
// 		if err = d.s.Insert(data); err != nil {
// 			return err
// 		}
// 		progress += limit + 1
// 		if progress >= rowCount {
// 			break
// 		}
// 	}
// 	return nil
// }

func DuneLogToETHLog(log dune.DuneLog) types.Log {
	data, err := hexutil.Decode(log.Data)
	if err != nil {
		return types.Log{}
	}
	topics := []common.Hash{}
	if log.Topic0 != "" {
		topics = append(topics, common.HexToHash(log.Topic0))
	}
	if log.Topic1 != "" {
		topics = append(topics, common.HexToHash(log.Topic1))
	}
	if log.Topic2 != "" {
		topics = append(topics, common.HexToHash(log.Topic2))
	}
	if log.Topic3 != "" {
		topics = append(topics, common.HexToHash(log.Topic3))
	}
	return types.Log{
		Address:     common.HexToAddress(log.ContractAddress),
		Topics:      topics,
		Data:        data,
		BlockNumber: log.BlockNumber,
		TxHash:      common.HexToHash(log.TxHash),
		TxIndex:     log.TxIndex,
		BlockHash:   common.HexToHash(log.BlockHash),
		Index:       log.Index,
	}
}

func (d *DuneWorker) backfillOneInch(l *zap.SugaredLogger, queryID int64, eventHash string) error {
	var progress uint64 = 0
	errCount := 0
	for {
		data := []storage.TradeLog{}
		logs := []dune.OneInchDuneLog{}
		time.Sleep(time.Second)
		rowCount, err := d.client.GetLastestExecuteResult(queryID, limit, progress, &logs)
		if err != nil {
			l.Errorw("error when collect data", "error", err)
			errCount++
			if errCount > retry {
				return err
			}
			continue
		}

		l.Infow("collected data", "progress", progress, "len", len(logs), "total", rowCount)
		for _, l := range logs {
			parse, err := OneInchDuneLogToTrade(l, eventHash)
			if err != nil {
				return err
			}
			data = append(data, parse)
		}
		if err = d.s.Insert(data); err != nil {
			return err
		}
		progress += limit + 1
		if progress >= rowCount {
			break
		}
	}
	return nil
}

type OneInchOrder struct {
	Maker      dune.BigInt `json:"maker"`
	MakerAsset dune.BigInt `json:"makerAsset"`
	TakerAsset dune.BigInt `json:"takerAsset"`
}

func OneInchDuneLogToTrade(l dune.OneInchDuneLog, eventHash string) (storage.TradeLog, error) {
	ts, err := time.Parse("2006-01-02 15:04:05.999 UTC", l.BlockTime)
	if err != nil {
		return storage.TradeLog{}, err
	}
	order := OneInchOrder{}
	if err := json.Unmarshal([]byte(l.Order), &order); err != nil {
		return storage.TradeLog{}, err
	}
	return storage.TradeLog{
		OrderHash:        l.Output2,
		Maker:            order.Maker.Hex(),
		MakerToken:       order.MakerAsset.Hex(),
		TakerToken:       order.TakerAsset.Hex(),
		MakerTokenAmount: l.Output0.String(),
		TakerTokenAmount: l.Output1.String(),
		ContractAddress:  l.ContractAddress,
		BlockNumber:      l.BlockNumber,
		TxHash:           l.TxHash,
		LogIndex:         l.EventIndex,
		Timestamp:        uint64(ts.UnixMilli()),
		EventHash:        eventHash,
	}, nil
}
