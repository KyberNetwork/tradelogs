package zxrfqv3

import (
	"github.com/KyberNetwork/tradelogs/pkg/decoder"
	"golang.org/x/crypto/sha3"
)

type FunctionName string
type FunctionABI string

const (
	settlerOtcSelfFundedFunction FunctionABI = "SETTLER_OTC_SELF_FUNDED(address,((address,uint256),uint256,uint256),address,bytes,address,uint256)"
	metatxnRFQVipFunction        FunctionABI = "METATXN_RFQ_VIP(address,((address,uint256),uint256,uint256),address,bytes,((address,uint256),uint256,uint256))"
	rfqVIPFunction               FunctionABI = "RFQ_VIP(address,((address,uint256),uint256,uint256),address,bytes,((address,uint256),uint256,uint256),bytes)"
	rfqFunction                  FunctionABI = "RFQ(address,((address,uint256),uint256,uint256),address,bytes,address,uint256)"

	settlerOtcSelfFundedName FunctionName = "SETTLER_OTC_SELF_FUNDED"
	metatxnRFQVipName        FunctionName = "METATXN_RFQ_VIP"
	rfqVIPName               FunctionName = "RFQ_VIP"
	rfqName                  FunctionName = "RFQ"

	executeFunctionName = "execute"
	actionParamName     = "actions"
)

var mSettlerActionName map[FunctionName]FunctionABI

func init() {
	mSettlerActionName = map[FunctionName]FunctionABI{
		settlerOtcSelfFundedName: settlerOtcSelfFundedFunction,
		metatxnRFQVipName:        metatxnRFQVipFunction,
		rfqVIPName:               rfqVIPFunction,
		rfqName:                  rfqFunction,
	}
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
