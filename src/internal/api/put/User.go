package put

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/cbugk/iamfeelingcody/src/internal/check"
	"github.com/cbugk/iamfeelingcody/src/internal/sqlc"
	"github.com/julienschmidt/httprouter"
)

func User(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	ctx := context.Background()

	r.ParseForm()
	name := r.Form.Get("name")

	if len(name) == 0 {
		// Name not provided
		w.WriteHeader(http.StatusNoContent)
	} else if _, err := sqlc.Q().GetGithubUser(ctx, name); err == nil {
		// User already created
		w.WriteHeader(http.StatusFound)
	} else {
		if err := check.CheckGithubUser(name); err == nil {
			// Github user's url exists
			w.WriteHeader(http.StatusCreated)
			if _, err = sqlc.Q().CreateGithubUser(ctx, name); err != nil {
				log.Fatal(err.Error())
			}
		} else if errors.Is(err, &check.ErrorGithubUserNotFound{}) {
			// Github user's url does not exist
			w.WriteHeader(http.StatusNoContent)
		} else {
			// Unanticipated error
			log.Fatal(err.Error())
		}
	}
}
