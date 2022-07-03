// Package routes handles all routes.  Note that including this package will set the current
// directory to the top of the source tree.  Any references to paths should be from the top
// of the tree.
package routes

import (
	"os"
	"path"
	"runtime"

	"github.com/PaulWaldo/gomoney/pkg/domain"
	"github.com/gin-gonic/gin"
)

// init changes directory to the top of the source tree.  All paths in the application
// need not divine where they are in a relative manner
func init() {
	// https://www.brandur.org/fragments/testing-go-project-root
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "..")
	err := os.Chdir(dir)
	if err != nil {
		panic(err)
	}
}

type Controller struct {
	router   *gin.Engine
	services *domain.Services
}

func NewController(router *gin.Engine, services *domain.Services) Controller {
	return Controller{
		router:   router,
		services: services,
	}
}
