package http

import (
	"net/http"
	"net/url"
	"reflect"
	"testing"

	"github.com/PaulWaldo/gomoney/mocks"
)

func TestNewAccountAPI(t *testing.T) {
	mockDb := mocks.AccountDB{}
	mockSvc := mocks.AccountSvc{}
	mux := http.NewServeMux()
	api := NewAccountAPI(mockDb, mockSvc, mux)
	if !reflect.DeepEqual(mockDb, api.db) {
		t.Errorf("Expexcting DB %v, but got %v", mockDb, api.db)
	}
	if !reflect.DeepEqual(mockSvc, api.svc) {
		t.Errorf("Expexcting Svc %v, but got %v", mockSvc, api.svc)
	}
	if !reflect.DeepEqual(mux, api.mux) {
		t.Errorf("Expexcting Mux %v, but got %v", mux, api.mux)
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

// func Test_accountAPI_registerHandlers(t *testing.T) {
// 	mockDb := mocks.AccountDB{}
// 	mockSvc := mocks.AccountSvc{}
// 	mux := http.DefaultServeMux
// 	api := NewAccountAPI(mockDb, mockSvc, *mux)
// 	tests := []struct {
// 		name string
// 		a    accountAPI
// 	}{
// 		{
// 			name: "test handlers are created",
// 			a: api,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			tt.a.registerHandlers()
// 		})
// 	}
// }

func Test_accountAPI_registerHandlers(t *testing.T) {
	a := accountAPI{
		mux: http.NewServeMux(),
		db:  mocks.AccountDB{},
		svc: mocks.AccountSvc{},
	}
	a.registerHandlers()
	url, err := url.Parse("http://localhost/accounts")
	if err != nil {
		t.Fatalf("Can't create URL: %s", err)
	}
	handler, pattern := a.mux.Handler(&http.Request{URL: url})
	if pattern != "/accounts" {
		t.Errorf("Expecting pattern but got %s", pattern)
	}
	if handler == nil {
		t.Errorf("Expecting handler, got %v", handler)
	}
}

// 	type fields struct {
// 		mux muxType
// 		db  domain.AccountDB
// 		svc domain.AccountSvc
// 	}
// 	tests := []struct {
// 		name   string
// 		fields fields
// 	}{
// 		{
// name: "verify all known handlers",
// fields: fields{
// 	mux: *http.DefaultServeMux,
// 	db: mocks.AccountDB{},
// 	svc: mocks.AccountSvc{},
// },
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			a := accountAPI{
// 				mux: tt.fields.mux,
// 				db:  tt.fields.db,
// 				svc: tt.fields.svc,
// 			}
// 			a.registerHandlers()
// 			a.mux.Handler(("/accounts")
// 		})
// 	}
// }
