package routes

import (
	"net/http"

	"github.com/cbugk/iamfeelingcody/internal/templates"
)

func hello(w http.ResponseWriter, r *http.Request) {
	templates.PageHello("Coder").Render(r.Context(), w)
}
