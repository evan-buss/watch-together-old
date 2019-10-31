package server

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	transcode "github.com/evan-buss/watch-together/transcoder"
)

var videoDir = os.Getenv("VIDEO_DIR")
var videoFile = os.Getenv("VIDEO_FILE")

// Send the home page (index.html)
func (s *Server) handleIndexPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./site/index.html")
}

// Send static media file assets
func (s *Server) handleStreamAssets(w http.ResponseWriter, r *http.Request) {
	file := filepath.Join(videoDir, r.URL.Path[len("/media/"):])
	fmt.Println("Requested: " + file)
	w.Header().Set("Cache-Control", "no-cache")
	http.ServeFile(w, r, file)
}

// Endpoint to start the media transcode
func (s *Server) handleTranscodeAction(w http.ResponseWriter, r *http.Request) {
	fmt.Println("STARTING TRANSCODE")
	go transcode.Transcode(filepath.Join(videoDir, videoFile))
}
