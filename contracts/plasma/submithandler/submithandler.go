// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package submithandler

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

// SubmitHandlerABI is the input ABI used to generate the binding from.
const SubmitHandlerABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"COST_URB_PREPARE\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CP_COMPUTATION\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_pos1\",\"type\":\"uint256\"},{\"name\":\"_pos2\",\"type\":\"uint256\"},{\"name\":\"_epochStateRoot\",\"type\":\"bytes32\"},{\"name\":\"_epochTransactionsRoot\",\"type\":\"bytes32\"},{\"name\":\"_epochReceiptsRoot\",\"type\":\"bytes32\"}],\"name\":\"submitNRE\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastAppliedForkNumber\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"currentFork\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"COST_URB\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"numEnterForORB\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"EROIdToFinalize\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"forks\",\"outputs\":[{\"name\":\"forkedBlock\",\"type\":\"uint64\"},{\"name\":\"firstEpoch\",\"type\":\"uint64\"},{\"name\":\"lastEpoch\",\"type\":\"uint64\"},{\"name\":\"firstBlock\",\"type\":\"uint64\"},{\"name\":\"lastBlock\",\"type\":\"uint64\"},{\"name\":\"lastFinalizedEpoch\",\"type\":\"uint64\"},{\"name\":\"lastFinalizedBlock\",\"type\":\"uint64\"},{\"name\":\"timestamp\",\"type\":\"uint64\"},{\"name\":\"firstEnterEpoch\",\"type\":\"uint64\"},{\"name\":\"lastEnterEpoch\",\"type\":\"uint64\"},{\"name\":\"nextBlockToRebase\",\"type\":\"uint64\"},{\"name\":\"rebased\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"operator\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_pos\",\"type\":\"uint256\"},{\"name\":\"_statesRoot\",\"type\":\"bytes32\"},{\"name\":\"_transactionsRoot\",\"type\":\"bytes32\"},{\"name\":\"_receiptsRoot\",\"type\":\"bytes32\"}],\"name\":\"submitURB\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"firstFilledORENumber\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"development\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CP_EXIT\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"COST_ERU\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"REQUEST_GAS\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"COST_NRB\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_pos\",\"type\":\"uint256\"},{\"name\":\"_statesRoot\",\"type\":\"bytes32\"},{\"name\":\"_transactionsRoot\",\"type\":\"bytes32\"},{\"name\":\"_receiptsRoot\",\"type\":\"bytes32\"}],\"name\":\"submitORB\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"NRELength\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CP_WITHHOLDING\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"COST_ORB\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"EROs\",\"outputs\":[{\"name\":\"timestamp\",\"type\":\"uint64\"},{\"name\":\"isExit\",\"type\":\"bool\"},{\"name\":\"isTransfer\",\"type\":\"bool\"},{\"name\":\"finalized\",\"type\":\"bool\"},{\"name\":\"challenged\",\"type\":\"bool\"},{\"name\":\"value\",\"type\":\"uint128\"},{\"name\":\"requestor\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"trieKey\",\"type\":\"bytes32\"},{\"name\":\"hash\",\"type\":\"bytes32\"},{\"name\":\"trieValue\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"lastNonEmptyRequestEpoch\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"etherToken\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"URBs\",\"outputs\":[{\"name\":\"submitted\",\"type\":\"bool\"},{\"name\":\"numEnter\",\"type\":\"uint64\"},{\"name\":\"epochNumber\",\"type\":\"uint64\"},{\"name\":\"requestStart\",\"type\":\"uint64\"},{\"name\":\"requestEnd\",\"type\":\"uint64\"},{\"name\":\"trie\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PREPARE_TIMEOUT\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ERUIdToFinalize\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastAppliedEpochNumber\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"firstNonEmptyRequestEpoch\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"COST_ERO\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"requestableContracts\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"NULL_ADDRESS\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"submitHandler\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"epochHandler\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"ORBs\",\"outputs\":[{\"name\":\"submitted\",\"type\":\"bool\"},{\"name\":\"numEnter\",\"type\":\"uint64\"},{\"name\":\"epochNumber\",\"type\":\"uint64\"},{\"name\":\"requestStart\",\"type\":\"uint64\"},{\"name\":\"requestEnd\",\"type\":\"uint64\"},{\"name\":\"trie\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"ERUs\",\"outputs\":[{\"name\":\"timestamp\",\"type\":\"uint64\"},{\"name\":\"isExit\",\"type\":\"bool\"},{\"name\":\"isTransfer\",\"type\":\"bool\"},{\"name\":\"finalized\",\"type\":\"bool\"},{\"name\":\"challenged\",\"type\":\"bool\"},{\"name\":\"value\",\"type\":\"uint128\"},{\"name\":\"requestor\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"trieKey\",\"type\":\"bytes32\"},{\"name\":\"hash\",\"type\":\"bytes32\"},{\"name\":\"trieValue\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastAppliedBlockNumber\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_epochHandler\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"userActivated\",\"type\":\"bool\"}],\"name\":\"SessionTimeout\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newFork\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"epochNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"forkedBlockNumber\",\"type\":\"uint256\"}],\"name\":\"Forked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"forkNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"epochNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"startBlockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"endBlockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"requestStart\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"requestEnd\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"epochIsEmpty\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"isRequest\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"userActivated\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"rebase\",\"type\":\"bool\"}],\"name\":\"EpochPrepared\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"forkNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"epochNumber\",\"type\":\"uint256\"}],\"name\":\"EpochFilling\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"forkNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"epochNumber\",\"type\":\"uint256\"}],\"name\":\"EpochFilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"forkNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"epochNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"startBlockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"endBlockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"requestStart\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"requestEnd\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"epochIsEmpty\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"isRequest\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"userActivated\",\"type\":\"bool\"}],\"name\":\"EpochRebased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"fork\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"epochNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"isRequest\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"userActivated\",\"type\":\"bool\"}],\"name\":\"BlockSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"requestor\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"weiAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"trieKey\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"trieValue\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"isExit\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"userActivated\",\"type\":\"bool\"}],\"name\":\"RequestCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"requestor\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"trieKey\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"trieValue\",\"type\":\"bytes32\"}],\"name\":\"ERUCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"forkNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"BlockFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"forkNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"epochNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"startBlockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"endBlockNumber\",\"type\":\"uint256\"}],\"name\":\"EpochFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"userActivated\",\"type\":\"bool\"}],\"name\":\"RequestFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"userActivated\",\"type\":\"bool\"}],\"name\":\"RequestApplied\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"userActivated\",\"type\":\"bool\"}],\"name\":\"RequestChallenged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"contractInRootchain\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"contractInChildchain\",\"type\":\"address\"}],\"name\":\"RequestableContractMapped\",\"type\":\"event\"}]"

// SubmitHandlerFuncSigs maps the 4-byte function signature to its string representation.
var SubmitHandlerFuncSigs = map[string]string{
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
	"0eaf45a8": "submitNRE(uint256,uint256,bytes32,bytes32,bytes32)",
	"a820c067": "submitORB(uint256,bytes32,bytes32,bytes32)",
	"6f3e4b90": "submitURB(uint256,bytes32,bytes32,bytes32)",
}

// SubmitHandlerBin is the compiled bytecode used for deploying new contracts.
var SubmitHandlerBin = "0x608060405234801561001057600080fd5b50604051611f1e380380611f1e8339818101604052602081101561003357600080fd5b5051600180546001600160a01b039092166001600160a01b03199283161790556002805490911630179055611eb18061006d6000396000f3fe60806040526004361061021a5760003560e01c8063ab96da2d11610123578063c8ad329f116100ab578063e259faf71161006f578063e259faf714610765578063e7b88b801461077a578063ea7f22a81461078f578063f4f31de4146107b9578063fb788a27146107e35761021a565b8063c8ad329f146106de578063ca6f6380146106f3578063d691acd81461021f578063da0185f81461071d578063de0ce17d146107505761021a565b8063b6715647116100f2578063b671564714610600578063b8066bcb1461062a578063c0e860641461063f578063c2bc88fa146106b4578063c54626cc146106c95761021a565b8063ab96da2d146104ba578063b17fa6e9146104cf578063b2ae9ba81461021f578063b443f3cc146104e45761021a565b8063570ca735116101a65780638155717d116101755780638155717d1461044c5780638b5172d0146104615780638eb288ca1461047657806394be3aa51461021f578063a820c0671461048b5761021a565b8063570ca735146103ad5780636f3e4b90146103de57806372ecb9a81461040d5780637b929c27146104375761021a565b8063183d2d1c116101ed578063183d2d1c146102b9578063192adc5b146102ce57806323691566146102e35780632dc6bb7b146102f85780634ba3a1261461030d5761021a565b8063033cfbed1461021f57806308c4fff0146102465780630eaf45a81461025b578063164bc2ae146102a4575b600080fd5b34801561022b57600080fd5b506102346107f8565b60408051918252519081900360200190f35b34801561025257600080fd5b50610234610802565b610290600480360360a081101561027157600080fd5b5080359060208101359060408101359060608101359060800135610807565b604080519115158252519081900360200190f35b3480156102b057600080fd5b506102346108bb565b3480156102c557600080fd5b506102346108c1565b3480156102da57600080fd5b506102346108c7565b3480156102ef57600080fd5b506102346108d1565b34801561030457600080fd5b506102346108d7565b34801561031957600080fd5b506103376004803603602081101561033057600080fd5b50356108dd565b604080516001600160401b039d8e1681529b8d1660208d0152998c168b8b0152978b1660608b0152958a1660808a015293891660a089015291881660c0880152871660e0870152861661010086015285166101208501529093166101408301529115156101608201529051908190036101800190f35b3480156103b957600080fd5b506103c2610954565b604080516001600160a01b039092168252519081900360200190f35b610290600480360360808110156103f457600080fd5b5080359060208101359060408101359060600135610968565b34801561041957600080fd5b506102346004803603602081101561043057600080fd5b5035610c74565b34801561044357600080fd5b50610290610c86565b34801561045857600080fd5b50610234610c8f565b34801561046d57600080fd5b50610234610c94565b34801561048257600080fd5b50610234610c9e565b610290600480360360808110156104a157600080fd5b5080359060208101359060408101359060600135610ca5565b3480156104c657600080fd5b50610234611101565b3480156104db57600080fd5b50610234611107565b3480156104f057600080fd5b5061050e6004803603602081101561050757600080fd5b503561110c565b604080516001600160401b038d1681528b15156020808301919091528b151592820192909252891515606082015288151560808201526001600160801b03881660a08201526001600160a01b0380881660c0830152861660e082015261010081018590526101208101849052610160610140820181815284519183019190915283519192909161018084019185019080838360005b838110156105bb5781810151838201526020016105a3565b50505050905090810190601f1680156105e85780820380516001836020036101000a031916815260200191505b509c5050505050505050505050505060405180910390f35b34801561060c57600080fd5b506102346004803603602081101561062357600080fd5b503561122d565b34801561063657600080fd5b506103c261123f565b34801561064b57600080fd5b506106696004803603602081101561066257600080fd5b503561124e565b6040805196151587526001600160401b03958616602088015293851686850152918416606086015290921660808401526001600160a01b0390911660a0830152519081900360c00190f35b3480156106c057600080fd5b506102346112b5565b3480156106d557600080fd5b506102346112ba565b3480156106ea57600080fd5b506102346112c0565b3480156106ff57600080fd5b506102346004803603602081101561071657600080fd5b50356112c6565b34801561072957600080fd5b506103c26004803603602081101561074057600080fd5b50356001600160a01b03166112d8565b34801561075c57600080fd5b506103c26112f3565b34801561077157600080fd5b506103c26112f8565b34801561078657600080fd5b506103c2611307565b34801561079b57600080fd5b50610669600480360360208110156107b257600080fd5b5035611316565b3480156107c557600080fd5b5061050e600480360360208110156107dc57600080fd5b5035611323565b3480156107ef57600080fd5b50610234611330565b6509184e72a00081565b600f81565b600080600061081588611336565b91509150816005541461086f576040805162461bcd60e51b815260206004820152601960248201527f63757272656e74466f726b203d3d20666f726b4e756d62657200000000000000604482015290519081900360640190fd5b60008061087b89611336565b600086815260076020526040902091935091506108a390848a8a8a878763ffffffff61134c16565b6108ab611492565b5060019998505050505050505050565b600f5481565b60055481565b6551dac207a00081565b600c5481565b60125481565b6007602052600090815260409020805460018201546002909201546001600160401b0380831693600160401b808504831694600160801b808204851695600160c01b92839004861695858116958581048216958482048316959182900483169484841694918204841693908204169160ff9104168c565b60005461010090046001600160a01b031681565b600080600061097687611336565b600554919350915060010182148080610990575082600554145b61099957600080fd5b60008381526007602052604090208115610a65575060058390556000838152600760205260409020426109ca611597565b6001830154600160c01b90046001600160401b031601116109ea57600080fd5b8054600160c01b90046001600160401b03168314610a0757600080fd5b8054604080518681526001600160401b03600160801b840481166020830152600160c01b90930490921682820152517f0647d42ab02f6e0ae76959757dcb6aa6feac1d4ba6f077f1223fb4b1b429f06c9181900360600190a1610a9a565b600181810154610a86916001600160401b039091169063ffffffff61159d16565b6001600160401b03168314610a9a57600080fd5b8054600160801b90046001600160401b0316600090815260038201602052604090208054600160d01b900460ff16610ad157600080fd5b8054600160d81b900460ff16610ae657600080fd5b600084815260048381016020526040822084548154600283018e9055600383018d90559282018b9055600160801b908190046001600160401b0390811667ffffffffffffffff199485161767ffffffffffffffff60801b19164282169290920291909117825560058201805461010060ff19909116600190811761ff001916919091179091558601805490931690881617909155905460ff16610be15760018301546001600160401b039081166000908152600485016020526040902054600b8054610be1939192600160401b9004909116908110610bc157fe5b6000918252602082208c926002909202019060099063ffffffff6115c216565b825460408051888152600160801b9092046001600160401b031660208301528181018790526001606083018190526080830152517f3d4a04291c66b06f39a4ecb817875b12b5485a05ec563133a56a905305c48e559160a0908290030190a18154600160401b90046001600160401b0316851415610c6157610c616116e4565b600196505050505050505b949350505050565b60066020526000908152604090205481565b60005460ff1681565b600a81565b6512309ce5400081565b620186a081565b6000806000610cb387611336565b60055491935091508214610cc657600080fd5b60008281526007602052604090206001908101546001600160401b03908116909101168114610cf457600080fd5b60008281526007602052604081209080841580610d42575060001985016000908152600760205260409020546001600160401b031615801590610d4257506002830154600160c01b900460ff165b15610ee357610d5e838a8a8a600160008063ffffffff61177816565b60005490955090925060ff16610dcc5760018301546001600160401b039081166000908152600485016020526040902054600a8054610dcc939192600160401b9004909116908110610dac57fe5b6000918252602082208b926002909202019060089063ffffffff6115c216565b506000838152600483016020526040902054600a8054600160401b9092046001600160401b03169182908110610dfe57fe5b60009182526020808320600292909202909101548483526003860182526040808420600190810180546001600160c01b0381166101009095046001600160401b03908116600160c01b9283900482160116029390931790925580518981529283018690528281018890526060830191909152608082019290925290517f3d4a04291c66b06f39a4ecb817875b12b5485a05ec563133a56a905305c48e559160a0908290030190a16000828152600384016020526040902054600160401b90046001600160401b0316841415610ed557610ed5611983565b600195505050505050610c6c565b610efa838a8a8a600160008163ffffffff61177816565b600019870160009081526007602090815260408083208484526004808a01845282852060028b01805460018301805467ffffffffffffffff19166001600160401b03600160801b9384900481169190911790915591540481168752918301909452918420548354600160401b918290049093160267ffffffffffffffff60401b199092169190911782559154929750929450919060ff16610fd7578054600a8054610fd792600160401b90046001600160401b0316908110610fb857fe5b600091825260209091208c91600202016008600163ffffffff6115c216565b6040805188815260208101869052808201889052600160608201526000608082015290517f3d4a04291c66b06f39a4ecb817875b12b5485a05ec563133a56a905305c48e559181900360a00190a16110378583600a63ffffffff611a0f16565b156110ef576000848152600386016020908152604080832080546001600160401b03808c16600160401b0267ffffffffffffffff60401b199092169190911780835560019283015484518e81529586018b905290821685850152606085018c905216608084015260a0830184905260c0830184905260e083015261010082019290925290517f030c1c69405c93021f28f57557240dee939a320b826a1fd0d39bf6e629ecab47918190036101200190a16110ef611c3a565b5060019b9a5050505050505050505050565b60045481565b601481565b6008818154811061111957fe5b6000918252602091829020600691909102018054600180830154600280850154600386015460048701546005880180546040805161010099831615999099026000190190911695909504601f81018b90048b0288018b019095528487526001600160401b0388169a50600160401b880460ff9081169a600160481b8a0482169a600160501b8b0483169a600160581b810490931699600160601b9093046001600160801b0316986001600160a01b03908116989716969093918301828280156112235780601f106111f857610100808354040283529160200191611223565b820191906000526020600020905b81548152906001019060200180831161120657829003601f168201915b505050505090508b565b600d6020526000908152604090205481565b6003546001600160a01b031681565b600b818154811061125b57fe5b60009182526020909120600290910201805460019091015460ff821692506001600160401b036101008304811692600160481b8104821692600160881b9091048216918116906001600160a01b03600160401b9091041686565b603c81565b60135481565b60105481565b600e6020526000908152604090205481565b6014602052600090815260409020546001600160a01b031681565b600081565b6002546001600160a01b031681565b6001546001600160a01b031681565b600a818154811061125b57fe5b6009818154811061111957fe5b60115481565b600160801b8104916001600160801b0390911690565b8654869061136b90600160801b90046001600160401b03166001611cce565b1461137557600080fd5b6001878101548391611396916001600160401b03169063ffffffff611cce16565b146113a057600080fd5b600086815260038801602052604090208054600160d01b900460ff16156113c657600080fd5b8054600160d81b900460ff16156113dc57600080fd5b8054600160e01b900460ff16156113f257600080fd5b80546001600160401b0316831461140857600080fd5b8054600160401b90046001600160401b0316821461142557600080fd5b60038101959095556004850193909355506005830155600690910180546001600160401b0342811667ffffffffffffffff19928316179092558454938216600160801b0267ffffffffffffffff60801b199094169390931784556001939093018054919093169116179055565b600154604080516b707265706172654f5245282960a01b8152815190819003600c018120600482526024820183526020820180516001600160e01b03166001600160e01b0319909216919091178152915181516000946060946001600160a01b039091169392918291908083835b6020831061151f5780518252601f199092019160209182019101611500565b6001836020036101000a038019825116818451168082178552505050505050905001915050600060405180830381855af49150503d806000811461157f576040519150601f19603f3d011682016040523d82523d6000602084013e611584565b606091505b50915091508161159357600080fd5b5050565b610e1090565b60008282016001600160401b0380851690821610156115bb57600080fd5b9392505050565b825460018401546001600160401b03600160881b909204821691166000836115ef57828203600101611600565b855461010090046001600160401b03165b90506000811161160f57600080fd5b60608160405190808252806020026020018201604052801561163b578160200160208202803883390190505b50905083805b8481116116c457861580611679575087818154811061165c57fe5b6000918252602090912060069091020154600160401b900460ff16155b156116bc5787818154811061168a57fe5b90600052602060002090600602016004015483878403815181106116aa57fe5b60209081029190910101526001909101905b600101611641565b50886116cf83611ce0565b146116d957600080fd5b505050505050505050565b6001546040805173707265706172654f52454166746572555245282960601b81528151908190036014018120600482526024820183526020820180516001600160e01b03166001600160e01b0319909216919091178152915181516000946060946001600160a01b039091169392918291908083836020831061151f5780518252601f199092019160209182019101611500565b86546001808901546001600160401b03600160801b9093048316926000926117a9929091169063ffffffff611cce16565b600083815260038b016020526040902080549192509060016001600160401b03600160401b9092048216011682141561181a575088546001600160401b03600193909301928316600160801b0267ffffffffffffffff60801b19909116178955600082815260038a01602052604090205b80546001600160401b031682101561183157600080fd5b838061184e57508054600160401b90046001600160401b03168211155b61185757600080fd5b805460ff600160d01b9091041615158615151461187357600080fd5b805460ff600160d81b9091041615158515151461188f57600080fd5b600082815260048b810160205260409091208054600282018c9055600382018b905591810189905567ffffffffffffffff199091166001600160401b038581169190911767ffffffffffffffff60801b1916600160801b42929092169190910217815560058101805460ff1916881580159190911761ff00191661010089151502179091556119575781546001830154825467ffffffffffffffff60401b1916600160401b6001600160401b03938416600160801b9093048416870192909203909216021781555b5050600198909801805467ffffffffffffffff19166001600160401b038a161790559795505050505050565b600154604080516b707265706172654e5245282960a01b8152815190819003600c018120600482526024820183526020820180516001600160e01b03166001600160e01b0319909216919091178152915181516000946060946001600160a01b039091169392918291908083836020831061151f5780518252601f199092019160209182019101611500565b60028301546001600160401b03600160801b909104811660008181526004850160205260408120549092165b6000818152600386016020526040902054600160c81b900460ff1615611c2e576000818152600386016020526040902054600160401b90046001600160401b03168210611aa25760020160008181526003860160205260409020546001600160401b031691505b5b6000818152600386016020526040902060010154600160c01b90046001600160401b0316158015611aeb57506000818152600386016020526040902054600160c81b900460ff165b15611b145760020160008181526003860160205260409020546001600160401b03169150611aa3565b6000818152600386016020526040902054600160c81b900460ff16611b3e576001925050506115bb565b6000818152600386016020526040902054600160401b90046001600160401b03165b808311611bca57600083815260048701602052604081205486548791600160401b90046001600160401b0316908110611b9557fe5b600091825260209091206002909102015461010090046001600160401b03161115611bbf57611bca565b600183019250611b60565b80831115611bf7575060020160008181526003860160205260409020546001600160401b03169150611a3b565b50506002850180546001600160401b03909216600160801b0267ffffffffffffffff60801b199092169190911790555060006115bb565b50600195945050505050565b6001546040805173707265706172654e52454166746572555245282960601b81528151908190036014018120600482526024820183526020820180516001600160e01b03166001600160e01b0319909216919091178152915181516000946060946001600160a01b039091169392918291908083836020831061151f5780518252601f199092019160209182019101611500565b6000828201838110156115bb57600080fd5b6000815160011415611d085781600081518110611cf957fe5b60200260200101519050611e77565b60606002835160010181611d1857fe5b04604051908082528060200260200182016040528015611d42578160200160208202803883390190505b50905060005b8351816001011015611dd557838181518110611d6057fe5b6020026020010151848260010181518110611d7757fe5b60200260200101516040516020018083815260200182815260200192505050604051602081830303815290604052805190602001208260028381611db757fe5b0481518110611dc257fe5b6020908102919091010152600201611d48565b6002845181611de057fe5b0660011415611e695783600185510381518110611df957fe5b602002602001015184600186510381518110611e1157fe5b60200260200101516040516020018083815260200182815260200192505050604051602081830303815290604052805190602001208260028381611e5157fe5b0481518110611e5c57fe5b6020026020010181815250505b611e7282611ce0565b925050505b91905056fea265627a7a72305820d90b68c5a7527cd89c24c96856f007035f6b23e238252b1ea80df4d759705d9a64736f6c63430005090032"

// DeploySubmitHandler deploys a new Ethereum contract, binding an instance of SubmitHandler to it.
func DeploySubmitHandler(auth *bind.TransactOpts, backend bind.ContractBackend, _epochHandler common.Address) (common.Address, *types.Transaction, *SubmitHandler, error) {
	parsed, err := abi.JSON(strings.NewReader(SubmitHandlerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SubmitHandlerBin), backend, _epochHandler)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SubmitHandler{SubmitHandlerCaller: SubmitHandlerCaller{contract: contract}, SubmitHandlerTransactor: SubmitHandlerTransactor{contract: contract}, SubmitHandlerFilterer: SubmitHandlerFilterer{contract: contract}}, nil
}

// SubmitHandler is an auto generated Go binding around an Ethereum contract.
type SubmitHandler struct {
	SubmitHandlerCaller     // Read-only binding to the contract
	SubmitHandlerTransactor // Write-only binding to the contract
	SubmitHandlerFilterer   // Log filterer for contract events
}

// SubmitHandlerCaller is an auto generated read-only Go binding around an Ethereum contract.
type SubmitHandlerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SubmitHandlerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SubmitHandlerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SubmitHandlerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SubmitHandlerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SubmitHandlerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SubmitHandlerSession struct {
	Contract     *SubmitHandler    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SubmitHandlerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SubmitHandlerCallerSession struct {
	Contract *SubmitHandlerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// SubmitHandlerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SubmitHandlerTransactorSession struct {
	Contract     *SubmitHandlerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// SubmitHandlerRaw is an auto generated low-level Go binding around an Ethereum contract.
type SubmitHandlerRaw struct {
	Contract *SubmitHandler // Generic contract binding to access the raw methods on
}

// SubmitHandlerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SubmitHandlerCallerRaw struct {
	Contract *SubmitHandlerCaller // Generic read-only contract binding to access the raw methods on
}

// SubmitHandlerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SubmitHandlerTransactorRaw struct {
	Contract *SubmitHandlerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSubmitHandler creates a new instance of SubmitHandler, bound to a specific deployed contract.
func NewSubmitHandler(address common.Address, backend bind.ContractBackend) (*SubmitHandler, error) {
	contract, err := bindSubmitHandler(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SubmitHandler{SubmitHandlerCaller: SubmitHandlerCaller{contract: contract}, SubmitHandlerTransactor: SubmitHandlerTransactor{contract: contract}, SubmitHandlerFilterer: SubmitHandlerFilterer{contract: contract}}, nil
}

// NewSubmitHandlerCaller creates a new read-only instance of SubmitHandler, bound to a specific deployed contract.
func NewSubmitHandlerCaller(address common.Address, caller bind.ContractCaller) (*SubmitHandlerCaller, error) {
	contract, err := bindSubmitHandler(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SubmitHandlerCaller{contract: contract}, nil
}

// NewSubmitHandlerTransactor creates a new write-only instance of SubmitHandler, bound to a specific deployed contract.
func NewSubmitHandlerTransactor(address common.Address, transactor bind.ContractTransactor) (*SubmitHandlerTransactor, error) {
	contract, err := bindSubmitHandler(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SubmitHandlerTransactor{contract: contract}, nil
}

// NewSubmitHandlerFilterer creates a new log filterer instance of SubmitHandler, bound to a specific deployed contract.
func NewSubmitHandlerFilterer(address common.Address, filterer bind.ContractFilterer) (*SubmitHandlerFilterer, error) {
	contract, err := bindSubmitHandler(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SubmitHandlerFilterer{contract: contract}, nil
}

// bindSubmitHandler binds a generic wrapper to an already deployed contract.
func bindSubmitHandler(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SubmitHandlerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SubmitHandler *SubmitHandlerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SubmitHandler.Contract.SubmitHandlerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SubmitHandler *SubmitHandlerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SubmitHandler.Contract.SubmitHandlerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SubmitHandler *SubmitHandlerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SubmitHandler.Contract.SubmitHandlerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SubmitHandler *SubmitHandlerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SubmitHandler.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SubmitHandler *SubmitHandlerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SubmitHandler.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SubmitHandler *SubmitHandlerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SubmitHandler.Contract.contract.Transact(opts, method, params...)
}

// COSTERO is a free data retrieval call binding the contract method 0xd691acd8.
//
// Solidity: function COST_ERO() constant returns(uint256)
func (_SubmitHandler *SubmitHandlerCaller) COSTERO(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SubmitHandler.contract.Call(opts, out, "COST_ERO")
	return *ret0, err
}

// COSTERO is a free data retrieval call binding the contract method 0xd691acd8.
//
// Solidity: function COST_ERO() constant returns(uint256)
func (_SubmitHandler *SubmitHandlerSession) COSTERO() (*big.Int, error) {
	return _SubmitHandler.Contract.COSTERO(&_SubmitHandler.CallOpts)
}

// COSTERO is a free data retrieval call binding the contract method 0xd691acd8.
//
// Solidity: function COST_ERO() constant returns(uint256)
func (_SubmitHandler *SubmitHandlerCallerSession) COSTERO() (*big.Int, error) {
	return _SubmitHandler.Contract.COSTERO(&_SubmitHandler.CallOpts)
}

// COSTERU is a free data retrieval call binding the contract method 0x8b5172d0.
//
// Solidity: function COST_ERU() constant returns(uint256)
func (_SubmitHandler *SubmitHandlerCaller) COSTERU(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SubmitHandler.contract.Call(opts, out, "COST_ERU")
	return *ret0, err
}

// COSTERU is a free data retrieval call binding the contract method 0x8b5172d0.
//
// Solidity: function COST_ERU() constant returns(uint256)
func (_SubmitHandler *SubmitHandlerSession) COSTERU() (*big.Int, error) {
	return _SubmitHandler.Contract.COSTERU(&_SubmitHandler.CallOpts)
}

// COSTERU is a free data retrieval call binding the contract method 0x8b5172d0.
//
// Solidity: function COST_ERU() constant returns(uint256)
func (_SubmitHandler *SubmitHandlerCallerSession) COSTERU() (*big.Int, error) {
	return _SubmitHandler.Contract.COSTERU(&_SubmitHandler.CallOpts)
}

// COSTNRB is a free data retrieval call binding the contract method 0x94be3aa5.
//
// Solidity: function COST_NRB() constant returns(uint256)
func (_SubmitHandler *SubmitHandlerCaller) COSTNRB(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SubmitHandler.contract.Call(opts, out, "COST_NRB")
	return *ret0, err
}

// COSTNRB is a free data retrieval call binding the contract method 0x94be3aa5.
//
// Solidity: function COST_NRB() constant returns(uint256)
func (_SubmitHandler *SubmitHandlerSession) COSTNRB() (*big.Int, error) {
	return _SubmitHandler.Contract.COSTNRB(&_SubmitHandler.CallOpts)
}

// COSTNRB is a free data retrieval call binding the contract method 0x94be3aa5.
//
// Solidity: function COST_NRB() constant returns(uint256)
func (_SubmitHandler *SubmitHandlerCallerSession) COSTNRB() (*big.Int, error) {
	return _SubmitHandler.Contract.COSTNRB(&_SubmitHandler.CallOpts)
}

// COSTORB is a free data retrieval call binding the contract method 0xb2ae9ba8.
//
// Solidity: function COST_ORB() constant returns(uint256)
func (_SubmitHandler *SubmitHandlerCaller) COSTORB(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SubmitHandler.contract.Call(opts, out, "COST_ORB")
	return *ret0, err
}

// COSTORB is a free data retrieval call binding the contract method 0xb2ae9ba8.
//
// Solidity: function COST_ORB() constant returns(uint256)
func (_SubmitHandler *SubmitHandlerSession) COSTORB() (*big.Int, error) {
	return _SubmitHandler.Contract.COSTORB(&_SubmitHandler.CallOpts)
}

// COSTORB is a free data retrieval call binding the contract method 0xb2ae9ba8.
//
// Solidity: function COST_ORB() constant returns(uint256)
func (_SubmitHandler *SubmitHandlerCallerSession) COSTORB() (*big.Int, error) {
	return _SubmitHandler.Contract.COSTORB(&_SubmitHandler.CallOpts)
}

// COSTURB is a free data retrieval call binding the contract method 0x192adc5b.
//
// Solidity: function COST_URB() constant returns(uint256)
func (_SubmitHandler *SubmitHandlerCaller) COSTURB(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SubmitHandler.contract.Call(opts, out, "COST_URB")
	return *ret0, err
}

// COSTURB is a free data retrieval call binding the contract method 0x192adc5b.
//
// Solidity: function COST_URB() constant returns(uint256)
func (_SubmitHandler *SubmitHandlerSession) COSTURB() (*big.Int, error) {
	return _SubmitHandler.Contract.COSTURB(&_SubmitHandler.CallOpts)
}

// COSTURB is a free data retrieval call binding the contract method 0x192adc5b.
//
// Solidity: function COST_URB() constant returns(uint256)
func (_SubmitHandler *SubmitHandlerCallerSession) COSTURB() (*big.Int, error) {
	return _SubmitHandler.Contract.COSTURB(&_SubmitHandler.CallOpts)
}

// COSTURBPREPARE is a free data retrieval call binding the contract method 0x033cfbed.
//
// Solidity: function COST_URB_PREPARE() constant returns(uint256)
func (_SubmitHandler *SubmitHandlerCaller) COSTURBPREPARE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SubmitHandler.contract.Call(opts, out, "COST_URB_PREPARE")
	return *ret0, err
}

// COSTURBPREPARE is a free data retrieval call binding the contract method 0x033cfbed.
//
// Solidity: function COST_URB_PREPARE() constant returns(uint256)
func (_SubmitHandler *SubmitHandlerSession) COSTURBPREPARE() (*big.Int, error) {
	return _SubmitHandler.Contract.COSTURBPREPARE(&_SubmitHandler.CallOpts)
}

// COSTURBPREPARE is a free data retrieval call binding the contract method 0x033cfbed.
//
// Solidity: function COST_URB_PREPARE() constant returns(uint256)
func (_SubmitHandler *SubmitHandlerCallerSession) COSTURBPREPARE() (*big.Int, error) {
	return _SubmitHandler.Contract.COSTURBPREPARE(&_SubmitHandler.CallOpts)
}

// CPCOMPUTATION is a free data retrieval call binding the contract method 0x08c4fff0.
//
// Solidity: function CP_COMPUTATION() constant returns(uint256)
func (_SubmitHandler *SubmitHandlerCaller) CPCOMPUTATION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SubmitHandler.contract.Call(opts, out, "CP_COMPUTATION")
	return *ret0, err
}

// CPCOMPUTATION is a free data retrieval call binding the contract method 0x08c4fff0.
//
// Solidity: function CP_COMPUTATION() constant returns(uint256)
func (_SubmitHandler *SubmitHandlerSession) CPCOMPUTATION() (*big.Int, error) {
	return _SubmitHandler.Contract.CPCOMPUTATION(&_SubmitHandler.CallOpts)
}

// CPCOMPUTATION is a free data retrieval call binding the contract method 0x08c4fff0.
//
// Solidity: function CP_COMPUTATION() constant returns(uint256)
func (_SubmitHandler *SubmitHandlerCallerSession) CPCOMPUTATION() (*big.Int, error) {
	return _SubmitHandler.Contract.CPCOMPUTATION(&_SubmitHandler.CallOpts)
}

// CPEXIT is a free data retrieval call binding the contract method 0x8155717d.
//
// Solidity: function CP_EXIT() constant returns(uint256)
func (_SubmitHandler *SubmitHandlerCaller) CPEXIT(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SubmitHandler.contract.Call(opts, out, "CP_EXIT")
	return *ret0, err
}

// CPEXIT is a free data retrieval call binding the contract method 0x8155717d.
//
// Solidity: function CP_EXIT() constant returns(uint256)
func (_SubmitHandler *SubmitHandlerSession) CPEXIT() (*big.Int, error) {
	return _SubmitHandler.Contract.CPEXIT(&_SubmitHandler.CallOpts)
}

// CPEXIT is a free data retrieval call binding the contract method 0x8155717d.
//
// Solidity: function CP_EXIT() constant returns(uint256)
func (_SubmitHandler *SubmitHandlerCallerSession) CPEXIT() (*big.Int, error) {
	return _SubmitHandler.Contract.CPEXIT(&_SubmitHandler.CallOpts)
}

// CPWITHHOLDING is a free data retrieval call binding the contract method 0xb17fa6e9.
//
// Solidity: function CP_WITHHOLDING() constant returns(uint256)
func (_SubmitHandler *SubmitHandlerCaller) CPWITHHOLDING(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SubmitHandler.contract.Call(opts, out, "CP_WITHHOLDING")
	return *ret0, err
}

// CPWITHHOLDING is a free data retrieval call binding the contract method 0xb17fa6e9.
//
// Solidity: function CP_WITHHOLDING() constant returns(uint256)
func (_SubmitHandler *SubmitHandlerSession) CPWITHHOLDING() (*big.Int, error) {
	return _SubmitHandler.Contract.CPWITHHOLDING(&_SubmitHandler.CallOpts)
}

// CPWITHHOLDING is a free data retrieval call binding the contract method 0xb17fa6e9.
//
// Solidity: function CP_WITHHOLDING() constant returns(uint256)
func (_SubmitHandler *SubmitHandlerCallerSession) CPWITHHOLDING() (*big.Int, error) {
	return _SubmitHandler.Contract.CPWITHHOLDING(&_SubmitHandler.CallOpts)
}

// EROIdToFinalize is a free data retrieval call binding the contract method 0x2dc6bb7b.
//
// Solidity: function EROIdToFinalize() constant returns(uint256)
func (_SubmitHandler *SubmitHandlerCaller) EROIdToFinalize(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SubmitHandler.contract.Call(opts, out, "EROIdToFinalize")
	return *ret0, err
}

// EROIdToFinalize is a free data retrieval call binding the contract method 0x2dc6bb7b.
//
// Solidity: function EROIdToFinalize() constant returns(uint256)
func (_SubmitHandler *SubmitHandlerSession) EROIdToFinalize() (*big.Int, error) {
	return _SubmitHandler.Contract.EROIdToFinalize(&_SubmitHandler.CallOpts)
}

// EROIdToFinalize is a free data retrieval call binding the contract method 0x2dc6bb7b.
//
// Solidity: function EROIdToFinalize() constant returns(uint256)
func (_SubmitHandler *SubmitHandlerCallerSession) EROIdToFinalize() (*big.Int, error) {
	return _SubmitHandler.Contract.EROIdToFinalize(&_SubmitHandler.CallOpts)
}

// EROs is a free data retrieval call binding the contract method 0xb443f3cc.
//
// Solidity: function EROs(uint256 ) constant returns(uint64 timestamp, bool isExit, bool isTransfer, bool finalized, bool challenged, uint128 value, address requestor, address to, bytes32 trieKey, bytes32 hash, bytes trieValue)
func (_SubmitHandler *SubmitHandlerCaller) EROs(opts *bind.CallOpts, arg0 *big.Int) (struct {
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
	err := _SubmitHandler.contract.Call(opts, out, "EROs", arg0)
	return *ret, err
}

// EROs is a free data retrieval call binding the contract method 0xb443f3cc.
//
// Solidity: function EROs(uint256 ) constant returns(uint64 timestamp, bool isExit, bool isTransfer, bool finalized, bool challenged, uint128 value, address requestor, address to, bytes32 trieKey, bytes32 hash, bytes trieValue)
func (_SubmitHandler *SubmitHandlerSession) EROs(arg0 *big.Int) (struct {
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
	return _SubmitHandler.Contract.EROs(&_SubmitHandler.CallOpts, arg0)
}

// EROs is a free data retrieval call binding the contract method 0xb443f3cc.
//
// Solidity: function EROs(uint256 ) constant returns(uint64 timestamp, bool isExit, bool isTransfer, bool finalized, bool challenged, uint128 value, address requestor, address to, bytes32 trieKey, bytes32 hash, bytes trieValue)
func (_SubmitHandler *SubmitHandlerCallerSession) EROs(arg0 *big.Int) (struct {
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
	return _SubmitHandler.Contract.EROs(&_SubmitHandler.CallOpts, arg0)
}

// ERUIdToFinalize is a free data retrieval call binding the contract method 0xc54626cc.
//
// Solidity: function ERUIdToFinalize() constant returns(uint256)
func (_SubmitHandler *SubmitHandlerCaller) ERUIdToFinalize(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SubmitHandler.contract.Call(opts, out, "ERUIdToFinalize")
	return *ret0, err
}

// ERUIdToFinalize is a free data retrieval call binding the contract method 0xc54626cc.
//
// Solidity: function ERUIdToFinalize() constant returns(uint256)
func (_SubmitHandler *SubmitHandlerSession) ERUIdToFinalize() (*big.Int, error) {
	return _SubmitHandler.Contract.ERUIdToFinalize(&_SubmitHandler.CallOpts)
}

// ERUIdToFinalize is a free data retrieval call binding the contract method 0xc54626cc.
//
// Solidity: function ERUIdToFinalize() constant returns(uint256)
func (_SubmitHandler *SubmitHandlerCallerSession) ERUIdToFinalize() (*big.Int, error) {
	return _SubmitHandler.Contract.ERUIdToFinalize(&_SubmitHandler.CallOpts)
}

// ERUs is a free data retrieval call binding the contract method 0xf4f31de4.
//
// Solidity: function ERUs(uint256 ) constant returns(uint64 timestamp, bool isExit, bool isTransfer, bool finalized, bool challenged, uint128 value, address requestor, address to, bytes32 trieKey, bytes32 hash, bytes trieValue)
func (_SubmitHandler *SubmitHandlerCaller) ERUs(opts *bind.CallOpts, arg0 *big.Int) (struct {
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
	err := _SubmitHandler.contract.Call(opts, out, "ERUs", arg0)
	return *ret, err
}

// ERUs is a free data retrieval call binding the contract method 0xf4f31de4.
//
// Solidity: function ERUs(uint256 ) constant returns(uint64 timestamp, bool isExit, bool isTransfer, bool finalized, bool challenged, uint128 value, address requestor, address to, bytes32 trieKey, bytes32 hash, bytes trieValue)
func (_SubmitHandler *SubmitHandlerSession) ERUs(arg0 *big.Int) (struct {
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
	return _SubmitHandler.Contract.ERUs(&_SubmitHandler.CallOpts, arg0)
}

// ERUs is a free data retrieval call binding the contract method 0xf4f31de4.
//
// Solidity: function ERUs(uint256 ) constant returns(uint64 timestamp, bool isExit, bool isTransfer, bool finalized, bool challenged, uint128 value, address requestor, address to, bytes32 trieKey, bytes32 hash, bytes trieValue)
func (_SubmitHandler *SubmitHandlerCallerSession) ERUs(arg0 *big.Int) (struct {
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
	return _SubmitHandler.Contract.ERUs(&_SubmitHandler.CallOpts, arg0)
}

// NRELength is a free data retrieval call binding the contract method 0xab96da2d.
//
// Solidity: function NRELength() constant returns(uint256)
func (_SubmitHandler *SubmitHandlerCaller) NRELength(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SubmitHandler.contract.Call(opts, out, "NRELength")
	return *ret0, err
}

// NRELength is a free data retrieval call binding the contract method 0xab96da2d.
//
// Solidity: function NRELength() constant returns(uint256)
func (_SubmitHandler *SubmitHandlerSession) NRELength() (*big.Int, error) {
	return _SubmitHandler.Contract.NRELength(&_SubmitHandler.CallOpts)
}

// NRELength is a free data retrieval call binding the contract method 0xab96da2d.
//
// Solidity: function NRELength() constant returns(uint256)
func (_SubmitHandler *SubmitHandlerCallerSession) NRELength() (*big.Int, error) {
	return _SubmitHandler.Contract.NRELength(&_SubmitHandler.CallOpts)
}

// NULLADDRESS is a free data retrieval call binding the contract method 0xde0ce17d.
//
// Solidity: function NULL_ADDRESS() constant returns(address)
func (_SubmitHandler *SubmitHandlerCaller) NULLADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SubmitHandler.contract.Call(opts, out, "NULL_ADDRESS")
	return *ret0, err
}

// NULLADDRESS is a free data retrieval call binding the contract method 0xde0ce17d.
//
// Solidity: function NULL_ADDRESS() constant returns(address)
func (_SubmitHandler *SubmitHandlerSession) NULLADDRESS() (common.Address, error) {
	return _SubmitHandler.Contract.NULLADDRESS(&_SubmitHandler.CallOpts)
}

// NULLADDRESS is a free data retrieval call binding the contract method 0xde0ce17d.
//
// Solidity: function NULL_ADDRESS() constant returns(address)
func (_SubmitHandler *SubmitHandlerCallerSession) NULLADDRESS() (common.Address, error) {
	return _SubmitHandler.Contract.NULLADDRESS(&_SubmitHandler.CallOpts)
}

// ORBs is a free data retrieval call binding the contract method 0xea7f22a8.
//
// Solidity: function ORBs(uint256 ) constant returns(bool submitted, uint64 numEnter, uint64 epochNumber, uint64 requestStart, uint64 requestEnd, address trie)
func (_SubmitHandler *SubmitHandlerCaller) ORBs(opts *bind.CallOpts, arg0 *big.Int) (struct {
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
	err := _SubmitHandler.contract.Call(opts, out, "ORBs", arg0)
	return *ret, err
}

// ORBs is a free data retrieval call binding the contract method 0xea7f22a8.
//
// Solidity: function ORBs(uint256 ) constant returns(bool submitted, uint64 numEnter, uint64 epochNumber, uint64 requestStart, uint64 requestEnd, address trie)
func (_SubmitHandler *SubmitHandlerSession) ORBs(arg0 *big.Int) (struct {
	Submitted    bool
	NumEnter     uint64
	EpochNumber  uint64
	RequestStart uint64
	RequestEnd   uint64
	Trie         common.Address
}, error) {
	return _SubmitHandler.Contract.ORBs(&_SubmitHandler.CallOpts, arg0)
}

// ORBs is a free data retrieval call binding the contract method 0xea7f22a8.
//
// Solidity: function ORBs(uint256 ) constant returns(bool submitted, uint64 numEnter, uint64 epochNumber, uint64 requestStart, uint64 requestEnd, address trie)
func (_SubmitHandler *SubmitHandlerCallerSession) ORBs(arg0 *big.Int) (struct {
	Submitted    bool
	NumEnter     uint64
	EpochNumber  uint64
	RequestStart uint64
	RequestEnd   uint64
	Trie         common.Address
}, error) {
	return _SubmitHandler.Contract.ORBs(&_SubmitHandler.CallOpts, arg0)
}

// PREPARETIMEOUT is a free data retrieval call binding the contract method 0xc2bc88fa.
//
// Solidity: function PREPARE_TIMEOUT() constant returns(uint256)
func (_SubmitHandler *SubmitHandlerCaller) PREPARETIMEOUT(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SubmitHandler.contract.Call(opts, out, "PREPARE_TIMEOUT")
	return *ret0, err
}

// PREPARETIMEOUT is a free data retrieval call binding the contract method 0xc2bc88fa.
//
// Solidity: function PREPARE_TIMEOUT() constant returns(uint256)
func (_SubmitHandler *SubmitHandlerSession) PREPARETIMEOUT() (*big.Int, error) {
	return _SubmitHandler.Contract.PREPARETIMEOUT(&_SubmitHandler.CallOpts)
}

// PREPARETIMEOUT is a free data retrieval call binding the contract method 0xc2bc88fa.
//
// Solidity: function PREPARE_TIMEOUT() constant returns(uint256)
func (_SubmitHandler *SubmitHandlerCallerSession) PREPARETIMEOUT() (*big.Int, error) {
	return _SubmitHandler.Contract.PREPARETIMEOUT(&_SubmitHandler.CallOpts)
}

// REQUESTGAS is a free data retrieval call binding the contract method 0x8eb288ca.
//
// Solidity: function REQUEST_GAS() constant returns(uint256)
func (_SubmitHandler *SubmitHandlerCaller) REQUESTGAS(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SubmitHandler.contract.Call(opts, out, "REQUEST_GAS")
	return *ret0, err
}

// REQUESTGAS is a free data retrieval call binding the contract method 0x8eb288ca.
//
// Solidity: function REQUEST_GAS() constant returns(uint256)
func (_SubmitHandler *SubmitHandlerSession) REQUESTGAS() (*big.Int, error) {
	return _SubmitHandler.Contract.REQUESTGAS(&_SubmitHandler.CallOpts)
}

// REQUESTGAS is a free data retrieval call binding the contract method 0x8eb288ca.
//
// Solidity: function REQUEST_GAS() constant returns(uint256)
func (_SubmitHandler *SubmitHandlerCallerSession) REQUESTGAS() (*big.Int, error) {
	return _SubmitHandler.Contract.REQUESTGAS(&_SubmitHandler.CallOpts)
}

// URBs is a free data retrieval call binding the contract method 0xc0e86064.
//
// Solidity: function URBs(uint256 ) constant returns(bool submitted, uint64 numEnter, uint64 epochNumber, uint64 requestStart, uint64 requestEnd, address trie)
func (_SubmitHandler *SubmitHandlerCaller) URBs(opts *bind.CallOpts, arg0 *big.Int) (struct {
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
	err := _SubmitHandler.contract.Call(opts, out, "URBs", arg0)
	return *ret, err
}

// URBs is a free data retrieval call binding the contract method 0xc0e86064.
//
// Solidity: function URBs(uint256 ) constant returns(bool submitted, uint64 numEnter, uint64 epochNumber, uint64 requestStart, uint64 requestEnd, address trie)
func (_SubmitHandler *SubmitHandlerSession) URBs(arg0 *big.Int) (struct {
	Submitted    bool
	NumEnter     uint64
	EpochNumber  uint64
	RequestStart uint64
	RequestEnd   uint64
	Trie         common.Address
}, error) {
	return _SubmitHandler.Contract.URBs(&_SubmitHandler.CallOpts, arg0)
}

// URBs is a free data retrieval call binding the contract method 0xc0e86064.
//
// Solidity: function URBs(uint256 ) constant returns(bool submitted, uint64 numEnter, uint64 epochNumber, uint64 requestStart, uint64 requestEnd, address trie)
func (_SubmitHandler *SubmitHandlerCallerSession) URBs(arg0 *big.Int) (struct {
	Submitted    bool
	NumEnter     uint64
	EpochNumber  uint64
	RequestStart uint64
	RequestEnd   uint64
	Trie         common.Address
}, error) {
	return _SubmitHandler.Contract.URBs(&_SubmitHandler.CallOpts, arg0)
}

// CurrentFork is a free data retrieval call binding the contract method 0x183d2d1c.
//
// Solidity: function currentFork() constant returns(uint256)
func (_SubmitHandler *SubmitHandlerCaller) CurrentFork(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SubmitHandler.contract.Call(opts, out, "currentFork")
	return *ret0, err
}

// CurrentFork is a free data retrieval call binding the contract method 0x183d2d1c.
//
// Solidity: function currentFork() constant returns(uint256)
func (_SubmitHandler *SubmitHandlerSession) CurrentFork() (*big.Int, error) {
	return _SubmitHandler.Contract.CurrentFork(&_SubmitHandler.CallOpts)
}

// CurrentFork is a free data retrieval call binding the contract method 0x183d2d1c.
//
// Solidity: function currentFork() constant returns(uint256)
func (_SubmitHandler *SubmitHandlerCallerSession) CurrentFork() (*big.Int, error) {
	return _SubmitHandler.Contract.CurrentFork(&_SubmitHandler.CallOpts)
}

// Development is a free data retrieval call binding the contract method 0x7b929c27.
//
// Solidity: function development() constant returns(bool)
func (_SubmitHandler *SubmitHandlerCaller) Development(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _SubmitHandler.contract.Call(opts, out, "development")
	return *ret0, err
}

// Development is a free data retrieval call binding the contract method 0x7b929c27.
//
// Solidity: function development() constant returns(bool)
func (_SubmitHandler *SubmitHandlerSession) Development() (bool, error) {
	return _SubmitHandler.Contract.Development(&_SubmitHandler.CallOpts)
}

// Development is a free data retrieval call binding the contract method 0x7b929c27.
//
// Solidity: function development() constant returns(bool)
func (_SubmitHandler *SubmitHandlerCallerSession) Development() (bool, error) {
	return _SubmitHandler.Contract.Development(&_SubmitHandler.CallOpts)
}

// EpochHandler is a free data retrieval call binding the contract method 0xe7b88b80.
//
// Solidity: function epochHandler() constant returns(address)
func (_SubmitHandler *SubmitHandlerCaller) EpochHandler(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SubmitHandler.contract.Call(opts, out, "epochHandler")
	return *ret0, err
}

// EpochHandler is a free data retrieval call binding the contract method 0xe7b88b80.
//
// Solidity: function epochHandler() constant returns(address)
func (_SubmitHandler *SubmitHandlerSession) EpochHandler() (common.Address, error) {
	return _SubmitHandler.Contract.EpochHandler(&_SubmitHandler.CallOpts)
}

// EpochHandler is a free data retrieval call binding the contract method 0xe7b88b80.
//
// Solidity: function epochHandler() constant returns(address)
func (_SubmitHandler *SubmitHandlerCallerSession) EpochHandler() (common.Address, error) {
	return _SubmitHandler.Contract.EpochHandler(&_SubmitHandler.CallOpts)
}

// EtherToken is a free data retrieval call binding the contract method 0xb8066bcb.
//
// Solidity: function etherToken() constant returns(address)
func (_SubmitHandler *SubmitHandlerCaller) EtherToken(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SubmitHandler.contract.Call(opts, out, "etherToken")
	return *ret0, err
}

// EtherToken is a free data retrieval call binding the contract method 0xb8066bcb.
//
// Solidity: function etherToken() constant returns(address)
func (_SubmitHandler *SubmitHandlerSession) EtherToken() (common.Address, error) {
	return _SubmitHandler.Contract.EtherToken(&_SubmitHandler.CallOpts)
}

// EtherToken is a free data retrieval call binding the contract method 0xb8066bcb.
//
// Solidity: function etherToken() constant returns(address)
func (_SubmitHandler *SubmitHandlerCallerSession) EtherToken() (common.Address, error) {
	return _SubmitHandler.Contract.EtherToken(&_SubmitHandler.CallOpts)
}

// FirstFilledORENumber is a free data retrieval call binding the contract method 0x72ecb9a8.
//
// Solidity: function firstFilledORENumber(uint256 ) constant returns(uint256)
func (_SubmitHandler *SubmitHandlerCaller) FirstFilledORENumber(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SubmitHandler.contract.Call(opts, out, "firstFilledORENumber", arg0)
	return *ret0, err
}

// FirstFilledORENumber is a free data retrieval call binding the contract method 0x72ecb9a8.
//
// Solidity: function firstFilledORENumber(uint256 ) constant returns(uint256)
func (_SubmitHandler *SubmitHandlerSession) FirstFilledORENumber(arg0 *big.Int) (*big.Int, error) {
	return _SubmitHandler.Contract.FirstFilledORENumber(&_SubmitHandler.CallOpts, arg0)
}

// FirstFilledORENumber is a free data retrieval call binding the contract method 0x72ecb9a8.
//
// Solidity: function firstFilledORENumber(uint256 ) constant returns(uint256)
func (_SubmitHandler *SubmitHandlerCallerSession) FirstFilledORENumber(arg0 *big.Int) (*big.Int, error) {
	return _SubmitHandler.Contract.FirstFilledORENumber(&_SubmitHandler.CallOpts, arg0)
}

// FirstNonEmptyRequestEpoch is a free data retrieval call binding the contract method 0xca6f6380.
//
// Solidity: function firstNonEmptyRequestEpoch(uint256 ) constant returns(uint256)
func (_SubmitHandler *SubmitHandlerCaller) FirstNonEmptyRequestEpoch(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SubmitHandler.contract.Call(opts, out, "firstNonEmptyRequestEpoch", arg0)
	return *ret0, err
}

// FirstNonEmptyRequestEpoch is a free data retrieval call binding the contract method 0xca6f6380.
//
// Solidity: function firstNonEmptyRequestEpoch(uint256 ) constant returns(uint256)
func (_SubmitHandler *SubmitHandlerSession) FirstNonEmptyRequestEpoch(arg0 *big.Int) (*big.Int, error) {
	return _SubmitHandler.Contract.FirstNonEmptyRequestEpoch(&_SubmitHandler.CallOpts, arg0)
}

// FirstNonEmptyRequestEpoch is a free data retrieval call binding the contract method 0xca6f6380.
//
// Solidity: function firstNonEmptyRequestEpoch(uint256 ) constant returns(uint256)
func (_SubmitHandler *SubmitHandlerCallerSession) FirstNonEmptyRequestEpoch(arg0 *big.Int) (*big.Int, error) {
	return _SubmitHandler.Contract.FirstNonEmptyRequestEpoch(&_SubmitHandler.CallOpts, arg0)
}

// Forks is a free data retrieval call binding the contract method 0x4ba3a126.
//
// Solidity: function forks(uint256 ) constant returns(uint64 forkedBlock, uint64 firstEpoch, uint64 lastEpoch, uint64 firstBlock, uint64 lastBlock, uint64 lastFinalizedEpoch, uint64 lastFinalizedBlock, uint64 timestamp, uint64 firstEnterEpoch, uint64 lastEnterEpoch, uint64 nextBlockToRebase, bool rebased)
func (_SubmitHandler *SubmitHandlerCaller) Forks(opts *bind.CallOpts, arg0 *big.Int) (struct {
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
	err := _SubmitHandler.contract.Call(opts, out, "forks", arg0)
	return *ret, err
}

// Forks is a free data retrieval call binding the contract method 0x4ba3a126.
//
// Solidity: function forks(uint256 ) constant returns(uint64 forkedBlock, uint64 firstEpoch, uint64 lastEpoch, uint64 firstBlock, uint64 lastBlock, uint64 lastFinalizedEpoch, uint64 lastFinalizedBlock, uint64 timestamp, uint64 firstEnterEpoch, uint64 lastEnterEpoch, uint64 nextBlockToRebase, bool rebased)
func (_SubmitHandler *SubmitHandlerSession) Forks(arg0 *big.Int) (struct {
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
	return _SubmitHandler.Contract.Forks(&_SubmitHandler.CallOpts, arg0)
}

// Forks is a free data retrieval call binding the contract method 0x4ba3a126.
//
// Solidity: function forks(uint256 ) constant returns(uint64 forkedBlock, uint64 firstEpoch, uint64 lastEpoch, uint64 firstBlock, uint64 lastBlock, uint64 lastFinalizedEpoch, uint64 lastFinalizedBlock, uint64 timestamp, uint64 firstEnterEpoch, uint64 lastEnterEpoch, uint64 nextBlockToRebase, bool rebased)
func (_SubmitHandler *SubmitHandlerCallerSession) Forks(arg0 *big.Int) (struct {
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
	return _SubmitHandler.Contract.Forks(&_SubmitHandler.CallOpts, arg0)
}

// LastAppliedBlockNumber is a free data retrieval call binding the contract method 0xfb788a27.
//
// Solidity: function lastAppliedBlockNumber() constant returns(uint256)
func (_SubmitHandler *SubmitHandlerCaller) LastAppliedBlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SubmitHandler.contract.Call(opts, out, "lastAppliedBlockNumber")
	return *ret0, err
}

// LastAppliedBlockNumber is a free data retrieval call binding the contract method 0xfb788a27.
//
// Solidity: function lastAppliedBlockNumber() constant returns(uint256)
func (_SubmitHandler *SubmitHandlerSession) LastAppliedBlockNumber() (*big.Int, error) {
	return _SubmitHandler.Contract.LastAppliedBlockNumber(&_SubmitHandler.CallOpts)
}

// LastAppliedBlockNumber is a free data retrieval call binding the contract method 0xfb788a27.
//
// Solidity: function lastAppliedBlockNumber() constant returns(uint256)
func (_SubmitHandler *SubmitHandlerCallerSession) LastAppliedBlockNumber() (*big.Int, error) {
	return _SubmitHandler.Contract.LastAppliedBlockNumber(&_SubmitHandler.CallOpts)
}

// LastAppliedEpochNumber is a free data retrieval call binding the contract method 0xc8ad329f.
//
// Solidity: function lastAppliedEpochNumber() constant returns(uint256)
func (_SubmitHandler *SubmitHandlerCaller) LastAppliedEpochNumber(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SubmitHandler.contract.Call(opts, out, "lastAppliedEpochNumber")
	return *ret0, err
}

// LastAppliedEpochNumber is a free data retrieval call binding the contract method 0xc8ad329f.
//
// Solidity: function lastAppliedEpochNumber() constant returns(uint256)
func (_SubmitHandler *SubmitHandlerSession) LastAppliedEpochNumber() (*big.Int, error) {
	return _SubmitHandler.Contract.LastAppliedEpochNumber(&_SubmitHandler.CallOpts)
}

// LastAppliedEpochNumber is a free data retrieval call binding the contract method 0xc8ad329f.
//
// Solidity: function lastAppliedEpochNumber() constant returns(uint256)
func (_SubmitHandler *SubmitHandlerCallerSession) LastAppliedEpochNumber() (*big.Int, error) {
	return _SubmitHandler.Contract.LastAppliedEpochNumber(&_SubmitHandler.CallOpts)
}

// LastAppliedForkNumber is a free data retrieval call binding the contract method 0x164bc2ae.
//
// Solidity: function lastAppliedForkNumber() constant returns(uint256)
func (_SubmitHandler *SubmitHandlerCaller) LastAppliedForkNumber(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SubmitHandler.contract.Call(opts, out, "lastAppliedForkNumber")
	return *ret0, err
}

// LastAppliedForkNumber is a free data retrieval call binding the contract method 0x164bc2ae.
//
// Solidity: function lastAppliedForkNumber() constant returns(uint256)
func (_SubmitHandler *SubmitHandlerSession) LastAppliedForkNumber() (*big.Int, error) {
	return _SubmitHandler.Contract.LastAppliedForkNumber(&_SubmitHandler.CallOpts)
}

// LastAppliedForkNumber is a free data retrieval call binding the contract method 0x164bc2ae.
//
// Solidity: function lastAppliedForkNumber() constant returns(uint256)
func (_SubmitHandler *SubmitHandlerCallerSession) LastAppliedForkNumber() (*big.Int, error) {
	return _SubmitHandler.Contract.LastAppliedForkNumber(&_SubmitHandler.CallOpts)
}

// LastNonEmptyRequestEpoch is a free data retrieval call binding the contract method 0xb6715647.
//
// Solidity: function lastNonEmptyRequestEpoch(uint256 ) constant returns(uint256)
func (_SubmitHandler *SubmitHandlerCaller) LastNonEmptyRequestEpoch(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SubmitHandler.contract.Call(opts, out, "lastNonEmptyRequestEpoch", arg0)
	return *ret0, err
}

// LastNonEmptyRequestEpoch is a free data retrieval call binding the contract method 0xb6715647.
//
// Solidity: function lastNonEmptyRequestEpoch(uint256 ) constant returns(uint256)
func (_SubmitHandler *SubmitHandlerSession) LastNonEmptyRequestEpoch(arg0 *big.Int) (*big.Int, error) {
	return _SubmitHandler.Contract.LastNonEmptyRequestEpoch(&_SubmitHandler.CallOpts, arg0)
}

// LastNonEmptyRequestEpoch is a free data retrieval call binding the contract method 0xb6715647.
//
// Solidity: function lastNonEmptyRequestEpoch(uint256 ) constant returns(uint256)
func (_SubmitHandler *SubmitHandlerCallerSession) LastNonEmptyRequestEpoch(arg0 *big.Int) (*big.Int, error) {
	return _SubmitHandler.Contract.LastNonEmptyRequestEpoch(&_SubmitHandler.CallOpts, arg0)
}

// NumEnterForORB is a free data retrieval call binding the contract method 0x23691566.
//
// Solidity: function numEnterForORB() constant returns(uint256)
func (_SubmitHandler *SubmitHandlerCaller) NumEnterForORB(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SubmitHandler.contract.Call(opts, out, "numEnterForORB")
	return *ret0, err
}

// NumEnterForORB is a free data retrieval call binding the contract method 0x23691566.
//
// Solidity: function numEnterForORB() constant returns(uint256)
func (_SubmitHandler *SubmitHandlerSession) NumEnterForORB() (*big.Int, error) {
	return _SubmitHandler.Contract.NumEnterForORB(&_SubmitHandler.CallOpts)
}

// NumEnterForORB is a free data retrieval call binding the contract method 0x23691566.
//
// Solidity: function numEnterForORB() constant returns(uint256)
func (_SubmitHandler *SubmitHandlerCallerSession) NumEnterForORB() (*big.Int, error) {
	return _SubmitHandler.Contract.NumEnterForORB(&_SubmitHandler.CallOpts)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() constant returns(address)
func (_SubmitHandler *SubmitHandlerCaller) Operator(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SubmitHandler.contract.Call(opts, out, "operator")
	return *ret0, err
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() constant returns(address)
func (_SubmitHandler *SubmitHandlerSession) Operator() (common.Address, error) {
	return _SubmitHandler.Contract.Operator(&_SubmitHandler.CallOpts)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() constant returns(address)
func (_SubmitHandler *SubmitHandlerCallerSession) Operator() (common.Address, error) {
	return _SubmitHandler.Contract.Operator(&_SubmitHandler.CallOpts)
}

// RequestableContracts is a free data retrieval call binding the contract method 0xda0185f8.
//
// Solidity: function requestableContracts(address ) constant returns(address)
func (_SubmitHandler *SubmitHandlerCaller) RequestableContracts(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SubmitHandler.contract.Call(opts, out, "requestableContracts", arg0)
	return *ret0, err
}

// RequestableContracts is a free data retrieval call binding the contract method 0xda0185f8.
//
// Solidity: function requestableContracts(address ) constant returns(address)
func (_SubmitHandler *SubmitHandlerSession) RequestableContracts(arg0 common.Address) (common.Address, error) {
	return _SubmitHandler.Contract.RequestableContracts(&_SubmitHandler.CallOpts, arg0)
}

// RequestableContracts is a free data retrieval call binding the contract method 0xda0185f8.
//
// Solidity: function requestableContracts(address ) constant returns(address)
func (_SubmitHandler *SubmitHandlerCallerSession) RequestableContracts(arg0 common.Address) (common.Address, error) {
	return _SubmitHandler.Contract.RequestableContracts(&_SubmitHandler.CallOpts, arg0)
}

// SubmitHandler is a free data retrieval call binding the contract method 0xe259faf7.
//
// Solidity: function submitHandler() constant returns(address)
func (_SubmitHandler *SubmitHandlerCaller) SubmitHandler(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SubmitHandler.contract.Call(opts, out, "submitHandler")
	return *ret0, err
}

// SubmitHandler is a free data retrieval call binding the contract method 0xe259faf7.
//
// Solidity: function submitHandler() constant returns(address)
func (_SubmitHandler *SubmitHandlerSession) SubmitHandler() (common.Address, error) {
	return _SubmitHandler.Contract.SubmitHandler(&_SubmitHandler.CallOpts)
}

// SubmitHandler is a free data retrieval call binding the contract method 0xe259faf7.
//
// Solidity: function submitHandler() constant returns(address)
func (_SubmitHandler *SubmitHandlerCallerSession) SubmitHandler() (common.Address, error) {
	return _SubmitHandler.Contract.SubmitHandler(&_SubmitHandler.CallOpts)
}

// SubmitNRE is a paid mutator transaction binding the contract method 0x0eaf45a8.
//
// Solidity: function submitNRE(uint256 _pos1, uint256 _pos2, bytes32 _epochStateRoot, bytes32 _epochTransactionsRoot, bytes32 _epochReceiptsRoot) returns(bool success)
func (_SubmitHandler *SubmitHandlerTransactor) SubmitNRE(opts *bind.TransactOpts, _pos1 *big.Int, _pos2 *big.Int, _epochStateRoot [32]byte, _epochTransactionsRoot [32]byte, _epochReceiptsRoot [32]byte) (*types.Transaction, error) {
	return _SubmitHandler.contract.Transact(opts, "submitNRE", _pos1, _pos2, _epochStateRoot, _epochTransactionsRoot, _epochReceiptsRoot)
}

// SubmitNRE is a paid mutator transaction binding the contract method 0x0eaf45a8.
//
// Solidity: function submitNRE(uint256 _pos1, uint256 _pos2, bytes32 _epochStateRoot, bytes32 _epochTransactionsRoot, bytes32 _epochReceiptsRoot) returns(bool success)
func (_SubmitHandler *SubmitHandlerSession) SubmitNRE(_pos1 *big.Int, _pos2 *big.Int, _epochStateRoot [32]byte, _epochTransactionsRoot [32]byte, _epochReceiptsRoot [32]byte) (*types.Transaction, error) {
	return _SubmitHandler.Contract.SubmitNRE(&_SubmitHandler.TransactOpts, _pos1, _pos2, _epochStateRoot, _epochTransactionsRoot, _epochReceiptsRoot)
}

// SubmitNRE is a paid mutator transaction binding the contract method 0x0eaf45a8.
//
// Solidity: function submitNRE(uint256 _pos1, uint256 _pos2, bytes32 _epochStateRoot, bytes32 _epochTransactionsRoot, bytes32 _epochReceiptsRoot) returns(bool success)
func (_SubmitHandler *SubmitHandlerTransactorSession) SubmitNRE(_pos1 *big.Int, _pos2 *big.Int, _epochStateRoot [32]byte, _epochTransactionsRoot [32]byte, _epochReceiptsRoot [32]byte) (*types.Transaction, error) {
	return _SubmitHandler.Contract.SubmitNRE(&_SubmitHandler.TransactOpts, _pos1, _pos2, _epochStateRoot, _epochTransactionsRoot, _epochReceiptsRoot)
}

// SubmitORB is a paid mutator transaction binding the contract method 0xa820c067.
//
// Solidity: function submitORB(uint256 _pos, bytes32 _statesRoot, bytes32 _transactionsRoot, bytes32 _receiptsRoot) returns(bool success)
func (_SubmitHandler *SubmitHandlerTransactor) SubmitORB(opts *bind.TransactOpts, _pos *big.Int, _statesRoot [32]byte, _transactionsRoot [32]byte, _receiptsRoot [32]byte) (*types.Transaction, error) {
	return _SubmitHandler.contract.Transact(opts, "submitORB", _pos, _statesRoot, _transactionsRoot, _receiptsRoot)
}

// SubmitORB is a paid mutator transaction binding the contract method 0xa820c067.
//
// Solidity: function submitORB(uint256 _pos, bytes32 _statesRoot, bytes32 _transactionsRoot, bytes32 _receiptsRoot) returns(bool success)
func (_SubmitHandler *SubmitHandlerSession) SubmitORB(_pos *big.Int, _statesRoot [32]byte, _transactionsRoot [32]byte, _receiptsRoot [32]byte) (*types.Transaction, error) {
	return _SubmitHandler.Contract.SubmitORB(&_SubmitHandler.TransactOpts, _pos, _statesRoot, _transactionsRoot, _receiptsRoot)
}

// SubmitORB is a paid mutator transaction binding the contract method 0xa820c067.
//
// Solidity: function submitORB(uint256 _pos, bytes32 _statesRoot, bytes32 _transactionsRoot, bytes32 _receiptsRoot) returns(bool success)
func (_SubmitHandler *SubmitHandlerTransactorSession) SubmitORB(_pos *big.Int, _statesRoot [32]byte, _transactionsRoot [32]byte, _receiptsRoot [32]byte) (*types.Transaction, error) {
	return _SubmitHandler.Contract.SubmitORB(&_SubmitHandler.TransactOpts, _pos, _statesRoot, _transactionsRoot, _receiptsRoot)
}

// SubmitURB is a paid mutator transaction binding the contract method 0x6f3e4b90.
//
// Solidity: function submitURB(uint256 _pos, bytes32 _statesRoot, bytes32 _transactionsRoot, bytes32 _receiptsRoot) returns(bool success)
func (_SubmitHandler *SubmitHandlerTransactor) SubmitURB(opts *bind.TransactOpts, _pos *big.Int, _statesRoot [32]byte, _transactionsRoot [32]byte, _receiptsRoot [32]byte) (*types.Transaction, error) {
	return _SubmitHandler.contract.Transact(opts, "submitURB", _pos, _statesRoot, _transactionsRoot, _receiptsRoot)
}

// SubmitURB is a paid mutator transaction binding the contract method 0x6f3e4b90.
//
// Solidity: function submitURB(uint256 _pos, bytes32 _statesRoot, bytes32 _transactionsRoot, bytes32 _receiptsRoot) returns(bool success)
func (_SubmitHandler *SubmitHandlerSession) SubmitURB(_pos *big.Int, _statesRoot [32]byte, _transactionsRoot [32]byte, _receiptsRoot [32]byte) (*types.Transaction, error) {
	return _SubmitHandler.Contract.SubmitURB(&_SubmitHandler.TransactOpts, _pos, _statesRoot, _transactionsRoot, _receiptsRoot)
}

// SubmitURB is a paid mutator transaction binding the contract method 0x6f3e4b90.
//
// Solidity: function submitURB(uint256 _pos, bytes32 _statesRoot, bytes32 _transactionsRoot, bytes32 _receiptsRoot) returns(bool success)
func (_SubmitHandler *SubmitHandlerTransactorSession) SubmitURB(_pos *big.Int, _statesRoot [32]byte, _transactionsRoot [32]byte, _receiptsRoot [32]byte) (*types.Transaction, error) {
	return _SubmitHandler.Contract.SubmitURB(&_SubmitHandler.TransactOpts, _pos, _statesRoot, _transactionsRoot, _receiptsRoot)
}

// SubmitHandlerBlockFinalizedIterator is returned from FilterBlockFinalized and is used to iterate over the raw logs and unpacked data for BlockFinalized events raised by the SubmitHandler contract.
type SubmitHandlerBlockFinalizedIterator struct {
	Event *SubmitHandlerBlockFinalized // Event containing the contract specifics and raw log

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
func (it *SubmitHandlerBlockFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubmitHandlerBlockFinalized)
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
		it.Event = new(SubmitHandlerBlockFinalized)
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
func (it *SubmitHandlerBlockFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubmitHandlerBlockFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubmitHandlerBlockFinalized represents a BlockFinalized event raised by the SubmitHandler contract.
type SubmitHandlerBlockFinalized struct {
	ForkNumber  *big.Int
	BlockNumber *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBlockFinalized is a free log retrieval operation binding the contract event 0xfb96205e4b3633fd57aa805b26b51ecf528714a10241a4af015929dce86768d9.
//
// Solidity: event BlockFinalized(uint256 forkNumber, uint256 blockNumber)
func (_SubmitHandler *SubmitHandlerFilterer) FilterBlockFinalized(opts *bind.FilterOpts) (*SubmitHandlerBlockFinalizedIterator, error) {

	logs, sub, err := _SubmitHandler.contract.FilterLogs(opts, "BlockFinalized")
	if err != nil {
		return nil, err
	}
	return &SubmitHandlerBlockFinalizedIterator{contract: _SubmitHandler.contract, event: "BlockFinalized", logs: logs, sub: sub}, nil
}

// WatchBlockFinalized is a free log subscription operation binding the contract event 0xfb96205e4b3633fd57aa805b26b51ecf528714a10241a4af015929dce86768d9.
//
// Solidity: event BlockFinalized(uint256 forkNumber, uint256 blockNumber)
func (_SubmitHandler *SubmitHandlerFilterer) WatchBlockFinalized(opts *bind.WatchOpts, sink chan<- *SubmitHandlerBlockFinalized) (event.Subscription, error) {

	logs, sub, err := _SubmitHandler.contract.WatchLogs(opts, "BlockFinalized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubmitHandlerBlockFinalized)
				if err := _SubmitHandler.contract.UnpackLog(event, "BlockFinalized", log); err != nil {
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
func (_SubmitHandler *SubmitHandlerFilterer) ParseBlockFinalized(log types.Log) (*SubmitHandlerBlockFinalized, error) {
	event := new(SubmitHandlerBlockFinalized)
	if err := _SubmitHandler.contract.UnpackLog(event, "BlockFinalized", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SubmitHandlerBlockSubmittedIterator is returned from FilterBlockSubmitted and is used to iterate over the raw logs and unpacked data for BlockSubmitted events raised by the SubmitHandler contract.
type SubmitHandlerBlockSubmittedIterator struct {
	Event *SubmitHandlerBlockSubmitted // Event containing the contract specifics and raw log

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
func (it *SubmitHandlerBlockSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubmitHandlerBlockSubmitted)
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
		it.Event = new(SubmitHandlerBlockSubmitted)
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
func (it *SubmitHandlerBlockSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubmitHandlerBlockSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubmitHandlerBlockSubmitted represents a BlockSubmitted event raised by the SubmitHandler contract.
type SubmitHandlerBlockSubmitted struct {
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
func (_SubmitHandler *SubmitHandlerFilterer) FilterBlockSubmitted(opts *bind.FilterOpts) (*SubmitHandlerBlockSubmittedIterator, error) {

	logs, sub, err := _SubmitHandler.contract.FilterLogs(opts, "BlockSubmitted")
	if err != nil {
		return nil, err
	}
	return &SubmitHandlerBlockSubmittedIterator{contract: _SubmitHandler.contract, event: "BlockSubmitted", logs: logs, sub: sub}, nil
}

// WatchBlockSubmitted is a free log subscription operation binding the contract event 0x3d4a04291c66b06f39a4ecb817875b12b5485a05ec563133a56a905305c48e55.
//
// Solidity: event BlockSubmitted(uint256 fork, uint256 epochNumber, uint256 blockNumber, bool isRequest, bool userActivated)
func (_SubmitHandler *SubmitHandlerFilterer) WatchBlockSubmitted(opts *bind.WatchOpts, sink chan<- *SubmitHandlerBlockSubmitted) (event.Subscription, error) {

	logs, sub, err := _SubmitHandler.contract.WatchLogs(opts, "BlockSubmitted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubmitHandlerBlockSubmitted)
				if err := _SubmitHandler.contract.UnpackLog(event, "BlockSubmitted", log); err != nil {
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
func (_SubmitHandler *SubmitHandlerFilterer) ParseBlockSubmitted(log types.Log) (*SubmitHandlerBlockSubmitted, error) {
	event := new(SubmitHandlerBlockSubmitted)
	if err := _SubmitHandler.contract.UnpackLog(event, "BlockSubmitted", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SubmitHandlerERUCreatedIterator is returned from FilterERUCreated and is used to iterate over the raw logs and unpacked data for ERUCreated events raised by the SubmitHandler contract.
type SubmitHandlerERUCreatedIterator struct {
	Event *SubmitHandlerERUCreated // Event containing the contract specifics and raw log

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
func (it *SubmitHandlerERUCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubmitHandlerERUCreated)
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
		it.Event = new(SubmitHandlerERUCreated)
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
func (it *SubmitHandlerERUCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubmitHandlerERUCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubmitHandlerERUCreated represents a ERUCreated event raised by the SubmitHandler contract.
type SubmitHandlerERUCreated struct {
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
func (_SubmitHandler *SubmitHandlerFilterer) FilterERUCreated(opts *bind.FilterOpts) (*SubmitHandlerERUCreatedIterator, error) {

	logs, sub, err := _SubmitHandler.contract.FilterLogs(opts, "ERUCreated")
	if err != nil {
		return nil, err
	}
	return &SubmitHandlerERUCreatedIterator{contract: _SubmitHandler.contract, event: "ERUCreated", logs: logs, sub: sub}, nil
}

// WatchERUCreated is a free log subscription operation binding the contract event 0xfcbdc2083dadd644b854d91b49aef8db06b8f5a3d5c1192de38ca0ba271d5a0d.
//
// Solidity: event ERUCreated(uint256 requestId, address requestor, address to, bytes trieKey, bytes32 trieValue)
func (_SubmitHandler *SubmitHandlerFilterer) WatchERUCreated(opts *bind.WatchOpts, sink chan<- *SubmitHandlerERUCreated) (event.Subscription, error) {

	logs, sub, err := _SubmitHandler.contract.WatchLogs(opts, "ERUCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubmitHandlerERUCreated)
				if err := _SubmitHandler.contract.UnpackLog(event, "ERUCreated", log); err != nil {
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
func (_SubmitHandler *SubmitHandlerFilterer) ParseERUCreated(log types.Log) (*SubmitHandlerERUCreated, error) {
	event := new(SubmitHandlerERUCreated)
	if err := _SubmitHandler.contract.UnpackLog(event, "ERUCreated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SubmitHandlerEpochFilledIterator is returned from FilterEpochFilled and is used to iterate over the raw logs and unpacked data for EpochFilled events raised by the SubmitHandler contract.
type SubmitHandlerEpochFilledIterator struct {
	Event *SubmitHandlerEpochFilled // Event containing the contract specifics and raw log

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
func (it *SubmitHandlerEpochFilledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubmitHandlerEpochFilled)
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
		it.Event = new(SubmitHandlerEpochFilled)
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
func (it *SubmitHandlerEpochFilledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubmitHandlerEpochFilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubmitHandlerEpochFilled represents a EpochFilled event raised by the SubmitHandler contract.
type SubmitHandlerEpochFilled struct {
	ForkNumber  *big.Int
	EpochNumber *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterEpochFilled is a free log retrieval operation binding the contract event 0x2fdeb407bf5c2b621f04b5c784822dae806c45b49a68aba413cc270128c96816.
//
// Solidity: event EpochFilled(uint256 forkNumber, uint256 epochNumber)
func (_SubmitHandler *SubmitHandlerFilterer) FilterEpochFilled(opts *bind.FilterOpts) (*SubmitHandlerEpochFilledIterator, error) {

	logs, sub, err := _SubmitHandler.contract.FilterLogs(opts, "EpochFilled")
	if err != nil {
		return nil, err
	}
	return &SubmitHandlerEpochFilledIterator{contract: _SubmitHandler.contract, event: "EpochFilled", logs: logs, sub: sub}, nil
}

// WatchEpochFilled is a free log subscription operation binding the contract event 0x2fdeb407bf5c2b621f04b5c784822dae806c45b49a68aba413cc270128c96816.
//
// Solidity: event EpochFilled(uint256 forkNumber, uint256 epochNumber)
func (_SubmitHandler *SubmitHandlerFilterer) WatchEpochFilled(opts *bind.WatchOpts, sink chan<- *SubmitHandlerEpochFilled) (event.Subscription, error) {

	logs, sub, err := _SubmitHandler.contract.WatchLogs(opts, "EpochFilled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubmitHandlerEpochFilled)
				if err := _SubmitHandler.contract.UnpackLog(event, "EpochFilled", log); err != nil {
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
func (_SubmitHandler *SubmitHandlerFilterer) ParseEpochFilled(log types.Log) (*SubmitHandlerEpochFilled, error) {
	event := new(SubmitHandlerEpochFilled)
	if err := _SubmitHandler.contract.UnpackLog(event, "EpochFilled", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SubmitHandlerEpochFillingIterator is returned from FilterEpochFilling and is used to iterate over the raw logs and unpacked data for EpochFilling events raised by the SubmitHandler contract.
type SubmitHandlerEpochFillingIterator struct {
	Event *SubmitHandlerEpochFilling // Event containing the contract specifics and raw log

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
func (it *SubmitHandlerEpochFillingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubmitHandlerEpochFilling)
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
		it.Event = new(SubmitHandlerEpochFilling)
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
func (it *SubmitHandlerEpochFillingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubmitHandlerEpochFillingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubmitHandlerEpochFilling represents a EpochFilling event raised by the SubmitHandler contract.
type SubmitHandlerEpochFilling struct {
	ForkNumber  *big.Int
	EpochNumber *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterEpochFilling is a free log retrieval operation binding the contract event 0x27b09f0953d27bbff306fe25b2987ac5a813248ac30cb2bbd5daf95e7b0e6dc0.
//
// Solidity: event EpochFilling(uint256 forkNumber, uint256 epochNumber)
func (_SubmitHandler *SubmitHandlerFilterer) FilterEpochFilling(opts *bind.FilterOpts) (*SubmitHandlerEpochFillingIterator, error) {

	logs, sub, err := _SubmitHandler.contract.FilterLogs(opts, "EpochFilling")
	if err != nil {
		return nil, err
	}
	return &SubmitHandlerEpochFillingIterator{contract: _SubmitHandler.contract, event: "EpochFilling", logs: logs, sub: sub}, nil
}

// WatchEpochFilling is a free log subscription operation binding the contract event 0x27b09f0953d27bbff306fe25b2987ac5a813248ac30cb2bbd5daf95e7b0e6dc0.
//
// Solidity: event EpochFilling(uint256 forkNumber, uint256 epochNumber)
func (_SubmitHandler *SubmitHandlerFilterer) WatchEpochFilling(opts *bind.WatchOpts, sink chan<- *SubmitHandlerEpochFilling) (event.Subscription, error) {

	logs, sub, err := _SubmitHandler.contract.WatchLogs(opts, "EpochFilling")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubmitHandlerEpochFilling)
				if err := _SubmitHandler.contract.UnpackLog(event, "EpochFilling", log); err != nil {
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
func (_SubmitHandler *SubmitHandlerFilterer) ParseEpochFilling(log types.Log) (*SubmitHandlerEpochFilling, error) {
	event := new(SubmitHandlerEpochFilling)
	if err := _SubmitHandler.contract.UnpackLog(event, "EpochFilling", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SubmitHandlerEpochFinalizedIterator is returned from FilterEpochFinalized and is used to iterate over the raw logs and unpacked data for EpochFinalized events raised by the SubmitHandler contract.
type SubmitHandlerEpochFinalizedIterator struct {
	Event *SubmitHandlerEpochFinalized // Event containing the contract specifics and raw log

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
func (it *SubmitHandlerEpochFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubmitHandlerEpochFinalized)
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
		it.Event = new(SubmitHandlerEpochFinalized)
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
func (it *SubmitHandlerEpochFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubmitHandlerEpochFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubmitHandlerEpochFinalized represents a EpochFinalized event raised by the SubmitHandler contract.
type SubmitHandlerEpochFinalized struct {
	ForkNumber       *big.Int
	EpochNumber      *big.Int
	StartBlockNumber *big.Int
	EndBlockNumber   *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterEpochFinalized is a free log retrieval operation binding the contract event 0x70801d4d63b3da6c19ba7349911f45bed5a99ccdfb51b8138c105872529bebd5.
//
// Solidity: event EpochFinalized(uint256 forkNumber, uint256 epochNumber, uint256 startBlockNumber, uint256 endBlockNumber)
func (_SubmitHandler *SubmitHandlerFilterer) FilterEpochFinalized(opts *bind.FilterOpts) (*SubmitHandlerEpochFinalizedIterator, error) {

	logs, sub, err := _SubmitHandler.contract.FilterLogs(opts, "EpochFinalized")
	if err != nil {
		return nil, err
	}
	return &SubmitHandlerEpochFinalizedIterator{contract: _SubmitHandler.contract, event: "EpochFinalized", logs: logs, sub: sub}, nil
}

// WatchEpochFinalized is a free log subscription operation binding the contract event 0x70801d4d63b3da6c19ba7349911f45bed5a99ccdfb51b8138c105872529bebd5.
//
// Solidity: event EpochFinalized(uint256 forkNumber, uint256 epochNumber, uint256 startBlockNumber, uint256 endBlockNumber)
func (_SubmitHandler *SubmitHandlerFilterer) WatchEpochFinalized(opts *bind.WatchOpts, sink chan<- *SubmitHandlerEpochFinalized) (event.Subscription, error) {

	logs, sub, err := _SubmitHandler.contract.WatchLogs(opts, "EpochFinalized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubmitHandlerEpochFinalized)
				if err := _SubmitHandler.contract.UnpackLog(event, "EpochFinalized", log); err != nil {
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
func (_SubmitHandler *SubmitHandlerFilterer) ParseEpochFinalized(log types.Log) (*SubmitHandlerEpochFinalized, error) {
	event := new(SubmitHandlerEpochFinalized)
	if err := _SubmitHandler.contract.UnpackLog(event, "EpochFinalized", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SubmitHandlerEpochPreparedIterator is returned from FilterEpochPrepared and is used to iterate over the raw logs and unpacked data for EpochPrepared events raised by the SubmitHandler contract.
type SubmitHandlerEpochPreparedIterator struct {
	Event *SubmitHandlerEpochPrepared // Event containing the contract specifics and raw log

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
func (it *SubmitHandlerEpochPreparedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubmitHandlerEpochPrepared)
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
		it.Event = new(SubmitHandlerEpochPrepared)
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
func (it *SubmitHandlerEpochPreparedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubmitHandlerEpochPreparedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubmitHandlerEpochPrepared represents a EpochPrepared event raised by the SubmitHandler contract.
type SubmitHandlerEpochPrepared struct {
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
func (_SubmitHandler *SubmitHandlerFilterer) FilterEpochPrepared(opts *bind.FilterOpts) (*SubmitHandlerEpochPreparedIterator, error) {

	logs, sub, err := _SubmitHandler.contract.FilterLogs(opts, "EpochPrepared")
	if err != nil {
		return nil, err
	}
	return &SubmitHandlerEpochPreparedIterator{contract: _SubmitHandler.contract, event: "EpochPrepared", logs: logs, sub: sub}, nil
}

// WatchEpochPrepared is a free log subscription operation binding the contract event 0x1a69c0760aa329b76f72579129869013ebd3d41594db019c0e997b939fcb32e3.
//
// Solidity: event EpochPrepared(uint256 forkNumber, uint256 epochNumber, uint256 startBlockNumber, uint256 endBlockNumber, uint256 requestStart, uint256 requestEnd, bool epochIsEmpty, bool isRequest, bool userActivated, bool rebase)
func (_SubmitHandler *SubmitHandlerFilterer) WatchEpochPrepared(opts *bind.WatchOpts, sink chan<- *SubmitHandlerEpochPrepared) (event.Subscription, error) {

	logs, sub, err := _SubmitHandler.contract.WatchLogs(opts, "EpochPrepared")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubmitHandlerEpochPrepared)
				if err := _SubmitHandler.contract.UnpackLog(event, "EpochPrepared", log); err != nil {
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
func (_SubmitHandler *SubmitHandlerFilterer) ParseEpochPrepared(log types.Log) (*SubmitHandlerEpochPrepared, error) {
	event := new(SubmitHandlerEpochPrepared)
	if err := _SubmitHandler.contract.UnpackLog(event, "EpochPrepared", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SubmitHandlerEpochRebasedIterator is returned from FilterEpochRebased and is used to iterate over the raw logs and unpacked data for EpochRebased events raised by the SubmitHandler contract.
type SubmitHandlerEpochRebasedIterator struct {
	Event *SubmitHandlerEpochRebased // Event containing the contract specifics and raw log

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
func (it *SubmitHandlerEpochRebasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubmitHandlerEpochRebased)
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
		it.Event = new(SubmitHandlerEpochRebased)
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
func (it *SubmitHandlerEpochRebasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubmitHandlerEpochRebasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubmitHandlerEpochRebased represents a EpochRebased event raised by the SubmitHandler contract.
type SubmitHandlerEpochRebased struct {
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
func (_SubmitHandler *SubmitHandlerFilterer) FilterEpochRebased(opts *bind.FilterOpts) (*SubmitHandlerEpochRebasedIterator, error) {

	logs, sub, err := _SubmitHandler.contract.FilterLogs(opts, "EpochRebased")
	if err != nil {
		return nil, err
	}
	return &SubmitHandlerEpochRebasedIterator{contract: _SubmitHandler.contract, event: "EpochRebased", logs: logs, sub: sub}, nil
}

// WatchEpochRebased is a free log subscription operation binding the contract event 0x030c1c69405c93021f28f57557240dee939a320b826a1fd0d39bf6e629ecab47.
//
// Solidity: event EpochRebased(uint256 forkNumber, uint256 epochNumber, uint256 startBlockNumber, uint256 endBlockNumber, uint256 requestStart, uint256 requestEnd, bool epochIsEmpty, bool isRequest, bool userActivated)
func (_SubmitHandler *SubmitHandlerFilterer) WatchEpochRebased(opts *bind.WatchOpts, sink chan<- *SubmitHandlerEpochRebased) (event.Subscription, error) {

	logs, sub, err := _SubmitHandler.contract.WatchLogs(opts, "EpochRebased")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubmitHandlerEpochRebased)
				if err := _SubmitHandler.contract.UnpackLog(event, "EpochRebased", log); err != nil {
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
func (_SubmitHandler *SubmitHandlerFilterer) ParseEpochRebased(log types.Log) (*SubmitHandlerEpochRebased, error) {
	event := new(SubmitHandlerEpochRebased)
	if err := _SubmitHandler.contract.UnpackLog(event, "EpochRebased", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SubmitHandlerForkedIterator is returned from FilterForked and is used to iterate over the raw logs and unpacked data for Forked events raised by the SubmitHandler contract.
type SubmitHandlerForkedIterator struct {
	Event *SubmitHandlerForked // Event containing the contract specifics and raw log

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
func (it *SubmitHandlerForkedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubmitHandlerForked)
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
		it.Event = new(SubmitHandlerForked)
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
func (it *SubmitHandlerForkedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubmitHandlerForkedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubmitHandlerForked represents a Forked event raised by the SubmitHandler contract.
type SubmitHandlerForked struct {
	NewFork           *big.Int
	EpochNumber       *big.Int
	ForkedBlockNumber *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterForked is a free log retrieval operation binding the contract event 0x0647d42ab02f6e0ae76959757dcb6aa6feac1d4ba6f077f1223fb4b1b429f06c.
//
// Solidity: event Forked(uint256 newFork, uint256 epochNumber, uint256 forkedBlockNumber)
func (_SubmitHandler *SubmitHandlerFilterer) FilterForked(opts *bind.FilterOpts) (*SubmitHandlerForkedIterator, error) {

	logs, sub, err := _SubmitHandler.contract.FilterLogs(opts, "Forked")
	if err != nil {
		return nil, err
	}
	return &SubmitHandlerForkedIterator{contract: _SubmitHandler.contract, event: "Forked", logs: logs, sub: sub}, nil
}

// WatchForked is a free log subscription operation binding the contract event 0x0647d42ab02f6e0ae76959757dcb6aa6feac1d4ba6f077f1223fb4b1b429f06c.
//
// Solidity: event Forked(uint256 newFork, uint256 epochNumber, uint256 forkedBlockNumber)
func (_SubmitHandler *SubmitHandlerFilterer) WatchForked(opts *bind.WatchOpts, sink chan<- *SubmitHandlerForked) (event.Subscription, error) {

	logs, sub, err := _SubmitHandler.contract.WatchLogs(opts, "Forked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubmitHandlerForked)
				if err := _SubmitHandler.contract.UnpackLog(event, "Forked", log); err != nil {
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
func (_SubmitHandler *SubmitHandlerFilterer) ParseForked(log types.Log) (*SubmitHandlerForked, error) {
	event := new(SubmitHandlerForked)
	if err := _SubmitHandler.contract.UnpackLog(event, "Forked", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SubmitHandlerRequestAppliedIterator is returned from FilterRequestApplied and is used to iterate over the raw logs and unpacked data for RequestApplied events raised by the SubmitHandler contract.
type SubmitHandlerRequestAppliedIterator struct {
	Event *SubmitHandlerRequestApplied // Event containing the contract specifics and raw log

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
func (it *SubmitHandlerRequestAppliedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubmitHandlerRequestApplied)
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
		it.Event = new(SubmitHandlerRequestApplied)
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
func (it *SubmitHandlerRequestAppliedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubmitHandlerRequestAppliedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubmitHandlerRequestApplied represents a RequestApplied event raised by the SubmitHandler contract.
type SubmitHandlerRequestApplied struct {
	RequestId     *big.Int
	UserActivated bool
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRequestApplied is a free log retrieval operation binding the contract event 0x6940a01870e576ceb735867e13863646d517ce10e66c0133186a4ebdfe9388c2.
//
// Solidity: event RequestApplied(uint256 requestId, bool userActivated)
func (_SubmitHandler *SubmitHandlerFilterer) FilterRequestApplied(opts *bind.FilterOpts) (*SubmitHandlerRequestAppliedIterator, error) {

	logs, sub, err := _SubmitHandler.contract.FilterLogs(opts, "RequestApplied")
	if err != nil {
		return nil, err
	}
	return &SubmitHandlerRequestAppliedIterator{contract: _SubmitHandler.contract, event: "RequestApplied", logs: logs, sub: sub}, nil
}

// WatchRequestApplied is a free log subscription operation binding the contract event 0x6940a01870e576ceb735867e13863646d517ce10e66c0133186a4ebdfe9388c2.
//
// Solidity: event RequestApplied(uint256 requestId, bool userActivated)
func (_SubmitHandler *SubmitHandlerFilterer) WatchRequestApplied(opts *bind.WatchOpts, sink chan<- *SubmitHandlerRequestApplied) (event.Subscription, error) {

	logs, sub, err := _SubmitHandler.contract.WatchLogs(opts, "RequestApplied")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubmitHandlerRequestApplied)
				if err := _SubmitHandler.contract.UnpackLog(event, "RequestApplied", log); err != nil {
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
func (_SubmitHandler *SubmitHandlerFilterer) ParseRequestApplied(log types.Log) (*SubmitHandlerRequestApplied, error) {
	event := new(SubmitHandlerRequestApplied)
	if err := _SubmitHandler.contract.UnpackLog(event, "RequestApplied", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SubmitHandlerRequestChallengedIterator is returned from FilterRequestChallenged and is used to iterate over the raw logs and unpacked data for RequestChallenged events raised by the SubmitHandler contract.
type SubmitHandlerRequestChallengedIterator struct {
	Event *SubmitHandlerRequestChallenged // Event containing the contract specifics and raw log

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
func (it *SubmitHandlerRequestChallengedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubmitHandlerRequestChallenged)
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
		it.Event = new(SubmitHandlerRequestChallenged)
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
func (it *SubmitHandlerRequestChallengedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubmitHandlerRequestChallengedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubmitHandlerRequestChallenged represents a RequestChallenged event raised by the SubmitHandler contract.
type SubmitHandlerRequestChallenged struct {
	RequestId     *big.Int
	UserActivated bool
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRequestChallenged is a free log retrieval operation binding the contract event 0xc8135db115644ed4ae193313c4c801235ef740d2a57a8d5e6fe26ab66635698a.
//
// Solidity: event RequestChallenged(uint256 requestId, bool userActivated)
func (_SubmitHandler *SubmitHandlerFilterer) FilterRequestChallenged(opts *bind.FilterOpts) (*SubmitHandlerRequestChallengedIterator, error) {

	logs, sub, err := _SubmitHandler.contract.FilterLogs(opts, "RequestChallenged")
	if err != nil {
		return nil, err
	}
	return &SubmitHandlerRequestChallengedIterator{contract: _SubmitHandler.contract, event: "RequestChallenged", logs: logs, sub: sub}, nil
}

// WatchRequestChallenged is a free log subscription operation binding the contract event 0xc8135db115644ed4ae193313c4c801235ef740d2a57a8d5e6fe26ab66635698a.
//
// Solidity: event RequestChallenged(uint256 requestId, bool userActivated)
func (_SubmitHandler *SubmitHandlerFilterer) WatchRequestChallenged(opts *bind.WatchOpts, sink chan<- *SubmitHandlerRequestChallenged) (event.Subscription, error) {

	logs, sub, err := _SubmitHandler.contract.WatchLogs(opts, "RequestChallenged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubmitHandlerRequestChallenged)
				if err := _SubmitHandler.contract.UnpackLog(event, "RequestChallenged", log); err != nil {
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
func (_SubmitHandler *SubmitHandlerFilterer) ParseRequestChallenged(log types.Log) (*SubmitHandlerRequestChallenged, error) {
	event := new(SubmitHandlerRequestChallenged)
	if err := _SubmitHandler.contract.UnpackLog(event, "RequestChallenged", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SubmitHandlerRequestCreatedIterator is returned from FilterRequestCreated and is used to iterate over the raw logs and unpacked data for RequestCreated events raised by the SubmitHandler contract.
type SubmitHandlerRequestCreatedIterator struct {
	Event *SubmitHandlerRequestCreated // Event containing the contract specifics and raw log

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
func (it *SubmitHandlerRequestCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubmitHandlerRequestCreated)
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
		it.Event = new(SubmitHandlerRequestCreated)
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
func (it *SubmitHandlerRequestCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubmitHandlerRequestCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubmitHandlerRequestCreated represents a RequestCreated event raised by the SubmitHandler contract.
type SubmitHandlerRequestCreated struct {
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
func (_SubmitHandler *SubmitHandlerFilterer) FilterRequestCreated(opts *bind.FilterOpts) (*SubmitHandlerRequestCreatedIterator, error) {

	logs, sub, err := _SubmitHandler.contract.FilterLogs(opts, "RequestCreated")
	if err != nil {
		return nil, err
	}
	return &SubmitHandlerRequestCreatedIterator{contract: _SubmitHandler.contract, event: "RequestCreated", logs: logs, sub: sub}, nil
}

// WatchRequestCreated is a free log subscription operation binding the contract event 0x879922cf5fcada9ebaf8bd7424dc62877f4b220cae07fb6695cc1e8f94c52b4d.
//
// Solidity: event RequestCreated(uint256 requestId, address requestor, address to, uint256 weiAmount, bytes32 trieKey, bytes trieValue, bool isExit, bool userActivated)
func (_SubmitHandler *SubmitHandlerFilterer) WatchRequestCreated(opts *bind.WatchOpts, sink chan<- *SubmitHandlerRequestCreated) (event.Subscription, error) {

	logs, sub, err := _SubmitHandler.contract.WatchLogs(opts, "RequestCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubmitHandlerRequestCreated)
				if err := _SubmitHandler.contract.UnpackLog(event, "RequestCreated", log); err != nil {
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
func (_SubmitHandler *SubmitHandlerFilterer) ParseRequestCreated(log types.Log) (*SubmitHandlerRequestCreated, error) {
	event := new(SubmitHandlerRequestCreated)
	if err := _SubmitHandler.contract.UnpackLog(event, "RequestCreated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SubmitHandlerRequestFinalizedIterator is returned from FilterRequestFinalized and is used to iterate over the raw logs and unpacked data for RequestFinalized events raised by the SubmitHandler contract.
type SubmitHandlerRequestFinalizedIterator struct {
	Event *SubmitHandlerRequestFinalized // Event containing the contract specifics and raw log

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
func (it *SubmitHandlerRequestFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubmitHandlerRequestFinalized)
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
		it.Event = new(SubmitHandlerRequestFinalized)
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
func (it *SubmitHandlerRequestFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubmitHandlerRequestFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubmitHandlerRequestFinalized represents a RequestFinalized event raised by the SubmitHandler contract.
type SubmitHandlerRequestFinalized struct {
	RequestId     *big.Int
	UserActivated bool
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRequestFinalized is a free log retrieval operation binding the contract event 0x134017cf3262b18f892ee95dde3b0aec9a80cc70a7c96f09c64bd237aceb0473.
//
// Solidity: event RequestFinalized(uint256 requestId, bool userActivated)
func (_SubmitHandler *SubmitHandlerFilterer) FilterRequestFinalized(opts *bind.FilterOpts) (*SubmitHandlerRequestFinalizedIterator, error) {

	logs, sub, err := _SubmitHandler.contract.FilterLogs(opts, "RequestFinalized")
	if err != nil {
		return nil, err
	}
	return &SubmitHandlerRequestFinalizedIterator{contract: _SubmitHandler.contract, event: "RequestFinalized", logs: logs, sub: sub}, nil
}

// WatchRequestFinalized is a free log subscription operation binding the contract event 0x134017cf3262b18f892ee95dde3b0aec9a80cc70a7c96f09c64bd237aceb0473.
//
// Solidity: event RequestFinalized(uint256 requestId, bool userActivated)
func (_SubmitHandler *SubmitHandlerFilterer) WatchRequestFinalized(opts *bind.WatchOpts, sink chan<- *SubmitHandlerRequestFinalized) (event.Subscription, error) {

	logs, sub, err := _SubmitHandler.contract.WatchLogs(opts, "RequestFinalized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubmitHandlerRequestFinalized)
				if err := _SubmitHandler.contract.UnpackLog(event, "RequestFinalized", log); err != nil {
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
func (_SubmitHandler *SubmitHandlerFilterer) ParseRequestFinalized(log types.Log) (*SubmitHandlerRequestFinalized, error) {
	event := new(SubmitHandlerRequestFinalized)
	if err := _SubmitHandler.contract.UnpackLog(event, "RequestFinalized", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SubmitHandlerRequestableContractMappedIterator is returned from FilterRequestableContractMapped and is used to iterate over the raw logs and unpacked data for RequestableContractMapped events raised by the SubmitHandler contract.
type SubmitHandlerRequestableContractMappedIterator struct {
	Event *SubmitHandlerRequestableContractMapped // Event containing the contract specifics and raw log

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
func (it *SubmitHandlerRequestableContractMappedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubmitHandlerRequestableContractMapped)
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
		it.Event = new(SubmitHandlerRequestableContractMapped)
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
func (it *SubmitHandlerRequestableContractMappedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubmitHandlerRequestableContractMappedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubmitHandlerRequestableContractMapped represents a RequestableContractMapped event raised by the SubmitHandler contract.
type SubmitHandlerRequestableContractMapped struct {
	ContractInRootchain  common.Address
	ContractInChildchain common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterRequestableContractMapped is a free log retrieval operation binding the contract event 0xc5ec2ed49686197edd2ed642c7e6096893cc81e6658cde2527030316037715d0.
//
// Solidity: event RequestableContractMapped(address contractInRootchain, address contractInChildchain)
func (_SubmitHandler *SubmitHandlerFilterer) FilterRequestableContractMapped(opts *bind.FilterOpts) (*SubmitHandlerRequestableContractMappedIterator, error) {

	logs, sub, err := _SubmitHandler.contract.FilterLogs(opts, "RequestableContractMapped")
	if err != nil {
		return nil, err
	}
	return &SubmitHandlerRequestableContractMappedIterator{contract: _SubmitHandler.contract, event: "RequestableContractMapped", logs: logs, sub: sub}, nil
}

// WatchRequestableContractMapped is a free log subscription operation binding the contract event 0xc5ec2ed49686197edd2ed642c7e6096893cc81e6658cde2527030316037715d0.
//
// Solidity: event RequestableContractMapped(address contractInRootchain, address contractInChildchain)
func (_SubmitHandler *SubmitHandlerFilterer) WatchRequestableContractMapped(opts *bind.WatchOpts, sink chan<- *SubmitHandlerRequestableContractMapped) (event.Subscription, error) {

	logs, sub, err := _SubmitHandler.contract.WatchLogs(opts, "RequestableContractMapped")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubmitHandlerRequestableContractMapped)
				if err := _SubmitHandler.contract.UnpackLog(event, "RequestableContractMapped", log); err != nil {
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
func (_SubmitHandler *SubmitHandlerFilterer) ParseRequestableContractMapped(log types.Log) (*SubmitHandlerRequestableContractMapped, error) {
	event := new(SubmitHandlerRequestableContractMapped)
	if err := _SubmitHandler.contract.UnpackLog(event, "RequestableContractMapped", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SubmitHandlerSessionTimeoutIterator is returned from FilterSessionTimeout and is used to iterate over the raw logs and unpacked data for SessionTimeout events raised by the SubmitHandler contract.
type SubmitHandlerSessionTimeoutIterator struct {
	Event *SubmitHandlerSessionTimeout // Event containing the contract specifics and raw log

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
func (it *SubmitHandlerSessionTimeoutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubmitHandlerSessionTimeout)
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
		it.Event = new(SubmitHandlerSessionTimeout)
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
func (it *SubmitHandlerSessionTimeoutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubmitHandlerSessionTimeoutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubmitHandlerSessionTimeout represents a SessionTimeout event raised by the SubmitHandler contract.
type SubmitHandlerSessionTimeout struct {
	UserActivated bool
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSessionTimeout is a free log retrieval operation binding the contract event 0x2122ec719581cd177f225f59a1ee005831211196831f8f1ccffa817d2e7bd108.
//
// Solidity: event SessionTimeout(bool userActivated)
func (_SubmitHandler *SubmitHandlerFilterer) FilterSessionTimeout(opts *bind.FilterOpts) (*SubmitHandlerSessionTimeoutIterator, error) {

	logs, sub, err := _SubmitHandler.contract.FilterLogs(opts, "SessionTimeout")
	if err != nil {
		return nil, err
	}
	return &SubmitHandlerSessionTimeoutIterator{contract: _SubmitHandler.contract, event: "SessionTimeout", logs: logs, sub: sub}, nil
}

// WatchSessionTimeout is a free log subscription operation binding the contract event 0x2122ec719581cd177f225f59a1ee005831211196831f8f1ccffa817d2e7bd108.
//
// Solidity: event SessionTimeout(bool userActivated)
func (_SubmitHandler *SubmitHandlerFilterer) WatchSessionTimeout(opts *bind.WatchOpts, sink chan<- *SubmitHandlerSessionTimeout) (event.Subscription, error) {

	logs, sub, err := _SubmitHandler.contract.WatchLogs(opts, "SessionTimeout")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubmitHandlerSessionTimeout)
				if err := _SubmitHandler.contract.UnpackLog(event, "SessionTimeout", log); err != nil {
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
func (_SubmitHandler *SubmitHandlerFilterer) ParseSessionTimeout(log types.Log) (*SubmitHandlerSessionTimeout, error) {
	event := new(SubmitHandlerSessionTimeout)
	if err := _SubmitHandler.contract.UnpackLog(event, "SessionTimeout", log); err != nil {
		return nil, err
	}
	return event, nil
}
