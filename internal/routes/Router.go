package routes

import (
	"github.com/julienschmidt/httprouter"
)

func Router() *httprouter.Router {
	mux := httprouter.New()

	mux.GET("/", root)
	mux.GET("/find", find)
	mux.GET("/found", foundUsers)
	mux.GET("/register/:user", registerUser)

	return mux
}
