package htmx

import (
	"context"
	"log"
	"net/http"

	"github.com/cbugk/iamfeelingcody/src/internal/sqlc"
	"github.com/cbugk/iamfeelingcody/src/internal/templ/page"
	"github.com/julienschmidt/httprouter"
)

func Users(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "text/html")

	if users, err := sqlc.Q().ListGithubUsers(context.Background()); err != nil {
		// Could not get users from DB
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err.Error())
	} else if users == nil {
		// No user present in DB
		page.NoUsers().Render(r.Context(), w)
	} else {
		// User list from DB
		page.FoundUsers(users).Render(r.Context(), w)
	}
}
