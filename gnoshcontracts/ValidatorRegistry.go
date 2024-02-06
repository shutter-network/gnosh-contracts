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

// IValidatorRegistryUpdate is an auto generated low-level Go binding around an user-defined struct.
type IValidatorRegistryUpdate struct {
	Message   []byte
	Signature []byte
}

// GnoshcontractsMetaData contains all meta data concerning the Gnoshcontracts contract.
var GnoshcontractsMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"getNumUpdates\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getUpdate\",\"inputs\":[{\"name\":\"i\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIValidatorRegistry.Update\",\"components\":[{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"update\",\"inputs\":[{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Updated\",\"inputs\":[{\"name\":\"message\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"signature\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false}]",
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

// GetNumUpdates is a free data retrieval call binding the contract method 0x00c7c019.
//
// Solidity: function getNumUpdates() view returns(uint256)
func (_Gnoshcontracts *GnoshcontractsCaller) GetNumUpdates(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Gnoshcontracts.contract.Call(opts, &out, "getNumUpdates")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNumUpdates is a free data retrieval call binding the contract method 0x00c7c019.
//
// Solidity: function getNumUpdates() view returns(uint256)
func (_Gnoshcontracts *GnoshcontractsSession) GetNumUpdates() (*big.Int, error) {
	return _Gnoshcontracts.Contract.GetNumUpdates(&_Gnoshcontracts.CallOpts)
}

// GetNumUpdates is a free data retrieval call binding the contract method 0x00c7c019.
//
// Solidity: function getNumUpdates() view returns(uint256)
func (_Gnoshcontracts *GnoshcontractsCallerSession) GetNumUpdates() (*big.Int, error) {
	return _Gnoshcontracts.Contract.GetNumUpdates(&_Gnoshcontracts.CallOpts)
}

// GetUpdate is a free data retrieval call binding the contract method 0x32cb25be.
//
// Solidity: function getUpdate(uint256 i) view returns((bytes,bytes))
func (_Gnoshcontracts *GnoshcontractsCaller) GetUpdate(opts *bind.CallOpts, i *big.Int) (IValidatorRegistryUpdate, error) {
	var out []interface{}
	err := _Gnoshcontracts.contract.Call(opts, &out, "getUpdate", i)

	if err != nil {
		return *new(IValidatorRegistryUpdate), err
	}

	out0 := *abi.ConvertType(out[0], new(IValidatorRegistryUpdate)).(*IValidatorRegistryUpdate)

	return out0, err

}

// GetUpdate is a free data retrieval call binding the contract method 0x32cb25be.
//
// Solidity: function getUpdate(uint256 i) view returns((bytes,bytes))
func (_Gnoshcontracts *GnoshcontractsSession) GetUpdate(i *big.Int) (IValidatorRegistryUpdate, error) {
	return _Gnoshcontracts.Contract.GetUpdate(&_Gnoshcontracts.CallOpts, i)
}

// GetUpdate is a free data retrieval call binding the contract method 0x32cb25be.
//
// Solidity: function getUpdate(uint256 i) view returns((bytes,bytes))
func (_Gnoshcontracts *GnoshcontractsCallerSession) GetUpdate(i *big.Int) (IValidatorRegistryUpdate, error) {
	return _Gnoshcontracts.Contract.GetUpdate(&_Gnoshcontracts.CallOpts, i)
}

// Update is a paid mutator transaction binding the contract method 0x3f37dce2.
//
// Solidity: function update(bytes message, bytes signature) returns()
func (_Gnoshcontracts *GnoshcontractsTransactor) Update(opts *bind.TransactOpts, message []byte, signature []byte) (*types.Transaction, error) {
	return _Gnoshcontracts.contract.Transact(opts, "update", message, signature)
}

// Update is a paid mutator transaction binding the contract method 0x3f37dce2.
//
// Solidity: function update(bytes message, bytes signature) returns()
func (_Gnoshcontracts *GnoshcontractsSession) Update(message []byte, signature []byte) (*types.Transaction, error) {
	return _Gnoshcontracts.Contract.Update(&_Gnoshcontracts.TransactOpts, message, signature)
}

// Update is a paid mutator transaction binding the contract method 0x3f37dce2.
//
// Solidity: function update(bytes message, bytes signature) returns()
func (_Gnoshcontracts *GnoshcontractsTransactorSession) Update(message []byte, signature []byte) (*types.Transaction, error) {
	return _Gnoshcontracts.Contract.Update(&_Gnoshcontracts.TransactOpts, message, signature)
}

// GnoshcontractsUpdatedIterator is returned from FilterUpdated and is used to iterate over the raw logs and unpacked data for Updated events raised by the Gnoshcontracts contract.
type GnoshcontractsUpdatedIterator struct {
	Event *GnoshcontractsUpdated // Event containing the contract specifics and raw log

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
func (it *GnoshcontractsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GnoshcontractsUpdated)
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
		it.Event = new(GnoshcontractsUpdated)
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
func (it *GnoshcontractsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GnoshcontractsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GnoshcontractsUpdated represents a Updated event raised by the Gnoshcontracts contract.
type GnoshcontractsUpdated struct {
	Message   []byte
	Signature []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUpdated is a free log retrieval operation binding the contract event 0x9796f15c93411b364b7f09bb591d0f77aa4dc399cf0481b8de1c3ce9f10a3a87.
//
// Solidity: event Updated(bytes message, bytes signature)
func (_Gnoshcontracts *GnoshcontractsFilterer) FilterUpdated(opts *bind.FilterOpts) (*GnoshcontractsUpdatedIterator, error) {

	logs, sub, err := _Gnoshcontracts.contract.FilterLogs(opts, "Updated")
	if err != nil {
		return nil, err
	}
	return &GnoshcontractsUpdatedIterator{contract: _Gnoshcontracts.contract, event: "Updated", logs: logs, sub: sub}, nil
}

// WatchUpdated is a free log subscription operation binding the contract event 0x9796f15c93411b364b7f09bb591d0f77aa4dc399cf0481b8de1c3ce9f10a3a87.
//
// Solidity: event Updated(bytes message, bytes signature)
func (_Gnoshcontracts *GnoshcontractsFilterer) WatchUpdated(opts *bind.WatchOpts, sink chan<- *GnoshcontractsUpdated) (event.Subscription, error) {

	logs, sub, err := _Gnoshcontracts.contract.WatchLogs(opts, "Updated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GnoshcontractsUpdated)
				if err := _Gnoshcontracts.contract.UnpackLog(event, "Updated", log); err != nil {
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

// ParseUpdated is a log parse operation binding the contract event 0x9796f15c93411b364b7f09bb591d0f77aa4dc399cf0481b8de1c3ce9f10a3a87.
//
// Solidity: event Updated(bytes message, bytes signature)
func (_Gnoshcontracts *GnoshcontractsFilterer) ParseUpdated(log types.Log) (*GnoshcontractsUpdated, error) {
	event := new(GnoshcontractsUpdated)
	if err := _Gnoshcontracts.contract.UnpackLog(event, "Updated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
