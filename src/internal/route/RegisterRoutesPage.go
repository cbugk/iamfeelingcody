package route

import (
	"github.com/cbugk/iamfeelingcody/src/internal/htmx"
	"github.com/julienschmidt/httprouter"
)

func RegisterRoutesPage(r *httprouter.Router) {
	r.GET("/", htmx.Index)
	r.GET("/random", htmx.Random)
	r.GET("/search", htmx.Search)
	r.GET("/user", htmx.User)
	r.GET("/users", htmx.Users)
}