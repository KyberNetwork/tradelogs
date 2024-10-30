package promotionparser

import (
	"errors"

	storageTypes "github.com/KyberNetwork/tradelogs/v2/pkg/storage/promotees"
	ethereumTypes "github.com/ethereum/go-ethereum/core/types"
)

var (
	ErrInvalidTopic = errors.New("invalid order topic")
)

type Parser interface {
	Parse(log ethereumTypes.Log, blockTime uint64) (storageTypes.Promotee, error)
	Topics() []string
	Contract() string
	LogFromContract(log ethereumTypes.Log) bool
	Address() string
}
