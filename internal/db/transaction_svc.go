package db

import (
	"github.com/PaulWaldo/gomoney/pkg/domain"
	"github.com/PaulWaldo/gomoney/pkg/domain/models"
	"gorm.io/gorm"
)

type transactionSvc struct {
	paginationScope func(*gorm.DB) *gorm.DB
	db              *gorm.DB
}

func NewTransactionSvc(db *gorm.DB) domain.TransactionSvc {
	return &transactionSvc{db: db}
}

func (ts *transactionSvc) SetPaginationScope(scope func(*gorm.DB) *gorm.DB) {
	ts.paginationScope = scope
}

func (ts *transactionSvc) Create(transaction *models.Transaction) error {
	res := ts.db.Create(&transaction)
	return res.Error
}

func (ts transactionSvc) Get(id uint) (models.Transaction, error) {
	var t models.Transaction
	res := ts.db.First(&t, id)
	return t, res.Error
}

// func convertTransactionsToAny(t []models.Transaction) []interface{} {
// 	a := make([]interface{}, len(t))
// 	for i, v := range t {
// 		a[i] = v
// 	}
// 	return a
// }

func (ts transactionSvc) ListByAccount(accountId uint) ([]models.Transaction, int64, error) {
	var txs []models.Transaction
	paginatedDb := ts.db
	if ts.paginationScope != nil {
		paginatedDb = ts.db.Scopes(ts.paginationScope)
	}
	var count int64
	// paginatedDb = paginatedDb.Debug()
	paginatedDb = paginatedDb.Where(&models.Transaction{AccountID: accountId})
	err := paginatedDb.Find(&txs).Offset(-1).Limit(-1).Count(&count).Error
	return txs, count, err
}

func (ts transactionSvc) List() ([]models.Transaction, int64, error) {
	var txs []models.Transaction
	paginatedDb := ts.db
	if ts.paginationScope != nil {
		paginatedDb = ts.db.Scopes(ts.paginationScope)
	}
	var count int64
	err := paginatedDb.Find(&txs).Offset(-1).Limit(-1).Count(&count).Error
	return txs, count, err
}

// func (ts transactionSvc) AddTransactions(a models.Account, transactions []models.Transaction) error {
// 	a.Transactions = append(a.Transactions, transactions...)
// 	err := ts.db.Save(a).Error
// 	return err
// }
