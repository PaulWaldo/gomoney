package domain

import "time"

type Transaction struct {
	ID     int       `json:"id"`
	Payee  string    `json:"payee"`
	Amount float64   `json:"amount"`
	Date   time.Time `json:"date"`
}

type TransactionSvc interface {
	Create(t Transaction) error
	Get(ID int) (Transaction, error)
}

type TransactionDB interface {
	Create(t Transaction) error
	Get(ID int) (Transaction, error)
}
