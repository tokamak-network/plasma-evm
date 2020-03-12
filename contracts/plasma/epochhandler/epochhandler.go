// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package epochhandler

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
var AddressBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a72315820451b843ffb424b86bcba31db5f97373aa893055a6c8333b114cc3f724489f96264736f6c634300050c0032"

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
var BMTBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a72315820e40cfcc115c891081c9c49cf067c463d7e9814fa3ea29c4a8d949bc82adcafb264736f6c634300050c0032"

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
const DataABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"APPLY_IN_CHILDCHAIN_SIGNATURE\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"APPLY_IN_ROOTCHAIN_SIGNATURE\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"NA\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"NA_TX_GAS_LIMIT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"NA_TX_GAS_PRICE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// DataFuncSigs maps the 4-byte function signature to its string representation.
var DataFuncSigs = map[string]string{
	"a89ca766": "APPLY_IN_CHILDCHAIN_SIGNATURE()",
	"a7b6ae28": "APPLY_IN_ROOTCHAIN_SIGNATURE()",
	"ab73ff05": "NA()",
	"90e84f56": "NA_TX_GAS_LIMIT()",
	"1927ac58": "NA_TX_GAS_PRICE()",
}

// DataBin is the compiled bytecode used for deploying new contracts.
var DataBin = "0x6101cf610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100615760003560e01c80631927ac581461006657806390e84f5614610080578063a7b6ae2814610088578063a89ca766146100ad578063ab73ff05146100b5575b600080fd5b61006e6100d9565b60408051918252519081900360200190f35b61006e6100e1565b6100906100e8565b604080516001600160e01b03199092168252519081900360200190f35b610090610103565b6100bd61011e565b604080516001600160a01b039092168252519081900360200190f35b633b9aca0081565b620186a081565b60405180603b6101608239603b019050604051809103902081565b60405180603c6101248239603c019050604051809103902081565b60008156fe6170706c7952657175657374496e4368696c64436861696e28626f6f6c2c75696e743235362c616464726573732c627974657333322c6279746573296170706c7952657175657374496e526f6f74436861696e28626f6f6c2c75696e743235362c616464726573732c627974657333322c627974657329a265627a7a723158204c9301a5408d81b999d438ab1834526edc3015412601cd9e2f1a4da5000c968364736f6c634300050c0032"

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
const EpochHandlerABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"forkNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"BlockFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fork\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epochNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isRequest\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"userActivated\",\"type\":\"bool\"}],\"name\":\"BlockSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"requestor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"trieKey\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"trieValue\",\"type\":\"bytes32\"}],\"name\":\"ERUCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"forkNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epochNumber\",\"type\":\"uint256\"}],\"name\":\"EpochFilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"forkNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epochNumber\",\"type\":\"uint256\"}],\"name\":\"EpochFilling\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"forkNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epochNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startBlockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"endBlockNumber\",\"type\":\"uint256\"}],\"name\":\"EpochFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"forkNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epochNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startBlockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"endBlockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requestStart\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requestEnd\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"epochIsEmpty\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isRequest\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"userActivated\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"rebase\",\"type\":\"bool\"}],\"name\":\"EpochPrepared\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"forkNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epochNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startBlockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"endBlockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requestStart\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requestEnd\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"epochIsEmpty\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isRequest\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"userActivated\",\"type\":\"bool\"}],\"name\":\"EpochRebased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newFork\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epochNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"forkedBlockNumber\",\"type\":\"uint256\"}],\"name\":\"Forked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_newOperator\",\"type\":\"address\"}],\"name\":\"OperatorChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"userActivated\",\"type\":\"bool\"}],\"name\":\"RequestApplied\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"userActivated\",\"type\":\"bool\"}],\"name\":\"RequestChallenged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"requestor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"weiAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"trieKey\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"trieValue\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isExit\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"userActivated\",\"type\":\"bool\"}],\"name\":\"RequestCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"userActivated\",\"type\":\"bool\"}],\"name\":\"RequestFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"contractInRootchain\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"contractInChildchain\",\"type\":\"address\"}],\"name\":\"RequestableContractMapped\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"userActivated\",\"type\":\"bool\"}],\"name\":\"SessionTimeout\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"COST_ERO\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"COST_ERU\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"COST_NRB\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"COST_ORB\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"COST_URB\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"COST_URB_PREPARE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CP_COMPUTATION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CP_EXIT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CP_WITHHOLDING\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"EROIdToFinalize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"EROs\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"isExit\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isTransfer\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"finalized\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"challenged\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"value\",\"type\":\"uint128\"},{\"internalType\":\"addresspayable\",\"name\":\"requestor\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"trieKey\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"trieValue\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ERUIdToFinalize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"ERUs\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"isExit\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isTransfer\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"finalized\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"challenged\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"value\",\"type\":\"uint128\"},{\"internalType\":\"addresspayable\",\"name\":\"requestor\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"trieKey\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"trieValue\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"NRELength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"NULL_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"ORBs\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"submitted\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"numEnter\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"epochNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"requestStart\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"requestEnd\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"trie\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PREPARE_TIMEOUT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"REQUEST_GAS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"URBs\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"submitted\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"numEnter\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"epochNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"requestStart\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"requestEnd\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"trie\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"currentFork\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"development\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"epochHandler\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"etherToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"firstFilledORENumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"firstNonEmptyRequestEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"forks\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"forkedBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"firstEpoch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"lastEpoch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"firstBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"lastBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"lastFinalizedEpoch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"lastFinalizedBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"firstEnterEpoch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"lastEnterEpoch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nextBlockToRebase\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"rebased\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isRootChain\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastAppliedBlockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastAppliedEpochNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastAppliedForkNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"lastNonEmptyRequestEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"numEnterForORB\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"operator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"prepareNRE\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"prepareNREAfterURE\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"prepareORE\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"prepareOREAfterURE\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"prepareToSubmitURB\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"requestableContracts\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"seigManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"submitHandler\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// EpochHandlerFuncSigs maps the 4-byte function signature to its string representation.
var EpochHandlerFuncSigs = map[string]string{
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
	"420bb4b8": "isRootChain()",
	"fb788a27": "lastAppliedBlockNumber()",
	"c8ad329f": "lastAppliedEpochNumber()",
	"164bc2ae": "lastAppliedForkNumber()",
	"b6715647": "lastNonEmptyRequestEpoch(uint256)",
	"23691566": "numEnterForORB()",
	"570ca735": "operator()",
	"03787fa2": "prepareNRE()",
	"5656225b": "prepareNREAfterURE()",
	"12d87f07": "prepareORE()",
	"5e9ef4f3": "prepareOREAfterURE()",
	"e6925d08": "prepareToSubmitURB()",
	"da0185f8": "requestableContracts(address)",
	"6fb7f558": "seigManager()",
	"e259faf7": "submitHandler()",
}

// EpochHandlerBin is the compiled bytecode used for deploying new contracts.
var EpochHandlerBin = "0x608060405234801561001057600080fd5b50600180546001600160a01b031916301790556127da806100326000396000f3fe6080604052600436106102465760003560e01c806394be3aa511610139578063c8ad329f116100b6578063e259faf71161007a578063e259faf71461074a578063e6925d081461075f578063e7b88b8014610767578063ea7f22a81461077c578063f4f31de4146107a6578063fb788a27146107d057610246565b8063c8ad329f146106c3578063ca6f6380146106d8578063d691acd81461024b578063da0185f814610702578063de0ce17d1461073557610246565b8063b6715647116100fd578063b6715647146105e5578063b8066bcb1461060f578063c0e8606414610624578063c2bc88fa14610699578063c54626cc146106ae57610246565b806394be3aa51461024b578063ab96da2d1461049f578063b17fa6e9146104b4578063b2ae9ba81461024b578063b443f3cc146104c957610246565b80634ba3a126116101c757806372ecb9a81161018b57806372ecb9a8146104215780637b929c271461044b5780638155717d146104605780638b5172d0146104755780638eb288ca1461048a57610246565b80634ba3a1261461032b5780635656225b146103cb578063570ca735146103d35780635e9ef4f3146104045780636fb7f5581461040c57610246565b8063183d2d1c1161020e578063183d2d1c146102ae578063192adc5b146102c357806323691566146102d85780632dc6bb7b146102ed578063420bb4b81461030257610246565b8063033cfbed1461024b57806303787fa21461027257806308c4fff01461027c57806312d87f0714610291578063164bc2ae14610299575b600080fd5b34801561025757600080fd5b506102606107e5565b60408051918252519081900360200190f35b61027a6107ef565b005b34801561028857600080fd5b50610260610972565b61027a610977565b3480156102a557600080fd5b50610260611173565b3480156102ba57600080fd5b50610260611179565b3480156102cf57600080fd5b5061026061117f565b3480156102e457600080fd5b50610260611189565b3480156102f957600080fd5b5061026061118f565b34801561030e57600080fd5b50610317611195565b604080519115158252519081900360200190f35b34801561033757600080fd5b506103556004803603602081101561034e57600080fd5b503561119a565b604080516001600160401b039d8e1681529b8d1660208d0152998c168b8b0152978b1660608b0152958a1660808a015293891660a089015291881660c0880152871660e0870152861661010086015285166101208501529093166101408301529115156101608201529051908190036101800190f35b61027a611211565b3480156103df57600080fd5b506103e86113eb565b604080516001600160a01b039092168252519081900360200190f35b61027a6113ff565b34801561041857600080fd5b506103e86115f6565b34801561042d57600080fd5b506102606004803603602081101561044457600080fd5b5035611605565b34801561045757600080fd5b50610317611617565b34801561046c57600080fd5b50610260611620565b34801561048157600080fd5b50610260611625565b34801561049657600080fd5b5061026061162f565b3480156104ab57600080fd5b50610260611636565b3480156104c057600080fd5b5061026061163c565b3480156104d557600080fd5b506104f3600480360360208110156104ec57600080fd5b5035611641565b604080516001600160401b038d1681528b15156020808301919091528b151592820192909252891515606082015288151560808201526001600160801b03881660a08201526001600160a01b0380881660c0830152861660e082015261010081018590526101208101849052610160610140820181815284519183019190915283519192909161018084019185019080838360005b838110156105a0578181015183820152602001610588565b50505050905090810190601f1680156105cd5780820380516001836020036101000a031916815260200191505b509c5050505050505050505050505060405180910390f35b3480156105f157600080fd5b506102606004803603602081101561060857600080fd5b5035611762565b34801561061b57600080fd5b506103e8611774565b34801561063057600080fd5b5061064e6004803603602081101561064757600080fd5b5035611783565b6040805196151587526001600160401b03958616602088015293851686850152918416606086015290921660808401526001600160a01b0390911660a0830152519081900360c00190f35b3480156106a557600080fd5b506102606117ea565b3480156106ba57600080fd5b506102606117ef565b3480156106cf57600080fd5b506102606117f5565b3480156106e457600080fd5b50610260600480360360208110156106fb57600080fd5b50356117fb565b34801561070e57600080fd5b506103e86004803603602081101561072557600080fd5b50356001600160a01b031661180d565b34801561074157600080fd5b506103e8611828565b34801561075657600080fd5b506103e861182d565b61027a61183c565b34801561077357600080fd5b506103e8611d45565b34801561078857600080fd5b5061064e6004803603602081101561079f57600080fd5b5035611d54565b3480156107b257600080fd5b506104f3600480360360208110156107c957600080fd5b5035611d61565b3480156107dc57600080fd5b50610260611d6e565b6509184e72a00081565b600654600081815260086020526040902090158061081857506002810154600160c01b900460ff165b61082157600080fd5b805460016001600160401b03600160801b90920482168101918216600081815260038501602052604090209190811461088857508254600160801b90046001600160401b039081166000908152600385016020526040902054600160401b90048116600101165b8154600160c81b60ff60c81b1990911617600160801b600160c01b031916600160801b426001600160401b0390811691909102919091176001600160401b03191682821617808455600554600160401b600160801b0319909116600160401b91840160001901831682021780855560065460408051918252878516602083015282851682820152928204909316606084015260006080840181905260a0840181905260c0840181905260e0840181905261010084015260ff600160e01b9091041615156101208301525160008051602061278683398151915291610140908290030190a150505050565b600f81565b60065460008181526008602052604090209015806109a057506002810154600160c01b900460ff165b6109a957600080fd5b8054600160801b908190046001600160401b03908116600181810180841660009081526003808801602052604080832091909501861682529381208454600160c81b60ff60c81b1990911617600160801b600160c01b0319164287169097029690961784558554600160d01b60ff60d01b19909116178655600d805487850180546001600160c01b0316600160c01b928916830217905591905591830154909492939190041615610af05760028401546001600160401b0316610a88576002840180546001600160401b0319166001600160401b038516179055610ac9565b600284810154600160401b90046001600160401b039081166000908152600387016020526040902090910180546001600160401b0319169185169190911790555b600284018054600160401b600160801b031916600160401b6001600160401b038616021790555b600654600090815260076020526040902054158015610b1857508154600160c01b900460ff16155b15610b3b5760065460009081526007602052604090206001600160401b03841690555b8154600090600160c01b900460ff16610b8557600180840154610b80916001600160401b03808316600160401b9093048116929092030116610b7b611d74565b611d79565b610b88565b60005b8354909150600160c01b900460ff1615610bef576001600160401b0360001985018116600090815260038701602052604090205484546001600160401b031916600160401b9182900483161791821602600160401b600160801b0319909116178355610c68565b6001600160401b03600019850181166000908152600387016020526040902054610c2991600160401b90910416600163ffffffff611db016565b8354600160401b600160801b03196001600160401b03199091166001600160401b0392831617908116600160401b918316840160001901909216021783555b6009541580610c9357506001830154600954600160401b9091046001600160401b0316600019909101145b15610caa57815460ff60c01b1916600160c01b1782555b8154600160c01b900460ff16610e89576006546000908152600760205260409020546001600160401b038516148015610ceb57508254600160e01b900460ff165b15610d4f576001808401805484830180546001600160401b03600160401b9093048316850183166001600160401b0319909116178082559254600160801b908190048316909401909116909202600160801b600160c01b0319909116179055610e84565b600654600090815260076020526040902054610db9576001808401805491840180546001600160401b0319166001600160401b03600160401b9094048416178082559154600160801b90819004909316909202600160801b600160c01b0319909116179055610e84565b8254600160c01b900460ff16610e29576001808401805484830180546001600160401b03600160401b909304831690940182166001600160401b0319909416939093178084559154600160801b908190048216850190911602600160801b600160c01b0319909116179055610e84565b6001808401805484830180546001600160401b03600160401b9093048316850183166001600160401b0319909116178082559254600160801b908190048316909401909116909202600160801b600160c01b03199091161790555b610f3c565b6001838101549083018054600160401b9092046001600160401b03166001600160401b03199092169190911790558254600160c01b900460ff16610f08576001838101549083018054600160801b600160c01b031916600160801b928390046001600160401b0390811685016000190116909202919091179055610f3c565b6001838101549083018054600160801b600160c01b031916600160801b928390046001600160401b03169092029190911790555b8254600160c01b900460ff16610fed576006546000908152600f6020526040902054610f80576006546000908152600f602052604090206001600160401b03851690555b6006546000908152600e60205260409020548015610fce57600081815260038701602052604090206002018054600160401b600160801b031916600160401b6001600160401b038816021790555b506006546000908152600e602052604090206001600160401b03851690555b8154600160c01b900460ff161561102b57600182018054600160401b6001600160401b03821602600160401b600160801b0319909116179055611099565b60095460018381018054600160401b600160801b031916600160401b6000199094016001600160401b03169390930292909217909155600b805461106f9083611dce565b8154811061107957fe5b60009182526020909120600290910201805460ff19169115159190911790555b6006548354600180860154604080519485526001600160401b03808a16602087015280851686830152600160401b80860482166060880152838216608088015290920490911660a085015260ff600160c01b84048116151560c086015260e08501929092526000610100850152600160e01b90920416151561012083015251600080516020612786833981519152918190036101400190a18254600160c01b900460ff161561116c578454600160801b600160c01b031916600160801b6001600160401b0386160217855561116c6107ef565b5050505050565b60105481565b60065481565b6551dac207a00081565b600d5481565b60135481565b600181565b6008602052600090815260409020805460018201546002909201546001600160401b0380831693600160401b808504831694600160801b808204851695600160c01b92839004861695858116958581048216958482048316959182900483169484841694918204841693908204169160ff9104168c565b6006546000818152600860208190526040822092611251918491849061123e90600163ffffffff611dce16565b8152602001908152602001600020611de3565b825460065460016001600160401b03600160801b90930483168101808416600081815260038901602090815260408083205481519788529187019390935290951684820152606084018590526080840185905260a0840185905285151560c085015260e0840185905261010084019490945261012083019190915291519293509091600080516020612786833981519152918190036101400190a181156113e6576001808401546001600160401b03838116600081815260038801602090815260408083208054600160401b9787168802600160401b600160801b03199091161781558a54600160801b8602600160801b600160c01b0319909116178b5560028b01805460ff60c01b1916600160c01b1790556006548154918901548351918252938101959095528086168583015286900485166060850152818516608085015294900490921660a082015260c081019390935260e08301819052610100830152517f030c1c69405c93021f28f57557240dee939a320b826a1fd0d39bf6e629ecab47918190036101200190a16113e66107ef565b505050565b60005461010090046001600160a01b031681565b6006546000818152600860208190526040822092611441918491849061142c90600163ffffffff611dce16565b8152602001908152602001600020600b6120a5565b8254909150600160801b90046001600160401b03166001018161147c5760065460009081526007602052604090206001600160401b03821690555b6006546001600160401b0380831660008181526003870160209081526040808320805460019182015483519889529388019590955284861687830152600160401b948590048616606088015282861660808801529390910490931660a085015285151560c085015260e0840182905261010084015261012083015251600080516020612786833981519152918190036101400190a181156113e6576001808401546001600160401b03838116600081815260038801602090815260408083208054600160401b9787168802600160401b600160801b03199091161781558a54600160801b8602600160801b600160c01b0319909116178b556006548154918901548351918252938101959095528086168583015286900485166060850152818516608085015294900490921660a082015260c0810184905260e0810193909352610100830152517f030c1c69405c93021f28f57557240dee939a320b826a1fd0d39bf6e629ecab47918190036101200190a16113e6611211565b6004546001600160a01b031681565b60076020526000908152604090205481565b60005460ff1681565b600a81565b6512309ce5400081565b620186a081565b60055481565b601481565b6009818154811061164e57fe5b6000918252602091829020600691909102018054600180830154600280850154600386015460048701546005880180546040805161010099831615999099026000190190911695909504601f81018b90048b0288018b019095528487526001600160401b0388169a50600160401b880460ff9081169a600160481b8a0482169a600160501b8b0483169a600160581b810490931699600160601b9093046001600160801b0316986001600160a01b03908116989716969093918301828280156117585780601f1061172d57610100808354040283529160200191611758565b820191906000526020600020905b81548152906001019060200180831161173b57829003601f168201915b505050505090508b565b600e6020526000908152604090205481565b6003546001600160a01b031681565b600c818154811061179057fe5b60009182526020909120600290910201805460019091015460ff821692506001600160401b036101008304811692600160481b8104821692600160881b9091048216918116906001600160a01b03600160401b9091041686565b603c81565b60145481565b60115481565b600f6020526000908152604090205481565b6015602052600090815260409020546001600160a01b031681565b600081565b6002546001600160a01b031681565b600654600081815260086020908152604080832060018086018552828520818301805484546001600160401b03600160801b928390048116860181166001600160401b0319909216919091178087558454600160c01b91831682026001600160c01b039182161780875582810484168c52600489018b52898c2054600160401b9085168102600160401b600160801b03199092169190911781810485168602600160801b600160c01b031991821617808955965498880180544287169586029a8890048716880290831617909316989098179091559093041688526003830190965293909520805460ff60d81b1960ff60d01b199590960260ff60c81b19909116600160c81b179092169190911792909216600160d01b1792909216600160d81b178155909215908161199f5783546001600160401b03600160401b918290048116600090815260038701602052604090206001908101549290920416016119a2565b60005b6001820180546001600160401b0319166001600160401b0392831617808255600a54600160401b600160801b0319909116600160401b60001990920184168202179182905581048216911611156119f557fe5b6001810154600090611a1a906001600160401b0380821691600160401b90041661249b565b845483546001600160401b0319166001600160401b03600160c01b909204821617808555919250611a6e91600191611a5991168463ffffffff611db016565b6001600160401b03169063ffffffff6124d516565b82546001600160401b0391909116600160401b02600160401b600160801b031990911617825582611b355784546001600160401b03600160401b9182900481166000908152600388016020526040902054611b3092611af492600192611adf9281048216911663ffffffff6124d516565b6001600160401b03169063ffffffff611db016565b86546001600160401b03600160401b90910481166000908152600389016020526040902060010154600160801b9004169063ffffffff611db016565b611b38565b60005b6001830180546001600160401b0392909216600160801b02600160801b600160c01b031990921691909117905560005b816001600160401b0316816001600160401b03161015611c945782546001906004870190600090611ba8906001600160401b03168563ffffffff611db016565b6001600160401b03908116825260208201929092526040016000908120600501805460ff191693151593909317909255845460019260048901929091611bef911685611db0565b6001600160401b0390811682526020820192909252604001600090812060050180549315156101000261ff00199094169390931790925560018501548554600160801b909104821684019260048901929091611c5291168563ffffffff611db016565b6001600160401b03908116825260208201929092526040016000208054600160401b600160801b031916600160401b9390921692909202179055600101611b68565b5060065484548354600185810154604080519290950182526001600160401b03600160401b94859004811660208401528381168387015284840481166060840152818116608084015293900490921660a0830152600060c083015260ff600160d01b82048116151560e0840152600160d81b820481161515610100840152600160e01b9091041615156101208201529051600080516020612786833981519152918190036101400190a15050505050565b6001546001600160a01b031681565b600b818154811061179057fe5b600a818154811061164e57fe5b60125481565b601490565b6000818381611d8457fe5b0615611d9c57818381611d9357fe5b04600101611da7565b818381611da557fe5b045b90505b92915050565b60008282016001600160401b038085169082161015611da757600080fd5b600082821115611ddd57600080fd5b50900390565b6002820154600090600160c01b900460ff1615611dff57600080fd5b8254600160801b90046001600160401b0316600090815260038401602052604090208054600160e01b900460ff168015611e4157508054600160d01b900460ff165b8015611e5657508054600160d81b900460ff16155b611e5f57600080fd5b5060018381015484546001600160401b03600160801b91829004811660009081526003808901602081815260408085208054600160401b600160801b031916988716600160401b02989098179097558a548a548616855260048b0182528785205490879004861690980180861685529181528684208054600160e01b600160c81b60ff60c81b199092169190911760ff60e01b191617600160801b600160c01b031916428716909702969096178655969093168083529088019095529283205491939092909160ff600160d01b9091041615611f3e5781600101611f40565b815b6000818152600388016020526040902054909150600160c81b900460ff16611fca575050506001848101805483546001600160401b039182166001600160401b0319909116178085559154600160c01b600160401b600160801b03199093169116600160401b021760ff60c01b199081168217909355600286018054909316179091559050611daa565b6000818152600387016020526040902054600160d01b900460ff1615611fec57fe5b60018781015461200d916001600160401b039091169063ffffffff611db016565b84546001600160401b0319166001600160401b039190911617845581811461204e5760008181526003870160205260409020546001600160401b031661205a565b85546001600160401b03165b600288018054600160801b600160c01b031916600160801b6001600160401b0393841681029190911791829055885483169104909116101561209857fe5b5060009695505050505050565b6002830154600090600160c01b900460ff16156120c157600080fd5b8354600160801b90046001600160401b0316600090815260038501602052604090208054600160d01b900460ff16801561210357508054600160d81b900460ff165b61210c57600080fd5b50835483546001600160401b039081166000908152600486016020908152604080832054600160801b95869004851660010180861685526003808c0185528386208054600160e01b600160d01b600160c81b60ff60c81b199093169290921760ff60d01b1916821760ff60e01b191617600160801b600160c01b031916428a16909a0299909917815592909616808652958a019093529083205490949193929160ff9104166121be57816001016121c0565b815b90506121ca6124f6565b600187018054600160401b600160801b031916600160401b6001600160401b0393841602176001600160401b03191692821692909217600160801b600160c01b031916600160801b93909116929092029190911790556000818152600388016020526040902054600160c81b900460ff166122a0575050815460ff60c01b1916600160c01b17808355600187810180546001600160401b03199093166001600160401b03938416178086559054600160401b931692909202600160401b600160801b031990921691909117909255509050612494565b6000818152600388016020526040902054600160d01b900460ff166122c157fe5b60028701548454600160c01b6001600160401b03600160401b909304929092168411820260ff60c01b19909116178086550460ff1615612344575050506001808601805483546001600160401b0319166001600160401b03918216178085559154600160401b911602600160401b600160801b0319909116179091559050612494565b805b6000828152600389016020526040902060010154600160c01b90046001600160401b031661237657600201612346565b600189810154612397916001600160401b039091169063ffffffff611db016565b85546001600160401b0319166001600160401b03918216178655600082815260038a016020908152604080832054841680845260048d019092528220548a5491938b92600160401b909204169081106123ec57fe5b906000526020600020906002020190505b805461010090046001600160401b031661245d57600191909101600081815260048b01602052604090205489549192918a91600160401b90046001600160401b031690811061244857fe5b906000526020600020906002020190506123fd565b5060028a0180546001600160401b03909216600160801b02600160801b600160c01b03199092169190911790555060009450505050505b9392505050565b6000611da76124a8611d74565b6124c960016124bd868863ffffffff611dce16565b9063ffffffff61273716565b9063ffffffff611d7916565b6000826001600160401b0316826001600160401b03161115611ddd57600080fd5b6000806000806125126001600654611dce90919063ffffffff16565b90505b60008181526008602052604081209061252d82612749565b82546001600160401b03600160801b90910481166001018116600090815260038501602052604081205492909116925090600160d01b900460ff16612583578254600160801b90046001600160401b0316612599565b8254600160801b90046001600160401b03166001015b6001600160401b031690505b80821115612653576125be84600163ffffffff611dce16565b600081815260086020526040902090945092506125da83612749565b83546001600160401b03600160801b90910481166001018116600090815260038601602052604090205491169250600160d01b900460ff1661262d578254600160801b90046001600160401b0316612643565b8254600160801b90046001600160401b03166001015b6001600160401b031690506125a5565b6000818152600384016020526040902054600160c81b900460ff16156127175782546001600160401b0390811660009081526004850160209081526040808320549093168083526003870190915291902054600160d01b900460ff166126b7576001015b6000908152600384016020908152604080832093835280832080546001600160401b0390811685526004909701909252909120546001928301549290910154600160401b9182900485169850918416965090049091169250612732915050565b61272884600163ffffffff611dce16565b9350505050612515565b909192565b600082820183811015611da757600080fd5b80546000906001600160401b031661276057600080fd5b5080546001600160401b0390811660009081526004830160205260409020541691905056fe1a69c0760aa329b76f72579129869013ebd3d41594db019c0e997b939fcb32e3a265627a7a72315820d026c7bc076d20dbc9dd79bc69a40fe3bae5cbc9bb207c931b92eb4ed325519564736f6c634300050c0032"

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

// CPEXIT is a free data retrieval call binding the contract method 0x8155717d.
//
// Solidity: function CP_EXIT() constant returns(uint256)
func (_EpochHandler *EpochHandlerCaller) CPEXIT(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _EpochHandler.contract.Call(opts, out, "CP_EXIT")
	return *ret0, err
}

// CPEXIT is a free data retrieval call binding the contract method 0x8155717d.
//
// Solidity: function CP_EXIT() constant returns(uint256)
func (_EpochHandler *EpochHandlerSession) CPEXIT() (*big.Int, error) {
	return _EpochHandler.Contract.CPEXIT(&_EpochHandler.CallOpts)
}

// CPEXIT is a free data retrieval call binding the contract method 0x8155717d.
//
// Solidity: function CP_EXIT() constant returns(uint256)
func (_EpochHandler *EpochHandlerCallerSession) CPEXIT() (*big.Int, error) {
	return _EpochHandler.Contract.CPEXIT(&_EpochHandler.CallOpts)
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

// EROIdToFinalize is a free data retrieval call binding the contract method 0x2dc6bb7b.
//
// Solidity: function EROIdToFinalize() constant returns(uint256)
func (_EpochHandler *EpochHandlerCaller) EROIdToFinalize(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _EpochHandler.contract.Call(opts, out, "EROIdToFinalize")
	return *ret0, err
}

// EROIdToFinalize is a free data retrieval call binding the contract method 0x2dc6bb7b.
//
// Solidity: function EROIdToFinalize() constant returns(uint256)
func (_EpochHandler *EpochHandlerSession) EROIdToFinalize() (*big.Int, error) {
	return _EpochHandler.Contract.EROIdToFinalize(&_EpochHandler.CallOpts)
}

// EROIdToFinalize is a free data retrieval call binding the contract method 0x2dc6bb7b.
//
// Solidity: function EROIdToFinalize() constant returns(uint256)
func (_EpochHandler *EpochHandlerCallerSession) EROIdToFinalize() (*big.Int, error) {
	return _EpochHandler.Contract.EROIdToFinalize(&_EpochHandler.CallOpts)
}

// EROs is a free data retrieval call binding the contract method 0xb443f3cc.
//
// Solidity: function EROs(uint256 ) constant returns(uint64 timestamp, bool isExit, bool isTransfer, bool finalized, bool challenged, uint128 value, address requestor, address to, bytes32 trieKey, bytes32 hash, bytes trieValue)
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
	err := _EpochHandler.contract.Call(opts, out, "EROs", arg0)
	return *ret, err
}

// EROs is a free data retrieval call binding the contract method 0xb443f3cc.
//
// Solidity: function EROs(uint256 ) constant returns(uint64 timestamp, bool isExit, bool isTransfer, bool finalized, bool challenged, uint128 value, address requestor, address to, bytes32 trieKey, bytes32 hash, bytes trieValue)
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
	Hash       [32]byte
	TrieValue  []byte
}, error) {
	return _EpochHandler.Contract.EROs(&_EpochHandler.CallOpts, arg0)
}

// EROs is a free data retrieval call binding the contract method 0xb443f3cc.
//
// Solidity: function EROs(uint256 ) constant returns(uint64 timestamp, bool isExit, bool isTransfer, bool finalized, bool challenged, uint128 value, address requestor, address to, bytes32 trieKey, bytes32 hash, bytes trieValue)
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
	Hash       [32]byte
	TrieValue  []byte
}, error) {
	return _EpochHandler.Contract.EROs(&_EpochHandler.CallOpts, arg0)
}

// ERUIdToFinalize is a free data retrieval call binding the contract method 0xc54626cc.
//
// Solidity: function ERUIdToFinalize() constant returns(uint256)
func (_EpochHandler *EpochHandlerCaller) ERUIdToFinalize(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _EpochHandler.contract.Call(opts, out, "ERUIdToFinalize")
	return *ret0, err
}

// ERUIdToFinalize is a free data retrieval call binding the contract method 0xc54626cc.
//
// Solidity: function ERUIdToFinalize() constant returns(uint256)
func (_EpochHandler *EpochHandlerSession) ERUIdToFinalize() (*big.Int, error) {
	return _EpochHandler.Contract.ERUIdToFinalize(&_EpochHandler.CallOpts)
}

// ERUIdToFinalize is a free data retrieval call binding the contract method 0xc54626cc.
//
// Solidity: function ERUIdToFinalize() constant returns(uint256)
func (_EpochHandler *EpochHandlerCallerSession) ERUIdToFinalize() (*big.Int, error) {
	return _EpochHandler.Contract.ERUIdToFinalize(&_EpochHandler.CallOpts)
}

// ERUs is a free data retrieval call binding the contract method 0xf4f31de4.
//
// Solidity: function ERUs(uint256 ) constant returns(uint64 timestamp, bool isExit, bool isTransfer, bool finalized, bool challenged, uint128 value, address requestor, address to, bytes32 trieKey, bytes32 hash, bytes trieValue)
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
	err := _EpochHandler.contract.Call(opts, out, "ERUs", arg0)
	return *ret, err
}

// ERUs is a free data retrieval call binding the contract method 0xf4f31de4.
//
// Solidity: function ERUs(uint256 ) constant returns(uint64 timestamp, bool isExit, bool isTransfer, bool finalized, bool challenged, uint128 value, address requestor, address to, bytes32 trieKey, bytes32 hash, bytes trieValue)
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
	Hash       [32]byte
	TrieValue  []byte
}, error) {
	return _EpochHandler.Contract.ERUs(&_EpochHandler.CallOpts, arg0)
}

// ERUs is a free data retrieval call binding the contract method 0xf4f31de4.
//
// Solidity: function ERUs(uint256 ) constant returns(uint64 timestamp, bool isExit, bool isTransfer, bool finalized, bool challenged, uint128 value, address requestor, address to, bytes32 trieKey, bytes32 hash, bytes trieValue)
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
	Hash       [32]byte
	TrieValue  []byte
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
// Solidity: function ORBs(uint256 ) constant returns(bool submitted, uint64 numEnter, uint64 epochNumber, uint64 requestStart, uint64 requestEnd, address trie)
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
// Solidity: function ORBs(uint256 ) constant returns(bool submitted, uint64 numEnter, uint64 epochNumber, uint64 requestStart, uint64 requestEnd, address trie)
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
// Solidity: function ORBs(uint256 ) constant returns(bool submitted, uint64 numEnter, uint64 epochNumber, uint64 requestStart, uint64 requestEnd, address trie)
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
// Solidity: function URBs(uint256 ) constant returns(bool submitted, uint64 numEnter, uint64 epochNumber, uint64 requestStart, uint64 requestEnd, address trie)
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
// Solidity: function URBs(uint256 ) constant returns(bool submitted, uint64 numEnter, uint64 epochNumber, uint64 requestStart, uint64 requestEnd, address trie)
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
// Solidity: function URBs(uint256 ) constant returns(bool submitted, uint64 numEnter, uint64 epochNumber, uint64 requestStart, uint64 requestEnd, address trie)
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

// EtherToken is a free data retrieval call binding the contract method 0xb8066bcb.
//
// Solidity: function etherToken() constant returns(address)
func (_EpochHandler *EpochHandlerCaller) EtherToken(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _EpochHandler.contract.Call(opts, out, "etherToken")
	return *ret0, err
}

// EtherToken is a free data retrieval call binding the contract method 0xb8066bcb.
//
// Solidity: function etherToken() constant returns(address)
func (_EpochHandler *EpochHandlerSession) EtherToken() (common.Address, error) {
	return _EpochHandler.Contract.EtherToken(&_EpochHandler.CallOpts)
}

// EtherToken is a free data retrieval call binding the contract method 0xb8066bcb.
//
// Solidity: function etherToken() constant returns(address)
func (_EpochHandler *EpochHandlerCallerSession) EtherToken() (common.Address, error) {
	return _EpochHandler.Contract.EtherToken(&_EpochHandler.CallOpts)
}

// FirstFilledORENumber is a free data retrieval call binding the contract method 0x72ecb9a8.
//
// Solidity: function firstFilledORENumber(uint256 ) constant returns(uint256)
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
// Solidity: function firstFilledORENumber(uint256 ) constant returns(uint256)
func (_EpochHandler *EpochHandlerSession) FirstFilledORENumber(arg0 *big.Int) (*big.Int, error) {
	return _EpochHandler.Contract.FirstFilledORENumber(&_EpochHandler.CallOpts, arg0)
}

// FirstFilledORENumber is a free data retrieval call binding the contract method 0x72ecb9a8.
//
// Solidity: function firstFilledORENumber(uint256 ) constant returns(uint256)
func (_EpochHandler *EpochHandlerCallerSession) FirstFilledORENumber(arg0 *big.Int) (*big.Int, error) {
	return _EpochHandler.Contract.FirstFilledORENumber(&_EpochHandler.CallOpts, arg0)
}

// FirstNonEmptyRequestEpoch is a free data retrieval call binding the contract method 0xca6f6380.
//
// Solidity: function firstNonEmptyRequestEpoch(uint256 ) constant returns(uint256)
func (_EpochHandler *EpochHandlerCaller) FirstNonEmptyRequestEpoch(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _EpochHandler.contract.Call(opts, out, "firstNonEmptyRequestEpoch", arg0)
	return *ret0, err
}

// FirstNonEmptyRequestEpoch is a free data retrieval call binding the contract method 0xca6f6380.
//
// Solidity: function firstNonEmptyRequestEpoch(uint256 ) constant returns(uint256)
func (_EpochHandler *EpochHandlerSession) FirstNonEmptyRequestEpoch(arg0 *big.Int) (*big.Int, error) {
	return _EpochHandler.Contract.FirstNonEmptyRequestEpoch(&_EpochHandler.CallOpts, arg0)
}

// FirstNonEmptyRequestEpoch is a free data retrieval call binding the contract method 0xca6f6380.
//
// Solidity: function firstNonEmptyRequestEpoch(uint256 ) constant returns(uint256)
func (_EpochHandler *EpochHandlerCallerSession) FirstNonEmptyRequestEpoch(arg0 *big.Int) (*big.Int, error) {
	return _EpochHandler.Contract.FirstNonEmptyRequestEpoch(&_EpochHandler.CallOpts, arg0)
}

// Forks is a free data retrieval call binding the contract method 0x4ba3a126.
//
// Solidity: function forks(uint256 ) constant returns(uint64 forkedBlock, uint64 firstEpoch, uint64 lastEpoch, uint64 firstBlock, uint64 lastBlock, uint64 lastFinalizedEpoch, uint64 lastFinalizedBlock, uint64 timestamp, uint64 firstEnterEpoch, uint64 lastEnterEpoch, uint64 nextBlockToRebase, bool rebased)
func (_EpochHandler *EpochHandlerCaller) Forks(opts *bind.CallOpts, arg0 *big.Int) (struct {
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
	err := _EpochHandler.contract.Call(opts, out, "forks", arg0)
	return *ret, err
}

// Forks is a free data retrieval call binding the contract method 0x4ba3a126.
//
// Solidity: function forks(uint256 ) constant returns(uint64 forkedBlock, uint64 firstEpoch, uint64 lastEpoch, uint64 firstBlock, uint64 lastBlock, uint64 lastFinalizedEpoch, uint64 lastFinalizedBlock, uint64 timestamp, uint64 firstEnterEpoch, uint64 lastEnterEpoch, uint64 nextBlockToRebase, bool rebased)
func (_EpochHandler *EpochHandlerSession) Forks(arg0 *big.Int) (struct {
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
	return _EpochHandler.Contract.Forks(&_EpochHandler.CallOpts, arg0)
}

// Forks is a free data retrieval call binding the contract method 0x4ba3a126.
//
// Solidity: function forks(uint256 ) constant returns(uint64 forkedBlock, uint64 firstEpoch, uint64 lastEpoch, uint64 firstBlock, uint64 lastBlock, uint64 lastFinalizedEpoch, uint64 lastFinalizedBlock, uint64 timestamp, uint64 firstEnterEpoch, uint64 lastEnterEpoch, uint64 nextBlockToRebase, bool rebased)
func (_EpochHandler *EpochHandlerCallerSession) Forks(arg0 *big.Int) (struct {
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
	return _EpochHandler.Contract.Forks(&_EpochHandler.CallOpts, arg0)
}

// IsRootChain is a free data retrieval call binding the contract method 0x420bb4b8.
//
// Solidity: function isRootChain() constant returns(bool)
func (_EpochHandler *EpochHandlerCaller) IsRootChain(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _EpochHandler.contract.Call(opts, out, "isRootChain")
	return *ret0, err
}

// IsRootChain is a free data retrieval call binding the contract method 0x420bb4b8.
//
// Solidity: function isRootChain() constant returns(bool)
func (_EpochHandler *EpochHandlerSession) IsRootChain() (bool, error) {
	return _EpochHandler.Contract.IsRootChain(&_EpochHandler.CallOpts)
}

// IsRootChain is a free data retrieval call binding the contract method 0x420bb4b8.
//
// Solidity: function isRootChain() constant returns(bool)
func (_EpochHandler *EpochHandlerCallerSession) IsRootChain() (bool, error) {
	return _EpochHandler.Contract.IsRootChain(&_EpochHandler.CallOpts)
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

// LastAppliedEpochNumber is a free data retrieval call binding the contract method 0xc8ad329f.
//
// Solidity: function lastAppliedEpochNumber() constant returns(uint256)
func (_EpochHandler *EpochHandlerCaller) LastAppliedEpochNumber(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _EpochHandler.contract.Call(opts, out, "lastAppliedEpochNumber")
	return *ret0, err
}

// LastAppliedEpochNumber is a free data retrieval call binding the contract method 0xc8ad329f.
//
// Solidity: function lastAppliedEpochNumber() constant returns(uint256)
func (_EpochHandler *EpochHandlerSession) LastAppliedEpochNumber() (*big.Int, error) {
	return _EpochHandler.Contract.LastAppliedEpochNumber(&_EpochHandler.CallOpts)
}

// LastAppliedEpochNumber is a free data retrieval call binding the contract method 0xc8ad329f.
//
// Solidity: function lastAppliedEpochNumber() constant returns(uint256)
func (_EpochHandler *EpochHandlerCallerSession) LastAppliedEpochNumber() (*big.Int, error) {
	return _EpochHandler.Contract.LastAppliedEpochNumber(&_EpochHandler.CallOpts)
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

// LastNonEmptyRequestEpoch is a free data retrieval call binding the contract method 0xb6715647.
//
// Solidity: function lastNonEmptyRequestEpoch(uint256 ) constant returns(uint256)
func (_EpochHandler *EpochHandlerCaller) LastNonEmptyRequestEpoch(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _EpochHandler.contract.Call(opts, out, "lastNonEmptyRequestEpoch", arg0)
	return *ret0, err
}

// LastNonEmptyRequestEpoch is a free data retrieval call binding the contract method 0xb6715647.
//
// Solidity: function lastNonEmptyRequestEpoch(uint256 ) constant returns(uint256)
func (_EpochHandler *EpochHandlerSession) LastNonEmptyRequestEpoch(arg0 *big.Int) (*big.Int, error) {
	return _EpochHandler.Contract.LastNonEmptyRequestEpoch(&_EpochHandler.CallOpts, arg0)
}

// LastNonEmptyRequestEpoch is a free data retrieval call binding the contract method 0xb6715647.
//
// Solidity: function lastNonEmptyRequestEpoch(uint256 ) constant returns(uint256)
func (_EpochHandler *EpochHandlerCallerSession) LastNonEmptyRequestEpoch(arg0 *big.Int) (*big.Int, error) {
	return _EpochHandler.Contract.LastNonEmptyRequestEpoch(&_EpochHandler.CallOpts, arg0)
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
// Solidity: function requestableContracts(address ) constant returns(address)
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
// Solidity: function requestableContracts(address ) constant returns(address)
func (_EpochHandler *EpochHandlerSession) RequestableContracts(arg0 common.Address) (common.Address, error) {
	return _EpochHandler.Contract.RequestableContracts(&_EpochHandler.CallOpts, arg0)
}

// RequestableContracts is a free data retrieval call binding the contract method 0xda0185f8.
//
// Solidity: function requestableContracts(address ) constant returns(address)
func (_EpochHandler *EpochHandlerCallerSession) RequestableContracts(arg0 common.Address) (common.Address, error) {
	return _EpochHandler.Contract.RequestableContracts(&_EpochHandler.CallOpts, arg0)
}

// SeigManager is a free data retrieval call binding the contract method 0x6fb7f558.
//
// Solidity: function seigManager() constant returns(address)
func (_EpochHandler *EpochHandlerCaller) SeigManager(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _EpochHandler.contract.Call(opts, out, "seigManager")
	return *ret0, err
}

// SeigManager is a free data retrieval call binding the contract method 0x6fb7f558.
//
// Solidity: function seigManager() constant returns(address)
func (_EpochHandler *EpochHandlerSession) SeigManager() (common.Address, error) {
	return _EpochHandler.Contract.SeigManager(&_EpochHandler.CallOpts)
}

// SeigManager is a free data retrieval call binding the contract method 0x6fb7f558.
//
// Solidity: function seigManager() constant returns(address)
func (_EpochHandler *EpochHandlerCallerSession) SeigManager() (common.Address, error) {
	return _EpochHandler.Contract.SeigManager(&_EpochHandler.CallOpts)
}

// SubmitHandler is a free data retrieval call binding the contract method 0xe259faf7.
//
// Solidity: function submitHandler() constant returns(address)
func (_EpochHandler *EpochHandlerCaller) SubmitHandler(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _EpochHandler.contract.Call(opts, out, "submitHandler")
	return *ret0, err
}

// SubmitHandler is a free data retrieval call binding the contract method 0xe259faf7.
//
// Solidity: function submitHandler() constant returns(address)
func (_EpochHandler *EpochHandlerSession) SubmitHandler() (common.Address, error) {
	return _EpochHandler.Contract.SubmitHandler(&_EpochHandler.CallOpts)
}

// SubmitHandler is a free data retrieval call binding the contract method 0xe259faf7.
//
// Solidity: function submitHandler() constant returns(address)
func (_EpochHandler *EpochHandlerCallerSession) SubmitHandler() (common.Address, error) {
	return _EpochHandler.Contract.SubmitHandler(&_EpochHandler.CallOpts)
}

// PrepareNRE is a paid mutator transaction binding the contract method 0x03787fa2.
//
// Solidity: function prepareNRE() returns()
func (_EpochHandler *EpochHandlerTransactor) PrepareNRE(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EpochHandler.contract.Transact(opts, "prepareNRE")
}

// PrepareNRE is a paid mutator transaction binding the contract method 0x03787fa2.
//
// Solidity: function prepareNRE() returns()
func (_EpochHandler *EpochHandlerSession) PrepareNRE() (*types.Transaction, error) {
	return _EpochHandler.Contract.PrepareNRE(&_EpochHandler.TransactOpts)
}

// PrepareNRE is a paid mutator transaction binding the contract method 0x03787fa2.
//
// Solidity: function prepareNRE() returns()
func (_EpochHandler *EpochHandlerTransactorSession) PrepareNRE() (*types.Transaction, error) {
	return _EpochHandler.Contract.PrepareNRE(&_EpochHandler.TransactOpts)
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

// PrepareORE is a paid mutator transaction binding the contract method 0x12d87f07.
//
// Solidity: function prepareORE() returns()
func (_EpochHandler *EpochHandlerTransactor) PrepareORE(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EpochHandler.contract.Transact(opts, "prepareORE")
}

// PrepareORE is a paid mutator transaction binding the contract method 0x12d87f07.
//
// Solidity: function prepareORE() returns()
func (_EpochHandler *EpochHandlerSession) PrepareORE() (*types.Transaction, error) {
	return _EpochHandler.Contract.PrepareORE(&_EpochHandler.TransactOpts)
}

// PrepareORE is a paid mutator transaction binding the contract method 0x12d87f07.
//
// Solidity: function prepareORE() returns()
func (_EpochHandler *EpochHandlerTransactorSession) PrepareORE() (*types.Transaction, error) {
	return _EpochHandler.Contract.PrepareORE(&_EpochHandler.TransactOpts)
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
// Solidity: event BlockFinalized(uint256 forkNumber, uint256 blockNumber)
func (_EpochHandler *EpochHandlerFilterer) FilterBlockFinalized(opts *bind.FilterOpts) (*EpochHandlerBlockFinalizedIterator, error) {

	logs, sub, err := _EpochHandler.contract.FilterLogs(opts, "BlockFinalized")
	if err != nil {
		return nil, err
	}
	return &EpochHandlerBlockFinalizedIterator{contract: _EpochHandler.contract, event: "BlockFinalized", logs: logs, sub: sub}, nil
}

// WatchBlockFinalized is a free log subscription operation binding the contract event 0xfb96205e4b3633fd57aa805b26b51ecf528714a10241a4af015929dce86768d9.
//
// Solidity: event BlockFinalized(uint256 forkNumber, uint256 blockNumber)
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

// ParseBlockFinalized is a log parse operation binding the contract event 0xfb96205e4b3633fd57aa805b26b51ecf528714a10241a4af015929dce86768d9.
//
// Solidity: event BlockFinalized(uint256 forkNumber, uint256 blockNumber)
func (_EpochHandler *EpochHandlerFilterer) ParseBlockFinalized(log types.Log) (*EpochHandlerBlockFinalized, error) {
	event := new(EpochHandlerBlockFinalized)
	if err := _EpochHandler.contract.UnpackLog(event, "BlockFinalized", log); err != nil {
		return nil, err
	}
	return event, nil
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
// Solidity: event BlockSubmitted(uint256 fork, uint256 epochNumber, uint256 blockNumber, bool isRequest, bool userActivated)
func (_EpochHandler *EpochHandlerFilterer) FilterBlockSubmitted(opts *bind.FilterOpts) (*EpochHandlerBlockSubmittedIterator, error) {

	logs, sub, err := _EpochHandler.contract.FilterLogs(opts, "BlockSubmitted")
	if err != nil {
		return nil, err
	}
	return &EpochHandlerBlockSubmittedIterator{contract: _EpochHandler.contract, event: "BlockSubmitted", logs: logs, sub: sub}, nil
}

// WatchBlockSubmitted is a free log subscription operation binding the contract event 0x3d4a04291c66b06f39a4ecb817875b12b5485a05ec563133a56a905305c48e55.
//
// Solidity: event BlockSubmitted(uint256 fork, uint256 epochNumber, uint256 blockNumber, bool isRequest, bool userActivated)
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

// ParseBlockSubmitted is a log parse operation binding the contract event 0x3d4a04291c66b06f39a4ecb817875b12b5485a05ec563133a56a905305c48e55.
//
// Solidity: event BlockSubmitted(uint256 fork, uint256 epochNumber, uint256 blockNumber, bool isRequest, bool userActivated)
func (_EpochHandler *EpochHandlerFilterer) ParseBlockSubmitted(log types.Log) (*EpochHandlerBlockSubmitted, error) {
	event := new(EpochHandlerBlockSubmitted)
	if err := _EpochHandler.contract.UnpackLog(event, "BlockSubmitted", log); err != nil {
		return nil, err
	}
	return event, nil
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
	TrieKey   []byte
	TrieValue [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterERUCreated is a free log retrieval operation binding the contract event 0xfcbdc2083dadd644b854d91b49aef8db06b8f5a3d5c1192de38ca0ba271d5a0d.
//
// Solidity: event ERUCreated(uint256 requestId, address requestor, address to, bytes trieKey, bytes32 trieValue)
func (_EpochHandler *EpochHandlerFilterer) FilterERUCreated(opts *bind.FilterOpts) (*EpochHandlerERUCreatedIterator, error) {

	logs, sub, err := _EpochHandler.contract.FilterLogs(opts, "ERUCreated")
	if err != nil {
		return nil, err
	}
	return &EpochHandlerERUCreatedIterator{contract: _EpochHandler.contract, event: "ERUCreated", logs: logs, sub: sub}, nil
}

// WatchERUCreated is a free log subscription operation binding the contract event 0xfcbdc2083dadd644b854d91b49aef8db06b8f5a3d5c1192de38ca0ba271d5a0d.
//
// Solidity: event ERUCreated(uint256 requestId, address requestor, address to, bytes trieKey, bytes32 trieValue)
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

// ParseERUCreated is a log parse operation binding the contract event 0xfcbdc2083dadd644b854d91b49aef8db06b8f5a3d5c1192de38ca0ba271d5a0d.
//
// Solidity: event ERUCreated(uint256 requestId, address requestor, address to, bytes trieKey, bytes32 trieValue)
func (_EpochHandler *EpochHandlerFilterer) ParseERUCreated(log types.Log) (*EpochHandlerERUCreated, error) {
	event := new(EpochHandlerERUCreated)
	if err := _EpochHandler.contract.UnpackLog(event, "ERUCreated", log); err != nil {
		return nil, err
	}
	return event, nil
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
// Solidity: event EpochFilled(uint256 forkNumber, uint256 epochNumber)
func (_EpochHandler *EpochHandlerFilterer) FilterEpochFilled(opts *bind.FilterOpts) (*EpochHandlerEpochFilledIterator, error) {

	logs, sub, err := _EpochHandler.contract.FilterLogs(opts, "EpochFilled")
	if err != nil {
		return nil, err
	}
	return &EpochHandlerEpochFilledIterator{contract: _EpochHandler.contract, event: "EpochFilled", logs: logs, sub: sub}, nil
}

// WatchEpochFilled is a free log subscription operation binding the contract event 0x2fdeb407bf5c2b621f04b5c784822dae806c45b49a68aba413cc270128c96816.
//
// Solidity: event EpochFilled(uint256 forkNumber, uint256 epochNumber)
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

// ParseEpochFilled is a log parse operation binding the contract event 0x2fdeb407bf5c2b621f04b5c784822dae806c45b49a68aba413cc270128c96816.
//
// Solidity: event EpochFilled(uint256 forkNumber, uint256 epochNumber)
func (_EpochHandler *EpochHandlerFilterer) ParseEpochFilled(log types.Log) (*EpochHandlerEpochFilled, error) {
	event := new(EpochHandlerEpochFilled)
	if err := _EpochHandler.contract.UnpackLog(event, "EpochFilled", log); err != nil {
		return nil, err
	}
	return event, nil
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
// Solidity: event EpochFilling(uint256 forkNumber, uint256 epochNumber)
func (_EpochHandler *EpochHandlerFilterer) FilterEpochFilling(opts *bind.FilterOpts) (*EpochHandlerEpochFillingIterator, error) {

	logs, sub, err := _EpochHandler.contract.FilterLogs(opts, "EpochFilling")
	if err != nil {
		return nil, err
	}
	return &EpochHandlerEpochFillingIterator{contract: _EpochHandler.contract, event: "EpochFilling", logs: logs, sub: sub}, nil
}

// WatchEpochFilling is a free log subscription operation binding the contract event 0x27b09f0953d27bbff306fe25b2987ac5a813248ac30cb2bbd5daf95e7b0e6dc0.
//
// Solidity: event EpochFilling(uint256 forkNumber, uint256 epochNumber)
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

// ParseEpochFilling is a log parse operation binding the contract event 0x27b09f0953d27bbff306fe25b2987ac5a813248ac30cb2bbd5daf95e7b0e6dc0.
//
// Solidity: event EpochFilling(uint256 forkNumber, uint256 epochNumber)
func (_EpochHandler *EpochHandlerFilterer) ParseEpochFilling(log types.Log) (*EpochHandlerEpochFilling, error) {
	event := new(EpochHandlerEpochFilling)
	if err := _EpochHandler.contract.UnpackLog(event, "EpochFilling", log); err != nil {
		return nil, err
	}
	return event, nil
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
// Solidity: event EpochFinalized(uint256 forkNumber, uint256 epochNumber, uint256 startBlockNumber, uint256 endBlockNumber)
func (_EpochHandler *EpochHandlerFilterer) FilterEpochFinalized(opts *bind.FilterOpts) (*EpochHandlerEpochFinalizedIterator, error) {

	logs, sub, err := _EpochHandler.contract.FilterLogs(opts, "EpochFinalized")
	if err != nil {
		return nil, err
	}
	return &EpochHandlerEpochFinalizedIterator{contract: _EpochHandler.contract, event: "EpochFinalized", logs: logs, sub: sub}, nil
}

// WatchEpochFinalized is a free log subscription operation binding the contract event 0x70801d4d63b3da6c19ba7349911f45bed5a99ccdfb51b8138c105872529bebd5.
//
// Solidity: event EpochFinalized(uint256 forkNumber, uint256 epochNumber, uint256 startBlockNumber, uint256 endBlockNumber)
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

// ParseEpochFinalized is a log parse operation binding the contract event 0x70801d4d63b3da6c19ba7349911f45bed5a99ccdfb51b8138c105872529bebd5.
//
// Solidity: event EpochFinalized(uint256 forkNumber, uint256 epochNumber, uint256 startBlockNumber, uint256 endBlockNumber)
func (_EpochHandler *EpochHandlerFilterer) ParseEpochFinalized(log types.Log) (*EpochHandlerEpochFinalized, error) {
	event := new(EpochHandlerEpochFinalized)
	if err := _EpochHandler.contract.UnpackLog(event, "EpochFinalized", log); err != nil {
		return nil, err
	}
	return event, nil
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
// Solidity: event EpochPrepared(uint256 forkNumber, uint256 epochNumber, uint256 startBlockNumber, uint256 endBlockNumber, uint256 requestStart, uint256 requestEnd, bool epochIsEmpty, bool isRequest, bool userActivated, bool rebase)
func (_EpochHandler *EpochHandlerFilterer) FilterEpochPrepared(opts *bind.FilterOpts) (*EpochHandlerEpochPreparedIterator, error) {

	logs, sub, err := _EpochHandler.contract.FilterLogs(opts, "EpochPrepared")
	if err != nil {
		return nil, err
	}
	return &EpochHandlerEpochPreparedIterator{contract: _EpochHandler.contract, event: "EpochPrepared", logs: logs, sub: sub}, nil
}

// WatchEpochPrepared is a free log subscription operation binding the contract event 0x1a69c0760aa329b76f72579129869013ebd3d41594db019c0e997b939fcb32e3.
//
// Solidity: event EpochPrepared(uint256 forkNumber, uint256 epochNumber, uint256 startBlockNumber, uint256 endBlockNumber, uint256 requestStart, uint256 requestEnd, bool epochIsEmpty, bool isRequest, bool userActivated, bool rebase)
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

// ParseEpochPrepared is a log parse operation binding the contract event 0x1a69c0760aa329b76f72579129869013ebd3d41594db019c0e997b939fcb32e3.
//
// Solidity: event EpochPrepared(uint256 forkNumber, uint256 epochNumber, uint256 startBlockNumber, uint256 endBlockNumber, uint256 requestStart, uint256 requestEnd, bool epochIsEmpty, bool isRequest, bool userActivated, bool rebase)
func (_EpochHandler *EpochHandlerFilterer) ParseEpochPrepared(log types.Log) (*EpochHandlerEpochPrepared, error) {
	event := new(EpochHandlerEpochPrepared)
	if err := _EpochHandler.contract.UnpackLog(event, "EpochPrepared", log); err != nil {
		return nil, err
	}
	return event, nil
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
// Solidity: event EpochRebased(uint256 forkNumber, uint256 epochNumber, uint256 startBlockNumber, uint256 endBlockNumber, uint256 requestStart, uint256 requestEnd, bool epochIsEmpty, bool isRequest, bool userActivated)
func (_EpochHandler *EpochHandlerFilterer) FilterEpochRebased(opts *bind.FilterOpts) (*EpochHandlerEpochRebasedIterator, error) {

	logs, sub, err := _EpochHandler.contract.FilterLogs(opts, "EpochRebased")
	if err != nil {
		return nil, err
	}
	return &EpochHandlerEpochRebasedIterator{contract: _EpochHandler.contract, event: "EpochRebased", logs: logs, sub: sub}, nil
}

// WatchEpochRebased is a free log subscription operation binding the contract event 0x030c1c69405c93021f28f57557240dee939a320b826a1fd0d39bf6e629ecab47.
//
// Solidity: event EpochRebased(uint256 forkNumber, uint256 epochNumber, uint256 startBlockNumber, uint256 endBlockNumber, uint256 requestStart, uint256 requestEnd, bool epochIsEmpty, bool isRequest, bool userActivated)
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

// ParseEpochRebased is a log parse operation binding the contract event 0x030c1c69405c93021f28f57557240dee939a320b826a1fd0d39bf6e629ecab47.
//
// Solidity: event EpochRebased(uint256 forkNumber, uint256 epochNumber, uint256 startBlockNumber, uint256 endBlockNumber, uint256 requestStart, uint256 requestEnd, bool epochIsEmpty, bool isRequest, bool userActivated)
func (_EpochHandler *EpochHandlerFilterer) ParseEpochRebased(log types.Log) (*EpochHandlerEpochRebased, error) {
	event := new(EpochHandlerEpochRebased)
	if err := _EpochHandler.contract.UnpackLog(event, "EpochRebased", log); err != nil {
		return nil, err
	}
	return event, nil
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
// Solidity: event Forked(uint256 newFork, uint256 epochNumber, uint256 forkedBlockNumber)
func (_EpochHandler *EpochHandlerFilterer) FilterForked(opts *bind.FilterOpts) (*EpochHandlerForkedIterator, error) {

	logs, sub, err := _EpochHandler.contract.FilterLogs(opts, "Forked")
	if err != nil {
		return nil, err
	}
	return &EpochHandlerForkedIterator{contract: _EpochHandler.contract, event: "Forked", logs: logs, sub: sub}, nil
}

// WatchForked is a free log subscription operation binding the contract event 0x0647d42ab02f6e0ae76959757dcb6aa6feac1d4ba6f077f1223fb4b1b429f06c.
//
// Solidity: event Forked(uint256 newFork, uint256 epochNumber, uint256 forkedBlockNumber)
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

// ParseForked is a log parse operation binding the contract event 0x0647d42ab02f6e0ae76959757dcb6aa6feac1d4ba6f077f1223fb4b1b429f06c.
//
// Solidity: event Forked(uint256 newFork, uint256 epochNumber, uint256 forkedBlockNumber)
func (_EpochHandler *EpochHandlerFilterer) ParseForked(log types.Log) (*EpochHandlerForked, error) {
	event := new(EpochHandlerForked)
	if err := _EpochHandler.contract.UnpackLog(event, "Forked", log); err != nil {
		return nil, err
	}
	return event, nil
}

// EpochHandlerOperatorChangedIterator is returned from FilterOperatorChanged and is used to iterate over the raw logs and unpacked data for OperatorChanged events raised by the EpochHandler contract.
type EpochHandlerOperatorChangedIterator struct {
	Event *EpochHandlerOperatorChanged // Event containing the contract specifics and raw log

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
func (it *EpochHandlerOperatorChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EpochHandlerOperatorChanged)
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
		it.Event = new(EpochHandlerOperatorChanged)
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
func (it *EpochHandlerOperatorChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EpochHandlerOperatorChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EpochHandlerOperatorChanged represents a OperatorChanged event raised by the EpochHandler contract.
type EpochHandlerOperatorChanged struct {
	NewOperator common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOperatorChanged is a free log retrieval operation binding the contract event 0x4721129e0e676ed6a92909bb24e853ccdd63ad72280cc2e974e38e480e0e6e54.
//
// Solidity: event OperatorChanged(address _newOperator)
func (_EpochHandler *EpochHandlerFilterer) FilterOperatorChanged(opts *bind.FilterOpts) (*EpochHandlerOperatorChangedIterator, error) {

	logs, sub, err := _EpochHandler.contract.FilterLogs(opts, "OperatorChanged")
	if err != nil {
		return nil, err
	}
	return &EpochHandlerOperatorChangedIterator{contract: _EpochHandler.contract, event: "OperatorChanged", logs: logs, sub: sub}, nil
}

// WatchOperatorChanged is a free log subscription operation binding the contract event 0x4721129e0e676ed6a92909bb24e853ccdd63ad72280cc2e974e38e480e0e6e54.
//
// Solidity: event OperatorChanged(address _newOperator)
func (_EpochHandler *EpochHandlerFilterer) WatchOperatorChanged(opts *bind.WatchOpts, sink chan<- *EpochHandlerOperatorChanged) (event.Subscription, error) {

	logs, sub, err := _EpochHandler.contract.WatchLogs(opts, "OperatorChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EpochHandlerOperatorChanged)
				if err := _EpochHandler.contract.UnpackLog(event, "OperatorChanged", log); err != nil {
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

// ParseOperatorChanged is a log parse operation binding the contract event 0x4721129e0e676ed6a92909bb24e853ccdd63ad72280cc2e974e38e480e0e6e54.
//
// Solidity: event OperatorChanged(address _newOperator)
func (_EpochHandler *EpochHandlerFilterer) ParseOperatorChanged(log types.Log) (*EpochHandlerOperatorChanged, error) {
	event := new(EpochHandlerOperatorChanged)
	if err := _EpochHandler.contract.UnpackLog(event, "OperatorChanged", log); err != nil {
		return nil, err
	}
	return event, nil
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
// Solidity: event RequestApplied(uint256 requestId, bool userActivated)
func (_EpochHandler *EpochHandlerFilterer) FilterRequestApplied(opts *bind.FilterOpts) (*EpochHandlerRequestAppliedIterator, error) {

	logs, sub, err := _EpochHandler.contract.FilterLogs(opts, "RequestApplied")
	if err != nil {
		return nil, err
	}
	return &EpochHandlerRequestAppliedIterator{contract: _EpochHandler.contract, event: "RequestApplied", logs: logs, sub: sub}, nil
}

// WatchRequestApplied is a free log subscription operation binding the contract event 0x6940a01870e576ceb735867e13863646d517ce10e66c0133186a4ebdfe9388c2.
//
// Solidity: event RequestApplied(uint256 requestId, bool userActivated)
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

// ParseRequestApplied is a log parse operation binding the contract event 0x6940a01870e576ceb735867e13863646d517ce10e66c0133186a4ebdfe9388c2.
//
// Solidity: event RequestApplied(uint256 requestId, bool userActivated)
func (_EpochHandler *EpochHandlerFilterer) ParseRequestApplied(log types.Log) (*EpochHandlerRequestApplied, error) {
	event := new(EpochHandlerRequestApplied)
	if err := _EpochHandler.contract.UnpackLog(event, "RequestApplied", log); err != nil {
		return nil, err
	}
	return event, nil
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
// Solidity: event RequestChallenged(uint256 requestId, bool userActivated)
func (_EpochHandler *EpochHandlerFilterer) FilterRequestChallenged(opts *bind.FilterOpts) (*EpochHandlerRequestChallengedIterator, error) {

	logs, sub, err := _EpochHandler.contract.FilterLogs(opts, "RequestChallenged")
	if err != nil {
		return nil, err
	}
	return &EpochHandlerRequestChallengedIterator{contract: _EpochHandler.contract, event: "RequestChallenged", logs: logs, sub: sub}, nil
}

// WatchRequestChallenged is a free log subscription operation binding the contract event 0xc8135db115644ed4ae193313c4c801235ef740d2a57a8d5e6fe26ab66635698a.
//
// Solidity: event RequestChallenged(uint256 requestId, bool userActivated)
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

// ParseRequestChallenged is a log parse operation binding the contract event 0xc8135db115644ed4ae193313c4c801235ef740d2a57a8d5e6fe26ab66635698a.
//
// Solidity: event RequestChallenged(uint256 requestId, bool userActivated)
func (_EpochHandler *EpochHandlerFilterer) ParseRequestChallenged(log types.Log) (*EpochHandlerRequestChallenged, error) {
	event := new(EpochHandlerRequestChallenged)
	if err := _EpochHandler.contract.UnpackLog(event, "RequestChallenged", log); err != nil {
		return nil, err
	}
	return event, nil
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
	TrieValue     []byte
	IsExit        bool
	UserActivated bool
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRequestCreated is a free log retrieval operation binding the contract event 0x879922cf5fcada9ebaf8bd7424dc62877f4b220cae07fb6695cc1e8f94c52b4d.
//
// Solidity: event RequestCreated(uint256 requestId, address requestor, address to, uint256 weiAmount, bytes32 trieKey, bytes trieValue, bool isExit, bool userActivated)
func (_EpochHandler *EpochHandlerFilterer) FilterRequestCreated(opts *bind.FilterOpts) (*EpochHandlerRequestCreatedIterator, error) {

	logs, sub, err := _EpochHandler.contract.FilterLogs(opts, "RequestCreated")
	if err != nil {
		return nil, err
	}
	return &EpochHandlerRequestCreatedIterator{contract: _EpochHandler.contract, event: "RequestCreated", logs: logs, sub: sub}, nil
}

// WatchRequestCreated is a free log subscription operation binding the contract event 0x879922cf5fcada9ebaf8bd7424dc62877f4b220cae07fb6695cc1e8f94c52b4d.
//
// Solidity: event RequestCreated(uint256 requestId, address requestor, address to, uint256 weiAmount, bytes32 trieKey, bytes trieValue, bool isExit, bool userActivated)
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

// ParseRequestCreated is a log parse operation binding the contract event 0x879922cf5fcada9ebaf8bd7424dc62877f4b220cae07fb6695cc1e8f94c52b4d.
//
// Solidity: event RequestCreated(uint256 requestId, address requestor, address to, uint256 weiAmount, bytes32 trieKey, bytes trieValue, bool isExit, bool userActivated)
func (_EpochHandler *EpochHandlerFilterer) ParseRequestCreated(log types.Log) (*EpochHandlerRequestCreated, error) {
	event := new(EpochHandlerRequestCreated)
	if err := _EpochHandler.contract.UnpackLog(event, "RequestCreated", log); err != nil {
		return nil, err
	}
	return event, nil
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
// Solidity: event RequestFinalized(uint256 requestId, bool userActivated)
func (_EpochHandler *EpochHandlerFilterer) FilterRequestFinalized(opts *bind.FilterOpts) (*EpochHandlerRequestFinalizedIterator, error) {

	logs, sub, err := _EpochHandler.contract.FilterLogs(opts, "RequestFinalized")
	if err != nil {
		return nil, err
	}
	return &EpochHandlerRequestFinalizedIterator{contract: _EpochHandler.contract, event: "RequestFinalized", logs: logs, sub: sub}, nil
}

// WatchRequestFinalized is a free log subscription operation binding the contract event 0x134017cf3262b18f892ee95dde3b0aec9a80cc70a7c96f09c64bd237aceb0473.
//
// Solidity: event RequestFinalized(uint256 requestId, bool userActivated)
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

// ParseRequestFinalized is a log parse operation binding the contract event 0x134017cf3262b18f892ee95dde3b0aec9a80cc70a7c96f09c64bd237aceb0473.
//
// Solidity: event RequestFinalized(uint256 requestId, bool userActivated)
func (_EpochHandler *EpochHandlerFilterer) ParseRequestFinalized(log types.Log) (*EpochHandlerRequestFinalized, error) {
	event := new(EpochHandlerRequestFinalized)
	if err := _EpochHandler.contract.UnpackLog(event, "RequestFinalized", log); err != nil {
		return nil, err
	}
	return event, nil
}

// EpochHandlerRequestableContractMappedIterator is returned from FilterRequestableContractMapped and is used to iterate over the raw logs and unpacked data for RequestableContractMapped events raised by the EpochHandler contract.
type EpochHandlerRequestableContractMappedIterator struct {
	Event *EpochHandlerRequestableContractMapped // Event containing the contract specifics and raw log

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
func (it *EpochHandlerRequestableContractMappedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EpochHandlerRequestableContractMapped)
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
		it.Event = new(EpochHandlerRequestableContractMapped)
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
func (it *EpochHandlerRequestableContractMappedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EpochHandlerRequestableContractMappedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EpochHandlerRequestableContractMapped represents a RequestableContractMapped event raised by the EpochHandler contract.
type EpochHandlerRequestableContractMapped struct {
	ContractInRootchain  common.Address
	ContractInChildchain common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterRequestableContractMapped is a free log retrieval operation binding the contract event 0xc5ec2ed49686197edd2ed642c7e6096893cc81e6658cde2527030316037715d0.
//
// Solidity: event RequestableContractMapped(address contractInRootchain, address contractInChildchain)
func (_EpochHandler *EpochHandlerFilterer) FilterRequestableContractMapped(opts *bind.FilterOpts) (*EpochHandlerRequestableContractMappedIterator, error) {

	logs, sub, err := _EpochHandler.contract.FilterLogs(opts, "RequestableContractMapped")
	if err != nil {
		return nil, err
	}
	return &EpochHandlerRequestableContractMappedIterator{contract: _EpochHandler.contract, event: "RequestableContractMapped", logs: logs, sub: sub}, nil
}

// WatchRequestableContractMapped is a free log subscription operation binding the contract event 0xc5ec2ed49686197edd2ed642c7e6096893cc81e6658cde2527030316037715d0.
//
// Solidity: event RequestableContractMapped(address contractInRootchain, address contractInChildchain)
func (_EpochHandler *EpochHandlerFilterer) WatchRequestableContractMapped(opts *bind.WatchOpts, sink chan<- *EpochHandlerRequestableContractMapped) (event.Subscription, error) {

	logs, sub, err := _EpochHandler.contract.WatchLogs(opts, "RequestableContractMapped")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EpochHandlerRequestableContractMapped)
				if err := _EpochHandler.contract.UnpackLog(event, "RequestableContractMapped", log); err != nil {
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
func (_EpochHandler *EpochHandlerFilterer) ParseRequestableContractMapped(log types.Log) (*EpochHandlerRequestableContractMapped, error) {
	event := new(EpochHandlerRequestableContractMapped)
	if err := _EpochHandler.contract.UnpackLog(event, "RequestableContractMapped", log); err != nil {
		return nil, err
	}
	return event, nil
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
// Solidity: event SessionTimeout(bool userActivated)
func (_EpochHandler *EpochHandlerFilterer) FilterSessionTimeout(opts *bind.FilterOpts) (*EpochHandlerSessionTimeoutIterator, error) {

	logs, sub, err := _EpochHandler.contract.FilterLogs(opts, "SessionTimeout")
	if err != nil {
		return nil, err
	}
	return &EpochHandlerSessionTimeoutIterator{contract: _EpochHandler.contract, event: "SessionTimeout", logs: logs, sub: sub}, nil
}

// WatchSessionTimeout is a free log subscription operation binding the contract event 0x2122ec719581cd177f225f59a1ee005831211196831f8f1ccffa817d2e7bd108.
//
// Solidity: event SessionTimeout(bool userActivated)
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

// ParseSessionTimeout is a log parse operation binding the contract event 0x2122ec719581cd177f225f59a1ee005831211196831f8f1ccffa817d2e7bd108.
//
// Solidity: event SessionTimeout(bool userActivated)
func (_EpochHandler *EpochHandlerFilterer) ParseSessionTimeout(log types.Log) (*EpochHandlerSessionTimeout, error) {
	event := new(EpochHandlerSessionTimeout)
	if err := _EpochHandler.contract.UnpackLog(event, "SessionTimeout", log); err != nil {
		return nil, err
	}
	return event, nil
}

// MathABI is the input ABI used to generate the binding from.
const MathABI = "[]"

// MathBin is the compiled bytecode used for deploying new contracts.
var MathBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a7231582059f5b6656555777713940676f444002442a3eab637cd25b4c9bb08f94083d96b64736f6c634300050c0032"

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
var RLPBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a723158200e62f7dc00fc4c8eac04f8d672dc03a017efe72ff7483585fb6978cb1b070d1264736f6c634300050c0032"

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
var RLPEncodeBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a723158205da3446f8619339802867bf627e0a042d4f4176598fe1e096c50366e6bfe24a164736f6c634300050c0032"

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

// RootChainEventABI is the input ABI used to generate the binding from.
const RootChainEventABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"forkNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"BlockFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fork\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epochNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isRequest\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"userActivated\",\"type\":\"bool\"}],\"name\":\"BlockSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"requestor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"trieKey\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"trieValue\",\"type\":\"bytes32\"}],\"name\":\"ERUCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"forkNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epochNumber\",\"type\":\"uint256\"}],\"name\":\"EpochFilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"forkNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epochNumber\",\"type\":\"uint256\"}],\"name\":\"EpochFilling\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"forkNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epochNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startBlockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"endBlockNumber\",\"type\":\"uint256\"}],\"name\":\"EpochFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"forkNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epochNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startBlockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"endBlockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requestStart\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requestEnd\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"epochIsEmpty\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isRequest\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"userActivated\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"rebase\",\"type\":\"bool\"}],\"name\":\"EpochPrepared\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"forkNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epochNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startBlockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"endBlockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requestStart\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requestEnd\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"epochIsEmpty\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isRequest\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"userActivated\",\"type\":\"bool\"}],\"name\":\"EpochRebased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newFork\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epochNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"forkedBlockNumber\",\"type\":\"uint256\"}],\"name\":\"Forked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_newOperator\",\"type\":\"address\"}],\"name\":\"OperatorChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"userActivated\",\"type\":\"bool\"}],\"name\":\"RequestApplied\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"userActivated\",\"type\":\"bool\"}],\"name\":\"RequestChallenged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"requestor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"weiAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"trieKey\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"trieValue\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isExit\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"userActivated\",\"type\":\"bool\"}],\"name\":\"RequestCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"userActivated\",\"type\":\"bool\"}],\"name\":\"RequestFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"contractInRootchain\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"contractInChildchain\",\"type\":\"address\"}],\"name\":\"RequestableContractMapped\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"userActivated\",\"type\":\"bool\"}],\"name\":\"SessionTimeout\",\"type\":\"event\"}]"

// RootChainEventBin is the compiled bytecode used for deploying new contracts.
var RootChainEventBin = "0x6080604052348015600f57600080fd5b50603e80601d6000396000f3fe6080604052600080fdfea265627a7a723158209eca3e6c0985f0e87007d8dee1de52ecce4b4b712a00ee740621f2a988b3d29764736f6c634300050c0032"

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

// RootChainEventOperatorChangedIterator is returned from FilterOperatorChanged and is used to iterate over the raw logs and unpacked data for OperatorChanged events raised by the RootChainEvent contract.
type RootChainEventOperatorChangedIterator struct {
	Event *RootChainEventOperatorChanged // Event containing the contract specifics and raw log

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
func (it *RootChainEventOperatorChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RootChainEventOperatorChanged)
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
		it.Event = new(RootChainEventOperatorChanged)
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
func (it *RootChainEventOperatorChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RootChainEventOperatorChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RootChainEventOperatorChanged represents a OperatorChanged event raised by the RootChainEvent contract.
type RootChainEventOperatorChanged struct {
	NewOperator common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOperatorChanged is a free log retrieval operation binding the contract event 0x4721129e0e676ed6a92909bb24e853ccdd63ad72280cc2e974e38e480e0e6e54.
//
// Solidity: event OperatorChanged(address _newOperator)
func (_RootChainEvent *RootChainEventFilterer) FilterOperatorChanged(opts *bind.FilterOpts) (*RootChainEventOperatorChangedIterator, error) {

	logs, sub, err := _RootChainEvent.contract.FilterLogs(opts, "OperatorChanged")
	if err != nil {
		return nil, err
	}
	return &RootChainEventOperatorChangedIterator{contract: _RootChainEvent.contract, event: "OperatorChanged", logs: logs, sub: sub}, nil
}

// WatchOperatorChanged is a free log subscription operation binding the contract event 0x4721129e0e676ed6a92909bb24e853ccdd63ad72280cc2e974e38e480e0e6e54.
//
// Solidity: event OperatorChanged(address _newOperator)
func (_RootChainEvent *RootChainEventFilterer) WatchOperatorChanged(opts *bind.WatchOpts, sink chan<- *RootChainEventOperatorChanged) (event.Subscription, error) {

	logs, sub, err := _RootChainEvent.contract.WatchLogs(opts, "OperatorChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RootChainEventOperatorChanged)
				if err := _RootChainEvent.contract.UnpackLog(event, "OperatorChanged", log); err != nil {
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

// ParseOperatorChanged is a log parse operation binding the contract event 0x4721129e0e676ed6a92909bb24e853ccdd63ad72280cc2e974e38e480e0e6e54.
//
// Solidity: event OperatorChanged(address _newOperator)
func (_RootChainEvent *RootChainEventFilterer) ParseOperatorChanged(log types.Log) (*RootChainEventOperatorChanged, error) {
	event := new(RootChainEventOperatorChanged)
	if err := _RootChainEvent.contract.UnpackLog(event, "OperatorChanged", log); err != nil {
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
const RootChainStorageABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"COST_ERO\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"COST_ERU\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"COST_NRB\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"COST_ORB\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"COST_URB\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"COST_URB_PREPARE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CP_COMPUTATION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CP_EXIT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CP_WITHHOLDING\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"EROIdToFinalize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"EROs\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"isExit\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isTransfer\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"finalized\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"challenged\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"value\",\"type\":\"uint128\"},{\"internalType\":\"addresspayable\",\"name\":\"requestor\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"trieKey\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"trieValue\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ERUIdToFinalize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"ERUs\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"isExit\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isTransfer\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"finalized\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"challenged\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"value\",\"type\":\"uint128\"},{\"internalType\":\"addresspayable\",\"name\":\"requestor\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"trieKey\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"trieValue\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"NRELength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"NULL_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"ORBs\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"submitted\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"numEnter\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"epochNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"requestStart\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"requestEnd\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"trie\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PREPARE_TIMEOUT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"REQUEST_GAS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"URBs\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"submitted\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"numEnter\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"epochNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"requestStart\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"requestEnd\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"trie\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"currentFork\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"development\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"epochHandler\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"etherToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"firstFilledORENumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"firstNonEmptyRequestEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"forks\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"forkedBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"firstEpoch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"lastEpoch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"firstBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"lastBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"lastFinalizedEpoch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"lastFinalizedBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"firstEnterEpoch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"lastEnterEpoch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nextBlockToRebase\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"rebased\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isRootChain\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastAppliedBlockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastAppliedEpochNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastAppliedForkNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"lastNonEmptyRequestEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"numEnterForORB\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"operator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"requestableContracts\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"seigManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"submitHandler\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

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
	"420bb4b8": "isRootChain()",
	"fb788a27": "lastAppliedBlockNumber()",
	"c8ad329f": "lastAppliedEpochNumber()",
	"164bc2ae": "lastAppliedForkNumber()",
	"b6715647": "lastNonEmptyRequestEpoch(uint256)",
	"23691566": "numEnterForORB()",
	"570ca735": "operator()",
	"da0185f8": "requestableContracts(address)",
	"6fb7f558": "seigManager()",
	"e259faf7": "submitHandler()",
}

// RootChainStorageBin is the compiled bytecode used for deploying new contracts.
var RootChainStorageBin = "0x608060405234801561001057600080fd5b50610955806100206000396000f3fe608060405234801561001057600080fd5b506004361061021c5760003560e01c8063b17fa6e911610125578063ca6f6380116100ad578063e259faf71161007c578063e259faf714610595578063e7b88b801461059d578063ea7f22a8146105a5578063f4f31de4146105c2578063fb788a27146105df5761021c565b8063ca6f63801461054a578063d691acd814610221578063da0185f814610567578063de0ce17d1461058d5761021c565b8063b8066bcb116100f4578063b8066bcb146104c1578063c0e86064146104c9578063c2bc88fa14610532578063c54626cc1461053a578063c8ad329f146105425761021c565b8063b17fa6e91461038c578063b2ae9ba814610221578063b443f3cc14610394578063b6715647146104a45761021c565b8063570ca735116101a85780638155717d116101775780638155717d1461036c5780638b5172d0146103745780638eb288ca1461037c57806394be3aa514610221578063ab96da2d146103845761021c565b8063570ca7351461031b5780636fb7f5581461033f57806372ecb9a8146103475780637b929c27146103645761021c565b8063192adc5b116101ef578063192adc5b14610253578063236915661461025b5780632dc6bb7b14610263578063420bb4b81461026b5780634ba3a126146102875761021c565b8063033cfbed1461022157806308c4fff01461023b578063164bc2ae14610243578063183d2d1c1461024b575b600080fd5b6102296105e7565b60408051918252519081900360200190f35b6102296105f1565b6102296105f6565b6102296105fc565b610229610602565b61022961060c565b610229610612565b610273610618565b604080519115158252519081900360200190f35b6102a46004803603602081101561029d57600080fd5b503561061d565b6040805167ffffffffffffffff9d8e1681529b8d1660208d0152998c168b8b0152978b1660608b0152958a1660808a015293891660a089015291881660c0880152871660e0870152861661010086015285166101208501529093166101408301529115156101608201529051908190036101800190f35b610323610695565b604080516001600160a01b039092168252519081900360200190f35b6103236106a9565b6102296004803603602081101561035d57600080fd5b50356106b8565b6102736106ca565b6102296106d3565b6102296106d8565b6102296106e2565b6102296106e9565b6102296106ef565b6103b1600480360360208110156103aa57600080fd5b50356106f4565b6040805167ffffffffffffffff8d1681528b15156020808301919091528b151592820192909252891515606082015288151560808201526001600160801b03881660a08201526001600160a01b0380881660c0830152861660e082015261010081018590526101208101849052610160610140820181815284519183019190915283519192909161018084019185019080838360005b8381101561045f578181015183820152602001610447565b50505050905090810190601f16801561048c5780820380516001836020036101000a031916815260200191505b509c5050505050505050505050505060405180910390f35b610229600480360360208110156104ba57600080fd5b5035610816565b610323610828565b6104e6600480360360208110156104df57600080fd5b5035610837565b60408051961515875267ffffffffffffffff958616602088015293851686850152918416606086015290921660808401526001600160a01b0390911660a0830152519081900360c00190f35b61022961089f565b6102296108a4565b6102296108aa565b6102296004803603602081101561056057600080fd5b50356108b0565b6103236004803603602081101561057d57600080fd5b50356001600160a01b03166108c2565b6103236108dd565b6103236108e2565b6103236108f1565b6104e6600480360360208110156105bb57600080fd5b5035610900565b6103b1600480360360208110156105d857600080fd5b503561090d565b61022961091a565b6509184e72a00081565b600f81565b60105481565b60065481565b6551dac207a00081565b600d5481565b60135481565b600181565b60086020526000908152604090208054600182015460029092015467ffffffffffffffff80831693600160401b808504831694600160801b808204851695600160c01b92839004861695858116958581048216958482048316959182900483169484841694918204841693908204169160ff9104168c565b60005461010090046001600160a01b031681565b6004546001600160a01b031681565b60076020526000908152604090205481565b60005460ff1681565b600a81565b6512309ce5400081565b620186a081565b60055481565b601481565b6009818154811061070157fe5b6000918252602091829020600691909102018054600180830154600280850154600386015460048701546005880180546040805161010099831615999099026000190190911695909504601f81018b90048b0288018b0190955284875267ffffffffffffffff88169a50600160401b880460ff9081169a600160481b8a0482169a600160501b8b0483169a600160581b810490931699600160601b9093046001600160801b0316986001600160a01b039081169897169690939183018282801561080c5780601f106107e15761010080835404028352916020019161080c565b820191906000526020600020905b8154815290600101906020018083116107ef57829003601f168201915b505050505090508b565b600e6020526000908152604090205481565b6003546001600160a01b031681565b600c818154811061084457fe5b60009182526020909120600290910201805460019091015460ff8216925067ffffffffffffffff6101008304811692600160481b8104821692600160881b9091048216918116906001600160a01b03600160401b9091041686565b603c81565b60145481565b60115481565b600f6020526000908152604090205481565b6015602052600090815260409020546001600160a01b031681565b600081565b6002546001600160a01b031681565b6001546001600160a01b031681565b600b818154811061084457fe5b600a818154811061070157fe5b6012548156fea265627a7a723158208658500c8e7688e6b2b532a144651710b608708a6e1f8775e1e146893a74b19164736f6c634300050c0032"

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

// IsRootChain is a free data retrieval call binding the contract method 0x420bb4b8.
//
// Solidity: function isRootChain() constant returns(bool)
func (_RootChainStorage *RootChainStorageCaller) IsRootChain(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _RootChainStorage.contract.Call(opts, out, "isRootChain")
	return *ret0, err
}

// IsRootChain is a free data retrieval call binding the contract method 0x420bb4b8.
//
// Solidity: function isRootChain() constant returns(bool)
func (_RootChainStorage *RootChainStorageSession) IsRootChain() (bool, error) {
	return _RootChainStorage.Contract.IsRootChain(&_RootChainStorage.CallOpts)
}

// IsRootChain is a free data retrieval call binding the contract method 0x420bb4b8.
//
// Solidity: function isRootChain() constant returns(bool)
func (_RootChainStorage *RootChainStorageCallerSession) IsRootChain() (bool, error) {
	return _RootChainStorage.Contract.IsRootChain(&_RootChainStorage.CallOpts)
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

// SeigManager is a free data retrieval call binding the contract method 0x6fb7f558.
//
// Solidity: function seigManager() constant returns(address)
func (_RootChainStorage *RootChainStorageCaller) SeigManager(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RootChainStorage.contract.Call(opts, out, "seigManager")
	return *ret0, err
}

// SeigManager is a free data retrieval call binding the contract method 0x6fb7f558.
//
// Solidity: function seigManager() constant returns(address)
func (_RootChainStorage *RootChainStorageSession) SeigManager() (common.Address, error) {
	return _RootChainStorage.Contract.SeigManager(&_RootChainStorage.CallOpts)
}

// SeigManager is a free data retrieval call binding the contract method 0x6fb7f558.
//
// Solidity: function seigManager() constant returns(address)
func (_RootChainStorage *RootChainStorageCallerSession) SeigManager() (common.Address, error) {
	return _RootChainStorage.Contract.SeigManager(&_RootChainStorage.CallOpts)
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
var SafeMathBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a7231582081abea5b57eb862cc535e559557e86893bc93410acfd1b3122d48a2c0cdee52a64736f6c634300050c0032"

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
