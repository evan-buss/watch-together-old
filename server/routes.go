package server

import "net/http"

// Server encapsulates the server's connections
type Server struct {
	Router *http.ServeMux
}

// Routes handles all application routing
func (s *Server) Routes() {
	s.Router.HandleFunc("/", s.handleIndexPage)
	s.Router.HandleFunc("/media/", s.handleStreamAssets)
	s.Router.HandleFunc("/transcode/", s.handleTranscodeAction)
}
