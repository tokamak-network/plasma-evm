// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package epochhandler

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
const BMTBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a72305820e0d1b0e70713d7d662ec7db5138b10cd728cf30852c05d1b9ecc7574a08be9b50029`

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

// DataABI is the input ABI used to generate the binding from.
const DataABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"NA_TX_GAS_PRICE\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"NA_TX_GAS_LIMIT\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"APPLY_IN_ROOTCHAIN_SIGNATURE\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"APPLY_IN_CHILDCHAIN_SIGNATURE\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"NA\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// DataBin is the compiled bytecode used for deploying new contracts.
const DataBin = `0x610208610030600b82828239805160001a6073146000811461002057610022565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600436106100835763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416631927ac58811461008857806390e84f56146100a6578063a7b6ae28146100ae578063a89ca766146100c3578063ab73ff05146100cb575b600080fd5b6100906100e0565b60405161009d919061017f565b60405180910390f35b6100906100e8565b6100b66100ef565b60405161009d9190610171565b6100b6610113565b6100d3610137565b60405161009d919061015d565b633b9aca0081565b620186a081565b7fd9afd3a90000000000000000000000000000000000000000000000000000000081565b7fe904e3d90000000000000000000000000000000000000000000000000000000081565b600081565b6101458161018d565b82525050565b610145816101a6565b610145816101cb565b6020810161016b828461013c565b92915050565b6020810161016b828461014b565b6020810161016b8284610154565b73ffffffffffffffffffffffffffffffffffffffff1690565b7fffffffff000000000000000000000000000000000000000000000000000000001690565b905600a265627a7a72305820460487125101bc5bab9c30a7bc7fd2458c964362244ccebf4d517e89b4d949f36c6578706572696d656e74616cf50037`

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

// EpochHandlerABI is the input ABI used to generate the binding from.
const EpochHandlerABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"COST_URB_PREPARE\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CP_COMPUTATION\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastAppliedForkNumber\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"currentFork\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"COST_URB\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastAppliedERU\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"numEnterForORB\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"forks\",\"outputs\":[{\"name\":\"forkedBlock\",\"type\":\"uint64\"},{\"name\":\"firstEpoch\",\"type\":\"uint64\"},{\"name\":\"lastEpoch\",\"type\":\"uint64\"},{\"name\":\"firstBlock\",\"type\":\"uint64\"},{\"name\":\"lastBlock\",\"type\":\"uint64\"},{\"name\":\"lastFinalizedBlock\",\"type\":\"uint64\"},{\"name\":\"timestamp\",\"type\":\"uint64\"},{\"name\":\"firstEnterEpoch\",\"type\":\"uint64\"},{\"name\":\"lastEnterEpoch\",\"type\":\"uint64\"},{\"name\":\"nextBlockToRebase\",\"type\":\"uint64\"},{\"name\":\"rebased\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"_prepareToSubmitORB\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"prepareNREAfterURE\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"operator\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"prepareOREAfterURE\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastAppliedERO\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"firstFilledORENumber\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"development\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"COST_ERU\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"REQUEST_GAS\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"COST_NRB\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"NRELength\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CP_WITHHOLDING\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"COST_ORB\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"EROs\",\"outputs\":[{\"name\":\"timestamp\",\"type\":\"uint64\"},{\"name\":\"isExit\",\"type\":\"bool\"},{\"name\":\"isTransfer\",\"type\":\"bool\"},{\"name\":\"finalized\",\"type\":\"bool\"},{\"name\":\"challenged\",\"type\":\"bool\"},{\"name\":\"value\",\"type\":\"uint128\"},{\"name\":\"requestor\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"trieKey\",\"type\":\"bytes32\"},{\"name\":\"trieValue\",\"type\":\"bytes32\"},{\"name\":\"hash\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"URBs\",\"outputs\":[{\"name\":\"submitted\",\"type\":\"bool\"},{\"name\":\"numEnter\",\"type\":\"uint64\"},{\"name\":\"epochNumber\",\"type\":\"uint64\"},{\"name\":\"requestStart\",\"type\":\"uint64\"},{\"name\":\"requestEnd\",\"type\":\"uint64\"},{\"name\":\"trie\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PREPARE_TIMEOUT\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"COST_ERO\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"requestableContracts\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"NULL_ADDRESS\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"prepareToSubmitURB\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"epochHandler\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"ORBs\",\"outputs\":[{\"name\":\"submitted\",\"type\":\"bool\"},{\"name\":\"numEnter\",\"type\":\"uint64\"},{\"name\":\"epochNumber\",\"type\":\"uint64\"},{\"name\":\"requestStart\",\"type\":\"uint64\"},{\"name\":\"requestEnd\",\"type\":\"uint64\"},{\"name\":\"trie\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"_prepareToSubmitNRB\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"ERUs\",\"outputs\":[{\"name\":\"timestamp\",\"type\":\"uint64\"},{\"name\":\"isExit\",\"type\":\"bool\"},{\"name\":\"isTransfer\",\"type\":\"bool\"},{\"name\":\"finalized\",\"type\":\"bool\"},{\"name\":\"challenged\",\"type\":\"bool\"},{\"name\":\"value\",\"type\":\"uint128\"},{\"name\":\"requestor\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"trieKey\",\"type\":\"bytes32\"},{\"name\":\"trieValue\",\"type\":\"bytes32\"},{\"name\":\"hash\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastAppliedBlockNumber\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"userActivated\",\"type\":\"bool\"}],\"name\":\"SessionTimeout\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newFork\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"epochNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"forkedBlockNumber\",\"type\":\"uint256\"}],\"name\":\"Forked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"forkNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"epochNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"startBlockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"endBlockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"requestStart\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"requestEnd\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"epochIsEmpty\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"isRequest\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"userActivated\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"rebase\",\"type\":\"bool\"}],\"name\":\"EpochPrepared\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"forkNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"epochNumber\",\"type\":\"uint256\"}],\"name\":\"EpochFilling\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"forkNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"epochNumber\",\"type\":\"uint256\"}],\"name\":\"EpochFilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"forkNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"epochNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"startBlockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"endBlockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"requestStart\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"requestEnd\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"epochIsEmpty\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"isRequest\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"userActivated\",\"type\":\"bool\"}],\"name\":\"EpochRebased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"fork\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"epochNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"isRequest\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"userActivated\",\"type\":\"bool\"}],\"name\":\"BlockSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"requestor\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"weiAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"trieKey\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"trieValue\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"isTransfer\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"isExit\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"userActivated\",\"type\":\"bool\"}],\"name\":\"RequestCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"requestor\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"trieKey\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"trieValue\",\"type\":\"bytes32\"}],\"name\":\"ERUCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"forkNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"BlockFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"forkNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"epochNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"startBlockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"endBlockNumber\",\"type\":\"uint256\"}],\"name\":\"EpochFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"userActivated\",\"type\":\"bool\"}],\"name\":\"RequestFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"userActivated\",\"type\":\"bool\"}],\"name\":\"RequestApplied\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"userActivated\",\"type\":\"bool\"}],\"name\":\"RequestChallenged\",\"type\":\"event\"}]"

// EpochHandlerBin is the compiled bytecode used for deploying new contracts.
const EpochHandlerBin = `0x608060405234801561001057600080fd5b5060018054600160a060020a031916301790556128c9806100326000396000f3006080604052600436106101a05763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663033cfbed81146101a557806308c4fff0146101cc578063164bc2ae146101e1578063183d2d1c146101f6578063192adc5b1461020b5780631f261d591461022057806323691566146102355780634ba3a1261461024a5780634dd594b5146102d05780635656225b146102da578063570ca735146102e25780635e9ef4f31461031357806365d724bc1461031b57806372ecb9a8146103305780637b929c27146103485780638b5172d0146103715780638eb288ca1461038657806394be3aa5146101a5578063ab96da2d1461039b578063b17fa6e9146103b0578063b2ae9ba8146101a5578063b443f3cc146103c5578063c0e860641461045a578063c2bc88fa146104bd578063d691acd8146101a5578063da0185f8146104d2578063de0ce17d146104f3578063e6925d0814610508578063e7b88b8014610510578063ea7f22a814610525578063f25fff571461053d578063f4f31de414610545578063fb788a271461055d575b600080fd5b3480156101b157600080fd5b506101ba610572565b60408051918252519081900360200190f35b3480156101d857600080fd5b506101ba61057e565b3480156101ed57600080fd5b506101ba610583565b34801561020257600080fd5b506101ba610589565b34801561021757600080fd5b506101ba61058f565b34801561022c57600080fd5b506101ba61059b565b34801561024157600080fd5b506101ba6105a1565b34801561025657600080fd5b506102626004356105a7565b604080516001604060020a039c8d1681529a8c1660208c0152988b168a8a0152968a1660608a0152948916608089015292881660a088015290871660c0870152861660e086015285166101008501529093166101208301529115156101408201529051908190036101600190f35b6102d8610619565b005b6102d8610af8565b3480156102ee57600080fd5b506102f7610ce7565b60408051600160a060020a039092168252519081900360200190f35b6102d8610cfb565b34801561032757600080fd5b506101ba610f08565b34801561033c57600080fd5b506101ba600435610f0e565b34801561035457600080fd5b5061035d610f20565b604080519115158252519081900360200190f35b34801561037d57600080fd5b506101ba610f29565b34801561039257600080fd5b506101ba610f35565b3480156103a757600080fd5b506101ba610f3c565b3480156103bc57600080fd5b506101ba610f42565b3480156103d157600080fd5b506103dd600435610f47565b604080516001604060020a03909c168c5299151560208c01529715158a8a015295151560608a015293151560808901526fffffffffffffffffffffffffffffffff90921660a0880152600160a060020a0390811660c08801521660e086015261010085015261012084015261014083015251908190036101600190f35b34801561046657600080fd5b50610472600435611009565b6040805196151587526001604060020a0395861660208801529385168685015291841660608601529092166080840152600160a060020a0390911660a0830152519081900360c00190f35b3480156104c957600080fd5b506101ba611085565b3480156104de57600080fd5b506102f7600160a060020a036004351661108b565b3480156104ff57600080fd5b506102f76110a6565b6102d86110ab565b34801561051c57600080fd5b506102f76115cc565b34801561053157600080fd5b506104726004356115db565b6102d86115e9565b34801561055157600080fd5b506103dd600435611786565b34801561056957600080fd5b506101ba611794565b67016345785d8a000081565b600181565b600b5481565b60035481565b670c7d713b49da000081565b600e5481565b600a5481565b6005602052600090815260409020805460018201546002909201546001604060020a0380831693604060020a808504831694608060020a80820485169560c060020a9283900486169585811695858104821695848204831695909104821693838316939182049092169160ff9104168b565b6003546000818152600560205260408120918190819081908190158061064a57506002860154608060020a900460ff165b151561065557600080fd5b8554608060020a908190046001604060020a039081166001818101808416600090815260038c01602052604080822094825281208054855460008051602061287e8339815191521660c060020a918290048816860188169098029790971785556002850180546101006201000062ff0000199092169190911761ff0019161790559284018054600160c060020a031642871690970296909617808755600a80546fffffffffffffffff000000000000000019909216604060020a92881683021788558290559554919a50929850909650929092041611156107d857600186015460c060020a90046001604060020a0316151561077257600186018054600160c060020a031660c060020a6001604060020a038816021790556107b9565b60028601546001604060020a03908116600090815260038801602052604090206001018054918716608060020a0260008051602061287e8339815191529092169190911790555b60028601805467ffffffffffffffff19166001604060020a0387161790555b6107eb84866001604060020a031661179a565b600284015460ff16156108695783546fffffffffffffffff00000000000000001981166001604060020a03918216604060020a02178086558454600160c060020a0360008051602061287e83398151915290921660c060020a918290048416608060020a908102919091179283169204909216909102178455610932565b60065484546001604060020a036000199092018216604060020a026fffffffffffffffff00000000000000001990911617808655845460c060020a900482166001908101909216608060020a0260008051602061287e8339815191529091161785556109016108d6611b7d565b86546001604060020a03808216604060020a909204811691909103600101169063ffffffff611b8316565b85546001604060020a03608060020a8204811692909201929092031660c060020a02600160c060020a039091161784555b60035484546002860154604080519384526001604060020a03808a166020860152608060020a840481168583015260c060020a8404811660608601528084166080860152604060020a90930490921660a084015260ff808216151560c0850152600160e085015260006101008501526401000000009091041615156101208301525160008051602061285e833981519152918190036101400190a1600284015460ff1615610a0c57855460008051602061287e83398151915216608060020a6001604060020a03871602178655610a076115e9565b610af0565b610a1584611bbe565b9150600090505b81816001604060020a03161015610af05783546001906004880190600090610a5490608060020a90046001604060020a031685611c33565b6001604060020a039081168252602082019290925260400160009081206004908101805460ff191694151594909417909355600187015487549083168501938a0192610aa891608060020a90041685611c33565b6001604060020a039081168252602082019290925260400160002080546fffffffffffffffff00000000000000001916604060020a9390921692909202179055600101610a1c565b505050505050565b60035460008181526005602081905260408220928291610b3b918591908490610b2890600163ffffffff611c5116565b8152602001908152602001600020611c68565b83546003805460016001604060020a03608060020a94859004811682018082166000818152958b0160209081526040808820548151978852918701929092529690960490911683860152606083018490526080830184905260a0830184905285151560c084015260e083018490526101008301939093526101208201529151929450925060008051602061285e83398151915291908190036101400190a18115610ce2576001808401546001604060020a038381166000818152600380890160209081526040808420805460c060020a9888168902600160c060020a039091161781558b54608060020a80880260008051602061287e833981519152909216919091178d5560028d01805470ff0000000000000000000000000000000019168217905593549054825191825292810195909552918104851684830152948504841660608401528484166080840152604060020a90940490921660a082015260c081019390935260e08301819052610100830152517f030c1c69405c93021f28f57557240dee939a320b826a1fd0d39bf6e629ecab47918190036101200190a1610ce26115e9565b505050565b6000546101009004600160a060020a031681565b60035460008181526005602081905260408220928291610d40918591908490610d2b90600163ffffffff611c5116565b8152602001908152602001600020600861203f565b8354909250608060020a90046001604060020a03166001019050811515610d7f5760035460009081526004602052604090206001604060020a03821690555b600380546001604060020a0380841660008181529387016020908152604080862054815195865291850192909252608060020a810483168483015260c060020a8104831660608501528083166080850152604060020a900490911660a083015284151560c0830152600160e08301819052610100830193909352610120820192909252905160008051602061285e833981519152918190036101400190a18115610ce2576001808401546001604060020a038381166000818152600380890160209081526040808420805460c060020a9888168902600160c060020a039091161781558b54608060020a80880260008051602061287e833981519152909216919091178d5593549054825191825292810195909552918104851684830152948504841660608401528484166080840152604060020a90940490921660a082015260c0810184905260e0810193909352610100830152517f030c1c69405c93021f28f57557240dee939a320b826a1fd0d39bf6e629ecab47918190036101200190a1610ce2610af8565b600d5481565b60046020526000908152604090205481565b60005460ff1681565b6702c68af0bb14000081565b620186a081565b60025481565b600381565b6006805482908110610f5557fe5b60009182526020909120600690910201805460018201546002830154600384015460048501546005909501546001604060020a0385169650604060020a850460ff90811696690100000000000000000087048216966a010000000000000000000081048316966b0100000000000000000000008204909316956c010000000000000000000000009091046fffffffffffffffffffffffffffffffff1694600160a060020a039384169493909116929091908b565b600980548290811061101757fe5b60009182526020909120600290910201805460019091015460ff821692506001604060020a0361010083048116926901000000000000000000810482169271010000000000000000000000000000000000909104821691811690600160a060020a03604060020a9091041686565b610e1081565b600f60205260009081526040902054600160a060020a031681565b600081565b600380546000818152600560209081526040808320600180860185528285208183015483546001604060020a03604060020a9283900481168501811667ffffffffffffffff1990921691909117808655835460c060020a9183168202600160c060020a039182161780865582810484168b52600488018a52888b2054608060020a90851686026fffffffffffffffff000000000000000019909216919091178581048516820260008051602061287e833981519152918216178088558789018054428816948502931692909217909155949094049092168952998301909652938620600281018054938201805496909a02959096169490941790975563ff0000001962ff00001961ff001990921661010017919091166201000017166301000000179092559093921591808361120d5785546001604060020a03604060020a918290048116600090815260038901602052604090205460019290041601611210565b60005b835467ffffffffffffffff19166001604060020a03918216178085556007546fffffffffffffffff000000000000000019909116604060020a600019909201831682021780865590810482169116111561126657fe5b8254611285906001604060020a0380821691604060020a9004166124f4565b8554845460008051602061287e8339815191521660c060020a9091046001604060020a03908116608060020a908102929092178087559294506112e7926001926112d29291041685611c33565b6001604060020a03169063ffffffff61252e16565b83546001604060020a039190911660c060020a02600160c060020a03909116178355836113af5785546001604060020a03604060020a909104811660009081526003880160205260409020546113aa91611375916001916113609160c060020a8104821691608060020a9091041663ffffffff61252e16565b6001604060020a03169063ffffffff611c3316565b87546001604060020a03604060020a9091048116600090815260038a016020526040902060010154169063ffffffff611c3316565b6113b2565b60005b60018401805467ffffffffffffffff19166001604060020a03929092169190911790555060005b816001604060020a0316816001604060020a0316101561150f578254600190600487019060009061141a90608060020a90046001604060020a031685611c33565b6001604060020a039081168252602082019290925260400160009081206004908101805460ff191694151594909417909355855460019389019261146691608060020a90041685611c33565b6001604060020a03908116825260208201929092526040016000908120600490810180549415156101000261ff001990951694909417909355600186015486549083168501938901926114c791608060020a9004168563ffffffff611c3316565b6001604060020a039081168252602082019290925260400160002080546fffffffffffffffff00000000000000001916604060020a93909216929092021790556001016113d9565b60035485548454600286015460408051600190950185526001604060020a03604060020a9485900481166020870152608060020a840481168683015260c060020a84048116606087015283811660808701529390920490921660a0840152600060c084015260ff6201000083048116151560e0850152630100000083048116151561010085015264010000000090920490911615156101208301525160008051602061285e833981519152918190036101400190a1505050505050565b600154600160a060020a031681565b600880548290811061101757fe5b60035460008181526005602052604081209181908190158061161657506002840154608060020a900460ff165b151561162157600080fd5b5050815460016001604060020a03608060020a9092048216810191821660008181526003860160205260409020929350811461168b57508254608060020a90046001604060020a03908116600090815260038501602052604090205460c060020a90048116600101165b6002808301805461ff001916610100908117918290556001850180546001604060020a0342811660c060020a908102600160c060020a03938416179093558754878216608060020a90810260008051602061287e83398151915290921691909117808a559654880160001901821684029690921695909517808855600354604080519182528a88166020830152928204871681840152929004909416606082015260006080820181905260a0820181905260c0820181905260e082018190529181019190915264010000000090910460ff161515610120820152905160008051602061285e833981519152918190036101400190a150505050565b6007805482908110610f5557fe5b600c5481565b600654600090819015156117bc5760028401805460ff19166001179055611b77565b600354600090815260056020526040902091506001198301151561185557600854600010156118405760088054600191906117fd908363ffffffff611c5116565b8154811061180757fe5b60009182526020808320600292909202909101805460ff1916931515939093179092556003548152600490915260409020839055611850565b60028401805460ff191660011790555b611b77565b506001198201600090815260038201602052604090206002820154608060020a900460ff16801561189c5750815460046001604060020a03604060020a9092048216011683145b156118b857506002198201600090815260038201602052604090205b600281015462010000900460ff1615156118d157600080fd5b8054600654604060020a9091046001604060020a031660001990910114156119035760028401805460ff191660011790555b600284015460ff161580156119275750600354600090815260046020526040902054155b156119415760035460009081526004602052604090208390555b600284015460ff161515611aea576003546000908152600460205260409020548314801561197b57506002810154640100000000900460ff165b156119cd578054845460016001604060020a03604060020a90930483168101831667ffffffffffffffff1992831617875580840154878201805491851690920190931692909116919091179055611ae5565b600354600090815260046020526040902054831415611a27578054845467ffffffffffffffff199081166001604060020a03604060020a909304831617865560018084015490870180549092169216919091179055611ae5565b600281015460ff161515611a9c57611a3e8161254a565b8154855467ffffffffffffffff19166001604060020a039182169290920116178455611a6981611bbe565b600182810154908601805467ffffffffffffffff19166001604060020a0392831693909301909116919091179055611ae5565b8054845460016001604060020a03604060020a90930483168101831667ffffffffffffffff19928316178755808401548782018054918516909201909316929091169190911790555b611b27565b8054845467ffffffffffffffff199081166001604060020a03604060020a9093048316178655600180840154908701805490921692169190911790555b600284015460ff161515611b77576008805460019190611b4d908363ffffffff611c5116565b81548110611b5757fe5b60009182526020909120600290910201805460ff19169115159190911790555b50505050565b6103e890565b60008183811515611b9057fe5b0615611baa578183811515611ba157fe5b04600101611bb7565b8183811515611bb557fe5b045b9392505050565b600281015460009060ff1680611bfa57506002820154640100000000900460ff168015611bfa5750815460c060020a90046001604060020a0316155b15611c0757506000611c2e565b5080546001604060020a03608060020a8204811660c060020a909204811691909103600101165b919050565b60008282016001604060020a038085169082161015611bb757600080fd5b60008083831115611c6157600080fd5b5050900390565b60008060008060008660020160109054906101000a900460ff16151515611c8e57600080fd5b8654608060020a90046001604060020a0316600090815260038801602052604090206002810154909450640100000000900460ff168015611cd95750600284015462010000900460ff165b8015611cf1575060028401546301000000900460ff16155b1515611cfc57600080fd5b8660010160009054906101000a90046001604060020a03168760030160008960000160109054906101000a90046001604060020a03166001604060020a0316815260200190815260200160002060000160186101000a8154816001604060020a0302191690836001604060020a031602179055508660000160109054906101000a90046001604060020a031660010192508560040160008760000160009054906101000a90046001604060020a03166001604060020a0316815260200190815260200160002060000160009054906101000a90046001604060020a03166001604060020a03169150866003016000846001604060020a03168152602001908152602001600020935060018460020160016101000a81548160ff02191690831515021790555060018460020160046101000a81548160ff021916908315150217905550428460010160186101000a8154816001604060020a0302191690836001604060020a0316021790555085600301600083815260200190815260200160002060020160029054906101000a900460ff1615611e9b5781600101611e9d565b815b6000818152600388016020526040902060020154909150610100900460ff161515611f455760018781018054865460008051602061287e833981519152166001604060020a03918216608060020a908102919091178089559254600160c060020a039093169290911660c060020a029190911786556002808701805460ff1916841790558901805470ff00000000000000000000000000000000191690911790559450612035565b600081815260038701602052604090206002015462010000900460ff1615611f6957fe5b600187810154611f8a916001604060020a039091169063ffffffff611c3316565b84546001604060020a0391909116608060020a0260008051602061287e833981519152909116178455808214611fe0576000818152600387016020526040902054608060020a90046001604060020a0316611fec565b85546001604060020a03165b6002880180546fffffffffffffffff00000000000000001916604060020a6001604060020a0393841681029190911791829055885483169104909116101561203057fe5b600094505b5050505092915050565b6000806000806000806000808a60020160109054906101000a900460ff1615151561206957600080fd5b8a54608060020a90046001604060020a0316600090815260038c0160205260409020600281015490975062010000900460ff1680156120b3575060028701546301000000900460ff165b15156120be57600080fd5b8a60000160109054906101000a90046001604060020a031660010195508960040160008b60000160009054906101000a90046001604060020a03166001604060020a0316815260200190815260200160002060000160009054906101000a90046001604060020a03166001604060020a031694508a6003016000876001604060020a03168152602001908152602001600020965060018760020160016101000a81548160ff02191690831515021790555060018760020160026101000a81548160ff02191690831515021790555060018760020160046101000a81548160ff021916908315150217905550428760010160186101000a8154816001604060020a0302191690836001604060020a0316021790555089600301600086815260200190815260200160002060020160029054906101000a900460ff166122055784600101612207565b845b93506122116125b5565b89546fffffffffffffffff00000000000000001916604060020a6001604060020a03928316021767ffffffffffffffff19908116928216929092178a5560018a0180549092169216919091179055600084815260038b016020526040902060020154610100900460ff1615156122df57600287018054600160ff1990911681179091558b81018054895460008051602061287e833981519152166001604060020a03918216608060020a0217808b559154600160c060020a03909216911660c060020a0217885597506124e6565b600084815260038b01602052604090206002015462010000900460ff16151561230457fe5b6002808b0154908801805460ff19166001604060020a039092168711919091179081905560ff161561237d5760018b81018054895460008051602061287e833981519152166001604060020a03918216608060020a0217808b559154600160c060020a03909216911660c060020a0217885597506124e6565b8392505b600084815260038b016020526040902060010154604060020a90046001604060020a031615156123b657600283019250612381565b60018b8101546123d7916001604060020a039091169063ffffffff611c3316565b875460008051602061287e83398151915216608060020a6001604060020a039283168102919091178955600085815260038d01602090815260408083205493909304841680835260048f01909152919020548b549194508b92604060020a9091041690811061244257fe5b906000526020600020906002020190505b805461010090046001604060020a031615156124b557600191909101600081815260048b01602052604090205489549192918a91604060020a90046001604060020a03169081106124a057fe5b90600052602060002090600202019050612453565b60028b0180546fffffffffffffffff00000000000000001916604060020a6001604060020a03851602179055600097505b505050505050509392505050565b6000611bb7612501611b7d565b6125226001612516868863ffffffff611c5116565b9063ffffffff61280d16565b9063ffffffff611b8316565b6000806001604060020a038085169084161115611c6157600080fd5b600281015460009060ff168061258657506002820154640100000000900460ff1680156125865750815460c060020a90046001604060020a0316155b1561259357506000611c2e565b505460016001604060020a03808316604060020a909304811692909203011690565b6000806000806000806000806000806125da6001600354611c5190919063ffffffff16565b96505b600087815260056020526040902095506125f68661281f565b86546001604060020a03608060020a9091048116600101811660009081526003890160205260409020600201549116955062010000900460ff1661264b578554608060020a90046001604060020a0316612661565b8554608060020a90046001604060020a03166001015b6001604060020a031693505b8385111561271d5761268687600163ffffffff611c5116565b600081815260056020526040902090975095506126a28661281f565b86546001604060020a03608060020a9091048116600101811660009081526003890160205260409020600201549116955062010000900460ff166126f7578554608060020a90046001604060020a031661270d565b8554608060020a90046001604060020a03166001015b6001604060020a0316935061266d565b6000848152600387016020526040902060020154610100900460ff16156127e95785546001604060020a03908116600090815260048801602090815260408083205490931680835260038a019091529190206002015490935062010000900460ff16151561278c576001830192505b50506000818152600385016020908152604080832085845281842080546001604060020a03608060020a82048116875260048b0190955292909420548154604060020a9182900485169d5084169b50909104909116975090612801565b6127fa87600163ffffffff611c5116565b96506125dd565b50505050505050909192565b600082820183811015611bb757600080fd5b80546000906001604060020a0316151561283857600080fd5b5080546001604060020a0390811660009081526004830160205260409020541691905056001a69c0760aa329b76f72579129869013ebd3d41594db019c0e997b939fcb32e3ffffffffffffffff0000000000000000ffffffffffffffffffffffffffffffffa165627a7a723058205ccae7113681780607c1fc3645b87329e1f6ec38d23475ecbfeae2752cea64af0029`

// DeployEpochHandler deploys a new Ethereum contract, binding an instance of EpochHandler to it.
func DeployEpochHandler(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *EpochHandler, error) {
	parsed, err := abi.JSON(strings.NewReader(EpochHandlerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(EpochHandlerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EpochHandler{EpochHandlerCaller: EpochHandlerCaller{contract: contract}, EpochHandlerTransactor: EpochHandlerTransactor{contract: contract}, EpochHandlerFilterer: EpochHandlerFilterer{contract: contract}}, nil
}

// EpochHandler is an auto generated Go binding around an Ethereum contract.
type EpochHandler struct {
	EpochHandlerCaller     // Read-only binding to the contract
	EpochHandlerTransactor // Write-only binding to the contract
	EpochHandlerFilterer   // Log filterer for contract events
}

// EpochHandlerCaller is an auto generated read-only Go binding around an Ethereum contract.
type EpochHandlerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EpochHandlerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EpochHandlerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EpochHandlerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EpochHandlerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EpochHandlerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EpochHandlerSession struct {
	Contract     *EpochHandler     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EpochHandlerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EpochHandlerCallerSession struct {
	Contract *EpochHandlerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// EpochHandlerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EpochHandlerTransactorSession struct {
	Contract     *EpochHandlerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// EpochHandlerRaw is an auto generated low-level Go binding around an Ethereum contract.
type EpochHandlerRaw struct {
	Contract *EpochHandler // Generic contract binding to access the raw methods on
}

// EpochHandlerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EpochHandlerCallerRaw struct {
	Contract *EpochHandlerCaller // Generic read-only contract binding to access the raw methods on
}

// EpochHandlerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EpochHandlerTransactorRaw struct {
	Contract *EpochHandlerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEpochHandler creates a new instance of EpochHandler, bound to a specific deployed contract.
func NewEpochHandler(address common.Address, backend bind.ContractBackend) (*EpochHandler, error) {
	contract, err := bindEpochHandler(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EpochHandler{EpochHandlerCaller: EpochHandlerCaller{contract: contract}, EpochHandlerTransactor: EpochHandlerTransactor{contract: contract}, EpochHandlerFilterer: EpochHandlerFilterer{contract: contract}}, nil
}

// NewEpochHandlerCaller creates a new read-only instance of EpochHandler, bound to a specific deployed contract.
func NewEpochHandlerCaller(address common.Address, caller bind.ContractCaller) (*EpochHandlerCaller, error) {
	contract, err := bindEpochHandler(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EpochHandlerCaller{contract: contract}, nil
}

// NewEpochHandlerTransactor creates a new write-only instance of EpochHandler, bound to a specific deployed contract.
func NewEpochHandlerTransactor(address common.Address, transactor bind.ContractTransactor) (*EpochHandlerTransactor, error) {
	contract, err := bindEpochHandler(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EpochHandlerTransactor{contract: contract}, nil
}

// NewEpochHandlerFilterer creates a new log filterer instance of EpochHandler, bound to a specific deployed contract.
func NewEpochHandlerFilterer(address common.Address, filterer bind.ContractFilterer) (*EpochHandlerFilterer, error) {
	contract, err := bindEpochHandler(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EpochHandlerFilterer{contract: contract}, nil
}

// bindEpochHandler binds a generic wrapper to an already deployed contract.
func bindEpochHandler(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EpochHandlerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EpochHandler *EpochHandlerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _EpochHandler.Contract.EpochHandlerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EpochHandler *EpochHandlerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EpochHandler.Contract.EpochHandlerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EpochHandler *EpochHandlerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EpochHandler.Contract.EpochHandlerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EpochHandler *EpochHandlerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _EpochHandler.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EpochHandler *EpochHandlerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EpochHandler.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EpochHandler *EpochHandlerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EpochHandler.Contract.contract.Transact(opts, method, params...)
}

// COSTERO is a free data retrieval call binding the contract method 0xd691acd8.
//
// Solidity: function COST_ERO() constant returns(uint256)
func (_EpochHandler *EpochHandlerCaller) COSTERO(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _EpochHandler.contract.Call(opts, out, "COST_ERO")
	return *ret0, err
}

// COSTERO is a free data retrieval call binding the contract method 0xd691acd8.
//
// Solidity: function COST_ERO() constant returns(uint256)
func (_EpochHandler *EpochHandlerSession) COSTERO() (*big.Int, error) {
	return _EpochHandler.Contract.COSTERO(&_EpochHandler.CallOpts)
}

// COSTERO is a free data retrieval call binding the contract method 0xd691acd8.
//
// Solidity: function COST_ERO() constant returns(uint256)
func (_EpochHandler *EpochHandlerCallerSession) COSTERO() (*big.Int, error) {
	return _EpochHandler.Contract.COSTERO(&_EpochHandler.CallOpts)
}

// COSTERU is a free data retrieval call binding the contract method 0x8b5172d0.
//
// Solidity: function COST_ERU() constant returns(uint256)
func (_EpochHandler *EpochHandlerCaller) COSTERU(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _EpochHandler.contract.Call(opts, out, "COST_ERU")
	return *ret0, err
}

// COSTERU is a free data retrieval call binding the contract method 0x8b5172d0.
//
// Solidity: function COST_ERU() constant returns(uint256)
func (_EpochHandler *EpochHandlerSession) COSTERU() (*big.Int, error) {
	return _EpochHandler.Contract.COSTERU(&_EpochHandler.CallOpts)
}

// COSTERU is a free data retrieval call binding the contract method 0x8b5172d0.
//
// Solidity: function COST_ERU() constant returns(uint256)
func (_EpochHandler *EpochHandlerCallerSession) COSTERU() (*big.Int, error) {
	return _EpochHandler.Contract.COSTERU(&_EpochHandler.CallOpts)
}

// COSTNRB is a free data retrieval call binding the contract method 0x94be3aa5.
//
// Solidity: function COST_NRB() constant returns(uint256)
func (_EpochHandler *EpochHandlerCaller) COSTNRB(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _EpochHandler.contract.Call(opts, out, "COST_NRB")
	return *ret0, err
}

// COSTNRB is a free data retrieval call binding the contract method 0x94be3aa5.
//
// Solidity: function COST_NRB() constant returns(uint256)
func (_EpochHandler *EpochHandlerSession) COSTNRB() (*big.Int, error) {
	return _EpochHandler.Contract.COSTNRB(&_EpochHandler.CallOpts)
}

// COSTNRB is a free data retrieval call binding the contract method 0x94be3aa5.
//
// Solidity: function COST_NRB() constant returns(uint256)
func (_EpochHandler *EpochHandlerCallerSession) COSTNRB() (*big.Int, error) {
	return _EpochHandler.Contract.COSTNRB(&_EpochHandler.CallOpts)
}

// COSTORB is a free data retrieval call binding the contract method 0xb2ae9ba8.
//
// Solidity: function COST_ORB() constant returns(uint256)
func (_EpochHandler *EpochHandlerCaller) COSTORB(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _EpochHandler.contract.Call(opts, out, "COST_ORB")
	return *ret0, err
}

// COSTORB is a free data retrieval call binding the contract method 0xb2ae9ba8.
//
// Solidity: function COST_ORB() constant returns(uint256)
func (_EpochHandler *EpochHandlerSession) COSTORB() (*big.Int, error) {
	return _EpochHandler.Contract.COSTORB(&_EpochHandler.CallOpts)
}

// COSTORB is a free data retrieval call binding the contract method 0xb2ae9ba8.
//
// Solidity: function COST_ORB() constant returns(uint256)
func (_EpochHandler *EpochHandlerCallerSession) COSTORB() (*big.Int, error) {
	return _EpochHandler.Contract.COSTORB(&_EpochHandler.CallOpts)
}

// COSTURB is a free data retrieval call binding the contract method 0x192adc5b.
//
// Solidity: function COST_URB() constant returns(uint256)
func (_EpochHandler *EpochHandlerCaller) COSTURB(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _EpochHandler.contract.Call(opts, out, "COST_URB")
	return *ret0, err
}

// COSTURB is a free data retrieval call binding the contract method 0x192adc5b.
//
// Solidity: function COST_URB() constant returns(uint256)
func (_EpochHandler *EpochHandlerSession) COSTURB() (*big.Int, error) {
	return _EpochHandler.Contract.COSTURB(&_EpochHandler.CallOpts)
}

// COSTURB is a free data retrieval call binding the contract method 0x192adc5b.
//
// Solidity: function COST_URB() constant returns(uint256)
func (_EpochHandler *EpochHandlerCallerSession) COSTURB() (*big.Int, error) {
	return _EpochHandler.Contract.COSTURB(&_EpochHandler.CallOpts)
}

// COSTURBPREPARE is a free data retrieval call binding the contract method 0x033cfbed.
//
// Solidity: function COST_URB_PREPARE() constant returns(uint256)
func (_EpochHandler *EpochHandlerCaller) COSTURBPREPARE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _EpochHandler.contract.Call(opts, out, "COST_URB_PREPARE")
	return *ret0, err
}

// COSTURBPREPARE is a free data retrieval call binding the contract method 0x033cfbed.
//
// Solidity: function COST_URB_PREPARE() constant returns(uint256)
func (_EpochHandler *EpochHandlerSession) COSTURBPREPARE() (*big.Int, error) {
	return _EpochHandler.Contract.COSTURBPREPARE(&_EpochHandler.CallOpts)
}

// COSTURBPREPARE is a free data retrieval call binding the contract method 0x033cfbed.
//
// Solidity: function COST_URB_PREPARE() constant returns(uint256)
func (_EpochHandler *EpochHandlerCallerSession) COSTURBPREPARE() (*big.Int, error) {
	return _EpochHandler.Contract.COSTURBPREPARE(&_EpochHandler.CallOpts)
}

// CPCOMPUTATION is a free data retrieval call binding the contract method 0x08c4fff0.
//
// Solidity: function CP_COMPUTATION() constant returns(uint256)
func (_EpochHandler *EpochHandlerCaller) CPCOMPUTATION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _EpochHandler.contract.Call(opts, out, "CP_COMPUTATION")
	return *ret0, err
}

// CPCOMPUTATION is a free data retrieval call binding the contract method 0x08c4fff0.
//
// Solidity: function CP_COMPUTATION() constant returns(uint256)
func (_EpochHandler *EpochHandlerSession) CPCOMPUTATION() (*big.Int, error) {
	return _EpochHandler.Contract.CPCOMPUTATION(&_EpochHandler.CallOpts)
}

// CPCOMPUTATION is a free data retrieval call binding the contract method 0x08c4fff0.
//
// Solidity: function CP_COMPUTATION() constant returns(uint256)
func (_EpochHandler *EpochHandlerCallerSession) CPCOMPUTATION() (*big.Int, error) {
	return _EpochHandler.Contract.CPCOMPUTATION(&_EpochHandler.CallOpts)
}

// CPWITHHOLDING is a free data retrieval call binding the contract method 0xb17fa6e9.
//
// Solidity: function CP_WITHHOLDING() constant returns(uint256)
func (_EpochHandler *EpochHandlerCaller) CPWITHHOLDING(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _EpochHandler.contract.Call(opts, out, "CP_WITHHOLDING")
	return *ret0, err
}

// CPWITHHOLDING is a free data retrieval call binding the contract method 0xb17fa6e9.
//
// Solidity: function CP_WITHHOLDING() constant returns(uint256)
func (_EpochHandler *EpochHandlerSession) CPWITHHOLDING() (*big.Int, error) {
	return _EpochHandler.Contract.CPWITHHOLDING(&_EpochHandler.CallOpts)
}

// CPWITHHOLDING is a free data retrieval call binding the contract method 0xb17fa6e9.
//
// Solidity: function CP_WITHHOLDING() constant returns(uint256)
func (_EpochHandler *EpochHandlerCallerSession) CPWITHHOLDING() (*big.Int, error) {
	return _EpochHandler.Contract.CPWITHHOLDING(&_EpochHandler.CallOpts)
}

// EROs is a free data retrieval call binding the contract method 0xb443f3cc.
//
// Solidity: function EROs( uint256) constant returns(timestamp uint64, isExit bool, isTransfer bool, finalized bool, challenged bool, value uint128, requestor address, to address, trieKey bytes32, trieValue bytes32, hash bytes32)
func (_EpochHandler *EpochHandlerCaller) EROs(opts *bind.CallOpts, arg0 *big.Int) (struct {
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
	err := _EpochHandler.contract.Call(opts, out, "EROs", arg0)
	return *ret, err
}

// EROs is a free data retrieval call binding the contract method 0xb443f3cc.
//
// Solidity: function EROs( uint256) constant returns(timestamp uint64, isExit bool, isTransfer bool, finalized bool, challenged bool, value uint128, requestor address, to address, trieKey bytes32, trieValue bytes32, hash bytes32)
func (_EpochHandler *EpochHandlerSession) EROs(arg0 *big.Int) (struct {
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
	return _EpochHandler.Contract.EROs(&_EpochHandler.CallOpts, arg0)
}

// EROs is a free data retrieval call binding the contract method 0xb443f3cc.
//
// Solidity: function EROs( uint256) constant returns(timestamp uint64, isExit bool, isTransfer bool, finalized bool, challenged bool, value uint128, requestor address, to address, trieKey bytes32, trieValue bytes32, hash bytes32)
func (_EpochHandler *EpochHandlerCallerSession) EROs(arg0 *big.Int) (struct {
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
	return _EpochHandler.Contract.EROs(&_EpochHandler.CallOpts, arg0)
}

// ERUs is a free data retrieval call binding the contract method 0xf4f31de4.
//
// Solidity: function ERUs( uint256) constant returns(timestamp uint64, isExit bool, isTransfer bool, finalized bool, challenged bool, value uint128, requestor address, to address, trieKey bytes32, trieValue bytes32, hash bytes32)
func (_EpochHandler *EpochHandlerCaller) ERUs(opts *bind.CallOpts, arg0 *big.Int) (struct {
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
	err := _EpochHandler.contract.Call(opts, out, "ERUs", arg0)
	return *ret, err
}

// ERUs is a free data retrieval call binding the contract method 0xf4f31de4.
//
// Solidity: function ERUs( uint256) constant returns(timestamp uint64, isExit bool, isTransfer bool, finalized bool, challenged bool, value uint128, requestor address, to address, trieKey bytes32, trieValue bytes32, hash bytes32)
func (_EpochHandler *EpochHandlerSession) ERUs(arg0 *big.Int) (struct {
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
	return _EpochHandler.Contract.ERUs(&_EpochHandler.CallOpts, arg0)
}

// ERUs is a free data retrieval call binding the contract method 0xf4f31de4.
//
// Solidity: function ERUs( uint256) constant returns(timestamp uint64, isExit bool, isTransfer bool, finalized bool, challenged bool, value uint128, requestor address, to address, trieKey bytes32, trieValue bytes32, hash bytes32)
func (_EpochHandler *EpochHandlerCallerSession) ERUs(arg0 *big.Int) (struct {
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
	return _EpochHandler.Contract.ERUs(&_EpochHandler.CallOpts, arg0)
}

// NRELength is a free data retrieval call binding the contract method 0xab96da2d.
//
// Solidity: function NRELength() constant returns(uint256)
func (_EpochHandler *EpochHandlerCaller) NRELength(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _EpochHandler.contract.Call(opts, out, "NRELength")
	return *ret0, err
}

// NRELength is a free data retrieval call binding the contract method 0xab96da2d.
//
// Solidity: function NRELength() constant returns(uint256)
func (_EpochHandler *EpochHandlerSession) NRELength() (*big.Int, error) {
	return _EpochHandler.Contract.NRELength(&_EpochHandler.CallOpts)
}

// NRELength is a free data retrieval call binding the contract method 0xab96da2d.
//
// Solidity: function NRELength() constant returns(uint256)
func (_EpochHandler *EpochHandlerCallerSession) NRELength() (*big.Int, error) {
	return _EpochHandler.Contract.NRELength(&_EpochHandler.CallOpts)
}

// NULLADDRESS is a free data retrieval call binding the contract method 0xde0ce17d.
//
// Solidity: function NULL_ADDRESS() constant returns(address)
func (_EpochHandler *EpochHandlerCaller) NULLADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _EpochHandler.contract.Call(opts, out, "NULL_ADDRESS")
	return *ret0, err
}

// NULLADDRESS is a free data retrieval call binding the contract method 0xde0ce17d.
//
// Solidity: function NULL_ADDRESS() constant returns(address)
func (_EpochHandler *EpochHandlerSession) NULLADDRESS() (common.Address, error) {
	return _EpochHandler.Contract.NULLADDRESS(&_EpochHandler.CallOpts)
}

// NULLADDRESS is a free data retrieval call binding the contract method 0xde0ce17d.
//
// Solidity: function NULL_ADDRESS() constant returns(address)
func (_EpochHandler *EpochHandlerCallerSession) NULLADDRESS() (common.Address, error) {
	return _EpochHandler.Contract.NULLADDRESS(&_EpochHandler.CallOpts)
}

// ORBs is a free data retrieval call binding the contract method 0xea7f22a8.
//
// Solidity: function ORBs( uint256) constant returns(submitted bool, numEnter uint64, epochNumber uint64, requestStart uint64, requestEnd uint64, trie address)
func (_EpochHandler *EpochHandlerCaller) ORBs(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Submitted    bool
	NumEnter     uint64
	EpochNumber  uint64
	RequestStart uint64
	RequestEnd   uint64
	Trie         common.Address
}, error) {
	ret := new(struct {
		Submitted    bool
		NumEnter     uint64
		EpochNumber  uint64
		RequestStart uint64
		RequestEnd   uint64
		Trie         common.Address
	})
	out := ret
	err := _EpochHandler.contract.Call(opts, out, "ORBs", arg0)
	return *ret, err
}

// ORBs is a free data retrieval call binding the contract method 0xea7f22a8.
//
// Solidity: function ORBs( uint256) constant returns(submitted bool, numEnter uint64, epochNumber uint64, requestStart uint64, requestEnd uint64, trie address)
func (_EpochHandler *EpochHandlerSession) ORBs(arg0 *big.Int) (struct {
	Submitted    bool
	NumEnter     uint64
	EpochNumber  uint64
	RequestStart uint64
	RequestEnd   uint64
	Trie         common.Address
}, error) {
	return _EpochHandler.Contract.ORBs(&_EpochHandler.CallOpts, arg0)
}

// ORBs is a free data retrieval call binding the contract method 0xea7f22a8.
//
// Solidity: function ORBs( uint256) constant returns(submitted bool, numEnter uint64, epochNumber uint64, requestStart uint64, requestEnd uint64, trie address)
func (_EpochHandler *EpochHandlerCallerSession) ORBs(arg0 *big.Int) (struct {
	Submitted    bool
	NumEnter     uint64
	EpochNumber  uint64
	RequestStart uint64
	RequestEnd   uint64
	Trie         common.Address
}, error) {
	return _EpochHandler.Contract.ORBs(&_EpochHandler.CallOpts, arg0)
}

// PREPARETIMEOUT is a free data retrieval call binding the contract method 0xc2bc88fa.
//
// Solidity: function PREPARE_TIMEOUT() constant returns(uint256)
func (_EpochHandler *EpochHandlerCaller) PREPARETIMEOUT(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _EpochHandler.contract.Call(opts, out, "PREPARE_TIMEOUT")
	return *ret0, err
}

// PREPARETIMEOUT is a free data retrieval call binding the contract method 0xc2bc88fa.
//
// Solidity: function PREPARE_TIMEOUT() constant returns(uint256)
func (_EpochHandler *EpochHandlerSession) PREPARETIMEOUT() (*big.Int, error) {
	return _EpochHandler.Contract.PREPARETIMEOUT(&_EpochHandler.CallOpts)
}

// PREPARETIMEOUT is a free data retrieval call binding the contract method 0xc2bc88fa.
//
// Solidity: function PREPARE_TIMEOUT() constant returns(uint256)
func (_EpochHandler *EpochHandlerCallerSession) PREPARETIMEOUT() (*big.Int, error) {
	return _EpochHandler.Contract.PREPARETIMEOUT(&_EpochHandler.CallOpts)
}

// REQUESTGAS is a free data retrieval call binding the contract method 0x8eb288ca.
//
// Solidity: function REQUEST_GAS() constant returns(uint256)
func (_EpochHandler *EpochHandlerCaller) REQUESTGAS(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _EpochHandler.contract.Call(opts, out, "REQUEST_GAS")
	return *ret0, err
}

// REQUESTGAS is a free data retrieval call binding the contract method 0x8eb288ca.
//
// Solidity: function REQUEST_GAS() constant returns(uint256)
func (_EpochHandler *EpochHandlerSession) REQUESTGAS() (*big.Int, error) {
	return _EpochHandler.Contract.REQUESTGAS(&_EpochHandler.CallOpts)
}

// REQUESTGAS is a free data retrieval call binding the contract method 0x8eb288ca.
//
// Solidity: function REQUEST_GAS() constant returns(uint256)
func (_EpochHandler *EpochHandlerCallerSession) REQUESTGAS() (*big.Int, error) {
	return _EpochHandler.Contract.REQUESTGAS(&_EpochHandler.CallOpts)
}

// URBs is a free data retrieval call binding the contract method 0xc0e86064.
//
// Solidity: function URBs( uint256) constant returns(submitted bool, numEnter uint64, epochNumber uint64, requestStart uint64, requestEnd uint64, trie address)
func (_EpochHandler *EpochHandlerCaller) URBs(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Submitted    bool
	NumEnter     uint64
	EpochNumber  uint64
	RequestStart uint64
	RequestEnd   uint64
	Trie         common.Address
}, error) {
	ret := new(struct {
		Submitted    bool
		NumEnter     uint64
		EpochNumber  uint64
		RequestStart uint64
		RequestEnd   uint64
		Trie         common.Address
	})
	out := ret
	err := _EpochHandler.contract.Call(opts, out, "URBs", arg0)
	return *ret, err
}

// URBs is a free data retrieval call binding the contract method 0xc0e86064.
//
// Solidity: function URBs( uint256) constant returns(submitted bool, numEnter uint64, epochNumber uint64, requestStart uint64, requestEnd uint64, trie address)
func (_EpochHandler *EpochHandlerSession) URBs(arg0 *big.Int) (struct {
	Submitted    bool
	NumEnter     uint64
	EpochNumber  uint64
	RequestStart uint64
	RequestEnd   uint64
	Trie         common.Address
}, error) {
	return _EpochHandler.Contract.URBs(&_EpochHandler.CallOpts, arg0)
}

// URBs is a free data retrieval call binding the contract method 0xc0e86064.
//
// Solidity: function URBs( uint256) constant returns(submitted bool, numEnter uint64, epochNumber uint64, requestStart uint64, requestEnd uint64, trie address)
func (_EpochHandler *EpochHandlerCallerSession) URBs(arg0 *big.Int) (struct {
	Submitted    bool
	NumEnter     uint64
	EpochNumber  uint64
	RequestStart uint64
	RequestEnd   uint64
	Trie         common.Address
}, error) {
	return _EpochHandler.Contract.URBs(&_EpochHandler.CallOpts, arg0)
}

// CurrentFork is a free data retrieval call binding the contract method 0x183d2d1c.
//
// Solidity: function currentFork() constant returns(uint256)
func (_EpochHandler *EpochHandlerCaller) CurrentFork(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _EpochHandler.contract.Call(opts, out, "currentFork")
	return *ret0, err
}

// CurrentFork is a free data retrieval call binding the contract method 0x183d2d1c.
//
// Solidity: function currentFork() constant returns(uint256)
func (_EpochHandler *EpochHandlerSession) CurrentFork() (*big.Int, error) {
	return _EpochHandler.Contract.CurrentFork(&_EpochHandler.CallOpts)
}

// CurrentFork is a free data retrieval call binding the contract method 0x183d2d1c.
//
// Solidity: function currentFork() constant returns(uint256)
func (_EpochHandler *EpochHandlerCallerSession) CurrentFork() (*big.Int, error) {
	return _EpochHandler.Contract.CurrentFork(&_EpochHandler.CallOpts)
}

// Development is a free data retrieval call binding the contract method 0x7b929c27.
//
// Solidity: function development() constant returns(bool)
func (_EpochHandler *EpochHandlerCaller) Development(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _EpochHandler.contract.Call(opts, out, "development")
	return *ret0, err
}

// Development is a free data retrieval call binding the contract method 0x7b929c27.
//
// Solidity: function development() constant returns(bool)
func (_EpochHandler *EpochHandlerSession) Development() (bool, error) {
	return _EpochHandler.Contract.Development(&_EpochHandler.CallOpts)
}

// Development is a free data retrieval call binding the contract method 0x7b929c27.
//
// Solidity: function development() constant returns(bool)
func (_EpochHandler *EpochHandlerCallerSession) Development() (bool, error) {
	return _EpochHandler.Contract.Development(&_EpochHandler.CallOpts)
}

// EpochHandler is a free data retrieval call binding the contract method 0xe7b88b80.
//
// Solidity: function epochHandler() constant returns(address)
func (_EpochHandler *EpochHandlerCaller) EpochHandler(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _EpochHandler.contract.Call(opts, out, "epochHandler")
	return *ret0, err
}

// EpochHandler is a free data retrieval call binding the contract method 0xe7b88b80.
//
// Solidity: function epochHandler() constant returns(address)
func (_EpochHandler *EpochHandlerSession) EpochHandler() (common.Address, error) {
	return _EpochHandler.Contract.EpochHandler(&_EpochHandler.CallOpts)
}

// EpochHandler is a free data retrieval call binding the contract method 0xe7b88b80.
//
// Solidity: function epochHandler() constant returns(address)
func (_EpochHandler *EpochHandlerCallerSession) EpochHandler() (common.Address, error) {
	return _EpochHandler.Contract.EpochHandler(&_EpochHandler.CallOpts)
}

// FirstFilledORENumber is a free data retrieval call binding the contract method 0x72ecb9a8.
//
// Solidity: function firstFilledORENumber( uint256) constant returns(uint256)
func (_EpochHandler *EpochHandlerCaller) FirstFilledORENumber(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _EpochHandler.contract.Call(opts, out, "firstFilledORENumber", arg0)
	return *ret0, err
}

// FirstFilledORENumber is a free data retrieval call binding the contract method 0x72ecb9a8.
//
// Solidity: function firstFilledORENumber( uint256) constant returns(uint256)
func (_EpochHandler *EpochHandlerSession) FirstFilledORENumber(arg0 *big.Int) (*big.Int, error) {
	return _EpochHandler.Contract.FirstFilledORENumber(&_EpochHandler.CallOpts, arg0)
}

// FirstFilledORENumber is a free data retrieval call binding the contract method 0x72ecb9a8.
//
// Solidity: function firstFilledORENumber( uint256) constant returns(uint256)
func (_EpochHandler *EpochHandlerCallerSession) FirstFilledORENumber(arg0 *big.Int) (*big.Int, error) {
	return _EpochHandler.Contract.FirstFilledORENumber(&_EpochHandler.CallOpts, arg0)
}

// Forks is a free data retrieval call binding the contract method 0x4ba3a126.
//
// Solidity: function forks( uint256) constant returns(forkedBlock uint64, firstEpoch uint64, lastEpoch uint64, firstBlock uint64, lastBlock uint64, lastFinalizedBlock uint64, timestamp uint64, firstEnterEpoch uint64, lastEnterEpoch uint64, nextBlockToRebase uint64, rebased bool)
func (_EpochHandler *EpochHandlerCaller) Forks(opts *bind.CallOpts, arg0 *big.Int) (struct {
	ForkedBlock        uint64
	FirstEpoch         uint64
	LastEpoch          uint64
	FirstBlock         uint64
	LastBlock          uint64
	LastFinalizedBlock uint64
	Timestamp          uint64
	FirstEnterEpoch    uint64
	LastEnterEpoch     uint64
	NextBlockToRebase  uint64
	Rebased            bool
}, error) {
	ret := new(struct {
		ForkedBlock        uint64
		FirstEpoch         uint64
		LastEpoch          uint64
		FirstBlock         uint64
		LastBlock          uint64
		LastFinalizedBlock uint64
		Timestamp          uint64
		FirstEnterEpoch    uint64
		LastEnterEpoch     uint64
		NextBlockToRebase  uint64
		Rebased            bool
	})
	out := ret
	err := _EpochHandler.contract.Call(opts, out, "forks", arg0)
	return *ret, err
}

// Forks is a free data retrieval call binding the contract method 0x4ba3a126.
//
// Solidity: function forks( uint256) constant returns(forkedBlock uint64, firstEpoch uint64, lastEpoch uint64, firstBlock uint64, lastBlock uint64, lastFinalizedBlock uint64, timestamp uint64, firstEnterEpoch uint64, lastEnterEpoch uint64, nextBlockToRebase uint64, rebased bool)
func (_EpochHandler *EpochHandlerSession) Forks(arg0 *big.Int) (struct {
	ForkedBlock        uint64
	FirstEpoch         uint64
	LastEpoch          uint64
	FirstBlock         uint64
	LastBlock          uint64
	LastFinalizedBlock uint64
	Timestamp          uint64
	FirstEnterEpoch    uint64
	LastEnterEpoch     uint64
	NextBlockToRebase  uint64
	Rebased            bool
}, error) {
	return _EpochHandler.Contract.Forks(&_EpochHandler.CallOpts, arg0)
}

// Forks is a free data retrieval call binding the contract method 0x4ba3a126.
//
// Solidity: function forks( uint256) constant returns(forkedBlock uint64, firstEpoch uint64, lastEpoch uint64, firstBlock uint64, lastBlock uint64, lastFinalizedBlock uint64, timestamp uint64, firstEnterEpoch uint64, lastEnterEpoch uint64, nextBlockToRebase uint64, rebased bool)
func (_EpochHandler *EpochHandlerCallerSession) Forks(arg0 *big.Int) (struct {
	ForkedBlock        uint64
	FirstEpoch         uint64
	LastEpoch          uint64
	FirstBlock         uint64
	LastBlock          uint64
	LastFinalizedBlock uint64
	Timestamp          uint64
	FirstEnterEpoch    uint64
	LastEnterEpoch     uint64
	NextBlockToRebase  uint64
	Rebased            bool
}, error) {
	return _EpochHandler.Contract.Forks(&_EpochHandler.CallOpts, arg0)
}

// LastAppliedBlockNumber is a free data retrieval call binding the contract method 0xfb788a27.
//
// Solidity: function lastAppliedBlockNumber() constant returns(uint256)
func (_EpochHandler *EpochHandlerCaller) LastAppliedBlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _EpochHandler.contract.Call(opts, out, "lastAppliedBlockNumber")
	return *ret0, err
}

// LastAppliedBlockNumber is a free data retrieval call binding the contract method 0xfb788a27.
//
// Solidity: function lastAppliedBlockNumber() constant returns(uint256)
func (_EpochHandler *EpochHandlerSession) LastAppliedBlockNumber() (*big.Int, error) {
	return _EpochHandler.Contract.LastAppliedBlockNumber(&_EpochHandler.CallOpts)
}

// LastAppliedBlockNumber is a free data retrieval call binding the contract method 0xfb788a27.
//
// Solidity: function lastAppliedBlockNumber() constant returns(uint256)
func (_EpochHandler *EpochHandlerCallerSession) LastAppliedBlockNumber() (*big.Int, error) {
	return _EpochHandler.Contract.LastAppliedBlockNumber(&_EpochHandler.CallOpts)
}

// LastAppliedERO is a free data retrieval call binding the contract method 0x65d724bc.
//
// Solidity: function lastAppliedERO() constant returns(uint256)
func (_EpochHandler *EpochHandlerCaller) LastAppliedERO(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _EpochHandler.contract.Call(opts, out, "lastAppliedERO")
	return *ret0, err
}

// LastAppliedERO is a free data retrieval call binding the contract method 0x65d724bc.
//
// Solidity: function lastAppliedERO() constant returns(uint256)
func (_EpochHandler *EpochHandlerSession) LastAppliedERO() (*big.Int, error) {
	return _EpochHandler.Contract.LastAppliedERO(&_EpochHandler.CallOpts)
}

// LastAppliedERO is a free data retrieval call binding the contract method 0x65d724bc.
//
// Solidity: function lastAppliedERO() constant returns(uint256)
func (_EpochHandler *EpochHandlerCallerSession) LastAppliedERO() (*big.Int, error) {
	return _EpochHandler.Contract.LastAppliedERO(&_EpochHandler.CallOpts)
}

// LastAppliedERU is a free data retrieval call binding the contract method 0x1f261d59.
//
// Solidity: function lastAppliedERU() constant returns(uint256)
func (_EpochHandler *EpochHandlerCaller) LastAppliedERU(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _EpochHandler.contract.Call(opts, out, "lastAppliedERU")
	return *ret0, err
}

// LastAppliedERU is a free data retrieval call binding the contract method 0x1f261d59.
//
// Solidity: function lastAppliedERU() constant returns(uint256)
func (_EpochHandler *EpochHandlerSession) LastAppliedERU() (*big.Int, error) {
	return _EpochHandler.Contract.LastAppliedERU(&_EpochHandler.CallOpts)
}

// LastAppliedERU is a free data retrieval call binding the contract method 0x1f261d59.
//
// Solidity: function lastAppliedERU() constant returns(uint256)
func (_EpochHandler *EpochHandlerCallerSession) LastAppliedERU() (*big.Int, error) {
	return _EpochHandler.Contract.LastAppliedERU(&_EpochHandler.CallOpts)
}

// LastAppliedForkNumber is a free data retrieval call binding the contract method 0x164bc2ae.
//
// Solidity: function lastAppliedForkNumber() constant returns(uint256)
func (_EpochHandler *EpochHandlerCaller) LastAppliedForkNumber(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _EpochHandler.contract.Call(opts, out, "lastAppliedForkNumber")
	return *ret0, err
}

// LastAppliedForkNumber is a free data retrieval call binding the contract method 0x164bc2ae.
//
// Solidity: function lastAppliedForkNumber() constant returns(uint256)
func (_EpochHandler *EpochHandlerSession) LastAppliedForkNumber() (*big.Int, error) {
	return _EpochHandler.Contract.LastAppliedForkNumber(&_EpochHandler.CallOpts)
}

// LastAppliedForkNumber is a free data retrieval call binding the contract method 0x164bc2ae.
//
// Solidity: function lastAppliedForkNumber() constant returns(uint256)
func (_EpochHandler *EpochHandlerCallerSession) LastAppliedForkNumber() (*big.Int, error) {
	return _EpochHandler.Contract.LastAppliedForkNumber(&_EpochHandler.CallOpts)
}

// NumEnterForORB is a free data retrieval call binding the contract method 0x23691566.
//
// Solidity: function numEnterForORB() constant returns(uint256)
func (_EpochHandler *EpochHandlerCaller) NumEnterForORB(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _EpochHandler.contract.Call(opts, out, "numEnterForORB")
	return *ret0, err
}

// NumEnterForORB is a free data retrieval call binding the contract method 0x23691566.
//
// Solidity: function numEnterForORB() constant returns(uint256)
func (_EpochHandler *EpochHandlerSession) NumEnterForORB() (*big.Int, error) {
	return _EpochHandler.Contract.NumEnterForORB(&_EpochHandler.CallOpts)
}

// NumEnterForORB is a free data retrieval call binding the contract method 0x23691566.
//
// Solidity: function numEnterForORB() constant returns(uint256)
func (_EpochHandler *EpochHandlerCallerSession) NumEnterForORB() (*big.Int, error) {
	return _EpochHandler.Contract.NumEnterForORB(&_EpochHandler.CallOpts)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() constant returns(address)
func (_EpochHandler *EpochHandlerCaller) Operator(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _EpochHandler.contract.Call(opts, out, "operator")
	return *ret0, err
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() constant returns(address)
func (_EpochHandler *EpochHandlerSession) Operator() (common.Address, error) {
	return _EpochHandler.Contract.Operator(&_EpochHandler.CallOpts)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() constant returns(address)
func (_EpochHandler *EpochHandlerCallerSession) Operator() (common.Address, error) {
	return _EpochHandler.Contract.Operator(&_EpochHandler.CallOpts)
}

// RequestableContracts is a free data retrieval call binding the contract method 0xda0185f8.
//
// Solidity: function requestableContracts( address) constant returns(address)
func (_EpochHandler *EpochHandlerCaller) RequestableContracts(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _EpochHandler.contract.Call(opts, out, "requestableContracts", arg0)
	return *ret0, err
}

// RequestableContracts is a free data retrieval call binding the contract method 0xda0185f8.
//
// Solidity: function requestableContracts( address) constant returns(address)
func (_EpochHandler *EpochHandlerSession) RequestableContracts(arg0 common.Address) (common.Address, error) {
	return _EpochHandler.Contract.RequestableContracts(&_EpochHandler.CallOpts, arg0)
}

// RequestableContracts is a free data retrieval call binding the contract method 0xda0185f8.
//
// Solidity: function requestableContracts( address) constant returns(address)
func (_EpochHandler *EpochHandlerCallerSession) RequestableContracts(arg0 common.Address) (common.Address, error) {
	return _EpochHandler.Contract.RequestableContracts(&_EpochHandler.CallOpts, arg0)
}

// PrepareToSubmitNRB is a paid mutator transaction binding the contract method 0xf25fff57.
//
// Solidity: function _prepareToSubmitNRB() returns()
func (_EpochHandler *EpochHandlerTransactor) PrepareToSubmitNRB(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EpochHandler.contract.Transact(opts, "_prepareToSubmitNRB")
}

// PrepareToSubmitNRB is a paid mutator transaction binding the contract method 0xf25fff57.
//
// Solidity: function _prepareToSubmitNRB() returns()
func (_EpochHandler *EpochHandlerSession) PrepareToSubmitNRB() (*types.Transaction, error) {
	return _EpochHandler.Contract.PrepareToSubmitNRB(&_EpochHandler.TransactOpts)
}

// PrepareToSubmitNRB is a paid mutator transaction binding the contract method 0xf25fff57.
//
// Solidity: function _prepareToSubmitNRB() returns()
func (_EpochHandler *EpochHandlerTransactorSession) PrepareToSubmitNRB() (*types.Transaction, error) {
	return _EpochHandler.Contract.PrepareToSubmitNRB(&_EpochHandler.TransactOpts)
}

// PrepareToSubmitORB is a paid mutator transaction binding the contract method 0x4dd594b5.
//
// Solidity: function _prepareToSubmitORB() returns()
func (_EpochHandler *EpochHandlerTransactor) PrepareToSubmitORB(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EpochHandler.contract.Transact(opts, "_prepareToSubmitORB")
}

// PrepareToSubmitORB is a paid mutator transaction binding the contract method 0x4dd594b5.
//
// Solidity: function _prepareToSubmitORB() returns()
func (_EpochHandler *EpochHandlerSession) PrepareToSubmitORB() (*types.Transaction, error) {
	return _EpochHandler.Contract.PrepareToSubmitORB(&_EpochHandler.TransactOpts)
}

// PrepareToSubmitORB is a paid mutator transaction binding the contract method 0x4dd594b5.
//
// Solidity: function _prepareToSubmitORB() returns()
func (_EpochHandler *EpochHandlerTransactorSession) PrepareToSubmitORB() (*types.Transaction, error) {
	return _EpochHandler.Contract.PrepareToSubmitORB(&_EpochHandler.TransactOpts)
}

// PrepareNREAfterURE is a paid mutator transaction binding the contract method 0x5656225b.
//
// Solidity: function prepareNREAfterURE() returns()
func (_EpochHandler *EpochHandlerTransactor) PrepareNREAfterURE(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EpochHandler.contract.Transact(opts, "prepareNREAfterURE")
}

// PrepareNREAfterURE is a paid mutator transaction binding the contract method 0x5656225b.
//
// Solidity: function prepareNREAfterURE() returns()
func (_EpochHandler *EpochHandlerSession) PrepareNREAfterURE() (*types.Transaction, error) {
	return _EpochHandler.Contract.PrepareNREAfterURE(&_EpochHandler.TransactOpts)
}

// PrepareNREAfterURE is a paid mutator transaction binding the contract method 0x5656225b.
//
// Solidity: function prepareNREAfterURE() returns()
func (_EpochHandler *EpochHandlerTransactorSession) PrepareNREAfterURE() (*types.Transaction, error) {
	return _EpochHandler.Contract.PrepareNREAfterURE(&_EpochHandler.TransactOpts)
}

// PrepareOREAfterURE is a paid mutator transaction binding the contract method 0x5e9ef4f3.
//
// Solidity: function prepareOREAfterURE() returns()
func (_EpochHandler *EpochHandlerTransactor) PrepareOREAfterURE(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EpochHandler.contract.Transact(opts, "prepareOREAfterURE")
}

// PrepareOREAfterURE is a paid mutator transaction binding the contract method 0x5e9ef4f3.
//
// Solidity: function prepareOREAfterURE() returns()
func (_EpochHandler *EpochHandlerSession) PrepareOREAfterURE() (*types.Transaction, error) {
	return _EpochHandler.Contract.PrepareOREAfterURE(&_EpochHandler.TransactOpts)
}

// PrepareOREAfterURE is a paid mutator transaction binding the contract method 0x5e9ef4f3.
//
// Solidity: function prepareOREAfterURE() returns()
func (_EpochHandler *EpochHandlerTransactorSession) PrepareOREAfterURE() (*types.Transaction, error) {
	return _EpochHandler.Contract.PrepareOREAfterURE(&_EpochHandler.TransactOpts)
}

// PrepareToSubmitURB is a paid mutator transaction binding the contract method 0xe6925d08.
//
// Solidity: function prepareToSubmitURB() returns()
func (_EpochHandler *EpochHandlerTransactor) PrepareToSubmitURB(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EpochHandler.contract.Transact(opts, "prepareToSubmitURB")
}

// PrepareToSubmitURB is a paid mutator transaction binding the contract method 0xe6925d08.
//
// Solidity: function prepareToSubmitURB() returns()
func (_EpochHandler *EpochHandlerSession) PrepareToSubmitURB() (*types.Transaction, error) {
	return _EpochHandler.Contract.PrepareToSubmitURB(&_EpochHandler.TransactOpts)
}

// PrepareToSubmitURB is a paid mutator transaction binding the contract method 0xe6925d08.
//
// Solidity: function prepareToSubmitURB() returns()
func (_EpochHandler *EpochHandlerTransactorSession) PrepareToSubmitURB() (*types.Transaction, error) {
	return _EpochHandler.Contract.PrepareToSubmitURB(&_EpochHandler.TransactOpts)
}

// EpochHandlerBlockFinalizedIterator is returned from FilterBlockFinalized and is used to iterate over the raw logs and unpacked data for BlockFinalized events raised by the EpochHandler contract.
type EpochHandlerBlockFinalizedIterator struct {
	Event *EpochHandlerBlockFinalized // Event containing the contract specifics and raw log

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
func (it *EpochHandlerBlockFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EpochHandlerBlockFinalized)
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
		it.Event = new(EpochHandlerBlockFinalized)
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
func (it *EpochHandlerBlockFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EpochHandlerBlockFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EpochHandlerBlockFinalized represents a BlockFinalized event raised by the EpochHandler contract.
type EpochHandlerBlockFinalized struct {
	ForkNumber  *big.Int
	BlockNumber *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBlockFinalized is a free log retrieval operation binding the contract event 0xfb96205e4b3633fd57aa805b26b51ecf528714a10241a4af015929dce86768d9.
//
// Solidity: e BlockFinalized(forkNumber uint256, blockNumber uint256)
func (_EpochHandler *EpochHandlerFilterer) FilterBlockFinalized(opts *bind.FilterOpts) (*EpochHandlerBlockFinalizedIterator, error) {

	logs, sub, err := _EpochHandler.contract.FilterLogs(opts, "BlockFinalized")
	if err != nil {
		return nil, err
	}
	return &EpochHandlerBlockFinalizedIterator{contract: _EpochHandler.contract, event: "BlockFinalized", logs: logs, sub: sub}, nil
}

// WatchBlockFinalized is a free log subscription operation binding the contract event 0xfb96205e4b3633fd57aa805b26b51ecf528714a10241a4af015929dce86768d9.
//
// Solidity: e BlockFinalized(forkNumber uint256, blockNumber uint256)
func (_EpochHandler *EpochHandlerFilterer) WatchBlockFinalized(opts *bind.WatchOpts, sink chan<- *EpochHandlerBlockFinalized) (event.Subscription, error) {

	logs, sub, err := _EpochHandler.contract.WatchLogs(opts, "BlockFinalized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EpochHandlerBlockFinalized)
				if err := _EpochHandler.contract.UnpackLog(event, "BlockFinalized", log); err != nil {
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

// EpochHandlerBlockSubmittedIterator is returned from FilterBlockSubmitted and is used to iterate over the raw logs and unpacked data for BlockSubmitted events raised by the EpochHandler contract.
type EpochHandlerBlockSubmittedIterator struct {
	Event *EpochHandlerBlockSubmitted // Event containing the contract specifics and raw log

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
func (it *EpochHandlerBlockSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EpochHandlerBlockSubmitted)
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
		it.Event = new(EpochHandlerBlockSubmitted)
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
func (it *EpochHandlerBlockSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EpochHandlerBlockSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EpochHandlerBlockSubmitted represents a BlockSubmitted event raised by the EpochHandler contract.
type EpochHandlerBlockSubmitted struct {
	Fork          *big.Int
	EpochNumber   *big.Int
	BlockNumber   *big.Int
	IsRequest     bool
	UserActivated bool
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterBlockSubmitted is a free log retrieval operation binding the contract event 0x3d4a04291c66b06f39a4ecb817875b12b5485a05ec563133a56a905305c48e55.
//
// Solidity: e BlockSubmitted(fork uint256, epochNumber uint256, blockNumber uint256, isRequest bool, userActivated bool)
func (_EpochHandler *EpochHandlerFilterer) FilterBlockSubmitted(opts *bind.FilterOpts) (*EpochHandlerBlockSubmittedIterator, error) {

	logs, sub, err := _EpochHandler.contract.FilterLogs(opts, "BlockSubmitted")
	if err != nil {
		return nil, err
	}
	return &EpochHandlerBlockSubmittedIterator{contract: _EpochHandler.contract, event: "BlockSubmitted", logs: logs, sub: sub}, nil
}

// WatchBlockSubmitted is a free log subscription operation binding the contract event 0x3d4a04291c66b06f39a4ecb817875b12b5485a05ec563133a56a905305c48e55.
//
// Solidity: e BlockSubmitted(fork uint256, epochNumber uint256, blockNumber uint256, isRequest bool, userActivated bool)
func (_EpochHandler *EpochHandlerFilterer) WatchBlockSubmitted(opts *bind.WatchOpts, sink chan<- *EpochHandlerBlockSubmitted) (event.Subscription, error) {

	logs, sub, err := _EpochHandler.contract.WatchLogs(opts, "BlockSubmitted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EpochHandlerBlockSubmitted)
				if err := _EpochHandler.contract.UnpackLog(event, "BlockSubmitted", log); err != nil {
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

// EpochHandlerERUCreatedIterator is returned from FilterERUCreated and is used to iterate over the raw logs and unpacked data for ERUCreated events raised by the EpochHandler contract.
type EpochHandlerERUCreatedIterator struct {
	Event *EpochHandlerERUCreated // Event containing the contract specifics and raw log

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
func (it *EpochHandlerERUCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EpochHandlerERUCreated)
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
		it.Event = new(EpochHandlerERUCreated)
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
func (it *EpochHandlerERUCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EpochHandlerERUCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EpochHandlerERUCreated represents a ERUCreated event raised by the EpochHandler contract.
type EpochHandlerERUCreated struct {
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
func (_EpochHandler *EpochHandlerFilterer) FilterERUCreated(opts *bind.FilterOpts) (*EpochHandlerERUCreatedIterator, error) {

	logs, sub, err := _EpochHandler.contract.FilterLogs(opts, "ERUCreated")
	if err != nil {
		return nil, err
	}
	return &EpochHandlerERUCreatedIterator{contract: _EpochHandler.contract, event: "ERUCreated", logs: logs, sub: sub}, nil
}

// WatchERUCreated is a free log subscription operation binding the contract event 0xd89c6857ed5778107e858511cdc309642a48c9d1717e813a25156f535acc8d36.
//
// Solidity: e ERUCreated(requestId uint256, requestor address, to address, trieKey bytes32, trieValue bytes32)
func (_EpochHandler *EpochHandlerFilterer) WatchERUCreated(opts *bind.WatchOpts, sink chan<- *EpochHandlerERUCreated) (event.Subscription, error) {

	logs, sub, err := _EpochHandler.contract.WatchLogs(opts, "ERUCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EpochHandlerERUCreated)
				if err := _EpochHandler.contract.UnpackLog(event, "ERUCreated", log); err != nil {
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

// EpochHandlerEpochFilledIterator is returned from FilterEpochFilled and is used to iterate over the raw logs and unpacked data for EpochFilled events raised by the EpochHandler contract.
type EpochHandlerEpochFilledIterator struct {
	Event *EpochHandlerEpochFilled // Event containing the contract specifics and raw log

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
func (it *EpochHandlerEpochFilledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EpochHandlerEpochFilled)
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
		it.Event = new(EpochHandlerEpochFilled)
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
func (it *EpochHandlerEpochFilledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EpochHandlerEpochFilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EpochHandlerEpochFilled represents a EpochFilled event raised by the EpochHandler contract.
type EpochHandlerEpochFilled struct {
	ForkNumber  *big.Int
	EpochNumber *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterEpochFilled is a free log retrieval operation binding the contract event 0x2fdeb407bf5c2b621f04b5c784822dae806c45b49a68aba413cc270128c96816.
//
// Solidity: e EpochFilled(forkNumber uint256, epochNumber uint256)
func (_EpochHandler *EpochHandlerFilterer) FilterEpochFilled(opts *bind.FilterOpts) (*EpochHandlerEpochFilledIterator, error) {

	logs, sub, err := _EpochHandler.contract.FilterLogs(opts, "EpochFilled")
	if err != nil {
		return nil, err
	}
	return &EpochHandlerEpochFilledIterator{contract: _EpochHandler.contract, event: "EpochFilled", logs: logs, sub: sub}, nil
}

// WatchEpochFilled is a free log subscription operation binding the contract event 0x2fdeb407bf5c2b621f04b5c784822dae806c45b49a68aba413cc270128c96816.
//
// Solidity: e EpochFilled(forkNumber uint256, epochNumber uint256)
func (_EpochHandler *EpochHandlerFilterer) WatchEpochFilled(opts *bind.WatchOpts, sink chan<- *EpochHandlerEpochFilled) (event.Subscription, error) {

	logs, sub, err := _EpochHandler.contract.WatchLogs(opts, "EpochFilled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EpochHandlerEpochFilled)
				if err := _EpochHandler.contract.UnpackLog(event, "EpochFilled", log); err != nil {
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

// EpochHandlerEpochFillingIterator is returned from FilterEpochFilling and is used to iterate over the raw logs and unpacked data for EpochFilling events raised by the EpochHandler contract.
type EpochHandlerEpochFillingIterator struct {
	Event *EpochHandlerEpochFilling // Event containing the contract specifics and raw log

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
func (it *EpochHandlerEpochFillingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EpochHandlerEpochFilling)
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
		it.Event = new(EpochHandlerEpochFilling)
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
func (it *EpochHandlerEpochFillingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EpochHandlerEpochFillingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EpochHandlerEpochFilling represents a EpochFilling event raised by the EpochHandler contract.
type EpochHandlerEpochFilling struct {
	ForkNumber  *big.Int
	EpochNumber *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterEpochFilling is a free log retrieval operation binding the contract event 0x27b09f0953d27bbff306fe25b2987ac5a813248ac30cb2bbd5daf95e7b0e6dc0.
//
// Solidity: e EpochFilling(forkNumber uint256, epochNumber uint256)
func (_EpochHandler *EpochHandlerFilterer) FilterEpochFilling(opts *bind.FilterOpts) (*EpochHandlerEpochFillingIterator, error) {

	logs, sub, err := _EpochHandler.contract.FilterLogs(opts, "EpochFilling")
	if err != nil {
		return nil, err
	}
	return &EpochHandlerEpochFillingIterator{contract: _EpochHandler.contract, event: "EpochFilling", logs: logs, sub: sub}, nil
}

// WatchEpochFilling is a free log subscription operation binding the contract event 0x27b09f0953d27bbff306fe25b2987ac5a813248ac30cb2bbd5daf95e7b0e6dc0.
//
// Solidity: e EpochFilling(forkNumber uint256, epochNumber uint256)
func (_EpochHandler *EpochHandlerFilterer) WatchEpochFilling(opts *bind.WatchOpts, sink chan<- *EpochHandlerEpochFilling) (event.Subscription, error) {

	logs, sub, err := _EpochHandler.contract.WatchLogs(opts, "EpochFilling")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EpochHandlerEpochFilling)
				if err := _EpochHandler.contract.UnpackLog(event, "EpochFilling", log); err != nil {
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

// EpochHandlerEpochFinalizedIterator is returned from FilterEpochFinalized and is used to iterate over the raw logs and unpacked data for EpochFinalized events raised by the EpochHandler contract.
type EpochHandlerEpochFinalizedIterator struct {
	Event *EpochHandlerEpochFinalized // Event containing the contract specifics and raw log

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
func (it *EpochHandlerEpochFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EpochHandlerEpochFinalized)
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
		it.Event = new(EpochHandlerEpochFinalized)
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
func (it *EpochHandlerEpochFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EpochHandlerEpochFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EpochHandlerEpochFinalized represents a EpochFinalized event raised by the EpochHandler contract.
type EpochHandlerEpochFinalized struct {
	ForkNumber       *big.Int
	EpochNumber      *big.Int
	StartBlockNumber *big.Int
	EndBlockNumber   *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterEpochFinalized is a free log retrieval operation binding the contract event 0x70801d4d63b3da6c19ba7349911f45bed5a99ccdfb51b8138c105872529bebd5.
//
// Solidity: e EpochFinalized(forkNumber uint256, epochNumber uint256, startBlockNumber uint256, endBlockNumber uint256)
func (_EpochHandler *EpochHandlerFilterer) FilterEpochFinalized(opts *bind.FilterOpts) (*EpochHandlerEpochFinalizedIterator, error) {

	logs, sub, err := _EpochHandler.contract.FilterLogs(opts, "EpochFinalized")
	if err != nil {
		return nil, err
	}
	return &EpochHandlerEpochFinalizedIterator{contract: _EpochHandler.contract, event: "EpochFinalized", logs: logs, sub: sub}, nil
}

// WatchEpochFinalized is a free log subscription operation binding the contract event 0x70801d4d63b3da6c19ba7349911f45bed5a99ccdfb51b8138c105872529bebd5.
//
// Solidity: e EpochFinalized(forkNumber uint256, epochNumber uint256, startBlockNumber uint256, endBlockNumber uint256)
func (_EpochHandler *EpochHandlerFilterer) WatchEpochFinalized(opts *bind.WatchOpts, sink chan<- *EpochHandlerEpochFinalized) (event.Subscription, error) {

	logs, sub, err := _EpochHandler.contract.WatchLogs(opts, "EpochFinalized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EpochHandlerEpochFinalized)
				if err := _EpochHandler.contract.UnpackLog(event, "EpochFinalized", log); err != nil {
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

// EpochHandlerEpochPreparedIterator is returned from FilterEpochPrepared and is used to iterate over the raw logs and unpacked data for EpochPrepared events raised by the EpochHandler contract.
type EpochHandlerEpochPreparedIterator struct {
	Event *EpochHandlerEpochPrepared // Event containing the contract specifics and raw log

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
func (it *EpochHandlerEpochPreparedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EpochHandlerEpochPrepared)
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
		it.Event = new(EpochHandlerEpochPrepared)
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
func (it *EpochHandlerEpochPreparedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EpochHandlerEpochPreparedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EpochHandlerEpochPrepared represents a EpochPrepared event raised by the EpochHandler contract.
type EpochHandlerEpochPrepared struct {
	ForkNumber       *big.Int
	EpochNumber      *big.Int
	StartBlockNumber *big.Int
	EndBlockNumber   *big.Int
	RequestStart     *big.Int
	RequestEnd       *big.Int
	EpochIsEmpty     bool
	IsRequest        bool
	UserActivated    bool
	Rebase           bool
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterEpochPrepared is a free log retrieval operation binding the contract event 0x1a69c0760aa329b76f72579129869013ebd3d41594db019c0e997b939fcb32e3.
//
// Solidity: e EpochPrepared(forkNumber uint256, epochNumber uint256, startBlockNumber uint256, endBlockNumber uint256, requestStart uint256, requestEnd uint256, epochIsEmpty bool, isRequest bool, userActivated bool, rebase bool)
func (_EpochHandler *EpochHandlerFilterer) FilterEpochPrepared(opts *bind.FilterOpts) (*EpochHandlerEpochPreparedIterator, error) {

	logs, sub, err := _EpochHandler.contract.FilterLogs(opts, "EpochPrepared")
	if err != nil {
		return nil, err
	}
	return &EpochHandlerEpochPreparedIterator{contract: _EpochHandler.contract, event: "EpochPrepared", logs: logs, sub: sub}, nil
}

// WatchEpochPrepared is a free log subscription operation binding the contract event 0x1a69c0760aa329b76f72579129869013ebd3d41594db019c0e997b939fcb32e3.
//
// Solidity: e EpochPrepared(forkNumber uint256, epochNumber uint256, startBlockNumber uint256, endBlockNumber uint256, requestStart uint256, requestEnd uint256, epochIsEmpty bool, isRequest bool, userActivated bool, rebase bool)
func (_EpochHandler *EpochHandlerFilterer) WatchEpochPrepared(opts *bind.WatchOpts, sink chan<- *EpochHandlerEpochPrepared) (event.Subscription, error) {

	logs, sub, err := _EpochHandler.contract.WatchLogs(opts, "EpochPrepared")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EpochHandlerEpochPrepared)
				if err := _EpochHandler.contract.UnpackLog(event, "EpochPrepared", log); err != nil {
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

// EpochHandlerEpochRebasedIterator is returned from FilterEpochRebased and is used to iterate over the raw logs and unpacked data for EpochRebased events raised by the EpochHandler contract.
type EpochHandlerEpochRebasedIterator struct {
	Event *EpochHandlerEpochRebased // Event containing the contract specifics and raw log

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
func (it *EpochHandlerEpochRebasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EpochHandlerEpochRebased)
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
		it.Event = new(EpochHandlerEpochRebased)
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
func (it *EpochHandlerEpochRebasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EpochHandlerEpochRebasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EpochHandlerEpochRebased represents a EpochRebased event raised by the EpochHandler contract.
type EpochHandlerEpochRebased struct {
	ForkNumber       *big.Int
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

// FilterEpochRebased is a free log retrieval operation binding the contract event 0x030c1c69405c93021f28f57557240dee939a320b826a1fd0d39bf6e629ecab47.
//
// Solidity: e EpochRebased(forkNumber uint256, epochNumber uint256, startBlockNumber uint256, endBlockNumber uint256, requestStart uint256, requestEnd uint256, epochIsEmpty bool, isRequest bool, userActivated bool)
func (_EpochHandler *EpochHandlerFilterer) FilterEpochRebased(opts *bind.FilterOpts) (*EpochHandlerEpochRebasedIterator, error) {

	logs, sub, err := _EpochHandler.contract.FilterLogs(opts, "EpochRebased")
	if err != nil {
		return nil, err
	}
	return &EpochHandlerEpochRebasedIterator{contract: _EpochHandler.contract, event: "EpochRebased", logs: logs, sub: sub}, nil
}

// WatchEpochRebased is a free log subscription operation binding the contract event 0x030c1c69405c93021f28f57557240dee939a320b826a1fd0d39bf6e629ecab47.
//
// Solidity: e EpochRebased(forkNumber uint256, epochNumber uint256, startBlockNumber uint256, endBlockNumber uint256, requestStart uint256, requestEnd uint256, epochIsEmpty bool, isRequest bool, userActivated bool)
func (_EpochHandler *EpochHandlerFilterer) WatchEpochRebased(opts *bind.WatchOpts, sink chan<- *EpochHandlerEpochRebased) (event.Subscription, error) {

	logs, sub, err := _EpochHandler.contract.WatchLogs(opts, "EpochRebased")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EpochHandlerEpochRebased)
				if err := _EpochHandler.contract.UnpackLog(event, "EpochRebased", log); err != nil {
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

// EpochHandlerForkedIterator is returned from FilterForked and is used to iterate over the raw logs and unpacked data for Forked events raised by the EpochHandler contract.
type EpochHandlerForkedIterator struct {
	Event *EpochHandlerForked // Event containing the contract specifics and raw log

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
func (it *EpochHandlerForkedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EpochHandlerForked)
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
		it.Event = new(EpochHandlerForked)
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
func (it *EpochHandlerForkedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EpochHandlerForkedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EpochHandlerForked represents a Forked event raised by the EpochHandler contract.
type EpochHandlerForked struct {
	NewFork           *big.Int
	EpochNumber       *big.Int
	ForkedBlockNumber *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterForked is a free log retrieval operation binding the contract event 0x0647d42ab02f6e0ae76959757dcb6aa6feac1d4ba6f077f1223fb4b1b429f06c.
//
// Solidity: e Forked(newFork uint256, epochNumber uint256, forkedBlockNumber uint256)
func (_EpochHandler *EpochHandlerFilterer) FilterForked(opts *bind.FilterOpts) (*EpochHandlerForkedIterator, error) {

	logs, sub, err := _EpochHandler.contract.FilterLogs(opts, "Forked")
	if err != nil {
		return nil, err
	}
	return &EpochHandlerForkedIterator{contract: _EpochHandler.contract, event: "Forked", logs: logs, sub: sub}, nil
}

// WatchForked is a free log subscription operation binding the contract event 0x0647d42ab02f6e0ae76959757dcb6aa6feac1d4ba6f077f1223fb4b1b429f06c.
//
// Solidity: e Forked(newFork uint256, epochNumber uint256, forkedBlockNumber uint256)
func (_EpochHandler *EpochHandlerFilterer) WatchForked(opts *bind.WatchOpts, sink chan<- *EpochHandlerForked) (event.Subscription, error) {

	logs, sub, err := _EpochHandler.contract.WatchLogs(opts, "Forked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EpochHandlerForked)
				if err := _EpochHandler.contract.UnpackLog(event, "Forked", log); err != nil {
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

// EpochHandlerRequestAppliedIterator is returned from FilterRequestApplied and is used to iterate over the raw logs and unpacked data for RequestApplied events raised by the EpochHandler contract.
type EpochHandlerRequestAppliedIterator struct {
	Event *EpochHandlerRequestApplied // Event containing the contract specifics and raw log

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
func (it *EpochHandlerRequestAppliedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EpochHandlerRequestApplied)
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
		it.Event = new(EpochHandlerRequestApplied)
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
func (it *EpochHandlerRequestAppliedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EpochHandlerRequestAppliedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EpochHandlerRequestApplied represents a RequestApplied event raised by the EpochHandler contract.
type EpochHandlerRequestApplied struct {
	RequestId     *big.Int
	UserActivated bool
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRequestApplied is a free log retrieval operation binding the contract event 0x6940a01870e576ceb735867e13863646d517ce10e66c0133186a4ebdfe9388c2.
//
// Solidity: e RequestApplied(requestId uint256, userActivated bool)
func (_EpochHandler *EpochHandlerFilterer) FilterRequestApplied(opts *bind.FilterOpts) (*EpochHandlerRequestAppliedIterator, error) {

	logs, sub, err := _EpochHandler.contract.FilterLogs(opts, "RequestApplied")
	if err != nil {
		return nil, err
	}
	return &EpochHandlerRequestAppliedIterator{contract: _EpochHandler.contract, event: "RequestApplied", logs: logs, sub: sub}, nil
}

// WatchRequestApplied is a free log subscription operation binding the contract event 0x6940a01870e576ceb735867e13863646d517ce10e66c0133186a4ebdfe9388c2.
//
// Solidity: e RequestApplied(requestId uint256, userActivated bool)
func (_EpochHandler *EpochHandlerFilterer) WatchRequestApplied(opts *bind.WatchOpts, sink chan<- *EpochHandlerRequestApplied) (event.Subscription, error) {

	logs, sub, err := _EpochHandler.contract.WatchLogs(opts, "RequestApplied")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EpochHandlerRequestApplied)
				if err := _EpochHandler.contract.UnpackLog(event, "RequestApplied", log); err != nil {
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

// EpochHandlerRequestChallengedIterator is returned from FilterRequestChallenged and is used to iterate over the raw logs and unpacked data for RequestChallenged events raised by the EpochHandler contract.
type EpochHandlerRequestChallengedIterator struct {
	Event *EpochHandlerRequestChallenged // Event containing the contract specifics and raw log

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
func (it *EpochHandlerRequestChallengedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EpochHandlerRequestChallenged)
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
		it.Event = new(EpochHandlerRequestChallenged)
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
func (it *EpochHandlerRequestChallengedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EpochHandlerRequestChallengedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EpochHandlerRequestChallenged represents a RequestChallenged event raised by the EpochHandler contract.
type EpochHandlerRequestChallenged struct {
	RequestId     *big.Int
	UserActivated bool
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRequestChallenged is a free log retrieval operation binding the contract event 0xc8135db115644ed4ae193313c4c801235ef740d2a57a8d5e6fe26ab66635698a.
//
// Solidity: e RequestChallenged(requestId uint256, userActivated bool)
func (_EpochHandler *EpochHandlerFilterer) FilterRequestChallenged(opts *bind.FilterOpts) (*EpochHandlerRequestChallengedIterator, error) {

	logs, sub, err := _EpochHandler.contract.FilterLogs(opts, "RequestChallenged")
	if err != nil {
		return nil, err
	}
	return &EpochHandlerRequestChallengedIterator{contract: _EpochHandler.contract, event: "RequestChallenged", logs: logs, sub: sub}, nil
}

// WatchRequestChallenged is a free log subscription operation binding the contract event 0xc8135db115644ed4ae193313c4c801235ef740d2a57a8d5e6fe26ab66635698a.
//
// Solidity: e RequestChallenged(requestId uint256, userActivated bool)
func (_EpochHandler *EpochHandlerFilterer) WatchRequestChallenged(opts *bind.WatchOpts, sink chan<- *EpochHandlerRequestChallenged) (event.Subscription, error) {

	logs, sub, err := _EpochHandler.contract.WatchLogs(opts, "RequestChallenged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EpochHandlerRequestChallenged)
				if err := _EpochHandler.contract.UnpackLog(event, "RequestChallenged", log); err != nil {
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

// EpochHandlerRequestCreatedIterator is returned from FilterRequestCreated and is used to iterate over the raw logs and unpacked data for RequestCreated events raised by the EpochHandler contract.
type EpochHandlerRequestCreatedIterator struct {
	Event *EpochHandlerRequestCreated // Event containing the contract specifics and raw log

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
func (it *EpochHandlerRequestCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EpochHandlerRequestCreated)
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
		it.Event = new(EpochHandlerRequestCreated)
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
func (it *EpochHandlerRequestCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EpochHandlerRequestCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EpochHandlerRequestCreated represents a RequestCreated event raised by the EpochHandler contract.
type EpochHandlerRequestCreated struct {
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
func (_EpochHandler *EpochHandlerFilterer) FilterRequestCreated(opts *bind.FilterOpts) (*EpochHandlerRequestCreatedIterator, error) {

	logs, sub, err := _EpochHandler.contract.FilterLogs(opts, "RequestCreated")
	if err != nil {
		return nil, err
	}
	return &EpochHandlerRequestCreatedIterator{contract: _EpochHandler.contract, event: "RequestCreated", logs: logs, sub: sub}, nil
}

// WatchRequestCreated is a free log subscription operation binding the contract event 0x9d57b50c5371c1c3fc64a8947cec60dbae09432e1e5d9ef048317ad7240353e3.
//
// Solidity: e RequestCreated(requestId uint256, requestor address, to address, weiAmount uint256, trieKey bytes32, trieValue bytes32, isTransfer bool, isExit bool, userActivated bool)
func (_EpochHandler *EpochHandlerFilterer) WatchRequestCreated(opts *bind.WatchOpts, sink chan<- *EpochHandlerRequestCreated) (event.Subscription, error) {

	logs, sub, err := _EpochHandler.contract.WatchLogs(opts, "RequestCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EpochHandlerRequestCreated)
				if err := _EpochHandler.contract.UnpackLog(event, "RequestCreated", log); err != nil {
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

// EpochHandlerRequestFinalizedIterator is returned from FilterRequestFinalized and is used to iterate over the raw logs and unpacked data for RequestFinalized events raised by the EpochHandler contract.
type EpochHandlerRequestFinalizedIterator struct {
	Event *EpochHandlerRequestFinalized // Event containing the contract specifics and raw log

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
func (it *EpochHandlerRequestFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EpochHandlerRequestFinalized)
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
		it.Event = new(EpochHandlerRequestFinalized)
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
func (it *EpochHandlerRequestFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EpochHandlerRequestFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EpochHandlerRequestFinalized represents a RequestFinalized event raised by the EpochHandler contract.
type EpochHandlerRequestFinalized struct {
	RequestId     *big.Int
	UserActivated bool
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRequestFinalized is a free log retrieval operation binding the contract event 0x134017cf3262b18f892ee95dde3b0aec9a80cc70a7c96f09c64bd237aceb0473.
//
// Solidity: e RequestFinalized(requestId uint256, userActivated bool)
func (_EpochHandler *EpochHandlerFilterer) FilterRequestFinalized(opts *bind.FilterOpts) (*EpochHandlerRequestFinalizedIterator, error) {

	logs, sub, err := _EpochHandler.contract.FilterLogs(opts, "RequestFinalized")
	if err != nil {
		return nil, err
	}
	return &EpochHandlerRequestFinalizedIterator{contract: _EpochHandler.contract, event: "RequestFinalized", logs: logs, sub: sub}, nil
}

// WatchRequestFinalized is a free log subscription operation binding the contract event 0x134017cf3262b18f892ee95dde3b0aec9a80cc70a7c96f09c64bd237aceb0473.
//
// Solidity: e RequestFinalized(requestId uint256, userActivated bool)
func (_EpochHandler *EpochHandlerFilterer) WatchRequestFinalized(opts *bind.WatchOpts, sink chan<- *EpochHandlerRequestFinalized) (event.Subscription, error) {

	logs, sub, err := _EpochHandler.contract.WatchLogs(opts, "RequestFinalized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EpochHandlerRequestFinalized)
				if err := _EpochHandler.contract.UnpackLog(event, "RequestFinalized", log); err != nil {
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

// EpochHandlerSessionTimeoutIterator is returned from FilterSessionTimeout and is used to iterate over the raw logs and unpacked data for SessionTimeout events raised by the EpochHandler contract.
type EpochHandlerSessionTimeoutIterator struct {
	Event *EpochHandlerSessionTimeout // Event containing the contract specifics and raw log

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
func (it *EpochHandlerSessionTimeoutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EpochHandlerSessionTimeout)
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
		it.Event = new(EpochHandlerSessionTimeout)
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
func (it *EpochHandlerSessionTimeoutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EpochHandlerSessionTimeoutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EpochHandlerSessionTimeout represents a SessionTimeout event raised by the EpochHandler contract.
type EpochHandlerSessionTimeout struct {
	UserActivated bool
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSessionTimeout is a free log retrieval operation binding the contract event 0x2122ec719581cd177f225f59a1ee005831211196831f8f1ccffa817d2e7bd108.
//
// Solidity: e SessionTimeout(userActivated bool)
func (_EpochHandler *EpochHandlerFilterer) FilterSessionTimeout(opts *bind.FilterOpts) (*EpochHandlerSessionTimeoutIterator, error) {

	logs, sub, err := _EpochHandler.contract.FilterLogs(opts, "SessionTimeout")
	if err != nil {
		return nil, err
	}
	return &EpochHandlerSessionTimeoutIterator{contract: _EpochHandler.contract, event: "SessionTimeout", logs: logs, sub: sub}, nil
}

// WatchSessionTimeout is a free log subscription operation binding the contract event 0x2122ec719581cd177f225f59a1ee005831211196831f8f1ccffa817d2e7bd108.
//
// Solidity: e SessionTimeout(userActivated bool)
func (_EpochHandler *EpochHandlerFilterer) WatchSessionTimeout(opts *bind.WatchOpts, sink chan<- *EpochHandlerSessionTimeout) (event.Subscription, error) {

	logs, sub, err := _EpochHandler.contract.WatchLogs(opts, "SessionTimeout")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EpochHandlerSessionTimeout)
				if err := _EpochHandler.contract.UnpackLog(event, "SessionTimeout", log); err != nil {
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

// RootChainEventABI is the input ABI used to generate the binding from.
const RootChainEventABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"userActivated\",\"type\":\"bool\"}],\"name\":\"SessionTimeout\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newFork\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"epochNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"forkedBlockNumber\",\"type\":\"uint256\"}],\"name\":\"Forked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"forkNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"epochNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"startBlockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"endBlockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"requestStart\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"requestEnd\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"epochIsEmpty\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"isRequest\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"userActivated\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"rebase\",\"type\":\"bool\"}],\"name\":\"EpochPrepared\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"forkNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"epochNumber\",\"type\":\"uint256\"}],\"name\":\"EpochFilling\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"forkNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"epochNumber\",\"type\":\"uint256\"}],\"name\":\"EpochFilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"forkNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"epochNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"startBlockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"endBlockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"requestStart\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"requestEnd\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"epochIsEmpty\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"isRequest\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"userActivated\",\"type\":\"bool\"}],\"name\":\"EpochRebased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"fork\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"epochNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"isRequest\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"userActivated\",\"type\":\"bool\"}],\"name\":\"BlockSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"requestor\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"weiAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"trieKey\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"trieValue\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"isTransfer\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"isExit\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"userActivated\",\"type\":\"bool\"}],\"name\":\"RequestCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"requestor\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"trieKey\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"trieValue\",\"type\":\"bytes32\"}],\"name\":\"ERUCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"forkNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"BlockFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"forkNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"epochNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"startBlockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"endBlockNumber\",\"type\":\"uint256\"}],\"name\":\"EpochFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"userActivated\",\"type\":\"bool\"}],\"name\":\"RequestFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"userActivated\",\"type\":\"bool\"}],\"name\":\"RequestApplied\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"userActivated\",\"type\":\"bool\"}],\"name\":\"RequestChallenged\",\"type\":\"event\"}]"

// RootChainEventBin is the compiled bytecode used for deploying new contracts.
const RootChainEventBin = `0x6080604052348015600f57600080fd5b50603580601d6000396000f3006080604052600080fd00a165627a7a72305820c785dc71652ad777820a42fb69b7a903356787b890d3a0076c6ea9e190b685530029`

// DeployRootChainEvent deploys a new Ethereum contract, binding an instance of RootChainEvent to it.
func DeployRootChainEvent(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RootChainEvent, error) {
	parsed, err := abi.JSON(strings.NewReader(RootChainEventABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RootChainEventBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RootChainEvent{RootChainEventCaller: RootChainEventCaller{contract: contract}, RootChainEventTransactor: RootChainEventTransactor{contract: contract}, RootChainEventFilterer: RootChainEventFilterer{contract: contract}}, nil
}

// RootChainEvent is an auto generated Go binding around an Ethereum contract.
type RootChainEvent struct {
	RootChainEventCaller     // Read-only binding to the contract
	RootChainEventTransactor // Write-only binding to the contract
	RootChainEventFilterer   // Log filterer for contract events
}

// RootChainEventCaller is an auto generated read-only Go binding around an Ethereum contract.
type RootChainEventCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RootChainEventTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RootChainEventTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RootChainEventFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RootChainEventFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RootChainEventSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RootChainEventSession struct {
	Contract     *RootChainEvent   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RootChainEventCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RootChainEventCallerSession struct {
	Contract *RootChainEventCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// RootChainEventTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RootChainEventTransactorSession struct {
	Contract     *RootChainEventTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// RootChainEventRaw is an auto generated low-level Go binding around an Ethereum contract.
type RootChainEventRaw struct {
	Contract *RootChainEvent // Generic contract binding to access the raw methods on
}

// RootChainEventCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RootChainEventCallerRaw struct {
	Contract *RootChainEventCaller // Generic read-only contract binding to access the raw methods on
}

// RootChainEventTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RootChainEventTransactorRaw struct {
	Contract *RootChainEventTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRootChainEvent creates a new instance of RootChainEvent, bound to a specific deployed contract.
func NewRootChainEvent(address common.Address, backend bind.ContractBackend) (*RootChainEvent, error) {
	contract, err := bindRootChainEvent(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RootChainEvent{RootChainEventCaller: RootChainEventCaller{contract: contract}, RootChainEventTransactor: RootChainEventTransactor{contract: contract}, RootChainEventFilterer: RootChainEventFilterer{contract: contract}}, nil
}

// NewRootChainEventCaller creates a new read-only instance of RootChainEvent, bound to a specific deployed contract.
func NewRootChainEventCaller(address common.Address, caller bind.ContractCaller) (*RootChainEventCaller, error) {
	contract, err := bindRootChainEvent(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RootChainEventCaller{contract: contract}, nil
}

// NewRootChainEventTransactor creates a new write-only instance of RootChainEvent, bound to a specific deployed contract.
func NewRootChainEventTransactor(address common.Address, transactor bind.ContractTransactor) (*RootChainEventTransactor, error) {
	contract, err := bindRootChainEvent(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RootChainEventTransactor{contract: contract}, nil
}

// NewRootChainEventFilterer creates a new log filterer instance of RootChainEvent, bound to a specific deployed contract.
func NewRootChainEventFilterer(address common.Address, filterer bind.ContractFilterer) (*RootChainEventFilterer, error) {
	contract, err := bindRootChainEvent(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RootChainEventFilterer{contract: contract}, nil
}

// bindRootChainEvent binds a generic wrapper to an already deployed contract.
func bindRootChainEvent(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RootChainEventABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RootChainEvent *RootChainEventRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RootChainEvent.Contract.RootChainEventCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RootChainEvent *RootChainEventRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RootChainEvent.Contract.RootChainEventTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RootChainEvent *RootChainEventRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RootChainEvent.Contract.RootChainEventTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RootChainEvent *RootChainEventCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RootChainEvent.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RootChainEvent *RootChainEventTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RootChainEvent.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RootChainEvent *RootChainEventTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RootChainEvent.Contract.contract.Transact(opts, method, params...)
}

// RootChainEventBlockFinalizedIterator is returned from FilterBlockFinalized and is used to iterate over the raw logs and unpacked data for BlockFinalized events raised by the RootChainEvent contract.
type RootChainEventBlockFinalizedIterator struct {
	Event *RootChainEventBlockFinalized // Event containing the contract specifics and raw log

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
func (it *RootChainEventBlockFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RootChainEventBlockFinalized)
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
		it.Event = new(RootChainEventBlockFinalized)
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
func (it *RootChainEventBlockFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RootChainEventBlockFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RootChainEventBlockFinalized represents a BlockFinalized event raised by the RootChainEvent contract.
type RootChainEventBlockFinalized struct {
	ForkNumber  *big.Int
	BlockNumber *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBlockFinalized is a free log retrieval operation binding the contract event 0xfb96205e4b3633fd57aa805b26b51ecf528714a10241a4af015929dce86768d9.
//
// Solidity: e BlockFinalized(forkNumber uint256, blockNumber uint256)
func (_RootChainEvent *RootChainEventFilterer) FilterBlockFinalized(opts *bind.FilterOpts) (*RootChainEventBlockFinalizedIterator, error) {

	logs, sub, err := _RootChainEvent.contract.FilterLogs(opts, "BlockFinalized")
	if err != nil {
		return nil, err
	}
	return &RootChainEventBlockFinalizedIterator{contract: _RootChainEvent.contract, event: "BlockFinalized", logs: logs, sub: sub}, nil
}

// WatchBlockFinalized is a free log subscription operation binding the contract event 0xfb96205e4b3633fd57aa805b26b51ecf528714a10241a4af015929dce86768d9.
//
// Solidity: e BlockFinalized(forkNumber uint256, blockNumber uint256)
func (_RootChainEvent *RootChainEventFilterer) WatchBlockFinalized(opts *bind.WatchOpts, sink chan<- *RootChainEventBlockFinalized) (event.Subscription, error) {

	logs, sub, err := _RootChainEvent.contract.WatchLogs(opts, "BlockFinalized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RootChainEventBlockFinalized)
				if err := _RootChainEvent.contract.UnpackLog(event, "BlockFinalized", log); err != nil {
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

// RootChainEventBlockSubmittedIterator is returned from FilterBlockSubmitted and is used to iterate over the raw logs and unpacked data for BlockSubmitted events raised by the RootChainEvent contract.
type RootChainEventBlockSubmittedIterator struct {
	Event *RootChainEventBlockSubmitted // Event containing the contract specifics and raw log

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
func (it *RootChainEventBlockSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RootChainEventBlockSubmitted)
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
		it.Event = new(RootChainEventBlockSubmitted)
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
func (it *RootChainEventBlockSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RootChainEventBlockSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RootChainEventBlockSubmitted represents a BlockSubmitted event raised by the RootChainEvent contract.
type RootChainEventBlockSubmitted struct {
	Fork          *big.Int
	EpochNumber   *big.Int
	BlockNumber   *big.Int
	IsRequest     bool
	UserActivated bool
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterBlockSubmitted is a free log retrieval operation binding the contract event 0x3d4a04291c66b06f39a4ecb817875b12b5485a05ec563133a56a905305c48e55.
//
// Solidity: e BlockSubmitted(fork uint256, epochNumber uint256, blockNumber uint256, isRequest bool, userActivated bool)
func (_RootChainEvent *RootChainEventFilterer) FilterBlockSubmitted(opts *bind.FilterOpts) (*RootChainEventBlockSubmittedIterator, error) {

	logs, sub, err := _RootChainEvent.contract.FilterLogs(opts, "BlockSubmitted")
	if err != nil {
		return nil, err
	}
	return &RootChainEventBlockSubmittedIterator{contract: _RootChainEvent.contract, event: "BlockSubmitted", logs: logs, sub: sub}, nil
}

// WatchBlockSubmitted is a free log subscription operation binding the contract event 0x3d4a04291c66b06f39a4ecb817875b12b5485a05ec563133a56a905305c48e55.
//
// Solidity: e BlockSubmitted(fork uint256, epochNumber uint256, blockNumber uint256, isRequest bool, userActivated bool)
func (_RootChainEvent *RootChainEventFilterer) WatchBlockSubmitted(opts *bind.WatchOpts, sink chan<- *RootChainEventBlockSubmitted) (event.Subscription, error) {

	logs, sub, err := _RootChainEvent.contract.WatchLogs(opts, "BlockSubmitted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RootChainEventBlockSubmitted)
				if err := _RootChainEvent.contract.UnpackLog(event, "BlockSubmitted", log); err != nil {
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

// RootChainEventERUCreatedIterator is returned from FilterERUCreated and is used to iterate over the raw logs and unpacked data for ERUCreated events raised by the RootChainEvent contract.
type RootChainEventERUCreatedIterator struct {
	Event *RootChainEventERUCreated // Event containing the contract specifics and raw log

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
func (it *RootChainEventERUCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RootChainEventERUCreated)
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
		it.Event = new(RootChainEventERUCreated)
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
func (it *RootChainEventERUCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RootChainEventERUCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RootChainEventERUCreated represents a ERUCreated event raised by the RootChainEvent contract.
type RootChainEventERUCreated struct {
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
func (_RootChainEvent *RootChainEventFilterer) FilterERUCreated(opts *bind.FilterOpts) (*RootChainEventERUCreatedIterator, error) {

	logs, sub, err := _RootChainEvent.contract.FilterLogs(opts, "ERUCreated")
	if err != nil {
		return nil, err
	}
	return &RootChainEventERUCreatedIterator{contract: _RootChainEvent.contract, event: "ERUCreated", logs: logs, sub: sub}, nil
}

// WatchERUCreated is a free log subscription operation binding the contract event 0xd89c6857ed5778107e858511cdc309642a48c9d1717e813a25156f535acc8d36.
//
// Solidity: e ERUCreated(requestId uint256, requestor address, to address, trieKey bytes32, trieValue bytes32)
func (_RootChainEvent *RootChainEventFilterer) WatchERUCreated(opts *bind.WatchOpts, sink chan<- *RootChainEventERUCreated) (event.Subscription, error) {

	logs, sub, err := _RootChainEvent.contract.WatchLogs(opts, "ERUCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RootChainEventERUCreated)
				if err := _RootChainEvent.contract.UnpackLog(event, "ERUCreated", log); err != nil {
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

// RootChainEventEpochFilledIterator is returned from FilterEpochFilled and is used to iterate over the raw logs and unpacked data for EpochFilled events raised by the RootChainEvent contract.
type RootChainEventEpochFilledIterator struct {
	Event *RootChainEventEpochFilled // Event containing the contract specifics and raw log

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
func (it *RootChainEventEpochFilledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RootChainEventEpochFilled)
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
		it.Event = new(RootChainEventEpochFilled)
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
func (it *RootChainEventEpochFilledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RootChainEventEpochFilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RootChainEventEpochFilled represents a EpochFilled event raised by the RootChainEvent contract.
type RootChainEventEpochFilled struct {
	ForkNumber  *big.Int
	EpochNumber *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterEpochFilled is a free log retrieval operation binding the contract event 0x2fdeb407bf5c2b621f04b5c784822dae806c45b49a68aba413cc270128c96816.
//
// Solidity: e EpochFilled(forkNumber uint256, epochNumber uint256)
func (_RootChainEvent *RootChainEventFilterer) FilterEpochFilled(opts *bind.FilterOpts) (*RootChainEventEpochFilledIterator, error) {

	logs, sub, err := _RootChainEvent.contract.FilterLogs(opts, "EpochFilled")
	if err != nil {
		return nil, err
	}
	return &RootChainEventEpochFilledIterator{contract: _RootChainEvent.contract, event: "EpochFilled", logs: logs, sub: sub}, nil
}

// WatchEpochFilled is a free log subscription operation binding the contract event 0x2fdeb407bf5c2b621f04b5c784822dae806c45b49a68aba413cc270128c96816.
//
// Solidity: e EpochFilled(forkNumber uint256, epochNumber uint256)
func (_RootChainEvent *RootChainEventFilterer) WatchEpochFilled(opts *bind.WatchOpts, sink chan<- *RootChainEventEpochFilled) (event.Subscription, error) {

	logs, sub, err := _RootChainEvent.contract.WatchLogs(opts, "EpochFilled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RootChainEventEpochFilled)
				if err := _RootChainEvent.contract.UnpackLog(event, "EpochFilled", log); err != nil {
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

// RootChainEventEpochFillingIterator is returned from FilterEpochFilling and is used to iterate over the raw logs and unpacked data for EpochFilling events raised by the RootChainEvent contract.
type RootChainEventEpochFillingIterator struct {
	Event *RootChainEventEpochFilling // Event containing the contract specifics and raw log

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
func (it *RootChainEventEpochFillingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RootChainEventEpochFilling)
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
		it.Event = new(RootChainEventEpochFilling)
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
func (it *RootChainEventEpochFillingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RootChainEventEpochFillingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RootChainEventEpochFilling represents a EpochFilling event raised by the RootChainEvent contract.
type RootChainEventEpochFilling struct {
	ForkNumber  *big.Int
	EpochNumber *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterEpochFilling is a free log retrieval operation binding the contract event 0x27b09f0953d27bbff306fe25b2987ac5a813248ac30cb2bbd5daf95e7b0e6dc0.
//
// Solidity: e EpochFilling(forkNumber uint256, epochNumber uint256)
func (_RootChainEvent *RootChainEventFilterer) FilterEpochFilling(opts *bind.FilterOpts) (*RootChainEventEpochFillingIterator, error) {

	logs, sub, err := _RootChainEvent.contract.FilterLogs(opts, "EpochFilling")
	if err != nil {
		return nil, err
	}
	return &RootChainEventEpochFillingIterator{contract: _RootChainEvent.contract, event: "EpochFilling", logs: logs, sub: sub}, nil
}

// WatchEpochFilling is a free log subscription operation binding the contract event 0x27b09f0953d27bbff306fe25b2987ac5a813248ac30cb2bbd5daf95e7b0e6dc0.
//
// Solidity: e EpochFilling(forkNumber uint256, epochNumber uint256)
func (_RootChainEvent *RootChainEventFilterer) WatchEpochFilling(opts *bind.WatchOpts, sink chan<- *RootChainEventEpochFilling) (event.Subscription, error) {

	logs, sub, err := _RootChainEvent.contract.WatchLogs(opts, "EpochFilling")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RootChainEventEpochFilling)
				if err := _RootChainEvent.contract.UnpackLog(event, "EpochFilling", log); err != nil {
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

// RootChainEventEpochFinalizedIterator is returned from FilterEpochFinalized and is used to iterate over the raw logs and unpacked data for EpochFinalized events raised by the RootChainEvent contract.
type RootChainEventEpochFinalizedIterator struct {
	Event *RootChainEventEpochFinalized // Event containing the contract specifics and raw log

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
func (it *RootChainEventEpochFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RootChainEventEpochFinalized)
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
		it.Event = new(RootChainEventEpochFinalized)
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
func (it *RootChainEventEpochFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RootChainEventEpochFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RootChainEventEpochFinalized represents a EpochFinalized event raised by the RootChainEvent contract.
type RootChainEventEpochFinalized struct {
	ForkNumber       *big.Int
	EpochNumber      *big.Int
	StartBlockNumber *big.Int
	EndBlockNumber   *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterEpochFinalized is a free log retrieval operation binding the contract event 0x70801d4d63b3da6c19ba7349911f45bed5a99ccdfb51b8138c105872529bebd5.
//
// Solidity: e EpochFinalized(forkNumber uint256, epochNumber uint256, startBlockNumber uint256, endBlockNumber uint256)
func (_RootChainEvent *RootChainEventFilterer) FilterEpochFinalized(opts *bind.FilterOpts) (*RootChainEventEpochFinalizedIterator, error) {

	logs, sub, err := _RootChainEvent.contract.FilterLogs(opts, "EpochFinalized")
	if err != nil {
		return nil, err
	}
	return &RootChainEventEpochFinalizedIterator{contract: _RootChainEvent.contract, event: "EpochFinalized", logs: logs, sub: sub}, nil
}

// WatchEpochFinalized is a free log subscription operation binding the contract event 0x70801d4d63b3da6c19ba7349911f45bed5a99ccdfb51b8138c105872529bebd5.
//
// Solidity: e EpochFinalized(forkNumber uint256, epochNumber uint256, startBlockNumber uint256, endBlockNumber uint256)
func (_RootChainEvent *RootChainEventFilterer) WatchEpochFinalized(opts *bind.WatchOpts, sink chan<- *RootChainEventEpochFinalized) (event.Subscription, error) {

	logs, sub, err := _RootChainEvent.contract.WatchLogs(opts, "EpochFinalized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RootChainEventEpochFinalized)
				if err := _RootChainEvent.contract.UnpackLog(event, "EpochFinalized", log); err != nil {
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

// RootChainEventEpochPreparedIterator is returned from FilterEpochPrepared and is used to iterate over the raw logs and unpacked data for EpochPrepared events raised by the RootChainEvent contract.
type RootChainEventEpochPreparedIterator struct {
	Event *RootChainEventEpochPrepared // Event containing the contract specifics and raw log

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
func (it *RootChainEventEpochPreparedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RootChainEventEpochPrepared)
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
		it.Event = new(RootChainEventEpochPrepared)
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
func (it *RootChainEventEpochPreparedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RootChainEventEpochPreparedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RootChainEventEpochPrepared represents a EpochPrepared event raised by the RootChainEvent contract.
type RootChainEventEpochPrepared struct {
	ForkNumber       *big.Int
	EpochNumber      *big.Int
	StartBlockNumber *big.Int
	EndBlockNumber   *big.Int
	RequestStart     *big.Int
	RequestEnd       *big.Int
	EpochIsEmpty     bool
	IsRequest        bool
	UserActivated    bool
	Rebase           bool
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterEpochPrepared is a free log retrieval operation binding the contract event 0x1a69c0760aa329b76f72579129869013ebd3d41594db019c0e997b939fcb32e3.
//
// Solidity: e EpochPrepared(forkNumber uint256, epochNumber uint256, startBlockNumber uint256, endBlockNumber uint256, requestStart uint256, requestEnd uint256, epochIsEmpty bool, isRequest bool, userActivated bool, rebase bool)
func (_RootChainEvent *RootChainEventFilterer) FilterEpochPrepared(opts *bind.FilterOpts) (*RootChainEventEpochPreparedIterator, error) {

	logs, sub, err := _RootChainEvent.contract.FilterLogs(opts, "EpochPrepared")
	if err != nil {
		return nil, err
	}
	return &RootChainEventEpochPreparedIterator{contract: _RootChainEvent.contract, event: "EpochPrepared", logs: logs, sub: sub}, nil
}

// WatchEpochPrepared is a free log subscription operation binding the contract event 0x1a69c0760aa329b76f72579129869013ebd3d41594db019c0e997b939fcb32e3.
//
// Solidity: e EpochPrepared(forkNumber uint256, epochNumber uint256, startBlockNumber uint256, endBlockNumber uint256, requestStart uint256, requestEnd uint256, epochIsEmpty bool, isRequest bool, userActivated bool, rebase bool)
func (_RootChainEvent *RootChainEventFilterer) WatchEpochPrepared(opts *bind.WatchOpts, sink chan<- *RootChainEventEpochPrepared) (event.Subscription, error) {

	logs, sub, err := _RootChainEvent.contract.WatchLogs(opts, "EpochPrepared")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RootChainEventEpochPrepared)
				if err := _RootChainEvent.contract.UnpackLog(event, "EpochPrepared", log); err != nil {
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

// RootChainEventEpochRebasedIterator is returned from FilterEpochRebased and is used to iterate over the raw logs and unpacked data for EpochRebased events raised by the RootChainEvent contract.
type RootChainEventEpochRebasedIterator struct {
	Event *RootChainEventEpochRebased // Event containing the contract specifics and raw log

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
func (it *RootChainEventEpochRebasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RootChainEventEpochRebased)
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
		it.Event = new(RootChainEventEpochRebased)
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
func (it *RootChainEventEpochRebasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RootChainEventEpochRebasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RootChainEventEpochRebased represents a EpochRebased event raised by the RootChainEvent contract.
type RootChainEventEpochRebased struct {
	ForkNumber       *big.Int
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

// FilterEpochRebased is a free log retrieval operation binding the contract event 0x030c1c69405c93021f28f57557240dee939a320b826a1fd0d39bf6e629ecab47.
//
// Solidity: e EpochRebased(forkNumber uint256, epochNumber uint256, startBlockNumber uint256, endBlockNumber uint256, requestStart uint256, requestEnd uint256, epochIsEmpty bool, isRequest bool, userActivated bool)
func (_RootChainEvent *RootChainEventFilterer) FilterEpochRebased(opts *bind.FilterOpts) (*RootChainEventEpochRebasedIterator, error) {

	logs, sub, err := _RootChainEvent.contract.FilterLogs(opts, "EpochRebased")
	if err != nil {
		return nil, err
	}
	return &RootChainEventEpochRebasedIterator{contract: _RootChainEvent.contract, event: "EpochRebased", logs: logs, sub: sub}, nil
}

// WatchEpochRebased is a free log subscription operation binding the contract event 0x030c1c69405c93021f28f57557240dee939a320b826a1fd0d39bf6e629ecab47.
//
// Solidity: e EpochRebased(forkNumber uint256, epochNumber uint256, startBlockNumber uint256, endBlockNumber uint256, requestStart uint256, requestEnd uint256, epochIsEmpty bool, isRequest bool, userActivated bool)
func (_RootChainEvent *RootChainEventFilterer) WatchEpochRebased(opts *bind.WatchOpts, sink chan<- *RootChainEventEpochRebased) (event.Subscription, error) {

	logs, sub, err := _RootChainEvent.contract.WatchLogs(opts, "EpochRebased")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RootChainEventEpochRebased)
				if err := _RootChainEvent.contract.UnpackLog(event, "EpochRebased", log); err != nil {
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

// RootChainEventForkedIterator is returned from FilterForked and is used to iterate over the raw logs and unpacked data for Forked events raised by the RootChainEvent contract.
type RootChainEventForkedIterator struct {
	Event *RootChainEventForked // Event containing the contract specifics and raw log

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
func (it *RootChainEventForkedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RootChainEventForked)
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
		it.Event = new(RootChainEventForked)
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
func (it *RootChainEventForkedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RootChainEventForkedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RootChainEventForked represents a Forked event raised by the RootChainEvent contract.
type RootChainEventForked struct {
	NewFork           *big.Int
	EpochNumber       *big.Int
	ForkedBlockNumber *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterForked is a free log retrieval operation binding the contract event 0x0647d42ab02f6e0ae76959757dcb6aa6feac1d4ba6f077f1223fb4b1b429f06c.
//
// Solidity: e Forked(newFork uint256, epochNumber uint256, forkedBlockNumber uint256)
func (_RootChainEvent *RootChainEventFilterer) FilterForked(opts *bind.FilterOpts) (*RootChainEventForkedIterator, error) {

	logs, sub, err := _RootChainEvent.contract.FilterLogs(opts, "Forked")
	if err != nil {
		return nil, err
	}
	return &RootChainEventForkedIterator{contract: _RootChainEvent.contract, event: "Forked", logs: logs, sub: sub}, nil
}

// WatchForked is a free log subscription operation binding the contract event 0x0647d42ab02f6e0ae76959757dcb6aa6feac1d4ba6f077f1223fb4b1b429f06c.
//
// Solidity: e Forked(newFork uint256, epochNumber uint256, forkedBlockNumber uint256)
func (_RootChainEvent *RootChainEventFilterer) WatchForked(opts *bind.WatchOpts, sink chan<- *RootChainEventForked) (event.Subscription, error) {

	logs, sub, err := _RootChainEvent.contract.WatchLogs(opts, "Forked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RootChainEventForked)
				if err := _RootChainEvent.contract.UnpackLog(event, "Forked", log); err != nil {
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

// RootChainEventRequestAppliedIterator is returned from FilterRequestApplied and is used to iterate over the raw logs and unpacked data for RequestApplied events raised by the RootChainEvent contract.
type RootChainEventRequestAppliedIterator struct {
	Event *RootChainEventRequestApplied // Event containing the contract specifics and raw log

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
func (it *RootChainEventRequestAppliedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RootChainEventRequestApplied)
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
		it.Event = new(RootChainEventRequestApplied)
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
func (it *RootChainEventRequestAppliedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RootChainEventRequestAppliedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RootChainEventRequestApplied represents a RequestApplied event raised by the RootChainEvent contract.
type RootChainEventRequestApplied struct {
	RequestId     *big.Int
	UserActivated bool
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRequestApplied is a free log retrieval operation binding the contract event 0x6940a01870e576ceb735867e13863646d517ce10e66c0133186a4ebdfe9388c2.
//
// Solidity: e RequestApplied(requestId uint256, userActivated bool)
func (_RootChainEvent *RootChainEventFilterer) FilterRequestApplied(opts *bind.FilterOpts) (*RootChainEventRequestAppliedIterator, error) {

	logs, sub, err := _RootChainEvent.contract.FilterLogs(opts, "RequestApplied")
	if err != nil {
		return nil, err
	}
	return &RootChainEventRequestAppliedIterator{contract: _RootChainEvent.contract, event: "RequestApplied", logs: logs, sub: sub}, nil
}

// WatchRequestApplied is a free log subscription operation binding the contract event 0x6940a01870e576ceb735867e13863646d517ce10e66c0133186a4ebdfe9388c2.
//
// Solidity: e RequestApplied(requestId uint256, userActivated bool)
func (_RootChainEvent *RootChainEventFilterer) WatchRequestApplied(opts *bind.WatchOpts, sink chan<- *RootChainEventRequestApplied) (event.Subscription, error) {

	logs, sub, err := _RootChainEvent.contract.WatchLogs(opts, "RequestApplied")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RootChainEventRequestApplied)
				if err := _RootChainEvent.contract.UnpackLog(event, "RequestApplied", log); err != nil {
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

// RootChainEventRequestChallengedIterator is returned from FilterRequestChallenged and is used to iterate over the raw logs and unpacked data for RequestChallenged events raised by the RootChainEvent contract.
type RootChainEventRequestChallengedIterator struct {
	Event *RootChainEventRequestChallenged // Event containing the contract specifics and raw log

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
func (it *RootChainEventRequestChallengedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RootChainEventRequestChallenged)
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
		it.Event = new(RootChainEventRequestChallenged)
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
func (it *RootChainEventRequestChallengedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RootChainEventRequestChallengedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RootChainEventRequestChallenged represents a RequestChallenged event raised by the RootChainEvent contract.
type RootChainEventRequestChallenged struct {
	RequestId     *big.Int
	UserActivated bool
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRequestChallenged is a free log retrieval operation binding the contract event 0xc8135db115644ed4ae193313c4c801235ef740d2a57a8d5e6fe26ab66635698a.
//
// Solidity: e RequestChallenged(requestId uint256, userActivated bool)
func (_RootChainEvent *RootChainEventFilterer) FilterRequestChallenged(opts *bind.FilterOpts) (*RootChainEventRequestChallengedIterator, error) {

	logs, sub, err := _RootChainEvent.contract.FilterLogs(opts, "RequestChallenged")
	if err != nil {
		return nil, err
	}
	return &RootChainEventRequestChallengedIterator{contract: _RootChainEvent.contract, event: "RequestChallenged", logs: logs, sub: sub}, nil
}

// WatchRequestChallenged is a free log subscription operation binding the contract event 0xc8135db115644ed4ae193313c4c801235ef740d2a57a8d5e6fe26ab66635698a.
//
// Solidity: e RequestChallenged(requestId uint256, userActivated bool)
func (_RootChainEvent *RootChainEventFilterer) WatchRequestChallenged(opts *bind.WatchOpts, sink chan<- *RootChainEventRequestChallenged) (event.Subscription, error) {

	logs, sub, err := _RootChainEvent.contract.WatchLogs(opts, "RequestChallenged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RootChainEventRequestChallenged)
				if err := _RootChainEvent.contract.UnpackLog(event, "RequestChallenged", log); err != nil {
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

// RootChainEventRequestCreatedIterator is returned from FilterRequestCreated and is used to iterate over the raw logs and unpacked data for RequestCreated events raised by the RootChainEvent contract.
type RootChainEventRequestCreatedIterator struct {
	Event *RootChainEventRequestCreated // Event containing the contract specifics and raw log

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
func (it *RootChainEventRequestCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RootChainEventRequestCreated)
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
		it.Event = new(RootChainEventRequestCreated)
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
func (it *RootChainEventRequestCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RootChainEventRequestCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RootChainEventRequestCreated represents a RequestCreated event raised by the RootChainEvent contract.
type RootChainEventRequestCreated struct {
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
func (_RootChainEvent *RootChainEventFilterer) FilterRequestCreated(opts *bind.FilterOpts) (*RootChainEventRequestCreatedIterator, error) {

	logs, sub, err := _RootChainEvent.contract.FilterLogs(opts, "RequestCreated")
	if err != nil {
		return nil, err
	}
	return &RootChainEventRequestCreatedIterator{contract: _RootChainEvent.contract, event: "RequestCreated", logs: logs, sub: sub}, nil
}

// WatchRequestCreated is a free log subscription operation binding the contract event 0x9d57b50c5371c1c3fc64a8947cec60dbae09432e1e5d9ef048317ad7240353e3.
//
// Solidity: e RequestCreated(requestId uint256, requestor address, to address, weiAmount uint256, trieKey bytes32, trieValue bytes32, isTransfer bool, isExit bool, userActivated bool)
func (_RootChainEvent *RootChainEventFilterer) WatchRequestCreated(opts *bind.WatchOpts, sink chan<- *RootChainEventRequestCreated) (event.Subscription, error) {

	logs, sub, err := _RootChainEvent.contract.WatchLogs(opts, "RequestCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RootChainEventRequestCreated)
				if err := _RootChainEvent.contract.UnpackLog(event, "RequestCreated", log); err != nil {
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

// RootChainEventRequestFinalizedIterator is returned from FilterRequestFinalized and is used to iterate over the raw logs and unpacked data for RequestFinalized events raised by the RootChainEvent contract.
type RootChainEventRequestFinalizedIterator struct {
	Event *RootChainEventRequestFinalized // Event containing the contract specifics and raw log

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
func (it *RootChainEventRequestFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RootChainEventRequestFinalized)
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
		it.Event = new(RootChainEventRequestFinalized)
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
func (it *RootChainEventRequestFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RootChainEventRequestFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RootChainEventRequestFinalized represents a RequestFinalized event raised by the RootChainEvent contract.
type RootChainEventRequestFinalized struct {
	RequestId     *big.Int
	UserActivated bool
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRequestFinalized is a free log retrieval operation binding the contract event 0x134017cf3262b18f892ee95dde3b0aec9a80cc70a7c96f09c64bd237aceb0473.
//
// Solidity: e RequestFinalized(requestId uint256, userActivated bool)
func (_RootChainEvent *RootChainEventFilterer) FilterRequestFinalized(opts *bind.FilterOpts) (*RootChainEventRequestFinalizedIterator, error) {

	logs, sub, err := _RootChainEvent.contract.FilterLogs(opts, "RequestFinalized")
	if err != nil {
		return nil, err
	}
	return &RootChainEventRequestFinalizedIterator{contract: _RootChainEvent.contract, event: "RequestFinalized", logs: logs, sub: sub}, nil
}

// WatchRequestFinalized is a free log subscription operation binding the contract event 0x134017cf3262b18f892ee95dde3b0aec9a80cc70a7c96f09c64bd237aceb0473.
//
// Solidity: e RequestFinalized(requestId uint256, userActivated bool)
func (_RootChainEvent *RootChainEventFilterer) WatchRequestFinalized(opts *bind.WatchOpts, sink chan<- *RootChainEventRequestFinalized) (event.Subscription, error) {

	logs, sub, err := _RootChainEvent.contract.WatchLogs(opts, "RequestFinalized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RootChainEventRequestFinalized)
				if err := _RootChainEvent.contract.UnpackLog(event, "RequestFinalized", log); err != nil {
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

// RootChainEventSessionTimeoutIterator is returned from FilterSessionTimeout and is used to iterate over the raw logs and unpacked data for SessionTimeout events raised by the RootChainEvent contract.
type RootChainEventSessionTimeoutIterator struct {
	Event *RootChainEventSessionTimeout // Event containing the contract specifics and raw log

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
func (it *RootChainEventSessionTimeoutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RootChainEventSessionTimeout)
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
		it.Event = new(RootChainEventSessionTimeout)
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
func (it *RootChainEventSessionTimeoutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RootChainEventSessionTimeoutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RootChainEventSessionTimeout represents a SessionTimeout event raised by the RootChainEvent contract.
type RootChainEventSessionTimeout struct {
	UserActivated bool
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSessionTimeout is a free log retrieval operation binding the contract event 0x2122ec719581cd177f225f59a1ee005831211196831f8f1ccffa817d2e7bd108.
//
// Solidity: e SessionTimeout(userActivated bool)
func (_RootChainEvent *RootChainEventFilterer) FilterSessionTimeout(opts *bind.FilterOpts) (*RootChainEventSessionTimeoutIterator, error) {

	logs, sub, err := _RootChainEvent.contract.FilterLogs(opts, "SessionTimeout")
	if err != nil {
		return nil, err
	}
	return &RootChainEventSessionTimeoutIterator{contract: _RootChainEvent.contract, event: "SessionTimeout", logs: logs, sub: sub}, nil
}

// WatchSessionTimeout is a free log subscription operation binding the contract event 0x2122ec719581cd177f225f59a1ee005831211196831f8f1ccffa817d2e7bd108.
//
// Solidity: e SessionTimeout(userActivated bool)
func (_RootChainEvent *RootChainEventFilterer) WatchSessionTimeout(opts *bind.WatchOpts, sink chan<- *RootChainEventSessionTimeout) (event.Subscription, error) {

	logs, sub, err := _RootChainEvent.contract.WatchLogs(opts, "SessionTimeout")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RootChainEventSessionTimeout)
				if err := _RootChainEvent.contract.UnpackLog(event, "SessionTimeout", log); err != nil {
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

// RootChainStorageABI is the input ABI used to generate the binding from.
const RootChainStorageABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"COST_URB_PREPARE\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CP_COMPUTATION\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastAppliedForkNumber\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"currentFork\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"COST_URB\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastAppliedERU\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"numEnterForORB\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"forks\",\"outputs\":[{\"name\":\"forkedBlock\",\"type\":\"uint64\"},{\"name\":\"firstEpoch\",\"type\":\"uint64\"},{\"name\":\"lastEpoch\",\"type\":\"uint64\"},{\"name\":\"firstBlock\",\"type\":\"uint64\"},{\"name\":\"lastBlock\",\"type\":\"uint64\"},{\"name\":\"lastFinalizedBlock\",\"type\":\"uint64\"},{\"name\":\"timestamp\",\"type\":\"uint64\"},{\"name\":\"firstEnterEpoch\",\"type\":\"uint64\"},{\"name\":\"lastEnterEpoch\",\"type\":\"uint64\"},{\"name\":\"nextBlockToRebase\",\"type\":\"uint64\"},{\"name\":\"rebased\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"operator\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastAppliedERO\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"firstFilledORENumber\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"development\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"COST_ERU\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"REQUEST_GAS\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"COST_NRB\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"NRELength\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CP_WITHHOLDING\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"COST_ORB\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"EROs\",\"outputs\":[{\"name\":\"timestamp\",\"type\":\"uint64\"},{\"name\":\"isExit\",\"type\":\"bool\"},{\"name\":\"isTransfer\",\"type\":\"bool\"},{\"name\":\"finalized\",\"type\":\"bool\"},{\"name\":\"challenged\",\"type\":\"bool\"},{\"name\":\"value\",\"type\":\"uint128\"},{\"name\":\"requestor\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"trieKey\",\"type\":\"bytes32\"},{\"name\":\"trieValue\",\"type\":\"bytes32\"},{\"name\":\"hash\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"URBs\",\"outputs\":[{\"name\":\"submitted\",\"type\":\"bool\"},{\"name\":\"numEnter\",\"type\":\"uint64\"},{\"name\":\"epochNumber\",\"type\":\"uint64\"},{\"name\":\"requestStart\",\"type\":\"uint64\"},{\"name\":\"requestEnd\",\"type\":\"uint64\"},{\"name\":\"trie\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PREPARE_TIMEOUT\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"COST_ERO\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"requestableContracts\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"NULL_ADDRESS\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"epochHandler\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"ORBs\",\"outputs\":[{\"name\":\"submitted\",\"type\":\"bool\"},{\"name\":\"numEnter\",\"type\":\"uint64\"},{\"name\":\"epochNumber\",\"type\":\"uint64\"},{\"name\":\"requestStart\",\"type\":\"uint64\"},{\"name\":\"requestEnd\",\"type\":\"uint64\"},{\"name\":\"trie\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"ERUs\",\"outputs\":[{\"name\":\"timestamp\",\"type\":\"uint64\"},{\"name\":\"isExit\",\"type\":\"bool\"},{\"name\":\"isTransfer\",\"type\":\"bool\"},{\"name\":\"finalized\",\"type\":\"bool\"},{\"name\":\"challenged\",\"type\":\"bool\"},{\"name\":\"value\",\"type\":\"uint128\"},{\"name\":\"requestor\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"trieKey\",\"type\":\"bytes32\"},{\"name\":\"trieValue\",\"type\":\"bytes32\"},{\"name\":\"hash\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastAppliedBlockNumber\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// RootChainStorageBin is the compiled bytecode used for deploying new contracts.
const RootChainStorageBin = `0x608060405234801561001057600080fd5b50610803806100206000396000f3006080604052600436106101695763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663033cfbed811461016e57806308c4fff014610195578063164bc2ae146101aa578063183d2d1c146101bf578063192adc5b146101d45780631f261d59146101e957806323691566146101fe5780634ba3a12614610213578063570ca7351461029a57806365d724bc146102cb57806372ecb9a8146102e05780637b929c27146102f85780638b5172d0146103215780638eb288ca1461033657806394be3aa51461016e578063ab96da2d1461034b578063b17fa6e914610360578063b2ae9ba81461016e578063b443f3cc14610375578063c0e860641461040b578063c2bc88fa1461046f578063d691acd81461016e578063da0185f814610484578063de0ce17d146104a5578063e7b88b80146104ba578063ea7f22a8146104cf578063f4f31de4146104e7578063fb788a27146104ff575b600080fd5b34801561017a57600080fd5b50610183610514565b60408051918252519081900360200190f35b3480156101a157600080fd5b50610183610520565b3480156101b657600080fd5b50610183610525565b3480156101cb57600080fd5b5061018361052b565b3480156101e057600080fd5b50610183610531565b3480156101f557600080fd5b5061018361053d565b34801561020a57600080fd5b50610183610543565b34801561021f57600080fd5b5061022b600435610549565b6040805167ffffffffffffffff9c8d1681529a8c1660208c0152988b168a8a0152968a1660608a0152948916608089015292881660a088015290871660c0870152861660e086015285166101008501529093166101208301529115156101408201529051908190036101600190f35b3480156102a657600080fd5b506102af6105e3565b60408051600160a060020a039092168252519081900360200190f35b3480156102d757600080fd5b506101836105f7565b3480156102ec57600080fd5b506101836004356105fd565b34801561030457600080fd5b5061030d61060f565b604080519115158252519081900360200190f35b34801561032d57600080fd5b50610183610618565b34801561034257600080fd5b50610183610624565b34801561035757600080fd5b5061018361062b565b34801561036c57600080fd5b50610183610631565b34801561038157600080fd5b5061038d600435610636565b6040805167ffffffffffffffff909c168c5299151560208c01529715158a8a015295151560608a015293151560808901526fffffffffffffffffffffffffffffffff90921660a0880152600160a060020a0390811660c08801521660e086015261010085015261012084015261014083015251908190036101600190f35b34801561041757600080fd5b506104236004356106fe565b60408051961515875267ffffffffffffffff95861660208801529385168685015291841660608601529092166080840152600160a060020a0390911660a0830152519081900360c00190f35b34801561047b57600080fd5b50610183610780565b34801561049057600080fd5b506102af600160a060020a0360043516610786565b3480156104b157600080fd5b506102af6107a1565b3480156104c657600080fd5b506102af6107a6565b3480156104db57600080fd5b506104236004356107b5565b3480156104f357600080fd5b5061038d6004356107c3565b34801561050b57600080fd5b506101836107d1565b67016345785d8a000081565b600181565b600b5481565b60035481565b670c7d713b49da000081565b600e5481565b600a5481565b60056020526000908152604090208054600182015460029092015467ffffffffffffffff808316936801000000000000000080850483169470010000000000000000000000000000000080820485169578010000000000000000000000000000000000000000000000009283900486169585811695858104821695848204831695909104821693838316939182049092169160ff9104168b565b6000546101009004600160a060020a031681565b600d5481565b60046020526000908152604090205481565b60005460ff1681565b6702c68af0bb14000081565b620186a081565b60025481565b600381565b600680548290811061064457fe5b600091825260209091206006909102018054600182015460028301546003840154600485015460059095015467ffffffffffffffff8516965068010000000000000000850460ff90811696690100000000000000000087048216966a010000000000000000000081048316966b0100000000000000000000008204909316956c010000000000000000000000009091046fffffffffffffffffffffffffffffffff1694600160a060020a039384169493909116929091908b565b600980548290811061070c57fe5b60009182526020909120600290910201805460019091015460ff8216925067ffffffffffffffff61010083048116926901000000000000000000810482169271010000000000000000000000000000000000909104821691811690600160a060020a03680100000000000000009091041686565b610e1081565b600f60205260009081526040902054600160a060020a031681565b600081565b600154600160a060020a031681565b600880548290811061070c57fe5b600780548290811061064457fe5b600c54815600a165627a7a7230582087f2547f36fc8a1fe4dbfa5c937dfb1cda519142940ca31ceb93fadd9fef11540029`

// DeployRootChainStorage deploys a new Ethereum contract, binding an instance of RootChainStorage to it.
func DeployRootChainStorage(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RootChainStorage, error) {
	parsed, err := abi.JSON(strings.NewReader(RootChainStorageABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RootChainStorageBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RootChainStorage{RootChainStorageCaller: RootChainStorageCaller{contract: contract}, RootChainStorageTransactor: RootChainStorageTransactor{contract: contract}, RootChainStorageFilterer: RootChainStorageFilterer{contract: contract}}, nil
}

// RootChainStorage is an auto generated Go binding around an Ethereum contract.
type RootChainStorage struct {
	RootChainStorageCaller     // Read-only binding to the contract
	RootChainStorageTransactor // Write-only binding to the contract
	RootChainStorageFilterer   // Log filterer for contract events
}

// RootChainStorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type RootChainStorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RootChainStorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RootChainStorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RootChainStorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RootChainStorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RootChainStorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RootChainStorageSession struct {
	Contract     *RootChainStorage // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RootChainStorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RootChainStorageCallerSession struct {
	Contract *RootChainStorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// RootChainStorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RootChainStorageTransactorSession struct {
	Contract     *RootChainStorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// RootChainStorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type RootChainStorageRaw struct {
	Contract *RootChainStorage // Generic contract binding to access the raw methods on
}

// RootChainStorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RootChainStorageCallerRaw struct {
	Contract *RootChainStorageCaller // Generic read-only contract binding to access the raw methods on
}

// RootChainStorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RootChainStorageTransactorRaw struct {
	Contract *RootChainStorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRootChainStorage creates a new instance of RootChainStorage, bound to a specific deployed contract.
func NewRootChainStorage(address common.Address, backend bind.ContractBackend) (*RootChainStorage, error) {
	contract, err := bindRootChainStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RootChainStorage{RootChainStorageCaller: RootChainStorageCaller{contract: contract}, RootChainStorageTransactor: RootChainStorageTransactor{contract: contract}, RootChainStorageFilterer: RootChainStorageFilterer{contract: contract}}, nil
}

// NewRootChainStorageCaller creates a new read-only instance of RootChainStorage, bound to a specific deployed contract.
func NewRootChainStorageCaller(address common.Address, caller bind.ContractCaller) (*RootChainStorageCaller, error) {
	contract, err := bindRootChainStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RootChainStorageCaller{contract: contract}, nil
}

// NewRootChainStorageTransactor creates a new write-only instance of RootChainStorage, bound to a specific deployed contract.
func NewRootChainStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*RootChainStorageTransactor, error) {
	contract, err := bindRootChainStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RootChainStorageTransactor{contract: contract}, nil
}

// NewRootChainStorageFilterer creates a new log filterer instance of RootChainStorage, bound to a specific deployed contract.
func NewRootChainStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*RootChainStorageFilterer, error) {
	contract, err := bindRootChainStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RootChainStorageFilterer{contract: contract}, nil
}

// bindRootChainStorage binds a generic wrapper to an already deployed contract.
func bindRootChainStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RootChainStorageABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RootChainStorage *RootChainStorageRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RootChainStorage.Contract.RootChainStorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RootChainStorage *RootChainStorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RootChainStorage.Contract.RootChainStorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RootChainStorage *RootChainStorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RootChainStorage.Contract.RootChainStorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RootChainStorage *RootChainStorageCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RootChainStorage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RootChainStorage *RootChainStorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RootChainStorage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RootChainStorage *RootChainStorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RootChainStorage.Contract.contract.Transact(opts, method, params...)
}

// COSTERO is a free data retrieval call binding the contract method 0xd691acd8.
//
// Solidity: function COST_ERO() constant returns(uint256)
func (_RootChainStorage *RootChainStorageCaller) COSTERO(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChainStorage.contract.Call(opts, out, "COST_ERO")
	return *ret0, err
}

// COSTERO is a free data retrieval call binding the contract method 0xd691acd8.
//
// Solidity: function COST_ERO() constant returns(uint256)
func (_RootChainStorage *RootChainStorageSession) COSTERO() (*big.Int, error) {
	return _RootChainStorage.Contract.COSTERO(&_RootChainStorage.CallOpts)
}

// COSTERO is a free data retrieval call binding the contract method 0xd691acd8.
//
// Solidity: function COST_ERO() constant returns(uint256)
func (_RootChainStorage *RootChainStorageCallerSession) COSTERO() (*big.Int, error) {
	return _RootChainStorage.Contract.COSTERO(&_RootChainStorage.CallOpts)
}

// COSTERU is a free data retrieval call binding the contract method 0x8b5172d0.
//
// Solidity: function COST_ERU() constant returns(uint256)
func (_RootChainStorage *RootChainStorageCaller) COSTERU(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChainStorage.contract.Call(opts, out, "COST_ERU")
	return *ret0, err
}

// COSTERU is a free data retrieval call binding the contract method 0x8b5172d0.
//
// Solidity: function COST_ERU() constant returns(uint256)
func (_RootChainStorage *RootChainStorageSession) COSTERU() (*big.Int, error) {
	return _RootChainStorage.Contract.COSTERU(&_RootChainStorage.CallOpts)
}

// COSTERU is a free data retrieval call binding the contract method 0x8b5172d0.
//
// Solidity: function COST_ERU() constant returns(uint256)
func (_RootChainStorage *RootChainStorageCallerSession) COSTERU() (*big.Int, error) {
	return _RootChainStorage.Contract.COSTERU(&_RootChainStorage.CallOpts)
}

// COSTNRB is a free data retrieval call binding the contract method 0x94be3aa5.
//
// Solidity: function COST_NRB() constant returns(uint256)
func (_RootChainStorage *RootChainStorageCaller) COSTNRB(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChainStorage.contract.Call(opts, out, "COST_NRB")
	return *ret0, err
}

// COSTNRB is a free data retrieval call binding the contract method 0x94be3aa5.
//
// Solidity: function COST_NRB() constant returns(uint256)
func (_RootChainStorage *RootChainStorageSession) COSTNRB() (*big.Int, error) {
	return _RootChainStorage.Contract.COSTNRB(&_RootChainStorage.CallOpts)
}

// COSTNRB is a free data retrieval call binding the contract method 0x94be3aa5.
//
// Solidity: function COST_NRB() constant returns(uint256)
func (_RootChainStorage *RootChainStorageCallerSession) COSTNRB() (*big.Int, error) {
	return _RootChainStorage.Contract.COSTNRB(&_RootChainStorage.CallOpts)
}

// COSTORB is a free data retrieval call binding the contract method 0xb2ae9ba8.
//
// Solidity: function COST_ORB() constant returns(uint256)
func (_RootChainStorage *RootChainStorageCaller) COSTORB(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChainStorage.contract.Call(opts, out, "COST_ORB")
	return *ret0, err
}

// COSTORB is a free data retrieval call binding the contract method 0xb2ae9ba8.
//
// Solidity: function COST_ORB() constant returns(uint256)
func (_RootChainStorage *RootChainStorageSession) COSTORB() (*big.Int, error) {
	return _RootChainStorage.Contract.COSTORB(&_RootChainStorage.CallOpts)
}

// COSTORB is a free data retrieval call binding the contract method 0xb2ae9ba8.
//
// Solidity: function COST_ORB() constant returns(uint256)
func (_RootChainStorage *RootChainStorageCallerSession) COSTORB() (*big.Int, error) {
	return _RootChainStorage.Contract.COSTORB(&_RootChainStorage.CallOpts)
}

// COSTURB is a free data retrieval call binding the contract method 0x192adc5b.
//
// Solidity: function COST_URB() constant returns(uint256)
func (_RootChainStorage *RootChainStorageCaller) COSTURB(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChainStorage.contract.Call(opts, out, "COST_URB")
	return *ret0, err
}

// COSTURB is a free data retrieval call binding the contract method 0x192adc5b.
//
// Solidity: function COST_URB() constant returns(uint256)
func (_RootChainStorage *RootChainStorageSession) COSTURB() (*big.Int, error) {
	return _RootChainStorage.Contract.COSTURB(&_RootChainStorage.CallOpts)
}

// COSTURB is a free data retrieval call binding the contract method 0x192adc5b.
//
// Solidity: function COST_URB() constant returns(uint256)
func (_RootChainStorage *RootChainStorageCallerSession) COSTURB() (*big.Int, error) {
	return _RootChainStorage.Contract.COSTURB(&_RootChainStorage.CallOpts)
}

// COSTURBPREPARE is a free data retrieval call binding the contract method 0x033cfbed.
//
// Solidity: function COST_URB_PREPARE() constant returns(uint256)
func (_RootChainStorage *RootChainStorageCaller) COSTURBPREPARE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChainStorage.contract.Call(opts, out, "COST_URB_PREPARE")
	return *ret0, err
}

// COSTURBPREPARE is a free data retrieval call binding the contract method 0x033cfbed.
//
// Solidity: function COST_URB_PREPARE() constant returns(uint256)
func (_RootChainStorage *RootChainStorageSession) COSTURBPREPARE() (*big.Int, error) {
	return _RootChainStorage.Contract.COSTURBPREPARE(&_RootChainStorage.CallOpts)
}

// COSTURBPREPARE is a free data retrieval call binding the contract method 0x033cfbed.
//
// Solidity: function COST_URB_PREPARE() constant returns(uint256)
func (_RootChainStorage *RootChainStorageCallerSession) COSTURBPREPARE() (*big.Int, error) {
	return _RootChainStorage.Contract.COSTURBPREPARE(&_RootChainStorage.CallOpts)
}

// CPCOMPUTATION is a free data retrieval call binding the contract method 0x08c4fff0.
//
// Solidity: function CP_COMPUTATION() constant returns(uint256)
func (_RootChainStorage *RootChainStorageCaller) CPCOMPUTATION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChainStorage.contract.Call(opts, out, "CP_COMPUTATION")
	return *ret0, err
}

// CPCOMPUTATION is a free data retrieval call binding the contract method 0x08c4fff0.
//
// Solidity: function CP_COMPUTATION() constant returns(uint256)
func (_RootChainStorage *RootChainStorageSession) CPCOMPUTATION() (*big.Int, error) {
	return _RootChainStorage.Contract.CPCOMPUTATION(&_RootChainStorage.CallOpts)
}

// CPCOMPUTATION is a free data retrieval call binding the contract method 0x08c4fff0.
//
// Solidity: function CP_COMPUTATION() constant returns(uint256)
func (_RootChainStorage *RootChainStorageCallerSession) CPCOMPUTATION() (*big.Int, error) {
	return _RootChainStorage.Contract.CPCOMPUTATION(&_RootChainStorage.CallOpts)
}

// CPWITHHOLDING is a free data retrieval call binding the contract method 0xb17fa6e9.
//
// Solidity: function CP_WITHHOLDING() constant returns(uint256)
func (_RootChainStorage *RootChainStorageCaller) CPWITHHOLDING(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChainStorage.contract.Call(opts, out, "CP_WITHHOLDING")
	return *ret0, err
}

// CPWITHHOLDING is a free data retrieval call binding the contract method 0xb17fa6e9.
//
// Solidity: function CP_WITHHOLDING() constant returns(uint256)
func (_RootChainStorage *RootChainStorageSession) CPWITHHOLDING() (*big.Int, error) {
	return _RootChainStorage.Contract.CPWITHHOLDING(&_RootChainStorage.CallOpts)
}

// CPWITHHOLDING is a free data retrieval call binding the contract method 0xb17fa6e9.
//
// Solidity: function CP_WITHHOLDING() constant returns(uint256)
func (_RootChainStorage *RootChainStorageCallerSession) CPWITHHOLDING() (*big.Int, error) {
	return _RootChainStorage.Contract.CPWITHHOLDING(&_RootChainStorage.CallOpts)
}

// EROs is a free data retrieval call binding the contract method 0xb443f3cc.
//
// Solidity: function EROs( uint256) constant returns(timestamp uint64, isExit bool, isTransfer bool, finalized bool, challenged bool, value uint128, requestor address, to address, trieKey bytes32, trieValue bytes32, hash bytes32)
func (_RootChainStorage *RootChainStorageCaller) EROs(opts *bind.CallOpts, arg0 *big.Int) (struct {
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
	err := _RootChainStorage.contract.Call(opts, out, "EROs", arg0)
	return *ret, err
}

// EROs is a free data retrieval call binding the contract method 0xb443f3cc.
//
// Solidity: function EROs( uint256) constant returns(timestamp uint64, isExit bool, isTransfer bool, finalized bool, challenged bool, value uint128, requestor address, to address, trieKey bytes32, trieValue bytes32, hash bytes32)
func (_RootChainStorage *RootChainStorageSession) EROs(arg0 *big.Int) (struct {
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
	return _RootChainStorage.Contract.EROs(&_RootChainStorage.CallOpts, arg0)
}

// EROs is a free data retrieval call binding the contract method 0xb443f3cc.
//
// Solidity: function EROs( uint256) constant returns(timestamp uint64, isExit bool, isTransfer bool, finalized bool, challenged bool, value uint128, requestor address, to address, trieKey bytes32, trieValue bytes32, hash bytes32)
func (_RootChainStorage *RootChainStorageCallerSession) EROs(arg0 *big.Int) (struct {
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
	return _RootChainStorage.Contract.EROs(&_RootChainStorage.CallOpts, arg0)
}

// ERUs is a free data retrieval call binding the contract method 0xf4f31de4.
//
// Solidity: function ERUs( uint256) constant returns(timestamp uint64, isExit bool, isTransfer bool, finalized bool, challenged bool, value uint128, requestor address, to address, trieKey bytes32, trieValue bytes32, hash bytes32)
func (_RootChainStorage *RootChainStorageCaller) ERUs(opts *bind.CallOpts, arg0 *big.Int) (struct {
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
	err := _RootChainStorage.contract.Call(opts, out, "ERUs", arg0)
	return *ret, err
}

// ERUs is a free data retrieval call binding the contract method 0xf4f31de4.
//
// Solidity: function ERUs( uint256) constant returns(timestamp uint64, isExit bool, isTransfer bool, finalized bool, challenged bool, value uint128, requestor address, to address, trieKey bytes32, trieValue bytes32, hash bytes32)
func (_RootChainStorage *RootChainStorageSession) ERUs(arg0 *big.Int) (struct {
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
	return _RootChainStorage.Contract.ERUs(&_RootChainStorage.CallOpts, arg0)
}

// ERUs is a free data retrieval call binding the contract method 0xf4f31de4.
//
// Solidity: function ERUs( uint256) constant returns(timestamp uint64, isExit bool, isTransfer bool, finalized bool, challenged bool, value uint128, requestor address, to address, trieKey bytes32, trieValue bytes32, hash bytes32)
func (_RootChainStorage *RootChainStorageCallerSession) ERUs(arg0 *big.Int) (struct {
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
	return _RootChainStorage.Contract.ERUs(&_RootChainStorage.CallOpts, arg0)
}

// NRELength is a free data retrieval call binding the contract method 0xab96da2d.
//
// Solidity: function NRELength() constant returns(uint256)
func (_RootChainStorage *RootChainStorageCaller) NRELength(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChainStorage.contract.Call(opts, out, "NRELength")
	return *ret0, err
}

// NRELength is a free data retrieval call binding the contract method 0xab96da2d.
//
// Solidity: function NRELength() constant returns(uint256)
func (_RootChainStorage *RootChainStorageSession) NRELength() (*big.Int, error) {
	return _RootChainStorage.Contract.NRELength(&_RootChainStorage.CallOpts)
}

// NRELength is a free data retrieval call binding the contract method 0xab96da2d.
//
// Solidity: function NRELength() constant returns(uint256)
func (_RootChainStorage *RootChainStorageCallerSession) NRELength() (*big.Int, error) {
	return _RootChainStorage.Contract.NRELength(&_RootChainStorage.CallOpts)
}

// NULLADDRESS is a free data retrieval call binding the contract method 0xde0ce17d.
//
// Solidity: function NULL_ADDRESS() constant returns(address)
func (_RootChainStorage *RootChainStorageCaller) NULLADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RootChainStorage.contract.Call(opts, out, "NULL_ADDRESS")
	return *ret0, err
}

// NULLADDRESS is a free data retrieval call binding the contract method 0xde0ce17d.
//
// Solidity: function NULL_ADDRESS() constant returns(address)
func (_RootChainStorage *RootChainStorageSession) NULLADDRESS() (common.Address, error) {
	return _RootChainStorage.Contract.NULLADDRESS(&_RootChainStorage.CallOpts)
}

// NULLADDRESS is a free data retrieval call binding the contract method 0xde0ce17d.
//
// Solidity: function NULL_ADDRESS() constant returns(address)
func (_RootChainStorage *RootChainStorageCallerSession) NULLADDRESS() (common.Address, error) {
	return _RootChainStorage.Contract.NULLADDRESS(&_RootChainStorage.CallOpts)
}

// ORBs is a free data retrieval call binding the contract method 0xea7f22a8.
//
// Solidity: function ORBs( uint256) constant returns(submitted bool, numEnter uint64, epochNumber uint64, requestStart uint64, requestEnd uint64, trie address)
func (_RootChainStorage *RootChainStorageCaller) ORBs(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Submitted    bool
	NumEnter     uint64
	EpochNumber  uint64
	RequestStart uint64
	RequestEnd   uint64
	Trie         common.Address
}, error) {
	ret := new(struct {
		Submitted    bool
		NumEnter     uint64
		EpochNumber  uint64
		RequestStart uint64
		RequestEnd   uint64
		Trie         common.Address
	})
	out := ret
	err := _RootChainStorage.contract.Call(opts, out, "ORBs", arg0)
	return *ret, err
}

// ORBs is a free data retrieval call binding the contract method 0xea7f22a8.
//
// Solidity: function ORBs( uint256) constant returns(submitted bool, numEnter uint64, epochNumber uint64, requestStart uint64, requestEnd uint64, trie address)
func (_RootChainStorage *RootChainStorageSession) ORBs(arg0 *big.Int) (struct {
	Submitted    bool
	NumEnter     uint64
	EpochNumber  uint64
	RequestStart uint64
	RequestEnd   uint64
	Trie         common.Address
}, error) {
	return _RootChainStorage.Contract.ORBs(&_RootChainStorage.CallOpts, arg0)
}

// ORBs is a free data retrieval call binding the contract method 0xea7f22a8.
//
// Solidity: function ORBs( uint256) constant returns(submitted bool, numEnter uint64, epochNumber uint64, requestStart uint64, requestEnd uint64, trie address)
func (_RootChainStorage *RootChainStorageCallerSession) ORBs(arg0 *big.Int) (struct {
	Submitted    bool
	NumEnter     uint64
	EpochNumber  uint64
	RequestStart uint64
	RequestEnd   uint64
	Trie         common.Address
}, error) {
	return _RootChainStorage.Contract.ORBs(&_RootChainStorage.CallOpts, arg0)
}

// PREPARETIMEOUT is a free data retrieval call binding the contract method 0xc2bc88fa.
//
// Solidity: function PREPARE_TIMEOUT() constant returns(uint256)
func (_RootChainStorage *RootChainStorageCaller) PREPARETIMEOUT(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChainStorage.contract.Call(opts, out, "PREPARE_TIMEOUT")
	return *ret0, err
}

// PREPARETIMEOUT is a free data retrieval call binding the contract method 0xc2bc88fa.
//
// Solidity: function PREPARE_TIMEOUT() constant returns(uint256)
func (_RootChainStorage *RootChainStorageSession) PREPARETIMEOUT() (*big.Int, error) {
	return _RootChainStorage.Contract.PREPARETIMEOUT(&_RootChainStorage.CallOpts)
}

// PREPARETIMEOUT is a free data retrieval call binding the contract method 0xc2bc88fa.
//
// Solidity: function PREPARE_TIMEOUT() constant returns(uint256)
func (_RootChainStorage *RootChainStorageCallerSession) PREPARETIMEOUT() (*big.Int, error) {
	return _RootChainStorage.Contract.PREPARETIMEOUT(&_RootChainStorage.CallOpts)
}

// REQUESTGAS is a free data retrieval call binding the contract method 0x8eb288ca.
//
// Solidity: function REQUEST_GAS() constant returns(uint256)
func (_RootChainStorage *RootChainStorageCaller) REQUESTGAS(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChainStorage.contract.Call(opts, out, "REQUEST_GAS")
	return *ret0, err
}

// REQUESTGAS is a free data retrieval call binding the contract method 0x8eb288ca.
//
// Solidity: function REQUEST_GAS() constant returns(uint256)
func (_RootChainStorage *RootChainStorageSession) REQUESTGAS() (*big.Int, error) {
	return _RootChainStorage.Contract.REQUESTGAS(&_RootChainStorage.CallOpts)
}

// REQUESTGAS is a free data retrieval call binding the contract method 0x8eb288ca.
//
// Solidity: function REQUEST_GAS() constant returns(uint256)
func (_RootChainStorage *RootChainStorageCallerSession) REQUESTGAS() (*big.Int, error) {
	return _RootChainStorage.Contract.REQUESTGAS(&_RootChainStorage.CallOpts)
}

// URBs is a free data retrieval call binding the contract method 0xc0e86064.
//
// Solidity: function URBs( uint256) constant returns(submitted bool, numEnter uint64, epochNumber uint64, requestStart uint64, requestEnd uint64, trie address)
func (_RootChainStorage *RootChainStorageCaller) URBs(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Submitted    bool
	NumEnter     uint64
	EpochNumber  uint64
	RequestStart uint64
	RequestEnd   uint64
	Trie         common.Address
}, error) {
	ret := new(struct {
		Submitted    bool
		NumEnter     uint64
		EpochNumber  uint64
		RequestStart uint64
		RequestEnd   uint64
		Trie         common.Address
	})
	out := ret
	err := _RootChainStorage.contract.Call(opts, out, "URBs", arg0)
	return *ret, err
}

// URBs is a free data retrieval call binding the contract method 0xc0e86064.
//
// Solidity: function URBs( uint256) constant returns(submitted bool, numEnter uint64, epochNumber uint64, requestStart uint64, requestEnd uint64, trie address)
func (_RootChainStorage *RootChainStorageSession) URBs(arg0 *big.Int) (struct {
	Submitted    bool
	NumEnter     uint64
	EpochNumber  uint64
	RequestStart uint64
	RequestEnd   uint64
	Trie         common.Address
}, error) {
	return _RootChainStorage.Contract.URBs(&_RootChainStorage.CallOpts, arg0)
}

// URBs is a free data retrieval call binding the contract method 0xc0e86064.
//
// Solidity: function URBs( uint256) constant returns(submitted bool, numEnter uint64, epochNumber uint64, requestStart uint64, requestEnd uint64, trie address)
func (_RootChainStorage *RootChainStorageCallerSession) URBs(arg0 *big.Int) (struct {
	Submitted    bool
	NumEnter     uint64
	EpochNumber  uint64
	RequestStart uint64
	RequestEnd   uint64
	Trie         common.Address
}, error) {
	return _RootChainStorage.Contract.URBs(&_RootChainStorage.CallOpts, arg0)
}

// CurrentFork is a free data retrieval call binding the contract method 0x183d2d1c.
//
// Solidity: function currentFork() constant returns(uint256)
func (_RootChainStorage *RootChainStorageCaller) CurrentFork(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChainStorage.contract.Call(opts, out, "currentFork")
	return *ret0, err
}

// CurrentFork is a free data retrieval call binding the contract method 0x183d2d1c.
//
// Solidity: function currentFork() constant returns(uint256)
func (_RootChainStorage *RootChainStorageSession) CurrentFork() (*big.Int, error) {
	return _RootChainStorage.Contract.CurrentFork(&_RootChainStorage.CallOpts)
}

// CurrentFork is a free data retrieval call binding the contract method 0x183d2d1c.
//
// Solidity: function currentFork() constant returns(uint256)
func (_RootChainStorage *RootChainStorageCallerSession) CurrentFork() (*big.Int, error) {
	return _RootChainStorage.Contract.CurrentFork(&_RootChainStorage.CallOpts)
}

// Development is a free data retrieval call binding the contract method 0x7b929c27.
//
// Solidity: function development() constant returns(bool)
func (_RootChainStorage *RootChainStorageCaller) Development(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _RootChainStorage.contract.Call(opts, out, "development")
	return *ret0, err
}

// Development is a free data retrieval call binding the contract method 0x7b929c27.
//
// Solidity: function development() constant returns(bool)
func (_RootChainStorage *RootChainStorageSession) Development() (bool, error) {
	return _RootChainStorage.Contract.Development(&_RootChainStorage.CallOpts)
}

// Development is a free data retrieval call binding the contract method 0x7b929c27.
//
// Solidity: function development() constant returns(bool)
func (_RootChainStorage *RootChainStorageCallerSession) Development() (bool, error) {
	return _RootChainStorage.Contract.Development(&_RootChainStorage.CallOpts)
}

// EpochHandler is a free data retrieval call binding the contract method 0xe7b88b80.
//
// Solidity: function epochHandler() constant returns(address)
func (_RootChainStorage *RootChainStorageCaller) EpochHandler(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RootChainStorage.contract.Call(opts, out, "epochHandler")
	return *ret0, err
}

// EpochHandler is a free data retrieval call binding the contract method 0xe7b88b80.
//
// Solidity: function epochHandler() constant returns(address)
func (_RootChainStorage *RootChainStorageSession) EpochHandler() (common.Address, error) {
	return _RootChainStorage.Contract.EpochHandler(&_RootChainStorage.CallOpts)
}

// EpochHandler is a free data retrieval call binding the contract method 0xe7b88b80.
//
// Solidity: function epochHandler() constant returns(address)
func (_RootChainStorage *RootChainStorageCallerSession) EpochHandler() (common.Address, error) {
	return _RootChainStorage.Contract.EpochHandler(&_RootChainStorage.CallOpts)
}

// FirstFilledORENumber is a free data retrieval call binding the contract method 0x72ecb9a8.
//
// Solidity: function firstFilledORENumber( uint256) constant returns(uint256)
func (_RootChainStorage *RootChainStorageCaller) FirstFilledORENumber(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChainStorage.contract.Call(opts, out, "firstFilledORENumber", arg0)
	return *ret0, err
}

// FirstFilledORENumber is a free data retrieval call binding the contract method 0x72ecb9a8.
//
// Solidity: function firstFilledORENumber( uint256) constant returns(uint256)
func (_RootChainStorage *RootChainStorageSession) FirstFilledORENumber(arg0 *big.Int) (*big.Int, error) {
	return _RootChainStorage.Contract.FirstFilledORENumber(&_RootChainStorage.CallOpts, arg0)
}

// FirstFilledORENumber is a free data retrieval call binding the contract method 0x72ecb9a8.
//
// Solidity: function firstFilledORENumber( uint256) constant returns(uint256)
func (_RootChainStorage *RootChainStorageCallerSession) FirstFilledORENumber(arg0 *big.Int) (*big.Int, error) {
	return _RootChainStorage.Contract.FirstFilledORENumber(&_RootChainStorage.CallOpts, arg0)
}

// Forks is a free data retrieval call binding the contract method 0x4ba3a126.
//
// Solidity: function forks( uint256) constant returns(forkedBlock uint64, firstEpoch uint64, lastEpoch uint64, firstBlock uint64, lastBlock uint64, lastFinalizedBlock uint64, timestamp uint64, firstEnterEpoch uint64, lastEnterEpoch uint64, nextBlockToRebase uint64, rebased bool)
func (_RootChainStorage *RootChainStorageCaller) Forks(opts *bind.CallOpts, arg0 *big.Int) (struct {
	ForkedBlock        uint64
	FirstEpoch         uint64
	LastEpoch          uint64
	FirstBlock         uint64
	LastBlock          uint64
	LastFinalizedBlock uint64
	Timestamp          uint64
	FirstEnterEpoch    uint64
	LastEnterEpoch     uint64
	NextBlockToRebase  uint64
	Rebased            bool
}, error) {
	ret := new(struct {
		ForkedBlock        uint64
		FirstEpoch         uint64
		LastEpoch          uint64
		FirstBlock         uint64
		LastBlock          uint64
		LastFinalizedBlock uint64
		Timestamp          uint64
		FirstEnterEpoch    uint64
		LastEnterEpoch     uint64
		NextBlockToRebase  uint64
		Rebased            bool
	})
	out := ret
	err := _RootChainStorage.contract.Call(opts, out, "forks", arg0)
	return *ret, err
}

// Forks is a free data retrieval call binding the contract method 0x4ba3a126.
//
// Solidity: function forks( uint256) constant returns(forkedBlock uint64, firstEpoch uint64, lastEpoch uint64, firstBlock uint64, lastBlock uint64, lastFinalizedBlock uint64, timestamp uint64, firstEnterEpoch uint64, lastEnterEpoch uint64, nextBlockToRebase uint64, rebased bool)
func (_RootChainStorage *RootChainStorageSession) Forks(arg0 *big.Int) (struct {
	ForkedBlock        uint64
	FirstEpoch         uint64
	LastEpoch          uint64
	FirstBlock         uint64
	LastBlock          uint64
	LastFinalizedBlock uint64
	Timestamp          uint64
	FirstEnterEpoch    uint64
	LastEnterEpoch     uint64
	NextBlockToRebase  uint64
	Rebased            bool
}, error) {
	return _RootChainStorage.Contract.Forks(&_RootChainStorage.CallOpts, arg0)
}

// Forks is a free data retrieval call binding the contract method 0x4ba3a126.
//
// Solidity: function forks( uint256) constant returns(forkedBlock uint64, firstEpoch uint64, lastEpoch uint64, firstBlock uint64, lastBlock uint64, lastFinalizedBlock uint64, timestamp uint64, firstEnterEpoch uint64, lastEnterEpoch uint64, nextBlockToRebase uint64, rebased bool)
func (_RootChainStorage *RootChainStorageCallerSession) Forks(arg0 *big.Int) (struct {
	ForkedBlock        uint64
	FirstEpoch         uint64
	LastEpoch          uint64
	FirstBlock         uint64
	LastBlock          uint64
	LastFinalizedBlock uint64
	Timestamp          uint64
	FirstEnterEpoch    uint64
	LastEnterEpoch     uint64
	NextBlockToRebase  uint64
	Rebased            bool
}, error) {
	return _RootChainStorage.Contract.Forks(&_RootChainStorage.CallOpts, arg0)
}

// LastAppliedBlockNumber is a free data retrieval call binding the contract method 0xfb788a27.
//
// Solidity: function lastAppliedBlockNumber() constant returns(uint256)
func (_RootChainStorage *RootChainStorageCaller) LastAppliedBlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChainStorage.contract.Call(opts, out, "lastAppliedBlockNumber")
	return *ret0, err
}

// LastAppliedBlockNumber is a free data retrieval call binding the contract method 0xfb788a27.
//
// Solidity: function lastAppliedBlockNumber() constant returns(uint256)
func (_RootChainStorage *RootChainStorageSession) LastAppliedBlockNumber() (*big.Int, error) {
	return _RootChainStorage.Contract.LastAppliedBlockNumber(&_RootChainStorage.CallOpts)
}

// LastAppliedBlockNumber is a free data retrieval call binding the contract method 0xfb788a27.
//
// Solidity: function lastAppliedBlockNumber() constant returns(uint256)
func (_RootChainStorage *RootChainStorageCallerSession) LastAppliedBlockNumber() (*big.Int, error) {
	return _RootChainStorage.Contract.LastAppliedBlockNumber(&_RootChainStorage.CallOpts)
}

// LastAppliedERO is a free data retrieval call binding the contract method 0x65d724bc.
//
// Solidity: function lastAppliedERO() constant returns(uint256)
func (_RootChainStorage *RootChainStorageCaller) LastAppliedERO(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChainStorage.contract.Call(opts, out, "lastAppliedERO")
	return *ret0, err
}

// LastAppliedERO is a free data retrieval call binding the contract method 0x65d724bc.
//
// Solidity: function lastAppliedERO() constant returns(uint256)
func (_RootChainStorage *RootChainStorageSession) LastAppliedERO() (*big.Int, error) {
	return _RootChainStorage.Contract.LastAppliedERO(&_RootChainStorage.CallOpts)
}

// LastAppliedERO is a free data retrieval call binding the contract method 0x65d724bc.
//
// Solidity: function lastAppliedERO() constant returns(uint256)
func (_RootChainStorage *RootChainStorageCallerSession) LastAppliedERO() (*big.Int, error) {
	return _RootChainStorage.Contract.LastAppliedERO(&_RootChainStorage.CallOpts)
}

// LastAppliedERU is a free data retrieval call binding the contract method 0x1f261d59.
//
// Solidity: function lastAppliedERU() constant returns(uint256)
func (_RootChainStorage *RootChainStorageCaller) LastAppliedERU(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChainStorage.contract.Call(opts, out, "lastAppliedERU")
	return *ret0, err
}

// LastAppliedERU is a free data retrieval call binding the contract method 0x1f261d59.
//
// Solidity: function lastAppliedERU() constant returns(uint256)
func (_RootChainStorage *RootChainStorageSession) LastAppliedERU() (*big.Int, error) {
	return _RootChainStorage.Contract.LastAppliedERU(&_RootChainStorage.CallOpts)
}

// LastAppliedERU is a free data retrieval call binding the contract method 0x1f261d59.
//
// Solidity: function lastAppliedERU() constant returns(uint256)
func (_RootChainStorage *RootChainStorageCallerSession) LastAppliedERU() (*big.Int, error) {
	return _RootChainStorage.Contract.LastAppliedERU(&_RootChainStorage.CallOpts)
}

// LastAppliedForkNumber is a free data retrieval call binding the contract method 0x164bc2ae.
//
// Solidity: function lastAppliedForkNumber() constant returns(uint256)
func (_RootChainStorage *RootChainStorageCaller) LastAppliedForkNumber(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChainStorage.contract.Call(opts, out, "lastAppliedForkNumber")
	return *ret0, err
}

// LastAppliedForkNumber is a free data retrieval call binding the contract method 0x164bc2ae.
//
// Solidity: function lastAppliedForkNumber() constant returns(uint256)
func (_RootChainStorage *RootChainStorageSession) LastAppliedForkNumber() (*big.Int, error) {
	return _RootChainStorage.Contract.LastAppliedForkNumber(&_RootChainStorage.CallOpts)
}

// LastAppliedForkNumber is a free data retrieval call binding the contract method 0x164bc2ae.
//
// Solidity: function lastAppliedForkNumber() constant returns(uint256)
func (_RootChainStorage *RootChainStorageCallerSession) LastAppliedForkNumber() (*big.Int, error) {
	return _RootChainStorage.Contract.LastAppliedForkNumber(&_RootChainStorage.CallOpts)
}

// NumEnterForORB is a free data retrieval call binding the contract method 0x23691566.
//
// Solidity: function numEnterForORB() constant returns(uint256)
func (_RootChainStorage *RootChainStorageCaller) NumEnterForORB(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChainStorage.contract.Call(opts, out, "numEnterForORB")
	return *ret0, err
}

// NumEnterForORB is a free data retrieval call binding the contract method 0x23691566.
//
// Solidity: function numEnterForORB() constant returns(uint256)
func (_RootChainStorage *RootChainStorageSession) NumEnterForORB() (*big.Int, error) {
	return _RootChainStorage.Contract.NumEnterForORB(&_RootChainStorage.CallOpts)
}

// NumEnterForORB is a free data retrieval call binding the contract method 0x23691566.
//
// Solidity: function numEnterForORB() constant returns(uint256)
func (_RootChainStorage *RootChainStorageCallerSession) NumEnterForORB() (*big.Int, error) {
	return _RootChainStorage.Contract.NumEnterForORB(&_RootChainStorage.CallOpts)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() constant returns(address)
func (_RootChainStorage *RootChainStorageCaller) Operator(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RootChainStorage.contract.Call(opts, out, "operator")
	return *ret0, err
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() constant returns(address)
func (_RootChainStorage *RootChainStorageSession) Operator() (common.Address, error) {
	return _RootChainStorage.Contract.Operator(&_RootChainStorage.CallOpts)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() constant returns(address)
func (_RootChainStorage *RootChainStorageCallerSession) Operator() (common.Address, error) {
	return _RootChainStorage.Contract.Operator(&_RootChainStorage.CallOpts)
}

// RequestableContracts is a free data retrieval call binding the contract method 0xda0185f8.
//
// Solidity: function requestableContracts( address) constant returns(address)
func (_RootChainStorage *RootChainStorageCaller) RequestableContracts(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RootChainStorage.contract.Call(opts, out, "requestableContracts", arg0)
	return *ret0, err
}

// RequestableContracts is a free data retrieval call binding the contract method 0xda0185f8.
//
// Solidity: function requestableContracts( address) constant returns(address)
func (_RootChainStorage *RootChainStorageSession) RequestableContracts(arg0 common.Address) (common.Address, error) {
	return _RootChainStorage.Contract.RequestableContracts(&_RootChainStorage.CallOpts, arg0)
}

// RequestableContracts is a free data retrieval call binding the contract method 0xda0185f8.
//
// Solidity: function requestableContracts( address) constant returns(address)
func (_RootChainStorage *RootChainStorageCallerSession) RequestableContracts(arg0 common.Address) (common.Address, error) {
	return _RootChainStorage.Contract.RequestableContracts(&_RootChainStorage.CallOpts, arg0)
}

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
const SafeMathBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a7230582077dd01411fe76a4319be092b6370ba52ea363d77483720cb3df4667b4db983540029`

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
