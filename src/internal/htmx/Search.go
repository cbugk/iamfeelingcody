package htmx

import (
	"context"
	"log"
	"net/http"

	"github.com/cbugk/iamfeelingcody/src/internal/sqlc"
	"github.com/cbugk/iamfeelingcody/src/internal/templ/page"
	"github.com/julienschmidt/httprouter"
)

func Search(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	users, err := sqlc.Q().ListGithubUsers(context.Background())
	if err != nil {
		log.Println(err)
		return
	}
	page.Search(users).Render(r.Context(), w)
}
