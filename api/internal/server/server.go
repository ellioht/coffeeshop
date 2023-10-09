package server

import (
	"context"
	"fmt"
	"github.com/ellioht/coffeeshop/internal/health"
	"github.com/go-chi/chi"
	"log"
	"net/http"
)

type Server struct {
	Router *chi.Mux
}

func (s *Server) Run(ctx context.Context) error {
	err := s.Setup(ctx)
	if err != nil {
		return err
	}

	server := http.Server{
		Addr:    ":80",
		Handler: s.Router,
	}

	// Start server in a goroutine
	go func() {
		err = server.ListenAndServe()
		if err != nil {
			panic(err)
		}
	}()

	log.Printf("Server listening on %s\n", ":80")

	<-ctx.Done()

	err = server.Shutdown(context.Background())
	if err != nil {
		return fmt.Errorf("could not gracefully shutdown the server: %v\n", err)
	}

	return nil
}

func (s *Server) Setup(ctx context.Context) error {
	s.Router = chi.NewRouter()
	health.NewHandler(s.Router)
	return nil
}
