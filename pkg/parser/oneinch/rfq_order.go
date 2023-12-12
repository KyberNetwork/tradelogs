package oneinch

import (
	"encoding/json"
	"errors"
	"github.com/KyberNetwork/tradelogs/internal/types"
	"github.com/KyberNetwork/tradelogs/pkg/storage"
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
		tradeLog.MakerTokenAmount = rfqOrder.MakingAmount.String()
		tradeLog.TakerTokenAmount = rfqOrder.TakingAmount.String()
	}

	return tradeLog, nil

}
