package types

import (
	tradingTypes "github.com/KyberNetwork/tradinglib/pkg/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	ethereumTypes "github.com/ethereum/go-ethereum/core/types"
	"math/big"
)

type TraceCallResponse struct {
	Jsonrpc string    `json:"jsonrpc"`
	ID      int       `json:"id"`
	Result  CallFrame `json:"result"`
}

type CallLog struct {
	Address common.Address `json:"address"`
	Topics  []common.Hash  `json:"topics"`
	Data    string         `json:"data"`
}

func (l CallLog) ToEthereumLog() ethereumTypes.Log {
	return ethereumTypes.Log{
		Address: l.Address,
		Topics:  l.Topics,
		Data:    common.Hex2Bytes(l.Data),
	}
}

// We will not using tradinglib.CallFrame because in the log field it required transactionHash
// But when we get traceCall from the response, it doesn't have transactionHash
type CallFrame struct {
	From    string      `json:"from"`
	Gas     string      `json:"gas"`
	GasUsed string      `json:"gasUsed"`
	To      string      `json:"to"`
	Input   string      `json:"input"`
	Output  string      `json:"output"`
	Calls   []CallFrame `json:"calls"`
	Value   string      `json:"value"`
	Type    string      `json:"type"`
	Logs    []CallLog   `json:"logs"`
}

func ConvertCallFrame(callFrame *tradingTypes.CallFrame) CallFrame {
	var result CallFrame

	if callFrame == nil {
		return CallFrame{}
	}

	for _, subCall := range callFrame.Calls {
		if subCall != nil {
			subCallFrame := ConvertCallFrame(subCall)
			result.Calls = append(result.Calls, subCallFrame)
		}
	}

	to := ""
	if callFrame.To != nil {
		to = callFrame.To.Hex()
	}
	gas := (*hexutil.Big)(big.NewInt(int64(callFrame.Gas))).String()
	gasUsed := (*hexutil.Big)(big.NewInt(int64(callFrame.GasUsed))).String()
	result.From = callFrame.From.String()
	result.Gas = gas
	result.GasUsed = gasUsed
	result.To = to
	result.Input = hexutil.Encode(callFrame.Input)
	result.Output = hexutil.Encode(callFrame.Output)
	result.Value = callFrame.Value.String()
	result.Type = callFrame.Type.String()

	for _, log := range callFrame.Logs {
		result.Logs = append(result.Logs, convertLog(log))
	}
	return result
}

func convertLog(log *ethereumTypes.Log) CallLog {
	return CallLog{
		Address: log.Address,
		Topics:  log.Topics,
		Data:    hexutil.Encode(log.Data),
	}
}
