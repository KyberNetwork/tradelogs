// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package oneinch

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

// OrderLibOrder is an auto generated low-level Go binding around an user-defined struct.
type OrderLibOrder struct {
	Salt          *big.Int
	MakerAsset    common.Address
	TakerAsset    common.Address
	Maker         common.Address
	Receiver      common.Address
	AllowedSender common.Address
	MakingAmount  *big.Int
	TakingAmount  *big.Int
	Offsets       *big.Int
	Interactions  []byte
}

// OrderRFQLibOrderRFQ is an auto generated low-level Go binding around an user-defined struct.
type OrderRFQLibOrderRFQ struct {
	Info          *big.Int
	MakerAsset    common.Address
	TakerAsset    common.Address
	Maker         common.Address
	AllowedSender common.Address
	MakingAmount  *big.Int
	TakingAmount  *big.Int
}

// OneinchMetaData contains all meta data concerning the Oneinch contract.
var OneinchMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIWETH\",\"name\":\"weth\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessDenied\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AdvanceNonceFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadyFilled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ArbitraryStaticCallFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadPool\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ETHTransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ETHTransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyPools\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EthDepositRejected\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GetAmountCallFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IncorrectDataLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMsgValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMsgValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidatedOrder\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MakingAmountExceeded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MakingAmountTooLow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyOneAmountShouldBeZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OrderExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PermitLengthTooLow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PredicateIsNotTrue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PrivateOrder\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RFQBadSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RFQPrivateOrder\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RFQSwapWithZeroAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RFQZeroTargetIsForbidden\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyDetected\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RemainingAmountIsZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReservesCallFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReturnAmountIsNotEnough\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SafePermitBadLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SafeTransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SafeTransferFromFailed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"res\",\"type\":\"bytes\"}],\"name\":\"SimulationResults\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SwapAmountTooLarge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SwapWithZeroAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TakingAmountExceeded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TakingAmountIncreased\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TakingAmountTooHigh\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferFromMakerToTakerFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferFromTakerToMakerFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnknownOrder\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WrongAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WrongGetter\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroMinReturn\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroReturnAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroTargetIsForbidden\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newNonce\",\"type\":\"uint256\"}],\"name\":\"NonceIncreased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"remainingRaw\",\"type\":\"uint256\"}],\"name\":\"OrderCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"remaining\",\"type\":\"uint256\"}],\"name\":\"OrderFilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"makingAmount\",\"type\":\"uint256\"}],\"name\":\"OrderFilledRFQ\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"amount\",\"type\":\"uint8\"}],\"name\":\"advanceNonce\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"offsets\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"and\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"arbitraryStaticCall\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"makerAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"allowedSender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"offsets\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"interactions\",\"type\":\"bytes\"}],\"internalType\":\"structOrderLib.Order\",\"name\":\"order\",\"type\":\"tuple\"}],\"name\":\"cancelOrder\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"orderRemaining\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"orderInfo\",\"type\":\"uint256\"}],\"name\":\"cancelOrderRFQ\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"orderInfo\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"additionalMask\",\"type\":\"uint256\"}],\"name\":\"cancelOrderRFQ\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"makerAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"allowedSender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"offsets\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"interactions\",\"type\":\"bytes\"}],\"internalType\":\"structOrderLib.Order\",\"name\":\"order\",\"type\":\"tuple\"}],\"name\":\"checkPredicate\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIClipperExchangeInterface\",\"name\":\"clipperExchange\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"dstToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"inputAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"outputAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"goodUntil\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"vs\",\"type\":\"bytes32\"}],\"name\":\"clipperSwap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIClipperExchangeInterface\",\"name\":\"clipperExchange\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"dstToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"inputAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"outputAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"goodUntil\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"vs\",\"type\":\"bytes32\"}],\"name\":\"clipperSwapTo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIClipperExchangeInterface\",\"name\":\"clipperExchange\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"dstToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"inputAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"outputAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"goodUntil\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"vs\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"permit\",\"type\":\"bytes\"}],\"name\":\"clipperSwapToWithPermit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"destroy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"eq\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"makerAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"allowedSender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"offsets\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"interactions\",\"type\":\"bytes\"}],\"internalType\":\"structOrderLib.Order\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"interaction\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"makingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"skipPermitAndThresholdAmount\",\"type\":\"uint256\"}],\"name\":\"fillOrder\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"info\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"makerAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"allowedSender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takingAmount\",\"type\":\"uint256\"}],\"internalType\":\"structOrderRFQLib.OrderRFQ\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"flagsAndAmount\",\"type\":\"uint256\"}],\"name\":\"fillOrderRFQ\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"info\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"makerAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"allowedSender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takingAmount\",\"type\":\"uint256\"}],\"internalType\":\"structOrderRFQLib.OrderRFQ\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"vs\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"flagsAndAmount\",\"type\":\"uint256\"}],\"name\":\"fillOrderRFQCompact\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"filledMakingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"filledTakingAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"info\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"makerAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"allowedSender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takingAmount\",\"type\":\"uint256\"}],\"internalType\":\"structOrderRFQLib.OrderRFQ\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"flagsAndAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"fillOrderRFQTo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"filledMakingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"filledTakingAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"info\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"makerAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"allowedSender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takingAmount\",\"type\":\"uint256\"}],\"internalType\":\"structOrderRFQLib.OrderRFQ\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"flagsAndAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"permit\",\"type\":\"bytes\"}],\"name\":\"fillOrderRFQToWithPermit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"makerAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"allowedSender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"offsets\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"interactions\",\"type\":\"bytes\"}],\"internalType\":\"structOrderLib.Order\",\"name\":\"order_\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"interaction\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"makingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"skipPermitAndThresholdAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"fillOrderTo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"actualMakingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualTakingAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"makerAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"allowedSender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"offsets\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"interactions\",\"type\":\"bytes\"}],\"internalType\":\"structOrderLib.Order\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"interaction\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"makingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"skipPermitAndThresholdAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"permit\",\"type\":\"bytes\"}],\"name\":\"fillOrderToWithPermit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"gt\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"makerAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"allowedSender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"offsets\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"interactions\",\"type\":\"bytes\"}],\"internalType\":\"structOrderLib.Order\",\"name\":\"order\",\"type\":\"tuple\"}],\"name\":\"hashOrder\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"increaseNonce\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"slot\",\"type\":\"uint256\"}],\"name\":\"invalidatorForOrderRFQ\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"lt\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"makerAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerNonce\",\"type\":\"uint256\"}],\"name\":\"nonceEquals\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"offsets\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"or\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"name\":\"remaining\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"name\":\"remainingRaw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"orderHashes\",\"type\":\"bytes32[]\"}],\"name\":\"remainingsRaw\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"rescueFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"simulate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIAggregationExecutor\",\"name\":\"executor\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"contractIERC20\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"dstToken\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"srcReceiver\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"dstReceiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minReturnAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"flags\",\"type\":\"uint256\"}],\"internalType\":\"structGenericRouter.SwapDescription\",\"name\":\"desc\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"permit\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"swap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"spentAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"timestampBelow\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timeNonceAccount\",\"type\":\"uint256\"}],\"name\":\"timestampBelowAndNonceEquals\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minReturn\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"pools\",\"type\":\"uint256[]\"}],\"name\":\"uniswapV3Swap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"amount0Delta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"amount1Delta\",\"type\":\"int256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"uniswapV3SwapCallback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minReturn\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"pools\",\"type\":\"uint256[]\"}],\"name\":\"uniswapV3SwapTo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minReturn\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"pools\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"permit\",\"type\":\"bytes\"}],\"name\":\"uniswapV3SwapToWithPermit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minReturn\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"pools\",\"type\":\"uint256[]\"}],\"name\":\"unoswap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minReturn\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"pools\",\"type\":\"uint256[]\"}],\"name\":\"unoswapTo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minReturn\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"pools\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"permit\",\"type\":\"bytes\"}],\"name\":\"unoswapToWithPermit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// OneinchABI is the input ABI used to generate the binding from.
// Deprecated: Use OneinchMetaData.ABI instead.
var OneinchABI = OneinchMetaData.ABI

// Oneinch is an auto generated Go binding around an Ethereum contract.
type Oneinch struct {
	OneinchCaller     // Read-only binding to the contract
	OneinchTransactor // Write-only binding to the contract
	OneinchFilterer   // Log filterer for contract events
}

// OneinchCaller is an auto generated read-only Go binding around an Ethereum contract.
type OneinchCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneinchTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OneinchTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneinchFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OneinchFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneinchSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OneinchSession struct {
	Contract     *Oneinch          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OneinchCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OneinchCallerSession struct {
	Contract *OneinchCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// OneinchTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OneinchTransactorSession struct {
	Contract     *OneinchTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// OneinchRaw is an auto generated low-level Go binding around an Ethereum contract.
type OneinchRaw struct {
	Contract *Oneinch // Generic contract binding to access the raw methods on
}

// OneinchCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OneinchCallerRaw struct {
	Contract *OneinchCaller // Generic read-only contract binding to access the raw methods on
}

// OneinchTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OneinchTransactorRaw struct {
	Contract *OneinchTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOneinch creates a new instance of Oneinch, bound to a specific deployed contract.
func NewOneinch(address common.Address, backend bind.ContractBackend) (*Oneinch, error) {
	contract, err := bindOneinch(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Oneinch{OneinchCaller: OneinchCaller{contract: contract}, OneinchTransactor: OneinchTransactor{contract: contract}, OneinchFilterer: OneinchFilterer{contract: contract}}, nil
}

// NewOneinchCaller creates a new read-only instance of Oneinch, bound to a specific deployed contract.
func NewOneinchCaller(address common.Address, caller bind.ContractCaller) (*OneinchCaller, error) {
	contract, err := bindOneinch(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OneinchCaller{contract: contract}, nil
}

// NewOneinchTransactor creates a new write-only instance of Oneinch, bound to a specific deployed contract.
func NewOneinchTransactor(address common.Address, transactor bind.ContractTransactor) (*OneinchTransactor, error) {
	contract, err := bindOneinch(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OneinchTransactor{contract: contract}, nil
}

// NewOneinchFilterer creates a new log filterer instance of Oneinch, bound to a specific deployed contract.
func NewOneinchFilterer(address common.Address, filterer bind.ContractFilterer) (*OneinchFilterer, error) {
	contract, err := bindOneinch(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OneinchFilterer{contract: contract}, nil
}

// bindOneinch binds a generic wrapper to an already deployed contract.
func bindOneinch(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OneinchMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Oneinch *OneinchRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Oneinch.Contract.OneinchCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Oneinch *OneinchRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Oneinch.Contract.OneinchTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Oneinch *OneinchRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Oneinch.Contract.OneinchTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Oneinch *OneinchCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Oneinch.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Oneinch *OneinchTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Oneinch.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Oneinch *OneinchTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Oneinch.Contract.contract.Transact(opts, method, params...)
}

// And is a free data retrieval call binding the contract method 0xbfa75143.
//
// Solidity: function and(uint256 offsets, bytes data) view returns(bool)
func (_Oneinch *OneinchCaller) And(opts *bind.CallOpts, offsets *big.Int, data []byte) (bool, error) {
	var out []interface{}
	err := _Oneinch.contract.Call(opts, &out, "and", offsets, data)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// And is a free data retrieval call binding the contract method 0xbfa75143.
//
// Solidity: function and(uint256 offsets, bytes data) view returns(bool)
func (_Oneinch *OneinchSession) And(offsets *big.Int, data []byte) (bool, error) {
	return _Oneinch.Contract.And(&_Oneinch.CallOpts, offsets, data)
}

// And is a free data retrieval call binding the contract method 0xbfa75143.
//
// Solidity: function and(uint256 offsets, bytes data) view returns(bool)
func (_Oneinch *OneinchCallerSession) And(offsets *big.Int, data []byte) (bool, error) {
	return _Oneinch.Contract.And(&_Oneinch.CallOpts, offsets, data)
}

// ArbitraryStaticCall is a free data retrieval call binding the contract method 0xbf15fcd8.
//
// Solidity: function arbitraryStaticCall(address target, bytes data) view returns(uint256)
func (_Oneinch *OneinchCaller) ArbitraryStaticCall(opts *bind.CallOpts, target common.Address, data []byte) (*big.Int, error) {
	var out []interface{}
	err := _Oneinch.contract.Call(opts, &out, "arbitraryStaticCall", target, data)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ArbitraryStaticCall is a free data retrieval call binding the contract method 0xbf15fcd8.
//
// Solidity: function arbitraryStaticCall(address target, bytes data) view returns(uint256)
func (_Oneinch *OneinchSession) ArbitraryStaticCall(target common.Address, data []byte) (*big.Int, error) {
	return _Oneinch.Contract.ArbitraryStaticCall(&_Oneinch.CallOpts, target, data)
}

// ArbitraryStaticCall is a free data retrieval call binding the contract method 0xbf15fcd8.
//
// Solidity: function arbitraryStaticCall(address target, bytes data) view returns(uint256)
func (_Oneinch *OneinchCallerSession) ArbitraryStaticCall(target common.Address, data []byte) (*big.Int, error) {
	return _Oneinch.Contract.ArbitraryStaticCall(&_Oneinch.CallOpts, target, data)
}

// CheckPredicate is a free data retrieval call binding the contract method 0x6c838250.
//
// Solidity: function checkPredicate((uint256,address,address,address,address,address,uint256,uint256,uint256,bytes) order) view returns(bool)
func (_Oneinch *OneinchCaller) CheckPredicate(opts *bind.CallOpts, order OrderLibOrder) (bool, error) {
	var out []interface{}
	err := _Oneinch.contract.Call(opts, &out, "checkPredicate", order)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckPredicate is a free data retrieval call binding the contract method 0x6c838250.
//
// Solidity: function checkPredicate((uint256,address,address,address,address,address,uint256,uint256,uint256,bytes) order) view returns(bool)
func (_Oneinch *OneinchSession) CheckPredicate(order OrderLibOrder) (bool, error) {
	return _Oneinch.Contract.CheckPredicate(&_Oneinch.CallOpts, order)
}

// CheckPredicate is a free data retrieval call binding the contract method 0x6c838250.
//
// Solidity: function checkPredicate((uint256,address,address,address,address,address,uint256,uint256,uint256,bytes) order) view returns(bool)
func (_Oneinch *OneinchCallerSession) CheckPredicate(order OrderLibOrder) (bool, error) {
	return _Oneinch.Contract.CheckPredicate(&_Oneinch.CallOpts, order)
}

// Eq is a free data retrieval call binding the contract method 0x6fe7b0ba.
//
// Solidity: function eq(uint256 value, bytes data) view returns(bool)
func (_Oneinch *OneinchCaller) Eq(opts *bind.CallOpts, value *big.Int, data []byte) (bool, error) {
	var out []interface{}
	err := _Oneinch.contract.Call(opts, &out, "eq", value, data)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Eq is a free data retrieval call binding the contract method 0x6fe7b0ba.
//
// Solidity: function eq(uint256 value, bytes data) view returns(bool)
func (_Oneinch *OneinchSession) Eq(value *big.Int, data []byte) (bool, error) {
	return _Oneinch.Contract.Eq(&_Oneinch.CallOpts, value, data)
}

// Eq is a free data retrieval call binding the contract method 0x6fe7b0ba.
//
// Solidity: function eq(uint256 value, bytes data) view returns(bool)
func (_Oneinch *OneinchCallerSession) Eq(value *big.Int, data []byte) (bool, error) {
	return _Oneinch.Contract.Eq(&_Oneinch.CallOpts, value, data)
}

// Gt is a free data retrieval call binding the contract method 0x4f38e2b8.
//
// Solidity: function gt(uint256 value, bytes data) view returns(bool)
func (_Oneinch *OneinchCaller) Gt(opts *bind.CallOpts, value *big.Int, data []byte) (bool, error) {
	var out []interface{}
	err := _Oneinch.contract.Call(opts, &out, "gt", value, data)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Gt is a free data retrieval call binding the contract method 0x4f38e2b8.
//
// Solidity: function gt(uint256 value, bytes data) view returns(bool)
func (_Oneinch *OneinchSession) Gt(value *big.Int, data []byte) (bool, error) {
	return _Oneinch.Contract.Gt(&_Oneinch.CallOpts, value, data)
}

// Gt is a free data retrieval call binding the contract method 0x4f38e2b8.
//
// Solidity: function gt(uint256 value, bytes data) view returns(bool)
func (_Oneinch *OneinchCallerSession) Gt(value *big.Int, data []byte) (bool, error) {
	return _Oneinch.Contract.Gt(&_Oneinch.CallOpts, value, data)
}

// HashOrder is a free data retrieval call binding the contract method 0x37e7316f.
//
// Solidity: function hashOrder((uint256,address,address,address,address,address,uint256,uint256,uint256,bytes) order) view returns(bytes32)
func (_Oneinch *OneinchCaller) HashOrder(opts *bind.CallOpts, order OrderLibOrder) ([32]byte, error) {
	var out []interface{}
	err := _Oneinch.contract.Call(opts, &out, "hashOrder", order)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HashOrder is a free data retrieval call binding the contract method 0x37e7316f.
//
// Solidity: function hashOrder((uint256,address,address,address,address,address,uint256,uint256,uint256,bytes) order) view returns(bytes32)
func (_Oneinch *OneinchSession) HashOrder(order OrderLibOrder) ([32]byte, error) {
	return _Oneinch.Contract.HashOrder(&_Oneinch.CallOpts, order)
}

// HashOrder is a free data retrieval call binding the contract method 0x37e7316f.
//
// Solidity: function hashOrder((uint256,address,address,address,address,address,uint256,uint256,uint256,bytes) order) view returns(bytes32)
func (_Oneinch *OneinchCallerSession) HashOrder(order OrderLibOrder) ([32]byte, error) {
	return _Oneinch.Contract.HashOrder(&_Oneinch.CallOpts, order)
}

// InvalidatorForOrderRFQ is a free data retrieval call binding the contract method 0x56f16124.
//
// Solidity: function invalidatorForOrderRFQ(address maker, uint256 slot) view returns(uint256)
func (_Oneinch *OneinchCaller) InvalidatorForOrderRFQ(opts *bind.CallOpts, maker common.Address, slot *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Oneinch.contract.Call(opts, &out, "invalidatorForOrderRFQ", maker, slot)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InvalidatorForOrderRFQ is a free data retrieval call binding the contract method 0x56f16124.
//
// Solidity: function invalidatorForOrderRFQ(address maker, uint256 slot) view returns(uint256)
func (_Oneinch *OneinchSession) InvalidatorForOrderRFQ(maker common.Address, slot *big.Int) (*big.Int, error) {
	return _Oneinch.Contract.InvalidatorForOrderRFQ(&_Oneinch.CallOpts, maker, slot)
}

// InvalidatorForOrderRFQ is a free data retrieval call binding the contract method 0x56f16124.
//
// Solidity: function invalidatorForOrderRFQ(address maker, uint256 slot) view returns(uint256)
func (_Oneinch *OneinchCallerSession) InvalidatorForOrderRFQ(maker common.Address, slot *big.Int) (*big.Int, error) {
	return _Oneinch.Contract.InvalidatorForOrderRFQ(&_Oneinch.CallOpts, maker, slot)
}

// Lt is a free data retrieval call binding the contract method 0xca4ece22.
//
// Solidity: function lt(uint256 value, bytes data) view returns(bool)
func (_Oneinch *OneinchCaller) Lt(opts *bind.CallOpts, value *big.Int, data []byte) (bool, error) {
	var out []interface{}
	err := _Oneinch.contract.Call(opts, &out, "lt", value, data)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Lt is a free data retrieval call binding the contract method 0xca4ece22.
//
// Solidity: function lt(uint256 value, bytes data) view returns(bool)
func (_Oneinch *OneinchSession) Lt(value *big.Int, data []byte) (bool, error) {
	return _Oneinch.Contract.Lt(&_Oneinch.CallOpts, value, data)
}

// Lt is a free data retrieval call binding the contract method 0xca4ece22.
//
// Solidity: function lt(uint256 value, bytes data) view returns(bool)
func (_Oneinch *OneinchCallerSession) Lt(value *big.Int, data []byte) (bool, error) {
	return _Oneinch.Contract.Lt(&_Oneinch.CallOpts, value, data)
}

// Nonce is a free data retrieval call binding the contract method 0x70ae92d2.
//
// Solidity: function nonce(address ) view returns(uint256)
func (_Oneinch *OneinchCaller) Nonce(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Oneinch.contract.Call(opts, &out, "nonce", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonce is a free data retrieval call binding the contract method 0x70ae92d2.
//
// Solidity: function nonce(address ) view returns(uint256)
func (_Oneinch *OneinchSession) Nonce(arg0 common.Address) (*big.Int, error) {
	return _Oneinch.Contract.Nonce(&_Oneinch.CallOpts, arg0)
}

// Nonce is a free data retrieval call binding the contract method 0x70ae92d2.
//
// Solidity: function nonce(address ) view returns(uint256)
func (_Oneinch *OneinchCallerSession) Nonce(arg0 common.Address) (*big.Int, error) {
	return _Oneinch.Contract.Nonce(&_Oneinch.CallOpts, arg0)
}

// NonceEquals is a free data retrieval call binding the contract method 0xcf6fc6e3.
//
// Solidity: function nonceEquals(address makerAddress, uint256 makerNonce) view returns(bool)
func (_Oneinch *OneinchCaller) NonceEquals(opts *bind.CallOpts, makerAddress common.Address, makerNonce *big.Int) (bool, error) {
	var out []interface{}
	err := _Oneinch.contract.Call(opts, &out, "nonceEquals", makerAddress, makerNonce)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// NonceEquals is a free data retrieval call binding the contract method 0xcf6fc6e3.
//
// Solidity: function nonceEquals(address makerAddress, uint256 makerNonce) view returns(bool)
func (_Oneinch *OneinchSession) NonceEquals(makerAddress common.Address, makerNonce *big.Int) (bool, error) {
	return _Oneinch.Contract.NonceEquals(&_Oneinch.CallOpts, makerAddress, makerNonce)
}

// NonceEquals is a free data retrieval call binding the contract method 0xcf6fc6e3.
//
// Solidity: function nonceEquals(address makerAddress, uint256 makerNonce) view returns(bool)
func (_Oneinch *OneinchCallerSession) NonceEquals(makerAddress common.Address, makerNonce *big.Int) (bool, error) {
	return _Oneinch.Contract.NonceEquals(&_Oneinch.CallOpts, makerAddress, makerNonce)
}

// Or is a free data retrieval call binding the contract method 0x74261145.
//
// Solidity: function or(uint256 offsets, bytes data) view returns(bool)
func (_Oneinch *OneinchCaller) Or(opts *bind.CallOpts, offsets *big.Int, data []byte) (bool, error) {
	var out []interface{}
	err := _Oneinch.contract.Call(opts, &out, "or", offsets, data)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Or is a free data retrieval call binding the contract method 0x74261145.
//
// Solidity: function or(uint256 offsets, bytes data) view returns(bool)
func (_Oneinch *OneinchSession) Or(offsets *big.Int, data []byte) (bool, error) {
	return _Oneinch.Contract.Or(&_Oneinch.CallOpts, offsets, data)
}

// Or is a free data retrieval call binding the contract method 0x74261145.
//
// Solidity: function or(uint256 offsets, bytes data) view returns(bool)
func (_Oneinch *OneinchCallerSession) Or(offsets *big.Int, data []byte) (bool, error) {
	return _Oneinch.Contract.Or(&_Oneinch.CallOpts, offsets, data)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Oneinch *OneinchCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Oneinch.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Oneinch *OneinchSession) Owner() (common.Address, error) {
	return _Oneinch.Contract.Owner(&_Oneinch.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Oneinch *OneinchCallerSession) Owner() (common.Address, error) {
	return _Oneinch.Contract.Owner(&_Oneinch.CallOpts)
}

// Remaining is a free data retrieval call binding the contract method 0xbc1ed74c.
//
// Solidity: function remaining(bytes32 orderHash) view returns(uint256)
func (_Oneinch *OneinchCaller) Remaining(opts *bind.CallOpts, orderHash [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Oneinch.contract.Call(opts, &out, "remaining", orderHash)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Remaining is a free data retrieval call binding the contract method 0xbc1ed74c.
//
// Solidity: function remaining(bytes32 orderHash) view returns(uint256)
func (_Oneinch *OneinchSession) Remaining(orderHash [32]byte) (*big.Int, error) {
	return _Oneinch.Contract.Remaining(&_Oneinch.CallOpts, orderHash)
}

// Remaining is a free data retrieval call binding the contract method 0xbc1ed74c.
//
// Solidity: function remaining(bytes32 orderHash) view returns(uint256)
func (_Oneinch *OneinchCallerSession) Remaining(orderHash [32]byte) (*big.Int, error) {
	return _Oneinch.Contract.Remaining(&_Oneinch.CallOpts, orderHash)
}

// RemainingRaw is a free data retrieval call binding the contract method 0x7e54f092.
//
// Solidity: function remainingRaw(bytes32 orderHash) view returns(uint256)
func (_Oneinch *OneinchCaller) RemainingRaw(opts *bind.CallOpts, orderHash [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Oneinch.contract.Call(opts, &out, "remainingRaw", orderHash)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RemainingRaw is a free data retrieval call binding the contract method 0x7e54f092.
//
// Solidity: function remainingRaw(bytes32 orderHash) view returns(uint256)
func (_Oneinch *OneinchSession) RemainingRaw(orderHash [32]byte) (*big.Int, error) {
	return _Oneinch.Contract.RemainingRaw(&_Oneinch.CallOpts, orderHash)
}

// RemainingRaw is a free data retrieval call binding the contract method 0x7e54f092.
//
// Solidity: function remainingRaw(bytes32 orderHash) view returns(uint256)
func (_Oneinch *OneinchCallerSession) RemainingRaw(orderHash [32]byte) (*big.Int, error) {
	return _Oneinch.Contract.RemainingRaw(&_Oneinch.CallOpts, orderHash)
}

// RemainingsRaw is a free data retrieval call binding the contract method 0x942461bb.
//
// Solidity: function remainingsRaw(bytes32[] orderHashes) view returns(uint256[])
func (_Oneinch *OneinchCaller) RemainingsRaw(opts *bind.CallOpts, orderHashes [][32]byte) ([]*big.Int, error) {
	var out []interface{}
	err := _Oneinch.contract.Call(opts, &out, "remainingsRaw", orderHashes)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// RemainingsRaw is a free data retrieval call binding the contract method 0x942461bb.
//
// Solidity: function remainingsRaw(bytes32[] orderHashes) view returns(uint256[])
func (_Oneinch *OneinchSession) RemainingsRaw(orderHashes [][32]byte) ([]*big.Int, error) {
	return _Oneinch.Contract.RemainingsRaw(&_Oneinch.CallOpts, orderHashes)
}

// RemainingsRaw is a free data retrieval call binding the contract method 0x942461bb.
//
// Solidity: function remainingsRaw(bytes32[] orderHashes) view returns(uint256[])
func (_Oneinch *OneinchCallerSession) RemainingsRaw(orderHashes [][32]byte) ([]*big.Int, error) {
	return _Oneinch.Contract.RemainingsRaw(&_Oneinch.CallOpts, orderHashes)
}

// TimestampBelow is a free data retrieval call binding the contract method 0x63592c2b.
//
// Solidity: function timestampBelow(uint256 time) view returns(bool)
func (_Oneinch *OneinchCaller) TimestampBelow(opts *bind.CallOpts, time *big.Int) (bool, error) {
	var out []interface{}
	err := _Oneinch.contract.Call(opts, &out, "timestampBelow", time)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TimestampBelow is a free data retrieval call binding the contract method 0x63592c2b.
//
// Solidity: function timestampBelow(uint256 time) view returns(bool)
func (_Oneinch *OneinchSession) TimestampBelow(time *big.Int) (bool, error) {
	return _Oneinch.Contract.TimestampBelow(&_Oneinch.CallOpts, time)
}

// TimestampBelow is a free data retrieval call binding the contract method 0x63592c2b.
//
// Solidity: function timestampBelow(uint256 time) view returns(bool)
func (_Oneinch *OneinchCallerSession) TimestampBelow(time *big.Int) (bool, error) {
	return _Oneinch.Contract.TimestampBelow(&_Oneinch.CallOpts, time)
}

// TimestampBelowAndNonceEquals is a free data retrieval call binding the contract method 0x2cc2878d.
//
// Solidity: function timestampBelowAndNonceEquals(uint256 timeNonceAccount) view returns(bool)
func (_Oneinch *OneinchCaller) TimestampBelowAndNonceEquals(opts *bind.CallOpts, timeNonceAccount *big.Int) (bool, error) {
	var out []interface{}
	err := _Oneinch.contract.Call(opts, &out, "timestampBelowAndNonceEquals", timeNonceAccount)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TimestampBelowAndNonceEquals is a free data retrieval call binding the contract method 0x2cc2878d.
//
// Solidity: function timestampBelowAndNonceEquals(uint256 timeNonceAccount) view returns(bool)
func (_Oneinch *OneinchSession) TimestampBelowAndNonceEquals(timeNonceAccount *big.Int) (bool, error) {
	return _Oneinch.Contract.TimestampBelowAndNonceEquals(&_Oneinch.CallOpts, timeNonceAccount)
}

// TimestampBelowAndNonceEquals is a free data retrieval call binding the contract method 0x2cc2878d.
//
// Solidity: function timestampBelowAndNonceEquals(uint256 timeNonceAccount) view returns(bool)
func (_Oneinch *OneinchCallerSession) TimestampBelowAndNonceEquals(timeNonceAccount *big.Int) (bool, error) {
	return _Oneinch.Contract.TimestampBelowAndNonceEquals(&_Oneinch.CallOpts, timeNonceAccount)
}

// AdvanceNonce is a paid mutator transaction binding the contract method 0x72c244a8.
//
// Solidity: function advanceNonce(uint8 amount) returns()
func (_Oneinch *OneinchTransactor) AdvanceNonce(opts *bind.TransactOpts, amount uint8) (*types.Transaction, error) {
	return _Oneinch.contract.Transact(opts, "advanceNonce", amount)
}

// AdvanceNonce is a paid mutator transaction binding the contract method 0x72c244a8.
//
// Solidity: function advanceNonce(uint8 amount) returns()
func (_Oneinch *OneinchSession) AdvanceNonce(amount uint8) (*types.Transaction, error) {
	return _Oneinch.Contract.AdvanceNonce(&_Oneinch.TransactOpts, amount)
}

// AdvanceNonce is a paid mutator transaction binding the contract method 0x72c244a8.
//
// Solidity: function advanceNonce(uint8 amount) returns()
func (_Oneinch *OneinchTransactorSession) AdvanceNonce(amount uint8) (*types.Transaction, error) {
	return _Oneinch.Contract.AdvanceNonce(&_Oneinch.TransactOpts, amount)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x2d9a56f6.
//
// Solidity: function cancelOrder((uint256,address,address,address,address,address,uint256,uint256,uint256,bytes) order) returns(uint256 orderRemaining, bytes32 orderHash)
func (_Oneinch *OneinchTransactor) CancelOrder(opts *bind.TransactOpts, order OrderLibOrder) (*types.Transaction, error) {
	return _Oneinch.contract.Transact(opts, "cancelOrder", order)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x2d9a56f6.
//
// Solidity: function cancelOrder((uint256,address,address,address,address,address,uint256,uint256,uint256,bytes) order) returns(uint256 orderRemaining, bytes32 orderHash)
func (_Oneinch *OneinchSession) CancelOrder(order OrderLibOrder) (*types.Transaction, error) {
	return _Oneinch.Contract.CancelOrder(&_Oneinch.TransactOpts, order)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x2d9a56f6.
//
// Solidity: function cancelOrder((uint256,address,address,address,address,address,uint256,uint256,uint256,bytes) order) returns(uint256 orderRemaining, bytes32 orderHash)
func (_Oneinch *OneinchTransactorSession) CancelOrder(order OrderLibOrder) (*types.Transaction, error) {
	return _Oneinch.Contract.CancelOrder(&_Oneinch.TransactOpts, order)
}

// CancelOrderRFQ is a paid mutator transaction binding the contract method 0x825caba1.
//
// Solidity: function cancelOrderRFQ(uint256 orderInfo) returns()
func (_Oneinch *OneinchTransactor) CancelOrderRFQ(opts *bind.TransactOpts, orderInfo *big.Int) (*types.Transaction, error) {
	return _Oneinch.contract.Transact(opts, "cancelOrderRFQ", orderInfo)
}

// CancelOrderRFQ is a paid mutator transaction binding the contract method 0x825caba1.
//
// Solidity: function cancelOrderRFQ(uint256 orderInfo) returns()
func (_Oneinch *OneinchSession) CancelOrderRFQ(orderInfo *big.Int) (*types.Transaction, error) {
	return _Oneinch.Contract.CancelOrderRFQ(&_Oneinch.TransactOpts, orderInfo)
}

// CancelOrderRFQ is a paid mutator transaction binding the contract method 0x825caba1.
//
// Solidity: function cancelOrderRFQ(uint256 orderInfo) returns()
func (_Oneinch *OneinchTransactorSession) CancelOrderRFQ(orderInfo *big.Int) (*types.Transaction, error) {
	return _Oneinch.Contract.CancelOrderRFQ(&_Oneinch.TransactOpts, orderInfo)
}

// CancelOrderRFQ0 is a paid mutator transaction binding the contract method 0xbddccd35.
//
// Solidity: function cancelOrderRFQ(uint256 orderInfo, uint256 additionalMask) returns()
func (_Oneinch *OneinchTransactor) CancelOrderRFQ0(opts *bind.TransactOpts, orderInfo *big.Int, additionalMask *big.Int) (*types.Transaction, error) {
	return _Oneinch.contract.Transact(opts, "cancelOrderRFQ0", orderInfo, additionalMask)
}

// CancelOrderRFQ0 is a paid mutator transaction binding the contract method 0xbddccd35.
//
// Solidity: function cancelOrderRFQ(uint256 orderInfo, uint256 additionalMask) returns()
func (_Oneinch *OneinchSession) CancelOrderRFQ0(orderInfo *big.Int, additionalMask *big.Int) (*types.Transaction, error) {
	return _Oneinch.Contract.CancelOrderRFQ0(&_Oneinch.TransactOpts, orderInfo, additionalMask)
}

// CancelOrderRFQ0 is a paid mutator transaction binding the contract method 0xbddccd35.
//
// Solidity: function cancelOrderRFQ(uint256 orderInfo, uint256 additionalMask) returns()
func (_Oneinch *OneinchTransactorSession) CancelOrderRFQ0(orderInfo *big.Int, additionalMask *big.Int) (*types.Transaction, error) {
	return _Oneinch.Contract.CancelOrderRFQ0(&_Oneinch.TransactOpts, orderInfo, additionalMask)
}

// ClipperSwap is a paid mutator transaction binding the contract method 0x84bd6d29.
//
// Solidity: function clipperSwap(address clipperExchange, address srcToken, address dstToken, uint256 inputAmount, uint256 outputAmount, uint256 goodUntil, bytes32 r, bytes32 vs) payable returns(uint256 returnAmount)
func (_Oneinch *OneinchTransactor) ClipperSwap(opts *bind.TransactOpts, clipperExchange common.Address, srcToken common.Address, dstToken common.Address, inputAmount *big.Int, outputAmount *big.Int, goodUntil *big.Int, r [32]byte, vs [32]byte) (*types.Transaction, error) {
	return _Oneinch.contract.Transact(opts, "clipperSwap", clipperExchange, srcToken, dstToken, inputAmount, outputAmount, goodUntil, r, vs)
}

// ClipperSwap is a paid mutator transaction binding the contract method 0x84bd6d29.
//
// Solidity: function clipperSwap(address clipperExchange, address srcToken, address dstToken, uint256 inputAmount, uint256 outputAmount, uint256 goodUntil, bytes32 r, bytes32 vs) payable returns(uint256 returnAmount)
func (_Oneinch *OneinchSession) ClipperSwap(clipperExchange common.Address, srcToken common.Address, dstToken common.Address, inputAmount *big.Int, outputAmount *big.Int, goodUntil *big.Int, r [32]byte, vs [32]byte) (*types.Transaction, error) {
	return _Oneinch.Contract.ClipperSwap(&_Oneinch.TransactOpts, clipperExchange, srcToken, dstToken, inputAmount, outputAmount, goodUntil, r, vs)
}

// ClipperSwap is a paid mutator transaction binding the contract method 0x84bd6d29.
//
// Solidity: function clipperSwap(address clipperExchange, address srcToken, address dstToken, uint256 inputAmount, uint256 outputAmount, uint256 goodUntil, bytes32 r, bytes32 vs) payable returns(uint256 returnAmount)
func (_Oneinch *OneinchTransactorSession) ClipperSwap(clipperExchange common.Address, srcToken common.Address, dstToken common.Address, inputAmount *big.Int, outputAmount *big.Int, goodUntil *big.Int, r [32]byte, vs [32]byte) (*types.Transaction, error) {
	return _Oneinch.Contract.ClipperSwap(&_Oneinch.TransactOpts, clipperExchange, srcToken, dstToken, inputAmount, outputAmount, goodUntil, r, vs)
}

// ClipperSwapTo is a paid mutator transaction binding the contract method 0x093d4fa5.
//
// Solidity: function clipperSwapTo(address clipperExchange, address recipient, address srcToken, address dstToken, uint256 inputAmount, uint256 outputAmount, uint256 goodUntil, bytes32 r, bytes32 vs) payable returns(uint256 returnAmount)
func (_Oneinch *OneinchTransactor) ClipperSwapTo(opts *bind.TransactOpts, clipperExchange common.Address, recipient common.Address, srcToken common.Address, dstToken common.Address, inputAmount *big.Int, outputAmount *big.Int, goodUntil *big.Int, r [32]byte, vs [32]byte) (*types.Transaction, error) {
	return _Oneinch.contract.Transact(opts, "clipperSwapTo", clipperExchange, recipient, srcToken, dstToken, inputAmount, outputAmount, goodUntil, r, vs)
}

// ClipperSwapTo is a paid mutator transaction binding the contract method 0x093d4fa5.
//
// Solidity: function clipperSwapTo(address clipperExchange, address recipient, address srcToken, address dstToken, uint256 inputAmount, uint256 outputAmount, uint256 goodUntil, bytes32 r, bytes32 vs) payable returns(uint256 returnAmount)
func (_Oneinch *OneinchSession) ClipperSwapTo(clipperExchange common.Address, recipient common.Address, srcToken common.Address, dstToken common.Address, inputAmount *big.Int, outputAmount *big.Int, goodUntil *big.Int, r [32]byte, vs [32]byte) (*types.Transaction, error) {
	return _Oneinch.Contract.ClipperSwapTo(&_Oneinch.TransactOpts, clipperExchange, recipient, srcToken, dstToken, inputAmount, outputAmount, goodUntil, r, vs)
}

// ClipperSwapTo is a paid mutator transaction binding the contract method 0x093d4fa5.
//
// Solidity: function clipperSwapTo(address clipperExchange, address recipient, address srcToken, address dstToken, uint256 inputAmount, uint256 outputAmount, uint256 goodUntil, bytes32 r, bytes32 vs) payable returns(uint256 returnAmount)
func (_Oneinch *OneinchTransactorSession) ClipperSwapTo(clipperExchange common.Address, recipient common.Address, srcToken common.Address, dstToken common.Address, inputAmount *big.Int, outputAmount *big.Int, goodUntil *big.Int, r [32]byte, vs [32]byte) (*types.Transaction, error) {
	return _Oneinch.Contract.ClipperSwapTo(&_Oneinch.TransactOpts, clipperExchange, recipient, srcToken, dstToken, inputAmount, outputAmount, goodUntil, r, vs)
}

// ClipperSwapToWithPermit is a paid mutator transaction binding the contract method 0xc805a666.
//
// Solidity: function clipperSwapToWithPermit(address clipperExchange, address recipient, address srcToken, address dstToken, uint256 inputAmount, uint256 outputAmount, uint256 goodUntil, bytes32 r, bytes32 vs, bytes permit) returns(uint256 returnAmount)
func (_Oneinch *OneinchTransactor) ClipperSwapToWithPermit(opts *bind.TransactOpts, clipperExchange common.Address, recipient common.Address, srcToken common.Address, dstToken common.Address, inputAmount *big.Int, outputAmount *big.Int, goodUntil *big.Int, r [32]byte, vs [32]byte, permit []byte) (*types.Transaction, error) {
	return _Oneinch.contract.Transact(opts, "clipperSwapToWithPermit", clipperExchange, recipient, srcToken, dstToken, inputAmount, outputAmount, goodUntil, r, vs, permit)
}

// ClipperSwapToWithPermit is a paid mutator transaction binding the contract method 0xc805a666.
//
// Solidity: function clipperSwapToWithPermit(address clipperExchange, address recipient, address srcToken, address dstToken, uint256 inputAmount, uint256 outputAmount, uint256 goodUntil, bytes32 r, bytes32 vs, bytes permit) returns(uint256 returnAmount)
func (_Oneinch *OneinchSession) ClipperSwapToWithPermit(clipperExchange common.Address, recipient common.Address, srcToken common.Address, dstToken common.Address, inputAmount *big.Int, outputAmount *big.Int, goodUntil *big.Int, r [32]byte, vs [32]byte, permit []byte) (*types.Transaction, error) {
	return _Oneinch.Contract.ClipperSwapToWithPermit(&_Oneinch.TransactOpts, clipperExchange, recipient, srcToken, dstToken, inputAmount, outputAmount, goodUntil, r, vs, permit)
}

// ClipperSwapToWithPermit is a paid mutator transaction binding the contract method 0xc805a666.
//
// Solidity: function clipperSwapToWithPermit(address clipperExchange, address recipient, address srcToken, address dstToken, uint256 inputAmount, uint256 outputAmount, uint256 goodUntil, bytes32 r, bytes32 vs, bytes permit) returns(uint256 returnAmount)
func (_Oneinch *OneinchTransactorSession) ClipperSwapToWithPermit(clipperExchange common.Address, recipient common.Address, srcToken common.Address, dstToken common.Address, inputAmount *big.Int, outputAmount *big.Int, goodUntil *big.Int, r [32]byte, vs [32]byte, permit []byte) (*types.Transaction, error) {
	return _Oneinch.Contract.ClipperSwapToWithPermit(&_Oneinch.TransactOpts, clipperExchange, recipient, srcToken, dstToken, inputAmount, outputAmount, goodUntil, r, vs, permit)
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_Oneinch *OneinchTransactor) Destroy(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Oneinch.contract.Transact(opts, "destroy")
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_Oneinch *OneinchSession) Destroy() (*types.Transaction, error) {
	return _Oneinch.Contract.Destroy(&_Oneinch.TransactOpts)
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_Oneinch *OneinchTransactorSession) Destroy() (*types.Transaction, error) {
	return _Oneinch.Contract.Destroy(&_Oneinch.TransactOpts)
}

// FillOrder is a paid mutator transaction binding the contract method 0x62e238bb.
//
// Solidity: function fillOrder((uint256,address,address,address,address,address,uint256,uint256,uint256,bytes) order, bytes signature, bytes interaction, uint256 makingAmount, uint256 takingAmount, uint256 skipPermitAndThresholdAmount) payable returns(uint256, uint256, bytes32)
func (_Oneinch *OneinchTransactor) FillOrder(opts *bind.TransactOpts, order OrderLibOrder, signature []byte, interaction []byte, makingAmount *big.Int, takingAmount *big.Int, skipPermitAndThresholdAmount *big.Int) (*types.Transaction, error) {
	return _Oneinch.contract.Transact(opts, "fillOrder", order, signature, interaction, makingAmount, takingAmount, skipPermitAndThresholdAmount)
}

// FillOrder is a paid mutator transaction binding the contract method 0x62e238bb.
//
// Solidity: function fillOrder((uint256,address,address,address,address,address,uint256,uint256,uint256,bytes) order, bytes signature, bytes interaction, uint256 makingAmount, uint256 takingAmount, uint256 skipPermitAndThresholdAmount) payable returns(uint256, uint256, bytes32)
func (_Oneinch *OneinchSession) FillOrder(order OrderLibOrder, signature []byte, interaction []byte, makingAmount *big.Int, takingAmount *big.Int, skipPermitAndThresholdAmount *big.Int) (*types.Transaction, error) {
	return _Oneinch.Contract.FillOrder(&_Oneinch.TransactOpts, order, signature, interaction, makingAmount, takingAmount, skipPermitAndThresholdAmount)
}

// FillOrder is a paid mutator transaction binding the contract method 0x62e238bb.
//
// Solidity: function fillOrder((uint256,address,address,address,address,address,uint256,uint256,uint256,bytes) order, bytes signature, bytes interaction, uint256 makingAmount, uint256 takingAmount, uint256 skipPermitAndThresholdAmount) payable returns(uint256, uint256, bytes32)
func (_Oneinch *OneinchTransactorSession) FillOrder(order OrderLibOrder, signature []byte, interaction []byte, makingAmount *big.Int, takingAmount *big.Int, skipPermitAndThresholdAmount *big.Int) (*types.Transaction, error) {
	return _Oneinch.Contract.FillOrder(&_Oneinch.TransactOpts, order, signature, interaction, makingAmount, takingAmount, skipPermitAndThresholdAmount)
}

// FillOrderRFQ is a paid mutator transaction binding the contract method 0x3eca9c0a.
//
// Solidity: function fillOrderRFQ((uint256,address,address,address,address,uint256,uint256) order, bytes signature, uint256 flagsAndAmount) payable returns(uint256, uint256, bytes32)
func (_Oneinch *OneinchTransactor) FillOrderRFQ(opts *bind.TransactOpts, order OrderRFQLibOrderRFQ, signature []byte, flagsAndAmount *big.Int) (*types.Transaction, error) {
	return _Oneinch.contract.Transact(opts, "fillOrderRFQ", order, signature, flagsAndAmount)
}

// FillOrderRFQ is a paid mutator transaction binding the contract method 0x3eca9c0a.
//
// Solidity: function fillOrderRFQ((uint256,address,address,address,address,uint256,uint256) order, bytes signature, uint256 flagsAndAmount) payable returns(uint256, uint256, bytes32)
func (_Oneinch *OneinchSession) FillOrderRFQ(order OrderRFQLibOrderRFQ, signature []byte, flagsAndAmount *big.Int) (*types.Transaction, error) {
	return _Oneinch.Contract.FillOrderRFQ(&_Oneinch.TransactOpts, order, signature, flagsAndAmount)
}

// FillOrderRFQ is a paid mutator transaction binding the contract method 0x3eca9c0a.
//
// Solidity: function fillOrderRFQ((uint256,address,address,address,address,uint256,uint256) order, bytes signature, uint256 flagsAndAmount) payable returns(uint256, uint256, bytes32)
func (_Oneinch *OneinchTransactorSession) FillOrderRFQ(order OrderRFQLibOrderRFQ, signature []byte, flagsAndAmount *big.Int) (*types.Transaction, error) {
	return _Oneinch.Contract.FillOrderRFQ(&_Oneinch.TransactOpts, order, signature, flagsAndAmount)
}

// FillOrderRFQCompact is a paid mutator transaction binding the contract method 0x9570eeee.
//
// Solidity: function fillOrderRFQCompact((uint256,address,address,address,address,uint256,uint256) order, bytes32 r, bytes32 vs, uint256 flagsAndAmount) payable returns(uint256 filledMakingAmount, uint256 filledTakingAmount, bytes32 orderHash)
func (_Oneinch *OneinchTransactor) FillOrderRFQCompact(opts *bind.TransactOpts, order OrderRFQLibOrderRFQ, r [32]byte, vs [32]byte, flagsAndAmount *big.Int) (*types.Transaction, error) {
	return _Oneinch.contract.Transact(opts, "fillOrderRFQCompact", order, r, vs, flagsAndAmount)
}

// FillOrderRFQCompact is a paid mutator transaction binding the contract method 0x9570eeee.
//
// Solidity: function fillOrderRFQCompact((uint256,address,address,address,address,uint256,uint256) order, bytes32 r, bytes32 vs, uint256 flagsAndAmount) payable returns(uint256 filledMakingAmount, uint256 filledTakingAmount, bytes32 orderHash)
func (_Oneinch *OneinchSession) FillOrderRFQCompact(order OrderRFQLibOrderRFQ, r [32]byte, vs [32]byte, flagsAndAmount *big.Int) (*types.Transaction, error) {
	return _Oneinch.Contract.FillOrderRFQCompact(&_Oneinch.TransactOpts, order, r, vs, flagsAndAmount)
}

// FillOrderRFQCompact is a paid mutator transaction binding the contract method 0x9570eeee.
//
// Solidity: function fillOrderRFQCompact((uint256,address,address,address,address,uint256,uint256) order, bytes32 r, bytes32 vs, uint256 flagsAndAmount) payable returns(uint256 filledMakingAmount, uint256 filledTakingAmount, bytes32 orderHash)
func (_Oneinch *OneinchTransactorSession) FillOrderRFQCompact(order OrderRFQLibOrderRFQ, r [32]byte, vs [32]byte, flagsAndAmount *big.Int) (*types.Transaction, error) {
	return _Oneinch.Contract.FillOrderRFQCompact(&_Oneinch.TransactOpts, order, r, vs, flagsAndAmount)
}

// FillOrderRFQTo is a paid mutator transaction binding the contract method 0x5a099843.
//
// Solidity: function fillOrderRFQTo((uint256,address,address,address,address,uint256,uint256) order, bytes signature, uint256 flagsAndAmount, address target) payable returns(uint256 filledMakingAmount, uint256 filledTakingAmount, bytes32 orderHash)
func (_Oneinch *OneinchTransactor) FillOrderRFQTo(opts *bind.TransactOpts, order OrderRFQLibOrderRFQ, signature []byte, flagsAndAmount *big.Int, target common.Address) (*types.Transaction, error) {
	return _Oneinch.contract.Transact(opts, "fillOrderRFQTo", order, signature, flagsAndAmount, target)
}

// FillOrderRFQTo is a paid mutator transaction binding the contract method 0x5a099843.
//
// Solidity: function fillOrderRFQTo((uint256,address,address,address,address,uint256,uint256) order, bytes signature, uint256 flagsAndAmount, address target) payable returns(uint256 filledMakingAmount, uint256 filledTakingAmount, bytes32 orderHash)
func (_Oneinch *OneinchSession) FillOrderRFQTo(order OrderRFQLibOrderRFQ, signature []byte, flagsAndAmount *big.Int, target common.Address) (*types.Transaction, error) {
	return _Oneinch.Contract.FillOrderRFQTo(&_Oneinch.TransactOpts, order, signature, flagsAndAmount, target)
}

// FillOrderRFQTo is a paid mutator transaction binding the contract method 0x5a099843.
//
// Solidity: function fillOrderRFQTo((uint256,address,address,address,address,uint256,uint256) order, bytes signature, uint256 flagsAndAmount, address target) payable returns(uint256 filledMakingAmount, uint256 filledTakingAmount, bytes32 orderHash)
func (_Oneinch *OneinchTransactorSession) FillOrderRFQTo(order OrderRFQLibOrderRFQ, signature []byte, flagsAndAmount *big.Int, target common.Address) (*types.Transaction, error) {
	return _Oneinch.Contract.FillOrderRFQTo(&_Oneinch.TransactOpts, order, signature, flagsAndAmount, target)
}

// FillOrderRFQToWithPermit is a paid mutator transaction binding the contract method 0x70ccbd31.
//
// Solidity: function fillOrderRFQToWithPermit((uint256,address,address,address,address,uint256,uint256) order, bytes signature, uint256 flagsAndAmount, address target, bytes permit) returns(uint256, uint256, bytes32)
func (_Oneinch *OneinchTransactor) FillOrderRFQToWithPermit(opts *bind.TransactOpts, order OrderRFQLibOrderRFQ, signature []byte, flagsAndAmount *big.Int, target common.Address, permit []byte) (*types.Transaction, error) {
	return _Oneinch.contract.Transact(opts, "fillOrderRFQToWithPermit", order, signature, flagsAndAmount, target, permit)
}

// FillOrderRFQToWithPermit is a paid mutator transaction binding the contract method 0x70ccbd31.
//
// Solidity: function fillOrderRFQToWithPermit((uint256,address,address,address,address,uint256,uint256) order, bytes signature, uint256 flagsAndAmount, address target, bytes permit) returns(uint256, uint256, bytes32)
func (_Oneinch *OneinchSession) FillOrderRFQToWithPermit(order OrderRFQLibOrderRFQ, signature []byte, flagsAndAmount *big.Int, target common.Address, permit []byte) (*types.Transaction, error) {
	return _Oneinch.Contract.FillOrderRFQToWithPermit(&_Oneinch.TransactOpts, order, signature, flagsAndAmount, target, permit)
}

// FillOrderRFQToWithPermit is a paid mutator transaction binding the contract method 0x70ccbd31.
//
// Solidity: function fillOrderRFQToWithPermit((uint256,address,address,address,address,uint256,uint256) order, bytes signature, uint256 flagsAndAmount, address target, bytes permit) returns(uint256, uint256, bytes32)
func (_Oneinch *OneinchTransactorSession) FillOrderRFQToWithPermit(order OrderRFQLibOrderRFQ, signature []byte, flagsAndAmount *big.Int, target common.Address, permit []byte) (*types.Transaction, error) {
	return _Oneinch.Contract.FillOrderRFQToWithPermit(&_Oneinch.TransactOpts, order, signature, flagsAndAmount, target, permit)
}

// FillOrderTo is a paid mutator transaction binding the contract method 0xe5d7bde6.
//
// Solidity: function fillOrderTo((uint256,address,address,address,address,address,uint256,uint256,uint256,bytes) order_, bytes signature, bytes interaction, uint256 makingAmount, uint256 takingAmount, uint256 skipPermitAndThresholdAmount, address target) payable returns(uint256 actualMakingAmount, uint256 actualTakingAmount, bytes32 orderHash)
func (_Oneinch *OneinchTransactor) FillOrderTo(opts *bind.TransactOpts, order_ OrderLibOrder, signature []byte, interaction []byte, makingAmount *big.Int, takingAmount *big.Int, skipPermitAndThresholdAmount *big.Int, target common.Address) (*types.Transaction, error) {
	return _Oneinch.contract.Transact(opts, "fillOrderTo", order_, signature, interaction, makingAmount, takingAmount, skipPermitAndThresholdAmount, target)
}

// FillOrderTo is a paid mutator transaction binding the contract method 0xe5d7bde6.
//
// Solidity: function fillOrderTo((uint256,address,address,address,address,address,uint256,uint256,uint256,bytes) order_, bytes signature, bytes interaction, uint256 makingAmount, uint256 takingAmount, uint256 skipPermitAndThresholdAmount, address target) payable returns(uint256 actualMakingAmount, uint256 actualTakingAmount, bytes32 orderHash)
func (_Oneinch *OneinchSession) FillOrderTo(order_ OrderLibOrder, signature []byte, interaction []byte, makingAmount *big.Int, takingAmount *big.Int, skipPermitAndThresholdAmount *big.Int, target common.Address) (*types.Transaction, error) {
	return _Oneinch.Contract.FillOrderTo(&_Oneinch.TransactOpts, order_, signature, interaction, makingAmount, takingAmount, skipPermitAndThresholdAmount, target)
}

// FillOrderTo is a paid mutator transaction binding the contract method 0xe5d7bde6.
//
// Solidity: function fillOrderTo((uint256,address,address,address,address,address,uint256,uint256,uint256,bytes) order_, bytes signature, bytes interaction, uint256 makingAmount, uint256 takingAmount, uint256 skipPermitAndThresholdAmount, address target) payable returns(uint256 actualMakingAmount, uint256 actualTakingAmount, bytes32 orderHash)
func (_Oneinch *OneinchTransactorSession) FillOrderTo(order_ OrderLibOrder, signature []byte, interaction []byte, makingAmount *big.Int, takingAmount *big.Int, skipPermitAndThresholdAmount *big.Int, target common.Address) (*types.Transaction, error) {
	return _Oneinch.Contract.FillOrderTo(&_Oneinch.TransactOpts, order_, signature, interaction, makingAmount, takingAmount, skipPermitAndThresholdAmount, target)
}

// FillOrderToWithPermit is a paid mutator transaction binding the contract method 0xd365c695.
//
// Solidity: function fillOrderToWithPermit((uint256,address,address,address,address,address,uint256,uint256,uint256,bytes) order, bytes signature, bytes interaction, uint256 makingAmount, uint256 takingAmount, uint256 skipPermitAndThresholdAmount, address target, bytes permit) returns(uint256, uint256, bytes32)
func (_Oneinch *OneinchTransactor) FillOrderToWithPermit(opts *bind.TransactOpts, order OrderLibOrder, signature []byte, interaction []byte, makingAmount *big.Int, takingAmount *big.Int, skipPermitAndThresholdAmount *big.Int, target common.Address, permit []byte) (*types.Transaction, error) {
	return _Oneinch.contract.Transact(opts, "fillOrderToWithPermit", order, signature, interaction, makingAmount, takingAmount, skipPermitAndThresholdAmount, target, permit)
}

// FillOrderToWithPermit is a paid mutator transaction binding the contract method 0xd365c695.
//
// Solidity: function fillOrderToWithPermit((uint256,address,address,address,address,address,uint256,uint256,uint256,bytes) order, bytes signature, bytes interaction, uint256 makingAmount, uint256 takingAmount, uint256 skipPermitAndThresholdAmount, address target, bytes permit) returns(uint256, uint256, bytes32)
func (_Oneinch *OneinchSession) FillOrderToWithPermit(order OrderLibOrder, signature []byte, interaction []byte, makingAmount *big.Int, takingAmount *big.Int, skipPermitAndThresholdAmount *big.Int, target common.Address, permit []byte) (*types.Transaction, error) {
	return _Oneinch.Contract.FillOrderToWithPermit(&_Oneinch.TransactOpts, order, signature, interaction, makingAmount, takingAmount, skipPermitAndThresholdAmount, target, permit)
}

// FillOrderToWithPermit is a paid mutator transaction binding the contract method 0xd365c695.
//
// Solidity: function fillOrderToWithPermit((uint256,address,address,address,address,address,uint256,uint256,uint256,bytes) order, bytes signature, bytes interaction, uint256 makingAmount, uint256 takingAmount, uint256 skipPermitAndThresholdAmount, address target, bytes permit) returns(uint256, uint256, bytes32)
func (_Oneinch *OneinchTransactorSession) FillOrderToWithPermit(order OrderLibOrder, signature []byte, interaction []byte, makingAmount *big.Int, takingAmount *big.Int, skipPermitAndThresholdAmount *big.Int, target common.Address, permit []byte) (*types.Transaction, error) {
	return _Oneinch.Contract.FillOrderToWithPermit(&_Oneinch.TransactOpts, order, signature, interaction, makingAmount, takingAmount, skipPermitAndThresholdAmount, target, permit)
}

// IncreaseNonce is a paid mutator transaction binding the contract method 0xc53a0292.
//
// Solidity: function increaseNonce() returns()
func (_Oneinch *OneinchTransactor) IncreaseNonce(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Oneinch.contract.Transact(opts, "increaseNonce")
}

// IncreaseNonce is a paid mutator transaction binding the contract method 0xc53a0292.
//
// Solidity: function increaseNonce() returns()
func (_Oneinch *OneinchSession) IncreaseNonce() (*types.Transaction, error) {
	return _Oneinch.Contract.IncreaseNonce(&_Oneinch.TransactOpts)
}

// IncreaseNonce is a paid mutator transaction binding the contract method 0xc53a0292.
//
// Solidity: function increaseNonce() returns()
func (_Oneinch *OneinchTransactorSession) IncreaseNonce() (*types.Transaction, error) {
	return _Oneinch.Contract.IncreaseNonce(&_Oneinch.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Oneinch *OneinchTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Oneinch.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Oneinch *OneinchSession) RenounceOwnership() (*types.Transaction, error) {
	return _Oneinch.Contract.RenounceOwnership(&_Oneinch.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Oneinch *OneinchTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Oneinch.Contract.RenounceOwnership(&_Oneinch.TransactOpts)
}

// RescueFunds is a paid mutator transaction binding the contract method 0x78e3214f.
//
// Solidity: function rescueFunds(address token, uint256 amount) returns()
func (_Oneinch *OneinchTransactor) RescueFunds(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Oneinch.contract.Transact(opts, "rescueFunds", token, amount)
}

// RescueFunds is a paid mutator transaction binding the contract method 0x78e3214f.
//
// Solidity: function rescueFunds(address token, uint256 amount) returns()
func (_Oneinch *OneinchSession) RescueFunds(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Oneinch.Contract.RescueFunds(&_Oneinch.TransactOpts, token, amount)
}

// RescueFunds is a paid mutator transaction binding the contract method 0x78e3214f.
//
// Solidity: function rescueFunds(address token, uint256 amount) returns()
func (_Oneinch *OneinchTransactorSession) RescueFunds(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Oneinch.Contract.RescueFunds(&_Oneinch.TransactOpts, token, amount)
}

// Simulate is a paid mutator transaction binding the contract method 0xbd61951d.
//
// Solidity: function simulate(address target, bytes data) returns()
func (_Oneinch *OneinchTransactor) Simulate(opts *bind.TransactOpts, target common.Address, data []byte) (*types.Transaction, error) {
	return _Oneinch.contract.Transact(opts, "simulate", target, data)
}

// Simulate is a paid mutator transaction binding the contract method 0xbd61951d.
//
// Solidity: function simulate(address target, bytes data) returns()
func (_Oneinch *OneinchSession) Simulate(target common.Address, data []byte) (*types.Transaction, error) {
	return _Oneinch.Contract.Simulate(&_Oneinch.TransactOpts, target, data)
}

// Simulate is a paid mutator transaction binding the contract method 0xbd61951d.
//
// Solidity: function simulate(address target, bytes data) returns()
func (_Oneinch *OneinchTransactorSession) Simulate(target common.Address, data []byte) (*types.Transaction, error) {
	return _Oneinch.Contract.Simulate(&_Oneinch.TransactOpts, target, data)
}

// Swap is a paid mutator transaction binding the contract method 0x12aa3caf.
//
// Solidity: function swap(address executor, (address,address,address,address,uint256,uint256,uint256) desc, bytes permit, bytes data) payable returns(uint256 returnAmount, uint256 spentAmount)
func (_Oneinch *OneinchTransactor) Swap(opts *bind.TransactOpts, executor common.Address, desc GenericRouterSwapDescription, permit []byte, data []byte) (*types.Transaction, error) {
	return _Oneinch.contract.Transact(opts, "swap", executor, desc, permit, data)
}

// Swap is a paid mutator transaction binding the contract method 0x12aa3caf.
//
// Solidity: function swap(address executor, (address,address,address,address,uint256,uint256,uint256) desc, bytes permit, bytes data) payable returns(uint256 returnAmount, uint256 spentAmount)
func (_Oneinch *OneinchSession) Swap(executor common.Address, desc GenericRouterSwapDescription, permit []byte, data []byte) (*types.Transaction, error) {
	return _Oneinch.Contract.Swap(&_Oneinch.TransactOpts, executor, desc, permit, data)
}

// Swap is a paid mutator transaction binding the contract method 0x12aa3caf.
//
// Solidity: function swap(address executor, (address,address,address,address,uint256,uint256,uint256) desc, bytes permit, bytes data) payable returns(uint256 returnAmount, uint256 spentAmount)
func (_Oneinch *OneinchTransactorSession) Swap(executor common.Address, desc GenericRouterSwapDescription, permit []byte, data []byte) (*types.Transaction, error) {
	return _Oneinch.Contract.Swap(&_Oneinch.TransactOpts, executor, desc, permit, data)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Oneinch *OneinchTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Oneinch.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Oneinch *OneinchSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Oneinch.Contract.TransferOwnership(&_Oneinch.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Oneinch *OneinchTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Oneinch.Contract.TransferOwnership(&_Oneinch.TransactOpts, newOwner)
}

// UniswapV3Swap is a paid mutator transaction binding the contract method 0xe449022e.
//
// Solidity: function uniswapV3Swap(uint256 amount, uint256 minReturn, uint256[] pools) payable returns(uint256 returnAmount)
func (_Oneinch *OneinchTransactor) UniswapV3Swap(opts *bind.TransactOpts, amount *big.Int, minReturn *big.Int, pools []*big.Int) (*types.Transaction, error) {
	return _Oneinch.contract.Transact(opts, "uniswapV3Swap", amount, minReturn, pools)
}

// UniswapV3Swap is a paid mutator transaction binding the contract method 0xe449022e.
//
// Solidity: function uniswapV3Swap(uint256 amount, uint256 minReturn, uint256[] pools) payable returns(uint256 returnAmount)
func (_Oneinch *OneinchSession) UniswapV3Swap(amount *big.Int, minReturn *big.Int, pools []*big.Int) (*types.Transaction, error) {
	return _Oneinch.Contract.UniswapV3Swap(&_Oneinch.TransactOpts, amount, minReturn, pools)
}

// UniswapV3Swap is a paid mutator transaction binding the contract method 0xe449022e.
//
// Solidity: function uniswapV3Swap(uint256 amount, uint256 minReturn, uint256[] pools) payable returns(uint256 returnAmount)
func (_Oneinch *OneinchTransactorSession) UniswapV3Swap(amount *big.Int, minReturn *big.Int, pools []*big.Int) (*types.Transaction, error) {
	return _Oneinch.Contract.UniswapV3Swap(&_Oneinch.TransactOpts, amount, minReturn, pools)
}

// UniswapV3SwapCallback is a paid mutator transaction binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes ) returns()
func (_Oneinch *OneinchTransactor) UniswapV3SwapCallback(opts *bind.TransactOpts, amount0Delta *big.Int, amount1Delta *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _Oneinch.contract.Transact(opts, "uniswapV3SwapCallback", amount0Delta, amount1Delta, arg2)
}

// UniswapV3SwapCallback is a paid mutator transaction binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes ) returns()
func (_Oneinch *OneinchSession) UniswapV3SwapCallback(amount0Delta *big.Int, amount1Delta *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _Oneinch.Contract.UniswapV3SwapCallback(&_Oneinch.TransactOpts, amount0Delta, amount1Delta, arg2)
}

// UniswapV3SwapCallback is a paid mutator transaction binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes ) returns()
func (_Oneinch *OneinchTransactorSession) UniswapV3SwapCallback(amount0Delta *big.Int, amount1Delta *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _Oneinch.Contract.UniswapV3SwapCallback(&_Oneinch.TransactOpts, amount0Delta, amount1Delta, arg2)
}

// UniswapV3SwapTo is a paid mutator transaction binding the contract method 0xbc80f1a8.
//
// Solidity: function uniswapV3SwapTo(address recipient, uint256 amount, uint256 minReturn, uint256[] pools) payable returns(uint256 returnAmount)
func (_Oneinch *OneinchTransactor) UniswapV3SwapTo(opts *bind.TransactOpts, recipient common.Address, amount *big.Int, minReturn *big.Int, pools []*big.Int) (*types.Transaction, error) {
	return _Oneinch.contract.Transact(opts, "uniswapV3SwapTo", recipient, amount, minReturn, pools)
}

// UniswapV3SwapTo is a paid mutator transaction binding the contract method 0xbc80f1a8.
//
// Solidity: function uniswapV3SwapTo(address recipient, uint256 amount, uint256 minReturn, uint256[] pools) payable returns(uint256 returnAmount)
func (_Oneinch *OneinchSession) UniswapV3SwapTo(recipient common.Address, amount *big.Int, minReturn *big.Int, pools []*big.Int) (*types.Transaction, error) {
	return _Oneinch.Contract.UniswapV3SwapTo(&_Oneinch.TransactOpts, recipient, amount, minReturn, pools)
}

// UniswapV3SwapTo is a paid mutator transaction binding the contract method 0xbc80f1a8.
//
// Solidity: function uniswapV3SwapTo(address recipient, uint256 amount, uint256 minReturn, uint256[] pools) payable returns(uint256 returnAmount)
func (_Oneinch *OneinchTransactorSession) UniswapV3SwapTo(recipient common.Address, amount *big.Int, minReturn *big.Int, pools []*big.Int) (*types.Transaction, error) {
	return _Oneinch.Contract.UniswapV3SwapTo(&_Oneinch.TransactOpts, recipient, amount, minReturn, pools)
}

// UniswapV3SwapToWithPermit is a paid mutator transaction binding the contract method 0x2521b930.
//
// Solidity: function uniswapV3SwapToWithPermit(address recipient, address srcToken, uint256 amount, uint256 minReturn, uint256[] pools, bytes permit) returns(uint256 returnAmount)
func (_Oneinch *OneinchTransactor) UniswapV3SwapToWithPermit(opts *bind.TransactOpts, recipient common.Address, srcToken common.Address, amount *big.Int, minReturn *big.Int, pools []*big.Int, permit []byte) (*types.Transaction, error) {
	return _Oneinch.contract.Transact(opts, "uniswapV3SwapToWithPermit", recipient, srcToken, amount, minReturn, pools, permit)
}

// UniswapV3SwapToWithPermit is a paid mutator transaction binding the contract method 0x2521b930.
//
// Solidity: function uniswapV3SwapToWithPermit(address recipient, address srcToken, uint256 amount, uint256 minReturn, uint256[] pools, bytes permit) returns(uint256 returnAmount)
func (_Oneinch *OneinchSession) UniswapV3SwapToWithPermit(recipient common.Address, srcToken common.Address, amount *big.Int, minReturn *big.Int, pools []*big.Int, permit []byte) (*types.Transaction, error) {
	return _Oneinch.Contract.UniswapV3SwapToWithPermit(&_Oneinch.TransactOpts, recipient, srcToken, amount, minReturn, pools, permit)
}

// UniswapV3SwapToWithPermit is a paid mutator transaction binding the contract method 0x2521b930.
//
// Solidity: function uniswapV3SwapToWithPermit(address recipient, address srcToken, uint256 amount, uint256 minReturn, uint256[] pools, bytes permit) returns(uint256 returnAmount)
func (_Oneinch *OneinchTransactorSession) UniswapV3SwapToWithPermit(recipient common.Address, srcToken common.Address, amount *big.Int, minReturn *big.Int, pools []*big.Int, permit []byte) (*types.Transaction, error) {
	return _Oneinch.Contract.UniswapV3SwapToWithPermit(&_Oneinch.TransactOpts, recipient, srcToken, amount, minReturn, pools, permit)
}

// Unoswap is a paid mutator transaction binding the contract method 0x0502b1c5.
//
// Solidity: function unoswap(address srcToken, uint256 amount, uint256 minReturn, uint256[] pools) payable returns(uint256 returnAmount)
func (_Oneinch *OneinchTransactor) Unoswap(opts *bind.TransactOpts, srcToken common.Address, amount *big.Int, minReturn *big.Int, pools []*big.Int) (*types.Transaction, error) {
	return _Oneinch.contract.Transact(opts, "unoswap", srcToken, amount, minReturn, pools)
}

// Unoswap is a paid mutator transaction binding the contract method 0x0502b1c5.
//
// Solidity: function unoswap(address srcToken, uint256 amount, uint256 minReturn, uint256[] pools) payable returns(uint256 returnAmount)
func (_Oneinch *OneinchSession) Unoswap(srcToken common.Address, amount *big.Int, minReturn *big.Int, pools []*big.Int) (*types.Transaction, error) {
	return _Oneinch.Contract.Unoswap(&_Oneinch.TransactOpts, srcToken, amount, minReturn, pools)
}

// Unoswap is a paid mutator transaction binding the contract method 0x0502b1c5.
//
// Solidity: function unoswap(address srcToken, uint256 amount, uint256 minReturn, uint256[] pools) payable returns(uint256 returnAmount)
func (_Oneinch *OneinchTransactorSession) Unoswap(srcToken common.Address, amount *big.Int, minReturn *big.Int, pools []*big.Int) (*types.Transaction, error) {
	return _Oneinch.Contract.Unoswap(&_Oneinch.TransactOpts, srcToken, amount, minReturn, pools)
}

// UnoswapTo is a paid mutator transaction binding the contract method 0xf78dc253.
//
// Solidity: function unoswapTo(address recipient, address srcToken, uint256 amount, uint256 minReturn, uint256[] pools) payable returns(uint256 returnAmount)
func (_Oneinch *OneinchTransactor) UnoswapTo(opts *bind.TransactOpts, recipient common.Address, srcToken common.Address, amount *big.Int, minReturn *big.Int, pools []*big.Int) (*types.Transaction, error) {
	return _Oneinch.contract.Transact(opts, "unoswapTo", recipient, srcToken, amount, minReturn, pools)
}

// UnoswapTo is a paid mutator transaction binding the contract method 0xf78dc253.
//
// Solidity: function unoswapTo(address recipient, address srcToken, uint256 amount, uint256 minReturn, uint256[] pools) payable returns(uint256 returnAmount)
func (_Oneinch *OneinchSession) UnoswapTo(recipient common.Address, srcToken common.Address, amount *big.Int, minReturn *big.Int, pools []*big.Int) (*types.Transaction, error) {
	return _Oneinch.Contract.UnoswapTo(&_Oneinch.TransactOpts, recipient, srcToken, amount, minReturn, pools)
}

// UnoswapTo is a paid mutator transaction binding the contract method 0xf78dc253.
//
// Solidity: function unoswapTo(address recipient, address srcToken, uint256 amount, uint256 minReturn, uint256[] pools) payable returns(uint256 returnAmount)
func (_Oneinch *OneinchTransactorSession) UnoswapTo(recipient common.Address, srcToken common.Address, amount *big.Int, minReturn *big.Int, pools []*big.Int) (*types.Transaction, error) {
	return _Oneinch.Contract.UnoswapTo(&_Oneinch.TransactOpts, recipient, srcToken, amount, minReturn, pools)
}

// UnoswapToWithPermit is a paid mutator transaction binding the contract method 0x3c15fd91.
//
// Solidity: function unoswapToWithPermit(address recipient, address srcToken, uint256 amount, uint256 minReturn, uint256[] pools, bytes permit) returns(uint256 returnAmount)
func (_Oneinch *OneinchTransactor) UnoswapToWithPermit(opts *bind.TransactOpts, recipient common.Address, srcToken common.Address, amount *big.Int, minReturn *big.Int, pools []*big.Int, permit []byte) (*types.Transaction, error) {
	return _Oneinch.contract.Transact(opts, "unoswapToWithPermit", recipient, srcToken, amount, minReturn, pools, permit)
}

// UnoswapToWithPermit is a paid mutator transaction binding the contract method 0x3c15fd91.
//
// Solidity: function unoswapToWithPermit(address recipient, address srcToken, uint256 amount, uint256 minReturn, uint256[] pools, bytes permit) returns(uint256 returnAmount)
func (_Oneinch *OneinchSession) UnoswapToWithPermit(recipient common.Address, srcToken common.Address, amount *big.Int, minReturn *big.Int, pools []*big.Int, permit []byte) (*types.Transaction, error) {
	return _Oneinch.Contract.UnoswapToWithPermit(&_Oneinch.TransactOpts, recipient, srcToken, amount, minReturn, pools, permit)
}

// UnoswapToWithPermit is a paid mutator transaction binding the contract method 0x3c15fd91.
//
// Solidity: function unoswapToWithPermit(address recipient, address srcToken, uint256 amount, uint256 minReturn, uint256[] pools, bytes permit) returns(uint256 returnAmount)
func (_Oneinch *OneinchTransactorSession) UnoswapToWithPermit(recipient common.Address, srcToken common.Address, amount *big.Int, minReturn *big.Int, pools []*big.Int, permit []byte) (*types.Transaction, error) {
	return _Oneinch.Contract.UnoswapToWithPermit(&_Oneinch.TransactOpts, recipient, srcToken, amount, minReturn, pools, permit)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Oneinch *OneinchTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Oneinch.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Oneinch *OneinchSession) Receive() (*types.Transaction, error) {
	return _Oneinch.Contract.Receive(&_Oneinch.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Oneinch *OneinchTransactorSession) Receive() (*types.Transaction, error) {
	return _Oneinch.Contract.Receive(&_Oneinch.TransactOpts)
}

// OneinchNonceIncreasedIterator is returned from FilterNonceIncreased and is used to iterate over the raw logs and unpacked data for NonceIncreased events raised by the Oneinch contract.
type OneinchNonceIncreasedIterator struct {
	Event *OneinchNonceIncreased // Event containing the contract specifics and raw log

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
func (it *OneinchNonceIncreasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OneinchNonceIncreased)
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
		it.Event = new(OneinchNonceIncreased)
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
func (it *OneinchNonceIncreasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OneinchNonceIncreasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OneinchNonceIncreased represents a NonceIncreased event raised by the Oneinch contract.
type OneinchNonceIncreased struct {
	Maker    common.Address
	NewNonce *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNonceIncreased is a free log retrieval operation binding the contract event 0xfc69110dd11eb791755e4abd6b7d281bae236de95736d38a23782814be5e10db.
//
// Solidity: event NonceIncreased(address indexed maker, uint256 newNonce)
func (_Oneinch *OneinchFilterer) FilterNonceIncreased(opts *bind.FilterOpts, maker []common.Address) (*OneinchNonceIncreasedIterator, error) {

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}

	logs, sub, err := _Oneinch.contract.FilterLogs(opts, "NonceIncreased", makerRule)
	if err != nil {
		return nil, err
	}
	return &OneinchNonceIncreasedIterator{contract: _Oneinch.contract, event: "NonceIncreased", logs: logs, sub: sub}, nil
}

// WatchNonceIncreased is a free log subscription operation binding the contract event 0xfc69110dd11eb791755e4abd6b7d281bae236de95736d38a23782814be5e10db.
//
// Solidity: event NonceIncreased(address indexed maker, uint256 newNonce)
func (_Oneinch *OneinchFilterer) WatchNonceIncreased(opts *bind.WatchOpts, sink chan<- *OneinchNonceIncreased, maker []common.Address) (event.Subscription, error) {

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}

	logs, sub, err := _Oneinch.contract.WatchLogs(opts, "NonceIncreased", makerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OneinchNonceIncreased)
				if err := _Oneinch.contract.UnpackLog(event, "NonceIncreased", log); err != nil {
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

// ParseNonceIncreased is a log parse operation binding the contract event 0xfc69110dd11eb791755e4abd6b7d281bae236de95736d38a23782814be5e10db.
//
// Solidity: event NonceIncreased(address indexed maker, uint256 newNonce)
func (_Oneinch *OneinchFilterer) ParseNonceIncreased(log types.Log) (*OneinchNonceIncreased, error) {
	event := new(OneinchNonceIncreased)
	if err := _Oneinch.contract.UnpackLog(event, "NonceIncreased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OneinchOrderCanceledIterator is returned from FilterOrderCanceled and is used to iterate over the raw logs and unpacked data for OrderCanceled events raised by the Oneinch contract.
type OneinchOrderCanceledIterator struct {
	Event *OneinchOrderCanceled // Event containing the contract specifics and raw log

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
func (it *OneinchOrderCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OneinchOrderCanceled)
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
		it.Event = new(OneinchOrderCanceled)
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
func (it *OneinchOrderCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OneinchOrderCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OneinchOrderCanceled represents a OrderCanceled event raised by the Oneinch contract.
type OneinchOrderCanceled struct {
	Maker        common.Address
	OrderHash    [32]byte
	RemainingRaw *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterOrderCanceled is a free log retrieval operation binding the contract event 0xcbfa7d191838ece7ba4783ca3a30afd316619b7f368094b57ee7ffde9a923db1.
//
// Solidity: event OrderCanceled(address indexed maker, bytes32 orderHash, uint256 remainingRaw)
func (_Oneinch *OneinchFilterer) FilterOrderCanceled(opts *bind.FilterOpts, maker []common.Address) (*OneinchOrderCanceledIterator, error) {

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}

	logs, sub, err := _Oneinch.contract.FilterLogs(opts, "OrderCanceled", makerRule)
	if err != nil {
		return nil, err
	}
	return &OneinchOrderCanceledIterator{contract: _Oneinch.contract, event: "OrderCanceled", logs: logs, sub: sub}, nil
}

// WatchOrderCanceled is a free log subscription operation binding the contract event 0xcbfa7d191838ece7ba4783ca3a30afd316619b7f368094b57ee7ffde9a923db1.
//
// Solidity: event OrderCanceled(address indexed maker, bytes32 orderHash, uint256 remainingRaw)
func (_Oneinch *OneinchFilterer) WatchOrderCanceled(opts *bind.WatchOpts, sink chan<- *OneinchOrderCanceled, maker []common.Address) (event.Subscription, error) {

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}

	logs, sub, err := _Oneinch.contract.WatchLogs(opts, "OrderCanceled", makerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OneinchOrderCanceled)
				if err := _Oneinch.contract.UnpackLog(event, "OrderCanceled", log); err != nil {
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

// ParseOrderCanceled is a log parse operation binding the contract event 0xcbfa7d191838ece7ba4783ca3a30afd316619b7f368094b57ee7ffde9a923db1.
//
// Solidity: event OrderCanceled(address indexed maker, bytes32 orderHash, uint256 remainingRaw)
func (_Oneinch *OneinchFilterer) ParseOrderCanceled(log types.Log) (*OneinchOrderCanceled, error) {
	event := new(OneinchOrderCanceled)
	if err := _Oneinch.contract.UnpackLog(event, "OrderCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OneinchOrderFilledIterator is returned from FilterOrderFilled and is used to iterate over the raw logs and unpacked data for OrderFilled events raised by the Oneinch contract.
type OneinchOrderFilledIterator struct {
	Event *OneinchOrderFilled // Event containing the contract specifics and raw log

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
func (it *OneinchOrderFilledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OneinchOrderFilled)
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
		it.Event = new(OneinchOrderFilled)
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
func (it *OneinchOrderFilledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OneinchOrderFilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OneinchOrderFilled represents a OrderFilled event raised by the Oneinch contract.
type OneinchOrderFilled struct {
	Maker     common.Address
	OrderHash [32]byte
	Remaining *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOrderFilled is a free log retrieval operation binding the contract event 0xb9ed0243fdf00f0545c63a0af8850c090d86bb46682baec4bf3c496814fe4f02.
//
// Solidity: event OrderFilled(address indexed maker, bytes32 orderHash, uint256 remaining)
func (_Oneinch *OneinchFilterer) FilterOrderFilled(opts *bind.FilterOpts, maker []common.Address) (*OneinchOrderFilledIterator, error) {

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}

	logs, sub, err := _Oneinch.contract.FilterLogs(opts, "OrderFilled", makerRule)
	if err != nil {
		return nil, err
	}
	return &OneinchOrderFilledIterator{contract: _Oneinch.contract, event: "OrderFilled", logs: logs, sub: sub}, nil
}

// WatchOrderFilled is a free log subscription operation binding the contract event 0xb9ed0243fdf00f0545c63a0af8850c090d86bb46682baec4bf3c496814fe4f02.
//
// Solidity: event OrderFilled(address indexed maker, bytes32 orderHash, uint256 remaining)
func (_Oneinch *OneinchFilterer) WatchOrderFilled(opts *bind.WatchOpts, sink chan<- *OneinchOrderFilled, maker []common.Address) (event.Subscription, error) {

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}

	logs, sub, err := _Oneinch.contract.WatchLogs(opts, "OrderFilled", makerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OneinchOrderFilled)
				if err := _Oneinch.contract.UnpackLog(event, "OrderFilled", log); err != nil {
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

// ParseOrderFilled is a log parse operation binding the contract event 0xb9ed0243fdf00f0545c63a0af8850c090d86bb46682baec4bf3c496814fe4f02.
//
// Solidity: event OrderFilled(address indexed maker, bytes32 orderHash, uint256 remaining)
func (_Oneinch *OneinchFilterer) ParseOrderFilled(log types.Log) (*OneinchOrderFilled, error) {
	event := new(OneinchOrderFilled)
	if err := _Oneinch.contract.UnpackLog(event, "OrderFilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OneinchOrderFilledRFQIterator is returned from FilterOrderFilledRFQ and is used to iterate over the raw logs and unpacked data for OrderFilledRFQ events raised by the Oneinch contract.
type OneinchOrderFilledRFQIterator struct {
	Event *OneinchOrderFilledRFQ // Event containing the contract specifics and raw log

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
func (it *OneinchOrderFilledRFQIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OneinchOrderFilledRFQ)
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
		it.Event = new(OneinchOrderFilledRFQ)
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
func (it *OneinchOrderFilledRFQIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OneinchOrderFilledRFQIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OneinchOrderFilledRFQ represents a OrderFilledRFQ event raised by the Oneinch contract.
type OneinchOrderFilledRFQ struct {
	OrderHash    [32]byte
	MakingAmount *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterOrderFilledRFQ is a free log retrieval operation binding the contract event 0xc3b639f02b125bfa160e50739b8c44eb2d1b6908e2b6d5925c6d770f2ca78127.
//
// Solidity: event OrderFilledRFQ(bytes32 orderHash, uint256 makingAmount)
func (_Oneinch *OneinchFilterer) FilterOrderFilledRFQ(opts *bind.FilterOpts) (*OneinchOrderFilledRFQIterator, error) {

	logs, sub, err := _Oneinch.contract.FilterLogs(opts, "OrderFilledRFQ")
	if err != nil {
		return nil, err
	}
	return &OneinchOrderFilledRFQIterator{contract: _Oneinch.contract, event: "OrderFilledRFQ", logs: logs, sub: sub}, nil
}

// WatchOrderFilledRFQ is a free log subscription operation binding the contract event 0xc3b639f02b125bfa160e50739b8c44eb2d1b6908e2b6d5925c6d770f2ca78127.
//
// Solidity: event OrderFilledRFQ(bytes32 orderHash, uint256 makingAmount)
func (_Oneinch *OneinchFilterer) WatchOrderFilledRFQ(opts *bind.WatchOpts, sink chan<- *OneinchOrderFilledRFQ) (event.Subscription, error) {

	logs, sub, err := _Oneinch.contract.WatchLogs(opts, "OrderFilledRFQ")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OneinchOrderFilledRFQ)
				if err := _Oneinch.contract.UnpackLog(event, "OrderFilledRFQ", log); err != nil {
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

// ParseOrderFilledRFQ is a log parse operation binding the contract event 0xc3b639f02b125bfa160e50739b8c44eb2d1b6908e2b6d5925c6d770f2ca78127.
//
// Solidity: event OrderFilledRFQ(bytes32 orderHash, uint256 makingAmount)
func (_Oneinch *OneinchFilterer) ParseOrderFilledRFQ(log types.Log) (*OneinchOrderFilledRFQ, error) {
	event := new(OneinchOrderFilledRFQ)
	if err := _Oneinch.contract.UnpackLog(event, "OrderFilledRFQ", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OneinchOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Oneinch contract.
type OneinchOwnershipTransferredIterator struct {
	Event *OneinchOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *OneinchOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OneinchOwnershipTransferred)
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
		it.Event = new(OneinchOwnershipTransferred)
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
func (it *OneinchOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OneinchOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OneinchOwnershipTransferred represents a OwnershipTransferred event raised by the Oneinch contract.
type OneinchOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Oneinch *OneinchFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OneinchOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Oneinch.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OneinchOwnershipTransferredIterator{contract: _Oneinch.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Oneinch *OneinchFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OneinchOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Oneinch.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OneinchOwnershipTransferred)
				if err := _Oneinch.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Oneinch *OneinchFilterer) ParseOwnershipTransferred(log types.Log) (*OneinchOwnershipTransferred, error) {
	event := new(OneinchOwnershipTransferred)
	if err := _Oneinch.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
