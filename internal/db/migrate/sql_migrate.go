package migrate

import (
	"database/sql"
	"fmt"
	"io"
	"os"
)

// Adapted from https://github.com/joncalhoun/migrate and https://www.calhoun.io/database-migrations-in-go/

type SQLMigration struct {
	ID       string
	Migrate  func(tx *sql.Tx) error
	Rollback func(tx *sql.Tx) error
}

func NewQueryMigration(id, upQuery, downQuery string) SQLMigration {
	queryFn := func(query string) func(tx *sql.Tx) error {
		if query == "" {
			return nil
		}
		return func(tx *sql.Tx) error {
			_, err := tx.Exec(query)
			return err
		}
	}
	m := SQLMigration{
		ID:       id,
		Migrate:  queryFn(upQuery),
		Rollback: queryFn(downQuery),
	}
	return m
}

func NewFileMigration(dir, id, upFile, downFile string) SQLMigration {
	fileFn := func(filename string) func(*sql.Tx) error {
		if filename == "" {
			return nil
		}
		f, err := os.Open(fmt.Sprintf("%s/%s", dir, filename))
		if err != nil {
			panic(err)
		}
		fileBytes, err := io.ReadAll(f)
		if err != nil {
			panic(err)
		}
		return func(tx *sql.Tx) error {
			_, err := tx.Exec(string(fileBytes))
			return err
		}
	}
	m := SQLMigration{
		ID:       id,
		Migrate:  fileFn(upFile),
		Rollback: fileFn(downFile),
	}
	return m
}

type Migrator struct {
	Migrations []SQLMigration
}

func (m Migrator) Migrate(db *sql.DB) error {
	if err := m.createMigrationTable(db); err != nil {
		return err
	}
	for _, migration := range m.Migrations {
		var id string
		err := db.QueryRow("SELECT id FROM migrations WHERE id=$1", migration.ID).Scan(&id)
		switch err {
		case sql.ErrNoRows:
			fmt.Printf("Running migration: %v\n", migration.ID)
			// we need to run the migration so we continue to code below
		case nil:
			fmt.Printf("Skipping migration: %v\n", migration.ID)
			continue
		default:
			return fmt.Errorf("looking up migration by id: %w", err)
		}
		err = m.runMigration(db, migration)
		if err != nil {
			return err
		}
	}
	return nil
}

func (m Migrator) createMigrationTable(db *sql.DB) error {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS migrations (id TEXT PRIMARY KEY)")
	if err != nil {
		return fmt.Errorf("creating migration table: %w", err)
	}
	return nil
}

func (m Migrator) runMigration(db *sql.DB, migration SQLMigration) error {
	errorf := func(err error) error { return fmt.Errorf("running migration: %w", err) }

	tx, err := db.Begin()
	if err != nil {
		return errorf(err)
	}
	_, err = tx.Exec("INSERT INTO migrations (id) VALUES ($1)", migration.ID)
	if err != nil {
		tx.Rollback()
		return errorf(err)
	}
	err = migration.Migrate(tx)
	if err != nil {
		tx.Rollback()
		return errorf(err)
	}
	err = tx.Commit()
	if err != nil {
		return errorf(err)
	}
	return nil
}
