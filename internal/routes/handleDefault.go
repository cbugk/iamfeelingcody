package routes

import (
	"fmt"
	"net/http"
)

func handleDefault(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "iamfeelingcody")
}
