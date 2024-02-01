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
	mux := httprouter.New()

	mux.ServeFiles("/static/*filepath", http.FS(embed.Dir()))
	mux.ServeFiles("/public/*filepath", http.Dir(binpath.PublicDir()))

	mux.GET("/", root)
	mux.GET("/api/v1/finduser", get.FindUser)

	mux.GET("/api/v1/users", get.Users)
	mux.POST("/api/v1/user", put.User)

	return mux
}
