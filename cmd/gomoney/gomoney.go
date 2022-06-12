package main

import (
	"fmt"
	"log"
	"net/http"

	// "github.com/PaulWaldo/gomoney/internal/app"
	"github.com/PaulWaldo/gomoney/internal/db"
	ihttp "github.com/PaulWaldo/gomoney/internal/http"
	"github.com/gin-gonic/gin"
	// "github.com/PaulWaldo/gomoney/internal/transactionstore"
	// "github.com/PaulWaldo/gomoney/pkg/domain"
)

// type moneyServer struct {
// 	store      *transactionstore.TransactionStore
// 	acctSvc    domain.AccountSvc
// 	accountAPI ihttp.AccountAPI
// }

// func NewMoneyServer() *moneyServer {
// 	store := transactionstore.New()
// 	db := db.NewMemoryStore()
// 	acctSvc := app.NewAccountSvc(db)
// 	mux := http.NewServeMux()
// 	accountAPI := ihttp.NewAccountAPI(db, acctSvc, mux)
// 	return &moneyServer{
// 		store:      store,
// 		acctSvc:    acctSvc,
// 		accountAPI: accountAPI,
// 	}
// }

type Routable interface {
	AddRoutes()
}

func main() {
	db, err := db.ConnectToDatabase()
	if err != nil {
		panic(fmt.Sprintf("Unable to connect to database: %s", err))
	}
	cont := ihttp.NewController(db)
	cont.AddRoutes(gin.Default()) 

	log.Print("Starting server")
	log.Fatal(http.ListenAndServe("localhost:8080", /*+os.Getenv("SERVERPORT"),
		mux*/nil))
}
