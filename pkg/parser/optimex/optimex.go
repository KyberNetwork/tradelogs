// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package optimex

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

// ITypesAffiliate is an auto generated low-level Go binding around an user-defined struct.
type ITypesAffiliate struct {
	AggregatedValue *big.Int
	Schema          string
	Data            []byte
}

// ITypesBundlePayment is an auto generated low-level Go binding around an user-defined struct.
type ITypesBundlePayment struct {
	TradeIds    [][32]byte
	SignedAt    uint64
	StartIdx    uint64
	PaymentTxId []byte
	Signature   []byte
}

// ITypesFeeDetails is an auto generated low-level Go binding around an user-defined struct.
type ITypesFeeDetails struct {
	TotalAmount *big.Int
	PFeeAmount  *big.Int
	AFeeAmount  *big.Int
	PFeeRate    *big.Int
	AFeeRate    *big.Int
}

// ITypesMPCInfo is an auto generated low-level Go binding around an user-defined struct.
type ITypesMPCInfo struct {
	MpcL2Address   common.Address
	ExpireTime     uint64
	MpcL2Pubkey    []byte
	MpcAssetPubkey []byte
}

// ITypesPMMSelection is an auto generated low-level Go binding around an user-defined struct.
type ITypesPMMSelection struct {
	RfqInfo ITypesRFQInfo
	PmmInfo ITypesSelectedPMMInfo
}

// ITypesPresign is an auto generated low-level Go binding around an user-defined struct.
type ITypesPresign struct {
	PmmId          [32]byte
	PmmRecvAddress []byte
	Presigns       [][]byte
}

// ITypesRFQInfo is an auto generated low-level Go binding around an user-defined struct.
type ITypesRFQInfo struct {
	MinAmountOut     *big.Int
	TradeTimeout     uint64
	RfqInfoSignature []byte
}

// ITypesScriptInfo is an auto generated low-level Go binding around an user-defined struct.
type ITypesScriptInfo struct {
	DepositInfo            [5][]byte
	UserEphemeralL2Address common.Address
	ScriptTimeout          uint64
}

// ITypesSelectedPMMInfo is an auto generated low-level Go binding around an user-defined struct.
type ITypesSelectedPMMInfo struct {
	AmountOut     *big.Int
	SelectedPMMId [32]byte
	Info          [2][]byte
	SigExpiry     uint64
}

// ITypesSettledPayment is an auto generated low-level Go binding around an user-defined struct.
type ITypesSettledPayment struct {
	BundlerHash [32]byte
	PaymentTxId []byte
	ReleaseTxId []byte
	IsConfirmed bool
}

// ITypesTokenInfo is an auto generated low-level Go binding around an user-defined struct.
type ITypesTokenInfo struct {
	Info     [5][]byte
	Decimals *big.Int
}

// ITypesTradeData is an auto generated low-level Go binding around an user-defined struct.
type ITypesTradeData struct {
	SessionId  *big.Int
	TradeInfo  ITypesTradeInfo
	ScriptInfo ITypesScriptInfo
}

// ITypesTradeInfo is an auto generated low-level Go binding around an user-defined struct.
type ITypesTradeInfo struct {
	AmountIn  *big.Int
	FromChain [3][]byte
	ToChain   [3][]byte
}

// OptimexMetaData contains all meta data concerning the Optimex contract.
var OptimexMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIManagement\",\"name\":\"management_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AddressZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BundlePaymentEmpty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InconsistentCoreType\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InconsistentPMM\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RegisteredAlready\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RouteNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RouteNotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"mpc\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"tradeId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pFeeRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"aFeeRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"list\",\"type\":\"bytes[]\"}],\"name\":\"ConfirmDeposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"mpc\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"tradeId\",\"type\":\"bytes32\"}],\"name\":\"ConfirmPayment\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"mpc\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"tradeId\",\"type\":\"bytes32\"}],\"name\":\"ConfirmSettlement\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"tradeId\",\"type\":\"bytes32\"}],\"name\":\"MakePayment\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"solver\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"tradeId\",\"type\":\"bytes32\"}],\"name\":\"SelectPMM\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"solver\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"tradeId\",\"type\":\"bytes32\"}],\"name\":\"SubmitTradeInfo\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"core\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"fromChain\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"toChain\",\"type\":\"bytes\"}],\"name\":\"UpdatedRoute\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"SIGNER\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32[]\",\"name\":\"tradeIds\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint64\",\"name\":\"signedAt\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"startIdx\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"paymentTxId\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structITypes.BundlePayment\",\"name\":\"bundle\",\"type\":\"tuple\"}],\"name\":\"bundlePayment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"tradeId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"depositFromList\",\"type\":\"bytes[]\"}],\"name\":\"confirmDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"tradeId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"confirmPayment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"tradeId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"releaseTxId\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"confirmSettlement\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"tradeId\",\"type\":\"bytes32\"}],\"name\":\"getAffiliateInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"aggregatedValue\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"schema\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structITypes.Affiliate\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"tradeId\",\"type\":\"bytes32\"}],\"name\":\"getCurrentStage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"tradeId\",\"type\":\"bytes32\"}],\"name\":\"getDepositAddressList\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"\",\"type\":\"bytes[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"tradeId\",\"type\":\"bytes32\"}],\"name\":\"getFeeDetails\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"totalAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pFeeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"aFeeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"pFeeRate\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"aFeeRate\",\"type\":\"uint128\"}],\"internalType\":\"structITypes.FeeDetails\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"fromChain\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"toChain\",\"type\":\"bytes\"}],\"name\":\"getHandler\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"handler\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"handlerType\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"tradeId\",\"type\":\"bytes32\"}],\"name\":\"getHandlerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"handler\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"handlerType\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"tradeId\",\"type\":\"bytes32\"}],\"name\":\"getLastSignedPayment\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"networkId\",\"type\":\"bytes\"}],\"name\":\"getLatestMPCInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"mpcL2Address\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"expireTime\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"mpcL2Pubkey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"mpcAssetPubkey\",\"type\":\"bytes\"}],\"internalType\":\"structITypes.MPCInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"networkId\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"}],\"name\":\"getMPCInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"mpcL2Address\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"expireTime\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"mpcL2Pubkey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"mpcAssetPubkey\",\"type\":\"bytes\"}],\"internalType\":\"structITypes.MPCInfo\",\"name\":\"info\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getManagementOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"fromChain\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"toChain\",\"type\":\"bytes\"}],\"name\":\"getMaxAffiliateFeeRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPFeeRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"pmmId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"fromIdx\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toIdx\",\"type\":\"uint256\"}],\"name\":\"getPMMAccounts\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"list\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"tradeId\",\"type\":\"bytes32\"}],\"name\":\"getPMMSelection\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"minAmountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"tradeTimeout\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"rfqInfoSignature\",\"type\":\"bytes\"}],\"internalType\":\"structITypes.RFQInfo\",\"name\":\"rfqInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"selectedPMMId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes[2]\",\"name\":\"info\",\"type\":\"bytes[2]\"},{\"internalType\":\"uint64\",\"name\":\"sigExpiry\",\"type\":\"uint64\"}],\"internalType\":\"structITypes.SelectedPMMInfo\",\"name\":\"pmmInfo\",\"type\":\"tuple\"}],\"internalType\":\"structITypes.PMMSelection\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"tradeId\",\"type\":\"bytes32\"}],\"name\":\"getPresigns\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"pmmId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"pmmRecvAddress\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"presigns\",\"type\":\"bytes[]\"}],\"internalType\":\"structITypes.Presign[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getProtocolState\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"tradeId\",\"type\":\"bytes32\"}],\"name\":\"getSettledPayment\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"bundlerHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"paymentTxId\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"releaseTxId\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"isConfirmed\",\"type\":\"bool\"}],\"internalType\":\"structITypes.SettledPayment\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fromIdx\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toIdx\",\"type\":\"uint256\"}],\"name\":\"getTokens\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes[5]\",\"name\":\"info\",\"type\":\"bytes[5]\"},{\"internalType\":\"uint256\",\"name\":\"decimals\",\"type\":\"uint256\"}],\"internalType\":\"structITypes.TokenInfo[]\",\"name\":\"list\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"tradeId\",\"type\":\"bytes32\"}],\"name\":\"getTradeData\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sessionId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"bytes[3]\",\"name\":\"fromChain\",\"type\":\"bytes[3]\"},{\"internalType\":\"bytes[3]\",\"name\":\"toChain\",\"type\":\"bytes[3]\"}],\"internalType\":\"structITypes.TradeInfo\",\"name\":\"tradeInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes[5]\",\"name\":\"depositInfo\",\"type\":\"bytes[5]\"},{\"internalType\":\"address\",\"name\":\"userEphemeralL2Address\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"scriptTimeout\",\"type\":\"uint64\"}],\"internalType\":\"structITypes.ScriptInfo\",\"name\":\"scriptInfo\",\"type\":\"tuple\"}],\"internalType\":\"structITypes.TradeData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isMPCNode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isSolver\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumITypes.STAGE\",\"name\":\"stage\",\"type\":\"uint8\"}],\"name\":\"isSuspended\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"networkId\",\"type\":\"bytes\"}],\"name\":\"isValidNetwork\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"pmmId\",\"type\":\"bytes32\"}],\"name\":\"isValidPMM\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"pmmId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isValidPMMAccount\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"networkId\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"}],\"name\":\"isValidPubkey\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"networkId\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"tokenId\",\"type\":\"bytes\"}],\"name\":\"isValidToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"management\",\"outputs\":[{\"internalType\":\"contractIManagement\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"pmmId\",\"type\":\"bytes32\"}],\"name\":\"numOfPMMAccounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numOfSupportedTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"tradeId\",\"type\":\"bytes32\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"minAmountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"tradeTimeout\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"rfqInfoSignature\",\"type\":\"bytes\"}],\"internalType\":\"structITypes.RFQInfo\",\"name\":\"rfqInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"selectedPMMId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes[2]\",\"name\":\"info\",\"type\":\"bytes[2]\"},{\"internalType\":\"uint64\",\"name\":\"sigExpiry\",\"type\":\"uint64\"}],\"internalType\":\"structITypes.SelectedPMMInfo\",\"name\":\"pmmInfo\",\"type\":\"tuple\"}],\"internalType\":\"structITypes.PMMSelection\",\"name\":\"info\",\"type\":\"tuple\"}],\"name\":\"selectPMM\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newManagement\",\"type\":\"address\"}],\"name\":\"setManagement\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"core\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"fromChain\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"toChain\",\"type\":\"bytes\"}],\"name\":\"setRoute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"tradeId\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sessionId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"bytes[3]\",\"name\":\"fromChain\",\"type\":\"bytes[3]\"},{\"internalType\":\"bytes[3]\",\"name\":\"toChain\",\"type\":\"bytes[3]\"}],\"internalType\":\"structITypes.TradeInfo\",\"name\":\"tradeInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes[5]\",\"name\":\"depositInfo\",\"type\":\"bytes[5]\"},{\"internalType\":\"address\",\"name\":\"userEphemeralL2Address\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"scriptTimeout\",\"type\":\"uint64\"}],\"internalType\":\"structITypes.ScriptInfo\",\"name\":\"scriptInfo\",\"type\":\"tuple\"}],\"internalType\":\"structITypes.TradeData\",\"name\":\"tradeData\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"aggregatedValue\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"schema\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structITypes.Affiliate\",\"name\":\"affiliateInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"pmmId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"pmmRecvAddress\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"presigns\",\"type\":\"bytes[]\"}],\"internalType\":\"structITypes.Presign[]\",\"name\":\"presignList\",\"type\":\"tuple[]\"}],\"name\":\"submitTrade\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// OptimexABI is the input ABI used to generate the binding from.
// Deprecated: Use OptimexMetaData.ABI instead.
var OptimexABI = OptimexMetaData.ABI

// Optimex is an auto generated Go binding around an Ethereum contract.
type Optimex struct {
	OptimexCaller     // Read-only binding to the contract
	OptimexTransactor // Write-only binding to the contract
	OptimexFilterer   // Log filterer for contract events
}

// OptimexCaller is an auto generated read-only Go binding around an Ethereum contract.
type OptimexCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OptimexTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OptimexTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OptimexFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OptimexFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OptimexSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OptimexSession struct {
	Contract     *Optimex          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OptimexCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OptimexCallerSession struct {
	Contract *OptimexCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// OptimexTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OptimexTransactorSession struct {
	Contract     *OptimexTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// OptimexRaw is an auto generated low-level Go binding around an Ethereum contract.
type OptimexRaw struct {
	Contract *Optimex // Generic contract binding to access the raw methods on
}

// OptimexCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OptimexCallerRaw struct {
	Contract *OptimexCaller // Generic read-only contract binding to access the raw methods on
}

// OptimexTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OptimexTransactorRaw struct {
	Contract *OptimexTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOptimex creates a new instance of Optimex, bound to a specific deployed contract.
func NewOptimex(address common.Address, backend bind.ContractBackend) (*Optimex, error) {
	contract, err := bindOptimex(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Optimex{OptimexCaller: OptimexCaller{contract: contract}, OptimexTransactor: OptimexTransactor{contract: contract}, OptimexFilterer: OptimexFilterer{contract: contract}}, nil
}

// NewOptimexCaller creates a new read-only instance of Optimex, bound to a specific deployed contract.
func NewOptimexCaller(address common.Address, caller bind.ContractCaller) (*OptimexCaller, error) {
	contract, err := bindOptimex(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OptimexCaller{contract: contract}, nil
}

// NewOptimexTransactor creates a new write-only instance of Optimex, bound to a specific deployed contract.
func NewOptimexTransactor(address common.Address, transactor bind.ContractTransactor) (*OptimexTransactor, error) {
	contract, err := bindOptimex(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OptimexTransactor{contract: contract}, nil
}

// NewOptimexFilterer creates a new log filterer instance of Optimex, bound to a specific deployed contract.
func NewOptimexFilterer(address common.Address, filterer bind.ContractFilterer) (*OptimexFilterer, error) {
	contract, err := bindOptimex(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OptimexFilterer{contract: contract}, nil
}

// bindOptimex binds a generic wrapper to an already deployed contract.
func bindOptimex(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OptimexMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Optimex *OptimexRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Optimex.Contract.OptimexCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Optimex *OptimexRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Optimex.Contract.OptimexTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Optimex *OptimexRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Optimex.Contract.OptimexTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Optimex *OptimexCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Optimex.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Optimex *OptimexTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Optimex.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Optimex *OptimexTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Optimex.Contract.contract.Transact(opts, method, params...)
}

// SIGNER is a free data retrieval call binding the contract method 0x582abd12.
//
// Solidity: function SIGNER() view returns(address)
func (_Optimex *OptimexCaller) SIGNER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Optimex.contract.Call(opts, &out, "SIGNER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SIGNER is a free data retrieval call binding the contract method 0x582abd12.
//
// Solidity: function SIGNER() view returns(address)
func (_Optimex *OptimexSession) SIGNER() (common.Address, error) {
	return _Optimex.Contract.SIGNER(&_Optimex.CallOpts)
}

// SIGNER is a free data retrieval call binding the contract method 0x582abd12.
//
// Solidity: function SIGNER() view returns(address)
func (_Optimex *OptimexCallerSession) SIGNER() (common.Address, error) {
	return _Optimex.Contract.SIGNER(&_Optimex.CallOpts)
}

// GetAffiliateInfo is a free data retrieval call binding the contract method 0x30644e3d.
//
// Solidity: function getAffiliateInfo(bytes32 tradeId) view returns((uint256,string,bytes))
func (_Optimex *OptimexCaller) GetAffiliateInfo(opts *bind.CallOpts, tradeId [32]byte) (ITypesAffiliate, error) {
	var out []interface{}
	err := _Optimex.contract.Call(opts, &out, "getAffiliateInfo", tradeId)

	if err != nil {
		return *new(ITypesAffiliate), err
	}

	out0 := *abi.ConvertType(out[0], new(ITypesAffiliate)).(*ITypesAffiliate)

	return out0, err

}

// GetAffiliateInfo is a free data retrieval call binding the contract method 0x30644e3d.
//
// Solidity: function getAffiliateInfo(bytes32 tradeId) view returns((uint256,string,bytes))
func (_Optimex *OptimexSession) GetAffiliateInfo(tradeId [32]byte) (ITypesAffiliate, error) {
	return _Optimex.Contract.GetAffiliateInfo(&_Optimex.CallOpts, tradeId)
}

// GetAffiliateInfo is a free data retrieval call binding the contract method 0x30644e3d.
//
// Solidity: function getAffiliateInfo(bytes32 tradeId) view returns((uint256,string,bytes))
func (_Optimex *OptimexCallerSession) GetAffiliateInfo(tradeId [32]byte) (ITypesAffiliate, error) {
	return _Optimex.Contract.GetAffiliateInfo(&_Optimex.CallOpts, tradeId)
}

// GetCurrentStage is a free data retrieval call binding the contract method 0x10178310.
//
// Solidity: function getCurrentStage(bytes32 tradeId) view returns(uint256)
func (_Optimex *OptimexCaller) GetCurrentStage(opts *bind.CallOpts, tradeId [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Optimex.contract.Call(opts, &out, "getCurrentStage", tradeId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentStage is a free data retrieval call binding the contract method 0x10178310.
//
// Solidity: function getCurrentStage(bytes32 tradeId) view returns(uint256)
func (_Optimex *OptimexSession) GetCurrentStage(tradeId [32]byte) (*big.Int, error) {
	return _Optimex.Contract.GetCurrentStage(&_Optimex.CallOpts, tradeId)
}

// GetCurrentStage is a free data retrieval call binding the contract method 0x10178310.
//
// Solidity: function getCurrentStage(bytes32 tradeId) view returns(uint256)
func (_Optimex *OptimexCallerSession) GetCurrentStage(tradeId [32]byte) (*big.Int, error) {
	return _Optimex.Contract.GetCurrentStage(&_Optimex.CallOpts, tradeId)
}

// GetDepositAddressList is a free data retrieval call binding the contract method 0xae4637a8.
//
// Solidity: function getDepositAddressList(bytes32 tradeId) view returns(bytes[])
func (_Optimex *OptimexCaller) GetDepositAddressList(opts *bind.CallOpts, tradeId [32]byte) ([][]byte, error) {
	var out []interface{}
	err := _Optimex.contract.Call(opts, &out, "getDepositAddressList", tradeId)

	if err != nil {
		return *new([][]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][]byte)).(*[][]byte)

	return out0, err

}

// GetDepositAddressList is a free data retrieval call binding the contract method 0xae4637a8.
//
// Solidity: function getDepositAddressList(bytes32 tradeId) view returns(bytes[])
func (_Optimex *OptimexSession) GetDepositAddressList(tradeId [32]byte) ([][]byte, error) {
	return _Optimex.Contract.GetDepositAddressList(&_Optimex.CallOpts, tradeId)
}

// GetDepositAddressList is a free data retrieval call binding the contract method 0xae4637a8.
//
// Solidity: function getDepositAddressList(bytes32 tradeId) view returns(bytes[])
func (_Optimex *OptimexCallerSession) GetDepositAddressList(tradeId [32]byte) ([][]byte, error) {
	return _Optimex.Contract.GetDepositAddressList(&_Optimex.CallOpts, tradeId)
}

// GetFeeDetails is a free data retrieval call binding the contract method 0x9108e266.
//
// Solidity: function getFeeDetails(bytes32 tradeId) view returns((uint256,uint256,uint256,uint128,uint128))
func (_Optimex *OptimexCaller) GetFeeDetails(opts *bind.CallOpts, tradeId [32]byte) (ITypesFeeDetails, error) {
	var out []interface{}
	err := _Optimex.contract.Call(opts, &out, "getFeeDetails", tradeId)

	if err != nil {
		return *new(ITypesFeeDetails), err
	}

	out0 := *abi.ConvertType(out[0], new(ITypesFeeDetails)).(*ITypesFeeDetails)

	return out0, err

}

// GetFeeDetails is a free data retrieval call binding the contract method 0x9108e266.
//
// Solidity: function getFeeDetails(bytes32 tradeId) view returns((uint256,uint256,uint256,uint128,uint128))
func (_Optimex *OptimexSession) GetFeeDetails(tradeId [32]byte) (ITypesFeeDetails, error) {
	return _Optimex.Contract.GetFeeDetails(&_Optimex.CallOpts, tradeId)
}

// GetFeeDetails is a free data retrieval call binding the contract method 0x9108e266.
//
// Solidity: function getFeeDetails(bytes32 tradeId) view returns((uint256,uint256,uint256,uint128,uint128))
func (_Optimex *OptimexCallerSession) GetFeeDetails(tradeId [32]byte) (ITypesFeeDetails, error) {
	return _Optimex.Contract.GetFeeDetails(&_Optimex.CallOpts, tradeId)
}

// GetHandler is a free data retrieval call binding the contract method 0x96c3b5a0.
//
// Solidity: function getHandler(bytes fromChain, bytes toChain) view returns(address handler, string handlerType)
func (_Optimex *OptimexCaller) GetHandler(opts *bind.CallOpts, fromChain []byte, toChain []byte) (struct {
	Handler     common.Address
	HandlerType string
}, error) {
	var out []interface{}
	err := _Optimex.contract.Call(opts, &out, "getHandler", fromChain, toChain)

	outstruct := new(struct {
		Handler     common.Address
		HandlerType string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Handler = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.HandlerType = *abi.ConvertType(out[1], new(string)).(*string)

	return *outstruct, err

}

// GetHandler is a free data retrieval call binding the contract method 0x96c3b5a0.
//
// Solidity: function getHandler(bytes fromChain, bytes toChain) view returns(address handler, string handlerType)
func (_Optimex *OptimexSession) GetHandler(fromChain []byte, toChain []byte) (struct {
	Handler     common.Address
	HandlerType string
}, error) {
	return _Optimex.Contract.GetHandler(&_Optimex.CallOpts, fromChain, toChain)
}

// GetHandler is a free data retrieval call binding the contract method 0x96c3b5a0.
//
// Solidity: function getHandler(bytes fromChain, bytes toChain) view returns(address handler, string handlerType)
func (_Optimex *OptimexCallerSession) GetHandler(fromChain []byte, toChain []byte) (struct {
	Handler     common.Address
	HandlerType string
}, error) {
	return _Optimex.Contract.GetHandler(&_Optimex.CallOpts, fromChain, toChain)
}

// GetHandlerOf is a free data retrieval call binding the contract method 0x84aadc97.
//
// Solidity: function getHandlerOf(bytes32 tradeId) view returns(address handler, string handlerType)
func (_Optimex *OptimexCaller) GetHandlerOf(opts *bind.CallOpts, tradeId [32]byte) (struct {
	Handler     common.Address
	HandlerType string
}, error) {
	var out []interface{}
	err := _Optimex.contract.Call(opts, &out, "getHandlerOf", tradeId)

	outstruct := new(struct {
		Handler     common.Address
		HandlerType string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Handler = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.HandlerType = *abi.ConvertType(out[1], new(string)).(*string)

	return *outstruct, err

}

// GetHandlerOf is a free data retrieval call binding the contract method 0x84aadc97.
//
// Solidity: function getHandlerOf(bytes32 tradeId) view returns(address handler, string handlerType)
func (_Optimex *OptimexSession) GetHandlerOf(tradeId [32]byte) (struct {
	Handler     common.Address
	HandlerType string
}, error) {
	return _Optimex.Contract.GetHandlerOf(&_Optimex.CallOpts, tradeId)
}

// GetHandlerOf is a free data retrieval call binding the contract method 0x84aadc97.
//
// Solidity: function getHandlerOf(bytes32 tradeId) view returns(address handler, string handlerType)
func (_Optimex *OptimexCallerSession) GetHandlerOf(tradeId [32]byte) (struct {
	Handler     common.Address
	HandlerType string
}, error) {
	return _Optimex.Contract.GetHandlerOf(&_Optimex.CallOpts, tradeId)
}

// GetLastSignedPayment is a free data retrieval call binding the contract method 0x140bf2a6.
//
// Solidity: function getLastSignedPayment(bytes32 tradeId) view returns(uint64)
func (_Optimex *OptimexCaller) GetLastSignedPayment(opts *bind.CallOpts, tradeId [32]byte) (uint64, error) {
	var out []interface{}
	err := _Optimex.contract.Call(opts, &out, "getLastSignedPayment", tradeId)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetLastSignedPayment is a free data retrieval call binding the contract method 0x140bf2a6.
//
// Solidity: function getLastSignedPayment(bytes32 tradeId) view returns(uint64)
func (_Optimex *OptimexSession) GetLastSignedPayment(tradeId [32]byte) (uint64, error) {
	return _Optimex.Contract.GetLastSignedPayment(&_Optimex.CallOpts, tradeId)
}

// GetLastSignedPayment is a free data retrieval call binding the contract method 0x140bf2a6.
//
// Solidity: function getLastSignedPayment(bytes32 tradeId) view returns(uint64)
func (_Optimex *OptimexCallerSession) GetLastSignedPayment(tradeId [32]byte) (uint64, error) {
	return _Optimex.Contract.GetLastSignedPayment(&_Optimex.CallOpts, tradeId)
}

// GetLatestMPCInfo is a free data retrieval call binding the contract method 0xe18dc3de.
//
// Solidity: function getLatestMPCInfo(bytes networkId) view returns((address,uint64,bytes,bytes))
func (_Optimex *OptimexCaller) GetLatestMPCInfo(opts *bind.CallOpts, networkId []byte) (ITypesMPCInfo, error) {
	var out []interface{}
	err := _Optimex.contract.Call(opts, &out, "getLatestMPCInfo", networkId)

	if err != nil {
		return *new(ITypesMPCInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(ITypesMPCInfo)).(*ITypesMPCInfo)

	return out0, err

}

// GetLatestMPCInfo is a free data retrieval call binding the contract method 0xe18dc3de.
//
// Solidity: function getLatestMPCInfo(bytes networkId) view returns((address,uint64,bytes,bytes))
func (_Optimex *OptimexSession) GetLatestMPCInfo(networkId []byte) (ITypesMPCInfo, error) {
	return _Optimex.Contract.GetLatestMPCInfo(&_Optimex.CallOpts, networkId)
}

// GetLatestMPCInfo is a free data retrieval call binding the contract method 0xe18dc3de.
//
// Solidity: function getLatestMPCInfo(bytes networkId) view returns((address,uint64,bytes,bytes))
func (_Optimex *OptimexCallerSession) GetLatestMPCInfo(networkId []byte) (ITypesMPCInfo, error) {
	return _Optimex.Contract.GetLatestMPCInfo(&_Optimex.CallOpts, networkId)
}

// GetMPCInfo is a free data retrieval call binding the contract method 0x69aed5b5.
//
// Solidity: function getMPCInfo(bytes networkId, bytes pubkey) view returns((address,uint64,bytes,bytes) info)
func (_Optimex *OptimexCaller) GetMPCInfo(opts *bind.CallOpts, networkId []byte, pubkey []byte) (ITypesMPCInfo, error) {
	var out []interface{}
	err := _Optimex.contract.Call(opts, &out, "getMPCInfo", networkId, pubkey)

	if err != nil {
		return *new(ITypesMPCInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(ITypesMPCInfo)).(*ITypesMPCInfo)

	return out0, err

}

// GetMPCInfo is a free data retrieval call binding the contract method 0x69aed5b5.
//
// Solidity: function getMPCInfo(bytes networkId, bytes pubkey) view returns((address,uint64,bytes,bytes) info)
func (_Optimex *OptimexSession) GetMPCInfo(networkId []byte, pubkey []byte) (ITypesMPCInfo, error) {
	return _Optimex.Contract.GetMPCInfo(&_Optimex.CallOpts, networkId, pubkey)
}

// GetMPCInfo is a free data retrieval call binding the contract method 0x69aed5b5.
//
// Solidity: function getMPCInfo(bytes networkId, bytes pubkey) view returns((address,uint64,bytes,bytes) info)
func (_Optimex *OptimexCallerSession) GetMPCInfo(networkId []byte, pubkey []byte) (ITypesMPCInfo, error) {
	return _Optimex.Contract.GetMPCInfo(&_Optimex.CallOpts, networkId, pubkey)
}

// GetManagementOwner is a free data retrieval call binding the contract method 0x50ae93e6.
//
// Solidity: function getManagementOwner() view returns(address)
func (_Optimex *OptimexCaller) GetManagementOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Optimex.contract.Call(opts, &out, "getManagementOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetManagementOwner is a free data retrieval call binding the contract method 0x50ae93e6.
//
// Solidity: function getManagementOwner() view returns(address)
func (_Optimex *OptimexSession) GetManagementOwner() (common.Address, error) {
	return _Optimex.Contract.GetManagementOwner(&_Optimex.CallOpts)
}

// GetManagementOwner is a free data retrieval call binding the contract method 0x50ae93e6.
//
// Solidity: function getManagementOwner() view returns(address)
func (_Optimex *OptimexCallerSession) GetManagementOwner() (common.Address, error) {
	return _Optimex.Contract.GetManagementOwner(&_Optimex.CallOpts)
}

// GetMaxAffiliateFeeRate is a free data retrieval call binding the contract method 0x0d459fec.
//
// Solidity: function getMaxAffiliateFeeRate(bytes fromChain, bytes toChain) view returns(uint256)
func (_Optimex *OptimexCaller) GetMaxAffiliateFeeRate(opts *bind.CallOpts, fromChain []byte, toChain []byte) (*big.Int, error) {
	var out []interface{}
	err := _Optimex.contract.Call(opts, &out, "getMaxAffiliateFeeRate", fromChain, toChain)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMaxAffiliateFeeRate is a free data retrieval call binding the contract method 0x0d459fec.
//
// Solidity: function getMaxAffiliateFeeRate(bytes fromChain, bytes toChain) view returns(uint256)
func (_Optimex *OptimexSession) GetMaxAffiliateFeeRate(fromChain []byte, toChain []byte) (*big.Int, error) {
	return _Optimex.Contract.GetMaxAffiliateFeeRate(&_Optimex.CallOpts, fromChain, toChain)
}

// GetMaxAffiliateFeeRate is a free data retrieval call binding the contract method 0x0d459fec.
//
// Solidity: function getMaxAffiliateFeeRate(bytes fromChain, bytes toChain) view returns(uint256)
func (_Optimex *OptimexCallerSession) GetMaxAffiliateFeeRate(fromChain []byte, toChain []byte) (*big.Int, error) {
	return _Optimex.Contract.GetMaxAffiliateFeeRate(&_Optimex.CallOpts, fromChain, toChain)
}

// GetPFeeRate is a free data retrieval call binding the contract method 0x6a0bd9c7.
//
// Solidity: function getPFeeRate() view returns(uint256)
func (_Optimex *OptimexCaller) GetPFeeRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Optimex.contract.Call(opts, &out, "getPFeeRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPFeeRate is a free data retrieval call binding the contract method 0x6a0bd9c7.
//
// Solidity: function getPFeeRate() view returns(uint256)
func (_Optimex *OptimexSession) GetPFeeRate() (*big.Int, error) {
	return _Optimex.Contract.GetPFeeRate(&_Optimex.CallOpts)
}

// GetPFeeRate is a free data retrieval call binding the contract method 0x6a0bd9c7.
//
// Solidity: function getPFeeRate() view returns(uint256)
func (_Optimex *OptimexCallerSession) GetPFeeRate() (*big.Int, error) {
	return _Optimex.Contract.GetPFeeRate(&_Optimex.CallOpts)
}

// GetPMMAccounts is a free data retrieval call binding the contract method 0xd0f8f46a.
//
// Solidity: function getPMMAccounts(bytes32 pmmId, uint256 fromIdx, uint256 toIdx) view returns(address[] list)
func (_Optimex *OptimexCaller) GetPMMAccounts(opts *bind.CallOpts, pmmId [32]byte, fromIdx *big.Int, toIdx *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _Optimex.contract.Call(opts, &out, "getPMMAccounts", pmmId, fromIdx, toIdx)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetPMMAccounts is a free data retrieval call binding the contract method 0xd0f8f46a.
//
// Solidity: function getPMMAccounts(bytes32 pmmId, uint256 fromIdx, uint256 toIdx) view returns(address[] list)
func (_Optimex *OptimexSession) GetPMMAccounts(pmmId [32]byte, fromIdx *big.Int, toIdx *big.Int) ([]common.Address, error) {
	return _Optimex.Contract.GetPMMAccounts(&_Optimex.CallOpts, pmmId, fromIdx, toIdx)
}

// GetPMMAccounts is a free data retrieval call binding the contract method 0xd0f8f46a.
//
// Solidity: function getPMMAccounts(bytes32 pmmId, uint256 fromIdx, uint256 toIdx) view returns(address[] list)
func (_Optimex *OptimexCallerSession) GetPMMAccounts(pmmId [32]byte, fromIdx *big.Int, toIdx *big.Int) ([]common.Address, error) {
	return _Optimex.Contract.GetPMMAccounts(&_Optimex.CallOpts, pmmId, fromIdx, toIdx)
}

// GetPMMSelection is a free data retrieval call binding the contract method 0x26cea89b.
//
// Solidity: function getPMMSelection(bytes32 tradeId) view returns(((uint256,uint64,bytes),(uint256,bytes32,bytes[2],uint64)))
func (_Optimex *OptimexCaller) GetPMMSelection(opts *bind.CallOpts, tradeId [32]byte) (ITypesPMMSelection, error) {
	var out []interface{}
	err := _Optimex.contract.Call(opts, &out, "getPMMSelection", tradeId)

	if err != nil {
		return *new(ITypesPMMSelection), err
	}

	out0 := *abi.ConvertType(out[0], new(ITypesPMMSelection)).(*ITypesPMMSelection)

	return out0, err

}

// GetPMMSelection is a free data retrieval call binding the contract method 0x26cea89b.
//
// Solidity: function getPMMSelection(bytes32 tradeId) view returns(((uint256,uint64,bytes),(uint256,bytes32,bytes[2],uint64)))
func (_Optimex *OptimexSession) GetPMMSelection(tradeId [32]byte) (ITypesPMMSelection, error) {
	return _Optimex.Contract.GetPMMSelection(&_Optimex.CallOpts, tradeId)
}

// GetPMMSelection is a free data retrieval call binding the contract method 0x26cea89b.
//
// Solidity: function getPMMSelection(bytes32 tradeId) view returns(((uint256,uint64,bytes),(uint256,bytes32,bytes[2],uint64)))
func (_Optimex *OptimexCallerSession) GetPMMSelection(tradeId [32]byte) (ITypesPMMSelection, error) {
	return _Optimex.Contract.GetPMMSelection(&_Optimex.CallOpts, tradeId)
}

// GetPresigns is a free data retrieval call binding the contract method 0xe1f65ebe.
//
// Solidity: function getPresigns(bytes32 tradeId) view returns((bytes32,bytes,bytes[])[])
func (_Optimex *OptimexCaller) GetPresigns(opts *bind.CallOpts, tradeId [32]byte) ([]ITypesPresign, error) {
	var out []interface{}
	err := _Optimex.contract.Call(opts, &out, "getPresigns", tradeId)

	if err != nil {
		return *new([]ITypesPresign), err
	}

	out0 := *abi.ConvertType(out[0], new([]ITypesPresign)).(*[]ITypesPresign)

	return out0, err

}

// GetPresigns is a free data retrieval call binding the contract method 0xe1f65ebe.
//
// Solidity: function getPresigns(bytes32 tradeId) view returns((bytes32,bytes,bytes[])[])
func (_Optimex *OptimexSession) GetPresigns(tradeId [32]byte) ([]ITypesPresign, error) {
	return _Optimex.Contract.GetPresigns(&_Optimex.CallOpts, tradeId)
}

// GetPresigns is a free data retrieval call binding the contract method 0xe1f65ebe.
//
// Solidity: function getPresigns(bytes32 tradeId) view returns((bytes32,bytes,bytes[])[])
func (_Optimex *OptimexCallerSession) GetPresigns(tradeId [32]byte) ([]ITypesPresign, error) {
	return _Optimex.Contract.GetPresigns(&_Optimex.CallOpts, tradeId)
}

// GetProtocolState is a free data retrieval call binding the contract method 0xd770adc7.
//
// Solidity: function getProtocolState() view returns(uint256)
func (_Optimex *OptimexCaller) GetProtocolState(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Optimex.contract.Call(opts, &out, "getProtocolState")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetProtocolState is a free data retrieval call binding the contract method 0xd770adc7.
//
// Solidity: function getProtocolState() view returns(uint256)
func (_Optimex *OptimexSession) GetProtocolState() (*big.Int, error) {
	return _Optimex.Contract.GetProtocolState(&_Optimex.CallOpts)
}

// GetProtocolState is a free data retrieval call binding the contract method 0xd770adc7.
//
// Solidity: function getProtocolState() view returns(uint256)
func (_Optimex *OptimexCallerSession) GetProtocolState() (*big.Int, error) {
	return _Optimex.Contract.GetProtocolState(&_Optimex.CallOpts)
}

// GetSettledPayment is a free data retrieval call binding the contract method 0xc98fa68c.
//
// Solidity: function getSettledPayment(bytes32 tradeId) view returns((bytes32,bytes,bytes,bool))
func (_Optimex *OptimexCaller) GetSettledPayment(opts *bind.CallOpts, tradeId [32]byte) (ITypesSettledPayment, error) {
	var out []interface{}
	err := _Optimex.contract.Call(opts, &out, "getSettledPayment", tradeId)

	if err != nil {
		return *new(ITypesSettledPayment), err
	}

	out0 := *abi.ConvertType(out[0], new(ITypesSettledPayment)).(*ITypesSettledPayment)

	return out0, err

}

// GetSettledPayment is a free data retrieval call binding the contract method 0xc98fa68c.
//
// Solidity: function getSettledPayment(bytes32 tradeId) view returns((bytes32,bytes,bytes,bool))
func (_Optimex *OptimexSession) GetSettledPayment(tradeId [32]byte) (ITypesSettledPayment, error) {
	return _Optimex.Contract.GetSettledPayment(&_Optimex.CallOpts, tradeId)
}

// GetSettledPayment is a free data retrieval call binding the contract method 0xc98fa68c.
//
// Solidity: function getSettledPayment(bytes32 tradeId) view returns((bytes32,bytes,bytes,bool))
func (_Optimex *OptimexCallerSession) GetSettledPayment(tradeId [32]byte) (ITypesSettledPayment, error) {
	return _Optimex.Contract.GetSettledPayment(&_Optimex.CallOpts, tradeId)
}

// GetTokens is a free data retrieval call binding the contract method 0x494cfc6c.
//
// Solidity: function getTokens(uint256 fromIdx, uint256 toIdx) view returns((bytes[5],uint256)[] list)
func (_Optimex *OptimexCaller) GetTokens(opts *bind.CallOpts, fromIdx *big.Int, toIdx *big.Int) ([]ITypesTokenInfo, error) {
	var out []interface{}
	err := _Optimex.contract.Call(opts, &out, "getTokens", fromIdx, toIdx)

	if err != nil {
		return *new([]ITypesTokenInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]ITypesTokenInfo)).(*[]ITypesTokenInfo)

	return out0, err

}

// GetTokens is a free data retrieval call binding the contract method 0x494cfc6c.
//
// Solidity: function getTokens(uint256 fromIdx, uint256 toIdx) view returns((bytes[5],uint256)[] list)
func (_Optimex *OptimexSession) GetTokens(fromIdx *big.Int, toIdx *big.Int) ([]ITypesTokenInfo, error) {
	return _Optimex.Contract.GetTokens(&_Optimex.CallOpts, fromIdx, toIdx)
}

// GetTokens is a free data retrieval call binding the contract method 0x494cfc6c.
//
// Solidity: function getTokens(uint256 fromIdx, uint256 toIdx) view returns((bytes[5],uint256)[] list)
func (_Optimex *OptimexCallerSession) GetTokens(fromIdx *big.Int, toIdx *big.Int) ([]ITypesTokenInfo, error) {
	return _Optimex.Contract.GetTokens(&_Optimex.CallOpts, fromIdx, toIdx)
}

// GetTradeData is a free data retrieval call binding the contract method 0x2afdebb1.
//
// Solidity: function getTradeData(bytes32 tradeId) view returns((uint256,(uint256,bytes[3],bytes[3]),(bytes[5],address,uint64)))
func (_Optimex *OptimexCaller) GetTradeData(opts *bind.CallOpts, tradeId [32]byte) (ITypesTradeData, error) {
	var out []interface{}
	err := _Optimex.contract.Call(opts, &out, "getTradeData", tradeId)

	if err != nil {
		return *new(ITypesTradeData), err
	}

	out0 := *abi.ConvertType(out[0], new(ITypesTradeData)).(*ITypesTradeData)

	return out0, err

}

// GetTradeData is a free data retrieval call binding the contract method 0x2afdebb1.
//
// Solidity: function getTradeData(bytes32 tradeId) view returns((uint256,(uint256,bytes[3],bytes[3]),(bytes[5],address,uint64)))
func (_Optimex *OptimexSession) GetTradeData(tradeId [32]byte) (ITypesTradeData, error) {
	return _Optimex.Contract.GetTradeData(&_Optimex.CallOpts, tradeId)
}

// GetTradeData is a free data retrieval call binding the contract method 0x2afdebb1.
//
// Solidity: function getTradeData(bytes32 tradeId) view returns((uint256,(uint256,bytes[3],bytes[3]),(bytes[5],address,uint64)))
func (_Optimex *OptimexCallerSession) GetTradeData(tradeId [32]byte) (ITypesTradeData, error) {
	return _Optimex.Contract.GetTradeData(&_Optimex.CallOpts, tradeId)
}

// IsMPCNode is a free data retrieval call binding the contract method 0xc6855320.
//
// Solidity: function isMPCNode(address account) view returns(bool)
func (_Optimex *OptimexCaller) IsMPCNode(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _Optimex.contract.Call(opts, &out, "isMPCNode", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMPCNode is a free data retrieval call binding the contract method 0xc6855320.
//
// Solidity: function isMPCNode(address account) view returns(bool)
func (_Optimex *OptimexSession) IsMPCNode(account common.Address) (bool, error) {
	return _Optimex.Contract.IsMPCNode(&_Optimex.CallOpts, account)
}

// IsMPCNode is a free data retrieval call binding the contract method 0xc6855320.
//
// Solidity: function isMPCNode(address account) view returns(bool)
func (_Optimex *OptimexCallerSession) IsMPCNode(account common.Address) (bool, error) {
	return _Optimex.Contract.IsMPCNode(&_Optimex.CallOpts, account)
}

// IsSolver is a free data retrieval call binding the contract method 0x02cc250d.
//
// Solidity: function isSolver(address account) view returns(bool)
func (_Optimex *OptimexCaller) IsSolver(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _Optimex.contract.Call(opts, &out, "isSolver", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSolver is a free data retrieval call binding the contract method 0x02cc250d.
//
// Solidity: function isSolver(address account) view returns(bool)
func (_Optimex *OptimexSession) IsSolver(account common.Address) (bool, error) {
	return _Optimex.Contract.IsSolver(&_Optimex.CallOpts, account)
}

// IsSolver is a free data retrieval call binding the contract method 0x02cc250d.
//
// Solidity: function isSolver(address account) view returns(bool)
func (_Optimex *OptimexCallerSession) IsSolver(account common.Address) (bool, error) {
	return _Optimex.Contract.IsSolver(&_Optimex.CallOpts, account)
}

// IsSuspended is a free data retrieval call binding the contract method 0xff909e01.
//
// Solidity: function isSuspended(uint8 stage) view returns(bool)
func (_Optimex *OptimexCaller) IsSuspended(opts *bind.CallOpts, stage uint8) (bool, error) {
	var out []interface{}
	err := _Optimex.contract.Call(opts, &out, "isSuspended", stage)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSuspended is a free data retrieval call binding the contract method 0xff909e01.
//
// Solidity: function isSuspended(uint8 stage) view returns(bool)
func (_Optimex *OptimexSession) IsSuspended(stage uint8) (bool, error) {
	return _Optimex.Contract.IsSuspended(&_Optimex.CallOpts, stage)
}

// IsSuspended is a free data retrieval call binding the contract method 0xff909e01.
//
// Solidity: function isSuspended(uint8 stage) view returns(bool)
func (_Optimex *OptimexCallerSession) IsSuspended(stage uint8) (bool, error) {
	return _Optimex.Contract.IsSuspended(&_Optimex.CallOpts, stage)
}

// IsValidNetwork is a free data retrieval call binding the contract method 0x85b0085e.
//
// Solidity: function isValidNetwork(bytes networkId) view returns(bool)
func (_Optimex *OptimexCaller) IsValidNetwork(opts *bind.CallOpts, networkId []byte) (bool, error) {
	var out []interface{}
	err := _Optimex.contract.Call(opts, &out, "isValidNetwork", networkId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidNetwork is a free data retrieval call binding the contract method 0x85b0085e.
//
// Solidity: function isValidNetwork(bytes networkId) view returns(bool)
func (_Optimex *OptimexSession) IsValidNetwork(networkId []byte) (bool, error) {
	return _Optimex.Contract.IsValidNetwork(&_Optimex.CallOpts, networkId)
}

// IsValidNetwork is a free data retrieval call binding the contract method 0x85b0085e.
//
// Solidity: function isValidNetwork(bytes networkId) view returns(bool)
func (_Optimex *OptimexCallerSession) IsValidNetwork(networkId []byte) (bool, error) {
	return _Optimex.Contract.IsValidNetwork(&_Optimex.CallOpts, networkId)
}

// IsValidPMM is a free data retrieval call binding the contract method 0xc779dbc0.
//
// Solidity: function isValidPMM(bytes32 pmmId) view returns(bool)
func (_Optimex *OptimexCaller) IsValidPMM(opts *bind.CallOpts, pmmId [32]byte) (bool, error) {
	var out []interface{}
	err := _Optimex.contract.Call(opts, &out, "isValidPMM", pmmId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidPMM is a free data retrieval call binding the contract method 0xc779dbc0.
//
// Solidity: function isValidPMM(bytes32 pmmId) view returns(bool)
func (_Optimex *OptimexSession) IsValidPMM(pmmId [32]byte) (bool, error) {
	return _Optimex.Contract.IsValidPMM(&_Optimex.CallOpts, pmmId)
}

// IsValidPMM is a free data retrieval call binding the contract method 0xc779dbc0.
//
// Solidity: function isValidPMM(bytes32 pmmId) view returns(bool)
func (_Optimex *OptimexCallerSession) IsValidPMM(pmmId [32]byte) (bool, error) {
	return _Optimex.Contract.IsValidPMM(&_Optimex.CallOpts, pmmId)
}

// IsValidPMMAccount is a free data retrieval call binding the contract method 0xe1156c4f.
//
// Solidity: function isValidPMMAccount(bytes32 pmmId, address account) view returns(bool)
func (_Optimex *OptimexCaller) IsValidPMMAccount(opts *bind.CallOpts, pmmId [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Optimex.contract.Call(opts, &out, "isValidPMMAccount", pmmId, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidPMMAccount is a free data retrieval call binding the contract method 0xe1156c4f.
//
// Solidity: function isValidPMMAccount(bytes32 pmmId, address account) view returns(bool)
func (_Optimex *OptimexSession) IsValidPMMAccount(pmmId [32]byte, account common.Address) (bool, error) {
	return _Optimex.Contract.IsValidPMMAccount(&_Optimex.CallOpts, pmmId, account)
}

// IsValidPMMAccount is a free data retrieval call binding the contract method 0xe1156c4f.
//
// Solidity: function isValidPMMAccount(bytes32 pmmId, address account) view returns(bool)
func (_Optimex *OptimexCallerSession) IsValidPMMAccount(pmmId [32]byte, account common.Address) (bool, error) {
	return _Optimex.Contract.IsValidPMMAccount(&_Optimex.CallOpts, pmmId, account)
}

// IsValidPubkey is a free data retrieval call binding the contract method 0x5f7e8e9d.
//
// Solidity: function isValidPubkey(bytes networkId, bytes pubkey) view returns(bool)
func (_Optimex *OptimexCaller) IsValidPubkey(opts *bind.CallOpts, networkId []byte, pubkey []byte) (bool, error) {
	var out []interface{}
	err := _Optimex.contract.Call(opts, &out, "isValidPubkey", networkId, pubkey)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidPubkey is a free data retrieval call binding the contract method 0x5f7e8e9d.
//
// Solidity: function isValidPubkey(bytes networkId, bytes pubkey) view returns(bool)
func (_Optimex *OptimexSession) IsValidPubkey(networkId []byte, pubkey []byte) (bool, error) {
	return _Optimex.Contract.IsValidPubkey(&_Optimex.CallOpts, networkId, pubkey)
}

// IsValidPubkey is a free data retrieval call binding the contract method 0x5f7e8e9d.
//
// Solidity: function isValidPubkey(bytes networkId, bytes pubkey) view returns(bool)
func (_Optimex *OptimexCallerSession) IsValidPubkey(networkId []byte, pubkey []byte) (bool, error) {
	return _Optimex.Contract.IsValidPubkey(&_Optimex.CallOpts, networkId, pubkey)
}

// IsValidToken is a free data retrieval call binding the contract method 0x6d25c123.
//
// Solidity: function isValidToken(bytes networkId, bytes tokenId) view returns(bool)
func (_Optimex *OptimexCaller) IsValidToken(opts *bind.CallOpts, networkId []byte, tokenId []byte) (bool, error) {
	var out []interface{}
	err := _Optimex.contract.Call(opts, &out, "isValidToken", networkId, tokenId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidToken is a free data retrieval call binding the contract method 0x6d25c123.
//
// Solidity: function isValidToken(bytes networkId, bytes tokenId) view returns(bool)
func (_Optimex *OptimexSession) IsValidToken(networkId []byte, tokenId []byte) (bool, error) {
	return _Optimex.Contract.IsValidToken(&_Optimex.CallOpts, networkId, tokenId)
}

// IsValidToken is a free data retrieval call binding the contract method 0x6d25c123.
//
// Solidity: function isValidToken(bytes networkId, bytes tokenId) view returns(bool)
func (_Optimex *OptimexCallerSession) IsValidToken(networkId []byte, tokenId []byte) (bool, error) {
	return _Optimex.Contract.IsValidToken(&_Optimex.CallOpts, networkId, tokenId)
}

// Management is a free data retrieval call binding the contract method 0x88a8d602.
//
// Solidity: function management() view returns(address)
func (_Optimex *OptimexCaller) Management(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Optimex.contract.Call(opts, &out, "management")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Management is a free data retrieval call binding the contract method 0x88a8d602.
//
// Solidity: function management() view returns(address)
func (_Optimex *OptimexSession) Management() (common.Address, error) {
	return _Optimex.Contract.Management(&_Optimex.CallOpts)
}

// Management is a free data retrieval call binding the contract method 0x88a8d602.
//
// Solidity: function management() view returns(address)
func (_Optimex *OptimexCallerSession) Management() (common.Address, error) {
	return _Optimex.Contract.Management(&_Optimex.CallOpts)
}

// NumOfPMMAccounts is a free data retrieval call binding the contract method 0xddcf7df1.
//
// Solidity: function numOfPMMAccounts(bytes32 pmmId) view returns(uint256)
func (_Optimex *OptimexCaller) NumOfPMMAccounts(opts *bind.CallOpts, pmmId [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Optimex.contract.Call(opts, &out, "numOfPMMAccounts", pmmId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumOfPMMAccounts is a free data retrieval call binding the contract method 0xddcf7df1.
//
// Solidity: function numOfPMMAccounts(bytes32 pmmId) view returns(uint256)
func (_Optimex *OptimexSession) NumOfPMMAccounts(pmmId [32]byte) (*big.Int, error) {
	return _Optimex.Contract.NumOfPMMAccounts(&_Optimex.CallOpts, pmmId)
}

// NumOfPMMAccounts is a free data retrieval call binding the contract method 0xddcf7df1.
//
// Solidity: function numOfPMMAccounts(bytes32 pmmId) view returns(uint256)
func (_Optimex *OptimexCallerSession) NumOfPMMAccounts(pmmId [32]byte) (*big.Int, error) {
	return _Optimex.Contract.NumOfPMMAccounts(&_Optimex.CallOpts, pmmId)
}

// NumOfSupportedTokens is a free data retrieval call binding the contract method 0x177b41ef.
//
// Solidity: function numOfSupportedTokens() view returns(uint256)
func (_Optimex *OptimexCaller) NumOfSupportedTokens(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Optimex.contract.Call(opts, &out, "numOfSupportedTokens")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumOfSupportedTokens is a free data retrieval call binding the contract method 0x177b41ef.
//
// Solidity: function numOfSupportedTokens() view returns(uint256)
func (_Optimex *OptimexSession) NumOfSupportedTokens() (*big.Int, error) {
	return _Optimex.Contract.NumOfSupportedTokens(&_Optimex.CallOpts)
}

// NumOfSupportedTokens is a free data retrieval call binding the contract method 0x177b41ef.
//
// Solidity: function numOfSupportedTokens() view returns(uint256)
func (_Optimex *OptimexCallerSession) NumOfSupportedTokens() (*big.Int, error) {
	return _Optimex.Contract.NumOfSupportedTokens(&_Optimex.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x767cb5bf.
//
// Solidity: function version(address ) view returns(uint256)
func (_Optimex *OptimexCaller) Version(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Optimex.contract.Call(opts, &out, "version", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x767cb5bf.
//
// Solidity: function version(address ) view returns(uint256)
func (_Optimex *OptimexSession) Version(arg0 common.Address) (*big.Int, error) {
	return _Optimex.Contract.Version(&_Optimex.CallOpts, arg0)
}

// Version is a free data retrieval call binding the contract method 0x767cb5bf.
//
// Solidity: function version(address ) view returns(uint256)
func (_Optimex *OptimexCallerSession) Version(arg0 common.Address) (*big.Int, error) {
	return _Optimex.Contract.Version(&_Optimex.CallOpts, arg0)
}

// BundlePayment is a paid mutator transaction binding the contract method 0x47012767.
//
// Solidity: function bundlePayment((bytes32[],uint64,uint64,bytes,bytes) bundle) returns()
func (_Optimex *OptimexTransactor) BundlePayment(opts *bind.TransactOpts, bundle ITypesBundlePayment) (*types.Transaction, error) {
	return _Optimex.contract.Transact(opts, "bundlePayment", bundle)
}

// BundlePayment is a paid mutator transaction binding the contract method 0x47012767.
//
// Solidity: function bundlePayment((bytes32[],uint64,uint64,bytes,bytes) bundle) returns()
func (_Optimex *OptimexSession) BundlePayment(bundle ITypesBundlePayment) (*types.Transaction, error) {
	return _Optimex.Contract.BundlePayment(&_Optimex.TransactOpts, bundle)
}

// BundlePayment is a paid mutator transaction binding the contract method 0x47012767.
//
// Solidity: function bundlePayment((bytes32[],uint64,uint64,bytes,bytes) bundle) returns()
func (_Optimex *OptimexTransactorSession) BundlePayment(bundle ITypesBundlePayment) (*types.Transaction, error) {
	return _Optimex.Contract.BundlePayment(&_Optimex.TransactOpts, bundle)
}

// ConfirmDeposit is a paid mutator transaction binding the contract method 0x7cb2e15f.
//
// Solidity: function confirmDeposit(bytes32 tradeId, bytes signature, bytes[] depositFromList) returns()
func (_Optimex *OptimexTransactor) ConfirmDeposit(opts *bind.TransactOpts, tradeId [32]byte, signature []byte, depositFromList [][]byte) (*types.Transaction, error) {
	return _Optimex.contract.Transact(opts, "confirmDeposit", tradeId, signature, depositFromList)
}

// ConfirmDeposit is a paid mutator transaction binding the contract method 0x7cb2e15f.
//
// Solidity: function confirmDeposit(bytes32 tradeId, bytes signature, bytes[] depositFromList) returns()
func (_Optimex *OptimexSession) ConfirmDeposit(tradeId [32]byte, signature []byte, depositFromList [][]byte) (*types.Transaction, error) {
	return _Optimex.Contract.ConfirmDeposit(&_Optimex.TransactOpts, tradeId, signature, depositFromList)
}

// ConfirmDeposit is a paid mutator transaction binding the contract method 0x7cb2e15f.
//
// Solidity: function confirmDeposit(bytes32 tradeId, bytes signature, bytes[] depositFromList) returns()
func (_Optimex *OptimexTransactorSession) ConfirmDeposit(tradeId [32]byte, signature []byte, depositFromList [][]byte) (*types.Transaction, error) {
	return _Optimex.Contract.ConfirmDeposit(&_Optimex.TransactOpts, tradeId, signature, depositFromList)
}

// ConfirmPayment is a paid mutator transaction binding the contract method 0x7ac64eac.
//
// Solidity: function confirmPayment(bytes32 tradeId, bytes signature) returns()
func (_Optimex *OptimexTransactor) ConfirmPayment(opts *bind.TransactOpts, tradeId [32]byte, signature []byte) (*types.Transaction, error) {
	return _Optimex.contract.Transact(opts, "confirmPayment", tradeId, signature)
}

// ConfirmPayment is a paid mutator transaction binding the contract method 0x7ac64eac.
//
// Solidity: function confirmPayment(bytes32 tradeId, bytes signature) returns()
func (_Optimex *OptimexSession) ConfirmPayment(tradeId [32]byte, signature []byte) (*types.Transaction, error) {
	return _Optimex.Contract.ConfirmPayment(&_Optimex.TransactOpts, tradeId, signature)
}

// ConfirmPayment is a paid mutator transaction binding the contract method 0x7ac64eac.
//
// Solidity: function confirmPayment(bytes32 tradeId, bytes signature) returns()
func (_Optimex *OptimexTransactorSession) ConfirmPayment(tradeId [32]byte, signature []byte) (*types.Transaction, error) {
	return _Optimex.Contract.ConfirmPayment(&_Optimex.TransactOpts, tradeId, signature)
}

// ConfirmSettlement is a paid mutator transaction binding the contract method 0x4fde3a7c.
//
// Solidity: function confirmSettlement(bytes32 tradeId, bytes releaseTxId, bytes signature) returns()
func (_Optimex *OptimexTransactor) ConfirmSettlement(opts *bind.TransactOpts, tradeId [32]byte, releaseTxId []byte, signature []byte) (*types.Transaction, error) {
	return _Optimex.contract.Transact(opts, "confirmSettlement", tradeId, releaseTxId, signature)
}

// ConfirmSettlement is a paid mutator transaction binding the contract method 0x4fde3a7c.
//
// Solidity: function confirmSettlement(bytes32 tradeId, bytes releaseTxId, bytes signature) returns()
func (_Optimex *OptimexSession) ConfirmSettlement(tradeId [32]byte, releaseTxId []byte, signature []byte) (*types.Transaction, error) {
	return _Optimex.Contract.ConfirmSettlement(&_Optimex.TransactOpts, tradeId, releaseTxId, signature)
}

// ConfirmSettlement is a paid mutator transaction binding the contract method 0x4fde3a7c.
//
// Solidity: function confirmSettlement(bytes32 tradeId, bytes releaseTxId, bytes signature) returns()
func (_Optimex *OptimexTransactorSession) ConfirmSettlement(tradeId [32]byte, releaseTxId []byte, signature []byte) (*types.Transaction, error) {
	return _Optimex.Contract.ConfirmSettlement(&_Optimex.TransactOpts, tradeId, releaseTxId, signature)
}

// SelectPMM is a paid mutator transaction binding the contract method 0x954505e8.
//
// Solidity: function selectPMM(bytes32 tradeId, ((uint256,uint64,bytes),(uint256,bytes32,bytes[2],uint64)) info) returns()
func (_Optimex *OptimexTransactor) SelectPMM(opts *bind.TransactOpts, tradeId [32]byte, info ITypesPMMSelection) (*types.Transaction, error) {
	return _Optimex.contract.Transact(opts, "selectPMM", tradeId, info)
}

// SelectPMM is a paid mutator transaction binding the contract method 0x954505e8.
//
// Solidity: function selectPMM(bytes32 tradeId, ((uint256,uint64,bytes),(uint256,bytes32,bytes[2],uint64)) info) returns()
func (_Optimex *OptimexSession) SelectPMM(tradeId [32]byte, info ITypesPMMSelection) (*types.Transaction, error) {
	return _Optimex.Contract.SelectPMM(&_Optimex.TransactOpts, tradeId, info)
}

// SelectPMM is a paid mutator transaction binding the contract method 0x954505e8.
//
// Solidity: function selectPMM(bytes32 tradeId, ((uint256,uint64,bytes),(uint256,bytes32,bytes[2],uint64)) info) returns()
func (_Optimex *OptimexTransactorSession) SelectPMM(tradeId [32]byte, info ITypesPMMSelection) (*types.Transaction, error) {
	return _Optimex.Contract.SelectPMM(&_Optimex.TransactOpts, tradeId, info)
}

// SetManagement is a paid mutator transaction binding the contract method 0xd4a22bde.
//
// Solidity: function setManagement(address newManagement) returns()
func (_Optimex *OptimexTransactor) SetManagement(opts *bind.TransactOpts, newManagement common.Address) (*types.Transaction, error) {
	return _Optimex.contract.Transact(opts, "setManagement", newManagement)
}

// SetManagement is a paid mutator transaction binding the contract method 0xd4a22bde.
//
// Solidity: function setManagement(address newManagement) returns()
func (_Optimex *OptimexSession) SetManagement(newManagement common.Address) (*types.Transaction, error) {
	return _Optimex.Contract.SetManagement(&_Optimex.TransactOpts, newManagement)
}

// SetManagement is a paid mutator transaction binding the contract method 0xd4a22bde.
//
// Solidity: function setManagement(address newManagement) returns()
func (_Optimex *OptimexTransactorSession) SetManagement(newManagement common.Address) (*types.Transaction, error) {
	return _Optimex.Contract.SetManagement(&_Optimex.TransactOpts, newManagement)
}

// SetRoute is a paid mutator transaction binding the contract method 0xd303de77.
//
// Solidity: function setRoute(address core, bytes fromChain, bytes toChain) returns()
func (_Optimex *OptimexTransactor) SetRoute(opts *bind.TransactOpts, core common.Address, fromChain []byte, toChain []byte) (*types.Transaction, error) {
	return _Optimex.contract.Transact(opts, "setRoute", core, fromChain, toChain)
}

// SetRoute is a paid mutator transaction binding the contract method 0xd303de77.
//
// Solidity: function setRoute(address core, bytes fromChain, bytes toChain) returns()
func (_Optimex *OptimexSession) SetRoute(core common.Address, fromChain []byte, toChain []byte) (*types.Transaction, error) {
	return _Optimex.Contract.SetRoute(&_Optimex.TransactOpts, core, fromChain, toChain)
}

// SetRoute is a paid mutator transaction binding the contract method 0xd303de77.
//
// Solidity: function setRoute(address core, bytes fromChain, bytes toChain) returns()
func (_Optimex *OptimexTransactorSession) SetRoute(core common.Address, fromChain []byte, toChain []byte) (*types.Transaction, error) {
	return _Optimex.Contract.SetRoute(&_Optimex.TransactOpts, core, fromChain, toChain)
}

// SubmitTrade is a paid mutator transaction binding the contract method 0xf8cd71b7.
//
// Solidity: function submitTrade(bytes32 tradeId, (uint256,(uint256,bytes[3],bytes[3]),(bytes[5],address,uint64)) tradeData, (uint256,string,bytes) affiliateInfo, (bytes32,bytes,bytes[])[] presignList) returns()
func (_Optimex *OptimexTransactor) SubmitTrade(opts *bind.TransactOpts, tradeId [32]byte, tradeData ITypesTradeData, affiliateInfo ITypesAffiliate, presignList []ITypesPresign) (*types.Transaction, error) {
	return _Optimex.contract.Transact(opts, "submitTrade", tradeId, tradeData, affiliateInfo, presignList)
}

// SubmitTrade is a paid mutator transaction binding the contract method 0xf8cd71b7.
//
// Solidity: function submitTrade(bytes32 tradeId, (uint256,(uint256,bytes[3],bytes[3]),(bytes[5],address,uint64)) tradeData, (uint256,string,bytes) affiliateInfo, (bytes32,bytes,bytes[])[] presignList) returns()
func (_Optimex *OptimexSession) SubmitTrade(tradeId [32]byte, tradeData ITypesTradeData, affiliateInfo ITypesAffiliate, presignList []ITypesPresign) (*types.Transaction, error) {
	return _Optimex.Contract.SubmitTrade(&_Optimex.TransactOpts, tradeId, tradeData, affiliateInfo, presignList)
}

// SubmitTrade is a paid mutator transaction binding the contract method 0xf8cd71b7.
//
// Solidity: function submitTrade(bytes32 tradeId, (uint256,(uint256,bytes[3],bytes[3]),(bytes[5],address,uint64)) tradeData, (uint256,string,bytes) affiliateInfo, (bytes32,bytes,bytes[])[] presignList) returns()
func (_Optimex *OptimexTransactorSession) SubmitTrade(tradeId [32]byte, tradeData ITypesTradeData, affiliateInfo ITypesAffiliate, presignList []ITypesPresign) (*types.Transaction, error) {
	return _Optimex.Contract.SubmitTrade(&_Optimex.TransactOpts, tradeId, tradeData, affiliateInfo, presignList)
}

// OptimexConfirmDepositIterator is returned from FilterConfirmDeposit and is used to iterate over the raw logs and unpacked data for ConfirmDeposit events raised by the Optimex contract.
type OptimexConfirmDepositIterator struct {
	Event *OptimexConfirmDeposit // Event containing the contract specifics and raw log

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
func (it *OptimexConfirmDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OptimexConfirmDeposit)
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
		it.Event = new(OptimexConfirmDeposit)
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
func (it *OptimexConfirmDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OptimexConfirmDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OptimexConfirmDeposit represents a ConfirmDeposit event raised by the Optimex contract.
type OptimexConfirmDeposit struct {
	Mpc      common.Address
	TradeId  [32]byte
	PFeeRate *big.Int
	AFeeRate *big.Int
	List     [][]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterConfirmDeposit is a free log retrieval operation binding the contract event 0xa00192ac6b20af83efb8ff3746ddbed8f2ac094469e27d58f4dca851800ddda6.
//
// Solidity: event ConfirmDeposit(address indexed mpc, bytes32 indexed tradeId, uint256 pFeeRate, uint256 aFeeRate, bytes[] list)
func (_Optimex *OptimexFilterer) FilterConfirmDeposit(opts *bind.FilterOpts, mpc []common.Address, tradeId [][32]byte) (*OptimexConfirmDepositIterator, error) {

	var mpcRule []interface{}
	for _, mpcItem := range mpc {
		mpcRule = append(mpcRule, mpcItem)
	}
	var tradeIdRule []interface{}
	for _, tradeIdItem := range tradeId {
		tradeIdRule = append(tradeIdRule, tradeIdItem)
	}

	logs, sub, err := _Optimex.contract.FilterLogs(opts, "ConfirmDeposit", mpcRule, tradeIdRule)
	if err != nil {
		return nil, err
	}
	return &OptimexConfirmDepositIterator{contract: _Optimex.contract, event: "ConfirmDeposit", logs: logs, sub: sub}, nil
}

// WatchConfirmDeposit is a free log subscription operation binding the contract event 0xa00192ac6b20af83efb8ff3746ddbed8f2ac094469e27d58f4dca851800ddda6.
//
// Solidity: event ConfirmDeposit(address indexed mpc, bytes32 indexed tradeId, uint256 pFeeRate, uint256 aFeeRate, bytes[] list)
func (_Optimex *OptimexFilterer) WatchConfirmDeposit(opts *bind.WatchOpts, sink chan<- *OptimexConfirmDeposit, mpc []common.Address, tradeId [][32]byte) (event.Subscription, error) {

	var mpcRule []interface{}
	for _, mpcItem := range mpc {
		mpcRule = append(mpcRule, mpcItem)
	}
	var tradeIdRule []interface{}
	for _, tradeIdItem := range tradeId {
		tradeIdRule = append(tradeIdRule, tradeIdItem)
	}

	logs, sub, err := _Optimex.contract.WatchLogs(opts, "ConfirmDeposit", mpcRule, tradeIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OptimexConfirmDeposit)
				if err := _Optimex.contract.UnpackLog(event, "ConfirmDeposit", log); err != nil {
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

// ParseConfirmDeposit is a log parse operation binding the contract event 0xa00192ac6b20af83efb8ff3746ddbed8f2ac094469e27d58f4dca851800ddda6.
//
// Solidity: event ConfirmDeposit(address indexed mpc, bytes32 indexed tradeId, uint256 pFeeRate, uint256 aFeeRate, bytes[] list)
func (_Optimex *OptimexFilterer) ParseConfirmDeposit(log types.Log) (*OptimexConfirmDeposit, error) {
	event := new(OptimexConfirmDeposit)
	if err := _Optimex.contract.UnpackLog(event, "ConfirmDeposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OptimexConfirmPaymentIterator is returned from FilterConfirmPayment and is used to iterate over the raw logs and unpacked data for ConfirmPayment events raised by the Optimex contract.
type OptimexConfirmPaymentIterator struct {
	Event *OptimexConfirmPayment // Event containing the contract specifics and raw log

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
func (it *OptimexConfirmPaymentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OptimexConfirmPayment)
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
		it.Event = new(OptimexConfirmPayment)
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
func (it *OptimexConfirmPaymentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OptimexConfirmPaymentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OptimexConfirmPayment represents a ConfirmPayment event raised by the Optimex contract.
type OptimexConfirmPayment struct {
	Mpc     common.Address
	TradeId [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterConfirmPayment is a free log retrieval operation binding the contract event 0x06de1d6c12fd91ba69a29e1add245f9a0466175bd93e81db6da34cb54dccdf40.
//
// Solidity: event ConfirmPayment(address indexed mpc, bytes32 indexed tradeId)
func (_Optimex *OptimexFilterer) FilterConfirmPayment(opts *bind.FilterOpts, mpc []common.Address, tradeId [][32]byte) (*OptimexConfirmPaymentIterator, error) {

	var mpcRule []interface{}
	for _, mpcItem := range mpc {
		mpcRule = append(mpcRule, mpcItem)
	}
	var tradeIdRule []interface{}
	for _, tradeIdItem := range tradeId {
		tradeIdRule = append(tradeIdRule, tradeIdItem)
	}

	logs, sub, err := _Optimex.contract.FilterLogs(opts, "ConfirmPayment", mpcRule, tradeIdRule)
	if err != nil {
		return nil, err
	}
	return &OptimexConfirmPaymentIterator{contract: _Optimex.contract, event: "ConfirmPayment", logs: logs, sub: sub}, nil
}

// WatchConfirmPayment is a free log subscription operation binding the contract event 0x06de1d6c12fd91ba69a29e1add245f9a0466175bd93e81db6da34cb54dccdf40.
//
// Solidity: event ConfirmPayment(address indexed mpc, bytes32 indexed tradeId)
func (_Optimex *OptimexFilterer) WatchConfirmPayment(opts *bind.WatchOpts, sink chan<- *OptimexConfirmPayment, mpc []common.Address, tradeId [][32]byte) (event.Subscription, error) {

	var mpcRule []interface{}
	for _, mpcItem := range mpc {
		mpcRule = append(mpcRule, mpcItem)
	}
	var tradeIdRule []interface{}
	for _, tradeIdItem := range tradeId {
		tradeIdRule = append(tradeIdRule, tradeIdItem)
	}

	logs, sub, err := _Optimex.contract.WatchLogs(opts, "ConfirmPayment", mpcRule, tradeIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OptimexConfirmPayment)
				if err := _Optimex.contract.UnpackLog(event, "ConfirmPayment", log); err != nil {
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

// ParseConfirmPayment is a log parse operation binding the contract event 0x06de1d6c12fd91ba69a29e1add245f9a0466175bd93e81db6da34cb54dccdf40.
//
// Solidity: event ConfirmPayment(address indexed mpc, bytes32 indexed tradeId)
func (_Optimex *OptimexFilterer) ParseConfirmPayment(log types.Log) (*OptimexConfirmPayment, error) {
	event := new(OptimexConfirmPayment)
	if err := _Optimex.contract.UnpackLog(event, "ConfirmPayment", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OptimexConfirmSettlementIterator is returned from FilterConfirmSettlement and is used to iterate over the raw logs and unpacked data for ConfirmSettlement events raised by the Optimex contract.
type OptimexConfirmSettlementIterator struct {
	Event *OptimexConfirmSettlement // Event containing the contract specifics and raw log

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
func (it *OptimexConfirmSettlementIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OptimexConfirmSettlement)
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
		it.Event = new(OptimexConfirmSettlement)
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
func (it *OptimexConfirmSettlementIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OptimexConfirmSettlementIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OptimexConfirmSettlement represents a ConfirmSettlement event raised by the Optimex contract.
type OptimexConfirmSettlement struct {
	Mpc     common.Address
	TradeId [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterConfirmSettlement is a free log retrieval operation binding the contract event 0x6292bb735dd6c41d01a708128a6cb40de00a8864b9208926d9fc1540497c90e6.
//
// Solidity: event ConfirmSettlement(address indexed mpc, bytes32 indexed tradeId)
func (_Optimex *OptimexFilterer) FilterConfirmSettlement(opts *bind.FilterOpts, mpc []common.Address, tradeId [][32]byte) (*OptimexConfirmSettlementIterator, error) {

	var mpcRule []interface{}
	for _, mpcItem := range mpc {
		mpcRule = append(mpcRule, mpcItem)
	}
	var tradeIdRule []interface{}
	for _, tradeIdItem := range tradeId {
		tradeIdRule = append(tradeIdRule, tradeIdItem)
	}

	logs, sub, err := _Optimex.contract.FilterLogs(opts, "ConfirmSettlement", mpcRule, tradeIdRule)
	if err != nil {
		return nil, err
	}
	return &OptimexConfirmSettlementIterator{contract: _Optimex.contract, event: "ConfirmSettlement", logs: logs, sub: sub}, nil
}

// WatchConfirmSettlement is a free log subscription operation binding the contract event 0x6292bb735dd6c41d01a708128a6cb40de00a8864b9208926d9fc1540497c90e6.
//
// Solidity: event ConfirmSettlement(address indexed mpc, bytes32 indexed tradeId)
func (_Optimex *OptimexFilterer) WatchConfirmSettlement(opts *bind.WatchOpts, sink chan<- *OptimexConfirmSettlement, mpc []common.Address, tradeId [][32]byte) (event.Subscription, error) {

	var mpcRule []interface{}
	for _, mpcItem := range mpc {
		mpcRule = append(mpcRule, mpcItem)
	}
	var tradeIdRule []interface{}
	for _, tradeIdItem := range tradeId {
		tradeIdRule = append(tradeIdRule, tradeIdItem)
	}

	logs, sub, err := _Optimex.contract.WatchLogs(opts, "ConfirmSettlement", mpcRule, tradeIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OptimexConfirmSettlement)
				if err := _Optimex.contract.UnpackLog(event, "ConfirmSettlement", log); err != nil {
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

// ParseConfirmSettlement is a log parse operation binding the contract event 0x6292bb735dd6c41d01a708128a6cb40de00a8864b9208926d9fc1540497c90e6.
//
// Solidity: event ConfirmSettlement(address indexed mpc, bytes32 indexed tradeId)
func (_Optimex *OptimexFilterer) ParseConfirmSettlement(log types.Log) (*OptimexConfirmSettlement, error) {
	event := new(OptimexConfirmSettlement)
	if err := _Optimex.contract.UnpackLog(event, "ConfirmSettlement", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OptimexMakePaymentIterator is returned from FilterMakePayment and is used to iterate over the raw logs and unpacked data for MakePayment events raised by the Optimex contract.
type OptimexMakePaymentIterator struct {
	Event *OptimexMakePayment // Event containing the contract specifics and raw log

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
func (it *OptimexMakePaymentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OptimexMakePayment)
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
		it.Event = new(OptimexMakePayment)
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
func (it *OptimexMakePaymentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OptimexMakePaymentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OptimexMakePayment represents a MakePayment event raised by the Optimex contract.
type OptimexMakePayment struct {
	Operator common.Address
	TradeId  [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMakePayment is a free log retrieval operation binding the contract event 0x469f9c81c5e357fe189ef62facf2bb3e1642cf8093da73ca51cb609d971377dd.
//
// Solidity: event MakePayment(address indexed operator, bytes32 indexed tradeId)
func (_Optimex *OptimexFilterer) FilterMakePayment(opts *bind.FilterOpts, operator []common.Address, tradeId [][32]byte) (*OptimexMakePaymentIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var tradeIdRule []interface{}
	for _, tradeIdItem := range tradeId {
		tradeIdRule = append(tradeIdRule, tradeIdItem)
	}

	logs, sub, err := _Optimex.contract.FilterLogs(opts, "MakePayment", operatorRule, tradeIdRule)
	if err != nil {
		return nil, err
	}
	return &OptimexMakePaymentIterator{contract: _Optimex.contract, event: "MakePayment", logs: logs, sub: sub}, nil
}

// WatchMakePayment is a free log subscription operation binding the contract event 0x469f9c81c5e357fe189ef62facf2bb3e1642cf8093da73ca51cb609d971377dd.
//
// Solidity: event MakePayment(address indexed operator, bytes32 indexed tradeId)
func (_Optimex *OptimexFilterer) WatchMakePayment(opts *bind.WatchOpts, sink chan<- *OptimexMakePayment, operator []common.Address, tradeId [][32]byte) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var tradeIdRule []interface{}
	for _, tradeIdItem := range tradeId {
		tradeIdRule = append(tradeIdRule, tradeIdItem)
	}

	logs, sub, err := _Optimex.contract.WatchLogs(opts, "MakePayment", operatorRule, tradeIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OptimexMakePayment)
				if err := _Optimex.contract.UnpackLog(event, "MakePayment", log); err != nil {
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

// ParseMakePayment is a log parse operation binding the contract event 0x469f9c81c5e357fe189ef62facf2bb3e1642cf8093da73ca51cb609d971377dd.
//
// Solidity: event MakePayment(address indexed operator, bytes32 indexed tradeId)
func (_Optimex *OptimexFilterer) ParseMakePayment(log types.Log) (*OptimexMakePayment, error) {
	event := new(OptimexMakePayment)
	if err := _Optimex.contract.UnpackLog(event, "MakePayment", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OptimexSelectPMMIterator is returned from FilterSelectPMM and is used to iterate over the raw logs and unpacked data for SelectPMM events raised by the Optimex contract.
type OptimexSelectPMMIterator struct {
	Event *OptimexSelectPMM // Event containing the contract specifics and raw log

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
func (it *OptimexSelectPMMIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OptimexSelectPMM)
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
		it.Event = new(OptimexSelectPMM)
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
func (it *OptimexSelectPMMIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OptimexSelectPMMIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OptimexSelectPMM represents a SelectPMM event raised by the Optimex contract.
type OptimexSelectPMM struct {
	Solver  common.Address
	TradeId [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSelectPMM is a free log retrieval operation binding the contract event 0x997438f726850dc0f689a2037c76da85881ee939121c2dce2ef362d02850f758.
//
// Solidity: event SelectPMM(address indexed solver, bytes32 indexed tradeId)
func (_Optimex *OptimexFilterer) FilterSelectPMM(opts *bind.FilterOpts, solver []common.Address, tradeId [][32]byte) (*OptimexSelectPMMIterator, error) {

	var solverRule []interface{}
	for _, solverItem := range solver {
		solverRule = append(solverRule, solverItem)
	}
	var tradeIdRule []interface{}
	for _, tradeIdItem := range tradeId {
		tradeIdRule = append(tradeIdRule, tradeIdItem)
	}

	logs, sub, err := _Optimex.contract.FilterLogs(opts, "SelectPMM", solverRule, tradeIdRule)
	if err != nil {
		return nil, err
	}
	return &OptimexSelectPMMIterator{contract: _Optimex.contract, event: "SelectPMM", logs: logs, sub: sub}, nil
}

// WatchSelectPMM is a free log subscription operation binding the contract event 0x997438f726850dc0f689a2037c76da85881ee939121c2dce2ef362d02850f758.
//
// Solidity: event SelectPMM(address indexed solver, bytes32 indexed tradeId)
func (_Optimex *OptimexFilterer) WatchSelectPMM(opts *bind.WatchOpts, sink chan<- *OptimexSelectPMM, solver []common.Address, tradeId [][32]byte) (event.Subscription, error) {

	var solverRule []interface{}
	for _, solverItem := range solver {
		solverRule = append(solverRule, solverItem)
	}
	var tradeIdRule []interface{}
	for _, tradeIdItem := range tradeId {
		tradeIdRule = append(tradeIdRule, tradeIdItem)
	}

	logs, sub, err := _Optimex.contract.WatchLogs(opts, "SelectPMM", solverRule, tradeIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OptimexSelectPMM)
				if err := _Optimex.contract.UnpackLog(event, "SelectPMM", log); err != nil {
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

// ParseSelectPMM is a log parse operation binding the contract event 0x997438f726850dc0f689a2037c76da85881ee939121c2dce2ef362d02850f758.
//
// Solidity: event SelectPMM(address indexed solver, bytes32 indexed tradeId)
func (_Optimex *OptimexFilterer) ParseSelectPMM(log types.Log) (*OptimexSelectPMM, error) {
	event := new(OptimexSelectPMM)
	if err := _Optimex.contract.UnpackLog(event, "SelectPMM", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OptimexSubmitTradeInfoIterator is returned from FilterSubmitTradeInfo and is used to iterate over the raw logs and unpacked data for SubmitTradeInfo events raised by the Optimex contract.
type OptimexSubmitTradeInfoIterator struct {
	Event *OptimexSubmitTradeInfo // Event containing the contract specifics and raw log

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
func (it *OptimexSubmitTradeInfoIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OptimexSubmitTradeInfo)
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
		it.Event = new(OptimexSubmitTradeInfo)
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
func (it *OptimexSubmitTradeInfoIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OptimexSubmitTradeInfoIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OptimexSubmitTradeInfo represents a SubmitTradeInfo event raised by the Optimex contract.
type OptimexSubmitTradeInfo struct {
	Solver  common.Address
	TradeId [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSubmitTradeInfo is a free log retrieval operation binding the contract event 0x730ae6437225bf566666bdac827d905d1fd8815c0ec5a3c35b290ccd2f8a1718.
//
// Solidity: event SubmitTradeInfo(address indexed solver, bytes32 indexed tradeId)
func (_Optimex *OptimexFilterer) FilterSubmitTradeInfo(opts *bind.FilterOpts, solver []common.Address, tradeId [][32]byte) (*OptimexSubmitTradeInfoIterator, error) {

	var solverRule []interface{}
	for _, solverItem := range solver {
		solverRule = append(solverRule, solverItem)
	}
	var tradeIdRule []interface{}
	for _, tradeIdItem := range tradeId {
		tradeIdRule = append(tradeIdRule, tradeIdItem)
	}

	logs, sub, err := _Optimex.contract.FilterLogs(opts, "SubmitTradeInfo", solverRule, tradeIdRule)
	if err != nil {
		return nil, err
	}
	return &OptimexSubmitTradeInfoIterator{contract: _Optimex.contract, event: "SubmitTradeInfo", logs: logs, sub: sub}, nil
}

// WatchSubmitTradeInfo is a free log subscription operation binding the contract event 0x730ae6437225bf566666bdac827d905d1fd8815c0ec5a3c35b290ccd2f8a1718.
//
// Solidity: event SubmitTradeInfo(address indexed solver, bytes32 indexed tradeId)
func (_Optimex *OptimexFilterer) WatchSubmitTradeInfo(opts *bind.WatchOpts, sink chan<- *OptimexSubmitTradeInfo, solver []common.Address, tradeId [][32]byte) (event.Subscription, error) {

	var solverRule []interface{}
	for _, solverItem := range solver {
		solverRule = append(solverRule, solverItem)
	}
	var tradeIdRule []interface{}
	for _, tradeIdItem := range tradeId {
		tradeIdRule = append(tradeIdRule, tradeIdItem)
	}

	logs, sub, err := _Optimex.contract.WatchLogs(opts, "SubmitTradeInfo", solverRule, tradeIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OptimexSubmitTradeInfo)
				if err := _Optimex.contract.UnpackLog(event, "SubmitTradeInfo", log); err != nil {
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

// ParseSubmitTradeInfo is a log parse operation binding the contract event 0x730ae6437225bf566666bdac827d905d1fd8815c0ec5a3c35b290ccd2f8a1718.
//
// Solidity: event SubmitTradeInfo(address indexed solver, bytes32 indexed tradeId)
func (_Optimex *OptimexFilterer) ParseSubmitTradeInfo(log types.Log) (*OptimexSubmitTradeInfo, error) {
	event := new(OptimexSubmitTradeInfo)
	if err := _Optimex.contract.UnpackLog(event, "SubmitTradeInfo", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OptimexUpdatedRouteIterator is returned from FilterUpdatedRoute and is used to iterate over the raw logs and unpacked data for UpdatedRoute events raised by the Optimex contract.
type OptimexUpdatedRouteIterator struct {
	Event *OptimexUpdatedRoute // Event containing the contract specifics and raw log

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
func (it *OptimexUpdatedRouteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OptimexUpdatedRoute)
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
		it.Event = new(OptimexUpdatedRoute)
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
func (it *OptimexUpdatedRouteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OptimexUpdatedRouteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OptimexUpdatedRoute represents a UpdatedRoute event raised by the Optimex contract.
type OptimexUpdatedRoute struct {
	Core      common.Address
	Version   *big.Int
	FromChain []byte
	ToChain   []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUpdatedRoute is a free log retrieval operation binding the contract event 0x5703d98db8e72b57f8044e76455805a023fcea4664019ff9fe2773566c77266f.
//
// Solidity: event UpdatedRoute(address indexed core, uint256 indexed version, bytes fromChain, bytes toChain)
func (_Optimex *OptimexFilterer) FilterUpdatedRoute(opts *bind.FilterOpts, core []common.Address, version []*big.Int) (*OptimexUpdatedRouteIterator, error) {

	var coreRule []interface{}
	for _, coreItem := range core {
		coreRule = append(coreRule, coreItem)
	}
	var versionRule []interface{}
	for _, versionItem := range version {
		versionRule = append(versionRule, versionItem)
	}

	logs, sub, err := _Optimex.contract.FilterLogs(opts, "UpdatedRoute", coreRule, versionRule)
	if err != nil {
		return nil, err
	}
	return &OptimexUpdatedRouteIterator{contract: _Optimex.contract, event: "UpdatedRoute", logs: logs, sub: sub}, nil
}

// WatchUpdatedRoute is a free log subscription operation binding the contract event 0x5703d98db8e72b57f8044e76455805a023fcea4664019ff9fe2773566c77266f.
//
// Solidity: event UpdatedRoute(address indexed core, uint256 indexed version, bytes fromChain, bytes toChain)
func (_Optimex *OptimexFilterer) WatchUpdatedRoute(opts *bind.WatchOpts, sink chan<- *OptimexUpdatedRoute, core []common.Address, version []*big.Int) (event.Subscription, error) {

	var coreRule []interface{}
	for _, coreItem := range core {
		coreRule = append(coreRule, coreItem)
	}
	var versionRule []interface{}
	for _, versionItem := range version {
		versionRule = append(versionRule, versionItem)
	}

	logs, sub, err := _Optimex.contract.WatchLogs(opts, "UpdatedRoute", coreRule, versionRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OptimexUpdatedRoute)
				if err := _Optimex.contract.UnpackLog(event, "UpdatedRoute", log); err != nil {
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

// ParseUpdatedRoute is a log parse operation binding the contract event 0x5703d98db8e72b57f8044e76455805a023fcea4664019ff9fe2773566c77266f.
//
// Solidity: event UpdatedRoute(address indexed core, uint256 indexed version, bytes fromChain, bytes toChain)
func (_Optimex *OptimexFilterer) ParseUpdatedRoute(log types.Log) (*OptimexUpdatedRoute, error) {
	event := new(OptimexUpdatedRoute)
	if err := _Optimex.contract.UnpackLog(event, "UpdatedRoute", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
