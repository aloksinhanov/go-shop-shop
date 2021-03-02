package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
)

type Server struct {
	name   string
	Router *mux.Router
}

func New(name string) *Server {
	return &Server{
		name:   name,
		Router: mux.NewRouter(),
	}
}

func (s Server) Start(ctx context.Context, port string) {
	log.Printf("Starting Server: %v", s.name)
	srv := http.Server{
		Addr:    ":" + port,
		Handler: s.Router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("Unable to start server: %v | error: %v", s.name, err)
		}
	}()

	//Listen for interruptions
	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGTERM, syscall.SIGKILL)

	<-done
	log.Printf("Recived interrupt. Shutting down: %v", s.name)
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	srv.Shutdown(ctx)
}
