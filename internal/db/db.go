package db

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/PaulWaldo/gomoney/pkg/domain"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func connectToDatabase(dsn string, gormConfig *gorm.Config) (*gorm.DB, error) {
	// In-memory sqlite if no database name is specified
	db, err := gorm.Open(sqlite.Open(dsn), gormConfig)
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&domain.Account{}, &domain.Transaction{})
	return db, nil
}

func populateDatabaseLarge(services domain.Services) error {
	const numModels = 20
	const numTransactions = 1000
	now := time.Now()
	day := 24 * time.Hour

	as := services.Account
	// accounts := make([]Account, numModels)
	for i := 0; i < numModels; i++ {
		transactions := make([]domain.Transaction, numTransactions)
		for j := 0; j < numTransactions; j++ {
			transactions[j] = domain.Transaction{
				Payee:  fmt.Sprintf("Transaction %d, Account %d", j+1, i+1),
				Type:   "W",
				Amount: float64(rand.Float32()) * 10000,
				Memo:   fmt.Sprintf("For stuff %d", j+1),
				Date:   now.Add(time.Duration(-(rand.Int31n(200))) * day),
			}
		}
		account := domain.Account{
			Name:         fmt.Sprintf("Account %d", i+1),
			Type:         domain.Checking.Slug,
			Transactions: transactions,
		}
		if err := as.Create(&account); err != nil {
			return err
		}
	}
	return nil
}

func populateDatabaseSmall(services domain.Services) error {
	as := services.Account
	var err error
	now := time.Now()
	day := 24 * time.Hour
	err = as.Create(&domain.Account{
		Name: "My Checking",
		Type: domain.Checking.Slug,
		Transactions: []domain.Transaction{
			{
				Payee:  "Grocery store",
				Type:   "W",
				Amount: 52.6,
				Memo:   "The usual suspects",
				Date:   now.Add(-3 * day),
			},
			{
				Payee:  "Liquor store",
				Type:   "W",
				Amount: 359,
				Memo:   "For the big party",
				Date:   now.Add(-2 * day),
			},
		},
	})
	if err != nil {
		return err
	}
	err = as.Create(&domain.Account{
		Name: "My Savings",
		Type: domain.Savings.Slug,
		Transactions: []domain.Transaction{
			{
				Payee:  "Me",
				Type:   "D",
				Amount: 500,
				Memo:   "Save for rainy day",
				Date:   now.Add(-1 * day),
			},
			{
				Payee:  "Me",
				Type:   "D",
				Amount: 500,
				Memo:   "Save for rainy day",
				Date:   now.Add(-30 * day),
			},
		},
	})
	if err != nil {
		return err
	}
	err = as.Create(&domain.Account{Name: "My Credit Card", Type: domain.CreditCard.Slug})
	if err != nil {
		return err
	}
	return nil
}

func NewSqliteInMemoryServices(gormConfig *gorm.Config, createDummyData bool) (*domain.Services, *gorm.DB, error) {
	db, err := connectToDatabase("file::memory:?cache=shared", gormConfig)
	if err != nil {
		return nil, nil, err
	}
	s := &domain.Services{Account: NewAccountSvc(db), Transaction: NewTransactionSvc(db)}
	if createDummyData {
		err = populateDatabaseLarge(*s)
		if err != nil {
			return nil, nil, err
		}
	}
	return s, db, nil
}
