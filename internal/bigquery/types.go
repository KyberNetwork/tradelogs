package bigquery

import (
	internalTypes "github.com/KyberNetwork/tradelogs/internal/types"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type BQLog struct {
	Index          int64     `bigquery:"log_index"`
	TxHash         string    `bigquery:"transaction_hash"`
	TxIndex        int64     `bigquery:"transaction_index"`
	Address        string    `bigquery:"address"`
	Data           string    `bigquery:"data"`
	Topics         []string  `bigquery:"topics"`
	BlockTimestamp time.Time `bigquery:"block_timestamp"`
	BlockNumber    int64     `bigquery:"block_number"`
	BlockHash      string    `bigquery:"block_hash"`
}

func hex2Bytes(raw string) []byte {
	if strings.HasPrefix(raw, "0x") {
		return common.Hex2Bytes(raw[2:])
	}

	return common.Hex2Bytes(raw)
}

func (log BQLog) ToEthLog() types.Log {
	topics := make([]common.Hash, len(log.Topics))
	for id, topic := range log.Topics {
		topics[id] = common.HexToHash(topic)
	}

	return types.Log{
		Address:     common.HexToAddress(log.Address),
		Topics:      topics,
		Data:        hex2Bytes(log.Data),
		BlockNumber: uint64(log.BlockNumber),
		TxHash:      common.HexToHash(log.TxHash),
		TxIndex:     uint(log.TxIndex),
		BlockHash:   common.HexToHash(log.BlockHash),
		Index:       uint(log.Index),
	}
}

type BQTraceCall struct {
	TransactionHash string    `bigquery:"transaction_hash"`
	FromAddress     string    `bigquery:"from_address"`
	ToAddress       string    `bigquery:"to_address"`
	Input           string    `bigquery:"input"`
	Output          string    `bigquery:"output"`
	BlockTimestamp  time.Time `bigquery:"block_timestamp"`
	BlockNumber     int64     `bigquery:"block_number"`
}

func (b BQTraceCall) ToTraceCall() internalTypes.CallFrame {
	return internalTypes.CallFrame{
		From:   b.FromAddress,
		To:     b.ToAddress,
		Input:  b.Input,
		Output: b.Output,
	}
}
