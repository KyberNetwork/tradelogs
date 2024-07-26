package zxrfqv3

import (
	"encoding/json"
	"fmt"
	"github.com/KyberNetwork/tradelogs/pkg/decoder"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type InputParamOfFillRfqOrderSelfFunded struct {
	Recipient      common.Address     `json:"recipient"`
	Permit         PermitTransferFrom `json:"permit"`
	Maker          common.Address     `json:"maker"`
	MakerSig       []byte             `json:"makerSig"`
	TakerToken     common.Address     `json:"takerToken"`
	MaxTakerAmount *big.Int           `json:"maxTakerAmount"`
}

type PermitTransferFrom struct {
	Permitted TokenPermissions
	Nonce     *big.Int
	Deadline  *big.Int
}

type TokenPermissions struct {
	Token  common.Address
	Amount *big.Int
}

type InputParamOfFillRfqOrderVIP struct {
	Recipient      common.Address     `json:"recipient"`
	MakerPermit    PermitTransferFrom `json:"permit"`
	Maker          common.Address     `json:"maker"`
	MakerSig       []byte             `json:"makerSig"`
	TakerPermit    PermitTransferFrom `json:"takerPermit"`
	MaxTakerAmount *big.Int           `json:"maxTakerAmount"`
}

func GetInputParamsOfFillRfqOrderSelfFunded(customABI *abi.ABI, actionName decoder.Bytes4, data []byte) (InputParamOfFillRfqOrderSelfFunded, error) {
	contractCall, err := decoder.DecodeCustomABI(customABI, actionName, data)
	if err != nil {
		return InputParamOfFillRfqOrderSelfFunded{}, err
	}
	var input InputParamOfFillRfqOrderSelfFunded
	var ok bool
	inputParam := contractCall.Params
	if len(inputParam) != 6 {
		return InputParamOfFillRfqOrderSelfFunded{}, fmt.Errorf("invalid number of input params, expect 6 but got %d", len(inputParam))
	}
	input.Recipient, ok = inputParam[0].Value.(common.Address)
	if !ok {
		return InputParamOfFillRfqOrderSelfFunded{}, fmt.Errorf("failed to convert reciptent to common.Address")
	}

	err = unpackValue(inputParam[1].Value, &input.Permit)
	if err != nil {
		return InputParamOfFillRfqOrderSelfFunded{}, fmt.Errorf("failed to convert permit to PermitTransferFrom, %w", err)
	}

	input.Maker, ok = inputParam[2].Value.(common.Address)
	if !ok {
		return InputParamOfFillRfqOrderSelfFunded{}, fmt.Errorf("failed to convert maker to common.Address")
	}

	input.MakerSig, ok = inputParam[3].Value.([]byte)
	if !ok {
		return InputParamOfFillRfqOrderSelfFunded{}, fmt.Errorf("failed to convert makerSig to []byte")
	}

	input.TakerToken, ok = inputParam[4].Value.(common.Address)
	if !ok {
		return InputParamOfFillRfqOrderSelfFunded{}, fmt.Errorf("failed to convert takerToken to common.Address")
	}

	input.MaxTakerAmount, ok = inputParam[5].Value.(*big.Int)
	if !ok {
		return InputParamOfFillRfqOrderSelfFunded{}, fmt.Errorf("failed to convert maxTakerAmount to *big.Int")
	}

	return input, nil
}

func GetInputParamsOfFillRfqOrderVIP(customABI *abi.ABI, actionName decoder.Bytes4, data []byte) (InputParamOfFillRfqOrderVIP, error) {
	contractCall, err := decoder.DecodeCustomABI(customABI, actionName, data)
	if err != nil {
		return InputParamOfFillRfqOrderVIP{}, err
	}
	var input InputParamOfFillRfqOrderVIP
	var ok bool
	inputParam := contractCall.Params
	if len(inputParam) != 6 {
		return InputParamOfFillRfqOrderVIP{}, fmt.Errorf("invalid number of input params, expect 6 but got %d", len(inputParam))
	}
	input.Recipient, ok = inputParam[0].Value.(common.Address)
	if !ok {
		return InputParamOfFillRfqOrderVIP{}, fmt.Errorf("failed to convert reciptent to common.Address")
	}

	err = unpackValue(inputParam[1].Value, &input.MakerPermit)
	if err != nil {
		return InputParamOfFillRfqOrderVIP{}, fmt.Errorf("failed to convert makerPermit to PermitTransferFrom, %w", err)
	}

	input.Maker, ok = inputParam[2].Value.(common.Address)
	if !ok {
		return InputParamOfFillRfqOrderVIP{}, fmt.Errorf("failed to convert maker to common.Address")
	}

	input.MakerSig, ok = inputParam[3].Value.([]byte)
	if !ok {
		return InputParamOfFillRfqOrderVIP{}, fmt.Errorf("failed to convert makerSig to []byte")
	}

	err = unpackValue(inputParam[4].Value, &input.TakerPermit)
	if err != nil {
		return InputParamOfFillRfqOrderVIP{}, fmt.Errorf("failed to convert takerPermit to PermitTransferFrom, %w", err)
	}

	input.MaxTakerAmount, ok = inputParam[5].Value.(*big.Int)
	if !ok {
		return InputParamOfFillRfqOrderVIP{}, fmt.Errorf("failed to convert maxTakerAmount to *big.Int")
	}

	return input, nil
}

func unpackValue(src, dest interface{}) error {
	data, err := json.Marshal(src)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, dest)
}
