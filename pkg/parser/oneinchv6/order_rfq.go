package oneinchv6

import (
	"encoding/json"
	"errors"

	tradingTypes "github.com/KyberNetwork/tradinglib/pkg/types"
	"github.com/ethereum/go-ethereum/common"

	"github.com/KyberNetwork/tradelogs/pkg/storage"
)

const (
	paramName = "order"
)

func ToTradeLog(tradeLog storage.TradeLog, contractCall *tradingTypes.ContractCall) (storage.TradeLog, error) {
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

		tradeLog.MakerToken = common.BigToAddress(rfqOrder.MakerAsset).String()
		tradeLog.TakerToken = common.BigToAddress(rfqOrder.TakerAsset).String()
		tradeLog.Maker = common.BigToAddress(rfqOrder.Maker).String()
		tradeLog.MakerTraits = rfqOrder.MakerTraits.String()
	}

	return tradeLog, nil

}
