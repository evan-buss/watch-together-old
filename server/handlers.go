package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"

	"github.com/evan-buss/watch-together/video"
	"github.com/go-chi/chi"
	"github.com/gorilla/websocket"
)

// Send static media file assets
func (s *Server) handleStreamAssets(w http.ResponseWriter, r *http.Request) {
	fileName := chi.URLParam(r, "fileName")
	file := filepath.Join(viper.GetString("video-dir"), fileName)
	fmt.Println("Requested: " + file)
	if strings.Contains(fileName, ".vtt") {
		w.Header().Set("Content-Type", "text/vtt")
	}
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

func (s *Server) handleLibrary(w http.ResponseWriter, r *http.Request) {
	fmt.Println("library")
	type movieMeta struct {
		Location string `json:"location"`
		Metadata int    `json:"metadata"`
	}
	var movies []movieMeta
	err := s.DB.Select(&movies, `SELECT * FROM movies;`)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), 404)
		return
	}

	fmt.Println(movies)

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(movies)
	if err != nil {
		http.Error(w, "json encoding error", http.StatusInternalServerError)
	}
}
