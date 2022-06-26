package db

import (
	"testing"
	"time"

	"github.com/PaulWaldo/gomoney/pkg/domain/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupSuite(t *testing.T) (teardown func(t *testing.T), db *gorm.DB) {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{SkipDefaultTransaction: true})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&models.Account{}, &models.Transaction{})

	// Return a function to teardown the test
	teardown = func(t *testing.T) {}
	return teardown, db
}

func setupTest(t *testing.T, db *gorm.DB) (teardown func(t *testing.T), tx *gorm.DB) {
	tx = db.Begin()
	teardown = func(t *testing.T) {
		tx.Rollback()
	}
	return teardown, tx
}

func Test_accountSvc_Create(t *testing.T) {
	teardownSuite, db := setupSuite(t)
	defer teardownSuite(t)

	type fields struct {
		// db *gorm.DB
	}
	type args struct {
		name        string
		accountType models.AccountType
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    uint
		wantErr bool
	}{
		{
			name:   "create success returns id",
			fields: fields{
				// db: db,
			},
			args:    args{name: "My Checking", accountType: models.Checking},
			want:    1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			teardownTest, tx := setupTest(t, db)
			defer teardownTest(t)

			as := NewAccountSvc(tx)
			newAcct, err := as.Create(tt.args.name, tt.args.accountType.Slug)
			if (err != nil) != tt.wantErr {
				t.Fatalf("accountSvc.Create() error = '%v', wantErr %v", err, tt.wantErr)
				return
			}
			var got models.Account
			tx.First(&got, newAcct.ID)
			if !(newAcct.ID == got.ID && newAcct.Name == got.Name && newAcct.Type == got.Type) {
				t.Fatalf("accountSvc.Create() = \n%v\n, want \n%v\n", got, newAcct)
			}
		})
	}
}

func Test_accountSvc_Get(t *testing.T) {
	teardownSuite, db := setupSuite(t)
	defer teardownSuite(t)
	teardownTest, tx := setupTest(t, db)
	defer teardownTest(t)

	as := NewAccountSvc(tx)

	type datum struct {
		account     models.Account
		generatedId uint
	}
	data := []datum{
		{account: models.Account{Name: "a1", Type: models.Checking.Slug}},
		{account: models.Account{Name: "a2", Type: models.Savings.Slug}},
		{account: models.Account{Name: "a3", Type: models.CreditCard.Slug}},
	}

	for i, d := range data {
		a, err := as.Create(d.account.Name, d.account.Type)
		if err != nil {
			t.Fatalf("accountSvc.Create() error = '%v'", err)
		}
		data[i].generatedId = a.ID
	}
	for _, d := range data {
		got, err := as.Get(d.generatedId)
		if err != nil {
			t.Fatalf("accountSvc.Get() error = '%v'", err)
		}
		if got.ID != d.generatedId {
			t.Fatalf("Expecting ID to be %d, got %d", d.generatedId, got.ID)
		}
		if got.Name != d.account.Name {
			t.Fatalf("Expecting name to be %s, got %s", d.account.Name, got.Name)
		}
		if got.Type != d.account.Type {
			t.Fatalf("Expecting type to be %s, got %s", d.account.Type, got.Type)
		}
	}
}

func Test_accountSvc_List(t *testing.T) {
	teardownSuite, db := setupSuite(t)
	defer teardownSuite(t)
	teardownTest, tx := setupTest(t, db)
	defer teardownTest(t)

	as := NewAccountSvc(tx)

	type datum struct {
		account     models.Account
		generatedId uint
	}
	data := []datum{
		{account: models.Account{Name: "a1", Type: models.Checking.Slug}},
		{account: models.Account{Name: "a2", Type: models.Savings.Slug}},
		{account: models.Account{Name: "a3", Type: models.CreditCard.Slug}},
	}

	for i, d := range data {
		a, err := as.Create(d.account.Name, d.account.Type)
		if err != nil {
			t.Fatalf("accountSvc.Create() error = '%v'", err)
		}
		data[i].generatedId = a.ID
	}
	accounts, err := as.List()
	if err != nil {
		t.Fatalf("accountSvc.List() error = '%v'", err)
	}
	if len(accounts) != len(data) {
		t.Fatalf("Expecting %d accounts from List, got %d", len(data), len(accounts))
	}
	for _, d := range data {
		var found = false
		for _, got := range accounts {
			if d.account.Name == got.Name {
				found = true

				if got.Name != d.account.Name {
					t.Fatalf("Expecting name to be %s, got %s", d.account.Name, got.Name)
				}
				if got.Type != d.account.Type {
					t.Fatalf("Expecting type to be %s, got %s", d.account.Type, got.Type)
				}
				break
			}
		}
		if !found {
			t.Fatalf("Expecting to find account %v in list %v, but was not present", d, accounts)
		}
	}
}

func Test_accountSvc_AddTransactions(t *testing.T) {
	teardownTest, tx := setupTest(t, db)
	defer teardownTest(t)

	as := NewAccountSvc(tx)
	acct, err := as.Create("Account", models.Checking.Slug)
	if err != nil {
		t.Fatalf("Error creating account %s", err)
	}
	txns := []models.Transaction{
		{Payee: "Me", Type: "D", Amount: 1234.56, Memo: "For testing services", Date: time.Now()},
		{Payee: "You", Type: "D", Amount: 1.23, Memo: "Tax", Date: time.Now()},
	}
	err = as.AddTransactions(acct, txns)
	if err != nil {
		t.Fatal(err)
	}
	var gotAcct models.Account
	err = tx.Model(&models.Account{}).Preload("Transactions").First(&gotAcct, acct.ID).Error
	if err != nil {
		t.Fatalf("Error getting updated account %s", err)
	}
	if len(gotAcct.Transactions) != len(txns) {
		t.Fatalf("Expecting %d transaction in Account, got %d", len(txns), len(acct.Transactions))
	}
}
