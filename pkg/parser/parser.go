package parser

import (
	"github.com/KyberNetwork/tradelogs/pkg/storage"
	"github.com/ethereum/go-ethereum/core/types"
)

type Parser interface {
	Parse(log types.Log, blockTime uint64) (storage.TradeLog, error)
	Topics() []string
}
