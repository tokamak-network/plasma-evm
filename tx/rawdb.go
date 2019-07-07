package tx

import (
	"encoding/binary"
	"github.com/Onther-Tech/plasma-evm/ethdb"
	"math"
	"math/big"

	"github.com/Onther-Tech/plasma-evm/common"
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

	numRawTxsPrefix = []byte("num-address-raw-txs") // numRawTxsPrefix + address -> # of raw transactions from the address

	// TODO: design new scheme to reduce I/O for resend + pending txs
	numConfirmedTxsPrefix = []byte("num-confirmed-raw-txs") // numConfirmedTxsPrefix + account address -> # of confirmed raw transactions
	confirmedTxsPrefix    = []byte("confirmed-raw-txs")     // confirmedTxsPrefix + account address + i (uint64 big endian) -> i-th confirmed raw transaction
	unconfirmedTxsPrefix  = []byte("unonfirmed-raw-txs")    // unconfirmedIndexPrefix + account address -> unconfirmed raw transactions
	pendingTxsPrefix      = []byte("pending-raw-txs")       // pendingTxsPrefix + account address -> (resend + pending) raw transactions

	rawTxHashPrefix = []byte("raw-tx-hash") // rawTxHashPrefix + account address + raw transaction hash -> raw transaction without index
)

func ReadGasPrice(db ethdb.Reader) *big.Int {
	data, _ := db.Get(gasPriceKey)

	if len(data) == 0 {
		return new(big.Int).SetInt64(10 * params.GWei)
	}

	var gasPrice big.Int
	if err := rlp.DecodeBytes(data, &gasPrice); err != nil {
		log.Crit("Failed to decode gas price", "err", err)
		return new(big.Int).SetInt64(10 * params.GWei)
	}

	return &gasPrice
}

func WriteGasPrice(db ethdb.KeyValueWriter, gasPrice *big.Int) {
	data, err := rlp.EncodeToBytes(gasPrice)
	if err != nil {
		log.Crit("Failed to encode gas price", "err", err)
	}
	if err := db.Put(gasPriceKey, data); err != nil {
		log.Crit("Failed to store gas price", "err", err)
	}
}

func ReadNumAddr(db ethdb.Reader) uint64 {
	data, _ := db.Get(numAddrKey)

	if len(data) == 0 {
		return 0
	}

	var n uint64
	if err := rlp.DecodeBytes(data, &n); err != nil {
		log.Crit("Failed to decode account number", "err", err)
		return MaxUint64
	}

	return n
}

func WriteNumAddr(db ethdb.KeyValueWriter, n uint64) {
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

func ReadAddr(db ethdb.Reader, i uint64) common.Address {
	data, _ := db.Get(addrKey(i))

	if len(data) == 0 {
		return common.Address{}
	}

	var addr common.Address
	if err := rlp.DecodeBytes(data, &addr); err != nil {
		log.Crit("Failed to decode account address", "err", err)
		return common.Address{}
	}

	return addr
}

func WriteAddr(db ethdb.KeyValueWriter, i uint64, addr common.Address) {
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

func ReadAddrNonce(db ethdb.Reader, addr common.Address) uint64 {
	data, _ := db.Get(addrNonceKey(addr))

	if len(data) == 0 {
		return 0
	}

	var nonce uint64
	if err := rlp.DecodeBytes(data, &nonce); err != nil {
		log.Crit("Failed to decode account nonce", "err", err)
		return 0
	}

	return nonce
}

func WriteAddrNonce(db ethdb.KeyValueWriter, addr common.Address, nonce uint64) {
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

func ReadNumRawTxs(db ethdb.Reader, addr common.Address) uint64 {
	data, _ := db.Get(numRawTxsKey(addr))

	if len(data) == 0 {
		return 0
	}

	var n uint64
	if err := rlp.DecodeBytes(data, &n); err != nil {
		log.Crit("Failed to decode number of raw transactions", "err", err)
		return 0
	}

	return n
}

func WriteNumRawTxs(db ethdb.KeyValueWriter, addr common.Address, n uint64) {
	data, err := rlp.EncodeToBytes(n)
	if err != nil {
		log.Crit("Failed to encode number of raw transactions", "err", err)
	}
	if err := db.Put(numRawTxsKey(addr), data); err != nil {
		log.Crit("Failed to store number of raw transactions", "err", err)
	}
}

func numConfirmedRawTxsKey(addr common.Address) []byte {
	return append(numConfirmedTxsPrefix, addr.Bytes()...)
}

func ReadNumConfirmedRawTxs(db ethdb.Reader, addr common.Address) uint64 {
	data, _ := db.Get(numConfirmedRawTxsKey(addr))

	if len(data) == 0 {
		return 0
	}

	var n uint64
	if err := rlp.DecodeBytes(data, &n); err != nil {
		log.Crit("Failed to decode number of confirmed raw transactions", "err", err)
		return MaxUint64
	}

	return n
}

func WriteNumConfirmedRawTxs(db ethdb.KeyValueWriter, addr common.Address, n uint64) {
	data, err := rlp.EncodeToBytes(n)
	if err != nil {
		log.Crit("Failed to encode number of raw transactions", "err", err)
	}
	if err := db.Put(numConfirmedRawTxsKey(addr), data); err != nil {
		log.Crit("Failed to store number of raw transactions", "err", err)
	}
}

func confirmedTxKey(addr common.Address, i uint64) []byte {
	return append(append(confirmedTxsPrefix, addr.Bytes()...), encodeNumber(i)...)
}

func ReadConfirmedTx(db ethdb.Reader, addr common.Address, i uint64) *RawTransaction {
	data, _ := db.Get(confirmedTxKey(addr, i))

	if len(data) == 0 {
		log.Crit("Failed to decode unconfirmed transactions")
	}

	var raw RawTransaction
	if err := rlp.DecodeBytes(data, &raw); err != nil {
		log.Crit("Failed to decode number of raw transactions", "err", err)
		return nil
	}

	return &raw
}

func WriteConfirmedTx(db ethdb.KeyValueWriter, addr common.Address, i uint64, raw *RawTransaction) {
	data, err := rlp.EncodeToBytes(raw)
	if err != nil {
		log.Crit("Failed to encode confirmed raw transaction", "err", err)
	}
	if err := db.Put(confirmedTxKey(addr, i), data); err != nil {
		log.Crit("Failed to store confirmed raw transaction", "err", err)
	}
}

func unconfirmedTxsKey(addr common.Address) []byte {
	return append(unconfirmedTxsPrefix, addr.Bytes()...)
}

func ReadUnconfirmedTxs(db ethdb.Reader, addr common.Address) RawTransactions {
	data, _ := db.Get(unconfirmedTxsKey(addr))

	if len(data) == 0 {
		return RawTransactions{}
	}

	var txs RawTransactions
	if err := rlp.DecodeBytes(data, &txs); err != nil {
		log.Crit("Failed to decode unconfirmed transactions", "err", err)
		return nil
	}

	return txs
}

func WriteUnconfirmedTxs(db ethdb.KeyValueWriter, addr common.Address, txs RawTransactions) {
	data, err := rlp.EncodeToBytes(txs)
	if err != nil {
		log.Crit("Failed to encode unconfirmed transactions", "err", err)
	}
	if err := db.Put(unconfirmedTxsKey(addr), data); err != nil {
		log.Crit("Failed to store unconfirmed transactions", "err", err)
	}
}

func pendingTxsKey(addr common.Address) []byte {
	return append(pendingTxsPrefix, addr.Bytes()...)
}

func ReadPendingTxs(db ethdb.Reader, addr common.Address) RawTransactions {
	data, _ := db.Get(pendingTxsKey(addr))

	if len(data) == 0 {
		return RawTransactions{}
	}

	var txs RawTransactions
	if err := rlp.DecodeBytes(data, &txs); err != nil {
		log.Crit("Failed to decode pending transactions", "err", err)
		return nil
	}

	return txs
}

func WritePendingTxs(db ethdb.KeyValueWriter, addr common.Address, txs RawTransactions) {
	data, err := rlp.EncodeToBytes(txs)
	if err != nil {
		log.Crit("Failed to encode pending transactions", "err", err)
	}
	if err := db.Put(pendingTxsKey(addr), data); err != nil {
		log.Crit("Failed to store pending transactions", "err", err)
	}
}

func rawTxHashKey(addr common.Address, rawHash common.Hash) []byte {
	return append(append(rawTxHashPrefix, addr.Bytes()...), rawHash.Bytes()...)
}

func ReadRawTxHash(db ethdb.Reader, addr common.Address, rawHash common.Hash) *RawTransaction {
	data, _ := db.Get(rawTxHashKey(addr, rawHash))

	if len(data) == 0 {
		return nil
	}

	var raw RawTransaction
	if err := rlp.DecodeBytes(data, &raw); err != nil {
		log.Crit("Failed to decode raw transaction", "err", err, "addr", addr)
		return nil
	}

	return &raw
}

func WriteRawTxHash(db ethdb.KeyValueWriter, addr common.Address, raw RawTransaction) {
	data, err := rlp.EncodeToBytes(raw)
	if err != nil {
		log.Crit("Failed to encode raw transaction", "err", err)
	}
	if err := db.Put(rawTxHashKey(addr, raw.Hash()), data); err != nil {
		log.Crit("Failed to store raw transaction", "err", err)
	}
}

// encodeBlockNumber encodes a number as big endian uint64
func encodeNumber(number uint64) []byte {
	enc := make([]byte, 8)
	binary.BigEndian.PutUint64(enc, number)
	return enc
}
