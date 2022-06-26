package db

import (
	"fmt"
	"os"
	"testing"

	"github.com/PaulWaldo/gomoney/pkg/domain/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}

func setup() {
	var err error
	db, err = gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{SkipDefaultTransaction: true})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&models.Account{}, &models.Transaction{})
	fmt.Printf("\033[1;36m%s\033[0m", "> Setup completed\n")
}

func teardown() {
	// Do something here.

	fmt.Printf("\033[1;36m%s\033[0m", "> Teardown completed")
	fmt.Printf("\n")
}
