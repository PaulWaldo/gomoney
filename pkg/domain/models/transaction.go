package models

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	Payee     string
	Type      string
	Amount    float64
	Memo      string
	Date      time.Time
	AccountID uint
	// TransferAccountId uint
}
