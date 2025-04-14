// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package zxotc

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

// LibNativeOrderOtcOrder is an auto generated low-level Go binding around an user-defined struct.
type LibNativeOrderOtcOrder struct {
	MakerToken     common.Address
	TakerToken     common.Address
	MakerAmount    *big.Int
	TakerAmount    *big.Int
	Maker          common.Address
	Taker          common.Address
	TxOrigin       common.Address
	ExpiryAndNonce *big.Int
}

// LibNativeOrderOtcOrderInfo is an auto generated low-level Go binding around an user-defined struct.
type LibNativeOrderOtcOrderInfo struct {
	OrderHash [32]byte
	Status    uint8
}

// LibSignatureSignature is an auto generated low-level Go binding around an user-defined struct.
type LibSignatureSignature struct {
	SignatureType uint8
	V             uint8
	R             [32]byte
	S             [32]byte
}

// ZeroXOTCMetaData contains all meta data concerning the ZeroXOTC contract.
var ZeroXOTCMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"zeroExAddress\",\"type\":\"address\"},{\"internalType\":\"contractIEtherTokenV06\",\"name\":\"weth\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"makerToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"takerToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"makerTokenFilledAmount\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"takerTokenFilledAmount\",\"type\":\"uint128\"}],\"name\":\"OtcOrderFilled\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"EIP712_DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"FEATURE_NAME\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"FEATURE_VERSION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"contractIERC20TokenV06\",\"name\":\"makerToken\",\"type\":\"address\"},{\"internalType\":\"contractIERC20TokenV06\",\"name\":\"takerToken\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"makerAmount\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"takerAmount\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"txOrigin\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expiryAndNonce\",\"type\":\"uint256\"}],\"internalType\":\"structLibNativeOrder.OtcOrder\",\"name\":\"order\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumLibSignature.SignatureType\",\"name\":\"signatureType\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structLibSignature.Signature\",\"name\":\"makerSignature\",\"type\":\"tuple\"},{\"internalType\":\"uint128\",\"name\":\"takerTokenFillAmount\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"useSelfBalance\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"_fillOtcOrder\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"takerTokenFilledAmount\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"makerTokenFilledAmount\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"contractIERC20TokenV06\",\"name\":\"makerToken\",\"type\":\"address\"},{\"internalType\":\"contractIERC20TokenV06\",\"name\":\"takerToken\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"makerAmount\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"takerAmount\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"txOrigin\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expiryAndNonce\",\"type\":\"uint256\"}],\"internalType\":\"structLibNativeOrder.OtcOrder[]\",\"name\":\"orders\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"enumLibSignature.SignatureType\",\"name\":\"signatureType\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structLibSignature.Signature[]\",\"name\":\"makerSignatures\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"enumLibSignature.SignatureType\",\"name\":\"signatureType\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structLibSignature.Signature[]\",\"name\":\"takerSignatures\",\"type\":\"tuple[]\"},{\"internalType\":\"bool[]\",\"name\":\"unwrapWeth\",\"type\":\"bool[]\"}],\"name\":\"batchFillTakerSignedOtcOrders\",\"outputs\":[{\"internalType\":\"bool[]\",\"name\":\"successes\",\"type\":\"bool[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"contractIERC20TokenV06\",\"name\":\"makerToken\",\"type\":\"address\"},{\"internalType\":\"contractIERC20TokenV06\",\"name\":\"takerToken\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"makerAmount\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"takerAmount\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"txOrigin\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expiryAndNonce\",\"type\":\"uint256\"}],\"internalType\":\"structLibNativeOrder.OtcOrder\",\"name\":\"order\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumLibSignature.SignatureType\",\"name\":\"signatureType\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structLibSignature.Signature\",\"name\":\"makerSignature\",\"type\":\"tuple\"},{\"internalType\":\"uint128\",\"name\":\"takerTokenFillAmount\",\"type\":\"uint128\"}],\"name\":\"fillOtcOrder\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"takerTokenFilledAmount\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"makerTokenFilledAmount\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"contractIERC20TokenV06\",\"name\":\"makerToken\",\"type\":\"address\"},{\"internalType\":\"contractIERC20TokenV06\",\"name\":\"takerToken\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"makerAmount\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"takerAmount\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"txOrigin\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expiryAndNonce\",\"type\":\"uint256\"}],\"internalType\":\"structLibNativeOrder.OtcOrder\",\"name\":\"order\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumLibSignature.SignatureType\",\"name\":\"signatureType\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structLibSignature.Signature\",\"name\":\"makerSignature\",\"type\":\"tuple\"},{\"internalType\":\"uint128\",\"name\":\"takerTokenFillAmount\",\"type\":\"uint128\"}],\"name\":\"fillOtcOrderForEth\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"takerTokenFilledAmount\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"makerTokenFilledAmount\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"contractIERC20TokenV06\",\"name\":\"makerToken\",\"type\":\"address\"},{\"internalType\":\"contractIERC20TokenV06\",\"name\":\"takerToken\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"makerAmount\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"takerAmount\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"txOrigin\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expiryAndNonce\",\"type\":\"uint256\"}],\"internalType\":\"structLibNativeOrder.OtcOrder\",\"name\":\"order\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumLibSignature.SignatureType\",\"name\":\"signatureType\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structLibSignature.Signature\",\"name\":\"makerSignature\",\"type\":\"tuple\"}],\"name\":\"fillOtcOrderWithEth\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"takerTokenFilledAmount\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"makerTokenFilledAmount\",\"type\":\"uint128\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"contractIERC20TokenV06\",\"name\":\"makerToken\",\"type\":\"address\"},{\"internalType\":\"contractIERC20TokenV06\",\"name\":\"takerToken\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"makerAmount\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"takerAmount\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"txOrigin\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expiryAndNonce\",\"type\":\"uint256\"}],\"internalType\":\"structLibNativeOrder.OtcOrder\",\"name\":\"order\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumLibSignature.SignatureType\",\"name\":\"signatureType\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structLibSignature.Signature\",\"name\":\"makerSignature\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumLibSignature.SignatureType\",\"name\":\"signatureType\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structLibSignature.Signature\",\"name\":\"takerSignature\",\"type\":\"tuple\"}],\"name\":\"fillTakerSignedOtcOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"contractIERC20TokenV06\",\"name\":\"makerToken\",\"type\":\"address\"},{\"internalType\":\"contractIERC20TokenV06\",\"name\":\"takerToken\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"makerAmount\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"takerAmount\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"txOrigin\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expiryAndNonce\",\"type\":\"uint256\"}],\"internalType\":\"structLibNativeOrder.OtcOrder\",\"name\":\"order\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumLibSignature.SignatureType\",\"name\":\"signatureType\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structLibSignature.Signature\",\"name\":\"makerSignature\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumLibSignature.SignatureType\",\"name\":\"signatureType\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structLibSignature.Signature\",\"name\":\"takerSignature\",\"type\":\"tuple\"}],\"name\":\"fillTakerSignedOtcOrderForEth\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"contractIERC20TokenV06\",\"name\":\"makerToken\",\"type\":\"address\"},{\"internalType\":\"contractIERC20TokenV06\",\"name\":\"takerToken\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"makerAmount\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"takerAmount\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"txOrigin\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expiryAndNonce\",\"type\":\"uint256\"}],\"internalType\":\"structLibNativeOrder.OtcOrder\",\"name\":\"order\",\"type\":\"tuple\"}],\"name\":\"getOtcOrderHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"contractIERC20TokenV06\",\"name\":\"makerToken\",\"type\":\"address\"},{\"internalType\":\"contractIERC20TokenV06\",\"name\":\"takerToken\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"makerAmount\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"takerAmount\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"txOrigin\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expiryAndNonce\",\"type\":\"uint256\"}],\"internalType\":\"structLibNativeOrder.OtcOrder\",\"name\":\"order\",\"type\":\"tuple\"}],\"name\":\"getOtcOrderInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"},{\"internalType\":\"enumLibNativeOrder.OrderStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"internalType\":\"structLibNativeOrder.OtcOrderInfo\",\"name\":\"orderInfo\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"txOrigin\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"nonceBucket\",\"type\":\"uint64\"}],\"name\":\"lastOtcTxOriginNonce\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"lastNonce\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"migrate\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"success\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ZeroXOTCABI is the input ABI used to generate the binding from.
// Deprecated: Use ZeroXOTCMetaData.ABI instead.
var ZeroXOTCABI = ZeroXOTCMetaData.ABI

// ZeroXOTC is an auto generated Go binding around an Ethereum contract.
type ZeroXOTC struct {
	ZeroXOTCCaller     // Read-only binding to the contract
	ZeroXOTCTransactor // Write-only binding to the contract
	ZeroXOTCFilterer   // Log filterer for contract events
}

// ZeroXOTCCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZeroXOTCCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZeroXOTCTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZeroXOTCTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZeroXOTCFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZeroXOTCFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZeroXOTCSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZeroXOTCSession struct {
	Contract     *ZeroXOTC         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZeroXOTCCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZeroXOTCCallerSession struct {
	Contract *ZeroXOTCCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ZeroXOTCTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZeroXOTCTransactorSession struct {
	Contract     *ZeroXOTCTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ZeroXOTCRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZeroXOTCRaw struct {
	Contract *ZeroXOTC // Generic contract binding to access the raw methods on
}

// ZeroXOTCCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZeroXOTCCallerRaw struct {
	Contract *ZeroXOTCCaller // Generic read-only contract binding to access the raw methods on
}

// ZeroXOTCTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZeroXOTCTransactorRaw struct {
	Contract *ZeroXOTCTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZeroXOTC creates a new instance of ZeroXOTC, bound to a specific deployed contract.
func NewZeroXOTC(address common.Address, backend bind.ContractBackend) (*ZeroXOTC, error) {
	contract, err := bindZeroXOTC(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZeroXOTC{ZeroXOTCCaller: ZeroXOTCCaller{contract: contract}, ZeroXOTCTransactor: ZeroXOTCTransactor{contract: contract}, ZeroXOTCFilterer: ZeroXOTCFilterer{contract: contract}}, nil
}

// NewZeroXOTCCaller creates a new read-only instance of ZeroXOTC, bound to a specific deployed contract.
func NewZeroXOTCCaller(address common.Address, caller bind.ContractCaller) (*ZeroXOTCCaller, error) {
	contract, err := bindZeroXOTC(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZeroXOTCCaller{contract: contract}, nil
}

// NewZeroXOTCTransactor creates a new write-only instance of ZeroXOTC, bound to a specific deployed contract.
func NewZeroXOTCTransactor(address common.Address, transactor bind.ContractTransactor) (*ZeroXOTCTransactor, error) {
	contract, err := bindZeroXOTC(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZeroXOTCTransactor{contract: contract}, nil
}

// NewZeroXOTCFilterer creates a new log filterer instance of ZeroXOTC, bound to a specific deployed contract.
func NewZeroXOTCFilterer(address common.Address, filterer bind.ContractFilterer) (*ZeroXOTCFilterer, error) {
	contract, err := bindZeroXOTC(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZeroXOTCFilterer{contract: contract}, nil
}

// bindZeroXOTC binds a generic wrapper to an already deployed contract.
func bindZeroXOTC(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ZeroXOTCMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZeroXOTC *ZeroXOTCRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZeroXOTC.Contract.ZeroXOTCCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZeroXOTC *ZeroXOTCRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZeroXOTC.Contract.ZeroXOTCTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZeroXOTC *ZeroXOTCRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZeroXOTC.Contract.ZeroXOTCTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZeroXOTC *ZeroXOTCCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZeroXOTC.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZeroXOTC *ZeroXOTCTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZeroXOTC.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZeroXOTC *ZeroXOTCTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZeroXOTC.Contract.contract.Transact(opts, method, params...)
}

// EIP712DOMAINSEPARATOR is a free data retrieval call binding the contract method 0xdab400f3.
//
// Solidity: function EIP712_DOMAIN_SEPARATOR() view returns(bytes32)
func (_ZeroXOTC *ZeroXOTCCaller) EIP712DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ZeroXOTC.contract.Call(opts, &out, "EIP712_DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// EIP712DOMAINSEPARATOR is a free data retrieval call binding the contract method 0xdab400f3.
//
// Solidity: function EIP712_DOMAIN_SEPARATOR() view returns(bytes32)
func (_ZeroXOTC *ZeroXOTCSession) EIP712DOMAINSEPARATOR() ([32]byte, error) {
	return _ZeroXOTC.Contract.EIP712DOMAINSEPARATOR(&_ZeroXOTC.CallOpts)
}

// EIP712DOMAINSEPARATOR is a free data retrieval call binding the contract method 0xdab400f3.
//
// Solidity: function EIP712_DOMAIN_SEPARATOR() view returns(bytes32)
func (_ZeroXOTC *ZeroXOTCCallerSession) EIP712DOMAINSEPARATOR() ([32]byte, error) {
	return _ZeroXOTC.Contract.EIP712DOMAINSEPARATOR(&_ZeroXOTC.CallOpts)
}

// FEATURENAME is a free data retrieval call binding the contract method 0x6ae4b4f7.
//
// Solidity: function FEATURE_NAME() view returns(string)
func (_ZeroXOTC *ZeroXOTCCaller) FEATURENAME(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ZeroXOTC.contract.Call(opts, &out, "FEATURE_NAME")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// FEATURENAME is a free data retrieval call binding the contract method 0x6ae4b4f7.
//
// Solidity: function FEATURE_NAME() view returns(string)
func (_ZeroXOTC *ZeroXOTCSession) FEATURENAME() (string, error) {
	return _ZeroXOTC.Contract.FEATURENAME(&_ZeroXOTC.CallOpts)
}

// FEATURENAME is a free data retrieval call binding the contract method 0x6ae4b4f7.
//
// Solidity: function FEATURE_NAME() view returns(string)
func (_ZeroXOTC *ZeroXOTCCallerSession) FEATURENAME() (string, error) {
	return _ZeroXOTC.Contract.FEATURENAME(&_ZeroXOTC.CallOpts)
}

// FEATUREVERSION is a free data retrieval call binding the contract method 0x031b905c.
//
// Solidity: function FEATURE_VERSION() view returns(uint256)
func (_ZeroXOTC *ZeroXOTCCaller) FEATUREVERSION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ZeroXOTC.contract.Call(opts, &out, "FEATURE_VERSION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FEATUREVERSION is a free data retrieval call binding the contract method 0x031b905c.
//
// Solidity: function FEATURE_VERSION() view returns(uint256)
func (_ZeroXOTC *ZeroXOTCSession) FEATUREVERSION() (*big.Int, error) {
	return _ZeroXOTC.Contract.FEATUREVERSION(&_ZeroXOTC.CallOpts)
}

// FEATUREVERSION is a free data retrieval call binding the contract method 0x031b905c.
//
// Solidity: function FEATURE_VERSION() view returns(uint256)
func (_ZeroXOTC *ZeroXOTCCallerSession) FEATUREVERSION() (*big.Int, error) {
	return _ZeroXOTC.Contract.FEATUREVERSION(&_ZeroXOTC.CallOpts)
}

// GetOtcOrderHash is a free data retrieval call binding the contract method 0x53476b89.
//
// Solidity: function getOtcOrderHash((address,address,uint128,uint128,address,address,address,uint256) order) view returns(bytes32 orderHash)
func (_ZeroXOTC *ZeroXOTCCaller) GetOtcOrderHash(opts *bind.CallOpts, order LibNativeOrderOtcOrder) ([32]byte, error) {
	var out []interface{}
	err := _ZeroXOTC.contract.Call(opts, &out, "getOtcOrderHash", order)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetOtcOrderHash is a free data retrieval call binding the contract method 0x53476b89.
//
// Solidity: function getOtcOrderHash((address,address,uint128,uint128,address,address,address,uint256) order) view returns(bytes32 orderHash)
func (_ZeroXOTC *ZeroXOTCSession) GetOtcOrderHash(order LibNativeOrderOtcOrder) ([32]byte, error) {
	return _ZeroXOTC.Contract.GetOtcOrderHash(&_ZeroXOTC.CallOpts, order)
}

// GetOtcOrderHash is a free data retrieval call binding the contract method 0x53476b89.
//
// Solidity: function getOtcOrderHash((address,address,uint128,uint128,address,address,address,uint256) order) view returns(bytes32 orderHash)
func (_ZeroXOTC *ZeroXOTCCallerSession) GetOtcOrderHash(order LibNativeOrderOtcOrder) ([32]byte, error) {
	return _ZeroXOTC.Contract.GetOtcOrderHash(&_ZeroXOTC.CallOpts, order)
}

// GetOtcOrderInfo is a free data retrieval call binding the contract method 0x8c807c43.
//
// Solidity: function getOtcOrderInfo((address,address,uint128,uint128,address,address,address,uint256) order) view returns((bytes32,uint8) orderInfo)
func (_ZeroXOTC *ZeroXOTCCaller) GetOtcOrderInfo(opts *bind.CallOpts, order LibNativeOrderOtcOrder) (LibNativeOrderOtcOrderInfo, error) {
	var out []interface{}
	err := _ZeroXOTC.contract.Call(opts, &out, "getOtcOrderInfo", order)

	if err != nil {
		return *new(LibNativeOrderOtcOrderInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(LibNativeOrderOtcOrderInfo)).(*LibNativeOrderOtcOrderInfo)

	return out0, err

}

// GetOtcOrderInfo is a free data retrieval call binding the contract method 0x8c807c43.
//
// Solidity: function getOtcOrderInfo((address,address,uint128,uint128,address,address,address,uint256) order) view returns((bytes32,uint8) orderInfo)
func (_ZeroXOTC *ZeroXOTCSession) GetOtcOrderInfo(order LibNativeOrderOtcOrder) (LibNativeOrderOtcOrderInfo, error) {
	return _ZeroXOTC.Contract.GetOtcOrderInfo(&_ZeroXOTC.CallOpts, order)
}

// GetOtcOrderInfo is a free data retrieval call binding the contract method 0x8c807c43.
//
// Solidity: function getOtcOrderInfo((address,address,uint128,uint128,address,address,address,uint256) order) view returns((bytes32,uint8) orderInfo)
func (_ZeroXOTC *ZeroXOTCCallerSession) GetOtcOrderInfo(order LibNativeOrderOtcOrder) (LibNativeOrderOtcOrderInfo, error) {
	return _ZeroXOTC.Contract.GetOtcOrderInfo(&_ZeroXOTC.CallOpts, order)
}

// LastOtcTxOriginNonce is a free data retrieval call binding the contract method 0x59ebfb45.
//
// Solidity: function lastOtcTxOriginNonce(address txOrigin, uint64 nonceBucket) view returns(uint128 lastNonce)
func (_ZeroXOTC *ZeroXOTCCaller) LastOtcTxOriginNonce(opts *bind.CallOpts, txOrigin common.Address, nonceBucket uint64) (*big.Int, error) {
	var out []interface{}
	err := _ZeroXOTC.contract.Call(opts, &out, "lastOtcTxOriginNonce", txOrigin, nonceBucket)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastOtcTxOriginNonce is a free data retrieval call binding the contract method 0x59ebfb45.
//
// Solidity: function lastOtcTxOriginNonce(address txOrigin, uint64 nonceBucket) view returns(uint128 lastNonce)
func (_ZeroXOTC *ZeroXOTCSession) LastOtcTxOriginNonce(txOrigin common.Address, nonceBucket uint64) (*big.Int, error) {
	return _ZeroXOTC.Contract.LastOtcTxOriginNonce(&_ZeroXOTC.CallOpts, txOrigin, nonceBucket)
}

// LastOtcTxOriginNonce is a free data retrieval call binding the contract method 0x59ebfb45.
//
// Solidity: function lastOtcTxOriginNonce(address txOrigin, uint64 nonceBucket) view returns(uint128 lastNonce)
func (_ZeroXOTC *ZeroXOTCCallerSession) LastOtcTxOriginNonce(txOrigin common.Address, nonceBucket uint64) (*big.Int, error) {
	return _ZeroXOTC.Contract.LastOtcTxOriginNonce(&_ZeroXOTC.CallOpts, txOrigin, nonceBucket)
}

// FillOtcOrder is a paid mutator transaction binding the contract method 0xe4ba8439.
//
// Solidity: function _fillOtcOrder((address,address,uint128,uint128,address,address,address,uint256) order, (uint8,uint8,bytes32,bytes32) makerSignature, uint128 takerTokenFillAmount, address taker, bool useSelfBalance, address recipient) returns(uint128 takerTokenFilledAmount, uint128 makerTokenFilledAmount)
func (_ZeroXOTC *ZeroXOTCTransactor) FillOtcOrder(opts *bind.TransactOpts, order LibNativeOrderOtcOrder, makerSignature LibSignatureSignature, takerTokenFillAmount *big.Int, taker common.Address, useSelfBalance bool, recipient common.Address) (*types.Transaction, error) {
	return _ZeroXOTC.contract.Transact(opts, "_fillOtcOrder", order, makerSignature, takerTokenFillAmount, taker, useSelfBalance, recipient)
}

// FillOtcOrder is a paid mutator transaction binding the contract method 0xe4ba8439.
//
// Solidity: function _fillOtcOrder((address,address,uint128,uint128,address,address,address,uint256) order, (uint8,uint8,bytes32,bytes32) makerSignature, uint128 takerTokenFillAmount, address taker, bool useSelfBalance, address recipient) returns(uint128 takerTokenFilledAmount, uint128 makerTokenFilledAmount)
func (_ZeroXOTC *ZeroXOTCSession) FillOtcOrder(order LibNativeOrderOtcOrder, makerSignature LibSignatureSignature, takerTokenFillAmount *big.Int, taker common.Address, useSelfBalance bool, recipient common.Address) (*types.Transaction, error) {
	return _ZeroXOTC.Contract.FillOtcOrder(&_ZeroXOTC.TransactOpts, order, makerSignature, takerTokenFillAmount, taker, useSelfBalance, recipient)
}

// FillOtcOrder is a paid mutator transaction binding the contract method 0xe4ba8439.
//
// Solidity: function _fillOtcOrder((address,address,uint128,uint128,address,address,address,uint256) order, (uint8,uint8,bytes32,bytes32) makerSignature, uint128 takerTokenFillAmount, address taker, bool useSelfBalance, address recipient) returns(uint128 takerTokenFilledAmount, uint128 makerTokenFilledAmount)
func (_ZeroXOTC *ZeroXOTCTransactorSession) FillOtcOrder(order LibNativeOrderOtcOrder, makerSignature LibSignatureSignature, takerTokenFillAmount *big.Int, taker common.Address, useSelfBalance bool, recipient common.Address) (*types.Transaction, error) {
	return _ZeroXOTC.Contract.FillOtcOrder(&_ZeroXOTC.TransactOpts, order, makerSignature, takerTokenFillAmount, taker, useSelfBalance, recipient)
}

// BatchFillTakerSignedOtcOrders is a paid mutator transaction binding the contract method 0xe52d1764.
//
// Solidity: function batchFillTakerSignedOtcOrders((address,address,uint128,uint128,address,address,address,uint256)[] orders, (uint8,uint8,bytes32,bytes32)[] makerSignatures, (uint8,uint8,bytes32,bytes32)[] takerSignatures, bool[] unwrapWeth) returns(bool[] successes)
func (_ZeroXOTC *ZeroXOTCTransactor) BatchFillTakerSignedOtcOrders(opts *bind.TransactOpts, orders []LibNativeOrderOtcOrder, makerSignatures []LibSignatureSignature, takerSignatures []LibSignatureSignature, unwrapWeth []bool) (*types.Transaction, error) {
	return _ZeroXOTC.contract.Transact(opts, "batchFillTakerSignedOtcOrders", orders, makerSignatures, takerSignatures, unwrapWeth)
}

// BatchFillTakerSignedOtcOrders is a paid mutator transaction binding the contract method 0xe52d1764.
//
// Solidity: function batchFillTakerSignedOtcOrders((address,address,uint128,uint128,address,address,address,uint256)[] orders, (uint8,uint8,bytes32,bytes32)[] makerSignatures, (uint8,uint8,bytes32,bytes32)[] takerSignatures, bool[] unwrapWeth) returns(bool[] successes)
func (_ZeroXOTC *ZeroXOTCSession) BatchFillTakerSignedOtcOrders(orders []LibNativeOrderOtcOrder, makerSignatures []LibSignatureSignature, takerSignatures []LibSignatureSignature, unwrapWeth []bool) (*types.Transaction, error) {
	return _ZeroXOTC.Contract.BatchFillTakerSignedOtcOrders(&_ZeroXOTC.TransactOpts, orders, makerSignatures, takerSignatures, unwrapWeth)
}

// BatchFillTakerSignedOtcOrders is a paid mutator transaction binding the contract method 0xe52d1764.
//
// Solidity: function batchFillTakerSignedOtcOrders((address,address,uint128,uint128,address,address,address,uint256)[] orders, (uint8,uint8,bytes32,bytes32)[] makerSignatures, (uint8,uint8,bytes32,bytes32)[] takerSignatures, bool[] unwrapWeth) returns(bool[] successes)
func (_ZeroXOTC *ZeroXOTCTransactorSession) BatchFillTakerSignedOtcOrders(orders []LibNativeOrderOtcOrder, makerSignatures []LibSignatureSignature, takerSignatures []LibSignatureSignature, unwrapWeth []bool) (*types.Transaction, error) {
	return _ZeroXOTC.Contract.BatchFillTakerSignedOtcOrders(&_ZeroXOTC.TransactOpts, orders, makerSignatures, takerSignatures, unwrapWeth)
}

// FillOtcOrder1 is a paid mutator transaction binding the contract method 0xdac748d4.
//
// Solidity: function fillOtcOrder((address,address,uint128,uint128,address,address,address,uint256) order, (uint8,uint8,bytes32,bytes32) makerSignature, uint128 takerTokenFillAmount) returns(uint128 takerTokenFilledAmount, uint128 makerTokenFilledAmount)
func (_ZeroXOTC *ZeroXOTCTransactor) FillOtcOrder1(opts *bind.TransactOpts, order LibNativeOrderOtcOrder, makerSignature LibSignatureSignature, takerTokenFillAmount *big.Int) (*types.Transaction, error) {
	return _ZeroXOTC.contract.Transact(opts, "fillOtcOrder", order, makerSignature, takerTokenFillAmount)
}

// FillOtcOrder1 is a paid mutator transaction binding the contract method 0xdac748d4.
//
// Solidity: function fillOtcOrder((address,address,uint128,uint128,address,address,address,uint256) order, (uint8,uint8,bytes32,bytes32) makerSignature, uint128 takerTokenFillAmount) returns(uint128 takerTokenFilledAmount, uint128 makerTokenFilledAmount)
func (_ZeroXOTC *ZeroXOTCSession) FillOtcOrder1(order LibNativeOrderOtcOrder, makerSignature LibSignatureSignature, takerTokenFillAmount *big.Int) (*types.Transaction, error) {
	return _ZeroXOTC.Contract.FillOtcOrder1(&_ZeroXOTC.TransactOpts, order, makerSignature, takerTokenFillAmount)
}

// FillOtcOrder1 is a paid mutator transaction binding the contract method 0xdac748d4.
//
// Solidity: function fillOtcOrder((address,address,uint128,uint128,address,address,address,uint256) order, (uint8,uint8,bytes32,bytes32) makerSignature, uint128 takerTokenFillAmount) returns(uint128 takerTokenFilledAmount, uint128 makerTokenFilledAmount)
func (_ZeroXOTC *ZeroXOTCTransactorSession) FillOtcOrder1(order LibNativeOrderOtcOrder, makerSignature LibSignatureSignature, takerTokenFillAmount *big.Int) (*types.Transaction, error) {
	return _ZeroXOTC.Contract.FillOtcOrder1(&_ZeroXOTC.TransactOpts, order, makerSignature, takerTokenFillAmount)
}

// FillOtcOrderForEth is a paid mutator transaction binding the contract method 0xa578efaf.
//
// Solidity: function fillOtcOrderForEth((address,address,uint128,uint128,address,address,address,uint256) order, (uint8,uint8,bytes32,bytes32) makerSignature, uint128 takerTokenFillAmount) returns(uint128 takerTokenFilledAmount, uint128 makerTokenFilledAmount)
func (_ZeroXOTC *ZeroXOTCTransactor) FillOtcOrderForEth(opts *bind.TransactOpts, order LibNativeOrderOtcOrder, makerSignature LibSignatureSignature, takerTokenFillAmount *big.Int) (*types.Transaction, error) {
	return _ZeroXOTC.contract.Transact(opts, "fillOtcOrderForEth", order, makerSignature, takerTokenFillAmount)
}

// FillOtcOrderForEth is a paid mutator transaction binding the contract method 0xa578efaf.
//
// Solidity: function fillOtcOrderForEth((address,address,uint128,uint128,address,address,address,uint256) order, (uint8,uint8,bytes32,bytes32) makerSignature, uint128 takerTokenFillAmount) returns(uint128 takerTokenFilledAmount, uint128 makerTokenFilledAmount)
func (_ZeroXOTC *ZeroXOTCSession) FillOtcOrderForEth(order LibNativeOrderOtcOrder, makerSignature LibSignatureSignature, takerTokenFillAmount *big.Int) (*types.Transaction, error) {
	return _ZeroXOTC.Contract.FillOtcOrderForEth(&_ZeroXOTC.TransactOpts, order, makerSignature, takerTokenFillAmount)
}

// FillOtcOrderForEth is a paid mutator transaction binding the contract method 0xa578efaf.
//
// Solidity: function fillOtcOrderForEth((address,address,uint128,uint128,address,address,address,uint256) order, (uint8,uint8,bytes32,bytes32) makerSignature, uint128 takerTokenFillAmount) returns(uint128 takerTokenFilledAmount, uint128 makerTokenFilledAmount)
func (_ZeroXOTC *ZeroXOTCTransactorSession) FillOtcOrderForEth(order LibNativeOrderOtcOrder, makerSignature LibSignatureSignature, takerTokenFillAmount *big.Int) (*types.Transaction, error) {
	return _ZeroXOTC.Contract.FillOtcOrderForEth(&_ZeroXOTC.TransactOpts, order, makerSignature, takerTokenFillAmount)
}

// FillOtcOrderWithEth is a paid mutator transaction binding the contract method 0x706394d5.
//
// Solidity: function fillOtcOrderWithEth((address,address,uint128,uint128,address,address,address,uint256) order, (uint8,uint8,bytes32,bytes32) makerSignature) payable returns(uint128 takerTokenFilledAmount, uint128 makerTokenFilledAmount)
func (_ZeroXOTC *ZeroXOTCTransactor) FillOtcOrderWithEth(opts *bind.TransactOpts, order LibNativeOrderOtcOrder, makerSignature LibSignatureSignature) (*types.Transaction, error) {
	return _ZeroXOTC.contract.Transact(opts, "fillOtcOrderWithEth", order, makerSignature)
}

// FillOtcOrderWithEth is a paid mutator transaction binding the contract method 0x706394d5.
//
// Solidity: function fillOtcOrderWithEth((address,address,uint128,uint128,address,address,address,uint256) order, (uint8,uint8,bytes32,bytes32) makerSignature) payable returns(uint128 takerTokenFilledAmount, uint128 makerTokenFilledAmount)
func (_ZeroXOTC *ZeroXOTCSession) FillOtcOrderWithEth(order LibNativeOrderOtcOrder, makerSignature LibSignatureSignature) (*types.Transaction, error) {
	return _ZeroXOTC.Contract.FillOtcOrderWithEth(&_ZeroXOTC.TransactOpts, order, makerSignature)
}

// FillOtcOrderWithEth is a paid mutator transaction binding the contract method 0x706394d5.
//
// Solidity: function fillOtcOrderWithEth((address,address,uint128,uint128,address,address,address,uint256) order, (uint8,uint8,bytes32,bytes32) makerSignature) payable returns(uint128 takerTokenFilledAmount, uint128 makerTokenFilledAmount)
func (_ZeroXOTC *ZeroXOTCTransactorSession) FillOtcOrderWithEth(order LibNativeOrderOtcOrder, makerSignature LibSignatureSignature) (*types.Transaction, error) {
	return _ZeroXOTC.Contract.FillOtcOrderWithEth(&_ZeroXOTC.TransactOpts, order, makerSignature)
}

// FillTakerSignedOtcOrder is a paid mutator transaction binding the contract method 0x4f948110.
//
// Solidity: function fillTakerSignedOtcOrder((address,address,uint128,uint128,address,address,address,uint256) order, (uint8,uint8,bytes32,bytes32) makerSignature, (uint8,uint8,bytes32,bytes32) takerSignature) returns()
func (_ZeroXOTC *ZeroXOTCTransactor) FillTakerSignedOtcOrder(opts *bind.TransactOpts, order LibNativeOrderOtcOrder, makerSignature LibSignatureSignature, takerSignature LibSignatureSignature) (*types.Transaction, error) {
	return _ZeroXOTC.contract.Transact(opts, "fillTakerSignedOtcOrder", order, makerSignature, takerSignature)
}

// FillTakerSignedOtcOrder is a paid mutator transaction binding the contract method 0x4f948110.
//
// Solidity: function fillTakerSignedOtcOrder((address,address,uint128,uint128,address,address,address,uint256) order, (uint8,uint8,bytes32,bytes32) makerSignature, (uint8,uint8,bytes32,bytes32) takerSignature) returns()
func (_ZeroXOTC *ZeroXOTCSession) FillTakerSignedOtcOrder(order LibNativeOrderOtcOrder, makerSignature LibSignatureSignature, takerSignature LibSignatureSignature) (*types.Transaction, error) {
	return _ZeroXOTC.Contract.FillTakerSignedOtcOrder(&_ZeroXOTC.TransactOpts, order, makerSignature, takerSignature)
}

// FillTakerSignedOtcOrder is a paid mutator transaction binding the contract method 0x4f948110.
//
// Solidity: function fillTakerSignedOtcOrder((address,address,uint128,uint128,address,address,address,uint256) order, (uint8,uint8,bytes32,bytes32) makerSignature, (uint8,uint8,bytes32,bytes32) takerSignature) returns()
func (_ZeroXOTC *ZeroXOTCTransactorSession) FillTakerSignedOtcOrder(order LibNativeOrderOtcOrder, makerSignature LibSignatureSignature, takerSignature LibSignatureSignature) (*types.Transaction, error) {
	return _ZeroXOTC.Contract.FillTakerSignedOtcOrder(&_ZeroXOTC.TransactOpts, order, makerSignature, takerSignature)
}

// FillTakerSignedOtcOrderForEth is a paid mutator transaction binding the contract method 0x724d3953.
//
// Solidity: function fillTakerSignedOtcOrderForEth((address,address,uint128,uint128,address,address,address,uint256) order, (uint8,uint8,bytes32,bytes32) makerSignature, (uint8,uint8,bytes32,bytes32) takerSignature) returns()
func (_ZeroXOTC *ZeroXOTCTransactor) FillTakerSignedOtcOrderForEth(opts *bind.TransactOpts, order LibNativeOrderOtcOrder, makerSignature LibSignatureSignature, takerSignature LibSignatureSignature) (*types.Transaction, error) {
	return _ZeroXOTC.contract.Transact(opts, "fillTakerSignedOtcOrderForEth", order, makerSignature, takerSignature)
}

// FillTakerSignedOtcOrderForEth is a paid mutator transaction binding the contract method 0x724d3953.
//
// Solidity: function fillTakerSignedOtcOrderForEth((address,address,uint128,uint128,address,address,address,uint256) order, (uint8,uint8,bytes32,bytes32) makerSignature, (uint8,uint8,bytes32,bytes32) takerSignature) returns()
func (_ZeroXOTC *ZeroXOTCSession) FillTakerSignedOtcOrderForEth(order LibNativeOrderOtcOrder, makerSignature LibSignatureSignature, takerSignature LibSignatureSignature) (*types.Transaction, error) {
	return _ZeroXOTC.Contract.FillTakerSignedOtcOrderForEth(&_ZeroXOTC.TransactOpts, order, makerSignature, takerSignature)
}

// FillTakerSignedOtcOrderForEth is a paid mutator transaction binding the contract method 0x724d3953.
//
// Solidity: function fillTakerSignedOtcOrderForEth((address,address,uint128,uint128,address,address,address,uint256) order, (uint8,uint8,bytes32,bytes32) makerSignature, (uint8,uint8,bytes32,bytes32) takerSignature) returns()
func (_ZeroXOTC *ZeroXOTCTransactorSession) FillTakerSignedOtcOrderForEth(order LibNativeOrderOtcOrder, makerSignature LibSignatureSignature, takerSignature LibSignatureSignature) (*types.Transaction, error) {
	return _ZeroXOTC.Contract.FillTakerSignedOtcOrderForEth(&_ZeroXOTC.TransactOpts, order, makerSignature, takerSignature)
}

// Migrate is a paid mutator transaction binding the contract method 0x8fd3ab80.
//
// Solidity: function migrate() returns(bytes4 success)
func (_ZeroXOTC *ZeroXOTCTransactor) Migrate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZeroXOTC.contract.Transact(opts, "migrate")
}

// Migrate is a paid mutator transaction binding the contract method 0x8fd3ab80.
//
// Solidity: function migrate() returns(bytes4 success)
func (_ZeroXOTC *ZeroXOTCSession) Migrate() (*types.Transaction, error) {
	return _ZeroXOTC.Contract.Migrate(&_ZeroXOTC.TransactOpts)
}

// Migrate is a paid mutator transaction binding the contract method 0x8fd3ab80.
//
// Solidity: function migrate() returns(bytes4 success)
func (_ZeroXOTC *ZeroXOTCTransactorSession) Migrate() (*types.Transaction, error) {
	return _ZeroXOTC.Contract.Migrate(&_ZeroXOTC.TransactOpts)
}

// ZeroXOTCOtcOrderFilledIterator is returned from FilterOtcOrderFilled and is used to iterate over the raw logs and unpacked data for OtcOrderFilled events raised by the ZeroXOTC contract.
type ZeroXOTCOtcOrderFilledIterator struct {
	Event *ZeroXOTCOtcOrderFilled // Event containing the contract specifics and raw log

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
func (it *ZeroXOTCOtcOrderFilledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZeroXOTCOtcOrderFilled)
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
		it.Event = new(ZeroXOTCOtcOrderFilled)
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
func (it *ZeroXOTCOtcOrderFilledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZeroXOTCOtcOrderFilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZeroXOTCOtcOrderFilled represents a OtcOrderFilled event raised by the ZeroXOTC contract.
type ZeroXOTCOtcOrderFilled struct {
	OrderHash              [32]byte
	Maker                  common.Address
	Taker                  common.Address
	MakerToken             common.Address
	TakerToken             common.Address
	MakerTokenFilledAmount *big.Int
	TakerTokenFilledAmount *big.Int
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterOtcOrderFilled is a free log retrieval operation binding the contract event 0xac75f773e3a92f1a02b12134d65e1f47f8a14eabe4eaf1e24624918e6a8b269f.
//
// Solidity: event OtcOrderFilled(bytes32 orderHash, address maker, address taker, address makerToken, address takerToken, uint128 makerTokenFilledAmount, uint128 takerTokenFilledAmount)
func (_ZeroXOTC *ZeroXOTCFilterer) FilterOtcOrderFilled(opts *bind.FilterOpts) (*ZeroXOTCOtcOrderFilledIterator, error) {

	logs, sub, err := _ZeroXOTC.contract.FilterLogs(opts, "OtcOrderFilled")
	if err != nil {
		return nil, err
	}
	return &ZeroXOTCOtcOrderFilledIterator{contract: _ZeroXOTC.contract, event: "OtcOrderFilled", logs: logs, sub: sub}, nil
}

// WatchOtcOrderFilled is a free log subscription operation binding the contract event 0xac75f773e3a92f1a02b12134d65e1f47f8a14eabe4eaf1e24624918e6a8b269f.
//
// Solidity: event OtcOrderFilled(bytes32 orderHash, address maker, address taker, address makerToken, address takerToken, uint128 makerTokenFilledAmount, uint128 takerTokenFilledAmount)
func (_ZeroXOTC *ZeroXOTCFilterer) WatchOtcOrderFilled(opts *bind.WatchOpts, sink chan<- *ZeroXOTCOtcOrderFilled) (event.Subscription, error) {

	logs, sub, err := _ZeroXOTC.contract.WatchLogs(opts, "OtcOrderFilled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZeroXOTCOtcOrderFilled)
				if err := _ZeroXOTC.contract.UnpackLog(event, "OtcOrderFilled", log); err != nil {
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

// ParseOtcOrderFilled is a log parse operation binding the contract event 0xac75f773e3a92f1a02b12134d65e1f47f8a14eabe4eaf1e24624918e6a8b269f.
//
// Solidity: event OtcOrderFilled(bytes32 orderHash, address maker, address taker, address makerToken, address takerToken, uint128 makerTokenFilledAmount, uint128 takerTokenFilledAmount)
func (_ZeroXOTC *ZeroXOTCFilterer) ParseOtcOrderFilled(log types.Log) (*ZeroXOTCOtcOrderFilled, error) {
	event := new(ZeroXOTCOtcOrderFilled)
	if err := _ZeroXOTC.contract.UnpackLog(event, "OtcOrderFilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
