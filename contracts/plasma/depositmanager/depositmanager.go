// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package depositmanager

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

// AddressABI is the input ABI used to generate the binding from.
const AddressABI = "[]"

// AddressBin is the compiled bytecode used for deploying new contracts.
var AddressBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a72315820caa83848f6f3a88a061ce398577e989d5da9a19232f9e00175de410a704b86e164736f6c634300050c0032"

// DeployAddress deploys a new Ethereum contract, binding an instance of Address to it.
func DeployAddress(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Address, error) {
	parsed, err := abi.JSON(strings.NewReader(AddressABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AddressBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Address{AddressCaller: AddressCaller{contract: contract}, AddressTransactor: AddressTransactor{contract: contract}, AddressFilterer: AddressFilterer{contract: contract}}, nil
}

// Address is an auto generated Go binding around an Ethereum contract.
type Address struct {
	AddressCaller     // Read-only binding to the contract
	AddressTransactor // Write-only binding to the contract
	AddressFilterer   // Log filterer for contract events
}

// AddressCaller is an auto generated read-only Go binding around an Ethereum contract.
type AddressCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AddressTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AddressFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AddressSession struct {
	Contract     *Address          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AddressCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AddressCallerSession struct {
	Contract *AddressCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// AddressTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AddressTransactorSession struct {
	Contract     *AddressTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// AddressRaw is an auto generated low-level Go binding around an Ethereum contract.
type AddressRaw struct {
	Contract *Address // Generic contract binding to access the raw methods on
}

// AddressCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AddressCallerRaw struct {
	Contract *AddressCaller // Generic read-only contract binding to access the raw methods on
}

// AddressTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AddressTransactorRaw struct {
	Contract *AddressTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAddress creates a new instance of Address, bound to a specific deployed contract.
func NewAddress(address common.Address, backend bind.ContractBackend) (*Address, error) {
	contract, err := bindAddress(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Address{AddressCaller: AddressCaller{contract: contract}, AddressTransactor: AddressTransactor{contract: contract}, AddressFilterer: AddressFilterer{contract: contract}}, nil
}

// NewAddressCaller creates a new read-only instance of Address, bound to a specific deployed contract.
func NewAddressCaller(address common.Address, caller bind.ContractCaller) (*AddressCaller, error) {
	contract, err := bindAddress(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AddressCaller{contract: contract}, nil
}

// NewAddressTransactor creates a new write-only instance of Address, bound to a specific deployed contract.
func NewAddressTransactor(address common.Address, transactor bind.ContractTransactor) (*AddressTransactor, error) {
	contract, err := bindAddress(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AddressTransactor{contract: contract}, nil
}

// NewAddressFilterer creates a new log filterer instance of Address, bound to a specific deployed contract.
func NewAddressFilterer(address common.Address, filterer bind.ContractFilterer) (*AddressFilterer, error) {
	contract, err := bindAddress(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AddressFilterer{contract: contract}, nil
}

// bindAddress binds a generic wrapper to an already deployed contract.
func bindAddress(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AddressABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Address *AddressRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Address.Contract.AddressCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Address *AddressRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Address.Contract.AddressTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Address *AddressRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Address.Contract.AddressTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Address *AddressCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Address.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Address *AddressTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Address.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Address *AddressTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Address.Contract.contract.Transact(opts, method, params...)
}

// AuthControllerABI is the input ABI used to generate the binding from.
const AuthControllerABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"renounceMinter\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"renouncePauser\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// AuthControllerFuncSigs maps the 4-byte function signature to its string representation.
var AuthControllerFuncSigs = map[string]string{
	"8f32d59b": "isOwner()",
	"8da5cb5b": "owner()",
	"5f112c68": "renounceMinter(address)",
	"715018a6": "renounceOwnership()",
	"38bf3cfa": "renounceOwnership(address)",
	"41eb24bb": "renouncePauser(address)",
	"f2fde38b": "transferOwnership(address)",
	"6d435421": "transferOwnership(address,address)",
}

// AuthControllerBin is the compiled bytecode used for deploying new contracts.
var AuthControllerBin = "0x608060405260006100176001600160e01b0361006616565b600080546001600160a01b0319166001600160a01b0383169081178255604051929350917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a35061006a565b3390565b61062f806100796000396000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c8063715018a61161005b578063715018a61461012f5780638da5cb5b146101375780638f32d59b1461015b578063f2fde38b1461017757610088565b806338bf3cfa1461008d57806341eb24bb146100b55780635f112c68146100db5780636d43542114610101575b600080fd5b6100b3600480360360208110156100a357600080fd5b50356001600160a01b031661019d565b005b6100b3600480360360208110156100cb57600080fd5b50356001600160a01b031661023a565b6100b3600480360360208110156100f157600080fd5b50356001600160a01b03166102bc565b6100b36004803603604081101561011757600080fd5b506001600160a01b038135811691602001351661033e565b6100b36103f9565b61013f61048a565b604080516001600160a01b039092168252519081900360200190f35b610163610499565b604080519115158252519081900360200190f35b6100b36004803603602081101561018d57600080fd5b50356001600160a01b03166104bd565b6101a5610499565b6101e4576040805162461bcd60e51b815260206004820181905260248201526000805160206105db833981519152604482015290519081900360640190fd5b806001600160a01b031663715018a66040518163ffffffff1660e01b8152600401600060405180830381600087803b15801561021f57600080fd5b505af1158015610233573d6000803e3d6000fd5b5050505050565b610242610499565b610281576040805162461bcd60e51b815260206004820181905260248201526000805160206105db833981519152604482015290519081900360640190fd5b806001600160a01b0316636ef8d66d6040518163ffffffff1660e01b8152600401600060405180830381600087803b15801561021f57600080fd5b6102c4610499565b610303576040805162461bcd60e51b815260206004820181905260248201526000805160206105db833981519152604482015290519081900360640190fd5b806001600160a01b031663986502756040518163ffffffff1660e01b8152600401600060405180830381600087803b15801561021f57600080fd5b610346610499565b610385576040805162461bcd60e51b815260206004820181905260248201526000805160206105db833981519152604482015290519081900360640190fd5b816001600160a01b031663f2fde38b826040518263ffffffff1660e01b815260040180826001600160a01b03166001600160a01b03168152602001915050600060405180830381600087803b1580156103dd57600080fd5b505af11580156103f1573d6000803e3d6000fd5b505050505050565b610401610499565b610440576040805162461bcd60e51b815260206004820181905260248201526000805160206105db833981519152604482015290519081900360640190fd5b600080546040516001600160a01b03909116907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3600080546001600160a01b0319169055565b6000546001600160a01b031690565b600080546001600160a01b03166104ae610510565b6001600160a01b031614905090565b6104c5610499565b610504576040805162461bcd60e51b815260206004820181905260248201526000805160206105db833981519152604482015290519081900360640190fd5b61050d81610514565b50565b3390565b6001600160a01b0381166105595760405162461bcd60e51b81526004018080602001828103825260268152602001806105b56026913960400191505060405180910390fd5b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b039290921691909117905556fe4f776e61626c653a206e6577206f776e657220697320746865207a65726f20616464726573734f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572a265627a7a72315820f5dbf027aa9fb336aafe169eae7ece5653dd467fdb8b8d677f8775fa7b10105264736f6c634300050c0032"

// DeployAuthController deploys a new Ethereum contract, binding an instance of AuthController to it.
func DeployAuthController(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *AuthController, error) {
	parsed, err := abi.JSON(strings.NewReader(AuthControllerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AuthControllerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AuthController{AuthControllerCaller: AuthControllerCaller{contract: contract}, AuthControllerTransactor: AuthControllerTransactor{contract: contract}, AuthControllerFilterer: AuthControllerFilterer{contract: contract}}, nil
}

// AuthController is an auto generated Go binding around an Ethereum contract.
type AuthController struct {
	AuthControllerCaller     // Read-only binding to the contract
	AuthControllerTransactor // Write-only binding to the contract
	AuthControllerFilterer   // Log filterer for contract events
}

// AuthControllerCaller is an auto generated read-only Go binding around an Ethereum contract.
type AuthControllerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuthControllerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AuthControllerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuthControllerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AuthControllerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuthControllerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AuthControllerSession struct {
	Contract     *AuthController   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AuthControllerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AuthControllerCallerSession struct {
	Contract *AuthControllerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// AuthControllerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AuthControllerTransactorSession struct {
	Contract     *AuthControllerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// AuthControllerRaw is an auto generated low-level Go binding around an Ethereum contract.
type AuthControllerRaw struct {
	Contract *AuthController // Generic contract binding to access the raw methods on
}

// AuthControllerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AuthControllerCallerRaw struct {
	Contract *AuthControllerCaller // Generic read-only contract binding to access the raw methods on
}

// AuthControllerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AuthControllerTransactorRaw struct {
	Contract *AuthControllerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAuthController creates a new instance of AuthController, bound to a specific deployed contract.
func NewAuthController(address common.Address, backend bind.ContractBackend) (*AuthController, error) {
	contract, err := bindAuthController(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AuthController{AuthControllerCaller: AuthControllerCaller{contract: contract}, AuthControllerTransactor: AuthControllerTransactor{contract: contract}, AuthControllerFilterer: AuthControllerFilterer{contract: contract}}, nil
}

// NewAuthControllerCaller creates a new read-only instance of AuthController, bound to a specific deployed contract.
func NewAuthControllerCaller(address common.Address, caller bind.ContractCaller) (*AuthControllerCaller, error) {
	contract, err := bindAuthController(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AuthControllerCaller{contract: contract}, nil
}

// NewAuthControllerTransactor creates a new write-only instance of AuthController, bound to a specific deployed contract.
func NewAuthControllerTransactor(address common.Address, transactor bind.ContractTransactor) (*AuthControllerTransactor, error) {
	contract, err := bindAuthController(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AuthControllerTransactor{contract: contract}, nil
}

// NewAuthControllerFilterer creates a new log filterer instance of AuthController, bound to a specific deployed contract.
func NewAuthControllerFilterer(address common.Address, filterer bind.ContractFilterer) (*AuthControllerFilterer, error) {
	contract, err := bindAuthController(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AuthControllerFilterer{contract: contract}, nil
}

// bindAuthController binds a generic wrapper to an already deployed contract.
func bindAuthController(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AuthControllerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AuthController *AuthControllerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AuthController.Contract.AuthControllerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AuthController *AuthControllerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AuthController.Contract.AuthControllerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AuthController *AuthControllerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AuthController.Contract.AuthControllerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AuthController *AuthControllerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AuthController.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AuthController *AuthControllerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AuthController.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AuthController *AuthControllerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AuthController.Contract.contract.Transact(opts, method, params...)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_AuthController *AuthControllerCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _AuthController.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_AuthController *AuthControllerSession) IsOwner() (bool, error) {
	return _AuthController.Contract.IsOwner(&_AuthController.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_AuthController *AuthControllerCallerSession) IsOwner() (bool, error) {
	return _AuthController.Contract.IsOwner(&_AuthController.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_AuthController *AuthControllerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _AuthController.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_AuthController *AuthControllerSession) Owner() (common.Address, error) {
	return _AuthController.Contract.Owner(&_AuthController.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_AuthController *AuthControllerCallerSession) Owner() (common.Address, error) {
	return _AuthController.Contract.Owner(&_AuthController.CallOpts)
}

// RenounceMinter is a paid mutator transaction binding the contract method 0x5f112c68.
//
// Solidity: function renounceMinter(address target) returns()
func (_AuthController *AuthControllerTransactor) RenounceMinter(opts *bind.TransactOpts, target common.Address) (*types.Transaction, error) {
	return _AuthController.contract.Transact(opts, "renounceMinter", target)
}

// RenounceMinter is a paid mutator transaction binding the contract method 0x5f112c68.
//
// Solidity: function renounceMinter(address target) returns()
func (_AuthController *AuthControllerSession) RenounceMinter(target common.Address) (*types.Transaction, error) {
	return _AuthController.Contract.RenounceMinter(&_AuthController.TransactOpts, target)
}

// RenounceMinter is a paid mutator transaction binding the contract method 0x5f112c68.
//
// Solidity: function renounceMinter(address target) returns()
func (_AuthController *AuthControllerTransactorSession) RenounceMinter(target common.Address) (*types.Transaction, error) {
	return _AuthController.Contract.RenounceMinter(&_AuthController.TransactOpts, target)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x38bf3cfa.
//
// Solidity: function renounceOwnership(address target) returns()
func (_AuthController *AuthControllerTransactor) RenounceOwnership(opts *bind.TransactOpts, target common.Address) (*types.Transaction, error) {
	return _AuthController.contract.Transact(opts, "renounceOwnership", target)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x38bf3cfa.
//
// Solidity: function renounceOwnership(address target) returns()
func (_AuthController *AuthControllerSession) RenounceOwnership(target common.Address) (*types.Transaction, error) {
	return _AuthController.Contract.RenounceOwnership(&_AuthController.TransactOpts, target)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x38bf3cfa.
//
// Solidity: function renounceOwnership(address target) returns()
func (_AuthController *AuthControllerTransactorSession) RenounceOwnership(target common.Address) (*types.Transaction, error) {
	return _AuthController.Contract.RenounceOwnership(&_AuthController.TransactOpts, target)
}

// RenounceOwnership0 is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AuthController *AuthControllerTransactor) RenounceOwnership0(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AuthController.contract.Transact(opts, "renounceOwnership0")
}

// RenounceOwnership0 is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AuthController *AuthControllerSession) RenounceOwnership0() (*types.Transaction, error) {
	return _AuthController.Contract.RenounceOwnership0(&_AuthController.TransactOpts)
}

// RenounceOwnership0 is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AuthController *AuthControllerTransactorSession) RenounceOwnership0() (*types.Transaction, error) {
	return _AuthController.Contract.RenounceOwnership0(&_AuthController.TransactOpts)
}

// RenouncePauser is a paid mutator transaction binding the contract method 0x41eb24bb.
//
// Solidity: function renouncePauser(address target) returns()
func (_AuthController *AuthControllerTransactor) RenouncePauser(opts *bind.TransactOpts, target common.Address) (*types.Transaction, error) {
	return _AuthController.contract.Transact(opts, "renouncePauser", target)
}

// RenouncePauser is a paid mutator transaction binding the contract method 0x41eb24bb.
//
// Solidity: function renouncePauser(address target) returns()
func (_AuthController *AuthControllerSession) RenouncePauser(target common.Address) (*types.Transaction, error) {
	return _AuthController.Contract.RenouncePauser(&_AuthController.TransactOpts, target)
}

// RenouncePauser is a paid mutator transaction binding the contract method 0x41eb24bb.
//
// Solidity: function renouncePauser(address target) returns()
func (_AuthController *AuthControllerTransactorSession) RenouncePauser(target common.Address) (*types.Transaction, error) {
	return _AuthController.Contract.RenouncePauser(&_AuthController.TransactOpts, target)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0x6d435421.
//
// Solidity: function transferOwnership(address target, address newOwner) returns()
func (_AuthController *AuthControllerTransactor) TransferOwnership(opts *bind.TransactOpts, target common.Address, newOwner common.Address) (*types.Transaction, error) {
	return _AuthController.contract.Transact(opts, "transferOwnership", target, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0x6d435421.
//
// Solidity: function transferOwnership(address target, address newOwner) returns()
func (_AuthController *AuthControllerSession) TransferOwnership(target common.Address, newOwner common.Address) (*types.Transaction, error) {
	return _AuthController.Contract.TransferOwnership(&_AuthController.TransactOpts, target, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0x6d435421.
//
// Solidity: function transferOwnership(address target, address newOwner) returns()
func (_AuthController *AuthControllerTransactorSession) TransferOwnership(target common.Address, newOwner common.Address) (*types.Transaction, error) {
	return _AuthController.Contract.TransferOwnership(&_AuthController.TransactOpts, target, newOwner)
}

// TransferOwnership0 is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AuthController *AuthControllerTransactor) TransferOwnership0(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _AuthController.contract.Transact(opts, "transferOwnership0", newOwner)
}

// TransferOwnership0 is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AuthController *AuthControllerSession) TransferOwnership0(newOwner common.Address) (*types.Transaction, error) {
	return _AuthController.Contract.TransferOwnership0(&_AuthController.TransactOpts, newOwner)
}

// TransferOwnership0 is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AuthController *AuthControllerTransactorSession) TransferOwnership0(newOwner common.Address) (*types.Transaction, error) {
	return _AuthController.Contract.TransferOwnership0(&_AuthController.TransactOpts, newOwner)
}

// AuthControllerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the AuthController contract.
type AuthControllerOwnershipTransferredIterator struct {
	Event *AuthControllerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *AuthControllerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuthControllerOwnershipTransferred)
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
		it.Event = new(AuthControllerOwnershipTransferred)
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
func (it *AuthControllerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuthControllerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuthControllerOwnershipTransferred represents a OwnershipTransferred event raised by the AuthController contract.
type AuthControllerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AuthController *AuthControllerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AuthControllerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AuthController.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AuthControllerOwnershipTransferredIterator{contract: _AuthController.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AuthController *AuthControllerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AuthControllerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AuthController.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuthControllerOwnershipTransferred)
				if err := _AuthController.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_AuthController *AuthControllerFilterer) ParseOwnershipTransferred(log types.Log) (*AuthControllerOwnershipTransferred, error) {
	event := new(AuthControllerOwnershipTransferred)
	if err := _AuthController.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

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

// DSMathABI is the input ABI used to generate the binding from.
const DSMathABI = "[]"

// DSMathBin is the compiled bytecode used for deploying new contracts.
var DSMathBin = "0x6080604052348015600f57600080fd5b50603e80601d6000396000f3fe6080604052600080fdfea265627a7a723158201d3383a4efea7d4c911511edcc2240f991199f32cf3b3c855dd9ffb45f3f45a164736f6c634300050c0032"

// DeployDSMath deploys a new Ethereum contract, binding an instance of DSMath to it.
func DeployDSMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DSMath, error) {
	parsed, err := abi.JSON(strings.NewReader(DSMathABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DSMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DSMath{DSMathCaller: DSMathCaller{contract: contract}, DSMathTransactor: DSMathTransactor{contract: contract}, DSMathFilterer: DSMathFilterer{contract: contract}}, nil
}

// DSMath is an auto generated Go binding around an Ethereum contract.
type DSMath struct {
	DSMathCaller     // Read-only binding to the contract
	DSMathTransactor // Write-only binding to the contract
	DSMathFilterer   // Log filterer for contract events
}

// DSMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type DSMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DSMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DSMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DSMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DSMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DSMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DSMathSession struct {
	Contract     *DSMath           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DSMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DSMathCallerSession struct {
	Contract *DSMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// DSMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DSMathTransactorSession struct {
	Contract     *DSMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DSMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type DSMathRaw struct {
	Contract *DSMath // Generic contract binding to access the raw methods on
}

// DSMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DSMathCallerRaw struct {
	Contract *DSMathCaller // Generic read-only contract binding to access the raw methods on
}

// DSMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DSMathTransactorRaw struct {
	Contract *DSMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDSMath creates a new instance of DSMath, bound to a specific deployed contract.
func NewDSMath(address common.Address, backend bind.ContractBackend) (*DSMath, error) {
	contract, err := bindDSMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DSMath{DSMathCaller: DSMathCaller{contract: contract}, DSMathTransactor: DSMathTransactor{contract: contract}, DSMathFilterer: DSMathFilterer{contract: contract}}, nil
}

// NewDSMathCaller creates a new read-only instance of DSMath, bound to a specific deployed contract.
func NewDSMathCaller(address common.Address, caller bind.ContractCaller) (*DSMathCaller, error) {
	contract, err := bindDSMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DSMathCaller{contract: contract}, nil
}

// NewDSMathTransactor creates a new write-only instance of DSMath, bound to a specific deployed contract.
func NewDSMathTransactor(address common.Address, transactor bind.ContractTransactor) (*DSMathTransactor, error) {
	contract, err := bindDSMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DSMathTransactor{contract: contract}, nil
}

// NewDSMathFilterer creates a new log filterer instance of DSMath, bound to a specific deployed contract.
func NewDSMathFilterer(address common.Address, filterer bind.ContractFilterer) (*DSMathFilterer, error) {
	contract, err := bindDSMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DSMathFilterer{contract: contract}, nil
}

// bindDSMath binds a generic wrapper to an already deployed contract.
func bindDSMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DSMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DSMath *DSMathRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DSMath.Contract.DSMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DSMath *DSMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DSMath.Contract.DSMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DSMath *DSMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DSMath.Contract.DSMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DSMath *DSMathCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DSMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DSMath *DSMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DSMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DSMath *DSMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DSMath.Contract.contract.Transact(opts, method, params...)
}

// DepositManagerABI is the input ABI used to generate the binding from.
const DepositManagerABI = "[{\"inputs\":[{\"internalType\":\"contractWTON\",\"name\":\"wton\",\"type\":\"address\"},{\"internalType\":\"contractRootChainRegistryI\",\"name\":\"registry\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"WITHDRAWAL_DELAY\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"rootchain\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"rootchain\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawalProcessed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"rootchain\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawalRequested\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"INTERFACE_ID_ON_APPROVE\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"WITHDRAWAL_DELAY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"rootchain\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"accStaked\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"wtonAmount\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"accStakedAccount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"wtonAmount\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"rootchain\",\"type\":\"address\"}],\"name\":\"accStakedRootChain\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"wtonAmount\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"rootchain\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"accUnstaked\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"wtonAmount\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"accUnstakedAccount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"wtonAmount\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"rootchain\",\"type\":\"address\"}],\"name\":\"accUnstakedRootChain\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"wtonAmount\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"rootchain\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"rootchain\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"numPendingRequests\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"rootchain\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"numRequests\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onApprove\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"rootchain\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"pendingUnstaked\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"wtonAmount\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"pendingUnstakedAccount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"wtonAmount\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"rootchain\",\"type\":\"address\"}],\"name\":\"pendingUnstakedRootChain\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"wtonAmount\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"rootchain\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"receiveTON\",\"type\":\"bool\"}],\"name\":\"processRequest\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"rootchain\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"n\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"receiveTON\",\"type\":\"bool\"}],\"name\":\"processRequests\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"rootchain\",\"type\":\"address\"}],\"name\":\"redeposit\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"rootchain\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"n\",\"type\":\"uint256\"}],\"name\":\"redepositMulti\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"registry\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"rootchain\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"requestWithdrawal\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"rootchain\",\"type\":\"address\"}],\"name\":\"requestWithdrawalAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"seigManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractSeigManagerI\",\"name\":\"seigManager\",\"type\":\"address\"}],\"name\":\"setSeigManager\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"rootchain\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"withdrawalRequest\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"withdrawableBlockNumber\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"},{\"internalType\":\"bool\",\"name\":\"processed\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"rootchain\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"withdrawalRequestIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"wton\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// DepositManagerFuncSigs maps the 4-byte function signature to its string representation.
var DepositManagerFuncSigs = map[string]string{
	"6cd28f9a": "INTERFACE_ID_ON_APPROVE()",
	"0ebb172a": "WITHDRAWAL_DELAY()",
	"2d2fab94": "accStaked(address,address)",
	"0055f5c1": "accStakedAccount(address)",
	"e8035ec9": "accStakedRootChain(address)",
	"9d91b87b": "accUnstaked(address,address)",
	"a3543989": "accUnstakedAccount(address)",
	"03dc6510": "accUnstakedRootChain(address)",
	"47e7ef24": "deposit(address,uint256)",
	"8f32d59b": "isOwner()",
	"5c0df46b": "numPendingRequests(address,address)",
	"f762eb57": "numRequests(address,address)",
	"4273ca16": "onApprove(address,address,uint256,bytes)",
	"8da5cb5b": "owner()",
	"2638fdf5": "pendingUnstaked(address,address)",
	"a0b2a913": "pendingUnstakedAccount(address)",
	"a8f3fb98": "pendingUnstakedRootChain(address)",
	"c20a44c6": "processRequest(address,bool)",
	"fb0713b1": "processRequests(address,uint256,bool)",
	"b8bee628": "redeposit(address)",
	"445e83b9": "redepositMulti(address,uint256)",
	"7b103999": "registry()",
	"715018a6": "renounceOwnership()",
	"da95ebf7": "requestWithdrawal(address,uint256)",
	"6b2160b7": "requestWithdrawalAll(address)",
	"6fb7f558": "seigManager()",
	"7657f20a": "setSeigManager(address)",
	"01ffc9a7": "supportsInterface(bytes4)",
	"f2fde38b": "transferOwnership(address)",
	"8fbef2d0": "withdrawalRequest(address,address,uint256)",
	"c647f26e": "withdrawalRequestIndex(address,address)",
	"8d62d949": "wton()",
}

// DepositManagerBin is the compiled bytecode used for deploying new contracts.
var DepositManagerBin = "0x60806040523480156200001157600080fd5b50604051620025e2380380620025e2833981810160405260608110156200003757600080fd5b508051602082015160409092015190919060006200005d6001600160e01b036200014216565b600080546001600160a01b0319166001600160a01b0383169081178255604051929350917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a350620000db7f01ffc9a7000000000000000000000000000000000000000000000000000000006001600160e01b036200014616565b620001096040518080620025ba6028913960405190819003602801902090506001600160e01b036200014616565b600280546001600160a01b039485166001600160a01b031991821617909155600380549390941692169190911790915560105562000218565b3390565b7fffffffff000000000000000000000000000000000000000000000000000000008082161415620001d857604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601c60248201527f4552433136353a20696e76616c696420696e7465726661636520696400000000604482015290519081900360640190fd5b7fffffffff00000000000000000000000000000000000000000000000000000000166000908152600160208190526040909120805460ff19169091179055565b61239280620002286000396000f3fe608060405234801561001057600080fd5b50600436106101ef5760003560e01c80638d62d9491161010f578063b8bee628116100a2578063e8035ec911610071578063e8035ec914610676578063f2fde38b1461069c578063f762eb57146106c2578063fb0713b1146106f0576101ef565b8063b8bee628146105c8578063c20a44c6146105ee578063c647f26e1461061c578063da95ebf71461064a576101ef565b80639d91b87b116100de5780639d91b87b14610528578063a0b2a91314610556578063a35439891461057c578063a8f3fb98146105a2576101ef565b80638d62d949146104ae5780638da5cb5b146104b65780638f32d59b146104be5780638fbef2d0146104c6576101ef565b806347e7ef24116101875780636fb7f558116101565780636fb7f55814610452578063715018a6146104765780637657f20a146104805780637b103999146104a6576101ef565b806347e7ef24146103ad5780635c0df46b146103d95780636b2160b7146104075780636cd28f9a1461042d576101ef565b80632638fdf5116101c35780632638fdf5146102955780632d2fab94146102c35780634273ca16146102f1578063445e83b914610381576101ef565b806255f5c1146101f457806301ffc9a71461022c57806303dc6510146102675780630ebb172a1461028d575b600080fd5b61021a6004803603602081101561020a57600080fd5b50356001600160a01b0316610724565b60408051918252519081900360200190f35b6102536004803603602081101561024257600080fd5b50356001600160e01b03191661073f565b604080519115158252519081900360200190f35b61021a6004803603602081101561027d57600080fd5b50356001600160a01b031661075e565b61021a610779565b61021a600480360360408110156102ab57600080fd5b506001600160a01b038135811691602001351661077f565b61021a600480360360408110156102d957600080fd5b506001600160a01b03813581169160200135166107ac565b6102536004803603608081101561030757600080fd5b6001600160a01b0382358116926020810135909116916040820135919081019060808101606082013564010000000081111561034257600080fd5b82018360208201111561035457600080fd5b8035906020019184600183028401116401000000008311171561037657600080fd5b5090925090506107d7565b6102536004803603604081101561039757600080fd5b506001600160a01b038135169060200135610884565b610253600480360360408110156103c357600080fd5b506001600160a01b0381351690602001356108c3565b61021a600480360360408110156103ef57600080fd5b506001600160a01b03813581169160200135166108d9565b6102536004803603602081101561041d57600080fd5b50356001600160a01b031661092f565b610435610a4e565b604080516001600160e01b03199092168252519081900360200190f35b61045a610a69565b604080516001600160a01b039092168252519081900360200190f35b61047e610a78565b005b61047e6004803603602081101561049657600080fd5b50356001600160a01b0316610b1b565b61045a610bde565b61045a610bed565b61045a610bfc565b610253610c0b565b6104fc600480360360608110156104dc57600080fd5b506001600160a01b03813581169160208101359091169060400135610c2f565b604080516001600160801b03948516815292909316602083015215158183015290519081900360600190f35b61021a6004803603604081101561053e57600080fd5b506001600160a01b0381358116916020013516610d2a565b61021a6004803603602081101561056c57600080fd5b50356001600160a01b0316610d55565b61021a6004803603602081101561059257600080fd5b50356001600160a01b0316610d70565b61021a600480360360208110156105b857600080fd5b50356001600160a01b0316610d8b565b610253600480360360208110156105de57600080fd5b50356001600160a01b0316610da6565b6102536004803603604081101561060457600080fd5b506001600160a01b0381351690602001351515610ddf565b61021a6004803603604081101561063257600080fd5b506001600160a01b0381358116916020013516610df2565b6102536004803603604081101561066057600080fd5b506001600160a01b038135169060200135610e1d565b61021a6004803603602081101561068c57600080fd5b50356001600160a01b0316610e29565b61047e600480360360208110156106b257600080fd5b50356001600160a01b0316610e44565b61021a600480360360408110156106d857600080fd5b506001600160a01b0381358116916020013516610ea9565b6102536004803603606081101561070657600080fd5b506001600160a01b0381351690602081013590604001351515610ed4565b6001600160a01b031660009081526007602052604090205490565b6001600160e01b03191660009081526001602052604090205460ff1690565b6001600160a01b03166000908152600c602052604090205490565b60105490565b6001600160a01b038083166000908152600860209081526040808320938516835292905220545b92915050565b6001600160a01b03918216600090815260056020908152604080832093909416825291909152205490565b6002546000906001600160a01b031633146108235760405162461bcd60e51b81526004018080602001828103825260318152602001806122db6031913960400191505060405180910390fd5b600061086484848080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250610f0692505050565b9050610871818887610f1e565b61087a57600080fd5b5095945050505050565b6001600160a01b0382166000908152600f602090815260408083203384529091528120546108b384828561116a565b6108bc57600080fd5b5092915050565b60006108d0833384610f1e565b6107a657600080fd5b6001600160a01b038083166000818152600e6020908152604080832094861680845294825280832054938352600f82528083209483529390529182205481610926576000925050506107a6565b90039392505050565b600354604080516302a1529960e01b81526001600160a01b0380851660048301529151600093859316916302a15299916024808301926020929190829003018186803b15801561097e57600080fd5b505afa158015610992573d6000803e3d6000fd5b505050506040513d60208110156109a857600080fd5b50516109b357600080fd5b60048054604080516367265c3b60e11b81526001600160a01b038781169482019490945233602482015290516000939092169163ce4cb87691604480820192602092909190829003018186803b158015610a0c57600080fd5b505afa158015610a20573d6000803e3d6000fd5b505050506040513d6020811015610a3657600080fd5b50519050610a448482611652565b9250505b50919050565b6040518060286121e382396028019050604051809103902081565b6004546001600160a01b031690565b610a80610c0b565b610ad1576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b600080546040516001600160a01b03909116907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3600080546001600160a01b0319169055565b610b23610c0b565b610b74576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b6004546001600160a01b031615610bbc5760405162461bcd60e51b815260040180806020018281038252602a81526020018061220b602a913960400191505060405180910390fd5b600480546001600160a01b0319166001600160a01b0392909216919091179055565b6003546001600160a01b031690565b6002546001600160a01b031690565b6000546001600160a01b031690565b600080546001600160a01b0316610c20611953565b6001600160a01b031614905090565b6001600160a01b038084166000908152600e60209081526040808320938616835292905290812080548291829185908110610c6657fe5b600091825260208083206002909202909101546001600160a01b03808a168452600e83526040808520918a1685529252912080546001600160801b0390921694509085908110610cb257fe5b60009182526020808320600292909202909101546001600160a01b038981168452600e83526040808520918a168552925291208054600160801b9092046001600160801b031693509085908110610d0557fe5b600091825260209091206001600290920201015492969195505060ff90911692509050565b6001600160a01b039182166000908152600b6020908152604080832093909416825291909152205490565b6001600160a01b03166000908152600a602052604090205490565b6001600160a01b03166000908152600d602052604090205490565b6001600160a01b031660009081526009602052604090205490565b6001600160a01b0381166000908152600f60209081526040808320338452909152812054610dd68382600161116a565b610a4857600080fd5b6000610deb8383611957565b9392505050565b6001600160a01b039182166000908152600f6020908152604080832093909416825291909152205490565b6000610deb8383611652565b6001600160a01b031660009081526006602052604090205490565b610e4c610c0b565b610e9d576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b610ea681611d10565b50565b6001600160a01b039182166000908152600e6020908152604080832093909416825291909152205490565b6000805b83811015610efb57610eea8584611957565b610ef357600080fd5b600101610ed8565b506001949350505050565b60008151602014610f1657600080fd5b506020015190565b600354604080516302a1529960e01b81526001600160a01b0380871660048301529151600093879316916302a15299916024808301926020929190829003018186803b158015610f6d57600080fd5b505afa158015610f81573d6000803e3d6000fd5b505050506040513d6020811015610f9757600080fd5b5051610fa257600080fd5b6001600160a01b03808616600090815260056020908152604080832093881683529290522054610fd8908463ffffffff611db016565b6001600160a01b038087166000818152600560209081526040808320948a16835293815283822094909455908152600690925290205461101e908463ffffffff611db016565b6001600160a01b0380871660009081526006602090815260408083209490945591871681526007909152205461105a908463ffffffff611db016565b6001600160a01b0380861660009081526007602052604090209190915560025461108d911685308663ffffffff611e0a16565b604080516001600160a01b038681168252602082018690528251908816927f8752a472e571a816aea92eec8dae9baf628e840f4929fbcc2d155e6233ff68a7928290030190a26004805460408051630412c6d560e41b81526001600160a01b03898116948201949094528784166024820152604481018790529051929091169163412c6d50916064808201926020929091908290030181600087803b15801561113557600080fd5b505af1158015611149573d6000803e3d6000fd5b505050506040513d602081101561115f57600080fd5b5051610efb57600080fd5b600354604080516302a1529960e01b81526001600160a01b0380871660048301529151600093879316916302a15299916024808301926020929190829003018186803b1580156111b957600080fd5b505afa1580156111cd573d6000803e3d6000fd5b505050506040513d60208110156111e357600080fd5b50516111ee57600080fd5b6001600160a01b0385166000908152600e60209081526040808320338452909152812054611263576040805162461bcd60e51b815260206004820152601a60248201527f4465706f7369744d616e616765723a206e6f2072657175657374000000000000604482015290519081900360640190fd5b6001600160a01b0386166000908152600e602090815260408083203384529091529020548590038411156112c85760405162461bcd60e51b81526004018080602001828103825260318152602001806122356031913960400191505060405180910390fd5b8484015b808610156113d9576001600160a01b0387166000908152600e60209081526040808320338452909152812080548890811061130357fe5b6000918252602090912060029091020180546001820154919250600160801b90046001600160801b03169060ff161561136d5760405162461bcd60e51b815260040180806020018281038252603181526020018061218c6031913960400191505060405180910390fd5b600081116113ac5760405162461bcd60e51b81526004018080602001828103825260288152602001806123366028913960400191505060405180910390fd5b6113bc848263ffffffff611db016565b6001928301805460ff191684179055979091019692506112cc9050565b6001600160a01b038716600090815260056020908152604080832033845290915290205461140d908363ffffffff611db016565b6001600160a01b03881660008181526005602090815260408083203384528252808320949094559181526006909152205461144e908363ffffffff611db016565b6001600160a01b038816600090815260066020908152604080832093909355338252600790522054611486908363ffffffff611db016565b336000818152600760209081526040808320949094556001600160a01b038b168252600881528382209282529190915220546114c8908363ffffffff611e6a16565b6001600160a01b038816600081815260086020908152604080832033845282528083209490945591815260099091522054611509908363ffffffff611e6a16565b6001600160a01b038816600090815260096020908152604080832093909355338252600a90522054611541908363ffffffff611e6a16565b336000818152600a60209081526040808320949094556001600160a01b038b16808352600f825284832084845282529184902080548a01905583519283528201859052825190927f8752a472e571a816aea92eec8dae9baf628e840f4929fbcc2d155e6233ff68a7928290030190a26004805460408051630412c6d560e41b81526001600160a01b038b811694820194909452336024820152604481018690529051929091169163412c6d50916064808201926020929091908290030181600087803b15801561161057600080fd5b505af1158015611624573d6000803e3d6000fd5b505050506040513d602081101561163a57600080fd5b505161164557600080fd5b5060019695505050505050565b600354604080516302a1529960e01b81526001600160a01b0380861660048301529151600093869316916302a15299916024808301926020929190829003018186803b1580156116a157600080fd5b505afa1580156116b5573d6000803e3d6000fd5b505050506040513d60208110156116cb57600080fd5b50516116d657600080fd5b600083116117155760405162461bcd60e51b815260040180806020018281038252602781526020018061228b6027913960400191505060405180910390fd5b6001600160a01b0384166000818152600e602090815260408083203380855290835281842082516060810184526010546001600160801b0343909101811682528a8116828701908152828601888152845460018181018755958a52888a2094516002909102909401805492518416600160801b029484166fffffffffffffffffffffffffffffffff1990931692909217909216929092178255519101805491151560ff199092169190911790559383526008825280832093835292905220546117e4908463ffffffff611db016565b6001600160a01b038516600081815260086020908152604080832033845282528083209490945591815260099091522054611825908463ffffffff611db016565b6001600160a01b038516600090815260096020908152604080832093909355338252600a9052205461185d908463ffffffff611db016565b336000818152600a602090815260409182902093909355805191825291810185905281516001600160a01b038716927f04c56a409d50971e45c5a2d96e5d557d2b0f1d66d40f14b141e4c958b0f39b32928290030190a26004805460408051637c287fd560e11b81526001600160a01b0388811694820194909452336024820152604481018790529051929091169163f850ffaa916064808201926020929091908290030181600087803b15801561191457600080fd5b505af1158015611928573d6000803e3d6000fd5b505050506040513d602081101561193e57600080fd5b505161194957600080fd5b5060019392505050565b3390565b6001600160a01b0382166000818152600f602090815260408083203380855290835281842054948452600e835281842090845290915281205490919081106119d05760405162461bcd60e51b81526004018080602001828103825260258152602001806122666025913960400191505060405180910390fd5b6001600160a01b0384166000908152600e6020908152604080832033845290915281208054839081106119ff57fe5b600091825260209091206002909102018054909150436001600160801b039091161115611a5d5760405162461bcd60e51b81526004018080602001828103825260298152602001806122b26029913960400191505060405180910390fd5b6001818101805460ff1916821790556001600160a01b0386166000818152600f602090815260408083203380855290835281842080549096019095558554938352600882528083209483529390529190912054600160801b9091046001600160801b031690611acc9082611e6a565b6001600160a01b038716600081815260086020908152604080832033845282528083209490945591815260099091522054611b0d908263ffffffff611e6a16565b6001600160a01b038716600090815260096020908152604080832093909355338252600a90522054611b45908263ffffffff611e6a16565b336000818152600a60209081526040808320949094556001600160a01b038a168252600b8152838220928252919091522054611b87908263ffffffff611db016565b6001600160a01b0387166000818152600b60209081526040808320338452825280832094909455918152600c9091522054611bc8908263ffffffff611db016565b6001600160a01b0387166000908152600c6020908152604080832093909355338252600d90522054611c00908263ffffffff611db016565b336000908152600d60205260409020558415611ca4576002546040805163e3b99e8560e01b81523360048201526024810184905290516001600160a01b039092169163e3b99e85916044808201926020929091908290030181600087803b158015611c6a57600080fd5b505af1158015611c7e573d6000803e3d6000fd5b505050506040513d6020811015611c9457600080fd5b5051611c9f57600080fd5b611cc1565b600254611cc1906001600160a01b0316338363ffffffff611eac16565b604080513381526020810183905281516001600160a01b038916927fcd1fce47d5ad89dd70b04c75bd6bdb8114d4d4ff7b4393f9fb5937e733ba9582928290030190a250600195945050505050565b6001600160a01b038116611d555760405162461bcd60e51b81526004018080602001828103825260268152602001806121bd6026913960400191505060405180910390fd5b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b0392909216919091179055565b600082820183811015610deb576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b604080516001600160a01b0385811660248301528416604482015260648082018490528251808303909101815260849091019091526020810180516001600160e01b03166323b872dd60e01b179052611e64908590611f03565b50505050565b6000610deb83836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f7700008152506120bb565b604080516001600160a01b038416602482015260448082018490528251808303909101815260649091019091526020810180516001600160e01b031663a9059cbb60e01b179052611efe908490611f03565b505050565b611f15826001600160a01b0316612152565b611f66576040805162461bcd60e51b815260206004820152601f60248201527f5361666545524332303a2063616c6c20746f206e6f6e2d636f6e747261637400604482015290519081900360640190fd5b60006060836001600160a01b0316836040518082805190602001908083835b60208310611fa45780518252601f199092019160209182019101611f85565b6001836020036101000a0380198251168184511680821785525050505050509050019150506000604051808303816000865af19150503d8060008114612006576040519150601f19603f3d011682016040523d82523d6000602084013e61200b565b606091505b509150915081612062576040805162461bcd60e51b815260206004820181905260248201527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564604482015290519081900360640190fd5b805115611e645780806020019051602081101561207e57600080fd5b5051611e645760405162461bcd60e51b815260040180806020018281038252602a81526020018061230c602a913960400191505060405180910390fd5b6000818484111561214a5760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b8381101561210f5781810151838201526020016120f7565b50505050905090810190601f16801561213c5780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b505050900390565b6000813f7fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470818114801590610a4457505015159291505056fe4465706f7369744d616e616765723a2070656e64696e67207265717565737420616c72656164792070726f6365737365644f776e61626c653a206e6577206f776e657220697320746865207a65726f20616464726573736f6e417070726f766528616464726573732c616464726573732c75696e743235362c6279746573294465706f7369744d616e616765723a20536569674d616e6167657220697320616c7265616479207365744465706f7369744d616e616765723a206e2065786365656473206e756d206f662070656e64696e672072657175657374734465706f7369744d616e616765723a206e6f207265717565737420746f2070726f636573734465706f7369744d616e616765723a20616d6f756e74206d757374206e6f74206265207a65726f4465706f7369744d616e616765723a207761697420666f72207769746864726177616c2064656c61794465706f7369744d616e616765723a206f6e6c79206163636570742057544f4e20617070726f76652063616c6c6261636b5361666545524332303a204552433230206f7065726174696f6e20646964206e6f7420737563636565644465706f7369744d616e616765723a206e6f2076616c69642070656e64696e672072657175657374a265627a7a723158209d76339bb8f6fb90cecf52c49b9c22b1b5a6d803116f62bed16df7c2fb1142d764736f6c634300050c00326f6e417070726f766528616464726573732c616464726573732c75696e743235362c627974657329"

// DeployDepositManager deploys a new Ethereum contract, binding an instance of DepositManager to it.
func DeployDepositManager(auth *bind.TransactOpts, backend bind.ContractBackend, wton common.Address, registry common.Address, WITHDRAWAL_DELAY *big.Int) (common.Address, *types.Transaction, *DepositManager, error) {
	parsed, err := abi.JSON(strings.NewReader(DepositManagerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DepositManagerBin), backend, wton, registry, WITHDRAWAL_DELAY)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DepositManager{DepositManagerCaller: DepositManagerCaller{contract: contract}, DepositManagerTransactor: DepositManagerTransactor{contract: contract}, DepositManagerFilterer: DepositManagerFilterer{contract: contract}}, nil
}

// DepositManager is an auto generated Go binding around an Ethereum contract.
type DepositManager struct {
	DepositManagerCaller     // Read-only binding to the contract
	DepositManagerTransactor // Write-only binding to the contract
	DepositManagerFilterer   // Log filterer for contract events
}

// DepositManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type DepositManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DepositManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DepositManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DepositManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DepositManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DepositManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DepositManagerSession struct {
	Contract     *DepositManager   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DepositManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DepositManagerCallerSession struct {
	Contract *DepositManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// DepositManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DepositManagerTransactorSession struct {
	Contract     *DepositManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// DepositManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type DepositManagerRaw struct {
	Contract *DepositManager // Generic contract binding to access the raw methods on
}

// DepositManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DepositManagerCallerRaw struct {
	Contract *DepositManagerCaller // Generic read-only contract binding to access the raw methods on
}

// DepositManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DepositManagerTransactorRaw struct {
	Contract *DepositManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDepositManager creates a new instance of DepositManager, bound to a specific deployed contract.
func NewDepositManager(address common.Address, backend bind.ContractBackend) (*DepositManager, error) {
	contract, err := bindDepositManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DepositManager{DepositManagerCaller: DepositManagerCaller{contract: contract}, DepositManagerTransactor: DepositManagerTransactor{contract: contract}, DepositManagerFilterer: DepositManagerFilterer{contract: contract}}, nil
}

// NewDepositManagerCaller creates a new read-only instance of DepositManager, bound to a specific deployed contract.
func NewDepositManagerCaller(address common.Address, caller bind.ContractCaller) (*DepositManagerCaller, error) {
	contract, err := bindDepositManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DepositManagerCaller{contract: contract}, nil
}

// NewDepositManagerTransactor creates a new write-only instance of DepositManager, bound to a specific deployed contract.
func NewDepositManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*DepositManagerTransactor, error) {
	contract, err := bindDepositManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DepositManagerTransactor{contract: contract}, nil
}

// NewDepositManagerFilterer creates a new log filterer instance of DepositManager, bound to a specific deployed contract.
func NewDepositManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*DepositManagerFilterer, error) {
	contract, err := bindDepositManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DepositManagerFilterer{contract: contract}, nil
}

// bindDepositManager binds a generic wrapper to an already deployed contract.
func bindDepositManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DepositManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DepositManager *DepositManagerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DepositManager.Contract.DepositManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DepositManager *DepositManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DepositManager.Contract.DepositManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DepositManager *DepositManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DepositManager.Contract.DepositManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DepositManager *DepositManagerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DepositManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DepositManager *DepositManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DepositManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DepositManager *DepositManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DepositManager.Contract.contract.Transact(opts, method, params...)
}

// INTERFACEIDONAPPROVE is a free data retrieval call binding the contract method 0x6cd28f9a.
//
// Solidity: function INTERFACE_ID_ON_APPROVE() constant returns(bytes4)
func (_DepositManager *DepositManagerCaller) INTERFACEIDONAPPROVE(opts *bind.CallOpts) ([4]byte, error) {
	var (
		ret0 = new([4]byte)
	)
	out := ret0
	err := _DepositManager.contract.Call(opts, out, "INTERFACE_ID_ON_APPROVE")
	return *ret0, err
}

// INTERFACEIDONAPPROVE is a free data retrieval call binding the contract method 0x6cd28f9a.
//
// Solidity: function INTERFACE_ID_ON_APPROVE() constant returns(bytes4)
func (_DepositManager *DepositManagerSession) INTERFACEIDONAPPROVE() ([4]byte, error) {
	return _DepositManager.Contract.INTERFACEIDONAPPROVE(&_DepositManager.CallOpts)
}

// INTERFACEIDONAPPROVE is a free data retrieval call binding the contract method 0x6cd28f9a.
//
// Solidity: function INTERFACE_ID_ON_APPROVE() constant returns(bytes4)
func (_DepositManager *DepositManagerCallerSession) INTERFACEIDONAPPROVE() ([4]byte, error) {
	return _DepositManager.Contract.INTERFACEIDONAPPROVE(&_DepositManager.CallOpts)
}

// WITHDRAWALDELAY is a free data retrieval call binding the contract method 0x0ebb172a.
//
// Solidity: function WITHDRAWAL_DELAY() constant returns(uint256)
func (_DepositManager *DepositManagerCaller) WITHDRAWALDELAY(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DepositManager.contract.Call(opts, out, "WITHDRAWAL_DELAY")
	return *ret0, err
}

// WITHDRAWALDELAY is a free data retrieval call binding the contract method 0x0ebb172a.
//
// Solidity: function WITHDRAWAL_DELAY() constant returns(uint256)
func (_DepositManager *DepositManagerSession) WITHDRAWALDELAY() (*big.Int, error) {
	return _DepositManager.Contract.WITHDRAWALDELAY(&_DepositManager.CallOpts)
}

// WITHDRAWALDELAY is a free data retrieval call binding the contract method 0x0ebb172a.
//
// Solidity: function WITHDRAWAL_DELAY() constant returns(uint256)
func (_DepositManager *DepositManagerCallerSession) WITHDRAWALDELAY() (*big.Int, error) {
	return _DepositManager.Contract.WITHDRAWALDELAY(&_DepositManager.CallOpts)
}

// AccStaked is a free data retrieval call binding the contract method 0x2d2fab94.
//
// Solidity: function accStaked(address rootchain, address account) constant returns(uint256 wtonAmount)
func (_DepositManager *DepositManagerCaller) AccStaked(opts *bind.CallOpts, rootchain common.Address, account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DepositManager.contract.Call(opts, out, "accStaked", rootchain, account)
	return *ret0, err
}

// AccStaked is a free data retrieval call binding the contract method 0x2d2fab94.
//
// Solidity: function accStaked(address rootchain, address account) constant returns(uint256 wtonAmount)
func (_DepositManager *DepositManagerSession) AccStaked(rootchain common.Address, account common.Address) (*big.Int, error) {
	return _DepositManager.Contract.AccStaked(&_DepositManager.CallOpts, rootchain, account)
}

// AccStaked is a free data retrieval call binding the contract method 0x2d2fab94.
//
// Solidity: function accStaked(address rootchain, address account) constant returns(uint256 wtonAmount)
func (_DepositManager *DepositManagerCallerSession) AccStaked(rootchain common.Address, account common.Address) (*big.Int, error) {
	return _DepositManager.Contract.AccStaked(&_DepositManager.CallOpts, rootchain, account)
}

// AccStakedAccount is a free data retrieval call binding the contract method 0x0055f5c1.
//
// Solidity: function accStakedAccount(address account) constant returns(uint256 wtonAmount)
func (_DepositManager *DepositManagerCaller) AccStakedAccount(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DepositManager.contract.Call(opts, out, "accStakedAccount", account)
	return *ret0, err
}

// AccStakedAccount is a free data retrieval call binding the contract method 0x0055f5c1.
//
// Solidity: function accStakedAccount(address account) constant returns(uint256 wtonAmount)
func (_DepositManager *DepositManagerSession) AccStakedAccount(account common.Address) (*big.Int, error) {
	return _DepositManager.Contract.AccStakedAccount(&_DepositManager.CallOpts, account)
}

// AccStakedAccount is a free data retrieval call binding the contract method 0x0055f5c1.
//
// Solidity: function accStakedAccount(address account) constant returns(uint256 wtonAmount)
func (_DepositManager *DepositManagerCallerSession) AccStakedAccount(account common.Address) (*big.Int, error) {
	return _DepositManager.Contract.AccStakedAccount(&_DepositManager.CallOpts, account)
}

// AccStakedRootChain is a free data retrieval call binding the contract method 0xe8035ec9.
//
// Solidity: function accStakedRootChain(address rootchain) constant returns(uint256 wtonAmount)
func (_DepositManager *DepositManagerCaller) AccStakedRootChain(opts *bind.CallOpts, rootchain common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DepositManager.contract.Call(opts, out, "accStakedRootChain", rootchain)
	return *ret0, err
}

// AccStakedRootChain is a free data retrieval call binding the contract method 0xe8035ec9.
//
// Solidity: function accStakedRootChain(address rootchain) constant returns(uint256 wtonAmount)
func (_DepositManager *DepositManagerSession) AccStakedRootChain(rootchain common.Address) (*big.Int, error) {
	return _DepositManager.Contract.AccStakedRootChain(&_DepositManager.CallOpts, rootchain)
}

// AccStakedRootChain is a free data retrieval call binding the contract method 0xe8035ec9.
//
// Solidity: function accStakedRootChain(address rootchain) constant returns(uint256 wtonAmount)
func (_DepositManager *DepositManagerCallerSession) AccStakedRootChain(rootchain common.Address) (*big.Int, error) {
	return _DepositManager.Contract.AccStakedRootChain(&_DepositManager.CallOpts, rootchain)
}

// AccUnstaked is a free data retrieval call binding the contract method 0x9d91b87b.
//
// Solidity: function accUnstaked(address rootchain, address account) constant returns(uint256 wtonAmount)
func (_DepositManager *DepositManagerCaller) AccUnstaked(opts *bind.CallOpts, rootchain common.Address, account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DepositManager.contract.Call(opts, out, "accUnstaked", rootchain, account)
	return *ret0, err
}

// AccUnstaked is a free data retrieval call binding the contract method 0x9d91b87b.
//
// Solidity: function accUnstaked(address rootchain, address account) constant returns(uint256 wtonAmount)
func (_DepositManager *DepositManagerSession) AccUnstaked(rootchain common.Address, account common.Address) (*big.Int, error) {
	return _DepositManager.Contract.AccUnstaked(&_DepositManager.CallOpts, rootchain, account)
}

// AccUnstaked is a free data retrieval call binding the contract method 0x9d91b87b.
//
// Solidity: function accUnstaked(address rootchain, address account) constant returns(uint256 wtonAmount)
func (_DepositManager *DepositManagerCallerSession) AccUnstaked(rootchain common.Address, account common.Address) (*big.Int, error) {
	return _DepositManager.Contract.AccUnstaked(&_DepositManager.CallOpts, rootchain, account)
}

// AccUnstakedAccount is a free data retrieval call binding the contract method 0xa3543989.
//
// Solidity: function accUnstakedAccount(address account) constant returns(uint256 wtonAmount)
func (_DepositManager *DepositManagerCaller) AccUnstakedAccount(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DepositManager.contract.Call(opts, out, "accUnstakedAccount", account)
	return *ret0, err
}

// AccUnstakedAccount is a free data retrieval call binding the contract method 0xa3543989.
//
// Solidity: function accUnstakedAccount(address account) constant returns(uint256 wtonAmount)
func (_DepositManager *DepositManagerSession) AccUnstakedAccount(account common.Address) (*big.Int, error) {
	return _DepositManager.Contract.AccUnstakedAccount(&_DepositManager.CallOpts, account)
}

// AccUnstakedAccount is a free data retrieval call binding the contract method 0xa3543989.
//
// Solidity: function accUnstakedAccount(address account) constant returns(uint256 wtonAmount)
func (_DepositManager *DepositManagerCallerSession) AccUnstakedAccount(account common.Address) (*big.Int, error) {
	return _DepositManager.Contract.AccUnstakedAccount(&_DepositManager.CallOpts, account)
}

// AccUnstakedRootChain is a free data retrieval call binding the contract method 0x03dc6510.
//
// Solidity: function accUnstakedRootChain(address rootchain) constant returns(uint256 wtonAmount)
func (_DepositManager *DepositManagerCaller) AccUnstakedRootChain(opts *bind.CallOpts, rootchain common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DepositManager.contract.Call(opts, out, "accUnstakedRootChain", rootchain)
	return *ret0, err
}

// AccUnstakedRootChain is a free data retrieval call binding the contract method 0x03dc6510.
//
// Solidity: function accUnstakedRootChain(address rootchain) constant returns(uint256 wtonAmount)
func (_DepositManager *DepositManagerSession) AccUnstakedRootChain(rootchain common.Address) (*big.Int, error) {
	return _DepositManager.Contract.AccUnstakedRootChain(&_DepositManager.CallOpts, rootchain)
}

// AccUnstakedRootChain is a free data retrieval call binding the contract method 0x03dc6510.
//
// Solidity: function accUnstakedRootChain(address rootchain) constant returns(uint256 wtonAmount)
func (_DepositManager *DepositManagerCallerSession) AccUnstakedRootChain(rootchain common.Address) (*big.Int, error) {
	return _DepositManager.Contract.AccUnstakedRootChain(&_DepositManager.CallOpts, rootchain)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_DepositManager *DepositManagerCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _DepositManager.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_DepositManager *DepositManagerSession) IsOwner() (bool, error) {
	return _DepositManager.Contract.IsOwner(&_DepositManager.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_DepositManager *DepositManagerCallerSession) IsOwner() (bool, error) {
	return _DepositManager.Contract.IsOwner(&_DepositManager.CallOpts)
}

// NumPendingRequests is a free data retrieval call binding the contract method 0x5c0df46b.
//
// Solidity: function numPendingRequests(address rootchain, address account) constant returns(uint256)
func (_DepositManager *DepositManagerCaller) NumPendingRequests(opts *bind.CallOpts, rootchain common.Address, account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DepositManager.contract.Call(opts, out, "numPendingRequests", rootchain, account)
	return *ret0, err
}

// NumPendingRequests is a free data retrieval call binding the contract method 0x5c0df46b.
//
// Solidity: function numPendingRequests(address rootchain, address account) constant returns(uint256)
func (_DepositManager *DepositManagerSession) NumPendingRequests(rootchain common.Address, account common.Address) (*big.Int, error) {
	return _DepositManager.Contract.NumPendingRequests(&_DepositManager.CallOpts, rootchain, account)
}

// NumPendingRequests is a free data retrieval call binding the contract method 0x5c0df46b.
//
// Solidity: function numPendingRequests(address rootchain, address account) constant returns(uint256)
func (_DepositManager *DepositManagerCallerSession) NumPendingRequests(rootchain common.Address, account common.Address) (*big.Int, error) {
	return _DepositManager.Contract.NumPendingRequests(&_DepositManager.CallOpts, rootchain, account)
}

// NumRequests is a free data retrieval call binding the contract method 0xf762eb57.
//
// Solidity: function numRequests(address rootchain, address account) constant returns(uint256)
func (_DepositManager *DepositManagerCaller) NumRequests(opts *bind.CallOpts, rootchain common.Address, account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DepositManager.contract.Call(opts, out, "numRequests", rootchain, account)
	return *ret0, err
}

// NumRequests is a free data retrieval call binding the contract method 0xf762eb57.
//
// Solidity: function numRequests(address rootchain, address account) constant returns(uint256)
func (_DepositManager *DepositManagerSession) NumRequests(rootchain common.Address, account common.Address) (*big.Int, error) {
	return _DepositManager.Contract.NumRequests(&_DepositManager.CallOpts, rootchain, account)
}

// NumRequests is a free data retrieval call binding the contract method 0xf762eb57.
//
// Solidity: function numRequests(address rootchain, address account) constant returns(uint256)
func (_DepositManager *DepositManagerCallerSession) NumRequests(rootchain common.Address, account common.Address) (*big.Int, error) {
	return _DepositManager.Contract.NumRequests(&_DepositManager.CallOpts, rootchain, account)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_DepositManager *DepositManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DepositManager.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_DepositManager *DepositManagerSession) Owner() (common.Address, error) {
	return _DepositManager.Contract.Owner(&_DepositManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_DepositManager *DepositManagerCallerSession) Owner() (common.Address, error) {
	return _DepositManager.Contract.Owner(&_DepositManager.CallOpts)
}

// PendingUnstaked is a free data retrieval call binding the contract method 0x2638fdf5.
//
// Solidity: function pendingUnstaked(address rootchain, address account) constant returns(uint256 wtonAmount)
func (_DepositManager *DepositManagerCaller) PendingUnstaked(opts *bind.CallOpts, rootchain common.Address, account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DepositManager.contract.Call(opts, out, "pendingUnstaked", rootchain, account)
	return *ret0, err
}

// PendingUnstaked is a free data retrieval call binding the contract method 0x2638fdf5.
//
// Solidity: function pendingUnstaked(address rootchain, address account) constant returns(uint256 wtonAmount)
func (_DepositManager *DepositManagerSession) PendingUnstaked(rootchain common.Address, account common.Address) (*big.Int, error) {
	return _DepositManager.Contract.PendingUnstaked(&_DepositManager.CallOpts, rootchain, account)
}

// PendingUnstaked is a free data retrieval call binding the contract method 0x2638fdf5.
//
// Solidity: function pendingUnstaked(address rootchain, address account) constant returns(uint256 wtonAmount)
func (_DepositManager *DepositManagerCallerSession) PendingUnstaked(rootchain common.Address, account common.Address) (*big.Int, error) {
	return _DepositManager.Contract.PendingUnstaked(&_DepositManager.CallOpts, rootchain, account)
}

// PendingUnstakedAccount is a free data retrieval call binding the contract method 0xa0b2a913.
//
// Solidity: function pendingUnstakedAccount(address account) constant returns(uint256 wtonAmount)
func (_DepositManager *DepositManagerCaller) PendingUnstakedAccount(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DepositManager.contract.Call(opts, out, "pendingUnstakedAccount", account)
	return *ret0, err
}

// PendingUnstakedAccount is a free data retrieval call binding the contract method 0xa0b2a913.
//
// Solidity: function pendingUnstakedAccount(address account) constant returns(uint256 wtonAmount)
func (_DepositManager *DepositManagerSession) PendingUnstakedAccount(account common.Address) (*big.Int, error) {
	return _DepositManager.Contract.PendingUnstakedAccount(&_DepositManager.CallOpts, account)
}

// PendingUnstakedAccount is a free data retrieval call binding the contract method 0xa0b2a913.
//
// Solidity: function pendingUnstakedAccount(address account) constant returns(uint256 wtonAmount)
func (_DepositManager *DepositManagerCallerSession) PendingUnstakedAccount(account common.Address) (*big.Int, error) {
	return _DepositManager.Contract.PendingUnstakedAccount(&_DepositManager.CallOpts, account)
}

// PendingUnstakedRootChain is a free data retrieval call binding the contract method 0xa8f3fb98.
//
// Solidity: function pendingUnstakedRootChain(address rootchain) constant returns(uint256 wtonAmount)
func (_DepositManager *DepositManagerCaller) PendingUnstakedRootChain(opts *bind.CallOpts, rootchain common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DepositManager.contract.Call(opts, out, "pendingUnstakedRootChain", rootchain)
	return *ret0, err
}

// PendingUnstakedRootChain is a free data retrieval call binding the contract method 0xa8f3fb98.
//
// Solidity: function pendingUnstakedRootChain(address rootchain) constant returns(uint256 wtonAmount)
func (_DepositManager *DepositManagerSession) PendingUnstakedRootChain(rootchain common.Address) (*big.Int, error) {
	return _DepositManager.Contract.PendingUnstakedRootChain(&_DepositManager.CallOpts, rootchain)
}

// PendingUnstakedRootChain is a free data retrieval call binding the contract method 0xa8f3fb98.
//
// Solidity: function pendingUnstakedRootChain(address rootchain) constant returns(uint256 wtonAmount)
func (_DepositManager *DepositManagerCallerSession) PendingUnstakedRootChain(rootchain common.Address) (*big.Int, error) {
	return _DepositManager.Contract.PendingUnstakedRootChain(&_DepositManager.CallOpts, rootchain)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() constant returns(address)
func (_DepositManager *DepositManagerCaller) Registry(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DepositManager.contract.Call(opts, out, "registry")
	return *ret0, err
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() constant returns(address)
func (_DepositManager *DepositManagerSession) Registry() (common.Address, error) {
	return _DepositManager.Contract.Registry(&_DepositManager.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() constant returns(address)
func (_DepositManager *DepositManagerCallerSession) Registry() (common.Address, error) {
	return _DepositManager.Contract.Registry(&_DepositManager.CallOpts)
}

// SeigManager is a free data retrieval call binding the contract method 0x6fb7f558.
//
// Solidity: function seigManager() constant returns(address)
func (_DepositManager *DepositManagerCaller) SeigManager(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DepositManager.contract.Call(opts, out, "seigManager")
	return *ret0, err
}

// SeigManager is a free data retrieval call binding the contract method 0x6fb7f558.
//
// Solidity: function seigManager() constant returns(address)
func (_DepositManager *DepositManagerSession) SeigManager() (common.Address, error) {
	return _DepositManager.Contract.SeigManager(&_DepositManager.CallOpts)
}

// SeigManager is a free data retrieval call binding the contract method 0x6fb7f558.
//
// Solidity: function seigManager() constant returns(address)
func (_DepositManager *DepositManagerCallerSession) SeigManager() (common.Address, error) {
	return _DepositManager.Contract.SeigManager(&_DepositManager.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) constant returns(bool)
func (_DepositManager *DepositManagerCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _DepositManager.contract.Call(opts, out, "supportsInterface", interfaceId)
	return *ret0, err
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) constant returns(bool)
func (_DepositManager *DepositManagerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _DepositManager.Contract.SupportsInterface(&_DepositManager.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) constant returns(bool)
func (_DepositManager *DepositManagerCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _DepositManager.Contract.SupportsInterface(&_DepositManager.CallOpts, interfaceId)
}

// WithdrawalRequest is a free data retrieval call binding the contract method 0x8fbef2d0.
//
// Solidity: function withdrawalRequest(address rootchain, address account, uint256 index) constant returns(uint128 withdrawableBlockNumber, uint128 amount, bool processed)
func (_DepositManager *DepositManagerCaller) WithdrawalRequest(opts *bind.CallOpts, rootchain common.Address, account common.Address, index *big.Int) (struct {
	WithdrawableBlockNumber *big.Int
	Amount                  *big.Int
	Processed               bool
}, error) {
	ret := new(struct {
		WithdrawableBlockNumber *big.Int
		Amount                  *big.Int
		Processed               bool
	})
	out := ret
	err := _DepositManager.contract.Call(opts, out, "withdrawalRequest", rootchain, account, index)
	return *ret, err
}

// WithdrawalRequest is a free data retrieval call binding the contract method 0x8fbef2d0.
//
// Solidity: function withdrawalRequest(address rootchain, address account, uint256 index) constant returns(uint128 withdrawableBlockNumber, uint128 amount, bool processed)
func (_DepositManager *DepositManagerSession) WithdrawalRequest(rootchain common.Address, account common.Address, index *big.Int) (struct {
	WithdrawableBlockNumber *big.Int
	Amount                  *big.Int
	Processed               bool
}, error) {
	return _DepositManager.Contract.WithdrawalRequest(&_DepositManager.CallOpts, rootchain, account, index)
}

// WithdrawalRequest is a free data retrieval call binding the contract method 0x8fbef2d0.
//
// Solidity: function withdrawalRequest(address rootchain, address account, uint256 index) constant returns(uint128 withdrawableBlockNumber, uint128 amount, bool processed)
func (_DepositManager *DepositManagerCallerSession) WithdrawalRequest(rootchain common.Address, account common.Address, index *big.Int) (struct {
	WithdrawableBlockNumber *big.Int
	Amount                  *big.Int
	Processed               bool
}, error) {
	return _DepositManager.Contract.WithdrawalRequest(&_DepositManager.CallOpts, rootchain, account, index)
}

// WithdrawalRequestIndex is a free data retrieval call binding the contract method 0xc647f26e.
//
// Solidity: function withdrawalRequestIndex(address rootchain, address account) constant returns(uint256 index)
func (_DepositManager *DepositManagerCaller) WithdrawalRequestIndex(opts *bind.CallOpts, rootchain common.Address, account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DepositManager.contract.Call(opts, out, "withdrawalRequestIndex", rootchain, account)
	return *ret0, err
}

// WithdrawalRequestIndex is a free data retrieval call binding the contract method 0xc647f26e.
//
// Solidity: function withdrawalRequestIndex(address rootchain, address account) constant returns(uint256 index)
func (_DepositManager *DepositManagerSession) WithdrawalRequestIndex(rootchain common.Address, account common.Address) (*big.Int, error) {
	return _DepositManager.Contract.WithdrawalRequestIndex(&_DepositManager.CallOpts, rootchain, account)
}

// WithdrawalRequestIndex is a free data retrieval call binding the contract method 0xc647f26e.
//
// Solidity: function withdrawalRequestIndex(address rootchain, address account) constant returns(uint256 index)
func (_DepositManager *DepositManagerCallerSession) WithdrawalRequestIndex(rootchain common.Address, account common.Address) (*big.Int, error) {
	return _DepositManager.Contract.WithdrawalRequestIndex(&_DepositManager.CallOpts, rootchain, account)
}

// Wton is a free data retrieval call binding the contract method 0x8d62d949.
//
// Solidity: function wton() constant returns(address)
func (_DepositManager *DepositManagerCaller) Wton(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DepositManager.contract.Call(opts, out, "wton")
	return *ret0, err
}

// Wton is a free data retrieval call binding the contract method 0x8d62d949.
//
// Solidity: function wton() constant returns(address)
func (_DepositManager *DepositManagerSession) Wton() (common.Address, error) {
	return _DepositManager.Contract.Wton(&_DepositManager.CallOpts)
}

// Wton is a free data retrieval call binding the contract method 0x8d62d949.
//
// Solidity: function wton() constant returns(address)
func (_DepositManager *DepositManagerCallerSession) Wton() (common.Address, error) {
	return _DepositManager.Contract.Wton(&_DepositManager.CallOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address rootchain, uint256 amount) returns(bool)
func (_DepositManager *DepositManagerTransactor) Deposit(opts *bind.TransactOpts, rootchain common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DepositManager.contract.Transact(opts, "deposit", rootchain, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address rootchain, uint256 amount) returns(bool)
func (_DepositManager *DepositManagerSession) Deposit(rootchain common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DepositManager.Contract.Deposit(&_DepositManager.TransactOpts, rootchain, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address rootchain, uint256 amount) returns(bool)
func (_DepositManager *DepositManagerTransactorSession) Deposit(rootchain common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DepositManager.Contract.Deposit(&_DepositManager.TransactOpts, rootchain, amount)
}

// OnApprove is a paid mutator transaction binding the contract method 0x4273ca16.
//
// Solidity: function onApprove(address owner, address spender, uint256 amount, bytes data) returns(bool)
func (_DepositManager *DepositManagerTransactor) OnApprove(opts *bind.TransactOpts, owner common.Address, spender common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _DepositManager.contract.Transact(opts, "onApprove", owner, spender, amount, data)
}

// OnApprove is a paid mutator transaction binding the contract method 0x4273ca16.
//
// Solidity: function onApprove(address owner, address spender, uint256 amount, bytes data) returns(bool)
func (_DepositManager *DepositManagerSession) OnApprove(owner common.Address, spender common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _DepositManager.Contract.OnApprove(&_DepositManager.TransactOpts, owner, spender, amount, data)
}

// OnApprove is a paid mutator transaction binding the contract method 0x4273ca16.
//
// Solidity: function onApprove(address owner, address spender, uint256 amount, bytes data) returns(bool)
func (_DepositManager *DepositManagerTransactorSession) OnApprove(owner common.Address, spender common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _DepositManager.Contract.OnApprove(&_DepositManager.TransactOpts, owner, spender, amount, data)
}

// ProcessRequest is a paid mutator transaction binding the contract method 0xc20a44c6.
//
// Solidity: function processRequest(address rootchain, bool receiveTON) returns(bool)
func (_DepositManager *DepositManagerTransactor) ProcessRequest(opts *bind.TransactOpts, rootchain common.Address, receiveTON bool) (*types.Transaction, error) {
	return _DepositManager.contract.Transact(opts, "processRequest", rootchain, receiveTON)
}

// ProcessRequest is a paid mutator transaction binding the contract method 0xc20a44c6.
//
// Solidity: function processRequest(address rootchain, bool receiveTON) returns(bool)
func (_DepositManager *DepositManagerSession) ProcessRequest(rootchain common.Address, receiveTON bool) (*types.Transaction, error) {
	return _DepositManager.Contract.ProcessRequest(&_DepositManager.TransactOpts, rootchain, receiveTON)
}

// ProcessRequest is a paid mutator transaction binding the contract method 0xc20a44c6.
//
// Solidity: function processRequest(address rootchain, bool receiveTON) returns(bool)
func (_DepositManager *DepositManagerTransactorSession) ProcessRequest(rootchain common.Address, receiveTON bool) (*types.Transaction, error) {
	return _DepositManager.Contract.ProcessRequest(&_DepositManager.TransactOpts, rootchain, receiveTON)
}

// ProcessRequests is a paid mutator transaction binding the contract method 0xfb0713b1.
//
// Solidity: function processRequests(address rootchain, uint256 n, bool receiveTON) returns(bool)
func (_DepositManager *DepositManagerTransactor) ProcessRequests(opts *bind.TransactOpts, rootchain common.Address, n *big.Int, receiveTON bool) (*types.Transaction, error) {
	return _DepositManager.contract.Transact(opts, "processRequests", rootchain, n, receiveTON)
}

// ProcessRequests is a paid mutator transaction binding the contract method 0xfb0713b1.
//
// Solidity: function processRequests(address rootchain, uint256 n, bool receiveTON) returns(bool)
func (_DepositManager *DepositManagerSession) ProcessRequests(rootchain common.Address, n *big.Int, receiveTON bool) (*types.Transaction, error) {
	return _DepositManager.Contract.ProcessRequests(&_DepositManager.TransactOpts, rootchain, n, receiveTON)
}

// ProcessRequests is a paid mutator transaction binding the contract method 0xfb0713b1.
//
// Solidity: function processRequests(address rootchain, uint256 n, bool receiveTON) returns(bool)
func (_DepositManager *DepositManagerTransactorSession) ProcessRequests(rootchain common.Address, n *big.Int, receiveTON bool) (*types.Transaction, error) {
	return _DepositManager.Contract.ProcessRequests(&_DepositManager.TransactOpts, rootchain, n, receiveTON)
}

// Redeposit is a paid mutator transaction binding the contract method 0xb8bee628.
//
// Solidity: function redeposit(address rootchain) returns(bool)
func (_DepositManager *DepositManagerTransactor) Redeposit(opts *bind.TransactOpts, rootchain common.Address) (*types.Transaction, error) {
	return _DepositManager.contract.Transact(opts, "redeposit", rootchain)
}

// Redeposit is a paid mutator transaction binding the contract method 0xb8bee628.
//
// Solidity: function redeposit(address rootchain) returns(bool)
func (_DepositManager *DepositManagerSession) Redeposit(rootchain common.Address) (*types.Transaction, error) {
	return _DepositManager.Contract.Redeposit(&_DepositManager.TransactOpts, rootchain)
}

// Redeposit is a paid mutator transaction binding the contract method 0xb8bee628.
//
// Solidity: function redeposit(address rootchain) returns(bool)
func (_DepositManager *DepositManagerTransactorSession) Redeposit(rootchain common.Address) (*types.Transaction, error) {
	return _DepositManager.Contract.Redeposit(&_DepositManager.TransactOpts, rootchain)
}

// RedepositMulti is a paid mutator transaction binding the contract method 0x445e83b9.
//
// Solidity: function redepositMulti(address rootchain, uint256 n) returns(bool)
func (_DepositManager *DepositManagerTransactor) RedepositMulti(opts *bind.TransactOpts, rootchain common.Address, n *big.Int) (*types.Transaction, error) {
	return _DepositManager.contract.Transact(opts, "redepositMulti", rootchain, n)
}

// RedepositMulti is a paid mutator transaction binding the contract method 0x445e83b9.
//
// Solidity: function redepositMulti(address rootchain, uint256 n) returns(bool)
func (_DepositManager *DepositManagerSession) RedepositMulti(rootchain common.Address, n *big.Int) (*types.Transaction, error) {
	return _DepositManager.Contract.RedepositMulti(&_DepositManager.TransactOpts, rootchain, n)
}

// RedepositMulti is a paid mutator transaction binding the contract method 0x445e83b9.
//
// Solidity: function redepositMulti(address rootchain, uint256 n) returns(bool)
func (_DepositManager *DepositManagerTransactorSession) RedepositMulti(rootchain common.Address, n *big.Int) (*types.Transaction, error) {
	return _DepositManager.Contract.RedepositMulti(&_DepositManager.TransactOpts, rootchain, n)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DepositManager *DepositManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DepositManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DepositManager *DepositManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _DepositManager.Contract.RenounceOwnership(&_DepositManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DepositManager *DepositManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _DepositManager.Contract.RenounceOwnership(&_DepositManager.TransactOpts)
}

// RequestWithdrawal is a paid mutator transaction binding the contract method 0xda95ebf7.
//
// Solidity: function requestWithdrawal(address rootchain, uint256 amount) returns(bool)
func (_DepositManager *DepositManagerTransactor) RequestWithdrawal(opts *bind.TransactOpts, rootchain common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DepositManager.contract.Transact(opts, "requestWithdrawal", rootchain, amount)
}

// RequestWithdrawal is a paid mutator transaction binding the contract method 0xda95ebf7.
//
// Solidity: function requestWithdrawal(address rootchain, uint256 amount) returns(bool)
func (_DepositManager *DepositManagerSession) RequestWithdrawal(rootchain common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DepositManager.Contract.RequestWithdrawal(&_DepositManager.TransactOpts, rootchain, amount)
}

// RequestWithdrawal is a paid mutator transaction binding the contract method 0xda95ebf7.
//
// Solidity: function requestWithdrawal(address rootchain, uint256 amount) returns(bool)
func (_DepositManager *DepositManagerTransactorSession) RequestWithdrawal(rootchain common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DepositManager.Contract.RequestWithdrawal(&_DepositManager.TransactOpts, rootchain, amount)
}

// RequestWithdrawalAll is a paid mutator transaction binding the contract method 0x6b2160b7.
//
// Solidity: function requestWithdrawalAll(address rootchain) returns(bool)
func (_DepositManager *DepositManagerTransactor) RequestWithdrawalAll(opts *bind.TransactOpts, rootchain common.Address) (*types.Transaction, error) {
	return _DepositManager.contract.Transact(opts, "requestWithdrawalAll", rootchain)
}

// RequestWithdrawalAll is a paid mutator transaction binding the contract method 0x6b2160b7.
//
// Solidity: function requestWithdrawalAll(address rootchain) returns(bool)
func (_DepositManager *DepositManagerSession) RequestWithdrawalAll(rootchain common.Address) (*types.Transaction, error) {
	return _DepositManager.Contract.RequestWithdrawalAll(&_DepositManager.TransactOpts, rootchain)
}

// RequestWithdrawalAll is a paid mutator transaction binding the contract method 0x6b2160b7.
//
// Solidity: function requestWithdrawalAll(address rootchain) returns(bool)
func (_DepositManager *DepositManagerTransactorSession) RequestWithdrawalAll(rootchain common.Address) (*types.Transaction, error) {
	return _DepositManager.Contract.RequestWithdrawalAll(&_DepositManager.TransactOpts, rootchain)
}

// SetSeigManager is a paid mutator transaction binding the contract method 0x7657f20a.
//
// Solidity: function setSeigManager(address seigManager) returns()
func (_DepositManager *DepositManagerTransactor) SetSeigManager(opts *bind.TransactOpts, seigManager common.Address) (*types.Transaction, error) {
	return _DepositManager.contract.Transact(opts, "setSeigManager", seigManager)
}

// SetSeigManager is a paid mutator transaction binding the contract method 0x7657f20a.
//
// Solidity: function setSeigManager(address seigManager) returns()
func (_DepositManager *DepositManagerSession) SetSeigManager(seigManager common.Address) (*types.Transaction, error) {
	return _DepositManager.Contract.SetSeigManager(&_DepositManager.TransactOpts, seigManager)
}

// SetSeigManager is a paid mutator transaction binding the contract method 0x7657f20a.
//
// Solidity: function setSeigManager(address seigManager) returns()
func (_DepositManager *DepositManagerTransactorSession) SetSeigManager(seigManager common.Address) (*types.Transaction, error) {
	return _DepositManager.Contract.SetSeigManager(&_DepositManager.TransactOpts, seigManager)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DepositManager *DepositManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _DepositManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DepositManager *DepositManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DepositManager.Contract.TransferOwnership(&_DepositManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DepositManager *DepositManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DepositManager.Contract.TransferOwnership(&_DepositManager.TransactOpts, newOwner)
}

// DepositManagerDepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the DepositManager contract.
type DepositManagerDepositedIterator struct {
	Event *DepositManagerDeposited // Event containing the contract specifics and raw log

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
func (it *DepositManagerDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DepositManagerDeposited)
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
		it.Event = new(DepositManagerDeposited)
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
func (it *DepositManagerDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DepositManagerDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DepositManagerDeposited represents a Deposited event raised by the DepositManager contract.
type DepositManagerDeposited struct {
	Rootchain common.Address
	Depositor common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDeposited is a free log retrieval operation binding the contract event 0x8752a472e571a816aea92eec8dae9baf628e840f4929fbcc2d155e6233ff68a7.
//
// Solidity: event Deposited(address indexed rootchain, address depositor, uint256 amount)
func (_DepositManager *DepositManagerFilterer) FilterDeposited(opts *bind.FilterOpts, rootchain []common.Address) (*DepositManagerDepositedIterator, error) {

	var rootchainRule []interface{}
	for _, rootchainItem := range rootchain {
		rootchainRule = append(rootchainRule, rootchainItem)
	}

	logs, sub, err := _DepositManager.contract.FilterLogs(opts, "Deposited", rootchainRule)
	if err != nil {
		return nil, err
	}
	return &DepositManagerDepositedIterator{contract: _DepositManager.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0x8752a472e571a816aea92eec8dae9baf628e840f4929fbcc2d155e6233ff68a7.
//
// Solidity: event Deposited(address indexed rootchain, address depositor, uint256 amount)
func (_DepositManager *DepositManagerFilterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *DepositManagerDeposited, rootchain []common.Address) (event.Subscription, error) {

	var rootchainRule []interface{}
	for _, rootchainItem := range rootchain {
		rootchainRule = append(rootchainRule, rootchainItem)
	}

	logs, sub, err := _DepositManager.contract.WatchLogs(opts, "Deposited", rootchainRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DepositManagerDeposited)
				if err := _DepositManager.contract.UnpackLog(event, "Deposited", log); err != nil {
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

// ParseDeposited is a log parse operation binding the contract event 0x8752a472e571a816aea92eec8dae9baf628e840f4929fbcc2d155e6233ff68a7.
//
// Solidity: event Deposited(address indexed rootchain, address depositor, uint256 amount)
func (_DepositManager *DepositManagerFilterer) ParseDeposited(log types.Log) (*DepositManagerDeposited, error) {
	event := new(DepositManagerDeposited)
	if err := _DepositManager.contract.UnpackLog(event, "Deposited", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DepositManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the DepositManager contract.
type DepositManagerOwnershipTransferredIterator struct {
	Event *DepositManagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *DepositManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DepositManagerOwnershipTransferred)
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
		it.Event = new(DepositManagerOwnershipTransferred)
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
func (it *DepositManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DepositManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DepositManagerOwnershipTransferred represents a OwnershipTransferred event raised by the DepositManager contract.
type DepositManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DepositManager *DepositManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*DepositManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DepositManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DepositManagerOwnershipTransferredIterator{contract: _DepositManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DepositManager *DepositManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DepositManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DepositManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DepositManagerOwnershipTransferred)
				if err := _DepositManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_DepositManager *DepositManagerFilterer) ParseOwnershipTransferred(log types.Log) (*DepositManagerOwnershipTransferred, error) {
	event := new(DepositManagerOwnershipTransferred)
	if err := _DepositManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DepositManagerWithdrawalProcessedIterator is returned from FilterWithdrawalProcessed and is used to iterate over the raw logs and unpacked data for WithdrawalProcessed events raised by the DepositManager contract.
type DepositManagerWithdrawalProcessedIterator struct {
	Event *DepositManagerWithdrawalProcessed // Event containing the contract specifics and raw log

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
func (it *DepositManagerWithdrawalProcessedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DepositManagerWithdrawalProcessed)
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
		it.Event = new(DepositManagerWithdrawalProcessed)
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
func (it *DepositManagerWithdrawalProcessedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DepositManagerWithdrawalProcessedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DepositManagerWithdrawalProcessed represents a WithdrawalProcessed event raised by the DepositManager contract.
type DepositManagerWithdrawalProcessed struct {
	Rootchain common.Address
	Depositor common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalProcessed is a free log retrieval operation binding the contract event 0xcd1fce47d5ad89dd70b04c75bd6bdb8114d4d4ff7b4393f9fb5937e733ba9582.
//
// Solidity: event WithdrawalProcessed(address indexed rootchain, address depositor, uint256 amount)
func (_DepositManager *DepositManagerFilterer) FilterWithdrawalProcessed(opts *bind.FilterOpts, rootchain []common.Address) (*DepositManagerWithdrawalProcessedIterator, error) {

	var rootchainRule []interface{}
	for _, rootchainItem := range rootchain {
		rootchainRule = append(rootchainRule, rootchainItem)
	}

	logs, sub, err := _DepositManager.contract.FilterLogs(opts, "WithdrawalProcessed", rootchainRule)
	if err != nil {
		return nil, err
	}
	return &DepositManagerWithdrawalProcessedIterator{contract: _DepositManager.contract, event: "WithdrawalProcessed", logs: logs, sub: sub}, nil
}

// WatchWithdrawalProcessed is a free log subscription operation binding the contract event 0xcd1fce47d5ad89dd70b04c75bd6bdb8114d4d4ff7b4393f9fb5937e733ba9582.
//
// Solidity: event WithdrawalProcessed(address indexed rootchain, address depositor, uint256 amount)
func (_DepositManager *DepositManagerFilterer) WatchWithdrawalProcessed(opts *bind.WatchOpts, sink chan<- *DepositManagerWithdrawalProcessed, rootchain []common.Address) (event.Subscription, error) {

	var rootchainRule []interface{}
	for _, rootchainItem := range rootchain {
		rootchainRule = append(rootchainRule, rootchainItem)
	}

	logs, sub, err := _DepositManager.contract.WatchLogs(opts, "WithdrawalProcessed", rootchainRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DepositManagerWithdrawalProcessed)
				if err := _DepositManager.contract.UnpackLog(event, "WithdrawalProcessed", log); err != nil {
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

// ParseWithdrawalProcessed is a log parse operation binding the contract event 0xcd1fce47d5ad89dd70b04c75bd6bdb8114d4d4ff7b4393f9fb5937e733ba9582.
//
// Solidity: event WithdrawalProcessed(address indexed rootchain, address depositor, uint256 amount)
func (_DepositManager *DepositManagerFilterer) ParseWithdrawalProcessed(log types.Log) (*DepositManagerWithdrawalProcessed, error) {
	event := new(DepositManagerWithdrawalProcessed)
	if err := _DepositManager.contract.UnpackLog(event, "WithdrawalProcessed", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DepositManagerWithdrawalRequestedIterator is returned from FilterWithdrawalRequested and is used to iterate over the raw logs and unpacked data for WithdrawalRequested events raised by the DepositManager contract.
type DepositManagerWithdrawalRequestedIterator struct {
	Event *DepositManagerWithdrawalRequested // Event containing the contract specifics and raw log

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
func (it *DepositManagerWithdrawalRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DepositManagerWithdrawalRequested)
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
		it.Event = new(DepositManagerWithdrawalRequested)
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
func (it *DepositManagerWithdrawalRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DepositManagerWithdrawalRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DepositManagerWithdrawalRequested represents a WithdrawalRequested event raised by the DepositManager contract.
type DepositManagerWithdrawalRequested struct {
	Rootchain common.Address
	Depositor common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalRequested is a free log retrieval operation binding the contract event 0x04c56a409d50971e45c5a2d96e5d557d2b0f1d66d40f14b141e4c958b0f39b32.
//
// Solidity: event WithdrawalRequested(address indexed rootchain, address depositor, uint256 amount)
func (_DepositManager *DepositManagerFilterer) FilterWithdrawalRequested(opts *bind.FilterOpts, rootchain []common.Address) (*DepositManagerWithdrawalRequestedIterator, error) {

	var rootchainRule []interface{}
	for _, rootchainItem := range rootchain {
		rootchainRule = append(rootchainRule, rootchainItem)
	}

	logs, sub, err := _DepositManager.contract.FilterLogs(opts, "WithdrawalRequested", rootchainRule)
	if err != nil {
		return nil, err
	}
	return &DepositManagerWithdrawalRequestedIterator{contract: _DepositManager.contract, event: "WithdrawalRequested", logs: logs, sub: sub}, nil
}

// WatchWithdrawalRequested is a free log subscription operation binding the contract event 0x04c56a409d50971e45c5a2d96e5d557d2b0f1d66d40f14b141e4c958b0f39b32.
//
// Solidity: event WithdrawalRequested(address indexed rootchain, address depositor, uint256 amount)
func (_DepositManager *DepositManagerFilterer) WatchWithdrawalRequested(opts *bind.WatchOpts, sink chan<- *DepositManagerWithdrawalRequested, rootchain []common.Address) (event.Subscription, error) {

	var rootchainRule []interface{}
	for _, rootchainItem := range rootchain {
		rootchainRule = append(rootchainRule, rootchainItem)
	}

	logs, sub, err := _DepositManager.contract.WatchLogs(opts, "WithdrawalRequested", rootchainRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DepositManagerWithdrawalRequested)
				if err := _DepositManager.contract.UnpackLog(event, "WithdrawalRequested", log); err != nil {
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

// ParseWithdrawalRequested is a log parse operation binding the contract event 0x04c56a409d50971e45c5a2d96e5d557d2b0f1d66d40f14b141e4c958b0f39b32.
//
// Solidity: event WithdrawalRequested(address indexed rootchain, address depositor, uint256 amount)
func (_DepositManager *DepositManagerFilterer) ParseWithdrawalRequested(log types.Log) (*DepositManagerWithdrawalRequested, error) {
	event := new(DepositManagerWithdrawalRequested)
	if err := _DepositManager.contract.UnpackLog(event, "WithdrawalRequested", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DepositManagerIABI is the input ABI used to generate the binding from.
const DepositManagerIABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"WITHDRAWAL_DELAY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"rootchain\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"accStaked\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"wtonAmount\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"accStakedAccount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"wtonAmount\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"rootchain\",\"type\":\"address\"}],\"name\":\"accStakedRootChain\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"wtonAmount\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"rootchain\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"accUnstaked\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"wtonAmount\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"accUnstakedAccount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"wtonAmount\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"rootchain\",\"type\":\"address\"}],\"name\":\"accUnstakedRootChain\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"wtonAmount\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"rootchain\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"rootchain\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"numPendingRequests\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"rootchain\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"numRequests\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"rootchain\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"pendingUnstaked\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"wtonAmount\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"pendingUnstakedAccount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"wtonAmount\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"rootchain\",\"type\":\"address\"}],\"name\":\"pendingUnstakedRootChain\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"wtonAmount\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"rootchain\",\"type\":\"address\"}],\"name\":\"processRequest\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"rootchain\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"n\",\"type\":\"uint256\"}],\"name\":\"processRequests\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"registry\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"rootchain\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"requestWithdrawal\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"rootchain\",\"type\":\"address\"}],\"name\":\"requestWithdrawalAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"seigManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"seigManager\",\"type\":\"address\"}],\"name\":\"setSeigManager\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"rootchain\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"withdrawalRequest\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"withdrawableBlockNumber\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"},{\"internalType\":\"bool\",\"name\":\"processed\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"rootchain\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"withdrawalRequestIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"wton\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// DepositManagerIFuncSigs maps the 4-byte function signature to its string representation.
var DepositManagerIFuncSigs = map[string]string{
	"0ebb172a": "WITHDRAWAL_DELAY()",
	"2d2fab94": "accStaked(address,address)",
	"0055f5c1": "accStakedAccount(address)",
	"e8035ec9": "accStakedRootChain(address)",
	"9d91b87b": "accUnstaked(address,address)",
	"a3543989": "accUnstakedAccount(address)",
	"03dc6510": "accUnstakedRootChain(address)",
	"47e7ef24": "deposit(address,uint256)",
	"5c0df46b": "numPendingRequests(address,address)",
	"f762eb57": "numRequests(address,address)",
	"8da5cb5b": "owner()",
	"2638fdf5": "pendingUnstaked(address,address)",
	"a0b2a913": "pendingUnstakedAccount(address)",
	"a8f3fb98": "pendingUnstakedRootChain(address)",
	"5d6f7cca": "processRequest(address)",
	"06260ceb": "processRequests(address,uint256)",
	"7b103999": "registry()",
	"da95ebf7": "requestWithdrawal(address,uint256)",
	"6b2160b7": "requestWithdrawalAll(address)",
	"6fb7f558": "seigManager()",
	"7657f20a": "setSeigManager(address)",
	"8fbef2d0": "withdrawalRequest(address,address,uint256)",
	"c647f26e": "withdrawalRequestIndex(address,address)",
	"8d62d949": "wton()",
}

// DepositManagerI is an auto generated Go binding around an Ethereum contract.
type DepositManagerI struct {
	DepositManagerICaller     // Read-only binding to the contract
	DepositManagerITransactor // Write-only binding to the contract
	DepositManagerIFilterer   // Log filterer for contract events
}

// DepositManagerICaller is an auto generated read-only Go binding around an Ethereum contract.
type DepositManagerICaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DepositManagerITransactor is an auto generated write-only Go binding around an Ethereum contract.
type DepositManagerITransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DepositManagerIFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DepositManagerIFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DepositManagerISession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DepositManagerISession struct {
	Contract     *DepositManagerI  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DepositManagerICallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DepositManagerICallerSession struct {
	Contract *DepositManagerICaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// DepositManagerITransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DepositManagerITransactorSession struct {
	Contract     *DepositManagerITransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// DepositManagerIRaw is an auto generated low-level Go binding around an Ethereum contract.
type DepositManagerIRaw struct {
	Contract *DepositManagerI // Generic contract binding to access the raw methods on
}

// DepositManagerICallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DepositManagerICallerRaw struct {
	Contract *DepositManagerICaller // Generic read-only contract binding to access the raw methods on
}

// DepositManagerITransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DepositManagerITransactorRaw struct {
	Contract *DepositManagerITransactor // Generic write-only contract binding to access the raw methods on
}

// NewDepositManagerI creates a new instance of DepositManagerI, bound to a specific deployed contract.
func NewDepositManagerI(address common.Address, backend bind.ContractBackend) (*DepositManagerI, error) {
	contract, err := bindDepositManagerI(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DepositManagerI{DepositManagerICaller: DepositManagerICaller{contract: contract}, DepositManagerITransactor: DepositManagerITransactor{contract: contract}, DepositManagerIFilterer: DepositManagerIFilterer{contract: contract}}, nil
}

// NewDepositManagerICaller creates a new read-only instance of DepositManagerI, bound to a specific deployed contract.
func NewDepositManagerICaller(address common.Address, caller bind.ContractCaller) (*DepositManagerICaller, error) {
	contract, err := bindDepositManagerI(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DepositManagerICaller{contract: contract}, nil
}

// NewDepositManagerITransactor creates a new write-only instance of DepositManagerI, bound to a specific deployed contract.
func NewDepositManagerITransactor(address common.Address, transactor bind.ContractTransactor) (*DepositManagerITransactor, error) {
	contract, err := bindDepositManagerI(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DepositManagerITransactor{contract: contract}, nil
}

// NewDepositManagerIFilterer creates a new log filterer instance of DepositManagerI, bound to a specific deployed contract.
func NewDepositManagerIFilterer(address common.Address, filterer bind.ContractFilterer) (*DepositManagerIFilterer, error) {
	contract, err := bindDepositManagerI(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DepositManagerIFilterer{contract: contract}, nil
}

// bindDepositManagerI binds a generic wrapper to an already deployed contract.
func bindDepositManagerI(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DepositManagerIABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DepositManagerI *DepositManagerIRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DepositManagerI.Contract.DepositManagerICaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DepositManagerI *DepositManagerIRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DepositManagerI.Contract.DepositManagerITransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DepositManagerI *DepositManagerIRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DepositManagerI.Contract.DepositManagerITransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DepositManagerI *DepositManagerICallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DepositManagerI.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DepositManagerI *DepositManagerITransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DepositManagerI.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DepositManagerI *DepositManagerITransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DepositManagerI.Contract.contract.Transact(opts, method, params...)
}

// WITHDRAWALDELAY is a free data retrieval call binding the contract method 0x0ebb172a.
//
// Solidity: function WITHDRAWAL_DELAY() constant returns(uint256)
func (_DepositManagerI *DepositManagerICaller) WITHDRAWALDELAY(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DepositManagerI.contract.Call(opts, out, "WITHDRAWAL_DELAY")
	return *ret0, err
}

// WITHDRAWALDELAY is a free data retrieval call binding the contract method 0x0ebb172a.
//
// Solidity: function WITHDRAWAL_DELAY() constant returns(uint256)
func (_DepositManagerI *DepositManagerISession) WITHDRAWALDELAY() (*big.Int, error) {
	return _DepositManagerI.Contract.WITHDRAWALDELAY(&_DepositManagerI.CallOpts)
}

// WITHDRAWALDELAY is a free data retrieval call binding the contract method 0x0ebb172a.
//
// Solidity: function WITHDRAWAL_DELAY() constant returns(uint256)
func (_DepositManagerI *DepositManagerICallerSession) WITHDRAWALDELAY() (*big.Int, error) {
	return _DepositManagerI.Contract.WITHDRAWALDELAY(&_DepositManagerI.CallOpts)
}

// AccStaked is a free data retrieval call binding the contract method 0x2d2fab94.
//
// Solidity: function accStaked(address rootchain, address account) constant returns(uint256 wtonAmount)
func (_DepositManagerI *DepositManagerICaller) AccStaked(opts *bind.CallOpts, rootchain common.Address, account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DepositManagerI.contract.Call(opts, out, "accStaked", rootchain, account)
	return *ret0, err
}

// AccStaked is a free data retrieval call binding the contract method 0x2d2fab94.
//
// Solidity: function accStaked(address rootchain, address account) constant returns(uint256 wtonAmount)
func (_DepositManagerI *DepositManagerISession) AccStaked(rootchain common.Address, account common.Address) (*big.Int, error) {
	return _DepositManagerI.Contract.AccStaked(&_DepositManagerI.CallOpts, rootchain, account)
}

// AccStaked is a free data retrieval call binding the contract method 0x2d2fab94.
//
// Solidity: function accStaked(address rootchain, address account) constant returns(uint256 wtonAmount)
func (_DepositManagerI *DepositManagerICallerSession) AccStaked(rootchain common.Address, account common.Address) (*big.Int, error) {
	return _DepositManagerI.Contract.AccStaked(&_DepositManagerI.CallOpts, rootchain, account)
}

// AccStakedAccount is a free data retrieval call binding the contract method 0x0055f5c1.
//
// Solidity: function accStakedAccount(address account) constant returns(uint256 wtonAmount)
func (_DepositManagerI *DepositManagerICaller) AccStakedAccount(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DepositManagerI.contract.Call(opts, out, "accStakedAccount", account)
	return *ret0, err
}

// AccStakedAccount is a free data retrieval call binding the contract method 0x0055f5c1.
//
// Solidity: function accStakedAccount(address account) constant returns(uint256 wtonAmount)
func (_DepositManagerI *DepositManagerISession) AccStakedAccount(account common.Address) (*big.Int, error) {
	return _DepositManagerI.Contract.AccStakedAccount(&_DepositManagerI.CallOpts, account)
}

// AccStakedAccount is a free data retrieval call binding the contract method 0x0055f5c1.
//
// Solidity: function accStakedAccount(address account) constant returns(uint256 wtonAmount)
func (_DepositManagerI *DepositManagerICallerSession) AccStakedAccount(account common.Address) (*big.Int, error) {
	return _DepositManagerI.Contract.AccStakedAccount(&_DepositManagerI.CallOpts, account)
}

// AccStakedRootChain is a free data retrieval call binding the contract method 0xe8035ec9.
//
// Solidity: function accStakedRootChain(address rootchain) constant returns(uint256 wtonAmount)
func (_DepositManagerI *DepositManagerICaller) AccStakedRootChain(opts *bind.CallOpts, rootchain common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DepositManagerI.contract.Call(opts, out, "accStakedRootChain", rootchain)
	return *ret0, err
}

// AccStakedRootChain is a free data retrieval call binding the contract method 0xe8035ec9.
//
// Solidity: function accStakedRootChain(address rootchain) constant returns(uint256 wtonAmount)
func (_DepositManagerI *DepositManagerISession) AccStakedRootChain(rootchain common.Address) (*big.Int, error) {
	return _DepositManagerI.Contract.AccStakedRootChain(&_DepositManagerI.CallOpts, rootchain)
}

// AccStakedRootChain is a free data retrieval call binding the contract method 0xe8035ec9.
//
// Solidity: function accStakedRootChain(address rootchain) constant returns(uint256 wtonAmount)
func (_DepositManagerI *DepositManagerICallerSession) AccStakedRootChain(rootchain common.Address) (*big.Int, error) {
	return _DepositManagerI.Contract.AccStakedRootChain(&_DepositManagerI.CallOpts, rootchain)
}

// AccUnstaked is a free data retrieval call binding the contract method 0x9d91b87b.
//
// Solidity: function accUnstaked(address rootchain, address account) constant returns(uint256 wtonAmount)
func (_DepositManagerI *DepositManagerICaller) AccUnstaked(opts *bind.CallOpts, rootchain common.Address, account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DepositManagerI.contract.Call(opts, out, "accUnstaked", rootchain, account)
	return *ret0, err
}

// AccUnstaked is a free data retrieval call binding the contract method 0x9d91b87b.
//
// Solidity: function accUnstaked(address rootchain, address account) constant returns(uint256 wtonAmount)
func (_DepositManagerI *DepositManagerISession) AccUnstaked(rootchain common.Address, account common.Address) (*big.Int, error) {
	return _DepositManagerI.Contract.AccUnstaked(&_DepositManagerI.CallOpts, rootchain, account)
}

// AccUnstaked is a free data retrieval call binding the contract method 0x9d91b87b.
//
// Solidity: function accUnstaked(address rootchain, address account) constant returns(uint256 wtonAmount)
func (_DepositManagerI *DepositManagerICallerSession) AccUnstaked(rootchain common.Address, account common.Address) (*big.Int, error) {
	return _DepositManagerI.Contract.AccUnstaked(&_DepositManagerI.CallOpts, rootchain, account)
}

// AccUnstakedAccount is a free data retrieval call binding the contract method 0xa3543989.
//
// Solidity: function accUnstakedAccount(address account) constant returns(uint256 wtonAmount)
func (_DepositManagerI *DepositManagerICaller) AccUnstakedAccount(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DepositManagerI.contract.Call(opts, out, "accUnstakedAccount", account)
	return *ret0, err
}

// AccUnstakedAccount is a free data retrieval call binding the contract method 0xa3543989.
//
// Solidity: function accUnstakedAccount(address account) constant returns(uint256 wtonAmount)
func (_DepositManagerI *DepositManagerISession) AccUnstakedAccount(account common.Address) (*big.Int, error) {
	return _DepositManagerI.Contract.AccUnstakedAccount(&_DepositManagerI.CallOpts, account)
}

// AccUnstakedAccount is a free data retrieval call binding the contract method 0xa3543989.
//
// Solidity: function accUnstakedAccount(address account) constant returns(uint256 wtonAmount)
func (_DepositManagerI *DepositManagerICallerSession) AccUnstakedAccount(account common.Address) (*big.Int, error) {
	return _DepositManagerI.Contract.AccUnstakedAccount(&_DepositManagerI.CallOpts, account)
}

// AccUnstakedRootChain is a free data retrieval call binding the contract method 0x03dc6510.
//
// Solidity: function accUnstakedRootChain(address rootchain) constant returns(uint256 wtonAmount)
func (_DepositManagerI *DepositManagerICaller) AccUnstakedRootChain(opts *bind.CallOpts, rootchain common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DepositManagerI.contract.Call(opts, out, "accUnstakedRootChain", rootchain)
	return *ret0, err
}

// AccUnstakedRootChain is a free data retrieval call binding the contract method 0x03dc6510.
//
// Solidity: function accUnstakedRootChain(address rootchain) constant returns(uint256 wtonAmount)
func (_DepositManagerI *DepositManagerISession) AccUnstakedRootChain(rootchain common.Address) (*big.Int, error) {
	return _DepositManagerI.Contract.AccUnstakedRootChain(&_DepositManagerI.CallOpts, rootchain)
}

// AccUnstakedRootChain is a free data retrieval call binding the contract method 0x03dc6510.
//
// Solidity: function accUnstakedRootChain(address rootchain) constant returns(uint256 wtonAmount)
func (_DepositManagerI *DepositManagerICallerSession) AccUnstakedRootChain(rootchain common.Address) (*big.Int, error) {
	return _DepositManagerI.Contract.AccUnstakedRootChain(&_DepositManagerI.CallOpts, rootchain)
}

// NumPendingRequests is a free data retrieval call binding the contract method 0x5c0df46b.
//
// Solidity: function numPendingRequests(address rootchain, address account) constant returns(uint256)
func (_DepositManagerI *DepositManagerICaller) NumPendingRequests(opts *bind.CallOpts, rootchain common.Address, account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DepositManagerI.contract.Call(opts, out, "numPendingRequests", rootchain, account)
	return *ret0, err
}

// NumPendingRequests is a free data retrieval call binding the contract method 0x5c0df46b.
//
// Solidity: function numPendingRequests(address rootchain, address account) constant returns(uint256)
func (_DepositManagerI *DepositManagerISession) NumPendingRequests(rootchain common.Address, account common.Address) (*big.Int, error) {
	return _DepositManagerI.Contract.NumPendingRequests(&_DepositManagerI.CallOpts, rootchain, account)
}

// NumPendingRequests is a free data retrieval call binding the contract method 0x5c0df46b.
//
// Solidity: function numPendingRequests(address rootchain, address account) constant returns(uint256)
func (_DepositManagerI *DepositManagerICallerSession) NumPendingRequests(rootchain common.Address, account common.Address) (*big.Int, error) {
	return _DepositManagerI.Contract.NumPendingRequests(&_DepositManagerI.CallOpts, rootchain, account)
}

// NumRequests is a free data retrieval call binding the contract method 0xf762eb57.
//
// Solidity: function numRequests(address rootchain, address account) constant returns(uint256)
func (_DepositManagerI *DepositManagerICaller) NumRequests(opts *bind.CallOpts, rootchain common.Address, account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DepositManagerI.contract.Call(opts, out, "numRequests", rootchain, account)
	return *ret0, err
}

// NumRequests is a free data retrieval call binding the contract method 0xf762eb57.
//
// Solidity: function numRequests(address rootchain, address account) constant returns(uint256)
func (_DepositManagerI *DepositManagerISession) NumRequests(rootchain common.Address, account common.Address) (*big.Int, error) {
	return _DepositManagerI.Contract.NumRequests(&_DepositManagerI.CallOpts, rootchain, account)
}

// NumRequests is a free data retrieval call binding the contract method 0xf762eb57.
//
// Solidity: function numRequests(address rootchain, address account) constant returns(uint256)
func (_DepositManagerI *DepositManagerICallerSession) NumRequests(rootchain common.Address, account common.Address) (*big.Int, error) {
	return _DepositManagerI.Contract.NumRequests(&_DepositManagerI.CallOpts, rootchain, account)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_DepositManagerI *DepositManagerICaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DepositManagerI.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_DepositManagerI *DepositManagerISession) Owner() (common.Address, error) {
	return _DepositManagerI.Contract.Owner(&_DepositManagerI.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_DepositManagerI *DepositManagerICallerSession) Owner() (common.Address, error) {
	return _DepositManagerI.Contract.Owner(&_DepositManagerI.CallOpts)
}

// PendingUnstaked is a free data retrieval call binding the contract method 0x2638fdf5.
//
// Solidity: function pendingUnstaked(address rootchain, address account) constant returns(uint256 wtonAmount)
func (_DepositManagerI *DepositManagerICaller) PendingUnstaked(opts *bind.CallOpts, rootchain common.Address, account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DepositManagerI.contract.Call(opts, out, "pendingUnstaked", rootchain, account)
	return *ret0, err
}

// PendingUnstaked is a free data retrieval call binding the contract method 0x2638fdf5.
//
// Solidity: function pendingUnstaked(address rootchain, address account) constant returns(uint256 wtonAmount)
func (_DepositManagerI *DepositManagerISession) PendingUnstaked(rootchain common.Address, account common.Address) (*big.Int, error) {
	return _DepositManagerI.Contract.PendingUnstaked(&_DepositManagerI.CallOpts, rootchain, account)
}

// PendingUnstaked is a free data retrieval call binding the contract method 0x2638fdf5.
//
// Solidity: function pendingUnstaked(address rootchain, address account) constant returns(uint256 wtonAmount)
func (_DepositManagerI *DepositManagerICallerSession) PendingUnstaked(rootchain common.Address, account common.Address) (*big.Int, error) {
	return _DepositManagerI.Contract.PendingUnstaked(&_DepositManagerI.CallOpts, rootchain, account)
}

// PendingUnstakedAccount is a free data retrieval call binding the contract method 0xa0b2a913.
//
// Solidity: function pendingUnstakedAccount(address account) constant returns(uint256 wtonAmount)
func (_DepositManagerI *DepositManagerICaller) PendingUnstakedAccount(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DepositManagerI.contract.Call(opts, out, "pendingUnstakedAccount", account)
	return *ret0, err
}

// PendingUnstakedAccount is a free data retrieval call binding the contract method 0xa0b2a913.
//
// Solidity: function pendingUnstakedAccount(address account) constant returns(uint256 wtonAmount)
func (_DepositManagerI *DepositManagerISession) PendingUnstakedAccount(account common.Address) (*big.Int, error) {
	return _DepositManagerI.Contract.PendingUnstakedAccount(&_DepositManagerI.CallOpts, account)
}

// PendingUnstakedAccount is a free data retrieval call binding the contract method 0xa0b2a913.
//
// Solidity: function pendingUnstakedAccount(address account) constant returns(uint256 wtonAmount)
func (_DepositManagerI *DepositManagerICallerSession) PendingUnstakedAccount(account common.Address) (*big.Int, error) {
	return _DepositManagerI.Contract.PendingUnstakedAccount(&_DepositManagerI.CallOpts, account)
}

// PendingUnstakedRootChain is a free data retrieval call binding the contract method 0xa8f3fb98.
//
// Solidity: function pendingUnstakedRootChain(address rootchain) constant returns(uint256 wtonAmount)
func (_DepositManagerI *DepositManagerICaller) PendingUnstakedRootChain(opts *bind.CallOpts, rootchain common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DepositManagerI.contract.Call(opts, out, "pendingUnstakedRootChain", rootchain)
	return *ret0, err
}

// PendingUnstakedRootChain is a free data retrieval call binding the contract method 0xa8f3fb98.
//
// Solidity: function pendingUnstakedRootChain(address rootchain) constant returns(uint256 wtonAmount)
func (_DepositManagerI *DepositManagerISession) PendingUnstakedRootChain(rootchain common.Address) (*big.Int, error) {
	return _DepositManagerI.Contract.PendingUnstakedRootChain(&_DepositManagerI.CallOpts, rootchain)
}

// PendingUnstakedRootChain is a free data retrieval call binding the contract method 0xa8f3fb98.
//
// Solidity: function pendingUnstakedRootChain(address rootchain) constant returns(uint256 wtonAmount)
func (_DepositManagerI *DepositManagerICallerSession) PendingUnstakedRootChain(rootchain common.Address) (*big.Int, error) {
	return _DepositManagerI.Contract.PendingUnstakedRootChain(&_DepositManagerI.CallOpts, rootchain)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() constant returns(address)
func (_DepositManagerI *DepositManagerICaller) Registry(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DepositManagerI.contract.Call(opts, out, "registry")
	return *ret0, err
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() constant returns(address)
func (_DepositManagerI *DepositManagerISession) Registry() (common.Address, error) {
	return _DepositManagerI.Contract.Registry(&_DepositManagerI.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() constant returns(address)
func (_DepositManagerI *DepositManagerICallerSession) Registry() (common.Address, error) {
	return _DepositManagerI.Contract.Registry(&_DepositManagerI.CallOpts)
}

// SeigManager is a free data retrieval call binding the contract method 0x6fb7f558.
//
// Solidity: function seigManager() constant returns(address)
func (_DepositManagerI *DepositManagerICaller) SeigManager(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DepositManagerI.contract.Call(opts, out, "seigManager")
	return *ret0, err
}

// SeigManager is a free data retrieval call binding the contract method 0x6fb7f558.
//
// Solidity: function seigManager() constant returns(address)
func (_DepositManagerI *DepositManagerISession) SeigManager() (common.Address, error) {
	return _DepositManagerI.Contract.SeigManager(&_DepositManagerI.CallOpts)
}

// SeigManager is a free data retrieval call binding the contract method 0x6fb7f558.
//
// Solidity: function seigManager() constant returns(address)
func (_DepositManagerI *DepositManagerICallerSession) SeigManager() (common.Address, error) {
	return _DepositManagerI.Contract.SeigManager(&_DepositManagerI.CallOpts)
}

// WithdrawalRequest is a free data retrieval call binding the contract method 0x8fbef2d0.
//
// Solidity: function withdrawalRequest(address rootchain, address account, uint256 index) constant returns(uint128 withdrawableBlockNumber, uint128 amount, bool processed)
func (_DepositManagerI *DepositManagerICaller) WithdrawalRequest(opts *bind.CallOpts, rootchain common.Address, account common.Address, index *big.Int) (struct {
	WithdrawableBlockNumber *big.Int
	Amount                  *big.Int
	Processed               bool
}, error) {
	ret := new(struct {
		WithdrawableBlockNumber *big.Int
		Amount                  *big.Int
		Processed               bool
	})
	out := ret
	err := _DepositManagerI.contract.Call(opts, out, "withdrawalRequest", rootchain, account, index)
	return *ret, err
}

// WithdrawalRequest is a free data retrieval call binding the contract method 0x8fbef2d0.
//
// Solidity: function withdrawalRequest(address rootchain, address account, uint256 index) constant returns(uint128 withdrawableBlockNumber, uint128 amount, bool processed)
func (_DepositManagerI *DepositManagerISession) WithdrawalRequest(rootchain common.Address, account common.Address, index *big.Int) (struct {
	WithdrawableBlockNumber *big.Int
	Amount                  *big.Int
	Processed               bool
}, error) {
	return _DepositManagerI.Contract.WithdrawalRequest(&_DepositManagerI.CallOpts, rootchain, account, index)
}

// WithdrawalRequest is a free data retrieval call binding the contract method 0x8fbef2d0.
//
// Solidity: function withdrawalRequest(address rootchain, address account, uint256 index) constant returns(uint128 withdrawableBlockNumber, uint128 amount, bool processed)
func (_DepositManagerI *DepositManagerICallerSession) WithdrawalRequest(rootchain common.Address, account common.Address, index *big.Int) (struct {
	WithdrawableBlockNumber *big.Int
	Amount                  *big.Int
	Processed               bool
}, error) {
	return _DepositManagerI.Contract.WithdrawalRequest(&_DepositManagerI.CallOpts, rootchain, account, index)
}

// WithdrawalRequestIndex is a free data retrieval call binding the contract method 0xc647f26e.
//
// Solidity: function withdrawalRequestIndex(address rootchain, address account) constant returns(uint256 index)
func (_DepositManagerI *DepositManagerICaller) WithdrawalRequestIndex(opts *bind.CallOpts, rootchain common.Address, account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DepositManagerI.contract.Call(opts, out, "withdrawalRequestIndex", rootchain, account)
	return *ret0, err
}

// WithdrawalRequestIndex is a free data retrieval call binding the contract method 0xc647f26e.
//
// Solidity: function withdrawalRequestIndex(address rootchain, address account) constant returns(uint256 index)
func (_DepositManagerI *DepositManagerISession) WithdrawalRequestIndex(rootchain common.Address, account common.Address) (*big.Int, error) {
	return _DepositManagerI.Contract.WithdrawalRequestIndex(&_DepositManagerI.CallOpts, rootchain, account)
}

// WithdrawalRequestIndex is a free data retrieval call binding the contract method 0xc647f26e.
//
// Solidity: function withdrawalRequestIndex(address rootchain, address account) constant returns(uint256 index)
func (_DepositManagerI *DepositManagerICallerSession) WithdrawalRequestIndex(rootchain common.Address, account common.Address) (*big.Int, error) {
	return _DepositManagerI.Contract.WithdrawalRequestIndex(&_DepositManagerI.CallOpts, rootchain, account)
}

// Wton is a free data retrieval call binding the contract method 0x8d62d949.
//
// Solidity: function wton() constant returns(address)
func (_DepositManagerI *DepositManagerICaller) Wton(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DepositManagerI.contract.Call(opts, out, "wton")
	return *ret0, err
}

// Wton is a free data retrieval call binding the contract method 0x8d62d949.
//
// Solidity: function wton() constant returns(address)
func (_DepositManagerI *DepositManagerISession) Wton() (common.Address, error) {
	return _DepositManagerI.Contract.Wton(&_DepositManagerI.CallOpts)
}

// Wton is a free data retrieval call binding the contract method 0x8d62d949.
//
// Solidity: function wton() constant returns(address)
func (_DepositManagerI *DepositManagerICallerSession) Wton() (common.Address, error) {
	return _DepositManagerI.Contract.Wton(&_DepositManagerI.CallOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address rootchain, uint256 amount) returns(bool)
func (_DepositManagerI *DepositManagerITransactor) Deposit(opts *bind.TransactOpts, rootchain common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DepositManagerI.contract.Transact(opts, "deposit", rootchain, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address rootchain, uint256 amount) returns(bool)
func (_DepositManagerI *DepositManagerISession) Deposit(rootchain common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DepositManagerI.Contract.Deposit(&_DepositManagerI.TransactOpts, rootchain, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address rootchain, uint256 amount) returns(bool)
func (_DepositManagerI *DepositManagerITransactorSession) Deposit(rootchain common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DepositManagerI.Contract.Deposit(&_DepositManagerI.TransactOpts, rootchain, amount)
}

// ProcessRequest is a paid mutator transaction binding the contract method 0x5d6f7cca.
//
// Solidity: function processRequest(address rootchain) returns(bool)
func (_DepositManagerI *DepositManagerITransactor) ProcessRequest(opts *bind.TransactOpts, rootchain common.Address) (*types.Transaction, error) {
	return _DepositManagerI.contract.Transact(opts, "processRequest", rootchain)
}

// ProcessRequest is a paid mutator transaction binding the contract method 0x5d6f7cca.
//
// Solidity: function processRequest(address rootchain) returns(bool)
func (_DepositManagerI *DepositManagerISession) ProcessRequest(rootchain common.Address) (*types.Transaction, error) {
	return _DepositManagerI.Contract.ProcessRequest(&_DepositManagerI.TransactOpts, rootchain)
}

// ProcessRequest is a paid mutator transaction binding the contract method 0x5d6f7cca.
//
// Solidity: function processRequest(address rootchain) returns(bool)
func (_DepositManagerI *DepositManagerITransactorSession) ProcessRequest(rootchain common.Address) (*types.Transaction, error) {
	return _DepositManagerI.Contract.ProcessRequest(&_DepositManagerI.TransactOpts, rootchain)
}

// ProcessRequests is a paid mutator transaction binding the contract method 0x06260ceb.
//
// Solidity: function processRequests(address rootchain, uint256 n) returns(bool)
func (_DepositManagerI *DepositManagerITransactor) ProcessRequests(opts *bind.TransactOpts, rootchain common.Address, n *big.Int) (*types.Transaction, error) {
	return _DepositManagerI.contract.Transact(opts, "processRequests", rootchain, n)
}

// ProcessRequests is a paid mutator transaction binding the contract method 0x06260ceb.
//
// Solidity: function processRequests(address rootchain, uint256 n) returns(bool)
func (_DepositManagerI *DepositManagerISession) ProcessRequests(rootchain common.Address, n *big.Int) (*types.Transaction, error) {
	return _DepositManagerI.Contract.ProcessRequests(&_DepositManagerI.TransactOpts, rootchain, n)
}

// ProcessRequests is a paid mutator transaction binding the contract method 0x06260ceb.
//
// Solidity: function processRequests(address rootchain, uint256 n) returns(bool)
func (_DepositManagerI *DepositManagerITransactorSession) ProcessRequests(rootchain common.Address, n *big.Int) (*types.Transaction, error) {
	return _DepositManagerI.Contract.ProcessRequests(&_DepositManagerI.TransactOpts, rootchain, n)
}

// RequestWithdrawal is a paid mutator transaction binding the contract method 0xda95ebf7.
//
// Solidity: function requestWithdrawal(address rootchain, uint256 amount) returns(bool)
func (_DepositManagerI *DepositManagerITransactor) RequestWithdrawal(opts *bind.TransactOpts, rootchain common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DepositManagerI.contract.Transact(opts, "requestWithdrawal", rootchain, amount)
}

// RequestWithdrawal is a paid mutator transaction binding the contract method 0xda95ebf7.
//
// Solidity: function requestWithdrawal(address rootchain, uint256 amount) returns(bool)
func (_DepositManagerI *DepositManagerISession) RequestWithdrawal(rootchain common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DepositManagerI.Contract.RequestWithdrawal(&_DepositManagerI.TransactOpts, rootchain, amount)
}

// RequestWithdrawal is a paid mutator transaction binding the contract method 0xda95ebf7.
//
// Solidity: function requestWithdrawal(address rootchain, uint256 amount) returns(bool)
func (_DepositManagerI *DepositManagerITransactorSession) RequestWithdrawal(rootchain common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DepositManagerI.Contract.RequestWithdrawal(&_DepositManagerI.TransactOpts, rootchain, amount)
}

// RequestWithdrawalAll is a paid mutator transaction binding the contract method 0x6b2160b7.
//
// Solidity: function requestWithdrawalAll(address rootchain) returns(bool)
func (_DepositManagerI *DepositManagerITransactor) RequestWithdrawalAll(opts *bind.TransactOpts, rootchain common.Address) (*types.Transaction, error) {
	return _DepositManagerI.contract.Transact(opts, "requestWithdrawalAll", rootchain)
}

// RequestWithdrawalAll is a paid mutator transaction binding the contract method 0x6b2160b7.
//
// Solidity: function requestWithdrawalAll(address rootchain) returns(bool)
func (_DepositManagerI *DepositManagerISession) RequestWithdrawalAll(rootchain common.Address) (*types.Transaction, error) {
	return _DepositManagerI.Contract.RequestWithdrawalAll(&_DepositManagerI.TransactOpts, rootchain)
}

// RequestWithdrawalAll is a paid mutator transaction binding the contract method 0x6b2160b7.
//
// Solidity: function requestWithdrawalAll(address rootchain) returns(bool)
func (_DepositManagerI *DepositManagerITransactorSession) RequestWithdrawalAll(rootchain common.Address) (*types.Transaction, error) {
	return _DepositManagerI.Contract.RequestWithdrawalAll(&_DepositManagerI.TransactOpts, rootchain)
}

// SetSeigManager is a paid mutator transaction binding the contract method 0x7657f20a.
//
// Solidity: function setSeigManager(address seigManager) returns()
func (_DepositManagerI *DepositManagerITransactor) SetSeigManager(opts *bind.TransactOpts, seigManager common.Address) (*types.Transaction, error) {
	return _DepositManagerI.contract.Transact(opts, "setSeigManager", seigManager)
}

// SetSeigManager is a paid mutator transaction binding the contract method 0x7657f20a.
//
// Solidity: function setSeigManager(address seigManager) returns()
func (_DepositManagerI *DepositManagerISession) SetSeigManager(seigManager common.Address) (*types.Transaction, error) {
	return _DepositManagerI.Contract.SetSeigManager(&_DepositManagerI.TransactOpts, seigManager)
}

// SetSeigManager is a paid mutator transaction binding the contract method 0x7657f20a.
//
// Solidity: function setSeigManager(address seigManager) returns()
func (_DepositManagerI *DepositManagerITransactorSession) SetSeigManager(seigManager common.Address) (*types.Transaction, error) {
	return _DepositManagerI.Contract.SetSeigManager(&_DepositManagerI.TransactOpts, seigManager)
}

// ERC165ABI is the input ABI used to generate the binding from.
const ERC165ABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ERC165FuncSigs maps the 4-byte function signature to its string representation.
var ERC165FuncSigs = map[string]string{
	"01ffc9a7": "supportsInterface(bytes4)",
}

// ERC165 is an auto generated Go binding around an Ethereum contract.
type ERC165 struct {
	ERC165Caller     // Read-only binding to the contract
	ERC165Transactor // Write-only binding to the contract
	ERC165Filterer   // Log filterer for contract events
}

// ERC165Caller is an auto generated read-only Go binding around an Ethereum contract.
type ERC165Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC165Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC165Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC165Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC165Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC165Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC165Session struct {
	Contract     *ERC165           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC165CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC165CallerSession struct {
	Contract *ERC165Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ERC165TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC165TransactorSession struct {
	Contract     *ERC165Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC165Raw is an auto generated low-level Go binding around an Ethereum contract.
type ERC165Raw struct {
	Contract *ERC165 // Generic contract binding to access the raw methods on
}

// ERC165CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC165CallerRaw struct {
	Contract *ERC165Caller // Generic read-only contract binding to access the raw methods on
}

// ERC165TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC165TransactorRaw struct {
	Contract *ERC165Transactor // Generic write-only contract binding to access the raw methods on
}

// NewERC165 creates a new instance of ERC165, bound to a specific deployed contract.
func NewERC165(address common.Address, backend bind.ContractBackend) (*ERC165, error) {
	contract, err := bindERC165(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC165{ERC165Caller: ERC165Caller{contract: contract}, ERC165Transactor: ERC165Transactor{contract: contract}, ERC165Filterer: ERC165Filterer{contract: contract}}, nil
}

// NewERC165Caller creates a new read-only instance of ERC165, bound to a specific deployed contract.
func NewERC165Caller(address common.Address, caller bind.ContractCaller) (*ERC165Caller, error) {
	contract, err := bindERC165(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC165Caller{contract: contract}, nil
}

// NewERC165Transactor creates a new write-only instance of ERC165, bound to a specific deployed contract.
func NewERC165Transactor(address common.Address, transactor bind.ContractTransactor) (*ERC165Transactor, error) {
	contract, err := bindERC165(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC165Transactor{contract: contract}, nil
}

// NewERC165Filterer creates a new log filterer instance of ERC165, bound to a specific deployed contract.
func NewERC165Filterer(address common.Address, filterer bind.ContractFilterer) (*ERC165Filterer, error) {
	contract, err := bindERC165(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC165Filterer{contract: contract}, nil
}

// bindERC165 binds a generic wrapper to an already deployed contract.
func bindERC165(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC165ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC165 *ERC165Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC165.Contract.ERC165Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC165 *ERC165Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC165.Contract.ERC165Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC165 *ERC165Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC165.Contract.ERC165Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC165 *ERC165CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC165.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC165 *ERC165TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC165.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC165 *ERC165TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC165.Contract.contract.Transact(opts, method, params...)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) constant returns(bool)
func (_ERC165 *ERC165Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ERC165.contract.Call(opts, out, "supportsInterface", interfaceId)
	return *ret0, err
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) constant returns(bool)
func (_ERC165 *ERC165Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ERC165.Contract.SupportsInterface(&_ERC165.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) constant returns(bool)
func (_ERC165 *ERC165CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ERC165.Contract.SupportsInterface(&_ERC165.CallOpts, interfaceId)
}

// ERC165CheckerABI is the input ABI used to generate the binding from.
const ERC165CheckerABI = "[]"

// ERC165CheckerBin is the compiled bytecode used for deploying new contracts.
var ERC165CheckerBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a72315820ea851d268567a3ed8f55b193eb1a6f3028b7565ea2152c791833f991fbaa44e864736f6c634300050c0032"

// DeployERC165Checker deploys a new Ethereum contract, binding an instance of ERC165Checker to it.
func DeployERC165Checker(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ERC165Checker, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC165CheckerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC165CheckerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC165Checker{ERC165CheckerCaller: ERC165CheckerCaller{contract: contract}, ERC165CheckerTransactor: ERC165CheckerTransactor{contract: contract}, ERC165CheckerFilterer: ERC165CheckerFilterer{contract: contract}}, nil
}

// ERC165Checker is an auto generated Go binding around an Ethereum contract.
type ERC165Checker struct {
	ERC165CheckerCaller     // Read-only binding to the contract
	ERC165CheckerTransactor // Write-only binding to the contract
	ERC165CheckerFilterer   // Log filterer for contract events
}

// ERC165CheckerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC165CheckerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC165CheckerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC165CheckerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC165CheckerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC165CheckerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC165CheckerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC165CheckerSession struct {
	Contract     *ERC165Checker    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC165CheckerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC165CheckerCallerSession struct {
	Contract *ERC165CheckerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// ERC165CheckerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC165CheckerTransactorSession struct {
	Contract     *ERC165CheckerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ERC165CheckerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC165CheckerRaw struct {
	Contract *ERC165Checker // Generic contract binding to access the raw methods on
}

// ERC165CheckerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC165CheckerCallerRaw struct {
	Contract *ERC165CheckerCaller // Generic read-only contract binding to access the raw methods on
}

// ERC165CheckerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC165CheckerTransactorRaw struct {
	Contract *ERC165CheckerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC165Checker creates a new instance of ERC165Checker, bound to a specific deployed contract.
func NewERC165Checker(address common.Address, backend bind.ContractBackend) (*ERC165Checker, error) {
	contract, err := bindERC165Checker(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC165Checker{ERC165CheckerCaller: ERC165CheckerCaller{contract: contract}, ERC165CheckerTransactor: ERC165CheckerTransactor{contract: contract}, ERC165CheckerFilterer: ERC165CheckerFilterer{contract: contract}}, nil
}

// NewERC165CheckerCaller creates a new read-only instance of ERC165Checker, bound to a specific deployed contract.
func NewERC165CheckerCaller(address common.Address, caller bind.ContractCaller) (*ERC165CheckerCaller, error) {
	contract, err := bindERC165Checker(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC165CheckerCaller{contract: contract}, nil
}

// NewERC165CheckerTransactor creates a new write-only instance of ERC165Checker, bound to a specific deployed contract.
func NewERC165CheckerTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC165CheckerTransactor, error) {
	contract, err := bindERC165Checker(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC165CheckerTransactor{contract: contract}, nil
}

// NewERC165CheckerFilterer creates a new log filterer instance of ERC165Checker, bound to a specific deployed contract.
func NewERC165CheckerFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC165CheckerFilterer, error) {
	contract, err := bindERC165Checker(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC165CheckerFilterer{contract: contract}, nil
}

// bindERC165Checker binds a generic wrapper to an already deployed contract.
func bindERC165Checker(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC165CheckerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC165Checker *ERC165CheckerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC165Checker.Contract.ERC165CheckerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC165Checker *ERC165CheckerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC165Checker.Contract.ERC165CheckerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC165Checker *ERC165CheckerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC165Checker.Contract.ERC165CheckerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC165Checker *ERC165CheckerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC165Checker.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC165Checker *ERC165CheckerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC165Checker.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC165Checker *ERC165CheckerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC165Checker.Contract.contract.Transact(opts, method, params...)
}

// ERC20ABI is the input ABI used to generate the binding from.
const ERC20ABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ERC20FuncSigs maps the 4-byte function signature to its string representation.
var ERC20FuncSigs = map[string]string{
	"dd62ed3e": "allowance(address,address)",
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"a457c2d7": "decreaseAllowance(address,uint256)",
	"39509351": "increaseAllowance(address,uint256)",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
}

// ERC20Bin is the compiled bytecode used for deploying new contracts.
var ERC20Bin = "0x608060405261083b806100136000396000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c806370a082311161005b57806370a0823114610149578063a457c2d71461016f578063a9059cbb1461019b578063dd62ed3e146101c757610088565b8063095ea7b31461008d57806318160ddd146100cd57806323b872dd146100e7578063395093511461011d575b600080fd5b6100b9600480360360408110156100a357600080fd5b506001600160a01b0381351690602001356101f5565b604080519115158252519081900360200190f35b6100d5610212565b60408051918252519081900360200190f35b6100b9600480360360608110156100fd57600080fd5b506001600160a01b03813581169160208101359091169060400135610218565b6100b96004803603604081101561013357600080fd5b506001600160a01b0381351690602001356102a5565b6100d56004803603602081101561015f57600080fd5b50356001600160a01b03166102f9565b6100b96004803603604081101561018557600080fd5b506001600160a01b038135169060200135610314565b6100b9600480360360408110156101b157600080fd5b506001600160a01b038135169060200135610382565b6100d5600480360360408110156101dd57600080fd5b506001600160a01b0381358116916020013516610396565b60006102096102026103c1565b84846103c5565b50600192915050565b60025490565b60006102258484846104b1565b61029b846102316103c1565b61029685604051806060016040528060288152602001610771602891396001600160a01b038a1660009081526001602052604081209061026f6103c1565b6001600160a01b03168152602081019190915260400160002054919063ffffffff61060d16565b6103c5565b5060019392505050565b60006102096102b26103c1565b8461029685600160006102c36103c1565b6001600160a01b03908116825260208083019390935260409182016000908120918c16815292529020549063ffffffff6106a416565b6001600160a01b031660009081526020819052604090205490565b60006102096103216103c1565b84610296856040518060600160405280602581526020016107e2602591396001600061034b6103c1565b6001600160a01b03908116825260208083019390935260409182016000908120918d1681529252902054919063ffffffff61060d16565b600061020961038f6103c1565b84846104b1565b6001600160a01b03918216600090815260016020908152604080832093909416825291909152205490565b3390565b6001600160a01b03831661040a5760405162461bcd60e51b81526004018080602001828103825260248152602001806107be6024913960400191505060405180910390fd5b6001600160a01b03821661044f5760405162461bcd60e51b81526004018080602001828103825260228152602001806107296022913960400191505060405180910390fd5b6001600160a01b03808416600081815260016020908152604080832094871680845294825291829020859055815185815291517f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259281900390910190a3505050565b6001600160a01b0383166104f65760405162461bcd60e51b81526004018080602001828103825260258152602001806107996025913960400191505060405180910390fd5b6001600160a01b03821661053b5760405162461bcd60e51b81526004018080602001828103825260238152602001806107066023913960400191505060405180910390fd5b61057e8160405180606001604052806026815260200161074b602691396001600160a01b038616600090815260208190526040902054919063ffffffff61060d16565b6001600160a01b0380851660009081526020819052604080822093909355908416815220546105b3908263ffffffff6106a416565b6001600160a01b038084166000818152602081815260409182902094909455805185815290519193928716927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef92918290030190a3505050565b6000818484111561069c5760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b83811015610661578181015183820152602001610649565b50505050905090810190601f16801561068e5780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b505050900390565b6000828201838110156106fe576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b939250505056fe45524332303a207472616e7366657220746f20746865207a65726f206164647265737345524332303a20617070726f766520746f20746865207a65726f206164647265737345524332303a207472616e7366657220616d6f756e7420657863656564732062616c616e636545524332303a207472616e7366657220616d6f756e74206578636565647320616c6c6f77616e636545524332303a207472616e736665722066726f6d20746865207a65726f206164647265737345524332303a20617070726f76652066726f6d20746865207a65726f206164647265737345524332303a2064656372656173656420616c6c6f77616e63652062656c6f77207a65726fa265627a7a7231582008ca9864b08faad547b35ef43549788be26de128b59e0167fefa4ac97cd6733964736f6c634300050c0032"

// DeployERC20 deploys a new Ethereum contract, binding an instance of ERC20 to it.
func DeployERC20(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ERC20, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC20Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC20{ERC20Caller: ERC20Caller{contract: contract}, ERC20Transactor: ERC20Transactor{contract: contract}, ERC20Filterer: ERC20Filterer{contract: contract}}, nil
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
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_ERC20 *ERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_ERC20 *ERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20.Contract.Allowance(&_ERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_ERC20 *ERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20.Contract.Allowance(&_ERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_ERC20 *ERC20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20.contract.Call(opts, out, "balanceOf", account)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_ERC20 *ERC20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _ERC20.Contract.BalanceOf(&_ERC20.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_ERC20 *ERC20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ERC20.Contract.BalanceOf(&_ERC20.CallOpts, account)
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
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC20 *ERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC20 *ERC20Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Approve(&_ERC20.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC20 *ERC20TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Approve(&_ERC20.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20 *ERC20Transactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20 *ERC20Session) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.DecreaseAllowance(&_ERC20.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20 *ERC20TransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.DecreaseAllowance(&_ERC20.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20 *ERC20Transactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20 *ERC20Session) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.IncreaseAllowance(&_ERC20.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20 *ERC20TransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.IncreaseAllowance(&_ERC20.TransactOpts, spender, addedValue)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20 *ERC20Transactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20 *ERC20Session) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Transfer(&_ERC20.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20 *ERC20TransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Transfer(&_ERC20.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20 *ERC20Transactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20 *ERC20Session) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.TransferFrom(&_ERC20.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20 *ERC20TransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.TransferFrom(&_ERC20.TransactOpts, sender, recipient, amount)
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

// ERC20BurnableABI is the input ABI used to generate the binding from.
const ERC20BurnableABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burnFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ERC20BurnableFuncSigs maps the 4-byte function signature to its string representation.
var ERC20BurnableFuncSigs = map[string]string{
	"dd62ed3e": "allowance(address,address)",
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"42966c68": "burn(uint256)",
	"79cc6790": "burnFrom(address,uint256)",
	"a457c2d7": "decreaseAllowance(address,uint256)",
	"39509351": "increaseAllowance(address,uint256)",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
}

// ERC20BurnableBin is the compiled bytecode used for deploying new contracts.
var ERC20BurnableBin = "0x6080604052610ab7806100136000396000f3fe608060405234801561001057600080fd5b506004361061009e5760003560e01c806370a082311161006657806370a082311461017e57806379cc6790146101a4578063a457c2d7146101d0578063a9059cbb146101fc578063dd62ed3e146102285761009e565b8063095ea7b3146100a357806318160ddd146100e357806323b872dd146100fd578063395093511461013357806342966c681461015f575b600080fd5b6100cf600480360360408110156100b957600080fd5b506001600160a01b038135169060200135610256565b604080519115158252519081900360200190f35b6100eb610273565b60408051918252519081900360200190f35b6100cf6004803603606081101561011357600080fd5b506001600160a01b03813581169160208101359091169060400135610279565b6100cf6004803603604081101561014957600080fd5b506001600160a01b038135169060200135610306565b61017c6004803603602081101561017557600080fd5b503561035a565b005b6100eb6004803603602081101561019457600080fd5b50356001600160a01b031661036e565b61017c600480360360408110156101ba57600080fd5b506001600160a01b038135169060200135610389565b6100cf600480360360408110156101e657600080fd5b506001600160a01b038135169060200135610397565b6100cf6004803603604081101561021257600080fd5b506001600160a01b038135169060200135610405565b6100eb6004803603604081101561023e57600080fd5b506001600160a01b0381358116916020013516610419565b600061026a610263610444565b8484610448565b50600192915050565b60025490565b6000610286848484610534565b6102fc84610292610444565b6102f7856040518060600160405280602881526020016109a8602891396001600160a01b038a166000908152600160205260408120906102d0610444565b6001600160a01b03168152602081019190915260400160002054919063ffffffff61069016565b610448565b5060019392505050565b600061026a610313610444565b846102f78560016000610324610444565b6001600160a01b03908116825260208083019390935260409182016000908120918c16815292529020549063ffffffff61072716565b61036b610365610444565b82610788565b50565b6001600160a01b031660009081526020819052604090205490565b6103938282610884565b5050565b600061026a6103a4610444565b846102f785604051806060016040528060258152602001610a5e60259139600160006103ce610444565b6001600160a01b03908116825260208083019390935260409182016000908120918d1681529252902054919063ffffffff61069016565b600061026a610412610444565b8484610534565b6001600160a01b03918216600090815260016020908152604080832093909416825291909152205490565b3390565b6001600160a01b03831661048d5760405162461bcd60e51b8152600401808060200182810382526024815260200180610a3a6024913960400191505060405180910390fd5b6001600160a01b0382166104d25760405162461bcd60e51b81526004018080602001828103825260228152602001806109606022913960400191505060405180910390fd5b6001600160a01b03808416600081815260016020908152604080832094871680845294825291829020859055815185815291517f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259281900390910190a3505050565b6001600160a01b0383166105795760405162461bcd60e51b8152600401808060200182810382526025815260200180610a156025913960400191505060405180910390fd5b6001600160a01b0382166105be5760405162461bcd60e51b815260040180806020018281038252602381526020018061091b6023913960400191505060405180910390fd5b61060181604051806060016040528060268152602001610982602691396001600160a01b038616600090815260208190526040902054919063ffffffff61069016565b6001600160a01b038085166000908152602081905260408082209390935590841681522054610636908263ffffffff61072716565b6001600160a01b038084166000818152602081815260409182902094909455805185815290519193928716927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef92918290030190a3505050565b6000818484111561071f5760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b838110156106e45781810151838201526020016106cc565b50505050905090810190601f1680156107115780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b505050900390565b600082820183811015610781576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b9392505050565b6001600160a01b0382166107cd5760405162461bcd60e51b81526004018080602001828103825260218152602001806109f46021913960400191505060405180910390fd5b6108108160405180606001604052806022815260200161093e602291396001600160a01b038516600090815260208190526040902054919063ffffffff61069016565b6001600160a01b03831660009081526020819052604090205560025461083c908263ffffffff6108d816565b6002556040805182815290516000916001600160a01b038516917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9181900360200190a35050565b61088e8282610788565b6103938261089a610444565b6102f7846040518060600160405280602481526020016109d0602491396001600160a01b0388166000908152600160205260408120906102d0610444565b600061078183836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f77000081525061069056fe45524332303a207472616e7366657220746f20746865207a65726f206164647265737345524332303a206275726e20616d6f756e7420657863656564732062616c616e636545524332303a20617070726f766520746f20746865207a65726f206164647265737345524332303a207472616e7366657220616d6f756e7420657863656564732062616c616e636545524332303a207472616e7366657220616d6f756e74206578636565647320616c6c6f77616e636545524332303a206275726e20616d6f756e74206578636565647320616c6c6f77616e636545524332303a206275726e2066726f6d20746865207a65726f206164647265737345524332303a207472616e736665722066726f6d20746865207a65726f206164647265737345524332303a20617070726f76652066726f6d20746865207a65726f206164647265737345524332303a2064656372656173656420616c6c6f77616e63652062656c6f77207a65726fa265627a7a72315820f1d69ab24aa4e03a9c57d269487eac94ddd30d01e9f8e7fb8561cbef87d6535a64736f6c634300050c0032"

// DeployERC20Burnable deploys a new Ethereum contract, binding an instance of ERC20Burnable to it.
func DeployERC20Burnable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ERC20Burnable, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20BurnableABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC20BurnableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC20Burnable{ERC20BurnableCaller: ERC20BurnableCaller{contract: contract}, ERC20BurnableTransactor: ERC20BurnableTransactor{contract: contract}, ERC20BurnableFilterer: ERC20BurnableFilterer{contract: contract}}, nil
}

// ERC20Burnable is an auto generated Go binding around an Ethereum contract.
type ERC20Burnable struct {
	ERC20BurnableCaller     // Read-only binding to the contract
	ERC20BurnableTransactor // Write-only binding to the contract
	ERC20BurnableFilterer   // Log filterer for contract events
}

// ERC20BurnableCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20BurnableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20BurnableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20BurnableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20BurnableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20BurnableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20BurnableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20BurnableSession struct {
	Contract     *ERC20Burnable    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20BurnableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20BurnableCallerSession struct {
	Contract *ERC20BurnableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// ERC20BurnableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20BurnableTransactorSession struct {
	Contract     *ERC20BurnableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ERC20BurnableRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20BurnableRaw struct {
	Contract *ERC20Burnable // Generic contract binding to access the raw methods on
}

// ERC20BurnableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20BurnableCallerRaw struct {
	Contract *ERC20BurnableCaller // Generic read-only contract binding to access the raw methods on
}

// ERC20BurnableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20BurnableTransactorRaw struct {
	Contract *ERC20BurnableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20Burnable creates a new instance of ERC20Burnable, bound to a specific deployed contract.
func NewERC20Burnable(address common.Address, backend bind.ContractBackend) (*ERC20Burnable, error) {
	contract, err := bindERC20Burnable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20Burnable{ERC20BurnableCaller: ERC20BurnableCaller{contract: contract}, ERC20BurnableTransactor: ERC20BurnableTransactor{contract: contract}, ERC20BurnableFilterer: ERC20BurnableFilterer{contract: contract}}, nil
}

// NewERC20BurnableCaller creates a new read-only instance of ERC20Burnable, bound to a specific deployed contract.
func NewERC20BurnableCaller(address common.Address, caller bind.ContractCaller) (*ERC20BurnableCaller, error) {
	contract, err := bindERC20Burnable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20BurnableCaller{contract: contract}, nil
}

// NewERC20BurnableTransactor creates a new write-only instance of ERC20Burnable, bound to a specific deployed contract.
func NewERC20BurnableTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC20BurnableTransactor, error) {
	contract, err := bindERC20Burnable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20BurnableTransactor{contract: contract}, nil
}

// NewERC20BurnableFilterer creates a new log filterer instance of ERC20Burnable, bound to a specific deployed contract.
func NewERC20BurnableFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC20BurnableFilterer, error) {
	contract, err := bindERC20Burnable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20BurnableFilterer{contract: contract}, nil
}

// bindERC20Burnable binds a generic wrapper to an already deployed contract.
func bindERC20Burnable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20BurnableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Burnable *ERC20BurnableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20Burnable.Contract.ERC20BurnableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Burnable *ERC20BurnableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Burnable.Contract.ERC20BurnableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Burnable *ERC20BurnableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Burnable.Contract.ERC20BurnableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Burnable *ERC20BurnableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20Burnable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Burnable *ERC20BurnableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Burnable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Burnable *ERC20BurnableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Burnable.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_ERC20Burnable *ERC20BurnableCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20Burnable.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_ERC20Burnable *ERC20BurnableSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20Burnable.Contract.Allowance(&_ERC20Burnable.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_ERC20Burnable *ERC20BurnableCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20Burnable.Contract.Allowance(&_ERC20Burnable.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_ERC20Burnable *ERC20BurnableCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20Burnable.contract.Call(opts, out, "balanceOf", account)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_ERC20Burnable *ERC20BurnableSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ERC20Burnable.Contract.BalanceOf(&_ERC20Burnable.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_ERC20Burnable *ERC20BurnableCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ERC20Burnable.Contract.BalanceOf(&_ERC20Burnable.CallOpts, account)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20Burnable *ERC20BurnableCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20Burnable.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20Burnable *ERC20BurnableSession) TotalSupply() (*big.Int, error) {
	return _ERC20Burnable.Contract.TotalSupply(&_ERC20Burnable.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20Burnable *ERC20BurnableCallerSession) TotalSupply() (*big.Int, error) {
	return _ERC20Burnable.Contract.TotalSupply(&_ERC20Burnable.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC20Burnable *ERC20BurnableTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC20Burnable *ERC20BurnableSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.Contract.Approve(&_ERC20Burnable.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC20Burnable *ERC20BurnableTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.Contract.Approve(&_ERC20Burnable.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_ERC20Burnable *ERC20BurnableTransactor) Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.contract.Transact(opts, "burn", amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_ERC20Burnable *ERC20BurnableSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.Contract.Burn(&_ERC20Burnable.TransactOpts, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_ERC20Burnable *ERC20BurnableTransactorSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.Contract.Burn(&_ERC20Burnable.TransactOpts, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_ERC20Burnable *ERC20BurnableTransactor) BurnFrom(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.contract.Transact(opts, "burnFrom", account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_ERC20Burnable *ERC20BurnableSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.Contract.BurnFrom(&_ERC20Burnable.TransactOpts, account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_ERC20Burnable *ERC20BurnableTransactorSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.Contract.BurnFrom(&_ERC20Burnable.TransactOpts, account, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20Burnable *ERC20BurnableTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20Burnable *ERC20BurnableSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.Contract.DecreaseAllowance(&_ERC20Burnable.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20Burnable *ERC20BurnableTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.Contract.DecreaseAllowance(&_ERC20Burnable.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20Burnable *ERC20BurnableTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20Burnable *ERC20BurnableSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.Contract.IncreaseAllowance(&_ERC20Burnable.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20Burnable *ERC20BurnableTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.Contract.IncreaseAllowance(&_ERC20Burnable.TransactOpts, spender, addedValue)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20Burnable *ERC20BurnableTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20Burnable *ERC20BurnableSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.Contract.Transfer(&_ERC20Burnable.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20Burnable *ERC20BurnableTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.Contract.Transfer(&_ERC20Burnable.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20Burnable *ERC20BurnableTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20Burnable *ERC20BurnableSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.Contract.TransferFrom(&_ERC20Burnable.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20Burnable *ERC20BurnableTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.Contract.TransferFrom(&_ERC20Burnable.TransactOpts, sender, recipient, amount)
}

// ERC20BurnableApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC20Burnable contract.
type ERC20BurnableApprovalIterator struct {
	Event *ERC20BurnableApproval // Event containing the contract specifics and raw log

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
func (it *ERC20BurnableApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20BurnableApproval)
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
		it.Event = new(ERC20BurnableApproval)
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
func (it *ERC20BurnableApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20BurnableApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20BurnableApproval represents a Approval event raised by the ERC20Burnable contract.
type ERC20BurnableApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20Burnable *ERC20BurnableFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ERC20BurnableApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20Burnable.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ERC20BurnableApprovalIterator{contract: _ERC20Burnable.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20Burnable *ERC20BurnableFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC20BurnableApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20Burnable.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20BurnableApproval)
				if err := _ERC20Burnable.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_ERC20Burnable *ERC20BurnableFilterer) ParseApproval(log types.Log) (*ERC20BurnableApproval, error) {
	event := new(ERC20BurnableApproval)
	if err := _ERC20Burnable.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ERC20BurnableTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC20Burnable contract.
type ERC20BurnableTransferIterator struct {
	Event *ERC20BurnableTransfer // Event containing the contract specifics and raw log

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
func (it *ERC20BurnableTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20BurnableTransfer)
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
		it.Event = new(ERC20BurnableTransfer)
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
func (it *ERC20BurnableTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20BurnableTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20BurnableTransfer represents a Transfer event raised by the ERC20Burnable contract.
type ERC20BurnableTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20Burnable *ERC20BurnableFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ERC20BurnableTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20Burnable.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERC20BurnableTransferIterator{contract: _ERC20Burnable.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20Burnable *ERC20BurnableFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC20BurnableTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20Burnable.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20BurnableTransfer)
				if err := _ERC20Burnable.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_ERC20Burnable *ERC20BurnableFilterer) ParseTransfer(log types.Log) (*ERC20BurnableTransfer, error) {
	event := new(ERC20BurnableTransfer)
	if err := _ERC20Burnable.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ERC20DetailedABI is the input ABI used to generate the binding from.
const ERC20DetailedABI = "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ERC20DetailedFuncSigs maps the 4-byte function signature to its string representation.
var ERC20DetailedFuncSigs = map[string]string{
	"dd62ed3e": "allowance(address,address)",
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"313ce567": "decimals()",
	"06fdde03": "name()",
	"95d89b41": "symbol()",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
}

// ERC20Detailed is an auto generated Go binding around an Ethereum contract.
type ERC20Detailed struct {
	ERC20DetailedCaller     // Read-only binding to the contract
	ERC20DetailedTransactor // Write-only binding to the contract
	ERC20DetailedFilterer   // Log filterer for contract events
}

// ERC20DetailedCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20DetailedCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20DetailedTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20DetailedTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20DetailedFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20DetailedFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20DetailedSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20DetailedSession struct {
	Contract     *ERC20Detailed    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20DetailedCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20DetailedCallerSession struct {
	Contract *ERC20DetailedCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// ERC20DetailedTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20DetailedTransactorSession struct {
	Contract     *ERC20DetailedTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ERC20DetailedRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20DetailedRaw struct {
	Contract *ERC20Detailed // Generic contract binding to access the raw methods on
}

// ERC20DetailedCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20DetailedCallerRaw struct {
	Contract *ERC20DetailedCaller // Generic read-only contract binding to access the raw methods on
}

// ERC20DetailedTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20DetailedTransactorRaw struct {
	Contract *ERC20DetailedTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20Detailed creates a new instance of ERC20Detailed, bound to a specific deployed contract.
func NewERC20Detailed(address common.Address, backend bind.ContractBackend) (*ERC20Detailed, error) {
	contract, err := bindERC20Detailed(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20Detailed{ERC20DetailedCaller: ERC20DetailedCaller{contract: contract}, ERC20DetailedTransactor: ERC20DetailedTransactor{contract: contract}, ERC20DetailedFilterer: ERC20DetailedFilterer{contract: contract}}, nil
}

// NewERC20DetailedCaller creates a new read-only instance of ERC20Detailed, bound to a specific deployed contract.
func NewERC20DetailedCaller(address common.Address, caller bind.ContractCaller) (*ERC20DetailedCaller, error) {
	contract, err := bindERC20Detailed(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20DetailedCaller{contract: contract}, nil
}

// NewERC20DetailedTransactor creates a new write-only instance of ERC20Detailed, bound to a specific deployed contract.
func NewERC20DetailedTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC20DetailedTransactor, error) {
	contract, err := bindERC20Detailed(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20DetailedTransactor{contract: contract}, nil
}

// NewERC20DetailedFilterer creates a new log filterer instance of ERC20Detailed, bound to a specific deployed contract.
func NewERC20DetailedFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC20DetailedFilterer, error) {
	contract, err := bindERC20Detailed(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20DetailedFilterer{contract: contract}, nil
}

// bindERC20Detailed binds a generic wrapper to an already deployed contract.
func bindERC20Detailed(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20DetailedABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Detailed *ERC20DetailedRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20Detailed.Contract.ERC20DetailedCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Detailed *ERC20DetailedRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Detailed.Contract.ERC20DetailedTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Detailed *ERC20DetailedRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Detailed.Contract.ERC20DetailedTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Detailed *ERC20DetailedCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20Detailed.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Detailed *ERC20DetailedTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Detailed.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Detailed *ERC20DetailedTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Detailed.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_ERC20Detailed *ERC20DetailedCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20Detailed.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_ERC20Detailed *ERC20DetailedSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20Detailed.Contract.Allowance(&_ERC20Detailed.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_ERC20Detailed *ERC20DetailedCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20Detailed.Contract.Allowance(&_ERC20Detailed.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_ERC20Detailed *ERC20DetailedCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20Detailed.contract.Call(opts, out, "balanceOf", account)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_ERC20Detailed *ERC20DetailedSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ERC20Detailed.Contract.BalanceOf(&_ERC20Detailed.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_ERC20Detailed *ERC20DetailedCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ERC20Detailed.Contract.BalanceOf(&_ERC20Detailed.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_ERC20Detailed *ERC20DetailedCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _ERC20Detailed.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_ERC20Detailed *ERC20DetailedSession) Decimals() (uint8, error) {
	return _ERC20Detailed.Contract.Decimals(&_ERC20Detailed.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_ERC20Detailed *ERC20DetailedCallerSession) Decimals() (uint8, error) {
	return _ERC20Detailed.Contract.Decimals(&_ERC20Detailed.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_ERC20Detailed *ERC20DetailedCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ERC20Detailed.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_ERC20Detailed *ERC20DetailedSession) Name() (string, error) {
	return _ERC20Detailed.Contract.Name(&_ERC20Detailed.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_ERC20Detailed *ERC20DetailedCallerSession) Name() (string, error) {
	return _ERC20Detailed.Contract.Name(&_ERC20Detailed.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_ERC20Detailed *ERC20DetailedCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ERC20Detailed.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_ERC20Detailed *ERC20DetailedSession) Symbol() (string, error) {
	return _ERC20Detailed.Contract.Symbol(&_ERC20Detailed.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_ERC20Detailed *ERC20DetailedCallerSession) Symbol() (string, error) {
	return _ERC20Detailed.Contract.Symbol(&_ERC20Detailed.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20Detailed *ERC20DetailedCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20Detailed.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20Detailed *ERC20DetailedSession) TotalSupply() (*big.Int, error) {
	return _ERC20Detailed.Contract.TotalSupply(&_ERC20Detailed.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20Detailed *ERC20DetailedCallerSession) TotalSupply() (*big.Int, error) {
	return _ERC20Detailed.Contract.TotalSupply(&_ERC20Detailed.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC20Detailed *ERC20DetailedTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Detailed.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC20Detailed *ERC20DetailedSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Detailed.Contract.Approve(&_ERC20Detailed.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC20Detailed *ERC20DetailedTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Detailed.Contract.Approve(&_ERC20Detailed.TransactOpts, spender, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20Detailed *ERC20DetailedTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Detailed.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20Detailed *ERC20DetailedSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Detailed.Contract.Transfer(&_ERC20Detailed.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20Detailed *ERC20DetailedTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Detailed.Contract.Transfer(&_ERC20Detailed.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20Detailed *ERC20DetailedTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Detailed.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20Detailed *ERC20DetailedSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Detailed.Contract.TransferFrom(&_ERC20Detailed.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20Detailed *ERC20DetailedTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Detailed.Contract.TransferFrom(&_ERC20Detailed.TransactOpts, sender, recipient, amount)
}

// ERC20DetailedApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC20Detailed contract.
type ERC20DetailedApprovalIterator struct {
	Event *ERC20DetailedApproval // Event containing the contract specifics and raw log

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
func (it *ERC20DetailedApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20DetailedApproval)
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
		it.Event = new(ERC20DetailedApproval)
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
func (it *ERC20DetailedApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20DetailedApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20DetailedApproval represents a Approval event raised by the ERC20Detailed contract.
type ERC20DetailedApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20Detailed *ERC20DetailedFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ERC20DetailedApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20Detailed.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ERC20DetailedApprovalIterator{contract: _ERC20Detailed.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20Detailed *ERC20DetailedFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC20DetailedApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20Detailed.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20DetailedApproval)
				if err := _ERC20Detailed.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_ERC20Detailed *ERC20DetailedFilterer) ParseApproval(log types.Log) (*ERC20DetailedApproval, error) {
	event := new(ERC20DetailedApproval)
	if err := _ERC20Detailed.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ERC20DetailedTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC20Detailed contract.
type ERC20DetailedTransferIterator struct {
	Event *ERC20DetailedTransfer // Event containing the contract specifics and raw log

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
func (it *ERC20DetailedTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20DetailedTransfer)
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
		it.Event = new(ERC20DetailedTransfer)
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
func (it *ERC20DetailedTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20DetailedTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20DetailedTransfer represents a Transfer event raised by the ERC20Detailed contract.
type ERC20DetailedTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20Detailed *ERC20DetailedFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ERC20DetailedTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20Detailed.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERC20DetailedTransferIterator{contract: _ERC20Detailed.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20Detailed *ERC20DetailedFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC20DetailedTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20Detailed.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20DetailedTransfer)
				if err := _ERC20Detailed.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_ERC20Detailed *ERC20DetailedFilterer) ParseTransfer(log types.Log) (*ERC20DetailedTransfer, error) {
	event := new(ERC20DetailedTransfer)
	if err := _ERC20Detailed.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ERC20MintableABI is the input ABI used to generate the binding from.
const ERC20MintableABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"MinterAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"MinterRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addMinter\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isMinter\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceMinter\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ERC20MintableFuncSigs maps the 4-byte function signature to its string representation.
var ERC20MintableFuncSigs = map[string]string{
	"983b2d56": "addMinter(address)",
	"dd62ed3e": "allowance(address,address)",
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"a457c2d7": "decreaseAllowance(address,uint256)",
	"39509351": "increaseAllowance(address,uint256)",
	"aa271e1a": "isMinter(address)",
	"40c10f19": "mint(address,uint256)",
	"98650275": "renounceMinter()",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
}

// ERC20MintableBin is the compiled bytecode used for deploying new contracts.
var ERC20MintableBin = "0x60806040526100266100186001600160e01b0361002b16565b6001600160e01b0361002f16565b6101a3565b3390565b61004781600361007e60201b610a6e1790919060201c565b6040516001600160a01b038216907f6ae172837ea30b801fbfcdd4108aa1d5bf8ff775444fd70256b44e6bf3dfc3f690600090a250565b61009182826001600160e01b0361012216565b156100fd57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f526f6c65733a206163636f756e7420616c72656164792068617320726f6c6500604482015290519081900360640190fd5b6001600160a01b0316600090815260209190915260409020805460ff19166001179055565b60006001600160a01b038216610183576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526022815260200180610eb16022913960400191505060405180910390fd5b506001600160a01b03166000908152602091909152604090205460ff1690565b610cff806101b26000396000f3fe608060405234801561001057600080fd5b50600436106100b45760003560e01c8063983b2d5611610071578063983b2d56146101c757806398650275146101ef578063a457c2d7146101f7578063a9059cbb14610223578063aa271e1a1461024f578063dd62ed3e14610275576100b4565b8063095ea7b3146100b957806318160ddd146100f957806323b872dd14610113578063395093511461014957806340c10f191461017557806370a08231146101a1575b600080fd5b6100e5600480360360408110156100cf57600080fd5b506001600160a01b0381351690602001356102a3565b604080519115158252519081900360200190f35b6101016102c0565b60408051918252519081900360200190f35b6100e56004803603606081101561012957600080fd5b506001600160a01b038135811691602081013590911690604001356102c6565b6100e56004803603604081101561015f57600080fd5b506001600160a01b038135169060200135610353565b6100e56004803603604081101561018b57600080fd5b506001600160a01b0381351690602001356103a7565b610101600480360360208110156101b757600080fd5b50356001600160a01b03166103fe565b6101ed600480360360208110156101dd57600080fd5b50356001600160a01b0316610419565b005b6101ed61046b565b6100e56004803603604081101561020d57600080fd5b506001600160a01b03813516906020013561047d565b6100e56004803603604081101561023957600080fd5b506001600160a01b0381351690602001356104eb565b6100e56004803603602081101561026557600080fd5b50356001600160a01b03166104ff565b6101016004803603604081101561028b57600080fd5b506001600160a01b0381358116916020013516610518565b60006102b76102b0610543565b8484610547565b50600192915050565b60025490565b60006102d3848484610633565b610349846102df610543565b61034485604051806060016040528060288152602001610c13602891396001600160a01b038a1660009081526001602052604081209061031d610543565b6001600160a01b03168152602081019190915260400160002054919063ffffffff61078f16565b610547565b5060019392505050565b60006102b7610360610543565b846103448560016000610371610543565b6001600160a01b03908116825260208083019390935260409182016000908120918c16815292529020549063ffffffff61082616565b60006103b96103b4610543565b6104ff565b6103f45760405162461bcd60e51b8152600401808060200182810382526030815260200180610bc26030913960400191505060405180910390fd5b6102b78383610887565b6001600160a01b031660009081526020819052604090205490565b6104246103b4610543565b61045f5760405162461bcd60e51b8152600401808060200182810382526030815260200180610bc26030913960400191505060405180910390fd5b61046881610977565b50565b61047b610476610543565b6109bf565b565b60006102b761048a610543565b8461034485604051806060016040528060258152602001610ca660259139600160006104b4610543565b6001600160a01b03908116825260208083019390935260409182016000908120918d1681529252902054919063ffffffff61078f16565b60006102b76104f8610543565b8484610633565b600061051260038363ffffffff610a0716565b92915050565b6001600160a01b03918216600090815260016020908152604080832093909416825291909152205490565b3390565b6001600160a01b03831661058c5760405162461bcd60e51b8152600401808060200182810382526024815260200180610c826024913960400191505060405180910390fd5b6001600160a01b0382166105d15760405162461bcd60e51b8152600401808060200182810382526022815260200180610b7a6022913960400191505060405180910390fd5b6001600160a01b03808416600081815260016020908152604080832094871680845294825291829020859055815185815291517f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259281900390910190a3505050565b6001600160a01b0383166106785760405162461bcd60e51b8152600401808060200182810382526025815260200180610c5d6025913960400191505060405180910390fd5b6001600160a01b0382166106bd5760405162461bcd60e51b8152600401808060200182810382526023815260200180610b576023913960400191505060405180910390fd5b61070081604051806060016040528060268152602001610b9c602691396001600160a01b038616600090815260208190526040902054919063ffffffff61078f16565b6001600160a01b038085166000908152602081905260408082209390935590841681522054610735908263ffffffff61082616565b6001600160a01b038084166000818152602081815260409182902094909455805185815290519193928716927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef92918290030190a3505050565b6000818484111561081e5760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b838110156107e35781810151838201526020016107cb565b50505050905090810190601f1680156108105780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b505050900390565b600082820183811015610880576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b9392505050565b6001600160a01b0382166108e2576040805162461bcd60e51b815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f206164647265737300604482015290519081900360640190fd5b6002546108f5908263ffffffff61082616565b6002556001600160a01b038216600090815260208190526040902054610921908263ffffffff61082616565b6001600160a01b0383166000818152602081815260408083209490945583518581529351929391927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9281900390910190a35050565b61098860038263ffffffff610a6e16565b6040516001600160a01b038216907f6ae172837ea30b801fbfcdd4108aa1d5bf8ff775444fd70256b44e6bf3dfc3f690600090a250565b6109d060038263ffffffff610aef16565b6040516001600160a01b038216907fe94479a9f7e1952cc78f2d6baab678adc1b772d936c6583def489e524cb6669290600090a250565b60006001600160a01b038216610a4e5760405162461bcd60e51b8152600401808060200182810382526022815260200180610c3b6022913960400191505060405180910390fd5b506001600160a01b03166000908152602091909152604090205460ff1690565b610a788282610a07565b15610aca576040805162461bcd60e51b815260206004820152601f60248201527f526f6c65733a206163636f756e7420616c72656164792068617320726f6c6500604482015290519081900360640190fd5b6001600160a01b0316600090815260209190915260409020805460ff19166001179055565b610af98282610a07565b610b345760405162461bcd60e51b8152600401808060200182810382526021815260200180610bf26021913960400191505060405180910390fd5b6001600160a01b0316600090815260209190915260409020805460ff1916905556fe45524332303a207472616e7366657220746f20746865207a65726f206164647265737345524332303a20617070726f766520746f20746865207a65726f206164647265737345524332303a207472616e7366657220616d6f756e7420657863656564732062616c616e63654d696e746572526f6c653a2063616c6c657220646f6573206e6f74206861766520746865204d696e74657220726f6c65526f6c65733a206163636f756e7420646f6573206e6f74206861766520726f6c6545524332303a207472616e7366657220616d6f756e74206578636565647320616c6c6f77616e6365526f6c65733a206163636f756e7420697320746865207a65726f206164647265737345524332303a207472616e736665722066726f6d20746865207a65726f206164647265737345524332303a20617070726f76652066726f6d20746865207a65726f206164647265737345524332303a2064656372656173656420616c6c6f77616e63652062656c6f77207a65726fa265627a7a723158209bffbae26a8473ab446e60484fc8f37931f262c4c630ff4ea7ef2fd99a2a4ba364736f6c634300050c0032526f6c65733a206163636f756e7420697320746865207a65726f2061646472657373"

// DeployERC20Mintable deploys a new Ethereum contract, binding an instance of ERC20Mintable to it.
func DeployERC20Mintable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ERC20Mintable, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20MintableABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC20MintableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC20Mintable{ERC20MintableCaller: ERC20MintableCaller{contract: contract}, ERC20MintableTransactor: ERC20MintableTransactor{contract: contract}, ERC20MintableFilterer: ERC20MintableFilterer{contract: contract}}, nil
}

// ERC20Mintable is an auto generated Go binding around an Ethereum contract.
type ERC20Mintable struct {
	ERC20MintableCaller     // Read-only binding to the contract
	ERC20MintableTransactor // Write-only binding to the contract
	ERC20MintableFilterer   // Log filterer for contract events
}

// ERC20MintableCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20MintableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20MintableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20MintableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20MintableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20MintableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20MintableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20MintableSession struct {
	Contract     *ERC20Mintable    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20MintableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20MintableCallerSession struct {
	Contract *ERC20MintableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// ERC20MintableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20MintableTransactorSession struct {
	Contract     *ERC20MintableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ERC20MintableRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20MintableRaw struct {
	Contract *ERC20Mintable // Generic contract binding to access the raw methods on
}

// ERC20MintableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20MintableCallerRaw struct {
	Contract *ERC20MintableCaller // Generic read-only contract binding to access the raw methods on
}

// ERC20MintableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20MintableTransactorRaw struct {
	Contract *ERC20MintableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20Mintable creates a new instance of ERC20Mintable, bound to a specific deployed contract.
func NewERC20Mintable(address common.Address, backend bind.ContractBackend) (*ERC20Mintable, error) {
	contract, err := bindERC20Mintable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20Mintable{ERC20MintableCaller: ERC20MintableCaller{contract: contract}, ERC20MintableTransactor: ERC20MintableTransactor{contract: contract}, ERC20MintableFilterer: ERC20MintableFilterer{contract: contract}}, nil
}

// NewERC20MintableCaller creates a new read-only instance of ERC20Mintable, bound to a specific deployed contract.
func NewERC20MintableCaller(address common.Address, caller bind.ContractCaller) (*ERC20MintableCaller, error) {
	contract, err := bindERC20Mintable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20MintableCaller{contract: contract}, nil
}

// NewERC20MintableTransactor creates a new write-only instance of ERC20Mintable, bound to a specific deployed contract.
func NewERC20MintableTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC20MintableTransactor, error) {
	contract, err := bindERC20Mintable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20MintableTransactor{contract: contract}, nil
}

// NewERC20MintableFilterer creates a new log filterer instance of ERC20Mintable, bound to a specific deployed contract.
func NewERC20MintableFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC20MintableFilterer, error) {
	contract, err := bindERC20Mintable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20MintableFilterer{contract: contract}, nil
}

// bindERC20Mintable binds a generic wrapper to an already deployed contract.
func bindERC20Mintable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20MintableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Mintable *ERC20MintableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20Mintable.Contract.ERC20MintableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Mintable *ERC20MintableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Mintable.Contract.ERC20MintableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Mintable *ERC20MintableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Mintable.Contract.ERC20MintableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Mintable *ERC20MintableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20Mintable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Mintable *ERC20MintableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Mintable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Mintable *ERC20MintableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Mintable.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_ERC20Mintable *ERC20MintableCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20Mintable.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_ERC20Mintable *ERC20MintableSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20Mintable.Contract.Allowance(&_ERC20Mintable.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_ERC20Mintable *ERC20MintableCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20Mintable.Contract.Allowance(&_ERC20Mintable.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_ERC20Mintable *ERC20MintableCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20Mintable.contract.Call(opts, out, "balanceOf", account)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_ERC20Mintable *ERC20MintableSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ERC20Mintable.Contract.BalanceOf(&_ERC20Mintable.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_ERC20Mintable *ERC20MintableCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ERC20Mintable.Contract.BalanceOf(&_ERC20Mintable.CallOpts, account)
}

// IsMinter is a free data retrieval call binding the contract method 0xaa271e1a.
//
// Solidity: function isMinter(address account) constant returns(bool)
func (_ERC20Mintable *ERC20MintableCaller) IsMinter(opts *bind.CallOpts, account common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ERC20Mintable.contract.Call(opts, out, "isMinter", account)
	return *ret0, err
}

// IsMinter is a free data retrieval call binding the contract method 0xaa271e1a.
//
// Solidity: function isMinter(address account) constant returns(bool)
func (_ERC20Mintable *ERC20MintableSession) IsMinter(account common.Address) (bool, error) {
	return _ERC20Mintable.Contract.IsMinter(&_ERC20Mintable.CallOpts, account)
}

// IsMinter is a free data retrieval call binding the contract method 0xaa271e1a.
//
// Solidity: function isMinter(address account) constant returns(bool)
func (_ERC20Mintable *ERC20MintableCallerSession) IsMinter(account common.Address) (bool, error) {
	return _ERC20Mintable.Contract.IsMinter(&_ERC20Mintable.CallOpts, account)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20Mintable *ERC20MintableCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20Mintable.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20Mintable *ERC20MintableSession) TotalSupply() (*big.Int, error) {
	return _ERC20Mintable.Contract.TotalSupply(&_ERC20Mintable.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20Mintable *ERC20MintableCallerSession) TotalSupply() (*big.Int, error) {
	return _ERC20Mintable.Contract.TotalSupply(&_ERC20Mintable.CallOpts)
}

// AddMinter is a paid mutator transaction binding the contract method 0x983b2d56.
//
// Solidity: function addMinter(address account) returns()
func (_ERC20Mintable *ERC20MintableTransactor) AddMinter(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _ERC20Mintable.contract.Transact(opts, "addMinter", account)
}

// AddMinter is a paid mutator transaction binding the contract method 0x983b2d56.
//
// Solidity: function addMinter(address account) returns()
func (_ERC20Mintable *ERC20MintableSession) AddMinter(account common.Address) (*types.Transaction, error) {
	return _ERC20Mintable.Contract.AddMinter(&_ERC20Mintable.TransactOpts, account)
}

// AddMinter is a paid mutator transaction binding the contract method 0x983b2d56.
//
// Solidity: function addMinter(address account) returns()
func (_ERC20Mintable *ERC20MintableTransactorSession) AddMinter(account common.Address) (*types.Transaction, error) {
	return _ERC20Mintable.Contract.AddMinter(&_ERC20Mintable.TransactOpts, account)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC20Mintable *ERC20MintableTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Mintable.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC20Mintable *ERC20MintableSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Mintable.Contract.Approve(&_ERC20Mintable.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC20Mintable *ERC20MintableTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Mintable.Contract.Approve(&_ERC20Mintable.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20Mintable *ERC20MintableTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20Mintable.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20Mintable *ERC20MintableSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20Mintable.Contract.DecreaseAllowance(&_ERC20Mintable.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20Mintable *ERC20MintableTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20Mintable.Contract.DecreaseAllowance(&_ERC20Mintable.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20Mintable *ERC20MintableTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20Mintable.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20Mintable *ERC20MintableSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20Mintable.Contract.IncreaseAllowance(&_ERC20Mintable.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20Mintable *ERC20MintableTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20Mintable.Contract.IncreaseAllowance(&_ERC20Mintable.TransactOpts, spender, addedValue)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns(bool)
func (_ERC20Mintable *ERC20MintableTransactor) Mint(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Mintable.contract.Transact(opts, "mint", account, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns(bool)
func (_ERC20Mintable *ERC20MintableSession) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Mintable.Contract.Mint(&_ERC20Mintable.TransactOpts, account, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns(bool)
func (_ERC20Mintable *ERC20MintableTransactorSession) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Mintable.Contract.Mint(&_ERC20Mintable.TransactOpts, account, amount)
}

// RenounceMinter is a paid mutator transaction binding the contract method 0x98650275.
//
// Solidity: function renounceMinter() returns()
func (_ERC20Mintable *ERC20MintableTransactor) RenounceMinter(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Mintable.contract.Transact(opts, "renounceMinter")
}

// RenounceMinter is a paid mutator transaction binding the contract method 0x98650275.
//
// Solidity: function renounceMinter() returns()
func (_ERC20Mintable *ERC20MintableSession) RenounceMinter() (*types.Transaction, error) {
	return _ERC20Mintable.Contract.RenounceMinter(&_ERC20Mintable.TransactOpts)
}

// RenounceMinter is a paid mutator transaction binding the contract method 0x98650275.
//
// Solidity: function renounceMinter() returns()
func (_ERC20Mintable *ERC20MintableTransactorSession) RenounceMinter() (*types.Transaction, error) {
	return _ERC20Mintable.Contract.RenounceMinter(&_ERC20Mintable.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20Mintable *ERC20MintableTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Mintable.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20Mintable *ERC20MintableSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Mintable.Contract.Transfer(&_ERC20Mintable.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20Mintable *ERC20MintableTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Mintable.Contract.Transfer(&_ERC20Mintable.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20Mintable *ERC20MintableTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Mintable.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20Mintable *ERC20MintableSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Mintable.Contract.TransferFrom(&_ERC20Mintable.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20Mintable *ERC20MintableTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Mintable.Contract.TransferFrom(&_ERC20Mintable.TransactOpts, sender, recipient, amount)
}

// ERC20MintableApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC20Mintable contract.
type ERC20MintableApprovalIterator struct {
	Event *ERC20MintableApproval // Event containing the contract specifics and raw log

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
func (it *ERC20MintableApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20MintableApproval)
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
		it.Event = new(ERC20MintableApproval)
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
func (it *ERC20MintableApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20MintableApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20MintableApproval represents a Approval event raised by the ERC20Mintable contract.
type ERC20MintableApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20Mintable *ERC20MintableFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ERC20MintableApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20Mintable.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ERC20MintableApprovalIterator{contract: _ERC20Mintable.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20Mintable *ERC20MintableFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC20MintableApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20Mintable.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20MintableApproval)
				if err := _ERC20Mintable.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_ERC20Mintable *ERC20MintableFilterer) ParseApproval(log types.Log) (*ERC20MintableApproval, error) {
	event := new(ERC20MintableApproval)
	if err := _ERC20Mintable.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ERC20MintableMinterAddedIterator is returned from FilterMinterAdded and is used to iterate over the raw logs and unpacked data for MinterAdded events raised by the ERC20Mintable contract.
type ERC20MintableMinterAddedIterator struct {
	Event *ERC20MintableMinterAdded // Event containing the contract specifics and raw log

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
func (it *ERC20MintableMinterAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20MintableMinterAdded)
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
		it.Event = new(ERC20MintableMinterAdded)
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
func (it *ERC20MintableMinterAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20MintableMinterAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20MintableMinterAdded represents a MinterAdded event raised by the ERC20Mintable contract.
type ERC20MintableMinterAdded struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMinterAdded is a free log retrieval operation binding the contract event 0x6ae172837ea30b801fbfcdd4108aa1d5bf8ff775444fd70256b44e6bf3dfc3f6.
//
// Solidity: event MinterAdded(address indexed account)
func (_ERC20Mintable *ERC20MintableFilterer) FilterMinterAdded(opts *bind.FilterOpts, account []common.Address) (*ERC20MintableMinterAddedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ERC20Mintable.contract.FilterLogs(opts, "MinterAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return &ERC20MintableMinterAddedIterator{contract: _ERC20Mintable.contract, event: "MinterAdded", logs: logs, sub: sub}, nil
}

// WatchMinterAdded is a free log subscription operation binding the contract event 0x6ae172837ea30b801fbfcdd4108aa1d5bf8ff775444fd70256b44e6bf3dfc3f6.
//
// Solidity: event MinterAdded(address indexed account)
func (_ERC20Mintable *ERC20MintableFilterer) WatchMinterAdded(opts *bind.WatchOpts, sink chan<- *ERC20MintableMinterAdded, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ERC20Mintable.contract.WatchLogs(opts, "MinterAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20MintableMinterAdded)
				if err := _ERC20Mintable.contract.UnpackLog(event, "MinterAdded", log); err != nil {
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

// ParseMinterAdded is a log parse operation binding the contract event 0x6ae172837ea30b801fbfcdd4108aa1d5bf8ff775444fd70256b44e6bf3dfc3f6.
//
// Solidity: event MinterAdded(address indexed account)
func (_ERC20Mintable *ERC20MintableFilterer) ParseMinterAdded(log types.Log) (*ERC20MintableMinterAdded, error) {
	event := new(ERC20MintableMinterAdded)
	if err := _ERC20Mintable.contract.UnpackLog(event, "MinterAdded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ERC20MintableMinterRemovedIterator is returned from FilterMinterRemoved and is used to iterate over the raw logs and unpacked data for MinterRemoved events raised by the ERC20Mintable contract.
type ERC20MintableMinterRemovedIterator struct {
	Event *ERC20MintableMinterRemoved // Event containing the contract specifics and raw log

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
func (it *ERC20MintableMinterRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20MintableMinterRemoved)
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
		it.Event = new(ERC20MintableMinterRemoved)
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
func (it *ERC20MintableMinterRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20MintableMinterRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20MintableMinterRemoved represents a MinterRemoved event raised by the ERC20Mintable contract.
type ERC20MintableMinterRemoved struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMinterRemoved is a free log retrieval operation binding the contract event 0xe94479a9f7e1952cc78f2d6baab678adc1b772d936c6583def489e524cb66692.
//
// Solidity: event MinterRemoved(address indexed account)
func (_ERC20Mintable *ERC20MintableFilterer) FilterMinterRemoved(opts *bind.FilterOpts, account []common.Address) (*ERC20MintableMinterRemovedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ERC20Mintable.contract.FilterLogs(opts, "MinterRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return &ERC20MintableMinterRemovedIterator{contract: _ERC20Mintable.contract, event: "MinterRemoved", logs: logs, sub: sub}, nil
}

// WatchMinterRemoved is a free log subscription operation binding the contract event 0xe94479a9f7e1952cc78f2d6baab678adc1b772d936c6583def489e524cb66692.
//
// Solidity: event MinterRemoved(address indexed account)
func (_ERC20Mintable *ERC20MintableFilterer) WatchMinterRemoved(opts *bind.WatchOpts, sink chan<- *ERC20MintableMinterRemoved, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ERC20Mintable.contract.WatchLogs(opts, "MinterRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20MintableMinterRemoved)
				if err := _ERC20Mintable.contract.UnpackLog(event, "MinterRemoved", log); err != nil {
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

// ParseMinterRemoved is a log parse operation binding the contract event 0xe94479a9f7e1952cc78f2d6baab678adc1b772d936c6583def489e524cb66692.
//
// Solidity: event MinterRemoved(address indexed account)
func (_ERC20Mintable *ERC20MintableFilterer) ParseMinterRemoved(log types.Log) (*ERC20MintableMinterRemoved, error) {
	event := new(ERC20MintableMinterRemoved)
	if err := _ERC20Mintable.contract.UnpackLog(event, "MinterRemoved", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ERC20MintableTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC20Mintable contract.
type ERC20MintableTransferIterator struct {
	Event *ERC20MintableTransfer // Event containing the contract specifics and raw log

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
func (it *ERC20MintableTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20MintableTransfer)
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
		it.Event = new(ERC20MintableTransfer)
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
func (it *ERC20MintableTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20MintableTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20MintableTransfer represents a Transfer event raised by the ERC20Mintable contract.
type ERC20MintableTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20Mintable *ERC20MintableFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ERC20MintableTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20Mintable.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERC20MintableTransferIterator{contract: _ERC20Mintable.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20Mintable *ERC20MintableFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC20MintableTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20Mintable.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20MintableTransfer)
				if err := _ERC20Mintable.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_ERC20Mintable *ERC20MintableFilterer) ParseTransfer(log types.Log) (*ERC20MintableTransfer, error) {
	event := new(ERC20MintableTransfer)
	if err := _ERC20Mintable.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ERC20OnApproveABI is the input ABI used to generate the binding from.
const ERC20OnApproveABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"INTERFACE_ID_ON_APPROVE\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"approveAndCall\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ERC20OnApproveFuncSigs maps the 4-byte function signature to its string representation.
var ERC20OnApproveFuncSigs = map[string]string{
	"6cd28f9a": "INTERFACE_ID_ON_APPROVE()",
	"dd62ed3e": "allowance(address,address)",
	"095ea7b3": "approve(address,uint256)",
	"cae9ca51": "approveAndCall(address,uint256,bytes)",
	"70a08231": "balanceOf(address)",
	"a457c2d7": "decreaseAllowance(address,uint256)",
	"39509351": "increaseAllowance(address,uint256)",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
}

// ERC20OnApproveBin is the compiled bytecode used for deploying new contracts.
var ERC20OnApproveBin = "0x6080604052610dc0806100136000396000f3fe608060405234801561001057600080fd5b506004361061009e5760003560e01c806370a082311161006657806370a0823114610184578063a457c2d7146101aa578063a9059cbb146101d6578063cae9ca5114610202578063dd62ed3e146102bd5761009e565b8063095ea7b3146100a357806318160ddd146100e357806323b872dd146100fd57806339509351146101335780636cd28f9a1461015f575b600080fd5b6100cf600480360360408110156100b957600080fd5b506001600160a01b0381351690602001356102eb565b604080519115158252519081900360200190f35b6100eb610308565b60408051918252519081900360200190f35b6100cf6004803603606081101561011357600080fd5b506001600160a01b0381358116916020810135909116906040013561030e565b6100cf6004803603604081101561014957600080fd5b506001600160a01b03813516906020013561039c565b6101676103f0565b604080516001600160e01b03199092168252519081900360200190f35b6100eb6004803603602081101561019a57600080fd5b50356001600160a01b031661040b565b6100cf600480360360408110156101c057600080fd5b506001600160a01b038135169060200135610426565b6100cf600480360360408110156101ec57600080fd5b506001600160a01b038135169060200135610494565b6100cf6004803603606081101561021857600080fd5b6001600160a01b038235169160208101359181019060608101604082013564010000000081111561024857600080fd5b82018360208201111561025a57600080fd5b8035906020019184600183028401116401000000008311171561027c57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506104a8945050505050565b6100eb600480360360408110156102d357600080fd5b506001600160a01b03813581169160200135166104c9565b60006102ff6102f86104f4565b84846104f8565b50600192915050565b60025490565b600061031b8484846105e4565b610391846103276104f4565b61038c85604051806060016040528060288152602001610cf6602891396001600160a01b038a166000908152600160205260408120906103656104f4565b6001600160a01b03168152602081019190915260400160002054919063ffffffff61074016565b6104f8565b5060015b9392505050565b60006102ff6103a96104f4565b8461038c85600160006103ba6104f4565b6001600160a01b03908116825260208083019390935260409182016000908120918c16815292529020549063ffffffff6107d716565b604051806028610cce82396028019050604051809103902081565b6001600160a01b031660009081526020819052604090205490565b60006102ff6104336104f4565b8461038c85604051806060016040528060258152602001610d67602591396001600061045d6104f4565b6001600160a01b03908116825260208083019390935260409182016000908120918d1681529252902054919063ffffffff61074016565b60006102ff6104a16104f4565b84846105e4565b60006104b484846102eb565b6104bd57600080fd5b61039533858585610831565b6001600160a01b03918216600090815260016020908152604080832093909416825291909152205490565b3390565b6001600160a01b03831661053d5760405162461bcd60e51b8152600401808060200182810382526024815260200180610d436024913960400191505060405180910390fd5b6001600160a01b0382166105825760405162461bcd60e51b8152600401808060200182810382526022815260200180610c556022913960400191505060405180910390fd5b6001600160a01b03808416600081815260016020908152604080832094871680845294825291829020859055815185815291517f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259281900390910190a3505050565b6001600160a01b0383166106295760405162461bcd60e51b8152600401808060200182810382526025815260200180610d1e6025913960400191505060405180910390fd5b6001600160a01b03821661066e5760405162461bcd60e51b8152600401808060200182810382526023815260200180610c326023913960400191505060405180910390fd5b6106b181604051806060016040528060268152602001610ca8602691396001600160a01b038616600090815260208190526040902054919063ffffffff61074016565b6001600160a01b0380851660009081526020819052604080822093909355908416815220546106e6908263ffffffff6107d716565b6001600160a01b038084166000818152602081815260409182902094909455805185815290519193928716927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef92918290030190a3505050565b600081848411156107cf5760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b8381101561079457818101518382015260200161077c565b50505050905090810190601f1680156107c15780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b505050900390565b600082820183811015610395576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b610853836040518080610cce6028913960280190506040518091039020610a87565b61088e5760405162461bcd60e51b8152600401808060200182810382526031815260200180610c776031913960400191505060405180910390fd5b60006060846001600160a01b03166040518080610cce60289139602801905060405180910390208787878760405160240180856001600160a01b03166001600160a01b03168152602001846001600160a01b03166001600160a01b0316815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b8381101561092f578181015183820152602001610917565b50505050905090810190601f16801561095c5780820380516001836020036101000a031916815260200191505b5060408051601f198184030181529181526020820180516001600160e01b03166001600160e01b0319909a16999099178952518151919890975087965094509250829150849050835b602083106109c45780518252601f1990920191602091820191016109a5565b6001836020036101000a0380198251168184511680821785525050505050509050019150506000604051808303816000865af19150503d8060008114610a26576040519150601f19603f3d011682016040523d82523d6000602084013e610a2b565b606091505b5091509150818190610a7e5760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561079457818101518382015260200161077c565b50505050505050565b6000610a9283610aa3565b801561039557506103958383610ad7565b6000610ab6826301ffc9a760e01b610ad7565b8015610ad15750610acf826001600160e01b0319610ad7565b155b92915050565b6000806000610ae68585610afd565b91509150818015610af45750805b95945050505050565b604080516001600160e01b0319831660248083019190915282518083039091018152604490910182526020810180516001600160e01b03166301ffc9a760e01b1781529151815160009384939284926060926001600160a01b038a169261753092879282918083835b60208310610b855780518252601f199092019160209182019101610b66565b6001836020036101000a0380198251168184511680821785525050505050509050019150506000604051808303818686fa925050503d8060008114610be6576040519150601f19603f3d011682016040523d82523d6000602084013e610beb565b606091505b5091509150602081511015610c095760008094509450505050610c2a565b81818060200190516020811015610c1f57600080fd5b505190955093505050505b925092905056fe45524332303a207472616e7366657220746f20746865207a65726f206164647265737345524332303a20617070726f766520746f20746865207a65726f206164647265737345524332304f6e417070726f76653a207370656e64657220646f65736e277420737570706f7274206f6e417070726f766545524332303a207472616e7366657220616d6f756e7420657863656564732062616c616e63656f6e417070726f766528616464726573732c616464726573732c75696e743235362c62797465732945524332303a207472616e7366657220616d6f756e74206578636565647320616c6c6f77616e636545524332303a207472616e736665722066726f6d20746865207a65726f206164647265737345524332303a20617070726f76652066726f6d20746865207a65726f206164647265737345524332303a2064656372656173656420616c6c6f77616e63652062656c6f77207a65726fa265627a7a723158205f38639f33e199b2a44d46fed86192ba1daf836a08c11f82aa877b89bb0dc9fc64736f6c634300050c0032"

// DeployERC20OnApprove deploys a new Ethereum contract, binding an instance of ERC20OnApprove to it.
func DeployERC20OnApprove(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ERC20OnApprove, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20OnApproveABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC20OnApproveBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC20OnApprove{ERC20OnApproveCaller: ERC20OnApproveCaller{contract: contract}, ERC20OnApproveTransactor: ERC20OnApproveTransactor{contract: contract}, ERC20OnApproveFilterer: ERC20OnApproveFilterer{contract: contract}}, nil
}

// ERC20OnApprove is an auto generated Go binding around an Ethereum contract.
type ERC20OnApprove struct {
	ERC20OnApproveCaller     // Read-only binding to the contract
	ERC20OnApproveTransactor // Write-only binding to the contract
	ERC20OnApproveFilterer   // Log filterer for contract events
}

// ERC20OnApproveCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20OnApproveCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20OnApproveTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20OnApproveTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20OnApproveFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20OnApproveFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20OnApproveSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20OnApproveSession struct {
	Contract     *ERC20OnApprove   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20OnApproveCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20OnApproveCallerSession struct {
	Contract *ERC20OnApproveCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// ERC20OnApproveTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20OnApproveTransactorSession struct {
	Contract     *ERC20OnApproveTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ERC20OnApproveRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20OnApproveRaw struct {
	Contract *ERC20OnApprove // Generic contract binding to access the raw methods on
}

// ERC20OnApproveCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20OnApproveCallerRaw struct {
	Contract *ERC20OnApproveCaller // Generic read-only contract binding to access the raw methods on
}

// ERC20OnApproveTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20OnApproveTransactorRaw struct {
	Contract *ERC20OnApproveTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20OnApprove creates a new instance of ERC20OnApprove, bound to a specific deployed contract.
func NewERC20OnApprove(address common.Address, backend bind.ContractBackend) (*ERC20OnApprove, error) {
	contract, err := bindERC20OnApprove(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20OnApprove{ERC20OnApproveCaller: ERC20OnApproveCaller{contract: contract}, ERC20OnApproveTransactor: ERC20OnApproveTransactor{contract: contract}, ERC20OnApproveFilterer: ERC20OnApproveFilterer{contract: contract}}, nil
}

// NewERC20OnApproveCaller creates a new read-only instance of ERC20OnApprove, bound to a specific deployed contract.
func NewERC20OnApproveCaller(address common.Address, caller bind.ContractCaller) (*ERC20OnApproveCaller, error) {
	contract, err := bindERC20OnApprove(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20OnApproveCaller{contract: contract}, nil
}

// NewERC20OnApproveTransactor creates a new write-only instance of ERC20OnApprove, bound to a specific deployed contract.
func NewERC20OnApproveTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC20OnApproveTransactor, error) {
	contract, err := bindERC20OnApprove(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20OnApproveTransactor{contract: contract}, nil
}

// NewERC20OnApproveFilterer creates a new log filterer instance of ERC20OnApprove, bound to a specific deployed contract.
func NewERC20OnApproveFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC20OnApproveFilterer, error) {
	contract, err := bindERC20OnApprove(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20OnApproveFilterer{contract: contract}, nil
}

// bindERC20OnApprove binds a generic wrapper to an already deployed contract.
func bindERC20OnApprove(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20OnApproveABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20OnApprove *ERC20OnApproveRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20OnApprove.Contract.ERC20OnApproveCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20OnApprove *ERC20OnApproveRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20OnApprove.Contract.ERC20OnApproveTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20OnApprove *ERC20OnApproveRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20OnApprove.Contract.ERC20OnApproveTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20OnApprove *ERC20OnApproveCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20OnApprove.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20OnApprove *ERC20OnApproveTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20OnApprove.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20OnApprove *ERC20OnApproveTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20OnApprove.Contract.contract.Transact(opts, method, params...)
}

// INTERFACEIDONAPPROVE is a free data retrieval call binding the contract method 0x6cd28f9a.
//
// Solidity: function INTERFACE_ID_ON_APPROVE() constant returns(bytes4)
func (_ERC20OnApprove *ERC20OnApproveCaller) INTERFACEIDONAPPROVE(opts *bind.CallOpts) ([4]byte, error) {
	var (
		ret0 = new([4]byte)
	)
	out := ret0
	err := _ERC20OnApprove.contract.Call(opts, out, "INTERFACE_ID_ON_APPROVE")
	return *ret0, err
}

// INTERFACEIDONAPPROVE is a free data retrieval call binding the contract method 0x6cd28f9a.
//
// Solidity: function INTERFACE_ID_ON_APPROVE() constant returns(bytes4)
func (_ERC20OnApprove *ERC20OnApproveSession) INTERFACEIDONAPPROVE() ([4]byte, error) {
	return _ERC20OnApprove.Contract.INTERFACEIDONAPPROVE(&_ERC20OnApprove.CallOpts)
}

// INTERFACEIDONAPPROVE is a free data retrieval call binding the contract method 0x6cd28f9a.
//
// Solidity: function INTERFACE_ID_ON_APPROVE() constant returns(bytes4)
func (_ERC20OnApprove *ERC20OnApproveCallerSession) INTERFACEIDONAPPROVE() ([4]byte, error) {
	return _ERC20OnApprove.Contract.INTERFACEIDONAPPROVE(&_ERC20OnApprove.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_ERC20OnApprove *ERC20OnApproveCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20OnApprove.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_ERC20OnApprove *ERC20OnApproveSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20OnApprove.Contract.Allowance(&_ERC20OnApprove.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_ERC20OnApprove *ERC20OnApproveCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20OnApprove.Contract.Allowance(&_ERC20OnApprove.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_ERC20OnApprove *ERC20OnApproveCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20OnApprove.contract.Call(opts, out, "balanceOf", account)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_ERC20OnApprove *ERC20OnApproveSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ERC20OnApprove.Contract.BalanceOf(&_ERC20OnApprove.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_ERC20OnApprove *ERC20OnApproveCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ERC20OnApprove.Contract.BalanceOf(&_ERC20OnApprove.CallOpts, account)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20OnApprove *ERC20OnApproveCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20OnApprove.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20OnApprove *ERC20OnApproveSession) TotalSupply() (*big.Int, error) {
	return _ERC20OnApprove.Contract.TotalSupply(&_ERC20OnApprove.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20OnApprove *ERC20OnApproveCallerSession) TotalSupply() (*big.Int, error) {
	return _ERC20OnApprove.Contract.TotalSupply(&_ERC20OnApprove.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC20OnApprove *ERC20OnApproveTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20OnApprove.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC20OnApprove *ERC20OnApproveSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20OnApprove.Contract.Approve(&_ERC20OnApprove.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC20OnApprove *ERC20OnApproveTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20OnApprove.Contract.Approve(&_ERC20OnApprove.TransactOpts, spender, amount)
}

// ApproveAndCall is a paid mutator transaction binding the contract method 0xcae9ca51.
//
// Solidity: function approveAndCall(address spender, uint256 amount, bytes data) returns(bool)
func (_ERC20OnApprove *ERC20OnApproveTransactor) ApproveAndCall(opts *bind.TransactOpts, spender common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _ERC20OnApprove.contract.Transact(opts, "approveAndCall", spender, amount, data)
}

// ApproveAndCall is a paid mutator transaction binding the contract method 0xcae9ca51.
//
// Solidity: function approveAndCall(address spender, uint256 amount, bytes data) returns(bool)
func (_ERC20OnApprove *ERC20OnApproveSession) ApproveAndCall(spender common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _ERC20OnApprove.Contract.ApproveAndCall(&_ERC20OnApprove.TransactOpts, spender, amount, data)
}

// ApproveAndCall is a paid mutator transaction binding the contract method 0xcae9ca51.
//
// Solidity: function approveAndCall(address spender, uint256 amount, bytes data) returns(bool)
func (_ERC20OnApprove *ERC20OnApproveTransactorSession) ApproveAndCall(spender common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _ERC20OnApprove.Contract.ApproveAndCall(&_ERC20OnApprove.TransactOpts, spender, amount, data)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20OnApprove *ERC20OnApproveTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20OnApprove.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20OnApprove *ERC20OnApproveSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20OnApprove.Contract.DecreaseAllowance(&_ERC20OnApprove.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20OnApprove *ERC20OnApproveTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20OnApprove.Contract.DecreaseAllowance(&_ERC20OnApprove.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20OnApprove *ERC20OnApproveTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20OnApprove.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20OnApprove *ERC20OnApproveSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20OnApprove.Contract.IncreaseAllowance(&_ERC20OnApprove.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20OnApprove *ERC20OnApproveTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20OnApprove.Contract.IncreaseAllowance(&_ERC20OnApprove.TransactOpts, spender, addedValue)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20OnApprove *ERC20OnApproveTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20OnApprove.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20OnApprove *ERC20OnApproveSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20OnApprove.Contract.Transfer(&_ERC20OnApprove.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20OnApprove *ERC20OnApproveTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20OnApprove.Contract.Transfer(&_ERC20OnApprove.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20OnApprove *ERC20OnApproveTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20OnApprove.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20OnApprove *ERC20OnApproveSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20OnApprove.Contract.TransferFrom(&_ERC20OnApprove.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20OnApprove *ERC20OnApproveTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20OnApprove.Contract.TransferFrom(&_ERC20OnApprove.TransactOpts, sender, recipient, amount)
}

// ERC20OnApproveApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC20OnApprove contract.
type ERC20OnApproveApprovalIterator struct {
	Event *ERC20OnApproveApproval // Event containing the contract specifics and raw log

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
func (it *ERC20OnApproveApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20OnApproveApproval)
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
		it.Event = new(ERC20OnApproveApproval)
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
func (it *ERC20OnApproveApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20OnApproveApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20OnApproveApproval represents a Approval event raised by the ERC20OnApprove contract.
type ERC20OnApproveApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20OnApprove *ERC20OnApproveFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ERC20OnApproveApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20OnApprove.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ERC20OnApproveApprovalIterator{contract: _ERC20OnApprove.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20OnApprove *ERC20OnApproveFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC20OnApproveApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20OnApprove.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20OnApproveApproval)
				if err := _ERC20OnApprove.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_ERC20OnApprove *ERC20OnApproveFilterer) ParseApproval(log types.Log) (*ERC20OnApproveApproval, error) {
	event := new(ERC20OnApproveApproval)
	if err := _ERC20OnApprove.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ERC20OnApproveTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC20OnApprove contract.
type ERC20OnApproveTransferIterator struct {
	Event *ERC20OnApproveTransfer // Event containing the contract specifics and raw log

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
func (it *ERC20OnApproveTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20OnApproveTransfer)
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
		it.Event = new(ERC20OnApproveTransfer)
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
func (it *ERC20OnApproveTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20OnApproveTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20OnApproveTransfer represents a Transfer event raised by the ERC20OnApprove contract.
type ERC20OnApproveTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20OnApprove *ERC20OnApproveFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ERC20OnApproveTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20OnApprove.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERC20OnApproveTransferIterator{contract: _ERC20OnApprove.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20OnApprove *ERC20OnApproveFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC20OnApproveTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20OnApprove.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20OnApproveTransfer)
				if err := _ERC20OnApprove.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_ERC20OnApprove *ERC20OnApproveFilterer) ParseTransfer(log types.Log) (*ERC20OnApproveTransfer, error) {
	event := new(ERC20OnApproveTransfer)
	if err := _ERC20OnApprove.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IERC165ABI is the input ABI used to generate the binding from.
const IERC165ABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// IERC165FuncSigs maps the 4-byte function signature to its string representation.
var IERC165FuncSigs = map[string]string{
	"01ffc9a7": "supportsInterface(bytes4)",
}

// IERC165 is an auto generated Go binding around an Ethereum contract.
type IERC165 struct {
	IERC165Caller     // Read-only binding to the contract
	IERC165Transactor // Write-only binding to the contract
	IERC165Filterer   // Log filterer for contract events
}

// IERC165Caller is an auto generated read-only Go binding around an Ethereum contract.
type IERC165Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC165Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC165Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC165Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC165Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC165Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC165Session struct {
	Contract     *IERC165          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC165CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC165CallerSession struct {
	Contract *IERC165Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// IERC165TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC165TransactorSession struct {
	Contract     *IERC165Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IERC165Raw is an auto generated low-level Go binding around an Ethereum contract.
type IERC165Raw struct {
	Contract *IERC165 // Generic contract binding to access the raw methods on
}

// IERC165CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC165CallerRaw struct {
	Contract *IERC165Caller // Generic read-only contract binding to access the raw methods on
}

// IERC165TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC165TransactorRaw struct {
	Contract *IERC165Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC165 creates a new instance of IERC165, bound to a specific deployed contract.
func NewIERC165(address common.Address, backend bind.ContractBackend) (*IERC165, error) {
	contract, err := bindIERC165(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC165{IERC165Caller: IERC165Caller{contract: contract}, IERC165Transactor: IERC165Transactor{contract: contract}, IERC165Filterer: IERC165Filterer{contract: contract}}, nil
}

// NewIERC165Caller creates a new read-only instance of IERC165, bound to a specific deployed contract.
func NewIERC165Caller(address common.Address, caller bind.ContractCaller) (*IERC165Caller, error) {
	contract, err := bindIERC165(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC165Caller{contract: contract}, nil
}

// NewIERC165Transactor creates a new write-only instance of IERC165, bound to a specific deployed contract.
func NewIERC165Transactor(address common.Address, transactor bind.ContractTransactor) (*IERC165Transactor, error) {
	contract, err := bindIERC165(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC165Transactor{contract: contract}, nil
}

// NewIERC165Filterer creates a new log filterer instance of IERC165, bound to a specific deployed contract.
func NewIERC165Filterer(address common.Address, filterer bind.ContractFilterer) (*IERC165Filterer, error) {
	contract, err := bindIERC165(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC165Filterer{contract: contract}, nil
}

// bindIERC165 binds a generic wrapper to an already deployed contract.
func bindIERC165(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC165ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC165 *IERC165Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IERC165.Contract.IERC165Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC165 *IERC165Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC165.Contract.IERC165Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC165 *IERC165Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC165.Contract.IERC165Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC165 *IERC165CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IERC165.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC165 *IERC165TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC165.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC165 *IERC165TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC165.Contract.contract.Transact(opts, method, params...)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) constant returns(bool)
func (_IERC165 *IERC165Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IERC165.contract.Call(opts, out, "supportsInterface", interfaceId)
	return *ret0, err
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) constant returns(bool)
func (_IERC165 *IERC165Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IERC165.Contract.SupportsInterface(&_IERC165.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) constant returns(bool)
func (_IERC165 *IERC165CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IERC165.Contract.SupportsInterface(&_IERC165.CallOpts, interfaceId)
}

// IERC20ABI is the input ABI used to generate the binding from.
const IERC20ABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IERC20FuncSigs maps the 4-byte function signature to its string representation.
var IERC20FuncSigs = map[string]string{
	"dd62ed3e": "allowance(address,address)",
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
}

// IERC20 is an auto generated Go binding around an Ethereum contract.
type IERC20 struct {
	IERC20Caller     // Read-only binding to the contract
	IERC20Transactor // Write-only binding to the contract
	IERC20Filterer   // Log filterer for contract events
}

// IERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type IERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC20Session struct {
	Contract     *IERC20           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC20CallerSession struct {
	Contract *IERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC20TransactorSession struct {
	Contract     *IERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type IERC20Raw struct {
	Contract *IERC20 // Generic contract binding to access the raw methods on
}

// IERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC20CallerRaw struct {
	Contract *IERC20Caller // Generic read-only contract binding to access the raw methods on
}

// IERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC20TransactorRaw struct {
	Contract *IERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC20 creates a new instance of IERC20, bound to a specific deployed contract.
func NewIERC20(address common.Address, backend bind.ContractBackend) (*IERC20, error) {
	contract, err := bindIERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC20{IERC20Caller: IERC20Caller{contract: contract}, IERC20Transactor: IERC20Transactor{contract: contract}, IERC20Filterer: IERC20Filterer{contract: contract}}, nil
}

// NewIERC20Caller creates a new read-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Caller(address common.Address, caller bind.ContractCaller) (*IERC20Caller, error) {
	contract, err := bindIERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Caller{contract: contract}, nil
}

// NewIERC20Transactor creates a new write-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*IERC20Transactor, error) {
	contract, err := bindIERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Transactor{contract: contract}, nil
}

// NewIERC20Filterer creates a new log filterer instance of IERC20, bound to a specific deployed contract.
func NewIERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*IERC20Filterer, error) {
	contract, err := bindIERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC20Filterer{contract: contract}, nil
}

// bindIERC20 binds a generic wrapper to an already deployed contract.
func bindIERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.IERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_IERC20 *IERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IERC20.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_IERC20 *IERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_IERC20 *IERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_IERC20 *IERC20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IERC20.contract.Call(opts, out, "balanceOf", account)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_IERC20 *IERC20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_IERC20 *IERC20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, account)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_IERC20 *IERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IERC20.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_IERC20 *IERC20Session) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_IERC20 *IERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, sender, recipient, amount)
}

// IERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IERC20 contract.
type IERC20ApprovalIterator struct {
	Event *IERC20Approval // Event containing the contract specifics and raw log

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
func (it *IERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Approval)
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
		it.Event = new(IERC20Approval)
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
func (it *IERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Approval represents a Approval event raised by the IERC20 contract.
type IERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*IERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &IERC20ApprovalIterator{contract: _IERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Approval)
				if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_IERC20 *IERC20Filterer) ParseApproval(log types.Log) (*IERC20Approval, error) {
	event := new(IERC20Approval)
	if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IERC20 contract.
type IERC20TransferIterator struct {
	Event *IERC20Transfer // Event containing the contract specifics and raw log

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
func (it *IERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Transfer)
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
		it.Event = new(IERC20Transfer)
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
func (it *IERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Transfer represents a Transfer event raised by the IERC20 contract.
type IERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IERC20TransferIterator{contract: _IERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Transfer)
				if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_IERC20 *IERC20Filterer) ParseTransfer(log types.Log) (*IERC20Transfer, error) {
	event := new(IERC20Transfer)
	if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// MinterRoleABI is the input ABI used to generate the binding from.
const MinterRoleABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"MinterAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"MinterRemoved\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addMinter\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isMinter\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceMinter\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// MinterRoleFuncSigs maps the 4-byte function signature to its string representation.
var MinterRoleFuncSigs = map[string]string{
	"983b2d56": "addMinter(address)",
	"aa271e1a": "isMinter(address)",
	"98650275": "renounceMinter()",
}

// MinterRole is an auto generated Go binding around an Ethereum contract.
type MinterRole struct {
	MinterRoleCaller     // Read-only binding to the contract
	MinterRoleTransactor // Write-only binding to the contract
	MinterRoleFilterer   // Log filterer for contract events
}

// MinterRoleCaller is an auto generated read-only Go binding around an Ethereum contract.
type MinterRoleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MinterRoleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MinterRoleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MinterRoleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MinterRoleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MinterRoleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MinterRoleSession struct {
	Contract     *MinterRole       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MinterRoleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MinterRoleCallerSession struct {
	Contract *MinterRoleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// MinterRoleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MinterRoleTransactorSession struct {
	Contract     *MinterRoleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// MinterRoleRaw is an auto generated low-level Go binding around an Ethereum contract.
type MinterRoleRaw struct {
	Contract *MinterRole // Generic contract binding to access the raw methods on
}

// MinterRoleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MinterRoleCallerRaw struct {
	Contract *MinterRoleCaller // Generic read-only contract binding to access the raw methods on
}

// MinterRoleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MinterRoleTransactorRaw struct {
	Contract *MinterRoleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMinterRole creates a new instance of MinterRole, bound to a specific deployed contract.
func NewMinterRole(address common.Address, backend bind.ContractBackend) (*MinterRole, error) {
	contract, err := bindMinterRole(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MinterRole{MinterRoleCaller: MinterRoleCaller{contract: contract}, MinterRoleTransactor: MinterRoleTransactor{contract: contract}, MinterRoleFilterer: MinterRoleFilterer{contract: contract}}, nil
}

// NewMinterRoleCaller creates a new read-only instance of MinterRole, bound to a specific deployed contract.
func NewMinterRoleCaller(address common.Address, caller bind.ContractCaller) (*MinterRoleCaller, error) {
	contract, err := bindMinterRole(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MinterRoleCaller{contract: contract}, nil
}

// NewMinterRoleTransactor creates a new write-only instance of MinterRole, bound to a specific deployed contract.
func NewMinterRoleTransactor(address common.Address, transactor bind.ContractTransactor) (*MinterRoleTransactor, error) {
	contract, err := bindMinterRole(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MinterRoleTransactor{contract: contract}, nil
}

// NewMinterRoleFilterer creates a new log filterer instance of MinterRole, bound to a specific deployed contract.
func NewMinterRoleFilterer(address common.Address, filterer bind.ContractFilterer) (*MinterRoleFilterer, error) {
	contract, err := bindMinterRole(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MinterRoleFilterer{contract: contract}, nil
}

// bindMinterRole binds a generic wrapper to an already deployed contract.
func bindMinterRole(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MinterRoleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MinterRole *MinterRoleRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MinterRole.Contract.MinterRoleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MinterRole *MinterRoleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MinterRole.Contract.MinterRoleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MinterRole *MinterRoleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MinterRole.Contract.MinterRoleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MinterRole *MinterRoleCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MinterRole.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MinterRole *MinterRoleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MinterRole.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MinterRole *MinterRoleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MinterRole.Contract.contract.Transact(opts, method, params...)
}

// IsMinter is a free data retrieval call binding the contract method 0xaa271e1a.
//
// Solidity: function isMinter(address account) constant returns(bool)
func (_MinterRole *MinterRoleCaller) IsMinter(opts *bind.CallOpts, account common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _MinterRole.contract.Call(opts, out, "isMinter", account)
	return *ret0, err
}

// IsMinter is a free data retrieval call binding the contract method 0xaa271e1a.
//
// Solidity: function isMinter(address account) constant returns(bool)
func (_MinterRole *MinterRoleSession) IsMinter(account common.Address) (bool, error) {
	return _MinterRole.Contract.IsMinter(&_MinterRole.CallOpts, account)
}

// IsMinter is a free data retrieval call binding the contract method 0xaa271e1a.
//
// Solidity: function isMinter(address account) constant returns(bool)
func (_MinterRole *MinterRoleCallerSession) IsMinter(account common.Address) (bool, error) {
	return _MinterRole.Contract.IsMinter(&_MinterRole.CallOpts, account)
}

// AddMinter is a paid mutator transaction binding the contract method 0x983b2d56.
//
// Solidity: function addMinter(address account) returns()
func (_MinterRole *MinterRoleTransactor) AddMinter(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _MinterRole.contract.Transact(opts, "addMinter", account)
}

// AddMinter is a paid mutator transaction binding the contract method 0x983b2d56.
//
// Solidity: function addMinter(address account) returns()
func (_MinterRole *MinterRoleSession) AddMinter(account common.Address) (*types.Transaction, error) {
	return _MinterRole.Contract.AddMinter(&_MinterRole.TransactOpts, account)
}

// AddMinter is a paid mutator transaction binding the contract method 0x983b2d56.
//
// Solidity: function addMinter(address account) returns()
func (_MinterRole *MinterRoleTransactorSession) AddMinter(account common.Address) (*types.Transaction, error) {
	return _MinterRole.Contract.AddMinter(&_MinterRole.TransactOpts, account)
}

// RenounceMinter is a paid mutator transaction binding the contract method 0x98650275.
//
// Solidity: function renounceMinter() returns()
func (_MinterRole *MinterRoleTransactor) RenounceMinter(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MinterRole.contract.Transact(opts, "renounceMinter")
}

// RenounceMinter is a paid mutator transaction binding the contract method 0x98650275.
//
// Solidity: function renounceMinter() returns()
func (_MinterRole *MinterRoleSession) RenounceMinter() (*types.Transaction, error) {
	return _MinterRole.Contract.RenounceMinter(&_MinterRole.TransactOpts)
}

// RenounceMinter is a paid mutator transaction binding the contract method 0x98650275.
//
// Solidity: function renounceMinter() returns()
func (_MinterRole *MinterRoleTransactorSession) RenounceMinter() (*types.Transaction, error) {
	return _MinterRole.Contract.RenounceMinter(&_MinterRole.TransactOpts)
}

// MinterRoleMinterAddedIterator is returned from FilterMinterAdded and is used to iterate over the raw logs and unpacked data for MinterAdded events raised by the MinterRole contract.
type MinterRoleMinterAddedIterator struct {
	Event *MinterRoleMinterAdded // Event containing the contract specifics and raw log

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
func (it *MinterRoleMinterAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MinterRoleMinterAdded)
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
		it.Event = new(MinterRoleMinterAdded)
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
func (it *MinterRoleMinterAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MinterRoleMinterAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MinterRoleMinterAdded represents a MinterAdded event raised by the MinterRole contract.
type MinterRoleMinterAdded struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMinterAdded is a free log retrieval operation binding the contract event 0x6ae172837ea30b801fbfcdd4108aa1d5bf8ff775444fd70256b44e6bf3dfc3f6.
//
// Solidity: event MinterAdded(address indexed account)
func (_MinterRole *MinterRoleFilterer) FilterMinterAdded(opts *bind.FilterOpts, account []common.Address) (*MinterRoleMinterAddedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _MinterRole.contract.FilterLogs(opts, "MinterAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return &MinterRoleMinterAddedIterator{contract: _MinterRole.contract, event: "MinterAdded", logs: logs, sub: sub}, nil
}

// WatchMinterAdded is a free log subscription operation binding the contract event 0x6ae172837ea30b801fbfcdd4108aa1d5bf8ff775444fd70256b44e6bf3dfc3f6.
//
// Solidity: event MinterAdded(address indexed account)
func (_MinterRole *MinterRoleFilterer) WatchMinterAdded(opts *bind.WatchOpts, sink chan<- *MinterRoleMinterAdded, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _MinterRole.contract.WatchLogs(opts, "MinterAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MinterRoleMinterAdded)
				if err := _MinterRole.contract.UnpackLog(event, "MinterAdded", log); err != nil {
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

// ParseMinterAdded is a log parse operation binding the contract event 0x6ae172837ea30b801fbfcdd4108aa1d5bf8ff775444fd70256b44e6bf3dfc3f6.
//
// Solidity: event MinterAdded(address indexed account)
func (_MinterRole *MinterRoleFilterer) ParseMinterAdded(log types.Log) (*MinterRoleMinterAdded, error) {
	event := new(MinterRoleMinterAdded)
	if err := _MinterRole.contract.UnpackLog(event, "MinterAdded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// MinterRoleMinterRemovedIterator is returned from FilterMinterRemoved and is used to iterate over the raw logs and unpacked data for MinterRemoved events raised by the MinterRole contract.
type MinterRoleMinterRemovedIterator struct {
	Event *MinterRoleMinterRemoved // Event containing the contract specifics and raw log

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
func (it *MinterRoleMinterRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MinterRoleMinterRemoved)
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
		it.Event = new(MinterRoleMinterRemoved)
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
func (it *MinterRoleMinterRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MinterRoleMinterRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MinterRoleMinterRemoved represents a MinterRemoved event raised by the MinterRole contract.
type MinterRoleMinterRemoved struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMinterRemoved is a free log retrieval operation binding the contract event 0xe94479a9f7e1952cc78f2d6baab678adc1b772d936c6583def489e524cb66692.
//
// Solidity: event MinterRemoved(address indexed account)
func (_MinterRole *MinterRoleFilterer) FilterMinterRemoved(opts *bind.FilterOpts, account []common.Address) (*MinterRoleMinterRemovedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _MinterRole.contract.FilterLogs(opts, "MinterRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return &MinterRoleMinterRemovedIterator{contract: _MinterRole.contract, event: "MinterRemoved", logs: logs, sub: sub}, nil
}

// WatchMinterRemoved is a free log subscription operation binding the contract event 0xe94479a9f7e1952cc78f2d6baab678adc1b772d936c6583def489e524cb66692.
//
// Solidity: event MinterRemoved(address indexed account)
func (_MinterRole *MinterRoleFilterer) WatchMinterRemoved(opts *bind.WatchOpts, sink chan<- *MinterRoleMinterRemoved, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _MinterRole.contract.WatchLogs(opts, "MinterRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MinterRoleMinterRemoved)
				if err := _MinterRole.contract.UnpackLog(event, "MinterRemoved", log); err != nil {
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

// ParseMinterRemoved is a log parse operation binding the contract event 0xe94479a9f7e1952cc78f2d6baab678adc1b772d936c6583def489e524cb66692.
//
// Solidity: event MinterRemoved(address indexed account)
func (_MinterRole *MinterRoleFilterer) ParseMinterRemoved(log types.Log) (*MinterRoleMinterRemoved, error) {
	event := new(MinterRoleMinterRemoved)
	if err := _MinterRole.contract.UnpackLog(event, "MinterRemoved", log); err != nil {
		return nil, err
	}
	return event, nil
}

// MinterRoleRenounceTargetABI is the input ABI used to generate the binding from.
const MinterRoleRenounceTargetABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"renounceMinter\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// MinterRoleRenounceTargetFuncSigs maps the 4-byte function signature to its string representation.
var MinterRoleRenounceTargetFuncSigs = map[string]string{
	"98650275": "renounceMinter()",
}

// MinterRoleRenounceTarget is an auto generated Go binding around an Ethereum contract.
type MinterRoleRenounceTarget struct {
	MinterRoleRenounceTargetCaller     // Read-only binding to the contract
	MinterRoleRenounceTargetTransactor // Write-only binding to the contract
	MinterRoleRenounceTargetFilterer   // Log filterer for contract events
}

// MinterRoleRenounceTargetCaller is an auto generated read-only Go binding around an Ethereum contract.
type MinterRoleRenounceTargetCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MinterRoleRenounceTargetTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MinterRoleRenounceTargetTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MinterRoleRenounceTargetFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MinterRoleRenounceTargetFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MinterRoleRenounceTargetSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MinterRoleRenounceTargetSession struct {
	Contract     *MinterRoleRenounceTarget // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// MinterRoleRenounceTargetCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MinterRoleRenounceTargetCallerSession struct {
	Contract *MinterRoleRenounceTargetCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// MinterRoleRenounceTargetTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MinterRoleRenounceTargetTransactorSession struct {
	Contract     *MinterRoleRenounceTargetTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// MinterRoleRenounceTargetRaw is an auto generated low-level Go binding around an Ethereum contract.
type MinterRoleRenounceTargetRaw struct {
	Contract *MinterRoleRenounceTarget // Generic contract binding to access the raw methods on
}

// MinterRoleRenounceTargetCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MinterRoleRenounceTargetCallerRaw struct {
	Contract *MinterRoleRenounceTargetCaller // Generic read-only contract binding to access the raw methods on
}

// MinterRoleRenounceTargetTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MinterRoleRenounceTargetTransactorRaw struct {
	Contract *MinterRoleRenounceTargetTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMinterRoleRenounceTarget creates a new instance of MinterRoleRenounceTarget, bound to a specific deployed contract.
func NewMinterRoleRenounceTarget(address common.Address, backend bind.ContractBackend) (*MinterRoleRenounceTarget, error) {
	contract, err := bindMinterRoleRenounceTarget(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MinterRoleRenounceTarget{MinterRoleRenounceTargetCaller: MinterRoleRenounceTargetCaller{contract: contract}, MinterRoleRenounceTargetTransactor: MinterRoleRenounceTargetTransactor{contract: contract}, MinterRoleRenounceTargetFilterer: MinterRoleRenounceTargetFilterer{contract: contract}}, nil
}

// NewMinterRoleRenounceTargetCaller creates a new read-only instance of MinterRoleRenounceTarget, bound to a specific deployed contract.
func NewMinterRoleRenounceTargetCaller(address common.Address, caller bind.ContractCaller) (*MinterRoleRenounceTargetCaller, error) {
	contract, err := bindMinterRoleRenounceTarget(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MinterRoleRenounceTargetCaller{contract: contract}, nil
}

// NewMinterRoleRenounceTargetTransactor creates a new write-only instance of MinterRoleRenounceTarget, bound to a specific deployed contract.
func NewMinterRoleRenounceTargetTransactor(address common.Address, transactor bind.ContractTransactor) (*MinterRoleRenounceTargetTransactor, error) {
	contract, err := bindMinterRoleRenounceTarget(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MinterRoleRenounceTargetTransactor{contract: contract}, nil
}

// NewMinterRoleRenounceTargetFilterer creates a new log filterer instance of MinterRoleRenounceTarget, bound to a specific deployed contract.
func NewMinterRoleRenounceTargetFilterer(address common.Address, filterer bind.ContractFilterer) (*MinterRoleRenounceTargetFilterer, error) {
	contract, err := bindMinterRoleRenounceTarget(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MinterRoleRenounceTargetFilterer{contract: contract}, nil
}

// bindMinterRoleRenounceTarget binds a generic wrapper to an already deployed contract.
func bindMinterRoleRenounceTarget(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MinterRoleRenounceTargetABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MinterRoleRenounceTarget *MinterRoleRenounceTargetRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MinterRoleRenounceTarget.Contract.MinterRoleRenounceTargetCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MinterRoleRenounceTarget *MinterRoleRenounceTargetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MinterRoleRenounceTarget.Contract.MinterRoleRenounceTargetTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MinterRoleRenounceTarget *MinterRoleRenounceTargetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MinterRoleRenounceTarget.Contract.MinterRoleRenounceTargetTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MinterRoleRenounceTarget *MinterRoleRenounceTargetCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MinterRoleRenounceTarget.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MinterRoleRenounceTarget *MinterRoleRenounceTargetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MinterRoleRenounceTarget.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MinterRoleRenounceTarget *MinterRoleRenounceTargetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MinterRoleRenounceTarget.Contract.contract.Transact(opts, method, params...)
}

// RenounceMinter is a paid mutator transaction binding the contract method 0x98650275.
//
// Solidity: function renounceMinter() returns()
func (_MinterRoleRenounceTarget *MinterRoleRenounceTargetTransactor) RenounceMinter(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MinterRoleRenounceTarget.contract.Transact(opts, "renounceMinter")
}

// RenounceMinter is a paid mutator transaction binding the contract method 0x98650275.
//
// Solidity: function renounceMinter() returns()
func (_MinterRoleRenounceTarget *MinterRoleRenounceTargetSession) RenounceMinter() (*types.Transaction, error) {
	return _MinterRoleRenounceTarget.Contract.RenounceMinter(&_MinterRoleRenounceTarget.TransactOpts)
}

// RenounceMinter is a paid mutator transaction binding the contract method 0x98650275.
//
// Solidity: function renounceMinter() returns()
func (_MinterRoleRenounceTarget *MinterRoleRenounceTargetTransactorSession) RenounceMinter() (*types.Transaction, error) {
	return _MinterRoleRenounceTarget.Contract.RenounceMinter(&_MinterRoleRenounceTarget.TransactOpts)
}

// OnApproveABI is the input ABI used to generate the binding from.
const OnApproveABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"constant\":true,\"inputs\":[],\"name\":\"INTERFACE_ID_ON_APPROVE\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onApprove\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// OnApproveFuncSigs maps the 4-byte function signature to its string representation.
var OnApproveFuncSigs = map[string]string{
	"6cd28f9a": "INTERFACE_ID_ON_APPROVE()",
	"4273ca16": "onApprove(address,address,uint256,bytes)",
	"01ffc9a7": "supportsInterface(bytes4)",
}

// OnApprove is an auto generated Go binding around an Ethereum contract.
type OnApprove struct {
	OnApproveCaller     // Read-only binding to the contract
	OnApproveTransactor // Write-only binding to the contract
	OnApproveFilterer   // Log filterer for contract events
}

// OnApproveCaller is an auto generated read-only Go binding around an Ethereum contract.
type OnApproveCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OnApproveTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OnApproveTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OnApproveFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OnApproveFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OnApproveSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OnApproveSession struct {
	Contract     *OnApprove        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OnApproveCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OnApproveCallerSession struct {
	Contract *OnApproveCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// OnApproveTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OnApproveTransactorSession struct {
	Contract     *OnApproveTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// OnApproveRaw is an auto generated low-level Go binding around an Ethereum contract.
type OnApproveRaw struct {
	Contract *OnApprove // Generic contract binding to access the raw methods on
}

// OnApproveCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OnApproveCallerRaw struct {
	Contract *OnApproveCaller // Generic read-only contract binding to access the raw methods on
}

// OnApproveTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OnApproveTransactorRaw struct {
	Contract *OnApproveTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOnApprove creates a new instance of OnApprove, bound to a specific deployed contract.
func NewOnApprove(address common.Address, backend bind.ContractBackend) (*OnApprove, error) {
	contract, err := bindOnApprove(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OnApprove{OnApproveCaller: OnApproveCaller{contract: contract}, OnApproveTransactor: OnApproveTransactor{contract: contract}, OnApproveFilterer: OnApproveFilterer{contract: contract}}, nil
}

// NewOnApproveCaller creates a new read-only instance of OnApprove, bound to a specific deployed contract.
func NewOnApproveCaller(address common.Address, caller bind.ContractCaller) (*OnApproveCaller, error) {
	contract, err := bindOnApprove(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OnApproveCaller{contract: contract}, nil
}

// NewOnApproveTransactor creates a new write-only instance of OnApprove, bound to a specific deployed contract.
func NewOnApproveTransactor(address common.Address, transactor bind.ContractTransactor) (*OnApproveTransactor, error) {
	contract, err := bindOnApprove(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OnApproveTransactor{contract: contract}, nil
}

// NewOnApproveFilterer creates a new log filterer instance of OnApprove, bound to a specific deployed contract.
func NewOnApproveFilterer(address common.Address, filterer bind.ContractFilterer) (*OnApproveFilterer, error) {
	contract, err := bindOnApprove(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OnApproveFilterer{contract: contract}, nil
}

// bindOnApprove binds a generic wrapper to an already deployed contract.
func bindOnApprove(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OnApproveABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OnApprove *OnApproveRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _OnApprove.Contract.OnApproveCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OnApprove *OnApproveRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OnApprove.Contract.OnApproveTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OnApprove *OnApproveRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OnApprove.Contract.OnApproveTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OnApprove *OnApproveCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _OnApprove.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OnApprove *OnApproveTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OnApprove.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OnApprove *OnApproveTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OnApprove.Contract.contract.Transact(opts, method, params...)
}

// INTERFACEIDONAPPROVE is a free data retrieval call binding the contract method 0x6cd28f9a.
//
// Solidity: function INTERFACE_ID_ON_APPROVE() constant returns(bytes4)
func (_OnApprove *OnApproveCaller) INTERFACEIDONAPPROVE(opts *bind.CallOpts) ([4]byte, error) {
	var (
		ret0 = new([4]byte)
	)
	out := ret0
	err := _OnApprove.contract.Call(opts, out, "INTERFACE_ID_ON_APPROVE")
	return *ret0, err
}

// INTERFACEIDONAPPROVE is a free data retrieval call binding the contract method 0x6cd28f9a.
//
// Solidity: function INTERFACE_ID_ON_APPROVE() constant returns(bytes4)
func (_OnApprove *OnApproveSession) INTERFACEIDONAPPROVE() ([4]byte, error) {
	return _OnApprove.Contract.INTERFACEIDONAPPROVE(&_OnApprove.CallOpts)
}

// INTERFACEIDONAPPROVE is a free data retrieval call binding the contract method 0x6cd28f9a.
//
// Solidity: function INTERFACE_ID_ON_APPROVE() constant returns(bytes4)
func (_OnApprove *OnApproveCallerSession) INTERFACEIDONAPPROVE() ([4]byte, error) {
	return _OnApprove.Contract.INTERFACEIDONAPPROVE(&_OnApprove.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) constant returns(bool)
func (_OnApprove *OnApproveCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _OnApprove.contract.Call(opts, out, "supportsInterface", interfaceId)
	return *ret0, err
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) constant returns(bool)
func (_OnApprove *OnApproveSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _OnApprove.Contract.SupportsInterface(&_OnApprove.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) constant returns(bool)
func (_OnApprove *OnApproveCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _OnApprove.Contract.SupportsInterface(&_OnApprove.CallOpts, interfaceId)
}

// OnApprove is a paid mutator transaction binding the contract method 0x4273ca16.
//
// Solidity: function onApprove(address owner, address spender, uint256 amount, bytes data) returns(bool)
func (_OnApprove *OnApproveTransactor) OnApprove(opts *bind.TransactOpts, owner common.Address, spender common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _OnApprove.contract.Transact(opts, "onApprove", owner, spender, amount, data)
}

// OnApprove is a paid mutator transaction binding the contract method 0x4273ca16.
//
// Solidity: function onApprove(address owner, address spender, uint256 amount, bytes data) returns(bool)
func (_OnApprove *OnApproveSession) OnApprove(owner common.Address, spender common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _OnApprove.Contract.OnApprove(&_OnApprove.TransactOpts, owner, spender, amount, data)
}

// OnApprove is a paid mutator transaction binding the contract method 0x4273ca16.
//
// Solidity: function onApprove(address owner, address spender, uint256 amount, bytes data) returns(bool)
func (_OnApprove *OnApproveTransactorSession) OnApprove(owner common.Address, spender common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _OnApprove.Contract.OnApprove(&_OnApprove.TransactOpts, owner, spender, amount, data)
}

// OnApproveConstantABI is the input ABI used to generate the binding from.
const OnApproveConstantABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"INTERFACE_ID_ON_APPROVE\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// OnApproveConstantFuncSigs maps the 4-byte function signature to its string representation.
var OnApproveConstantFuncSigs = map[string]string{
	"6cd28f9a": "INTERFACE_ID_ON_APPROVE()",
}

// OnApproveConstantBin is the compiled bytecode used for deploying new contracts.
var OnApproveConstantBin = "0x6080604052348015600f57600080fd5b5060c78061001e6000396000f3fe6080604052348015600f57600080fd5b506004361060285760003560e01c80636cd28f9a14602d575b600080fd5b60336050565b604080516001600160e01b03199092168252519081900360200190f35b604051806028606b8239602801905060405180910390208156fe6f6e417070726f766528616464726573732c616464726573732c75696e743235362c627974657329a265627a7a723158206d94182e39f46cad6425154650e56ecd40259e07a64364bd131e67fe91c85c7664736f6c634300050c0032"

// DeployOnApproveConstant deploys a new Ethereum contract, binding an instance of OnApproveConstant to it.
func DeployOnApproveConstant(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *OnApproveConstant, error) {
	parsed, err := abi.JSON(strings.NewReader(OnApproveConstantABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OnApproveConstantBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OnApproveConstant{OnApproveConstantCaller: OnApproveConstantCaller{contract: contract}, OnApproveConstantTransactor: OnApproveConstantTransactor{contract: contract}, OnApproveConstantFilterer: OnApproveConstantFilterer{contract: contract}}, nil
}

// OnApproveConstant is an auto generated Go binding around an Ethereum contract.
type OnApproveConstant struct {
	OnApproveConstantCaller     // Read-only binding to the contract
	OnApproveConstantTransactor // Write-only binding to the contract
	OnApproveConstantFilterer   // Log filterer for contract events
}

// OnApproveConstantCaller is an auto generated read-only Go binding around an Ethereum contract.
type OnApproveConstantCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OnApproveConstantTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OnApproveConstantTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OnApproveConstantFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OnApproveConstantFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OnApproveConstantSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OnApproveConstantSession struct {
	Contract     *OnApproveConstant // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// OnApproveConstantCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OnApproveConstantCallerSession struct {
	Contract *OnApproveConstantCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// OnApproveConstantTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OnApproveConstantTransactorSession struct {
	Contract     *OnApproveConstantTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// OnApproveConstantRaw is an auto generated low-level Go binding around an Ethereum contract.
type OnApproveConstantRaw struct {
	Contract *OnApproveConstant // Generic contract binding to access the raw methods on
}

// OnApproveConstantCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OnApproveConstantCallerRaw struct {
	Contract *OnApproveConstantCaller // Generic read-only contract binding to access the raw methods on
}

// OnApproveConstantTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OnApproveConstantTransactorRaw struct {
	Contract *OnApproveConstantTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOnApproveConstant creates a new instance of OnApproveConstant, bound to a specific deployed contract.
func NewOnApproveConstant(address common.Address, backend bind.ContractBackend) (*OnApproveConstant, error) {
	contract, err := bindOnApproveConstant(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OnApproveConstant{OnApproveConstantCaller: OnApproveConstantCaller{contract: contract}, OnApproveConstantTransactor: OnApproveConstantTransactor{contract: contract}, OnApproveConstantFilterer: OnApproveConstantFilterer{contract: contract}}, nil
}

// NewOnApproveConstantCaller creates a new read-only instance of OnApproveConstant, bound to a specific deployed contract.
func NewOnApproveConstantCaller(address common.Address, caller bind.ContractCaller) (*OnApproveConstantCaller, error) {
	contract, err := bindOnApproveConstant(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OnApproveConstantCaller{contract: contract}, nil
}

// NewOnApproveConstantTransactor creates a new write-only instance of OnApproveConstant, bound to a specific deployed contract.
func NewOnApproveConstantTransactor(address common.Address, transactor bind.ContractTransactor) (*OnApproveConstantTransactor, error) {
	contract, err := bindOnApproveConstant(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OnApproveConstantTransactor{contract: contract}, nil
}

// NewOnApproveConstantFilterer creates a new log filterer instance of OnApproveConstant, bound to a specific deployed contract.
func NewOnApproveConstantFilterer(address common.Address, filterer bind.ContractFilterer) (*OnApproveConstantFilterer, error) {
	contract, err := bindOnApproveConstant(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OnApproveConstantFilterer{contract: contract}, nil
}

// bindOnApproveConstant binds a generic wrapper to an already deployed contract.
func bindOnApproveConstant(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OnApproveConstantABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OnApproveConstant *OnApproveConstantRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _OnApproveConstant.Contract.OnApproveConstantCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OnApproveConstant *OnApproveConstantRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OnApproveConstant.Contract.OnApproveConstantTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OnApproveConstant *OnApproveConstantRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OnApproveConstant.Contract.OnApproveConstantTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OnApproveConstant *OnApproveConstantCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _OnApproveConstant.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OnApproveConstant *OnApproveConstantTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OnApproveConstant.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OnApproveConstant *OnApproveConstantTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OnApproveConstant.Contract.contract.Transact(opts, method, params...)
}

// INTERFACEIDONAPPROVE is a free data retrieval call binding the contract method 0x6cd28f9a.
//
// Solidity: function INTERFACE_ID_ON_APPROVE() constant returns(bytes4)
func (_OnApproveConstant *OnApproveConstantCaller) INTERFACEIDONAPPROVE(opts *bind.CallOpts) ([4]byte, error) {
	var (
		ret0 = new([4]byte)
	)
	out := ret0
	err := _OnApproveConstant.contract.Call(opts, out, "INTERFACE_ID_ON_APPROVE")
	return *ret0, err
}

// INTERFACEIDONAPPROVE is a free data retrieval call binding the contract method 0x6cd28f9a.
//
// Solidity: function INTERFACE_ID_ON_APPROVE() constant returns(bytes4)
func (_OnApproveConstant *OnApproveConstantSession) INTERFACEIDONAPPROVE() ([4]byte, error) {
	return _OnApproveConstant.Contract.INTERFACEIDONAPPROVE(&_OnApproveConstant.CallOpts)
}

// INTERFACEIDONAPPROVE is a free data retrieval call binding the contract method 0x6cd28f9a.
//
// Solidity: function INTERFACE_ID_ON_APPROVE() constant returns(bytes4)
func (_OnApproveConstant *OnApproveConstantCallerSession) INTERFACEIDONAPPROVE() ([4]byte, error) {
	return _OnApproveConstant.Contract.INTERFACEIDONAPPROVE(&_OnApproveConstant.CallOpts)
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

// OwnableTargetABI is the input ABI used to generate the binding from.
const OwnableTargetABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// OwnableTargetFuncSigs maps the 4-byte function signature to its string representation.
var OwnableTargetFuncSigs = map[string]string{
	"715018a6": "renounceOwnership()",
	"f2fde38b": "transferOwnership(address)",
}

// OwnableTarget is an auto generated Go binding around an Ethereum contract.
type OwnableTarget struct {
	OwnableTargetCaller     // Read-only binding to the contract
	OwnableTargetTransactor // Write-only binding to the contract
	OwnableTargetFilterer   // Log filterer for contract events
}

// OwnableTargetCaller is an auto generated read-only Go binding around an Ethereum contract.
type OwnableTargetCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableTargetTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OwnableTargetTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableTargetFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OwnableTargetFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableTargetSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OwnableTargetSession struct {
	Contract     *OwnableTarget    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwnableTargetCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OwnableTargetCallerSession struct {
	Contract *OwnableTargetCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// OwnableTargetTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OwnableTargetTransactorSession struct {
	Contract     *OwnableTargetTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// OwnableTargetRaw is an auto generated low-level Go binding around an Ethereum contract.
type OwnableTargetRaw struct {
	Contract *OwnableTarget // Generic contract binding to access the raw methods on
}

// OwnableTargetCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OwnableTargetCallerRaw struct {
	Contract *OwnableTargetCaller // Generic read-only contract binding to access the raw methods on
}

// OwnableTargetTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OwnableTargetTransactorRaw struct {
	Contract *OwnableTargetTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOwnableTarget creates a new instance of OwnableTarget, bound to a specific deployed contract.
func NewOwnableTarget(address common.Address, backend bind.ContractBackend) (*OwnableTarget, error) {
	contract, err := bindOwnableTarget(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OwnableTarget{OwnableTargetCaller: OwnableTargetCaller{contract: contract}, OwnableTargetTransactor: OwnableTargetTransactor{contract: contract}, OwnableTargetFilterer: OwnableTargetFilterer{contract: contract}}, nil
}

// NewOwnableTargetCaller creates a new read-only instance of OwnableTarget, bound to a specific deployed contract.
func NewOwnableTargetCaller(address common.Address, caller bind.ContractCaller) (*OwnableTargetCaller, error) {
	contract, err := bindOwnableTarget(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableTargetCaller{contract: contract}, nil
}

// NewOwnableTargetTransactor creates a new write-only instance of OwnableTarget, bound to a specific deployed contract.
func NewOwnableTargetTransactor(address common.Address, transactor bind.ContractTransactor) (*OwnableTargetTransactor, error) {
	contract, err := bindOwnableTarget(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableTargetTransactor{contract: contract}, nil
}

// NewOwnableTargetFilterer creates a new log filterer instance of OwnableTarget, bound to a specific deployed contract.
func NewOwnableTargetFilterer(address common.Address, filterer bind.ContractFilterer) (*OwnableTargetFilterer, error) {
	contract, err := bindOwnableTarget(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OwnableTargetFilterer{contract: contract}, nil
}

// bindOwnableTarget binds a generic wrapper to an already deployed contract.
func bindOwnableTarget(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnableTargetABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OwnableTarget *OwnableTargetRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _OwnableTarget.Contract.OwnableTargetCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OwnableTarget *OwnableTargetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OwnableTarget.Contract.OwnableTargetTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OwnableTarget *OwnableTargetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OwnableTarget.Contract.OwnableTargetTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OwnableTarget *OwnableTargetCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _OwnableTarget.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OwnableTarget *OwnableTargetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OwnableTarget.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OwnableTarget *OwnableTargetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OwnableTarget.Contract.contract.Transact(opts, method, params...)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OwnableTarget *OwnableTargetTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OwnableTarget.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OwnableTarget *OwnableTargetSession) RenounceOwnership() (*types.Transaction, error) {
	return _OwnableTarget.Contract.RenounceOwnership(&_OwnableTarget.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OwnableTarget *OwnableTargetTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _OwnableTarget.Contract.RenounceOwnership(&_OwnableTarget.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OwnableTarget *OwnableTargetTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _OwnableTarget.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OwnableTarget *OwnableTargetSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _OwnableTarget.Contract.TransferOwnership(&_OwnableTarget.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OwnableTarget *OwnableTargetTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _OwnableTarget.Contract.TransferOwnership(&_OwnableTarget.TransactOpts, newOwner)
}

// PauserRoleRenounceTargetABI is the input ABI used to generate the binding from.
const PauserRoleRenounceTargetABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"renouncePauser\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// PauserRoleRenounceTargetFuncSigs maps the 4-byte function signature to its string representation.
var PauserRoleRenounceTargetFuncSigs = map[string]string{
	"6ef8d66d": "renouncePauser()",
}

// PauserRoleRenounceTarget is an auto generated Go binding around an Ethereum contract.
type PauserRoleRenounceTarget struct {
	PauserRoleRenounceTargetCaller     // Read-only binding to the contract
	PauserRoleRenounceTargetTransactor // Write-only binding to the contract
	PauserRoleRenounceTargetFilterer   // Log filterer for contract events
}

// PauserRoleRenounceTargetCaller is an auto generated read-only Go binding around an Ethereum contract.
type PauserRoleRenounceTargetCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PauserRoleRenounceTargetTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PauserRoleRenounceTargetTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PauserRoleRenounceTargetFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PauserRoleRenounceTargetFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PauserRoleRenounceTargetSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PauserRoleRenounceTargetSession struct {
	Contract     *PauserRoleRenounceTarget // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// PauserRoleRenounceTargetCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PauserRoleRenounceTargetCallerSession struct {
	Contract *PauserRoleRenounceTargetCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// PauserRoleRenounceTargetTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PauserRoleRenounceTargetTransactorSession struct {
	Contract     *PauserRoleRenounceTargetTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// PauserRoleRenounceTargetRaw is an auto generated low-level Go binding around an Ethereum contract.
type PauserRoleRenounceTargetRaw struct {
	Contract *PauserRoleRenounceTarget // Generic contract binding to access the raw methods on
}

// PauserRoleRenounceTargetCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PauserRoleRenounceTargetCallerRaw struct {
	Contract *PauserRoleRenounceTargetCaller // Generic read-only contract binding to access the raw methods on
}

// PauserRoleRenounceTargetTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PauserRoleRenounceTargetTransactorRaw struct {
	Contract *PauserRoleRenounceTargetTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPauserRoleRenounceTarget creates a new instance of PauserRoleRenounceTarget, bound to a specific deployed contract.
func NewPauserRoleRenounceTarget(address common.Address, backend bind.ContractBackend) (*PauserRoleRenounceTarget, error) {
	contract, err := bindPauserRoleRenounceTarget(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PauserRoleRenounceTarget{PauserRoleRenounceTargetCaller: PauserRoleRenounceTargetCaller{contract: contract}, PauserRoleRenounceTargetTransactor: PauserRoleRenounceTargetTransactor{contract: contract}, PauserRoleRenounceTargetFilterer: PauserRoleRenounceTargetFilterer{contract: contract}}, nil
}

// NewPauserRoleRenounceTargetCaller creates a new read-only instance of PauserRoleRenounceTarget, bound to a specific deployed contract.
func NewPauserRoleRenounceTargetCaller(address common.Address, caller bind.ContractCaller) (*PauserRoleRenounceTargetCaller, error) {
	contract, err := bindPauserRoleRenounceTarget(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PauserRoleRenounceTargetCaller{contract: contract}, nil
}

// NewPauserRoleRenounceTargetTransactor creates a new write-only instance of PauserRoleRenounceTarget, bound to a specific deployed contract.
func NewPauserRoleRenounceTargetTransactor(address common.Address, transactor bind.ContractTransactor) (*PauserRoleRenounceTargetTransactor, error) {
	contract, err := bindPauserRoleRenounceTarget(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PauserRoleRenounceTargetTransactor{contract: contract}, nil
}

// NewPauserRoleRenounceTargetFilterer creates a new log filterer instance of PauserRoleRenounceTarget, bound to a specific deployed contract.
func NewPauserRoleRenounceTargetFilterer(address common.Address, filterer bind.ContractFilterer) (*PauserRoleRenounceTargetFilterer, error) {
	contract, err := bindPauserRoleRenounceTarget(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PauserRoleRenounceTargetFilterer{contract: contract}, nil
}

// bindPauserRoleRenounceTarget binds a generic wrapper to an already deployed contract.
func bindPauserRoleRenounceTarget(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PauserRoleRenounceTargetABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PauserRoleRenounceTarget *PauserRoleRenounceTargetRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PauserRoleRenounceTarget.Contract.PauserRoleRenounceTargetCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PauserRoleRenounceTarget *PauserRoleRenounceTargetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PauserRoleRenounceTarget.Contract.PauserRoleRenounceTargetTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PauserRoleRenounceTarget *PauserRoleRenounceTargetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PauserRoleRenounceTarget.Contract.PauserRoleRenounceTargetTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PauserRoleRenounceTarget *PauserRoleRenounceTargetCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PauserRoleRenounceTarget.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PauserRoleRenounceTarget *PauserRoleRenounceTargetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PauserRoleRenounceTarget.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PauserRoleRenounceTarget *PauserRoleRenounceTargetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PauserRoleRenounceTarget.Contract.contract.Transact(opts, method, params...)
}

// RenouncePauser is a paid mutator transaction binding the contract method 0x6ef8d66d.
//
// Solidity: function renouncePauser() returns()
func (_PauserRoleRenounceTarget *PauserRoleRenounceTargetTransactor) RenouncePauser(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PauserRoleRenounceTarget.contract.Transact(opts, "renouncePauser")
}

// RenouncePauser is a paid mutator transaction binding the contract method 0x6ef8d66d.
//
// Solidity: function renouncePauser() returns()
func (_PauserRoleRenounceTarget *PauserRoleRenounceTargetSession) RenouncePauser() (*types.Transaction, error) {
	return _PauserRoleRenounceTarget.Contract.RenouncePauser(&_PauserRoleRenounceTarget.TransactOpts)
}

// RenouncePauser is a paid mutator transaction binding the contract method 0x6ef8d66d.
//
// Solidity: function renouncePauser() returns()
func (_PauserRoleRenounceTarget *PauserRoleRenounceTargetTransactorSession) RenouncePauser() (*types.Transaction, error) {
	return _PauserRoleRenounceTarget.Contract.RenouncePauser(&_PauserRoleRenounceTarget.TransactOpts)
}

// ReentrancyGuardABI is the input ABI used to generate the binding from.
const ReentrancyGuardABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// ReentrancyGuard is an auto generated Go binding around an Ethereum contract.
type ReentrancyGuard struct {
	ReentrancyGuardCaller     // Read-only binding to the contract
	ReentrancyGuardTransactor // Write-only binding to the contract
	ReentrancyGuardFilterer   // Log filterer for contract events
}

// ReentrancyGuardCaller is an auto generated read-only Go binding around an Ethereum contract.
type ReentrancyGuardCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReentrancyGuardTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ReentrancyGuardTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReentrancyGuardFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ReentrancyGuardFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReentrancyGuardSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ReentrancyGuardSession struct {
	Contract     *ReentrancyGuard  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ReentrancyGuardCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ReentrancyGuardCallerSession struct {
	Contract *ReentrancyGuardCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// ReentrancyGuardTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ReentrancyGuardTransactorSession struct {
	Contract     *ReentrancyGuardTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ReentrancyGuardRaw is an auto generated low-level Go binding around an Ethereum contract.
type ReentrancyGuardRaw struct {
	Contract *ReentrancyGuard // Generic contract binding to access the raw methods on
}

// ReentrancyGuardCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ReentrancyGuardCallerRaw struct {
	Contract *ReentrancyGuardCaller // Generic read-only contract binding to access the raw methods on
}

// ReentrancyGuardTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ReentrancyGuardTransactorRaw struct {
	Contract *ReentrancyGuardTransactor // Generic write-only contract binding to access the raw methods on
}

// NewReentrancyGuard creates a new instance of ReentrancyGuard, bound to a specific deployed contract.
func NewReentrancyGuard(address common.Address, backend bind.ContractBackend) (*ReentrancyGuard, error) {
	contract, err := bindReentrancyGuard(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ReentrancyGuard{ReentrancyGuardCaller: ReentrancyGuardCaller{contract: contract}, ReentrancyGuardTransactor: ReentrancyGuardTransactor{contract: contract}, ReentrancyGuardFilterer: ReentrancyGuardFilterer{contract: contract}}, nil
}

// NewReentrancyGuardCaller creates a new read-only instance of ReentrancyGuard, bound to a specific deployed contract.
func NewReentrancyGuardCaller(address common.Address, caller bind.ContractCaller) (*ReentrancyGuardCaller, error) {
	contract, err := bindReentrancyGuard(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ReentrancyGuardCaller{contract: contract}, nil
}

// NewReentrancyGuardTransactor creates a new write-only instance of ReentrancyGuard, bound to a specific deployed contract.
func NewReentrancyGuardTransactor(address common.Address, transactor bind.ContractTransactor) (*ReentrancyGuardTransactor, error) {
	contract, err := bindReentrancyGuard(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ReentrancyGuardTransactor{contract: contract}, nil
}

// NewReentrancyGuardFilterer creates a new log filterer instance of ReentrancyGuard, bound to a specific deployed contract.
func NewReentrancyGuardFilterer(address common.Address, filterer bind.ContractFilterer) (*ReentrancyGuardFilterer, error) {
	contract, err := bindReentrancyGuard(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ReentrancyGuardFilterer{contract: contract}, nil
}

// bindReentrancyGuard binds a generic wrapper to an already deployed contract.
func bindReentrancyGuard(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ReentrancyGuardABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReentrancyGuard *ReentrancyGuardRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ReentrancyGuard.Contract.ReentrancyGuardCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReentrancyGuard *ReentrancyGuardRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReentrancyGuard.Contract.ReentrancyGuardTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReentrancyGuard *ReentrancyGuardRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReentrancyGuard.Contract.ReentrancyGuardTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReentrancyGuard *ReentrancyGuardCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ReentrancyGuard.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReentrancyGuard *ReentrancyGuardTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReentrancyGuard.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReentrancyGuard *ReentrancyGuardTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReentrancyGuard.Contract.contract.Transact(opts, method, params...)
}

// RolesABI is the input ABI used to generate the binding from.
const RolesABI = "[]"

// RolesBin is the compiled bytecode used for deploying new contracts.
var RolesBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a72315820451a2112c879141f72bbe0ec00d853511a9d98646af7e3081d90a3d70270eb1564736f6c634300050c0032"

// DeployRoles deploys a new Ethereum contract, binding an instance of Roles to it.
func DeployRoles(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Roles, error) {
	parsed, err := abi.JSON(strings.NewReader(RolesABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RolesBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Roles{RolesCaller: RolesCaller{contract: contract}, RolesTransactor: RolesTransactor{contract: contract}, RolesFilterer: RolesFilterer{contract: contract}}, nil
}

// Roles is an auto generated Go binding around an Ethereum contract.
type Roles struct {
	RolesCaller     // Read-only binding to the contract
	RolesTransactor // Write-only binding to the contract
	RolesFilterer   // Log filterer for contract events
}

// RolesCaller is an auto generated read-only Go binding around an Ethereum contract.
type RolesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RolesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RolesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RolesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RolesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RolesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RolesSession struct {
	Contract     *Roles            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RolesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RolesCallerSession struct {
	Contract *RolesCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// RolesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RolesTransactorSession struct {
	Contract     *RolesTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RolesRaw is an auto generated low-level Go binding around an Ethereum contract.
type RolesRaw struct {
	Contract *Roles // Generic contract binding to access the raw methods on
}

// RolesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RolesCallerRaw struct {
	Contract *RolesCaller // Generic read-only contract binding to access the raw methods on
}

// RolesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RolesTransactorRaw struct {
	Contract *RolesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRoles creates a new instance of Roles, bound to a specific deployed contract.
func NewRoles(address common.Address, backend bind.ContractBackend) (*Roles, error) {
	contract, err := bindRoles(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Roles{RolesCaller: RolesCaller{contract: contract}, RolesTransactor: RolesTransactor{contract: contract}, RolesFilterer: RolesFilterer{contract: contract}}, nil
}

// NewRolesCaller creates a new read-only instance of Roles, bound to a specific deployed contract.
func NewRolesCaller(address common.Address, caller bind.ContractCaller) (*RolesCaller, error) {
	contract, err := bindRoles(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RolesCaller{contract: contract}, nil
}

// NewRolesTransactor creates a new write-only instance of Roles, bound to a specific deployed contract.
func NewRolesTransactor(address common.Address, transactor bind.ContractTransactor) (*RolesTransactor, error) {
	contract, err := bindRoles(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RolesTransactor{contract: contract}, nil
}

// NewRolesFilterer creates a new log filterer instance of Roles, bound to a specific deployed contract.
func NewRolesFilterer(address common.Address, filterer bind.ContractFilterer) (*RolesFilterer, error) {
	contract, err := bindRoles(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RolesFilterer{contract: contract}, nil
}

// bindRoles binds a generic wrapper to an already deployed contract.
func bindRoles(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RolesABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Roles *RolesRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Roles.Contract.RolesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Roles *RolesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Roles.Contract.RolesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Roles *RolesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Roles.Contract.RolesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Roles *RolesCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Roles.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Roles *RolesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Roles.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Roles *RolesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Roles.Contract.contract.Transact(opts, method, params...)
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

// SafeERC20ABI is the input ABI used to generate the binding from.
const SafeERC20ABI = "[]"

// SafeERC20Bin is the compiled bytecode used for deploying new contracts.
var SafeERC20Bin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a7231582045a04cb654e4c37bcd57cef5b1e4313fd04e0fbb06c03b04321592cfa5eb149764736f6c634300050c0032"

// DeploySafeERC20 deploys a new Ethereum contract, binding an instance of SafeERC20 to it.
func DeploySafeERC20(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeERC20, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeERC20ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeERC20Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeERC20{SafeERC20Caller: SafeERC20Caller{contract: contract}, SafeERC20Transactor: SafeERC20Transactor{contract: contract}, SafeERC20Filterer: SafeERC20Filterer{contract: contract}}, nil
}

// SafeERC20 is an auto generated Go binding around an Ethereum contract.
type SafeERC20 struct {
	SafeERC20Caller     // Read-only binding to the contract
	SafeERC20Transactor // Write-only binding to the contract
	SafeERC20Filterer   // Log filterer for contract events
}

// SafeERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type SafeERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeERC20Session struct {
	Contract     *SafeERC20        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeERC20CallerSession struct {
	Contract *SafeERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// SafeERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeERC20TransactorSession struct {
	Contract     *SafeERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// SafeERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type SafeERC20Raw struct {
	Contract *SafeERC20 // Generic contract binding to access the raw methods on
}

// SafeERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeERC20CallerRaw struct {
	Contract *SafeERC20Caller // Generic read-only contract binding to access the raw methods on
}

// SafeERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeERC20TransactorRaw struct {
	Contract *SafeERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeERC20 creates a new instance of SafeERC20, bound to a specific deployed contract.
func NewSafeERC20(address common.Address, backend bind.ContractBackend) (*SafeERC20, error) {
	contract, err := bindSafeERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeERC20{SafeERC20Caller: SafeERC20Caller{contract: contract}, SafeERC20Transactor: SafeERC20Transactor{contract: contract}, SafeERC20Filterer: SafeERC20Filterer{contract: contract}}, nil
}

// NewSafeERC20Caller creates a new read-only instance of SafeERC20, bound to a specific deployed contract.
func NewSafeERC20Caller(address common.Address, caller bind.ContractCaller) (*SafeERC20Caller, error) {
	contract, err := bindSafeERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeERC20Caller{contract: contract}, nil
}

// NewSafeERC20Transactor creates a new write-only instance of SafeERC20, bound to a specific deployed contract.
func NewSafeERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*SafeERC20Transactor, error) {
	contract, err := bindSafeERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeERC20Transactor{contract: contract}, nil
}

// NewSafeERC20Filterer creates a new log filterer instance of SafeERC20, bound to a specific deployed contract.
func NewSafeERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*SafeERC20Filterer, error) {
	contract, err := bindSafeERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeERC20Filterer{contract: contract}, nil
}

// bindSafeERC20 binds a generic wrapper to an already deployed contract.
func bindSafeERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeERC20 *SafeERC20Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeERC20.Contract.SafeERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeERC20 *SafeERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeERC20.Contract.SafeERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeERC20 *SafeERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeERC20.Contract.SafeERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeERC20 *SafeERC20CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeERC20 *SafeERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeERC20 *SafeERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeERC20.Contract.contract.Transact(opts, method, params...)
}

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
var SafeMathBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a7231582045f810ccb41cd151344ccf006ddfbe7c3c58a867d26e345c6576ae09178171dc64736f6c634300050c0032"

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

// SeigTokenABI is the input ABI used to generate the binding from.
const SeigTokenABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"INTERFACE_ID_ON_APPROVE\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"approveAndCall\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"renounceMinter\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"renouncePauser\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"seigManager\",\"outputs\":[{\"internalType\":\"contractSeigManagerI\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractSeigManagerI\",\"name\":\"_seigManager\",\"type\":\"address\"}],\"name\":\"setSeigManager\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// SeigTokenFuncSigs maps the 4-byte function signature to its string representation.
var SeigTokenFuncSigs = map[string]string{
	"6cd28f9a": "INTERFACE_ID_ON_APPROVE()",
	"dd62ed3e": "allowance(address,address)",
	"095ea7b3": "approve(address,uint256)",
	"cae9ca51": "approveAndCall(address,uint256,bytes)",
	"70a08231": "balanceOf(address)",
	"a457c2d7": "decreaseAllowance(address,uint256)",
	"39509351": "increaseAllowance(address,uint256)",
	"8f32d59b": "isOwner()",
	"8da5cb5b": "owner()",
	"5f112c68": "renounceMinter(address)",
	"715018a6": "renounceOwnership()",
	"38bf3cfa": "renounceOwnership(address)",
	"41eb24bb": "renouncePauser(address)",
	"6fb7f558": "seigManager()",
	"7657f20a": "setSeigManager(address)",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
	"f2fde38b": "transferOwnership(address)",
	"6d435421": "transferOwnership(address,address)",
}

// SeigTokenBin is the compiled bytecode used for deploying new contracts.
var SeigTokenBin = "0x608060405260006100176001600160e01b0361006a16565b600380546001600160a01b0319166001600160a01b038316908117909155604051919250906000907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a35061006e565b3390565b6114e68061007d6000396000f3fe608060405234801561001057600080fd5b506004361061012c5760003560e01c806370a08231116100ad578063a457c2d711610071578063a457c2d71461033c578063a9059cbb14610368578063cae9ca5114610394578063dd62ed3e1461044f578063f2fde38b1461047d5761012c565b806370a08231146102d8578063715018a6146102fe5780637657f20a146103065780638da5cb5b1461032c5780638f32d59b146103345761012c565b806341eb24bb116100f457806341eb24bb146102155780635f112c681461023b5780636cd28f9a146102615780636d435421146102865780636fb7f558146102b45761012c565b8063095ea7b31461013157806318160ddd1461017157806323b872dd1461018b57806338bf3cfa146101c157806339509351146101e9575b600080fd5b61015d6004803603604081101561014757600080fd5b506001600160a01b0381351690602001356104a3565b604080519115158252519081900360200190f35b6101796104c0565b60408051918252519081900360200190f35b61015d600480360360608110156101a157600080fd5b506001600160a01b038135811691602081013590911690604001356104c6565b6101e7600480360360208110156101d757600080fd5b50356001600160a01b0316610554565b005b61015d600480360360408110156101ff57600080fd5b506001600160a01b0381351690602001356105f1565b6101e76004803603602081101561022b57600080fd5b50356001600160a01b0316610645565b6101e76004803603602081101561025157600080fd5b50356001600160a01b03166106c7565b610269610749565b604080516001600160e01b03199092168252519081900360200190f35b6101e76004803603604081101561029c57600080fd5b506001600160a01b0381358116916020013516610764565b6102bc61081f565b604080516001600160a01b039092168252519081900360200190f35b610179600480360360208110156102ee57600080fd5b50356001600160a01b031661082e565b6101e7610849565b6101e76004803603602081101561031c57600080fd5b50356001600160a01b03166108da565b6102bc610943565b61015d610952565b61015d6004803603604081101561035257600080fd5b506001600160a01b038135169060200135610978565b61015d6004803603604081101561037e57600080fd5b506001600160a01b0381351690602001356109e6565b61015d600480360360608110156103aa57600080fd5b6001600160a01b03823516916020810135918101906060810160408201356401000000008111156103da57600080fd5b8201836020820111156103ec57600080fd5b8035906020019184600183028401116401000000008311171561040e57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506109fa945050505050565b6101796004803603604081101561046557600080fd5b506001600160a01b0381358116916020013516610a1b565b6101e76004803603602081101561049357600080fd5b50356001600160a01b0316610a46565b60006104b76104b0610a99565b8484610a9d565b50600192915050565b60025490565b60006104d3848484610b89565b610549846104df610a99565b610544856040518060600160405280602881526020016113fc602891396001600160a01b038a1660009081526001602052604081209061051d610a99565b6001600160a01b03168152602081019190915260400160002054919063ffffffff610bf316565b610a9d565b5060015b9392505050565b61055c610952565b61059b576040805162461bcd60e51b81526020600482018190526024820152600080516020611424833981519152604482015290519081900360640190fd5b806001600160a01b031663715018a66040518163ffffffff1660e01b8152600401600060405180830381600087803b1580156105d657600080fd5b505af11580156105ea573d6000803e3d6000fd5b5050505050565b60006104b76105fe610a99565b84610544856001600061060f610a99565b6001600160a01b03908116825260208083019390935260409182016000908120918c16815292529020549063ffffffff610c8a16565b61064d610952565b61068c576040805162461bcd60e51b81526020600482018190526024820152600080516020611424833981519152604482015290519081900360640190fd5b806001600160a01b0316636ef8d66d6040518163ffffffff1660e01b8152600401600060405180830381600087803b1580156105d657600080fd5b6106cf610952565b61070e576040805162461bcd60e51b81526020600482018190526024820152600080516020611424833981519152604482015290519081900360640190fd5b806001600160a01b031663986502756040518163ffffffff1660e01b8152600401600060405180830381600087803b1580156105d657600080fd5b6040518060286113a482396028019050604051809103902081565b61076c610952565b6107ab576040805162461bcd60e51b81526020600482018190526024820152600080516020611424833981519152604482015290519081900360640190fd5b816001600160a01b031663f2fde38b826040518263ffffffff1660e01b815260040180826001600160a01b03166001600160a01b03168152602001915050600060405180830381600087803b15801561080357600080fd5b505af1158015610817573d6000803e3d6000fd5b505050505050565b6004546001600160a01b031681565b6001600160a01b031660009081526020819052604090205490565b610851610952565b610890576040805162461bcd60e51b81526020600482018190526024820152600080516020611424833981519152604482015290519081900360640190fd5b6003546040516000916001600160a01b0316907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3600380546001600160a01b0319169055565b6108e2610952565b610921576040805162461bcd60e51b81526020600482018190526024820152600080516020611424833981519152604482015290519081900360640190fd5b600480546001600160a01b0319166001600160a01b0392909216919091179055565b6003546001600160a01b031690565b6003546000906001600160a01b0316610969610a99565b6001600160a01b031614905090565b60006104b7610985610a99565b846105448560405180606001604052806025815260200161148d60259139600160006109af610a99565b6001600160a01b03908116825260208083019390935260409182016000908120918d1681529252902054919063ffffffff610bf316565b60006104b76109f3610a99565b8484610b89565b6000610a0684846104a3565b610a0f57600080fd5b61054d33858585610ce4565b6001600160a01b03918216600090815260016020908152604080832093909416825291909152205490565b610a4e610952565b610a8d576040805162461bcd60e51b81526020600482018190526024820152600080516020611424833981519152604482015290519081900360640190fd5b610a9681610f3a565b50565b3390565b6001600160a01b038316610ae25760405162461bcd60e51b81526004018080602001828103825260248152602001806114696024913960400191505060405180910390fd5b6001600160a01b038216610b275760405162461bcd60e51b815260040180806020018281038252602281526020018061132b6022913960400191505060405180910390fd5b6001600160a01b03808416600081815260016020908152604080832094871680845294825291829020859055815185815291517f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259281900390910190a3505050565b336001600160a01b0384161480610ba85750336001600160a01b038316145b610be35760405162461bcd60e51b81526004018080602001828103825260308152602001806113cc6030913960400191505060405180910390fd5b610bee838383610fdb565b505050565b60008184841115610c825760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b83811015610c47578181015183820152602001610c2f565b50505050905090810190601f168015610c745780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b505050900390565b60008282018381101561054d576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b610d068360405180806113a46028913960280190506040518091039020611137565b610d415760405162461bcd60e51b815260040180806020018281038252603181526020018061134d6031913960400191505060405180910390fd5b60006060846001600160a01b031660405180806113a460289139602801905060405180910390208787878760405160240180856001600160a01b03166001600160a01b03168152602001846001600160a01b03166001600160a01b0316815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b83811015610de2578181015183820152602001610dca565b50505050905090810190601f168015610e0f5780820380516001836020036101000a031916815260200191505b5060408051601f198184030181529181526020820180516001600160e01b03166001600160e01b0319909a16999099178952518151919890975087965094509250829150849050835b60208310610e775780518252601f199092019160209182019101610e58565b6001836020036101000a0380198251168184511680821785525050505050509050019150506000604051808303816000865af19150503d8060008114610ed9576040519150601f19603f3d011682016040523d82523d6000602084013e610ede565b606091505b5091509150818190610f315760405162461bcd60e51b8152602060048201818152835160248401528351909283926044909101919085019080838360008315610c47578181015183820152602001610c2f565b50505050505050565b6001600160a01b038116610f7f5760405162461bcd60e51b81526004018080602001828103825260268152602001806113056026913960400191505060405180910390fd5b6003546040516001600160a01b038084169216907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a3600380546001600160a01b0319166001600160a01b0392909216919091179055565b6001600160a01b0383166110205760405162461bcd60e51b81526004018080602001828103825260258152602001806114446025913960400191505060405180910390fd5b6001600160a01b0382166110655760405162461bcd60e51b81526004018080602001828103825260238152602001806112e26023913960400191505060405180910390fd5b6110a88160405180606001604052806026815260200161137e602691396001600160a01b038616600090815260208190526040902054919063ffffffff610bf316565b6001600160a01b0380851660009081526020819052604080822093909355908416815220546110dd908263ffffffff610c8a16565b6001600160a01b038084166000818152602081815260409182902094909455805185815290519193928716927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef92918290030190a3505050565b600061114283611153565b801561054d575061054d8383611187565b6000611166826301ffc9a760e01b611187565b8015611181575061117f826001600160e01b0319611187565b155b92915050565b600080600061119685856111ad565b915091508180156111a45750805b95945050505050565b604080516001600160e01b0319831660248083019190915282518083039091018152604490910182526020810180516001600160e01b03166301ffc9a760e01b1781529151815160009384939284926060926001600160a01b038a169261753092879282918083835b602083106112355780518252601f199092019160209182019101611216565b6001836020036101000a0380198251168184511680821785525050505050509050019150506000604051808303818686fa925050503d8060008114611296576040519150601f19603f3d011682016040523d82523d6000602084013e61129b565b606091505b50915091506020815110156112b957600080945094505050506112da565b818180602001905160208110156112cf57600080fd5b505190955093505050505b925092905056fe45524332303a207472616e7366657220746f20746865207a65726f20616464726573734f776e61626c653a206e6577206f776e657220697320746865207a65726f206164647265737345524332303a20617070726f766520746f20746865207a65726f206164647265737345524332304f6e417070726f76653a207370656e64657220646f65736e277420737570706f7274206f6e417070726f766545524332303a207472616e7366657220616d6f756e7420657863656564732062616c616e63656f6e417070726f766528616464726573732c616464726573732c75696e743235362c62797465732953656967546f6b656e3a206f6e6c792073656e646572206f7220726563697069656e742063616e207472616e7366657245524332303a207472616e7366657220616d6f756e74206578636565647320616c6c6f77616e63654f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657245524332303a207472616e736665722066726f6d20746865207a65726f206164647265737345524332303a20617070726f76652066726f6d20746865207a65726f206164647265737345524332303a2064656372656173656420616c6c6f77616e63652062656c6f77207a65726fa265627a7a7231582076279835faafec28aae7facfa0911aaea8f78f2ebf1707778dba15d0dbc0713964736f6c634300050c0032"

// DeploySeigToken deploys a new Ethereum contract, binding an instance of SeigToken to it.
func DeploySeigToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SeigToken, error) {
	parsed, err := abi.JSON(strings.NewReader(SeigTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SeigTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SeigToken{SeigTokenCaller: SeigTokenCaller{contract: contract}, SeigTokenTransactor: SeigTokenTransactor{contract: contract}, SeigTokenFilterer: SeigTokenFilterer{contract: contract}}, nil
}

// SeigToken is an auto generated Go binding around an Ethereum contract.
type SeigToken struct {
	SeigTokenCaller     // Read-only binding to the contract
	SeigTokenTransactor // Write-only binding to the contract
	SeigTokenFilterer   // Log filterer for contract events
}

// SeigTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type SeigTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SeigTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SeigTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SeigTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SeigTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SeigTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SeigTokenSession struct {
	Contract     *SeigToken        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SeigTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SeigTokenCallerSession struct {
	Contract *SeigTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// SeigTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SeigTokenTransactorSession struct {
	Contract     *SeigTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// SeigTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type SeigTokenRaw struct {
	Contract *SeigToken // Generic contract binding to access the raw methods on
}

// SeigTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SeigTokenCallerRaw struct {
	Contract *SeigTokenCaller // Generic read-only contract binding to access the raw methods on
}

// SeigTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SeigTokenTransactorRaw struct {
	Contract *SeigTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSeigToken creates a new instance of SeigToken, bound to a specific deployed contract.
func NewSeigToken(address common.Address, backend bind.ContractBackend) (*SeigToken, error) {
	contract, err := bindSeigToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SeigToken{SeigTokenCaller: SeigTokenCaller{contract: contract}, SeigTokenTransactor: SeigTokenTransactor{contract: contract}, SeigTokenFilterer: SeigTokenFilterer{contract: contract}}, nil
}

// NewSeigTokenCaller creates a new read-only instance of SeigToken, bound to a specific deployed contract.
func NewSeigTokenCaller(address common.Address, caller bind.ContractCaller) (*SeigTokenCaller, error) {
	contract, err := bindSeigToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SeigTokenCaller{contract: contract}, nil
}

// NewSeigTokenTransactor creates a new write-only instance of SeigToken, bound to a specific deployed contract.
func NewSeigTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*SeigTokenTransactor, error) {
	contract, err := bindSeigToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SeigTokenTransactor{contract: contract}, nil
}

// NewSeigTokenFilterer creates a new log filterer instance of SeigToken, bound to a specific deployed contract.
func NewSeigTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*SeigTokenFilterer, error) {
	contract, err := bindSeigToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SeigTokenFilterer{contract: contract}, nil
}

// bindSeigToken binds a generic wrapper to an already deployed contract.
func bindSeigToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SeigTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SeigToken *SeigTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SeigToken.Contract.SeigTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SeigToken *SeigTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SeigToken.Contract.SeigTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SeigToken *SeigTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SeigToken.Contract.SeigTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SeigToken *SeigTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SeigToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SeigToken *SeigTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SeigToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SeigToken *SeigTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SeigToken.Contract.contract.Transact(opts, method, params...)
}

// INTERFACEIDONAPPROVE is a free data retrieval call binding the contract method 0x6cd28f9a.
//
// Solidity: function INTERFACE_ID_ON_APPROVE() constant returns(bytes4)
func (_SeigToken *SeigTokenCaller) INTERFACEIDONAPPROVE(opts *bind.CallOpts) ([4]byte, error) {
	var (
		ret0 = new([4]byte)
	)
	out := ret0
	err := _SeigToken.contract.Call(opts, out, "INTERFACE_ID_ON_APPROVE")
	return *ret0, err
}

// INTERFACEIDONAPPROVE is a free data retrieval call binding the contract method 0x6cd28f9a.
//
// Solidity: function INTERFACE_ID_ON_APPROVE() constant returns(bytes4)
func (_SeigToken *SeigTokenSession) INTERFACEIDONAPPROVE() ([4]byte, error) {
	return _SeigToken.Contract.INTERFACEIDONAPPROVE(&_SeigToken.CallOpts)
}

// INTERFACEIDONAPPROVE is a free data retrieval call binding the contract method 0x6cd28f9a.
//
// Solidity: function INTERFACE_ID_ON_APPROVE() constant returns(bytes4)
func (_SeigToken *SeigTokenCallerSession) INTERFACEIDONAPPROVE() ([4]byte, error) {
	return _SeigToken.Contract.INTERFACEIDONAPPROVE(&_SeigToken.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_SeigToken *SeigTokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SeigToken.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_SeigToken *SeigTokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _SeigToken.Contract.Allowance(&_SeigToken.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_SeigToken *SeigTokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _SeigToken.Contract.Allowance(&_SeigToken.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_SeigToken *SeigTokenCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SeigToken.contract.Call(opts, out, "balanceOf", account)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_SeigToken *SeigTokenSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _SeigToken.Contract.BalanceOf(&_SeigToken.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_SeigToken *SeigTokenCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _SeigToken.Contract.BalanceOf(&_SeigToken.CallOpts, account)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_SeigToken *SeigTokenCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _SeigToken.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_SeigToken *SeigTokenSession) IsOwner() (bool, error) {
	return _SeigToken.Contract.IsOwner(&_SeigToken.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_SeigToken *SeigTokenCallerSession) IsOwner() (bool, error) {
	return _SeigToken.Contract.IsOwner(&_SeigToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_SeigToken *SeigTokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SeigToken.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_SeigToken *SeigTokenSession) Owner() (common.Address, error) {
	return _SeigToken.Contract.Owner(&_SeigToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_SeigToken *SeigTokenCallerSession) Owner() (common.Address, error) {
	return _SeigToken.Contract.Owner(&_SeigToken.CallOpts)
}

// SeigManager is a free data retrieval call binding the contract method 0x6fb7f558.
//
// Solidity: function seigManager() constant returns(address)
func (_SeigToken *SeigTokenCaller) SeigManager(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SeigToken.contract.Call(opts, out, "seigManager")
	return *ret0, err
}

// SeigManager is a free data retrieval call binding the contract method 0x6fb7f558.
//
// Solidity: function seigManager() constant returns(address)
func (_SeigToken *SeigTokenSession) SeigManager() (common.Address, error) {
	return _SeigToken.Contract.SeigManager(&_SeigToken.CallOpts)
}

// SeigManager is a free data retrieval call binding the contract method 0x6fb7f558.
//
// Solidity: function seigManager() constant returns(address)
func (_SeigToken *SeigTokenCallerSession) SeigManager() (common.Address, error) {
	return _SeigToken.Contract.SeigManager(&_SeigToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_SeigToken *SeigTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SeigToken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_SeigToken *SeigTokenSession) TotalSupply() (*big.Int, error) {
	return _SeigToken.Contract.TotalSupply(&_SeigToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_SeigToken *SeigTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _SeigToken.Contract.TotalSupply(&_SeigToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_SeigToken *SeigTokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SeigToken.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_SeigToken *SeigTokenSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SeigToken.Contract.Approve(&_SeigToken.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_SeigToken *SeigTokenTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SeigToken.Contract.Approve(&_SeigToken.TransactOpts, spender, amount)
}

// ApproveAndCall is a paid mutator transaction binding the contract method 0xcae9ca51.
//
// Solidity: function approveAndCall(address spender, uint256 amount, bytes data) returns(bool)
func (_SeigToken *SeigTokenTransactor) ApproveAndCall(opts *bind.TransactOpts, spender common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _SeigToken.contract.Transact(opts, "approveAndCall", spender, amount, data)
}

// ApproveAndCall is a paid mutator transaction binding the contract method 0xcae9ca51.
//
// Solidity: function approveAndCall(address spender, uint256 amount, bytes data) returns(bool)
func (_SeigToken *SeigTokenSession) ApproveAndCall(spender common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _SeigToken.Contract.ApproveAndCall(&_SeigToken.TransactOpts, spender, amount, data)
}

// ApproveAndCall is a paid mutator transaction binding the contract method 0xcae9ca51.
//
// Solidity: function approveAndCall(address spender, uint256 amount, bytes data) returns(bool)
func (_SeigToken *SeigTokenTransactorSession) ApproveAndCall(spender common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _SeigToken.Contract.ApproveAndCall(&_SeigToken.TransactOpts, spender, amount, data)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_SeigToken *SeigTokenTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _SeigToken.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_SeigToken *SeigTokenSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _SeigToken.Contract.DecreaseAllowance(&_SeigToken.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_SeigToken *SeigTokenTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _SeigToken.Contract.DecreaseAllowance(&_SeigToken.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_SeigToken *SeigTokenTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _SeigToken.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_SeigToken *SeigTokenSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _SeigToken.Contract.IncreaseAllowance(&_SeigToken.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_SeigToken *SeigTokenTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _SeigToken.Contract.IncreaseAllowance(&_SeigToken.TransactOpts, spender, addedValue)
}

// RenounceMinter is a paid mutator transaction binding the contract method 0x5f112c68.
//
// Solidity: function renounceMinter(address target) returns()
func (_SeigToken *SeigTokenTransactor) RenounceMinter(opts *bind.TransactOpts, target common.Address) (*types.Transaction, error) {
	return _SeigToken.contract.Transact(opts, "renounceMinter", target)
}

// RenounceMinter is a paid mutator transaction binding the contract method 0x5f112c68.
//
// Solidity: function renounceMinter(address target) returns()
func (_SeigToken *SeigTokenSession) RenounceMinter(target common.Address) (*types.Transaction, error) {
	return _SeigToken.Contract.RenounceMinter(&_SeigToken.TransactOpts, target)
}

// RenounceMinter is a paid mutator transaction binding the contract method 0x5f112c68.
//
// Solidity: function renounceMinter(address target) returns()
func (_SeigToken *SeigTokenTransactorSession) RenounceMinter(target common.Address) (*types.Transaction, error) {
	return _SeigToken.Contract.RenounceMinter(&_SeigToken.TransactOpts, target)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x38bf3cfa.
//
// Solidity: function renounceOwnership(address target) returns()
func (_SeigToken *SeigTokenTransactor) RenounceOwnership(opts *bind.TransactOpts, target common.Address) (*types.Transaction, error) {
	return _SeigToken.contract.Transact(opts, "renounceOwnership", target)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x38bf3cfa.
//
// Solidity: function renounceOwnership(address target) returns()
func (_SeigToken *SeigTokenSession) RenounceOwnership(target common.Address) (*types.Transaction, error) {
	return _SeigToken.Contract.RenounceOwnership(&_SeigToken.TransactOpts, target)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x38bf3cfa.
//
// Solidity: function renounceOwnership(address target) returns()
func (_SeigToken *SeigTokenTransactorSession) RenounceOwnership(target common.Address) (*types.Transaction, error) {
	return _SeigToken.Contract.RenounceOwnership(&_SeigToken.TransactOpts, target)
}

// RenounceOwnership0 is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SeigToken *SeigTokenTransactor) RenounceOwnership0(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SeigToken.contract.Transact(opts, "renounceOwnership0")
}

// RenounceOwnership0 is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SeigToken *SeigTokenSession) RenounceOwnership0() (*types.Transaction, error) {
	return _SeigToken.Contract.RenounceOwnership0(&_SeigToken.TransactOpts)
}

// RenounceOwnership0 is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SeigToken *SeigTokenTransactorSession) RenounceOwnership0() (*types.Transaction, error) {
	return _SeigToken.Contract.RenounceOwnership0(&_SeigToken.TransactOpts)
}

// RenouncePauser is a paid mutator transaction binding the contract method 0x41eb24bb.
//
// Solidity: function renouncePauser(address target) returns()
func (_SeigToken *SeigTokenTransactor) RenouncePauser(opts *bind.TransactOpts, target common.Address) (*types.Transaction, error) {
	return _SeigToken.contract.Transact(opts, "renouncePauser", target)
}

// RenouncePauser is a paid mutator transaction binding the contract method 0x41eb24bb.
//
// Solidity: function renouncePauser(address target) returns()
func (_SeigToken *SeigTokenSession) RenouncePauser(target common.Address) (*types.Transaction, error) {
	return _SeigToken.Contract.RenouncePauser(&_SeigToken.TransactOpts, target)
}

// RenouncePauser is a paid mutator transaction binding the contract method 0x41eb24bb.
//
// Solidity: function renouncePauser(address target) returns()
func (_SeigToken *SeigTokenTransactorSession) RenouncePauser(target common.Address) (*types.Transaction, error) {
	return _SeigToken.Contract.RenouncePauser(&_SeigToken.TransactOpts, target)
}

// SetSeigManager is a paid mutator transaction binding the contract method 0x7657f20a.
//
// Solidity: function setSeigManager(address _seigManager) returns()
func (_SeigToken *SeigTokenTransactor) SetSeigManager(opts *bind.TransactOpts, _seigManager common.Address) (*types.Transaction, error) {
	return _SeigToken.contract.Transact(opts, "setSeigManager", _seigManager)
}

// SetSeigManager is a paid mutator transaction binding the contract method 0x7657f20a.
//
// Solidity: function setSeigManager(address _seigManager) returns()
func (_SeigToken *SeigTokenSession) SetSeigManager(_seigManager common.Address) (*types.Transaction, error) {
	return _SeigToken.Contract.SetSeigManager(&_SeigToken.TransactOpts, _seigManager)
}

// SetSeigManager is a paid mutator transaction binding the contract method 0x7657f20a.
//
// Solidity: function setSeigManager(address _seigManager) returns()
func (_SeigToken *SeigTokenTransactorSession) SetSeigManager(_seigManager common.Address) (*types.Transaction, error) {
	return _SeigToken.Contract.SetSeigManager(&_SeigToken.TransactOpts, _seigManager)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_SeigToken *SeigTokenTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SeigToken.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_SeigToken *SeigTokenSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SeigToken.Contract.Transfer(&_SeigToken.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_SeigToken *SeigTokenTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SeigToken.Contract.Transfer(&_SeigToken.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_SeigToken *SeigTokenTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SeigToken.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_SeigToken *SeigTokenSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SeigToken.Contract.TransferFrom(&_SeigToken.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_SeigToken *SeigTokenTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SeigToken.Contract.TransferFrom(&_SeigToken.TransactOpts, sender, recipient, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0x6d435421.
//
// Solidity: function transferOwnership(address target, address newOwner) returns()
func (_SeigToken *SeigTokenTransactor) TransferOwnership(opts *bind.TransactOpts, target common.Address, newOwner common.Address) (*types.Transaction, error) {
	return _SeigToken.contract.Transact(opts, "transferOwnership", target, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0x6d435421.
//
// Solidity: function transferOwnership(address target, address newOwner) returns()
func (_SeigToken *SeigTokenSession) TransferOwnership(target common.Address, newOwner common.Address) (*types.Transaction, error) {
	return _SeigToken.Contract.TransferOwnership(&_SeigToken.TransactOpts, target, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0x6d435421.
//
// Solidity: function transferOwnership(address target, address newOwner) returns()
func (_SeigToken *SeigTokenTransactorSession) TransferOwnership(target common.Address, newOwner common.Address) (*types.Transaction, error) {
	return _SeigToken.Contract.TransferOwnership(&_SeigToken.TransactOpts, target, newOwner)
}

// TransferOwnership0 is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SeigToken *SeigTokenTransactor) TransferOwnership0(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _SeigToken.contract.Transact(opts, "transferOwnership0", newOwner)
}

// TransferOwnership0 is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SeigToken *SeigTokenSession) TransferOwnership0(newOwner common.Address) (*types.Transaction, error) {
	return _SeigToken.Contract.TransferOwnership0(&_SeigToken.TransactOpts, newOwner)
}

// TransferOwnership0 is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SeigToken *SeigTokenTransactorSession) TransferOwnership0(newOwner common.Address) (*types.Transaction, error) {
	return _SeigToken.Contract.TransferOwnership0(&_SeigToken.TransactOpts, newOwner)
}

// SeigTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the SeigToken contract.
type SeigTokenApprovalIterator struct {
	Event *SeigTokenApproval // Event containing the contract specifics and raw log

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
func (it *SeigTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SeigTokenApproval)
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
		it.Event = new(SeigTokenApproval)
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
func (it *SeigTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SeigTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SeigTokenApproval represents a Approval event raised by the SeigToken contract.
type SeigTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_SeigToken *SeigTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*SeigTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _SeigToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &SeigTokenApprovalIterator{contract: _SeigToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_SeigToken *SeigTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *SeigTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _SeigToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SeigTokenApproval)
				if err := _SeigToken.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_SeigToken *SeigTokenFilterer) ParseApproval(log types.Log) (*SeigTokenApproval, error) {
	event := new(SeigTokenApproval)
	if err := _SeigToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SeigTokenOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the SeigToken contract.
type SeigTokenOwnershipTransferredIterator struct {
	Event *SeigTokenOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *SeigTokenOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SeigTokenOwnershipTransferred)
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
		it.Event = new(SeigTokenOwnershipTransferred)
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
func (it *SeigTokenOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SeigTokenOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SeigTokenOwnershipTransferred represents a OwnershipTransferred event raised by the SeigToken contract.
type SeigTokenOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SeigToken *SeigTokenFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SeigTokenOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SeigToken.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SeigTokenOwnershipTransferredIterator{contract: _SeigToken.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SeigToken *SeigTokenFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SeigTokenOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SeigToken.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SeigTokenOwnershipTransferred)
				if err := _SeigToken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_SeigToken *SeigTokenFilterer) ParseOwnershipTransferred(log types.Log) (*SeigTokenOwnershipTransferred, error) {
	event := new(SeigTokenOwnershipTransferred)
	if err := _SeigToken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SeigTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the SeigToken contract.
type SeigTokenTransferIterator struct {
	Event *SeigTokenTransfer // Event containing the contract specifics and raw log

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
func (it *SeigTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SeigTokenTransfer)
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
		it.Event = new(SeigTokenTransfer)
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
func (it *SeigTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SeigTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SeigTokenTransfer represents a Transfer event raised by the SeigToken contract.
type SeigTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_SeigToken *SeigTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*SeigTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SeigToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &SeigTokenTransferIterator{contract: _SeigToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_SeigToken *SeigTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *SeigTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SeigToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SeigTokenTransfer)
				if err := _SeigToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_SeigToken *SeigTokenFilterer) ParseTransfer(log types.Log) (*SeigTokenTransfer, error) {
	event := new(SeigTokenTransfer)
	if err := _SeigToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WTONABI is the input ABI used to generate the binding from.
const WTONABI = "[{\"inputs\":[{\"internalType\":\"contractERC20Mintable\",\"name\":\"_ton\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"MinterAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"MinterRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"INTERFACE_ID_ON_APPROVE\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addMinter\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"approveAndCall\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burnFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isMinter\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tonAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onApprove\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"renounceMinter\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceMinter\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"renouncePauser\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"seigManager\",\"outputs\":[{\"internalType\":\"contractSeigManagerI\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractSeigManagerI\",\"name\":\"_seigManager\",\"type\":\"address\"}],\"name\":\"setSeigManager\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tonAmount\",\"type\":\"uint256\"}],\"name\":\"swapFromTON\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tonAmount\",\"type\":\"uint256\"}],\"name\":\"swapFromTONAndTransfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"wtonAmount\",\"type\":\"uint256\"}],\"name\":\"swapToTON\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"wtonAmount\",\"type\":\"uint256\"}],\"name\":\"swapToTONAndTransfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ton\",\"outputs\":[{\"internalType\":\"contractERC20Mintable\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// WTONFuncSigs maps the 4-byte function signature to its string representation.
var WTONFuncSigs = map[string]string{
	"6cd28f9a": "INTERFACE_ID_ON_APPROVE()",
	"983b2d56": "addMinter(address)",
	"dd62ed3e": "allowance(address,address)",
	"095ea7b3": "approve(address,uint256)",
	"cae9ca51": "approveAndCall(address,uint256,bytes)",
	"70a08231": "balanceOf(address)",
	"42966c68": "burn(uint256)",
	"79cc6790": "burnFrom(address,uint256)",
	"313ce567": "decimals()",
	"a457c2d7": "decreaseAllowance(address,uint256)",
	"39509351": "increaseAllowance(address,uint256)",
	"aa271e1a": "isMinter(address)",
	"8f32d59b": "isOwner()",
	"40c10f19": "mint(address,uint256)",
	"06fdde03": "name()",
	"4273ca16": "onApprove(address,address,uint256,bytes)",
	"8da5cb5b": "owner()",
	"98650275": "renounceMinter()",
	"5f112c68": "renounceMinter(address)",
	"715018a6": "renounceOwnership()",
	"38bf3cfa": "renounceOwnership(address)",
	"41eb24bb": "renouncePauser(address)",
	"6fb7f558": "seigManager()",
	"7657f20a": "setSeigManager(address)",
	"01ffc9a7": "supportsInterface(bytes4)",
	"e34869d7": "swapFromTON(uint256)",
	"588420b7": "swapFromTONAndTransfer(address,uint256)",
	"f53fe70f": "swapToTON(uint256)",
	"e3b99e85": "swapToTONAndTransfer(address,uint256)",
	"95d89b41": "symbol()",
	"cc48b947": "ton()",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
	"f2fde38b": "transferOwnership(address)",
	"6d435421": "transferOwnership(address,address)",
}

// WTONBin is the compiled bytecode used for deploying new contracts.
var WTONBin = "0x60806040523480156200001157600080fd5b5060405162002dd338038062002dd3833981810160405260208110156200003757600080fd5b5051604080518082018252600b81527f5772617070656420544f4e0000000000000000000000000000000000000000006020828101919091528251808401909352600483527f57544f4e00000000000000000000000000000000000000000000000000000000908301526000805460ff1916600117905590601b620000d7620000c86001600160e01b03620002d616565b6001600160e01b03620002db16565b6000620000ec6001600160e01b03620002d616565b600580546001600160a01b0319166001600160a01b038316908117909155604051919250906000907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a35082516200014f90600690602086019062000526565b5081516200016590600790602085019062000526565b506008805460ff191660ff9290921691909117905550620001a890507f01ffc9a7000000000000000000000000000000000000000000000000000000006200032d565b620001d6604051808062002d896028913960405190819003602801902090506001600160e01b036200032d16565b806001600160a01b031663313ce5676040518163ffffffff1660e01b815260040160206040518083038186803b1580156200021057600080fd5b505afa15801562000225573d6000803e3d6000fd5b505050506040513d60208110156200023c57600080fd5b505160ff16601214620002b057604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f57544f4e3a20646563696d616c73206f6620544f4e206d757374206265203138604482015290519081900360640190fd5b600a80546001600160a01b0319166001600160a01b0392909216919091179055620005c8565b335b90565b620002f6816004620003fc60201b62001f5f1790919060201c565b6040516001600160a01b038216907f6ae172837ea30b801fbfcdd4108aa1d5bf8ff775444fd70256b44e6bf3dfc3f690600090a250565b7fffffffff000000000000000000000000000000000000000000000000000000008082161415620003bf57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601c60248201527f4552433136353a20696e76616c696420696e7465726661636520696400000000604482015290519081900360640190fd5b7fffffffff00000000000000000000000000000000000000000000000000000000166000908152600960205260409020805460ff19166001179055565b6200041182826001600160e01b03620004a316565b156200047e57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f526f6c65733a206163636f756e7420616c72656164792068617320726f6c6500604482015290519081900360640190fd5b6001600160a01b0316600090815260209190915260409020805460ff19166001179055565b60006001600160a01b03821662000506576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602281526020018062002db16022913960400191505060405180910390fd5b506001600160a01b03166000908152602091909152604090205460ff1690565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200056957805160ff191683800117855562000599565b8280016001018555821562000599579182015b82811115620005995782518255916020019190600101906200057c565b50620005a7929150620005ab565b5090565b620002d891905b80821115620005a75760008155600101620005b2565b6127b180620005d86000396000f3fe608060405234801561001057600080fd5b506004361061021c5760003560e01c8063715018a611610125578063a9059cbb116100ad578063dd62ed3e1161007c578063dd62ed3e14610796578063e34869d7146107c4578063e3b99e85146107e1578063f2fde38b1461080d578063f53fe70f146108335761021c565b8063a9059cbb14610681578063aa271e1a146106ad578063cae9ca51146106d3578063cc48b9471461078e5761021c565b80638f32d59b116100f45780638f32d59b1461061757806395d89b411461061f578063983b2d5614610627578063986502751461064d578063a457c2d7146106555761021c565b8063715018a6146105b55780637657f20a146105bd57806379cc6790146105e35780638da5cb5b1461060f5761021c565b806341eb24bb116101a85780635f112c68116101775780635f112c68146104f25780636cd28f9a146105185780636d4354211461053d5780636fb7f5581461056b57806370a082311461058f5761021c565b806341eb24bb146103f35780634273ca161461041957806342966c68146104a9578063588420b7146104c65761021c565b806323b872dd116101ef57806323b872dd1461031f578063313ce5671461035557806338bf3cfa14610373578063395093511461039b57806340c10f19146103c75761021c565b806301ffc9a71461022157806306fdde031461025c578063095ea7b3146102d957806318160ddd14610305575b600080fd5b6102486004803603602081101561023757600080fd5b50356001600160e01b031916610850565b604080519115158252519081900360200190f35b61026461086f565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561029e578181015183820152602001610286565b50505050905090810190601f1680156102cb5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b610248600480360360408110156102ef57600080fd5b506001600160a01b038135169060200135610905565b61030d610922565b60408051918252519081900360200190f35b6102486004803603606081101561033557600080fd5b506001600160a01b03813581169160208101359091169060400135610928565b61035d6109b6565b6040805160ff9092168252519081900360200190f35b6103996004803603602081101561038957600080fd5b50356001600160a01b03166109bf565b005b610248600480360360408110156103b157600080fd5b506001600160a01b038135169060200135610a5c565b610248600480360360408110156103dd57600080fd5b506001600160a01b038135169060200135610ab0565b6103996004803603602081101561040957600080fd5b50356001600160a01b0316610b07565b6102486004803603608081101561042f57600080fd5b6001600160a01b0382358116926020810135909116916040820135919081019060808101606082013564010000000081111561046a57600080fd5b82018360208201111561047c57600080fd5b8035906020019184600183028401116401000000008311171561049e57600080fd5b509092509050610b89565b610399600480360360208110156104bf57600080fd5b5035610c68565b610248600480360360408110156104dc57600080fd5b506001600160a01b038135169060200135610c7c565b6103996004803603602081101561050857600080fd5b50356001600160a01b0316610cec565b610520610d6e565b604080516001600160e01b03199092168252519081900360200190f35b6103996004803603604081101561055357600080fd5b506001600160a01b0381358116916020013516610d89565b610573610e44565b604080516001600160a01b039092168252519081900360200190f35b61030d600480360360208110156105a557600080fd5b50356001600160a01b0316610e58565b610399610e73565b610399600480360360208110156105d357600080fd5b50356001600160a01b0316610f04565b610399600480360360408110156105f957600080fd5b506001600160a01b038135169060200135610f73565b610573610f9e565b610248610fad565b610264610fd3565b6103996004803603602081101561063d57600080fd5b50356001600160a01b0316611034565b610399611083565b6102486004803603604081101561066b57600080fd5b506001600160a01b038135169060200135611095565b6102486004803603604081101561069757600080fd5b506001600160a01b038135169060200135611103565b610248600480360360208110156106c357600080fd5b50356001600160a01b0316611117565b610248600480360360608110156106e957600080fd5b6001600160a01b038235169160208101359181019060608101604082013564010000000081111561071957600080fd5b82018360208201111561072b57600080fd5b8035906020019184600183028401116401000000008311171561074d57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550611130945050505050565b610573611151565b61030d600480360360408110156107ac57600080fd5b506001600160a01b0381358116916020013516611160565b610248600480360360208110156107da57600080fd5b503561118b565b610248600480360360408110156107f757600080fd5b506001600160a01b0381351690602001356111fa565b6103996004803603602081101561082357600080fd5b50356001600160a01b0316611255565b6102486004803603602081101561084957600080fd5b50356112a5565b6001600160e01b03191660009081526009602052604090205460ff1690565b60068054604080516020601f60026000196101006001881615020190951694909404938401819004810282018101909252828152606093909290918301828280156108fb5780601f106108d0576101008083540402835291602001916108fb565b820191906000526020600020905b8154815290600101906020018083116108de57829003601f168201915b5050505050905090565b6000610919610912611300565b8484611304565b50600192915050565b60035490565b60006109358484846113f0565b6109ab84610941611300565b6109a685604051806060016040528060288152602001612610602891396001600160a01b038a1660009081526002602052604081209061097f611300565b6001600160a01b03168152602081019190915260400160002054919063ffffffff61145a16565b611304565b5060015b9392505050565b60085460ff1690565b6109c7610fad565b610a06576040805162461bcd60e51b81526020600482018190526024820152600080516020612638833981519152604482015290519081900360640190fd5b806001600160a01b031663715018a66040518163ffffffff1660e01b8152600401600060405180830381600087803b158015610a4157600080fd5b505af1158015610a55573d6000803e3d6000fd5b5050505050565b6000610919610a69611300565b846109a68560026000610a7a611300565b6001600160a01b03908116825260208083019390935260409182016000908120918c16815292529020549063ffffffff6114f116565b6000610ac2610abd611300565b611117565b610afd5760405162461bcd60e51b81526004018080602001828103825260308152602001806125bf6030913960400191505060405180910390fd5b610919838361154b565b610b0f610fad565b610b4e576040805162461bcd60e51b81526020600482018190526024820152600080516020612638833981519152604482015290519081900360640190fd5b806001600160a01b0316636ef8d66d6040518163ffffffff1660e01b8152600401600060405180830381600087803b158015610a4157600080fd5b600a546000906001600160a01b03163314610bd55760405162461bcd60e51b81526004018080602001828103825260268152602001806127086026913960400191505060405180910390fd5b610be0868786611555565b506000610bec85611587565b9050600080610c3086868080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061159092505050565b91509150610c3f898385611304565b6060610c4a826115b2565b9050610c588a8486846115dc565b5060019998505050505050505050565b610c79610c73611300565b82611832565b50565b6000805460ff16610cc2576040805162461bcd60e51b815260206004820152601f6024820152600080516020612486833981519152604482015290519081900360640190fd5b6000805460ff19169055610cd7338484611555565b90506000805460ff1916600117905592915050565b610cf4610fad565b610d33576040805162461bcd60e51b81526020600482018190526024820152600080516020612638833981519152604482015290519081900360640190fd5b806001600160a01b031663986502756040518163ffffffff1660e01b8152600401600060405180830381600087803b158015610a4157600080fd5b60405180602861256782396028019050604051809103902081565b610d91610fad565b610dd0576040805162461bcd60e51b81526020600482018190526024820152600080516020612638833981519152604482015290519081900360640190fd5b816001600160a01b031663f2fde38b826040518263ffffffff1660e01b815260040180826001600160a01b03166001600160a01b03168152602001915050600060405180830381600087803b158015610e2857600080fd5b505af1158015610e3c573d6000803e3d6000fd5b505050505050565b60085461010090046001600160a01b031681565b6001600160a01b031660009081526001602052604090205490565b610e7b610fad565b610eba576040805162461bcd60e51b81526020600482018190526024820152600080516020612638833981519152604482015290519081900360640190fd5b6005546040516000916001600160a01b0316907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3600580546001600160a01b0319169055565b610f0c610fad565b610f4b576040805162461bcd60e51b81526020600482018190526024820152600080516020612638833981519152604482015290519081900360640190fd5b600880546001600160a01b0390921661010002610100600160a81b0319909216919091179055565b610f7c33611117565b15610f9057610f8b8282611832565b610f9a565b610f9a828261183c565b5050565b6005546001600160a01b031690565b6005546000906001600160a01b0316610fc4611300565b6001600160a01b031614905090565b60078054604080516020601f60026000196101006001881615020190951694909404938401819004810282018101909252828152606093909290918301828280156108fb5780601f106108d0576101008083540402835291602001916108fb565b61103f610abd611300565b61107a5760405162461bcd60e51b81526004018080602001828103825260308152602001806125bf6030913960400191505060405180910390fd5b610c7981611846565b61109361108e611300565b61188e565b565b60006109196110a2611300565b846109a68560405180606001604052806025815260200161275860259139600260006110cc611300565b6001600160a01b03908116825260208083019390935260409182016000908120918d1681529252902054919063ffffffff61145a16565b6000610919611110611300565b84846113f0565b600061112a60048363ffffffff6118d616565b92915050565b600061113c8484610905565b61114557600080fd5b6109af338585856115dc565b600a546001600160a01b031681565b6001600160a01b03918216600090815260026020908152604080832093909416825291909152205490565b6000805460ff166111d1576040805162461bcd60e51b815260206004820152601f6024820152600080516020612486833981519152604482015290519081900360640190fd5b6000805460ff191690556111e6338084611555565b90506000805460ff19166001179055919050565b6000805460ff16611240576040805162461bcd60e51b815260206004820152601f6024820152600080516020612486833981519152604482015290519081900360640190fd5b6000805460ff19169055610cd783338461193d565b61125d610fad565b61129c576040805162461bcd60e51b81526020600482018190526024820152600080516020612638833981519152604482015290519081900360640190fd5b610c7981611aa2565b6000805460ff166112eb576040805162461bcd60e51b815260206004820152601f6024820152600080516020612486833981519152604482015290519081900360640190fd5b6000805460ff191690556111e633808461193d565b3390565b6001600160a01b0383166113495760405162461bcd60e51b81526004018080602001828103825260248152602001806126e46024913960400191505060405180910390fd5b6001600160a01b03821661138e5760405162461bcd60e51b81526004018080602001828103825260228152602001806124ee6022913960400191505060405180910390fd5b6001600160a01b03808416600081815260026020908152604080832094871680845294825291829020859055815185815291517f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259281900390910190a3505050565b336001600160a01b038416148061140f5750336001600160a01b038316145b61144a5760405162461bcd60e51b815260040180806020018281038252603081526020018061258f6030913960400191505060405180910390fd5b611455838383611b43565b505050565b600081848411156114e95760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b838110156114ae578181015183820152602001611496565b50505050905090810190601f1680156114db5780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b505050900390565b6000828201838110156109af576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b610f9a8282611ca1565b60006115698361156484611587565b61154b565b600a546109ab906001600160a01b031685308563ffffffff611d9316565b633b9aca000290565b60008082516040146115a157600080fd5b505060208101516040909101519091565b60408051602080825281830190925260609160208201818038833950505060208101929092525090565b6115fe8360405180806125676028913960280190506040518091039020611df3565b6116395760405162461bcd60e51b81526004018080602001828103825260318152602001806125106031913960400191505060405180910390fd5b60006060846001600160a01b0316604051808061256760289139602801905060405180910390208787878760405160240180856001600160a01b03166001600160a01b03168152602001846001600160a01b03166001600160a01b0316815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b838110156116da5781810151838201526020016116c2565b50505050905090810190601f1680156117075780820380516001836020036101000a031916815260200191505b5060408051601f198184030181529181526020820180516001600160e01b03166001600160e01b0319909a16999099178952518151919890975087965094509250829150849050835b6020831061176f5780518252601f199092019160209182019101611750565b6001836020036101000a0380198251168184511680821785525050505050509050019150506000604051808303816000865af19150503d80600081146117d1576040519150601f19603f3d011682016040523d82523d6000602084013e6117d6565b606091505b50915091508181906118295760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156114ae578181015183820152602001611496565b50505050505050565b610f9a8282611e0f565b610f9a8282611f0b565b61185760048263ffffffff611f5f16565b6040516001600160a01b038216907f6ae172837ea30b801fbfcdd4108aa1d5bf8ff775444fd70256b44e6bf3dfc3f690600090a250565b61189f60048263ffffffff611fe016565b6040516001600160a01b038216907fe94479a9f7e1952cc78f2d6baab678adc1b772d936c6583def489e524cb6669290600090a250565b60006001600160a01b03821661191d5760405162461bcd60e51b81526004018080602001828103825260228152602001806126586022913960400191505060405180910390fd5b506001600160a01b03166000908152602091909152604090205460ff1690565b60006119498383611832565b600061195483612047565b600a54604080516370a0823160e01b815230600482015290519293506000926001600160a01b03909216916370a0823191602480820192602092909190829003018186803b1580156119a557600080fd5b505afa1580156119b9573d6000803e3d6000fd5b505050506040513d60208110156119cf57600080fd5b5051905081811015611a7957600a546001600160a01b03166340c10f19306119fd858563ffffffff61205116565b6040518363ffffffff1660e01b815260040180836001600160a01b03166001600160a01b0316815260200182815260200192505050602060405180830381600087803b158015611a4c57600080fd5b505af1158015611a60573d6000803e3d6000fd5b505050506040513d6020811015611a7657600080fd5b50505b600a54611a96906001600160a01b0316878463ffffffff61209316565b50600195945050505050565b6001600160a01b038116611ae75760405162461bcd60e51b81526004018080602001828103825260268152602001806124c86026913960400191505060405180910390fd5b6005546040516001600160a01b038084169216907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a3600580546001600160a01b0319166001600160a01b0392909216919091179055565b6001600160a01b038316611b885760405162461bcd60e51b81526004018080602001828103825260258152602001806126bf6025913960400191505060405180910390fd5b6001600160a01b038216611bcd5760405162461bcd60e51b81526004018080602001828103825260238152602001806124636023913960400191505060405180910390fd5b611c1081604051806060016040528060268152602001612541602691396001600160a01b038616600090815260016020526040902054919063ffffffff61145a16565b6001600160a01b038085166000908152600160205260408082209390935590841681522054611c45908263ffffffff6114f116565b6001600160a01b0380841660008181526001602090815260409182902094909455805185815290519193928716927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef92918290030190a3505050565b6001600160a01b038216611cfc576040805162461bcd60e51b815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f206164647265737300604482015290519081900360640190fd5b600354611d0f908263ffffffff6114f116565b6003556001600160a01b038216600090815260016020526040902054611d3b908263ffffffff6114f116565b6001600160a01b03831660008181526001602090815260408083209490945583518581529351929391927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9281900390910190a35050565b604080516001600160a01b0385811660248301528416604482015260648082018490528251808303909101815260849091019091526020810180516001600160e01b03166323b872dd60e01b179052611ded9085906120e1565b50505050565b6000611dfe83612299565b80156109af57506109af83836122cc565b6001600160a01b038216611e545760405162461bcd60e51b815260040180806020018281038252602181526020018061269e6021913960400191505060405180910390fd5b611e97816040518060600160405280602281526020016124a6602291396001600160a01b038516600090815260016020526040902054919063ffffffff61145a16565b6001600160a01b038316600090815260016020526040902055600354611ec3908263ffffffff61205116565b6003556040805182815290516000916001600160a01b038516917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9181900360200190a35050565b611f158282611832565b610f9a82611f21611300565b6109a68460405180606001604052806024815260200161267a602491396001600160a01b03881660009081526002602052604081209061097f611300565b611f6982826118d6565b15611fbb576040805162461bcd60e51b815260206004820152601f60248201527f526f6c65733a206163636f756e7420616c72656164792068617320726f6c6500604482015290519081900360640190fd5b6001600160a01b0316600090815260209190915260409020805460ff19166001179055565b611fea82826118d6565b6120255760405162461bcd60e51b81526004018080602001828103825260218152602001806125ef6021913960400191505060405180910390fd5b6001600160a01b0316600090815260209190915260409020805460ff19169055565b633b9aca00900490565b60006109af83836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f77000081525061145a565b604080516001600160a01b038416602482015260448082018490528251808303909101815260649091019091526020810180516001600160e01b031663a9059cbb60e01b1790526114559084905b6120f3826001600160a01b03166122f2565b612144576040805162461bcd60e51b815260206004820152601f60248201527f5361666545524332303a2063616c6c20746f206e6f6e2d636f6e747261637400604482015290519081900360640190fd5b60006060836001600160a01b0316836040518082805190602001908083835b602083106121825780518252601f199092019160209182019101612163565b6001836020036101000a0380198251168184511680821785525050505050509050019150506000604051808303816000865af19150503d80600081146121e4576040519150601f19603f3d011682016040523d82523d6000602084013e6121e9565b606091505b509150915081612240576040805162461bcd60e51b815260206004820181905260248201527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564604482015290519081900360640190fd5b805115611ded5780806020019051602081101561225c57600080fd5b5051611ded5760405162461bcd60e51b815260040180806020018281038252602a81526020018061272e602a913960400191505060405180910390fd5b60006122ac826301ffc9a760e01b6122cc565b801561112a57506122c5826001600160e01b03196122cc565b1592915050565b60008060006122db858561232e565b915091508180156122e95750805b95945050505050565b6000813f7fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a47081811480159061232657508115155b949350505050565b604080516001600160e01b0319831660248083019190915282518083039091018152604490910182526020810180516001600160e01b03166301ffc9a760e01b1781529151815160009384939284926060926001600160a01b038a169261753092879282918083835b602083106123b65780518252601f199092019160209182019101612397565b6001836020036101000a0380198251168184511680821785525050505050509050019150506000604051808303818686fa925050503d8060008114612417576040519150601f19603f3d011682016040523d82523d6000602084013e61241c565b606091505b509150915060208151101561243a576000809450945050505061245b565b8181806020019051602081101561245057600080fd5b505190955093505050505b925092905056fe45524332303a207472616e7366657220746f20746865207a65726f20616464726573735265656e7472616e637947756172643a207265656e7472616e742063616c6c0045524332303a206275726e20616d6f756e7420657863656564732062616c616e63654f776e61626c653a206e6577206f776e657220697320746865207a65726f206164647265737345524332303a20617070726f766520746f20746865207a65726f206164647265737345524332304f6e417070726f76653a207370656e64657220646f65736e277420737570706f7274206f6e417070726f766545524332303a207472616e7366657220616d6f756e7420657863656564732062616c616e63656f6e417070726f766528616464726573732c616464726573732c75696e743235362c62797465732953656967546f6b656e3a206f6e6c792073656e646572206f7220726563697069656e742063616e207472616e736665724d696e746572526f6c653a2063616c6c657220646f6573206e6f74206861766520746865204d696e74657220726f6c65526f6c65733a206163636f756e7420646f6573206e6f74206861766520726f6c6545524332303a207472616e7366657220616d6f756e74206578636565647320616c6c6f77616e63654f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572526f6c65733a206163636f756e7420697320746865207a65726f206164647265737345524332303a206275726e20616d6f756e74206578636565647320616c6c6f77616e636545524332303a206275726e2066726f6d20746865207a65726f206164647265737345524332303a207472616e736665722066726f6d20746865207a65726f206164647265737345524332303a20617070726f76652066726f6d20746865207a65726f206164647265737357544f4e3a206f6e6c792061636365707420544f4e20617070726f76652063616c6c6261636b5361666545524332303a204552433230206f7065726174696f6e20646964206e6f74207375636365656445524332303a2064656372656173656420616c6c6f77616e63652062656c6f77207a65726fa265627a7a723158202bfe31d10b317d00b331127684b94dbdcf8eb8656dc5af1f2f15bd57a3cd166664736f6c634300050c00326f6e417070726f766528616464726573732c616464726573732c75696e743235362c627974657329526f6c65733a206163636f756e7420697320746865207a65726f2061646472657373"

// DeployWTON deploys a new Ethereum contract, binding an instance of WTON to it.
func DeployWTON(auth *bind.TransactOpts, backend bind.ContractBackend, _ton common.Address) (common.Address, *types.Transaction, *WTON, error) {
	parsed, err := abi.JSON(strings.NewReader(WTONABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(WTONBin), backend, _ton)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &WTON{WTONCaller: WTONCaller{contract: contract}, WTONTransactor: WTONTransactor{contract: contract}, WTONFilterer: WTONFilterer{contract: contract}}, nil
}

// WTON is an auto generated Go binding around an Ethereum contract.
type WTON struct {
	WTONCaller     // Read-only binding to the contract
	WTONTransactor // Write-only binding to the contract
	WTONFilterer   // Log filterer for contract events
}

// WTONCaller is an auto generated read-only Go binding around an Ethereum contract.
type WTONCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WTONTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WTONTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WTONFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WTONFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WTONSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WTONSession struct {
	Contract     *WTON             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WTONCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WTONCallerSession struct {
	Contract *WTONCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// WTONTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WTONTransactorSession struct {
	Contract     *WTONTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WTONRaw is an auto generated low-level Go binding around an Ethereum contract.
type WTONRaw struct {
	Contract *WTON // Generic contract binding to access the raw methods on
}

// WTONCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WTONCallerRaw struct {
	Contract *WTONCaller // Generic read-only contract binding to access the raw methods on
}

// WTONTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WTONTransactorRaw struct {
	Contract *WTONTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWTON creates a new instance of WTON, bound to a specific deployed contract.
func NewWTON(address common.Address, backend bind.ContractBackend) (*WTON, error) {
	contract, err := bindWTON(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WTON{WTONCaller: WTONCaller{contract: contract}, WTONTransactor: WTONTransactor{contract: contract}, WTONFilterer: WTONFilterer{contract: contract}}, nil
}

// NewWTONCaller creates a new read-only instance of WTON, bound to a specific deployed contract.
func NewWTONCaller(address common.Address, caller bind.ContractCaller) (*WTONCaller, error) {
	contract, err := bindWTON(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WTONCaller{contract: contract}, nil
}

// NewWTONTransactor creates a new write-only instance of WTON, bound to a specific deployed contract.
func NewWTONTransactor(address common.Address, transactor bind.ContractTransactor) (*WTONTransactor, error) {
	contract, err := bindWTON(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WTONTransactor{contract: contract}, nil
}

// NewWTONFilterer creates a new log filterer instance of WTON, bound to a specific deployed contract.
func NewWTONFilterer(address common.Address, filterer bind.ContractFilterer) (*WTONFilterer, error) {
	contract, err := bindWTON(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WTONFilterer{contract: contract}, nil
}

// bindWTON binds a generic wrapper to an already deployed contract.
func bindWTON(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(WTONABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WTON *WTONRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _WTON.Contract.WTONCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WTON *WTONRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WTON.Contract.WTONTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WTON *WTONRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WTON.Contract.WTONTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WTON *WTONCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _WTON.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WTON *WTONTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WTON.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WTON *WTONTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WTON.Contract.contract.Transact(opts, method, params...)
}

// INTERFACEIDONAPPROVE is a free data retrieval call binding the contract method 0x6cd28f9a.
//
// Solidity: function INTERFACE_ID_ON_APPROVE() constant returns(bytes4)
func (_WTON *WTONCaller) INTERFACEIDONAPPROVE(opts *bind.CallOpts) ([4]byte, error) {
	var (
		ret0 = new([4]byte)
	)
	out := ret0
	err := _WTON.contract.Call(opts, out, "INTERFACE_ID_ON_APPROVE")
	return *ret0, err
}

// INTERFACEIDONAPPROVE is a free data retrieval call binding the contract method 0x6cd28f9a.
//
// Solidity: function INTERFACE_ID_ON_APPROVE() constant returns(bytes4)
func (_WTON *WTONSession) INTERFACEIDONAPPROVE() ([4]byte, error) {
	return _WTON.Contract.INTERFACEIDONAPPROVE(&_WTON.CallOpts)
}

// INTERFACEIDONAPPROVE is a free data retrieval call binding the contract method 0x6cd28f9a.
//
// Solidity: function INTERFACE_ID_ON_APPROVE() constant returns(bytes4)
func (_WTON *WTONCallerSession) INTERFACEIDONAPPROVE() ([4]byte, error) {
	return _WTON.Contract.INTERFACEIDONAPPROVE(&_WTON.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_WTON *WTONCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _WTON.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_WTON *WTONSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _WTON.Contract.Allowance(&_WTON.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_WTON *WTONCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _WTON.Contract.Allowance(&_WTON.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_WTON *WTONCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _WTON.contract.Call(opts, out, "balanceOf", account)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_WTON *WTONSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _WTON.Contract.BalanceOf(&_WTON.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_WTON *WTONCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _WTON.Contract.BalanceOf(&_WTON.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_WTON *WTONCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _WTON.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_WTON *WTONSession) Decimals() (uint8, error) {
	return _WTON.Contract.Decimals(&_WTON.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_WTON *WTONCallerSession) Decimals() (uint8, error) {
	return _WTON.Contract.Decimals(&_WTON.CallOpts)
}

// IsMinter is a free data retrieval call binding the contract method 0xaa271e1a.
//
// Solidity: function isMinter(address account) constant returns(bool)
func (_WTON *WTONCaller) IsMinter(opts *bind.CallOpts, account common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _WTON.contract.Call(opts, out, "isMinter", account)
	return *ret0, err
}

// IsMinter is a free data retrieval call binding the contract method 0xaa271e1a.
//
// Solidity: function isMinter(address account) constant returns(bool)
func (_WTON *WTONSession) IsMinter(account common.Address) (bool, error) {
	return _WTON.Contract.IsMinter(&_WTON.CallOpts, account)
}

// IsMinter is a free data retrieval call binding the contract method 0xaa271e1a.
//
// Solidity: function isMinter(address account) constant returns(bool)
func (_WTON *WTONCallerSession) IsMinter(account common.Address) (bool, error) {
	return _WTON.Contract.IsMinter(&_WTON.CallOpts, account)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_WTON *WTONCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _WTON.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_WTON *WTONSession) IsOwner() (bool, error) {
	return _WTON.Contract.IsOwner(&_WTON.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_WTON *WTONCallerSession) IsOwner() (bool, error) {
	return _WTON.Contract.IsOwner(&_WTON.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_WTON *WTONCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _WTON.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_WTON *WTONSession) Name() (string, error) {
	return _WTON.Contract.Name(&_WTON.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_WTON *WTONCallerSession) Name() (string, error) {
	return _WTON.Contract.Name(&_WTON.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_WTON *WTONCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _WTON.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_WTON *WTONSession) Owner() (common.Address, error) {
	return _WTON.Contract.Owner(&_WTON.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_WTON *WTONCallerSession) Owner() (common.Address, error) {
	return _WTON.Contract.Owner(&_WTON.CallOpts)
}

// SeigManager is a free data retrieval call binding the contract method 0x6fb7f558.
//
// Solidity: function seigManager() constant returns(address)
func (_WTON *WTONCaller) SeigManager(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _WTON.contract.Call(opts, out, "seigManager")
	return *ret0, err
}

// SeigManager is a free data retrieval call binding the contract method 0x6fb7f558.
//
// Solidity: function seigManager() constant returns(address)
func (_WTON *WTONSession) SeigManager() (common.Address, error) {
	return _WTON.Contract.SeigManager(&_WTON.CallOpts)
}

// SeigManager is a free data retrieval call binding the contract method 0x6fb7f558.
//
// Solidity: function seigManager() constant returns(address)
func (_WTON *WTONCallerSession) SeigManager() (common.Address, error) {
	return _WTON.Contract.SeigManager(&_WTON.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) constant returns(bool)
func (_WTON *WTONCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _WTON.contract.Call(opts, out, "supportsInterface", interfaceId)
	return *ret0, err
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) constant returns(bool)
func (_WTON *WTONSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _WTON.Contract.SupportsInterface(&_WTON.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) constant returns(bool)
func (_WTON *WTONCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _WTON.Contract.SupportsInterface(&_WTON.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_WTON *WTONCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _WTON.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_WTON *WTONSession) Symbol() (string, error) {
	return _WTON.Contract.Symbol(&_WTON.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_WTON *WTONCallerSession) Symbol() (string, error) {
	return _WTON.Contract.Symbol(&_WTON.CallOpts)
}

// Ton is a free data retrieval call binding the contract method 0xcc48b947.
//
// Solidity: function ton() constant returns(address)
func (_WTON *WTONCaller) Ton(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _WTON.contract.Call(opts, out, "ton")
	return *ret0, err
}

// Ton is a free data retrieval call binding the contract method 0xcc48b947.
//
// Solidity: function ton() constant returns(address)
func (_WTON *WTONSession) Ton() (common.Address, error) {
	return _WTON.Contract.Ton(&_WTON.CallOpts)
}

// Ton is a free data retrieval call binding the contract method 0xcc48b947.
//
// Solidity: function ton() constant returns(address)
func (_WTON *WTONCallerSession) Ton() (common.Address, error) {
	return _WTON.Contract.Ton(&_WTON.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_WTON *WTONCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _WTON.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_WTON *WTONSession) TotalSupply() (*big.Int, error) {
	return _WTON.Contract.TotalSupply(&_WTON.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_WTON *WTONCallerSession) TotalSupply() (*big.Int, error) {
	return _WTON.Contract.TotalSupply(&_WTON.CallOpts)
}

// AddMinter is a paid mutator transaction binding the contract method 0x983b2d56.
//
// Solidity: function addMinter(address account) returns()
func (_WTON *WTONTransactor) AddMinter(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _WTON.contract.Transact(opts, "addMinter", account)
}

// AddMinter is a paid mutator transaction binding the contract method 0x983b2d56.
//
// Solidity: function addMinter(address account) returns()
func (_WTON *WTONSession) AddMinter(account common.Address) (*types.Transaction, error) {
	return _WTON.Contract.AddMinter(&_WTON.TransactOpts, account)
}

// AddMinter is a paid mutator transaction binding the contract method 0x983b2d56.
//
// Solidity: function addMinter(address account) returns()
func (_WTON *WTONTransactorSession) AddMinter(account common.Address) (*types.Transaction, error) {
	return _WTON.Contract.AddMinter(&_WTON.TransactOpts, account)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_WTON *WTONTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WTON.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_WTON *WTONSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WTON.Contract.Approve(&_WTON.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_WTON *WTONTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WTON.Contract.Approve(&_WTON.TransactOpts, spender, amount)
}

// ApproveAndCall is a paid mutator transaction binding the contract method 0xcae9ca51.
//
// Solidity: function approveAndCall(address spender, uint256 amount, bytes data) returns(bool)
func (_WTON *WTONTransactor) ApproveAndCall(opts *bind.TransactOpts, spender common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _WTON.contract.Transact(opts, "approveAndCall", spender, amount, data)
}

// ApproveAndCall is a paid mutator transaction binding the contract method 0xcae9ca51.
//
// Solidity: function approveAndCall(address spender, uint256 amount, bytes data) returns(bool)
func (_WTON *WTONSession) ApproveAndCall(spender common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _WTON.Contract.ApproveAndCall(&_WTON.TransactOpts, spender, amount, data)
}

// ApproveAndCall is a paid mutator transaction binding the contract method 0xcae9ca51.
//
// Solidity: function approveAndCall(address spender, uint256 amount, bytes data) returns(bool)
func (_WTON *WTONTransactorSession) ApproveAndCall(spender common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _WTON.Contract.ApproveAndCall(&_WTON.TransactOpts, spender, amount, data)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_WTON *WTONTransactor) Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _WTON.contract.Transact(opts, "burn", amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_WTON *WTONSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _WTON.Contract.Burn(&_WTON.TransactOpts, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_WTON *WTONTransactorSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _WTON.Contract.Burn(&_WTON.TransactOpts, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_WTON *WTONTransactor) BurnFrom(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WTON.contract.Transact(opts, "burnFrom", account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_WTON *WTONSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WTON.Contract.BurnFrom(&_WTON.TransactOpts, account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_WTON *WTONTransactorSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WTON.Contract.BurnFrom(&_WTON.TransactOpts, account, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_WTON *WTONTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _WTON.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_WTON *WTONSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _WTON.Contract.DecreaseAllowance(&_WTON.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_WTON *WTONTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _WTON.Contract.DecreaseAllowance(&_WTON.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_WTON *WTONTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _WTON.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_WTON *WTONSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _WTON.Contract.IncreaseAllowance(&_WTON.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_WTON *WTONTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _WTON.Contract.IncreaseAllowance(&_WTON.TransactOpts, spender, addedValue)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns(bool)
func (_WTON *WTONTransactor) Mint(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WTON.contract.Transact(opts, "mint", account, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns(bool)
func (_WTON *WTONSession) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WTON.Contract.Mint(&_WTON.TransactOpts, account, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns(bool)
func (_WTON *WTONTransactorSession) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WTON.Contract.Mint(&_WTON.TransactOpts, account, amount)
}

// OnApprove is a paid mutator transaction binding the contract method 0x4273ca16.
//
// Solidity: function onApprove(address owner, address spender, uint256 tonAmount, bytes data) returns(bool)
func (_WTON *WTONTransactor) OnApprove(opts *bind.TransactOpts, owner common.Address, spender common.Address, tonAmount *big.Int, data []byte) (*types.Transaction, error) {
	return _WTON.contract.Transact(opts, "onApprove", owner, spender, tonAmount, data)
}

// OnApprove is a paid mutator transaction binding the contract method 0x4273ca16.
//
// Solidity: function onApprove(address owner, address spender, uint256 tonAmount, bytes data) returns(bool)
func (_WTON *WTONSession) OnApprove(owner common.Address, spender common.Address, tonAmount *big.Int, data []byte) (*types.Transaction, error) {
	return _WTON.Contract.OnApprove(&_WTON.TransactOpts, owner, spender, tonAmount, data)
}

// OnApprove is a paid mutator transaction binding the contract method 0x4273ca16.
//
// Solidity: function onApprove(address owner, address spender, uint256 tonAmount, bytes data) returns(bool)
func (_WTON *WTONTransactorSession) OnApprove(owner common.Address, spender common.Address, tonAmount *big.Int, data []byte) (*types.Transaction, error) {
	return _WTON.Contract.OnApprove(&_WTON.TransactOpts, owner, spender, tonAmount, data)
}

// RenounceMinter is a paid mutator transaction binding the contract method 0x5f112c68.
//
// Solidity: function renounceMinter(address target) returns()
func (_WTON *WTONTransactor) RenounceMinter(opts *bind.TransactOpts, target common.Address) (*types.Transaction, error) {
	return _WTON.contract.Transact(opts, "renounceMinter", target)
}

// RenounceMinter is a paid mutator transaction binding the contract method 0x5f112c68.
//
// Solidity: function renounceMinter(address target) returns()
func (_WTON *WTONSession) RenounceMinter(target common.Address) (*types.Transaction, error) {
	return _WTON.Contract.RenounceMinter(&_WTON.TransactOpts, target)
}

// RenounceMinter is a paid mutator transaction binding the contract method 0x5f112c68.
//
// Solidity: function renounceMinter(address target) returns()
func (_WTON *WTONTransactorSession) RenounceMinter(target common.Address) (*types.Transaction, error) {
	return _WTON.Contract.RenounceMinter(&_WTON.TransactOpts, target)
}

// RenounceMinter0 is a paid mutator transaction binding the contract method 0x98650275.
//
// Solidity: function renounceMinter() returns()
func (_WTON *WTONTransactor) RenounceMinter0(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WTON.contract.Transact(opts, "renounceMinter0")
}

// RenounceMinter0 is a paid mutator transaction binding the contract method 0x98650275.
//
// Solidity: function renounceMinter() returns()
func (_WTON *WTONSession) RenounceMinter0() (*types.Transaction, error) {
	return _WTON.Contract.RenounceMinter0(&_WTON.TransactOpts)
}

// RenounceMinter0 is a paid mutator transaction binding the contract method 0x98650275.
//
// Solidity: function renounceMinter() returns()
func (_WTON *WTONTransactorSession) RenounceMinter0() (*types.Transaction, error) {
	return _WTON.Contract.RenounceMinter0(&_WTON.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x38bf3cfa.
//
// Solidity: function renounceOwnership(address target) returns()
func (_WTON *WTONTransactor) RenounceOwnership(opts *bind.TransactOpts, target common.Address) (*types.Transaction, error) {
	return _WTON.contract.Transact(opts, "renounceOwnership", target)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x38bf3cfa.
//
// Solidity: function renounceOwnership(address target) returns()
func (_WTON *WTONSession) RenounceOwnership(target common.Address) (*types.Transaction, error) {
	return _WTON.Contract.RenounceOwnership(&_WTON.TransactOpts, target)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x38bf3cfa.
//
// Solidity: function renounceOwnership(address target) returns()
func (_WTON *WTONTransactorSession) RenounceOwnership(target common.Address) (*types.Transaction, error) {
	return _WTON.Contract.RenounceOwnership(&_WTON.TransactOpts, target)
}

// RenounceOwnership0 is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_WTON *WTONTransactor) RenounceOwnership0(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WTON.contract.Transact(opts, "renounceOwnership0")
}

// RenounceOwnership0 is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_WTON *WTONSession) RenounceOwnership0() (*types.Transaction, error) {
	return _WTON.Contract.RenounceOwnership0(&_WTON.TransactOpts)
}

// RenounceOwnership0 is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_WTON *WTONTransactorSession) RenounceOwnership0() (*types.Transaction, error) {
	return _WTON.Contract.RenounceOwnership0(&_WTON.TransactOpts)
}

// RenouncePauser is a paid mutator transaction binding the contract method 0x41eb24bb.
//
// Solidity: function renouncePauser(address target) returns()
func (_WTON *WTONTransactor) RenouncePauser(opts *bind.TransactOpts, target common.Address) (*types.Transaction, error) {
	return _WTON.contract.Transact(opts, "renouncePauser", target)
}

// RenouncePauser is a paid mutator transaction binding the contract method 0x41eb24bb.
//
// Solidity: function renouncePauser(address target) returns()
func (_WTON *WTONSession) RenouncePauser(target common.Address) (*types.Transaction, error) {
	return _WTON.Contract.RenouncePauser(&_WTON.TransactOpts, target)
}

// RenouncePauser is a paid mutator transaction binding the contract method 0x41eb24bb.
//
// Solidity: function renouncePauser(address target) returns()
func (_WTON *WTONTransactorSession) RenouncePauser(target common.Address) (*types.Transaction, error) {
	return _WTON.Contract.RenouncePauser(&_WTON.TransactOpts, target)
}

// SetSeigManager is a paid mutator transaction binding the contract method 0x7657f20a.
//
// Solidity: function setSeigManager(address _seigManager) returns()
func (_WTON *WTONTransactor) SetSeigManager(opts *bind.TransactOpts, _seigManager common.Address) (*types.Transaction, error) {
	return _WTON.contract.Transact(opts, "setSeigManager", _seigManager)
}

// SetSeigManager is a paid mutator transaction binding the contract method 0x7657f20a.
//
// Solidity: function setSeigManager(address _seigManager) returns()
func (_WTON *WTONSession) SetSeigManager(_seigManager common.Address) (*types.Transaction, error) {
	return _WTON.Contract.SetSeigManager(&_WTON.TransactOpts, _seigManager)
}

// SetSeigManager is a paid mutator transaction binding the contract method 0x7657f20a.
//
// Solidity: function setSeigManager(address _seigManager) returns()
func (_WTON *WTONTransactorSession) SetSeigManager(_seigManager common.Address) (*types.Transaction, error) {
	return _WTON.Contract.SetSeigManager(&_WTON.TransactOpts, _seigManager)
}

// SwapFromTON is a paid mutator transaction binding the contract method 0xe34869d7.
//
// Solidity: function swapFromTON(uint256 tonAmount) returns(bool)
func (_WTON *WTONTransactor) SwapFromTON(opts *bind.TransactOpts, tonAmount *big.Int) (*types.Transaction, error) {
	return _WTON.contract.Transact(opts, "swapFromTON", tonAmount)
}

// SwapFromTON is a paid mutator transaction binding the contract method 0xe34869d7.
//
// Solidity: function swapFromTON(uint256 tonAmount) returns(bool)
func (_WTON *WTONSession) SwapFromTON(tonAmount *big.Int) (*types.Transaction, error) {
	return _WTON.Contract.SwapFromTON(&_WTON.TransactOpts, tonAmount)
}

// SwapFromTON is a paid mutator transaction binding the contract method 0xe34869d7.
//
// Solidity: function swapFromTON(uint256 tonAmount) returns(bool)
func (_WTON *WTONTransactorSession) SwapFromTON(tonAmount *big.Int) (*types.Transaction, error) {
	return _WTON.Contract.SwapFromTON(&_WTON.TransactOpts, tonAmount)
}

// SwapFromTONAndTransfer is a paid mutator transaction binding the contract method 0x588420b7.
//
// Solidity: function swapFromTONAndTransfer(address to, uint256 tonAmount) returns(bool)
func (_WTON *WTONTransactor) SwapFromTONAndTransfer(opts *bind.TransactOpts, to common.Address, tonAmount *big.Int) (*types.Transaction, error) {
	return _WTON.contract.Transact(opts, "swapFromTONAndTransfer", to, tonAmount)
}

// SwapFromTONAndTransfer is a paid mutator transaction binding the contract method 0x588420b7.
//
// Solidity: function swapFromTONAndTransfer(address to, uint256 tonAmount) returns(bool)
func (_WTON *WTONSession) SwapFromTONAndTransfer(to common.Address, tonAmount *big.Int) (*types.Transaction, error) {
	return _WTON.Contract.SwapFromTONAndTransfer(&_WTON.TransactOpts, to, tonAmount)
}

// SwapFromTONAndTransfer is a paid mutator transaction binding the contract method 0x588420b7.
//
// Solidity: function swapFromTONAndTransfer(address to, uint256 tonAmount) returns(bool)
func (_WTON *WTONTransactorSession) SwapFromTONAndTransfer(to common.Address, tonAmount *big.Int) (*types.Transaction, error) {
	return _WTON.Contract.SwapFromTONAndTransfer(&_WTON.TransactOpts, to, tonAmount)
}

// SwapToTON is a paid mutator transaction binding the contract method 0xf53fe70f.
//
// Solidity: function swapToTON(uint256 wtonAmount) returns(bool)
func (_WTON *WTONTransactor) SwapToTON(opts *bind.TransactOpts, wtonAmount *big.Int) (*types.Transaction, error) {
	return _WTON.contract.Transact(opts, "swapToTON", wtonAmount)
}

// SwapToTON is a paid mutator transaction binding the contract method 0xf53fe70f.
//
// Solidity: function swapToTON(uint256 wtonAmount) returns(bool)
func (_WTON *WTONSession) SwapToTON(wtonAmount *big.Int) (*types.Transaction, error) {
	return _WTON.Contract.SwapToTON(&_WTON.TransactOpts, wtonAmount)
}

// SwapToTON is a paid mutator transaction binding the contract method 0xf53fe70f.
//
// Solidity: function swapToTON(uint256 wtonAmount) returns(bool)
func (_WTON *WTONTransactorSession) SwapToTON(wtonAmount *big.Int) (*types.Transaction, error) {
	return _WTON.Contract.SwapToTON(&_WTON.TransactOpts, wtonAmount)
}

// SwapToTONAndTransfer is a paid mutator transaction binding the contract method 0xe3b99e85.
//
// Solidity: function swapToTONAndTransfer(address to, uint256 wtonAmount) returns(bool)
func (_WTON *WTONTransactor) SwapToTONAndTransfer(opts *bind.TransactOpts, to common.Address, wtonAmount *big.Int) (*types.Transaction, error) {
	return _WTON.contract.Transact(opts, "swapToTONAndTransfer", to, wtonAmount)
}

// SwapToTONAndTransfer is a paid mutator transaction binding the contract method 0xe3b99e85.
//
// Solidity: function swapToTONAndTransfer(address to, uint256 wtonAmount) returns(bool)
func (_WTON *WTONSession) SwapToTONAndTransfer(to common.Address, wtonAmount *big.Int) (*types.Transaction, error) {
	return _WTON.Contract.SwapToTONAndTransfer(&_WTON.TransactOpts, to, wtonAmount)
}

// SwapToTONAndTransfer is a paid mutator transaction binding the contract method 0xe3b99e85.
//
// Solidity: function swapToTONAndTransfer(address to, uint256 wtonAmount) returns(bool)
func (_WTON *WTONTransactorSession) SwapToTONAndTransfer(to common.Address, wtonAmount *big.Int) (*types.Transaction, error) {
	return _WTON.Contract.SwapToTONAndTransfer(&_WTON.TransactOpts, to, wtonAmount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_WTON *WTONTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WTON.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_WTON *WTONSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WTON.Contract.Transfer(&_WTON.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_WTON *WTONTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WTON.Contract.Transfer(&_WTON.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_WTON *WTONTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WTON.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_WTON *WTONSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WTON.Contract.TransferFrom(&_WTON.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_WTON *WTONTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WTON.Contract.TransferFrom(&_WTON.TransactOpts, sender, recipient, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0x6d435421.
//
// Solidity: function transferOwnership(address target, address newOwner) returns()
func (_WTON *WTONTransactor) TransferOwnership(opts *bind.TransactOpts, target common.Address, newOwner common.Address) (*types.Transaction, error) {
	return _WTON.contract.Transact(opts, "transferOwnership", target, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0x6d435421.
//
// Solidity: function transferOwnership(address target, address newOwner) returns()
func (_WTON *WTONSession) TransferOwnership(target common.Address, newOwner common.Address) (*types.Transaction, error) {
	return _WTON.Contract.TransferOwnership(&_WTON.TransactOpts, target, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0x6d435421.
//
// Solidity: function transferOwnership(address target, address newOwner) returns()
func (_WTON *WTONTransactorSession) TransferOwnership(target common.Address, newOwner common.Address) (*types.Transaction, error) {
	return _WTON.Contract.TransferOwnership(&_WTON.TransactOpts, target, newOwner)
}

// TransferOwnership0 is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_WTON *WTONTransactor) TransferOwnership0(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _WTON.contract.Transact(opts, "transferOwnership0", newOwner)
}

// TransferOwnership0 is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_WTON *WTONSession) TransferOwnership0(newOwner common.Address) (*types.Transaction, error) {
	return _WTON.Contract.TransferOwnership0(&_WTON.TransactOpts, newOwner)
}

// TransferOwnership0 is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_WTON *WTONTransactorSession) TransferOwnership0(newOwner common.Address) (*types.Transaction, error) {
	return _WTON.Contract.TransferOwnership0(&_WTON.TransactOpts, newOwner)
}

// WTONApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the WTON contract.
type WTONApprovalIterator struct {
	Event *WTONApproval // Event containing the contract specifics and raw log

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
func (it *WTONApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WTONApproval)
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
		it.Event = new(WTONApproval)
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
func (it *WTONApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WTONApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WTONApproval represents a Approval event raised by the WTON contract.
type WTONApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_WTON *WTONFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*WTONApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _WTON.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &WTONApprovalIterator{contract: _WTON.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_WTON *WTONFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *WTONApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _WTON.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WTONApproval)
				if err := _WTON.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_WTON *WTONFilterer) ParseApproval(log types.Log) (*WTONApproval, error) {
	event := new(WTONApproval)
	if err := _WTON.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WTONMinterAddedIterator is returned from FilterMinterAdded and is used to iterate over the raw logs and unpacked data for MinterAdded events raised by the WTON contract.
type WTONMinterAddedIterator struct {
	Event *WTONMinterAdded // Event containing the contract specifics and raw log

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
func (it *WTONMinterAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WTONMinterAdded)
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
		it.Event = new(WTONMinterAdded)
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
func (it *WTONMinterAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WTONMinterAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WTONMinterAdded represents a MinterAdded event raised by the WTON contract.
type WTONMinterAdded struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMinterAdded is a free log retrieval operation binding the contract event 0x6ae172837ea30b801fbfcdd4108aa1d5bf8ff775444fd70256b44e6bf3dfc3f6.
//
// Solidity: event MinterAdded(address indexed account)
func (_WTON *WTONFilterer) FilterMinterAdded(opts *bind.FilterOpts, account []common.Address) (*WTONMinterAddedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _WTON.contract.FilterLogs(opts, "MinterAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return &WTONMinterAddedIterator{contract: _WTON.contract, event: "MinterAdded", logs: logs, sub: sub}, nil
}

// WatchMinterAdded is a free log subscription operation binding the contract event 0x6ae172837ea30b801fbfcdd4108aa1d5bf8ff775444fd70256b44e6bf3dfc3f6.
//
// Solidity: event MinterAdded(address indexed account)
func (_WTON *WTONFilterer) WatchMinterAdded(opts *bind.WatchOpts, sink chan<- *WTONMinterAdded, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _WTON.contract.WatchLogs(opts, "MinterAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WTONMinterAdded)
				if err := _WTON.contract.UnpackLog(event, "MinterAdded", log); err != nil {
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

// ParseMinterAdded is a log parse operation binding the contract event 0x6ae172837ea30b801fbfcdd4108aa1d5bf8ff775444fd70256b44e6bf3dfc3f6.
//
// Solidity: event MinterAdded(address indexed account)
func (_WTON *WTONFilterer) ParseMinterAdded(log types.Log) (*WTONMinterAdded, error) {
	event := new(WTONMinterAdded)
	if err := _WTON.contract.UnpackLog(event, "MinterAdded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WTONMinterRemovedIterator is returned from FilterMinterRemoved and is used to iterate over the raw logs and unpacked data for MinterRemoved events raised by the WTON contract.
type WTONMinterRemovedIterator struct {
	Event *WTONMinterRemoved // Event containing the contract specifics and raw log

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
func (it *WTONMinterRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WTONMinterRemoved)
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
		it.Event = new(WTONMinterRemoved)
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
func (it *WTONMinterRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WTONMinterRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WTONMinterRemoved represents a MinterRemoved event raised by the WTON contract.
type WTONMinterRemoved struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMinterRemoved is a free log retrieval operation binding the contract event 0xe94479a9f7e1952cc78f2d6baab678adc1b772d936c6583def489e524cb66692.
//
// Solidity: event MinterRemoved(address indexed account)
func (_WTON *WTONFilterer) FilterMinterRemoved(opts *bind.FilterOpts, account []common.Address) (*WTONMinterRemovedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _WTON.contract.FilterLogs(opts, "MinterRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return &WTONMinterRemovedIterator{contract: _WTON.contract, event: "MinterRemoved", logs: logs, sub: sub}, nil
}

// WatchMinterRemoved is a free log subscription operation binding the contract event 0xe94479a9f7e1952cc78f2d6baab678adc1b772d936c6583def489e524cb66692.
//
// Solidity: event MinterRemoved(address indexed account)
func (_WTON *WTONFilterer) WatchMinterRemoved(opts *bind.WatchOpts, sink chan<- *WTONMinterRemoved, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _WTON.contract.WatchLogs(opts, "MinterRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WTONMinterRemoved)
				if err := _WTON.contract.UnpackLog(event, "MinterRemoved", log); err != nil {
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

// ParseMinterRemoved is a log parse operation binding the contract event 0xe94479a9f7e1952cc78f2d6baab678adc1b772d936c6583def489e524cb66692.
//
// Solidity: event MinterRemoved(address indexed account)
func (_WTON *WTONFilterer) ParseMinterRemoved(log types.Log) (*WTONMinterRemoved, error) {
	event := new(WTONMinterRemoved)
	if err := _WTON.contract.UnpackLog(event, "MinterRemoved", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WTONOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the WTON contract.
type WTONOwnershipTransferredIterator struct {
	Event *WTONOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *WTONOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WTONOwnershipTransferred)
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
		it.Event = new(WTONOwnershipTransferred)
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
func (it *WTONOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WTONOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WTONOwnershipTransferred represents a OwnershipTransferred event raised by the WTON contract.
type WTONOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_WTON *WTONFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*WTONOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _WTON.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &WTONOwnershipTransferredIterator{contract: _WTON.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_WTON *WTONFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *WTONOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _WTON.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WTONOwnershipTransferred)
				if err := _WTON.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_WTON *WTONFilterer) ParseOwnershipTransferred(log types.Log) (*WTONOwnershipTransferred, error) {
	event := new(WTONOwnershipTransferred)
	if err := _WTON.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WTONTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the WTON contract.
type WTONTransferIterator struct {
	Event *WTONTransfer // Event containing the contract specifics and raw log

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
func (it *WTONTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WTONTransfer)
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
		it.Event = new(WTONTransfer)
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
func (it *WTONTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WTONTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WTONTransfer represents a Transfer event raised by the WTON contract.
type WTONTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_WTON *WTONFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*WTONTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _WTON.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &WTONTransferIterator{contract: _WTON.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_WTON *WTONFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *WTONTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _WTON.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WTONTransfer)
				if err := _WTON.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_WTON *WTONFilterer) ParseTransfer(log types.Log) (*WTONTransfer, error) {
	event := new(WTONTransfer)
	if err := _WTON.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}
