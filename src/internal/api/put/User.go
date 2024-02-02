package put

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/cbugk/iamfeelingcody/src/internal/check"
	"github.com/cbugk/iamfeelingcody/src/internal/sqlc"
	"github.com/julienschmidt/httprouter"
)

func User(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/json")

	r.ParseForm()
	name := r.Form.Get("name")

	if len(name) == 0 {
		// Name not provided
		w.WriteHeader(http.StatusNoContent)
		w.Write([]byte("{}"))
	} else if user, err := sqlc.Q().GetGithubUser(ctx, name); err == nil {
		// User already created
		if m, err := json.Marshal(user); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("{}"))
			log.Println(err.Error())
		} else {
			w.WriteHeader(http.StatusFound)
			w.Write(m)
		}
	} else if err := check.CheckGithubUser(name); err == nil {
		// Github user's url exists
		if user, err = sqlc.Q().CreateGithubUser(ctx, name); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("{}"))
			log.Println(err.Error())
		} else if m, err := json.Marshal(user); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("{}"))
			log.Println(err.Error())
		} else {
			w.WriteHeader(http.StatusCreated)
			w.Write(m)
		}
	} else if errors.Is(err, &check.ErrorGithubUserNotFound{}) {
		// Github user's url does not exist
		w.WriteHeader(http.StatusNoContent)
		w.Write([]byte("{}"))
	} else {
		// Unanticipated error
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("{}"))
		log.Fatalln(err.Error())
	}
}
