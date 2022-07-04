package db

import (
	"github.com/PaulWaldo/gomoney/pkg/domain"
	"github.com/PaulWaldo/gomoney/pkg/domain/models"
	"gorm.io/gorm"
)

type accountSvc struct {
	db *gorm.DB
}

func NewAccountSvc(db *gorm.DB) domain.AccountSvc {
	return accountSvc{db: db}
}

func (as accountSvc) Create(account *models.Account) error {
	res := as.db.Create(&account)
	return res.Error
}

func (as accountSvc) Get(id uint) (models.Account, error) {
	var acct models.Account
	res := as.db.Preload("Transactions").First(&acct, id)
	return acct, res.Error
}

func (as accountSvc) List() ([]models.Account, error) {
	var accounts []models.Account
	res := as.db.Preload("Transactions").Find(&accounts)
	return accounts, res.Error
}
func (as accountSvc) AddTransactions(a models.Account, transactions []models.Transaction) error {
	a.Transactions = append(a.Transactions, transactions...)
	err := as.db.Save(a).Error
	return err
}
