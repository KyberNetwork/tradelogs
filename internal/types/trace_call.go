package types

import (
	"github.com/ethereum/go-ethereum/common"
	ethereumTypes "github.com/ethereum/go-ethereum/core/types"
)

type TraceCallResponse struct {
	Jsonrpc string    `json:"jsonrpc"`
	ID      int       `json:"id"`
	Result  CallFrame `json:"result"`
}

type CallLog struct {
	Index   uint           `json:"index"`
	Address common.Address `json:"address"`
	Topics  []common.Hash  `json:"topics"`
	Data    string         `json:"data"`
}

func (l CallLog) ToEthereumLog() ethereumTypes.Log {
	return ethereumTypes.Log{
		Index:   l.Index,
		Address: l.Address,
		Topics:  l.Topics,
		Data:    common.Hex2Bytes(l.Data),
	}
}

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
