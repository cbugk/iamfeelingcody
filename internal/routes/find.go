package routes

import (
	"errors"
	"log"
	"net/http"

	"github.com/cbugk/iamfeelingcody/internal/misc"
	"github.com/cbugk/iamfeelingcody/internal/templates"
	"github.com/julienschmidt/httprouter"
)

func find(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Empty username
	if username := r.URL.Query().Get("username"); username == "" {
		templates.PageFindErrProvideUsername().Render(r.Context(), w)
	} else {
		uri, err := misc.CheckGithubUser(username)

		if err == nil {
			templates.PageGithubUserFound(uri).Render(r.Context(), w)
		} else if errors.Is(err, &misc.ErrorGithubUserNotFound{}) {
			templates.PageGithubUserNotFound(uri).Render(r.Context(), w)
		} else {
			log.Fatal(err)
		}
	}
}
