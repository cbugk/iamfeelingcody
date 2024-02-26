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
		// Could not get users from DB
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err.Error())
	} else if m, err := json.Marshal(users); err != nil {
		// Could not Marshal user list
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err.Error())
	} else if users == nil {
		// No user present in DB
		w.Write([]byte("[]"))
	} else {
		// User list from DB
		w.WriteHeader(http.StatusOK)
		w.Write(m)
	}
}
