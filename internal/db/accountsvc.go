package db

import (
	"github.com/PaulWaldo/gomoney/pkg/domain"
	"gorm.io/gorm"
)

type accountSvc struct {
	db *gorm.DB
}

func NewAccountSvc(db *gorm.DB) domain.AccountSvc {
	return accountSvc{db: db}
}

func (as accountSvc) Create(account *domain.Account) error {
	res := as.db.Create(&account)
	return res.Error
}

func (as accountSvc) Get(id uint) (domain.Account, error) {
	var acct domain.Account
	res := as.db.Preload("Transactions").First(&acct, id)
	return acct, res.Error
}

func (as accountSvc) List() ([]domain.Account, error) {
	var accounts []domain.Account
	res := as.db.Preload("Transactions").Find(&accounts)
	return accounts, res.Error
}
func (as accountSvc) AddTransactions(a domain.Account, transactions []domain.Transaction) error {
	a.Transactions = append(a.Transactions, transactions...)
	err := as.db.Save(a).Error
	return err
}
