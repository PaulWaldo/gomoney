package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"

	"github.com/PaulWaldo/gomoney/mocks"
	"github.com/PaulWaldo/gomoney/pkg/domain"
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
		a    AccountAPI
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.a.handleAccountCreate(tt.args.w, tt.args.r)
		})
	}
}

func Test_accountAPI_registerHandlers(t *testing.T) {
	a := AccountAPI{
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

func TestHandleSlothfulMessage(t *testing.T) {
	input := createAccountRequest{Name: "testacct", AccountType: domain.Checking}
	var b bytes.Buffer
	encoder := json.NewEncoder(&b)
	if err := encoder.Encode(input); err != nil {
		t.Errorf("Error encoding input: %s", err)
	}

	wr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/sloth", &b)

	var expectedAccountID domain.AccountIDType = 123
	a := AccountAPI{
		mux: http.NewServeMux(),
		db:  mocks.AccountDB{},
		svc: mocks.AccountSvc{CreateAccountResp: expectedAccountID},
	}
	a.handleAccountCreate(wr, req)
	if wr.Code != http.StatusOK {
		t.Errorf("got HTTP status code %d, expected 200", wr.Code)
	}
	fmt.Println(wr.Body.String())

	var j createAccountResponse
	if err := json.NewDecoder(wr.Body).Decode(&j); err != nil {
		t.Errorf("Error decoding output: %s", err)
	}

	if j.ID != expectedAccountID {
		t.Errorf("Expecting created ID to be 0, got %d", j.ID)
	}
}

func Test_accountAPI_handleAccountCreate(t *testing.T) {
	const expectedAccountID domain.AccountIDType = 123
	var b bytes.Buffer
	type fields struct {
		mux               muxType
		db                domain.AccountDB
		svc               domain.AccountSvc
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
		expectedAccountID domain.AccountIDType
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Happy path",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodGet, "/accounts", &b),
				expectedAccountID: 123,
			},
			fields: fields{
				mux: http.NewServeMux(),
				db:  mocks.AccountDB{},
				svc: mocks.AccountSvc{CreateAccountResp: expectedAccountID},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := AccountAPI{
				mux: tt.fields.mux,
				db:  tt.fields.db,
				svc: tt.fields.svc,
			}
			a.handleAccountCreate(tt.args.w, tt.args.r)
	if http.ResponseRecorder(tt.args.w) != http.StatusOK {
		t.Errorf("got HTTP status code %d, expected 200", wr.Code)
	}
	fmt.Println(wr.Body.String())

	var j createAccountResponse
	if err := json.NewDecoder(wr.Body).Decode(&j); err != nil {
		t.Errorf("Error decoding output: %s", err)
	}

	if j.ID != expectedAccountID {
		t.Errorf("Expecting created ID to be 0, got %d", j.ID)
	}
		})
	}
}
