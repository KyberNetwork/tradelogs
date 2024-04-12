package oneinchv6

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
		var rfqOrder IOrderMixinOrder
		bytes, err := json.Marshal(param.Value)
		if err != nil {
			return tradeLog, err
		}

		if err := json.Unmarshal(bytes, &rfqOrder); err != nil {
			return tradeLog, err
		}

		tradeLog.MakerToken = "0x" + rfqOrder.MakerAsset.Text(16)
		tradeLog.TakerToken = "0x" + rfqOrder.TakerAsset.Text(16)
		tradeLog.Maker = "0x" + rfqOrder.Maker.Text(16)
		tradeLog.MakerTraits = rfqOrder.MakerTraits.String()
	}

	return tradeLog, nil

}
