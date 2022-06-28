package db

import (
	"testing"
	"time"

	"github.com/PaulWaldo/gomoney/pkg/domain/models"
)

// func TestNewTransactionSvc(t *testing.T) {
// 	type args struct {
// 		db *gorm.DB
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want domain.TransactionSvc
// 	}{
// 		{
// 			name: "New service stores database",
// 			args: args{db: &gorm.DB{}},
// 			want: domain.TransactionSvc{},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := NewTransactionSvc(tt.args.db); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("NewTransactionSvc() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

func Test_transactionSvc_Create(t *testing.T) {
	type fields struct {
		// db *gorm.DB
	}
	type args struct {
		transaction *models.Transaction
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "No error on creation",
			args: args{transaction: &models.Transaction{
				Payee:  "payee",
				Type:   "t",
				Amount: 123.45,
				Memo:   "memo",
			}},
		},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			teardownTest, tx := setupTest(t, db)
			defer teardownTest(t)
			ts := NewTransactionSvc(tx)
			if err := ts.Create(tt.args.transaction); (err != nil) != tt.wantErr {
				t.Errorf("transactionSvc.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_transactionSvc_Get(t *testing.T) {
	teardownTest, tx := setupTest(t, db)
	defer teardownTest(t)

	var expected = []models.Transaction{
		{Payee: "payee1", Type: "D", Amount: 123.45, Memo: "memo 1", Date: time.Now()},
		{Payee: "payee2", Type: "W", Amount: 678.90, Memo: "memo 2", Date: time.Now()},
	}
	tx.Save(&expected)

	type fields struct{}
	type args struct {
		id uint
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.Transaction
		wantErr bool
	}{
		{
			name:    "Get TX1",
			args:    args{1},
			want:    &expected[0],
			wantErr: false,
		},
		{
			name:    "Get TX2",
			args:    args{2},
			want:    &expected[1],
			wantErr: false,
		},
		{
			name:    "Get missing TX",
			args:    args{666},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := NewTransactionSvc(tx)
			got, err := ts.Get(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("transactionSvc.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				return
			}

			if got.ID != tt.want.ID {
				t.Errorf("Expecting ID to be %d, got %d", tt.want.ID, got.ID)
			}
			if got.Payee != tt.want.Payee {
				t.Errorf("Expecting payee to be %s, got %s", tt.want.Payee, got.Payee)
			}
			if got.Amount != tt.want.Amount {
				t.Errorf("Expecting amount to be %f, got %f", tt.want.Amount, got.Amount)
			}
		})
	}
}

// func Test_transactionSvc_Get(t *testing.T) {
// 	type fields struct {
// 		db *gorm.DB
// 	}
// 	type args struct {
// 		id uint
// 	}
// 	tests := []struct {
// 		name    string
// 		fields  fields
// 		args    args
// 		want    *models.Transaction
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			ts := transactionSvc{
// 				db: tt.fields.db,
// 			}
// 			got, err := ts.Get(tt.args.id)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("transactionSvc.Get() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("transactionSvc.Get() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
