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

func connectToDatabase() (*gorm.DB, error) {
	// In-memory sqlite if no database name is specified
	dsn := "file::memory:?cache=shared"
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&models.Account{})
	return db, nil
}


func NewSqliteInMemoryServices() (*Services, error) {
	db, err := connectToDatabase()
	if err != nil {
		return nil, err
	}
	s := &Services {Account: NewAccountSvc(db)}
	return s, nil
}