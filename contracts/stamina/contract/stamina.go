// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// StaminaABI is the input ABI used to generate the binding from.
const StaminaABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"initialized\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getStamina\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"delegatee\",\"type\":\"address\"}],\"name\":\"resetStamina\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"delegatee\",\"type\":\"address\"}],\"name\":\"getTotalDeposit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"t\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"delegater\",\"type\":\"address\"}],\"name\":\"getDelegatee\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"minDeposit\",\"type\":\"uint256\"}],\"name\":\"init\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"delegatee\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"addStamina\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"depositor\",\"type\":\"address\"},{\"name\":\"delegatee\",\"type\":\"address\"}],\"name\":\"getDeposit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"delegatee\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"subtractStamina\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MIN_DEPOSIT\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"delegater\",\"type\":\"address\"}],\"name\":\"setDelegatee\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"delegatee\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"delegatee\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"depositor\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"delegatee\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"depositor\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"delegatee\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"delegater\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"oldDelegatee\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"newDelegatee\",\"type\":\"address\"}],\"name\":\"DelegateeChanged\",\"type\":\"event\"}]"

// StaminaBin is the compiled bytecode used for deploying new contracts.
const StaminaBin = `0x608060405261dead60045534801561001657600080fd5b506107c3806100266000396000f3006080604052600436106100cf5763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663158ef93e81146100d45780633900e4ec146100fd57806362eb5d9814610142578063857184d11461017757806392d0d153146101aa5780639b4e735f146101bf578063b7b0422d1461020e578063bcac973614610238578063c35082a914610271578063d1c0c042146102ac578063e1e158a5146102e5578063e842a64b146102fa578063f340fa011461032d578063f3fef3a314610353575b600080fd5b3480156100e057600080fd5b506100e961038c565b604080519115158252519081900360200190f35b34801561010957600080fd5b506101306004803603602081101561012057600080fd5b5035600160a060020a0316610395565b60408051918252519081900360200190f35b34801561014e57600080fd5b506101756004803603602081101561016557600080fd5b5035600160a060020a03166103b0565b005b34801561018357600080fd5b506101306004803603602081101561019a57600080fd5b5035600160a060020a03166103d6565b3480156101b657600080fd5b506101306103f1565b3480156101cb57600080fd5b506101f2600480360360208110156101e257600080fd5b5035600160a060020a03166103f7565b60408051600160a060020a039092168252519081900360200190f35b34801561021a57600080fd5b506101756004803603602081101561023157600080fd5b5035610415565b34801561024457600080fd5b506100e96004803603604081101561025b57600080fd5b50600160a060020a038135169060200135610437565b34801561027d57600080fd5b506101306004803603604081101561029457600080fd5b50600160a060020a03813581169160200135166104bc565b3480156102b857600080fd5b506100e9600480360360408110156102cf57600080fd5b50600160a060020a0381351690602001356104e7565b3480156102f157600080fd5b50610130610536565b34801561030657600080fd5b506100e96004803603602081101561031d57600080fd5b5035600160a060020a031661053c565b6100e96004803603602081101561034357600080fd5b5035600160a060020a03166105c5565b34801561035f57600080fd5b506100e96004803603604081101561037657600080fd5b50600160a060020a03813516906020013561069a565b60055460ff1681565b600160a060020a031660009081526001602052604090205490565b600160a060020a0316600090815260026020908152604080832054600190925290912055565b600160a060020a031660009081526002602052604090205490565b60045481565b600160a060020a039081166000908152602081905260409020541690565b60055460ff161561042557600080fd5b6006556005805460ff19166001179055565b600160a060020a0382166000908152600260209081526040808320546001909252822054838101811061046957600080fd5b8084018281111561049457600160a060020a03861660009081526001602052604090208390556104b0565b600160a060020a03861660009081526001602052604090208190555b50600195945050505050565b600160a060020a03918216600090815260036020908152604080832093909416825291909152205490565b600160a060020a038216600090815260016020526040812054828103811161050e57600080fd5b600160a060020a03939093166000908152600160208190526040909120929093039091555090565b60065481565b600160a060020a03808216600081815260208181526040808320805473ffffffffffffffffffffffffffffffffffffffff1981163390811790925582519586529095169184018290528381019490945292519092917f5884d7e3ec123de8e772bcf576c18dcdad75b056c4314f999ed966693419c692919081900360600190a150600192915050565b6006546000903410156105d757600080fd5b600160a060020a0382166000818152600260209081526040808320543384526003835281842094845293909152902054348201821061061557600080fd5b348101811061062357600080fd5b600160a060020a03841660008181526002602090815260408083203487810190915533808552600384528285208686528452938290208682019055815190815290517f8752a472e571a816aea92eec8dae9baf628e840f4929fbcc2d155e6233ff68a7929181900390910190a35060019392505050565b600160a060020a038216600081815260026020908152604080832054338452600383528184209484529390915281205490919083820382116106db57600080fd5b83810381116106e957600080fd5b600160a060020a0385166000818152600260209081526040808320888703905533808452600383528184209484529390915280822087850390555186156108fc0291879190818181858888f1935050505015801561074b573d6000803e3d6000fd5b50604080518581529051600160a060020a0387169133917fd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb9181900360200190a35060019493505050505600a165627a7a72305820a011f3b07b2d2070faaf83531209f631ad3e33cd6aa2a5012e67d91316a8ef4d0029`

// DeployStamina deploys a new Ethereum contract, binding an instance of Stamina to it.
func DeployStamina(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Stamina, error) {
	parsed, err := abi.JSON(strings.NewReader(StaminaABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(StaminaBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Stamina{StaminaCaller: StaminaCaller{contract: contract}, StaminaTransactor: StaminaTransactor{contract: contract}, StaminaFilterer: StaminaFilterer{contract: contract}}, nil
}

// Stamina is an auto generated Go binding around an Ethereum contract.
type Stamina struct {
	StaminaCaller     // Read-only binding to the contract
	StaminaTransactor // Write-only binding to the contract
	StaminaFilterer   // Log filterer for contract events
}

// StaminaCaller is an auto generated read-only Go binding around an Ethereum contract.
type StaminaCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StaminaTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StaminaTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StaminaFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StaminaFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StaminaSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StaminaSession struct {
	Contract     *Stamina          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StaminaCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StaminaCallerSession struct {
	Contract *StaminaCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// StaminaTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StaminaTransactorSession struct {
	Contract     *StaminaTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// StaminaRaw is an auto generated low-level Go binding around an Ethereum contract.
type StaminaRaw struct {
	Contract *Stamina // Generic contract binding to access the raw methods on
}

// StaminaCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StaminaCallerRaw struct {
	Contract *StaminaCaller // Generic read-only contract binding to access the raw methods on
}

// StaminaTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StaminaTransactorRaw struct {
	Contract *StaminaTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStamina creates a new instance of Stamina, bound to a specific deployed contract.
func NewStamina(address common.Address, backend bind.ContractBackend) (*Stamina, error) {
	contract, err := bindStamina(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Stamina{StaminaCaller: StaminaCaller{contract: contract}, StaminaTransactor: StaminaTransactor{contract: contract}, StaminaFilterer: StaminaFilterer{contract: contract}}, nil
}

// NewStaminaCaller creates a new read-only instance of Stamina, bound to a specific deployed contract.
func NewStaminaCaller(address common.Address, caller bind.ContractCaller) (*StaminaCaller, error) {
	contract, err := bindStamina(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StaminaCaller{contract: contract}, nil
}

// NewStaminaTransactor creates a new write-only instance of Stamina, bound to a specific deployed contract.
func NewStaminaTransactor(address common.Address, transactor bind.ContractTransactor) (*StaminaTransactor, error) {
	contract, err := bindStamina(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StaminaTransactor{contract: contract}, nil
}

// NewStaminaFilterer creates a new log filterer instance of Stamina, bound to a specific deployed contract.
func NewStaminaFilterer(address common.Address, filterer bind.ContractFilterer) (*StaminaFilterer, error) {
	contract, err := bindStamina(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StaminaFilterer{contract: contract}, nil
}

// bindStamina binds a generic wrapper to an already deployed contract.
func bindStamina(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StaminaABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Stamina *StaminaRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Stamina.Contract.StaminaCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Stamina *StaminaRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stamina.Contract.StaminaTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Stamina *StaminaRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Stamina.Contract.StaminaTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Stamina *StaminaCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Stamina.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Stamina *StaminaTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stamina.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Stamina *StaminaTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Stamina.Contract.contract.Transact(opts, method, params...)
}

// MINDEPOSIT is a free data retrieval call binding the contract method 0xe1e158a5.
//
// Solidity: function MIN_DEPOSIT() constant returns(uint256)
func (_Stamina *StaminaCaller) MINDEPOSIT(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Stamina.contract.Call(opts, out, "MIN_DEPOSIT")
	return *ret0, err
}

// MINDEPOSIT is a free data retrieval call binding the contract method 0xe1e158a5.
//
// Solidity: function MIN_DEPOSIT() constant returns(uint256)
func (_Stamina *StaminaSession) MINDEPOSIT() (*big.Int, error) {
	return _Stamina.Contract.MINDEPOSIT(&_Stamina.CallOpts)
}

// MINDEPOSIT is a free data retrieval call binding the contract method 0xe1e158a5.
//
// Solidity: function MIN_DEPOSIT() constant returns(uint256)
func (_Stamina *StaminaCallerSession) MINDEPOSIT() (*big.Int, error) {
	return _Stamina.Contract.MINDEPOSIT(&_Stamina.CallOpts)
}

// GetDelegatee is a free data retrieval call binding the contract method 0x9b4e735f.
//
// Solidity: function getDelegatee(delegater address) constant returns(address)
func (_Stamina *StaminaCaller) GetDelegatee(opts *bind.CallOpts, delegater common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Stamina.contract.Call(opts, out, "getDelegatee", delegater)
	return *ret0, err
}

// GetDelegatee is a free data retrieval call binding the contract method 0x9b4e735f.
//
// Solidity: function getDelegatee(delegater address) constant returns(address)
func (_Stamina *StaminaSession) GetDelegatee(delegater common.Address) (common.Address, error) {
	return _Stamina.Contract.GetDelegatee(&_Stamina.CallOpts, delegater)
}

// GetDelegatee is a free data retrieval call binding the contract method 0x9b4e735f.
//
// Solidity: function getDelegatee(delegater address) constant returns(address)
func (_Stamina *StaminaCallerSession) GetDelegatee(delegater common.Address) (common.Address, error) {
	return _Stamina.Contract.GetDelegatee(&_Stamina.CallOpts, delegater)
}

// GetDeposit is a free data retrieval call binding the contract method 0xc35082a9.
//
// Solidity: function getDeposit(depositor address, delegatee address) constant returns(uint256)
func (_Stamina *StaminaCaller) GetDeposit(opts *bind.CallOpts, depositor common.Address, delegatee common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Stamina.contract.Call(opts, out, "getDeposit", depositor, delegatee)
	return *ret0, err
}

// GetDeposit is a free data retrieval call binding the contract method 0xc35082a9.
//
// Solidity: function getDeposit(depositor address, delegatee address) constant returns(uint256)
func (_Stamina *StaminaSession) GetDeposit(depositor common.Address, delegatee common.Address) (*big.Int, error) {
	return _Stamina.Contract.GetDeposit(&_Stamina.CallOpts, depositor, delegatee)
}

// GetDeposit is a free data retrieval call binding the contract method 0xc35082a9.
//
// Solidity: function getDeposit(depositor address, delegatee address) constant returns(uint256)
func (_Stamina *StaminaCallerSession) GetDeposit(depositor common.Address, delegatee common.Address) (*big.Int, error) {
	return _Stamina.Contract.GetDeposit(&_Stamina.CallOpts, depositor, delegatee)
}

// GetStamina is a free data retrieval call binding the contract method 0x3900e4ec.
//
// Solidity: function getStamina(addr address) constant returns(uint256)
func (_Stamina *StaminaCaller) GetStamina(opts *bind.CallOpts, addr common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Stamina.contract.Call(opts, out, "getStamina", addr)
	return *ret0, err
}

// GetStamina is a free data retrieval call binding the contract method 0x3900e4ec.
//
// Solidity: function getStamina(addr address) constant returns(uint256)
func (_Stamina *StaminaSession) GetStamina(addr common.Address) (*big.Int, error) {
	return _Stamina.Contract.GetStamina(&_Stamina.CallOpts, addr)
}

// GetStamina is a free data retrieval call binding the contract method 0x3900e4ec.
//
// Solidity: function getStamina(addr address) constant returns(uint256)
func (_Stamina *StaminaCallerSession) GetStamina(addr common.Address) (*big.Int, error) {
	return _Stamina.Contract.GetStamina(&_Stamina.CallOpts, addr)
}

// GetTotalDeposit is a free data retrieval call binding the contract method 0x857184d1.
//
// Solidity: function getTotalDeposit(delegatee address) constant returns(uint256)
func (_Stamina *StaminaCaller) GetTotalDeposit(opts *bind.CallOpts, delegatee common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Stamina.contract.Call(opts, out, "getTotalDeposit", delegatee)
	return *ret0, err
}

// GetTotalDeposit is a free data retrieval call binding the contract method 0x857184d1.
//
// Solidity: function getTotalDeposit(delegatee address) constant returns(uint256)
func (_Stamina *StaminaSession) GetTotalDeposit(delegatee common.Address) (*big.Int, error) {
	return _Stamina.Contract.GetTotalDeposit(&_Stamina.CallOpts, delegatee)
}

// GetTotalDeposit is a free data retrieval call binding the contract method 0x857184d1.
//
// Solidity: function getTotalDeposit(delegatee address) constant returns(uint256)
func (_Stamina *StaminaCallerSession) GetTotalDeposit(delegatee common.Address) (*big.Int, error) {
	return _Stamina.Contract.GetTotalDeposit(&_Stamina.CallOpts, delegatee)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() constant returns(bool)
func (_Stamina *StaminaCaller) Initialized(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Stamina.contract.Call(opts, out, "initialized")
	return *ret0, err
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() constant returns(bool)
func (_Stamina *StaminaSession) Initialized() (bool, error) {
	return _Stamina.Contract.Initialized(&_Stamina.CallOpts)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() constant returns(bool)
func (_Stamina *StaminaCallerSession) Initialized() (bool, error) {
	return _Stamina.Contract.Initialized(&_Stamina.CallOpts)
}

// T is a free data retrieval call binding the contract method 0x92d0d153.
//
// Solidity: function t() constant returns(uint256)
func (_Stamina *StaminaCaller) T(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Stamina.contract.Call(opts, out, "t")
	return *ret0, err
}

// T is a free data retrieval call binding the contract method 0x92d0d153.
//
// Solidity: function t() constant returns(uint256)
func (_Stamina *StaminaSession) T() (*big.Int, error) {
	return _Stamina.Contract.T(&_Stamina.CallOpts)
}

// T is a free data retrieval call binding the contract method 0x92d0d153.
//
// Solidity: function t() constant returns(uint256)
func (_Stamina *StaminaCallerSession) T() (*big.Int, error) {
	return _Stamina.Contract.T(&_Stamina.CallOpts)
}

// AddStamina is a paid mutator transaction binding the contract method 0xbcac9736.
//
// Solidity: function addStamina(delegatee address, amount uint256) returns(bool)
func (_Stamina *StaminaTransactor) AddStamina(opts *bind.TransactOpts, delegatee common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Stamina.contract.Transact(opts, "addStamina", delegatee, amount)
}

// AddStamina is a paid mutator transaction binding the contract method 0xbcac9736.
//
// Solidity: function addStamina(delegatee address, amount uint256) returns(bool)
func (_Stamina *StaminaSession) AddStamina(delegatee common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Stamina.Contract.AddStamina(&_Stamina.TransactOpts, delegatee, amount)
}

// AddStamina is a paid mutator transaction binding the contract method 0xbcac9736.
//
// Solidity: function addStamina(delegatee address, amount uint256) returns(bool)
func (_Stamina *StaminaTransactorSession) AddStamina(delegatee common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Stamina.Contract.AddStamina(&_Stamina.TransactOpts, delegatee, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xf340fa01.
//
// Solidity: function deposit(delegatee address) returns(bool)
func (_Stamina *StaminaTransactor) Deposit(opts *bind.TransactOpts, delegatee common.Address) (*types.Transaction, error) {
	return _Stamina.contract.Transact(opts, "deposit", delegatee)
}

// Deposit is a paid mutator transaction binding the contract method 0xf340fa01.
//
// Solidity: function deposit(delegatee address) returns(bool)
func (_Stamina *StaminaSession) Deposit(delegatee common.Address) (*types.Transaction, error) {
	return _Stamina.Contract.Deposit(&_Stamina.TransactOpts, delegatee)
}

// Deposit is a paid mutator transaction binding the contract method 0xf340fa01.
//
// Solidity: function deposit(delegatee address) returns(bool)
func (_Stamina *StaminaTransactorSession) Deposit(delegatee common.Address) (*types.Transaction, error) {
	return _Stamina.Contract.Deposit(&_Stamina.TransactOpts, delegatee)
}

// Init is a paid mutator transaction binding the contract method 0xb7b0422d.
//
// Solidity: function init(minDeposit uint256) returns()
func (_Stamina *StaminaTransactor) Init(opts *bind.TransactOpts, minDeposit *big.Int) (*types.Transaction, error) {
	return _Stamina.contract.Transact(opts, "init", minDeposit)
}

// Init is a paid mutator transaction binding the contract method 0xb7b0422d.
//
// Solidity: function init(minDeposit uint256) returns()
func (_Stamina *StaminaSession) Init(minDeposit *big.Int) (*types.Transaction, error) {
	return _Stamina.Contract.Init(&_Stamina.TransactOpts, minDeposit)
}

// Init is a paid mutator transaction binding the contract method 0xb7b0422d.
//
// Solidity: function init(minDeposit uint256) returns()
func (_Stamina *StaminaTransactorSession) Init(minDeposit *big.Int) (*types.Transaction, error) {
	return _Stamina.Contract.Init(&_Stamina.TransactOpts, minDeposit)
}

// ResetStamina is a paid mutator transaction binding the contract method 0x62eb5d98.
//
// Solidity: function resetStamina(delegatee address) returns()
func (_Stamina *StaminaTransactor) ResetStamina(opts *bind.TransactOpts, delegatee common.Address) (*types.Transaction, error) {
	return _Stamina.contract.Transact(opts, "resetStamina", delegatee)
}

// ResetStamina is a paid mutator transaction binding the contract method 0x62eb5d98.
//
// Solidity: function resetStamina(delegatee address) returns()
func (_Stamina *StaminaSession) ResetStamina(delegatee common.Address) (*types.Transaction, error) {
	return _Stamina.Contract.ResetStamina(&_Stamina.TransactOpts, delegatee)
}

// ResetStamina is a paid mutator transaction binding the contract method 0x62eb5d98.
//
// Solidity: function resetStamina(delegatee address) returns()
func (_Stamina *StaminaTransactorSession) ResetStamina(delegatee common.Address) (*types.Transaction, error) {
	return _Stamina.Contract.ResetStamina(&_Stamina.TransactOpts, delegatee)
}

// SetDelegatee is a paid mutator transaction binding the contract method 0xe842a64b.
//
// Solidity: function setDelegatee(delegater address) returns(bool)
func (_Stamina *StaminaTransactor) SetDelegatee(opts *bind.TransactOpts, delegater common.Address) (*types.Transaction, error) {
	return _Stamina.contract.Transact(opts, "setDelegatee", delegater)
}

// SetDelegatee is a paid mutator transaction binding the contract method 0xe842a64b.
//
// Solidity: function setDelegatee(delegater address) returns(bool)
func (_Stamina *StaminaSession) SetDelegatee(delegater common.Address) (*types.Transaction, error) {
	return _Stamina.Contract.SetDelegatee(&_Stamina.TransactOpts, delegater)
}

// SetDelegatee is a paid mutator transaction binding the contract method 0xe842a64b.
//
// Solidity: function setDelegatee(delegater address) returns(bool)
func (_Stamina *StaminaTransactorSession) SetDelegatee(delegater common.Address) (*types.Transaction, error) {
	return _Stamina.Contract.SetDelegatee(&_Stamina.TransactOpts, delegater)
}

// SubtractStamina is a paid mutator transaction binding the contract method 0xd1c0c042.
//
// Solidity: function subtractStamina(delegatee address, amount uint256) returns(bool)
func (_Stamina *StaminaTransactor) SubtractStamina(opts *bind.TransactOpts, delegatee common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Stamina.contract.Transact(opts, "subtractStamina", delegatee, amount)
}

// SubtractStamina is a paid mutator transaction binding the contract method 0xd1c0c042.
//
// Solidity: function subtractStamina(delegatee address, amount uint256) returns(bool)
func (_Stamina *StaminaSession) SubtractStamina(delegatee common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Stamina.Contract.SubtractStamina(&_Stamina.TransactOpts, delegatee, amount)
}

// SubtractStamina is a paid mutator transaction binding the contract method 0xd1c0c042.
//
// Solidity: function subtractStamina(delegatee address, amount uint256) returns(bool)
func (_Stamina *StaminaTransactorSession) SubtractStamina(delegatee common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Stamina.Contract.SubtractStamina(&_Stamina.TransactOpts, delegatee, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(delegatee address, amount uint256) returns(bool)
func (_Stamina *StaminaTransactor) Withdraw(opts *bind.TransactOpts, delegatee common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Stamina.contract.Transact(opts, "withdraw", delegatee, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(delegatee address, amount uint256) returns(bool)
func (_Stamina *StaminaSession) Withdraw(delegatee common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Stamina.Contract.Withdraw(&_Stamina.TransactOpts, delegatee, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(delegatee address, amount uint256) returns(bool)
func (_Stamina *StaminaTransactorSession) Withdraw(delegatee common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Stamina.Contract.Withdraw(&_Stamina.TransactOpts, delegatee, amount)
}

// StaminaDelegateeChangedIterator is returned from FilterDelegateeChanged and is used to iterate over the raw logs and unpacked data for DelegateeChanged events raised by the Stamina contract.
type StaminaDelegateeChangedIterator struct {
	Event *StaminaDelegateeChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StaminaDelegateeChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StaminaDelegateeChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StaminaDelegateeChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StaminaDelegateeChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StaminaDelegateeChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StaminaDelegateeChanged represents a DelegateeChanged event raised by the Stamina contract.
type StaminaDelegateeChanged struct {
	Delegater    common.Address
	OldDelegatee common.Address
	NewDelegatee common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDelegateeChanged is a free log retrieval operation binding the contract event 0x5884d7e3ec123de8e772bcf576c18dcdad75b056c4314f999ed966693419c692.
//
// Solidity: e DelegateeChanged(delegater address, oldDelegatee address, newDelegatee address)
func (_Stamina *StaminaFilterer) FilterDelegateeChanged(opts *bind.FilterOpts) (*StaminaDelegateeChangedIterator, error) {

	logs, sub, err := _Stamina.contract.FilterLogs(opts, "DelegateeChanged")
	if err != nil {
		return nil, err
	}
	return &StaminaDelegateeChangedIterator{contract: _Stamina.contract, event: "DelegateeChanged", logs: logs, sub: sub}, nil
}

// WatchDelegateeChanged is a free log subscription operation binding the contract event 0x5884d7e3ec123de8e772bcf576c18dcdad75b056c4314f999ed966693419c692.
//
// Solidity: e DelegateeChanged(delegater address, oldDelegatee address, newDelegatee address)
func (_Stamina *StaminaFilterer) WatchDelegateeChanged(opts *bind.WatchOpts, sink chan<- *StaminaDelegateeChanged) (event.Subscription, error) {

	logs, sub, err := _Stamina.contract.WatchLogs(opts, "DelegateeChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StaminaDelegateeChanged)
				if err := _Stamina.contract.UnpackLog(event, "DelegateeChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// StaminaDepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the Stamina contract.
type StaminaDepositedIterator struct {
	Event *StaminaDeposited // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StaminaDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StaminaDeposited)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StaminaDeposited)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StaminaDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StaminaDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StaminaDeposited represents a Deposited event raised by the Stamina contract.
type StaminaDeposited struct {
	Depositor common.Address
	Delegatee common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDeposited is a free log retrieval operation binding the contract event 0x8752a472e571a816aea92eec8dae9baf628e840f4929fbcc2d155e6233ff68a7.
//
// Solidity: e Deposited(depositor indexed address, delegatee indexed address, amount uint256)
func (_Stamina *StaminaFilterer) FilterDeposited(opts *bind.FilterOpts, depositor []common.Address, delegatee []common.Address) (*StaminaDepositedIterator, error) {

	var depositorRule []interface{}
	for _, depositorItem := range depositor {
		depositorRule = append(depositorRule, depositorItem)
	}
	var delegateeRule []interface{}
	for _, delegateeItem := range delegatee {
		delegateeRule = append(delegateeRule, delegateeItem)
	}

	logs, sub, err := _Stamina.contract.FilterLogs(opts, "Deposited", depositorRule, delegateeRule)
	if err != nil {
		return nil, err
	}
	return &StaminaDepositedIterator{contract: _Stamina.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0x8752a472e571a816aea92eec8dae9baf628e840f4929fbcc2d155e6233ff68a7.
//
// Solidity: e Deposited(depositor indexed address, delegatee indexed address, amount uint256)
func (_Stamina *StaminaFilterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *StaminaDeposited, depositor []common.Address, delegatee []common.Address) (event.Subscription, error) {

	var depositorRule []interface{}
	for _, depositorItem := range depositor {
		depositorRule = append(depositorRule, depositorItem)
	}
	var delegateeRule []interface{}
	for _, delegateeItem := range delegatee {
		delegateeRule = append(delegateeRule, delegateeItem)
	}

	logs, sub, err := _Stamina.contract.WatchLogs(opts, "Deposited", depositorRule, delegateeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StaminaDeposited)
				if err := _Stamina.contract.UnpackLog(event, "Deposited", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// StaminaWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the Stamina contract.
type StaminaWithdrawnIterator struct {
	Event *StaminaWithdrawn // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StaminaWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StaminaWithdrawn)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StaminaWithdrawn)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StaminaWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StaminaWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StaminaWithdrawn represents a Withdrawn event raised by the Stamina contract.
type StaminaWithdrawn struct {
	Depositor common.Address
	Delegatee common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0xd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb.
//
// Solidity: e Withdrawn(depositor indexed address, delegatee indexed address, amount uint256)
func (_Stamina *StaminaFilterer) FilterWithdrawn(opts *bind.FilterOpts, depositor []common.Address, delegatee []common.Address) (*StaminaWithdrawnIterator, error) {

	var depositorRule []interface{}
	for _, depositorItem := range depositor {
		depositorRule = append(depositorRule, depositorItem)
	}
	var delegateeRule []interface{}
	for _, delegateeItem := range delegatee {
		delegateeRule = append(delegateeRule, delegateeItem)
	}

	logs, sub, err := _Stamina.contract.FilterLogs(opts, "Withdrawn", depositorRule, delegateeRule)
	if err != nil {
		return nil, err
	}
	return &StaminaWithdrawnIterator{contract: _Stamina.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0xd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb.
//
// Solidity: e Withdrawn(depositor indexed address, delegatee indexed address, amount uint256)
func (_Stamina *StaminaFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *StaminaWithdrawn, depositor []common.Address, delegatee []common.Address) (event.Subscription, error) {

	var depositorRule []interface{}
	for _, depositorItem := range depositor {
		depositorRule = append(depositorRule, depositorItem)
	}
	var delegateeRule []interface{}
	for _, delegateeItem := range delegatee {
		delegateeRule = append(delegateeRule, delegateeItem)
	}

	logs, sub, err := _Stamina.contract.WatchLogs(opts, "Withdrawn", depositorRule, delegateeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StaminaWithdrawn)
				if err := _Stamina.contract.UnpackLog(event, "Withdrawn", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}
