package routes

import (
	"net/http"

	"github.com/cbugk/iamfeelingcody/internal/misc"
	"github.com/julienschmidt/httprouter"
)

func registerUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	db := misc.GetDB()
	db.UserURIs = append(db.UserURIs, misc.GithubUserURI(p.ByName("user")))
}
