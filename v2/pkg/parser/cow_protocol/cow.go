// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package cowprotocol

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

// GPv2InteractionData is an auto generated low-level Go binding around an user-defined struct.
type GPv2InteractionData struct {
	Target   common.Address
	Value    *big.Int
	CallData []byte
}

// GPv2TradeData is an auto generated low-level Go binding around an user-defined struct.
type GPv2TradeData struct {
	SellTokenIndex *big.Int
	BuyTokenIndex  *big.Int
	Receiver       common.Address
	SellAmount     *big.Int
	BuyAmount      *big.Int
	ValidTo        uint32
	AppData        [32]byte
	FeeAmount      *big.Int
	Flags          *big.Int
	ExecutedAmount *big.Int
	Signature      []byte
}

// IVaultBatchSwapStep is an auto generated low-level Go binding around an user-defined struct.
type IVaultBatchSwapStep struct {
	PoolId        [32]byte
	AssetInIndex  *big.Int
	AssetOutIndex *big.Int
	Amount        *big.Int
	UserData      []byte
}

// CowProtocolMetaData contains all meta data concerning the CowProtocol contract.
var CowProtocolMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractGPv2Authentication\",\"name\":\"authenticator_\",\"type\":\"address\"},{\"internalType\":\"contractIVault\",\"name\":\"vault_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"}],\"name\":\"Interaction\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"orderUid\",\"type\":\"bytes\"}],\"name\":\"OrderInvalidated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"orderUid\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"signed\",\"type\":\"bool\"}],\"name\":\"PreSignature\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"solver\",\"type\":\"address\"}],\"name\":\"Settlement\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"sellToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"buyToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sellAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"buyAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"orderUid\",\"type\":\"bytes\"}],\"name\":\"Trade\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"authenticator\",\"outputs\":[{\"internalType\":\"contractGPv2Authentication\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"domainSeparator\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"filledAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"orderUids\",\"type\":\"bytes[]\"}],\"name\":\"freeFilledAmountStorage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"orderUids\",\"type\":\"bytes[]\"}],\"name\":\"freePreSignatureStorage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"getStorageAt\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"orderUid\",\"type\":\"bytes\"}],\"name\":\"invalidateOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"preSignature\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"orderUid\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"signed\",\"type\":\"bool\"}],\"name\":\"setPreSignature\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"clearingPrices\",\"type\":\"uint256[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sellTokenIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"buyTokenIndex\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"sellAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"buyAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"validTo\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"appData\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"flags\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"executedAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structGPv2Trade.Data[]\",\"name\":\"trades\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structGPv2Interaction.Data[][3]\",\"name\":\"interactions\",\"type\":\"tuple[][3]\"}],\"name\":\"settle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"targetContract\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"calldataPayload\",\"type\":\"bytes\"}],\"name\":\"simulateDelegatecall\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"response\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"targetContract\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"calldataPayload\",\"type\":\"bytes\"}],\"name\":\"simulateDelegatecallInternal\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"response\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"poolId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"assetInIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetOutIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"userData\",\"type\":\"bytes\"}],\"internalType\":\"structIVault.BatchSwapStep[]\",\"name\":\"swaps\",\"type\":\"tuple[]\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sellTokenIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"buyTokenIndex\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"sellAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"buyAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"validTo\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"appData\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"flags\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"executedAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structGPv2Trade.Data\",\"name\":\"trade\",\"type\":\"tuple\"}],\"name\":\"swap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vault\",\"outputs\":[{\"internalType\":\"contractIVault\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vaultRelayer\",\"outputs\":[{\"internalType\":\"contractGPv2VaultRelayer\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// CowProtocolABI is the input ABI used to generate the binding from.
// Deprecated: Use CowProtocolMetaData.ABI instead.
var CowProtocolABI = CowProtocolMetaData.ABI

// CowProtocol is an auto generated Go binding around an Ethereum contract.
type CowProtocol struct {
	CowProtocolCaller     // Read-only binding to the contract
	CowProtocolTransactor // Write-only binding to the contract
	CowProtocolFilterer   // Log filterer for contract events
}

// CowProtocolCaller is an auto generated read-only Go binding around an Ethereum contract.
type CowProtocolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CowProtocolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CowProtocolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CowProtocolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CowProtocolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CowProtocolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CowProtocolSession struct {
	Contract     *CowProtocol      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CowProtocolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CowProtocolCallerSession struct {
	Contract *CowProtocolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// CowProtocolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CowProtocolTransactorSession struct {
	Contract     *CowProtocolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// CowProtocolRaw is an auto generated low-level Go binding around an Ethereum contract.
type CowProtocolRaw struct {
	Contract *CowProtocol // Generic contract binding to access the raw methods on
}

// CowProtocolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CowProtocolCallerRaw struct {
	Contract *CowProtocolCaller // Generic read-only contract binding to access the raw methods on
}

// CowProtocolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CowProtocolTransactorRaw struct {
	Contract *CowProtocolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCowProtocol creates a new instance of CowProtocol, bound to a specific deployed contract.
func NewCowProtocol(address common.Address, backend bind.ContractBackend) (*CowProtocol, error) {
	contract, err := bindCowProtocol(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CowProtocol{CowProtocolCaller: CowProtocolCaller{contract: contract}, CowProtocolTransactor: CowProtocolTransactor{contract: contract}, CowProtocolFilterer: CowProtocolFilterer{contract: contract}}, nil
}

// NewCowProtocolCaller creates a new read-only instance of CowProtocol, bound to a specific deployed contract.
func NewCowProtocolCaller(address common.Address, caller bind.ContractCaller) (*CowProtocolCaller, error) {
	contract, err := bindCowProtocol(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CowProtocolCaller{contract: contract}, nil
}

// NewCowProtocolTransactor creates a new write-only instance of CowProtocol, bound to a specific deployed contract.
func NewCowProtocolTransactor(address common.Address, transactor bind.ContractTransactor) (*CowProtocolTransactor, error) {
	contract, err := bindCowProtocol(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CowProtocolTransactor{contract: contract}, nil
}

// NewCowProtocolFilterer creates a new log filterer instance of CowProtocol, bound to a specific deployed contract.
func NewCowProtocolFilterer(address common.Address, filterer bind.ContractFilterer) (*CowProtocolFilterer, error) {
	contract, err := bindCowProtocol(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CowProtocolFilterer{contract: contract}, nil
}

// bindCowProtocol binds a generic wrapper to an already deployed contract.
func bindCowProtocol(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CowProtocolMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CowProtocol *CowProtocolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CowProtocol.Contract.CowProtocolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CowProtocol *CowProtocolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CowProtocol.Contract.CowProtocolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CowProtocol *CowProtocolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CowProtocol.Contract.CowProtocolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CowProtocol *CowProtocolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CowProtocol.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CowProtocol *CowProtocolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CowProtocol.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CowProtocol *CowProtocolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CowProtocol.Contract.contract.Transact(opts, method, params...)
}

// Authenticator is a free data retrieval call binding the contract method 0x2335c76b.
//
// Solidity: function authenticator() view returns(address)
func (_CowProtocol *CowProtocolCaller) Authenticator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CowProtocol.contract.Call(opts, &out, "authenticator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Authenticator is a free data retrieval call binding the contract method 0x2335c76b.
//
// Solidity: function authenticator() view returns(address)
func (_CowProtocol *CowProtocolSession) Authenticator() (common.Address, error) {
	return _CowProtocol.Contract.Authenticator(&_CowProtocol.CallOpts)
}

// Authenticator is a free data retrieval call binding the contract method 0x2335c76b.
//
// Solidity: function authenticator() view returns(address)
func (_CowProtocol *CowProtocolCallerSession) Authenticator() (common.Address, error) {
	return _CowProtocol.Contract.Authenticator(&_CowProtocol.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_CowProtocol *CowProtocolCaller) DomainSeparator(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CowProtocol.contract.Call(opts, &out, "domainSeparator")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_CowProtocol *CowProtocolSession) DomainSeparator() ([32]byte, error) {
	return _CowProtocol.Contract.DomainSeparator(&_CowProtocol.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_CowProtocol *CowProtocolCallerSession) DomainSeparator() ([32]byte, error) {
	return _CowProtocol.Contract.DomainSeparator(&_CowProtocol.CallOpts)
}

// FilledAmount is a free data retrieval call binding the contract method 0x2479fb6e.
//
// Solidity: function filledAmount(bytes ) view returns(uint256)
func (_CowProtocol *CowProtocolCaller) FilledAmount(opts *bind.CallOpts, arg0 []byte) (*big.Int, error) {
	var out []interface{}
	err := _CowProtocol.contract.Call(opts, &out, "filledAmount", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FilledAmount is a free data retrieval call binding the contract method 0x2479fb6e.
//
// Solidity: function filledAmount(bytes ) view returns(uint256)
func (_CowProtocol *CowProtocolSession) FilledAmount(arg0 []byte) (*big.Int, error) {
	return _CowProtocol.Contract.FilledAmount(&_CowProtocol.CallOpts, arg0)
}

// FilledAmount is a free data retrieval call binding the contract method 0x2479fb6e.
//
// Solidity: function filledAmount(bytes ) view returns(uint256)
func (_CowProtocol *CowProtocolCallerSession) FilledAmount(arg0 []byte) (*big.Int, error) {
	return _CowProtocol.Contract.FilledAmount(&_CowProtocol.CallOpts, arg0)
}

// GetStorageAt is a free data retrieval call binding the contract method 0x5624b25b.
//
// Solidity: function getStorageAt(uint256 offset, uint256 length) view returns(bytes)
func (_CowProtocol *CowProtocolCaller) GetStorageAt(opts *bind.CallOpts, offset *big.Int, length *big.Int) ([]byte, error) {
	var out []interface{}
	err := _CowProtocol.contract.Call(opts, &out, "getStorageAt", offset, length)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetStorageAt is a free data retrieval call binding the contract method 0x5624b25b.
//
// Solidity: function getStorageAt(uint256 offset, uint256 length) view returns(bytes)
func (_CowProtocol *CowProtocolSession) GetStorageAt(offset *big.Int, length *big.Int) ([]byte, error) {
	return _CowProtocol.Contract.GetStorageAt(&_CowProtocol.CallOpts, offset, length)
}

// GetStorageAt is a free data retrieval call binding the contract method 0x5624b25b.
//
// Solidity: function getStorageAt(uint256 offset, uint256 length) view returns(bytes)
func (_CowProtocol *CowProtocolCallerSession) GetStorageAt(offset *big.Int, length *big.Int) ([]byte, error) {
	return _CowProtocol.Contract.GetStorageAt(&_CowProtocol.CallOpts, offset, length)
}

// PreSignature is a free data retrieval call binding the contract method 0xd08d33d1.
//
// Solidity: function preSignature(bytes ) view returns(uint256)
func (_CowProtocol *CowProtocolCaller) PreSignature(opts *bind.CallOpts, arg0 []byte) (*big.Int, error) {
	var out []interface{}
	err := _CowProtocol.contract.Call(opts, &out, "preSignature", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreSignature is a free data retrieval call binding the contract method 0xd08d33d1.
//
// Solidity: function preSignature(bytes ) view returns(uint256)
func (_CowProtocol *CowProtocolSession) PreSignature(arg0 []byte) (*big.Int, error) {
	return _CowProtocol.Contract.PreSignature(&_CowProtocol.CallOpts, arg0)
}

// PreSignature is a free data retrieval call binding the contract method 0xd08d33d1.
//
// Solidity: function preSignature(bytes ) view returns(uint256)
func (_CowProtocol *CowProtocolCallerSession) PreSignature(arg0 []byte) (*big.Int, error) {
	return _CowProtocol.Contract.PreSignature(&_CowProtocol.CallOpts, arg0)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_CowProtocol *CowProtocolCaller) Vault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CowProtocol.contract.Call(opts, &out, "vault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_CowProtocol *CowProtocolSession) Vault() (common.Address, error) {
	return _CowProtocol.Contract.Vault(&_CowProtocol.CallOpts)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_CowProtocol *CowProtocolCallerSession) Vault() (common.Address, error) {
	return _CowProtocol.Contract.Vault(&_CowProtocol.CallOpts)
}

// VaultRelayer is a free data retrieval call binding the contract method 0x9b552cc2.
//
// Solidity: function vaultRelayer() view returns(address)
func (_CowProtocol *CowProtocolCaller) VaultRelayer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CowProtocol.contract.Call(opts, &out, "vaultRelayer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VaultRelayer is a free data retrieval call binding the contract method 0x9b552cc2.
//
// Solidity: function vaultRelayer() view returns(address)
func (_CowProtocol *CowProtocolSession) VaultRelayer() (common.Address, error) {
	return _CowProtocol.Contract.VaultRelayer(&_CowProtocol.CallOpts)
}

// VaultRelayer is a free data retrieval call binding the contract method 0x9b552cc2.
//
// Solidity: function vaultRelayer() view returns(address)
func (_CowProtocol *CowProtocolCallerSession) VaultRelayer() (common.Address, error) {
	return _CowProtocol.Contract.VaultRelayer(&_CowProtocol.CallOpts)
}

// FreeFilledAmountStorage is a paid mutator transaction binding the contract method 0xed9f35ce.
//
// Solidity: function freeFilledAmountStorage(bytes[] orderUids) returns()
func (_CowProtocol *CowProtocolTransactor) FreeFilledAmountStorage(opts *bind.TransactOpts, orderUids [][]byte) (*types.Transaction, error) {
	return _CowProtocol.contract.Transact(opts, "freeFilledAmountStorage", orderUids)
}

// FreeFilledAmountStorage is a paid mutator transaction binding the contract method 0xed9f35ce.
//
// Solidity: function freeFilledAmountStorage(bytes[] orderUids) returns()
func (_CowProtocol *CowProtocolSession) FreeFilledAmountStorage(orderUids [][]byte) (*types.Transaction, error) {
	return _CowProtocol.Contract.FreeFilledAmountStorage(&_CowProtocol.TransactOpts, orderUids)
}

// FreeFilledAmountStorage is a paid mutator transaction binding the contract method 0xed9f35ce.
//
// Solidity: function freeFilledAmountStorage(bytes[] orderUids) returns()
func (_CowProtocol *CowProtocolTransactorSession) FreeFilledAmountStorage(orderUids [][]byte) (*types.Transaction, error) {
	return _CowProtocol.Contract.FreeFilledAmountStorage(&_CowProtocol.TransactOpts, orderUids)
}

// FreePreSignatureStorage is a paid mutator transaction binding the contract method 0xa2a7d51b.
//
// Solidity: function freePreSignatureStorage(bytes[] orderUids) returns()
func (_CowProtocol *CowProtocolTransactor) FreePreSignatureStorage(opts *bind.TransactOpts, orderUids [][]byte) (*types.Transaction, error) {
	return _CowProtocol.contract.Transact(opts, "freePreSignatureStorage", orderUids)
}

// FreePreSignatureStorage is a paid mutator transaction binding the contract method 0xa2a7d51b.
//
// Solidity: function freePreSignatureStorage(bytes[] orderUids) returns()
func (_CowProtocol *CowProtocolSession) FreePreSignatureStorage(orderUids [][]byte) (*types.Transaction, error) {
	return _CowProtocol.Contract.FreePreSignatureStorage(&_CowProtocol.TransactOpts, orderUids)
}

// FreePreSignatureStorage is a paid mutator transaction binding the contract method 0xa2a7d51b.
//
// Solidity: function freePreSignatureStorage(bytes[] orderUids) returns()
func (_CowProtocol *CowProtocolTransactorSession) FreePreSignatureStorage(orderUids [][]byte) (*types.Transaction, error) {
	return _CowProtocol.Contract.FreePreSignatureStorage(&_CowProtocol.TransactOpts, orderUids)
}

// InvalidateOrder is a paid mutator transaction binding the contract method 0x15337bc0.
//
// Solidity: function invalidateOrder(bytes orderUid) returns()
func (_CowProtocol *CowProtocolTransactor) InvalidateOrder(opts *bind.TransactOpts, orderUid []byte) (*types.Transaction, error) {
	return _CowProtocol.contract.Transact(opts, "invalidateOrder", orderUid)
}

// InvalidateOrder is a paid mutator transaction binding the contract method 0x15337bc0.
//
// Solidity: function invalidateOrder(bytes orderUid) returns()
func (_CowProtocol *CowProtocolSession) InvalidateOrder(orderUid []byte) (*types.Transaction, error) {
	return _CowProtocol.Contract.InvalidateOrder(&_CowProtocol.TransactOpts, orderUid)
}

// InvalidateOrder is a paid mutator transaction binding the contract method 0x15337bc0.
//
// Solidity: function invalidateOrder(bytes orderUid) returns()
func (_CowProtocol *CowProtocolTransactorSession) InvalidateOrder(orderUid []byte) (*types.Transaction, error) {
	return _CowProtocol.Contract.InvalidateOrder(&_CowProtocol.TransactOpts, orderUid)
}

// SetPreSignature is a paid mutator transaction binding the contract method 0xec6cb13f.
//
// Solidity: function setPreSignature(bytes orderUid, bool signed) returns()
func (_CowProtocol *CowProtocolTransactor) SetPreSignature(opts *bind.TransactOpts, orderUid []byte, signed bool) (*types.Transaction, error) {
	return _CowProtocol.contract.Transact(opts, "setPreSignature", orderUid, signed)
}

// SetPreSignature is a paid mutator transaction binding the contract method 0xec6cb13f.
//
// Solidity: function setPreSignature(bytes orderUid, bool signed) returns()
func (_CowProtocol *CowProtocolSession) SetPreSignature(orderUid []byte, signed bool) (*types.Transaction, error) {
	return _CowProtocol.Contract.SetPreSignature(&_CowProtocol.TransactOpts, orderUid, signed)
}

// SetPreSignature is a paid mutator transaction binding the contract method 0xec6cb13f.
//
// Solidity: function setPreSignature(bytes orderUid, bool signed) returns()
func (_CowProtocol *CowProtocolTransactorSession) SetPreSignature(orderUid []byte, signed bool) (*types.Transaction, error) {
	return _CowProtocol.Contract.SetPreSignature(&_CowProtocol.TransactOpts, orderUid, signed)
}

// Settle is a paid mutator transaction binding the contract method 0x13d79a0b.
//
// Solidity: function settle(address[] tokens, uint256[] clearingPrices, (uint256,uint256,address,uint256,uint256,uint32,bytes32,uint256,uint256,uint256,bytes)[] trades, (address,uint256,bytes)[][3] interactions) returns()
func (_CowProtocol *CowProtocolTransactor) Settle(opts *bind.TransactOpts, tokens []common.Address, clearingPrices []*big.Int, trades []GPv2TradeData, interactions [3][]GPv2InteractionData) (*types.Transaction, error) {
	return _CowProtocol.contract.Transact(opts, "settle", tokens, clearingPrices, trades, interactions)
}

// Settle is a paid mutator transaction binding the contract method 0x13d79a0b.
//
// Solidity: function settle(address[] tokens, uint256[] clearingPrices, (uint256,uint256,address,uint256,uint256,uint32,bytes32,uint256,uint256,uint256,bytes)[] trades, (address,uint256,bytes)[][3] interactions) returns()
func (_CowProtocol *CowProtocolSession) Settle(tokens []common.Address, clearingPrices []*big.Int, trades []GPv2TradeData, interactions [3][]GPv2InteractionData) (*types.Transaction, error) {
	return _CowProtocol.Contract.Settle(&_CowProtocol.TransactOpts, tokens, clearingPrices, trades, interactions)
}

// Settle is a paid mutator transaction binding the contract method 0x13d79a0b.
//
// Solidity: function settle(address[] tokens, uint256[] clearingPrices, (uint256,uint256,address,uint256,uint256,uint32,bytes32,uint256,uint256,uint256,bytes)[] trades, (address,uint256,bytes)[][3] interactions) returns()
func (_CowProtocol *CowProtocolTransactorSession) Settle(tokens []common.Address, clearingPrices []*big.Int, trades []GPv2TradeData, interactions [3][]GPv2InteractionData) (*types.Transaction, error) {
	return _CowProtocol.Contract.Settle(&_CowProtocol.TransactOpts, tokens, clearingPrices, trades, interactions)
}

// SimulateDelegatecall is a paid mutator transaction binding the contract method 0xf84436bd.
//
// Solidity: function simulateDelegatecall(address targetContract, bytes calldataPayload) returns(bytes response)
func (_CowProtocol *CowProtocolTransactor) SimulateDelegatecall(opts *bind.TransactOpts, targetContract common.Address, calldataPayload []byte) (*types.Transaction, error) {
	return _CowProtocol.contract.Transact(opts, "simulateDelegatecall", targetContract, calldataPayload)
}

// SimulateDelegatecall is a paid mutator transaction binding the contract method 0xf84436bd.
//
// Solidity: function simulateDelegatecall(address targetContract, bytes calldataPayload) returns(bytes response)
func (_CowProtocol *CowProtocolSession) SimulateDelegatecall(targetContract common.Address, calldataPayload []byte) (*types.Transaction, error) {
	return _CowProtocol.Contract.SimulateDelegatecall(&_CowProtocol.TransactOpts, targetContract, calldataPayload)
}

// SimulateDelegatecall is a paid mutator transaction binding the contract method 0xf84436bd.
//
// Solidity: function simulateDelegatecall(address targetContract, bytes calldataPayload) returns(bytes response)
func (_CowProtocol *CowProtocolTransactorSession) SimulateDelegatecall(targetContract common.Address, calldataPayload []byte) (*types.Transaction, error) {
	return _CowProtocol.Contract.SimulateDelegatecall(&_CowProtocol.TransactOpts, targetContract, calldataPayload)
}

// SimulateDelegatecallInternal is a paid mutator transaction binding the contract method 0x43218e19.
//
// Solidity: function simulateDelegatecallInternal(address targetContract, bytes calldataPayload) returns(bytes response)
func (_CowProtocol *CowProtocolTransactor) SimulateDelegatecallInternal(opts *bind.TransactOpts, targetContract common.Address, calldataPayload []byte) (*types.Transaction, error) {
	return _CowProtocol.contract.Transact(opts, "simulateDelegatecallInternal", targetContract, calldataPayload)
}

// SimulateDelegatecallInternal is a paid mutator transaction binding the contract method 0x43218e19.
//
// Solidity: function simulateDelegatecallInternal(address targetContract, bytes calldataPayload) returns(bytes response)
func (_CowProtocol *CowProtocolSession) SimulateDelegatecallInternal(targetContract common.Address, calldataPayload []byte) (*types.Transaction, error) {
	return _CowProtocol.Contract.SimulateDelegatecallInternal(&_CowProtocol.TransactOpts, targetContract, calldataPayload)
}

// SimulateDelegatecallInternal is a paid mutator transaction binding the contract method 0x43218e19.
//
// Solidity: function simulateDelegatecallInternal(address targetContract, bytes calldataPayload) returns(bytes response)
func (_CowProtocol *CowProtocolTransactorSession) SimulateDelegatecallInternal(targetContract common.Address, calldataPayload []byte) (*types.Transaction, error) {
	return _CowProtocol.Contract.SimulateDelegatecallInternal(&_CowProtocol.TransactOpts, targetContract, calldataPayload)
}

// Swap is a paid mutator transaction binding the contract method 0x845a101f.
//
// Solidity: function swap((bytes32,uint256,uint256,uint256,bytes)[] swaps, address[] tokens, (uint256,uint256,address,uint256,uint256,uint32,bytes32,uint256,uint256,uint256,bytes) trade) returns()
func (_CowProtocol *CowProtocolTransactor) Swap(opts *bind.TransactOpts, swaps []IVaultBatchSwapStep, tokens []common.Address, trade GPv2TradeData) (*types.Transaction, error) {
	return _CowProtocol.contract.Transact(opts, "swap", swaps, tokens, trade)
}

// Swap is a paid mutator transaction binding the contract method 0x845a101f.
//
// Solidity: function swap((bytes32,uint256,uint256,uint256,bytes)[] swaps, address[] tokens, (uint256,uint256,address,uint256,uint256,uint32,bytes32,uint256,uint256,uint256,bytes) trade) returns()
func (_CowProtocol *CowProtocolSession) Swap(swaps []IVaultBatchSwapStep, tokens []common.Address, trade GPv2TradeData) (*types.Transaction, error) {
	return _CowProtocol.Contract.Swap(&_CowProtocol.TransactOpts, swaps, tokens, trade)
}

// Swap is a paid mutator transaction binding the contract method 0x845a101f.
//
// Solidity: function swap((bytes32,uint256,uint256,uint256,bytes)[] swaps, address[] tokens, (uint256,uint256,address,uint256,uint256,uint32,bytes32,uint256,uint256,uint256,bytes) trade) returns()
func (_CowProtocol *CowProtocolTransactorSession) Swap(swaps []IVaultBatchSwapStep, tokens []common.Address, trade GPv2TradeData) (*types.Transaction, error) {
	return _CowProtocol.Contract.Swap(&_CowProtocol.TransactOpts, swaps, tokens, trade)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_CowProtocol *CowProtocolTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CowProtocol.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_CowProtocol *CowProtocolSession) Receive() (*types.Transaction, error) {
	return _CowProtocol.Contract.Receive(&_CowProtocol.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_CowProtocol *CowProtocolTransactorSession) Receive() (*types.Transaction, error) {
	return _CowProtocol.Contract.Receive(&_CowProtocol.TransactOpts)
}

// CowProtocolInteractionIterator is returned from FilterInteraction and is used to iterate over the raw logs and unpacked data for Interaction events raised by the CowProtocol contract.
type CowProtocolInteractionIterator struct {
	Event *CowProtocolInteraction // Event containing the contract specifics and raw log

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
func (it *CowProtocolInteractionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CowProtocolInteraction)
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
		it.Event = new(CowProtocolInteraction)
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
func (it *CowProtocolInteractionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CowProtocolInteractionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CowProtocolInteraction represents a Interaction event raised by the CowProtocol contract.
type CowProtocolInteraction struct {
	Target   common.Address
	Value    *big.Int
	Selector [4]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterInteraction is a free log retrieval operation binding the contract event 0xed99827efb37016f2275f98c4bcf71c7551c75d59e9b450f79fa32e60be672c2.
//
// Solidity: event Interaction(address indexed target, uint256 value, bytes4 selector)
func (_CowProtocol *CowProtocolFilterer) FilterInteraction(opts *bind.FilterOpts, target []common.Address) (*CowProtocolInteractionIterator, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _CowProtocol.contract.FilterLogs(opts, "Interaction", targetRule)
	if err != nil {
		return nil, err
	}
	return &CowProtocolInteractionIterator{contract: _CowProtocol.contract, event: "Interaction", logs: logs, sub: sub}, nil
}

// WatchInteraction is a free log subscription operation binding the contract event 0xed99827efb37016f2275f98c4bcf71c7551c75d59e9b450f79fa32e60be672c2.
//
// Solidity: event Interaction(address indexed target, uint256 value, bytes4 selector)
func (_CowProtocol *CowProtocolFilterer) WatchInteraction(opts *bind.WatchOpts, sink chan<- *CowProtocolInteraction, target []common.Address) (event.Subscription, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _CowProtocol.contract.WatchLogs(opts, "Interaction", targetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CowProtocolInteraction)
				if err := _CowProtocol.contract.UnpackLog(event, "Interaction", log); err != nil {
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

// ParseInteraction is a log parse operation binding the contract event 0xed99827efb37016f2275f98c4bcf71c7551c75d59e9b450f79fa32e60be672c2.
//
// Solidity: event Interaction(address indexed target, uint256 value, bytes4 selector)
func (_CowProtocol *CowProtocolFilterer) ParseInteraction(log types.Log) (*CowProtocolInteraction, error) {
	event := new(CowProtocolInteraction)
	if err := _CowProtocol.contract.UnpackLog(event, "Interaction", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CowProtocolOrderInvalidatedIterator is returned from FilterOrderInvalidated and is used to iterate over the raw logs and unpacked data for OrderInvalidated events raised by the CowProtocol contract.
type CowProtocolOrderInvalidatedIterator struct {
	Event *CowProtocolOrderInvalidated // Event containing the contract specifics and raw log

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
func (it *CowProtocolOrderInvalidatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CowProtocolOrderInvalidated)
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
		it.Event = new(CowProtocolOrderInvalidated)
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
func (it *CowProtocolOrderInvalidatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CowProtocolOrderInvalidatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CowProtocolOrderInvalidated represents a OrderInvalidated event raised by the CowProtocol contract.
type CowProtocolOrderInvalidated struct {
	Owner    common.Address
	OrderUid []byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOrderInvalidated is a free log retrieval operation binding the contract event 0x875b6cb035bbd4ac6500fabc6d1e4ca5bdc58a3e2b424ccb5c24cdbebeb009a9.
//
// Solidity: event OrderInvalidated(address indexed owner, bytes orderUid)
func (_CowProtocol *CowProtocolFilterer) FilterOrderInvalidated(opts *bind.FilterOpts, owner []common.Address) (*CowProtocolOrderInvalidatedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _CowProtocol.contract.FilterLogs(opts, "OrderInvalidated", ownerRule)
	if err != nil {
		return nil, err
	}
	return &CowProtocolOrderInvalidatedIterator{contract: _CowProtocol.contract, event: "OrderInvalidated", logs: logs, sub: sub}, nil
}

// WatchOrderInvalidated is a free log subscription operation binding the contract event 0x875b6cb035bbd4ac6500fabc6d1e4ca5bdc58a3e2b424ccb5c24cdbebeb009a9.
//
// Solidity: event OrderInvalidated(address indexed owner, bytes orderUid)
func (_CowProtocol *CowProtocolFilterer) WatchOrderInvalidated(opts *bind.WatchOpts, sink chan<- *CowProtocolOrderInvalidated, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _CowProtocol.contract.WatchLogs(opts, "OrderInvalidated", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CowProtocolOrderInvalidated)
				if err := _CowProtocol.contract.UnpackLog(event, "OrderInvalidated", log); err != nil {
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

// ParseOrderInvalidated is a log parse operation binding the contract event 0x875b6cb035bbd4ac6500fabc6d1e4ca5bdc58a3e2b424ccb5c24cdbebeb009a9.
//
// Solidity: event OrderInvalidated(address indexed owner, bytes orderUid)
func (_CowProtocol *CowProtocolFilterer) ParseOrderInvalidated(log types.Log) (*CowProtocolOrderInvalidated, error) {
	event := new(CowProtocolOrderInvalidated)
	if err := _CowProtocol.contract.UnpackLog(event, "OrderInvalidated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CowProtocolPreSignatureIterator is returned from FilterPreSignature and is used to iterate over the raw logs and unpacked data for PreSignature events raised by the CowProtocol contract.
type CowProtocolPreSignatureIterator struct {
	Event *CowProtocolPreSignature // Event containing the contract specifics and raw log

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
func (it *CowProtocolPreSignatureIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CowProtocolPreSignature)
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
		it.Event = new(CowProtocolPreSignature)
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
func (it *CowProtocolPreSignatureIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CowProtocolPreSignatureIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CowProtocolPreSignature represents a PreSignature event raised by the CowProtocol contract.
type CowProtocolPreSignature struct {
	Owner    common.Address
	OrderUid []byte
	Signed   bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPreSignature is a free log retrieval operation binding the contract event 0x01bf7c8b0ca55deecbea89d7e58295b7ffbf685fd0d96801034ba8c6ffe1c68d.
//
// Solidity: event PreSignature(address indexed owner, bytes orderUid, bool signed)
func (_CowProtocol *CowProtocolFilterer) FilterPreSignature(opts *bind.FilterOpts, owner []common.Address) (*CowProtocolPreSignatureIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _CowProtocol.contract.FilterLogs(opts, "PreSignature", ownerRule)
	if err != nil {
		return nil, err
	}
	return &CowProtocolPreSignatureIterator{contract: _CowProtocol.contract, event: "PreSignature", logs: logs, sub: sub}, nil
}

// WatchPreSignature is a free log subscription operation binding the contract event 0x01bf7c8b0ca55deecbea89d7e58295b7ffbf685fd0d96801034ba8c6ffe1c68d.
//
// Solidity: event PreSignature(address indexed owner, bytes orderUid, bool signed)
func (_CowProtocol *CowProtocolFilterer) WatchPreSignature(opts *bind.WatchOpts, sink chan<- *CowProtocolPreSignature, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _CowProtocol.contract.WatchLogs(opts, "PreSignature", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CowProtocolPreSignature)
				if err := _CowProtocol.contract.UnpackLog(event, "PreSignature", log); err != nil {
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

// ParsePreSignature is a log parse operation binding the contract event 0x01bf7c8b0ca55deecbea89d7e58295b7ffbf685fd0d96801034ba8c6ffe1c68d.
//
// Solidity: event PreSignature(address indexed owner, bytes orderUid, bool signed)
func (_CowProtocol *CowProtocolFilterer) ParsePreSignature(log types.Log) (*CowProtocolPreSignature, error) {
	event := new(CowProtocolPreSignature)
	if err := _CowProtocol.contract.UnpackLog(event, "PreSignature", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CowProtocolSettlementIterator is returned from FilterSettlement and is used to iterate over the raw logs and unpacked data for Settlement events raised by the CowProtocol contract.
type CowProtocolSettlementIterator struct {
	Event *CowProtocolSettlement // Event containing the contract specifics and raw log

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
func (it *CowProtocolSettlementIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CowProtocolSettlement)
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
		it.Event = new(CowProtocolSettlement)
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
func (it *CowProtocolSettlementIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CowProtocolSettlementIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CowProtocolSettlement represents a Settlement event raised by the CowProtocol contract.
type CowProtocolSettlement struct {
	Solver common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSettlement is a free log retrieval operation binding the contract event 0x40338ce1a7c49204f0099533b1e9a7ee0a3d261f84974ab7af36105b8c4e9db4.
//
// Solidity: event Settlement(address indexed solver)
func (_CowProtocol *CowProtocolFilterer) FilterSettlement(opts *bind.FilterOpts, solver []common.Address) (*CowProtocolSettlementIterator, error) {

	var solverRule []interface{}
	for _, solverItem := range solver {
		solverRule = append(solverRule, solverItem)
	}

	logs, sub, err := _CowProtocol.contract.FilterLogs(opts, "Settlement", solverRule)
	if err != nil {
		return nil, err
	}
	return &CowProtocolSettlementIterator{contract: _CowProtocol.contract, event: "Settlement", logs: logs, sub: sub}, nil
}

// WatchSettlement is a free log subscription operation binding the contract event 0x40338ce1a7c49204f0099533b1e9a7ee0a3d261f84974ab7af36105b8c4e9db4.
//
// Solidity: event Settlement(address indexed solver)
func (_CowProtocol *CowProtocolFilterer) WatchSettlement(opts *bind.WatchOpts, sink chan<- *CowProtocolSettlement, solver []common.Address) (event.Subscription, error) {

	var solverRule []interface{}
	for _, solverItem := range solver {
		solverRule = append(solverRule, solverItem)
	}

	logs, sub, err := _CowProtocol.contract.WatchLogs(opts, "Settlement", solverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CowProtocolSettlement)
				if err := _CowProtocol.contract.UnpackLog(event, "Settlement", log); err != nil {
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

// ParseSettlement is a log parse operation binding the contract event 0x40338ce1a7c49204f0099533b1e9a7ee0a3d261f84974ab7af36105b8c4e9db4.
//
// Solidity: event Settlement(address indexed solver)
func (_CowProtocol *CowProtocolFilterer) ParseSettlement(log types.Log) (*CowProtocolSettlement, error) {
	event := new(CowProtocolSettlement)
	if err := _CowProtocol.contract.UnpackLog(event, "Settlement", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CowProtocolTradeIterator is returned from FilterTrade and is used to iterate over the raw logs and unpacked data for Trade events raised by the CowProtocol contract.
type CowProtocolTradeIterator struct {
	Event *CowProtocolTrade // Event containing the contract specifics and raw log

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
func (it *CowProtocolTradeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CowProtocolTrade)
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
		it.Event = new(CowProtocolTrade)
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
func (it *CowProtocolTradeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CowProtocolTradeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CowProtocolTrade represents a Trade event raised by the CowProtocol contract.
type CowProtocolTrade struct {
	Owner      common.Address
	SellToken  common.Address
	BuyToken   common.Address
	SellAmount *big.Int
	BuyAmount  *big.Int
	FeeAmount  *big.Int
	OrderUid   []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTrade is a free log retrieval operation binding the contract event 0xa07a543ab8a018198e99ca0184c93fe9050a79400a0a723441f84de1d972cc17.
//
// Solidity: event Trade(address indexed owner, address sellToken, address buyToken, uint256 sellAmount, uint256 buyAmount, uint256 feeAmount, bytes orderUid)
func (_CowProtocol *CowProtocolFilterer) FilterTrade(opts *bind.FilterOpts, owner []common.Address) (*CowProtocolTradeIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _CowProtocol.contract.FilterLogs(opts, "Trade", ownerRule)
	if err != nil {
		return nil, err
	}
	return &CowProtocolTradeIterator{contract: _CowProtocol.contract, event: "Trade", logs: logs, sub: sub}, nil
}

// WatchTrade is a free log subscription operation binding the contract event 0xa07a543ab8a018198e99ca0184c93fe9050a79400a0a723441f84de1d972cc17.
//
// Solidity: event Trade(address indexed owner, address sellToken, address buyToken, uint256 sellAmount, uint256 buyAmount, uint256 feeAmount, bytes orderUid)
func (_CowProtocol *CowProtocolFilterer) WatchTrade(opts *bind.WatchOpts, sink chan<- *CowProtocolTrade, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _CowProtocol.contract.WatchLogs(opts, "Trade", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CowProtocolTrade)
				if err := _CowProtocol.contract.UnpackLog(event, "Trade", log); err != nil {
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

// ParseTrade is a log parse operation binding the contract event 0xa07a543ab8a018198e99ca0184c93fe9050a79400a0a723441f84de1d972cc17.
//
// Solidity: event Trade(address indexed owner, address sellToken, address buyToken, uint256 sellAmount, uint256 buyAmount, uint256 feeAmount, bytes orderUid)
func (_CowProtocol *CowProtocolFilterer) ParseTrade(log types.Log) (*CowProtocolTrade, error) {
	event := new(CowProtocolTrade)
	if err := _CowProtocol.contract.UnpackLog(event, "Trade", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
