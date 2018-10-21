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
const DataABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"NA_TX_GAS_PRICE\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"V\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"NA_TX_GAS_LIMIT\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"APPLY_IN_ROOTCHAIN_SIGNATURE\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"APPLY_IN_CHILDCHAIN_SIGNATURE\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"NA\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// DataBin is the compiled bytecode used for deploying new contracts.
const DataBin = `0x610220610030600b82828239805160001a6073146000811461002057610022565bfe5b5030600052607381538281f300730000000000000000000000000000000000000000301460806040526004361061008e5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416631927ac588114610093578063205db2c1146100b157806390e84f56146100b9578063a7b6ae28146100c1578063a89ca766146100d6578063ab73ff05146100de575b600080fd5b61009b6100f3565b6040516100a89190610197565b60405180910390f35b61009b6100fb565b61009b610100565b6100c9610107565b6040516100a89190610189565b6100c961012b565b6100e661014f565b6040516100a89190610175565b633b9aca0081565b604381565b620186a081565b7fd9afd3a90000000000000000000000000000000000000000000000000000000081565b7fe904e3d90000000000000000000000000000000000000000000000000000000081565b600081565b61015d816101a5565b82525050565b61015d816101be565b61015d816101e3565b602081016101838284610154565b92915050565b602081016101838284610163565b60208101610183828461016c565b73ffffffffffffffffffffffffffffffffffffffff1690565b7fffffffff000000000000000000000000000000000000000000000000000000001690565b905600a265627a7a723058205665d2dbfc2f467d22635f8ed716cc0c1147e48c39fb5b62c98395f3b497407a6c6578706572696d656e74616cf50037`

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

// V is a free data retrieval call binding the contract method 0x205db2c1.
//
// Solidity: function V() constant returns(uint256)
func (_Data *DataCaller) V(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Data.contract.Call(opts, out, "V")
	return *ret0, err
}

// V is a free data retrieval call binding the contract method 0x205db2c1.
//
// Solidity: function V() constant returns(uint256)
func (_Data *DataSession) V() (*big.Int, error) {
	return _Data.Contract.V(&_Data.CallOpts)
}

// V is a free data retrieval call binding the contract method 0x205db2c1.
//
// Solidity: function V() constant returns(uint256)
func (_Data *DataCallerSession) V() (*big.Int, error) {
	return _Data.Contract.V(&_Data.CallOpts)
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

// PatriciaTreeABI is the input ABI used to generate the binding from.
const PatriciaTreeABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"key\",\"type\":\"bytes\"},{\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"insert\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"getNode\",\"outputs\":[{\"name\":\"nodes\",\"type\":\"bytes32[2]\"},{\"name\":\"labelDatas\",\"type\":\"bytes32[2]\"},{\"name\":\"labelLengths\",\"type\":\"uint256[2]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"key\",\"type\":\"bytes\"}],\"name\":\"getProof\",\"outputs\":[{\"name\":\"branchMask\",\"type\":\"uint256\"},{\"name\":\"_siblings\",\"type\":\"bytes32[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getRootHash\",\"outputs\":[{\"name\":\"h\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getRootEdge\",\"outputs\":[{\"name\":\"node\",\"type\":\"bytes32\"},{\"name\":\"labelData\",\"type\":\"bytes32\"},{\"name\":\"labelLength\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"rootHash\",\"type\":\"bytes32\"},{\"name\":\"key\",\"type\":\"bytes\"},{\"name\":\"value\",\"type\":\"bytes\"},{\"name\":\"branchMask\",\"type\":\"uint256\"},{\"name\":\"siblings\",\"type\":\"bytes32[]\"}],\"name\":\"verifyProof\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// PatriciaTreeBin is the compiled bytecode used for deploying new contracts.
const PatriciaTreeBin = `0x608060405234801561001057600080fd5b506114fd806100206000396000f3006080604052600436106100775763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166320ba5b60811461007c57806350c946fe1461009e578063693ac4fb146100d657806380759f1f14610104578063a43914da14610126578063f7e498f61461014a575b600080fd5b34801561008857600080fd5b5061009c61009736600461123a565b610177565b005b3480156100aa57600080fd5b506100be6100b9366004611122565b61018d565b6040516100cd9392919061139d565b60405180910390f35b3480156100e257600080fd5b506100f66100f1366004611205565b610278565b6040516100cd92919061140f565b34801561011057600080fd5b5061011961052b565b6040516100cd91906113d9565b34801561013257600080fd5b5061013b610531565b6040516100cd939291906113e7565b34801561015657600080fd5b5061016a610165366004611148565b61057b565b6040516100cd91906113c5565b6101896000838363ffffffff6107dc16565b5050565b610195610f62565b61019d610f62565b6101a5610f62565b6101ad610f7d565b6000858152600460209081526040808320815160608101909252909290918391908201908390600290835b828210156102265760408051808201825260038402860180548252825180840190935260018082015484526002909101546020848101919091528083019390935290835290920191016101d8565b5050509152505080515151855280515160209081015151855281515181015181015184528151810151518682015281518101518101515185820152905181015181015181015190830152509193909250565b60006060610284610f96565b61028c610fad565b610294610fc8565b60008061029f610f96565b6102a7610f96565b60006102b1610f96565b6000805415156102c057600080fd5b60408051908101604052808e6040518082805190602001908083835b602083106102fb5780518252601f1990920191602091820191016102dc565b518151602093840361010090810a600019018019909316929091169190911790915260408051939095018390039092208752958601525080518082018252600154815281518083019092526002548252600354828601529384015250919b509950505b6020890151610374908b9063ffffffff61095216565b6020808c0151810151908301519297509095501461038e57fe5b6020840151151561039e576104a7565b60208501519690960160ff81900360020a9b909b179a600101956103c184610980565b8a5160009081526004602052604090209194509250610427906001859003600281106103e957fe5b6040805180820182526003929092029290920180548252825180840190935260018101548352600201546020838101919091528101919091526109e7565b60018701968990610100811061043957fe5b6020908102919091019190915289516000908152600490915260409020836002811061046157fe5b604080518082018252600392909202929092018054825282518084019093526001810154835260020154602083810191909152810191909152919950909750889061035e565b600086111561051c57856040519080825280602002602001820160405280156104da578160200160208202803883390190505b509a50600090505b8581101561051c57878161010081106104f757fe5b60200201518b8281518110151561050a57fe5b602090810290910101526001016104e2565b50505050505050505050915091565b60005490565b600080600061053e610fad565b5050604080518082018252600154815281518083019092526002548252600354602083810191825282018390529051915190519194909350909150565b6000610585610f96565b61058d610fad565b600080600061059a610f62565b60408051908101604052808c6040518082805190602001908083835b602083106105d55780518252601f1990920191602091820191016105b6565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405180910390206000191681526020016101008152509550896040518082805190602001908083835b602083106106435780518252601f199092019160209182019101610624565b5181516020939093036101000a600019018019909116921691909117905260405192018290039091208852506000955050505b88156107af5761068589610a84565b60ff1692508260019060020a0219891698506106ad8360ff0387610add90919063ffffffff16565b602087018190529096506106c090610980565b602087015291506106d0856109e7565b8183600281106106dc57fe5b602002015287518890858103600019019081106106f557fe5b90602001906020020151818360010360028110151561071057fe5b6020908102919091019190915281518282015160408051808501939093528281019190915280518083038201815260609092019081905281519192909182918401908083835b602083106107755780518252601f199092019160209182019101610756565b5181516020939093036101000a60001901801990911692169190911790526040519201829003909120885250505060019390930192610676565b602085018690526107bf856109e7565b8c146107ca57600080fd5b5060019b9a5050505050505050505050565b6107e4610f96565b60006107ee610fad565b6040805190810160405280866040518082805190602001908083835b602083106108295780518252601f19909201916020918201910161080a565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405180910390206000191681526020016101008152509250836040518082805190602001908083835b602083106108975780518252601f199092019160209182019101610878565b5181516020939093036101000a600019018019909116921691909117905260405192018290039091208954909550151592506108df9150505760208101839052818152610922565b60408051808201825260018801548152815180830190925260028801548252600388015460208381019190915281019190915261091f9087908585610b51565b90505b61092b816109e7565b86558051600187015560209081015180516002880155015160039095019490945550505050565b61095a610f96565b610962610f96565b610975846109708686610d8e565b610add565b915091509250929050565b600061098a610f96565b602083015160001061099b57600080fd5b50508051604080518082019091528251600202815260209283015160001901928101929092527f8000000000000000000000000000000000000000000000000000000000000000900491565b8051602080830151808201519051604080518085019590955284810192909252606080850191909152815180850390910181526080909301908190528251600093928291908401908083835b60208310610a525780518252601f199092019160209182019101610a33565b5181516020939093036101000a6000190180199091169216919091179052604051920182900390912095945050505050565b60008080831515610a9457600080fd5b5082905060805b600160ff821610610ad65760001960ff821660020a0182161515610ac9579182019160ff811660020a909104905b600260ff90911604610a9b565b5050919050565b610ae5610f96565b610aed610f96565b83602001518311158015610b0357506101008311155b1515610b0b57fe5b60208201839052821515610b225760008252610b34565b835160ff84900360020a600119021682525b60208085015184900390820152925160029290920a909102825291565b610b59610fad565b610b61610f96565b610b69610f96565b6000610b73610f96565b6000610b7d610f7d565b610b85610f7d565b6020808c0151810151908b01511015610b9a57fe5b610ba88a8c60200151610952565b602081015191985096501515610bc057889250610d6a565b6020808c01518101519088015110610cde576020860151600110610be057fe5b8a51600090815260048d0160209081526040808320815160608101909252909290918391908201908390600290835b82821015610c5d576040805180820182526003840286018054825282518084019093526001808201548452600290910154602084810191909152808301939093529083529092019101610c0f565b50505050815250509150610c7086610980565b83519196509450610c94908d908760028110610c8857fe5b6020020151868c610b51565b82518660028110610ca157fe5b602090810291909101919091528b51600090815260048e019091526040812090610ccb8282610fe9565b5050610cd78c83610e00565b9250610d6a565b610ce786610980565b604080518082019091528b8152602081018290528351929750909550908660028110610d0f57fe5b602002018190525060408051908101604052808c60000151600019168152602001610d458d602001518a60200151600101610e66565b90528151600187900360028110610d5857fe5b6020020152610d678c82610e00565b92505b50506040805180820190915290815260208101949094525091979650505050505050565b60008060008360200151856020015110610dac578360200151610db2565b84602001515b9150811515610dc45760009250610df8565b50825184511861010082900360020a60000316801515610de657819250610df8565b610def81610e9a565b60ff0360ff1692505b505092915050565b600080610e0c83610eed565b83515160008281526004968701602090815260409091208251815591810151805160018401558101516002830155945185015180516003830155850151805196820196909655949093015160059094019390935550919050565b610e6e610f96565b6020830151821115610e7f57600080fd5b60208084015183900390820152915160029190910a02815290565b60008080831515610eaa57600080fd5b5082905060805b600160ff821610610ad65760ff811660020a600019810102821615610ee0579182019160ff811660020a909104905b600260ff90911604610eb1565b8051600090610f0290825b60200201516109e7565b8251610f0f906001610ef8565b6040805160208082019490945280820192909252805180830382018152606090920190819052815191929091829184019080838360208310610a525780518252601f199092019160209182019101610a33565b60408051808201825290600290829080388339509192915050565b60c060405190810160405280610f91611013565b905290565b604080518082019091526000808252602082015290565b60408051606081019091526000815260208101610f91610f96565b61200060405190810160405280610100906020820280388339509192915050565b50600080825560018201819055600282018190556003820181905560048201819055600590910155565b60c0604051908101604052806002905b61102b610fad565b8152602001906001900390816110235790505090565b6000601f8201831361105257600080fd5b813561106561106082611456565b61142f565b9150818183526020840193506020810190508385602084028201111561108a57600080fd5b60005b838110156110b657816110a088826110c0565b845250602092830192919091019060010161108d565b5050505092915050565b60006110cc823561149f565b9392505050565b6000601f820183136110e457600080fd5b81356110f261106082611477565b9150808252602083016020830185838301111561110e57600080fd5b6111198382846114b7565b50505092915050565b60006020828403121561113457600080fd5b600061114084846110c0565b949350505050565b600080600080600060a0868803121561116057600080fd5b600061116c88886110c0565b955050602086013567ffffffffffffffff81111561118957600080fd5b611195888289016110d3565b945050604086013567ffffffffffffffff8111156111b257600080fd5b6111be888289016110d3565b93505060606111cf888289016110c0565b925050608086013567ffffffffffffffff8111156111ec57600080fd5b6111f888828901611041565b9150509295509295909350565b60006020828403121561121757600080fd5b813567ffffffffffffffff81111561122e57600080fd5b611140848285016110d3565b6000806040838503121561124d57600080fd5b823567ffffffffffffffff81111561126457600080fd5b611270858286016110d3565b925050602083013567ffffffffffffffff81111561128d57600080fd5b611299858286016110d3565b9150509250929050565b6112ac816114a8565b6112b58261149f565b60005b828110156112e5576112cb858351611394565b6112d4826114a2565b6020959095019491506001016112b8565b5050505050565b60006112f7826114ae565b808452602084019350611309836114a2565b60005b828110156113395761131f868351611394565b611328826114a2565b60209690960195915060010161130c565b5093949350505050565b61134c816114a8565b6113558261149f565b60005b828110156112e55761136b858351611394565b611374826114a2565b602095909501949150600101611358565b61138e816114b2565b82525050565b61138e8161149f565b60c081016113ab82866112a3565b6113b860408301856112a3565b6111406080830184611343565b602081016113d38284611385565b92915050565b602081016113d38284611394565b606081016113f58286611394565b6114026020830185611394565b6111406040830184611394565b6040810161141d8285611394565b818103602083015261114081846112ec565b60405181810167ffffffffffffffff8111828210171561144e57600080fd5b604052919050565b600067ffffffffffffffff82111561146d57600080fd5b5060209081020190565b600067ffffffffffffffff82111561148e57600080fd5b506020601f91909101601f19160190565b90565b60200190565b50600290565b5190565b151590565b828183375060009101525600a265627a7a72305820910360ca0293b325436c58e5dd43837d4b06187240d93f41d7aaedf1a71ac5026c6578706572696d656e74616cf50037`

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
const RLPBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a72305820bafe640993f280e5e4ee91ea453f50f96545dd8fa5298b941cff5366d4c882ff0029`

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
const RootChainABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"COST_URB_PREPARE\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CP_COMPUTATION\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"epochs\",\"outputs\":[{\"name\":\"requestStart\",\"type\":\"uint64\"},{\"name\":\"requestEnd\",\"type\":\"uint64\"},{\"name\":\"startBlockNumber\",\"type\":\"uint64\"},{\"name\":\"endBlockNumber\",\"type\":\"uint64\"},{\"name\":\"forkedBlockNumber\",\"type\":\"uint64\"},{\"name\":\"firstRequestBlockId\",\"type\":\"uint64\"},{\"name\":\"limit\",\"type\":\"uint64\"},{\"name\":\"timestamp\",\"type\":\"uint64\"},{\"name\":\"isEmpty\",\"type\":\"bool\"},{\"name\":\"initialized\",\"type\":\"bool\"},{\"name\":\"isRequest\",\"type\":\"bool\"},{\"name\":\"userActivated\",\"type\":\"bool\"},{\"name\":\"finalized\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastAppliedForkNumber\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"currentFork\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"COST_URB\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastAppliedERU\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"firstEpoch\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_isTransfer\",\"type\":\"bool\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_trieKey\",\"type\":\"bytes32\"},{\"name\":\"_trieValue\",\"type\":\"bytes32\"}],\"name\":\"startEnter\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_isTransfer\",\"type\":\"bool\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_trieKey\",\"type\":\"bytes32\"},{\"name\":\"_trieValue\",\"type\":\"bytes32\"}],\"name\":\"makeERU\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"operator\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"NRBEpochLength\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"applyRequest\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_blockNumber\",\"type\":\"uint256\"},{\"name\":\"_key\",\"type\":\"bytes\"},{\"name\":\"_txByte\",\"type\":\"bytes\"},{\"name\":\"_branchMask\",\"type\":\"uint256\"},{\"name\":\"_siblings\",\"type\":\"bytes32[]\"}],\"name\":\"challengeNullAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastAppliedERO\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"finalizeBlock\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_forkNumber\",\"type\":\"uint256\"},{\"name\":\"_blockNumber\",\"type\":\"uint256\"}],\"name\":\"revertBlock\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"currentEpoch\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"lastFinalizedBlock\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"development\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"firstFilledORBEpochNumber\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"COST_ERU\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"REQUEST_GAS\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_statesRoot\",\"type\":\"bytes32\"},{\"name\":\"_transactionsRoot\",\"type\":\"bytes32\"},{\"name\":\"_intermediateStatesRoot\",\"type\":\"bytes32\"}],\"name\":\"submitNRB\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAX_REQUESTS\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"COST_NRB\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_statesRoot\",\"type\":\"bytes32\"},{\"name\":\"_transactionsRoot\",\"type\":\"bytes32\"},{\"name\":\"_intermediateStatesRoot\",\"type\":\"bytes32\"}],\"name\":\"submitURB\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CP_WITHHOLDING\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"COST_ORB\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"EROs\",\"outputs\":[{\"name\":\"timestamp\",\"type\":\"uint64\"},{\"name\":\"isExit\",\"type\":\"bool\"},{\"name\":\"isTransfer\",\"type\":\"bool\"},{\"name\":\"finalized\",\"type\":\"bool\"},{\"name\":\"challenged\",\"type\":\"bool\"},{\"name\":\"value\",\"type\":\"uint128\"},{\"name\":\"requestor\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"trieKey\",\"type\":\"bytes32\"},{\"name\":\"trieValue\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getNumEROs\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"URBs\",\"outputs\":[{\"name\":\"submitted\",\"type\":\"bool\"},{\"name\":\"epochNumber\",\"type\":\"uint64\"},{\"name\":\"requestStart\",\"type\":\"uint64\"},{\"name\":\"requestEnd\",\"type\":\"uint64\"},{\"name\":\"trie\",\"type\":\"address\"},{\"name\":\"transactionsRoot\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"state\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PREPARE_TIMEOUT\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_rootchain\",\"type\":\"address\"},{\"name\":\"_childchain\",\"type\":\"address\"}],\"name\":\"mapRequestableContractByOperator\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_isTransfer\",\"type\":\"bool\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_trieKey\",\"type\":\"bytes32\"},{\"name\":\"_trieValue\",\"type\":\"bytes32\"}],\"name\":\"startExit\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_forkNumber\",\"type\":\"uint256\"}],\"name\":\"forked\",\"outputs\":[{\"name\":\"result\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_requestId\",\"type\":\"uint256\"}],\"name\":\"getEROBytes\",\"outputs\":[{\"name\":\"out\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"COST_ERO\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"requestableContracts\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"NULL_ADDRESS\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_statesRoot\",\"type\":\"bytes32\"},{\"name\":\"_transactionsRoot\",\"type\":\"bytes32\"},{\"name\":\"_intermediateStatesRoot\",\"type\":\"bytes32\"}],\"name\":\"submitORB\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"prepareToSubmitURB\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getNumORBs\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"ORBs\",\"outputs\":[{\"name\":\"submitted\",\"type\":\"bool\"},{\"name\":\"epochNumber\",\"type\":\"uint64\"},{\"name\":\"requestStart\",\"type\":\"uint64\"},{\"name\":\"requestEnd\",\"type\":\"uint64\"},{\"name\":\"trie\",\"type\":\"address\"},{\"name\":\"transactionsRoot\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"highestBlockNumber\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_requestId\",\"type\":\"uint256\"},{\"name\":\"_userActivated\",\"type\":\"bool\"}],\"name\":\"getRequestFinalized\",\"outputs\":[{\"name\":\"finalized\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"ERUs\",\"outputs\":[{\"name\":\"timestamp\",\"type\":\"uint64\"},{\"name\":\"isExit\",\"type\":\"bool\"},{\"name\":\"isTransfer\",\"type\":\"bool\"},{\"name\":\"finalized\",\"type\":\"bool\"},{\"name\":\"challenged\",\"type\":\"bool\"},{\"name\":\"value\",\"type\":\"uint128\"},{\"name\":\"requestor\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"trieKey\",\"type\":\"bytes32\"},{\"name\":\"trieValue\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"blocks\",\"outputs\":[{\"name\":\"forkNumber\",\"type\":\"uint64\"},{\"name\":\"epochNumber\",\"type\":\"uint64\"},{\"name\":\"requestBlockId\",\"type\":\"uint64\"},{\"name\":\"timestamp\",\"type\":\"uint64\"},{\"name\":\"statesRoot\",\"type\":\"bytes32\"},{\"name\":\"transactionsRoot\",\"type\":\"bytes32\"},{\"name\":\"intermediateStatesRoot\",\"type\":\"bytes32\"},{\"name\":\"isRequest\",\"type\":\"bool\"},{\"name\":\"userActivated\",\"type\":\"bool\"},{\"name\":\"challenged\",\"type\":\"bool\"},{\"name\":\"challenging\",\"type\":\"bool\"},{\"name\":\"finalized\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastAppliedBlockNumber\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_development\",\"type\":\"bool\"},{\"name\":\"_NRBEpochLength\",\"type\":\"uint256\"},{\"name\":\"_statesRoot\",\"type\":\"bytes32\"},{\"name\":\"_transactionsRoot\",\"type\":\"bytes32\"},{\"name\":\"_intermediateStatesRoot\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"userActivated\",\"type\":\"bool\"}],\"name\":\"SessionTimeout\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"state\",\"type\":\"uint8\"}],\"name\":\"StateChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newFork\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"forkedBlockNumber\",\"type\":\"uint256\"}],\"name\":\"Forked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"epochNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"startBlockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"endBlockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"requestStart\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"requestEnd\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"epochIsEmpty\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"isRequest\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"userActivated\",\"type\":\"bool\"}],\"name\":\"EpochPrepared\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"fork\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"NRBSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"fork\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"ORBSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"fork\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"URBSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"requestor\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"weiAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"trieKey\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"trieValue\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"isTransfer\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"isExit\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"userActivated\",\"type\":\"bool\"}],\"name\":\"RequestCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"requestor\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"trieKey\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"trieValue\",\"type\":\"bytes32\"}],\"name\":\"ERUCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"forkNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"BlockFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"forkNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"epochNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"firstBlockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"lastBlockNumber\",\"type\":\"uint256\"}],\"name\":\"EpochFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"userActivated\",\"type\":\"bool\"}],\"name\":\"RequestFinalized\",\"type\":\"event\"}]"

// RootChainBin is the compiled bytecode used for deploying new contracts.
const RootChainBin = `0x60806040526000805460ff191660011790553480156200001e57600080fd5b5060405160a0806200644c83398101604090815281516020808401518385015160608601516080909601516000805460ff19168615151761010060a860020a031916610100338102919091178255600985905560018054835260078752888320838052875297822097880184905560028801899055600388018390557f3b543ac468519aa6396aba1a2611925510398aac61d1cc7b45f433e69042eb588054600160c060020a03167801000000000000000000000000000000000000000000000000426001604060020a0316021790557f3b543ac468519aa6396aba1a2611925510398aac61d1cc7b45f433e69042eb5980547f5eff886ea0ce6ca488a3d6e336d6c0f75f46d19b42c06ce5ee98e42c96d256c790975264010000000061ffff1990971690911764ff00000000191686179055949692959194929390926200016c918391906200018b810204565b6200017f640100000000620001f8810204565b50505050505062000446565b60048201805464ff000000001916640100000000179055600180546000908152600660209081526040918290208490559154815190815291820183905280517ffb96205e4b3633fd57aa805b26b51ecf528714a10241a4af015929dce86768d99281900390910190a15050565b600480546001908101918290558054600090815260086020908152604080832085845290915290209181146200027957506001805460009081526008602090815260408083206004546000190184529091529020546001604060020a0378010000000000000000000000000000000000000000000000009091048116909101165b60028201805461010061ff00199091161790558154608060020a60c060020a0319167001000000000000000000000000000000006001604060020a038381169190910291909117808455600954600160c060020a0391821678010000000000000000000000000000000000000000000000009185016000190184168202178555600185018054909216429093160291909117905560008054819060a860020a60ff02191675010000000000000000000000000000000000000000008202179055506000546040517f551dc40198cc79684bb69e4931dba4ac16e4598792ee1c0a5000aeea366d7bb6917501000000000000000000000000000000000000000000900460ff1690808260028111156200038d57fe5b60ff16815260200191505060405180910390a16004548254604080519283526001604060020a037001000000000000000000000000000000008304811660208501527801000000000000000000000000000000000000000000000000909204909116828201526000606083018190526080830181905260a0830181905260c0830181905260e0830152517f99d777cc3a347ecdaa852c6438ff2afa73f4f68b801c1e81050e817fdb6171d2918190036101000190a15050565b615ff680620004566000396000f30060806040526004361061025b5763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663033cfbed811461026057806308c4fff0146102875780630a88bd041461029c578063164bc2ae14610338578063183d2d1c1461034d578063192adc5b146103625780631f261d59146103775780632e7ab9481461038c5780633cfb871e146103a457806340a029a1146103d7578063570ca73514610403578063584e349e1461043457806361d29e83146104495780636299fb241461045e57806365d724bc146104a057806375395a58146104b557806376024792146104ca57806376671808146104e557806376e61c36146104fa5780637b929c271461051257806383bcbcff146105275780638b5172d01461053f5780638eb288ca1461055457806393248d0214610569578063935212221461057a57806394be3aa51461026057806397be455d1461058f578063b17fa6e9146105a0578063b2ae9ba814610260578063b443f3cc146105b5578063b540adba1461063b578063c0e8606414610650578063c19d93fb146106b4578063c2bc88fa146106ed578063cb5d742f14610702578063cc9374b214610729578063ce8a2bc214610748578063d1723a9614610760578063d691acd814610260578063da0185f8146107ed578063de0ce17d1461080e578063e67123c414610823578063e6925d0814610834578063ea0c73f614610849578063ea7f22a81461085e578063f1f3c46c14610876578063f28a7afa1461088e578063f4f31de4146108ab578063f4f911db146108c3578063fb788a2714610955575b600080fd5b34801561026c57600080fd5b5061027561096a565b60408051918252519081900360200190f35b34801561029357600080fd5b50610275610976565b3480156102a857600080fd5b506102b760043560243561097b565b604080516001604060020a039e8f1681529c8e1660208e01529a8d168c8c0152988c1660608c0152968b1660808b0152948a1660a08a015292891660c0890152971660e08701529515156101008601529415156101208501529315156101408401529215156101608301529115156101808201529051908190036101a00190f35b34801561034457600080fd5b50610275610a14565b34801561035957600080fd5b50610275610a1a565b34801561036e57600080fd5b50610275610a20565b34801561038357600080fd5b50610275610a2c565b34801561039857600080fd5b50610275600435610a32565b6103c36004351515600160a060020a0360243516604435606435610a44565b604080519115158252519081900360200190f35b3480156103e357600080fd5b506103c36004351515600160a060020a0360243516604435606435610ae4565b34801561040f57600080fd5b50610418610b9a565b60408051600160a060020a039092168252519081900360200190f35b34801561044057600080fd5b50610275610bae565b34801561045557600080fd5b506103c3610bb4565b34801561046a57600080fd5b5061049e60048035906024803580820192908101359160443580820192908101359160643591608435918201910135611174565b005b3480156104ac57600080fd5b506102756113d9565b3480156104c157600080fd5b506103c36113df565b3480156104d657600080fd5b5061049e6004356024356113fb565b3480156104f157600080fd5b506102756113ff565b34801561050657600080fd5b50610275600435611405565b34801561051e57600080fd5b506103c3611417565b34801561053357600080fd5b50610275600435611420565b34801561054b57600080fd5b50610275611432565b34801561056057600080fd5b5061027561143e565b6103c3600435602435604435611445565b34801561058657600080fd5b50610275611561565b6103c3600435602435604435611567565b3480156105ac57600080fd5b5061027561177f565b3480156105c157600080fd5b506105cd600435611784565b604080516001604060020a03909b168b5298151560208b015296151589890152941515606089015292151560808801526001608060020a0390911660a0870152600160a060020a0390811660c08701521660e085015261010084015261012083015251908190036101400190f35b34801561064757600080fd5b5061027561182c565b34801561065c57600080fd5b50610668600435611832565b6040805196151587526001604060020a03958616602088015293851686850152919093166060850152600160a060020a03909216608084015260a0830191909152519081900360c00190f35b3480156106c057600080fd5b506106c96118ab565b604051808260028111156106d957fe5b60ff16815260200191505060405180910390f35b3480156106f957600080fd5b506102756118bb565b34801561070e57600080fd5b506103c3600160a060020a03600435811690602435166118c1565b6103c36004351515600160a060020a0360243516604435606435611965565b34801561075457600080fd5b506103c3600435611a1c565b34801561076c57600080fd5b50610778600435611a28565b6040805160208082528351818301528351919283929083019185019080838360005b838110156107b257818101518382015260200161079a565b50505050905090810190601f1680156107df5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156107f957600080fd5b50610418600160a060020a0360043516611b53565b34801561081a57600080fd5b50610418611b6e565b6103c3600435602435604435611b73565b34801561084057600080fd5b506103c3611cfe565b34801561085557600080fd5b50610275611d89565b34801561086a57600080fd5b50610668600435611d8f565b34801561088257600080fd5b50610275600435611d9d565b34801561089a57600080fd5b506103c36004356024351515611daf565b3480156108b757600080fd5b506105cd600435611e01565b3480156108cf57600080fd5b506108de600435602435611e0f565b604080516001604060020a039d8e1681529b8d1660208d0152998c168b8b015297909a1660608a0152608089019590955260a088019390935260c0870191909152151560e0860152151561010085015215156101208401529215156101408301529115156101608201529051908190036101800190f35b34801561096157600080fd5b50610275611e9b565b67016345785d8a000081565b600181565b60086020908152600092835260408084209091529082529020805460018201546002909201546001604060020a0380831693604060020a808504831694608060020a80820485169560c060020a92839004861695858116959485048116949283048116939092049091169060ff80821691610100810482169162010000820481169163010000008104821691640100000000909104168d565b600e5481565b60015481565b670c7d713b49da000081565b60115481565b60026020526000908152604090205481565b6000806000610a5c600a600c89898989600080611ea1565b60408051838152336020820152600160a060020a038a1681830152606081018390526080810189905260a081018890528a151560c0820152600060e0820181905261010082015290519294509092507f9d57b50c5371c1c3fc64a8947cec60dbae09432e1e5d9ef048317ad7240353e391908190036101200190a15060019695505050505050565b600080806702c68af0bb14000034811115610afe57600080fd5b610b11600b600d8a8a8a8a600180611ea1565b60408051838152336020820152600160a060020a038b168183015260608101839052608081018a905260a081018990528b151560c0820152600160e0820181905261010082015290519295509093507f9d57b50c5371c1c3fc64a8947cec60dbae09432e1e5d9ef048317ad7240353e391908190036101200190a1506001979650505050505050565b6000546101009004600160a060020a031681565b60095481565b600e54600f5460008281526005602052604081205490929183918290819081908190819081908190881115610be857600080fd5b60008a81526007602090815260408083208b8452825280832080548e855260088452828520604060020a9091046001604060020a0316808652935292206004830154919b5091975090955060ff161515610cec575b60008a81526008602090815260408083208c845290915290206002015462010000900460ff161515610ca45760008a81526008602090815260408083208c8452909152902060020154610100900460ff161515610c9957600080fd5b886001019850610c3d565b935460008a81526008602090815260408083208c845282528083208d845260078352818420608060020a9095046001604060020a031680855294909252909120919850909550935b6004860154640100000000900460ff161515610d0757600080fd5b6004860154610100900460ff1615610f4257601154600b549097508710610d2d57600080fd5b600b805488908110610d3b57fe5b90600052602060002090600502019350600d8660000160109054906101000a90046001604060020a03166001604060020a0316815481101515610d7a57fe5b6000918252602090912060039091020180549093507101000000000000000000000000000000000090046001604060020a0316871415610dfb57600185015460006001604060020a03909116118015610de6575060018501546000196001604060020a03918216011688145b15610df35760018a01600e555b60018801600f555b600187016011558354604060020a900460ff1615610edf57604080516101408101825285546001604060020a0381168252604060020a810460ff908116151560208401526901000000000000000000820481161515938301939093526a010000000000000000000081048316151560608301526b010000000000000000000000810490921615156080820152606060020a9091046001608060020a031660a08201526001850154600160a060020a0390811660c083015260028601541660e082015260038501546101008201526004850154610120820152610edd9088612441565b505b83546aff0000000000000000000019166a0100000000000000000000178455604080518881526001602082015281517f134017cf3262b18f892ee95dde3b0aec9a80cc70a7c96f09c64bd237aceb0473929181900390910190a160019a50611167565b601054600a549097508710610f5657600080fd5b600a805488908110610f6457fe5b90600052602060002090600502019150600c8660000160109054906101000a90046001604060020a03166001604060020a0316815481101515610fa357fe5b6000918252602090912060039091020180549091507101000000000000000000000000000000000090046001604060020a031687141561102457600185015460006001604060020a0390911611801561100f575060018501546000196001604060020a03918216011688145b1561101c5760018a01600e555b60018801600f555b600187016010558154604060020a900460ff161561110857604080516101408101825283546001604060020a0381168252604060020a810460ff908116151560208401526901000000000000000000820481161515938301939093526a010000000000000000000081048316151560608301526b010000000000000000000000810490921615156080820152606060020a9091046001608060020a031660a08201526001830154600160a060020a0390811660c083015260028401541660e0820152600383015461010082015260048301546101208201526111069088612441565b505b81546aff0000000000000000000019166a0100000000000000000000178255604080518881526000602082015281517f134017cf3262b18f892ee95dde3b0aec9a80cc70a7c96f09c64bd237aceb0473929181900390910190a160019a505b5050505050505050505090565b60008061117f61487e565b60015460009081526007602090815260408083208e84529091529020600481015490935060ff16156111b057600080fd5b8254426001604060020a0360c060020a90920491909116600101116111d457600080fd5b6004830154610100900460ff161561122b578254600d80549091608060020a90046001604060020a031690811061120757fe5b6000918252602090912060016003909202010154600160a060020a0316915061126c565b8254600c80549091608060020a90046001604060020a031690811061124c57fe5b6000918252602090912060016003909202010154600160a060020a031691505b6112a588888080601f0160208091040260200160405190810160405280939291908181526020018383808284375061256c945050505050565b90506112b0816126b5565b15156112bb57600080fd5b81600160a060020a031663f7e498f684600201548c8c8c8c8c8c8c6040518963ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180896000191660001916815260200180602001806020018681526020018060200184810384528b8b8281815260200192508082843790910185810384528981526020019050898980828437909101858103835286815260209081019150879087028082843782019150509b505050505050505050505050602060405180830381600087803b15801561139557600080fd5b505af11580156113a9573d6000803e3d6000fd5b505050506040513d60208110156113bf57600080fd5b505115156113cc57600080fd5b5050505050505050505050565b60105481565b60006113e96126de565b15156113f457600080fd5b5060015b90565b5050565b60045481565b60066020526000908152604090205481565b60005460ff1681565b60036020526000908152604090205481565b6702c68af0bb14000081565b620186a081565b60008054819081906101009004600160a060020a0316331461146657600080fd5b60008060005460a860020a900460ff16600281111561148157fe5b1461148b57600080fd5b67016345785d8a0000348111156114a157600080fd5b6114a96126de565b50600154600090815260086020908152604080832060045484529091529020600281015490945062010000900460ff16156114e357600080fd5b6114f38888886000806000612844565b600154604080519182526020820183905280519295507fd85b68c66f9db94728c533f39125381246238815483496c57a62b56f70dea8f192918290030190a1835460c060020a90046001604060020a031683141561155357611553612a9f565b506001979650505050505050565b6103e881565b6000808080808060028060005460a860020a900460ff16600281111561158957fe5b1461159357600080fd5b670c7d713b49da0000348111156115a957600080fd5b6001805460009081526007602090815260408083206005835281842054845290915290206004015460ff161597506115e8908c908c908c90808c612844565b9550851561176c57600180546000908152600860209081526040808320600454845282529182902082516101a08101845281546001604060020a038082168352604060020a808304821695840195909552608060020a80830482169684019690965260c060020a9182900481166060840152958301548087166080840152938404861660a0830152938304851660c08201529290910490921660e0820152600282015460ff8082161515610100808501919091528204811615156101208401526201000082048116151561014084015263010000008204811615156101608401526401000000009091041615156101808201529095506116e790612f15565b855460018054600090815260056020526040902054929650608060020a9091046001604060020a03169091030192508284141561172657611726612f48565b600154604080519182526020820188905280517f041d165424b249ff6b188265c7ea6f322c4f47171efbac0768888b5fa3ac13d59281900390910190a160019750611771565b600097505b505050505050509392505050565b600381565b600a80548290811061179257fe5b6000918252602090912060059091020180546001820154600283015460038401546004909401546001604060020a0384169550604060020a840460ff90811695690100000000000000000086048216956a010000000000000000000081048316956b010000000000000000000000820490931694606060020a9091046001608060020a031693600160a060020a039384169390911691908a565b600a5490565b600d80548290811061184057fe5b600091825260209091206003909102018054600182015460029092015460ff821693506001604060020a0361010083048116936901000000000000000000840482169371010000000000000000000000000000000000900490911691600160a060020a039091169086565b60005460a860020a900460ff1681565b610e1081565b600080546101009004600160a060020a031633146118de57600080fd5b6118f083600160a060020a0316613134565b15156118fb57600080fd5b600160a060020a03838116600090815260126020526040902054161561192057600080fd5b50600160a060020a038281166000908152601260205260409020805473ffffffffffffffffffffffffffffffffffffffff191691831691909117905560015b92915050565b6000808067016345785d8a00003481111561197f57600080fd5b611993600a600c8a8a8a8a60016000611ea1565b60408051838152336020820152600160a060020a038b168183015260608101839052608081018a905260a081018990528b151560c0820152600160e0820152600061010082015290519295509093507f9d57b50c5371c1c3fc64a8947cec60dbae09432e1e5d9ef048317ad7240353e391908190036101200190a1506001979650505050505050565b6001548114155b919050565b60606000600a83815481101515611a3b57fe5b600091825260208083206005929092029091016002810154600160a060020a039081168085526012845260408086205481516101408101835285546001604060020a0381168252604060020a810460ff9081161515988301989098526901000000000000000000810488161515938201939093526a010000000000000000000083048716151560608201526b010000000000000000000000830490961615156080870152606060020a9091046001608060020a031660a08601526001840154831660c086015260e085019190915260038301546101008501526004830154610120850152919450611b4a93611b459388939192611b38921661313c565b919063ffffffff6131cd16565b61322c565b91505b50919050565b601260205260009081526040902054600160a060020a031681565b600081565b600080548190819081906101009004600160a060020a03163314611b9657600080fd5b60018060005460a860020a900460ff166002811115611bb157fe5b14611bbb57600080fd5b67016345785d8a000034811115611bd157600080fd5b611bd96126de565b50600154600090815260086020908152604080832060045484529091529020600281015490955062010000900460ff161515611c1457600080fd5b611c248989896001600080612844565b600154604080519182526020820183905280519296507f041d165424b249ff6b188265c7ea6f322c4f47171efbac0768888b5fa3ac13d592918290030190a160005460ff161515611cce576001546000908152600760209081526040808320878452909152902054600c80549091608060020a90046001604060020a0316908110611cab57fe5b6000918252602090912060039091020160028101549093508814611cce57600080fd5b845460c060020a90046001604060020a0316841415611cef57611cef612f48565b50600198975050505050505050565b600080546101009004600160a060020a03163314611d1b57600080fd5b60028060005460a860020a900460ff166002811115611d3657fe5b1415611d4157600080fd5b6000805475ff00000000000000000000000000000000000000000019167502000000000000000000000000000000000000000000179055611d806133e2565b600191505b5090565b600c5490565b600c80548290811061184057fe5b60056020526000908152604090205481565b60008115611dca57600b805484908110611dc557fe5b506000525b600a805484908110611dd857fe5b60009182526020909120600590910201546a0100000000000000000000900460ff169392505050565b600b80548290811061179257fe5b6007602090815260009283526040808420909152908252902080546001820154600283015460038401546004909401546001604060020a0380851695604060020a8604821695608060020a810483169560c060020a90910490921693919260ff80821691610100810482169162010000820481169163010000008104821691640100000000909104168c565b600f5481565b60008060008060008615611eee5785611ed157611ecc3467016345785d8a000063ffffffff6135b416565b611ee9565b611ee9346702c68af0bb14000063ffffffff6135b416565b611ef0565b345b93508a80611f175750600160a060020a038a81166000908152601260205260409020541615155b1515611f1f57fe5b8a15611f47578315801590611f32575088155b8015611f3c575087155b1515611f4757600080fd5b8c54611f568e600183016148e6565b94508c85815481101515611f6657fe5b600091825260209091206005909102016001810180543373ffffffffffffffffffffffffffffffffffffffff19918216179091558154600283018054909216600160a060020a038e1617909155600382018b9055600482018a90557fffffffff00000000000000000000000000000000ffffffffffffffffffffffff16606060020a6001608060020a038716021767ffffffffffffffff1916426001604060020a03161768ff00000000000000001916604060020a891580159182029290921769ff000000000000000000191669010000000000000000008f1515021783559194509061205157508a155b1561212857604080516101408101825284546001604060020a0381168252604060020a810460ff908116151560208401526901000000000000000000820481161515938301939093526a010000000000000000000081048316151560608301526b010000000000000000000000810490921615156080820152606060020a9091046001608060020a031660a08201526001840154600160a060020a0390811660c083015260028501541660e0820152600384015461010082015260048401546101208201526121209086612441565b151561212857fe5b8b541515612149578b5461213f8d60018301614912565b5060009150612152565b8b546000190191505b8b8281548110151561216057fe5b60009182526020909120600390910201805490915060ff16806121bc5750805460016001604060020a036901000000000000000000830481167101000000000000000000000000000000000090930481169290920301166103e8145b1561221e578b548c906121d28260018301614912565b815481106121dc57fe5b60009182526020909120600390910201805470ffffffffffffffff000000000000000000191669010000000000000000006001604060020a0388160217815590505b612227816135cb565b805478ffffffffffffffff00000000000000000000000000000000001916710100000000000000000000000000000000006001604060020a038716021781558a1561234d57604080516101408101825284546001604060020a0381168252604060020a810460ff908116151560208401526901000000000000000000820481161515938301939093526a010000000000000000000081048316151560608301526b010000000000000000000000810490921615156080820152606060020a9091046001608060020a031660a08201526001840154600160a060020a0390811660c083015260028501541660e0820152600384015461010082015260048401546101208201526123489061233a908c61313c565b82908763ffffffff61363716565b612431565b600160a060020a038a81166000908152601260209081526040918290205482516101408101845287546001604060020a0381168252604060020a810460ff9081161515948301949094526901000000000000000000810484161515948201949094526a010000000000000000000084048316151560608201526b010000000000000000000000840490921615156080830152606060020a9092046001608060020a031660a08201526001860154831660c08201526002860154831660e0820152600386015461010082015260048601546101208201526124319261233a921661313c565b5050509850989650505050505050565b60008260400151156124a1578260e00151600160a060020a03166108fc8460a001516001608060020a03169081150290604051600060405180830381858888f19350505050158015612497573d6000803e3d6000fd5b506001905061195f565b60e083015160208085015160c0860151610100870151610120880151604080517fd9afd3a9000000000000000000000000000000000000000000000000000000008152941515600486015260248501899052600160a060020a039384166044860152606485019290925260848401525193169263d9afd3a99260a4808401939192918290030181600087803b15801561253957600080fd5b505af115801561254d573d6000803e3d6000fd5b505050506040513d602081101561256357600080fd5b50519392505050565b61257461487e565b606061259060096125848561387f565b9063ffffffff6138ad16565b90506125b38160008151811015156125a457fe5b9060200190602002015161393b565b6001604060020a0316825280516125d190829060019081106125a457fe5b602083015280516125e990829060029081106125a457fe5b6001604060020a031660408301528051612619908290600390811061260a57fe5b9060200190602002015161395f565b600160a060020a03166060830152805161263a90829060049081106125a457fe5b60808301528051612661908290600590811061265257fe5b9060200190602002015161397b565b60a0830152805161267990829060069081106125a457fe5b60c0830152805161269190829060079081106125a457fe5b60e083015280516126a990829060089081106125a457fe5b61010083015250919050565b600060438260c001511480156126cd575060e0820151155b801561195f57505061010001511590565b6000808080600260005460a860020a900460ff1660028111156126fd57fe5b141561270c576000935061283e565b600180546000908152600660209081526040808320546005909252909120549101935083111561273f576000935061283e565b6001546000908152600760209081526040808320868452909152902060048101549092506301000000900460ff161561277b576000935061283e565b600482015460ff16156127c4578154426001604060020a0360c060020a9092049190911660010111156127b1576000935061283e565b6127bb82846139cf565b6001935061283e565b50805460016001604060020a03604060020a909204821601166127e681613a3c565b156128065781546127bb90604060020a90046001604060020a0316613b88565b8154426001604060020a0360c060020a90920491909116600301111561282f576000935061283e565b61283982846139cf565b600193505b50505090565b60008060008060008780156128565750865b801561285f5750855b156129be57600180546128779163ffffffff613ced16565b6000818152600260208181526040808420546008835281852081865290925290922090810154929650909450925062010000900460ff1680156128d557506001820154426001604060020a0360c060020a90920491909116610e1001105b1561291c5760008481526008602090815260408083208684528252808320838155600181018490556002908101805464ffffffffff19169055878452909152812055612a91565b60018481556000858152600860208181526040808420600283528185205485528252808420546000198a0185529282528084208885528252928390209093018054608060020a9092046001604060020a031667ffffffffffffffff1990921682179055815187815292830181905281519097507f18013fce596c7fc273e36aaa4d9ba6f94db4e483239db76e94fe9eb6769df2a89281900390910190a16129e3565b600180546000908152600560205260409020546129e09163ffffffff613ced16565b94505b5060018054600090815260076020908152604080832088845282528083208085018f9055600281018e9055600381018d905580546001604060020a0342811660c060020a02600160c060020a039092169190911780835560048054909216604060020a026fffffffffffffffff000000000000000019909116178255810180548c15156101000261ff00198f151560ff19909316929092179190911617905593548352600590915290208590555b505050509695505050505050565b60048054600181810192839055805460009081526008602090815260408083209583528582528083209483529490529283205482546001604060020a0360c060020a92839004811684018116608060020a0277ffffffffffffffff000000000000000000000000000000001990921691909117845560028401805461ff001962ff000019909116620100001716610100179055918301805442909316909102600160c060020a039092169190911790559080612b5a83613d06565b600283015460ff1615612bde578254600160c060020a0377ffffffffffffffff00000000000000000000000000000000196fffffffffffffffff00000000000000001983166001604060020a03938416604060020a0217908116608060020a918290048416600019018416820217918216910490911660c060020a02178355612c62565b600a5483546fffffffffffffffff00000000000000001916604060020a6000199092016001604060020a03908116830291909117808655600192612c3192828116919092048216038301166103e8614137565b84546001604060020a03608060020a8204811692909201929092031660c060020a02600160c060020a039091161783555b600080546001919075ff000000000000000000000000000000000000000000191660a860020a8302179055506000546040517f551dc40198cc79684bb69e4931dba4ac16e4598792ee1c0a5000aeea366d7bb69160a860020a900460ff169080826002811115612cce57fe5b60ff16815260200191505060405180910390a160045483546002850154604080519384526001604060020a03608060020a84048116602086015260c060020a84048116858301528084166060860152604060020a909304909216608084015260ff16151560a0830152600160c0830152600060e0830152517f99d777cc3a347ecdaa852c6438ff2afa73f4f68b801c1e81050e817fdb6171d2918190036101000190a1600283015460ff1615612d8b57612d86612f48565b612f10565b604080516101a08101825284546001604060020a038082168352604060020a80830482166020850152608060020a80840483169585019590955260c060020a928390048216606085015260018801548083166080860152908104821660a0850152938404811660c0840152920490911660e0820152600284015460ff808216151561010080850191909152820481161515610120840152620100008204811615156101408401526301000000820481161515610160840152640100000000909104161515610180820152612e5e90612f15565b9150600090505b81816001604060020a03161015612f10576001805460009081526007602081815260408084208854608060020a908190046001604060020a03908116890181168752918452828620600401805460ff19168817905589870154875487529484528286208a5482900483168901831687529093529320805477ffffffffffffffff000000000000000000000000000000001916604060020a909304841686019093160217905501612e65565b505050565b600081610100015115612f2a57506000611a23565b81604001518260600151036001016001604060020a03169050919050565b60048054600190810191829055805460009081526008602090815260408083208584529091529020918114612fb357506001805460009081526008602090815260408083206004546000190184529091529020546001604060020a0360c060020a9091048116909101165b60028201805461010061ff0019909116179055815477ffffffffffffffff000000000000000000000000000000001916608060020a6001604060020a038381169190910291909117808455600954600160c060020a0391821660c060020a9185016000190184168202178555600185018054909216429093160291909117905560008054819075ff000000000000000000000000000000000000000000191660a860020a8202179055506000546040517f551dc40198cc79684bb69e4931dba4ac16e4598792ee1c0a5000aeea366d7bb69160a860020a900460ff16908082600281111561309d57fe5b60ff16815260200191505060405180910390a16004548254604080519283526001604060020a03608060020a83048116602085015260c060020a909204909116828201526000606083018190526080830181905260a0830181905260c0830181905260e0830152517f99d777cc3a347ecdaa852c6438ff2afa73f4f68b801c1e81050e817fdb6171d2918190036101000190a15050565b6000903b1190565b61314461493e565b6020808401511515908201526040808401511580159183019190915260c080850151600160a060020a03169083015260a0808501516001608060020a031690830152610100808501519083015261012080850151908301526131b85760e080840151600160a060020a03169082015261195f565b600160a060020a03821660e082015292915050565b6131d561487e565b633b9aca006020820152620186a0604082015260e0840151600160a060020a0316606082015260a08401516001608060020a03166080820152613219848484614171565b60a0820152604360c08201529392505050565b6040805160098082526101408201909252606091829190816020015b60608152602001906001900390816132485750508351909150613273906001604060020a0316614289565b81600081518110151561328257fe5b9060200190602002018190525061329c8360200151614289565b8160018151811015156132ab57fe5b6020908102909101015260408301516132cc906001604060020a0316614289565b8160028151811015156132db57fe5b6020908102909101015260608301516132fc90600160a060020a031661429c565b81600381518110151561330b57fe5b60209081029091010152608083015161332390614289565b81600481518110151561333257fe5b6020908102909101015260a083015161334a906142cc565b81600581518110151561335957fe5b6020908102909101015260c083015161337190614289565b81600681518110151561338057fe5b6020908102909101015260e083015161339890614289565b8160078151811015156133a757fe5b602090810290910101526101008301516133c090614289565b8160088151811015156133cf57fe5b60209081029091010152611b4a8161434f565b60018054600081815260066020908152604080832054600783528184208185528352818420805495870180865260028552838620604060020a9097046001604060020a03169687905560088552838620878752909452919093209294909391929083141561345e57805467ffffffffffffffff191681556134ac565b6001805460009081526008602090815260408083206002835281842054845290915290205482546001604060020a03604060020a90920482169092011667ffffffffffffffff199091161781555b60028101805463ff0000001962ff0000199091166201000017166301000000179055600181810180546001604060020a03421660c060020a02600160c060020a03909116179055600b546135059163ffffffff6135b416565b81546fffffffffffffffff00000000000000001916604060020a6001604060020a0392831681029190911777ffffffffffffffff000000000000000000000000000000001916608060020a600189810185169190910291909117808555909261357e9282821692048116919091038301166103e8614137565b82546001604060020a03608060020a8204811692909201929092031660c060020a02600160c060020a0390911617905550505050565b600080838311156135c457600080fd5b5050900390565b6001810154600160a060020a03161515613634576135e7614992565b604051809103906000f080158015613603573d6000803e3d6000fd5b5060018201805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03929092169190911790555b50565b60018301546000906060908190600160a060020a0316151561365857600080fd5b855461367b908590690100000000000000000090046001604060020a03166135b4565b925061368683614289565b9150613697611b45868660006131cd565b6001870154604080517f20ba5b6000000000000000000000000000000000000000000000000000000000815260048101918252855160448201528551939450600160a060020a03909216926320ba5b60928692869290918291602482019160640190602087019080838360005b8381101561371c578181015183820152602001613704565b50505050905090810190601f1680156137495780820380516001836020036101000a031916815260200191505b50838103825284518152845160209182019186019080838360005b8381101561377c578181015183820152602001613764565b50505050905090810190601f1680156137a95780820380516001836020036101000a031916815260200191505b50945050505050600060405180830381600087803b1580156137ca57600080fd5b505af11580156137de573d6000803e3d6000fd5b5050506001870154604080517f80759f1f0000000000000000000000000000000000000000000000000000000081529051600160a060020a0390921692506380759f1f9160048083019260209291908290030181600087803b15801561384357600080fd5b505af1158015613857573d6000803e3d6000fd5b505050506040513d602081101561386d57600080fd5b50516002909601959095555050505050565b6138876149a2565b5080516040805180820190915260208084018083529082018390529091905b5050919050565b60606138b76149b9565b6000836040519080825280602002602001820160405280156138f357816020015b6138e06149a2565b8152602001906001900390816138d85790505b5092506138ff856143b1565b91505b8381101561393357613913826143d6565b838281518110151561392157fe5b60209081029091010152600101613902565b505092915050565b600080600061394984614408565b90516020919091036101000a9004949350505050565b60008061396b83614408565b5051606060020a90049392505050565b602081015160609080151561398f57611b4d565b806040519080825280601f01601f1916602001820160405280156139bd578160200160208202803883390190505b509150611b4d8360000151838361446b565b60048201805464ff000000001916640100000000179055600180546000908152600660209081526040918290208490559154815190815291820183905280517ffb96205e4b3633fd57aa805b26b51ecf528714a10241a4af015929dce86768d99281900390910190a15050565b6000806000600454841115613a5457600092506138a6565b60015460009081526008602090815260408083208784529091529020600281015490925062010000900460ff161515613a9057600092506138a6565b600282015460ff1615613aca574260018360010160189054906101000a90046001604060020a03166001604060020a0316011192506138a6565b6001546000908152600560205260409020548254608060020a90046001604060020a03161115613afd57600092506138a6565b5060015460009081526007602090815260408083208454608060020a90046001604060020a0316845290915290206004810154640100000000900460ff1615613b4957600192506138a6565b60048101546301000000900460ff1615613b6657600092506138a6565b5442600160c060020a9092046001604060020a03169190910111159392505050565b6001546000908152600860209081526040808320848452909152812080549091608060020a9091046001604060020a03169080805b845460c060020a90046001604060020a03168411613c49576001546000908152600760209081526040808320878452909152902060048101549092506301000000900460ff1680613c185750600482015462010000900460ff165b15613c265760019250613c49565b60048201805464ff00000000191664010000000017905560019390930192613bbd565b82613c545783613c59565b600184035b8554909150608060020a90046001604060020a03168110613ce55760018054600090815260066020908152604091829020849055915487548251918252928101899052608060020a9092046001604060020a03168282015260608201839052517f70801d4d63b3da6c19ba7349911f45bed5a99ccdfb51b8138c105872529bebd5916080908290030190a15b505050505050565b600082820183811015613cff57600080fd5b9392505050565b600a546000901515613d265760028201805460ff191660011790556113fb565b50600154600090815260086020908152604080832060045460011901845290915290208054600a54604060020a9091046001604060020a03166000199091011415613d7b5760028201805460ff191660011790555b60015460021415613d8b576113fb565b600282015460ff1615613f225780548254604060020a9091046001604060020a031667ffffffffffffffff19909116178255600281015460ff1615613e085760018181015490830180546fffffffffffffffff00000000000000001916604060020a928390046001604060020a0316909202919091179055613f1d565b604080516101a08101825282546001604060020a038082168352604060020a80830482166020850152608060020a80840483169585019590955260c060020a928390048216606085015260018601548083166080860152908104821660a0850152938404811660c0840152920490911660e0820152600282015460ff808216151561010080850191909152820481161515610120840152620100008204811615156101408401526301000000820481161515610160840152640100000000909104161515610180820152613edb90612f15565b8160010160089054906101000a90046001604060020a0316018260010160086101000a8154816001604060020a0302191690836001604060020a031602179055505b6140f3565b6001546000908152600360205260409020541515613f54576004546001546000908152600360205260409020556140f3565b600281015460ff1615613fbe578054825467ffffffffffffffff19166001604060020a03604060020a9283900481169190911784556001808401549085018054918490049092169092026fffffffffffffffff0000000000000000199092169190911790556140f3565b8054825460016001604060020a03604060020a9384900481168201811667ffffffffffffffff19909316929092178555604080516101a0810182528554808516825285810485166020830152608060020a80820486169383019390935260c060020a9081900485166060830152928601548085166080830152948504841660a0820152908404831660c082015292041660e0820152600282015460ff8082161515610100848101919091528204811615156101208401526201000082048116151561014084015263010000008204811615156101608401526401000000009091041615156101808201526140b190612f15565b8160010160089054906101000a90046001604060020a0316018260010160086101000a8154816001604060020a0302191690836001604060020a031602179055505b600c54600010156113fb57600c805460019190600019810190811061411457fe5b60009182526020909120600390910201805460ff19169115159190911790555050565b6000818381151561414457fe5b061561415e57818381151561415557fe5b04600101613cff565b818381151561416957fe5b049392505050565b6060600084604001511561418457614281565b826141af577fe904e3d9000000000000000000000000000000000000000000000000000000006141d1565b7fd9afd3a9000000000000000000000000000000000000000000000000000000005b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19169050808560200151614203576000614206565b60015b60c08701516101008801516101208901516040805177ffffffffffffffffffffffffffffffffffffffffffffffff19909616602087015260ff94909416602886015260488501899052600160a060020a039092166068850152608884015260a8808401919091528151808403909101815260c8909201905291505b509392505050565b606061195f614297836144a9565b6142cc565b60408051741400000000000000000000000000000000000000008318601482015260348101909152606090611b4a815b60608151600114801561432b575081517f7f00000000000000000000000000000000000000000000000000000000000000908390600090811061430b57fe5b90602001015160f860020a900460f860020a02600160f860020a03191611155b15614337575080611a23565b61195f6143498351608060ff166145da565b83614702565b604080516000808252602082019092526060915b83518110156143975761438d82858381518110151561437e57fe5b90602001906020020151614702565b9150600101614363565b6143a9614349835160c060ff166145da565b949350505050565b6143b96149b9565b60006143c48361480c565b83519383529092016020820152919050565b6143de6149a2565b602082015160006143ee82614852565b828452602080850182905292019390910192909252919050565b805180516000918291821a9082608083101561442a5781945060019350614463565b60b88310156144485760018660200151039350816001019450614463565b60b78303905080600187602001510303935080820160010194505b505050915091565b6020601f820104836020840160005b838110156144965760208102838101519083015260010161447a565b5060008551602001860152505050505050565b60408051602080825281830190925260609160009183918291849180820161040080388339019050509250856020840152600093505b602084101561453d5782848151811015156144f657fe5b60209101015160f860020a90819004027fff0000000000000000000000000000000000000000000000000000000000000016156145325761453d565b6001909301926144df565b836020036040519080825280601f01601f19166020018201604052801561456e578160200160208202803883390190505b509150600090505b81518110156145d157825160018501948491811061459057fe5b90602001015160f860020a900460f860020a0282828151811015156145b157fe5b906020010190600160f860020a031916908160001a905350600101614576565b50949350505050565b60608080604060020a851061465057604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600e60248201527f696e70757420746f6f206c6f6e67000000000000000000000000000000000000604482015290519081900360640190fd5b6040805160018082528183019092529060208083019080388339019050509150603785116146b05783850160f860020a0282600081518110151561469057fe5b906020010190600160f860020a031916908160001a905350819250613933565b6146b9856144a9565b90508381510160370160f860020a028260008151811015156146d757fe5b906020010190600160f860020a031916908160001a9053506146f98282614702565b95945050505050565b60608060008084518651016040519080825280601f01601f19166020018201604052801561473a578160200160208202803883390190505b50925060009150600090505b85518110156147a257858181518110151561475d57fe5b90602001015160f860020a900460f860020a02838381518110151561477e57fe5b906020010190600160f860020a031916908160001a90535060019182019101614746565b5060005b84518110156148025784818151811015156147bd57fe5b90602001015160f860020a900460f860020a0283838151811015156147de57fe5b906020010190600160f860020a031916908160001a905350600191820191016147a6565b5090949350505050565b8051805160009190821a90608082101561482957600092506138a6565b60b8821080614844575060c08210158015614844575060f882105b156138a657600192506138a6565b8051600090811a608081101561486b5760019150611b4d565b60b8811015611b4d57607e190192915050565b6101206040519081016040528060006001604060020a031681526020016000815260200160006001604060020a031681526020016000600160a060020a0316815260200160008152602001606081526020016000815260200160008152602001600081525090565b815481835581811115612f1057600502816005028360005260206000209182019101612f1091906149da565b815481835581811115612f1057600302816003028360005260206000209182019101612f109190614a51565b6040805161014081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c0810182905260e08101829052610100810182905261012081019190915290565b60405161151d80614aae83390190565b604080518082019091526000808252602082015290565b6060604051908101604052806149cd6149a2565b8152602001600081525090565b6113f891905b80821115611d855780547fffffffff0000000000000000000000000000000000000000000000000000000016815560018101805473ffffffffffffffffffffffffffffffffffffffff19908116909155600282018054909116905560006003820181905560048201556005016149e0565b6113f891905b80821115611d8557805478ffffffffffffffffffffffffffffffffffffffffffffffffff1916815560018101805473ffffffffffffffffffffffffffffffffffffffff1916905560006002820155600301614a575600608060405234801561001057600080fd5b506114fd806100206000396000f3006080604052600436106100775763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166320ba5b60811461007c57806350c946fe1461009e578063693ac4fb146100d657806380759f1f14610104578063a43914da14610126578063f7e498f61461014a575b600080fd5b34801561008857600080fd5b5061009c61009736600461123a565b610177565b005b3480156100aa57600080fd5b506100be6100b9366004611122565b61018d565b6040516100cd9392919061139d565b60405180910390f35b3480156100e257600080fd5b506100f66100f1366004611205565b610278565b6040516100cd92919061140f565b34801561011057600080fd5b5061011961052b565b6040516100cd91906113d9565b34801561013257600080fd5b5061013b610531565b6040516100cd939291906113e7565b34801561015657600080fd5b5061016a610165366004611148565b61057b565b6040516100cd91906113c5565b6101896000838363ffffffff6107dc16565b5050565b610195610f62565b61019d610f62565b6101a5610f62565b6101ad610f7d565b6000858152600460209081526040808320815160608101909252909290918391908201908390600290835b828210156102265760408051808201825260038402860180548252825180840190935260018082015484526002909101546020848101919091528083019390935290835290920191016101d8565b5050509152505080515151855280515160209081015151855281515181015181015184528151810151518682015281518101518101515185820152905181015181015181015190830152509193909250565b60006060610284610f96565b61028c610fad565b610294610fc8565b60008061029f610f96565b6102a7610f96565b60006102b1610f96565b6000805415156102c057600080fd5b60408051908101604052808e6040518082805190602001908083835b602083106102fb5780518252601f1990920191602091820191016102dc565b518151602093840361010090810a600019018019909316929091169190911790915260408051939095018390039092208752958601525080518082018252600154815281518083019092526002548252600354828601529384015250919b509950505b6020890151610374908b9063ffffffff61095216565b6020808c0151810151908301519297509095501461038e57fe5b6020840151151561039e576104a7565b60208501519690960160ff81900360020a9b909b179a600101956103c184610980565b8a5160009081526004602052604090209194509250610427906001859003600281106103e957fe5b6040805180820182526003929092029290920180548252825180840190935260018101548352600201546020838101919091528101919091526109e7565b60018701968990610100811061043957fe5b6020908102919091019190915289516000908152600490915260409020836002811061046157fe5b604080518082018252600392909202929092018054825282518084019093526001810154835260020154602083810191909152810191909152919950909750889061035e565b600086111561051c57856040519080825280602002602001820160405280156104da578160200160208202803883390190505b509a50600090505b8581101561051c57878161010081106104f757fe5b60200201518b8281518110151561050a57fe5b602090810290910101526001016104e2565b50505050505050505050915091565b60005490565b600080600061053e610fad565b5050604080518082018252600154815281518083019092526002548252600354602083810191825282018390529051915190519194909350909150565b6000610585610f96565b61058d610fad565b600080600061059a610f62565b60408051908101604052808c6040518082805190602001908083835b602083106105d55780518252601f1990920191602091820191016105b6565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405180910390206000191681526020016101008152509550896040518082805190602001908083835b602083106106435780518252601f199092019160209182019101610624565b5181516020939093036101000a600019018019909116921691909117905260405192018290039091208852506000955050505b88156107af5761068589610a84565b60ff1692508260019060020a0219891698506106ad8360ff0387610add90919063ffffffff16565b602087018190529096506106c090610980565b602087015291506106d0856109e7565b8183600281106106dc57fe5b602002015287518890858103600019019081106106f557fe5b90602001906020020151818360010360028110151561071057fe5b6020908102919091019190915281518282015160408051808501939093528281019190915280518083038201815260609092019081905281519192909182918401908083835b602083106107755780518252601f199092019160209182019101610756565b5181516020939093036101000a60001901801990911692169190911790526040519201829003909120885250505060019390930192610676565b602085018690526107bf856109e7565b8c146107ca57600080fd5b5060019b9a5050505050505050505050565b6107e4610f96565b60006107ee610fad565b6040805190810160405280866040518082805190602001908083835b602083106108295780518252601f19909201916020918201910161080a565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405180910390206000191681526020016101008152509250836040518082805190602001908083835b602083106108975780518252601f199092019160209182019101610878565b5181516020939093036101000a600019018019909116921691909117905260405192018290039091208954909550151592506108df9150505760208101839052818152610922565b60408051808201825260018801548152815180830190925260028801548252600388015460208381019190915281019190915261091f9087908585610b51565b90505b61092b816109e7565b86558051600187015560209081015180516002880155015160039095019490945550505050565b61095a610f96565b610962610f96565b610975846109708686610d8e565b610add565b915091509250929050565b600061098a610f96565b602083015160001061099b57600080fd5b50508051604080518082019091528251600202815260209283015160001901928101929092527f8000000000000000000000000000000000000000000000000000000000000000900491565b8051602080830151808201519051604080518085019590955284810192909252606080850191909152815180850390910181526080909301908190528251600093928291908401908083835b60208310610a525780518252601f199092019160209182019101610a33565b5181516020939093036101000a6000190180199091169216919091179052604051920182900390912095945050505050565b60008080831515610a9457600080fd5b5082905060805b600160ff821610610ad65760001960ff821660020a0182161515610ac9579182019160ff811660020a909104905b600260ff90911604610a9b565b5050919050565b610ae5610f96565b610aed610f96565b83602001518311158015610b0357506101008311155b1515610b0b57fe5b60208201839052821515610b225760008252610b34565b835160ff84900360020a600119021682525b60208085015184900390820152925160029290920a909102825291565b610b59610fad565b610b61610f96565b610b69610f96565b6000610b73610f96565b6000610b7d610f7d565b610b85610f7d565b6020808c0151810151908b01511015610b9a57fe5b610ba88a8c60200151610952565b602081015191985096501515610bc057889250610d6a565b6020808c01518101519088015110610cde576020860151600110610be057fe5b8a51600090815260048d0160209081526040808320815160608101909252909290918391908201908390600290835b82821015610c5d576040805180820182526003840286018054825282518084019093526001808201548452600290910154602084810191909152808301939093529083529092019101610c0f565b50505050815250509150610c7086610980565b83519196509450610c94908d908760028110610c8857fe5b6020020151868c610b51565b82518660028110610ca157fe5b602090810291909101919091528b51600090815260048e019091526040812090610ccb8282610fe9565b5050610cd78c83610e00565b9250610d6a565b610ce786610980565b604080518082019091528b8152602081018290528351929750909550908660028110610d0f57fe5b602002018190525060408051908101604052808c60000151600019168152602001610d458d602001518a60200151600101610e66565b90528151600187900360028110610d5857fe5b6020020152610d678c82610e00565b92505b50506040805180820190915290815260208101949094525091979650505050505050565b60008060008360200151856020015110610dac578360200151610db2565b84602001515b9150811515610dc45760009250610df8565b50825184511861010082900360020a60000316801515610de657819250610df8565b610def81610e9a565b60ff0360ff1692505b505092915050565b600080610e0c83610eed565b83515160008281526004968701602090815260409091208251815591810151805160018401558101516002830155945185015180516003830155850151805196820196909655949093015160059094019390935550919050565b610e6e610f96565b6020830151821115610e7f57600080fd5b60208084015183900390820152915160029190910a02815290565b60008080831515610eaa57600080fd5b5082905060805b600160ff821610610ad65760ff811660020a600019810102821615610ee0579182019160ff811660020a909104905b600260ff90911604610eb1565b8051600090610f0290825b60200201516109e7565b8251610f0f906001610ef8565b6040805160208082019490945280820192909252805180830382018152606090920190819052815191929091829184019080838360208310610a525780518252601f199092019160209182019101610a33565b60408051808201825290600290829080388339509192915050565b60c060405190810160405280610f91611013565b905290565b604080518082019091526000808252602082015290565b60408051606081019091526000815260208101610f91610f96565b61200060405190810160405280610100906020820280388339509192915050565b50600080825560018201819055600282018190556003820181905560048201819055600590910155565b60c0604051908101604052806002905b61102b610fad565b8152602001906001900390816110235790505090565b6000601f8201831361105257600080fd5b813561106561106082611456565b61142f565b9150818183526020840193506020810190508385602084028201111561108a57600080fd5b60005b838110156110b657816110a088826110c0565b845250602092830192919091019060010161108d565b5050505092915050565b60006110cc823561149f565b9392505050565b6000601f820183136110e457600080fd5b81356110f261106082611477565b9150808252602083016020830185838301111561110e57600080fd5b6111198382846114b7565b50505092915050565b60006020828403121561113457600080fd5b600061114084846110c0565b949350505050565b600080600080600060a0868803121561116057600080fd5b600061116c88886110c0565b955050602086013567ffffffffffffffff81111561118957600080fd5b611195888289016110d3565b945050604086013567ffffffffffffffff8111156111b257600080fd5b6111be888289016110d3565b93505060606111cf888289016110c0565b925050608086013567ffffffffffffffff8111156111ec57600080fd5b6111f888828901611041565b9150509295509295909350565b60006020828403121561121757600080fd5b813567ffffffffffffffff81111561122e57600080fd5b611140848285016110d3565b6000806040838503121561124d57600080fd5b823567ffffffffffffffff81111561126457600080fd5b611270858286016110d3565b925050602083013567ffffffffffffffff81111561128d57600080fd5b611299858286016110d3565b9150509250929050565b6112ac816114a8565b6112b58261149f565b60005b828110156112e5576112cb858351611394565b6112d4826114a2565b6020959095019491506001016112b8565b5050505050565b60006112f7826114ae565b808452602084019350611309836114a2565b60005b828110156113395761131f868351611394565b611328826114a2565b60209690960195915060010161130c565b5093949350505050565b61134c816114a8565b6113558261149f565b60005b828110156112e55761136b858351611394565b611374826114a2565b602095909501949150600101611358565b61138e816114b2565b82525050565b61138e8161149f565b60c081016113ab82866112a3565b6113b860408301856112a3565b6111406080830184611343565b602081016113d38284611385565b92915050565b602081016113d38284611394565b606081016113f58286611394565b6114026020830185611394565b6111406040830184611394565b6040810161141d8285611394565b818103602083015261114081846112ec565b60405181810167ffffffffffffffff8111828210171561144e57600080fd5b604052919050565b600067ffffffffffffffff82111561146d57600080fd5b5060209081020190565b600067ffffffffffffffff82111561148e57600080fd5b506020601f91909101601f19160190565b90565b60200190565b50600290565b5190565b151590565b828183375060009101525600a265627a7a72305820910360ca0293b325436c58e5dd43837d4b06187240d93f41d7aaedf1a71ac5026c6578706572696d656e74616cf50037a165627a7a72305820c5fd38a84d1722b65b5d3a4223813d4bba0a04697474a8c10c45abfd5ed71ad40029`

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
// Solidity: function EROs( uint256) constant returns(timestamp uint64, isExit bool, isTransfer bool, finalized bool, challenged bool, value uint128, requestor address, to address, trieKey bytes32, trieValue bytes32)
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
	})
	out := ret
	err := _RootChain.contract.Call(opts, out, "EROs", arg0)
	return *ret, err
}

// EROs is a free data retrieval call binding the contract method 0xb443f3cc.
//
// Solidity: function EROs( uint256) constant returns(timestamp uint64, isExit bool, isTransfer bool, finalized bool, challenged bool, value uint128, requestor address, to address, trieKey bytes32, trieValue bytes32)
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
}, error) {
	return _RootChain.Contract.EROs(&_RootChain.CallOpts, arg0)
}

// EROs is a free data retrieval call binding the contract method 0xb443f3cc.
//
// Solidity: function EROs( uint256) constant returns(timestamp uint64, isExit bool, isTransfer bool, finalized bool, challenged bool, value uint128, requestor address, to address, trieKey bytes32, trieValue bytes32)
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
}, error) {
	return _RootChain.Contract.EROs(&_RootChain.CallOpts, arg0)
}

// ERUs is a free data retrieval call binding the contract method 0xf4f31de4.
//
// Solidity: function ERUs( uint256) constant returns(timestamp uint64, isExit bool, isTransfer bool, finalized bool, challenged bool, value uint128, requestor address, to address, trieKey bytes32, trieValue bytes32)
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
	})
	out := ret
	err := _RootChain.contract.Call(opts, out, "ERUs", arg0)
	return *ret, err
}

// ERUs is a free data retrieval call binding the contract method 0xf4f31de4.
//
// Solidity: function ERUs( uint256) constant returns(timestamp uint64, isExit bool, isTransfer bool, finalized bool, challenged bool, value uint128, requestor address, to address, trieKey bytes32, trieValue bytes32)
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
}, error) {
	return _RootChain.Contract.ERUs(&_RootChain.CallOpts, arg0)
}

// ERUs is a free data retrieval call binding the contract method 0xf4f31de4.
//
// Solidity: function ERUs( uint256) constant returns(timestamp uint64, isExit bool, isTransfer bool, finalized bool, challenged bool, value uint128, requestor address, to address, trieKey bytes32, trieValue bytes32)
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
// Solidity: function ORBs( uint256) constant returns(submitted bool, epochNumber uint64, requestStart uint64, requestEnd uint64, trie address, transactionsRoot bytes32)
func (_RootChain *RootChainCaller) ORBs(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Submitted        bool
	EpochNumber      uint64
	RequestStart     uint64
	RequestEnd       uint64
	Trie             common.Address
	TransactionsRoot [32]byte
}, error) {
	ret := new(struct {
		Submitted        bool
		EpochNumber      uint64
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
// Solidity: function ORBs( uint256) constant returns(submitted bool, epochNumber uint64, requestStart uint64, requestEnd uint64, trie address, transactionsRoot bytes32)
func (_RootChain *RootChainSession) ORBs(arg0 *big.Int) (struct {
	Submitted        bool
	EpochNumber      uint64
	RequestStart     uint64
	RequestEnd       uint64
	Trie             common.Address
	TransactionsRoot [32]byte
}, error) {
	return _RootChain.Contract.ORBs(&_RootChain.CallOpts, arg0)
}

// ORBs is a free data retrieval call binding the contract method 0xea7f22a8.
//
// Solidity: function ORBs( uint256) constant returns(submitted bool, epochNumber uint64, requestStart uint64, requestEnd uint64, trie address, transactionsRoot bytes32)
func (_RootChain *RootChainCallerSession) ORBs(arg0 *big.Int) (struct {
	Submitted        bool
	EpochNumber      uint64
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
// Solidity: function URBs( uint256) constant returns(submitted bool, epochNumber uint64, requestStart uint64, requestEnd uint64, trie address, transactionsRoot bytes32)
func (_RootChain *RootChainCaller) URBs(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Submitted        bool
	EpochNumber      uint64
	RequestStart     uint64
	RequestEnd       uint64
	Trie             common.Address
	TransactionsRoot [32]byte
}, error) {
	ret := new(struct {
		Submitted        bool
		EpochNumber      uint64
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
// Solidity: function URBs( uint256) constant returns(submitted bool, epochNumber uint64, requestStart uint64, requestEnd uint64, trie address, transactionsRoot bytes32)
func (_RootChain *RootChainSession) URBs(arg0 *big.Int) (struct {
	Submitted        bool
	EpochNumber      uint64
	RequestStart     uint64
	RequestEnd       uint64
	Trie             common.Address
	TransactionsRoot [32]byte
}, error) {
	return _RootChain.Contract.URBs(&_RootChain.CallOpts, arg0)
}

// URBs is a free data retrieval call binding the contract method 0xc0e86064.
//
// Solidity: function URBs( uint256) constant returns(submitted bool, epochNumber uint64, requestStart uint64, requestEnd uint64, trie address, transactionsRoot bytes32)
func (_RootChain *RootChainCallerSession) URBs(arg0 *big.Int) (struct {
	Submitted        bool
	EpochNumber      uint64
	RequestStart     uint64
	RequestEnd       uint64
	Trie             common.Address
	TransactionsRoot [32]byte
}, error) {
	return _RootChain.Contract.URBs(&_RootChain.CallOpts, arg0)
}

// Blocks is a free data retrieval call binding the contract method 0xf4f911db.
//
// Solidity: function blocks( uint256,  uint256) constant returns(forkNumber uint64, epochNumber uint64, requestBlockId uint64, timestamp uint64, statesRoot bytes32, transactionsRoot bytes32, intermediateStatesRoot bytes32, isRequest bool, userActivated bool, challenged bool, challenging bool, finalized bool)
func (_RootChain *RootChainCaller) Blocks(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (struct {
	ForkNumber             uint64
	EpochNumber            uint64
	RequestBlockId         uint64
	Timestamp              uint64
	StatesRoot             [32]byte
	TransactionsRoot       [32]byte
	IntermediateStatesRoot [32]byte
	IsRequest              bool
	UserActivated          bool
	Challenged             bool
	Challenging            bool
	Finalized              bool
}, error) {
	ret := new(struct {
		ForkNumber             uint64
		EpochNumber            uint64
		RequestBlockId         uint64
		Timestamp              uint64
		StatesRoot             [32]byte
		TransactionsRoot       [32]byte
		IntermediateStatesRoot [32]byte
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
// Solidity: function blocks( uint256,  uint256) constant returns(forkNumber uint64, epochNumber uint64, requestBlockId uint64, timestamp uint64, statesRoot bytes32, transactionsRoot bytes32, intermediateStatesRoot bytes32, isRequest bool, userActivated bool, challenged bool, challenging bool, finalized bool)
func (_RootChain *RootChainSession) Blocks(arg0 *big.Int, arg1 *big.Int) (struct {
	ForkNumber             uint64
	EpochNumber            uint64
	RequestBlockId         uint64
	Timestamp              uint64
	StatesRoot             [32]byte
	TransactionsRoot       [32]byte
	IntermediateStatesRoot [32]byte
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
// Solidity: function blocks( uint256,  uint256) constant returns(forkNumber uint64, epochNumber uint64, requestBlockId uint64, timestamp uint64, statesRoot bytes32, transactionsRoot bytes32, intermediateStatesRoot bytes32, isRequest bool, userActivated bool, challenged bool, challenging bool, finalized bool)
func (_RootChain *RootChainCallerSession) Blocks(arg0 *big.Int, arg1 *big.Int) (struct {
	ForkNumber             uint64
	EpochNumber            uint64
	RequestBlockId         uint64
	Timestamp              uint64
	StatesRoot             [32]byte
	TransactionsRoot       [32]byte
	IntermediateStatesRoot [32]byte
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

// StartExit is a paid mutator transaction binding the contract method 0xcc9374b2.
//
// Solidity: function startExit(_isTransfer bool, _to address, _trieKey bytes32, _trieValue bytes32) returns(success bool)
func (_RootChain *RootChainTransactor) StartExit(opts *bind.TransactOpts, _isTransfer bool, _to common.Address, _trieKey [32]byte, _trieValue [32]byte) (*types.Transaction, error) {
	return _RootChain.contract.Transact(opts, "startExit", _isTransfer, _to, _trieKey, _trieValue)
}

// StartExit is a paid mutator transaction binding the contract method 0xcc9374b2.
//
// Solidity: function startExit(_isTransfer bool, _to address, _trieKey bytes32, _trieValue bytes32) returns(success bool)
func (_RootChain *RootChainSession) StartExit(_isTransfer bool, _to common.Address, _trieKey [32]byte, _trieValue [32]byte) (*types.Transaction, error) {
	return _RootChain.Contract.StartExit(&_RootChain.TransactOpts, _isTransfer, _to, _trieKey, _trieValue)
}

// StartExit is a paid mutator transaction binding the contract method 0xcc9374b2.
//
// Solidity: function startExit(_isTransfer bool, _to address, _trieKey bytes32, _trieValue bytes32) returns(success bool)
func (_RootChain *RootChainTransactorSession) StartExit(_isTransfer bool, _to common.Address, _trieKey [32]byte, _trieValue [32]byte) (*types.Transaction, error) {
	return _RootChain.Contract.StartExit(&_RootChain.TransactOpts, _isTransfer, _to, _trieKey, _trieValue)
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
