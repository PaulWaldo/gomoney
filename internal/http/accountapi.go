package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/PaulWaldo/gomoney/pkg/domain"
)

type createAccountRequest struct {
	Name        string             `json:"name"`
	AccountType domain.AccountType `json:"accountType"`
}

type createAccountResponse struct {
	ID domain.AccountIDType `json:"accountId"`
}

// type AccountAPI interface {
// 	// create(req createAccountRequest) (createAccountResponse, error)
// }

type muxType = *http.ServeMux

type AccountAPI struct {
	mux muxType
	db  domain.AccountDB
	svc domain.AccountSvc
}

func NewAccountAPI(db domain.AccountDB, svc domain.AccountSvc, mux muxType) AccountAPI {
	a := AccountAPI{db: db, svc: svc, mux: mux}
	a.registerHandlers()
	return a
}

const path = "/accounts"

func (a AccountAPI) handleAccountCreate(w http.ResponseWriter, r *http.Request) {
	var request createAccountRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&request)
	if err != nil {
		http.Error(w, fmt.Sprintf("Bad Request (unable to decode input): %s", err), http.StatusBadRequest)
		return
	}
	id, err := a.svc.Create(request.Name, request.AccountType)
	if err != nil {
		http.Error(w, fmt.Sprintf("Internal Server Error: %s", err), http.StatusInternalServerError)
		return
	}
	response := createAccountResponse{ID: id}
	encoder := json.NewEncoder(w)
	encoder.Encode(response)
}

func (a AccountAPI) registerHandlers() {
	a.mux.HandleFunc(path, a.handleAccountCreate)
}