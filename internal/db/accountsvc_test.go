package db

// import (
// 	"database/sql"
// 	"testing"

// 	"github.com/DATA-DOG/go-sqlmock"
// 	"github.com/PaulWaldo/gomoney/pkg/domain/models"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/require"
// )

// func NewMock(t *testing.T) (*sql.DB, sqlmock.Sqlmock) {
// 	db, mock, err := sqlmock.New()
// 	if err != nil {
// 		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
// 	}

// 	return db, mock
// }

// func TestList(t *testing.T) {
// 	db, mock := NewMock(t)
// 	as := NewAccountSvc(db)

// 	query := "SELECT account_id, name, type FROM accounts"

// 	rows := sqlmock.NewRows([]string{"account_id", "name", "type"}).
// 		AddRow(1, "name", "type").AddRow(2, "jjj", "kkkk")

// 	mock.ExpectQuery(query).WillReturnRows(rows)

// 	accounts, err := as.List()

// 	assert.Equal(t, 2, len(accounts))
// 	assert.NoError(t, err)
// 	if err := mock.ExpectationsWereMet(); err != nil {
// 		t.Errorf("there were unfulfilled expectations: %s", err)
// 	}
// }

// func TestGet(t *testing.T) {
// 	db, mock := NewMock(t)
// 	as := NewAccountSvc(db)

// 	rows := sqlmock.NewRows([]string{"account_id", "name", "type"}).
// 		AddRow(1, "the_name", "the_type")
// 	query := "SELECT account_id, name, type FROM accounts WHERE account_id = \\?"
// 	mock.ExpectPrepare(query)
// 	mock.ExpectQuery(query).WithArgs(1).WillReturnRows(rows)

// 	got, err := as.Get(1)

// 	assert.NoError(t, err)
// 	assert.EqualValues(t, 1, got.ID)
// 	assert.EqualValues(t, "the_name", got.Name)
// 	assert.EqualValues(t, "the_type", got.Type)
// 	if err := mock.ExpectationsWereMet(); err != nil {
// 		t.Errorf("there were unfulfilled expectations: %s", err)
// 	}
// }

// func TestCreate(t *testing.T) {
// 	db, mock := NewMock(t)
// 	as := NewAccountSvc(db)

// 	acct := models.Account{Name: "fred", Type: models.Checking.Slug}
// 	query := "INSERT INTO accounts\\(name, type\\) VALUES\\(\\$1, \\$2\\)"
// 	mock.ExpectBegin()
// 	prep := mock.ExpectPrepare(query)
// 	prep.ExpectExec().WithArgs(acct.Name, acct.Type).WillReturnResult(sqlmock.NewResult(1, 0))
// 	mock.ExpectCommit()

// 	err := as.Create(&acct)

// 	assert.NoError(t, err)
// 	err = mock.ExpectationsWereMet()
// 	require.NoError(t, err, "there were unfulfilled expectations: %s", err)
// }

// // func TestUpdate(t *testing.T) {
// // 	db, mock := NewMock(t)
// // 	as := NewAccountSvc(db)

// // 	acct := models.Account{Name: "fred", Type: models.Checking.Slug}
// // 	query := `UPDATE accounts SET name = \$1, type = \$2 WHERE account_id = \$3`
// // 	mock.ExpectBegin()
// // 	prep := mock.ExpectPrepare(query)
// // 	prep.ExpectExec().WithArgs(acct.Name, acct.Type, 1).WillReturnResult(sqlmock.NewResult(1, 0))
// // 	mock.ExpectCommit()

// // 	err := as.Update(&acct)

// // 	assert.NoError(t, err)
// // 	err = mock.ExpectationsWereMet()
// // 	require.NoError(t, err, "there were unfulfilled expectations: %s", err)
// // }
