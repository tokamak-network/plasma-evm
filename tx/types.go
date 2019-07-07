package tx

import (
	"context"
	"fmt"
	"github.com/Onther-Tech/plasma-evm"
	"github.com/pkg/errors"
	"math/big"
	"sync"

	"github.com/Onther-Tech/plasma-evm/common"
	"github.com/Onther-Tech/plasma-evm/core/types"
	"github.com/Onther-Tech/plasma-evm/ethclient"
)

const (
	MaxNumPending = 128 // The maximum number of transactions that a raw transaction can have.
)

var (
	ErrTooManyPending = errors.New("Too many pending transactions")
)

// TODO: Belows should be moved into core/types package.

// (nonce, price) is set by TransactionManager.
// RawTransaction represents metadata for ethereum transaction.
// A RawTransaction with nonce and gas price is signed and sent to ethereum JSONRPC server, and it is stored in PendingTxs.
// PendingTxs is checked if it is mined and confirmed. Unconfirmed and mined pending transactions
type RawTransaction struct {
	Index               uint64
	ConfirmedIndex      uint64
	ResendCount         uint64
	LastSentBlockNumber uint64

	// transaction field
	Nonce     *big.Int
	From      common.Address
	GasLimit  uint64
	Recipient *common.Address
	Amount    *big.Int
	Payload   []byte

	AllowRevert bool

	PendingTxs       types.Transactions
	MinedTxHash      common.Hash
	MinedBlockNumber *big.Int
	Reverted         bool

	Caption string

	sendLock sync.Mutex
	lock     sync.RWMutex
}

func NewRawTransaction(from common.Address, gasLimit uint64, receipt *common.Address, amount *big.Int, payload []byte, allowRevert bool, caption string) *RawTransaction {
	rawTx := &RawTransaction{
		From:        from,
		GasLimit:    gasLimit,
		Recipient:   receipt,
		Amount:      amount,
		Payload:     payload,
		AllowRevert: allowRevert,
		Caption:     caption,
	}

	return rawTx
}

func (raw *RawTransaction) getCaption() string {
	if raw.ResendCount == 0 {
		return raw.Caption
	}

	return fmt.Sprintf("%s (resend: %d)", raw.Caption, raw.ResendCount)
}

func (raw *RawTransaction) ToTransaction(gasPrice *big.Int) *types.Transaction {
	var to common.Address
	if raw.Recipient != nil {
		to = *raw.Recipient
	}

	return types.NewTransaction(raw.Nonce.Uint64(), to, raw.Amount, raw.GasLimit, gasPrice, raw.Payload)
}

func (raw *RawTransaction) Hash() common.Hash {
	return types.NewTransaction(0, *raw.Recipient, raw.Amount, raw.GasLimit, big.NewInt(0), raw.Payload).Hash()
}

func (raw *RawTransaction) Equal(tx *types.Transaction, chainId *big.Int) bool {
	from, err := types.Sender(types.NewEIP155Signer(chainId), tx)
	if err != nil {
		return false
	}

	converted := toRawTransaction(from, tx, false, "")
	return raw.Hash() == converted.Hash()
}

func (raw *RawTransaction) AddPending(tx *types.Transaction) error {
	raw.lock.Lock()
	defer raw.lock.Unlock()

	if len(raw.PendingTxs) >= MaxNumPending {
		return ErrTooManyPending
	}

	raw.PendingTxs = append(raw.PendingTxs, tx)
	return nil
}

// CheckMined clears all pending transactions and sets mined transaction hash if transaction is mined .
func (raw *RawTransaction) CheckMined(backend *ethclient.Client, force bool) (mined bool, err error) {
	raw.lock.Lock()
	defer raw.lock.Unlock()

	// short circuit if already mined
	if (raw.MinedTxHash != common.Hash{}) {
		return true, nil
	}

	for _, tx := range raw.PendingTxs {
		receipt, err := backend.TransactionReceipt(context.Background(), tx.Hash())

		if receipt == nil {
			continue
		}

		if err != nil {
			return false, err
		}

		raw.Reverted = receipt.Status == 0
		raw.MinedBlockNumber = receipt.BlockNumber
		raw.MinedTxHash = tx.Hash()

		mined = true
		break
	}

	if mined {
		raw.PendingTxs = make(types.Transactions, 0)
	}

	return mined, nil
}

func (raw *RawTransaction) Mined(backend *ethclient.Client) bool {
	raw.lock.Lock()
	defer raw.lock.Unlock()

	return (raw.MinedTxHash != common.Hash{})
}

// Removed returns whether the mined transaction is removed from ethereum blockchain.
func (raw *RawTransaction) Removed(backend *ethclient.Client) (removed bool, err error) {
	receipt, err := backend.TransactionReceipt(context.Background(), raw.MinedTxHash)

	if err == ethereum.NotFound {
		return true, nil
	}

	if err != nil {
		return false, err
	}

	if receipt == nil {
		return true, nil
	}

	return false, nil
}

func (raw *RawTransaction) PrepareToResend() {
	raw.ResendCount++
	raw.MinedTxHash = common.Hash{}
	raw.LastSentBlockNumber = 0
	raw.ConfirmedIndex = 0
	raw.Nonce = new(big.Int)
	raw.MinedBlockNumber = new(big.Int)
	raw.Reverted = false
	raw.PendingTxs = make(types.Transactions, 0)
}

func (raw *RawTransaction) Confirmed(backend *ethclient.Client, currentBlockNumber *big.Int) bool {
	raw.lock.Lock()
	defer raw.lock.Unlock()

	// short circuit if transaction is not mined yet.
	if (raw.MinedTxHash == common.Hash{}) {
		return false
	}

	receipt, _ := backend.TransactionReceipt(context.Background(), raw.MinedTxHash)
	if receipt == nil {
		return false
	}
	raw.MinedBlockNumber = receipt.BlockNumber

	if new(big.Int).Add(receipt.BlockNumber, big.NewInt(Confirmation)).Cmp(currentBlockNumber) > 0 {
		return false
	}

	return true
}

func (raw *RawTransaction) HasPending(tx *types.Transaction) bool {
	raw.lock.Lock()
	defer raw.lock.Unlock()

	for _, pending := range raw.PendingTxs {
		if pending.Hash() == tx.Hash() {
			return true
		}
	}

	return false
}

func toRawTransaction(from common.Address, tx *types.Transaction, allowRevert bool, caption string) *RawTransaction {
	return NewRawTransaction(from, tx.Gas(), tx.To(), tx.Value(), tx.Data(), allowRevert, caption)
}

type RawTransactions []*RawTransaction

type RawTransactionsByIndex RawTransactions

func (r RawTransactionsByIndex) Len() int           { return len(r) }
func (r RawTransactionsByIndex) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
func (r RawTransactionsByIndex) Less(i, j int) bool { return r[i].Index < r[j].Index }
