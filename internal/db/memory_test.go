package db

import (
	"reflect"
	"testing"

	"github.com/PaulWaldo/gomoney/pkg/domain"
)

func TestNewMemoryStore(t *testing.T) {
	tests := []struct {
		name string
		want domain.AccountDB
	}{
		{
			name: "initializes",
			want: &memoryStore{
				nextId:   0,
				accounts: make(map[domain.AccountIDType]*domain.Account),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMemoryStore(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMemoryStore() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_memoryStore_Create(t *testing.T) {
	ms := memoryStore{
		accounts: make(map[domain.AccountIDType]*domain.Account),
		nextId:   0,
	}
	id, err := ms.Create("acctName1", domain.Checking)
	if err != nil {
		t.Fatal("Got error creating account")
	}
	if id != 0 {
		t.Fatalf("expecting creation id 0, got %d", id)
	}
	if ms.accounts[0].AccountType != domain.Checking {
		t.Fatalf("expecting account type (0) to be %v, but got %v", domain.Checking, ms.accounts[0].AccountType)
	}
	if ms.accounts[0].Name != "acctName1" {
		t.Fatalf("expecting accout name (0) to be acctName1, but got %s", ms.accounts[0].Name)
	}

	id, err = ms.Create("acctName2", domain.Savings)
	if err != nil {
		t.Fatal("Got error creating account")
	}
	if id != 1 {
		t.Fatalf("expecting creation id 1, got %d", id)
	}
	if ms.accounts[1].AccountType != domain.Savings {
		t.Fatalf("expecting account type (1) to be %v, but got %v", domain.Checking, ms.accounts[0].AccountType)
	}
	if ms.accounts[1].Name != "acctName2" {
		t.Fatalf("expecting accout name (1) to be acctName2, but got %s", ms.accounts[0].Name)
	}
}

func Test_memoryStore_Get(t *testing.T) {
	type fields struct {
		Account map[domain.AccountIDType]*domain.Account
		nextId  domain.AccountIDType
	}
	type args struct {
		id domain.AccountIDType
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *domain.Account
		wantErr bool
	}{
		{
			name: "can retreive",
			fields: fields{
				nextId: 2,
				Account: map[domain.AccountIDType]*domain.Account{
					0: {
						ID:          0,
						Name:        "num1",
						AccountType: domain.Checking,
					},
					1: {
						ID:          1,
						Name:        "num2",
						AccountType: domain.Savings,
					},
				},
			},
			args: args{id: 1},
			want: &domain.Account{
				ID:          1,
				Name:        "num2",
				AccountType: domain.Savings,
			},
			wantErr: false,
		},
		{
			name: "cannot retreive",
			fields: fields{
				nextId: 2,
				Account: map[domain.AccountIDType]*domain.Account{
					0: {
						ID:          0,
						Name:        "num1",
						AccountType: domain.Checking,
					},
					1: {
						ID:          1,
						Name:        "num2",
						AccountType: domain.Savings,
					},
				},
			},
			args:    args{id: 100},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ms := &memoryStore{
				accounts: tt.fields.Account,
				nextId:   tt.fields.nextId,
			}
			got, err := ms.Get(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("memoryStore.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("memoryStore.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_memoryStore_List(t *testing.T) {
	type fields struct {
		accounts map[domain.AccountIDType]*domain.Account
		nextId   domain.AccountIDType
	}
	expectedAccounts := map[domain.AccountIDType]*domain.Account{
		0: {
			ID:          0,
			Name:        "num1",
			AccountType: domain.Checking,
		},
		1: {
			ID:          1,
			Name:        "num2",
			AccountType: domain.Savings,
		},
	}
	tests := []struct {
		name    string
		fields  fields
		want    []*domain.Account
		wantErr bool
	}{
		{
			name:    "show all records",
			fields:  fields{accounts: expectedAccounts, nextId: 2},
			want:    []*domain.Account{
				expectedAccounts[0], expectedAccounts[1],
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ms := &memoryStore{
				accounts: tt.fields.accounts,
				nextId:   tt.fields.nextId,
			}
			got, err := ms.List()
			if (err != nil) != tt.wantErr {
				t.Errorf("memoryStore.List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("memoryStore.List() = %v, want %v", got, tt.want)
			}
		})
	}
}
