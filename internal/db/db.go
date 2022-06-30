package db

import (
	"time"

	"github.com/PaulWaldo/gomoney/pkg/domain"
	"github.com/PaulWaldo/gomoney/pkg/domain/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func connectToDatabase(dsn string, gormConfig *gorm.Config) (*gorm.DB, error) {
	// In-memory sqlite if no database name is specified
	db, err := gorm.Open(sqlite.Open(dsn), gormConfig)
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&models.Account{}, &models.Transaction{})
	return db, nil
}

func populateDatabase(services domain.Services) error {
	as := services.Account
	var err error
	now := time.Now()
	day := 24 * time.Hour
	err = as.Create(&models.Account{
		Name: "My Checking",
		Type: models.Checking.Slug,
		Transactions: []models.Transaction{
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
	err = as.Create(&models.Account{
		Name: "My Savings",
		Type: models.Savings.Slug,
		Transactions: []models.Transaction{
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
	err = as.Create(&models.Account{Name: "My Credit Card", Type: models.CreditCard.Slug})
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
		err = populateDatabase(*s)
		if err != nil {
			return nil, nil, err
		}
	}
	return s, db, nil
}
