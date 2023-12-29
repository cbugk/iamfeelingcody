package main

import (
	"fmt"
	"net/http"

	"github.com/cbugk/iamfeelingcody/internal/routes"
)

func main() {
	port := 8000
	addr := fmt.Sprintf("localhost:%d", port)
	fmt.Printf("Server listening on http://%s\n", addr)

	http.HandleFunc("/", routes.Route)
	if err := http.ListenAndServe(addr, nil); err != nil {
		panic(err)
	}
}
