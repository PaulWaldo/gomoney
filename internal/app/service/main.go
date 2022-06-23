package service

import (
	"github.com/PaulWaldo/gomoney/internal/db/models"
	"github.com/PaulWaldo/gomoney/pkg/domain"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Services struct {
	Account domain.AccountSvc
}

func connectToDatabase(dsn string) (*gorm.DB, error) {
	// In-memory sqlite if no database name is specified
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&models.Account{})
	return db, nil
}

func populateDatabase(services Services) error {
	as := services.Account
	var err error
	_, err = as.Create("My Checking", models.Checking.Slug)
	if err != nil {
		return err
	}
	_, err = as.Create("My Savings", models.Savings.Slug)
	if err != nil {
		return err
	}
	_, err = as.Create("My Credit Card", models.CreditCard.Slug)
	if err != nil {
		return err
	}
	return nil
}

func NewSqliteInMemoryServices() (*Services, error) {
	db, err := connectToDatabase("file::memory:?cache=shared")
	if err != nil {
		return nil, err
	}
	s := &Services{Account: NewAccountSvc(db)}
	err = populateDatabase(*s)
	if err != nil {
		return nil, err
	}
	return s, nil
}
