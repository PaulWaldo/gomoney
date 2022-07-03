// Package routes handles all routes
package routes

import (
	"reflect"
	"testing"

	"github.com/PaulWaldo/gomoney/pkg/domain"
	"github.com/gin-gonic/gin"
)

func TestNewController(t *testing.T) {
	type args struct {
		router   *gin.Engine
		services *domain.Services
	}
	tests := []struct {
		name string
		args args
		want Controller
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewController(tt.args.router, tt.args.services); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewController() = %v, want %v", got, tt.want)
			}
		})
	}
}
