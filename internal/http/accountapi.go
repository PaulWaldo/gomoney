package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/PaulWaldo/gomoney/pkg/domain"
)

// swagger:model CreateAccount
type createAccountRequest struct {
	Name        string `json:"name"`
	AccountType string `json:"accountType"`
}

type createAccountResponse struct {
	ID domain.AccountIDType `json:"accountId"`
}

// type AccountAPI interface {
// 	// create(req createAccountRequest) (createAccountResponse, error)
// }

type muxType = *http.ServeMux

type AccountAPI struct {
	Mux muxType
	db  domain.AccountDB
	svc domain.AccountSvc
}

func NewAccountAPI(db domain.AccountDB, svc domain.AccountSvc, mux muxType) AccountAPI {
	a := AccountAPI{db: db, svc: svc, Mux: mux}
	a.registerHandlers()
	return a
}

const path = "/accounts"

// swagger:route GET /accounts admin listCompany
// Get companies list
//
// security:
// - apiKey: []
// responses:
//  401: CommonError
//  200: GetCompanies
func (a AccountAPI) handleAccountCreate(w http.ResponseWriter, r *http.Request) {
	var request createAccountRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&request)
	if err != nil {
		http.Error(w, fmt.Sprintf("Bad Request (unable to decode input): %s", err), http.StatusBadRequest)
		return
	}
	acctType, err := domain.AccountTypeFromString(request.AccountType)
	if err != nil {
		http.Error(w, fmt.Sprintf("Bad Request: %s", err), http.StatusBadRequest)
		return
	}
	id, err := a.svc.Create(request.Name, acctType)
	if err != nil {
		http.Error(w, fmt.Sprintf("Internal Server Error: %s", err), http.StatusInternalServerError)
		return
	}
	response := createAccountResponse{ID: id}
	encoder := json.NewEncoder(w)
	encoder.Encode(response)
}

func (a AccountAPI) accountHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		a.handleAccountCreate(w, r)
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
}

func (a AccountAPI) registerHandlers() {
	a.Mux.HandleFunc(path, a.accountHandler)
}
