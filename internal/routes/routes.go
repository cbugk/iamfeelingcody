package routes

import (
	"fmt"
	"net/http"

	"github.com/cbugk/iamfeelingcody/internal/templates"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "iamfeelingcody")
	})

	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		component := templates.Hello("Coder")
		component.Render(r.Context(), w)
	})

	return mux
}
