package tx

import (
	"context"
	"math/big"
	"sync"

	"github.com/Onther-Tech/plasma-evm/common"
	"github.com/Onther-Tech/plasma-evm/core/types"
	"github.com/Onther-Tech/plasma-evm/ethclient"
	"github.com/Onther-Tech/plasma-evm/log"
)

// TODO: Belows should be moved into core/types package.

// (nonces, price) is set by TransactionManager.
type RawTransaction struct {
	Index uint64

	// transaction field
	Nonce     *big.Int
	From      common.Address
	GasLimit  uint64
	Recipient *common.Address
	Amount    *big.Int
	Payload   []byte

	AllowRevert bool

	PendingTxs  types.Transactions
	MinedTxHash common.Hash
	Reverted    bool

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

func (raw *RawTransaction) AddPending(tx *types.Transaction) {
	raw.lock.Lock()
	defer raw.lock.Unlock()

	raw.PendingTxs = append(raw.PendingTxs, tx)
}

// ClearPendings clears all pending transactions and sets mined transaction hash if transaction is mined .
func (raw *RawTransaction) ClearPendings(backend *ethclient.Client, force bool) (bool, error) {
	raw.lock.Lock()
	defer raw.lock.Unlock()

	mined := false

	for _, tx := range raw.PendingTxs {
		receipt, err := backend.TransactionReceipt(context.Background(), tx.Hash())

		if receipt == nil {
			continue
		}

		if err != nil {
			log.Error("Failed to get transaction receipt", "err", err)
			continue
		}

		raw.Reverted = receipt.Status == 0
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
