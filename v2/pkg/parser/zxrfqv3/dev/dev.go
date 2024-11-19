// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package dev

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

// SettlerAllowedSlippage is an auto generated low-level Go binding around an user-defined struct.
type SettlerAllowedSlippage struct {
	BuyToken     common.Address
	Recipient    common.Address
	MinAmountOut *big.Int
}

// DevMetaData contains all meta data concerning the Dev contract.
var DevMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"uniFactory\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"poolInitCodeHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"dai\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"},{\"internalType\":\"bytes4\",\"name\":\"action\",\"type\":\"bytes4\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"ActionInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ConfusedDeputy\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForwarderNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidOffset\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSignatureLen\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTarget\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"oldOperator\",\"type\":\"address\"}],\"name\":\"OperatorNotSpent\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"oldOperator\",\"type\":\"address\"}],\"name\":\"ReentrantCallback\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"oldWitness\",\"type\":\"bytes32\"}],\"name\":\"ReentrantMetatransaction\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"TooMuchSlippage\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"oldWitness\",\"type\":\"bytes32\"}],\"name\":\"WitnessNotSpent\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroSwapAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"actions\",\"type\":\"bytes[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"buyToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minAmountOut\",\"type\":\"uint256\"}],\"internalType\":\"structSettler.AllowedSlippage\",\"name\":\"slippage\",\"type\":\"tuple\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"actions\",\"type\":\"bytes[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"buyToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minAmountOut\",\"type\":\"uint256\"}],\"internalType\":\"structSettler.AllowedSlippage\",\"name\":\"slippage\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"msgSender\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"executeMetaTxn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"amount0Delta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"amount1Delta\",\"type\":\"int256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"uniswapV3SwapCallback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// DevABI is the input ABI used to generate the binding from.
// Deprecated: Use DevMetaData.ABI instead.
var DevABI = DevMetaData.ABI

// Dev is an auto generated Go binding around an Ethereum contract.
type Dev struct {
	DevCaller     // Read-only binding to the contract
	DevTransactor // Write-only binding to the contract
	DevFilterer   // Log filterer for contract events
}

// DevCaller is an auto generated read-only Go binding around an Ethereum contract.
type DevCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DevTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DevTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DevFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DevFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DevSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DevSession struct {
	Contract     *Dev              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DevCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DevCallerSession struct {
	Contract *DevCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// DevTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DevTransactorSession struct {
	Contract     *DevTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DevRaw is an auto generated low-level Go binding around an Ethereum contract.
type DevRaw struct {
	Contract *Dev // Generic contract binding to access the raw methods on
}

// DevCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DevCallerRaw struct {
	Contract *DevCaller // Generic read-only contract binding to access the raw methods on
}

// DevTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DevTransactorRaw struct {
	Contract *DevTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDev creates a new instance of Dev, bound to a specific deployed contract.
func NewDev(address common.Address, backend bind.ContractBackend) (*Dev, error) {
	contract, err := bindDev(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Dev{DevCaller: DevCaller{contract: contract}, DevTransactor: DevTransactor{contract: contract}, DevFilterer: DevFilterer{contract: contract}}, nil
}

// NewDevCaller creates a new read-only instance of Dev, bound to a specific deployed contract.
func NewDevCaller(address common.Address, caller bind.ContractCaller) (*DevCaller, error) {
	contract, err := bindDev(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DevCaller{contract: contract}, nil
}

// NewDevTransactor creates a new write-only instance of Dev, bound to a specific deployed contract.
func NewDevTransactor(address common.Address, transactor bind.ContractTransactor) (*DevTransactor, error) {
	contract, err := bindDev(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DevTransactor{contract: contract}, nil
}

// NewDevFilterer creates a new log filterer instance of Dev, bound to a specific deployed contract.
func NewDevFilterer(address common.Address, filterer bind.ContractFilterer) (*DevFilterer, error) {
	contract, err := bindDev(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DevFilterer{contract: contract}, nil
}

// bindDev binds a generic wrapper to an already deployed contract.
func bindDev(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DevMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dev *DevRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Dev.Contract.DevCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dev *DevRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dev.Contract.DevTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dev *DevRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Dev.Contract.DevTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dev *DevCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Dev.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dev *DevTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dev.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dev *DevTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Dev.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) pure returns()
func (_Dev *DevCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) error {
	var out []interface{}
	err := _Dev.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return err
	}

	return err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) pure returns()
func (_Dev *DevSession) BalanceOf(arg0 common.Address) error {
	return _Dev.Contract.BalanceOf(&_Dev.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) pure returns()
func (_Dev *DevCallerSession) BalanceOf(arg0 common.Address) error {
	return _Dev.Contract.BalanceOf(&_Dev.CallOpts, arg0)
}

// Execute is a paid mutator transaction binding the contract method 0xe4b11b29.
//
// Solidity: function execute(bytes[] actions, (address,address,uint256) slippage) payable returns()
func (_Dev *DevTransactor) Execute(opts *bind.TransactOpts, actions [][]byte, slippage SettlerAllowedSlippage) (*types.Transaction, error) {
	return _Dev.contract.Transact(opts, "execute", actions, slippage)
}

// Execute is a paid mutator transaction binding the contract method 0xe4b11b29.
//
// Solidity: function execute(bytes[] actions, (address,address,uint256) slippage) payable returns()
func (_Dev *DevSession) Execute(actions [][]byte, slippage SettlerAllowedSlippage) (*types.Transaction, error) {
	return _Dev.Contract.Execute(&_Dev.TransactOpts, actions, slippage)
}

// Execute is a paid mutator transaction binding the contract method 0xe4b11b29.
//
// Solidity: function execute(bytes[] actions, (address,address,uint256) slippage) payable returns()
func (_Dev *DevTransactorSession) Execute(actions [][]byte, slippage SettlerAllowedSlippage) (*types.Transaction, error) {
	return _Dev.Contract.Execute(&_Dev.TransactOpts, actions, slippage)
}

// ExecuteMetaTxn is a paid mutator transaction binding the contract method 0xdd304e43.
//
// Solidity: function executeMetaTxn(bytes[] actions, (address,address,uint256) slippage, address msgSender, bytes sig) returns()
func (_Dev *DevTransactor) ExecuteMetaTxn(opts *bind.TransactOpts, actions [][]byte, slippage SettlerAllowedSlippage, msgSender common.Address, sig []byte) (*types.Transaction, error) {
	return _Dev.contract.Transact(opts, "executeMetaTxn", actions, slippage, msgSender, sig)
}

// ExecuteMetaTxn is a paid mutator transaction binding the contract method 0xdd304e43.
//
// Solidity: function executeMetaTxn(bytes[] actions, (address,address,uint256) slippage, address msgSender, bytes sig) returns()
func (_Dev *DevSession) ExecuteMetaTxn(actions [][]byte, slippage SettlerAllowedSlippage, msgSender common.Address, sig []byte) (*types.Transaction, error) {
	return _Dev.Contract.ExecuteMetaTxn(&_Dev.TransactOpts, actions, slippage, msgSender, sig)
}

// ExecuteMetaTxn is a paid mutator transaction binding the contract method 0xdd304e43.
//
// Solidity: function executeMetaTxn(bytes[] actions, (address,address,uint256) slippage, address msgSender, bytes sig) returns()
func (_Dev *DevTransactorSession) ExecuteMetaTxn(actions [][]byte, slippage SettlerAllowedSlippage, msgSender common.Address, sig []byte) (*types.Transaction, error) {
	return _Dev.Contract.ExecuteMetaTxn(&_Dev.TransactOpts, actions, slippage, msgSender, sig)
}

// UniswapV3SwapCallback is a paid mutator transaction binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes data) returns()
func (_Dev *DevTransactor) UniswapV3SwapCallback(opts *bind.TransactOpts, amount0Delta *big.Int, amount1Delta *big.Int, data []byte) (*types.Transaction, error) {
	return _Dev.contract.Transact(opts, "uniswapV3SwapCallback", amount0Delta, amount1Delta, data)
}

// UniswapV3SwapCallback is a paid mutator transaction binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes data) returns()
func (_Dev *DevSession) UniswapV3SwapCallback(amount0Delta *big.Int, amount1Delta *big.Int, data []byte) (*types.Transaction, error) {
	return _Dev.Contract.UniswapV3SwapCallback(&_Dev.TransactOpts, amount0Delta, amount1Delta, data)
}

// UniswapV3SwapCallback is a paid mutator transaction binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes data) returns()
func (_Dev *DevTransactorSession) UniswapV3SwapCallback(amount0Delta *big.Int, amount1Delta *big.Int, data []byte) (*types.Transaction, error) {
	return _Dev.Contract.UniswapV3SwapCallback(&_Dev.TransactOpts, amount0Delta, amount1Delta, data)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Dev *DevTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dev.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Dev *DevSession) Receive() (*types.Transaction, error) {
	return _Dev.Contract.Receive(&_Dev.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Dev *DevTransactorSession) Receive() (*types.Transaction, error) {
	return _Dev.Contract.Receive(&_Dev.TransactOpts)
}
