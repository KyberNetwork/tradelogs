// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package cowtransferparser

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

// TransferMetaData contains all meta data concerning the Transfer contract.
var TransferMetaData = &bind.MetaData{
	ABI: "[{\"name\":\"Transfer\",\"inputs\":[{\"type\":\"address\",\"name\":\"_from\",\"indexed\":true},{\"type\":\"address\",\"name\":\"_to\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"_value\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"Approval\",\"inputs\":[{\"type\":\"address\",\"name\":\"_owner\",\"indexed\":true},{\"type\":\"address\",\"name\":\"_spender\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"_value\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"UpdateMiningParameters\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"time\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"rate\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"supply\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"SetMinter\",\"inputs\":[{\"type\":\"address\",\"name\":\"minter\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"SetAdmin\",\"inputs\":[{\"type\":\"address\",\"name\":\"admin\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"outputs\":[],\"inputs\":[{\"type\":\"string\",\"name\":\"_name\"},{\"type\":\"string\",\"name\":\"_symbol\"},{\"type\":\"uint256\",\"name\":\"_decimals\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"name\":\"update_mining_parameters\",\"outputs\":[],\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":148748},{\"name\":\"start_epoch_time_write\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":149603},{\"name\":\"future_epoch_time_write\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":149806},{\"name\":\"available_supply\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":4018},{\"name\":\"mintable_in_timeframe\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"start\"},{\"type\":\"uint256\",\"name\":\"end\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2216141},{\"name\":\"set_minter\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"_minter\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":38698},{\"name\":\"set_admin\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"_admin\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":37837},{\"name\":\"totalSupply\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1421},{\"name\":\"allowance\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"_owner\"},{\"type\":\"address\",\"name\":\"_spender\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1759},{\"name\":\"transfer\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"_to\"},{\"type\":\"uint256\",\"name\":\"_value\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":75139},{\"name\":\"transferFrom\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"_from\"},{\"type\":\"address\",\"name\":\"_to\"},{\"type\":\"uint256\",\"name\":\"_value\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":111433},{\"name\":\"approve\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"_spender\"},{\"type\":\"uint256\",\"name\":\"_value\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":39288},{\"name\":\"mint\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"_to\"},{\"type\":\"uint256\",\"name\":\"_value\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":228030},{\"name\":\"burn\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"_value\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":74999},{\"name\":\"set_name\",\"outputs\":[],\"inputs\":[{\"type\":\"string\",\"name\":\"_name\"},{\"type\":\"string\",\"name\":\"_symbol\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":178270},{\"name\":\"name\",\"outputs\":[{\"type\":\"string\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":8063},{\"name\":\"symbol\",\"outputs\":[{\"type\":\"string\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":7116},{\"name\":\"decimals\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1721},{\"name\":\"balanceOf\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"arg0\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1905},{\"name\":\"minter\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1781},{\"name\":\"admin\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1811},{\"name\":\"mining_epoch\",\"outputs\":[{\"type\":\"int128\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1841},{\"name\":\"start_epoch_time\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1871},{\"name\":\"rate\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1901}]",
}

// TransferABI is the input ABI used to generate the binding from.
// Deprecated: Use TransferMetaData.ABI instead.
var TransferABI = TransferMetaData.ABI

// Transfer is an auto generated Go binding around an Ethereum contract.
type Transfer struct {
	TransferCaller     // Read-only binding to the contract
	TransferTransactor // Write-only binding to the contract
	TransferFilterer   // Log filterer for contract events
}

// TransferCaller is an auto generated read-only Go binding around an Ethereum contract.
type TransferCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransferTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TransferTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransferFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TransferFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransferSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TransferSession struct {
	Contract     *Transfer         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TransferCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TransferCallerSession struct {
	Contract *TransferCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// TransferTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TransferTransactorSession struct {
	Contract     *TransferTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// TransferRaw is an auto generated low-level Go binding around an Ethereum contract.
type TransferRaw struct {
	Contract *Transfer // Generic contract binding to access the raw methods on
}

// TransferCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TransferCallerRaw struct {
	Contract *TransferCaller // Generic read-only contract binding to access the raw methods on
}

// TransferTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TransferTransactorRaw struct {
	Contract *TransferTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTransfer creates a new instance of Transfer, bound to a specific deployed contract.
func NewTransfer(address common.Address, backend bind.ContractBackend) (*Transfer, error) {
	contract, err := bindTransfer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Transfer{TransferCaller: TransferCaller{contract: contract}, TransferTransactor: TransferTransactor{contract: contract}, TransferFilterer: TransferFilterer{contract: contract}}, nil
}

// NewTransferCaller creates a new read-only instance of Transfer, bound to a specific deployed contract.
func NewTransferCaller(address common.Address, caller bind.ContractCaller) (*TransferCaller, error) {
	contract, err := bindTransfer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TransferCaller{contract: contract}, nil
}

// NewTransferTransactor creates a new write-only instance of Transfer, bound to a specific deployed contract.
func NewTransferTransactor(address common.Address, transactor bind.ContractTransactor) (*TransferTransactor, error) {
	contract, err := bindTransfer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TransferTransactor{contract: contract}, nil
}

// NewTransferFilterer creates a new log filterer instance of Transfer, bound to a specific deployed contract.
func NewTransferFilterer(address common.Address, filterer bind.ContractFilterer) (*TransferFilterer, error) {
	contract, err := bindTransfer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TransferFilterer{contract: contract}, nil
}

// bindTransfer binds a generic wrapper to an already deployed contract.
func bindTransfer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TransferMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Transfer *TransferRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Transfer.Contract.TransferCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Transfer *TransferRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Transfer.Contract.TransferTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Transfer *TransferRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Transfer.Contract.TransferTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Transfer *TransferCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Transfer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Transfer *TransferTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Transfer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Transfer *TransferTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Transfer.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Transfer *TransferCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Transfer.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Transfer *TransferSession) Admin() (common.Address, error) {
	return _Transfer.Contract.Admin(&_Transfer.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Transfer *TransferCallerSession) Admin() (common.Address, error) {
	return _Transfer.Contract.Admin(&_Transfer.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_Transfer *TransferCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Transfer.contract.Call(opts, &out, "allowance", _owner, _spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_Transfer *TransferSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _Transfer.Contract.Allowance(&_Transfer.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_Transfer *TransferCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _Transfer.Contract.Allowance(&_Transfer.CallOpts, _owner, _spender)
}

// AvailableSupply is a free data retrieval call binding the contract method 0x24f92a25.
//
// Solidity: function available_supply() view returns(uint256)
func (_Transfer *TransferCaller) AvailableSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Transfer.contract.Call(opts, &out, "available_supply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AvailableSupply is a free data retrieval call binding the contract method 0x24f92a25.
//
// Solidity: function available_supply() view returns(uint256)
func (_Transfer *TransferSession) AvailableSupply() (*big.Int, error) {
	return _Transfer.Contract.AvailableSupply(&_Transfer.CallOpts)
}

// AvailableSupply is a free data retrieval call binding the contract method 0x24f92a25.
//
// Solidity: function available_supply() view returns(uint256)
func (_Transfer *TransferCallerSession) AvailableSupply() (*big.Int, error) {
	return _Transfer.Contract.AvailableSupply(&_Transfer.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address arg0) view returns(uint256)
func (_Transfer *TransferCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Transfer.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address arg0) view returns(uint256)
func (_Transfer *TransferSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _Transfer.Contract.BalanceOf(&_Transfer.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address arg0) view returns(uint256)
func (_Transfer *TransferCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _Transfer.Contract.BalanceOf(&_Transfer.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_Transfer *TransferCaller) Decimals(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Transfer.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_Transfer *TransferSession) Decimals() (*big.Int, error) {
	return _Transfer.Contract.Decimals(&_Transfer.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_Transfer *TransferCallerSession) Decimals() (*big.Int, error) {
	return _Transfer.Contract.Decimals(&_Transfer.CallOpts)
}

// MiningEpoch is a free data retrieval call binding the contract method 0xf9a40bf6.
//
// Solidity: function mining_epoch() view returns(int128)
func (_Transfer *TransferCaller) MiningEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Transfer.contract.Call(opts, &out, "mining_epoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MiningEpoch is a free data retrieval call binding the contract method 0xf9a40bf6.
//
// Solidity: function mining_epoch() view returns(int128)
func (_Transfer *TransferSession) MiningEpoch() (*big.Int, error) {
	return _Transfer.Contract.MiningEpoch(&_Transfer.CallOpts)
}

// MiningEpoch is a free data retrieval call binding the contract method 0xf9a40bf6.
//
// Solidity: function mining_epoch() view returns(int128)
func (_Transfer *TransferCallerSession) MiningEpoch() (*big.Int, error) {
	return _Transfer.Contract.MiningEpoch(&_Transfer.CallOpts)
}

// MintableInTimeframe is a free data retrieval call binding the contract method 0xd725a9ca.
//
// Solidity: function mintable_in_timeframe(uint256 start, uint256 end) view returns(uint256)
func (_Transfer *TransferCaller) MintableInTimeframe(opts *bind.CallOpts, start *big.Int, end *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Transfer.contract.Call(opts, &out, "mintable_in_timeframe", start, end)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MintableInTimeframe is a free data retrieval call binding the contract method 0xd725a9ca.
//
// Solidity: function mintable_in_timeframe(uint256 start, uint256 end) view returns(uint256)
func (_Transfer *TransferSession) MintableInTimeframe(start *big.Int, end *big.Int) (*big.Int, error) {
	return _Transfer.Contract.MintableInTimeframe(&_Transfer.CallOpts, start, end)
}

// MintableInTimeframe is a free data retrieval call binding the contract method 0xd725a9ca.
//
// Solidity: function mintable_in_timeframe(uint256 start, uint256 end) view returns(uint256)
func (_Transfer *TransferCallerSession) MintableInTimeframe(start *big.Int, end *big.Int) (*big.Int, error) {
	return _Transfer.Contract.MintableInTimeframe(&_Transfer.CallOpts, start, end)
}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_Transfer *TransferCaller) Minter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Transfer.contract.Call(opts, &out, "minter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_Transfer *TransferSession) Minter() (common.Address, error) {
	return _Transfer.Contract.Minter(&_Transfer.CallOpts)
}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_Transfer *TransferCallerSession) Minter() (common.Address, error) {
	return _Transfer.Contract.Minter(&_Transfer.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Transfer *TransferCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Transfer.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Transfer *TransferSession) Name() (string, error) {
	return _Transfer.Contract.Name(&_Transfer.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Transfer *TransferCallerSession) Name() (string, error) {
	return _Transfer.Contract.Name(&_Transfer.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() view returns(uint256)
func (_Transfer *TransferCaller) Rate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Transfer.contract.Call(opts, &out, "rate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() view returns(uint256)
func (_Transfer *TransferSession) Rate() (*big.Int, error) {
	return _Transfer.Contract.Rate(&_Transfer.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() view returns(uint256)
func (_Transfer *TransferCallerSession) Rate() (*big.Int, error) {
	return _Transfer.Contract.Rate(&_Transfer.CallOpts)
}

// StartEpochTime is a free data retrieval call binding the contract method 0x7375be26.
//
// Solidity: function start_epoch_time() view returns(uint256)
func (_Transfer *TransferCaller) StartEpochTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Transfer.contract.Call(opts, &out, "start_epoch_time")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StartEpochTime is a free data retrieval call binding the contract method 0x7375be26.
//
// Solidity: function start_epoch_time() view returns(uint256)
func (_Transfer *TransferSession) StartEpochTime() (*big.Int, error) {
	return _Transfer.Contract.StartEpochTime(&_Transfer.CallOpts)
}

// StartEpochTime is a free data retrieval call binding the contract method 0x7375be26.
//
// Solidity: function start_epoch_time() view returns(uint256)
func (_Transfer *TransferCallerSession) StartEpochTime() (*big.Int, error) {
	return _Transfer.Contract.StartEpochTime(&_Transfer.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Transfer *TransferCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Transfer.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Transfer *TransferSession) Symbol() (string, error) {
	return _Transfer.Contract.Symbol(&_Transfer.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Transfer *TransferCallerSession) Symbol() (string, error) {
	return _Transfer.Contract.Symbol(&_Transfer.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Transfer *TransferCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Transfer.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Transfer *TransferSession) TotalSupply() (*big.Int, error) {
	return _Transfer.Contract.TotalSupply(&_Transfer.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Transfer *TransferCallerSession) TotalSupply() (*big.Int, error) {
	return _Transfer.Contract.TotalSupply(&_Transfer.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_Transfer *TransferTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Transfer.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_Transfer *TransferSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Transfer.Contract.Approve(&_Transfer.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_Transfer *TransferTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Transfer.Contract.Approve(&_Transfer.TransactOpts, _spender, _value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 _value) returns(bool)
func (_Transfer *TransferTransactor) Burn(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _Transfer.contract.Transact(opts, "burn", _value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 _value) returns(bool)
func (_Transfer *TransferSession) Burn(_value *big.Int) (*types.Transaction, error) {
	return _Transfer.Contract.Burn(&_Transfer.TransactOpts, _value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 _value) returns(bool)
func (_Transfer *TransferTransactorSession) Burn(_value *big.Int) (*types.Transaction, error) {
	return _Transfer.Contract.Burn(&_Transfer.TransactOpts, _value)
}

// FutureEpochTimeWrite is a paid mutator transaction binding the contract method 0xb26b238e.
//
// Solidity: function future_epoch_time_write() returns(uint256)
func (_Transfer *TransferTransactor) FutureEpochTimeWrite(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Transfer.contract.Transact(opts, "future_epoch_time_write")
}

// FutureEpochTimeWrite is a paid mutator transaction binding the contract method 0xb26b238e.
//
// Solidity: function future_epoch_time_write() returns(uint256)
func (_Transfer *TransferSession) FutureEpochTimeWrite() (*types.Transaction, error) {
	return _Transfer.Contract.FutureEpochTimeWrite(&_Transfer.TransactOpts)
}

// FutureEpochTimeWrite is a paid mutator transaction binding the contract method 0xb26b238e.
//
// Solidity: function future_epoch_time_write() returns(uint256)
func (_Transfer *TransferTransactorSession) FutureEpochTimeWrite() (*types.Transaction, error) {
	return _Transfer.Contract.FutureEpochTimeWrite(&_Transfer.TransactOpts)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _value) returns(bool)
func (_Transfer *TransferTransactor) Mint(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Transfer.contract.Transact(opts, "mint", _to, _value)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _value) returns(bool)
func (_Transfer *TransferSession) Mint(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Transfer.Contract.Mint(&_Transfer.TransactOpts, _to, _value)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _value) returns(bool)
func (_Transfer *TransferTransactorSession) Mint(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Transfer.Contract.Mint(&_Transfer.TransactOpts, _to, _value)
}

// SetAdmin is a paid mutator transaction binding the contract method 0xe9333fab.
//
// Solidity: function set_admin(address _admin) returns()
func (_Transfer *TransferTransactor) SetAdmin(opts *bind.TransactOpts, _admin common.Address) (*types.Transaction, error) {
	return _Transfer.contract.Transact(opts, "set_admin", _admin)
}

// SetAdmin is a paid mutator transaction binding the contract method 0xe9333fab.
//
// Solidity: function set_admin(address _admin) returns()
func (_Transfer *TransferSession) SetAdmin(_admin common.Address) (*types.Transaction, error) {
	return _Transfer.Contract.SetAdmin(&_Transfer.TransactOpts, _admin)
}

// SetAdmin is a paid mutator transaction binding the contract method 0xe9333fab.
//
// Solidity: function set_admin(address _admin) returns()
func (_Transfer *TransferTransactorSession) SetAdmin(_admin common.Address) (*types.Transaction, error) {
	return _Transfer.Contract.SetAdmin(&_Transfer.TransactOpts, _admin)
}

// SetMinter is a paid mutator transaction binding the contract method 0x1652e9fc.
//
// Solidity: function set_minter(address _minter) returns()
func (_Transfer *TransferTransactor) SetMinter(opts *bind.TransactOpts, _minter common.Address) (*types.Transaction, error) {
	return _Transfer.contract.Transact(opts, "set_minter", _minter)
}

// SetMinter is a paid mutator transaction binding the contract method 0x1652e9fc.
//
// Solidity: function set_minter(address _minter) returns()
func (_Transfer *TransferSession) SetMinter(_minter common.Address) (*types.Transaction, error) {
	return _Transfer.Contract.SetMinter(&_Transfer.TransactOpts, _minter)
}

// SetMinter is a paid mutator transaction binding the contract method 0x1652e9fc.
//
// Solidity: function set_minter(address _minter) returns()
func (_Transfer *TransferTransactorSession) SetMinter(_minter common.Address) (*types.Transaction, error) {
	return _Transfer.Contract.SetMinter(&_Transfer.TransactOpts, _minter)
}

// SetName is a paid mutator transaction binding the contract method 0xe1430e06.
//
// Solidity: function set_name(string _name, string _symbol) returns()
func (_Transfer *TransferTransactor) SetName(opts *bind.TransactOpts, _name string, _symbol string) (*types.Transaction, error) {
	return _Transfer.contract.Transact(opts, "set_name", _name, _symbol)
}

// SetName is a paid mutator transaction binding the contract method 0xe1430e06.
//
// Solidity: function set_name(string _name, string _symbol) returns()
func (_Transfer *TransferSession) SetName(_name string, _symbol string) (*types.Transaction, error) {
	return _Transfer.Contract.SetName(&_Transfer.TransactOpts, _name, _symbol)
}

// SetName is a paid mutator transaction binding the contract method 0xe1430e06.
//
// Solidity: function set_name(string _name, string _symbol) returns()
func (_Transfer *TransferTransactorSession) SetName(_name string, _symbol string) (*types.Transaction, error) {
	return _Transfer.Contract.SetName(&_Transfer.TransactOpts, _name, _symbol)
}

// StartEpochTimeWrite is a paid mutator transaction binding the contract method 0xadc4cf43.
//
// Solidity: function start_epoch_time_write() returns(uint256)
func (_Transfer *TransferTransactor) StartEpochTimeWrite(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Transfer.contract.Transact(opts, "start_epoch_time_write")
}

// StartEpochTimeWrite is a paid mutator transaction binding the contract method 0xadc4cf43.
//
// Solidity: function start_epoch_time_write() returns(uint256)
func (_Transfer *TransferSession) StartEpochTimeWrite() (*types.Transaction, error) {
	return _Transfer.Contract.StartEpochTimeWrite(&_Transfer.TransactOpts)
}

// StartEpochTimeWrite is a paid mutator transaction binding the contract method 0xadc4cf43.
//
// Solidity: function start_epoch_time_write() returns(uint256)
func (_Transfer *TransferTransactorSession) StartEpochTimeWrite() (*types.Transaction, error) {
	return _Transfer.Contract.StartEpochTimeWrite(&_Transfer.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_Transfer *TransferTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Transfer.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_Transfer *TransferSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Transfer.Contract.Transfer(&_Transfer.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_Transfer *TransferTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Transfer.Contract.Transfer(&_Transfer.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_Transfer *TransferTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Transfer.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_Transfer *TransferSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Transfer.Contract.TransferFrom(&_Transfer.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_Transfer *TransferTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Transfer.Contract.TransferFrom(&_Transfer.TransactOpts, _from, _to, _value)
}

// UpdateMiningParameters is a paid mutator transaction binding the contract method 0xd43b40fa.
//
// Solidity: function update_mining_parameters() returns()
func (_Transfer *TransferTransactor) UpdateMiningParameters(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Transfer.contract.Transact(opts, "update_mining_parameters")
}

// UpdateMiningParameters is a paid mutator transaction binding the contract method 0xd43b40fa.
//
// Solidity: function update_mining_parameters() returns()
func (_Transfer *TransferSession) UpdateMiningParameters() (*types.Transaction, error) {
	return _Transfer.Contract.UpdateMiningParameters(&_Transfer.TransactOpts)
}

// UpdateMiningParameters is a paid mutator transaction binding the contract method 0xd43b40fa.
//
// Solidity: function update_mining_parameters() returns()
func (_Transfer *TransferTransactorSession) UpdateMiningParameters() (*types.Transaction, error) {
	return _Transfer.Contract.UpdateMiningParameters(&_Transfer.TransactOpts)
}

// TransferApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Transfer contract.
type TransferApprovalIterator struct {
	Event *TransferApproval // Event containing the contract specifics and raw log

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
func (it *TransferApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TransferApproval)
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
		it.Event = new(TransferApproval)
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
func (it *TransferApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TransferApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TransferApproval represents a Approval event raised by the Transfer contract.
type TransferApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _spender, uint256 _value)
func (_Transfer *TransferFilterer) FilterApproval(opts *bind.FilterOpts, _owner []common.Address, _spender []common.Address) (*TransferApprovalIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _Transfer.contract.FilterLogs(opts, "Approval", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return &TransferApprovalIterator{contract: _Transfer.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _spender, uint256 _value)
func (_Transfer *TransferFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *TransferApproval, _owner []common.Address, _spender []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _Transfer.contract.WatchLogs(opts, "Approval", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TransferApproval)
				if err := _Transfer.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _spender, uint256 _value)
func (_Transfer *TransferFilterer) ParseApproval(log types.Log) (*TransferApproval, error) {
	event := new(TransferApproval)
	if err := _Transfer.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TransferSetAdminIterator is returned from FilterSetAdmin and is used to iterate over the raw logs and unpacked data for SetAdmin events raised by the Transfer contract.
type TransferSetAdminIterator struct {
	Event *TransferSetAdmin // Event containing the contract specifics and raw log

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
func (it *TransferSetAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TransferSetAdmin)
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
		it.Event = new(TransferSetAdmin)
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
func (it *TransferSetAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TransferSetAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TransferSetAdmin represents a SetAdmin event raised by the Transfer contract.
type TransferSetAdmin struct {
	Admin common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterSetAdmin is a free log retrieval operation binding the contract event 0x5a272403b402d892977df56625f4164ccaf70ca3863991c43ecfe76a6905b0a1.
//
// Solidity: event SetAdmin(address admin)
func (_Transfer *TransferFilterer) FilterSetAdmin(opts *bind.FilterOpts) (*TransferSetAdminIterator, error) {

	logs, sub, err := _Transfer.contract.FilterLogs(opts, "SetAdmin")
	if err != nil {
		return nil, err
	}
	return &TransferSetAdminIterator{contract: _Transfer.contract, event: "SetAdmin", logs: logs, sub: sub}, nil
}

// WatchSetAdmin is a free log subscription operation binding the contract event 0x5a272403b402d892977df56625f4164ccaf70ca3863991c43ecfe76a6905b0a1.
//
// Solidity: event SetAdmin(address admin)
func (_Transfer *TransferFilterer) WatchSetAdmin(opts *bind.WatchOpts, sink chan<- *TransferSetAdmin) (event.Subscription, error) {

	logs, sub, err := _Transfer.contract.WatchLogs(opts, "SetAdmin")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TransferSetAdmin)
				if err := _Transfer.contract.UnpackLog(event, "SetAdmin", log); err != nil {
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

// ParseSetAdmin is a log parse operation binding the contract event 0x5a272403b402d892977df56625f4164ccaf70ca3863991c43ecfe76a6905b0a1.
//
// Solidity: event SetAdmin(address admin)
func (_Transfer *TransferFilterer) ParseSetAdmin(log types.Log) (*TransferSetAdmin, error) {
	event := new(TransferSetAdmin)
	if err := _Transfer.contract.UnpackLog(event, "SetAdmin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TransferSetMinterIterator is returned from FilterSetMinter and is used to iterate over the raw logs and unpacked data for SetMinter events raised by the Transfer contract.
type TransferSetMinterIterator struct {
	Event *TransferSetMinter // Event containing the contract specifics and raw log

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
func (it *TransferSetMinterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TransferSetMinter)
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
		it.Event = new(TransferSetMinter)
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
func (it *TransferSetMinterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TransferSetMinterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TransferSetMinter represents a SetMinter event raised by the Transfer contract.
type TransferSetMinter struct {
	Minter common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetMinter is a free log retrieval operation binding the contract event 0xcec52196e972044edde8689a1b608e459c5946b7f3e5c8cd3d6d8e126d422e1c.
//
// Solidity: event SetMinter(address minter)
func (_Transfer *TransferFilterer) FilterSetMinter(opts *bind.FilterOpts) (*TransferSetMinterIterator, error) {

	logs, sub, err := _Transfer.contract.FilterLogs(opts, "SetMinter")
	if err != nil {
		return nil, err
	}
	return &TransferSetMinterIterator{contract: _Transfer.contract, event: "SetMinter", logs: logs, sub: sub}, nil
}

// WatchSetMinter is a free log subscription operation binding the contract event 0xcec52196e972044edde8689a1b608e459c5946b7f3e5c8cd3d6d8e126d422e1c.
//
// Solidity: event SetMinter(address minter)
func (_Transfer *TransferFilterer) WatchSetMinter(opts *bind.WatchOpts, sink chan<- *TransferSetMinter) (event.Subscription, error) {

	logs, sub, err := _Transfer.contract.WatchLogs(opts, "SetMinter")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TransferSetMinter)
				if err := _Transfer.contract.UnpackLog(event, "SetMinter", log); err != nil {
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

// ParseSetMinter is a log parse operation binding the contract event 0xcec52196e972044edde8689a1b608e459c5946b7f3e5c8cd3d6d8e126d422e1c.
//
// Solidity: event SetMinter(address minter)
func (_Transfer *TransferFilterer) ParseSetMinter(log types.Log) (*TransferSetMinter, error) {
	event := new(TransferSetMinter)
	if err := _Transfer.contract.UnpackLog(event, "SetMinter", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TransferTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Transfer contract.
type TransferTransferIterator struct {
	Event *TransferTransfer // Event containing the contract specifics and raw log

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
func (it *TransferTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TransferTransfer)
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
		it.Event = new(TransferTransfer)
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
func (it *TransferTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TransferTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TransferTransfer represents a Transfer event raised by the Transfer contract.
type TransferTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _value)
func (_Transfer *TransferFilterer) FilterTransfer(opts *bind.FilterOpts, _from []common.Address, _to []common.Address) (*TransferTransferIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _Transfer.contract.FilterLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &TransferTransferIterator{contract: _Transfer.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _value)
func (_Transfer *TransferFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *TransferTransfer, _from []common.Address, _to []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _Transfer.contract.WatchLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TransferTransfer)
				if err := _Transfer.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _value)
func (_Transfer *TransferFilterer) ParseTransfer(log types.Log) (*TransferTransfer, error) {
	event := new(TransferTransfer)
	if err := _Transfer.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TransferUpdateMiningParametersIterator is returned from FilterUpdateMiningParameters and is used to iterate over the raw logs and unpacked data for UpdateMiningParameters events raised by the Transfer contract.
type TransferUpdateMiningParametersIterator struct {
	Event *TransferUpdateMiningParameters // Event containing the contract specifics and raw log

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
func (it *TransferUpdateMiningParametersIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TransferUpdateMiningParameters)
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
		it.Event = new(TransferUpdateMiningParameters)
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
func (it *TransferUpdateMiningParametersIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TransferUpdateMiningParametersIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TransferUpdateMiningParameters represents a UpdateMiningParameters event raised by the Transfer contract.
type TransferUpdateMiningParameters struct {
	Time   *big.Int
	Rate   *big.Int
	Supply *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterUpdateMiningParameters is a free log retrieval operation binding the contract event 0x27e46362a1e6129b6dd539c984ce739291a97128dfcaeca1255e8ac83abd9441.
//
// Solidity: event UpdateMiningParameters(uint256 time, uint256 rate, uint256 supply)
func (_Transfer *TransferFilterer) FilterUpdateMiningParameters(opts *bind.FilterOpts) (*TransferUpdateMiningParametersIterator, error) {

	logs, sub, err := _Transfer.contract.FilterLogs(opts, "UpdateMiningParameters")
	if err != nil {
		return nil, err
	}
	return &TransferUpdateMiningParametersIterator{contract: _Transfer.contract, event: "UpdateMiningParameters", logs: logs, sub: sub}, nil
}

// WatchUpdateMiningParameters is a free log subscription operation binding the contract event 0x27e46362a1e6129b6dd539c984ce739291a97128dfcaeca1255e8ac83abd9441.
//
// Solidity: event UpdateMiningParameters(uint256 time, uint256 rate, uint256 supply)
func (_Transfer *TransferFilterer) WatchUpdateMiningParameters(opts *bind.WatchOpts, sink chan<- *TransferUpdateMiningParameters) (event.Subscription, error) {

	logs, sub, err := _Transfer.contract.WatchLogs(opts, "UpdateMiningParameters")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TransferUpdateMiningParameters)
				if err := _Transfer.contract.UnpackLog(event, "UpdateMiningParameters", log); err != nil {
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

// ParseUpdateMiningParameters is a log parse operation binding the contract event 0x27e46362a1e6129b6dd539c984ce739291a97128dfcaeca1255e8ac83abd9441.
//
// Solidity: event UpdateMiningParameters(uint256 time, uint256 rate, uint256 supply)
func (_Transfer *TransferFilterer) ParseUpdateMiningParameters(log types.Log) (*TransferUpdateMiningParameters, error) {
	event := new(TransferUpdateMiningParameters)
	if err := _Transfer.contract.UnpackLog(event, "UpdateMiningParameters", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
