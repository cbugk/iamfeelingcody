package routes

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func root(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintln(w, "iamfeelingcody")
}
