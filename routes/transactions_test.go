package routes

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	// "github.com/stretchr/testify/assert"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"

	"github.com/PaulWaldo/gomoney/constants"
	"github.com/PaulWaldo/gomoney/internal/db"
	"github.com/PaulWaldo/gomoney/pkg/domain/models"
)

func TestController_AddTransactionRoutes(t *testing.T) {
	services, db, err := db.NewSqliteInMemoryServices(&gorm.Config{}, false)
	if err != nil {
		panic(err)
	}
	var transactions = []models.Transaction{{Payee: "p1"}, {Payee: "p2"}}
	err = db.Create(&transactions).Error
	require.NoError(t, err, "error creating transactions: %s", err)

	r := gin.Default()
	controller := NewController(r, services)

	// controller := Controller{
	// 	router: gin.Default(),
	// 	services: &domain.Services{
	// 		Account: mocks.AccountSvc{},
	// 		Transaction: ,
	// 	},
	// }
	controller.AddTransactionRoutes()
	w := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodGet, constants.TransactionsURL, nil)
	require.NoErrorf(t, err, "Got error creating request: %s", err)
	controller.router.ServeHTTP(w, req)
	require.Equal(t, http.StatusOK, w.Code, "expecting response code %d, got %d", http.StatusInternalServerError, w.Code)
	var response map[string]interface{}
	err = json.NewEncoder(w).Encode(response)
	require.NoErrorf(t, err, "Got error encoding response: %s", err)
	t.Logf("response=%v", response)
	t.Logf("response body=%s", w.Body.String())

	require.NotEmptyf(t, response, "output:\n%s", response)
}
