// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package zxrfq_v3_helper

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

// PermitTransferFrom is an auto generated low-level Go binding around an user-defined struct.
type PermitTransferFrom struct {
	Permitted TokenPermissions
	Nonce     *big.Int
	Deadline  *big.Int
}

// TokenPermissions is an auto generated low-level Go binding around an user-defined struct.
type TokenPermissions struct {
	Token  common.Address
	Amount *big.Int
}

// ZxrfqV3HelperMetaData contains all meta data concerning the ZxrfqV3Helper contract.
var ZxrfqV3HelperMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structTokenPermissions\",\"name\":\"permitted\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structPermitTransferFrom\",\"name\":\"permit\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"makerSig\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"maxTakerAmount\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"msgSender\",\"type\":\"uint256\"}],\"name\":\"DecodeParamOfFillOtcOrderSelfFunded\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ZxrfqV3HelperABI is the input ABI used to generate the binding from.
// Deprecated: Use ZxrfqV3HelperMetaData.ABI instead.
var ZxrfqV3HelperABI = ZxrfqV3HelperMetaData.ABI

// ZxrfqV3Helper is an auto generated Go binding around an Ethereum contract.
type ZxrfqV3Helper struct {
	ZxrfqV3HelperCaller     // Read-only binding to the contract
	ZxrfqV3HelperTransactor // Write-only binding to the contract
	ZxrfqV3HelperFilterer   // Log filterer for contract events
}

// ZxrfqV3HelperCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZxrfqV3HelperCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZxrfqV3HelperTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZxrfqV3HelperTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZxrfqV3HelperFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZxrfqV3HelperFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZxrfqV3HelperSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZxrfqV3HelperSession struct {
	Contract     *ZxrfqV3Helper    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZxrfqV3HelperCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZxrfqV3HelperCallerSession struct {
	Contract *ZxrfqV3HelperCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// ZxrfqV3HelperTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZxrfqV3HelperTransactorSession struct {
	Contract     *ZxrfqV3HelperTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ZxrfqV3HelperRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZxrfqV3HelperRaw struct {
	Contract *ZxrfqV3Helper // Generic contract binding to access the raw methods on
}

// ZxrfqV3HelperCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZxrfqV3HelperCallerRaw struct {
	Contract *ZxrfqV3HelperCaller // Generic read-only contract binding to access the raw methods on
}

// ZxrfqV3HelperTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZxrfqV3HelperTransactorRaw struct {
	Contract *ZxrfqV3HelperTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZxrfqV3Helper creates a new instance of ZxrfqV3Helper, bound to a specific deployed contract.
func NewZxrfqV3Helper(address common.Address, backend bind.ContractBackend) (*ZxrfqV3Helper, error) {
	contract, err := bindZxrfqV3Helper(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZxrfqV3Helper{ZxrfqV3HelperCaller: ZxrfqV3HelperCaller{contract: contract}, ZxrfqV3HelperTransactor: ZxrfqV3HelperTransactor{contract: contract}, ZxrfqV3HelperFilterer: ZxrfqV3HelperFilterer{contract: contract}}, nil
}

// NewZxrfqV3HelperCaller creates a new read-only instance of ZxrfqV3Helper, bound to a specific deployed contract.
func NewZxrfqV3HelperCaller(address common.Address, caller bind.ContractCaller) (*ZxrfqV3HelperCaller, error) {
	contract, err := bindZxrfqV3Helper(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZxrfqV3HelperCaller{contract: contract}, nil
}

// NewZxrfqV3HelperTransactor creates a new write-only instance of ZxrfqV3Helper, bound to a specific deployed contract.
func NewZxrfqV3HelperTransactor(address common.Address, transactor bind.ContractTransactor) (*ZxrfqV3HelperTransactor, error) {
	contract, err := bindZxrfqV3Helper(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZxrfqV3HelperTransactor{contract: contract}, nil
}

// NewZxrfqV3HelperFilterer creates a new log filterer instance of ZxrfqV3Helper, bound to a specific deployed contract.
func NewZxrfqV3HelperFilterer(address common.Address, filterer bind.ContractFilterer) (*ZxrfqV3HelperFilterer, error) {
	contract, err := bindZxrfqV3Helper(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZxrfqV3HelperFilterer{contract: contract}, nil
}

// bindZxrfqV3Helper binds a generic wrapper to an already deployed contract.
func bindZxrfqV3Helper(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ZxrfqV3HelperMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZxrfqV3Helper *ZxrfqV3HelperRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZxrfqV3Helper.Contract.ZxrfqV3HelperCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZxrfqV3Helper *ZxrfqV3HelperRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZxrfqV3Helper.Contract.ZxrfqV3HelperTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZxrfqV3Helper *ZxrfqV3HelperRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZxrfqV3Helper.Contract.ZxrfqV3HelperTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZxrfqV3Helper *ZxrfqV3HelperCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZxrfqV3Helper.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZxrfqV3Helper *ZxrfqV3HelperTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZxrfqV3Helper.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZxrfqV3Helper *ZxrfqV3HelperTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZxrfqV3Helper.Contract.contract.Transact(opts, method, params...)
}

// DecodeParamOfFillOtcOrderSelfFunded is a paid mutator transaction binding the contract method 0x0a164181.
//
// Solidity: function DecodeParamOfFillOtcOrderSelfFunded(address recipient, ((address,uint256),uint256,uint256) permit, address maker, bytes makerSig, address maxTakerAmount, uint256 msgSender) returns()
func (_ZxrfqV3Helper *ZxrfqV3HelperTransactor) DecodeParamOfFillOtcOrderSelfFunded(opts *bind.TransactOpts, recipient common.Address, permit PermitTransferFrom, maker common.Address, makerSig []byte, maxTakerAmount common.Address, msgSender *big.Int) (*types.Transaction, error) {
	return _ZxrfqV3Helper.contract.Transact(opts, "DecodeParamOfFillOtcOrderSelfFunded", recipient, permit, maker, makerSig, maxTakerAmount, msgSender)
}

// DecodeParamOfFillOtcOrderSelfFunded is a paid mutator transaction binding the contract method 0x0a164181.
//
// Solidity: function DecodeParamOfFillOtcOrderSelfFunded(address recipient, ((address,uint256),uint256,uint256) permit, address maker, bytes makerSig, address maxTakerAmount, uint256 msgSender) returns()
func (_ZxrfqV3Helper *ZxrfqV3HelperSession) DecodeParamOfFillOtcOrderSelfFunded(recipient common.Address, permit PermitTransferFrom, maker common.Address, makerSig []byte, maxTakerAmount common.Address, msgSender *big.Int) (*types.Transaction, error) {
	return _ZxrfqV3Helper.Contract.DecodeParamOfFillOtcOrderSelfFunded(&_ZxrfqV3Helper.TransactOpts, recipient, permit, maker, makerSig, maxTakerAmount, msgSender)
}

// DecodeParamOfFillOtcOrderSelfFunded is a paid mutator transaction binding the contract method 0x0a164181.
//
// Solidity: function DecodeParamOfFillOtcOrderSelfFunded(address recipient, ((address,uint256),uint256,uint256) permit, address maker, bytes makerSig, address maxTakerAmount, uint256 msgSender) returns()
func (_ZxrfqV3Helper *ZxrfqV3HelperTransactorSession) DecodeParamOfFillOtcOrderSelfFunded(recipient common.Address, permit PermitTransferFrom, maker common.Address, makerSig []byte, maxTakerAmount common.Address, msgSender *big.Int) (*types.Transaction, error) {
	return _ZxrfqV3Helper.Contract.DecodeParamOfFillOtcOrderSelfFunded(&_ZxrfqV3Helper.TransactOpts, recipient, permit, maker, makerSig, maxTakerAmount, msgSender)
}
