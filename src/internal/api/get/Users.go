package get

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/cbugk/iamfeelingcody/src/internal/sqlc"
	"github.com/julienschmidt/httprouter"
)

func Users(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	if users, err := sqlc.Q().ListGithubUsers(context.Background()); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("{}"))
		log.Println(err.Error())
	} else if m, err := json.Marshal(users); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("{}"))
		log.Println(err.Error())
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write(m)
	}
}
