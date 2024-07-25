package zxrfqv3

import (
	"encoding/hex"
	"github.com/KyberNetwork/tradelogs/pkg/decoder"
	"golang.org/x/crypto/sha3"
	"log"
)

const (
	executeFunctionName = "execute"
	balanceOf           = "balanceOf"
	actionParamName     = "actions"
)

type FunctionName string
type FunctionABI string

const (
	settlerOtcSelfFundedFunction     FunctionABI = "SETTLER_OTC_SELF_FUNDED(address,((address,uint256),uint256,uint256),address,bytes,address,uint256)"
	metatxnSettlerOtcPermit2Function FunctionABI = "METATXN_SETTLER_OTC_PERMIT2(address,((address,uint256),uint256,uint256),address,bytes,((address,uint256),uint256,uint256))"
)

const (
	settlerOtcSelfFundedName     FunctionName = "SETTLER_OTC_SELF_FUNDED"
	metatxnSettlerOtcPermit2Name FunctionName = "METATXN_SETTLER_OTC_PERMIT2"
)

const (
	MethodIdDecodeParamOfFillOrderSelfFundedHex = "0a164181"
)

var mSettlerActionName map[FunctionName]FunctionABI
var methodIdDecodeParamOfFillOrderSelfFunded decoder.Bytes4

func init() {
	mSettlerActionName = map[FunctionName]FunctionABI{
		settlerOtcSelfFundedName:     settlerOtcSelfFundedFunction,
		metatxnSettlerOtcPermit2Name: metatxnSettlerOtcPermit2Function,
	}

	byteMethodId, err := hex.DecodeString(MethodIdDecodeParamOfFillOrderSelfFundedHex)
	if err != nil {
		log.Fatalf("failed to decode method id: %s", err)
	}
	methodId, err := decoder.GetBytes4(byteMethodId)
	if err != nil {
		log.Fatalf("failed to get method id: %s", err)
	}
	methodIdDecodeParamOfFillOrderSelfFunded = methodId
}

func getSettlerAction() map[decoder.Bytes4]FunctionName {
	mSettlerAction := make(map[decoder.Bytes4]FunctionName)

	for name, sig := range mSettlerActionName {
		hash := sha3.NewLegacyKeccak256()
		hash.Write([]byte(sig))
		hashBytes := hash.Sum(nil)

		// Extract the first 4 bytes (function selector)
		selector := decoder.Bytes4{hashBytes[0], hashBytes[1], hashBytes[2], hashBytes[3]}

		mSettlerAction[selector] = name
	}
	return mSettlerAction
}
