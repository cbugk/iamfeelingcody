package get

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"net/http"

	"github.com/cbugk/iamfeelingcody/src/internal/sqlc"
	"github.com/cbugk/iamfeelingcody/src/internal/templ"
	"github.com/julienschmidt/httprouter"
)

func FindUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Empty name
	if name := r.URL.Query().Get("name"); name == "" {
		templ.PageFindErrProvideName().Render(r.Context(), w)
	} else {
		user, err := sqlc.Q().GetGithubUser(context.Background(), name)
		if err == nil {
			templ.PageGithubUserFound(user).Render(r.Context(), w)
		} else if errors.As(err, &sql.ErrNoRows) {
			templ.PageGithubUserNotFound(user).Render(r.Context(), w)
		} else {
			log.Fatal(err.Error())
		}
	}
}
