package stamina

//go:generate abigen --sol contract/Stamina.sol --pkg contract --out contract/stamina.go

import (
	"github.com/Onther-Tech/plasma-evm/accounts/abi/bind"
	"github.com/Onther-Tech/plasma-evm/common"
	"github.com/Onther-Tech/plasma-evm/contracts/stamina/contract"
)

type Stamina struct {
	*contract.StaminaSession
	contractBackend bind.ContractBackend
}

func NewStamina(transactOpts *bind.TransactOpts, contractAddr common.Address, contractBackend bind.ContractBackend) (*Stamina, error) {
	stamina, err := contract.NewStamina(contractAddr, contractBackend)
	if err != nil {
		return nil, err
	}

	return &Stamina{
		&contract.StaminaSession{
			Contract:     stamina,
			TransactOpts: *transactOpts,
		},
		contractBackend,
	}, nil
}
