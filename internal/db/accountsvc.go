package db

import (
	"database/sql"
	"fmt"

	"github.com/PaulWaldo/gomoney/pkg/domain"
	"github.com/PaulWaldo/gomoney/pkg/domain/models"
)

var _ domain.AccountSvc = accountSvc{}

const createColumnNames = "name, type, memo, routing, acct_number, hidden, net_worth_include, budget_include"
const readColumnNames = "account_id, " + createColumnNames

type accountSvc struct {
	db *sql.DB
}

func NewAccountSvc(db *sql.DB) domain.AccountSvc {
	return accountSvc{db: db}
}

func (as accountSvc) Create(account *models.Account) error {
	err := WithTransaction(as.db, func(tx Transaction) error {
		query := fmt.Sprintf("INSERT INTO accounts(%s) VALUES($1, $2, $3, $4, $5, $6, $7, $8)", createColumnNames)
		stmt, err := tx.Prepare(query)
		if err != nil {
			return err
		}
		defer stmt.Close()
		res, err := stmt.Exec(
			account.Name,
			account.Type,
			account.Memo,
			account.Routing,
			account.AccountNumber,
			account.Hidden,
			account.NetWorthInclude,
			account.BudgetInclude,
		)
		if err != nil {
			return err
		}
		id, err := res.LastInsertId()
		if err != nil {
			return err
		}
		account.ID = id
		rows, err := res.RowsAffected()
		if err != nil {
			return err
		}

		fmt.Printf("Create accound ID %d, affected rows: %d\n", id, rows)
		return nil
	})
	return err
}

// const selectStmt = `
// 		SELECT
// 			account_id,
// 			name,
// 			type,
// 			memo,
// 			acct_number,
// 			hidden,
// 			net_worth_include,
// 			budget_include
// 		FROM accounts
// `

func (as accountSvc) Get(id int64) (models.Account, error) {
	cmd := fmt.Sprintf("SELECT %s FROM accounts WHERE account_id = ?", readColumnNames)
	stmt, err := as.db.Prepare(cmd)
	if err != nil {
		return models.Account{}, err
	}
	var got models.Account
	err = stmt.QueryRow(id).Scan(
		&got.ID,
		&got.Name,
		&got.Type,
		&got.Memo,
		&got.Routing,
		&got.AccountNumber,
		&got.Hidden,
		&got.NetWorthInclude,
		&got.BudgetInclude,
	)
	if err != nil {
		return models.Account{}, err
	}
	return got, nil
}

func (as accountSvc) List() ([]models.Account, error) {
	cmd := fmt.Sprintf("SELECT %s FROM accounts", readColumnNames)
	rows, err := as.db.Query(cmd)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var accounts []models.Account
	for rows.Next() {
		var a models.Account
		err = rows.Scan(
			&a.ID,
			&a.Name,
			&a.Type,
			&a.Memo,
			&a.Routing,
			&a.AccountNumber,
			&a.Hidden,
			&a.NetWorthInclude,
			&a.BudgetInclude,
		)
		if err != nil {
			return nil, err
		}
		accounts = append(accounts, a)
	}
	return accounts, nil
}

func (as accountSvc) AddTransactions(a models.Account, transactions []models.Transaction) error {
	return nil
	// a.Transactions = append(a.Transactions, transactions...)
	// err := as.Save(a).Error
	// return err
}

func (as accountSvc) Update(a *models.Account) error {
	err := WithTransaction(as.db, func(tx Transaction) error {
		query := `UPDATE accounts SET name = $1, type = $2 WHERE account_id = $3`
		stmt, err := tx.Prepare(query)
		if err != nil {
			return err
		}
		defer stmt.Close()
		_, err = stmt.Exec(a.Name, a.Type, a.ID)
		if err != nil {
			return err
		}
		return nil
	})
	return err
}
