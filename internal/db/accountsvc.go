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

func (as accountSvc) Create(name string, accountType string) (uint, error) {
	acct := models.Account{Name: name, Type: accountType}
	res := as.db.Create(&acct)
	return acct.ID, res.Error
}

func (as accountSvc) Get(id uint) (models.Account, error) {
	var acct models.Account
	res := as.db.First(&acct, id)
	return acct, res.Error
}

func (as accountSvc) List() ([]models.Account, error) {
	var accounts []models.Account
	res := as.db.Find(&accounts)
	return accounts, res.Error
}
