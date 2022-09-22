package db

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"database/sql"

	"github.com/PaulWaldo/gomoney/internal/db/migrate"
	"github.com/PaulWaldo/gomoney/pkg/domain"
	"github.com/PaulWaldo/gomoney/pkg/domain/models"
	_ "github.com/mattn/go-sqlite3"
)

var migrator migrate.Migrator

func prepareMigrations(dir string) {
	cwd, _ := os.Getwd()
	fmt.Printf("CWD is %s\n", cwd)
	migrator = migrate.Migrator{Migrations: []migrate.SQLMigration{
		migrate.NewFileMigration(dir, "001_account_transaction", "001_account_transaction.sql", ""),
	}}
}

func connectToDatabase(driverName, dsn, migDir string) (*sql.DB, error) {
	prepareMigrations(migDir)
	// Assumes that dsn has been validated as an existing file
	db, err := sql.Open(driverName, dsn)
	if err != nil {
		return nil, err
	}

	err = migrator.Migrate(db)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func populateDatabaseLarge(services domain.Services) error {
	const numModels = 20
	const numTransactions = 1000
	now := time.Now()
	day := 24 * time.Hour

	as := services.Account
	ts := services.Transaction
	// accounts := make([]models.Account, numModels)
	for i := 0; i < numModels; i++ {
		account := models.Account{
			Name: fmt.Sprintf("Account %d", i+1),
			Type: models.Checking.Slug,
			// Transactions: transactions,
		}
		if err := as.Create(&account); err != nil {
			return err
		}
		transactions := make([]models.Transaction, numTransactions)
		for j := 0; j < numTransactions; j++ {
			transactions[j] = models.Transaction{
				Payee:     fmt.Sprintf("Transaction %d, Account %d", j+1, i+1),
				Type:      "W",
				Amount:    float64(rand.Float32()) * 10000,
				Memo:      fmt.Sprintf("For stuff %d", j+1),
				Date:      now.Add(time.Duration(-(rand.Int31n(200))) * day),
				AccountID: account.ID,
			}
			if err := ts.Create(&(transactions[j])); err != nil {
				return err
			}
		}
	}
	return nil
}

func populateDatabaseSmall(services domain.Services) error {
	as := services.Account
	var err error
	now := time.Now()
	day := 24 * time.Hour
	err = as.Create(&models.Account{
		Name: "My Checkingxxx",
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

func enforcesForeignKeys(db *sql.DB) (bool, error) {
	var p int
	err := db.QueryRow("PRAGMA foreign_keys;").Scan(&p)
	if err != nil {
		return false, err
	}
	return p == 1, nil
}

func prepareServices(driverName, dsn, migDir string) (*domain.Services, error) {
	db, err := connectToDatabase("sqlite3", dsn, migDir)
	if err != nil {
		return nil, err
	}

	// Check for Foreign Key support
	enforced, err := enforcesForeignKeys(db)
	if err != nil {
		return nil, err
	}
	if !enforced {
		_, err = db.Exec("PRAGMA foreign_keys=TRUE;")
		if err != nil {
			return nil, err
		}
		enforced, err = enforcesForeignKeys(db)
		if err != nil {
			return nil, err
		}
		if !enforced {
			panic("Foreign Key support is not available.  Refusing to continue")
		}
	}

	s := &domain.Services{
		Db:          db,
		Account:     NewAccountSvc(db),
		Transaction: NewTransactionSvc(db),
	}
	return s, nil
}

func NewSqliteInMemoryServices(migDir string, createDummyData bool) (*domain.Services, error) {
	s, err := prepareServices("sqlite3", ":memory:", migDir)
	if err != nil {
		return nil, err
	}
	if createDummyData {
		err = populateDatabaseLarge(*s)
		if err != nil {
			return nil, err
		}
	}
	return s, nil
}

func NewSqliteDiskServices(dsn, migDir string) (*domain.Services, error) {
	return prepareServices("sqlite3", dsn, migDir)
}
