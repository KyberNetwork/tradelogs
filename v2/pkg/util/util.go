package util

import (
	"errors"

	"github.com/KyberNetwork/tradelogs/v2/pkg/types"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

const (
	NativeTokenAddress = "0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee"
)

type TransferType string

const (
	TransferTypeNative TransferType = "native"
	TransferTypeERC20  TransferType = "ERC20"
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
