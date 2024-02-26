package route

import (
	"github.com/cbugk/iamfeelingcody/src/internal/api/v1/htmx/post"
	"github.com/cbugk/iamfeelingcody/src/internal/api/v1/json/delete"
	"github.com/cbugk/iamfeelingcody/src/internal/api/v1/json/get"
	"github.com/cbugk/iamfeelingcody/src/internal/api/v1/json/put"
	"github.com/julienschmidt/httprouter"
)

func RegisterRoutesAPI(r *httprouter.Router) {
	// TODO OpenAPI3 integration

	// JSON methods
	r.DELETE("/api/v1/json/user", delete.User)
	r.GET("/api/v1/json/user", get.User)
	r.GET("/api/v1/json/users", get.Users)
	r.GET("/api/v1/json/random", get.Random)
	r.PUT("/api/v1/json/user", put.User)

	// HTMX methods
	r.POST("/api/v1/htmx/search", post.Search)
}
