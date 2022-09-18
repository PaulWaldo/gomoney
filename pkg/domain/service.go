package domain

import (
	"database/sql"

	"github.com/PaulWaldo/gomoney/pkg/domain/models"
)

type AccountSvc interface {
	Create(account *models.Account) error
	Get(id int64) (models.Account, error)
	List() ([]models.Account, error)
	AddTransactions(models.Account, []models.Transaction) error
	Update(*models.Account) error
}

type TransactionSvc interface {
	Create(transaction *models.Transaction) error
	Get(id int64) (models.Transaction, error)
	List() ([]models.Transaction, error)
	ListByAccount(accountId int64) ([]models.Transaction, error)
	Update(*models.Transaction) error
}

type Services struct {
	Account     AccountSvc
	Transaction TransactionSvc
	Db          *sql.DB
}
