package bigquery

import (
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type BQLog struct {
	Index          uint      `bigquery:"log_index"`
	TxHash         string    `bigquery:"transaction_hash"`
	TxIndex        uint      `bigquery:"transaction_index"`
	Address        string    `bigquery:"address"`
	Data           string    `bigquery:"data"`
	Topics         []string  `bigquery:"topics"`
	BlockTimestamp time.Time `bigquery:"block_timestamp"`
	BlockNumber    uint64    `bigquery:"block_number"`
	BlockHash      string    `bigquery:"block_hash"`
}

func (log BQLog) ToEthLog() types.Log {
	topics := make([]common.Hash, len(log.Topics))
	for id, topic := range log.Topics {
		topics[id] = common.HexToHash(topic)
	}

	return types.Log{
		Address:     common.HexToAddress(log.Address),
		Topics:      topics,
		Data:        common.Hex2Bytes(log.Data),
		BlockNumber: log.BlockNumber,
		TxHash:      common.HexToHash(log.TxHash),
		TxIndex:     log.TxIndex,
		BlockHash:   common.HexToHash(log.BlockHash),
		Index:       log.Index,
	}
}
