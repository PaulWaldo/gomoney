package migrate_test

import (
	"database/sql"
	"testing"

	"github.com/PaulWaldo/gomoney/internal/db/migrate"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/require"
)

func sqliteInMem(t *testing.T) *sql.DB {
	db, err := sql.Open("sqlite3", ":memory:")
	require.NoError(t, err, "Got error opening inmem database: %s", err)
	t.Cleanup(func() {
		err = db.Close()
		require.NoError(t, err, "Got error closing inmem database: %s", err)
	})
	return db
}

func TestSqlMigrate(t *testing.T) {
	t.Run("single statement migration", func(t *testing.T) {
		db := sqliteInMem(t)
		m := migrate.Migrator{
			Migrations: []migrate.SQLMigration{
				migrate.NewQueryMigration("001_create_widget", createWidgetTableSql, ""),
			}}
		err := m.Migrate(db)
		require.NoError(t, err)
	})

	t.Run("existing migrations", func(t *testing.T) {
		db := sqliteInMem(t)
		migrator := migrate.Migrator{
			Migrations: []migrate.SQLMigration{
				migrate.NewQueryMigration("001_create_widget", createWidgetTableSql, ""),
			},
		}

		//Apply a single migration first
		err := migrator.Migrate(db)
		require.NoError(t, err)
		_, err = db.Exec("INSERT INTO widgets (name) VALUES ($1) ", "existing_test")
		require.NoError(t, err,
			"db.Exec() err = %v; want nil", err)

		// Add another migration to the list
		migrator = migrate.Migrator{
			Migrations: []migrate.SQLMigration{
				migrate.NewQueryMigration("001_create_widget", createWidgetTableSql, ""),
				migrate.NewQueryMigration("002_create_users", createUsersSql, ""),
			},
		}
		err = migrator.Migrate(db)
		require.NoError(t, err, "Migrate() err = %v; want nil", err)
		_, err = db.Exec("INSERT INTO users (email) VALUES ($1) ", "abc@test.com")
		require.NoError(t, err, "db.Exec() err = %v; want nil", err)
	})

	t.Run("file", func(t *testing.T) {
		db := sqliteInMem(t)
		// cwd, _ := os.Getwd()
		// fmt.Printf("CWD is %s\n", cwd)
		migrator := migrate.Migrator{
			Migrations: []migrate.SQLMigration{
				migrate.NewFileMigration("testdata", "001_create_widgets", "users.sql", ""),
			},
		}
		err := migrator.Migrate(db)
		if err != nil {
			t.Fatalf("Migrate() err = %v; want nil", err)
		}
		_, err = db.Exec("INSERT INTO users (name) VALUES ($1)", "fred")
		if err != nil {
			t.Fatalf("db.Exec() err = %v; want nil", err)
		}
	})
	//	t.Run("rollback", func(t *testing.T) {
	//		db := sqliteInMem(t)
	//		migrator := migrate.Migrator{
	//			Migrations: []migrate.SQLMigration{
	//				migrate.NewQueryMigration("001_create_courses", createWidgetTableSql, dropWidgetsSql),
	//			},
	//		}
	//		err := migrator.Migrate(db)
	//		if err != nil {
	//			t.Fatalf("Migrate() err = %v; want nil", err)
	//		}
	//		_, err = db.Exec("INSERT INTO courses (name) VALUES ($1) ", "cor_test")
	//		if err != nil {
	//			t.Fatalf("db.Exec() err = %v; want nil", err)
	//		}
	//		err = migrator.Rollback(db, "sqlite3")
	//		if err != nil {
	//			t.Fatalf("Rollback() err = %v; want nil", err)
	//		}
	//		var count int
	//		err = db.QueryRow("SELECT COUNT(id) FROM courses;").Scan(&count)
	//		if err == nil {
	//			// Want an error here
	//			t.Fatalf("db.QueryRow() err = nil; want table missing error")
	//		}
	//		// Don't want to test inner workings of lib, so let's just migrate again and verify we have a table now
	//		err = migrator.Migrate(db, "sqlite3")
	//		if err != nil {
	//			t.Fatalf("Migrate() err = %v; want nil", err)
	//		}
	//		_, err = db.Exec("INSERT INTO courses (name) VALUES ($1) ", "cor_test")
	//		if err != nil {
	//			t.Fatalf("db.Exec() err = %v; want nil", err)
	//		}
	//		err = db.QueryRow("SELECT COUNT(*) FROM courses;").Scan(&count)
	//		if err != nil {
	//			// Want an error here
	//			t.Fatalf("db.QueryRow() err = %v; want nil", err)
	//		}
	//		if count != 1 {
	//			t.Fatalf("count = %d; want %d", count, 1)
	//		}
	//	})
}

var (
	createWidgetTableSql = `
CREATE TABLE widgets (
	id serial PRIMARY KEY,
	name text
);`
	createUsersSql = `
CREATE TABLE users (
  id serial PRIMARY KEY,
  email text UNIQUE NOT NULL
);`
	dropWidgetsSql = `DROP TABLE widgets;`
	dropUsersSql   = `DROP TABLE users;`
)
