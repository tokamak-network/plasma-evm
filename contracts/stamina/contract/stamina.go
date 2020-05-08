// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

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

// StaminaABI is the input ABI used to generate the binding from.
const StaminaABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldDelegatee\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newDelegatee\",\"type\":\"address\"}],\"name\":\"DelegateeChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegatee\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegatee\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"recovered\",\"type\":\"bool\"}],\"name\":\"StaminaAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegatee\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"StaminaSubtracted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegatee\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requestBlockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"withdrawalIndex\",\"type\":\"uint256\"}],\"name\":\"WithdrawalRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegatee\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"withdrawalIndex\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"MIN_DEPOSIT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"RECOVER_EPOCH_LENGTH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"WITHDRAWAL_DELAY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegatee\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"addStamina\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegatee\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"development\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"}],\"name\":\"getDelegatee\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"delegatee\",\"type\":\"address\"}],\"name\":\"getDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"}],\"name\":\"getLastProcessedWithdrawalIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegatee\",\"type\":\"address\"}],\"name\":\"getLastRecoveryBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegatee\",\"type\":\"address\"}],\"name\":\"getNumRecovery\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"}],\"name\":\"getNumWithdrawals\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getStamina\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegatee\",\"type\":\"address\"}],\"name\":\"getTotalDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"withdrawalIndex\",\"type\":\"uint256\"}],\"name\":\"getWithdrawal\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"requestBlockNumber\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"delegatee\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"processed\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minDeposit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"recoveryEpochLength\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"withdrawalDelay\",\"type\":\"uint256\"}],\"name\":\"init\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"initialized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegatee\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"requestWithdrawal\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"}],\"name\":\"setDelegator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegatee\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"subtractStamina\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// StaminaFuncSigs maps the 4-byte function signature to its string representation.
var StaminaFuncSigs = map[string]string{
	"e1e158a5": "MIN_DEPOSIT()",
	"1556d8ac": "RECOVER_EPOCH_LENGTH()",
	"0ebb172a": "WITHDRAWAL_DELAY()",
	"bcac9736": "addStamina(address,uint256)",
	"f340fa01": "deposit(address)",
	"7b929c27": "development()",
	"9b4e735f": "getDelegatee(address)",
	"c35082a9": "getDeposit(address,address)",
	"22347f95": "getLastProcessedWithdrawalIndex(address)",
	"937aaef1": "getLastRecoveryBlock(address)",
	"b69ad63b": "getNumRecovery(address)",
	"d898ae1c": "getNumWithdrawals(address)",
	"3900e4ec": "getStamina(address)",
	"857184d1": "getTotalDeposit(address)",
	"5be4f765": "getWithdrawal(address,uint256)",
	"8cd8db8a": "init(uint256,uint256,uint256)",
	"158ef93e": "initialized()",
	"da95ebf7": "requestWithdrawal(address,uint256)",
	"83cd9cc3": "setDelegator(address)",
	"d1c0c042": "subtractStamina(address,uint256)",
	"3ccfd60b": "withdraw()",
}

// StaminaBin is the compiled bytecode used for deploying new contracts.
var StaminaBin = "0x6080604052600c805460ff1916600117905534801561001d57600080fd5b50610fad8061002d6000396000f3fe60806040526004361061012a5760003560e01c80638cd8db8a116100ab578063c35082a91161006f578063c35082a914610425578063d1c0c04214610460578063d898ae1c14610499578063da95ebf7146104cc578063e1e158a514610505578063f340fa011461051a5761012a565b80638cd8db8a146102ff578063937aaef1146103375780639b4e735f1461036a578063b69ad63b146103b9578063bcac9736146103ec5761012a565b80633ccfd60b116100f25780633ccfd60b146101fa5780635be4f7651461020f5780637b929c271461028457806383cd9cc314610299578063857184d1146102cc5761012a565b80630ebb172a1461012f5780631556d8ac14610156578063158ef93e1461016b57806322347f95146101945780633900e4ec146101c7575b600080fd5b34801561013b57600080fd5b50610144610540565b60408051918252519081900360200190f35b34801561016257600080fd5b50610144610546565b34801561017757600080fd5b5061018061054c565b604080519115158252519081900360200190f35b3480156101a057600080fd5b50610144600480360360208110156101b757600080fd5b50356001600160a01b0316610555565b3480156101d357600080fd5b50610144600480360360208110156101ea57600080fd5b50356001600160a01b0316610570565b34801561020657600080fd5b5061018061058b565b34801561021b57600080fd5b506102486004803603604081101561023257600080fd5b506001600160a01b03813516906020013561074c565b604080516001600160801b0395861681529390941660208401526001600160a01b03909116828401521515606082015290519081900360800190f35b34801561029057600080fd5b50610180610808565b3480156102a557600080fd5b50610180600480360360208110156102bc57600080fd5b50356001600160a01b0316610811565b3480156102d857600080fd5b50610144600480360360208110156102ef57600080fd5b50356001600160a01b0316610892565b34801561030b57600080fd5b506103356004803603606081101561032257600080fd5b50803590602081013590604001356108ad565b005b34801561034357600080fd5b506101446004803603602081101561035a57600080fd5b50356001600160a01b031661090e565b34801561037657600080fd5b5061039d6004803603602081101561038d57600080fd5b50356001600160a01b0316610929565b604080516001600160a01b039092168252519081900360200190f35b3480156103c557600080fd5b50610144600480360360208110156103dc57600080fd5b50356001600160a01b0316610947565b3480156103f857600080fd5b506101806004803603604081101561040f57600080fd5b506001600160a01b038135169060200135610962565b34801561043157600080fd5b506101446004803603604081101561044857600080fd5b506001600160a01b0381358116916020013516610aeb565b34801561046c57600080fd5b506101806004803603604081101561048357600080fd5b506001600160a01b038135169060200135610b16565b3480156104a557600080fd5b50610144600480360360208110156104bc57600080fd5b50356001600160a01b0316610bb4565b3480156104d857600080fd5b50610180600480360360408110156104ef57600080fd5b506001600160a01b038135169060200135610bcf565b34801561051157600080fd5b50610144610dae565b6101806004803603602081101561053057600080fd5b50356001600160a01b0316610db4565b600b5481565b600a5481565b60085460ff1681565b6001600160a01b031660009081526006602052604090205490565b6001600160a01b031660009081526001602052604090205490565b33600090815260056020526040812080546105a557600080fd5b3360009081526006602052604081205490811580156105ed5750826000815481106105cc57fe5b906000526020600020906002020160010160149054906101000a900460ff16155b156105fa5750600061061c565b816106165782546002111561060e57600080fd5b50600161061c565b50600181015b8254811061062957600080fd5b33600090815260056020526040812080548390811061064457fe5b906000526020600020906002020190508060010160149054906101000a900460ff161561067057600080fd5b600b54815443600160801b9091046001600160801b0316909101111561069557600080fd5b805460018201805460ff60a01b1916600160a01b17905533600081815260066020526040808220869055516001600160801b039093169283156108fc0291849190818181858888f193505050501580156106f3573d6000803e3d6000fd5b506001820154604080518381526020810186905281516001600160a01b039093169233927f91fb9d98b786c57d74c099ccd2beca1739e9f6a81fb49001ca465c4b7591bbe2928290030190a36001955050505050505b90565b60008060008061075b86610bb4565b851061076657600080fd5b61076e610ef0565b6001600160a01b038716600090815260056020526040902080548790811061079257fe5b600091825260209182902060408051608081018252600290930290910180546001600160801b03808216808652600160801b909204169484018590526001909101546001600160a01b03811692840183905260ff600160a01b90910416151560609093018390529a929950975095509350505050565b600c5460ff1681565b60085460009060ff1661082357600080fd5b6001600160a01b03828116600081815260208181526040918290208054336001600160a01b0319821681179092558351951680865291850152815190937f5884d7e3ec123de8e772bcf576c18dcdad75b056c4314f999ed966693419c69292908290030190a250600192915050565b6001600160a01b031660009081526002602052604090205490565b60085460ff16156108bd57600080fd5b600083116108ca57600080fd5b600082116108d757600080fd5b600081116108e457600080fd5b8082600202106108f357600080fd5b600992909255600a55600b556008805460ff19166001179055565b6001600160a01b031660009081526004602052604090205490565b6001600160a01b039081166000908152602081905260409020541690565b6001600160a01b031660009081526007602052604090205490565b600c5460009060ff1680610974575033155b61097d57600080fd5b600a546001600160a01b03841660009081526004602052604090205443910111610a20576001600160a01b038316600081815260026020908152604080832054600180845282852091909155600483528184204390556007835281842080548201905581519384529183019190915280517f85bf8701ef98ea32e97f08708da81c7daa93e87ea3e2fd661801cce6d36f68099281900390910190a2506001610ae5565b6001600160a01b0383166000908152600260209081526040808320546001909252909120548381018110610a5357600080fd5b80840182811115610a7e576001600160a01b0386166000908152600160205260409020839055610a9a565b6001600160a01b03861660009081526001602052604090208190555b604080518681526000602082015281516001600160a01b038916927f85bf8701ef98ea32e97f08708da81c7daa93e87ea3e2fd661801cce6d36f6809928290030190a2600193505050505b92915050565b6001600160a01b03918216600090815260036020908152604080832093909416825291909152205490565b600c5460009060ff1680610b28575033155b610b3157600080fd5b6001600160a01b0383166000908152600160205260409020548281038111610b5857600080fd5b6001600160a01b0384166000818152600160209081526040918290208685039055815186815291517f66649d0546ffaed7a9e91793ec2fba0941afa9ebed5b599a8031611ad911fd2f9281900390910190a25060019392505050565b6001600160a01b031660009081526005602052604090205490565b60085460009060ff16610be157600080fd5b60008211610bee57600080fd5b6001600160a01b038316600081815260026020908152604080832054338452600383528184209484529382528083205460019092529091205481610c3157600080fd5b8285840310610c3f57600080fd5b8185830310610c4d57600080fd5b6001600160a01b038616600081815260026020908152604080832089880390553383526003825280832093835292905220858303905584811115610cad576001600160a01b03861660009081526001602052604090208582039055610cc7565b6001600160a01b0386166000908152600160205260408120555b336000908152600560205260408120805490918282610ce98260018301610f17565b81548110610cf357fe5b60009182526020918290206002919091020180546fffffffffffffffffffffffffffffffff19166001600160801b038b8116919091178116600160801b43928316021782556001820180546001600160a01b0319166001600160a01b038e16908117909155604080518d81529485019290925283820186905290519193509133917f3aeb15af61588a39bcfafb19ed853140d195c2a924537afbf9a6d04348e76a69916060908290030190a350600198975050505050505050565b60095481565b60085460009060ff16610dc657600080fd5b600954341015610dd557600080fd5b6001600160a01b03821660008181526002602090815260408083205433845260038352818420948452938252808320546001909252909120543483018310610e1c57600080fd5b8134830111610e2a57600080fd5b8034820111610e3857600080fd5b6001600160a01b03851660008181526002602090815260408083203488810190915533845260038352818420948452938252808320868501905560018252808320938501909355600490522054610ea5576001600160a01b03851660009081526004602052604090204390555b6040805134815290516001600160a01b0387169133917f8752a472e571a816aea92eec8dae9baf628e840f4929fbcc2d155e6233ff68a79181900360200190a3506001949350505050565b60408051608081018252600080825260208201819052918101829052606081019190915290565b815481835581811115610f4357600202816002028360005260206000209182019101610f439190610f48565b505050565b61074991905b80821115610f7457600081556001810180546001600160a81b0319169055600201610f4e565b509056fea265627a7a72315820c78285bb85a38fbf4c44937f68e217e65725afcc6a0a0652f35e8a64e7cee11a64736f6c634300050c0032"

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
// Solidity: function getDelegatee(address delegator) constant returns(address)
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
// Solidity: function getDelegatee(address delegator) constant returns(address)
func (_Stamina *StaminaSession) GetDelegatee(delegator common.Address) (common.Address, error) {
	return _Stamina.Contract.GetDelegatee(&_Stamina.CallOpts, delegator)
}

// GetDelegatee is a free data retrieval call binding the contract method 0x9b4e735f.
//
// Solidity: function getDelegatee(address delegator) constant returns(address)
func (_Stamina *StaminaCallerSession) GetDelegatee(delegator common.Address) (common.Address, error) {
	return _Stamina.Contract.GetDelegatee(&_Stamina.CallOpts, delegator)
}

// GetDeposit is a free data retrieval call binding the contract method 0xc35082a9.
//
// Solidity: function getDeposit(address depositor, address delegatee) constant returns(uint256)
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
// Solidity: function getDeposit(address depositor, address delegatee) constant returns(uint256)
func (_Stamina *StaminaSession) GetDeposit(depositor common.Address, delegatee common.Address) (*big.Int, error) {
	return _Stamina.Contract.GetDeposit(&_Stamina.CallOpts, depositor, delegatee)
}

// GetDeposit is a free data retrieval call binding the contract method 0xc35082a9.
//
// Solidity: function getDeposit(address depositor, address delegatee) constant returns(uint256)
func (_Stamina *StaminaCallerSession) GetDeposit(depositor common.Address, delegatee common.Address) (*big.Int, error) {
	return _Stamina.Contract.GetDeposit(&_Stamina.CallOpts, depositor, delegatee)
}

// GetLastProcessedWithdrawalIndex is a free data retrieval call binding the contract method 0x22347f95.
//
// Solidity: function getLastProcessedWithdrawalIndex(address depositor) constant returns(uint256)
func (_Stamina *StaminaCaller) GetLastProcessedWithdrawalIndex(opts *bind.CallOpts, depositor common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Stamina.contract.Call(opts, out, "getLastProcessedWithdrawalIndex", depositor)
	return *ret0, err
}

// GetLastProcessedWithdrawalIndex is a free data retrieval call binding the contract method 0x22347f95.
//
// Solidity: function getLastProcessedWithdrawalIndex(address depositor) constant returns(uint256)
func (_Stamina *StaminaSession) GetLastProcessedWithdrawalIndex(depositor common.Address) (*big.Int, error) {
	return _Stamina.Contract.GetLastProcessedWithdrawalIndex(&_Stamina.CallOpts, depositor)
}

// GetLastProcessedWithdrawalIndex is a free data retrieval call binding the contract method 0x22347f95.
//
// Solidity: function getLastProcessedWithdrawalIndex(address depositor) constant returns(uint256)
func (_Stamina *StaminaCallerSession) GetLastProcessedWithdrawalIndex(depositor common.Address) (*big.Int, error) {
	return _Stamina.Contract.GetLastProcessedWithdrawalIndex(&_Stamina.CallOpts, depositor)
}

// GetLastRecoveryBlock is a free data retrieval call binding the contract method 0x937aaef1.
//
// Solidity: function getLastRecoveryBlock(address delegatee) constant returns(uint256)
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
// Solidity: function getLastRecoveryBlock(address delegatee) constant returns(uint256)
func (_Stamina *StaminaSession) GetLastRecoveryBlock(delegatee common.Address) (*big.Int, error) {
	return _Stamina.Contract.GetLastRecoveryBlock(&_Stamina.CallOpts, delegatee)
}

// GetLastRecoveryBlock is a free data retrieval call binding the contract method 0x937aaef1.
//
// Solidity: function getLastRecoveryBlock(address delegatee) constant returns(uint256)
func (_Stamina *StaminaCallerSession) GetLastRecoveryBlock(delegatee common.Address) (*big.Int, error) {
	return _Stamina.Contract.GetLastRecoveryBlock(&_Stamina.CallOpts, delegatee)
}

// GetNumRecovery is a free data retrieval call binding the contract method 0xb69ad63b.
//
// Solidity: function getNumRecovery(address delegatee) constant returns(uint256)
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
// Solidity: function getNumRecovery(address delegatee) constant returns(uint256)
func (_Stamina *StaminaSession) GetNumRecovery(delegatee common.Address) (*big.Int, error) {
	return _Stamina.Contract.GetNumRecovery(&_Stamina.CallOpts, delegatee)
}

// GetNumRecovery is a free data retrieval call binding the contract method 0xb69ad63b.
//
// Solidity: function getNumRecovery(address delegatee) constant returns(uint256)
func (_Stamina *StaminaCallerSession) GetNumRecovery(delegatee common.Address) (*big.Int, error) {
	return _Stamina.Contract.GetNumRecovery(&_Stamina.CallOpts, delegatee)
}

// GetNumWithdrawals is a free data retrieval call binding the contract method 0xd898ae1c.
//
// Solidity: function getNumWithdrawals(address depositor) constant returns(uint256)
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
// Solidity: function getNumWithdrawals(address depositor) constant returns(uint256)
func (_Stamina *StaminaSession) GetNumWithdrawals(depositor common.Address) (*big.Int, error) {
	return _Stamina.Contract.GetNumWithdrawals(&_Stamina.CallOpts, depositor)
}

// GetNumWithdrawals is a free data retrieval call binding the contract method 0xd898ae1c.
//
// Solidity: function getNumWithdrawals(address depositor) constant returns(uint256)
func (_Stamina *StaminaCallerSession) GetNumWithdrawals(depositor common.Address) (*big.Int, error) {
	return _Stamina.Contract.GetNumWithdrawals(&_Stamina.CallOpts, depositor)
}

// GetStamina is a free data retrieval call binding the contract method 0x3900e4ec.
//
// Solidity: function getStamina(address addr) constant returns(uint256)
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
// Solidity: function getStamina(address addr) constant returns(uint256)
func (_Stamina *StaminaSession) GetStamina(addr common.Address) (*big.Int, error) {
	return _Stamina.Contract.GetStamina(&_Stamina.CallOpts, addr)
}

// GetStamina is a free data retrieval call binding the contract method 0x3900e4ec.
//
// Solidity: function getStamina(address addr) constant returns(uint256)
func (_Stamina *StaminaCallerSession) GetStamina(addr common.Address) (*big.Int, error) {
	return _Stamina.Contract.GetStamina(&_Stamina.CallOpts, addr)
}

// GetTotalDeposit is a free data retrieval call binding the contract method 0x857184d1.
//
// Solidity: function getTotalDeposit(address delegatee) constant returns(uint256)
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
// Solidity: function getTotalDeposit(address delegatee) constant returns(uint256)
func (_Stamina *StaminaSession) GetTotalDeposit(delegatee common.Address) (*big.Int, error) {
	return _Stamina.Contract.GetTotalDeposit(&_Stamina.CallOpts, delegatee)
}

// GetTotalDeposit is a free data retrieval call binding the contract method 0x857184d1.
//
// Solidity: function getTotalDeposit(address delegatee) constant returns(uint256)
func (_Stamina *StaminaCallerSession) GetTotalDeposit(delegatee common.Address) (*big.Int, error) {
	return _Stamina.Contract.GetTotalDeposit(&_Stamina.CallOpts, delegatee)
}

// GetWithdrawal is a free data retrieval call binding the contract method 0x5be4f765.
//
// Solidity: function getWithdrawal(address depositor, uint256 withdrawalIndex) constant returns(uint128 amount, uint128 requestBlockNumber, address delegatee, bool processed)
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
// Solidity: function getWithdrawal(address depositor, uint256 withdrawalIndex) constant returns(uint128 amount, uint128 requestBlockNumber, address delegatee, bool processed)
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
// Solidity: function getWithdrawal(address depositor, uint256 withdrawalIndex) constant returns(uint128 amount, uint128 requestBlockNumber, address delegatee, bool processed)
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
// Solidity: function addStamina(address delegatee, uint256 amount) returns(bool)
func (_Stamina *StaminaTransactor) AddStamina(opts *bind.TransactOpts, delegatee common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Stamina.contract.Transact(opts, "addStamina", delegatee, amount)
}

// AddStamina is a paid mutator transaction binding the contract method 0xbcac9736.
//
// Solidity: function addStamina(address delegatee, uint256 amount) returns(bool)
func (_Stamina *StaminaSession) AddStamina(delegatee common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Stamina.Contract.AddStamina(&_Stamina.TransactOpts, delegatee, amount)
}

// AddStamina is a paid mutator transaction binding the contract method 0xbcac9736.
//
// Solidity: function addStamina(address delegatee, uint256 amount) returns(bool)
func (_Stamina *StaminaTransactorSession) AddStamina(delegatee common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Stamina.Contract.AddStamina(&_Stamina.TransactOpts, delegatee, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xf340fa01.
//
// Solidity: function deposit(address delegatee) returns(bool)
func (_Stamina *StaminaTransactor) Deposit(opts *bind.TransactOpts, delegatee common.Address) (*types.Transaction, error) {
	return _Stamina.contract.Transact(opts, "deposit", delegatee)
}

// Deposit is a paid mutator transaction binding the contract method 0xf340fa01.
//
// Solidity: function deposit(address delegatee) returns(bool)
func (_Stamina *StaminaSession) Deposit(delegatee common.Address) (*types.Transaction, error) {
	return _Stamina.Contract.Deposit(&_Stamina.TransactOpts, delegatee)
}

// Deposit is a paid mutator transaction binding the contract method 0xf340fa01.
//
// Solidity: function deposit(address delegatee) returns(bool)
func (_Stamina *StaminaTransactorSession) Deposit(delegatee common.Address) (*types.Transaction, error) {
	return _Stamina.Contract.Deposit(&_Stamina.TransactOpts, delegatee)
}

// Init is a paid mutator transaction binding the contract method 0x8cd8db8a.
//
// Solidity: function init(uint256 minDeposit, uint256 recoveryEpochLength, uint256 withdrawalDelay) returns()
func (_Stamina *StaminaTransactor) Init(opts *bind.TransactOpts, minDeposit *big.Int, recoveryEpochLength *big.Int, withdrawalDelay *big.Int) (*types.Transaction, error) {
	return _Stamina.contract.Transact(opts, "init", minDeposit, recoveryEpochLength, withdrawalDelay)
}

// Init is a paid mutator transaction binding the contract method 0x8cd8db8a.
//
// Solidity: function init(uint256 minDeposit, uint256 recoveryEpochLength, uint256 withdrawalDelay) returns()
func (_Stamina *StaminaSession) Init(minDeposit *big.Int, recoveryEpochLength *big.Int, withdrawalDelay *big.Int) (*types.Transaction, error) {
	return _Stamina.Contract.Init(&_Stamina.TransactOpts, minDeposit, recoveryEpochLength, withdrawalDelay)
}

// Init is a paid mutator transaction binding the contract method 0x8cd8db8a.
//
// Solidity: function init(uint256 minDeposit, uint256 recoveryEpochLength, uint256 withdrawalDelay) returns()
func (_Stamina *StaminaTransactorSession) Init(minDeposit *big.Int, recoveryEpochLength *big.Int, withdrawalDelay *big.Int) (*types.Transaction, error) {
	return _Stamina.Contract.Init(&_Stamina.TransactOpts, minDeposit, recoveryEpochLength, withdrawalDelay)
}

// RequestWithdrawal is a paid mutator transaction binding the contract method 0xda95ebf7.
//
// Solidity: function requestWithdrawal(address delegatee, uint256 amount) returns(bool)
func (_Stamina *StaminaTransactor) RequestWithdrawal(opts *bind.TransactOpts, delegatee common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Stamina.contract.Transact(opts, "requestWithdrawal", delegatee, amount)
}

// RequestWithdrawal is a paid mutator transaction binding the contract method 0xda95ebf7.
//
// Solidity: function requestWithdrawal(address delegatee, uint256 amount) returns(bool)
func (_Stamina *StaminaSession) RequestWithdrawal(delegatee common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Stamina.Contract.RequestWithdrawal(&_Stamina.TransactOpts, delegatee, amount)
}

// RequestWithdrawal is a paid mutator transaction binding the contract method 0xda95ebf7.
//
// Solidity: function requestWithdrawal(address delegatee, uint256 amount) returns(bool)
func (_Stamina *StaminaTransactorSession) RequestWithdrawal(delegatee common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Stamina.Contract.RequestWithdrawal(&_Stamina.TransactOpts, delegatee, amount)
}

// SetDelegator is a paid mutator transaction binding the contract method 0x83cd9cc3.
//
// Solidity: function setDelegator(address delegator) returns(bool)
func (_Stamina *StaminaTransactor) SetDelegator(opts *bind.TransactOpts, delegator common.Address) (*types.Transaction, error) {
	return _Stamina.contract.Transact(opts, "setDelegator", delegator)
}

// SetDelegator is a paid mutator transaction binding the contract method 0x83cd9cc3.
//
// Solidity: function setDelegator(address delegator) returns(bool)
func (_Stamina *StaminaSession) SetDelegator(delegator common.Address) (*types.Transaction, error) {
	return _Stamina.Contract.SetDelegator(&_Stamina.TransactOpts, delegator)
}

// SetDelegator is a paid mutator transaction binding the contract method 0x83cd9cc3.
//
// Solidity: function setDelegator(address delegator) returns(bool)
func (_Stamina *StaminaTransactorSession) SetDelegator(delegator common.Address) (*types.Transaction, error) {
	return _Stamina.Contract.SetDelegator(&_Stamina.TransactOpts, delegator)
}

// SubtractStamina is a paid mutator transaction binding the contract method 0xd1c0c042.
//
// Solidity: function subtractStamina(address delegatee, uint256 amount) returns(bool)
func (_Stamina *StaminaTransactor) SubtractStamina(opts *bind.TransactOpts, delegatee common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Stamina.contract.Transact(opts, "subtractStamina", delegatee, amount)
}

// SubtractStamina is a paid mutator transaction binding the contract method 0xd1c0c042.
//
// Solidity: function subtractStamina(address delegatee, uint256 amount) returns(bool)
func (_Stamina *StaminaSession) SubtractStamina(delegatee common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Stamina.Contract.SubtractStamina(&_Stamina.TransactOpts, delegatee, amount)
}

// SubtractStamina is a paid mutator transaction binding the contract method 0xd1c0c042.
//
// Solidity: function subtractStamina(address delegatee, uint256 amount) returns(bool)
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
// Solidity: event DelegateeChanged(address indexed delegator, address oldDelegatee, address newDelegatee)
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
// Solidity: event DelegateeChanged(address indexed delegator, address oldDelegatee, address newDelegatee)
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

// ParseDelegateeChanged is a log parse operation binding the contract event 0x5884d7e3ec123de8e772bcf576c18dcdad75b056c4314f999ed966693419c692.
//
// Solidity: event DelegateeChanged(address indexed delegator, address oldDelegatee, address newDelegatee)
func (_Stamina *StaminaFilterer) ParseDelegateeChanged(log types.Log) (*StaminaDelegateeChanged, error) {
	event := new(StaminaDelegateeChanged)
	if err := _Stamina.contract.UnpackLog(event, "DelegateeChanged", log); err != nil {
		return nil, err
	}
	return event, nil
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
// Solidity: event Deposited(address indexed depositor, address indexed delegatee, uint256 amount)
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
// Solidity: event Deposited(address indexed depositor, address indexed delegatee, uint256 amount)
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

// ParseDeposited is a log parse operation binding the contract event 0x8752a472e571a816aea92eec8dae9baf628e840f4929fbcc2d155e6233ff68a7.
//
// Solidity: event Deposited(address indexed depositor, address indexed delegatee, uint256 amount)
func (_Stamina *StaminaFilterer) ParseDeposited(log types.Log) (*StaminaDeposited, error) {
	event := new(StaminaDeposited)
	if err := _Stamina.contract.UnpackLog(event, "Deposited", log); err != nil {
		return nil, err
	}
	return event, nil
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
// Solidity: event StaminaAdded(address indexed delegatee, uint256 amount, bool recovered)
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
// Solidity: event StaminaAdded(address indexed delegatee, uint256 amount, bool recovered)
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

// ParseStaminaAdded is a log parse operation binding the contract event 0x85bf8701ef98ea32e97f08708da81c7daa93e87ea3e2fd661801cce6d36f6809.
//
// Solidity: event StaminaAdded(address indexed delegatee, uint256 amount, bool recovered)
func (_Stamina *StaminaFilterer) ParseStaminaAdded(log types.Log) (*StaminaStaminaAdded, error) {
	event := new(StaminaStaminaAdded)
	if err := _Stamina.contract.UnpackLog(event, "StaminaAdded", log); err != nil {
		return nil, err
	}
	return event, nil
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
// Solidity: event StaminaSubtracted(address indexed delegatee, uint256 amount)
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
// Solidity: event StaminaSubtracted(address indexed delegatee, uint256 amount)
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

// ParseStaminaSubtracted is a log parse operation binding the contract event 0x66649d0546ffaed7a9e91793ec2fba0941afa9ebed5b599a8031611ad911fd2f.
//
// Solidity: event StaminaSubtracted(address indexed delegatee, uint256 amount)
func (_Stamina *StaminaFilterer) ParseStaminaSubtracted(log types.Log) (*StaminaStaminaSubtracted, error) {
	event := new(StaminaStaminaSubtracted)
	if err := _Stamina.contract.UnpackLog(event, "StaminaSubtracted", log); err != nil {
		return nil, err
	}
	return event, nil
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
// Solidity: event WithdrawalRequested(address indexed depositor, address indexed delegatee, uint256 amount, uint256 requestBlockNumber, uint256 withdrawalIndex)
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
// Solidity: event WithdrawalRequested(address indexed depositor, address indexed delegatee, uint256 amount, uint256 requestBlockNumber, uint256 withdrawalIndex)
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

// ParseWithdrawalRequested is a log parse operation binding the contract event 0x3aeb15af61588a39bcfafb19ed853140d195c2a924537afbf9a6d04348e76a69.
//
// Solidity: event WithdrawalRequested(address indexed depositor, address indexed delegatee, uint256 amount, uint256 requestBlockNumber, uint256 withdrawalIndex)
func (_Stamina *StaminaFilterer) ParseWithdrawalRequested(log types.Log) (*StaminaWithdrawalRequested, error) {
	event := new(StaminaWithdrawalRequested)
	if err := _Stamina.contract.UnpackLog(event, "WithdrawalRequested", log); err != nil {
		return nil, err
	}
	return event, nil
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
// Solidity: event Withdrawn(address indexed depositor, address indexed delegatee, uint256 amount, uint256 withdrawalIndex)
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
// Solidity: event Withdrawn(address indexed depositor, address indexed delegatee, uint256 amount, uint256 withdrawalIndex)
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

// ParseWithdrawn is a log parse operation binding the contract event 0x91fb9d98b786c57d74c099ccd2beca1739e9f6a81fb49001ca465c4b7591bbe2.
//
// Solidity: event Withdrawn(address indexed depositor, address indexed delegatee, uint256 amount, uint256 withdrawalIndex)
func (_Stamina *StaminaFilterer) ParseWithdrawn(log types.Log) (*StaminaWithdrawn, error) {
	event := new(StaminaWithdrawn)
	if err := _Stamina.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	return event, nil
}
