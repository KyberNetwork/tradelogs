package decoder

import (
	"encoding/hex"
	"fmt"
	tradingTypes "github.com/KyberNetwork/tradinglib/pkg/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

func Decode(ABI *abi.ABI, input string) (*tradingTypes.ContractCall, error) {
	if ABI == nil {
		return nil, fmt.Errorf("missing abi")
	}
	inputBytes := common.FromHex(input)
	method, err := ABI.MethodById(inputBytes)
	if err != nil {
		return nil, err
	}

	bytes, err := hex.DecodeString(input[10:])
	if err != nil {
		return nil, err
	}

	inputs, err := method.Inputs.Unpack(bytes)
	if err != nil {
		return nil, err
	}

	nonIndexedArgs := method.Inputs.NonIndexed()

	contractCall := &tradingTypes.ContractCall{
		Name: method.Name,
	}

	for i, input := range inputs {
		arg := nonIndexedArgs[i]
		param := tradingTypes.ContractCallParam{
			Name:  arg.Name,
			Value: input,
			Type:  arg.Type.String(),
		}
		contractCall.Params = append(contractCall.Params, param)
	}

	return contractCall, nil
}
