package db

import (
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/PaulWaldo/gomoney/pkg/domain"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB
var services *domain.Services

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}

func setup() {
	var err error
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Warn, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,        // Disable color
		},
	)
	services, db, err = NewSqliteInMemoryServices(&gorm.Config{
		SkipDefaultTransaction: false,
		Logger:                 newLogger,
	}, false)
	if err != nil {
		panic(err)
	}
	fmt.Printf("\033[1;36m%s\033[0m", "> Setup completed\n")
}

func teardown() {
	// Do something here.

	fmt.Printf("\033[1;36m%s\033[0m", "> Teardown completed")
	fmt.Printf("\n")
}

func setupTest(t *testing.T, db *gorm.DB) (teardown func(t *testing.T), tx *gorm.DB) {
	tx = db.Begin()
	teardown = func(t *testing.T) {
		tx.Rollback()
	}
	return teardown, tx
}
