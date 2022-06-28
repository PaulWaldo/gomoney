package db

import (
	"github.com/PaulWaldo/gomoney/pkg/domain"
	"github.com/PaulWaldo/gomoney/pkg/domain/models"
	"gorm.io/gorm"
)

type transactionSvc struct {
	db *gorm.DB
}

func NewTransactionSvc(db *gorm.DB) domain.TransactionSvc {
	return transactionSvc{db: db}
}

func (ts transactionSvc) Create(transaction *models.Transaction) error {
	res := ts.db.Create(&transaction)
	return res.Error
}

func (ts transactionSvc) Get(id uint) (*models.Transaction, error) {
	var t models.Transaction
	res := ts.db.First(&t, id)
	return &t, res.Error
}

// func (ts transactionSvc) Get(id uint) (models.Account, error) {
// 	var acct models.Account
// 	res := ts.db.First(&acct, id)
// 	return acct, res.Error
// }

// func (ts transactionSvc) List() ([]models.Account, error) {
// 	var accounts []models.Account
// 	res := ts.db.Find(&accounts)
// 	return accounts, res.Error
// }
// func (ts transactionSvc) AddTransactions(a models.Account, transactions []models.Transaction) error {
// 	a.Transactions = append(a.Transactions, transactions...)
// 	err := ts.db.Save(a).Error
// 	return err
// }
