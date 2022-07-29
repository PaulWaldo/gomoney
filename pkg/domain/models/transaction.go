package models

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	ID        uint           `gorm:"primarykey" json:"id,omitempty"`
	CreatedAt time.Time      `json:"created_at,omitempty"`
	UpdatedAt time.Time      `json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
	Payee     string         `json:"payee,omitempty"`
	Type      string         `json:"type,omitempty"`
	Amount    float64        `json:"amount,omitempty"`
	Memo      string         `json:"memo,omitempty"`
	Date      time.Time      `json:"date,omitempty"`
	AccountID uint           `json:"account_id,omitempty"`
	// TransferAccountId uint
}
