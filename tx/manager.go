package tx

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/Onther-Tech/plasma-evm"
	"github.com/Onther-Tech/plasma-evm/accounts"
	"github.com/Onther-Tech/plasma-evm/accounts/keystore"
	"github.com/Onther-Tech/plasma-evm/common"
	"github.com/Onther-Tech/plasma-evm/core"
	"github.com/Onther-Tech/plasma-evm/core/types"
	"github.com/Onther-Tech/plasma-evm/ethclient"
	"github.com/Onther-Tech/plasma-evm/ethdb"
	"github.com/Onther-Tech/plasma-evm/log"
	"github.com/Onther-Tech/plasma-evm/params"
	"github.com/hashicorp/golang-lru"
)

const (
	MaxNumTask    = 500
	MaxNumKnownTx = 5
	SendDelay     = 2

	numKnownErrLimit = 256

	// TODO: make below configurable
	Confirmation      = 32
	ConfirmationDelay = 4
)

var (
	ErrLockedAccount    = errors.New("account is locked")
	ErrUnknownAccount   = errors.New("account not found in keystore")
	ErrKnownTransaction = errors.New("known transaction")
	ErrDuplicateRaw     = errors.New("duplicate raw transaction")
	ErrNoDuplicateRaw   = errors.New("there is no duplicate raw transaction")
)

// TODO: Add JSONRPC API for TransactionManager
type TransactionManager struct {
	config *Config

	ks      *keystore.KeyStore
	backend *ethclient.Client
	db      ethdb.Database

	currentBlockNumber *big.Int // current block number of root chian network
	gasPrice           *big.Int

	addresses []common.Address // list of account address

	nonce map[common.Address]uint64 // account nonce

	lastInspectTime time.Time

	numKnownErrCache *lru.Cache // Cache for the most recent block bodies

	taskCh chan *RawTransaction

	lock         sync.RWMutex
	gasPriceLock sync.Mutex
	quit         chan struct{}
}

func NewTransactionManager(ks *keystore.KeyStore, backend *ethclient.Client, db ethdb.Database, config *Config) (*TransactionManager, error) {
	numKnownErrCache, _ := lru.New(numKnownErrLimit)

	tm := &TransactionManager{
		config: config,

		ks:      ks,
		db:      db,
		backend: backend,

		currentBlockNumber: new(big.Int),
		gasPrice:           new(big.Int),

		nonce: make(map[common.Address]uint64),

		numKnownErrCache: numKnownErrCache,

		taskCh: make(chan *RawTransaction, MaxNumTask),

		quit: make(chan struct{}),
	}

	gasPrice := ReadGasPrice(db)

	if config.MinGasPrice.Cmp(config.MaxGasPrice) > 0 {
		return nil, errors.New("min gas price cannot exceed max gas price")
	}

	if config.GasPrice.Cmp(big.NewInt(0)) == 0 {
		gasPrice = new(big.Int).Set(DefaultConfig.GasPrice)
		log.Info("Use default gas price", "gasprice", gasPrice)
	}

	if gasPrice.Cmp(config.MinGasPrice) < 0 {
		gasPrice = new(big.Int).Set(config.MinGasPrice)
		log.Warn("Gas price is below the min gas price.")
	}

	if gasPrice.Cmp(config.MaxGasPrice) > 0 {
		gasPrice = new(big.Int).Set(config.MaxGasPrice)
		log.Warn("Gas price is above the max gas price.")
	}

	tm.gasPrice = gasPrice

	numAddrs := ReadNumAddr(db)

	if numAddrs == MaxUint64 {
		return nil, errors.New("failed to read account number in database")
	}

	var (
		i   uint64
		err error
	)

	for i = 0; i < numAddrs; i++ {
		addr := ReadAddr(db, uint64(i))
		tm.addresses = append(tm.addresses, addr)

		if _, ok := tm.nonce[addr]; ok {
			log.Error("Duplicated account found", "addr", addr)
			continue
		}

		numConfirmedRawTxs := ReadNumConfirmedRawTxs(tm.db, addr)
		if numConfirmedRawTxs == MaxUint64 {
			return nil, errors.New(fmt.Sprintf("failed to read number of confirmed raw transaction of %s", addr.String()))
		}

		log.Info("Previous account loaded", "addr", addr, "numConfirmedRawTxs", numConfirmedRawTxs)

		tm.nonce[addr] = ReadAddrNonce(db, addr)
		if tm.nonce[addr] == 0 {
			tm.nonce[addr], err = backend.NonceAt(context.Background(), addr, nil)
			if err != nil {
				log.Error("Failed to read account nonce", "err", err)
				return nil, err
			}
			WriteAddrNonce(db, addr, tm.nonce[addr])
		}
		log.Info("Previous transactions are loaded", "addr", addr, "nonce", tm.nonce[addr])
		tm.inspect(addr)
	}

	log.Info("Transaction manager loaded", "numAccounts", numAddrs)

	return tm, nil
}

// Add adds raw transaction to confirmed.
func (tm *TransactionManager) Add(account accounts.Account, raw *RawTransaction, duplicate bool) error {
	tm.lock.Lock()
	defer tm.lock.Unlock()

	addr := account.Address

	tm.inspect(addr)

	if !tm.ks.HasAddress(addr) {
		return ErrUnknownAccount
	}

	// Update database for the first raw transaction from the account.
	if tm.indexOf(addr) < 0 {

		n := len(tm.addresses)
		WriteNumAddr(tm.db, uint64(n+1))

		tm.addresses = append(tm.addresses, addr)
		WriteAddr(tm.db, uint64(n), addr)

		log.Debug("New account is added to transaction manager", "addr", addr)
	}

	if !duplicate {
		// add unique aw transaction
		if previous := ReadRawTxHash(tm.db, addr, raw.Hash()); previous != nil {
			return ErrDuplicateRaw
		}
		WriteRawTxHash(tm.db, addr, *raw)
	} else {
		// add duplicate raw transaction
		if previous := ReadRawTxHash(tm.db, addr, raw.Hash()); previous == nil {
			return ErrNoDuplicateRaw
		}
	}

	// assign index
	i := ReadNumRawTxs(tm.db, addr)
	raw.Index = i
	WriteNumRawTxs(tm.db, addr, i+1)

	// assign nonce
	raw.Nonce = big.NewInt(int64(tm.nonce[addr]))
	tm.nonce[addr]++
	WriteAddrNonce(tm.db, addr, tm.nonce[addr])

	// enqueue raw transaction
	pending := ReadPendingTxs(tm.db, addr)
	pending = append(pending, raw)
	WritePendingTxs(tm.db, addr, pending)

	log.Info("Raw transaction added", "caption", raw.getCaption(), "from", raw.From)

	return nil
}

func (tm *TransactionManager) Has(account accounts.Account, tx *types.Transaction) bool {
	tm.lock.Lock()
	defer tm.lock.Unlock()

	raw := toRawTransaction(account.Address, tx, false, "")

	return ReadRawTxHash(tm.db, account.Address, raw.Hash()) != nil
}

func (tm *TransactionManager) Start() {
	go tm.confirmLoop()

	// send a single raw transaction to root chain.
	// TODO: make it safe under root chain provider disconnect
	send := func(addr common.Address, raw *RawTransaction, pending RawTransactions) (common.Hash, error) {
		raw.sendLock.Lock()
		defer raw.sendLock.Unlock()

		// short circuit if transaction was already mined
		if raw.Mined(tm.backend) {
			return raw.MinedTxHash, nil
		}

		// subscribe new block mined event
		newHeaderEvents := make(chan *types.Header)
		newHeaderSub, err := tm.backend.SubscribeNewHead(context.Background(), newHeaderEvents)

		close := func() {
			defer func() {
				if err := recover(); err != nil {
					log.Error("New block event unsubscription", "err", err)
				}
			}()
			newHeaderSub.Unsubscribe()
		}

		defer close()

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
			blockNumber := tm.currentBlockNumber.Uint64()

			// short circuit
			if raw.LastSentBlockNumber != 0 && raw.LastSentBlockNumber+SendDelay <= blockNumber {
				return common.Hash{}, nil
			}

			tx := raw.ToTransaction(tm.gasPrice)
			signedTx, err := tm.ks.SignTx(from, tx, tm.config.ChainId)

			if err != nil {
				log.Error("failed to sign transaction", "err", err, "raw", raw.Hash(), "caption", raw.getCaption(), "tx", tx.Hash())
				return signedTx.Hash(), err
			}

			// short circuit raw transaction already has same transaction.
			if raw.HasPending(signedTx) {
				return signedTx.Hash(), nil
			}

			err = raw.AddPending(signedTx)
			if err != nil {
				log.Error(err.Error(), "raw", raw.Hash(), "caption", raw.getCaption(), "tx", tx.Hash())
				return signedTx.Hash(), err
			}
			raw.LastSentBlockNumber = blockNumber

			tm.lock.Lock()
			WritePendingTxs(tm.db, addr, pending)
			tm.lock.Unlock()

			err = tm.backend.SendTransaction(context.Background(), signedTx)

			if err == nil {
				log.Info("Transaction sent", "hash", signedTx.Hash(), "nonce", raw.Nonce, "caption", raw.getCaption())
				return signedTx.Hash(), nil
			}

			errMessage := strings.ToLower(err.Error())

			// short circuit if operator has not enough ether
			if strings.Contains(errMessage, "insufficient funds for gas * price + value") {
				return signedTx.Hash(), core.ErrInsufficientFunds
			}

			if strings.Contains(errMessage, "replacement transaction underpriced") {
				return signedTx.Hash(), core.ErrReplaceUnderpriced
			}

			if strings.Contains(errMessage, "transaction underpriced") {
				return signedTx.Hash(), core.ErrReplaceUnderpriced
			}

			// resubmit transaction at most MAX_NUM_KNOWN_TX times.
			if strings.Contains(errMessage, "known transaction") {
				var cnt uint
				cached, ok := tm.numKnownErrCache.Get(signedTx.Hash())

				if !ok {
					cnt = 1
				} else {
					cnt = cached.(uint) + 1
				}

				tm.numKnownErrCache.Add(signedTx.Hash(), cnt)

				if cnt >= MaxNumKnownTx {
					log.Debug("Transaction is in pool")
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

			// resubmit transaction with new nonce.
			if strings.Contains(errMessage, "nonce too low") || strings.Contains(errMessage, "nonce is too low") || strings.Contains(errMessage, "nonce too high") || strings.Contains(errMessage, "nonce is too high") {
				previousNonce := raw.Nonce.Uint64()

				// update nonce
				tm.nonce[addr], err = tm.backend.NonceAt(context.Background(), addr, nil)

				if err != nil {
					log.Error("Failed to read account nonce", "err", err)
					return f()
				}

				log.Warn("Account nonce has changed by another transaction", "previousNonce", previousNonce, "currentNonce", tm.nonce[addr])
				raw.Nonce = big.NewInt(int64(tm.nonce[addr]))
				WriteAddrNonce(tm.db, addr, tm.nonce[addr])
				WritePendingTxs(tm.db, addr, pending)

				return f()
			}

			// return unknown error
			log.Error("Failed to send transaction to root chain.", "err", err)
			return signedTx.Hash(), err
		}

		return f()
	}

	// send loop
	go func() {
		ticker := time.NewTicker(tm.config.Interval)
		defer ticker.Stop()

		for {
			select {
			case _, ok := <-ticker.C:
				if !ok {
					continue
				}

				for _, addr := range tm.addresses {
					log.Trace("TransactionManager iterates", "addr", addr)

					tm.clearQueue(addr)
					tm.confirmQueue(addr)

					pending := ReadPendingTxs(tm.db, addr)
					if len(pending) == 0 {
						continue
					}

					sort.Sort(RawTransactionsByIndex(pending))

					raw := pending[0]
					hash, err := send(addr, raw, pending)

					// resubmit transaction in pending intarval loop
					if err == core.ErrReplaceUnderpriced {
						log.Debug("Gas price is adjusted for underpriced transaction error")
						tm.adjustGasPrice(raw, false)
						hash, err = send(addr, raw, pending)
						continue
					}

					// short circuit if operator has not enough fund.
					if err == core.ErrInsufficientFunds || err == core.ErrReplaceUnderpriced {
						log.Error("Account doesn't have enough fund to run the chain.", "addr", addr)
						hash, err = send(addr, raw, pending)
						continue
					}

					receipt, err2 := tm.backend.TransactionReceipt(context.Background(), hash)
					adjusted := false

					// short circuit if tx is already mined
					if receipt != nil {
						log.Debug("Raw transaction is already mined", "caption", raw.getCaption(), "hash", receipt.TxHash.String())
						continue
					}

					var cnt uint
					cached, ok := tm.numKnownErrCache.Get(hash)

					if ok {
						cnt = cached.(uint)
					}

					if receipt == nil && err == ErrKnownTransaction && cnt <= MaxNumKnownTx {
						log.Debug("Transaction is in pool")
						tm.numKnownErrCache.Add(hash, cnt+1)
						continue
					}

					// handle previous submit errors
					if err == ErrKnownTransaction {
						log.Debug("Gas price is adjusted for known transaction error")
						adjusted = true
					}

					if err2 == ethereum.NotFound {
						log.Warn("Ethereum Transaction not found. It may be pending", "err", err2, "caption", raw.getCaption(), "hash", hash.Hex())
						adjusted = true
						tm.adjustGasPrice(raw, false)
					}

					if err != nil && !adjusted {
						log.Debug("Gas price is adjusted for unknown transaction error")
						tm.adjustGasPrice(raw, false)
					}

					hash, err = send(addr, raw, pending)

					if err != nil && err != ErrKnownTransaction {
						log.Error("Failed to submit block to root chain.", "err", err)
					}
				}

			case <-tm.quit:
				log.Info("TransactionManager stopped")
				return
			}
		}
	}()
}

// adjustGasPrice adjust gas prices at reasonable prices.
func (tm *TransactionManager) adjustGasPrice(raw *RawTransaction, decrease bool) {
	tm.gasPriceLock.Lock()
	defer tm.gasPriceLock.Unlock()

	var previousTxGasPrice *big.Int

	if raw.Mined(tm.backend) {
		minedTx, _, _ := tm.backend.TransactionByHash(context.Background(), raw.MinedTxHash)
		previousTxGasPrice = new(big.Int).Set(minedTx.GasPrice())
	} else {
		if len(raw.PendingTxs) == 0 {
			previousTxGasPrice = new(big.Int).Set(tm.gasPrice)
		} else {
			lastPendingTx := raw.PendingTxs[len(raw.PendingTxs)-1]
			previousTxGasPrice = new(big.Int).Set(lastPendingTx.GasPrice())
		}
	}

	previousGwei := gasPriceToString(tm.gasPrice)

	if decrease {
		if tm.gasPrice.Cmp(previousTxGasPrice) < 0 {
			tm.gasPrice = new(big.Int).Set(previousTxGasPrice)
		} else {
			tm.gasPrice.Mul(new(big.Int).Div(tm.gasPrice, big.NewInt(10)), big.NewInt(4))
		}

		if tm.gasPrice.Cmp(tm.config.MinGasPrice) < 0 {
			tm.gasPrice.Set(tm.config.MinGasPrice)
		}
	} else {
		if tm.gasPrice.Cmp(previousTxGasPrice) > 0 {
			tm.gasPrice = new(big.Int).Set(previousTxGasPrice)
		} else {
			tm.gasPrice.Mul(new(big.Int).Div(tm.gasPrice, big.NewInt(10)), big.NewInt(12))
		}

		if tm.gasPrice.Cmp(tm.config.MaxGasPrice) > 0 {
			tm.gasPrice.Set(tm.config.MaxGasPrice)
		}
	}

	WriteGasPrice(tm.db, tm.gasPrice)

	previousTxGasPriceGwei := gasPriceToString(previousTxGasPrice)
	adjustGwei := gasPriceToString(tm.gasPrice)
	log.Info("Gas price adjusted", "caption", raw.getCaption(), "decrease", decrease, "previousTxGasPriceGwei ", previousTxGasPriceGwei, "previous", previousGwei, "adjusted", adjustGwei)
}

// clearQueue check raw transaction is mined. Mined raw transactions move to unconfirmed pending.
// Before confirmed, if the mined raw transaction is removed from root chain network, it goes back to the pending again.
func (tm *TransactionManager) clearQueue(addr common.Address) {
	tm.lock.Lock()
	defer tm.lock.Unlock()

	pending := ReadPendingTxs(tm.db, addr)
	sort.Sort(RawTransactionsByIndex(pending))

	// short circuit if pending is nil or empty.
	if pending == nil || len(pending) == 0 {
		return
	}

	// check pending transaction is mined or not
	for _, raw := range pending {
		ok, err := raw.CheckMined(tm.backend, false)
		if err != nil {
			log.Error("Failed to clear pending transactions. Check rootchain provider", "err", err, "caption", raw.getCaption())
			break
		}

		if !ok {
			break
		}

		log.Info("Transaction is mined", "nonce", raw.Nonce, "caption", raw.getCaption(), "reverted", raw.Reverted, "from", addr, "hash", raw.MinedTxHash.String())

		if raw.Reverted {
			log.Error("Transaction is reverted", "caption", raw.getCaption(), "hash", raw.MinedTxHash.String())
		}
		tm.adjustGasPrice(raw, true)
	}

	// remove mined raw transactions
	var minedRaws RawTransactions
	i := 0
	for ; i < len(pending); i++ {
		raw := pending[i]
		if raw == nil {
			break
		}

		if !raw.Mined(tm.backend) {
			if raw.Reverted && !raw.AllowRevert {
				log.Error("Transaction reverted", "caption", raw.getCaption(), "hash", raw.Hash())
			}
			break
		}
		minedRaws = append(minedRaws, raw)
	}

	// update database
	l := len(minedRaws)
	if l != 0 {
		unconfirmed := ReadUnconfirmedTxs(tm.db, addr)
		unconfirmed = append(unconfirmed, minedRaws...)

		pending = pending[l:]
		WritePendingTxs(tm.db, addr, pending)
		WriteUnconfirmedTxs(tm.db, addr, unconfirmed)
	}
}

// confirmQueue check mined raw transaction is confirmed.
// If unconfirmed transaction is removed from canonical chain, insert it into pending pending.
func (tm *TransactionManager) confirmQueue(addr common.Address) {
	tm.lock.Lock()
	defer tm.lock.Unlock()

	if time.Since(tm.lastInspectTime) > time.Second*5 {
		tm.inspect(addr)
	}

	unconfirmed := ReadUnconfirmedTxs(tm.db, addr)
	pending := ReadPendingTxs(tm.db, addr)
	sort.Sort(RawTransactionsByIndex(pending))

	// short circuit if unconfirmed is nil or empty.
	if unconfirmed == nil || len(unconfirmed) == 0 {
		return
	}

	// re-add removed raw transaction into queue
	var lastRemovedRaw *RawTransaction
	var newUnconfirmed RawTransactions
	for _, raw := range unconfirmed {
		removed, err := raw.Removed(tm.backend)
		if err != nil {
			log.Error("Failed to check transaction is removed", "err", err)
			break
		}

		if removed {
			log.Info("Raw transaction is removed", "addr", addr, "caption", raw.getCaption())
			raw.PrepareToResend()
			pending = append(pending, raw)
		} else {
			newUnconfirmed = append(newUnconfirmed, raw)
		}

		lastRemovedRaw = raw
	}

	// update database
	if lastRemovedRaw != nil {
		unconfirmed = newUnconfirmed
		sort.Sort(RawTransactionsByIndex(pending))

		WriteUnconfirmedTxs(tm.db, addr, unconfirmed)
		WritePendingTxs(tm.db, addr, pending)
	}

	// remove already confirmed raw transactions
	numConfirmed := ReadNumConfirmedRawTxs(tm.db, addr)
	currentBlockNumber := new(big.Int).Set(tm.currentBlockNumber)
	i := 0
	for ; i < len(unconfirmed); i++ {
		raw := unconfirmed[i]
		log.Trace("check raw is confirmed", "addr", addr, "caption", raw.getCaption())

		if !raw.Confirmed(tm.backend, currentBlockNumber) {
			break
		}

		log.Info("Transaction is confirmed", "addr", addr, "caption", raw.getCaption())
		raw.ConfirmedIndex = numConfirmed
		WriteConfirmedTx(tm.db, addr, numConfirmed, raw)
		numConfirmed++
	}

	// update database
	if i != 0 {
		unconfirmed = unconfirmed[i:]
		WriteNumConfirmedRawTxs(tm.db, addr, numConfirmed)
		WriteUnconfirmedTxs(tm.db, addr, unconfirmed)
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

// TODO: use SubscribeNewHead with disconnection handling
func (tm *TransactionManager) confirmLoop() {
	closed := false

	newHeaderCh := make(chan *types.Header)
	sub, err := tm.backend.SubscribeNewHead(context.Background(), newHeaderCh)

	close := func() {
		defer func() {
			if err := recover(); err != nil {
				log.Error("New block event unsubscription", "err", err)
			}
		}()
		sub.Unsubscribe()
	}

	defer close()

	if err != nil {
		log.Error("Failed to subscribe root chian new block event", "err", err)
		return
	}

	reconnTimer := time.NewTimer(0)
	<-reconnTimer.C

	reconn := func() {
		if closed {
			return
		}

		sub2, err := tm.backend.SubscribeNewHead(context.Background(), newHeaderCh)

		if err != nil {
			log.Error("Failed to re-subscribe root chian new block event", "err", err)
			reconnTimer.Reset(5 * time.Second)
		} else {
			sub = sub2
			log.Info("Re-subscribe root chian new block event", "err", err)
		}
	}

	var lastConfirmed uint64

	for {
		select {
		case <-tm.quit:
			closed = true
			reconnTimer.Stop()
			return

		case b, ok := <-newHeaderCh:
			if !ok {
				continue
			}

			if lastConfirmed == 0 {
				lastConfirmed = b.Number.Uint64()
			}

			log.Info("New root chain block mined", "number", b.Number)

			tm.lock.Lock()
			tm.currentBlockNumber.Set(b.Number)
			tm.lock.Unlock()

			if lastConfirmed+ConfirmationDelay < b.Number.Uint64() {
				continue
			}

			lastConfirmed = b.Number.Uint64()
			for _, addr := range tm.addresses {
				tm.confirmQueue(addr)
			}

		case err := <-sub.Err():
			log.Error("New block event unsubscribed", "err", err)
			reconn()

		case _, ok := <-reconnTimer.C:
			if ok {
				reconn()
			}
		}
	}
}

func (tm *TransactionManager) Stop() {
	close(tm.quit)
}

func (tm *TransactionManager) inspect(addr common.Address) {
	confirmed := ReadNumConfirmedRawTxs(tm.db, addr)
	unconfirmed := len(ReadUnconfirmedTxs(tm.db, addr))
	pending := len(ReadPendingTxs(tm.db, addr))
	log.Debug("Inspect queue", "addr", addr, "total", int(confirmed)+unconfirmed+pending, "confirmed", confirmed, "unconfiemd", unconfirmed, "pending", pending)

	tm.lastInspectTime = time.Now()
}

func gasPriceToString(gp *big.Int) string {
	return new(big.Float).Quo(new(big.Float).SetInt(gp), new(big.Float).SetInt64(params.GWei)).String() + " Gwei"
}
