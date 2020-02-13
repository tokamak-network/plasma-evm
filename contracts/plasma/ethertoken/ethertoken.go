// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ethertoken

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

// BasicTokenABI is the input ABI used to generate the binding from.
const BasicTokenABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// BasicTokenFuncSigs maps the 4-byte function signature to its string representation.
var BasicTokenFuncSigs = map[string]string{
	"70a08231": "balanceOf(address)",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
}

// BasicTokenBin is the compiled bytecode used for deploying new contracts.
var BasicTokenBin = "0x608060405234801561001057600080fd5b5061021e806100206000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c806318160ddd1461004657806370a0823114610060578063a9059cbb14610086575b600080fd5b61004e6100c6565b60408051918252519081900360200190f35b61004e6004803603602081101561007657600080fd5b50356001600160a01b03166100cc565b6100b26004803603604081101561009c57600080fd5b506001600160a01b0381351690602001356100e7565b604080519115158252519081900360200190f35b60015490565b6001600160a01b031660009081526020819052604090205490565b3360009081526020819052604081205482111561010357600080fd5b6001600160a01b03831661011657600080fd5b33600090815260208190526040902054610136908363ffffffff6101c416565b33600090815260208190526040808220929092556001600160a01b03851681522054610168908363ffffffff6101d616565b6001600160a01b038416600081815260208181526040918290209390935580518581529051919233927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9281900390910190a350600192915050565b6000828211156101d057fe5b50900390565b818101828110156101e357fe5b9291505056fea265627a7a72315820fab8a7f7ff4ea62418cee2b1150fb40410bbb8ecc14cb2ba7c51a3a9f1c3f1ed64736f6c634300050c0032"

// DeployBasicToken deploys a new Ethereum contract, binding an instance of BasicToken to it.
func DeployBasicToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BasicToken, error) {
	parsed, err := abi.JSON(strings.NewReader(BasicTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BasicTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BasicToken{BasicTokenCaller: BasicTokenCaller{contract: contract}, BasicTokenTransactor: BasicTokenTransactor{contract: contract}, BasicTokenFilterer: BasicTokenFilterer{contract: contract}}, nil
}

// BasicToken is an auto generated Go binding around an Ethereum contract.
type BasicToken struct {
	BasicTokenCaller     // Read-only binding to the contract
	BasicTokenTransactor // Write-only binding to the contract
	BasicTokenFilterer   // Log filterer for contract events
}

// BasicTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type BasicTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasicTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BasicTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasicTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BasicTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasicTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BasicTokenSession struct {
	Contract     *BasicToken       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BasicTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BasicTokenCallerSession struct {
	Contract *BasicTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// BasicTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BasicTokenTransactorSession struct {
	Contract     *BasicTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// BasicTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type BasicTokenRaw struct {
	Contract *BasicToken // Generic contract binding to access the raw methods on
}

// BasicTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BasicTokenCallerRaw struct {
	Contract *BasicTokenCaller // Generic read-only contract binding to access the raw methods on
}

// BasicTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BasicTokenTransactorRaw struct {
	Contract *BasicTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBasicToken creates a new instance of BasicToken, bound to a specific deployed contract.
func NewBasicToken(address common.Address, backend bind.ContractBackend) (*BasicToken, error) {
	contract, err := bindBasicToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BasicToken{BasicTokenCaller: BasicTokenCaller{contract: contract}, BasicTokenTransactor: BasicTokenTransactor{contract: contract}, BasicTokenFilterer: BasicTokenFilterer{contract: contract}}, nil
}

// NewBasicTokenCaller creates a new read-only instance of BasicToken, bound to a specific deployed contract.
func NewBasicTokenCaller(address common.Address, caller bind.ContractCaller) (*BasicTokenCaller, error) {
	contract, err := bindBasicToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BasicTokenCaller{contract: contract}, nil
}

// NewBasicTokenTransactor creates a new write-only instance of BasicToken, bound to a specific deployed contract.
func NewBasicTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*BasicTokenTransactor, error) {
	contract, err := bindBasicToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BasicTokenTransactor{contract: contract}, nil
}

// NewBasicTokenFilterer creates a new log filterer instance of BasicToken, bound to a specific deployed contract.
func NewBasicTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*BasicTokenFilterer, error) {
	contract, err := bindBasicToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BasicTokenFilterer{contract: contract}, nil
}

// bindBasicToken binds a generic wrapper to an already deployed contract.
func bindBasicToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BasicTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BasicToken *BasicTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BasicToken.Contract.BasicTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BasicToken *BasicTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BasicToken.Contract.BasicTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BasicToken *BasicTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BasicToken.Contract.BasicTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BasicToken *BasicTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BasicToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BasicToken *BasicTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BasicToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BasicToken *BasicTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BasicToken.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) constant returns(uint256)
func (_BasicToken *BasicTokenCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BasicToken.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) constant returns(uint256)
func (_BasicToken *BasicTokenSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _BasicToken.Contract.BalanceOf(&_BasicToken.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) constant returns(uint256)
func (_BasicToken *BasicTokenCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _BasicToken.Contract.BalanceOf(&_BasicToken.CallOpts, _owner)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_BasicToken *BasicTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BasicToken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_BasicToken *BasicTokenSession) TotalSupply() (*big.Int, error) {
	return _BasicToken.Contract.TotalSupply(&_BasicToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_BasicToken *BasicTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _BasicToken.Contract.TotalSupply(&_BasicToken.CallOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_BasicToken *BasicTokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BasicToken.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_BasicToken *BasicTokenSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BasicToken.Contract.Transfer(&_BasicToken.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_BasicToken *BasicTokenTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BasicToken.Contract.Transfer(&_BasicToken.TransactOpts, _to, _value)
}

// BasicTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the BasicToken contract.
type BasicTokenTransferIterator struct {
	Event *BasicTokenTransfer // Event containing the contract specifics and raw log

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
func (it *BasicTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasicTokenTransfer)
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
		it.Event = new(BasicTokenTransfer)
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
func (it *BasicTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasicTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasicTokenTransfer represents a Transfer event raised by the BasicToken contract.
type BasicTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_BasicToken *BasicTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BasicTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BasicToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BasicTokenTransferIterator{contract: _BasicToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_BasicToken *BasicTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *BasicTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BasicToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasicTokenTransfer)
				if err := _BasicToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_BasicToken *BasicTokenFilterer) ParseTransfer(log types.Log) (*BasicTokenTransfer, error) {
	event := new(BasicTokenTransfer)
	if err := _BasicToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ERC20ABI is the input ABI used to generate the binding from.
const ERC20ABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_who\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ERC20FuncSigs maps the 4-byte function signature to its string representation.
var ERC20FuncSigs = map[string]string{
	"dd62ed3e": "allowance(address,address)",
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
}

// ERC20 is an auto generated Go binding around an Ethereum contract.
type ERC20 struct {
	ERC20Caller     // Read-only binding to the contract
	ERC20Transactor // Write-only binding to the contract
	ERC20Filterer   // Log filterer for contract events
}

// ERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20Session struct {
	Contract     *ERC20            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20CallerSession struct {
	Contract *ERC20Caller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20TransactorSession struct {
	Contract     *ERC20Transactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20Raw struct {
	Contract *ERC20 // Generic contract binding to access the raw methods on
}

// ERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20CallerRaw struct {
	Contract *ERC20Caller // Generic read-only contract binding to access the raw methods on
}

// ERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20TransactorRaw struct {
	Contract *ERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20 creates a new instance of ERC20, bound to a specific deployed contract.
func NewERC20(address common.Address, backend bind.ContractBackend) (*ERC20, error) {
	contract, err := bindERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20{ERC20Caller: ERC20Caller{contract: contract}, ERC20Transactor: ERC20Transactor{contract: contract}, ERC20Filterer: ERC20Filterer{contract: contract}}, nil
}

// NewERC20Caller creates a new read-only instance of ERC20, bound to a specific deployed contract.
func NewERC20Caller(address common.Address, caller bind.ContractCaller) (*ERC20Caller, error) {
	contract, err := bindERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20Caller{contract: contract}, nil
}

// NewERC20Transactor creates a new write-only instance of ERC20, bound to a specific deployed contract.
func NewERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*ERC20Transactor, error) {
	contract, err := bindERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20Transactor{contract: contract}, nil
}

// NewERC20Filterer creates a new log filterer instance of ERC20, bound to a specific deployed contract.
func NewERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*ERC20Filterer, error) {
	contract, err := bindERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20Filterer{contract: contract}, nil
}

// bindERC20 binds a generic wrapper to an already deployed contract.
func bindERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20 *ERC20Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20.Contract.ERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20 *ERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20.Contract.ERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20 *ERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20.Contract.ERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20 *ERC20CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20 *ERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20 *ERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) constant returns(uint256)
func (_ERC20 *ERC20Caller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) constant returns(uint256)
func (_ERC20 *ERC20Session) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _ERC20.Contract.Allowance(&_ERC20.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) constant returns(uint256)
func (_ERC20 *ERC20CallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _ERC20.Contract.Allowance(&_ERC20.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _who) constant returns(uint256)
func (_ERC20 *ERC20Caller) BalanceOf(opts *bind.CallOpts, _who common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20.contract.Call(opts, out, "balanceOf", _who)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _who) constant returns(uint256)
func (_ERC20 *ERC20Session) BalanceOf(_who common.Address) (*big.Int, error) {
	return _ERC20.Contract.BalanceOf(&_ERC20.CallOpts, _who)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _who) constant returns(uint256)
func (_ERC20 *ERC20CallerSession) BalanceOf(_who common.Address) (*big.Int, error) {
	return _ERC20.Contract.BalanceOf(&_ERC20.CallOpts, _who)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20 *ERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20 *ERC20Session) TotalSupply() (*big.Int, error) {
	return _ERC20.Contract.TotalSupply(&_ERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20 *ERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _ERC20.Contract.TotalSupply(&_ERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_ERC20 *ERC20Transactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_ERC20 *ERC20Session) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Approve(&_ERC20.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_ERC20 *ERC20TransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Approve(&_ERC20.TransactOpts, _spender, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_ERC20 *ERC20Transactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_ERC20 *ERC20Session) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Transfer(&_ERC20.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_ERC20 *ERC20TransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Transfer(&_ERC20.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_ERC20 *ERC20Transactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_ERC20 *ERC20Session) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.TransferFrom(&_ERC20.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_ERC20 *ERC20TransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.TransferFrom(&_ERC20.TransactOpts, _from, _to, _value)
}

// ERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC20 contract.
type ERC20ApprovalIterator struct {
	Event *ERC20Approval // Event containing the contract specifics and raw log

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
func (it *ERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20Approval)
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
		it.Event = new(ERC20Approval)
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
func (it *ERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20Approval represents a Approval event raised by the ERC20 contract.
type ERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20 *ERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ERC20ApprovalIterator{contract: _ERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20 *ERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20Approval)
				if err := _ERC20.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20 *ERC20Filterer) ParseApproval(log types.Log) (*ERC20Approval, error) {
	event := new(ERC20Approval)
	if err := _ERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC20 contract.
type ERC20TransferIterator struct {
	Event *ERC20Transfer // Event containing the contract specifics and raw log

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
func (it *ERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20Transfer)
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
		it.Event = new(ERC20Transfer)
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
func (it *ERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20Transfer represents a Transfer event raised by the ERC20 contract.
type ERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20 *ERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TransferIterator{contract: _ERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20 *ERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20Transfer)
				if err := _ERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20 *ERC20Filterer) ParseTransfer(log types.Log) (*ERC20Transfer, error) {
	event := new(ERC20Transfer)
	if err := _ERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ERC20BasicABI is the input ABI used to generate the binding from.
const ERC20BasicABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_who\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ERC20BasicFuncSigs maps the 4-byte function signature to its string representation.
var ERC20BasicFuncSigs = map[string]string{
	"70a08231": "balanceOf(address)",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
}

// ERC20Basic is an auto generated Go binding around an Ethereum contract.
type ERC20Basic struct {
	ERC20BasicCaller     // Read-only binding to the contract
	ERC20BasicTransactor // Write-only binding to the contract
	ERC20BasicFilterer   // Log filterer for contract events
}

// ERC20BasicCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20BasicCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20BasicTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20BasicTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20BasicFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20BasicFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20BasicSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20BasicSession struct {
	Contract     *ERC20Basic       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20BasicCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20BasicCallerSession struct {
	Contract *ERC20BasicCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ERC20BasicTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20BasicTransactorSession struct {
	Contract     *ERC20BasicTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ERC20BasicRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20BasicRaw struct {
	Contract *ERC20Basic // Generic contract binding to access the raw methods on
}

// ERC20BasicCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20BasicCallerRaw struct {
	Contract *ERC20BasicCaller // Generic read-only contract binding to access the raw methods on
}

// ERC20BasicTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20BasicTransactorRaw struct {
	Contract *ERC20BasicTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20Basic creates a new instance of ERC20Basic, bound to a specific deployed contract.
func NewERC20Basic(address common.Address, backend bind.ContractBackend) (*ERC20Basic, error) {
	contract, err := bindERC20Basic(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20Basic{ERC20BasicCaller: ERC20BasicCaller{contract: contract}, ERC20BasicTransactor: ERC20BasicTransactor{contract: contract}, ERC20BasicFilterer: ERC20BasicFilterer{contract: contract}}, nil
}

// NewERC20BasicCaller creates a new read-only instance of ERC20Basic, bound to a specific deployed contract.
func NewERC20BasicCaller(address common.Address, caller bind.ContractCaller) (*ERC20BasicCaller, error) {
	contract, err := bindERC20Basic(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20BasicCaller{contract: contract}, nil
}

// NewERC20BasicTransactor creates a new write-only instance of ERC20Basic, bound to a specific deployed contract.
func NewERC20BasicTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC20BasicTransactor, error) {
	contract, err := bindERC20Basic(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20BasicTransactor{contract: contract}, nil
}

// NewERC20BasicFilterer creates a new log filterer instance of ERC20Basic, bound to a specific deployed contract.
func NewERC20BasicFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC20BasicFilterer, error) {
	contract, err := bindERC20Basic(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20BasicFilterer{contract: contract}, nil
}

// bindERC20Basic binds a generic wrapper to an already deployed contract.
func bindERC20Basic(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20BasicABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Basic *ERC20BasicRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20Basic.Contract.ERC20BasicCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Basic *ERC20BasicRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Basic.Contract.ERC20BasicTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Basic *ERC20BasicRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Basic.Contract.ERC20BasicTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Basic *ERC20BasicCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20Basic.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Basic *ERC20BasicTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Basic.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Basic *ERC20BasicTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Basic.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _who) constant returns(uint256)
func (_ERC20Basic *ERC20BasicCaller) BalanceOf(opts *bind.CallOpts, _who common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20Basic.contract.Call(opts, out, "balanceOf", _who)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _who) constant returns(uint256)
func (_ERC20Basic *ERC20BasicSession) BalanceOf(_who common.Address) (*big.Int, error) {
	return _ERC20Basic.Contract.BalanceOf(&_ERC20Basic.CallOpts, _who)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _who) constant returns(uint256)
func (_ERC20Basic *ERC20BasicCallerSession) BalanceOf(_who common.Address) (*big.Int, error) {
	return _ERC20Basic.Contract.BalanceOf(&_ERC20Basic.CallOpts, _who)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20Basic *ERC20BasicCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20Basic.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20Basic *ERC20BasicSession) TotalSupply() (*big.Int, error) {
	return _ERC20Basic.Contract.TotalSupply(&_ERC20Basic.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20Basic *ERC20BasicCallerSession) TotalSupply() (*big.Int, error) {
	return _ERC20Basic.Contract.TotalSupply(&_ERC20Basic.CallOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_ERC20Basic *ERC20BasicTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20Basic.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_ERC20Basic *ERC20BasicSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20Basic.Contract.Transfer(&_ERC20Basic.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_ERC20Basic *ERC20BasicTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20Basic.Contract.Transfer(&_ERC20Basic.TransactOpts, _to, _value)
}

// ERC20BasicTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC20Basic contract.
type ERC20BasicTransferIterator struct {
	Event *ERC20BasicTransfer // Event containing the contract specifics and raw log

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
func (it *ERC20BasicTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20BasicTransfer)
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
		it.Event = new(ERC20BasicTransfer)
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
func (it *ERC20BasicTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20BasicTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20BasicTransfer represents a Transfer event raised by the ERC20Basic contract.
type ERC20BasicTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20Basic *ERC20BasicFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ERC20BasicTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20Basic.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERC20BasicTransferIterator{contract: _ERC20Basic.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20Basic *ERC20BasicFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC20BasicTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20Basic.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20BasicTransfer)
				if err := _ERC20Basic.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20Basic *ERC20BasicFilterer) ParseTransfer(log types.Log) (*ERC20BasicTransfer, error) {
	event := new(ERC20BasicTransfer)
	if err := _ERC20Basic.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// EtherTokenABI is the input ABI used to generate the binding from.
const EtherTokenABI = "[{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_development\",\"type\":\"bool\"},{\"internalType\":\"contractERC20\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_swapEnabled\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Depositted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"_isExit\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_requestor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_trieKey\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"RequestCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"isExit\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"requestor\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"trieKey\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"trieValue\",\"type\":\"bytes\"}],\"name\":\"applyRequestInChildChain\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"isExit\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"requestor\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"trieKey\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"trieValue\",\"type\":\"bytes\"}],\"name\":\"applyRequestInRootChain\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"trieValue\",\"type\":\"bytes\"}],\"name\":\"decodeTrieValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"v\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseApproval\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"development\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_who\",\"type\":\"address\"}],\"name\":\"getBalanceTrieKey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseApproval\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rootchain\",\"type\":\"address\"}],\"name\":\"init\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"initialized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rootchain\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"swapEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"swapFromEth\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"swapToEth\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractERC20\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// EtherTokenFuncSigs maps the 4-byte function signature to its string representation.
var EtherTokenFuncSigs = map[string]string{
	"dd62ed3e": "allowance(address,address)",
	"141ecf46": "applyRequestInChildChain(bool,uint256,address,bytes32,bytes)",
	"a9f79308": "applyRequestInRootChain(bool,uint256,address,bytes32,bytes)",
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"b9e59d09": "decodeTrieValue(bytes)",
	"66188463": "decreaseApproval(address,uint256)",
	"b6b55f25": "deposit(uint256)",
	"7b929c27": "development()",
	"b18fcfdf": "getBalanceTrieKey(address)",
	"d73dd623": "increaseApproval(address,uint256)",
	"19ab453c": "init(address)",
	"158ef93e": "initialized()",
	"bcc7874d": "rootchain()",
	"6ddd1713": "swapEnabled()",
	"59310da7": "swapFromEth()",
	"b06c6729": "swapToEth(uint256)",
	"fc0c546a": "token()",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
	"2e1a7d4d": "withdraw(uint256)",
}

// EtherTokenBin is the compiled bytecode used for deploying new contracts.
var EtherTokenBin = "0x608060405234801561001057600080fd5b506040516111ae3803806111ae8339818101604052606081101561003357600080fd5b50805160208201516040909201516003805461ff00191661010084151502179055600480546001600160a01b0319166001600160a01b03851690811790915591929115838061009957508080156100875750815b80610099575080158015610099575081155b6100a257600080fd5b5060048054911515740100000000000000000000000000000000000000000260ff60a01b1990921691909117905550506110cd806100e16000396000f3fe6080604052600436106101355760003560e01c80637b929c27116100ab578063b6b55f251161006f578063b6b55f25146104d4578063b9e59d09146104fe578063bcc7874d146105b1578063d73dd623146105e2578063dd62ed3e1461061b578063fc0c546a1461065657610135565b80637b929c2714610389578063a9059cbb1461039e578063a9f79308146103d7578063b06c672914610477578063b18fcfdf146104a157610135565b806323b872dd116100fd57806323b872dd1461029b5780632e1a7d4d146102de57806359310da71461013557806366188463146103085780636ddd17131461034157806370a082311461035657610135565b8063095ea7b31461013f578063141ecf461461018c578063158ef93e1461022c57806318160ddd1461024157806319ab453c14610268575b61013d61066b565b005b34801561014b57600080fd5b506101786004803603604081101561016257600080fd5b506001600160a01b03813516906020013561068d565b604080519115158252519081900360200190f35b34801561019857600080fd5b50610178600480360360a08110156101af57600080fd5b81351515916020810135916001600160a01b036040830135169160608101359181019060a0810160808201356401000000008111156101ed57600080fd5b8201836020820111156101ff57600080fd5b8035906020019184600183028401116401000000008311171561022157600080fd5b5090925090506106f3565b34801561023857600080fd5b506101786107e8565b34801561024d57600080fd5b506102566107f1565b60408051918252519081900360200190f35b34801561027457600080fd5b506101786004803603602081101561028b57600080fd5b50356001600160a01b03166107f7565b3480156102a757600080fd5b50610178600480360360608110156102be57600080fd5b506001600160a01b0381358116916020810135909116906040013561083d565b3480156102ea57600080fd5b506101786004803603602081101561030157600080fd5b503561099e565b34801561031457600080fd5b506101786004803603604081101561032b57600080fd5b506001600160a01b038135169060200135610a87565b34801561034d57600080fd5b50610178610b76565b34801561036257600080fd5b506102566004803603602081101561037957600080fd5b50356001600160a01b0316610b86565b34801561039557600080fd5b50610178610ba1565b3480156103aa57600080fd5b50610178600480360360408110156103c157600080fd5b506001600160a01b038135169060200135610baf565b3480156103e357600080fd5b50610178600480360360a08110156103fa57600080fd5b81351515916020810135916001600160a01b036040830135169160608101359181019060a08101608082013564010000000081111561043857600080fd5b82018360208201111561044a57600080fd5b8035906020019184600183028401116401000000008311171561046c57600080fd5b509092509050610c7a565b34801561048357600080fd5b5061013d6004803603602081101561049a57600080fd5b5035610d1a565b3480156104ad57600080fd5b50610256600480360360208110156104c457600080fd5b50356001600160a01b0316610d51565b3480156104e057600080fd5b50610178600480360360208110156104f757600080fd5b5035610d93565b34801561050a57600080fd5b506102566004803603602081101561052157600080fd5b81019060208101813564010000000081111561053c57600080fd5b82018360208201111561054e57600080fd5b8035906020019184600183028401116401000000008311171561057057600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610e45945050505050565b3480156105bd57600080fd5b506105c6610e5d565b604080516001600160a01b039092168252519081900360200190f35b3480156105ee57600080fd5b506101786004803603604081101561060557600080fd5b506001600160a01b038135169060200135610e72565b34801561062757600080fd5b506102566004803603604081101561063e57600080fd5b506001600160a01b0381358116916020013516610f0b565b34801561066257600080fd5b506105c6610f36565b600454600160a01b900460ff1661068157600080fd5b61068b3334610f45565b565b3360008181526002602090815260408083206001600160a01b038716808552908352818420869055815186815291519394909390927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925928290030190a350600192915050565b600354600090610100900460ff168061070a575033155b61071357600080fd5b61071c85610d51565b841461072757600080fd5b600061076884848080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250610e4592505050565b9050871561077f5761077a8682610fc8565b610789565b6107898682610f45565b6040805189151581526001600160a01b03881660208201528082018790526060810183905290517f4b0c23cd90bac23cdaa0958fbd0306f53e20cd5e65676d21b9ceb109784e15b89181900360800190a1506001979650505050505050565b60035460ff1681565b60015490565b60035460009060ff161561080a57600080fd5b6003805460ff196001600160a01b03909416620100000262010000600160b01b0319909116179290921660011790915590565b6001600160a01b03831660009081526020819052604081205482111561086257600080fd5b6001600160a01b038416600090815260026020908152604080832033845290915290205482111561089257600080fd5b6001600160a01b0383166108a557600080fd5b6001600160a01b0384166000908152602081905260409020546108ce908363ffffffff61105316565b6001600160a01b038086166000908152602081905260408082209390935590851681522054610903908363ffffffff61106516565b6001600160a01b03808516600090815260208181526040808320949094559187168152600282528281203382529091522054610945908363ffffffff61105316565b6001600160a01b0380861660008181526002602090815260408083203384528252918290209490945580518681529051928716939192600080516020611079833981519152929181900390910190a35060019392505050565b60035460009060ff166109b057600080fd5b6109ba3383610fc8565b604080513381526020810184905281517f7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5929181900390910190a1600480546040805163a9059cbb60e01b8152339381019390935260248301859052516001600160a01b039091169163a9059cbb9160448083019260209291908290030181600087803b158015610a4a57600080fd5b505af1158015610a5e573d6000803e3d6000fd5b505050506040513d6020811015610a7457600080fd5b5051610a7f57600080fd5b506001919050565b3360009081526002602090815260408083206001600160a01b0386168452909152812054808310610adb573360009081526002602090815260408083206001600160a01b0388168452909152812055610b10565b610aeb818463ffffffff61105316565b3360009081526002602090815260408083206001600160a01b03891684529091529020555b3360008181526002602090815260408083206001600160a01b0389168085529083529281902054815190815290519293927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929181900390910190a35060019392505050565b600454600160a01b900460ff1681565b6001600160a01b031660009081526020819052604090205490565b600354610100900460ff1681565b33600090815260208190526040812054821115610bcb57600080fd5b6001600160a01b038316610bde57600080fd5b33600090815260208190526040902054610bfe908363ffffffff61105316565b33600090815260208190526040808220929092556001600160a01b03851681522054610c30908363ffffffff61106516565b6001600160a01b038416600081815260208181526040918290209390935580518581529051919233926000805160206110798339815191529281900390910190a350600192915050565b60035460009060ff16610c8c57600080fd5b6003546201000090046001600160a01b03163314610ca957600080fd5b610cb285610d51565b8414610cbd57600080fd5b6000610cfe84848080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250610e4592505050565b90508715610d105761077a8682610f45565b6107898682610fc8565b600454600160a01b900460ff16610d3057600080fd5b610d3b33308361083d565b610d4457600080fd5b610d4e3382610fc8565b50565b60408051600060208083019190915260609390931b6bffffffffffffffffffffffff191681830152815180820360340181526054909101909152805191012090565b60035460009060ff16610da557600080fd5b610daf3383610f45565b604080513381526020810184905281517f46fc0825df492a7a7f1fb1b690f6912ecb69575484fbb24be0f398ae70d97db4929181900390910190a160048054604080516323b872dd60e01b8152339381019390935230602484015260448301859052516001600160a01b03909116916323b872dd9160648083019260209291908290030181600087803b158015610a4a57600080fd5b60008151602014610e5557600080fd5b506020015190565b6003546201000090046001600160a01b031681565b3360009081526002602090815260408083206001600160a01b0386168452909152812054610ea6908363ffffffff61106516565b3360008181526002602090815260408083206001600160a01b0389168085529083529281902085905580519485525191937f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929081900390910190a350600192915050565b6001600160a01b03918216600090815260026020908152604080832093909416825291909152205490565b6004546001600160a01b031681565b600154610f58908263ffffffff61106516565b6001556001600160a01b038216600090815260208190526040902054610f84908263ffffffff61106516565b6001600160a01b0383166000818152602081815260408083209490945583518581529351929391926000805160206110798339815191529281900390910190a35050565b6001600160a01b038216600090815260208190526040902054610ff1908263ffffffff61105316565b6001600160a01b03831660009081526020819052604090205560015461101d908263ffffffff61105316565b6001556040805182815290516000916001600160a01b038516916000805160206110798339815191529181900360200190a35050565b60008282111561105f57fe5b50900390565b8181018281101561107257fe5b9291505056feddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3efa265627a7a723158209a1372fa2218ea99ed56f23e57a34ed8b2f0e1049050353422da8e310723a1b464736f6c634300050c0032"

// DeployEtherToken deploys a new Ethereum contract, binding an instance of EtherToken to it.
func DeployEtherToken(auth *bind.TransactOpts, backend bind.ContractBackend, _development bool, _token common.Address, _swapEnabled bool) (common.Address, *types.Transaction, *EtherToken, error) {
	parsed, err := abi.JSON(strings.NewReader(EtherTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(EtherTokenBin), backend, _development, _token, _swapEnabled)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EtherToken{EtherTokenCaller: EtherTokenCaller{contract: contract}, EtherTokenTransactor: EtherTokenTransactor{contract: contract}, EtherTokenFilterer: EtherTokenFilterer{contract: contract}}, nil
}

// EtherToken is an auto generated Go binding around an Ethereum contract.
type EtherToken struct {
	EtherTokenCaller     // Read-only binding to the contract
	EtherTokenTransactor // Write-only binding to the contract
	EtherTokenFilterer   // Log filterer for contract events
}

// EtherTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type EtherTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EtherTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EtherTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EtherTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EtherTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EtherTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EtherTokenSession struct {
	Contract     *EtherToken       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EtherTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EtherTokenCallerSession struct {
	Contract *EtherTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// EtherTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EtherTokenTransactorSession struct {
	Contract     *EtherTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// EtherTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type EtherTokenRaw struct {
	Contract *EtherToken // Generic contract binding to access the raw methods on
}

// EtherTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EtherTokenCallerRaw struct {
	Contract *EtherTokenCaller // Generic read-only contract binding to access the raw methods on
}

// EtherTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EtherTokenTransactorRaw struct {
	Contract *EtherTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEtherToken creates a new instance of EtherToken, bound to a specific deployed contract.
func NewEtherToken(address common.Address, backend bind.ContractBackend) (*EtherToken, error) {
	contract, err := bindEtherToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EtherToken{EtherTokenCaller: EtherTokenCaller{contract: contract}, EtherTokenTransactor: EtherTokenTransactor{contract: contract}, EtherTokenFilterer: EtherTokenFilterer{contract: contract}}, nil
}

// NewEtherTokenCaller creates a new read-only instance of EtherToken, bound to a specific deployed contract.
func NewEtherTokenCaller(address common.Address, caller bind.ContractCaller) (*EtherTokenCaller, error) {
	contract, err := bindEtherToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EtherTokenCaller{contract: contract}, nil
}

// NewEtherTokenTransactor creates a new write-only instance of EtherToken, bound to a specific deployed contract.
func NewEtherTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*EtherTokenTransactor, error) {
	contract, err := bindEtherToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EtherTokenTransactor{contract: contract}, nil
}

// NewEtherTokenFilterer creates a new log filterer instance of EtherToken, bound to a specific deployed contract.
func NewEtherTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*EtherTokenFilterer, error) {
	contract, err := bindEtherToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EtherTokenFilterer{contract: contract}, nil
}

// bindEtherToken binds a generic wrapper to an already deployed contract.
func bindEtherToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EtherTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EtherToken *EtherTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _EtherToken.Contract.EtherTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EtherToken *EtherTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EtherToken.Contract.EtherTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EtherToken *EtherTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EtherToken.Contract.EtherTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EtherToken *EtherTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _EtherToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EtherToken *EtherTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EtherToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EtherToken *EtherTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EtherToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) constant returns(uint256)
func (_EtherToken *EtherTokenCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _EtherToken.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) constant returns(uint256)
func (_EtherToken *EtherTokenSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _EtherToken.Contract.Allowance(&_EtherToken.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) constant returns(uint256)
func (_EtherToken *EtherTokenCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _EtherToken.Contract.Allowance(&_EtherToken.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) constant returns(uint256)
func (_EtherToken *EtherTokenCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _EtherToken.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) constant returns(uint256)
func (_EtherToken *EtherTokenSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _EtherToken.Contract.BalanceOf(&_EtherToken.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) constant returns(uint256)
func (_EtherToken *EtherTokenCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _EtherToken.Contract.BalanceOf(&_EtherToken.CallOpts, _owner)
}

// DecodeTrieValue is a free data retrieval call binding the contract method 0xb9e59d09.
//
// Solidity: function decodeTrieValue(bytes trieValue) constant returns(uint256 v)
func (_EtherToken *EtherTokenCaller) DecodeTrieValue(opts *bind.CallOpts, trieValue []byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _EtherToken.contract.Call(opts, out, "decodeTrieValue", trieValue)
	return *ret0, err
}

// DecodeTrieValue is a free data retrieval call binding the contract method 0xb9e59d09.
//
// Solidity: function decodeTrieValue(bytes trieValue) constant returns(uint256 v)
func (_EtherToken *EtherTokenSession) DecodeTrieValue(trieValue []byte) (*big.Int, error) {
	return _EtherToken.Contract.DecodeTrieValue(&_EtherToken.CallOpts, trieValue)
}

// DecodeTrieValue is a free data retrieval call binding the contract method 0xb9e59d09.
//
// Solidity: function decodeTrieValue(bytes trieValue) constant returns(uint256 v)
func (_EtherToken *EtherTokenCallerSession) DecodeTrieValue(trieValue []byte) (*big.Int, error) {
	return _EtherToken.Contract.DecodeTrieValue(&_EtherToken.CallOpts, trieValue)
}

// Development is a free data retrieval call binding the contract method 0x7b929c27.
//
// Solidity: function development() constant returns(bool)
func (_EtherToken *EtherTokenCaller) Development(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _EtherToken.contract.Call(opts, out, "development")
	return *ret0, err
}

// Development is a free data retrieval call binding the contract method 0x7b929c27.
//
// Solidity: function development() constant returns(bool)
func (_EtherToken *EtherTokenSession) Development() (bool, error) {
	return _EtherToken.Contract.Development(&_EtherToken.CallOpts)
}

// Development is a free data retrieval call binding the contract method 0x7b929c27.
//
// Solidity: function development() constant returns(bool)
func (_EtherToken *EtherTokenCallerSession) Development() (bool, error) {
	return _EtherToken.Contract.Development(&_EtherToken.CallOpts)
}

// GetBalanceTrieKey is a free data retrieval call binding the contract method 0xb18fcfdf.
//
// Solidity: function getBalanceTrieKey(address _who) constant returns(bytes32)
func (_EtherToken *EtherTokenCaller) GetBalanceTrieKey(opts *bind.CallOpts, _who common.Address) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _EtherToken.contract.Call(opts, out, "getBalanceTrieKey", _who)
	return *ret0, err
}

// GetBalanceTrieKey is a free data retrieval call binding the contract method 0xb18fcfdf.
//
// Solidity: function getBalanceTrieKey(address _who) constant returns(bytes32)
func (_EtherToken *EtherTokenSession) GetBalanceTrieKey(_who common.Address) ([32]byte, error) {
	return _EtherToken.Contract.GetBalanceTrieKey(&_EtherToken.CallOpts, _who)
}

// GetBalanceTrieKey is a free data retrieval call binding the contract method 0xb18fcfdf.
//
// Solidity: function getBalanceTrieKey(address _who) constant returns(bytes32)
func (_EtherToken *EtherTokenCallerSession) GetBalanceTrieKey(_who common.Address) ([32]byte, error) {
	return _EtherToken.Contract.GetBalanceTrieKey(&_EtherToken.CallOpts, _who)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() constant returns(bool)
func (_EtherToken *EtherTokenCaller) Initialized(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _EtherToken.contract.Call(opts, out, "initialized")
	return *ret0, err
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() constant returns(bool)
func (_EtherToken *EtherTokenSession) Initialized() (bool, error) {
	return _EtherToken.Contract.Initialized(&_EtherToken.CallOpts)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() constant returns(bool)
func (_EtherToken *EtherTokenCallerSession) Initialized() (bool, error) {
	return _EtherToken.Contract.Initialized(&_EtherToken.CallOpts)
}

// Rootchain is a free data retrieval call binding the contract method 0xbcc7874d.
//
// Solidity: function rootchain() constant returns(address)
func (_EtherToken *EtherTokenCaller) Rootchain(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _EtherToken.contract.Call(opts, out, "rootchain")
	return *ret0, err
}

// Rootchain is a free data retrieval call binding the contract method 0xbcc7874d.
//
// Solidity: function rootchain() constant returns(address)
func (_EtherToken *EtherTokenSession) Rootchain() (common.Address, error) {
	return _EtherToken.Contract.Rootchain(&_EtherToken.CallOpts)
}

// Rootchain is a free data retrieval call binding the contract method 0xbcc7874d.
//
// Solidity: function rootchain() constant returns(address)
func (_EtherToken *EtherTokenCallerSession) Rootchain() (common.Address, error) {
	return _EtherToken.Contract.Rootchain(&_EtherToken.CallOpts)
}

// SwapEnabled is a free data retrieval call binding the contract method 0x6ddd1713.
//
// Solidity: function swapEnabled() constant returns(bool)
func (_EtherToken *EtherTokenCaller) SwapEnabled(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _EtherToken.contract.Call(opts, out, "swapEnabled")
	return *ret0, err
}

// SwapEnabled is a free data retrieval call binding the contract method 0x6ddd1713.
//
// Solidity: function swapEnabled() constant returns(bool)
func (_EtherToken *EtherTokenSession) SwapEnabled() (bool, error) {
	return _EtherToken.Contract.SwapEnabled(&_EtherToken.CallOpts)
}

// SwapEnabled is a free data retrieval call binding the contract method 0x6ddd1713.
//
// Solidity: function swapEnabled() constant returns(bool)
func (_EtherToken *EtherTokenCallerSession) SwapEnabled() (bool, error) {
	return _EtherToken.Contract.SwapEnabled(&_EtherToken.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_EtherToken *EtherTokenCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _EtherToken.contract.Call(opts, out, "token")
	return *ret0, err
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_EtherToken *EtherTokenSession) Token() (common.Address, error) {
	return _EtherToken.Contract.Token(&_EtherToken.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_EtherToken *EtherTokenCallerSession) Token() (common.Address, error) {
	return _EtherToken.Contract.Token(&_EtherToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_EtherToken *EtherTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _EtherToken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_EtherToken *EtherTokenSession) TotalSupply() (*big.Int, error) {
	return _EtherToken.Contract.TotalSupply(&_EtherToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_EtherToken *EtherTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _EtherToken.Contract.TotalSupply(&_EtherToken.CallOpts)
}

// ApplyRequestInChildChain is a paid mutator transaction binding the contract method 0x141ecf46.
//
// Solidity: function applyRequestInChildChain(bool isExit, uint256 requestId, address requestor, bytes32 trieKey, bytes trieValue) returns(bool success)
func (_EtherToken *EtherTokenTransactor) ApplyRequestInChildChain(opts *bind.TransactOpts, isExit bool, requestId *big.Int, requestor common.Address, trieKey [32]byte, trieValue []byte) (*types.Transaction, error) {
	return _EtherToken.contract.Transact(opts, "applyRequestInChildChain", isExit, requestId, requestor, trieKey, trieValue)
}

// ApplyRequestInChildChain is a paid mutator transaction binding the contract method 0x141ecf46.
//
// Solidity: function applyRequestInChildChain(bool isExit, uint256 requestId, address requestor, bytes32 trieKey, bytes trieValue) returns(bool success)
func (_EtherToken *EtherTokenSession) ApplyRequestInChildChain(isExit bool, requestId *big.Int, requestor common.Address, trieKey [32]byte, trieValue []byte) (*types.Transaction, error) {
	return _EtherToken.Contract.ApplyRequestInChildChain(&_EtherToken.TransactOpts, isExit, requestId, requestor, trieKey, trieValue)
}

// ApplyRequestInChildChain is a paid mutator transaction binding the contract method 0x141ecf46.
//
// Solidity: function applyRequestInChildChain(bool isExit, uint256 requestId, address requestor, bytes32 trieKey, bytes trieValue) returns(bool success)
func (_EtherToken *EtherTokenTransactorSession) ApplyRequestInChildChain(isExit bool, requestId *big.Int, requestor common.Address, trieKey [32]byte, trieValue []byte) (*types.Transaction, error) {
	return _EtherToken.Contract.ApplyRequestInChildChain(&_EtherToken.TransactOpts, isExit, requestId, requestor, trieKey, trieValue)
}

// ApplyRequestInRootChain is a paid mutator transaction binding the contract method 0xa9f79308.
//
// Solidity: function applyRequestInRootChain(bool isExit, uint256 requestId, address requestor, bytes32 trieKey, bytes trieValue) returns(bool success)
func (_EtherToken *EtherTokenTransactor) ApplyRequestInRootChain(opts *bind.TransactOpts, isExit bool, requestId *big.Int, requestor common.Address, trieKey [32]byte, trieValue []byte) (*types.Transaction, error) {
	return _EtherToken.contract.Transact(opts, "applyRequestInRootChain", isExit, requestId, requestor, trieKey, trieValue)
}

// ApplyRequestInRootChain is a paid mutator transaction binding the contract method 0xa9f79308.
//
// Solidity: function applyRequestInRootChain(bool isExit, uint256 requestId, address requestor, bytes32 trieKey, bytes trieValue) returns(bool success)
func (_EtherToken *EtherTokenSession) ApplyRequestInRootChain(isExit bool, requestId *big.Int, requestor common.Address, trieKey [32]byte, trieValue []byte) (*types.Transaction, error) {
	return _EtherToken.Contract.ApplyRequestInRootChain(&_EtherToken.TransactOpts, isExit, requestId, requestor, trieKey, trieValue)
}

// ApplyRequestInRootChain is a paid mutator transaction binding the contract method 0xa9f79308.
//
// Solidity: function applyRequestInRootChain(bool isExit, uint256 requestId, address requestor, bytes32 trieKey, bytes trieValue) returns(bool success)
func (_EtherToken *EtherTokenTransactorSession) ApplyRequestInRootChain(isExit bool, requestId *big.Int, requestor common.Address, trieKey [32]byte, trieValue []byte) (*types.Transaction, error) {
	return _EtherToken.Contract.ApplyRequestInRootChain(&_EtherToken.TransactOpts, isExit, requestId, requestor, trieKey, trieValue)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_EtherToken *EtherTokenTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _EtherToken.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_EtherToken *EtherTokenSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _EtherToken.Contract.Approve(&_EtherToken.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_EtherToken *EtherTokenTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _EtherToken.Contract.Approve(&_EtherToken.TransactOpts, _spender, _value)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(address _spender, uint256 _subtractedValue) returns(bool)
func (_EtherToken *EtherTokenTransactor) DecreaseApproval(opts *bind.TransactOpts, _spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _EtherToken.contract.Transact(opts, "decreaseApproval", _spender, _subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(address _spender, uint256 _subtractedValue) returns(bool)
func (_EtherToken *EtherTokenSession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _EtherToken.Contract.DecreaseApproval(&_EtherToken.TransactOpts, _spender, _subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(address _spender, uint256 _subtractedValue) returns(bool)
func (_EtherToken *EtherTokenTransactorSession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _EtherToken.Contract.DecreaseApproval(&_EtherToken.TransactOpts, _spender, _subtractedValue)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _amount) returns(bool)
func (_EtherToken *EtherTokenTransactor) Deposit(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _EtherToken.contract.Transact(opts, "deposit", _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _amount) returns(bool)
func (_EtherToken *EtherTokenSession) Deposit(_amount *big.Int) (*types.Transaction, error) {
	return _EtherToken.Contract.Deposit(&_EtherToken.TransactOpts, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _amount) returns(bool)
func (_EtherToken *EtherTokenTransactorSession) Deposit(_amount *big.Int) (*types.Transaction, error) {
	return _EtherToken.Contract.Deposit(&_EtherToken.TransactOpts, _amount)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(address _spender, uint256 _addedValue) returns(bool)
func (_EtherToken *EtherTokenTransactor) IncreaseApproval(opts *bind.TransactOpts, _spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _EtherToken.contract.Transact(opts, "increaseApproval", _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(address _spender, uint256 _addedValue) returns(bool)
func (_EtherToken *EtherTokenSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _EtherToken.Contract.IncreaseApproval(&_EtherToken.TransactOpts, _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(address _spender, uint256 _addedValue) returns(bool)
func (_EtherToken *EtherTokenTransactorSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _EtherToken.Contract.IncreaseApproval(&_EtherToken.TransactOpts, _spender, _addedValue)
}

// Init is a paid mutator transaction binding the contract method 0x19ab453c.
//
// Solidity: function init(address _rootchain) returns(bool)
func (_EtherToken *EtherTokenTransactor) Init(opts *bind.TransactOpts, _rootchain common.Address) (*types.Transaction, error) {
	return _EtherToken.contract.Transact(opts, "init", _rootchain)
}

// Init is a paid mutator transaction binding the contract method 0x19ab453c.
//
// Solidity: function init(address _rootchain) returns(bool)
func (_EtherToken *EtherTokenSession) Init(_rootchain common.Address) (*types.Transaction, error) {
	return _EtherToken.Contract.Init(&_EtherToken.TransactOpts, _rootchain)
}

// Init is a paid mutator transaction binding the contract method 0x19ab453c.
//
// Solidity: function init(address _rootchain) returns(bool)
func (_EtherToken *EtherTokenTransactorSession) Init(_rootchain common.Address) (*types.Transaction, error) {
	return _EtherToken.Contract.Init(&_EtherToken.TransactOpts, _rootchain)
}

// SwapFromEth is a paid mutator transaction binding the contract method 0x59310da7.
//
// Solidity: function swapFromEth() returns()
func (_EtherToken *EtherTokenTransactor) SwapFromEth(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EtherToken.contract.Transact(opts, "swapFromEth")
}

// SwapFromEth is a paid mutator transaction binding the contract method 0x59310da7.
//
// Solidity: function swapFromEth() returns()
func (_EtherToken *EtherTokenSession) SwapFromEth() (*types.Transaction, error) {
	return _EtherToken.Contract.SwapFromEth(&_EtherToken.TransactOpts)
}

// SwapFromEth is a paid mutator transaction binding the contract method 0x59310da7.
//
// Solidity: function swapFromEth() returns()
func (_EtherToken *EtherTokenTransactorSession) SwapFromEth() (*types.Transaction, error) {
	return _EtherToken.Contract.SwapFromEth(&_EtherToken.TransactOpts)
}

// SwapToEth is a paid mutator transaction binding the contract method 0xb06c6729.
//
// Solidity: function swapToEth(uint256 _amount) returns()
func (_EtherToken *EtherTokenTransactor) SwapToEth(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _EtherToken.contract.Transact(opts, "swapToEth", _amount)
}

// SwapToEth is a paid mutator transaction binding the contract method 0xb06c6729.
//
// Solidity: function swapToEth(uint256 _amount) returns()
func (_EtherToken *EtherTokenSession) SwapToEth(_amount *big.Int) (*types.Transaction, error) {
	return _EtherToken.Contract.SwapToEth(&_EtherToken.TransactOpts, _amount)
}

// SwapToEth is a paid mutator transaction binding the contract method 0xb06c6729.
//
// Solidity: function swapToEth(uint256 _amount) returns()
func (_EtherToken *EtherTokenTransactorSession) SwapToEth(_amount *big.Int) (*types.Transaction, error) {
	return _EtherToken.Contract.SwapToEth(&_EtherToken.TransactOpts, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_EtherToken *EtherTokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _EtherToken.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_EtherToken *EtherTokenSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _EtherToken.Contract.Transfer(&_EtherToken.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_EtherToken *EtherTokenTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _EtherToken.Contract.Transfer(&_EtherToken.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_EtherToken *EtherTokenTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _EtherToken.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_EtherToken *EtherTokenSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _EtherToken.Contract.TransferFrom(&_EtherToken.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_EtherToken *EtherTokenTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _EtherToken.Contract.TransferFrom(&_EtherToken.TransactOpts, _from, _to, _value)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amount) returns(bool)
func (_EtherToken *EtherTokenTransactor) Withdraw(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _EtherToken.contract.Transact(opts, "withdraw", _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amount) returns(bool)
func (_EtherToken *EtherTokenSession) Withdraw(_amount *big.Int) (*types.Transaction, error) {
	return _EtherToken.Contract.Withdraw(&_EtherToken.TransactOpts, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amount) returns(bool)
func (_EtherToken *EtherTokenTransactorSession) Withdraw(_amount *big.Int) (*types.Transaction, error) {
	return _EtherToken.Contract.Withdraw(&_EtherToken.TransactOpts, _amount)
}

// EtherTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the EtherToken contract.
type EtherTokenApprovalIterator struct {
	Event *EtherTokenApproval // Event containing the contract specifics and raw log

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
func (it *EtherTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherTokenApproval)
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
		it.Event = new(EtherTokenApproval)
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
func (it *EtherTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherTokenApproval represents a Approval event raised by the EtherToken contract.
type EtherTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_EtherToken *EtherTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*EtherTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _EtherToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &EtherTokenApprovalIterator{contract: _EtherToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_EtherToken *EtherTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *EtherTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _EtherToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherTokenApproval)
				if err := _EtherToken.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_EtherToken *EtherTokenFilterer) ParseApproval(log types.Log) (*EtherTokenApproval, error) {
	event := new(EtherTokenApproval)
	if err := _EtherToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// EtherTokenDeposittedIterator is returned from FilterDepositted and is used to iterate over the raw logs and unpacked data for Depositted events raised by the EtherToken contract.
type EtherTokenDeposittedIterator struct {
	Event *EtherTokenDepositted // Event containing the contract specifics and raw log

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
func (it *EtherTokenDeposittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherTokenDepositted)
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
		it.Event = new(EtherTokenDepositted)
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
func (it *EtherTokenDeposittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherTokenDeposittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherTokenDepositted represents a Depositted event raised by the EtherToken contract.
type EtherTokenDepositted struct {
	From  common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterDepositted is a free log retrieval operation binding the contract event 0x46fc0825df492a7a7f1fb1b690f6912ecb69575484fbb24be0f398ae70d97db4.
//
// Solidity: event Depositted(address _from, uint256 _value)
func (_EtherToken *EtherTokenFilterer) FilterDepositted(opts *bind.FilterOpts) (*EtherTokenDeposittedIterator, error) {

	logs, sub, err := _EtherToken.contract.FilterLogs(opts, "Depositted")
	if err != nil {
		return nil, err
	}
	return &EtherTokenDeposittedIterator{contract: _EtherToken.contract, event: "Depositted", logs: logs, sub: sub}, nil
}

// WatchDepositted is a free log subscription operation binding the contract event 0x46fc0825df492a7a7f1fb1b690f6912ecb69575484fbb24be0f398ae70d97db4.
//
// Solidity: event Depositted(address _from, uint256 _value)
func (_EtherToken *EtherTokenFilterer) WatchDepositted(opts *bind.WatchOpts, sink chan<- *EtherTokenDepositted) (event.Subscription, error) {

	logs, sub, err := _EtherToken.contract.WatchLogs(opts, "Depositted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherTokenDepositted)
				if err := _EtherToken.contract.UnpackLog(event, "Depositted", log); err != nil {
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

// ParseDepositted is a log parse operation binding the contract event 0x46fc0825df492a7a7f1fb1b690f6912ecb69575484fbb24be0f398ae70d97db4.
//
// Solidity: event Depositted(address _from, uint256 _value)
func (_EtherToken *EtherTokenFilterer) ParseDepositted(log types.Log) (*EtherTokenDepositted, error) {
	event := new(EtherTokenDepositted)
	if err := _EtherToken.contract.UnpackLog(event, "Depositted", log); err != nil {
		return nil, err
	}
	return event, nil
}

// EtherTokenRequestCreatedIterator is returned from FilterRequestCreated and is used to iterate over the raw logs and unpacked data for RequestCreated events raised by the EtherToken contract.
type EtherTokenRequestCreatedIterator struct {
	Event *EtherTokenRequestCreated // Event containing the contract specifics and raw log

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
func (it *EtherTokenRequestCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherTokenRequestCreated)
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
		it.Event = new(EtherTokenRequestCreated)
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
func (it *EtherTokenRequestCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherTokenRequestCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherTokenRequestCreated represents a RequestCreated event raised by the EtherToken contract.
type EtherTokenRequestCreated struct {
	IsExit    bool
	Requestor common.Address
	TrieKey   [32]byte
	Value     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRequestCreated is a free log retrieval operation binding the contract event 0x4b0c23cd90bac23cdaa0958fbd0306f53e20cd5e65676d21b9ceb109784e15b8.
//
// Solidity: event RequestCreated(bool _isExit, address _requestor, bytes32 _trieKey, uint256 _value)
func (_EtherToken *EtherTokenFilterer) FilterRequestCreated(opts *bind.FilterOpts) (*EtherTokenRequestCreatedIterator, error) {

	logs, sub, err := _EtherToken.contract.FilterLogs(opts, "RequestCreated")
	if err != nil {
		return nil, err
	}
	return &EtherTokenRequestCreatedIterator{contract: _EtherToken.contract, event: "RequestCreated", logs: logs, sub: sub}, nil
}

// WatchRequestCreated is a free log subscription operation binding the contract event 0x4b0c23cd90bac23cdaa0958fbd0306f53e20cd5e65676d21b9ceb109784e15b8.
//
// Solidity: event RequestCreated(bool _isExit, address _requestor, bytes32 _trieKey, uint256 _value)
func (_EtherToken *EtherTokenFilterer) WatchRequestCreated(opts *bind.WatchOpts, sink chan<- *EtherTokenRequestCreated) (event.Subscription, error) {

	logs, sub, err := _EtherToken.contract.WatchLogs(opts, "RequestCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherTokenRequestCreated)
				if err := _EtherToken.contract.UnpackLog(event, "RequestCreated", log); err != nil {
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

// ParseRequestCreated is a log parse operation binding the contract event 0x4b0c23cd90bac23cdaa0958fbd0306f53e20cd5e65676d21b9ceb109784e15b8.
//
// Solidity: event RequestCreated(bool _isExit, address _requestor, bytes32 _trieKey, uint256 _value)
func (_EtherToken *EtherTokenFilterer) ParseRequestCreated(log types.Log) (*EtherTokenRequestCreated, error) {
	event := new(EtherTokenRequestCreated)
	if err := _EtherToken.contract.UnpackLog(event, "RequestCreated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// EtherTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the EtherToken contract.
type EtherTokenTransferIterator struct {
	Event *EtherTokenTransfer // Event containing the contract specifics and raw log

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
func (it *EtherTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherTokenTransfer)
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
		it.Event = new(EtherTokenTransfer)
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
func (it *EtherTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherTokenTransfer represents a Transfer event raised by the EtherToken contract.
type EtherTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_EtherToken *EtherTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*EtherTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EtherToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &EtherTokenTransferIterator{contract: _EtherToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_EtherToken *EtherTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *EtherTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EtherToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherTokenTransfer)
				if err := _EtherToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_EtherToken *EtherTokenFilterer) ParseTransfer(log types.Log) (*EtherTokenTransfer, error) {
	event := new(EtherTokenTransfer)
	if err := _EtherToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// EtherTokenWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the EtherToken contract.
type EtherTokenWithdrawnIterator struct {
	Event *EtherTokenWithdrawn // Event containing the contract specifics and raw log

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
func (it *EtherTokenWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherTokenWithdrawn)
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
		it.Event = new(EtherTokenWithdrawn)
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
func (it *EtherTokenWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherTokenWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherTokenWithdrawn represents a Withdrawn event raised by the EtherToken contract.
type EtherTokenWithdrawn struct {
	From  common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address _from, uint256 _value)
func (_EtherToken *EtherTokenFilterer) FilterWithdrawn(opts *bind.FilterOpts) (*EtherTokenWithdrawnIterator, error) {

	logs, sub, err := _EtherToken.contract.FilterLogs(opts, "Withdrawn")
	if err != nil {
		return nil, err
	}
	return &EtherTokenWithdrawnIterator{contract: _EtherToken.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address _from, uint256 _value)
func (_EtherToken *EtherTokenFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *EtherTokenWithdrawn) (event.Subscription, error) {

	logs, sub, err := _EtherToken.contract.WatchLogs(opts, "Withdrawn")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherTokenWithdrawn)
				if err := _EtherToken.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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

// ParseWithdrawn is a log parse operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address _from, uint256 _value)
func (_EtherToken *EtherTokenFilterer) ParseWithdrawn(log types.Log) (*EtherTokenWithdrawn, error) {
	event := new(EtherTokenWithdrawn)
	if err := _EtherToken.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RequestableERC20WrapperABI is the input ABI used to generate the binding from.
const RequestableERC20WrapperABI = "[{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_development\",\"type\":\"bool\"},{\"internalType\":\"contractERC20\",\"name\":\"_token\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Depositted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"_isExit\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_requestor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_trieKey\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"RequestCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"isExit\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"requestor\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"trieKey\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"trieValue\",\"type\":\"bytes\"}],\"name\":\"applyRequestInChildChain\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"isExit\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"requestor\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"trieKey\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"trieValue\",\"type\":\"bytes\"}],\"name\":\"applyRequestInRootChain\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"trieValue\",\"type\":\"bytes\"}],\"name\":\"decodeTrieValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"v\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseApproval\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"development\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_who\",\"type\":\"address\"}],\"name\":\"getBalanceTrieKey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseApproval\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rootchain\",\"type\":\"address\"}],\"name\":\"init\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"initialized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rootchain\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractERC20\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// RequestableERC20WrapperFuncSigs maps the 4-byte function signature to its string representation.
var RequestableERC20WrapperFuncSigs = map[string]string{
	"dd62ed3e": "allowance(address,address)",
	"141ecf46": "applyRequestInChildChain(bool,uint256,address,bytes32,bytes)",
	"a9f79308": "applyRequestInRootChain(bool,uint256,address,bytes32,bytes)",
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"b9e59d09": "decodeTrieValue(bytes)",
	"66188463": "decreaseApproval(address,uint256)",
	"b6b55f25": "deposit(uint256)",
	"7b929c27": "development()",
	"b18fcfdf": "getBalanceTrieKey(address)",
	"d73dd623": "increaseApproval(address,uint256)",
	"19ab453c": "init(address)",
	"158ef93e": "initialized()",
	"bcc7874d": "rootchain()",
	"fc0c546a": "token()",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
	"2e1a7d4d": "withdraw(uint256)",
}

// RequestableERC20WrapperBin is the compiled bytecode used for deploying new contracts.
var RequestableERC20WrapperBin = "0x608060405234801561001057600080fd5b50604051610f99380380610f998339818101604052604081101561003357600080fd5b5080516020909101516003805461ff0019166101009315159390930292909217909155600480546001600160a01b0319166001600160a01b03909216919091179055610f15806100846000396000f3fe608060405234801561001057600080fd5b50600436106101215760003560e01c80637b929c27116100ad578063b9e59d0911610071578063b9e59d09146103f0578063bcc7874d14610496578063d73dd623146104ba578063dd62ed3e146104e6578063fc0c546a1461051457610121565b80637b929c27146102e6578063a9059cbb146102ee578063a9f793081461031a578063b18fcfdf146103ad578063b6b55f25146103d357610121565b806319ab453c116100f457806319ab453c1461021b57806323b872dd146102415780632e1a7d4d14610277578063661884631461029457806370a08231146102c057610121565b8063095ea7b314610126578063141ecf4614610166578063158ef93e146101f957806318160ddd14610201575b600080fd5b6101526004803603604081101561013c57600080fd5b506001600160a01b03813516906020013561051c565b604080519115158252519081900360200190f35b610152600480360360a081101561017c57600080fd5b81351515916020810135916001600160a01b036040830135169160608101359181019060a0810160808201356401000000008111156101ba57600080fd5b8201836020820111156101cc57600080fd5b803590602001918460018302840111640100000000831117156101ee57600080fd5b509092509050610582565b610152610677565b610209610680565b60408051918252519081900360200190f35b6101526004803603602081101561023157600080fd5b50356001600160a01b0316610686565b6101526004803603606081101561025757600080fd5b506001600160a01b038135811691602081013590911690604001356106cc565b6101526004803603602081101561028d57600080fd5b503561082d565b610152600480360360408110156102aa57600080fd5b506001600160a01b038135169060200135610916565b610209600480360360208110156102d657600080fd5b50356001600160a01b0316610a05565b610152610a20565b6101526004803603604081101561030457600080fd5b506001600160a01b038135169060200135610a2e565b610152600480360360a081101561033057600080fd5b81351515916020810135916001600160a01b036040830135169160608101359181019060a08101608082013564010000000081111561036e57600080fd5b82018360208201111561038057600080fd5b803590602001918460018302840111640100000000831117156103a257600080fd5b509092509050610af9565b610209600480360360208110156103c357600080fd5b50356001600160a01b0316610b99565b610152600480360360208110156103e957600080fd5b5035610bdb565b6102096004803603602081101561040657600080fd5b81019060208101813564010000000081111561042157600080fd5b82018360208201111561043357600080fd5b8035906020019184600183028401116401000000008311171561045557600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610c8d945050505050565b61049e610ca5565b604080516001600160a01b039092168252519081900360200190f35b610152600480360360408110156104d057600080fd5b506001600160a01b038135169060200135610cba565b610209600480360360408110156104fc57600080fd5b506001600160a01b0381358116916020013516610d53565b61049e610d7e565b3360008181526002602090815260408083206001600160a01b038716808552908352818420869055815186815291519394909390927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925928290030190a350600192915050565b600354600090610100900460ff1680610599575033155b6105a257600080fd5b6105ab85610b99565b84146105b657600080fd5b60006105f784848080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250610c8d92505050565b9050871561060e576106098682610d8d565b610618565b6106188682610e18565b6040805189151581526001600160a01b03881660208201528082018790526060810183905290517f4b0c23cd90bac23cdaa0958fbd0306f53e20cd5e65676d21b9ceb109784e15b89181900360800190a1506001979650505050505050565b60035460ff1681565b60015490565b60035460009060ff161561069957600080fd5b6003805460ff196001600160a01b03909416620100000262010000600160b01b0319909116179290921660011790915590565b6001600160a01b0383166000908152602081905260408120548211156106f157600080fd5b6001600160a01b038416600090815260026020908152604080832033845290915290205482111561072157600080fd5b6001600160a01b03831661073457600080fd5b6001600160a01b03841660009081526020819052604090205461075d908363ffffffff610e9b16565b6001600160a01b038086166000908152602081905260408082209390935590851681522054610792908363ffffffff610ead16565b6001600160a01b038085166000908152602081815260408083209490945591871681526002825282812033825290915220546107d4908363ffffffff610e9b16565b6001600160a01b0380861660008181526002602090815260408083203384528252918290209490945580518681529051928716939192600080516020610ec1833981519152929181900390910190a35060019392505050565b60035460009060ff1661083f57600080fd5b6108493383610d8d565b604080513381526020810184905281517f7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5929181900390910190a1600480546040805163a9059cbb60e01b8152339381019390935260248301859052516001600160a01b039091169163a9059cbb9160448083019260209291908290030181600087803b1580156108d957600080fd5b505af11580156108ed573d6000803e3d6000fd5b505050506040513d602081101561090357600080fd5b505161090e57600080fd5b506001919050565b3360009081526002602090815260408083206001600160a01b038616845290915281205480831061096a573360009081526002602090815260408083206001600160a01b038816845290915281205561099f565b61097a818463ffffffff610e9b16565b3360009081526002602090815260408083206001600160a01b03891684529091529020555b3360008181526002602090815260408083206001600160a01b0389168085529083529281902054815190815290519293927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929181900390910190a35060019392505050565b6001600160a01b031660009081526020819052604090205490565b600354610100900460ff1681565b33600090815260208190526040812054821115610a4a57600080fd5b6001600160a01b038316610a5d57600080fd5b33600090815260208190526040902054610a7d908363ffffffff610e9b16565b33600090815260208190526040808220929092556001600160a01b03851681522054610aaf908363ffffffff610ead16565b6001600160a01b03841660008181526020818152604091829020939093558051858152905191923392600080516020610ec18339815191529281900390910190a350600192915050565b60035460009060ff16610b0b57600080fd5b6003546201000090046001600160a01b03163314610b2857600080fd5b610b3185610b99565b8414610b3c57600080fd5b6000610b7d84848080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250610c8d92505050565b90508715610b8f576106098682610e18565b6106188682610d8d565b60408051600060208083019190915260609390931b6bffffffffffffffffffffffff191681830152815180820360340181526054909101909152805191012090565b60035460009060ff16610bed57600080fd5b610bf73383610e18565b604080513381526020810184905281517f46fc0825df492a7a7f1fb1b690f6912ecb69575484fbb24be0f398ae70d97db4929181900390910190a160048054604080516323b872dd60e01b8152339381019390935230602484015260448301859052516001600160a01b03909116916323b872dd9160648083019260209291908290030181600087803b1580156108d957600080fd5b60008151602014610c9d57600080fd5b506020015190565b6003546201000090046001600160a01b031681565b3360009081526002602090815260408083206001600160a01b0386168452909152812054610cee908363ffffffff610ead16565b3360008181526002602090815260408083206001600160a01b0389168085529083529281902085905580519485525191937f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929081900390910190a350600192915050565b6001600160a01b03918216600090815260026020908152604080832093909416825291909152205490565b6004546001600160a01b031681565b6001600160a01b038216600090815260208190526040902054610db6908263ffffffff610e9b16565b6001600160a01b038316600090815260208190526040902055600154610de2908263ffffffff610e9b16565b6001556040805182815290516000916001600160a01b03851691600080516020610ec18339815191529181900360200190a35050565b600154610e2b908263ffffffff610ead16565b6001556001600160a01b038216600090815260208190526040902054610e57908263ffffffff610ead16565b6001600160a01b038316600081815260208181526040808320949094558351858152935192939192600080516020610ec18339815191529281900390910190a35050565b600082821115610ea757fe5b50900390565b81810182811015610eba57fe5b9291505056feddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3efa265627a7a72315820cf7f97c42d1f68f0429bad58c03e7bcc858d14f5ceab5a78e7f5e8dbe5433f9264736f6c634300050c0032"

// DeployRequestableERC20Wrapper deploys a new Ethereum contract, binding an instance of RequestableERC20Wrapper to it.
func DeployRequestableERC20Wrapper(auth *bind.TransactOpts, backend bind.ContractBackend, _development bool, _token common.Address) (common.Address, *types.Transaction, *RequestableERC20Wrapper, error) {
	parsed, err := abi.JSON(strings.NewReader(RequestableERC20WrapperABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RequestableERC20WrapperBin), backend, _development, _token)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RequestableERC20Wrapper{RequestableERC20WrapperCaller: RequestableERC20WrapperCaller{contract: contract}, RequestableERC20WrapperTransactor: RequestableERC20WrapperTransactor{contract: contract}, RequestableERC20WrapperFilterer: RequestableERC20WrapperFilterer{contract: contract}}, nil
}

// RequestableERC20Wrapper is an auto generated Go binding around an Ethereum contract.
type RequestableERC20Wrapper struct {
	RequestableERC20WrapperCaller     // Read-only binding to the contract
	RequestableERC20WrapperTransactor // Write-only binding to the contract
	RequestableERC20WrapperFilterer   // Log filterer for contract events
}

// RequestableERC20WrapperCaller is an auto generated read-only Go binding around an Ethereum contract.
type RequestableERC20WrapperCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RequestableERC20WrapperTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RequestableERC20WrapperTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RequestableERC20WrapperFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RequestableERC20WrapperFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RequestableERC20WrapperSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RequestableERC20WrapperSession struct {
	Contract     *RequestableERC20Wrapper // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// RequestableERC20WrapperCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RequestableERC20WrapperCallerSession struct {
	Contract *RequestableERC20WrapperCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// RequestableERC20WrapperTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RequestableERC20WrapperTransactorSession struct {
	Contract     *RequestableERC20WrapperTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// RequestableERC20WrapperRaw is an auto generated low-level Go binding around an Ethereum contract.
type RequestableERC20WrapperRaw struct {
	Contract *RequestableERC20Wrapper // Generic contract binding to access the raw methods on
}

// RequestableERC20WrapperCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RequestableERC20WrapperCallerRaw struct {
	Contract *RequestableERC20WrapperCaller // Generic read-only contract binding to access the raw methods on
}

// RequestableERC20WrapperTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RequestableERC20WrapperTransactorRaw struct {
	Contract *RequestableERC20WrapperTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRequestableERC20Wrapper creates a new instance of RequestableERC20Wrapper, bound to a specific deployed contract.
func NewRequestableERC20Wrapper(address common.Address, backend bind.ContractBackend) (*RequestableERC20Wrapper, error) {
	contract, err := bindRequestableERC20Wrapper(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RequestableERC20Wrapper{RequestableERC20WrapperCaller: RequestableERC20WrapperCaller{contract: contract}, RequestableERC20WrapperTransactor: RequestableERC20WrapperTransactor{contract: contract}, RequestableERC20WrapperFilterer: RequestableERC20WrapperFilterer{contract: contract}}, nil
}

// NewRequestableERC20WrapperCaller creates a new read-only instance of RequestableERC20Wrapper, bound to a specific deployed contract.
func NewRequestableERC20WrapperCaller(address common.Address, caller bind.ContractCaller) (*RequestableERC20WrapperCaller, error) {
	contract, err := bindRequestableERC20Wrapper(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RequestableERC20WrapperCaller{contract: contract}, nil
}

// NewRequestableERC20WrapperTransactor creates a new write-only instance of RequestableERC20Wrapper, bound to a specific deployed contract.
func NewRequestableERC20WrapperTransactor(address common.Address, transactor bind.ContractTransactor) (*RequestableERC20WrapperTransactor, error) {
	contract, err := bindRequestableERC20Wrapper(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RequestableERC20WrapperTransactor{contract: contract}, nil
}

// NewRequestableERC20WrapperFilterer creates a new log filterer instance of RequestableERC20Wrapper, bound to a specific deployed contract.
func NewRequestableERC20WrapperFilterer(address common.Address, filterer bind.ContractFilterer) (*RequestableERC20WrapperFilterer, error) {
	contract, err := bindRequestableERC20Wrapper(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RequestableERC20WrapperFilterer{contract: contract}, nil
}

// bindRequestableERC20Wrapper binds a generic wrapper to an already deployed contract.
func bindRequestableERC20Wrapper(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RequestableERC20WrapperABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RequestableERC20Wrapper *RequestableERC20WrapperRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RequestableERC20Wrapper.Contract.RequestableERC20WrapperCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RequestableERC20Wrapper *RequestableERC20WrapperRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RequestableERC20Wrapper.Contract.RequestableERC20WrapperTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RequestableERC20Wrapper *RequestableERC20WrapperRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RequestableERC20Wrapper.Contract.RequestableERC20WrapperTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RequestableERC20Wrapper *RequestableERC20WrapperCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RequestableERC20Wrapper.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RequestableERC20Wrapper *RequestableERC20WrapperTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RequestableERC20Wrapper.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RequestableERC20Wrapper *RequestableERC20WrapperTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RequestableERC20Wrapper.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) constant returns(uint256)
func (_RequestableERC20Wrapper *RequestableERC20WrapperCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RequestableERC20Wrapper.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) constant returns(uint256)
func (_RequestableERC20Wrapper *RequestableERC20WrapperSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _RequestableERC20Wrapper.Contract.Allowance(&_RequestableERC20Wrapper.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) constant returns(uint256)
func (_RequestableERC20Wrapper *RequestableERC20WrapperCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _RequestableERC20Wrapper.Contract.Allowance(&_RequestableERC20Wrapper.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) constant returns(uint256)
func (_RequestableERC20Wrapper *RequestableERC20WrapperCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RequestableERC20Wrapper.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) constant returns(uint256)
func (_RequestableERC20Wrapper *RequestableERC20WrapperSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _RequestableERC20Wrapper.Contract.BalanceOf(&_RequestableERC20Wrapper.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) constant returns(uint256)
func (_RequestableERC20Wrapper *RequestableERC20WrapperCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _RequestableERC20Wrapper.Contract.BalanceOf(&_RequestableERC20Wrapper.CallOpts, _owner)
}

// DecodeTrieValue is a free data retrieval call binding the contract method 0xb9e59d09.
//
// Solidity: function decodeTrieValue(bytes trieValue) constant returns(uint256 v)
func (_RequestableERC20Wrapper *RequestableERC20WrapperCaller) DecodeTrieValue(opts *bind.CallOpts, trieValue []byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RequestableERC20Wrapper.contract.Call(opts, out, "decodeTrieValue", trieValue)
	return *ret0, err
}

// DecodeTrieValue is a free data retrieval call binding the contract method 0xb9e59d09.
//
// Solidity: function decodeTrieValue(bytes trieValue) constant returns(uint256 v)
func (_RequestableERC20Wrapper *RequestableERC20WrapperSession) DecodeTrieValue(trieValue []byte) (*big.Int, error) {
	return _RequestableERC20Wrapper.Contract.DecodeTrieValue(&_RequestableERC20Wrapper.CallOpts, trieValue)
}

// DecodeTrieValue is a free data retrieval call binding the contract method 0xb9e59d09.
//
// Solidity: function decodeTrieValue(bytes trieValue) constant returns(uint256 v)
func (_RequestableERC20Wrapper *RequestableERC20WrapperCallerSession) DecodeTrieValue(trieValue []byte) (*big.Int, error) {
	return _RequestableERC20Wrapper.Contract.DecodeTrieValue(&_RequestableERC20Wrapper.CallOpts, trieValue)
}

// Development is a free data retrieval call binding the contract method 0x7b929c27.
//
// Solidity: function development() constant returns(bool)
func (_RequestableERC20Wrapper *RequestableERC20WrapperCaller) Development(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _RequestableERC20Wrapper.contract.Call(opts, out, "development")
	return *ret0, err
}

// Development is a free data retrieval call binding the contract method 0x7b929c27.
//
// Solidity: function development() constant returns(bool)
func (_RequestableERC20Wrapper *RequestableERC20WrapperSession) Development() (bool, error) {
	return _RequestableERC20Wrapper.Contract.Development(&_RequestableERC20Wrapper.CallOpts)
}

// Development is a free data retrieval call binding the contract method 0x7b929c27.
//
// Solidity: function development() constant returns(bool)
func (_RequestableERC20Wrapper *RequestableERC20WrapperCallerSession) Development() (bool, error) {
	return _RequestableERC20Wrapper.Contract.Development(&_RequestableERC20Wrapper.CallOpts)
}

// GetBalanceTrieKey is a free data retrieval call binding the contract method 0xb18fcfdf.
//
// Solidity: function getBalanceTrieKey(address _who) constant returns(bytes32)
func (_RequestableERC20Wrapper *RequestableERC20WrapperCaller) GetBalanceTrieKey(opts *bind.CallOpts, _who common.Address) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _RequestableERC20Wrapper.contract.Call(opts, out, "getBalanceTrieKey", _who)
	return *ret0, err
}

// GetBalanceTrieKey is a free data retrieval call binding the contract method 0xb18fcfdf.
//
// Solidity: function getBalanceTrieKey(address _who) constant returns(bytes32)
func (_RequestableERC20Wrapper *RequestableERC20WrapperSession) GetBalanceTrieKey(_who common.Address) ([32]byte, error) {
	return _RequestableERC20Wrapper.Contract.GetBalanceTrieKey(&_RequestableERC20Wrapper.CallOpts, _who)
}

// GetBalanceTrieKey is a free data retrieval call binding the contract method 0xb18fcfdf.
//
// Solidity: function getBalanceTrieKey(address _who) constant returns(bytes32)
func (_RequestableERC20Wrapper *RequestableERC20WrapperCallerSession) GetBalanceTrieKey(_who common.Address) ([32]byte, error) {
	return _RequestableERC20Wrapper.Contract.GetBalanceTrieKey(&_RequestableERC20Wrapper.CallOpts, _who)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() constant returns(bool)
func (_RequestableERC20Wrapper *RequestableERC20WrapperCaller) Initialized(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _RequestableERC20Wrapper.contract.Call(opts, out, "initialized")
	return *ret0, err
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() constant returns(bool)
func (_RequestableERC20Wrapper *RequestableERC20WrapperSession) Initialized() (bool, error) {
	return _RequestableERC20Wrapper.Contract.Initialized(&_RequestableERC20Wrapper.CallOpts)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() constant returns(bool)
func (_RequestableERC20Wrapper *RequestableERC20WrapperCallerSession) Initialized() (bool, error) {
	return _RequestableERC20Wrapper.Contract.Initialized(&_RequestableERC20Wrapper.CallOpts)
}

// Rootchain is a free data retrieval call binding the contract method 0xbcc7874d.
//
// Solidity: function rootchain() constant returns(address)
func (_RequestableERC20Wrapper *RequestableERC20WrapperCaller) Rootchain(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RequestableERC20Wrapper.contract.Call(opts, out, "rootchain")
	return *ret0, err
}

// Rootchain is a free data retrieval call binding the contract method 0xbcc7874d.
//
// Solidity: function rootchain() constant returns(address)
func (_RequestableERC20Wrapper *RequestableERC20WrapperSession) Rootchain() (common.Address, error) {
	return _RequestableERC20Wrapper.Contract.Rootchain(&_RequestableERC20Wrapper.CallOpts)
}

// Rootchain is a free data retrieval call binding the contract method 0xbcc7874d.
//
// Solidity: function rootchain() constant returns(address)
func (_RequestableERC20Wrapper *RequestableERC20WrapperCallerSession) Rootchain() (common.Address, error) {
	return _RequestableERC20Wrapper.Contract.Rootchain(&_RequestableERC20Wrapper.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_RequestableERC20Wrapper *RequestableERC20WrapperCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RequestableERC20Wrapper.contract.Call(opts, out, "token")
	return *ret0, err
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_RequestableERC20Wrapper *RequestableERC20WrapperSession) Token() (common.Address, error) {
	return _RequestableERC20Wrapper.Contract.Token(&_RequestableERC20Wrapper.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_RequestableERC20Wrapper *RequestableERC20WrapperCallerSession) Token() (common.Address, error) {
	return _RequestableERC20Wrapper.Contract.Token(&_RequestableERC20Wrapper.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_RequestableERC20Wrapper *RequestableERC20WrapperCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RequestableERC20Wrapper.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_RequestableERC20Wrapper *RequestableERC20WrapperSession) TotalSupply() (*big.Int, error) {
	return _RequestableERC20Wrapper.Contract.TotalSupply(&_RequestableERC20Wrapper.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_RequestableERC20Wrapper *RequestableERC20WrapperCallerSession) TotalSupply() (*big.Int, error) {
	return _RequestableERC20Wrapper.Contract.TotalSupply(&_RequestableERC20Wrapper.CallOpts)
}

// ApplyRequestInChildChain is a paid mutator transaction binding the contract method 0x141ecf46.
//
// Solidity: function applyRequestInChildChain(bool isExit, uint256 requestId, address requestor, bytes32 trieKey, bytes trieValue) returns(bool success)
func (_RequestableERC20Wrapper *RequestableERC20WrapperTransactor) ApplyRequestInChildChain(opts *bind.TransactOpts, isExit bool, requestId *big.Int, requestor common.Address, trieKey [32]byte, trieValue []byte) (*types.Transaction, error) {
	return _RequestableERC20Wrapper.contract.Transact(opts, "applyRequestInChildChain", isExit, requestId, requestor, trieKey, trieValue)
}

// ApplyRequestInChildChain is a paid mutator transaction binding the contract method 0x141ecf46.
//
// Solidity: function applyRequestInChildChain(bool isExit, uint256 requestId, address requestor, bytes32 trieKey, bytes trieValue) returns(bool success)
func (_RequestableERC20Wrapper *RequestableERC20WrapperSession) ApplyRequestInChildChain(isExit bool, requestId *big.Int, requestor common.Address, trieKey [32]byte, trieValue []byte) (*types.Transaction, error) {
	return _RequestableERC20Wrapper.Contract.ApplyRequestInChildChain(&_RequestableERC20Wrapper.TransactOpts, isExit, requestId, requestor, trieKey, trieValue)
}

// ApplyRequestInChildChain is a paid mutator transaction binding the contract method 0x141ecf46.
//
// Solidity: function applyRequestInChildChain(bool isExit, uint256 requestId, address requestor, bytes32 trieKey, bytes trieValue) returns(bool success)
func (_RequestableERC20Wrapper *RequestableERC20WrapperTransactorSession) ApplyRequestInChildChain(isExit bool, requestId *big.Int, requestor common.Address, trieKey [32]byte, trieValue []byte) (*types.Transaction, error) {
	return _RequestableERC20Wrapper.Contract.ApplyRequestInChildChain(&_RequestableERC20Wrapper.TransactOpts, isExit, requestId, requestor, trieKey, trieValue)
}

// ApplyRequestInRootChain is a paid mutator transaction binding the contract method 0xa9f79308.
//
// Solidity: function applyRequestInRootChain(bool isExit, uint256 requestId, address requestor, bytes32 trieKey, bytes trieValue) returns(bool success)
func (_RequestableERC20Wrapper *RequestableERC20WrapperTransactor) ApplyRequestInRootChain(opts *bind.TransactOpts, isExit bool, requestId *big.Int, requestor common.Address, trieKey [32]byte, trieValue []byte) (*types.Transaction, error) {
	return _RequestableERC20Wrapper.contract.Transact(opts, "applyRequestInRootChain", isExit, requestId, requestor, trieKey, trieValue)
}

// ApplyRequestInRootChain is a paid mutator transaction binding the contract method 0xa9f79308.
//
// Solidity: function applyRequestInRootChain(bool isExit, uint256 requestId, address requestor, bytes32 trieKey, bytes trieValue) returns(bool success)
func (_RequestableERC20Wrapper *RequestableERC20WrapperSession) ApplyRequestInRootChain(isExit bool, requestId *big.Int, requestor common.Address, trieKey [32]byte, trieValue []byte) (*types.Transaction, error) {
	return _RequestableERC20Wrapper.Contract.ApplyRequestInRootChain(&_RequestableERC20Wrapper.TransactOpts, isExit, requestId, requestor, trieKey, trieValue)
}

// ApplyRequestInRootChain is a paid mutator transaction binding the contract method 0xa9f79308.
//
// Solidity: function applyRequestInRootChain(bool isExit, uint256 requestId, address requestor, bytes32 trieKey, bytes trieValue) returns(bool success)
func (_RequestableERC20Wrapper *RequestableERC20WrapperTransactorSession) ApplyRequestInRootChain(isExit bool, requestId *big.Int, requestor common.Address, trieKey [32]byte, trieValue []byte) (*types.Transaction, error) {
	return _RequestableERC20Wrapper.Contract.ApplyRequestInRootChain(&_RequestableERC20Wrapper.TransactOpts, isExit, requestId, requestor, trieKey, trieValue)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_RequestableERC20Wrapper *RequestableERC20WrapperTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _RequestableERC20Wrapper.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_RequestableERC20Wrapper *RequestableERC20WrapperSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _RequestableERC20Wrapper.Contract.Approve(&_RequestableERC20Wrapper.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_RequestableERC20Wrapper *RequestableERC20WrapperTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _RequestableERC20Wrapper.Contract.Approve(&_RequestableERC20Wrapper.TransactOpts, _spender, _value)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(address _spender, uint256 _subtractedValue) returns(bool)
func (_RequestableERC20Wrapper *RequestableERC20WrapperTransactor) DecreaseApproval(opts *bind.TransactOpts, _spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _RequestableERC20Wrapper.contract.Transact(opts, "decreaseApproval", _spender, _subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(address _spender, uint256 _subtractedValue) returns(bool)
func (_RequestableERC20Wrapper *RequestableERC20WrapperSession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _RequestableERC20Wrapper.Contract.DecreaseApproval(&_RequestableERC20Wrapper.TransactOpts, _spender, _subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(address _spender, uint256 _subtractedValue) returns(bool)
func (_RequestableERC20Wrapper *RequestableERC20WrapperTransactorSession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _RequestableERC20Wrapper.Contract.DecreaseApproval(&_RequestableERC20Wrapper.TransactOpts, _spender, _subtractedValue)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _amount) returns(bool)
func (_RequestableERC20Wrapper *RequestableERC20WrapperTransactor) Deposit(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _RequestableERC20Wrapper.contract.Transact(opts, "deposit", _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _amount) returns(bool)
func (_RequestableERC20Wrapper *RequestableERC20WrapperSession) Deposit(_amount *big.Int) (*types.Transaction, error) {
	return _RequestableERC20Wrapper.Contract.Deposit(&_RequestableERC20Wrapper.TransactOpts, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _amount) returns(bool)
func (_RequestableERC20Wrapper *RequestableERC20WrapperTransactorSession) Deposit(_amount *big.Int) (*types.Transaction, error) {
	return _RequestableERC20Wrapper.Contract.Deposit(&_RequestableERC20Wrapper.TransactOpts, _amount)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(address _spender, uint256 _addedValue) returns(bool)
func (_RequestableERC20Wrapper *RequestableERC20WrapperTransactor) IncreaseApproval(opts *bind.TransactOpts, _spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _RequestableERC20Wrapper.contract.Transact(opts, "increaseApproval", _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(address _spender, uint256 _addedValue) returns(bool)
func (_RequestableERC20Wrapper *RequestableERC20WrapperSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _RequestableERC20Wrapper.Contract.IncreaseApproval(&_RequestableERC20Wrapper.TransactOpts, _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(address _spender, uint256 _addedValue) returns(bool)
func (_RequestableERC20Wrapper *RequestableERC20WrapperTransactorSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _RequestableERC20Wrapper.Contract.IncreaseApproval(&_RequestableERC20Wrapper.TransactOpts, _spender, _addedValue)
}

// Init is a paid mutator transaction binding the contract method 0x19ab453c.
//
// Solidity: function init(address _rootchain) returns(bool)
func (_RequestableERC20Wrapper *RequestableERC20WrapperTransactor) Init(opts *bind.TransactOpts, _rootchain common.Address) (*types.Transaction, error) {
	return _RequestableERC20Wrapper.contract.Transact(opts, "init", _rootchain)
}

// Init is a paid mutator transaction binding the contract method 0x19ab453c.
//
// Solidity: function init(address _rootchain) returns(bool)
func (_RequestableERC20Wrapper *RequestableERC20WrapperSession) Init(_rootchain common.Address) (*types.Transaction, error) {
	return _RequestableERC20Wrapper.Contract.Init(&_RequestableERC20Wrapper.TransactOpts, _rootchain)
}

// Init is a paid mutator transaction binding the contract method 0x19ab453c.
//
// Solidity: function init(address _rootchain) returns(bool)
func (_RequestableERC20Wrapper *RequestableERC20WrapperTransactorSession) Init(_rootchain common.Address) (*types.Transaction, error) {
	return _RequestableERC20Wrapper.Contract.Init(&_RequestableERC20Wrapper.TransactOpts, _rootchain)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_RequestableERC20Wrapper *RequestableERC20WrapperTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _RequestableERC20Wrapper.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_RequestableERC20Wrapper *RequestableERC20WrapperSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _RequestableERC20Wrapper.Contract.Transfer(&_RequestableERC20Wrapper.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_RequestableERC20Wrapper *RequestableERC20WrapperTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _RequestableERC20Wrapper.Contract.Transfer(&_RequestableERC20Wrapper.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_RequestableERC20Wrapper *RequestableERC20WrapperTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _RequestableERC20Wrapper.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_RequestableERC20Wrapper *RequestableERC20WrapperSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _RequestableERC20Wrapper.Contract.TransferFrom(&_RequestableERC20Wrapper.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_RequestableERC20Wrapper *RequestableERC20WrapperTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _RequestableERC20Wrapper.Contract.TransferFrom(&_RequestableERC20Wrapper.TransactOpts, _from, _to, _value)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amount) returns(bool)
func (_RequestableERC20Wrapper *RequestableERC20WrapperTransactor) Withdraw(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _RequestableERC20Wrapper.contract.Transact(opts, "withdraw", _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amount) returns(bool)
func (_RequestableERC20Wrapper *RequestableERC20WrapperSession) Withdraw(_amount *big.Int) (*types.Transaction, error) {
	return _RequestableERC20Wrapper.Contract.Withdraw(&_RequestableERC20Wrapper.TransactOpts, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amount) returns(bool)
func (_RequestableERC20Wrapper *RequestableERC20WrapperTransactorSession) Withdraw(_amount *big.Int) (*types.Transaction, error) {
	return _RequestableERC20Wrapper.Contract.Withdraw(&_RequestableERC20Wrapper.TransactOpts, _amount)
}

// RequestableERC20WrapperApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the RequestableERC20Wrapper contract.
type RequestableERC20WrapperApprovalIterator struct {
	Event *RequestableERC20WrapperApproval // Event containing the contract specifics and raw log

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
func (it *RequestableERC20WrapperApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RequestableERC20WrapperApproval)
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
		it.Event = new(RequestableERC20WrapperApproval)
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
func (it *RequestableERC20WrapperApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RequestableERC20WrapperApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RequestableERC20WrapperApproval represents a Approval event raised by the RequestableERC20Wrapper contract.
type RequestableERC20WrapperApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_RequestableERC20Wrapper *RequestableERC20WrapperFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*RequestableERC20WrapperApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _RequestableERC20Wrapper.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &RequestableERC20WrapperApprovalIterator{contract: _RequestableERC20Wrapper.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_RequestableERC20Wrapper *RequestableERC20WrapperFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *RequestableERC20WrapperApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _RequestableERC20Wrapper.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RequestableERC20WrapperApproval)
				if err := _RequestableERC20Wrapper.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_RequestableERC20Wrapper *RequestableERC20WrapperFilterer) ParseApproval(log types.Log) (*RequestableERC20WrapperApproval, error) {
	event := new(RequestableERC20WrapperApproval)
	if err := _RequestableERC20Wrapper.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RequestableERC20WrapperDeposittedIterator is returned from FilterDepositted and is used to iterate over the raw logs and unpacked data for Depositted events raised by the RequestableERC20Wrapper contract.
type RequestableERC20WrapperDeposittedIterator struct {
	Event *RequestableERC20WrapperDepositted // Event containing the contract specifics and raw log

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
func (it *RequestableERC20WrapperDeposittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RequestableERC20WrapperDepositted)
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
		it.Event = new(RequestableERC20WrapperDepositted)
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
func (it *RequestableERC20WrapperDeposittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RequestableERC20WrapperDeposittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RequestableERC20WrapperDepositted represents a Depositted event raised by the RequestableERC20Wrapper contract.
type RequestableERC20WrapperDepositted struct {
	From  common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterDepositted is a free log retrieval operation binding the contract event 0x46fc0825df492a7a7f1fb1b690f6912ecb69575484fbb24be0f398ae70d97db4.
//
// Solidity: event Depositted(address _from, uint256 _value)
func (_RequestableERC20Wrapper *RequestableERC20WrapperFilterer) FilterDepositted(opts *bind.FilterOpts) (*RequestableERC20WrapperDeposittedIterator, error) {

	logs, sub, err := _RequestableERC20Wrapper.contract.FilterLogs(opts, "Depositted")
	if err != nil {
		return nil, err
	}
	return &RequestableERC20WrapperDeposittedIterator{contract: _RequestableERC20Wrapper.contract, event: "Depositted", logs: logs, sub: sub}, nil
}

// WatchDepositted is a free log subscription operation binding the contract event 0x46fc0825df492a7a7f1fb1b690f6912ecb69575484fbb24be0f398ae70d97db4.
//
// Solidity: event Depositted(address _from, uint256 _value)
func (_RequestableERC20Wrapper *RequestableERC20WrapperFilterer) WatchDepositted(opts *bind.WatchOpts, sink chan<- *RequestableERC20WrapperDepositted) (event.Subscription, error) {

	logs, sub, err := _RequestableERC20Wrapper.contract.WatchLogs(opts, "Depositted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RequestableERC20WrapperDepositted)
				if err := _RequestableERC20Wrapper.contract.UnpackLog(event, "Depositted", log); err != nil {
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

// ParseDepositted is a log parse operation binding the contract event 0x46fc0825df492a7a7f1fb1b690f6912ecb69575484fbb24be0f398ae70d97db4.
//
// Solidity: event Depositted(address _from, uint256 _value)
func (_RequestableERC20Wrapper *RequestableERC20WrapperFilterer) ParseDepositted(log types.Log) (*RequestableERC20WrapperDepositted, error) {
	event := new(RequestableERC20WrapperDepositted)
	if err := _RequestableERC20Wrapper.contract.UnpackLog(event, "Depositted", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RequestableERC20WrapperRequestCreatedIterator is returned from FilterRequestCreated and is used to iterate over the raw logs and unpacked data for RequestCreated events raised by the RequestableERC20Wrapper contract.
type RequestableERC20WrapperRequestCreatedIterator struct {
	Event *RequestableERC20WrapperRequestCreated // Event containing the contract specifics and raw log

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
func (it *RequestableERC20WrapperRequestCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RequestableERC20WrapperRequestCreated)
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
		it.Event = new(RequestableERC20WrapperRequestCreated)
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
func (it *RequestableERC20WrapperRequestCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RequestableERC20WrapperRequestCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RequestableERC20WrapperRequestCreated represents a RequestCreated event raised by the RequestableERC20Wrapper contract.
type RequestableERC20WrapperRequestCreated struct {
	IsExit    bool
	Requestor common.Address
	TrieKey   [32]byte
	Value     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRequestCreated is a free log retrieval operation binding the contract event 0x4b0c23cd90bac23cdaa0958fbd0306f53e20cd5e65676d21b9ceb109784e15b8.
//
// Solidity: event RequestCreated(bool _isExit, address _requestor, bytes32 _trieKey, uint256 _value)
func (_RequestableERC20Wrapper *RequestableERC20WrapperFilterer) FilterRequestCreated(opts *bind.FilterOpts) (*RequestableERC20WrapperRequestCreatedIterator, error) {

	logs, sub, err := _RequestableERC20Wrapper.contract.FilterLogs(opts, "RequestCreated")
	if err != nil {
		return nil, err
	}
	return &RequestableERC20WrapperRequestCreatedIterator{contract: _RequestableERC20Wrapper.contract, event: "RequestCreated", logs: logs, sub: sub}, nil
}

// WatchRequestCreated is a free log subscription operation binding the contract event 0x4b0c23cd90bac23cdaa0958fbd0306f53e20cd5e65676d21b9ceb109784e15b8.
//
// Solidity: event RequestCreated(bool _isExit, address _requestor, bytes32 _trieKey, uint256 _value)
func (_RequestableERC20Wrapper *RequestableERC20WrapperFilterer) WatchRequestCreated(opts *bind.WatchOpts, sink chan<- *RequestableERC20WrapperRequestCreated) (event.Subscription, error) {

	logs, sub, err := _RequestableERC20Wrapper.contract.WatchLogs(opts, "RequestCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RequestableERC20WrapperRequestCreated)
				if err := _RequestableERC20Wrapper.contract.UnpackLog(event, "RequestCreated", log); err != nil {
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

// ParseRequestCreated is a log parse operation binding the contract event 0x4b0c23cd90bac23cdaa0958fbd0306f53e20cd5e65676d21b9ceb109784e15b8.
//
// Solidity: event RequestCreated(bool _isExit, address _requestor, bytes32 _trieKey, uint256 _value)
func (_RequestableERC20Wrapper *RequestableERC20WrapperFilterer) ParseRequestCreated(log types.Log) (*RequestableERC20WrapperRequestCreated, error) {
	event := new(RequestableERC20WrapperRequestCreated)
	if err := _RequestableERC20Wrapper.contract.UnpackLog(event, "RequestCreated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RequestableERC20WrapperTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the RequestableERC20Wrapper contract.
type RequestableERC20WrapperTransferIterator struct {
	Event *RequestableERC20WrapperTransfer // Event containing the contract specifics and raw log

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
func (it *RequestableERC20WrapperTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RequestableERC20WrapperTransfer)
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
		it.Event = new(RequestableERC20WrapperTransfer)
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
func (it *RequestableERC20WrapperTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RequestableERC20WrapperTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RequestableERC20WrapperTransfer represents a Transfer event raised by the RequestableERC20Wrapper contract.
type RequestableERC20WrapperTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_RequestableERC20Wrapper *RequestableERC20WrapperFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*RequestableERC20WrapperTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _RequestableERC20Wrapper.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &RequestableERC20WrapperTransferIterator{contract: _RequestableERC20Wrapper.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_RequestableERC20Wrapper *RequestableERC20WrapperFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *RequestableERC20WrapperTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _RequestableERC20Wrapper.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RequestableERC20WrapperTransfer)
				if err := _RequestableERC20Wrapper.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_RequestableERC20Wrapper *RequestableERC20WrapperFilterer) ParseTransfer(log types.Log) (*RequestableERC20WrapperTransfer, error) {
	event := new(RequestableERC20WrapperTransfer)
	if err := _RequestableERC20Wrapper.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RequestableERC20WrapperWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the RequestableERC20Wrapper contract.
type RequestableERC20WrapperWithdrawnIterator struct {
	Event *RequestableERC20WrapperWithdrawn // Event containing the contract specifics and raw log

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
func (it *RequestableERC20WrapperWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RequestableERC20WrapperWithdrawn)
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
		it.Event = new(RequestableERC20WrapperWithdrawn)
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
func (it *RequestableERC20WrapperWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RequestableERC20WrapperWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RequestableERC20WrapperWithdrawn represents a Withdrawn event raised by the RequestableERC20Wrapper contract.
type RequestableERC20WrapperWithdrawn struct {
	From  common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address _from, uint256 _value)
func (_RequestableERC20Wrapper *RequestableERC20WrapperFilterer) FilterWithdrawn(opts *bind.FilterOpts) (*RequestableERC20WrapperWithdrawnIterator, error) {

	logs, sub, err := _RequestableERC20Wrapper.contract.FilterLogs(opts, "Withdrawn")
	if err != nil {
		return nil, err
	}
	return &RequestableERC20WrapperWithdrawnIterator{contract: _RequestableERC20Wrapper.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address _from, uint256 _value)
func (_RequestableERC20Wrapper *RequestableERC20WrapperFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *RequestableERC20WrapperWithdrawn) (event.Subscription, error) {

	logs, sub, err := _RequestableERC20Wrapper.contract.WatchLogs(opts, "Withdrawn")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RequestableERC20WrapperWithdrawn)
				if err := _RequestableERC20Wrapper.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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

// ParseWithdrawn is a log parse operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address _from, uint256 _value)
func (_RequestableERC20Wrapper *RequestableERC20WrapperFilterer) ParseWithdrawn(log types.Log) (*RequestableERC20WrapperWithdrawn, error) {
	event := new(RequestableERC20WrapperWithdrawn)
	if err := _RequestableERC20Wrapper.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RequestableIABI is the input ABI used to generate the binding from.
const RequestableIABI = "[{\"constant\":false,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"isExit\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"requestor\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"trieKey\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"trieValue\",\"type\":\"bytes\"}],\"name\":\"applyRequestInChildChain\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"isExit\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"requestor\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"trieKey\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"trieValue\",\"type\":\"bytes\"}],\"name\":\"applyRequestInRootChain\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

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

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
var SafeMathBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a723158201b7c830ded9c443d2743d32c4392152cca155e66431e8b7b6f29c49182b1f85464736f6c634300050c0032"

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

// StandardTokenABI is the input ABI used to generate the binding from.
const StandardTokenABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseApproval\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseApproval\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// StandardTokenFuncSigs maps the 4-byte function signature to its string representation.
var StandardTokenFuncSigs = map[string]string{
	"dd62ed3e": "allowance(address,address)",
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"66188463": "decreaseApproval(address,uint256)",
	"d73dd623": "increaseApproval(address,uint256)",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
}

// StandardTokenBin is the compiled bytecode used for deploying new contracts.
var StandardTokenBin = "0x608060405234801561001057600080fd5b506106d9806100206000396000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c806370a082311161005b57806370a0823114610149578063a9059cbb1461016f578063d73dd6231461019b578063dd62ed3e146101c757610088565b8063095ea7b31461008d57806318160ddd146100cd57806323b872dd146100e7578063661884631461011d575b600080fd5b6100b9600480360360408110156100a357600080fd5b506001600160a01b0381351690602001356101f5565b604080519115158252519081900360200190f35b6100d561025b565b60408051918252519081900360200190f35b6100b9600480360360608110156100fd57600080fd5b506001600160a01b03813581169160208101359091169060400135610261565b6100b96004803603604081101561013357600080fd5b506001600160a01b0381351690602001356103d4565b6100d56004803603602081101561015f57600080fd5b50356001600160a01b03166104c3565b6100b96004803603604081101561018557600080fd5b506001600160a01b0381351690602001356104de565b6100b9600480360360408110156101b157600080fd5b506001600160a01b0381351690602001356105bb565b6100d5600480360360408110156101dd57600080fd5b506001600160a01b0381358116916020013516610654565b3360008181526002602090815260408083206001600160a01b038716808552908352818420869055815186815291519394909390927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925928290030190a350600192915050565b60015490565b6001600160a01b03831660009081526020819052604081205482111561028657600080fd5b6001600160a01b03841660009081526002602090815260408083203384529091529020548211156102b657600080fd5b6001600160a01b0383166102c957600080fd5b6001600160a01b0384166000908152602081905260409020546102f2908363ffffffff61067f16565b6001600160a01b038086166000908152602081905260408082209390935590851681522054610327908363ffffffff61069116565b6001600160a01b03808516600090815260208181526040808320949094559187168152600282528281203382529091522054610369908363ffffffff61067f16565b6001600160a01b03808616600081815260026020908152604080832033845282529182902094909455805186815290519287169391927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef929181900390910190a35060019392505050565b3360009081526002602090815260408083206001600160a01b0386168452909152812054808310610428573360009081526002602090815260408083206001600160a01b038816845290915281205561045d565b610438818463ffffffff61067f16565b3360009081526002602090815260408083206001600160a01b03891684529091529020555b3360008181526002602090815260408083206001600160a01b0389168085529083529281902054815190815290519293927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929181900390910190a35060019392505050565b6001600160a01b031660009081526020819052604090205490565b336000908152602081905260408120548211156104fa57600080fd5b6001600160a01b03831661050d57600080fd5b3360009081526020819052604090205461052d908363ffffffff61067f16565b33600090815260208190526040808220929092556001600160a01b0385168152205461055f908363ffffffff61069116565b6001600160a01b038416600081815260208181526040918290209390935580518581529051919233927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9281900390910190a350600192915050565b3360009081526002602090815260408083206001600160a01b03861684529091528120546105ef908363ffffffff61069116565b3360008181526002602090815260408083206001600160a01b0389168085529083529281902085905580519485525191937f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929081900390910190a350600192915050565b6001600160a01b03918216600090815260026020908152604080832093909416825291909152205490565b60008282111561068b57fe5b50900390565b8181018281101561069e57fe5b9291505056fea265627a7a72315820676e61f9766dfb7ce861217ef75a6f0704bff0e9f8ce0c7ce7521b999097bc6f64736f6c634300050c0032"

// DeployStandardToken deploys a new Ethereum contract, binding an instance of StandardToken to it.
func DeployStandardToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *StandardToken, error) {
	parsed, err := abi.JSON(strings.NewReader(StandardTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(StandardTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &StandardToken{StandardTokenCaller: StandardTokenCaller{contract: contract}, StandardTokenTransactor: StandardTokenTransactor{contract: contract}, StandardTokenFilterer: StandardTokenFilterer{contract: contract}}, nil
}

// StandardToken is an auto generated Go binding around an Ethereum contract.
type StandardToken struct {
	StandardTokenCaller     // Read-only binding to the contract
	StandardTokenTransactor // Write-only binding to the contract
	StandardTokenFilterer   // Log filterer for contract events
}

// StandardTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type StandardTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StandardTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StandardTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StandardTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StandardTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StandardTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StandardTokenSession struct {
	Contract     *StandardToken    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StandardTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StandardTokenCallerSession struct {
	Contract *StandardTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// StandardTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StandardTokenTransactorSession struct {
	Contract     *StandardTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// StandardTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type StandardTokenRaw struct {
	Contract *StandardToken // Generic contract binding to access the raw methods on
}

// StandardTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StandardTokenCallerRaw struct {
	Contract *StandardTokenCaller // Generic read-only contract binding to access the raw methods on
}

// StandardTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StandardTokenTransactorRaw struct {
	Contract *StandardTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStandardToken creates a new instance of StandardToken, bound to a specific deployed contract.
func NewStandardToken(address common.Address, backend bind.ContractBackend) (*StandardToken, error) {
	contract, err := bindStandardToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StandardToken{StandardTokenCaller: StandardTokenCaller{contract: contract}, StandardTokenTransactor: StandardTokenTransactor{contract: contract}, StandardTokenFilterer: StandardTokenFilterer{contract: contract}}, nil
}

// NewStandardTokenCaller creates a new read-only instance of StandardToken, bound to a specific deployed contract.
func NewStandardTokenCaller(address common.Address, caller bind.ContractCaller) (*StandardTokenCaller, error) {
	contract, err := bindStandardToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StandardTokenCaller{contract: contract}, nil
}

// NewStandardTokenTransactor creates a new write-only instance of StandardToken, bound to a specific deployed contract.
func NewStandardTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*StandardTokenTransactor, error) {
	contract, err := bindStandardToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StandardTokenTransactor{contract: contract}, nil
}

// NewStandardTokenFilterer creates a new log filterer instance of StandardToken, bound to a specific deployed contract.
func NewStandardTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*StandardTokenFilterer, error) {
	contract, err := bindStandardToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StandardTokenFilterer{contract: contract}, nil
}

// bindStandardToken binds a generic wrapper to an already deployed contract.
func bindStandardToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StandardTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StandardToken *StandardTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _StandardToken.Contract.StandardTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StandardToken *StandardTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StandardToken.Contract.StandardTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StandardToken *StandardTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StandardToken.Contract.StandardTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StandardToken *StandardTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _StandardToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StandardToken *StandardTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StandardToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StandardToken *StandardTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StandardToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) constant returns(uint256)
func (_StandardToken *StandardTokenCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StandardToken.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) constant returns(uint256)
func (_StandardToken *StandardTokenSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _StandardToken.Contract.Allowance(&_StandardToken.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) constant returns(uint256)
func (_StandardToken *StandardTokenCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _StandardToken.Contract.Allowance(&_StandardToken.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) constant returns(uint256)
func (_StandardToken *StandardTokenCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StandardToken.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) constant returns(uint256)
func (_StandardToken *StandardTokenSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _StandardToken.Contract.BalanceOf(&_StandardToken.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) constant returns(uint256)
func (_StandardToken *StandardTokenCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _StandardToken.Contract.BalanceOf(&_StandardToken.CallOpts, _owner)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_StandardToken *StandardTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StandardToken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_StandardToken *StandardTokenSession) TotalSupply() (*big.Int, error) {
	return _StandardToken.Contract.TotalSupply(&_StandardToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_StandardToken *StandardTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _StandardToken.Contract.TotalSupply(&_StandardToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_StandardToken *StandardTokenTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_StandardToken *StandardTokenSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.Approve(&_StandardToken.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_StandardToken *StandardTokenTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.Approve(&_StandardToken.TransactOpts, _spender, _value)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(address _spender, uint256 _subtractedValue) returns(bool)
func (_StandardToken *StandardTokenTransactor) DecreaseApproval(opts *bind.TransactOpts, _spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _StandardToken.contract.Transact(opts, "decreaseApproval", _spender, _subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(address _spender, uint256 _subtractedValue) returns(bool)
func (_StandardToken *StandardTokenSession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.DecreaseApproval(&_StandardToken.TransactOpts, _spender, _subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(address _spender, uint256 _subtractedValue) returns(bool)
func (_StandardToken *StandardTokenTransactorSession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.DecreaseApproval(&_StandardToken.TransactOpts, _spender, _subtractedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(address _spender, uint256 _addedValue) returns(bool)
func (_StandardToken *StandardTokenTransactor) IncreaseApproval(opts *bind.TransactOpts, _spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _StandardToken.contract.Transact(opts, "increaseApproval", _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(address _spender, uint256 _addedValue) returns(bool)
func (_StandardToken *StandardTokenSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.IncreaseApproval(&_StandardToken.TransactOpts, _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(address _spender, uint256 _addedValue) returns(bool)
func (_StandardToken *StandardTokenTransactorSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.IncreaseApproval(&_StandardToken.TransactOpts, _spender, _addedValue)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_StandardToken *StandardTokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_StandardToken *StandardTokenSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.Transfer(&_StandardToken.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_StandardToken *StandardTokenTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.Transfer(&_StandardToken.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_StandardToken *StandardTokenTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_StandardToken *StandardTokenSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.TransferFrom(&_StandardToken.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_StandardToken *StandardTokenTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.TransferFrom(&_StandardToken.TransactOpts, _from, _to, _value)
}

// StandardTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the StandardToken contract.
type StandardTokenApprovalIterator struct {
	Event *StandardTokenApproval // Event containing the contract specifics and raw log

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
func (it *StandardTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StandardTokenApproval)
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
		it.Event = new(StandardTokenApproval)
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
func (it *StandardTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StandardTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StandardTokenApproval represents a Approval event raised by the StandardToken contract.
type StandardTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_StandardToken *StandardTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*StandardTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _StandardToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &StandardTokenApprovalIterator{contract: _StandardToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_StandardToken *StandardTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *StandardTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _StandardToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StandardTokenApproval)
				if err := _StandardToken.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_StandardToken *StandardTokenFilterer) ParseApproval(log types.Log) (*StandardTokenApproval, error) {
	event := new(StandardTokenApproval)
	if err := _StandardToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// StandardTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the StandardToken contract.
type StandardTokenTransferIterator struct {
	Event *StandardTokenTransfer // Event containing the contract specifics and raw log

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
func (it *StandardTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StandardTokenTransfer)
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
		it.Event = new(StandardTokenTransfer)
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
func (it *StandardTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StandardTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StandardTokenTransfer represents a Transfer event raised by the StandardToken contract.
type StandardTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_StandardToken *StandardTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*StandardTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _StandardToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &StandardTokenTransferIterator{contract: _StandardToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_StandardToken *StandardTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *StandardTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _StandardToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StandardTokenTransfer)
				if err := _StandardToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_StandardToken *StandardTokenFilterer) ParseTransfer(log types.Log) (*StandardTokenTransfer, error) {
	event := new(StandardTokenTransfer)
	if err := _StandardToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}
