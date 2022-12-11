package models

import (
	"time"
)

type Transaction struct {
	ID        int64
	Payee     string    `json:"payee,omitempty"`
	Type      string    `json:"type,omitempty"`
	Amount    float64   `json:"amount,omitempty"`
	Memo      string    `json:"memo,omitempty"`
	Date      time.Time `json:"date,omitempty"`
	Balance   float64
	AccountID int64 `json:"account_id,omitempty"`
	// TransferAccountId uint
}
