package routes

import (
	"errors"
	"log"
	"net/http"

	"github.com/cbugk/iamfeelingcody/internal/model"
	"github.com/cbugk/iamfeelingcody/internal/templ"
	"github.com/julienschmidt/httprouter"
)

func find(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Empty name
	if name := r.URL.Query().Get("name"); name == "" {
		templ.PageFindErrProvideName().Render(r.Context(), w)
	} else {
		user := model.MakeGithubUser(name)

		err := model.CheckGithubUser(user)

		if err == nil {
			templ.PageGithubUserFound(user).Render(r.Context(), w)
		} else if errors.Is(err, &model.ErrorGithubUserNotFound{}) {
			templ.PageGithubUserNotFound(user).Render(r.Context(), w)
		} else {
			log.Fatal(err.Error())
		}
	}
}
