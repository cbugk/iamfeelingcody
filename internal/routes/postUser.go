package routes

import (
	"errors"
	"log"
	"net/http"

	"github.com/cbugk/iamfeelingcody/internal/misc"
	"github.com/julienschmidt/httprouter"
)

func postUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	db := misc.GetDB()

	r.ParseForm()
	username := r.Form.Get("username")
	_, err := misc.CheckGithubUser(username)

	if err == nil {
		w.WriteHeader(http.StatusCreated)
		db.UserURIs = append(db.UserURIs, misc.GithubUserURI(username))
	} else if errors.Is(err, &misc.ErrorGithubUserNotFound{}) {
		w.WriteHeader(http.StatusNoContent)
	} else {
		log.Fatal(err)
	}
}
