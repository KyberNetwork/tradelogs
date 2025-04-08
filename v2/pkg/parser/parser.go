package parser

import (
	"errors"

	storageTypes "github.com/KyberNetwork/tradelogs/v2/pkg/storage/tradelogs/types"
	"github.com/KyberNetwork/tradelogs/v2/pkg/types"
	ethereumTypes "github.com/ethereum/go-ethereum/core/types"
)

var (
	ErrInvalidTopic  = errors.New("invalid order topic")
	ErrNotFoundTrade = errors.New("not found log")
)

type Parser interface {
	Parse(log ethereumTypes.Log, blockTime uint64) ([]storageTypes.TradeLog, error)
	Topics() []string
	Exchange() string
	UseTraceCall() bool
	ParseWithCallFrame(callFrame types.CallFrame, log ethereumTypes.Log, blockTime uint64) ([]storageTypes.TradeLog, error)
	LogFromExchange(log ethereumTypes.Log) bool
	Address() string
	ParseTransferEvent(txHash string, block_number, timestamp uint64, call types.CallFrame) []interface{}
}
