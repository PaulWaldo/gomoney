package mocks

import "github.com/PaulWaldo/gomoney/pkg/domain/models"

type TransactionSvc struct {
	CreateErr error
	GetResp   models.Transaction
	GetErr    error
	ListResp  []models.Transaction
	ListErr   error
}

func (ts TransactionSvc) Create(transaction *models.Transaction) error {
	return ts.CreateErr
}

func (ts TransactionSvc) Get(id uint) (models.Transaction, error) {
	return ts.GetResp, ts.GetErr
}

func (ts TransactionSvc) List() ([]models.Transaction, error) {
	return ts.ListResp, ts.ListErr
}
