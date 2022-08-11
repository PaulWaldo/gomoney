package main

import (
	"log"
	"os"
	"time"

	"github.com/PaulWaldo/gomoney/internal/application"
	"github.com/PaulWaldo/gomoney/internal/db"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	useDefaultTransactions := true
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,        // Disable color
		},
	)
	services, _, err := db.NewSqliteInMemoryServices(&gorm.Config{
		SkipDefaultTransaction: true,
		Logger:                 newLogger,
	}, useDefaultTransactions)
	if err != nil {
		panic(err)
	}
	transactions, _, err := services.Transaction.List()
	if err != nil {
		panic(err)
	}
	accounts, err := services.Account.List()
	if err != nil {
		panic(err)
	}
	appData := &application.AppData{Accounts: accounts, Transactions: transactions, Service: *services}
	application.RunApp(appData)
}
