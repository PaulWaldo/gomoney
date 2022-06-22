package mocks

import (
	"github.com/PaulWaldo/gomoney/internal/db/models"
)

type AccountSvc struct {
	GetAccountResp    *models.Account
	GetAccountErr     error
	ListAccountResp   []*models.Account
	ListAccountErr    error
	CreateAccountResp uint
	CreateAccountErr  error
	DeleteAccountErr  error
}

func (as AccountSvc) Create(name string, accountType models.AccountType) (uint, error) {
	return as.CreateAccountResp, as.CreateAccountErr
}

func (as AccountSvc) Get(id uint) (*models.Account, error) {
	return as.GetAccountResp, as.GetAccountErr
}

func (as AccountSvc) List() ([]*models.Account, error) {
	return as.ListAccountResp, as.ListAccountErr
}

type AccountDB struct {
	GetAccountResp    *models.Account
	GetAccountErr     error
	ListAccountResp   []*models.Account
	ListAccountErr    error
	CreateAccountResp uint
	CreateAccountErr  error
	DeleteAccountErr  error
}

func (ad AccountDB) Create(name string, accountType models.AccountType) (uint, error) {
	return ad.CreateAccountResp, ad.CreateAccountErr
}

func (as AccountDB) Get(id uint) (*models.Account, error) {
	return as.GetAccountResp, as.GetAccountErr
}

func (as AccountDB) List() ([]*models.Account, error) {
	return as.ListAccountResp, as.ListAccountErr
}
