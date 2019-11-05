package server

import (
	"github.com/evan-buss/watch-together/server/chat"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// Server encapsulates the server's outside connections
type Server struct {
	Router *chi.Mux
	Hub    *chat.Hub
}

// Routes handles all application routing
func (s *Server) Routes() {

	s.Router.Use(middleware.Logger)

	s.Router.Get("/", s.handleIndexPage)
	s.Router.Get("/media/{fileName}", s.handleStreamAssets)
	s.Router.Get("/transcode/", s.handleTranscodeAction)
	s.Router.Get("/ws", s.handleWebsockets())
}
