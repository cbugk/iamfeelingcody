package htmx

import (
	"net/http"

	"github.com/cbugk/iamfeelingcody/src/internal/templ/page"
	"github.com/julienschmidt/httprouter"
)

func Find(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	page.Find().Render(r.Context(), w)
}
