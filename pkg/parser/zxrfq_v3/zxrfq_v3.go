// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package zxrfq_v3

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

// ZxrfqV3MetaData contains all meta data concerning the ZxrfqV3 contract.
var ZxrfqV3MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"uniFactory\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"poolInitCodeHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"dai\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"},{\"internalType\":\"bytes4\",\"name\":\"action\",\"type\":\"bytes4\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"ActionInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ConfusedDeputy\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForwarderNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidOffset\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSignatureLen\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTarget\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"oldOperator\",\"type\":\"address\"}],\"name\":\"OperatorNotSpent\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"oldOperator\",\"type\":\"address\"}],\"name\":\"ReentrantCallback\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"oldWitness\",\"type\":\"bytes32\"}],\"name\":\"ReentrantMetatransaction\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"TooMuchSlippage\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"oldWitness\",\"type\":\"bytes32\"}],\"name\":\"WitnessNotSpent\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroSwapAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"actions\",\"type\":\"bytes[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"buyToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minAmountOut\",\"type\":\"uint256\"}],\"internalType\":\"structSettler.AllowedSlippage\",\"name\":\"slippage\",\"type\":\"tuple\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"actions\",\"type\":\"bytes[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"buyToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minAmountOut\",\"type\":\"uint256\"}],\"internalType\":\"structSettler.AllowedSlippage\",\"name\":\"slippage\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"msgSender\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"executeMetaTxn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"amount0Delta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"amount1Delta\",\"type\":\"int256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"uniswapV3SwapCallback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// ZxrfqV3ABI is the input ABI used to generate the binding from.
// Deprecated: Use ZxrfqV3MetaData.ABI instead.
var ZxrfqV3ABI = ZxrfqV3MetaData.ABI

// ZxrfqV3 is an auto generated Go binding around an Ethereum contract.
type ZxrfqV3 struct {
	ZxrfqV3Caller     // Read-only binding to the contract
	ZxrfqV3Transactor // Write-only binding to the contract
	ZxrfqV3Filterer   // Log filterer for contract events
}

// ZxrfqV3Caller is an auto generated read-only Go binding around an Ethereum contract.
type ZxrfqV3Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZxrfqV3Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ZxrfqV3Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZxrfqV3Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZxrfqV3Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZxrfqV3Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZxrfqV3Session struct {
	Contract     *ZxrfqV3          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZxrfqV3CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZxrfqV3CallerSession struct {
	Contract *ZxrfqV3Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ZxrfqV3TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZxrfqV3TransactorSession struct {
	Contract     *ZxrfqV3Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ZxrfqV3Raw is an auto generated low-level Go binding around an Ethereum contract.
type ZxrfqV3Raw struct {
	Contract *ZxrfqV3 // Generic contract binding to access the raw methods on
}

// ZxrfqV3CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZxrfqV3CallerRaw struct {
	Contract *ZxrfqV3Caller // Generic read-only contract binding to access the raw methods on
}

// ZxrfqV3TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZxrfqV3TransactorRaw struct {
	Contract *ZxrfqV3Transactor // Generic write-only contract binding to access the raw methods on
}

// NewZxrfqV3 creates a new instance of ZxrfqV3, bound to a specific deployed contract.
func NewZxrfqV3(address common.Address, backend bind.ContractBackend) (*ZxrfqV3, error) {
	contract, err := bindZxrfqV3(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZxrfqV3{ZxrfqV3Caller: ZxrfqV3Caller{contract: contract}, ZxrfqV3Transactor: ZxrfqV3Transactor{contract: contract}, ZxrfqV3Filterer: ZxrfqV3Filterer{contract: contract}}, nil
}

// NewZxrfqV3Caller creates a new read-only instance of ZxrfqV3, bound to a specific deployed contract.
func NewZxrfqV3Caller(address common.Address, caller bind.ContractCaller) (*ZxrfqV3Caller, error) {
	contract, err := bindZxrfqV3(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZxrfqV3Caller{contract: contract}, nil
}

// NewZxrfqV3Transactor creates a new write-only instance of ZxrfqV3, bound to a specific deployed contract.
func NewZxrfqV3Transactor(address common.Address, transactor bind.ContractTransactor) (*ZxrfqV3Transactor, error) {
	contract, err := bindZxrfqV3(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZxrfqV3Transactor{contract: contract}, nil
}

// NewZxrfqV3Filterer creates a new log filterer instance of ZxrfqV3, bound to a specific deployed contract.
func NewZxrfqV3Filterer(address common.Address, filterer bind.ContractFilterer) (*ZxrfqV3Filterer, error) {
	contract, err := bindZxrfqV3(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZxrfqV3Filterer{contract: contract}, nil
}

// bindZxrfqV3 binds a generic wrapper to an already deployed contract.
func bindZxrfqV3(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ZxrfqV3MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZxrfqV3 *ZxrfqV3Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZxrfqV3.Contract.ZxrfqV3Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZxrfqV3 *ZxrfqV3Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZxrfqV3.Contract.ZxrfqV3Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZxrfqV3 *ZxrfqV3Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZxrfqV3.Contract.ZxrfqV3Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZxrfqV3 *ZxrfqV3CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZxrfqV3.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZxrfqV3 *ZxrfqV3TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZxrfqV3.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZxrfqV3 *ZxrfqV3TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZxrfqV3.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) pure returns()
func (_ZxrfqV3 *ZxrfqV3Caller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) error {
	var out []interface{}
	err := _ZxrfqV3.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return err
	}

	return err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) pure returns()
func (_ZxrfqV3 *ZxrfqV3Session) BalanceOf(arg0 common.Address) error {
	return _ZxrfqV3.Contract.BalanceOf(&_ZxrfqV3.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) pure returns()
func (_ZxrfqV3 *ZxrfqV3CallerSession) BalanceOf(arg0 common.Address) error {
	return _ZxrfqV3.Contract.BalanceOf(&_ZxrfqV3.CallOpts, arg0)
}

// Execute is a paid mutator transaction binding the contract method 0xe4b11b29.
//
// Solidity: function execute(bytes[] actions, (address,address,uint256) slippage) payable returns()
func (_ZxrfqV3 *ZxrfqV3Transactor) Execute(opts *bind.TransactOpts, actions [][]byte, slippage SettlerAllowedSlippage) (*types.Transaction, error) {
	return _ZxrfqV3.contract.Transact(opts, "execute", actions, slippage)
}

// Execute is a paid mutator transaction binding the contract method 0xe4b11b29.
//
// Solidity: function execute(bytes[] actions, (address,address,uint256) slippage) payable returns()
func (_ZxrfqV3 *ZxrfqV3Session) Execute(actions [][]byte, slippage SettlerAllowedSlippage) (*types.Transaction, error) {
	return _ZxrfqV3.Contract.Execute(&_ZxrfqV3.TransactOpts, actions, slippage)
}

// Execute is a paid mutator transaction binding the contract method 0xe4b11b29.
//
// Solidity: function execute(bytes[] actions, (address,address,uint256) slippage) payable returns()
func (_ZxrfqV3 *ZxrfqV3TransactorSession) Execute(actions [][]byte, slippage SettlerAllowedSlippage) (*types.Transaction, error) {
	return _ZxrfqV3.Contract.Execute(&_ZxrfqV3.TransactOpts, actions, slippage)
}

// ExecuteMetaTxn is a paid mutator transaction binding the contract method 0xdd304e43.
//
// Solidity: function executeMetaTxn(bytes[] actions, (address,address,uint256) slippage, address msgSender, bytes sig) returns()
func (_ZxrfqV3 *ZxrfqV3Transactor) ExecuteMetaTxn(opts *bind.TransactOpts, actions [][]byte, slippage SettlerAllowedSlippage, msgSender common.Address, sig []byte) (*types.Transaction, error) {
	return _ZxrfqV3.contract.Transact(opts, "executeMetaTxn", actions, slippage, msgSender, sig)
}

// ExecuteMetaTxn is a paid mutator transaction binding the contract method 0xdd304e43.
//
// Solidity: function executeMetaTxn(bytes[] actions, (address,address,uint256) slippage, address msgSender, bytes sig) returns()
func (_ZxrfqV3 *ZxrfqV3Session) ExecuteMetaTxn(actions [][]byte, slippage SettlerAllowedSlippage, msgSender common.Address, sig []byte) (*types.Transaction, error) {
	return _ZxrfqV3.Contract.ExecuteMetaTxn(&_ZxrfqV3.TransactOpts, actions, slippage, msgSender, sig)
}

// ExecuteMetaTxn is a paid mutator transaction binding the contract method 0xdd304e43.
//
// Solidity: function executeMetaTxn(bytes[] actions, (address,address,uint256) slippage, address msgSender, bytes sig) returns()
func (_ZxrfqV3 *ZxrfqV3TransactorSession) ExecuteMetaTxn(actions [][]byte, slippage SettlerAllowedSlippage, msgSender common.Address, sig []byte) (*types.Transaction, error) {
	return _ZxrfqV3.Contract.ExecuteMetaTxn(&_ZxrfqV3.TransactOpts, actions, slippage, msgSender, sig)
}

// UniswapV3SwapCallback is a paid mutator transaction binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes data) returns()
func (_ZxrfqV3 *ZxrfqV3Transactor) UniswapV3SwapCallback(opts *bind.TransactOpts, amount0Delta *big.Int, amount1Delta *big.Int, data []byte) (*types.Transaction, error) {
	return _ZxrfqV3.contract.Transact(opts, "uniswapV3SwapCallback", amount0Delta, amount1Delta, data)
}

// UniswapV3SwapCallback is a paid mutator transaction binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes data) returns()
func (_ZxrfqV3 *ZxrfqV3Session) UniswapV3SwapCallback(amount0Delta *big.Int, amount1Delta *big.Int, data []byte) (*types.Transaction, error) {
	return _ZxrfqV3.Contract.UniswapV3SwapCallback(&_ZxrfqV3.TransactOpts, amount0Delta, amount1Delta, data)
}

// UniswapV3SwapCallback is a paid mutator transaction binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes data) returns()
func (_ZxrfqV3 *ZxrfqV3TransactorSession) UniswapV3SwapCallback(amount0Delta *big.Int, amount1Delta *big.Int, data []byte) (*types.Transaction, error) {
	return _ZxrfqV3.Contract.UniswapV3SwapCallback(&_ZxrfqV3.TransactOpts, amount0Delta, amount1Delta, data)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ZxrfqV3 *ZxrfqV3Transactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZxrfqV3.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ZxrfqV3 *ZxrfqV3Session) Receive() (*types.Transaction, error) {
	return _ZxrfqV3.Contract.Receive(&_ZxrfqV3.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ZxrfqV3 *ZxrfqV3TransactorSession) Receive() (*types.Transaction, error) {
	return _ZxrfqV3.Contract.Receive(&_ZxrfqV3.TransactOpts)
}
