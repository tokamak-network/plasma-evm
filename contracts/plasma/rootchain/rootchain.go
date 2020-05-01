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

// DataEpoch is an auto generated low-level Go binding around an user-defined struct.
type DataEpoch struct {
	StartBlockNumber uint64
	EndBlockNumber   uint64
	Timestamp        uint64
	IsEmpty          bool
	Initialized      bool
	IsRequest        bool
	UserActivated    bool
	Rebase           bool
	RE               DataRequestEpochMeta
	NRE              DataNonRequestEpochMeta
}

// DataNonRequestEpochMeta is an auto generated low-level Go binding around an user-defined struct.
type DataNonRequestEpochMeta struct {
	EpochStateRoot        [32]byte
	EpochTransactionsRoot [32]byte
	EpochReceiptsRoot     [32]byte
	SubmittedAt           uint64
	FinalizedAt           uint64
	Finalized             bool
	Challenging           bool
	Challenged            bool
}

// DataPlasmaBlock is an auto generated low-level Go binding around an user-defined struct.
type DataPlasmaBlock struct {
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

// DataRequestEpochMeta is an auto generated low-level Go binding around an user-defined struct.
type DataRequestEpochMeta struct {
	RequestStart        uint64
	RequestEnd          uint64
	FirstRequestBlockId uint64
	NumEnter            uint64
	NextEnterEpoch      uint64
	NextEpoch           uint64
}

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

// MapperRoleABI is the input ABI used to generate the binding from.
const MapperRoleABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"MapperAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"MapperRemoved\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addMapper\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isMapper\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceMapper\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// MapperRoleFuncSigs maps the 4-byte function signature to its string representation.
var MapperRoleFuncSigs = map[string]string{
	"80e3e81d": "addMapper(address)",
	"c0c49c2a": "isMapper(address)",
	"3565fb0d": "renounceMapper()",
}

// MapperRole is an auto generated Go binding around an Ethereum contract.
type MapperRole struct {
	MapperRoleCaller     // Read-only binding to the contract
	MapperRoleTransactor // Write-only binding to the contract
	MapperRoleFilterer   // Log filterer for contract events
}

// MapperRoleCaller is an auto generated read-only Go binding around an Ethereum contract.
type MapperRoleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MapperRoleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MapperRoleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MapperRoleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MapperRoleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MapperRoleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MapperRoleSession struct {
	Contract     *MapperRole       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MapperRoleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MapperRoleCallerSession struct {
	Contract *MapperRoleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// MapperRoleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MapperRoleTransactorSession struct {
	Contract     *MapperRoleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// MapperRoleRaw is an auto generated low-level Go binding around an Ethereum contract.
type MapperRoleRaw struct {
	Contract *MapperRole // Generic contract binding to access the raw methods on
}

// MapperRoleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MapperRoleCallerRaw struct {
	Contract *MapperRoleCaller // Generic read-only contract binding to access the raw methods on
}

// MapperRoleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MapperRoleTransactorRaw struct {
	Contract *MapperRoleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMapperRole creates a new instance of MapperRole, bound to a specific deployed contract.
func NewMapperRole(address common.Address, backend bind.ContractBackend) (*MapperRole, error) {
	contract, err := bindMapperRole(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MapperRole{MapperRoleCaller: MapperRoleCaller{contract: contract}, MapperRoleTransactor: MapperRoleTransactor{contract: contract}, MapperRoleFilterer: MapperRoleFilterer{contract: contract}}, nil
}

// NewMapperRoleCaller creates a new read-only instance of MapperRole, bound to a specific deployed contract.
func NewMapperRoleCaller(address common.Address, caller bind.ContractCaller) (*MapperRoleCaller, error) {
	contract, err := bindMapperRole(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MapperRoleCaller{contract: contract}, nil
}

// NewMapperRoleTransactor creates a new write-only instance of MapperRole, bound to a specific deployed contract.
func NewMapperRoleTransactor(address common.Address, transactor bind.ContractTransactor) (*MapperRoleTransactor, error) {
	contract, err := bindMapperRole(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MapperRoleTransactor{contract: contract}, nil
}

// NewMapperRoleFilterer creates a new log filterer instance of MapperRole, bound to a specific deployed contract.
func NewMapperRoleFilterer(address common.Address, filterer bind.ContractFilterer) (*MapperRoleFilterer, error) {
	contract, err := bindMapperRole(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MapperRoleFilterer{contract: contract}, nil
}

// bindMapperRole binds a generic wrapper to an already deployed contract.
func bindMapperRole(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MapperRoleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MapperRole *MapperRoleRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MapperRole.Contract.MapperRoleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MapperRole *MapperRoleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MapperRole.Contract.MapperRoleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MapperRole *MapperRoleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MapperRole.Contract.MapperRoleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MapperRole *MapperRoleCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MapperRole.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MapperRole *MapperRoleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MapperRole.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MapperRole *MapperRoleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MapperRole.Contract.contract.Transact(opts, method, params...)
}

// IsMapper is a free data retrieval call binding the contract method 0xc0c49c2a.
//
// Solidity: function isMapper(address account) constant returns(bool)
func (_MapperRole *MapperRoleCaller) IsMapper(opts *bind.CallOpts, account common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _MapperRole.contract.Call(opts, out, "isMapper", account)
	return *ret0, err
}

// IsMapper is a free data retrieval call binding the contract method 0xc0c49c2a.
//
// Solidity: function isMapper(address account) constant returns(bool)
func (_MapperRole *MapperRoleSession) IsMapper(account common.Address) (bool, error) {
	return _MapperRole.Contract.IsMapper(&_MapperRole.CallOpts, account)
}

// IsMapper is a free data retrieval call binding the contract method 0xc0c49c2a.
//
// Solidity: function isMapper(address account) constant returns(bool)
func (_MapperRole *MapperRoleCallerSession) IsMapper(account common.Address) (bool, error) {
	return _MapperRole.Contract.IsMapper(&_MapperRole.CallOpts, account)
}

// AddMapper is a paid mutator transaction binding the contract method 0x80e3e81d.
//
// Solidity: function addMapper(address account) returns()
func (_MapperRole *MapperRoleTransactor) AddMapper(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _MapperRole.contract.Transact(opts, "addMapper", account)
}

// AddMapper is a paid mutator transaction binding the contract method 0x80e3e81d.
//
// Solidity: function addMapper(address account) returns()
func (_MapperRole *MapperRoleSession) AddMapper(account common.Address) (*types.Transaction, error) {
	return _MapperRole.Contract.AddMapper(&_MapperRole.TransactOpts, account)
}

// AddMapper is a paid mutator transaction binding the contract method 0x80e3e81d.
//
// Solidity: function addMapper(address account) returns()
func (_MapperRole *MapperRoleTransactorSession) AddMapper(account common.Address) (*types.Transaction, error) {
	return _MapperRole.Contract.AddMapper(&_MapperRole.TransactOpts, account)
}

// RenounceMapper is a paid mutator transaction binding the contract method 0x3565fb0d.
//
// Solidity: function renounceMapper() returns()
func (_MapperRole *MapperRoleTransactor) RenounceMapper(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MapperRole.contract.Transact(opts, "renounceMapper")
}

// RenounceMapper is a paid mutator transaction binding the contract method 0x3565fb0d.
//
// Solidity: function renounceMapper() returns()
func (_MapperRole *MapperRoleSession) RenounceMapper() (*types.Transaction, error) {
	return _MapperRole.Contract.RenounceMapper(&_MapperRole.TransactOpts)
}

// RenounceMapper is a paid mutator transaction binding the contract method 0x3565fb0d.
//
// Solidity: function renounceMapper() returns()
func (_MapperRole *MapperRoleTransactorSession) RenounceMapper() (*types.Transaction, error) {
	return _MapperRole.Contract.RenounceMapper(&_MapperRole.TransactOpts)
}

// MapperRoleMapperAddedIterator is returned from FilterMapperAdded and is used to iterate over the raw logs and unpacked data for MapperAdded events raised by the MapperRole contract.
type MapperRoleMapperAddedIterator struct {
	Event *MapperRoleMapperAdded // Event containing the contract specifics and raw log

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
func (it *MapperRoleMapperAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MapperRoleMapperAdded)
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
		it.Event = new(MapperRoleMapperAdded)
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
func (it *MapperRoleMapperAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MapperRoleMapperAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MapperRoleMapperAdded represents a MapperAdded event raised by the MapperRole contract.
type MapperRoleMapperAdded struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMapperAdded is a free log retrieval operation binding the contract event 0xa199526a01dbcba68aa3ab9d7d06c6692c83fb1a8bcf2184e6d94ad34f5aaf95.
//
// Solidity: event MapperAdded(address indexed account)
func (_MapperRole *MapperRoleFilterer) FilterMapperAdded(opts *bind.FilterOpts, account []common.Address) (*MapperRoleMapperAddedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _MapperRole.contract.FilterLogs(opts, "MapperAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return &MapperRoleMapperAddedIterator{contract: _MapperRole.contract, event: "MapperAdded", logs: logs, sub: sub}, nil
}

// WatchMapperAdded is a free log subscription operation binding the contract event 0xa199526a01dbcba68aa3ab9d7d06c6692c83fb1a8bcf2184e6d94ad34f5aaf95.
//
// Solidity: event MapperAdded(address indexed account)
func (_MapperRole *MapperRoleFilterer) WatchMapperAdded(opts *bind.WatchOpts, sink chan<- *MapperRoleMapperAdded, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _MapperRole.contract.WatchLogs(opts, "MapperAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MapperRoleMapperAdded)
				if err := _MapperRole.contract.UnpackLog(event, "MapperAdded", log); err != nil {
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

// ParseMapperAdded is a log parse operation binding the contract event 0xa199526a01dbcba68aa3ab9d7d06c6692c83fb1a8bcf2184e6d94ad34f5aaf95.
//
// Solidity: event MapperAdded(address indexed account)
func (_MapperRole *MapperRoleFilterer) ParseMapperAdded(log types.Log) (*MapperRoleMapperAdded, error) {
	event := new(MapperRoleMapperAdded)
	if err := _MapperRole.contract.UnpackLog(event, "MapperAdded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// MapperRoleMapperRemovedIterator is returned from FilterMapperRemoved and is used to iterate over the raw logs and unpacked data for MapperRemoved events raised by the MapperRole contract.
type MapperRoleMapperRemovedIterator struct {
	Event *MapperRoleMapperRemoved // Event containing the contract specifics and raw log

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
func (it *MapperRoleMapperRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MapperRoleMapperRemoved)
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
		it.Event = new(MapperRoleMapperRemoved)
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
func (it *MapperRoleMapperRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MapperRoleMapperRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MapperRoleMapperRemoved represents a MapperRemoved event raised by the MapperRole contract.
type MapperRoleMapperRemoved struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMapperRemoved is a free log retrieval operation binding the contract event 0xf012269e1b6bb33c7800e7e54167ddd74f0ae7244108030b6876821d3779822f.
//
// Solidity: event MapperRemoved(address indexed account)
func (_MapperRole *MapperRoleFilterer) FilterMapperRemoved(opts *bind.FilterOpts, account []common.Address) (*MapperRoleMapperRemovedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _MapperRole.contract.FilterLogs(opts, "MapperRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return &MapperRoleMapperRemovedIterator{contract: _MapperRole.contract, event: "MapperRemoved", logs: logs, sub: sub}, nil
}

// WatchMapperRemoved is a free log subscription operation binding the contract event 0xf012269e1b6bb33c7800e7e54167ddd74f0ae7244108030b6876821d3779822f.
//
// Solidity: event MapperRemoved(address indexed account)
func (_MapperRole *MapperRoleFilterer) WatchMapperRemoved(opts *bind.WatchOpts, sink chan<- *MapperRoleMapperRemoved, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _MapperRole.contract.WatchLogs(opts, "MapperRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MapperRoleMapperRemoved)
				if err := _MapperRole.contract.UnpackLog(event, "MapperRemoved", log); err != nil {
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

// ParseMapperRemoved is a log parse operation binding the contract event 0xf012269e1b6bb33c7800e7e54167ddd74f0ae7244108030b6876821d3779822f.
//
// Solidity: event MapperRemoved(address indexed account)
func (_MapperRole *MapperRoleFilterer) ParseMapperRemoved(log types.Log) (*MapperRoleMapperRemoved, error) {
	event := new(MapperRoleMapperRemoved)
	if err := _MapperRole.contract.UnpackLog(event, "MapperRemoved", log); err != nil {
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

// RolesABI is the input ABI used to generate the binding from.
const RolesABI = "[]"

// RolesBin is the compiled bytecode used for deploying new contracts.
var RolesBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a723158201aa75146c2e109912d51a929e368d1efc36fee31e25f784cae25ad0120ec320564736f6c634300050c0032"

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

// RootChainABI is the input ABI used to generate the binding from.
const RootChainABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_epochHandler\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_submitHandler\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_etherToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_development\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"_NRELength\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_statesRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_transactionsRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_receiptsRoot\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"forkNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"BlockFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fork\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epochNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isRequest\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"userActivated\",\"type\":\"bool\"}],\"name\":\"BlockSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"requestor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"trieKey\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"trieValue\",\"type\":\"bytes32\"}],\"name\":\"ERUCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"forkNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epochNumber\",\"type\":\"uint256\"}],\"name\":\"EpochFilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"forkNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epochNumber\",\"type\":\"uint256\"}],\"name\":\"EpochFilling\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"forkNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epochNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startBlockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"endBlockNumber\",\"type\":\"uint256\"}],\"name\":\"EpochFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"forkNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epochNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startBlockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"endBlockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requestStart\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requestEnd\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"epochIsEmpty\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isRequest\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"userActivated\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"rebase\",\"type\":\"bool\"}],\"name\":\"EpochPrepared\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"forkNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epochNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startBlockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"endBlockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requestStart\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requestEnd\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"epochIsEmpty\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isRequest\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"userActivated\",\"type\":\"bool\"}],\"name\":\"EpochRebased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newFork\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epochNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"forkedBlockNumber\",\"type\":\"uint256\"}],\"name\":\"Forked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"MapperAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"MapperRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_newOperator\",\"type\":\"address\"}],\"name\":\"OperatorChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"userActivated\",\"type\":\"bool\"}],\"name\":\"RequestApplied\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"userActivated\",\"type\":\"bool\"}],\"name\":\"RequestChallenged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"requestor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"weiAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"trieKey\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"trieValue\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isExit\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"userActivated\",\"type\":\"bool\"}],\"name\":\"RequestCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"userActivated\",\"type\":\"bool\"}],\"name\":\"RequestFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"contractInRootchain\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"contractInChildchain\",\"type\":\"address\"}],\"name\":\"RequestableContractMapped\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"userActivated\",\"type\":\"bool\"}],\"name\":\"SessionTimeout\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"SubmitterAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"SubmitterRemoved\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"COST_ERO\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"COST_ERU\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"COST_NRB\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"COST_ORB\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"COST_URB\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"COST_URB_PREPARE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CP_COMPUTATION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CP_EXIT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CP_WITHHOLDING\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"EROIdToFinalize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"EROs\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"isExit\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isTransfer\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"finalized\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"challenged\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"value\",\"type\":\"uint128\"},{\"internalType\":\"addresspayable\",\"name\":\"requestor\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"trieKey\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"trieValue\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ERUIdToFinalize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"ERUs\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"isExit\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isTransfer\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"finalized\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"challenged\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"value\",\"type\":\"uint128\"},{\"internalType\":\"addresspayable\",\"name\":\"requestor\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"trieKey\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"trieValue\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAX_REQUESTS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"maxRequests\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"NRELength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"NULL_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"ORBs\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"submitted\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"numEnter\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"epochNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"requestStart\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"requestEnd\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"trie\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PREPARE_TIMEOUT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"REQUEST_GAS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"URBs\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"submitted\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"numEnter\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"epochNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"requestStart\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"requestEnd\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"trie\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addMapper\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addSubmitter\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_forkNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_receiptData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_proof\",\"type\":\"bytes\"}],\"name\":\"challengeExit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_key\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_txByte\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_branchMask\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"_siblings\",\"type\":\"bytes32[]\"}],\"name\":\"challengeNullAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"changeOperator\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"currentFork\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"development\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"epochHandler\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"etherToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"finalizeBlock\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"finalizeRequest\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"n\",\"type\":\"uint256\"}],\"name\":\"finalizeRequests\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"firstFilledORENumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"firstNonEmptyRequestEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_forkNumber\",\"type\":\"uint256\"}],\"name\":\"forked\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"result\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"forks\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"forkedBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"firstEpoch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"lastEpoch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"firstBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"lastBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"lastFinalizedEpoch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"lastFinalizedBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"firstEnterEpoch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"lastEnterEpoch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nextBlockToRebase\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"rebased\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"forkNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"getBlock\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"epochNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"requestBlockId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"finalizedAt\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"referenceBlock\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"statesRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"transactionsRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"receiptsRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"isRequest\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"userActivated\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"challenged\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"challenging\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"finalized\",\"type\":\"bool\"}],\"internalType\":\"structData.PlasmaBlock\",\"name\":\"\",\"type\":\"tuple\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"forkNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"getBlockFinalizedAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"}],\"name\":\"getEROBytes\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"out\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"forkNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochNumber\",\"type\":\"uint256\"}],\"name\":\"getEpoch\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"startBlockNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"endBlockNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"isEmpty\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"initialized\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isRequest\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"userActivated\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"rebase\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"requestStart\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"requestEnd\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"firstRequestBlockId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"numEnter\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nextEnterEpoch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nextEpoch\",\"type\":\"uint64\"}],\"internalType\":\"structData.RequestEpochMeta\",\"name\":\"RE\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"epochStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"epochTransactionsRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"epochReceiptsRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"submittedAt\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"finalizedAt\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"finalized\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"challenging\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"challenged\",\"type\":\"bool\"}],\"internalType\":\"structData.NonRequestEpochMeta\",\"name\":\"NRE\",\"type\":\"tuple\"}],\"internalType\":\"structData.Epoch\",\"name\":\"epoch\",\"type\":\"tuple\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getLastEpoch\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"startBlockNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"endBlockNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"isEmpty\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"initialized\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isRequest\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"userActivated\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"rebase\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"requestStart\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"requestEnd\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"firstRequestBlockId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"numEnter\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nextEnterEpoch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nextEpoch\",\"type\":\"uint64\"}],\"internalType\":\"structData.RequestEpochMeta\",\"name\":\"RE\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"epochStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"epochTransactionsRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"epochReceiptsRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"submittedAt\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"finalizedAt\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"finalized\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"challenging\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"challenged\",\"type\":\"bool\"}],\"internalType\":\"structData.NonRequestEpochMeta\",\"name\":\"NRE\",\"type\":\"tuple\"}],\"internalType\":\"structData.Epoch\",\"name\":\"\",\"type\":\"tuple\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"forkNumber\",\"type\":\"uint256\"}],\"name\":\"getLastFinalizedBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"forkNumber\",\"type\":\"uint256\"}],\"name\":\"getLastFinalizedEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getNumEROs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getNumORBs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_userActivated\",\"type\":\"bool\"}],\"name\":\"getRequestFinalized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"finalized\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isMapper\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isRootChain\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isSubmitter\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastAppliedBlockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastAppliedEpochNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastAppliedForkNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"forkNumber\",\"type\":\"uint256\"}],\"name\":\"lastBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"lastBlock\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"forkNumber\",\"type\":\"uint256\"}],\"name\":\"lastEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"lastBlock\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"lastNonEmptyRequestEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_trieKey\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_trieValue\",\"type\":\"bytes\"}],\"name\":\"makeERU\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rootchain\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_childchain\",\"type\":\"address\"}],\"name\":\"mapRequestableContractByOperator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"numEnterForORB\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"operator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"prepareToSubmitURB\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceMapper\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceSubmitter\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"requestableContracts\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"seigManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"setSeigManager\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_trieKey\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_trieValue\",\"type\":\"bytes\"}],\"name\":\"startEnter\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_trieKey\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_trieValue\",\"type\":\"bytes\"}],\"name\":\"startExit\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"submitHandler\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pos1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_pos2\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_epochStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_epochTransactionsRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_epochReceiptsRoot\",\"type\":\"bytes32\"}],\"name\":\"submitNRE\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pos\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_statesRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_transactionsRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_receiptsRoot\",\"type\":\"bytes32\"}],\"name\":\"submitORB\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pos\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_statesRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_transactionsRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_receiptsRoot\",\"type\":\"bytes32\"}],\"name\":\"submitURB\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"}]"

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
	"80e3e81d": "addMapper(address)",
	"072900f9": "addSubmitter(address)",
	"404f7d66": "challengeExit(uint256,uint256,uint256,bytes,bytes)",
	"6299fb24": "challengeNullAddress(uint256,bytes,bytes,uint256,bytes32[])",
	"06394c9b": "changeOperator(address)",
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
	"c0c49c2a": "isMapper(address)",
	"420bb4b8": "isRootChain()",
	"a926fdbc": "isSubmitter(address)",
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
	"3565fb0d": "renounceMapper()",
	"5e0ca71b": "renounceSubmitter()",
	"da0185f8": "requestableContracts(address)",
	"6fb7f558": "seigManager()",
	"7657f20a": "setSeigManager(address)",
	"2aa196c8": "startEnter(address,bytes32,bytes)",
	"e7f96505": "startExit(address,bytes32,bytes)",
	"e259faf7": "submitHandler()",
	"0eaf45a8": "submitNRE(uint256,uint256,bytes32,bytes32,bytes32)",
	"a820c067": "submitORB(uint256,bytes32,bytes32,bytes32)",
	"6f3e4b90": "submitURB(uint256,bytes32,bytes32,bytes32)",
}

// RootChainBin is the compiled bytecode used for deploying new contracts.
var RootChainBin = "0x60806040523480156200001157600080fd5b5060405162005f9438038062005f9483398101604081905262000034916200063e565b62000048336001600160e01b036200026316565b6200005c336001600160e01b03620002b516565b6200007b886001600160a01b03166200030760201b620035731760201c565b6200008557600080fd5b620000a4876001600160a01b03166200030760201b620035731760201c565b620000ae57600080fd5b620000cd866001600160a01b03166200030760201b620035731760201c565b620000d757600080fd5b600180546001600160a01b03199081166001600160a01b038b8116919091179092556002805482168a8416178155600380549092169289169290921781556000805460ff191688151517610100600160a81b031916610100330217815560058790556006548152600860209081526040808320838052600480820184528285208088018b90558087018a90559081018890559481019092528083208054790100000000000000000000000000000000000000000000000000600160801b600160c01b0319909116700100000000000000000000000000000000426001600160401b0316021760ff60c81b191617905593825292812080547a010000000000000000000000000000000000000000000000000000780100000000000000000000000000000000000000000000000060ff60c01b199092169190911760ff60d01b19161790556200022a90839083906200030d565b620002408260006001600160e01b03620003f316565b620002536001600160e01b03620004c916565b50505050505050505050620008b8565b6200027e8160166200059860201b620038421790919060201c565b6040516001600160a01b038216907fa199526a01dbcba68aa3ab9d7d06c6692c83fb1a8bcf2184e6d94ad34f5aaf9590600090a250565b620002d08160176200059860201b620038421790919060201c565b6040516001600160a01b038216907fb079bc2cbde1f186e0b351d4a87c4597e3ed098f571548617449e73506428d8b90600090a250565b3b151590565b60058201805464010000000060ff60201b1990911617905581546001600160c01b03167801000000000000000000000000000000000000000000000000426001600160401b0390811691909102919091178355600184018054600160801b600160c01b03191670010000000000000000000000000000000084841602178082558454600160401b600160801b0319909116921668010000000000000000029190911790556006546040517ffb96205e4b3633fd57aa805b26b51ecf528714a10241a4af015929dce86768d991620003e6918490620007ad565b60405180910390a1505050565b6000818152600383016020526040908190206006808201805470010000000000000000000000000000000060ff60801b199091168117600160401b600160801b031990811668010000000000000000426001600160401b03908116820292909217909455855460018a018054600160801b600160c01b0319169186900483169094021790911681881684021790915591548354945193947f70801d4d63b3da6c19ba7349911f45bed5a99ccdfb51b8138c105872529bebd594620003e6949293889383811693919091041690620007cc565b5050565b6001546040516000916060916001600160a01b0390911690620004ec90620007a0565b60408051918290038220600483526024830182526020830180516001600160e01b03167fffffffff00000000000000000000000000000000000000000000000000000000909216919091179052516200054691906200078b565b600060405180830381855af49150503d806000811462000583576040519150601f19603f3d011682016040523d82523d6000602084013e62000588565b606091505b509150915081620004c557600080fd5b620005ad82826001600160e01b03620005dd16565b15620005b857600080fd5b6001600160a01b0316600090815260209190915260409020805460ff19166001179055565b60006001600160a01b038216620005f357600080fd5b506001600160a01b03811660009081526020839052604090205460ff165b92915050565b8051620006118162000888565b80516200061181620008a2565b80516200061181620008ad565b600080600080600080600080610100898b0312156200065c57600080fd5b60006200066a8b8b62000617565b98505060206200067d8b828c0162000617565b9750506040620006908b828c0162000617565b9650506060620006a38b828c0162000624565b9550506080620006b68b828c0162000631565b94505060a0620006c98b828c0162000631565b93505060c0620006dc8b828c0162000631565b92505060e0620006ef8b828c0162000631565b9150509295985092959890939650565b60006200070c8262000812565b62000718818562000816565b93506200072a81856020860162000855565b9290920192915050565b600062000743600c8362000816565b7f707265706172654e5245282900000000000000000000000000000000000000008152600c0192915050565b6200077a816200082d565b82525050565b6200077a8162000848565b6000620007998284620006ff565b9392505050565b6000620006118262000734565b60408101620007bd82856200076f565b6200079960208301846200076f565b60808101620007dc82876200076f565b620007eb60208301866200076f565b620007fa604083018562000780565b62000809606083018462000780565b95945050505050565b5190565b919050565b6000620006118262000830565b151590565b90565b6001600160a01b031690565b6001600160401b031690565b600062000611826200083c565b60005b838110156200087257818101518382015260200162000858565b8381111562000882576000848401525b50505050565b62000893816200081b565b81146200089f57600080fd5b50565b620008938162000828565b62000893816200082d565b6156cc80620008c86000396000f3fe6080604052600436106104105760003560e01c80638b5172d01161021e578063c8ad329f11610123578063e259faf7116100ab578063ea0c73f61161007a578063ea0c73f614610ab5578063ea7f22a814610aca578063f28a7afa14610aea578063f4f31de414610b0a578063fb788a2714610b2a57610410565b8063e259faf714610a70578063e6925d0814610a85578063e7b88b8014610a8d578063e7f9650514610aa257610410565b8063d1723a96116100f2578063d1723a96146109ee578063d636857e14610a1b578063d691acd81461044b578063da0185f814610a3b578063de0ce17d14610a5b57610410565b8063c8ad329f14610979578063ca6f63801461098e578063cb5d742f146109ae578063ce8a2bc2146109ce57610410565b8063b2ae9ba8116101a6578063b8066bcb11610175578063b8066bcb146108e8578063c0c49c2a146108fd578063c0e860641461091d578063c2bc88fa1461094f578063c54626cc1461096457610410565b8063b2ae9ba81461044b578063b443f3cc1461087c578063b540adba146108b3578063b6715647146108c857610410565b806399bd3600116101ed57806399bd36001461080a578063a820c0671461081f578063a926fdbc14610832578063ab96da2d14610852578063b17fa6e91461086757610410565b80638b5172d01461044b5780638eb288ca146107e057806393521222146107f557806394be3aa51461044b57610410565b8063404f7d66116103245780636f3e4b90116102ac5780637657f20a1161027b5780637657f20a1461076357806378fe705f146107835780637b929c271461079657806380e3e81d146107ab5780638155717d146107cb57610410565b80636f3e4b90146107065780636fb7f5581461071957806372ecb9a81461072e57806375395a581461074e57610410565b806354768571116102f3578063547685711461066f578063570ca7351461068f5780635b884682146106b15780635e0ca71b146106d15780636299fb24146106e657610410565b8063404f7d66146105d5578063420bb4b8146105f55780634a44bdb81461060a5780634ba3a1261461063757610410565b8063183d2d1c116103a75780632aa196c8116103765780632aa196c8146105565780632b25a38b146105695780632dc6bb7b146105965780633565fb0d146105ab578063398bac63146105c057610410565b8063183d2d1c1461050c578063192adc5b1461044b5780631ec2042b14610521578063236915661461054157610410565b806308c4fff0116103e357806308c4fff0146104a25780630eaf45a8146104b757806311e4c914146104d7578063164bc2ae146104f757610410565b8063019dc09914610415578063033cfbed1461044b57806306394c9b14610460578063072900f914610482575b600080fd5b34801561042157600080fd5b506104356104303660046148cd565b610b3f565b604051610442919061526a565b60405180910390f35b34801561045757600080fd5b50610435610b67565b34801561046c57600080fd5b5061048061047b3660046147eb565b610b6c565b005b34801561048e57600080fd5b5061048061049d3660046147eb565b610be3565b3480156104ae57600080fd5b50610435610c0b565b6104ca6104c5366004614a73565b610c10565b604051610442919061515f565b3480156104e357600080fd5b506104356104f23660046148cd565b610c51565b34801561050357600080fd5b50610435610c73565b34801561051857600080fd5b50610435610c79565b34801561052d57600080fd5b5061043561053c3660046148cd565b610c7f565b34801561054d57600080fd5b50610435610c9d565b6104ca61056436600461484b565b610ca3565b34801561057557600080fd5b50610589610584366004614a43565b610d65565b604051610442919061524c565b3480156105a257600080fd5b506104356110cd565b3480156105b757600080fd5b506104806110d3565b3480156105cc57600080fd5b506105896110de565b3480156105e157600080fd5b506104806105f0366004614ae8565b611479565b34801561060157600080fd5b506104ca611705565b34801561061657600080fd5b5061062a610625366004614a43565b61170a565b604051610442919061525b565b34801561064357600080fd5b506106576106523660046148cd565b6117fa565b6040516104429c9b9a999897969594939291906154b8565b34801561067b57600080fd5b506104ca61068a3660046148cd565b611871565b34801561069b57600080fd5b506106a461189f565b6040516104429190615136565b3480156106bd57600080fd5b506104356106cc366004614a43565b6118b3565b3480156106dd57600080fd5b506104806118e4565b3480156106f257600080fd5b5061048061070136600461497c565b6118ed565b6104ca61071436600461491b565b611950565b34801561072557600080fd5b506106a4611958565b34801561073a57600080fd5b506104356107493660046148cd565b611967565b34801561075a57600080fd5b506104ca611979565b34801561076f57600080fd5b5061048061077e3660046147eb565b611992565b6104ca61079136600461484b565b6119d0565b3480156107a257600080fd5b506104ca611a3e565b3480156107b757600080fd5b506104806107c63660046147eb565b611a47565b3480156107d757600080fd5b50610435611a6c565b3480156107ec57600080fd5b50610435611a71565b34801561080157600080fd5b50610435611a78565b34801561081657600080fd5b506104ca611a87565b6104ca61082d36600461491b565b612327565b34801561083e57600080fd5b506104ca61084d3660046147eb565b612366565b34801561085e57600080fd5b50610435612379565b34801561087357600080fd5b5061043561237f565b34801561088857600080fd5b5061089c6108973660046148cd565b612384565b6040516104429b9a9998979695949392919061540e565b3480156108bf57600080fd5b506104356124a5565b3480156108d457600080fd5b506104356108e33660046148cd565b6124ab565b3480156108f457600080fd5b506106a46124bd565b34801561090957600080fd5b506104ca6109183660046147eb565b6124cc565b34801561092957600080fd5b5061093d6109383660046148cd565b6124df565b604051610442969594939291906151b4565b34801561095b57600080fd5b50610435612546565b34801561097057600080fd5b5061043561254b565b34801561098557600080fd5b50610435612551565b34801561099a57600080fd5b506104356109a93660046148cd565b612557565b3480156109ba57600080fd5b506104ca6109c9366004614811565b612569565b3480156109da57600080fd5b506104ca6109e93660046148cd565b61262f565b3480156109fa57600080fd5b50610a0e610a093660046148cd565b612637565b604051610442919061522b565b348015610a2757600080fd5b50610435610a363660046148cd565b6127e0565b348015610a4757600080fd5b506106a4610a563660046147eb565b612805565b348015610a6757600080fd5b506106a4610b67565b348015610a7c57600080fd5b506106a4612820565b61048061282f565b348015610a9957600080fd5b506106a4612836565b6104ca610ab036600461484b565b612845565b348015610ac157600080fd5b506104356128a1565b348015610ad657600080fd5b5061093d610ae53660046148cd565b6128a7565b348015610af657600080fd5b506104ca610b053660046148eb565b6128b4565b348015610b1657600080fd5b5061089c610b253660046148cd565b6128fd565b348015610b3657600080fd5b5061043561290a565b600081815260086020526040902060010154600160401b90046001600160401b03165b919050565b600081565b60005461010090046001600160a01b03163314610b8857600080fd5b60008054610100600160a81b0319166101006001600160a01b038416021790556040517f4721129e0e676ed6a92909bb24e853ccdd63ad72280cc2e974e38e480e0e6e5490610bd8908390615136565b60405180910390a150565b60005461010090046001600160a01b03163314610bff57600080fd5b610c0881612910565b50565b600f81565b6000610c1b33612366565b610c2457600080fd5b6000805460ff16610c3957610c37612958565b505b610c468787878787612ace565b979650505050505050565b600090815260086020526040902054600160801b90046001600160401b031690565b60105481565b60065481565b6000908152600860205260409020600101546001600160401b031690565b600d5481565b60008034610cb96009600b888489898880612bad565b600d805460010190556006546000908152600860205260408082209051929450917f6940a01870e576ceb735867e13863646d517ce10e66c0133186a4ebdfe9388c291610d089186919061532c565b60405180910390a17f879922cf5fcada9ebaf8bd7424dc62877f4b220cae07fb6695cc1e8f94c52b4d833389858a8a600080604051610d4e9897969594939291906152f6565b60405180910390a1600193505050505b9392505050565b610d6d6142cd565b600860008481526020019081526020016000206003016000838152602001908152602001600020604051806101400160405290816000820160009054906101000a90046001600160401b03166001600160401b03166001600160401b031681526020016000820160089054906101000a90046001600160401b03166001600160401b03166001600160401b031681526020016000820160109054906101000a90046001600160401b03166001600160401b03166001600160401b031681526020016000820160189054906101000a900460ff161515151581526020016000820160199054906101000a900460ff1615151515815260200160008201601a9054906101000a900460ff1615151515815260200160008201601b9054906101000a900460ff1615151515815260200160008201601c9054906101000a900460ff16151515158152602001600182016040518060c00160405290816000820160009054906101000a90046001600160401b03166001600160401b03166001600160401b031681526020016000820160089054906101000a90046001600160401b03166001600160401b03166001600160401b031681526020016000820160109054906101000a90046001600160401b03166001600160401b03166001600160401b031681526020016000820160189054906101000a90046001600160401b03166001600160401b03166001600160401b031681526020016001820160009054906101000a90046001600160401b03166001600160401b03166001600160401b031681526020016001820160089054906101000a90046001600160401b03166001600160401b03166001600160401b031681525050815260200160038201604051806101000160405290816000820154815260200160018201548152602001600282015481526020016003820160009054906101000a90046001600160401b03166001600160401b03166001600160401b031681526020016003820160089054906101000a90046001600160401b03166001600160401b03166001600160401b031681526020016003820160109054906101000a900460ff161515151581526020016003820160119054906101000a900460ff161515151581526020016003820160129054906101000a900460ff1615151515815250508152505090505b92915050565b60135481565b6110dc336131cf565b565b6110e66142cd565b600860006006548152602001908152602001600020600301600060086000600654815260200190815260200160002060000160109054906101000a90046001600160401b03166001600160401b03168152602001908152602001600020604051806101400160405290816000820160009054906101000a90046001600160401b03166001600160401b03166001600160401b031681526020016000820160089054906101000a90046001600160401b03166001600160401b03166001600160401b031681526020016000820160109054906101000a90046001600160401b03166001600160401b03166001600160401b031681526020016000820160189054906101000a900460ff161515151581526020016000820160199054906101000a900460ff1615151515815260200160008201601a9054906101000a900460ff1615151515815260200160008201601b9054906101000a900460ff1615151515815260200160008201601c9054906101000a900460ff16151515158152602001600182016040518060c00160405290816000820160009054906101000a90046001600160401b03166001600160401b03166001600160401b031681526020016000820160089054906101000a90046001600160401b03166001600160401b03166001600160401b031681526020016000820160109054906101000a90046001600160401b03166001600160401b03166001600160401b031681526020016000820160189054906101000a90046001600160401b03166001600160401b03166001600160401b031681526020016001820160009054906101000a90046001600160401b03166001600160401b03166001600160401b031681526020016001820160089054906101000a90046001600160401b03166001600160401b03166001600160401b031681525050815260200160038201604051806101000160405290816000820154815260200160018201548152602001600282015481526020016003820160009054906101000a90046001600160401b03166001600160401b03166001600160401b031681526020016003820160089054906101000a90046001600160401b03166001600160401b03166001600160401b031681526020016003820160109054906101000a900460ff161515151581526020016003820160119054906101000a900460ff161515151581526020016003820160129054906101000a900460ff1615151515815250508152505090505b90565b600087815260086020908152604080832089845260048101909252909120600581015460ff166114a857600080fd5b6005810154640100000000900460ff166114c157600080fd5b6005810154600090610100900460ff1680156115cd5761159a83600c8560000160089054906101000a90046001600160401b03166001600160401b03168154811061150857fe5b9060005260206000209060020201600a8c8c8c8080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050508b8b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061321792505050565b60405190925033906108fc9060009081818181818888f193505050501580156115c7573d6000803e3d6000fd5b506116bf565b61169083600b8560000160089054906101000a90046001600160401b03166001600160401b0316815481106115fe57fe5b906000526020600020906002020160098c8c8c8080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050508b8b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061321792505050565b60405190925033906108fc9060009081818181818888f193505050501580156116bd573d6000803e3d6000fd5b505b7fc8135db115644ed4ae193313c4c801235ef740d2a57a8d5e6fe26ab66635698a82826040516116f092919061532c565b60405180910390a15050505050505050505050565b600181565b61171261432d565b506000918252600860209081526040808420928452600492830182529283902083516101a08101855281546001600160401b038082168352600160401b8204811694830194909452600160801b8104841695820195909552600160c01b9094048216606085015260018101549091166080840152600281015460a0840152600381015460c08401529081015460e08301526005015460ff80821615156101008085019190915282048116151561012084015262010000820481161515610140840152630100000082048116151561016084015264010000000090910416151561018082015290565b6008602052600090815260409020805460018201546002909201546001600160401b0380831693600160401b808504831694600160801b808204851695600160c01b92839004861695858116958581048216958482048316959182900483169484841694918204841693908204169160ff9104168c565b6000805b8281101561189657611885611a87565b61188e57600080fd5b600101611875565b50600192915050565b60005461010090046001600160a01b031681565b600091825260086020908152604080842092845260049092019052902054600160c01b90046001600160401b031690565b6110dc33613322565b60065460009081526008602090815260408083208b845260048101909252909120600581015460ff161561192057600080fd5b8054426001600160401b03600160801b90920491909116600f011161194457600080fd5b50505050505050505050565b600080610410565b6004546001600160a01b031681565b60076020526000908152604090205481565b6000611983612958565b61198c57600080fd5b50600190565b60005461010090046001600160a01b031633146119ae57600080fd5b600480546001600160a01b0319166001600160a01b0392909216919091179055565b60008060006119e9600a600c8860008989600180612bad565b90507f879922cf5fcada9ebaf8bd7424dc62877f4b220cae07fb6695cc1e8f94c52b4d81338860008989600180604051611a2a989796959493929190615278565b60405180910390a150600195945050505050565b60005460ff1681565b60005461010090046001600160a01b03163314611a6357600080fd5b610c088161336a565b600a81565b620186a081565b6000611a826133b2565b905090565b60105460009081526008602052604081206011546001820154601254849392916001600160401b03161015611abb57600080fd5b601254600090815260048301602090815260408083208484526003860190925290912083546001600160401b031615801590611b05575083546012546001600160401b0390911611155b15611b705750506010805460010190819055600090815260086020908152604080832080546001600160401b03600160401b82048116808752600384018652848720600160c01b90930490911660128190556011829055865260048301909452919093209294509092505b600582015460ff16611c025782611b9a576010546000908152600f60205260409020549250611bc1565b600092835260038401602052604090922060020154600160401b90046001600160401b0316915b82611bcb57600080fd5b50506000818152600383016020908152604080832080546001600160401b0316601281905584526004860190925290912090611c22565b5080546001600160401b0316600081815260038501602052604090209092505b60118390558054600160c01b900460ff1615611c3d57600080fd5b8054600160d01b900460ff16611c5257600080fd5b600582015460ff16611c6357600080fd5b6005820154640100000000900460ff16611c7c57600080fd5b8154426001600160401b03600160c01b90920491909116600a011115611ca157600080fd5b6005820154610100900460ff161561201557601454600a549095508510611cc757600080fd5b6000600a8681548110611cd657fe5b906000526020600020906006020190506000600c8460000160089054906101000a90046001600160401b03166001600160401b031681548110611d1557fe5b600091825260209091206002909102018054909150600160881b90046001600160401b03168710801590611d56575060018101546001600160401b03168711155b611d5f57600080fd5b60018101546001600160401b0316871415611dba5785546001600160401b031615801590611da1575085546012546000196001600160401b0392831601909116145b15611db0576010805460010190555b6012805460010190555b600187016014558154600160401b900460ff168015611de257508154600160581b900460ff16155b15611fb957604080516101608101825283546001600160401b0381168252600160401b810460ff9081161515602080850191909152600160481b83048216151584860152600160501b8304821615156060850152600160581b830490911615156080840152600160601b9091046001600160801b031660a08301526001808601546001600160a01b0390811660c085015260028088015490911660e085015260038701546101008086019190915260048801546101208601526005880180548751948116159092026000190190911691909104601f8101849004840283018401909552848252611f43948c9493889361014086019390929091830182828015611f2c5780601f10611f0157610100808354040283529160200191611f2c565b820191906000526020600020905b815481529060010190602001808311611f0f57829003601f168201915b5050505050815250506133b790919063ffffffff16565b5060018201546040516001600160a01b03909116906108fc9060009081818181818888f19350505050158015611f7d573d6000803e3d6000fd5b507f6940a01870e576ceb735867e13863646d517ce10e66c0133186a4ebdfe9388c2876001604051611fb092919061532c565b60405180910390a15b815460ff60501b1916600160501b1782556040517f134017cf3262b18f892ee95dde3b0aec9a80cc70a7c96f09c64bd237aceb047390611ffd90899060019061532c565b60405180910390a16001975050505050505050611476565b601354600954909550851061202957600080fd5b8154600b8054600092600160401b90046001600160401b031690811061204b57fe5b600091825260209091206002909102018054909150600160881b90046001600160401b0316861080159061208c575060018101546001600160401b03168611155b61209557600080fd5b60018101546001600160401b03168614156120f05784546001600160401b0316158015906120d7575084546012546000196001600160401b0392831601909116145b156120e6576010805460010190555b6012805460010190555b6000600987815481106120ff57fe5b6000918252602090912060018901601355600690910201805460ff60501b1916600160501b17808255909150600160401b900460ff16801561214a57508054600160581b900460ff16155b156122df57604080516101608101825282546001600160401b0381168252600160401b810460ff9081161515602080850191909152600160481b83048216151584860152600160501b8304821615156060850152600160581b830490911615156080840152600160601b9091046001600160801b031660a08301526001808501546001600160a01b0390811660c085015260028087015490911660e085015260038601546101008086019190915260048701546101208601526005870180548751948116159092026000190190911691909104601f8101849004840283018401909552848252612269948c9493879361014086019390929091830182828015611f2c5780601f10611f0157610100808354040283529160200191611f2c565b5060018101546040516001600160a01b03909116906108fc9060009081818181818888f193505050501580156122a3573d6000803e3d6000fd5b507f6940a01870e576ceb735867e13863646d517ce10e66c0133186a4ebdfe9388c28760006040516122d692919061532c565b60405180910390a15b7f134017cf3262b18f892ee95dde3b0aec9a80cc70a7c96f09c64bd237aceb047387600060405161231192919061532c565b60405180910390a1600197505050505050505090565b600061233233612366565b61233b57600080fd5b6000805460ff166123505761234e612958565b505b61235c86868686613464565b9695505050505050565b60006110c760178363ffffffff61353e16565b60055481565b601481565b6009818154811061239157fe5b6000918252602091829020600691909102018054600180830154600280850154600386015460048701546005880180546040805161010099831615999099026000190190911695909504601f81018b90048b0288018b019095528487526001600160401b0388169a50600160401b880460ff9081169a600160481b8a0482169a600160501b8b0483169a600160581b810490931699600160601b9093046001600160801b0316986001600160a01b039081169897169690939183018282801561249b5780601f106124705761010080835404028352916020019161249b565b820191906000526020600020905b81548152906001019060200180831161247e57829003601f168201915b505050505090508b565b60095490565b600e6020526000908152604090205481565b6003546001600160a01b031681565b60006110c760168363ffffffff61353e16565b600c81815481106124ec57fe5b60009182526020909120600290910201805460019091015460ff821692506001600160401b036101008304811692600160481b8104821692600160881b9091048216918116906001600160a01b03600160401b9091041686565b603c81565b60145481565b60115481565b600f6020526000908152604090205481565b6000612574336124cc565b61257d57600080fd5b61258f836001600160a01b0316613573565b61259857600080fd5b6001600160a01b0383811660009081526015602052604090205416156125bd57600080fd5b6001600160a01b038381166000908152601560205260409081902080546001600160a01b03191692851692909217909155517fc5ec2ed49686197edd2ed642c7e6096893cc81e6658cde2527030316037715d09061261e9085908590615144565b60405180910390a150600192915050565b600654141590565b606060006009838154811061264857fe5b600091825260208083206002600690930201828101546001600160a01b039081168086526015845260408087205481516101608101835285546001600160401b038116825260ff600160401b820481161515838a0152600160481b82048116151583860152600160501b8204811615156060840152600160581b820416151560808301526001600160801b03600160601b9091041660a0820152600180870154861660c083015260e082019490945260038601546101008281019190915260048701546101208301526005870180548551601f60001998831615909402979097011699909904908101889004880285018801909352828452949850610d5e976127db978c9791966127ce96939093169492938b93610140860193909291908301828280156127b75780601f1061278c576101008083540402835291602001916127b7565b820191906000526020600020905b81548152906001019060200180831161279a57829003601f168201915b50505050508152505061357990919063ffffffff16565b919063ffffffff61364716565b61369f565b600090815260086020526040902060010154600160801b90046001600160401b031690565b6015602052600090815260409020546001600160a01b031681565b6002546001600160a01b031681565b6000610410565b6001546001600160a01b031681565b600080600061285f6009600b886000898960016000612bad565b90507f879922cf5fcada9ebaf8bd7424dc62877f4b220cae07fb6695cc1e8f94c52b4d8133886000898960016000604051611a2a989796959493929190615278565b600b5490565b600b81815481106124ec57fe5b600081156128ce57600a83815481106128c957fe5b506000525b600983815481106128db57fe5b6000918252602090912060069091020154600160501b900460ff169392505050565b600a818154811061239157fe5b60125481565b61292160178263ffffffff61384216565b6040516001600160a01b038216907fb079bc2cbde1f186e0b351d4a87c4597e3ed098f571548617449e73506428d8b90600090a250565b600654600090815260086020526040812080546001600160401b031615612983576000915050611476565b80546001808301546000926129b5926001600160401b03600160c01b909204821692600160801b90048216011661387b565b60018301549091506001600160401b03168111156129d857600092505050611476565b6000818152600483016020526040902060058101546301000000900460ff1615612a085760009350505050611476565b600581015460ff1615612a58578054426001600160401b03600160801b90920491909116600f011115612a415760009350505050611476565b612a4c838284613892565b60019350505050611476565b6001808401546001600160401b03600160401b9091048116909101165b6000818152600385016020526040902054600160d01b900460ff1615612a9d57600101612a75565b612aa78482613952565b15612ac357612ab68482613a0c565b6001945050505050611476565b600094505050505090565b60025460405160009182916060916001600160a01b031690612aef9061510a565b60405180910390208989898989604051602401612b10959493929190615397565b60408051601f198184030181529181526020820180516001600160e01b03166001600160e01b0319909416939093179092529051612b4e91906150fe565b600060405180830381855af49150503d8060008114612b89576040519150601f19603f3d011682016040523d82523d6000602084013e612b8e565b606091505b509150915081612b9d57600080fd5b6001925050505b95945050505050565b600061040084511115612bbf57600080fd5b6003546001600160a01b03888116911614808015612bdb575083155b80612bff57506001600160a01b038881166000908152601560205260409020541615155b612c0857600080fd5b8954612c178b60018301614399565b915060008a8381548110612c2757fe5b600091825260209182902060069190910201600181018054336001600160a01b0319918216179091556002820180549091166001600160a01b038d16179055805467ffffffffffffffff1916426001600160401b03161768ff00000000000000001916600160401b881515021769ff0000000000000000001916600160481b85151502176fffffffffffffffffffffffffffffffff60601b1916600160601b6001600160801b038c1602178155600381018990558751909250612cf2916005840191908901906143ca565b5084612e1b57604080516101608101825282546001600160401b0381168252600160401b810460ff9081161515602080850191909152600160481b83048216151584860152600160501b8304821615156060850152600160581b830490911615156080840152600160601b9091046001600160801b031660a08301526001808501546001600160a01b0390811660c085015260028087015490911660e085015260038601546101008086019190915260048701546101208601526005870180548751948116159092026000190190911691909104601f8101849004840283018401909552848252612e1294889493879361014086019390929091830182828015611f2c5780601f10611f0157610100808354040283529160200191611f2c565b612e1b57600080fd5b895460009015612e30578a5460001901612e3f565b8a54612e3f8c60018301614448565b905060008b8281548110612e4f57fe5b60009182526020909120600290910201805490915060ff1680612e975750612e756133b2565b81546001808401546001600160401b03600160881b9093048316908316030116145b15612efa57805460ff1916600190811782558c548d91612eba9083908301614448565b81548110612ec457fe5b60009182526020909120600290910201805467ffffffffffffffff60881b1916600160881b6001600160401b0388160217815590505b612f0381610c08565b60018101805467ffffffffffffffff19166001600160401b03871617905586612f515780546001600160401b0361010080830482166001019091160268ffffffffffffffff00199091161781555b838015612f5c575086155b1561309357604080516101608101825284546001600160401b0381168252600160401b810460ff9081161515602080850191909152600160481b83048216151584860152600160501b8304821615156060850152600160581b830490911615156080840152600160601b9091046001600160801b031660a08301526001808701546001600160a01b0390811660c085015260028089015490911660e085015260038801546101008086019190915260048901546101208601526005890180548751948116159092026000190190911691909104601f810184900484028301840190955284825261308e94889461307f94339491938793610140860193928301828280156127b75780601f1061278c576101008083540402835291602001916127b7565b8391908863ffffffff613acc16565b6131bf565b6001600160a01b038b81166000908152601560209081526040918290205482516101608101845287546001600160401b0381168252600160401b810460ff908116151583860152600160481b82048116151583870152600160501b8204811615156060840152600160581b82041615156080830152600160601b90046001600160801b031660a0820152600180890154861660c08301526002808a0154871660e084015260038a01546101008085019190915260048b015461012085015260058b0180548851948116159092026000190190911691909104601f81018690048602830186019096528582526131bf968a9661307f9695909116948793610140860193909290918301828280156127b75780601f1061278c576101008083540402835291602001916127b7565b5050505098975050505050505050565b6131e060168263ffffffff613aee16565b6040516001600160a01b038216907ff012269e1b6bb33c7800e7e54167ddd74f0ae7244108030b6876821d3779822f90600090a250565b845460018601546001600160401b03600160881b90920482168501911681111561324057600080fd5b600085828154811061324e57fe5b6000918252602090912089546006909202019150426001600160401b03600160c01b90920491909116600a011161328457600080fd5b6005880154640100000000900460ff1661329d57600080fd5b8054600160581b900460ff16156132b357600080fd5b8054600160501b900460ff16156132c957600080fd5b835160208501206132d985613b23565b156132e357600080fd5b60005460ff16613306576132fd81878b6004015487613b60565b61330657600080fd5b50805460ff60581b1916600160581b1790559695505050505050565b61333360178263ffffffff613aee16565b6040516001600160a01b038216907ff84a004e1673d2f349a7c93c72b3794b8eba6d2f9338044d8c8cd260e51a57a190600090a250565b61337b60168263ffffffff61384216565b6040516001600160a01b038216907fa199526a01dbcba68aa3ab9d7d06c6692c83fb1a8bcf2184e6d94ad34f5aaf9590600090a250565b601490565b600062019a285a116133c857600080fd5b8260e001516001600160a01b031663a9f793088460200151848660c001518761010001518861014001516040518663ffffffff1660e01b815260040161341295949392919061516d565b602060405180830381600087803b15801561342c57600080fd5b505af1158015613440573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250610d5e91908101906148af565b60025460405160009182916060916001600160a01b03169061348590615120565b6040518091039020888888886040516024016134a49493929190615347565b60408051601f198184030181529181526020820180516001600160e01b03166001600160e01b03199094169390931790925290516134e291906150fe565b600060405180830381855af49150503d806000811461351d576040519150601f19603f3d011682016040523d82523d6000602084013e613522565b606091505b50915091508161353157600080fd5b5060019695505050505050565b60006001600160a01b03821661355357600080fd5b506001600160a01b03166000908152602091909152604090205460ff1690565b3b151590565b613581614474565b6020808401805115159183019190915260408085015115159083015260c0808501516001600160a01b031690830152511580156135bf575082604001515b156136095760c08301516001600160a01b031660e082015261014083015160208101516001600160801b0381166135f557600080fd5b6001600160801b031660a0830152506110c7565b6001600160a01b03821660e082015260a0808401516001600160801b0316908201526101008084015190820152610140808401519082015292915050565b61364f6144d2565b633b9aca006020820152620186a0604082015260e08401516001600160a01b0316606082015260a08401516001600160801b03166080820152613693848484613c2b565b60a08201529392505050565b6040805160098082526101408201909252606091829190816020015b60608152602001906001900390816136bb57505083519091506136e6906001600160401b0316613d01565b816000815181106136f357fe5b602002602001018190525061370b8360200151613d01565b8160018151811061371857fe5b602002602001018190525061373983604001516001600160401b0316613d01565b8160028151811061374657fe5b602002602001018190525061376783606001516001600160a01b0316613d14565b8160038151811061377457fe5b602002602001018190525061378c8360800151613d01565b8160048151811061379957fe5b60200260200101819052506137b18360a00151613d33565b816005815181106137be57fe5b60200260200101819052506137d68360c00151613d01565b816006815181106137e357fe5b60200260200101819052506137fb8360e00151613d01565b8160078151811061380857fe5b6020026020010181905250613821836101000151613d01565b8160088151811061382e57fe5b6020026020010181905250610d5e81613d89565b61384c828261353e565b1561385657600080fd5b6001600160a01b0316600090815260209190915260409020805460ff19166001179055565b60008183101561388b5781610d5e565b5090919050565b60058201805464010000000064ff000000001990911617905581546001600160c01b0316600160c01b426001600160401b039081169190910291909117835560018401805467ffffffffffffffff60801b1916600160801b8484160217808255845467ffffffffffffffff60401b199091169216600160401b029190911790556006546040517ffb96205e4b3633fd57aa805b26b51ecf528714a10241a4af015929dce86768d99161394591849061537c565b60405180910390a1505050565b8154600090600160801b90046001600160401b0316821115613976575060006110c7565b600082815260038401602052604090208054600160c81b900460ff166139a05760009150506110c7565b8054600160d01b900460ff16156139bb5760009150506110c7565b6006810154600160881b900460ff16806139e057506006810154600160901b900460ff165b156139ef5760009150506110c7565b60060154426001600160401b039091166014011115905092915050565b60008181526003830160205260409081902060068082018054600160801b60ff60801b19909116811767ffffffffffffffff60401b19908116600160401b426001600160401b03908116820292909217909455855460018a01805467ffffffffffffffff60801b19169186900483169094021790911681881684021790915591548354945193947f70801d4d63b3da6c19ba7349911f45bed5a99ccdfb51b8138c105872529bebd5946139459492938893838116939190910416906153d9565b613ae0613adb83836000613647565b613de6565b836004018190555050505050565b613af8828261353e565b613b0157600080fd5b6001600160a01b0316600090815260209190915260409020805460ff19169055565b60006060613b416004613b3585613e02565b9063ffffffff613e2516565b9050610d5e81600081518110613b5357fe5b6020026020010151613eb2565b60006020825181613b6d57fe5b0615613b7857600080fd5b60006020835181613b8557fe5b04905060108110613b9557600080fd5b60008660205b836020028111613c1d5785810151925060028806613be3578183604051602001613bc69291906150bc565b604051602081830303815290604052805190602001209150613c0f565b8282604051602001613bf69291906150bc565b6040516020818303038152906040528051906020012091505b600288049750602001613b9b565b509094149695505050505050565b606083604001518015613c4057508360200151155b15613c4a57610d5e565b600082613c6a57604051613c5d90615115565b6040518091039020613c7f565b604051613c769061512b565b60405180910390205b9050808560200151613c92576000613c95565b60015b60ff1660001b858760c001516001600160a01b0316886101000151896101400151604051602001613cca959493929190615203565b60408051601f1981840301815290829052613ce892916020016150e2565b6040516020818303038152906040529150509392505050565b60606110c7613d0f83613ed6565b613d33565b60408051600560a21b8318601482015260348101909152606090610d5e815b606081516001148015613d655750607f60f81b82600081518110613d5357fe5b01602001516001600160f81b03191611155b15613d71575080610b62565b6110c7613d838351608060ff16613fcf565b83614098565b604080516000808252602082019092526060915b8351811015613dcd57613dc382858381518110613db657fe5b6020026020010151614098565b9150600101613d9d565b50610d5e613de0825160c060ff16613fcf565b82614098565b60006060613df38361369f565b80516020909101209392505050565b613e0a614539565b50805160408051808201909152602092830181529182015290565b606081604051908082528060200260200182016040528015613e6157816020015b613e4e614539565b815260200190600190039081613e465790505b509050613e6c614553565b613e758461417d565b905060005b83811015613eaa57613e8b826141a2565b838281518110613e9757fe5b6020908102919091010152600101613e7a565b505092915050565b6000806000613ec0846141d4565b90516020919091036101000a9004949350505050565b6040805160208082528183019092526060916000918391602082018180388339019050509050836020820152600091505b6020821015613f4057808281518110613f1c57fe5b01602001516001600160f81b03191615613f3557613f40565b600190910190613f07565b6060826020036040519080825280601f01601f191660200182016040528015613f70576020820181803883390190505b50905060005b8151811015613fc6578251600185019484918110613f9057fe5b602001015160f81c60f81b828281518110613fa757fe5b60200101906001600160f81b031916908160001a905350600101613f76565b50949350505050565b6060600160401b8310613ffd5760405162461bcd60e51b8152600401613ff49061523c565b60405180910390fd5b604080516001808252818301909252606091602082018180388339019050509050603784116140575782840160f81b8160008151811061403957fe5b60200101906001600160f81b031916908160001a90535090506110c7565b606061406285613ed6565b90508381510160370160f81b8260008151811061407b57fe5b60200101906001600160f81b031916908160001a905350612ba482825b60608082518451016040519080825280601f01601f1916602001820160405280156140ca576020820181803883390190505b5090506000805b8551811015614120578581815181106140e657fe5b602001015160f81c60f81b8383815181106140fd57fe5b60200101906001600160f81b031916908160001a905350600191820191016140d1565b5060005b84518110156141735784818151811061413957fe5b602001015160f81c60f81b83838151811061415057fe5b60200101906001600160f81b031916908160001a90535060019182019101614124565b5090949350505050565b614185614553565b600061419083614233565b83519383529092016020820152919050565b6141aa614539565b602082015160006141ba8261429d565b828452602080850182905292019390910192909252919050565b805180516000918291821a9060808210156141f65792506001915061422e9050565b60b8821015614214576001856020015103925080600101935061422b565b602085015182820160b51901945082900360b60192505b50505b915091565b8051805160009190821a90608082101561425257600092505050610b62565b60b882108061426d575060c0821015801561426d575060f882105b1561427d57600192505050610b62565b60c0821015614292575060b519019050610b62565b5060f5190192915050565b8051600090811a60808110156142b657600191506142c7565b60b88110156142c757607e19810191505b50919050565b6040805161014081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c0810182905260e0810191909152610100810161431b614573565b81526020016143286145a8565b905290565b604080516101a081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c0810182905260e08101829052610100810182905261012081018290526101408101829052610160810182905261018081019190915290565b8154818355818111156143c5576006028160060283600052602060002091820191016143c591906145ec565b505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061440b57805160ff1916838001178555614438565b82800160010185558215614438579182015b8281111561443857825182559160200191906001019061441d565b5061444492915061464e565b5090565b8154818355818111156143c5576002028160020283600052602060002091820191016143c59190614668565b604080516101608101825260008082526020820181905291810182905260608082018390526080820183905260a0820183905260c0820183905260e08201839052610100820183905261012082019290925261014081019190915290565b60405180610120016040528060006001600160401b031681526020016000815260200160006001600160401b0316815260200160006001600160a01b0316815260200160008152602001606081526020016000815260200160008152602001600081525090565b604051806040016040528060008152602001600081525090565b6040518060400160405280614566614539565b8152602001600081525090565b6040805160c081018252600080825260208201819052918101829052606081018290526080810182905260a081019190915290565b6040805161010081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c0810182905260e081019190915290565b61147691905b808211156144445780546001600160e01b03191681556001810180546001600160a01b0319908116909155600282018054909116905560006003820181905560048201819055614645600583018261469e565b506006016145f2565b61147691905b808211156144445760008155600101614654565b61147691905b808211156144445780546001600160c81b03191681556001810180546001600160e01b031916905560020161466e565b50805460018160011615610100020316600290046000825580601f106146c45750610c08565b601f016020900490600052602060002090810190610c08919061464e565b80356110c781615663565b60008083601f8401126146ff57600080fd5b5081356001600160401b0381111561471657600080fd5b60208301915083602082028301111561472e57600080fd5b9250929050565b80356110c781615677565b80516110c781615677565b80356110c781615680565b60008083601f84011261476857600080fd5b5081356001600160401b0381111561477f57600080fd5b60208301915083600182028301111561472e57600080fd5b600082601f8301126147a857600080fd5b81356147bb6147b682615580565b61555a565b915080825260208301602083018583830111156147d757600080fd5b6147e2838284615621565b50505092915050565b6000602082840312156147fd57600080fd5b600061480984846146e2565b949350505050565b6000806040838503121561482457600080fd5b600061483085856146e2565b9250506020614841858286016146e2565b9150509250929050565b60008060006060848603121561486057600080fd5b600061486c86866146e2565b935050602061487d8682870161474b565b92505060408401356001600160401b0381111561489957600080fd5b6148a586828701614797565b9150509250925092565b6000602082840312156148c157600080fd5b60006148098484614740565b6000602082840312156148df57600080fd5b6000614809848461474b565b600080604083850312156148fe57600080fd5b600061490a858561474b565b925050602061484185828601614735565b6000806000806080858703121561493157600080fd5b600061493d878761474b565b945050602061494e8782880161474b565b935050604061495f8782880161474b565b92505060606149708782880161474b565b91505092959194509250565b60008060008060008060008060a0898b03121561499857600080fd5b60006149a48b8b61474b565b98505060208901356001600160401b038111156149c057600080fd5b6149cc8b828c01614756565b975097505060408901356001600160401b038111156149ea57600080fd5b6149f68b828c01614756565b95509550506060614a098b828c0161474b565b93505060808901356001600160401b03811115614a2557600080fd5b614a318b828c016146ed565b92509250509295985092959890939650565b60008060408385031215614a5657600080fd5b6000614a62858561474b565b92505060206148418582860161474b565b600080600080600060a08688031215614a8b57600080fd5b6000614a97888861474b565b9550506020614aa88882890161474b565b9450506040614ab98882890161474b565b9350506060614aca8882890161474b565b9250506080614adb8882890161474b565b9150509295509295909350565b600080600080600080600060a0888a031215614b0357600080fd5b6000614b0f8a8a61474b565b9750506020614b208a828b0161474b565b9650506040614b318a828b0161474b565b95505060608801356001600160401b03811115614b4d57600080fd5b614b598a828b01614756565b945094505060808801356001600160401b03811115614b7757600080fd5b614b838a828b01614756565b925092505092959891949750929550565b614b9d816155f5565b82525050565b614b9d816155b4565b614b9d816155bf565b614b9d81611476565b614b9d614bca82611476565b611476565b614b9d614bca826155c4565b6000614be6826155a7565b614bf081856155ab565b9350614c0081856020860161562d565b614c0981615659565b9093019392505050565b6000614c1e826155a7565b614c288185610b62565b9350614c3881856020860161562d565b9290920192915050565b614b9d81615600565b6000614c58603283610b62565b7f7375626d69744e52452875696e743235362c75696e743235362c627974657333815271322c627974657333322c627974657333322960701b602082015260320192915050565b6000614cac603c83610b62565b7f6170706c7952657175657374496e4368696c64436861696e28626f6f6c2c756981527f6e743235362c616464726573732c627974657333322c627974657329000000006020820152603c0192915050565b6000614d0b600e836155ab565b6d696e70757420746f6f206c6f6e6760901b815260200192915050565b6000614d35602a83610b62565b7f7375626d69744f52422875696e743235362c627974657333322c627974657333815269322c627974657333322960b01b6020820152602a0192915050565b6000614d81603b83610b62565b7f6170706c7952657175657374496e526f6f74436861696e28626f6f6c2c75696e81527f743235362c616464726573732c627974657333322c62797465732900000000006020820152603b0192915050565b80516102c0830190614de584826150b3565b506020820151614df860208501826150b3565b506040820151614e0b60408501826150b3565b506060820151614e1e6060850182614bac565b506080820151614e316080850182614bac565b5060a0820151614e4460a0850182614bac565b5060c0820151614e5760c0850182614bac565b5060e0820151614e6a60e0850182614bac565b50610100820151614e7f610100850182615031565b50610120820151614e946101c0850182614e9a565b50505050565b8051610100830190614eac8482614bb5565b506020820151614ebf6020850182614bb5565b506040820151614ed26040850182614bb5565b506060820151614ee560608501826150b3565b506080820151614ef860808501826150b3565b5060a0820151614f0b60a0850182614bac565b5060c0820151614f1e60c0850182614bac565b5060e0820151614e9460e0850182614bac565b80516101a0830190614f4384826150b3565b506020820151614f5660208501826150b3565b506040820151614f6960408501826150b3565b506060820151614f7c60608501826150b3565b506080820151614f8f60808501826150b3565b5060a0820151614fa260a0850182614bb5565b5060c0820151614fb560c0850182614bb5565b5060e0820151614fc860e0850182614bb5565b50610100820151614fdd610100850182614bac565b50610120820151614ff2610120850182614bac565b50610140820151615007610140850182614bac565b5061016082015161501c610160850182614bac565b50610180820151614e94610180850182614bac565b805160c083019061504284826150b3565b50602082015161505560208501826150b3565b50604082015161506860408501826150b3565b50606082015161507b60608501826150b3565b50608082015161508e60808501826150b3565b5060a0820151614e9460a08501826150b3565b614b9d816155d1565b614b9d81615616565b614b9d816155e9565b60006150c88285614bbe565b6020820191506150d88284614bbe565b5060200192915050565b60006150ee8285614bcf565b6004820191506148098284614c13565b6000610d5e8284614c13565b60006110c782614c4b565b60006110c782614c9f565b60006110c782614d28565b60006110c782614d74565b602081016110c78284614ba3565b604081016151528285614ba3565b610d5e6020830184614ba3565b602081016110c78284614bac565b60a0810161517b8288614bac565b6151886020830187614bb5565b6151956040830186614b94565b6151a26060830185614bb5565b8181036080830152610c468184614bdb565b60c081016151c28289614bac565b6151cf60208301886150b3565b6151dc60408301876150b3565b6151e960608301866150b3565b6151f660808301856150b3565b610c4660a0830184614ba3565b60a081016152118288614bb5565b61521e6020830187614bb5565b6151956040830186614bb5565b60208082528101610d5e8184614bdb565b602080825281016110c781614cfe565b6102c081016110c78284614dd3565b6101a081016110c78284614f31565b602081016110c78284614bb5565b6101008101615287828b614bb5565b615294602083018a614b94565b6152a16040830189614ba3565b6152ae6060830188614c42565b6152bb6080830187614bb5565b81810360a08301526152cd8186614bdb565b90506152dc60c0830185614bac565b6152e960e0830184614bac565b9998505050505050505050565b6101008101615305828b614bb5565b615312602083018a614b94565b61531f6040830189614ba3565b6152ae6060830188614bb5565b6040810161533a8285614bb5565b610d5e6020830184614bac565b608081016153558287614bb5565b6153626020830186614bb5565b61536f6040830185614bb5565b612ba46060830184614bb5565b6040810161538a8285614bb5565b610d5e6020830184614bb5565b60a081016153a58288614bb5565b6153b26020830187614bb5565b6153bf6040830186614bb5565b6153cc6060830185614bb5565b61235c6080830184614bb5565b608081016153e78287614bb5565b6153f46020830186614bb5565b61540160408301856150aa565b612ba460608301846150aa565b610160810161541d828e6150b3565b61542a602083018d614bac565b615437604083018c614bac565b615444606083018b614bac565b615451608083018a614bac565b61545e60a08301896150a1565b61546b60c0830188614ba3565b61547860e0830187614ba3565b615486610100830186614bb5565b615494610120830185614bb5565b8181036101408301526154a78184614bdb565b9d9c50505050505050505050505050565b61018081016154c7828f6150b3565b6154d4602083018e6150b3565b6154e1604083018d6150b3565b6154ee606083018c6150b3565b6154fb608083018b6150b3565b61550860a083018a6150b3565b61551560c08301896150b3565b61552260e08301886150b3565b6155306101008301876150b3565b61553e6101208301866150b3565b61554c6101408301856150b3565b6154a7610160830184614bac565b6040518181016001600160401b038111828210171561557857600080fd5b604052919050565b60006001600160401b0382111561559657600080fd5b506020601f91909101601f19160190565b5190565b90815260200190565b60006110c7826155dd565b151590565b6001600160e01b03191690565b6001600160801b031690565b6001600160a01b031690565b6001600160401b031690565b60006110c78261560b565b60006110c782611476565b60006110c7826155b4565b60006110c7826155e9565b82818337506000910152565b60005b83811015615648578181015183820152602001615630565b83811115614e945750506000910152565b601f01601f191690565b61566c816155b4565b8114610c0857600080fd5b61566c816155bf565b61566c8161147656fea365627a7a723158206eecf876ed164303d30cc733813ff412105ba8be23583e9e6eed2b5ae04b06016c6578706572696d656e74616cf564736f6c634300050c0040"

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
// Solidity: function getBlock(uint256 forkNumber, uint256 blockNumber) constant returns(DataPlasmaBlock)
func (_RootChain *RootChainCaller) GetBlock(opts *bind.CallOpts, forkNumber *big.Int, blockNumber *big.Int) (DataPlasmaBlock, error) {
	var (
		ret0 = new(DataPlasmaBlock)
	)
	out := ret0
	err := _RootChain.contract.Call(opts, out, "getBlock", forkNumber, blockNumber)
	return *ret0, err
}

// GetBlock is a free data retrieval call binding the contract method 0x4a44bdb8.
//
// Solidity: function getBlock(uint256 forkNumber, uint256 blockNumber) constant returns(DataPlasmaBlock)
func (_RootChain *RootChainSession) GetBlock(forkNumber *big.Int, blockNumber *big.Int) (DataPlasmaBlock, error) {
	return _RootChain.Contract.GetBlock(&_RootChain.CallOpts, forkNumber, blockNumber)
}

// GetBlock is a free data retrieval call binding the contract method 0x4a44bdb8.
//
// Solidity: function getBlock(uint256 forkNumber, uint256 blockNumber) constant returns(DataPlasmaBlock)
func (_RootChain *RootChainCallerSession) GetBlock(forkNumber *big.Int, blockNumber *big.Int) (DataPlasmaBlock, error) {
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
// Solidity: function getEpoch(uint256 forkNumber, uint256 epochNumber) constant returns(DataEpoch epoch)
func (_RootChain *RootChainCaller) GetEpoch(opts *bind.CallOpts, forkNumber *big.Int, epochNumber *big.Int) (DataEpoch, error) {
	var (
		ret0 = new(DataEpoch)
	)
	out := ret0
	err := _RootChain.contract.Call(opts, out, "getEpoch", forkNumber, epochNumber)
	return *ret0, err
}

// GetEpoch is a free data retrieval call binding the contract method 0x2b25a38b.
//
// Solidity: function getEpoch(uint256 forkNumber, uint256 epochNumber) constant returns(DataEpoch epoch)
func (_RootChain *RootChainSession) GetEpoch(forkNumber *big.Int, epochNumber *big.Int) (DataEpoch, error) {
	return _RootChain.Contract.GetEpoch(&_RootChain.CallOpts, forkNumber, epochNumber)
}

// GetEpoch is a free data retrieval call binding the contract method 0x2b25a38b.
//
// Solidity: function getEpoch(uint256 forkNumber, uint256 epochNumber) constant returns(DataEpoch epoch)
func (_RootChain *RootChainCallerSession) GetEpoch(forkNumber *big.Int, epochNumber *big.Int) (DataEpoch, error) {
	return _RootChain.Contract.GetEpoch(&_RootChain.CallOpts, forkNumber, epochNumber)
}

// GetLastEpoch is a free data retrieval call binding the contract method 0x398bac63.
//
// Solidity: function getLastEpoch() constant returns(DataEpoch)
func (_RootChain *RootChainCaller) GetLastEpoch(opts *bind.CallOpts) (DataEpoch, error) {
	var (
		ret0 = new(DataEpoch)
	)
	out := ret0
	err := _RootChain.contract.Call(opts, out, "getLastEpoch")
	return *ret0, err
}

// GetLastEpoch is a free data retrieval call binding the contract method 0x398bac63.
//
// Solidity: function getLastEpoch() constant returns(DataEpoch)
func (_RootChain *RootChainSession) GetLastEpoch() (DataEpoch, error) {
	return _RootChain.Contract.GetLastEpoch(&_RootChain.CallOpts)
}

// GetLastEpoch is a free data retrieval call binding the contract method 0x398bac63.
//
// Solidity: function getLastEpoch() constant returns(DataEpoch)
func (_RootChain *RootChainCallerSession) GetLastEpoch() (DataEpoch, error) {
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

// IsMapper is a free data retrieval call binding the contract method 0xc0c49c2a.
//
// Solidity: function isMapper(address account) constant returns(bool)
func (_RootChain *RootChainCaller) IsMapper(opts *bind.CallOpts, account common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _RootChain.contract.Call(opts, out, "isMapper", account)
	return *ret0, err
}

// IsMapper is a free data retrieval call binding the contract method 0xc0c49c2a.
//
// Solidity: function isMapper(address account) constant returns(bool)
func (_RootChain *RootChainSession) IsMapper(account common.Address) (bool, error) {
	return _RootChain.Contract.IsMapper(&_RootChain.CallOpts, account)
}

// IsMapper is a free data retrieval call binding the contract method 0xc0c49c2a.
//
// Solidity: function isMapper(address account) constant returns(bool)
func (_RootChain *RootChainCallerSession) IsMapper(account common.Address) (bool, error) {
	return _RootChain.Contract.IsMapper(&_RootChain.CallOpts, account)
}

// IsRootChain is a free data retrieval call binding the contract method 0x420bb4b8.
//
// Solidity: function isRootChain() constant returns(bool)
func (_RootChain *RootChainCaller) IsRootChain(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _RootChain.contract.Call(opts, out, "isRootChain")
	return *ret0, err
}

// IsRootChain is a free data retrieval call binding the contract method 0x420bb4b8.
//
// Solidity: function isRootChain() constant returns(bool)
func (_RootChain *RootChainSession) IsRootChain() (bool, error) {
	return _RootChain.Contract.IsRootChain(&_RootChain.CallOpts)
}

// IsRootChain is a free data retrieval call binding the contract method 0x420bb4b8.
//
// Solidity: function isRootChain() constant returns(bool)
func (_RootChain *RootChainCallerSession) IsRootChain() (bool, error) {
	return _RootChain.Contract.IsRootChain(&_RootChain.CallOpts)
}

// IsSubmitter is a free data retrieval call binding the contract method 0xa926fdbc.
//
// Solidity: function isSubmitter(address account) constant returns(bool)
func (_RootChain *RootChainCaller) IsSubmitter(opts *bind.CallOpts, account common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _RootChain.contract.Call(opts, out, "isSubmitter", account)
	return *ret0, err
}

// IsSubmitter is a free data retrieval call binding the contract method 0xa926fdbc.
//
// Solidity: function isSubmitter(address account) constant returns(bool)
func (_RootChain *RootChainSession) IsSubmitter(account common.Address) (bool, error) {
	return _RootChain.Contract.IsSubmitter(&_RootChain.CallOpts, account)
}

// IsSubmitter is a free data retrieval call binding the contract method 0xa926fdbc.
//
// Solidity: function isSubmitter(address account) constant returns(bool)
func (_RootChain *RootChainCallerSession) IsSubmitter(account common.Address) (bool, error) {
	return _RootChain.Contract.IsSubmitter(&_RootChain.CallOpts, account)
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

// SeigManager is a free data retrieval call binding the contract method 0x6fb7f558.
//
// Solidity: function seigManager() constant returns(address)
func (_RootChain *RootChainCaller) SeigManager(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RootChain.contract.Call(opts, out, "seigManager")
	return *ret0, err
}

// SeigManager is a free data retrieval call binding the contract method 0x6fb7f558.
//
// Solidity: function seigManager() constant returns(address)
func (_RootChain *RootChainSession) SeigManager() (common.Address, error) {
	return _RootChain.Contract.SeigManager(&_RootChain.CallOpts)
}

// SeigManager is a free data retrieval call binding the contract method 0x6fb7f558.
//
// Solidity: function seigManager() constant returns(address)
func (_RootChain *RootChainCallerSession) SeigManager() (common.Address, error) {
	return _RootChain.Contract.SeigManager(&_RootChain.CallOpts)
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

// AddMapper is a paid mutator transaction binding the contract method 0x80e3e81d.
//
// Solidity: function addMapper(address account) returns()
func (_RootChain *RootChainTransactor) AddMapper(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _RootChain.contract.Transact(opts, "addMapper", account)
}

// AddMapper is a paid mutator transaction binding the contract method 0x80e3e81d.
//
// Solidity: function addMapper(address account) returns()
func (_RootChain *RootChainSession) AddMapper(account common.Address) (*types.Transaction, error) {
	return _RootChain.Contract.AddMapper(&_RootChain.TransactOpts, account)
}

// AddMapper is a paid mutator transaction binding the contract method 0x80e3e81d.
//
// Solidity: function addMapper(address account) returns()
func (_RootChain *RootChainTransactorSession) AddMapper(account common.Address) (*types.Transaction, error) {
	return _RootChain.Contract.AddMapper(&_RootChain.TransactOpts, account)
}

// AddSubmitter is a paid mutator transaction binding the contract method 0x072900f9.
//
// Solidity: function addSubmitter(address account) returns()
func (_RootChain *RootChainTransactor) AddSubmitter(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _RootChain.contract.Transact(opts, "addSubmitter", account)
}

// AddSubmitter is a paid mutator transaction binding the contract method 0x072900f9.
//
// Solidity: function addSubmitter(address account) returns()
func (_RootChain *RootChainSession) AddSubmitter(account common.Address) (*types.Transaction, error) {
	return _RootChain.Contract.AddSubmitter(&_RootChain.TransactOpts, account)
}

// AddSubmitter is a paid mutator transaction binding the contract method 0x072900f9.
//
// Solidity: function addSubmitter(address account) returns()
func (_RootChain *RootChainTransactorSession) AddSubmitter(account common.Address) (*types.Transaction, error) {
	return _RootChain.Contract.AddSubmitter(&_RootChain.TransactOpts, account)
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

// ChangeOperator is a paid mutator transaction binding the contract method 0x06394c9b.
//
// Solidity: function changeOperator(address _operator) returns()
func (_RootChain *RootChainTransactor) ChangeOperator(opts *bind.TransactOpts, _operator common.Address) (*types.Transaction, error) {
	return _RootChain.contract.Transact(opts, "changeOperator", _operator)
}

// ChangeOperator is a paid mutator transaction binding the contract method 0x06394c9b.
//
// Solidity: function changeOperator(address _operator) returns()
func (_RootChain *RootChainSession) ChangeOperator(_operator common.Address) (*types.Transaction, error) {
	return _RootChain.Contract.ChangeOperator(&_RootChain.TransactOpts, _operator)
}

// ChangeOperator is a paid mutator transaction binding the contract method 0x06394c9b.
//
// Solidity: function changeOperator(address _operator) returns()
func (_RootChain *RootChainTransactorSession) ChangeOperator(_operator common.Address) (*types.Transaction, error) {
	return _RootChain.Contract.ChangeOperator(&_RootChain.TransactOpts, _operator)
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

// RenounceMapper is a paid mutator transaction binding the contract method 0x3565fb0d.
//
// Solidity: function renounceMapper() returns()
func (_RootChain *RootChainTransactor) RenounceMapper(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RootChain.contract.Transact(opts, "renounceMapper")
}

// RenounceMapper is a paid mutator transaction binding the contract method 0x3565fb0d.
//
// Solidity: function renounceMapper() returns()
func (_RootChain *RootChainSession) RenounceMapper() (*types.Transaction, error) {
	return _RootChain.Contract.RenounceMapper(&_RootChain.TransactOpts)
}

// RenounceMapper is a paid mutator transaction binding the contract method 0x3565fb0d.
//
// Solidity: function renounceMapper() returns()
func (_RootChain *RootChainTransactorSession) RenounceMapper() (*types.Transaction, error) {
	return _RootChain.Contract.RenounceMapper(&_RootChain.TransactOpts)
}

// RenounceSubmitter is a paid mutator transaction binding the contract method 0x5e0ca71b.
//
// Solidity: function renounceSubmitter() returns()
func (_RootChain *RootChainTransactor) RenounceSubmitter(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RootChain.contract.Transact(opts, "renounceSubmitter")
}

// RenounceSubmitter is a paid mutator transaction binding the contract method 0x5e0ca71b.
//
// Solidity: function renounceSubmitter() returns()
func (_RootChain *RootChainSession) RenounceSubmitter() (*types.Transaction, error) {
	return _RootChain.Contract.RenounceSubmitter(&_RootChain.TransactOpts)
}

// RenounceSubmitter is a paid mutator transaction binding the contract method 0x5e0ca71b.
//
// Solidity: function renounceSubmitter() returns()
func (_RootChain *RootChainTransactorSession) RenounceSubmitter() (*types.Transaction, error) {
	return _RootChain.Contract.RenounceSubmitter(&_RootChain.TransactOpts)
}

// SetSeigManager is a paid mutator transaction binding the contract method 0x7657f20a.
//
// Solidity: function setSeigManager(address account) returns()
func (_RootChain *RootChainTransactor) SetSeigManager(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _RootChain.contract.Transact(opts, "setSeigManager", account)
}

// SetSeigManager is a paid mutator transaction binding the contract method 0x7657f20a.
//
// Solidity: function setSeigManager(address account) returns()
func (_RootChain *RootChainSession) SetSeigManager(account common.Address) (*types.Transaction, error) {
	return _RootChain.Contract.SetSeigManager(&_RootChain.TransactOpts, account)
}

// SetSeigManager is a paid mutator transaction binding the contract method 0x7657f20a.
//
// Solidity: function setSeigManager(address account) returns()
func (_RootChain *RootChainTransactorSession) SetSeigManager(account common.Address) (*types.Transaction, error) {
	return _RootChain.Contract.SetSeigManager(&_RootChain.TransactOpts, account)
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

// RootChainMapperAddedIterator is returned from FilterMapperAdded and is used to iterate over the raw logs and unpacked data for MapperAdded events raised by the RootChain contract.
type RootChainMapperAddedIterator struct {
	Event *RootChainMapperAdded // Event containing the contract specifics and raw log

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
func (it *RootChainMapperAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RootChainMapperAdded)
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
		it.Event = new(RootChainMapperAdded)
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
func (it *RootChainMapperAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RootChainMapperAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RootChainMapperAdded represents a MapperAdded event raised by the RootChain contract.
type RootChainMapperAdded struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMapperAdded is a free log retrieval operation binding the contract event 0xa199526a01dbcba68aa3ab9d7d06c6692c83fb1a8bcf2184e6d94ad34f5aaf95.
//
// Solidity: event MapperAdded(address indexed account)
func (_RootChain *RootChainFilterer) FilterMapperAdded(opts *bind.FilterOpts, account []common.Address) (*RootChainMapperAddedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _RootChain.contract.FilterLogs(opts, "MapperAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return &RootChainMapperAddedIterator{contract: _RootChain.contract, event: "MapperAdded", logs: logs, sub: sub}, nil
}

// WatchMapperAdded is a free log subscription operation binding the contract event 0xa199526a01dbcba68aa3ab9d7d06c6692c83fb1a8bcf2184e6d94ad34f5aaf95.
//
// Solidity: event MapperAdded(address indexed account)
func (_RootChain *RootChainFilterer) WatchMapperAdded(opts *bind.WatchOpts, sink chan<- *RootChainMapperAdded, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _RootChain.contract.WatchLogs(opts, "MapperAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RootChainMapperAdded)
				if err := _RootChain.contract.UnpackLog(event, "MapperAdded", log); err != nil {
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

// ParseMapperAdded is a log parse operation binding the contract event 0xa199526a01dbcba68aa3ab9d7d06c6692c83fb1a8bcf2184e6d94ad34f5aaf95.
//
// Solidity: event MapperAdded(address indexed account)
func (_RootChain *RootChainFilterer) ParseMapperAdded(log types.Log) (*RootChainMapperAdded, error) {
	event := new(RootChainMapperAdded)
	if err := _RootChain.contract.UnpackLog(event, "MapperAdded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RootChainMapperRemovedIterator is returned from FilterMapperRemoved and is used to iterate over the raw logs and unpacked data for MapperRemoved events raised by the RootChain contract.
type RootChainMapperRemovedIterator struct {
	Event *RootChainMapperRemoved // Event containing the contract specifics and raw log

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
func (it *RootChainMapperRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RootChainMapperRemoved)
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
		it.Event = new(RootChainMapperRemoved)
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
func (it *RootChainMapperRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RootChainMapperRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RootChainMapperRemoved represents a MapperRemoved event raised by the RootChain contract.
type RootChainMapperRemoved struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMapperRemoved is a free log retrieval operation binding the contract event 0xf012269e1b6bb33c7800e7e54167ddd74f0ae7244108030b6876821d3779822f.
//
// Solidity: event MapperRemoved(address indexed account)
func (_RootChain *RootChainFilterer) FilterMapperRemoved(opts *bind.FilterOpts, account []common.Address) (*RootChainMapperRemovedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _RootChain.contract.FilterLogs(opts, "MapperRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return &RootChainMapperRemovedIterator{contract: _RootChain.contract, event: "MapperRemoved", logs: logs, sub: sub}, nil
}

// WatchMapperRemoved is a free log subscription operation binding the contract event 0xf012269e1b6bb33c7800e7e54167ddd74f0ae7244108030b6876821d3779822f.
//
// Solidity: event MapperRemoved(address indexed account)
func (_RootChain *RootChainFilterer) WatchMapperRemoved(opts *bind.WatchOpts, sink chan<- *RootChainMapperRemoved, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _RootChain.contract.WatchLogs(opts, "MapperRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RootChainMapperRemoved)
				if err := _RootChain.contract.UnpackLog(event, "MapperRemoved", log); err != nil {
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

// ParseMapperRemoved is a log parse operation binding the contract event 0xf012269e1b6bb33c7800e7e54167ddd74f0ae7244108030b6876821d3779822f.
//
// Solidity: event MapperRemoved(address indexed account)
func (_RootChain *RootChainFilterer) ParseMapperRemoved(log types.Log) (*RootChainMapperRemoved, error) {
	event := new(RootChainMapperRemoved)
	if err := _RootChain.contract.UnpackLog(event, "MapperRemoved", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RootChainOperatorChangedIterator is returned from FilterOperatorChanged and is used to iterate over the raw logs and unpacked data for OperatorChanged events raised by the RootChain contract.
type RootChainOperatorChangedIterator struct {
	Event *RootChainOperatorChanged // Event containing the contract specifics and raw log

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
func (it *RootChainOperatorChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RootChainOperatorChanged)
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
		it.Event = new(RootChainOperatorChanged)
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
func (it *RootChainOperatorChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RootChainOperatorChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RootChainOperatorChanged represents a OperatorChanged event raised by the RootChain contract.
type RootChainOperatorChanged struct {
	NewOperator common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOperatorChanged is a free log retrieval operation binding the contract event 0x4721129e0e676ed6a92909bb24e853ccdd63ad72280cc2e974e38e480e0e6e54.
//
// Solidity: event OperatorChanged(address _newOperator)
func (_RootChain *RootChainFilterer) FilterOperatorChanged(opts *bind.FilterOpts) (*RootChainOperatorChangedIterator, error) {

	logs, sub, err := _RootChain.contract.FilterLogs(opts, "OperatorChanged")
	if err != nil {
		return nil, err
	}
	return &RootChainOperatorChangedIterator{contract: _RootChain.contract, event: "OperatorChanged", logs: logs, sub: sub}, nil
}

// WatchOperatorChanged is a free log subscription operation binding the contract event 0x4721129e0e676ed6a92909bb24e853ccdd63ad72280cc2e974e38e480e0e6e54.
//
// Solidity: event OperatorChanged(address _newOperator)
func (_RootChain *RootChainFilterer) WatchOperatorChanged(opts *bind.WatchOpts, sink chan<- *RootChainOperatorChanged) (event.Subscription, error) {

	logs, sub, err := _RootChain.contract.WatchLogs(opts, "OperatorChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RootChainOperatorChanged)
				if err := _RootChain.contract.UnpackLog(event, "OperatorChanged", log); err != nil {
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
func (_RootChain *RootChainFilterer) ParseOperatorChanged(log types.Log) (*RootChainOperatorChanged, error) {
	event := new(RootChainOperatorChanged)
	if err := _RootChain.contract.UnpackLog(event, "OperatorChanged", log); err != nil {
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

// RootChainSubmitterAddedIterator is returned from FilterSubmitterAdded and is used to iterate over the raw logs and unpacked data for SubmitterAdded events raised by the RootChain contract.
type RootChainSubmitterAddedIterator struct {
	Event *RootChainSubmitterAdded // Event containing the contract specifics and raw log

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
func (it *RootChainSubmitterAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RootChainSubmitterAdded)
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
		it.Event = new(RootChainSubmitterAdded)
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
func (it *RootChainSubmitterAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RootChainSubmitterAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RootChainSubmitterAdded represents a SubmitterAdded event raised by the RootChain contract.
type RootChainSubmitterAdded struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSubmitterAdded is a free log retrieval operation binding the contract event 0xb079bc2cbde1f186e0b351d4a87c4597e3ed098f571548617449e73506428d8b.
//
// Solidity: event SubmitterAdded(address indexed account)
func (_RootChain *RootChainFilterer) FilterSubmitterAdded(opts *bind.FilterOpts, account []common.Address) (*RootChainSubmitterAddedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _RootChain.contract.FilterLogs(opts, "SubmitterAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return &RootChainSubmitterAddedIterator{contract: _RootChain.contract, event: "SubmitterAdded", logs: logs, sub: sub}, nil
}

// WatchSubmitterAdded is a free log subscription operation binding the contract event 0xb079bc2cbde1f186e0b351d4a87c4597e3ed098f571548617449e73506428d8b.
//
// Solidity: event SubmitterAdded(address indexed account)
func (_RootChain *RootChainFilterer) WatchSubmitterAdded(opts *bind.WatchOpts, sink chan<- *RootChainSubmitterAdded, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _RootChain.contract.WatchLogs(opts, "SubmitterAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RootChainSubmitterAdded)
				if err := _RootChain.contract.UnpackLog(event, "SubmitterAdded", log); err != nil {
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

// ParseSubmitterAdded is a log parse operation binding the contract event 0xb079bc2cbde1f186e0b351d4a87c4597e3ed098f571548617449e73506428d8b.
//
// Solidity: event SubmitterAdded(address indexed account)
func (_RootChain *RootChainFilterer) ParseSubmitterAdded(log types.Log) (*RootChainSubmitterAdded, error) {
	event := new(RootChainSubmitterAdded)
	if err := _RootChain.contract.UnpackLog(event, "SubmitterAdded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RootChainSubmitterRemovedIterator is returned from FilterSubmitterRemoved and is used to iterate over the raw logs and unpacked data for SubmitterRemoved events raised by the RootChain contract.
type RootChainSubmitterRemovedIterator struct {
	Event *RootChainSubmitterRemoved // Event containing the contract specifics and raw log

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
func (it *RootChainSubmitterRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RootChainSubmitterRemoved)
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
		it.Event = new(RootChainSubmitterRemoved)
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
func (it *RootChainSubmitterRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RootChainSubmitterRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RootChainSubmitterRemoved represents a SubmitterRemoved event raised by the RootChain contract.
type RootChainSubmitterRemoved struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSubmitterRemoved is a free log retrieval operation binding the contract event 0xf84a004e1673d2f349a7c93c72b3794b8eba6d2f9338044d8c8cd260e51a57a1.
//
// Solidity: event SubmitterRemoved(address indexed account)
func (_RootChain *RootChainFilterer) FilterSubmitterRemoved(opts *bind.FilterOpts, account []common.Address) (*RootChainSubmitterRemovedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _RootChain.contract.FilterLogs(opts, "SubmitterRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return &RootChainSubmitterRemovedIterator{contract: _RootChain.contract, event: "SubmitterRemoved", logs: logs, sub: sub}, nil
}

// WatchSubmitterRemoved is a free log subscription operation binding the contract event 0xf84a004e1673d2f349a7c93c72b3794b8eba6d2f9338044d8c8cd260e51a57a1.
//
// Solidity: event SubmitterRemoved(address indexed account)
func (_RootChain *RootChainFilterer) WatchSubmitterRemoved(opts *bind.WatchOpts, sink chan<- *RootChainSubmitterRemoved, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _RootChain.contract.WatchLogs(opts, "SubmitterRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RootChainSubmitterRemoved)
				if err := _RootChain.contract.UnpackLog(event, "SubmitterRemoved", log); err != nil {
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

// ParseSubmitterRemoved is a log parse operation binding the contract event 0xf84a004e1673d2f349a7c93c72b3794b8eba6d2f9338044d8c8cd260e51a57a1.
//
// Solidity: event SubmitterRemoved(address indexed account)
func (_RootChain *RootChainFilterer) ParseSubmitterRemoved(log types.Log) (*RootChainSubmitterRemoved, error) {
	event := new(RootChainSubmitterRemoved)
	if err := _RootChain.contract.UnpackLog(event, "SubmitterRemoved", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RootChainBaseABI is the input ABI used to generate the binding from.
const RootChainBaseABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"forkNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"BlockFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fork\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epochNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isRequest\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"userActivated\",\"type\":\"bool\"}],\"name\":\"BlockSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"requestor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"trieKey\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"trieValue\",\"type\":\"bytes32\"}],\"name\":\"ERUCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"forkNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epochNumber\",\"type\":\"uint256\"}],\"name\":\"EpochFilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"forkNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epochNumber\",\"type\":\"uint256\"}],\"name\":\"EpochFilling\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"forkNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epochNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startBlockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"endBlockNumber\",\"type\":\"uint256\"}],\"name\":\"EpochFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"forkNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epochNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startBlockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"endBlockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requestStart\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requestEnd\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"epochIsEmpty\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isRequest\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"userActivated\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"rebase\",\"type\":\"bool\"}],\"name\":\"EpochPrepared\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"forkNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epochNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startBlockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"endBlockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requestStart\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requestEnd\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"epochIsEmpty\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isRequest\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"userActivated\",\"type\":\"bool\"}],\"name\":\"EpochRebased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newFork\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epochNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"forkedBlockNumber\",\"type\":\"uint256\"}],\"name\":\"Forked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_newOperator\",\"type\":\"address\"}],\"name\":\"OperatorChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"userActivated\",\"type\":\"bool\"}],\"name\":\"RequestApplied\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"userActivated\",\"type\":\"bool\"}],\"name\":\"RequestChallenged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"requestor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"weiAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"trieKey\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"trieValue\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isExit\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"userActivated\",\"type\":\"bool\"}],\"name\":\"RequestCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"userActivated\",\"type\":\"bool\"}],\"name\":\"RequestFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"contractInRootchain\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"contractInChildchain\",\"type\":\"address\"}],\"name\":\"RequestableContractMapped\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"userActivated\",\"type\":\"bool\"}],\"name\":\"SessionTimeout\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"COST_ERO\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"COST_ERU\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"COST_NRB\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"COST_ORB\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"COST_URB\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"COST_URB_PREPARE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CP_COMPUTATION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CP_EXIT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CP_WITHHOLDING\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"EROIdToFinalize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"EROs\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"isExit\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isTransfer\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"finalized\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"challenged\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"value\",\"type\":\"uint128\"},{\"internalType\":\"addresspayable\",\"name\":\"requestor\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"trieKey\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"trieValue\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ERUIdToFinalize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"ERUs\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"isExit\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isTransfer\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"finalized\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"challenged\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"value\",\"type\":\"uint128\"},{\"internalType\":\"addresspayable\",\"name\":\"requestor\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"trieKey\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"trieValue\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"NRELength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"NULL_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"ORBs\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"submitted\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"numEnter\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"epochNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"requestStart\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"requestEnd\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"trie\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PREPARE_TIMEOUT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"REQUEST_GAS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"URBs\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"submitted\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"numEnter\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"epochNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"requestStart\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"requestEnd\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"trie\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"currentFork\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"development\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"epochHandler\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"etherToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"firstFilledORENumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"firstNonEmptyRequestEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"forks\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"forkedBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"firstEpoch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"lastEpoch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"firstBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"lastBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"lastFinalizedEpoch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"lastFinalizedBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"firstEnterEpoch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"lastEnterEpoch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nextBlockToRebase\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"rebased\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isRootChain\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastAppliedBlockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastAppliedEpochNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastAppliedForkNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"lastNonEmptyRequestEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"numEnterForORB\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"operator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"requestableContracts\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"seigManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"submitHandler\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

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

// RootChainBaseBin is the compiled bytecode used for deploying new contracts.
var RootChainBaseBin = "0x608060405234801561001057600080fd5b50610927806100206000396000f3fe608060405234801561001057600080fd5b506004361061021c5760003560e01c8063b17fa6e911610125578063ca6f6380116100ad578063e259faf71161007c578063e259faf714610585578063e7b88b801461058d578063ea7f22a814610595578063f4f31de4146105b2578063fb788a27146105cf5761021c565b8063ca6f63801461053a578063d691acd814610221578063da0185f814610557578063de0ce17d1461057d5761021c565b8063b8066bcb116100f4578063b8066bcb146104b1578063c0e86064146104b9578063c2bc88fa14610522578063c54626cc1461052a578063c8ad329f146105325761021c565b8063b17fa6e91461037c578063b2ae9ba814610221578063b443f3cc14610384578063b6715647146104945761021c565b8063570ca735116101a85780638155717d116101775780638155717d146103645780638b5172d0146102215780638eb288ca1461036c57806394be3aa514610221578063ab96da2d146103745761021c565b8063570ca735146103135780636fb7f5581461033757806372ecb9a81461033f5780637b929c271461035c5761021c565b8063192adc5b116101ef578063192adc5b1461022157806323691566146102535780632dc6bb7b1461025b578063420bb4b8146102635780634ba3a1261461027f5761021c565b8063033cfbed1461022157806308c4fff01461023b578063164bc2ae14610243578063183d2d1c1461024b575b600080fd5b6102296105d7565b60408051918252519081900360200190f35b6102296105dc565b6102296105e1565b6102296105e7565b6102296105ed565b6102296105f3565b61026b6105f9565b604080519115158252519081900360200190f35b61029c6004803603602081101561029557600080fd5b50356105fe565b6040805167ffffffffffffffff9d8e1681529b8d1660208d0152998c168b8b0152978b1660608b0152958a1660808a015293891660a089015291881660c0880152871660e0870152861661010086015285166101208501529093166101408301529115156101608201529051908190036101800190f35b61031b610676565b604080516001600160a01b039092168252519081900360200190f35b61031b61068a565b6102296004803603602081101561035557600080fd5b5035610699565b61026b6106ab565b6102296106b4565b6102296106b9565b6102296106c0565b6102296106c6565b6103a16004803603602081101561039a57600080fd5b50356106cb565b6040805167ffffffffffffffff8d1681528b15156020808301919091528b151592820192909252891515606082015288151560808201526001600160801b03881660a08201526001600160a01b0380881660c0830152861660e082015261010081018590526101208101849052610160610140820181815284519183019190915283519192909161018084019185019080838360005b8381101561044f578181015183820152602001610437565b50505050905090810190601f16801561047c5780820380516001836020036101000a031916815260200191505b509c5050505050505050505050505060405180910390f35b610229600480360360208110156104aa57600080fd5b50356107ed565b61031b6107ff565b6104d6600480360360208110156104cf57600080fd5b503561080e565b60408051961515875267ffffffffffffffff958616602088015293851686850152918416606086015290921660808401526001600160a01b0390911660a0830152519081900360c00190f35b610229610876565b61022961087b565b610229610881565b6102296004803603602081101561055057600080fd5b5035610887565b61031b6004803603602081101561056d57600080fd5b50356001600160a01b0316610899565b61031b6105d7565b61031b6108b4565b61031b6108c3565b6104d6600480360360208110156105ab57600080fd5b50356108d2565b6103a1600480360360208110156105c857600080fd5b50356108df565b6102296108ec565b600081565b600f81565b60105481565b60065481565b600d5481565b60135481565b600181565b60086020526000908152604090208054600182015460029092015467ffffffffffffffff80831693600160401b808504831694600160801b808204851695600160c01b92839004861695858116958581048216958482048316959182900483169484841694918204841693908204169160ff9104168c565b60005461010090046001600160a01b031681565b6004546001600160a01b031681565b60076020526000908152604090205481565b60005460ff1681565b600a81565b620186a081565b60055481565b601481565b600981815481106106d857fe5b6000918252602091829020600691909102018054600180830154600280850154600386015460048701546005880180546040805161010099831615999099026000190190911695909504601f81018b90048b0288018b0190955284875267ffffffffffffffff88169a50600160401b880460ff9081169a600160481b8a0482169a600160501b8b0483169a600160581b810490931699600160601b9093046001600160801b0316986001600160a01b03908116989716969093918301828280156107e35780601f106107b8576101008083540402835291602001916107e3565b820191906000526020600020905b8154815290600101906020018083116107c657829003601f168201915b505050505090508b565b600e6020526000908152604090205481565b6003546001600160a01b031681565b600c818154811061081b57fe5b60009182526020909120600290910201805460019091015460ff8216925067ffffffffffffffff6101008304811692600160481b8104821692600160881b9091048216918116906001600160a01b03600160401b9091041686565b603c81565b60145481565b60115481565b600f6020526000908152604090205481565b6015602052600090815260409020546001600160a01b031681565b6002546001600160a01b031681565b6001546001600160a01b031681565b600b818154811061081b57fe5b600a81815481106106d857fe5b6012548156fea265627a7a72315820caa1bece88c45e45da9bb43f2c28f1168891cf641eeeb099e1f946f89f73c30f64736f6c634300050c0032"

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

// IsRootChain is a free data retrieval call binding the contract method 0x420bb4b8.
//
// Solidity: function isRootChain() constant returns(bool)
func (_RootChainBase *RootChainBaseCaller) IsRootChain(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _RootChainBase.contract.Call(opts, out, "isRootChain")
	return *ret0, err
}

// IsRootChain is a free data retrieval call binding the contract method 0x420bb4b8.
//
// Solidity: function isRootChain() constant returns(bool)
func (_RootChainBase *RootChainBaseSession) IsRootChain() (bool, error) {
	return _RootChainBase.Contract.IsRootChain(&_RootChainBase.CallOpts)
}

// IsRootChain is a free data retrieval call binding the contract method 0x420bb4b8.
//
// Solidity: function isRootChain() constant returns(bool)
func (_RootChainBase *RootChainBaseCallerSession) IsRootChain() (bool, error) {
	return _RootChainBase.Contract.IsRootChain(&_RootChainBase.CallOpts)
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

// SeigManager is a free data retrieval call binding the contract method 0x6fb7f558.
//
// Solidity: function seigManager() constant returns(address)
func (_RootChainBase *RootChainBaseCaller) SeigManager(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RootChainBase.contract.Call(opts, out, "seigManager")
	return *ret0, err
}

// SeigManager is a free data retrieval call binding the contract method 0x6fb7f558.
//
// Solidity: function seigManager() constant returns(address)
func (_RootChainBase *RootChainBaseSession) SeigManager() (common.Address, error) {
	return _RootChainBase.Contract.SeigManager(&_RootChainBase.CallOpts)
}

// SeigManager is a free data retrieval call binding the contract method 0x6fb7f558.
//
// Solidity: function seigManager() constant returns(address)
func (_RootChainBase *RootChainBaseCallerSession) SeigManager() (common.Address, error) {
	return _RootChainBase.Contract.SeigManager(&_RootChainBase.CallOpts)
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

// RootChainBaseOperatorChangedIterator is returned from FilterOperatorChanged and is used to iterate over the raw logs and unpacked data for OperatorChanged events raised by the RootChainBase contract.
type RootChainBaseOperatorChangedIterator struct {
	Event *RootChainBaseOperatorChanged // Event containing the contract specifics and raw log

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
func (it *RootChainBaseOperatorChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RootChainBaseOperatorChanged)
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
		it.Event = new(RootChainBaseOperatorChanged)
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
func (it *RootChainBaseOperatorChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RootChainBaseOperatorChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RootChainBaseOperatorChanged represents a OperatorChanged event raised by the RootChainBase contract.
type RootChainBaseOperatorChanged struct {
	NewOperator common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOperatorChanged is a free log retrieval operation binding the contract event 0x4721129e0e676ed6a92909bb24e853ccdd63ad72280cc2e974e38e480e0e6e54.
//
// Solidity: event OperatorChanged(address _newOperator)
func (_RootChainBase *RootChainBaseFilterer) FilterOperatorChanged(opts *bind.FilterOpts) (*RootChainBaseOperatorChangedIterator, error) {

	logs, sub, err := _RootChainBase.contract.FilterLogs(opts, "OperatorChanged")
	if err != nil {
		return nil, err
	}
	return &RootChainBaseOperatorChangedIterator{contract: _RootChainBase.contract, event: "OperatorChanged", logs: logs, sub: sub}, nil
}

// WatchOperatorChanged is a free log subscription operation binding the contract event 0x4721129e0e676ed6a92909bb24e853ccdd63ad72280cc2e974e38e480e0e6e54.
//
// Solidity: event OperatorChanged(address _newOperator)
func (_RootChainBase *RootChainBaseFilterer) WatchOperatorChanged(opts *bind.WatchOpts, sink chan<- *RootChainBaseOperatorChanged) (event.Subscription, error) {

	logs, sub, err := _RootChainBase.contract.WatchLogs(opts, "OperatorChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RootChainBaseOperatorChanged)
				if err := _RootChainBase.contract.UnpackLog(event, "OperatorChanged", log); err != nil {
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
func (_RootChainBase *RootChainBaseFilterer) ParseOperatorChanged(log types.Log) (*RootChainBaseOperatorChanged, error) {
	event := new(RootChainBaseOperatorChanged)
	if err := _RootChainBase.contract.UnpackLog(event, "OperatorChanged", log); err != nil {
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
var RootChainStorageBin = "0x608060405234801561001057600080fd5b50610927806100206000396000f3fe608060405234801561001057600080fd5b506004361061021c5760003560e01c8063b17fa6e911610125578063ca6f6380116100ad578063e259faf71161007c578063e259faf714610585578063e7b88b801461058d578063ea7f22a814610595578063f4f31de4146105b2578063fb788a27146105cf5761021c565b8063ca6f63801461053a578063d691acd814610221578063da0185f814610557578063de0ce17d1461057d5761021c565b8063b8066bcb116100f4578063b8066bcb146104b1578063c0e86064146104b9578063c2bc88fa14610522578063c54626cc1461052a578063c8ad329f146105325761021c565b8063b17fa6e91461037c578063b2ae9ba814610221578063b443f3cc14610384578063b6715647146104945761021c565b8063570ca735116101a85780638155717d116101775780638155717d146103645780638b5172d0146102215780638eb288ca1461036c57806394be3aa514610221578063ab96da2d146103745761021c565b8063570ca735146103135780636fb7f5581461033757806372ecb9a81461033f5780637b929c271461035c5761021c565b8063192adc5b116101ef578063192adc5b1461022157806323691566146102535780632dc6bb7b1461025b578063420bb4b8146102635780634ba3a1261461027f5761021c565b8063033cfbed1461022157806308c4fff01461023b578063164bc2ae14610243578063183d2d1c1461024b575b600080fd5b6102296105d7565b60408051918252519081900360200190f35b6102296105dc565b6102296105e1565b6102296105e7565b6102296105ed565b6102296105f3565b61026b6105f9565b604080519115158252519081900360200190f35b61029c6004803603602081101561029557600080fd5b50356105fe565b6040805167ffffffffffffffff9d8e1681529b8d1660208d0152998c168b8b0152978b1660608b0152958a1660808a015293891660a089015291881660c0880152871660e0870152861661010086015285166101208501529093166101408301529115156101608201529051908190036101800190f35b61031b610676565b604080516001600160a01b039092168252519081900360200190f35b61031b61068a565b6102296004803603602081101561035557600080fd5b5035610699565b61026b6106ab565b6102296106b4565b6102296106b9565b6102296106c0565b6102296106c6565b6103a16004803603602081101561039a57600080fd5b50356106cb565b6040805167ffffffffffffffff8d1681528b15156020808301919091528b151592820192909252891515606082015288151560808201526001600160801b03881660a08201526001600160a01b0380881660c0830152861660e082015261010081018590526101208101849052610160610140820181815284519183019190915283519192909161018084019185019080838360005b8381101561044f578181015183820152602001610437565b50505050905090810190601f16801561047c5780820380516001836020036101000a031916815260200191505b509c5050505050505050505050505060405180910390f35b610229600480360360208110156104aa57600080fd5b50356107ed565b61031b6107ff565b6104d6600480360360208110156104cf57600080fd5b503561080e565b60408051961515875267ffffffffffffffff958616602088015293851686850152918416606086015290921660808401526001600160a01b0390911660a0830152519081900360c00190f35b610229610876565b61022961087b565b610229610881565b6102296004803603602081101561055057600080fd5b5035610887565b61031b6004803603602081101561056d57600080fd5b50356001600160a01b0316610899565b61031b6105d7565b61031b6108b4565b61031b6108c3565b6104d6600480360360208110156105ab57600080fd5b50356108d2565b6103a1600480360360208110156105c857600080fd5b50356108df565b6102296108ec565b600081565b600f81565b60105481565b60065481565b600d5481565b60135481565b600181565b60086020526000908152604090208054600182015460029092015467ffffffffffffffff80831693600160401b808504831694600160801b808204851695600160c01b92839004861695858116958581048216958482048316959182900483169484841694918204841693908204169160ff9104168c565b60005461010090046001600160a01b031681565b6004546001600160a01b031681565b60076020526000908152604090205481565b60005460ff1681565b600a81565b620186a081565b60055481565b601481565b600981815481106106d857fe5b6000918252602091829020600691909102018054600180830154600280850154600386015460048701546005880180546040805161010099831615999099026000190190911695909504601f81018b90048b0288018b0190955284875267ffffffffffffffff88169a50600160401b880460ff9081169a600160481b8a0482169a600160501b8b0483169a600160581b810490931699600160601b9093046001600160801b0316986001600160a01b03908116989716969093918301828280156107e35780601f106107b8576101008083540402835291602001916107e3565b820191906000526020600020905b8154815290600101906020018083116107c657829003601f168201915b505050505090508b565b600e6020526000908152604090205481565b6003546001600160a01b031681565b600c818154811061081b57fe5b60009182526020909120600290910201805460019091015460ff8216925067ffffffffffffffff6101008304811692600160481b8104821692600160881b9091048216918116906001600160a01b03600160401b9091041686565b603c81565b60145481565b60115481565b600f6020526000908152604090205481565b6015602052600090815260409020546001600160a01b031681565b6002546001600160a01b031681565b6001546001600160a01b031681565b600b818154811061081b57fe5b600a81815481106106d857fe5b6012548156fea265627a7a7231582050fb340a9c8ce204a949c7618e67ebead6f9d9e0b47c5a41d2ac573628ef8de264736f6c634300050c0032"

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

// SubmitterRoleABI is the input ABI used to generate the binding from.
const SubmitterRoleABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"SubmitterAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"SubmitterRemoved\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addSubmitter\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isSubmitter\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceSubmitter\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// SubmitterRoleFuncSigs maps the 4-byte function signature to its string representation.
var SubmitterRoleFuncSigs = map[string]string{
	"072900f9": "addSubmitter(address)",
	"a926fdbc": "isSubmitter(address)",
	"5e0ca71b": "renounceSubmitter()",
}

// SubmitterRole is an auto generated Go binding around an Ethereum contract.
type SubmitterRole struct {
	SubmitterRoleCaller     // Read-only binding to the contract
	SubmitterRoleTransactor // Write-only binding to the contract
	SubmitterRoleFilterer   // Log filterer for contract events
}

// SubmitterRoleCaller is an auto generated read-only Go binding around an Ethereum contract.
type SubmitterRoleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SubmitterRoleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SubmitterRoleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SubmitterRoleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SubmitterRoleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SubmitterRoleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SubmitterRoleSession struct {
	Contract     *SubmitterRole    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SubmitterRoleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SubmitterRoleCallerSession struct {
	Contract *SubmitterRoleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// SubmitterRoleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SubmitterRoleTransactorSession struct {
	Contract     *SubmitterRoleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// SubmitterRoleRaw is an auto generated low-level Go binding around an Ethereum contract.
type SubmitterRoleRaw struct {
	Contract *SubmitterRole // Generic contract binding to access the raw methods on
}

// SubmitterRoleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SubmitterRoleCallerRaw struct {
	Contract *SubmitterRoleCaller // Generic read-only contract binding to access the raw methods on
}

// SubmitterRoleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SubmitterRoleTransactorRaw struct {
	Contract *SubmitterRoleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSubmitterRole creates a new instance of SubmitterRole, bound to a specific deployed contract.
func NewSubmitterRole(address common.Address, backend bind.ContractBackend) (*SubmitterRole, error) {
	contract, err := bindSubmitterRole(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SubmitterRole{SubmitterRoleCaller: SubmitterRoleCaller{contract: contract}, SubmitterRoleTransactor: SubmitterRoleTransactor{contract: contract}, SubmitterRoleFilterer: SubmitterRoleFilterer{contract: contract}}, nil
}

// NewSubmitterRoleCaller creates a new read-only instance of SubmitterRole, bound to a specific deployed contract.
func NewSubmitterRoleCaller(address common.Address, caller bind.ContractCaller) (*SubmitterRoleCaller, error) {
	contract, err := bindSubmitterRole(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SubmitterRoleCaller{contract: contract}, nil
}

// NewSubmitterRoleTransactor creates a new write-only instance of SubmitterRole, bound to a specific deployed contract.
func NewSubmitterRoleTransactor(address common.Address, transactor bind.ContractTransactor) (*SubmitterRoleTransactor, error) {
	contract, err := bindSubmitterRole(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SubmitterRoleTransactor{contract: contract}, nil
}

// NewSubmitterRoleFilterer creates a new log filterer instance of SubmitterRole, bound to a specific deployed contract.
func NewSubmitterRoleFilterer(address common.Address, filterer bind.ContractFilterer) (*SubmitterRoleFilterer, error) {
	contract, err := bindSubmitterRole(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SubmitterRoleFilterer{contract: contract}, nil
}

// bindSubmitterRole binds a generic wrapper to an already deployed contract.
func bindSubmitterRole(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SubmitterRoleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SubmitterRole *SubmitterRoleRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SubmitterRole.Contract.SubmitterRoleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SubmitterRole *SubmitterRoleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SubmitterRole.Contract.SubmitterRoleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SubmitterRole *SubmitterRoleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SubmitterRole.Contract.SubmitterRoleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SubmitterRole *SubmitterRoleCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SubmitterRole.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SubmitterRole *SubmitterRoleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SubmitterRole.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SubmitterRole *SubmitterRoleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SubmitterRole.Contract.contract.Transact(opts, method, params...)
}

// IsSubmitter is a free data retrieval call binding the contract method 0xa926fdbc.
//
// Solidity: function isSubmitter(address account) constant returns(bool)
func (_SubmitterRole *SubmitterRoleCaller) IsSubmitter(opts *bind.CallOpts, account common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _SubmitterRole.contract.Call(opts, out, "isSubmitter", account)
	return *ret0, err
}

// IsSubmitter is a free data retrieval call binding the contract method 0xa926fdbc.
//
// Solidity: function isSubmitter(address account) constant returns(bool)
func (_SubmitterRole *SubmitterRoleSession) IsSubmitter(account common.Address) (bool, error) {
	return _SubmitterRole.Contract.IsSubmitter(&_SubmitterRole.CallOpts, account)
}

// IsSubmitter is a free data retrieval call binding the contract method 0xa926fdbc.
//
// Solidity: function isSubmitter(address account) constant returns(bool)
func (_SubmitterRole *SubmitterRoleCallerSession) IsSubmitter(account common.Address) (bool, error) {
	return _SubmitterRole.Contract.IsSubmitter(&_SubmitterRole.CallOpts, account)
}

// AddSubmitter is a paid mutator transaction binding the contract method 0x072900f9.
//
// Solidity: function addSubmitter(address account) returns()
func (_SubmitterRole *SubmitterRoleTransactor) AddSubmitter(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _SubmitterRole.contract.Transact(opts, "addSubmitter", account)
}

// AddSubmitter is a paid mutator transaction binding the contract method 0x072900f9.
//
// Solidity: function addSubmitter(address account) returns()
func (_SubmitterRole *SubmitterRoleSession) AddSubmitter(account common.Address) (*types.Transaction, error) {
	return _SubmitterRole.Contract.AddSubmitter(&_SubmitterRole.TransactOpts, account)
}

// AddSubmitter is a paid mutator transaction binding the contract method 0x072900f9.
//
// Solidity: function addSubmitter(address account) returns()
func (_SubmitterRole *SubmitterRoleTransactorSession) AddSubmitter(account common.Address) (*types.Transaction, error) {
	return _SubmitterRole.Contract.AddSubmitter(&_SubmitterRole.TransactOpts, account)
}

// RenounceSubmitter is a paid mutator transaction binding the contract method 0x5e0ca71b.
//
// Solidity: function renounceSubmitter() returns()
func (_SubmitterRole *SubmitterRoleTransactor) RenounceSubmitter(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SubmitterRole.contract.Transact(opts, "renounceSubmitter")
}

// RenounceSubmitter is a paid mutator transaction binding the contract method 0x5e0ca71b.
//
// Solidity: function renounceSubmitter() returns()
func (_SubmitterRole *SubmitterRoleSession) RenounceSubmitter() (*types.Transaction, error) {
	return _SubmitterRole.Contract.RenounceSubmitter(&_SubmitterRole.TransactOpts)
}

// RenounceSubmitter is a paid mutator transaction binding the contract method 0x5e0ca71b.
//
// Solidity: function renounceSubmitter() returns()
func (_SubmitterRole *SubmitterRoleTransactorSession) RenounceSubmitter() (*types.Transaction, error) {
	return _SubmitterRole.Contract.RenounceSubmitter(&_SubmitterRole.TransactOpts)
}

// SubmitterRoleSubmitterAddedIterator is returned from FilterSubmitterAdded and is used to iterate over the raw logs and unpacked data for SubmitterAdded events raised by the SubmitterRole contract.
type SubmitterRoleSubmitterAddedIterator struct {
	Event *SubmitterRoleSubmitterAdded // Event containing the contract specifics and raw log

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
func (it *SubmitterRoleSubmitterAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubmitterRoleSubmitterAdded)
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
		it.Event = new(SubmitterRoleSubmitterAdded)
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
func (it *SubmitterRoleSubmitterAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubmitterRoleSubmitterAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubmitterRoleSubmitterAdded represents a SubmitterAdded event raised by the SubmitterRole contract.
type SubmitterRoleSubmitterAdded struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSubmitterAdded is a free log retrieval operation binding the contract event 0xb079bc2cbde1f186e0b351d4a87c4597e3ed098f571548617449e73506428d8b.
//
// Solidity: event SubmitterAdded(address indexed account)
func (_SubmitterRole *SubmitterRoleFilterer) FilterSubmitterAdded(opts *bind.FilterOpts, account []common.Address) (*SubmitterRoleSubmitterAddedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _SubmitterRole.contract.FilterLogs(opts, "SubmitterAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return &SubmitterRoleSubmitterAddedIterator{contract: _SubmitterRole.contract, event: "SubmitterAdded", logs: logs, sub: sub}, nil
}

// WatchSubmitterAdded is a free log subscription operation binding the contract event 0xb079bc2cbde1f186e0b351d4a87c4597e3ed098f571548617449e73506428d8b.
//
// Solidity: event SubmitterAdded(address indexed account)
func (_SubmitterRole *SubmitterRoleFilterer) WatchSubmitterAdded(opts *bind.WatchOpts, sink chan<- *SubmitterRoleSubmitterAdded, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _SubmitterRole.contract.WatchLogs(opts, "SubmitterAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubmitterRoleSubmitterAdded)
				if err := _SubmitterRole.contract.UnpackLog(event, "SubmitterAdded", log); err != nil {
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

// ParseSubmitterAdded is a log parse operation binding the contract event 0xb079bc2cbde1f186e0b351d4a87c4597e3ed098f571548617449e73506428d8b.
//
// Solidity: event SubmitterAdded(address indexed account)
func (_SubmitterRole *SubmitterRoleFilterer) ParseSubmitterAdded(log types.Log) (*SubmitterRoleSubmitterAdded, error) {
	event := new(SubmitterRoleSubmitterAdded)
	if err := _SubmitterRole.contract.UnpackLog(event, "SubmitterAdded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SubmitterRoleSubmitterRemovedIterator is returned from FilterSubmitterRemoved and is used to iterate over the raw logs and unpacked data for SubmitterRemoved events raised by the SubmitterRole contract.
type SubmitterRoleSubmitterRemovedIterator struct {
	Event *SubmitterRoleSubmitterRemoved // Event containing the contract specifics and raw log

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
func (it *SubmitterRoleSubmitterRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubmitterRoleSubmitterRemoved)
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
		it.Event = new(SubmitterRoleSubmitterRemoved)
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
func (it *SubmitterRoleSubmitterRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubmitterRoleSubmitterRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubmitterRoleSubmitterRemoved represents a SubmitterRemoved event raised by the SubmitterRole contract.
type SubmitterRoleSubmitterRemoved struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSubmitterRemoved is a free log retrieval operation binding the contract event 0xf84a004e1673d2f349a7c93c72b3794b8eba6d2f9338044d8c8cd260e51a57a1.
//
// Solidity: event SubmitterRemoved(address indexed account)
func (_SubmitterRole *SubmitterRoleFilterer) FilterSubmitterRemoved(opts *bind.FilterOpts, account []common.Address) (*SubmitterRoleSubmitterRemovedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _SubmitterRole.contract.FilterLogs(opts, "SubmitterRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return &SubmitterRoleSubmitterRemovedIterator{contract: _SubmitterRole.contract, event: "SubmitterRemoved", logs: logs, sub: sub}, nil
}

// WatchSubmitterRemoved is a free log subscription operation binding the contract event 0xf84a004e1673d2f349a7c93c72b3794b8eba6d2f9338044d8c8cd260e51a57a1.
//
// Solidity: event SubmitterRemoved(address indexed account)
func (_SubmitterRole *SubmitterRoleFilterer) WatchSubmitterRemoved(opts *bind.WatchOpts, sink chan<- *SubmitterRoleSubmitterRemoved, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _SubmitterRole.contract.WatchLogs(opts, "SubmitterRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubmitterRoleSubmitterRemoved)
				if err := _SubmitterRole.contract.UnpackLog(event, "SubmitterRemoved", log); err != nil {
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

// ParseSubmitterRemoved is a log parse operation binding the contract event 0xf84a004e1673d2f349a7c93c72b3794b8eba6d2f9338044d8c8cd260e51a57a1.
//
// Solidity: event SubmitterRemoved(address indexed account)
func (_SubmitterRole *SubmitterRoleFilterer) ParseSubmitterRemoved(log types.Log) (*SubmitterRoleSubmitterRemoved, error) {
	event := new(SubmitterRoleSubmitterRemoved)
	if err := _SubmitterRole.contract.UnpackLog(event, "SubmitterRemoved", log); err != nil {
		return nil, err
	}
	return event, nil
}
