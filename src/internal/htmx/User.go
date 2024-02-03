package htmx

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"net/http"

	"github.com/cbugk/iamfeelingcody/src/internal/sqlc"
	"github.com/cbugk/iamfeelingcody/src/internal/templ/page"
	"github.com/julienschmidt/httprouter"
)

func User(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "text/html")

	if name := r.URL.Query().Get("name"); name == "" {
		// Empty name
		page.FindErrProvideName().Render(r.Context(), w)
	} else if _, err := sqlc.Q().GetGithubUser(context.Background(), name); err != nil {
		if errors.As(err, &sql.ErrNoRows) {
			// Github user not found
			page.GithubUserNotFound(name).Render(r.Context(), w)
		} else {
			// Unknown error
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err.Error())
		}
	} else {
		// Github user found
		page.GithubUserFound(name).Render(r.Context(), w)
	}
}
