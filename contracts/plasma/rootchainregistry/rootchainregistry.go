// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package rootchainregistry

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

// ContextABI is the input ABI used to generate the binding from.
const ContextABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// Context is an auto generated Go binding around an Ethereum contract.
type Context struct {
	ContextCaller     // Read-only binding to the contract
	ContextTransactor // Write-only binding to the contract
	ContextFilterer   // Log filterer for contract events
}

// ContextCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContextCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContextTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContextFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContextSession struct {
	Contract     *Context          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContextCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContextCallerSession struct {
	Contract *ContextCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ContextTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContextTransactorSession struct {
	Contract     *ContextTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ContextRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContextRaw struct {
	Contract *Context // Generic contract binding to access the raw methods on
}

// ContextCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContextCallerRaw struct {
	Contract *ContextCaller // Generic read-only contract binding to access the raw methods on
}

// ContextTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContextTransactorRaw struct {
	Contract *ContextTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContext creates a new instance of Context, bound to a specific deployed contract.
func NewContext(address common.Address, backend bind.ContractBackend) (*Context, error) {
	contract, err := bindContext(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Context{ContextCaller: ContextCaller{contract: contract}, ContextTransactor: ContextTransactor{contract: contract}, ContextFilterer: ContextFilterer{contract: contract}}, nil
}

// NewContextCaller creates a new read-only instance of Context, bound to a specific deployed contract.
func NewContextCaller(address common.Address, caller bind.ContractCaller) (*ContextCaller, error) {
	contract, err := bindContext(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContextCaller{contract: contract}, nil
}

// NewContextTransactor creates a new write-only instance of Context, bound to a specific deployed contract.
func NewContextTransactor(address common.Address, transactor bind.ContractTransactor) (*ContextTransactor, error) {
	contract, err := bindContext(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContextTransactor{contract: contract}, nil
}

// NewContextFilterer creates a new log filterer instance of Context, bound to a specific deployed contract.
func NewContextFilterer(address common.Address, filterer bind.ContractFilterer) (*ContextFilterer, error) {
	contract, err := bindContext(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContextFilterer{contract: contract}, nil
}

// bindContext binds a generic wrapper to an already deployed contract.
func bindContext(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContextABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Context *ContextRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Context.Contract.ContextCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Context *ContextRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Context.Contract.ContextTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Context *ContextRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Context.Contract.ContextTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Context *ContextCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Context.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Context *ContextTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Context.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Context *ContextTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Context.Contract.contract.Transact(opts, method, params...)
}

// OwnableABI is the input ABI used to generate the binding from.
const OwnableABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// OwnableFuncSigs maps the 4-byte function signature to its string representation.
var OwnableFuncSigs = map[string]string{
	"8f32d59b": "isOwner()",
	"8da5cb5b": "owner()",
	"715018a6": "renounceOwnership()",
	"f2fde38b": "transferOwnership(address)",
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

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Ownable *OwnableCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Ownable.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Ownable *OwnableSession) IsOwner() (bool, error) {
	return _Ownable.Contract.IsOwner(&_Ownable.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Ownable *OwnableCallerSession) IsOwner() (bool, error) {
	return _Ownable.Contract.IsOwner(&_Ownable.CallOpts)
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
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
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

// RootChainIABI is the input ABI used to generate the binding from.
const RootChainIABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"currentFork\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isRootChain\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"forkNumber\",\"type\":\"uint256\"}],\"name\":\"lastEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"operator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// RootChainIFuncSigs maps the 4-byte function signature to its string representation.
var RootChainIFuncSigs = map[string]string{
	"183d2d1c": "currentFork()",
	"420bb4b8": "isRootChain()",
	"11e4c914": "lastEpoch(uint256)",
	"570ca735": "operator()",
}

// RootChainI is an auto generated Go binding around an Ethereum contract.
type RootChainI struct {
	RootChainICaller     // Read-only binding to the contract
	RootChainITransactor // Write-only binding to the contract
	RootChainIFilterer   // Log filterer for contract events
}

// RootChainICaller is an auto generated read-only Go binding around an Ethereum contract.
type RootChainICaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RootChainITransactor is an auto generated write-only Go binding around an Ethereum contract.
type RootChainITransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RootChainIFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RootChainIFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RootChainISession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RootChainISession struct {
	Contract     *RootChainI       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RootChainICallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RootChainICallerSession struct {
	Contract *RootChainICaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// RootChainITransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RootChainITransactorSession struct {
	Contract     *RootChainITransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// RootChainIRaw is an auto generated low-level Go binding around an Ethereum contract.
type RootChainIRaw struct {
	Contract *RootChainI // Generic contract binding to access the raw methods on
}

// RootChainICallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RootChainICallerRaw struct {
	Contract *RootChainICaller // Generic read-only contract binding to access the raw methods on
}

// RootChainITransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RootChainITransactorRaw struct {
	Contract *RootChainITransactor // Generic write-only contract binding to access the raw methods on
}

// NewRootChainI creates a new instance of RootChainI, bound to a specific deployed contract.
func NewRootChainI(address common.Address, backend bind.ContractBackend) (*RootChainI, error) {
	contract, err := bindRootChainI(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RootChainI{RootChainICaller: RootChainICaller{contract: contract}, RootChainITransactor: RootChainITransactor{contract: contract}, RootChainIFilterer: RootChainIFilterer{contract: contract}}, nil
}

// NewRootChainICaller creates a new read-only instance of RootChainI, bound to a specific deployed contract.
func NewRootChainICaller(address common.Address, caller bind.ContractCaller) (*RootChainICaller, error) {
	contract, err := bindRootChainI(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RootChainICaller{contract: contract}, nil
}

// NewRootChainITransactor creates a new write-only instance of RootChainI, bound to a specific deployed contract.
func NewRootChainITransactor(address common.Address, transactor bind.ContractTransactor) (*RootChainITransactor, error) {
	contract, err := bindRootChainI(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RootChainITransactor{contract: contract}, nil
}

// NewRootChainIFilterer creates a new log filterer instance of RootChainI, bound to a specific deployed contract.
func NewRootChainIFilterer(address common.Address, filterer bind.ContractFilterer) (*RootChainIFilterer, error) {
	contract, err := bindRootChainI(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RootChainIFilterer{contract: contract}, nil
}

// bindRootChainI binds a generic wrapper to an already deployed contract.
func bindRootChainI(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RootChainIABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RootChainI *RootChainIRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RootChainI.Contract.RootChainICaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RootChainI *RootChainIRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RootChainI.Contract.RootChainITransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RootChainI *RootChainIRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RootChainI.Contract.RootChainITransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RootChainI *RootChainICallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RootChainI.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RootChainI *RootChainITransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RootChainI.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RootChainI *RootChainITransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RootChainI.Contract.contract.Transact(opts, method, params...)
}

// CurrentFork is a free data retrieval call binding the contract method 0x183d2d1c.
//
// Solidity: function currentFork() constant returns(uint256)
func (_RootChainI *RootChainICaller) CurrentFork(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChainI.contract.Call(opts, out, "currentFork")
	return *ret0, err
}

// CurrentFork is a free data retrieval call binding the contract method 0x183d2d1c.
//
// Solidity: function currentFork() constant returns(uint256)
func (_RootChainI *RootChainISession) CurrentFork() (*big.Int, error) {
	return _RootChainI.Contract.CurrentFork(&_RootChainI.CallOpts)
}

// CurrentFork is a free data retrieval call binding the contract method 0x183d2d1c.
//
// Solidity: function currentFork() constant returns(uint256)
func (_RootChainI *RootChainICallerSession) CurrentFork() (*big.Int, error) {
	return _RootChainI.Contract.CurrentFork(&_RootChainI.CallOpts)
}

// IsRootChain is a free data retrieval call binding the contract method 0x420bb4b8.
//
// Solidity: function isRootChain() constant returns(bool)
func (_RootChainI *RootChainICaller) IsRootChain(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _RootChainI.contract.Call(opts, out, "isRootChain")
	return *ret0, err
}

// IsRootChain is a free data retrieval call binding the contract method 0x420bb4b8.
//
// Solidity: function isRootChain() constant returns(bool)
func (_RootChainI *RootChainISession) IsRootChain() (bool, error) {
	return _RootChainI.Contract.IsRootChain(&_RootChainI.CallOpts)
}

// IsRootChain is a free data retrieval call binding the contract method 0x420bb4b8.
//
// Solidity: function isRootChain() constant returns(bool)
func (_RootChainI *RootChainICallerSession) IsRootChain() (bool, error) {
	return _RootChainI.Contract.IsRootChain(&_RootChainI.CallOpts)
}

// LastEpoch is a free data retrieval call binding the contract method 0x11e4c914.
//
// Solidity: function lastEpoch(uint256 forkNumber) constant returns(uint256)
func (_RootChainI *RootChainICaller) LastEpoch(opts *bind.CallOpts, forkNumber *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChainI.contract.Call(opts, out, "lastEpoch", forkNumber)
	return *ret0, err
}

// LastEpoch is a free data retrieval call binding the contract method 0x11e4c914.
//
// Solidity: function lastEpoch(uint256 forkNumber) constant returns(uint256)
func (_RootChainI *RootChainISession) LastEpoch(forkNumber *big.Int) (*big.Int, error) {
	return _RootChainI.Contract.LastEpoch(&_RootChainI.CallOpts, forkNumber)
}

// LastEpoch is a free data retrieval call binding the contract method 0x11e4c914.
//
// Solidity: function lastEpoch(uint256 forkNumber) constant returns(uint256)
func (_RootChainI *RootChainICallerSession) LastEpoch(forkNumber *big.Int) (*big.Int, error) {
	return _RootChainI.Contract.LastEpoch(&_RootChainI.CallOpts, forkNumber)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() constant returns(address)
func (_RootChainI *RootChainICaller) Operator(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RootChainI.contract.Call(opts, out, "operator")
	return *ret0, err
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() constant returns(address)
func (_RootChainI *RootChainISession) Operator() (common.Address, error) {
	return _RootChainI.Contract.Operator(&_RootChainI.CallOpts)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() constant returns(address)
func (_RootChainI *RootChainICallerSession) Operator() (common.Address, error) {
	return _RootChainI.Contract.Operator(&_RootChainI.CallOpts)
}

// RootChainRegistryABI is the input ABI used to generate the binding from.
const RootChainRegistryABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"rootchain\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"seigManager\",\"type\":\"address\"}],\"name\":\"deployCoinage\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"numRootChains\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"rootchain\",\"type\":\"address\"}],\"name\":\"register\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"rootchain\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"seigManager\",\"type\":\"address\"}],\"name\":\"registerAndDeployCoinage\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"rootchain\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"seigManager\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"commissionRate\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isCommissionRateNegative\",\"type\":\"bool\"}],\"name\":\"registerAndDeployCoinageAndSetCommissionRate\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"rootchainByIndex\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"rootchain\",\"type\":\"address\"}],\"name\":\"rootchains\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"rootchain\",\"type\":\"address\"}],\"name\":\"unregister\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// RootChainRegistryFuncSigs maps the 4-byte function signature to its string representation.
var RootChainRegistryFuncSigs = map[string]string{
	"85108604": "deployCoinage(address,address)",
	"8f32d59b": "isOwner()",
	"b2b604d0": "numRootChains()",
	"8da5cb5b": "owner()",
	"4420e486": "register(address)",
	"bcb1a71e": "registerAndDeployCoinage(address,address)",
	"3eb2a66c": "registerAndDeployCoinageAndSetCommissionRate(address,address,uint256,bool)",
	"715018a6": "renounceOwnership()",
	"821f602c": "rootchainByIndex(uint256)",
	"02a15299": "rootchains(address)",
	"f2fde38b": "transferOwnership(address)",
	"2ec2c246": "unregister(address)",
}

// RootChainRegistryBin is the compiled bytecode used for deploying new contracts.
var RootChainRegistryBin = "0x608060405260006100176001600160e01b0361006616565b600080546001600160a01b0319166001600160a01b0383169081178255604051929350917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a35061006a565b3390565b610b5e806100796000396000f3fe608060405234801561001057600080fd5b50600436106100b45760003560e01c8063851086041161007157806385108604146101c05780638da5cb5b146101ee5780638f32d59b146101f6578063b2b604d0146101fe578063bcb1a71e14610218578063f2fde38b14610246576100b4565b806302a15299146100b95780632ec2c246146100f35780633eb2a66c146101195780634420e48614610157578063715018a61461017d578063821f602c14610187575b600080fd5b6100df600480360360208110156100cf57600080fd5b50356001600160a01b031661026c565b604080519115158252519081900360200190f35b6100df6004803603602081101561010957600080fd5b50356001600160a01b031661028a565b6100df6004803603608081101561012f57600080fd5b506001600160a01b03813581169160208101359091169060408101359060600135151561032e565b6100df6004803603602081101561016d57600080fd5b50356001600160a01b031661043a565b610185610510565b005b6101a46004803603602081101561019d57600080fd5b50356105b3565b604080516001600160a01b039092168252519081900360200190f35b6100df600480360360408110156101d657600080fd5b506001600160a01b03813581169160200135166105ce565b6101a46106a6565b6100df6106b5565b6102066106d9565b60408051918252519081900360200190f35b6100df6004803603604081101561022e57600080fd5b506001600160a01b03813581169160200135166106df565b6101856004803603602081101561025c57600080fd5b50356001600160a01b03166107d4565b6001600160a01b031660009081526001602052604090205460ff1690565b60006102946106b5565b6102e5576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b6001600160a01b03821660009081526001602052604090205460ff1661030a57600080fd5b6001600160a01b039091166000908152600160205260409020805460ff1916905590565b6000846103396106b5565b806103b95750336001600160a01b0316816001600160a01b031663570ca7356040518163ffffffff1660e01b815260040160206040518083038186803b15801561038257600080fd5b505afa158015610396573d6000803e3d6000fd5b505050506040513d60208110156103ac57600080fd5b50516001600160a01b0316145b6103f45760405162461bcd60e51b8152600401808060200182810382526027815260200180610b036027913960400191505060405180910390fd5b6103fd86610839565b61040657600080fd5b610410868661091b565b61041957600080fd5b610425868686866109a8565b61042e57600080fd5b50600195945050505050565b6000816104456106b5565b806104c55750336001600160a01b0316816001600160a01b031663570ca7356040518163ffffffff1660e01b815260040160206040518083038186803b15801561048e57600080fd5b505afa1580156104a2573d6000803e3d6000fd5b505050506040513d60208110156104b857600080fd5b50516001600160a01b0316145b6105005760405162461bcd60e51b8152600401808060200182810382526027815260200180610b036027913960400191505060405180910390fd5b61050983610839565b9392505050565b6105186106b5565b610569576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b600080546040516001600160a01b03909116907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3600080546001600160a01b0319169055565b6000908152600360205260409020546001600160a01b031690565b6000826105d96106b5565b806106595750336001600160a01b0316816001600160a01b031663570ca7356040518163ffffffff1660e01b815260040160206040518083038186803b15801561062257600080fd5b505afa158015610636573d6000803e3d6000fd5b505050506040513d602081101561064c57600080fd5b50516001600160a01b0316145b6106945760405162461bcd60e51b8152600401808060200182810382526027815260200180610b036027913960400191505060405180910390fd5b61069e848461091b565b949350505050565b6000546001600160a01b031690565b600080546001600160a01b03166106ca610a38565b6001600160a01b031614905090565b60025490565b6000826106ea6106b5565b8061076a5750336001600160a01b0316816001600160a01b031663570ca7356040518163ffffffff1660e01b815260040160206040518083038186803b15801561073357600080fd5b505afa158015610747573d6000803e3d6000fd5b505050506040513d602081101561075d57600080fd5b50516001600160a01b0316145b6107a55760405162461bcd60e51b8152600401808060200182810382526027815260200180610b036027913960400191505060405180910390fd5b6107ae84610839565b6107b757600080fd5b6107c1848461091b565b6107ca57600080fd5b5060019392505050565b6107dc6106b5565b61082d576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b61083681610a3c565b50565b6001600160a01b03811660009081526001602052604081205460ff161561085f57600080fd5b816001600160a01b031663420bb4b86040518163ffffffff1660e01b815260040160206040518083038186803b15801561089857600080fd5b505afa1580156108ac573d6000803e3d6000fd5b505050506040513d60208110156108c257600080fd5b50516108cd57600080fd5b506001600160a01b03166000818152600160208181526040808420805460ff191684179055600280548552600390925290922080546001600160a01b03191690931790925580548201905590565b6000816001600160a01b031663833a774f846040518263ffffffff1660e01b815260040180826001600160a01b03166001600160a01b03168152602001915050602060405180830381600087803b15801561097557600080fd5b505af1158015610989573d6000803e3d6000fd5b505050506040513d602081101561099f57600080fd5b50519392505050565b6040805163211276b360e11b81526001600160a01b0386811660048301526024820185905283151560448301529151600092861691634224ed6691606480830192602092919082900301818787803b158015610a0357600080fd5b505af1158015610a17573d6000803e3d6000fd5b505050506040513d6020811015610a2d57600080fd5b505195945050505050565b3390565b6001600160a01b038116610a815760405162461bcd60e51b8152600401808060200182810382526026815260200180610add6026913960400191505060405180910390fd5b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b039290921691909117905556fe4f776e61626c653a206e6577206f776e657220697320746865207a65726f206164647265737373656e646572206973206e656974686572206f70657261746f72206e6f72206f70657261746f72a265627a7a72315820c2d60bfe96bb4663b748cdbc2620c53b9a65806673ef97d3301d083a07c6ad7964736f6c634300050c0032"

// DeployRootChainRegistry deploys a new Ethereum contract, binding an instance of RootChainRegistry to it.
func DeployRootChainRegistry(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RootChainRegistry, error) {
	parsed, err := abi.JSON(strings.NewReader(RootChainRegistryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RootChainRegistryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RootChainRegistry{RootChainRegistryCaller: RootChainRegistryCaller{contract: contract}, RootChainRegistryTransactor: RootChainRegistryTransactor{contract: contract}, RootChainRegistryFilterer: RootChainRegistryFilterer{contract: contract}}, nil
}

// RootChainRegistry is an auto generated Go binding around an Ethereum contract.
type RootChainRegistry struct {
	RootChainRegistryCaller     // Read-only binding to the contract
	RootChainRegistryTransactor // Write-only binding to the contract
	RootChainRegistryFilterer   // Log filterer for contract events
}

// RootChainRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type RootChainRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RootChainRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RootChainRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RootChainRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RootChainRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RootChainRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RootChainRegistrySession struct {
	Contract     *RootChainRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// RootChainRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RootChainRegistryCallerSession struct {
	Contract *RootChainRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// RootChainRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RootChainRegistryTransactorSession struct {
	Contract     *RootChainRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// RootChainRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type RootChainRegistryRaw struct {
	Contract *RootChainRegistry // Generic contract binding to access the raw methods on
}

// RootChainRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RootChainRegistryCallerRaw struct {
	Contract *RootChainRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// RootChainRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RootChainRegistryTransactorRaw struct {
	Contract *RootChainRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRootChainRegistry creates a new instance of RootChainRegistry, bound to a specific deployed contract.
func NewRootChainRegistry(address common.Address, backend bind.ContractBackend) (*RootChainRegistry, error) {
	contract, err := bindRootChainRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RootChainRegistry{RootChainRegistryCaller: RootChainRegistryCaller{contract: contract}, RootChainRegistryTransactor: RootChainRegistryTransactor{contract: contract}, RootChainRegistryFilterer: RootChainRegistryFilterer{contract: contract}}, nil
}

// NewRootChainRegistryCaller creates a new read-only instance of RootChainRegistry, bound to a specific deployed contract.
func NewRootChainRegistryCaller(address common.Address, caller bind.ContractCaller) (*RootChainRegistryCaller, error) {
	contract, err := bindRootChainRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RootChainRegistryCaller{contract: contract}, nil
}

// NewRootChainRegistryTransactor creates a new write-only instance of RootChainRegistry, bound to a specific deployed contract.
func NewRootChainRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*RootChainRegistryTransactor, error) {
	contract, err := bindRootChainRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RootChainRegistryTransactor{contract: contract}, nil
}

// NewRootChainRegistryFilterer creates a new log filterer instance of RootChainRegistry, bound to a specific deployed contract.
func NewRootChainRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*RootChainRegistryFilterer, error) {
	contract, err := bindRootChainRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RootChainRegistryFilterer{contract: contract}, nil
}

// bindRootChainRegistry binds a generic wrapper to an already deployed contract.
func bindRootChainRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RootChainRegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RootChainRegistry *RootChainRegistryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RootChainRegistry.Contract.RootChainRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RootChainRegistry *RootChainRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RootChainRegistry.Contract.RootChainRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RootChainRegistry *RootChainRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RootChainRegistry.Contract.RootChainRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RootChainRegistry *RootChainRegistryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RootChainRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RootChainRegistry *RootChainRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RootChainRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RootChainRegistry *RootChainRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RootChainRegistry.Contract.contract.Transact(opts, method, params...)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_RootChainRegistry *RootChainRegistryCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _RootChainRegistry.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_RootChainRegistry *RootChainRegistrySession) IsOwner() (bool, error) {
	return _RootChainRegistry.Contract.IsOwner(&_RootChainRegistry.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_RootChainRegistry *RootChainRegistryCallerSession) IsOwner() (bool, error) {
	return _RootChainRegistry.Contract.IsOwner(&_RootChainRegistry.CallOpts)
}

// NumRootChains is a free data retrieval call binding the contract method 0xb2b604d0.
//
// Solidity: function numRootChains() constant returns(uint256)
func (_RootChainRegistry *RootChainRegistryCaller) NumRootChains(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChainRegistry.contract.Call(opts, out, "numRootChains")
	return *ret0, err
}

// NumRootChains is a free data retrieval call binding the contract method 0xb2b604d0.
//
// Solidity: function numRootChains() constant returns(uint256)
func (_RootChainRegistry *RootChainRegistrySession) NumRootChains() (*big.Int, error) {
	return _RootChainRegistry.Contract.NumRootChains(&_RootChainRegistry.CallOpts)
}

// NumRootChains is a free data retrieval call binding the contract method 0xb2b604d0.
//
// Solidity: function numRootChains() constant returns(uint256)
func (_RootChainRegistry *RootChainRegistryCallerSession) NumRootChains() (*big.Int, error) {
	return _RootChainRegistry.Contract.NumRootChains(&_RootChainRegistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_RootChainRegistry *RootChainRegistryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RootChainRegistry.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_RootChainRegistry *RootChainRegistrySession) Owner() (common.Address, error) {
	return _RootChainRegistry.Contract.Owner(&_RootChainRegistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_RootChainRegistry *RootChainRegistryCallerSession) Owner() (common.Address, error) {
	return _RootChainRegistry.Contract.Owner(&_RootChainRegistry.CallOpts)
}

// RootchainByIndex is a free data retrieval call binding the contract method 0x821f602c.
//
// Solidity: function rootchainByIndex(uint256 index) constant returns(address)
func (_RootChainRegistry *RootChainRegistryCaller) RootchainByIndex(opts *bind.CallOpts, index *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RootChainRegistry.contract.Call(opts, out, "rootchainByIndex", index)
	return *ret0, err
}

// RootchainByIndex is a free data retrieval call binding the contract method 0x821f602c.
//
// Solidity: function rootchainByIndex(uint256 index) constant returns(address)
func (_RootChainRegistry *RootChainRegistrySession) RootchainByIndex(index *big.Int) (common.Address, error) {
	return _RootChainRegistry.Contract.RootchainByIndex(&_RootChainRegistry.CallOpts, index)
}

// RootchainByIndex is a free data retrieval call binding the contract method 0x821f602c.
//
// Solidity: function rootchainByIndex(uint256 index) constant returns(address)
func (_RootChainRegistry *RootChainRegistryCallerSession) RootchainByIndex(index *big.Int) (common.Address, error) {
	return _RootChainRegistry.Contract.RootchainByIndex(&_RootChainRegistry.CallOpts, index)
}

// Rootchains is a free data retrieval call binding the contract method 0x02a15299.
//
// Solidity: function rootchains(address rootchain) constant returns(bool)
func (_RootChainRegistry *RootChainRegistryCaller) Rootchains(opts *bind.CallOpts, rootchain common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _RootChainRegistry.contract.Call(opts, out, "rootchains", rootchain)
	return *ret0, err
}

// Rootchains is a free data retrieval call binding the contract method 0x02a15299.
//
// Solidity: function rootchains(address rootchain) constant returns(bool)
func (_RootChainRegistry *RootChainRegistrySession) Rootchains(rootchain common.Address) (bool, error) {
	return _RootChainRegistry.Contract.Rootchains(&_RootChainRegistry.CallOpts, rootchain)
}

// Rootchains is a free data retrieval call binding the contract method 0x02a15299.
//
// Solidity: function rootchains(address rootchain) constant returns(bool)
func (_RootChainRegistry *RootChainRegistryCallerSession) Rootchains(rootchain common.Address) (bool, error) {
	return _RootChainRegistry.Contract.Rootchains(&_RootChainRegistry.CallOpts, rootchain)
}

// DeployCoinage is a paid mutator transaction binding the contract method 0x85108604.
//
// Solidity: function deployCoinage(address rootchain, address seigManager) returns(bool)
func (_RootChainRegistry *RootChainRegistryTransactor) DeployCoinage(opts *bind.TransactOpts, rootchain common.Address, seigManager common.Address) (*types.Transaction, error) {
	return _RootChainRegistry.contract.Transact(opts, "deployCoinage", rootchain, seigManager)
}

// DeployCoinage is a paid mutator transaction binding the contract method 0x85108604.
//
// Solidity: function deployCoinage(address rootchain, address seigManager) returns(bool)
func (_RootChainRegistry *RootChainRegistrySession) DeployCoinage(rootchain common.Address, seigManager common.Address) (*types.Transaction, error) {
	return _RootChainRegistry.Contract.DeployCoinage(&_RootChainRegistry.TransactOpts, rootchain, seigManager)
}

// DeployCoinage is a paid mutator transaction binding the contract method 0x85108604.
//
// Solidity: function deployCoinage(address rootchain, address seigManager) returns(bool)
func (_RootChainRegistry *RootChainRegistryTransactorSession) DeployCoinage(rootchain common.Address, seigManager common.Address) (*types.Transaction, error) {
	return _RootChainRegistry.Contract.DeployCoinage(&_RootChainRegistry.TransactOpts, rootchain, seigManager)
}

// Register is a paid mutator transaction binding the contract method 0x4420e486.
//
// Solidity: function register(address rootchain) returns(bool)
func (_RootChainRegistry *RootChainRegistryTransactor) Register(opts *bind.TransactOpts, rootchain common.Address) (*types.Transaction, error) {
	return _RootChainRegistry.contract.Transact(opts, "register", rootchain)
}

// Register is a paid mutator transaction binding the contract method 0x4420e486.
//
// Solidity: function register(address rootchain) returns(bool)
func (_RootChainRegistry *RootChainRegistrySession) Register(rootchain common.Address) (*types.Transaction, error) {
	return _RootChainRegistry.Contract.Register(&_RootChainRegistry.TransactOpts, rootchain)
}

// Register is a paid mutator transaction binding the contract method 0x4420e486.
//
// Solidity: function register(address rootchain) returns(bool)
func (_RootChainRegistry *RootChainRegistryTransactorSession) Register(rootchain common.Address) (*types.Transaction, error) {
	return _RootChainRegistry.Contract.Register(&_RootChainRegistry.TransactOpts, rootchain)
}

// RegisterAndDeployCoinage is a paid mutator transaction binding the contract method 0xbcb1a71e.
//
// Solidity: function registerAndDeployCoinage(address rootchain, address seigManager) returns(bool)
func (_RootChainRegistry *RootChainRegistryTransactor) RegisterAndDeployCoinage(opts *bind.TransactOpts, rootchain common.Address, seigManager common.Address) (*types.Transaction, error) {
	return _RootChainRegistry.contract.Transact(opts, "registerAndDeployCoinage", rootchain, seigManager)
}

// RegisterAndDeployCoinage is a paid mutator transaction binding the contract method 0xbcb1a71e.
//
// Solidity: function registerAndDeployCoinage(address rootchain, address seigManager) returns(bool)
func (_RootChainRegistry *RootChainRegistrySession) RegisterAndDeployCoinage(rootchain common.Address, seigManager common.Address) (*types.Transaction, error) {
	return _RootChainRegistry.Contract.RegisterAndDeployCoinage(&_RootChainRegistry.TransactOpts, rootchain, seigManager)
}

// RegisterAndDeployCoinage is a paid mutator transaction binding the contract method 0xbcb1a71e.
//
// Solidity: function registerAndDeployCoinage(address rootchain, address seigManager) returns(bool)
func (_RootChainRegistry *RootChainRegistryTransactorSession) RegisterAndDeployCoinage(rootchain common.Address, seigManager common.Address) (*types.Transaction, error) {
	return _RootChainRegistry.Contract.RegisterAndDeployCoinage(&_RootChainRegistry.TransactOpts, rootchain, seigManager)
}

// RegisterAndDeployCoinageAndSetCommissionRate is a paid mutator transaction binding the contract method 0x3eb2a66c.
//
// Solidity: function registerAndDeployCoinageAndSetCommissionRate(address rootchain, address seigManager, uint256 commissionRate, bool isCommissionRateNegative) returns(bool)
func (_RootChainRegistry *RootChainRegistryTransactor) RegisterAndDeployCoinageAndSetCommissionRate(opts *bind.TransactOpts, rootchain common.Address, seigManager common.Address, commissionRate *big.Int, isCommissionRateNegative bool) (*types.Transaction, error) {
	return _RootChainRegistry.contract.Transact(opts, "registerAndDeployCoinageAndSetCommissionRate", rootchain, seigManager, commissionRate, isCommissionRateNegative)
}

// RegisterAndDeployCoinageAndSetCommissionRate is a paid mutator transaction binding the contract method 0x3eb2a66c.
//
// Solidity: function registerAndDeployCoinageAndSetCommissionRate(address rootchain, address seigManager, uint256 commissionRate, bool isCommissionRateNegative) returns(bool)
func (_RootChainRegistry *RootChainRegistrySession) RegisterAndDeployCoinageAndSetCommissionRate(rootchain common.Address, seigManager common.Address, commissionRate *big.Int, isCommissionRateNegative bool) (*types.Transaction, error) {
	return _RootChainRegistry.Contract.RegisterAndDeployCoinageAndSetCommissionRate(&_RootChainRegistry.TransactOpts, rootchain, seigManager, commissionRate, isCommissionRateNegative)
}

// RegisterAndDeployCoinageAndSetCommissionRate is a paid mutator transaction binding the contract method 0x3eb2a66c.
//
// Solidity: function registerAndDeployCoinageAndSetCommissionRate(address rootchain, address seigManager, uint256 commissionRate, bool isCommissionRateNegative) returns(bool)
func (_RootChainRegistry *RootChainRegistryTransactorSession) RegisterAndDeployCoinageAndSetCommissionRate(rootchain common.Address, seigManager common.Address, commissionRate *big.Int, isCommissionRateNegative bool) (*types.Transaction, error) {
	return _RootChainRegistry.Contract.RegisterAndDeployCoinageAndSetCommissionRate(&_RootChainRegistry.TransactOpts, rootchain, seigManager, commissionRate, isCommissionRateNegative)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RootChainRegistry *RootChainRegistryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RootChainRegistry.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RootChainRegistry *RootChainRegistrySession) RenounceOwnership() (*types.Transaction, error) {
	return _RootChainRegistry.Contract.RenounceOwnership(&_RootChainRegistry.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RootChainRegistry *RootChainRegistryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _RootChainRegistry.Contract.RenounceOwnership(&_RootChainRegistry.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RootChainRegistry *RootChainRegistryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _RootChainRegistry.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RootChainRegistry *RootChainRegistrySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _RootChainRegistry.Contract.TransferOwnership(&_RootChainRegistry.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RootChainRegistry *RootChainRegistryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _RootChainRegistry.Contract.TransferOwnership(&_RootChainRegistry.TransactOpts, newOwner)
}

// Unregister is a paid mutator transaction binding the contract method 0x2ec2c246.
//
// Solidity: function unregister(address rootchain) returns(bool)
func (_RootChainRegistry *RootChainRegistryTransactor) Unregister(opts *bind.TransactOpts, rootchain common.Address) (*types.Transaction, error) {
	return _RootChainRegistry.contract.Transact(opts, "unregister", rootchain)
}

// Unregister is a paid mutator transaction binding the contract method 0x2ec2c246.
//
// Solidity: function unregister(address rootchain) returns(bool)
func (_RootChainRegistry *RootChainRegistrySession) Unregister(rootchain common.Address) (*types.Transaction, error) {
	return _RootChainRegistry.Contract.Unregister(&_RootChainRegistry.TransactOpts, rootchain)
}

// Unregister is a paid mutator transaction binding the contract method 0x2ec2c246.
//
// Solidity: function unregister(address rootchain) returns(bool)
func (_RootChainRegistry *RootChainRegistryTransactorSession) Unregister(rootchain common.Address) (*types.Transaction, error) {
	return _RootChainRegistry.Contract.Unregister(&_RootChainRegistry.TransactOpts, rootchain)
}

// RootChainRegistryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the RootChainRegistry contract.
type RootChainRegistryOwnershipTransferredIterator struct {
	Event *RootChainRegistryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *RootChainRegistryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RootChainRegistryOwnershipTransferred)
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
		it.Event = new(RootChainRegistryOwnershipTransferred)
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
func (it *RootChainRegistryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RootChainRegistryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RootChainRegistryOwnershipTransferred represents a OwnershipTransferred event raised by the RootChainRegistry contract.
type RootChainRegistryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_RootChainRegistry *RootChainRegistryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*RootChainRegistryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RootChainRegistry.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RootChainRegistryOwnershipTransferredIterator{contract: _RootChainRegistry.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_RootChainRegistry *RootChainRegistryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RootChainRegistryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RootChainRegistry.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RootChainRegistryOwnershipTransferred)
				if err := _RootChainRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_RootChainRegistry *RootChainRegistryFilterer) ParseOwnershipTransferred(log types.Log) (*RootChainRegistryOwnershipTransferred, error) {
	event := new(RootChainRegistryOwnershipTransferred)
	if err := _RootChainRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RootChainRegistryIABI is the input ABI used to generate the binding from.
const RootChainRegistryIABI = "[{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"rootchain\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"seigManager\",\"type\":\"address\"}],\"name\":\"deployCoinage\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"numRootChains\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"rootchain\",\"type\":\"address\"}],\"name\":\"register\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"rootchain\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"seigManager\",\"type\":\"address\"}],\"name\":\"registerAndDeployCoinage\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"rootchainByIndex\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"rootchain\",\"type\":\"address\"}],\"name\":\"rootchains\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"rootchain\",\"type\":\"address\"}],\"name\":\"unregister\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// RootChainRegistryIFuncSigs maps the 4-byte function signature to its string representation.
var RootChainRegistryIFuncSigs = map[string]string{
	"85108604": "deployCoinage(address,address)",
	"b2b604d0": "numRootChains()",
	"4420e486": "register(address)",
	"bcb1a71e": "registerAndDeployCoinage(address,address)",
	"821f602c": "rootchainByIndex(uint256)",
	"02a15299": "rootchains(address)",
	"2ec2c246": "unregister(address)",
}

// RootChainRegistryI is an auto generated Go binding around an Ethereum contract.
type RootChainRegistryI struct {
	RootChainRegistryICaller     // Read-only binding to the contract
	RootChainRegistryITransactor // Write-only binding to the contract
	RootChainRegistryIFilterer   // Log filterer for contract events
}

// RootChainRegistryICaller is an auto generated read-only Go binding around an Ethereum contract.
type RootChainRegistryICaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RootChainRegistryITransactor is an auto generated write-only Go binding around an Ethereum contract.
type RootChainRegistryITransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RootChainRegistryIFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RootChainRegistryIFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RootChainRegistryISession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RootChainRegistryISession struct {
	Contract     *RootChainRegistryI // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// RootChainRegistryICallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RootChainRegistryICallerSession struct {
	Contract *RootChainRegistryICaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// RootChainRegistryITransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RootChainRegistryITransactorSession struct {
	Contract     *RootChainRegistryITransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// RootChainRegistryIRaw is an auto generated low-level Go binding around an Ethereum contract.
type RootChainRegistryIRaw struct {
	Contract *RootChainRegistryI // Generic contract binding to access the raw methods on
}

// RootChainRegistryICallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RootChainRegistryICallerRaw struct {
	Contract *RootChainRegistryICaller // Generic read-only contract binding to access the raw methods on
}

// RootChainRegistryITransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RootChainRegistryITransactorRaw struct {
	Contract *RootChainRegistryITransactor // Generic write-only contract binding to access the raw methods on
}

// NewRootChainRegistryI creates a new instance of RootChainRegistryI, bound to a specific deployed contract.
func NewRootChainRegistryI(address common.Address, backend bind.ContractBackend) (*RootChainRegistryI, error) {
	contract, err := bindRootChainRegistryI(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RootChainRegistryI{RootChainRegistryICaller: RootChainRegistryICaller{contract: contract}, RootChainRegistryITransactor: RootChainRegistryITransactor{contract: contract}, RootChainRegistryIFilterer: RootChainRegistryIFilterer{contract: contract}}, nil
}

// NewRootChainRegistryICaller creates a new read-only instance of RootChainRegistryI, bound to a specific deployed contract.
func NewRootChainRegistryICaller(address common.Address, caller bind.ContractCaller) (*RootChainRegistryICaller, error) {
	contract, err := bindRootChainRegistryI(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RootChainRegistryICaller{contract: contract}, nil
}

// NewRootChainRegistryITransactor creates a new write-only instance of RootChainRegistryI, bound to a specific deployed contract.
func NewRootChainRegistryITransactor(address common.Address, transactor bind.ContractTransactor) (*RootChainRegistryITransactor, error) {
	contract, err := bindRootChainRegistryI(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RootChainRegistryITransactor{contract: contract}, nil
}

// NewRootChainRegistryIFilterer creates a new log filterer instance of RootChainRegistryI, bound to a specific deployed contract.
func NewRootChainRegistryIFilterer(address common.Address, filterer bind.ContractFilterer) (*RootChainRegistryIFilterer, error) {
	contract, err := bindRootChainRegistryI(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RootChainRegistryIFilterer{contract: contract}, nil
}

// bindRootChainRegistryI binds a generic wrapper to an already deployed contract.
func bindRootChainRegistryI(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RootChainRegistryIABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RootChainRegistryI *RootChainRegistryIRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RootChainRegistryI.Contract.RootChainRegistryICaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RootChainRegistryI *RootChainRegistryIRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RootChainRegistryI.Contract.RootChainRegistryITransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RootChainRegistryI *RootChainRegistryIRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RootChainRegistryI.Contract.RootChainRegistryITransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RootChainRegistryI *RootChainRegistryICallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RootChainRegistryI.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RootChainRegistryI *RootChainRegistryITransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RootChainRegistryI.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RootChainRegistryI *RootChainRegistryITransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RootChainRegistryI.Contract.contract.Transact(opts, method, params...)
}

// NumRootChains is a free data retrieval call binding the contract method 0xb2b604d0.
//
// Solidity: function numRootChains() constant returns(uint256)
func (_RootChainRegistryI *RootChainRegistryICaller) NumRootChains(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChainRegistryI.contract.Call(opts, out, "numRootChains")
	return *ret0, err
}

// NumRootChains is a free data retrieval call binding the contract method 0xb2b604d0.
//
// Solidity: function numRootChains() constant returns(uint256)
func (_RootChainRegistryI *RootChainRegistryISession) NumRootChains() (*big.Int, error) {
	return _RootChainRegistryI.Contract.NumRootChains(&_RootChainRegistryI.CallOpts)
}

// NumRootChains is a free data retrieval call binding the contract method 0xb2b604d0.
//
// Solidity: function numRootChains() constant returns(uint256)
func (_RootChainRegistryI *RootChainRegistryICallerSession) NumRootChains() (*big.Int, error) {
	return _RootChainRegistryI.Contract.NumRootChains(&_RootChainRegistryI.CallOpts)
}

// RootchainByIndex is a free data retrieval call binding the contract method 0x821f602c.
//
// Solidity: function rootchainByIndex(uint256 index) constant returns(address)
func (_RootChainRegistryI *RootChainRegistryICaller) RootchainByIndex(opts *bind.CallOpts, index *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RootChainRegistryI.contract.Call(opts, out, "rootchainByIndex", index)
	return *ret0, err
}

// RootchainByIndex is a free data retrieval call binding the contract method 0x821f602c.
//
// Solidity: function rootchainByIndex(uint256 index) constant returns(address)
func (_RootChainRegistryI *RootChainRegistryISession) RootchainByIndex(index *big.Int) (common.Address, error) {
	return _RootChainRegistryI.Contract.RootchainByIndex(&_RootChainRegistryI.CallOpts, index)
}

// RootchainByIndex is a free data retrieval call binding the contract method 0x821f602c.
//
// Solidity: function rootchainByIndex(uint256 index) constant returns(address)
func (_RootChainRegistryI *RootChainRegistryICallerSession) RootchainByIndex(index *big.Int) (common.Address, error) {
	return _RootChainRegistryI.Contract.RootchainByIndex(&_RootChainRegistryI.CallOpts, index)
}

// Rootchains is a free data retrieval call binding the contract method 0x02a15299.
//
// Solidity: function rootchains(address rootchain) constant returns(bool)
func (_RootChainRegistryI *RootChainRegistryICaller) Rootchains(opts *bind.CallOpts, rootchain common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _RootChainRegistryI.contract.Call(opts, out, "rootchains", rootchain)
	return *ret0, err
}

// Rootchains is a free data retrieval call binding the contract method 0x02a15299.
//
// Solidity: function rootchains(address rootchain) constant returns(bool)
func (_RootChainRegistryI *RootChainRegistryISession) Rootchains(rootchain common.Address) (bool, error) {
	return _RootChainRegistryI.Contract.Rootchains(&_RootChainRegistryI.CallOpts, rootchain)
}

// Rootchains is a free data retrieval call binding the contract method 0x02a15299.
//
// Solidity: function rootchains(address rootchain) constant returns(bool)
func (_RootChainRegistryI *RootChainRegistryICallerSession) Rootchains(rootchain common.Address) (bool, error) {
	return _RootChainRegistryI.Contract.Rootchains(&_RootChainRegistryI.CallOpts, rootchain)
}

// DeployCoinage is a paid mutator transaction binding the contract method 0x85108604.
//
// Solidity: function deployCoinage(address rootchain, address seigManager) returns(bool)
func (_RootChainRegistryI *RootChainRegistryITransactor) DeployCoinage(opts *bind.TransactOpts, rootchain common.Address, seigManager common.Address) (*types.Transaction, error) {
	return _RootChainRegistryI.contract.Transact(opts, "deployCoinage", rootchain, seigManager)
}

// DeployCoinage is a paid mutator transaction binding the contract method 0x85108604.
//
// Solidity: function deployCoinage(address rootchain, address seigManager) returns(bool)
func (_RootChainRegistryI *RootChainRegistryISession) DeployCoinage(rootchain common.Address, seigManager common.Address) (*types.Transaction, error) {
	return _RootChainRegistryI.Contract.DeployCoinage(&_RootChainRegistryI.TransactOpts, rootchain, seigManager)
}

// DeployCoinage is a paid mutator transaction binding the contract method 0x85108604.
//
// Solidity: function deployCoinage(address rootchain, address seigManager) returns(bool)
func (_RootChainRegistryI *RootChainRegistryITransactorSession) DeployCoinage(rootchain common.Address, seigManager common.Address) (*types.Transaction, error) {
	return _RootChainRegistryI.Contract.DeployCoinage(&_RootChainRegistryI.TransactOpts, rootchain, seigManager)
}

// Register is a paid mutator transaction binding the contract method 0x4420e486.
//
// Solidity: function register(address rootchain) returns(bool)
func (_RootChainRegistryI *RootChainRegistryITransactor) Register(opts *bind.TransactOpts, rootchain common.Address) (*types.Transaction, error) {
	return _RootChainRegistryI.contract.Transact(opts, "register", rootchain)
}

// Register is a paid mutator transaction binding the contract method 0x4420e486.
//
// Solidity: function register(address rootchain) returns(bool)
func (_RootChainRegistryI *RootChainRegistryISession) Register(rootchain common.Address) (*types.Transaction, error) {
	return _RootChainRegistryI.Contract.Register(&_RootChainRegistryI.TransactOpts, rootchain)
}

// Register is a paid mutator transaction binding the contract method 0x4420e486.
//
// Solidity: function register(address rootchain) returns(bool)
func (_RootChainRegistryI *RootChainRegistryITransactorSession) Register(rootchain common.Address) (*types.Transaction, error) {
	return _RootChainRegistryI.Contract.Register(&_RootChainRegistryI.TransactOpts, rootchain)
}

// RegisterAndDeployCoinage is a paid mutator transaction binding the contract method 0xbcb1a71e.
//
// Solidity: function registerAndDeployCoinage(address rootchain, address seigManager) returns(bool)
func (_RootChainRegistryI *RootChainRegistryITransactor) RegisterAndDeployCoinage(opts *bind.TransactOpts, rootchain common.Address, seigManager common.Address) (*types.Transaction, error) {
	return _RootChainRegistryI.contract.Transact(opts, "registerAndDeployCoinage", rootchain, seigManager)
}

// RegisterAndDeployCoinage is a paid mutator transaction binding the contract method 0xbcb1a71e.
//
// Solidity: function registerAndDeployCoinage(address rootchain, address seigManager) returns(bool)
func (_RootChainRegistryI *RootChainRegistryISession) RegisterAndDeployCoinage(rootchain common.Address, seigManager common.Address) (*types.Transaction, error) {
	return _RootChainRegistryI.Contract.RegisterAndDeployCoinage(&_RootChainRegistryI.TransactOpts, rootchain, seigManager)
}

// RegisterAndDeployCoinage is a paid mutator transaction binding the contract method 0xbcb1a71e.
//
// Solidity: function registerAndDeployCoinage(address rootchain, address seigManager) returns(bool)
func (_RootChainRegistryI *RootChainRegistryITransactorSession) RegisterAndDeployCoinage(rootchain common.Address, seigManager common.Address) (*types.Transaction, error) {
	return _RootChainRegistryI.Contract.RegisterAndDeployCoinage(&_RootChainRegistryI.TransactOpts, rootchain, seigManager)
}

// Unregister is a paid mutator transaction binding the contract method 0x2ec2c246.
//
// Solidity: function unregister(address rootchain) returns(bool)
func (_RootChainRegistryI *RootChainRegistryITransactor) Unregister(opts *bind.TransactOpts, rootchain common.Address) (*types.Transaction, error) {
	return _RootChainRegistryI.contract.Transact(opts, "unregister", rootchain)
}

// Unregister is a paid mutator transaction binding the contract method 0x2ec2c246.
//
// Solidity: function unregister(address rootchain) returns(bool)
func (_RootChainRegistryI *RootChainRegistryISession) Unregister(rootchain common.Address) (*types.Transaction, error) {
	return _RootChainRegistryI.Contract.Unregister(&_RootChainRegistryI.TransactOpts, rootchain)
}

// Unregister is a paid mutator transaction binding the contract method 0x2ec2c246.
//
// Solidity: function unregister(address rootchain) returns(bool)
func (_RootChainRegistryI *RootChainRegistryITransactorSession) Unregister(rootchain common.Address) (*types.Transaction, error) {
	return _RootChainRegistryI.Contract.Unregister(&_RootChainRegistryI.TransactOpts, rootchain)
}

// SeigManagerIABI is the input ABI used to generate the binding from.
const SeigManagerIABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"DEFAULT_FACTOR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"rootchain\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"additionalTotBurnAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totAmount\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"rootchain\",\"type\":\"address\"}],\"name\":\"coinages\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"rootchain\",\"type\":\"address\"}],\"name\":\"commissionRates\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"rootchain\",\"type\":\"address\"}],\"name\":\"deployCoinage\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"depositManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"rootchain\",\"type\":\"address\"}],\"name\":\"lastCommitBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastSeigBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"onCommit\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"rootchain\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"onDeposit\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"onTransfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"rootchain\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"onWithdraw\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pausedBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"powerton\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"registry\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"seigPerBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"rootchain\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"commission\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isCommissionRateNegative\",\"type\":\"bool\"}],\"name\":\"setCommissionRate\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"rootchain\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"stakeOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ton\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tot\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"rootchain\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"uncomittedStakeOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"unpausedBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"wton\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// SeigManagerIFuncSigs maps the 4-byte function signature to its string representation.
var SeigManagerIFuncSigs = map[string]string{
	"8bf91dc4": "DEFAULT_FACTOR()",
	"48c8577e": "additionalTotBurnAmount(address,address,uint256)",
	"4c063c19": "coinages(address)",
	"7b056c1b": "commissionRates(address)",
	"833a774f": "deployCoinage(address)",
	"6c7ac9d8": "depositManager()",
	"c59f1046": "lastCommitBlock(address)",
	"f35c89e8": "lastSeigBlock()",
	"359c4d59": "onCommit()",
	"412c6d50": "onDeposit(address,address,uint256)",
	"4a393149": "onTransfer(address,address,uint256)",
	"f850ffaa": "onWithdraw(address,address,uint256)",
	"32053c99": "pausedBlock()",
	"3e832e1d": "powerton()",
	"7b103999": "registry()",
	"5f40a349": "seigPerBlock()",
	"4224ed66": "setCommissionRate(address,uint256,bool)",
	"ce4cb876": "stakeOf(address,address)",
	"cc48b947": "ton()",
	"a16d6aa7": "tot()",
	"fa9789c8": "uncomittedStakeOf(address,address)",
	"1cc47890": "unpausedBlock()",
	"8d62d949": "wton()",
}

// SeigManagerI is an auto generated Go binding around an Ethereum contract.
type SeigManagerI struct {
	SeigManagerICaller     // Read-only binding to the contract
	SeigManagerITransactor // Write-only binding to the contract
	SeigManagerIFilterer   // Log filterer for contract events
}

// SeigManagerICaller is an auto generated read-only Go binding around an Ethereum contract.
type SeigManagerICaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SeigManagerITransactor is an auto generated write-only Go binding around an Ethereum contract.
type SeigManagerITransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SeigManagerIFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SeigManagerIFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SeigManagerISession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SeigManagerISession struct {
	Contract     *SeigManagerI     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SeigManagerICallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SeigManagerICallerSession struct {
	Contract *SeigManagerICaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// SeigManagerITransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SeigManagerITransactorSession struct {
	Contract     *SeigManagerITransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// SeigManagerIRaw is an auto generated low-level Go binding around an Ethereum contract.
type SeigManagerIRaw struct {
	Contract *SeigManagerI // Generic contract binding to access the raw methods on
}

// SeigManagerICallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SeigManagerICallerRaw struct {
	Contract *SeigManagerICaller // Generic read-only contract binding to access the raw methods on
}

// SeigManagerITransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SeigManagerITransactorRaw struct {
	Contract *SeigManagerITransactor // Generic write-only contract binding to access the raw methods on
}

// NewSeigManagerI creates a new instance of SeigManagerI, bound to a specific deployed contract.
func NewSeigManagerI(address common.Address, backend bind.ContractBackend) (*SeigManagerI, error) {
	contract, err := bindSeigManagerI(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SeigManagerI{SeigManagerICaller: SeigManagerICaller{contract: contract}, SeigManagerITransactor: SeigManagerITransactor{contract: contract}, SeigManagerIFilterer: SeigManagerIFilterer{contract: contract}}, nil
}

// NewSeigManagerICaller creates a new read-only instance of SeigManagerI, bound to a specific deployed contract.
func NewSeigManagerICaller(address common.Address, caller bind.ContractCaller) (*SeigManagerICaller, error) {
	contract, err := bindSeigManagerI(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SeigManagerICaller{contract: contract}, nil
}

// NewSeigManagerITransactor creates a new write-only instance of SeigManagerI, bound to a specific deployed contract.
func NewSeigManagerITransactor(address common.Address, transactor bind.ContractTransactor) (*SeigManagerITransactor, error) {
	contract, err := bindSeigManagerI(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SeigManagerITransactor{contract: contract}, nil
}

// NewSeigManagerIFilterer creates a new log filterer instance of SeigManagerI, bound to a specific deployed contract.
func NewSeigManagerIFilterer(address common.Address, filterer bind.ContractFilterer) (*SeigManagerIFilterer, error) {
	contract, err := bindSeigManagerI(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SeigManagerIFilterer{contract: contract}, nil
}

// bindSeigManagerI binds a generic wrapper to an already deployed contract.
func bindSeigManagerI(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SeigManagerIABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SeigManagerI *SeigManagerIRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SeigManagerI.Contract.SeigManagerICaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SeigManagerI *SeigManagerIRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SeigManagerI.Contract.SeigManagerITransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SeigManagerI *SeigManagerIRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SeigManagerI.Contract.SeigManagerITransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SeigManagerI *SeigManagerICallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SeigManagerI.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SeigManagerI *SeigManagerITransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SeigManagerI.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SeigManagerI *SeigManagerITransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SeigManagerI.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTFACTOR is a free data retrieval call binding the contract method 0x8bf91dc4.
//
// Solidity: function DEFAULT_FACTOR() constant returns(uint256)
func (_SeigManagerI *SeigManagerICaller) DEFAULTFACTOR(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SeigManagerI.contract.Call(opts, out, "DEFAULT_FACTOR")
	return *ret0, err
}

// DEFAULTFACTOR is a free data retrieval call binding the contract method 0x8bf91dc4.
//
// Solidity: function DEFAULT_FACTOR() constant returns(uint256)
func (_SeigManagerI *SeigManagerISession) DEFAULTFACTOR() (*big.Int, error) {
	return _SeigManagerI.Contract.DEFAULTFACTOR(&_SeigManagerI.CallOpts)
}

// DEFAULTFACTOR is a free data retrieval call binding the contract method 0x8bf91dc4.
//
// Solidity: function DEFAULT_FACTOR() constant returns(uint256)
func (_SeigManagerI *SeigManagerICallerSession) DEFAULTFACTOR() (*big.Int, error) {
	return _SeigManagerI.Contract.DEFAULTFACTOR(&_SeigManagerI.CallOpts)
}

// AdditionalTotBurnAmount is a free data retrieval call binding the contract method 0x48c8577e.
//
// Solidity: function additionalTotBurnAmount(address rootchain, address account, uint256 amount) constant returns(uint256 totAmount)
func (_SeigManagerI *SeigManagerICaller) AdditionalTotBurnAmount(opts *bind.CallOpts, rootchain common.Address, account common.Address, amount *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SeigManagerI.contract.Call(opts, out, "additionalTotBurnAmount", rootchain, account, amount)
	return *ret0, err
}

// AdditionalTotBurnAmount is a free data retrieval call binding the contract method 0x48c8577e.
//
// Solidity: function additionalTotBurnAmount(address rootchain, address account, uint256 amount) constant returns(uint256 totAmount)
func (_SeigManagerI *SeigManagerISession) AdditionalTotBurnAmount(rootchain common.Address, account common.Address, amount *big.Int) (*big.Int, error) {
	return _SeigManagerI.Contract.AdditionalTotBurnAmount(&_SeigManagerI.CallOpts, rootchain, account, amount)
}

// AdditionalTotBurnAmount is a free data retrieval call binding the contract method 0x48c8577e.
//
// Solidity: function additionalTotBurnAmount(address rootchain, address account, uint256 amount) constant returns(uint256 totAmount)
func (_SeigManagerI *SeigManagerICallerSession) AdditionalTotBurnAmount(rootchain common.Address, account common.Address, amount *big.Int) (*big.Int, error) {
	return _SeigManagerI.Contract.AdditionalTotBurnAmount(&_SeigManagerI.CallOpts, rootchain, account, amount)
}

// Coinages is a free data retrieval call binding the contract method 0x4c063c19.
//
// Solidity: function coinages(address rootchain) constant returns(address)
func (_SeigManagerI *SeigManagerICaller) Coinages(opts *bind.CallOpts, rootchain common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SeigManagerI.contract.Call(opts, out, "coinages", rootchain)
	return *ret0, err
}

// Coinages is a free data retrieval call binding the contract method 0x4c063c19.
//
// Solidity: function coinages(address rootchain) constant returns(address)
func (_SeigManagerI *SeigManagerISession) Coinages(rootchain common.Address) (common.Address, error) {
	return _SeigManagerI.Contract.Coinages(&_SeigManagerI.CallOpts, rootchain)
}

// Coinages is a free data retrieval call binding the contract method 0x4c063c19.
//
// Solidity: function coinages(address rootchain) constant returns(address)
func (_SeigManagerI *SeigManagerICallerSession) Coinages(rootchain common.Address) (common.Address, error) {
	return _SeigManagerI.Contract.Coinages(&_SeigManagerI.CallOpts, rootchain)
}

// CommissionRates is a free data retrieval call binding the contract method 0x7b056c1b.
//
// Solidity: function commissionRates(address rootchain) constant returns(uint256)
func (_SeigManagerI *SeigManagerICaller) CommissionRates(opts *bind.CallOpts, rootchain common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SeigManagerI.contract.Call(opts, out, "commissionRates", rootchain)
	return *ret0, err
}

// CommissionRates is a free data retrieval call binding the contract method 0x7b056c1b.
//
// Solidity: function commissionRates(address rootchain) constant returns(uint256)
func (_SeigManagerI *SeigManagerISession) CommissionRates(rootchain common.Address) (*big.Int, error) {
	return _SeigManagerI.Contract.CommissionRates(&_SeigManagerI.CallOpts, rootchain)
}

// CommissionRates is a free data retrieval call binding the contract method 0x7b056c1b.
//
// Solidity: function commissionRates(address rootchain) constant returns(uint256)
func (_SeigManagerI *SeigManagerICallerSession) CommissionRates(rootchain common.Address) (*big.Int, error) {
	return _SeigManagerI.Contract.CommissionRates(&_SeigManagerI.CallOpts, rootchain)
}

// DepositManager is a free data retrieval call binding the contract method 0x6c7ac9d8.
//
// Solidity: function depositManager() constant returns(address)
func (_SeigManagerI *SeigManagerICaller) DepositManager(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SeigManagerI.contract.Call(opts, out, "depositManager")
	return *ret0, err
}

// DepositManager is a free data retrieval call binding the contract method 0x6c7ac9d8.
//
// Solidity: function depositManager() constant returns(address)
func (_SeigManagerI *SeigManagerISession) DepositManager() (common.Address, error) {
	return _SeigManagerI.Contract.DepositManager(&_SeigManagerI.CallOpts)
}

// DepositManager is a free data retrieval call binding the contract method 0x6c7ac9d8.
//
// Solidity: function depositManager() constant returns(address)
func (_SeigManagerI *SeigManagerICallerSession) DepositManager() (common.Address, error) {
	return _SeigManagerI.Contract.DepositManager(&_SeigManagerI.CallOpts)
}

// LastCommitBlock is a free data retrieval call binding the contract method 0xc59f1046.
//
// Solidity: function lastCommitBlock(address rootchain) constant returns(uint256)
func (_SeigManagerI *SeigManagerICaller) LastCommitBlock(opts *bind.CallOpts, rootchain common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SeigManagerI.contract.Call(opts, out, "lastCommitBlock", rootchain)
	return *ret0, err
}

// LastCommitBlock is a free data retrieval call binding the contract method 0xc59f1046.
//
// Solidity: function lastCommitBlock(address rootchain) constant returns(uint256)
func (_SeigManagerI *SeigManagerISession) LastCommitBlock(rootchain common.Address) (*big.Int, error) {
	return _SeigManagerI.Contract.LastCommitBlock(&_SeigManagerI.CallOpts, rootchain)
}

// LastCommitBlock is a free data retrieval call binding the contract method 0xc59f1046.
//
// Solidity: function lastCommitBlock(address rootchain) constant returns(uint256)
func (_SeigManagerI *SeigManagerICallerSession) LastCommitBlock(rootchain common.Address) (*big.Int, error) {
	return _SeigManagerI.Contract.LastCommitBlock(&_SeigManagerI.CallOpts, rootchain)
}

// LastSeigBlock is a free data retrieval call binding the contract method 0xf35c89e8.
//
// Solidity: function lastSeigBlock() constant returns(uint256)
func (_SeigManagerI *SeigManagerICaller) LastSeigBlock(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SeigManagerI.contract.Call(opts, out, "lastSeigBlock")
	return *ret0, err
}

// LastSeigBlock is a free data retrieval call binding the contract method 0xf35c89e8.
//
// Solidity: function lastSeigBlock() constant returns(uint256)
func (_SeigManagerI *SeigManagerISession) LastSeigBlock() (*big.Int, error) {
	return _SeigManagerI.Contract.LastSeigBlock(&_SeigManagerI.CallOpts)
}

// LastSeigBlock is a free data retrieval call binding the contract method 0xf35c89e8.
//
// Solidity: function lastSeigBlock() constant returns(uint256)
func (_SeigManagerI *SeigManagerICallerSession) LastSeigBlock() (*big.Int, error) {
	return _SeigManagerI.Contract.LastSeigBlock(&_SeigManagerI.CallOpts)
}

// PausedBlock is a free data retrieval call binding the contract method 0x32053c99.
//
// Solidity: function pausedBlock() constant returns(uint256)
func (_SeigManagerI *SeigManagerICaller) PausedBlock(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SeigManagerI.contract.Call(opts, out, "pausedBlock")
	return *ret0, err
}

// PausedBlock is a free data retrieval call binding the contract method 0x32053c99.
//
// Solidity: function pausedBlock() constant returns(uint256)
func (_SeigManagerI *SeigManagerISession) PausedBlock() (*big.Int, error) {
	return _SeigManagerI.Contract.PausedBlock(&_SeigManagerI.CallOpts)
}

// PausedBlock is a free data retrieval call binding the contract method 0x32053c99.
//
// Solidity: function pausedBlock() constant returns(uint256)
func (_SeigManagerI *SeigManagerICallerSession) PausedBlock() (*big.Int, error) {
	return _SeigManagerI.Contract.PausedBlock(&_SeigManagerI.CallOpts)
}

// Powerton is a free data retrieval call binding the contract method 0x3e832e1d.
//
// Solidity: function powerton() constant returns(address)
func (_SeigManagerI *SeigManagerICaller) Powerton(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SeigManagerI.contract.Call(opts, out, "powerton")
	return *ret0, err
}

// Powerton is a free data retrieval call binding the contract method 0x3e832e1d.
//
// Solidity: function powerton() constant returns(address)
func (_SeigManagerI *SeigManagerISession) Powerton() (common.Address, error) {
	return _SeigManagerI.Contract.Powerton(&_SeigManagerI.CallOpts)
}

// Powerton is a free data retrieval call binding the contract method 0x3e832e1d.
//
// Solidity: function powerton() constant returns(address)
func (_SeigManagerI *SeigManagerICallerSession) Powerton() (common.Address, error) {
	return _SeigManagerI.Contract.Powerton(&_SeigManagerI.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() constant returns(address)
func (_SeigManagerI *SeigManagerICaller) Registry(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SeigManagerI.contract.Call(opts, out, "registry")
	return *ret0, err
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() constant returns(address)
func (_SeigManagerI *SeigManagerISession) Registry() (common.Address, error) {
	return _SeigManagerI.Contract.Registry(&_SeigManagerI.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() constant returns(address)
func (_SeigManagerI *SeigManagerICallerSession) Registry() (common.Address, error) {
	return _SeigManagerI.Contract.Registry(&_SeigManagerI.CallOpts)
}

// SeigPerBlock is a free data retrieval call binding the contract method 0x5f40a349.
//
// Solidity: function seigPerBlock() constant returns(uint256)
func (_SeigManagerI *SeigManagerICaller) SeigPerBlock(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SeigManagerI.contract.Call(opts, out, "seigPerBlock")
	return *ret0, err
}

// SeigPerBlock is a free data retrieval call binding the contract method 0x5f40a349.
//
// Solidity: function seigPerBlock() constant returns(uint256)
func (_SeigManagerI *SeigManagerISession) SeigPerBlock() (*big.Int, error) {
	return _SeigManagerI.Contract.SeigPerBlock(&_SeigManagerI.CallOpts)
}

// SeigPerBlock is a free data retrieval call binding the contract method 0x5f40a349.
//
// Solidity: function seigPerBlock() constant returns(uint256)
func (_SeigManagerI *SeigManagerICallerSession) SeigPerBlock() (*big.Int, error) {
	return _SeigManagerI.Contract.SeigPerBlock(&_SeigManagerI.CallOpts)
}

// StakeOf is a free data retrieval call binding the contract method 0xce4cb876.
//
// Solidity: function stakeOf(address rootchain, address account) constant returns(uint256)
func (_SeigManagerI *SeigManagerICaller) StakeOf(opts *bind.CallOpts, rootchain common.Address, account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SeigManagerI.contract.Call(opts, out, "stakeOf", rootchain, account)
	return *ret0, err
}

// StakeOf is a free data retrieval call binding the contract method 0xce4cb876.
//
// Solidity: function stakeOf(address rootchain, address account) constant returns(uint256)
func (_SeigManagerI *SeigManagerISession) StakeOf(rootchain common.Address, account common.Address) (*big.Int, error) {
	return _SeigManagerI.Contract.StakeOf(&_SeigManagerI.CallOpts, rootchain, account)
}

// StakeOf is a free data retrieval call binding the contract method 0xce4cb876.
//
// Solidity: function stakeOf(address rootchain, address account) constant returns(uint256)
func (_SeigManagerI *SeigManagerICallerSession) StakeOf(rootchain common.Address, account common.Address) (*big.Int, error) {
	return _SeigManagerI.Contract.StakeOf(&_SeigManagerI.CallOpts, rootchain, account)
}

// Ton is a free data retrieval call binding the contract method 0xcc48b947.
//
// Solidity: function ton() constant returns(address)
func (_SeigManagerI *SeigManagerICaller) Ton(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SeigManagerI.contract.Call(opts, out, "ton")
	return *ret0, err
}

// Ton is a free data retrieval call binding the contract method 0xcc48b947.
//
// Solidity: function ton() constant returns(address)
func (_SeigManagerI *SeigManagerISession) Ton() (common.Address, error) {
	return _SeigManagerI.Contract.Ton(&_SeigManagerI.CallOpts)
}

// Ton is a free data retrieval call binding the contract method 0xcc48b947.
//
// Solidity: function ton() constant returns(address)
func (_SeigManagerI *SeigManagerICallerSession) Ton() (common.Address, error) {
	return _SeigManagerI.Contract.Ton(&_SeigManagerI.CallOpts)
}

// Tot is a free data retrieval call binding the contract method 0xa16d6aa7.
//
// Solidity: function tot() constant returns(address)
func (_SeigManagerI *SeigManagerICaller) Tot(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SeigManagerI.contract.Call(opts, out, "tot")
	return *ret0, err
}

// Tot is a free data retrieval call binding the contract method 0xa16d6aa7.
//
// Solidity: function tot() constant returns(address)
func (_SeigManagerI *SeigManagerISession) Tot() (common.Address, error) {
	return _SeigManagerI.Contract.Tot(&_SeigManagerI.CallOpts)
}

// Tot is a free data retrieval call binding the contract method 0xa16d6aa7.
//
// Solidity: function tot() constant returns(address)
func (_SeigManagerI *SeigManagerICallerSession) Tot() (common.Address, error) {
	return _SeigManagerI.Contract.Tot(&_SeigManagerI.CallOpts)
}

// UncomittedStakeOf is a free data retrieval call binding the contract method 0xfa9789c8.
//
// Solidity: function uncomittedStakeOf(address rootchain, address account) constant returns(uint256)
func (_SeigManagerI *SeigManagerICaller) UncomittedStakeOf(opts *bind.CallOpts, rootchain common.Address, account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SeigManagerI.contract.Call(opts, out, "uncomittedStakeOf", rootchain, account)
	return *ret0, err
}

// UncomittedStakeOf is a free data retrieval call binding the contract method 0xfa9789c8.
//
// Solidity: function uncomittedStakeOf(address rootchain, address account) constant returns(uint256)
func (_SeigManagerI *SeigManagerISession) UncomittedStakeOf(rootchain common.Address, account common.Address) (*big.Int, error) {
	return _SeigManagerI.Contract.UncomittedStakeOf(&_SeigManagerI.CallOpts, rootchain, account)
}

// UncomittedStakeOf is a free data retrieval call binding the contract method 0xfa9789c8.
//
// Solidity: function uncomittedStakeOf(address rootchain, address account) constant returns(uint256)
func (_SeigManagerI *SeigManagerICallerSession) UncomittedStakeOf(rootchain common.Address, account common.Address) (*big.Int, error) {
	return _SeigManagerI.Contract.UncomittedStakeOf(&_SeigManagerI.CallOpts, rootchain, account)
}

// UnpausedBlock is a free data retrieval call binding the contract method 0x1cc47890.
//
// Solidity: function unpausedBlock() constant returns(uint256)
func (_SeigManagerI *SeigManagerICaller) UnpausedBlock(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SeigManagerI.contract.Call(opts, out, "unpausedBlock")
	return *ret0, err
}

// UnpausedBlock is a free data retrieval call binding the contract method 0x1cc47890.
//
// Solidity: function unpausedBlock() constant returns(uint256)
func (_SeigManagerI *SeigManagerISession) UnpausedBlock() (*big.Int, error) {
	return _SeigManagerI.Contract.UnpausedBlock(&_SeigManagerI.CallOpts)
}

// UnpausedBlock is a free data retrieval call binding the contract method 0x1cc47890.
//
// Solidity: function unpausedBlock() constant returns(uint256)
func (_SeigManagerI *SeigManagerICallerSession) UnpausedBlock() (*big.Int, error) {
	return _SeigManagerI.Contract.UnpausedBlock(&_SeigManagerI.CallOpts)
}

// Wton is a free data retrieval call binding the contract method 0x8d62d949.
//
// Solidity: function wton() constant returns(address)
func (_SeigManagerI *SeigManagerICaller) Wton(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SeigManagerI.contract.Call(opts, out, "wton")
	return *ret0, err
}

// Wton is a free data retrieval call binding the contract method 0x8d62d949.
//
// Solidity: function wton() constant returns(address)
func (_SeigManagerI *SeigManagerISession) Wton() (common.Address, error) {
	return _SeigManagerI.Contract.Wton(&_SeigManagerI.CallOpts)
}

// Wton is a free data retrieval call binding the contract method 0x8d62d949.
//
// Solidity: function wton() constant returns(address)
func (_SeigManagerI *SeigManagerICallerSession) Wton() (common.Address, error) {
	return _SeigManagerI.Contract.Wton(&_SeigManagerI.CallOpts)
}

// DeployCoinage is a paid mutator transaction binding the contract method 0x833a774f.
//
// Solidity: function deployCoinage(address rootchain) returns(bool)
func (_SeigManagerI *SeigManagerITransactor) DeployCoinage(opts *bind.TransactOpts, rootchain common.Address) (*types.Transaction, error) {
	return _SeigManagerI.contract.Transact(opts, "deployCoinage", rootchain)
}

// DeployCoinage is a paid mutator transaction binding the contract method 0x833a774f.
//
// Solidity: function deployCoinage(address rootchain) returns(bool)
func (_SeigManagerI *SeigManagerISession) DeployCoinage(rootchain common.Address) (*types.Transaction, error) {
	return _SeigManagerI.Contract.DeployCoinage(&_SeigManagerI.TransactOpts, rootchain)
}

// DeployCoinage is a paid mutator transaction binding the contract method 0x833a774f.
//
// Solidity: function deployCoinage(address rootchain) returns(bool)
func (_SeigManagerI *SeigManagerITransactorSession) DeployCoinage(rootchain common.Address) (*types.Transaction, error) {
	return _SeigManagerI.Contract.DeployCoinage(&_SeigManagerI.TransactOpts, rootchain)
}

// OnCommit is a paid mutator transaction binding the contract method 0x359c4d59.
//
// Solidity: function onCommit() returns(bool)
func (_SeigManagerI *SeigManagerITransactor) OnCommit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SeigManagerI.contract.Transact(opts, "onCommit")
}

// OnCommit is a paid mutator transaction binding the contract method 0x359c4d59.
//
// Solidity: function onCommit() returns(bool)
func (_SeigManagerI *SeigManagerISession) OnCommit() (*types.Transaction, error) {
	return _SeigManagerI.Contract.OnCommit(&_SeigManagerI.TransactOpts)
}

// OnCommit is a paid mutator transaction binding the contract method 0x359c4d59.
//
// Solidity: function onCommit() returns(bool)
func (_SeigManagerI *SeigManagerITransactorSession) OnCommit() (*types.Transaction, error) {
	return _SeigManagerI.Contract.OnCommit(&_SeigManagerI.TransactOpts)
}

// OnDeposit is a paid mutator transaction binding the contract method 0x412c6d50.
//
// Solidity: function onDeposit(address rootchain, address account, uint256 amount) returns(bool)
func (_SeigManagerI *SeigManagerITransactor) OnDeposit(opts *bind.TransactOpts, rootchain common.Address, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SeigManagerI.contract.Transact(opts, "onDeposit", rootchain, account, amount)
}

// OnDeposit is a paid mutator transaction binding the contract method 0x412c6d50.
//
// Solidity: function onDeposit(address rootchain, address account, uint256 amount) returns(bool)
func (_SeigManagerI *SeigManagerISession) OnDeposit(rootchain common.Address, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SeigManagerI.Contract.OnDeposit(&_SeigManagerI.TransactOpts, rootchain, account, amount)
}

// OnDeposit is a paid mutator transaction binding the contract method 0x412c6d50.
//
// Solidity: function onDeposit(address rootchain, address account, uint256 amount) returns(bool)
func (_SeigManagerI *SeigManagerITransactorSession) OnDeposit(rootchain common.Address, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SeigManagerI.Contract.OnDeposit(&_SeigManagerI.TransactOpts, rootchain, account, amount)
}

// OnTransfer is a paid mutator transaction binding the contract method 0x4a393149.
//
// Solidity: function onTransfer(address sender, address recipient, uint256 amount) returns(bool)
func (_SeigManagerI *SeigManagerITransactor) OnTransfer(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SeigManagerI.contract.Transact(opts, "onTransfer", sender, recipient, amount)
}

// OnTransfer is a paid mutator transaction binding the contract method 0x4a393149.
//
// Solidity: function onTransfer(address sender, address recipient, uint256 amount) returns(bool)
func (_SeigManagerI *SeigManagerISession) OnTransfer(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SeigManagerI.Contract.OnTransfer(&_SeigManagerI.TransactOpts, sender, recipient, amount)
}

// OnTransfer is a paid mutator transaction binding the contract method 0x4a393149.
//
// Solidity: function onTransfer(address sender, address recipient, uint256 amount) returns(bool)
func (_SeigManagerI *SeigManagerITransactorSession) OnTransfer(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SeigManagerI.Contract.OnTransfer(&_SeigManagerI.TransactOpts, sender, recipient, amount)
}

// OnWithdraw is a paid mutator transaction binding the contract method 0xf850ffaa.
//
// Solidity: function onWithdraw(address rootchain, address account, uint256 amount) returns(bool)
func (_SeigManagerI *SeigManagerITransactor) OnWithdraw(opts *bind.TransactOpts, rootchain common.Address, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SeigManagerI.contract.Transact(opts, "onWithdraw", rootchain, account, amount)
}

// OnWithdraw is a paid mutator transaction binding the contract method 0xf850ffaa.
//
// Solidity: function onWithdraw(address rootchain, address account, uint256 amount) returns(bool)
func (_SeigManagerI *SeigManagerISession) OnWithdraw(rootchain common.Address, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SeigManagerI.Contract.OnWithdraw(&_SeigManagerI.TransactOpts, rootchain, account, amount)
}

// OnWithdraw is a paid mutator transaction binding the contract method 0xf850ffaa.
//
// Solidity: function onWithdraw(address rootchain, address account, uint256 amount) returns(bool)
func (_SeigManagerI *SeigManagerITransactorSession) OnWithdraw(rootchain common.Address, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SeigManagerI.Contract.OnWithdraw(&_SeigManagerI.TransactOpts, rootchain, account, amount)
}

// SetCommissionRate is a paid mutator transaction binding the contract method 0x4224ed66.
//
// Solidity: function setCommissionRate(address rootchain, uint256 commission, bool isCommissionRateNegative) returns(bool)
func (_SeigManagerI *SeigManagerITransactor) SetCommissionRate(opts *bind.TransactOpts, rootchain common.Address, commission *big.Int, isCommissionRateNegative bool) (*types.Transaction, error) {
	return _SeigManagerI.contract.Transact(opts, "setCommissionRate", rootchain, commission, isCommissionRateNegative)
}

// SetCommissionRate is a paid mutator transaction binding the contract method 0x4224ed66.
//
// Solidity: function setCommissionRate(address rootchain, uint256 commission, bool isCommissionRateNegative) returns(bool)
func (_SeigManagerI *SeigManagerISession) SetCommissionRate(rootchain common.Address, commission *big.Int, isCommissionRateNegative bool) (*types.Transaction, error) {
	return _SeigManagerI.Contract.SetCommissionRate(&_SeigManagerI.TransactOpts, rootchain, commission, isCommissionRateNegative)
}

// SetCommissionRate is a paid mutator transaction binding the contract method 0x4224ed66.
//
// Solidity: function setCommissionRate(address rootchain, uint256 commission, bool isCommissionRateNegative) returns(bool)
func (_SeigManagerI *SeigManagerITransactorSession) SetCommissionRate(rootchain common.Address, commission *big.Int, isCommissionRateNegative bool) (*types.Transaction, error) {
	return _SeigManagerI.Contract.SetCommissionRate(&_SeigManagerI.TransactOpts, rootchain, commission, isCommissionRateNegative)
}
