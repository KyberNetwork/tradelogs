package zxrfqv3

import (
	"fmt"
	"sync"

	"github.com/KyberNetwork/tradelogs/pkg/parser/zxrfqv3/dev"
	"github.com/KyberNetwork/tradelogs/pkg/parser/zxrfqv3/gasless"
	newgasless "github.com/KyberNetwork/tradelogs/pkg/parser/zxrfqv3/new_gasless"
	"github.com/KyberNetwork/tradelogs/pkg/parser/zxrfqv3/swap"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

type ContractType int

const (
	DevContract ContractType = iota + 1
	SwapContract
	GaslessContract
	NewGaslessContract
)

type ContractABI struct {
	Address      common.Address
	ContractType ContractType
}

type ContractABIs struct {
	m            sync.RWMutex
	mAddressABIs map[common.Address]*abi.ABI
}

func NewABI(contractABI ContractABI) (*abi.ABI, error) {
	if isZeroAddress(contractABI.Address) {
		return nil, fmt.Errorf("ContractType: 1, error %w", ErrContractAddressIsZeroAddress)
	}
	switch contractABI.ContractType {
	case DevContract:
		ab, err := dev.DevMetaData.GetAbi()
		if err != nil {
			return nil, fmt.Errorf("get dev contract abi error: %w", err)
		}
		return ab, nil
	case SwapContract:
		ab, err := swap.SwapMetaData.GetAbi()
		if err != nil {
			return nil, fmt.Errorf("get swap contract abi error: %w", err)
		}
		return ab, nil
	case GaslessContract:
		ab, err := gasless.GaslessMetaData.GetAbi()
		if err != nil {
			return nil, fmt.Errorf("get gasless contract abi error: %w", err)
		}
		return ab, nil
	case NewGaslessContract:
		ab, err := newgasless.NewGaslessMetaData.GetAbi()
		if err != nil {
			return nil, fmt.Errorf("get new gasless contract abi error: %w", err)
		}
		return ab, nil
	default:
		return nil, fmt.Errorf("current type %v, only support dev, swap, gassless type", contractABI.ContractType)
	}
}

func (contractABIs *ContractABIs) getABI(address common.Address) (*abi.ABI, bool) {
	contractABIs.m.Lock()
	defer contractABIs.m.Unlock()
	contractABI, ok := contractABIs.mAddressABIs[address]
	return contractABI, ok
}

func (contractABIs *ContractABIs) addContractABI(address common.Address, newABI *abi.ABI) {
	contractABIs.m.Lock()
	defer contractABIs.m.Unlock()
	_, ok := contractABIs.mAddressABIs[address]
	if !ok {
		contractABIs.mAddressABIs[address] = newABI
	}
}

func (contractABIs *ContractABIs) containAddress(address common.Address) bool {
	contractABIs.m.Lock()
	defer contractABIs.m.Unlock()
	_, ok := contractABIs.mAddressABIs[address]
	return ok
}

func (contractABIs *ContractABIs) getAddresses() []string {
	contractABIs.m.Lock()
	defer contractABIs.m.Unlock()
	res := make([]string, 0, len(contractABIs.mAddressABIs))
	for addr := range contractABIs.mAddressABIs {
		res = append(res, addr.String())
	}
	return res
}
