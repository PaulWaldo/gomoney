package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	// "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/PaulWaldo/gomoney/constants"
	"github.com/PaulWaldo/gomoney/mocks"
	"github.com/PaulWaldo/gomoney/pkg/domain"
	"github.com/PaulWaldo/gomoney/pkg/domain/models"
	"github.com/gin-gonic/gin"
)

func TestController_AddTransactionRoutes(t *testing.T) {
	controller := Controller{
		router: gin.Default(),
		services: &domain.Services{
			Account: mocks.AccountSvc{},
			Transaction: mocks.TransactionSvc{
				ListResp: []models.Transaction{{Payee: "p1"}, {Payee: "p2"}},
			},
		},
	}
	controller.AddTransactionRoutes()
	w := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodGet, constants.TransactionsURL, nil)
	require.NoErrorf(t, err, "Got error creating request: %s", err)
	controller.router.ServeHTTP(w, req)
	require.Equal(t, http.StatusOK, w.Code, "expecting response code %d, got %d", http.StatusInternalServerError, w.Code)
	var response map[string]interface{}
	err = json.NewEncoder(w).Encode(response)
	require.NoErrorf(t, err, "Got error encoding response: %s", err)
	fmt.Printf("response=%v", response)
	fmt.Printf("response body=%s", w.Body.String())

	require.Error(t, err)
}
