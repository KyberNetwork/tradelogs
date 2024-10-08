// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package uniswapxv1

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

// Uniswapxv1MetaData contains all meta data concerning the Uniswapxv1 contract.
var Uniswapxv1MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIPermit2\",\"name\":\"_permit2\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_protocolFeeOwner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"DeadlineBeforeEndTime\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DeadlinePassed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"duplicateToken\",\"type\":\"address\"}],\"name\":\"DuplicateFeeOutput\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EndTimeBeforeStartTime\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"FeeTooLarge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IncorrectAmounts\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InputAndOutputDecay\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientEth\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"}],\"name\":\"InvalidFeeToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidReactor\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NativeTransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoExclusiveOverride\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OrderEndTimeBeforeStartTime\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"filler\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"swapper\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"Fill\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldFeeController\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newFeeController\",\"type\":\"address\"}],\"name\":\"ProtocolFeeControllerSet\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"order\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"internalType\":\"structSignedOrder\",\"name\":\"order\",\"type\":\"tuple\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"order\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"internalType\":\"structSignedOrder[]\",\"name\":\"orders\",\"type\":\"tuple[]\"}],\"name\":\"executeBatch\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"order\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"internalType\":\"structSignedOrder[]\",\"name\":\"orders\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"callbackData\",\"type\":\"bytes\"}],\"name\":\"executeBatchWithCallback\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"order\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"internalType\":\"structSignedOrder\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"callbackData\",\"type\":\"bytes\"}],\"name\":\"executeWithCallback\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeController\",\"outputs\":[{\"internalType\":\"contractIProtocolFeeController\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"permit2\",\"outputs\":[{\"internalType\":\"contractIPermit2\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newFeeController\",\"type\":\"address\"}],\"name\":\"setProtocolFeeController\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// Uniswapxv1ABI is the input ABI used to generate the binding from.
// Deprecated: Use Uniswapxv1MetaData.ABI instead.
var Uniswapxv1ABI = Uniswapxv1MetaData.ABI

// Uniswapxv1 is an auto generated Go binding around an Ethereum contract.
type Uniswapxv1 struct {
	Uniswapxv1Caller     // Read-only binding to the contract
	Uniswapxv1Transactor // Write-only binding to the contract
	Uniswapxv1Filterer   // Log filterer for contract events
}

// Uniswapxv1Caller is an auto generated read-only Go binding around an Ethereum contract.
type Uniswapxv1Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Uniswapxv1Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Uniswapxv1Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Uniswapxv1Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Uniswapxv1Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Uniswapxv1Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Uniswapxv1Session struct {
	Contract     *Uniswapxv1       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Uniswapxv1CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Uniswapxv1CallerSession struct {
	Contract *Uniswapxv1Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// Uniswapxv1TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Uniswapxv1TransactorSession struct {
	Contract     *Uniswapxv1Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// Uniswapxv1Raw is an auto generated low-level Go binding around an Ethereum contract.
type Uniswapxv1Raw struct {
	Contract *Uniswapxv1 // Generic contract binding to access the raw methods on
}

// Uniswapxv1CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Uniswapxv1CallerRaw struct {
	Contract *Uniswapxv1Caller // Generic read-only contract binding to access the raw methods on
}

// Uniswapxv1TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Uniswapxv1TransactorRaw struct {
	Contract *Uniswapxv1Transactor // Generic write-only contract binding to access the raw methods on
}

// NewUniswapxv1 creates a new instance of Uniswapxv1, bound to a specific deployed contract.
func NewUniswapxv1(address common.Address, backend bind.ContractBackend) (*Uniswapxv1, error) {
	contract, err := bindUniswapxv1(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Uniswapxv1{Uniswapxv1Caller: Uniswapxv1Caller{contract: contract}, Uniswapxv1Transactor: Uniswapxv1Transactor{contract: contract}, Uniswapxv1Filterer: Uniswapxv1Filterer{contract: contract}}, nil
}

// NewUniswapxv1Caller creates a new read-only instance of Uniswapxv1, bound to a specific deployed contract.
func NewUniswapxv1Caller(address common.Address, caller bind.ContractCaller) (*Uniswapxv1Caller, error) {
	contract, err := bindUniswapxv1(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Uniswapxv1Caller{contract: contract}, nil
}

// NewUniswapxv1Transactor creates a new write-only instance of Uniswapxv1, bound to a specific deployed contract.
func NewUniswapxv1Transactor(address common.Address, transactor bind.ContractTransactor) (*Uniswapxv1Transactor, error) {
	contract, err := bindUniswapxv1(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Uniswapxv1Transactor{contract: contract}, nil
}

// NewUniswapxv1Filterer creates a new log filterer instance of Uniswapxv1, bound to a specific deployed contract.
func NewUniswapxv1Filterer(address common.Address, filterer bind.ContractFilterer) (*Uniswapxv1Filterer, error) {
	contract, err := bindUniswapxv1(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Uniswapxv1Filterer{contract: contract}, nil
}

// bindUniswapxv1 binds a generic wrapper to an already deployed contract.
func bindUniswapxv1(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Uniswapxv1MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Uniswapxv1 *Uniswapxv1Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Uniswapxv1.Contract.Uniswapxv1Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Uniswapxv1 *Uniswapxv1Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Uniswapxv1.Contract.Uniswapxv1Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Uniswapxv1 *Uniswapxv1Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Uniswapxv1.Contract.Uniswapxv1Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Uniswapxv1 *Uniswapxv1CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Uniswapxv1.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Uniswapxv1 *Uniswapxv1TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Uniswapxv1.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Uniswapxv1 *Uniswapxv1TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Uniswapxv1.Contract.contract.Transact(opts, method, params...)
}

// FeeController is a free data retrieval call binding the contract method 0x6999b377.
//
// Solidity: function feeController() view returns(address)
func (_Uniswapxv1 *Uniswapxv1Caller) FeeController(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Uniswapxv1.contract.Call(opts, &out, "feeController")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeController is a free data retrieval call binding the contract method 0x6999b377.
//
// Solidity: function feeController() view returns(address)
func (_Uniswapxv1 *Uniswapxv1Session) FeeController() (common.Address, error) {
	return _Uniswapxv1.Contract.FeeController(&_Uniswapxv1.CallOpts)
}

// FeeController is a free data retrieval call binding the contract method 0x6999b377.
//
// Solidity: function feeController() view returns(address)
func (_Uniswapxv1 *Uniswapxv1CallerSession) FeeController() (common.Address, error) {
	return _Uniswapxv1.Contract.FeeController(&_Uniswapxv1.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Uniswapxv1 *Uniswapxv1Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Uniswapxv1.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Uniswapxv1 *Uniswapxv1Session) Owner() (common.Address, error) {
	return _Uniswapxv1.Contract.Owner(&_Uniswapxv1.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Uniswapxv1 *Uniswapxv1CallerSession) Owner() (common.Address, error) {
	return _Uniswapxv1.Contract.Owner(&_Uniswapxv1.CallOpts)
}

// Permit2 is a free data retrieval call binding the contract method 0x12261ee7.
//
// Solidity: function permit2() view returns(address)
func (_Uniswapxv1 *Uniswapxv1Caller) Permit2(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Uniswapxv1.contract.Call(opts, &out, "permit2")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Permit2 is a free data retrieval call binding the contract method 0x12261ee7.
//
// Solidity: function permit2() view returns(address)
func (_Uniswapxv1 *Uniswapxv1Session) Permit2() (common.Address, error) {
	return _Uniswapxv1.Contract.Permit2(&_Uniswapxv1.CallOpts)
}

// Permit2 is a free data retrieval call binding the contract method 0x12261ee7.
//
// Solidity: function permit2() view returns(address)
func (_Uniswapxv1 *Uniswapxv1CallerSession) Permit2() (common.Address, error) {
	return _Uniswapxv1.Contract.Permit2(&_Uniswapxv1.CallOpts)
}

// Execute is a paid mutator transaction binding the contract method 0x3f62192e.
//
// Solidity: function execute((bytes,bytes) order) payable returns()
func (_Uniswapxv1 *Uniswapxv1Transactor) Execute(opts *bind.TransactOpts, order SignedOrder) (*types.Transaction, error) {
	return _Uniswapxv1.contract.Transact(opts, "execute", order)
}

// Execute is a paid mutator transaction binding the contract method 0x3f62192e.
//
// Solidity: function execute((bytes,bytes) order) payable returns()
func (_Uniswapxv1 *Uniswapxv1Session) Execute(order SignedOrder) (*types.Transaction, error) {
	return _Uniswapxv1.Contract.Execute(&_Uniswapxv1.TransactOpts, order)
}

// Execute is a paid mutator transaction binding the contract method 0x3f62192e.
//
// Solidity: function execute((bytes,bytes) order) payable returns()
func (_Uniswapxv1 *Uniswapxv1TransactorSession) Execute(order SignedOrder) (*types.Transaction, error) {
	return _Uniswapxv1.Contract.Execute(&_Uniswapxv1.TransactOpts, order)
}

// ExecuteBatch is a paid mutator transaction binding the contract method 0x0d7a16c3.
//
// Solidity: function executeBatch((bytes,bytes)[] orders) payable returns()
func (_Uniswapxv1 *Uniswapxv1Transactor) ExecuteBatch(opts *bind.TransactOpts, orders []SignedOrder) (*types.Transaction, error) {
	return _Uniswapxv1.contract.Transact(opts, "executeBatch", orders)
}

// ExecuteBatch is a paid mutator transaction binding the contract method 0x0d7a16c3.
//
// Solidity: function executeBatch((bytes,bytes)[] orders) payable returns()
func (_Uniswapxv1 *Uniswapxv1Session) ExecuteBatch(orders []SignedOrder) (*types.Transaction, error) {
	return _Uniswapxv1.Contract.ExecuteBatch(&_Uniswapxv1.TransactOpts, orders)
}

// ExecuteBatch is a paid mutator transaction binding the contract method 0x0d7a16c3.
//
// Solidity: function executeBatch((bytes,bytes)[] orders) payable returns()
func (_Uniswapxv1 *Uniswapxv1TransactorSession) ExecuteBatch(orders []SignedOrder) (*types.Transaction, error) {
	return _Uniswapxv1.Contract.ExecuteBatch(&_Uniswapxv1.TransactOpts, orders)
}

// ExecuteBatchWithCallback is a paid mutator transaction binding the contract method 0x13fb72c7.
//
// Solidity: function executeBatchWithCallback((bytes,bytes)[] orders, bytes callbackData) payable returns()
func (_Uniswapxv1 *Uniswapxv1Transactor) ExecuteBatchWithCallback(opts *bind.TransactOpts, orders []SignedOrder, callbackData []byte) (*types.Transaction, error) {
	return _Uniswapxv1.contract.Transact(opts, "executeBatchWithCallback", orders, callbackData)
}

// ExecuteBatchWithCallback is a paid mutator transaction binding the contract method 0x13fb72c7.
//
// Solidity: function executeBatchWithCallback((bytes,bytes)[] orders, bytes callbackData) payable returns()
func (_Uniswapxv1 *Uniswapxv1Session) ExecuteBatchWithCallback(orders []SignedOrder, callbackData []byte) (*types.Transaction, error) {
	return _Uniswapxv1.Contract.ExecuteBatchWithCallback(&_Uniswapxv1.TransactOpts, orders, callbackData)
}

// ExecuteBatchWithCallback is a paid mutator transaction binding the contract method 0x13fb72c7.
//
// Solidity: function executeBatchWithCallback((bytes,bytes)[] orders, bytes callbackData) payable returns()
func (_Uniswapxv1 *Uniswapxv1TransactorSession) ExecuteBatchWithCallback(orders []SignedOrder, callbackData []byte) (*types.Transaction, error) {
	return _Uniswapxv1.Contract.ExecuteBatchWithCallback(&_Uniswapxv1.TransactOpts, orders, callbackData)
}

// ExecuteWithCallback is a paid mutator transaction binding the contract method 0x0d335884.
//
// Solidity: function executeWithCallback((bytes,bytes) order, bytes callbackData) payable returns()
func (_Uniswapxv1 *Uniswapxv1Transactor) ExecuteWithCallback(opts *bind.TransactOpts, order SignedOrder, callbackData []byte) (*types.Transaction, error) {
	return _Uniswapxv1.contract.Transact(opts, "executeWithCallback", order, callbackData)
}

// ExecuteWithCallback is a paid mutator transaction binding the contract method 0x0d335884.
//
// Solidity: function executeWithCallback((bytes,bytes) order, bytes callbackData) payable returns()
func (_Uniswapxv1 *Uniswapxv1Session) ExecuteWithCallback(order SignedOrder, callbackData []byte) (*types.Transaction, error) {
	return _Uniswapxv1.Contract.ExecuteWithCallback(&_Uniswapxv1.TransactOpts, order, callbackData)
}

// ExecuteWithCallback is a paid mutator transaction binding the contract method 0x0d335884.
//
// Solidity: function executeWithCallback((bytes,bytes) order, bytes callbackData) payable returns()
func (_Uniswapxv1 *Uniswapxv1TransactorSession) ExecuteWithCallback(order SignedOrder, callbackData []byte) (*types.Transaction, error) {
	return _Uniswapxv1.Contract.ExecuteWithCallback(&_Uniswapxv1.TransactOpts, order, callbackData)
}

// SetProtocolFeeController is a paid mutator transaction binding the contract method 0x2d771389.
//
// Solidity: function setProtocolFeeController(address _newFeeController) returns()
func (_Uniswapxv1 *Uniswapxv1Transactor) SetProtocolFeeController(opts *bind.TransactOpts, _newFeeController common.Address) (*types.Transaction, error) {
	return _Uniswapxv1.contract.Transact(opts, "setProtocolFeeController", _newFeeController)
}

// SetProtocolFeeController is a paid mutator transaction binding the contract method 0x2d771389.
//
// Solidity: function setProtocolFeeController(address _newFeeController) returns()
func (_Uniswapxv1 *Uniswapxv1Session) SetProtocolFeeController(_newFeeController common.Address) (*types.Transaction, error) {
	return _Uniswapxv1.Contract.SetProtocolFeeController(&_Uniswapxv1.TransactOpts, _newFeeController)
}

// SetProtocolFeeController is a paid mutator transaction binding the contract method 0x2d771389.
//
// Solidity: function setProtocolFeeController(address _newFeeController) returns()
func (_Uniswapxv1 *Uniswapxv1TransactorSession) SetProtocolFeeController(_newFeeController common.Address) (*types.Transaction, error) {
	return _Uniswapxv1.Contract.SetProtocolFeeController(&_Uniswapxv1.TransactOpts, _newFeeController)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Uniswapxv1 *Uniswapxv1Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Uniswapxv1.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Uniswapxv1 *Uniswapxv1Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Uniswapxv1.Contract.TransferOwnership(&_Uniswapxv1.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Uniswapxv1 *Uniswapxv1TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Uniswapxv1.Contract.TransferOwnership(&_Uniswapxv1.TransactOpts, newOwner)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Uniswapxv1 *Uniswapxv1Transactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Uniswapxv1.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Uniswapxv1 *Uniswapxv1Session) Receive() (*types.Transaction, error) {
	return _Uniswapxv1.Contract.Receive(&_Uniswapxv1.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Uniswapxv1 *Uniswapxv1TransactorSession) Receive() (*types.Transaction, error) {
	return _Uniswapxv1.Contract.Receive(&_Uniswapxv1.TransactOpts)
}

// Uniswapxv1FillIterator is returned from FilterFill and is used to iterate over the raw logs and unpacked data for Fill events raised by the Uniswapxv1 contract.
type Uniswapxv1FillIterator struct {
	Event *Uniswapxv1Fill // Event containing the contract specifics and raw log

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
func (it *Uniswapxv1FillIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Uniswapxv1Fill)
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
		it.Event = new(Uniswapxv1Fill)
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
func (it *Uniswapxv1FillIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Uniswapxv1FillIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Uniswapxv1Fill represents a Fill event raised by the Uniswapxv1 contract.
type Uniswapxv1Fill struct {
	OrderHash [32]byte
	Filler    common.Address
	Swapper   common.Address
	Nonce     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterFill is a free log retrieval operation binding the contract event 0x78ad7ec0e9f89e74012afa58738b6b661c024cb0fd185ee2f616c0a28924bd66.
//
// Solidity: event Fill(bytes32 indexed orderHash, address indexed filler, address indexed swapper, uint256 nonce)
func (_Uniswapxv1 *Uniswapxv1Filterer) FilterFill(opts *bind.FilterOpts, orderHash [][32]byte, filler []common.Address, swapper []common.Address) (*Uniswapxv1FillIterator, error) {

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

	logs, sub, err := _Uniswapxv1.contract.FilterLogs(opts, "Fill", orderHashRule, fillerRule, swapperRule)
	if err != nil {
		return nil, err
	}
	return &Uniswapxv1FillIterator{contract: _Uniswapxv1.contract, event: "Fill", logs: logs, sub: sub}, nil
}

// WatchFill is a free log subscription operation binding the contract event 0x78ad7ec0e9f89e74012afa58738b6b661c024cb0fd185ee2f616c0a28924bd66.
//
// Solidity: event Fill(bytes32 indexed orderHash, address indexed filler, address indexed swapper, uint256 nonce)
func (_Uniswapxv1 *Uniswapxv1Filterer) WatchFill(opts *bind.WatchOpts, sink chan<- *Uniswapxv1Fill, orderHash [][32]byte, filler []common.Address, swapper []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Uniswapxv1.contract.WatchLogs(opts, "Fill", orderHashRule, fillerRule, swapperRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Uniswapxv1Fill)
				if err := _Uniswapxv1.contract.UnpackLog(event, "Fill", log); err != nil {
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
func (_Uniswapxv1 *Uniswapxv1Filterer) ParseFill(log types.Log) (*Uniswapxv1Fill, error) {
	event := new(Uniswapxv1Fill)
	if err := _Uniswapxv1.contract.UnpackLog(event, "Fill", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Uniswapxv1OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Uniswapxv1 contract.
type Uniswapxv1OwnershipTransferredIterator struct {
	Event *Uniswapxv1OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *Uniswapxv1OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Uniswapxv1OwnershipTransferred)
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
		it.Event = new(Uniswapxv1OwnershipTransferred)
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
func (it *Uniswapxv1OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Uniswapxv1OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Uniswapxv1OwnershipTransferred represents a OwnershipTransferred event raised by the Uniswapxv1 contract.
type Uniswapxv1OwnershipTransferred struct {
	User     common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed user, address indexed newOwner)
func (_Uniswapxv1 *Uniswapxv1Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, user []common.Address, newOwner []common.Address) (*Uniswapxv1OwnershipTransferredIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Uniswapxv1.contract.FilterLogs(opts, "OwnershipTransferred", userRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &Uniswapxv1OwnershipTransferredIterator{contract: _Uniswapxv1.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed user, address indexed newOwner)
func (_Uniswapxv1 *Uniswapxv1Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *Uniswapxv1OwnershipTransferred, user []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Uniswapxv1.contract.WatchLogs(opts, "OwnershipTransferred", userRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Uniswapxv1OwnershipTransferred)
				if err := _Uniswapxv1.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Uniswapxv1 *Uniswapxv1Filterer) ParseOwnershipTransferred(log types.Log) (*Uniswapxv1OwnershipTransferred, error) {
	event := new(Uniswapxv1OwnershipTransferred)
	if err := _Uniswapxv1.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Uniswapxv1ProtocolFeeControllerSetIterator is returned from FilterProtocolFeeControllerSet and is used to iterate over the raw logs and unpacked data for ProtocolFeeControllerSet events raised by the Uniswapxv1 contract.
type Uniswapxv1ProtocolFeeControllerSetIterator struct {
	Event *Uniswapxv1ProtocolFeeControllerSet // Event containing the contract specifics and raw log

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
func (it *Uniswapxv1ProtocolFeeControllerSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Uniswapxv1ProtocolFeeControllerSet)
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
		it.Event = new(Uniswapxv1ProtocolFeeControllerSet)
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
func (it *Uniswapxv1ProtocolFeeControllerSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Uniswapxv1ProtocolFeeControllerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Uniswapxv1ProtocolFeeControllerSet represents a ProtocolFeeControllerSet event raised by the Uniswapxv1 contract.
type Uniswapxv1ProtocolFeeControllerSet struct {
	OldFeeController common.Address
	NewFeeController common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterProtocolFeeControllerSet is a free log retrieval operation binding the contract event 0xb904ae9529e373e48bc82df4326cceaf1b4c472babf37f5b7dec46fecc6b53e0.
//
// Solidity: event ProtocolFeeControllerSet(address oldFeeController, address newFeeController)
func (_Uniswapxv1 *Uniswapxv1Filterer) FilterProtocolFeeControllerSet(opts *bind.FilterOpts) (*Uniswapxv1ProtocolFeeControllerSetIterator, error) {

	logs, sub, err := _Uniswapxv1.contract.FilterLogs(opts, "ProtocolFeeControllerSet")
	if err != nil {
		return nil, err
	}
	return &Uniswapxv1ProtocolFeeControllerSetIterator{contract: _Uniswapxv1.contract, event: "ProtocolFeeControllerSet", logs: logs, sub: sub}, nil
}

// WatchProtocolFeeControllerSet is a free log subscription operation binding the contract event 0xb904ae9529e373e48bc82df4326cceaf1b4c472babf37f5b7dec46fecc6b53e0.
//
// Solidity: event ProtocolFeeControllerSet(address oldFeeController, address newFeeController)
func (_Uniswapxv1 *Uniswapxv1Filterer) WatchProtocolFeeControllerSet(opts *bind.WatchOpts, sink chan<- *Uniswapxv1ProtocolFeeControllerSet) (event.Subscription, error) {

	logs, sub, err := _Uniswapxv1.contract.WatchLogs(opts, "ProtocolFeeControllerSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Uniswapxv1ProtocolFeeControllerSet)
				if err := _Uniswapxv1.contract.UnpackLog(event, "ProtocolFeeControllerSet", log); err != nil {
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
func (_Uniswapxv1 *Uniswapxv1Filterer) ParseProtocolFeeControllerSet(log types.Log) (*Uniswapxv1ProtocolFeeControllerSet, error) {
	event := new(Uniswapxv1ProtocolFeeControllerSet)
	if err := _Uniswapxv1.contract.UnpackLog(event, "ProtocolFeeControllerSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
