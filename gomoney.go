package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/PaulWaldo/gomoney/internal/app"
	"github.com/PaulWaldo/gomoney/internal/db"
	"github.com/PaulWaldo/gomoney/internal/transactionstore"
	"github.com/PaulWaldo/gomoney/pkg/domain"
)

type moneyServer struct {
	store   *transactionstore.TransactionStore
	acctSvc domain.AccountSvc
}

func NewMoneyServer() *moneyServer {
	store := transactionstore.New()
	db := db.NewMemoryStore()
	return &moneyServer{store: store, acctSvc: app.NewAccountSvc(db)}
}

func (ms *moneyServer) transactionHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		ms.createTransactionHandler(w, req)
	}
}

func (ms *moneyServer) createTransactionHandler(w http.ResponseWriter, req *http.Request) {
	var tcr transactionstore.TransactionCreateRequest
	dec := json.NewDecoder(req.Body)
	err := dec.Decode(&tcr)
	if err != nil {
		http.Error(w, "Bad request: "+err.Error(), http.StatusBadRequest)
		return
	}
	resp := ms.store.CreateTransaction(tcr)
	enc := json.NewEncoder(w)
	enc.Encode(resp)
}

func main() {
	mux := http.NewServeMux()
	server := NewMoneyServer()
	mux.HandleFunc("/task/", server.transactionHandler)
	log.Fatal(http.ListenAndServe("localhost:8080"+os.Getenv("SERVERPORT"),
		mux))
}
