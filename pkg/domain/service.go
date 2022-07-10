package domain

import (
	"github.com/PaulWaldo/gomoney/pkg/domain/models"
	"github.com/PaulWaldo/gomoney/utils"
	"gorm.io/gorm"
)

type AccountSvc interface {
	Create(account *models.Account) error
	Get(id uint) (models.Account, error)
	List() ([]models.Account, error)
	AddTransactions(a models.Account, t []models.Transaction) error
}

type TransactionSvc interface {
	Create(transaction *models.Transaction) error
	Get(id uint) (models.Transaction, error)
	List() (utils.PaginatedResponse, error)
	SetPaginationScope(scope func(*gorm.DB) *gorm.DB)
}

type Services struct {
	Account     AccountSvc
	Transaction TransactionSvc
}
