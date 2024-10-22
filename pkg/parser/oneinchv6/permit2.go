package oneinchv6

import (
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

const (
	argsExtensionLengthOffset  = 224
	argsExtensionLengthMask    = 0xffffff
	permit2WitnessProxyAddress = "0x8b543dff08ed4ba13ee96f533638ef54591aee71"
)

type (
	dynamicField int
	Permit       struct {
		Permitted struct {
			Token  common.Address `json:"token"`
			Amount *big.Int       `json:"amount"`
		} `json:"permitted"`
		Nonce    *big.Int `json:"nonce"`
		Deadline *big.Int `json:"deadline"`
	}
)

const (
	MakerAssetSuffix dynamicField = iota
	TakerAssetSuffix
	MakingAmountData
	TakingAmountData
	Predicate
	MakerPermit
	PreInteractionData
	PostInteractionData
	CustomData
)

func parseArgs(takerTraits *big.Int, args []byte) (extension []byte) {
	// (TakerTraits.unwrap(takerTraits) & _ARGS_HAS_TARGET) != 0; _ARGS_HAS_TARGET = 1<<251
	if new(big.Int).And(takerTraits, new(big.Int).Lsh(big.NewInt(1), 251)).Cmp(big.NewInt(0)) != 0 {
		args = args[20:]
	}

	// (TakerTraits.unwrap(takerTraits) >> _ARGS_EXTENSION_LENGTH_OFFSET) & _ARGS_EXTENSION_LENGTH_MASK;
	extensionLength := new(big.Int).And(new(big.Int).Rsh(takerTraits, argsExtensionLengthOffset), big.NewInt(argsExtensionLengthMask)).Uint64()
	if extensionLength > 0 && extensionLength <= uint64(len(args)) {
		extension = args[:extensionLength]
	}
	return
}

func getMakerAssetSuffix(extension []byte, field dynamicField) []byte {
	if len(extension) < 32 {
		return []byte{}
	}

	// Simulate loading offsets and extracting concatenated data
	offsets := new(big.Int).SetBytes(extension[:32])
	concat := extension[32:]

	bitShift := uint(field * 32)
	mask := big.NewInt(0xFFFFFFFF)

	// Extract the begin value by shifting and masking
	// let begin := and(0xffffffff, shr(bitShift, shl(32, offsets)))
	begin := new(big.Int).Rsh(new(big.Int).Lsh(offsets, 32), bitShift)
	begin.And(begin, mask)

	// Extract the end value by shifting and masking
	//let end := and(0xffffffff, shr(bitShift, offsets))
	end := new(big.Int).Rsh(offsets, bitShift)
	end.And(end, mask)

	beginInt := begin.Uint64()
	endInt := end.Uint64()

	if endInt > uint64(len(concat)) {
		return []byte{}
	}

	return concat[beginInt:endInt]
}

func decodeMakerAssetSuffix(takerTraits *big.Int, args []byte) (common.Address, error) {
	// get the extension from the args
	extension := parseArgs(takerTraits, args)

	// get the suffix from extension
	suffix := getMakerAssetSuffix(extension, MakerAssetSuffix)
	if len(suffix) == 0 {
		return common.Address{}, fmt.Errorf("suffix is empty")
	}

	permitTransferFromType, err := abi.NewType("tuple", "", []abi.ArgumentMarshaling{
		{Name: "permitted", Type: "tuple", Components: []abi.ArgumentMarshaling{
			{Name: "token", Type: "address"},
			{Name: "amount", Type: "uint256"}}},
		{Name: "nonce", Type: "uint256"},
		{Name: "deadline", Type: "uint256"},
	})
	if err != nil {
		return common.Address{}, err
	}
	permitArgument := abi.Arguments{
		{Name: "permit", Type: permitTransferFromType},
	}

	// Decode the input data (suffix) into the PermitTransferFrom struct
	result, err := permitArgument.Unpack(suffix)
	if err != nil {
		return common.Address{}, err
	}
	marshal, err := json.Marshal(result)
	if err != nil {
		return common.Address{}, err
	}

	var p []Permit
	err = json.Unmarshal(marshal, &p)
	if err != nil {
		return common.Address{}, err
	}
	if len(p) == 0 {
		return common.Address{}, nil
	}
	return p[0].Permitted.Token, err
}
