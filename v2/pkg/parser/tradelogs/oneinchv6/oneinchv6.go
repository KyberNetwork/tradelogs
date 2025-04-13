// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package oneinchv6

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

// GenericRouterSwapDescription is an auto generated low-level Go binding around an user-defined struct.
type GenericRouterSwapDescription struct {
	SrcToken        common.Address
	DstToken        common.Address
	SrcReceiver     common.Address
	DstReceiver     common.Address
	Amount          *big.Int
	MinReturnAmount *big.Int
	Flags           *big.Int
}

// IOrderMixinOrder is an auto generated low-level Go binding around an user-defined struct.
type IOrderMixinOrder struct {
	Salt         *big.Int
	Maker        *big.Int
	Receiver     *big.Int
	MakerAsset   *big.Int
	TakerAsset   *big.Int
	MakingAmount *big.Int
	TakingAmount *big.Int
	MakerTraits  *big.Int
}

// Oneinchv6MetaData contains all meta data concerning the Oneinchv6 contract.
var Oneinchv6MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIWETH\",\"name\":\"weth\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AdvanceEpochFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ArbitraryStaticCallFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadCurveSwapSelector\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadPool\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BitInvalidatedOrder\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ETHTransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ETHTransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EpochManagerAndBitInvalidatorsAreIncompatible\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EthDepositRejected\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMsgValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMsgValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPermit2Transfer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidShortString\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidatedOrder\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MakingAmountTooLow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MismatchArraysLengths\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OrderExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OrderIsNotSuitableForMassInvalidation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PartialFillNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Permit2TransferAmountTooHigh\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PredicateIsNotTrue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PrivateOrder\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyDetected\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RemainingInvalidatedOrder\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReservesCallFailed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"result\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minReturn\",\"type\":\"uint256\"}],\"name\":\"ReturnAmountIsNotEnough\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SafeTransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SafeTransferFromFailed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"res\",\"type\":\"bytes\"}],\"name\":\"SimulationResults\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"str\",\"type\":\"string\"}],\"name\":\"StringTooLong\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SwapWithZeroAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TakingAmountExceeded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TakingAmountTooHigh\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferFromMakerToTakerFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferFromTakerToMakerFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WrongSeriesNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroMinReturn\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"slotIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"slotValue\",\"type\":\"uint256\"}],\"name\":\"BitInvalidatorUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"series\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newEpoch\",\"type\":\"uint256\"}],\"name\":\"EpochIncreased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"name\":\"OrderCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"remainingAmount\",\"type\":\"uint256\"}],\"name\":\"OrderFilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint96\",\"name\":\"series\",\"type\":\"uint96\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"advanceEpoch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"offsets\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"and\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"arbitraryStaticCall\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"slot\",\"type\":\"uint256\"}],\"name\":\"bitInvalidatorForOrder\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"MakerTraits\",\"name\":\"makerTraits\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"additionalMask\",\"type\":\"uint256\"}],\"name\":\"bitsInvalidateForOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"MakerTraits\",\"name\":\"makerTraits\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"name\":\"cancelOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"MakerTraits[]\",\"name\":\"makerTraits\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"orderHashes\",\"type\":\"bytes32[]\"}],\"name\":\"cancelOrders\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"predicate\",\"type\":\"bytes\"}],\"name\":\"checkPredicate\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIClipperExchange\",\"name\":\"clipperExchange\",\"type\":\"address\"},{\"internalType\":\"Address\",\"name\":\"srcToken\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"dstToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"inputAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"outputAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"goodUntil\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"vs\",\"type\":\"bytes32\"}],\"name\":\"clipperSwap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIClipperExchange\",\"name\":\"clipperExchange\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"Address\",\"name\":\"srcToken\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"dstToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"inputAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"outputAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"goodUntil\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"vs\",\"type\":\"bytes32\"}],\"name\":\"clipperSwapTo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"inCoin\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"dx\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"curveSwapCallback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"series\",\"type\":\"uint96\"}],\"name\":\"epoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"series\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerEpoch\",\"type\":\"uint256\"}],\"name\":\"epochEquals\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"eq\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minReturn\",\"type\":\"uint256\"},{\"internalType\":\"Address\",\"name\":\"dex\",\"type\":\"uint256\"}],\"name\":\"ethUnoswap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minReturn\",\"type\":\"uint256\"},{\"internalType\":\"Address\",\"name\":\"dex\",\"type\":\"uint256\"},{\"internalType\":\"Address\",\"name\":\"dex2\",\"type\":\"uint256\"}],\"name\":\"ethUnoswap2\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minReturn\",\"type\":\"uint256\"},{\"internalType\":\"Address\",\"name\":\"dex\",\"type\":\"uint256\"},{\"internalType\":\"Address\",\"name\":\"dex2\",\"type\":\"uint256\"},{\"internalType\":\"Address\",\"name\":\"dex3\",\"type\":\"uint256\"}],\"name\":\"ethUnoswap3\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"Address\",\"name\":\"to\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minReturn\",\"type\":\"uint256\"},{\"internalType\":\"Address\",\"name\":\"dex\",\"type\":\"uint256\"}],\"name\":\"ethUnoswapTo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"Address\",\"name\":\"to\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minReturn\",\"type\":\"uint256\"},{\"internalType\":\"Address\",\"name\":\"dex\",\"type\":\"uint256\"},{\"internalType\":\"Address\",\"name\":\"dex2\",\"type\":\"uint256\"}],\"name\":\"ethUnoswapTo2\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"Address\",\"name\":\"to\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minReturn\",\"type\":\"uint256\"},{\"internalType\":\"Address\",\"name\":\"dex\",\"type\":\"uint256\"},{\"internalType\":\"Address\",\"name\":\"dex2\",\"type\":\"uint256\"},{\"internalType\":\"Address\",\"name\":\"dex3\",\"type\":\"uint256\"}],\"name\":\"ethUnoswapTo3\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"Address\",\"name\":\"maker\",\"type\":\"uint256\"},{\"internalType\":\"Address\",\"name\":\"receiver\",\"type\":\"uint256\"},{\"internalType\":\"Address\",\"name\":\"makerAsset\",\"type\":\"uint256\"},{\"internalType\":\"Address\",\"name\":\"takerAsset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takingAmount\",\"type\":\"uint256\"},{\"internalType\":\"MakerTraits\",\"name\":\"makerTraits\",\"type\":\"uint256\"}],\"internalType\":\"structIOrderMixin.Order\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"TakerTraits\",\"name\":\"takerTraits\",\"type\":\"uint256\"}],\"name\":\"fillContractOrder\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"Address\",\"name\":\"maker\",\"type\":\"uint256\"},{\"internalType\":\"Address\",\"name\":\"receiver\",\"type\":\"uint256\"},{\"internalType\":\"Address\",\"name\":\"makerAsset\",\"type\":\"uint256\"},{\"internalType\":\"Address\",\"name\":\"takerAsset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takingAmount\",\"type\":\"uint256\"},{\"internalType\":\"MakerTraits\",\"name\":\"makerTraits\",\"type\":\"uint256\"}],\"internalType\":\"structIOrderMixin.Order\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"TakerTraits\",\"name\":\"takerTraits\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"args\",\"type\":\"bytes\"}],\"name\":\"fillContractOrderArgs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"Address\",\"name\":\"maker\",\"type\":\"uint256\"},{\"internalType\":\"Address\",\"name\":\"receiver\",\"type\":\"uint256\"},{\"internalType\":\"Address\",\"name\":\"makerAsset\",\"type\":\"uint256\"},{\"internalType\":\"Address\",\"name\":\"takerAsset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takingAmount\",\"type\":\"uint256\"},{\"internalType\":\"MakerTraits\",\"name\":\"makerTraits\",\"type\":\"uint256\"}],\"internalType\":\"structIOrderMixin.Order\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"vs\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"TakerTraits\",\"name\":\"takerTraits\",\"type\":\"uint256\"}],\"name\":\"fillOrder\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"Address\",\"name\":\"maker\",\"type\":\"uint256\"},{\"internalType\":\"Address\",\"name\":\"receiver\",\"type\":\"uint256\"},{\"internalType\":\"Address\",\"name\":\"makerAsset\",\"type\":\"uint256\"},{\"internalType\":\"Address\",\"name\":\"takerAsset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takingAmount\",\"type\":\"uint256\"},{\"internalType\":\"MakerTraits\",\"name\":\"makerTraits\",\"type\":\"uint256\"}],\"internalType\":\"structIOrderMixin.Order\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"vs\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"TakerTraits\",\"name\":\"takerTraits\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"args\",\"type\":\"bytes\"}],\"name\":\"fillOrderArgs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"gt\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"Address\",\"name\":\"maker\",\"type\":\"uint256\"},{\"internalType\":\"Address\",\"name\":\"receiver\",\"type\":\"uint256\"},{\"internalType\":\"Address\",\"name\":\"makerAsset\",\"type\":\"uint256\"},{\"internalType\":\"Address\",\"name\":\"takerAsset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takingAmount\",\"type\":\"uint256\"},{\"internalType\":\"MakerTraits\",\"name\":\"makerTraits\",\"type\":\"uint256\"}],\"internalType\":\"structIOrderMixin.Order\",\"name\":\"order\",\"type\":\"tuple\"}],\"name\":\"hashOrder\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint96\",\"name\":\"series\",\"type\":\"uint96\"}],\"name\":\"increaseEpoch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"lt\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"not\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"offsets\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"or\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"permit\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"action\",\"type\":\"bytes\"}],\"name\":\"permitAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"name\":\"rawRemainingInvalidatorForOrder\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"name\":\"remainingInvalidatorForOrder\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"rescueFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"simulate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIAggregationExecutor\",\"name\":\"executor\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"contractIERC20\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"dstToken\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"srcReceiver\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"dstReceiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minReturnAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"flags\",\"type\":\"uint256\"}],\"internalType\":\"structGenericRouter.SwapDescription\",\"name\":\"desc\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"swap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"spentAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"amount0Delta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"amount1Delta\",\"type\":\"int256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"uniswapV3SwapCallback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"Address\",\"name\":\"token\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minReturn\",\"type\":\"uint256\"},{\"internalType\":\"Address\",\"name\":\"dex\",\"type\":\"uint256\"}],\"name\":\"unoswap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"Address\",\"name\":\"token\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minReturn\",\"type\":\"uint256\"},{\"internalType\":\"Address\",\"name\":\"dex\",\"type\":\"uint256\"},{\"internalType\":\"Address\",\"name\":\"dex2\",\"type\":\"uint256\"}],\"name\":\"unoswap2\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"Address\",\"name\":\"token\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minReturn\",\"type\":\"uint256\"},{\"internalType\":\"Address\",\"name\":\"dex\",\"type\":\"uint256\"},{\"internalType\":\"Address\",\"name\":\"dex2\",\"type\":\"uint256\"},{\"internalType\":\"Address\",\"name\":\"dex3\",\"type\":\"uint256\"}],\"name\":\"unoswap3\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"Address\",\"name\":\"to\",\"type\":\"uint256\"},{\"internalType\":\"Address\",\"name\":\"token\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minReturn\",\"type\":\"uint256\"},{\"internalType\":\"Address\",\"name\":\"dex\",\"type\":\"uint256\"}],\"name\":\"unoswapTo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"Address\",\"name\":\"to\",\"type\":\"uint256\"},{\"internalType\":\"Address\",\"name\":\"token\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minReturn\",\"type\":\"uint256\"},{\"internalType\":\"Address\",\"name\":\"dex\",\"type\":\"uint256\"},{\"internalType\":\"Address\",\"name\":\"dex2\",\"type\":\"uint256\"}],\"name\":\"unoswapTo2\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"Address\",\"name\":\"to\",\"type\":\"uint256\"},{\"internalType\":\"Address\",\"name\":\"token\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minReturn\",\"type\":\"uint256\"},{\"internalType\":\"Address\",\"name\":\"dex\",\"type\":\"uint256\"},{\"internalType\":\"Address\",\"name\":\"dex2\",\"type\":\"uint256\"},{\"internalType\":\"Address\",\"name\":\"dex3\",\"type\":\"uint256\"}],\"name\":\"unoswapTo3\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// Oneinchv6ABI is the input ABI used to generate the binding from.
// Deprecated: Use Oneinchv6MetaData.ABI instead.
var Oneinchv6ABI = Oneinchv6MetaData.ABI

// Oneinchv6 is an auto generated Go binding around an Ethereum contract.
type Oneinchv6 struct {
	Oneinchv6Caller     // Read-only binding to the contract
	Oneinchv6Transactor // Write-only binding to the contract
	Oneinchv6Filterer   // Log filterer for contract events
}

// Oneinchv6Caller is an auto generated read-only Go binding around an Ethereum contract.
type Oneinchv6Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Oneinchv6Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Oneinchv6Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Oneinchv6Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Oneinchv6Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Oneinchv6Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Oneinchv6Session struct {
	Contract     *Oneinchv6        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Oneinchv6CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Oneinchv6CallerSession struct {
	Contract *Oneinchv6Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// Oneinchv6TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Oneinchv6TransactorSession struct {
	Contract     *Oneinchv6Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// Oneinchv6Raw is an auto generated low-level Go binding around an Ethereum contract.
type Oneinchv6Raw struct {
	Contract *Oneinchv6 // Generic contract binding to access the raw methods on
}

// Oneinchv6CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Oneinchv6CallerRaw struct {
	Contract *Oneinchv6Caller // Generic read-only contract binding to access the raw methods on
}

// Oneinchv6TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Oneinchv6TransactorRaw struct {
	Contract *Oneinchv6Transactor // Generic write-only contract binding to access the raw methods on
}

// NewOneinchv6 creates a new instance of Oneinchv6, bound to a specific deployed contract.
func NewOneinchv6(address common.Address, backend bind.ContractBackend) (*Oneinchv6, error) {
	contract, err := bindOneinchv6(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Oneinchv6{Oneinchv6Caller: Oneinchv6Caller{contract: contract}, Oneinchv6Transactor: Oneinchv6Transactor{contract: contract}, Oneinchv6Filterer: Oneinchv6Filterer{contract: contract}}, nil
}

// NewOneinchv6Caller creates a new read-only instance of Oneinchv6, bound to a specific deployed contract.
func NewOneinchv6Caller(address common.Address, caller bind.ContractCaller) (*Oneinchv6Caller, error) {
	contract, err := bindOneinchv6(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Oneinchv6Caller{contract: contract}, nil
}

// NewOneinchv6Transactor creates a new write-only instance of Oneinchv6, bound to a specific deployed contract.
func NewOneinchv6Transactor(address common.Address, transactor bind.ContractTransactor) (*Oneinchv6Transactor, error) {
	contract, err := bindOneinchv6(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Oneinchv6Transactor{contract: contract}, nil
}

// NewOneinchv6Filterer creates a new log filterer instance of Oneinchv6, bound to a specific deployed contract.
func NewOneinchv6Filterer(address common.Address, filterer bind.ContractFilterer) (*Oneinchv6Filterer, error) {
	contract, err := bindOneinchv6(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Oneinchv6Filterer{contract: contract}, nil
}

// bindOneinchv6 binds a generic wrapper to an already deployed contract.
func bindOneinchv6(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Oneinchv6MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Oneinchv6 *Oneinchv6Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Oneinchv6.Contract.Oneinchv6Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Oneinchv6 *Oneinchv6Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Oneinchv6.Contract.Oneinchv6Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Oneinchv6 *Oneinchv6Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Oneinchv6.Contract.Oneinchv6Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Oneinchv6 *Oneinchv6CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Oneinchv6.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Oneinchv6 *Oneinchv6TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Oneinchv6.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Oneinchv6 *Oneinchv6TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Oneinchv6.Contract.contract.Transact(opts, method, params...)
}

// And is a free data retrieval call binding the contract method 0xbfa75143.
//
// Solidity: function and(uint256 offsets, bytes data) view returns(bool)
func (_Oneinchv6 *Oneinchv6Caller) And(opts *bind.CallOpts, offsets *big.Int, data []byte) (bool, error) {
	var out []interface{}
	err := _Oneinchv6.contract.Call(opts, &out, "and", offsets, data)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// And is a free data retrieval call binding the contract method 0xbfa75143.
//
// Solidity: function and(uint256 offsets, bytes data) view returns(bool)
func (_Oneinchv6 *Oneinchv6Session) And(offsets *big.Int, data []byte) (bool, error) {
	return _Oneinchv6.Contract.And(&_Oneinchv6.CallOpts, offsets, data)
}

// And is a free data retrieval call binding the contract method 0xbfa75143.
//
// Solidity: function and(uint256 offsets, bytes data) view returns(bool)
func (_Oneinchv6 *Oneinchv6CallerSession) And(offsets *big.Int, data []byte) (bool, error) {
	return _Oneinchv6.Contract.And(&_Oneinchv6.CallOpts, offsets, data)
}

// ArbitraryStaticCall is a free data retrieval call binding the contract method 0xbf15fcd8.
//
// Solidity: function arbitraryStaticCall(address target, bytes data) view returns(uint256)
func (_Oneinchv6 *Oneinchv6Caller) ArbitraryStaticCall(opts *bind.CallOpts, target common.Address, data []byte) (*big.Int, error) {
	var out []interface{}
	err := _Oneinchv6.contract.Call(opts, &out, "arbitraryStaticCall", target, data)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ArbitraryStaticCall is a free data retrieval call binding the contract method 0xbf15fcd8.
//
// Solidity: function arbitraryStaticCall(address target, bytes data) view returns(uint256)
func (_Oneinchv6 *Oneinchv6Session) ArbitraryStaticCall(target common.Address, data []byte) (*big.Int, error) {
	return _Oneinchv6.Contract.ArbitraryStaticCall(&_Oneinchv6.CallOpts, target, data)
}

// ArbitraryStaticCall is a free data retrieval call binding the contract method 0xbf15fcd8.
//
// Solidity: function arbitraryStaticCall(address target, bytes data) view returns(uint256)
func (_Oneinchv6 *Oneinchv6CallerSession) ArbitraryStaticCall(target common.Address, data []byte) (*big.Int, error) {
	return _Oneinchv6.Contract.ArbitraryStaticCall(&_Oneinchv6.CallOpts, target, data)
}

// BitInvalidatorForOrder is a free data retrieval call binding the contract method 0x143e86a7.
//
// Solidity: function bitInvalidatorForOrder(address maker, uint256 slot) view returns(uint256)
func (_Oneinchv6 *Oneinchv6Caller) BitInvalidatorForOrder(opts *bind.CallOpts, maker common.Address, slot *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Oneinchv6.contract.Call(opts, &out, "bitInvalidatorForOrder", maker, slot)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BitInvalidatorForOrder is a free data retrieval call binding the contract method 0x143e86a7.
//
// Solidity: function bitInvalidatorForOrder(address maker, uint256 slot) view returns(uint256)
func (_Oneinchv6 *Oneinchv6Session) BitInvalidatorForOrder(maker common.Address, slot *big.Int) (*big.Int, error) {
	return _Oneinchv6.Contract.BitInvalidatorForOrder(&_Oneinchv6.CallOpts, maker, slot)
}

// BitInvalidatorForOrder is a free data retrieval call binding the contract method 0x143e86a7.
//
// Solidity: function bitInvalidatorForOrder(address maker, uint256 slot) view returns(uint256)
func (_Oneinchv6 *Oneinchv6CallerSession) BitInvalidatorForOrder(maker common.Address, slot *big.Int) (*big.Int, error) {
	return _Oneinchv6.Contract.BitInvalidatorForOrder(&_Oneinchv6.CallOpts, maker, slot)
}

// CheckPredicate is a free data retrieval call binding the contract method 0x15169dec.
//
// Solidity: function checkPredicate(bytes predicate) view returns(bool)
func (_Oneinchv6 *Oneinchv6Caller) CheckPredicate(opts *bind.CallOpts, predicate []byte) (bool, error) {
	var out []interface{}
	err := _Oneinchv6.contract.Call(opts, &out, "checkPredicate", predicate)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckPredicate is a free data retrieval call binding the contract method 0x15169dec.
//
// Solidity: function checkPredicate(bytes predicate) view returns(bool)
func (_Oneinchv6 *Oneinchv6Session) CheckPredicate(predicate []byte) (bool, error) {
	return _Oneinchv6.Contract.CheckPredicate(&_Oneinchv6.CallOpts, predicate)
}

// CheckPredicate is a free data retrieval call binding the contract method 0x15169dec.
//
// Solidity: function checkPredicate(bytes predicate) view returns(bool)
func (_Oneinchv6 *Oneinchv6CallerSession) CheckPredicate(predicate []byte) (bool, error) {
	return _Oneinchv6.Contract.CheckPredicate(&_Oneinchv6.CallOpts, predicate)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_Oneinchv6 *Oneinchv6Caller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _Oneinchv6.contract.Call(opts, &out, "eip712Domain")

	outstruct := new(struct {
		Fields            [1]byte
		Name              string
		Version           string
		ChainId           *big.Int
		VerifyingContract common.Address
		Salt              [32]byte
		Extensions        []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Fields = *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Version = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.ChainId = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.VerifyingContract = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Salt = *abi.ConvertType(out[5], new([32]byte)).(*[32]byte)
	outstruct.Extensions = *abi.ConvertType(out[6], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_Oneinchv6 *Oneinchv6Session) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _Oneinchv6.Contract.Eip712Domain(&_Oneinchv6.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_Oneinchv6 *Oneinchv6CallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _Oneinchv6.Contract.Eip712Domain(&_Oneinchv6.CallOpts)
}

// Epoch is a free data retrieval call binding the contract method 0xfcea9e4e.
//
// Solidity: function epoch(address maker, uint96 series) view returns(uint256)
func (_Oneinchv6 *Oneinchv6Caller) Epoch(opts *bind.CallOpts, maker common.Address, series *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Oneinchv6.contract.Call(opts, &out, "epoch", maker, series)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Epoch is a free data retrieval call binding the contract method 0xfcea9e4e.
//
// Solidity: function epoch(address maker, uint96 series) view returns(uint256)
func (_Oneinchv6 *Oneinchv6Session) Epoch(maker common.Address, series *big.Int) (*big.Int, error) {
	return _Oneinchv6.Contract.Epoch(&_Oneinchv6.CallOpts, maker, series)
}

// Epoch is a free data retrieval call binding the contract method 0xfcea9e4e.
//
// Solidity: function epoch(address maker, uint96 series) view returns(uint256)
func (_Oneinchv6 *Oneinchv6CallerSession) Epoch(maker common.Address, series *big.Int) (*big.Int, error) {
	return _Oneinchv6.Contract.Epoch(&_Oneinchv6.CallOpts, maker, series)
}

// EpochEquals is a free data retrieval call binding the contract method 0xce3d710a.
//
// Solidity: function epochEquals(address maker, uint256 series, uint256 makerEpoch) view returns(bool)
func (_Oneinchv6 *Oneinchv6Caller) EpochEquals(opts *bind.CallOpts, maker common.Address, series *big.Int, makerEpoch *big.Int) (bool, error) {
	var out []interface{}
	err := _Oneinchv6.contract.Call(opts, &out, "epochEquals", maker, series, makerEpoch)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// EpochEquals is a free data retrieval call binding the contract method 0xce3d710a.
//
// Solidity: function epochEquals(address maker, uint256 series, uint256 makerEpoch) view returns(bool)
func (_Oneinchv6 *Oneinchv6Session) EpochEquals(maker common.Address, series *big.Int, makerEpoch *big.Int) (bool, error) {
	return _Oneinchv6.Contract.EpochEquals(&_Oneinchv6.CallOpts, maker, series, makerEpoch)
}

// EpochEquals is a free data retrieval call binding the contract method 0xce3d710a.
//
// Solidity: function epochEquals(address maker, uint256 series, uint256 makerEpoch) view returns(bool)
func (_Oneinchv6 *Oneinchv6CallerSession) EpochEquals(maker common.Address, series *big.Int, makerEpoch *big.Int) (bool, error) {
	return _Oneinchv6.Contract.EpochEquals(&_Oneinchv6.CallOpts, maker, series, makerEpoch)
}

// Eq is a free data retrieval call binding the contract method 0x6fe7b0ba.
//
// Solidity: function eq(uint256 value, bytes data) view returns(bool)
func (_Oneinchv6 *Oneinchv6Caller) Eq(opts *bind.CallOpts, value *big.Int, data []byte) (bool, error) {
	var out []interface{}
	err := _Oneinchv6.contract.Call(opts, &out, "eq", value, data)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Eq is a free data retrieval call binding the contract method 0x6fe7b0ba.
//
// Solidity: function eq(uint256 value, bytes data) view returns(bool)
func (_Oneinchv6 *Oneinchv6Session) Eq(value *big.Int, data []byte) (bool, error) {
	return _Oneinchv6.Contract.Eq(&_Oneinchv6.CallOpts, value, data)
}

// Eq is a free data retrieval call binding the contract method 0x6fe7b0ba.
//
// Solidity: function eq(uint256 value, bytes data) view returns(bool)
func (_Oneinchv6 *Oneinchv6CallerSession) Eq(value *big.Int, data []byte) (bool, error) {
	return _Oneinchv6.Contract.Eq(&_Oneinchv6.CallOpts, value, data)
}

// Gt is a free data retrieval call binding the contract method 0x4f38e2b8.
//
// Solidity: function gt(uint256 value, bytes data) view returns(bool)
func (_Oneinchv6 *Oneinchv6Caller) Gt(opts *bind.CallOpts, value *big.Int, data []byte) (bool, error) {
	var out []interface{}
	err := _Oneinchv6.contract.Call(opts, &out, "gt", value, data)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Gt is a free data retrieval call binding the contract method 0x4f38e2b8.
//
// Solidity: function gt(uint256 value, bytes data) view returns(bool)
func (_Oneinchv6 *Oneinchv6Session) Gt(value *big.Int, data []byte) (bool, error) {
	return _Oneinchv6.Contract.Gt(&_Oneinchv6.CallOpts, value, data)
}

// Gt is a free data retrieval call binding the contract method 0x4f38e2b8.
//
// Solidity: function gt(uint256 value, bytes data) view returns(bool)
func (_Oneinchv6 *Oneinchv6CallerSession) Gt(value *big.Int, data []byte) (bool, error) {
	return _Oneinchv6.Contract.Gt(&_Oneinchv6.CallOpts, value, data)
}

// HashOrder is a free data retrieval call binding the contract method 0x802b2ef1.
//
// Solidity: function hashOrder((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) order) view returns(bytes32)
func (_Oneinchv6 *Oneinchv6Caller) HashOrder(opts *bind.CallOpts, order IOrderMixinOrder) ([32]byte, error) {
	var out []interface{}
	err := _Oneinchv6.contract.Call(opts, &out, "hashOrder", order)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HashOrder is a free data retrieval call binding the contract method 0x802b2ef1.
//
// Solidity: function hashOrder((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) order) view returns(bytes32)
func (_Oneinchv6 *Oneinchv6Session) HashOrder(order IOrderMixinOrder) ([32]byte, error) {
	return _Oneinchv6.Contract.HashOrder(&_Oneinchv6.CallOpts, order)
}

// HashOrder is a free data retrieval call binding the contract method 0x802b2ef1.
//
// Solidity: function hashOrder((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) order) view returns(bytes32)
func (_Oneinchv6 *Oneinchv6CallerSession) HashOrder(order IOrderMixinOrder) ([32]byte, error) {
	return _Oneinchv6.Contract.HashOrder(&_Oneinchv6.CallOpts, order)
}

// Lt is a free data retrieval call binding the contract method 0xca4ece22.
//
// Solidity: function lt(uint256 value, bytes data) view returns(bool)
func (_Oneinchv6 *Oneinchv6Caller) Lt(opts *bind.CallOpts, value *big.Int, data []byte) (bool, error) {
	var out []interface{}
	err := _Oneinchv6.contract.Call(opts, &out, "lt", value, data)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Lt is a free data retrieval call binding the contract method 0xca4ece22.
//
// Solidity: function lt(uint256 value, bytes data) view returns(bool)
func (_Oneinchv6 *Oneinchv6Session) Lt(value *big.Int, data []byte) (bool, error) {
	return _Oneinchv6.Contract.Lt(&_Oneinchv6.CallOpts, value, data)
}

// Lt is a free data retrieval call binding the contract method 0xca4ece22.
//
// Solidity: function lt(uint256 value, bytes data) view returns(bool)
func (_Oneinchv6 *Oneinchv6CallerSession) Lt(value *big.Int, data []byte) (bool, error) {
	return _Oneinchv6.Contract.Lt(&_Oneinchv6.CallOpts, value, data)
}

// Not is a free data retrieval call binding the contract method 0xbf797959.
//
// Solidity: function not(bytes data) view returns(bool)
func (_Oneinchv6 *Oneinchv6Caller) Not(opts *bind.CallOpts, data []byte) (bool, error) {
	var out []interface{}
	err := _Oneinchv6.contract.Call(opts, &out, "not", data)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Not is a free data retrieval call binding the contract method 0xbf797959.
//
// Solidity: function not(bytes data) view returns(bool)
func (_Oneinchv6 *Oneinchv6Session) Not(data []byte) (bool, error) {
	return _Oneinchv6.Contract.Not(&_Oneinchv6.CallOpts, data)
}

// Not is a free data retrieval call binding the contract method 0xbf797959.
//
// Solidity: function not(bytes data) view returns(bool)
func (_Oneinchv6 *Oneinchv6CallerSession) Not(data []byte) (bool, error) {
	return _Oneinchv6.Contract.Not(&_Oneinchv6.CallOpts, data)
}

// Or is a free data retrieval call binding the contract method 0x74261145.
//
// Solidity: function or(uint256 offsets, bytes data) view returns(bool)
func (_Oneinchv6 *Oneinchv6Caller) Or(opts *bind.CallOpts, offsets *big.Int, data []byte) (bool, error) {
	var out []interface{}
	err := _Oneinchv6.contract.Call(opts, &out, "or", offsets, data)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Or is a free data retrieval call binding the contract method 0x74261145.
//
// Solidity: function or(uint256 offsets, bytes data) view returns(bool)
func (_Oneinchv6 *Oneinchv6Session) Or(offsets *big.Int, data []byte) (bool, error) {
	return _Oneinchv6.Contract.Or(&_Oneinchv6.CallOpts, offsets, data)
}

// Or is a free data retrieval call binding the contract method 0x74261145.
//
// Solidity: function or(uint256 offsets, bytes data) view returns(bool)
func (_Oneinchv6 *Oneinchv6CallerSession) Or(offsets *big.Int, data []byte) (bool, error) {
	return _Oneinchv6.Contract.Or(&_Oneinchv6.CallOpts, offsets, data)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Oneinchv6 *Oneinchv6Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Oneinchv6.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Oneinchv6 *Oneinchv6Session) Owner() (common.Address, error) {
	return _Oneinchv6.Contract.Owner(&_Oneinchv6.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Oneinchv6 *Oneinchv6CallerSession) Owner() (common.Address, error) {
	return _Oneinchv6.Contract.Owner(&_Oneinchv6.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Oneinchv6 *Oneinchv6Caller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Oneinchv6.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Oneinchv6 *Oneinchv6Session) Paused() (bool, error) {
	return _Oneinchv6.Contract.Paused(&_Oneinchv6.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Oneinchv6 *Oneinchv6CallerSession) Paused() (bool, error) {
	return _Oneinchv6.Contract.Paused(&_Oneinchv6.CallOpts)
}

// RawRemainingInvalidatorForOrder is a free data retrieval call binding the contract method 0xc2a40753.
//
// Solidity: function rawRemainingInvalidatorForOrder(address maker, bytes32 orderHash) view returns(uint256)
func (_Oneinchv6 *Oneinchv6Caller) RawRemainingInvalidatorForOrder(opts *bind.CallOpts, maker common.Address, orderHash [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Oneinchv6.contract.Call(opts, &out, "rawRemainingInvalidatorForOrder", maker, orderHash)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RawRemainingInvalidatorForOrder is a free data retrieval call binding the contract method 0xc2a40753.
//
// Solidity: function rawRemainingInvalidatorForOrder(address maker, bytes32 orderHash) view returns(uint256)
func (_Oneinchv6 *Oneinchv6Session) RawRemainingInvalidatorForOrder(maker common.Address, orderHash [32]byte) (*big.Int, error) {
	return _Oneinchv6.Contract.RawRemainingInvalidatorForOrder(&_Oneinchv6.CallOpts, maker, orderHash)
}

// RawRemainingInvalidatorForOrder is a free data retrieval call binding the contract method 0xc2a40753.
//
// Solidity: function rawRemainingInvalidatorForOrder(address maker, bytes32 orderHash) view returns(uint256)
func (_Oneinchv6 *Oneinchv6CallerSession) RawRemainingInvalidatorForOrder(maker common.Address, orderHash [32]byte) (*big.Int, error) {
	return _Oneinchv6.Contract.RawRemainingInvalidatorForOrder(&_Oneinchv6.CallOpts, maker, orderHash)
}

// RemainingInvalidatorForOrder is a free data retrieval call binding the contract method 0x435b9789.
//
// Solidity: function remainingInvalidatorForOrder(address maker, bytes32 orderHash) view returns(uint256)
func (_Oneinchv6 *Oneinchv6Caller) RemainingInvalidatorForOrder(opts *bind.CallOpts, maker common.Address, orderHash [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Oneinchv6.contract.Call(opts, &out, "remainingInvalidatorForOrder", maker, orderHash)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RemainingInvalidatorForOrder is a free data retrieval call binding the contract method 0x435b9789.
//
// Solidity: function remainingInvalidatorForOrder(address maker, bytes32 orderHash) view returns(uint256)
func (_Oneinchv6 *Oneinchv6Session) RemainingInvalidatorForOrder(maker common.Address, orderHash [32]byte) (*big.Int, error) {
	return _Oneinchv6.Contract.RemainingInvalidatorForOrder(&_Oneinchv6.CallOpts, maker, orderHash)
}

// RemainingInvalidatorForOrder is a free data retrieval call binding the contract method 0x435b9789.
//
// Solidity: function remainingInvalidatorForOrder(address maker, bytes32 orderHash) view returns(uint256)
func (_Oneinchv6 *Oneinchv6CallerSession) RemainingInvalidatorForOrder(maker common.Address, orderHash [32]byte) (*big.Int, error) {
	return _Oneinchv6.Contract.RemainingInvalidatorForOrder(&_Oneinchv6.CallOpts, maker, orderHash)
}

// AdvanceEpoch is a paid mutator transaction binding the contract method 0x0d2c7c16.
//
// Solidity: function advanceEpoch(uint96 series, uint256 amount) returns()
func (_Oneinchv6 *Oneinchv6Transactor) AdvanceEpoch(opts *bind.TransactOpts, series *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Oneinchv6.contract.Transact(opts, "advanceEpoch", series, amount)
}

// AdvanceEpoch is a paid mutator transaction binding the contract method 0x0d2c7c16.
//
// Solidity: function advanceEpoch(uint96 series, uint256 amount) returns()
func (_Oneinchv6 *Oneinchv6Session) AdvanceEpoch(series *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Oneinchv6.Contract.AdvanceEpoch(&_Oneinchv6.TransactOpts, series, amount)
}

// AdvanceEpoch is a paid mutator transaction binding the contract method 0x0d2c7c16.
//
// Solidity: function advanceEpoch(uint96 series, uint256 amount) returns()
func (_Oneinchv6 *Oneinchv6TransactorSession) AdvanceEpoch(series *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Oneinchv6.Contract.AdvanceEpoch(&_Oneinchv6.TransactOpts, series, amount)
}

// BitsInvalidateForOrder is a paid mutator transaction binding the contract method 0x05b1ea03.
//
// Solidity: function bitsInvalidateForOrder(uint256 makerTraits, uint256 additionalMask) returns()
func (_Oneinchv6 *Oneinchv6Transactor) BitsInvalidateForOrder(opts *bind.TransactOpts, makerTraits *big.Int, additionalMask *big.Int) (*types.Transaction, error) {
	return _Oneinchv6.contract.Transact(opts, "bitsInvalidateForOrder", makerTraits, additionalMask)
}

// BitsInvalidateForOrder is a paid mutator transaction binding the contract method 0x05b1ea03.
//
// Solidity: function bitsInvalidateForOrder(uint256 makerTraits, uint256 additionalMask) returns()
func (_Oneinchv6 *Oneinchv6Session) BitsInvalidateForOrder(makerTraits *big.Int, additionalMask *big.Int) (*types.Transaction, error) {
	return _Oneinchv6.Contract.BitsInvalidateForOrder(&_Oneinchv6.TransactOpts, makerTraits, additionalMask)
}

// BitsInvalidateForOrder is a paid mutator transaction binding the contract method 0x05b1ea03.
//
// Solidity: function bitsInvalidateForOrder(uint256 makerTraits, uint256 additionalMask) returns()
func (_Oneinchv6 *Oneinchv6TransactorSession) BitsInvalidateForOrder(makerTraits *big.Int, additionalMask *big.Int) (*types.Transaction, error) {
	return _Oneinchv6.Contract.BitsInvalidateForOrder(&_Oneinchv6.TransactOpts, makerTraits, additionalMask)
}

// CancelOrder is a paid mutator transaction binding the contract method 0xb68fb020.
//
// Solidity: function cancelOrder(uint256 makerTraits, bytes32 orderHash) returns()
func (_Oneinchv6 *Oneinchv6Transactor) CancelOrder(opts *bind.TransactOpts, makerTraits *big.Int, orderHash [32]byte) (*types.Transaction, error) {
	return _Oneinchv6.contract.Transact(opts, "cancelOrder", makerTraits, orderHash)
}

// CancelOrder is a paid mutator transaction binding the contract method 0xb68fb020.
//
// Solidity: function cancelOrder(uint256 makerTraits, bytes32 orderHash) returns()
func (_Oneinchv6 *Oneinchv6Session) CancelOrder(makerTraits *big.Int, orderHash [32]byte) (*types.Transaction, error) {
	return _Oneinchv6.Contract.CancelOrder(&_Oneinchv6.TransactOpts, makerTraits, orderHash)
}

// CancelOrder is a paid mutator transaction binding the contract method 0xb68fb020.
//
// Solidity: function cancelOrder(uint256 makerTraits, bytes32 orderHash) returns()
func (_Oneinchv6 *Oneinchv6TransactorSession) CancelOrder(makerTraits *big.Int, orderHash [32]byte) (*types.Transaction, error) {
	return _Oneinchv6.Contract.CancelOrder(&_Oneinchv6.TransactOpts, makerTraits, orderHash)
}

// CancelOrders is a paid mutator transaction binding the contract method 0x89e7c650.
//
// Solidity: function cancelOrders(uint256[] makerTraits, bytes32[] orderHashes) returns()
func (_Oneinchv6 *Oneinchv6Transactor) CancelOrders(opts *bind.TransactOpts, makerTraits []*big.Int, orderHashes [][32]byte) (*types.Transaction, error) {
	return _Oneinchv6.contract.Transact(opts, "cancelOrders", makerTraits, orderHashes)
}

// CancelOrders is a paid mutator transaction binding the contract method 0x89e7c650.
//
// Solidity: function cancelOrders(uint256[] makerTraits, bytes32[] orderHashes) returns()
func (_Oneinchv6 *Oneinchv6Session) CancelOrders(makerTraits []*big.Int, orderHashes [][32]byte) (*types.Transaction, error) {
	return _Oneinchv6.Contract.CancelOrders(&_Oneinchv6.TransactOpts, makerTraits, orderHashes)
}

// CancelOrders is a paid mutator transaction binding the contract method 0x89e7c650.
//
// Solidity: function cancelOrders(uint256[] makerTraits, bytes32[] orderHashes) returns()
func (_Oneinchv6 *Oneinchv6TransactorSession) CancelOrders(makerTraits []*big.Int, orderHashes [][32]byte) (*types.Transaction, error) {
	return _Oneinchv6.Contract.CancelOrders(&_Oneinchv6.TransactOpts, makerTraits, orderHashes)
}

// ClipperSwap is a paid mutator transaction binding the contract method 0xd2d374e5.
//
// Solidity: function clipperSwap(address clipperExchange, uint256 srcToken, address dstToken, uint256 inputAmount, uint256 outputAmount, uint256 goodUntil, bytes32 r, bytes32 vs) payable returns(uint256 returnAmount)
func (_Oneinchv6 *Oneinchv6Transactor) ClipperSwap(opts *bind.TransactOpts, clipperExchange common.Address, srcToken *big.Int, dstToken common.Address, inputAmount *big.Int, outputAmount *big.Int, goodUntil *big.Int, r [32]byte, vs [32]byte) (*types.Transaction, error) {
	return _Oneinchv6.contract.Transact(opts, "clipperSwap", clipperExchange, srcToken, dstToken, inputAmount, outputAmount, goodUntil, r, vs)
}

// ClipperSwap is a paid mutator transaction binding the contract method 0xd2d374e5.
//
// Solidity: function clipperSwap(address clipperExchange, uint256 srcToken, address dstToken, uint256 inputAmount, uint256 outputAmount, uint256 goodUntil, bytes32 r, bytes32 vs) payable returns(uint256 returnAmount)
func (_Oneinchv6 *Oneinchv6Session) ClipperSwap(clipperExchange common.Address, srcToken *big.Int, dstToken common.Address, inputAmount *big.Int, outputAmount *big.Int, goodUntil *big.Int, r [32]byte, vs [32]byte) (*types.Transaction, error) {
	return _Oneinchv6.Contract.ClipperSwap(&_Oneinchv6.TransactOpts, clipperExchange, srcToken, dstToken, inputAmount, outputAmount, goodUntil, r, vs)
}

// ClipperSwap is a paid mutator transaction binding the contract method 0xd2d374e5.
//
// Solidity: function clipperSwap(address clipperExchange, uint256 srcToken, address dstToken, uint256 inputAmount, uint256 outputAmount, uint256 goodUntil, bytes32 r, bytes32 vs) payable returns(uint256 returnAmount)
func (_Oneinchv6 *Oneinchv6TransactorSession) ClipperSwap(clipperExchange common.Address, srcToken *big.Int, dstToken common.Address, inputAmount *big.Int, outputAmount *big.Int, goodUntil *big.Int, r [32]byte, vs [32]byte) (*types.Transaction, error) {
	return _Oneinchv6.Contract.ClipperSwap(&_Oneinchv6.TransactOpts, clipperExchange, srcToken, dstToken, inputAmount, outputAmount, goodUntil, r, vs)
}

// ClipperSwapTo is a paid mutator transaction binding the contract method 0xc4d652af.
//
// Solidity: function clipperSwapTo(address clipperExchange, address recipient, uint256 srcToken, address dstToken, uint256 inputAmount, uint256 outputAmount, uint256 goodUntil, bytes32 r, bytes32 vs) payable returns(uint256 returnAmount)
func (_Oneinchv6 *Oneinchv6Transactor) ClipperSwapTo(opts *bind.TransactOpts, clipperExchange common.Address, recipient common.Address, srcToken *big.Int, dstToken common.Address, inputAmount *big.Int, outputAmount *big.Int, goodUntil *big.Int, r [32]byte, vs [32]byte) (*types.Transaction, error) {
	return _Oneinchv6.contract.Transact(opts, "clipperSwapTo", clipperExchange, recipient, srcToken, dstToken, inputAmount, outputAmount, goodUntil, r, vs)
}

// ClipperSwapTo is a paid mutator transaction binding the contract method 0xc4d652af.
//
// Solidity: function clipperSwapTo(address clipperExchange, address recipient, uint256 srcToken, address dstToken, uint256 inputAmount, uint256 outputAmount, uint256 goodUntil, bytes32 r, bytes32 vs) payable returns(uint256 returnAmount)
func (_Oneinchv6 *Oneinchv6Session) ClipperSwapTo(clipperExchange common.Address, recipient common.Address, srcToken *big.Int, dstToken common.Address, inputAmount *big.Int, outputAmount *big.Int, goodUntil *big.Int, r [32]byte, vs [32]byte) (*types.Transaction, error) {
	return _Oneinchv6.Contract.ClipperSwapTo(&_Oneinchv6.TransactOpts, clipperExchange, recipient, srcToken, dstToken, inputAmount, outputAmount, goodUntil, r, vs)
}

// ClipperSwapTo is a paid mutator transaction binding the contract method 0xc4d652af.
//
// Solidity: function clipperSwapTo(address clipperExchange, address recipient, uint256 srcToken, address dstToken, uint256 inputAmount, uint256 outputAmount, uint256 goodUntil, bytes32 r, bytes32 vs) payable returns(uint256 returnAmount)
func (_Oneinchv6 *Oneinchv6TransactorSession) ClipperSwapTo(clipperExchange common.Address, recipient common.Address, srcToken *big.Int, dstToken common.Address, inputAmount *big.Int, outputAmount *big.Int, goodUntil *big.Int, r [32]byte, vs [32]byte) (*types.Transaction, error) {
	return _Oneinchv6.Contract.ClipperSwapTo(&_Oneinchv6.TransactOpts, clipperExchange, recipient, srcToken, dstToken, inputAmount, outputAmount, goodUntil, r, vs)
}

// CurveSwapCallback is a paid mutator transaction binding the contract method 0xe413f48d.
//
// Solidity: function curveSwapCallback(address , address , address inCoin, uint256 dx, uint256 ) returns()
func (_Oneinchv6 *Oneinchv6Transactor) CurveSwapCallback(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, inCoin common.Address, dx *big.Int, arg4 *big.Int) (*types.Transaction, error) {
	return _Oneinchv6.contract.Transact(opts, "curveSwapCallback", arg0, arg1, inCoin, dx, arg4)
}

// CurveSwapCallback is a paid mutator transaction binding the contract method 0xe413f48d.
//
// Solidity: function curveSwapCallback(address , address , address inCoin, uint256 dx, uint256 ) returns()
func (_Oneinchv6 *Oneinchv6Session) CurveSwapCallback(arg0 common.Address, arg1 common.Address, inCoin common.Address, dx *big.Int, arg4 *big.Int) (*types.Transaction, error) {
	return _Oneinchv6.Contract.CurveSwapCallback(&_Oneinchv6.TransactOpts, arg0, arg1, inCoin, dx, arg4)
}

// CurveSwapCallback is a paid mutator transaction binding the contract method 0xe413f48d.
//
// Solidity: function curveSwapCallback(address , address , address inCoin, uint256 dx, uint256 ) returns()
func (_Oneinchv6 *Oneinchv6TransactorSession) CurveSwapCallback(arg0 common.Address, arg1 common.Address, inCoin common.Address, dx *big.Int, arg4 *big.Int) (*types.Transaction, error) {
	return _Oneinchv6.Contract.CurveSwapCallback(&_Oneinchv6.TransactOpts, arg0, arg1, inCoin, dx, arg4)
}

// EthUnoswap is a paid mutator transaction binding the contract method 0xa76dfc3b.
//
// Solidity: function ethUnoswap(uint256 minReturn, uint256 dex) payable returns(uint256 returnAmount)
func (_Oneinchv6 *Oneinchv6Transactor) EthUnoswap(opts *bind.TransactOpts, minReturn *big.Int, dex *big.Int) (*types.Transaction, error) {
	return _Oneinchv6.contract.Transact(opts, "ethUnoswap", minReturn, dex)
}

// EthUnoswap is a paid mutator transaction binding the contract method 0xa76dfc3b.
//
// Solidity: function ethUnoswap(uint256 minReturn, uint256 dex) payable returns(uint256 returnAmount)
func (_Oneinchv6 *Oneinchv6Session) EthUnoswap(minReturn *big.Int, dex *big.Int) (*types.Transaction, error) {
	return _Oneinchv6.Contract.EthUnoswap(&_Oneinchv6.TransactOpts, minReturn, dex)
}

// EthUnoswap is a paid mutator transaction binding the contract method 0xa76dfc3b.
//
// Solidity: function ethUnoswap(uint256 minReturn, uint256 dex) payable returns(uint256 returnAmount)
func (_Oneinchv6 *Oneinchv6TransactorSession) EthUnoswap(minReturn *big.Int, dex *big.Int) (*types.Transaction, error) {
	return _Oneinchv6.Contract.EthUnoswap(&_Oneinchv6.TransactOpts, minReturn, dex)
}

// EthUnoswap2 is a paid mutator transaction binding the contract method 0x89af926a.
//
// Solidity: function ethUnoswap2(uint256 minReturn, uint256 dex, uint256 dex2) payable returns(uint256 returnAmount)
func (_Oneinchv6 *Oneinchv6Transactor) EthUnoswap2(opts *bind.TransactOpts, minReturn *big.Int, dex *big.Int, dex2 *big.Int) (*types.Transaction, error) {
	return _Oneinchv6.contract.Transact(opts, "ethUnoswap2", minReturn, dex, dex2)
}

// EthUnoswap2 is a paid mutator transaction binding the contract method 0x89af926a.
//
// Solidity: function ethUnoswap2(uint256 minReturn, uint256 dex, uint256 dex2) payable returns(uint256 returnAmount)
func (_Oneinchv6 *Oneinchv6Session) EthUnoswap2(minReturn *big.Int, dex *big.Int, dex2 *big.Int) (*types.Transaction, error) {
	return _Oneinchv6.Contract.EthUnoswap2(&_Oneinchv6.TransactOpts, minReturn, dex, dex2)
}

// EthUnoswap2 is a paid mutator transaction binding the contract method 0x89af926a.
//
// Solidity: function ethUnoswap2(uint256 minReturn, uint256 dex, uint256 dex2) payable returns(uint256 returnAmount)
func (_Oneinchv6 *Oneinchv6TransactorSession) EthUnoswap2(minReturn *big.Int, dex *big.Int, dex2 *big.Int) (*types.Transaction, error) {
	return _Oneinchv6.Contract.EthUnoswap2(&_Oneinchv6.TransactOpts, minReturn, dex, dex2)
}

// EthUnoswap3 is a paid mutator transaction binding the contract method 0x188ac35d.
//
// Solidity: function ethUnoswap3(uint256 minReturn, uint256 dex, uint256 dex2, uint256 dex3) payable returns(uint256 returnAmount)
func (_Oneinchv6 *Oneinchv6Transactor) EthUnoswap3(opts *bind.TransactOpts, minReturn *big.Int, dex *big.Int, dex2 *big.Int, dex3 *big.Int) (*types.Transaction, error) {
	return _Oneinchv6.contract.Transact(opts, "ethUnoswap3", minReturn, dex, dex2, dex3)
}

// EthUnoswap3 is a paid mutator transaction binding the contract method 0x188ac35d.
//
// Solidity: function ethUnoswap3(uint256 minReturn, uint256 dex, uint256 dex2, uint256 dex3) payable returns(uint256 returnAmount)
func (_Oneinchv6 *Oneinchv6Session) EthUnoswap3(minReturn *big.Int, dex *big.Int, dex2 *big.Int, dex3 *big.Int) (*types.Transaction, error) {
	return _Oneinchv6.Contract.EthUnoswap3(&_Oneinchv6.TransactOpts, minReturn, dex, dex2, dex3)
}

// EthUnoswap3 is a paid mutator transaction binding the contract method 0x188ac35d.
//
// Solidity: function ethUnoswap3(uint256 minReturn, uint256 dex, uint256 dex2, uint256 dex3) payable returns(uint256 returnAmount)
func (_Oneinchv6 *Oneinchv6TransactorSession) EthUnoswap3(minReturn *big.Int, dex *big.Int, dex2 *big.Int, dex3 *big.Int) (*types.Transaction, error) {
	return _Oneinchv6.Contract.EthUnoswap3(&_Oneinchv6.TransactOpts, minReturn, dex, dex2, dex3)
}

// EthUnoswapTo is a paid mutator transaction binding the contract method 0x175accdc.
//
// Solidity: function ethUnoswapTo(uint256 to, uint256 minReturn, uint256 dex) payable returns(uint256 returnAmount)
func (_Oneinchv6 *Oneinchv6Transactor) EthUnoswapTo(opts *bind.TransactOpts, to *big.Int, minReturn *big.Int, dex *big.Int) (*types.Transaction, error) {
	return _Oneinchv6.contract.Transact(opts, "ethUnoswapTo", to, minReturn, dex)
}

// EthUnoswapTo is a paid mutator transaction binding the contract method 0x175accdc.
//
// Solidity: function ethUnoswapTo(uint256 to, uint256 minReturn, uint256 dex) payable returns(uint256 returnAmount)
func (_Oneinchv6 *Oneinchv6Session) EthUnoswapTo(to *big.Int, minReturn *big.Int, dex *big.Int) (*types.Transaction, error) {
	return _Oneinchv6.Contract.EthUnoswapTo(&_Oneinchv6.TransactOpts, to, minReturn, dex)
}

// EthUnoswapTo is a paid mutator transaction binding the contract method 0x175accdc.
//
// Solidity: function ethUnoswapTo(uint256 to, uint256 minReturn, uint256 dex) payable returns(uint256 returnAmount)
func (_Oneinchv6 *Oneinchv6TransactorSession) EthUnoswapTo(to *big.Int, minReturn *big.Int, dex *big.Int) (*types.Transaction, error) {
	return _Oneinchv6.Contract.EthUnoswapTo(&_Oneinchv6.TransactOpts, to, minReturn, dex)
}

// EthUnoswapTo2 is a paid mutator transaction binding the contract method 0x0f449d71.
//
// Solidity: function ethUnoswapTo2(uint256 to, uint256 minReturn, uint256 dex, uint256 dex2) payable returns(uint256 returnAmount)
func (_Oneinchv6 *Oneinchv6Transactor) EthUnoswapTo2(opts *bind.TransactOpts, to *big.Int, minReturn *big.Int, dex *big.Int, dex2 *big.Int) (*types.Transaction, error) {
	return _Oneinchv6.contract.Transact(opts, "ethUnoswapTo2", to, minReturn, dex, dex2)
}

// EthUnoswapTo2 is a paid mutator transaction binding the contract method 0x0f449d71.
//
// Solidity: function ethUnoswapTo2(uint256 to, uint256 minReturn, uint256 dex, uint256 dex2) payable returns(uint256 returnAmount)
func (_Oneinchv6 *Oneinchv6Session) EthUnoswapTo2(to *big.Int, minReturn *big.Int, dex *big.Int, dex2 *big.Int) (*types.Transaction, error) {
	return _Oneinchv6.Contract.EthUnoswapTo2(&_Oneinchv6.TransactOpts, to, minReturn, dex, dex2)
}

// EthUnoswapTo2 is a paid mutator transaction binding the contract method 0x0f449d71.
//
// Solidity: function ethUnoswapTo2(uint256 to, uint256 minReturn, uint256 dex, uint256 dex2) payable returns(uint256 returnAmount)
func (_Oneinchv6 *Oneinchv6TransactorSession) EthUnoswapTo2(to *big.Int, minReturn *big.Int, dex *big.Int, dex2 *big.Int) (*types.Transaction, error) {
	return _Oneinchv6.Contract.EthUnoswapTo2(&_Oneinchv6.TransactOpts, to, minReturn, dex, dex2)
}

// EthUnoswapTo3 is a paid mutator transaction binding the contract method 0x493189f0.
//
// Solidity: function ethUnoswapTo3(uint256 to, uint256 minReturn, uint256 dex, uint256 dex2, uint256 dex3) payable returns(uint256 returnAmount)
func (_Oneinchv6 *Oneinchv6Transactor) EthUnoswapTo3(opts *bind.TransactOpts, to *big.Int, minReturn *big.Int, dex *big.Int, dex2 *big.Int, dex3 *big.Int) (*types.Transaction, error) {
	return _Oneinchv6.contract.Transact(opts, "ethUnoswapTo3", to, minReturn, dex, dex2, dex3)
}

// EthUnoswapTo3 is a paid mutator transaction binding the contract method 0x493189f0.
//
// Solidity: function ethUnoswapTo3(uint256 to, uint256 minReturn, uint256 dex, uint256 dex2, uint256 dex3) payable returns(uint256 returnAmount)
func (_Oneinchv6 *Oneinchv6Session) EthUnoswapTo3(to *big.Int, minReturn *big.Int, dex *big.Int, dex2 *big.Int, dex3 *big.Int) (*types.Transaction, error) {
	return _Oneinchv6.Contract.EthUnoswapTo3(&_Oneinchv6.TransactOpts, to, minReturn, dex, dex2, dex3)
}

// EthUnoswapTo3 is a paid mutator transaction binding the contract method 0x493189f0.
//
// Solidity: function ethUnoswapTo3(uint256 to, uint256 minReturn, uint256 dex, uint256 dex2, uint256 dex3) payable returns(uint256 returnAmount)
func (_Oneinchv6 *Oneinchv6TransactorSession) EthUnoswapTo3(to *big.Int, minReturn *big.Int, dex *big.Int, dex2 *big.Int, dex3 *big.Int) (*types.Transaction, error) {
	return _Oneinchv6.Contract.EthUnoswapTo3(&_Oneinchv6.TransactOpts, to, minReturn, dex, dex2, dex3)
}

// FillContractOrder is a paid mutator transaction binding the contract method 0xcc713a04.
//
// Solidity: function fillContractOrder((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) order, bytes signature, uint256 amount, uint256 takerTraits) returns(uint256, uint256, bytes32)
func (_Oneinchv6 *Oneinchv6Transactor) FillContractOrder(opts *bind.TransactOpts, order IOrderMixinOrder, signature []byte, amount *big.Int, takerTraits *big.Int) (*types.Transaction, error) {
	return _Oneinchv6.contract.Transact(opts, "fillContractOrder", order, signature, amount, takerTraits)
}

// FillContractOrder is a paid mutator transaction binding the contract method 0xcc713a04.
//
// Solidity: function fillContractOrder((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) order, bytes signature, uint256 amount, uint256 takerTraits) returns(uint256, uint256, bytes32)
func (_Oneinchv6 *Oneinchv6Session) FillContractOrder(order IOrderMixinOrder, signature []byte, amount *big.Int, takerTraits *big.Int) (*types.Transaction, error) {
	return _Oneinchv6.Contract.FillContractOrder(&_Oneinchv6.TransactOpts, order, signature, amount, takerTraits)
}

// FillContractOrder is a paid mutator transaction binding the contract method 0xcc713a04.
//
// Solidity: function fillContractOrder((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) order, bytes signature, uint256 amount, uint256 takerTraits) returns(uint256, uint256, bytes32)
func (_Oneinchv6 *Oneinchv6TransactorSession) FillContractOrder(order IOrderMixinOrder, signature []byte, amount *big.Int, takerTraits *big.Int) (*types.Transaction, error) {
	return _Oneinchv6.Contract.FillContractOrder(&_Oneinchv6.TransactOpts, order, signature, amount, takerTraits)
}

// FillContractOrderArgs is a paid mutator transaction binding the contract method 0x56a75868.
//
// Solidity: function fillContractOrderArgs((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) order, bytes signature, uint256 amount, uint256 takerTraits, bytes args) returns(uint256, uint256, bytes32)
func (_Oneinchv6 *Oneinchv6Transactor) FillContractOrderArgs(opts *bind.TransactOpts, order IOrderMixinOrder, signature []byte, amount *big.Int, takerTraits *big.Int, args []byte) (*types.Transaction, error) {
	return _Oneinchv6.contract.Transact(opts, "fillContractOrderArgs", order, signature, amount, takerTraits, args)
}

// FillContractOrderArgs is a paid mutator transaction binding the contract method 0x56a75868.
//
// Solidity: function fillContractOrderArgs((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) order, bytes signature, uint256 amount, uint256 takerTraits, bytes args) returns(uint256, uint256, bytes32)
func (_Oneinchv6 *Oneinchv6Session) FillContractOrderArgs(order IOrderMixinOrder, signature []byte, amount *big.Int, takerTraits *big.Int, args []byte) (*types.Transaction, error) {
	return _Oneinchv6.Contract.FillContractOrderArgs(&_Oneinchv6.TransactOpts, order, signature, amount, takerTraits, args)
}

// FillContractOrderArgs is a paid mutator transaction binding the contract method 0x56a75868.
//
// Solidity: function fillContractOrderArgs((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) order, bytes signature, uint256 amount, uint256 takerTraits, bytes args) returns(uint256, uint256, bytes32)
func (_Oneinchv6 *Oneinchv6TransactorSession) FillContractOrderArgs(order IOrderMixinOrder, signature []byte, amount *big.Int, takerTraits *big.Int, args []byte) (*types.Transaction, error) {
	return _Oneinchv6.Contract.FillContractOrderArgs(&_Oneinchv6.TransactOpts, order, signature, amount, takerTraits, args)
}

// FillOrder is a paid mutator transaction binding the contract method 0x9fda64bd.
//
// Solidity: function fillOrder((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) order, bytes32 r, bytes32 vs, uint256 amount, uint256 takerTraits) payable returns(uint256, uint256, bytes32)
func (_Oneinchv6 *Oneinchv6Transactor) FillOrder(opts *bind.TransactOpts, order IOrderMixinOrder, r [32]byte, vs [32]byte, amount *big.Int, takerTraits *big.Int) (*types.Transaction, error) {
	return _Oneinchv6.contract.Transact(opts, "fillOrder", order, r, vs, amount, takerTraits)
}

// FillOrder is a paid mutator transaction binding the contract method 0x9fda64bd.
//
// Solidity: function fillOrder((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) order, bytes32 r, bytes32 vs, uint256 amount, uint256 takerTraits) payable returns(uint256, uint256, bytes32)
func (_Oneinchv6 *Oneinchv6Session) FillOrder(order IOrderMixinOrder, r [32]byte, vs [32]byte, amount *big.Int, takerTraits *big.Int) (*types.Transaction, error) {
	return _Oneinchv6.Contract.FillOrder(&_Oneinchv6.TransactOpts, order, r, vs, amount, takerTraits)
}

// FillOrder is a paid mutator transaction binding the contract method 0x9fda64bd.
//
// Solidity: function fillOrder((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) order, bytes32 r, bytes32 vs, uint256 amount, uint256 takerTraits) payable returns(uint256, uint256, bytes32)
func (_Oneinchv6 *Oneinchv6TransactorSession) FillOrder(order IOrderMixinOrder, r [32]byte, vs [32]byte, amount *big.Int, takerTraits *big.Int) (*types.Transaction, error) {
	return _Oneinchv6.Contract.FillOrder(&_Oneinchv6.TransactOpts, order, r, vs, amount, takerTraits)
}

// FillOrderArgs is a paid mutator transaction binding the contract method 0xf497df75.
//
// Solidity: function fillOrderArgs((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) order, bytes32 r, bytes32 vs, uint256 amount, uint256 takerTraits, bytes args) payable returns(uint256, uint256, bytes32)
func (_Oneinchv6 *Oneinchv6Transactor) FillOrderArgs(opts *bind.TransactOpts, order IOrderMixinOrder, r [32]byte, vs [32]byte, amount *big.Int, takerTraits *big.Int, args []byte) (*types.Transaction, error) {
	return _Oneinchv6.contract.Transact(opts, "fillOrderArgs", order, r, vs, amount, takerTraits, args)
}

// FillOrderArgs is a paid mutator transaction binding the contract method 0xf497df75.
//
// Solidity: function fillOrderArgs((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) order, bytes32 r, bytes32 vs, uint256 amount, uint256 takerTraits, bytes args) payable returns(uint256, uint256, bytes32)
func (_Oneinchv6 *Oneinchv6Session) FillOrderArgs(order IOrderMixinOrder, r [32]byte, vs [32]byte, amount *big.Int, takerTraits *big.Int, args []byte) (*types.Transaction, error) {
	return _Oneinchv6.Contract.FillOrderArgs(&_Oneinchv6.TransactOpts, order, r, vs, amount, takerTraits, args)
}

// FillOrderArgs is a paid mutator transaction binding the contract method 0xf497df75.
//
// Solidity: function fillOrderArgs((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) order, bytes32 r, bytes32 vs, uint256 amount, uint256 takerTraits, bytes args) payable returns(uint256, uint256, bytes32)
func (_Oneinchv6 *Oneinchv6TransactorSession) FillOrderArgs(order IOrderMixinOrder, r [32]byte, vs [32]byte, amount *big.Int, takerTraits *big.Int, args []byte) (*types.Transaction, error) {
	return _Oneinchv6.Contract.FillOrderArgs(&_Oneinchv6.TransactOpts, order, r, vs, amount, takerTraits, args)
}

// IncreaseEpoch is a paid mutator transaction binding the contract method 0xc3cf8043.
//
// Solidity: function increaseEpoch(uint96 series) returns()
func (_Oneinchv6 *Oneinchv6Transactor) IncreaseEpoch(opts *bind.TransactOpts, series *big.Int) (*types.Transaction, error) {
	return _Oneinchv6.contract.Transact(opts, "increaseEpoch", series)
}

// IncreaseEpoch is a paid mutator transaction binding the contract method 0xc3cf8043.
//
// Solidity: function increaseEpoch(uint96 series) returns()
func (_Oneinchv6 *Oneinchv6Session) IncreaseEpoch(series *big.Int) (*types.Transaction, error) {
	return _Oneinchv6.Contract.IncreaseEpoch(&_Oneinchv6.TransactOpts, series)
}

// IncreaseEpoch is a paid mutator transaction binding the contract method 0xc3cf8043.
//
// Solidity: function increaseEpoch(uint96 series) returns()
func (_Oneinchv6 *Oneinchv6TransactorSession) IncreaseEpoch(series *big.Int) (*types.Transaction, error) {
	return _Oneinchv6.Contract.IncreaseEpoch(&_Oneinchv6.TransactOpts, series)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Oneinchv6 *Oneinchv6Transactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Oneinchv6.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Oneinchv6 *Oneinchv6Session) Pause() (*types.Transaction, error) {
	return _Oneinchv6.Contract.Pause(&_Oneinchv6.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Oneinchv6 *Oneinchv6TransactorSession) Pause() (*types.Transaction, error) {
	return _Oneinchv6.Contract.Pause(&_Oneinchv6.TransactOpts)
}

// PermitAndCall is a paid mutator transaction binding the contract method 0x5816d723.
//
// Solidity: function permitAndCall(bytes permit, bytes action) payable returns()
func (_Oneinchv6 *Oneinchv6Transactor) PermitAndCall(opts *bind.TransactOpts, permit []byte, action []byte) (*types.Transaction, error) {
	return _Oneinchv6.contract.Transact(opts, "permitAndCall", permit, action)
}

// PermitAndCall is a paid mutator transaction binding the contract method 0x5816d723.
//
// Solidity: function permitAndCall(bytes permit, bytes action) payable returns()
func (_Oneinchv6 *Oneinchv6Session) PermitAndCall(permit []byte, action []byte) (*types.Transaction, error) {
	return _Oneinchv6.Contract.PermitAndCall(&_Oneinchv6.TransactOpts, permit, action)
}

// PermitAndCall is a paid mutator transaction binding the contract method 0x5816d723.
//
// Solidity: function permitAndCall(bytes permit, bytes action) payable returns()
func (_Oneinchv6 *Oneinchv6TransactorSession) PermitAndCall(permit []byte, action []byte) (*types.Transaction, error) {
	return _Oneinchv6.Contract.PermitAndCall(&_Oneinchv6.TransactOpts, permit, action)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Oneinchv6 *Oneinchv6Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Oneinchv6.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Oneinchv6 *Oneinchv6Session) RenounceOwnership() (*types.Transaction, error) {
	return _Oneinchv6.Contract.RenounceOwnership(&_Oneinchv6.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Oneinchv6 *Oneinchv6TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Oneinchv6.Contract.RenounceOwnership(&_Oneinchv6.TransactOpts)
}

// RescueFunds is a paid mutator transaction binding the contract method 0x78e3214f.
//
// Solidity: function rescueFunds(address token, uint256 amount) returns()
func (_Oneinchv6 *Oneinchv6Transactor) RescueFunds(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Oneinchv6.contract.Transact(opts, "rescueFunds", token, amount)
}

// RescueFunds is a paid mutator transaction binding the contract method 0x78e3214f.
//
// Solidity: function rescueFunds(address token, uint256 amount) returns()
func (_Oneinchv6 *Oneinchv6Session) RescueFunds(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Oneinchv6.Contract.RescueFunds(&_Oneinchv6.TransactOpts, token, amount)
}

// RescueFunds is a paid mutator transaction binding the contract method 0x78e3214f.
//
// Solidity: function rescueFunds(address token, uint256 amount) returns()
func (_Oneinchv6 *Oneinchv6TransactorSession) RescueFunds(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Oneinchv6.Contract.RescueFunds(&_Oneinchv6.TransactOpts, token, amount)
}

// Simulate is a paid mutator transaction binding the contract method 0xbd61951d.
//
// Solidity: function simulate(address target, bytes data) returns()
func (_Oneinchv6 *Oneinchv6Transactor) Simulate(opts *bind.TransactOpts, target common.Address, data []byte) (*types.Transaction, error) {
	return _Oneinchv6.contract.Transact(opts, "simulate", target, data)
}

// Simulate is a paid mutator transaction binding the contract method 0xbd61951d.
//
// Solidity: function simulate(address target, bytes data) returns()
func (_Oneinchv6 *Oneinchv6Session) Simulate(target common.Address, data []byte) (*types.Transaction, error) {
	return _Oneinchv6.Contract.Simulate(&_Oneinchv6.TransactOpts, target, data)
}

// Simulate is a paid mutator transaction binding the contract method 0xbd61951d.
//
// Solidity: function simulate(address target, bytes data) returns()
func (_Oneinchv6 *Oneinchv6TransactorSession) Simulate(target common.Address, data []byte) (*types.Transaction, error) {
	return _Oneinchv6.Contract.Simulate(&_Oneinchv6.TransactOpts, target, data)
}

// Swap is a paid mutator transaction binding the contract method 0x07ed2379.
//
// Solidity: function swap(address executor, (address,address,address,address,uint256,uint256,uint256) desc, bytes data) payable returns(uint256 returnAmount, uint256 spentAmount)
func (_Oneinchv6 *Oneinchv6Transactor) Swap(opts *bind.TransactOpts, executor common.Address, desc GenericRouterSwapDescription, data []byte) (*types.Transaction, error) {
	return _Oneinchv6.contract.Transact(opts, "swap", executor, desc, data)
}

// Swap is a paid mutator transaction binding the contract method 0x07ed2379.
//
// Solidity: function swap(address executor, (address,address,address,address,uint256,uint256,uint256) desc, bytes data) payable returns(uint256 returnAmount, uint256 spentAmount)
func (_Oneinchv6 *Oneinchv6Session) Swap(executor common.Address, desc GenericRouterSwapDescription, data []byte) (*types.Transaction, error) {
	return _Oneinchv6.Contract.Swap(&_Oneinchv6.TransactOpts, executor, desc, data)
}

// Swap is a paid mutator transaction binding the contract method 0x07ed2379.
//
// Solidity: function swap(address executor, (address,address,address,address,uint256,uint256,uint256) desc, bytes data) payable returns(uint256 returnAmount, uint256 spentAmount)
func (_Oneinchv6 *Oneinchv6TransactorSession) Swap(executor common.Address, desc GenericRouterSwapDescription, data []byte) (*types.Transaction, error) {
	return _Oneinchv6.Contract.Swap(&_Oneinchv6.TransactOpts, executor, desc, data)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Oneinchv6 *Oneinchv6Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Oneinchv6.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Oneinchv6 *Oneinchv6Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Oneinchv6.Contract.TransferOwnership(&_Oneinchv6.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Oneinchv6 *Oneinchv6TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Oneinchv6.Contract.TransferOwnership(&_Oneinchv6.TransactOpts, newOwner)
}

// UniswapV3SwapCallback is a paid mutator transaction binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes ) returns()
func (_Oneinchv6 *Oneinchv6Transactor) UniswapV3SwapCallback(opts *bind.TransactOpts, amount0Delta *big.Int, amount1Delta *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _Oneinchv6.contract.Transact(opts, "uniswapV3SwapCallback", amount0Delta, amount1Delta, arg2)
}

// UniswapV3SwapCallback is a paid mutator transaction binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes ) returns()
func (_Oneinchv6 *Oneinchv6Session) UniswapV3SwapCallback(amount0Delta *big.Int, amount1Delta *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _Oneinchv6.Contract.UniswapV3SwapCallback(&_Oneinchv6.TransactOpts, amount0Delta, amount1Delta, arg2)
}

// UniswapV3SwapCallback is a paid mutator transaction binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes ) returns()
func (_Oneinchv6 *Oneinchv6TransactorSession) UniswapV3SwapCallback(amount0Delta *big.Int, amount1Delta *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _Oneinchv6.Contract.UniswapV3SwapCallback(&_Oneinchv6.TransactOpts, amount0Delta, amount1Delta, arg2)
}

// Unoswap is a paid mutator transaction binding the contract method 0x83800a8e.
//
// Solidity: function unoswap(uint256 token, uint256 amount, uint256 minReturn, uint256 dex) returns(uint256 returnAmount)
func (_Oneinchv6 *Oneinchv6Transactor) Unoswap(opts *bind.TransactOpts, token *big.Int, amount *big.Int, minReturn *big.Int, dex *big.Int) (*types.Transaction, error) {
	return _Oneinchv6.contract.Transact(opts, "unoswap", token, amount, minReturn, dex)
}

// Unoswap is a paid mutator transaction binding the contract method 0x83800a8e.
//
// Solidity: function unoswap(uint256 token, uint256 amount, uint256 minReturn, uint256 dex) returns(uint256 returnAmount)
func (_Oneinchv6 *Oneinchv6Session) Unoswap(token *big.Int, amount *big.Int, minReturn *big.Int, dex *big.Int) (*types.Transaction, error) {
	return _Oneinchv6.Contract.Unoswap(&_Oneinchv6.TransactOpts, token, amount, minReturn, dex)
}

// Unoswap is a paid mutator transaction binding the contract method 0x83800a8e.
//
// Solidity: function unoswap(uint256 token, uint256 amount, uint256 minReturn, uint256 dex) returns(uint256 returnAmount)
func (_Oneinchv6 *Oneinchv6TransactorSession) Unoswap(token *big.Int, amount *big.Int, minReturn *big.Int, dex *big.Int) (*types.Transaction, error) {
	return _Oneinchv6.Contract.Unoswap(&_Oneinchv6.TransactOpts, token, amount, minReturn, dex)
}

// Unoswap2 is a paid mutator transaction binding the contract method 0x8770ba91.
//
// Solidity: function unoswap2(uint256 token, uint256 amount, uint256 minReturn, uint256 dex, uint256 dex2) returns(uint256 returnAmount)
func (_Oneinchv6 *Oneinchv6Transactor) Unoswap2(opts *bind.TransactOpts, token *big.Int, amount *big.Int, minReturn *big.Int, dex *big.Int, dex2 *big.Int) (*types.Transaction, error) {
	return _Oneinchv6.contract.Transact(opts, "unoswap2", token, amount, minReturn, dex, dex2)
}

// Unoswap2 is a paid mutator transaction binding the contract method 0x8770ba91.
//
// Solidity: function unoswap2(uint256 token, uint256 amount, uint256 minReturn, uint256 dex, uint256 dex2) returns(uint256 returnAmount)
func (_Oneinchv6 *Oneinchv6Session) Unoswap2(token *big.Int, amount *big.Int, minReturn *big.Int, dex *big.Int, dex2 *big.Int) (*types.Transaction, error) {
	return _Oneinchv6.Contract.Unoswap2(&_Oneinchv6.TransactOpts, token, amount, minReturn, dex, dex2)
}

// Unoswap2 is a paid mutator transaction binding the contract method 0x8770ba91.
//
// Solidity: function unoswap2(uint256 token, uint256 amount, uint256 minReturn, uint256 dex, uint256 dex2) returns(uint256 returnAmount)
func (_Oneinchv6 *Oneinchv6TransactorSession) Unoswap2(token *big.Int, amount *big.Int, minReturn *big.Int, dex *big.Int, dex2 *big.Int) (*types.Transaction, error) {
	return _Oneinchv6.Contract.Unoswap2(&_Oneinchv6.TransactOpts, token, amount, minReturn, dex, dex2)
}

// Unoswap3 is a paid mutator transaction binding the contract method 0x19367472.
//
// Solidity: function unoswap3(uint256 token, uint256 amount, uint256 minReturn, uint256 dex, uint256 dex2, uint256 dex3) returns(uint256 returnAmount)
func (_Oneinchv6 *Oneinchv6Transactor) Unoswap3(opts *bind.TransactOpts, token *big.Int, amount *big.Int, minReturn *big.Int, dex *big.Int, dex2 *big.Int, dex3 *big.Int) (*types.Transaction, error) {
	return _Oneinchv6.contract.Transact(opts, "unoswap3", token, amount, minReturn, dex, dex2, dex3)
}

// Unoswap3 is a paid mutator transaction binding the contract method 0x19367472.
//
// Solidity: function unoswap3(uint256 token, uint256 amount, uint256 minReturn, uint256 dex, uint256 dex2, uint256 dex3) returns(uint256 returnAmount)
func (_Oneinchv6 *Oneinchv6Session) Unoswap3(token *big.Int, amount *big.Int, minReturn *big.Int, dex *big.Int, dex2 *big.Int, dex3 *big.Int) (*types.Transaction, error) {
	return _Oneinchv6.Contract.Unoswap3(&_Oneinchv6.TransactOpts, token, amount, minReturn, dex, dex2, dex3)
}

// Unoswap3 is a paid mutator transaction binding the contract method 0x19367472.
//
// Solidity: function unoswap3(uint256 token, uint256 amount, uint256 minReturn, uint256 dex, uint256 dex2, uint256 dex3) returns(uint256 returnAmount)
func (_Oneinchv6 *Oneinchv6TransactorSession) Unoswap3(token *big.Int, amount *big.Int, minReturn *big.Int, dex *big.Int, dex2 *big.Int, dex3 *big.Int) (*types.Transaction, error) {
	return _Oneinchv6.Contract.Unoswap3(&_Oneinchv6.TransactOpts, token, amount, minReturn, dex, dex2, dex3)
}

// UnoswapTo is a paid mutator transaction binding the contract method 0xe2c95c82.
//
// Solidity: function unoswapTo(uint256 to, uint256 token, uint256 amount, uint256 minReturn, uint256 dex) returns(uint256 returnAmount)
func (_Oneinchv6 *Oneinchv6Transactor) UnoswapTo(opts *bind.TransactOpts, to *big.Int, token *big.Int, amount *big.Int, minReturn *big.Int, dex *big.Int) (*types.Transaction, error) {
	return _Oneinchv6.contract.Transact(opts, "unoswapTo", to, token, amount, minReturn, dex)
}

// UnoswapTo is a paid mutator transaction binding the contract method 0xe2c95c82.
//
// Solidity: function unoswapTo(uint256 to, uint256 token, uint256 amount, uint256 minReturn, uint256 dex) returns(uint256 returnAmount)
func (_Oneinchv6 *Oneinchv6Session) UnoswapTo(to *big.Int, token *big.Int, amount *big.Int, minReturn *big.Int, dex *big.Int) (*types.Transaction, error) {
	return _Oneinchv6.Contract.UnoswapTo(&_Oneinchv6.TransactOpts, to, token, amount, minReturn, dex)
}

// UnoswapTo is a paid mutator transaction binding the contract method 0xe2c95c82.
//
// Solidity: function unoswapTo(uint256 to, uint256 token, uint256 amount, uint256 minReturn, uint256 dex) returns(uint256 returnAmount)
func (_Oneinchv6 *Oneinchv6TransactorSession) UnoswapTo(to *big.Int, token *big.Int, amount *big.Int, minReturn *big.Int, dex *big.Int) (*types.Transaction, error) {
	return _Oneinchv6.Contract.UnoswapTo(&_Oneinchv6.TransactOpts, to, token, amount, minReturn, dex)
}

// UnoswapTo2 is a paid mutator transaction binding the contract method 0xea76dddf.
//
// Solidity: function unoswapTo2(uint256 to, uint256 token, uint256 amount, uint256 minReturn, uint256 dex, uint256 dex2) returns(uint256 returnAmount)
func (_Oneinchv6 *Oneinchv6Transactor) UnoswapTo2(opts *bind.TransactOpts, to *big.Int, token *big.Int, amount *big.Int, minReturn *big.Int, dex *big.Int, dex2 *big.Int) (*types.Transaction, error) {
	return _Oneinchv6.contract.Transact(opts, "unoswapTo2", to, token, amount, minReturn, dex, dex2)
}

// UnoswapTo2 is a paid mutator transaction binding the contract method 0xea76dddf.
//
// Solidity: function unoswapTo2(uint256 to, uint256 token, uint256 amount, uint256 minReturn, uint256 dex, uint256 dex2) returns(uint256 returnAmount)
func (_Oneinchv6 *Oneinchv6Session) UnoswapTo2(to *big.Int, token *big.Int, amount *big.Int, minReturn *big.Int, dex *big.Int, dex2 *big.Int) (*types.Transaction, error) {
	return _Oneinchv6.Contract.UnoswapTo2(&_Oneinchv6.TransactOpts, to, token, amount, minReturn, dex, dex2)
}

// UnoswapTo2 is a paid mutator transaction binding the contract method 0xea76dddf.
//
// Solidity: function unoswapTo2(uint256 to, uint256 token, uint256 amount, uint256 minReturn, uint256 dex, uint256 dex2) returns(uint256 returnAmount)
func (_Oneinchv6 *Oneinchv6TransactorSession) UnoswapTo2(to *big.Int, token *big.Int, amount *big.Int, minReturn *big.Int, dex *big.Int, dex2 *big.Int) (*types.Transaction, error) {
	return _Oneinchv6.Contract.UnoswapTo2(&_Oneinchv6.TransactOpts, to, token, amount, minReturn, dex, dex2)
}

// UnoswapTo3 is a paid mutator transaction binding the contract method 0xf7a70056.
//
// Solidity: function unoswapTo3(uint256 to, uint256 token, uint256 amount, uint256 minReturn, uint256 dex, uint256 dex2, uint256 dex3) returns(uint256 returnAmount)
func (_Oneinchv6 *Oneinchv6Transactor) UnoswapTo3(opts *bind.TransactOpts, to *big.Int, token *big.Int, amount *big.Int, minReturn *big.Int, dex *big.Int, dex2 *big.Int, dex3 *big.Int) (*types.Transaction, error) {
	return _Oneinchv6.contract.Transact(opts, "unoswapTo3", to, token, amount, minReturn, dex, dex2, dex3)
}

// UnoswapTo3 is a paid mutator transaction binding the contract method 0xf7a70056.
//
// Solidity: function unoswapTo3(uint256 to, uint256 token, uint256 amount, uint256 minReturn, uint256 dex, uint256 dex2, uint256 dex3) returns(uint256 returnAmount)
func (_Oneinchv6 *Oneinchv6Session) UnoswapTo3(to *big.Int, token *big.Int, amount *big.Int, minReturn *big.Int, dex *big.Int, dex2 *big.Int, dex3 *big.Int) (*types.Transaction, error) {
	return _Oneinchv6.Contract.UnoswapTo3(&_Oneinchv6.TransactOpts, to, token, amount, minReturn, dex, dex2, dex3)
}

// UnoswapTo3 is a paid mutator transaction binding the contract method 0xf7a70056.
//
// Solidity: function unoswapTo3(uint256 to, uint256 token, uint256 amount, uint256 minReturn, uint256 dex, uint256 dex2, uint256 dex3) returns(uint256 returnAmount)
func (_Oneinchv6 *Oneinchv6TransactorSession) UnoswapTo3(to *big.Int, token *big.Int, amount *big.Int, minReturn *big.Int, dex *big.Int, dex2 *big.Int, dex3 *big.Int) (*types.Transaction, error) {
	return _Oneinchv6.Contract.UnoswapTo3(&_Oneinchv6.TransactOpts, to, token, amount, minReturn, dex, dex2, dex3)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Oneinchv6 *Oneinchv6Transactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Oneinchv6.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Oneinchv6 *Oneinchv6Session) Unpause() (*types.Transaction, error) {
	return _Oneinchv6.Contract.Unpause(&_Oneinchv6.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Oneinchv6 *Oneinchv6TransactorSession) Unpause() (*types.Transaction, error) {
	return _Oneinchv6.Contract.Unpause(&_Oneinchv6.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Oneinchv6 *Oneinchv6Transactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Oneinchv6.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Oneinchv6 *Oneinchv6Session) Receive() (*types.Transaction, error) {
	return _Oneinchv6.Contract.Receive(&_Oneinchv6.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Oneinchv6 *Oneinchv6TransactorSession) Receive() (*types.Transaction, error) {
	return _Oneinchv6.Contract.Receive(&_Oneinchv6.TransactOpts)
}

// Oneinchv6BitInvalidatorUpdatedIterator is returned from FilterBitInvalidatorUpdated and is used to iterate over the raw logs and unpacked data for BitInvalidatorUpdated events raised by the Oneinchv6 contract.
type Oneinchv6BitInvalidatorUpdatedIterator struct {
	Event *Oneinchv6BitInvalidatorUpdated // Event containing the contract specifics and raw log

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
func (it *Oneinchv6BitInvalidatorUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Oneinchv6BitInvalidatorUpdated)
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
		it.Event = new(Oneinchv6BitInvalidatorUpdated)
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
func (it *Oneinchv6BitInvalidatorUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Oneinchv6BitInvalidatorUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Oneinchv6BitInvalidatorUpdated represents a BitInvalidatorUpdated event raised by the Oneinchv6 contract.
type Oneinchv6BitInvalidatorUpdated struct {
	Maker     common.Address
	SlotIndex *big.Int
	SlotValue *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterBitInvalidatorUpdated is a free log retrieval operation binding the contract event 0xcda0f7e73d07bdb14b141f2cf4745926629a1b63e7c6a3dd8a80232cb459a850.
//
// Solidity: event BitInvalidatorUpdated(address indexed maker, uint256 slotIndex, uint256 slotValue)
func (_Oneinchv6 *Oneinchv6Filterer) FilterBitInvalidatorUpdated(opts *bind.FilterOpts, maker []common.Address) (*Oneinchv6BitInvalidatorUpdatedIterator, error) {

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}

	logs, sub, err := _Oneinchv6.contract.FilterLogs(opts, "BitInvalidatorUpdated", makerRule)
	if err != nil {
		return nil, err
	}
	return &Oneinchv6BitInvalidatorUpdatedIterator{contract: _Oneinchv6.contract, event: "BitInvalidatorUpdated", logs: logs, sub: sub}, nil
}

// WatchBitInvalidatorUpdated is a free log subscription operation binding the contract event 0xcda0f7e73d07bdb14b141f2cf4745926629a1b63e7c6a3dd8a80232cb459a850.
//
// Solidity: event BitInvalidatorUpdated(address indexed maker, uint256 slotIndex, uint256 slotValue)
func (_Oneinchv6 *Oneinchv6Filterer) WatchBitInvalidatorUpdated(opts *bind.WatchOpts, sink chan<- *Oneinchv6BitInvalidatorUpdated, maker []common.Address) (event.Subscription, error) {

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}

	logs, sub, err := _Oneinchv6.contract.WatchLogs(opts, "BitInvalidatorUpdated", makerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Oneinchv6BitInvalidatorUpdated)
				if err := _Oneinchv6.contract.UnpackLog(event, "BitInvalidatorUpdated", log); err != nil {
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

// ParseBitInvalidatorUpdated is a log parse operation binding the contract event 0xcda0f7e73d07bdb14b141f2cf4745926629a1b63e7c6a3dd8a80232cb459a850.
//
// Solidity: event BitInvalidatorUpdated(address indexed maker, uint256 slotIndex, uint256 slotValue)
func (_Oneinchv6 *Oneinchv6Filterer) ParseBitInvalidatorUpdated(log types.Log) (*Oneinchv6BitInvalidatorUpdated, error) {
	event := new(Oneinchv6BitInvalidatorUpdated)
	if err := _Oneinchv6.contract.UnpackLog(event, "BitInvalidatorUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Oneinchv6EIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the Oneinchv6 contract.
type Oneinchv6EIP712DomainChangedIterator struct {
	Event *Oneinchv6EIP712DomainChanged // Event containing the contract specifics and raw log

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
func (it *Oneinchv6EIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Oneinchv6EIP712DomainChanged)
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
		it.Event = new(Oneinchv6EIP712DomainChanged)
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
func (it *Oneinchv6EIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Oneinchv6EIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Oneinchv6EIP712DomainChanged represents a EIP712DomainChanged event raised by the Oneinchv6 contract.
type Oneinchv6EIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_Oneinchv6 *Oneinchv6Filterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*Oneinchv6EIP712DomainChangedIterator, error) {

	logs, sub, err := _Oneinchv6.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &Oneinchv6EIP712DomainChangedIterator{contract: _Oneinchv6.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_Oneinchv6 *Oneinchv6Filterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *Oneinchv6EIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _Oneinchv6.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Oneinchv6EIP712DomainChanged)
				if err := _Oneinchv6.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
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

// ParseEIP712DomainChanged is a log parse operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_Oneinchv6 *Oneinchv6Filterer) ParseEIP712DomainChanged(log types.Log) (*Oneinchv6EIP712DomainChanged, error) {
	event := new(Oneinchv6EIP712DomainChanged)
	if err := _Oneinchv6.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Oneinchv6EpochIncreasedIterator is returned from FilterEpochIncreased and is used to iterate over the raw logs and unpacked data for EpochIncreased events raised by the Oneinchv6 contract.
type Oneinchv6EpochIncreasedIterator struct {
	Event *Oneinchv6EpochIncreased // Event containing the contract specifics and raw log

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
func (it *Oneinchv6EpochIncreasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Oneinchv6EpochIncreased)
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
		it.Event = new(Oneinchv6EpochIncreased)
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
func (it *Oneinchv6EpochIncreasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Oneinchv6EpochIncreasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Oneinchv6EpochIncreased represents a EpochIncreased event raised by the Oneinchv6 contract.
type Oneinchv6EpochIncreased struct {
	Maker    common.Address
	Series   *big.Int
	NewEpoch *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterEpochIncreased is a free log retrieval operation binding the contract event 0x099133aefc2c2d1e56f8ef3622ec8e80979a0713fc9c4e1497740efcf8099396.
//
// Solidity: event EpochIncreased(address indexed maker, uint256 series, uint256 newEpoch)
func (_Oneinchv6 *Oneinchv6Filterer) FilterEpochIncreased(opts *bind.FilterOpts, maker []common.Address) (*Oneinchv6EpochIncreasedIterator, error) {

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}

	logs, sub, err := _Oneinchv6.contract.FilterLogs(opts, "EpochIncreased", makerRule)
	if err != nil {
		return nil, err
	}
	return &Oneinchv6EpochIncreasedIterator{contract: _Oneinchv6.contract, event: "EpochIncreased", logs: logs, sub: sub}, nil
}

// WatchEpochIncreased is a free log subscription operation binding the contract event 0x099133aefc2c2d1e56f8ef3622ec8e80979a0713fc9c4e1497740efcf8099396.
//
// Solidity: event EpochIncreased(address indexed maker, uint256 series, uint256 newEpoch)
func (_Oneinchv6 *Oneinchv6Filterer) WatchEpochIncreased(opts *bind.WatchOpts, sink chan<- *Oneinchv6EpochIncreased, maker []common.Address) (event.Subscription, error) {

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}

	logs, sub, err := _Oneinchv6.contract.WatchLogs(opts, "EpochIncreased", makerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Oneinchv6EpochIncreased)
				if err := _Oneinchv6.contract.UnpackLog(event, "EpochIncreased", log); err != nil {
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

// ParseEpochIncreased is a log parse operation binding the contract event 0x099133aefc2c2d1e56f8ef3622ec8e80979a0713fc9c4e1497740efcf8099396.
//
// Solidity: event EpochIncreased(address indexed maker, uint256 series, uint256 newEpoch)
func (_Oneinchv6 *Oneinchv6Filterer) ParseEpochIncreased(log types.Log) (*Oneinchv6EpochIncreased, error) {
	event := new(Oneinchv6EpochIncreased)
	if err := _Oneinchv6.contract.UnpackLog(event, "EpochIncreased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Oneinchv6OrderCancelledIterator is returned from FilterOrderCancelled and is used to iterate over the raw logs and unpacked data for OrderCancelled events raised by the Oneinchv6 contract.
type Oneinchv6OrderCancelledIterator struct {
	Event *Oneinchv6OrderCancelled // Event containing the contract specifics and raw log

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
func (it *Oneinchv6OrderCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Oneinchv6OrderCancelled)
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
		it.Event = new(Oneinchv6OrderCancelled)
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
func (it *Oneinchv6OrderCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Oneinchv6OrderCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Oneinchv6OrderCancelled represents a OrderCancelled event raised by the Oneinchv6 contract.
type Oneinchv6OrderCancelled struct {
	OrderHash [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOrderCancelled is a free log retrieval operation binding the contract event 0x5152abf959f6564662358c2e52b702259b78bac5ee7842a0f01937e670efcc7d.
//
// Solidity: event OrderCancelled(bytes32 orderHash)
func (_Oneinchv6 *Oneinchv6Filterer) FilterOrderCancelled(opts *bind.FilterOpts) (*Oneinchv6OrderCancelledIterator, error) {

	logs, sub, err := _Oneinchv6.contract.FilterLogs(opts, "OrderCancelled")
	if err != nil {
		return nil, err
	}
	return &Oneinchv6OrderCancelledIterator{contract: _Oneinchv6.contract, event: "OrderCancelled", logs: logs, sub: sub}, nil
}

// WatchOrderCancelled is a free log subscription operation binding the contract event 0x5152abf959f6564662358c2e52b702259b78bac5ee7842a0f01937e670efcc7d.
//
// Solidity: event OrderCancelled(bytes32 orderHash)
func (_Oneinchv6 *Oneinchv6Filterer) WatchOrderCancelled(opts *bind.WatchOpts, sink chan<- *Oneinchv6OrderCancelled) (event.Subscription, error) {

	logs, sub, err := _Oneinchv6.contract.WatchLogs(opts, "OrderCancelled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Oneinchv6OrderCancelled)
				if err := _Oneinchv6.contract.UnpackLog(event, "OrderCancelled", log); err != nil {
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

// ParseOrderCancelled is a log parse operation binding the contract event 0x5152abf959f6564662358c2e52b702259b78bac5ee7842a0f01937e670efcc7d.
//
// Solidity: event OrderCancelled(bytes32 orderHash)
func (_Oneinchv6 *Oneinchv6Filterer) ParseOrderCancelled(log types.Log) (*Oneinchv6OrderCancelled, error) {
	event := new(Oneinchv6OrderCancelled)
	if err := _Oneinchv6.contract.UnpackLog(event, "OrderCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Oneinchv6OrderFilledIterator is returned from FilterOrderFilled and is used to iterate over the raw logs and unpacked data for OrderFilled events raised by the Oneinchv6 contract.
type Oneinchv6OrderFilledIterator struct {
	Event *Oneinchv6OrderFilled // Event containing the contract specifics and raw log

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
func (it *Oneinchv6OrderFilledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Oneinchv6OrderFilled)
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
		it.Event = new(Oneinchv6OrderFilled)
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
func (it *Oneinchv6OrderFilledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Oneinchv6OrderFilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Oneinchv6OrderFilled represents a OrderFilled event raised by the Oneinchv6 contract.
type Oneinchv6OrderFilled struct {
	OrderHash       [32]byte
	RemainingAmount *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterOrderFilled is a free log retrieval operation binding the contract event 0xfec331350fce78ba658e082a71da20ac9f8d798a99b3c79681c8440cbfe77e07.
//
// Solidity: event OrderFilled(bytes32 orderHash, uint256 remainingAmount)
func (_Oneinchv6 *Oneinchv6Filterer) FilterOrderFilled(opts *bind.FilterOpts) (*Oneinchv6OrderFilledIterator, error) {

	logs, sub, err := _Oneinchv6.contract.FilterLogs(opts, "OrderFilled")
	if err != nil {
		return nil, err
	}
	return &Oneinchv6OrderFilledIterator{contract: _Oneinchv6.contract, event: "OrderFilled", logs: logs, sub: sub}, nil
}

// WatchOrderFilled is a free log subscription operation binding the contract event 0xfec331350fce78ba658e082a71da20ac9f8d798a99b3c79681c8440cbfe77e07.
//
// Solidity: event OrderFilled(bytes32 orderHash, uint256 remainingAmount)
func (_Oneinchv6 *Oneinchv6Filterer) WatchOrderFilled(opts *bind.WatchOpts, sink chan<- *Oneinchv6OrderFilled) (event.Subscription, error) {

	logs, sub, err := _Oneinchv6.contract.WatchLogs(opts, "OrderFilled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Oneinchv6OrderFilled)
				if err := _Oneinchv6.contract.UnpackLog(event, "OrderFilled", log); err != nil {
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

// ParseOrderFilled is a log parse operation binding the contract event 0xfec331350fce78ba658e082a71da20ac9f8d798a99b3c79681c8440cbfe77e07.
//
// Solidity: event OrderFilled(bytes32 orderHash, uint256 remainingAmount)
func (_Oneinchv6 *Oneinchv6Filterer) ParseOrderFilled(log types.Log) (*Oneinchv6OrderFilled, error) {
	event := new(Oneinchv6OrderFilled)
	if err := _Oneinchv6.contract.UnpackLog(event, "OrderFilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Oneinchv6OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Oneinchv6 contract.
type Oneinchv6OwnershipTransferredIterator struct {
	Event *Oneinchv6OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *Oneinchv6OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Oneinchv6OwnershipTransferred)
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
		it.Event = new(Oneinchv6OwnershipTransferred)
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
func (it *Oneinchv6OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Oneinchv6OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Oneinchv6OwnershipTransferred represents a OwnershipTransferred event raised by the Oneinchv6 contract.
type Oneinchv6OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Oneinchv6 *Oneinchv6Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*Oneinchv6OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Oneinchv6.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &Oneinchv6OwnershipTransferredIterator{contract: _Oneinchv6.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Oneinchv6 *Oneinchv6Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *Oneinchv6OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Oneinchv6.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Oneinchv6OwnershipTransferred)
				if err := _Oneinchv6.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Oneinchv6 *Oneinchv6Filterer) ParseOwnershipTransferred(log types.Log) (*Oneinchv6OwnershipTransferred, error) {
	event := new(Oneinchv6OwnershipTransferred)
	if err := _Oneinchv6.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Oneinchv6PausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Oneinchv6 contract.
type Oneinchv6PausedIterator struct {
	Event *Oneinchv6Paused // Event containing the contract specifics and raw log

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
func (it *Oneinchv6PausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Oneinchv6Paused)
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
		it.Event = new(Oneinchv6Paused)
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
func (it *Oneinchv6PausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Oneinchv6PausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Oneinchv6Paused represents a Paused event raised by the Oneinchv6 contract.
type Oneinchv6Paused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Oneinchv6 *Oneinchv6Filterer) FilterPaused(opts *bind.FilterOpts) (*Oneinchv6PausedIterator, error) {

	logs, sub, err := _Oneinchv6.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &Oneinchv6PausedIterator{contract: _Oneinchv6.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Oneinchv6 *Oneinchv6Filterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *Oneinchv6Paused) (event.Subscription, error) {

	logs, sub, err := _Oneinchv6.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Oneinchv6Paused)
				if err := _Oneinchv6.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Oneinchv6 *Oneinchv6Filterer) ParsePaused(log types.Log) (*Oneinchv6Paused, error) {
	event := new(Oneinchv6Paused)
	if err := _Oneinchv6.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Oneinchv6UnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Oneinchv6 contract.
type Oneinchv6UnpausedIterator struct {
	Event *Oneinchv6Unpaused // Event containing the contract specifics and raw log

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
func (it *Oneinchv6UnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Oneinchv6Unpaused)
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
		it.Event = new(Oneinchv6Unpaused)
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
func (it *Oneinchv6UnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Oneinchv6UnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Oneinchv6Unpaused represents a Unpaused event raised by the Oneinchv6 contract.
type Oneinchv6Unpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Oneinchv6 *Oneinchv6Filterer) FilterUnpaused(opts *bind.FilterOpts) (*Oneinchv6UnpausedIterator, error) {

	logs, sub, err := _Oneinchv6.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &Oneinchv6UnpausedIterator{contract: _Oneinchv6.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Oneinchv6 *Oneinchv6Filterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *Oneinchv6Unpaused) (event.Subscription, error) {

	logs, sub, err := _Oneinchv6.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Oneinchv6Unpaused)
				if err := _Oneinchv6.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Oneinchv6 *Oneinchv6Filterer) ParseUnpaused(log types.Log) (*Oneinchv6Unpaused, error) {
	event := new(Oneinchv6Unpaused)
	if err := _Oneinchv6.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
