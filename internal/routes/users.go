package routes

import (
	"net/http"

	"github.com/cbugk/iamfeelingcody/internal/db"
	"github.com/cbugk/iamfeelingcody/internal/templ"
	"github.com/julienschmidt/httprouter"
)

func users(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	templ.PageFoundUsers(db.AllGithubUsers()).Render(r.Context(), w)
}
