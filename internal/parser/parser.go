package parser

import (
	"github.com/KyberNetwork/tradelogs/internal/storage"
	"github.com/ethereum/go-ethereum/core/types"
)

type Parser interface {
	Parse(log types.Log, blockNumber, blockTime uint64) (storage.TradeLogs, error)
	Topics() []string
}
