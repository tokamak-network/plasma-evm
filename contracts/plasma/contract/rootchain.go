// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

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
const AddressBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a72305820e2c15d31406954a71a6485749836de23d45c0bf8fb01c435f387d7b23a9999af0029`

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

// BitsABI is the input ABI used to generate the binding from.
const BitsABI = "[]"

// BitsBin is the compiled bytecode used for deploying new contracts.
const BitsBin = `0x605a602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a265627a7a723058207a644ef047535603cff62c8178055146e8a5c41033b2f6e6f04d09fa09d6d5976c6578706572696d656e74616cf50037`

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
const DataBin = `0x610205610030600b82828239805160001a6073146000811461002057610022565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600436106100835763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416631927ac58811461008857806390e84f56146100a6578063a7b6ae28146100ae578063a89ca766146100c3578063ab73ff05146100cb575b600080fd5b6100906100e0565b60405161009d919061017c565b60405180910390f35b6100906100e5565b6100b66100ec565b60405161009d919061016e565b6100b6610110565b6100d3610134565b60405161009d919061015a565b600181565b620186a081565b7fd9afd3a90000000000000000000000000000000000000000000000000000000081565b7fe904e3d90000000000000000000000000000000000000000000000000000000081565b600081565b6101428161018a565b82525050565b610142816101a3565b610142816101c8565b602081016101688284610139565b92915050565b602081016101688284610148565b602081016101688284610151565b73ffffffffffffffffffffffffffffffffffffffff1690565b7fffffffff000000000000000000000000000000000000000000000000000000001690565b905600a265627a7a723058203e7fa85538a76b4ec6666b8a3a269ba7898be482923a98e33147f252cf7267c56c6578706572696d656e74616cf50037`

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
const MathBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a723058206821729e7a1c1a61076ce529d53dbbe3faa6286298adb57eced5b9334578de930029`

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

// PatriciaTreeABI is the input ABI used to generate the binding from.
const PatriciaTreeABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"key\",\"type\":\"bytes\"},{\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"insert\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"getNode\",\"outputs\":[{\"name\":\"nodes\",\"type\":\"bytes32[2]\"},{\"name\":\"labelDatas\",\"type\":\"bytes32[2]\"},{\"name\":\"labelLengths\",\"type\":\"uint256[2]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"key\",\"type\":\"bytes\"}],\"name\":\"getProof\",\"outputs\":[{\"name\":\"branchMask\",\"type\":\"uint256\"},{\"name\":\"_siblings\",\"type\":\"bytes32[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getRootHash\",\"outputs\":[{\"name\":\"h\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getRootEdge\",\"outputs\":[{\"name\":\"node\",\"type\":\"bytes32\"},{\"name\":\"labelData\",\"type\":\"bytes32\"},{\"name\":\"labelLength\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"rootHash\",\"type\":\"bytes32\"},{\"name\":\"key\",\"type\":\"bytes\"},{\"name\":\"value\",\"type\":\"bytes\"},{\"name\":\"branchMask\",\"type\":\"uint256\"},{\"name\":\"siblings\",\"type\":\"bytes32[]\"}],\"name\":\"verifyProof\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// PatriciaTreeBin is the compiled bytecode used for deploying new contracts.
const PatriciaTreeBin = `0x608060405234801561001057600080fd5b5061124f806100206000396000f3006080604052600436106100775763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166320ba5b60811461007c57806350c946fe1461009e578063693ac4fb146100d657806380759f1f14610104578063a43914da14610126578063f7e498f61461014a575b600080fd5b34801561008857600080fd5b5061009c610097366004610f92565b610177565b005b3480156100aa57600080fd5b506100be6100b9366004610e82565b61018d565b6040516100cd939291906110f5565b60405180910390f35b3480156100e257600080fd5b506100f66100f1366004610f5d565b610278565b6040516100cd929190611161565b34801561011057600080fd5b506101196104d7565b6040516100cd919061112b565b34801561013257600080fd5b5061013b6104dd565b6040516100cd93929190611139565b34801561015657600080fd5b5061016a610165366004610ea0565b610527565b6040516100cd919061111d565b6101896000838363ffffffff61067c16565b5050565b610195610cc2565b61019d610cc2565b6101a5610cc2565b6101ad610cdd565b6000858152600460209081526040808320815160608101909252909290918391908201908390600290835b828210156102265760408051808201825260038402860180548252825180840190935260018082015484526002909101546020848101919091528083019390935290835290920191016101d8565b5050509152505080515151855280515160209081015151855281515181015181015184528151810151518682015281518101518101515185820152905181015181015181015190830152509193909250565b60008054606090151561028a57600080fd5b610292610cf6565b60408051908101604052808580519060200120815260200161010081525090506102ba610d0d565b506040805180820182526001548152815180830190925260025482526003546020838101919091528101919091526102f0610d28565b6000806102fb610cf6565b610303610cf6565b600061030d610cf6565b6020880151610323908a9063ffffffff61073b16565b6020808b0151810151908301519296509094501461033d57fe5b6020830151151561034d57610454565b60208401519590950160ff81900360020a9a909a17996001019461037083610769565b8951600090815260046020526040902091935091506103d69060018490036002811061039857fe5b6040805180820182526003929092029290920180548252825180840190935260018101548352600201546020838101919091528101919091526107d0565b6001860195889061010081106103e857fe5b6020908102919091019190915288516000908152600490915260409020826002811061041057fe5b60408051808201825260039290920292909201805482528251808401909352600181015483526002015460208381019190915281019190915290985096508761030d565b60008511156104c95784604051908082528060200260200182016040528015610487578160200160208202803883390190505b50995060005b858110156104c757878161010081106104a257fe5b60200201518b828151811015156104b557fe5b6020908102909101015260010161048d565b505b505050505050505050915091565b60005490565b60008060006104ea610d0d565b5050604080518082018252600154815281518083019092526002548252600354602083810191825282018390529051915190519194909350909150565b6000610531610cf6565b6040805190810160405280878051906020012081526020016101008152509050610559610d0d565b85516020870120815260005b851561065257600061057687610812565b60ff1690508060019060020a02198716965061059e8160ff038561086690919063ffffffff16565b602085018190529094506000906105b490610769565b602086015290506105c3610cc2565b6105cc856107d0565b8183600281106105d857fe5b602002015287518890600019868203019081106105f157fe5b90602001906020020151818360010360028110151561060c57fe5b6020908102919091019190915281519181015160408051808401949094528381019190915280518084038201815260609093019052815191012084525050600101610565565b5060208101829052610663816107d0565b881461066e57600080fd5b506001979650505050505050565b610684610cf6565b506040805180820190915282516020808501919091208252610100818301528251908301206106b1610d0d565b855415156106c8576020810183905281815261070b565b60408051808201825260018801548152815180830190925260028801548252600388015460208381019190915281019190915261070890879085856108da565b90505b610714816107d0565b86558051600187015560209081015180516002880155015160039095019490945550505050565b610743610cf6565b61074b610cf6565b61075e846107598686610b14565b610866565b915091509250929050565b6000610773610cf6565b602083015160001061078457600080fd5b50508051604080518082019091528251600202815260209283015160001901928101929092527f8000000000000000000000000000000000000000000000000000000000000000900491565b80516020918201518083015190516040805180860194909452838101929092526060808401919091528151808403909101815260809092019052805191012090565b600081151561082057600080fd5b8160805b600160ff82161061085f5760001960ff821660020a0182161515610852579182019160ff811660020a909104905b600260ff90911604610824565b5050919050565b61086e610cf6565b610876610cf6565b8360200151831115801561088c57506101008311155b151561089457fe5b602082018390528215156108ab57600082526108bd565b835160ff84900360020a600119021682525b60208085015184900390820152925160029290920a909102825291565b6108e2610d0d565b6020808501518101519084015110156108f757fe5b6108ff610cf6565b610907610cf6565b6000610911610cf6565b61091f87896020015161073b565b602081015191955093506000901515610939575085610aef565b6020808a01518101519086015110610a6057602084015160011061095957fe5b610961610cdd565b8951600090815260048c0160209081526040808320815160608101909252909290918391908201908390600290835b828210156109de576040805180820182526003840286018054825282518084019093526001808201548452600290910154602084810191909152808301939093529083529092019101610990565b505050508152505090506109f185610769565b82519195509350610a15908c908660028110610a0957fe5b6020020151858b6108da565b81518560028110610a2257fe5b602090810291909101919091528a51600090815260048d019091526040812090610a4c8282610d49565b5050610a588b82610b84565b915050610aef565b610a6984610769565b9093509150610a76610cdd565b604080518082019091528881526020810184905281518560028110610a9757fe5b602002018190525060408051908101604052808b600001518152602001610ac98c602001518960200151600101610bea565b90528151600186900360028110610adc57fe5b6020020152610aeb8b82610b84565b9150505b604080519081016040528082815260200186815250955050505050505b949350505050565b6000808260200151846020015110610b30578260200151610b36565b83602001515b9050801515610b49576000915050610b7e565b825184511861010082900360020a60000316801515610b6a57509050610b7e565b610b7381610c1e565b60ff0360ff16925050505b92915050565b600080610b9083610c6c565b83515160008281526004968701602090815260409091208251815591810151805160018401558101516002830155945185015180516003830155850151805196820196909655949093015160059094019390935550919050565b610bf2610cf6565b6020830151821115610c0357600080fd5b60208084015183900390820152915160029190910a02815290565b6000811515610c2c57600080fd5b8160805b600160ff82161061085f5760ff811660020a600019810102821615610c5f579182019160ff811660020a909104905b600260ff90911604610c30565b8051600090610c8190825b60200201516107d0565b8251610c8e906001610c77565b6040516020018083815260200182815260200192505050604051602081830303815290604052805190602001209050919050565b60408051808201825290600290829080388339509192915050565b60c060405190810160405280610cf1610d73565b905290565b604080518082019091526000808252602082015290565b60408051606081019091526000815260208101610cf1610cf6565b61200060405190810160405280610100906020820280388339509192915050565b50600080825560018201819055600282018190556003820181905560048201819055600590910155565b60c0604051908101604052806002905b610d8b610d0d565b815260200190600190039081610d835790505090565b6000601f82018313610db257600080fd5b8135610dc5610dc0826111a8565b611181565b91508181835260208401935060208101905083856020840282011115610dea57600080fd5b60005b83811015610e165781610e008882610e20565b8452506020928301929190910190600101610ded565b5050505092915050565b6000610e2c82356111f1565b9392505050565b6000601f82018313610e4457600080fd5b8135610e52610dc0826111c9565b91508082526020830160208301858383011115610e6e57600080fd5b610e79838284611209565b50505092915050565b600060208284031215610e9457600080fd5b6000610b0c8484610e20565b600080600080600060a08688031215610eb857600080fd5b6000610ec48888610e20565b955050602086013567ffffffffffffffff811115610ee157600080fd5b610eed88828901610e33565b945050604086013567ffffffffffffffff811115610f0a57600080fd5b610f1688828901610e33565b9350506060610f2788828901610e20565b925050608086013567ffffffffffffffff811115610f4457600080fd5b610f5088828901610da1565b9150509295509295909350565b600060208284031215610f6f57600080fd5b813567ffffffffffffffff811115610f8657600080fd5b610b0c84828501610e33565b60008060408385031215610fa557600080fd5b823567ffffffffffffffff811115610fbc57600080fd5b610fc885828601610e33565b925050602083013567ffffffffffffffff811115610fe557600080fd5b610ff185828601610e33565b9150509250929050565b611004816111fa565b61100d826111f1565b60005b8281101561103d576110238583516110ec565b61102c826111f4565b602095909501949150600101611010565b5050505050565b600061104f82611200565b808452602084019350611061836111f4565b60005b82811015611091576110778683516110ec565b611080826111f4565b602096909601959150600101611064565b5093949350505050565b6110a4816111fa565b6110ad826111f1565b60005b8281101561103d576110c38583516110ec565b6110cc826111f4565b6020959095019491506001016110b0565b6110e681611204565b82525050565b6110e6816111f1565b60c081016111038286610ffb565b6111106040830185610ffb565b610b0c608083018461109b565b60208101610b7e82846110dd565b60208101610b7e82846110ec565b6060810161114782866110ec565b61115460208301856110ec565b610b0c60408301846110ec565b6040810161116f82856110ec565b8181036020830152610b0c8184611044565b60405181810167ffffffffffffffff811182821017156111a057600080fd5b604052919050565b600067ffffffffffffffff8211156111bf57600080fd5b5060209081020190565b600067ffffffffffffffff8211156111e057600080fd5b506020601f91909101601f19160190565b90565b60200190565b50600290565b5190565b151590565b828183375060009101525600a265627a7a7230582018b914ecdc6c0bae868236caa7c87914f895736b958199d65b8e37a35f3a8b3c6c6578706572696d656e74616cf50037`

// DeployPatriciaTree deploys a new Ethereum contract, binding an instance of PatriciaTree to it.
func DeployPatriciaTree(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PatriciaTree, error) {
	parsed, err := abi.JSON(strings.NewReader(PatriciaTreeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PatriciaTreeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PatriciaTree{PatriciaTreeCaller: PatriciaTreeCaller{contract: contract}, PatriciaTreeTransactor: PatriciaTreeTransactor{contract: contract}, PatriciaTreeFilterer: PatriciaTreeFilterer{contract: contract}}, nil
}

// PatriciaTree is an auto generated Go binding around an Ethereum contract.
type PatriciaTree struct {
	PatriciaTreeCaller     // Read-only binding to the contract
	PatriciaTreeTransactor // Write-only binding to the contract
	PatriciaTreeFilterer   // Log filterer for contract events
}

// PatriciaTreeCaller is an auto generated read-only Go binding around an Ethereum contract.
type PatriciaTreeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PatriciaTreeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PatriciaTreeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PatriciaTreeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PatriciaTreeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PatriciaTreeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PatriciaTreeSession struct {
	Contract     *PatriciaTree     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PatriciaTreeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PatriciaTreeCallerSession struct {
	Contract *PatriciaTreeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// PatriciaTreeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PatriciaTreeTransactorSession struct {
	Contract     *PatriciaTreeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// PatriciaTreeRaw is an auto generated low-level Go binding around an Ethereum contract.
type PatriciaTreeRaw struct {
	Contract *PatriciaTree // Generic contract binding to access the raw methods on
}

// PatriciaTreeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PatriciaTreeCallerRaw struct {
	Contract *PatriciaTreeCaller // Generic read-only contract binding to access the raw methods on
}

// PatriciaTreeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PatriciaTreeTransactorRaw struct {
	Contract *PatriciaTreeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPatriciaTree creates a new instance of PatriciaTree, bound to a specific deployed contract.
func NewPatriciaTree(address common.Address, backend bind.ContractBackend) (*PatriciaTree, error) {
	contract, err := bindPatriciaTree(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PatriciaTree{PatriciaTreeCaller: PatriciaTreeCaller{contract: contract}, PatriciaTreeTransactor: PatriciaTreeTransactor{contract: contract}, PatriciaTreeFilterer: PatriciaTreeFilterer{contract: contract}}, nil
}

// NewPatriciaTreeCaller creates a new read-only instance of PatriciaTree, bound to a specific deployed contract.
func NewPatriciaTreeCaller(address common.Address, caller bind.ContractCaller) (*PatriciaTreeCaller, error) {
	contract, err := bindPatriciaTree(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PatriciaTreeCaller{contract: contract}, nil
}

// NewPatriciaTreeTransactor creates a new write-only instance of PatriciaTree, bound to a specific deployed contract.
func NewPatriciaTreeTransactor(address common.Address, transactor bind.ContractTransactor) (*PatriciaTreeTransactor, error) {
	contract, err := bindPatriciaTree(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PatriciaTreeTransactor{contract: contract}, nil
}

// NewPatriciaTreeFilterer creates a new log filterer instance of PatriciaTree, bound to a specific deployed contract.
func NewPatriciaTreeFilterer(address common.Address, filterer bind.ContractFilterer) (*PatriciaTreeFilterer, error) {
	contract, err := bindPatriciaTree(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PatriciaTreeFilterer{contract: contract}, nil
}

// bindPatriciaTree binds a generic wrapper to an already deployed contract.
func bindPatriciaTree(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PatriciaTreeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PatriciaTree *PatriciaTreeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PatriciaTree.Contract.PatriciaTreeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PatriciaTree *PatriciaTreeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PatriciaTree.Contract.PatriciaTreeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PatriciaTree *PatriciaTreeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PatriciaTree.Contract.PatriciaTreeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PatriciaTree *PatriciaTreeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PatriciaTree.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PatriciaTree *PatriciaTreeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PatriciaTree.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PatriciaTree *PatriciaTreeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PatriciaTree.Contract.contract.Transact(opts, method, params...)
}

// GetNode is a free data retrieval call binding the contract method 0x50c946fe.
//
// Solidity: function getNode(hash bytes32) constant returns(nodes bytes32[2], labelDatas bytes32[2], labelLengths uint256[2])
func (_PatriciaTree *PatriciaTreeCaller) GetNode(opts *bind.CallOpts, hash [32]byte) (struct {
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
	err := _PatriciaTree.contract.Call(opts, out, "getNode", hash)
	return *ret, err
}

// GetNode is a free data retrieval call binding the contract method 0x50c946fe.
//
// Solidity: function getNode(hash bytes32) constant returns(nodes bytes32[2], labelDatas bytes32[2], labelLengths uint256[2])
func (_PatriciaTree *PatriciaTreeSession) GetNode(hash [32]byte) (struct {
	Nodes        [2][32]byte
	LabelDatas   [2][32]byte
	LabelLengths [2]*big.Int
}, error) {
	return _PatriciaTree.Contract.GetNode(&_PatriciaTree.CallOpts, hash)
}

// GetNode is a free data retrieval call binding the contract method 0x50c946fe.
//
// Solidity: function getNode(hash bytes32) constant returns(nodes bytes32[2], labelDatas bytes32[2], labelLengths uint256[2])
func (_PatriciaTree *PatriciaTreeCallerSession) GetNode(hash [32]byte) (struct {
	Nodes        [2][32]byte
	LabelDatas   [2][32]byte
	LabelLengths [2]*big.Int
}, error) {
	return _PatriciaTree.Contract.GetNode(&_PatriciaTree.CallOpts, hash)
}

// GetProof is a free data retrieval call binding the contract method 0x693ac4fb.
//
// Solidity: function getProof(key bytes) constant returns(branchMask uint256, _siblings bytes32[])
func (_PatriciaTree *PatriciaTreeCaller) GetProof(opts *bind.CallOpts, key []byte) (struct {
	BranchMask *big.Int
	Siblings   [][32]byte
}, error) {
	ret := new(struct {
		BranchMask *big.Int
		Siblings   [][32]byte
	})
	out := ret
	err := _PatriciaTree.contract.Call(opts, out, "getProof", key)
	return *ret, err
}

// GetProof is a free data retrieval call binding the contract method 0x693ac4fb.
//
// Solidity: function getProof(key bytes) constant returns(branchMask uint256, _siblings bytes32[])
func (_PatriciaTree *PatriciaTreeSession) GetProof(key []byte) (struct {
	BranchMask *big.Int
	Siblings   [][32]byte
}, error) {
	return _PatriciaTree.Contract.GetProof(&_PatriciaTree.CallOpts, key)
}

// GetProof is a free data retrieval call binding the contract method 0x693ac4fb.
//
// Solidity: function getProof(key bytes) constant returns(branchMask uint256, _siblings bytes32[])
func (_PatriciaTree *PatriciaTreeCallerSession) GetProof(key []byte) (struct {
	BranchMask *big.Int
	Siblings   [][32]byte
}, error) {
	return _PatriciaTree.Contract.GetProof(&_PatriciaTree.CallOpts, key)
}

// GetRootEdge is a free data retrieval call binding the contract method 0xa43914da.
//
// Solidity: function getRootEdge() constant returns(node bytes32, labelData bytes32, labelLength uint256)
func (_PatriciaTree *PatriciaTreeCaller) GetRootEdge(opts *bind.CallOpts) (struct {
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
	err := _PatriciaTree.contract.Call(opts, out, "getRootEdge")
	return *ret, err
}

// GetRootEdge is a free data retrieval call binding the contract method 0xa43914da.
//
// Solidity: function getRootEdge() constant returns(node bytes32, labelData bytes32, labelLength uint256)
func (_PatriciaTree *PatriciaTreeSession) GetRootEdge() (struct {
	Node        [32]byte
	LabelData   [32]byte
	LabelLength *big.Int
}, error) {
	return _PatriciaTree.Contract.GetRootEdge(&_PatriciaTree.CallOpts)
}

// GetRootEdge is a free data retrieval call binding the contract method 0xa43914da.
//
// Solidity: function getRootEdge() constant returns(node bytes32, labelData bytes32, labelLength uint256)
func (_PatriciaTree *PatriciaTreeCallerSession) GetRootEdge() (struct {
	Node        [32]byte
	LabelData   [32]byte
	LabelLength *big.Int
}, error) {
	return _PatriciaTree.Contract.GetRootEdge(&_PatriciaTree.CallOpts)
}

// GetRootHash is a free data retrieval call binding the contract method 0x80759f1f.
//
// Solidity: function getRootHash() constant returns(h bytes32)
func (_PatriciaTree *PatriciaTreeCaller) GetRootHash(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _PatriciaTree.contract.Call(opts, out, "getRootHash")
	return *ret0, err
}

// GetRootHash is a free data retrieval call binding the contract method 0x80759f1f.
//
// Solidity: function getRootHash() constant returns(h bytes32)
func (_PatriciaTree *PatriciaTreeSession) GetRootHash() ([32]byte, error) {
	return _PatriciaTree.Contract.GetRootHash(&_PatriciaTree.CallOpts)
}

// GetRootHash is a free data retrieval call binding the contract method 0x80759f1f.
//
// Solidity: function getRootHash() constant returns(h bytes32)
func (_PatriciaTree *PatriciaTreeCallerSession) GetRootHash() ([32]byte, error) {
	return _PatriciaTree.Contract.GetRootHash(&_PatriciaTree.CallOpts)
}

// VerifyProof is a free data retrieval call binding the contract method 0xf7e498f6.
//
// Solidity: function verifyProof(rootHash bytes32, key bytes, value bytes, branchMask uint256, siblings bytes32[]) constant returns(success bool)
func (_PatriciaTree *PatriciaTreeCaller) VerifyProof(opts *bind.CallOpts, rootHash [32]byte, key []byte, value []byte, branchMask *big.Int, siblings [][32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _PatriciaTree.contract.Call(opts, out, "verifyProof", rootHash, key, value, branchMask, siblings)
	return *ret0, err
}

// VerifyProof is a free data retrieval call binding the contract method 0xf7e498f6.
//
// Solidity: function verifyProof(rootHash bytes32, key bytes, value bytes, branchMask uint256, siblings bytes32[]) constant returns(success bool)
func (_PatriciaTree *PatriciaTreeSession) VerifyProof(rootHash [32]byte, key []byte, value []byte, branchMask *big.Int, siblings [][32]byte) (bool, error) {
	return _PatriciaTree.Contract.VerifyProof(&_PatriciaTree.CallOpts, rootHash, key, value, branchMask, siblings)
}

// VerifyProof is a free data retrieval call binding the contract method 0xf7e498f6.
//
// Solidity: function verifyProof(rootHash bytes32, key bytes, value bytes, branchMask uint256, siblings bytes32[]) constant returns(success bool)
func (_PatriciaTree *PatriciaTreeCallerSession) VerifyProof(rootHash [32]byte, key []byte, value []byte, branchMask *big.Int, siblings [][32]byte) (bool, error) {
	return _PatriciaTree.Contract.VerifyProof(&_PatriciaTree.CallOpts, rootHash, key, value, branchMask, siblings)
}

// Insert is a paid mutator transaction binding the contract method 0x20ba5b60.
//
// Solidity: function insert(key bytes, value bytes) returns()
func (_PatriciaTree *PatriciaTreeTransactor) Insert(opts *bind.TransactOpts, key []byte, value []byte) (*types.Transaction, error) {
	return _PatriciaTree.contract.Transact(opts, "insert", key, value)
}

// Insert is a paid mutator transaction binding the contract method 0x20ba5b60.
//
// Solidity: function insert(key bytes, value bytes) returns()
func (_PatriciaTree *PatriciaTreeSession) Insert(key []byte, value []byte) (*types.Transaction, error) {
	return _PatriciaTree.Contract.Insert(&_PatriciaTree.TransactOpts, key, value)
}

// Insert is a paid mutator transaction binding the contract method 0x20ba5b60.
//
// Solidity: function insert(key bytes, value bytes) returns()
func (_PatriciaTree *PatriciaTreeTransactorSession) Insert(key []byte, value []byte) (*types.Transaction, error) {
	return _PatriciaTree.Contract.Insert(&_PatriciaTree.TransactOpts, key, value)
}

// PatriciaTreeDataABI is the input ABI used to generate the binding from.
const PatriciaTreeDataABI = "[]"

// PatriciaTreeDataBin is the compiled bytecode used for deploying new contracts.
const PatriciaTreeDataBin = `0x605a602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a265627a7a72305820c1bcc829b65d2265560e832a095ceb3b9129c0717854dd5a99bc47e5216bd9526c6578706572696d656e74616cf50037`

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
const RLPBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a72305820230548a37307bfa6c80e09c03d291d975ca5f35fca5a2faeb37542bf111860bd0029`

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
const RLPEncodeBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a72305820f2ceb9bea42b35ab26e1fabebadc09fd1ed9cd72d112da05517f68206ee7825e0029`

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
const RootChainABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"COST_URB_PREPARE\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CP_COMPUTATION\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"epochs\",\"outputs\":[{\"name\":\"requestStart\",\"type\":\"uint64\"},{\"name\":\"requestEnd\",\"type\":\"uint64\"},{\"name\":\"startBlockNumber\",\"type\":\"uint64\"},{\"name\":\"endBlockNumber\",\"type\":\"uint64\"},{\"name\":\"forkedBlockNumber\",\"type\":\"uint64\"},{\"name\":\"firstRequestBlockId\",\"type\":\"uint64\"},{\"name\":\"limit\",\"type\":\"uint64\"},{\"name\":\"timestamp\",\"type\":\"uint64\"},{\"name\":\"isEmpty\",\"type\":\"bool\"},{\"name\":\"initialized\",\"type\":\"bool\"},{\"name\":\"isRequest\",\"type\":\"bool\"},{\"name\":\"userActivated\",\"type\":\"bool\"},{\"name\":\"finalized\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastAppliedForkNumber\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"currentFork\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"COST_URB\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastAppliedERU\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_trieKey\",\"type\":\"bytes32\"},{\"name\":\"_trieValue\",\"type\":\"bytes32\"}],\"name\":\"startExit\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_trieKey\",\"type\":\"bytes32\"},{\"name\":\"_trieValue\",\"type\":\"bytes32\"}],\"name\":\"startEnter\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_forkNumber\",\"type\":\"uint256\"},{\"name\":\"_epochNumber\",\"type\":\"uint256\"}],\"name\":\"getEpoch\",\"outputs\":[{\"name\":\"requestStart\",\"type\":\"uint64\"},{\"name\":\"requestEnd\",\"type\":\"uint64\"},{\"name\":\"startBlockNumber\",\"type\":\"uint64\"},{\"name\":\"endBlockNumber\",\"type\":\"uint64\"},{\"name\":\"forkedBlockNumber\",\"type\":\"uint64\"},{\"name\":\"isEmpty\",\"type\":\"bool\"},{\"name\":\"initialized\",\"type\":\"bool\"},{\"name\":\"isRequest\",\"type\":\"bool\"},{\"name\":\"userActivated\",\"type\":\"bool\"},{\"name\":\"finalized\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"firstEpoch\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"operator\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"NRBEpochLength\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"applyRequest\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_blockNumber\",\"type\":\"uint256\"},{\"name\":\"_key\",\"type\":\"bytes\"},{\"name\":\"_txByte\",\"type\":\"bytes\"},{\"name\":\"_branchMask\",\"type\":\"uint256\"},{\"name\":\"_siblings\",\"type\":\"bytes32[]\"}],\"name\":\"challengeNullAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastAppliedERO\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"finalizeBlock\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_forkNumber\",\"type\":\"uint256\"},{\"name\":\"_blockNumber\",\"type\":\"uint256\"}],\"name\":\"revertBlock\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"currentEpoch\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"lastFinalizedBlock\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"development\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"COST_ERU\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"REQUEST_GAS\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_statesRoot\",\"type\":\"bytes32\"},{\"name\":\"_transactionsRoot\",\"type\":\"bytes32\"},{\"name\":\"_intermediateStatesRoot\",\"type\":\"bytes32\"}],\"name\":\"submitNRB\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAX_REQUESTS\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"COST_NRB\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_statesRoot\",\"type\":\"bytes32\"},{\"name\":\"_transactionsRoot\",\"type\":\"bytes32\"},{\"name\":\"_intermediateStatesRoot\",\"type\":\"bytes32\"}],\"name\":\"submitURB\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"finalizeRequest\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_requestId\",\"type\":\"uint256\"},{\"name\":\"_userActivated\",\"type\":\"bool\"}],\"name\":\"getRequestApplied\",\"outputs\":[{\"name\":\"applied\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CP_WITHHOLDING\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"COST_ORB\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_trieKey\",\"type\":\"bytes32\"},{\"name\":\"_trieValue\",\"type\":\"bytes32\"}],\"name\":\"makeERU\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"EROs\",\"outputs\":[{\"name\":\"timestamp\",\"type\":\"uint64\"},{\"name\":\"isExit\",\"type\":\"bool\"},{\"name\":\"applied\",\"type\":\"bool\"},{\"name\":\"finalized\",\"type\":\"bool\"},{\"name\":\"challenged\",\"type\":\"bool\"},{\"name\":\"value\",\"type\":\"uint128\"},{\"name\":\"requestor\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"trieKey\",\"type\":\"bytes32\"},{\"name\":\"trieValue\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getNumEROs\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"URBs\",\"outputs\":[{\"name\":\"submitted\",\"type\":\"bool\"},{\"name\":\"requestStart\",\"type\":\"uint64\"},{\"name\":\"requestEnd\",\"type\":\"uint64\"},{\"name\":\"trie\",\"type\":\"address\"},{\"name\":\"transactionsRoot\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"state\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PREPARE_TIMEOUT\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_rootchain\",\"type\":\"address\"},{\"name\":\"_childchain\",\"type\":\"address\"}],\"name\":\"mapRequestableContractByOperator\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_forkNumber\",\"type\":\"uint256\"}],\"name\":\"forked\",\"outputs\":[{\"name\":\"result\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"COST_ERO\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"requestableContracts\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"NULL_ADDRESS\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_statesRoot\",\"type\":\"bytes32\"},{\"name\":\"_transactionsRoot\",\"type\":\"bytes32\"},{\"name\":\"_intermediateStatesRoot\",\"type\":\"bytes32\"}],\"name\":\"submitORB\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"prepareToSubmitURB\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"ORBs\",\"outputs\":[{\"name\":\"submitted\",\"type\":\"bool\"},{\"name\":\"requestStart\",\"type\":\"uint64\"},{\"name\":\"requestEnd\",\"type\":\"uint64\"},{\"name\":\"trie\",\"type\":\"address\"},{\"name\":\"transactionsRoot\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"highestBlockNumber\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_requestId\",\"type\":\"uint256\"},{\"name\":\"_userActivated\",\"type\":\"bool\"}],\"name\":\"getRequestFinalized\",\"outputs\":[{\"name\":\"finalized\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_target\",\"type\":\"address\"}],\"name\":\"mapRequestableContract\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"ERUs\",\"outputs\":[{\"name\":\"timestamp\",\"type\":\"uint64\"},{\"name\":\"isExit\",\"type\":\"bool\"},{\"name\":\"applied\",\"type\":\"bool\"},{\"name\":\"finalized\",\"type\":\"bool\"},{\"name\":\"challenged\",\"type\":\"bool\"},{\"name\":\"value\",\"type\":\"uint128\"},{\"name\":\"requestor\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"trieKey\",\"type\":\"bytes32\"},{\"name\":\"trieValue\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"blocks\",\"outputs\":[{\"name\":\"statesRoot\",\"type\":\"bytes32\"},{\"name\":\"transactionsRoot\",\"type\":\"bytes32\"},{\"name\":\"intermediateStatesRoot\",\"type\":\"bytes32\"},{\"name\":\"forkNumber\",\"type\":\"uint64\"},{\"name\":\"epochNumber\",\"type\":\"uint64\"},{\"name\":\"timestamp\",\"type\":\"uint64\"},{\"name\":\"requestBlockId\",\"type\":\"uint64\"},{\"name\":\"isRequest\",\"type\":\"bool\"},{\"name\":\"userActivated\",\"type\":\"bool\"},{\"name\":\"challenged\",\"type\":\"bool\"},{\"name\":\"challenging\",\"type\":\"bool\"},{\"name\":\"finalized\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastAppliedBlockNumber\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_development\",\"type\":\"bool\"},{\"name\":\"_NRBEpochLength\",\"type\":\"uint256\"},{\"name\":\"_statesRoot\",\"type\":\"bytes32\"},{\"name\":\"_transactionsRoot\",\"type\":\"bytes32\"},{\"name\":\"_intermediateStatesRoot\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"userActivated\",\"type\":\"bool\"}],\"name\":\"SessionTimeout\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"state\",\"type\":\"uint8\"}],\"name\":\"StateChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newFork\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"forkedBlockNumber\",\"type\":\"uint256\"}],\"name\":\"Forked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"forkNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"epochNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"isEmpty\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"startBlockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"endBlockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"requestStart\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"requestEnd\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"isRequest\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"userActivated\",\"type\":\"bool\"}],\"name\":\"EpochPrepared\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"fork\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"NRBSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"fork\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"ORBSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"fork\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"URBSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"requestor\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"weiAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"trieKey\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"trieValue\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"isExit\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"userActivated\",\"type\":\"bool\"}],\"name\":\"RequestCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"requestor\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"trieKey\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"trieValue\",\"type\":\"bytes32\"}],\"name\":\"ERUCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"forkNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"BlockFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"forkNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"epochNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"firstBlockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"lastBlockNumber\",\"type\":\"uint256\"}],\"name\":\"EpochFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"userActivated\",\"type\":\"bool\"}],\"name\":\"RequestFinalized\",\"type\":\"event\"}]"

// RootChainBin is the compiled bytecode used for deploying new contracts.
const RootChainBin = `0x60806040526000805460ff191660011790553480156200001e57600080fd5b5060405160a08062006341833981018060405260a08110156200004057600080fd5b50805160208083015160408085015160608601516080909601516000805460ff19168715151761010060a860020a0319166101003381029190911782556008869055600180548352600688528583208380528852948220848155948501899055600285018390557fc65916d663d52b0a18b63681e34ad6e3e8bb58d57f662b8e7e045ab09fab038780547f6d5257204ebe7d88fd91ae87941cb2dd9d8062b64ae5a2bd2d28ec40b9fbf6df90985264010000000061ff001990981690911764ff00000000191687179055959693959194909291620001249183919062000143810204565b62000137640100000000620001b0810204565b5050505050506200040e565b60048201805464ff000000001916640100000000179055600180546000908152600560209081526040918290208490559154815190815291820183905280517ffb96205e4b3633fd57aa805b26b51ecf528714a10241a4af015929dce86768d99281900390910190a15050565b600180546000908152600460205260408120549091620001df919064010000000062003d63620003f482021704565b600380546001908101918290558054600090815260076020908152604080832094835293905291822060028101805461ff00191661010017905580546001604060020a0380861670010000000000000000000000000000000002608060020a60c060020a0319909216919091178083556008546000199087010182167801000000000000000000000000000000000000000000000000908102600160c060020a03928316178455938301805442909316909402911617909155815492935091819060a860020a60ff02191675010000000000000000000000000000000000000000008202179055506000546040517f551dc40198cc79684bb69e4931dba4ac16e4598792ee1c0a5000aeea366d7bb6917501000000000000000000000000000000000000000000900460ff1690808260028111156200031a57fe5b60ff16815260200191505060405180910390a16001546003546002830154835460408051948552602085019390935260ff9091161515838301526001604060020a03700100000000000000000000000000000000820481166060850152780100000000000000000000000000000000000000000000000082048116608085015280821660a0850152680100000000000000009091041660c0830152600060e08301819052610100830152517fe0bc0e53be29ccc81a897770a6c14b5b81310036fabf148eb7e92723c3d79e1b918190036101200190a15050565b6000828201838110156200040757600080fd5b9392505050565b615f23806200041e6000396000f3006080604052600436106102665763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663033cfbed811461026b57806308c4fff0146102925780630a88bd04146102a7578063164bc2ae14610358578063183d2d1c1461036d578063192adc5b146103825780631f261d591461039757806324839704146103ac57806325119293146103f25780632b25a38b146104245780632e7ab948146104bb578063570ca735146104e5578063584e349e1461051657806361d29e831461052b5780636299fb241461054057806365d724bc1461066e57806375395a5814610683578063760247921461069857806376671808146106c857806376e61c36146106dd5780637b929c27146107075780638b5172d01461071c5780638eb288ca1461073157806393248d0214610746578063935212221461076f57806394be3aa51461026b57806397be455d1461078457806399bd3600146107ad578063ae08ffe9146107c2578063b17fa6e9146107f4578063b2ae9ba81461026b578063b41e69dd14610809578063b443f3cc14610848578063b540adba146108e0578063c0e86064146108f5578063c19d93fb14610961578063c2bc88fa1461099a578063cb5d742f146109af578063ce8a2bc2146109ea578063d691acd81461026b578063da0185f814610a14578063de0ce17d14610a47578063e67123c414610a5c578063e6925d0814610a85578063ea7f22a814610a9a578063f1f3c46c14610ac4578063f28a7afa14610aee578063f3aba34414610b20578063f4f31de414610b53578063f4f911db14610b7d578063fb788a2714610c22575b600080fd5b34801561027757600080fd5b50610280610c37565b60408051918252519081900360200190f35b34801561029e57600080fd5b50610280610c43565b3480156102b357600080fd5b506102d7600480360360408110156102ca57600080fd5b5080359060200135610c48565b604080516001604060020a039e8f1681529c8e1660208e01529a8d168c8c0152988c1660608c0152968b1660808b0152948a1660a08a015292891660c0890152971660e08701529515156101008601529415156101208501529315156101408401529215156101608301529115156101808201529051908190036101a00190f35b34801561036457600080fd5b50610280610ce1565b34801561037957600080fd5b50610280610ce7565b34801561038e57600080fd5b50610280610ced565b3480156103a357600080fd5b50610280610cf9565b6103de600480360360608110156103c257600080fd5b50600160a060020a038135169060208101359060400135610cff565b604080519115158252519081900360200190f35b6103de6004803603606081101561040857600080fd5b50600160a060020a038135169060208101359060400135610dac565b34801561043057600080fd5b506104546004803603604081101561044757600080fd5b5080359060200135610e41565b604080516001604060020a039b8c168152998b1660208b0152978a1689890152958916606089015293909716608087015290151560a0860152151560c085015293151560e08401529215156101008301529115156101208201529051908190036101400190f35b3480156104c757600080fd5b50610280600480360360208110156104de57600080fd5b5035610ec2565b3480156104f157600080fd5b506104fa610ed4565b60408051600160a060020a039092168252519081900360200190f35b34801561052257600080fd5b50610280610ee8565b34801561053757600080fd5b506103de610eee565b34801561054c57600080fd5b5061066c600480360360a081101561056357600080fd5b8135919081019060408101602082013564010000000081111561058557600080fd5b82018360208201111561059757600080fd5b803590602001918460018302840111640100000000831117156105b957600080fd5b9193909290916020810190356401000000008111156105d757600080fd5b8201836020820111156105e957600080fd5b8035906020019184600183028401116401000000008311171561060b57600080fd5b9193909282359260408101906020013564010000000081111561062d57600080fd5b82018360208201111561063f57600080fd5b8035906020019184602083028401116401000000008311171561066157600080fd5b5090925090506114d6565b005b34801561067a57600080fd5b50610280611769565b34801561068f57600080fd5b506103de61176f565b3480156106a457600080fd5b5061066c600480360360408110156106bb57600080fd5b508035906020013561178a565b3480156106d457600080fd5b5061028061178e565b3480156106e957600080fd5b506102806004803603602081101561070057600080fd5b5035611794565b34801561071357600080fd5b506103de6117a6565b34801561072857600080fd5b506102806117af565b34801561073d57600080fd5b506102806117bb565b6103de6004803603606081101561075c57600080fd5b50803590602081013590604001356117c2565b34801561077b57600080fd5b506102806119cf565b6103de6004803603606081101561079a57600080fd5b50803590602081013590604001356119d5565b3480156107b957600080fd5b506103de611be0565b3480156107ce57600080fd5b506103de600480360360408110156107e557600080fd5b50803590602001351515611be5565b34801561080057600080fd5b50610280611c32565b34801561081557600080fd5b506103de6004803603606081101561082c57600080fd5b50600160a060020a038135169060208101359060400135611c37565b34801561085457600080fd5b506108726004803603602081101561086b57600080fd5b5035611ce3565b604080516001604060020a03909b168b5298151560208b015296151589890152941515606089015292151560808801526001608060020a0390911660a0870152600160a060020a0390811660c08701521660e085015261010084015261012083015251908190036101400190f35b3480156108ec57600080fd5b50610280611d85565b34801561090157600080fd5b5061091f6004803603602081101561091857600080fd5b5035611d8b565b6040805195151586526001604060020a0394851660208701529290931684830152600160a060020a031660608401526080830191909152519081900360a00190f35b34801561096d57600080fd5b50610976611de3565b6040518082600281111561098657fe5b60ff16815260200191505060405180910390f35b3480156109a657600080fd5b50610280611df3565b3480156109bb57600080fd5b506103de600480360360408110156109d257600080fd5b50600160a060020a0381358116916020013516611df9565b3480156109f657600080fd5b506103de60048036036020811015610a0d57600080fd5b5035611e97565b348015610a2057600080fd5b506104fa60048036036020811015610a3757600080fd5b5035600160a060020a0316611ea3565b348015610a5357600080fd5b506104fa611ebe565b6103de60048036036060811015610a7257600080fd5b5080359060208101359060400135611ec3565b348015610a9157600080fd5b506103de61205b565b348015610aa657600080fd5b5061091f60048036036020811015610abd57600080fd5b50356120e6565b348015610ad057600080fd5b5061028060048036036020811015610ae757600080fd5b50356120f4565b348015610afa57600080fd5b506103de60048036036040811015610b1157600080fd5b50803590602001351515612106565b348015610b2c57600080fd5b506103de60048036036020811015610b4357600080fd5b5035600160a060020a0316612158565b348015610b5f57600080fd5b5061087260048036036020811015610b7657600080fd5b50356121d0565b348015610b8957600080fd5b50610bad60048036036040811015610ba057600080fd5b50803590602001356121de565b604080519c8d5260208d019b909b528b8b01999099526001604060020a0397881660608c015295871660808b015293861660a08a01529190941660c088015292151560e08701529115156101008601529015156101208501521515610140840152151561016083015251908190036101800190f35b348015610c2e57600080fd5b5061028061226c565b67016345785d8a000081565b600181565b60076020908152600092835260408084209091529082529020805460018201546002909201546001604060020a0380831693604060020a808504831694608060020a80820485169560c060020a92839004861695858116959485048116949283048116939092049091169060ff80821691610100810482169162010000820481169163010000008104821691640100000000909104168d565b600d5481565b60015481565b670c7d713b49da000081565b60105481565b600067016345785d8a000034811115610d1757600080fd5b600080610d2d6009600b89898960016000612272565b60408051838152336020820152600160a060020a038b168183015260608101839052608081018a905260a08101899052600160c0820152600060e082015290519294509092507f3362048aaeec35587d767ff0979557a3eceaa0ec93ab377113afde2d2fc026f091908190036101000190a15060019695505050505050565b6000806000610dc36009600b888888600080612272565b60408051838152336020820152600160a060020a038a1681830152606081018390526080810189905260a08101889052600060c0820181905260e082015290519294509092507f3362048aaeec35587d767ff0979557a3eceaa0ec93ab377113afde2d2fc026f091908190036101000190a150600195945050505050565b60009182526007602090815260408084209284529190529020805460018201546002909201546001604060020a0380831694604060020a8404821694608060020a850483169460c060020a900483169392169160ff808216926101008304821692620100008104831692630100000082048116926401000000009092041690565b60026020526000908152604090205481565b6000546101009004600160a060020a031681565b60085481565b600d54600e5460008281526004602052604081205490929183918290821115610f1657600080fd5b60008481526006602090815260408083208584528252808320600381015488855260078452828520604060020a9091046001604060020a03168086529352922060048301549195509060ff161515611016575b600086815260076020908152604080832088845290915290206002015462010000900460ff161515610fd0576000868152600760209081526040808320888452909152902060020154610100900460ff161515610fc557600080fd5b846001019450610f69565b546000868152600760209081526040808320888452825280832089845260068352818420608060020a9095046001604060020a0316808552949092529091209194509091505b6004820154640100000000900460ff16151561103157600080fd5b6004820154610100900460ff161561128d57601054600a54909350831061105757600080fd5b6000600a8481548110151561106857fe5b906000526020600020906005020190506000600c8460030160189054906101000a90046001604060020a03166001604060020a03168154811015156110a957fe5b600091825260209091206003909102018054909150604860020a90046001604060020a031685141561111c57600183015460006001604060020a03909116118015611107575060018301546000196001604060020a03918216011686145b156111145760018801600d555b60018601600e555b600185016010558154604060020a900460ff16801561114457508154604860020a900460ff16155b15611222578154604860020a69ff00000000000000000019909116811780845560408051610140810182526001604060020a038316815260ff604060020a840481161515602083015293830484161515918101919091526a010000000000000000000082048316151560608201526b010000000000000000000000820490921615156080830152606060020a90046001608060020a031660a08201526001830154600160a060020a0390811660c083015260028401541660e08201526003830154610100820152600483015461012082015261122090866127f0565b505b81546aff0000000000000000000019166a0100000000000000000000178255604080518681526001602082015281517f134017cf3262b18f892ee95dde3b0aec9a80cc70a7c96f09c64bd237aceb0473929181900390910190a16001985050505050505050506114d3565b600f5460095490935083106112a157600080fd5b60006009848154811015156112b257fe5b906000526020600020906005020190506000600b8460030160189054906101000a90046001604060020a03166001604060020a03168154811015156112f357fe5b600091825260209091206003909102018054909150604860020a90046001604060020a031685141561136657600183015460006001604060020a03909116118015611351575060018301546000196001604060020a03918216011686145b1561135e5760018801600d555b60018601600e555b60018501600f558154604060020a900460ff16801561138e57508154604860020a900460ff16155b1561146c578154604860020a69ff00000000000000000019909116811780845560408051610140810182526001604060020a038316815260ff604060020a840481161515602083015293830484161515918101919091526a010000000000000000000082048316151560608201526b010000000000000000000000820490921615156080830152606060020a90046001608060020a031660a08201526001830154600160a060020a0390811660c083015260028401541660e08201526003830154610100820152600483015461012082015261146a90866127f0565b505b81546aff0000000000000000000019166a0100000000000000000000178255604080518681526000602082015281517f134017cf3262b18f892ee95dde3b0aec9a80cc70a7c96f09c64bd237aceb0473929181900390910190a16001985050505050505050505b90565b60015460009081526006602090815260408083208b84529091529020600481015460ff161561150457600080fd5b6003810154426001604060020a03608060020a909204919091166001011161152b57600080fd5b6004810154600090610100900460ff1615611588576003820154600c8054909160c060020a90046001604060020a031690811061156457fe5b6000918252602090912060016003909202010154600160a060020a031690506115cc565b6003820154600b8054909160c060020a90046001604060020a03169081106115ac57fe5b6000918252602090912060016003909202010154600160a060020a031690505b6115d4614a5f565b61161388888080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061293392505050565b905061161e81612a7c565b151561162957600080fd5b81600160a060020a031663f7e498f684600101548c8c8c8c8c8c8c6040518963ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018089815260200180602001806020018681526020018060200184810384528b8b82818152602001925080828437600083820152601f01601f191690910185810384528981526020019050898980828437600083820152601f01601f19169091018581038352868152602090810191508790870280828437600081840152601f19601f8201169050808301925050509b50505050505050505050505060206040518083038186803b15801561172557600080fd5b505afa158015611739573d6000803e3d6000fd5b505050506040513d602081101561174f57600080fd5b5051151561175c57600080fd5b5050505050505050505050565b600f5481565b6000611779612aa5565b151561178457600080fd5b50600190565b5050565b60035481565b60056020526000908152604090205481565b60005460ff1681565b6702c68af0bb14000081565b620186a081565b600080546101009004600160a060020a031633146117df57600080fd5b60008060005460a860020a900460ff1660028111156117fa57fe5b1461180457600080fd5b67016345785d8a00003481111561181a57600080fd5b611822612aa5565b5060006118358787876000806000612c2b565b600154604080519182526020820183905280519293507fd85b68c66f9db94728c533f39125381246238815483496c57a62b56f70dea8f192918290030190a1600154600090815260076020908152604080832060035484529091529020600281015462010000900460ff16156118aa57600080fd5b604080516101a08101825282546001604060020a038082168352604060020a80830482166020850152608060020a80840483169585019590955260c060020a928390048216606085015260018601548083166080860152908104821660a0850152938404811660c0840152920490911660e0820152600282015460ff80821615156101008085019190915282048116151561012084015262010000820481161515610140840152630100000082048116151561016084015264010000000090910416151561018082015260009061198090612ecf565b825460018054600090815260046020526040902054929350608060020a9091046001604060020a031690910301808214156119bd576119bd612f19565b60019650505050505b50509392505050565b6103e881565b600060028060005460a860020a900460ff1660028111156119f257fe5b146119fc57600080fd5b670c7d713b49da000034811115611a1257600080fd5b600180546000908152600660209081526040808320600480845282852054855292528220015460ff161591611a4e908990899089908087612c2b565b90508015611bd2576001805460009081526007602090815260408083206003548452825280832081516101a08101835281546001604060020a038082168352604060020a808304821696840196909652608060020a80830482169584019590955260c060020a9182900481166060840152968301548088166080840152948504871660a0830152928404861660c08201529190920490931660e0840152600281015460ff80821615156101008087019190915282048116151561012086015262010000820481161515610140860152630100000082048116151561016086015264010000000090910416151561018084015291611b4a90612ecf565b825460018054600090815260046020526040902054929350608060020a9091046001604060020a03169091030180821415611b8757611b87613204565b600154604080519182526020820186905280517f041d165424b249ff6b188265c7ea6f322c4f47171efbac0768888b5fa3ac13d59281900390910190a16001975050505050506119c6565b506000979650505050505050565b600190565b60008115611c0057600a805484908110611bfb57fe5b506000525b6009805484908110611c0e57fe5b6000918252602090912060059091020154604860020a900460ff1690505b92915050565b600381565b60006702c68af0bb14000034811115611c4f57600080fd5b600080611c64600a600c898989600180612272565b60408051838152336020820152600160a060020a038b168183015260608101839052608081018a905260a08101899052600160c0820181905260e082015290519294509092507f3362048aaeec35587d767ff0979557a3eceaa0ec93ab377113afde2d2fc026f091908190036101000190a15060019695505050505050565b6009805482908110611cf157fe5b6000918252602090912060059091020180546001820154600283015460038401546004909401546001604060020a0384169550604060020a840460ff90811695604860020a86048216956a010000000000000000000081048316956b010000000000000000000000820490931694606060020a9091046001608060020a031693600160a060020a039384169390911691908a565b60095490565b600c805482908110611d9957fe5b600091825260209091206003909102018054600182015460029092015460ff821693506001604060020a036101008304811693604860020a9093041691600160a060020a03169085565b60005460a860020a900460ff1681565b610e1081565b600080546101009004600160a060020a03163314611e1657600080fd5b611e2883600160a060020a03166133ed565b1515611e3357600080fd5b600160a060020a038381166000908152601160205260409020541615611e5857600080fd5b50600160a060020a039182166000908152601160205260409020805473ffffffffffffffffffffffffffffffffffffffff191691909216179055600190565b6001548114155b919050565b601160205260009081526040902054600160a060020a031681565b600081565b600080546101009004600160a060020a03163314611ee057600080fd5b60018060005460a860020a900460ff166002811115611efb57fe5b14611f0557600080fd5b67016345785d8a000034811115611f1b57600080fd5b611f23612aa5565b506000611f368787876001600080612c2b565b600154604080519182526020820183905280519293507f041d165424b249ff6b188265c7ea6f322c4f47171efbac0768888b5fa3ac13d592918290030190a1600154600090815260076020908152604080832060035484529091529020600281015462010000900460ff161515611fac57600080fd5b60005460ff16151561201b576001546000908152600660209081526040808320858452909152812060030154600b8054909160c060020a90046001604060020a0316908110611ff757fe5b9060005260206000209060030201905080600201548814151561201957600080fd5b505b600154600090815260046020526040902054815460c060020a90046001604060020a0316141561204d5761204d613204565b506001979650505050505050565b600080546101009004600160a060020a0316331461207857600080fd5b60028060005460a860020a900460ff16600281111561209357fe5b141561209e57600080fd5b6000805475ff000000000000000000000000000000000000000000191675020000000000000000000000000000000000000000001790556120dd6133f5565b600191505b5090565b600b805482908110611d9957fe5b60046020526000908152604090205481565b6000811561212157600a80548490811061211c57fe5b506000525b600980548490811061212f57fe5b60009182526020909120600590910201546a0100000000000000000000900460ff169392505050565b6000612163336133ed565b151561216e57600080fd5b33600090815260116020526040902054600160a060020a03161561219157600080fd5b503360009081526011602052604090208054600160a060020a03831673ffffffffffffffffffffffffffffffffffffffff199091161790556001919050565b600a805482908110611cf157fe5b6006602090815260009283526040808420909152908252902080546001820154600283015460038401546004909401549293919290916001604060020a0380821692604060020a8304821692608060020a810483169260c060020a909104169060ff80821691610100810482169162010000820481169163010000008104821691640100000000909104168c565b600e5481565b60008083156122ba578261229d576122983467016345785d8a000063ffffffff6135ca16565b6122b5565b6122b5346702c68af0bb14000063ffffffff6135ca16565b6122bc565b345b9050600160a060020a038716331480156122d557508015155b80156122df575085155b80156122e9575084155b806123225750600160a060020a03878116600090815260116020526040902054161580159061231757508515155b801561232257508415155b151561232d57600080fd5b885461233c8a60018301614ac7565b91506000898381548110151561234e57fe5b600091825260209091206005909102016001810180543373ffffffffffffffffffffffffffffffffffffffff19918216179091558154600283018054909216600160a060020a038c161790915560038201899055600482018890557fffffffff00000000000000000000000000000000ffffffffffffffffffffffff16606060020a6001608060020a038516021767ffffffffffffffff1916426001604060020a03161768ff00000000000000001916604060020a87151590810291909117825590915061250c57805469ff0000000000000000001916604860020a178155600160a060020a038816331461250c57604080516101408101825282546001604060020a0381168252604060020a810460ff90811615156020840152604860020a820481161515938301939093526a010000000000000000000081048316151560608301526b010000000000000000000000810490921615156080820152606060020a9091046001608060020a031660a08201526001820154600160a060020a0390811660c083015260028301541660e08201526003820154610100820152600482015461012082015261250190846127f0565b151561250c57600080fd5b885460009015156125305789546125268b60018301614af8565b5060009050612538565b508854600019015b60008a8281548110151561254857fe5b90600052602060002090600302019050612561816135df565b805460ff1680612594575080546001604060020a0361010082048116604860020a909204811691909103600101166103e8145b156125ee578a548b906125aa8260018301614af8565b815481106125b457fe5b60009182526020909120600390910201805468ffffffffffffffff0019166101006001604060020a0388160217815590506125ee816135df565b805470ffffffffffffffff0000000000000000001916604860020a6001604060020a0387160217815533600160a060020a038b16141561270357604080516101408101825284546001604060020a0381168252604060020a810460ff90811615156020840152604860020a820481161515938301939093526a010000000000000000000081048316151560608301526b010000000000000000000000810490921615156080820152606060020a9091046001608060020a031660a08201526001840154600160a060020a0390811660c083015260028501541660e0820152600384015461010082015260048401546101208201526126fe906126f0908c61364b565b82908763ffffffff6136ad16565b6127e1565b600160a060020a038a81166000908152601160209081526040918290205482516101408101845287546001604060020a0381168252604060020a810460ff908116151594830194909452604860020a810484161515948201949094526a010000000000000000000084048316151560608201526b010000000000000000000000840490921615156080830152606060020a9092046001608060020a031660a08201526001860154831660c08201526002860154831660e0820152600386015461010082015260048601546101208201526127e1926126f0921661364b565b50505097509795505050505050565b60008260e00151600160a060020a03168360c00151600160a060020a03161415612868578260e00151600160a060020a03166108fc8460a001516001608060020a03169081150290604051600060405180830381858888f1935050505015801561285e573d6000803e3d6000fd5b5060019050611c2c565b60e083015160208085015160c0860151610100870151610120880151604080517fd9afd3a9000000000000000000000000000000000000000000000000000000008152941515600486015260248501899052600160a060020a039384166044860152606485019290925260848401525193169263d9afd3a99260a4808401939192918290030181600087803b15801561290057600080fd5b505af1158015612914573d6000803e3d6000fd5b505050506040513d602081101561292a57600080fd5b50519392505050565b61293b614a5f565b6060612957600961294b85613909565b9063ffffffff61392c16565b905061297a81600081518110151561296b57fe5b906020019060200201516139ba565b6001604060020a031682528051612998908290600190811061296b57fe5b602083015280516129b0908290600290811061296b57fe5b6001604060020a0316604083015280516129e090829060039081106129d157fe5b906020019060200201516139de565b600160a060020a031660608301528051612a01908290600490811061296b57fe5b60808301528051612a289082906005908110612a1957fe5b906020019060200201516139fa565b60a08301528051612a40908290600690811061296b57fe5b60c08301528051612a58908290600790811061296b57fe5b60e08301528051612a70908290600890811061296b57fe5b61010083015250919050565b60008160c001516000148015612a94575060e0820151155b8015611c2c57505061010001511590565b6000600260005460a860020a900460ff166002811115612ac157fe5b1415612acf575060006114d3565b60018054600090815260056020908152604080832054600490925290912054910190811115612b025760009150506114d3565b6001546000908152600660209081526040808320848452909152902060048101546301000000900460ff1615612b3d576000925050506114d3565b600481015460ff1615612b8d576003810154426001604060020a03608060020a909204919091166001011115612b78576000925050506114d3565b612b828183613a52565b6001925050506114d3565b600381015460016001604060020a03604060020a90920482160116612bb181613abf565b15612be0576003820154612bd490604060020a90046001604060020a0316613c02565b600193505050506114d3565b4260038360030160109054906101000a90046001604060020a03166001604060020a0316011115612c1757600093505050506114d3565b612c218284613a52565b6001935050505090565b6000838015612c375750825b8015612c405750815b15612da6576000612c5c60018054613d6390919063ffffffff16565b6000818152600260208181526040808420546007835281852081865290925290922090810154929350909162010000900460ff168015612cb757506001810154426001604060020a0360c060020a90920491909116610e1001105b15612d0157506000828152600760209081526040808320938352928152828220828155600181018390556002908101805464ffffffffff1916905593825292909252812055612ec5565b6001838155600084815260076020818152604080842060028352818520548552825280842054600019890185529282528084208785528252928390209093018054608060020a9092046001604060020a031667ffffffffffffffff1990921682179055815186815292830181905281519096507f18013fce596c7fc273e36aaa4d9ba6f94db4e483239db76e94fe9eb6769df2a89281900390910190a1505050612dcb565b60018054600090815260046020526040902054612dc89163ffffffff613d6316565b90505b6001546000908152600760209081526040808320600354845290915290206002015460ff6201000090910416151584151514612e0657600080fd5b60018054600090815260066020908152604080832085845282528083208b81558085018b9055600281018a9055600380820180546001604060020a03428116608060020a0277ffffffffffffffff000000000000000000000000000000001990921691909117808355925416604060020a026fffffffffffffffff000000000000000019909216919091179055600490810180548915156101000261ff00198c151560ff19909316929092179190911617905593548352929052208190555b9695505050505050565b600081606001516001604060020a031682604001516001604060020a03161415612efb57506000611e9e565b81604001518260600151036001016001604060020a03169050919050565b6003805460019081019182905580546000818152600760209081526040808320958352948152848220928252600490529290922054612f5d9163ffffffff613d6316565b81546001604060020a0391909116608060020a0277ffffffffffffffff000000000000000000000000000000001990911617815560028101805461ff001962ff000019909116620100001716610100179055612fb881613d7c565b600281015460ff161561303c578054600160c060020a0377ffffffffffffffff00000000000000000000000000000000196fffffffffffffffff00000000000000001983166001604060020a03938416604060020a0217908116608060020a918290048416600019018416820217918216910490911660c060020a021781556130c0565b60095481546fffffffffffffffff00000000000000001916604060020a6000199092016001604060020a0390811683029190911780845560019261308f92828116919092048216038301166103e8613ff4565b82546001604060020a03608060020a8204811692909201929092031660c060020a02600160c060020a039091161781555b600080546001919075ff000000000000000000000000000000000000000000191660a860020a8302179055506000546040517f551dc40198cc79684bb69e4931dba4ac16e4598792ee1c0a5000aeea366d7bb69160a860020a900460ff16908082600281111561312c57fe5b60ff16815260200191505060405180910390a1600180546003546002840154845460408051948552602085019390935260ff9091161515838301526001604060020a03608060020a82048116606085015260c060020a82048116608085015280821660a0850152604060020a9091041660c083015260e0820192909252600061010082015290517fe0bc0e53be29ccc81a897770a6c14b5b81310036fabf148eb7e92723c3d79e1b918190036101200190a1600281015460ff16156131f8576131f3613204565b613201565b6132018161402e565b50565b600180546000908152600460205260408120549091613229919063ffffffff613d6316565b600380546001908101918290558054600090815260076020908152604080832094835293905291822060028101805461ff00191661010017905580546001604060020a03808616608060020a0277ffffffffffffffff00000000000000000000000000000000199092169190911780835560085460001990870101821660c060020a908102600160c060020a03928316178455938301805442909316909402911617909155815492935091819075ff000000000000000000000000000000000000000000191660a860020a8202179055506000546040517f551dc40198cc79684bb69e4931dba4ac16e4598792ee1c0a5000aeea366d7bb69160a860020a900460ff16908082600281111561333a57fe5b60ff16815260200191505060405180910390a16001546003546002830154835460408051948552602085019390935260ff9091161515838301526001604060020a03608060020a82048116606085015260c060020a82048116608085015280821660a0850152604060020a9091041660c0830152600060e08301819052610100830152517fe0bc0e53be29ccc81a897770a6c14b5b81310036fabf148eb7e92723c3d79e1b918190036101200190a15050565b6000903b1190565b60018054600081815260056020908152604080832054600683528184208185528352818420600381015495870180865260028552838620604060020a9097046001604060020a03169687905560078552838620878752909452919093209294909391929083141561347457805467ffffffffffffffff191681556134c2565b6001805460009081526007602090815260408083206002835281842054845290915290205482546001604060020a03604060020a90920482169092011667ffffffffffffffff199091161781555b60028101805463ff0000001962ff0000199091166201000017166301000000179055600181810180546001604060020a03421660c060020a02600160c060020a03909116179055600a5461351b9163ffffffff6135ca16565b81546fffffffffffffffff00000000000000001916604060020a6001604060020a0392831681029190911777ffffffffffffffff000000000000000000000000000000001916608060020a60018981018516919091029190911780855590926135949282821692048116919091038301166103e8613ff4565b82546001604060020a03608060020a8204811692909201929092031660c060020a02600160c060020a0390911617905550505050565b6000828211156135d957600080fd5b50900390565b6001810154600160a060020a03161515613201576135fb614b24565b604051809103906000f080158015613617573d6000803e3d6000fd5b50600191909101805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b613653614b34565b60208084015115159082015260c080840151600160a060020a039081169183019190915260a0808501516001608060020a03169083015261010080850151908301526101209384015193820193909352911660e082015290565b6001830154600160a060020a031615156136c657600080fd5b604080516020808252818301909252606091602082018180388339505085549192506000916137059150849061010090046001604060020a03166135ca565b9050806020830152606061372361371e868660006141a7565b6141d6565b6001870154604080517f20ba5b6000000000000000000000000000000000000000000000000000000000815260048101918252865160448201528651939450600160a060020a03909216926320ba5b60928792869290918291602482019160640190602087019080838360005b838110156137a8578181015183820152602001613790565b50505050905090810190601f1680156137d55780820380516001836020036101000a031916815260200191505b50838103825284518152845160209182019186019080838360005b838110156138085781810151838201526020016137f0565b50505050905090810190601f1680156138355780820380516001836020036101000a031916815260200191505b50945050505050600060405180830381600087803b15801561385657600080fd5b505af115801561386a573d6000803e3d6000fd5b5050506001870154604080517f80759f1f0000000000000000000000000000000000000000000000000000000081529051600160a060020a0390921692506380759f1f916004808301926020929190829003018186803b1580156138cd57600080fd5b505afa1580156138e1573d6000803e3d6000fd5b505050506040513d60208110156138f757600080fd5b50516002909601959095555050505050565b613911614b88565b50805160408051808201909152602092830181529182015290565b60608160405190808252806020026020018201604052801561396857816020015b613955614b88565b81526020019060019003908161394d5790505b509050613973614b9f565b61397c8461438c565b905060005b838110156139b257613992826143b1565b83828151811015156139a057fe5b60209081029091010152600101613981565b505092915050565b60008060006139c8846143e3565b90516020919091036101000a9004949350505050565b6000806139ea836143e3565b5051606060020a90049392505050565b6020810151606090801515613a0f5750611e9e565b806040519080825280601f01601f191660200182016040528015613a3a576020820181803883390190505b509150613a4c83600001518383614442565b50919050565b60048201805464ff000000001916640100000000179055600180546000908152600560209081526040918290208490559154815190815291820183905280517ffb96205e4b3633fd57aa805b26b51ecf528714a10241a4af015929dce86768d99281900390910190a15050565b6000600354821115613ad357506000611e9e565b60015460009081526007602090815260408083208584529091529020600281015462010000900460ff161515613b0d576000915050611e9e565b600281015460ff1615613b3d576001908101544260c060020a9091046001604060020a0316909101119050611e9e565b6001546000908152600460205260409020548154608060020a90046001604060020a03161115613b71576000915050611e9e565b60015460009081526006602090815260408083208454608060020a90046001604060020a0316845290915290206004810154640100000000900460ff1615613bbe57600192505050611e9e565b60048101546301000000900460ff1615613bdd57600092505050611e9e565b60030154426001608060020a9092046001604060020a03169190910111159392505050565b6001546000908152600760209081526040808320848452909152812080549091608060020a9091046001604060020a0316905b825460c060020a90046001604060020a03168211613cbe576001546000908152600660209081526040808320858452909152902060048101546301000000900460ff1680613c8d5750600481015462010000900460ff165b15613c9c576001915050613cbe565b600401805464ff00000000191664010000000017905560019190910190613c35565b600081613ccb5782613cd0565b600183035b8454909150608060020a90046001604060020a03168110613d5c5760018054600090815260056020908152604091829020849055915486548251918252928101889052608060020a9092046001604060020a03168282015260608201839052517f70801d4d63b3da6c19ba7349911f45bed5a99ccdfb51b8138c105872529bebd5916080908290030190a15b5050505050565b600082820183811015613d7557600080fd5b9392505050565b60035460021415613da4576009541515613da45760028101805460ff19166001179055613201565b60015460009081526007602090815260408083206003546001190184529091529020600281015460ff1615613e30578054825467ffffffffffffffff19166001604060020a03604060020a9283900481169190911784556001808401549085018054918490049092169092026fffffffffffffffff000000000000000019909216919091179055613fcb565b8054825467ffffffffffffffff1916604060020a918290046001604060020a0390811660019081018216929092178555604080516101a0810182528554808416825285810484166020830152608060020a80820485169383019390935260c060020a9081900484166060830152848701548085166080840152958604841660a0830152918504831660c082015293041660e0830152600283015460ff80821615156101008086019190915282048116151561012085015262010000820481161515610140850152630100000082048116151561016085015264010000000090910416151561018083015290613f2490612ecf565b60018381015485820180546fffffffffffffffff00000000000000001916604060020a928390046001604060020a03908116959095019590950384169091029390931790925581546000908152600660209081526040808320865460c060020a9081900486168552925290912060030154600b8054909392909104909116908110613fab57fe5b60009182526020909120600390910201805460ff19169115159190911790555b81546009546001604060020a03909116141561178a5760028201805460ff191660011790555050565b6000818381151561400157fe5b061561401b57818381151561401257fe5b04600101613d75565b818381151561402657fe5b049392505050565b60005b604080516101a08101825283546001604060020a038082168352604060020a80830482166020850152608060020a80840483169585019590955260c060020a928390048216606085015260018701548083166080860152908104821660a0850152938404811660c0840152920490911660e0820152600283015460ff80821615156101008085019190915282048116151561012084015262010000820481161515610140840152630100000082048116151561016084015264010000000090910416151561018082015261410490612ecf565b816001604060020a0316101561178a5760018054600090815260066020818152604080842087546001604060020a03608060020a918290048116890181168752918452828620600401805460ff191688179055888701548754875294845282862089549190910482168801821686529092529092206003018054604060020a9092048316850190921660c060020a02600160c060020a0390911617905501614031565b6141af614a5f565b60016020820152620186a060408201526141ca848484614480565b60a08201529392505050565b6040805160098082526101408201909252606091829190816020015b60608152602001906001900390816141f2575050835190915061421d906001604060020a031661456b565b81600081518110151561422c57fe5b90602001906020020181905250614246836020015161456b565b81600181518110151561425557fe5b602090810290910101526040830151614276906001604060020a031661456b565b81600281518110151561428557fe5b6020908102909101015260608301516142a690600160a060020a031661457e565b8160038151811015156142b557fe5b6020908102909101015260808301516142cd9061456b565b8160048151811015156142dc57fe5b6020908102909101015260a08301516142f4906145ae565b81600581518110151561430357fe5b6020908102909101015260c083015161431b9061456b565b81600681518110151561432a57fe5b6020908102909101015260e08301516143429061456b565b81600781518110151561435157fe5b6020908102909101015261010083015161436a9061456b565b81600881518110151561437957fe5b60209081029091010152613d7581614631565b614394614b9f565b600061439f83614692565b83519383529092016020820152919050565b6143b9614b88565b602082015160006143c9826146e3565b828452602080850182905292019390910192909252919050565b805180516000918291821a9060808210156144055792506001915061443d9050565b60b8821015614423576001856020015103925080600101935061443a565b602085015182820160b51901945082900360b60192505b50505b915091565b6020601f820104836020840160005b8381101561446d57602081028381015190830152600101614451565b5060008551602001860152505050505050565b60606000826144af577fe904e3d9000000000000000000000000000000000000000000000000000000006144d1565b7fd9afd3a9000000000000000000000000000000000000000000000000000000005b60208681015160c088015161010089015161012090990151604080517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19969096169486019490945291151560f860020a0260288501526029840197909752600160a060020a03909616606060020a026049830152605d820196909652607d8082019590955285518082039095018552609d01909452509092915050565b6060611c2c6145798361470f565b6145ae565b60408051741400000000000000000000000000000000000000008318601482015260348101909152606090613d75815b60608151600114801561460d575081517f7f0000000000000000000000000000000000000000000000000000000000000090839060009081106145ed57fe5b90602001015160f860020a900460f860020a02600160f860020a03191611155b15614619575080611e9e565b611c2c61462b8351608060ff16614838565b83614960565b604080516000808252602082019092526060915b83518110156146795761466f82858381518110151561466057fe5b90602001906020020151614960565b9150600101614645565b50613d7561468c825160c060ff16614838565b82614960565b8051805160009190821a9060808210156146b157600092505050611e9e565b60b88210806146cc575060c082101580156146cc575060f882105b156146dc57600192505050611e9e565b5050919050565b8051600090811a60808110156146fc5760019150613a4c565b60b8811015613a4c57607e190192915050565b6040805160208082528183019092526060916000918391602082018180388339019050509050836020820152600091505b602082101561479e57808281518110151561475757fe5b60209101015160f860020a90819004027fff0000000000000000000000000000000000000000000000000000000000000016156147935761479e565b600190910190614740565b6060826020036040519080825280601f01601f1916602001820160405280156147ce576020820181803883390190505b50905060005b815181101561482f5782516001850194849181106147ee57fe5b90602001015160f860020a900460f860020a02828281518110151561480f57fe5b906020010190600160f860020a031916908160001a9053506001016147d4565b50949350505050565b6060604060020a83106148ac57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600e60248201527f696e70757420746f6f206c6f6e67000000000000000000000000000000000000604482015290519081900360640190fd5b6040805160018082528183019092526060916020820181803883390190505090506037841161490c5782840160f860020a028160008151811015156148ed57fe5b906020010190600160f860020a031916908160001a9053509050611c2c565b60606149178561470f565b90508381510160370160f860020a0282600081518110151561493557fe5b906020010190600160f860020a031916908160001a9053506149578282614960565b95945050505050565b60608082518451016040519080825280601f01601f191660200182016040528015614992576020820181803883390190505b5090506000805b85518110156149f55785818151811015156149b057fe5b90602001015160f860020a900460f860020a0283838151811015156149d157fe5b906020010190600160f860020a031916908160001a90535060019182019101614999565b5060005b8451811015614a55578481815181101515614a1057fe5b90602001015160f860020a900460f860020a028383815181101515614a3157fe5b906020010190600160f860020a031916908160001a905350600191820191016149f9565b5090949350505050565b6101206040519081016040528060006001604060020a031681526020016000815260200160006001604060020a031681526020016000600160a060020a0316815260200160008152602001606081526020016000815260200160008152602001600081525090565b815481835581811115614af357600502816005028360005260206000209182019101614af39190614bc0565b505050565b815481835581811115614af357600302816003028360005260206000209182019101614af39190614c34565b60405161126f80614c8983390190565b6040805161014081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c0810182905260e08101829052610100810182905261012081019190915290565b604080518082019091526000808252602082015290565b606060405190810160405280614bb3614b88565b8152602001600081525090565b6114d391905b808211156120e25780547bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916815560018101805473ffffffffffffffffffffffffffffffffffffffff1990811690915560028201805490911690556000600382018190556004820155600501614bc6565b6114d391905b808211156120e257805470ffffffffffffffffffffffffffffffffff1916815560018101805473ffffffffffffffffffffffffffffffffffffffff1916905560006002820155600301614c3a5600608060405234801561001057600080fd5b5061124f806100206000396000f3006080604052600436106100775763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166320ba5b60811461007c57806350c946fe1461009e578063693ac4fb146100d657806380759f1f14610104578063a43914da14610126578063f7e498f61461014a575b600080fd5b34801561008857600080fd5b5061009c610097366004610f92565b610177565b005b3480156100aa57600080fd5b506100be6100b9366004610e82565b61018d565b6040516100cd939291906110f5565b60405180910390f35b3480156100e257600080fd5b506100f66100f1366004610f5d565b610278565b6040516100cd929190611161565b34801561011057600080fd5b506101196104d7565b6040516100cd919061112b565b34801561013257600080fd5b5061013b6104dd565b6040516100cd93929190611139565b34801561015657600080fd5b5061016a610165366004610ea0565b610527565b6040516100cd919061111d565b6101896000838363ffffffff61067c16565b5050565b610195610cc2565b61019d610cc2565b6101a5610cc2565b6101ad610cdd565b6000858152600460209081526040808320815160608101909252909290918391908201908390600290835b828210156102265760408051808201825260038402860180548252825180840190935260018082015484526002909101546020848101919091528083019390935290835290920191016101d8565b5050509152505080515151855280515160209081015151855281515181015181015184528151810151518682015281518101518101515185820152905181015181015181015190830152509193909250565b60008054606090151561028a57600080fd5b610292610cf6565b60408051908101604052808580519060200120815260200161010081525090506102ba610d0d565b506040805180820182526001548152815180830190925260025482526003546020838101919091528101919091526102f0610d28565b6000806102fb610cf6565b610303610cf6565b600061030d610cf6565b6020880151610323908a9063ffffffff61073b16565b6020808b0151810151908301519296509094501461033d57fe5b6020830151151561034d57610454565b60208401519590950160ff81900360020a9a909a17996001019461037083610769565b8951600090815260046020526040902091935091506103d69060018490036002811061039857fe5b6040805180820182526003929092029290920180548252825180840190935260018101548352600201546020838101919091528101919091526107d0565b6001860195889061010081106103e857fe5b6020908102919091019190915288516000908152600490915260409020826002811061041057fe5b60408051808201825260039290920292909201805482528251808401909352600181015483526002015460208381019190915281019190915290985096508761030d565b60008511156104c95784604051908082528060200260200182016040528015610487578160200160208202803883390190505b50995060005b858110156104c757878161010081106104a257fe5b60200201518b828151811015156104b557fe5b6020908102909101015260010161048d565b505b505050505050505050915091565b60005490565b60008060006104ea610d0d565b5050604080518082018252600154815281518083019092526002548252600354602083810191825282018390529051915190519194909350909150565b6000610531610cf6565b6040805190810160405280878051906020012081526020016101008152509050610559610d0d565b85516020870120815260005b851561065257600061057687610812565b60ff1690508060019060020a02198716965061059e8160ff038561086690919063ffffffff16565b602085018190529094506000906105b490610769565b602086015290506105c3610cc2565b6105cc856107d0565b8183600281106105d857fe5b602002015287518890600019868203019081106105f157fe5b90602001906020020151818360010360028110151561060c57fe5b6020908102919091019190915281519181015160408051808401949094528381019190915280518084038201815260609093019052815191012084525050600101610565565b5060208101829052610663816107d0565b881461066e57600080fd5b506001979650505050505050565b610684610cf6565b506040805180820190915282516020808501919091208252610100818301528251908301206106b1610d0d565b855415156106c8576020810183905281815261070b565b60408051808201825260018801548152815180830190925260028801548252600388015460208381019190915281019190915261070890879085856108da565b90505b610714816107d0565b86558051600187015560209081015180516002880155015160039095019490945550505050565b610743610cf6565b61074b610cf6565b61075e846107598686610b14565b610866565b915091509250929050565b6000610773610cf6565b602083015160001061078457600080fd5b50508051604080518082019091528251600202815260209283015160001901928101929092527f8000000000000000000000000000000000000000000000000000000000000000900491565b80516020918201518083015190516040805180860194909452838101929092526060808401919091528151808403909101815260809092019052805191012090565b600081151561082057600080fd5b8160805b600160ff82161061085f5760001960ff821660020a0182161515610852579182019160ff811660020a909104905b600260ff90911604610824565b5050919050565b61086e610cf6565b610876610cf6565b8360200151831115801561088c57506101008311155b151561089457fe5b602082018390528215156108ab57600082526108bd565b835160ff84900360020a600119021682525b60208085015184900390820152925160029290920a909102825291565b6108e2610d0d565b6020808501518101519084015110156108f757fe5b6108ff610cf6565b610907610cf6565b6000610911610cf6565b61091f87896020015161073b565b602081015191955093506000901515610939575085610aef565b6020808a01518101519086015110610a6057602084015160011061095957fe5b610961610cdd565b8951600090815260048c0160209081526040808320815160608101909252909290918391908201908390600290835b828210156109de576040805180820182526003840286018054825282518084019093526001808201548452600290910154602084810191909152808301939093529083529092019101610990565b505050508152505090506109f185610769565b82519195509350610a15908c908660028110610a0957fe5b6020020151858b6108da565b81518560028110610a2257fe5b602090810291909101919091528a51600090815260048d019091526040812090610a4c8282610d49565b5050610a588b82610b84565b915050610aef565b610a6984610769565b9093509150610a76610cdd565b604080518082019091528881526020810184905281518560028110610a9757fe5b602002018190525060408051908101604052808b600001518152602001610ac98c602001518960200151600101610bea565b90528151600186900360028110610adc57fe5b6020020152610aeb8b82610b84565b9150505b604080519081016040528082815260200186815250955050505050505b949350505050565b6000808260200151846020015110610b30578260200151610b36565b83602001515b9050801515610b49576000915050610b7e565b825184511861010082900360020a60000316801515610b6a57509050610b7e565b610b7381610c1e565b60ff0360ff16925050505b92915050565b600080610b9083610c6c565b83515160008281526004968701602090815260409091208251815591810151805160018401558101516002830155945185015180516003830155850151805196820196909655949093015160059094019390935550919050565b610bf2610cf6565b6020830151821115610c0357600080fd5b60208084015183900390820152915160029190910a02815290565b6000811515610c2c57600080fd5b8160805b600160ff82161061085f5760ff811660020a600019810102821615610c5f579182019160ff811660020a909104905b600260ff90911604610c30565b8051600090610c8190825b60200201516107d0565b8251610c8e906001610c77565b6040516020018083815260200182815260200192505050604051602081830303815290604052805190602001209050919050565b60408051808201825290600290829080388339509192915050565b60c060405190810160405280610cf1610d73565b905290565b604080518082019091526000808252602082015290565b60408051606081019091526000815260208101610cf1610cf6565b61200060405190810160405280610100906020820280388339509192915050565b50600080825560018201819055600282018190556003820181905560048201819055600590910155565b60c0604051908101604052806002905b610d8b610d0d565b815260200190600190039081610d835790505090565b6000601f82018313610db257600080fd5b8135610dc5610dc0826111a8565b611181565b91508181835260208401935060208101905083856020840282011115610dea57600080fd5b60005b83811015610e165781610e008882610e20565b8452506020928301929190910190600101610ded565b5050505092915050565b6000610e2c82356111f1565b9392505050565b6000601f82018313610e4457600080fd5b8135610e52610dc0826111c9565b91508082526020830160208301858383011115610e6e57600080fd5b610e79838284611209565b50505092915050565b600060208284031215610e9457600080fd5b6000610b0c8484610e20565b600080600080600060a08688031215610eb857600080fd5b6000610ec48888610e20565b955050602086013567ffffffffffffffff811115610ee157600080fd5b610eed88828901610e33565b945050604086013567ffffffffffffffff811115610f0a57600080fd5b610f1688828901610e33565b9350506060610f2788828901610e20565b925050608086013567ffffffffffffffff811115610f4457600080fd5b610f5088828901610da1565b9150509295509295909350565b600060208284031215610f6f57600080fd5b813567ffffffffffffffff811115610f8657600080fd5b610b0c84828501610e33565b60008060408385031215610fa557600080fd5b823567ffffffffffffffff811115610fbc57600080fd5b610fc885828601610e33565b925050602083013567ffffffffffffffff811115610fe557600080fd5b610ff185828601610e33565b9150509250929050565b611004816111fa565b61100d826111f1565b60005b8281101561103d576110238583516110ec565b61102c826111f4565b602095909501949150600101611010565b5050505050565b600061104f82611200565b808452602084019350611061836111f4565b60005b82811015611091576110778683516110ec565b611080826111f4565b602096909601959150600101611064565b5093949350505050565b6110a4816111fa565b6110ad826111f1565b60005b8281101561103d576110c38583516110ec565b6110cc826111f4565b6020959095019491506001016110b0565b6110e681611204565b82525050565b6110e6816111f1565b60c081016111038286610ffb565b6111106040830185610ffb565b610b0c608083018461109b565b60208101610b7e82846110dd565b60208101610b7e82846110ec565b6060810161114782866110ec565b61115460208301856110ec565b610b0c60408301846110ec565b6040810161116f82856110ec565b8181036020830152610b0c8184611044565b60405181810167ffffffffffffffff811182821017156111a057600080fd5b604052919050565b600067ffffffffffffffff8211156111bf57600080fd5b5060209081020190565b600067ffffffffffffffff8211156111e057600080fd5b506020601f91909101601f19160190565b90565b60200190565b50600290565b5190565b151590565b828183375060009101525600a265627a7a7230582018b914ecdc6c0bae868236caa7c87914f895736b958199d65b8e37a35f3a8b3c6c6578706572696d656e74616cf50037a165627a7a72305820a2b8ccc4015f7841a41ce5aebcce2d00eb779b8b0c9862a87e5b46fb3425b4640029`

// DeployRootChain deploys a new Ethereum contract, binding an instance of RootChain to it.
func DeployRootChain(auth *bind.TransactOpts, backend bind.ContractBackend, _development bool, _NRBEpochLength *big.Int, _statesRoot [32]byte, _transactionsRoot [32]byte, _intermediateStatesRoot [32]byte) (common.Address, *types.Transaction, *RootChain, error) {
	parsed, err := abi.JSON(strings.NewReader(RootChainABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RootChainBin), backend, _development, _NRBEpochLength, _statesRoot, _transactionsRoot, _intermediateStatesRoot)
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
// Solidity: function EROs( uint256) constant returns(timestamp uint64, isExit bool, applied bool, finalized bool, challenged bool, value uint128, requestor address, to address, trieKey bytes32, trieValue bytes32)
func (_RootChain *RootChainCaller) EROs(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Timestamp  uint64
	IsExit     bool
	Applied    bool
	Finalized  bool
	Challenged bool
	Value      *big.Int
	Requestor  common.Address
	To         common.Address
	TrieKey    [32]byte
	TrieValue  [32]byte
}, error) {
	ret := new(struct {
		Timestamp  uint64
		IsExit     bool
		Applied    bool
		Finalized  bool
		Challenged bool
		Value      *big.Int
		Requestor  common.Address
		To         common.Address
		TrieKey    [32]byte
		TrieValue  [32]byte
	})
	out := ret
	err := _RootChain.contract.Call(opts, out, "EROs", arg0)
	return *ret, err
}

// EROs is a free data retrieval call binding the contract method 0xb443f3cc.
//
// Solidity: function EROs( uint256) constant returns(timestamp uint64, isExit bool, applied bool, finalized bool, challenged bool, value uint128, requestor address, to address, trieKey bytes32, trieValue bytes32)
func (_RootChain *RootChainSession) EROs(arg0 *big.Int) (struct {
	Timestamp  uint64
	IsExit     bool
	Applied    bool
	Finalized  bool
	Challenged bool
	Value      *big.Int
	Requestor  common.Address
	To         common.Address
	TrieKey    [32]byte
	TrieValue  [32]byte
}, error) {
	return _RootChain.Contract.EROs(&_RootChain.CallOpts, arg0)
}

// EROs is a free data retrieval call binding the contract method 0xb443f3cc.
//
// Solidity: function EROs( uint256) constant returns(timestamp uint64, isExit bool, applied bool, finalized bool, challenged bool, value uint128, requestor address, to address, trieKey bytes32, trieValue bytes32)
func (_RootChain *RootChainCallerSession) EROs(arg0 *big.Int) (struct {
	Timestamp  uint64
	IsExit     bool
	Applied    bool
	Finalized  bool
	Challenged bool
	Value      *big.Int
	Requestor  common.Address
	To         common.Address
	TrieKey    [32]byte
	TrieValue  [32]byte
}, error) {
	return _RootChain.Contract.EROs(&_RootChain.CallOpts, arg0)
}

// ERUs is a free data retrieval call binding the contract method 0xf4f31de4.
//
// Solidity: function ERUs( uint256) constant returns(timestamp uint64, isExit bool, applied bool, finalized bool, challenged bool, value uint128, requestor address, to address, trieKey bytes32, trieValue bytes32)
func (_RootChain *RootChainCaller) ERUs(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Timestamp  uint64
	IsExit     bool
	Applied    bool
	Finalized  bool
	Challenged bool
	Value      *big.Int
	Requestor  common.Address
	To         common.Address
	TrieKey    [32]byte
	TrieValue  [32]byte
}, error) {
	ret := new(struct {
		Timestamp  uint64
		IsExit     bool
		Applied    bool
		Finalized  bool
		Challenged bool
		Value      *big.Int
		Requestor  common.Address
		To         common.Address
		TrieKey    [32]byte
		TrieValue  [32]byte
	})
	out := ret
	err := _RootChain.contract.Call(opts, out, "ERUs", arg0)
	return *ret, err
}

// ERUs is a free data retrieval call binding the contract method 0xf4f31de4.
//
// Solidity: function ERUs( uint256) constant returns(timestamp uint64, isExit bool, applied bool, finalized bool, challenged bool, value uint128, requestor address, to address, trieKey bytes32, trieValue bytes32)
func (_RootChain *RootChainSession) ERUs(arg0 *big.Int) (struct {
	Timestamp  uint64
	IsExit     bool
	Applied    bool
	Finalized  bool
	Challenged bool
	Value      *big.Int
	Requestor  common.Address
	To         common.Address
	TrieKey    [32]byte
	TrieValue  [32]byte
}, error) {
	return _RootChain.Contract.ERUs(&_RootChain.CallOpts, arg0)
}

// ERUs is a free data retrieval call binding the contract method 0xf4f31de4.
//
// Solidity: function ERUs( uint256) constant returns(timestamp uint64, isExit bool, applied bool, finalized bool, challenged bool, value uint128, requestor address, to address, trieKey bytes32, trieValue bytes32)
func (_RootChain *RootChainCallerSession) ERUs(arg0 *big.Int) (struct {
	Timestamp  uint64
	IsExit     bool
	Applied    bool
	Finalized  bool
	Challenged bool
	Value      *big.Int
	Requestor  common.Address
	To         common.Address
	TrieKey    [32]byte
	TrieValue  [32]byte
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
// Solidity: function ORBs( uint256) constant returns(submitted bool, requestStart uint64, requestEnd uint64, trie address, transactionsRoot bytes32)
func (_RootChain *RootChainCaller) ORBs(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Submitted        bool
	RequestStart     uint64
	RequestEnd       uint64
	Trie             common.Address
	TransactionsRoot [32]byte
}, error) {
	ret := new(struct {
		Submitted        bool
		RequestStart     uint64
		RequestEnd       uint64
		Trie             common.Address
		TransactionsRoot [32]byte
	})
	out := ret
	err := _RootChain.contract.Call(opts, out, "ORBs", arg0)
	return *ret, err
}

// ORBs is a free data retrieval call binding the contract method 0xea7f22a8.
//
// Solidity: function ORBs( uint256) constant returns(submitted bool, requestStart uint64, requestEnd uint64, trie address, transactionsRoot bytes32)
func (_RootChain *RootChainSession) ORBs(arg0 *big.Int) (struct {
	Submitted        bool
	RequestStart     uint64
	RequestEnd       uint64
	Trie             common.Address
	TransactionsRoot [32]byte
}, error) {
	return _RootChain.Contract.ORBs(&_RootChain.CallOpts, arg0)
}

// ORBs is a free data retrieval call binding the contract method 0xea7f22a8.
//
// Solidity: function ORBs( uint256) constant returns(submitted bool, requestStart uint64, requestEnd uint64, trie address, transactionsRoot bytes32)
func (_RootChain *RootChainCallerSession) ORBs(arg0 *big.Int) (struct {
	Submitted        bool
	RequestStart     uint64
	RequestEnd       uint64
	Trie             common.Address
	TransactionsRoot [32]byte
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
// Solidity: function URBs( uint256) constant returns(submitted bool, requestStart uint64, requestEnd uint64, trie address, transactionsRoot bytes32)
func (_RootChain *RootChainCaller) URBs(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Submitted        bool
	RequestStart     uint64
	RequestEnd       uint64
	Trie             common.Address
	TransactionsRoot [32]byte
}, error) {
	ret := new(struct {
		Submitted        bool
		RequestStart     uint64
		RequestEnd       uint64
		Trie             common.Address
		TransactionsRoot [32]byte
	})
	out := ret
	err := _RootChain.contract.Call(opts, out, "URBs", arg0)
	return *ret, err
}

// URBs is a free data retrieval call binding the contract method 0xc0e86064.
//
// Solidity: function URBs( uint256) constant returns(submitted bool, requestStart uint64, requestEnd uint64, trie address, transactionsRoot bytes32)
func (_RootChain *RootChainSession) URBs(arg0 *big.Int) (struct {
	Submitted        bool
	RequestStart     uint64
	RequestEnd       uint64
	Trie             common.Address
	TransactionsRoot [32]byte
}, error) {
	return _RootChain.Contract.URBs(&_RootChain.CallOpts, arg0)
}

// URBs is a free data retrieval call binding the contract method 0xc0e86064.
//
// Solidity: function URBs( uint256) constant returns(submitted bool, requestStart uint64, requestEnd uint64, trie address, transactionsRoot bytes32)
func (_RootChain *RootChainCallerSession) URBs(arg0 *big.Int) (struct {
	Submitted        bool
	RequestStart     uint64
	RequestEnd       uint64
	Trie             common.Address
	TransactionsRoot [32]byte
}, error) {
	return _RootChain.Contract.URBs(&_RootChain.CallOpts, arg0)
}

// Blocks is a free data retrieval call binding the contract method 0xf4f911db.
//
// Solidity: function blocks( uint256,  uint256) constant returns(statesRoot bytes32, transactionsRoot bytes32, intermediateStatesRoot bytes32, forkNumber uint64, epochNumber uint64, timestamp uint64, requestBlockId uint64, isRequest bool, userActivated bool, challenged bool, challenging bool, finalized bool)
func (_RootChain *RootChainCaller) Blocks(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (struct {
	StatesRoot             [32]byte
	TransactionsRoot       [32]byte
	IntermediateStatesRoot [32]byte
	ForkNumber             uint64
	EpochNumber            uint64
	Timestamp              uint64
	RequestBlockId         uint64
	IsRequest              bool
	UserActivated          bool
	Challenged             bool
	Challenging            bool
	Finalized              bool
}, error) {
	ret := new(struct {
		StatesRoot             [32]byte
		TransactionsRoot       [32]byte
		IntermediateStatesRoot [32]byte
		ForkNumber             uint64
		EpochNumber            uint64
		Timestamp              uint64
		RequestBlockId         uint64
		IsRequest              bool
		UserActivated          bool
		Challenged             bool
		Challenging            bool
		Finalized              bool
	})
	out := ret
	err := _RootChain.contract.Call(opts, out, "blocks", arg0, arg1)
	return *ret, err
}

// Blocks is a free data retrieval call binding the contract method 0xf4f911db.
//
// Solidity: function blocks( uint256,  uint256) constant returns(statesRoot bytes32, transactionsRoot bytes32, intermediateStatesRoot bytes32, forkNumber uint64, epochNumber uint64, timestamp uint64, requestBlockId uint64, isRequest bool, userActivated bool, challenged bool, challenging bool, finalized bool)
func (_RootChain *RootChainSession) Blocks(arg0 *big.Int, arg1 *big.Int) (struct {
	StatesRoot             [32]byte
	TransactionsRoot       [32]byte
	IntermediateStatesRoot [32]byte
	ForkNumber             uint64
	EpochNumber            uint64
	Timestamp              uint64
	RequestBlockId         uint64
	IsRequest              bool
	UserActivated          bool
	Challenged             bool
	Challenging            bool
	Finalized              bool
}, error) {
	return _RootChain.Contract.Blocks(&_RootChain.CallOpts, arg0, arg1)
}

// Blocks is a free data retrieval call binding the contract method 0xf4f911db.
//
// Solidity: function blocks( uint256,  uint256) constant returns(statesRoot bytes32, transactionsRoot bytes32, intermediateStatesRoot bytes32, forkNumber uint64, epochNumber uint64, timestamp uint64, requestBlockId uint64, isRequest bool, userActivated bool, challenged bool, challenging bool, finalized bool)
func (_RootChain *RootChainCallerSession) Blocks(arg0 *big.Int, arg1 *big.Int) (struct {
	StatesRoot             [32]byte
	TransactionsRoot       [32]byte
	IntermediateStatesRoot [32]byte
	ForkNumber             uint64
	EpochNumber            uint64
	Timestamp              uint64
	RequestBlockId         uint64
	IsRequest              bool
	UserActivated          bool
	Challenged             bool
	Challenging            bool
	Finalized              bool
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

// GetEpoch is a free data retrieval call binding the contract method 0x2b25a38b.
//
// Solidity: function getEpoch(_forkNumber uint256, _epochNumber uint256) constant returns(requestStart uint64, requestEnd uint64, startBlockNumber uint64, endBlockNumber uint64, forkedBlockNumber uint64, isEmpty bool, initialized bool, isRequest bool, userActivated bool, finalized bool)
func (_RootChain *RootChainCaller) GetEpoch(opts *bind.CallOpts, _forkNumber *big.Int, _epochNumber *big.Int) (struct {
	RequestStart      uint64
	RequestEnd        uint64
	StartBlockNumber  uint64
	EndBlockNumber    uint64
	ForkedBlockNumber uint64
	IsEmpty           bool
	Initialized       bool
	IsRequest         bool
	UserActivated     bool
	Finalized         bool
}, error) {
	ret := new(struct {
		RequestStart      uint64
		RequestEnd        uint64
		StartBlockNumber  uint64
		EndBlockNumber    uint64
		ForkedBlockNumber uint64
		IsEmpty           bool
		Initialized       bool
		IsRequest         bool
		UserActivated     bool
		Finalized         bool
	})
	out := ret
	err := _RootChain.contract.Call(opts, out, "getEpoch", _forkNumber, _epochNumber)
	return *ret, err
}

// GetEpoch is a free data retrieval call binding the contract method 0x2b25a38b.
//
// Solidity: function getEpoch(_forkNumber uint256, _epochNumber uint256) constant returns(requestStart uint64, requestEnd uint64, startBlockNumber uint64, endBlockNumber uint64, forkedBlockNumber uint64, isEmpty bool, initialized bool, isRequest bool, userActivated bool, finalized bool)
func (_RootChain *RootChainSession) GetEpoch(_forkNumber *big.Int, _epochNumber *big.Int) (struct {
	RequestStart      uint64
	RequestEnd        uint64
	StartBlockNumber  uint64
	EndBlockNumber    uint64
	ForkedBlockNumber uint64
	IsEmpty           bool
	Initialized       bool
	IsRequest         bool
	UserActivated     bool
	Finalized         bool
}, error) {
	return _RootChain.Contract.GetEpoch(&_RootChain.CallOpts, _forkNumber, _epochNumber)
}

// GetEpoch is a free data retrieval call binding the contract method 0x2b25a38b.
//
// Solidity: function getEpoch(_forkNumber uint256, _epochNumber uint256) constant returns(requestStart uint64, requestEnd uint64, startBlockNumber uint64, endBlockNumber uint64, forkedBlockNumber uint64, isEmpty bool, initialized bool, isRequest bool, userActivated bool, finalized bool)
func (_RootChain *RootChainCallerSession) GetEpoch(_forkNumber *big.Int, _epochNumber *big.Int) (struct {
	RequestStart      uint64
	RequestEnd        uint64
	StartBlockNumber  uint64
	EndBlockNumber    uint64
	ForkedBlockNumber uint64
	IsEmpty           bool
	Initialized       bool
	IsRequest         bool
	UserActivated     bool
	Finalized         bool
}, error) {
	return _RootChain.Contract.GetEpoch(&_RootChain.CallOpts, _forkNumber, _epochNumber)
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

// GetRequestApplied is a free data retrieval call binding the contract method 0xae08ffe9.
//
// Solidity: function getRequestApplied(_requestId uint256, _userActivated bool) constant returns(applied bool)
func (_RootChain *RootChainCaller) GetRequestApplied(opts *bind.CallOpts, _requestId *big.Int, _userActivated bool) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _RootChain.contract.Call(opts, out, "getRequestApplied", _requestId, _userActivated)
	return *ret0, err
}

// GetRequestApplied is a free data retrieval call binding the contract method 0xae08ffe9.
//
// Solidity: function getRequestApplied(_requestId uint256, _userActivated bool) constant returns(applied bool)
func (_RootChain *RootChainSession) GetRequestApplied(_requestId *big.Int, _userActivated bool) (bool, error) {
	return _RootChain.Contract.GetRequestApplied(&_RootChain.CallOpts, _requestId, _userActivated)
}

// GetRequestApplied is a free data retrieval call binding the contract method 0xae08ffe9.
//
// Solidity: function getRequestApplied(_requestId uint256, _userActivated bool) constant returns(applied bool)
func (_RootChain *RootChainCallerSession) GetRequestApplied(_requestId *big.Int, _userActivated bool) (bool, error) {
	return _RootChain.Contract.GetRequestApplied(&_RootChain.CallOpts, _requestId, _userActivated)
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

// FinalizeRequest is a paid mutator transaction binding the contract method 0x99bd3600.
//
// Solidity: function finalizeRequest() returns(success bool)
func (_RootChain *RootChainTransactor) FinalizeRequest(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RootChain.contract.Transact(opts, "finalizeRequest")
}

// FinalizeRequest is a paid mutator transaction binding the contract method 0x99bd3600.
//
// Solidity: function finalizeRequest() returns(success bool)
func (_RootChain *RootChainSession) FinalizeRequest() (*types.Transaction, error) {
	return _RootChain.Contract.FinalizeRequest(&_RootChain.TransactOpts)
}

// FinalizeRequest is a paid mutator transaction binding the contract method 0x99bd3600.
//
// Solidity: function finalizeRequest() returns(success bool)
func (_RootChain *RootChainTransactorSession) FinalizeRequest() (*types.Transaction, error) {
	return _RootChain.Contract.FinalizeRequest(&_RootChain.TransactOpts)
}

// Forked is a paid mutator transaction binding the contract method 0xce8a2bc2.
//
// Solidity: function forked(_forkNumber uint256) returns(result bool)
func (_RootChain *RootChainTransactor) Forked(opts *bind.TransactOpts, _forkNumber *big.Int) (*types.Transaction, error) {
	return _RootChain.contract.Transact(opts, "forked", _forkNumber)
}

// Forked is a paid mutator transaction binding the contract method 0xce8a2bc2.
//
// Solidity: function forked(_forkNumber uint256) returns(result bool)
func (_RootChain *RootChainSession) Forked(_forkNumber *big.Int) (*types.Transaction, error) {
	return _RootChain.Contract.Forked(&_RootChain.TransactOpts, _forkNumber)
}

// Forked is a paid mutator transaction binding the contract method 0xce8a2bc2.
//
// Solidity: function forked(_forkNumber uint256) returns(result bool)
func (_RootChain *RootChainTransactorSession) Forked(_forkNumber *big.Int) (*types.Transaction, error) {
	return _RootChain.Contract.Forked(&_RootChain.TransactOpts, _forkNumber)
}

// MakeERU is a paid mutator transaction binding the contract method 0xb41e69dd.
//
// Solidity: function makeERU(_to address, _trieKey bytes32, _trieValue bytes32) returns(success bool)
func (_RootChain *RootChainTransactor) MakeERU(opts *bind.TransactOpts, _to common.Address, _trieKey [32]byte, _trieValue [32]byte) (*types.Transaction, error) {
	return _RootChain.contract.Transact(opts, "makeERU", _to, _trieKey, _trieValue)
}

// MakeERU is a paid mutator transaction binding the contract method 0xb41e69dd.
//
// Solidity: function makeERU(_to address, _trieKey bytes32, _trieValue bytes32) returns(success bool)
func (_RootChain *RootChainSession) MakeERU(_to common.Address, _trieKey [32]byte, _trieValue [32]byte) (*types.Transaction, error) {
	return _RootChain.Contract.MakeERU(&_RootChain.TransactOpts, _to, _trieKey, _trieValue)
}

// MakeERU is a paid mutator transaction binding the contract method 0xb41e69dd.
//
// Solidity: function makeERU(_to address, _trieKey bytes32, _trieValue bytes32) returns(success bool)
func (_RootChain *RootChainTransactorSession) MakeERU(_to common.Address, _trieKey [32]byte, _trieValue [32]byte) (*types.Transaction, error) {
	return _RootChain.Contract.MakeERU(&_RootChain.TransactOpts, _to, _trieKey, _trieValue)
}

// MapRequestableContract is a paid mutator transaction binding the contract method 0xf3aba344.
//
// Solidity: function mapRequestableContract(_target address) returns(success bool)
func (_RootChain *RootChainTransactor) MapRequestableContract(opts *bind.TransactOpts, _target common.Address) (*types.Transaction, error) {
	return _RootChain.contract.Transact(opts, "mapRequestableContract", _target)
}

// MapRequestableContract is a paid mutator transaction binding the contract method 0xf3aba344.
//
// Solidity: function mapRequestableContract(_target address) returns(success bool)
func (_RootChain *RootChainSession) MapRequestableContract(_target common.Address) (*types.Transaction, error) {
	return _RootChain.Contract.MapRequestableContract(&_RootChain.TransactOpts, _target)
}

// MapRequestableContract is a paid mutator transaction binding the contract method 0xf3aba344.
//
// Solidity: function mapRequestableContract(_target address) returns(success bool)
func (_RootChain *RootChainTransactorSession) MapRequestableContract(_target common.Address) (*types.Transaction, error) {
	return _RootChain.Contract.MapRequestableContract(&_RootChain.TransactOpts, _target)
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

// RevertBlock is a paid mutator transaction binding the contract method 0x76024792.
//
// Solidity: function revertBlock(_forkNumber uint256, _blockNumber uint256) returns()
func (_RootChain *RootChainTransactor) RevertBlock(opts *bind.TransactOpts, _forkNumber *big.Int, _blockNumber *big.Int) (*types.Transaction, error) {
	return _RootChain.contract.Transact(opts, "revertBlock", _forkNumber, _blockNumber)
}

// RevertBlock is a paid mutator transaction binding the contract method 0x76024792.
//
// Solidity: function revertBlock(_forkNumber uint256, _blockNumber uint256) returns()
func (_RootChain *RootChainSession) RevertBlock(_forkNumber *big.Int, _blockNumber *big.Int) (*types.Transaction, error) {
	return _RootChain.Contract.RevertBlock(&_RootChain.TransactOpts, _forkNumber, _blockNumber)
}

// RevertBlock is a paid mutator transaction binding the contract method 0x76024792.
//
// Solidity: function revertBlock(_forkNumber uint256, _blockNumber uint256) returns()
func (_RootChain *RootChainTransactorSession) RevertBlock(_forkNumber *big.Int, _blockNumber *big.Int) (*types.Transaction, error) {
	return _RootChain.Contract.RevertBlock(&_RootChain.TransactOpts, _forkNumber, _blockNumber)
}

// StartEnter is a paid mutator transaction binding the contract method 0x25119293.
//
// Solidity: function startEnter(_to address, _trieKey bytes32, _trieValue bytes32) returns(success bool)
func (_RootChain *RootChainTransactor) StartEnter(opts *bind.TransactOpts, _to common.Address, _trieKey [32]byte, _trieValue [32]byte) (*types.Transaction, error) {
	return _RootChain.contract.Transact(opts, "startEnter", _to, _trieKey, _trieValue)
}

// StartEnter is a paid mutator transaction binding the contract method 0x25119293.
//
// Solidity: function startEnter(_to address, _trieKey bytes32, _trieValue bytes32) returns(success bool)
func (_RootChain *RootChainSession) StartEnter(_to common.Address, _trieKey [32]byte, _trieValue [32]byte) (*types.Transaction, error) {
	return _RootChain.Contract.StartEnter(&_RootChain.TransactOpts, _to, _trieKey, _trieValue)
}

// StartEnter is a paid mutator transaction binding the contract method 0x25119293.
//
// Solidity: function startEnter(_to address, _trieKey bytes32, _trieValue bytes32) returns(success bool)
func (_RootChain *RootChainTransactorSession) StartEnter(_to common.Address, _trieKey [32]byte, _trieValue [32]byte) (*types.Transaction, error) {
	return _RootChain.Contract.StartEnter(&_RootChain.TransactOpts, _to, _trieKey, _trieValue)
}

// StartExit is a paid mutator transaction binding the contract method 0x24839704.
//
// Solidity: function startExit(_to address, _trieKey bytes32, _trieValue bytes32) returns(success bool)
func (_RootChain *RootChainTransactor) StartExit(opts *bind.TransactOpts, _to common.Address, _trieKey [32]byte, _trieValue [32]byte) (*types.Transaction, error) {
	return _RootChain.contract.Transact(opts, "startExit", _to, _trieKey, _trieValue)
}

// StartExit is a paid mutator transaction binding the contract method 0x24839704.
//
// Solidity: function startExit(_to address, _trieKey bytes32, _trieValue bytes32) returns(success bool)
func (_RootChain *RootChainSession) StartExit(_to common.Address, _trieKey [32]byte, _trieValue [32]byte) (*types.Transaction, error) {
	return _RootChain.Contract.StartExit(&_RootChain.TransactOpts, _to, _trieKey, _trieValue)
}

// StartExit is a paid mutator transaction binding the contract method 0x24839704.
//
// Solidity: function startExit(_to address, _trieKey bytes32, _trieValue bytes32) returns(success bool)
func (_RootChain *RootChainTransactorSession) StartExit(_to common.Address, _trieKey [32]byte, _trieValue [32]byte) (*types.Transaction, error) {
	return _RootChain.Contract.StartExit(&_RootChain.TransactOpts, _to, _trieKey, _trieValue)
}

// SubmitNRB is a paid mutator transaction binding the contract method 0x93248d02.
//
// Solidity: function submitNRB(_statesRoot bytes32, _transactionsRoot bytes32, _intermediateStatesRoot bytes32) returns(success bool)
func (_RootChain *RootChainTransactor) SubmitNRB(opts *bind.TransactOpts, _statesRoot [32]byte, _transactionsRoot [32]byte, _intermediateStatesRoot [32]byte) (*types.Transaction, error) {
	return _RootChain.contract.Transact(opts, "submitNRB", _statesRoot, _transactionsRoot, _intermediateStatesRoot)
}

// SubmitNRB is a paid mutator transaction binding the contract method 0x93248d02.
//
// Solidity: function submitNRB(_statesRoot bytes32, _transactionsRoot bytes32, _intermediateStatesRoot bytes32) returns(success bool)
func (_RootChain *RootChainSession) SubmitNRB(_statesRoot [32]byte, _transactionsRoot [32]byte, _intermediateStatesRoot [32]byte) (*types.Transaction, error) {
	return _RootChain.Contract.SubmitNRB(&_RootChain.TransactOpts, _statesRoot, _transactionsRoot, _intermediateStatesRoot)
}

// SubmitNRB is a paid mutator transaction binding the contract method 0x93248d02.
//
// Solidity: function submitNRB(_statesRoot bytes32, _transactionsRoot bytes32, _intermediateStatesRoot bytes32) returns(success bool)
func (_RootChain *RootChainTransactorSession) SubmitNRB(_statesRoot [32]byte, _transactionsRoot [32]byte, _intermediateStatesRoot [32]byte) (*types.Transaction, error) {
	return _RootChain.Contract.SubmitNRB(&_RootChain.TransactOpts, _statesRoot, _transactionsRoot, _intermediateStatesRoot)
}

// SubmitORB is a paid mutator transaction binding the contract method 0xe67123c4.
//
// Solidity: function submitORB(_statesRoot bytes32, _transactionsRoot bytes32, _intermediateStatesRoot bytes32) returns(success bool)
func (_RootChain *RootChainTransactor) SubmitORB(opts *bind.TransactOpts, _statesRoot [32]byte, _transactionsRoot [32]byte, _intermediateStatesRoot [32]byte) (*types.Transaction, error) {
	return _RootChain.contract.Transact(opts, "submitORB", _statesRoot, _transactionsRoot, _intermediateStatesRoot)
}

// SubmitORB is a paid mutator transaction binding the contract method 0xe67123c4.
//
// Solidity: function submitORB(_statesRoot bytes32, _transactionsRoot bytes32, _intermediateStatesRoot bytes32) returns(success bool)
func (_RootChain *RootChainSession) SubmitORB(_statesRoot [32]byte, _transactionsRoot [32]byte, _intermediateStatesRoot [32]byte) (*types.Transaction, error) {
	return _RootChain.Contract.SubmitORB(&_RootChain.TransactOpts, _statesRoot, _transactionsRoot, _intermediateStatesRoot)
}

// SubmitORB is a paid mutator transaction binding the contract method 0xe67123c4.
//
// Solidity: function submitORB(_statesRoot bytes32, _transactionsRoot bytes32, _intermediateStatesRoot bytes32) returns(success bool)
func (_RootChain *RootChainTransactorSession) SubmitORB(_statesRoot [32]byte, _transactionsRoot [32]byte, _intermediateStatesRoot [32]byte) (*types.Transaction, error) {
	return _RootChain.Contract.SubmitORB(&_RootChain.TransactOpts, _statesRoot, _transactionsRoot, _intermediateStatesRoot)
}

// SubmitURB is a paid mutator transaction binding the contract method 0x97be455d.
//
// Solidity: function submitURB(_statesRoot bytes32, _transactionsRoot bytes32, _intermediateStatesRoot bytes32) returns(success bool)
func (_RootChain *RootChainTransactor) SubmitURB(opts *bind.TransactOpts, _statesRoot [32]byte, _transactionsRoot [32]byte, _intermediateStatesRoot [32]byte) (*types.Transaction, error) {
	return _RootChain.contract.Transact(opts, "submitURB", _statesRoot, _transactionsRoot, _intermediateStatesRoot)
}

// SubmitURB is a paid mutator transaction binding the contract method 0x97be455d.
//
// Solidity: function submitURB(_statesRoot bytes32, _transactionsRoot bytes32, _intermediateStatesRoot bytes32) returns(success bool)
func (_RootChain *RootChainSession) SubmitURB(_statesRoot [32]byte, _transactionsRoot [32]byte, _intermediateStatesRoot [32]byte) (*types.Transaction, error) {
	return _RootChain.Contract.SubmitURB(&_RootChain.TransactOpts, _statesRoot, _transactionsRoot, _intermediateStatesRoot)
}

// SubmitURB is a paid mutator transaction binding the contract method 0x97be455d.
//
// Solidity: function submitURB(_statesRoot bytes32, _transactionsRoot bytes32, _intermediateStatesRoot bytes32) returns(success bool)
func (_RootChain *RootChainTransactorSession) SubmitURB(_statesRoot [32]byte, _transactionsRoot [32]byte, _intermediateStatesRoot [32]byte) (*types.Transaction, error) {
	return _RootChain.Contract.SubmitURB(&_RootChain.TransactOpts, _statesRoot, _transactionsRoot, _intermediateStatesRoot)
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
	FirstBlockNumber *big.Int
	LastBlockNumber  *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterEpochFinalized is a free log retrieval operation binding the contract event 0x70801d4d63b3da6c19ba7349911f45bed5a99ccdfb51b8138c105872529bebd5.
//
// Solidity: e EpochFinalized(forkNumber uint256, epochNumber uint256, firstBlockNumber uint256, lastBlockNumber uint256)
func (_RootChain *RootChainFilterer) FilterEpochFinalized(opts *bind.FilterOpts) (*RootChainEpochFinalizedIterator, error) {

	logs, sub, err := _RootChain.contract.FilterLogs(opts, "EpochFinalized")
	if err != nil {
		return nil, err
	}
	return &RootChainEpochFinalizedIterator{contract: _RootChain.contract, event: "EpochFinalized", logs: logs, sub: sub}, nil
}

// WatchEpochFinalized is a free log subscription operation binding the contract event 0x70801d4d63b3da6c19ba7349911f45bed5a99ccdfb51b8138c105872529bebd5.
//
// Solidity: e EpochFinalized(forkNumber uint256, epochNumber uint256, firstBlockNumber uint256, lastBlockNumber uint256)
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
	ForkNumber       *big.Int
	EpochNumber      *big.Int
	IsEmpty          bool
	StartBlockNumber *big.Int
	EndBlockNumber   *big.Int
	RequestStart     *big.Int
	RequestEnd       *big.Int
	IsRequest        bool
	UserActivated    bool
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterEpochPrepared is a free log retrieval operation binding the contract event 0xe0bc0e53be29ccc81a897770a6c14b5b81310036fabf148eb7e92723c3d79e1b.
//
// Solidity: e EpochPrepared(forkNumber uint256, epochNumber uint256, isEmpty bool, startBlockNumber uint256, endBlockNumber uint256, requestStart uint256, requestEnd uint256, isRequest bool, userActivated bool)
func (_RootChain *RootChainFilterer) FilterEpochPrepared(opts *bind.FilterOpts) (*RootChainEpochPreparedIterator, error) {

	logs, sub, err := _RootChain.contract.FilterLogs(opts, "EpochPrepared")
	if err != nil {
		return nil, err
	}
	return &RootChainEpochPreparedIterator{contract: _RootChain.contract, event: "EpochPrepared", logs: logs, sub: sub}, nil
}

// WatchEpochPrepared is a free log subscription operation binding the contract event 0xe0bc0e53be29ccc81a897770a6c14b5b81310036fabf148eb7e92723c3d79e1b.
//
// Solidity: e EpochPrepared(forkNumber uint256, epochNumber uint256, isEmpty bool, startBlockNumber uint256, endBlockNumber uint256, requestStart uint256, requestEnd uint256, isRequest bool, userActivated bool)
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
	IsExit        bool
	UserActivated bool
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRequestCreated is a free log retrieval operation binding the contract event 0x3362048aaeec35587d767ff0979557a3eceaa0ec93ab377113afde2d2fc026f0.
//
// Solidity: e RequestCreated(requestId uint256, requestor address, to address, weiAmount uint256, trieKey bytes32, trieValue bytes32, isExit bool, userActivated bool)
func (_RootChain *RootChainFilterer) FilterRequestCreated(opts *bind.FilterOpts) (*RootChainRequestCreatedIterator, error) {

	logs, sub, err := _RootChain.contract.FilterLogs(opts, "RequestCreated")
	if err != nil {
		return nil, err
	}
	return &RootChainRequestCreatedIterator{contract: _RootChain.contract, event: "RequestCreated", logs: logs, sub: sub}, nil
}

// WatchRequestCreated is a free log subscription operation binding the contract event 0x3362048aaeec35587d767ff0979557a3eceaa0ec93ab377113afde2d2fc026f0.
//
// Solidity: e RequestCreated(requestId uint256, requestor address, to address, weiAmount uint256, trieKey bytes32, trieValue bytes32, isExit bool, userActivated bool)
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
const SafeMathBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a7230582055a16021c31a45d118cd984993fd2642aaad49aa15fcbff9f296fa9598670d900029`

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
