package promotionparser

import (
	"errors"

	"github.com/KyberNetwork/tradelogs/pkg/promotionstorage"
	"github.com/KyberNetwork/tradelogs/pkg/types"
	ethereumTypes "github.com/ethereum/go-ethereum/core/types"
)

const (
	Promotion1InchV2 = "oneinchv2"
	Promotion1InchV1 = "oneinchv1"

	AddrPr1InchV2 = "0xf55684bc536487394b423e70567413fab8e45e26"
	AddrPr1InchV1 = "0xcb8308fcb7bc2f84ed1bea2c016991d34de5cc77"
)

var (
	ErrInvalidTopic = errors.New("invalid order topic")
)

type Parser interface {
	Parse(log ethereumTypes.Log, blockTime uint64) (promotionstorage.Promotee, error)
	Topics() []string
	Contract() string
	UseTraceCall() bool
	ParseWithCallFrame(callFrame types.CallFrame, log ethereumTypes.Log, blockTime uint64) (promotionstorage.Promotee, error)
	LogFromContract(log ethereumTypes.Log) bool
	Address() string
}
