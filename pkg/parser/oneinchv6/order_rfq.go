package oneinchv6

import (
	"encoding/json"
	"errors"
	"github.com/KyberNetwork/tradelogs/pkg/storage"
	tradingTypes "github.com/KyberNetwork/tradinglib/pkg/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
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
		tradeLog.Taker = common.BigToAddress(rfqOrder.Receiver).String()
		tradeLog.MakerTraits = rfqOrder.MakerTraits.String()
		makerTraitsOption, err := DecodeMarkerTraits(math.PaddedBigBytes(rfqOrder.MakerTraits, 32))
		if err != nil {
			return tradeLog, err
		}
		tradeLog.Expiry = uint64(makerTraitsOption.Expiration)
	}

	return tradeLog, nil

}
