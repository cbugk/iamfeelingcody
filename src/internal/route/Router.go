package route

import (
	"github.com/julienschmidt/httprouter"
)

func Router() *httprouter.Router {
	r := httprouter.New()

	// /static
	// /public
	RegisterServeFiles(r)

	// /api/v1
	RegisterRoutesAPI(r)

	// /
	RegisterRoutesHTMX(r)

	return r
}
