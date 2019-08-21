// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package token

import (
	"math/big"
	"strings"

	"github.com/Onther-Tech/plasma-evm"
	"github.com/Onther-Tech/plasma-evm/accounts/abi"
	"github.com/Onther-Tech/plasma-evm/accounts/abi/bind"
	"github.com/Onther-Tech/plasma-evm/common"
	"github.com/Onther-Tech/plasma-evm/core/types"
	"github.com/Onther-Tech/plasma-evm/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// OwnableABI is the input ABI used to generate the binding from.
const OwnableABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// OwnableFuncSigs maps the 4-byte function signature to its string representation.
var OwnableFuncSigs = map[string]string{
	"8da5cb5b": "owner()",
	"715018a6": "renounceOwnership()",
	"f2fde38b": "transferOwnership(address)",
}

// OwnableBin is the compiled bytecode used for deploying new contracts.
var OwnableBin = "0x608060405234801561001057600080fd5b50600080546001600160a01b031916331790556101ce806100326000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c8063715018a6146100465780638da5cb5b14610050578063f2fde38b14610074575b600080fd5b61004e61009a565b005b6100586100f9565b604080516001600160a01b039092168252519081900360200190f35b61004e6004803603602081101561008a57600080fd5b50356001600160a01b0316610108565b6000546001600160a01b031633146100b157600080fd5b600080546040516001600160a01b03909116917ff8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c6482091a2600080546001600160a01b0319169055565b6000546001600160a01b031681565b6000546001600160a01b0316331461011f57600080fd5b6101288161012b565b50565b6001600160a01b03811661013e57600080fd5b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b039290921691909117905556fea265627a7a723058208980494d4d1a24286f8824dbad433f3b9f5dbb6efee260fce80d85896c0cfdfb64736f6c63430005090032"

// DeployOwnable deploys a new Ethereum contract, binding an instance of Ownable to it.
func DeployOwnable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Ownable, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnableABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OwnableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Ownable{OwnableCaller: OwnableCaller{contract: contract}, OwnableTransactor: OwnableTransactor{contract: contract}, OwnableFilterer: OwnableFilterer{contract: contract}}, nil
}

// Ownable is an auto generated Go binding around an Ethereum contract.
type Ownable struct {
	OwnableCaller     // Read-only binding to the contract
	OwnableTransactor // Write-only binding to the contract
	OwnableFilterer   // Log filterer for contract events
}

// OwnableCaller is an auto generated read-only Go binding around an Ethereum contract.
type OwnableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OwnableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OwnableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OwnableSession struct {
	Contract     *Ownable          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwnableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OwnableCallerSession struct {
	Contract *OwnableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// OwnableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OwnableTransactorSession struct {
	Contract     *OwnableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// OwnableRaw is an auto generated low-level Go binding around an Ethereum contract.
type OwnableRaw struct {
	Contract *Ownable // Generic contract binding to access the raw methods on
}

// OwnableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OwnableCallerRaw struct {
	Contract *OwnableCaller // Generic read-only contract binding to access the raw methods on
}

// OwnableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OwnableTransactorRaw struct {
	Contract *OwnableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOwnable creates a new instance of Ownable, bound to a specific deployed contract.
func NewOwnable(address common.Address, backend bind.ContractBackend) (*Ownable, error) {
	contract, err := bindOwnable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ownable{OwnableCaller: OwnableCaller{contract: contract}, OwnableTransactor: OwnableTransactor{contract: contract}, OwnableFilterer: OwnableFilterer{contract: contract}}, nil
}

// NewOwnableCaller creates a new read-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableCaller(address common.Address, caller bind.ContractCaller) (*OwnableCaller, error) {
	contract, err := bindOwnable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableCaller{contract: contract}, nil
}

// NewOwnableTransactor creates a new write-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableTransactor(address common.Address, transactor bind.ContractTransactor) (*OwnableTransactor, error) {
	contract, err := bindOwnable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableTransactor{contract: contract}, nil
}

// NewOwnableFilterer creates a new log filterer instance of Ownable, bound to a specific deployed contract.
func NewOwnableFilterer(address common.Address, filterer bind.ContractFilterer) (*OwnableFilterer, error) {
	contract, err := bindOwnable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OwnableFilterer{contract: contract}, nil
}

// bindOwnable binds a generic wrapper to an already deployed contract.
func bindOwnable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.OwnableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Ownable *OwnableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Ownable.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Ownable *OwnableSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Ownable *OwnableCallerSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ownable.Contract.RenounceOwnership(&_Ownable.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ownable.Contract.RenounceOwnership(&_Ownable.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_Ownable *OwnableTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_Ownable *OwnableSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_Ownable *OwnableTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, _newOwner)
}

// OwnableOwnershipRenouncedIterator is returned from FilterOwnershipRenounced and is used to iterate over the raw logs and unpacked data for OwnershipRenounced events raised by the Ownable contract.
type OwnableOwnershipRenouncedIterator struct {
	Event *OwnableOwnershipRenounced // Event containing the contract specifics and raw log

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
func (it *OwnableOwnershipRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnableOwnershipRenounced)
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
		it.Event = new(OwnableOwnershipRenounced)
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
func (it *OwnableOwnershipRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnableOwnershipRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnableOwnershipRenounced represents a OwnershipRenounced event raised by the Ownable contract.
type OwnableOwnershipRenounced struct {
	PreviousOwner common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipRenounced is a free log retrieval operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: event OwnershipRenounced(address indexed previousOwner)
func (_Ownable *OwnableFilterer) FilterOwnershipRenounced(opts *bind.FilterOpts, previousOwner []common.Address) (*OwnableOwnershipRenouncedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _Ownable.contract.FilterLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OwnableOwnershipRenouncedIterator{contract: _Ownable.contract, event: "OwnershipRenounced", logs: logs, sub: sub}, nil
}

// WatchOwnershipRenounced is a free log subscription operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: event OwnershipRenounced(address indexed previousOwner)
func (_Ownable *OwnableFilterer) WatchOwnershipRenounced(opts *bind.WatchOpts, sink chan<- *OwnableOwnershipRenounced, previousOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _Ownable.contract.WatchLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnableOwnershipRenounced)
				if err := _Ownable.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
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

// ParseOwnershipRenounced is a log parse operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: event OwnershipRenounced(address indexed previousOwner)
func (_Ownable *OwnableFilterer) ParseOwnershipRenounced(log types.Log) (*OwnableOwnershipRenounced, error) {
	event := new(OwnableOwnershipRenounced)
	if err := _Ownable.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
		return nil, err
	}
	return event, nil
}

// OwnableOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Ownable contract.
type OwnableOwnershipTransferredIterator struct {
	Event *OwnableOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *OwnableOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnableOwnershipTransferred)
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
		it.Event = new(OwnableOwnershipTransferred)
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
func (it *OwnableOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnableOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnableOwnershipTransferred represents a OwnershipTransferred event raised by the Ownable contract.
type OwnableOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable *OwnableFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OwnableOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OwnableOwnershipTransferredIterator{contract: _Ownable.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable *OwnableFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OwnableOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnableOwnershipTransferred)
				if err := _Ownable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Ownable *OwnableFilterer) ParseOwnershipTransferred(log types.Log) (*OwnableOwnershipTransferred, error) {
	event := new(OwnableOwnershipTransferred)
	if err := _Ownable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RequestableIABI is the input ABI used to generate the binding from.
const RequestableIABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"isExit\",\"type\":\"bool\"},{\"name\":\"requestId\",\"type\":\"uint256\"},{\"name\":\"requestor\",\"type\":\"address\"},{\"name\":\"trieKey\",\"type\":\"bytes32\"},{\"name\":\"trieValue\",\"type\":\"bytes\"}],\"name\":\"applyRequestInChildChain\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"isExit\",\"type\":\"bool\"},{\"name\":\"requestId\",\"type\":\"uint256\"},{\"name\":\"requestor\",\"type\":\"address\"},{\"name\":\"trieKey\",\"type\":\"bytes32\"},{\"name\":\"trieValue\",\"type\":\"bytes\"}],\"name\":\"applyRequestInRootChain\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// RequestableIFuncSigs maps the 4-byte function signature to its string representation.
var RequestableIFuncSigs = map[string]string{
	"141ecf46": "applyRequestInChildChain(bool,uint256,address,bytes32,bytes)",
	"a9f79308": "applyRequestInRootChain(bool,uint256,address,bytes32,bytes)",
}

// RequestableI is an auto generated Go binding around an Ethereum contract.
type RequestableI struct {
	RequestableICaller     // Read-only binding to the contract
	RequestableITransactor // Write-only binding to the contract
	RequestableIFilterer   // Log filterer for contract events
}

// RequestableICaller is an auto generated read-only Go binding around an Ethereum contract.
type RequestableICaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RequestableITransactor is an auto generated write-only Go binding around an Ethereum contract.
type RequestableITransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RequestableIFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RequestableIFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RequestableISession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RequestableISession struct {
	Contract     *RequestableI     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RequestableICallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RequestableICallerSession struct {
	Contract *RequestableICaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// RequestableITransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RequestableITransactorSession struct {
	Contract     *RequestableITransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// RequestableIRaw is an auto generated low-level Go binding around an Ethereum contract.
type RequestableIRaw struct {
	Contract *RequestableI // Generic contract binding to access the raw methods on
}

// RequestableICallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RequestableICallerRaw struct {
	Contract *RequestableICaller // Generic read-only contract binding to access the raw methods on
}

// RequestableITransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RequestableITransactorRaw struct {
	Contract *RequestableITransactor // Generic write-only contract binding to access the raw methods on
}

// NewRequestableI creates a new instance of RequestableI, bound to a specific deployed contract.
func NewRequestableI(address common.Address, backend bind.ContractBackend) (*RequestableI, error) {
	contract, err := bindRequestableI(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RequestableI{RequestableICaller: RequestableICaller{contract: contract}, RequestableITransactor: RequestableITransactor{contract: contract}, RequestableIFilterer: RequestableIFilterer{contract: contract}}, nil
}

// NewRequestableICaller creates a new read-only instance of RequestableI, bound to a specific deployed contract.
func NewRequestableICaller(address common.Address, caller bind.ContractCaller) (*RequestableICaller, error) {
	contract, err := bindRequestableI(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RequestableICaller{contract: contract}, nil
}

// NewRequestableITransactor creates a new write-only instance of RequestableI, bound to a specific deployed contract.
func NewRequestableITransactor(address common.Address, transactor bind.ContractTransactor) (*RequestableITransactor, error) {
	contract, err := bindRequestableI(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RequestableITransactor{contract: contract}, nil
}

// NewRequestableIFilterer creates a new log filterer instance of RequestableI, bound to a specific deployed contract.
func NewRequestableIFilterer(address common.Address, filterer bind.ContractFilterer) (*RequestableIFilterer, error) {
	contract, err := bindRequestableI(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RequestableIFilterer{contract: contract}, nil
}

// bindRequestableI binds a generic wrapper to an already deployed contract.
func bindRequestableI(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RequestableIABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RequestableI *RequestableIRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RequestableI.Contract.RequestableICaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RequestableI *RequestableIRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RequestableI.Contract.RequestableITransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RequestableI *RequestableIRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RequestableI.Contract.RequestableITransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RequestableI *RequestableICallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RequestableI.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RequestableI *RequestableITransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RequestableI.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RequestableI *RequestableITransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RequestableI.Contract.contract.Transact(opts, method, params...)
}

// ApplyRequestInChildChain is a paid mutator transaction binding the contract method 0x141ecf46.
//
// Solidity: function applyRequestInChildChain(bool isExit, uint256 requestId, address requestor, bytes32 trieKey, bytes trieValue) returns(bool success)
func (_RequestableI *RequestableITransactor) ApplyRequestInChildChain(opts *bind.TransactOpts, isExit bool, requestId *big.Int, requestor common.Address, trieKey [32]byte, trieValue []byte) (*types.Transaction, error) {
	return _RequestableI.contract.Transact(opts, "applyRequestInChildChain", isExit, requestId, requestor, trieKey, trieValue)
}

// ApplyRequestInChildChain is a paid mutator transaction binding the contract method 0x141ecf46.
//
// Solidity: function applyRequestInChildChain(bool isExit, uint256 requestId, address requestor, bytes32 trieKey, bytes trieValue) returns(bool success)
func (_RequestableI *RequestableISession) ApplyRequestInChildChain(isExit bool, requestId *big.Int, requestor common.Address, trieKey [32]byte, trieValue []byte) (*types.Transaction, error) {
	return _RequestableI.Contract.ApplyRequestInChildChain(&_RequestableI.TransactOpts, isExit, requestId, requestor, trieKey, trieValue)
}

// ApplyRequestInChildChain is a paid mutator transaction binding the contract method 0x141ecf46.
//
// Solidity: function applyRequestInChildChain(bool isExit, uint256 requestId, address requestor, bytes32 trieKey, bytes trieValue) returns(bool success)
func (_RequestableI *RequestableITransactorSession) ApplyRequestInChildChain(isExit bool, requestId *big.Int, requestor common.Address, trieKey [32]byte, trieValue []byte) (*types.Transaction, error) {
	return _RequestableI.Contract.ApplyRequestInChildChain(&_RequestableI.TransactOpts, isExit, requestId, requestor, trieKey, trieValue)
}

// ApplyRequestInRootChain is a paid mutator transaction binding the contract method 0xa9f79308.
//
// Solidity: function applyRequestInRootChain(bool isExit, uint256 requestId, address requestor, bytes32 trieKey, bytes trieValue) returns(bool success)
func (_RequestableI *RequestableITransactor) ApplyRequestInRootChain(opts *bind.TransactOpts, isExit bool, requestId *big.Int, requestor common.Address, trieKey [32]byte, trieValue []byte) (*types.Transaction, error) {
	return _RequestableI.contract.Transact(opts, "applyRequestInRootChain", isExit, requestId, requestor, trieKey, trieValue)
}

// ApplyRequestInRootChain is a paid mutator transaction binding the contract method 0xa9f79308.
//
// Solidity: function applyRequestInRootChain(bool isExit, uint256 requestId, address requestor, bytes32 trieKey, bytes trieValue) returns(bool success)
func (_RequestableI *RequestableISession) ApplyRequestInRootChain(isExit bool, requestId *big.Int, requestor common.Address, trieKey [32]byte, trieValue []byte) (*types.Transaction, error) {
	return _RequestableI.Contract.ApplyRequestInRootChain(&_RequestableI.TransactOpts, isExit, requestId, requestor, trieKey, trieValue)
}

// ApplyRequestInRootChain is a paid mutator transaction binding the contract method 0xa9f79308.
//
// Solidity: function applyRequestInRootChain(bool isExit, uint256 requestId, address requestor, bytes32 trieKey, bytes trieValue) returns(bool success)
func (_RequestableI *RequestableITransactorSession) ApplyRequestInRootChain(isExit bool, requestId *big.Int, requestor common.Address, trieKey [32]byte, trieValue []byte) (*types.Transaction, error) {
	return _RequestableI.Contract.ApplyRequestInRootChain(&_RequestableI.TransactOpts, isExit, requestId, requestor, trieKey, trieValue)
}

// RequestableSimpleTokenABI is the input ABI used to generate the binding from.
const RequestableSimpleTokenABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"isExit\",\"type\":\"bool\"},{\"name\":\"requestId\",\"type\":\"uint256\"},{\"name\":\"requestor\",\"type\":\"address\"},{\"name\":\"trieKey\",\"type\":\"bytes32\"},{\"name\":\"trieValue\",\"type\":\"bytes\"}],\"name\":\"applyRequestInChildChain\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"balances\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"isExit\",\"type\":\"bool\"},{\"name\":\"requestId\",\"type\":\"uint256\"},{\"name\":\"requestor\",\"type\":\"address\"},{\"name\":\"trieKey\",\"type\":\"bytes32\"},{\"name\":\"trieValue\",\"type\":\"bytes\"}],\"name\":\"applyRequestInRootChain\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"who\",\"type\":\"address\"}],\"name\":\"getBalanceTrieKey\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"trieValue\",\"type\":\"bytes\"}],\"name\":\"decodeTrieValue\",\"outputs\":[{\"name\":\"v\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_isExit\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"_requestor\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_trieKey\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"_trieValue\",\"type\":\"bytes\"}],\"name\":\"Request\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// RequestableSimpleTokenFuncSigs maps the 4-byte function signature to its string representation.
var RequestableSimpleTokenFuncSigs = map[string]string{
	"141ecf46": "applyRequestInChildChain(bool,uint256,address,bytes32,bytes)",
	"a9f79308": "applyRequestInRootChain(bool,uint256,address,bytes32,bytes)",
	"27e235e3": "balances(address)",
	"b9e59d09": "decodeTrieValue(bytes)",
	"b18fcfdf": "getBalanceTrieKey(address)",
	"40c10f19": "mint(address,uint256)",
	"8da5cb5b": "owner()",
	"715018a6": "renounceOwnership()",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
	"f2fde38b": "transferOwnership(address)",
}

// RequestableSimpleTokenBin is the compiled bytecode used for deploying new contracts.
var RequestableSimpleTokenBin = "0x6080604052600080546001600160a01b03191633179055610e80806100256000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c80638da5cb5b116100715780638da5cb5b146101cb578063a9059cbb146101ef578063a9f793081461021b578063b18fcfdf146102ae578063b9e59d09146102d4578063f2fde38b1461037a576100a9565b8063141ecf46146100ae57806318160ddd1461015557806327e235e31461016f57806340c10f1914610195578063715018a6146101c3575b600080fd5b610141600480360360a08110156100c457600080fd5b81351515916020810135916001600160a01b036040830135169160608101359181019060a08101608082013564010000000081111561010257600080fd5b82018360208201111561011457600080fd5b8035906020019184600183028401116401000000008311171561013657600080fd5b5090925090506103a0565b604080519115158252519081900360200190f35b61015d610754565b60408051918252519081900360200190f35b61015d6004803603602081101561018557600080fd5b50356001600160a01b031661075a565b6101c1600480360360408110156101ab57600080fd5b506001600160a01b03813516906020013561076c565b005b6101c1610867565b6101d36108c6565b604080516001600160a01b039092168252519081900360200190f35b6101c16004803603604081101561020557600080fd5b506001600160a01b0381351690602001356108d5565b610141600480360360a081101561023157600080fd5b81351515916020810135916001600160a01b036040830135169160608101359181019060a08101608082013564010000000081111561026f57600080fd5b82018360208201111561028157600080fd5b803590602001918460018302840111640100000000831117156102a357600080fd5b509092509050610988565b61015d600480360360208110156102c457600080fd5b50356001600160a01b0316610bea565b61015d600480360360208110156102ea57600080fd5b81019060208101813564010000000081111561030557600080fd5b82018360208201111561031757600080fd5b8035906020019184600183028401116401000000008311171561033957600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610c29945050505050565b6101c16004803603602081101561039057600080fd5b50356001600160a01b0316610c73565b60008581526003602052604081205460ff16156103ee5760405162461bcd60e51b815260040180806020018281038252602f815260200180610d6c602f913960400191505060405180910390fd5b86156105cf578361044a576000546001600160a01b038681169116146104455760405162461bcd60e51b8152600401808060200182810382526029815260200180610e236029913960400191505060405180910390fd5b6105ca565b600184141561048a5760405162461bcd60e51b8152600401808060200182810382526031815260200180610df26031913960400191505060405180910390fd5b8361049486610bea565b1415610593576104d983838080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250610c2992505050565b6001600160a01b038616600090815260026020526040902054101561052f5760405162461bcd60e51b815260040180806020018281038252602f815260200180610dc3602f913960400191505060405180910390fd5b61056e83838080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250610c2992505050565b6001600160a01b038616600090815260026020526040902080549190910390556105ca565b60405162461bcd60e51b8152600401808060200182810382526028815260200180610d9b6028913960400191505060405180910390fd5b6106a3565b836105f457600080546001600160a01b0319166001600160a01b0387161790556106a3565b60018414156106345760405162461bcd60e51b8152600401808060200182810382526031815260200180610df26031913960400191505060405180910390fd5b8361063e86610bea565b14156105935761068383838080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250610c2992505050565b6001600160a01b0386166000908152600260205260409020805490910190555b600086815260036020908152604091829020805460ff19166001179055815189151581526001600160a01b038816918101919091529081018590526080606082018181529082018490527faf049e1d32035f6a4403975248f256e6bc3ac4f5417198f75d25dc5e6522666b918991889188918891889160a08201848480828437600083820152604051601f909101601f19169092018290039850909650505050505050a15060019695505050505050565b60015481565b60026020526000908152604090205481565b6000546001600160a01b0316331461078357600080fd5b600154610796908263ffffffff610c9616565b6001556001600160a01b0382166000908152600260205260409020546107c2908263ffffffff610c9616565b6001600160a01b03831660008181526002602090815260409182902093909355805191825291810183905281517f0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885929181900390910190a160408051600081526001600160a01b038416602082015280820183905290517fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9181900360600190a15050565b6000546001600160a01b0316331461087e57600080fd5b600080546040516001600160a01b03909116917ff8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c6482091a2600080546001600160a01b0319169055565b6000546001600160a01b031681565b336000908152600260205260409020546108f5908263ffffffff610caf16565b33600090815260026020526040808220929092556001600160a01b03841681522054610927908263ffffffff610c9616565b6001600160a01b03831660008181526002602090815260409182902093909355805133815292830191909152818101839052517fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9181900360600190a15050565b60008581526003602052604081205460ff16156109d65760405162461bcd60e51b815260040180806020018281038252602f815260200180610d6c602f913960400191505060405180910390fd5b8615610a825783610a0157600080546001600160a01b0319166001600160a01b0387161790556105ca565b6001841415610a0f576105ca565b83610a1986610bea565b14156105ca57610a5e83838080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250610c2992505050565b6001600160a01b0386166000908152600260205260409020805490910190556105ca565b83610ad3576000546001600160a01b038681169116146105ca5760405162461bcd60e51b8152600401808060200182810382526029815260200180610e236029913960400191505060405180910390fd5b6001841415610b135760405162461bcd60e51b8152600401808060200182810382526031815260200180610df26031913960400191505060405180910390fd5b83610b1d86610bea565b141561059357610b6283838080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250610c2992505050565b6001600160a01b0386166000908152600260205260409020541015610b8657600080fd5b610bc583838080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250610c2992505050565b6001600160a01b038616600090815260026020526040902080549190910390556106a3565b60408051606092831b6bffffffffffffffffffffffff191660208083019190915260028284015282518083038401815293909101909152815191012090565b60008151602014610c6b5760405162461bcd60e51b8152600401808060200182810382526039815260200180610d336039913960400191505060405180910390fd5b506020015190565b6000546001600160a01b03163314610c8a57600080fd5b610c9381610cc4565b50565b600082820183811015610ca857600080fd5b9392505050565b600082821115610cbe57600080fd5b50900390565b6001600160a01b038116610cd757600080fd5b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b039290921691909117905556fe5265717565737461626c6553696d706c65546f6b656e3a206c656e677468206f6620747269652076616c7565206d75737420626520307832305265717565737461626c6553696d706c65546f6b656e3a207265717565737420616c7265616479206170706c6965645265717565737461626c6553696d706c65546f6b656e3a20756e6b6e6f776e2074726965206b65795265717565737453696d70696c65546f6b656e3a206e6f7420656e6f7567682062616c616e636520746f20657869745265717565737461626c6553696d706c65546f6b656e3a2063616e6e6f7420656e74657220746f74616c20737570706c795265717565737453696d706c65546f6b656e3a2072657175657374206d757374206265206f776e6572a265627a7a723058205a68b72a626b5da09b1815631a42e5d737fca125c49dc19d2c6694cb6d05b9a764736f6c63430005090032"

// DeployRequestableSimpleToken deploys a new Ethereum contract, binding an instance of RequestableSimpleToken to it.
func DeployRequestableSimpleToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RequestableSimpleToken, error) {
	parsed, err := abi.JSON(strings.NewReader(RequestableSimpleTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RequestableSimpleTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RequestableSimpleToken{RequestableSimpleTokenCaller: RequestableSimpleTokenCaller{contract: contract}, RequestableSimpleTokenTransactor: RequestableSimpleTokenTransactor{contract: contract}, RequestableSimpleTokenFilterer: RequestableSimpleTokenFilterer{contract: contract}}, nil
}

// RequestableSimpleToken is an auto generated Go binding around an Ethereum contract.
type RequestableSimpleToken struct {
	RequestableSimpleTokenCaller     // Read-only binding to the contract
	RequestableSimpleTokenTransactor // Write-only binding to the contract
	RequestableSimpleTokenFilterer   // Log filterer for contract events
}

// RequestableSimpleTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type RequestableSimpleTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RequestableSimpleTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RequestableSimpleTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RequestableSimpleTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RequestableSimpleTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RequestableSimpleTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RequestableSimpleTokenSession struct {
	Contract     *RequestableSimpleToken // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// RequestableSimpleTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RequestableSimpleTokenCallerSession struct {
	Contract *RequestableSimpleTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// RequestableSimpleTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RequestableSimpleTokenTransactorSession struct {
	Contract     *RequestableSimpleTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// RequestableSimpleTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type RequestableSimpleTokenRaw struct {
	Contract *RequestableSimpleToken // Generic contract binding to access the raw methods on
}

// RequestableSimpleTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RequestableSimpleTokenCallerRaw struct {
	Contract *RequestableSimpleTokenCaller // Generic read-only contract binding to access the raw methods on
}

// RequestableSimpleTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RequestableSimpleTokenTransactorRaw struct {
	Contract *RequestableSimpleTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRequestableSimpleToken creates a new instance of RequestableSimpleToken, bound to a specific deployed contract.
func NewRequestableSimpleToken(address common.Address, backend bind.ContractBackend) (*RequestableSimpleToken, error) {
	contract, err := bindRequestableSimpleToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RequestableSimpleToken{RequestableSimpleTokenCaller: RequestableSimpleTokenCaller{contract: contract}, RequestableSimpleTokenTransactor: RequestableSimpleTokenTransactor{contract: contract}, RequestableSimpleTokenFilterer: RequestableSimpleTokenFilterer{contract: contract}}, nil
}

// NewRequestableSimpleTokenCaller creates a new read-only instance of RequestableSimpleToken, bound to a specific deployed contract.
func NewRequestableSimpleTokenCaller(address common.Address, caller bind.ContractCaller) (*RequestableSimpleTokenCaller, error) {
	contract, err := bindRequestableSimpleToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RequestableSimpleTokenCaller{contract: contract}, nil
}

// NewRequestableSimpleTokenTransactor creates a new write-only instance of RequestableSimpleToken, bound to a specific deployed contract.
func NewRequestableSimpleTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*RequestableSimpleTokenTransactor, error) {
	contract, err := bindRequestableSimpleToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RequestableSimpleTokenTransactor{contract: contract}, nil
}

// NewRequestableSimpleTokenFilterer creates a new log filterer instance of RequestableSimpleToken, bound to a specific deployed contract.
func NewRequestableSimpleTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*RequestableSimpleTokenFilterer, error) {
	contract, err := bindRequestableSimpleToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RequestableSimpleTokenFilterer{contract: contract}, nil
}

// bindRequestableSimpleToken binds a generic wrapper to an already deployed contract.
func bindRequestableSimpleToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RequestableSimpleTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RequestableSimpleToken *RequestableSimpleTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RequestableSimpleToken.Contract.RequestableSimpleTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RequestableSimpleToken *RequestableSimpleTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RequestableSimpleToken.Contract.RequestableSimpleTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RequestableSimpleToken *RequestableSimpleTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RequestableSimpleToken.Contract.RequestableSimpleTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RequestableSimpleToken *RequestableSimpleTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RequestableSimpleToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RequestableSimpleToken *RequestableSimpleTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RequestableSimpleToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RequestableSimpleToken *RequestableSimpleTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RequestableSimpleToken.Contract.contract.Transact(opts, method, params...)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) constant returns(uint256)
func (_RequestableSimpleToken *RequestableSimpleTokenCaller) Balances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RequestableSimpleToken.contract.Call(opts, out, "balances", arg0)
	return *ret0, err
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) constant returns(uint256)
func (_RequestableSimpleToken *RequestableSimpleTokenSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _RequestableSimpleToken.Contract.Balances(&_RequestableSimpleToken.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) constant returns(uint256)
func (_RequestableSimpleToken *RequestableSimpleTokenCallerSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _RequestableSimpleToken.Contract.Balances(&_RequestableSimpleToken.CallOpts, arg0)
}

// DecodeTrieValue is a free data retrieval call binding the contract method 0xb9e59d09.
//
// Solidity: function decodeTrieValue(bytes trieValue) constant returns(uint256 v)
func (_RequestableSimpleToken *RequestableSimpleTokenCaller) DecodeTrieValue(opts *bind.CallOpts, trieValue []byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RequestableSimpleToken.contract.Call(opts, out, "decodeTrieValue", trieValue)
	return *ret0, err
}

// DecodeTrieValue is a free data retrieval call binding the contract method 0xb9e59d09.
//
// Solidity: function decodeTrieValue(bytes trieValue) constant returns(uint256 v)
func (_RequestableSimpleToken *RequestableSimpleTokenSession) DecodeTrieValue(trieValue []byte) (*big.Int, error) {
	return _RequestableSimpleToken.Contract.DecodeTrieValue(&_RequestableSimpleToken.CallOpts, trieValue)
}

// DecodeTrieValue is a free data retrieval call binding the contract method 0xb9e59d09.
//
// Solidity: function decodeTrieValue(bytes trieValue) constant returns(uint256 v)
func (_RequestableSimpleToken *RequestableSimpleTokenCallerSession) DecodeTrieValue(trieValue []byte) (*big.Int, error) {
	return _RequestableSimpleToken.Contract.DecodeTrieValue(&_RequestableSimpleToken.CallOpts, trieValue)
}

// GetBalanceTrieKey is a free data retrieval call binding the contract method 0xb18fcfdf.
//
// Solidity: function getBalanceTrieKey(address who) constant returns(bytes32)
func (_RequestableSimpleToken *RequestableSimpleTokenCaller) GetBalanceTrieKey(opts *bind.CallOpts, who common.Address) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _RequestableSimpleToken.contract.Call(opts, out, "getBalanceTrieKey", who)
	return *ret0, err
}

// GetBalanceTrieKey is a free data retrieval call binding the contract method 0xb18fcfdf.
//
// Solidity: function getBalanceTrieKey(address who) constant returns(bytes32)
func (_RequestableSimpleToken *RequestableSimpleTokenSession) GetBalanceTrieKey(who common.Address) ([32]byte, error) {
	return _RequestableSimpleToken.Contract.GetBalanceTrieKey(&_RequestableSimpleToken.CallOpts, who)
}

// GetBalanceTrieKey is a free data retrieval call binding the contract method 0xb18fcfdf.
//
// Solidity: function getBalanceTrieKey(address who) constant returns(bytes32)
func (_RequestableSimpleToken *RequestableSimpleTokenCallerSession) GetBalanceTrieKey(who common.Address) ([32]byte, error) {
	return _RequestableSimpleToken.Contract.GetBalanceTrieKey(&_RequestableSimpleToken.CallOpts, who)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_RequestableSimpleToken *RequestableSimpleTokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RequestableSimpleToken.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_RequestableSimpleToken *RequestableSimpleTokenSession) Owner() (common.Address, error) {
	return _RequestableSimpleToken.Contract.Owner(&_RequestableSimpleToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_RequestableSimpleToken *RequestableSimpleTokenCallerSession) Owner() (common.Address, error) {
	return _RequestableSimpleToken.Contract.Owner(&_RequestableSimpleToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_RequestableSimpleToken *RequestableSimpleTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RequestableSimpleToken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_RequestableSimpleToken *RequestableSimpleTokenSession) TotalSupply() (*big.Int, error) {
	return _RequestableSimpleToken.Contract.TotalSupply(&_RequestableSimpleToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_RequestableSimpleToken *RequestableSimpleTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _RequestableSimpleToken.Contract.TotalSupply(&_RequestableSimpleToken.CallOpts)
}

// ApplyRequestInChildChain is a paid mutator transaction binding the contract method 0x141ecf46.
//
// Solidity: function applyRequestInChildChain(bool isExit, uint256 requestId, address requestor, bytes32 trieKey, bytes trieValue) returns(bool success)
func (_RequestableSimpleToken *RequestableSimpleTokenTransactor) ApplyRequestInChildChain(opts *bind.TransactOpts, isExit bool, requestId *big.Int, requestor common.Address, trieKey [32]byte, trieValue []byte) (*types.Transaction, error) {
	return _RequestableSimpleToken.contract.Transact(opts, "applyRequestInChildChain", isExit, requestId, requestor, trieKey, trieValue)
}

// ApplyRequestInChildChain is a paid mutator transaction binding the contract method 0x141ecf46.
//
// Solidity: function applyRequestInChildChain(bool isExit, uint256 requestId, address requestor, bytes32 trieKey, bytes trieValue) returns(bool success)
func (_RequestableSimpleToken *RequestableSimpleTokenSession) ApplyRequestInChildChain(isExit bool, requestId *big.Int, requestor common.Address, trieKey [32]byte, trieValue []byte) (*types.Transaction, error) {
	return _RequestableSimpleToken.Contract.ApplyRequestInChildChain(&_RequestableSimpleToken.TransactOpts, isExit, requestId, requestor, trieKey, trieValue)
}

// ApplyRequestInChildChain is a paid mutator transaction binding the contract method 0x141ecf46.
//
// Solidity: function applyRequestInChildChain(bool isExit, uint256 requestId, address requestor, bytes32 trieKey, bytes trieValue) returns(bool success)
func (_RequestableSimpleToken *RequestableSimpleTokenTransactorSession) ApplyRequestInChildChain(isExit bool, requestId *big.Int, requestor common.Address, trieKey [32]byte, trieValue []byte) (*types.Transaction, error) {
	return _RequestableSimpleToken.Contract.ApplyRequestInChildChain(&_RequestableSimpleToken.TransactOpts, isExit, requestId, requestor, trieKey, trieValue)
}

// ApplyRequestInRootChain is a paid mutator transaction binding the contract method 0xa9f79308.
//
// Solidity: function applyRequestInRootChain(bool isExit, uint256 requestId, address requestor, bytes32 trieKey, bytes trieValue) returns(bool success)
func (_RequestableSimpleToken *RequestableSimpleTokenTransactor) ApplyRequestInRootChain(opts *bind.TransactOpts, isExit bool, requestId *big.Int, requestor common.Address, trieKey [32]byte, trieValue []byte) (*types.Transaction, error) {
	return _RequestableSimpleToken.contract.Transact(opts, "applyRequestInRootChain", isExit, requestId, requestor, trieKey, trieValue)
}

// ApplyRequestInRootChain is a paid mutator transaction binding the contract method 0xa9f79308.
//
// Solidity: function applyRequestInRootChain(bool isExit, uint256 requestId, address requestor, bytes32 trieKey, bytes trieValue) returns(bool success)
func (_RequestableSimpleToken *RequestableSimpleTokenSession) ApplyRequestInRootChain(isExit bool, requestId *big.Int, requestor common.Address, trieKey [32]byte, trieValue []byte) (*types.Transaction, error) {
	return _RequestableSimpleToken.Contract.ApplyRequestInRootChain(&_RequestableSimpleToken.TransactOpts, isExit, requestId, requestor, trieKey, trieValue)
}

// ApplyRequestInRootChain is a paid mutator transaction binding the contract method 0xa9f79308.
//
// Solidity: function applyRequestInRootChain(bool isExit, uint256 requestId, address requestor, bytes32 trieKey, bytes trieValue) returns(bool success)
func (_RequestableSimpleToken *RequestableSimpleTokenTransactorSession) ApplyRequestInRootChain(isExit bool, requestId *big.Int, requestor common.Address, trieKey [32]byte, trieValue []byte) (*types.Transaction, error) {
	return _RequestableSimpleToken.Contract.ApplyRequestInRootChain(&_RequestableSimpleToken.TransactOpts, isExit, requestId, requestor, trieKey, trieValue)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _value) returns()
func (_RequestableSimpleToken *RequestableSimpleTokenTransactor) Mint(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _RequestableSimpleToken.contract.Transact(opts, "mint", _to, _value)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _value) returns()
func (_RequestableSimpleToken *RequestableSimpleTokenSession) Mint(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _RequestableSimpleToken.Contract.Mint(&_RequestableSimpleToken.TransactOpts, _to, _value)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _value) returns()
func (_RequestableSimpleToken *RequestableSimpleTokenTransactorSession) Mint(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _RequestableSimpleToken.Contract.Mint(&_RequestableSimpleToken.TransactOpts, _to, _value)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RequestableSimpleToken *RequestableSimpleTokenTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RequestableSimpleToken.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RequestableSimpleToken *RequestableSimpleTokenSession) RenounceOwnership() (*types.Transaction, error) {
	return _RequestableSimpleToken.Contract.RenounceOwnership(&_RequestableSimpleToken.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RequestableSimpleToken *RequestableSimpleTokenTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _RequestableSimpleToken.Contract.RenounceOwnership(&_RequestableSimpleToken.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns()
func (_RequestableSimpleToken *RequestableSimpleTokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _RequestableSimpleToken.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns()
func (_RequestableSimpleToken *RequestableSimpleTokenSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _RequestableSimpleToken.Contract.Transfer(&_RequestableSimpleToken.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns()
func (_RequestableSimpleToken *RequestableSimpleTokenTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _RequestableSimpleToken.Contract.Transfer(&_RequestableSimpleToken.TransactOpts, _to, _value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_RequestableSimpleToken *RequestableSimpleTokenTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _RequestableSimpleToken.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_RequestableSimpleToken *RequestableSimpleTokenSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _RequestableSimpleToken.Contract.TransferOwnership(&_RequestableSimpleToken.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_RequestableSimpleToken *RequestableSimpleTokenTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _RequestableSimpleToken.Contract.TransferOwnership(&_RequestableSimpleToken.TransactOpts, _newOwner)
}

// RequestableSimpleTokenMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the RequestableSimpleToken contract.
type RequestableSimpleTokenMintIterator struct {
	Event *RequestableSimpleTokenMint // Event containing the contract specifics and raw log

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
func (it *RequestableSimpleTokenMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RequestableSimpleTokenMint)
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
		it.Event = new(RequestableSimpleTokenMint)
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
func (it *RequestableSimpleTokenMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RequestableSimpleTokenMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RequestableSimpleTokenMint represents a Mint event raised by the RequestableSimpleToken contract.
type RequestableSimpleTokenMint struct {
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(address _to, uint256 _value)
func (_RequestableSimpleToken *RequestableSimpleTokenFilterer) FilterMint(opts *bind.FilterOpts) (*RequestableSimpleTokenMintIterator, error) {

	logs, sub, err := _RequestableSimpleToken.contract.FilterLogs(opts, "Mint")
	if err != nil {
		return nil, err
	}
	return &RequestableSimpleTokenMintIterator{contract: _RequestableSimpleToken.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(address _to, uint256 _value)
func (_RequestableSimpleToken *RequestableSimpleTokenFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *RequestableSimpleTokenMint) (event.Subscription, error) {

	logs, sub, err := _RequestableSimpleToken.contract.WatchLogs(opts, "Mint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RequestableSimpleTokenMint)
				if err := _RequestableSimpleToken.contract.UnpackLog(event, "Mint", log); err != nil {
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

// ParseMint is a log parse operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(address _to, uint256 _value)
func (_RequestableSimpleToken *RequestableSimpleTokenFilterer) ParseMint(log types.Log) (*RequestableSimpleTokenMint, error) {
	event := new(RequestableSimpleTokenMint)
	if err := _RequestableSimpleToken.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RequestableSimpleTokenOwnershipRenouncedIterator is returned from FilterOwnershipRenounced and is used to iterate over the raw logs and unpacked data for OwnershipRenounced events raised by the RequestableSimpleToken contract.
type RequestableSimpleTokenOwnershipRenouncedIterator struct {
	Event *RequestableSimpleTokenOwnershipRenounced // Event containing the contract specifics and raw log

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
func (it *RequestableSimpleTokenOwnershipRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RequestableSimpleTokenOwnershipRenounced)
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
		it.Event = new(RequestableSimpleTokenOwnershipRenounced)
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
func (it *RequestableSimpleTokenOwnershipRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RequestableSimpleTokenOwnershipRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RequestableSimpleTokenOwnershipRenounced represents a OwnershipRenounced event raised by the RequestableSimpleToken contract.
type RequestableSimpleTokenOwnershipRenounced struct {
	PreviousOwner common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipRenounced is a free log retrieval operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: event OwnershipRenounced(address indexed previousOwner)
func (_RequestableSimpleToken *RequestableSimpleTokenFilterer) FilterOwnershipRenounced(opts *bind.FilterOpts, previousOwner []common.Address) (*RequestableSimpleTokenOwnershipRenouncedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _RequestableSimpleToken.contract.FilterLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RequestableSimpleTokenOwnershipRenouncedIterator{contract: _RequestableSimpleToken.contract, event: "OwnershipRenounced", logs: logs, sub: sub}, nil
}

// WatchOwnershipRenounced is a free log subscription operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: event OwnershipRenounced(address indexed previousOwner)
func (_RequestableSimpleToken *RequestableSimpleTokenFilterer) WatchOwnershipRenounced(opts *bind.WatchOpts, sink chan<- *RequestableSimpleTokenOwnershipRenounced, previousOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _RequestableSimpleToken.contract.WatchLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RequestableSimpleTokenOwnershipRenounced)
				if err := _RequestableSimpleToken.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
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

// ParseOwnershipRenounced is a log parse operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: event OwnershipRenounced(address indexed previousOwner)
func (_RequestableSimpleToken *RequestableSimpleTokenFilterer) ParseOwnershipRenounced(log types.Log) (*RequestableSimpleTokenOwnershipRenounced, error) {
	event := new(RequestableSimpleTokenOwnershipRenounced)
	if err := _RequestableSimpleToken.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RequestableSimpleTokenOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the RequestableSimpleToken contract.
type RequestableSimpleTokenOwnershipTransferredIterator struct {
	Event *RequestableSimpleTokenOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *RequestableSimpleTokenOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RequestableSimpleTokenOwnershipTransferred)
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
		it.Event = new(RequestableSimpleTokenOwnershipTransferred)
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
func (it *RequestableSimpleTokenOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RequestableSimpleTokenOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RequestableSimpleTokenOwnershipTransferred represents a OwnershipTransferred event raised by the RequestableSimpleToken contract.
type RequestableSimpleTokenOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_RequestableSimpleToken *RequestableSimpleTokenFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*RequestableSimpleTokenOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RequestableSimpleToken.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RequestableSimpleTokenOwnershipTransferredIterator{contract: _RequestableSimpleToken.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_RequestableSimpleToken *RequestableSimpleTokenFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RequestableSimpleTokenOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RequestableSimpleToken.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RequestableSimpleTokenOwnershipTransferred)
				if err := _RequestableSimpleToken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_RequestableSimpleToken *RequestableSimpleTokenFilterer) ParseOwnershipTransferred(log types.Log) (*RequestableSimpleTokenOwnershipTransferred, error) {
	event := new(RequestableSimpleTokenOwnershipTransferred)
	if err := _RequestableSimpleToken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RequestableSimpleTokenRequestIterator is returned from FilterRequest and is used to iterate over the raw logs and unpacked data for Request events raised by the RequestableSimpleToken contract.
type RequestableSimpleTokenRequestIterator struct {
	Event *RequestableSimpleTokenRequest // Event containing the contract specifics and raw log

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
func (it *RequestableSimpleTokenRequestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RequestableSimpleTokenRequest)
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
		it.Event = new(RequestableSimpleTokenRequest)
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
func (it *RequestableSimpleTokenRequestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RequestableSimpleTokenRequestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RequestableSimpleTokenRequest represents a Request event raised by the RequestableSimpleToken contract.
type RequestableSimpleTokenRequest struct {
	IsExit    bool
	Requestor common.Address
	TrieKey   [32]byte
	TrieValue []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRequest is a free log retrieval operation binding the contract event 0xaf049e1d32035f6a4403975248f256e6bc3ac4f5417198f75d25dc5e6522666b.
//
// Solidity: event Request(bool _isExit, address _requestor, bytes32 _trieKey, bytes _trieValue)
func (_RequestableSimpleToken *RequestableSimpleTokenFilterer) FilterRequest(opts *bind.FilterOpts) (*RequestableSimpleTokenRequestIterator, error) {

	logs, sub, err := _RequestableSimpleToken.contract.FilterLogs(opts, "Request")
	if err != nil {
		return nil, err
	}
	return &RequestableSimpleTokenRequestIterator{contract: _RequestableSimpleToken.contract, event: "Request", logs: logs, sub: sub}, nil
}

// WatchRequest is a free log subscription operation binding the contract event 0xaf049e1d32035f6a4403975248f256e6bc3ac4f5417198f75d25dc5e6522666b.
//
// Solidity: event Request(bool _isExit, address _requestor, bytes32 _trieKey, bytes _trieValue)
func (_RequestableSimpleToken *RequestableSimpleTokenFilterer) WatchRequest(opts *bind.WatchOpts, sink chan<- *RequestableSimpleTokenRequest) (event.Subscription, error) {

	logs, sub, err := _RequestableSimpleToken.contract.WatchLogs(opts, "Request")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RequestableSimpleTokenRequest)
				if err := _RequestableSimpleToken.contract.UnpackLog(event, "Request", log); err != nil {
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

// ParseRequest is a log parse operation binding the contract event 0xaf049e1d32035f6a4403975248f256e6bc3ac4f5417198f75d25dc5e6522666b.
//
// Solidity: event Request(bool _isExit, address _requestor, bytes32 _trieKey, bytes _trieValue)
func (_RequestableSimpleToken *RequestableSimpleTokenFilterer) ParseRequest(log types.Log) (*RequestableSimpleTokenRequest, error) {
	event := new(RequestableSimpleTokenRequest)
	if err := _RequestableSimpleToken.contract.UnpackLog(event, "Request", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RequestableSimpleTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the RequestableSimpleToken contract.
type RequestableSimpleTokenTransferIterator struct {
	Event *RequestableSimpleTokenTransfer // Event containing the contract specifics and raw log

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
func (it *RequestableSimpleTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RequestableSimpleTokenTransfer)
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
		it.Event = new(RequestableSimpleTokenTransfer)
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
func (it *RequestableSimpleTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RequestableSimpleTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RequestableSimpleTokenTransfer represents a Transfer event raised by the RequestableSimpleToken contract.
type RequestableSimpleTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address _from, address _to, uint256 _value)
func (_RequestableSimpleToken *RequestableSimpleTokenFilterer) FilterTransfer(opts *bind.FilterOpts) (*RequestableSimpleTokenTransferIterator, error) {

	logs, sub, err := _RequestableSimpleToken.contract.FilterLogs(opts, "Transfer")
	if err != nil {
		return nil, err
	}
	return &RequestableSimpleTokenTransferIterator{contract: _RequestableSimpleToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address _from, address _to, uint256 _value)
func (_RequestableSimpleToken *RequestableSimpleTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *RequestableSimpleTokenTransfer) (event.Subscription, error) {

	logs, sub, err := _RequestableSimpleToken.contract.WatchLogs(opts, "Transfer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RequestableSimpleTokenTransfer)
				if err := _RequestableSimpleToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address _from, address _to, uint256 _value)
func (_RequestableSimpleToken *RequestableSimpleTokenFilterer) ParseTransfer(log types.Log) (*RequestableSimpleTokenTransfer, error) {
	event := new(RequestableSimpleTokenTransfer)
	if err := _RequestableSimpleToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
var SafeMathBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a723058201da5f5be29b57b3320a859ffcc42395bd583e549ffc40551affe4a01a35b887d64736f6c63430005090032"

// DeploySafeMath deploys a new Ethereum contract, binding an instance of SafeMath to it.
func DeploySafeMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMath, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// SafeMath is an auto generated Go binding around an Ethereum contract.
type SafeMath struct {
	SafeMathCaller     // Read-only binding to the contract
	SafeMathTransactor // Write-only binding to the contract
	SafeMathFilterer   // Log filterer for contract events
}

// SafeMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeMathSession struct {
	Contract     *SafeMath         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeMathCallerSession struct {
	Contract *SafeMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SafeMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeMathTransactorSession struct {
	Contract     *SafeMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SafeMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeMathRaw struct {
	Contract *SafeMath // Generic contract binding to access the raw methods on
}

// SafeMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeMathCallerRaw struct {
	Contract *SafeMathCaller // Generic read-only contract binding to access the raw methods on
}

// SafeMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeMathTransactorRaw struct {
	Contract *SafeMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeMath creates a new instance of SafeMath, bound to a specific deployed contract.
func NewSafeMath(address common.Address, backend bind.ContractBackend) (*SafeMath, error) {
	contract, err := bindSafeMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// NewSafeMathCaller creates a new read-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathCaller(address common.Address, caller bind.ContractCaller) (*SafeMathCaller, error) {
	contract, err := bindSafeMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathCaller{contract: contract}, nil
}

// NewSafeMathTransactor creates a new write-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathTransactor, error) {
	contract, err := bindSafeMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathTransactor{contract: contract}, nil
}

// NewSafeMathFilterer creates a new log filterer instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeMathFilterer, error) {
	contract, err := bindSafeMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeMathFilterer{contract: contract}, nil
}

// bindSafeMath binds a generic wrapper to an already deployed contract.
func bindSafeMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.SafeMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transact(opts, method, params...)
}
