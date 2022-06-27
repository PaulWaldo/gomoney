package domain

import "github.com/PaulWaldo/gomoney/pkg/domain/models"

type AccountSvc interface {
	Create(account *models.Account) error
	Get(id uint) (models.Account, error)
	List() ([]models.Account, error)
	AddTransactions(a models.Account, t []models.Transaction) error
}

type TransactionSvc interface {
}

type Services struct {
	Account AccountSvc
}
