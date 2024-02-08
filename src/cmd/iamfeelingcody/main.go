package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"net/http"

	"path/filepath"

	"github.com/cbugk/iamfeelingcody/src/internal/ralpv"
	"github.com/cbugk/iamfeelingcody/src/internal/route"
	"github.com/cbugk/iamfeelingcody/src/internal/routine"
	"github.com/cbugk/iamfeelingcody/src/internal/sqlc"
	"github.com/cbugk/iamfeelingcody/src/pkg/binpath"
	pkgRoutine "github.com/cbugk/iamfeelingcody/src/pkg/routine"
	"github.com/dchest/uniuri"
)

func main() {
	workerThreadCount := 10

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
	pkgRoutine.GracefulShutdown(server, idleConnsClosed)

	// Channel for user finders
	names := make(chan string, workerThreadCount)

	// Send random name to channel
	go func() {
		for {
			names <- fmt.Sprint(uniuri.NewLenChars(rand.Intn(1), ralpv.Alpnum), uniuri.NewLenChars(rand.Intn(38), ralpv.Alpnumdash))
		}
	}()

	// Run worker
	for i := 0; i < workerThreadCount; i++ {
		routine.TryAddUser(names)
	}

	go func() {
		if err := server.ListenAndServe(); err != nil &&
			!errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Http server failed to start: %v", err.Error())
		}
	}()

	<-idleConnsClosed
	log.Println("Service stopped")
}
