package app

import "github.com/PaulWaldo/gomoney/pkg/domain"

type accountSvc struct {
	db domain.AccountDB
}

func NewAccountSvc(db domain.AccountDB) domain.AccountSvc {
	return accountSvc{db: db}
}

func (as accountSvc) Create(name string, accountType domain.AccountType) (domain.AccountIDType, error) {
	return as.db.Create(name, accountType)
}

func (as accountSvc) Get(id domain.AccountIDType) (*domain.Account, error) {
	return as.db.Get(id)
}

func (as accountSvc) List() ([]*domain.Account, error) {
	return as.db.List()
}