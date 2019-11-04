package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/evan-buss/watch-together/server/chat"
	"github.com/evan-buss/watch-together/video"
	"github.com/gorilla/websocket"
)

var videoDir = os.Getenv("VIDEO_DIR")
var videoFile = os.Getenv("VIDEO_FILE")

// Send the home page (index.html)
func (s *Server) handleIndexPage(w http.ResponseWriter, r *http.Request) {
	// TODO: Hook this up to the frontend svelte application
	http.ServeFile(w, r, "/home/evan/Documents/watch-together/server/index.html")
}

// Send static media file assets
func (s *Server) handleStreamAssets(w http.ResponseWriter, r *http.Request) {
	file := filepath.Join(videoDir, r.URL.Path[len("/media/"):])
	fmt.Println("Requested: " + file)
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	http.ServeFile(w, r, file)
}

// Endpoint to start the media transcode
func (s *Server) handleTranscodeAction(w http.ResponseWriter, r *http.Request) {
	fmt.Println("STARTING TRANSCODE")
	go video.Transcode(filepath.Join(videoDir, videoFile))
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func (s *Server) handleWebsockets() http.HandlerFunc {
	// Allow same origin requests
	upgrader.CheckOrigin = func(req *http.Request) bool {
		return true
	}

	return func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println(err)
			return
		}
		client := &chat.Client{Hub: s.Hub, Conn: conn, Send: make(chan []byte, 256)}
		s.Hub.Register <- client

		go client.ReadPump()
		go client.WritePump()
	}
}
