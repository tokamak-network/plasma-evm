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

// StaminaABI is the input ABI used to generate the binding from.
const StaminaABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"WITHDRAWAL_DELAY\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"RECOVER_EPOCH_LENGTH\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"initialized\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getStamina\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"depositor\",\"type\":\"address\"},{\"name\":\"withdrawalIndex\",\"type\":\"uint256\"}],\"name\":\"getWithdrawal\",\"outputs\":[{\"name\":\"amount\",\"type\":\"uint128\"},{\"name\":\"requestBlockNumber\",\"type\":\"uint128\"},{\"name\":\"delegatee\",\"type\":\"address\"},{\"name\":\"processed\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"development\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"delegator\",\"type\":\"address\"}],\"name\":\"setDelegator\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"delegatee\",\"type\":\"address\"}],\"name\":\"getTotalDeposit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"minDeposit\",\"type\":\"uint256\"},{\"name\":\"recoveryEpochLength\",\"type\":\"uint256\"},{\"name\":\"withdrawalDelay\",\"type\":\"uint256\"}],\"name\":\"init\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"delegatee\",\"type\":\"address\"}],\"name\":\"getLastRecoveryBlock\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"delegator\",\"type\":\"address\"}],\"name\":\"getDelegatee\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"delegatee\",\"type\":\"address\"}],\"name\":\"getNumRecovery\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"delegatee\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"addStamina\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"depositor\",\"type\":\"address\"},{\"name\":\"delegatee\",\"type\":\"address\"}],\"name\":\"getDeposit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"delegatee\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"subtractStamina\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"depositor\",\"type\":\"address\"}],\"name\":\"getNumWithdrawals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"delegatee\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"requestWithdrawal\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MIN_DEPOSIT\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"delegatee\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"depositor\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"delegatee\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"oldDelegatee\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"newDelegatee\",\"type\":\"address\"}],\"name\":\"DelegateeChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"depositor\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"delegatee\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"requestBlockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"withdrawalIndex\",\"type\":\"uint256\"}],\"name\":\"WithdrawalRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"depositor\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"delegatee\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"withdrawalIndex\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"delegatee\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"recovered\",\"type\":\"bool\"}],\"name\":\"StaminaAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"delegatee\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"StaminaSubtracted\",\"type\":\"event\"}]"

// StaminaBin is the compiled bytecode used for deploying new contracts.
const StaminaBin = `0x6080604052600c805460ff1916600117905534801561001d57600080fd5b50610f388061002d6000396000f3006080604052600436106101115763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416630ebb172a81146101165780631556d8ac1461013d578063158ef93e146101525780633900e4ec1461017b5780633ccfd60b1461019c5780635be4f765146101b15780637b929c271461021a57806383cd9cc31461022f578063857184d1146102505780638cd8db8a14610271578063937aaef1146102915780639b4e735f146102b2578063b69ad63b146102ef578063bcac973614610310578063c35082a914610334578063d1c0c0421461035b578063d898ae1c1461037f578063da95ebf7146103a0578063e1e158a5146103c4578063f340fa01146103d9575b600080fd5b34801561012257600080fd5b5061012b6103ed565b60408051918252519081900360200190f35b34801561014957600080fd5b5061012b6103f3565b34801561015e57600080fd5b506101676103f9565b604080519115158252519081900360200190f35b34801561018757600080fd5b5061012b600160a060020a0360043516610402565b3480156101a857600080fd5b5061016761041d565b3480156101bd57600080fd5b506101d5600160a060020a0360043516602435610633565b604080516fffffffffffffffffffffffffffffffff9586168152939094166020840152600160a060020a03909116828401521515606082015290519081900360800190f35b34801561022657600080fd5b50610167610716565b34801561023b57600080fd5b50610167600160a060020a036004351661071f565b34801561025c57600080fd5b5061012b600160a060020a03600435166107b2565b34801561027d57600080fd5b5061028f6004356024356044356107cd565b005b34801561029d57600080fd5b5061012b600160a060020a036004351661082e565b3480156102be57600080fd5b506102d3600160a060020a0360043516610849565b60408051600160a060020a039092168252519081900360200190f35b3480156102fb57600080fd5b5061012b600160a060020a0360043516610867565b34801561031c57600080fd5b50610167600160a060020a0360043516602435610882565b34801561034057600080fd5b5061012b600160a060020a0360043581169060243516610a1a565b34801561036757600080fd5b50610167600160a060020a0360043516602435610a45565b34801561038b57600080fd5b5061012b600160a060020a0360043516610ae8565b3480156103ac57600080fd5b50610167600160a060020a0360043516602435610b03565b3480156103d057600080fd5b5061012b610d24565b610167600160a060020a0360043516610d2a565b600b5481565b600a5481565b60085460ff1681565b600160a060020a031660009081526001602052604090205490565b33600090815260056020526040812080548290819081908190811061044157600080fd5b3360009081526006602052604090205493508315801561048c575084600081548110151561046b57fe5b906000526020600020906002020160010160149054906101000a900460ff16155b1561049a57600092506104c0565b8315156104b9578454600211156104b057600080fd5b600192506104c0565b8360010192505b845483106104cd57600080fd5b3360009081526005602052604090208054849081106104e857fe5b906000526020600020906002020191508160010160149054906101000a900460ff1615151561051657600080fd5b600b548254437001000000000000000000000000000000009091046fffffffffffffffffffffffffffffffff16909101111561055157600080fd5b50805460018201805474ff000000000000000000000000000000000000000019167401000000000000000000000000000000000000000017905533600081815260066020526040808220869055516fffffffffffffffffffffffffffffffff9093169283156108fc0291849190818181858888f193505050501580156105db573d6000803e3d6000fd5b50600182015460408051838152602081018690528151600160a060020a039093169233927f91fb9d98b786c57d74c099ccd2beca1739e9f6a81fb49001ca465c4b7591bbe2928290030190a360019550505050505090565b600080600080610641610e73565b61064a87610ae8565b861061065557600080fd5b600160a060020a038716600090815260056020526040902080548790811061067957fe5b600091825260209182902060408051608081018252600290930290910180546fffffffffffffffffffffffffffffffff80821680865270010000000000000000000000000000000090920416948401859052600190910154600160a060020a03811692840183905260ff7401000000000000000000000000000000000000000090910416151560609093018390529a929950975095509350505050565b600c5460ff1681565b600854600090819060ff16151561073557600080fd5b50600160a060020a038281166000818152602081815260409182902080543373ffffffffffffffffffffffffffffffffffffffff19821681179092558351951680865291850152815190937f5884d7e3ec123de8e772bcf576c18dcdad75b056c4314f999ed966693419c69292908290030190a250600192915050565b600160a060020a031660009081526002602052604090205490565b60085460ff16156107dd57600080fd5b600083116107ea57600080fd5b600082116107f757600080fd5b6000811161080457600080fd5b60028202811161081357600080fd5b600992909255600a55600b556008805460ff19166001179055565b600160a060020a031660009081526004602052604090205490565b600160a060020a039081166000908152602081905260409020541690565b600160a060020a031660009081526007602052604090205490565b600c5460009081908190819060ff168061089a575033155b15156108a557600080fd5b600a54600160a060020a0387166000908152600460205260409020544391011161094957600160a060020a038616600081815260026020908152604080832054600180845282852091909155600483528184204390556007835281842080548201905581519384529183019190915280517f85bf8701ef98ea32e97f08708da81c7daa93e87ea3e2fd661801cce6d36f68099281900390910190a260019350610a11565b600160a060020a0386166000908152600260209081526040808320546001909252909120549093509150848201821061098157600080fd5b50808401828111156109ad57600160a060020a03861660009081526001602052604090208390556109c9565b600160a060020a03861660009081526001602052604090208190555b60408051868152600060208201528151600160a060020a038916927f85bf8701ef98ea32e97f08708da81c7daa93e87ea3e2fd661801cce6d36f6809928290030190a2600193505b50505092915050565b600160a060020a03918216600090815260036020908152604080832093909416825291909152205490565b600c54600090819060ff1680610a59575033155b1515610a6457600080fd5b50600160a060020a0383166000908152600160205260409020548281038111610a8c57600080fd5b600160a060020a0384166000818152600160209081526040918290208685039055815186815291517f66649d0546ffaed7a9e91793ec2fba0941afa9ebed5b599a8031611ad911fd2f9281900390910190a25060019392505050565b600160a060020a031660009081526005602052604090205490565b6000806000806000806000600860009054906101000a900460ff161515610b2957600080fd5b60008811610b3657600080fd5b600160a060020a0389166000818152600260209081526040808320543384526003835281842094845293825280832054600190925282205492985096509094508511610b8157600080fd5b8786038611610b8f57600080fd5b8785038511610b9d57600080fd5b600160a060020a03891660008181526002602090815260408083208c8b0390553383526003825280832093835292905220888603905587841115610bfd57600160a060020a03891660009081526001602052604090208885039055610c17565b600160a060020a0389166000908152600160205260408120555b336000908152600560205260409020805490935091508282610c3c8260018301610e9a565b81548110610c4657fe5b60009182526020918290206002919091020180546fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff8b8116919091178116700100000000000000000000000000000000439283160217825560018201805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a038e16908117909155604080518d81529485019290925283820186905290519193509133917f3aeb15af61588a39bcfafb19ed853140d195c2a924537afbf9a6d04348e76a69916060908290030190a350600198975050505050505050565b60095481565b60085460009081908190819060ff161515610d4457600080fd5b600954341015610d5357600080fd5b505050600160a060020a03821660008181526002602090815260408083205433845260038352818420948452938252808320546001909252909120543483018310610d9d57600080fd5b3482018210610dab57600080fd5b3481018110610db957600080fd5b600160a060020a038516600081815260026020908152604080832034888101909155338452600383528184209484529382528083208685019055600182528083209385019093556004905220541515610e2857600160a060020a03851660009081526004602052604090204390555b604080513481529051600160a060020a0387169133917f8752a472e571a816aea92eec8dae9baf628e840f4929fbcc2d155e6233ff68a79181900360200190a3506001949350505050565b60408051608081018252600080825260208201819052918101829052606081019190915290565b815481835581811115610ec657600202816002028360005260206000209182019101610ec69190610ecb565b505050565b610f0991905b80821115610f05576000815560018101805474ffffffffffffffffffffffffffffffffffffffffff19169055600201610ed1565b5090565b905600a165627a7a723058207756db8d0b66ca08e1c2d55388ff88403e3a0156e622cdf21e7c37b165814dec0029`

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

// SetDelegator is a paid mutator transaction binding the contract method 0x83cd9cc3.
//
// Solidity: function setDelegator(delegator address) returns(bool)
func (_Stamina *StaminaTransactor) SetDelegator(opts *bind.TransactOpts, delegator common.Address) (*types.Transaction, error) {
	return _Stamina.contract.Transact(opts, "setDelegator", delegator)
}

// SetDelegator is a paid mutator transaction binding the contract method 0x83cd9cc3.
//
// Solidity: function setDelegator(delegator address) returns(bool)
func (_Stamina *StaminaSession) SetDelegator(delegator common.Address) (*types.Transaction, error) {
	return _Stamina.Contract.SetDelegator(&_Stamina.TransactOpts, delegator)
}

// SetDelegator is a paid mutator transaction binding the contract method 0x83cd9cc3.
//
// Solidity: function setDelegator(delegator address) returns(bool)
func (_Stamina *StaminaTransactorSession) SetDelegator(delegator common.Address) (*types.Transaction, error) {
	return _Stamina.Contract.SetDelegator(&_Stamina.TransactOpts, delegator)
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
// Solidity: e DelegateeChanged(delegator indexed address, oldDelegatee address, newDelegatee address)
func (_Stamina *StaminaFilterer) FilterDelegateeChanged(opts *bind.FilterOpts, delegator []common.Address) (*StaminaDelegateeChangedIterator, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}

	logs, sub, err := _Stamina.contract.FilterLogs(opts, "DelegateeChanged", delegatorRule)
	if err != nil {
		return nil, err
	}
	return &StaminaDelegateeChangedIterator{contract: _Stamina.contract, event: "DelegateeChanged", logs: logs, sub: sub}, nil
}

// WatchDelegateeChanged is a free log subscription operation binding the contract event 0x5884d7e3ec123de8e772bcf576c18dcdad75b056c4314f999ed966693419c692.
//
// Solidity: e DelegateeChanged(delegator indexed address, oldDelegatee address, newDelegatee address)
func (_Stamina *StaminaFilterer) WatchDelegateeChanged(opts *bind.WatchOpts, sink chan<- *StaminaDelegateeChanged, delegator []common.Address) (event.Subscription, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}

	logs, sub, err := _Stamina.contract.WatchLogs(opts, "DelegateeChanged", delegatorRule)
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

// StaminaStaminaAddedIterator is returned from FilterStaminaAdded and is used to iterate over the raw logs and unpacked data for StaminaAdded events raised by the Stamina contract.
type StaminaStaminaAddedIterator struct {
	Event *StaminaStaminaAdded // Event containing the contract specifics and raw log

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
func (it *StaminaStaminaAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StaminaStaminaAdded)
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
		it.Event = new(StaminaStaminaAdded)
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
func (it *StaminaStaminaAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StaminaStaminaAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StaminaStaminaAdded represents a StaminaAdded event raised by the Stamina contract.
type StaminaStaminaAdded struct {
	Delegatee common.Address
	Amount    *big.Int
	Recovered bool
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterStaminaAdded is a free log retrieval operation binding the contract event 0x85bf8701ef98ea32e97f08708da81c7daa93e87ea3e2fd661801cce6d36f6809.
//
// Solidity: e StaminaAdded(delegatee indexed address, amount uint256, recovered bool)
func (_Stamina *StaminaFilterer) FilterStaminaAdded(opts *bind.FilterOpts, delegatee []common.Address) (*StaminaStaminaAddedIterator, error) {

	var delegateeRule []interface{}
	for _, delegateeItem := range delegatee {
		delegateeRule = append(delegateeRule, delegateeItem)
	}

	logs, sub, err := _Stamina.contract.FilterLogs(opts, "StaminaAdded", delegateeRule)
	if err != nil {
		return nil, err
	}
	return &StaminaStaminaAddedIterator{contract: _Stamina.contract, event: "StaminaAdded", logs: logs, sub: sub}, nil
}

// WatchStaminaAdded is a free log subscription operation binding the contract event 0x85bf8701ef98ea32e97f08708da81c7daa93e87ea3e2fd661801cce6d36f6809.
//
// Solidity: e StaminaAdded(delegatee indexed address, amount uint256, recovered bool)
func (_Stamina *StaminaFilterer) WatchStaminaAdded(opts *bind.WatchOpts, sink chan<- *StaminaStaminaAdded, delegatee []common.Address) (event.Subscription, error) {

	var delegateeRule []interface{}
	for _, delegateeItem := range delegatee {
		delegateeRule = append(delegateeRule, delegateeItem)
	}

	logs, sub, err := _Stamina.contract.WatchLogs(opts, "StaminaAdded", delegateeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StaminaStaminaAdded)
				if err := _Stamina.contract.UnpackLog(event, "StaminaAdded", log); err != nil {
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

// StaminaStaminaSubtractedIterator is returned from FilterStaminaSubtracted and is used to iterate over the raw logs and unpacked data for StaminaSubtracted events raised by the Stamina contract.
type StaminaStaminaSubtractedIterator struct {
	Event *StaminaStaminaSubtracted // Event containing the contract specifics and raw log

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
func (it *StaminaStaminaSubtractedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StaminaStaminaSubtracted)
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
		it.Event = new(StaminaStaminaSubtracted)
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
func (it *StaminaStaminaSubtractedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StaminaStaminaSubtractedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StaminaStaminaSubtracted represents a StaminaSubtracted event raised by the Stamina contract.
type StaminaStaminaSubtracted struct {
	Delegatee common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterStaminaSubtracted is a free log retrieval operation binding the contract event 0x66649d0546ffaed7a9e91793ec2fba0941afa9ebed5b599a8031611ad911fd2f.
//
// Solidity: e StaminaSubtracted(delegatee indexed address, amount uint256)
func (_Stamina *StaminaFilterer) FilterStaminaSubtracted(opts *bind.FilterOpts, delegatee []common.Address) (*StaminaStaminaSubtractedIterator, error) {

	var delegateeRule []interface{}
	for _, delegateeItem := range delegatee {
		delegateeRule = append(delegateeRule, delegateeItem)
	}

	logs, sub, err := _Stamina.contract.FilterLogs(opts, "StaminaSubtracted", delegateeRule)
	if err != nil {
		return nil, err
	}
	return &StaminaStaminaSubtractedIterator{contract: _Stamina.contract, event: "StaminaSubtracted", logs: logs, sub: sub}, nil
}

// WatchStaminaSubtracted is a free log subscription operation binding the contract event 0x66649d0546ffaed7a9e91793ec2fba0941afa9ebed5b599a8031611ad911fd2f.
//
// Solidity: e StaminaSubtracted(delegatee indexed address, amount uint256)
func (_Stamina *StaminaFilterer) WatchStaminaSubtracted(opts *bind.WatchOpts, sink chan<- *StaminaStaminaSubtracted, delegatee []common.Address) (event.Subscription, error) {

	var delegateeRule []interface{}
	for _, delegateeItem := range delegatee {
		delegateeRule = append(delegateeRule, delegateeItem)
	}

	logs, sub, err := _Stamina.contract.WatchLogs(opts, "StaminaSubtracted", delegateeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StaminaStaminaSubtracted)
				if err := _Stamina.contract.UnpackLog(event, "StaminaSubtracted", log); err != nil {
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
	Depositor       common.Address
	Delegatee       common.Address
	Amount          *big.Int
	WithdrawalIndex *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0x91fb9d98b786c57d74c099ccd2beca1739e9f6a81fb49001ca465c4b7591bbe2.
//
// Solidity: e Withdrawn(depositor indexed address, delegatee indexed address, amount uint256, withdrawalIndex uint256)
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

// WatchWithdrawn is a free log subscription operation binding the contract event 0x91fb9d98b786c57d74c099ccd2beca1739e9f6a81fb49001ca465c4b7591bbe2.
//
// Solidity: e Withdrawn(depositor indexed address, delegatee indexed address, amount uint256, withdrawalIndex uint256)
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
