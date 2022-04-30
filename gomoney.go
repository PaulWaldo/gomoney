package main

import (
	"log"
	"net/http"

	"github.com/PaulWaldo/gomoney/internal/app"
	"github.com/PaulWaldo/gomoney/internal/db"
	ihttp "github.com/PaulWaldo/gomoney/internal/http"
	"github.com/PaulWaldo/gomoney/internal/transactionstore"
	"github.com/PaulWaldo/gomoney/pkg/domain"
)

type moneyServer struct {
	store      *transactionstore.TransactionStore
	acctSvc    domain.AccountSvc
	accountAPI ihttp.AccountAPI
}

func NewMoneyServer() *moneyServer {
	store := transactionstore.New()
	db := db.NewMemoryStore()
	acctSvc := app.NewAccountSvc(db)
	mux := http.NewServeMux()
	accountAPI := ihttp.NewAccountAPI(db, acctSvc, mux)
	return &moneyServer{
		store:      store,
		acctSvc:    acctSvc,
		accountAPI: accountAPI,
	}
}

// func (ms *moneyServer) transactionHandler(w http.ResponseWriter, req *http.Request) {
// 	if req.Method == http.MethodPost {
// 		ms.createTransactionHandler(w, req)
// 	}
// }

// func (ms *moneyServer) createTransactionHandler(w http.ResponseWriter, req *http.Request) {
// 	var tcr transactionstore.TransactionCreateRequest
// 	dec := json.NewDecoder(req.Body)
// 	err := dec.Decode(&tcr)
// 	if err != nil {
// 		http.Error(w, "Bad request: "+err.Error(), http.StatusBadRequest)
// 		return
// 	}
// 	resp := ms.store.CreateTransaction(tcr)
// 	enc := json.NewEncoder(w)
// 	enc.Encode(resp)
// }

func main() {
	/*server := */
	NewMoneyServer()
	// mux.HandleFunc("/task/", server.transactionHandler)
	log.Print("Starting server")
	log.Fatal(http.ListenAndServe("localhost:8080", /*+os.Getenv("SERVERPORT"),
		mux*/nil))
}
