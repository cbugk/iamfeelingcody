package routes

import (
	"net/http"

	"github.com/cbugk/iamfeelingcody/internal/misc"
	"github.com/cbugk/iamfeelingcody/internal/templates"
	"github.com/julienschmidt/httprouter"
)

func findUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	templates.PageFind(misc.GithubUserURI(p.ByName("user"))).Render(r.Context(), w)
}
