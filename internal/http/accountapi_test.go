package http

import (
	"net/http"
	"reflect"
	"testing"

	// "github.com/PaulWaldo/gomoney/mocks"
	"github.com/PaulWaldo/gomoney/pkg/domain"
)

func TestNewAccountAPI(t *testing.T) {
	// mockDb := mocks.AccountDB{}
	// mockSvc := mocks.AccountSvc{}
	type args struct {
		db  domain.AccountDB
		svc domain.AccountSvc
	}
	tests := []struct {
		name string
		args args
		want accountAPI
	}{
		// {
		// 	name: "stores parameters",
		// 	args: args{db: mockDb, svc: mockSvc},
		// 	want: accountAPI{db: mockDb, svc: mockSvc},
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAccountAPI(tt.args.db, tt.args.svc); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAccountAPI() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_accountAPI_createHandler(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		a    accountAPI
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.a.createHandler(tt.args.w, tt.args.r)
		})
	}
}

func Test_accountAPI_registerHandlers(t *testing.T) {
	tests := []struct {
		name string
		a    accountAPI
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.a.registerHandlers()
		})
	}
}
