package models

import (
	"github.com/PaulWaldo/gomoney/pkg/domain"
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	domain.Transaction
	// TransferAccountId uint
}
