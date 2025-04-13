// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package oneinchv2

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

// PromotionMetaData contains all meta data concerning the Promotion contract.
var PromotionMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIVotable\",\"name\":\"token_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"resolverPercentageThreshold_\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AlreadyRegistered\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BalanceLessThanThreshold\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ETHTransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IndexOutOfBounds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotWhitelisted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OutputArrayTooSmall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PopFromEmptyArray\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SafeTransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SamePromotee\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"promoter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"promotee\",\"type\":\"address\"}],\"name\":\"Promotion\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"Registered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"resolverPercentageThreshold\",\"type\":\"uint256\"}],\"name\":\"ResolverPercentageThresholdSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"Unregistered\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BASIS_POINTS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"clean\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"getPromotees\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"promotees\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getWhitelist\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"promotee\",\"type\":\"address\"}],\"name\":\"promote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"promotions\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"rescueFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"resolverPercentageThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"resolverPercentageThreshold_\",\"type\":\"uint256\"}],\"name\":\"setResolverPercentageThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractIVotable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// PromotionABI is the input ABI used to generate the binding from.
// Deprecated: Use PromotionMetaData.ABI instead.
var PromotionABI = PromotionMetaData.ABI

// Promotion is an auto generated Go binding around an Ethereum contract.
type Promotion struct {
	PromotionCaller     // Read-only binding to the contract
	PromotionTransactor // Write-only binding to the contract
	PromotionFilterer   // Log filterer for contract events
}

// PromotionCaller is an auto generated read-only Go binding around an Ethereum contract.
type PromotionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PromotionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PromotionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PromotionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PromotionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PromotionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PromotionSession struct {
	Contract     *Promotion        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PromotionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PromotionCallerSession struct {
	Contract *PromotionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// PromotionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PromotionTransactorSession struct {
	Contract     *PromotionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// PromotionRaw is an auto generated low-level Go binding around an Ethereum contract.
type PromotionRaw struct {
	Contract *Promotion // Generic contract binding to access the raw methods on
}

// PromotionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PromotionCallerRaw struct {
	Contract *PromotionCaller // Generic read-only contract binding to access the raw methods on
}

// PromotionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PromotionTransactorRaw struct {
	Contract *PromotionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPromotion creates a new instance of Promotion, bound to a specific deployed contract.
func NewPromotion(address common.Address, backend bind.ContractBackend) (*Promotion, error) {
	contract, err := bindPromotion(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Promotion{PromotionCaller: PromotionCaller{contract: contract}, PromotionTransactor: PromotionTransactor{contract: contract}, PromotionFilterer: PromotionFilterer{contract: contract}}, nil
}

// NewPromotionCaller creates a new read-only instance of Promotion, bound to a specific deployed contract.
func NewPromotionCaller(address common.Address, caller bind.ContractCaller) (*PromotionCaller, error) {
	contract, err := bindPromotion(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PromotionCaller{contract: contract}, nil
}

// NewPromotionTransactor creates a new write-only instance of Promotion, bound to a specific deployed contract.
func NewPromotionTransactor(address common.Address, transactor bind.ContractTransactor) (*PromotionTransactor, error) {
	contract, err := bindPromotion(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PromotionTransactor{contract: contract}, nil
}

// NewPromotionFilterer creates a new log filterer instance of Promotion, bound to a specific deployed contract.
func NewPromotionFilterer(address common.Address, filterer bind.ContractFilterer) (*PromotionFilterer, error) {
	contract, err := bindPromotion(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PromotionFilterer{contract: contract}, nil
}

// bindPromotion binds a generic wrapper to an already deployed contract.
func bindPromotion(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PromotionMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Promotion *PromotionRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Promotion.Contract.PromotionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Promotion *PromotionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Promotion.Contract.PromotionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Promotion *PromotionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Promotion.Contract.PromotionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Promotion *PromotionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Promotion.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Promotion *PromotionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Promotion.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Promotion *PromotionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Promotion.Contract.contract.Transact(opts, method, params...)
}

// BASISPOINTS is a free data retrieval call binding the contract method 0xe1f1c4a7.
//
// Solidity: function BASIS_POINTS() view returns(uint256)
func (_Promotion *PromotionCaller) BASISPOINTS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Promotion.contract.Call(opts, &out, "BASIS_POINTS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BASISPOINTS is a free data retrieval call binding the contract method 0xe1f1c4a7.
//
// Solidity: function BASIS_POINTS() view returns(uint256)
func (_Promotion *PromotionSession) BASISPOINTS() (*big.Int, error) {
	return _Promotion.Contract.BASISPOINTS(&_Promotion.CallOpts)
}

// BASISPOINTS is a free data retrieval call binding the contract method 0xe1f1c4a7.
//
// Solidity: function BASIS_POINTS() view returns(uint256)
func (_Promotion *PromotionCallerSession) BASISPOINTS() (*big.Int, error) {
	return _Promotion.Contract.BASISPOINTS(&_Promotion.CallOpts)
}

// GetPromotees is a free data retrieval call binding the contract method 0xe5136e94.
//
// Solidity: function getPromotees(uint256 chainId) view returns(address[] promotees)
func (_Promotion *PromotionCaller) GetPromotees(opts *bind.CallOpts, chainId *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _Promotion.contract.Call(opts, &out, "getPromotees", chainId)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetPromotees is a free data retrieval call binding the contract method 0xe5136e94.
//
// Solidity: function getPromotees(uint256 chainId) view returns(address[] promotees)
func (_Promotion *PromotionSession) GetPromotees(chainId *big.Int) ([]common.Address, error) {
	return _Promotion.Contract.GetPromotees(&_Promotion.CallOpts, chainId)
}

// GetPromotees is a free data retrieval call binding the contract method 0xe5136e94.
//
// Solidity: function getPromotees(uint256 chainId) view returns(address[] promotees)
func (_Promotion *PromotionCallerSession) GetPromotees(chainId *big.Int) ([]common.Address, error) {
	return _Promotion.Contract.GetPromotees(&_Promotion.CallOpts, chainId)
}

// GetWhitelist is a free data retrieval call binding the contract method 0xd01f63f5.
//
// Solidity: function getWhitelist() view returns(address[])
func (_Promotion *PromotionCaller) GetWhitelist(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Promotion.contract.Call(opts, &out, "getWhitelist")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetWhitelist is a free data retrieval call binding the contract method 0xd01f63f5.
//
// Solidity: function getWhitelist() view returns(address[])
func (_Promotion *PromotionSession) GetWhitelist() ([]common.Address, error) {
	return _Promotion.Contract.GetWhitelist(&_Promotion.CallOpts)
}

// GetWhitelist is a free data retrieval call binding the contract method 0xd01f63f5.
//
// Solidity: function getWhitelist() view returns(address[])
func (_Promotion *PromotionCallerSession) GetWhitelist() ([]common.Address, error) {
	return _Promotion.Contract.GetWhitelist(&_Promotion.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Promotion *PromotionCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Promotion.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Promotion *PromotionSession) Owner() (common.Address, error) {
	return _Promotion.Contract.Owner(&_Promotion.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Promotion *PromotionCallerSession) Owner() (common.Address, error) {
	return _Promotion.Contract.Owner(&_Promotion.CallOpts)
}

// Promotions is a free data retrieval call binding the contract method 0x2af05166.
//
// Solidity: function promotions(address , uint256 ) view returns(address)
func (_Promotion *PromotionCaller) Promotions(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Promotion.contract.Call(opts, &out, "promotions", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Promotions is a free data retrieval call binding the contract method 0x2af05166.
//
// Solidity: function promotions(address , uint256 ) view returns(address)
func (_Promotion *PromotionSession) Promotions(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _Promotion.Contract.Promotions(&_Promotion.CallOpts, arg0, arg1)
}

// Promotions is a free data retrieval call binding the contract method 0x2af05166.
//
// Solidity: function promotions(address , uint256 ) view returns(address)
func (_Promotion *PromotionCallerSession) Promotions(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _Promotion.Contract.Promotions(&_Promotion.CallOpts, arg0, arg1)
}

// ResolverPercentageThreshold is a free data retrieval call binding the contract method 0x960462d1.
//
// Solidity: function resolverPercentageThreshold() view returns(uint256)
func (_Promotion *PromotionCaller) ResolverPercentageThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Promotion.contract.Call(opts, &out, "resolverPercentageThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ResolverPercentageThreshold is a free data retrieval call binding the contract method 0x960462d1.
//
// Solidity: function resolverPercentageThreshold() view returns(uint256)
func (_Promotion *PromotionSession) ResolverPercentageThreshold() (*big.Int, error) {
	return _Promotion.Contract.ResolverPercentageThreshold(&_Promotion.CallOpts)
}

// ResolverPercentageThreshold is a free data retrieval call binding the contract method 0x960462d1.
//
// Solidity: function resolverPercentageThreshold() view returns(uint256)
func (_Promotion *PromotionCallerSession) ResolverPercentageThreshold() (*big.Int, error) {
	return _Promotion.Contract.ResolverPercentageThreshold(&_Promotion.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Promotion *PromotionCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Promotion.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Promotion *PromotionSession) Token() (common.Address, error) {
	return _Promotion.Contract.Token(&_Promotion.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Promotion *PromotionCallerSession) Token() (common.Address, error) {
	return _Promotion.Contract.Token(&_Promotion.CallOpts)
}

// Clean is a paid mutator transaction binding the contract method 0xfc4333cd.
//
// Solidity: function clean() returns()
func (_Promotion *PromotionTransactor) Clean(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Promotion.contract.Transact(opts, "clean")
}

// Clean is a paid mutator transaction binding the contract method 0xfc4333cd.
//
// Solidity: function clean() returns()
func (_Promotion *PromotionSession) Clean() (*types.Transaction, error) {
	return _Promotion.Contract.Clean(&_Promotion.TransactOpts)
}

// Clean is a paid mutator transaction binding the contract method 0xfc4333cd.
//
// Solidity: function clean() returns()
func (_Promotion *PromotionTransactorSession) Clean() (*types.Transaction, error) {
	return _Promotion.Contract.Clean(&_Promotion.TransactOpts)
}

// Promote is a paid mutator transaction binding the contract method 0xf204bdb9.
//
// Solidity: function promote(uint256 chainId, address promotee) returns()
func (_Promotion *PromotionTransactor) Promote(opts *bind.TransactOpts, chainId *big.Int, promotee common.Address) (*types.Transaction, error) {
	return _Promotion.contract.Transact(opts, "promote", chainId, promotee)
}

// Promote is a paid mutator transaction binding the contract method 0xf204bdb9.
//
// Solidity: function promote(uint256 chainId, address promotee) returns()
func (_Promotion *PromotionSession) Promote(chainId *big.Int, promotee common.Address) (*types.Transaction, error) {
	return _Promotion.Contract.Promote(&_Promotion.TransactOpts, chainId, promotee)
}

// Promote is a paid mutator transaction binding the contract method 0xf204bdb9.
//
// Solidity: function promote(uint256 chainId, address promotee) returns()
func (_Promotion *PromotionTransactorSession) Promote(chainId *big.Int, promotee common.Address) (*types.Transaction, error) {
	return _Promotion.Contract.Promote(&_Promotion.TransactOpts, chainId, promotee)
}

// Register is a paid mutator transaction binding the contract method 0x1aa3a008.
//
// Solidity: function register() returns()
func (_Promotion *PromotionTransactor) Register(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Promotion.contract.Transact(opts, "register")
}

// Register is a paid mutator transaction binding the contract method 0x1aa3a008.
//
// Solidity: function register() returns()
func (_Promotion *PromotionSession) Register() (*types.Transaction, error) {
	return _Promotion.Contract.Register(&_Promotion.TransactOpts)
}

// Register is a paid mutator transaction binding the contract method 0x1aa3a008.
//
// Solidity: function register() returns()
func (_Promotion *PromotionTransactorSession) Register() (*types.Transaction, error) {
	return _Promotion.Contract.Register(&_Promotion.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Promotion *PromotionTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Promotion.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Promotion *PromotionSession) RenounceOwnership() (*types.Transaction, error) {
	return _Promotion.Contract.RenounceOwnership(&_Promotion.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Promotion *PromotionTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Promotion.Contract.RenounceOwnership(&_Promotion.TransactOpts)
}

// RescueFunds is a paid mutator transaction binding the contract method 0x78e3214f.
//
// Solidity: function rescueFunds(address token_, uint256 amount) returns()
func (_Promotion *PromotionTransactor) RescueFunds(opts *bind.TransactOpts, token_ common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Promotion.contract.Transact(opts, "rescueFunds", token_, amount)
}

// RescueFunds is a paid mutator transaction binding the contract method 0x78e3214f.
//
// Solidity: function rescueFunds(address token_, uint256 amount) returns()
func (_Promotion *PromotionSession) RescueFunds(token_ common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Promotion.Contract.RescueFunds(&_Promotion.TransactOpts, token_, amount)
}

// RescueFunds is a paid mutator transaction binding the contract method 0x78e3214f.
//
// Solidity: function rescueFunds(address token_, uint256 amount) returns()
func (_Promotion *PromotionTransactorSession) RescueFunds(token_ common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Promotion.Contract.RescueFunds(&_Promotion.TransactOpts, token_, amount)
}

// SetResolverPercentageThreshold is a paid mutator transaction binding the contract method 0xcfaa0951.
//
// Solidity: function setResolverPercentageThreshold(uint256 resolverPercentageThreshold_) returns()
func (_Promotion *PromotionTransactor) SetResolverPercentageThreshold(opts *bind.TransactOpts, resolverPercentageThreshold_ *big.Int) (*types.Transaction, error) {
	return _Promotion.contract.Transact(opts, "setResolverPercentageThreshold", resolverPercentageThreshold_)
}

// SetResolverPercentageThreshold is a paid mutator transaction binding the contract method 0xcfaa0951.
//
// Solidity: function setResolverPercentageThreshold(uint256 resolverPercentageThreshold_) returns()
func (_Promotion *PromotionSession) SetResolverPercentageThreshold(resolverPercentageThreshold_ *big.Int) (*types.Transaction, error) {
	return _Promotion.Contract.SetResolverPercentageThreshold(&_Promotion.TransactOpts, resolverPercentageThreshold_)
}

// SetResolverPercentageThreshold is a paid mutator transaction binding the contract method 0xcfaa0951.
//
// Solidity: function setResolverPercentageThreshold(uint256 resolverPercentageThreshold_) returns()
func (_Promotion *PromotionTransactorSession) SetResolverPercentageThreshold(resolverPercentageThreshold_ *big.Int) (*types.Transaction, error) {
	return _Promotion.Contract.SetResolverPercentageThreshold(&_Promotion.TransactOpts, resolverPercentageThreshold_)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Promotion *PromotionTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Promotion.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Promotion *PromotionSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Promotion.Contract.TransferOwnership(&_Promotion.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Promotion *PromotionTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Promotion.Contract.TransferOwnership(&_Promotion.TransactOpts, newOwner)
}

// PromotionOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Promotion contract.
type PromotionOwnershipTransferredIterator struct {
	Event *PromotionOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PromotionOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PromotionOwnershipTransferred)
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
		it.Event = new(PromotionOwnershipTransferred)
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
func (it *PromotionOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PromotionOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PromotionOwnershipTransferred represents a OwnershipTransferred event raised by the Promotion contract.
type PromotionOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Promotion *PromotionFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PromotionOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Promotion.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PromotionOwnershipTransferredIterator{contract: _Promotion.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Promotion *PromotionFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PromotionOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Promotion.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PromotionOwnershipTransferred)
				if err := _Promotion.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Promotion *PromotionFilterer) ParseOwnershipTransferred(log types.Log) (*PromotionOwnershipTransferred, error) {
	event := new(PromotionOwnershipTransferred)
	if err := _Promotion.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PromotionPromotionIterator is returned from FilterPromotion and is used to iterate over the raw logs and unpacked data for Promotion events raised by the Promotion contract.
type PromotionPromotionIterator struct {
	Event *PromotionPromotion // Event containing the contract specifics and raw log

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
func (it *PromotionPromotionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PromotionPromotion)
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
		it.Event = new(PromotionPromotion)
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
func (it *PromotionPromotionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PromotionPromotionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PromotionPromotion represents a Promotion event raised by the Promotion contract.
type PromotionPromotion struct {
	Promoter common.Address
	ChainId  *big.Int
	Promotee common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPromotion is a free log retrieval operation binding the contract event 0xb863cf86b291171e4b0332ea12b59af030f98a2c74a6d51effaf1109ae4c7f1e.
//
// Solidity: event Promotion(address promoter, uint256 chainId, address promotee)
func (_Promotion *PromotionFilterer) FilterPromotion(opts *bind.FilterOpts) (*PromotionPromotionIterator, error) {

	logs, sub, err := _Promotion.contract.FilterLogs(opts, "Promotion")
	if err != nil {
		return nil, err
	}
	return &PromotionPromotionIterator{contract: _Promotion.contract, event: "Promotion", logs: logs, sub: sub}, nil
}

// WatchPromotion is a free log subscription operation binding the contract event 0xb863cf86b291171e4b0332ea12b59af030f98a2c74a6d51effaf1109ae4c7f1e.
//
// Solidity: event Promotion(address promoter, uint256 chainId, address promotee)
func (_Promotion *PromotionFilterer) WatchPromotion(opts *bind.WatchOpts, sink chan<- *PromotionPromotion) (event.Subscription, error) {

	logs, sub, err := _Promotion.contract.WatchLogs(opts, "Promotion")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PromotionPromotion)
				if err := _Promotion.contract.UnpackLog(event, "Promotion", log); err != nil {
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

// ParsePromotion is a log parse operation binding the contract event 0xb863cf86b291171e4b0332ea12b59af030f98a2c74a6d51effaf1109ae4c7f1e.
//
// Solidity: event Promotion(address promoter, uint256 chainId, address promotee)
func (_Promotion *PromotionFilterer) ParsePromotion(log types.Log) (*PromotionPromotion, error) {
	event := new(PromotionPromotion)
	if err := _Promotion.contract.UnpackLog(event, "Promotion", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PromotionRegisteredIterator is returned from FilterRegistered and is used to iterate over the raw logs and unpacked data for Registered events raised by the Promotion contract.
type PromotionRegisteredIterator struct {
	Event *PromotionRegistered // Event containing the contract specifics and raw log

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
func (it *PromotionRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PromotionRegistered)
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
		it.Event = new(PromotionRegistered)
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
func (it *PromotionRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PromotionRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PromotionRegistered represents a Registered event raised by the Promotion contract.
type PromotionRegistered struct {
	Addr common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterRegistered is a free log retrieval operation binding the contract event 0x2d3734a8e47ac8316e500ac231c90a6e1848ca2285f40d07eaa52005e4b3a0e9.
//
// Solidity: event Registered(address addr)
func (_Promotion *PromotionFilterer) FilterRegistered(opts *bind.FilterOpts) (*PromotionRegisteredIterator, error) {

	logs, sub, err := _Promotion.contract.FilterLogs(opts, "Registered")
	if err != nil {
		return nil, err
	}
	return &PromotionRegisteredIterator{contract: _Promotion.contract, event: "Registered", logs: logs, sub: sub}, nil
}

// WatchRegistered is a free log subscription operation binding the contract event 0x2d3734a8e47ac8316e500ac231c90a6e1848ca2285f40d07eaa52005e4b3a0e9.
//
// Solidity: event Registered(address addr)
func (_Promotion *PromotionFilterer) WatchRegistered(opts *bind.WatchOpts, sink chan<- *PromotionRegistered) (event.Subscription, error) {

	logs, sub, err := _Promotion.contract.WatchLogs(opts, "Registered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PromotionRegistered)
				if err := _Promotion.contract.UnpackLog(event, "Registered", log); err != nil {
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

// ParseRegistered is a log parse operation binding the contract event 0x2d3734a8e47ac8316e500ac231c90a6e1848ca2285f40d07eaa52005e4b3a0e9.
//
// Solidity: event Registered(address addr)
func (_Promotion *PromotionFilterer) ParseRegistered(log types.Log) (*PromotionRegistered, error) {
	event := new(PromotionRegistered)
	if err := _Promotion.contract.UnpackLog(event, "Registered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PromotionResolverPercentageThresholdSetIterator is returned from FilterResolverPercentageThresholdSet and is used to iterate over the raw logs and unpacked data for ResolverPercentageThresholdSet events raised by the Promotion contract.
type PromotionResolverPercentageThresholdSetIterator struct {
	Event *PromotionResolverPercentageThresholdSet // Event containing the contract specifics and raw log

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
func (it *PromotionResolverPercentageThresholdSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PromotionResolverPercentageThresholdSet)
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
		it.Event = new(PromotionResolverPercentageThresholdSet)
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
func (it *PromotionResolverPercentageThresholdSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PromotionResolverPercentageThresholdSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PromotionResolverPercentageThresholdSet represents a ResolverPercentageThresholdSet event raised by the Promotion contract.
type PromotionResolverPercentageThresholdSet struct {
	ResolverPercentageThreshold *big.Int
	Raw                         types.Log // Blockchain specific contextual infos
}

// FilterResolverPercentageThresholdSet is a free log retrieval operation binding the contract event 0xcbcd22f10b3a748cd1a2d091ce0b2108352aa00351115f173e159e24d97d280e.
//
// Solidity: event ResolverPercentageThresholdSet(uint256 resolverPercentageThreshold)
func (_Promotion *PromotionFilterer) FilterResolverPercentageThresholdSet(opts *bind.FilterOpts) (*PromotionResolverPercentageThresholdSetIterator, error) {

	logs, sub, err := _Promotion.contract.FilterLogs(opts, "ResolverPercentageThresholdSet")
	if err != nil {
		return nil, err
	}
	return &PromotionResolverPercentageThresholdSetIterator{contract: _Promotion.contract, event: "ResolverPercentageThresholdSet", logs: logs, sub: sub}, nil
}

// WatchResolverPercentageThresholdSet is a free log subscription operation binding the contract event 0xcbcd22f10b3a748cd1a2d091ce0b2108352aa00351115f173e159e24d97d280e.
//
// Solidity: event ResolverPercentageThresholdSet(uint256 resolverPercentageThreshold)
func (_Promotion *PromotionFilterer) WatchResolverPercentageThresholdSet(opts *bind.WatchOpts, sink chan<- *PromotionResolverPercentageThresholdSet) (event.Subscription, error) {

	logs, sub, err := _Promotion.contract.WatchLogs(opts, "ResolverPercentageThresholdSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PromotionResolverPercentageThresholdSet)
				if err := _Promotion.contract.UnpackLog(event, "ResolverPercentageThresholdSet", log); err != nil {
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

// ParseResolverPercentageThresholdSet is a log parse operation binding the contract event 0xcbcd22f10b3a748cd1a2d091ce0b2108352aa00351115f173e159e24d97d280e.
//
// Solidity: event ResolverPercentageThresholdSet(uint256 resolverPercentageThreshold)
func (_Promotion *PromotionFilterer) ParseResolverPercentageThresholdSet(log types.Log) (*PromotionResolverPercentageThresholdSet, error) {
	event := new(PromotionResolverPercentageThresholdSet)
	if err := _Promotion.contract.UnpackLog(event, "ResolverPercentageThresholdSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PromotionUnregisteredIterator is returned from FilterUnregistered and is used to iterate over the raw logs and unpacked data for Unregistered events raised by the Promotion contract.
type PromotionUnregisteredIterator struct {
	Event *PromotionUnregistered // Event containing the contract specifics and raw log

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
func (it *PromotionUnregisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PromotionUnregistered)
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
		it.Event = new(PromotionUnregistered)
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
func (it *PromotionUnregisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PromotionUnregisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PromotionUnregistered represents a Unregistered event raised by the Promotion contract.
type PromotionUnregistered struct {
	Addr common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterUnregistered is a free log retrieval operation binding the contract event 0x75cd6de711483e11488a1cd9b66172abccb9e5c19572f92015a7880f0c8c0edc.
//
// Solidity: event Unregistered(address addr)
func (_Promotion *PromotionFilterer) FilterUnregistered(opts *bind.FilterOpts) (*PromotionUnregisteredIterator, error) {

	logs, sub, err := _Promotion.contract.FilterLogs(opts, "Unregistered")
	if err != nil {
		return nil, err
	}
	return &PromotionUnregisteredIterator{contract: _Promotion.contract, event: "Unregistered", logs: logs, sub: sub}, nil
}

// WatchUnregistered is a free log subscription operation binding the contract event 0x75cd6de711483e11488a1cd9b66172abccb9e5c19572f92015a7880f0c8c0edc.
//
// Solidity: event Unregistered(address addr)
func (_Promotion *PromotionFilterer) WatchUnregistered(opts *bind.WatchOpts, sink chan<- *PromotionUnregistered) (event.Subscription, error) {

	logs, sub, err := _Promotion.contract.WatchLogs(opts, "Unregistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PromotionUnregistered)
				if err := _Promotion.contract.UnpackLog(event, "Unregistered", log); err != nil {
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

// ParseUnregistered is a log parse operation binding the contract event 0x75cd6de711483e11488a1cd9b66172abccb9e5c19572f92015a7880f0c8c0edc.
//
// Solidity: event Unregistered(address addr)
func (_Promotion *PromotionFilterer) ParseUnregistered(log types.Log) (*PromotionUnregistered, error) {
	event := new(PromotionUnregistered)
	if err := _Promotion.contract.UnpackLog(event, "Unregistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
