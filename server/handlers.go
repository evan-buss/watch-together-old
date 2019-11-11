package server

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"path/filepath"

	"github.com/evan-buss/watch-together/video"
	"github.com/go-chi/chi"
	"github.com/gorilla/websocket"
)

// Send static media file assets
func (s *Server) handleStreamAssets(w http.ResponseWriter, r *http.Request) {
	file := filepath.Join(viper.GetString("video-dir"), chi.URLParam(r, "fileName"))
	fmt.Println("Requested: " + file)
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	http.ServeFile(w, r, file)
}

// Endpoint to start the media transcode
func (s *Server) handleTranscodeAction(w http.ResponseWriter, r *http.Request) {
	fmt.Println("STARTING TRANSCODE")
	go video.Transcode(filepath.Join(viper.GetString("video-dir"), viper.GetString("file")))
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
		client := &Client{Hub: s.Hub, Conn: conn, Send: make(chan []byte, 256)}
		s.Hub.Register <- client

		go client.ReadPump()
		go client.WritePump()
	}
}
