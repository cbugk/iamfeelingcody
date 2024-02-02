package route

import (
	"net/http"

	"github.com/cbugk/iamfeelingcody/src/internal/api/get"
	"github.com/cbugk/iamfeelingcody/src/internal/api/put"
	"github.com/cbugk/iamfeelingcody/src/pkg/binpath"
	"github.com/cbugk/iamfeelingcody/src/pkg/embed"
	"github.com/julienschmidt/httprouter"
)

func Router() *httprouter.Router {
	r := httprouter.New()

	r.ServeFiles("/static/*filepath", http.FS(embed.Dir()))
	r.ServeFiles("/public/*filepath", http.Dir(binpath.PublicDir()))

	r.GET("/", root)
	r.GET("/api/v1/finduser", get.FindUser)

	r.GET("/api/v1/users", get.Users)
	r.PUT("/api/v1/user", put.User)

	return r
}
