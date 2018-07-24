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
const StaminaABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"WITHDRAWAL_DELAY\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"RECOVER_EPOCH_LENGTH\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"initialized\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getStamina\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"depositor\",\"type\":\"address\"},{\"name\":\"withdrawalIndex\",\"type\":\"uint256\"}],\"name\":\"getWithdrawal\",\"outputs\":[{\"name\":\"amount\",\"type\":\"uint128\"},{\"name\":\"requestBlockNumber\",\"type\":\"uint128\"},{\"name\":\"delegatee\",\"type\":\"address\"},{\"name\":\"processed\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"development\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"delegatee\",\"type\":\"address\"}],\"name\":\"getTotalDeposit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"minDeposit\",\"type\":\"uint256\"},{\"name\":\"recoveryEpochLength\",\"type\":\"uint256\"},{\"name\":\"withdrawalDelay\",\"type\":\"uint256\"}],\"name\":\"init\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"delegatee\",\"type\":\"address\"}],\"name\":\"getLastRecoveryBlock\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"delegator\",\"type\":\"address\"}],\"name\":\"getDelegatee\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"delegatee\",\"type\":\"address\"}],\"name\":\"getNumRecovery\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"delegatee\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"addStamina\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"depositor\",\"type\":\"address\"},{\"name\":\"delegatee\",\"type\":\"address\"}],\"name\":\"getDeposit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"delegatee\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"subtractStamina\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"depositor\",\"type\":\"address\"}],\"name\":\"getNumWithdrawals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"delegatee\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"requestWithdrawal\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MIN_DEPOSIT\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"delegator\",\"type\":\"address\"}],\"name\":\"setDelegatee\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"delegatee\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"depositor\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"delegatee\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"oldDelegatee\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"newDelegatee\",\"type\":\"address\"}],\"name\":\"DelegateeChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"depositor\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"delegatee\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"requestBlockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"withdrawalIndex\",\"type\":\"uint256\"}],\"name\":\"WithdrawalRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"depositor\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"delegatee\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"withdrawalIndex\",\"type\":\"uint256\"}],\"name\":\"Withdrawan\",\"type\":\"event\"}]"

// StaminaBin is the compiled bytecode used for deploying new contracts.
const StaminaBin = `0x6080604052600c805460ff1916600117905534801561001d57600080fd5b50610f5b8061002d6000396000f3006080604052600436106101115763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416630ebb172a81146101165780631556d8ac1461013d578063158ef93e146101525780633900e4ec1461017b5780633ccfd60b146101ae5780635be4f765146101c35780637b929c2714610241578063857184d1146102565780638cd8db8a14610289578063937aaef1146102c15780639b4e735f146102f4578063b69ad63b14610343578063bcac973614610376578063c35082a9146103af578063d1c0c042146103ea578063d898ae1c14610423578063da95ebf714610456578063e1e158a51461048f578063e842a64b146104a4578063f340fa01146104d7575b600080fd5b34801561012257600080fd5b5061012b6104fd565b60408051918252519081900360200190f35b34801561014957600080fd5b5061012b610503565b34801561015e57600080fd5b50610167610509565b604080519115158252519081900360200190f35b34801561018757600080fd5b5061012b6004803603602081101561019e57600080fd5b5035600160a060020a0316610512565b3480156101ba57600080fd5b5061016761052d565b3480156101cf57600080fd5b506101fc600480360360408110156101e657600080fd5b50600160a060020a038135169060200135610737565b604080516fffffffffffffffffffffffffffffffff9586168152939094166020840152600160a060020a03909116828401521515606082015290519081900360800190f35b34801561024d57600080fd5b5061016761081a565b34801561026257600080fd5b5061012b6004803603602081101561027957600080fd5b5035600160a060020a0316610823565b34801561029557600080fd5b506102bf600480360360608110156102ac57600080fd5b508035906020810135906040013561083e565b005b3480156102cd57600080fd5b5061012b600480360360208110156102e457600080fd5b5035600160a060020a031661089f565b34801561030057600080fd5b506103276004803603602081101561031757600080fd5b5035600160a060020a03166108ba565b60408051600160a060020a039092168252519081900360200190f35b34801561034f57600080fd5b5061012b6004803603602081101561036657600080fd5b5035600160a060020a03166108d8565b34801561038257600080fd5b506101676004803603604081101561039957600080fd5b50600160a060020a0381351690602001356108f3565b3480156103bb57600080fd5b5061012b600480360360408110156103d257600080fd5b50600160a060020a0381358116916020013516610a02565b3480156103f657600080fd5b506101676004803603604081101561040d57600080fd5b50600160a060020a038135169060200135610a2d565b34801561042f57600080fd5b5061012b6004803603602081101561044657600080fd5b5035600160a060020a0316610a99565b34801561046257600080fd5b506101676004803603604081101561047957600080fd5b50600160a060020a038135169060200135610ab4565b34801561049b57600080fd5b5061012b610cba565b3480156104b057600080fd5b50610167600480360360208110156104c757600080fd5b5035600160a060020a0316610cc0565b610167600480360360208110156104ed57600080fd5b5035600160a060020a0316610d59565b600b5481565b600a5481565b60085460ff1681565b600160a060020a031660009081526001602052604090205490565b3360009081526005602052604081208054821061054957600080fd5b336000908152600660205260408120549081158015610593575082600081548110151561057257fe5b906000526020600020906002020160010160149054906101000a900460ff16155b156105a0575060006105c4565b8115156105be578254600211156105b657600080fd5b5060016105c4565b50600181015b825481106105d157600080fd5b3360009081526005602052604081208054839081106105ec57fe5b906000526020600020906002020190508060010160149054906101000a900460ff1615151561061a57600080fd5b600b548154437001000000000000000000000000000000009091046fffffffffffffffffffffffffffffffff16909101111561065557600080fd5b805460018201805474ff000000000000000000000000000000000000000019167401000000000000000000000000000000000000000017905533600081815260066020526040808220869055516fffffffffffffffffffffffffffffffff9093169283156108fc0291849190818181858888f193505050501580156106de573d6000803e3d6000fd5b50600182015460408051838152602081018690528151600160a060020a039093169233927ff105379f79c77d26f97e62cc481d8493005422dde6ad256be000ea973120e879928290030190a36001955050505050505b90565b60008060008061074686610a99565b851061075157600080fd5b610759610e99565b600160a060020a038716600090815260056020526040902080548790811061077d57fe5b600091825260209182902060408051608081018252600290930290910180546fffffffffffffffffffffffffffffffff80821680865270010000000000000000000000000000000090920416948401859052600190910154600160a060020a03811692840183905260ff7401000000000000000000000000000000000000000090910416151560609093018390529a929950975095509350505050565b600c5460ff1681565b600160a060020a031660009081526002602052604090205490565b60085460ff161561084e57600080fd5b6000831161085b57600080fd5b6000821161086857600080fd5b6000811161087557600080fd5b60028202811161088457600080fd5b600992909255600a55600b556008805460ff19166001179055565b600160a060020a031660009081526004602052604090205490565b600160a060020a039081166000908152602081905260409020541690565b600160a060020a031660009081526007602052604090205490565b600c5460009060ff1680610905575033155b151561091057600080fd5b600a54600160a060020a0384166000908152600460205260409020544391011161097a5750600160a060020a0382166000908152600260209081526040808320546001808452828520919091556004835281842043905560079092529091208054820190556109fc565b600160a060020a03831660009081526002602090815260408083205460019092529091205483810181106109ad57600080fd5b808401828111156109d857600160a060020a03861660009081526001602052604090208390556109f4565b600160a060020a03861660009081526001602052604090208190555b600193505050505b92915050565b600160a060020a03918216600090815260036020908152604080832093909416825291909152205490565b600c5460009060ff1680610a3f575033155b1515610a4a57600080fd5b600160a060020a0383166000908152600160205260409020548281038111610a7157600080fd5b600160a060020a03939093166000908152600160208190526040909120929093039091555090565b600160a060020a031660009081526005602052604090205490565b60085460009060ff161515610ac857600080fd5b60008211610ad557600080fd5b600160a060020a0383166000818152600260209081526040808320543384526003835281842094845293825280832054600190925282205490918211610b1a57600080fd5b8483038311610b2857600080fd5b8482038211610b3657600080fd5b600160a060020a038616600081815260026020908152604080832089880390553383526003825280832093835292905220858303905584811115610b9657600160a060020a03861660009081526001602052604090208582039055610bb0565b600160a060020a0386166000908152600160205260408120555b336000908152600560205260408120805490918282610bd28260018301610ec0565b81548110610bdc57fe5b60009182526020918290206002919091020180546fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff8b8116919091178116700100000000000000000000000000000000439283160217825560018201805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a038e16908117909155604080518d81529485019290925283820186905290519193509133917f3aeb15af61588a39bcfafb19ed853140d195c2a924537afbf9a6d04348e76a69916060908290030190a350600198975050505050505050565b60095481565b60085460009060ff161515610cd457600080fd5b600160a060020a0380831660008181526020818152604091829020805473ffffffffffffffffffffffffffffffffffffffff19811633908117909255835194855290941690830181905282820193909352517f5884d7e3ec123de8e772bcf576c18dcdad75b056c4314f999ed966693419c6929181900360600190a150600192915050565b60085460009060ff161515610d6d57600080fd5b600954341015610d7c57600080fd5b600160a060020a03821660008181526002602090815260408083205433845260038352818420948452938252808320546001909252909120543483018310610dc357600080fd5b3482018210610dd157600080fd5b3481018110610ddf57600080fd5b600160a060020a038516600081815260026020908152604080832034888101909155338452600383528184209484529382528083208685019055600182528083209385019093556004905220541515610e4e57600160a060020a03851660009081526004602052604090204390555b604080513481529051600160a060020a0387169133917f8752a472e571a816aea92eec8dae9baf628e840f4929fbcc2d155e6233ff68a79181900360200190a3506001949350505050565b60408051608081018252600080825260208201819052918101829052606081019190915290565b815481835581811115610eec57600202816002028360005260206000209182019101610eec9190610ef1565b505050565b61073491905b80821115610f2b576000815560018101805474ffffffffffffffffffffffffffffffffffffffffff19169055600201610ef7565b50905600a165627a7a72305820a02bfd230d4c61a1e1f620f19aa68a9f2b4e21138813f3eb509c3191aa204f1b0029`

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

// RECOVEREPOCHLENGTH is a free data retrieval call binding the contract method 0x1556d8ac.
//
// Solidity: function RECOVER_EPOCH_LENGTH() constant returns(uint256)
func (_Stamina *StaminaCaller) RECOVEREPOCHLENGTH(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Stamina.contract.Call(opts, out, "RECOVER_EPOCH_LENGTH")
	return *ret0, err
}

// RECOVEREPOCHLENGTH is a free data retrieval call binding the contract method 0x1556d8ac.
//
// Solidity: function RECOVER_EPOCH_LENGTH() constant returns(uint256)
func (_Stamina *StaminaSession) RECOVEREPOCHLENGTH() (*big.Int, error) {
	return _Stamina.Contract.RECOVEREPOCHLENGTH(&_Stamina.CallOpts)
}

// RECOVEREPOCHLENGTH is a free data retrieval call binding the contract method 0x1556d8ac.
//
// Solidity: function RECOVER_EPOCH_LENGTH() constant returns(uint256)
func (_Stamina *StaminaCallerSession) RECOVEREPOCHLENGTH() (*big.Int, error) {
	return _Stamina.Contract.RECOVEREPOCHLENGTH(&_Stamina.CallOpts)
}

// WITHDRAWALDELAY is a free data retrieval call binding the contract method 0x0ebb172a.
//
// Solidity: function WITHDRAWAL_DELAY() constant returns(uint256)
func (_Stamina *StaminaCaller) WITHDRAWALDELAY(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Stamina.contract.Call(opts, out, "WITHDRAWAL_DELAY")
	return *ret0, err
}

// WITHDRAWALDELAY is a free data retrieval call binding the contract method 0x0ebb172a.
//
// Solidity: function WITHDRAWAL_DELAY() constant returns(uint256)
func (_Stamina *StaminaSession) WITHDRAWALDELAY() (*big.Int, error) {
	return _Stamina.Contract.WITHDRAWALDELAY(&_Stamina.CallOpts)
}

// WITHDRAWALDELAY is a free data retrieval call binding the contract method 0x0ebb172a.
//
// Solidity: function WITHDRAWAL_DELAY() constant returns(uint256)
func (_Stamina *StaminaCallerSession) WITHDRAWALDELAY() (*big.Int, error) {
	return _Stamina.Contract.WITHDRAWALDELAY(&_Stamina.CallOpts)
}

// Development is a free data retrieval call binding the contract method 0x7b929c27.
//
// Solidity: function development() constant returns(bool)
func (_Stamina *StaminaCaller) Development(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Stamina.contract.Call(opts, out, "development")
	return *ret0, err
}

// Development is a free data retrieval call binding the contract method 0x7b929c27.
//
// Solidity: function development() constant returns(bool)
func (_Stamina *StaminaSession) Development() (bool, error) {
	return _Stamina.Contract.Development(&_Stamina.CallOpts)
}

// Development is a free data retrieval call binding the contract method 0x7b929c27.
//
// Solidity: function development() constant returns(bool)
func (_Stamina *StaminaCallerSession) Development() (bool, error) {
	return _Stamina.Contract.Development(&_Stamina.CallOpts)
}

// GetDelegatee is a free data retrieval call binding the contract method 0x9b4e735f.
//
// Solidity: function getDelegatee(delegator address) constant returns(address)
func (_Stamina *StaminaCaller) GetDelegatee(opts *bind.CallOpts, delegator common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Stamina.contract.Call(opts, out, "getDelegatee", delegator)
	return *ret0, err
}

// GetDelegatee is a free data retrieval call binding the contract method 0x9b4e735f.
//
// Solidity: function getDelegatee(delegator address) constant returns(address)
func (_Stamina *StaminaSession) GetDelegatee(delegator common.Address) (common.Address, error) {
	return _Stamina.Contract.GetDelegatee(&_Stamina.CallOpts, delegator)
}

// GetDelegatee is a free data retrieval call binding the contract method 0x9b4e735f.
//
// Solidity: function getDelegatee(delegator address) constant returns(address)
func (_Stamina *StaminaCallerSession) GetDelegatee(delegator common.Address) (common.Address, error) {
	return _Stamina.Contract.GetDelegatee(&_Stamina.CallOpts, delegator)
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

// GetLastRecoveryBlock is a free data retrieval call binding the contract method 0x937aaef1.
//
// Solidity: function getLastRecoveryBlock(delegatee address) constant returns(uint256)
func (_Stamina *StaminaCaller) GetLastRecoveryBlock(opts *bind.CallOpts, delegatee common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Stamina.contract.Call(opts, out, "getLastRecoveryBlock", delegatee)
	return *ret0, err
}

// GetLastRecoveryBlock is a free data retrieval call binding the contract method 0x937aaef1.
//
// Solidity: function getLastRecoveryBlock(delegatee address) constant returns(uint256)
func (_Stamina *StaminaSession) GetLastRecoveryBlock(delegatee common.Address) (*big.Int, error) {
	return _Stamina.Contract.GetLastRecoveryBlock(&_Stamina.CallOpts, delegatee)
}

// GetLastRecoveryBlock is a free data retrieval call binding the contract method 0x937aaef1.
//
// Solidity: function getLastRecoveryBlock(delegatee address) constant returns(uint256)
func (_Stamina *StaminaCallerSession) GetLastRecoveryBlock(delegatee common.Address) (*big.Int, error) {
	return _Stamina.Contract.GetLastRecoveryBlock(&_Stamina.CallOpts, delegatee)
}

// GetNumRecovery is a free data retrieval call binding the contract method 0xb69ad63b.
//
// Solidity: function getNumRecovery(delegatee address) constant returns(uint256)
func (_Stamina *StaminaCaller) GetNumRecovery(opts *bind.CallOpts, delegatee common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Stamina.contract.Call(opts, out, "getNumRecovery", delegatee)
	return *ret0, err
}

// GetNumRecovery is a free data retrieval call binding the contract method 0xb69ad63b.
//
// Solidity: function getNumRecovery(delegatee address) constant returns(uint256)
func (_Stamina *StaminaSession) GetNumRecovery(delegatee common.Address) (*big.Int, error) {
	return _Stamina.Contract.GetNumRecovery(&_Stamina.CallOpts, delegatee)
}

// GetNumRecovery is a free data retrieval call binding the contract method 0xb69ad63b.
//
// Solidity: function getNumRecovery(delegatee address) constant returns(uint256)
func (_Stamina *StaminaCallerSession) GetNumRecovery(delegatee common.Address) (*big.Int, error) {
	return _Stamina.Contract.GetNumRecovery(&_Stamina.CallOpts, delegatee)
}

// GetNumWithdrawals is a free data retrieval call binding the contract method 0xd898ae1c.
//
// Solidity: function getNumWithdrawals(depositor address) constant returns(uint256)
func (_Stamina *StaminaCaller) GetNumWithdrawals(opts *bind.CallOpts, depositor common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Stamina.contract.Call(opts, out, "getNumWithdrawals", depositor)
	return *ret0, err
}

// GetNumWithdrawals is a free data retrieval call binding the contract method 0xd898ae1c.
//
// Solidity: function getNumWithdrawals(depositor address) constant returns(uint256)
func (_Stamina *StaminaSession) GetNumWithdrawals(depositor common.Address) (*big.Int, error) {
	return _Stamina.Contract.GetNumWithdrawals(&_Stamina.CallOpts, depositor)
}

// GetNumWithdrawals is a free data retrieval call binding the contract method 0xd898ae1c.
//
// Solidity: function getNumWithdrawals(depositor address) constant returns(uint256)
func (_Stamina *StaminaCallerSession) GetNumWithdrawals(depositor common.Address) (*big.Int, error) {
	return _Stamina.Contract.GetNumWithdrawals(&_Stamina.CallOpts, depositor)
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

// GetWithdrawal is a free data retrieval call binding the contract method 0x5be4f765.
//
// Solidity: function getWithdrawal(depositor address, withdrawalIndex uint256) constant returns(amount uint128, requestBlockNumber uint128, delegatee address, processed bool)
func (_Stamina *StaminaCaller) GetWithdrawal(opts *bind.CallOpts, depositor common.Address, withdrawalIndex *big.Int) (struct {
	Amount             *big.Int
	RequestBlockNumber *big.Int
	Delegatee          common.Address
	Processed          bool
}, error) {
	ret := new(struct {
		Amount             *big.Int
		RequestBlockNumber *big.Int
		Delegatee          common.Address
		Processed          bool
	})
	out := ret
	err := _Stamina.contract.Call(opts, out, "getWithdrawal", depositor, withdrawalIndex)
	return *ret, err
}

// GetWithdrawal is a free data retrieval call binding the contract method 0x5be4f765.
//
// Solidity: function getWithdrawal(depositor address, withdrawalIndex uint256) constant returns(amount uint128, requestBlockNumber uint128, delegatee address, processed bool)
func (_Stamina *StaminaSession) GetWithdrawal(depositor common.Address, withdrawalIndex *big.Int) (struct {
	Amount             *big.Int
	RequestBlockNumber *big.Int
	Delegatee          common.Address
	Processed          bool
}, error) {
	return _Stamina.Contract.GetWithdrawal(&_Stamina.CallOpts, depositor, withdrawalIndex)
}

// GetWithdrawal is a free data retrieval call binding the contract method 0x5be4f765.
//
// Solidity: function getWithdrawal(depositor address, withdrawalIndex uint256) constant returns(amount uint128, requestBlockNumber uint128, delegatee address, processed bool)
func (_Stamina *StaminaCallerSession) GetWithdrawal(depositor common.Address, withdrawalIndex *big.Int) (struct {
	Amount             *big.Int
	RequestBlockNumber *big.Int
	Delegatee          common.Address
	Processed          bool
}, error) {
	return _Stamina.Contract.GetWithdrawal(&_Stamina.CallOpts, depositor, withdrawalIndex)
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

// Init is a paid mutator transaction binding the contract method 0x8cd8db8a.
//
// Solidity: function init(minDeposit uint256, recoveryEpochLength uint256, withdrawalDelay uint256) returns()
func (_Stamina *StaminaTransactor) Init(opts *bind.TransactOpts, minDeposit *big.Int, recoveryEpochLength *big.Int, withdrawalDelay *big.Int) (*types.Transaction, error) {
	return _Stamina.contract.Transact(opts, "init", minDeposit, recoveryEpochLength, withdrawalDelay)
}

// Init is a paid mutator transaction binding the contract method 0x8cd8db8a.
//
// Solidity: function init(minDeposit uint256, recoveryEpochLength uint256, withdrawalDelay uint256) returns()
func (_Stamina *StaminaSession) Init(minDeposit *big.Int, recoveryEpochLength *big.Int, withdrawalDelay *big.Int) (*types.Transaction, error) {
	return _Stamina.Contract.Init(&_Stamina.TransactOpts, minDeposit, recoveryEpochLength, withdrawalDelay)
}

// Init is a paid mutator transaction binding the contract method 0x8cd8db8a.
//
// Solidity: function init(minDeposit uint256, recoveryEpochLength uint256, withdrawalDelay uint256) returns()
func (_Stamina *StaminaTransactorSession) Init(minDeposit *big.Int, recoveryEpochLength *big.Int, withdrawalDelay *big.Int) (*types.Transaction, error) {
	return _Stamina.Contract.Init(&_Stamina.TransactOpts, minDeposit, recoveryEpochLength, withdrawalDelay)
}

// RequestWithdrawal is a paid mutator transaction binding the contract method 0xda95ebf7.
//
// Solidity: function requestWithdrawal(delegatee address, amount uint256) returns(bool)
func (_Stamina *StaminaTransactor) RequestWithdrawal(opts *bind.TransactOpts, delegatee common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Stamina.contract.Transact(opts, "requestWithdrawal", delegatee, amount)
}

// RequestWithdrawal is a paid mutator transaction binding the contract method 0xda95ebf7.
//
// Solidity: function requestWithdrawal(delegatee address, amount uint256) returns(bool)
func (_Stamina *StaminaSession) RequestWithdrawal(delegatee common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Stamina.Contract.RequestWithdrawal(&_Stamina.TransactOpts, delegatee, amount)
}

// RequestWithdrawal is a paid mutator transaction binding the contract method 0xda95ebf7.
//
// Solidity: function requestWithdrawal(delegatee address, amount uint256) returns(bool)
func (_Stamina *StaminaTransactorSession) RequestWithdrawal(delegatee common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Stamina.Contract.RequestWithdrawal(&_Stamina.TransactOpts, delegatee, amount)
}

// SetDelegatee is a paid mutator transaction binding the contract method 0xe842a64b.
//
// Solidity: function setDelegatee(delegator address) returns(bool)
func (_Stamina *StaminaTransactor) SetDelegatee(opts *bind.TransactOpts, delegator common.Address) (*types.Transaction, error) {
	return _Stamina.contract.Transact(opts, "setDelegatee", delegator)
}

// SetDelegatee is a paid mutator transaction binding the contract method 0xe842a64b.
//
// Solidity: function setDelegatee(delegator address) returns(bool)
func (_Stamina *StaminaSession) SetDelegatee(delegator common.Address) (*types.Transaction, error) {
	return _Stamina.Contract.SetDelegatee(&_Stamina.TransactOpts, delegator)
}

// SetDelegatee is a paid mutator transaction binding the contract method 0xe842a64b.
//
// Solidity: function setDelegatee(delegator address) returns(bool)
func (_Stamina *StaminaTransactorSession) SetDelegatee(delegator common.Address) (*types.Transaction, error) {
	return _Stamina.Contract.SetDelegatee(&_Stamina.TransactOpts, delegator)
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

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns(bool)
func (_Stamina *StaminaTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stamina.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns(bool)
func (_Stamina *StaminaSession) Withdraw() (*types.Transaction, error) {
	return _Stamina.Contract.Withdraw(&_Stamina.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns(bool)
func (_Stamina *StaminaTransactorSession) Withdraw() (*types.Transaction, error) {
	return _Stamina.Contract.Withdraw(&_Stamina.TransactOpts)
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
	Delegator    common.Address
	OldDelegatee common.Address
	NewDelegatee common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDelegateeChanged is a free log retrieval operation binding the contract event 0x5884d7e3ec123de8e772bcf576c18dcdad75b056c4314f999ed966693419c692.
//
// Solidity: e DelegateeChanged(delegator address, oldDelegatee address, newDelegatee address)
func (_Stamina *StaminaFilterer) FilterDelegateeChanged(opts *bind.FilterOpts) (*StaminaDelegateeChangedIterator, error) {

	logs, sub, err := _Stamina.contract.FilterLogs(opts, "DelegateeChanged")
	if err != nil {
		return nil, err
	}
	return &StaminaDelegateeChangedIterator{contract: _Stamina.contract, event: "DelegateeChanged", logs: logs, sub: sub}, nil
}

// WatchDelegateeChanged is a free log subscription operation binding the contract event 0x5884d7e3ec123de8e772bcf576c18dcdad75b056c4314f999ed966693419c692.
//
// Solidity: e DelegateeChanged(delegator address, oldDelegatee address, newDelegatee address)
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

// StaminaWithdrawalRequestedIterator is returned from FilterWithdrawalRequested and is used to iterate over the raw logs and unpacked data for WithdrawalRequested events raised by the Stamina contract.
type StaminaWithdrawalRequestedIterator struct {
	Event *StaminaWithdrawalRequested // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StaminaWithdrawalRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StaminaWithdrawalRequested)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StaminaWithdrawalRequested)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StaminaWithdrawalRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StaminaWithdrawalRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StaminaWithdrawalRequested represents a WithdrawalRequested event raised by the Stamina contract.
type StaminaWithdrawalRequested struct {
	Depositor          common.Address
	Delegatee          common.Address
	Amount             *big.Int
	RequestBlockNumber *big.Int
	WithdrawalIndex    *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalRequested is a free log retrieval operation binding the contract event 0x3aeb15af61588a39bcfafb19ed853140d195c2a924537afbf9a6d04348e76a69.
//
// Solidity: e WithdrawalRequested(depositor indexed address, delegatee indexed address, amount uint256, requestBlockNumber uint256, withdrawalIndex uint256)
func (_Stamina *StaminaFilterer) FilterWithdrawalRequested(opts *bind.FilterOpts, depositor []common.Address, delegatee []common.Address) (*StaminaWithdrawalRequestedIterator, error) {

	var depositorRule []interface{}
	for _, depositorItem := range depositor {
		depositorRule = append(depositorRule, depositorItem)
	}
	var delegateeRule []interface{}
	for _, delegateeItem := range delegatee {
		delegateeRule = append(delegateeRule, delegateeItem)
	}

	logs, sub, err := _Stamina.contract.FilterLogs(opts, "WithdrawalRequested", depositorRule, delegateeRule)
	if err != nil {
		return nil, err
	}
	return &StaminaWithdrawalRequestedIterator{contract: _Stamina.contract, event: "WithdrawalRequested", logs: logs, sub: sub}, nil
}

// WatchWithdrawalRequested is a free log subscription operation binding the contract event 0x3aeb15af61588a39bcfafb19ed853140d195c2a924537afbf9a6d04348e76a69.
//
// Solidity: e WithdrawalRequested(depositor indexed address, delegatee indexed address, amount uint256, requestBlockNumber uint256, withdrawalIndex uint256)
func (_Stamina *StaminaFilterer) WatchWithdrawalRequested(opts *bind.WatchOpts, sink chan<- *StaminaWithdrawalRequested, depositor []common.Address, delegatee []common.Address) (event.Subscription, error) {

	var depositorRule []interface{}
	for _, depositorItem := range depositor {
		depositorRule = append(depositorRule, depositorItem)
	}
	var delegateeRule []interface{}
	for _, delegateeItem := range delegatee {
		delegateeRule = append(delegateeRule, delegateeItem)
	}

	logs, sub, err := _Stamina.contract.WatchLogs(opts, "WithdrawalRequested", depositorRule, delegateeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StaminaWithdrawalRequested)
				if err := _Stamina.contract.UnpackLog(event, "WithdrawalRequested", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// StaminaWithdrawanIterator is returned from FilterWithdrawan and is used to iterate over the raw logs and unpacked data for Withdrawan events raised by the Stamina contract.
type StaminaWithdrawanIterator struct {
	Event *StaminaWithdrawan // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StaminaWithdrawanIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StaminaWithdrawan)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StaminaWithdrawan)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StaminaWithdrawanIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StaminaWithdrawanIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StaminaWithdrawan represents a Withdrawan event raised by the Stamina contract.
type StaminaWithdrawan struct {
	Depositor       common.Address
	Delegatee       common.Address
	Amount          *big.Int
	WithdrawalIndex *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterWithdrawan is a free log retrieval operation binding the contract event 0xf105379f79c77d26f97e62cc481d8493005422dde6ad256be000ea973120e879.
//
// Solidity: e Withdrawan(depositor indexed address, delegatee indexed address, amount uint256, withdrawalIndex uint256)
func (_Stamina *StaminaFilterer) FilterWithdrawan(opts *bind.FilterOpts, depositor []common.Address, delegatee []common.Address) (*StaminaWithdrawanIterator, error) {

	var depositorRule []interface{}
	for _, depositorItem := range depositor {
		depositorRule = append(depositorRule, depositorItem)
	}
	var delegateeRule []interface{}
	for _, delegateeItem := range delegatee {
		delegateeRule = append(delegateeRule, delegateeItem)
	}

	logs, sub, err := _Stamina.contract.FilterLogs(opts, "Withdrawan", depositorRule, delegateeRule)
	if err != nil {
		return nil, err
	}
	return &StaminaWithdrawanIterator{contract: _Stamina.contract, event: "Withdrawan", logs: logs, sub: sub}, nil
}

// WatchWithdrawan is a free log subscription operation binding the contract event 0xf105379f79c77d26f97e62cc481d8493005422dde6ad256be000ea973120e879.
//
// Solidity: e Withdrawan(depositor indexed address, delegatee indexed address, amount uint256, withdrawalIndex uint256)
func (_Stamina *StaminaFilterer) WatchWithdrawan(opts *bind.WatchOpts, sink chan<- *StaminaWithdrawan, depositor []common.Address, delegatee []common.Address) (event.Subscription, error) {

	var depositorRule []interface{}
	for _, depositorItem := range depositor {
		depositorRule = append(depositorRule, depositorItem)
	}
	var delegateeRule []interface{}
	for _, delegateeItem := range delegatee {
		delegateeRule = append(delegateeRule, delegateeItem)
	}

	logs, sub, err := _Stamina.contract.WatchLogs(opts, "Withdrawan", depositorRule, delegateeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StaminaWithdrawan)
				if err := _Stamina.contract.UnpackLog(event, "Withdrawan", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}
