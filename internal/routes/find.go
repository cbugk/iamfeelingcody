package routes

import (
	"net/http"

	"github.com/cbugk/iamfeelingcody/internal/misc"
	"github.com/cbugk/iamfeelingcody/internal/templates"
	"github.com/julienschmidt/httprouter"
)

func find(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	templates.PageFind(misc.FindGithubUser(1)).Render(r.Context(), w)
}
