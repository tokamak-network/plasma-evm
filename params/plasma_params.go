package params

import (
	"github.com/Onther-Tech/plasma-evm/common"
	"github.com/Onther-Tech/plasma-evm/crypto"
)

var (
	NullAddress 	= common.Address{0x0000000000000000000000000000000000000000}
	NullKey, _ 		= crypto.HexToECDSA("0x00")
	NRBepochLength	= int32(4)
	ORBepochLength	= int32(1)
	IsNRB			= true
	IsORB			= false
	NumNRBmined 	= int32(0)
	NumORBmined		= int32(0)
)