// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package gnoshcontracts

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

// GnoshcontractsMetaData contains all meta data concerning the Gnoshcontracts contract.
var GnoshcontractsMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"submitDecryptionProgress\",\"inputs\":[{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"submitEncryptedTransaction\",\"inputs\":[{\"name\":\"eon\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"identityPrefix\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"encryptedTransaction\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"event\",\"name\":\"DecryptionProgressSubmitted\",\"inputs\":[{\"name\":\"message\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TransactionSubmitted\",\"inputs\":[{\"name\":\"eon\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"identityPrefix\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"encryptedTransaction\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"gasLimit\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"InsufficientFee\",\"inputs\":[]}]",
}

// GnoshcontractsABI is the input ABI used to generate the binding from.
// Deprecated: Use GnoshcontractsMetaData.ABI instead.
var GnoshcontractsABI = GnoshcontractsMetaData.ABI

// Gnoshcontracts is an auto generated Go binding around an Ethereum contract.
type Gnoshcontracts struct {
	GnoshcontractsCaller     // Read-only binding to the contract
	GnoshcontractsTransactor // Write-only binding to the contract
	GnoshcontractsFilterer   // Log filterer for contract events
}

// GnoshcontractsCaller is an auto generated read-only Go binding around an Ethereum contract.
type GnoshcontractsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GnoshcontractsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GnoshcontractsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GnoshcontractsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GnoshcontractsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GnoshcontractsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GnoshcontractsSession struct {
	Contract     *Gnoshcontracts   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GnoshcontractsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GnoshcontractsCallerSession struct {
	Contract *GnoshcontractsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// GnoshcontractsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GnoshcontractsTransactorSession struct {
	Contract     *GnoshcontractsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// GnoshcontractsRaw is an auto generated low-level Go binding around an Ethereum contract.
type GnoshcontractsRaw struct {
	Contract *Gnoshcontracts // Generic contract binding to access the raw methods on
}

// GnoshcontractsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GnoshcontractsCallerRaw struct {
	Contract *GnoshcontractsCaller // Generic read-only contract binding to access the raw methods on
}

// GnoshcontractsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GnoshcontractsTransactorRaw struct {
	Contract *GnoshcontractsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGnoshcontracts creates a new instance of Gnoshcontracts, bound to a specific deployed contract.
func NewGnoshcontracts(address common.Address, backend bind.ContractBackend) (*Gnoshcontracts, error) {
	contract, err := bindGnoshcontracts(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Gnoshcontracts{GnoshcontractsCaller: GnoshcontractsCaller{contract: contract}, GnoshcontractsTransactor: GnoshcontractsTransactor{contract: contract}, GnoshcontractsFilterer: GnoshcontractsFilterer{contract: contract}}, nil
}

// NewGnoshcontractsCaller creates a new read-only instance of Gnoshcontracts, bound to a specific deployed contract.
func NewGnoshcontractsCaller(address common.Address, caller bind.ContractCaller) (*GnoshcontractsCaller, error) {
	contract, err := bindGnoshcontracts(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GnoshcontractsCaller{contract: contract}, nil
}

// NewGnoshcontractsTransactor creates a new write-only instance of Gnoshcontracts, bound to a specific deployed contract.
func NewGnoshcontractsTransactor(address common.Address, transactor bind.ContractTransactor) (*GnoshcontractsTransactor, error) {
	contract, err := bindGnoshcontracts(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GnoshcontractsTransactor{contract: contract}, nil
}

// NewGnoshcontractsFilterer creates a new log filterer instance of Gnoshcontracts, bound to a specific deployed contract.
func NewGnoshcontractsFilterer(address common.Address, filterer bind.ContractFilterer) (*GnoshcontractsFilterer, error) {
	contract, err := bindGnoshcontracts(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GnoshcontractsFilterer{contract: contract}, nil
}

// bindGnoshcontracts binds a generic wrapper to an already deployed contract.
func bindGnoshcontracts(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := GnoshcontractsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Gnoshcontracts *GnoshcontractsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Gnoshcontracts.Contract.GnoshcontractsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Gnoshcontracts *GnoshcontractsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Gnoshcontracts.Contract.GnoshcontractsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Gnoshcontracts *GnoshcontractsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Gnoshcontracts.Contract.GnoshcontractsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Gnoshcontracts *GnoshcontractsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Gnoshcontracts.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Gnoshcontracts *GnoshcontractsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Gnoshcontracts.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Gnoshcontracts *GnoshcontractsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Gnoshcontracts.Contract.contract.Transact(opts, method, params...)
}

// SubmitDecryptionProgress is a paid mutator transaction binding the contract method 0x2d32522e.
//
// Solidity: function submitDecryptionProgress(bytes message) returns()
func (_Gnoshcontracts *GnoshcontractsTransactor) SubmitDecryptionProgress(opts *bind.TransactOpts, message []byte) (*types.Transaction, error) {
	return _Gnoshcontracts.contract.Transact(opts, "submitDecryptionProgress", message)
}

// SubmitDecryptionProgress is a paid mutator transaction binding the contract method 0x2d32522e.
//
// Solidity: function submitDecryptionProgress(bytes message) returns()
func (_Gnoshcontracts *GnoshcontractsSession) SubmitDecryptionProgress(message []byte) (*types.Transaction, error) {
	return _Gnoshcontracts.Contract.SubmitDecryptionProgress(&_Gnoshcontracts.TransactOpts, message)
}

// SubmitDecryptionProgress is a paid mutator transaction binding the contract method 0x2d32522e.
//
// Solidity: function submitDecryptionProgress(bytes message) returns()
func (_Gnoshcontracts *GnoshcontractsTransactorSession) SubmitDecryptionProgress(message []byte) (*types.Transaction, error) {
	return _Gnoshcontracts.Contract.SubmitDecryptionProgress(&_Gnoshcontracts.TransactOpts, message)
}

// SubmitEncryptedTransaction is a paid mutator transaction binding the contract method 0x6a69d2e1.
//
// Solidity: function submitEncryptedTransaction(uint64 eon, bytes32 identityPrefix, bytes encryptedTransaction, uint256 gasLimit) payable returns()
func (_Gnoshcontracts *GnoshcontractsTransactor) SubmitEncryptedTransaction(opts *bind.TransactOpts, eon uint64, identityPrefix [32]byte, encryptedTransaction []byte, gasLimit *big.Int) (*types.Transaction, error) {
	return _Gnoshcontracts.contract.Transact(opts, "submitEncryptedTransaction", eon, identityPrefix, encryptedTransaction, gasLimit)
}

// SubmitEncryptedTransaction is a paid mutator transaction binding the contract method 0x6a69d2e1.
//
// Solidity: function submitEncryptedTransaction(uint64 eon, bytes32 identityPrefix, bytes encryptedTransaction, uint256 gasLimit) payable returns()
func (_Gnoshcontracts *GnoshcontractsSession) SubmitEncryptedTransaction(eon uint64, identityPrefix [32]byte, encryptedTransaction []byte, gasLimit *big.Int) (*types.Transaction, error) {
	return _Gnoshcontracts.Contract.SubmitEncryptedTransaction(&_Gnoshcontracts.TransactOpts, eon, identityPrefix, encryptedTransaction, gasLimit)
}

// SubmitEncryptedTransaction is a paid mutator transaction binding the contract method 0x6a69d2e1.
//
// Solidity: function submitEncryptedTransaction(uint64 eon, bytes32 identityPrefix, bytes encryptedTransaction, uint256 gasLimit) payable returns()
func (_Gnoshcontracts *GnoshcontractsTransactorSession) SubmitEncryptedTransaction(eon uint64, identityPrefix [32]byte, encryptedTransaction []byte, gasLimit *big.Int) (*types.Transaction, error) {
	return _Gnoshcontracts.Contract.SubmitEncryptedTransaction(&_Gnoshcontracts.TransactOpts, eon, identityPrefix, encryptedTransaction, gasLimit)
}

// GnoshcontractsDecryptionProgressSubmittedIterator is returned from FilterDecryptionProgressSubmitted and is used to iterate over the raw logs and unpacked data for DecryptionProgressSubmitted events raised by the Gnoshcontracts contract.
type GnoshcontractsDecryptionProgressSubmittedIterator struct {
	Event *GnoshcontractsDecryptionProgressSubmitted // Event containing the contract specifics and raw log

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
func (it *GnoshcontractsDecryptionProgressSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GnoshcontractsDecryptionProgressSubmitted)
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
		it.Event = new(GnoshcontractsDecryptionProgressSubmitted)
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
func (it *GnoshcontractsDecryptionProgressSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GnoshcontractsDecryptionProgressSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GnoshcontractsDecryptionProgressSubmitted represents a DecryptionProgressSubmitted event raised by the Gnoshcontracts contract.
type GnoshcontractsDecryptionProgressSubmitted struct {
	Message []byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDecryptionProgressSubmitted is a free log retrieval operation binding the contract event 0xa9a0645b33a70f18b8d490681d637cb46a859ec51707787e6f46b942f90e8f59.
//
// Solidity: event DecryptionProgressSubmitted(bytes message)
func (_Gnoshcontracts *GnoshcontractsFilterer) FilterDecryptionProgressSubmitted(opts *bind.FilterOpts) (*GnoshcontractsDecryptionProgressSubmittedIterator, error) {

	logs, sub, err := _Gnoshcontracts.contract.FilterLogs(opts, "DecryptionProgressSubmitted")
	if err != nil {
		return nil, err
	}
	return &GnoshcontractsDecryptionProgressSubmittedIterator{contract: _Gnoshcontracts.contract, event: "DecryptionProgressSubmitted", logs: logs, sub: sub}, nil
}

// WatchDecryptionProgressSubmitted is a free log subscription operation binding the contract event 0xa9a0645b33a70f18b8d490681d637cb46a859ec51707787e6f46b942f90e8f59.
//
// Solidity: event DecryptionProgressSubmitted(bytes message)
func (_Gnoshcontracts *GnoshcontractsFilterer) WatchDecryptionProgressSubmitted(opts *bind.WatchOpts, sink chan<- *GnoshcontractsDecryptionProgressSubmitted) (event.Subscription, error) {

	logs, sub, err := _Gnoshcontracts.contract.WatchLogs(opts, "DecryptionProgressSubmitted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GnoshcontractsDecryptionProgressSubmitted)
				if err := _Gnoshcontracts.contract.UnpackLog(event, "DecryptionProgressSubmitted", log); err != nil {
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

// ParseDecryptionProgressSubmitted is a log parse operation binding the contract event 0xa9a0645b33a70f18b8d490681d637cb46a859ec51707787e6f46b942f90e8f59.
//
// Solidity: event DecryptionProgressSubmitted(bytes message)
func (_Gnoshcontracts *GnoshcontractsFilterer) ParseDecryptionProgressSubmitted(log types.Log) (*GnoshcontractsDecryptionProgressSubmitted, error) {
	event := new(GnoshcontractsDecryptionProgressSubmitted)
	if err := _Gnoshcontracts.contract.UnpackLog(event, "DecryptionProgressSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GnoshcontractsTransactionSubmittedIterator is returned from FilterTransactionSubmitted and is used to iterate over the raw logs and unpacked data for TransactionSubmitted events raised by the Gnoshcontracts contract.
type GnoshcontractsTransactionSubmittedIterator struct {
	Event *GnoshcontractsTransactionSubmitted // Event containing the contract specifics and raw log

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
func (it *GnoshcontractsTransactionSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GnoshcontractsTransactionSubmitted)
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
		it.Event = new(GnoshcontractsTransactionSubmitted)
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
func (it *GnoshcontractsTransactionSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GnoshcontractsTransactionSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GnoshcontractsTransactionSubmitted represents a TransactionSubmitted event raised by the Gnoshcontracts contract.
type GnoshcontractsTransactionSubmitted struct {
	Eon                  uint64
	IdentityPrefix       [32]byte
	Sender               common.Address
	EncryptedTransaction []byte
	GasLimit             *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterTransactionSubmitted is a free log retrieval operation binding the contract event 0x6515f8e10d22a184f86cfbaeb024db7afde82add43a1b1c065e8d202e43ef1a0.
//
// Solidity: event TransactionSubmitted(uint64 eon, bytes32 identityPrefix, address sender, bytes encryptedTransaction, uint256 gasLimit)
func (_Gnoshcontracts *GnoshcontractsFilterer) FilterTransactionSubmitted(opts *bind.FilterOpts) (*GnoshcontractsTransactionSubmittedIterator, error) {

	logs, sub, err := _Gnoshcontracts.contract.FilterLogs(opts, "TransactionSubmitted")
	if err != nil {
		return nil, err
	}
	return &GnoshcontractsTransactionSubmittedIterator{contract: _Gnoshcontracts.contract, event: "TransactionSubmitted", logs: logs, sub: sub}, nil
}

// WatchTransactionSubmitted is a free log subscription operation binding the contract event 0x6515f8e10d22a184f86cfbaeb024db7afde82add43a1b1c065e8d202e43ef1a0.
//
// Solidity: event TransactionSubmitted(uint64 eon, bytes32 identityPrefix, address sender, bytes encryptedTransaction, uint256 gasLimit)
func (_Gnoshcontracts *GnoshcontractsFilterer) WatchTransactionSubmitted(opts *bind.WatchOpts, sink chan<- *GnoshcontractsTransactionSubmitted) (event.Subscription, error) {

	logs, sub, err := _Gnoshcontracts.contract.WatchLogs(opts, "TransactionSubmitted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GnoshcontractsTransactionSubmitted)
				if err := _Gnoshcontracts.contract.UnpackLog(event, "TransactionSubmitted", log); err != nil {
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

// ParseTransactionSubmitted is a log parse operation binding the contract event 0x6515f8e10d22a184f86cfbaeb024db7afde82add43a1b1c065e8d202e43ef1a0.
//
// Solidity: event TransactionSubmitted(uint64 eon, bytes32 identityPrefix, address sender, bytes encryptedTransaction, uint256 gasLimit)
func (_Gnoshcontracts *GnoshcontractsFilterer) ParseTransactionSubmitted(log types.Log) (*GnoshcontractsTransactionSubmitted, error) {
	event := new(GnoshcontractsTransactionSubmitted)
	if err := _Gnoshcontracts.contract.UnpackLog(event, "TransactionSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
