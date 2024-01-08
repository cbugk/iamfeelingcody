package routes

import (
	"net/http"
	"github.com/cbugk/iamfeelingcody/internal/db"
	"github.com/cbugk/iamfeelingcody/internal/templates"
	"github.com/julienschmidt/httprouter"
)

func users(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	templates.PageFoundUsers(db.AllGithubUsers()).Render(r.Context(), w)
}
