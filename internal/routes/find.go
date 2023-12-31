package routes

import (
	"net/http"

	"github.com/cbugk/iamfeelingcody/internal/misc"
	"github.com/cbugk/iamfeelingcody/internal/templates"
)

func find(w http.ResponseWriter, r *http.Request) {
	templates.PageFind(misc.FindGithubUser(1)).Render(r.Context(), w)
}
