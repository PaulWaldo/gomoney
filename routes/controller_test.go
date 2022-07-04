// Package routes handles all routes
package routes

import (
	"testing"

	"github.com/PaulWaldo/gomoney/pkg/domain"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func TestNewControllerStoresParams(t *testing.T) {
	c := NewController(gin.Default(), &domain.Services{})
	require.NotNil(t, c.router, "Expecting router to not be nil")
	require.NotNil(t, c.services, "Expecting services to not be nil")
}
