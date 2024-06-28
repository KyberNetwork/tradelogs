package oneinch

import (
	"encoding/json"
	"errors"
	"github.com/KyberNetwork/tradelogs/pkg/storage"
	tradingTypes "github.com/KyberNetwork/tradinglib/pkg/types"
)

const (
	paramName = "order"
)

func GetExpiry(rfqOrder OrderRFQLibOrderRFQ) uint64 {
	if rfqOrder.Info == nil {
		return 0
	}
	expiry := rfqOrder.Info.Rsh(rfqOrder.Info, 64)
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
		tradeLog.Expiry = GetExpiry(rfqOrder)
	}

	return tradeLog, nil
}
