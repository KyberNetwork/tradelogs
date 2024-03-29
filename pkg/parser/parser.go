package parser

import (
	"github.com/KyberNetwork/tradelogs/pkg/storage"
	"github.com/ethereum/go-ethereum/core/types"
)

const (
	ExZeroX         = "zerox"
	ExParaswap      = "paraswap"
	ExTokenlon      = "tokenlon"
	ExZeroXRFQ      = "zerox_rfq"
	ExKs            = "kyberswap"
	ExHashflow      = "hashflow"
	ExKsRFQ         = "kyberswapRFQ"
	Ex1Inch         = "1inch"
	ExHashflowV3    = "hashflowV3"
	ExParaswapTaker = "paraswap_taker"
	ExUniswapX      = "uniswapx"
	ExNative        = "native"
)

type Parser interface {
	Parse(log types.Log, blockTime uint64) (storage.TradeLog, error)
	Topics() []string
	Exchange() string
}
