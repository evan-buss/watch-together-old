package server

import (
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/gobuffalo/packr/v2"
	"github.com/jmoiron/sqlx"
)

// Server encapsulates the server's outside connections
type Server struct {
	Router *chi.Mux
	Hub    *Hub
	DB     *sqlx.DB
}

// Middlewares adds middlwares to our HTTP handlers
func (s *Server) Middlewares() {
	s.Router.Use(middleware.RequestID)
	s.Router.Use(middleware.RealIP)
	s.Router.Use(middleware.Logger)
	s.Router.Use(middleware.Recoverer)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	s.Router.Use(middleware.Timeout(60 * time.Second))

	// CORS Config
	s.Router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{http.MethodHead, http.MethodOptions, http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodPatch},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
		MaxAge:           300,
	}).Handler)
}

// Routes handles all application routing
func (s *Server) Routes() {

	box := packr.New("watch-together", "../web/public")

	s.Router.Handle("/*", http.FileServer(box))
	s.Router.Get("/media/{fileName}", s.handleStreamAssets)
	s.Router.Get("/transcode/", s.handleTranscodeAction)
	s.Router.Get("/ws", s.handleWebsockets())
	s.Router.Get("/library", s.handleGetLibrary)
	s.Router.Post("/library", s.handleUpdateLibrary)
}
