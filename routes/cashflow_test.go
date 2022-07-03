package routes

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/PaulWaldo/gomoney/mocks"
	"github.com/PaulWaldo/gomoney/pkg/domain"
	"github.com/PaulWaldo/gomoney/pkg/domain/models"
	"github.com/gin-gonic/gin"
)

func TestController_AddCashFlowRoutes(t *testing.T) {
	type fields struct {
		router   *gin.Engine
		services *domain.Services
	}
	type args struct{ url string }
	tests := []struct {
		name           string
		fields         fields
		args           args
		wantStatusCode int
		wantSubstrings []string
	}{
		{
			name: "Account is displayed",
			fields: fields{
				router: gin.Default(),
				services: &domain.Services{
					Account: mocks.AccountSvc{
						ListAccountResp: []models.Account{
							{Name: "xyzzy"},
							{Name: "Tardis"},
						},
						ListAccountErr: nil,
					},
				},
			},
			args:           args{url: "/cashflow"},
			wantStatusCode: http.StatusOK,
			wantSubstrings: []string{"xyzzy", "Tardis"},
		},
		{
			name: "Database error gives Internal Server Error",
			fields: fields{
				router: gin.Default(),
				services: &domain.Services{
					Account: mocks.AccountSvc{
						ListAccountResp: []models.Account{},
						ListAccountErr:  errors.New("Simulated DB error"),
					},
				},
			},
			args:           args{url: "/cashflow"},
			wantStatusCode: http.StatusInternalServerError,
			wantSubstrings: []string{},
		},
		{
			name: "/ redirects to cashflow",
			fields: fields{
				router: gin.Default(),
				services: &domain.Services{
					Account: mocks.AccountSvc{
						ListAccountResp: []models.Account{},
						ListAccountErr:  nil,
					},
				},
			},
			args:           args{url: "/"},
			wantStatusCode: http.StatusMovedPermanently,
			wantSubstrings: []string{},
		},
		{
			name: "/index.html redirects to cashflow",
			fields: fields{
				router: gin.Default(),
				services: &domain.Services{
					Account: mocks.AccountSvc{
						ListAccountResp: []models.Account{},
						ListAccountErr:  nil,
					},
				},
			},
			args:           args{url: "/index.html"},
			wantStatusCode: http.StatusMovedPermanently,
			wantSubstrings: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			controller := Controller{
				router:   tt.fields.router,
				services: tt.fields.services,
			}
			controller.AddCashFlowRoutes()

			w := httptest.NewRecorder()
			req, err := http.NewRequest(http.MethodGet, tt.args.url, nil)
			require.NoErrorf(t, err, "Got error creating request: %s", err)
			tt.fields.router.ServeHTTP(w, req)
			require.Equal(t, tt.wantStatusCode, w.Code, "expecting response code %d, got %d", tt.wantStatusCode, w.Code)
			for _, sub := range tt.wantSubstrings {
				assert.Contains(t, w.Body.String(), sub)
			}
		})
	}
}
