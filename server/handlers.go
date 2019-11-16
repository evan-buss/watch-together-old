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

type movieMeta struct {
	RowID    int    `json:"id" sql:"rowid"`
	Location string `json:"location"`
	Metadata int    `json:"metadata"`
}

func (s *Server) handleGetLibrary(w http.ResponseWriter, r *http.Request) {
	fmt.Println("library")

	var movies []movieMeta
	err := s.DB.Select(&movies, `SELECT rowid, * FROM movies;`)
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

func (s *Server) handleUpdateLibrary(w http.ResponseWriter, r *http.Request) {
	type Request struct {
		ID       int `json:"id"`
		Metadata int `json:"metadata"`
	}

	var req Request
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Println(err)
	}
	_, err = s.DB.Exec(`UPDATE movies SET metadata = (?) WHERE rowid = (?)`, req.Metadata, req.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var updatedVal movieMeta
	err = s.DB.Get(&updatedVal, `SELECT rowid, * FROM movies WHERE rowid = (?)`, req.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// http.Error(w, http.StatusText(http.StatusOK), http.StatusOK)
	err = json.NewEncoder(w).Encode(updatedVal)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
