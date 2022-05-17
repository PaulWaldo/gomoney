package http

import (
	"bytes"
	"encoding/json"
	"errors"
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

func TestHandleCreateAccountMessage_decodeInputProblem(t *testing.T) {
	var b bytes.Buffer

	wr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/accounts", &b)

	var expectedAccountID domain.AccountIDType = 123
	a := AccountAPI{
		mux: http.NewServeMux(),
		db:  mocks.AccountDB{},
		svc: mocks.AccountSvc{CreateAccountResp: expectedAccountID},
	}
	a.handleAccountCreate(wr, req)
	if wr.Code != http.StatusBadRequest {
		fmt.Println(wr.Body)
		t.Errorf("got HTTP status code %d, expected %d", wr.Code, http.StatusBadRequest)
	}
}

func TestHandleCreateAccountMessage_happyPath(t *testing.T) {
	input := createAccountRequest{Name: "testacct", AccountType: domain.Checking}
	var b bytes.Buffer
	encoder := json.NewEncoder(&b)
	if err := encoder.Encode(input); err != nil {
		t.Errorf("Error encoding input: %s", err)
	}

	wr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/accounts", &b)

	var expectedAccountID domain.AccountIDType = 123
	a := AccountAPI{
		mux: http.NewServeMux(),
		db:  mocks.AccountDB{},
		svc: mocks.AccountSvc{CreateAccountResp: expectedAccountID},
	}
	a.handleAccountCreate(wr, req)
	if wr.Code != http.StatusOK {
		t.Errorf("got HTTP status code %d, expected %d", wr.Code, http.StatusOK)
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

func TestHandleCreateAccountMessage_serviceFailure(t *testing.T) {
	input := createAccountRequest{Name: "testacct", AccountType: domain.Checking}
	var b bytes.Buffer
	encoder := json.NewEncoder(&b)
	if err := encoder.Encode(input); err != nil {
		t.Errorf("Error encoding input: %s", err)
	}

	wr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/accounts", &b)

	a := AccountAPI{
		mux: http.NewServeMux(),
		db:  mocks.AccountDB{},
		svc: mocks.AccountSvc{CreateAccountErr: errors.New("Zoiks")},
	}
	a.handleAccountCreate(wr, req)
	if wr.Code != http.StatusInternalServerError {
		t.Errorf("got HTTP status code %d, expected %d", wr.Code, http.StatusInternalServerError)
	}
	fmt.Println(wr.Body.String())
}
