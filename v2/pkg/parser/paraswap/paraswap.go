// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package paraswap

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

// AugustusRFQOrder is an auto generated low-level Go binding around an user-defined struct.
type AugustusRFQOrder struct {
	NonceAndMeta *big.Int
	Expiry       *big.Int
	MakerAsset   common.Address
	TakerAsset   common.Address
	Maker        common.Address
	Taker        common.Address
	MakerAmount  *big.Int
	TakerAmount  *big.Int
}

// AugustusRFQOrderInfo is an auto generated low-level Go binding around an user-defined struct.
type AugustusRFQOrderInfo struct {
	Order                AugustusRFQOrder
	Signature            []byte
	TakerTokenFillAmount *big.Int
	PermitTakerAsset     []byte
	PermitMakerAsset     []byte
}

// AugustusRFQOrderNFT is an auto generated low-level Go binding around an user-defined struct.
type AugustusRFQOrderNFT struct {
	NonceAndMeta *big.Int
	Expiry       *big.Int
	MakerAsset   *big.Int
	MakerAssetId *big.Int
	TakerAsset   *big.Int
	TakerAssetId *big.Int
	Maker        common.Address
	Taker        common.Address
	MakerAmount  *big.Int
	TakerAmount  *big.Int
}

// AugustusRFQOrderNFTInfo is an auto generated low-level Go binding around an user-defined struct.
type AugustusRFQOrderNFTInfo struct {
	Order                AugustusRFQOrderNFT
	Signature            []byte
	TakerTokenFillAmount *big.Int
	PermitTakerAsset     []byte
	PermitMakerAsset     []byte
}

// OrderFilledMetaData contains all meta data concerning the OrderFilled contract.
var OrderFilledMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"}],\"name\":\"OrderCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"makerAsset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"makerAmount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"takerAsset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"takerAmount\",\"type\":\"uint256\"}],\"name\":\"OrderFilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"makerAsset\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"makerAssetId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"makerAmount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"takerAsset\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"takerAssetId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"takerAmount\",\"type\":\"uint256\"}],\"name\":\"OrderFilledNFT\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"FILLED_ORDER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RFQ_LIMIT_NFT_ORDER_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RFQ_LIMIT_ORDER_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UNFILLED_ORDER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"nonceAndMeta\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"expiry\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"makerAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAmount\",\"type\":\"uint256\"}],\"internalType\":\"structAugustusRFQ.Order\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"takerTokenFillAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"permitTakerAsset\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"permitMakerAsset\",\"type\":\"bytes\"}],\"internalType\":\"structAugustusRFQ.OrderInfo[]\",\"name\":\"orderInfos\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"batchFillOrderWithTarget\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"nonceAndMeta\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"expiry\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"makerAsset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerAssetId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAsset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAmount\",\"type\":\"uint256\"}],\"internalType\":\"structAugustusRFQ.OrderNFT\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"takerTokenFillAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"permitTakerAsset\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"permitMakerAsset\",\"type\":\"bytes\"}],\"internalType\":\"structAugustusRFQ.OrderNFTInfo[]\",\"name\":\"orderInfos\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"batchFillOrderWithTargetNFT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"name\":\"cancelOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"orderHashes\",\"type\":\"bytes32[]\"}],\"name\":\"cancelOrders\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"nonceAndMeta\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"expiry\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"makerAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAmount\",\"type\":\"uint256\"}],\"internalType\":\"structAugustusRFQ.Order\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"fillOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"nonceAndMeta\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"expiry\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"makerAsset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerAssetId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAsset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAmount\",\"type\":\"uint256\"}],\"internalType\":\"structAugustusRFQ.OrderNFT\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"fillOrderNFT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"nonceAndMeta\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"expiry\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"makerAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAmount\",\"type\":\"uint256\"}],\"internalType\":\"structAugustusRFQ.Order\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"fillOrderWithTarget\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"nonceAndMeta\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"expiry\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"makerAsset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerAssetId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAsset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAmount\",\"type\":\"uint256\"}],\"internalType\":\"structAugustusRFQ.OrderNFT\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"fillOrderWithTargetNFT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"orderHashes\",\"type\":\"bytes32[]\"}],\"name\":\"getRemainingOrderBalance\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"remainingBalances\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"nonceAndMeta\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"expiry\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"makerAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAmount\",\"type\":\"uint256\"}],\"internalType\":\"structAugustusRFQ.Order\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"takerTokenFillAmount\",\"type\":\"uint256\"}],\"name\":\"partialFillOrder\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"makerTokenFilledAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"nonceAndMeta\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"expiry\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"makerAsset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerAssetId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAsset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAmount\",\"type\":\"uint256\"}],\"internalType\":\"structAugustusRFQ.OrderNFT\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"takerTokenFillAmount\",\"type\":\"uint256\"}],\"name\":\"partialFillOrderNFT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"makerTokenFilledAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"nonceAndMeta\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"expiry\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"makerAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAmount\",\"type\":\"uint256\"}],\"internalType\":\"structAugustusRFQ.Order\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"takerTokenFillAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"partialFillOrderWithTarget\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"makerTokenFilledAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"nonceAndMeta\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"expiry\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"makerAsset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerAssetId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAsset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAmount\",\"type\":\"uint256\"}],\"internalType\":\"structAugustusRFQ.OrderNFT\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"takerTokenFillAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"partialFillOrderWithTargetNFT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"makerTokenFilledAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"nonceAndMeta\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"expiry\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"makerAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAmount\",\"type\":\"uint256\"}],\"internalType\":\"structAugustusRFQ.Order\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"takerTokenFillAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"permitTakerAsset\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"permitMakerAsset\",\"type\":\"bytes\"}],\"name\":\"partialFillOrderWithTargetPermit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"makerTokenFilledAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"nonceAndMeta\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"expiry\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"makerAsset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerAssetId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAsset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAmount\",\"type\":\"uint256\"}],\"internalType\":\"structAugustusRFQ.OrderNFT\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"takerTokenFillAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"permitTakerAsset\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"permitMakerAsset\",\"type\":\"bytes\"}],\"name\":\"partialFillOrderWithTargetPermitNFT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"makerTokenFilledAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"remaining\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"nonceAndMeta\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"expiry\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"makerAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAmount\",\"type\":\"uint256\"}],\"internalType\":\"structAugustusRFQ.Order\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"takerTokenFillAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"permitTakerAsset\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"permitMakerAsset\",\"type\":\"bytes\"}],\"internalType\":\"structAugustusRFQ.OrderInfo[]\",\"name\":\"orderInfos\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"makerFillAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"tryBatchFillOrderMakerAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"nonceAndMeta\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"expiry\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"makerAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAmount\",\"type\":\"uint256\"}],\"internalType\":\"structAugustusRFQ.Order\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"takerTokenFillAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"permitTakerAsset\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"permitMakerAsset\",\"type\":\"bytes\"}],\"internalType\":\"structAugustusRFQ.OrderInfo[]\",\"name\":\"orderInfos\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"takerFillAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"tryBatchFillOrderTakerAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// OrderFilledABI is the input ABI used to generate the binding from.
// Deprecated: Use OrderFilledMetaData.ABI instead.
var OrderFilledABI = OrderFilledMetaData.ABI

// OrderFilled is an auto generated Go binding around an Ethereum contract.
type OrderFilled struct {
	OrderFilledCaller     // Read-only binding to the contract
	OrderFilledTransactor // Write-only binding to the contract
	OrderFilledFilterer   // Log filterer for contract events
}

// OrderFilledCaller is an auto generated read-only Go binding around an Ethereum contract.
type OrderFilledCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OrderFilledTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OrderFilledTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OrderFilledFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OrderFilledFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OrderFilledSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OrderFilledSession struct {
	Contract     *OrderFilled      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OrderFilledCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OrderFilledCallerSession struct {
	Contract *OrderFilledCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// OrderFilledTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OrderFilledTransactorSession struct {
	Contract     *OrderFilledTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// OrderFilledRaw is an auto generated low-level Go binding around an Ethereum contract.
type OrderFilledRaw struct {
	Contract *OrderFilled // Generic contract binding to access the raw methods on
}

// OrderFilledCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OrderFilledCallerRaw struct {
	Contract *OrderFilledCaller // Generic read-only contract binding to access the raw methods on
}

// OrderFilledTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OrderFilledTransactorRaw struct {
	Contract *OrderFilledTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOrderFilled creates a new instance of OrderFilled, bound to a specific deployed contract.
func NewOrderFilled(address common.Address, backend bind.ContractBackend) (*OrderFilled, error) {
	contract, err := bindOrderFilled(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OrderFilled{OrderFilledCaller: OrderFilledCaller{contract: contract}, OrderFilledTransactor: OrderFilledTransactor{contract: contract}, OrderFilledFilterer: OrderFilledFilterer{contract: contract}}, nil
}

// NewOrderFilledCaller creates a new read-only instance of OrderFilled, bound to a specific deployed contract.
func NewOrderFilledCaller(address common.Address, caller bind.ContractCaller) (*OrderFilledCaller, error) {
	contract, err := bindOrderFilled(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OrderFilledCaller{contract: contract}, nil
}

// NewOrderFilledTransactor creates a new write-only instance of OrderFilled, bound to a specific deployed contract.
func NewOrderFilledTransactor(address common.Address, transactor bind.ContractTransactor) (*OrderFilledTransactor, error) {
	contract, err := bindOrderFilled(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OrderFilledTransactor{contract: contract}, nil
}

// NewOrderFilledFilterer creates a new log filterer instance of OrderFilled, bound to a specific deployed contract.
func NewOrderFilledFilterer(address common.Address, filterer bind.ContractFilterer) (*OrderFilledFilterer, error) {
	contract, err := bindOrderFilled(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OrderFilledFilterer{contract: contract}, nil
}

// bindOrderFilled binds a generic wrapper to an already deployed contract.
func bindOrderFilled(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OrderFilledMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OrderFilled *OrderFilledRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OrderFilled.Contract.OrderFilledCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OrderFilled *OrderFilledRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OrderFilled.Contract.OrderFilledTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OrderFilled *OrderFilledRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OrderFilled.Contract.OrderFilledTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OrderFilled *OrderFilledCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OrderFilled.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OrderFilled *OrderFilledTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OrderFilled.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OrderFilled *OrderFilledTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OrderFilled.Contract.contract.Transact(opts, method, params...)
}

// FILLEDORDER is a free data retrieval call binding the contract method 0x2ea1ee84.
//
// Solidity: function FILLED_ORDER() view returns(uint256)
func (_OrderFilled *OrderFilledCaller) FILLEDORDER(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OrderFilled.contract.Call(opts, &out, "FILLED_ORDER")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FILLEDORDER is a free data retrieval call binding the contract method 0x2ea1ee84.
//
// Solidity: function FILLED_ORDER() view returns(uint256)
func (_OrderFilled *OrderFilledSession) FILLEDORDER() (*big.Int, error) {
	return _OrderFilled.Contract.FILLEDORDER(&_OrderFilled.CallOpts)
}

// FILLEDORDER is a free data retrieval call binding the contract method 0x2ea1ee84.
//
// Solidity: function FILLED_ORDER() view returns(uint256)
func (_OrderFilled *OrderFilledCallerSession) FILLEDORDER() (*big.Int, error) {
	return _OrderFilled.Contract.FILLEDORDER(&_OrderFilled.CallOpts)
}

// RFQLIMITNFTORDERTYPEHASH is a free data retrieval call binding the contract method 0xde77aaf9.
//
// Solidity: function RFQ_LIMIT_NFT_ORDER_TYPEHASH() view returns(bytes32)
func (_OrderFilled *OrderFilledCaller) RFQLIMITNFTORDERTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _OrderFilled.contract.Call(opts, &out, "RFQ_LIMIT_NFT_ORDER_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RFQLIMITNFTORDERTYPEHASH is a free data retrieval call binding the contract method 0xde77aaf9.
//
// Solidity: function RFQ_LIMIT_NFT_ORDER_TYPEHASH() view returns(bytes32)
func (_OrderFilled *OrderFilledSession) RFQLIMITNFTORDERTYPEHASH() ([32]byte, error) {
	return _OrderFilled.Contract.RFQLIMITNFTORDERTYPEHASH(&_OrderFilled.CallOpts)
}

// RFQLIMITNFTORDERTYPEHASH is a free data retrieval call binding the contract method 0xde77aaf9.
//
// Solidity: function RFQ_LIMIT_NFT_ORDER_TYPEHASH() view returns(bytes32)
func (_OrderFilled *OrderFilledCallerSession) RFQLIMITNFTORDERTYPEHASH() ([32]byte, error) {
	return _OrderFilled.Contract.RFQLIMITNFTORDERTYPEHASH(&_OrderFilled.CallOpts)
}

// RFQLIMITORDERTYPEHASH is a free data retrieval call binding the contract method 0x8a3ae43d.
//
// Solidity: function RFQ_LIMIT_ORDER_TYPEHASH() view returns(bytes32)
func (_OrderFilled *OrderFilledCaller) RFQLIMITORDERTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _OrderFilled.contract.Call(opts, &out, "RFQ_LIMIT_ORDER_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RFQLIMITORDERTYPEHASH is a free data retrieval call binding the contract method 0x8a3ae43d.
//
// Solidity: function RFQ_LIMIT_ORDER_TYPEHASH() view returns(bytes32)
func (_OrderFilled *OrderFilledSession) RFQLIMITORDERTYPEHASH() ([32]byte, error) {
	return _OrderFilled.Contract.RFQLIMITORDERTYPEHASH(&_OrderFilled.CallOpts)
}

// RFQLIMITORDERTYPEHASH is a free data retrieval call binding the contract method 0x8a3ae43d.
//
// Solidity: function RFQ_LIMIT_ORDER_TYPEHASH() view returns(bytes32)
func (_OrderFilled *OrderFilledCallerSession) RFQLIMITORDERTYPEHASH() ([32]byte, error) {
	return _OrderFilled.Contract.RFQLIMITORDERTYPEHASH(&_OrderFilled.CallOpts)
}

// UNFILLEDORDER is a free data retrieval call binding the contract method 0xff7e506a.
//
// Solidity: function UNFILLED_ORDER() view returns(uint256)
func (_OrderFilled *OrderFilledCaller) UNFILLEDORDER(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OrderFilled.contract.Call(opts, &out, "UNFILLED_ORDER")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UNFILLEDORDER is a free data retrieval call binding the contract method 0xff7e506a.
//
// Solidity: function UNFILLED_ORDER() view returns(uint256)
func (_OrderFilled *OrderFilledSession) UNFILLEDORDER() (*big.Int, error) {
	return _OrderFilled.Contract.UNFILLEDORDER(&_OrderFilled.CallOpts)
}

// UNFILLEDORDER is a free data retrieval call binding the contract method 0xff7e506a.
//
// Solidity: function UNFILLED_ORDER() view returns(uint256)
func (_OrderFilled *OrderFilledCallerSession) UNFILLEDORDER() (*big.Int, error) {
	return _OrderFilled.Contract.UNFILLEDORDER(&_OrderFilled.CallOpts)
}

// GetRemainingOrderBalance is a free data retrieval call binding the contract method 0x01568b83.
//
// Solidity: function getRemainingOrderBalance(address maker, bytes32[] orderHashes) view returns(uint256[] remainingBalances)
func (_OrderFilled *OrderFilledCaller) GetRemainingOrderBalance(opts *bind.CallOpts, maker common.Address, orderHashes [][32]byte) ([]*big.Int, error) {
	var out []interface{}
	err := _OrderFilled.contract.Call(opts, &out, "getRemainingOrderBalance", maker, orderHashes)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetRemainingOrderBalance is a free data retrieval call binding the contract method 0x01568b83.
//
// Solidity: function getRemainingOrderBalance(address maker, bytes32[] orderHashes) view returns(uint256[] remainingBalances)
func (_OrderFilled *OrderFilledSession) GetRemainingOrderBalance(maker common.Address, orderHashes [][32]byte) ([]*big.Int, error) {
	return _OrderFilled.Contract.GetRemainingOrderBalance(&_OrderFilled.CallOpts, maker, orderHashes)
}

// GetRemainingOrderBalance is a free data retrieval call binding the contract method 0x01568b83.
//
// Solidity: function getRemainingOrderBalance(address maker, bytes32[] orderHashes) view returns(uint256[] remainingBalances)
func (_OrderFilled *OrderFilledCallerSession) GetRemainingOrderBalance(maker common.Address, orderHashes [][32]byte) ([]*big.Int, error) {
	return _OrderFilled.Contract.GetRemainingOrderBalance(&_OrderFilled.CallOpts, maker, orderHashes)
}

// Remaining is a free data retrieval call binding the contract method 0x0b57f091.
//
// Solidity: function remaining(address , bytes32 ) view returns(uint256)
func (_OrderFilled *OrderFilledCaller) Remaining(opts *bind.CallOpts, arg0 common.Address, arg1 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _OrderFilled.contract.Call(opts, &out, "remaining", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Remaining is a free data retrieval call binding the contract method 0x0b57f091.
//
// Solidity: function remaining(address , bytes32 ) view returns(uint256)
func (_OrderFilled *OrderFilledSession) Remaining(arg0 common.Address, arg1 [32]byte) (*big.Int, error) {
	return _OrderFilled.Contract.Remaining(&_OrderFilled.CallOpts, arg0, arg1)
}

// Remaining is a free data retrieval call binding the contract method 0x0b57f091.
//
// Solidity: function remaining(address , bytes32 ) view returns(uint256)
func (_OrderFilled *OrderFilledCallerSession) Remaining(arg0 common.Address, arg1 [32]byte) (*big.Int, error) {
	return _OrderFilled.Contract.Remaining(&_OrderFilled.CallOpts, arg0, arg1)
}

// BatchFillOrderWithTarget is a paid mutator transaction binding the contract method 0x077822bd.
//
// Solidity: function batchFillOrderWithTarget(((uint256,uint128,address,address,address,address,uint256,uint256),bytes,uint256,bytes,bytes)[] orderInfos, address target) returns()
func (_OrderFilled *OrderFilledTransactor) BatchFillOrderWithTarget(opts *bind.TransactOpts, orderInfos []AugustusRFQOrderInfo, target common.Address) (*types.Transaction, error) {
	return _OrderFilled.contract.Transact(opts, "batchFillOrderWithTarget", orderInfos, target)
}

// BatchFillOrderWithTarget is a paid mutator transaction binding the contract method 0x077822bd.
//
// Solidity: function batchFillOrderWithTarget(((uint256,uint128,address,address,address,address,uint256,uint256),bytes,uint256,bytes,bytes)[] orderInfos, address target) returns()
func (_OrderFilled *OrderFilledSession) BatchFillOrderWithTarget(orderInfos []AugustusRFQOrderInfo, target common.Address) (*types.Transaction, error) {
	return _OrderFilled.Contract.BatchFillOrderWithTarget(&_OrderFilled.TransactOpts, orderInfos, target)
}

// BatchFillOrderWithTarget is a paid mutator transaction binding the contract method 0x077822bd.
//
// Solidity: function batchFillOrderWithTarget(((uint256,uint128,address,address,address,address,uint256,uint256),bytes,uint256,bytes,bytes)[] orderInfos, address target) returns()
func (_OrderFilled *OrderFilledTransactorSession) BatchFillOrderWithTarget(orderInfos []AugustusRFQOrderInfo, target common.Address) (*types.Transaction, error) {
	return _OrderFilled.Contract.BatchFillOrderWithTarget(&_OrderFilled.TransactOpts, orderInfos, target)
}

// BatchFillOrderWithTargetNFT is a paid mutator transaction binding the contract method 0xc8b81d63.
//
// Solidity: function batchFillOrderWithTargetNFT(((uint256,uint128,uint256,uint256,uint256,uint256,address,address,uint256,uint256),bytes,uint256,bytes,bytes)[] orderInfos, address target) returns()
func (_OrderFilled *OrderFilledTransactor) BatchFillOrderWithTargetNFT(opts *bind.TransactOpts, orderInfos []AugustusRFQOrderNFTInfo, target common.Address) (*types.Transaction, error) {
	return _OrderFilled.contract.Transact(opts, "batchFillOrderWithTargetNFT", orderInfos, target)
}

// BatchFillOrderWithTargetNFT is a paid mutator transaction binding the contract method 0xc8b81d63.
//
// Solidity: function batchFillOrderWithTargetNFT(((uint256,uint128,uint256,uint256,uint256,uint256,address,address,uint256,uint256),bytes,uint256,bytes,bytes)[] orderInfos, address target) returns()
func (_OrderFilled *OrderFilledSession) BatchFillOrderWithTargetNFT(orderInfos []AugustusRFQOrderNFTInfo, target common.Address) (*types.Transaction, error) {
	return _OrderFilled.Contract.BatchFillOrderWithTargetNFT(&_OrderFilled.TransactOpts, orderInfos, target)
}

// BatchFillOrderWithTargetNFT is a paid mutator transaction binding the contract method 0xc8b81d63.
//
// Solidity: function batchFillOrderWithTargetNFT(((uint256,uint128,uint256,uint256,uint256,uint256,address,address,uint256,uint256),bytes,uint256,bytes,bytes)[] orderInfos, address target) returns()
func (_OrderFilled *OrderFilledTransactorSession) BatchFillOrderWithTargetNFT(orderInfos []AugustusRFQOrderNFTInfo, target common.Address) (*types.Transaction, error) {
	return _OrderFilled.Contract.BatchFillOrderWithTargetNFT(&_OrderFilled.TransactOpts, orderInfos, target)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x7489ec23.
//
// Solidity: function cancelOrder(bytes32 orderHash) returns()
func (_OrderFilled *OrderFilledTransactor) CancelOrder(opts *bind.TransactOpts, orderHash [32]byte) (*types.Transaction, error) {
	return _OrderFilled.contract.Transact(opts, "cancelOrder", orderHash)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x7489ec23.
//
// Solidity: function cancelOrder(bytes32 orderHash) returns()
func (_OrderFilled *OrderFilledSession) CancelOrder(orderHash [32]byte) (*types.Transaction, error) {
	return _OrderFilled.Contract.CancelOrder(&_OrderFilled.TransactOpts, orderHash)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x7489ec23.
//
// Solidity: function cancelOrder(bytes32 orderHash) returns()
func (_OrderFilled *OrderFilledTransactorSession) CancelOrder(orderHash [32]byte) (*types.Transaction, error) {
	return _OrderFilled.Contract.CancelOrder(&_OrderFilled.TransactOpts, orderHash)
}

// CancelOrders is a paid mutator transaction binding the contract method 0x21c77c96.
//
// Solidity: function cancelOrders(bytes32[] orderHashes) returns()
func (_OrderFilled *OrderFilledTransactor) CancelOrders(opts *bind.TransactOpts, orderHashes [][32]byte) (*types.Transaction, error) {
	return _OrderFilled.contract.Transact(opts, "cancelOrders", orderHashes)
}

// CancelOrders is a paid mutator transaction binding the contract method 0x21c77c96.
//
// Solidity: function cancelOrders(bytes32[] orderHashes) returns()
func (_OrderFilled *OrderFilledSession) CancelOrders(orderHashes [][32]byte) (*types.Transaction, error) {
	return _OrderFilled.Contract.CancelOrders(&_OrderFilled.TransactOpts, orderHashes)
}

// CancelOrders is a paid mutator transaction binding the contract method 0x21c77c96.
//
// Solidity: function cancelOrders(bytes32[] orderHashes) returns()
func (_OrderFilled *OrderFilledTransactorSession) CancelOrders(orderHashes [][32]byte) (*types.Transaction, error) {
	return _OrderFilled.Contract.CancelOrders(&_OrderFilled.TransactOpts, orderHashes)
}

// FillOrder is a paid mutator transaction binding the contract method 0x98f9b46b.
//
// Solidity: function fillOrder((uint256,uint128,address,address,address,address,uint256,uint256) order, bytes signature) returns()
func (_OrderFilled *OrderFilledTransactor) FillOrder(opts *bind.TransactOpts, order AugustusRFQOrder, signature []byte) (*types.Transaction, error) {
	return _OrderFilled.contract.Transact(opts, "fillOrder", order, signature)
}

// FillOrder is a paid mutator transaction binding the contract method 0x98f9b46b.
//
// Solidity: function fillOrder((uint256,uint128,address,address,address,address,uint256,uint256) order, bytes signature) returns()
func (_OrderFilled *OrderFilledSession) FillOrder(order AugustusRFQOrder, signature []byte) (*types.Transaction, error) {
	return _OrderFilled.Contract.FillOrder(&_OrderFilled.TransactOpts, order, signature)
}

// FillOrder is a paid mutator transaction binding the contract method 0x98f9b46b.
//
// Solidity: function fillOrder((uint256,uint128,address,address,address,address,uint256,uint256) order, bytes signature) returns()
func (_OrderFilled *OrderFilledTransactorSession) FillOrder(order AugustusRFQOrder, signature []byte) (*types.Transaction, error) {
	return _OrderFilled.Contract.FillOrder(&_OrderFilled.TransactOpts, order, signature)
}

// FillOrderNFT is a paid mutator transaction binding the contract method 0xbbbc2372.
//
// Solidity: function fillOrderNFT((uint256,uint128,uint256,uint256,uint256,uint256,address,address,uint256,uint256) order, bytes signature) returns()
func (_OrderFilled *OrderFilledTransactor) FillOrderNFT(opts *bind.TransactOpts, order AugustusRFQOrderNFT, signature []byte) (*types.Transaction, error) {
	return _OrderFilled.contract.Transact(opts, "fillOrderNFT", order, signature)
}

// FillOrderNFT is a paid mutator transaction binding the contract method 0xbbbc2372.
//
// Solidity: function fillOrderNFT((uint256,uint128,uint256,uint256,uint256,uint256,address,address,uint256,uint256) order, bytes signature) returns()
func (_OrderFilled *OrderFilledSession) FillOrderNFT(order AugustusRFQOrderNFT, signature []byte) (*types.Transaction, error) {
	return _OrderFilled.Contract.FillOrderNFT(&_OrderFilled.TransactOpts, order, signature)
}

// FillOrderNFT is a paid mutator transaction binding the contract method 0xbbbc2372.
//
// Solidity: function fillOrderNFT((uint256,uint128,uint256,uint256,uint256,uint256,address,address,uint256,uint256) order, bytes signature) returns()
func (_OrderFilled *OrderFilledTransactorSession) FillOrderNFT(order AugustusRFQOrderNFT, signature []byte) (*types.Transaction, error) {
	return _OrderFilled.Contract.FillOrderNFT(&_OrderFilled.TransactOpts, order, signature)
}

// FillOrderWithTarget is a paid mutator transaction binding the contract method 0x00154008.
//
// Solidity: function fillOrderWithTarget((uint256,uint128,address,address,address,address,uint256,uint256) order, bytes signature, address target) returns()
func (_OrderFilled *OrderFilledTransactor) FillOrderWithTarget(opts *bind.TransactOpts, order AugustusRFQOrder, signature []byte, target common.Address) (*types.Transaction, error) {
	return _OrderFilled.contract.Transact(opts, "fillOrderWithTarget", order, signature, target)
}

// FillOrderWithTarget is a paid mutator transaction binding the contract method 0x00154008.
//
// Solidity: function fillOrderWithTarget((uint256,uint128,address,address,address,address,uint256,uint256) order, bytes signature, address target) returns()
func (_OrderFilled *OrderFilledSession) FillOrderWithTarget(order AugustusRFQOrder, signature []byte, target common.Address) (*types.Transaction, error) {
	return _OrderFilled.Contract.FillOrderWithTarget(&_OrderFilled.TransactOpts, order, signature, target)
}

// FillOrderWithTarget is a paid mutator transaction binding the contract method 0x00154008.
//
// Solidity: function fillOrderWithTarget((uint256,uint128,address,address,address,address,uint256,uint256) order, bytes signature, address target) returns()
func (_OrderFilled *OrderFilledTransactorSession) FillOrderWithTarget(order AugustusRFQOrder, signature []byte, target common.Address) (*types.Transaction, error) {
	return _OrderFilled.Contract.FillOrderWithTarget(&_OrderFilled.TransactOpts, order, signature, target)
}

// FillOrderWithTargetNFT is a paid mutator transaction binding the contract method 0x3c3694ab.
//
// Solidity: function fillOrderWithTargetNFT((uint256,uint128,uint256,uint256,uint256,uint256,address,address,uint256,uint256) order, bytes signature, address target) returns()
func (_OrderFilled *OrderFilledTransactor) FillOrderWithTargetNFT(opts *bind.TransactOpts, order AugustusRFQOrderNFT, signature []byte, target common.Address) (*types.Transaction, error) {
	return _OrderFilled.contract.Transact(opts, "fillOrderWithTargetNFT", order, signature, target)
}

// FillOrderWithTargetNFT is a paid mutator transaction binding the contract method 0x3c3694ab.
//
// Solidity: function fillOrderWithTargetNFT((uint256,uint128,uint256,uint256,uint256,uint256,address,address,uint256,uint256) order, bytes signature, address target) returns()
func (_OrderFilled *OrderFilledSession) FillOrderWithTargetNFT(order AugustusRFQOrderNFT, signature []byte, target common.Address) (*types.Transaction, error) {
	return _OrderFilled.Contract.FillOrderWithTargetNFT(&_OrderFilled.TransactOpts, order, signature, target)
}

// FillOrderWithTargetNFT is a paid mutator transaction binding the contract method 0x3c3694ab.
//
// Solidity: function fillOrderWithTargetNFT((uint256,uint128,uint256,uint256,uint256,uint256,address,address,uint256,uint256) order, bytes signature, address target) returns()
func (_OrderFilled *OrderFilledTransactorSession) FillOrderWithTargetNFT(order AugustusRFQOrderNFT, signature []byte, target common.Address) (*types.Transaction, error) {
	return _OrderFilled.Contract.FillOrderWithTargetNFT(&_OrderFilled.TransactOpts, order, signature, target)
}

// PartialFillOrder is a paid mutator transaction binding the contract method 0xc88ae6dc.
//
// Solidity: function partialFillOrder((uint256,uint128,address,address,address,address,uint256,uint256) order, bytes signature, uint256 takerTokenFillAmount) returns(uint256 makerTokenFilledAmount)
func (_OrderFilled *OrderFilledTransactor) PartialFillOrder(opts *bind.TransactOpts, order AugustusRFQOrder, signature []byte, takerTokenFillAmount *big.Int) (*types.Transaction, error) {
	return _OrderFilled.contract.Transact(opts, "partialFillOrder", order, signature, takerTokenFillAmount)
}

// PartialFillOrder is a paid mutator transaction binding the contract method 0xc88ae6dc.
//
// Solidity: function partialFillOrder((uint256,uint128,address,address,address,address,uint256,uint256) order, bytes signature, uint256 takerTokenFillAmount) returns(uint256 makerTokenFilledAmount)
func (_OrderFilled *OrderFilledSession) PartialFillOrder(order AugustusRFQOrder, signature []byte, takerTokenFillAmount *big.Int) (*types.Transaction, error) {
	return _OrderFilled.Contract.PartialFillOrder(&_OrderFilled.TransactOpts, order, signature, takerTokenFillAmount)
}

// PartialFillOrder is a paid mutator transaction binding the contract method 0xc88ae6dc.
//
// Solidity: function partialFillOrder((uint256,uint128,address,address,address,address,uint256,uint256) order, bytes signature, uint256 takerTokenFillAmount) returns(uint256 makerTokenFilledAmount)
func (_OrderFilled *OrderFilledTransactorSession) PartialFillOrder(order AugustusRFQOrder, signature []byte, takerTokenFillAmount *big.Int) (*types.Transaction, error) {
	return _OrderFilled.Contract.PartialFillOrder(&_OrderFilled.TransactOpts, order, signature, takerTokenFillAmount)
}

// PartialFillOrderNFT is a paid mutator transaction binding the contract method 0xb28ace5f.
//
// Solidity: function partialFillOrderNFT((uint256,uint128,uint256,uint256,uint256,uint256,address,address,uint256,uint256) order, bytes signature, uint256 takerTokenFillAmount) returns(uint256 makerTokenFilledAmount)
func (_OrderFilled *OrderFilledTransactor) PartialFillOrderNFT(opts *bind.TransactOpts, order AugustusRFQOrderNFT, signature []byte, takerTokenFillAmount *big.Int) (*types.Transaction, error) {
	return _OrderFilled.contract.Transact(opts, "partialFillOrderNFT", order, signature, takerTokenFillAmount)
}

// PartialFillOrderNFT is a paid mutator transaction binding the contract method 0xb28ace5f.
//
// Solidity: function partialFillOrderNFT((uint256,uint128,uint256,uint256,uint256,uint256,address,address,uint256,uint256) order, bytes signature, uint256 takerTokenFillAmount) returns(uint256 makerTokenFilledAmount)
func (_OrderFilled *OrderFilledSession) PartialFillOrderNFT(order AugustusRFQOrderNFT, signature []byte, takerTokenFillAmount *big.Int) (*types.Transaction, error) {
	return _OrderFilled.Contract.PartialFillOrderNFT(&_OrderFilled.TransactOpts, order, signature, takerTokenFillAmount)
}

// PartialFillOrderNFT is a paid mutator transaction binding the contract method 0xb28ace5f.
//
// Solidity: function partialFillOrderNFT((uint256,uint128,uint256,uint256,uint256,uint256,address,address,uint256,uint256) order, bytes signature, uint256 takerTokenFillAmount) returns(uint256 makerTokenFilledAmount)
func (_OrderFilled *OrderFilledTransactorSession) PartialFillOrderNFT(order AugustusRFQOrderNFT, signature []byte, takerTokenFillAmount *big.Int) (*types.Transaction, error) {
	return _OrderFilled.Contract.PartialFillOrderNFT(&_OrderFilled.TransactOpts, order, signature, takerTokenFillAmount)
}

// PartialFillOrderWithTarget is a paid mutator transaction binding the contract method 0x24abf828.
//
// Solidity: function partialFillOrderWithTarget((uint256,uint128,address,address,address,address,uint256,uint256) order, bytes signature, uint256 takerTokenFillAmount, address target) returns(uint256 makerTokenFilledAmount)
func (_OrderFilled *OrderFilledTransactor) PartialFillOrderWithTarget(opts *bind.TransactOpts, order AugustusRFQOrder, signature []byte, takerTokenFillAmount *big.Int, target common.Address) (*types.Transaction, error) {
	return _OrderFilled.contract.Transact(opts, "partialFillOrderWithTarget", order, signature, takerTokenFillAmount, target)
}

// PartialFillOrderWithTarget is a paid mutator transaction binding the contract method 0x24abf828.
//
// Solidity: function partialFillOrderWithTarget((uint256,uint128,address,address,address,address,uint256,uint256) order, bytes signature, uint256 takerTokenFillAmount, address target) returns(uint256 makerTokenFilledAmount)
func (_OrderFilled *OrderFilledSession) PartialFillOrderWithTarget(order AugustusRFQOrder, signature []byte, takerTokenFillAmount *big.Int, target common.Address) (*types.Transaction, error) {
	return _OrderFilled.Contract.PartialFillOrderWithTarget(&_OrderFilled.TransactOpts, order, signature, takerTokenFillAmount, target)
}

// PartialFillOrderWithTarget is a paid mutator transaction binding the contract method 0x24abf828.
//
// Solidity: function partialFillOrderWithTarget((uint256,uint128,address,address,address,address,uint256,uint256) order, bytes signature, uint256 takerTokenFillAmount, address target) returns(uint256 makerTokenFilledAmount)
func (_OrderFilled *OrderFilledTransactorSession) PartialFillOrderWithTarget(order AugustusRFQOrder, signature []byte, takerTokenFillAmount *big.Int, target common.Address) (*types.Transaction, error) {
	return _OrderFilled.Contract.PartialFillOrderWithTarget(&_OrderFilled.TransactOpts, order, signature, takerTokenFillAmount, target)
}

// PartialFillOrderWithTargetNFT is a paid mutator transaction binding the contract method 0x30201ad3.
//
// Solidity: function partialFillOrderWithTargetNFT((uint256,uint128,uint256,uint256,uint256,uint256,address,address,uint256,uint256) order, bytes signature, uint256 takerTokenFillAmount, address target) returns(uint256 makerTokenFilledAmount)
func (_OrderFilled *OrderFilledTransactor) PartialFillOrderWithTargetNFT(opts *bind.TransactOpts, order AugustusRFQOrderNFT, signature []byte, takerTokenFillAmount *big.Int, target common.Address) (*types.Transaction, error) {
	return _OrderFilled.contract.Transact(opts, "partialFillOrderWithTargetNFT", order, signature, takerTokenFillAmount, target)
}

// PartialFillOrderWithTargetNFT is a paid mutator transaction binding the contract method 0x30201ad3.
//
// Solidity: function partialFillOrderWithTargetNFT((uint256,uint128,uint256,uint256,uint256,uint256,address,address,uint256,uint256) order, bytes signature, uint256 takerTokenFillAmount, address target) returns(uint256 makerTokenFilledAmount)
func (_OrderFilled *OrderFilledSession) PartialFillOrderWithTargetNFT(order AugustusRFQOrderNFT, signature []byte, takerTokenFillAmount *big.Int, target common.Address) (*types.Transaction, error) {
	return _OrderFilled.Contract.PartialFillOrderWithTargetNFT(&_OrderFilled.TransactOpts, order, signature, takerTokenFillAmount, target)
}

// PartialFillOrderWithTargetNFT is a paid mutator transaction binding the contract method 0x30201ad3.
//
// Solidity: function partialFillOrderWithTargetNFT((uint256,uint128,uint256,uint256,uint256,uint256,address,address,uint256,uint256) order, bytes signature, uint256 takerTokenFillAmount, address target) returns(uint256 makerTokenFilledAmount)
func (_OrderFilled *OrderFilledTransactorSession) PartialFillOrderWithTargetNFT(order AugustusRFQOrderNFT, signature []byte, takerTokenFillAmount *big.Int, target common.Address) (*types.Transaction, error) {
	return _OrderFilled.Contract.PartialFillOrderWithTargetNFT(&_OrderFilled.TransactOpts, order, signature, takerTokenFillAmount, target)
}

// PartialFillOrderWithTargetPermit is a paid mutator transaction binding the contract method 0xda6b84af.
//
// Solidity: function partialFillOrderWithTargetPermit((uint256,uint128,address,address,address,address,uint256,uint256) order, bytes signature, uint256 takerTokenFillAmount, address target, bytes permitTakerAsset, bytes permitMakerAsset) returns(uint256 makerTokenFilledAmount)
func (_OrderFilled *OrderFilledTransactor) PartialFillOrderWithTargetPermit(opts *bind.TransactOpts, order AugustusRFQOrder, signature []byte, takerTokenFillAmount *big.Int, target common.Address, permitTakerAsset []byte, permitMakerAsset []byte) (*types.Transaction, error) {
	return _OrderFilled.contract.Transact(opts, "partialFillOrderWithTargetPermit", order, signature, takerTokenFillAmount, target, permitTakerAsset, permitMakerAsset)
}

// PartialFillOrderWithTargetPermit is a paid mutator transaction binding the contract method 0xda6b84af.
//
// Solidity: function partialFillOrderWithTargetPermit((uint256,uint128,address,address,address,address,uint256,uint256) order, bytes signature, uint256 takerTokenFillAmount, address target, bytes permitTakerAsset, bytes permitMakerAsset) returns(uint256 makerTokenFilledAmount)
func (_OrderFilled *OrderFilledSession) PartialFillOrderWithTargetPermit(order AugustusRFQOrder, signature []byte, takerTokenFillAmount *big.Int, target common.Address, permitTakerAsset []byte, permitMakerAsset []byte) (*types.Transaction, error) {
	return _OrderFilled.Contract.PartialFillOrderWithTargetPermit(&_OrderFilled.TransactOpts, order, signature, takerTokenFillAmount, target, permitTakerAsset, permitMakerAsset)
}

// PartialFillOrderWithTargetPermit is a paid mutator transaction binding the contract method 0xda6b84af.
//
// Solidity: function partialFillOrderWithTargetPermit((uint256,uint128,address,address,address,address,uint256,uint256) order, bytes signature, uint256 takerTokenFillAmount, address target, bytes permitTakerAsset, bytes permitMakerAsset) returns(uint256 makerTokenFilledAmount)
func (_OrderFilled *OrderFilledTransactorSession) PartialFillOrderWithTargetPermit(order AugustusRFQOrder, signature []byte, takerTokenFillAmount *big.Int, target common.Address, permitTakerAsset []byte, permitMakerAsset []byte) (*types.Transaction, error) {
	return _OrderFilled.Contract.PartialFillOrderWithTargetPermit(&_OrderFilled.TransactOpts, order, signature, takerTokenFillAmount, target, permitTakerAsset, permitMakerAsset)
}

// PartialFillOrderWithTargetPermitNFT is a paid mutator transaction binding the contract method 0xf6c1b371.
//
// Solidity: function partialFillOrderWithTargetPermitNFT((uint256,uint128,uint256,uint256,uint256,uint256,address,address,uint256,uint256) order, bytes signature, uint256 takerTokenFillAmount, address target, bytes permitTakerAsset, bytes permitMakerAsset) returns(uint256 makerTokenFilledAmount)
func (_OrderFilled *OrderFilledTransactor) PartialFillOrderWithTargetPermitNFT(opts *bind.TransactOpts, order AugustusRFQOrderNFT, signature []byte, takerTokenFillAmount *big.Int, target common.Address, permitTakerAsset []byte, permitMakerAsset []byte) (*types.Transaction, error) {
	return _OrderFilled.contract.Transact(opts, "partialFillOrderWithTargetPermitNFT", order, signature, takerTokenFillAmount, target, permitTakerAsset, permitMakerAsset)
}

// PartialFillOrderWithTargetPermitNFT is a paid mutator transaction binding the contract method 0xf6c1b371.
//
// Solidity: function partialFillOrderWithTargetPermitNFT((uint256,uint128,uint256,uint256,uint256,uint256,address,address,uint256,uint256) order, bytes signature, uint256 takerTokenFillAmount, address target, bytes permitTakerAsset, bytes permitMakerAsset) returns(uint256 makerTokenFilledAmount)
func (_OrderFilled *OrderFilledSession) PartialFillOrderWithTargetPermitNFT(order AugustusRFQOrderNFT, signature []byte, takerTokenFillAmount *big.Int, target common.Address, permitTakerAsset []byte, permitMakerAsset []byte) (*types.Transaction, error) {
	return _OrderFilled.Contract.PartialFillOrderWithTargetPermitNFT(&_OrderFilled.TransactOpts, order, signature, takerTokenFillAmount, target, permitTakerAsset, permitMakerAsset)
}

// PartialFillOrderWithTargetPermitNFT is a paid mutator transaction binding the contract method 0xf6c1b371.
//
// Solidity: function partialFillOrderWithTargetPermitNFT((uint256,uint128,uint256,uint256,uint256,uint256,address,address,uint256,uint256) order, bytes signature, uint256 takerTokenFillAmount, address target, bytes permitTakerAsset, bytes permitMakerAsset) returns(uint256 makerTokenFilledAmount)
func (_OrderFilled *OrderFilledTransactorSession) PartialFillOrderWithTargetPermitNFT(order AugustusRFQOrderNFT, signature []byte, takerTokenFillAmount *big.Int, target common.Address, permitTakerAsset []byte, permitMakerAsset []byte) (*types.Transaction, error) {
	return _OrderFilled.Contract.PartialFillOrderWithTargetPermitNFT(&_OrderFilled.TransactOpts, order, signature, takerTokenFillAmount, target, permitTakerAsset, permitMakerAsset)
}

// TryBatchFillOrderMakerAmount is a paid mutator transaction binding the contract method 0x01fb36ba.
//
// Solidity: function tryBatchFillOrderMakerAmount(((uint256,uint128,address,address,address,address,uint256,uint256),bytes,uint256,bytes,bytes)[] orderInfos, uint256 makerFillAmount, address target) returns()
func (_OrderFilled *OrderFilledTransactor) TryBatchFillOrderMakerAmount(opts *bind.TransactOpts, orderInfos []AugustusRFQOrderInfo, makerFillAmount *big.Int, target common.Address) (*types.Transaction, error) {
	return _OrderFilled.contract.Transact(opts, "tryBatchFillOrderMakerAmount", orderInfos, makerFillAmount, target)
}

// TryBatchFillOrderMakerAmount is a paid mutator transaction binding the contract method 0x01fb36ba.
//
// Solidity: function tryBatchFillOrderMakerAmount(((uint256,uint128,address,address,address,address,uint256,uint256),bytes,uint256,bytes,bytes)[] orderInfos, uint256 makerFillAmount, address target) returns()
func (_OrderFilled *OrderFilledSession) TryBatchFillOrderMakerAmount(orderInfos []AugustusRFQOrderInfo, makerFillAmount *big.Int, target common.Address) (*types.Transaction, error) {
	return _OrderFilled.Contract.TryBatchFillOrderMakerAmount(&_OrderFilled.TransactOpts, orderInfos, makerFillAmount, target)
}

// TryBatchFillOrderMakerAmount is a paid mutator transaction binding the contract method 0x01fb36ba.
//
// Solidity: function tryBatchFillOrderMakerAmount(((uint256,uint128,address,address,address,address,uint256,uint256),bytes,uint256,bytes,bytes)[] orderInfos, uint256 makerFillAmount, address target) returns()
func (_OrderFilled *OrderFilledTransactorSession) TryBatchFillOrderMakerAmount(orderInfos []AugustusRFQOrderInfo, makerFillAmount *big.Int, target common.Address) (*types.Transaction, error) {
	return _OrderFilled.Contract.TryBatchFillOrderMakerAmount(&_OrderFilled.TransactOpts, orderInfos, makerFillAmount, target)
}

// TryBatchFillOrderTakerAmount is a paid mutator transaction binding the contract method 0x1c64b820.
//
// Solidity: function tryBatchFillOrderTakerAmount(((uint256,uint128,address,address,address,address,uint256,uint256),bytes,uint256,bytes,bytes)[] orderInfos, uint256 takerFillAmount, address target) returns()
func (_OrderFilled *OrderFilledTransactor) TryBatchFillOrderTakerAmount(opts *bind.TransactOpts, orderInfos []AugustusRFQOrderInfo, takerFillAmount *big.Int, target common.Address) (*types.Transaction, error) {
	return _OrderFilled.contract.Transact(opts, "tryBatchFillOrderTakerAmount", orderInfos, takerFillAmount, target)
}

// TryBatchFillOrderTakerAmount is a paid mutator transaction binding the contract method 0x1c64b820.
//
// Solidity: function tryBatchFillOrderTakerAmount(((uint256,uint128,address,address,address,address,uint256,uint256),bytes,uint256,bytes,bytes)[] orderInfos, uint256 takerFillAmount, address target) returns()
func (_OrderFilled *OrderFilledSession) TryBatchFillOrderTakerAmount(orderInfos []AugustusRFQOrderInfo, takerFillAmount *big.Int, target common.Address) (*types.Transaction, error) {
	return _OrderFilled.Contract.TryBatchFillOrderTakerAmount(&_OrderFilled.TransactOpts, orderInfos, takerFillAmount, target)
}

// TryBatchFillOrderTakerAmount is a paid mutator transaction binding the contract method 0x1c64b820.
//
// Solidity: function tryBatchFillOrderTakerAmount(((uint256,uint128,address,address,address,address,uint256,uint256),bytes,uint256,bytes,bytes)[] orderInfos, uint256 takerFillAmount, address target) returns()
func (_OrderFilled *OrderFilledTransactorSession) TryBatchFillOrderTakerAmount(orderInfos []AugustusRFQOrderInfo, takerFillAmount *big.Int, target common.Address) (*types.Transaction, error) {
	return _OrderFilled.Contract.TryBatchFillOrderTakerAmount(&_OrderFilled.TransactOpts, orderInfos, takerFillAmount, target)
}

// OrderFilledOrderCancelledIterator is returned from FilterOrderCancelled and is used to iterate over the raw logs and unpacked data for OrderCancelled events raised by the OrderFilled contract.
type OrderFilledOrderCancelledIterator struct {
	Event *OrderFilledOrderCancelled // Event containing the contract specifics and raw log

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
func (it *OrderFilledOrderCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OrderFilledOrderCancelled)
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
		it.Event = new(OrderFilledOrderCancelled)
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
func (it *OrderFilledOrderCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OrderFilledOrderCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OrderFilledOrderCancelled represents a OrderCancelled event raised by the OrderFilled contract.
type OrderFilledOrderCancelled struct {
	OrderHash [32]byte
	Maker     common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOrderCancelled is a free log retrieval operation binding the contract event 0xa6eb7cdc219e1518ced964e9a34e61d68a94e4f1569db3e84256ba981ba52753.
//
// Solidity: event OrderCancelled(bytes32 indexed orderHash, address indexed maker)
func (_OrderFilled *OrderFilledFilterer) FilterOrderCancelled(opts *bind.FilterOpts, orderHash [][32]byte, maker []common.Address) (*OrderFilledOrderCancelledIterator, error) {

	var orderHashRule []interface{}
	for _, orderHashItem := range orderHash {
		orderHashRule = append(orderHashRule, orderHashItem)
	}
	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}

	logs, sub, err := _OrderFilled.contract.FilterLogs(opts, "OrderCancelled", orderHashRule, makerRule)
	if err != nil {
		return nil, err
	}
	return &OrderFilledOrderCancelledIterator{contract: _OrderFilled.contract, event: "OrderCancelled", logs: logs, sub: sub}, nil
}

// WatchOrderCancelled is a free log subscription operation binding the contract event 0xa6eb7cdc219e1518ced964e9a34e61d68a94e4f1569db3e84256ba981ba52753.
//
// Solidity: event OrderCancelled(bytes32 indexed orderHash, address indexed maker)
func (_OrderFilled *OrderFilledFilterer) WatchOrderCancelled(opts *bind.WatchOpts, sink chan<- *OrderFilledOrderCancelled, orderHash [][32]byte, maker []common.Address) (event.Subscription, error) {

	var orderHashRule []interface{}
	for _, orderHashItem := range orderHash {
		orderHashRule = append(orderHashRule, orderHashItem)
	}
	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}

	logs, sub, err := _OrderFilled.contract.WatchLogs(opts, "OrderCancelled", orderHashRule, makerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OrderFilledOrderCancelled)
				if err := _OrderFilled.contract.UnpackLog(event, "OrderCancelled", log); err != nil {
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

// ParseOrderCancelled is a log parse operation binding the contract event 0xa6eb7cdc219e1518ced964e9a34e61d68a94e4f1569db3e84256ba981ba52753.
//
// Solidity: event OrderCancelled(bytes32 indexed orderHash, address indexed maker)
func (_OrderFilled *OrderFilledFilterer) ParseOrderCancelled(log types.Log) (*OrderFilledOrderCancelled, error) {
	event := new(OrderFilledOrderCancelled)
	if err := _OrderFilled.contract.UnpackLog(event, "OrderCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OrderFilledOrderFilledIterator is returned from FilterOrderFilled and is used to iterate over the raw logs and unpacked data for OrderFilled events raised by the OrderFilled contract.
type OrderFilledOrderFilledIterator struct {
	Event *OrderFilledOrderFilled // Event containing the contract specifics and raw log

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
func (it *OrderFilledOrderFilledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OrderFilledOrderFilled)
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
		it.Event = new(OrderFilledOrderFilled)
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
func (it *OrderFilledOrderFilledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OrderFilledOrderFilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OrderFilledOrderFilled represents a OrderFilled event raised by the OrderFilled contract.
type OrderFilledOrderFilled struct {
	OrderHash   [32]byte
	Maker       common.Address
	MakerAsset  common.Address
	MakerAmount *big.Int
	Taker       common.Address
	TakerAsset  common.Address
	TakerAmount *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOrderFilled is a free log retrieval operation binding the contract event 0x6621486d9c28838df4a87d2cca5007bc2aaf6a5b5de083b1db8faf709302c473.
//
// Solidity: event OrderFilled(bytes32 indexed orderHash, address indexed maker, address makerAsset, uint256 makerAmount, address indexed taker, address takerAsset, uint256 takerAmount)
func (_OrderFilled *OrderFilledFilterer) FilterOrderFilled(opts *bind.FilterOpts, orderHash [][32]byte, maker []common.Address, taker []common.Address) (*OrderFilledOrderFilledIterator, error) {

	var orderHashRule []interface{}
	for _, orderHashItem := range orderHash {
		orderHashRule = append(orderHashRule, orderHashItem)
	}
	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}

	var takerRule []interface{}
	for _, takerItem := range taker {
		takerRule = append(takerRule, takerItem)
	}

	logs, sub, err := _OrderFilled.contract.FilterLogs(opts, "OrderFilled", orderHashRule, makerRule, takerRule)
	if err != nil {
		return nil, err
	}
	return &OrderFilledOrderFilledIterator{contract: _OrderFilled.contract, event: "OrderFilled", logs: logs, sub: sub}, nil
}

// WatchOrderFilled is a free log subscription operation binding the contract event 0x6621486d9c28838df4a87d2cca5007bc2aaf6a5b5de083b1db8faf709302c473.
//
// Solidity: event OrderFilled(bytes32 indexed orderHash, address indexed maker, address makerAsset, uint256 makerAmount, address indexed taker, address takerAsset, uint256 takerAmount)
func (_OrderFilled *OrderFilledFilterer) WatchOrderFilled(opts *bind.WatchOpts, sink chan<- *OrderFilledOrderFilled, orderHash [][32]byte, maker []common.Address, taker []common.Address) (event.Subscription, error) {

	var orderHashRule []interface{}
	for _, orderHashItem := range orderHash {
		orderHashRule = append(orderHashRule, orderHashItem)
	}
	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}

	var takerRule []interface{}
	for _, takerItem := range taker {
		takerRule = append(takerRule, takerItem)
	}

	logs, sub, err := _OrderFilled.contract.WatchLogs(opts, "OrderFilled", orderHashRule, makerRule, takerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OrderFilledOrderFilled)
				if err := _OrderFilled.contract.UnpackLog(event, "OrderFilled", log); err != nil {
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

// ParseOrderFilled is a log parse operation binding the contract event 0x6621486d9c28838df4a87d2cca5007bc2aaf6a5b5de083b1db8faf709302c473.
//
// Solidity: event OrderFilled(bytes32 indexed orderHash, address indexed maker, address makerAsset, uint256 makerAmount, address indexed taker, address takerAsset, uint256 takerAmount)
func (_OrderFilled *OrderFilledFilterer) ParseOrderFilled(log types.Log) (*OrderFilledOrderFilled, error) {
	event := new(OrderFilledOrderFilled)
	if err := _OrderFilled.contract.UnpackLog(event, "OrderFilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OrderFilledOrderFilledNFTIterator is returned from FilterOrderFilledNFT and is used to iterate over the raw logs and unpacked data for OrderFilledNFT events raised by the OrderFilled contract.
type OrderFilledOrderFilledNFTIterator struct {
	Event *OrderFilledOrderFilledNFT // Event containing the contract specifics and raw log

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
func (it *OrderFilledOrderFilledNFTIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OrderFilledOrderFilledNFT)
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
		it.Event = new(OrderFilledOrderFilledNFT)
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
func (it *OrderFilledOrderFilledNFTIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OrderFilledOrderFilledNFTIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OrderFilledOrderFilledNFT represents a OrderFilledNFT event raised by the OrderFilled contract.
type OrderFilledOrderFilledNFT struct {
	OrderHash    [32]byte
	Maker        common.Address
	MakerAsset   *big.Int
	MakerAssetId *big.Int
	MakerAmount  *big.Int
	Taker        common.Address
	TakerAsset   *big.Int
	TakerAssetId *big.Int
	TakerAmount  *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterOrderFilledNFT is a free log retrieval operation binding the contract event 0x58454165245cb04f65f0d8e8e867125cee46b1b301053018898247b85cb4a9bc.
//
// Solidity: event OrderFilledNFT(bytes32 indexed orderHash, address indexed maker, uint256 makerAsset, uint256 makerAssetId, uint256 makerAmount, address indexed taker, uint256 takerAsset, uint256 takerAssetId, uint256 takerAmount)
func (_OrderFilled *OrderFilledFilterer) FilterOrderFilledNFT(opts *bind.FilterOpts, orderHash [][32]byte, maker []common.Address, taker []common.Address) (*OrderFilledOrderFilledNFTIterator, error) {

	var orderHashRule []interface{}
	for _, orderHashItem := range orderHash {
		orderHashRule = append(orderHashRule, orderHashItem)
	}
	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}

	var takerRule []interface{}
	for _, takerItem := range taker {
		takerRule = append(takerRule, takerItem)
	}

	logs, sub, err := _OrderFilled.contract.FilterLogs(opts, "OrderFilledNFT", orderHashRule, makerRule, takerRule)
	if err != nil {
		return nil, err
	}
	return &OrderFilledOrderFilledNFTIterator{contract: _OrderFilled.contract, event: "OrderFilledNFT", logs: logs, sub: sub}, nil
}

// WatchOrderFilledNFT is a free log subscription operation binding the contract event 0x58454165245cb04f65f0d8e8e867125cee46b1b301053018898247b85cb4a9bc.
//
// Solidity: event OrderFilledNFT(bytes32 indexed orderHash, address indexed maker, uint256 makerAsset, uint256 makerAssetId, uint256 makerAmount, address indexed taker, uint256 takerAsset, uint256 takerAssetId, uint256 takerAmount)
func (_OrderFilled *OrderFilledFilterer) WatchOrderFilledNFT(opts *bind.WatchOpts, sink chan<- *OrderFilledOrderFilledNFT, orderHash [][32]byte, maker []common.Address, taker []common.Address) (event.Subscription, error) {

	var orderHashRule []interface{}
	for _, orderHashItem := range orderHash {
		orderHashRule = append(orderHashRule, orderHashItem)
	}
	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}

	var takerRule []interface{}
	for _, takerItem := range taker {
		takerRule = append(takerRule, takerItem)
	}

	logs, sub, err := _OrderFilled.contract.WatchLogs(opts, "OrderFilledNFT", orderHashRule, makerRule, takerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OrderFilledOrderFilledNFT)
				if err := _OrderFilled.contract.UnpackLog(event, "OrderFilledNFT", log); err != nil {
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

// ParseOrderFilledNFT is a log parse operation binding the contract event 0x58454165245cb04f65f0d8e8e867125cee46b1b301053018898247b85cb4a9bc.
//
// Solidity: event OrderFilledNFT(bytes32 indexed orderHash, address indexed maker, uint256 makerAsset, uint256 makerAssetId, uint256 makerAmount, address indexed taker, uint256 takerAsset, uint256 takerAssetId, uint256 takerAmount)
func (_OrderFilled *OrderFilledFilterer) ParseOrderFilledNFT(log types.Log) (*OrderFilledOrderFilledNFT, error) {
	event := new(OrderFilledOrderFilledNFT)
	if err := _OrderFilled.contract.UnpackLog(event, "OrderFilledNFT", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
