package route

import (
	"net/http"

	"github.com/cbugk/iamfeelingcody/src/internal/api/delete"
	"github.com/cbugk/iamfeelingcody/src/internal/api/get"
	"github.com/cbugk/iamfeelingcody/src/internal/api/put"
	"github.com/cbugk/iamfeelingcody/src/pkg/binpath"
	"github.com/cbugk/iamfeelingcody/src/pkg/embed"
	"github.com/julienschmidt/httprouter"
)

func Router() *httprouter.Router {
	r := httprouter.New()

	// Serve files
	r.ServeFiles("/static/*filepath", http.FS(embed.Dir()))
	r.ServeFiles("/public/*filepath", http.Dir(binpath.PublicDir()))

	// API
	r.DELETE("/api/v1/user", delete.User)
	r.GET("/api/v1/user", get.User)
	r.GET("/api/v1/users", get.Users)
	r.PUT("/api/v1/user", put.User)

	// HTMX
	r.GET("/", root)

	return r
}
