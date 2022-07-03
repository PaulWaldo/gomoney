package mocks

import "github.com/PaulWaldo/gomoney/pkg/domain/models"

type AccountSvc struct {
	GetAccountResp      models.Account
	GetAccountErr       error
	ListAccountResp     []models.Account
	ListAccountErr      error
	CreateAccountErr    error
	DeleteAccountErr    error
	AddTransactinsError error
}

func (as AccountSvc) Create(account *models.Account) error {
	return as.CreateAccountErr
}

func (as AccountSvc) Get(id uint) (models.Account, error) {
	return as.GetAccountResp, as.GetAccountErr
}

func (as AccountSvc) List() ([]models.Account, error) {
	return as.ListAccountResp, as.ListAccountErr
}

func (as AccountSvc) AddTransactions(a models.Account, t []models.Transaction) error {
	return as.AddTransactinsError
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
