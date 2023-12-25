package main

import (
	"fmt"
	"net/http"

	"github.com/cbugk/iamfeelingcody/internal/routes"
)

func main() {
	router := routes.NewRouter()

	port := 8000
	addr := fmt.Sprintf("localhost:%d", port)
	fmt.Printf("Server listening on http://%s\n", addr)

	if err := http.ListenAndServe(addr, router); err != nil {
		panic(err)
	}
}
