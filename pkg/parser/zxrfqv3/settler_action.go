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
	settlerOtcSelfFundedFunction FunctionABI = "SETTLER_OTC_SELF_FUNDED(address,((address,uint256),uint256,uint256),address,bytes,address,uint256)"
	metatxnRFQVipFunction        FunctionABI = "METATXN_RFQ_VIP(address,((address,uint256),uint256,uint256),address,bytes,((address,uint256),uint256,uint256),bytes)"
	rfqVIPFunction               FunctionABI = "RFQ_VIP(address,((address,uint256),uint256,uint256),address,bytes,((address,uint256),uint256,uint256))"
)

const (
	settlerOtcSelfFundedName FunctionName = "SETTLER_OTC_SELF_FUNDED"
	metatxnRFQVipName        FunctionName = "METATXN_RFQ_VIP"
	rfqVIPName               FunctionName = "RFQ_VIP"
)

var mSettlerActionName map[FunctionName]FunctionABI

const (
	MethodIdDecodeParamOfFillOrderSelfFundedHex = "0a164181"
	methodIdDecodeParamOfFillOrderVIPHex        = "bcd804f7"
)

var (
	methodIdDecodeParamOfFillOrderSelfFunded decoder.Bytes4
	methodIdDecodeParamOfFillOrderVIP        decoder.Bytes4
)

func init() {
	mSettlerActionName = map[FunctionName]FunctionABI{
		settlerOtcSelfFundedName: settlerOtcSelfFundedFunction,
		metatxnRFQVipName:        metatxnRFQVipFunction,
		rfqVIPName:               rfqVIPFunction,
	}

	methodIdDecodeParamOfFillOrderSelfFunded = loadMethodId(MethodIdDecodeParamOfFillOrderSelfFundedHex)
	methodIdDecodeParamOfFillOrderVIP = loadMethodId(methodIdDecodeParamOfFillOrderVIPHex)
}

func loadMethodId(methodHex string) decoder.Bytes4 {
	byteMethodId, err := hex.DecodeString(methodHex)
	if err != nil {
		log.Fatalf("failed to decode method id %v, err: %s", methodHex, err)
	}
	methodId, err := decoder.GetBytes4(byteMethodId)
	if err != nil {
		log.Fatalf("failed to get method id %v, err: %s", methodHex, err)
	}
	return methodId
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
