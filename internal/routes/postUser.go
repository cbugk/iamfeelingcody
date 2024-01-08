package routes

import (
	"errors"
	"log"
	"net/http"

	"github.com/cbugk/iamfeelingcody/internal/db"
	"github.com/cbugk/iamfeelingcody/internal/model"
	"github.com/julienschmidt/httprouter"
)

func postUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	r.ParseForm()
	name := r.Form.Get("name")
	user := model.MakeGithubUser(name)

	if err := model.CheckGithubUser(user); err == nil {
		w.WriteHeader(http.StatusCreated)
		db.AddGithubUser(user)
	} else if errors.Is(err, &model.ErrorGithubUserNotFound{}) {
		w.WriteHeader(http.StatusNoContent)
	} else {
		log.Fatal(err.Error())
	}
}
