package core

import (
	"errors"
	"math/big"

	"github.com/Onther-Tech/plasma-evm/common"
	"github.com/Onther-Tech/plasma-evm/core/vm"
	"github.com/Onther-Tech/plasma-evm/params"
)

var (
	errUpdateStamina = errors.New("failed to update stamina")
)

func GetDelegatee(evm *vm.EVM, from common.Address) (common.Address, error) {
	data, err := params.StaminaABI.Pack("getDelegatee", from)
	if err != nil {
		return common.Address{}, err
	}

	ret, _, err := evm.StaticCall(params.BlockchainAccount, params.StaminaAddress, data, 1000000)

	if err != nil {
		return common.Address{}, err
	}

	return common.BytesToAddress(ret), nil
}

func GetStamina(evm *vm.EVM, delegatee common.Address) (*big.Int, error) {
	data, err := params.StaminaABI.Pack("getStamina", delegatee)
	if err != nil {
		return big.NewInt(0), err
	}

	ret, _, err := evm.StaticCall(params.BlockchainAccount, params.StaminaAddress, data, 1000000)

	if err != nil {
		return big.NewInt(0), err
	}

	stamina := new(big.Int)
	stamina.SetBytes(ret)

	return stamina, nil
}

func AddStamina(evm *vm.EVM, delegatee common.Address, gas *big.Int) error {
	data, err := params.StaminaABI.Pack("addStamina", delegatee, gas)
	if err != nil {
		return err
	}

	res, _, err := evm.Call(params.BlockchainAccount, params.StaminaAddress, data, 1000000, big.NewInt(0))

	if new(big.Int).SetBytes(res).Cmp(common.Big1) != 0 {
		return errUpdateStamina
	}

	return err
}

func SubtractStamina(evm *vm.EVM, delegatee common.Address, gas *big.Int) error {
	data, err := params.StaminaABI.Pack("subtractStamina", delegatee, gas)
	if err != nil {
		return err
	}

	res, _, err := evm.Call(params.BlockchainAccount, params.StaminaAddress, data, 1000000, big.NewInt(0))

	if new(big.Int).SetBytes(res).Cmp(common.Big1) != 0 {
		return errUpdateStamina
	}

	return err
}
