package routes

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/cbugk/iamfeelingcody/internal/glob"
	"github.com/cbugk/iamfeelingcody/internal/model"
	"github.com/julienschmidt/httprouter"
)

func postUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	ctx := context.Background()

	r.ParseForm()
	name := r.Form.Get("name")

	if err := model.CheckGithubUser(name); err == nil {
		w.WriteHeader(http.StatusCreated)
		_, err = glob.Q().CreateGithubUser(ctx, name)
		if err != nil {
			log.Fatal(err.Error())
		}
	} else if errors.Is(err, &model.ErrorGithubUserNotFound{}) {
		w.WriteHeader(http.StatusNoContent)
	} else {
		log.Fatal(err.Error())
	}
}
