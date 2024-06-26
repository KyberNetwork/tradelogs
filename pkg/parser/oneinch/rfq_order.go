package oneinch

import (
	"encoding/json"
	"errors"
	"github.com/KyberNetwork/tradelogs/pkg/decoder"
	"github.com/KyberNetwork/tradelogs/pkg/storage"
	tradingTypes "github.com/KyberNetwork/tradinglib/pkg/types"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
)

const (
	paramName = "order"
)

type OrderRFQ struct {
	Info          float64 `json:"info"`
	MakerAsset    string  `json:"makerAsset"`
	TakerAsset    string  `json:"takerAsset"`
	Maker         string  `json:"maker"`
	AllowedSender string  `json:"allowedSender"`
	MakingAmount  float64 `json:"makingAmount"`
	TakingAmount  float64 `json:"takingAmount"`
}

func (o *OrderRFQ) GetExpiry() uint64 {
	infoFloat := big.NewFloat(o.Info)
	info := new(big.Int)
	info, _ = infoFloat.Int(info)
	expiry := info.Rsh(info, 64)
	return expiry.Uint64()
}

func ToTradeLog(tradeLog storage.TradeLog, contractCall *tradingTypes.ContractCall) (storage.TradeLog, error) {
	if contractCall == nil {
		return tradeLog, errors.New("contract call is empty")
	}
	for _, param := range contractCall.Params {
		if param.Name != paramName {
			continue
		}
		var rfqOrder OrderRFQLibOrderRFQ
		bytes, err := json.Marshal(param.Value)
		if err != nil {
			return tradeLog, err
		}

		if err := json.Unmarshal(bytes, &rfqOrder); err != nil {
			return tradeLog, err
		}

		tradeLog.MakerToken = rfqOrder.MakerAsset.String()
		tradeLog.TakerToken = rfqOrder.TakerAsset.String()
		tradeLog.Maker = rfqOrder.Maker.String()
	}

	return tradeLog, nil
}

func (p *Parser) getRFQOrderParams(callFrame *tradingTypes.CallFrame) (*OrderRFQ, error) {
	var (
		err error
	)
	contractCall := callFrame.ContractCall
	if contractCall == nil {
		contractCall, err = decoder.Decode(p.abi, hexutil.Encode(callFrame.Input))
		if err != nil {
			return nil, err
		}
		if contractCall == nil {
			return nil, errors.New("missing contract_call")
		}
	}
	for _, param := range contractCall.Params {
		if param.Name != paramName {
			continue
		}

		var rfqOrder OrderRFQ
		bytes, err := json.Marshal(param.Value)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(bytes, &rfqOrder); err != nil {
			return nil, err
		}

		return &rfqOrder, nil
	}
	return nil, nil
}
