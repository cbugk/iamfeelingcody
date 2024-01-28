package routes

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/cbugk/iamfeelingcody/internal/check"
	"github.com/cbugk/iamfeelingcody/internal/sqlc"
	"github.com/julienschmidt/httprouter"
)

func postUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	ctx := context.Background()

	r.ParseForm()
	name := r.Form.Get("name")

	if err := check.CheckGithubUser(name); err == nil {
		w.WriteHeader(http.StatusCreated)
		_, err = sqlc.Q().CreateGithubUser(ctx, name)
		if err != nil {
			log.Fatal(err.Error())
		}
	} else if errors.Is(err, &check.ErrorGithubUserNotFound{}) {
		w.WriteHeader(http.StatusNoContent)
	} else {
		log.Fatal(err.Error())
	}
}
