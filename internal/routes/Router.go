package routes

import (
	"net/http"

	"github.com/cbugk/iamfeelingcody/internal/embed"
	"github.com/julienschmidt/httprouter"
)

func Router() *httprouter.Router {
	mux := httprouter.New()

	mux.ServeFiles("/static/*filepath", http.FS(embed.Dir()))
	mux.GET("/", root)
	mux.GET("/find", find)

	mux.GET("/users", users)
	mux.POST("/user", putUser)

	return mux
}
