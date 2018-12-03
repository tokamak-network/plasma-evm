// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package rootchain

import (
	"math/big"
	"strings"

	ethereum "github.com/Onther-Tech/plasma-evm"
	"github.com/Onther-Tech/plasma-evm/accounts/abi"
	"github.com/Onther-Tech/plasma-evm/accounts/abi/bind"
	"github.com/Onther-Tech/plasma-evm/common"
	"github.com/Onther-Tech/plasma-evm/core/types"
	"github.com/Onther-Tech/plasma-evm/event"
)

// AddressABI is the input ABI used to generate the binding from.
const AddressABI = "[]"

// AddressBin is the compiled bytecode used for deploying new contracts.
const AddressBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a72305820e665db6b585c5509fcce0f747a603f8d16fd8ea63367357e2ac604507ea5d5910029`

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

// BMTABI is the input ABI used to generate the binding from.
const BMTABI = "[]"

// BMTBin is the compiled bytecode used for deploying new contracts.
const BMTBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a723058201abd1de65faff869f0a2f6a1a1f86d3bba1b44a4a014202b79ab869ae22381ee0029`

// DeployBMT deploys a new Ethereum contract, binding an instance of BMT to it.
func DeployBMT(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BMT, error) {
	parsed, err := abi.JSON(strings.NewReader(BMTABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BMTBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BMT{BMTCaller: BMTCaller{contract: contract}, BMTTransactor: BMTTransactor{contract: contract}, BMTFilterer: BMTFilterer{contract: contract}}, nil
}

// BMT is an auto generated Go binding around an Ethereum contract.
type BMT struct {
	BMTCaller     // Read-only binding to the contract
	BMTTransactor // Write-only binding to the contract
	BMTFilterer   // Log filterer for contract events
}

// BMTCaller is an auto generated read-only Go binding around an Ethereum contract.
type BMTCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BMTTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BMTTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BMTFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BMTFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BMTSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BMTSession struct {
	Contract     *BMT              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BMTCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BMTCallerSession struct {
	Contract *BMTCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BMTTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BMTTransactorSession struct {
	Contract     *BMTTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BMTRaw is an auto generated low-level Go binding around an Ethereum contract.
type BMTRaw struct {
	Contract *BMT // Generic contract binding to access the raw methods on
}

// BMTCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BMTCallerRaw struct {
	Contract *BMTCaller // Generic read-only contract binding to access the raw methods on
}

// BMTTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BMTTransactorRaw struct {
	Contract *BMTTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBMT creates a new instance of BMT, bound to a specific deployed contract.
func NewBMT(address common.Address, backend bind.ContractBackend) (*BMT, error) {
	contract, err := bindBMT(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BMT{BMTCaller: BMTCaller{contract: contract}, BMTTransactor: BMTTransactor{contract: contract}, BMTFilterer: BMTFilterer{contract: contract}}, nil
}

// NewBMTCaller creates a new read-only instance of BMT, bound to a specific deployed contract.
func NewBMTCaller(address common.Address, caller bind.ContractCaller) (*BMTCaller, error) {
	contract, err := bindBMT(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BMTCaller{contract: contract}, nil
}

// NewBMTTransactor creates a new write-only instance of BMT, bound to a specific deployed contract.
func NewBMTTransactor(address common.Address, transactor bind.ContractTransactor) (*BMTTransactor, error) {
	contract, err := bindBMT(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BMTTransactor{contract: contract}, nil
}

// NewBMTFilterer creates a new log filterer instance of BMT, bound to a specific deployed contract.
func NewBMTFilterer(address common.Address, filterer bind.ContractFilterer) (*BMTFilterer, error) {
	contract, err := bindBMT(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BMTFilterer{contract: contract}, nil
}

// bindBMT binds a generic wrapper to an already deployed contract.
func bindBMT(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BMTABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BMT *BMTRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BMT.Contract.BMTCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BMT *BMTRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BMT.Contract.BMTTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BMT *BMTRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BMT.Contract.BMTTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BMT *BMTCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BMT.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BMT *BMTTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BMT.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BMT *BMTTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BMT.Contract.contract.Transact(opts, method, params...)
}

// BitsABI is the input ABI used to generate the binding from.
const BitsABI = "[]"

// BitsBin is the compiled bytecode used for deploying new contracts.
const BitsBin = `0x605a602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a265627a7a72305820e434e72c3930cb53abc0f7accaa2f51065f7905f260d556d109c4751718d04e76c6578706572696d656e74616cf50037`

// DeployBits deploys a new Ethereum contract, binding an instance of Bits to it.
func DeployBits(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Bits, error) {
	parsed, err := abi.JSON(strings.NewReader(BitsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BitsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Bits{BitsCaller: BitsCaller{contract: contract}, BitsTransactor: BitsTransactor{contract: contract}, BitsFilterer: BitsFilterer{contract: contract}}, nil
}

// Bits is an auto generated Go binding around an Ethereum contract.
type Bits struct {
	BitsCaller     // Read-only binding to the contract
	BitsTransactor // Write-only binding to the contract
	BitsFilterer   // Log filterer for contract events
}

// BitsCaller is an auto generated read-only Go binding around an Ethereum contract.
type BitsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BitsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BitsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BitsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BitsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BitsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BitsSession struct {
	Contract     *Bits             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BitsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BitsCallerSession struct {
	Contract *BitsCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BitsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BitsTransactorSession struct {
	Contract     *BitsTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BitsRaw is an auto generated low-level Go binding around an Ethereum contract.
type BitsRaw struct {
	Contract *Bits // Generic contract binding to access the raw methods on
}

// BitsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BitsCallerRaw struct {
	Contract *BitsCaller // Generic read-only contract binding to access the raw methods on
}

// BitsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BitsTransactorRaw struct {
	Contract *BitsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBits creates a new instance of Bits, bound to a specific deployed contract.
func NewBits(address common.Address, backend bind.ContractBackend) (*Bits, error) {
	contract, err := bindBits(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bits{BitsCaller: BitsCaller{contract: contract}, BitsTransactor: BitsTransactor{contract: contract}, BitsFilterer: BitsFilterer{contract: contract}}, nil
}

// NewBitsCaller creates a new read-only instance of Bits, bound to a specific deployed contract.
func NewBitsCaller(address common.Address, caller bind.ContractCaller) (*BitsCaller, error) {
	contract, err := bindBits(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BitsCaller{contract: contract}, nil
}

// NewBitsTransactor creates a new write-only instance of Bits, bound to a specific deployed contract.
func NewBitsTransactor(address common.Address, transactor bind.ContractTransactor) (*BitsTransactor, error) {
	contract, err := bindBits(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BitsTransactor{contract: contract}, nil
}

// NewBitsFilterer creates a new log filterer instance of Bits, bound to a specific deployed contract.
func NewBitsFilterer(address common.Address, filterer bind.ContractFilterer) (*BitsFilterer, error) {
	contract, err := bindBits(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BitsFilterer{contract: contract}, nil
}

// bindBits binds a generic wrapper to an already deployed contract.
func bindBits(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BitsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bits *BitsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Bits.Contract.BitsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bits *BitsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bits.Contract.BitsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bits *BitsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bits.Contract.BitsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bits *BitsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Bits.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bits *BitsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bits.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bits *BitsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bits.Contract.contract.Transact(opts, method, params...)
}

// DataABI is the input ABI used to generate the binding from.
const DataABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"NA_TX_GAS_PRICE\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"NA_TX_GAS_LIMIT\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"APPLY_IN_ROOTCHAIN_SIGNATURE\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"APPLY_IN_CHILDCHAIN_SIGNATURE\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"NA\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// DataBin is the compiled bytecode used for deploying new contracts.
const DataBin = `0x610208610030600b82828239805160001a6073146000811461002057610022565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600436106100835763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416631927ac58811461008857806390e84f56146100a6578063a7b6ae28146100ae578063a89ca766146100c3578063ab73ff05146100cb575b600080fd5b6100906100e0565b60405161009d919061017f565b60405180910390f35b6100906100e8565b6100b66100ef565b60405161009d9190610171565b6100b6610113565b6100d3610137565b60405161009d919061015d565b633b9aca0081565b620186a081565b7fd9afd3a90000000000000000000000000000000000000000000000000000000081565b7fe904e3d90000000000000000000000000000000000000000000000000000000081565b600081565b6101458161018d565b82525050565b610145816101a6565b610145816101cb565b6020810161016b828461013c565b92915050565b6020810161016b828461014b565b6020810161016b8284610154565b73ffffffffffffffffffffffffffffffffffffffff1690565b7fffffffff000000000000000000000000000000000000000000000000000000001690565b905600a265627a7a723058202a0f2da848b99afc90f90f7b9e003c712f741641309d5950895134d99a3ab88b6c6578706572696d656e74616cf50037`

// DeployData deploys a new Ethereum contract, binding an instance of Data to it.
func DeployData(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Data, error) {
	parsed, err := abi.JSON(strings.NewReader(DataABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DataBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Data{DataCaller: DataCaller{contract: contract}, DataTransactor: DataTransactor{contract: contract}, DataFilterer: DataFilterer{contract: contract}}, nil
}

// Data is an auto generated Go binding around an Ethereum contract.
type Data struct {
	DataCaller     // Read-only binding to the contract
	DataTransactor // Write-only binding to the contract
	DataFilterer   // Log filterer for contract events
}

// DataCaller is an auto generated read-only Go binding around an Ethereum contract.
type DataCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DataTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DataTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DataFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DataFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DataSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DataSession struct {
	Contract     *Data             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DataCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DataCallerSession struct {
	Contract *DataCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// DataTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DataTransactorSession struct {
	Contract     *DataTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DataRaw is an auto generated low-level Go binding around an Ethereum contract.
type DataRaw struct {
	Contract *Data // Generic contract binding to access the raw methods on
}

// DataCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DataCallerRaw struct {
	Contract *DataCaller // Generic read-only contract binding to access the raw methods on
}

// DataTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DataTransactorRaw struct {
	Contract *DataTransactor // Generic write-only contract binding to access the raw methods on
}

// NewData creates a new instance of Data, bound to a specific deployed contract.
func NewData(address common.Address, backend bind.ContractBackend) (*Data, error) {
	contract, err := bindData(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Data{DataCaller: DataCaller{contract: contract}, DataTransactor: DataTransactor{contract: contract}, DataFilterer: DataFilterer{contract: contract}}, nil
}

// NewDataCaller creates a new read-only instance of Data, bound to a specific deployed contract.
func NewDataCaller(address common.Address, caller bind.ContractCaller) (*DataCaller, error) {
	contract, err := bindData(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DataCaller{contract: contract}, nil
}

// NewDataTransactor creates a new write-only instance of Data, bound to a specific deployed contract.
func NewDataTransactor(address common.Address, transactor bind.ContractTransactor) (*DataTransactor, error) {
	contract, err := bindData(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DataTransactor{contract: contract}, nil
}

// NewDataFilterer creates a new log filterer instance of Data, bound to a specific deployed contract.
func NewDataFilterer(address common.Address, filterer bind.ContractFilterer) (*DataFilterer, error) {
	contract, err := bindData(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DataFilterer{contract: contract}, nil
}

// bindData binds a generic wrapper to an already deployed contract.
func bindData(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DataABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Data *DataRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Data.Contract.DataCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Data *DataRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Data.Contract.DataTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Data *DataRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Data.Contract.DataTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Data *DataCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Data.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Data *DataTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Data.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Data *DataTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Data.Contract.contract.Transact(opts, method, params...)
}

// APPLYINCHILDCHAINSIGNATURE is a free data retrieval call binding the contract method 0xa89ca766.
//
// Solidity: function APPLY_IN_CHILDCHAIN_SIGNATURE() constant returns(bytes4)
func (_Data *DataCaller) APPLYINCHILDCHAINSIGNATURE(opts *bind.CallOpts) ([4]byte, error) {
	var (
		ret0 = new([4]byte)
	)
	out := ret0
	err := _Data.contract.Call(opts, out, "APPLY_IN_CHILDCHAIN_SIGNATURE")
	return *ret0, err
}

// APPLYINCHILDCHAINSIGNATURE is a free data retrieval call binding the contract method 0xa89ca766.
//
// Solidity: function APPLY_IN_CHILDCHAIN_SIGNATURE() constant returns(bytes4)
func (_Data *DataSession) APPLYINCHILDCHAINSIGNATURE() ([4]byte, error) {
	return _Data.Contract.APPLYINCHILDCHAINSIGNATURE(&_Data.CallOpts)
}

// APPLYINCHILDCHAINSIGNATURE is a free data retrieval call binding the contract method 0xa89ca766.
//
// Solidity: function APPLY_IN_CHILDCHAIN_SIGNATURE() constant returns(bytes4)
func (_Data *DataCallerSession) APPLYINCHILDCHAINSIGNATURE() ([4]byte, error) {
	return _Data.Contract.APPLYINCHILDCHAINSIGNATURE(&_Data.CallOpts)
}

// APPLYINROOTCHAINSIGNATURE is a free data retrieval call binding the contract method 0xa7b6ae28.
//
// Solidity: function APPLY_IN_ROOTCHAIN_SIGNATURE() constant returns(bytes4)
func (_Data *DataCaller) APPLYINROOTCHAINSIGNATURE(opts *bind.CallOpts) ([4]byte, error) {
	var (
		ret0 = new([4]byte)
	)
	out := ret0
	err := _Data.contract.Call(opts, out, "APPLY_IN_ROOTCHAIN_SIGNATURE")
	return *ret0, err
}

// APPLYINROOTCHAINSIGNATURE is a free data retrieval call binding the contract method 0xa7b6ae28.
//
// Solidity: function APPLY_IN_ROOTCHAIN_SIGNATURE() constant returns(bytes4)
func (_Data *DataSession) APPLYINROOTCHAINSIGNATURE() ([4]byte, error) {
	return _Data.Contract.APPLYINROOTCHAINSIGNATURE(&_Data.CallOpts)
}

// APPLYINROOTCHAINSIGNATURE is a free data retrieval call binding the contract method 0xa7b6ae28.
//
// Solidity: function APPLY_IN_ROOTCHAIN_SIGNATURE() constant returns(bytes4)
func (_Data *DataCallerSession) APPLYINROOTCHAINSIGNATURE() ([4]byte, error) {
	return _Data.Contract.APPLYINROOTCHAINSIGNATURE(&_Data.CallOpts)
}

// NA is a free data retrieval call binding the contract method 0xab73ff05.
//
// Solidity: function NA() constant returns(address)
func (_Data *DataCaller) NA(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Data.contract.Call(opts, out, "NA")
	return *ret0, err
}

// NA is a free data retrieval call binding the contract method 0xab73ff05.
//
// Solidity: function NA() constant returns(address)
func (_Data *DataSession) NA() (common.Address, error) {
	return _Data.Contract.NA(&_Data.CallOpts)
}

// NA is a free data retrieval call binding the contract method 0xab73ff05.
//
// Solidity: function NA() constant returns(address)
func (_Data *DataCallerSession) NA() (common.Address, error) {
	return _Data.Contract.NA(&_Data.CallOpts)
}

// NATXGASLIMIT is a free data retrieval call binding the contract method 0x90e84f56.
//
// Solidity: function NA_TX_GAS_LIMIT() constant returns(uint256)
func (_Data *DataCaller) NATXGASLIMIT(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Data.contract.Call(opts, out, "NA_TX_GAS_LIMIT")
	return *ret0, err
}

// NATXGASLIMIT is a free data retrieval call binding the contract method 0x90e84f56.
//
// Solidity: function NA_TX_GAS_LIMIT() constant returns(uint256)
func (_Data *DataSession) NATXGASLIMIT() (*big.Int, error) {
	return _Data.Contract.NATXGASLIMIT(&_Data.CallOpts)
}

// NATXGASLIMIT is a free data retrieval call binding the contract method 0x90e84f56.
//
// Solidity: function NA_TX_GAS_LIMIT() constant returns(uint256)
func (_Data *DataCallerSession) NATXGASLIMIT() (*big.Int, error) {
	return _Data.Contract.NATXGASLIMIT(&_Data.CallOpts)
}

// NATXGASPRICE is a free data retrieval call binding the contract method 0x1927ac58.
//
// Solidity: function NA_TX_GAS_PRICE() constant returns(uint256)
func (_Data *DataCaller) NATXGASPRICE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Data.contract.Call(opts, out, "NA_TX_GAS_PRICE")
	return *ret0, err
}

// NATXGASPRICE is a free data retrieval call binding the contract method 0x1927ac58.
//
// Solidity: function NA_TX_GAS_PRICE() constant returns(uint256)
func (_Data *DataSession) NATXGASPRICE() (*big.Int, error) {
	return _Data.Contract.NATXGASPRICE(&_Data.CallOpts)
}

// NATXGASPRICE is a free data retrieval call binding the contract method 0x1927ac58.
//
// Solidity: function NA_TX_GAS_PRICE() constant returns(uint256)
func (_Data *DataCallerSession) NATXGASPRICE() (*big.Int, error) {
	return _Data.Contract.NATXGASPRICE(&_Data.CallOpts)
}

// MathABI is the input ABI used to generate the binding from.
const MathABI = "[]"

// MathBin is the compiled bytecode used for deploying new contracts.
const MathBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a7230582070fa6eb0cb7a631d2978dfd027d28c10cc0b3967dae2288b2a5ff0d6c9a476fc0029`

// DeployMath deploys a new Ethereum contract, binding an instance of Math to it.
func DeployMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Math, error) {
	parsed, err := abi.JSON(strings.NewReader(MathABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Math{MathCaller: MathCaller{contract: contract}, MathTransactor: MathTransactor{contract: contract}, MathFilterer: MathFilterer{contract: contract}}, nil
}

// Math is an auto generated Go binding around an Ethereum contract.
type Math struct {
	MathCaller     // Read-only binding to the contract
	MathTransactor // Write-only binding to the contract
	MathFilterer   // Log filterer for contract events
}

// MathCaller is an auto generated read-only Go binding around an Ethereum contract.
type MathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MathSession struct {
	Contract     *Math             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MathCallerSession struct {
	Contract *MathCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MathTransactorSession struct {
	Contract     *MathTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MathRaw is an auto generated low-level Go binding around an Ethereum contract.
type MathRaw struct {
	Contract *Math // Generic contract binding to access the raw methods on
}

// MathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MathCallerRaw struct {
	Contract *MathCaller // Generic read-only contract binding to access the raw methods on
}

// MathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MathTransactorRaw struct {
	Contract *MathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMath creates a new instance of Math, bound to a specific deployed contract.
func NewMath(address common.Address, backend bind.ContractBackend) (*Math, error) {
	contract, err := bindMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Math{MathCaller: MathCaller{contract: contract}, MathTransactor: MathTransactor{contract: contract}, MathFilterer: MathFilterer{contract: contract}}, nil
}

// NewMathCaller creates a new read-only instance of Math, bound to a specific deployed contract.
func NewMathCaller(address common.Address, caller bind.ContractCaller) (*MathCaller, error) {
	contract, err := bindMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MathCaller{contract: contract}, nil
}

// NewMathTransactor creates a new write-only instance of Math, bound to a specific deployed contract.
func NewMathTransactor(address common.Address, transactor bind.ContractTransactor) (*MathTransactor, error) {
	contract, err := bindMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MathTransactor{contract: contract}, nil
}

// NewMathFilterer creates a new log filterer instance of Math, bound to a specific deployed contract.
func NewMathFilterer(address common.Address, filterer bind.ContractFilterer) (*MathFilterer, error) {
	contract, err := bindMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MathFilterer{contract: contract}, nil
}

// bindMath binds a generic wrapper to an already deployed contract.
func bindMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Math *MathRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Math.Contract.MathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Math *MathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Math.Contract.MathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Math *MathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Math.Contract.MathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Math *MathCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Math.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Math *MathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Math.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Math *MathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Math.Contract.contract.Transact(opts, method, params...)
}

// PatriciaTreeDataABI is the input ABI used to generate the binding from.
const PatriciaTreeDataABI = "[]"

// PatriciaTreeDataBin is the compiled bytecode used for deploying new contracts.
const PatriciaTreeDataBin = `0x605a602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a265627a7a72305820c34ea91844b4a0d17c23e00f9df039da3c30e086b87834c80f795a7d441fe61b6c6578706572696d656e74616cf50037`

// DeployPatriciaTreeData deploys a new Ethereum contract, binding an instance of PatriciaTreeData to it.
func DeployPatriciaTreeData(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PatriciaTreeData, error) {
	parsed, err := abi.JSON(strings.NewReader(PatriciaTreeDataABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PatriciaTreeDataBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PatriciaTreeData{PatriciaTreeDataCaller: PatriciaTreeDataCaller{contract: contract}, PatriciaTreeDataTransactor: PatriciaTreeDataTransactor{contract: contract}, PatriciaTreeDataFilterer: PatriciaTreeDataFilterer{contract: contract}}, nil
}

// PatriciaTreeData is an auto generated Go binding around an Ethereum contract.
type PatriciaTreeData struct {
	PatriciaTreeDataCaller     // Read-only binding to the contract
	PatriciaTreeDataTransactor // Write-only binding to the contract
	PatriciaTreeDataFilterer   // Log filterer for contract events
}

// PatriciaTreeDataCaller is an auto generated read-only Go binding around an Ethereum contract.
type PatriciaTreeDataCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PatriciaTreeDataTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PatriciaTreeDataTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PatriciaTreeDataFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PatriciaTreeDataFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PatriciaTreeDataSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PatriciaTreeDataSession struct {
	Contract     *PatriciaTreeData // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PatriciaTreeDataCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PatriciaTreeDataCallerSession struct {
	Contract *PatriciaTreeDataCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// PatriciaTreeDataTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PatriciaTreeDataTransactorSession struct {
	Contract     *PatriciaTreeDataTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// PatriciaTreeDataRaw is an auto generated low-level Go binding around an Ethereum contract.
type PatriciaTreeDataRaw struct {
	Contract *PatriciaTreeData // Generic contract binding to access the raw methods on
}

// PatriciaTreeDataCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PatriciaTreeDataCallerRaw struct {
	Contract *PatriciaTreeDataCaller // Generic read-only contract binding to access the raw methods on
}

// PatriciaTreeDataTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PatriciaTreeDataTransactorRaw struct {
	Contract *PatriciaTreeDataTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPatriciaTreeData creates a new instance of PatriciaTreeData, bound to a specific deployed contract.
func NewPatriciaTreeData(address common.Address, backend bind.ContractBackend) (*PatriciaTreeData, error) {
	contract, err := bindPatriciaTreeData(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PatriciaTreeData{PatriciaTreeDataCaller: PatriciaTreeDataCaller{contract: contract}, PatriciaTreeDataTransactor: PatriciaTreeDataTransactor{contract: contract}, PatriciaTreeDataFilterer: PatriciaTreeDataFilterer{contract: contract}}, nil
}

// NewPatriciaTreeDataCaller creates a new read-only instance of PatriciaTreeData, bound to a specific deployed contract.
func NewPatriciaTreeDataCaller(address common.Address, caller bind.ContractCaller) (*PatriciaTreeDataCaller, error) {
	contract, err := bindPatriciaTreeData(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PatriciaTreeDataCaller{contract: contract}, nil
}

// NewPatriciaTreeDataTransactor creates a new write-only instance of PatriciaTreeData, bound to a specific deployed contract.
func NewPatriciaTreeDataTransactor(address common.Address, transactor bind.ContractTransactor) (*PatriciaTreeDataTransactor, error) {
	contract, err := bindPatriciaTreeData(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PatriciaTreeDataTransactor{contract: contract}, nil
}

// NewPatriciaTreeDataFilterer creates a new log filterer instance of PatriciaTreeData, bound to a specific deployed contract.
func NewPatriciaTreeDataFilterer(address common.Address, filterer bind.ContractFilterer) (*PatriciaTreeDataFilterer, error) {
	contract, err := bindPatriciaTreeData(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PatriciaTreeDataFilterer{contract: contract}, nil
}

// bindPatriciaTreeData binds a generic wrapper to an already deployed contract.
func bindPatriciaTreeData(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PatriciaTreeDataABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PatriciaTreeData *PatriciaTreeDataRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PatriciaTreeData.Contract.PatriciaTreeDataCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PatriciaTreeData *PatriciaTreeDataRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PatriciaTreeData.Contract.PatriciaTreeDataTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PatriciaTreeData *PatriciaTreeDataRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PatriciaTreeData.Contract.PatriciaTreeDataTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PatriciaTreeData *PatriciaTreeDataCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PatriciaTreeData.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PatriciaTreeData *PatriciaTreeDataTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PatriciaTreeData.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PatriciaTreeData *PatriciaTreeDataTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PatriciaTreeData.Contract.contract.Transact(opts, method, params...)
}

// PatriciaTreeFaceABI is the input ABI used to generate the binding from.
const PatriciaTreeFaceABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"key\",\"type\":\"bytes\"},{\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"insert\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"getNode\",\"outputs\":[{\"name\":\"nodes\",\"type\":\"bytes32[2]\"},{\"name\":\"labelDatas\",\"type\":\"bytes32[2]\"},{\"name\":\"labelLengths\",\"type\":\"uint256[2]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"key\",\"type\":\"bytes\"}],\"name\":\"getProof\",\"outputs\":[{\"name\":\"branchMask\",\"type\":\"uint256\"},{\"name\":\"_siblings\",\"type\":\"bytes32[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getRootHash\",\"outputs\":[{\"name\":\"h\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getRootEdge\",\"outputs\":[{\"name\":\"node\",\"type\":\"bytes32\"},{\"name\":\"labelData\",\"type\":\"bytes32\"},{\"name\":\"labelLength\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"rootHash\",\"type\":\"bytes32\"},{\"name\":\"key\",\"type\":\"bytes\"},{\"name\":\"value\",\"type\":\"bytes\"},{\"name\":\"branchMask\",\"type\":\"uint256\"},{\"name\":\"siblings\",\"type\":\"bytes32[]\"}],\"name\":\"verifyProof\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// PatriciaTreeFaceBin is the compiled bytecode used for deploying new contracts.
const PatriciaTreeFaceBin = `0x`

// DeployPatriciaTreeFace deploys a new Ethereum contract, binding an instance of PatriciaTreeFace to it.
func DeployPatriciaTreeFace(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PatriciaTreeFace, error) {
	parsed, err := abi.JSON(strings.NewReader(PatriciaTreeFaceABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PatriciaTreeFaceBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PatriciaTreeFace{PatriciaTreeFaceCaller: PatriciaTreeFaceCaller{contract: contract}, PatriciaTreeFaceTransactor: PatriciaTreeFaceTransactor{contract: contract}, PatriciaTreeFaceFilterer: PatriciaTreeFaceFilterer{contract: contract}}, nil
}

// PatriciaTreeFace is an auto generated Go binding around an Ethereum contract.
type PatriciaTreeFace struct {
	PatriciaTreeFaceCaller     // Read-only binding to the contract
	PatriciaTreeFaceTransactor // Write-only binding to the contract
	PatriciaTreeFaceFilterer   // Log filterer for contract events
}

// PatriciaTreeFaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type PatriciaTreeFaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PatriciaTreeFaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PatriciaTreeFaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PatriciaTreeFaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PatriciaTreeFaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PatriciaTreeFaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PatriciaTreeFaceSession struct {
	Contract     *PatriciaTreeFace // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PatriciaTreeFaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PatriciaTreeFaceCallerSession struct {
	Contract *PatriciaTreeFaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// PatriciaTreeFaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PatriciaTreeFaceTransactorSession struct {
	Contract     *PatriciaTreeFaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// PatriciaTreeFaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type PatriciaTreeFaceRaw struct {
	Contract *PatriciaTreeFace // Generic contract binding to access the raw methods on
}

// PatriciaTreeFaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PatriciaTreeFaceCallerRaw struct {
	Contract *PatriciaTreeFaceCaller // Generic read-only contract binding to access the raw methods on
}

// PatriciaTreeFaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PatriciaTreeFaceTransactorRaw struct {
	Contract *PatriciaTreeFaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPatriciaTreeFace creates a new instance of PatriciaTreeFace, bound to a specific deployed contract.
func NewPatriciaTreeFace(address common.Address, backend bind.ContractBackend) (*PatriciaTreeFace, error) {
	contract, err := bindPatriciaTreeFace(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PatriciaTreeFace{PatriciaTreeFaceCaller: PatriciaTreeFaceCaller{contract: contract}, PatriciaTreeFaceTransactor: PatriciaTreeFaceTransactor{contract: contract}, PatriciaTreeFaceFilterer: PatriciaTreeFaceFilterer{contract: contract}}, nil
}

// NewPatriciaTreeFaceCaller creates a new read-only instance of PatriciaTreeFace, bound to a specific deployed contract.
func NewPatriciaTreeFaceCaller(address common.Address, caller bind.ContractCaller) (*PatriciaTreeFaceCaller, error) {
	contract, err := bindPatriciaTreeFace(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PatriciaTreeFaceCaller{contract: contract}, nil
}

// NewPatriciaTreeFaceTransactor creates a new write-only instance of PatriciaTreeFace, bound to a specific deployed contract.
func NewPatriciaTreeFaceTransactor(address common.Address, transactor bind.ContractTransactor) (*PatriciaTreeFaceTransactor, error) {
	contract, err := bindPatriciaTreeFace(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PatriciaTreeFaceTransactor{contract: contract}, nil
}

// NewPatriciaTreeFaceFilterer creates a new log filterer instance of PatriciaTreeFace, bound to a specific deployed contract.
func NewPatriciaTreeFaceFilterer(address common.Address, filterer bind.ContractFilterer) (*PatriciaTreeFaceFilterer, error) {
	contract, err := bindPatriciaTreeFace(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PatriciaTreeFaceFilterer{contract: contract}, nil
}

// bindPatriciaTreeFace binds a generic wrapper to an already deployed contract.
func bindPatriciaTreeFace(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PatriciaTreeFaceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PatriciaTreeFace *PatriciaTreeFaceRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PatriciaTreeFace.Contract.PatriciaTreeFaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PatriciaTreeFace *PatriciaTreeFaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PatriciaTreeFace.Contract.PatriciaTreeFaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PatriciaTreeFace *PatriciaTreeFaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PatriciaTreeFace.Contract.PatriciaTreeFaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PatriciaTreeFace *PatriciaTreeFaceCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PatriciaTreeFace.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PatriciaTreeFace *PatriciaTreeFaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PatriciaTreeFace.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PatriciaTreeFace *PatriciaTreeFaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PatriciaTreeFace.Contract.contract.Transact(opts, method, params...)
}

// GetNode is a free data retrieval call binding the contract method 0x50c946fe.
//
// Solidity: function getNode(hash bytes32) constant returns(nodes bytes32[2], labelDatas bytes32[2], labelLengths uint256[2])
func (_PatriciaTreeFace *PatriciaTreeFaceCaller) GetNode(opts *bind.CallOpts, hash [32]byte) (struct {
	Nodes        [2][32]byte
	LabelDatas   [2][32]byte
	LabelLengths [2]*big.Int
}, error) {
	ret := new(struct {
		Nodes        [2][32]byte
		LabelDatas   [2][32]byte
		LabelLengths [2]*big.Int
	})
	out := ret
	err := _PatriciaTreeFace.contract.Call(opts, out, "getNode", hash)
	return *ret, err
}

// GetNode is a free data retrieval call binding the contract method 0x50c946fe.
//
// Solidity: function getNode(hash bytes32) constant returns(nodes bytes32[2], labelDatas bytes32[2], labelLengths uint256[2])
func (_PatriciaTreeFace *PatriciaTreeFaceSession) GetNode(hash [32]byte) (struct {
	Nodes        [2][32]byte
	LabelDatas   [2][32]byte
	LabelLengths [2]*big.Int
}, error) {
	return _PatriciaTreeFace.Contract.GetNode(&_PatriciaTreeFace.CallOpts, hash)
}

// GetNode is a free data retrieval call binding the contract method 0x50c946fe.
//
// Solidity: function getNode(hash bytes32) constant returns(nodes bytes32[2], labelDatas bytes32[2], labelLengths uint256[2])
func (_PatriciaTreeFace *PatriciaTreeFaceCallerSession) GetNode(hash [32]byte) (struct {
	Nodes        [2][32]byte
	LabelDatas   [2][32]byte
	LabelLengths [2]*big.Int
}, error) {
	return _PatriciaTreeFace.Contract.GetNode(&_PatriciaTreeFace.CallOpts, hash)
}

// GetProof is a free data retrieval call binding the contract method 0x693ac4fb.
//
// Solidity: function getProof(key bytes) constant returns(branchMask uint256, _siblings bytes32[])
func (_PatriciaTreeFace *PatriciaTreeFaceCaller) GetProof(opts *bind.CallOpts, key []byte) (struct {
	BranchMask *big.Int
	Siblings   [][32]byte
}, error) {
	ret := new(struct {
		BranchMask *big.Int
		Siblings   [][32]byte
	})
	out := ret
	err := _PatriciaTreeFace.contract.Call(opts, out, "getProof", key)
	return *ret, err
}

// GetProof is a free data retrieval call binding the contract method 0x693ac4fb.
//
// Solidity: function getProof(key bytes) constant returns(branchMask uint256, _siblings bytes32[])
func (_PatriciaTreeFace *PatriciaTreeFaceSession) GetProof(key []byte) (struct {
	BranchMask *big.Int
	Siblings   [][32]byte
}, error) {
	return _PatriciaTreeFace.Contract.GetProof(&_PatriciaTreeFace.CallOpts, key)
}

// GetProof is a free data retrieval call binding the contract method 0x693ac4fb.
//
// Solidity: function getProof(key bytes) constant returns(branchMask uint256, _siblings bytes32[])
func (_PatriciaTreeFace *PatriciaTreeFaceCallerSession) GetProof(key []byte) (struct {
	BranchMask *big.Int
	Siblings   [][32]byte
}, error) {
	return _PatriciaTreeFace.Contract.GetProof(&_PatriciaTreeFace.CallOpts, key)
}

// GetRootEdge is a free data retrieval call binding the contract method 0xa43914da.
//
// Solidity: function getRootEdge() constant returns(node bytes32, labelData bytes32, labelLength uint256)
func (_PatriciaTreeFace *PatriciaTreeFaceCaller) GetRootEdge(opts *bind.CallOpts) (struct {
	Node        [32]byte
	LabelData   [32]byte
	LabelLength *big.Int
}, error) {
	ret := new(struct {
		Node        [32]byte
		LabelData   [32]byte
		LabelLength *big.Int
	})
	out := ret
	err := _PatriciaTreeFace.contract.Call(opts, out, "getRootEdge")
	return *ret, err
}

// GetRootEdge is a free data retrieval call binding the contract method 0xa43914da.
//
// Solidity: function getRootEdge() constant returns(node bytes32, labelData bytes32, labelLength uint256)
func (_PatriciaTreeFace *PatriciaTreeFaceSession) GetRootEdge() (struct {
	Node        [32]byte
	LabelData   [32]byte
	LabelLength *big.Int
}, error) {
	return _PatriciaTreeFace.Contract.GetRootEdge(&_PatriciaTreeFace.CallOpts)
}

// GetRootEdge is a free data retrieval call binding the contract method 0xa43914da.
//
// Solidity: function getRootEdge() constant returns(node bytes32, labelData bytes32, labelLength uint256)
func (_PatriciaTreeFace *PatriciaTreeFaceCallerSession) GetRootEdge() (struct {
	Node        [32]byte
	LabelData   [32]byte
	LabelLength *big.Int
}, error) {
	return _PatriciaTreeFace.Contract.GetRootEdge(&_PatriciaTreeFace.CallOpts)
}

// GetRootHash is a free data retrieval call binding the contract method 0x80759f1f.
//
// Solidity: function getRootHash() constant returns(h bytes32)
func (_PatriciaTreeFace *PatriciaTreeFaceCaller) GetRootHash(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _PatriciaTreeFace.contract.Call(opts, out, "getRootHash")
	return *ret0, err
}

// GetRootHash is a free data retrieval call binding the contract method 0x80759f1f.
//
// Solidity: function getRootHash() constant returns(h bytes32)
func (_PatriciaTreeFace *PatriciaTreeFaceSession) GetRootHash() ([32]byte, error) {
	return _PatriciaTreeFace.Contract.GetRootHash(&_PatriciaTreeFace.CallOpts)
}

// GetRootHash is a free data retrieval call binding the contract method 0x80759f1f.
//
// Solidity: function getRootHash() constant returns(h bytes32)
func (_PatriciaTreeFace *PatriciaTreeFaceCallerSession) GetRootHash() ([32]byte, error) {
	return _PatriciaTreeFace.Contract.GetRootHash(&_PatriciaTreeFace.CallOpts)
}

// VerifyProof is a free data retrieval call binding the contract method 0xf7e498f6.
//
// Solidity: function verifyProof(rootHash bytes32, key bytes, value bytes, branchMask uint256, siblings bytes32[]) constant returns(success bool)
func (_PatriciaTreeFace *PatriciaTreeFaceCaller) VerifyProof(opts *bind.CallOpts, rootHash [32]byte, key []byte, value []byte, branchMask *big.Int, siblings [][32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _PatriciaTreeFace.contract.Call(opts, out, "verifyProof", rootHash, key, value, branchMask, siblings)
	return *ret0, err
}

// VerifyProof is a free data retrieval call binding the contract method 0xf7e498f6.
//
// Solidity: function verifyProof(rootHash bytes32, key bytes, value bytes, branchMask uint256, siblings bytes32[]) constant returns(success bool)
func (_PatriciaTreeFace *PatriciaTreeFaceSession) VerifyProof(rootHash [32]byte, key []byte, value []byte, branchMask *big.Int, siblings [][32]byte) (bool, error) {
	return _PatriciaTreeFace.Contract.VerifyProof(&_PatriciaTreeFace.CallOpts, rootHash, key, value, branchMask, siblings)
}

// VerifyProof is a free data retrieval call binding the contract method 0xf7e498f6.
//
// Solidity: function verifyProof(rootHash bytes32, key bytes, value bytes, branchMask uint256, siblings bytes32[]) constant returns(success bool)
func (_PatriciaTreeFace *PatriciaTreeFaceCallerSession) VerifyProof(rootHash [32]byte, key []byte, value []byte, branchMask *big.Int, siblings [][32]byte) (bool, error) {
	return _PatriciaTreeFace.Contract.VerifyProof(&_PatriciaTreeFace.CallOpts, rootHash, key, value, branchMask, siblings)
}

// Insert is a paid mutator transaction binding the contract method 0x20ba5b60.
//
// Solidity: function insert(key bytes, value bytes) returns()
func (_PatriciaTreeFace *PatriciaTreeFaceTransactor) Insert(opts *bind.TransactOpts, key []byte, value []byte) (*types.Transaction, error) {
	return _PatriciaTreeFace.contract.Transact(opts, "insert", key, value)
}

// Insert is a paid mutator transaction binding the contract method 0x20ba5b60.
//
// Solidity: function insert(key bytes, value bytes) returns()
func (_PatriciaTreeFace *PatriciaTreeFaceSession) Insert(key []byte, value []byte) (*types.Transaction, error) {
	return _PatriciaTreeFace.Contract.Insert(&_PatriciaTreeFace.TransactOpts, key, value)
}

// Insert is a paid mutator transaction binding the contract method 0x20ba5b60.
//
// Solidity: function insert(key bytes, value bytes) returns()
func (_PatriciaTreeFace *PatriciaTreeFaceTransactorSession) Insert(key []byte, value []byte) (*types.Transaction, error) {
	return _PatriciaTreeFace.Contract.Insert(&_PatriciaTreeFace.TransactOpts, key, value)
}

// RLPABI is the input ABI used to generate the binding from.
const RLPABI = "[]"

// RLPBin is the compiled bytecode used for deploying new contracts.
const RLPBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a723058206d058bb901168727f1a2b64180e1b4e71111721ff46ab883a41ce354f753c6e80029`

// DeployRLP deploys a new Ethereum contract, binding an instance of RLP to it.
func DeployRLP(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RLP, error) {
	parsed, err := abi.JSON(strings.NewReader(RLPABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RLPBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RLP{RLPCaller: RLPCaller{contract: contract}, RLPTransactor: RLPTransactor{contract: contract}, RLPFilterer: RLPFilterer{contract: contract}}, nil
}

// RLP is an auto generated Go binding around an Ethereum contract.
type RLP struct {
	RLPCaller     // Read-only binding to the contract
	RLPTransactor // Write-only binding to the contract
	RLPFilterer   // Log filterer for contract events
}

// RLPCaller is an auto generated read-only Go binding around an Ethereum contract.
type RLPCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RLPTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RLPTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RLPFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RLPFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RLPSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RLPSession struct {
	Contract     *RLP              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RLPCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RLPCallerSession struct {
	Contract *RLPCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// RLPTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RLPTransactorSession struct {
	Contract     *RLPTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RLPRaw is an auto generated low-level Go binding around an Ethereum contract.
type RLPRaw struct {
	Contract *RLP // Generic contract binding to access the raw methods on
}

// RLPCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RLPCallerRaw struct {
	Contract *RLPCaller // Generic read-only contract binding to access the raw methods on
}

// RLPTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RLPTransactorRaw struct {
	Contract *RLPTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRLP creates a new instance of RLP, bound to a specific deployed contract.
func NewRLP(address common.Address, backend bind.ContractBackend) (*RLP, error) {
	contract, err := bindRLP(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RLP{RLPCaller: RLPCaller{contract: contract}, RLPTransactor: RLPTransactor{contract: contract}, RLPFilterer: RLPFilterer{contract: contract}}, nil
}

// NewRLPCaller creates a new read-only instance of RLP, bound to a specific deployed contract.
func NewRLPCaller(address common.Address, caller bind.ContractCaller) (*RLPCaller, error) {
	contract, err := bindRLP(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RLPCaller{contract: contract}, nil
}

// NewRLPTransactor creates a new write-only instance of RLP, bound to a specific deployed contract.
func NewRLPTransactor(address common.Address, transactor bind.ContractTransactor) (*RLPTransactor, error) {
	contract, err := bindRLP(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RLPTransactor{contract: contract}, nil
}

// NewRLPFilterer creates a new log filterer instance of RLP, bound to a specific deployed contract.
func NewRLPFilterer(address common.Address, filterer bind.ContractFilterer) (*RLPFilterer, error) {
	contract, err := bindRLP(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RLPFilterer{contract: contract}, nil
}

// bindRLP binds a generic wrapper to an already deployed contract.
func bindRLP(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RLPABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RLP *RLPRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RLP.Contract.RLPCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RLP *RLPRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RLP.Contract.RLPTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RLP *RLPRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RLP.Contract.RLPTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RLP *RLPCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RLP.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RLP *RLPTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RLP.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RLP *RLPTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RLP.Contract.contract.Transact(opts, method, params...)
}

// RLPEncodeABI is the input ABI used to generate the binding from.
const RLPEncodeABI = "[]"

// RLPEncodeBin is the compiled bytecode used for deploying new contracts.
const RLPEncodeBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a72305820860672072daa09b905ffc40c8e11d71c1c9f8c3df13678ac1b9cc692f19d205a0029`

// DeployRLPEncode deploys a new Ethereum contract, binding an instance of RLPEncode to it.
func DeployRLPEncode(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RLPEncode, error) {
	parsed, err := abi.JSON(strings.NewReader(RLPEncodeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RLPEncodeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RLPEncode{RLPEncodeCaller: RLPEncodeCaller{contract: contract}, RLPEncodeTransactor: RLPEncodeTransactor{contract: contract}, RLPEncodeFilterer: RLPEncodeFilterer{contract: contract}}, nil
}

// RLPEncode is an auto generated Go binding around an Ethereum contract.
type RLPEncode struct {
	RLPEncodeCaller     // Read-only binding to the contract
	RLPEncodeTransactor // Write-only binding to the contract
	RLPEncodeFilterer   // Log filterer for contract events
}

// RLPEncodeCaller is an auto generated read-only Go binding around an Ethereum contract.
type RLPEncodeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RLPEncodeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RLPEncodeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RLPEncodeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RLPEncodeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RLPEncodeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RLPEncodeSession struct {
	Contract     *RLPEncode        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RLPEncodeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RLPEncodeCallerSession struct {
	Contract *RLPEncodeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// RLPEncodeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RLPEncodeTransactorSession struct {
	Contract     *RLPEncodeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// RLPEncodeRaw is an auto generated low-level Go binding around an Ethereum contract.
type RLPEncodeRaw struct {
	Contract *RLPEncode // Generic contract binding to access the raw methods on
}

// RLPEncodeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RLPEncodeCallerRaw struct {
	Contract *RLPEncodeCaller // Generic read-only contract binding to access the raw methods on
}

// RLPEncodeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RLPEncodeTransactorRaw struct {
	Contract *RLPEncodeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRLPEncode creates a new instance of RLPEncode, bound to a specific deployed contract.
func NewRLPEncode(address common.Address, backend bind.ContractBackend) (*RLPEncode, error) {
	contract, err := bindRLPEncode(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RLPEncode{RLPEncodeCaller: RLPEncodeCaller{contract: contract}, RLPEncodeTransactor: RLPEncodeTransactor{contract: contract}, RLPEncodeFilterer: RLPEncodeFilterer{contract: contract}}, nil
}

// NewRLPEncodeCaller creates a new read-only instance of RLPEncode, bound to a specific deployed contract.
func NewRLPEncodeCaller(address common.Address, caller bind.ContractCaller) (*RLPEncodeCaller, error) {
	contract, err := bindRLPEncode(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RLPEncodeCaller{contract: contract}, nil
}

// NewRLPEncodeTransactor creates a new write-only instance of RLPEncode, bound to a specific deployed contract.
func NewRLPEncodeTransactor(address common.Address, transactor bind.ContractTransactor) (*RLPEncodeTransactor, error) {
	contract, err := bindRLPEncode(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RLPEncodeTransactor{contract: contract}, nil
}

// NewRLPEncodeFilterer creates a new log filterer instance of RLPEncode, bound to a specific deployed contract.
func NewRLPEncodeFilterer(address common.Address, filterer bind.ContractFilterer) (*RLPEncodeFilterer, error) {
	contract, err := bindRLPEncode(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RLPEncodeFilterer{contract: contract}, nil
}

// bindRLPEncode binds a generic wrapper to an already deployed contract.
func bindRLPEncode(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RLPEncodeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RLPEncode *RLPEncodeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RLPEncode.Contract.RLPEncodeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RLPEncode *RLPEncodeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RLPEncode.Contract.RLPEncodeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RLPEncode *RLPEncodeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RLPEncode.Contract.RLPEncodeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RLPEncode *RLPEncodeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RLPEncode.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RLPEncode *RLPEncodeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RLPEncode.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RLPEncode *RLPEncodeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RLPEncode.Contract.contract.Transact(opts, method, params...)
}

// RequestableContractIABI is the input ABI used to generate the binding from.
const RequestableContractIABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"isExit\",\"type\":\"bool\"},{\"name\":\"requestId\",\"type\":\"uint256\"},{\"name\":\"requestor\",\"type\":\"address\"},{\"name\":\"trieKey\",\"type\":\"bytes32\"},{\"name\":\"trieValue\",\"type\":\"bytes32\"}],\"name\":\"applyRequestInRootChain\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"isExit\",\"type\":\"bool\"},{\"name\":\"requestId\",\"type\":\"uint256\"},{\"name\":\"requestor\",\"type\":\"address\"},{\"name\":\"trieKey\",\"type\":\"bytes32\"},{\"name\":\"trieValue\",\"type\":\"bytes32\"}],\"name\":\"applyRequestInChildChain\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// RequestableContractIBin is the compiled bytecode used for deploying new contracts.
const RequestableContractIBin = `0x`

// DeployRequestableContractI deploys a new Ethereum contract, binding an instance of RequestableContractI to it.
func DeployRequestableContractI(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RequestableContractI, error) {
	parsed, err := abi.JSON(strings.NewReader(RequestableContractIABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RequestableContractIBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RequestableContractI{RequestableContractICaller: RequestableContractICaller{contract: contract}, RequestableContractITransactor: RequestableContractITransactor{contract: contract}, RequestableContractIFilterer: RequestableContractIFilterer{contract: contract}}, nil
}

// RequestableContractI is an auto generated Go binding around an Ethereum contract.
type RequestableContractI struct {
	RequestableContractICaller     // Read-only binding to the contract
	RequestableContractITransactor // Write-only binding to the contract
	RequestableContractIFilterer   // Log filterer for contract events
}

// RequestableContractICaller is an auto generated read-only Go binding around an Ethereum contract.
type RequestableContractICaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RequestableContractITransactor is an auto generated write-only Go binding around an Ethereum contract.
type RequestableContractITransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RequestableContractIFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RequestableContractIFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RequestableContractISession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RequestableContractISession struct {
	Contract     *RequestableContractI // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// RequestableContractICallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RequestableContractICallerSession struct {
	Contract *RequestableContractICaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// RequestableContractITransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RequestableContractITransactorSession struct {
	Contract     *RequestableContractITransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// RequestableContractIRaw is an auto generated low-level Go binding around an Ethereum contract.
type RequestableContractIRaw struct {
	Contract *RequestableContractI // Generic contract binding to access the raw methods on
}

// RequestableContractICallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RequestableContractICallerRaw struct {
	Contract *RequestableContractICaller // Generic read-only contract binding to access the raw methods on
}

// RequestableContractITransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RequestableContractITransactorRaw struct {
	Contract *RequestableContractITransactor // Generic write-only contract binding to access the raw methods on
}

// NewRequestableContractI creates a new instance of RequestableContractI, bound to a specific deployed contract.
func NewRequestableContractI(address common.Address, backend bind.ContractBackend) (*RequestableContractI, error) {
	contract, err := bindRequestableContractI(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RequestableContractI{RequestableContractICaller: RequestableContractICaller{contract: contract}, RequestableContractITransactor: RequestableContractITransactor{contract: contract}, RequestableContractIFilterer: RequestableContractIFilterer{contract: contract}}, nil
}

// NewRequestableContractICaller creates a new read-only instance of RequestableContractI, bound to a specific deployed contract.
func NewRequestableContractICaller(address common.Address, caller bind.ContractCaller) (*RequestableContractICaller, error) {
	contract, err := bindRequestableContractI(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RequestableContractICaller{contract: contract}, nil
}

// NewRequestableContractITransactor creates a new write-only instance of RequestableContractI, bound to a specific deployed contract.
func NewRequestableContractITransactor(address common.Address, transactor bind.ContractTransactor) (*RequestableContractITransactor, error) {
	contract, err := bindRequestableContractI(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RequestableContractITransactor{contract: contract}, nil
}

// NewRequestableContractIFilterer creates a new log filterer instance of RequestableContractI, bound to a specific deployed contract.
func NewRequestableContractIFilterer(address common.Address, filterer bind.ContractFilterer) (*RequestableContractIFilterer, error) {
	contract, err := bindRequestableContractI(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RequestableContractIFilterer{contract: contract}, nil
}

// bindRequestableContractI binds a generic wrapper to an already deployed contract.
func bindRequestableContractI(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RequestableContractIABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RequestableContractI *RequestableContractIRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RequestableContractI.Contract.RequestableContractICaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RequestableContractI *RequestableContractIRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RequestableContractI.Contract.RequestableContractITransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RequestableContractI *RequestableContractIRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RequestableContractI.Contract.RequestableContractITransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RequestableContractI *RequestableContractICallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RequestableContractI.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RequestableContractI *RequestableContractITransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RequestableContractI.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RequestableContractI *RequestableContractITransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RequestableContractI.Contract.contract.Transact(opts, method, params...)
}

// ApplyRequestInChildChain is a paid mutator transaction binding the contract method 0xe904e3d9.
//
// Solidity: function applyRequestInChildChain(isExit bool, requestId uint256, requestor address, trieKey bytes32, trieValue bytes32) returns(success bool)
func (_RequestableContractI *RequestableContractITransactor) ApplyRequestInChildChain(opts *bind.TransactOpts, isExit bool, requestId *big.Int, requestor common.Address, trieKey [32]byte, trieValue [32]byte) (*types.Transaction, error) {
	return _RequestableContractI.contract.Transact(opts, "applyRequestInChildChain", isExit, requestId, requestor, trieKey, trieValue)
}

// ApplyRequestInChildChain is a paid mutator transaction binding the contract method 0xe904e3d9.
//
// Solidity: function applyRequestInChildChain(isExit bool, requestId uint256, requestor address, trieKey bytes32, trieValue bytes32) returns(success bool)
func (_RequestableContractI *RequestableContractISession) ApplyRequestInChildChain(isExit bool, requestId *big.Int, requestor common.Address, trieKey [32]byte, trieValue [32]byte) (*types.Transaction, error) {
	return _RequestableContractI.Contract.ApplyRequestInChildChain(&_RequestableContractI.TransactOpts, isExit, requestId, requestor, trieKey, trieValue)
}

// ApplyRequestInChildChain is a paid mutator transaction binding the contract method 0xe904e3d9.
//
// Solidity: function applyRequestInChildChain(isExit bool, requestId uint256, requestor address, trieKey bytes32, trieValue bytes32) returns(success bool)
func (_RequestableContractI *RequestableContractITransactorSession) ApplyRequestInChildChain(isExit bool, requestId *big.Int, requestor common.Address, trieKey [32]byte, trieValue [32]byte) (*types.Transaction, error) {
	return _RequestableContractI.Contract.ApplyRequestInChildChain(&_RequestableContractI.TransactOpts, isExit, requestId, requestor, trieKey, trieValue)
}

// ApplyRequestInRootChain is a paid mutator transaction binding the contract method 0xd9afd3a9.
//
// Solidity: function applyRequestInRootChain(isExit bool, requestId uint256, requestor address, trieKey bytes32, trieValue bytes32) returns(success bool)
func (_RequestableContractI *RequestableContractITransactor) ApplyRequestInRootChain(opts *bind.TransactOpts, isExit bool, requestId *big.Int, requestor common.Address, trieKey [32]byte, trieValue [32]byte) (*types.Transaction, error) {
	return _RequestableContractI.contract.Transact(opts, "applyRequestInRootChain", isExit, requestId, requestor, trieKey, trieValue)
}

// ApplyRequestInRootChain is a paid mutator transaction binding the contract method 0xd9afd3a9.
//
// Solidity: function applyRequestInRootChain(isExit bool, requestId uint256, requestor address, trieKey bytes32, trieValue bytes32) returns(success bool)
func (_RequestableContractI *RequestableContractISession) ApplyRequestInRootChain(isExit bool, requestId *big.Int, requestor common.Address, trieKey [32]byte, trieValue [32]byte) (*types.Transaction, error) {
	return _RequestableContractI.Contract.ApplyRequestInRootChain(&_RequestableContractI.TransactOpts, isExit, requestId, requestor, trieKey, trieValue)
}

// ApplyRequestInRootChain is a paid mutator transaction binding the contract method 0xd9afd3a9.
//
// Solidity: function applyRequestInRootChain(isExit bool, requestId uint256, requestor address, trieKey bytes32, trieValue bytes32) returns(success bool)
func (_RequestableContractI *RequestableContractITransactorSession) ApplyRequestInRootChain(isExit bool, requestId *big.Int, requestor common.Address, trieKey [32]byte, trieValue [32]byte) (*types.Transaction, error) {
	return _RequestableContractI.Contract.ApplyRequestInRootChain(&_RequestableContractI.TransactOpts, isExit, requestId, requestor, trieKey, trieValue)
}

// RootChainABI is the input ABI used to generate the binding from.
const RootChainABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"COST_URB_PREPARE\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CP_COMPUTATION\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"epochs\",\"outputs\":[{\"name\":\"requestStart\",\"type\":\"uint64\"},{\"name\":\"requestEnd\",\"type\":\"uint64\"},{\"name\":\"startBlockNumber\",\"type\":\"uint64\"},{\"name\":\"endBlockNumber\",\"type\":\"uint64\"},{\"name\":\"forkedBlockNumber\",\"type\":\"uint64\"},{\"name\":\"firstRequestBlockId\",\"type\":\"uint64\"},{\"name\":\"limit\",\"type\":\"uint64\"},{\"name\":\"timestamp\",\"type\":\"uint64\"},{\"name\":\"isEmpty\",\"type\":\"bool\"},{\"name\":\"initialized\",\"type\":\"bool\"},{\"name\":\"isRequest\",\"type\":\"bool\"},{\"name\":\"userActivated\",\"type\":\"bool\"},{\"name\":\"finalized\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastAppliedForkNumber\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"currentFork\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"COST_URB\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastAppliedERU\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"firstEpoch\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_isTransfer\",\"type\":\"bool\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_trieKey\",\"type\":\"bytes32\"},{\"name\":\"_trieValue\",\"type\":\"bytes32\"}],\"name\":\"startEnter\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_forkNumber\",\"type\":\"uint256\"},{\"name\":\"_blockNumber\",\"type\":\"uint256\"},{\"name\":\"_index\",\"type\":\"uint256\"},{\"name\":\"_receiptData\",\"type\":\"bytes\"},{\"name\":\"_proof\",\"type\":\"bytes\"}],\"name\":\"challengeExit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_isTransfer\",\"type\":\"bool\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_trieKey\",\"type\":\"bytes32\"},{\"name\":\"_trieValue\",\"type\":\"bytes32\"}],\"name\":\"makeERU\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"operator\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"NRBEpochLength\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"applyRequest\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_blockNumber\",\"type\":\"uint256\"},{\"name\":\"_key\",\"type\":\"bytes\"},{\"name\":\"_txByte\",\"type\":\"bytes\"},{\"name\":\"_branchMask\",\"type\":\"uint256\"},{\"name\":\"_siblings\",\"type\":\"bytes32[]\"}],\"name\":\"challengeNullAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastAppliedERO\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"finalizeBlock\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"currentEpoch\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"lastFinalizedBlock\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"development\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"firstFilledORBEpochNumber\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"COST_ERU\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"REQUEST_GAS\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_statesRoot\",\"type\":\"bytes32\"},{\"name\":\"_transactionsRoot\",\"type\":\"bytes32\"},{\"name\":\"_receiptsRoot\",\"type\":\"bytes32\"}],\"name\":\"submitNRB\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAX_REQUESTS\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"COST_NRB\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_statesRoot\",\"type\":\"bytes32\"},{\"name\":\"_transactionsRoot\",\"type\":\"bytes32\"},{\"name\":\"_receiptsRoot\",\"type\":\"bytes32\"}],\"name\":\"submitURB\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CP_WITHHOLDING\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"COST_ORB\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"EROs\",\"outputs\":[{\"name\":\"timestamp\",\"type\":\"uint64\"},{\"name\":\"isExit\",\"type\":\"bool\"},{\"name\":\"isTransfer\",\"type\":\"bool\"},{\"name\":\"finalized\",\"type\":\"bool\"},{\"name\":\"challenged\",\"type\":\"bool\"},{\"name\":\"value\",\"type\":\"uint128\"},{\"name\":\"requestor\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"trieKey\",\"type\":\"bytes32\"},{\"name\":\"trieValue\",\"type\":\"bytes32\"},{\"name\":\"hash\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getNumEROs\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"URBs\",\"outputs\":[{\"name\":\"submitted\",\"type\":\"bool\"},{\"name\":\"epochNumber\",\"type\":\"uint64\"},{\"name\":\"requestStart\",\"type\":\"uint64\"},{\"name\":\"requestEnd\",\"type\":\"uint64\"},{\"name\":\"trie\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"state\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PREPARE_TIMEOUT\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_rootchain\",\"type\":\"address\"},{\"name\":\"_childchain\",\"type\":\"address\"}],\"name\":\"mapRequestableContractByOperator\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_forkNumber\",\"type\":\"uint256\"}],\"name\":\"forked\",\"outputs\":[{\"name\":\"result\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_requestId\",\"type\":\"uint256\"}],\"name\":\"getEROBytes\",\"outputs\":[{\"name\":\"out\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"COST_ERO\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"requestableContracts\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_trieKey\",\"type\":\"bytes32\"},{\"name\":\"_trieValue\",\"type\":\"bytes32\"}],\"name\":\"startExit\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"NULL_ADDRESS\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_statesRoot\",\"type\":\"bytes32\"},{\"name\":\"_transactionsRoot\",\"type\":\"bytes32\"},{\"name\":\"_receiptsRoot\",\"type\":\"bytes32\"}],\"name\":\"submitORB\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"prepareToSubmitURB\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getNumORBs\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"ORBs\",\"outputs\":[{\"name\":\"submitted\",\"type\":\"bool\"},{\"name\":\"epochNumber\",\"type\":\"uint64\"},{\"name\":\"requestStart\",\"type\":\"uint64\"},{\"name\":\"requestEnd\",\"type\":\"uint64\"},{\"name\":\"trie\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"highestBlockNumber\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_requestId\",\"type\":\"uint256\"},{\"name\":\"_userActivated\",\"type\":\"bool\"}],\"name\":\"getRequestFinalized\",\"outputs\":[{\"name\":\"finalized\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"ERUs\",\"outputs\":[{\"name\":\"timestamp\",\"type\":\"uint64\"},{\"name\":\"isExit\",\"type\":\"bool\"},{\"name\":\"isTransfer\",\"type\":\"bool\"},{\"name\":\"finalized\",\"type\":\"bool\"},{\"name\":\"challenged\",\"type\":\"bool\"},{\"name\":\"value\",\"type\":\"uint128\"},{\"name\":\"requestor\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"trieKey\",\"type\":\"bytes32\"},{\"name\":\"trieValue\",\"type\":\"bytes32\"},{\"name\":\"hash\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"blocks\",\"outputs\":[{\"name\":\"forkNumber\",\"type\":\"uint64\"},{\"name\":\"epochNumber\",\"type\":\"uint64\"},{\"name\":\"requestBlockId\",\"type\":\"uint64\"},{\"name\":\"timestamp\",\"type\":\"uint64\"},{\"name\":\"statesRoot\",\"type\":\"bytes32\"},{\"name\":\"transactionsRoot\",\"type\":\"bytes32\"},{\"name\":\"receiptsRoot\",\"type\":\"bytes32\"},{\"name\":\"isRequest\",\"type\":\"bool\"},{\"name\":\"userActivated\",\"type\":\"bool\"},{\"name\":\"challenged\",\"type\":\"bool\"},{\"name\":\"challenging\",\"type\":\"bool\"},{\"name\":\"finalized\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastAppliedBlockNumber\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_development\",\"type\":\"bool\"},{\"name\":\"_NRBEpochLength\",\"type\":\"uint256\"},{\"name\":\"_statesRoot\",\"type\":\"bytes32\"},{\"name\":\"_transactionsRoot\",\"type\":\"bytes32\"},{\"name\":\"_receiptsRoot\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"userActivated\",\"type\":\"bool\"}],\"name\":\"SessionTimeout\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"state\",\"type\":\"uint8\"}],\"name\":\"StateChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newFork\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"forkedBlockNumber\",\"type\":\"uint256\"}],\"name\":\"Forked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"epochNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"startBlockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"endBlockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"requestStart\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"requestEnd\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"epochIsEmpty\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"isRequest\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"userActivated\",\"type\":\"bool\"}],\"name\":\"EpochPrepared\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"fork\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"NRBSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"fork\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"ORBSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"fork\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"URBSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"requestor\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"weiAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"trieKey\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"trieValue\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"isTransfer\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"isExit\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"userActivated\",\"type\":\"bool\"}],\"name\":\"RequestCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"requestor\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"trieKey\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"trieValue\",\"type\":\"bytes32\"}],\"name\":\"ERUCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"forkNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"BlockFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"forkNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"epochNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"startBlockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"endBlockNumber\",\"type\":\"uint256\"}],\"name\":\"EpochFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"userActivated\",\"type\":\"bool\"}],\"name\":\"RequestFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"userActivated\",\"type\":\"bool\"}],\"name\":\"RequestApplied\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"userActivated\",\"type\":\"bool\"}],\"name\":\"RequestChallenged\",\"type\":\"event\"}]"

// RootChainBin is the compiled bytecode used for deploying new contracts.
const RootChainBin = `0x60806040523480156200001157600080fd5b5060405160a0806200566b83398101604090815281516020808401518385015160608601516080909601516000805460ff19168615151761010060a860020a031916610100338102919091178255600985905560018054835260078752888320838052875297822097880184905560028801899055600388018390557f3b543ac468519aa6396aba1a2611925510398aac61d1cc7b45f433e69042eb588054600160c060020a03167801000000000000000000000000000000000000000000000000426001604060020a0316021790557f3b543ac468519aa6396aba1a2611925510398aac61d1cc7b45f433e69042eb5980547f5eff886ea0ce6ca488a3d6e336d6c0f75f46d19b42c06ce5ee98e42c96d256c790975264010000000061ffff1990971690911764ff00000000191686179055949692959194929390926200015f918391906200017e810204565b62000172640100000000620001eb810204565b50505050505062000439565b60048201805464ff000000001916640100000000179055600180546000908152600660209081526040918290208490559154815190815291820183905280517ffb96205e4b3633fd57aa805b26b51ecf528714a10241a4af015929dce86768d99281900390910190a15050565b600480546001908101918290558054600090815260086020908152604080832085845290915290209181146200026c57506001805460009081526008602090815260408083206004546000190184529091529020546001604060020a0378010000000000000000000000000000000000000000000000009091048116909101165b60028201805461010061ff00199091161790558154608060020a60c060020a0319167001000000000000000000000000000000006001604060020a038381169190910291909117808455600954600160c060020a0391821678010000000000000000000000000000000000000000000000009185016000190184168202178555600185018054909216429093160291909117905560008054819060a860020a60ff02191675010000000000000000000000000000000000000000008202179055506000546040517f551dc40198cc79684bb69e4931dba4ac16e4598792ee1c0a5000aeea366d7bb6917501000000000000000000000000000000000000000000900460ff1690808260028111156200038057fe5b60ff16815260200191505060405180910390a16004548254604080519283526001604060020a037001000000000000000000000000000000008304811660208501527801000000000000000000000000000000000000000000000000909204909116828201526000606083018190526080830181905260a0830181905260c0830181905260e0830152517f99d777cc3a347ecdaa852c6438ff2afa73f4f68b801c1e81050e817fdb6171d2918190036101000190a15050565b61522280620004496000396000f30060806040526004361061025b5763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663033cfbed811461026057806308c4fff0146102875780630a88bd041461029c578063164bc2ae14610338578063183d2d1c1461034d578063192adc5b146103625780631f261d59146103775780632e7ab9481461038c5780633cfb871e146103a4578063404f7d66146103d757806340a029a114610411578063570ca7351461043d578063584e349e1461046e57806361d29e83146104835780636299fb241461049857806365d724bc146104d857806375395a58146104ed578063766718081461050257806376e61c36146105175780637b929c271461052f57806383bcbcff146105445780638b5172d01461055c5780638eb288ca1461057157806393248d0214610586578063935212221461059757806394be3aa51461026057806397be455d146105ac578063b17fa6e9146105bd578063b2ae9ba814610260578063b443f3cc146105d2578063b540adba1461065e578063c0e8606414610673578063c19d93fb146106cd578063c2bc88fa14610706578063cb5d742f1461071b578063ce8a2bc214610742578063d1723a961461075a578063d691acd814610260578063da0185f8146107e7578063dac3092014610808578063de0ce17d14610825578063e67123c41461083a578063e6925d081461084b578063ea0c73f614610860578063ea7f22a814610875578063f1f3c46c1461088d578063f28a7afa146108a5578063f4f31de4146108c2578063f4f911db146108da578063fb788a271461096c575b600080fd5b34801561026c57600080fd5b50610275610981565b60408051918252519081900360200190f35b34801561029357600080fd5b5061027561098d565b3480156102a857600080fd5b506102b7600435602435610992565b604080516001604060020a039e8f1681529c8e1660208e01529a8d168c8c0152988c1660608c0152968b1660808b0152948a1660a08a015292891660c0890152971660e08701529515156101008601529415156101208501529315156101408401529215156101608301529115156101808201529051908190036101a00190f35b34801561034457600080fd5b50610275610a2b565b34801561035957600080fd5b50610275610a31565b34801561036e57600080fd5b50610275610a37565b34801561038357600080fd5b50610275610a43565b34801561039857600080fd5b50610275600435610a49565b6103c36004351515600160a060020a0360243516604435606435610a5b565b604080519115158252519081900360200190f35b3480156103e357600080fd5b5061040f6004803590602480359160443591606435808201929081013591608435908101910135610b31565b005b34801561041d57600080fd5b506103c36004351515600160a060020a0360243516604435606435610d89565b34801561044957600080fd5b50610452610e4a565b60408051600160a060020a039092168252519081900360200190f35b34801561047a57600080fd5b50610275610e5e565b34801561048f57600080fd5b506103c3610e64565b3480156104a457600080fd5b5061040f60048035906024803580820192908101359160443580820192908101359160643591608435918201910135611540565b3480156104e457600080fd5b506102756117a5565b3480156104f957600080fd5b506103c36117ab565b34801561050e57600080fd5b506102756117c7565b34801561052357600080fd5b506102756004356117cd565b34801561053b57600080fd5b506103c36117df565b34801561055057600080fd5b506102756004356117e8565b34801561056857600080fd5b506102756117fa565b34801561057d57600080fd5b50610275611806565b6103c360043560243560443561180d565b3480156105a357600080fd5b50610275611929565b6103c360043560243560443561192f565b3480156105c957600080fd5b50610275611b47565b3480156105de57600080fd5b506105ea600435611b4c565b604080516001604060020a03909c168c5299151560208c01529715158a8a015295151560608a015293151560808901526001608060020a0390921660a0880152600160a060020a0390811660c08801521660e086015261010085015261012084015261014083015251908190036101600190f35b34801561066a57600080fd5b50610275611bfd565b34801561067f57600080fd5b5061068b600435611c03565b6040805195151586526001604060020a0394851660208701529284168584015292166060840152600160a060020a039091166080830152519081900360a00190f35b3480156106d957600080fd5b506106e2611c67565b604051808260028111156106f257fe5b60ff16815260200191505060405180910390f35b34801561071257600080fd5b50610275611c77565b34801561072757600080fd5b506103c3600160a060020a0360043581169060243516611c7d565b34801561074e57600080fd5b506103c3600435611d21565b34801561076657600080fd5b50610772600435611d2d565b6040805160208082528351818301528351919283929083019185019080838360005b838110156107ac578181015183820152602001610794565b50505050905090810190601f1680156107d95780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156107f357600080fd5b50610452600160a060020a0360043516611e64565b6103c3600160a060020a0360043516602435604435606435611e7f565b34801561083157600080fd5b50610452611f43565b6103c3600435602435604435611f48565b34801561085757600080fd5b506103c36121a5565b34801561086c57600080fd5b50610275612230565b34801561088157600080fd5b5061068b600435612236565b34801561089957600080fd5b50610275600435612244565b3480156108b157600080fd5b506103c36004356024351515612256565b3480156108ce57600080fd5b506105ea6004356122a8565b3480156108e657600080fd5b506108f56004356024356122b6565b604080516001604060020a039d8e1681529b8d1660208d0152998c168b8b015297909a1660608a0152608089019590955260a088019390935260c0870191909152151560e0860152151561010085015215156101208401529215156101408301529115156101608201529051908190036101800190f35b34801561097857600080fd5b50610275612342565b67016345785d8a000081565b600181565b60086020908152600092835260408084209091529082529020805460018201546002909201546001604060020a0380831693604060020a808504831694608060020a80820485169560c060020a92839004861695858116959485048116949283048116939092049091169060ff80821691610100810482169162010000820481169163010000008104821691640100000000909104168d565b600e5481565b60015481565b670c7d713b49da000081565b60115481565b60026020526000908152604090205481565b60008034610a72600a600c8989858a8a8980612348565b604080518281526000602082015281519294507f6940a01870e576ceb735867e13863646d517ce10e66c0133186a4ebdfe9388c2929081900390910190a160408051838152336020820152600160a060020a03881681830152606081018390526080810187905260a0810186905288151560c0820152600060e0820181905261010082015290517f9d57b50c5371c1c3fc64a8947cec60dbae09432e1e5d9ef048317ad7240353e3918190036101200190a15060019695505050505050565b600087815260076020908152604080832089845290915281206004810154909190819060ff161515610b6257600080fd5b6004830154640100000000900460ff161515610b7d57600080fd5b506004820154610100900460ff168015610c6b578254600d8054610c3292869291608060020a9091046001604060020a0316908110610bb857fe5b9060005260206000209060020201600b8b8b8b8080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050508a8a8080601f016020809104026020016040519081016040528093929190818152602001838380828437506128d8945050505050565b60405190925033906000906702c68af0bb1400009082818181858883f19350505050158015610c65573d6000803e3d6000fd5b50610d41565b8254600c8054610d0c92869291608060020a9091046001604060020a0316908110610c9257fe5b9060005260206000209060020201600a8b8b8b8080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050508a8a8080601f016020809104026020016040519081016040528093929190818152602001838380828437506128d8945050505050565b604051909250339060009067016345785d8a00009082818181858883f19350505050158015610d3f573d6000803e3d6000fd5b505b60408051838152821515602082015281517fc8135db115644ed4ae193313c4c801235ef740d2a57a8d5e6fe26ab66635698a929181900390910190a150505050505050505050565b600080806702c68af0bb14000034811115610da357600080fd5b6702c68af0bb14000034039150610dc4600b600d8a8a868b8b600180612348565b60408051828152336020820152600160a060020a038a1681830152606081018590526080810189905260a081018890528a151560c0820152600160e0820181905261010082015290519194507f9d57b50c5371c1c3fc64a8947cec60dbae09432e1e5d9ef048317ad7240353e391908190036101200190a1506001979650505050505050565b6000546101009004600160a060020a031681565b60095481565b600e54600f5460008281526005602052604081205490929183918290819081908190819081908190881115610e9857600080fd5b60008a81526007602090815260408083208b8452825280832080548e855260088452828520604060020a9091046001604060020a0316808652935292206004830154919b5091975090955060ff161515610f9c575b60008a81526008602090815260408083208c845290915290206002015462010000900460ff161515610f545760008a81526008602090815260408083208c8452909152902060020154610100900460ff161515610f4957600080fd5b886001019850610eed565b935460008a81526008602090815260408083208c845282528083208d845260078352818420608060020a9095046001604060020a031680855294909252909120919850909550935b6004860154640100000000900460ff161515610fb757600080fd5b6004860154610100900460ff161561128057601154600b549097508710610fdd57600080fd5b600b805488908110610feb57fe5b90600052602060002090600602019350600d8660000160109054906101000a90046001604060020a03166001604060020a031681548110151561102a57fe5b600091825260209091206002909102018054909350608860020a90046001604060020a031687141561109d57600185015460006001604060020a03909116118015611088575060018501546000196001604060020a03918216011688145b156110955760018a01600e555b60018801600f555b600187016011558354604060020a900460ff1680156110c557508354605860020a900460ff16155b1561121d57604080516101608101825285546001604060020a0381168252604060020a810460ff908116151560208401526901000000000000000000820481161515938301939093526a01000000000000000000008104831615156060830152605860020a8104909216151560808201526c010000000000000000000000009091046001608060020a031660a08201526001850154600160a060020a0390811660c083015260028601541660e08201526003850154610100820152600485015461012082015260058501546101408201526111a090886129eb565b506001840154604051600160a060020a03909116906000906702c68af0bb1400009082818181858883f193505050501580156111e0573d6000803e3d6000fd5b50604080518881526001602082015281517f6940a01870e576ceb735867e13863646d517ce10e66c0133186a4ebdfe9388c2929181900390910190a15b83546aff0000000000000000000019166a0100000000000000000000178455604080518881526001602082015281517f134017cf3262b18f892ee95dde3b0aec9a80cc70a7c96f09c64bd237aceb0473929181900390910190a160019a50611533565b601054600a54909750871061129457600080fd5b600a8054889081106112a257fe5b90600052602060002090600602019150600c8660000160109054906101000a90046001604060020a03166001604060020a03168154811015156112e157fe5b600091825260209091206002909102018054909150608860020a90046001604060020a031687141561135457600185015460006001604060020a0390911611801561133f575060018501546000196001604060020a03918216011688145b1561134c5760018a01600e555b60018801600f555b600187016010558154604060020a900460ff16801561137c57508154605860020a900460ff16155b156114d457604080516101608101825283546001604060020a0381168252604060020a810460ff908116151560208401526901000000000000000000820481161515938301939093526a01000000000000000000008104831615156060830152605860020a8104909216151560808201526c010000000000000000000000009091046001608060020a031660a08201526001830154600160a060020a0390811660c083015260028401541660e082015260038301546101008201526004830154610120820152600583015461014082015261145790886129eb565b506001820154604051600160a060020a039091169060009067016345785d8a00009082818181858883f19350505050158015611497573d6000803e3d6000fd5b50604080518881526000602082015281517f6940a01870e576ceb735867e13863646d517ce10e66c0133186a4ebdfe9388c2929181900390910190a15b81546aff0000000000000000000019166a0100000000000000000000178255604080518881526000602082015281517f134017cf3262b18f892ee95dde3b0aec9a80cc70a7c96f09c64bd237aceb0473929181900390910190a160019a505b5050505050505050505090565b60008061154b614fd2565b60015460009081526007602090815260408083208e84529091529020600481015490935060ff161561157c57600080fd5b8254426001604060020a0360c060020a90920491909116600101116115a057600080fd5b6004830154610100900460ff16156115f7578254600d80549091608060020a90046001604060020a03169081106115d357fe5b6000918252602090912060016002909202010154600160a060020a03169150611638565b8254600c80549091608060020a90046001604060020a031690811061161857fe5b6000918252602090912060016002909202010154600160a060020a031691505b61167188888080601f01602080910402602001604051908101604052809392919081815260200183838082843750612b16945050505050565b905061167c81612c5f565b151561168757600080fd5b81600160a060020a031663f7e498f684600201548c8c8c8c8c8c8c6040518963ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180896000191660001916815260200180602001806020018681526020018060200184810384528b8b8281815260200192508082843790910185810384528981526020019050898980828437909101858103835286815260209081019150879087028082843782019150509b505050505050505050505050602060405180830381600087803b15801561176157600080fd5b505af1158015611775573d6000803e3d6000fd5b505050506040513d602081101561178b57600080fd5b5051151561179857600080fd5b5050505050505050505050565b60105481565b60006117b5612c88565b15156117c057600080fd5b5060015b90565b60045481565b60066020526000908152604090205481565b60005460ff1681565b60036020526000908152604090205481565b6702c68af0bb14000081565b620186a081565b60008054819081906101009004600160a060020a0316331461182e57600080fd5b60008060005460a860020a900460ff16600281111561184957fe5b1461185357600080fd5b67016345785d8a00003481111561186957600080fd5b611871612c88565b50600154600090815260086020908152604080832060045484529091529020600281015490945062010000900460ff16156118ab57600080fd5b6118bb8888886000806000612dee565b600154604080519182526020820183905280519295507fd85b68c66f9db94728c533f39125381246238815483496c57a62b56f70dea8f192918290030190a1835460c060020a90046001604060020a031683141561191b5761191b613049565b506001979650505050505050565b6103e881565b6000808080808060028060005460a860020a900460ff16600281111561195157fe5b1461195b57600080fd5b670c7d713b49da00003481111561197157600080fd5b6001805460009081526007602090815260408083206005835281842054845290915290206004015460ff161597506119b0908c908c908c90808c612dee565b95508515611b3457600180546000908152600860209081526040808320600454845282529182902082516101a08101845281546001604060020a038082168352604060020a808304821695840195909552608060020a80830482169684019690965260c060020a9182900481166060840152958301548087166080840152938404861660a0830152938304851660c08201529290910490921660e0820152600282015460ff808216151561010080850191909152820481161515610120840152620100008204811615156101408401526301000000820481161515610160840152640100000000909104161515610180820152909550611aaf906134bf565b855460018054600090815260056020526040902054929650608060020a9091046001604060020a031690910301925082841415611aee57611aee6134f2565b600154604080519182526020820188905280517f041d165424b249ff6b188265c7ea6f322c4f47171efbac0768888b5fa3ac13d59281900390910190a160019750611b39565b600097505b505050505050509392505050565b600381565b600a805482908110611b5a57fe5b60009182526020909120600690910201805460018201546002830154600384015460048501546005909501546001604060020a0385169650604060020a850460ff90811696690100000000000000000087048216966a01000000000000000000008104831696605860020a8204909316956c010000000000000000000000009091046001608060020a031694600160a060020a039384169493909116929091908b565b600a5490565b600d805482908110611c1157fe5b60009182526020909120600290910201805460019091015460ff821692506001604060020a03610100830481169269010000000000000000008104821692608860020a90910490911690600160a060020a031685565b60005460a860020a900460ff1681565b610e1081565b600080546101009004600160a060020a03163314611c9a57600080fd5b611cac83600160a060020a03166136de565b1515611cb757600080fd5b600160a060020a038381166000908152601260205260409020541615611cdc57600080fd5b50600160a060020a038281166000908152601260205260409020805473ffffffffffffffffffffffffffffffffffffffff191691831691909117905560015b92915050565b6001548114155b919050565b60606000600a83815481101515611d4057fe5b600091825260208083206006929092029091016002810154600160a060020a039081168085526012845260408086205481516101608101835285546001604060020a0381168252604060020a810460ff9081161515988301989098526901000000000000000000810488161515938201939093526a01000000000000000000008304871615156060820152605860020a8304909616151560808701526c010000000000000000000000009091046001608060020a031660a08601526001840154831660c086015260e0850191909152600383015461010085015260048301546101208501526005830154610140850152919450611e5b93611e569388939192611e4992166136e6565b919063ffffffff61377716565b6137cf565b91505b50919050565b601260205260009081526040902054600160a060020a031681565b6000808067016345785d8a000034811115611e9957600080fd5b841515611ea557600080fd5b869150611ebe600a600c60008b868b8b60016000612348565b60408051828152336020820152600160a060020a038b1681830152606081018590526080810189905260a08101889052600060c08201819052600160e083015261010082015290519194507f9d57b50c5371c1c3fc64a8947cec60dbae09432e1e5d9ef048317ad7240353e391908190036101200190a1506001979650505050505050565b600081565b600080600080600080606060008060019054906101000a9004600160a060020a0316600160a060020a031633600160a060020a0316141515611f8957600080fd5b60018060005460a860020a900460ff166002811115611fa457fe5b14611fae57600080fd5b67016345785d8a000034811115611fc457600080fd5b611fcc612c88565b50600154600090815260086020908152604080832060045484529091529020600281015490995062010000900460ff16151561200757600080fd5b6120178d8d8d6001600080612dee565b60015460408051918252602082018390528051929a507f041d165424b249ff6b188265c7ea6f322c4f47171efbac0768888b5fa3ac13d592918290030190a160005460ff1615156121715760015460009081526007602090815260408083208b8452909152902054600c80549091608060020a90046001604060020a031690811061209e57fe5b600091825260209182902060029190910201805460408051690100000000000000000083046001604060020a03908116608860020a909404168381036001018083528087028301909601909252929a50909850965090801561210a578160200160208202803883390190505b5093508592505b84831161215d57600a80548490811061212657fe5b9060005260206000209060060201600501548487850381518110151561214857fe5b60209081029091010152600190920191612111565b8b61216785613985565b1461217157600080fd5b885460c060020a90046001604060020a0316881415612192576121926134f2565b5060019c9b505050505050505050505050565b600080546101009004600160a060020a031633146121c257600080fd5b60028060005460a860020a900460ff1660028111156121dd57fe5b14156121e857600080fd5b6000805475ff00000000000000000000000000000000000000000019167502000000000000000000000000000000000000000000179055612227613bd2565b600191505b5090565b600c5490565b600c805482908110611c1157fe5b60056020526000908152604090205481565b6000811561227157600b80548490811061226c57fe5b506000525b600a80548490811061227f57fe5b60009182526020909120600690910201546a0100000000000000000000900460ff169392505050565b600b805482908110611b5a57fe5b6007602090815260009283526040808420909152908252902080546001820154600283015460038401546004909401546001604060020a0380851695604060020a8604821695608060020a810483169560c060020a90910490921693919260ff80821691610100810482169162010000820481169163010000008104821691640100000000909104168c565b600f5481565b6000806000808a806123735750600160a060020a038a81166000908152601260205260409020541615155b151561237e57600080fd5b8a156123a6578815801590612391575087155b801561239b575086155b15156123a657600080fd5b8c546123b58e6001830161503a565b93508c848154811015156123c557fe5b90600052602060002090600602019250338360010160006101000a815481600160a060020a030219169083600160a060020a031602179055508883600001600c6101000a8154816001608060020a0302191690836001608060020a03160217905550898360020160006101000a815481600160a060020a030219169083600160a060020a031602179055508783600301816000191690555086836004018160001916905550428360000160006101000a8154816001604060020a0302191690836001604060020a03160217905550858360000160086101000a81548160ff0219169083151502179055508a8360000160096101000a81548160ff021916908315150217905550851580156124d757508a155b156125bd57604080516101608101825284546001604060020a0381168252604060020a810460ff908116151560208401526901000000000000000000820481161515938301939093526a01000000000000000000008104831615156060830152605860020a8104909216151560808201526c010000000000000000000000009091046001608060020a031660a08201526001840154600160a060020a0390811660c083015260028501541660e08201526003840154610100820152600484015461012082015260058401546101408201526125b290856129eb565b15156125bd57600080fd5b8b5415156125de578b546125d48d60018301615066565b50600091506125e7565b8b546000190191505b8b828154811015156125f557fe5b60009182526020909120600290910201805490915060ff16806126435750805460016001604060020a03690100000000000000000083048116608860020a90930481169290920301166103e8145b156126a5578b548c906126598260018301615066565b8154811061266357fe5b60009182526020909120600290910201805470ffffffffffffffff000000000000000000191669010000000000000000006001604060020a0387160217815590505b6126ae81613da4565b805478ffffffffffffffff00000000000000000000000000000000001916608860020a6001604060020a038616021781558a156127d557604080516101608101825284546001604060020a0381168252604060020a810460ff908116151560208401526901000000000000000000820481161515938301939093526a01000000000000000000008104831615156060830152605860020a8104909216151560808201526c010000000000000000000000009091046001608060020a031660a08201526001840154600160a060020a0390811660c083015260028501541660e08201526003840154610100820152600484015461012082015260058401546101408201526127d09084906127c1908d6136e6565b8391908763ffffffff613da716565b6128c8565b600160a060020a038a81166000908152601260209081526040918290205482516101608101845287546001604060020a0381168252604060020a810460ff9081161515948301949094526901000000000000000000810484161515948201949094526a01000000000000000000008404831615156060820152605860020a8404909216151560808301526c010000000000000000000000009092046001608060020a031660a08201526001860154831660c08201526002860154831660e08201526003860154610100820152600486015461012082015260058601546101408201526128c89286926127c19291166136e6565b5050509998505050505050505050565b84546001604060020a036901000000000000000000820481168501916000918291608860020a90041683111561290d57600080fd5b846040518082805190602001908083835b6020831061293d5780518252601f19909201916020918201910161291e565b6001836020036101000a0380198251168184511680821785525050505050509050019150506040518091039020915061297585613dc9565b1561297f57600080fd5b60005460ff1615156129a65761299b82878b6003015487613def565b15156129a657600080fd5b86838154811015156129b457fe5b6000918252602090912060069091020180546bff00000000000000000000001916605860020a178155905050509695505050505050565b6000826040015115612a4b578260e00151600160a060020a03166108fc8460a001516001608060020a03169081150290604051600060405180830381858888f19350505050158015612a41573d6000803e3d6000fd5b5060019050611d1b565b60e083015160208085015160c0860151610100870151610120880151604080517fd9afd3a9000000000000000000000000000000000000000000000000000000008152941515600486015260248501899052600160a060020a039384166044860152606485019290925260848401525193169263d9afd3a99260a4808401939192918290030181600087803b158015612ae357600080fd5b505af1158015612af7573d6000803e3d6000fd5b505050506040513d6020811015612b0d57600080fd5b50519392505050565b612b1e614fd2565b6060612b3a6009612b2e85613f48565b9063ffffffff613f6b16565b9050612b5d816000815181101515612b4e57fe5b90602001906020020151613ff9565b6001604060020a031682528051612b7b9082906001908110612b4e57fe5b60208301528051612b939082906002908110612b4e57fe5b6001604060020a031660408301528051612bc39082906003908110612bb457fe5b9060200190602002015161401d565b600160a060020a031660608301528051612be49082906004908110612b4e57fe5b60808301528051612c0b9082906005908110612bfc57fe5b90602001906020020151614042565b60a08301528051612c239082906006908110612b4e57fe5b60c08301528051612c3b9082906007908110612b4e57fe5b60e08301528051612c539082906008908110612b4e57fe5b61010083015250919050565b60008160c001516000148015612c77575060e0820151155b8015611d1b57505061010001511590565b6000808080600260005460a860020a900460ff166002811115612ca757fe5b1415612cb65760009350612de8565b6001805460009081526006602090815260408083205460059092529091205491019350831115612ce95760009350612de8565b6001546000908152600760209081526040808320868452909152902060048101549092506301000000900460ff1615612d255760009350612de8565b600482015460ff1615612d6e578154426001604060020a0360c060020a909204919091166001011115612d5b5760009350612de8565b612d658284614096565b60019350612de8565b50805460016001604060020a03604060020a90920482160116612d9081614103565b15612db0578154612d6590604060020a90046001604060020a031661424f565b8154426001604060020a0360c060020a909204919091166003011115612dd95760009350612de8565b612de38284614096565b600193505b50505090565b6000806000806000878015612e005750865b8015612e095750855b15612f685760018054612e219163ffffffff6143ba16565b6000818152600260208181526040808420546008835281852081865290925290922090810154929650909450925062010000900460ff168015612e7f57506001820154426001604060020a0360c060020a90920491909116610e1001105b15612ec65760008481526008602090815260408083208684528252808320838155600181018490556002908101805464ffffffffff1916905587845290915281205561303b565b60018481556000858152600860208181526040808420600283528185205485528252808420546000198a0185529282528084208885528252928390209093018054608060020a9092046001604060020a031667ffffffffffffffff1990921682179055815187815292830181905281519097507f18013fce596c7fc273e36aaa4d9ba6f94db4e483239db76e94fe9eb6769df2a89281900390910190a1612f8d565b60018054600090815260056020526040902054612f8a9163ffffffff6143ba16565b94505b5060018054600090815260076020908152604080832088845282528083208085018f9055600281018e9055600381018d905580546001604060020a0342811660c060020a02600160c060020a039092169190911780835560048054909216604060020a026fffffffffffffffff000000000000000019909116178255810180548c15156101000261ff00198f151560ff19909316929092179190911617905593548352600590915290208590555b505050509695505050505050565b60048054600181810192839055805460009081526008602090815260408083209583528582528083209483529490529283205482546001604060020a0360c060020a92839004811684018116608060020a0277ffffffffffffffff000000000000000000000000000000001990921691909117845560028401805461ff001962ff000019909116620100001716610100179055918301805442909316909102600160c060020a039092169190911790559080613104836143d3565b600283015460ff1615613188578254600160c060020a0377ffffffffffffffff00000000000000000000000000000000196fffffffffffffffff00000000000000001983166001604060020a03938416604060020a0217908116608060020a918290048416600019018416820217918216910490911660c060020a0217835561320c565b600a5483546fffffffffffffffff00000000000000001916604060020a6000199092016001604060020a039081168302919091178086556001926131db92828116919092048216038301166103e8614805565b84546001604060020a03608060020a8204811692909201929092031660c060020a02600160c060020a039091161783555b600080546001919075ff000000000000000000000000000000000000000000191660a860020a8302179055506000546040517f551dc40198cc79684bb69e4931dba4ac16e4598792ee1c0a5000aeea366d7bb69160a860020a900460ff16908082600281111561327857fe5b60ff16815260200191505060405180910390a160045483546002850154604080519384526001604060020a03608060020a84048116602086015260c060020a84048116858301528084166060860152604060020a909304909216608084015260ff16151560a0830152600160c0830152600060e0830152517f99d777cc3a347ecdaa852c6438ff2afa73f4f68b801c1e81050e817fdb6171d2918190036101000190a1600283015460ff1615613335576133306134f2565b6134ba565b604080516101a08101825284546001604060020a038082168352604060020a80830482166020850152608060020a80840483169585019590955260c060020a928390048216606085015260018801548083166080860152908104821660a0850152938404811660c0840152920490911660e0820152600284015460ff808216151561010080850191909152820481161515610120840152620100008204811615156101408401526301000000820481161515610160840152640100000000909104161515610180820152613408906134bf565b9150600090505b81816001604060020a031610156134ba576001805460009081526007602081815260408084208854608060020a908190046001604060020a03908116890181168752918452828620600401805460ff19168817905589870154875487529484528286208a5482900483168901831687529093529320805477ffffffffffffffff000000000000000000000000000000001916604060020a90930484168601909316021790550161340f565b505050565b6000816101000151156134d457506000611d28565b81604001518260600151036001016001604060020a03169050919050565b6004805460019081019182905580546000908152600860209081526040808320858452909152902091811461355d57506001805460009081526008602090815260408083206004546000190184529091529020546001604060020a0360c060020a9091048116909101165b60028201805461010061ff0019909116179055815477ffffffffffffffff000000000000000000000000000000001916608060020a6001604060020a038381169190910291909117808455600954600160c060020a0391821660c060020a9185016000190184168202178555600185018054909216429093160291909117905560008054819075ff000000000000000000000000000000000000000000191660a860020a8202179055506000546040517f551dc40198cc79684bb69e4931dba4ac16e4598792ee1c0a5000aeea366d7bb69160a860020a900460ff16908082600281111561364757fe5b60ff16815260200191505060405180910390a16004548254604080519283526001604060020a03608060020a83048116602085015260c060020a909204909116828201526000606083018190526080830181905260a0830181905260c0830181905260e0830152517f99d777cc3a347ecdaa852c6438ff2afa73f4f68b801c1e81050e817fdb6171d2918190036101000190a15050565b6000903b1190565b6136ee615092565b6020808401511515908201526040808401511580159183019190915260c080850151600160a060020a03169083015260a0808501516001608060020a031690830152610100808501519083015261012080850151908301526137625760e080840151600160a060020a031690820152611d1b565b600160a060020a03821660e082015292915050565b61377f614fd2565b633b9aca006020820152620186a0604082015260e0840151600160a060020a0316606082015260a08401516001608060020a031660808201526137c384848461483f565b60a08201529392505050565b6040805160098082526101408201909252606091829190816020015b60608152602001906001900390816137eb5750508351909150613816906001604060020a031661493c565b81600081518110151561382557fe5b9060200190602002018190525061383f836020015161493c565b81600181518110151561384e57fe5b60209081029091010152604083015161386f906001604060020a031661493c565b81600281518110151561387e57fe5b60209081029091010152606083015161389f90600160a060020a031661494f565b8160038151811015156138ae57fe5b6020908102909101015260808301516138c69061493c565b8160048151811015156138d557fe5b6020908102909101015260a08301516138ed9061497f565b8160058151811015156138fc57fe5b6020908102909101015260c08301516139149061493c565b81600681518110151561392357fe5b6020908102909101015260e083015161393b9061493c565b81600781518110151561394a57fe5b602090810290910101526101008301516139639061493c565b81600881518110151561397257fe5b60209081029091010152611e5b81614a02565b6000606060008351600114156139b5578360008151811015156139a457fe5b906020019060200201519250613bcb565b8351600290600101046040519080825280602002602001820160405280156139e7578160200160208202803883390190505b5091505b8351816001011015613ad5578381815181101515613a0557fe5b906020019060200201518482600101815181101515613a2057fe5b6020908102909101810151604080518084019490945283810191909152805180840382018152606090930190819052825190918291908401908083835b60208310613a7c5780518252601f199092019160209182019101613a5d565b6001836020036101000a038019825116818451168082178552505050505050905001915050604051809103902082600283811515613ab657fe5b04815181101515613ac357fe5b602090810290910101526002016139eb565b83516002900660011415613bbf57835184906000198101908110613af557fe5b90602001906020020151846001865103815181101515613b1157fe5b6020908102909101810151604080518084019490945283810191909152805180840382018152606090930190819052825190918291908401908083835b60208310613b6d5780518252601f199092019160209182019101613b4e565b6001836020036101000a038019825116818451168082178552505050505050905001915050604051809103902082600283811515613ba757fe5b04815181101515613bb457fe5b602090810290910101525b613bc882613985565b92505b5050919050565b60018054600081815260066020908152604080832054600783528184208185528352818420805495870180865260028552838620604060020a9097046001604060020a031696879055600885528386208787529094529190932092949093919290831415613c4e57805467ffffffffffffffff19168155613c9c565b6001805460009081526008602090815260408083206002835281842054845290915290205482546001604060020a03604060020a90920482169092011667ffffffffffffffff199091161781555b60028101805463ff0000001962ff0000199091166201000017166301000000179055600181810180546001604060020a03421660c060020a02600160c060020a03909116179055600b54613cf59163ffffffff614a5c16565b81546fffffffffffffffff00000000000000001916604060020a6001604060020a0392831681029190911777ffffffffffffffff000000000000000000000000000000001916608060020a6001898101851691909102919091178085559092613d6e9282821692048116919091038301166103e8614805565b82546001604060020a03608060020a8204811692909201929092031660c060020a02600160c060020a0390911617905550505050565b50565b613dbb613db683836000613777565b614a73565b600590930192909255505050565b60006060613ddb6004612b2e85613f48565b9050611e5b816000815181101515612b4e57fe5b6000806000808451610200141515613e0657600080fd5b5086905060205b6102008111613f3a57848101519250600287061515613eab57604080516020808201859052818301869052825180830384018152606090920192839052815191929182918401908083835b60208310613e775780518252601f199092019160209182019101613e58565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405180910390209150613f2c565b604080516020808201869052818301859052825180830384018152606090920192839052815191929182918401908083835b60208310613efc5780518252601f199092019160209182019101613edd565b6001836020036101000a038019825116818451168082178552505050505050905001915050604051809103902091505b600287049650602001613e0d565b509390931495945050505050565b613f506150ee565b50805160408051808201909152602092830181529182015290565b6060613f75615105565b600083604051908082528060200260200182016040528015613fb157816020015b613f9e6150ee565b815260200190600190039081613f965790505b509250613fbd85614ae5565b91505b83811015613ff157613fd182614b0a565b8382815181101515613fdf57fe5b60209081029091010152600101613fc0565b505092915050565b600080600061400784614b3c565b90516020919091036101000a9004949350505050565b60008061402983614b3c565b50516c0100000000000000000000000090049392505050565b602081015160609080151561405657611e5e565b806040519080825280601f01601f191660200182016040528015614084578160200160208202803883390190505b509150611e5e83600001518383614b9f565b60048201805464ff000000001916640100000000179055600180546000908152600660209081526040918290208490559154815190815291820183905280517ffb96205e4b3633fd57aa805b26b51ecf528714a10241a4af015929dce86768d99281900390910190a15050565b600080600060045484111561411b5760009250613bcb565b60015460009081526008602090815260408083208784529091529020600281015490925062010000900460ff1615156141575760009250613bcb565b600282015460ff1615614191574260018360010160189054906101000a90046001604060020a03166001604060020a031601119250613bcb565b6001546000908152600560205260409020548254608060020a90046001604060020a031611156141c45760009250613bcb565b5060015460009081526007602090815260408083208454608060020a90046001604060020a0316845290915290206004810154640100000000900460ff16156142105760019250613bcb565b60048101546301000000900460ff161561422d5760009250613bcb565b5442600160c060020a9092046001604060020a03169190910111159392505050565b600154600090815260086020908152604080832084845290915281206002810154909190819062010000900460ff161561428857600080fd5b8254608060020a90046001604060020a031691505b825460c060020a90046001604060020a0316821161432357506001546000908152600760209081526040808320848452909152902060048101546301000000900460ff16806142f65750600481015462010000900460ff165b1561430057614323565b60048101805464ff0000000019166401000000001790556001919091019061429d565b8254600019909201916001604060020a03608060020a9091041682106143b45760018054600090815260066020908152604091829020859055915485548251918252928101879052608060020a9092046001604060020a03168282015260608201849052517f70801d4d63b3da6c19ba7349911f45bed5a99ccdfb51b8138c105872529bebd5916080908290030190a15b50505050565b6000828201838110156143cc57600080fd5b9392505050565b600a5460009015156143f35760028201805460ff19166001179055614801565b50600154600090815260086020908152604080832060045460011901845290915290208054600a54604060020a9091046001604060020a031660001990910114156144485760028201805460ff191660011790555b6001546002141561445857614801565b600282015460ff16156145ef5780548254604060020a9091046001604060020a031667ffffffffffffffff19909116178255600281015460ff16156144d55760018181015490830180546fffffffffffffffff00000000000000001916604060020a928390046001604060020a03169092029190911790556145ea565b604080516101a08101825282546001604060020a038082168352604060020a80830482166020850152608060020a80840483169585019590955260c060020a928390048216606085015260018601548083166080860152908104821660a0850152938404811660c0840152920490911660e0820152600282015460ff8082161515610100808501919091528204811615156101208401526201000082048116151561014084015263010000008204811615156101608401526401000000009091041615156101808201526145a8906134bf565b8160010160089054906101000a90046001604060020a0316018260010160086101000a8154816001604060020a0302191690836001604060020a031602179055505b6147c0565b6001546000908152600360205260409020541515614621576004546001546000908152600360205260409020556147c0565b600281015460ff161561468b578054825467ffffffffffffffff19166001604060020a03604060020a9283900481169190911784556001808401549085018054918490049092169092026fffffffffffffffff0000000000000000199092169190911790556147c0565b8054825460016001604060020a03604060020a9384900481168201811667ffffffffffffffff19909316929092178555604080516101a0810182528554808516825285810485166020830152608060020a80820486169383019390935260c060020a9081900485166060830152928601548085166080830152948504841660a0820152908404831660c082015292041660e0820152600282015460ff80821615156101008481019190915282048116151561012084015262010000820481161515610140840152630100000082048116151561016084015264010000000090910416151561018082015261477e906134bf565b8160010160089054906101000a90046001604060020a0316018260010160086101000a8154816001604060020a0302191690836001604060020a031602179055505b600c546000101561480157600c80546001919060001981019081106147e157fe5b60009182526020909120600290910201805460ff19169115159190911790555b5050565b6000818381151561481257fe5b061561482c57818381151561482357fe5b046001016143cc565b818381151561483757fe5b049392505050565b6060600084604001511561485257614934565b8261487d577fe904e3d90000000000000000000000000000000000000000000000000000000061489f565b7fd9afd3a9000000000000000000000000000000000000000000000000000000005b90508085602001516148b25760006148b5565b60015b60c0870151610100880151610120890151604080517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19909616602087015260ff94909416602486015260448501899052600160a060020a039092166064850152608484015260a4808401919091528151808403909101815260c4909201905291505b509392505050565b6060611d1b61494a83614bdd565b61497f565b60408051741400000000000000000000000000000000000000008318601482015260348101909152606090611e5b815b6060815160011480156149de575081517f7f0000000000000000000000000000000000000000000000000000000000000090839060009081106149be57fe5b90602001015160f860020a900460f860020a02600160f860020a03191611155b156149ea575080611d28565b611d1b6149fc8351608060ff16614d0e565b83614e36565b604080516000808252602082019092526060915b8351811015614a4a57614a40828583815181101515614a3157fe5b90602001906020020151614e36565b9150600101614a16565b613bc86149fc835160c060ff16614d0e565b60008083831115614a6c57600080fd5b5050900390565b60006060614a80836137cf565b9050806040518082805190602001908083835b60208310614ab25780518252601f199092019160209182019101614a93565b5181516020939093036101000a600019018019909116921691909117905260405192018290039091209695505050505050565b614aed615105565b6000614af883614f40565b83519383529092016020820152919050565b614b126150ee565b60208201516000614b2282614fa6565b828452602080850182905292019390910192909252919050565b805180516000918291821a90826080831015614b5e5781945060019350614b97565b60b8831015614b7c5760018660200151039350816001019450614b97565b60b78303905080600187602001510303935080820160010194505b505050915091565b6020601f820104836020840160005b83811015614bca57602081028381015190830152600101614bae565b5060008551602001860152505050505050565b60408051602080825281830190925260609160009183918291849180820161040080388339019050509250856020840152600093505b6020841015614c71578284815181101515614c2a57fe5b60209101015160f860020a90819004027fff000000000000000000000000000000000000000000000000000000000000001615614c6657614c71565b600190930192614c13565b836020036040519080825280601f01601f191660200182016040528015614ca2578160200160208202803883390190505b509150600090505b8151811015614d05578251600185019484918110614cc457fe5b90602001015160f860020a900460f860020a028282815181101515614ce557fe5b906020010190600160f860020a031916908160001a905350600101614caa565b50949350505050565b60608080604060020a8510614d8457604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600e60248201527f696e70757420746f6f206c6f6e67000000000000000000000000000000000000604482015290519081900360640190fd5b604080516001808252818301909252906020808301908038833901905050915060378511614de45783850160f860020a02826000815181101515614dc457fe5b906020010190600160f860020a031916908160001a905350819250613ff1565b614ded85614bdd565b90508381510160370160f860020a02826000815181101515614e0b57fe5b906020010190600160f860020a031916908160001a905350614e2d8282614e36565b95945050505050565b60608060008084518651016040519080825280601f01601f191660200182016040528015614e6e578160200160208202803883390190505b50925060009150600090505b8551811015614ed6578581815181101515614e9157fe5b90602001015160f860020a900460f860020a028383815181101515614eb257fe5b906020010190600160f860020a031916908160001a90535060019182019101614e7a565b5060005b8451811015614f36578481815181101515614ef157fe5b90602001015160f860020a900460f860020a028383815181101515614f1257fe5b906020010190600160f860020a031916908160001a90535060019182019101614eda565b5090949350505050565b8051805160009190821a906080821015614f5d5760009250613bcb565b60b8821080614f78575060c08210158015614f78575060f882105b15614f865760019250613bcb565b60c0821015614f9b5760b51982019250613bcb565b5060f5190192915050565b8051600090811a6080811015614fbf5760019150611e5e565b60b8811015611e5e57607e190192915050565b6101206040519081016040528060006001604060020a031681526020016000815260200160006001604060020a031681526020016000600160a060020a0316815260200160008152602001606081526020016000815260200160008152602001600081525090565b8154818355818111156134ba576006028160060283600052602060002091820191016134ba9190615126565b8154818355818111156134ba576002028160020283600052602060002091820191016134ba91906151a1565b6040805161016081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c0810182905260e081018290526101008101829052610120810182905261014081019190915290565b604080518082019091526000808252602082015290565b6060604051908101604052806151196150ee565b8152602001600081525090565b6117c491905b8082111561222c5780547bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916815560018101805473ffffffffffffffffffffffffffffffffffffffff19908116909155600282018054909116905560006003820181905560048201819055600582015560060161512c565b6117c491905b8082111561222c57805478ffffffffffffffffffffffffffffffffffffffffffffffffff1916815560018101805473ffffffffffffffffffffffffffffffffffffffff191690556002016151a75600a165627a7a7230582098e18ca594a4f922196052b7524aee5b77d44bc0f7d3a062640ad810412673350029`

// DeployRootChain deploys a new Ethereum contract, binding an instance of RootChain to it.
func DeployRootChain(auth *bind.TransactOpts, backend bind.ContractBackend, _development bool, _NRBEpochLength *big.Int, _statesRoot [32]byte, _transactionsRoot [32]byte, _receiptsRoot [32]byte) (common.Address, *types.Transaction, *RootChain, error) {
	parsed, err := abi.JSON(strings.NewReader(RootChainABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RootChainBin), backend, _development, _NRBEpochLength, _statesRoot, _transactionsRoot, _receiptsRoot)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RootChain{RootChainCaller: RootChainCaller{contract: contract}, RootChainTransactor: RootChainTransactor{contract: contract}, RootChainFilterer: RootChainFilterer{contract: contract}}, nil
}

// RootChain is an auto generated Go binding around an Ethereum contract.
type RootChain struct {
	RootChainCaller     // Read-only binding to the contract
	RootChainTransactor // Write-only binding to the contract
	RootChainFilterer   // Log filterer for contract events
}

// RootChainCaller is an auto generated read-only Go binding around an Ethereum contract.
type RootChainCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RootChainTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RootChainTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RootChainFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RootChainFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RootChainSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RootChainSession struct {
	Contract     *RootChain        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RootChainCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RootChainCallerSession struct {
	Contract *RootChainCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// RootChainTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RootChainTransactorSession struct {
	Contract     *RootChainTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// RootChainRaw is an auto generated low-level Go binding around an Ethereum contract.
type RootChainRaw struct {
	Contract *RootChain // Generic contract binding to access the raw methods on
}

// RootChainCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RootChainCallerRaw struct {
	Contract *RootChainCaller // Generic read-only contract binding to access the raw methods on
}

// RootChainTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RootChainTransactorRaw struct {
	Contract *RootChainTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRootChain creates a new instance of RootChain, bound to a specific deployed contract.
func NewRootChain(address common.Address, backend bind.ContractBackend) (*RootChain, error) {
	contract, err := bindRootChain(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RootChain{RootChainCaller: RootChainCaller{contract: contract}, RootChainTransactor: RootChainTransactor{contract: contract}, RootChainFilterer: RootChainFilterer{contract: contract}}, nil
}

// NewRootChainCaller creates a new read-only instance of RootChain, bound to a specific deployed contract.
func NewRootChainCaller(address common.Address, caller bind.ContractCaller) (*RootChainCaller, error) {
	contract, err := bindRootChain(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RootChainCaller{contract: contract}, nil
}

// NewRootChainTransactor creates a new write-only instance of RootChain, bound to a specific deployed contract.
func NewRootChainTransactor(address common.Address, transactor bind.ContractTransactor) (*RootChainTransactor, error) {
	contract, err := bindRootChain(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RootChainTransactor{contract: contract}, nil
}

// NewRootChainFilterer creates a new log filterer instance of RootChain, bound to a specific deployed contract.
func NewRootChainFilterer(address common.Address, filterer bind.ContractFilterer) (*RootChainFilterer, error) {
	contract, err := bindRootChain(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RootChainFilterer{contract: contract}, nil
}

// bindRootChain binds a generic wrapper to an already deployed contract.
func bindRootChain(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RootChainABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RootChain *RootChainRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RootChain.Contract.RootChainCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RootChain *RootChainRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RootChain.Contract.RootChainTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RootChain *RootChainRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RootChain.Contract.RootChainTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RootChain *RootChainCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RootChain.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RootChain *RootChainTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RootChain.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RootChain *RootChainTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RootChain.Contract.contract.Transact(opts, method, params...)
}

// COSTERO is a free data retrieval call binding the contract method 0xd691acd8.
//
// Solidity: function COST_ERO() constant returns(uint256)
func (_RootChain *RootChainCaller) COSTERO(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChain.contract.Call(opts, out, "COST_ERO")
	return *ret0, err
}

// COSTERO is a free data retrieval call binding the contract method 0xd691acd8.
//
// Solidity: function COST_ERO() constant returns(uint256)
func (_RootChain *RootChainSession) COSTERO() (*big.Int, error) {
	return _RootChain.Contract.COSTERO(&_RootChain.CallOpts)
}

// COSTERO is a free data retrieval call binding the contract method 0xd691acd8.
//
// Solidity: function COST_ERO() constant returns(uint256)
func (_RootChain *RootChainCallerSession) COSTERO() (*big.Int, error) {
	return _RootChain.Contract.COSTERO(&_RootChain.CallOpts)
}

// COSTERU is a free data retrieval call binding the contract method 0x8b5172d0.
//
// Solidity: function COST_ERU() constant returns(uint256)
func (_RootChain *RootChainCaller) COSTERU(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChain.contract.Call(opts, out, "COST_ERU")
	return *ret0, err
}

// COSTERU is a free data retrieval call binding the contract method 0x8b5172d0.
//
// Solidity: function COST_ERU() constant returns(uint256)
func (_RootChain *RootChainSession) COSTERU() (*big.Int, error) {
	return _RootChain.Contract.COSTERU(&_RootChain.CallOpts)
}

// COSTERU is a free data retrieval call binding the contract method 0x8b5172d0.
//
// Solidity: function COST_ERU() constant returns(uint256)
func (_RootChain *RootChainCallerSession) COSTERU() (*big.Int, error) {
	return _RootChain.Contract.COSTERU(&_RootChain.CallOpts)
}

// COSTNRB is a free data retrieval call binding the contract method 0x94be3aa5.
//
// Solidity: function COST_NRB() constant returns(uint256)
func (_RootChain *RootChainCaller) COSTNRB(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChain.contract.Call(opts, out, "COST_NRB")
	return *ret0, err
}

// COSTNRB is a free data retrieval call binding the contract method 0x94be3aa5.
//
// Solidity: function COST_NRB() constant returns(uint256)
func (_RootChain *RootChainSession) COSTNRB() (*big.Int, error) {
	return _RootChain.Contract.COSTNRB(&_RootChain.CallOpts)
}

// COSTNRB is a free data retrieval call binding the contract method 0x94be3aa5.
//
// Solidity: function COST_NRB() constant returns(uint256)
func (_RootChain *RootChainCallerSession) COSTNRB() (*big.Int, error) {
	return _RootChain.Contract.COSTNRB(&_RootChain.CallOpts)
}

// COSTORB is a free data retrieval call binding the contract method 0xb2ae9ba8.
//
// Solidity: function COST_ORB() constant returns(uint256)
func (_RootChain *RootChainCaller) COSTORB(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChain.contract.Call(opts, out, "COST_ORB")
	return *ret0, err
}

// COSTORB is a free data retrieval call binding the contract method 0xb2ae9ba8.
//
// Solidity: function COST_ORB() constant returns(uint256)
func (_RootChain *RootChainSession) COSTORB() (*big.Int, error) {
	return _RootChain.Contract.COSTORB(&_RootChain.CallOpts)
}

// COSTORB is a free data retrieval call binding the contract method 0xb2ae9ba8.
//
// Solidity: function COST_ORB() constant returns(uint256)
func (_RootChain *RootChainCallerSession) COSTORB() (*big.Int, error) {
	return _RootChain.Contract.COSTORB(&_RootChain.CallOpts)
}

// COSTURB is a free data retrieval call binding the contract method 0x192adc5b.
//
// Solidity: function COST_URB() constant returns(uint256)
func (_RootChain *RootChainCaller) COSTURB(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChain.contract.Call(opts, out, "COST_URB")
	return *ret0, err
}

// COSTURB is a free data retrieval call binding the contract method 0x192adc5b.
//
// Solidity: function COST_URB() constant returns(uint256)
func (_RootChain *RootChainSession) COSTURB() (*big.Int, error) {
	return _RootChain.Contract.COSTURB(&_RootChain.CallOpts)
}

// COSTURB is a free data retrieval call binding the contract method 0x192adc5b.
//
// Solidity: function COST_URB() constant returns(uint256)
func (_RootChain *RootChainCallerSession) COSTURB() (*big.Int, error) {
	return _RootChain.Contract.COSTURB(&_RootChain.CallOpts)
}

// COSTURBPREPARE is a free data retrieval call binding the contract method 0x033cfbed.
//
// Solidity: function COST_URB_PREPARE() constant returns(uint256)
func (_RootChain *RootChainCaller) COSTURBPREPARE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChain.contract.Call(opts, out, "COST_URB_PREPARE")
	return *ret0, err
}

// COSTURBPREPARE is a free data retrieval call binding the contract method 0x033cfbed.
//
// Solidity: function COST_URB_PREPARE() constant returns(uint256)
func (_RootChain *RootChainSession) COSTURBPREPARE() (*big.Int, error) {
	return _RootChain.Contract.COSTURBPREPARE(&_RootChain.CallOpts)
}

// COSTURBPREPARE is a free data retrieval call binding the contract method 0x033cfbed.
//
// Solidity: function COST_URB_PREPARE() constant returns(uint256)
func (_RootChain *RootChainCallerSession) COSTURBPREPARE() (*big.Int, error) {
	return _RootChain.Contract.COSTURBPREPARE(&_RootChain.CallOpts)
}

// CPCOMPUTATION is a free data retrieval call binding the contract method 0x08c4fff0.
//
// Solidity: function CP_COMPUTATION() constant returns(uint256)
func (_RootChain *RootChainCaller) CPCOMPUTATION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChain.contract.Call(opts, out, "CP_COMPUTATION")
	return *ret0, err
}

// CPCOMPUTATION is a free data retrieval call binding the contract method 0x08c4fff0.
//
// Solidity: function CP_COMPUTATION() constant returns(uint256)
func (_RootChain *RootChainSession) CPCOMPUTATION() (*big.Int, error) {
	return _RootChain.Contract.CPCOMPUTATION(&_RootChain.CallOpts)
}

// CPCOMPUTATION is a free data retrieval call binding the contract method 0x08c4fff0.
//
// Solidity: function CP_COMPUTATION() constant returns(uint256)
func (_RootChain *RootChainCallerSession) CPCOMPUTATION() (*big.Int, error) {
	return _RootChain.Contract.CPCOMPUTATION(&_RootChain.CallOpts)
}

// CPWITHHOLDING is a free data retrieval call binding the contract method 0xb17fa6e9.
//
// Solidity: function CP_WITHHOLDING() constant returns(uint256)
func (_RootChain *RootChainCaller) CPWITHHOLDING(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChain.contract.Call(opts, out, "CP_WITHHOLDING")
	return *ret0, err
}

// CPWITHHOLDING is a free data retrieval call binding the contract method 0xb17fa6e9.
//
// Solidity: function CP_WITHHOLDING() constant returns(uint256)
func (_RootChain *RootChainSession) CPWITHHOLDING() (*big.Int, error) {
	return _RootChain.Contract.CPWITHHOLDING(&_RootChain.CallOpts)
}

// CPWITHHOLDING is a free data retrieval call binding the contract method 0xb17fa6e9.
//
// Solidity: function CP_WITHHOLDING() constant returns(uint256)
func (_RootChain *RootChainCallerSession) CPWITHHOLDING() (*big.Int, error) {
	return _RootChain.Contract.CPWITHHOLDING(&_RootChain.CallOpts)
}

// EROs is a free data retrieval call binding the contract method 0xb443f3cc.
//
// Solidity: function EROs( uint256) constant returns(timestamp uint64, isExit bool, isTransfer bool, finalized bool, challenged bool, value uint128, requestor address, to address, trieKey bytes32, trieValue bytes32, hash bytes32)
func (_RootChain *RootChainCaller) EROs(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Timestamp  uint64
	IsExit     bool
	IsTransfer bool
	Finalized  bool
	Challenged bool
	Value      *big.Int
	Requestor  common.Address
	To         common.Address
	TrieKey    [32]byte
	TrieValue  [32]byte
	Hash       [32]byte
}, error) {
	ret := new(struct {
		Timestamp  uint64
		IsExit     bool
		IsTransfer bool
		Finalized  bool
		Challenged bool
		Value      *big.Int
		Requestor  common.Address
		To         common.Address
		TrieKey    [32]byte
		TrieValue  [32]byte
		Hash       [32]byte
	})
	out := ret
	err := _RootChain.contract.Call(opts, out, "EROs", arg0)
	return *ret, err
}

// EROs is a free data retrieval call binding the contract method 0xb443f3cc.
//
// Solidity: function EROs( uint256) constant returns(timestamp uint64, isExit bool, isTransfer bool, finalized bool, challenged bool, value uint128, requestor address, to address, trieKey bytes32, trieValue bytes32, hash bytes32)
func (_RootChain *RootChainSession) EROs(arg0 *big.Int) (struct {
	Timestamp  uint64
	IsExit     bool
	IsTransfer bool
	Finalized  bool
	Challenged bool
	Value      *big.Int
	Requestor  common.Address
	To         common.Address
	TrieKey    [32]byte
	TrieValue  [32]byte
	Hash       [32]byte
}, error) {
	return _RootChain.Contract.EROs(&_RootChain.CallOpts, arg0)
}

// EROs is a free data retrieval call binding the contract method 0xb443f3cc.
//
// Solidity: function EROs( uint256) constant returns(timestamp uint64, isExit bool, isTransfer bool, finalized bool, challenged bool, value uint128, requestor address, to address, trieKey bytes32, trieValue bytes32, hash bytes32)
func (_RootChain *RootChainCallerSession) EROs(arg0 *big.Int) (struct {
	Timestamp  uint64
	IsExit     bool
	IsTransfer bool
	Finalized  bool
	Challenged bool
	Value      *big.Int
	Requestor  common.Address
	To         common.Address
	TrieKey    [32]byte
	TrieValue  [32]byte
	Hash       [32]byte
}, error) {
	return _RootChain.Contract.EROs(&_RootChain.CallOpts, arg0)
}

// ERUs is a free data retrieval call binding the contract method 0xf4f31de4.
//
// Solidity: function ERUs( uint256) constant returns(timestamp uint64, isExit bool, isTransfer bool, finalized bool, challenged bool, value uint128, requestor address, to address, trieKey bytes32, trieValue bytes32, hash bytes32)
func (_RootChain *RootChainCaller) ERUs(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Timestamp  uint64
	IsExit     bool
	IsTransfer bool
	Finalized  bool
	Challenged bool
	Value      *big.Int
	Requestor  common.Address
	To         common.Address
	TrieKey    [32]byte
	TrieValue  [32]byte
	Hash       [32]byte
}, error) {
	ret := new(struct {
		Timestamp  uint64
		IsExit     bool
		IsTransfer bool
		Finalized  bool
		Challenged bool
		Value      *big.Int
		Requestor  common.Address
		To         common.Address
		TrieKey    [32]byte
		TrieValue  [32]byte
		Hash       [32]byte
	})
	out := ret
	err := _RootChain.contract.Call(opts, out, "ERUs", arg0)
	return *ret, err
}

// ERUs is a free data retrieval call binding the contract method 0xf4f31de4.
//
// Solidity: function ERUs( uint256) constant returns(timestamp uint64, isExit bool, isTransfer bool, finalized bool, challenged bool, value uint128, requestor address, to address, trieKey bytes32, trieValue bytes32, hash bytes32)
func (_RootChain *RootChainSession) ERUs(arg0 *big.Int) (struct {
	Timestamp  uint64
	IsExit     bool
	IsTransfer bool
	Finalized  bool
	Challenged bool
	Value      *big.Int
	Requestor  common.Address
	To         common.Address
	TrieKey    [32]byte
	TrieValue  [32]byte
	Hash       [32]byte
}, error) {
	return _RootChain.Contract.ERUs(&_RootChain.CallOpts, arg0)
}

// ERUs is a free data retrieval call binding the contract method 0xf4f31de4.
//
// Solidity: function ERUs( uint256) constant returns(timestamp uint64, isExit bool, isTransfer bool, finalized bool, challenged bool, value uint128, requestor address, to address, trieKey bytes32, trieValue bytes32, hash bytes32)
func (_RootChain *RootChainCallerSession) ERUs(arg0 *big.Int) (struct {
	Timestamp  uint64
	IsExit     bool
	IsTransfer bool
	Finalized  bool
	Challenged bool
	Value      *big.Int
	Requestor  common.Address
	To         common.Address
	TrieKey    [32]byte
	TrieValue  [32]byte
	Hash       [32]byte
}, error) {
	return _RootChain.Contract.ERUs(&_RootChain.CallOpts, arg0)
}

// MAXREQUESTS is a free data retrieval call binding the contract method 0x93521222.
//
// Solidity: function MAX_REQUESTS() constant returns(uint256)
func (_RootChain *RootChainCaller) MAXREQUESTS(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChain.contract.Call(opts, out, "MAX_REQUESTS")
	return *ret0, err
}

// MAXREQUESTS is a free data retrieval call binding the contract method 0x93521222.
//
// Solidity: function MAX_REQUESTS() constant returns(uint256)
func (_RootChain *RootChainSession) MAXREQUESTS() (*big.Int, error) {
	return _RootChain.Contract.MAXREQUESTS(&_RootChain.CallOpts)
}

// MAXREQUESTS is a free data retrieval call binding the contract method 0x93521222.
//
// Solidity: function MAX_REQUESTS() constant returns(uint256)
func (_RootChain *RootChainCallerSession) MAXREQUESTS() (*big.Int, error) {
	return _RootChain.Contract.MAXREQUESTS(&_RootChain.CallOpts)
}

// NRBEpochLength is a free data retrieval call binding the contract method 0x584e349e.
//
// Solidity: function NRBEpochLength() constant returns(uint256)
func (_RootChain *RootChainCaller) NRBEpochLength(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChain.contract.Call(opts, out, "NRBEpochLength")
	return *ret0, err
}

// NRBEpochLength is a free data retrieval call binding the contract method 0x584e349e.
//
// Solidity: function NRBEpochLength() constant returns(uint256)
func (_RootChain *RootChainSession) NRBEpochLength() (*big.Int, error) {
	return _RootChain.Contract.NRBEpochLength(&_RootChain.CallOpts)
}

// NRBEpochLength is a free data retrieval call binding the contract method 0x584e349e.
//
// Solidity: function NRBEpochLength() constant returns(uint256)
func (_RootChain *RootChainCallerSession) NRBEpochLength() (*big.Int, error) {
	return _RootChain.Contract.NRBEpochLength(&_RootChain.CallOpts)
}

// NULLADDRESS is a free data retrieval call binding the contract method 0xde0ce17d.
//
// Solidity: function NULL_ADDRESS() constant returns(address)
func (_RootChain *RootChainCaller) NULLADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RootChain.contract.Call(opts, out, "NULL_ADDRESS")
	return *ret0, err
}

// NULLADDRESS is a free data retrieval call binding the contract method 0xde0ce17d.
//
// Solidity: function NULL_ADDRESS() constant returns(address)
func (_RootChain *RootChainSession) NULLADDRESS() (common.Address, error) {
	return _RootChain.Contract.NULLADDRESS(&_RootChain.CallOpts)
}

// NULLADDRESS is a free data retrieval call binding the contract method 0xde0ce17d.
//
// Solidity: function NULL_ADDRESS() constant returns(address)
func (_RootChain *RootChainCallerSession) NULLADDRESS() (common.Address, error) {
	return _RootChain.Contract.NULLADDRESS(&_RootChain.CallOpts)
}

// ORBs is a free data retrieval call binding the contract method 0xea7f22a8.
//
// Solidity: function ORBs( uint256) constant returns(submitted bool, epochNumber uint64, requestStart uint64, requestEnd uint64, trie address)
func (_RootChain *RootChainCaller) ORBs(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Submitted    bool
	EpochNumber  uint64
	RequestStart uint64
	RequestEnd   uint64
	Trie         common.Address
}, error) {
	ret := new(struct {
		Submitted    bool
		EpochNumber  uint64
		RequestStart uint64
		RequestEnd   uint64
		Trie         common.Address
	})
	out := ret
	err := _RootChain.contract.Call(opts, out, "ORBs", arg0)
	return *ret, err
}

// ORBs is a free data retrieval call binding the contract method 0xea7f22a8.
//
// Solidity: function ORBs( uint256) constant returns(submitted bool, epochNumber uint64, requestStart uint64, requestEnd uint64, trie address)
func (_RootChain *RootChainSession) ORBs(arg0 *big.Int) (struct {
	Submitted    bool
	EpochNumber  uint64
	RequestStart uint64
	RequestEnd   uint64
	Trie         common.Address
}, error) {
	return _RootChain.Contract.ORBs(&_RootChain.CallOpts, arg0)
}

// ORBs is a free data retrieval call binding the contract method 0xea7f22a8.
//
// Solidity: function ORBs( uint256) constant returns(submitted bool, epochNumber uint64, requestStart uint64, requestEnd uint64, trie address)
func (_RootChain *RootChainCallerSession) ORBs(arg0 *big.Int) (struct {
	Submitted    bool
	EpochNumber  uint64
	RequestStart uint64
	RequestEnd   uint64
	Trie         common.Address
}, error) {
	return _RootChain.Contract.ORBs(&_RootChain.CallOpts, arg0)
}

// PREPARETIMEOUT is a free data retrieval call binding the contract method 0xc2bc88fa.
//
// Solidity: function PREPARE_TIMEOUT() constant returns(uint256)
func (_RootChain *RootChainCaller) PREPARETIMEOUT(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChain.contract.Call(opts, out, "PREPARE_TIMEOUT")
	return *ret0, err
}

// PREPARETIMEOUT is a free data retrieval call binding the contract method 0xc2bc88fa.
//
// Solidity: function PREPARE_TIMEOUT() constant returns(uint256)
func (_RootChain *RootChainSession) PREPARETIMEOUT() (*big.Int, error) {
	return _RootChain.Contract.PREPARETIMEOUT(&_RootChain.CallOpts)
}

// PREPARETIMEOUT is a free data retrieval call binding the contract method 0xc2bc88fa.
//
// Solidity: function PREPARE_TIMEOUT() constant returns(uint256)
func (_RootChain *RootChainCallerSession) PREPARETIMEOUT() (*big.Int, error) {
	return _RootChain.Contract.PREPARETIMEOUT(&_RootChain.CallOpts)
}

// REQUESTGAS is a free data retrieval call binding the contract method 0x8eb288ca.
//
// Solidity: function REQUEST_GAS() constant returns(uint256)
func (_RootChain *RootChainCaller) REQUESTGAS(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChain.contract.Call(opts, out, "REQUEST_GAS")
	return *ret0, err
}

// REQUESTGAS is a free data retrieval call binding the contract method 0x8eb288ca.
//
// Solidity: function REQUEST_GAS() constant returns(uint256)
func (_RootChain *RootChainSession) REQUESTGAS() (*big.Int, error) {
	return _RootChain.Contract.REQUESTGAS(&_RootChain.CallOpts)
}

// REQUESTGAS is a free data retrieval call binding the contract method 0x8eb288ca.
//
// Solidity: function REQUEST_GAS() constant returns(uint256)
func (_RootChain *RootChainCallerSession) REQUESTGAS() (*big.Int, error) {
	return _RootChain.Contract.REQUESTGAS(&_RootChain.CallOpts)
}

// URBs is a free data retrieval call binding the contract method 0xc0e86064.
//
// Solidity: function URBs( uint256) constant returns(submitted bool, epochNumber uint64, requestStart uint64, requestEnd uint64, trie address)
func (_RootChain *RootChainCaller) URBs(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Submitted    bool
	EpochNumber  uint64
	RequestStart uint64
	RequestEnd   uint64
	Trie         common.Address
}, error) {
	ret := new(struct {
		Submitted    bool
		EpochNumber  uint64
		RequestStart uint64
		RequestEnd   uint64
		Trie         common.Address
	})
	out := ret
	err := _RootChain.contract.Call(opts, out, "URBs", arg0)
	return *ret, err
}

// URBs is a free data retrieval call binding the contract method 0xc0e86064.
//
// Solidity: function URBs( uint256) constant returns(submitted bool, epochNumber uint64, requestStart uint64, requestEnd uint64, trie address)
func (_RootChain *RootChainSession) URBs(arg0 *big.Int) (struct {
	Submitted    bool
	EpochNumber  uint64
	RequestStart uint64
	RequestEnd   uint64
	Trie         common.Address
}, error) {
	return _RootChain.Contract.URBs(&_RootChain.CallOpts, arg0)
}

// URBs is a free data retrieval call binding the contract method 0xc0e86064.
//
// Solidity: function URBs( uint256) constant returns(submitted bool, epochNumber uint64, requestStart uint64, requestEnd uint64, trie address)
func (_RootChain *RootChainCallerSession) URBs(arg0 *big.Int) (struct {
	Submitted    bool
	EpochNumber  uint64
	RequestStart uint64
	RequestEnd   uint64
	Trie         common.Address
}, error) {
	return _RootChain.Contract.URBs(&_RootChain.CallOpts, arg0)
}

// Blocks is a free data retrieval call binding the contract method 0xf4f911db.
//
// Solidity: function blocks( uint256,  uint256) constant returns(forkNumber uint64, epochNumber uint64, requestBlockId uint64, timestamp uint64, statesRoot bytes32, transactionsRoot bytes32, receiptsRoot bytes32, isRequest bool, userActivated bool, challenged bool, challenging bool, finalized bool)
func (_RootChain *RootChainCaller) Blocks(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (struct {
	ForkNumber       uint64
	EpochNumber      uint64
	RequestBlockId   uint64
	Timestamp        uint64
	StatesRoot       [32]byte
	TransactionsRoot [32]byte
	ReceiptsRoot     [32]byte
	IsRequest        bool
	UserActivated    bool
	Challenged       bool
	Challenging      bool
	Finalized        bool
}, error) {
	ret := new(struct {
		ForkNumber       uint64
		EpochNumber      uint64
		RequestBlockId   uint64
		Timestamp        uint64
		StatesRoot       [32]byte
		TransactionsRoot [32]byte
		ReceiptsRoot     [32]byte
		IsRequest        bool
		UserActivated    bool
		Challenged       bool
		Challenging      bool
		Finalized        bool
	})
	out := ret
	err := _RootChain.contract.Call(opts, out, "blocks", arg0, arg1)
	return *ret, err
}

// Blocks is a free data retrieval call binding the contract method 0xf4f911db.
//
// Solidity: function blocks( uint256,  uint256) constant returns(forkNumber uint64, epochNumber uint64, requestBlockId uint64, timestamp uint64, statesRoot bytes32, transactionsRoot bytes32, receiptsRoot bytes32, isRequest bool, userActivated bool, challenged bool, challenging bool, finalized bool)
func (_RootChain *RootChainSession) Blocks(arg0 *big.Int, arg1 *big.Int) (struct {
	ForkNumber       uint64
	EpochNumber      uint64
	RequestBlockId   uint64
	Timestamp        uint64
	StatesRoot       [32]byte
	TransactionsRoot [32]byte
	ReceiptsRoot     [32]byte
	IsRequest        bool
	UserActivated    bool
	Challenged       bool
	Challenging      bool
	Finalized        bool
}, error) {
	return _RootChain.Contract.Blocks(&_RootChain.CallOpts, arg0, arg1)
}

// Blocks is a free data retrieval call binding the contract method 0xf4f911db.
//
// Solidity: function blocks( uint256,  uint256) constant returns(forkNumber uint64, epochNumber uint64, requestBlockId uint64, timestamp uint64, statesRoot bytes32, transactionsRoot bytes32, receiptsRoot bytes32, isRequest bool, userActivated bool, challenged bool, challenging bool, finalized bool)
func (_RootChain *RootChainCallerSession) Blocks(arg0 *big.Int, arg1 *big.Int) (struct {
	ForkNumber       uint64
	EpochNumber      uint64
	RequestBlockId   uint64
	Timestamp        uint64
	StatesRoot       [32]byte
	TransactionsRoot [32]byte
	ReceiptsRoot     [32]byte
	IsRequest        bool
	UserActivated    bool
	Challenged       bool
	Challenging      bool
	Finalized        bool
}, error) {
	return _RootChain.Contract.Blocks(&_RootChain.CallOpts, arg0, arg1)
}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() constant returns(uint256)
func (_RootChain *RootChainCaller) CurrentEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChain.contract.Call(opts, out, "currentEpoch")
	return *ret0, err
}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() constant returns(uint256)
func (_RootChain *RootChainSession) CurrentEpoch() (*big.Int, error) {
	return _RootChain.Contract.CurrentEpoch(&_RootChain.CallOpts)
}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() constant returns(uint256)
func (_RootChain *RootChainCallerSession) CurrentEpoch() (*big.Int, error) {
	return _RootChain.Contract.CurrentEpoch(&_RootChain.CallOpts)
}

// CurrentFork is a free data retrieval call binding the contract method 0x183d2d1c.
//
// Solidity: function currentFork() constant returns(uint256)
func (_RootChain *RootChainCaller) CurrentFork(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChain.contract.Call(opts, out, "currentFork")
	return *ret0, err
}

// CurrentFork is a free data retrieval call binding the contract method 0x183d2d1c.
//
// Solidity: function currentFork() constant returns(uint256)
func (_RootChain *RootChainSession) CurrentFork() (*big.Int, error) {
	return _RootChain.Contract.CurrentFork(&_RootChain.CallOpts)
}

// CurrentFork is a free data retrieval call binding the contract method 0x183d2d1c.
//
// Solidity: function currentFork() constant returns(uint256)
func (_RootChain *RootChainCallerSession) CurrentFork() (*big.Int, error) {
	return _RootChain.Contract.CurrentFork(&_RootChain.CallOpts)
}

// Development is a free data retrieval call binding the contract method 0x7b929c27.
//
// Solidity: function development() constant returns(bool)
func (_RootChain *RootChainCaller) Development(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _RootChain.contract.Call(opts, out, "development")
	return *ret0, err
}

// Development is a free data retrieval call binding the contract method 0x7b929c27.
//
// Solidity: function development() constant returns(bool)
func (_RootChain *RootChainSession) Development() (bool, error) {
	return _RootChain.Contract.Development(&_RootChain.CallOpts)
}

// Development is a free data retrieval call binding the contract method 0x7b929c27.
//
// Solidity: function development() constant returns(bool)
func (_RootChain *RootChainCallerSession) Development() (bool, error) {
	return _RootChain.Contract.Development(&_RootChain.CallOpts)
}

// Epochs is a free data retrieval call binding the contract method 0x0a88bd04.
//
// Solidity: function epochs( uint256,  uint256) constant returns(requestStart uint64, requestEnd uint64, startBlockNumber uint64, endBlockNumber uint64, forkedBlockNumber uint64, firstRequestBlockId uint64, limit uint64, timestamp uint64, isEmpty bool, initialized bool, isRequest bool, userActivated bool, finalized bool)
func (_RootChain *RootChainCaller) Epochs(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (struct {
	RequestStart        uint64
	RequestEnd          uint64
	StartBlockNumber    uint64
	EndBlockNumber      uint64
	ForkedBlockNumber   uint64
	FirstRequestBlockId uint64
	Limit               uint64
	Timestamp           uint64
	IsEmpty             bool
	Initialized         bool
	IsRequest           bool
	UserActivated       bool
	Finalized           bool
}, error) {
	ret := new(struct {
		RequestStart        uint64
		RequestEnd          uint64
		StartBlockNumber    uint64
		EndBlockNumber      uint64
		ForkedBlockNumber   uint64
		FirstRequestBlockId uint64
		Limit               uint64
		Timestamp           uint64
		IsEmpty             bool
		Initialized         bool
		IsRequest           bool
		UserActivated       bool
		Finalized           bool
	})
	out := ret
	err := _RootChain.contract.Call(opts, out, "epochs", arg0, arg1)
	return *ret, err
}

// Epochs is a free data retrieval call binding the contract method 0x0a88bd04.
//
// Solidity: function epochs( uint256,  uint256) constant returns(requestStart uint64, requestEnd uint64, startBlockNumber uint64, endBlockNumber uint64, forkedBlockNumber uint64, firstRequestBlockId uint64, limit uint64, timestamp uint64, isEmpty bool, initialized bool, isRequest bool, userActivated bool, finalized bool)
func (_RootChain *RootChainSession) Epochs(arg0 *big.Int, arg1 *big.Int) (struct {
	RequestStart        uint64
	RequestEnd          uint64
	StartBlockNumber    uint64
	EndBlockNumber      uint64
	ForkedBlockNumber   uint64
	FirstRequestBlockId uint64
	Limit               uint64
	Timestamp           uint64
	IsEmpty             bool
	Initialized         bool
	IsRequest           bool
	UserActivated       bool
	Finalized           bool
}, error) {
	return _RootChain.Contract.Epochs(&_RootChain.CallOpts, arg0, arg1)
}

// Epochs is a free data retrieval call binding the contract method 0x0a88bd04.
//
// Solidity: function epochs( uint256,  uint256) constant returns(requestStart uint64, requestEnd uint64, startBlockNumber uint64, endBlockNumber uint64, forkedBlockNumber uint64, firstRequestBlockId uint64, limit uint64, timestamp uint64, isEmpty bool, initialized bool, isRequest bool, userActivated bool, finalized bool)
func (_RootChain *RootChainCallerSession) Epochs(arg0 *big.Int, arg1 *big.Int) (struct {
	RequestStart        uint64
	RequestEnd          uint64
	StartBlockNumber    uint64
	EndBlockNumber      uint64
	ForkedBlockNumber   uint64
	FirstRequestBlockId uint64
	Limit               uint64
	Timestamp           uint64
	IsEmpty             bool
	Initialized         bool
	IsRequest           bool
	UserActivated       bool
	Finalized           bool
}, error) {
	return _RootChain.Contract.Epochs(&_RootChain.CallOpts, arg0, arg1)
}

// FirstEpoch is a free data retrieval call binding the contract method 0x2e7ab948.
//
// Solidity: function firstEpoch( uint256) constant returns(uint256)
func (_RootChain *RootChainCaller) FirstEpoch(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChain.contract.Call(opts, out, "firstEpoch", arg0)
	return *ret0, err
}

// FirstEpoch is a free data retrieval call binding the contract method 0x2e7ab948.
//
// Solidity: function firstEpoch( uint256) constant returns(uint256)
func (_RootChain *RootChainSession) FirstEpoch(arg0 *big.Int) (*big.Int, error) {
	return _RootChain.Contract.FirstEpoch(&_RootChain.CallOpts, arg0)
}

// FirstEpoch is a free data retrieval call binding the contract method 0x2e7ab948.
//
// Solidity: function firstEpoch( uint256) constant returns(uint256)
func (_RootChain *RootChainCallerSession) FirstEpoch(arg0 *big.Int) (*big.Int, error) {
	return _RootChain.Contract.FirstEpoch(&_RootChain.CallOpts, arg0)
}

// FirstFilledORBEpochNumber is a free data retrieval call binding the contract method 0x83bcbcff.
//
// Solidity: function firstFilledORBEpochNumber( uint256) constant returns(uint256)
func (_RootChain *RootChainCaller) FirstFilledORBEpochNumber(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChain.contract.Call(opts, out, "firstFilledORBEpochNumber", arg0)
	return *ret0, err
}

// FirstFilledORBEpochNumber is a free data retrieval call binding the contract method 0x83bcbcff.
//
// Solidity: function firstFilledORBEpochNumber( uint256) constant returns(uint256)
func (_RootChain *RootChainSession) FirstFilledORBEpochNumber(arg0 *big.Int) (*big.Int, error) {
	return _RootChain.Contract.FirstFilledORBEpochNumber(&_RootChain.CallOpts, arg0)
}

// FirstFilledORBEpochNumber is a free data retrieval call binding the contract method 0x83bcbcff.
//
// Solidity: function firstFilledORBEpochNumber( uint256) constant returns(uint256)
func (_RootChain *RootChainCallerSession) FirstFilledORBEpochNumber(arg0 *big.Int) (*big.Int, error) {
	return _RootChain.Contract.FirstFilledORBEpochNumber(&_RootChain.CallOpts, arg0)
}

// Forked is a free data retrieval call binding the contract method 0xce8a2bc2.
//
// Solidity: function forked(_forkNumber uint256) constant returns(result bool)
func (_RootChain *RootChainCaller) Forked(opts *bind.CallOpts, _forkNumber *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _RootChain.contract.Call(opts, out, "forked", _forkNumber)
	return *ret0, err
}

// Forked is a free data retrieval call binding the contract method 0xce8a2bc2.
//
// Solidity: function forked(_forkNumber uint256) constant returns(result bool)
func (_RootChain *RootChainSession) Forked(_forkNumber *big.Int) (bool, error) {
	return _RootChain.Contract.Forked(&_RootChain.CallOpts, _forkNumber)
}

// Forked is a free data retrieval call binding the contract method 0xce8a2bc2.
//
// Solidity: function forked(_forkNumber uint256) constant returns(result bool)
func (_RootChain *RootChainCallerSession) Forked(_forkNumber *big.Int) (bool, error) {
	return _RootChain.Contract.Forked(&_RootChain.CallOpts, _forkNumber)
}

// GetEROBytes is a free data retrieval call binding the contract method 0xd1723a96.
//
// Solidity: function getEROBytes(_requestId uint256) constant returns(out bytes)
func (_RootChain *RootChainCaller) GetEROBytes(opts *bind.CallOpts, _requestId *big.Int) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _RootChain.contract.Call(opts, out, "getEROBytes", _requestId)
	return *ret0, err
}

// GetEROBytes is a free data retrieval call binding the contract method 0xd1723a96.
//
// Solidity: function getEROBytes(_requestId uint256) constant returns(out bytes)
func (_RootChain *RootChainSession) GetEROBytes(_requestId *big.Int) ([]byte, error) {
	return _RootChain.Contract.GetEROBytes(&_RootChain.CallOpts, _requestId)
}

// GetEROBytes is a free data retrieval call binding the contract method 0xd1723a96.
//
// Solidity: function getEROBytes(_requestId uint256) constant returns(out bytes)
func (_RootChain *RootChainCallerSession) GetEROBytes(_requestId *big.Int) ([]byte, error) {
	return _RootChain.Contract.GetEROBytes(&_RootChain.CallOpts, _requestId)
}

// GetNumEROs is a free data retrieval call binding the contract method 0xb540adba.
//
// Solidity: function getNumEROs() constant returns(uint256)
func (_RootChain *RootChainCaller) GetNumEROs(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChain.contract.Call(opts, out, "getNumEROs")
	return *ret0, err
}

// GetNumEROs is a free data retrieval call binding the contract method 0xb540adba.
//
// Solidity: function getNumEROs() constant returns(uint256)
func (_RootChain *RootChainSession) GetNumEROs() (*big.Int, error) {
	return _RootChain.Contract.GetNumEROs(&_RootChain.CallOpts)
}

// GetNumEROs is a free data retrieval call binding the contract method 0xb540adba.
//
// Solidity: function getNumEROs() constant returns(uint256)
func (_RootChain *RootChainCallerSession) GetNumEROs() (*big.Int, error) {
	return _RootChain.Contract.GetNumEROs(&_RootChain.CallOpts)
}

// GetNumORBs is a free data retrieval call binding the contract method 0xea0c73f6.
//
// Solidity: function getNumORBs() constant returns(uint256)
func (_RootChain *RootChainCaller) GetNumORBs(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChain.contract.Call(opts, out, "getNumORBs")
	return *ret0, err
}

// GetNumORBs is a free data retrieval call binding the contract method 0xea0c73f6.
//
// Solidity: function getNumORBs() constant returns(uint256)
func (_RootChain *RootChainSession) GetNumORBs() (*big.Int, error) {
	return _RootChain.Contract.GetNumORBs(&_RootChain.CallOpts)
}

// GetNumORBs is a free data retrieval call binding the contract method 0xea0c73f6.
//
// Solidity: function getNumORBs() constant returns(uint256)
func (_RootChain *RootChainCallerSession) GetNumORBs() (*big.Int, error) {
	return _RootChain.Contract.GetNumORBs(&_RootChain.CallOpts)
}

// GetRequestFinalized is a free data retrieval call binding the contract method 0xf28a7afa.
//
// Solidity: function getRequestFinalized(_requestId uint256, _userActivated bool) constant returns(finalized bool)
func (_RootChain *RootChainCaller) GetRequestFinalized(opts *bind.CallOpts, _requestId *big.Int, _userActivated bool) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _RootChain.contract.Call(opts, out, "getRequestFinalized", _requestId, _userActivated)
	return *ret0, err
}

// GetRequestFinalized is a free data retrieval call binding the contract method 0xf28a7afa.
//
// Solidity: function getRequestFinalized(_requestId uint256, _userActivated bool) constant returns(finalized bool)
func (_RootChain *RootChainSession) GetRequestFinalized(_requestId *big.Int, _userActivated bool) (bool, error) {
	return _RootChain.Contract.GetRequestFinalized(&_RootChain.CallOpts, _requestId, _userActivated)
}

// GetRequestFinalized is a free data retrieval call binding the contract method 0xf28a7afa.
//
// Solidity: function getRequestFinalized(_requestId uint256, _userActivated bool) constant returns(finalized bool)
func (_RootChain *RootChainCallerSession) GetRequestFinalized(_requestId *big.Int, _userActivated bool) (bool, error) {
	return _RootChain.Contract.GetRequestFinalized(&_RootChain.CallOpts, _requestId, _userActivated)
}

// HighestBlockNumber is a free data retrieval call binding the contract method 0xf1f3c46c.
//
// Solidity: function highestBlockNumber( uint256) constant returns(uint256)
func (_RootChain *RootChainCaller) HighestBlockNumber(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChain.contract.Call(opts, out, "highestBlockNumber", arg0)
	return *ret0, err
}

// HighestBlockNumber is a free data retrieval call binding the contract method 0xf1f3c46c.
//
// Solidity: function highestBlockNumber( uint256) constant returns(uint256)
func (_RootChain *RootChainSession) HighestBlockNumber(arg0 *big.Int) (*big.Int, error) {
	return _RootChain.Contract.HighestBlockNumber(&_RootChain.CallOpts, arg0)
}

// HighestBlockNumber is a free data retrieval call binding the contract method 0xf1f3c46c.
//
// Solidity: function highestBlockNumber( uint256) constant returns(uint256)
func (_RootChain *RootChainCallerSession) HighestBlockNumber(arg0 *big.Int) (*big.Int, error) {
	return _RootChain.Contract.HighestBlockNumber(&_RootChain.CallOpts, arg0)
}

// LastAppliedBlockNumber is a free data retrieval call binding the contract method 0xfb788a27.
//
// Solidity: function lastAppliedBlockNumber() constant returns(uint256)
func (_RootChain *RootChainCaller) LastAppliedBlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChain.contract.Call(opts, out, "lastAppliedBlockNumber")
	return *ret0, err
}

// LastAppliedBlockNumber is a free data retrieval call binding the contract method 0xfb788a27.
//
// Solidity: function lastAppliedBlockNumber() constant returns(uint256)
func (_RootChain *RootChainSession) LastAppliedBlockNumber() (*big.Int, error) {
	return _RootChain.Contract.LastAppliedBlockNumber(&_RootChain.CallOpts)
}

// LastAppliedBlockNumber is a free data retrieval call binding the contract method 0xfb788a27.
//
// Solidity: function lastAppliedBlockNumber() constant returns(uint256)
func (_RootChain *RootChainCallerSession) LastAppliedBlockNumber() (*big.Int, error) {
	return _RootChain.Contract.LastAppliedBlockNumber(&_RootChain.CallOpts)
}

// LastAppliedERO is a free data retrieval call binding the contract method 0x65d724bc.
//
// Solidity: function lastAppliedERO() constant returns(uint256)
func (_RootChain *RootChainCaller) LastAppliedERO(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChain.contract.Call(opts, out, "lastAppliedERO")
	return *ret0, err
}

// LastAppliedERO is a free data retrieval call binding the contract method 0x65d724bc.
//
// Solidity: function lastAppliedERO() constant returns(uint256)
func (_RootChain *RootChainSession) LastAppliedERO() (*big.Int, error) {
	return _RootChain.Contract.LastAppliedERO(&_RootChain.CallOpts)
}

// LastAppliedERO is a free data retrieval call binding the contract method 0x65d724bc.
//
// Solidity: function lastAppliedERO() constant returns(uint256)
func (_RootChain *RootChainCallerSession) LastAppliedERO() (*big.Int, error) {
	return _RootChain.Contract.LastAppliedERO(&_RootChain.CallOpts)
}

// LastAppliedERU is a free data retrieval call binding the contract method 0x1f261d59.
//
// Solidity: function lastAppliedERU() constant returns(uint256)
func (_RootChain *RootChainCaller) LastAppliedERU(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChain.contract.Call(opts, out, "lastAppliedERU")
	return *ret0, err
}

// LastAppliedERU is a free data retrieval call binding the contract method 0x1f261d59.
//
// Solidity: function lastAppliedERU() constant returns(uint256)
func (_RootChain *RootChainSession) LastAppliedERU() (*big.Int, error) {
	return _RootChain.Contract.LastAppliedERU(&_RootChain.CallOpts)
}

// LastAppliedERU is a free data retrieval call binding the contract method 0x1f261d59.
//
// Solidity: function lastAppliedERU() constant returns(uint256)
func (_RootChain *RootChainCallerSession) LastAppliedERU() (*big.Int, error) {
	return _RootChain.Contract.LastAppliedERU(&_RootChain.CallOpts)
}

// LastAppliedForkNumber is a free data retrieval call binding the contract method 0x164bc2ae.
//
// Solidity: function lastAppliedForkNumber() constant returns(uint256)
func (_RootChain *RootChainCaller) LastAppliedForkNumber(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChain.contract.Call(opts, out, "lastAppliedForkNumber")
	return *ret0, err
}

// LastAppliedForkNumber is a free data retrieval call binding the contract method 0x164bc2ae.
//
// Solidity: function lastAppliedForkNumber() constant returns(uint256)
func (_RootChain *RootChainSession) LastAppliedForkNumber() (*big.Int, error) {
	return _RootChain.Contract.LastAppliedForkNumber(&_RootChain.CallOpts)
}

// LastAppliedForkNumber is a free data retrieval call binding the contract method 0x164bc2ae.
//
// Solidity: function lastAppliedForkNumber() constant returns(uint256)
func (_RootChain *RootChainCallerSession) LastAppliedForkNumber() (*big.Int, error) {
	return _RootChain.Contract.LastAppliedForkNumber(&_RootChain.CallOpts)
}

// LastFinalizedBlock is a free data retrieval call binding the contract method 0x76e61c36.
//
// Solidity: function lastFinalizedBlock( uint256) constant returns(uint256)
func (_RootChain *RootChainCaller) LastFinalizedBlock(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChain.contract.Call(opts, out, "lastFinalizedBlock", arg0)
	return *ret0, err
}

// LastFinalizedBlock is a free data retrieval call binding the contract method 0x76e61c36.
//
// Solidity: function lastFinalizedBlock( uint256) constant returns(uint256)
func (_RootChain *RootChainSession) LastFinalizedBlock(arg0 *big.Int) (*big.Int, error) {
	return _RootChain.Contract.LastFinalizedBlock(&_RootChain.CallOpts, arg0)
}

// LastFinalizedBlock is a free data retrieval call binding the contract method 0x76e61c36.
//
// Solidity: function lastFinalizedBlock( uint256) constant returns(uint256)
func (_RootChain *RootChainCallerSession) LastFinalizedBlock(arg0 *big.Int) (*big.Int, error) {
	return _RootChain.Contract.LastFinalizedBlock(&_RootChain.CallOpts, arg0)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() constant returns(address)
func (_RootChain *RootChainCaller) Operator(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RootChain.contract.Call(opts, out, "operator")
	return *ret0, err
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() constant returns(address)
func (_RootChain *RootChainSession) Operator() (common.Address, error) {
	return _RootChain.Contract.Operator(&_RootChain.CallOpts)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() constant returns(address)
func (_RootChain *RootChainCallerSession) Operator() (common.Address, error) {
	return _RootChain.Contract.Operator(&_RootChain.CallOpts)
}

// RequestableContracts is a free data retrieval call binding the contract method 0xda0185f8.
//
// Solidity: function requestableContracts( address) constant returns(address)
func (_RootChain *RootChainCaller) RequestableContracts(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RootChain.contract.Call(opts, out, "requestableContracts", arg0)
	return *ret0, err
}

// RequestableContracts is a free data retrieval call binding the contract method 0xda0185f8.
//
// Solidity: function requestableContracts( address) constant returns(address)
func (_RootChain *RootChainSession) RequestableContracts(arg0 common.Address) (common.Address, error) {
	return _RootChain.Contract.RequestableContracts(&_RootChain.CallOpts, arg0)
}

// RequestableContracts is a free data retrieval call binding the contract method 0xda0185f8.
//
// Solidity: function requestableContracts( address) constant returns(address)
func (_RootChain *RootChainCallerSession) RequestableContracts(arg0 common.Address) (common.Address, error) {
	return _RootChain.Contract.RequestableContracts(&_RootChain.CallOpts, arg0)
}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() constant returns(uint8)
func (_RootChain *RootChainCaller) State(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _RootChain.contract.Call(opts, out, "state")
	return *ret0, err
}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() constant returns(uint8)
func (_RootChain *RootChainSession) State() (uint8, error) {
	return _RootChain.Contract.State(&_RootChain.CallOpts)
}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() constant returns(uint8)
func (_RootChain *RootChainCallerSession) State() (uint8, error) {
	return _RootChain.Contract.State(&_RootChain.CallOpts)
}

// ApplyRequest is a paid mutator transaction binding the contract method 0x61d29e83.
//
// Solidity: function applyRequest() returns(success bool)
func (_RootChain *RootChainTransactor) ApplyRequest(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RootChain.contract.Transact(opts, "applyRequest")
}

// ApplyRequest is a paid mutator transaction binding the contract method 0x61d29e83.
//
// Solidity: function applyRequest() returns(success bool)
func (_RootChain *RootChainSession) ApplyRequest() (*types.Transaction, error) {
	return _RootChain.Contract.ApplyRequest(&_RootChain.TransactOpts)
}

// ApplyRequest is a paid mutator transaction binding the contract method 0x61d29e83.
//
// Solidity: function applyRequest() returns(success bool)
func (_RootChain *RootChainTransactorSession) ApplyRequest() (*types.Transaction, error) {
	return _RootChain.Contract.ApplyRequest(&_RootChain.TransactOpts)
}

// ChallengeExit is a paid mutator transaction binding the contract method 0x404f7d66.
//
// Solidity: function challengeExit(_forkNumber uint256, _blockNumber uint256, _index uint256, _receiptData bytes, _proof bytes) returns()
func (_RootChain *RootChainTransactor) ChallengeExit(opts *bind.TransactOpts, _forkNumber *big.Int, _blockNumber *big.Int, _index *big.Int, _receiptData []byte, _proof []byte) (*types.Transaction, error) {
	return _RootChain.contract.Transact(opts, "challengeExit", _forkNumber, _blockNumber, _index, _receiptData, _proof)
}

// ChallengeExit is a paid mutator transaction binding the contract method 0x404f7d66.
//
// Solidity: function challengeExit(_forkNumber uint256, _blockNumber uint256, _index uint256, _receiptData bytes, _proof bytes) returns()
func (_RootChain *RootChainSession) ChallengeExit(_forkNumber *big.Int, _blockNumber *big.Int, _index *big.Int, _receiptData []byte, _proof []byte) (*types.Transaction, error) {
	return _RootChain.Contract.ChallengeExit(&_RootChain.TransactOpts, _forkNumber, _blockNumber, _index, _receiptData, _proof)
}

// ChallengeExit is a paid mutator transaction binding the contract method 0x404f7d66.
//
// Solidity: function challengeExit(_forkNumber uint256, _blockNumber uint256, _index uint256, _receiptData bytes, _proof bytes) returns()
func (_RootChain *RootChainTransactorSession) ChallengeExit(_forkNumber *big.Int, _blockNumber *big.Int, _index *big.Int, _receiptData []byte, _proof []byte) (*types.Transaction, error) {
	return _RootChain.Contract.ChallengeExit(&_RootChain.TransactOpts, _forkNumber, _blockNumber, _index, _receiptData, _proof)
}

// ChallengeNullAddress is a paid mutator transaction binding the contract method 0x6299fb24.
//
// Solidity: function challengeNullAddress(_blockNumber uint256, _key bytes, _txByte bytes, _branchMask uint256, _siblings bytes32[]) returns()
func (_RootChain *RootChainTransactor) ChallengeNullAddress(opts *bind.TransactOpts, _blockNumber *big.Int, _key []byte, _txByte []byte, _branchMask *big.Int, _siblings [][32]byte) (*types.Transaction, error) {
	return _RootChain.contract.Transact(opts, "challengeNullAddress", _blockNumber, _key, _txByte, _branchMask, _siblings)
}

// ChallengeNullAddress is a paid mutator transaction binding the contract method 0x6299fb24.
//
// Solidity: function challengeNullAddress(_blockNumber uint256, _key bytes, _txByte bytes, _branchMask uint256, _siblings bytes32[]) returns()
func (_RootChain *RootChainSession) ChallengeNullAddress(_blockNumber *big.Int, _key []byte, _txByte []byte, _branchMask *big.Int, _siblings [][32]byte) (*types.Transaction, error) {
	return _RootChain.Contract.ChallengeNullAddress(&_RootChain.TransactOpts, _blockNumber, _key, _txByte, _branchMask, _siblings)
}

// ChallengeNullAddress is a paid mutator transaction binding the contract method 0x6299fb24.
//
// Solidity: function challengeNullAddress(_blockNumber uint256, _key bytes, _txByte bytes, _branchMask uint256, _siblings bytes32[]) returns()
func (_RootChain *RootChainTransactorSession) ChallengeNullAddress(_blockNumber *big.Int, _key []byte, _txByte []byte, _branchMask *big.Int, _siblings [][32]byte) (*types.Transaction, error) {
	return _RootChain.Contract.ChallengeNullAddress(&_RootChain.TransactOpts, _blockNumber, _key, _txByte, _branchMask, _siblings)
}

// FinalizeBlock is a paid mutator transaction binding the contract method 0x75395a58.
//
// Solidity: function finalizeBlock() returns(success bool)
func (_RootChain *RootChainTransactor) FinalizeBlock(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RootChain.contract.Transact(opts, "finalizeBlock")
}

// FinalizeBlock is a paid mutator transaction binding the contract method 0x75395a58.
//
// Solidity: function finalizeBlock() returns(success bool)
func (_RootChain *RootChainSession) FinalizeBlock() (*types.Transaction, error) {
	return _RootChain.Contract.FinalizeBlock(&_RootChain.TransactOpts)
}

// FinalizeBlock is a paid mutator transaction binding the contract method 0x75395a58.
//
// Solidity: function finalizeBlock() returns(success bool)
func (_RootChain *RootChainTransactorSession) FinalizeBlock() (*types.Transaction, error) {
	return _RootChain.Contract.FinalizeBlock(&_RootChain.TransactOpts)
}

// MakeERU is a paid mutator transaction binding the contract method 0x40a029a1.
//
// Solidity: function makeERU(_isTransfer bool, _to address, _trieKey bytes32, _trieValue bytes32) returns(success bool)
func (_RootChain *RootChainTransactor) MakeERU(opts *bind.TransactOpts, _isTransfer bool, _to common.Address, _trieKey [32]byte, _trieValue [32]byte) (*types.Transaction, error) {
	return _RootChain.contract.Transact(opts, "makeERU", _isTransfer, _to, _trieKey, _trieValue)
}

// MakeERU is a paid mutator transaction binding the contract method 0x40a029a1.
//
// Solidity: function makeERU(_isTransfer bool, _to address, _trieKey bytes32, _trieValue bytes32) returns(success bool)
func (_RootChain *RootChainSession) MakeERU(_isTransfer bool, _to common.Address, _trieKey [32]byte, _trieValue [32]byte) (*types.Transaction, error) {
	return _RootChain.Contract.MakeERU(&_RootChain.TransactOpts, _isTransfer, _to, _trieKey, _trieValue)
}

// MakeERU is a paid mutator transaction binding the contract method 0x40a029a1.
//
// Solidity: function makeERU(_isTransfer bool, _to address, _trieKey bytes32, _trieValue bytes32) returns(success bool)
func (_RootChain *RootChainTransactorSession) MakeERU(_isTransfer bool, _to common.Address, _trieKey [32]byte, _trieValue [32]byte) (*types.Transaction, error) {
	return _RootChain.Contract.MakeERU(&_RootChain.TransactOpts, _isTransfer, _to, _trieKey, _trieValue)
}

// MapRequestableContractByOperator is a paid mutator transaction binding the contract method 0xcb5d742f.
//
// Solidity: function mapRequestableContractByOperator(_rootchain address, _childchain address) returns(success bool)
func (_RootChain *RootChainTransactor) MapRequestableContractByOperator(opts *bind.TransactOpts, _rootchain common.Address, _childchain common.Address) (*types.Transaction, error) {
	return _RootChain.contract.Transact(opts, "mapRequestableContractByOperator", _rootchain, _childchain)
}

// MapRequestableContractByOperator is a paid mutator transaction binding the contract method 0xcb5d742f.
//
// Solidity: function mapRequestableContractByOperator(_rootchain address, _childchain address) returns(success bool)
func (_RootChain *RootChainSession) MapRequestableContractByOperator(_rootchain common.Address, _childchain common.Address) (*types.Transaction, error) {
	return _RootChain.Contract.MapRequestableContractByOperator(&_RootChain.TransactOpts, _rootchain, _childchain)
}

// MapRequestableContractByOperator is a paid mutator transaction binding the contract method 0xcb5d742f.
//
// Solidity: function mapRequestableContractByOperator(_rootchain address, _childchain address) returns(success bool)
func (_RootChain *RootChainTransactorSession) MapRequestableContractByOperator(_rootchain common.Address, _childchain common.Address) (*types.Transaction, error) {
	return _RootChain.Contract.MapRequestableContractByOperator(&_RootChain.TransactOpts, _rootchain, _childchain)
}

// PrepareToSubmitURB is a paid mutator transaction binding the contract method 0xe6925d08.
//
// Solidity: function prepareToSubmitURB() returns(success bool)
func (_RootChain *RootChainTransactor) PrepareToSubmitURB(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RootChain.contract.Transact(opts, "prepareToSubmitURB")
}

// PrepareToSubmitURB is a paid mutator transaction binding the contract method 0xe6925d08.
//
// Solidity: function prepareToSubmitURB() returns(success bool)
func (_RootChain *RootChainSession) PrepareToSubmitURB() (*types.Transaction, error) {
	return _RootChain.Contract.PrepareToSubmitURB(&_RootChain.TransactOpts)
}

// PrepareToSubmitURB is a paid mutator transaction binding the contract method 0xe6925d08.
//
// Solidity: function prepareToSubmitURB() returns(success bool)
func (_RootChain *RootChainTransactorSession) PrepareToSubmitURB() (*types.Transaction, error) {
	return _RootChain.Contract.PrepareToSubmitURB(&_RootChain.TransactOpts)
}

// StartEnter is a paid mutator transaction binding the contract method 0x3cfb871e.
//
// Solidity: function startEnter(_isTransfer bool, _to address, _trieKey bytes32, _trieValue bytes32) returns(success bool)
func (_RootChain *RootChainTransactor) StartEnter(opts *bind.TransactOpts, _isTransfer bool, _to common.Address, _trieKey [32]byte, _trieValue [32]byte) (*types.Transaction, error) {
	return _RootChain.contract.Transact(opts, "startEnter", _isTransfer, _to, _trieKey, _trieValue)
}

// StartEnter is a paid mutator transaction binding the contract method 0x3cfb871e.
//
// Solidity: function startEnter(_isTransfer bool, _to address, _trieKey bytes32, _trieValue bytes32) returns(success bool)
func (_RootChain *RootChainSession) StartEnter(_isTransfer bool, _to common.Address, _trieKey [32]byte, _trieValue [32]byte) (*types.Transaction, error) {
	return _RootChain.Contract.StartEnter(&_RootChain.TransactOpts, _isTransfer, _to, _trieKey, _trieValue)
}

// StartEnter is a paid mutator transaction binding the contract method 0x3cfb871e.
//
// Solidity: function startEnter(_isTransfer bool, _to address, _trieKey bytes32, _trieValue bytes32) returns(success bool)
func (_RootChain *RootChainTransactorSession) StartEnter(_isTransfer bool, _to common.Address, _trieKey [32]byte, _trieValue [32]byte) (*types.Transaction, error) {
	return _RootChain.Contract.StartEnter(&_RootChain.TransactOpts, _isTransfer, _to, _trieKey, _trieValue)
}

// StartExit is a paid mutator transaction binding the contract method 0xdac30920.
//
// Solidity: function startExit(_to address, _value uint256, _trieKey bytes32, _trieValue bytes32) returns(success bool)
func (_RootChain *RootChainTransactor) StartExit(opts *bind.TransactOpts, _to common.Address, _value *big.Int, _trieKey [32]byte, _trieValue [32]byte) (*types.Transaction, error) {
	return _RootChain.contract.Transact(opts, "startExit", _to, _value, _trieKey, _trieValue)
}

// StartExit is a paid mutator transaction binding the contract method 0xdac30920.
//
// Solidity: function startExit(_to address, _value uint256, _trieKey bytes32, _trieValue bytes32) returns(success bool)
func (_RootChain *RootChainSession) StartExit(_to common.Address, _value *big.Int, _trieKey [32]byte, _trieValue [32]byte) (*types.Transaction, error) {
	return _RootChain.Contract.StartExit(&_RootChain.TransactOpts, _to, _value, _trieKey, _trieValue)
}

// StartExit is a paid mutator transaction binding the contract method 0xdac30920.
//
// Solidity: function startExit(_to address, _value uint256, _trieKey bytes32, _trieValue bytes32) returns(success bool)
func (_RootChain *RootChainTransactorSession) StartExit(_to common.Address, _value *big.Int, _trieKey [32]byte, _trieValue [32]byte) (*types.Transaction, error) {
	return _RootChain.Contract.StartExit(&_RootChain.TransactOpts, _to, _value, _trieKey, _trieValue)
}

// SubmitNRB is a paid mutator transaction binding the contract method 0x93248d02.
//
// Solidity: function submitNRB(_statesRoot bytes32, _transactionsRoot bytes32, _receiptsRoot bytes32) returns(success bool)
func (_RootChain *RootChainTransactor) SubmitNRB(opts *bind.TransactOpts, _statesRoot [32]byte, _transactionsRoot [32]byte, _receiptsRoot [32]byte) (*types.Transaction, error) {
	return _RootChain.contract.Transact(opts, "submitNRB", _statesRoot, _transactionsRoot, _receiptsRoot)
}

// SubmitNRB is a paid mutator transaction binding the contract method 0x93248d02.
//
// Solidity: function submitNRB(_statesRoot bytes32, _transactionsRoot bytes32, _receiptsRoot bytes32) returns(success bool)
func (_RootChain *RootChainSession) SubmitNRB(_statesRoot [32]byte, _transactionsRoot [32]byte, _receiptsRoot [32]byte) (*types.Transaction, error) {
	return _RootChain.Contract.SubmitNRB(&_RootChain.TransactOpts, _statesRoot, _transactionsRoot, _receiptsRoot)
}

// SubmitNRB is a paid mutator transaction binding the contract method 0x93248d02.
//
// Solidity: function submitNRB(_statesRoot bytes32, _transactionsRoot bytes32, _receiptsRoot bytes32) returns(success bool)
func (_RootChain *RootChainTransactorSession) SubmitNRB(_statesRoot [32]byte, _transactionsRoot [32]byte, _receiptsRoot [32]byte) (*types.Transaction, error) {
	return _RootChain.Contract.SubmitNRB(&_RootChain.TransactOpts, _statesRoot, _transactionsRoot, _receiptsRoot)
}

// SubmitORB is a paid mutator transaction binding the contract method 0xe67123c4.
//
// Solidity: function submitORB(_statesRoot bytes32, _transactionsRoot bytes32, _receiptsRoot bytes32) returns(success bool)
func (_RootChain *RootChainTransactor) SubmitORB(opts *bind.TransactOpts, _statesRoot [32]byte, _transactionsRoot [32]byte, _receiptsRoot [32]byte) (*types.Transaction, error) {
	return _RootChain.contract.Transact(opts, "submitORB", _statesRoot, _transactionsRoot, _receiptsRoot)
}

// SubmitORB is a paid mutator transaction binding the contract method 0xe67123c4.
//
// Solidity: function submitORB(_statesRoot bytes32, _transactionsRoot bytes32, _receiptsRoot bytes32) returns(success bool)
func (_RootChain *RootChainSession) SubmitORB(_statesRoot [32]byte, _transactionsRoot [32]byte, _receiptsRoot [32]byte) (*types.Transaction, error) {
	return _RootChain.Contract.SubmitORB(&_RootChain.TransactOpts, _statesRoot, _transactionsRoot, _receiptsRoot)
}

// SubmitORB is a paid mutator transaction binding the contract method 0xe67123c4.
//
// Solidity: function submitORB(_statesRoot bytes32, _transactionsRoot bytes32, _receiptsRoot bytes32) returns(success bool)
func (_RootChain *RootChainTransactorSession) SubmitORB(_statesRoot [32]byte, _transactionsRoot [32]byte, _receiptsRoot [32]byte) (*types.Transaction, error) {
	return _RootChain.Contract.SubmitORB(&_RootChain.TransactOpts, _statesRoot, _transactionsRoot, _receiptsRoot)
}

// SubmitURB is a paid mutator transaction binding the contract method 0x97be455d.
//
// Solidity: function submitURB(_statesRoot bytes32, _transactionsRoot bytes32, _receiptsRoot bytes32) returns(success bool)
func (_RootChain *RootChainTransactor) SubmitURB(opts *bind.TransactOpts, _statesRoot [32]byte, _transactionsRoot [32]byte, _receiptsRoot [32]byte) (*types.Transaction, error) {
	return _RootChain.contract.Transact(opts, "submitURB", _statesRoot, _transactionsRoot, _receiptsRoot)
}

// SubmitURB is a paid mutator transaction binding the contract method 0x97be455d.
//
// Solidity: function submitURB(_statesRoot bytes32, _transactionsRoot bytes32, _receiptsRoot bytes32) returns(success bool)
func (_RootChain *RootChainSession) SubmitURB(_statesRoot [32]byte, _transactionsRoot [32]byte, _receiptsRoot [32]byte) (*types.Transaction, error) {
	return _RootChain.Contract.SubmitURB(&_RootChain.TransactOpts, _statesRoot, _transactionsRoot, _receiptsRoot)
}

// SubmitURB is a paid mutator transaction binding the contract method 0x97be455d.
//
// Solidity: function submitURB(_statesRoot bytes32, _transactionsRoot bytes32, _receiptsRoot bytes32) returns(success bool)
func (_RootChain *RootChainTransactorSession) SubmitURB(_statesRoot [32]byte, _transactionsRoot [32]byte, _receiptsRoot [32]byte) (*types.Transaction, error) {
	return _RootChain.Contract.SubmitURB(&_RootChain.TransactOpts, _statesRoot, _transactionsRoot, _receiptsRoot)
}

// RootChainBlockFinalizedIterator is returned from FilterBlockFinalized and is used to iterate over the raw logs and unpacked data for BlockFinalized events raised by the RootChain contract.
type RootChainBlockFinalizedIterator struct {
	Event *RootChainBlockFinalized // Event containing the contract specifics and raw log

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
func (it *RootChainBlockFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RootChainBlockFinalized)
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
		it.Event = new(RootChainBlockFinalized)
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
func (it *RootChainBlockFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RootChainBlockFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RootChainBlockFinalized represents a BlockFinalized event raised by the RootChain contract.
type RootChainBlockFinalized struct {
	ForkNumber  *big.Int
	BlockNumber *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBlockFinalized is a free log retrieval operation binding the contract event 0xfb96205e4b3633fd57aa805b26b51ecf528714a10241a4af015929dce86768d9.
//
// Solidity: e BlockFinalized(forkNumber uint256, blockNumber uint256)
func (_RootChain *RootChainFilterer) FilterBlockFinalized(opts *bind.FilterOpts) (*RootChainBlockFinalizedIterator, error) {

	logs, sub, err := _RootChain.contract.FilterLogs(opts, "BlockFinalized")
	if err != nil {
		return nil, err
	}
	return &RootChainBlockFinalizedIterator{contract: _RootChain.contract, event: "BlockFinalized", logs: logs, sub: sub}, nil
}

// WatchBlockFinalized is a free log subscription operation binding the contract event 0xfb96205e4b3633fd57aa805b26b51ecf528714a10241a4af015929dce86768d9.
//
// Solidity: e BlockFinalized(forkNumber uint256, blockNumber uint256)
func (_RootChain *RootChainFilterer) WatchBlockFinalized(opts *bind.WatchOpts, sink chan<- *RootChainBlockFinalized) (event.Subscription, error) {

	logs, sub, err := _RootChain.contract.WatchLogs(opts, "BlockFinalized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RootChainBlockFinalized)
				if err := _RootChain.contract.UnpackLog(event, "BlockFinalized", log); err != nil {
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

// RootChainERUCreatedIterator is returned from FilterERUCreated and is used to iterate over the raw logs and unpacked data for ERUCreated events raised by the RootChain contract.
type RootChainERUCreatedIterator struct {
	Event *RootChainERUCreated // Event containing the contract specifics and raw log

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
func (it *RootChainERUCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RootChainERUCreated)
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
		it.Event = new(RootChainERUCreated)
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
func (it *RootChainERUCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RootChainERUCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RootChainERUCreated represents a ERUCreated event raised by the RootChain contract.
type RootChainERUCreated struct {
	RequestId *big.Int
	Requestor common.Address
	To        common.Address
	TrieKey   [32]byte
	TrieValue [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterERUCreated is a free log retrieval operation binding the contract event 0xd89c6857ed5778107e858511cdc309642a48c9d1717e813a25156f535acc8d36.
//
// Solidity: e ERUCreated(requestId uint256, requestor address, to address, trieKey bytes32, trieValue bytes32)
func (_RootChain *RootChainFilterer) FilterERUCreated(opts *bind.FilterOpts) (*RootChainERUCreatedIterator, error) {

	logs, sub, err := _RootChain.contract.FilterLogs(opts, "ERUCreated")
	if err != nil {
		return nil, err
	}
	return &RootChainERUCreatedIterator{contract: _RootChain.contract, event: "ERUCreated", logs: logs, sub: sub}, nil
}

// WatchERUCreated is a free log subscription operation binding the contract event 0xd89c6857ed5778107e858511cdc309642a48c9d1717e813a25156f535acc8d36.
//
// Solidity: e ERUCreated(requestId uint256, requestor address, to address, trieKey bytes32, trieValue bytes32)
func (_RootChain *RootChainFilterer) WatchERUCreated(opts *bind.WatchOpts, sink chan<- *RootChainERUCreated) (event.Subscription, error) {

	logs, sub, err := _RootChain.contract.WatchLogs(opts, "ERUCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RootChainERUCreated)
				if err := _RootChain.contract.UnpackLog(event, "ERUCreated", log); err != nil {
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

// RootChainEpochFinalizedIterator is returned from FilterEpochFinalized and is used to iterate over the raw logs and unpacked data for EpochFinalized events raised by the RootChain contract.
type RootChainEpochFinalizedIterator struct {
	Event *RootChainEpochFinalized // Event containing the contract specifics and raw log

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
func (it *RootChainEpochFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RootChainEpochFinalized)
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
		it.Event = new(RootChainEpochFinalized)
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
func (it *RootChainEpochFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RootChainEpochFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RootChainEpochFinalized represents a EpochFinalized event raised by the RootChain contract.
type RootChainEpochFinalized struct {
	ForkNumber       *big.Int
	EpochNumber      *big.Int
	StartBlockNumber *big.Int
	EndBlockNumber   *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterEpochFinalized is a free log retrieval operation binding the contract event 0x70801d4d63b3da6c19ba7349911f45bed5a99ccdfb51b8138c105872529bebd5.
//
// Solidity: e EpochFinalized(forkNumber uint256, epochNumber uint256, startBlockNumber uint256, endBlockNumber uint256)
func (_RootChain *RootChainFilterer) FilterEpochFinalized(opts *bind.FilterOpts) (*RootChainEpochFinalizedIterator, error) {

	logs, sub, err := _RootChain.contract.FilterLogs(opts, "EpochFinalized")
	if err != nil {
		return nil, err
	}
	return &RootChainEpochFinalizedIterator{contract: _RootChain.contract, event: "EpochFinalized", logs: logs, sub: sub}, nil
}

// WatchEpochFinalized is a free log subscription operation binding the contract event 0x70801d4d63b3da6c19ba7349911f45bed5a99ccdfb51b8138c105872529bebd5.
//
// Solidity: e EpochFinalized(forkNumber uint256, epochNumber uint256, startBlockNumber uint256, endBlockNumber uint256)
func (_RootChain *RootChainFilterer) WatchEpochFinalized(opts *bind.WatchOpts, sink chan<- *RootChainEpochFinalized) (event.Subscription, error) {

	logs, sub, err := _RootChain.contract.WatchLogs(opts, "EpochFinalized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RootChainEpochFinalized)
				if err := _RootChain.contract.UnpackLog(event, "EpochFinalized", log); err != nil {
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

// RootChainEpochPreparedIterator is returned from FilterEpochPrepared and is used to iterate over the raw logs and unpacked data for EpochPrepared events raised by the RootChain contract.
type RootChainEpochPreparedIterator struct {
	Event *RootChainEpochPrepared // Event containing the contract specifics and raw log

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
func (it *RootChainEpochPreparedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RootChainEpochPrepared)
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
		it.Event = new(RootChainEpochPrepared)
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
func (it *RootChainEpochPreparedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RootChainEpochPreparedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RootChainEpochPrepared represents a EpochPrepared event raised by the RootChain contract.
type RootChainEpochPrepared struct {
	EpochNumber      *big.Int
	StartBlockNumber *big.Int
	EndBlockNumber   *big.Int
	RequestStart     *big.Int
	RequestEnd       *big.Int
	EpochIsEmpty     bool
	IsRequest        bool
	UserActivated    bool
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterEpochPrepared is a free log retrieval operation binding the contract event 0x99d777cc3a347ecdaa852c6438ff2afa73f4f68b801c1e81050e817fdb6171d2.
//
// Solidity: e EpochPrepared(epochNumber uint256, startBlockNumber uint256, endBlockNumber uint256, requestStart uint256, requestEnd uint256, epochIsEmpty bool, isRequest bool, userActivated bool)
func (_RootChain *RootChainFilterer) FilterEpochPrepared(opts *bind.FilterOpts) (*RootChainEpochPreparedIterator, error) {

	logs, sub, err := _RootChain.contract.FilterLogs(opts, "EpochPrepared")
	if err != nil {
		return nil, err
	}
	return &RootChainEpochPreparedIterator{contract: _RootChain.contract, event: "EpochPrepared", logs: logs, sub: sub}, nil
}

// WatchEpochPrepared is a free log subscription operation binding the contract event 0x99d777cc3a347ecdaa852c6438ff2afa73f4f68b801c1e81050e817fdb6171d2.
//
// Solidity: e EpochPrepared(epochNumber uint256, startBlockNumber uint256, endBlockNumber uint256, requestStart uint256, requestEnd uint256, epochIsEmpty bool, isRequest bool, userActivated bool)
func (_RootChain *RootChainFilterer) WatchEpochPrepared(opts *bind.WatchOpts, sink chan<- *RootChainEpochPrepared) (event.Subscription, error) {

	logs, sub, err := _RootChain.contract.WatchLogs(opts, "EpochPrepared")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RootChainEpochPrepared)
				if err := _RootChain.contract.UnpackLog(event, "EpochPrepared", log); err != nil {
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

// RootChainForkedIterator is returned from FilterForked and is used to iterate over the raw logs and unpacked data for Forked events raised by the RootChain contract.
type RootChainForkedIterator struct {
	Event *RootChainForked // Event containing the contract specifics and raw log

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
func (it *RootChainForkedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RootChainForked)
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
		it.Event = new(RootChainForked)
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
func (it *RootChainForkedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RootChainForkedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RootChainForked represents a Forked event raised by the RootChain contract.
type RootChainForked struct {
	NewFork           *big.Int
	ForkedBlockNumber *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterForked is a free log retrieval operation binding the contract event 0x18013fce596c7fc273e36aaa4d9ba6f94db4e483239db76e94fe9eb6769df2a8.
//
// Solidity: e Forked(newFork uint256, forkedBlockNumber uint256)
func (_RootChain *RootChainFilterer) FilterForked(opts *bind.FilterOpts) (*RootChainForkedIterator, error) {

	logs, sub, err := _RootChain.contract.FilterLogs(opts, "Forked")
	if err != nil {
		return nil, err
	}
	return &RootChainForkedIterator{contract: _RootChain.contract, event: "Forked", logs: logs, sub: sub}, nil
}

// WatchForked is a free log subscription operation binding the contract event 0x18013fce596c7fc273e36aaa4d9ba6f94db4e483239db76e94fe9eb6769df2a8.
//
// Solidity: e Forked(newFork uint256, forkedBlockNumber uint256)
func (_RootChain *RootChainFilterer) WatchForked(opts *bind.WatchOpts, sink chan<- *RootChainForked) (event.Subscription, error) {

	logs, sub, err := _RootChain.contract.WatchLogs(opts, "Forked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RootChainForked)
				if err := _RootChain.contract.UnpackLog(event, "Forked", log); err != nil {
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

// RootChainNRBSubmittedIterator is returned from FilterNRBSubmitted and is used to iterate over the raw logs and unpacked data for NRBSubmitted events raised by the RootChain contract.
type RootChainNRBSubmittedIterator struct {
	Event *RootChainNRBSubmitted // Event containing the contract specifics and raw log

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
func (it *RootChainNRBSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RootChainNRBSubmitted)
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
		it.Event = new(RootChainNRBSubmitted)
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
func (it *RootChainNRBSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RootChainNRBSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RootChainNRBSubmitted represents a NRBSubmitted event raised by the RootChain contract.
type RootChainNRBSubmitted struct {
	Fork        *big.Int
	BlockNumber *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterNRBSubmitted is a free log retrieval operation binding the contract event 0xd85b68c66f9db94728c533f39125381246238815483496c57a62b56f70dea8f1.
//
// Solidity: e NRBSubmitted(fork uint256, blockNumber uint256)
func (_RootChain *RootChainFilterer) FilterNRBSubmitted(opts *bind.FilterOpts) (*RootChainNRBSubmittedIterator, error) {

	logs, sub, err := _RootChain.contract.FilterLogs(opts, "NRBSubmitted")
	if err != nil {
		return nil, err
	}
	return &RootChainNRBSubmittedIterator{contract: _RootChain.contract, event: "NRBSubmitted", logs: logs, sub: sub}, nil
}

// WatchNRBSubmitted is a free log subscription operation binding the contract event 0xd85b68c66f9db94728c533f39125381246238815483496c57a62b56f70dea8f1.
//
// Solidity: e NRBSubmitted(fork uint256, blockNumber uint256)
func (_RootChain *RootChainFilterer) WatchNRBSubmitted(opts *bind.WatchOpts, sink chan<- *RootChainNRBSubmitted) (event.Subscription, error) {

	logs, sub, err := _RootChain.contract.WatchLogs(opts, "NRBSubmitted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RootChainNRBSubmitted)
				if err := _RootChain.contract.UnpackLog(event, "NRBSubmitted", log); err != nil {
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

// RootChainORBSubmittedIterator is returned from FilterORBSubmitted and is used to iterate over the raw logs and unpacked data for ORBSubmitted events raised by the RootChain contract.
type RootChainORBSubmittedIterator struct {
	Event *RootChainORBSubmitted // Event containing the contract specifics and raw log

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
func (it *RootChainORBSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RootChainORBSubmitted)
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
		it.Event = new(RootChainORBSubmitted)
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
func (it *RootChainORBSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RootChainORBSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RootChainORBSubmitted represents a ORBSubmitted event raised by the RootChain contract.
type RootChainORBSubmitted struct {
	Fork        *big.Int
	BlockNumber *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterORBSubmitted is a free log retrieval operation binding the contract event 0x041d165424b249ff6b188265c7ea6f322c4f47171efbac0768888b5fa3ac13d5.
//
// Solidity: e ORBSubmitted(fork uint256, blockNumber uint256)
func (_RootChain *RootChainFilterer) FilterORBSubmitted(opts *bind.FilterOpts) (*RootChainORBSubmittedIterator, error) {

	logs, sub, err := _RootChain.contract.FilterLogs(opts, "ORBSubmitted")
	if err != nil {
		return nil, err
	}
	return &RootChainORBSubmittedIterator{contract: _RootChain.contract, event: "ORBSubmitted", logs: logs, sub: sub}, nil
}

// WatchORBSubmitted is a free log subscription operation binding the contract event 0x041d165424b249ff6b188265c7ea6f322c4f47171efbac0768888b5fa3ac13d5.
//
// Solidity: e ORBSubmitted(fork uint256, blockNumber uint256)
func (_RootChain *RootChainFilterer) WatchORBSubmitted(opts *bind.WatchOpts, sink chan<- *RootChainORBSubmitted) (event.Subscription, error) {

	logs, sub, err := _RootChain.contract.WatchLogs(opts, "ORBSubmitted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RootChainORBSubmitted)
				if err := _RootChain.contract.UnpackLog(event, "ORBSubmitted", log); err != nil {
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

// RootChainRequestAppliedIterator is returned from FilterRequestApplied and is used to iterate over the raw logs and unpacked data for RequestApplied events raised by the RootChain contract.
type RootChainRequestAppliedIterator struct {
	Event *RootChainRequestApplied // Event containing the contract specifics and raw log

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
func (it *RootChainRequestAppliedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RootChainRequestApplied)
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
		it.Event = new(RootChainRequestApplied)
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
func (it *RootChainRequestAppliedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RootChainRequestAppliedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RootChainRequestApplied represents a RequestApplied event raised by the RootChain contract.
type RootChainRequestApplied struct {
	RequestId     *big.Int
	UserActivated bool
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRequestApplied is a free log retrieval operation binding the contract event 0x6940a01870e576ceb735867e13863646d517ce10e66c0133186a4ebdfe9388c2.
//
// Solidity: e RequestApplied(requestId uint256, userActivated bool)
func (_RootChain *RootChainFilterer) FilterRequestApplied(opts *bind.FilterOpts) (*RootChainRequestAppliedIterator, error) {

	logs, sub, err := _RootChain.contract.FilterLogs(opts, "RequestApplied")
	if err != nil {
		return nil, err
	}
	return &RootChainRequestAppliedIterator{contract: _RootChain.contract, event: "RequestApplied", logs: logs, sub: sub}, nil
}

// WatchRequestApplied is a free log subscription operation binding the contract event 0x6940a01870e576ceb735867e13863646d517ce10e66c0133186a4ebdfe9388c2.
//
// Solidity: e RequestApplied(requestId uint256, userActivated bool)
func (_RootChain *RootChainFilterer) WatchRequestApplied(opts *bind.WatchOpts, sink chan<- *RootChainRequestApplied) (event.Subscription, error) {

	logs, sub, err := _RootChain.contract.WatchLogs(opts, "RequestApplied")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RootChainRequestApplied)
				if err := _RootChain.contract.UnpackLog(event, "RequestApplied", log); err != nil {
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

// RootChainRequestChallengedIterator is returned from FilterRequestChallenged and is used to iterate over the raw logs and unpacked data for RequestChallenged events raised by the RootChain contract.
type RootChainRequestChallengedIterator struct {
	Event *RootChainRequestChallenged // Event containing the contract specifics and raw log

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
func (it *RootChainRequestChallengedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RootChainRequestChallenged)
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
		it.Event = new(RootChainRequestChallenged)
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
func (it *RootChainRequestChallengedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RootChainRequestChallengedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RootChainRequestChallenged represents a RequestChallenged event raised by the RootChain contract.
type RootChainRequestChallenged struct {
	RequestId     *big.Int
	UserActivated bool
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRequestChallenged is a free log retrieval operation binding the contract event 0xc8135db115644ed4ae193313c4c801235ef740d2a57a8d5e6fe26ab66635698a.
//
// Solidity: e RequestChallenged(requestId uint256, userActivated bool)
func (_RootChain *RootChainFilterer) FilterRequestChallenged(opts *bind.FilterOpts) (*RootChainRequestChallengedIterator, error) {

	logs, sub, err := _RootChain.contract.FilterLogs(opts, "RequestChallenged")
	if err != nil {
		return nil, err
	}
	return &RootChainRequestChallengedIterator{contract: _RootChain.contract, event: "RequestChallenged", logs: logs, sub: sub}, nil
}

// WatchRequestChallenged is a free log subscription operation binding the contract event 0xc8135db115644ed4ae193313c4c801235ef740d2a57a8d5e6fe26ab66635698a.
//
// Solidity: e RequestChallenged(requestId uint256, userActivated bool)
func (_RootChain *RootChainFilterer) WatchRequestChallenged(opts *bind.WatchOpts, sink chan<- *RootChainRequestChallenged) (event.Subscription, error) {

	logs, sub, err := _RootChain.contract.WatchLogs(opts, "RequestChallenged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RootChainRequestChallenged)
				if err := _RootChain.contract.UnpackLog(event, "RequestChallenged", log); err != nil {
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

// RootChainRequestCreatedIterator is returned from FilterRequestCreated and is used to iterate over the raw logs and unpacked data for RequestCreated events raised by the RootChain contract.
type RootChainRequestCreatedIterator struct {
	Event *RootChainRequestCreated // Event containing the contract specifics and raw log

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
func (it *RootChainRequestCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RootChainRequestCreated)
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
		it.Event = new(RootChainRequestCreated)
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
func (it *RootChainRequestCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RootChainRequestCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RootChainRequestCreated represents a RequestCreated event raised by the RootChain contract.
type RootChainRequestCreated struct {
	RequestId     *big.Int
	Requestor     common.Address
	To            common.Address
	WeiAmount     *big.Int
	TrieKey       [32]byte
	TrieValue     [32]byte
	IsTransfer    bool
	IsExit        bool
	UserActivated bool
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRequestCreated is a free log retrieval operation binding the contract event 0x9d57b50c5371c1c3fc64a8947cec60dbae09432e1e5d9ef048317ad7240353e3.
//
// Solidity: e RequestCreated(requestId uint256, requestor address, to address, weiAmount uint256, trieKey bytes32, trieValue bytes32, isTransfer bool, isExit bool, userActivated bool)
func (_RootChain *RootChainFilterer) FilterRequestCreated(opts *bind.FilterOpts) (*RootChainRequestCreatedIterator, error) {

	logs, sub, err := _RootChain.contract.FilterLogs(opts, "RequestCreated")
	if err != nil {
		return nil, err
	}
	return &RootChainRequestCreatedIterator{contract: _RootChain.contract, event: "RequestCreated", logs: logs, sub: sub}, nil
}

// WatchRequestCreated is a free log subscription operation binding the contract event 0x9d57b50c5371c1c3fc64a8947cec60dbae09432e1e5d9ef048317ad7240353e3.
//
// Solidity: e RequestCreated(requestId uint256, requestor address, to address, weiAmount uint256, trieKey bytes32, trieValue bytes32, isTransfer bool, isExit bool, userActivated bool)
func (_RootChain *RootChainFilterer) WatchRequestCreated(opts *bind.WatchOpts, sink chan<- *RootChainRequestCreated) (event.Subscription, error) {

	logs, sub, err := _RootChain.contract.WatchLogs(opts, "RequestCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RootChainRequestCreated)
				if err := _RootChain.contract.UnpackLog(event, "RequestCreated", log); err != nil {
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

// RootChainRequestFinalizedIterator is returned from FilterRequestFinalized and is used to iterate over the raw logs and unpacked data for RequestFinalized events raised by the RootChain contract.
type RootChainRequestFinalizedIterator struct {
	Event *RootChainRequestFinalized // Event containing the contract specifics and raw log

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
func (it *RootChainRequestFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RootChainRequestFinalized)
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
		it.Event = new(RootChainRequestFinalized)
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
func (it *RootChainRequestFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RootChainRequestFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RootChainRequestFinalized represents a RequestFinalized event raised by the RootChain contract.
type RootChainRequestFinalized struct {
	RequestId     *big.Int
	UserActivated bool
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRequestFinalized is a free log retrieval operation binding the contract event 0x134017cf3262b18f892ee95dde3b0aec9a80cc70a7c96f09c64bd237aceb0473.
//
// Solidity: e RequestFinalized(requestId uint256, userActivated bool)
func (_RootChain *RootChainFilterer) FilterRequestFinalized(opts *bind.FilterOpts) (*RootChainRequestFinalizedIterator, error) {

	logs, sub, err := _RootChain.contract.FilterLogs(opts, "RequestFinalized")
	if err != nil {
		return nil, err
	}
	return &RootChainRequestFinalizedIterator{contract: _RootChain.contract, event: "RequestFinalized", logs: logs, sub: sub}, nil
}

// WatchRequestFinalized is a free log subscription operation binding the contract event 0x134017cf3262b18f892ee95dde3b0aec9a80cc70a7c96f09c64bd237aceb0473.
//
// Solidity: e RequestFinalized(requestId uint256, userActivated bool)
func (_RootChain *RootChainFilterer) WatchRequestFinalized(opts *bind.WatchOpts, sink chan<- *RootChainRequestFinalized) (event.Subscription, error) {

	logs, sub, err := _RootChain.contract.WatchLogs(opts, "RequestFinalized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RootChainRequestFinalized)
				if err := _RootChain.contract.UnpackLog(event, "RequestFinalized", log); err != nil {
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

// RootChainSessionTimeoutIterator is returned from FilterSessionTimeout and is used to iterate over the raw logs and unpacked data for SessionTimeout events raised by the RootChain contract.
type RootChainSessionTimeoutIterator struct {
	Event *RootChainSessionTimeout // Event containing the contract specifics and raw log

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
func (it *RootChainSessionTimeoutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RootChainSessionTimeout)
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
		it.Event = new(RootChainSessionTimeout)
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
func (it *RootChainSessionTimeoutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RootChainSessionTimeoutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RootChainSessionTimeout represents a SessionTimeout event raised by the RootChain contract.
type RootChainSessionTimeout struct {
	UserActivated bool
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSessionTimeout is a free log retrieval operation binding the contract event 0x2122ec719581cd177f225f59a1ee005831211196831f8f1ccffa817d2e7bd108.
//
// Solidity: e SessionTimeout(userActivated bool)
func (_RootChain *RootChainFilterer) FilterSessionTimeout(opts *bind.FilterOpts) (*RootChainSessionTimeoutIterator, error) {

	logs, sub, err := _RootChain.contract.FilterLogs(opts, "SessionTimeout")
	if err != nil {
		return nil, err
	}
	return &RootChainSessionTimeoutIterator{contract: _RootChain.contract, event: "SessionTimeout", logs: logs, sub: sub}, nil
}

// WatchSessionTimeout is a free log subscription operation binding the contract event 0x2122ec719581cd177f225f59a1ee005831211196831f8f1ccffa817d2e7bd108.
//
// Solidity: e SessionTimeout(userActivated bool)
func (_RootChain *RootChainFilterer) WatchSessionTimeout(opts *bind.WatchOpts, sink chan<- *RootChainSessionTimeout) (event.Subscription, error) {

	logs, sub, err := _RootChain.contract.WatchLogs(opts, "SessionTimeout")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RootChainSessionTimeout)
				if err := _RootChain.contract.UnpackLog(event, "SessionTimeout", log); err != nil {
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

// RootChainStateChangedIterator is returned from FilterStateChanged and is used to iterate over the raw logs and unpacked data for StateChanged events raised by the RootChain contract.
type RootChainStateChangedIterator struct {
	Event *RootChainStateChanged // Event containing the contract specifics and raw log

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
func (it *RootChainStateChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RootChainStateChanged)
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
		it.Event = new(RootChainStateChanged)
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
func (it *RootChainStateChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RootChainStateChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RootChainStateChanged represents a StateChanged event raised by the RootChain contract.
type RootChainStateChanged struct {
	State uint8
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterStateChanged is a free log retrieval operation binding the contract event 0x551dc40198cc79684bb69e4931dba4ac16e4598792ee1c0a5000aeea366d7bb6.
//
// Solidity: e StateChanged(state uint8)
func (_RootChain *RootChainFilterer) FilterStateChanged(opts *bind.FilterOpts) (*RootChainStateChangedIterator, error) {

	logs, sub, err := _RootChain.contract.FilterLogs(opts, "StateChanged")
	if err != nil {
		return nil, err
	}
	return &RootChainStateChangedIterator{contract: _RootChain.contract, event: "StateChanged", logs: logs, sub: sub}, nil
}

// WatchStateChanged is a free log subscription operation binding the contract event 0x551dc40198cc79684bb69e4931dba4ac16e4598792ee1c0a5000aeea366d7bb6.
//
// Solidity: e StateChanged(state uint8)
func (_RootChain *RootChainFilterer) WatchStateChanged(opts *bind.WatchOpts, sink chan<- *RootChainStateChanged) (event.Subscription, error) {

	logs, sub, err := _RootChain.contract.WatchLogs(opts, "StateChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RootChainStateChanged)
				if err := _RootChain.contract.UnpackLog(event, "StateChanged", log); err != nil {
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

// RootChainURBSubmittedIterator is returned from FilterURBSubmitted and is used to iterate over the raw logs and unpacked data for URBSubmitted events raised by the RootChain contract.
type RootChainURBSubmittedIterator struct {
	Event *RootChainURBSubmitted // Event containing the contract specifics and raw log

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
func (it *RootChainURBSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RootChainURBSubmitted)
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
		it.Event = new(RootChainURBSubmitted)
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
func (it *RootChainURBSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RootChainURBSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RootChainURBSubmitted represents a URBSubmitted event raised by the RootChain contract.
type RootChainURBSubmitted struct {
	Fork        *big.Int
	BlockNumber *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterURBSubmitted is a free log retrieval operation binding the contract event 0xd8803c978d8d54db576ad0fe876c746dc02c25d13bd7847815860b723304c34f.
//
// Solidity: e URBSubmitted(fork uint256, blockNumber uint256)
func (_RootChain *RootChainFilterer) FilterURBSubmitted(opts *bind.FilterOpts) (*RootChainURBSubmittedIterator, error) {

	logs, sub, err := _RootChain.contract.FilterLogs(opts, "URBSubmitted")
	if err != nil {
		return nil, err
	}
	return &RootChainURBSubmittedIterator{contract: _RootChain.contract, event: "URBSubmitted", logs: logs, sub: sub}, nil
}

// WatchURBSubmitted is a free log subscription operation binding the contract event 0xd8803c978d8d54db576ad0fe876c746dc02c25d13bd7847815860b723304c34f.
//
// Solidity: e URBSubmitted(fork uint256, blockNumber uint256)
func (_RootChain *RootChainFilterer) WatchURBSubmitted(opts *bind.WatchOpts, sink chan<- *RootChainURBSubmitted) (event.Subscription, error) {

	logs, sub, err := _RootChain.contract.WatchLogs(opts, "URBSubmitted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RootChainURBSubmitted)
				if err := _RootChain.contract.UnpackLog(event, "URBSubmitted", log); err != nil {
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

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
const SafeMathBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a72305820f488c403a399e816f111cdb06a588a102aa04f3e383556673ac23fff503673840029`

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
