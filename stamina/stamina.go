package stamina

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
	staminaCommon "github.com/ethereum/go-ethereum/stamina/common"
)

func GetDelegatee(evm *vm.EVM, from common.Address) (common.Address, error) {
	data, err := staminaCommon.StaminaABI.Pack("getDelegatee", from)
	if err != nil {
		return common.Address{}, err
	}

	ret, _, err := evm.StaticCall(staminaCommon.BlockchainAccount, staminaCommon.StaminaContractAddress, data, 1000000)

	if err != nil {
		return common.Address{}, err
	}

	return common.BytesToAddress(ret), nil
}

func GetStamina(evm *vm.EVM, delegatee common.Address) (*big.Int, error) {
	data, err := staminaCommon.StaminaABI.Pack("getStamina", delegatee)
	if err != nil {
		return big.NewInt(0), err
	}

	ret, _, err := evm.StaticCall(staminaCommon.BlockchainAccount, staminaCommon.StaminaContractAddress, data, 1000000)

	if err != nil {
		return big.NewInt(0), err
	}

	stamina := new(big.Int)
	stamina.SetBytes(ret)

	return stamina, nil
}

func AddStamina(evm *vm.EVM, delegatee common.Address, gas *big.Int) error {
	data, err := staminaCommon.StaminaABI.Pack("addStamina", delegatee, gas)
	if err != nil {
		return err
	}

	_, _, err = evm.Call(staminaCommon.BlockchainAccount, staminaCommon.StaminaContractAddress, data, 1000000, big.NewInt(0))

	return err
}

func SubtractStamina(evm *vm.EVM, delegatee common.Address, gas *big.Int) error {
	data, err := staminaCommon.StaminaABI.Pack("subtractStamina", delegatee, gas)
	if err != nil {
		return err
	}

	_, _, err = evm.Call(staminaCommon.BlockchainAccount, staminaCommon.StaminaContractAddress, data, 1000000, big.NewInt(0))

	return err
}
