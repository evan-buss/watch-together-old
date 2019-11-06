package server

import (
	"net/http"

	"github.com/evan-buss/watch-together/server/chat"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/gobuffalo/packr/v2"
)

// Server encapsulates the server's outside connections
type Server struct {
	Router *chi.Mux
	Hub    *chat.Hub
}

// TODO: Unable to serve site to local network. FileServe doesn't seem to be working properly

// Routes handles all application routing
func (s *Server) Routes() {

	box := packr.New("watch-together", "../web/public")

	s.Router.Use(middleware.Logger)

	s.Router.Handle("/*", http.FileServer(box))
	s.Router.HandleFunc("/media/{fileName}", s.handleStreamAssets)
	s.Router.HandleFunc("/transcode/", s.handleTranscodeAction)
	s.Router.HandleFunc("/ws", s.handleWebsockets())
}
