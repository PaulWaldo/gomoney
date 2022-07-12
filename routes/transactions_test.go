package routes

import (
	"encoding/json"
	"fmt"
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

	const (
		numItems = 20
		pageSize = 7
	)
	transactions := make([]models.Transaction, numItems)
	for i := 0; i < numItems; i++ {
		transactions[i] = models.Transaction{Payee: fmt.Sprintf("Payee %d", i)}
	}
	err = db.Create(&transactions).Error
	require.NoError(t, err, "error creating transactions: %s", err)

	r := gin.Default()
	controller := NewController(r, services)
	controller.AddTransactionRoutes()

	w := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodGet, constants.TransactionsURL, nil)
	require.NoErrorf(t, err, "Got error creating request: %s", err)
	values := req.URL.Query()
	values.Add("per_page", fmt.Sprintf("%d", pageSize))
	values.Add("page", "1")
	req.URL.RawQuery = values.Encode()
	controller.router.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code, "expecting response code %d, got %d", http.StatusInternalServerError, w.Code)
	t.Logf("response body=%s", w.Body.String())

	var response utils.PaginatedResponse
	err = json.NewDecoder(w.Body).Decode(&response)
	require.NoErrorf(t, err, "Got error encoding response: %s", err)
	t.Logf("converted response=%v", response)

	data := response.Data
	require.Lenf(t, data, pageSize, "expecting returned transaction length to be %d, but was %d", pageSize, len(data))
	for i, gotTx := range data {
		assert.Equal(t, transactions[gotTx.ID-1].Payee, data[i].Payee,
			"expecting element %d's payee to be %s, but got %s",
			transactions[i].Payee, data[i].Payee)
	}

	count := response.Count
	assert.EqualValuesf(t, numItems, count, "expecting response.count to be %d, but got %d", numItems, count)
}
