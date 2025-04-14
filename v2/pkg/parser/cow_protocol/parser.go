package cowtradeparser

import (
	cowStorage "github.com/KyberNetwork/tradelogs/v2/pkg/storage/cow_protocol"
	ethereumTypes "github.com/ethereum/go-ethereum/core/types"
)

type TradeParser interface {
	Parse(log ethereumTypes.Log, blockTime uint64) (cowStorage.CowTrade, error)
	Topics() []string
	IsMatchLog(log ethereumTypes.Log) bool
}

type TransferParser interface {
	Parse(log ethereumTypes.Log, blockTime uint64) (cowStorage.CowTransfer, error)
	Topics() []string
	IsMatchLog(log ethereumTypes.Log) bool
}
