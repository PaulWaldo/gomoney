package http

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/PaulWaldo/gomoney/mocks"
)

func TestNewAccountAPI(t *testing.T) {
	mockDb := mocks.AccountDB{}
	mockSvc := mocks.AccountSvc{}
	api := NewAccountAPI(mockDb, mockSvc)
	if !reflect.DeepEqual(mockDb, api.db) {
		t.Errorf("Expexcting DB %v, but got %v", mockDb, api.db)
	}
	if !reflect.DeepEqual(mockSvc, api.svc) {
		t.Errorf("Expexcting Svc %v, but got %v", mockSvc, api.svc)
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
