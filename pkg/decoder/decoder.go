package decoder

import (
	"encoding/hex"
	"fmt"
	tradingTypes "github.com/KyberNetwork/tradinglib/pkg/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

type Bytes4 [4]byte

// convert bytes4 to []byte
func (b Bytes4) Bytes() []byte {
	return b[:]
}

func GetBytes4(data []byte) (Bytes4, error) {
	if len(data) < 4 {
		return Bytes4{}, fmt.Errorf("need atleast 4 bytes, got %d", len(data))
	}
	return Bytes4{data[0], data[1], data[2], data[3]}, nil
}

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
	return decode(method, bytes)
}

func decode(method *abi.Method, data []byte) (*tradingTypes.ContractCall, error) {
	inputs, err := method.Inputs.Unpack(data)
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
			Type:  arg.Type.GetType().String(),
		}
		contractCall.Params = append(contractCall.Params, param)
	}

	return contractCall, nil
}

func DecodeCustomAbi(ABI *abi.ABI, methodId Bytes4, rawData []byte) (*tradingTypes.ContractCall, error) {
	if ABI == nil {
		return nil, fmt.Errorf("missing abi")
	}
	method, err := ABI.MethodById(methodId.Bytes())
	if err != nil {
		return nil, err
	}
	return decode(method, rawData)
}
