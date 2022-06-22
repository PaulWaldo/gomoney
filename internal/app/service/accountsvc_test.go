package service

import (
	"reflect"
	"testing"

	"github.com/PaulWaldo/gomoney/internal/db/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupSuite(t *testing.T) (teardown func(t *testing.T), db *gorm.DB) {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{SkipDefaultTransaction: true})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&models.Account{})

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

// func Test_accountSvc_Get(t *testing.T) {
// 	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	db.AutoMigrate(&models.Account{})

// 	var want = &models.Account{Name: "test", Type: models.Checking.String()}
// 	result := db.Create(want)
// 	if result.Error != nil {
// 		t.Error(result.Error)
// 	}
// 	sut := NewAccountSvc(db)
// 	got, err := sut.Get(models.AccountIDType(want.ID))
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	if want.ID != got.ID {
// 		t.Errorf("Expecting ID %d, got %d", want.ID, got.ID)
// 	}
// 	if want.Name != got.Name {
// 		t.Errorf("Expecting ID %s, got %s", want.Name, got.Name)
// 	}
// 	if want.Type != got.Type {
// 		t.Errorf("Expecting ID %s, got %s", want.Type, got.Type)
// 	}
// }

// func Test_accountSvc_Create(t *testing.T) {
// 	type fields struct {
// 		DB models.AccountDB
// 	}
// 	type args struct {
// 		name        string
// 		accountType models.AccountType
// 	}
// 	tests := []struct {
// 		name    string
// 		fields  fields
// 		args    args
// 		want    models.AccountIDType
// 		wantErr bool
// 	}{
// 		{
// 			name: "create success returns id",
// 			fields: fields{
// 				DB: mocks.AccountDB{
// 					CreateAccountResp: 123,
// 					CreateAccountErr:  nil,
// 				},
// 			},
// 			want:    123,
// 			wantErr: false,
// 		},
// 		{
// 			name: "create failure returns error",
// 			fields: fields{
// 				DB: mocks.AccountDB{
// 					CreateAccountResp: 123,
// 					CreateAccountErr:  errors.New("abc"),
// 				},
// 			},
// 			want:    123,
// 			wantErr: true,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			as := accountSvc{
// 				db: tt.fields.DB,
// 			}
// 			got, err := as.Create(tt.args.name, tt.args.accountType)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("accountSvc.Create() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("accountSvc.Create() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func Test_accountSvc_Get(t *testing.T) {
// 	type fields struct {
// 		DB models.AccountDB
// 	}
// 	type args struct {
// 		ID models.AccountIDType
// 	}
// 	tests := []struct {
// 		name    string
// 		fields  fields
// 		args    args
// 		want    *models.Account
// 		wantErr bool
// 	}{
// 		{
// 			name: "get available",
// 			fields: fields{
// 				DB: mocks.AccountDB{
// 					GetAccountResp: &models.Account{
// 						ID:          123,
// 						Name:        "my acct",
// 						AccountType: models.Checking,
// 					},
// 				},
// 			},
// 			args: args{ID: 123},
// 			want: &models.Account{
// 				ID:          123,
// 				Name:        "my acct",
// 				AccountType: models.Checking,
// 			},
// 			wantErr: false,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			as := accountSvc{
// 				db: tt.fields.DB,
// 			}
// 			got, err := as.Get(tt.args.ID)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("accountSvc.Get() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("accountSvc.Get() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func TestAccountSvc_List(t *testing.T) {
// 	expectedAccounts := []*models.Account{
// 		{
// 			ID:          0,
// 			Name:        "acct1",
// 			AccountType: models.Checking,
// 		},
// 		{
// 			ID:          1,
// 			Name:        "acct2",
// 			AccountType: models.Savings,
// 		},
// 		{
// 			ID:          2,
// 			Name:        "acct3",
// 			AccountType: models.CreditCard,
// 		},
// 	}

// 	type fields struct {
// 		DB models.AccountDB
// 	}
// 	tests := []struct {
// 		name    string
// 		fields  fields
// 		want    []*models.Account
// 		wantErr bool
// 	}{
// 		{
// 			name: "list all accounts",
// 			fields: fields{
// 				DB: mocks.AccountDB{
// 					ListAccountResp: expectedAccounts,
// 					ListAccountErr:  nil,
// 				},
// 			},
// 			want:    expectedAccounts,
// 			wantErr: false,
// 		},
// 		{
// 			name: "list error",
// 			fields: fields{
// 				DB: mocks.AccountDB{
// 					ListAccountResp: []*models.Account{},
// 					ListAccountErr:  errors.New("mocked error"),
// 				},
// 			},
// 			want:    []*models.Account{},
// 			wantErr: true,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			as := accountSvc{
// 				db: tt.fields.DB,
// 			}
// 			got, err := as.List()
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("AccountSvc.List() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("AccountSvc.List() = %+v, want %+v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func setupSuite(t *testing.T) (func(t *testing.T), *gorm.DB) {
// 	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	if err := db.AutoMigrate(&models.Account{}); err != nil {
// 		t.Error(err)
// 	}

// 	// Return a function to teardown the test
// 	teardownSuite := func(t *testing.T) {
// 	}
// 	return teardownSuite, db
// }

// func setupTest(t *testing.T) func(t *testing.T) {
// 	return func(t *testing.T) {}
// }

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
			args: args{name:"My Checking", accountType: models.Checking},
			want:    1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			teardownTest, tx := setupTest(t, db)
			defer teardownTest(t)

			as := accountSvc{
				db: tx,
			}

			got, err := as.Create(tt.args.name, tt.args.accountType)
			if (err != nil) != tt.wantErr {
				t.Errorf("accountSvc.Create() error = '%v', wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("accountSvc.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_accountSvc_Get(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		id uint
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.Account
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			as := accountSvc{
				db: tt.fields.db,
			}
			got, err := as.Get(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("accountSvc.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("accountSvc.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_accountSvc_List(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	tests := []struct {
		name    string
		fields  fields
		want    []*models.Account
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			as := accountSvc{
				db: tt.fields.db,
			}
			got, err := as.List()
			if (err != nil) != tt.wantErr {
				t.Errorf("accountSvc.List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("accountSvc.List() = %v, want %v", got, tt.want)
			}
		})
	}
}
