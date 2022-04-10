package mocks

import "github.com/PaulWaldo/gomoney/pkg/domain"

type AccountSvc struct {
	GetAccountResp    domain.Account
	GetAccountErr     error
	ListAccountResp   []*domain.Account
	ListAccountErr    error
	CreateAccountResp domain.AccountIDType
	CreateAccountErr  error
	DeleteAccountErr  error
}

func (as AccountSvc) Create(a domain.Account) (domain.AccountIDType, error) {
	return as.CreateAccountResp, as.CreateAccountErr
}

func (as AccountSvc) Get(ID domain.AccountIDType) (domain.Account, error) {
	return as.GetAccountResp, as.GetAccountErr
}

type AccountDB struct {
	GetAccountResp    *domain.Account
	GetAccountErr     error
	ListAccountResp   []*domain.Account
	ListAccountErr    error
	CreateAccountResp domain.AccountIDType
	CreateAccountErr  error
	DeleteAccountErr  error
}

func (ad AccountDB) Create(name string, accountType domain.AccountType) (domain.AccountIDType, error) {
	return ad.CreateAccountResp, ad.CreateAccountErr
}

func (as AccountDB) Get(id domain.AccountIDType) (*domain.Account, error) {
	return as.GetAccountResp, as.GetAccountErr
}

func (as AccountDB)List()([]*domain.Account, error) {
	return as.ListAccountResp, as.ListAccountErr
}
