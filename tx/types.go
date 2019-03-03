package tx

import (
	"context"
	"github.com/Onther-Tech/plasma-evm/log"
	"math/big"
	"sync"

	"github.com/Onther-Tech/plasma-evm/common"
	"github.com/Onther-Tech/plasma-evm/core/types"
	"github.com/Onther-Tech/plasma-evm/ethclient"
)

// TODO: Belows should be moved into core/types package.

// (nonces, price) is set by TransactionManager.
type RawTransaction struct {
	Nonce       *big.Int
	From        common.Address
	GasLimit    uint64
	Recipient   *common.Address
	Amount      *big.Int
	Payload     []byte
	AllowRevert bool

	PendingTxs  types.Transactions
	MinedTxHash common.Hash
	Reverted    bool

	Caption string

	lock sync.RWMutex
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

// minedTransaction returns mined transaction of a raw transaction.
func (raw *RawTransaction) minedTransaction(backend *ethclient.Client) (int, *types.Transaction) {
	for i, tx := range raw.PendingTxs {
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

		return i, tx
	}

	return -1, nil
}

// ClearPendings clears all pending transactions. It returns true if transaction is mined and removed in pending tx.
func (raw *RawTransaction) ClearPendings(backend *ethclient.Client, force bool) (bool, error) {
	raw.lock.Lock()
	defer raw.lock.Unlock()

	i, _ := raw.minedTransaction(backend)

	if !force && i < 0 {
		return false, nil
	}

	raw.PendingTxs = make(types.Transactions, 0)

	return i >= 0, nil
}

func (raw *RawTransaction) Mined(backend *ethclient.Client) bool {
	raw.lock.Lock()
	defer raw.lock.Unlock()

	if (raw.MinedTxHash != common.Hash{}) {
		return true
	}

	i, _ := raw.minedTransaction(backend)

	return i >= 0
}

func toRawTransaction(from common.Address, tx *types.Transaction, allowRevert bool, caption string) *RawTransaction {
	return NewRawTransaction(from, tx.Gas(), tx.To(), tx.Value(), tx.Data(), allowRevert, caption)
}

type RawTransactions []*RawTransaction
