package db

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/PaulWaldo/gomoney/pkg/domain"
	"github.com/PaulWaldo/gomoney/pkg/domain/models"
)

var _ domain.TransactionSvc = transactionSvc{}

type transactionSvc struct {
	db *sql.DB
}

func NewTransactionSvc(db *sql.DB) domain.TransactionSvc {
	return &transactionSvc{db: db}
}

func (ts transactionSvc) Create(t *models.Transaction) error {
	err := WithTransaction(ts.db, func(tx Transaction) error {
		query := `
			INSERT INTO
				transactions(
					payee,
					type,
					amount,
					memo,
					date,
					account_id
				)
			VALUES($1, $2, $3, $4, $5, $6)`
		stmt, err := tx.Prepare(query)
		if err != nil {
			return err
		}
		defer stmt.Close()

		res, err := stmt.Exec(t.Payee, t.Type, t.Amount, t.Memo, t.Date, t.AccountID)
		if err != nil {
			return err
		}
		id, err := res.LastInsertId()
		if err != nil {
			return err
		}
		t.ID = id
		rows, err := res.RowsAffected()
		if err != nil {
			return err
		}

		fmt.Printf("Create transaction ID %d, affected rows: %d\n", id, rows)
		return nil
	})
	return err
}

func (ts transactionSvc) Get(id int64) (models.Transaction, error) {
	var t models.Transaction
	stmt, err := ts.db.Prepare(`
		SELECT
			transaction_id,
			payee,
			type,
			amount,
			memo,
			date,
			account_id
		FROM transactions
		WHERE account_id = ?
	`)
	if err != nil {
		return t, err
	}
	var got models.Transaction
	err = stmt.QueryRow(id).Scan(&got.ID, &got.Payee, &got.Type, &got.Amount, &got.Memo, &got.Date, &got.AccountID)
	if err != nil {
		return t, err
	}
	return got, nil
}

func (ts transactionSvc) ListByAccount(accountId int64) ([]models.Transaction, error) {
	var txs []models.Transaction
	query := `
		SELECT
			transaction_id,
			payee,
			type,
			amount,
			memo,
			date,
			account_id,
			SUM (amount) OVER (
				ORDER BY  date, transaction_id
			) AS balance
		FROM transactions
		WHERE account_id = ?
		ORDER BY date`

	stmt, err := ts.db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(accountId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var got models.Transaction
		err = rows.Scan(&got.ID, &got.Payee, &got.Type, &got.Amount, &got.Memo, &got.Date, &got.AccountID, &got.Balance)
		if err != nil {
			return txs, err
		}
		txs = append(txs, got)
	}
	return txs, nil
}

func (ts transactionSvc) List() ([]models.Transaction, error) {
	var txs []models.Transaction
	rows, err := ts.db.Query(`
		SELECT
			transaction_id,
			payee,
			type,
			amount,
			memo,
			date,
			account_id
		FROM transactions
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var got models.Transaction
		err = rows.Scan(&got.ID, &got.Payee, &got.Type, &got.Amount, &got.Memo, &got.Date, &got.AccountID)
		if err != nil {
			return txs, err
		}
		txs = append(txs, got)
	}
	return txs, nil
}

func (ts transactionSvc) AddTransactions(a models.Account, transactions []models.Transaction) error {
	// a.Transactions = append(a.Transactions, transactions...)
	// err := ts.db.Save(a).Error
	return errors.New("oops")
}

func (ts transactionSvc) Update(t *models.Transaction) error {
	err := WithTransaction(ts.db, func(tx Transaction) error {
		query := `
			UPDATE transactions
			SET
				payee = $1,
				type = $2,
				amount = $3,
				memo = $4,
				date = $5,
				account_id = $6
			WHERE transaction_id = $7
		`
		stmt, err := tx.Prepare(query)
		if err != nil {
			return err
		}
		defer stmt.Close()
		_, err = stmt.Exec(t.Payee, t.Type, t.Amount, t.Memo, t.Date, t.AccountID, t.ID)
		if err != nil {
			return err
		}
		return nil
	})
	return err
}
