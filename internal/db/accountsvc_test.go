package db

import (
	"testing"
	"time"

	"github.com/PaulWaldo/gomoney/pkg/domain/models"
)


func Test_accountSvc_Create(t *testing.T) {
	// teardownSuite, db := setupSuite(t)
	// defer teardownSuite(t)

	type fields struct {
		// db *gorm.DB
	}
	type args struct {
		account models.Account
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    uint
		wantErr bool
	}{
		{
			name:    "create success returns id",
			fields:  fields{},
			args:    args{account: models.Account{Name: "My Checking", Type: models.Checking.Slug}},
			want:    1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			teardownTest, tx := setupTest(t, db)
			defer teardownTest(t)

			as := NewAccountSvc(tx)
			err := as.Create(&tt.args.account)
			if (err != nil) != tt.wantErr {
				t.Fatalf("accountSvc.Create() error = '%v', wantErr %v", err, tt.wantErr)
				return
			}
			var got models.Account
			tx.First(&got, tt.args.account.ID)
			if !(tt.args.account.ID == got.ID && tt.args.account.Name == got.Name && tt.args.account.Type == got.Type) {
				t.Fatalf("accountSvc.Create() = \n%v\n, want \n%v\n", got, tt.args.account)
			}
		})
	}
}

func Test_accountSvc_Get(t *testing.T) {
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
		err := as.Create(&d.account)
		if err != nil {
			t.Fatalf("accountSvc.Create() error = '%v'", err)
		}
		data[i].generatedId = d.account.ID
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
		err := as.Create(&d.account)
		if err != nil {
			t.Fatalf("accountSvc.Create() error = '%v'", err)
		}
		data[i].generatedId = d.account.ID
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
	acct := models.Account{Name: "Account", Type: models.Checking.Slug}
	err := as.Create(&acct)
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
