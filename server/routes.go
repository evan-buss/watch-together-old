package server

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/gobuffalo/packr/v2"
)

// Server encapsulates the server's outside connections
type Server struct {
	Router *chi.Mux
	Hub    *Hub
}

// Routes handles all application routing
func (s *Server) Routes() {

	box := packr.New("watch-together", "../web/public")

	s.Router.Use(middleware.Logger)

	s.Router.Handle("/*", http.FileServer(box))
	s.Router.Get("/media/{fileName}", s.handleStreamAssets)
	s.Router.Get("/transcode/", s.handleTranscodeAction)
	s.Router.Get("/ws", s.handleWebsockets())
}
