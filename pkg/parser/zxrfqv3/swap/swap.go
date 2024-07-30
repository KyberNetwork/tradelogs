// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package swap

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

// SwapMetaData contains all meta data concerning the Swap contract.
var SwapMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes20\",\"name\":\"gitCommit\",\"type\":\"bytes20\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"},{\"internalType\":\"bytes4\",\"name\":\"action\",\"type\":\"bytes4\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"ActionInvalid\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"callbackInt\",\"type\":\"uint256\"}],\"name\":\"CallbackNotSpent\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ConfusedDeputy\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForwarderNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidOffset\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSignatureLen\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTarget\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotConverged\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PayerSpent\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"callbackInt\",\"type\":\"uint256\"}],\"name\":\"ReentrantCallback\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"oldPayer\",\"type\":\"address\"}],\"name\":\"ReentrantPayer\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"SignatureExpired\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"TooMuchSlippage\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"forkId\",\"type\":\"uint8\"}],\"name\":\"UnknownForkId\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes20\",\"name\":\"\",\"type\":\"bytes20\"}],\"name\":\"GitCommit\",\"type\":\"event\"},{\"stateMutability\":\"nonpayable\",\"type\":\"fallback\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"buyToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minAmountOut\",\"type\":\"uint256\"}],\"internalType\":\"structSettlerBase.AllowedSlippage\",\"name\":\"slippage\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"actions\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"execute\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// SwapABI is the input ABI used to generate the binding from.
// Deprecated: Use SwapMetaData.ABI instead.
var SwapABI = SwapMetaData.ABI

// Swap is an auto generated Go binding around an Ethereum contract.
type Swap struct {
	SwapCaller     // Read-only binding to the contract
	SwapTransactor // Write-only binding to the contract
	SwapFilterer   // Log filterer for contract events
}

// SwapCaller is an auto generated read-only Go binding around an Ethereum contract.
type SwapCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SwapTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SwapFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SwapSession struct {
	Contract     *Swap             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SwapCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SwapCallerSession struct {
	Contract *SwapCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// SwapTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SwapTransactorSession struct {
	Contract     *SwapTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SwapRaw is an auto generated low-level Go binding around an Ethereum contract.
type SwapRaw struct {
	Contract *Swap // Generic contract binding to access the raw methods on
}

// SwapCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SwapCallerRaw struct {
	Contract *SwapCaller // Generic read-only contract binding to access the raw methods on
}

// SwapTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SwapTransactorRaw struct {
	Contract *SwapTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSwap creates a new instance of Swap, bound to a specific deployed contract.
func NewSwap(address common.Address, backend bind.ContractBackend) (*Swap, error) {
	contract, err := bindSwap(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Swap{SwapCaller: SwapCaller{contract: contract}, SwapTransactor: SwapTransactor{contract: contract}, SwapFilterer: SwapFilterer{contract: contract}}, nil
}

// NewSwapCaller creates a new read-only instance of Swap, bound to a specific deployed contract.
func NewSwapCaller(address common.Address, caller bind.ContractCaller) (*SwapCaller, error) {
	contract, err := bindSwap(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SwapCaller{contract: contract}, nil
}

// NewSwapTransactor creates a new write-only instance of Swap, bound to a specific deployed contract.
func NewSwapTransactor(address common.Address, transactor bind.ContractTransactor) (*SwapTransactor, error) {
	contract, err := bindSwap(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SwapTransactor{contract: contract}, nil
}

// NewSwapFilterer creates a new log filterer instance of Swap, bound to a specific deployed contract.
func NewSwapFilterer(address common.Address, filterer bind.ContractFilterer) (*SwapFilterer, error) {
	contract, err := bindSwap(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SwapFilterer{contract: contract}, nil
}

// bindSwap binds a generic wrapper to an already deployed contract.
func bindSwap(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SwapMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Swap *SwapRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Swap.Contract.SwapCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Swap *SwapRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Swap.Contract.SwapTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Swap *SwapRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Swap.Contract.SwapTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Swap *SwapCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Swap.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Swap *SwapTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Swap.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Swap *SwapTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Swap.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) pure returns()
func (_Swap *SwapCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) error {
	var out []interface{}
	err := _Swap.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return err
	}

	return err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) pure returns()
func (_Swap *SwapSession) BalanceOf(arg0 common.Address) error {
	return _Swap.Contract.BalanceOf(&_Swap.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) pure returns()
func (_Swap *SwapCallerSession) BalanceOf(arg0 common.Address) error {
	return _Swap.Contract.BalanceOf(&_Swap.CallOpts, arg0)
}

// Execute is a paid mutator transaction binding the contract method 0x1fff991f.
//
// Solidity: function execute((address,address,uint256) slippage, bytes[] actions, bytes32 ) payable returns(bool)
func (_Swap *SwapTransactor) Execute(opts *bind.TransactOpts, slippage SettlerBaseAllowedSlippage, actions [][]byte, arg2 [32]byte) (*types.Transaction, error) {
	return _Swap.contract.Transact(opts, "execute", slippage, actions, arg2)
}

// Execute is a paid mutator transaction binding the contract method 0x1fff991f.
//
// Solidity: function execute((address,address,uint256) slippage, bytes[] actions, bytes32 ) payable returns(bool)
func (_Swap *SwapSession) Execute(slippage SettlerBaseAllowedSlippage, actions [][]byte, arg2 [32]byte) (*types.Transaction, error) {
	return _Swap.Contract.Execute(&_Swap.TransactOpts, slippage, actions, arg2)
}

// Execute is a paid mutator transaction binding the contract method 0x1fff991f.
//
// Solidity: function execute((address,address,uint256) slippage, bytes[] actions, bytes32 ) payable returns(bool)
func (_Swap *SwapTransactorSession) Execute(slippage SettlerBaseAllowedSlippage, actions [][]byte, arg2 [32]byte) (*types.Transaction, error) {
	return _Swap.Contract.Execute(&_Swap.TransactOpts, slippage, actions, arg2)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_Swap *SwapTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Swap.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_Swap *SwapSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Swap.Contract.Fallback(&_Swap.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_Swap *SwapTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Swap.Contract.Fallback(&_Swap.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Swap *SwapTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Swap.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Swap *SwapSession) Receive() (*types.Transaction, error) {
	return _Swap.Contract.Receive(&_Swap.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Swap *SwapTransactorSession) Receive() (*types.Transaction, error) {
	return _Swap.Contract.Receive(&_Swap.TransactOpts)
}

// SwapGitCommitIterator is returned from FilterGitCommit and is used to iterate over the raw logs and unpacked data for GitCommit events raised by the Swap contract.
type SwapGitCommitIterator struct {
	Event *SwapGitCommit // Event containing the contract specifics and raw log

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
func (it *SwapGitCommitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapGitCommit)
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
		it.Event = new(SwapGitCommit)
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
func (it *SwapGitCommitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapGitCommitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapGitCommit represents a GitCommit event raised by the Swap contract.
type SwapGitCommit struct {
	Arg0 [20]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterGitCommit is a free log retrieval operation binding the contract event 0x16fbd3a71aa6d159973eb9ff1e1199f9fe242767e6f30ac662a492f92ac70411.
//
// Solidity: event GitCommit(bytes20 indexed arg0)
func (_Swap *SwapFilterer) FilterGitCommit(opts *bind.FilterOpts, arg0 [][20]byte) (*SwapGitCommitIterator, error) {

	var arg0Rule []interface{}
	for _, arg0Item := range arg0 {
		arg0Rule = append(arg0Rule, arg0Item)
	}

	logs, sub, err := _Swap.contract.FilterLogs(opts, "GitCommit", arg0Rule)
	if err != nil {
		return nil, err
	}
	return &SwapGitCommitIterator{contract: _Swap.contract, event: "GitCommit", logs: logs, sub: sub}, nil
}

// WatchGitCommit is a free log subscription operation binding the contract event 0x16fbd3a71aa6d159973eb9ff1e1199f9fe242767e6f30ac662a492f92ac70411.
//
// Solidity: event GitCommit(bytes20 indexed arg0)
func (_Swap *SwapFilterer) WatchGitCommit(opts *bind.WatchOpts, sink chan<- *SwapGitCommit, arg0 [][20]byte) (event.Subscription, error) {

	var arg0Rule []interface{}
	for _, arg0Item := range arg0 {
		arg0Rule = append(arg0Rule, arg0Item)
	}

	logs, sub, err := _Swap.contract.WatchLogs(opts, "GitCommit", arg0Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapGitCommit)
				if err := _Swap.contract.UnpackLog(event, "GitCommit", log); err != nil {
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
func (_Swap *SwapFilterer) ParseGitCommit(log types.Log) (*SwapGitCommit, error) {
	event := new(SwapGitCommit)
	if err := _Swap.contract.UnpackLog(event, "GitCommit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
