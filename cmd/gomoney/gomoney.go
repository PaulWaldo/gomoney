package main

import (
	// "log"

	"github.com/PaulWaldo/gomoney/internal/db"
	"github.com/PaulWaldo/gomoney/ui"
	"gorm.io/gorm"
)

func main() {
	services, _, err := db.NewSqliteInMemoryServices(&gorm.Config{}, true)
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
	appData := ui.AppData{Accounts: accounts, Transactions: transactions}
	appData.RunApp()
}
