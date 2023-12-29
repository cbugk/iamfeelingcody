package routes

import (
	"net/http"
	"regexp"
)

var rRoot = regexp.MustCompile(`^/?$`)
var rFind = regexp.MustCompile(`^\/find\/?$`)
var rHello = regexp.MustCompile(`^\/hello\/?.*$`)

func Route(w http.ResponseWriter, r *http.Request) {
	switch {
	case rFind.MatchString(r.URL.Path):
		find(w, r)
	case rHello.MatchString(r.URL.Path):
		hello(w, r)
	case rRoot.MatchString(r.URL.Path):
		root(w, r)
	default:
		notFound(w, r)
	}
}
