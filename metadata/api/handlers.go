package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/evan-buss/watch-together/metadata/scraper/data"
)

func (s *Server) handleMovies(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Query().Get("title")
	fmt.Println(title)
	if title != "" {
		var results []data.ImdbData

		err := s.db.Select(&results, `SELECT (title) FROM movies WHERE title LIKE (?)`, "%"+title+"%")
		if err != nil {
			log.Println(err)
		}
		if len(results) == 0 {
			http.Error(w, "No matching results", 404)
			return
		}

		fmt.Println(results)
		responder(w, r, results, 200)
	}

}

func responder(w http.ResponseWriter, r *http.Request, obj interface{}, statusCode int) {
	dat, err := json.Marshal(obj)
	if err != nil {
		http.Error(w, "Couldn't encode json", http.StatusInternalServerError)
		return
	}

	fmt.Println(statusCode)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusTeapot)
	w.Write(dat)
}
