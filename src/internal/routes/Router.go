package routes

import (
	"net/http"

	"github.com/cbugk/iamfeelingcody/src/internal/binpath"
	"github.com/cbugk/iamfeelingcody/src/internal/embed"
	"github.com/julienschmidt/httprouter"
)

func Router() *httprouter.Router {
	mux := httprouter.New()

	mux.ServeFiles("/static/*filepath", http.FS(embed.Dir()))
	mux.ServeFiles("/public/*filepath", http.Dir(binpath.PublicDir()))

	mux.GET("/", root)
	mux.GET("/find", find)

	mux.GET("/users", users)
	mux.POST("/user", putUser)

	return mux
}
