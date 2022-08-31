package domain

import (
	"github.com/PaulWaldo/gomoney/pkg/domain/models"
)

type AccountSvc interface {
	Create(account *models.Account) error
	Get(id uint) (models.Account, error)
	List() ([]models.Account, error)
	AddTransactions(models.Account, []models.Transaction) error
	Update(*models.Account) error
}

type TransactionSvc interface {
	Create(transaction *models.Transaction) error
	Get(id uint) (models.Transaction, error)
	List() ([]models.Transaction, error)
	ListByAccount(accountId uint) ([]models.Transaction, error)
	// Update(*models.Account) error
}

type Services struct {
	Account     AccountSvc
	Transaction TransactionSvc
}
