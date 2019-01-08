package common

import (
	"strings"

	"github.com/Onther-Tech/plasma-evm/accounts/abi"
	"github.com/Onther-Tech/plasma-evm/common"
	"github.com/Onther-Tech/plasma-evm/contracts/stamina/contract"
)

type accountWrapper struct {
	address common.Address
}

func (a accountWrapper) Address() common.Address {
	return a.address
}

var BlockchainAccount = accountWrapper{common.HexToAddress("0x00")}
var StaminaAccount = accountWrapper{StaminaContractAddress}
var StaminaABI, _ = abi.JSON(strings.NewReader(contract.StaminaABI))
