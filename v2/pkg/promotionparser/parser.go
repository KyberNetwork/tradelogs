package promotionparser

import (
	"errors"

	storageTypes "github.com/KyberNetwork/tradelogs/v2/pkg/storage/promotees"
	"github.com/KyberNetwork/tradelogs/v2/pkg/types"
	ethereumTypes "github.com/ethereum/go-ethereum/core/types"
)

var (
	ErrInvalidTopic = errors.New("invalid order topic")
)

type Parser interface {
	Parse(log ethereumTypes.Log, blockTime uint64) (storageTypes.Promotee, error)
	Topics() []string
	Contract() string
	UseTraceCall() bool
	ParseWithCallFrame(callFrame types.CallFrame, log ethereumTypes.Log, blockTime uint64) (storageTypes.Promotee, error)
	LogFromContract(log ethereumTypes.Log) bool
	Address() string
}
