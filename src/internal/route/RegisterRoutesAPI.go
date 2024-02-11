package route

import (
	"github.com/cbugk/iamfeelingcody/src/internal/api/delete"
	"github.com/cbugk/iamfeelingcody/src/internal/api/get"
	"github.com/cbugk/iamfeelingcody/src/internal/api/put"
	"github.com/julienschmidt/httprouter"
)

func RegisterRoutesAPI(r *httprouter.Router) {
	// TODO OpenAPI3 integration

	// User methods
	r.DELETE("/api/v1/user", delete.User)
	r.GET("/api/v1/user", get.User)
	r.GET("/api/v1/users", get.Users)
	r.GET("/api/v1/random", get.Random)
	r.PUT("/api/v1/user", put.User)
}
