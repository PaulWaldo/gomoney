package app

import "github.com/PaulWaldo/gomoney/pkg/domain"

type accountSvc struct {
	DB domain.AccountDB
}

func NewAccountSvc(db domain.AccountDB) domain.AccountSvc {
	return accountSvc{DB: db}
}

func (as accountSvc) Create(a domain.Account) (domain.AccountIDType, error) {
	return as.DB.Create(a)
}

func (as accountSvc) Get(ID domain.AccountIDType) (domain.Account, error) {
	return as.DB.Get(ID)
}
