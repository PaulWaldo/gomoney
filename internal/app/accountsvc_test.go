package app

import (
	"reflect"
	"testing"

	"github.com/PaulWaldo/gomoney/pkg/domain"
)

func TestNewAccountSvc(t *testing.T) {
	type args struct {
		db domain.AccountDB
	}
	tests := []struct {
		name string
		args args
		want domain.AccountSvc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAccountSvc(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAccountSvc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_accountSvc_Create(t *testing.T) {
	type fields struct {
		DB domain.AccountDB
	}
	type args struct {
		a domain.Account
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    domain.AccountIDType
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			as := accountSvc{
				DB: tt.fields.DB,
			}
			got, err := as.Create(tt.args.a)
			if (err != nil) != tt.wantErr {
				t.Errorf("accountSvc.Create() error = %v, wantErr %v", err, tt.wantErr)
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
		DB domain.AccountDB
	}
	type args struct {
		ID domain.AccountIDType
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    domain.Account
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			as := accountSvc{
				DB: tt.fields.DB,
			}
			got, err := as.Get(tt.args.ID)
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
