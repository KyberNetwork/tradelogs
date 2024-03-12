package decoder

import (
	"encoding/hex"
	"fmt"
	"github.com/KyberNetwork/tradelogs/internal/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

func Decode(ABI *abi.ABI, input string) (*types.ContractCall, error) {
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

	contractCall := &types.ContractCall{
		Name: method.Name,
	}

	for i, input := range inputs {
		arg := nonIndexedArgs[i]
		param := types.ContractCallParam{
			Name:  arg.Name,
			Value: input,
			Type:  arg.Type.String(),
		}
		contractCall.Params = append(contractCall.Params, param)
	}

	return contractCall, nil
}
