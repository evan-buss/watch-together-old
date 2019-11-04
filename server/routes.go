package server

import (
	"net/http"

	"github.com/evan-buss/watch-together/server/chat"
)

// Server encapsulates the server's outside connections
type Server struct {
	Router *http.ServeMux
	Hub    *chat.Hub
}

// Routes handles all application routing
func (s *Server) Routes() {
	s.Router.HandleFunc("/", s.handleIndexPage)
	s.Router.HandleFunc("/media/", s.handleStreamAssets)
	s.Router.HandleFunc("/transcode/", s.handleTranscodeAction)
	s.Router.HandleFunc("/ws", s.handleWebsockets())
}
