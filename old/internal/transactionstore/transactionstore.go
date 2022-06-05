package transactionstore

import (
	"fmt"
	"sync"
	"time"

	"github.com/PaulWaldo/gomoney/pkg/domain"
)

type TransactionStore struct {
	sync.Mutex
	transactions map[int]domain.Transaction
	nextId       int
}

// New returns a new TransactionStore
func New() *TransactionStore {
	ts := &TransactionStore{}
	ts.transactions = make(map[int]domain.Transaction)
	ts.nextId = 0
	return ts
}

type TransactionCreateRequest struct {
	Payee  string    `json:"payee"`
	Amount float64   `json:"amount"`
	Date   time.Time `json:"date"`
}

type TransactionCreateResponse struct {
	ID string `json:"id"`
}

// CreateTransaction creates a new Transaction in the store
func (ts *TransactionStore) CreateTransaction(req TransactionCreateRequest) int {
	ts.Lock()
	defer ts.Unlock()

	trans := domain.Transaction{
		ID:     ts.nextId,
		Payee:  req.Payee,
		Amount: req.Amount,
		Date:   req.Date,
	}

	ts.transactions[ts.nextId] = trans
	ts.nextId++
	return trans.ID
}

// GetTransaction returns a transaction specified by Id from the store.
// If no such transaction exists, an error is returned
func (ts *TransactionStore) GetTransaction(id int) (domain.Transaction, error) {
	ts.Lock()
	defer ts.Unlock()

	trans, ok := ts.transactions[id]
	if !ok {
		return domain.Transaction{}, fmt.Errorf("transaction with id=%d not found", id)
	}
	return trans, nil
}
