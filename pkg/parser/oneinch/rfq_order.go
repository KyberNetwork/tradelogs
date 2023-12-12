package oneinch

import (
	"encoding/json"
	"errors"
	"github.com/KyberNetwork/tradelogs/internal/types"
	"github.com/KyberNetwork/tradelogs/pkg/storage"
	"math/big"
)

const (
	paramName = "order"
)

func ToTradeLog(tradeLog storage.TradeLog, contractCall *types.ContractCall) (storage.TradeLog, error) {
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

		// takerTokenAmount = (makingAmount * orderTakerAmount + orderMakerAmount - 1)/orderMakerAmount
		orderMakerAmount := rfqOrder.MakingAmount
		orderTakerAmount := rfqOrder.TakingAmount
		makingAmount, ok := new(big.Int).SetString(tradeLog.MakerTokenAmount, 10)
		if !ok {
			return tradeLog, errors.New("failed to convert maker_token_amount")
		}

		takerTokenAmount := big.NewInt(0).Mul(makingAmount, orderTakerAmount)
		takerTokenAmount = takerTokenAmount.Add(takerTokenAmount, orderMakerAmount)
		takerTokenAmount = takerTokenAmount.Sub(takerTokenAmount, big.NewInt(1))
		takerTokenAmount = takerTokenAmount.Div(takerTokenAmount, orderMakerAmount)

		tradeLog.TakerTokenAmount = takerTokenAmount.String()
	}

	return tradeLog, nil

}
