package domain

import (
	"time"

	"gorm.io/gorm"
)

type AccountSvc interface {
	Create(account *Account) error
	Get(id uint) (Account, error)
	List() ([]Account, error)
	AddTransactions(a Account, t []Transaction) error
}

type Transaction struct {
	Payee     string    `json:"payee,omitempty"`
	Type      string    `json:"type,omitempty"`
	Amount    float64   `json:"amount,omitempty"`
	Memo      string    `json:"memo,omitempty"`
	Date      time.Time `json:"date,omitempty"`
	AccountID uint      `json:"account_id,omitempty"`
}

type TransactionSvc interface {
	Create(transaction *Transaction) error
	Get(id uint) (Transaction, error)
	List() ([]Transaction, int64, error)
	SetPaginationScope(scope func(*gorm.DB) *gorm.DB)
}

type Services struct {
	Account     AccountSvc
	Transaction TransactionSvc
}
