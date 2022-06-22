package service

import (
	"github.com/PaulWaldo/gomoney/internal/db/models"
	"github.com/PaulWaldo/gomoney/pkg/domain"
	"gorm.io/gorm"
)

type accountSvc struct {
	db *gorm.DB
}

func NewAccountSvc(db *gorm.DB) domain.AccountSvc {
	return accountSvc{db: db}
}

func (as accountSvc) Create(name string, accountType models.AccountType) (uint, error) {
	acct := models.Account{
		Name: name, Type: accountType}
	res := as.db.Create(&acct)
	return acct.ID, res.Error
}

func (as accountSvc) Get(id uint) (*models.Account, error) {
	var acct models.Account
	res := as.db.First(&acct, id)
	return &acct, res.Error
}

func (as accountSvc) List() ([]*models.Account, error) {
	return []*models.Account{}, nil
}
