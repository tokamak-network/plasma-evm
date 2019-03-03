package tx

import (
	"encoding/binary"
	"github.com/Onther-Tech/plasma-evm/params"
	"math/big"

	"github.com/Onther-Tech/plasma-evm/common"
	"github.com/Onther-Tech/plasma-evm/core/rawdb"
	"github.com/Onther-Tech/plasma-evm/core/types"
	"github.com/Onther-Tech/plasma-evm/log"
	"github.com/Onther-Tech/plasma-evm/rlp"
)

// TODO: Belows should be move into rawdb package.

const (
	MaxAccountNum = 256
)

var (
	gasPriceKey     = []byte("gas-price")
	numAddrKey      = []byte("num-address")   // numAddrKey -> the number of accounts
	addrPrefix      = []byte("address")       // addrPrefix + i (uint64 big endian) -> i-th account to send transaction
	addrNoncePrefix = []byte("address-nonce") // addrPrefix + address -> i-th account's nonce
	rawTxsPrefix    = []byte("raw-txs")       // rawTxsPrefix + account address -> queued raw transactions to send to root chain
	txQueuePrefix   = []byte("tx-queue")      // txQueuePrefix + raw tx hash -> sent transactions to root chain
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
	if err := rlp.DecodeBytes(data, n); err != nil {
		log.Error("Failed to decode account number", "err", err)
		return MaxAccountNum
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

// rawTxsKey = rawTxsPrefix + address
func rawTxsKey(addr common.Address) []byte {
	return append(rawTxsPrefix, addr.Bytes()...)
}

func ReadRawTxs(db rawdb.DatabaseReader, addr common.Address) RawTransactions {
	data, _ := db.Get(rawTxsKey(addr))

	if len(data) == 0 {
		return RawTransactions{}
	}

	var rawTxs RawTransactions
	if err := rlp.DecodeBytes(data, &rawTxs); err != nil {
		log.Error("Failed to decode raw transactions", "err", err)
		return RawTransactions{}
	}

	return rawTxs
}

func WriteRawTxs(db rawdb.DatabaseWriter, addr common.Address, rawTxs RawTransactions) {
	data, err := rlp.EncodeToBytes(rawTxs)
	if err != nil {
		log.Crit("Failed to encode raw transactions", "err", err)
	}
	if err := db.Put(rawTxsKey(addr), data); err != nil {
		log.Crit("Failed to store raw transactions", "err", err)
	}
}

// txQueueKey = txQueuePrefix + raw tx hash
func txQueueKey(hash common.Hash) []byte {
	return append(txQueuePrefix, hash.Bytes()...)
}

func ReadTxQueue(db rawdb.DatabaseReader, hash common.Hash) types.Transactions {
	data, _ := db.Get(txQueueKey(hash))

	if len(data) == 0 {
		return types.Transactions{}
	}

	var txs types.Transactions
	if err := rlp.DecodeBytes(data, &txs); err != nil {
		log.Error("Failed to decode transaction queue", "err", err)
		return types.Transactions{}
	}

	return txs
}

func WriteTxQueue(db rawdb.DatabaseWriter, hash common.Hash, txs types.Transactions) {
	data, err := rlp.EncodeToBytes(txs)
	if err != nil {
		log.Crit("Failed to encode transaction queue", "err", err)
	}
	if err := db.Put(txQueueKey(hash), data); err != nil {
		log.Crit("Failed to store transaction queue", "err", err)
	}
}

// encodeBlockNumber encodes a number as big endian uint64
func encodeNumber(number uint64) []byte {
	enc := make([]byte, 8)
	binary.BigEndian.PutUint64(enc, number)
	return enc
}
