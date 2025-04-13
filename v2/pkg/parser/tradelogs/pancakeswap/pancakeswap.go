// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package pancakeswap

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

// SignedOrder is an auto generated low-level Go binding around an user-defined struct.
type SignedOrder struct {
	Order []byte
	Sig   []byte
}

// PancakeswapMetaData contains all meta data concerning the Pancakeswap contract.
var PancakeswapMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIPermit2\",\"name\":\"_permit2\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_protocolFeeOwner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"DeadlineBeforeEndTime\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DeadlinePassed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"duplicateToken\",\"type\":\"address\"}],\"name\":\"DuplicateFeeOutput\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EndTimeBeforeStartTime\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"FeeTooLarge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IncorrectAmounts\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InputAndOutputDecay\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientEth\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"}],\"name\":\"InvalidFeeToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidReactor\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NativeTransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoExclusiveOverride\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OrderEndTimeBeforeStartTime\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"filler\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"swapper\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"Fill\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldFeeController\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newFeeController\",\"type\":\"address\"}],\"name\":\"ProtocolFeeControllerSet\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"order\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"internalType\":\"structSignedOrder\",\"name\":\"order\",\"type\":\"tuple\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"order\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"internalType\":\"structSignedOrder[]\",\"name\":\"orders\",\"type\":\"tuple[]\"}],\"name\":\"executeBatch\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"order\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"internalType\":\"structSignedOrder[]\",\"name\":\"orders\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"callbackData\",\"type\":\"bytes\"}],\"name\":\"executeBatchWithCallback\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"order\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"internalType\":\"structSignedOrder\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"callbackData\",\"type\":\"bytes\"}],\"name\":\"executeWithCallback\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeController\",\"outputs\":[{\"internalType\":\"contractIProtocolFeeController\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"permit2\",\"outputs\":[{\"internalType\":\"contractIPermit2\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newFeeController\",\"type\":\"address\"}],\"name\":\"setProtocolFeeController\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// PancakeswapABI is the input ABI used to generate the binding from.
// Deprecated: Use PancakeswapMetaData.ABI instead.
var PancakeswapABI = PancakeswapMetaData.ABI

// Pancakeswap is an auto generated Go binding around an Ethereum contract.
type Pancakeswap struct {
	PancakeswapCaller     // Read-only binding to the contract
	PancakeswapTransactor // Write-only binding to the contract
	PancakeswapFilterer   // Log filterer for contract events
}

// PancakeswapCaller is an auto generated read-only Go binding around an Ethereum contract.
type PancakeswapCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PancakeswapTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PancakeswapTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PancakeswapFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PancakeswapFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PancakeswapSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PancakeswapSession struct {
	Contract     *Pancakeswap      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PancakeswapCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PancakeswapCallerSession struct {
	Contract *PancakeswapCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// PancakeswapTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PancakeswapTransactorSession struct {
	Contract     *PancakeswapTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// PancakeswapRaw is an auto generated low-level Go binding around an Ethereum contract.
type PancakeswapRaw struct {
	Contract *Pancakeswap // Generic contract binding to access the raw methods on
}

// PancakeswapCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PancakeswapCallerRaw struct {
	Contract *PancakeswapCaller // Generic read-only contract binding to access the raw methods on
}

// PancakeswapTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PancakeswapTransactorRaw struct {
	Contract *PancakeswapTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPancakeswap creates a new instance of Pancakeswap, bound to a specific deployed contract.
func NewPancakeswap(address common.Address, backend bind.ContractBackend) (*Pancakeswap, error) {
	contract, err := bindPancakeswap(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Pancakeswap{PancakeswapCaller: PancakeswapCaller{contract: contract}, PancakeswapTransactor: PancakeswapTransactor{contract: contract}, PancakeswapFilterer: PancakeswapFilterer{contract: contract}}, nil
}

// NewPancakeswapCaller creates a new read-only instance of Pancakeswap, bound to a specific deployed contract.
func NewPancakeswapCaller(address common.Address, caller bind.ContractCaller) (*PancakeswapCaller, error) {
	contract, err := bindPancakeswap(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PancakeswapCaller{contract: contract}, nil
}

// NewPancakeswapTransactor creates a new write-only instance of Pancakeswap, bound to a specific deployed contract.
func NewPancakeswapTransactor(address common.Address, transactor bind.ContractTransactor) (*PancakeswapTransactor, error) {
	contract, err := bindPancakeswap(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PancakeswapTransactor{contract: contract}, nil
}

// NewPancakeswapFilterer creates a new log filterer instance of Pancakeswap, bound to a specific deployed contract.
func NewPancakeswapFilterer(address common.Address, filterer bind.ContractFilterer) (*PancakeswapFilterer, error) {
	contract, err := bindPancakeswap(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PancakeswapFilterer{contract: contract}, nil
}

// bindPancakeswap binds a generic wrapper to an already deployed contract.
func bindPancakeswap(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PancakeswapMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pancakeswap *PancakeswapRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pancakeswap.Contract.PancakeswapCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pancakeswap *PancakeswapRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pancakeswap.Contract.PancakeswapTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pancakeswap *PancakeswapRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pancakeswap.Contract.PancakeswapTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pancakeswap *PancakeswapCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pancakeswap.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pancakeswap *PancakeswapTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pancakeswap.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pancakeswap *PancakeswapTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pancakeswap.Contract.contract.Transact(opts, method, params...)
}

// FeeController is a free data retrieval call binding the contract method 0x6999b377.
//
// Solidity: function feeController() view returns(address)
func (_Pancakeswap *PancakeswapCaller) FeeController(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pancakeswap.contract.Call(opts, &out, "feeController")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeController is a free data retrieval call binding the contract method 0x6999b377.
//
// Solidity: function feeController() view returns(address)
func (_Pancakeswap *PancakeswapSession) FeeController() (common.Address, error) {
	return _Pancakeswap.Contract.FeeController(&_Pancakeswap.CallOpts)
}

// FeeController is a free data retrieval call binding the contract method 0x6999b377.
//
// Solidity: function feeController() view returns(address)
func (_Pancakeswap *PancakeswapCallerSession) FeeController() (common.Address, error) {
	return _Pancakeswap.Contract.FeeController(&_Pancakeswap.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Pancakeswap *PancakeswapCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pancakeswap.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Pancakeswap *PancakeswapSession) Owner() (common.Address, error) {
	return _Pancakeswap.Contract.Owner(&_Pancakeswap.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Pancakeswap *PancakeswapCallerSession) Owner() (common.Address, error) {
	return _Pancakeswap.Contract.Owner(&_Pancakeswap.CallOpts)
}

// Permit2 is a free data retrieval call binding the contract method 0x12261ee7.
//
// Solidity: function permit2() view returns(address)
func (_Pancakeswap *PancakeswapCaller) Permit2(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pancakeswap.contract.Call(opts, &out, "permit2")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Permit2 is a free data retrieval call binding the contract method 0x12261ee7.
//
// Solidity: function permit2() view returns(address)
func (_Pancakeswap *PancakeswapSession) Permit2() (common.Address, error) {
	return _Pancakeswap.Contract.Permit2(&_Pancakeswap.CallOpts)
}

// Permit2 is a free data retrieval call binding the contract method 0x12261ee7.
//
// Solidity: function permit2() view returns(address)
func (_Pancakeswap *PancakeswapCallerSession) Permit2() (common.Address, error) {
	return _Pancakeswap.Contract.Permit2(&_Pancakeswap.CallOpts)
}

// Execute is a paid mutator transaction binding the contract method 0x3f62192e.
//
// Solidity: function execute((bytes,bytes) order) payable returns()
func (_Pancakeswap *PancakeswapTransactor) Execute(opts *bind.TransactOpts, order SignedOrder) (*types.Transaction, error) {
	return _Pancakeswap.contract.Transact(opts, "execute", order)
}

// Execute is a paid mutator transaction binding the contract method 0x3f62192e.
//
// Solidity: function execute((bytes,bytes) order) payable returns()
func (_Pancakeswap *PancakeswapSession) Execute(order SignedOrder) (*types.Transaction, error) {
	return _Pancakeswap.Contract.Execute(&_Pancakeswap.TransactOpts, order)
}

// Execute is a paid mutator transaction binding the contract method 0x3f62192e.
//
// Solidity: function execute((bytes,bytes) order) payable returns()
func (_Pancakeswap *PancakeswapTransactorSession) Execute(order SignedOrder) (*types.Transaction, error) {
	return _Pancakeswap.Contract.Execute(&_Pancakeswap.TransactOpts, order)
}

// ExecuteBatch is a paid mutator transaction binding the contract method 0x0d7a16c3.
//
// Solidity: function executeBatch((bytes,bytes)[] orders) payable returns()
func (_Pancakeswap *PancakeswapTransactor) ExecuteBatch(opts *bind.TransactOpts, orders []SignedOrder) (*types.Transaction, error) {
	return _Pancakeswap.contract.Transact(opts, "executeBatch", orders)
}

// ExecuteBatch is a paid mutator transaction binding the contract method 0x0d7a16c3.
//
// Solidity: function executeBatch((bytes,bytes)[] orders) payable returns()
func (_Pancakeswap *PancakeswapSession) ExecuteBatch(orders []SignedOrder) (*types.Transaction, error) {
	return _Pancakeswap.Contract.ExecuteBatch(&_Pancakeswap.TransactOpts, orders)
}

// ExecuteBatch is a paid mutator transaction binding the contract method 0x0d7a16c3.
//
// Solidity: function executeBatch((bytes,bytes)[] orders) payable returns()
func (_Pancakeswap *PancakeswapTransactorSession) ExecuteBatch(orders []SignedOrder) (*types.Transaction, error) {
	return _Pancakeswap.Contract.ExecuteBatch(&_Pancakeswap.TransactOpts, orders)
}

// ExecuteBatchWithCallback is a paid mutator transaction binding the contract method 0x13fb72c7.
//
// Solidity: function executeBatchWithCallback((bytes,bytes)[] orders, bytes callbackData) payable returns()
func (_Pancakeswap *PancakeswapTransactor) ExecuteBatchWithCallback(opts *bind.TransactOpts, orders []SignedOrder, callbackData []byte) (*types.Transaction, error) {
	return _Pancakeswap.contract.Transact(opts, "executeBatchWithCallback", orders, callbackData)
}

// ExecuteBatchWithCallback is a paid mutator transaction binding the contract method 0x13fb72c7.
//
// Solidity: function executeBatchWithCallback((bytes,bytes)[] orders, bytes callbackData) payable returns()
func (_Pancakeswap *PancakeswapSession) ExecuteBatchWithCallback(orders []SignedOrder, callbackData []byte) (*types.Transaction, error) {
	return _Pancakeswap.Contract.ExecuteBatchWithCallback(&_Pancakeswap.TransactOpts, orders, callbackData)
}

// ExecuteBatchWithCallback is a paid mutator transaction binding the contract method 0x13fb72c7.
//
// Solidity: function executeBatchWithCallback((bytes,bytes)[] orders, bytes callbackData) payable returns()
func (_Pancakeswap *PancakeswapTransactorSession) ExecuteBatchWithCallback(orders []SignedOrder, callbackData []byte) (*types.Transaction, error) {
	return _Pancakeswap.Contract.ExecuteBatchWithCallback(&_Pancakeswap.TransactOpts, orders, callbackData)
}

// ExecuteWithCallback is a paid mutator transaction binding the contract method 0x0d335884.
//
// Solidity: function executeWithCallback((bytes,bytes) order, bytes callbackData) payable returns()
func (_Pancakeswap *PancakeswapTransactor) ExecuteWithCallback(opts *bind.TransactOpts, order SignedOrder, callbackData []byte) (*types.Transaction, error) {
	return _Pancakeswap.contract.Transact(opts, "executeWithCallback", order, callbackData)
}

// ExecuteWithCallback is a paid mutator transaction binding the contract method 0x0d335884.
//
// Solidity: function executeWithCallback((bytes,bytes) order, bytes callbackData) payable returns()
func (_Pancakeswap *PancakeswapSession) ExecuteWithCallback(order SignedOrder, callbackData []byte) (*types.Transaction, error) {
	return _Pancakeswap.Contract.ExecuteWithCallback(&_Pancakeswap.TransactOpts, order, callbackData)
}

// ExecuteWithCallback is a paid mutator transaction binding the contract method 0x0d335884.
//
// Solidity: function executeWithCallback((bytes,bytes) order, bytes callbackData) payable returns()
func (_Pancakeswap *PancakeswapTransactorSession) ExecuteWithCallback(order SignedOrder, callbackData []byte) (*types.Transaction, error) {
	return _Pancakeswap.Contract.ExecuteWithCallback(&_Pancakeswap.TransactOpts, order, callbackData)
}

// SetProtocolFeeController is a paid mutator transaction binding the contract method 0x2d771389.
//
// Solidity: function setProtocolFeeController(address _newFeeController) returns()
func (_Pancakeswap *PancakeswapTransactor) SetProtocolFeeController(opts *bind.TransactOpts, _newFeeController common.Address) (*types.Transaction, error) {
	return _Pancakeswap.contract.Transact(opts, "setProtocolFeeController", _newFeeController)
}

// SetProtocolFeeController is a paid mutator transaction binding the contract method 0x2d771389.
//
// Solidity: function setProtocolFeeController(address _newFeeController) returns()
func (_Pancakeswap *PancakeswapSession) SetProtocolFeeController(_newFeeController common.Address) (*types.Transaction, error) {
	return _Pancakeswap.Contract.SetProtocolFeeController(&_Pancakeswap.TransactOpts, _newFeeController)
}

// SetProtocolFeeController is a paid mutator transaction binding the contract method 0x2d771389.
//
// Solidity: function setProtocolFeeController(address _newFeeController) returns()
func (_Pancakeswap *PancakeswapTransactorSession) SetProtocolFeeController(_newFeeController common.Address) (*types.Transaction, error) {
	return _Pancakeswap.Contract.SetProtocolFeeController(&_Pancakeswap.TransactOpts, _newFeeController)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Pancakeswap *PancakeswapTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Pancakeswap.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Pancakeswap *PancakeswapSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Pancakeswap.Contract.TransferOwnership(&_Pancakeswap.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Pancakeswap *PancakeswapTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Pancakeswap.Contract.TransferOwnership(&_Pancakeswap.TransactOpts, newOwner)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Pancakeswap *PancakeswapTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pancakeswap.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Pancakeswap *PancakeswapSession) Receive() (*types.Transaction, error) {
	return _Pancakeswap.Contract.Receive(&_Pancakeswap.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Pancakeswap *PancakeswapTransactorSession) Receive() (*types.Transaction, error) {
	return _Pancakeswap.Contract.Receive(&_Pancakeswap.TransactOpts)
}

// PancakeswapFillIterator is returned from FilterFill and is used to iterate over the raw logs and unpacked data for Fill events raised by the Pancakeswap contract.
type PancakeswapFillIterator struct {
	Event *PancakeswapFill // Event containing the contract specifics and raw log

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
func (it *PancakeswapFillIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PancakeswapFill)
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
		it.Event = new(PancakeswapFill)
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
func (it *PancakeswapFillIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PancakeswapFillIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PancakeswapFill represents a Fill event raised by the Pancakeswap contract.
type PancakeswapFill struct {
	OrderHash [32]byte
	Filler    common.Address
	Swapper   common.Address
	Nonce     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterFill is a free log retrieval operation binding the contract event 0x78ad7ec0e9f89e74012afa58738b6b661c024cb0fd185ee2f616c0a28924bd66.
//
// Solidity: event Fill(bytes32 indexed orderHash, address indexed filler, address indexed swapper, uint256 nonce)
func (_Pancakeswap *PancakeswapFilterer) FilterFill(opts *bind.FilterOpts, orderHash [][32]byte, filler []common.Address, swapper []common.Address) (*PancakeswapFillIterator, error) {

	var orderHashRule []interface{}
	for _, orderHashItem := range orderHash {
		orderHashRule = append(orderHashRule, orderHashItem)
	}
	var fillerRule []interface{}
	for _, fillerItem := range filler {
		fillerRule = append(fillerRule, fillerItem)
	}
	var swapperRule []interface{}
	for _, swapperItem := range swapper {
		swapperRule = append(swapperRule, swapperItem)
	}

	logs, sub, err := _Pancakeswap.contract.FilterLogs(opts, "Fill", orderHashRule, fillerRule, swapperRule)
	if err != nil {
		return nil, err
	}
	return &PancakeswapFillIterator{contract: _Pancakeswap.contract, event: "Fill", logs: logs, sub: sub}, nil
}

// WatchFill is a free log subscription operation binding the contract event 0x78ad7ec0e9f89e74012afa58738b6b661c024cb0fd185ee2f616c0a28924bd66.
//
// Solidity: event Fill(bytes32 indexed orderHash, address indexed filler, address indexed swapper, uint256 nonce)
func (_Pancakeswap *PancakeswapFilterer) WatchFill(opts *bind.WatchOpts, sink chan<- *PancakeswapFill, orderHash [][32]byte, filler []common.Address, swapper []common.Address) (event.Subscription, error) {

	var orderHashRule []interface{}
	for _, orderHashItem := range orderHash {
		orderHashRule = append(orderHashRule, orderHashItem)
	}
	var fillerRule []interface{}
	for _, fillerItem := range filler {
		fillerRule = append(fillerRule, fillerItem)
	}
	var swapperRule []interface{}
	for _, swapperItem := range swapper {
		swapperRule = append(swapperRule, swapperItem)
	}

	logs, sub, err := _Pancakeswap.contract.WatchLogs(opts, "Fill", orderHashRule, fillerRule, swapperRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PancakeswapFill)
				if err := _Pancakeswap.contract.UnpackLog(event, "Fill", log); err != nil {
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

// ParseFill is a log parse operation binding the contract event 0x78ad7ec0e9f89e74012afa58738b6b661c024cb0fd185ee2f616c0a28924bd66.
//
// Solidity: event Fill(bytes32 indexed orderHash, address indexed filler, address indexed swapper, uint256 nonce)
func (_Pancakeswap *PancakeswapFilterer) ParseFill(log types.Log) (*PancakeswapFill, error) {
	event := new(PancakeswapFill)
	if err := _Pancakeswap.contract.UnpackLog(event, "Fill", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PancakeswapOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Pancakeswap contract.
type PancakeswapOwnershipTransferredIterator struct {
	Event *PancakeswapOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PancakeswapOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PancakeswapOwnershipTransferred)
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
		it.Event = new(PancakeswapOwnershipTransferred)
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
func (it *PancakeswapOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PancakeswapOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PancakeswapOwnershipTransferred represents a OwnershipTransferred event raised by the Pancakeswap contract.
type PancakeswapOwnershipTransferred struct {
	User     common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed user, address indexed newOwner)
func (_Pancakeswap *PancakeswapFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, user []common.Address, newOwner []common.Address) (*PancakeswapOwnershipTransferredIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Pancakeswap.contract.FilterLogs(opts, "OwnershipTransferred", userRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PancakeswapOwnershipTransferredIterator{contract: _Pancakeswap.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed user, address indexed newOwner)
func (_Pancakeswap *PancakeswapFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PancakeswapOwnershipTransferred, user []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Pancakeswap.contract.WatchLogs(opts, "OwnershipTransferred", userRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PancakeswapOwnershipTransferred)
				if err := _Pancakeswap.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
// Solidity: event OwnershipTransferred(address indexed user, address indexed newOwner)
func (_Pancakeswap *PancakeswapFilterer) ParseOwnershipTransferred(log types.Log) (*PancakeswapOwnershipTransferred, error) {
	event := new(PancakeswapOwnershipTransferred)
	if err := _Pancakeswap.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PancakeswapProtocolFeeControllerSetIterator is returned from FilterProtocolFeeControllerSet and is used to iterate over the raw logs and unpacked data for ProtocolFeeControllerSet events raised by the Pancakeswap contract.
type PancakeswapProtocolFeeControllerSetIterator struct {
	Event *PancakeswapProtocolFeeControllerSet // Event containing the contract specifics and raw log

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
func (it *PancakeswapProtocolFeeControllerSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PancakeswapProtocolFeeControllerSet)
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
		it.Event = new(PancakeswapProtocolFeeControllerSet)
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
func (it *PancakeswapProtocolFeeControllerSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PancakeswapProtocolFeeControllerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PancakeswapProtocolFeeControllerSet represents a ProtocolFeeControllerSet event raised by the Pancakeswap contract.
type PancakeswapProtocolFeeControllerSet struct {
	OldFeeController common.Address
	NewFeeController common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterProtocolFeeControllerSet is a free log retrieval operation binding the contract event 0xb904ae9529e373e48bc82df4326cceaf1b4c472babf37f5b7dec46fecc6b53e0.
//
// Solidity: event ProtocolFeeControllerSet(address oldFeeController, address newFeeController)
func (_Pancakeswap *PancakeswapFilterer) FilterProtocolFeeControllerSet(opts *bind.FilterOpts) (*PancakeswapProtocolFeeControllerSetIterator, error) {

	logs, sub, err := _Pancakeswap.contract.FilterLogs(opts, "ProtocolFeeControllerSet")
	if err != nil {
		return nil, err
	}
	return &PancakeswapProtocolFeeControllerSetIterator{contract: _Pancakeswap.contract, event: "ProtocolFeeControllerSet", logs: logs, sub: sub}, nil
}

// WatchProtocolFeeControllerSet is a free log subscription operation binding the contract event 0xb904ae9529e373e48bc82df4326cceaf1b4c472babf37f5b7dec46fecc6b53e0.
//
// Solidity: event ProtocolFeeControllerSet(address oldFeeController, address newFeeController)
func (_Pancakeswap *PancakeswapFilterer) WatchProtocolFeeControllerSet(opts *bind.WatchOpts, sink chan<- *PancakeswapProtocolFeeControllerSet) (event.Subscription, error) {

	logs, sub, err := _Pancakeswap.contract.WatchLogs(opts, "ProtocolFeeControllerSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PancakeswapProtocolFeeControllerSet)
				if err := _Pancakeswap.contract.UnpackLog(event, "ProtocolFeeControllerSet", log); err != nil {
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

// ParseProtocolFeeControllerSet is a log parse operation binding the contract event 0xb904ae9529e373e48bc82df4326cceaf1b4c472babf37f5b7dec46fecc6b53e0.
//
// Solidity: event ProtocolFeeControllerSet(address oldFeeController, address newFeeController)
func (_Pancakeswap *PancakeswapFilterer) ParseProtocolFeeControllerSet(log types.Log) (*PancakeswapProtocolFeeControllerSet, error) {
	event := new(PancakeswapProtocolFeeControllerSet)
	if err := _Pancakeswap.contract.UnpackLog(event, "ProtocolFeeControllerSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
