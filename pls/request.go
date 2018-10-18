// dummy requests transactions for included ORB, URB

package pls

import (
	"container/heap"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/big"
	"sync/atomic"

	"github.com/Onther-Tech/plasma-evm/common"
	"github.com/Onther-Tech/plasma-evm/common/hexutil"
	"github.com/Onther-Tech/plasma-evm/crypto"
	"github.com/Onther-Tech/plasma-evm/crypto/sha3"
	"github.com/Onther-Tech/plasma-evm/rlp"
)

var (
	ErrInvalidSig = errors.New("invalid transaction v, r, s values")
)

type RequestTransaction struct {
	data txdata
	// caches
	hash atomic.Value
	size atomic.Value
	from atomic.Value
}

type txdata struct {
	AccountNonce uint64          `json:"nonce"    gencodec:"required"`
	Price        *big.Int        `json:"gasPrice" gencodec:"required"`
	GasLimit     uint64          `json:"gas"      gencodec:"required"`
	Recipient    *common.Address `json:"to"       rlp:"nil"` // nil means contract creation
	Amount       *big.Int        `json:"value"    gencodec:"required"`
	Payload      []byte          `json:"input"    gencodec:"required"`

	// Signature values
	V *big.Int `json:"v" gencodec:"required"`
	R *big.Int `json:"r" gencodec:"required"`
	S *big.Int `json:"s" gencodec:"required"`

	// This is only used when marshaling to JSON.
	Hash *common.Hash `json:"hash" rlp:"-"`
}

func NewRequest(nonce uint64, to common.Address, amount *big.Int, gasLimit uint64, gasPrice *big.Int, data []byte) *RequestTransaction {
	return newRequest(nonce, &to, amount, gasLimit, gasPrice, data)
}

func newRequest(nonce uint64, to *common.Address, amount *big.Int, gasLimit uint64, gasPrice *big.Int, data []byte) *RequestTransaction {
	if len(data) > 0 {
		data = common.CopyBytes(data)
	}
	d := txdata{
		AccountNonce: 0,
		Recipient:    to,
		Payload:      data,
		Amount:       new(big.Int),
		GasLimit:     gasLimit,
		Price:        new(big.Int), // It should be Ninimum Gas Price in place on PlasmaChain
		V:            common.Big0,
		R:            common.Big0,
		S:            common.Big0,
	}
	if amount != nil {
		d.Amount.Set(amount)
	}
	if gasPrice != nil {
		d.Price.Set(gasPrice)
	}

	return &RequestTransaction{data: d}
}


// ChainId returns which chain id this transaction was signed for (if at all)
func (req *RequestTransaction) ChainId() *big.Int {
	return deriveChainId(req.data.V)
}

// deriveChainId derives the chain id from the given v parameter
func deriveChainId(v *big.Int) *big.Int {
	if v.BitLen() <= 64 {
		v := v.Uint64()
		if v == 27 || v == 28 {
			return new(big.Int)
		}
		return new(big.Int).SetUint64((v - 35) / 2)
	}
	v = new(big.Int).Sub(v, big.NewInt(35))
	return v.Div(v, big.NewInt(2))
}

// Protected returns whether the transaction is protected from replay protection.
func (req *RequestTransaction) Protected() bool {
	return isProtectedV(req.data.V)
}

func isProtectedV(V *big.Int) bool {
	if V.BitLen() <= 8 {
		v := V.Uint64()
		return v != 27 && v != 28
	}
	// anything not 27 or 28 is considered protected
	return true
}

// EncodeRLP implements rlp.Encoder
func (req *RequestTransaction) EncodeRLP(w io.Writer) error {
	return rlp.Encode(w, &req.data)
}

// DecodeRLP implements rlp.Decoder
func (req *RequestTransaction) DecodeRLP(s *rlp.Stream) error {
	_, size, _ := s.Kind()
	err := s.Decode(&req.data)
	if err == nil {
		req.size.Store(common.StorageSize(rlp.ListSize(size)))
	}

	return err
}

// MarshalJSON encodes the web3 RPC transaction format.
func (req *RequestTransaction) MarshalJSON() ([]byte, error) {
	hash := req.Hash()
	data := req.data
	data.Hash = &hash
	return data.MarshalJSON()
}

// UnmarshalJSON decodes the web3 RPC transaction format.
func (req *RequestTransaction) UnmarshalJSON(input []byte) error {
	var dec txdata
	if err := dec.UnmarshalJSON(input); err != nil {
		return err
	}
	var V byte
	if isProtectedV(dec.V) {
		chainID := deriveChainId(dec.V).Uint64()
		V = byte(dec.V.Uint64() - 35 - 2*chainID)
	} else {
		V = byte(dec.V.Uint64() - 27)
	}
	if !crypto.ValidateSignatureValues(V, dec.R, dec.S, false) {
		return ErrInvalidSig
	}
	*req = RequestTransaction{data: dec}
	return nil
}

func (req *RequestTransaction) Data() []byte       { return common.CopyBytes(req.data.Payload) }
func (req *RequestTransaction) Gas() uint64        { return req.data.GasLimit }
func (req *RequestTransaction) GasPrice() *big.Int { return new(big.Int).Set(req.data.Price) }
func (req *RequestTransaction) Value() *big.Int    { return new(big.Int).Set(req.data.Amount) }
func (req *RequestTransaction) Nonce() uint64      { return req.data.AccountNonce }
func (req *RequestTransaction) CheckNonce() bool   { return true }

// To returns the recipient address of the transaction.
// It returns nil if the transaction is a contract creation.
func (req *RequestTransaction) To() *common.Address {
	if req.data.Recipient == nil {
		return nil
	}
	to := *req.data.Recipient
	return &to
}

// Hash hashes the RLP encoding of tx.
// It uniquely identifies the transaction.
func (req *RequestTransaction) Hash() common.Hash {
	if hash := req.hash.Load(); hash != nil {
		return hash.(common.Hash)
	}
	v := rlpHash(req)
	req.hash.Store(v)
	return v
}

func rlpHash(x interface{}) (h common.Hash) {
	hw := sha3.NewKeccak256()
	rlp.Encode(hw, x)
	hw.Sum(h[:0])
	return h
}

// Set Null Address if nil num, will be error Address
func (req *RequestTransaction) From() error {
	if hash := req.from.Load(); hash != nil {
		return fmt.Errorf("This is Not Nil, address is : %s.", hash)
	}
	req.from.Store(common.NullAddress)
	return nil
}

// Size returns the true RLP encoded storage size of the transaction, either by
// encoding and returning it, or returning a previsouly cached value.
func (req *RequestTransaction) Size() common.StorageSize {
	if size := req.size.Load(); size != nil {
		return size.(common.StorageSize)
	}
	c := writeCounter(0)
	rlp.Encode(&c, &req.data)
	req.size.Store(common.StorageSize(c))
	return common.StorageSize(c)
}

type writeCounter common.StorageSize

func (c *writeCounter) Write(b []byte) (int, error) {
	*c += writeCounter(len(b))
	return len(b), nil
}

// AsMessage returns the transaction as a core.Message.
//
// AsMessage requires a signer to derive the sender.
//
// XXX Rename message to something less arbitrary?
func (req *RequestTransaction) AsMessage() (Message, error) {
	msg := Message{
		nonce:      req.data.AccountNonce,
		gasLimit:   req.data.GasLimit,
		gasPrice:   new(big.Int).Set(req.data.Price),
		to:         req.data.Recipient,
		amount:     req.data.Amount,
		data:       req.data.Payload,
		checkNonce: true,
	}

	var err error
	msg.from, err = common.NullAddress, nil
	return msg, err
}

// WithSignature returns a new transaction with the given signature.
// This signature needs to be formatted as described in the yellow paper (v+27).
func (req *RequestTransaction) WithSignature() (*RequestTransaction, error) {
	r, s, v := req.data.R, req.data.S, req.data.V
	cpy := &RequestTransaction{data: req.data}
	cpy.data.R, cpy.data.S, cpy.data.V = r, s, v
	return cpy, nil
}

// Cost returns amount + gasprice * gaslimit.
func (req *RequestTransaction) Cost() *big.Int {
	total := new(big.Int).Mul(req.data.Price, new(big.Int).SetUint64(req.data.GasLimit))
	total.Add(total, req.data.Amount)
	return total
}

func (req *RequestTransaction) RawSignatureValues() (*big.Int, *big.Int, *big.Int) {
	return req.data.V, req.data.R, req.data.S
}

//Equals tests for equality of two Contents
func (q *RequestTransaction) Equals(other *RequestTransaction) (bool, error) {
	return q == other, nil
}

func (q *RequestTransaction) GetRlp() []byte {
	enc, _ := rlp.EncodeToBytes(q)
	return enc
}


// Transactions is a Transaction slice type for basic sorting.
type  RequestTransactions []*RequestTransaction

// Len returns the length of s.
func (reqs RequestTransactions) Len() int { return len(reqs) }

// Swap swaps the i'th and the j'th element in s.
func (reqs RequestTransactions) Swap(i, j int) { reqs[i], reqs[j] = reqs[j], reqs[i] }

// GetRlp implements Rlpable and returns the i'th element of s in rlp.
func (reqs RequestTransactions) GetRlp(i int) []byte {
	enc, _ := rlp.EncodeToBytes(reqs[i])
	return enc
}

// TxDifference returns a new set which is the difference between a and b.
func TxDifference(a, b RequestTransactions) RequestTransactions {
	keep := make(RequestTransactions, 0, len(a))

	remove := make(map[common.Hash]struct{})
	for _, tx := range b {
		remove[tx.Hash()] = struct{}{}
	}

	for _, tx := range a {
		if _, ok := remove[tx.Hash()]; !ok {
			keep = append(keep, tx)
		}
	}

	return keep
}

// TxByNonce implements the sort interface to allow sorting a list of transactions
// by their nonces. This is usually only useful for sorting transactions from a
// single account, otherwise a nonce comparison doesn't make much sense.
type RqByNonce RequestTransactions

func (s RqByNonce) Len() int           { return len(s) }
func (s RqByNonce) Less(i, j int) bool { return s[i].data.AccountNonce < s[j].data.AccountNonce }
func (s RqByNonce) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

// TxByPrice implements both the sort and the heap interface, making it useful
// for all at once sorting as well as individually adding and removing elements.
type RqByPrice RequestTransactions

func (s RqByPrice) Len() int           { return len(s) }
func (s RqByPrice) Less(i, j int) bool { return s[i].data.Price.Cmp(s[j].data.Price) > 0 }
func (s RqByPrice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func (s *RqByPrice) Push(x interface{}) {
	*s = append(*s, x.(*RequestTransaction))
}

func (s *RqByPrice) Pop() interface{} {
	old := *s
	n := len(old)
	x := old[n-1]
	*s = old[0 : n-1]
	return x
}

// Modified None singer

// TransactionsByPriceAndNonce represents a set of transactions that can return
// transactions in a profit-maximizing sorted order, while supporting removing
// entire batches of transactions for non-executable accounts.
type TransactionsByPriceAndNonce struct {
	txs    map[common.Address]RequestTransactions // Per account nonce-sorted list of transactions
	heads  RqByPrice                              // Next transaction for each unique account (price heap)
}

// NewTransactionsByPriceAndNonce creates a transaction set that can retrieve
// price sorted transactions in a nonce-honouring way.
//
// Note, the input map is reowned so the caller should not interact any more with
// if after providing it to the constructor.
func NewTransactionsByPriceAndNonce(txs map[common.Address]RequestTransactions) *TransactionsByPriceAndNonce {
	// Initialize a price based heap with the head transactions
	heads := make(RqByPrice, 0, len(txs))
	for _, accTxs := range txs {
		heads = append(heads, accTxs[0])
	}
	// Just Passing input, no process in this method
	return &TransactionsByPriceAndNonce{
		txs:    txs,
		heads:  heads,
	}
}

// Peek returns the next transaction by price.
func (t *TransactionsByPriceAndNonce) Peek() *RequestTransaction {
	if len(t.heads) == 0 {
		return nil
	}
	return t.heads[0]
}

// Shift replaces the current best head with the next one from the same account.
func (t *TransactionsByPriceAndNonce) Shift() {
	// acc is always NullAddress
	acc := common.NullAddress
	if txs, ok := t.txs[acc]; ok && len(txs) > 0 {
		t.heads[0], t.txs[acc] = txs[0], txs[1:]
		heap.Fix(&t.heads, 0)
	} else {
		heap.Pop(&t.heads)
	}
}

// Pop removes the best transaction, *not* replacing it with the next one from
// the same account. This should be used when a transaction cannot be executed
// and hence all subsequent ones should be discarded from the same account.
func (t *TransactionsByPriceAndNonce) Pop() {
	heap.Pop(&t.heads)
}

// Message is a fully derived transaction and implements core.Message
//
// NOTE: In a future PR this will be removed.
type Message struct {
	to         *common.Address
	from       common.Address
	nonce      uint64
	amount     *big.Int
	gasLimit   uint64
	gasPrice   *big.Int
	data       []byte
	checkNonce bool
}

func NewMessage(from common.Address, to *common.Address, nonce uint64, amount *big.Int, gasLimit uint64, gasPrice *big.Int, data []byte, checkNonce bool) Message {
	return Message{
		from:       from,
		to:         to,
		nonce:      nonce,
		amount:     amount,
		gasLimit:   gasLimit,
		gasPrice:   gasPrice,
		data:       data,
		checkNonce: checkNonce,
	}
}

func (m Message) From() common.Address { return m.from }
func (m Message) To() *common.Address  { return m.to }
func (m Message) GasPrice() *big.Int   { return m.gasPrice }
func (m Message) Value() *big.Int      { return m.amount }
func (m Message) Gas() uint64          { return m.gasLimit }
func (m Message) Nonce() uint64        { return m.nonce }
func (m Message) Data() []byte         { return m.data }
func (m Message) CheckNonce() bool     { return m.checkNonce }


// Marshaling method implemented from `gen_to_json.go`
//
var _ = (*txdataMarshaling)(nil)

type txdataMarshaling struct {
	AccountNonce hexutil.Uint64
	Price        *hexutil.Big
	GasLimit     hexutil.Uint64
	Amount       *hexutil.Big
	Payload      hexutil.Bytes
	V            *hexutil.Big
	R            *hexutil.Big
	S            *hexutil.Big
}


func (t txdata) MarshalJSON() ([]byte, error) {
	type txdata struct {
		AccountNonce hexutil.Uint64  `json:"nonce"    gencodec:"required"`
		Price        *hexutil.Big    `json:"gasPrice" gencodec:"required"`
		GasLimit     hexutil.Uint64  `json:"gas"      gencodec:"required"`
		Recipient    *common.Address `json:"to"       rlp:"nil"`
		Amount       *hexutil.Big    `json:"value"    gencodec:"required"`
		Payload      hexutil.Bytes   `json:"input"    gencodec:"required"`
		V            *hexutil.Big    `json:"v" gencodec:"required"`
		R            *hexutil.Big    `json:"r" gencodec:"required"`
		S            *hexutil.Big    `json:"s" gencodec:"required"`
		Hash         *common.Hash    `json:"hash" rlp:"-"`
	}
	var enc txdata
	enc.AccountNonce = hexutil.Uint64(t.AccountNonce)
	enc.Price = (*hexutil.Big)(t.Price)
	enc.GasLimit = hexutil.Uint64(t.GasLimit)
	enc.Recipient = t.Recipient
	enc.Amount = (*hexutil.Big)(t.Amount)
	enc.Payload = t.Payload
	enc.V = (*hexutil.Big)(t.V)
	enc.R = (*hexutil.Big)(t.R)
	enc.S = (*hexutil.Big)(t.S)
	enc.Hash = t.Hash
	return json.Marshal(&enc)
}

func (t *txdata) UnmarshalJSON(input []byte) error {
	type txdata struct {
		AccountNonce *hexutil.Uint64 `json:"nonce"    gencodec:"required"`
		Price        *hexutil.Big    `json:"gasPrice" gencodec:"required"`
		GasLimit     *hexutil.Uint64 `json:"gas"      gencodec:"required"`
		Recipient    *common.Address `json:"to"       rlp:"nil"`
		Amount       *hexutil.Big    `json:"value"    gencodec:"required"`
		Payload      *hexutil.Bytes  `json:"input"    gencodec:"required"`
		V            *hexutil.Big    `json:"v" gencodec:"required"`
		R            *hexutil.Big    `json:"r" gencodec:"required"`
		S            *hexutil.Big    `json:"s" gencodec:"required"`
		Hash         *common.Hash    `json:"hash" rlp:"-"`
	}
	var dec txdata
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.AccountNonce == nil {
		return errors.New("missing required field 'nonce' for txdata")
	}
	t.AccountNonce = uint64(*dec.AccountNonce)
	if dec.Price == nil {
		return errors.New("missing required field 'gasPrice' for txdata")
	}
	t.Price = (*big.Int)(dec.Price)
	if dec.GasLimit == nil {
		return errors.New("missing required field 'gas' for txdata")
	}
	t.GasLimit = uint64(*dec.GasLimit)
	if dec.Recipient != nil {
		t.Recipient = dec.Recipient
	}
	if dec.Amount == nil {
		return errors.New("missing required field 'value' for txdata")
	}
	t.Amount = (*big.Int)(dec.Amount)
	if dec.Payload == nil {
		return errors.New("missing required field 'input' for txdata")
	}
	t.Payload = *dec.Payload
	if dec.V == nil {
		return errors.New("missing required field 'v' for txdata")
	}
	t.V = (*big.Int)(dec.V)
	if dec.R == nil {
		return errors.New("missing required field 'r' for txdata")
	}
	t.R = (*big.Int)(dec.R)
	if dec.S == nil {
		return errors.New("missing required field 's' for txdata")
	}
	t.S = (*big.Int)(dec.S)
	if dec.Hash != nil {
		t.Hash = dec.Hash
	}
	return nil
}