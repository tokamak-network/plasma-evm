// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package rootchain

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
var AddressBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a723058206fc5747bf3aa42e5e3b7cbe598456385bf018515ad193018c47587792e19655e64736f6c63430005090032"

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
var BMTBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a72305820ed3ff2f9a57068d474c79eee405d5ab813895bb146476d4963833cc4d0c2a3a464736f6c63430005090032"

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

// DataFuncSigs maps the 4-byte function signature to its string representation.
var DataFuncSigs = map[string]string{
	"a89ca766": "APPLY_IN_CHILDCHAIN_SIGNATURE()",
	"a7b6ae28": "APPLY_IN_ROOTCHAIN_SIGNATURE()",
	"ab73ff05": "NA()",
	"90e84f56": "NA_TX_GAS_LIMIT()",
	"1927ac58": "NA_TX_GAS_PRICE()",
}

// DataBin is the compiled bytecode used for deploying new contracts.
var DataBin = "0x6101cf610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100615760003560e01c80631927ac581461006657806390e84f5614610080578063a7b6ae2814610088578063a89ca766146100ad578063ab73ff05146100b5575b600080fd5b61006e6100d9565b60408051918252519081900360200190f35b61006e6100e1565b6100906100e8565b604080516001600160e01b03199092168252519081900360200190f35b610090610103565b6100bd61011e565b604080516001600160a01b039092168252519081900360200190f35b633b9aca0081565b620186a081565b60405180603b6101608239603b019050604051809103902081565b60405180603c6101248239603c019050604051809103902081565b60008156fe6170706c7952657175657374496e4368696c64436861696e28626f6f6c2c75696e743235362c616464726573732c627974657333322c6279746573296170706c7952657175657374496e526f6f74436861696e28626f6f6c2c75696e743235362c616464726573732c627974657333322c627974657329a265627a7a72305820aaf0362f1175f8c322b4e191a40a7239a330933a21369a714c9db14d4d7fe57164736f6c63430005090032"

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
var MathBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a723058206088d9a17c63d845fc404648ce68aba3c54f7dbdff2d9924b40a8968ead52c2464736f6c63430005090032"

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
var RLPBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a723058204db3d0cb00c0c2e12d8032084b4855a375dca50756a4038557090ded2527ab8f64736f6c63430005090032"

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
var RLPEncodeBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a72305820cbe97fecd12ec670b6bce1a4cfbc545712415858846d4d018e052f1fbd2c5ef064736f6c63430005090032"

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

// RootChainABI is the input ABI used to generate the binding from.
const RootChainABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"forkNumber\",\"type\":\"uint256\"}],\"name\":\"getLastFinalizedEpoch\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"COST_URB_PREPARE\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CP_COMPUTATION\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_pos1\",\"type\":\"uint256\"},{\"name\":\"_pos2\",\"type\":\"uint256\"},{\"name\":\"_epochStateRoot\",\"type\":\"bytes32\"},{\"name\":\"_epochTransactionsRoot\",\"type\":\"bytes32\"},{\"name\":\"_epochReceiptsRoot\",\"type\":\"bytes32\"}],\"name\":\"submitNRE\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"forkNumber\",\"type\":\"uint256\"}],\"name\":\"lastEpoch\",\"outputs\":[{\"name\":\"lastBlock\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastAppliedForkNumber\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"currentFork\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"COST_URB\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"forkNumber\",\"type\":\"uint256\"}],\"name\":\"lastBlock\",\"outputs\":[{\"name\":\"lastBlock\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"numEnterForORB\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_trieKey\",\"type\":\"bytes32\"},{\"name\":\"_trieValue\",\"type\":\"bytes\"}],\"name\":\"startEnter\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"forkNumber\",\"type\":\"uint256\"},{\"name\":\"epochNumber\",\"type\":\"uint256\"}],\"name\":\"getEpoch\",\"outputs\":[{\"components\":[{\"name\":\"startBlockNumber\",\"type\":\"uint64\"},{\"name\":\"endBlockNumber\",\"type\":\"uint64\"},{\"name\":\"timestamp\",\"type\":\"uint64\"},{\"name\":\"isEmpty\",\"type\":\"bool\"},{\"name\":\"initialized\",\"type\":\"bool\"},{\"name\":\"isRequest\",\"type\":\"bool\"},{\"name\":\"userActivated\",\"type\":\"bool\"},{\"name\":\"rebase\",\"type\":\"bool\"},{\"components\":[{\"name\":\"requestStart\",\"type\":\"uint64\"},{\"name\":\"requestEnd\",\"type\":\"uint64\"},{\"name\":\"firstRequestBlockId\",\"type\":\"uint64\"},{\"name\":\"numEnter\",\"type\":\"uint64\"},{\"name\":\"nextEnterEpoch\",\"type\":\"uint64\"},{\"name\":\"nextEpoch\",\"type\":\"uint64\"}],\"name\":\"RE\",\"type\":\"tuple\"},{\"components\":[{\"name\":\"epochStateRoot\",\"type\":\"bytes32\"},{\"name\":\"epochTransactionsRoot\",\"type\":\"bytes32\"},{\"name\":\"epochReceiptsRoot\",\"type\":\"bytes32\"},{\"name\":\"submittedAt\",\"type\":\"uint64\"},{\"name\":\"finalizedAt\",\"type\":\"uint64\"},{\"name\":\"finalized\",\"type\":\"bool\"},{\"name\":\"challenging\",\"type\":\"bool\"},{\"name\":\"challenged\",\"type\":\"bool\"}],\"name\":\"NRE\",\"type\":\"tuple\"}],\"name\":\"epoch\",\"type\":\"tuple\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"EROIdToFinalize\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getLastEpoch\",\"outputs\":[{\"components\":[{\"name\":\"startBlockNumber\",\"type\":\"uint64\"},{\"name\":\"endBlockNumber\",\"type\":\"uint64\"},{\"name\":\"timestamp\",\"type\":\"uint64\"},{\"name\":\"isEmpty\",\"type\":\"bool\"},{\"name\":\"initialized\",\"type\":\"bool\"},{\"name\":\"isRequest\",\"type\":\"bool\"},{\"name\":\"userActivated\",\"type\":\"bool\"},{\"name\":\"rebase\",\"type\":\"bool\"},{\"components\":[{\"name\":\"requestStart\",\"type\":\"uint64\"},{\"name\":\"requestEnd\",\"type\":\"uint64\"},{\"name\":\"firstRequestBlockId\",\"type\":\"uint64\"},{\"name\":\"numEnter\",\"type\":\"uint64\"},{\"name\":\"nextEnterEpoch\",\"type\":\"uint64\"},{\"name\":\"nextEpoch\",\"type\":\"uint64\"}],\"name\":\"RE\",\"type\":\"tuple\"},{\"components\":[{\"name\":\"epochStateRoot\",\"type\":\"bytes32\"},{\"name\":\"epochTransactionsRoot\",\"type\":\"bytes32\"},{\"name\":\"epochReceiptsRoot\",\"type\":\"bytes32\"},{\"name\":\"submittedAt\",\"type\":\"uint64\"},{\"name\":\"finalizedAt\",\"type\":\"uint64\"},{\"name\":\"finalized\",\"type\":\"bool\"},{\"name\":\"challenging\",\"type\":\"bool\"},{\"name\":\"challenged\",\"type\":\"bool\"}],\"name\":\"NRE\",\"type\":\"tuple\"}],\"name\":\"\",\"type\":\"tuple\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_forkNumber\",\"type\":\"uint256\"},{\"name\":\"_blockNumber\",\"type\":\"uint256\"},{\"name\":\"_index\",\"type\":\"uint256\"},{\"name\":\"_receiptData\",\"type\":\"bytes\"},{\"name\":\"_proof\",\"type\":\"bytes\"}],\"name\":\"challengeExit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"forkNumber\",\"type\":\"uint256\"},{\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"getBlock\",\"outputs\":[{\"components\":[{\"name\":\"epochNumber\",\"type\":\"uint64\"},{\"name\":\"requestBlockId\",\"type\":\"uint64\"},{\"name\":\"timestamp\",\"type\":\"uint64\"},{\"name\":\"finalizedAt\",\"type\":\"uint64\"},{\"name\":\"referenceBlock\",\"type\":\"uint64\"},{\"name\":\"statesRoot\",\"type\":\"bytes32\"},{\"name\":\"transactionsRoot\",\"type\":\"bytes32\"},{\"name\":\"receiptsRoot\",\"type\":\"bytes32\"},{\"name\":\"isRequest\",\"type\":\"bool\"},{\"name\":\"userActivated\",\"type\":\"bool\"},{\"name\":\"challenged\",\"type\":\"bool\"},{\"name\":\"challenging\",\"type\":\"bool\"},{\"name\":\"finalized\",\"type\":\"bool\"}],\"name\":\"\",\"type\":\"tuple\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"forks\",\"outputs\":[{\"name\":\"forkedBlock\",\"type\":\"uint64\"},{\"name\":\"firstEpoch\",\"type\":\"uint64\"},{\"name\":\"lastEpoch\",\"type\":\"uint64\"},{\"name\":\"firstBlock\",\"type\":\"uint64\"},{\"name\":\"lastBlock\",\"type\":\"uint64\"},{\"name\":\"lastFinalizedEpoch\",\"type\":\"uint64\"},{\"name\":\"lastFinalizedBlock\",\"type\":\"uint64\"},{\"name\":\"timestamp\",\"type\":\"uint64\"},{\"name\":\"firstEnterEpoch\",\"type\":\"uint64\"},{\"name\":\"lastEnterEpoch\",\"type\":\"uint64\"},{\"name\":\"nextBlockToRebase\",\"type\":\"uint64\"},{\"name\":\"rebased\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"n\",\"type\":\"uint256\"}],\"name\":\"finalizeRequests\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"operator\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"forkNumber\",\"type\":\"uint256\"},{\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"getBlockFinalizedAt\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_blockNumber\",\"type\":\"uint256\"},{\"name\":\"_key\",\"type\":\"bytes\"},{\"name\":\"_txByte\",\"type\":\"bytes\"},{\"name\":\"_branchMask\",\"type\":\"uint256\"},{\"name\":\"_siblings\",\"type\":\"bytes32[]\"}],\"name\":\"challengeNullAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_pos\",\"type\":\"uint256\"},{\"name\":\"_statesRoot\",\"type\":\"bytes32\"},{\"name\":\"_transactionsRoot\",\"type\":\"bytes32\"},{\"name\":\"_receiptsRoot\",\"type\":\"bytes32\"}],\"name\":\"submitURB\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"firstFilledORENumber\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"finalizeBlock\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_trieKey\",\"type\":\"bytes32\"},{\"name\":\"_trieValue\",\"type\":\"bytes\"}],\"name\":\"makeERU\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"development\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CP_EXIT\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"COST_ERU\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"REQUEST_GAS\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAX_REQUESTS\",\"outputs\":[{\"name\":\"maxRequests\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"COST_NRB\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"finalizeRequest\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_pos\",\"type\":\"uint256\"},{\"name\":\"_statesRoot\",\"type\":\"bytes32\"},{\"name\":\"_transactionsRoot\",\"type\":\"bytes32\"},{\"name\":\"_receiptsRoot\",\"type\":\"bytes32\"}],\"name\":\"submitORB\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"NRELength\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CP_WITHHOLDING\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"COST_ORB\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"EROs\",\"outputs\":[{\"name\":\"timestamp\",\"type\":\"uint64\"},{\"name\":\"isExit\",\"type\":\"bool\"},{\"name\":\"isTransfer\",\"type\":\"bool\"},{\"name\":\"finalized\",\"type\":\"bool\"},{\"name\":\"challenged\",\"type\":\"bool\"},{\"name\":\"value\",\"type\":\"uint128\"},{\"name\":\"requestor\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"trieKey\",\"type\":\"bytes32\"},{\"name\":\"hash\",\"type\":\"bytes32\"},{\"name\":\"trieValue\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getNumEROs\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"lastNonEmptyRequestEpoch\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"etherToken\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"URBs\",\"outputs\":[{\"name\":\"submitted\",\"type\":\"bool\"},{\"name\":\"numEnter\",\"type\":\"uint64\"},{\"name\":\"epochNumber\",\"type\":\"uint64\"},{\"name\":\"requestStart\",\"type\":\"uint64\"},{\"name\":\"requestEnd\",\"type\":\"uint64\"},{\"name\":\"trie\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PREPARE_TIMEOUT\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ERUIdToFinalize\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastAppliedEpochNumber\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"firstNonEmptyRequestEpoch\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_rootchain\",\"type\":\"address\"},{\"name\":\"_childchain\",\"type\":\"address\"}],\"name\":\"mapRequestableContractByOperator\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_forkNumber\",\"type\":\"uint256\"}],\"name\":\"forked\",\"outputs\":[{\"name\":\"result\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_requestId\",\"type\":\"uint256\"}],\"name\":\"getEROBytes\",\"outputs\":[{\"name\":\"out\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"forkNumber\",\"type\":\"uint256\"}],\"name\":\"getLastFinalizedBlock\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"COST_ERO\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"requestableContracts\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"NULL_ADDRESS\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"submitHandler\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"prepareToSubmitURB\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"epochHandler\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_trieKey\",\"type\":\"bytes32\"},{\"name\":\"_trieValue\",\"type\":\"bytes\"}],\"name\":\"startExit\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getNumORBs\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"ORBs\",\"outputs\":[{\"name\":\"submitted\",\"type\":\"bool\"},{\"name\":\"numEnter\",\"type\":\"uint64\"},{\"name\":\"epochNumber\",\"type\":\"uint64\"},{\"name\":\"requestStart\",\"type\":\"uint64\"},{\"name\":\"requestEnd\",\"type\":\"uint64\"},{\"name\":\"trie\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_requestId\",\"type\":\"uint256\"},{\"name\":\"_userActivated\",\"type\":\"bool\"}],\"name\":\"getRequestFinalized\",\"outputs\":[{\"name\":\"finalized\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"ERUs\",\"outputs\":[{\"name\":\"timestamp\",\"type\":\"uint64\"},{\"name\":\"isExit\",\"type\":\"bool\"},{\"name\":\"isTransfer\",\"type\":\"bool\"},{\"name\":\"finalized\",\"type\":\"bool\"},{\"name\":\"challenged\",\"type\":\"bool\"},{\"name\":\"value\",\"type\":\"uint128\"},{\"name\":\"requestor\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"trieKey\",\"type\":\"bytes32\"},{\"name\":\"hash\",\"type\":\"bytes32\"},{\"name\":\"trieValue\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastAppliedBlockNumber\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_epochHandler\",\"type\":\"address\"},{\"name\":\"_submitHandler\",\"type\":\"address\"},{\"name\":\"_etherToken\",\"type\":\"address\"},{\"name\":\"_development\",\"type\":\"bool\"},{\"name\":\"_NRELength\",\"type\":\"uint256\"},{\"name\":\"_statesRoot\",\"type\":\"bytes32\"},{\"name\":\"_transactionsRoot\",\"type\":\"bytes32\"},{\"name\":\"_receiptsRoot\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"userActivated\",\"type\":\"bool\"}],\"name\":\"SessionTimeout\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newFork\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"epochNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"forkedBlockNumber\",\"type\":\"uint256\"}],\"name\":\"Forked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"forkNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"epochNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"startBlockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"endBlockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"requestStart\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"requestEnd\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"epochIsEmpty\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"isRequest\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"userActivated\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"rebase\",\"type\":\"bool\"}],\"name\":\"EpochPrepared\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"forkNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"epochNumber\",\"type\":\"uint256\"}],\"name\":\"EpochFilling\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"forkNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"epochNumber\",\"type\":\"uint256\"}],\"name\":\"EpochFilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"forkNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"epochNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"startBlockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"endBlockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"requestStart\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"requestEnd\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"epochIsEmpty\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"isRequest\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"userActivated\",\"type\":\"bool\"}],\"name\":\"EpochRebased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"fork\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"epochNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"isRequest\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"userActivated\",\"type\":\"bool\"}],\"name\":\"BlockSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"requestor\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"weiAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"trieKey\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"trieValue\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"isExit\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"userActivated\",\"type\":\"bool\"}],\"name\":\"RequestCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"requestor\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"trieKey\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"trieValue\",\"type\":\"bytes32\"}],\"name\":\"ERUCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"forkNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"BlockFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"forkNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"epochNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"startBlockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"endBlockNumber\",\"type\":\"uint256\"}],\"name\":\"EpochFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"userActivated\",\"type\":\"bool\"}],\"name\":\"RequestFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"userActivated\",\"type\":\"bool\"}],\"name\":\"RequestApplied\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"userActivated\",\"type\":\"bool\"}],\"name\":\"RequestChallenged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"contractInRootchain\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"contractInChildchain\",\"type\":\"address\"}],\"name\":\"RequestableContractMapped\",\"type\":\"event\"}]"

// RootChainFuncSigs maps the 4-byte function signature to its string representation.
var RootChainFuncSigs = map[string]string{
	"d691acd8": "COST_ERO()",
	"8b5172d0": "COST_ERU()",
	"94be3aa5": "COST_NRB()",
	"b2ae9ba8": "COST_ORB()",
	"192adc5b": "COST_URB()",
	"033cfbed": "COST_URB_PREPARE()",
	"08c4fff0": "CP_COMPUTATION()",
	"8155717d": "CP_EXIT()",
	"b17fa6e9": "CP_WITHHOLDING()",
	"2dc6bb7b": "EROIdToFinalize()",
	"b443f3cc": "EROs(uint256)",
	"c54626cc": "ERUIdToFinalize()",
	"f4f31de4": "ERUs(uint256)",
	"93521222": "MAX_REQUESTS()",
	"ab96da2d": "NRELength()",
	"de0ce17d": "NULL_ADDRESS()",
	"ea7f22a8": "ORBs(uint256)",
	"c2bc88fa": "PREPARE_TIMEOUT()",
	"8eb288ca": "REQUEST_GAS()",
	"c0e86064": "URBs(uint256)",
	"404f7d66": "challengeExit(uint256,uint256,uint256,bytes,bytes)",
	"6299fb24": "challengeNullAddress(uint256,bytes,bytes,uint256,bytes32[])",
	"183d2d1c": "currentFork()",
	"7b929c27": "development()",
	"e7b88b80": "epochHandler()",
	"b8066bcb": "etherToken()",
	"75395a58": "finalizeBlock()",
	"99bd3600": "finalizeRequest()",
	"54768571": "finalizeRequests(uint256)",
	"72ecb9a8": "firstFilledORENumber(uint256)",
	"ca6f6380": "firstNonEmptyRequestEpoch(uint256)",
	"ce8a2bc2": "forked(uint256)",
	"4ba3a126": "forks(uint256)",
	"4a44bdb8": "getBlock(uint256,uint256)",
	"5b884682": "getBlockFinalizedAt(uint256,uint256)",
	"d1723a96": "getEROBytes(uint256)",
	"2b25a38b": "getEpoch(uint256,uint256)",
	"398bac63": "getLastEpoch()",
	"d636857e": "getLastFinalizedBlock(uint256)",
	"019dc099": "getLastFinalizedEpoch(uint256)",
	"b540adba": "getNumEROs()",
	"ea0c73f6": "getNumORBs()",
	"f28a7afa": "getRequestFinalized(uint256,bool)",
	"fb788a27": "lastAppliedBlockNumber()",
	"c8ad329f": "lastAppliedEpochNumber()",
	"164bc2ae": "lastAppliedForkNumber()",
	"1ec2042b": "lastBlock(uint256)",
	"11e4c914": "lastEpoch(uint256)",
	"b6715647": "lastNonEmptyRequestEpoch(uint256)",
	"78fe705f": "makeERU(address,bytes32,bytes)",
	"cb5d742f": "mapRequestableContractByOperator(address,address)",
	"23691566": "numEnterForORB()",
	"570ca735": "operator()",
	"e6925d08": "prepareToSubmitURB()",
	"da0185f8": "requestableContracts(address)",
	"2aa196c8": "startEnter(address,bytes32,bytes)",
	"e7f96505": "startExit(address,bytes32,bytes)",
	"e259faf7": "submitHandler()",
	"0eaf45a8": "submitNRE(uint256,uint256,bytes32,bytes32,bytes32)",
	"a820c067": "submitORB(uint256,bytes32,bytes32,bytes32)",
	"6f3e4b90": "submitURB(uint256,bytes32,bytes32,bytes32)",
}

// RootChainBin is the compiled bytecode used for deploying new contracts.
var RootChainBin = "0x60806040523480156200001157600080fd5b5060405162005a9338038062005a938339810160408190526200003491620004f6565b62000053886001600160a01b03166200023c60201b620032291760201c565b6200005d57600080fd5b6200007c876001600160a01b03166200023c60201b620032291760201c565b6200008657600080fd5b620000a5866001600160a01b03166200023c60201b620032291760201c565b620000af57600080fd5b600180546001600160a01b03199081166001600160a01b038b8116919091179092556002805482168a8416178155600380549092169289169290921781556000805460ff191688151517610100600160a81b03191661010033021781556004878155600554825260076020908152604080842084805280840183528185208088018b90558087018a90559384018890559485019091528083208054790100000000000000000000000000000000000000000000000000600160801b600160c01b0319909116700100000000000000000000000000000000426001600160401b0316021760ff60c81b191617905593825292812080547a010000000000000000000000000000000000000000000000000000780100000000000000000000000000000000000000000000000060ff60c01b199092169190911760ff60d01b191617905590919062000203908390839062000242565b620002198260006001600160e01b036200032716565b6200022c6001600160e01b03620003fa16565b5050505050505050505062000770565b3b151590565b6005828101805464010000000060ff60201b1990911617905582546001600160c01b03167801000000000000000000000000000000000000000000000000426001600160401b0390811691909102919091178455600185018054600160801b600160c01b03191670010000000000000000000000000000000085841602178082558554600160401b600160801b031990911692166801000000000000000002919091179055546040517ffb96205e4b3633fd57aa805b26b51ecf528714a10241a4af015929dce86768d9916200031a91849062000665565b60405180910390a1505050565b60008181526003830160205260409081902060068101805470010000000000000000000000000000000060ff60801b199091168117600160401b600160801b031990811668010000000000000000426001600160401b039081168202929092179094558454600189018054600160801b600160c01b031916918690048316909402179091168187168402179091556005548354945193947f70801d4d63b3da6c19ba7349911f45bed5a99ccdfb51b8138c105872529bebd5946200031a94929388938382169304169062000684565b5050565b6001546040516000916060916001600160a01b03909116906200041d9062000658565b60408051918290038220600483526024830182526020830180516001600160e01b03167fffffffff000000000000000000000000000000000000000000000000000000009092169190911790525162000477919062000643565b600060405180830381855af49150503d8060008114620004b4576040519150601f19603f3d011682016040523d82523d6000602084013e620004b9565b606091505b509150915081620003f657600080fd5b8051620004d68162000740565b92915050565b8051620004d6816200075a565b8051620004d68162000765565b600080600080600080600080610100898b0312156200051457600080fd5b6000620005228b8b620004c9565b9850506020620005358b828c01620004c9565b9750506040620005488b828c01620004c9565b96505060606200055b8b828c01620004dc565b95505060806200056e8b828c01620004e9565b94505060a0620005818b828c01620004e9565b93505060c0620005948b828c01620004e9565b92505060e0620005a78b828c01620004e9565b9150509295985092959890939650565b6000620005c482620006ca565b620005d08185620006ce565b9350620005e28185602086016200070d565b9290920192915050565b6000620005fb600c83620006ce565b7f707265706172654e5245282900000000000000000000000000000000000000008152600c0192915050565b6200063281620006e5565b82525050565b620006328162000700565b6000620006518284620005b7565b9392505050565b6000620004d682620005ec565b6040810162000675828562000627565b62000651602083018462000627565b6080810162000694828762000627565b620006a3602083018662000627565b620006b2604083018562000638565b620006c1606083018462000638565b95945050505050565b5190565b919050565b6000620004d682620006e8565b151590565b90565b6001600160a01b031690565b6001600160401b031690565b6000620004d682620006f4565b60005b838110156200072a57818101518382015260200162000710565b838111156200073a576000848401525b50505050565b6200074b81620006d3565b81146200075757600080fd5b50565b6200074b81620006e0565b6200074b81620006e5565b61531380620007806000396000f3fe6080604052600436106103a25760003560e01c806394be3aa5116101e7578063cb5d742f1161010d578063e6925d08116100a0578063ea7f22a81161006f578063ea7f22a814610972578063f28a7afa14610992578063f4f31de4146109b2578063fb788a27146109d2576103a2565b8063e6925d081461092d578063e7b88b8014610935578063e7f965051461094a578063ea0c73f61461095d576103a2565b8063d691acd8116100dc578063d691acd8146103dd578063da0185f8146108e3578063de0ce17d14610903578063e259faf714610918576103a2565b8063cb5d742f14610856578063ce8a2bc214610876578063d1723a9614610896578063d636857e146108c3576103a2565b8063b540adba11610185578063c2bc88fa11610154578063c2bc88fa146107f7578063c54626cc1461080c578063c8ad329f14610821578063ca6f638014610836576103a2565b8063b540adba1461077b578063b671564714610790578063b8066bcb146107b0578063c0e86064146107c5576103a2565b8063ab96da2d116101c1578063ab96da2d1461071a578063b17fa6e91461072f578063b2ae9ba8146103dd578063b443f3cc14610744576103a2565b806394be3aa5146103dd57806399bd3600146106f2578063a820c06714610707576103a2565b80634a44bdb8116102cc57806372ecb9a81161026a5780638155717d116102395780638155717d1461069e5780638b5172d0146106b35780638eb288ca146106c857806393521222146106dd576103a2565b806372ecb9a81461064157806375395a581461066157806378fe705f146106765780637b929c2714610689576103a2565b8063570ca735116102a6578063570ca735146105cc5780635b884682146105ee5780636299fb241461060e5780636f3e4b901461062e576103a2565b80634a44bdb8146105475780634ba3a1261461057457806354768571146105ac576103a2565b8063192adc5b116103445780632b25a38b116103135780632b25a38b146104ce5780632dc6bb7b146104fb578063398bac6314610510578063404f7d6614610525576103a2565b8063192adc5b146104715780631ec2042b1461048657806323691566146104a65780632aa196c8146104bb576103a2565b80630eaf45a8116103805780630eaf45a81461040757806311e4c91414610427578063164bc2ae14610447578063183d2d1c1461045c576103a2565b8063019dc099146103a7578063033cfbed146103dd57806308c4fff0146103f2575b600080fd5b3480156103b357600080fd5b506103c76103c2366004614514565b6109e7565b6040516103d49190614eb1565b60405180910390f35b3480156103e957600080fd5b506103c7610a0f565b3480156103fe57600080fd5b506103c7610a19565b61041a6104153660046146ba565b610a1e565b6040516103d49190614da6565b34801561043357600080fd5b506103c7610442366004614514565b610a7b565b34801561045357600080fd5b506103c7610a9d565b34801561046857600080fd5b506103c7610aa3565b34801561047d57600080fd5b506103c7610aa9565b34801561049257600080fd5b506103c76104a1366004614514565b610ab3565b3480156104b257600080fd5b506103c7610ad1565b61041a6104c9366004614492565b610ad7565b3480156104da57600080fd5b506104ee6104e936600461468a565b610b99565b6040516103d49190614e93565b34801561050757600080fd5b506103c7610f01565b34801561051c57600080fd5b506104ee610f07565b34801561053157600080fd5b5061054561054036600461472f565b6112a2565b005b34801561055357600080fd5b5061056761056236600461468a565b611536565b6040516103d49190614ea2565b34801561058057600080fd5b5061059461058f366004614514565b611626565b6040516103d49c9b9a999897969594939291906150ff565b3480156105b857600080fd5b5061041a6105c7366004614514565b61169d565b3480156105d857600080fd5b506105e16116cb565b6040516103d49190614d7d565b3480156105fa57600080fd5b506103c761060936600461468a565b6116df565b34801561061a57600080fd5b506105456106293660046145c3565b611710565b61041a61063c366004614562565b611776565b34801561064d57600080fd5b506103c761065c366004614514565b61178c565b34801561066d57600080fd5b5061041a61179e565b61041a610684366004614492565b6117b7565b34801561069557600080fd5b5061041a611838565b3480156106aa57600080fd5b506103c7611841565b3480156106bf57600080fd5b506103c7611846565b3480156106d457600080fd5b506103c7611850565b3480156106e957600080fd5b506103c7611857565b3480156106fe57600080fd5b5061041a611866565b61041a610715366004614562565b61210e565b34801561072657600080fd5b506103c7612169565b34801561073b57600080fd5b506103c761216f565b34801561075057600080fd5b5061076461075f366004614514565b612174565b6040516103d49b9a99989796959493929190615055565b34801561078757600080fd5b506103c7612295565b34801561079c57600080fd5b506103c76107ab366004614514565b61229b565b3480156107bc57600080fd5b506105e16122ad565b3480156107d157600080fd5b506107e56107e0366004614514565b6122bc565b6040516103d496959493929190614dfb565b34801561080357600080fd5b506103c7612323565b34801561081857600080fd5b506103c7612328565b34801561082d57600080fd5b506103c761232e565b34801561084257600080fd5b506103c7610851366004614514565b612334565b34801561086257600080fd5b5061041a610871366004614458565b612346565b34801561088257600080fd5b5061041a610891366004614514565b612415565b3480156108a257600080fd5b506108b66108b1366004614514565b61241d565b6040516103d49190614e72565b3480156108cf57600080fd5b506103c76108de366004614514565b6125c6565b3480156108ef57600080fd5b506105e16108fe366004614432565b6125eb565b34801561090f57600080fd5b506105e1612606565b34801561092457600080fd5b506105e161260b565b61054561261a565b34801561094157600080fd5b506105e161262e565b61041a610958366004614492565b61263d565b34801561096957600080fd5b506103c76126ac565b34801561097e57600080fd5b506107e561098d366004614514565b6126b2565b34801561099e57600080fd5b5061041a6109ad366004614532565b6126bf565b3480156109be57600080fd5b506107646109cd366004614514565b612708565b3480156109de57600080fd5b506103c7612715565b600081815260076020526040902060010154600160401b90046001600160401b03165b919050565b6509184e72a00081565b600f81565b6000805461010090046001600160a01b03163314610a3b57600080fd5b6509184e72a00080341015610a4f57600080fd5b60005460ff16610a6357610a6161271b565b505b610a708787878787612891565b979650505050505050565b600090815260076020526040902054600160801b90046001600160401b031690565b600f5481565b60055481565b6551dac207a00081565b6000908152600760205260409020600101546001600160401b031690565b600c5481565b60008034610aed6008600a888489898880612970565b600c805460010190556005546000908152600760205260408082209051929450917f6940a01870e576ceb735867e13863646d517ce10e66c0133186a4ebdfe9388c291610b3c91869190614f73565b60405180910390a17f879922cf5fcada9ebaf8bd7424dc62877f4b220cae07fb6695cc1e8f94c52b4d833389858a8a600080604051610b82989796959493929190614f3d565b60405180910390a1600193505050505b9392505050565b610ba1613f14565b600760008481526020019081526020016000206003016000838152602001908152602001600020604051806101400160405290816000820160009054906101000a90046001600160401b03166001600160401b03166001600160401b031681526020016000820160089054906101000a90046001600160401b03166001600160401b03166001600160401b031681526020016000820160109054906101000a90046001600160401b03166001600160401b03166001600160401b031681526020016000820160189054906101000a900460ff161515151581526020016000820160199054906101000a900460ff1615151515815260200160008201601a9054906101000a900460ff1615151515815260200160008201601b9054906101000a900460ff1615151515815260200160008201601c9054906101000a900460ff16151515158152602001600182016040518060c00160405290816000820160009054906101000a90046001600160401b03166001600160401b03166001600160401b031681526020016000820160089054906101000a90046001600160401b03166001600160401b03166001600160401b031681526020016000820160109054906101000a90046001600160401b03166001600160401b03166001600160401b031681526020016000820160189054906101000a90046001600160401b03166001600160401b03166001600160401b031681526020016001820160009054906101000a90046001600160401b03166001600160401b03166001600160401b031681526020016001820160089054906101000a90046001600160401b03166001600160401b03166001600160401b031681525050815260200160038201604051806101000160405290816000820154815260200160018201548152602001600282015481526020016003820160009054906101000a90046001600160401b03166001600160401b03166001600160401b031681526020016003820160089054906101000a90046001600160401b03166001600160401b03166001600160401b031681526020016003820160109054906101000a900460ff161515151581526020016003820160119054906101000a900460ff161515151581526020016003820160129054906101000a900460ff1615151515815250508152505090505b92915050565b60125481565b610f0f613f14565b600760006005548152602001908152602001600020600301600060076000600554815260200190815260200160002060000160109054906101000a90046001600160401b03166001600160401b03168152602001908152602001600020604051806101400160405290816000820160009054906101000a90046001600160401b03166001600160401b03166001600160401b031681526020016000820160089054906101000a90046001600160401b03166001600160401b03166001600160401b031681526020016000820160109054906101000a90046001600160401b03166001600160401b03166001600160401b031681526020016000820160189054906101000a900460ff161515151581526020016000820160199054906101000a900460ff1615151515815260200160008201601a9054906101000a900460ff1615151515815260200160008201601b9054906101000a900460ff1615151515815260200160008201601c9054906101000a900460ff16151515158152602001600182016040518060c00160405290816000820160009054906101000a90046001600160401b03166001600160401b03166001600160401b031681526020016000820160089054906101000a90046001600160401b03166001600160401b03166001600160401b031681526020016000820160109054906101000a90046001600160401b03166001600160401b03166001600160401b031681526020016000820160189054906101000a90046001600160401b03166001600160401b03166001600160401b031681526020016001820160009054906101000a90046001600160401b03166001600160401b03166001600160401b031681526020016001820160089054906101000a90046001600160401b03166001600160401b03166001600160401b031681525050815260200160038201604051806101000160405290816000820154815260200160018201548152602001600282015481526020016003820160009054906101000a90046001600160401b03166001600160401b03166001600160401b031681526020016003820160089054906101000a90046001600160401b03166001600160401b03166001600160401b031681526020016003820160109054906101000a900460ff161515151581526020016003820160119054906101000a900460ff161515151581526020016003820160129054906101000a900460ff1615151515815250508152505090505b90565b600087815260076020908152604080832089845260048101909252909120600581015460ff166112d157600080fd5b6005810154640100000000900460ff166112ea57600080fd5b6005810154600090610100900460ff1680156113fa576113c383600b8560000160089054906101000a90046001600160401b03166001600160401b03168154811061133157fe5b906000526020600020906002020160098c8c8c8080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050508b8b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250612f9292505050565b60405190925033906000906512309ce540009082818181858883f193505050501580156113f4573d6000803e3d6000fd5b506114f0565b6114bd83600a8560000160089054906101000a90046001600160401b03166001600160401b03168154811061142b57fe5b906000526020600020906002020160088c8c8c8080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050508b8b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250612f9292505050565b60405190925033906000906509184e72a0009082818181858883f193505050501580156114ee573d6000803e3d6000fd5b505b7fc8135db115644ed4ae193313c4c801235ef740d2a57a8d5e6fe26ab66635698a8282604051611521929190614f73565b60405180910390a15050505050505050505050565b61153e613f74565b506000918252600760209081526040808420928452600492830182529283902083516101a08101855281546001600160401b038082168352600160401b8204811694830194909452600160801b8104841695820195909552600160c01b9094048216606085015260018101549091166080840152600281015460a0840152600381015460c08401529081015460e08301526005015460ff80821615156101008085019190915282048116151561012084015262010000820481161515610140840152630100000082048116151561016084015264010000000090910416151561018082015290565b6007602052600090815260409020805460018201546002909201546001600160401b0380831693600160401b808504831694600160801b808204851695600160c01b92839004861695858116958581048216958482048316959182900483169484841694918204841693908204169160ff9104168c565b6000805b828110156116c2576116b1611866565b6116ba57600080fd5b6001016116a1565b50600192915050565b60005461010090046001600160a01b031681565b600091825260076020908152604080842092845260049092019052902054600160c01b90046001600160401b031690565b6005805460009081526007602090815260408083208c8452600481019092529091209182015490919060ff161561174657600080fd5b8054426001600160401b03600160801b90920491909116600f011161176a57600080fd5b50505050505050505050565b60006551dac207a000803410156103a257600080fd5b60066020526000908152604090205481565b60006117a861271b565b6117b157600080fd5b50600190565b60006512309ce54000803410156117cd57600080fd5b60006117e36009600b8860008989600180612970565b90507f879922cf5fcada9ebaf8bd7424dc62877f4b220cae07fb6695cc1e8f94c52b4d81338860008989600180604051611824989796959493929190614ebf565b60405180910390a150600195945050505050565b60005460ff1681565b600a81565b6512309ce5400081565b620186a081565b600061186161309d565b905090565b600f5460009081526007602052604081206010546001820154601154849392916001600160401b0316101561189a57600080fd5b601154600090815260048301602090815260408083208484526003860190925290912083546001600160401b0316158015906118e4575083546011546001600160401b0390911611155b1561194f575050600f805460010190819055600090815260076020908152604080832080546001600160401b03600160401b82048116808752600384018652848720600160c01b90930490911660118190556010829055865260048301909452919093209294509092505b600582015460ff166119e1578261197957600f546000908152600e602052604090205492506119a0565b600092835260038401602052604090922060020154600160401b90046001600160401b0316915b826119aa57600080fd5b50506000818152600383016020908152604080832080546001600160401b0316601181905584526004860190925290912090611a01565b5080546001600160401b0316600081815260038501602052604090209092505b60108390558054600160c01b900460ff1615611a1c57600080fd5b8054600160d01b900460ff16611a3157600080fd5b600582015460ff16611a4257600080fd5b6005820154640100000000900460ff16611a5b57600080fd5b8154426001600160401b03600160c01b90920491909116600a011115611a8057600080fd5b6005820154610100900460ff1615611df8576013546009549095508510611aa657600080fd5b600060098681548110611ab557fe5b906000526020600020906006020190506000600b8460000160089054906101000a90046001600160401b03166001600160401b031681548110611af457fe5b600091825260209091206002909102018054909150600160881b90046001600160401b03168710801590611b35575060018101546001600160401b03168711155b611b3e57600080fd5b60018101546001600160401b0316871415611b995785546001600160401b031615801590611b80575085546011546000196001600160401b0392831601909116145b15611b8f57600f805460010190555b6011805460010190555b600187016013558154600160401b900460ff168015611bc157508154600160581b900460ff16155b15611d9c57604080516101608101825283546001600160401b0381168252600160401b810460ff9081161515602080850191909152600160481b83048216151584860152600160501b8304821615156060850152600160581b830490911615156080840152600160601b9091046001600160801b031660a08301526001808601546001600160a01b0390811660c085015260028088015490911660e085015260038701546101008086019190915260048801546101208601526005880180548751948116159092026000190190911691909104601f8101849004840283018401909552848252611d22948c9493889361014086019390929091830182828015611d0b5780601f10611ce057610100808354040283529160200191611d0b565b820191906000526020600020905b815481529060010190602001808311611cee57829003601f168201915b5050505050815250506130a290919063ffffffff16565b5060018201546040516001600160a01b03909116906000906512309ce540009082818181858883f19350505050158015611d60573d6000803e3d6000fd5b507f6940a01870e576ceb735867e13863646d517ce10e66c0133186a4ebdfe9388c2876001604051611d93929190614f73565b60405180910390a15b815460ff60501b1916600160501b1782556040517f134017cf3262b18f892ee95dde3b0aec9a80cc70a7c96f09c64bd237aceb047390611de0908990600190614f73565b60405180910390a1600197505050505050505061129f565b6012546008549095508510611e0c57600080fd5b8154600a8054600092600160401b90046001600160401b0316908110611e2e57fe5b600091825260209091206002909102018054909150600160881b90046001600160401b03168610801590611e6f575060018101546001600160401b03168611155b611e7857600080fd5b60018101546001600160401b0316861415611ed35784546001600160401b031615801590611eba575084546011546000196001600160401b0392831601909116145b15611ec957600f805460010190555b6011805460010190555b600060088781548110611ee257fe5b6000918252602090912060018901601255600690910201805460ff60501b1916600160501b17808255909150600160401b900460ff168015611f2d57508054600160581b900460ff16155b156120c657604080516101608101825282546001600160401b0381168252600160401b810460ff9081161515602080850191909152600160481b83048216151584860152600160501b8304821615156060850152600160581b830490911615156080840152600160601b9091046001600160801b031660a08301526001808501546001600160a01b0390811660c085015260028087015490911660e085015260038601546101008086019190915260048701546101208601526005870180548751948116159092026000190190911691909104601f810184900484028301840190955284825261204c948c9493879361014086019390929091830182828015611d0b5780601f10611ce057610100808354040283529160200191611d0b565b5060018101546040516001600160a01b03909116906000906509184e72a0009082818181858883f1935050505015801561208a573d6000803e3d6000fd5b507f6940a01870e576ceb735867e13863646d517ce10e66c0133186a4ebdfe9388c28760006040516120bd929190614f73565b60405180910390a15b7f134017cf3262b18f892ee95dde3b0aec9a80cc70a7c96f09c64bd237aceb04738760006040516120f8929190614f73565b60405180910390a1600197505050505050505090565b6000805461010090046001600160a01b0316331461212b57600080fd5b6509184e72a0008034101561213f57600080fd5b60005460ff166121535761215161271b565b505b61215f8686868661314f565b9695505050505050565b60045481565b601481565b6008818154811061218157fe5b6000918252602091829020600691909102018054600180830154600280850154600386015460048701546005880180546040805161010099831615999099026000190190911695909504601f81018b90048b0288018b019095528487526001600160401b0388169a50600160401b880460ff9081169a600160481b8a0482169a600160501b8b0483169a600160581b810490931699600160601b9093046001600160801b0316986001600160a01b039081169897169690939183018282801561228b5780601f106122605761010080835404028352916020019161228b565b820191906000526020600020905b81548152906001019060200180831161226e57829003601f168201915b505050505090508b565b60085490565b600d6020526000908152604090205481565b6003546001600160a01b031681565b600b81815481106122c957fe5b60009182526020909120600290910201805460019091015460ff821692506001600160401b036101008304811692600160481b8104821692600160881b9091048216918116906001600160a01b03600160401b9091041686565b603c81565b60135481565b60105481565b600e6020526000908152604090205481565b6000805461010090046001600160a01b0316331461236357600080fd5b612375836001600160a01b0316613229565b61237e57600080fd5b6001600160a01b0383811660009081526014602052604090205416156123a357600080fd5b6001600160a01b038381166000908152601460205260409081902080546001600160a01b03191692851692909217909155517fc5ec2ed49686197edd2ed642c7e6096893cc81e6658cde2527030316037715d0906124049085908590614d8b565b60405180910390a150600192915050565b600554141590565b606060006008838154811061242e57fe5b600091825260208083206002600690930201828101546001600160a01b039081168086526014845260408087205481516101608101835285546001600160401b038116825260ff600160401b820481161515838a0152600160481b82048116151583860152600160501b8204811615156060840152600160581b820416151560808301526001600160801b03600160601b9091041660a0820152600180870154861660c083015260e082019490945260038601546101008281019190915260048701546101208301526005870180548551601f60001998831615909402979097011699909904908101889004880285018801909352828452949850610b92976125c1978c9791966125b496939093169492938b936101408601939092919083018282801561259d5780601f106125725761010080835404028352916020019161259d565b820191906000526020600020905b81548152906001019060200180831161258057829003601f168201915b50505050508152505061322f90919063ffffffff16565b919063ffffffff6132fd16565b613355565b600090815260076020526040902060010154600160801b90046001600160401b031690565b6014602052600090815260409020546001600160a01b031681565b600081565b6002546001600160a01b031681565b6509184e72a000803410156103a257600080fd5b6001546001600160a01b031681565b60006509184e72a0008034101561265357600080fd5b600061266a6008600a886000898960016000612970565b90507f879922cf5fcada9ebaf8bd7424dc62877f4b220cae07fb6695cc1e8f94c52b4d8133886000898960016000604051611824989796959493929190614ebf565b600a5490565b600a81815481106122c957fe5b600081156126d957600983815481106126d457fe5b506000525b600883815481106126e657fe5b6000918252602090912060069091020154600160501b900460ff169392505050565b6009818154811061218157fe5b60115481565b600554600090815260076020526040812080546001600160401b03161561274657600091505061129f565b8054600180830154600092612778926001600160401b03600160c01b909204821692600160801b9004821601166134f8565b60018301549091506001600160401b031681111561279b5760009250505061129f565b6000818152600483016020526040902060058101546301000000900460ff16156127cb576000935050505061129f565b600581015460ff161561281b578054426001600160401b03600160801b90920491909116600f011115612804576000935050505061129f565b61280f83828461350f565b6001935050505061129f565b6001808401546001600160401b03600160401b9091048116909101165b6000818152600385016020526040902054600160d01b900460ff161561286057600101612838565b61286a84826135ce565b15612886576128798482613688565b600194505050505061129f565b600094505050505090565b60025460405160009182916060916001600160a01b0316906128b290614d51565b604051809103902089898989896040516024016128d3959493929190614fde565b60408051601f198184030181529181526020820180516001600160e01b03166001600160e01b03199094169390931790925290516129119190614d45565b600060405180830381855af49150503d806000811461294c576040519150601f19603f3d011682016040523d82523d6000602084013e612951565b606091505b50915091508161296057600080fd5b6001925050505b95945050505050565b60006104008451111561298257600080fd5b6003546001600160a01b0388811691161480801561299e575083155b806129c257506001600160a01b038881166000908152601460205260409020541615155b6129cb57600080fd5b89546129da8b60018301613fe0565b915060008a83815481106129ea57fe5b600091825260209182902060069190910201600181018054336001600160a01b0319918216179091556002820180549091166001600160a01b038d16179055805467ffffffffffffffff1916426001600160401b03161768ff00000000000000001916600160401b881515021769ff0000000000000000001916600160481b85151502176fffffffffffffffffffffffffffffffff60601b1916600160601b6001600160801b038c1602178155600381018990558751909250612ab591600584019190890190614011565b5084612bde57604080516101608101825282546001600160401b0381168252600160401b810460ff9081161515602080850191909152600160481b83048216151584860152600160501b8304821615156060850152600160581b830490911615156080840152600160601b9091046001600160801b031660a08301526001808501546001600160a01b0390811660c085015260028087015490911660e085015260038601546101008086019190915260048701546101208601526005870180548751948116159092026000190190911691909104601f8101849004840283018401909552848252612bd594889493879361014086019390929091830182828015611d0b5780601f10611ce057610100808354040283529160200191611d0b565b612bde57600080fd5b895460009015612bf3578a5460001901612c02565b8a54612c028c6001830161408f565b905060008b8281548110612c1257fe5b60009182526020909120600290910201805490915060ff1680612c5a5750612c3861309d565b81546001808401546001600160401b03600160881b9093048316908316030116145b15612cbd57805460ff1916600190811782558c548d91612c7d908390830161408f565b81548110612c8757fe5b60009182526020909120600290910201805467ffffffffffffffff60881b1916600160881b6001600160401b0388160217815590505b612cc681613745565b60018101805467ffffffffffffffff19166001600160401b03871617905586612d145780546001600160401b0361010080830482166001019091160268ffffffffffffffff00199091161781555b838015612d1f575086155b15612e5657604080516101608101825284546001600160401b0381168252600160401b810460ff9081161515602080850191909152600160481b83048216151584860152600160501b8304821615156060850152600160581b830490911615156080840152600160601b9091046001600160801b031660a08301526001808701546001600160a01b0390811660c085015260028089015490911660e085015260038801546101008086019190915260048901546101208601526005890180548751948116159092026000190190911691909104601f8101849004840283018401909552848252612e51948894612e42943394919387936101408601939283018282801561259d5780601f106125725761010080835404028352916020019161259d565b8391908863ffffffff61374816565b612f82565b6001600160a01b038b81166000908152601460209081526040918290205482516101608101845287546001600160401b0381168252600160401b810460ff908116151583860152600160481b82048116151583870152600160501b8204811615156060840152600160581b82041615156080830152600160601b90046001600160801b031660a0820152600180890154861660c08301526002808a0154871660e084015260038a01546101008085019190915260048b015461012085015260058b0180548851948116159092026000190190911691909104601f8101869004860283018601909652858252612f82968a96612e4296959091169487936101408601939092909183018282801561259d5780601f106125725761010080835404028352916020019161259d565b5050505098975050505050505050565b845460018601546001600160401b03600160881b909204821685019116811115612fbb57600080fd5b6000858281548110612fc957fe5b6000918252602090912089546006909202019150426001600160401b03600160c01b90920491909116600a0111612fff57600080fd5b6005880154640100000000900460ff1661301857600080fd5b8054600160581b900460ff161561302e57600080fd5b8054600160501b900460ff161561304457600080fd5b835160208501206130548561376a565b1561305e57600080fd5b60005460ff166130815761307881878b60040154876137a7565b61308157600080fd5b50805460ff60581b1916600160581b1790559695505050505050565b601490565b600062019a285a116130b357600080fd5b8260e001516001600160a01b031663a9f793088460200151848660c001518761010001518861014001516040518663ffffffff1660e01b81526004016130fd959493929190614db4565b602060405180830381600087803b15801561311757600080fd5b505af115801561312b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250610b9291908101906144f6565b60025460405160009182916060916001600160a01b03169061317090614d67565b60405180910390208888888860405160240161318f9493929190614f8e565b60408051601f198184030181529181526020820180516001600160e01b03166001600160e01b03199094169390931790925290516131cd9190614d45565b600060405180830381855af49150503d8060008114613208576040519150601f19603f3d011682016040523d82523d6000602084013e61320d565b606091505b50915091508161321c57600080fd5b5060019695505050505050565b3b151590565b6132376140bb565b6020808401805115159183019190915260408085015115159083015260c0808501516001600160a01b03169083015251158015613275575082604001515b156132bf5760c08301516001600160a01b031660e082015261014083015160208101516001600160801b0381166132ab57600080fd5b6001600160801b031660a083015250610efb565b6001600160a01b03821660e082015260a0808401516001600160801b0316908201526101008084015190820152610140808401519082015292915050565b613305614119565b633b9aca006020820152620186a0604082015260e08401516001600160a01b0316606082015260a08401516001600160801b03166080820152613349848484613872565b60a08201529392505050565b6040805160098082526101408201909252606091829190816020015b6060815260200190600190039081613371575050835190915061339c906001600160401b0316613948565b816000815181106133a957fe5b60200260200101819052506133c18360200151613948565b816001815181106133ce57fe5b60200260200101819052506133ef83604001516001600160401b0316613948565b816002815181106133fc57fe5b602002602001018190525061341d83606001516001600160a01b031661395b565b8160038151811061342a57fe5b60200260200101819052506134428360800151613948565b8160048151811061344f57fe5b60200260200101819052506134678360a0015161397a565b8160058151811061347457fe5b602002602001018190525061348c8360c00151613948565b8160068151811061349957fe5b60200260200101819052506134b18360e00151613948565b816007815181106134be57fe5b60200260200101819052506134d7836101000151613948565b816008815181106134e457fe5b6020026020010181905250610b92816139d0565b6000818310156135085781610b92565b5090919050565b6005828101805464010000000064ff000000001990911617905582546001600160c01b0316600160c01b426001600160401b039081169190910291909117845560018501805467ffffffffffffffff60801b1916600160801b8584160217808255855467ffffffffffffffff60401b199091169216600160401b02919091179055546040517ffb96205e4b3633fd57aa805b26b51ecf528714a10241a4af015929dce86768d9916135c1918490614fc3565b60405180910390a1505050565b8154600090600160801b90046001600160401b03168211156135f257506000610efb565b600082815260038401602052604090208054600160c81b900460ff1661361c576000915050610efb565b8054600160d01b900460ff1615613637576000915050610efb565b6006810154600160881b900460ff168061365c57506006810154600160901b900460ff165b1561366b576000915050610efb565b60060154426001600160401b039091166014011115905092915050565b600081815260038301602052604090819020600681018054600160801b60ff60801b19909116811767ffffffffffffffff60401b19908116600160401b426001600160401b03908116820292909217909455845460018901805467ffffffffffffffff60801b1916918690048316909402179091168187168402179091556005548354945193947f70801d4d63b3da6c19ba7349911f45bed5a99ccdfb51b8138c105872529bebd5946135c1949293889383821693041690615020565b50565b61375c613757838360006132fd565b613a2d565b836004018190555050505050565b60006060613788600461377c85613a49565b9063ffffffff613a6c16565b9050610b928160008151811061379a57fe5b6020026020010151613af9565b600060208251816137b457fe5b06156137bf57600080fd5b600060208351816137cc57fe5b049050601081106137dc57600080fd5b60008660205b836020028111613864578581015192506002880661382a57818360405160200161380d929190614d03565b604051602081830303815290604052805190602001209150613856565b828260405160200161383d929190614d03565b6040516020818303038152906040528051906020012091505b6002880497506020016137e2565b509094149695505050505050565b60608360400151801561388757508360200151155b1561389157610b92565b6000826138b1576040516138a490614d5c565b60405180910390206138c6565b6040516138bd90614d72565b60405180910390205b90508085602001516138d95760006138dc565b60015b60ff1660001b858760c001516001600160a01b0316886101000151896101400151604051602001613911959493929190614e4a565b60408051601f198184030181529082905261392f9291602001614d29565b6040516020818303038152906040529150509392505050565b6060610efb61395683613b1d565b61397a565b60408051600560a21b8318601482015260348101909152606090610b92815b6060815160011480156139ac5750607f60f81b8260008151811061399a57fe5b01602001516001600160f81b03191611155b156139b8575080610a0a565b610efb6139ca8351608060ff16613c16565b83613cdf565b604080516000808252602082019092526060915b8351811015613a1457613a0a828583815181106139fd57fe5b6020026020010151613cdf565b91506001016139e4565b50610b92613a27825160c060ff16613c16565b82613cdf565b60006060613a3a83613355565b80516020909101209392505050565b613a51614180565b50805160408051808201909152602092830181529182015290565b606081604051908082528060200260200182016040528015613aa857816020015b613a95614180565b815260200190600190039081613a8d5790505b509050613ab361419a565b613abc84613dc4565b905060005b83811015613af157613ad282613de9565b838281518110613ade57fe5b6020908102919091010152600101613ac1565b505092915050565b6000806000613b0784613e1b565b90516020919091036101000a9004949350505050565b6040805160208082528183019092526060916000918391602082018180388339019050509050836020820152600091505b6020821015613b8757808281518110613b6357fe5b01602001516001600160f81b03191615613b7c57613b87565b600190910190613b4e565b6060826020036040519080825280601f01601f191660200182016040528015613bb7576020820181803883390190505b50905060005b8151811015613c0d578251600185019484918110613bd757fe5b602001015160f81c60f81b828281518110613bee57fe5b60200101906001600160f81b031916908160001a905350600101613bbd565b50949350505050565b6060600160401b8310613c445760405162461bcd60e51b8152600401613c3b90614e83565b60405180910390fd5b60408051600180825281830190925260609160208201818038833901905050905060378411613c9e5782840160f81b81600081518110613c8057fe5b60200101906001600160f81b031916908160001a9053509050610efb565b6060613ca985613b1d565b90508381510160370160f81b82600081518110613cc257fe5b60200101906001600160f81b031916908160001a90535061296782825b60608082518451016040519080825280601f01601f191660200182016040528015613d11576020820181803883390190505b5090506000805b8551811015613d6757858181518110613d2d57fe5b602001015160f81c60f81b838381518110613d4457fe5b60200101906001600160f81b031916908160001a90535060019182019101613d18565b5060005b8451811015613dba57848181518110613d8057fe5b602001015160f81c60f81b838381518110613d9757fe5b60200101906001600160f81b031916908160001a90535060019182019101613d6b565b5090949350505050565b613dcc61419a565b6000613dd783613e7a565b83519383529092016020820152919050565b613df1614180565b60208201516000613e0182613ee4565b828452602080850182905292019390910192909252919050565b805180516000918291821a906080821015613e3d57925060019150613e759050565b60b8821015613e5b5760018560200151039250806001019350613e72565b602085015182820160b51901945082900360b60192505b50505b915091565b8051805160009190821a906080821015613e9957600092505050610a0a565b60b8821080613eb4575060c08210158015613eb4575060f882105b15613ec457600192505050610a0a565b60c0821015613ed9575060b519019050610a0a565b5060f5190192915050565b8051600090811a6080811015613efd5760019150613f0e565b60b8811015613f0e57607e19810191505b50919050565b604080516102c081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c0810182905260e08101919091526101008101613f626141ba565b8152602001613f6f6141ef565b905290565b604080516101a081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c0810182905260e08101829052610100810182905261012081018290526101408101829052610160810182905261018081019190915290565b81548183558181111561400c5760060281600602836000526020600020918201910161400c9190614233565b505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061405257805160ff191683800117855561407f565b8280016001018555821561407f579182015b8281111561407f578251825591602001919060010190614064565b5061408b929150614295565b5090565b81548183558181111561400c5760020281600202836000526020600020918201910161400c91906142af565b604080516101608101825260008082526020820181905291810182905260608082018390526080820183905260a0820183905260c0820183905260e08201839052610100820183905261012082019290925261014081019190915290565b60405180610120016040528060006001600160401b031681526020016000815260200160006001600160401b0316815260200160006001600160a01b0316815260200160008152602001606081526020016000815260200160008152602001600081525090565b604051806040016040528060008152602001600081525090565b60405180606001604052806141ad614180565b8152602001600081525090565b6040805160c081018252600080825260208201819052918101829052606081018290526080810182905260a081019190915290565b6040805161010081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c0810182905260e081019190915290565b61129f91905b8082111561408b5780546001600160e01b03191681556001810180546001600160a01b031990811690915560028201805490911690556000600382018190556004820181905561428c60058301826142e5565b50600601614239565b61129f91905b8082111561408b576000815560010161429b565b61129f91905b8082111561408b5780546001600160c81b03191681556001810180546001600160e01b03191690556002016142b5565b50805460018160011615610100020316600290046000825580601f1061430b5750613745565b601f0160209004906000526020600020908101906137459190614295565b8035610efb816152aa565b60008083601f84011261434657600080fd5b5081356001600160401b0381111561435d57600080fd5b60208301915083602082028301111561437557600080fd5b9250929050565b8035610efb816152be565b8051610efb816152be565b8035610efb816152c7565b60008083601f8401126143af57600080fd5b5081356001600160401b038111156143c657600080fd5b60208301915083600182028301111561437557600080fd5b600082601f8301126143ef57600080fd5b81356144026143fd826151c7565b6151a1565b9150808252602083016020830185838301111561441e57600080fd5b614429838284615268565b50505092915050565b60006020828403121561444457600080fd5b60006144508484614329565b949350505050565b6000806040838503121561446b57600080fd5b60006144778585614329565b925050602061448885828601614329565b9150509250929050565b6000806000606084860312156144a757600080fd5b60006144b38686614329565b93505060206144c486828701614392565b92505060408401356001600160401b038111156144e057600080fd5b6144ec868287016143de565b9150509250925092565b60006020828403121561450857600080fd5b60006144508484614387565b60006020828403121561452657600080fd5b60006144508484614392565b6000806040838503121561454557600080fd5b60006145518585614392565b92505060206144888582860161437c565b6000806000806080858703121561457857600080fd5b60006145848787614392565b945050602061459587828801614392565b93505060406145a687828801614392565b92505060606145b787828801614392565b91505092959194509250565b60008060008060008060008060a0898b0312156145df57600080fd5b60006145eb8b8b614392565b98505060208901356001600160401b0381111561460757600080fd5b6146138b828c0161439d565b975097505060408901356001600160401b0381111561463157600080fd5b61463d8b828c0161439d565b955095505060606146508b828c01614392565b93505060808901356001600160401b0381111561466c57600080fd5b6146788b828c01614334565b92509250509295985092959890939650565b6000806040838503121561469d57600080fd5b60006146a98585614392565b925050602061448885828601614392565b600080600080600060a086880312156146d257600080fd5b60006146de8888614392565b95505060206146ef88828901614392565b945050604061470088828901614392565b935050606061471188828901614392565b925050608061472288828901614392565b9150509295509295909350565b600080600080600080600060a0888a03121561474a57600080fd5b60006147568a8a614392565b97505060206147678a828b01614392565b96505060406147788a828b01614392565b95505060608801356001600160401b0381111561479457600080fd5b6147a08a828b0161439d565b945094505060808801356001600160401b038111156147be57600080fd5b6147ca8a828b0161439d565b925092505092959891949750929550565b6147e48161523c565b82525050565b6147e4816151fb565b6147e481615206565b6147e48161129f565b6147e46148118261129f565b61129f565b6147e46148118261520b565b600061482d826151ee565b61483781856151f2565b9350614847818560208601615274565b614850816152a0565b9093019392505050565b6000614865826151ee565b61486f8185610a0a565b935061487f818560208601615274565b9290920192915050565b6147e481615247565b600061489f603283610a0a565b7f7375626d69744e52452875696e743235362c75696e743235362c627974657333815271322c627974657333322c627974657333322960701b602082015260320192915050565b60006148f3603c83610a0a565b7f6170706c7952657175657374496e4368696c64436861696e28626f6f6c2c756981527f6e743235362c616464726573732c627974657333322c627974657329000000006020820152603c0192915050565b6000614952600e836151f2565b6d696e70757420746f6f206c6f6e6760901b815260200192915050565b600061497c602a83610a0a565b7f7375626d69744f52422875696e743235362c627974657333322c627974657333815269322c627974657333322960b01b6020820152602a0192915050565b60006149c8603b83610a0a565b7f6170706c7952657175657374496e526f6f74436861696e28626f6f6c2c75696e81527f743235362c616464726573732c627974657333322c62797465732900000000006020820152603b0192915050565b80516102c0830190614a2c8482614cfa565b506020820151614a3f6020850182614cfa565b506040820151614a526040850182614cfa565b506060820151614a6560608501826147f3565b506080820151614a7860808501826147f3565b5060a0820151614a8b60a08501826147f3565b5060c0820151614a9e60c08501826147f3565b5060e0820151614ab160e08501826147f3565b50610100820151614ac6610100850182614c78565b50610120820151614adb6101c0850182614ae1565b50505050565b8051610100830190614af384826147fc565b506020820151614b0660208501826147fc565b506040820151614b1960408501826147fc565b506060820151614b2c6060850182614cfa565b506080820151614b3f6080850182614cfa565b5060a0820151614b5260a08501826147f3565b5060c0820151614b6560c08501826147f3565b5060e0820151614adb60e08501826147f3565b80516101a0830190614b8a8482614cfa565b506020820151614b9d6020850182614cfa565b506040820151614bb06040850182614cfa565b506060820151614bc36060850182614cfa565b506080820151614bd66080850182614cfa565b5060a0820151614be960a08501826147fc565b5060c0820151614bfc60c08501826147fc565b5060e0820151614c0f60e08501826147fc565b50610100820151614c246101008501826147f3565b50610120820151614c396101208501826147f3565b50610140820151614c4e6101408501826147f3565b50610160820151614c636101608501826147f3565b50610180820151614adb6101808501826147f3565b805160c0830190614c898482614cfa565b506020820151614c9c6020850182614cfa565b506040820151614caf6040850182614cfa565b506060820151614cc26060850182614cfa565b506080820151614cd56080850182614cfa565b5060a0820151614adb60a0850182614cfa565b6147e481615218565b6147e48161525d565b6147e481615230565b6000614d0f8285614805565b602082019150614d1f8284614805565b5060200192915050565b6000614d358285614816565b600482019150614450828461485a565b6000610b92828461485a565b6000610efb82614892565b6000610efb826148e6565b6000610efb8261496f565b6000610efb826149bb565b60208101610efb82846147ea565b60408101614d9982856147ea565b610b9260208301846147ea565b60208101610efb82846147f3565b60a08101614dc282886147f3565b614dcf60208301876147fc565b614ddc60408301866147db565b614de960608301856147fc565b8181036080830152610a708184614822565b60c08101614e0982896147f3565b614e166020830188614cfa565b614e236040830187614cfa565b614e306060830186614cfa565b614e3d6080830185614cfa565b610a7060a08301846147ea565b60a08101614e5882886147fc565b614e6560208301876147fc565b614ddc60408301866147fc565b60208082528101610b928184614822565b60208082528101610efb81614945565b6102c08101610efb8284614a1a565b6101a08101610efb8284614b78565b60208101610efb82846147fc565b6101008101614ece828b6147fc565b614edb602083018a6147db565b614ee860408301896147ea565b614ef56060830188614889565b614f0260808301876147fc565b81810360a0830152614f148186614822565b9050614f2360c08301856147f3565b614f3060e08301846147f3565b9998505050505050505050565b6101008101614f4c828b6147fc565b614f59602083018a6147db565b614f6660408301896147ea565b614ef560608301886147fc565b60408101614f8182856147fc565b610b9260208301846147f3565b60808101614f9c82876147fc565b614fa960208301866147fc565b614fb660408301856147fc565b61296760608301846147fc565b60408101614fd182856147fc565b610b9260208301846147fc565b60a08101614fec82886147fc565b614ff960208301876147fc565b61500660408301866147fc565b61501360608301856147fc565b61215f60808301846147fc565b6080810161502e82876147fc565b61503b60208301866147fc565b6150486040830185614cf1565b6129676060830184614cf1565b6101608101615064828e614cfa565b615071602083018d6147f3565b61507e604083018c6147f3565b61508b606083018b6147f3565b615098608083018a6147f3565b6150a560a0830189614ce8565b6150b260c08301886147ea565b6150bf60e08301876147ea565b6150cd6101008301866147fc565b6150db6101208301856147fc565b8181036101408301526150ee8184614822565b9d9c50505050505050505050505050565b610180810161510e828f614cfa565b61511b602083018e614cfa565b615128604083018d614cfa565b615135606083018c614cfa565b615142608083018b614cfa565b61514f60a083018a614cfa565b61515c60c0830189614cfa565b61516960e0830188614cfa565b615177610100830187614cfa565b615185610120830186614cfa565b615193610140830185614cfa565b6150ee6101608301846147f3565b6040518181016001600160401b03811182821017156151bf57600080fd5b604052919050565b60006001600160401b038211156151dd57600080fd5b506020601f91909101601f19160190565b5190565b90815260200190565b6000610efb82615224565b151590565b6001600160e01b03191690565b6001600160801b031690565b6001600160a01b031690565b6001600160401b031690565b6000610efb82615252565b6000610efb8261129f565b6000610efb826151fb565b6000610efb82615230565b82818337506000910152565b60005b8381101561528f578181015183820152602001615277565b83811115614adb5750506000910152565b601f01601f191690565b6152b3816151fb565b811461374557600080fd5b6152b381615206565b6152b38161129f56fea365627a7a72305820dceaaa144785d50c1e631881b9d1c666eb36d5a371014e408dd638a76cc0ae3d6c6578706572696d656e74616cf564736f6c63430005090040"

// DeployRootChain deploys a new Ethereum contract, binding an instance of RootChain to it.
func DeployRootChain(auth *bind.TransactOpts, backend bind.ContractBackend, _epochHandler common.Address, _submitHandler common.Address, _etherToken common.Address, _development bool, _NRELength *big.Int, _statesRoot [32]byte, _transactionsRoot [32]byte, _receiptsRoot [32]byte) (common.Address, *types.Transaction, *RootChain, error) {
	parsed, err := abi.JSON(strings.NewReader(RootChainABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RootChainBin), backend, _epochHandler, _submitHandler, _etherToken, _development, _NRELength, _statesRoot, _transactionsRoot, _receiptsRoot)
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

// Struct2 is an auto generated low-level Go binding around an user-defined struct.
type Struct2 struct {
	EpochStateRoot        [32]byte
	EpochTransactionsRoot [32]byte
	EpochReceiptsRoot     [32]byte
	SubmittedAt           uint64
	FinalizedAt           uint64
	Finalized             bool
	Challenging           bool
	Challenged            bool
}

// Struct3 is an auto generated low-level Go binding around an user-defined struct.
type Struct3 struct {
	StartBlockNumber uint64
	EndBlockNumber   uint64
	Timestamp        uint64
	IsEmpty          bool
	Initialized      bool
	IsRequest        bool
	UserActivated    bool
	Rebase           bool
	RE               Struct1
	NRE              Struct2
}

// Struct0 is an auto generated low-level Go binding around an user-defined struct.
type Struct0 struct {
	EpochNumber      uint64
	RequestBlockId   uint64
	Timestamp        uint64
	FinalizedAt      uint64
	ReferenceBlock   uint64
	StatesRoot       [32]byte
	TransactionsRoot [32]byte
	ReceiptsRoot     [32]byte
	IsRequest        bool
	UserActivated    bool
	Challenged       bool
	Challenging      bool
	Finalized        bool
}

// Struct1 is an auto generated low-level Go binding around an user-defined struct.
type Struct1 struct {
	RequestStart        uint64
	RequestEnd          uint64
	FirstRequestBlockId uint64
	NumEnter            uint64
	NextEnterEpoch      uint64
	NextEpoch           uint64
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

// CPEXIT is a free data retrieval call binding the contract method 0x8155717d.
//
// Solidity: function CP_EXIT() constant returns(uint256)
func (_RootChain *RootChainCaller) CPEXIT(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChain.contract.Call(opts, out, "CP_EXIT")
	return *ret0, err
}

// CPEXIT is a free data retrieval call binding the contract method 0x8155717d.
//
// Solidity: function CP_EXIT() constant returns(uint256)
func (_RootChain *RootChainSession) CPEXIT() (*big.Int, error) {
	return _RootChain.Contract.CPEXIT(&_RootChain.CallOpts)
}

// CPEXIT is a free data retrieval call binding the contract method 0x8155717d.
//
// Solidity: function CP_EXIT() constant returns(uint256)
func (_RootChain *RootChainCallerSession) CPEXIT() (*big.Int, error) {
	return _RootChain.Contract.CPEXIT(&_RootChain.CallOpts)
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

// EROIdToFinalize is a free data retrieval call binding the contract method 0x2dc6bb7b.
//
// Solidity: function EROIdToFinalize() constant returns(uint256)
func (_RootChain *RootChainCaller) EROIdToFinalize(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChain.contract.Call(opts, out, "EROIdToFinalize")
	return *ret0, err
}

// EROIdToFinalize is a free data retrieval call binding the contract method 0x2dc6bb7b.
//
// Solidity: function EROIdToFinalize() constant returns(uint256)
func (_RootChain *RootChainSession) EROIdToFinalize() (*big.Int, error) {
	return _RootChain.Contract.EROIdToFinalize(&_RootChain.CallOpts)
}

// EROIdToFinalize is a free data retrieval call binding the contract method 0x2dc6bb7b.
//
// Solidity: function EROIdToFinalize() constant returns(uint256)
func (_RootChain *RootChainCallerSession) EROIdToFinalize() (*big.Int, error) {
	return _RootChain.Contract.EROIdToFinalize(&_RootChain.CallOpts)
}

// EROs is a free data retrieval call binding the contract method 0xb443f3cc.
//
// Solidity: function EROs(uint256 ) constant returns(uint64 timestamp, bool isExit, bool isTransfer, bool finalized, bool challenged, uint128 value, address requestor, address to, bytes32 trieKey, bytes32 hash, bytes trieValue)
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
	Hash       [32]byte
	TrieValue  []byte
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
		Hash       [32]byte
		TrieValue  []byte
	})
	out := ret
	err := _RootChain.contract.Call(opts, out, "EROs", arg0)
	return *ret, err
}

// EROs is a free data retrieval call binding the contract method 0xb443f3cc.
//
// Solidity: function EROs(uint256 ) constant returns(uint64 timestamp, bool isExit, bool isTransfer, bool finalized, bool challenged, uint128 value, address requestor, address to, bytes32 trieKey, bytes32 hash, bytes trieValue)
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
	Hash       [32]byte
	TrieValue  []byte
}, error) {
	return _RootChain.Contract.EROs(&_RootChain.CallOpts, arg0)
}

// EROs is a free data retrieval call binding the contract method 0xb443f3cc.
//
// Solidity: function EROs(uint256 ) constant returns(uint64 timestamp, bool isExit, bool isTransfer, bool finalized, bool challenged, uint128 value, address requestor, address to, bytes32 trieKey, bytes32 hash, bytes trieValue)
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
	Hash       [32]byte
	TrieValue  []byte
}, error) {
	return _RootChain.Contract.EROs(&_RootChain.CallOpts, arg0)
}

// ERUIdToFinalize is a free data retrieval call binding the contract method 0xc54626cc.
//
// Solidity: function ERUIdToFinalize() constant returns(uint256)
func (_RootChain *RootChainCaller) ERUIdToFinalize(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChain.contract.Call(opts, out, "ERUIdToFinalize")
	return *ret0, err
}

// ERUIdToFinalize is a free data retrieval call binding the contract method 0xc54626cc.
//
// Solidity: function ERUIdToFinalize() constant returns(uint256)
func (_RootChain *RootChainSession) ERUIdToFinalize() (*big.Int, error) {
	return _RootChain.Contract.ERUIdToFinalize(&_RootChain.CallOpts)
}

// ERUIdToFinalize is a free data retrieval call binding the contract method 0xc54626cc.
//
// Solidity: function ERUIdToFinalize() constant returns(uint256)
func (_RootChain *RootChainCallerSession) ERUIdToFinalize() (*big.Int, error) {
	return _RootChain.Contract.ERUIdToFinalize(&_RootChain.CallOpts)
}

// ERUs is a free data retrieval call binding the contract method 0xf4f31de4.
//
// Solidity: function ERUs(uint256 ) constant returns(uint64 timestamp, bool isExit, bool isTransfer, bool finalized, bool challenged, uint128 value, address requestor, address to, bytes32 trieKey, bytes32 hash, bytes trieValue)
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
	Hash       [32]byte
	TrieValue  []byte
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
		Hash       [32]byte
		TrieValue  []byte
	})
	out := ret
	err := _RootChain.contract.Call(opts, out, "ERUs", arg0)
	return *ret, err
}

// ERUs is a free data retrieval call binding the contract method 0xf4f31de4.
//
// Solidity: function ERUs(uint256 ) constant returns(uint64 timestamp, bool isExit, bool isTransfer, bool finalized, bool challenged, uint128 value, address requestor, address to, bytes32 trieKey, bytes32 hash, bytes trieValue)
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
	Hash       [32]byte
	TrieValue  []byte
}, error) {
	return _RootChain.Contract.ERUs(&_RootChain.CallOpts, arg0)
}

// ERUs is a free data retrieval call binding the contract method 0xf4f31de4.
//
// Solidity: function ERUs(uint256 ) constant returns(uint64 timestamp, bool isExit, bool isTransfer, bool finalized, bool challenged, uint128 value, address requestor, address to, bytes32 trieKey, bytes32 hash, bytes trieValue)
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
	Hash       [32]byte
	TrieValue  []byte
}, error) {
	return _RootChain.Contract.ERUs(&_RootChain.CallOpts, arg0)
}

// MAXREQUESTS is a free data retrieval call binding the contract method 0x93521222.
//
// Solidity: function MAX_REQUESTS() constant returns(uint256 maxRequests)
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
// Solidity: function MAX_REQUESTS() constant returns(uint256 maxRequests)
func (_RootChain *RootChainSession) MAXREQUESTS() (*big.Int, error) {
	return _RootChain.Contract.MAXREQUESTS(&_RootChain.CallOpts)
}

// MAXREQUESTS is a free data retrieval call binding the contract method 0x93521222.
//
// Solidity: function MAX_REQUESTS() constant returns(uint256 maxRequests)
func (_RootChain *RootChainCallerSession) MAXREQUESTS() (*big.Int, error) {
	return _RootChain.Contract.MAXREQUESTS(&_RootChain.CallOpts)
}

// NRELength is a free data retrieval call binding the contract method 0xab96da2d.
//
// Solidity: function NRELength() constant returns(uint256)
func (_RootChain *RootChainCaller) NRELength(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChain.contract.Call(opts, out, "NRELength")
	return *ret0, err
}

// NRELength is a free data retrieval call binding the contract method 0xab96da2d.
//
// Solidity: function NRELength() constant returns(uint256)
func (_RootChain *RootChainSession) NRELength() (*big.Int, error) {
	return _RootChain.Contract.NRELength(&_RootChain.CallOpts)
}

// NRELength is a free data retrieval call binding the contract method 0xab96da2d.
//
// Solidity: function NRELength() constant returns(uint256)
func (_RootChain *RootChainCallerSession) NRELength() (*big.Int, error) {
	return _RootChain.Contract.NRELength(&_RootChain.CallOpts)
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
// Solidity: function ORBs(uint256 ) constant returns(bool submitted, uint64 numEnter, uint64 epochNumber, uint64 requestStart, uint64 requestEnd, address trie)
func (_RootChain *RootChainCaller) ORBs(opts *bind.CallOpts, arg0 *big.Int) (struct {
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
	err := _RootChain.contract.Call(opts, out, "ORBs", arg0)
	return *ret, err
}

// ORBs is a free data retrieval call binding the contract method 0xea7f22a8.
//
// Solidity: function ORBs(uint256 ) constant returns(bool submitted, uint64 numEnter, uint64 epochNumber, uint64 requestStart, uint64 requestEnd, address trie)
func (_RootChain *RootChainSession) ORBs(arg0 *big.Int) (struct {
	Submitted    bool
	NumEnter     uint64
	EpochNumber  uint64
	RequestStart uint64
	RequestEnd   uint64
	Trie         common.Address
}, error) {
	return _RootChain.Contract.ORBs(&_RootChain.CallOpts, arg0)
}

// ORBs is a free data retrieval call binding the contract method 0xea7f22a8.
//
// Solidity: function ORBs(uint256 ) constant returns(bool submitted, uint64 numEnter, uint64 epochNumber, uint64 requestStart, uint64 requestEnd, address trie)
func (_RootChain *RootChainCallerSession) ORBs(arg0 *big.Int) (struct {
	Submitted    bool
	NumEnter     uint64
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
// Solidity: function URBs(uint256 ) constant returns(bool submitted, uint64 numEnter, uint64 epochNumber, uint64 requestStart, uint64 requestEnd, address trie)
func (_RootChain *RootChainCaller) URBs(opts *bind.CallOpts, arg0 *big.Int) (struct {
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
	err := _RootChain.contract.Call(opts, out, "URBs", arg0)
	return *ret, err
}

// URBs is a free data retrieval call binding the contract method 0xc0e86064.
//
// Solidity: function URBs(uint256 ) constant returns(bool submitted, uint64 numEnter, uint64 epochNumber, uint64 requestStart, uint64 requestEnd, address trie)
func (_RootChain *RootChainSession) URBs(arg0 *big.Int) (struct {
	Submitted    bool
	NumEnter     uint64
	EpochNumber  uint64
	RequestStart uint64
	RequestEnd   uint64
	Trie         common.Address
}, error) {
	return _RootChain.Contract.URBs(&_RootChain.CallOpts, arg0)
}

// URBs is a free data retrieval call binding the contract method 0xc0e86064.
//
// Solidity: function URBs(uint256 ) constant returns(bool submitted, uint64 numEnter, uint64 epochNumber, uint64 requestStart, uint64 requestEnd, address trie)
func (_RootChain *RootChainCallerSession) URBs(arg0 *big.Int) (struct {
	Submitted    bool
	NumEnter     uint64
	EpochNumber  uint64
	RequestStart uint64
	RequestEnd   uint64
	Trie         common.Address
}, error) {
	return _RootChain.Contract.URBs(&_RootChain.CallOpts, arg0)
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

// EpochHandler is a free data retrieval call binding the contract method 0xe7b88b80.
//
// Solidity: function epochHandler() constant returns(address)
func (_RootChain *RootChainCaller) EpochHandler(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RootChain.contract.Call(opts, out, "epochHandler")
	return *ret0, err
}

// EpochHandler is a free data retrieval call binding the contract method 0xe7b88b80.
//
// Solidity: function epochHandler() constant returns(address)
func (_RootChain *RootChainSession) EpochHandler() (common.Address, error) {
	return _RootChain.Contract.EpochHandler(&_RootChain.CallOpts)
}

// EpochHandler is a free data retrieval call binding the contract method 0xe7b88b80.
//
// Solidity: function epochHandler() constant returns(address)
func (_RootChain *RootChainCallerSession) EpochHandler() (common.Address, error) {
	return _RootChain.Contract.EpochHandler(&_RootChain.CallOpts)
}

// EtherToken is a free data retrieval call binding the contract method 0xb8066bcb.
//
// Solidity: function etherToken() constant returns(address)
func (_RootChain *RootChainCaller) EtherToken(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RootChain.contract.Call(opts, out, "etherToken")
	return *ret0, err
}

// EtherToken is a free data retrieval call binding the contract method 0xb8066bcb.
//
// Solidity: function etherToken() constant returns(address)
func (_RootChain *RootChainSession) EtherToken() (common.Address, error) {
	return _RootChain.Contract.EtherToken(&_RootChain.CallOpts)
}

// EtherToken is a free data retrieval call binding the contract method 0xb8066bcb.
//
// Solidity: function etherToken() constant returns(address)
func (_RootChain *RootChainCallerSession) EtherToken() (common.Address, error) {
	return _RootChain.Contract.EtherToken(&_RootChain.CallOpts)
}

// FirstFilledORENumber is a free data retrieval call binding the contract method 0x72ecb9a8.
//
// Solidity: function firstFilledORENumber(uint256 ) constant returns(uint256)
func (_RootChain *RootChainCaller) FirstFilledORENumber(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChain.contract.Call(opts, out, "firstFilledORENumber", arg0)
	return *ret0, err
}

// FirstFilledORENumber is a free data retrieval call binding the contract method 0x72ecb9a8.
//
// Solidity: function firstFilledORENumber(uint256 ) constant returns(uint256)
func (_RootChain *RootChainSession) FirstFilledORENumber(arg0 *big.Int) (*big.Int, error) {
	return _RootChain.Contract.FirstFilledORENumber(&_RootChain.CallOpts, arg0)
}

// FirstFilledORENumber is a free data retrieval call binding the contract method 0x72ecb9a8.
//
// Solidity: function firstFilledORENumber(uint256 ) constant returns(uint256)
func (_RootChain *RootChainCallerSession) FirstFilledORENumber(arg0 *big.Int) (*big.Int, error) {
	return _RootChain.Contract.FirstFilledORENumber(&_RootChain.CallOpts, arg0)
}

// FirstNonEmptyRequestEpoch is a free data retrieval call binding the contract method 0xca6f6380.
//
// Solidity: function firstNonEmptyRequestEpoch(uint256 ) constant returns(uint256)
func (_RootChain *RootChainCaller) FirstNonEmptyRequestEpoch(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChain.contract.Call(opts, out, "firstNonEmptyRequestEpoch", arg0)
	return *ret0, err
}

// FirstNonEmptyRequestEpoch is a free data retrieval call binding the contract method 0xca6f6380.
//
// Solidity: function firstNonEmptyRequestEpoch(uint256 ) constant returns(uint256)
func (_RootChain *RootChainSession) FirstNonEmptyRequestEpoch(arg0 *big.Int) (*big.Int, error) {
	return _RootChain.Contract.FirstNonEmptyRequestEpoch(&_RootChain.CallOpts, arg0)
}

// FirstNonEmptyRequestEpoch is a free data retrieval call binding the contract method 0xca6f6380.
//
// Solidity: function firstNonEmptyRequestEpoch(uint256 ) constant returns(uint256)
func (_RootChain *RootChainCallerSession) FirstNonEmptyRequestEpoch(arg0 *big.Int) (*big.Int, error) {
	return _RootChain.Contract.FirstNonEmptyRequestEpoch(&_RootChain.CallOpts, arg0)
}

// Forked is a free data retrieval call binding the contract method 0xce8a2bc2.
//
// Solidity: function forked(uint256 _forkNumber) constant returns(bool result)
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
// Solidity: function forked(uint256 _forkNumber) constant returns(bool result)
func (_RootChain *RootChainSession) Forked(_forkNumber *big.Int) (bool, error) {
	return _RootChain.Contract.Forked(&_RootChain.CallOpts, _forkNumber)
}

// Forked is a free data retrieval call binding the contract method 0xce8a2bc2.
//
// Solidity: function forked(uint256 _forkNumber) constant returns(bool result)
func (_RootChain *RootChainCallerSession) Forked(_forkNumber *big.Int) (bool, error) {
	return _RootChain.Contract.Forked(&_RootChain.CallOpts, _forkNumber)
}

// Forks is a free data retrieval call binding the contract method 0x4ba3a126.
//
// Solidity: function forks(uint256 ) constant returns(uint64 forkedBlock, uint64 firstEpoch, uint64 lastEpoch, uint64 firstBlock, uint64 lastBlock, uint64 lastFinalizedEpoch, uint64 lastFinalizedBlock, uint64 timestamp, uint64 firstEnterEpoch, uint64 lastEnterEpoch, uint64 nextBlockToRebase, bool rebased)
func (_RootChain *RootChainCaller) Forks(opts *bind.CallOpts, arg0 *big.Int) (struct {
	ForkedBlock        uint64
	FirstEpoch         uint64
	LastEpoch          uint64
	FirstBlock         uint64
	LastBlock          uint64
	LastFinalizedEpoch uint64
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
		LastFinalizedEpoch uint64
		LastFinalizedBlock uint64
		Timestamp          uint64
		FirstEnterEpoch    uint64
		LastEnterEpoch     uint64
		NextBlockToRebase  uint64
		Rebased            bool
	})
	out := ret
	err := _RootChain.contract.Call(opts, out, "forks", arg0)
	return *ret, err
}

// Forks is a free data retrieval call binding the contract method 0x4ba3a126.
//
// Solidity: function forks(uint256 ) constant returns(uint64 forkedBlock, uint64 firstEpoch, uint64 lastEpoch, uint64 firstBlock, uint64 lastBlock, uint64 lastFinalizedEpoch, uint64 lastFinalizedBlock, uint64 timestamp, uint64 firstEnterEpoch, uint64 lastEnterEpoch, uint64 nextBlockToRebase, bool rebased)
func (_RootChain *RootChainSession) Forks(arg0 *big.Int) (struct {
	ForkedBlock        uint64
	FirstEpoch         uint64
	LastEpoch          uint64
	FirstBlock         uint64
	LastBlock          uint64
	LastFinalizedEpoch uint64
	LastFinalizedBlock uint64
	Timestamp          uint64
	FirstEnterEpoch    uint64
	LastEnterEpoch     uint64
	NextBlockToRebase  uint64
	Rebased            bool
}, error) {
	return _RootChain.Contract.Forks(&_RootChain.CallOpts, arg0)
}

// Forks is a free data retrieval call binding the contract method 0x4ba3a126.
//
// Solidity: function forks(uint256 ) constant returns(uint64 forkedBlock, uint64 firstEpoch, uint64 lastEpoch, uint64 firstBlock, uint64 lastBlock, uint64 lastFinalizedEpoch, uint64 lastFinalizedBlock, uint64 timestamp, uint64 firstEnterEpoch, uint64 lastEnterEpoch, uint64 nextBlockToRebase, bool rebased)
func (_RootChain *RootChainCallerSession) Forks(arg0 *big.Int) (struct {
	ForkedBlock        uint64
	FirstEpoch         uint64
	LastEpoch          uint64
	FirstBlock         uint64
	LastBlock          uint64
	LastFinalizedEpoch uint64
	LastFinalizedBlock uint64
	Timestamp          uint64
	FirstEnterEpoch    uint64
	LastEnterEpoch     uint64
	NextBlockToRebase  uint64
	Rebased            bool
}, error) {
	return _RootChain.Contract.Forks(&_RootChain.CallOpts, arg0)
}

// GetBlock is a free data retrieval call binding the contract method 0x4a44bdb8.
//
// Solidity: function getBlock(uint256 forkNumber, uint256 blockNumber) constant returns(Struct0)
func (_RootChain *RootChainCaller) GetBlock(opts *bind.CallOpts, forkNumber *big.Int, blockNumber *big.Int) (Struct0, error) {
	var (
		ret0 = new(Struct0)
	)
	out := ret0
	err := _RootChain.contract.Call(opts, out, "getBlock", forkNumber, blockNumber)
	return *ret0, err
}

// GetBlock is a free data retrieval call binding the contract method 0x4a44bdb8.
//
// Solidity: function getBlock(uint256 forkNumber, uint256 blockNumber) constant returns(Struct0)
func (_RootChain *RootChainSession) GetBlock(forkNumber *big.Int, blockNumber *big.Int) (Struct0, error) {
	return _RootChain.Contract.GetBlock(&_RootChain.CallOpts, forkNumber, blockNumber)
}

// GetBlock is a free data retrieval call binding the contract method 0x4a44bdb8.
//
// Solidity: function getBlock(uint256 forkNumber, uint256 blockNumber) constant returns(Struct0)
func (_RootChain *RootChainCallerSession) GetBlock(forkNumber *big.Int, blockNumber *big.Int) (Struct0, error) {
	return _RootChain.Contract.GetBlock(&_RootChain.CallOpts, forkNumber, blockNumber)
}

// GetBlockFinalizedAt is a free data retrieval call binding the contract method 0x5b884682.
//
// Solidity: function getBlockFinalizedAt(uint256 forkNumber, uint256 blockNumber) constant returns(uint256)
func (_RootChain *RootChainCaller) GetBlockFinalizedAt(opts *bind.CallOpts, forkNumber *big.Int, blockNumber *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChain.contract.Call(opts, out, "getBlockFinalizedAt", forkNumber, blockNumber)
	return *ret0, err
}

// GetBlockFinalizedAt is a free data retrieval call binding the contract method 0x5b884682.
//
// Solidity: function getBlockFinalizedAt(uint256 forkNumber, uint256 blockNumber) constant returns(uint256)
func (_RootChain *RootChainSession) GetBlockFinalizedAt(forkNumber *big.Int, blockNumber *big.Int) (*big.Int, error) {
	return _RootChain.Contract.GetBlockFinalizedAt(&_RootChain.CallOpts, forkNumber, blockNumber)
}

// GetBlockFinalizedAt is a free data retrieval call binding the contract method 0x5b884682.
//
// Solidity: function getBlockFinalizedAt(uint256 forkNumber, uint256 blockNumber) constant returns(uint256)
func (_RootChain *RootChainCallerSession) GetBlockFinalizedAt(forkNumber *big.Int, blockNumber *big.Int) (*big.Int, error) {
	return _RootChain.Contract.GetBlockFinalizedAt(&_RootChain.CallOpts, forkNumber, blockNumber)
}

// GetEROBytes is a free data retrieval call binding the contract method 0xd1723a96.
//
// Solidity: function getEROBytes(uint256 _requestId) constant returns(bytes out)
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
// Solidity: function getEROBytes(uint256 _requestId) constant returns(bytes out)
func (_RootChain *RootChainSession) GetEROBytes(_requestId *big.Int) ([]byte, error) {
	return _RootChain.Contract.GetEROBytes(&_RootChain.CallOpts, _requestId)
}

// GetEROBytes is a free data retrieval call binding the contract method 0xd1723a96.
//
// Solidity: function getEROBytes(uint256 _requestId) constant returns(bytes out)
func (_RootChain *RootChainCallerSession) GetEROBytes(_requestId *big.Int) ([]byte, error) {
	return _RootChain.Contract.GetEROBytes(&_RootChain.CallOpts, _requestId)
}

// GetEpoch is a free data retrieval call binding the contract method 0x2b25a38b.
//
// Solidity: function getEpoch(uint256 forkNumber, uint256 epochNumber) constant returns(Struct3 epoch)
func (_RootChain *RootChainCaller) GetEpoch(opts *bind.CallOpts, forkNumber *big.Int, epochNumber *big.Int) (Struct3, error) {
	var (
		ret0 = new(Struct3)
	)
	out := ret0
	err := _RootChain.contract.Call(opts, out, "getEpoch", forkNumber, epochNumber)
	return *ret0, err
}

// GetEpoch is a free data retrieval call binding the contract method 0x2b25a38b.
//
// Solidity: function getEpoch(uint256 forkNumber, uint256 epochNumber) constant returns(Struct3 epoch)
func (_RootChain *RootChainSession) GetEpoch(forkNumber *big.Int, epochNumber *big.Int) (Struct3, error) {
	return _RootChain.Contract.GetEpoch(&_RootChain.CallOpts, forkNumber, epochNumber)
}

// GetEpoch is a free data retrieval call binding the contract method 0x2b25a38b.
//
// Solidity: function getEpoch(uint256 forkNumber, uint256 epochNumber) constant returns(Struct3 epoch)
func (_RootChain *RootChainCallerSession) GetEpoch(forkNumber *big.Int, epochNumber *big.Int) (Struct3, error) {
	return _RootChain.Contract.GetEpoch(&_RootChain.CallOpts, forkNumber, epochNumber)
}

// GetLastEpoch is a free data retrieval call binding the contract method 0x398bac63.
//
// Solidity: function getLastEpoch() constant returns(Struct3)
func (_RootChain *RootChainCaller) GetLastEpoch(opts *bind.CallOpts) (Struct3, error) {
	var (
		ret0 = new(Struct3)
	)
	out := ret0
	err := _RootChain.contract.Call(opts, out, "getLastEpoch")
	return *ret0, err
}

// GetLastEpoch is a free data retrieval call binding the contract method 0x398bac63.
//
// Solidity: function getLastEpoch() constant returns(Struct3)
func (_RootChain *RootChainSession) GetLastEpoch() (Struct3, error) {
	return _RootChain.Contract.GetLastEpoch(&_RootChain.CallOpts)
}

// GetLastEpoch is a free data retrieval call binding the contract method 0x398bac63.
//
// Solidity: function getLastEpoch() constant returns(Struct3)
func (_RootChain *RootChainCallerSession) GetLastEpoch() (Struct3, error) {
	return _RootChain.Contract.GetLastEpoch(&_RootChain.CallOpts)
}

// GetLastFinalizedBlock is a free data retrieval call binding the contract method 0xd636857e.
//
// Solidity: function getLastFinalizedBlock(uint256 forkNumber) constant returns(uint256)
func (_RootChain *RootChainCaller) GetLastFinalizedBlock(opts *bind.CallOpts, forkNumber *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChain.contract.Call(opts, out, "getLastFinalizedBlock", forkNumber)
	return *ret0, err
}

// GetLastFinalizedBlock is a free data retrieval call binding the contract method 0xd636857e.
//
// Solidity: function getLastFinalizedBlock(uint256 forkNumber) constant returns(uint256)
func (_RootChain *RootChainSession) GetLastFinalizedBlock(forkNumber *big.Int) (*big.Int, error) {
	return _RootChain.Contract.GetLastFinalizedBlock(&_RootChain.CallOpts, forkNumber)
}

// GetLastFinalizedBlock is a free data retrieval call binding the contract method 0xd636857e.
//
// Solidity: function getLastFinalizedBlock(uint256 forkNumber) constant returns(uint256)
func (_RootChain *RootChainCallerSession) GetLastFinalizedBlock(forkNumber *big.Int) (*big.Int, error) {
	return _RootChain.Contract.GetLastFinalizedBlock(&_RootChain.CallOpts, forkNumber)
}

// GetLastFinalizedEpoch is a free data retrieval call binding the contract method 0x019dc099.
//
// Solidity: function getLastFinalizedEpoch(uint256 forkNumber) constant returns(uint256)
func (_RootChain *RootChainCaller) GetLastFinalizedEpoch(opts *bind.CallOpts, forkNumber *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChain.contract.Call(opts, out, "getLastFinalizedEpoch", forkNumber)
	return *ret0, err
}

// GetLastFinalizedEpoch is a free data retrieval call binding the contract method 0x019dc099.
//
// Solidity: function getLastFinalizedEpoch(uint256 forkNumber) constant returns(uint256)
func (_RootChain *RootChainSession) GetLastFinalizedEpoch(forkNumber *big.Int) (*big.Int, error) {
	return _RootChain.Contract.GetLastFinalizedEpoch(&_RootChain.CallOpts, forkNumber)
}

// GetLastFinalizedEpoch is a free data retrieval call binding the contract method 0x019dc099.
//
// Solidity: function getLastFinalizedEpoch(uint256 forkNumber) constant returns(uint256)
func (_RootChain *RootChainCallerSession) GetLastFinalizedEpoch(forkNumber *big.Int) (*big.Int, error) {
	return _RootChain.Contract.GetLastFinalizedEpoch(&_RootChain.CallOpts, forkNumber)
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
// Solidity: function getRequestFinalized(uint256 _requestId, bool _userActivated) constant returns(bool finalized)
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
// Solidity: function getRequestFinalized(uint256 _requestId, bool _userActivated) constant returns(bool finalized)
func (_RootChain *RootChainSession) GetRequestFinalized(_requestId *big.Int, _userActivated bool) (bool, error) {
	return _RootChain.Contract.GetRequestFinalized(&_RootChain.CallOpts, _requestId, _userActivated)
}

// GetRequestFinalized is a free data retrieval call binding the contract method 0xf28a7afa.
//
// Solidity: function getRequestFinalized(uint256 _requestId, bool _userActivated) constant returns(bool finalized)
func (_RootChain *RootChainCallerSession) GetRequestFinalized(_requestId *big.Int, _userActivated bool) (bool, error) {
	return _RootChain.Contract.GetRequestFinalized(&_RootChain.CallOpts, _requestId, _userActivated)
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

// LastAppliedEpochNumber is a free data retrieval call binding the contract method 0xc8ad329f.
//
// Solidity: function lastAppliedEpochNumber() constant returns(uint256)
func (_RootChain *RootChainCaller) LastAppliedEpochNumber(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChain.contract.Call(opts, out, "lastAppliedEpochNumber")
	return *ret0, err
}

// LastAppliedEpochNumber is a free data retrieval call binding the contract method 0xc8ad329f.
//
// Solidity: function lastAppliedEpochNumber() constant returns(uint256)
func (_RootChain *RootChainSession) LastAppliedEpochNumber() (*big.Int, error) {
	return _RootChain.Contract.LastAppliedEpochNumber(&_RootChain.CallOpts)
}

// LastAppliedEpochNumber is a free data retrieval call binding the contract method 0xc8ad329f.
//
// Solidity: function lastAppliedEpochNumber() constant returns(uint256)
func (_RootChain *RootChainCallerSession) LastAppliedEpochNumber() (*big.Int, error) {
	return _RootChain.Contract.LastAppliedEpochNumber(&_RootChain.CallOpts)
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

// LastBlock is a free data retrieval call binding the contract method 0x1ec2042b.
//
// Solidity: function lastBlock(uint256 forkNumber) constant returns(uint256 lastBlock)
func (_RootChain *RootChainCaller) LastBlock(opts *bind.CallOpts, forkNumber *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChain.contract.Call(opts, out, "lastBlock", forkNumber)
	return *ret0, err
}

// LastBlock is a free data retrieval call binding the contract method 0x1ec2042b.
//
// Solidity: function lastBlock(uint256 forkNumber) constant returns(uint256 lastBlock)
func (_RootChain *RootChainSession) LastBlock(forkNumber *big.Int) (*big.Int, error) {
	return _RootChain.Contract.LastBlock(&_RootChain.CallOpts, forkNumber)
}

// LastBlock is a free data retrieval call binding the contract method 0x1ec2042b.
//
// Solidity: function lastBlock(uint256 forkNumber) constant returns(uint256 lastBlock)
func (_RootChain *RootChainCallerSession) LastBlock(forkNumber *big.Int) (*big.Int, error) {
	return _RootChain.Contract.LastBlock(&_RootChain.CallOpts, forkNumber)
}

// LastEpoch is a free data retrieval call binding the contract method 0x11e4c914.
//
// Solidity: function lastEpoch(uint256 forkNumber) constant returns(uint256 lastBlock)
func (_RootChain *RootChainCaller) LastEpoch(opts *bind.CallOpts, forkNumber *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChain.contract.Call(opts, out, "lastEpoch", forkNumber)
	return *ret0, err
}

// LastEpoch is a free data retrieval call binding the contract method 0x11e4c914.
//
// Solidity: function lastEpoch(uint256 forkNumber) constant returns(uint256 lastBlock)
func (_RootChain *RootChainSession) LastEpoch(forkNumber *big.Int) (*big.Int, error) {
	return _RootChain.Contract.LastEpoch(&_RootChain.CallOpts, forkNumber)
}

// LastEpoch is a free data retrieval call binding the contract method 0x11e4c914.
//
// Solidity: function lastEpoch(uint256 forkNumber) constant returns(uint256 lastBlock)
func (_RootChain *RootChainCallerSession) LastEpoch(forkNumber *big.Int) (*big.Int, error) {
	return _RootChain.Contract.LastEpoch(&_RootChain.CallOpts, forkNumber)
}

// LastNonEmptyRequestEpoch is a free data retrieval call binding the contract method 0xb6715647.
//
// Solidity: function lastNonEmptyRequestEpoch(uint256 ) constant returns(uint256)
func (_RootChain *RootChainCaller) LastNonEmptyRequestEpoch(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChain.contract.Call(opts, out, "lastNonEmptyRequestEpoch", arg0)
	return *ret0, err
}

// LastNonEmptyRequestEpoch is a free data retrieval call binding the contract method 0xb6715647.
//
// Solidity: function lastNonEmptyRequestEpoch(uint256 ) constant returns(uint256)
func (_RootChain *RootChainSession) LastNonEmptyRequestEpoch(arg0 *big.Int) (*big.Int, error) {
	return _RootChain.Contract.LastNonEmptyRequestEpoch(&_RootChain.CallOpts, arg0)
}

// LastNonEmptyRequestEpoch is a free data retrieval call binding the contract method 0xb6715647.
//
// Solidity: function lastNonEmptyRequestEpoch(uint256 ) constant returns(uint256)
func (_RootChain *RootChainCallerSession) LastNonEmptyRequestEpoch(arg0 *big.Int) (*big.Int, error) {
	return _RootChain.Contract.LastNonEmptyRequestEpoch(&_RootChain.CallOpts, arg0)
}

// NumEnterForORB is a free data retrieval call binding the contract method 0x23691566.
//
// Solidity: function numEnterForORB() constant returns(uint256)
func (_RootChain *RootChainCaller) NumEnterForORB(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChain.contract.Call(opts, out, "numEnterForORB")
	return *ret0, err
}

// NumEnterForORB is a free data retrieval call binding the contract method 0x23691566.
//
// Solidity: function numEnterForORB() constant returns(uint256)
func (_RootChain *RootChainSession) NumEnterForORB() (*big.Int, error) {
	return _RootChain.Contract.NumEnterForORB(&_RootChain.CallOpts)
}

// NumEnterForORB is a free data retrieval call binding the contract method 0x23691566.
//
// Solidity: function numEnterForORB() constant returns(uint256)
func (_RootChain *RootChainCallerSession) NumEnterForORB() (*big.Int, error) {
	return _RootChain.Contract.NumEnterForORB(&_RootChain.CallOpts)
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
// Solidity: function requestableContracts(address ) constant returns(address)
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
// Solidity: function requestableContracts(address ) constant returns(address)
func (_RootChain *RootChainSession) RequestableContracts(arg0 common.Address) (common.Address, error) {
	return _RootChain.Contract.RequestableContracts(&_RootChain.CallOpts, arg0)
}

// RequestableContracts is a free data retrieval call binding the contract method 0xda0185f8.
//
// Solidity: function requestableContracts(address ) constant returns(address)
func (_RootChain *RootChainCallerSession) RequestableContracts(arg0 common.Address) (common.Address, error) {
	return _RootChain.Contract.RequestableContracts(&_RootChain.CallOpts, arg0)
}

// SubmitHandler is a free data retrieval call binding the contract method 0xe259faf7.
//
// Solidity: function submitHandler() constant returns(address)
func (_RootChain *RootChainCaller) SubmitHandler(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RootChain.contract.Call(opts, out, "submitHandler")
	return *ret0, err
}

// SubmitHandler is a free data retrieval call binding the contract method 0xe259faf7.
//
// Solidity: function submitHandler() constant returns(address)
func (_RootChain *RootChainSession) SubmitHandler() (common.Address, error) {
	return _RootChain.Contract.SubmitHandler(&_RootChain.CallOpts)
}

// SubmitHandler is a free data retrieval call binding the contract method 0xe259faf7.
//
// Solidity: function submitHandler() constant returns(address)
func (_RootChain *RootChainCallerSession) SubmitHandler() (common.Address, error) {
	return _RootChain.Contract.SubmitHandler(&_RootChain.CallOpts)
}

// ChallengeExit is a paid mutator transaction binding the contract method 0x404f7d66.
//
// Solidity: function challengeExit(uint256 _forkNumber, uint256 _blockNumber, uint256 _index, bytes _receiptData, bytes _proof) returns()
func (_RootChain *RootChainTransactor) ChallengeExit(opts *bind.TransactOpts, _forkNumber *big.Int, _blockNumber *big.Int, _index *big.Int, _receiptData []byte, _proof []byte) (*types.Transaction, error) {
	return _RootChain.contract.Transact(opts, "challengeExit", _forkNumber, _blockNumber, _index, _receiptData, _proof)
}

// ChallengeExit is a paid mutator transaction binding the contract method 0x404f7d66.
//
// Solidity: function challengeExit(uint256 _forkNumber, uint256 _blockNumber, uint256 _index, bytes _receiptData, bytes _proof) returns()
func (_RootChain *RootChainSession) ChallengeExit(_forkNumber *big.Int, _blockNumber *big.Int, _index *big.Int, _receiptData []byte, _proof []byte) (*types.Transaction, error) {
	return _RootChain.Contract.ChallengeExit(&_RootChain.TransactOpts, _forkNumber, _blockNumber, _index, _receiptData, _proof)
}

// ChallengeExit is a paid mutator transaction binding the contract method 0x404f7d66.
//
// Solidity: function challengeExit(uint256 _forkNumber, uint256 _blockNumber, uint256 _index, bytes _receiptData, bytes _proof) returns()
func (_RootChain *RootChainTransactorSession) ChallengeExit(_forkNumber *big.Int, _blockNumber *big.Int, _index *big.Int, _receiptData []byte, _proof []byte) (*types.Transaction, error) {
	return _RootChain.Contract.ChallengeExit(&_RootChain.TransactOpts, _forkNumber, _blockNumber, _index, _receiptData, _proof)
}

// ChallengeNullAddress is a paid mutator transaction binding the contract method 0x6299fb24.
//
// Solidity: function challengeNullAddress(uint256 _blockNumber, bytes _key, bytes _txByte, uint256 _branchMask, bytes32[] _siblings) returns()
func (_RootChain *RootChainTransactor) ChallengeNullAddress(opts *bind.TransactOpts, _blockNumber *big.Int, _key []byte, _txByte []byte, _branchMask *big.Int, _siblings [][32]byte) (*types.Transaction, error) {
	return _RootChain.contract.Transact(opts, "challengeNullAddress", _blockNumber, _key, _txByte, _branchMask, _siblings)
}

// ChallengeNullAddress is a paid mutator transaction binding the contract method 0x6299fb24.
//
// Solidity: function challengeNullAddress(uint256 _blockNumber, bytes _key, bytes _txByte, uint256 _branchMask, bytes32[] _siblings) returns()
func (_RootChain *RootChainSession) ChallengeNullAddress(_blockNumber *big.Int, _key []byte, _txByte []byte, _branchMask *big.Int, _siblings [][32]byte) (*types.Transaction, error) {
	return _RootChain.Contract.ChallengeNullAddress(&_RootChain.TransactOpts, _blockNumber, _key, _txByte, _branchMask, _siblings)
}

// ChallengeNullAddress is a paid mutator transaction binding the contract method 0x6299fb24.
//
// Solidity: function challengeNullAddress(uint256 _blockNumber, bytes _key, bytes _txByte, uint256 _branchMask, bytes32[] _siblings) returns()
func (_RootChain *RootChainTransactorSession) ChallengeNullAddress(_blockNumber *big.Int, _key []byte, _txByte []byte, _branchMask *big.Int, _siblings [][32]byte) (*types.Transaction, error) {
	return _RootChain.Contract.ChallengeNullAddress(&_RootChain.TransactOpts, _blockNumber, _key, _txByte, _branchMask, _siblings)
}

// FinalizeBlock is a paid mutator transaction binding the contract method 0x75395a58.
//
// Solidity: function finalizeBlock() returns(bool success)
func (_RootChain *RootChainTransactor) FinalizeBlock(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RootChain.contract.Transact(opts, "finalizeBlock")
}

// FinalizeBlock is a paid mutator transaction binding the contract method 0x75395a58.
//
// Solidity: function finalizeBlock() returns(bool success)
func (_RootChain *RootChainSession) FinalizeBlock() (*types.Transaction, error) {
	return _RootChain.Contract.FinalizeBlock(&_RootChain.TransactOpts)
}

// FinalizeBlock is a paid mutator transaction binding the contract method 0x75395a58.
//
// Solidity: function finalizeBlock() returns(bool success)
func (_RootChain *RootChainTransactorSession) FinalizeBlock() (*types.Transaction, error) {
	return _RootChain.Contract.FinalizeBlock(&_RootChain.TransactOpts)
}

// FinalizeRequest is a paid mutator transaction binding the contract method 0x99bd3600.
//
// Solidity: function finalizeRequest() returns(bool success)
func (_RootChain *RootChainTransactor) FinalizeRequest(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RootChain.contract.Transact(opts, "finalizeRequest")
}

// FinalizeRequest is a paid mutator transaction binding the contract method 0x99bd3600.
//
// Solidity: function finalizeRequest() returns(bool success)
func (_RootChain *RootChainSession) FinalizeRequest() (*types.Transaction, error) {
	return _RootChain.Contract.FinalizeRequest(&_RootChain.TransactOpts)
}

// FinalizeRequest is a paid mutator transaction binding the contract method 0x99bd3600.
//
// Solidity: function finalizeRequest() returns(bool success)
func (_RootChain *RootChainTransactorSession) FinalizeRequest() (*types.Transaction, error) {
	return _RootChain.Contract.FinalizeRequest(&_RootChain.TransactOpts)
}

// FinalizeRequests is a paid mutator transaction binding the contract method 0x54768571.
//
// Solidity: function finalizeRequests(uint256 n) returns(bool success)
func (_RootChain *RootChainTransactor) FinalizeRequests(opts *bind.TransactOpts, n *big.Int) (*types.Transaction, error) {
	return _RootChain.contract.Transact(opts, "finalizeRequests", n)
}

// FinalizeRequests is a paid mutator transaction binding the contract method 0x54768571.
//
// Solidity: function finalizeRequests(uint256 n) returns(bool success)
func (_RootChain *RootChainSession) FinalizeRequests(n *big.Int) (*types.Transaction, error) {
	return _RootChain.Contract.FinalizeRequests(&_RootChain.TransactOpts, n)
}

// FinalizeRequests is a paid mutator transaction binding the contract method 0x54768571.
//
// Solidity: function finalizeRequests(uint256 n) returns(bool success)
func (_RootChain *RootChainTransactorSession) FinalizeRequests(n *big.Int) (*types.Transaction, error) {
	return _RootChain.Contract.FinalizeRequests(&_RootChain.TransactOpts, n)
}

// MakeERU is a paid mutator transaction binding the contract method 0x78fe705f.
//
// Solidity: function makeERU(address _to, bytes32 _trieKey, bytes _trieValue) returns(bool success)
func (_RootChain *RootChainTransactor) MakeERU(opts *bind.TransactOpts, _to common.Address, _trieKey [32]byte, _trieValue []byte) (*types.Transaction, error) {
	return _RootChain.contract.Transact(opts, "makeERU", _to, _trieKey, _trieValue)
}

// MakeERU is a paid mutator transaction binding the contract method 0x78fe705f.
//
// Solidity: function makeERU(address _to, bytes32 _trieKey, bytes _trieValue) returns(bool success)
func (_RootChain *RootChainSession) MakeERU(_to common.Address, _trieKey [32]byte, _trieValue []byte) (*types.Transaction, error) {
	return _RootChain.Contract.MakeERU(&_RootChain.TransactOpts, _to, _trieKey, _trieValue)
}

// MakeERU is a paid mutator transaction binding the contract method 0x78fe705f.
//
// Solidity: function makeERU(address _to, bytes32 _trieKey, bytes _trieValue) returns(bool success)
func (_RootChain *RootChainTransactorSession) MakeERU(_to common.Address, _trieKey [32]byte, _trieValue []byte) (*types.Transaction, error) {
	return _RootChain.Contract.MakeERU(&_RootChain.TransactOpts, _to, _trieKey, _trieValue)
}

// MapRequestableContractByOperator is a paid mutator transaction binding the contract method 0xcb5d742f.
//
// Solidity: function mapRequestableContractByOperator(address _rootchain, address _childchain) returns(bool success)
func (_RootChain *RootChainTransactor) MapRequestableContractByOperator(opts *bind.TransactOpts, _rootchain common.Address, _childchain common.Address) (*types.Transaction, error) {
	return _RootChain.contract.Transact(opts, "mapRequestableContractByOperator", _rootchain, _childchain)
}

// MapRequestableContractByOperator is a paid mutator transaction binding the contract method 0xcb5d742f.
//
// Solidity: function mapRequestableContractByOperator(address _rootchain, address _childchain) returns(bool success)
func (_RootChain *RootChainSession) MapRequestableContractByOperator(_rootchain common.Address, _childchain common.Address) (*types.Transaction, error) {
	return _RootChain.Contract.MapRequestableContractByOperator(&_RootChain.TransactOpts, _rootchain, _childchain)
}

// MapRequestableContractByOperator is a paid mutator transaction binding the contract method 0xcb5d742f.
//
// Solidity: function mapRequestableContractByOperator(address _rootchain, address _childchain) returns(bool success)
func (_RootChain *RootChainTransactorSession) MapRequestableContractByOperator(_rootchain common.Address, _childchain common.Address) (*types.Transaction, error) {
	return _RootChain.Contract.MapRequestableContractByOperator(&_RootChain.TransactOpts, _rootchain, _childchain)
}

// PrepareToSubmitURB is a paid mutator transaction binding the contract method 0xe6925d08.
//
// Solidity: function prepareToSubmitURB() returns()
func (_RootChain *RootChainTransactor) PrepareToSubmitURB(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RootChain.contract.Transact(opts, "prepareToSubmitURB")
}

// PrepareToSubmitURB is a paid mutator transaction binding the contract method 0xe6925d08.
//
// Solidity: function prepareToSubmitURB() returns()
func (_RootChain *RootChainSession) PrepareToSubmitURB() (*types.Transaction, error) {
	return _RootChain.Contract.PrepareToSubmitURB(&_RootChain.TransactOpts)
}

// PrepareToSubmitURB is a paid mutator transaction binding the contract method 0xe6925d08.
//
// Solidity: function prepareToSubmitURB() returns()
func (_RootChain *RootChainTransactorSession) PrepareToSubmitURB() (*types.Transaction, error) {
	return _RootChain.Contract.PrepareToSubmitURB(&_RootChain.TransactOpts)
}

// StartEnter is a paid mutator transaction binding the contract method 0x2aa196c8.
//
// Solidity: function startEnter(address _to, bytes32 _trieKey, bytes _trieValue) returns(bool success)
func (_RootChain *RootChainTransactor) StartEnter(opts *bind.TransactOpts, _to common.Address, _trieKey [32]byte, _trieValue []byte) (*types.Transaction, error) {
	return _RootChain.contract.Transact(opts, "startEnter", _to, _trieKey, _trieValue)
}

// StartEnter is a paid mutator transaction binding the contract method 0x2aa196c8.
//
// Solidity: function startEnter(address _to, bytes32 _trieKey, bytes _trieValue) returns(bool success)
func (_RootChain *RootChainSession) StartEnter(_to common.Address, _trieKey [32]byte, _trieValue []byte) (*types.Transaction, error) {
	return _RootChain.Contract.StartEnter(&_RootChain.TransactOpts, _to, _trieKey, _trieValue)
}

// StartEnter is a paid mutator transaction binding the contract method 0x2aa196c8.
//
// Solidity: function startEnter(address _to, bytes32 _trieKey, bytes _trieValue) returns(bool success)
func (_RootChain *RootChainTransactorSession) StartEnter(_to common.Address, _trieKey [32]byte, _trieValue []byte) (*types.Transaction, error) {
	return _RootChain.Contract.StartEnter(&_RootChain.TransactOpts, _to, _trieKey, _trieValue)
}

// StartExit is a paid mutator transaction binding the contract method 0xe7f96505.
//
// Solidity: function startExit(address _to, bytes32 _trieKey, bytes _trieValue) returns(bool success)
func (_RootChain *RootChainTransactor) StartExit(opts *bind.TransactOpts, _to common.Address, _trieKey [32]byte, _trieValue []byte) (*types.Transaction, error) {
	return _RootChain.contract.Transact(opts, "startExit", _to, _trieKey, _trieValue)
}

// StartExit is a paid mutator transaction binding the contract method 0xe7f96505.
//
// Solidity: function startExit(address _to, bytes32 _trieKey, bytes _trieValue) returns(bool success)
func (_RootChain *RootChainSession) StartExit(_to common.Address, _trieKey [32]byte, _trieValue []byte) (*types.Transaction, error) {
	return _RootChain.Contract.StartExit(&_RootChain.TransactOpts, _to, _trieKey, _trieValue)
}

// StartExit is a paid mutator transaction binding the contract method 0xe7f96505.
//
// Solidity: function startExit(address _to, bytes32 _trieKey, bytes _trieValue) returns(bool success)
func (_RootChain *RootChainTransactorSession) StartExit(_to common.Address, _trieKey [32]byte, _trieValue []byte) (*types.Transaction, error) {
	return _RootChain.Contract.StartExit(&_RootChain.TransactOpts, _to, _trieKey, _trieValue)
}

// SubmitNRE is a paid mutator transaction binding the contract method 0x0eaf45a8.
//
// Solidity: function submitNRE(uint256 _pos1, uint256 _pos2, bytes32 _epochStateRoot, bytes32 _epochTransactionsRoot, bytes32 _epochReceiptsRoot) returns(bool success)
func (_RootChain *RootChainTransactor) SubmitNRE(opts *bind.TransactOpts, _pos1 *big.Int, _pos2 *big.Int, _epochStateRoot [32]byte, _epochTransactionsRoot [32]byte, _epochReceiptsRoot [32]byte) (*types.Transaction, error) {
	return _RootChain.contract.Transact(opts, "submitNRE", _pos1, _pos2, _epochStateRoot, _epochTransactionsRoot, _epochReceiptsRoot)
}

// SubmitNRE is a paid mutator transaction binding the contract method 0x0eaf45a8.
//
// Solidity: function submitNRE(uint256 _pos1, uint256 _pos2, bytes32 _epochStateRoot, bytes32 _epochTransactionsRoot, bytes32 _epochReceiptsRoot) returns(bool success)
func (_RootChain *RootChainSession) SubmitNRE(_pos1 *big.Int, _pos2 *big.Int, _epochStateRoot [32]byte, _epochTransactionsRoot [32]byte, _epochReceiptsRoot [32]byte) (*types.Transaction, error) {
	return _RootChain.Contract.SubmitNRE(&_RootChain.TransactOpts, _pos1, _pos2, _epochStateRoot, _epochTransactionsRoot, _epochReceiptsRoot)
}

// SubmitNRE is a paid mutator transaction binding the contract method 0x0eaf45a8.
//
// Solidity: function submitNRE(uint256 _pos1, uint256 _pos2, bytes32 _epochStateRoot, bytes32 _epochTransactionsRoot, bytes32 _epochReceiptsRoot) returns(bool success)
func (_RootChain *RootChainTransactorSession) SubmitNRE(_pos1 *big.Int, _pos2 *big.Int, _epochStateRoot [32]byte, _epochTransactionsRoot [32]byte, _epochReceiptsRoot [32]byte) (*types.Transaction, error) {
	return _RootChain.Contract.SubmitNRE(&_RootChain.TransactOpts, _pos1, _pos2, _epochStateRoot, _epochTransactionsRoot, _epochReceiptsRoot)
}

// SubmitORB is a paid mutator transaction binding the contract method 0xa820c067.
//
// Solidity: function submitORB(uint256 _pos, bytes32 _statesRoot, bytes32 _transactionsRoot, bytes32 _receiptsRoot) returns(bool success)
func (_RootChain *RootChainTransactor) SubmitORB(opts *bind.TransactOpts, _pos *big.Int, _statesRoot [32]byte, _transactionsRoot [32]byte, _receiptsRoot [32]byte) (*types.Transaction, error) {
	return _RootChain.contract.Transact(opts, "submitORB", _pos, _statesRoot, _transactionsRoot, _receiptsRoot)
}

// SubmitORB is a paid mutator transaction binding the contract method 0xa820c067.
//
// Solidity: function submitORB(uint256 _pos, bytes32 _statesRoot, bytes32 _transactionsRoot, bytes32 _receiptsRoot) returns(bool success)
func (_RootChain *RootChainSession) SubmitORB(_pos *big.Int, _statesRoot [32]byte, _transactionsRoot [32]byte, _receiptsRoot [32]byte) (*types.Transaction, error) {
	return _RootChain.Contract.SubmitORB(&_RootChain.TransactOpts, _pos, _statesRoot, _transactionsRoot, _receiptsRoot)
}

// SubmitORB is a paid mutator transaction binding the contract method 0xa820c067.
//
// Solidity: function submitORB(uint256 _pos, bytes32 _statesRoot, bytes32 _transactionsRoot, bytes32 _receiptsRoot) returns(bool success)
func (_RootChain *RootChainTransactorSession) SubmitORB(_pos *big.Int, _statesRoot [32]byte, _transactionsRoot [32]byte, _receiptsRoot [32]byte) (*types.Transaction, error) {
	return _RootChain.Contract.SubmitORB(&_RootChain.TransactOpts, _pos, _statesRoot, _transactionsRoot, _receiptsRoot)
}

// SubmitURB is a paid mutator transaction binding the contract method 0x6f3e4b90.
//
// Solidity: function submitURB(uint256 _pos, bytes32 _statesRoot, bytes32 _transactionsRoot, bytes32 _receiptsRoot) returns(bool success)
func (_RootChain *RootChainTransactor) SubmitURB(opts *bind.TransactOpts, _pos *big.Int, _statesRoot [32]byte, _transactionsRoot [32]byte, _receiptsRoot [32]byte) (*types.Transaction, error) {
	return _RootChain.contract.Transact(opts, "submitURB", _pos, _statesRoot, _transactionsRoot, _receiptsRoot)
}

// SubmitURB is a paid mutator transaction binding the contract method 0x6f3e4b90.
//
// Solidity: function submitURB(uint256 _pos, bytes32 _statesRoot, bytes32 _transactionsRoot, bytes32 _receiptsRoot) returns(bool success)
func (_RootChain *RootChainSession) SubmitURB(_pos *big.Int, _statesRoot [32]byte, _transactionsRoot [32]byte, _receiptsRoot [32]byte) (*types.Transaction, error) {
	return _RootChain.Contract.SubmitURB(&_RootChain.TransactOpts, _pos, _statesRoot, _transactionsRoot, _receiptsRoot)
}

// SubmitURB is a paid mutator transaction binding the contract method 0x6f3e4b90.
//
// Solidity: function submitURB(uint256 _pos, bytes32 _statesRoot, bytes32 _transactionsRoot, bytes32 _receiptsRoot) returns(bool success)
func (_RootChain *RootChainTransactorSession) SubmitURB(_pos *big.Int, _statesRoot [32]byte, _transactionsRoot [32]byte, _receiptsRoot [32]byte) (*types.Transaction, error) {
	return _RootChain.Contract.SubmitURB(&_RootChain.TransactOpts, _pos, _statesRoot, _transactionsRoot, _receiptsRoot)
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
// Solidity: event BlockFinalized(uint256 forkNumber, uint256 blockNumber)
func (_RootChain *RootChainFilterer) FilterBlockFinalized(opts *bind.FilterOpts) (*RootChainBlockFinalizedIterator, error) {

	logs, sub, err := _RootChain.contract.FilterLogs(opts, "BlockFinalized")
	if err != nil {
		return nil, err
	}
	return &RootChainBlockFinalizedIterator{contract: _RootChain.contract, event: "BlockFinalized", logs: logs, sub: sub}, nil
}

// WatchBlockFinalized is a free log subscription operation binding the contract event 0xfb96205e4b3633fd57aa805b26b51ecf528714a10241a4af015929dce86768d9.
//
// Solidity: event BlockFinalized(uint256 forkNumber, uint256 blockNumber)
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

// ParseBlockFinalized is a log parse operation binding the contract event 0xfb96205e4b3633fd57aa805b26b51ecf528714a10241a4af015929dce86768d9.
//
// Solidity: event BlockFinalized(uint256 forkNumber, uint256 blockNumber)
func (_RootChain *RootChainFilterer) ParseBlockFinalized(log types.Log) (*RootChainBlockFinalized, error) {
	event := new(RootChainBlockFinalized)
	if err := _RootChain.contract.UnpackLog(event, "BlockFinalized", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RootChainBlockSubmittedIterator is returned from FilterBlockSubmitted and is used to iterate over the raw logs and unpacked data for BlockSubmitted events raised by the RootChain contract.
type RootChainBlockSubmittedIterator struct {
	Event *RootChainBlockSubmitted // Event containing the contract specifics and raw log

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
func (it *RootChainBlockSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RootChainBlockSubmitted)
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
		it.Event = new(RootChainBlockSubmitted)
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
func (it *RootChainBlockSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RootChainBlockSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RootChainBlockSubmitted represents a BlockSubmitted event raised by the RootChain contract.
type RootChainBlockSubmitted struct {
	Fork          *big.Int
	EpochNumber   *big.Int
	BlockNumber   *big.Int
	IsRequest     bool
	UserActivated bool
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterBlockSubmitted is a free log retrieval operation binding the contract event 0x3d4a04291c66b06f39a4ecb817875b12b5485a05ec563133a56a905305c48e55.
//
// Solidity: event BlockSubmitted(uint256 fork, uint256 epochNumber, uint256 blockNumber, bool isRequest, bool userActivated)
func (_RootChain *RootChainFilterer) FilterBlockSubmitted(opts *bind.FilterOpts) (*RootChainBlockSubmittedIterator, error) {

	logs, sub, err := _RootChain.contract.FilterLogs(opts, "BlockSubmitted")
	if err != nil {
		return nil, err
	}
	return &RootChainBlockSubmittedIterator{contract: _RootChain.contract, event: "BlockSubmitted", logs: logs, sub: sub}, nil
}

// WatchBlockSubmitted is a free log subscription operation binding the contract event 0x3d4a04291c66b06f39a4ecb817875b12b5485a05ec563133a56a905305c48e55.
//
// Solidity: event BlockSubmitted(uint256 fork, uint256 epochNumber, uint256 blockNumber, bool isRequest, bool userActivated)
func (_RootChain *RootChainFilterer) WatchBlockSubmitted(opts *bind.WatchOpts, sink chan<- *RootChainBlockSubmitted) (event.Subscription, error) {

	logs, sub, err := _RootChain.contract.WatchLogs(opts, "BlockSubmitted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RootChainBlockSubmitted)
				if err := _RootChain.contract.UnpackLog(event, "BlockSubmitted", log); err != nil {
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

// ParseBlockSubmitted is a log parse operation binding the contract event 0x3d4a04291c66b06f39a4ecb817875b12b5485a05ec563133a56a905305c48e55.
//
// Solidity: event BlockSubmitted(uint256 fork, uint256 epochNumber, uint256 blockNumber, bool isRequest, bool userActivated)
func (_RootChain *RootChainFilterer) ParseBlockSubmitted(log types.Log) (*RootChainBlockSubmitted, error) {
	event := new(RootChainBlockSubmitted)
	if err := _RootChain.contract.UnpackLog(event, "BlockSubmitted", log); err != nil {
		return nil, err
	}
	return event, nil
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
	TrieKey   []byte
	TrieValue [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterERUCreated is a free log retrieval operation binding the contract event 0xfcbdc2083dadd644b854d91b49aef8db06b8f5a3d5c1192de38ca0ba271d5a0d.
//
// Solidity: event ERUCreated(uint256 requestId, address requestor, address to, bytes trieKey, bytes32 trieValue)
func (_RootChain *RootChainFilterer) FilterERUCreated(opts *bind.FilterOpts) (*RootChainERUCreatedIterator, error) {

	logs, sub, err := _RootChain.contract.FilterLogs(opts, "ERUCreated")
	if err != nil {
		return nil, err
	}
	return &RootChainERUCreatedIterator{contract: _RootChain.contract, event: "ERUCreated", logs: logs, sub: sub}, nil
}

// WatchERUCreated is a free log subscription operation binding the contract event 0xfcbdc2083dadd644b854d91b49aef8db06b8f5a3d5c1192de38ca0ba271d5a0d.
//
// Solidity: event ERUCreated(uint256 requestId, address requestor, address to, bytes trieKey, bytes32 trieValue)
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

// ParseERUCreated is a log parse operation binding the contract event 0xfcbdc2083dadd644b854d91b49aef8db06b8f5a3d5c1192de38ca0ba271d5a0d.
//
// Solidity: event ERUCreated(uint256 requestId, address requestor, address to, bytes trieKey, bytes32 trieValue)
func (_RootChain *RootChainFilterer) ParseERUCreated(log types.Log) (*RootChainERUCreated, error) {
	event := new(RootChainERUCreated)
	if err := _RootChain.contract.UnpackLog(event, "ERUCreated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RootChainEpochFilledIterator is returned from FilterEpochFilled and is used to iterate over the raw logs and unpacked data for EpochFilled events raised by the RootChain contract.
type RootChainEpochFilledIterator struct {
	Event *RootChainEpochFilled // Event containing the contract specifics and raw log

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
func (it *RootChainEpochFilledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RootChainEpochFilled)
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
		it.Event = new(RootChainEpochFilled)
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
func (it *RootChainEpochFilledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RootChainEpochFilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RootChainEpochFilled represents a EpochFilled event raised by the RootChain contract.
type RootChainEpochFilled struct {
	ForkNumber  *big.Int
	EpochNumber *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterEpochFilled is a free log retrieval operation binding the contract event 0x2fdeb407bf5c2b621f04b5c784822dae806c45b49a68aba413cc270128c96816.
//
// Solidity: event EpochFilled(uint256 forkNumber, uint256 epochNumber)
func (_RootChain *RootChainFilterer) FilterEpochFilled(opts *bind.FilterOpts) (*RootChainEpochFilledIterator, error) {

	logs, sub, err := _RootChain.contract.FilterLogs(opts, "EpochFilled")
	if err != nil {
		return nil, err
	}
	return &RootChainEpochFilledIterator{contract: _RootChain.contract, event: "EpochFilled", logs: logs, sub: sub}, nil
}

// WatchEpochFilled is a free log subscription operation binding the contract event 0x2fdeb407bf5c2b621f04b5c784822dae806c45b49a68aba413cc270128c96816.
//
// Solidity: event EpochFilled(uint256 forkNumber, uint256 epochNumber)
func (_RootChain *RootChainFilterer) WatchEpochFilled(opts *bind.WatchOpts, sink chan<- *RootChainEpochFilled) (event.Subscription, error) {

	logs, sub, err := _RootChain.contract.WatchLogs(opts, "EpochFilled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RootChainEpochFilled)
				if err := _RootChain.contract.UnpackLog(event, "EpochFilled", log); err != nil {
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

// ParseEpochFilled is a log parse operation binding the contract event 0x2fdeb407bf5c2b621f04b5c784822dae806c45b49a68aba413cc270128c96816.
//
// Solidity: event EpochFilled(uint256 forkNumber, uint256 epochNumber)
func (_RootChain *RootChainFilterer) ParseEpochFilled(log types.Log) (*RootChainEpochFilled, error) {
	event := new(RootChainEpochFilled)
	if err := _RootChain.contract.UnpackLog(event, "EpochFilled", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RootChainEpochFillingIterator is returned from FilterEpochFilling and is used to iterate over the raw logs and unpacked data for EpochFilling events raised by the RootChain contract.
type RootChainEpochFillingIterator struct {
	Event *RootChainEpochFilling // Event containing the contract specifics and raw log

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
func (it *RootChainEpochFillingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RootChainEpochFilling)
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
		it.Event = new(RootChainEpochFilling)
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
func (it *RootChainEpochFillingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RootChainEpochFillingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RootChainEpochFilling represents a EpochFilling event raised by the RootChain contract.
type RootChainEpochFilling struct {
	ForkNumber  *big.Int
	EpochNumber *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterEpochFilling is a free log retrieval operation binding the contract event 0x27b09f0953d27bbff306fe25b2987ac5a813248ac30cb2bbd5daf95e7b0e6dc0.
//
// Solidity: event EpochFilling(uint256 forkNumber, uint256 epochNumber)
func (_RootChain *RootChainFilterer) FilterEpochFilling(opts *bind.FilterOpts) (*RootChainEpochFillingIterator, error) {

	logs, sub, err := _RootChain.contract.FilterLogs(opts, "EpochFilling")
	if err != nil {
		return nil, err
	}
	return &RootChainEpochFillingIterator{contract: _RootChain.contract, event: "EpochFilling", logs: logs, sub: sub}, nil
}

// WatchEpochFilling is a free log subscription operation binding the contract event 0x27b09f0953d27bbff306fe25b2987ac5a813248ac30cb2bbd5daf95e7b0e6dc0.
//
// Solidity: event EpochFilling(uint256 forkNumber, uint256 epochNumber)
func (_RootChain *RootChainFilterer) WatchEpochFilling(opts *bind.WatchOpts, sink chan<- *RootChainEpochFilling) (event.Subscription, error) {

	logs, sub, err := _RootChain.contract.WatchLogs(opts, "EpochFilling")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RootChainEpochFilling)
				if err := _RootChain.contract.UnpackLog(event, "EpochFilling", log); err != nil {
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

// ParseEpochFilling is a log parse operation binding the contract event 0x27b09f0953d27bbff306fe25b2987ac5a813248ac30cb2bbd5daf95e7b0e6dc0.
//
// Solidity: event EpochFilling(uint256 forkNumber, uint256 epochNumber)
func (_RootChain *RootChainFilterer) ParseEpochFilling(log types.Log) (*RootChainEpochFilling, error) {
	event := new(RootChainEpochFilling)
	if err := _RootChain.contract.UnpackLog(event, "EpochFilling", log); err != nil {
		return nil, err
	}
	return event, nil
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
// Solidity: event EpochFinalized(uint256 forkNumber, uint256 epochNumber, uint256 startBlockNumber, uint256 endBlockNumber)
func (_RootChain *RootChainFilterer) FilterEpochFinalized(opts *bind.FilterOpts) (*RootChainEpochFinalizedIterator, error) {

	logs, sub, err := _RootChain.contract.FilterLogs(opts, "EpochFinalized")
	if err != nil {
		return nil, err
	}
	return &RootChainEpochFinalizedIterator{contract: _RootChain.contract, event: "EpochFinalized", logs: logs, sub: sub}, nil
}

// WatchEpochFinalized is a free log subscription operation binding the contract event 0x70801d4d63b3da6c19ba7349911f45bed5a99ccdfb51b8138c105872529bebd5.
//
// Solidity: event EpochFinalized(uint256 forkNumber, uint256 epochNumber, uint256 startBlockNumber, uint256 endBlockNumber)
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

// ParseEpochFinalized is a log parse operation binding the contract event 0x70801d4d63b3da6c19ba7349911f45bed5a99ccdfb51b8138c105872529bebd5.
//
// Solidity: event EpochFinalized(uint256 forkNumber, uint256 epochNumber, uint256 startBlockNumber, uint256 endBlockNumber)
func (_RootChain *RootChainFilterer) ParseEpochFinalized(log types.Log) (*RootChainEpochFinalized, error) {
	event := new(RootChainEpochFinalized)
	if err := _RootChain.contract.UnpackLog(event, "EpochFinalized", log); err != nil {
		return nil, err
	}
	return event, nil
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
// Solidity: event EpochPrepared(uint256 forkNumber, uint256 epochNumber, uint256 startBlockNumber, uint256 endBlockNumber, uint256 requestStart, uint256 requestEnd, bool epochIsEmpty, bool isRequest, bool userActivated, bool rebase)
func (_RootChain *RootChainFilterer) FilterEpochPrepared(opts *bind.FilterOpts) (*RootChainEpochPreparedIterator, error) {

	logs, sub, err := _RootChain.contract.FilterLogs(opts, "EpochPrepared")
	if err != nil {
		return nil, err
	}
	return &RootChainEpochPreparedIterator{contract: _RootChain.contract, event: "EpochPrepared", logs: logs, sub: sub}, nil
}

// WatchEpochPrepared is a free log subscription operation binding the contract event 0x1a69c0760aa329b76f72579129869013ebd3d41594db019c0e997b939fcb32e3.
//
// Solidity: event EpochPrepared(uint256 forkNumber, uint256 epochNumber, uint256 startBlockNumber, uint256 endBlockNumber, uint256 requestStart, uint256 requestEnd, bool epochIsEmpty, bool isRequest, bool userActivated, bool rebase)
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

// ParseEpochPrepared is a log parse operation binding the contract event 0x1a69c0760aa329b76f72579129869013ebd3d41594db019c0e997b939fcb32e3.
//
// Solidity: event EpochPrepared(uint256 forkNumber, uint256 epochNumber, uint256 startBlockNumber, uint256 endBlockNumber, uint256 requestStart, uint256 requestEnd, bool epochIsEmpty, bool isRequest, bool userActivated, bool rebase)
func (_RootChain *RootChainFilterer) ParseEpochPrepared(log types.Log) (*RootChainEpochPrepared, error) {
	event := new(RootChainEpochPrepared)
	if err := _RootChain.contract.UnpackLog(event, "EpochPrepared", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RootChainEpochRebasedIterator is returned from FilterEpochRebased and is used to iterate over the raw logs and unpacked data for EpochRebased events raised by the RootChain contract.
type RootChainEpochRebasedIterator struct {
	Event *RootChainEpochRebased // Event containing the contract specifics and raw log

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
func (it *RootChainEpochRebasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RootChainEpochRebased)
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
		it.Event = new(RootChainEpochRebased)
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
func (it *RootChainEpochRebasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RootChainEpochRebasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RootChainEpochRebased represents a EpochRebased event raised by the RootChain contract.
type RootChainEpochRebased struct {
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
// Solidity: event EpochRebased(uint256 forkNumber, uint256 epochNumber, uint256 startBlockNumber, uint256 endBlockNumber, uint256 requestStart, uint256 requestEnd, bool epochIsEmpty, bool isRequest, bool userActivated)
func (_RootChain *RootChainFilterer) FilterEpochRebased(opts *bind.FilterOpts) (*RootChainEpochRebasedIterator, error) {

	logs, sub, err := _RootChain.contract.FilterLogs(opts, "EpochRebased")
	if err != nil {
		return nil, err
	}
	return &RootChainEpochRebasedIterator{contract: _RootChain.contract, event: "EpochRebased", logs: logs, sub: sub}, nil
}

// WatchEpochRebased is a free log subscription operation binding the contract event 0x030c1c69405c93021f28f57557240dee939a320b826a1fd0d39bf6e629ecab47.
//
// Solidity: event EpochRebased(uint256 forkNumber, uint256 epochNumber, uint256 startBlockNumber, uint256 endBlockNumber, uint256 requestStart, uint256 requestEnd, bool epochIsEmpty, bool isRequest, bool userActivated)
func (_RootChain *RootChainFilterer) WatchEpochRebased(opts *bind.WatchOpts, sink chan<- *RootChainEpochRebased) (event.Subscription, error) {

	logs, sub, err := _RootChain.contract.WatchLogs(opts, "EpochRebased")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RootChainEpochRebased)
				if err := _RootChain.contract.UnpackLog(event, "EpochRebased", log); err != nil {
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

// ParseEpochRebased is a log parse operation binding the contract event 0x030c1c69405c93021f28f57557240dee939a320b826a1fd0d39bf6e629ecab47.
//
// Solidity: event EpochRebased(uint256 forkNumber, uint256 epochNumber, uint256 startBlockNumber, uint256 endBlockNumber, uint256 requestStart, uint256 requestEnd, bool epochIsEmpty, bool isRequest, bool userActivated)
func (_RootChain *RootChainFilterer) ParseEpochRebased(log types.Log) (*RootChainEpochRebased, error) {
	event := new(RootChainEpochRebased)
	if err := _RootChain.contract.UnpackLog(event, "EpochRebased", log); err != nil {
		return nil, err
	}
	return event, nil
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
	EpochNumber       *big.Int
	ForkedBlockNumber *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterForked is a free log retrieval operation binding the contract event 0x0647d42ab02f6e0ae76959757dcb6aa6feac1d4ba6f077f1223fb4b1b429f06c.
//
// Solidity: event Forked(uint256 newFork, uint256 epochNumber, uint256 forkedBlockNumber)
func (_RootChain *RootChainFilterer) FilterForked(opts *bind.FilterOpts) (*RootChainForkedIterator, error) {

	logs, sub, err := _RootChain.contract.FilterLogs(opts, "Forked")
	if err != nil {
		return nil, err
	}
	return &RootChainForkedIterator{contract: _RootChain.contract, event: "Forked", logs: logs, sub: sub}, nil
}

// WatchForked is a free log subscription operation binding the contract event 0x0647d42ab02f6e0ae76959757dcb6aa6feac1d4ba6f077f1223fb4b1b429f06c.
//
// Solidity: event Forked(uint256 newFork, uint256 epochNumber, uint256 forkedBlockNumber)
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

// ParseForked is a log parse operation binding the contract event 0x0647d42ab02f6e0ae76959757dcb6aa6feac1d4ba6f077f1223fb4b1b429f06c.
//
// Solidity: event Forked(uint256 newFork, uint256 epochNumber, uint256 forkedBlockNumber)
func (_RootChain *RootChainFilterer) ParseForked(log types.Log) (*RootChainForked, error) {
	event := new(RootChainForked)
	if err := _RootChain.contract.UnpackLog(event, "Forked", log); err != nil {
		return nil, err
	}
	return event, nil
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
// Solidity: event RequestApplied(uint256 requestId, bool userActivated)
func (_RootChain *RootChainFilterer) FilterRequestApplied(opts *bind.FilterOpts) (*RootChainRequestAppliedIterator, error) {

	logs, sub, err := _RootChain.contract.FilterLogs(opts, "RequestApplied")
	if err != nil {
		return nil, err
	}
	return &RootChainRequestAppliedIterator{contract: _RootChain.contract, event: "RequestApplied", logs: logs, sub: sub}, nil
}

// WatchRequestApplied is a free log subscription operation binding the contract event 0x6940a01870e576ceb735867e13863646d517ce10e66c0133186a4ebdfe9388c2.
//
// Solidity: event RequestApplied(uint256 requestId, bool userActivated)
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

// ParseRequestApplied is a log parse operation binding the contract event 0x6940a01870e576ceb735867e13863646d517ce10e66c0133186a4ebdfe9388c2.
//
// Solidity: event RequestApplied(uint256 requestId, bool userActivated)
func (_RootChain *RootChainFilterer) ParseRequestApplied(log types.Log) (*RootChainRequestApplied, error) {
	event := new(RootChainRequestApplied)
	if err := _RootChain.contract.UnpackLog(event, "RequestApplied", log); err != nil {
		return nil, err
	}
	return event, nil
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
// Solidity: event RequestChallenged(uint256 requestId, bool userActivated)
func (_RootChain *RootChainFilterer) FilterRequestChallenged(opts *bind.FilterOpts) (*RootChainRequestChallengedIterator, error) {

	logs, sub, err := _RootChain.contract.FilterLogs(opts, "RequestChallenged")
	if err != nil {
		return nil, err
	}
	return &RootChainRequestChallengedIterator{contract: _RootChain.contract, event: "RequestChallenged", logs: logs, sub: sub}, nil
}

// WatchRequestChallenged is a free log subscription operation binding the contract event 0xc8135db115644ed4ae193313c4c801235ef740d2a57a8d5e6fe26ab66635698a.
//
// Solidity: event RequestChallenged(uint256 requestId, bool userActivated)
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

// ParseRequestChallenged is a log parse operation binding the contract event 0xc8135db115644ed4ae193313c4c801235ef740d2a57a8d5e6fe26ab66635698a.
//
// Solidity: event RequestChallenged(uint256 requestId, bool userActivated)
func (_RootChain *RootChainFilterer) ParseRequestChallenged(log types.Log) (*RootChainRequestChallenged, error) {
	event := new(RootChainRequestChallenged)
	if err := _RootChain.contract.UnpackLog(event, "RequestChallenged", log); err != nil {
		return nil, err
	}
	return event, nil
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
	TrieValue     []byte
	IsExit        bool
	UserActivated bool
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRequestCreated is a free log retrieval operation binding the contract event 0x879922cf5fcada9ebaf8bd7424dc62877f4b220cae07fb6695cc1e8f94c52b4d.
//
// Solidity: event RequestCreated(uint256 requestId, address requestor, address to, uint256 weiAmount, bytes32 trieKey, bytes trieValue, bool isExit, bool userActivated)
func (_RootChain *RootChainFilterer) FilterRequestCreated(opts *bind.FilterOpts) (*RootChainRequestCreatedIterator, error) {

	logs, sub, err := _RootChain.contract.FilterLogs(opts, "RequestCreated")
	if err != nil {
		return nil, err
	}
	return &RootChainRequestCreatedIterator{contract: _RootChain.contract, event: "RequestCreated", logs: logs, sub: sub}, nil
}

// WatchRequestCreated is a free log subscription operation binding the contract event 0x879922cf5fcada9ebaf8bd7424dc62877f4b220cae07fb6695cc1e8f94c52b4d.
//
// Solidity: event RequestCreated(uint256 requestId, address requestor, address to, uint256 weiAmount, bytes32 trieKey, bytes trieValue, bool isExit, bool userActivated)
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

// ParseRequestCreated is a log parse operation binding the contract event 0x879922cf5fcada9ebaf8bd7424dc62877f4b220cae07fb6695cc1e8f94c52b4d.
//
// Solidity: event RequestCreated(uint256 requestId, address requestor, address to, uint256 weiAmount, bytes32 trieKey, bytes trieValue, bool isExit, bool userActivated)
func (_RootChain *RootChainFilterer) ParseRequestCreated(log types.Log) (*RootChainRequestCreated, error) {
	event := new(RootChainRequestCreated)
	if err := _RootChain.contract.UnpackLog(event, "RequestCreated", log); err != nil {
		return nil, err
	}
	return event, nil
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
// Solidity: event RequestFinalized(uint256 requestId, bool userActivated)
func (_RootChain *RootChainFilterer) FilterRequestFinalized(opts *bind.FilterOpts) (*RootChainRequestFinalizedIterator, error) {

	logs, sub, err := _RootChain.contract.FilterLogs(opts, "RequestFinalized")
	if err != nil {
		return nil, err
	}
	return &RootChainRequestFinalizedIterator{contract: _RootChain.contract, event: "RequestFinalized", logs: logs, sub: sub}, nil
}

// WatchRequestFinalized is a free log subscription operation binding the contract event 0x134017cf3262b18f892ee95dde3b0aec9a80cc70a7c96f09c64bd237aceb0473.
//
// Solidity: event RequestFinalized(uint256 requestId, bool userActivated)
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

// ParseRequestFinalized is a log parse operation binding the contract event 0x134017cf3262b18f892ee95dde3b0aec9a80cc70a7c96f09c64bd237aceb0473.
//
// Solidity: event RequestFinalized(uint256 requestId, bool userActivated)
func (_RootChain *RootChainFilterer) ParseRequestFinalized(log types.Log) (*RootChainRequestFinalized, error) {
	event := new(RootChainRequestFinalized)
	if err := _RootChain.contract.UnpackLog(event, "RequestFinalized", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RootChainRequestableContractMappedIterator is returned from FilterRequestableContractMapped and is used to iterate over the raw logs and unpacked data for RequestableContractMapped events raised by the RootChain contract.
type RootChainRequestableContractMappedIterator struct {
	Event *RootChainRequestableContractMapped // Event containing the contract specifics and raw log

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
func (it *RootChainRequestableContractMappedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RootChainRequestableContractMapped)
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
		it.Event = new(RootChainRequestableContractMapped)
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
func (it *RootChainRequestableContractMappedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RootChainRequestableContractMappedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RootChainRequestableContractMapped represents a RequestableContractMapped event raised by the RootChain contract.
type RootChainRequestableContractMapped struct {
	ContractInRootchain  common.Address
	ContractInChildchain common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterRequestableContractMapped is a free log retrieval operation binding the contract event 0xc5ec2ed49686197edd2ed642c7e6096893cc81e6658cde2527030316037715d0.
//
// Solidity: event RequestableContractMapped(address contractInRootchain, address contractInChildchain)
func (_RootChain *RootChainFilterer) FilterRequestableContractMapped(opts *bind.FilterOpts) (*RootChainRequestableContractMappedIterator, error) {

	logs, sub, err := _RootChain.contract.FilterLogs(opts, "RequestableContractMapped")
	if err != nil {
		return nil, err
	}
	return &RootChainRequestableContractMappedIterator{contract: _RootChain.contract, event: "RequestableContractMapped", logs: logs, sub: sub}, nil
}

// WatchRequestableContractMapped is a free log subscription operation binding the contract event 0xc5ec2ed49686197edd2ed642c7e6096893cc81e6658cde2527030316037715d0.
//
// Solidity: event RequestableContractMapped(address contractInRootchain, address contractInChildchain)
func (_RootChain *RootChainFilterer) WatchRequestableContractMapped(opts *bind.WatchOpts, sink chan<- *RootChainRequestableContractMapped) (event.Subscription, error) {

	logs, sub, err := _RootChain.contract.WatchLogs(opts, "RequestableContractMapped")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RootChainRequestableContractMapped)
				if err := _RootChain.contract.UnpackLog(event, "RequestableContractMapped", log); err != nil {
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

// ParseRequestableContractMapped is a log parse operation binding the contract event 0xc5ec2ed49686197edd2ed642c7e6096893cc81e6658cde2527030316037715d0.
//
// Solidity: event RequestableContractMapped(address contractInRootchain, address contractInChildchain)
func (_RootChain *RootChainFilterer) ParseRequestableContractMapped(log types.Log) (*RootChainRequestableContractMapped, error) {
	event := new(RootChainRequestableContractMapped)
	if err := _RootChain.contract.UnpackLog(event, "RequestableContractMapped", log); err != nil {
		return nil, err
	}
	return event, nil
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
// Solidity: event SessionTimeout(bool userActivated)
func (_RootChain *RootChainFilterer) FilterSessionTimeout(opts *bind.FilterOpts) (*RootChainSessionTimeoutIterator, error) {

	logs, sub, err := _RootChain.contract.FilterLogs(opts, "SessionTimeout")
	if err != nil {
		return nil, err
	}
	return &RootChainSessionTimeoutIterator{contract: _RootChain.contract, event: "SessionTimeout", logs: logs, sub: sub}, nil
}

// WatchSessionTimeout is a free log subscription operation binding the contract event 0x2122ec719581cd177f225f59a1ee005831211196831f8f1ccffa817d2e7bd108.
//
// Solidity: event SessionTimeout(bool userActivated)
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

// ParseSessionTimeout is a log parse operation binding the contract event 0x2122ec719581cd177f225f59a1ee005831211196831f8f1ccffa817d2e7bd108.
//
// Solidity: event SessionTimeout(bool userActivated)
func (_RootChain *RootChainFilterer) ParseSessionTimeout(log types.Log) (*RootChainSessionTimeout, error) {
	event := new(RootChainSessionTimeout)
	if err := _RootChain.contract.UnpackLog(event, "SessionTimeout", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RootChainBaseABI is the input ABI used to generate the binding from.
const RootChainBaseABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"COST_URB_PREPARE\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CP_COMPUTATION\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastAppliedForkNumber\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"currentFork\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"COST_URB\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"numEnterForORB\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"EROIdToFinalize\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"forks\",\"outputs\":[{\"name\":\"forkedBlock\",\"type\":\"uint64\"},{\"name\":\"firstEpoch\",\"type\":\"uint64\"},{\"name\":\"lastEpoch\",\"type\":\"uint64\"},{\"name\":\"firstBlock\",\"type\":\"uint64\"},{\"name\":\"lastBlock\",\"type\":\"uint64\"},{\"name\":\"lastFinalizedEpoch\",\"type\":\"uint64\"},{\"name\":\"lastFinalizedBlock\",\"type\":\"uint64\"},{\"name\":\"timestamp\",\"type\":\"uint64\"},{\"name\":\"firstEnterEpoch\",\"type\":\"uint64\"},{\"name\":\"lastEnterEpoch\",\"type\":\"uint64\"},{\"name\":\"nextBlockToRebase\",\"type\":\"uint64\"},{\"name\":\"rebased\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"operator\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"firstFilledORENumber\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"development\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CP_EXIT\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"COST_ERU\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"REQUEST_GAS\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"COST_NRB\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"NRELength\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CP_WITHHOLDING\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"COST_ORB\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"EROs\",\"outputs\":[{\"name\":\"timestamp\",\"type\":\"uint64\"},{\"name\":\"isExit\",\"type\":\"bool\"},{\"name\":\"isTransfer\",\"type\":\"bool\"},{\"name\":\"finalized\",\"type\":\"bool\"},{\"name\":\"challenged\",\"type\":\"bool\"},{\"name\":\"value\",\"type\":\"uint128\"},{\"name\":\"requestor\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"trieKey\",\"type\":\"bytes32\"},{\"name\":\"hash\",\"type\":\"bytes32\"},{\"name\":\"trieValue\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"lastNonEmptyRequestEpoch\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"etherToken\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"URBs\",\"outputs\":[{\"name\":\"submitted\",\"type\":\"bool\"},{\"name\":\"numEnter\",\"type\":\"uint64\"},{\"name\":\"epochNumber\",\"type\":\"uint64\"},{\"name\":\"requestStart\",\"type\":\"uint64\"},{\"name\":\"requestEnd\",\"type\":\"uint64\"},{\"name\":\"trie\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PREPARE_TIMEOUT\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ERUIdToFinalize\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastAppliedEpochNumber\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"firstNonEmptyRequestEpoch\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"COST_ERO\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"requestableContracts\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"NULL_ADDRESS\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"submitHandler\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"epochHandler\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"ORBs\",\"outputs\":[{\"name\":\"submitted\",\"type\":\"bool\"},{\"name\":\"numEnter\",\"type\":\"uint64\"},{\"name\":\"epochNumber\",\"type\":\"uint64\"},{\"name\":\"requestStart\",\"type\":\"uint64\"},{\"name\":\"requestEnd\",\"type\":\"uint64\"},{\"name\":\"trie\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"ERUs\",\"outputs\":[{\"name\":\"timestamp\",\"type\":\"uint64\"},{\"name\":\"isExit\",\"type\":\"bool\"},{\"name\":\"isTransfer\",\"type\":\"bool\"},{\"name\":\"finalized\",\"type\":\"bool\"},{\"name\":\"challenged\",\"type\":\"bool\"},{\"name\":\"value\",\"type\":\"uint128\"},{\"name\":\"requestor\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"trieKey\",\"type\":\"bytes32\"},{\"name\":\"hash\",\"type\":\"bytes32\"},{\"name\":\"trieValue\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastAppliedBlockNumber\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"userActivated\",\"type\":\"bool\"}],\"name\":\"SessionTimeout\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newFork\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"epochNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"forkedBlockNumber\",\"type\":\"uint256\"}],\"name\":\"Forked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"forkNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"epochNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"startBlockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"endBlockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"requestStart\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"requestEnd\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"epochIsEmpty\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"isRequest\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"userActivated\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"rebase\",\"type\":\"bool\"}],\"name\":\"EpochPrepared\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"forkNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"epochNumber\",\"type\":\"uint256\"}],\"name\":\"EpochFilling\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"forkNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"epochNumber\",\"type\":\"uint256\"}],\"name\":\"EpochFilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"forkNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"epochNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"startBlockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"endBlockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"requestStart\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"requestEnd\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"epochIsEmpty\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"isRequest\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"userActivated\",\"type\":\"bool\"}],\"name\":\"EpochRebased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"fork\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"epochNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"isRequest\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"userActivated\",\"type\":\"bool\"}],\"name\":\"BlockSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"requestor\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"weiAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"trieKey\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"trieValue\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"isExit\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"userActivated\",\"type\":\"bool\"}],\"name\":\"RequestCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"requestor\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"trieKey\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"trieValue\",\"type\":\"bytes32\"}],\"name\":\"ERUCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"forkNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"BlockFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"forkNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"epochNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"startBlockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"endBlockNumber\",\"type\":\"uint256\"}],\"name\":\"EpochFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"userActivated\",\"type\":\"bool\"}],\"name\":\"RequestFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"userActivated\",\"type\":\"bool\"}],\"name\":\"RequestApplied\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"userActivated\",\"type\":\"bool\"}],\"name\":\"RequestChallenged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"contractInRootchain\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"contractInChildchain\",\"type\":\"address\"}],\"name\":\"RequestableContractMapped\",\"type\":\"event\"}]"

// RootChainBaseFuncSigs maps the 4-byte function signature to its string representation.
var RootChainBaseFuncSigs = map[string]string{
	"d691acd8": "COST_ERO()",
	"8b5172d0": "COST_ERU()",
	"94be3aa5": "COST_NRB()",
	"b2ae9ba8": "COST_ORB()",
	"192adc5b": "COST_URB()",
	"033cfbed": "COST_URB_PREPARE()",
	"08c4fff0": "CP_COMPUTATION()",
	"8155717d": "CP_EXIT()",
	"b17fa6e9": "CP_WITHHOLDING()",
	"2dc6bb7b": "EROIdToFinalize()",
	"b443f3cc": "EROs(uint256)",
	"c54626cc": "ERUIdToFinalize()",
	"f4f31de4": "ERUs(uint256)",
	"ab96da2d": "NRELength()",
	"de0ce17d": "NULL_ADDRESS()",
	"ea7f22a8": "ORBs(uint256)",
	"c2bc88fa": "PREPARE_TIMEOUT()",
	"8eb288ca": "REQUEST_GAS()",
	"c0e86064": "URBs(uint256)",
	"183d2d1c": "currentFork()",
	"7b929c27": "development()",
	"e7b88b80": "epochHandler()",
	"b8066bcb": "etherToken()",
	"72ecb9a8": "firstFilledORENumber(uint256)",
	"ca6f6380": "firstNonEmptyRequestEpoch(uint256)",
	"4ba3a126": "forks(uint256)",
	"fb788a27": "lastAppliedBlockNumber()",
	"c8ad329f": "lastAppliedEpochNumber()",
	"164bc2ae": "lastAppliedForkNumber()",
	"b6715647": "lastNonEmptyRequestEpoch(uint256)",
	"23691566": "numEnterForORB()",
	"570ca735": "operator()",
	"da0185f8": "requestableContracts(address)",
	"e259faf7": "submitHandler()",
}

// RootChainBaseBin is the compiled bytecode used for deploying new contracts.
var RootChainBaseBin = "0x608060405234801561001057600080fd5b5061091b806100206000396000f3fe608060405234801561001057600080fd5b50600436106102065760003560e01c8063b2ae9ba81161011a578063ca6f6380116100ad578063e259faf71161007c578063e259faf71461056f578063e7b88b8014610577578063ea7f22a81461057f578063f4f31de41461059c578063fb788a27146105b957610206565b8063ca6f638014610524578063d691acd81461020b578063da0185f814610541578063de0ce17d1461056757610206565b8063c0e86064116100e9578063c0e86064146104a3578063c2bc88fa1461050c578063c54626cc14610514578063c8ad329f1461051c57610206565b8063b2ae9ba81461020b578063b443f3cc1461036e578063b67156471461047e578063b8066bcb1461049b57610206565b8063570ca7351161019d5780638b5172d01161016c5780638b5172d01461034e5780638eb288ca1461035657806394be3aa51461020b578063ab96da2d1461035e578063b17fa6e91461036657610206565b8063570ca735146102e957806372ecb9a81461030d5780637b929c271461032a5780638155717d1461034657610206565b8063192adc5b116101d9578063192adc5b1461023d57806323691566146102455780632dc6bb7b1461024d5780634ba3a1261461025557610206565b8063033cfbed1461020b57806308c4fff014610225578063164bc2ae1461022d578063183d2d1c14610235575b600080fd5b6102136105c1565b60408051918252519081900360200190f35b6102136105cb565b6102136105d0565b6102136105d6565b6102136105dc565b6102136105e6565b6102136105ec565b6102726004803603602081101561026b57600080fd5b50356105f2565b6040805167ffffffffffffffff9d8e1681529b8d1660208d0152998c168b8b0152978b1660608b0152958a1660808a015293891660a089015291881660c0880152871660e0870152861661010086015285166101208501529093166101408301529115156101608201529051908190036101800190f35b6102f161066a565b604080516001600160a01b039092168252519081900360200190f35b6102136004803603602081101561032357600080fd5b503561067e565b610332610690565b604080519115158252519081900360200190f35b610213610699565b61021361069e565b6102136106a8565b6102136106af565b6102136106b5565b61038b6004803603602081101561038457600080fd5b50356106ba565b6040805167ffffffffffffffff8d1681528b15156020808301919091528b151592820192909252891515606082015288151560808201526001600160801b03881660a08201526001600160a01b0380881660c0830152861660e082015261010081018590526101208101849052610160610140820181815284519183019190915283519192909161018084019185019080838360005b83811015610439578181015183820152602001610421565b50505050905090810190601f1680156104665780820380516001836020036101000a031916815260200191505b509c5050505050505050505050505060405180910390f35b6102136004803603602081101561049457600080fd5b50356107dc565b6102f16107ee565b6104c0600480360360208110156104b957600080fd5b50356107fd565b60408051961515875267ffffffffffffffff958616602088015293851686850152918416606086015290921660808401526001600160a01b0390911660a0830152519081900360c00190f35b610213610865565b61021361086a565b610213610870565b6102136004803603602081101561053a57600080fd5b5035610876565b6102f16004803603602081101561055757600080fd5b50356001600160a01b0316610888565b6102f16108a3565b6102f16108a8565b6102f16108b7565b6104c06004803603602081101561059557600080fd5b50356108c6565b61038b600480360360208110156105b257600080fd5b50356108d3565b6102136108e0565b6509184e72a00081565b600f81565b600f5481565b60055481565b6551dac207a00081565b600c5481565b60125481565b60076020526000908152604090208054600182015460029092015467ffffffffffffffff80831693600160401b808504831694600160801b808204851695600160c01b92839004861695858116958581048216958482048316959182900483169484841694918204841693908204169160ff9104168c565b60005461010090046001600160a01b031681565b60066020526000908152604090205481565b60005460ff1681565b600a81565b6512309ce5400081565b620186a081565b60045481565b601481565b600881815481106106c757fe5b6000918252602091829020600691909102018054600180830154600280850154600386015460048701546005880180546040805161010099831615999099026000190190911695909504601f81018b90048b0288018b0190955284875267ffffffffffffffff88169a50600160401b880460ff9081169a600160481b8a0482169a600160501b8b0483169a600160581b810490931699600160601b9093046001600160801b0316986001600160a01b03908116989716969093918301828280156107d25780601f106107a7576101008083540402835291602001916107d2565b820191906000526020600020905b8154815290600101906020018083116107b557829003601f168201915b505050505090508b565b600d6020526000908152604090205481565b6003546001600160a01b031681565b600b818154811061080a57fe5b60009182526020909120600290910201805460019091015460ff8216925067ffffffffffffffff6101008304811692600160481b8104821692600160881b9091048216918116906001600160a01b03600160401b9091041686565b603c81565b60135481565b60105481565b600e6020526000908152604090205481565b6014602052600090815260409020546001600160a01b031681565b600081565b6002546001600160a01b031681565b6001546001600160a01b031681565b600a818154811061080a57fe5b600981815481106106c757fe5b6011548156fea265627a7a7230582014fc077e3ae750b6d424fa40d5949b94805eb51f509f71e0a24f4230ec5738e964736f6c63430005090032"

// DeployRootChainBase deploys a new Ethereum contract, binding an instance of RootChainBase to it.
func DeployRootChainBase(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RootChainBase, error) {
	parsed, err := abi.JSON(strings.NewReader(RootChainBaseABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RootChainBaseBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RootChainBase{RootChainBaseCaller: RootChainBaseCaller{contract: contract}, RootChainBaseTransactor: RootChainBaseTransactor{contract: contract}, RootChainBaseFilterer: RootChainBaseFilterer{contract: contract}}, nil
}

// RootChainBase is an auto generated Go binding around an Ethereum contract.
type RootChainBase struct {
	RootChainBaseCaller     // Read-only binding to the contract
	RootChainBaseTransactor // Write-only binding to the contract
	RootChainBaseFilterer   // Log filterer for contract events
}

// RootChainBaseCaller is an auto generated read-only Go binding around an Ethereum contract.
type RootChainBaseCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RootChainBaseTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RootChainBaseTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RootChainBaseFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RootChainBaseFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RootChainBaseSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RootChainBaseSession struct {
	Contract     *RootChainBase    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RootChainBaseCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RootChainBaseCallerSession struct {
	Contract *RootChainBaseCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// RootChainBaseTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RootChainBaseTransactorSession struct {
	Contract     *RootChainBaseTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// RootChainBaseRaw is an auto generated low-level Go binding around an Ethereum contract.
type RootChainBaseRaw struct {
	Contract *RootChainBase // Generic contract binding to access the raw methods on
}

// RootChainBaseCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RootChainBaseCallerRaw struct {
	Contract *RootChainBaseCaller // Generic read-only contract binding to access the raw methods on
}

// RootChainBaseTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RootChainBaseTransactorRaw struct {
	Contract *RootChainBaseTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRootChainBase creates a new instance of RootChainBase, bound to a specific deployed contract.
func NewRootChainBase(address common.Address, backend bind.ContractBackend) (*RootChainBase, error) {
	contract, err := bindRootChainBase(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RootChainBase{RootChainBaseCaller: RootChainBaseCaller{contract: contract}, RootChainBaseTransactor: RootChainBaseTransactor{contract: contract}, RootChainBaseFilterer: RootChainBaseFilterer{contract: contract}}, nil
}

// NewRootChainBaseCaller creates a new read-only instance of RootChainBase, bound to a specific deployed contract.
func NewRootChainBaseCaller(address common.Address, caller bind.ContractCaller) (*RootChainBaseCaller, error) {
	contract, err := bindRootChainBase(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RootChainBaseCaller{contract: contract}, nil
}

// NewRootChainBaseTransactor creates a new write-only instance of RootChainBase, bound to a specific deployed contract.
func NewRootChainBaseTransactor(address common.Address, transactor bind.ContractTransactor) (*RootChainBaseTransactor, error) {
	contract, err := bindRootChainBase(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RootChainBaseTransactor{contract: contract}, nil
}

// NewRootChainBaseFilterer creates a new log filterer instance of RootChainBase, bound to a specific deployed contract.
func NewRootChainBaseFilterer(address common.Address, filterer bind.ContractFilterer) (*RootChainBaseFilterer, error) {
	contract, err := bindRootChainBase(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RootChainBaseFilterer{contract: contract}, nil
}

// bindRootChainBase binds a generic wrapper to an already deployed contract.
func bindRootChainBase(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RootChainBaseABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RootChainBase *RootChainBaseRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RootChainBase.Contract.RootChainBaseCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RootChainBase *RootChainBaseRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RootChainBase.Contract.RootChainBaseTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RootChainBase *RootChainBaseRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RootChainBase.Contract.RootChainBaseTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RootChainBase *RootChainBaseCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RootChainBase.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RootChainBase *RootChainBaseTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RootChainBase.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RootChainBase *RootChainBaseTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RootChainBase.Contract.contract.Transact(opts, method, params...)
}

// COSTERO is a free data retrieval call binding the contract method 0xd691acd8.
//
// Solidity: function COST_ERO() constant returns(uint256)
func (_RootChainBase *RootChainBaseCaller) COSTERO(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChainBase.contract.Call(opts, out, "COST_ERO")
	return *ret0, err
}

// COSTERO is a free data retrieval call binding the contract method 0xd691acd8.
//
// Solidity: function COST_ERO() constant returns(uint256)
func (_RootChainBase *RootChainBaseSession) COSTERO() (*big.Int, error) {
	return _RootChainBase.Contract.COSTERO(&_RootChainBase.CallOpts)
}

// COSTERO is a free data retrieval call binding the contract method 0xd691acd8.
//
// Solidity: function COST_ERO() constant returns(uint256)
func (_RootChainBase *RootChainBaseCallerSession) COSTERO() (*big.Int, error) {
	return _RootChainBase.Contract.COSTERO(&_RootChainBase.CallOpts)
}

// COSTERU is a free data retrieval call binding the contract method 0x8b5172d0.
//
// Solidity: function COST_ERU() constant returns(uint256)
func (_RootChainBase *RootChainBaseCaller) COSTERU(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChainBase.contract.Call(opts, out, "COST_ERU")
	return *ret0, err
}

// COSTERU is a free data retrieval call binding the contract method 0x8b5172d0.
//
// Solidity: function COST_ERU() constant returns(uint256)
func (_RootChainBase *RootChainBaseSession) COSTERU() (*big.Int, error) {
	return _RootChainBase.Contract.COSTERU(&_RootChainBase.CallOpts)
}

// COSTERU is a free data retrieval call binding the contract method 0x8b5172d0.
//
// Solidity: function COST_ERU() constant returns(uint256)
func (_RootChainBase *RootChainBaseCallerSession) COSTERU() (*big.Int, error) {
	return _RootChainBase.Contract.COSTERU(&_RootChainBase.CallOpts)
}

// COSTNRB is a free data retrieval call binding the contract method 0x94be3aa5.
//
// Solidity: function COST_NRB() constant returns(uint256)
func (_RootChainBase *RootChainBaseCaller) COSTNRB(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChainBase.contract.Call(opts, out, "COST_NRB")
	return *ret0, err
}

// COSTNRB is a free data retrieval call binding the contract method 0x94be3aa5.
//
// Solidity: function COST_NRB() constant returns(uint256)
func (_RootChainBase *RootChainBaseSession) COSTNRB() (*big.Int, error) {
	return _RootChainBase.Contract.COSTNRB(&_RootChainBase.CallOpts)
}

// COSTNRB is a free data retrieval call binding the contract method 0x94be3aa5.
//
// Solidity: function COST_NRB() constant returns(uint256)
func (_RootChainBase *RootChainBaseCallerSession) COSTNRB() (*big.Int, error) {
	return _RootChainBase.Contract.COSTNRB(&_RootChainBase.CallOpts)
}

// COSTORB is a free data retrieval call binding the contract method 0xb2ae9ba8.
//
// Solidity: function COST_ORB() constant returns(uint256)
func (_RootChainBase *RootChainBaseCaller) COSTORB(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChainBase.contract.Call(opts, out, "COST_ORB")
	return *ret0, err
}

// COSTORB is a free data retrieval call binding the contract method 0xb2ae9ba8.
//
// Solidity: function COST_ORB() constant returns(uint256)
func (_RootChainBase *RootChainBaseSession) COSTORB() (*big.Int, error) {
	return _RootChainBase.Contract.COSTORB(&_RootChainBase.CallOpts)
}

// COSTORB is a free data retrieval call binding the contract method 0xb2ae9ba8.
//
// Solidity: function COST_ORB() constant returns(uint256)
func (_RootChainBase *RootChainBaseCallerSession) COSTORB() (*big.Int, error) {
	return _RootChainBase.Contract.COSTORB(&_RootChainBase.CallOpts)
}

// COSTURB is a free data retrieval call binding the contract method 0x192adc5b.
//
// Solidity: function COST_URB() constant returns(uint256)
func (_RootChainBase *RootChainBaseCaller) COSTURB(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChainBase.contract.Call(opts, out, "COST_URB")
	return *ret0, err
}

// COSTURB is a free data retrieval call binding the contract method 0x192adc5b.
//
// Solidity: function COST_URB() constant returns(uint256)
func (_RootChainBase *RootChainBaseSession) COSTURB() (*big.Int, error) {
	return _RootChainBase.Contract.COSTURB(&_RootChainBase.CallOpts)
}

// COSTURB is a free data retrieval call binding the contract method 0x192adc5b.
//
// Solidity: function COST_URB() constant returns(uint256)
func (_RootChainBase *RootChainBaseCallerSession) COSTURB() (*big.Int, error) {
	return _RootChainBase.Contract.COSTURB(&_RootChainBase.CallOpts)
}

// COSTURBPREPARE is a free data retrieval call binding the contract method 0x033cfbed.
//
// Solidity: function COST_URB_PREPARE() constant returns(uint256)
func (_RootChainBase *RootChainBaseCaller) COSTURBPREPARE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChainBase.contract.Call(opts, out, "COST_URB_PREPARE")
	return *ret0, err
}

// COSTURBPREPARE is a free data retrieval call binding the contract method 0x033cfbed.
//
// Solidity: function COST_URB_PREPARE() constant returns(uint256)
func (_RootChainBase *RootChainBaseSession) COSTURBPREPARE() (*big.Int, error) {
	return _RootChainBase.Contract.COSTURBPREPARE(&_RootChainBase.CallOpts)
}

// COSTURBPREPARE is a free data retrieval call binding the contract method 0x033cfbed.
//
// Solidity: function COST_URB_PREPARE() constant returns(uint256)
func (_RootChainBase *RootChainBaseCallerSession) COSTURBPREPARE() (*big.Int, error) {
	return _RootChainBase.Contract.COSTURBPREPARE(&_RootChainBase.CallOpts)
}

// CPCOMPUTATION is a free data retrieval call binding the contract method 0x08c4fff0.
//
// Solidity: function CP_COMPUTATION() constant returns(uint256)
func (_RootChainBase *RootChainBaseCaller) CPCOMPUTATION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChainBase.contract.Call(opts, out, "CP_COMPUTATION")
	return *ret0, err
}

// CPCOMPUTATION is a free data retrieval call binding the contract method 0x08c4fff0.
//
// Solidity: function CP_COMPUTATION() constant returns(uint256)
func (_RootChainBase *RootChainBaseSession) CPCOMPUTATION() (*big.Int, error) {
	return _RootChainBase.Contract.CPCOMPUTATION(&_RootChainBase.CallOpts)
}

// CPCOMPUTATION is a free data retrieval call binding the contract method 0x08c4fff0.
//
// Solidity: function CP_COMPUTATION() constant returns(uint256)
func (_RootChainBase *RootChainBaseCallerSession) CPCOMPUTATION() (*big.Int, error) {
	return _RootChainBase.Contract.CPCOMPUTATION(&_RootChainBase.CallOpts)
}

// CPEXIT is a free data retrieval call binding the contract method 0x8155717d.
//
// Solidity: function CP_EXIT() constant returns(uint256)
func (_RootChainBase *RootChainBaseCaller) CPEXIT(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChainBase.contract.Call(opts, out, "CP_EXIT")
	return *ret0, err
}

// CPEXIT is a free data retrieval call binding the contract method 0x8155717d.
//
// Solidity: function CP_EXIT() constant returns(uint256)
func (_RootChainBase *RootChainBaseSession) CPEXIT() (*big.Int, error) {
	return _RootChainBase.Contract.CPEXIT(&_RootChainBase.CallOpts)
}

// CPEXIT is a free data retrieval call binding the contract method 0x8155717d.
//
// Solidity: function CP_EXIT() constant returns(uint256)
func (_RootChainBase *RootChainBaseCallerSession) CPEXIT() (*big.Int, error) {
	return _RootChainBase.Contract.CPEXIT(&_RootChainBase.CallOpts)
}

// CPWITHHOLDING is a free data retrieval call binding the contract method 0xb17fa6e9.
//
// Solidity: function CP_WITHHOLDING() constant returns(uint256)
func (_RootChainBase *RootChainBaseCaller) CPWITHHOLDING(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChainBase.contract.Call(opts, out, "CP_WITHHOLDING")
	return *ret0, err
}

// CPWITHHOLDING is a free data retrieval call binding the contract method 0xb17fa6e9.
//
// Solidity: function CP_WITHHOLDING() constant returns(uint256)
func (_RootChainBase *RootChainBaseSession) CPWITHHOLDING() (*big.Int, error) {
	return _RootChainBase.Contract.CPWITHHOLDING(&_RootChainBase.CallOpts)
}

// CPWITHHOLDING is a free data retrieval call binding the contract method 0xb17fa6e9.
//
// Solidity: function CP_WITHHOLDING() constant returns(uint256)
func (_RootChainBase *RootChainBaseCallerSession) CPWITHHOLDING() (*big.Int, error) {
	return _RootChainBase.Contract.CPWITHHOLDING(&_RootChainBase.CallOpts)
}

// EROIdToFinalize is a free data retrieval call binding the contract method 0x2dc6bb7b.
//
// Solidity: function EROIdToFinalize() constant returns(uint256)
func (_RootChainBase *RootChainBaseCaller) EROIdToFinalize(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChainBase.contract.Call(opts, out, "EROIdToFinalize")
	return *ret0, err
}

// EROIdToFinalize is a free data retrieval call binding the contract method 0x2dc6bb7b.
//
// Solidity: function EROIdToFinalize() constant returns(uint256)
func (_RootChainBase *RootChainBaseSession) EROIdToFinalize() (*big.Int, error) {
	return _RootChainBase.Contract.EROIdToFinalize(&_RootChainBase.CallOpts)
}

// EROIdToFinalize is a free data retrieval call binding the contract method 0x2dc6bb7b.
//
// Solidity: function EROIdToFinalize() constant returns(uint256)
func (_RootChainBase *RootChainBaseCallerSession) EROIdToFinalize() (*big.Int, error) {
	return _RootChainBase.Contract.EROIdToFinalize(&_RootChainBase.CallOpts)
}

// EROs is a free data retrieval call binding the contract method 0xb443f3cc.
//
// Solidity: function EROs(uint256 ) constant returns(uint64 timestamp, bool isExit, bool isTransfer, bool finalized, bool challenged, uint128 value, address requestor, address to, bytes32 trieKey, bytes32 hash, bytes trieValue)
func (_RootChainBase *RootChainBaseCaller) EROs(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Timestamp  uint64
	IsExit     bool
	IsTransfer bool
	Finalized  bool
	Challenged bool
	Value      *big.Int
	Requestor  common.Address
	To         common.Address
	TrieKey    [32]byte
	Hash       [32]byte
	TrieValue  []byte
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
		Hash       [32]byte
		TrieValue  []byte
	})
	out := ret
	err := _RootChainBase.contract.Call(opts, out, "EROs", arg0)
	return *ret, err
}

// EROs is a free data retrieval call binding the contract method 0xb443f3cc.
//
// Solidity: function EROs(uint256 ) constant returns(uint64 timestamp, bool isExit, bool isTransfer, bool finalized, bool challenged, uint128 value, address requestor, address to, bytes32 trieKey, bytes32 hash, bytes trieValue)
func (_RootChainBase *RootChainBaseSession) EROs(arg0 *big.Int) (struct {
	Timestamp  uint64
	IsExit     bool
	IsTransfer bool
	Finalized  bool
	Challenged bool
	Value      *big.Int
	Requestor  common.Address
	To         common.Address
	TrieKey    [32]byte
	Hash       [32]byte
	TrieValue  []byte
}, error) {
	return _RootChainBase.Contract.EROs(&_RootChainBase.CallOpts, arg0)
}

// EROs is a free data retrieval call binding the contract method 0xb443f3cc.
//
// Solidity: function EROs(uint256 ) constant returns(uint64 timestamp, bool isExit, bool isTransfer, bool finalized, bool challenged, uint128 value, address requestor, address to, bytes32 trieKey, bytes32 hash, bytes trieValue)
func (_RootChainBase *RootChainBaseCallerSession) EROs(arg0 *big.Int) (struct {
	Timestamp  uint64
	IsExit     bool
	IsTransfer bool
	Finalized  bool
	Challenged bool
	Value      *big.Int
	Requestor  common.Address
	To         common.Address
	TrieKey    [32]byte
	Hash       [32]byte
	TrieValue  []byte
}, error) {
	return _RootChainBase.Contract.EROs(&_RootChainBase.CallOpts, arg0)
}

// ERUIdToFinalize is a free data retrieval call binding the contract method 0xc54626cc.
//
// Solidity: function ERUIdToFinalize() constant returns(uint256)
func (_RootChainBase *RootChainBaseCaller) ERUIdToFinalize(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChainBase.contract.Call(opts, out, "ERUIdToFinalize")
	return *ret0, err
}

// ERUIdToFinalize is a free data retrieval call binding the contract method 0xc54626cc.
//
// Solidity: function ERUIdToFinalize() constant returns(uint256)
func (_RootChainBase *RootChainBaseSession) ERUIdToFinalize() (*big.Int, error) {
	return _RootChainBase.Contract.ERUIdToFinalize(&_RootChainBase.CallOpts)
}

// ERUIdToFinalize is a free data retrieval call binding the contract method 0xc54626cc.
//
// Solidity: function ERUIdToFinalize() constant returns(uint256)
func (_RootChainBase *RootChainBaseCallerSession) ERUIdToFinalize() (*big.Int, error) {
	return _RootChainBase.Contract.ERUIdToFinalize(&_RootChainBase.CallOpts)
}

// ERUs is a free data retrieval call binding the contract method 0xf4f31de4.
//
// Solidity: function ERUs(uint256 ) constant returns(uint64 timestamp, bool isExit, bool isTransfer, bool finalized, bool challenged, uint128 value, address requestor, address to, bytes32 trieKey, bytes32 hash, bytes trieValue)
func (_RootChainBase *RootChainBaseCaller) ERUs(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Timestamp  uint64
	IsExit     bool
	IsTransfer bool
	Finalized  bool
	Challenged bool
	Value      *big.Int
	Requestor  common.Address
	To         common.Address
	TrieKey    [32]byte
	Hash       [32]byte
	TrieValue  []byte
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
		Hash       [32]byte
		TrieValue  []byte
	})
	out := ret
	err := _RootChainBase.contract.Call(opts, out, "ERUs", arg0)
	return *ret, err
}

// ERUs is a free data retrieval call binding the contract method 0xf4f31de4.
//
// Solidity: function ERUs(uint256 ) constant returns(uint64 timestamp, bool isExit, bool isTransfer, bool finalized, bool challenged, uint128 value, address requestor, address to, bytes32 trieKey, bytes32 hash, bytes trieValue)
func (_RootChainBase *RootChainBaseSession) ERUs(arg0 *big.Int) (struct {
	Timestamp  uint64
	IsExit     bool
	IsTransfer bool
	Finalized  bool
	Challenged bool
	Value      *big.Int
	Requestor  common.Address
	To         common.Address
	TrieKey    [32]byte
	Hash       [32]byte
	TrieValue  []byte
}, error) {
	return _RootChainBase.Contract.ERUs(&_RootChainBase.CallOpts, arg0)
}

// ERUs is a free data retrieval call binding the contract method 0xf4f31de4.
//
// Solidity: function ERUs(uint256 ) constant returns(uint64 timestamp, bool isExit, bool isTransfer, bool finalized, bool challenged, uint128 value, address requestor, address to, bytes32 trieKey, bytes32 hash, bytes trieValue)
func (_RootChainBase *RootChainBaseCallerSession) ERUs(arg0 *big.Int) (struct {
	Timestamp  uint64
	IsExit     bool
	IsTransfer bool
	Finalized  bool
	Challenged bool
	Value      *big.Int
	Requestor  common.Address
	To         common.Address
	TrieKey    [32]byte
	Hash       [32]byte
	TrieValue  []byte
}, error) {
	return _RootChainBase.Contract.ERUs(&_RootChainBase.CallOpts, arg0)
}

// NRELength is a free data retrieval call binding the contract method 0xab96da2d.
//
// Solidity: function NRELength() constant returns(uint256)
func (_RootChainBase *RootChainBaseCaller) NRELength(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChainBase.contract.Call(opts, out, "NRELength")
	return *ret0, err
}

// NRELength is a free data retrieval call binding the contract method 0xab96da2d.
//
// Solidity: function NRELength() constant returns(uint256)
func (_RootChainBase *RootChainBaseSession) NRELength() (*big.Int, error) {
	return _RootChainBase.Contract.NRELength(&_RootChainBase.CallOpts)
}

// NRELength is a free data retrieval call binding the contract method 0xab96da2d.
//
// Solidity: function NRELength() constant returns(uint256)
func (_RootChainBase *RootChainBaseCallerSession) NRELength() (*big.Int, error) {
	return _RootChainBase.Contract.NRELength(&_RootChainBase.CallOpts)
}

// NULLADDRESS is a free data retrieval call binding the contract method 0xde0ce17d.
//
// Solidity: function NULL_ADDRESS() constant returns(address)
func (_RootChainBase *RootChainBaseCaller) NULLADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RootChainBase.contract.Call(opts, out, "NULL_ADDRESS")
	return *ret0, err
}

// NULLADDRESS is a free data retrieval call binding the contract method 0xde0ce17d.
//
// Solidity: function NULL_ADDRESS() constant returns(address)
func (_RootChainBase *RootChainBaseSession) NULLADDRESS() (common.Address, error) {
	return _RootChainBase.Contract.NULLADDRESS(&_RootChainBase.CallOpts)
}

// NULLADDRESS is a free data retrieval call binding the contract method 0xde0ce17d.
//
// Solidity: function NULL_ADDRESS() constant returns(address)
func (_RootChainBase *RootChainBaseCallerSession) NULLADDRESS() (common.Address, error) {
	return _RootChainBase.Contract.NULLADDRESS(&_RootChainBase.CallOpts)
}

// ORBs is a free data retrieval call binding the contract method 0xea7f22a8.
//
// Solidity: function ORBs(uint256 ) constant returns(bool submitted, uint64 numEnter, uint64 epochNumber, uint64 requestStart, uint64 requestEnd, address trie)
func (_RootChainBase *RootChainBaseCaller) ORBs(opts *bind.CallOpts, arg0 *big.Int) (struct {
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
	err := _RootChainBase.contract.Call(opts, out, "ORBs", arg0)
	return *ret, err
}

// ORBs is a free data retrieval call binding the contract method 0xea7f22a8.
//
// Solidity: function ORBs(uint256 ) constant returns(bool submitted, uint64 numEnter, uint64 epochNumber, uint64 requestStart, uint64 requestEnd, address trie)
func (_RootChainBase *RootChainBaseSession) ORBs(arg0 *big.Int) (struct {
	Submitted    bool
	NumEnter     uint64
	EpochNumber  uint64
	RequestStart uint64
	RequestEnd   uint64
	Trie         common.Address
}, error) {
	return _RootChainBase.Contract.ORBs(&_RootChainBase.CallOpts, arg0)
}

// ORBs is a free data retrieval call binding the contract method 0xea7f22a8.
//
// Solidity: function ORBs(uint256 ) constant returns(bool submitted, uint64 numEnter, uint64 epochNumber, uint64 requestStart, uint64 requestEnd, address trie)
func (_RootChainBase *RootChainBaseCallerSession) ORBs(arg0 *big.Int) (struct {
	Submitted    bool
	NumEnter     uint64
	EpochNumber  uint64
	RequestStart uint64
	RequestEnd   uint64
	Trie         common.Address
}, error) {
	return _RootChainBase.Contract.ORBs(&_RootChainBase.CallOpts, arg0)
}

// PREPARETIMEOUT is a free data retrieval call binding the contract method 0xc2bc88fa.
//
// Solidity: function PREPARE_TIMEOUT() constant returns(uint256)
func (_RootChainBase *RootChainBaseCaller) PREPARETIMEOUT(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChainBase.contract.Call(opts, out, "PREPARE_TIMEOUT")
	return *ret0, err
}

// PREPARETIMEOUT is a free data retrieval call binding the contract method 0xc2bc88fa.
//
// Solidity: function PREPARE_TIMEOUT() constant returns(uint256)
func (_RootChainBase *RootChainBaseSession) PREPARETIMEOUT() (*big.Int, error) {
	return _RootChainBase.Contract.PREPARETIMEOUT(&_RootChainBase.CallOpts)
}

// PREPARETIMEOUT is a free data retrieval call binding the contract method 0xc2bc88fa.
//
// Solidity: function PREPARE_TIMEOUT() constant returns(uint256)
func (_RootChainBase *RootChainBaseCallerSession) PREPARETIMEOUT() (*big.Int, error) {
	return _RootChainBase.Contract.PREPARETIMEOUT(&_RootChainBase.CallOpts)
}

// REQUESTGAS is a free data retrieval call binding the contract method 0x8eb288ca.
//
// Solidity: function REQUEST_GAS() constant returns(uint256)
func (_RootChainBase *RootChainBaseCaller) REQUESTGAS(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChainBase.contract.Call(opts, out, "REQUEST_GAS")
	return *ret0, err
}

// REQUESTGAS is a free data retrieval call binding the contract method 0x8eb288ca.
//
// Solidity: function REQUEST_GAS() constant returns(uint256)
func (_RootChainBase *RootChainBaseSession) REQUESTGAS() (*big.Int, error) {
	return _RootChainBase.Contract.REQUESTGAS(&_RootChainBase.CallOpts)
}

// REQUESTGAS is a free data retrieval call binding the contract method 0x8eb288ca.
//
// Solidity: function REQUEST_GAS() constant returns(uint256)
func (_RootChainBase *RootChainBaseCallerSession) REQUESTGAS() (*big.Int, error) {
	return _RootChainBase.Contract.REQUESTGAS(&_RootChainBase.CallOpts)
}

// URBs is a free data retrieval call binding the contract method 0xc0e86064.
//
// Solidity: function URBs(uint256 ) constant returns(bool submitted, uint64 numEnter, uint64 epochNumber, uint64 requestStart, uint64 requestEnd, address trie)
func (_RootChainBase *RootChainBaseCaller) URBs(opts *bind.CallOpts, arg0 *big.Int) (struct {
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
	err := _RootChainBase.contract.Call(opts, out, "URBs", arg0)
	return *ret, err
}

// URBs is a free data retrieval call binding the contract method 0xc0e86064.
//
// Solidity: function URBs(uint256 ) constant returns(bool submitted, uint64 numEnter, uint64 epochNumber, uint64 requestStart, uint64 requestEnd, address trie)
func (_RootChainBase *RootChainBaseSession) URBs(arg0 *big.Int) (struct {
	Submitted    bool
	NumEnter     uint64
	EpochNumber  uint64
	RequestStart uint64
	RequestEnd   uint64
	Trie         common.Address
}, error) {
	return _RootChainBase.Contract.URBs(&_RootChainBase.CallOpts, arg0)
}

// URBs is a free data retrieval call binding the contract method 0xc0e86064.
//
// Solidity: function URBs(uint256 ) constant returns(bool submitted, uint64 numEnter, uint64 epochNumber, uint64 requestStart, uint64 requestEnd, address trie)
func (_RootChainBase *RootChainBaseCallerSession) URBs(arg0 *big.Int) (struct {
	Submitted    bool
	NumEnter     uint64
	EpochNumber  uint64
	RequestStart uint64
	RequestEnd   uint64
	Trie         common.Address
}, error) {
	return _RootChainBase.Contract.URBs(&_RootChainBase.CallOpts, arg0)
}

// CurrentFork is a free data retrieval call binding the contract method 0x183d2d1c.
//
// Solidity: function currentFork() constant returns(uint256)
func (_RootChainBase *RootChainBaseCaller) CurrentFork(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChainBase.contract.Call(opts, out, "currentFork")
	return *ret0, err
}

// CurrentFork is a free data retrieval call binding the contract method 0x183d2d1c.
//
// Solidity: function currentFork() constant returns(uint256)
func (_RootChainBase *RootChainBaseSession) CurrentFork() (*big.Int, error) {
	return _RootChainBase.Contract.CurrentFork(&_RootChainBase.CallOpts)
}

// CurrentFork is a free data retrieval call binding the contract method 0x183d2d1c.
//
// Solidity: function currentFork() constant returns(uint256)
func (_RootChainBase *RootChainBaseCallerSession) CurrentFork() (*big.Int, error) {
	return _RootChainBase.Contract.CurrentFork(&_RootChainBase.CallOpts)
}

// Development is a free data retrieval call binding the contract method 0x7b929c27.
//
// Solidity: function development() constant returns(bool)
func (_RootChainBase *RootChainBaseCaller) Development(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _RootChainBase.contract.Call(opts, out, "development")
	return *ret0, err
}

// Development is a free data retrieval call binding the contract method 0x7b929c27.
//
// Solidity: function development() constant returns(bool)
func (_RootChainBase *RootChainBaseSession) Development() (bool, error) {
	return _RootChainBase.Contract.Development(&_RootChainBase.CallOpts)
}

// Development is a free data retrieval call binding the contract method 0x7b929c27.
//
// Solidity: function development() constant returns(bool)
func (_RootChainBase *RootChainBaseCallerSession) Development() (bool, error) {
	return _RootChainBase.Contract.Development(&_RootChainBase.CallOpts)
}

// EpochHandler is a free data retrieval call binding the contract method 0xe7b88b80.
//
// Solidity: function epochHandler() constant returns(address)
func (_RootChainBase *RootChainBaseCaller) EpochHandler(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RootChainBase.contract.Call(opts, out, "epochHandler")
	return *ret0, err
}

// EpochHandler is a free data retrieval call binding the contract method 0xe7b88b80.
//
// Solidity: function epochHandler() constant returns(address)
func (_RootChainBase *RootChainBaseSession) EpochHandler() (common.Address, error) {
	return _RootChainBase.Contract.EpochHandler(&_RootChainBase.CallOpts)
}

// EpochHandler is a free data retrieval call binding the contract method 0xe7b88b80.
//
// Solidity: function epochHandler() constant returns(address)
func (_RootChainBase *RootChainBaseCallerSession) EpochHandler() (common.Address, error) {
	return _RootChainBase.Contract.EpochHandler(&_RootChainBase.CallOpts)
}

// EtherToken is a free data retrieval call binding the contract method 0xb8066bcb.
//
// Solidity: function etherToken() constant returns(address)
func (_RootChainBase *RootChainBaseCaller) EtherToken(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RootChainBase.contract.Call(opts, out, "etherToken")
	return *ret0, err
}

// EtherToken is a free data retrieval call binding the contract method 0xb8066bcb.
//
// Solidity: function etherToken() constant returns(address)
func (_RootChainBase *RootChainBaseSession) EtherToken() (common.Address, error) {
	return _RootChainBase.Contract.EtherToken(&_RootChainBase.CallOpts)
}

// EtherToken is a free data retrieval call binding the contract method 0xb8066bcb.
//
// Solidity: function etherToken() constant returns(address)
func (_RootChainBase *RootChainBaseCallerSession) EtherToken() (common.Address, error) {
	return _RootChainBase.Contract.EtherToken(&_RootChainBase.CallOpts)
}

// FirstFilledORENumber is a free data retrieval call binding the contract method 0x72ecb9a8.
//
// Solidity: function firstFilledORENumber(uint256 ) constant returns(uint256)
func (_RootChainBase *RootChainBaseCaller) FirstFilledORENumber(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChainBase.contract.Call(opts, out, "firstFilledORENumber", arg0)
	return *ret0, err
}

// FirstFilledORENumber is a free data retrieval call binding the contract method 0x72ecb9a8.
//
// Solidity: function firstFilledORENumber(uint256 ) constant returns(uint256)
func (_RootChainBase *RootChainBaseSession) FirstFilledORENumber(arg0 *big.Int) (*big.Int, error) {
	return _RootChainBase.Contract.FirstFilledORENumber(&_RootChainBase.CallOpts, arg0)
}

// FirstFilledORENumber is a free data retrieval call binding the contract method 0x72ecb9a8.
//
// Solidity: function firstFilledORENumber(uint256 ) constant returns(uint256)
func (_RootChainBase *RootChainBaseCallerSession) FirstFilledORENumber(arg0 *big.Int) (*big.Int, error) {
	return _RootChainBase.Contract.FirstFilledORENumber(&_RootChainBase.CallOpts, arg0)
}

// FirstNonEmptyRequestEpoch is a free data retrieval call binding the contract method 0xca6f6380.
//
// Solidity: function firstNonEmptyRequestEpoch(uint256 ) constant returns(uint256)
func (_RootChainBase *RootChainBaseCaller) FirstNonEmptyRequestEpoch(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChainBase.contract.Call(opts, out, "firstNonEmptyRequestEpoch", arg0)
	return *ret0, err
}

// FirstNonEmptyRequestEpoch is a free data retrieval call binding the contract method 0xca6f6380.
//
// Solidity: function firstNonEmptyRequestEpoch(uint256 ) constant returns(uint256)
func (_RootChainBase *RootChainBaseSession) FirstNonEmptyRequestEpoch(arg0 *big.Int) (*big.Int, error) {
	return _RootChainBase.Contract.FirstNonEmptyRequestEpoch(&_RootChainBase.CallOpts, arg0)
}

// FirstNonEmptyRequestEpoch is a free data retrieval call binding the contract method 0xca6f6380.
//
// Solidity: function firstNonEmptyRequestEpoch(uint256 ) constant returns(uint256)
func (_RootChainBase *RootChainBaseCallerSession) FirstNonEmptyRequestEpoch(arg0 *big.Int) (*big.Int, error) {
	return _RootChainBase.Contract.FirstNonEmptyRequestEpoch(&_RootChainBase.CallOpts, arg0)
}

// Forks is a free data retrieval call binding the contract method 0x4ba3a126.
//
// Solidity: function forks(uint256 ) constant returns(uint64 forkedBlock, uint64 firstEpoch, uint64 lastEpoch, uint64 firstBlock, uint64 lastBlock, uint64 lastFinalizedEpoch, uint64 lastFinalizedBlock, uint64 timestamp, uint64 firstEnterEpoch, uint64 lastEnterEpoch, uint64 nextBlockToRebase, bool rebased)
func (_RootChainBase *RootChainBaseCaller) Forks(opts *bind.CallOpts, arg0 *big.Int) (struct {
	ForkedBlock        uint64
	FirstEpoch         uint64
	LastEpoch          uint64
	FirstBlock         uint64
	LastBlock          uint64
	LastFinalizedEpoch uint64
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
		LastFinalizedEpoch uint64
		LastFinalizedBlock uint64
		Timestamp          uint64
		FirstEnterEpoch    uint64
		LastEnterEpoch     uint64
		NextBlockToRebase  uint64
		Rebased            bool
	})
	out := ret
	err := _RootChainBase.contract.Call(opts, out, "forks", arg0)
	return *ret, err
}

// Forks is a free data retrieval call binding the contract method 0x4ba3a126.
//
// Solidity: function forks(uint256 ) constant returns(uint64 forkedBlock, uint64 firstEpoch, uint64 lastEpoch, uint64 firstBlock, uint64 lastBlock, uint64 lastFinalizedEpoch, uint64 lastFinalizedBlock, uint64 timestamp, uint64 firstEnterEpoch, uint64 lastEnterEpoch, uint64 nextBlockToRebase, bool rebased)
func (_RootChainBase *RootChainBaseSession) Forks(arg0 *big.Int) (struct {
	ForkedBlock        uint64
	FirstEpoch         uint64
	LastEpoch          uint64
	FirstBlock         uint64
	LastBlock          uint64
	LastFinalizedEpoch uint64
	LastFinalizedBlock uint64
	Timestamp          uint64
	FirstEnterEpoch    uint64
	LastEnterEpoch     uint64
	NextBlockToRebase  uint64
	Rebased            bool
}, error) {
	return _RootChainBase.Contract.Forks(&_RootChainBase.CallOpts, arg0)
}

// Forks is a free data retrieval call binding the contract method 0x4ba3a126.
//
// Solidity: function forks(uint256 ) constant returns(uint64 forkedBlock, uint64 firstEpoch, uint64 lastEpoch, uint64 firstBlock, uint64 lastBlock, uint64 lastFinalizedEpoch, uint64 lastFinalizedBlock, uint64 timestamp, uint64 firstEnterEpoch, uint64 lastEnterEpoch, uint64 nextBlockToRebase, bool rebased)
func (_RootChainBase *RootChainBaseCallerSession) Forks(arg0 *big.Int) (struct {
	ForkedBlock        uint64
	FirstEpoch         uint64
	LastEpoch          uint64
	FirstBlock         uint64
	LastBlock          uint64
	LastFinalizedEpoch uint64
	LastFinalizedBlock uint64
	Timestamp          uint64
	FirstEnterEpoch    uint64
	LastEnterEpoch     uint64
	NextBlockToRebase  uint64
	Rebased            bool
}, error) {
	return _RootChainBase.Contract.Forks(&_RootChainBase.CallOpts, arg0)
}

// LastAppliedBlockNumber is a free data retrieval call binding the contract method 0xfb788a27.
//
// Solidity: function lastAppliedBlockNumber() constant returns(uint256)
func (_RootChainBase *RootChainBaseCaller) LastAppliedBlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChainBase.contract.Call(opts, out, "lastAppliedBlockNumber")
	return *ret0, err
}

// LastAppliedBlockNumber is a free data retrieval call binding the contract method 0xfb788a27.
//
// Solidity: function lastAppliedBlockNumber() constant returns(uint256)
func (_RootChainBase *RootChainBaseSession) LastAppliedBlockNumber() (*big.Int, error) {
	return _RootChainBase.Contract.LastAppliedBlockNumber(&_RootChainBase.CallOpts)
}

// LastAppliedBlockNumber is a free data retrieval call binding the contract method 0xfb788a27.
//
// Solidity: function lastAppliedBlockNumber() constant returns(uint256)
func (_RootChainBase *RootChainBaseCallerSession) LastAppliedBlockNumber() (*big.Int, error) {
	return _RootChainBase.Contract.LastAppliedBlockNumber(&_RootChainBase.CallOpts)
}

// LastAppliedEpochNumber is a free data retrieval call binding the contract method 0xc8ad329f.
//
// Solidity: function lastAppliedEpochNumber() constant returns(uint256)
func (_RootChainBase *RootChainBaseCaller) LastAppliedEpochNumber(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChainBase.contract.Call(opts, out, "lastAppliedEpochNumber")
	return *ret0, err
}

// LastAppliedEpochNumber is a free data retrieval call binding the contract method 0xc8ad329f.
//
// Solidity: function lastAppliedEpochNumber() constant returns(uint256)
func (_RootChainBase *RootChainBaseSession) LastAppliedEpochNumber() (*big.Int, error) {
	return _RootChainBase.Contract.LastAppliedEpochNumber(&_RootChainBase.CallOpts)
}

// LastAppliedEpochNumber is a free data retrieval call binding the contract method 0xc8ad329f.
//
// Solidity: function lastAppliedEpochNumber() constant returns(uint256)
func (_RootChainBase *RootChainBaseCallerSession) LastAppliedEpochNumber() (*big.Int, error) {
	return _RootChainBase.Contract.LastAppliedEpochNumber(&_RootChainBase.CallOpts)
}

// LastAppliedForkNumber is a free data retrieval call binding the contract method 0x164bc2ae.
//
// Solidity: function lastAppliedForkNumber() constant returns(uint256)
func (_RootChainBase *RootChainBaseCaller) LastAppliedForkNumber(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChainBase.contract.Call(opts, out, "lastAppliedForkNumber")
	return *ret0, err
}

// LastAppliedForkNumber is a free data retrieval call binding the contract method 0x164bc2ae.
//
// Solidity: function lastAppliedForkNumber() constant returns(uint256)
func (_RootChainBase *RootChainBaseSession) LastAppliedForkNumber() (*big.Int, error) {
	return _RootChainBase.Contract.LastAppliedForkNumber(&_RootChainBase.CallOpts)
}

// LastAppliedForkNumber is a free data retrieval call binding the contract method 0x164bc2ae.
//
// Solidity: function lastAppliedForkNumber() constant returns(uint256)
func (_RootChainBase *RootChainBaseCallerSession) LastAppliedForkNumber() (*big.Int, error) {
	return _RootChainBase.Contract.LastAppliedForkNumber(&_RootChainBase.CallOpts)
}

// LastNonEmptyRequestEpoch is a free data retrieval call binding the contract method 0xb6715647.
//
// Solidity: function lastNonEmptyRequestEpoch(uint256 ) constant returns(uint256)
func (_RootChainBase *RootChainBaseCaller) LastNonEmptyRequestEpoch(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChainBase.contract.Call(opts, out, "lastNonEmptyRequestEpoch", arg0)
	return *ret0, err
}

// LastNonEmptyRequestEpoch is a free data retrieval call binding the contract method 0xb6715647.
//
// Solidity: function lastNonEmptyRequestEpoch(uint256 ) constant returns(uint256)
func (_RootChainBase *RootChainBaseSession) LastNonEmptyRequestEpoch(arg0 *big.Int) (*big.Int, error) {
	return _RootChainBase.Contract.LastNonEmptyRequestEpoch(&_RootChainBase.CallOpts, arg0)
}

// LastNonEmptyRequestEpoch is a free data retrieval call binding the contract method 0xb6715647.
//
// Solidity: function lastNonEmptyRequestEpoch(uint256 ) constant returns(uint256)
func (_RootChainBase *RootChainBaseCallerSession) LastNonEmptyRequestEpoch(arg0 *big.Int) (*big.Int, error) {
	return _RootChainBase.Contract.LastNonEmptyRequestEpoch(&_RootChainBase.CallOpts, arg0)
}

// NumEnterForORB is a free data retrieval call binding the contract method 0x23691566.
//
// Solidity: function numEnterForORB() constant returns(uint256)
func (_RootChainBase *RootChainBaseCaller) NumEnterForORB(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChainBase.contract.Call(opts, out, "numEnterForORB")
	return *ret0, err
}

// NumEnterForORB is a free data retrieval call binding the contract method 0x23691566.
//
// Solidity: function numEnterForORB() constant returns(uint256)
func (_RootChainBase *RootChainBaseSession) NumEnterForORB() (*big.Int, error) {
	return _RootChainBase.Contract.NumEnterForORB(&_RootChainBase.CallOpts)
}

// NumEnterForORB is a free data retrieval call binding the contract method 0x23691566.
//
// Solidity: function numEnterForORB() constant returns(uint256)
func (_RootChainBase *RootChainBaseCallerSession) NumEnterForORB() (*big.Int, error) {
	return _RootChainBase.Contract.NumEnterForORB(&_RootChainBase.CallOpts)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() constant returns(address)
func (_RootChainBase *RootChainBaseCaller) Operator(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RootChainBase.contract.Call(opts, out, "operator")
	return *ret0, err
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() constant returns(address)
func (_RootChainBase *RootChainBaseSession) Operator() (common.Address, error) {
	return _RootChainBase.Contract.Operator(&_RootChainBase.CallOpts)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() constant returns(address)
func (_RootChainBase *RootChainBaseCallerSession) Operator() (common.Address, error) {
	return _RootChainBase.Contract.Operator(&_RootChainBase.CallOpts)
}

// RequestableContracts is a free data retrieval call binding the contract method 0xda0185f8.
//
// Solidity: function requestableContracts(address ) constant returns(address)
func (_RootChainBase *RootChainBaseCaller) RequestableContracts(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RootChainBase.contract.Call(opts, out, "requestableContracts", arg0)
	return *ret0, err
}

// RequestableContracts is a free data retrieval call binding the contract method 0xda0185f8.
//
// Solidity: function requestableContracts(address ) constant returns(address)
func (_RootChainBase *RootChainBaseSession) RequestableContracts(arg0 common.Address) (common.Address, error) {
	return _RootChainBase.Contract.RequestableContracts(&_RootChainBase.CallOpts, arg0)
}

// RequestableContracts is a free data retrieval call binding the contract method 0xda0185f8.
//
// Solidity: function requestableContracts(address ) constant returns(address)
func (_RootChainBase *RootChainBaseCallerSession) RequestableContracts(arg0 common.Address) (common.Address, error) {
	return _RootChainBase.Contract.RequestableContracts(&_RootChainBase.CallOpts, arg0)
}

// SubmitHandler is a free data retrieval call binding the contract method 0xe259faf7.
//
// Solidity: function submitHandler() constant returns(address)
func (_RootChainBase *RootChainBaseCaller) SubmitHandler(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RootChainBase.contract.Call(opts, out, "submitHandler")
	return *ret0, err
}

// SubmitHandler is a free data retrieval call binding the contract method 0xe259faf7.
//
// Solidity: function submitHandler() constant returns(address)
func (_RootChainBase *RootChainBaseSession) SubmitHandler() (common.Address, error) {
	return _RootChainBase.Contract.SubmitHandler(&_RootChainBase.CallOpts)
}

// SubmitHandler is a free data retrieval call binding the contract method 0xe259faf7.
//
// Solidity: function submitHandler() constant returns(address)
func (_RootChainBase *RootChainBaseCallerSession) SubmitHandler() (common.Address, error) {
	return _RootChainBase.Contract.SubmitHandler(&_RootChainBase.CallOpts)
}

// RootChainBaseBlockFinalizedIterator is returned from FilterBlockFinalized and is used to iterate over the raw logs and unpacked data for BlockFinalized events raised by the RootChainBase contract.
type RootChainBaseBlockFinalizedIterator struct {
	Event *RootChainBaseBlockFinalized // Event containing the contract specifics and raw log

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
func (it *RootChainBaseBlockFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RootChainBaseBlockFinalized)
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
		it.Event = new(RootChainBaseBlockFinalized)
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
func (it *RootChainBaseBlockFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RootChainBaseBlockFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RootChainBaseBlockFinalized represents a BlockFinalized event raised by the RootChainBase contract.
type RootChainBaseBlockFinalized struct {
	ForkNumber  *big.Int
	BlockNumber *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBlockFinalized is a free log retrieval operation binding the contract event 0xfb96205e4b3633fd57aa805b26b51ecf528714a10241a4af015929dce86768d9.
//
// Solidity: event BlockFinalized(uint256 forkNumber, uint256 blockNumber)
func (_RootChainBase *RootChainBaseFilterer) FilterBlockFinalized(opts *bind.FilterOpts) (*RootChainBaseBlockFinalizedIterator, error) {

	logs, sub, err := _RootChainBase.contract.FilterLogs(opts, "BlockFinalized")
	if err != nil {
		return nil, err
	}
	return &RootChainBaseBlockFinalizedIterator{contract: _RootChainBase.contract, event: "BlockFinalized", logs: logs, sub: sub}, nil
}

// WatchBlockFinalized is a free log subscription operation binding the contract event 0xfb96205e4b3633fd57aa805b26b51ecf528714a10241a4af015929dce86768d9.
//
// Solidity: event BlockFinalized(uint256 forkNumber, uint256 blockNumber)
func (_RootChainBase *RootChainBaseFilterer) WatchBlockFinalized(opts *bind.WatchOpts, sink chan<- *RootChainBaseBlockFinalized) (event.Subscription, error) {

	logs, sub, err := _RootChainBase.contract.WatchLogs(opts, "BlockFinalized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RootChainBaseBlockFinalized)
				if err := _RootChainBase.contract.UnpackLog(event, "BlockFinalized", log); err != nil {
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

// ParseBlockFinalized is a log parse operation binding the contract event 0xfb96205e4b3633fd57aa805b26b51ecf528714a10241a4af015929dce86768d9.
//
// Solidity: event BlockFinalized(uint256 forkNumber, uint256 blockNumber)
func (_RootChainBase *RootChainBaseFilterer) ParseBlockFinalized(log types.Log) (*RootChainBaseBlockFinalized, error) {
	event := new(RootChainBaseBlockFinalized)
	if err := _RootChainBase.contract.UnpackLog(event, "BlockFinalized", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RootChainBaseBlockSubmittedIterator is returned from FilterBlockSubmitted and is used to iterate over the raw logs and unpacked data for BlockSubmitted events raised by the RootChainBase contract.
type RootChainBaseBlockSubmittedIterator struct {
	Event *RootChainBaseBlockSubmitted // Event containing the contract specifics and raw log

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
func (it *RootChainBaseBlockSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RootChainBaseBlockSubmitted)
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
		it.Event = new(RootChainBaseBlockSubmitted)
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
func (it *RootChainBaseBlockSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RootChainBaseBlockSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RootChainBaseBlockSubmitted represents a BlockSubmitted event raised by the RootChainBase contract.
type RootChainBaseBlockSubmitted struct {
	Fork          *big.Int
	EpochNumber   *big.Int
	BlockNumber   *big.Int
	IsRequest     bool
	UserActivated bool
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterBlockSubmitted is a free log retrieval operation binding the contract event 0x3d4a04291c66b06f39a4ecb817875b12b5485a05ec563133a56a905305c48e55.
//
// Solidity: event BlockSubmitted(uint256 fork, uint256 epochNumber, uint256 blockNumber, bool isRequest, bool userActivated)
func (_RootChainBase *RootChainBaseFilterer) FilterBlockSubmitted(opts *bind.FilterOpts) (*RootChainBaseBlockSubmittedIterator, error) {

	logs, sub, err := _RootChainBase.contract.FilterLogs(opts, "BlockSubmitted")
	if err != nil {
		return nil, err
	}
	return &RootChainBaseBlockSubmittedIterator{contract: _RootChainBase.contract, event: "BlockSubmitted", logs: logs, sub: sub}, nil
}

// WatchBlockSubmitted is a free log subscription operation binding the contract event 0x3d4a04291c66b06f39a4ecb817875b12b5485a05ec563133a56a905305c48e55.
//
// Solidity: event BlockSubmitted(uint256 fork, uint256 epochNumber, uint256 blockNumber, bool isRequest, bool userActivated)
func (_RootChainBase *RootChainBaseFilterer) WatchBlockSubmitted(opts *bind.WatchOpts, sink chan<- *RootChainBaseBlockSubmitted) (event.Subscription, error) {

	logs, sub, err := _RootChainBase.contract.WatchLogs(opts, "BlockSubmitted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RootChainBaseBlockSubmitted)
				if err := _RootChainBase.contract.UnpackLog(event, "BlockSubmitted", log); err != nil {
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

// ParseBlockSubmitted is a log parse operation binding the contract event 0x3d4a04291c66b06f39a4ecb817875b12b5485a05ec563133a56a905305c48e55.
//
// Solidity: event BlockSubmitted(uint256 fork, uint256 epochNumber, uint256 blockNumber, bool isRequest, bool userActivated)
func (_RootChainBase *RootChainBaseFilterer) ParseBlockSubmitted(log types.Log) (*RootChainBaseBlockSubmitted, error) {
	event := new(RootChainBaseBlockSubmitted)
	if err := _RootChainBase.contract.UnpackLog(event, "BlockSubmitted", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RootChainBaseERUCreatedIterator is returned from FilterERUCreated and is used to iterate over the raw logs and unpacked data for ERUCreated events raised by the RootChainBase contract.
type RootChainBaseERUCreatedIterator struct {
	Event *RootChainBaseERUCreated // Event containing the contract specifics and raw log

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
func (it *RootChainBaseERUCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RootChainBaseERUCreated)
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
		it.Event = new(RootChainBaseERUCreated)
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
func (it *RootChainBaseERUCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RootChainBaseERUCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RootChainBaseERUCreated represents a ERUCreated event raised by the RootChainBase contract.
type RootChainBaseERUCreated struct {
	RequestId *big.Int
	Requestor common.Address
	To        common.Address
	TrieKey   []byte
	TrieValue [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterERUCreated is a free log retrieval operation binding the contract event 0xfcbdc2083dadd644b854d91b49aef8db06b8f5a3d5c1192de38ca0ba271d5a0d.
//
// Solidity: event ERUCreated(uint256 requestId, address requestor, address to, bytes trieKey, bytes32 trieValue)
func (_RootChainBase *RootChainBaseFilterer) FilterERUCreated(opts *bind.FilterOpts) (*RootChainBaseERUCreatedIterator, error) {

	logs, sub, err := _RootChainBase.contract.FilterLogs(opts, "ERUCreated")
	if err != nil {
		return nil, err
	}
	return &RootChainBaseERUCreatedIterator{contract: _RootChainBase.contract, event: "ERUCreated", logs: logs, sub: sub}, nil
}

// WatchERUCreated is a free log subscription operation binding the contract event 0xfcbdc2083dadd644b854d91b49aef8db06b8f5a3d5c1192de38ca0ba271d5a0d.
//
// Solidity: event ERUCreated(uint256 requestId, address requestor, address to, bytes trieKey, bytes32 trieValue)
func (_RootChainBase *RootChainBaseFilterer) WatchERUCreated(opts *bind.WatchOpts, sink chan<- *RootChainBaseERUCreated) (event.Subscription, error) {

	logs, sub, err := _RootChainBase.contract.WatchLogs(opts, "ERUCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RootChainBaseERUCreated)
				if err := _RootChainBase.contract.UnpackLog(event, "ERUCreated", log); err != nil {
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

// ParseERUCreated is a log parse operation binding the contract event 0xfcbdc2083dadd644b854d91b49aef8db06b8f5a3d5c1192de38ca0ba271d5a0d.
//
// Solidity: event ERUCreated(uint256 requestId, address requestor, address to, bytes trieKey, bytes32 trieValue)
func (_RootChainBase *RootChainBaseFilterer) ParseERUCreated(log types.Log) (*RootChainBaseERUCreated, error) {
	event := new(RootChainBaseERUCreated)
	if err := _RootChainBase.contract.UnpackLog(event, "ERUCreated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RootChainBaseEpochFilledIterator is returned from FilterEpochFilled and is used to iterate over the raw logs and unpacked data for EpochFilled events raised by the RootChainBase contract.
type RootChainBaseEpochFilledIterator struct {
	Event *RootChainBaseEpochFilled // Event containing the contract specifics and raw log

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
func (it *RootChainBaseEpochFilledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RootChainBaseEpochFilled)
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
		it.Event = new(RootChainBaseEpochFilled)
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
func (it *RootChainBaseEpochFilledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RootChainBaseEpochFilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RootChainBaseEpochFilled represents a EpochFilled event raised by the RootChainBase contract.
type RootChainBaseEpochFilled struct {
	ForkNumber  *big.Int
	EpochNumber *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterEpochFilled is a free log retrieval operation binding the contract event 0x2fdeb407bf5c2b621f04b5c784822dae806c45b49a68aba413cc270128c96816.
//
// Solidity: event EpochFilled(uint256 forkNumber, uint256 epochNumber)
func (_RootChainBase *RootChainBaseFilterer) FilterEpochFilled(opts *bind.FilterOpts) (*RootChainBaseEpochFilledIterator, error) {

	logs, sub, err := _RootChainBase.contract.FilterLogs(opts, "EpochFilled")
	if err != nil {
		return nil, err
	}
	return &RootChainBaseEpochFilledIterator{contract: _RootChainBase.contract, event: "EpochFilled", logs: logs, sub: sub}, nil
}

// WatchEpochFilled is a free log subscription operation binding the contract event 0x2fdeb407bf5c2b621f04b5c784822dae806c45b49a68aba413cc270128c96816.
//
// Solidity: event EpochFilled(uint256 forkNumber, uint256 epochNumber)
func (_RootChainBase *RootChainBaseFilterer) WatchEpochFilled(opts *bind.WatchOpts, sink chan<- *RootChainBaseEpochFilled) (event.Subscription, error) {

	logs, sub, err := _RootChainBase.contract.WatchLogs(opts, "EpochFilled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RootChainBaseEpochFilled)
				if err := _RootChainBase.contract.UnpackLog(event, "EpochFilled", log); err != nil {
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

// ParseEpochFilled is a log parse operation binding the contract event 0x2fdeb407bf5c2b621f04b5c784822dae806c45b49a68aba413cc270128c96816.
//
// Solidity: event EpochFilled(uint256 forkNumber, uint256 epochNumber)
func (_RootChainBase *RootChainBaseFilterer) ParseEpochFilled(log types.Log) (*RootChainBaseEpochFilled, error) {
	event := new(RootChainBaseEpochFilled)
	if err := _RootChainBase.contract.UnpackLog(event, "EpochFilled", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RootChainBaseEpochFillingIterator is returned from FilterEpochFilling and is used to iterate over the raw logs and unpacked data for EpochFilling events raised by the RootChainBase contract.
type RootChainBaseEpochFillingIterator struct {
	Event *RootChainBaseEpochFilling // Event containing the contract specifics and raw log

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
func (it *RootChainBaseEpochFillingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RootChainBaseEpochFilling)
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
		it.Event = new(RootChainBaseEpochFilling)
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
func (it *RootChainBaseEpochFillingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RootChainBaseEpochFillingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RootChainBaseEpochFilling represents a EpochFilling event raised by the RootChainBase contract.
type RootChainBaseEpochFilling struct {
	ForkNumber  *big.Int
	EpochNumber *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterEpochFilling is a free log retrieval operation binding the contract event 0x27b09f0953d27bbff306fe25b2987ac5a813248ac30cb2bbd5daf95e7b0e6dc0.
//
// Solidity: event EpochFilling(uint256 forkNumber, uint256 epochNumber)
func (_RootChainBase *RootChainBaseFilterer) FilterEpochFilling(opts *bind.FilterOpts) (*RootChainBaseEpochFillingIterator, error) {

	logs, sub, err := _RootChainBase.contract.FilterLogs(opts, "EpochFilling")
	if err != nil {
		return nil, err
	}
	return &RootChainBaseEpochFillingIterator{contract: _RootChainBase.contract, event: "EpochFilling", logs: logs, sub: sub}, nil
}

// WatchEpochFilling is a free log subscription operation binding the contract event 0x27b09f0953d27bbff306fe25b2987ac5a813248ac30cb2bbd5daf95e7b0e6dc0.
//
// Solidity: event EpochFilling(uint256 forkNumber, uint256 epochNumber)
func (_RootChainBase *RootChainBaseFilterer) WatchEpochFilling(opts *bind.WatchOpts, sink chan<- *RootChainBaseEpochFilling) (event.Subscription, error) {

	logs, sub, err := _RootChainBase.contract.WatchLogs(opts, "EpochFilling")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RootChainBaseEpochFilling)
				if err := _RootChainBase.contract.UnpackLog(event, "EpochFilling", log); err != nil {
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

// ParseEpochFilling is a log parse operation binding the contract event 0x27b09f0953d27bbff306fe25b2987ac5a813248ac30cb2bbd5daf95e7b0e6dc0.
//
// Solidity: event EpochFilling(uint256 forkNumber, uint256 epochNumber)
func (_RootChainBase *RootChainBaseFilterer) ParseEpochFilling(log types.Log) (*RootChainBaseEpochFilling, error) {
	event := new(RootChainBaseEpochFilling)
	if err := _RootChainBase.contract.UnpackLog(event, "EpochFilling", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RootChainBaseEpochFinalizedIterator is returned from FilterEpochFinalized and is used to iterate over the raw logs and unpacked data for EpochFinalized events raised by the RootChainBase contract.
type RootChainBaseEpochFinalizedIterator struct {
	Event *RootChainBaseEpochFinalized // Event containing the contract specifics and raw log

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
func (it *RootChainBaseEpochFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RootChainBaseEpochFinalized)
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
		it.Event = new(RootChainBaseEpochFinalized)
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
func (it *RootChainBaseEpochFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RootChainBaseEpochFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RootChainBaseEpochFinalized represents a EpochFinalized event raised by the RootChainBase contract.
type RootChainBaseEpochFinalized struct {
	ForkNumber       *big.Int
	EpochNumber      *big.Int
	StartBlockNumber *big.Int
	EndBlockNumber   *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterEpochFinalized is a free log retrieval operation binding the contract event 0x70801d4d63b3da6c19ba7349911f45bed5a99ccdfb51b8138c105872529bebd5.
//
// Solidity: event EpochFinalized(uint256 forkNumber, uint256 epochNumber, uint256 startBlockNumber, uint256 endBlockNumber)
func (_RootChainBase *RootChainBaseFilterer) FilterEpochFinalized(opts *bind.FilterOpts) (*RootChainBaseEpochFinalizedIterator, error) {

	logs, sub, err := _RootChainBase.contract.FilterLogs(opts, "EpochFinalized")
	if err != nil {
		return nil, err
	}
	return &RootChainBaseEpochFinalizedIterator{contract: _RootChainBase.contract, event: "EpochFinalized", logs: logs, sub: sub}, nil
}

// WatchEpochFinalized is a free log subscription operation binding the contract event 0x70801d4d63b3da6c19ba7349911f45bed5a99ccdfb51b8138c105872529bebd5.
//
// Solidity: event EpochFinalized(uint256 forkNumber, uint256 epochNumber, uint256 startBlockNumber, uint256 endBlockNumber)
func (_RootChainBase *RootChainBaseFilterer) WatchEpochFinalized(opts *bind.WatchOpts, sink chan<- *RootChainBaseEpochFinalized) (event.Subscription, error) {

	logs, sub, err := _RootChainBase.contract.WatchLogs(opts, "EpochFinalized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RootChainBaseEpochFinalized)
				if err := _RootChainBase.contract.UnpackLog(event, "EpochFinalized", log); err != nil {
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

// ParseEpochFinalized is a log parse operation binding the contract event 0x70801d4d63b3da6c19ba7349911f45bed5a99ccdfb51b8138c105872529bebd5.
//
// Solidity: event EpochFinalized(uint256 forkNumber, uint256 epochNumber, uint256 startBlockNumber, uint256 endBlockNumber)
func (_RootChainBase *RootChainBaseFilterer) ParseEpochFinalized(log types.Log) (*RootChainBaseEpochFinalized, error) {
	event := new(RootChainBaseEpochFinalized)
	if err := _RootChainBase.contract.UnpackLog(event, "EpochFinalized", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RootChainBaseEpochPreparedIterator is returned from FilterEpochPrepared and is used to iterate over the raw logs and unpacked data for EpochPrepared events raised by the RootChainBase contract.
type RootChainBaseEpochPreparedIterator struct {
	Event *RootChainBaseEpochPrepared // Event containing the contract specifics and raw log

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
func (it *RootChainBaseEpochPreparedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RootChainBaseEpochPrepared)
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
		it.Event = new(RootChainBaseEpochPrepared)
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
func (it *RootChainBaseEpochPreparedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RootChainBaseEpochPreparedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RootChainBaseEpochPrepared represents a EpochPrepared event raised by the RootChainBase contract.
type RootChainBaseEpochPrepared struct {
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
// Solidity: event EpochPrepared(uint256 forkNumber, uint256 epochNumber, uint256 startBlockNumber, uint256 endBlockNumber, uint256 requestStart, uint256 requestEnd, bool epochIsEmpty, bool isRequest, bool userActivated, bool rebase)
func (_RootChainBase *RootChainBaseFilterer) FilterEpochPrepared(opts *bind.FilterOpts) (*RootChainBaseEpochPreparedIterator, error) {

	logs, sub, err := _RootChainBase.contract.FilterLogs(opts, "EpochPrepared")
	if err != nil {
		return nil, err
	}
	return &RootChainBaseEpochPreparedIterator{contract: _RootChainBase.contract, event: "EpochPrepared", logs: logs, sub: sub}, nil
}

// WatchEpochPrepared is a free log subscription operation binding the contract event 0x1a69c0760aa329b76f72579129869013ebd3d41594db019c0e997b939fcb32e3.
//
// Solidity: event EpochPrepared(uint256 forkNumber, uint256 epochNumber, uint256 startBlockNumber, uint256 endBlockNumber, uint256 requestStart, uint256 requestEnd, bool epochIsEmpty, bool isRequest, bool userActivated, bool rebase)
func (_RootChainBase *RootChainBaseFilterer) WatchEpochPrepared(opts *bind.WatchOpts, sink chan<- *RootChainBaseEpochPrepared) (event.Subscription, error) {

	logs, sub, err := _RootChainBase.contract.WatchLogs(opts, "EpochPrepared")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RootChainBaseEpochPrepared)
				if err := _RootChainBase.contract.UnpackLog(event, "EpochPrepared", log); err != nil {
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

// ParseEpochPrepared is a log parse operation binding the contract event 0x1a69c0760aa329b76f72579129869013ebd3d41594db019c0e997b939fcb32e3.
//
// Solidity: event EpochPrepared(uint256 forkNumber, uint256 epochNumber, uint256 startBlockNumber, uint256 endBlockNumber, uint256 requestStart, uint256 requestEnd, bool epochIsEmpty, bool isRequest, bool userActivated, bool rebase)
func (_RootChainBase *RootChainBaseFilterer) ParseEpochPrepared(log types.Log) (*RootChainBaseEpochPrepared, error) {
	event := new(RootChainBaseEpochPrepared)
	if err := _RootChainBase.contract.UnpackLog(event, "EpochPrepared", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RootChainBaseEpochRebasedIterator is returned from FilterEpochRebased and is used to iterate over the raw logs and unpacked data for EpochRebased events raised by the RootChainBase contract.
type RootChainBaseEpochRebasedIterator struct {
	Event *RootChainBaseEpochRebased // Event containing the contract specifics and raw log

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
func (it *RootChainBaseEpochRebasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RootChainBaseEpochRebased)
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
		it.Event = new(RootChainBaseEpochRebased)
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
func (it *RootChainBaseEpochRebasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RootChainBaseEpochRebasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RootChainBaseEpochRebased represents a EpochRebased event raised by the RootChainBase contract.
type RootChainBaseEpochRebased struct {
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
// Solidity: event EpochRebased(uint256 forkNumber, uint256 epochNumber, uint256 startBlockNumber, uint256 endBlockNumber, uint256 requestStart, uint256 requestEnd, bool epochIsEmpty, bool isRequest, bool userActivated)
func (_RootChainBase *RootChainBaseFilterer) FilterEpochRebased(opts *bind.FilterOpts) (*RootChainBaseEpochRebasedIterator, error) {

	logs, sub, err := _RootChainBase.contract.FilterLogs(opts, "EpochRebased")
	if err != nil {
		return nil, err
	}
	return &RootChainBaseEpochRebasedIterator{contract: _RootChainBase.contract, event: "EpochRebased", logs: logs, sub: sub}, nil
}

// WatchEpochRebased is a free log subscription operation binding the contract event 0x030c1c69405c93021f28f57557240dee939a320b826a1fd0d39bf6e629ecab47.
//
// Solidity: event EpochRebased(uint256 forkNumber, uint256 epochNumber, uint256 startBlockNumber, uint256 endBlockNumber, uint256 requestStart, uint256 requestEnd, bool epochIsEmpty, bool isRequest, bool userActivated)
func (_RootChainBase *RootChainBaseFilterer) WatchEpochRebased(opts *bind.WatchOpts, sink chan<- *RootChainBaseEpochRebased) (event.Subscription, error) {

	logs, sub, err := _RootChainBase.contract.WatchLogs(opts, "EpochRebased")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RootChainBaseEpochRebased)
				if err := _RootChainBase.contract.UnpackLog(event, "EpochRebased", log); err != nil {
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

// ParseEpochRebased is a log parse operation binding the contract event 0x030c1c69405c93021f28f57557240dee939a320b826a1fd0d39bf6e629ecab47.
//
// Solidity: event EpochRebased(uint256 forkNumber, uint256 epochNumber, uint256 startBlockNumber, uint256 endBlockNumber, uint256 requestStart, uint256 requestEnd, bool epochIsEmpty, bool isRequest, bool userActivated)
func (_RootChainBase *RootChainBaseFilterer) ParseEpochRebased(log types.Log) (*RootChainBaseEpochRebased, error) {
	event := new(RootChainBaseEpochRebased)
	if err := _RootChainBase.contract.UnpackLog(event, "EpochRebased", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RootChainBaseForkedIterator is returned from FilterForked and is used to iterate over the raw logs and unpacked data for Forked events raised by the RootChainBase contract.
type RootChainBaseForkedIterator struct {
	Event *RootChainBaseForked // Event containing the contract specifics and raw log

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
func (it *RootChainBaseForkedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RootChainBaseForked)
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
		it.Event = new(RootChainBaseForked)
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
func (it *RootChainBaseForkedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RootChainBaseForkedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RootChainBaseForked represents a Forked event raised by the RootChainBase contract.
type RootChainBaseForked struct {
	NewFork           *big.Int
	EpochNumber       *big.Int
	ForkedBlockNumber *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterForked is a free log retrieval operation binding the contract event 0x0647d42ab02f6e0ae76959757dcb6aa6feac1d4ba6f077f1223fb4b1b429f06c.
//
// Solidity: event Forked(uint256 newFork, uint256 epochNumber, uint256 forkedBlockNumber)
func (_RootChainBase *RootChainBaseFilterer) FilterForked(opts *bind.FilterOpts) (*RootChainBaseForkedIterator, error) {

	logs, sub, err := _RootChainBase.contract.FilterLogs(opts, "Forked")
	if err != nil {
		return nil, err
	}
	return &RootChainBaseForkedIterator{contract: _RootChainBase.contract, event: "Forked", logs: logs, sub: sub}, nil
}

// WatchForked is a free log subscription operation binding the contract event 0x0647d42ab02f6e0ae76959757dcb6aa6feac1d4ba6f077f1223fb4b1b429f06c.
//
// Solidity: event Forked(uint256 newFork, uint256 epochNumber, uint256 forkedBlockNumber)
func (_RootChainBase *RootChainBaseFilterer) WatchForked(opts *bind.WatchOpts, sink chan<- *RootChainBaseForked) (event.Subscription, error) {

	logs, sub, err := _RootChainBase.contract.WatchLogs(opts, "Forked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RootChainBaseForked)
				if err := _RootChainBase.contract.UnpackLog(event, "Forked", log); err != nil {
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

// ParseForked is a log parse operation binding the contract event 0x0647d42ab02f6e0ae76959757dcb6aa6feac1d4ba6f077f1223fb4b1b429f06c.
//
// Solidity: event Forked(uint256 newFork, uint256 epochNumber, uint256 forkedBlockNumber)
func (_RootChainBase *RootChainBaseFilterer) ParseForked(log types.Log) (*RootChainBaseForked, error) {
	event := new(RootChainBaseForked)
	if err := _RootChainBase.contract.UnpackLog(event, "Forked", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RootChainBaseRequestAppliedIterator is returned from FilterRequestApplied and is used to iterate over the raw logs and unpacked data for RequestApplied events raised by the RootChainBase contract.
type RootChainBaseRequestAppliedIterator struct {
	Event *RootChainBaseRequestApplied // Event containing the contract specifics and raw log

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
func (it *RootChainBaseRequestAppliedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RootChainBaseRequestApplied)
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
		it.Event = new(RootChainBaseRequestApplied)
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
func (it *RootChainBaseRequestAppliedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RootChainBaseRequestAppliedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RootChainBaseRequestApplied represents a RequestApplied event raised by the RootChainBase contract.
type RootChainBaseRequestApplied struct {
	RequestId     *big.Int
	UserActivated bool
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRequestApplied is a free log retrieval operation binding the contract event 0x6940a01870e576ceb735867e13863646d517ce10e66c0133186a4ebdfe9388c2.
//
// Solidity: event RequestApplied(uint256 requestId, bool userActivated)
func (_RootChainBase *RootChainBaseFilterer) FilterRequestApplied(opts *bind.FilterOpts) (*RootChainBaseRequestAppliedIterator, error) {

	logs, sub, err := _RootChainBase.contract.FilterLogs(opts, "RequestApplied")
	if err != nil {
		return nil, err
	}
	return &RootChainBaseRequestAppliedIterator{contract: _RootChainBase.contract, event: "RequestApplied", logs: logs, sub: sub}, nil
}

// WatchRequestApplied is a free log subscription operation binding the contract event 0x6940a01870e576ceb735867e13863646d517ce10e66c0133186a4ebdfe9388c2.
//
// Solidity: event RequestApplied(uint256 requestId, bool userActivated)
func (_RootChainBase *RootChainBaseFilterer) WatchRequestApplied(opts *bind.WatchOpts, sink chan<- *RootChainBaseRequestApplied) (event.Subscription, error) {

	logs, sub, err := _RootChainBase.contract.WatchLogs(opts, "RequestApplied")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RootChainBaseRequestApplied)
				if err := _RootChainBase.contract.UnpackLog(event, "RequestApplied", log); err != nil {
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

// ParseRequestApplied is a log parse operation binding the contract event 0x6940a01870e576ceb735867e13863646d517ce10e66c0133186a4ebdfe9388c2.
//
// Solidity: event RequestApplied(uint256 requestId, bool userActivated)
func (_RootChainBase *RootChainBaseFilterer) ParseRequestApplied(log types.Log) (*RootChainBaseRequestApplied, error) {
	event := new(RootChainBaseRequestApplied)
	if err := _RootChainBase.contract.UnpackLog(event, "RequestApplied", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RootChainBaseRequestChallengedIterator is returned from FilterRequestChallenged and is used to iterate over the raw logs and unpacked data for RequestChallenged events raised by the RootChainBase contract.
type RootChainBaseRequestChallengedIterator struct {
	Event *RootChainBaseRequestChallenged // Event containing the contract specifics and raw log

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
func (it *RootChainBaseRequestChallengedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RootChainBaseRequestChallenged)
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
		it.Event = new(RootChainBaseRequestChallenged)
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
func (it *RootChainBaseRequestChallengedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RootChainBaseRequestChallengedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RootChainBaseRequestChallenged represents a RequestChallenged event raised by the RootChainBase contract.
type RootChainBaseRequestChallenged struct {
	RequestId     *big.Int
	UserActivated bool
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRequestChallenged is a free log retrieval operation binding the contract event 0xc8135db115644ed4ae193313c4c801235ef740d2a57a8d5e6fe26ab66635698a.
//
// Solidity: event RequestChallenged(uint256 requestId, bool userActivated)
func (_RootChainBase *RootChainBaseFilterer) FilterRequestChallenged(opts *bind.FilterOpts) (*RootChainBaseRequestChallengedIterator, error) {

	logs, sub, err := _RootChainBase.contract.FilterLogs(opts, "RequestChallenged")
	if err != nil {
		return nil, err
	}
	return &RootChainBaseRequestChallengedIterator{contract: _RootChainBase.contract, event: "RequestChallenged", logs: logs, sub: sub}, nil
}

// WatchRequestChallenged is a free log subscription operation binding the contract event 0xc8135db115644ed4ae193313c4c801235ef740d2a57a8d5e6fe26ab66635698a.
//
// Solidity: event RequestChallenged(uint256 requestId, bool userActivated)
func (_RootChainBase *RootChainBaseFilterer) WatchRequestChallenged(opts *bind.WatchOpts, sink chan<- *RootChainBaseRequestChallenged) (event.Subscription, error) {

	logs, sub, err := _RootChainBase.contract.WatchLogs(opts, "RequestChallenged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RootChainBaseRequestChallenged)
				if err := _RootChainBase.contract.UnpackLog(event, "RequestChallenged", log); err != nil {
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

// ParseRequestChallenged is a log parse operation binding the contract event 0xc8135db115644ed4ae193313c4c801235ef740d2a57a8d5e6fe26ab66635698a.
//
// Solidity: event RequestChallenged(uint256 requestId, bool userActivated)
func (_RootChainBase *RootChainBaseFilterer) ParseRequestChallenged(log types.Log) (*RootChainBaseRequestChallenged, error) {
	event := new(RootChainBaseRequestChallenged)
	if err := _RootChainBase.contract.UnpackLog(event, "RequestChallenged", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RootChainBaseRequestCreatedIterator is returned from FilterRequestCreated and is used to iterate over the raw logs and unpacked data for RequestCreated events raised by the RootChainBase contract.
type RootChainBaseRequestCreatedIterator struct {
	Event *RootChainBaseRequestCreated // Event containing the contract specifics and raw log

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
func (it *RootChainBaseRequestCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RootChainBaseRequestCreated)
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
		it.Event = new(RootChainBaseRequestCreated)
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
func (it *RootChainBaseRequestCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RootChainBaseRequestCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RootChainBaseRequestCreated represents a RequestCreated event raised by the RootChainBase contract.
type RootChainBaseRequestCreated struct {
	RequestId     *big.Int
	Requestor     common.Address
	To            common.Address
	WeiAmount     *big.Int
	TrieKey       [32]byte
	TrieValue     []byte
	IsExit        bool
	UserActivated bool
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRequestCreated is a free log retrieval operation binding the contract event 0x879922cf5fcada9ebaf8bd7424dc62877f4b220cae07fb6695cc1e8f94c52b4d.
//
// Solidity: event RequestCreated(uint256 requestId, address requestor, address to, uint256 weiAmount, bytes32 trieKey, bytes trieValue, bool isExit, bool userActivated)
func (_RootChainBase *RootChainBaseFilterer) FilterRequestCreated(opts *bind.FilterOpts) (*RootChainBaseRequestCreatedIterator, error) {

	logs, sub, err := _RootChainBase.contract.FilterLogs(opts, "RequestCreated")
	if err != nil {
		return nil, err
	}
	return &RootChainBaseRequestCreatedIterator{contract: _RootChainBase.contract, event: "RequestCreated", logs: logs, sub: sub}, nil
}

// WatchRequestCreated is a free log subscription operation binding the contract event 0x879922cf5fcada9ebaf8bd7424dc62877f4b220cae07fb6695cc1e8f94c52b4d.
//
// Solidity: event RequestCreated(uint256 requestId, address requestor, address to, uint256 weiAmount, bytes32 trieKey, bytes trieValue, bool isExit, bool userActivated)
func (_RootChainBase *RootChainBaseFilterer) WatchRequestCreated(opts *bind.WatchOpts, sink chan<- *RootChainBaseRequestCreated) (event.Subscription, error) {

	logs, sub, err := _RootChainBase.contract.WatchLogs(opts, "RequestCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RootChainBaseRequestCreated)
				if err := _RootChainBase.contract.UnpackLog(event, "RequestCreated", log); err != nil {
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

// ParseRequestCreated is a log parse operation binding the contract event 0x879922cf5fcada9ebaf8bd7424dc62877f4b220cae07fb6695cc1e8f94c52b4d.
//
// Solidity: event RequestCreated(uint256 requestId, address requestor, address to, uint256 weiAmount, bytes32 trieKey, bytes trieValue, bool isExit, bool userActivated)
func (_RootChainBase *RootChainBaseFilterer) ParseRequestCreated(log types.Log) (*RootChainBaseRequestCreated, error) {
	event := new(RootChainBaseRequestCreated)
	if err := _RootChainBase.contract.UnpackLog(event, "RequestCreated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RootChainBaseRequestFinalizedIterator is returned from FilterRequestFinalized and is used to iterate over the raw logs and unpacked data for RequestFinalized events raised by the RootChainBase contract.
type RootChainBaseRequestFinalizedIterator struct {
	Event *RootChainBaseRequestFinalized // Event containing the contract specifics and raw log

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
func (it *RootChainBaseRequestFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RootChainBaseRequestFinalized)
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
		it.Event = new(RootChainBaseRequestFinalized)
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
func (it *RootChainBaseRequestFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RootChainBaseRequestFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RootChainBaseRequestFinalized represents a RequestFinalized event raised by the RootChainBase contract.
type RootChainBaseRequestFinalized struct {
	RequestId     *big.Int
	UserActivated bool
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRequestFinalized is a free log retrieval operation binding the contract event 0x134017cf3262b18f892ee95dde3b0aec9a80cc70a7c96f09c64bd237aceb0473.
//
// Solidity: event RequestFinalized(uint256 requestId, bool userActivated)
func (_RootChainBase *RootChainBaseFilterer) FilterRequestFinalized(opts *bind.FilterOpts) (*RootChainBaseRequestFinalizedIterator, error) {

	logs, sub, err := _RootChainBase.contract.FilterLogs(opts, "RequestFinalized")
	if err != nil {
		return nil, err
	}
	return &RootChainBaseRequestFinalizedIterator{contract: _RootChainBase.contract, event: "RequestFinalized", logs: logs, sub: sub}, nil
}

// WatchRequestFinalized is a free log subscription operation binding the contract event 0x134017cf3262b18f892ee95dde3b0aec9a80cc70a7c96f09c64bd237aceb0473.
//
// Solidity: event RequestFinalized(uint256 requestId, bool userActivated)
func (_RootChainBase *RootChainBaseFilterer) WatchRequestFinalized(opts *bind.WatchOpts, sink chan<- *RootChainBaseRequestFinalized) (event.Subscription, error) {

	logs, sub, err := _RootChainBase.contract.WatchLogs(opts, "RequestFinalized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RootChainBaseRequestFinalized)
				if err := _RootChainBase.contract.UnpackLog(event, "RequestFinalized", log); err != nil {
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

// ParseRequestFinalized is a log parse operation binding the contract event 0x134017cf3262b18f892ee95dde3b0aec9a80cc70a7c96f09c64bd237aceb0473.
//
// Solidity: event RequestFinalized(uint256 requestId, bool userActivated)
func (_RootChainBase *RootChainBaseFilterer) ParseRequestFinalized(log types.Log) (*RootChainBaseRequestFinalized, error) {
	event := new(RootChainBaseRequestFinalized)
	if err := _RootChainBase.contract.UnpackLog(event, "RequestFinalized", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RootChainBaseRequestableContractMappedIterator is returned from FilterRequestableContractMapped and is used to iterate over the raw logs and unpacked data for RequestableContractMapped events raised by the RootChainBase contract.
type RootChainBaseRequestableContractMappedIterator struct {
	Event *RootChainBaseRequestableContractMapped // Event containing the contract specifics and raw log

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
func (it *RootChainBaseRequestableContractMappedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RootChainBaseRequestableContractMapped)
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
		it.Event = new(RootChainBaseRequestableContractMapped)
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
func (it *RootChainBaseRequestableContractMappedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RootChainBaseRequestableContractMappedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RootChainBaseRequestableContractMapped represents a RequestableContractMapped event raised by the RootChainBase contract.
type RootChainBaseRequestableContractMapped struct {
	ContractInRootchain  common.Address
	ContractInChildchain common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterRequestableContractMapped is a free log retrieval operation binding the contract event 0xc5ec2ed49686197edd2ed642c7e6096893cc81e6658cde2527030316037715d0.
//
// Solidity: event RequestableContractMapped(address contractInRootchain, address contractInChildchain)
func (_RootChainBase *RootChainBaseFilterer) FilterRequestableContractMapped(opts *bind.FilterOpts) (*RootChainBaseRequestableContractMappedIterator, error) {

	logs, sub, err := _RootChainBase.contract.FilterLogs(opts, "RequestableContractMapped")
	if err != nil {
		return nil, err
	}
	return &RootChainBaseRequestableContractMappedIterator{contract: _RootChainBase.contract, event: "RequestableContractMapped", logs: logs, sub: sub}, nil
}

// WatchRequestableContractMapped is a free log subscription operation binding the contract event 0xc5ec2ed49686197edd2ed642c7e6096893cc81e6658cde2527030316037715d0.
//
// Solidity: event RequestableContractMapped(address contractInRootchain, address contractInChildchain)
func (_RootChainBase *RootChainBaseFilterer) WatchRequestableContractMapped(opts *bind.WatchOpts, sink chan<- *RootChainBaseRequestableContractMapped) (event.Subscription, error) {

	logs, sub, err := _RootChainBase.contract.WatchLogs(opts, "RequestableContractMapped")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RootChainBaseRequestableContractMapped)
				if err := _RootChainBase.contract.UnpackLog(event, "RequestableContractMapped", log); err != nil {
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

// ParseRequestableContractMapped is a log parse operation binding the contract event 0xc5ec2ed49686197edd2ed642c7e6096893cc81e6658cde2527030316037715d0.
//
// Solidity: event RequestableContractMapped(address contractInRootchain, address contractInChildchain)
func (_RootChainBase *RootChainBaseFilterer) ParseRequestableContractMapped(log types.Log) (*RootChainBaseRequestableContractMapped, error) {
	event := new(RootChainBaseRequestableContractMapped)
	if err := _RootChainBase.contract.UnpackLog(event, "RequestableContractMapped", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RootChainBaseSessionTimeoutIterator is returned from FilterSessionTimeout and is used to iterate over the raw logs and unpacked data for SessionTimeout events raised by the RootChainBase contract.
type RootChainBaseSessionTimeoutIterator struct {
	Event *RootChainBaseSessionTimeout // Event containing the contract specifics and raw log

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
func (it *RootChainBaseSessionTimeoutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RootChainBaseSessionTimeout)
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
		it.Event = new(RootChainBaseSessionTimeout)
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
func (it *RootChainBaseSessionTimeoutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RootChainBaseSessionTimeoutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RootChainBaseSessionTimeout represents a SessionTimeout event raised by the RootChainBase contract.
type RootChainBaseSessionTimeout struct {
	UserActivated bool
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSessionTimeout is a free log retrieval operation binding the contract event 0x2122ec719581cd177f225f59a1ee005831211196831f8f1ccffa817d2e7bd108.
//
// Solidity: event SessionTimeout(bool userActivated)
func (_RootChainBase *RootChainBaseFilterer) FilterSessionTimeout(opts *bind.FilterOpts) (*RootChainBaseSessionTimeoutIterator, error) {

	logs, sub, err := _RootChainBase.contract.FilterLogs(opts, "SessionTimeout")
	if err != nil {
		return nil, err
	}
	return &RootChainBaseSessionTimeoutIterator{contract: _RootChainBase.contract, event: "SessionTimeout", logs: logs, sub: sub}, nil
}

// WatchSessionTimeout is a free log subscription operation binding the contract event 0x2122ec719581cd177f225f59a1ee005831211196831f8f1ccffa817d2e7bd108.
//
// Solidity: event SessionTimeout(bool userActivated)
func (_RootChainBase *RootChainBaseFilterer) WatchSessionTimeout(opts *bind.WatchOpts, sink chan<- *RootChainBaseSessionTimeout) (event.Subscription, error) {

	logs, sub, err := _RootChainBase.contract.WatchLogs(opts, "SessionTimeout")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RootChainBaseSessionTimeout)
				if err := _RootChainBase.contract.UnpackLog(event, "SessionTimeout", log); err != nil {
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

// ParseSessionTimeout is a log parse operation binding the contract event 0x2122ec719581cd177f225f59a1ee005831211196831f8f1ccffa817d2e7bd108.
//
// Solidity: event SessionTimeout(bool userActivated)
func (_RootChainBase *RootChainBaseFilterer) ParseSessionTimeout(log types.Log) (*RootChainBaseSessionTimeout, error) {
	event := new(RootChainBaseSessionTimeout)
	if err := _RootChainBase.contract.UnpackLog(event, "SessionTimeout", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RootChainEventABI is the input ABI used to generate the binding from.
const RootChainEventABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"userActivated\",\"type\":\"bool\"}],\"name\":\"SessionTimeout\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newFork\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"epochNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"forkedBlockNumber\",\"type\":\"uint256\"}],\"name\":\"Forked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"forkNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"epochNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"startBlockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"endBlockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"requestStart\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"requestEnd\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"epochIsEmpty\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"isRequest\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"userActivated\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"rebase\",\"type\":\"bool\"}],\"name\":\"EpochPrepared\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"forkNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"epochNumber\",\"type\":\"uint256\"}],\"name\":\"EpochFilling\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"forkNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"epochNumber\",\"type\":\"uint256\"}],\"name\":\"EpochFilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"forkNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"epochNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"startBlockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"endBlockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"requestStart\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"requestEnd\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"epochIsEmpty\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"isRequest\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"userActivated\",\"type\":\"bool\"}],\"name\":\"EpochRebased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"fork\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"epochNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"isRequest\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"userActivated\",\"type\":\"bool\"}],\"name\":\"BlockSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"requestor\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"weiAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"trieKey\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"trieValue\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"isExit\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"userActivated\",\"type\":\"bool\"}],\"name\":\"RequestCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"requestor\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"trieKey\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"trieValue\",\"type\":\"bytes32\"}],\"name\":\"ERUCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"forkNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"BlockFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"forkNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"epochNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"startBlockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"endBlockNumber\",\"type\":\"uint256\"}],\"name\":\"EpochFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"userActivated\",\"type\":\"bool\"}],\"name\":\"RequestFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"userActivated\",\"type\":\"bool\"}],\"name\":\"RequestApplied\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"userActivated\",\"type\":\"bool\"}],\"name\":\"RequestChallenged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"contractInRootchain\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"contractInChildchain\",\"type\":\"address\"}],\"name\":\"RequestableContractMapped\",\"type\":\"event\"}]"

// RootChainEventBin is the compiled bytecode used for deploying new contracts.
var RootChainEventBin = "0x6080604052348015600f57600080fd5b50603e80601d6000396000f3fe6080604052600080fdfea265627a7a723058206ee580968ba091e6149f220d7bda19eec17b48b5deecd8f525ce1bf39929c05e64736f6c63430005090032"

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
// Solidity: event BlockFinalized(uint256 forkNumber, uint256 blockNumber)
func (_RootChainEvent *RootChainEventFilterer) FilterBlockFinalized(opts *bind.FilterOpts) (*RootChainEventBlockFinalizedIterator, error) {

	logs, sub, err := _RootChainEvent.contract.FilterLogs(opts, "BlockFinalized")
	if err != nil {
		return nil, err
	}
	return &RootChainEventBlockFinalizedIterator{contract: _RootChainEvent.contract, event: "BlockFinalized", logs: logs, sub: sub}, nil
}

// WatchBlockFinalized is a free log subscription operation binding the contract event 0xfb96205e4b3633fd57aa805b26b51ecf528714a10241a4af015929dce86768d9.
//
// Solidity: event BlockFinalized(uint256 forkNumber, uint256 blockNumber)
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

// ParseBlockFinalized is a log parse operation binding the contract event 0xfb96205e4b3633fd57aa805b26b51ecf528714a10241a4af015929dce86768d9.
//
// Solidity: event BlockFinalized(uint256 forkNumber, uint256 blockNumber)
func (_RootChainEvent *RootChainEventFilterer) ParseBlockFinalized(log types.Log) (*RootChainEventBlockFinalized, error) {
	event := new(RootChainEventBlockFinalized)
	if err := _RootChainEvent.contract.UnpackLog(event, "BlockFinalized", log); err != nil {
		return nil, err
	}
	return event, nil
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
// Solidity: event BlockSubmitted(uint256 fork, uint256 epochNumber, uint256 blockNumber, bool isRequest, bool userActivated)
func (_RootChainEvent *RootChainEventFilterer) FilterBlockSubmitted(opts *bind.FilterOpts) (*RootChainEventBlockSubmittedIterator, error) {

	logs, sub, err := _RootChainEvent.contract.FilterLogs(opts, "BlockSubmitted")
	if err != nil {
		return nil, err
	}
	return &RootChainEventBlockSubmittedIterator{contract: _RootChainEvent.contract, event: "BlockSubmitted", logs: logs, sub: sub}, nil
}

// WatchBlockSubmitted is a free log subscription operation binding the contract event 0x3d4a04291c66b06f39a4ecb817875b12b5485a05ec563133a56a905305c48e55.
//
// Solidity: event BlockSubmitted(uint256 fork, uint256 epochNumber, uint256 blockNumber, bool isRequest, bool userActivated)
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

// ParseBlockSubmitted is a log parse operation binding the contract event 0x3d4a04291c66b06f39a4ecb817875b12b5485a05ec563133a56a905305c48e55.
//
// Solidity: event BlockSubmitted(uint256 fork, uint256 epochNumber, uint256 blockNumber, bool isRequest, bool userActivated)
func (_RootChainEvent *RootChainEventFilterer) ParseBlockSubmitted(log types.Log) (*RootChainEventBlockSubmitted, error) {
	event := new(RootChainEventBlockSubmitted)
	if err := _RootChainEvent.contract.UnpackLog(event, "BlockSubmitted", log); err != nil {
		return nil, err
	}
	return event, nil
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
	TrieKey   []byte
	TrieValue [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterERUCreated is a free log retrieval operation binding the contract event 0xfcbdc2083dadd644b854d91b49aef8db06b8f5a3d5c1192de38ca0ba271d5a0d.
//
// Solidity: event ERUCreated(uint256 requestId, address requestor, address to, bytes trieKey, bytes32 trieValue)
func (_RootChainEvent *RootChainEventFilterer) FilterERUCreated(opts *bind.FilterOpts) (*RootChainEventERUCreatedIterator, error) {

	logs, sub, err := _RootChainEvent.contract.FilterLogs(opts, "ERUCreated")
	if err != nil {
		return nil, err
	}
	return &RootChainEventERUCreatedIterator{contract: _RootChainEvent.contract, event: "ERUCreated", logs: logs, sub: sub}, nil
}

// WatchERUCreated is a free log subscription operation binding the contract event 0xfcbdc2083dadd644b854d91b49aef8db06b8f5a3d5c1192de38ca0ba271d5a0d.
//
// Solidity: event ERUCreated(uint256 requestId, address requestor, address to, bytes trieKey, bytes32 trieValue)
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

// ParseERUCreated is a log parse operation binding the contract event 0xfcbdc2083dadd644b854d91b49aef8db06b8f5a3d5c1192de38ca0ba271d5a0d.
//
// Solidity: event ERUCreated(uint256 requestId, address requestor, address to, bytes trieKey, bytes32 trieValue)
func (_RootChainEvent *RootChainEventFilterer) ParseERUCreated(log types.Log) (*RootChainEventERUCreated, error) {
	event := new(RootChainEventERUCreated)
	if err := _RootChainEvent.contract.UnpackLog(event, "ERUCreated", log); err != nil {
		return nil, err
	}
	return event, nil
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
// Solidity: event EpochFilled(uint256 forkNumber, uint256 epochNumber)
func (_RootChainEvent *RootChainEventFilterer) FilterEpochFilled(opts *bind.FilterOpts) (*RootChainEventEpochFilledIterator, error) {

	logs, sub, err := _RootChainEvent.contract.FilterLogs(opts, "EpochFilled")
	if err != nil {
		return nil, err
	}
	return &RootChainEventEpochFilledIterator{contract: _RootChainEvent.contract, event: "EpochFilled", logs: logs, sub: sub}, nil
}

// WatchEpochFilled is a free log subscription operation binding the contract event 0x2fdeb407bf5c2b621f04b5c784822dae806c45b49a68aba413cc270128c96816.
//
// Solidity: event EpochFilled(uint256 forkNumber, uint256 epochNumber)
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

// ParseEpochFilled is a log parse operation binding the contract event 0x2fdeb407bf5c2b621f04b5c784822dae806c45b49a68aba413cc270128c96816.
//
// Solidity: event EpochFilled(uint256 forkNumber, uint256 epochNumber)
func (_RootChainEvent *RootChainEventFilterer) ParseEpochFilled(log types.Log) (*RootChainEventEpochFilled, error) {
	event := new(RootChainEventEpochFilled)
	if err := _RootChainEvent.contract.UnpackLog(event, "EpochFilled", log); err != nil {
		return nil, err
	}
	return event, nil
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
// Solidity: event EpochFilling(uint256 forkNumber, uint256 epochNumber)
func (_RootChainEvent *RootChainEventFilterer) FilterEpochFilling(opts *bind.FilterOpts) (*RootChainEventEpochFillingIterator, error) {

	logs, sub, err := _RootChainEvent.contract.FilterLogs(opts, "EpochFilling")
	if err != nil {
		return nil, err
	}
	return &RootChainEventEpochFillingIterator{contract: _RootChainEvent.contract, event: "EpochFilling", logs: logs, sub: sub}, nil
}

// WatchEpochFilling is a free log subscription operation binding the contract event 0x27b09f0953d27bbff306fe25b2987ac5a813248ac30cb2bbd5daf95e7b0e6dc0.
//
// Solidity: event EpochFilling(uint256 forkNumber, uint256 epochNumber)
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

// ParseEpochFilling is a log parse operation binding the contract event 0x27b09f0953d27bbff306fe25b2987ac5a813248ac30cb2bbd5daf95e7b0e6dc0.
//
// Solidity: event EpochFilling(uint256 forkNumber, uint256 epochNumber)
func (_RootChainEvent *RootChainEventFilterer) ParseEpochFilling(log types.Log) (*RootChainEventEpochFilling, error) {
	event := new(RootChainEventEpochFilling)
	if err := _RootChainEvent.contract.UnpackLog(event, "EpochFilling", log); err != nil {
		return nil, err
	}
	return event, nil
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
// Solidity: event EpochFinalized(uint256 forkNumber, uint256 epochNumber, uint256 startBlockNumber, uint256 endBlockNumber)
func (_RootChainEvent *RootChainEventFilterer) FilterEpochFinalized(opts *bind.FilterOpts) (*RootChainEventEpochFinalizedIterator, error) {

	logs, sub, err := _RootChainEvent.contract.FilterLogs(opts, "EpochFinalized")
	if err != nil {
		return nil, err
	}
	return &RootChainEventEpochFinalizedIterator{contract: _RootChainEvent.contract, event: "EpochFinalized", logs: logs, sub: sub}, nil
}

// WatchEpochFinalized is a free log subscription operation binding the contract event 0x70801d4d63b3da6c19ba7349911f45bed5a99ccdfb51b8138c105872529bebd5.
//
// Solidity: event EpochFinalized(uint256 forkNumber, uint256 epochNumber, uint256 startBlockNumber, uint256 endBlockNumber)
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

// ParseEpochFinalized is a log parse operation binding the contract event 0x70801d4d63b3da6c19ba7349911f45bed5a99ccdfb51b8138c105872529bebd5.
//
// Solidity: event EpochFinalized(uint256 forkNumber, uint256 epochNumber, uint256 startBlockNumber, uint256 endBlockNumber)
func (_RootChainEvent *RootChainEventFilterer) ParseEpochFinalized(log types.Log) (*RootChainEventEpochFinalized, error) {
	event := new(RootChainEventEpochFinalized)
	if err := _RootChainEvent.contract.UnpackLog(event, "EpochFinalized", log); err != nil {
		return nil, err
	}
	return event, nil
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
// Solidity: event EpochPrepared(uint256 forkNumber, uint256 epochNumber, uint256 startBlockNumber, uint256 endBlockNumber, uint256 requestStart, uint256 requestEnd, bool epochIsEmpty, bool isRequest, bool userActivated, bool rebase)
func (_RootChainEvent *RootChainEventFilterer) FilterEpochPrepared(opts *bind.FilterOpts) (*RootChainEventEpochPreparedIterator, error) {

	logs, sub, err := _RootChainEvent.contract.FilterLogs(opts, "EpochPrepared")
	if err != nil {
		return nil, err
	}
	return &RootChainEventEpochPreparedIterator{contract: _RootChainEvent.contract, event: "EpochPrepared", logs: logs, sub: sub}, nil
}

// WatchEpochPrepared is a free log subscription operation binding the contract event 0x1a69c0760aa329b76f72579129869013ebd3d41594db019c0e997b939fcb32e3.
//
// Solidity: event EpochPrepared(uint256 forkNumber, uint256 epochNumber, uint256 startBlockNumber, uint256 endBlockNumber, uint256 requestStart, uint256 requestEnd, bool epochIsEmpty, bool isRequest, bool userActivated, bool rebase)
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

// ParseEpochPrepared is a log parse operation binding the contract event 0x1a69c0760aa329b76f72579129869013ebd3d41594db019c0e997b939fcb32e3.
//
// Solidity: event EpochPrepared(uint256 forkNumber, uint256 epochNumber, uint256 startBlockNumber, uint256 endBlockNumber, uint256 requestStart, uint256 requestEnd, bool epochIsEmpty, bool isRequest, bool userActivated, bool rebase)
func (_RootChainEvent *RootChainEventFilterer) ParseEpochPrepared(log types.Log) (*RootChainEventEpochPrepared, error) {
	event := new(RootChainEventEpochPrepared)
	if err := _RootChainEvent.contract.UnpackLog(event, "EpochPrepared", log); err != nil {
		return nil, err
	}
	return event, nil
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
// Solidity: event EpochRebased(uint256 forkNumber, uint256 epochNumber, uint256 startBlockNumber, uint256 endBlockNumber, uint256 requestStart, uint256 requestEnd, bool epochIsEmpty, bool isRequest, bool userActivated)
func (_RootChainEvent *RootChainEventFilterer) FilterEpochRebased(opts *bind.FilterOpts) (*RootChainEventEpochRebasedIterator, error) {

	logs, sub, err := _RootChainEvent.contract.FilterLogs(opts, "EpochRebased")
	if err != nil {
		return nil, err
	}
	return &RootChainEventEpochRebasedIterator{contract: _RootChainEvent.contract, event: "EpochRebased", logs: logs, sub: sub}, nil
}

// WatchEpochRebased is a free log subscription operation binding the contract event 0x030c1c69405c93021f28f57557240dee939a320b826a1fd0d39bf6e629ecab47.
//
// Solidity: event EpochRebased(uint256 forkNumber, uint256 epochNumber, uint256 startBlockNumber, uint256 endBlockNumber, uint256 requestStart, uint256 requestEnd, bool epochIsEmpty, bool isRequest, bool userActivated)
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

// ParseEpochRebased is a log parse operation binding the contract event 0x030c1c69405c93021f28f57557240dee939a320b826a1fd0d39bf6e629ecab47.
//
// Solidity: event EpochRebased(uint256 forkNumber, uint256 epochNumber, uint256 startBlockNumber, uint256 endBlockNumber, uint256 requestStart, uint256 requestEnd, bool epochIsEmpty, bool isRequest, bool userActivated)
func (_RootChainEvent *RootChainEventFilterer) ParseEpochRebased(log types.Log) (*RootChainEventEpochRebased, error) {
	event := new(RootChainEventEpochRebased)
	if err := _RootChainEvent.contract.UnpackLog(event, "EpochRebased", log); err != nil {
		return nil, err
	}
	return event, nil
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
// Solidity: event Forked(uint256 newFork, uint256 epochNumber, uint256 forkedBlockNumber)
func (_RootChainEvent *RootChainEventFilterer) FilterForked(opts *bind.FilterOpts) (*RootChainEventForkedIterator, error) {

	logs, sub, err := _RootChainEvent.contract.FilterLogs(opts, "Forked")
	if err != nil {
		return nil, err
	}
	return &RootChainEventForkedIterator{contract: _RootChainEvent.contract, event: "Forked", logs: logs, sub: sub}, nil
}

// WatchForked is a free log subscription operation binding the contract event 0x0647d42ab02f6e0ae76959757dcb6aa6feac1d4ba6f077f1223fb4b1b429f06c.
//
// Solidity: event Forked(uint256 newFork, uint256 epochNumber, uint256 forkedBlockNumber)
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

// ParseForked is a log parse operation binding the contract event 0x0647d42ab02f6e0ae76959757dcb6aa6feac1d4ba6f077f1223fb4b1b429f06c.
//
// Solidity: event Forked(uint256 newFork, uint256 epochNumber, uint256 forkedBlockNumber)
func (_RootChainEvent *RootChainEventFilterer) ParseForked(log types.Log) (*RootChainEventForked, error) {
	event := new(RootChainEventForked)
	if err := _RootChainEvent.contract.UnpackLog(event, "Forked", log); err != nil {
		return nil, err
	}
	return event, nil
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
// Solidity: event RequestApplied(uint256 requestId, bool userActivated)
func (_RootChainEvent *RootChainEventFilterer) FilterRequestApplied(opts *bind.FilterOpts) (*RootChainEventRequestAppliedIterator, error) {

	logs, sub, err := _RootChainEvent.contract.FilterLogs(opts, "RequestApplied")
	if err != nil {
		return nil, err
	}
	return &RootChainEventRequestAppliedIterator{contract: _RootChainEvent.contract, event: "RequestApplied", logs: logs, sub: sub}, nil
}

// WatchRequestApplied is a free log subscription operation binding the contract event 0x6940a01870e576ceb735867e13863646d517ce10e66c0133186a4ebdfe9388c2.
//
// Solidity: event RequestApplied(uint256 requestId, bool userActivated)
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

// ParseRequestApplied is a log parse operation binding the contract event 0x6940a01870e576ceb735867e13863646d517ce10e66c0133186a4ebdfe9388c2.
//
// Solidity: event RequestApplied(uint256 requestId, bool userActivated)
func (_RootChainEvent *RootChainEventFilterer) ParseRequestApplied(log types.Log) (*RootChainEventRequestApplied, error) {
	event := new(RootChainEventRequestApplied)
	if err := _RootChainEvent.contract.UnpackLog(event, "RequestApplied", log); err != nil {
		return nil, err
	}
	return event, nil
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
// Solidity: event RequestChallenged(uint256 requestId, bool userActivated)
func (_RootChainEvent *RootChainEventFilterer) FilterRequestChallenged(opts *bind.FilterOpts) (*RootChainEventRequestChallengedIterator, error) {

	logs, sub, err := _RootChainEvent.contract.FilterLogs(opts, "RequestChallenged")
	if err != nil {
		return nil, err
	}
	return &RootChainEventRequestChallengedIterator{contract: _RootChainEvent.contract, event: "RequestChallenged", logs: logs, sub: sub}, nil
}

// WatchRequestChallenged is a free log subscription operation binding the contract event 0xc8135db115644ed4ae193313c4c801235ef740d2a57a8d5e6fe26ab66635698a.
//
// Solidity: event RequestChallenged(uint256 requestId, bool userActivated)
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

// ParseRequestChallenged is a log parse operation binding the contract event 0xc8135db115644ed4ae193313c4c801235ef740d2a57a8d5e6fe26ab66635698a.
//
// Solidity: event RequestChallenged(uint256 requestId, bool userActivated)
func (_RootChainEvent *RootChainEventFilterer) ParseRequestChallenged(log types.Log) (*RootChainEventRequestChallenged, error) {
	event := new(RootChainEventRequestChallenged)
	if err := _RootChainEvent.contract.UnpackLog(event, "RequestChallenged", log); err != nil {
		return nil, err
	}
	return event, nil
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
	TrieValue     []byte
	IsExit        bool
	UserActivated bool
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRequestCreated is a free log retrieval operation binding the contract event 0x879922cf5fcada9ebaf8bd7424dc62877f4b220cae07fb6695cc1e8f94c52b4d.
//
// Solidity: event RequestCreated(uint256 requestId, address requestor, address to, uint256 weiAmount, bytes32 trieKey, bytes trieValue, bool isExit, bool userActivated)
func (_RootChainEvent *RootChainEventFilterer) FilterRequestCreated(opts *bind.FilterOpts) (*RootChainEventRequestCreatedIterator, error) {

	logs, sub, err := _RootChainEvent.contract.FilterLogs(opts, "RequestCreated")
	if err != nil {
		return nil, err
	}
	return &RootChainEventRequestCreatedIterator{contract: _RootChainEvent.contract, event: "RequestCreated", logs: logs, sub: sub}, nil
}

// WatchRequestCreated is a free log subscription operation binding the contract event 0x879922cf5fcada9ebaf8bd7424dc62877f4b220cae07fb6695cc1e8f94c52b4d.
//
// Solidity: event RequestCreated(uint256 requestId, address requestor, address to, uint256 weiAmount, bytes32 trieKey, bytes trieValue, bool isExit, bool userActivated)
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

// ParseRequestCreated is a log parse operation binding the contract event 0x879922cf5fcada9ebaf8bd7424dc62877f4b220cae07fb6695cc1e8f94c52b4d.
//
// Solidity: event RequestCreated(uint256 requestId, address requestor, address to, uint256 weiAmount, bytes32 trieKey, bytes trieValue, bool isExit, bool userActivated)
func (_RootChainEvent *RootChainEventFilterer) ParseRequestCreated(log types.Log) (*RootChainEventRequestCreated, error) {
	event := new(RootChainEventRequestCreated)
	if err := _RootChainEvent.contract.UnpackLog(event, "RequestCreated", log); err != nil {
		return nil, err
	}
	return event, nil
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
// Solidity: event RequestFinalized(uint256 requestId, bool userActivated)
func (_RootChainEvent *RootChainEventFilterer) FilterRequestFinalized(opts *bind.FilterOpts) (*RootChainEventRequestFinalizedIterator, error) {

	logs, sub, err := _RootChainEvent.contract.FilterLogs(opts, "RequestFinalized")
	if err != nil {
		return nil, err
	}
	return &RootChainEventRequestFinalizedIterator{contract: _RootChainEvent.contract, event: "RequestFinalized", logs: logs, sub: sub}, nil
}

// WatchRequestFinalized is a free log subscription operation binding the contract event 0x134017cf3262b18f892ee95dde3b0aec9a80cc70a7c96f09c64bd237aceb0473.
//
// Solidity: event RequestFinalized(uint256 requestId, bool userActivated)
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

// ParseRequestFinalized is a log parse operation binding the contract event 0x134017cf3262b18f892ee95dde3b0aec9a80cc70a7c96f09c64bd237aceb0473.
//
// Solidity: event RequestFinalized(uint256 requestId, bool userActivated)
func (_RootChainEvent *RootChainEventFilterer) ParseRequestFinalized(log types.Log) (*RootChainEventRequestFinalized, error) {
	event := new(RootChainEventRequestFinalized)
	if err := _RootChainEvent.contract.UnpackLog(event, "RequestFinalized", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RootChainEventRequestableContractMappedIterator is returned from FilterRequestableContractMapped and is used to iterate over the raw logs and unpacked data for RequestableContractMapped events raised by the RootChainEvent contract.
type RootChainEventRequestableContractMappedIterator struct {
	Event *RootChainEventRequestableContractMapped // Event containing the contract specifics and raw log

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
func (it *RootChainEventRequestableContractMappedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RootChainEventRequestableContractMapped)
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
		it.Event = new(RootChainEventRequestableContractMapped)
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
func (it *RootChainEventRequestableContractMappedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RootChainEventRequestableContractMappedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RootChainEventRequestableContractMapped represents a RequestableContractMapped event raised by the RootChainEvent contract.
type RootChainEventRequestableContractMapped struct {
	ContractInRootchain  common.Address
	ContractInChildchain common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterRequestableContractMapped is a free log retrieval operation binding the contract event 0xc5ec2ed49686197edd2ed642c7e6096893cc81e6658cde2527030316037715d0.
//
// Solidity: event RequestableContractMapped(address contractInRootchain, address contractInChildchain)
func (_RootChainEvent *RootChainEventFilterer) FilterRequestableContractMapped(opts *bind.FilterOpts) (*RootChainEventRequestableContractMappedIterator, error) {

	logs, sub, err := _RootChainEvent.contract.FilterLogs(opts, "RequestableContractMapped")
	if err != nil {
		return nil, err
	}
	return &RootChainEventRequestableContractMappedIterator{contract: _RootChainEvent.contract, event: "RequestableContractMapped", logs: logs, sub: sub}, nil
}

// WatchRequestableContractMapped is a free log subscription operation binding the contract event 0xc5ec2ed49686197edd2ed642c7e6096893cc81e6658cde2527030316037715d0.
//
// Solidity: event RequestableContractMapped(address contractInRootchain, address contractInChildchain)
func (_RootChainEvent *RootChainEventFilterer) WatchRequestableContractMapped(opts *bind.WatchOpts, sink chan<- *RootChainEventRequestableContractMapped) (event.Subscription, error) {

	logs, sub, err := _RootChainEvent.contract.WatchLogs(opts, "RequestableContractMapped")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RootChainEventRequestableContractMapped)
				if err := _RootChainEvent.contract.UnpackLog(event, "RequestableContractMapped", log); err != nil {
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

// ParseRequestableContractMapped is a log parse operation binding the contract event 0xc5ec2ed49686197edd2ed642c7e6096893cc81e6658cde2527030316037715d0.
//
// Solidity: event RequestableContractMapped(address contractInRootchain, address contractInChildchain)
func (_RootChainEvent *RootChainEventFilterer) ParseRequestableContractMapped(log types.Log) (*RootChainEventRequestableContractMapped, error) {
	event := new(RootChainEventRequestableContractMapped)
	if err := _RootChainEvent.contract.UnpackLog(event, "RequestableContractMapped", log); err != nil {
		return nil, err
	}
	return event, nil
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
// Solidity: event SessionTimeout(bool userActivated)
func (_RootChainEvent *RootChainEventFilterer) FilterSessionTimeout(opts *bind.FilterOpts) (*RootChainEventSessionTimeoutIterator, error) {

	logs, sub, err := _RootChainEvent.contract.FilterLogs(opts, "SessionTimeout")
	if err != nil {
		return nil, err
	}
	return &RootChainEventSessionTimeoutIterator{contract: _RootChainEvent.contract, event: "SessionTimeout", logs: logs, sub: sub}, nil
}

// WatchSessionTimeout is a free log subscription operation binding the contract event 0x2122ec719581cd177f225f59a1ee005831211196831f8f1ccffa817d2e7bd108.
//
// Solidity: event SessionTimeout(bool userActivated)
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

// ParseSessionTimeout is a log parse operation binding the contract event 0x2122ec719581cd177f225f59a1ee005831211196831f8f1ccffa817d2e7bd108.
//
// Solidity: event SessionTimeout(bool userActivated)
func (_RootChainEvent *RootChainEventFilterer) ParseSessionTimeout(log types.Log) (*RootChainEventSessionTimeout, error) {
	event := new(RootChainEventSessionTimeout)
	if err := _RootChainEvent.contract.UnpackLog(event, "SessionTimeout", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RootChainStorageABI is the input ABI used to generate the binding from.
const RootChainStorageABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"COST_URB_PREPARE\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CP_COMPUTATION\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastAppliedForkNumber\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"currentFork\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"COST_URB\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"numEnterForORB\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"EROIdToFinalize\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"forks\",\"outputs\":[{\"name\":\"forkedBlock\",\"type\":\"uint64\"},{\"name\":\"firstEpoch\",\"type\":\"uint64\"},{\"name\":\"lastEpoch\",\"type\":\"uint64\"},{\"name\":\"firstBlock\",\"type\":\"uint64\"},{\"name\":\"lastBlock\",\"type\":\"uint64\"},{\"name\":\"lastFinalizedEpoch\",\"type\":\"uint64\"},{\"name\":\"lastFinalizedBlock\",\"type\":\"uint64\"},{\"name\":\"timestamp\",\"type\":\"uint64\"},{\"name\":\"firstEnterEpoch\",\"type\":\"uint64\"},{\"name\":\"lastEnterEpoch\",\"type\":\"uint64\"},{\"name\":\"nextBlockToRebase\",\"type\":\"uint64\"},{\"name\":\"rebased\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"operator\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"firstFilledORENumber\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"development\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CP_EXIT\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"COST_ERU\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"REQUEST_GAS\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"COST_NRB\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"NRELength\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CP_WITHHOLDING\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"COST_ORB\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"EROs\",\"outputs\":[{\"name\":\"timestamp\",\"type\":\"uint64\"},{\"name\":\"isExit\",\"type\":\"bool\"},{\"name\":\"isTransfer\",\"type\":\"bool\"},{\"name\":\"finalized\",\"type\":\"bool\"},{\"name\":\"challenged\",\"type\":\"bool\"},{\"name\":\"value\",\"type\":\"uint128\"},{\"name\":\"requestor\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"trieKey\",\"type\":\"bytes32\"},{\"name\":\"hash\",\"type\":\"bytes32\"},{\"name\":\"trieValue\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"lastNonEmptyRequestEpoch\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"etherToken\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"URBs\",\"outputs\":[{\"name\":\"submitted\",\"type\":\"bool\"},{\"name\":\"numEnter\",\"type\":\"uint64\"},{\"name\":\"epochNumber\",\"type\":\"uint64\"},{\"name\":\"requestStart\",\"type\":\"uint64\"},{\"name\":\"requestEnd\",\"type\":\"uint64\"},{\"name\":\"trie\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PREPARE_TIMEOUT\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ERUIdToFinalize\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastAppliedEpochNumber\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"firstNonEmptyRequestEpoch\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"COST_ERO\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"requestableContracts\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"NULL_ADDRESS\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"submitHandler\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"epochHandler\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"ORBs\",\"outputs\":[{\"name\":\"submitted\",\"type\":\"bool\"},{\"name\":\"numEnter\",\"type\":\"uint64\"},{\"name\":\"epochNumber\",\"type\":\"uint64\"},{\"name\":\"requestStart\",\"type\":\"uint64\"},{\"name\":\"requestEnd\",\"type\":\"uint64\"},{\"name\":\"trie\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"ERUs\",\"outputs\":[{\"name\":\"timestamp\",\"type\":\"uint64\"},{\"name\":\"isExit\",\"type\":\"bool\"},{\"name\":\"isTransfer\",\"type\":\"bool\"},{\"name\":\"finalized\",\"type\":\"bool\"},{\"name\":\"challenged\",\"type\":\"bool\"},{\"name\":\"value\",\"type\":\"uint128\"},{\"name\":\"requestor\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"trieKey\",\"type\":\"bytes32\"},{\"name\":\"hash\",\"type\":\"bytes32\"},{\"name\":\"trieValue\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastAppliedBlockNumber\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// RootChainStorageFuncSigs maps the 4-byte function signature to its string representation.
var RootChainStorageFuncSigs = map[string]string{
	"d691acd8": "COST_ERO()",
	"8b5172d0": "COST_ERU()",
	"94be3aa5": "COST_NRB()",
	"b2ae9ba8": "COST_ORB()",
	"192adc5b": "COST_URB()",
	"033cfbed": "COST_URB_PREPARE()",
	"08c4fff0": "CP_COMPUTATION()",
	"8155717d": "CP_EXIT()",
	"b17fa6e9": "CP_WITHHOLDING()",
	"2dc6bb7b": "EROIdToFinalize()",
	"b443f3cc": "EROs(uint256)",
	"c54626cc": "ERUIdToFinalize()",
	"f4f31de4": "ERUs(uint256)",
	"ab96da2d": "NRELength()",
	"de0ce17d": "NULL_ADDRESS()",
	"ea7f22a8": "ORBs(uint256)",
	"c2bc88fa": "PREPARE_TIMEOUT()",
	"8eb288ca": "REQUEST_GAS()",
	"c0e86064": "URBs(uint256)",
	"183d2d1c": "currentFork()",
	"7b929c27": "development()",
	"e7b88b80": "epochHandler()",
	"b8066bcb": "etherToken()",
	"72ecb9a8": "firstFilledORENumber(uint256)",
	"ca6f6380": "firstNonEmptyRequestEpoch(uint256)",
	"4ba3a126": "forks(uint256)",
	"fb788a27": "lastAppliedBlockNumber()",
	"c8ad329f": "lastAppliedEpochNumber()",
	"164bc2ae": "lastAppliedForkNumber()",
	"b6715647": "lastNonEmptyRequestEpoch(uint256)",
	"23691566": "numEnterForORB()",
	"570ca735": "operator()",
	"da0185f8": "requestableContracts(address)",
	"e259faf7": "submitHandler()",
}

// RootChainStorageBin is the compiled bytecode used for deploying new contracts.
var RootChainStorageBin = "0x608060405234801561001057600080fd5b5061091b806100206000396000f3fe608060405234801561001057600080fd5b50600436106102065760003560e01c8063b2ae9ba81161011a578063ca6f6380116100ad578063e259faf71161007c578063e259faf71461056f578063e7b88b8014610577578063ea7f22a81461057f578063f4f31de41461059c578063fb788a27146105b957610206565b8063ca6f638014610524578063d691acd81461020b578063da0185f814610541578063de0ce17d1461056757610206565b8063c0e86064116100e9578063c0e86064146104a3578063c2bc88fa1461050c578063c54626cc14610514578063c8ad329f1461051c57610206565b8063b2ae9ba81461020b578063b443f3cc1461036e578063b67156471461047e578063b8066bcb1461049b57610206565b8063570ca7351161019d5780638b5172d01161016c5780638b5172d01461034e5780638eb288ca1461035657806394be3aa51461020b578063ab96da2d1461035e578063b17fa6e91461036657610206565b8063570ca735146102e957806372ecb9a81461030d5780637b929c271461032a5780638155717d1461034657610206565b8063192adc5b116101d9578063192adc5b1461023d57806323691566146102455780632dc6bb7b1461024d5780634ba3a1261461025557610206565b8063033cfbed1461020b57806308c4fff014610225578063164bc2ae1461022d578063183d2d1c14610235575b600080fd5b6102136105c1565b60408051918252519081900360200190f35b6102136105cb565b6102136105d0565b6102136105d6565b6102136105dc565b6102136105e6565b6102136105ec565b6102726004803603602081101561026b57600080fd5b50356105f2565b6040805167ffffffffffffffff9d8e1681529b8d1660208d0152998c168b8b0152978b1660608b0152958a1660808a015293891660a089015291881660c0880152871660e0870152861661010086015285166101208501529093166101408301529115156101608201529051908190036101800190f35b6102f161066a565b604080516001600160a01b039092168252519081900360200190f35b6102136004803603602081101561032357600080fd5b503561067e565b610332610690565b604080519115158252519081900360200190f35b610213610699565b61021361069e565b6102136106a8565b6102136106af565b6102136106b5565b61038b6004803603602081101561038457600080fd5b50356106ba565b6040805167ffffffffffffffff8d1681528b15156020808301919091528b151592820192909252891515606082015288151560808201526001600160801b03881660a08201526001600160a01b0380881660c0830152861660e082015261010081018590526101208101849052610160610140820181815284519183019190915283519192909161018084019185019080838360005b83811015610439578181015183820152602001610421565b50505050905090810190601f1680156104665780820380516001836020036101000a031916815260200191505b509c5050505050505050505050505060405180910390f35b6102136004803603602081101561049457600080fd5b50356107dc565b6102f16107ee565b6104c0600480360360208110156104b957600080fd5b50356107fd565b60408051961515875267ffffffffffffffff958616602088015293851686850152918416606086015290921660808401526001600160a01b0390911660a0830152519081900360c00190f35b610213610865565b61021361086a565b610213610870565b6102136004803603602081101561053a57600080fd5b5035610876565b6102f16004803603602081101561055757600080fd5b50356001600160a01b0316610888565b6102f16108a3565b6102f16108a8565b6102f16108b7565b6104c06004803603602081101561059557600080fd5b50356108c6565b61038b600480360360208110156105b257600080fd5b50356108d3565b6102136108e0565b6509184e72a00081565b600f81565b600f5481565b60055481565b6551dac207a00081565b600c5481565b60125481565b60076020526000908152604090208054600182015460029092015467ffffffffffffffff80831693600160401b808504831694600160801b808204851695600160c01b92839004861695858116958581048216958482048316959182900483169484841694918204841693908204169160ff9104168c565b60005461010090046001600160a01b031681565b60066020526000908152604090205481565b60005460ff1681565b600a81565b6512309ce5400081565b620186a081565b60045481565b601481565b600881815481106106c757fe5b6000918252602091829020600691909102018054600180830154600280850154600386015460048701546005880180546040805161010099831615999099026000190190911695909504601f81018b90048b0288018b0190955284875267ffffffffffffffff88169a50600160401b880460ff9081169a600160481b8a0482169a600160501b8b0483169a600160581b810490931699600160601b9093046001600160801b0316986001600160a01b03908116989716969093918301828280156107d25780601f106107a7576101008083540402835291602001916107d2565b820191906000526020600020905b8154815290600101906020018083116107b557829003601f168201915b505050505090508b565b600d6020526000908152604090205481565b6003546001600160a01b031681565b600b818154811061080a57fe5b60009182526020909120600290910201805460019091015460ff8216925067ffffffffffffffff6101008304811692600160481b8104821692600160881b9091048216918116906001600160a01b03600160401b9091041686565b603c81565b60135481565b60105481565b600e6020526000908152604090205481565b6014602052600090815260409020546001600160a01b031681565b600081565b6002546001600160a01b031681565b6001546001600160a01b031681565b600a818154811061080a57fe5b600981815481106106c757fe5b6011548156fea265627a7a7230582032b282f30304a93e27cf3923dd26b846ceee3a3ccb8149f011de32df3002ce4764736f6c63430005090032"

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

// CPEXIT is a free data retrieval call binding the contract method 0x8155717d.
//
// Solidity: function CP_EXIT() constant returns(uint256)
func (_RootChainStorage *RootChainStorageCaller) CPEXIT(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChainStorage.contract.Call(opts, out, "CP_EXIT")
	return *ret0, err
}

// CPEXIT is a free data retrieval call binding the contract method 0x8155717d.
//
// Solidity: function CP_EXIT() constant returns(uint256)
func (_RootChainStorage *RootChainStorageSession) CPEXIT() (*big.Int, error) {
	return _RootChainStorage.Contract.CPEXIT(&_RootChainStorage.CallOpts)
}

// CPEXIT is a free data retrieval call binding the contract method 0x8155717d.
//
// Solidity: function CP_EXIT() constant returns(uint256)
func (_RootChainStorage *RootChainStorageCallerSession) CPEXIT() (*big.Int, error) {
	return _RootChainStorage.Contract.CPEXIT(&_RootChainStorage.CallOpts)
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

// EROIdToFinalize is a free data retrieval call binding the contract method 0x2dc6bb7b.
//
// Solidity: function EROIdToFinalize() constant returns(uint256)
func (_RootChainStorage *RootChainStorageCaller) EROIdToFinalize(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChainStorage.contract.Call(opts, out, "EROIdToFinalize")
	return *ret0, err
}

// EROIdToFinalize is a free data retrieval call binding the contract method 0x2dc6bb7b.
//
// Solidity: function EROIdToFinalize() constant returns(uint256)
func (_RootChainStorage *RootChainStorageSession) EROIdToFinalize() (*big.Int, error) {
	return _RootChainStorage.Contract.EROIdToFinalize(&_RootChainStorage.CallOpts)
}

// EROIdToFinalize is a free data retrieval call binding the contract method 0x2dc6bb7b.
//
// Solidity: function EROIdToFinalize() constant returns(uint256)
func (_RootChainStorage *RootChainStorageCallerSession) EROIdToFinalize() (*big.Int, error) {
	return _RootChainStorage.Contract.EROIdToFinalize(&_RootChainStorage.CallOpts)
}

// EROs is a free data retrieval call binding the contract method 0xb443f3cc.
//
// Solidity: function EROs(uint256 ) constant returns(uint64 timestamp, bool isExit, bool isTransfer, bool finalized, bool challenged, uint128 value, address requestor, address to, bytes32 trieKey, bytes32 hash, bytes trieValue)
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
	Hash       [32]byte
	TrieValue  []byte
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
		Hash       [32]byte
		TrieValue  []byte
	})
	out := ret
	err := _RootChainStorage.contract.Call(opts, out, "EROs", arg0)
	return *ret, err
}

// EROs is a free data retrieval call binding the contract method 0xb443f3cc.
//
// Solidity: function EROs(uint256 ) constant returns(uint64 timestamp, bool isExit, bool isTransfer, bool finalized, bool challenged, uint128 value, address requestor, address to, bytes32 trieKey, bytes32 hash, bytes trieValue)
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
	Hash       [32]byte
	TrieValue  []byte
}, error) {
	return _RootChainStorage.Contract.EROs(&_RootChainStorage.CallOpts, arg0)
}

// EROs is a free data retrieval call binding the contract method 0xb443f3cc.
//
// Solidity: function EROs(uint256 ) constant returns(uint64 timestamp, bool isExit, bool isTransfer, bool finalized, bool challenged, uint128 value, address requestor, address to, bytes32 trieKey, bytes32 hash, bytes trieValue)
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
	Hash       [32]byte
	TrieValue  []byte
}, error) {
	return _RootChainStorage.Contract.EROs(&_RootChainStorage.CallOpts, arg0)
}

// ERUIdToFinalize is a free data retrieval call binding the contract method 0xc54626cc.
//
// Solidity: function ERUIdToFinalize() constant returns(uint256)
func (_RootChainStorage *RootChainStorageCaller) ERUIdToFinalize(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChainStorage.contract.Call(opts, out, "ERUIdToFinalize")
	return *ret0, err
}

// ERUIdToFinalize is a free data retrieval call binding the contract method 0xc54626cc.
//
// Solidity: function ERUIdToFinalize() constant returns(uint256)
func (_RootChainStorage *RootChainStorageSession) ERUIdToFinalize() (*big.Int, error) {
	return _RootChainStorage.Contract.ERUIdToFinalize(&_RootChainStorage.CallOpts)
}

// ERUIdToFinalize is a free data retrieval call binding the contract method 0xc54626cc.
//
// Solidity: function ERUIdToFinalize() constant returns(uint256)
func (_RootChainStorage *RootChainStorageCallerSession) ERUIdToFinalize() (*big.Int, error) {
	return _RootChainStorage.Contract.ERUIdToFinalize(&_RootChainStorage.CallOpts)
}

// ERUs is a free data retrieval call binding the contract method 0xf4f31de4.
//
// Solidity: function ERUs(uint256 ) constant returns(uint64 timestamp, bool isExit, bool isTransfer, bool finalized, bool challenged, uint128 value, address requestor, address to, bytes32 trieKey, bytes32 hash, bytes trieValue)
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
	Hash       [32]byte
	TrieValue  []byte
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
		Hash       [32]byte
		TrieValue  []byte
	})
	out := ret
	err := _RootChainStorage.contract.Call(opts, out, "ERUs", arg0)
	return *ret, err
}

// ERUs is a free data retrieval call binding the contract method 0xf4f31de4.
//
// Solidity: function ERUs(uint256 ) constant returns(uint64 timestamp, bool isExit, bool isTransfer, bool finalized, bool challenged, uint128 value, address requestor, address to, bytes32 trieKey, bytes32 hash, bytes trieValue)
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
	Hash       [32]byte
	TrieValue  []byte
}, error) {
	return _RootChainStorage.Contract.ERUs(&_RootChainStorage.CallOpts, arg0)
}

// ERUs is a free data retrieval call binding the contract method 0xf4f31de4.
//
// Solidity: function ERUs(uint256 ) constant returns(uint64 timestamp, bool isExit, bool isTransfer, bool finalized, bool challenged, uint128 value, address requestor, address to, bytes32 trieKey, bytes32 hash, bytes trieValue)
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
	Hash       [32]byte
	TrieValue  []byte
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
// Solidity: function ORBs(uint256 ) constant returns(bool submitted, uint64 numEnter, uint64 epochNumber, uint64 requestStart, uint64 requestEnd, address trie)
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
// Solidity: function ORBs(uint256 ) constant returns(bool submitted, uint64 numEnter, uint64 epochNumber, uint64 requestStart, uint64 requestEnd, address trie)
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
// Solidity: function ORBs(uint256 ) constant returns(bool submitted, uint64 numEnter, uint64 epochNumber, uint64 requestStart, uint64 requestEnd, address trie)
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
// Solidity: function URBs(uint256 ) constant returns(bool submitted, uint64 numEnter, uint64 epochNumber, uint64 requestStart, uint64 requestEnd, address trie)
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
// Solidity: function URBs(uint256 ) constant returns(bool submitted, uint64 numEnter, uint64 epochNumber, uint64 requestStart, uint64 requestEnd, address trie)
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
// Solidity: function URBs(uint256 ) constant returns(bool submitted, uint64 numEnter, uint64 epochNumber, uint64 requestStart, uint64 requestEnd, address trie)
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

// EtherToken is a free data retrieval call binding the contract method 0xb8066bcb.
//
// Solidity: function etherToken() constant returns(address)
func (_RootChainStorage *RootChainStorageCaller) EtherToken(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RootChainStorage.contract.Call(opts, out, "etherToken")
	return *ret0, err
}

// EtherToken is a free data retrieval call binding the contract method 0xb8066bcb.
//
// Solidity: function etherToken() constant returns(address)
func (_RootChainStorage *RootChainStorageSession) EtherToken() (common.Address, error) {
	return _RootChainStorage.Contract.EtherToken(&_RootChainStorage.CallOpts)
}

// EtherToken is a free data retrieval call binding the contract method 0xb8066bcb.
//
// Solidity: function etherToken() constant returns(address)
func (_RootChainStorage *RootChainStorageCallerSession) EtherToken() (common.Address, error) {
	return _RootChainStorage.Contract.EtherToken(&_RootChainStorage.CallOpts)
}

// FirstFilledORENumber is a free data retrieval call binding the contract method 0x72ecb9a8.
//
// Solidity: function firstFilledORENumber(uint256 ) constant returns(uint256)
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
// Solidity: function firstFilledORENumber(uint256 ) constant returns(uint256)
func (_RootChainStorage *RootChainStorageSession) FirstFilledORENumber(arg0 *big.Int) (*big.Int, error) {
	return _RootChainStorage.Contract.FirstFilledORENumber(&_RootChainStorage.CallOpts, arg0)
}

// FirstFilledORENumber is a free data retrieval call binding the contract method 0x72ecb9a8.
//
// Solidity: function firstFilledORENumber(uint256 ) constant returns(uint256)
func (_RootChainStorage *RootChainStorageCallerSession) FirstFilledORENumber(arg0 *big.Int) (*big.Int, error) {
	return _RootChainStorage.Contract.FirstFilledORENumber(&_RootChainStorage.CallOpts, arg0)
}

// FirstNonEmptyRequestEpoch is a free data retrieval call binding the contract method 0xca6f6380.
//
// Solidity: function firstNonEmptyRequestEpoch(uint256 ) constant returns(uint256)
func (_RootChainStorage *RootChainStorageCaller) FirstNonEmptyRequestEpoch(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChainStorage.contract.Call(opts, out, "firstNonEmptyRequestEpoch", arg0)
	return *ret0, err
}

// FirstNonEmptyRequestEpoch is a free data retrieval call binding the contract method 0xca6f6380.
//
// Solidity: function firstNonEmptyRequestEpoch(uint256 ) constant returns(uint256)
func (_RootChainStorage *RootChainStorageSession) FirstNonEmptyRequestEpoch(arg0 *big.Int) (*big.Int, error) {
	return _RootChainStorage.Contract.FirstNonEmptyRequestEpoch(&_RootChainStorage.CallOpts, arg0)
}

// FirstNonEmptyRequestEpoch is a free data retrieval call binding the contract method 0xca6f6380.
//
// Solidity: function firstNonEmptyRequestEpoch(uint256 ) constant returns(uint256)
func (_RootChainStorage *RootChainStorageCallerSession) FirstNonEmptyRequestEpoch(arg0 *big.Int) (*big.Int, error) {
	return _RootChainStorage.Contract.FirstNonEmptyRequestEpoch(&_RootChainStorage.CallOpts, arg0)
}

// Forks is a free data retrieval call binding the contract method 0x4ba3a126.
//
// Solidity: function forks(uint256 ) constant returns(uint64 forkedBlock, uint64 firstEpoch, uint64 lastEpoch, uint64 firstBlock, uint64 lastBlock, uint64 lastFinalizedEpoch, uint64 lastFinalizedBlock, uint64 timestamp, uint64 firstEnterEpoch, uint64 lastEnterEpoch, uint64 nextBlockToRebase, bool rebased)
func (_RootChainStorage *RootChainStorageCaller) Forks(opts *bind.CallOpts, arg0 *big.Int) (struct {
	ForkedBlock        uint64
	FirstEpoch         uint64
	LastEpoch          uint64
	FirstBlock         uint64
	LastBlock          uint64
	LastFinalizedEpoch uint64
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
		LastFinalizedEpoch uint64
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
// Solidity: function forks(uint256 ) constant returns(uint64 forkedBlock, uint64 firstEpoch, uint64 lastEpoch, uint64 firstBlock, uint64 lastBlock, uint64 lastFinalizedEpoch, uint64 lastFinalizedBlock, uint64 timestamp, uint64 firstEnterEpoch, uint64 lastEnterEpoch, uint64 nextBlockToRebase, bool rebased)
func (_RootChainStorage *RootChainStorageSession) Forks(arg0 *big.Int) (struct {
	ForkedBlock        uint64
	FirstEpoch         uint64
	LastEpoch          uint64
	FirstBlock         uint64
	LastBlock          uint64
	LastFinalizedEpoch uint64
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
// Solidity: function forks(uint256 ) constant returns(uint64 forkedBlock, uint64 firstEpoch, uint64 lastEpoch, uint64 firstBlock, uint64 lastBlock, uint64 lastFinalizedEpoch, uint64 lastFinalizedBlock, uint64 timestamp, uint64 firstEnterEpoch, uint64 lastEnterEpoch, uint64 nextBlockToRebase, bool rebased)
func (_RootChainStorage *RootChainStorageCallerSession) Forks(arg0 *big.Int) (struct {
	ForkedBlock        uint64
	FirstEpoch         uint64
	LastEpoch          uint64
	FirstBlock         uint64
	LastBlock          uint64
	LastFinalizedEpoch uint64
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

// LastAppliedEpochNumber is a free data retrieval call binding the contract method 0xc8ad329f.
//
// Solidity: function lastAppliedEpochNumber() constant returns(uint256)
func (_RootChainStorage *RootChainStorageCaller) LastAppliedEpochNumber(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChainStorage.contract.Call(opts, out, "lastAppliedEpochNumber")
	return *ret0, err
}

// LastAppliedEpochNumber is a free data retrieval call binding the contract method 0xc8ad329f.
//
// Solidity: function lastAppliedEpochNumber() constant returns(uint256)
func (_RootChainStorage *RootChainStorageSession) LastAppliedEpochNumber() (*big.Int, error) {
	return _RootChainStorage.Contract.LastAppliedEpochNumber(&_RootChainStorage.CallOpts)
}

// LastAppliedEpochNumber is a free data retrieval call binding the contract method 0xc8ad329f.
//
// Solidity: function lastAppliedEpochNumber() constant returns(uint256)
func (_RootChainStorage *RootChainStorageCallerSession) LastAppliedEpochNumber() (*big.Int, error) {
	return _RootChainStorage.Contract.LastAppliedEpochNumber(&_RootChainStorage.CallOpts)
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

// LastNonEmptyRequestEpoch is a free data retrieval call binding the contract method 0xb6715647.
//
// Solidity: function lastNonEmptyRequestEpoch(uint256 ) constant returns(uint256)
func (_RootChainStorage *RootChainStorageCaller) LastNonEmptyRequestEpoch(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChainStorage.contract.Call(opts, out, "lastNonEmptyRequestEpoch", arg0)
	return *ret0, err
}

// LastNonEmptyRequestEpoch is a free data retrieval call binding the contract method 0xb6715647.
//
// Solidity: function lastNonEmptyRequestEpoch(uint256 ) constant returns(uint256)
func (_RootChainStorage *RootChainStorageSession) LastNonEmptyRequestEpoch(arg0 *big.Int) (*big.Int, error) {
	return _RootChainStorage.Contract.LastNonEmptyRequestEpoch(&_RootChainStorage.CallOpts, arg0)
}

// LastNonEmptyRequestEpoch is a free data retrieval call binding the contract method 0xb6715647.
//
// Solidity: function lastNonEmptyRequestEpoch(uint256 ) constant returns(uint256)
func (_RootChainStorage *RootChainStorageCallerSession) LastNonEmptyRequestEpoch(arg0 *big.Int) (*big.Int, error) {
	return _RootChainStorage.Contract.LastNonEmptyRequestEpoch(&_RootChainStorage.CallOpts, arg0)
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
// Solidity: function requestableContracts(address ) constant returns(address)
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
// Solidity: function requestableContracts(address ) constant returns(address)
func (_RootChainStorage *RootChainStorageSession) RequestableContracts(arg0 common.Address) (common.Address, error) {
	return _RootChainStorage.Contract.RequestableContracts(&_RootChainStorage.CallOpts, arg0)
}

// RequestableContracts is a free data retrieval call binding the contract method 0xda0185f8.
//
// Solidity: function requestableContracts(address ) constant returns(address)
func (_RootChainStorage *RootChainStorageCallerSession) RequestableContracts(arg0 common.Address) (common.Address, error) {
	return _RootChainStorage.Contract.RequestableContracts(&_RootChainStorage.CallOpts, arg0)
}

// SubmitHandler is a free data retrieval call binding the contract method 0xe259faf7.
//
// Solidity: function submitHandler() constant returns(address)
func (_RootChainStorage *RootChainStorageCaller) SubmitHandler(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RootChainStorage.contract.Call(opts, out, "submitHandler")
	return *ret0, err
}

// SubmitHandler is a free data retrieval call binding the contract method 0xe259faf7.
//
// Solidity: function submitHandler() constant returns(address)
func (_RootChainStorage *RootChainStorageSession) SubmitHandler() (common.Address, error) {
	return _RootChainStorage.Contract.SubmitHandler(&_RootChainStorage.CallOpts)
}

// SubmitHandler is a free data retrieval call binding the contract method 0xe259faf7.
//
// Solidity: function submitHandler() constant returns(address)
func (_RootChainStorage *RootChainStorageCallerSession) SubmitHandler() (common.Address, error) {
	return _RootChainStorage.Contract.SubmitHandler(&_RootChainStorage.CallOpts)
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
