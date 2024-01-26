package routes

import (
	"context"
	"log"
	"net/http"

	"github.com/cbugk/iamfeelingcody/internal/sqlc"
	"github.com/cbugk/iamfeelingcody/internal/templ"
	"github.com/julienschmidt/httprouter"
)

func users(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	users, err := sqlc.Q().ListGithubUsers(context.Background())
	if err != nil {
		log.Fatal(err.Error())
	}
	templ.PageFoundUsers(users).Render(r.Context(), w)
}
