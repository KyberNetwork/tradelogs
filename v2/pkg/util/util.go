package util

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/KyberNetwork/tradelogs/v2/pkg/types"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

var (
	ErrInvalidTopic  = errors.New("invalid order topic")
	ErrNotFoundTrade = errors.New("not found log")
)

func AssignLogIndexes(cf *types.CallFrame, index int) int {
	subCallIndex := hexutil.Uint(0)
	for i := range cf.Logs {
		for subCallIndex < cf.Logs[i].Position {
			index = AssignLogIndexes(&cf.Calls[subCallIndex], index)
			subCallIndex++
		}
		cf.Logs[i].Index = index
		index++
	}
	for subCallIndex < hexutil.Uint(len(cf.Calls)) {
		index = AssignLogIndexes(&cf.Calls[subCallIndex], index)
		subCallIndex++
	}
	return index
}

func ConvertHexToDecimal(value string) (string, error) {
	if len(value) <= 2 {
		return value, nil
	}
	amount := new(big.Int)
	amount, success := amount.SetString(value[2:], 16)
	amountStr := amount.String()
	if !success {
		amountStr = value
		return amountStr, fmt.Errorf("cannot convert hex to decimal")
	}
	return amountStr, nil
}
