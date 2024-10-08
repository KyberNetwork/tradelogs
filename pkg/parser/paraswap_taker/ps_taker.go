// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package pstaker

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

// IBalancerV2VaultBatchSwapStep is an auto generated low-level Go binding around an user-defined struct.
type IBalancerV2VaultBatchSwapStep struct {
	PoolId        [32]byte
	AssetInIndex  *big.Int
	AssetOutIndex *big.Int
	Amount        *big.Int
	UserData      []byte
}

// IBalancerV2VaultFundManagement is an auto generated low-level Go binding around an user-defined struct.
type IBalancerV2VaultFundManagement struct {
	Sender              common.Address
	FromInternalBalance bool
	Recipient           common.Address
	ToInternalBalance   bool
}

// UtilsDirectBalancerV2 is an auto generated low-level Go binding around an user-defined struct.
type UtilsDirectBalancerV2 struct {
	Swaps          []IBalancerV2VaultBatchSwapStep
	Assets         []common.Address
	Funds          IBalancerV2VaultFundManagement
	Limits         []*big.Int
	FromAmount     *big.Int
	ToAmount       *big.Int
	ExpectedAmount *big.Int
	Deadline       *big.Int
	FeePercent     *big.Int
	Vault          common.Address
	Partner        common.Address
	IsApproved     bool
	Beneficiary    common.Address
	Permit         []byte
	Uuid           [16]byte
}

// UtilsDirectCurveV1 is an auto generated low-level Go binding around an user-defined struct.
type UtilsDirectCurveV1 struct {
	FromToken      common.Address
	ToToken        common.Address
	Exchange       common.Address
	FromAmount     *big.Int
	ToAmount       *big.Int
	ExpectedAmount *big.Int
	FeePercent     *big.Int
	I              *big.Int
	J              *big.Int
	Partner        common.Address
	IsApproved     bool
	SwapType       uint8
	Beneficiary    common.Address
	NeedWrapNative bool
	Permit         []byte
	Uuid           [16]byte
}

// UtilsDirectCurveV2 is an auto generated low-level Go binding around an user-defined struct.
type UtilsDirectCurveV2 struct {
	FromToken      common.Address
	ToToken        common.Address
	Exchange       common.Address
	PoolAddress    common.Address
	FromAmount     *big.Int
	ToAmount       *big.Int
	ExpectedAmount *big.Int
	FeePercent     *big.Int
	I              *big.Int
	J              *big.Int
	Partner        common.Address
	IsApproved     bool
	SwapType       uint8
	Beneficiary    common.Address
	NeedWrapNative bool
	Permit         []byte
	Uuid           [16]byte
}

// UtilsDirectUniV3 is an auto generated low-level Go binding around an user-defined struct.
type UtilsDirectUniV3 struct {
	FromToken      common.Address
	ToToken        common.Address
	Exchange       common.Address
	FromAmount     *big.Int
	ToAmount       *big.Int
	ExpectedAmount *big.Int
	FeePercent     *big.Int
	Deadline       *big.Int
	Partner        common.Address
	IsApproved     bool
	Beneficiary    common.Address
	Path           []byte
	Permit         []byte
	Uuid           [16]byte
}

// PsTakerMetaData contains all meta data concerning the PsTaker contract.
var PsTakerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_weth\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_partnerSharePercent\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxFeePercent\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_paraswapReferralShare\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_paraswapSlippageShare\",\"type\":\"uint256\"},{\"internalType\":\"contractIFeeClaimer\",\"name\":\"_feeClaimer\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes16\",\"name\":\"uuid\",\"type\":\"bytes16\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"partner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feePercent\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"initiator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"srcToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"destToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"srcAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"receivedAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expectedAmount\",\"type\":\"uint256\"}],\"name\":\"BoughtV3\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes16\",\"name\":\"uuid\",\"type\":\"bytes16\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"partner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feePercent\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"initiator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumDirectSwap.DirectSwapKind\",\"name\":\"kind\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"srcToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"destToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"srcAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"receivedAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expectedAmount\",\"type\":\"uint256\"}],\"name\":\"SwappedDirect\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes16\",\"name\":\"uuid\",\"type\":\"bytes16\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"partner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feePercent\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"initiator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"srcToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"destToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"srcAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"receivedAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expectedAmount\",\"type\":\"uint256\"}],\"name\":\"SwappedV3\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ROUTER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WHITELISTED_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"poolId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"assetInIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetOutIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"userData\",\"type\":\"bytes\"}],\"internalType\":\"structIBalancerV2Vault.BatchSwapStep[]\",\"name\":\"swaps\",\"type\":\"tuple[]\"},{\"internalType\":\"address[]\",\"name\":\"assets\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"fromInternalBalance\",\"type\":\"bool\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"toInternalBalance\",\"type\":\"bool\"}],\"internalType\":\"structIBalancerV2Vault.FundManagement\",\"name\":\"funds\",\"type\":\"tuple\"},{\"internalType\":\"int256[]\",\"name\":\"limits\",\"type\":\"int256[]\"},{\"internalType\":\"uint256\",\"name\":\"fromAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feePercent\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"partner\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isApproved\",\"type\":\"bool\"},{\"internalType\":\"addresspayable\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"permit\",\"type\":\"bytes\"},{\"internalType\":\"bytes16\",\"name\":\"uuid\",\"type\":\"bytes16\"}],\"internalType\":\"structUtils.DirectBalancerV2\",\"name\":\"data\",\"type\":\"tuple\"}],\"name\":\"directBalancerV2GivenInSwap\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"poolId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"assetInIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetOutIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"userData\",\"type\":\"bytes\"}],\"internalType\":\"structIBalancerV2Vault.BatchSwapStep[]\",\"name\":\"swaps\",\"type\":\"tuple[]\"},{\"internalType\":\"address[]\",\"name\":\"assets\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"fromInternalBalance\",\"type\":\"bool\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"toInternalBalance\",\"type\":\"bool\"}],\"internalType\":\"structIBalancerV2Vault.FundManagement\",\"name\":\"funds\",\"type\":\"tuple\"},{\"internalType\":\"int256[]\",\"name\":\"limits\",\"type\":\"int256[]\"},{\"internalType\":\"uint256\",\"name\":\"fromAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feePercent\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"partner\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isApproved\",\"type\":\"bool\"},{\"internalType\":\"addresspayable\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"permit\",\"type\":\"bytes\"},{\"internalType\":\"bytes16\",\"name\":\"uuid\",\"type\":\"bytes16\"}],\"internalType\":\"structUtils.DirectBalancerV2\",\"name\":\"data\",\"type\":\"tuple\"}],\"name\":\"directBalancerV2GivenOutSwap\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"toToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"exchange\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"fromAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feePercent\",\"type\":\"uint256\"},{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"j\",\"type\":\"int128\"},{\"internalType\":\"addresspayable\",\"name\":\"partner\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isApproved\",\"type\":\"bool\"},{\"internalType\":\"enumUtils.CurveSwapType\",\"name\":\"swapType\",\"type\":\"uint8\"},{\"internalType\":\"addresspayable\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"needWrapNative\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"permit\",\"type\":\"bytes\"},{\"internalType\":\"bytes16\",\"name\":\"uuid\",\"type\":\"bytes16\"}],\"internalType\":\"structUtils.DirectCurveV1\",\"name\":\"data\",\"type\":\"tuple\"}],\"name\":\"directCurveV1Swap\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"toToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"exchange\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"poolAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"fromAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feePercent\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"j\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"partner\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isApproved\",\"type\":\"bool\"},{\"internalType\":\"enumUtils.CurveSwapType\",\"name\":\"swapType\",\"type\":\"uint8\"},{\"internalType\":\"addresspayable\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"needWrapNative\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"permit\",\"type\":\"bytes\"},{\"internalType\":\"bytes16\",\"name\":\"uuid\",\"type\":\"bytes16\"}],\"internalType\":\"structUtils.DirectCurveV2\",\"name\":\"data\",\"type\":\"tuple\"}],\"name\":\"directCurveV2Swap\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"toToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"exchange\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"fromAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feePercent\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"partner\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isApproved\",\"type\":\"bool\"},{\"internalType\":\"addresspayable\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"path\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"permit\",\"type\":\"bytes\"},{\"internalType\":\"bytes16\",\"name\":\"uuid\",\"type\":\"bytes16\"}],\"internalType\":\"structUtils.DirectUniV3\",\"name\":\"data\",\"type\":\"tuple\"}],\"name\":\"directUniV3Buy\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"toToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"exchange\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"fromAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feePercent\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"partner\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isApproved\",\"type\":\"bool\"},{\"internalType\":\"addresspayable\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"path\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"permit\",\"type\":\"bytes\"},{\"internalType\":\"bytes16\",\"name\":\"uuid\",\"type\":\"bytes16\"}],\"internalType\":\"structUtils.DirectUniV3\",\"name\":\"data\",\"type\":\"tuple\"}],\"name\":\"directUniV3Swap\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeClaimer\",\"outputs\":[{\"internalType\":\"contractIFeeClaimer\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getKey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxFeePercent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paraswapReferralShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paraswapSlippageShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"partnerSharePercent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"weth\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// PsTakerABI is the input ABI used to generate the binding from.
// Deprecated: Use PsTakerMetaData.ABI instead.
var PsTakerABI = PsTakerMetaData.ABI

// PsTaker is an auto generated Go binding around an Ethereum contract.
type PsTaker struct {
	PsTakerCaller     // Read-only binding to the contract
	PsTakerTransactor // Write-only binding to the contract
	PsTakerFilterer   // Log filterer for contract events
}

// PsTakerCaller is an auto generated read-only Go binding around an Ethereum contract.
type PsTakerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PsTakerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PsTakerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PsTakerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PsTakerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PsTakerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PsTakerSession struct {
	Contract     *PsTaker          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PsTakerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PsTakerCallerSession struct {
	Contract *PsTakerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// PsTakerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PsTakerTransactorSession struct {
	Contract     *PsTakerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// PsTakerRaw is an auto generated low-level Go binding around an Ethereum contract.
type PsTakerRaw struct {
	Contract *PsTaker // Generic contract binding to access the raw methods on
}

// PsTakerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PsTakerCallerRaw struct {
	Contract *PsTakerCaller // Generic read-only contract binding to access the raw methods on
}

// PsTakerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PsTakerTransactorRaw struct {
	Contract *PsTakerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPsTaker creates a new instance of PsTaker, bound to a specific deployed contract.
func NewPsTaker(address common.Address, backend bind.ContractBackend) (*PsTaker, error) {
	contract, err := bindPsTaker(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PsTaker{PsTakerCaller: PsTakerCaller{contract: contract}, PsTakerTransactor: PsTakerTransactor{contract: contract}, PsTakerFilterer: PsTakerFilterer{contract: contract}}, nil
}

// NewPsTakerCaller creates a new read-only instance of PsTaker, bound to a specific deployed contract.
func NewPsTakerCaller(address common.Address, caller bind.ContractCaller) (*PsTakerCaller, error) {
	contract, err := bindPsTaker(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PsTakerCaller{contract: contract}, nil
}

// NewPsTakerTransactor creates a new write-only instance of PsTaker, bound to a specific deployed contract.
func NewPsTakerTransactor(address common.Address, transactor bind.ContractTransactor) (*PsTakerTransactor, error) {
	contract, err := bindPsTaker(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PsTakerTransactor{contract: contract}, nil
}

// NewPsTakerFilterer creates a new log filterer instance of PsTaker, bound to a specific deployed contract.
func NewPsTakerFilterer(address common.Address, filterer bind.ContractFilterer) (*PsTakerFilterer, error) {
	contract, err := bindPsTaker(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PsTakerFilterer{contract: contract}, nil
}

// bindPsTaker binds a generic wrapper to an already deployed contract.
func bindPsTaker(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PsTakerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PsTaker *PsTakerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PsTaker.Contract.PsTakerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PsTaker *PsTakerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PsTaker.Contract.PsTakerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PsTaker *PsTakerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PsTaker.Contract.PsTakerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PsTaker *PsTakerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PsTaker.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PsTaker *PsTakerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PsTaker.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PsTaker *PsTakerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PsTaker.Contract.contract.Transact(opts, method, params...)
}

// ROUTERROLE is a free data retrieval call binding the contract method 0x30d643b5.
//
// Solidity: function ROUTER_ROLE() view returns(bytes32)
func (_PsTaker *PsTakerCaller) ROUTERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _PsTaker.contract.Call(opts, &out, "ROUTER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ROUTERROLE is a free data retrieval call binding the contract method 0x30d643b5.
//
// Solidity: function ROUTER_ROLE() view returns(bytes32)
func (_PsTaker *PsTakerSession) ROUTERROLE() ([32]byte, error) {
	return _PsTaker.Contract.ROUTERROLE(&_PsTaker.CallOpts)
}

// ROUTERROLE is a free data retrieval call binding the contract method 0x30d643b5.
//
// Solidity: function ROUTER_ROLE() view returns(bytes32)
func (_PsTaker *PsTakerCallerSession) ROUTERROLE() ([32]byte, error) {
	return _PsTaker.Contract.ROUTERROLE(&_PsTaker.CallOpts)
}

// WHITELISTEDROLE is a free data retrieval call binding the contract method 0x7a3226ec.
//
// Solidity: function WHITELISTED_ROLE() view returns(bytes32)
func (_PsTaker *PsTakerCaller) WHITELISTEDROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _PsTaker.contract.Call(opts, &out, "WHITELISTED_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// WHITELISTEDROLE is a free data retrieval call binding the contract method 0x7a3226ec.
//
// Solidity: function WHITELISTED_ROLE() view returns(bytes32)
func (_PsTaker *PsTakerSession) WHITELISTEDROLE() ([32]byte, error) {
	return _PsTaker.Contract.WHITELISTEDROLE(&_PsTaker.CallOpts)
}

// WHITELISTEDROLE is a free data retrieval call binding the contract method 0x7a3226ec.
//
// Solidity: function WHITELISTED_ROLE() view returns(bytes32)
func (_PsTaker *PsTakerCallerSession) WHITELISTEDROLE() ([32]byte, error) {
	return _PsTaker.Contract.WHITELISTEDROLE(&_PsTaker.CallOpts)
}

// FeeClaimer is a free data retrieval call binding the contract method 0x81cbd3ea.
//
// Solidity: function feeClaimer() view returns(address)
func (_PsTaker *PsTakerCaller) FeeClaimer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PsTaker.contract.Call(opts, &out, "feeClaimer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeClaimer is a free data retrieval call binding the contract method 0x81cbd3ea.
//
// Solidity: function feeClaimer() view returns(address)
func (_PsTaker *PsTakerSession) FeeClaimer() (common.Address, error) {
	return _PsTaker.Contract.FeeClaimer(&_PsTaker.CallOpts)
}

// FeeClaimer is a free data retrieval call binding the contract method 0x81cbd3ea.
//
// Solidity: function feeClaimer() view returns(address)
func (_PsTaker *PsTakerCallerSession) FeeClaimer() (common.Address, error) {
	return _PsTaker.Contract.FeeClaimer(&_PsTaker.CallOpts)
}

// GetKey is a free data retrieval call binding the contract method 0x82678dd6.
//
// Solidity: function getKey() pure returns(bytes32)
func (_PsTaker *PsTakerCaller) GetKey(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _PsTaker.contract.Call(opts, &out, "getKey")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetKey is a free data retrieval call binding the contract method 0x82678dd6.
//
// Solidity: function getKey() pure returns(bytes32)
func (_PsTaker *PsTakerSession) GetKey() ([32]byte, error) {
	return _PsTaker.Contract.GetKey(&_PsTaker.CallOpts)
}

// GetKey is a free data retrieval call binding the contract method 0x82678dd6.
//
// Solidity: function getKey() pure returns(bytes32)
func (_PsTaker *PsTakerCallerSession) GetKey() ([32]byte, error) {
	return _PsTaker.Contract.GetKey(&_PsTaker.CallOpts)
}

// Initialize is a free data retrieval call binding the contract method 0x439fab91.
//
// Solidity: function initialize(bytes ) pure returns()
func (_PsTaker *PsTakerCaller) Initialize(opts *bind.CallOpts, arg0 []byte) error {
	var out []interface{}
	err := _PsTaker.contract.Call(opts, &out, "initialize", arg0)

	if err != nil {
		return err
	}

	return err

}

// Initialize is a free data retrieval call binding the contract method 0x439fab91.
//
// Solidity: function initialize(bytes ) pure returns()
func (_PsTaker *PsTakerSession) Initialize(arg0 []byte) error {
	return _PsTaker.Contract.Initialize(&_PsTaker.CallOpts, arg0)
}

// Initialize is a free data retrieval call binding the contract method 0x439fab91.
//
// Solidity: function initialize(bytes ) pure returns()
func (_PsTaker *PsTakerCallerSession) Initialize(arg0 []byte) error {
	return _PsTaker.Contract.Initialize(&_PsTaker.CallOpts, arg0)
}

// MaxFeePercent is a free data retrieval call binding the contract method 0xd830a05b.
//
// Solidity: function maxFeePercent() view returns(uint256)
func (_PsTaker *PsTakerCaller) MaxFeePercent(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PsTaker.contract.Call(opts, &out, "maxFeePercent")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxFeePercent is a free data retrieval call binding the contract method 0xd830a05b.
//
// Solidity: function maxFeePercent() view returns(uint256)
func (_PsTaker *PsTakerSession) MaxFeePercent() (*big.Int, error) {
	return _PsTaker.Contract.MaxFeePercent(&_PsTaker.CallOpts)
}

// MaxFeePercent is a free data retrieval call binding the contract method 0xd830a05b.
//
// Solidity: function maxFeePercent() view returns(uint256)
func (_PsTaker *PsTakerCallerSession) MaxFeePercent() (*big.Int, error) {
	return _PsTaker.Contract.MaxFeePercent(&_PsTaker.CallOpts)
}

// ParaswapReferralShare is a free data retrieval call binding the contract method 0xd555d4f9.
//
// Solidity: function paraswapReferralShare() view returns(uint256)
func (_PsTaker *PsTakerCaller) ParaswapReferralShare(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PsTaker.contract.Call(opts, &out, "paraswapReferralShare")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ParaswapReferralShare is a free data retrieval call binding the contract method 0xd555d4f9.
//
// Solidity: function paraswapReferralShare() view returns(uint256)
func (_PsTaker *PsTakerSession) ParaswapReferralShare() (*big.Int, error) {
	return _PsTaker.Contract.ParaswapReferralShare(&_PsTaker.CallOpts)
}

// ParaswapReferralShare is a free data retrieval call binding the contract method 0xd555d4f9.
//
// Solidity: function paraswapReferralShare() view returns(uint256)
func (_PsTaker *PsTakerCallerSession) ParaswapReferralShare() (*big.Int, error) {
	return _PsTaker.Contract.ParaswapReferralShare(&_PsTaker.CallOpts)
}

// ParaswapSlippageShare is a free data retrieval call binding the contract method 0xc25ff026.
//
// Solidity: function paraswapSlippageShare() view returns(uint256)
func (_PsTaker *PsTakerCaller) ParaswapSlippageShare(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PsTaker.contract.Call(opts, &out, "paraswapSlippageShare")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ParaswapSlippageShare is a free data retrieval call binding the contract method 0xc25ff026.
//
// Solidity: function paraswapSlippageShare() view returns(uint256)
func (_PsTaker *PsTakerSession) ParaswapSlippageShare() (*big.Int, error) {
	return _PsTaker.Contract.ParaswapSlippageShare(&_PsTaker.CallOpts)
}

// ParaswapSlippageShare is a free data retrieval call binding the contract method 0xc25ff026.
//
// Solidity: function paraswapSlippageShare() view returns(uint256)
func (_PsTaker *PsTakerCallerSession) ParaswapSlippageShare() (*big.Int, error) {
	return _PsTaker.Contract.ParaswapSlippageShare(&_PsTaker.CallOpts)
}

// PartnerSharePercent is a free data retrieval call binding the contract method 0x12070a41.
//
// Solidity: function partnerSharePercent() view returns(uint256)
func (_PsTaker *PsTakerCaller) PartnerSharePercent(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PsTaker.contract.Call(opts, &out, "partnerSharePercent")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PartnerSharePercent is a free data retrieval call binding the contract method 0x12070a41.
//
// Solidity: function partnerSharePercent() view returns(uint256)
func (_PsTaker *PsTakerSession) PartnerSharePercent() (*big.Int, error) {
	return _PsTaker.Contract.PartnerSharePercent(&_PsTaker.CallOpts)
}

// PartnerSharePercent is a free data retrieval call binding the contract method 0x12070a41.
//
// Solidity: function partnerSharePercent() view returns(uint256)
func (_PsTaker *PsTakerCallerSession) PartnerSharePercent() (*big.Int, error) {
	return _PsTaker.Contract.PartnerSharePercent(&_PsTaker.CallOpts)
}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_PsTaker *PsTakerCaller) Weth(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PsTaker.contract.Call(opts, &out, "weth")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_PsTaker *PsTakerSession) Weth() (common.Address, error) {
	return _PsTaker.Contract.Weth(&_PsTaker.CallOpts)
}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_PsTaker *PsTakerCallerSession) Weth() (common.Address, error) {
	return _PsTaker.Contract.Weth(&_PsTaker.CallOpts)
}

// DirectBalancerV2GivenInSwap is a paid mutator transaction binding the contract method 0xb22f4db8.
//
// Solidity: function directBalancerV2GivenInSwap(((bytes32,uint256,uint256,uint256,bytes)[],address[],(address,bool,address,bool),int256[],uint256,uint256,uint256,uint256,uint256,address,address,bool,address,bytes,bytes16) data) payable returns()
func (_PsTaker *PsTakerTransactor) DirectBalancerV2GivenInSwap(opts *bind.TransactOpts, data UtilsDirectBalancerV2) (*types.Transaction, error) {
	return _PsTaker.contract.Transact(opts, "directBalancerV2GivenInSwap", data)
}

// DirectBalancerV2GivenInSwap is a paid mutator transaction binding the contract method 0xb22f4db8.
//
// Solidity: function directBalancerV2GivenInSwap(((bytes32,uint256,uint256,uint256,bytes)[],address[],(address,bool,address,bool),int256[],uint256,uint256,uint256,uint256,uint256,address,address,bool,address,bytes,bytes16) data) payable returns()
func (_PsTaker *PsTakerSession) DirectBalancerV2GivenInSwap(data UtilsDirectBalancerV2) (*types.Transaction, error) {
	return _PsTaker.Contract.DirectBalancerV2GivenInSwap(&_PsTaker.TransactOpts, data)
}

// DirectBalancerV2GivenInSwap is a paid mutator transaction binding the contract method 0xb22f4db8.
//
// Solidity: function directBalancerV2GivenInSwap(((bytes32,uint256,uint256,uint256,bytes)[],address[],(address,bool,address,bool),int256[],uint256,uint256,uint256,uint256,uint256,address,address,bool,address,bytes,bytes16) data) payable returns()
func (_PsTaker *PsTakerTransactorSession) DirectBalancerV2GivenInSwap(data UtilsDirectBalancerV2) (*types.Transaction, error) {
	return _PsTaker.Contract.DirectBalancerV2GivenInSwap(&_PsTaker.TransactOpts, data)
}

// DirectBalancerV2GivenOutSwap is a paid mutator transaction binding the contract method 0x19fc5be0.
//
// Solidity: function directBalancerV2GivenOutSwap(((bytes32,uint256,uint256,uint256,bytes)[],address[],(address,bool,address,bool),int256[],uint256,uint256,uint256,uint256,uint256,address,address,bool,address,bytes,bytes16) data) payable returns()
func (_PsTaker *PsTakerTransactor) DirectBalancerV2GivenOutSwap(opts *bind.TransactOpts, data UtilsDirectBalancerV2) (*types.Transaction, error) {
	return _PsTaker.contract.Transact(opts, "directBalancerV2GivenOutSwap", data)
}

// DirectBalancerV2GivenOutSwap is a paid mutator transaction binding the contract method 0x19fc5be0.
//
// Solidity: function directBalancerV2GivenOutSwap(((bytes32,uint256,uint256,uint256,bytes)[],address[],(address,bool,address,bool),int256[],uint256,uint256,uint256,uint256,uint256,address,address,bool,address,bytes,bytes16) data) payable returns()
func (_PsTaker *PsTakerSession) DirectBalancerV2GivenOutSwap(data UtilsDirectBalancerV2) (*types.Transaction, error) {
	return _PsTaker.Contract.DirectBalancerV2GivenOutSwap(&_PsTaker.TransactOpts, data)
}

// DirectBalancerV2GivenOutSwap is a paid mutator transaction binding the contract method 0x19fc5be0.
//
// Solidity: function directBalancerV2GivenOutSwap(((bytes32,uint256,uint256,uint256,bytes)[],address[],(address,bool,address,bool),int256[],uint256,uint256,uint256,uint256,uint256,address,address,bool,address,bytes,bytes16) data) payable returns()
func (_PsTaker *PsTakerTransactorSession) DirectBalancerV2GivenOutSwap(data UtilsDirectBalancerV2) (*types.Transaction, error) {
	return _PsTaker.Contract.DirectBalancerV2GivenOutSwap(&_PsTaker.TransactOpts, data)
}

// DirectCurveV1Swap is a paid mutator transaction binding the contract method 0x3865bde6.
//
// Solidity: function directCurveV1Swap((address,address,address,uint256,uint256,uint256,uint256,int128,int128,address,bool,uint8,address,bool,bytes,bytes16) data) payable returns()
func (_PsTaker *PsTakerTransactor) DirectCurveV1Swap(opts *bind.TransactOpts, data UtilsDirectCurveV1) (*types.Transaction, error) {
	return _PsTaker.contract.Transact(opts, "directCurveV1Swap", data)
}

// DirectCurveV1Swap is a paid mutator transaction binding the contract method 0x3865bde6.
//
// Solidity: function directCurveV1Swap((address,address,address,uint256,uint256,uint256,uint256,int128,int128,address,bool,uint8,address,bool,bytes,bytes16) data) payable returns()
func (_PsTaker *PsTakerSession) DirectCurveV1Swap(data UtilsDirectCurveV1) (*types.Transaction, error) {
	return _PsTaker.Contract.DirectCurveV1Swap(&_PsTaker.TransactOpts, data)
}

// DirectCurveV1Swap is a paid mutator transaction binding the contract method 0x3865bde6.
//
// Solidity: function directCurveV1Swap((address,address,address,uint256,uint256,uint256,uint256,int128,int128,address,bool,uint8,address,bool,bytes,bytes16) data) payable returns()
func (_PsTaker *PsTakerTransactorSession) DirectCurveV1Swap(data UtilsDirectCurveV1) (*types.Transaction, error) {
	return _PsTaker.Contract.DirectCurveV1Swap(&_PsTaker.TransactOpts, data)
}

// DirectCurveV2Swap is a paid mutator transaction binding the contract method 0x58f15100.
//
// Solidity: function directCurveV2Swap((address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,address,bool,uint8,address,bool,bytes,bytes16) data) payable returns()
func (_PsTaker *PsTakerTransactor) DirectCurveV2Swap(opts *bind.TransactOpts, data UtilsDirectCurveV2) (*types.Transaction, error) {
	return _PsTaker.contract.Transact(opts, "directCurveV2Swap", data)
}

// DirectCurveV2Swap is a paid mutator transaction binding the contract method 0x58f15100.
//
// Solidity: function directCurveV2Swap((address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,address,bool,uint8,address,bool,bytes,bytes16) data) payable returns()
func (_PsTaker *PsTakerSession) DirectCurveV2Swap(data UtilsDirectCurveV2) (*types.Transaction, error) {
	return _PsTaker.Contract.DirectCurveV2Swap(&_PsTaker.TransactOpts, data)
}

// DirectCurveV2Swap is a paid mutator transaction binding the contract method 0x58f15100.
//
// Solidity: function directCurveV2Swap((address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,address,bool,uint8,address,bool,bytes,bytes16) data) payable returns()
func (_PsTaker *PsTakerTransactorSession) DirectCurveV2Swap(data UtilsDirectCurveV2) (*types.Transaction, error) {
	return _PsTaker.Contract.DirectCurveV2Swap(&_PsTaker.TransactOpts, data)
}

// DirectUniV3Buy is a paid mutator transaction binding the contract method 0x87a63926.
//
// Solidity: function directUniV3Buy((address,address,address,uint256,uint256,uint256,uint256,uint256,address,bool,address,bytes,bytes,bytes16) data) payable returns()
func (_PsTaker *PsTakerTransactor) DirectUniV3Buy(opts *bind.TransactOpts, data UtilsDirectUniV3) (*types.Transaction, error) {
	return _PsTaker.contract.Transact(opts, "directUniV3Buy", data)
}

// DirectUniV3Buy is a paid mutator transaction binding the contract method 0x87a63926.
//
// Solidity: function directUniV3Buy((address,address,address,uint256,uint256,uint256,uint256,uint256,address,bool,address,bytes,bytes,bytes16) data) payable returns()
func (_PsTaker *PsTakerSession) DirectUniV3Buy(data UtilsDirectUniV3) (*types.Transaction, error) {
	return _PsTaker.Contract.DirectUniV3Buy(&_PsTaker.TransactOpts, data)
}

// DirectUniV3Buy is a paid mutator transaction binding the contract method 0x87a63926.
//
// Solidity: function directUniV3Buy((address,address,address,uint256,uint256,uint256,uint256,uint256,address,bool,address,bytes,bytes,bytes16) data) payable returns()
func (_PsTaker *PsTakerTransactorSession) DirectUniV3Buy(data UtilsDirectUniV3) (*types.Transaction, error) {
	return _PsTaker.Contract.DirectUniV3Buy(&_PsTaker.TransactOpts, data)
}

// DirectUniV3Swap is a paid mutator transaction binding the contract method 0xa6886da9.
//
// Solidity: function directUniV3Swap((address,address,address,uint256,uint256,uint256,uint256,uint256,address,bool,address,bytes,bytes,bytes16) data) payable returns()
func (_PsTaker *PsTakerTransactor) DirectUniV3Swap(opts *bind.TransactOpts, data UtilsDirectUniV3) (*types.Transaction, error) {
	return _PsTaker.contract.Transact(opts, "directUniV3Swap", data)
}

// DirectUniV3Swap is a paid mutator transaction binding the contract method 0xa6886da9.
//
// Solidity: function directUniV3Swap((address,address,address,uint256,uint256,uint256,uint256,uint256,address,bool,address,bytes,bytes,bytes16) data) payable returns()
func (_PsTaker *PsTakerSession) DirectUniV3Swap(data UtilsDirectUniV3) (*types.Transaction, error) {
	return _PsTaker.Contract.DirectUniV3Swap(&_PsTaker.TransactOpts, data)
}

// DirectUniV3Swap is a paid mutator transaction binding the contract method 0xa6886da9.
//
// Solidity: function directUniV3Swap((address,address,address,uint256,uint256,uint256,uint256,uint256,address,bool,address,bytes,bytes,bytes16) data) payable returns()
func (_PsTaker *PsTakerTransactorSession) DirectUniV3Swap(data UtilsDirectUniV3) (*types.Transaction, error) {
	return _PsTaker.Contract.DirectUniV3Swap(&_PsTaker.TransactOpts, data)
}

// PsTakerBoughtV3Iterator is returned from FilterBoughtV3 and is used to iterate over the raw logs and unpacked data for BoughtV3 events raised by the PsTaker contract.
type PsTakerBoughtV3Iterator struct {
	Event *PsTakerBoughtV3 // Event containing the contract specifics and raw log

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
func (it *PsTakerBoughtV3Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PsTakerBoughtV3)
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
		it.Event = new(PsTakerBoughtV3)
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
func (it *PsTakerBoughtV3Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PsTakerBoughtV3Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PsTakerBoughtV3 represents a BoughtV3 event raised by the PsTaker contract.
type PsTakerBoughtV3 struct {
	Uuid           [16]byte
	Partner        common.Address
	FeePercent     *big.Int
	Initiator      common.Address
	Beneficiary    common.Address
	SrcToken       common.Address
	DestToken      common.Address
	SrcAmount      *big.Int
	ReceivedAmount *big.Int
	ExpectedAmount *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBoughtV3 is a free log retrieval operation binding the contract event 0x4cc7e95e48af62690313a0733e93308ac9a73326bc3c29f1788b1191c376d5b6.
//
// Solidity: event BoughtV3(bytes16 uuid, address partner, uint256 feePercent, address initiator, address indexed beneficiary, address indexed srcToken, address indexed destToken, uint256 srcAmount, uint256 receivedAmount, uint256 expectedAmount)
func (_PsTaker *PsTakerFilterer) FilterBoughtV3(opts *bind.FilterOpts, beneficiary []common.Address, srcToken []common.Address, destToken []common.Address) (*PsTakerBoughtV3Iterator, error) {

	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}
	var srcTokenRule []interface{}
	for _, srcTokenItem := range srcToken {
		srcTokenRule = append(srcTokenRule, srcTokenItem)
	}
	var destTokenRule []interface{}
	for _, destTokenItem := range destToken {
		destTokenRule = append(destTokenRule, destTokenItem)
	}

	logs, sub, err := _PsTaker.contract.FilterLogs(opts, "BoughtV3", beneficiaryRule, srcTokenRule, destTokenRule)
	if err != nil {
		return nil, err
	}
	return &PsTakerBoughtV3Iterator{contract: _PsTaker.contract, event: "BoughtV3", logs: logs, sub: sub}, nil
}

// WatchBoughtV3 is a free log subscription operation binding the contract event 0x4cc7e95e48af62690313a0733e93308ac9a73326bc3c29f1788b1191c376d5b6.
//
// Solidity: event BoughtV3(bytes16 uuid, address partner, uint256 feePercent, address initiator, address indexed beneficiary, address indexed srcToken, address indexed destToken, uint256 srcAmount, uint256 receivedAmount, uint256 expectedAmount)
func (_PsTaker *PsTakerFilterer) WatchBoughtV3(opts *bind.WatchOpts, sink chan<- *PsTakerBoughtV3, beneficiary []common.Address, srcToken []common.Address, destToken []common.Address) (event.Subscription, error) {

	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}
	var srcTokenRule []interface{}
	for _, srcTokenItem := range srcToken {
		srcTokenRule = append(srcTokenRule, srcTokenItem)
	}
	var destTokenRule []interface{}
	for _, destTokenItem := range destToken {
		destTokenRule = append(destTokenRule, destTokenItem)
	}

	logs, sub, err := _PsTaker.contract.WatchLogs(opts, "BoughtV3", beneficiaryRule, srcTokenRule, destTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PsTakerBoughtV3)
				if err := _PsTaker.contract.UnpackLog(event, "BoughtV3", log); err != nil {
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

// ParseBoughtV3 is a log parse operation binding the contract event 0x4cc7e95e48af62690313a0733e93308ac9a73326bc3c29f1788b1191c376d5b6.
//
// Solidity: event BoughtV3(bytes16 uuid, address partner, uint256 feePercent, address initiator, address indexed beneficiary, address indexed srcToken, address indexed destToken, uint256 srcAmount, uint256 receivedAmount, uint256 expectedAmount)
func (_PsTaker *PsTakerFilterer) ParseBoughtV3(log types.Log) (*PsTakerBoughtV3, error) {
	event := new(PsTakerBoughtV3)
	if err := _PsTaker.contract.UnpackLog(event, "BoughtV3", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PsTakerSwappedDirectIterator is returned from FilterSwappedDirect and is used to iterate over the raw logs and unpacked data for SwappedDirect events raised by the PsTaker contract.
type PsTakerSwappedDirectIterator struct {
	Event *PsTakerSwappedDirect // Event containing the contract specifics and raw log

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
func (it *PsTakerSwappedDirectIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PsTakerSwappedDirect)
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
		it.Event = new(PsTakerSwappedDirect)
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
func (it *PsTakerSwappedDirectIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PsTakerSwappedDirectIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PsTakerSwappedDirect represents a SwappedDirect event raised by the PsTaker contract.
type PsTakerSwappedDirect struct {
	Uuid           [16]byte
	Partner        common.Address
	FeePercent     *big.Int
	Initiator      common.Address
	Kind           uint8
	Beneficiary    common.Address
	SrcToken       common.Address
	DestToken      common.Address
	SrcAmount      *big.Int
	ReceivedAmount *big.Int
	ExpectedAmount *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterSwappedDirect is a free log retrieval operation binding the contract event 0xd2d73da2b5fd52cd654d8fd1b514ad57355bad741de639e3a1c3a20dd9f17347.
//
// Solidity: event SwappedDirect(bytes16 uuid, address partner, uint256 feePercent, address initiator, uint8 kind, address indexed beneficiary, address indexed srcToken, address indexed destToken, uint256 srcAmount, uint256 receivedAmount, uint256 expectedAmount)
func (_PsTaker *PsTakerFilterer) FilterSwappedDirect(opts *bind.FilterOpts, beneficiary []common.Address, srcToken []common.Address, destToken []common.Address) (*PsTakerSwappedDirectIterator, error) {

	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}
	var srcTokenRule []interface{}
	for _, srcTokenItem := range srcToken {
		srcTokenRule = append(srcTokenRule, srcTokenItem)
	}
	var destTokenRule []interface{}
	for _, destTokenItem := range destToken {
		destTokenRule = append(destTokenRule, destTokenItem)
	}

	logs, sub, err := _PsTaker.contract.FilterLogs(opts, "SwappedDirect", beneficiaryRule, srcTokenRule, destTokenRule)
	if err != nil {
		return nil, err
	}
	return &PsTakerSwappedDirectIterator{contract: _PsTaker.contract, event: "SwappedDirect", logs: logs, sub: sub}, nil
}

// WatchSwappedDirect is a free log subscription operation binding the contract event 0xd2d73da2b5fd52cd654d8fd1b514ad57355bad741de639e3a1c3a20dd9f17347.
//
// Solidity: event SwappedDirect(bytes16 uuid, address partner, uint256 feePercent, address initiator, uint8 kind, address indexed beneficiary, address indexed srcToken, address indexed destToken, uint256 srcAmount, uint256 receivedAmount, uint256 expectedAmount)
func (_PsTaker *PsTakerFilterer) WatchSwappedDirect(opts *bind.WatchOpts, sink chan<- *PsTakerSwappedDirect, beneficiary []common.Address, srcToken []common.Address, destToken []common.Address) (event.Subscription, error) {

	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}
	var srcTokenRule []interface{}
	for _, srcTokenItem := range srcToken {
		srcTokenRule = append(srcTokenRule, srcTokenItem)
	}
	var destTokenRule []interface{}
	for _, destTokenItem := range destToken {
		destTokenRule = append(destTokenRule, destTokenItem)
	}

	logs, sub, err := _PsTaker.contract.WatchLogs(opts, "SwappedDirect", beneficiaryRule, srcTokenRule, destTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PsTakerSwappedDirect)
				if err := _PsTaker.contract.UnpackLog(event, "SwappedDirect", log); err != nil {
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

// ParseSwappedDirect is a log parse operation binding the contract event 0xd2d73da2b5fd52cd654d8fd1b514ad57355bad741de639e3a1c3a20dd9f17347.
//
// Solidity: event SwappedDirect(bytes16 uuid, address partner, uint256 feePercent, address initiator, uint8 kind, address indexed beneficiary, address indexed srcToken, address indexed destToken, uint256 srcAmount, uint256 receivedAmount, uint256 expectedAmount)
func (_PsTaker *PsTakerFilterer) ParseSwappedDirect(log types.Log) (*PsTakerSwappedDirect, error) {
	event := new(PsTakerSwappedDirect)
	if err := _PsTaker.contract.UnpackLog(event, "SwappedDirect", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PsTakerSwappedV3Iterator is returned from FilterSwappedV3 and is used to iterate over the raw logs and unpacked data for SwappedV3 events raised by the PsTaker contract.
type PsTakerSwappedV3Iterator struct {
	Event *PsTakerSwappedV3 // Event containing the contract specifics and raw log

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
func (it *PsTakerSwappedV3Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PsTakerSwappedV3)
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
		it.Event = new(PsTakerSwappedV3)
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
func (it *PsTakerSwappedV3Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PsTakerSwappedV3Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PsTakerSwappedV3 represents a SwappedV3 event raised by the PsTaker contract.
type PsTakerSwappedV3 struct {
	Uuid           [16]byte
	Partner        common.Address
	FeePercent     *big.Int
	Initiator      common.Address
	Beneficiary    common.Address
	SrcToken       common.Address
	DestToken      common.Address
	SrcAmount      *big.Int
	ReceivedAmount *big.Int
	ExpectedAmount *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterSwappedV3 is a free log retrieval operation binding the contract event 0xe00361d207b252a464323eb23d45d42583e391f2031acdd2e9fa36efddd43cb0.
//
// Solidity: event SwappedV3(bytes16 uuid, address partner, uint256 feePercent, address initiator, address indexed beneficiary, address indexed srcToken, address indexed destToken, uint256 srcAmount, uint256 receivedAmount, uint256 expectedAmount)
func (_PsTaker *PsTakerFilterer) FilterSwappedV3(opts *bind.FilterOpts, beneficiary []common.Address, srcToken []common.Address, destToken []common.Address) (*PsTakerSwappedV3Iterator, error) {

	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}
	var srcTokenRule []interface{}
	for _, srcTokenItem := range srcToken {
		srcTokenRule = append(srcTokenRule, srcTokenItem)
	}
	var destTokenRule []interface{}
	for _, destTokenItem := range destToken {
		destTokenRule = append(destTokenRule, destTokenItem)
	}

	logs, sub, err := _PsTaker.contract.FilterLogs(opts, "SwappedV3", beneficiaryRule, srcTokenRule, destTokenRule)
	if err != nil {
		return nil, err
	}
	return &PsTakerSwappedV3Iterator{contract: _PsTaker.contract, event: "SwappedV3", logs: logs, sub: sub}, nil
}

// WatchSwappedV3 is a free log subscription operation binding the contract event 0xe00361d207b252a464323eb23d45d42583e391f2031acdd2e9fa36efddd43cb0.
//
// Solidity: event SwappedV3(bytes16 uuid, address partner, uint256 feePercent, address initiator, address indexed beneficiary, address indexed srcToken, address indexed destToken, uint256 srcAmount, uint256 receivedAmount, uint256 expectedAmount)
func (_PsTaker *PsTakerFilterer) WatchSwappedV3(opts *bind.WatchOpts, sink chan<- *PsTakerSwappedV3, beneficiary []common.Address, srcToken []common.Address, destToken []common.Address) (event.Subscription, error) {

	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}
	var srcTokenRule []interface{}
	for _, srcTokenItem := range srcToken {
		srcTokenRule = append(srcTokenRule, srcTokenItem)
	}
	var destTokenRule []interface{}
	for _, destTokenItem := range destToken {
		destTokenRule = append(destTokenRule, destTokenItem)
	}

	logs, sub, err := _PsTaker.contract.WatchLogs(opts, "SwappedV3", beneficiaryRule, srcTokenRule, destTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PsTakerSwappedV3)
				if err := _PsTaker.contract.UnpackLog(event, "SwappedV3", log); err != nil {
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

// ParseSwappedV3 is a log parse operation binding the contract event 0xe00361d207b252a464323eb23d45d42583e391f2031acdd2e9fa36efddd43cb0.
//
// Solidity: event SwappedV3(bytes16 uuid, address partner, uint256 feePercent, address initiator, address indexed beneficiary, address indexed srcToken, address indexed destToken, uint256 srcAmount, uint256 receivedAmount, uint256 expectedAmount)
func (_PsTaker *PsTakerFilterer) ParseSwappedV3(log types.Log) (*PsTakerSwappedV3, error) {
	event := new(PsTakerSwappedV3)
	if err := _PsTaker.contract.UnpackLog(event, "SwappedV3", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
