package oneinchv6

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"strings"

	"github.com/KyberNetwork/tradelogs/pkg/storage"
	tradingTypes "github.com/KyberNetwork/tradinglib/pkg/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
)

const (
	paramName            = "order"
	takerTraitsParamName = "takerTraits"
	argsParamName        = "args"
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
		makerTraitsOption, err := DecodeMarkerTraits(math.PaddedBigBytes(rfqOrder.MakerTraits, 32))
		if err != nil {
			return tradeLog, err
		}
		tradeLog.Expiry = uint64(makerTraitsOption.Expiration)

		// if maker is Permit2WitnessProxy
		if strings.EqualFold(tradeLog.MakerToken, permit2WitnessProxyAddress) {
			tradeLog.MakerToken, err = getMakerAsset(contractCall)
			if err != nil {
				return tradeLog, err
			}
		}
	}

	return tradeLog, nil

}

func getMakerAsset(contractCall *tradingTypes.ContractCall) (string, error) {
	var (
		takerTraits *big.Int
		args        []byte
	)
	for _, param := range contractCall.Params {
		switch param.Name {
		case takerTraitsParamName:
			takerTraits = param.Value.(*big.Int)
		case argsParamName:
			args = param.Value.([]byte)
		}
	}
	if takerTraits == nil || len(args) == 0 {
		return "", fmt.Errorf("taker traits or args is empty")
	}
	token, err := decodeMakerAssetSuffix(takerTraits, args)
	if err != nil {
		return "", fmt.Errorf("fail to decode marker asset suffix: %w", err)
	}
	return token.String(), nil
}
