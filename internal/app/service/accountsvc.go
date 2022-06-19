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

func (as accountSvc) Create(name string, accountType domain.AccountType) (domain.AccountIDType, error) {
	acct := domain.Account{Name: name, AccountType: accountType}
	res := as.db.Create(acct)
	return acct.ID, res.Error
}

func (as accountSvc) Get(id domain.AccountIDType) (*models.Account, error) {
	var acct models.Account
	res := as.db.First(&acct, id)
	return &acct, res.Error
}

func (as accountSvc) List() ([]*models.Account, error) {
	return []*models.Account{}, nil
}
