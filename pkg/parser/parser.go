package parser

import (
	"github.com/KyberNetwork/tradelogs/pkg/storage"
	"github.com/KyberNetwork/tradelogs/pkg/types"
	ethereumTypes "github.com/ethereum/go-ethereum/core/types"
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
	Ex1InchV6       = "1inchV6"
	ExHashflowV3    = "hashflowV3"
	ExParaswapTaker = "paraswap_taker"
	ExUniswapX      = "uniswapx"
	ExNative        = "native"
	ExBebop         = "bebop"
)

type Parser interface {
	Parse(log ethereumTypes.Log, blockTime uint64) (storage.TradeLog, error)
	Topics() []string
	Exchange() string
	UseTraceCall() bool
	ParseWithCallFrame(_ types.CallFrame, log ethereumTypes.Log) (storage.TradeLog, error)
}
