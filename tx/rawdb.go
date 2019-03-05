package tx

import (
	"encoding/binary"
	"math"
	"math/big"

	"github.com/Onther-Tech/plasma-evm/common"
	"github.com/Onther-Tech/plasma-evm/core/rawdb"
	"github.com/Onther-Tech/plasma-evm/log"
	"github.com/Onther-Tech/plasma-evm/params"
	"github.com/Onther-Tech/plasma-evm/rlp"
)

// TODO: Belows should be move into rawdb package.

const (
	MaxUint64 = math.MaxUint64
)

var (
	gasPriceKey = []byte("gas-price")

	numAddrKey      = []byte("num-address")   // numAddrKey -> the number of accounts
	addrPrefix      = []byte("address")       // addrPrefix + i (uint64 big endian) -> i-th account to send transaction
	addrNoncePrefix = []byte("address-nonce") // addrPrefix + address -> i-th account's nonce

	numRawTxsPrefix        = []byte("num-raw-txs")      // numRawTxsPrefix + account address -> number of raw transactions
	lastPendingIndexPrefix = []byte("last-pending-raw") // lastPendingIndexKey + account address -> index of last pending transaction

	rawTxPrefix = []byte("raw-tx") // rawTxPrefix + account address + i (uint64 big endian) -> i-th raw transaction
)

func ReadGasPrice(db rawdb.DatabaseReader) *big.Int {
	data, _ := db.Get(gasPriceKey)

	if len(data) == 0 {
		return new(big.Int).SetInt64(params.GWei)
	}

	var gasPrice big.Int
	if err := rlp.DecodeBytes(data, &gasPrice); err != nil {
		log.Error("Failed to decode gas price", "err", err)
		return new(big.Int).SetInt64(params.GWei)
	}

	return &gasPrice
}

func WriteGasPrice(db rawdb.DatabaseWriter, gasPrice *big.Int) {
	data, err := rlp.EncodeToBytes(gasPrice)
	if err != nil {
		log.Crit("Failed to encode gas price", "err", err)
	}
	if err := db.Put(gasPriceKey, data); err != nil {
		log.Crit("Failed to store gas price", "err", err)
	}
}

func ReadNumAddr(db rawdb.DatabaseReader) uint64 {
	data, _ := db.Get(numAddrKey)

	if len(data) == 0 {
		return 0
	}

	var n uint64
	if err := rlp.DecodeBytes(data, &n); err != nil {
		log.Error("Failed to decode account number", "err", err)
		return MaxUint64
	}

	return n
}

func WriteNumAddr(db rawdb.DatabaseWriter, n uint64) {
	data, err := rlp.EncodeToBytes(n)
	if err != nil {
		log.Crit("Failed to encode account number", "err", err)
	}
	if err := db.Put(numAddrKey, data); err != nil {
		log.Crit("Failed to store account number", "err", err)
	}
}

func addrKey(i uint64) []byte {
	return append(addrPrefix, encodeNumber(i)...)
}

func ReadAddr(db rawdb.DatabaseReader, i uint64) common.Address {
	data, _ := db.Get(addrKey(i))

	if len(data) == 0 {
		return common.Address{}
	}

	var addr common.Address
	if err := rlp.DecodeBytes(data, &addr); err != nil {
		log.Error("Failed to decode account address", "err", err)
		return common.Address{}
	}

	return addr
}

func WriteAddr(db rawdb.DatabaseWriter, i uint64, addr common.Address) {
	data, err := rlp.EncodeToBytes(addr)
	if err != nil {
		log.Crit("Failed to encode account address", "err", err)
	}
	if err := db.Put(addrKey(i), data); err != nil {
		log.Crit("Failed to store account address", "err", err)
	}
}

func addrNonceKey(addr common.Address) []byte {
	return append(addrNoncePrefix, addr.Bytes()...)
}

func ReadAddrNonce(db rawdb.DatabaseReader, addr common.Address) uint64 {
	data, _ := db.Get(addrNonceKey(addr))

	if len(data) == 0 {
		return 0
	}

	var nonce uint64
	if err := rlp.DecodeBytes(data, &nonce); err != nil {
		log.Error("Failed to decode account nonce", "err", err)
		return 0
	}

	return nonce
}

func WriteAddrNonce(db rawdb.DatabaseWriter, addr common.Address, nonce uint64) {
	data, err := rlp.EncodeToBytes(nonce)
	if err != nil {
		log.Crit("Failed to encode account nonce", "err", err)
	}
	if err := db.Put(addrNonceKey(addr), data); err != nil {
		log.Crit("Failed to store account nonce", "err", err)
	}
}

func numRawTxsKey(addr common.Address) []byte {
	return append(numRawTxsPrefix, addr.Bytes()...)
}

func ReadNumRawTxs(db rawdb.DatabaseReader, addr common.Address) uint64 {
	data, _ := db.Get(numRawTxsKey(addr))

	if len(data) == 0 {
		return 0
	}

	var n uint64
	if err := rlp.DecodeBytes(data, &n); err != nil {
		log.Error("Failed to decode number of raw transactions", "err", err)
		return MaxUint64
	}

	return n
}

func WriteNumRawTxs(db rawdb.DatabaseWriter, addr common.Address, n uint64) {
	data, err := rlp.EncodeToBytes(n)
	if err != nil {
		log.Crit("Failed to encode number of raw transactions", "err", err)
	}
	if err := db.Put(numRawTxsKey(addr), data); err != nil {
		log.Crit("Failed to store number of raw transactions", "err", err)
	}
}

func lastPendingIndexKey(addr common.Address) []byte {
	return append(lastPendingIndexPrefix, addr.Bytes()...)
}

func ReadLastPendingIndex(db rawdb.DatabaseReader, addr common.Address) uint64 {
	data, _ := db.Get(lastPendingIndexKey(addr))

	if len(data) == 0 {
		return MaxUint64
	}

	var i uint64
	if err := rlp.DecodeBytes(data, &i); err != nil {
		log.Error("Failed to decode number of raw transactions", "err", err)
		return MaxUint64
	}

	return i
}

func WriteLastPendingIndex(db rawdb.DatabaseWriter, addr common.Address, i uint64) {
	data, err := rlp.EncodeToBytes(i)
	if err != nil {
		log.Crit("Failed to encode number of raw transactions", "err", err)
	}
	if err := db.Put(lastPendingIndexKey(addr), data); err != nil {
		log.Crit("Failed to store number of raw transactions", "err", err)
	}
}

func rawTxKey(addr common.Address, i uint64) []byte {
	return append(append(rawTxPrefix, addr.Bytes()...), encodeNumber(i)...)
}

func ReadRawTx(db rawdb.DatabaseReader, addr common.Address, i uint64) *RawTransaction {
	data, _ := db.Get(rawTxKey(addr, i))

	if len(data) == 0 {
		return nil
	}

	var raw RawTransaction
	if err := rlp.DecodeBytes(data, &raw); err != nil {
		log.Error("Failed to decode raw transaction", "err", err, "addr", addr, "i", i)
		return nil
	}

	return &raw
}

func WriteRawTx(db rawdb.DatabaseWriter, addr common.Address, raw RawTransaction) {
	data, err := rlp.EncodeToBytes(raw)
	if err != nil {
		log.Crit("Failed to encode raw transaction", "err", err)
	}
	if err := db.Put(rawTxKey(addr, raw.Index), data); err != nil {
		log.Crit("Failed to store raw transaction", "err", err)
	}
}

// encodeBlockNumber encodes a number as big endian uint64
func encodeNumber(number uint64) []byte {
	enc := make([]byte, 8)
	binary.BigEndian.PutUint64(enc, number)
	return enc
}
