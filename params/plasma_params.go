package params

import (
	"github.com/Onther-Tech/plasma-evm/common"
	"github.com/Onther-Tech/plasma-evm/crypto"
)

var (
	NullAddress = common.Address{0x0000000000000000000000000000000000000000}
	NullKey, _ = crypto.HexToECDSA("0x00")
)