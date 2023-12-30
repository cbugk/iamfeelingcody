package routes

import (
	"net/http"

	"github.com/cbugk/iamfeelingcody/internal/templates"
)

func notFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	templates.PageNotFound().Render(r.Context(), w)
}
