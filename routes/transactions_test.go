package routes

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	// "github.com/stretchr/testify/assert"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"

	"github.com/PaulWaldo/gomoney/constants"
	"github.com/PaulWaldo/gomoney/internal/db"
	"github.com/PaulWaldo/gomoney/pkg/domain/models"
	"github.com/PaulWaldo/gomoney/utils"
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
	controller.AddTransactionRoutes()

	w := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodGet, constants.TransactionsURL, nil)
	require.NoErrorf(t, err, "Got error creating request: %s", err)
	controller.router.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code, "expecting response code %d, got %d", http.StatusInternalServerError, w.Code)
	t.Logf("response body=%s", w.Body.String())

	// Convert JSON to map
	var response utils.PaginatedResponse
	err = json.NewEncoder(w.Body).Encode(&response)
	// fmt.Printf("Transaction body: %s", w.Body.String())
	// err = json.Unmarshal([]byte(w.Body.Bytes()), &response)
	require.NoErrorf(t, err, "Got error encoding response: %s", err)
	t.Logf("response=%v", response)

	data := response.Data
	require.Lenf(t, data, len(transactions), "expecting returned transaction length to be %d, but was %d", len(transactions), len(data))
	for i := range transactions {
		assert.Equal(t, transactions[i].Payee, data[i].Payee,
			"expecting element %d's payee to be %s, but got %s",
			transactions[i].Payee, data[i].Payee)
	}
}
