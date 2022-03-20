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
