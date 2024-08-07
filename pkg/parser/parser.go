package parser

import (
	"github.com/KyberNetwork/tradelogs/pkg/storage"
	tradingTypes "github.com/KyberNetwork/tradinglib/pkg/types"
	ethereumTypes "github.com/ethereum/go-ethereum/core/types"
)

const (
	ExZeroX         = "zerox"
	EXZeroXV3       = "zeroxV3"
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

	Addr1InchV6      = "0x111111125421cA6dc452d289314280a0f8842A65"
	AddrBebop        = "0xbbbbbBB520d69a9775E85b458C58c648259FAD5F"
	AddrHashflowV3   = "0x24b9d98FABF4DA1F69eE10775F240AE3dA6856fd"
	AddrKyberswap    = "0x6131B5fae19EA4f9D964eAc0408E4408b66337b5"
	AddrKyberswapRFQ = "0x7A819Fa46734a49D0112796f9377E024c350FB26"
	AddrParaswap     = "0xe92b586627ccA7a83dC919cc7127196d70f55a06"
	AddrUniswapX     = "0x00000011F84B9aa48e5f8aA8B9897600006289Be"
	Addr0x           = "0xDef1C0ded9bec7F1a1670819833240f027b25EfF"
	Deployer0xV3     = "0x00000000000004533Fe15556B1E086BB1A72cEae"
)

type Parser interface {
	Parse(log ethereumTypes.Log, blockTime uint64) (storage.TradeLog, error)
	Topics() []string
	Exchange() string
	UseTraceCall() bool
	ParseWithCallFrame(callFrame *tradingTypes.CallFrame, log ethereumTypes.Log, blockTime uint64) (storage.TradeLog, error)
	LogFromExchange(log ethereumTypes.Log) bool
}
