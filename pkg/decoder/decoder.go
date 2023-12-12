package decoder

import (
	"encoding/hex"
	"github.com/KyberNetwork/tradelogs/internal/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"strings"
)

func Decode(abiStr string, input string) (*types.ContractCall, error) {
	ABI, err := abi.JSON(strings.NewReader(abiStr))
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
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
