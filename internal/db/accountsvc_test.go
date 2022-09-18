package db

// import (
// 	"testing"
// 	"time"

// 	"github.com/PaulWaldo/gomoney/pkg/domain/models"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/require"
// )

// func Test_accountSvc_Create(t *testing.T) {
// 	// teardownSuite, db := setupSuite(t)
// 	// defer teardownSuite(t)

// 	type fields struct {
// 		// db *gorm.DB
// 	}
// 	type args struct {
// 		account models.Account
// 	}
// 	tests := []struct {
// 		name    string
// 		fields  fields
// 		args    args
// 		want    uint
// 		wantErr bool
// 	}{
// 		{
// 			name:    "create success returns id",
// 			fields:  fields{},
// 			args:    args{account: models.Account{Name: "My Checking", Type: models.Checking.Slug}},
// 			want:    1,
// 			wantErr: false,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			teardownTest, tx := setupTest(t, db)
// 			defer teardownTest(t)

// 			as := NewAccountSvc(tx)
// 			err := as.Create(&tt.args.account)
// 			if (err != nil) != tt.wantErr {
// 				t.Fatalf("accountSvc.Create() error = '%v', wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			var got models.Account
// 			tx.First(&got, tt.args.account.ID)
// 			if !(tt.args.account.ID == got.ID && tt.args.account.Name == got.Name && tt.args.account.Type == got.Type) {
// 				t.Fatalf("accountSvc.Create() = \n%v\n, want \n%v\n", got, tt.args.account)
// 			}
// 		})
// 	}
// }

// func Test_accountSvc_Get(t *testing.T) {
// 	teardownTest, tx := setupTest(t, db)
// 	defer teardownTest(t)

// 	as := NewAccountSvc(tx)

// 	type datum struct {
// 		account     models.Account
// 		generatedId uint
// 	}
// 	data := []datum{
// 		{account: models.Account{
// 			Name:         "a1",
// 			Type:         models.Checking.Slug,
// 			Transactions: []models.Transaction{{Payee: "p1"}},
// 		}},
// 		{account: models.Account{Name: "a2", Type: models.Savings.Slug}},
// 		{account: models.Account{Name: "a3", Type: models.CreditCard.Slug}},
// 	}

// 	for i, d := range data {
// 		err := as.Create(&d.account)
// 		if err != nil {
// 			t.Fatalf("accountSvc.Create() error = '%v'", err)
// 		}
// 		data[i].generatedId = d.account.ID
// 	}
// 	for i, d := range data {
// 		got, err := as.Get(d.generatedId)
// 		if err != nil {
// 			t.Fatalf("accountSvc.Get() error = '%v'", err)
// 		}
// 		if i == 0 {
// 			if len(got.Transactions) != 1 {
// 				t.Errorf("Expecting 1 transaction for first item, got %d", len(got.Transactions))
// 			}
// 		}
// 		if got.ID != d.generatedId {
// 			t.Fatalf("Expecting ID to be %d, got %d", d.generatedId, got.ID)
// 		}
// 		if got.Name != d.account.Name {
// 			t.Fatalf("Expecting name to be %s, got %s", d.account.Name, got.Name)
// 		}
// 		if got.Type != d.account.Type {
// 			t.Fatalf("Expecting type to be %s, got %s", d.account.Type, got.Type)
// 		}
// 	}
// }

// func Test_accountSvc_List(t *testing.T) {
// 	teardownTest, tx := setupTest(t, db)
// 	defer teardownTest(t)

// 	as := NewAccountSvc(tx)

// 	type datum struct {
// 		account     models.Account
// 		generatedId uint
// 	}
// 	data := []datum{
// 		{account: models.Account{Name: "a1", Type: models.Checking.Slug}},
// 		{account: models.Account{Name: "a2", Type: models.Savings.Slug}},
// 		{account: models.Account{Name: "a3", Type: models.CreditCard.Slug}},
// 	}

// 	for i, d := range data {
// 		err := as.Create(&d.account)
// 		if err != nil {
// 			t.Fatalf("accountSvc.Create() error = '%v'", err)
// 		}
// 		data[i].generatedId = d.account.ID
// 	}
// 	accounts, err := as.List()
// 	if err != nil {
// 		t.Fatalf("accountSvc.List() error = '%v'", err)
// 	}
// 	if len(accounts) != len(data) {
// 		t.Fatalf("Expecting %d accounts from List, got %d", len(data), len(accounts))
// 	}
// 	for _, d := range data {
// 		var found = false
// 		for _, got := range accounts {
// 			if d.account.Name == got.Name {
// 				found = true

// 				if got.Name != d.account.Name {
// 					t.Fatalf("Expecting name to be %s, got %s", d.account.Name, got.Name)
// 				}
// 				if got.Type != d.account.Type {
// 					t.Fatalf("Expecting type to be %s, got %s", d.account.Type, got.Type)
// 				}
// 				break
// 			}
// 		}
// 		if !found {
// 			t.Fatalf("Expecting to find account %v in list %v, but was not present", d, accounts)
// 		}
// 	}
// }

// func Test_accountSvc_AddTransactions(t *testing.T) {
// 	teardownTest, tx := setupTest(t, db)
// 	defer teardownTest(t)

// 	as := NewAccountSvc(tx)
// 	acct := models.Account{Name: "Account", Type: models.Checking.Slug}
// 	err := as.Create(&acct)
// 	if err != nil {
// 		t.Fatalf("Error creating account %s", err)
// 	}
// 	txns := []models.Transaction{
// 		{Payee: "Me", Type: "D", Amount: 1234.56, Memo: "For testing services", Date: time.Now()},
// 		{Payee: "You", Type: "D", Amount: 1.23, Memo: "Tax", Date: time.Now()},
// 	}
// 	err = as.AddTransactions(acct, txns)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	var gotAcct models.Account
// 	err = tx.Model(&models.Account{}).Preload("Transactions").First(&gotAcct, acct.ID).Error
// 	if err != nil {
// 		t.Fatalf("Error getting updated account %s", err)
// 	}
// 	if len(gotAcct.Transactions) != len(txns) {
// 		t.Fatalf("Expecting %d transaction in Account, got %d", len(txns), len(acct.Transactions))
// 	}
// }

// func Test_accountSvc_Update(t *testing.T) {
// 	teardownTest, tx := setupTest(t, db)
// 	defer teardownTest(t)

// 	saved := models.Account{
// 		Name: "Account",
// 		Type: models.Checking.Slug,
// 	}
// 	tx = tx.Create(&saved)
// 	assert.NoError(t, tx.Error, "Unable to save initial record: %s")

// 	saved.Name = "Changed"
// 	saved.Type = models.Savings.Slug
// 	transactions := []models.Transaction{{Payee: "p1"}, {Payee: "p2"}}
// 	saved.Transactions = transactions
// 	as := NewAccountSvc(tx)
// 	err := as.Update(&saved)
// 	assert.NoError(t, err, "Update threw error: %s", err)

// 	var loaded = models.Account{}
// 	err = tx.Model(&models.Account{}).Preload("Transactions").First(&loaded, saved.ID).Error
// 	require.NoError(t, err, "Result load threw error: %s", err)
// 	assert.Equal(t, loaded.Name, "Changed")
// 	assert.Equal(t, loaded.Type, models.Savings.Slug)
// 	require.Equal(t, len(transactions), len(loaded.Transactions), "Expecting %d transactions, got %d", len(transactions), len(loaded.Transactions))
// 	// assert.ElementsMatch(t, transactions, loaded.Transactions)
// }

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/PaulWaldo/gomoney/pkg/domain/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func NewMock(t *testing.T) (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}

func TestList(t *testing.T) {
	db, mock := NewMock(t)
	as := NewAccountSvc(db)

	query := "SELECT account_id, name, type FROM accounts"

	rows := sqlmock.NewRows([]string{"account_id", "name", "type"}).
		AddRow(1, "name", "type").AddRow(2, "jjj", "kkkk")

	mock.ExpectQuery(query).WillReturnRows(rows)

	accounts, err := as.List()

	assert.Equal(t, 2, len(accounts))
	assert.NoError(t, err)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGet(t *testing.T) {
	db, mock := NewMock(t)
	as := NewAccountSvc(db)

	rows := sqlmock.NewRows([]string{"account_id", "name", "type"}).
		AddRow(1, "the_name", "the_type")
	query := "SELECT account_id, name, type FROM accounts WHERE account_id = \\?"
	mock.ExpectPrepare(query)
	mock.ExpectQuery(query).WithArgs(1).WillReturnRows(rows)

	got, err := as.Get(1)

	assert.NoError(t, err)
	assert.EqualValues(t, 1, got.ID)
	assert.EqualValues(t, "the_name", got.Name)
	assert.EqualValues(t, "the_type", got.Type)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestCreate(t *testing.T) {
	db, mock := NewMock(t)
	as := NewAccountSvc(db)

	acct := models.Account{Name: "fred", Type: models.Checking.Slug}
	query := "INSERT INTO accounts\\(name, type\\) VALUES\\(\\$1, \\$2\\)"
	mock.ExpectBegin()
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(acct.Name, acct.Type).WillReturnResult(sqlmock.NewResult(1, 0))
	mock.ExpectCommit()

	err := as.Create(&acct)

	assert.NoError(t, err)
	err = mock.ExpectationsWereMet()
	require.NoError(t, err, "there were unfulfilled expectations: %s", err)
}

// func TestUpdate(t *testing.T) {
// 	db, mock := NewMock(t)
// 	as := NewAccountSvc(db)

// 	acct := models.Account{Name: "fred", Type: models.Checking.Slug}
// 	query := `UPDATE accounts SET name = \$1, type = \$2 WHERE account_id = \$3`
// 	mock.ExpectBegin()
// 	prep := mock.ExpectPrepare(query)
// 	prep.ExpectExec().WithArgs(acct.Name, acct.Type, 1).WillReturnResult(sqlmock.NewResult(1, 0))
// 	mock.ExpectCommit()

// 	err := as.Update(&acct)

// 	assert.NoError(t, err)
// 	err = mock.ExpectationsWereMet()
// 	require.NoError(t, err, "there were unfulfilled expectations: %s", err)
// }
