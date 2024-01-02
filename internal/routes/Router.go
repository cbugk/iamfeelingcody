package routes

import (
	"github.com/julienschmidt/httprouter"
)

func Router() *httprouter.Router {
	mux := httprouter.New()

	mux.GET("/", root)

	mux.GET("/find", find)

	return mux
}
