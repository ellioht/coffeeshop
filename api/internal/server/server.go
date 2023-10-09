package server

import (
	"context"
	"fmt"
	"github.com/ellioht/coffeeshop/internal/coffee"
	"github.com/ellioht/coffeeshop/internal/health"
	"github.com/ellioht/coffeeshop/internal/middleware"
	"github.com/go-chi/chi"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"net/http"
)

type Server struct {
	Router *chi.Mux
	Pool   *pgxpool.Pool
}
type Deps struct {
	mw     *middleware.Middlewares
	coffee *coffee.Service
}

func (s *Server) Run(ctx context.Context) error {

	pool, err := connectToDB(ctx)
	if err != nil {
		panic(err)
	}
	defer pool.Close()
	s.Pool = pool

	err = s.Setup(ctx)
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
	var deps Deps
	deps.mw = &middleware.Middlewares{}

	s.Router = chi.NewRouter()

	coffeeDb := coffee.NewCoffeeDB(s.Pool)
	deps.coffee = coffee.NewService(coffeeDb)

	health.NewHandler(s.Router)
	coffee.NewHandler(s.Router, deps.coffee)
	return nil
}

func connectToDB(ctx context.Context) (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(ctx, "postgres://coffee:coffee_pw@localhost:5432/coffee")
	if err != nil {
		return nil, err
	}
	return pool, nil
}
