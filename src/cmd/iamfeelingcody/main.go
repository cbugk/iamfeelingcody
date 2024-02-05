package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"path/filepath"

	"github.com/cbugk/iamfeelingcody/src/internal/route"
	"github.com/cbugk/iamfeelingcody/src/internal/routine"
	"github.com/cbugk/iamfeelingcody/src/internal/sqlc"
	"github.com/cbugk/iamfeelingcody/src/pkg/binpath"
	pkgRoutine "github.com/cbugk/iamfeelingcody/src/pkg/routine"
)

func main() {
	workerThreadCount := 1

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
	// Run task assigner
	routine.AssignNextName(names)
	// Run workers
	for i := 0; i < workerThreadCount; i++ {
		routine.TryAddUser(names)
	}

	//http.HandleFunc("/", route.Route)
	//if err := http.ListenAndServe(addr, nil); err != nil &&
	if err := server.ListenAndServe(); err != nil &&
		errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("Http server failed to start: %v", err.Error())
	}

	<-idleConnsClosed
	log.Println("Service stopped")
}
