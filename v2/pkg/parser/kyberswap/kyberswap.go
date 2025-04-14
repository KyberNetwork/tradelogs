// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package kyberswap

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// MetaAggregationRouterV2SwapDescriptionV2 is an auto generated low-level Go binding around an user-defined struct.
type MetaAggregationRouterV2SwapDescriptionV2 struct {
	SrcToken        common.Address
	DstToken        common.Address
	SrcReceivers    []common.Address
	SrcAmounts      []*big.Int
	FeeReceivers    []common.Address
	FeeAmounts      []*big.Int
	DstReceiver     common.Address
	Amount          *big.Int
	MinReturnAmount *big.Int
	Flags           *big.Int
	Permit          []byte
}

// MetaAggregationRouterV2SwapExecutionParams is an auto generated low-level Go binding around an user-defined struct.
type MetaAggregationRouterV2SwapExecutionParams struct {
	CallTarget    common.Address
	ApproveTarget common.Address
	TargetData    []byte
	Desc          MetaAggregationRouterV2SwapDescriptionV2
	ClientData    []byte
}

// SwappedMetaData contains all meta data concerning the Swapped contract.
var SwappedMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_WETH\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"clientData\",\"type\":\"bytes\"}],\"name\":\"ClientData\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"Error\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"output\",\"type\":\"address\"}],\"name\":\"Exchange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"recipients\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isBps\",\"type\":\"bool\"}],\"name\":\"Fee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"srcToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"dstToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"dstReceiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"spentAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"}],\"name\":\"Swapped\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"WETH\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isWhitelist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"rescueFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"callTarget\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"approveTarget\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"targetData\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"contractIERC20\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"dstToken\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"srcReceivers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"srcAmounts\",\"type\":\"uint256[]\"},{\"internalType\":\"address[]\",\"name\":\"feeReceivers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"feeAmounts\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"dstReceiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minReturnAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"flags\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"permit\",\"type\":\"bytes\"}],\"internalType\":\"structMetaAggregationRouterV2.SwapDescriptionV2\",\"name\":\"desc\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"clientData\",\"type\":\"bytes\"}],\"internalType\":\"structMetaAggregationRouterV2.SwapExecutionParams\",\"name\":\"execution\",\"type\":\"tuple\"}],\"name\":\"swap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasUsed\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"callTarget\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"approveTarget\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"targetData\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"contractIERC20\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"dstToken\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"srcReceivers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"srcAmounts\",\"type\":\"uint256[]\"},{\"internalType\":\"address[]\",\"name\":\"feeReceivers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"feeAmounts\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"dstReceiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minReturnAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"flags\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"permit\",\"type\":\"bytes\"}],\"internalType\":\"structMetaAggregationRouterV2.SwapDescriptionV2\",\"name\":\"desc\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"clientData\",\"type\":\"bytes\"}],\"internalType\":\"structMetaAggregationRouterV2.SwapExecutionParams\",\"name\":\"execution\",\"type\":\"tuple\"}],\"name\":\"swapGeneric\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasUsed\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIAggregationExecutor\",\"name\":\"caller\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"contractIERC20\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"dstToken\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"srcReceivers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"srcAmounts\",\"type\":\"uint256[]\"},{\"internalType\":\"address[]\",\"name\":\"feeReceivers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"feeAmounts\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"dstReceiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minReturnAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"flags\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"permit\",\"type\":\"bytes\"}],\"internalType\":\"structMetaAggregationRouterV2.SwapDescriptionV2\",\"name\":\"desc\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"executorData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"clientData\",\"type\":\"bytes\"}],\"name\":\"swapSimpleMode\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasUsed\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"addr\",\"type\":\"address[]\"},{\"internalType\":\"bool[]\",\"name\":\"value\",\"type\":\"bool[]\"}],\"name\":\"updateWhitelist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// SwappedABI is the input ABI used to generate the binding from.
// Deprecated: Use SwappedMetaData.ABI instead.
var SwappedABI = SwappedMetaData.ABI

// Swapped is an auto generated Go binding around an Ethereum contract.
type Swapped struct {
	SwappedCaller     // Read-only binding to the contract
	SwappedTransactor // Write-only binding to the contract
	SwappedFilterer   // Log filterer for contract events
}

// SwappedCaller is an auto generated read-only Go binding around an Ethereum contract.
type SwappedCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwappedTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SwappedTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwappedFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SwappedFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwappedSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SwappedSession struct {
	Contract     *Swapped          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SwappedCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SwappedCallerSession struct {
	Contract *SwappedCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// SwappedTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SwappedTransactorSession struct {
	Contract     *SwappedTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// SwappedRaw is an auto generated low-level Go binding around an Ethereum contract.
type SwappedRaw struct {
	Contract *Swapped // Generic contract binding to access the raw methods on
}

// SwappedCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SwappedCallerRaw struct {
	Contract *SwappedCaller // Generic read-only contract binding to access the raw methods on
}

// SwappedTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SwappedTransactorRaw struct {
	Contract *SwappedTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSwapped creates a new instance of Swapped, bound to a specific deployed contract.
func NewSwapped(address common.Address, backend bind.ContractBackend) (*Swapped, error) {
	contract, err := bindSwapped(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Swapped{SwappedCaller: SwappedCaller{contract: contract}, SwappedTransactor: SwappedTransactor{contract: contract}, SwappedFilterer: SwappedFilterer{contract: contract}}, nil
}

// NewSwappedCaller creates a new read-only instance of Swapped, bound to a specific deployed contract.
func NewSwappedCaller(address common.Address, caller bind.ContractCaller) (*SwappedCaller, error) {
	contract, err := bindSwapped(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SwappedCaller{contract: contract}, nil
}

// NewSwappedTransactor creates a new write-only instance of Swapped, bound to a specific deployed contract.
func NewSwappedTransactor(address common.Address, transactor bind.ContractTransactor) (*SwappedTransactor, error) {
	contract, err := bindSwapped(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SwappedTransactor{contract: contract}, nil
}

// NewSwappedFilterer creates a new log filterer instance of Swapped, bound to a specific deployed contract.
func NewSwappedFilterer(address common.Address, filterer bind.ContractFilterer) (*SwappedFilterer, error) {
	contract, err := bindSwapped(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SwappedFilterer{contract: contract}, nil
}

// bindSwapped binds a generic wrapper to an already deployed contract.
func bindSwapped(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SwappedMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Swapped *SwappedRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Swapped.Contract.SwappedCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Swapped *SwappedRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Swapped.Contract.SwappedTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Swapped *SwappedRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Swapped.Contract.SwappedTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Swapped *SwappedCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Swapped.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Swapped *SwappedTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Swapped.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Swapped *SwappedTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Swapped.Contract.contract.Transact(opts, method, params...)
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_Swapped *SwappedCaller) WETH(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Swapped.contract.Call(opts, &out, "WETH")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_Swapped *SwappedSession) WETH() (common.Address, error) {
	return _Swapped.Contract.WETH(&_Swapped.CallOpts)
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_Swapped *SwappedCallerSession) WETH() (common.Address, error) {
	return _Swapped.Contract.WETH(&_Swapped.CallOpts)
}

// IsWhitelist is a free data retrieval call binding the contract method 0xc683630d.
//
// Solidity: function isWhitelist(address ) view returns(bool)
func (_Swapped *SwappedCaller) IsWhitelist(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Swapped.contract.Call(opts, &out, "isWhitelist", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsWhitelist is a free data retrieval call binding the contract method 0xc683630d.
//
// Solidity: function isWhitelist(address ) view returns(bool)
func (_Swapped *SwappedSession) IsWhitelist(arg0 common.Address) (bool, error) {
	return _Swapped.Contract.IsWhitelist(&_Swapped.CallOpts, arg0)
}

// IsWhitelist is a free data retrieval call binding the contract method 0xc683630d.
//
// Solidity: function isWhitelist(address ) view returns(bool)
func (_Swapped *SwappedCallerSession) IsWhitelist(arg0 common.Address) (bool, error) {
	return _Swapped.Contract.IsWhitelist(&_Swapped.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Swapped *SwappedCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Swapped.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Swapped *SwappedSession) Owner() (common.Address, error) {
	return _Swapped.Contract.Owner(&_Swapped.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Swapped *SwappedCallerSession) Owner() (common.Address, error) {
	return _Swapped.Contract.Owner(&_Swapped.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Swapped *SwappedTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Swapped.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Swapped *SwappedSession) RenounceOwnership() (*types.Transaction, error) {
	return _Swapped.Contract.RenounceOwnership(&_Swapped.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Swapped *SwappedTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Swapped.Contract.RenounceOwnership(&_Swapped.TransactOpts)
}

// RescueFunds is a paid mutator transaction binding the contract method 0x78e3214f.
//
// Solidity: function rescueFunds(address token, uint256 amount) returns()
func (_Swapped *SwappedTransactor) RescueFunds(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Swapped.contract.Transact(opts, "rescueFunds", token, amount)
}

// RescueFunds is a paid mutator transaction binding the contract method 0x78e3214f.
//
// Solidity: function rescueFunds(address token, uint256 amount) returns()
func (_Swapped *SwappedSession) RescueFunds(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Swapped.Contract.RescueFunds(&_Swapped.TransactOpts, token, amount)
}

// RescueFunds is a paid mutator transaction binding the contract method 0x78e3214f.
//
// Solidity: function rescueFunds(address token, uint256 amount) returns()
func (_Swapped *SwappedTransactorSession) RescueFunds(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Swapped.Contract.RescueFunds(&_Swapped.TransactOpts, token, amount)
}

// Swap is a paid mutator transaction binding the contract method 0xe21fd0e9.
//
// Solidity: function swap((address,address,bytes,(address,address,address[],uint256[],address[],uint256[],address,uint256,uint256,uint256,bytes),bytes) execution) payable returns(uint256 returnAmount, uint256 gasUsed)
func (_Swapped *SwappedTransactor) Swap(opts *bind.TransactOpts, execution MetaAggregationRouterV2SwapExecutionParams) (*types.Transaction, error) {
	return _Swapped.contract.Transact(opts, "swap", execution)
}

// Swap is a paid mutator transaction binding the contract method 0xe21fd0e9.
//
// Solidity: function swap((address,address,bytes,(address,address,address[],uint256[],address[],uint256[],address,uint256,uint256,uint256,bytes),bytes) execution) payable returns(uint256 returnAmount, uint256 gasUsed)
func (_Swapped *SwappedSession) Swap(execution MetaAggregationRouterV2SwapExecutionParams) (*types.Transaction, error) {
	return _Swapped.Contract.Swap(&_Swapped.TransactOpts, execution)
}

// Swap is a paid mutator transaction binding the contract method 0xe21fd0e9.
//
// Solidity: function swap((address,address,bytes,(address,address,address[],uint256[],address[],uint256[],address,uint256,uint256,uint256,bytes),bytes) execution) payable returns(uint256 returnAmount, uint256 gasUsed)
func (_Swapped *SwappedTransactorSession) Swap(execution MetaAggregationRouterV2SwapExecutionParams) (*types.Transaction, error) {
	return _Swapped.Contract.Swap(&_Swapped.TransactOpts, execution)
}

// SwapGeneric is a paid mutator transaction binding the contract method 0x59e50fed.
//
// Solidity: function swapGeneric((address,address,bytes,(address,address,address[],uint256[],address[],uint256[],address,uint256,uint256,uint256,bytes),bytes) execution) payable returns(uint256 returnAmount, uint256 gasUsed)
func (_Swapped *SwappedTransactor) SwapGeneric(opts *bind.TransactOpts, execution MetaAggregationRouterV2SwapExecutionParams) (*types.Transaction, error) {
	return _Swapped.contract.Transact(opts, "swapGeneric", execution)
}

// SwapGeneric is a paid mutator transaction binding the contract method 0x59e50fed.
//
// Solidity: function swapGeneric((address,address,bytes,(address,address,address[],uint256[],address[],uint256[],address,uint256,uint256,uint256,bytes),bytes) execution) payable returns(uint256 returnAmount, uint256 gasUsed)
func (_Swapped *SwappedSession) SwapGeneric(execution MetaAggregationRouterV2SwapExecutionParams) (*types.Transaction, error) {
	return _Swapped.Contract.SwapGeneric(&_Swapped.TransactOpts, execution)
}

// SwapGeneric is a paid mutator transaction binding the contract method 0x59e50fed.
//
// Solidity: function swapGeneric((address,address,bytes,(address,address,address[],uint256[],address[],uint256[],address,uint256,uint256,uint256,bytes),bytes) execution) payable returns(uint256 returnAmount, uint256 gasUsed)
func (_Swapped *SwappedTransactorSession) SwapGeneric(execution MetaAggregationRouterV2SwapExecutionParams) (*types.Transaction, error) {
	return _Swapped.Contract.SwapGeneric(&_Swapped.TransactOpts, execution)
}

// SwapSimpleMode is a paid mutator transaction binding the contract method 0x8af033fb.
//
// Solidity: function swapSimpleMode(address caller, (address,address,address[],uint256[],address[],uint256[],address,uint256,uint256,uint256,bytes) desc, bytes executorData, bytes clientData) returns(uint256 returnAmount, uint256 gasUsed)
func (_Swapped *SwappedTransactor) SwapSimpleMode(opts *bind.TransactOpts, caller common.Address, desc MetaAggregationRouterV2SwapDescriptionV2, executorData []byte, clientData []byte) (*types.Transaction, error) {
	return _Swapped.contract.Transact(opts, "swapSimpleMode", caller, desc, executorData, clientData)
}

// SwapSimpleMode is a paid mutator transaction binding the contract method 0x8af033fb.
//
// Solidity: function swapSimpleMode(address caller, (address,address,address[],uint256[],address[],uint256[],address,uint256,uint256,uint256,bytes) desc, bytes executorData, bytes clientData) returns(uint256 returnAmount, uint256 gasUsed)
func (_Swapped *SwappedSession) SwapSimpleMode(caller common.Address, desc MetaAggregationRouterV2SwapDescriptionV2, executorData []byte, clientData []byte) (*types.Transaction, error) {
	return _Swapped.Contract.SwapSimpleMode(&_Swapped.TransactOpts, caller, desc, executorData, clientData)
}

// SwapSimpleMode is a paid mutator transaction binding the contract method 0x8af033fb.
//
// Solidity: function swapSimpleMode(address caller, (address,address,address[],uint256[],address[],uint256[],address,uint256,uint256,uint256,bytes) desc, bytes executorData, bytes clientData) returns(uint256 returnAmount, uint256 gasUsed)
func (_Swapped *SwappedTransactorSession) SwapSimpleMode(caller common.Address, desc MetaAggregationRouterV2SwapDescriptionV2, executorData []byte, clientData []byte) (*types.Transaction, error) {
	return _Swapped.Contract.SwapSimpleMode(&_Swapped.TransactOpts, caller, desc, executorData, clientData)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Swapped *SwappedTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Swapped.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Swapped *SwappedSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Swapped.Contract.TransferOwnership(&_Swapped.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Swapped *SwappedTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Swapped.Contract.TransferOwnership(&_Swapped.TransactOpts, newOwner)
}

// UpdateWhitelist is a paid mutator transaction binding the contract method 0x33320de3.
//
// Solidity: function updateWhitelist(address[] addr, bool[] value) returns()
func (_Swapped *SwappedTransactor) UpdateWhitelist(opts *bind.TransactOpts, addr []common.Address, value []bool) (*types.Transaction, error) {
	return _Swapped.contract.Transact(opts, "updateWhitelist", addr, value)
}

// UpdateWhitelist is a paid mutator transaction binding the contract method 0x33320de3.
//
// Solidity: function updateWhitelist(address[] addr, bool[] value) returns()
func (_Swapped *SwappedSession) UpdateWhitelist(addr []common.Address, value []bool) (*types.Transaction, error) {
	return _Swapped.Contract.UpdateWhitelist(&_Swapped.TransactOpts, addr, value)
}

// UpdateWhitelist is a paid mutator transaction binding the contract method 0x33320de3.
//
// Solidity: function updateWhitelist(address[] addr, bool[] value) returns()
func (_Swapped *SwappedTransactorSession) UpdateWhitelist(addr []common.Address, value []bool) (*types.Transaction, error) {
	return _Swapped.Contract.UpdateWhitelist(&_Swapped.TransactOpts, addr, value)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Swapped *SwappedTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Swapped.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Swapped *SwappedSession) Receive() (*types.Transaction, error) {
	return _Swapped.Contract.Receive(&_Swapped.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Swapped *SwappedTransactorSession) Receive() (*types.Transaction, error) {
	return _Swapped.Contract.Receive(&_Swapped.TransactOpts)
}

// SwappedClientDataIterator is returned from FilterClientData and is used to iterate over the raw logs and unpacked data for ClientData events raised by the Swapped contract.
type SwappedClientDataIterator struct {
	Event *SwappedClientData // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SwappedClientDataIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwappedClientData)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SwappedClientData)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SwappedClientDataIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwappedClientDataIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwappedClientData represents a ClientData event raised by the Swapped contract.
type SwappedClientData struct {
	ClientData []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterClientData is a free log retrieval operation binding the contract event 0x095e66fa4dd6a6f7b43fb8444a7bd0edb870508c7abf639bc216efb0bcff9779.
//
// Solidity: event ClientData(bytes clientData)
func (_Swapped *SwappedFilterer) FilterClientData(opts *bind.FilterOpts) (*SwappedClientDataIterator, error) {

	logs, sub, err := _Swapped.contract.FilterLogs(opts, "ClientData")
	if err != nil {
		return nil, err
	}
	return &SwappedClientDataIterator{contract: _Swapped.contract, event: "ClientData", logs: logs, sub: sub}, nil
}

// WatchClientData is a free log subscription operation binding the contract event 0x095e66fa4dd6a6f7b43fb8444a7bd0edb870508c7abf639bc216efb0bcff9779.
//
// Solidity: event ClientData(bytes clientData)
func (_Swapped *SwappedFilterer) WatchClientData(opts *bind.WatchOpts, sink chan<- *SwappedClientData) (event.Subscription, error) {

	logs, sub, err := _Swapped.contract.WatchLogs(opts, "ClientData")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwappedClientData)
				if err := _Swapped.contract.UnpackLog(event, "ClientData", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseClientData is a log parse operation binding the contract event 0x095e66fa4dd6a6f7b43fb8444a7bd0edb870508c7abf639bc216efb0bcff9779.
//
// Solidity: event ClientData(bytes clientData)
func (_Swapped *SwappedFilterer) ParseClientData(log types.Log) (*SwappedClientData, error) {
	event := new(SwappedClientData)
	if err := _Swapped.contract.UnpackLog(event, "ClientData", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwappedErrorIterator is returned from FilterError and is used to iterate over the raw logs and unpacked data for Error events raised by the Swapped contract.
type SwappedErrorIterator struct {
	Event *SwappedError // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SwappedErrorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwappedError)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SwappedError)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SwappedErrorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwappedErrorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwappedError represents a Error event raised by the Swapped contract.
type SwappedError struct {
	Reason string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterError is a free log retrieval operation binding the contract event 0x08c379a0afcc32b1a39302f7cb8073359698411ab5fd6e3edb2c02c0b5fba8aa.
//
// Solidity: event Error(string reason)
func (_Swapped *SwappedFilterer) FilterError(opts *bind.FilterOpts) (*SwappedErrorIterator, error) {

	logs, sub, err := _Swapped.contract.FilterLogs(opts, "Error")
	if err != nil {
		return nil, err
	}
	return &SwappedErrorIterator{contract: _Swapped.contract, event: "Error", logs: logs, sub: sub}, nil
}

// WatchError is a free log subscription operation binding the contract event 0x08c379a0afcc32b1a39302f7cb8073359698411ab5fd6e3edb2c02c0b5fba8aa.
//
// Solidity: event Error(string reason)
func (_Swapped *SwappedFilterer) WatchError(opts *bind.WatchOpts, sink chan<- *SwappedError) (event.Subscription, error) {

	logs, sub, err := _Swapped.contract.WatchLogs(opts, "Error")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwappedError)
				if err := _Swapped.contract.UnpackLog(event, "Error", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseError is a log parse operation binding the contract event 0x08c379a0afcc32b1a39302f7cb8073359698411ab5fd6e3edb2c02c0b5fba8aa.
//
// Solidity: event Error(string reason)
func (_Swapped *SwappedFilterer) ParseError(log types.Log) (*SwappedError, error) {
	event := new(SwappedError)
	if err := _Swapped.contract.UnpackLog(event, "Error", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwappedExchangeIterator is returned from FilterExchange and is used to iterate over the raw logs and unpacked data for Exchange events raised by the Swapped contract.
type SwappedExchangeIterator struct {
	Event *SwappedExchange // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SwappedExchangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwappedExchange)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SwappedExchange)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SwappedExchangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwappedExchangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwappedExchange represents a Exchange event raised by the Swapped contract.
type SwappedExchange struct {
	Pair      common.Address
	AmountOut *big.Int
	Output    common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterExchange is a free log retrieval operation binding the contract event 0xddac40937f35385a34f721af292e5a83fc5b840f722bff57c2fc71adba708c48.
//
// Solidity: event Exchange(address pair, uint256 amountOut, address output)
func (_Swapped *SwappedFilterer) FilterExchange(opts *bind.FilterOpts) (*SwappedExchangeIterator, error) {

	logs, sub, err := _Swapped.contract.FilterLogs(opts, "Exchange")
	if err != nil {
		return nil, err
	}
	return &SwappedExchangeIterator{contract: _Swapped.contract, event: "Exchange", logs: logs, sub: sub}, nil
}

// WatchExchange is a free log subscription operation binding the contract event 0xddac40937f35385a34f721af292e5a83fc5b840f722bff57c2fc71adba708c48.
//
// Solidity: event Exchange(address pair, uint256 amountOut, address output)
func (_Swapped *SwappedFilterer) WatchExchange(opts *bind.WatchOpts, sink chan<- *SwappedExchange) (event.Subscription, error) {

	logs, sub, err := _Swapped.contract.WatchLogs(opts, "Exchange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwappedExchange)
				if err := _Swapped.contract.UnpackLog(event, "Exchange", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseExchange is a log parse operation binding the contract event 0xddac40937f35385a34f721af292e5a83fc5b840f722bff57c2fc71adba708c48.
//
// Solidity: event Exchange(address pair, uint256 amountOut, address output)
func (_Swapped *SwappedFilterer) ParseExchange(log types.Log) (*SwappedExchange, error) {
	event := new(SwappedExchange)
	if err := _Swapped.contract.UnpackLog(event, "Exchange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwappedFeeIterator is returned from FilterFee and is used to iterate over the raw logs and unpacked data for Fee events raised by the Swapped contract.
type SwappedFeeIterator struct {
	Event *SwappedFee // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SwappedFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwappedFee)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SwappedFee)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SwappedFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwappedFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwappedFee represents a Fee event raised by the Swapped contract.
type SwappedFee struct {
	Token       common.Address
	TotalAmount *big.Int
	TotalFee    *big.Int
	Recipients  []common.Address
	Amounts     []*big.Int
	IsBps       bool
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterFee is a free log retrieval operation binding the contract event 0x4c39b7ce5f4f514f45cb6f82b171b8b0b7f2cbf488ad28e4eff451588e2f014b.
//
// Solidity: event Fee(address token, uint256 totalAmount, uint256 totalFee, address[] recipients, uint256[] amounts, bool isBps)
func (_Swapped *SwappedFilterer) FilterFee(opts *bind.FilterOpts) (*SwappedFeeIterator, error) {

	logs, sub, err := _Swapped.contract.FilterLogs(opts, "Fee")
	if err != nil {
		return nil, err
	}
	return &SwappedFeeIterator{contract: _Swapped.contract, event: "Fee", logs: logs, sub: sub}, nil
}

// WatchFee is a free log subscription operation binding the contract event 0x4c39b7ce5f4f514f45cb6f82b171b8b0b7f2cbf488ad28e4eff451588e2f014b.
//
// Solidity: event Fee(address token, uint256 totalAmount, uint256 totalFee, address[] recipients, uint256[] amounts, bool isBps)
func (_Swapped *SwappedFilterer) WatchFee(opts *bind.WatchOpts, sink chan<- *SwappedFee) (event.Subscription, error) {

	logs, sub, err := _Swapped.contract.WatchLogs(opts, "Fee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwappedFee)
				if err := _Swapped.contract.UnpackLog(event, "Fee", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFee is a log parse operation binding the contract event 0x4c39b7ce5f4f514f45cb6f82b171b8b0b7f2cbf488ad28e4eff451588e2f014b.
//
// Solidity: event Fee(address token, uint256 totalAmount, uint256 totalFee, address[] recipients, uint256[] amounts, bool isBps)
func (_Swapped *SwappedFilterer) ParseFee(log types.Log) (*SwappedFee, error) {
	event := new(SwappedFee)
	if err := _Swapped.contract.UnpackLog(event, "Fee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwappedOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Swapped contract.
type SwappedOwnershipTransferredIterator struct {
	Event *SwappedOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SwappedOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwappedOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SwappedOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SwappedOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwappedOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwappedOwnershipTransferred represents a OwnershipTransferred event raised by the Swapped contract.
type SwappedOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Swapped *SwappedFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SwappedOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Swapped.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SwappedOwnershipTransferredIterator{contract: _Swapped.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Swapped *SwappedFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SwappedOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Swapped.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwappedOwnershipTransferred)
				if err := _Swapped.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Swapped *SwappedFilterer) ParseOwnershipTransferred(log types.Log) (*SwappedOwnershipTransferred, error) {
	event := new(SwappedOwnershipTransferred)
	if err := _Swapped.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwappedSwappedIterator is returned from FilterSwapped and is used to iterate over the raw logs and unpacked data for Swapped events raised by the Swapped contract.
type SwappedSwappedIterator struct {
	Event *SwappedSwapped // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SwappedSwappedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwappedSwapped)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SwappedSwapped)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SwappedSwappedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwappedSwappedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwappedSwapped represents a Swapped event raised by the Swapped contract.
type SwappedSwapped struct {
	Sender       common.Address
	SrcToken     common.Address
	DstToken     common.Address
	DstReceiver  common.Address
	SpentAmount  *big.Int
	ReturnAmount *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSwapped is a free log retrieval operation binding the contract event 0xd6d4f5681c246c9f42c203e287975af1601f8df8035a9251f79aab5c8f09e2f8.
//
// Solidity: event Swapped(address sender, address srcToken, address dstToken, address dstReceiver, uint256 spentAmount, uint256 returnAmount)
func (_Swapped *SwappedFilterer) FilterSwapped(opts *bind.FilterOpts) (*SwappedSwappedIterator, error) {

	logs, sub, err := _Swapped.contract.FilterLogs(opts, "Swapped")
	if err != nil {
		return nil, err
	}
	return &SwappedSwappedIterator{contract: _Swapped.contract, event: "Swapped", logs: logs, sub: sub}, nil
}

// WatchSwapped is a free log subscription operation binding the contract event 0xd6d4f5681c246c9f42c203e287975af1601f8df8035a9251f79aab5c8f09e2f8.
//
// Solidity: event Swapped(address sender, address srcToken, address dstToken, address dstReceiver, uint256 spentAmount, uint256 returnAmount)
func (_Swapped *SwappedFilterer) WatchSwapped(opts *bind.WatchOpts, sink chan<- *SwappedSwapped) (event.Subscription, error) {

	logs, sub, err := _Swapped.contract.WatchLogs(opts, "Swapped")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwappedSwapped)
				if err := _Swapped.contract.UnpackLog(event, "Swapped", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSwapped is a log parse operation binding the contract event 0xd6d4f5681c246c9f42c203e287975af1601f8df8035a9251f79aab5c8f09e2f8.
//
// Solidity: event Swapped(address sender, address srcToken, address dstToken, address dstReceiver, uint256 spentAmount, uint256 returnAmount)
func (_Swapped *SwappedFilterer) ParseSwapped(log types.Log) (*SwappedSwapped, error) {
	event := new(SwappedSwapped)
	if err := _Swapped.contract.UnpackLog(event, "Swapped", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
