// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package erc20transfer

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

// ERC20TransferMetaData contains all meta data concerning the ERC20Transfer contract.
var ERC20TransferMetaData = &bind.MetaData{
	ABI: "[{\"name\":\"Transfer\",\"inputs\":[{\"type\":\"address\",\"name\":\"_from\",\"indexed\":true},{\"type\":\"address\",\"name\":\"_to\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"_value\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"Approval\",\"inputs\":[{\"type\":\"address\",\"name\":\"_owner\",\"indexed\":true},{\"type\":\"address\",\"name\":\"_spender\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"_value\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"UpdateMiningParameters\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"time\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"rate\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"supply\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"SetMinter\",\"inputs\":[{\"type\":\"address\",\"name\":\"minter\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"SetAdmin\",\"inputs\":[{\"type\":\"address\",\"name\":\"admin\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"outputs\":[],\"inputs\":[{\"type\":\"string\",\"name\":\"_name\"},{\"type\":\"string\",\"name\":\"_symbol\"},{\"type\":\"uint256\",\"name\":\"_decimals\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"name\":\"update_mining_parameters\",\"outputs\":[],\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":148748},{\"name\":\"start_epoch_time_write\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":149603},{\"name\":\"future_epoch_time_write\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":149806},{\"name\":\"available_supply\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":4018},{\"name\":\"mintable_in_timeframe\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"start\"},{\"type\":\"uint256\",\"name\":\"end\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2216141},{\"name\":\"set_minter\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"_minter\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":38698},{\"name\":\"set_admin\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"_admin\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":37837},{\"name\":\"totalSupply\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1421},{\"name\":\"allowance\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"_owner\"},{\"type\":\"address\",\"name\":\"_spender\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1759},{\"name\":\"transfer\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"_to\"},{\"type\":\"uint256\",\"name\":\"_value\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":75139},{\"name\":\"transferFrom\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"_from\"},{\"type\":\"address\",\"name\":\"_to\"},{\"type\":\"uint256\",\"name\":\"_value\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":111433},{\"name\":\"approve\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"_spender\"},{\"type\":\"uint256\",\"name\":\"_value\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":39288},{\"name\":\"mint\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"_to\"},{\"type\":\"uint256\",\"name\":\"_value\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":228030},{\"name\":\"burn\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"_value\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":74999},{\"name\":\"set_name\",\"outputs\":[],\"inputs\":[{\"type\":\"string\",\"name\":\"_name\"},{\"type\":\"string\",\"name\":\"_symbol\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":178270},{\"name\":\"name\",\"outputs\":[{\"type\":\"string\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":8063},{\"name\":\"symbol\",\"outputs\":[{\"type\":\"string\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":7116},{\"name\":\"decimals\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1721},{\"name\":\"balanceOf\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"arg0\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1905},{\"name\":\"minter\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1781},{\"name\":\"admin\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1811},{\"name\":\"mining_epoch\",\"outputs\":[{\"type\":\"int128\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1841},{\"name\":\"start_epoch_time\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1871},{\"name\":\"rate\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1901}]",
}

// ERC20TransferABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC20TransferMetaData.ABI instead.
var ERC20TransferABI = ERC20TransferMetaData.ABI

// ERC20Transfer is an auto generated Go binding around an Ethereum contract.
type ERC20Transfer struct {
	ERC20TransferCaller     // Read-only binding to the contract
	ERC20TransferTransactor // Write-only binding to the contract
	ERC20TransferFilterer   // Log filterer for contract events
}

// ERC20TransferCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20TransferCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20TransferTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20TransferTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20TransferFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20TransferFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20TransferSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20TransferSession struct {
	Contract     *ERC20Transfer    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20TransferCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20TransferCallerSession struct {
	Contract *ERC20TransferCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// ERC20TransferTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20TransferTransactorSession struct {
	Contract     *ERC20TransferTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ERC20TransferRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20TransferRaw struct {
	Contract *ERC20Transfer // Generic contract binding to access the raw methods on
}

// ERC20TransferCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20TransferCallerRaw struct {
	Contract *ERC20TransferCaller // Generic read-only contract binding to access the raw methods on
}

// ERC20TransferTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20TransferTransactorRaw struct {
	Contract *ERC20TransferTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20Transfer creates a new instance of ERC20Transfer, bound to a specific deployed contract.
func NewERC20Transfer(address common.Address, backend bind.ContractBackend) (*ERC20Transfer, error) {
	contract, err := bindERC20Transfer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20Transfer{ERC20TransferCaller: ERC20TransferCaller{contract: contract}, ERC20TransferTransactor: ERC20TransferTransactor{contract: contract}, ERC20TransferFilterer: ERC20TransferFilterer{contract: contract}}, nil
}

// NewERC20TransferCaller creates a new read-only instance of ERC20Transfer, bound to a specific deployed contract.
func NewERC20TransferCaller(address common.Address, caller bind.ContractCaller) (*ERC20TransferCaller, error) {
	contract, err := bindERC20Transfer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20TransferCaller{contract: contract}, nil
}

// NewERC20TransferTransactor creates a new write-only instance of ERC20Transfer, bound to a specific deployed contract.
func NewERC20TransferTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC20TransferTransactor, error) {
	contract, err := bindERC20Transfer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20TransferTransactor{contract: contract}, nil
}

// NewERC20TransferFilterer creates a new log filterer instance of ERC20Transfer, bound to a specific deployed contract.
func NewERC20TransferFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC20TransferFilterer, error) {
	contract, err := bindERC20Transfer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20TransferFilterer{contract: contract}, nil
}

// bindERC20Transfer binds a generic wrapper to an already deployed contract.
func bindERC20Transfer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ERC20TransferMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Transfer *ERC20TransferRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20Transfer.Contract.ERC20TransferCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Transfer *ERC20TransferRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Transfer.Contract.ERC20TransferTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Transfer *ERC20TransferRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Transfer.Contract.ERC20TransferTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Transfer *ERC20TransferCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20Transfer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Transfer *ERC20TransferTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Transfer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Transfer *ERC20TransferTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Transfer.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_ERC20Transfer *ERC20TransferCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC20Transfer.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_ERC20Transfer *ERC20TransferSession) Admin() (common.Address, error) {
	return _ERC20Transfer.Contract.Admin(&_ERC20Transfer.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_ERC20Transfer *ERC20TransferCallerSession) Admin() (common.Address, error) {
	return _ERC20Transfer.Contract.Admin(&_ERC20Transfer.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_ERC20Transfer *ERC20TransferCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC20Transfer.contract.Call(opts, &out, "allowance", _owner, _spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_ERC20Transfer *ERC20TransferSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _ERC20Transfer.Contract.Allowance(&_ERC20Transfer.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_ERC20Transfer *ERC20TransferCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _ERC20Transfer.Contract.Allowance(&_ERC20Transfer.CallOpts, _owner, _spender)
}

// AvailableSupply is a free data retrieval call binding the contract method 0x24f92a25.
//
// Solidity: function available_supply() view returns(uint256)
func (_ERC20Transfer *ERC20TransferCaller) AvailableSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20Transfer.contract.Call(opts, &out, "available_supply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AvailableSupply is a free data retrieval call binding the contract method 0x24f92a25.
//
// Solidity: function available_supply() view returns(uint256)
func (_ERC20Transfer *ERC20TransferSession) AvailableSupply() (*big.Int, error) {
	return _ERC20Transfer.Contract.AvailableSupply(&_ERC20Transfer.CallOpts)
}

// AvailableSupply is a free data retrieval call binding the contract method 0x24f92a25.
//
// Solidity: function available_supply() view returns(uint256)
func (_ERC20Transfer *ERC20TransferCallerSession) AvailableSupply() (*big.Int, error) {
	return _ERC20Transfer.Contract.AvailableSupply(&_ERC20Transfer.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address arg0) view returns(uint256)
func (_ERC20Transfer *ERC20TransferCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC20Transfer.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address arg0) view returns(uint256)
func (_ERC20Transfer *ERC20TransferSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _ERC20Transfer.Contract.BalanceOf(&_ERC20Transfer.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address arg0) view returns(uint256)
func (_ERC20Transfer *ERC20TransferCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _ERC20Transfer.Contract.BalanceOf(&_ERC20Transfer.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_ERC20Transfer *ERC20TransferCaller) Decimals(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20Transfer.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_ERC20Transfer *ERC20TransferSession) Decimals() (*big.Int, error) {
	return _ERC20Transfer.Contract.Decimals(&_ERC20Transfer.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_ERC20Transfer *ERC20TransferCallerSession) Decimals() (*big.Int, error) {
	return _ERC20Transfer.Contract.Decimals(&_ERC20Transfer.CallOpts)
}

// MiningEpoch is a free data retrieval call binding the contract method 0xf9a40bf6.
//
// Solidity: function mining_epoch() view returns(int128)
func (_ERC20Transfer *ERC20TransferCaller) MiningEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20Transfer.contract.Call(opts, &out, "mining_epoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MiningEpoch is a free data retrieval call binding the contract method 0xf9a40bf6.
//
// Solidity: function mining_epoch() view returns(int128)
func (_ERC20Transfer *ERC20TransferSession) MiningEpoch() (*big.Int, error) {
	return _ERC20Transfer.Contract.MiningEpoch(&_ERC20Transfer.CallOpts)
}

// MiningEpoch is a free data retrieval call binding the contract method 0xf9a40bf6.
//
// Solidity: function mining_epoch() view returns(int128)
func (_ERC20Transfer *ERC20TransferCallerSession) MiningEpoch() (*big.Int, error) {
	return _ERC20Transfer.Contract.MiningEpoch(&_ERC20Transfer.CallOpts)
}

// MintableInTimeframe is a free data retrieval call binding the contract method 0xd725a9ca.
//
// Solidity: function mintable_in_timeframe(uint256 start, uint256 end) view returns(uint256)
func (_ERC20Transfer *ERC20TransferCaller) MintableInTimeframe(opts *bind.CallOpts, start *big.Int, end *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ERC20Transfer.contract.Call(opts, &out, "mintable_in_timeframe", start, end)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MintableInTimeframe is a free data retrieval call binding the contract method 0xd725a9ca.
//
// Solidity: function mintable_in_timeframe(uint256 start, uint256 end) view returns(uint256)
func (_ERC20Transfer *ERC20TransferSession) MintableInTimeframe(start *big.Int, end *big.Int) (*big.Int, error) {
	return _ERC20Transfer.Contract.MintableInTimeframe(&_ERC20Transfer.CallOpts, start, end)
}

// MintableInTimeframe is a free data retrieval call binding the contract method 0xd725a9ca.
//
// Solidity: function mintable_in_timeframe(uint256 start, uint256 end) view returns(uint256)
func (_ERC20Transfer *ERC20TransferCallerSession) MintableInTimeframe(start *big.Int, end *big.Int) (*big.Int, error) {
	return _ERC20Transfer.Contract.MintableInTimeframe(&_ERC20Transfer.CallOpts, start, end)
}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_ERC20Transfer *ERC20TransferCaller) Minter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC20Transfer.contract.Call(opts, &out, "minter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_ERC20Transfer *ERC20TransferSession) Minter() (common.Address, error) {
	return _ERC20Transfer.Contract.Minter(&_ERC20Transfer.CallOpts)
}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_ERC20Transfer *ERC20TransferCallerSession) Minter() (common.Address, error) {
	return _ERC20Transfer.Contract.Minter(&_ERC20Transfer.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20Transfer *ERC20TransferCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERC20Transfer.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20Transfer *ERC20TransferSession) Name() (string, error) {
	return _ERC20Transfer.Contract.Name(&_ERC20Transfer.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20Transfer *ERC20TransferCallerSession) Name() (string, error) {
	return _ERC20Transfer.Contract.Name(&_ERC20Transfer.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() view returns(uint256)
func (_ERC20Transfer *ERC20TransferCaller) Rate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20Transfer.contract.Call(opts, &out, "rate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() view returns(uint256)
func (_ERC20Transfer *ERC20TransferSession) Rate() (*big.Int, error) {
	return _ERC20Transfer.Contract.Rate(&_ERC20Transfer.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() view returns(uint256)
func (_ERC20Transfer *ERC20TransferCallerSession) Rate() (*big.Int, error) {
	return _ERC20Transfer.Contract.Rate(&_ERC20Transfer.CallOpts)
}

// StartEpochTime is a free data retrieval call binding the contract method 0x7375be26.
//
// Solidity: function start_epoch_time() view returns(uint256)
func (_ERC20Transfer *ERC20TransferCaller) StartEpochTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20Transfer.contract.Call(opts, &out, "start_epoch_time")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StartEpochTime is a free data retrieval call binding the contract method 0x7375be26.
//
// Solidity: function start_epoch_time() view returns(uint256)
func (_ERC20Transfer *ERC20TransferSession) StartEpochTime() (*big.Int, error) {
	return _ERC20Transfer.Contract.StartEpochTime(&_ERC20Transfer.CallOpts)
}

// StartEpochTime is a free data retrieval call binding the contract method 0x7375be26.
//
// Solidity: function start_epoch_time() view returns(uint256)
func (_ERC20Transfer *ERC20TransferCallerSession) StartEpochTime() (*big.Int, error) {
	return _ERC20Transfer.Contract.StartEpochTime(&_ERC20Transfer.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20Transfer *ERC20TransferCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERC20Transfer.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20Transfer *ERC20TransferSession) Symbol() (string, error) {
	return _ERC20Transfer.Contract.Symbol(&_ERC20Transfer.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20Transfer *ERC20TransferCallerSession) Symbol() (string, error) {
	return _ERC20Transfer.Contract.Symbol(&_ERC20Transfer.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20Transfer *ERC20TransferCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20Transfer.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20Transfer *ERC20TransferSession) TotalSupply() (*big.Int, error) {
	return _ERC20Transfer.Contract.TotalSupply(&_ERC20Transfer.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20Transfer *ERC20TransferCallerSession) TotalSupply() (*big.Int, error) {
	return _ERC20Transfer.Contract.TotalSupply(&_ERC20Transfer.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_ERC20Transfer *ERC20TransferTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20Transfer.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_ERC20Transfer *ERC20TransferSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20Transfer.Contract.Approve(&_ERC20Transfer.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_ERC20Transfer *ERC20TransferTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20Transfer.Contract.Approve(&_ERC20Transfer.TransactOpts, _spender, _value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 _value) returns(bool)
func (_ERC20Transfer *ERC20TransferTransactor) Burn(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _ERC20Transfer.contract.Transact(opts, "burn", _value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 _value) returns(bool)
func (_ERC20Transfer *ERC20TransferSession) Burn(_value *big.Int) (*types.Transaction, error) {
	return _ERC20Transfer.Contract.Burn(&_ERC20Transfer.TransactOpts, _value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 _value) returns(bool)
func (_ERC20Transfer *ERC20TransferTransactorSession) Burn(_value *big.Int) (*types.Transaction, error) {
	return _ERC20Transfer.Contract.Burn(&_ERC20Transfer.TransactOpts, _value)
}

// FutureEpochTimeWrite is a paid mutator transaction binding the contract method 0xb26b238e.
//
// Solidity: function future_epoch_time_write() returns(uint256)
func (_ERC20Transfer *ERC20TransferTransactor) FutureEpochTimeWrite(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Transfer.contract.Transact(opts, "future_epoch_time_write")
}

// FutureEpochTimeWrite is a paid mutator transaction binding the contract method 0xb26b238e.
//
// Solidity: function future_epoch_time_write() returns(uint256)
func (_ERC20Transfer *ERC20TransferSession) FutureEpochTimeWrite() (*types.Transaction, error) {
	return _ERC20Transfer.Contract.FutureEpochTimeWrite(&_ERC20Transfer.TransactOpts)
}

// FutureEpochTimeWrite is a paid mutator transaction binding the contract method 0xb26b238e.
//
// Solidity: function future_epoch_time_write() returns(uint256)
func (_ERC20Transfer *ERC20TransferTransactorSession) FutureEpochTimeWrite() (*types.Transaction, error) {
	return _ERC20Transfer.Contract.FutureEpochTimeWrite(&_ERC20Transfer.TransactOpts)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _value) returns(bool)
func (_ERC20Transfer *ERC20TransferTransactor) Mint(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20Transfer.contract.Transact(opts, "mint", _to, _value)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _value) returns(bool)
func (_ERC20Transfer *ERC20TransferSession) Mint(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20Transfer.Contract.Mint(&_ERC20Transfer.TransactOpts, _to, _value)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _value) returns(bool)
func (_ERC20Transfer *ERC20TransferTransactorSession) Mint(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20Transfer.Contract.Mint(&_ERC20Transfer.TransactOpts, _to, _value)
}

// SetAdmin is a paid mutator transaction binding the contract method 0xe9333fab.
//
// Solidity: function set_admin(address _admin) returns()
func (_ERC20Transfer *ERC20TransferTransactor) SetAdmin(opts *bind.TransactOpts, _admin common.Address) (*types.Transaction, error) {
	return _ERC20Transfer.contract.Transact(opts, "set_admin", _admin)
}

// SetAdmin is a paid mutator transaction binding the contract method 0xe9333fab.
//
// Solidity: function set_admin(address _admin) returns()
func (_ERC20Transfer *ERC20TransferSession) SetAdmin(_admin common.Address) (*types.Transaction, error) {
	return _ERC20Transfer.Contract.SetAdmin(&_ERC20Transfer.TransactOpts, _admin)
}

// SetAdmin is a paid mutator transaction binding the contract method 0xe9333fab.
//
// Solidity: function set_admin(address _admin) returns()
func (_ERC20Transfer *ERC20TransferTransactorSession) SetAdmin(_admin common.Address) (*types.Transaction, error) {
	return _ERC20Transfer.Contract.SetAdmin(&_ERC20Transfer.TransactOpts, _admin)
}

// SetMinter is a paid mutator transaction binding the contract method 0x1652e9fc.
//
// Solidity: function set_minter(address _minter) returns()
func (_ERC20Transfer *ERC20TransferTransactor) SetMinter(opts *bind.TransactOpts, _minter common.Address) (*types.Transaction, error) {
	return _ERC20Transfer.contract.Transact(opts, "set_minter", _minter)
}

// SetMinter is a paid mutator transaction binding the contract method 0x1652e9fc.
//
// Solidity: function set_minter(address _minter) returns()
func (_ERC20Transfer *ERC20TransferSession) SetMinter(_minter common.Address) (*types.Transaction, error) {
	return _ERC20Transfer.Contract.SetMinter(&_ERC20Transfer.TransactOpts, _minter)
}

// SetMinter is a paid mutator transaction binding the contract method 0x1652e9fc.
//
// Solidity: function set_minter(address _minter) returns()
func (_ERC20Transfer *ERC20TransferTransactorSession) SetMinter(_minter common.Address) (*types.Transaction, error) {
	return _ERC20Transfer.Contract.SetMinter(&_ERC20Transfer.TransactOpts, _minter)
}

// SetName is a paid mutator transaction binding the contract method 0xe1430e06.
//
// Solidity: function set_name(string _name, string _symbol) returns()
func (_ERC20Transfer *ERC20TransferTransactor) SetName(opts *bind.TransactOpts, _name string, _symbol string) (*types.Transaction, error) {
	return _ERC20Transfer.contract.Transact(opts, "set_name", _name, _symbol)
}

// SetName is a paid mutator transaction binding the contract method 0xe1430e06.
//
// Solidity: function set_name(string _name, string _symbol) returns()
func (_ERC20Transfer *ERC20TransferSession) SetName(_name string, _symbol string) (*types.Transaction, error) {
	return _ERC20Transfer.Contract.SetName(&_ERC20Transfer.TransactOpts, _name, _symbol)
}

// SetName is a paid mutator transaction binding the contract method 0xe1430e06.
//
// Solidity: function set_name(string _name, string _symbol) returns()
func (_ERC20Transfer *ERC20TransferTransactorSession) SetName(_name string, _symbol string) (*types.Transaction, error) {
	return _ERC20Transfer.Contract.SetName(&_ERC20Transfer.TransactOpts, _name, _symbol)
}

// StartEpochTimeWrite is a paid mutator transaction binding the contract method 0xadc4cf43.
//
// Solidity: function start_epoch_time_write() returns(uint256)
func (_ERC20Transfer *ERC20TransferTransactor) StartEpochTimeWrite(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Transfer.contract.Transact(opts, "start_epoch_time_write")
}

// StartEpochTimeWrite is a paid mutator transaction binding the contract method 0xadc4cf43.
//
// Solidity: function start_epoch_time_write() returns(uint256)
func (_ERC20Transfer *ERC20TransferSession) StartEpochTimeWrite() (*types.Transaction, error) {
	return _ERC20Transfer.Contract.StartEpochTimeWrite(&_ERC20Transfer.TransactOpts)
}

// StartEpochTimeWrite is a paid mutator transaction binding the contract method 0xadc4cf43.
//
// Solidity: function start_epoch_time_write() returns(uint256)
func (_ERC20Transfer *ERC20TransferTransactorSession) StartEpochTimeWrite() (*types.Transaction, error) {
	return _ERC20Transfer.Contract.StartEpochTimeWrite(&_ERC20Transfer.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_ERC20Transfer *ERC20TransferTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20Transfer.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_ERC20Transfer *ERC20TransferSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20Transfer.Contract.Transfer(&_ERC20Transfer.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_ERC20Transfer *ERC20TransferTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20Transfer.Contract.Transfer(&_ERC20Transfer.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_ERC20Transfer *ERC20TransferTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20Transfer.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_ERC20Transfer *ERC20TransferSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20Transfer.Contract.TransferFrom(&_ERC20Transfer.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_ERC20Transfer *ERC20TransferTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20Transfer.Contract.TransferFrom(&_ERC20Transfer.TransactOpts, _from, _to, _value)
}

// UpdateMiningParameters is a paid mutator transaction binding the contract method 0xd43b40fa.
//
// Solidity: function update_mining_parameters() returns()
func (_ERC20Transfer *ERC20TransferTransactor) UpdateMiningParameters(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Transfer.contract.Transact(opts, "update_mining_parameters")
}

// UpdateMiningParameters is a paid mutator transaction binding the contract method 0xd43b40fa.
//
// Solidity: function update_mining_parameters() returns()
func (_ERC20Transfer *ERC20TransferSession) UpdateMiningParameters() (*types.Transaction, error) {
	return _ERC20Transfer.Contract.UpdateMiningParameters(&_ERC20Transfer.TransactOpts)
}

// UpdateMiningParameters is a paid mutator transaction binding the contract method 0xd43b40fa.
//
// Solidity: function update_mining_parameters() returns()
func (_ERC20Transfer *ERC20TransferTransactorSession) UpdateMiningParameters() (*types.Transaction, error) {
	return _ERC20Transfer.Contract.UpdateMiningParameters(&_ERC20Transfer.TransactOpts)
}

// ERC20TransferApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC20Transfer contract.
type ERC20TransferApprovalIterator struct {
	Event *ERC20TransferApproval // Event containing the contract specifics and raw log

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
func (it *ERC20TransferApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TransferApproval)
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
		it.Event = new(ERC20TransferApproval)
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
func (it *ERC20TransferApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TransferApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TransferApproval represents a Approval event raised by the ERC20Transfer contract.
type ERC20TransferApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _spender, uint256 _value)
func (_ERC20Transfer *ERC20TransferFilterer) FilterApproval(opts *bind.FilterOpts, _owner []common.Address, _spender []common.Address) (*ERC20TransferApprovalIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _ERC20Transfer.contract.FilterLogs(opts, "Approval", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TransferApprovalIterator{contract: _ERC20Transfer.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _spender, uint256 _value)
func (_ERC20Transfer *ERC20TransferFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC20TransferApproval, _owner []common.Address, _spender []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _ERC20Transfer.contract.WatchLogs(opts, "Approval", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TransferApproval)
				if err := _ERC20Transfer.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_ERC20Transfer *ERC20TransferFilterer) ParseApproval(log types.Log) (*ERC20TransferApproval, error) {
	event := new(ERC20TransferApproval)
	if err := _ERC20Transfer.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TransferSetAdminIterator is returned from FilterSetAdmin and is used to iterate over the raw logs and unpacked data for SetAdmin events raised by the ERC20Transfer contract.
type ERC20TransferSetAdminIterator struct {
	Event *ERC20TransferSetAdmin // Event containing the contract specifics and raw log

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
func (it *ERC20TransferSetAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TransferSetAdmin)
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
		it.Event = new(ERC20TransferSetAdmin)
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
func (it *ERC20TransferSetAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TransferSetAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TransferSetAdmin represents a SetAdmin event raised by the ERC20Transfer contract.
type ERC20TransferSetAdmin struct {
	Admin common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterSetAdmin is a free log retrieval operation binding the contract event 0x5a272403b402d892977df56625f4164ccaf70ca3863991c43ecfe76a6905b0a1.
//
// Solidity: event SetAdmin(address admin)
func (_ERC20Transfer *ERC20TransferFilterer) FilterSetAdmin(opts *bind.FilterOpts) (*ERC20TransferSetAdminIterator, error) {

	logs, sub, err := _ERC20Transfer.contract.FilterLogs(opts, "SetAdmin")
	if err != nil {
		return nil, err
	}
	return &ERC20TransferSetAdminIterator{contract: _ERC20Transfer.contract, event: "SetAdmin", logs: logs, sub: sub}, nil
}

// WatchSetAdmin is a free log subscription operation binding the contract event 0x5a272403b402d892977df56625f4164ccaf70ca3863991c43ecfe76a6905b0a1.
//
// Solidity: event SetAdmin(address admin)
func (_ERC20Transfer *ERC20TransferFilterer) WatchSetAdmin(opts *bind.WatchOpts, sink chan<- *ERC20TransferSetAdmin) (event.Subscription, error) {

	logs, sub, err := _ERC20Transfer.contract.WatchLogs(opts, "SetAdmin")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TransferSetAdmin)
				if err := _ERC20Transfer.contract.UnpackLog(event, "SetAdmin", log); err != nil {
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
func (_ERC20Transfer *ERC20TransferFilterer) ParseSetAdmin(log types.Log) (*ERC20TransferSetAdmin, error) {
	event := new(ERC20TransferSetAdmin)
	if err := _ERC20Transfer.contract.UnpackLog(event, "SetAdmin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TransferSetMinterIterator is returned from FilterSetMinter and is used to iterate over the raw logs and unpacked data for SetMinter events raised by the ERC20Transfer contract.
type ERC20TransferSetMinterIterator struct {
	Event *ERC20TransferSetMinter // Event containing the contract specifics and raw log

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
func (it *ERC20TransferSetMinterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TransferSetMinter)
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
		it.Event = new(ERC20TransferSetMinter)
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
func (it *ERC20TransferSetMinterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TransferSetMinterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TransferSetMinter represents a SetMinter event raised by the ERC20Transfer contract.
type ERC20TransferSetMinter struct {
	Minter common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetMinter is a free log retrieval operation binding the contract event 0xcec52196e972044edde8689a1b608e459c5946b7f3e5c8cd3d6d8e126d422e1c.
//
// Solidity: event SetMinter(address minter)
func (_ERC20Transfer *ERC20TransferFilterer) FilterSetMinter(opts *bind.FilterOpts) (*ERC20TransferSetMinterIterator, error) {

	logs, sub, err := _ERC20Transfer.contract.FilterLogs(opts, "SetMinter")
	if err != nil {
		return nil, err
	}
	return &ERC20TransferSetMinterIterator{contract: _ERC20Transfer.contract, event: "SetMinter", logs: logs, sub: sub}, nil
}

// WatchSetMinter is a free log subscription operation binding the contract event 0xcec52196e972044edde8689a1b608e459c5946b7f3e5c8cd3d6d8e126d422e1c.
//
// Solidity: event SetMinter(address minter)
func (_ERC20Transfer *ERC20TransferFilterer) WatchSetMinter(opts *bind.WatchOpts, sink chan<- *ERC20TransferSetMinter) (event.Subscription, error) {

	logs, sub, err := _ERC20Transfer.contract.WatchLogs(opts, "SetMinter")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TransferSetMinter)
				if err := _ERC20Transfer.contract.UnpackLog(event, "SetMinter", log); err != nil {
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
func (_ERC20Transfer *ERC20TransferFilterer) ParseSetMinter(log types.Log) (*ERC20TransferSetMinter, error) {
	event := new(ERC20TransferSetMinter)
	if err := _ERC20Transfer.contract.UnpackLog(event, "SetMinter", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TransferTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC20Transfer contract.
type ERC20TransferTransferIterator struct {
	Event *ERC20TransferTransfer // Event containing the contract specifics and raw log

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
func (it *ERC20TransferTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TransferTransfer)
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
		it.Event = new(ERC20TransferTransfer)
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
func (it *ERC20TransferTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TransferTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TransferTransfer represents a Transfer event raised by the ERC20Transfer contract.
type ERC20TransferTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _value)
func (_ERC20Transfer *ERC20TransferFilterer) FilterTransfer(opts *bind.FilterOpts, _from []common.Address, _to []common.Address) (*ERC20TransferTransferIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _ERC20Transfer.contract.FilterLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TransferTransferIterator{contract: _ERC20Transfer.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _value)
func (_ERC20Transfer *ERC20TransferFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC20TransferTransfer, _from []common.Address, _to []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _ERC20Transfer.contract.WatchLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TransferTransfer)
				if err := _ERC20Transfer.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_ERC20Transfer *ERC20TransferFilterer) ParseTransfer(log types.Log) (*ERC20TransferTransfer, error) {
	event := new(ERC20TransferTransfer)
	if err := _ERC20Transfer.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TransferUpdateMiningParametersIterator is returned from FilterUpdateMiningParameters and is used to iterate over the raw logs and unpacked data for UpdateMiningParameters events raised by the ERC20Transfer contract.
type ERC20TransferUpdateMiningParametersIterator struct {
	Event *ERC20TransferUpdateMiningParameters // Event containing the contract specifics and raw log

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
func (it *ERC20TransferUpdateMiningParametersIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TransferUpdateMiningParameters)
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
		it.Event = new(ERC20TransferUpdateMiningParameters)
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
func (it *ERC20TransferUpdateMiningParametersIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TransferUpdateMiningParametersIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TransferUpdateMiningParameters represents a UpdateMiningParameters event raised by the ERC20Transfer contract.
type ERC20TransferUpdateMiningParameters struct {
	Time   *big.Int
	Rate   *big.Int
	Supply *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterUpdateMiningParameters is a free log retrieval operation binding the contract event 0x27e46362a1e6129b6dd539c984ce739291a97128dfcaeca1255e8ac83abd9441.
//
// Solidity: event UpdateMiningParameters(uint256 time, uint256 rate, uint256 supply)
func (_ERC20Transfer *ERC20TransferFilterer) FilterUpdateMiningParameters(opts *bind.FilterOpts) (*ERC20TransferUpdateMiningParametersIterator, error) {

	logs, sub, err := _ERC20Transfer.contract.FilterLogs(opts, "UpdateMiningParameters")
	if err != nil {
		return nil, err
	}
	return &ERC20TransferUpdateMiningParametersIterator{contract: _ERC20Transfer.contract, event: "UpdateMiningParameters", logs: logs, sub: sub}, nil
}

// WatchUpdateMiningParameters is a free log subscription operation binding the contract event 0x27e46362a1e6129b6dd539c984ce739291a97128dfcaeca1255e8ac83abd9441.
//
// Solidity: event UpdateMiningParameters(uint256 time, uint256 rate, uint256 supply)
func (_ERC20Transfer *ERC20TransferFilterer) WatchUpdateMiningParameters(opts *bind.WatchOpts, sink chan<- *ERC20TransferUpdateMiningParameters) (event.Subscription, error) {

	logs, sub, err := _ERC20Transfer.contract.WatchLogs(opts, "UpdateMiningParameters")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TransferUpdateMiningParameters)
				if err := _ERC20Transfer.contract.UnpackLog(event, "UpdateMiningParameters", log); err != nil {
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
func (_ERC20Transfer *ERC20TransferFilterer) ParseUpdateMiningParameters(log types.Log) (*ERC20TransferUpdateMiningParameters, error) {
	event := new(ERC20TransferUpdateMiningParameters)
	if err := _ERC20Transfer.contract.UnpackLog(event, "UpdateMiningParameters", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
