package routes

import (
	"net/http"

	"github.com/cbugk/iamfeelingcody/internal/templates"
)

func handleHello(w http.ResponseWriter, r *http.Request) {
	templates.PageHello("Coder").Render(r.Context(), w)
}
