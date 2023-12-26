package routes

import (
	"net/http"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handleDefault)

	mux.HandleFunc("/hello", handleHello)

	mux.HandleFunc("/find", handleFind)

	return mux
}
