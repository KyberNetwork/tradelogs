// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package newgasless

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

// NewGaslessMetaData contains all meta data concerning the NewGasless contract.
var NewGaslessMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes20\",\"name\":\"gitCommit\",\"type\":\"bytes20\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"},{\"internalType\":\"bytes4\",\"name\":\"action\",\"type\":\"bytes4\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"ActionInvalid\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"sellToken\",\"type\":\"address\"}],\"name\":\"BoughtSellToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"callbackInt\",\"type\":\"uint256\"}],\"name\":\"CallbackNotSpent\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ConfusedDeputy\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"DeltaNotPositive\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForwarderNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidOffset\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"prev\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"solver\",\"type\":\"address\"}],\"name\":\"InvalidSolver\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTarget\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotConverged\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PayerSpent\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PermissionDenied\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"callbackInt\",\"type\":\"uint256\"}],\"name\":\"ReentrantCallback\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"oldWitness\",\"type\":\"bytes32\"}],\"name\":\"ReentrantMetatransaction\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"oldPayer\",\"type\":\"address\"}],\"name\":\"ReentrantPayer\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"TooMuchSlippage\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"forkId\",\"type\":\"uint8\"}],\"name\":\"UnknownForkId\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"oldWitness\",\"type\":\"bytes32\"}],\"name\":\"WitnessNotSpent\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"ZeroSellAmount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes20\",\"name\":\"\",\"type\":\"bytes20\"}],\"name\":\"GitCommit\",\"type\":\"event\"},{\"stateMutability\":\"nonpayable\",\"type\":\"fallback\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"buyToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minAmountOut\",\"type\":\"uint256\"}],\"internalType\":\"structSettlerBase.AllowedSlippage\",\"name\":\"slippage\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"actions\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"msgSender\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"executeMetaTxn\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSolvers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner_\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rebateClaimer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"prev\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"solver\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"addNotRemove\",\"type\":\"bool\"}],\"name\":\"setSolver\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// NewGaslessABI is the input ABI used to generate the binding from.
// Deprecated: Use NewGaslessMetaData.ABI instead.
var NewGaslessABI = NewGaslessMetaData.ABI

// NewGasless is an auto generated Go binding around an Ethereum contract.
type NewGasless struct {
	NewGaslessCaller     // Read-only binding to the contract
	NewGaslessTransactor // Write-only binding to the contract
	NewGaslessFilterer   // Log filterer for contract events
}

// NewGaslessCaller is an auto generated read-only Go binding around an Ethereum contract.
type NewGaslessCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NewGaslessTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NewGaslessTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NewGaslessFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NewGaslessFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NewGaslessSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NewGaslessSession struct {
	Contract     *NewGasless       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NewGaslessCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NewGaslessCallerSession struct {
	Contract *NewGaslessCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// NewGaslessTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NewGaslessTransactorSession struct {
	Contract     *NewGaslessTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// NewGaslessRaw is an auto generated low-level Go binding around an Ethereum contract.
type NewGaslessRaw struct {
	Contract *NewGasless // Generic contract binding to access the raw methods on
}

// NewGaslessCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NewGaslessCallerRaw struct {
	Contract *NewGaslessCaller // Generic read-only contract binding to access the raw methods on
}

// NewGaslessTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NewGaslessTransactorRaw struct {
	Contract *NewGaslessTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNewGasless creates a new instance of NewGasless, bound to a specific deployed contract.
func NewNewGasless(address common.Address, backend bind.ContractBackend) (*NewGasless, error) {
	contract, err := bindNewGasless(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NewGasless{NewGaslessCaller: NewGaslessCaller{contract: contract}, NewGaslessTransactor: NewGaslessTransactor{contract: contract}, NewGaslessFilterer: NewGaslessFilterer{contract: contract}}, nil
}

// NewNewGaslessCaller creates a new read-only instance of NewGasless, bound to a specific deployed contract.
func NewNewGaslessCaller(address common.Address, caller bind.ContractCaller) (*NewGaslessCaller, error) {
	contract, err := bindNewGasless(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NewGaslessCaller{contract: contract}, nil
}

// NewNewGaslessTransactor creates a new write-only instance of NewGasless, bound to a specific deployed contract.
func NewNewGaslessTransactor(address common.Address, transactor bind.ContractTransactor) (*NewGaslessTransactor, error) {
	contract, err := bindNewGasless(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NewGaslessTransactor{contract: contract}, nil
}

// NewNewGaslessFilterer creates a new log filterer instance of NewGasless, bound to a specific deployed contract.
func NewNewGaslessFilterer(address common.Address, filterer bind.ContractFilterer) (*NewGaslessFilterer, error) {
	contract, err := bindNewGasless(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NewGaslessFilterer{contract: contract}, nil
}

// bindNewGasless binds a generic wrapper to an already deployed contract.
func bindNewGasless(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := NewGaslessMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NewGasless *NewGaslessRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NewGasless.Contract.NewGaslessCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NewGasless *NewGaslessRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NewGasless.Contract.NewGaslessTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NewGasless *NewGaslessRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NewGasless.Contract.NewGaslessTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NewGasless *NewGaslessCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NewGasless.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NewGasless *NewGaslessTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NewGasless.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NewGasless *NewGaslessTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NewGasless.Contract.contract.Transact(opts, method, params...)
}

// GetSolvers is a free data retrieval call binding the contract method 0x8bc1e8eb.
//
// Solidity: function getSolvers() view returns(address[])
func (_NewGasless *NewGaslessCaller) GetSolvers(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _NewGasless.contract.Call(opts, &out, "getSolvers")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetSolvers is a free data retrieval call binding the contract method 0x8bc1e8eb.
//
// Solidity: function getSolvers() view returns(address[])
func (_NewGasless *NewGaslessSession) GetSolvers() ([]common.Address, error) {
	return _NewGasless.Contract.GetSolvers(&_NewGasless.CallOpts)
}

// GetSolvers is a free data retrieval call binding the contract method 0x8bc1e8eb.
//
// Solidity: function getSolvers() view returns(address[])
func (_NewGasless *NewGaslessCallerSession) GetSolvers() ([]common.Address, error) {
	return _NewGasless.Contract.GetSolvers(&_NewGasless.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address owner_)
func (_NewGasless *NewGaslessCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NewGasless.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address owner_)
func (_NewGasless *NewGaslessSession) Owner() (common.Address, error) {
	return _NewGasless.Contract.Owner(&_NewGasless.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address owner_)
func (_NewGasless *NewGaslessCallerSession) Owner() (common.Address, error) {
	return _NewGasless.Contract.Owner(&_NewGasless.CallOpts)
}

// RebateClaimer is a free data retrieval call binding the contract method 0x67c4a3b0.
//
// Solidity: function rebateClaimer() view returns(address)
func (_NewGasless *NewGaslessCaller) RebateClaimer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NewGasless.contract.Call(opts, &out, "rebateClaimer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RebateClaimer is a free data retrieval call binding the contract method 0x67c4a3b0.
//
// Solidity: function rebateClaimer() view returns(address)
func (_NewGasless *NewGaslessSession) RebateClaimer() (common.Address, error) {
	return _NewGasless.Contract.RebateClaimer(&_NewGasless.CallOpts)
}

// RebateClaimer is a free data retrieval call binding the contract method 0x67c4a3b0.
//
// Solidity: function rebateClaimer() view returns(address)
func (_NewGasless *NewGaslessCallerSession) RebateClaimer() (common.Address, error) {
	return _NewGasless.Contract.RebateClaimer(&_NewGasless.CallOpts)
}

// ExecuteMetaTxn is a paid mutator transaction binding the contract method 0xfd3ad6d4.
//
// Solidity: function executeMetaTxn((address,address,uint256) slippage, bytes[] actions, bytes32 , address msgSender, bytes sig) returns(bool)
func (_NewGasless *NewGaslessTransactor) ExecuteMetaTxn(opts *bind.TransactOpts, slippage SettlerBaseAllowedSlippage, actions [][]byte, arg2 [32]byte, msgSender common.Address, sig []byte) (*types.Transaction, error) {
	return _NewGasless.contract.Transact(opts, "executeMetaTxn", slippage, actions, arg2, msgSender, sig)
}

// ExecuteMetaTxn is a paid mutator transaction binding the contract method 0xfd3ad6d4.
//
// Solidity: function executeMetaTxn((address,address,uint256) slippage, bytes[] actions, bytes32 , address msgSender, bytes sig) returns(bool)
func (_NewGasless *NewGaslessSession) ExecuteMetaTxn(slippage SettlerBaseAllowedSlippage, actions [][]byte, arg2 [32]byte, msgSender common.Address, sig []byte) (*types.Transaction, error) {
	return _NewGasless.Contract.ExecuteMetaTxn(&_NewGasless.TransactOpts, slippage, actions, arg2, msgSender, sig)
}

// ExecuteMetaTxn is a paid mutator transaction binding the contract method 0xfd3ad6d4.
//
// Solidity: function executeMetaTxn((address,address,uint256) slippage, bytes[] actions, bytes32 , address msgSender, bytes sig) returns(bool)
func (_NewGasless *NewGaslessTransactorSession) ExecuteMetaTxn(slippage SettlerBaseAllowedSlippage, actions [][]byte, arg2 [32]byte, msgSender common.Address, sig []byte) (*types.Transaction, error) {
	return _NewGasless.Contract.ExecuteMetaTxn(&_NewGasless.TransactOpts, slippage, actions, arg2, msgSender, sig)
}

// SetSolver is a paid mutator transaction binding the contract method 0x4b7758a5.
//
// Solidity: function setSolver(address prev, address solver, bool addNotRemove) returns()
func (_NewGasless *NewGaslessTransactor) SetSolver(opts *bind.TransactOpts, prev common.Address, solver common.Address, addNotRemove bool) (*types.Transaction, error) {
	return _NewGasless.contract.Transact(opts, "setSolver", prev, solver, addNotRemove)
}

// SetSolver is a paid mutator transaction binding the contract method 0x4b7758a5.
//
// Solidity: function setSolver(address prev, address solver, bool addNotRemove) returns()
func (_NewGasless *NewGaslessSession) SetSolver(prev common.Address, solver common.Address, addNotRemove bool) (*types.Transaction, error) {
	return _NewGasless.Contract.SetSolver(&_NewGasless.TransactOpts, prev, solver, addNotRemove)
}

// SetSolver is a paid mutator transaction binding the contract method 0x4b7758a5.
//
// Solidity: function setSolver(address prev, address solver, bool addNotRemove) returns()
func (_NewGasless *NewGaslessTransactorSession) SetSolver(prev common.Address, solver common.Address, addNotRemove bool) (*types.Transaction, error) {
	return _NewGasless.Contract.SetSolver(&_NewGasless.TransactOpts, prev, solver, addNotRemove)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_NewGasless *NewGaslessTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _NewGasless.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_NewGasless *NewGaslessSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _NewGasless.Contract.Fallback(&_NewGasless.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_NewGasless *NewGaslessTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _NewGasless.Contract.Fallback(&_NewGasless.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_NewGasless *NewGaslessTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NewGasless.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_NewGasless *NewGaslessSession) Receive() (*types.Transaction, error) {
	return _NewGasless.Contract.Receive(&_NewGasless.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_NewGasless *NewGaslessTransactorSession) Receive() (*types.Transaction, error) {
	return _NewGasless.Contract.Receive(&_NewGasless.TransactOpts)
}

// NewGaslessGitCommitIterator is returned from FilterGitCommit and is used to iterate over the raw logs and unpacked data for GitCommit events raised by the NewGasless contract.
type NewGaslessGitCommitIterator struct {
	Event *NewGaslessGitCommit // Event containing the contract specifics and raw log

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
func (it *NewGaslessGitCommitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NewGaslessGitCommit)
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
		it.Event = new(NewGaslessGitCommit)
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
func (it *NewGaslessGitCommitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NewGaslessGitCommitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NewGaslessGitCommit represents a GitCommit event raised by the NewGasless contract.
type NewGaslessGitCommit struct {
	Arg0 [20]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterGitCommit is a free log retrieval operation binding the contract event 0x16fbd3a71aa6d159973eb9ff1e1199f9fe242767e6f30ac662a492f92ac70411.
//
// Solidity: event GitCommit(bytes20 indexed arg0)
func (_NewGasless *NewGaslessFilterer) FilterGitCommit(opts *bind.FilterOpts, arg0 [][20]byte) (*NewGaslessGitCommitIterator, error) {

	var arg0Rule []interface{}
	for _, arg0Item := range arg0 {
		arg0Rule = append(arg0Rule, arg0Item)
	}

	logs, sub, err := _NewGasless.contract.FilterLogs(opts, "GitCommit", arg0Rule)
	if err != nil {
		return nil, err
	}
	return &NewGaslessGitCommitIterator{contract: _NewGasless.contract, event: "GitCommit", logs: logs, sub: sub}, nil
}

// WatchGitCommit is a free log subscription operation binding the contract event 0x16fbd3a71aa6d159973eb9ff1e1199f9fe242767e6f30ac662a492f92ac70411.
//
// Solidity: event GitCommit(bytes20 indexed arg0)
func (_NewGasless *NewGaslessFilterer) WatchGitCommit(opts *bind.WatchOpts, sink chan<- *NewGaslessGitCommit, arg0 [][20]byte) (event.Subscription, error) {

	var arg0Rule []interface{}
	for _, arg0Item := range arg0 {
		arg0Rule = append(arg0Rule, arg0Item)
	}

	logs, sub, err := _NewGasless.contract.WatchLogs(opts, "GitCommit", arg0Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NewGaslessGitCommit)
				if err := _NewGasless.contract.UnpackLog(event, "GitCommit", log); err != nil {
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
func (_NewGasless *NewGaslessFilterer) ParseGitCommit(log types.Log) (*NewGaslessGitCommit, error) {
	event := new(NewGaslessGitCommit)
	if err := _NewGasless.contract.UnpackLog(event, "GitCommit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
