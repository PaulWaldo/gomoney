package db

import (
	"fmt"
	"testing"
	"time"

	"github.com/PaulWaldo/gomoney/pkg/domain/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_transactionSvc_Create(t *testing.T) {
	type fields struct{}
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
	var expected = []models.Transaction{
		{Payee: "payee1", Type: "D", Amount: 123.45, Memo: "memo 1", Date: time.Now()},
		{Payee: "payee2", Type: "W", Amount: 678.90, Memo: "memo 2", Date: time.Now()},
	}

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
			teardownTest, tx := setupTest(t, db)
			defer teardownTest(t)
			tx.Save(&expected)

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

func TestList(t *testing.T) {
	teardownTest, tx := setupTest(t, db)
	defer teardownTest(t)
	const numTxs = 10
	toAdd := make([]models.Transaction, numTxs)
	for i := 0; i < numTxs; i++ {
		toAdd[i] = models.Transaction{Payee: fmt.Sprintf("Payee %d", i)}
	}
	err := tx.Create(toAdd).Error
	require.NoErrorf(t, err, "got error creating initial data: %s", err)

	svc := NewTransactionSvc(tx)
	// svc.SetPaginationScope(scope func(*gorm.DB) *gorm.DB)
	txns, count, err := svc.List()
	require.NoErrorf(t, err, "got error callint List: %s", err)
	require.NotNil(t, txns, "List response is nil")
	assert.EqualValues(t, numTxs, count, "expecting Count to be %d but got %d", numTxs, count)
	assert.Equal(t, numTxs, len(txns), "expecting num Data items to be %d but got %d", numTxs, len(txns))
}

// func Test_transactionSvc_List(t *testing.T) {
// 	type fields struct{}
// 	type initialState []models.Transaction
// 	tests := []struct {
// 		name    string
// 		fields  fields
// 		toAdd   initialState
// 		want    []models.Transaction
// 		wantErr bool
// 	}{
// 		{
// 			name:    "list all in database",
// 			toAdd:   initialState{{Payee: "p1"}, {Payee: "p2"}},
// 			want:    []models.Transaction{{Payee: "p1"}, {Payee: "p2"}},
// 			wantErr: false,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			teardownTest, tx := setupTest(t, db)
// 			defer teardownTest(t)
// 			err := tx.Create(tt.toAdd).Error
// 			require.NoErrorf(t, err, "got error creating initial data: %s", err)

// 			ts := transactionSvc{
// 				db: tx,
// 			}
// 			got, err := ts.List()
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("transactionSvc.List() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}

// 			require.Equal(t, len(tt.want), len(got), "expecting %d elements, got %d", len(tt.want), len(got))
// 			for i := range got {
// 				require.Equal(t, tt.want[i].Payee, got[i].Payee, "expecting payee %s, but got %s", tt.want[i].Payee, got[i].Payee)
// 			}
// 		})
// 	}
// }

func Test__transactionSvc_ListByAccount_ReturnsOnlySelectedTransactions(t *testing.T) {
	teardownTest, tx := setupTest(t, db)
	defer teardownTest(t)
	accounts := []models.Account{{Name: "acct1"}, {Name: "acct2"}}
	tx.Create(&accounts)
	txns := [] models.Transaction{
		// Account 1
		{Payee: "acct1payee1", AccountID: accounts[0].ID},
		{Payee: "acct1payee2", AccountID: accounts[0].ID},
		// Account 2
		{Payee: "acct2payee1", AccountID: accounts[1].ID},
		{Payee: "acct2payee2", AccountID: accounts[1].ID},
	}
	tx.Create((txns))

	svc := NewTransactionSvc(tx)
	got,count,err := svc.ListByAccount(accounts[1].ID)
	require.NoError(t, err)
	require.EqualValuesf(t, 2, count, "expecting 2 transactions for account ID accounts[1].ID, got %d", count)
	assert.Equal(t, txns[2].Payee, got[0].Payee, "Expecting retrieved payee to be %s, got %s", txns[2].Payee, got[0].Payee)
	assert.Equal(t, txns[3].Payee, got[1].Payee, "Expecting retrieved payee to be %s, got %s", txns[3].Payee, got[1].Payee)
}
