package zxrfqv3

import (
	"encoding/json"
	"fmt"
	"github.com/KyberNetwork/tradelogs/v2/pkg/abitypes"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type rfqArguments struct {
	fillRfqOrderSelfFundedArguments abi.Arguments
	fillRfqOrderVIPArguments        abi.Arguments
}

type InputParamOfFillRfqOrderSelfFunded struct {
	Recipient          common.Address     `json:"recipient"`
	PermitTransferFrom PermitTransferFrom `json:"permit"`
	Maker              common.Address     `json:"maker"`
	MakerSig           []byte             `json:"makerSig"`
	TakerToken         common.Address     `json:"takerToken"`
	MaxTakerAmount     *big.Int           `json:"maxTakerAmount"`
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
	MakerPermit    PermitTransferFrom `json:"makerPermit"`
	Maker          common.Address     `json:"maker"`
	MakerSig       []byte             `json:"makerSig"`
	TakerPermit    PermitTransferFrom `json:"takerPermit"`
	MaxTakerAmount *big.Int           `json:"maxTakerAmount"`
}

func newRfqArguments() (rfqArguments, error) {
	permitTransferFromType, err := abi.NewType("tuple", "", []abi.ArgumentMarshaling{
		{Name: "permitted", Type: "tuple", Components: []abi.ArgumentMarshaling{
			{Name: "token", Type: "address"},
			{Name: "amount", Type: "uint256"}}},
		{Name: "nonce", Type: "uint256"},
		{Name: "deadline", Type: "uint256"},
	})

	if err != nil {
		return rfqArguments{}, err
	}

	fillRfqOrderSelfFundedArguments := abi.Arguments{
		{Name: "recipient", Type: abitypes.Address},
		{Name: "permit", Type: permitTransferFromType},
		{Name: "maker", Type: abitypes.Address},
		{Name: "makerSig", Type: abitypes.Bytes},
		{Name: "takerToken", Type: abitypes.Address},
		{Name: "maxTakerAmount", Type: abitypes.Uint256},
	}

	fillRfqOrderVIPArguments := abi.Arguments{
		{Name: "recipient", Type: abitypes.Address},
		{Name: "makerPermit", Type: permitTransferFromType},
		{Name: "maker", Type: abitypes.Address},
		{Name: "makerSig", Type: abitypes.Bytes},
		{Name: "takerPermit", Type: permitTransferFromType},
		{Name: "maxTakerAmount", Type: abitypes.Uint256},
	}

	return rfqArguments{
		fillRfqOrderSelfFundedArguments: fillRfqOrderSelfFundedArguments,
		fillRfqOrderVIPArguments:        fillRfqOrderVIPArguments,
	}, nil
}

func (arg rfqArguments) UnpackInputParamsOfFillRfqOrderSelfFunded(bytes []byte) (InputParamOfFillRfqOrderSelfFunded, error) {
	mNameAndValue := make(map[string]interface{})
	err := arg.fillRfqOrderSelfFundedArguments.UnpackIntoMap(mNameAndValue, bytes)
	if err != nil {
		return InputParamOfFillRfqOrderSelfFunded{}, fmt.Errorf("failed to unpack into map: %v", err)
	}
	data, err := json.Marshal(mNameAndValue)
	if err != nil {
		return InputParamOfFillRfqOrderSelfFunded{}, err
	}
	var input InputParamOfFillRfqOrderSelfFunded
	err = json.Unmarshal(data, &input)
	return input, err
}

func (arg rfqArguments) UnpackInputParamsOfFillRfqOrderVIP(bytes []byte) (InputParamOfFillRfqOrderVIP, error) {
	mNameAndValue := make(map[string]interface{})
	err := arg.fillRfqOrderVIPArguments.UnpackIntoMap(mNameAndValue, bytes)
	if err != nil {
		return InputParamOfFillRfqOrderVIP{}, fmt.Errorf("failed to unpack into map: %v", err)
	}
	data, err := json.Marshal(mNameAndValue)
	if err != nil {
		return InputParamOfFillRfqOrderVIP{}, err
	}
	var input InputParamOfFillRfqOrderVIP
	err = json.Unmarshal(data, &input)
	return input, err
}
