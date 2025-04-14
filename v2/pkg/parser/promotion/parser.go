package promotion

import (
	storageTypes "github.com/KyberNetwork/tradelogs/v2/pkg/storage/promotees"
	ethereumTypes "github.com/ethereum/go-ethereum/core/types"
)

type Parser interface {
	Parse(log ethereumTypes.Log, blockTime uint64) (storageTypes.Promotee, error)
	Topics() []string
	Contract() string
	IsMatchLog(log ethereumTypes.Log) bool
	Address() string
}
