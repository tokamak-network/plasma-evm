package tx

import (
	"context"
	"errors"
	"math/big"
	"strings"
	"sync"
	"time"

	"github.com/Onther-Tech/plasma-evm/accounts"
	"github.com/Onther-Tech/plasma-evm/accounts/keystore"
	"github.com/Onther-Tech/plasma-evm/common"
	"github.com/Onther-Tech/plasma-evm/core"
	"github.com/Onther-Tech/plasma-evm/core/types"
	"github.com/Onther-Tech/plasma-evm/ethclient"
	"github.com/Onther-Tech/plasma-evm/ethdb"
	"github.com/Onther-Tech/plasma-evm/log"
	"github.com/Onther-Tech/plasma-evm/params"
)

const (
	MaxNumPending = 10 // The maximum number of transactions that a raw transaction can have.
	MaxNumTask    = 500
	MaxNumKnownTx = 5
)

var (
	ErrLockedAccount    = errors.New("account is locked")
	ErrUnknownAccount   = errors.New("account not found in keystore")
	ErrKnownTransaction = errors.New("known transaction")
)

type TransactionManager struct {
	config *Config

	ks      *keystore.KeyStore
	backend *ethclient.Client
	db      ethdb.Database

	addresses []common.Address // list of account address

	queue  map[common.Address]RawTransactions // raw transactions to be sent
	nonces map[common.Address]uint64          // account nonce

	numKnownErr map[common.Hash]uint64 // number of know transaction error

	gasPrice *big.Int

	taskCh chan *RawTransaction

	lock sync.RWMutex
	quit chan struct{}
}

func NewTransactionManager(ks *keystore.KeyStore, backend *ethclient.Client, db ethdb.Database, config *Config) (*TransactionManager, error) {
	tm := &TransactionManager{
		config: config,

		ks:      ks,
		db:      db,
		backend: backend,

		queue:  make(map[common.Address]RawTransactions),
		nonces: make(map[common.Address]uint64),

		numKnownErr: make(map[common.Hash]uint64),

		gasPrice: ReadGasPrice(db),

		taskCh: make(chan *RawTransaction, MaxNumTask),

		quit: make(chan struct{}),
	}

	n := ReadNumAddr(db)

	if n == MaxAccountNum {
		return nil, errors.New("failed to read account number in database")
	}

	log.Info("Transaction manager loaded", "numAccounts", n)

	var (
		i   uint64
		err error
	)

	for i = 0; i < n; i++ {
		addr := ReadAddr(db, uint64(i))

		if tm.queue[addr] != nil {
			log.Error("Duplicated account", "addr", addr)
			return nil, errors.New("duplicated account")
		}

		tm.queue[addr] = ReadRawTxs(db, addr)
		log.Info("Previous transactions are loaded", "addr", addr, "txs", len(tm.queue[addr]))

		tm.addresses = append(tm.addresses, addr)

		tm.nonces[addr] = ReadAddrNonce(db, addr)
		if tm.nonces[addr] == 0 {
			tm.nonces[addr], err = backend.NonceAt(context.Background(), addr, nil)
			if err != nil {
				log.Error("Failed to read account nonce", "err", err)
				return nil, err
			}
			WriteAddrNonce(db, addr, tm.nonces[addr])
		}
	}

	return tm, nil
}

// Add adds raw transaction to the queue.
func (tm *TransactionManager) Add(account accounts.Account, raw *RawTransaction) error {
	tm.lock.Lock()
	defer tm.lock.Unlock()

	addr := account.Address

	if !tm.ks.HasAddress(addr) {
		return ErrUnknownAccount
	}

	// Update database for the first raw transaction from the account.
	if tm.queue[addr] == nil {
		n := len(tm.queue)
		WriteNumAddr(tm.db, uint64(n+1))

		tm.addresses = append(tm.addresses, addr)
		WriteAddr(tm.db, uint64(n), addr)

		tm.queue[addr] = make(RawTransactions, 0)
		log.Debug("New account is added to transaction manager", "addr", addr)
	}

	raw.Nonce = big.NewInt(int64(tm.nonces[addr]))
	tm.nonces[addr]++

	i := tm.indexOf(addr)
	if i < 0 {
		log.Error("Failed to find account nonce", "addr", addr)
	} else {
		WriteAddrNonce(tm.db, addr, tm.nonces[addr])
	}

	// enqueue raw transaction
	tm.queue[addr] = append(tm.queue[addr], raw)
	WriteRawTxs(tm.db, addr, tm.queue[addr])

	return nil
}

// Count returns the number of raw transactions corresponding to the transaction.
func (tm *TransactionManager) Count(account accounts.Account, tx *types.Transaction) uint64 {
	tm.lock.Lock()
	defer tm.lock.Unlock()

	if tm.queue[account.Address] == nil {
		return 0
	}

	var count uint64
	hash := tx.Hash()

	for _, queuedTx := range tm.queue[account.Address] {
		if queuedTx.Hash() == hash {
			count++
		}
	}

	return count
}

func (tm *TransactionManager) Start() {
	// adjust coordinates gas prices at reasonable prices.
	adjust := func(decrease bool) {
		previous := new(big.Int).Set(tm.gasPrice)

		if decrease {
			tm.gasPrice.Mul(new(big.Int).Div(tm.gasPrice, big.NewInt(4)), big.NewInt(3))
			if tm.gasPrice.Cmp(tm.config.MinGasPrice) < 0 {
				tm.gasPrice.Set(tm.config.MinGasPrice)
			}
		} else {
			tm.gasPrice.Mul(new(big.Int).Div(tm.gasPrice, big.NewInt(2)), big.NewInt(3))
			if tm.gasPrice.Cmp(tm.config.MaxGasPrice) > 0 {
				tm.gasPrice.Set(tm.config.MaxGasPrice)
			}
		}

		WriteGasPrice(tm.db, tm.gasPrice)

		previousGwei := new(big.Float).Quo(new(big.Float).SetInt(previous), new(big.Float).SetInt64(params.GWei)).String() + " Gwei"
		adjustGwei := new(big.Float).Quo(new(big.Float).SetInt(tm.gasPrice), new(big.Float).SetInt64(params.GWei)).String() + " Gwei"

		if previous.Cmp(tm.gasPrice) != 0 {
			log.Info("Gas price adjusted", "previous", previousGwei, "adjusted", adjustGwei)
		}
	}

	// delete mined transactions from queue.
	clearQueue := func(addr common.Address) {
		tm.lock.Lock()
		defer tm.lock.Unlock()

		queue := tm.queue[addr]

		// short circuit if queue is nil or empty.
		if queue == nil || len(queue) == 0 {
			return
		}

		for _, raw := range queue {
			ok, err := raw.ClearPendings(tm.backend, false)
			if err != nil {
				log.Error("Failed to clear pending transaction. Check rootchain provider", "err", err)
				break
			}

			if !ok {
				break
			}

			log.Info("Pending transactions are removed", "addr", addr)
		}

		// update database
		WriteRawTxs(tm.db, addr, queue)
	}

	// send a single raw transaction to root chain.
	send := func(addr common.Address, raw *RawTransaction) (common.Hash, error) {
		tm.lock.Lock()
		defer tm.lock.Unlock()

		// short circuit if transacrtion was already mined
		if raw.Mined(tm.backend) {
			return raw.MinedTxHash, nil
		}

		// subscribe new block mined event
		newHeaderEvents := make(chan *types.Header)
		newHeaderSub, err := tm.backend.SubscribeNewHead(context.Background(), newHeaderEvents)
		defer newHeaderSub.Unsubscribe()

		if err != nil {
			log.Error("Failed to subscribe new block event", "err", err)
		}

		clearHeaderEvent := func() {
			for len(newHeaderEvents) > 0 {
				<-newHeaderEvents
			}
		}

		// account to send transaction
		from := accounts.Account{Address: addr}

		// helper to avoid recursive read lock
		var f func() (common.Hash, error)

		f = func() (common.Hash, error) {
			tx := raw.ToTransaction(tm.gasPrice)
			signedTx, err := tm.ks.SignTx(from, tx, tm.config.ChainId)

			if err != nil {
				log.Error("failed to sign transaction", "err", err, "raw", raw.Hash(), "tx", tx.Hash())
				return signedTx.Hash(), err
			}

			raw.AddPending(signedTx)
			WriteRawTxs(tm.db, addr, tm.queue[addr])

			err = tm.backend.SendTransaction(context.Background(), signedTx)

			if err == nil {
				log.Info("Transaction sent", "hash", signedTx.Hash(), "raw", raw)
				return signedTx.Hash(), nil
			}

			errMessage := strings.ToLower(err.Error())

			// short circuit if operator has not enough ether
			if strings.Contains(errMessage, "insufficient funds for gas * price + value") {
				return signedTx.Hash(), core.ErrInsufficientFunds
			}

			// resubmit transaction in pending intarval loop
			if strings.Contains(errMessage, "replacement transaction underpriced") {
				return signedTx.Hash(), core.ErrReplaceUnderpriced
			}

			// resubmit transaction in pending intarval loop
			if strings.Contains(errMessage, "transaction underpriced") {
				return signedTx.Hash(), core.ErrReplaceUnderpriced
			}

			// resubmit transaction at most MAX_NUM_KNOWN_TX times.
			if strings.Contains(errMessage, "known transaction") {
				tm.numKnownErr[signedTx.Hash()]++

				if tm.numKnownErr[signedTx.Hash()] == MaxNumKnownTx {
					tm.numKnownErr[signedTx.Hash()] = 0
					return signedTx.Hash(), ErrKnownTransaction
				}

				clearHeaderEvent()

				select {
				case <-newHeaderEvents:
					return signedTx.Hash(), ErrKnownTransaction
				case <-tm.quit:
					return signedTx.Hash(), nil
				}

			}

			// resubmit transaction with nonce increased.
			if strings.Contains(errMessage, "nonce too low") {
				// increase nonce immediately if only 1 transaction is pending.
				if len(raw.PendingTxs) == 1 {
					tm.nonces[addr], err = tm.backend.NonceAt(context.Background(), addr, nil)
					if err != nil {
						log.Error("Failed to read account nonce", "err", err)
					} else {
						raw.Nonce = big.NewInt(int64(tm.nonces[addr]))
						WriteAddrNonce(tm.db, addr, tm.nonces[addr])
					}
					return f()
				}

				// if more than 1 transactions are pending, increase nonce carefully.
				// TODO: count and increase nonces
				return signedTx.Hash(), nil
			}

			// return unknown error
			log.Error("Failed to send transaction to root chain.", "err", err)
			return signedTx.Hash(), err
		}

		return f()
	}

	ticker := time.NewTicker(tm.config.Resubmit)
	defer ticker.Stop()

	for {
		select {
		case _, ok := <-ticker.C:
			if !ok {
				continue
			}

			ctx := context.Background()

			for addr, _ := range tm.queue {
				go func(addr common.Address) {
					log.Trace("TransactionManager iterates", "addr", addr)
					queue := tm.queue[addr]

					clearQueue(addr)

					if len(queue) == 0 {
						return
					}

					var raw *RawTransaction

					// find next pending raw transaction
					for _, pending := range queue {
						if !pending.Mined(tm.backend) {
							raw = pending
							break
						}
					}

					if raw == nil {
						return
					}

					hash, err := send(addr, raw)

					// resubmit transaction in pending intarval loop
					if err == core.ErrReplaceUnderpriced {
						log.Debug("Gas price is adjusted for underpriced transaction error")
						adjust(false)
						hash, err = send(addr, raw)
						return
					}

					// short circuit if operator has not enough fund.
					if err == core.ErrInsufficientFunds || err == core.ErrReplaceUnderpriced {
						log.Error("Operator doesn't have enough fund to run the chain.")
						hash, err = send(addr, raw)
						return
					}

					receipt, err2 := tm.backend.TransactionReceipt(ctx, hash)

					if receipt == nil && err == ErrKnownTransaction && tm.numKnownErr[hash] <= MaxNumKnownTx {
						tm.numKnownErr[hash]++
						return
					}

					if receipt == nil {
						log.Warn("Trasaction not found. It may be pending", "err", err2, "hash", hash.Hex())
					} else if err2 != nil {
						log.Error("Send transaction failed", "err", err2, "hash", hash.Hex())
					} else if receipt != nil {
						log.Info("Transaction is mined", "addr", addr, "hash", hash.Hex(), "caption", raw.Caption)

						if !raw.AllowRevert && receipt.Status == 0 {
							// TODO: Do something..
						}

						return
					}

					adjusted := false

					// handle previous submit errors
					if err == ErrKnownTransaction {
						log.Debug("Gas price is adjusted for known transaction error")
						adjusted = true
					}

					if err != nil && !adjusted {
						log.Debug("Gas price is adjusted for unknown transaction error")
						adjust(false)
					}

					hash, err = send(addr, raw)

					if err != nil && err != ErrKnownTransaction {
						log.Error("Failed to submit block to root chain.", "err", err)
					}
				}(addr)
			}

		case <-tm.quit:
			log.Info("TransactionManager stopped")
			return
		}
	}
}

func (tm *TransactionManager) indexOf(addr common.Address) int {
	var i int
	for i = 0; i < len(tm.addresses); i++ {
		if tm.addresses[i] == addr {
			return i
		}
	}

	return -1
}

func (tm *TransactionManager) Stop() {
	close(tm.quit)
}
