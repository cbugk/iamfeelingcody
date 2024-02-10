package routine

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func GracefulShutdown(s *http.Server, c chan<- struct{}) {
	go func() {
		sigInt := make(chan os.Signal, 1)
		signal.Notify(sigInt, os.Interrupt)
		signal.Notify(sigInt, syscall.SIGTERM)
		<-sigInt

		log.Println("service interrupt recieved")

		IsControlCRecieved = true

		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()

		if err := s.Shutdown(ctx); err != nil {
			log.Printf("Http server shutdown error: %v", err)
		}

		log.Println("Http server shutdown complete")

		close(c)
	}()
}
