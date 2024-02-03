package route

import (
	"net/http"

	"github.com/cbugk/iamfeelingcody/src/pkg/binpath"
	"github.com/cbugk/iamfeelingcody/src/pkg/embed"
	"github.com/julienschmidt/httprouter"
)

func RegisterServeFiles(r *httprouter.Router) {
	r.ServeFiles("/static/*filepath", http.FS(embed.Dir()))
	r.ServeFiles("/public/*filepath", http.Dir(binpath.PublicDir()))
}
