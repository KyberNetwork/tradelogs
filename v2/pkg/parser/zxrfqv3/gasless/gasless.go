// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package gasless

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

// SettlerBaseAllowedSlippage is an auto generated low-level Go binding around an user-defined struct.
type SettlerBaseAllowedSlippage struct {
	Recipient    common.Address
	BuyToken     common.Address
	MinAmountOut *big.Int
}

// GaslessMetaData contains all meta data concerning the Gasless contract.
var GaslessMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes20\",\"name\":\"gitCommit\",\"type\":\"bytes20\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"},{\"internalType\":\"bytes4\",\"name\":\"action\",\"type\":\"bytes4\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"ActionInvalid\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"callbackInt\",\"type\":\"uint256\"}],\"name\":\"CallbackNotSpent\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ConfusedDeputy\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForwarderNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidOffset\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTarget\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotConverged\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PayerSpent\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"callbackInt\",\"type\":\"uint256\"}],\"name\":\"ReentrantCallback\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"oldWitness\",\"type\":\"bytes32\"}],\"name\":\"ReentrantMetatransaction\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"oldPayer\",\"type\":\"address\"}],\"name\":\"ReentrantPayer\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"TooMuchSlippage\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"forkId\",\"type\":\"uint8\"}],\"name\":\"UnknownForkId\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"oldWitness\",\"type\":\"bytes32\"}],\"name\":\"WitnessNotSpent\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes20\",\"name\":\"\",\"type\":\"bytes20\"}],\"name\":\"GitCommit\",\"type\":\"event\"},{\"stateMutability\":\"nonpayable\",\"type\":\"fallback\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"buyToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minAmountOut\",\"type\":\"uint256\"}],\"internalType\":\"structSettlerBase.AllowedSlippage\",\"name\":\"slippage\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"actions\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"msgSender\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"executeMetaTxn\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// GaslessABI is the input ABI used to generate the binding from.
// Deprecated: Use GaslessMetaData.ABI instead.
var GaslessABI = GaslessMetaData.ABI

// Gasless is an auto generated Go binding around an Ethereum contract.
type Gasless struct {
	GaslessCaller     // Read-only binding to the contract
	GaslessTransactor // Write-only binding to the contract
	GaslessFilterer   // Log filterer for contract events
}

// GaslessCaller is an auto generated read-only Go binding around an Ethereum contract.
type GaslessCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GaslessTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GaslessTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GaslessFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GaslessFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GaslessSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GaslessSession struct {
	Contract     *Gasless          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GaslessCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GaslessCallerSession struct {
	Contract *GaslessCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// GaslessTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GaslessTransactorSession struct {
	Contract     *GaslessTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// GaslessRaw is an auto generated low-level Go binding around an Ethereum contract.
type GaslessRaw struct {
	Contract *Gasless // Generic contract binding to access the raw methods on
}

// GaslessCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GaslessCallerRaw struct {
	Contract *GaslessCaller // Generic read-only contract binding to access the raw methods on
}

// GaslessTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GaslessTransactorRaw struct {
	Contract *GaslessTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGasless creates a new instance of Gasless, bound to a specific deployed contract.
func NewGasless(address common.Address, backend bind.ContractBackend) (*Gasless, error) {
	contract, err := bindGasless(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Gasless{GaslessCaller: GaslessCaller{contract: contract}, GaslessTransactor: GaslessTransactor{contract: contract}, GaslessFilterer: GaslessFilterer{contract: contract}}, nil
}

// NewGaslessCaller creates a new read-only instance of Gasless, bound to a specific deployed contract.
func NewGaslessCaller(address common.Address, caller bind.ContractCaller) (*GaslessCaller, error) {
	contract, err := bindGasless(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GaslessCaller{contract: contract}, nil
}

// NewGaslessTransactor creates a new write-only instance of Gasless, bound to a specific deployed contract.
func NewGaslessTransactor(address common.Address, transactor bind.ContractTransactor) (*GaslessTransactor, error) {
	contract, err := bindGasless(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GaslessTransactor{contract: contract}, nil
}

// NewGaslessFilterer creates a new log filterer instance of Gasless, bound to a specific deployed contract.
func NewGaslessFilterer(address common.Address, filterer bind.ContractFilterer) (*GaslessFilterer, error) {
	contract, err := bindGasless(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GaslessFilterer{contract: contract}, nil
}

// bindGasless binds a generic wrapper to an already deployed contract.
func bindGasless(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := GaslessMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Gasless *GaslessRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Gasless.Contract.GaslessCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Gasless *GaslessRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Gasless.Contract.GaslessTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Gasless *GaslessRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Gasless.Contract.GaslessTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Gasless *GaslessCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Gasless.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Gasless *GaslessTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Gasless.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Gasless *GaslessTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Gasless.Contract.contract.Transact(opts, method, params...)
}

// ExecuteMetaTxn is a paid mutator transaction binding the contract method 0xfd3ad6d4.
//
// Solidity: function executeMetaTxn((address,address,uint256) slippage, bytes[] actions, bytes32 , address msgSender, bytes sig) returns(bool)
func (_Gasless *GaslessTransactor) ExecuteMetaTxn(opts *bind.TransactOpts, slippage SettlerBaseAllowedSlippage, actions [][]byte, arg2 [32]byte, msgSender common.Address, sig []byte) (*types.Transaction, error) {
	return _Gasless.contract.Transact(opts, "executeMetaTxn", slippage, actions, arg2, msgSender, sig)
}

// ExecuteMetaTxn is a paid mutator transaction binding the contract method 0xfd3ad6d4.
//
// Solidity: function executeMetaTxn((address,address,uint256) slippage, bytes[] actions, bytes32 , address msgSender, bytes sig) returns(bool)
func (_Gasless *GaslessSession) ExecuteMetaTxn(slippage SettlerBaseAllowedSlippage, actions [][]byte, arg2 [32]byte, msgSender common.Address, sig []byte) (*types.Transaction, error) {
	return _Gasless.Contract.ExecuteMetaTxn(&_Gasless.TransactOpts, slippage, actions, arg2, msgSender, sig)
}

// ExecuteMetaTxn is a paid mutator transaction binding the contract method 0xfd3ad6d4.
//
// Solidity: function executeMetaTxn((address,address,uint256) slippage, bytes[] actions, bytes32 , address msgSender, bytes sig) returns(bool)
func (_Gasless *GaslessTransactorSession) ExecuteMetaTxn(slippage SettlerBaseAllowedSlippage, actions [][]byte, arg2 [32]byte, msgSender common.Address, sig []byte) (*types.Transaction, error) {
	return _Gasless.Contract.ExecuteMetaTxn(&_Gasless.TransactOpts, slippage, actions, arg2, msgSender, sig)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_Gasless *GaslessTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Gasless.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_Gasless *GaslessSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Gasless.Contract.Fallback(&_Gasless.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_Gasless *GaslessTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Gasless.Contract.Fallback(&_Gasless.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Gasless *GaslessTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Gasless.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Gasless *GaslessSession) Receive() (*types.Transaction, error) {
	return _Gasless.Contract.Receive(&_Gasless.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Gasless *GaslessTransactorSession) Receive() (*types.Transaction, error) {
	return _Gasless.Contract.Receive(&_Gasless.TransactOpts)
}

// GaslessGitCommitIterator is returned from FilterGitCommit and is used to iterate over the raw logs and unpacked data for GitCommit events raised by the Gasless contract.
type GaslessGitCommitIterator struct {
	Event *GaslessGitCommit // Event containing the contract specifics and raw log

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
func (it *GaslessGitCommitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GaslessGitCommit)
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
		it.Event = new(GaslessGitCommit)
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
func (it *GaslessGitCommitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GaslessGitCommitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GaslessGitCommit represents a GitCommit event raised by the Gasless contract.
type GaslessGitCommit struct {
	Arg0 [20]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterGitCommit is a free log retrieval operation binding the contract event 0x16fbd3a71aa6d159973eb9ff1e1199f9fe242767e6f30ac662a492f92ac70411.
//
// Solidity: event GitCommit(bytes20 indexed arg0)
func (_Gasless *GaslessFilterer) FilterGitCommit(opts *bind.FilterOpts, arg0 [][20]byte) (*GaslessGitCommitIterator, error) {

	var arg0Rule []interface{}
	for _, arg0Item := range arg0 {
		arg0Rule = append(arg0Rule, arg0Item)
	}

	logs, sub, err := _Gasless.contract.FilterLogs(opts, "GitCommit", arg0Rule)
	if err != nil {
		return nil, err
	}
	return &GaslessGitCommitIterator{contract: _Gasless.contract, event: "GitCommit", logs: logs, sub: sub}, nil
}

// WatchGitCommit is a free log subscription operation binding the contract event 0x16fbd3a71aa6d159973eb9ff1e1199f9fe242767e6f30ac662a492f92ac70411.
//
// Solidity: event GitCommit(bytes20 indexed arg0)
func (_Gasless *GaslessFilterer) WatchGitCommit(opts *bind.WatchOpts, sink chan<- *GaslessGitCommit, arg0 [][20]byte) (event.Subscription, error) {

	var arg0Rule []interface{}
	for _, arg0Item := range arg0 {
		arg0Rule = append(arg0Rule, arg0Item)
	}

	logs, sub, err := _Gasless.contract.WatchLogs(opts, "GitCommit", arg0Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GaslessGitCommit)
				if err := _Gasless.contract.UnpackLog(event, "GitCommit", log); err != nil {
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

// ParseGitCommit is a log parse operation binding the contract event 0x16fbd3a71aa6d159973eb9ff1e1199f9fe242767e6f30ac662a492f92ac70411.
//
// Solidity: event GitCommit(bytes20 indexed arg0)
func (_Gasless *GaslessFilterer) ParseGitCommit(log types.Log) (*GaslessGitCommit, error) {
	event := new(GaslessGitCommit)
	if err := _Gasless.contract.UnpackLog(event, "GitCommit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
