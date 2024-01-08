package routes

import (
	"errors"
	"log"
	"net/http"

	"github.com/cbugk/iamfeelingcody/internal/model"
	"github.com/cbugk/iamfeelingcody/internal/templates"
	"github.com/julienschmidt/httprouter"
)

func find(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Empty name
	if name := r.URL.Query().Get("name"); name == "" {
		templates.PageFindErrProvideName().Render(r.Context(), w)
	} else {
		user := model.MakeGithubUser(name)

		err := model.CheckGithubUser(user)

		if err == nil {
			templates.PageGithubUserFound(user).Render(r.Context(), w)
		} else if errors.Is(err, &model.ErrorGithubUserNotFound{}) {
			templates.PageGithubUserNotFound(user).Render(r.Context(), w)
		} else {
			log.Fatal(err.Error())
		}
	}
}
