package route

import (
	"github.com/cbugk/iamfeelingcody/src/internal/htmx"
	"github.com/julienschmidt/httprouter"
)

func RegisterRoutesHTMX(r *httprouter.Router) {
	r.GET("/", htmx.Index)
	r.GET("/find", htmx.Find)
	r.GET("/user", htmx.User)
	r.GET("/users", htmx.Users)
}
