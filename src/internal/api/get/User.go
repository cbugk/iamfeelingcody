package get

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/cbugk/iamfeelingcody/src/internal/sqlc"
	"github.com/julienschmidt/httprouter"
)

func User(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	if name := r.URL.Query().Get("name"); name == "" {
		// Empty name
		if m, err := json.Marshal(map[string]interface{}{
			"Error":     "provide name",
			"ErrorCode": 1,
		}); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("{}"))
			log.Println(err.Error())
		} else {
			w.WriteHeader(http.StatusNotAcceptable)
			w.Write(m)
		}
	} else if user, err := sqlc.Q().GetGithubUser(context.Background(), name); err == nil {
		// Github user found
		if m, err := json.Marshal(user); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("{}"))
			log.Println(err.Error())
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write(m)
		}
	} else if errors.As(err, &sql.ErrNoRows) {
		// Github user not found
		if m, err := json.Marshal(map[string]interface{}{
			"Error":   "user not found",
			"ErrCode": 2,
		}); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("{}"))
			log.Println(err.Error())
		} else {
			w.WriteHeader(http.StatusNoContent)
			w.Write(m)
		}
	} else {
		// Unknown error
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("{}"))
		log.Println(err.Error())
	}
}
