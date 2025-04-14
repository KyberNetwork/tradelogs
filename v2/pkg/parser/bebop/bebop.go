// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bebop

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

// OrderAggregate is an auto generated low-level Go binding around an user-defined struct.
type OrderAggregate struct {
	Expiry         *big.Int
	TakerAddress   common.Address
	MakerAddresses []common.Address
	MakerNonces    []*big.Int
	TakerTokens    [][]common.Address
	MakerTokens    [][]common.Address
	TakerAmounts   [][]*big.Int
	MakerAmounts   [][]*big.Int
	Receiver       common.Address
	Commands       []byte
	Flags          *big.Int
}

// OrderMulti is an auto generated low-level Go binding around an user-defined struct.
type OrderMulti struct {
	Expiry       *big.Int
	TakerAddress common.Address
	MakerAddress common.Address
	MakerNonce   *big.Int
	TakerTokens  []common.Address
	MakerTokens  []common.Address
	TakerAmounts []*big.Int
	MakerAmounts []*big.Int
	Receiver     common.Address
	Commands     []byte
	Flags        *big.Int
}

// OrderSingle is an auto generated low-level Go binding around an user-defined struct.
type OrderSingle struct {
	Expiry         *big.Int
	TakerAddress   common.Address
	MakerAddress   common.Address
	MakerNonce     *big.Int
	TakerToken     common.Address
	MakerToken     common.Address
	TakerAmount    *big.Int
	MakerAmount    *big.Int
	Receiver       common.Address
	PackedCommands *big.Int
	Flags          *big.Int
}

// SignatureMakerSignature is an auto generated low-level Go binding around an user-defined struct.
type SignatureMakerSignature struct {
	SignatureBytes []byte
	Flags          *big.Int
}

// SignatureMultiTokensPermit2Signature is an auto generated low-level Go binding around an user-defined struct.
type SignatureMultiTokensPermit2Signature struct {
	SignatureBytes []byte
	Deadline       *big.Int
	Nonces         []*big.Int
}

// SignaturePermit2Signature is an auto generated low-level Go binding around an user-defined struct.
type SignaturePermit2Signature struct {
	SignatureBytes []byte
	Deadline       *big.Int
	Nonce          *big.Int
}

// SignaturePermitSignature is an auto generated low-level Go binding around an user-defined struct.
type SignaturePermitSignature struct {
	SignatureBytes []byte
	Deadline       *big.Int
}

// TransferOldAggregateQuote is an auto generated low-level Go binding around an user-defined struct.
type TransferOldAggregateQuote struct {
	UseOldAmount bool
	MakerAmounts [][]*big.Int
	MakerNonces  []*big.Int
}

// TransferOldMultiQuote is an auto generated low-level Go binding around an user-defined struct.
type TransferOldMultiQuote struct {
	UseOldAmount bool
	MakerAmounts []*big.Int
	MakerNonce   *big.Int
}

// TransferOldSingleQuote is an auto generated low-level Go binding around an user-defined struct.
type TransferOldSingleQuote struct {
	UseOldAmount bool
	MakerAmount  *big.Int
	MakerNonce   *big.Int
}

// BebopMetaData contains all meta data concerning the Bebop contract.
var BebopMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_wrappedNativeToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_permit2\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_daiAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"CommandsLengthsMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedToSendNativeToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCommand\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCommandsLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidEIP1271Signature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidEIP721Signature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidETHSIGNSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidFlags\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPendingTransfersLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPermit2Commands\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPermit2TransfersLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSender\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSignatureType\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSignatureValueS\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSignatureValueV\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MakerAmountsLengthsMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ManyToManyNotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughNativeToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NullBeneficiary\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OrderExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OrderInvalidSigner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OrdersLengthsMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PartialFillNotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PartnerAlreadyRegistered\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PartnerFeeTooHigh\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokensLengthsMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnsafeCast\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UpdatedMakerAmountsTooLow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WrongWrappedTokenAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroMakerAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroNonce\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"eventId\",\"type\":\"uint128\"}],\"name\":\"BebopOrder\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"OrderSignerRegistered\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"taker_address\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"maker_addresses\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"maker_nonces\",\"type\":\"uint256[]\"},{\"internalType\":\"address[][]\",\"name\":\"taker_tokens\",\"type\":\"address[][]\"},{\"internalType\":\"address[][]\",\"name\":\"maker_tokens\",\"type\":\"address[][]\"},{\"internalType\":\"uint256[][]\",\"name\":\"taker_amounts\",\"type\":\"uint256[][]\"},{\"internalType\":\"uint256[][]\",\"name\":\"maker_amounts\",\"type\":\"uint256[][]\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"commands\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"flags\",\"type\":\"uint256\"}],\"internalType\":\"structOrder.Aggregate\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"partnerId\",\"type\":\"uint64\"},{\"internalType\":\"uint256[][]\",\"name\":\"updatedMakerAmounts\",\"type\":\"uint256[][]\"},{\"internalType\":\"uint256[]\",\"name\":\"updatedMakerNonces\",\"type\":\"uint256[]\"}],\"name\":\"hashAggregateOrder\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"taker_address\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"maker_address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maker_nonce\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"taker_tokens\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"maker_tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"taker_amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"maker_amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"commands\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"flags\",\"type\":\"uint256\"}],\"internalType\":\"structOrder.Multi\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"partnerId\",\"type\":\"uint64\"},{\"internalType\":\"uint256[]\",\"name\":\"updatedMakerAmounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"updatedMakerNonce\",\"type\":\"uint256\"}],\"name\":\"hashMultiOrder\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"taker_address\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"maker_address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maker_nonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"taker_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"maker_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"taker_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maker_amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"packed_commands\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"flags\",\"type\":\"uint256\"}],\"internalType\":\"structOrder.Single\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"partnerId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"updatedMakerAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updatedMakerNonce\",\"type\":\"uint256\"}],\"name\":\"hashSingleOrder\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"partners\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"fee\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"registered\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"registerAllowedOrderSigner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"partnerId\",\"type\":\"uint64\"},{\"internalType\":\"uint16\",\"name\":\"fee\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"}],\"name\":\"registerPartner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"taker_address\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"maker_addresses\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"maker_nonces\",\"type\":\"uint256[]\"},{\"internalType\":\"address[][]\",\"name\":\"taker_tokens\",\"type\":\"address[][]\"},{\"internalType\":\"address[][]\",\"name\":\"maker_tokens\",\"type\":\"address[][]\"},{\"internalType\":\"uint256[][]\",\"name\":\"taker_amounts\",\"type\":\"uint256[][]\"},{\"internalType\":\"uint256[][]\",\"name\":\"maker_amounts\",\"type\":\"uint256[][]\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"commands\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"flags\",\"type\":\"uint256\"}],\"internalType\":\"structOrder.Aggregate\",\"name\":\"order\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"signatureBytes\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"flags\",\"type\":\"uint256\"}],\"internalType\":\"structSignature.MakerSignature[]\",\"name\":\"makersSignatures\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"filledTakerAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"useOldAmount\",\"type\":\"bool\"},{\"internalType\":\"uint256[][]\",\"name\":\"makerAmounts\",\"type\":\"uint256[][]\"},{\"internalType\":\"uint256[]\",\"name\":\"makerNonces\",\"type\":\"uint256[]\"}],\"internalType\":\"structTransfer.OldAggregateQuote\",\"name\":\"takerQuoteInfo\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"takerSignature\",\"type\":\"bytes\"}],\"name\":\"settleAggregate\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"taker_address\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"maker_addresses\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"maker_nonces\",\"type\":\"uint256[]\"},{\"internalType\":\"address[][]\",\"name\":\"taker_tokens\",\"type\":\"address[][]\"},{\"internalType\":\"address[][]\",\"name\":\"maker_tokens\",\"type\":\"address[][]\"},{\"internalType\":\"uint256[][]\",\"name\":\"taker_amounts\",\"type\":\"uint256[][]\"},{\"internalType\":\"uint256[][]\",\"name\":\"maker_amounts\",\"type\":\"uint256[][]\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"commands\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"flags\",\"type\":\"uint256\"}],\"internalType\":\"structOrder.Aggregate\",\"name\":\"order\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"signatureBytes\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"flags\",\"type\":\"uint256\"}],\"internalType\":\"structSignature.MakerSignature[]\",\"name\":\"makersSignatures\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"filledTakerAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"useOldAmount\",\"type\":\"bool\"},{\"internalType\":\"uint256[][]\",\"name\":\"makerAmounts\",\"type\":\"uint256[][]\"},{\"internalType\":\"uint256[]\",\"name\":\"makerNonces\",\"type\":\"uint256[]\"}],\"internalType\":\"structTransfer.OldAggregateQuote\",\"name\":\"takerQuoteInfo\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"takerSignature\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"signatureBytes\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structSignature.PermitSignature\",\"name\":\"takerPermitSignature\",\"type\":\"tuple\"}],\"name\":\"settleAggregateAndSignPermit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"taker_address\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"maker_addresses\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"maker_nonces\",\"type\":\"uint256[]\"},{\"internalType\":\"address[][]\",\"name\":\"taker_tokens\",\"type\":\"address[][]\"},{\"internalType\":\"address[][]\",\"name\":\"maker_tokens\",\"type\":\"address[][]\"},{\"internalType\":\"uint256[][]\",\"name\":\"taker_amounts\",\"type\":\"uint256[][]\"},{\"internalType\":\"uint256[][]\",\"name\":\"maker_amounts\",\"type\":\"uint256[][]\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"commands\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"flags\",\"type\":\"uint256\"}],\"internalType\":\"structOrder.Aggregate\",\"name\":\"order\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"signatureBytes\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"flags\",\"type\":\"uint256\"}],\"internalType\":\"structSignature.MakerSignature[]\",\"name\":\"makersSignatures\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"filledTakerAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"useOldAmount\",\"type\":\"bool\"},{\"internalType\":\"uint256[][]\",\"name\":\"makerAmounts\",\"type\":\"uint256[][]\"},{\"internalType\":\"uint256[]\",\"name\":\"makerNonces\",\"type\":\"uint256[]\"}],\"internalType\":\"structTransfer.OldAggregateQuote\",\"name\":\"takerQuoteInfo\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"takerSignature\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"signatureBytes\",\"type\":\"bytes\"},{\"internalType\":\"uint48\",\"name\":\"deadline\",\"type\":\"uint48\"},{\"internalType\":\"uint48[]\",\"name\":\"nonces\",\"type\":\"uint48[]\"}],\"internalType\":\"structSignature.MultiTokensPermit2Signature\",\"name\":\"infoPermit2\",\"type\":\"tuple\"}],\"name\":\"settleAggregateAndSignPermit2\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"taker_address\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"maker_address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maker_nonce\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"taker_tokens\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"maker_tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"taker_amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"maker_amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"commands\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"flags\",\"type\":\"uint256\"}],\"internalType\":\"structOrder.Multi\",\"name\":\"order\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"signatureBytes\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"flags\",\"type\":\"uint256\"}],\"internalType\":\"structSignature.MakerSignature\",\"name\":\"makerSignature\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"filledTakerAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"useOldAmount\",\"type\":\"bool\"},{\"internalType\":\"uint256[]\",\"name\":\"makerAmounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"makerNonce\",\"type\":\"uint256\"}],\"internalType\":\"structTransfer.OldMultiQuote\",\"name\":\"takerQuoteInfo\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"takerSignature\",\"type\":\"bytes\"}],\"name\":\"settleMulti\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"taker_address\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"maker_address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maker_nonce\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"taker_tokens\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"maker_tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"taker_amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"maker_amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"commands\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"flags\",\"type\":\"uint256\"}],\"internalType\":\"structOrder.Multi\",\"name\":\"order\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"signatureBytes\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"flags\",\"type\":\"uint256\"}],\"internalType\":\"structSignature.MakerSignature\",\"name\":\"makerSignature\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"filledTakerAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"useOldAmount\",\"type\":\"bool\"},{\"internalType\":\"uint256[]\",\"name\":\"makerAmounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"makerNonce\",\"type\":\"uint256\"}],\"internalType\":\"structTransfer.OldMultiQuote\",\"name\":\"takerQuoteInfo\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"takerSignature\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"signatureBytes\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structSignature.PermitSignature\",\"name\":\"takerPermitSignature\",\"type\":\"tuple\"}],\"name\":\"settleMultiAndSignPermit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"taker_address\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"maker_address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maker_nonce\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"taker_tokens\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"maker_tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"taker_amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"maker_amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"commands\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"flags\",\"type\":\"uint256\"}],\"internalType\":\"structOrder.Multi\",\"name\":\"order\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"signatureBytes\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"flags\",\"type\":\"uint256\"}],\"internalType\":\"structSignature.MakerSignature\",\"name\":\"makerSignature\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"filledTakerAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"useOldAmount\",\"type\":\"bool\"},{\"internalType\":\"uint256[]\",\"name\":\"makerAmounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"makerNonce\",\"type\":\"uint256\"}],\"internalType\":\"structTransfer.OldMultiQuote\",\"name\":\"takerQuoteInfo\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"takerSignature\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"signatureBytes\",\"type\":\"bytes\"},{\"internalType\":\"uint48\",\"name\":\"deadline\",\"type\":\"uint48\"},{\"internalType\":\"uint48[]\",\"name\":\"nonces\",\"type\":\"uint48[]\"}],\"internalType\":\"structSignature.MultiTokensPermit2Signature\",\"name\":\"infoPermit2\",\"type\":\"tuple\"}],\"name\":\"settleMultiAndSignPermit2\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"taker_address\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"maker_address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maker_nonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"taker_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"maker_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"taker_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maker_amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"packed_commands\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"flags\",\"type\":\"uint256\"}],\"internalType\":\"structOrder.Single\",\"name\":\"order\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"signatureBytes\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"flags\",\"type\":\"uint256\"}],\"internalType\":\"structSignature.MakerSignature\",\"name\":\"makerSignature\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"filledTakerAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"useOldAmount\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"makerAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerNonce\",\"type\":\"uint256\"}],\"internalType\":\"structTransfer.OldSingleQuote\",\"name\":\"takerQuoteInfo\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"takerSignature\",\"type\":\"bytes\"}],\"name\":\"settleSingle\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"taker_address\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"maker_address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maker_nonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"taker_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"maker_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"taker_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maker_amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"packed_commands\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"flags\",\"type\":\"uint256\"}],\"internalType\":\"structOrder.Single\",\"name\":\"order\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"signatureBytes\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"flags\",\"type\":\"uint256\"}],\"internalType\":\"structSignature.MakerSignature\",\"name\":\"makerSignature\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"filledTakerAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"useOldAmount\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"makerAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerNonce\",\"type\":\"uint256\"}],\"internalType\":\"structTransfer.OldSingleQuote\",\"name\":\"takerQuoteInfo\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"takerSignature\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"signatureBytes\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structSignature.PermitSignature\",\"name\":\"takerPermitSignature\",\"type\":\"tuple\"}],\"name\":\"settleSingleAndSignPermit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"taker_address\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"maker_address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maker_nonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"taker_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"maker_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"taker_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maker_amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"packed_commands\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"flags\",\"type\":\"uint256\"}],\"internalType\":\"structOrder.Single\",\"name\":\"order\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"signatureBytes\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"flags\",\"type\":\"uint256\"}],\"internalType\":\"structSignature.MakerSignature\",\"name\":\"makerSignature\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"filledTakerAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"useOldAmount\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"makerAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerNonce\",\"type\":\"uint256\"}],\"internalType\":\"structTransfer.OldSingleQuote\",\"name\":\"takerQuoteInfo\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"takerSignature\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"signatureBytes\",\"type\":\"bytes\"},{\"internalType\":\"uint48\",\"name\":\"deadline\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"nonce\",\"type\":\"uint48\"}],\"internalType\":\"structSignature.Permit2Signature\",\"name\":\"takerPermit2Signature\",\"type\":\"tuple\"}],\"name\":\"settleSingleAndSignPermit2\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"taker_address\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"maker_addresses\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"maker_nonces\",\"type\":\"uint256[]\"},{\"internalType\":\"address[][]\",\"name\":\"taker_tokens\",\"type\":\"address[][]\"},{\"internalType\":\"address[][]\",\"name\":\"maker_tokens\",\"type\":\"address[][]\"},{\"internalType\":\"uint256[][]\",\"name\":\"taker_amounts\",\"type\":\"uint256[][]\"},{\"internalType\":\"uint256[][]\",\"name\":\"maker_amounts\",\"type\":\"uint256[][]\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"commands\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"flags\",\"type\":\"uint256\"}],\"internalType\":\"structOrder.Aggregate\",\"name\":\"order\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"signatureBytes\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"flags\",\"type\":\"uint256\"}],\"internalType\":\"structSignature.MakerSignature[]\",\"name\":\"makersSignatures\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"filledTakerAmount\",\"type\":\"uint256\"}],\"name\":\"swapAggregate\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"taker_address\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"maker_address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maker_nonce\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"taker_tokens\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"maker_tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"taker_amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"maker_amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"commands\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"flags\",\"type\":\"uint256\"}],\"internalType\":\"structOrder.Multi\",\"name\":\"order\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"signatureBytes\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"flags\",\"type\":\"uint256\"}],\"internalType\":\"structSignature.MakerSignature\",\"name\":\"makerSignature\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"filledTakerAmount\",\"type\":\"uint256\"}],\"name\":\"swapMulti\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"taker_address\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"maker_address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maker_nonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"taker_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"maker_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"taker_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maker_amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"packed_commands\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"flags\",\"type\":\"uint256\"}],\"internalType\":\"structOrder.Single\",\"name\":\"order\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"signatureBytes\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"flags\",\"type\":\"uint256\"}],\"internalType\":\"structSignature.MakerSignature\",\"name\":\"makerSignature\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"filledTakerAmount\",\"type\":\"uint256\"}],\"name\":\"swapSingle\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"taker_address\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"maker_address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maker_nonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"taker_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"maker_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"taker_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maker_amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"packed_commands\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"flags\",\"type\":\"uint256\"}],\"internalType\":\"structOrder.Single\",\"name\":\"order\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"signatureBytes\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"flags\",\"type\":\"uint256\"}],\"internalType\":\"structSignature.MakerSignature\",\"name\":\"makerSignature\",\"type\":\"tuple\"}],\"name\":\"swapSingleFromContract\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// BebopABI is the input ABI used to generate the binding from.
// Deprecated: Use BebopMetaData.ABI instead.
var BebopABI = BebopMetaData.ABI

// Bebop is an auto generated Go binding around an Ethereum contract.
type Bebop struct {
	BebopCaller     // Read-only binding to the contract
	BebopTransactor // Write-only binding to the contract
	BebopFilterer   // Log filterer for contract events
}

// BebopCaller is an auto generated read-only Go binding around an Ethereum contract.
type BebopCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BebopTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BebopTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BebopFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BebopFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BebopSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BebopSession struct {
	Contract     *Bebop            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BebopCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BebopCallerSession struct {
	Contract *BebopCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BebopTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BebopTransactorSession struct {
	Contract     *BebopTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BebopRaw is an auto generated low-level Go binding around an Ethereum contract.
type BebopRaw struct {
	Contract *Bebop // Generic contract binding to access the raw methods on
}

// BebopCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BebopCallerRaw struct {
	Contract *BebopCaller // Generic read-only contract binding to access the raw methods on
}

// BebopTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BebopTransactorRaw struct {
	Contract *BebopTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBebop creates a new instance of Bebop, bound to a specific deployed contract.
func NewBebop(address common.Address, backend bind.ContractBackend) (*Bebop, error) {
	contract, err := bindBebop(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bebop{BebopCaller: BebopCaller{contract: contract}, BebopTransactor: BebopTransactor{contract: contract}, BebopFilterer: BebopFilterer{contract: contract}}, nil
}

// NewBebopCaller creates a new read-only instance of Bebop, bound to a specific deployed contract.
func NewBebopCaller(address common.Address, caller bind.ContractCaller) (*BebopCaller, error) {
	contract, err := bindBebop(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BebopCaller{contract: contract}, nil
}

// NewBebopTransactor creates a new write-only instance of Bebop, bound to a specific deployed contract.
func NewBebopTransactor(address common.Address, transactor bind.ContractTransactor) (*BebopTransactor, error) {
	contract, err := bindBebop(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BebopTransactor{contract: contract}, nil
}

// NewBebopFilterer creates a new log filterer instance of Bebop, bound to a specific deployed contract.
func NewBebopFilterer(address common.Address, filterer bind.ContractFilterer) (*BebopFilterer, error) {
	contract, err := bindBebop(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BebopFilterer{contract: contract}, nil
}

// bindBebop binds a generic wrapper to an already deployed contract.
func bindBebop(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BebopMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bebop *BebopRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bebop.Contract.BebopCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bebop *BebopRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bebop.Contract.BebopTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bebop *BebopRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bebop.Contract.BebopTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bebop *BebopCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bebop.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bebop *BebopTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bebop.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bebop *BebopTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bebop.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Bebop *BebopCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Bebop.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Bebop *BebopSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _Bebop.Contract.DOMAINSEPARATOR(&_Bebop.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Bebop *BebopCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _Bebop.Contract.DOMAINSEPARATOR(&_Bebop.CallOpts)
}

// HashAggregateOrder is a free data retrieval call binding the contract method 0x8cce3140.
//
// Solidity: function hashAggregateOrder((uint256,address,address[],uint256[],address[][],address[][],uint256[][],uint256[][],address,bytes,uint256) order, uint64 partnerId, uint256[][] updatedMakerAmounts, uint256[] updatedMakerNonces) view returns(bytes32)
func (_Bebop *BebopCaller) HashAggregateOrder(opts *bind.CallOpts, order OrderAggregate, partnerId uint64, updatedMakerAmounts [][]*big.Int, updatedMakerNonces []*big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Bebop.contract.Call(opts, &out, "hashAggregateOrder", order, partnerId, updatedMakerAmounts, updatedMakerNonces)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HashAggregateOrder is a free data retrieval call binding the contract method 0x8cce3140.
//
// Solidity: function hashAggregateOrder((uint256,address,address[],uint256[],address[][],address[][],uint256[][],uint256[][],address,bytes,uint256) order, uint64 partnerId, uint256[][] updatedMakerAmounts, uint256[] updatedMakerNonces) view returns(bytes32)
func (_Bebop *BebopSession) HashAggregateOrder(order OrderAggregate, partnerId uint64, updatedMakerAmounts [][]*big.Int, updatedMakerNonces []*big.Int) ([32]byte, error) {
	return _Bebop.Contract.HashAggregateOrder(&_Bebop.CallOpts, order, partnerId, updatedMakerAmounts, updatedMakerNonces)
}

// HashAggregateOrder is a free data retrieval call binding the contract method 0x8cce3140.
//
// Solidity: function hashAggregateOrder((uint256,address,address[],uint256[],address[][],address[][],uint256[][],uint256[][],address,bytes,uint256) order, uint64 partnerId, uint256[][] updatedMakerAmounts, uint256[] updatedMakerNonces) view returns(bytes32)
func (_Bebop *BebopCallerSession) HashAggregateOrder(order OrderAggregate, partnerId uint64, updatedMakerAmounts [][]*big.Int, updatedMakerNonces []*big.Int) ([32]byte, error) {
	return _Bebop.Contract.HashAggregateOrder(&_Bebop.CallOpts, order, partnerId, updatedMakerAmounts, updatedMakerNonces)
}

// HashMultiOrder is a free data retrieval call binding the contract method 0x1bceedd7.
//
// Solidity: function hashMultiOrder((uint256,address,address,uint256,address[],address[],uint256[],uint256[],address,bytes,uint256) order, uint64 partnerId, uint256[] updatedMakerAmounts, uint256 updatedMakerNonce) view returns(bytes32)
func (_Bebop *BebopCaller) HashMultiOrder(opts *bind.CallOpts, order OrderMulti, partnerId uint64, updatedMakerAmounts []*big.Int, updatedMakerNonce *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Bebop.contract.Call(opts, &out, "hashMultiOrder", order, partnerId, updatedMakerAmounts, updatedMakerNonce)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HashMultiOrder is a free data retrieval call binding the contract method 0x1bceedd7.
//
// Solidity: function hashMultiOrder((uint256,address,address,uint256,address[],address[],uint256[],uint256[],address,bytes,uint256) order, uint64 partnerId, uint256[] updatedMakerAmounts, uint256 updatedMakerNonce) view returns(bytes32)
func (_Bebop *BebopSession) HashMultiOrder(order OrderMulti, partnerId uint64, updatedMakerAmounts []*big.Int, updatedMakerNonce *big.Int) ([32]byte, error) {
	return _Bebop.Contract.HashMultiOrder(&_Bebop.CallOpts, order, partnerId, updatedMakerAmounts, updatedMakerNonce)
}

// HashMultiOrder is a free data retrieval call binding the contract method 0x1bceedd7.
//
// Solidity: function hashMultiOrder((uint256,address,address,uint256,address[],address[],uint256[],uint256[],address,bytes,uint256) order, uint64 partnerId, uint256[] updatedMakerAmounts, uint256 updatedMakerNonce) view returns(bytes32)
func (_Bebop *BebopCallerSession) HashMultiOrder(order OrderMulti, partnerId uint64, updatedMakerAmounts []*big.Int, updatedMakerNonce *big.Int) ([32]byte, error) {
	return _Bebop.Contract.HashMultiOrder(&_Bebop.CallOpts, order, partnerId, updatedMakerAmounts, updatedMakerNonce)
}

// HashSingleOrder is a free data retrieval call binding the contract method 0x529c4a2f.
//
// Solidity: function hashSingleOrder((uint256,address,address,uint256,address,address,uint256,uint256,address,uint256,uint256) order, uint64 partnerId, uint256 updatedMakerAmount, uint256 updatedMakerNonce) view returns(bytes32)
func (_Bebop *BebopCaller) HashSingleOrder(opts *bind.CallOpts, order OrderSingle, partnerId uint64, updatedMakerAmount *big.Int, updatedMakerNonce *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Bebop.contract.Call(opts, &out, "hashSingleOrder", order, partnerId, updatedMakerAmount, updatedMakerNonce)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HashSingleOrder is a free data retrieval call binding the contract method 0x529c4a2f.
//
// Solidity: function hashSingleOrder((uint256,address,address,uint256,address,address,uint256,uint256,address,uint256,uint256) order, uint64 partnerId, uint256 updatedMakerAmount, uint256 updatedMakerNonce) view returns(bytes32)
func (_Bebop *BebopSession) HashSingleOrder(order OrderSingle, partnerId uint64, updatedMakerAmount *big.Int, updatedMakerNonce *big.Int) ([32]byte, error) {
	return _Bebop.Contract.HashSingleOrder(&_Bebop.CallOpts, order, partnerId, updatedMakerAmount, updatedMakerNonce)
}

// HashSingleOrder is a free data retrieval call binding the contract method 0x529c4a2f.
//
// Solidity: function hashSingleOrder((uint256,address,address,uint256,address,address,uint256,uint256,address,uint256,uint256) order, uint64 partnerId, uint256 updatedMakerAmount, uint256 updatedMakerNonce) view returns(bytes32)
func (_Bebop *BebopCallerSession) HashSingleOrder(order OrderSingle, partnerId uint64, updatedMakerAmount *big.Int, updatedMakerNonce *big.Int) ([32]byte, error) {
	return _Bebop.Contract.HashSingleOrder(&_Bebop.CallOpts, order, partnerId, updatedMakerAmount, updatedMakerNonce)
}

// Partners is a free data retrieval call binding the contract method 0x6d2a03d8.
//
// Solidity: function partners(uint64 ) view returns(uint16 fee, address beneficiary, bool registered)
func (_Bebop *BebopCaller) Partners(opts *bind.CallOpts, arg0 uint64) (struct {
	Fee         uint16
	Beneficiary common.Address
	Registered  bool
}, error) {
	var out []interface{}
	err := _Bebop.contract.Call(opts, &out, "partners", arg0)

	outstruct := new(struct {
		Fee         uint16
		Beneficiary common.Address
		Registered  bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Fee = *abi.ConvertType(out[0], new(uint16)).(*uint16)
	outstruct.Beneficiary = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Registered = *abi.ConvertType(out[2], new(bool)).(*bool)

	return *outstruct, err

}

// Partners is a free data retrieval call binding the contract method 0x6d2a03d8.
//
// Solidity: function partners(uint64 ) view returns(uint16 fee, address beneficiary, bool registered)
func (_Bebop *BebopSession) Partners(arg0 uint64) (struct {
	Fee         uint16
	Beneficiary common.Address
	Registered  bool
}, error) {
	return _Bebop.Contract.Partners(&_Bebop.CallOpts, arg0)
}

// Partners is a free data retrieval call binding the contract method 0x6d2a03d8.
//
// Solidity: function partners(uint64 ) view returns(uint16 fee, address beneficiary, bool registered)
func (_Bebop *BebopCallerSession) Partners(arg0 uint64) (struct {
	Fee         uint16
	Beneficiary common.Address
	Registered  bool
}, error) {
	return _Bebop.Contract.Partners(&_Bebop.CallOpts, arg0)
}

// RegisterAllowedOrderSigner is a paid mutator transaction binding the contract method 0xea7faa61.
//
// Solidity: function registerAllowedOrderSigner(address signer, bool allowed) returns()
func (_Bebop *BebopTransactor) RegisterAllowedOrderSigner(opts *bind.TransactOpts, signer common.Address, allowed bool) (*types.Transaction, error) {
	return _Bebop.contract.Transact(opts, "registerAllowedOrderSigner", signer, allowed)
}

// RegisterAllowedOrderSigner is a paid mutator transaction binding the contract method 0xea7faa61.
//
// Solidity: function registerAllowedOrderSigner(address signer, bool allowed) returns()
func (_Bebop *BebopSession) RegisterAllowedOrderSigner(signer common.Address, allowed bool) (*types.Transaction, error) {
	return _Bebop.Contract.RegisterAllowedOrderSigner(&_Bebop.TransactOpts, signer, allowed)
}

// RegisterAllowedOrderSigner is a paid mutator transaction binding the contract method 0xea7faa61.
//
// Solidity: function registerAllowedOrderSigner(address signer, bool allowed) returns()
func (_Bebop *BebopTransactorSession) RegisterAllowedOrderSigner(signer common.Address, allowed bool) (*types.Transaction, error) {
	return _Bebop.Contract.RegisterAllowedOrderSigner(&_Bebop.TransactOpts, signer, allowed)
}

// RegisterPartner is a paid mutator transaction binding the contract method 0x9080ef02.
//
// Solidity: function registerPartner(uint64 partnerId, uint16 fee, address beneficiary) returns()
func (_Bebop *BebopTransactor) RegisterPartner(opts *bind.TransactOpts, partnerId uint64, fee uint16, beneficiary common.Address) (*types.Transaction, error) {
	return _Bebop.contract.Transact(opts, "registerPartner", partnerId, fee, beneficiary)
}

// RegisterPartner is a paid mutator transaction binding the contract method 0x9080ef02.
//
// Solidity: function registerPartner(uint64 partnerId, uint16 fee, address beneficiary) returns()
func (_Bebop *BebopSession) RegisterPartner(partnerId uint64, fee uint16, beneficiary common.Address) (*types.Transaction, error) {
	return _Bebop.Contract.RegisterPartner(&_Bebop.TransactOpts, partnerId, fee, beneficiary)
}

// RegisterPartner is a paid mutator transaction binding the contract method 0x9080ef02.
//
// Solidity: function registerPartner(uint64 partnerId, uint16 fee, address beneficiary) returns()
func (_Bebop *BebopTransactorSession) RegisterPartner(partnerId uint64, fee uint16, beneficiary common.Address) (*types.Transaction, error) {
	return _Bebop.Contract.RegisterPartner(&_Bebop.TransactOpts, partnerId, fee, beneficiary)
}

// SettleAggregate is a paid mutator transaction binding the contract method 0x14e7a7ab.
//
// Solidity: function settleAggregate((uint256,address,address[],uint256[],address[][],address[][],uint256[][],uint256[][],address,bytes,uint256) order, (bytes,uint256)[] makersSignatures, uint256 filledTakerAmount, (bool,uint256[][],uint256[]) takerQuoteInfo, bytes takerSignature) payable returns()
func (_Bebop *BebopTransactor) SettleAggregate(opts *bind.TransactOpts, order OrderAggregate, makersSignatures []SignatureMakerSignature, filledTakerAmount *big.Int, takerQuoteInfo TransferOldAggregateQuote, takerSignature []byte) (*types.Transaction, error) {
	return _Bebop.contract.Transact(opts, "settleAggregate", order, makersSignatures, filledTakerAmount, takerQuoteInfo, takerSignature)
}

// SettleAggregate is a paid mutator transaction binding the contract method 0x14e7a7ab.
//
// Solidity: function settleAggregate((uint256,address,address[],uint256[],address[][],address[][],uint256[][],uint256[][],address,bytes,uint256) order, (bytes,uint256)[] makersSignatures, uint256 filledTakerAmount, (bool,uint256[][],uint256[]) takerQuoteInfo, bytes takerSignature) payable returns()
func (_Bebop *BebopSession) SettleAggregate(order OrderAggregate, makersSignatures []SignatureMakerSignature, filledTakerAmount *big.Int, takerQuoteInfo TransferOldAggregateQuote, takerSignature []byte) (*types.Transaction, error) {
	return _Bebop.Contract.SettleAggregate(&_Bebop.TransactOpts, order, makersSignatures, filledTakerAmount, takerQuoteInfo, takerSignature)
}

// SettleAggregate is a paid mutator transaction binding the contract method 0x14e7a7ab.
//
// Solidity: function settleAggregate((uint256,address,address[],uint256[],address[][],address[][],uint256[][],uint256[][],address,bytes,uint256) order, (bytes,uint256)[] makersSignatures, uint256 filledTakerAmount, (bool,uint256[][],uint256[]) takerQuoteInfo, bytes takerSignature) payable returns()
func (_Bebop *BebopTransactorSession) SettleAggregate(order OrderAggregate, makersSignatures []SignatureMakerSignature, filledTakerAmount *big.Int, takerQuoteInfo TransferOldAggregateQuote, takerSignature []byte) (*types.Transaction, error) {
	return _Bebop.Contract.SettleAggregate(&_Bebop.TransactOpts, order, makersSignatures, filledTakerAmount, takerQuoteInfo, takerSignature)
}

// SettleAggregateAndSignPermit is a paid mutator transaction binding the contract method 0x9093bf7b.
//
// Solidity: function settleAggregateAndSignPermit((uint256,address,address[],uint256[],address[][],address[][],uint256[][],uint256[][],address,bytes,uint256) order, (bytes,uint256)[] makersSignatures, uint256 filledTakerAmount, (bool,uint256[][],uint256[]) takerQuoteInfo, bytes takerSignature, (bytes,uint256) takerPermitSignature) payable returns()
func (_Bebop *BebopTransactor) SettleAggregateAndSignPermit(opts *bind.TransactOpts, order OrderAggregate, makersSignatures []SignatureMakerSignature, filledTakerAmount *big.Int, takerQuoteInfo TransferOldAggregateQuote, takerSignature []byte, takerPermitSignature SignaturePermitSignature) (*types.Transaction, error) {
	return _Bebop.contract.Transact(opts, "settleAggregateAndSignPermit", order, makersSignatures, filledTakerAmount, takerQuoteInfo, takerSignature, takerPermitSignature)
}

// SettleAggregateAndSignPermit is a paid mutator transaction binding the contract method 0x9093bf7b.
//
// Solidity: function settleAggregateAndSignPermit((uint256,address,address[],uint256[],address[][],address[][],uint256[][],uint256[][],address,bytes,uint256) order, (bytes,uint256)[] makersSignatures, uint256 filledTakerAmount, (bool,uint256[][],uint256[]) takerQuoteInfo, bytes takerSignature, (bytes,uint256) takerPermitSignature) payable returns()
func (_Bebop *BebopSession) SettleAggregateAndSignPermit(order OrderAggregate, makersSignatures []SignatureMakerSignature, filledTakerAmount *big.Int, takerQuoteInfo TransferOldAggregateQuote, takerSignature []byte, takerPermitSignature SignaturePermitSignature) (*types.Transaction, error) {
	return _Bebop.Contract.SettleAggregateAndSignPermit(&_Bebop.TransactOpts, order, makersSignatures, filledTakerAmount, takerQuoteInfo, takerSignature, takerPermitSignature)
}

// SettleAggregateAndSignPermit is a paid mutator transaction binding the contract method 0x9093bf7b.
//
// Solidity: function settleAggregateAndSignPermit((uint256,address,address[],uint256[],address[][],address[][],uint256[][],uint256[][],address,bytes,uint256) order, (bytes,uint256)[] makersSignatures, uint256 filledTakerAmount, (bool,uint256[][],uint256[]) takerQuoteInfo, bytes takerSignature, (bytes,uint256) takerPermitSignature) payable returns()
func (_Bebop *BebopTransactorSession) SettleAggregateAndSignPermit(order OrderAggregate, makersSignatures []SignatureMakerSignature, filledTakerAmount *big.Int, takerQuoteInfo TransferOldAggregateQuote, takerSignature []byte, takerPermitSignature SignaturePermitSignature) (*types.Transaction, error) {
	return _Bebop.Contract.SettleAggregateAndSignPermit(&_Bebop.TransactOpts, order, makersSignatures, filledTakerAmount, takerQuoteInfo, takerSignature, takerPermitSignature)
}

// SettleAggregateAndSignPermit2 is a paid mutator transaction binding the contract method 0x72fc32e5.
//
// Solidity: function settleAggregateAndSignPermit2((uint256,address,address[],uint256[],address[][],address[][],uint256[][],uint256[][],address,bytes,uint256) order, (bytes,uint256)[] makersSignatures, uint256 filledTakerAmount, (bool,uint256[][],uint256[]) takerQuoteInfo, bytes takerSignature, (bytes,uint48,uint48[]) infoPermit2) payable returns()
func (_Bebop *BebopTransactor) SettleAggregateAndSignPermit2(opts *bind.TransactOpts, order OrderAggregate, makersSignatures []SignatureMakerSignature, filledTakerAmount *big.Int, takerQuoteInfo TransferOldAggregateQuote, takerSignature []byte, infoPermit2 SignatureMultiTokensPermit2Signature) (*types.Transaction, error) {
	return _Bebop.contract.Transact(opts, "settleAggregateAndSignPermit2", order, makersSignatures, filledTakerAmount, takerQuoteInfo, takerSignature, infoPermit2)
}

// SettleAggregateAndSignPermit2 is a paid mutator transaction binding the contract method 0x72fc32e5.
//
// Solidity: function settleAggregateAndSignPermit2((uint256,address,address[],uint256[],address[][],address[][],uint256[][],uint256[][],address,bytes,uint256) order, (bytes,uint256)[] makersSignatures, uint256 filledTakerAmount, (bool,uint256[][],uint256[]) takerQuoteInfo, bytes takerSignature, (bytes,uint48,uint48[]) infoPermit2) payable returns()
func (_Bebop *BebopSession) SettleAggregateAndSignPermit2(order OrderAggregate, makersSignatures []SignatureMakerSignature, filledTakerAmount *big.Int, takerQuoteInfo TransferOldAggregateQuote, takerSignature []byte, infoPermit2 SignatureMultiTokensPermit2Signature) (*types.Transaction, error) {
	return _Bebop.Contract.SettleAggregateAndSignPermit2(&_Bebop.TransactOpts, order, makersSignatures, filledTakerAmount, takerQuoteInfo, takerSignature, infoPermit2)
}

// SettleAggregateAndSignPermit2 is a paid mutator transaction binding the contract method 0x72fc32e5.
//
// Solidity: function settleAggregateAndSignPermit2((uint256,address,address[],uint256[],address[][],address[][],uint256[][],uint256[][],address,bytes,uint256) order, (bytes,uint256)[] makersSignatures, uint256 filledTakerAmount, (bool,uint256[][],uint256[]) takerQuoteInfo, bytes takerSignature, (bytes,uint48,uint48[]) infoPermit2) payable returns()
func (_Bebop *BebopTransactorSession) SettleAggregateAndSignPermit2(order OrderAggregate, makersSignatures []SignatureMakerSignature, filledTakerAmount *big.Int, takerQuoteInfo TransferOldAggregateQuote, takerSignature []byte, infoPermit2 SignatureMultiTokensPermit2Signature) (*types.Transaction, error) {
	return _Bebop.Contract.SettleAggregateAndSignPermit2(&_Bebop.TransactOpts, order, makersSignatures, filledTakerAmount, takerQuoteInfo, takerSignature, infoPermit2)
}

// SettleMulti is a paid mutator transaction binding the contract method 0xefe34fe6.
//
// Solidity: function settleMulti((uint256,address,address,uint256,address[],address[],uint256[],uint256[],address,bytes,uint256) order, (bytes,uint256) makerSignature, uint256 filledTakerAmount, (bool,uint256[],uint256) takerQuoteInfo, bytes takerSignature) payable returns()
func (_Bebop *BebopTransactor) SettleMulti(opts *bind.TransactOpts, order OrderMulti, makerSignature SignatureMakerSignature, filledTakerAmount *big.Int, takerQuoteInfo TransferOldMultiQuote, takerSignature []byte) (*types.Transaction, error) {
	return _Bebop.contract.Transact(opts, "settleMulti", order, makerSignature, filledTakerAmount, takerQuoteInfo, takerSignature)
}

// SettleMulti is a paid mutator transaction binding the contract method 0xefe34fe6.
//
// Solidity: function settleMulti((uint256,address,address,uint256,address[],address[],uint256[],uint256[],address,bytes,uint256) order, (bytes,uint256) makerSignature, uint256 filledTakerAmount, (bool,uint256[],uint256) takerQuoteInfo, bytes takerSignature) payable returns()
func (_Bebop *BebopSession) SettleMulti(order OrderMulti, makerSignature SignatureMakerSignature, filledTakerAmount *big.Int, takerQuoteInfo TransferOldMultiQuote, takerSignature []byte) (*types.Transaction, error) {
	return _Bebop.Contract.SettleMulti(&_Bebop.TransactOpts, order, makerSignature, filledTakerAmount, takerQuoteInfo, takerSignature)
}

// SettleMulti is a paid mutator transaction binding the contract method 0xefe34fe6.
//
// Solidity: function settleMulti((uint256,address,address,uint256,address[],address[],uint256[],uint256[],address,bytes,uint256) order, (bytes,uint256) makerSignature, uint256 filledTakerAmount, (bool,uint256[],uint256) takerQuoteInfo, bytes takerSignature) payable returns()
func (_Bebop *BebopTransactorSession) SettleMulti(order OrderMulti, makerSignature SignatureMakerSignature, filledTakerAmount *big.Int, takerQuoteInfo TransferOldMultiQuote, takerSignature []byte) (*types.Transaction, error) {
	return _Bebop.Contract.SettleMulti(&_Bebop.TransactOpts, order, makerSignature, filledTakerAmount, takerQuoteInfo, takerSignature)
}

// SettleMultiAndSignPermit is a paid mutator transaction binding the contract method 0x384eada0.
//
// Solidity: function settleMultiAndSignPermit((uint256,address,address,uint256,address[],address[],uint256[],uint256[],address,bytes,uint256) order, (bytes,uint256) makerSignature, uint256 filledTakerAmount, (bool,uint256[],uint256) takerQuoteInfo, bytes takerSignature, (bytes,uint256) takerPermitSignature) payable returns()
func (_Bebop *BebopTransactor) SettleMultiAndSignPermit(opts *bind.TransactOpts, order OrderMulti, makerSignature SignatureMakerSignature, filledTakerAmount *big.Int, takerQuoteInfo TransferOldMultiQuote, takerSignature []byte, takerPermitSignature SignaturePermitSignature) (*types.Transaction, error) {
	return _Bebop.contract.Transact(opts, "settleMultiAndSignPermit", order, makerSignature, filledTakerAmount, takerQuoteInfo, takerSignature, takerPermitSignature)
}

// SettleMultiAndSignPermit is a paid mutator transaction binding the contract method 0x384eada0.
//
// Solidity: function settleMultiAndSignPermit((uint256,address,address,uint256,address[],address[],uint256[],uint256[],address,bytes,uint256) order, (bytes,uint256) makerSignature, uint256 filledTakerAmount, (bool,uint256[],uint256) takerQuoteInfo, bytes takerSignature, (bytes,uint256) takerPermitSignature) payable returns()
func (_Bebop *BebopSession) SettleMultiAndSignPermit(order OrderMulti, makerSignature SignatureMakerSignature, filledTakerAmount *big.Int, takerQuoteInfo TransferOldMultiQuote, takerSignature []byte, takerPermitSignature SignaturePermitSignature) (*types.Transaction, error) {
	return _Bebop.Contract.SettleMultiAndSignPermit(&_Bebop.TransactOpts, order, makerSignature, filledTakerAmount, takerQuoteInfo, takerSignature, takerPermitSignature)
}

// SettleMultiAndSignPermit is a paid mutator transaction binding the contract method 0x384eada0.
//
// Solidity: function settleMultiAndSignPermit((uint256,address,address,uint256,address[],address[],uint256[],uint256[],address,bytes,uint256) order, (bytes,uint256) makerSignature, uint256 filledTakerAmount, (bool,uint256[],uint256) takerQuoteInfo, bytes takerSignature, (bytes,uint256) takerPermitSignature) payable returns()
func (_Bebop *BebopTransactorSession) SettleMultiAndSignPermit(order OrderMulti, makerSignature SignatureMakerSignature, filledTakerAmount *big.Int, takerQuoteInfo TransferOldMultiQuote, takerSignature []byte, takerPermitSignature SignaturePermitSignature) (*types.Transaction, error) {
	return _Bebop.Contract.SettleMultiAndSignPermit(&_Bebop.TransactOpts, order, makerSignature, filledTakerAmount, takerQuoteInfo, takerSignature, takerPermitSignature)
}

// SettleMultiAndSignPermit2 is a paid mutator transaction binding the contract method 0x66a65e41.
//
// Solidity: function settleMultiAndSignPermit2((uint256,address,address,uint256,address[],address[],uint256[],uint256[],address,bytes,uint256) order, (bytes,uint256) makerSignature, uint256 filledTakerAmount, (bool,uint256[],uint256) takerQuoteInfo, bytes takerSignature, (bytes,uint48,uint48[]) infoPermit2) payable returns()
func (_Bebop *BebopTransactor) SettleMultiAndSignPermit2(opts *bind.TransactOpts, order OrderMulti, makerSignature SignatureMakerSignature, filledTakerAmount *big.Int, takerQuoteInfo TransferOldMultiQuote, takerSignature []byte, infoPermit2 SignatureMultiTokensPermit2Signature) (*types.Transaction, error) {
	return _Bebop.contract.Transact(opts, "settleMultiAndSignPermit2", order, makerSignature, filledTakerAmount, takerQuoteInfo, takerSignature, infoPermit2)
}

// SettleMultiAndSignPermit2 is a paid mutator transaction binding the contract method 0x66a65e41.
//
// Solidity: function settleMultiAndSignPermit2((uint256,address,address,uint256,address[],address[],uint256[],uint256[],address,bytes,uint256) order, (bytes,uint256) makerSignature, uint256 filledTakerAmount, (bool,uint256[],uint256) takerQuoteInfo, bytes takerSignature, (bytes,uint48,uint48[]) infoPermit2) payable returns()
func (_Bebop *BebopSession) SettleMultiAndSignPermit2(order OrderMulti, makerSignature SignatureMakerSignature, filledTakerAmount *big.Int, takerQuoteInfo TransferOldMultiQuote, takerSignature []byte, infoPermit2 SignatureMultiTokensPermit2Signature) (*types.Transaction, error) {
	return _Bebop.Contract.SettleMultiAndSignPermit2(&_Bebop.TransactOpts, order, makerSignature, filledTakerAmount, takerQuoteInfo, takerSignature, infoPermit2)
}

// SettleMultiAndSignPermit2 is a paid mutator transaction binding the contract method 0x66a65e41.
//
// Solidity: function settleMultiAndSignPermit2((uint256,address,address,uint256,address[],address[],uint256[],uint256[],address,bytes,uint256) order, (bytes,uint256) makerSignature, uint256 filledTakerAmount, (bool,uint256[],uint256) takerQuoteInfo, bytes takerSignature, (bytes,uint48,uint48[]) infoPermit2) payable returns()
func (_Bebop *BebopTransactorSession) SettleMultiAndSignPermit2(order OrderMulti, makerSignature SignatureMakerSignature, filledTakerAmount *big.Int, takerQuoteInfo TransferOldMultiQuote, takerSignature []byte, infoPermit2 SignatureMultiTokensPermit2Signature) (*types.Transaction, error) {
	return _Bebop.Contract.SettleMultiAndSignPermit2(&_Bebop.TransactOpts, order, makerSignature, filledTakerAmount, takerQuoteInfo, takerSignature, infoPermit2)
}

// SettleSingle is a paid mutator transaction binding the contract method 0x1a499026.
//
// Solidity: function settleSingle((uint256,address,address,uint256,address,address,uint256,uint256,address,uint256,uint256) order, (bytes,uint256) makerSignature, uint256 filledTakerAmount, (bool,uint256,uint256) takerQuoteInfo, bytes takerSignature) payable returns()
func (_Bebop *BebopTransactor) SettleSingle(opts *bind.TransactOpts, order OrderSingle, makerSignature SignatureMakerSignature, filledTakerAmount *big.Int, takerQuoteInfo TransferOldSingleQuote, takerSignature []byte) (*types.Transaction, error) {
	return _Bebop.contract.Transact(opts, "settleSingle", order, makerSignature, filledTakerAmount, takerQuoteInfo, takerSignature)
}

// SettleSingle is a paid mutator transaction binding the contract method 0x1a499026.
//
// Solidity: function settleSingle((uint256,address,address,uint256,address,address,uint256,uint256,address,uint256,uint256) order, (bytes,uint256) makerSignature, uint256 filledTakerAmount, (bool,uint256,uint256) takerQuoteInfo, bytes takerSignature) payable returns()
func (_Bebop *BebopSession) SettleSingle(order OrderSingle, makerSignature SignatureMakerSignature, filledTakerAmount *big.Int, takerQuoteInfo TransferOldSingleQuote, takerSignature []byte) (*types.Transaction, error) {
	return _Bebop.Contract.SettleSingle(&_Bebop.TransactOpts, order, makerSignature, filledTakerAmount, takerQuoteInfo, takerSignature)
}

// SettleSingle is a paid mutator transaction binding the contract method 0x1a499026.
//
// Solidity: function settleSingle((uint256,address,address,uint256,address,address,uint256,uint256,address,uint256,uint256) order, (bytes,uint256) makerSignature, uint256 filledTakerAmount, (bool,uint256,uint256) takerQuoteInfo, bytes takerSignature) payable returns()
func (_Bebop *BebopTransactorSession) SettleSingle(order OrderSingle, makerSignature SignatureMakerSignature, filledTakerAmount *big.Int, takerQuoteInfo TransferOldSingleQuote, takerSignature []byte) (*types.Transaction, error) {
	return _Bebop.Contract.SettleSingle(&_Bebop.TransactOpts, order, makerSignature, filledTakerAmount, takerQuoteInfo, takerSignature)
}

// SettleSingleAndSignPermit is a paid mutator transaction binding the contract method 0xef7d27ad.
//
// Solidity: function settleSingleAndSignPermit((uint256,address,address,uint256,address,address,uint256,uint256,address,uint256,uint256) order, (bytes,uint256) makerSignature, uint256 filledTakerAmount, (bool,uint256,uint256) takerQuoteInfo, bytes takerSignature, (bytes,uint256) takerPermitSignature) payable returns()
func (_Bebop *BebopTransactor) SettleSingleAndSignPermit(opts *bind.TransactOpts, order OrderSingle, makerSignature SignatureMakerSignature, filledTakerAmount *big.Int, takerQuoteInfo TransferOldSingleQuote, takerSignature []byte, takerPermitSignature SignaturePermitSignature) (*types.Transaction, error) {
	return _Bebop.contract.Transact(opts, "settleSingleAndSignPermit", order, makerSignature, filledTakerAmount, takerQuoteInfo, takerSignature, takerPermitSignature)
}

// SettleSingleAndSignPermit is a paid mutator transaction binding the contract method 0xef7d27ad.
//
// Solidity: function settleSingleAndSignPermit((uint256,address,address,uint256,address,address,uint256,uint256,address,uint256,uint256) order, (bytes,uint256) makerSignature, uint256 filledTakerAmount, (bool,uint256,uint256) takerQuoteInfo, bytes takerSignature, (bytes,uint256) takerPermitSignature) payable returns()
func (_Bebop *BebopSession) SettleSingleAndSignPermit(order OrderSingle, makerSignature SignatureMakerSignature, filledTakerAmount *big.Int, takerQuoteInfo TransferOldSingleQuote, takerSignature []byte, takerPermitSignature SignaturePermitSignature) (*types.Transaction, error) {
	return _Bebop.Contract.SettleSingleAndSignPermit(&_Bebop.TransactOpts, order, makerSignature, filledTakerAmount, takerQuoteInfo, takerSignature, takerPermitSignature)
}

// SettleSingleAndSignPermit is a paid mutator transaction binding the contract method 0xef7d27ad.
//
// Solidity: function settleSingleAndSignPermit((uint256,address,address,uint256,address,address,uint256,uint256,address,uint256,uint256) order, (bytes,uint256) makerSignature, uint256 filledTakerAmount, (bool,uint256,uint256) takerQuoteInfo, bytes takerSignature, (bytes,uint256) takerPermitSignature) payable returns()
func (_Bebop *BebopTransactorSession) SettleSingleAndSignPermit(order OrderSingle, makerSignature SignatureMakerSignature, filledTakerAmount *big.Int, takerQuoteInfo TransferOldSingleQuote, takerSignature []byte, takerPermitSignature SignaturePermitSignature) (*types.Transaction, error) {
	return _Bebop.Contract.SettleSingleAndSignPermit(&_Bebop.TransactOpts, order, makerSignature, filledTakerAmount, takerQuoteInfo, takerSignature, takerPermitSignature)
}

// SettleSingleAndSignPermit2 is a paid mutator transaction binding the contract method 0x38ec0211.
//
// Solidity: function settleSingleAndSignPermit2((uint256,address,address,uint256,address,address,uint256,uint256,address,uint256,uint256) order, (bytes,uint256) makerSignature, uint256 filledTakerAmount, (bool,uint256,uint256) takerQuoteInfo, bytes takerSignature, (bytes,uint48,uint48) takerPermit2Signature) payable returns()
func (_Bebop *BebopTransactor) SettleSingleAndSignPermit2(opts *bind.TransactOpts, order OrderSingle, makerSignature SignatureMakerSignature, filledTakerAmount *big.Int, takerQuoteInfo TransferOldSingleQuote, takerSignature []byte, takerPermit2Signature SignaturePermit2Signature) (*types.Transaction, error) {
	return _Bebop.contract.Transact(opts, "settleSingleAndSignPermit2", order, makerSignature, filledTakerAmount, takerQuoteInfo, takerSignature, takerPermit2Signature)
}

// SettleSingleAndSignPermit2 is a paid mutator transaction binding the contract method 0x38ec0211.
//
// Solidity: function settleSingleAndSignPermit2((uint256,address,address,uint256,address,address,uint256,uint256,address,uint256,uint256) order, (bytes,uint256) makerSignature, uint256 filledTakerAmount, (bool,uint256,uint256) takerQuoteInfo, bytes takerSignature, (bytes,uint48,uint48) takerPermit2Signature) payable returns()
func (_Bebop *BebopSession) SettleSingleAndSignPermit2(order OrderSingle, makerSignature SignatureMakerSignature, filledTakerAmount *big.Int, takerQuoteInfo TransferOldSingleQuote, takerSignature []byte, takerPermit2Signature SignaturePermit2Signature) (*types.Transaction, error) {
	return _Bebop.Contract.SettleSingleAndSignPermit2(&_Bebop.TransactOpts, order, makerSignature, filledTakerAmount, takerQuoteInfo, takerSignature, takerPermit2Signature)
}

// SettleSingleAndSignPermit2 is a paid mutator transaction binding the contract method 0x38ec0211.
//
// Solidity: function settleSingleAndSignPermit2((uint256,address,address,uint256,address,address,uint256,uint256,address,uint256,uint256) order, (bytes,uint256) makerSignature, uint256 filledTakerAmount, (bool,uint256,uint256) takerQuoteInfo, bytes takerSignature, (bytes,uint48,uint48) takerPermit2Signature) payable returns()
func (_Bebop *BebopTransactorSession) SettleSingleAndSignPermit2(order OrderSingle, makerSignature SignatureMakerSignature, filledTakerAmount *big.Int, takerQuoteInfo TransferOldSingleQuote, takerSignature []byte, takerPermit2Signature SignaturePermit2Signature) (*types.Transaction, error) {
	return _Bebop.Contract.SettleSingleAndSignPermit2(&_Bebop.TransactOpts, order, makerSignature, filledTakerAmount, takerQuoteInfo, takerSignature, takerPermit2Signature)
}

// SwapAggregate is a paid mutator transaction binding the contract method 0xa2f74893.
//
// Solidity: function swapAggregate((uint256,address,address[],uint256[],address[][],address[][],uint256[][],uint256[][],address,bytes,uint256) order, (bytes,uint256)[] makersSignatures, uint256 filledTakerAmount) payable returns()
func (_Bebop *BebopTransactor) SwapAggregate(opts *bind.TransactOpts, order OrderAggregate, makersSignatures []SignatureMakerSignature, filledTakerAmount *big.Int) (*types.Transaction, error) {
	return _Bebop.contract.Transact(opts, "swapAggregate", order, makersSignatures, filledTakerAmount)
}

// SwapAggregate is a paid mutator transaction binding the contract method 0xa2f74893.
//
// Solidity: function swapAggregate((uint256,address,address[],uint256[],address[][],address[][],uint256[][],uint256[][],address,bytes,uint256) order, (bytes,uint256)[] makersSignatures, uint256 filledTakerAmount) payable returns()
func (_Bebop *BebopSession) SwapAggregate(order OrderAggregate, makersSignatures []SignatureMakerSignature, filledTakerAmount *big.Int) (*types.Transaction, error) {
	return _Bebop.Contract.SwapAggregate(&_Bebop.TransactOpts, order, makersSignatures, filledTakerAmount)
}

// SwapAggregate is a paid mutator transaction binding the contract method 0xa2f74893.
//
// Solidity: function swapAggregate((uint256,address,address[],uint256[],address[][],address[][],uint256[][],uint256[][],address,bytes,uint256) order, (bytes,uint256)[] makersSignatures, uint256 filledTakerAmount) payable returns()
func (_Bebop *BebopTransactorSession) SwapAggregate(order OrderAggregate, makersSignatures []SignatureMakerSignature, filledTakerAmount *big.Int) (*types.Transaction, error) {
	return _Bebop.Contract.SwapAggregate(&_Bebop.TransactOpts, order, makersSignatures, filledTakerAmount)
}

// SwapMulti is a paid mutator transaction binding the contract method 0x602926bf.
//
// Solidity: function swapMulti((uint256,address,address,uint256,address[],address[],uint256[],uint256[],address,bytes,uint256) order, (bytes,uint256) makerSignature, uint256 filledTakerAmount) payable returns()
func (_Bebop *BebopTransactor) SwapMulti(opts *bind.TransactOpts, order OrderMulti, makerSignature SignatureMakerSignature, filledTakerAmount *big.Int) (*types.Transaction, error) {
	return _Bebop.contract.Transact(opts, "swapMulti", order, makerSignature, filledTakerAmount)
}

// SwapMulti is a paid mutator transaction binding the contract method 0x602926bf.
//
// Solidity: function swapMulti((uint256,address,address,uint256,address[],address[],uint256[],uint256[],address,bytes,uint256) order, (bytes,uint256) makerSignature, uint256 filledTakerAmount) payable returns()
func (_Bebop *BebopSession) SwapMulti(order OrderMulti, makerSignature SignatureMakerSignature, filledTakerAmount *big.Int) (*types.Transaction, error) {
	return _Bebop.Contract.SwapMulti(&_Bebop.TransactOpts, order, makerSignature, filledTakerAmount)
}

// SwapMulti is a paid mutator transaction binding the contract method 0x602926bf.
//
// Solidity: function swapMulti((uint256,address,address,uint256,address[],address[],uint256[],uint256[],address,bytes,uint256) order, (bytes,uint256) makerSignature, uint256 filledTakerAmount) payable returns()
func (_Bebop *BebopTransactorSession) SwapMulti(order OrderMulti, makerSignature SignatureMakerSignature, filledTakerAmount *big.Int) (*types.Transaction, error) {
	return _Bebop.Contract.SwapMulti(&_Bebop.TransactOpts, order, makerSignature, filledTakerAmount)
}

// SwapSingle is a paid mutator transaction binding the contract method 0x4dcebcba.
//
// Solidity: function swapSingle((uint256,address,address,uint256,address,address,uint256,uint256,address,uint256,uint256) order, (bytes,uint256) makerSignature, uint256 filledTakerAmount) payable returns()
func (_Bebop *BebopTransactor) SwapSingle(opts *bind.TransactOpts, order OrderSingle, makerSignature SignatureMakerSignature, filledTakerAmount *big.Int) (*types.Transaction, error) {
	return _Bebop.contract.Transact(opts, "swapSingle", order, makerSignature, filledTakerAmount)
}

// SwapSingle is a paid mutator transaction binding the contract method 0x4dcebcba.
//
// Solidity: function swapSingle((uint256,address,address,uint256,address,address,uint256,uint256,address,uint256,uint256) order, (bytes,uint256) makerSignature, uint256 filledTakerAmount) payable returns()
func (_Bebop *BebopSession) SwapSingle(order OrderSingle, makerSignature SignatureMakerSignature, filledTakerAmount *big.Int) (*types.Transaction, error) {
	return _Bebop.Contract.SwapSingle(&_Bebop.TransactOpts, order, makerSignature, filledTakerAmount)
}

// SwapSingle is a paid mutator transaction binding the contract method 0x4dcebcba.
//
// Solidity: function swapSingle((uint256,address,address,uint256,address,address,uint256,uint256,address,uint256,uint256) order, (bytes,uint256) makerSignature, uint256 filledTakerAmount) payable returns()
func (_Bebop *BebopTransactorSession) SwapSingle(order OrderSingle, makerSignature SignatureMakerSignature, filledTakerAmount *big.Int) (*types.Transaction, error) {
	return _Bebop.Contract.SwapSingle(&_Bebop.TransactOpts, order, makerSignature, filledTakerAmount)
}

// SwapSingleFromContract is a paid mutator transaction binding the contract method 0xc38a4474.
//
// Solidity: function swapSingleFromContract((uint256,address,address,uint256,address,address,uint256,uint256,address,uint256,uint256) order, (bytes,uint256) makerSignature) payable returns()
func (_Bebop *BebopTransactor) SwapSingleFromContract(opts *bind.TransactOpts, order OrderSingle, makerSignature SignatureMakerSignature) (*types.Transaction, error) {
	return _Bebop.contract.Transact(opts, "swapSingleFromContract", order, makerSignature)
}

// SwapSingleFromContract is a paid mutator transaction binding the contract method 0xc38a4474.
//
// Solidity: function swapSingleFromContract((uint256,address,address,uint256,address,address,uint256,uint256,address,uint256,uint256) order, (bytes,uint256) makerSignature) payable returns()
func (_Bebop *BebopSession) SwapSingleFromContract(order OrderSingle, makerSignature SignatureMakerSignature) (*types.Transaction, error) {
	return _Bebop.Contract.SwapSingleFromContract(&_Bebop.TransactOpts, order, makerSignature)
}

// SwapSingleFromContract is a paid mutator transaction binding the contract method 0xc38a4474.
//
// Solidity: function swapSingleFromContract((uint256,address,address,uint256,address,address,uint256,uint256,address,uint256,uint256) order, (bytes,uint256) makerSignature) payable returns()
func (_Bebop *BebopTransactorSession) SwapSingleFromContract(order OrderSingle, makerSignature SignatureMakerSignature) (*types.Transaction, error) {
	return _Bebop.Contract.SwapSingleFromContract(&_Bebop.TransactOpts, order, makerSignature)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Bebop *BebopTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bebop.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Bebop *BebopSession) Receive() (*types.Transaction, error) {
	return _Bebop.Contract.Receive(&_Bebop.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Bebop *BebopTransactorSession) Receive() (*types.Transaction, error) {
	return _Bebop.Contract.Receive(&_Bebop.TransactOpts)
}

// BebopBebopOrderIterator is returned from FilterBebopOrder and is used to iterate over the raw logs and unpacked data for BebopOrder events raised by the Bebop contract.
type BebopBebopOrderIterator struct {
	Event *BebopBebopOrder // Event containing the contract specifics and raw log

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
func (it *BebopBebopOrderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BebopBebopOrder)
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
		it.Event = new(BebopBebopOrder)
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
func (it *BebopBebopOrderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BebopBebopOrderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BebopBebopOrder represents a BebopOrder event raised by the Bebop contract.
type BebopBebopOrder struct {
	EventId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBebopOrder is a free log retrieval operation binding the contract event 0xadd7095becdaa725f0f33243630938c861b0bba83dfd217d4055701aa768ec2e.
//
// Solidity: event BebopOrder(uint128 indexed eventId)
func (_Bebop *BebopFilterer) FilterBebopOrder(opts *bind.FilterOpts, eventId []*big.Int) (*BebopBebopOrderIterator, error) {

	var eventIdRule []interface{}
	for _, eventIdItem := range eventId {
		eventIdRule = append(eventIdRule, eventIdItem)
	}

	logs, sub, err := _Bebop.contract.FilterLogs(opts, "BebopOrder", eventIdRule)
	if err != nil {
		return nil, err
	}
	return &BebopBebopOrderIterator{contract: _Bebop.contract, event: "BebopOrder", logs: logs, sub: sub}, nil
}

// WatchBebopOrder is a free log subscription operation binding the contract event 0xadd7095becdaa725f0f33243630938c861b0bba83dfd217d4055701aa768ec2e.
//
// Solidity: event BebopOrder(uint128 indexed eventId)
func (_Bebop *BebopFilterer) WatchBebopOrder(opts *bind.WatchOpts, sink chan<- *BebopBebopOrder, eventId []*big.Int) (event.Subscription, error) {

	var eventIdRule []interface{}
	for _, eventIdItem := range eventId {
		eventIdRule = append(eventIdRule, eventIdItem)
	}

	logs, sub, err := _Bebop.contract.WatchLogs(opts, "BebopOrder", eventIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BebopBebopOrder)
				if err := _Bebop.contract.UnpackLog(event, "BebopOrder", log); err != nil {
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

// ParseBebopOrder is a log parse operation binding the contract event 0xadd7095becdaa725f0f33243630938c861b0bba83dfd217d4055701aa768ec2e.
//
// Solidity: event BebopOrder(uint128 indexed eventId)
func (_Bebop *BebopFilterer) ParseBebopOrder(log types.Log) (*BebopBebopOrder, error) {
	event := new(BebopBebopOrder)
	if err := _Bebop.contract.UnpackLog(event, "BebopOrder", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BebopOrderSignerRegisteredIterator is returned from FilterOrderSignerRegistered and is used to iterate over the raw logs and unpacked data for OrderSignerRegistered events raised by the Bebop contract.
type BebopOrderSignerRegisteredIterator struct {
	Event *BebopOrderSignerRegistered // Event containing the contract specifics and raw log

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
func (it *BebopOrderSignerRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BebopOrderSignerRegistered)
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
		it.Event = new(BebopOrderSignerRegistered)
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
func (it *BebopOrderSignerRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BebopOrderSignerRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BebopOrderSignerRegistered represents a OrderSignerRegistered event raised by the Bebop contract.
type BebopOrderSignerRegistered struct {
	Maker   common.Address
	Signer  common.Address
	Allowed bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterOrderSignerRegistered is a free log retrieval operation binding the contract event 0x6ea9dbe8b2cc119348716a9220a0742ad62b7884ecb0ff4b32cd508121fd9379.
//
// Solidity: event OrderSignerRegistered(address maker, address signer, bool allowed)
func (_Bebop *BebopFilterer) FilterOrderSignerRegistered(opts *bind.FilterOpts) (*BebopOrderSignerRegisteredIterator, error) {

	logs, sub, err := _Bebop.contract.FilterLogs(opts, "OrderSignerRegistered")
	if err != nil {
		return nil, err
	}
	return &BebopOrderSignerRegisteredIterator{contract: _Bebop.contract, event: "OrderSignerRegistered", logs: logs, sub: sub}, nil
}

// WatchOrderSignerRegistered is a free log subscription operation binding the contract event 0x6ea9dbe8b2cc119348716a9220a0742ad62b7884ecb0ff4b32cd508121fd9379.
//
// Solidity: event OrderSignerRegistered(address maker, address signer, bool allowed)
func (_Bebop *BebopFilterer) WatchOrderSignerRegistered(opts *bind.WatchOpts, sink chan<- *BebopOrderSignerRegistered) (event.Subscription, error) {

	logs, sub, err := _Bebop.contract.WatchLogs(opts, "OrderSignerRegistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BebopOrderSignerRegistered)
				if err := _Bebop.contract.UnpackLog(event, "OrderSignerRegistered", log); err != nil {
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

// ParseOrderSignerRegistered is a log parse operation binding the contract event 0x6ea9dbe8b2cc119348716a9220a0742ad62b7884ecb0ff4b32cd508121fd9379.
//
// Solidity: event OrderSignerRegistered(address maker, address signer, bool allowed)
func (_Bebop *BebopFilterer) ParseOrderSignerRegistered(log types.Log) (*BebopOrderSignerRegistered, error) {
	event := new(BebopOrderSignerRegistered)
	if err := _Bebop.contract.UnpackLog(event, "OrderSignerRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
