package service

import (
	"reflect"
	"testing"

	"github.com/PaulWaldo/gomoney/internal/db/models"
	"github.com/PaulWaldo/gomoney/pkg/domain"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// func TestNewAccountSvc(t *testing.T) {
// 	db := mocks.AccountDB{}
// 	svc := NewAccountSvc(db)
// 	if svc == nil {
// 		t.Error("Expecting an accountsvc, got nil")
// 	}
// }

func TestMyGoodness(t *testing.T) {
	// db, mock, _ := sqlmock.New()
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&models.Account{})

	var want = &models.Account{Name: "test", Type: domain.Checking.String()}
	result := db.Create(want)
	if result.Error != nil {
		t.Error(result.Error)
	}
	sut := NewAccountSvc(db)
	got, err := sut.Get(domain.AccountIDType(want.ID))
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("accountSvc.Create() = \n%v, \nwant \n%v", got, want)
	}
}

// func Test_accountSvc_Create(t *testing.T) {
// 	type fields struct {
// 		DB domain.AccountDB
// 	}
// 	type args struct {
// 		name        string
// 		accountType domain.AccountType
// 	}
// 	tests := []struct {
// 		name    string
// 		fields  fields
// 		args    args
// 		want    domain.AccountIDType
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
// 		DB domain.AccountDB
// 	}
// 	type args struct {
// 		ID domain.AccountIDType
// 	}
// 	tests := []struct {
// 		name    string
// 		fields  fields
// 		args    args
// 		want    *domain.Account
// 		wantErr bool
// 	}{
// 		{
// 			name: "get available",
// 			fields: fields{
// 				DB: mocks.AccountDB{
// 					GetAccountResp: &domain.Account{
// 						ID:          123,
// 						Name:        "my acct",
// 						AccountType: domain.Checking,
// 					},
// 				},
// 			},
// 			args: args{ID: 123},
// 			want: &domain.Account{
// 				ID:          123,
// 				Name:        "my acct",
// 				AccountType: domain.Checking,
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
// 	expectedAccounts := []*domain.Account{
// 		{
// 			ID:          0,
// 			Name:        "acct1",
// 			AccountType: domain.Checking,
// 		},
// 		{
// 			ID:          1,
// 			Name:        "acct2",
// 			AccountType: domain.Savings,
// 		},
// 		{
// 			ID:          2,
// 			Name:        "acct3",
// 			AccountType: domain.CreditCard,
// 		},
// 	}

// 	type fields struct {
// 		DB domain.AccountDB
// 	}
// 	tests := []struct {
// 		name    string
// 		fields  fields
// 		want    []*domain.Account
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
// 					ListAccountResp: []*domain.Account{},
// 					ListAccountErr:  errors.New("mocked error"),
// 				},
// 			},
// 			want:    []*domain.Account{},
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
