package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/jmoiron/sqlx"
	// Use sqlite database driver
	_ "github.com/mattn/go-sqlite3"
)

// Server contains our routes and db connection
type Server struct {
	r  *chi.Mux
	db *sqlx.DB
}

func main() {
	server := Server{
		r:  chi.NewRouter(),
		db: sqlx.MustConnect("sqlite3", "../scraper/ImdbData.db"),
	}

	defer server.db.Close()

	server.Middlewares()
	server.Routes()
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), server.r))
}

// Middlewares sets up our server middlewares
func (s *Server) Middlewares() {
	// A good base middleware stack
	s.r.Use(middleware.RequestID)
	s.r.Use(middleware.RealIP)
	s.r.Use(middleware.Logger)
	s.r.Use(middleware.Recoverer)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	s.r.Use(middleware.Timeout(60 * time.Second))
}

// Routes sets up our API routes
func (s *Server) Routes() {
	s.r.Route("/movie", func(r chi.Router) {
		r.Get("/title", s.handleTitleSearch)
		r.Get("/year", s.handleYearSearch)
	})
}
