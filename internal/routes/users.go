package routes

import (
	"net/http"

	"github.com/cbugk/iamfeelingcody/internal/misc"
	"github.com/cbugk/iamfeelingcody/internal/templates"
	"github.com/julienschmidt/httprouter"
)

func users(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	templates.PageFoundUsers(misc.GetDB().UserURIs).Render(r.Context(), w)
}
