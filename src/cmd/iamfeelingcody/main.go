package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"github.com/cbugk/iamfeelingcody/src/internal/route"
	"github.com/cbugk/iamfeelingcody/src/internal/sqlc"
	"github.com/cbugk/iamfeelingcody/src/pkg/binpath"
)

func main() {
	fmt.Println("Started iamfeelingcody")

	// initialize db
	sqlc.InitDB(filepath.Join(binpath.Dir(), "iamfeelingcody.sqlite"))

	port := 8000
	addr := fmt.Sprintf("localhost:%d", port)
	fmt.Printf("Server listening on http://%s\n", addr)

	server := &http.Server{
		Addr:    addr,
		Handler: route.Router(),
	}

	idleConnsClosed := make(chan struct{})
	go func() {
		sigInt := make(chan os.Signal, 1)
		signal.Notify(sigInt, os.Interrupt)
		signal.Notify(sigInt, syscall.SIGTERM)
		<-sigInt

		log.Println("service interrupt recieved")

		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			log.Printf("Http server shutdown error: %v", err)
		}

		log.Println("Http server shutdown complete")

		close(idleConnsClosed)
	}()

	//http.HandleFunc("/", route.Route)
	//if err := http.ListenAndServe(addr, nil); err != nil &&
	if err := server.ListenAndServe(); err != nil &&
		errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("Http server failed to start: %v", err.Error())
	}

	<-idleConnsClosed
	log.Println("Service stopped")
}
