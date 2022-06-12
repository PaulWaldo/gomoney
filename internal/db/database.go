package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectToDatabase() (db *gorm.DB, err error) {
	// In-memory sqlite if no database name is specified
	dsn := "file::memory:?cache=shared"
	db, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	return db,err
}
