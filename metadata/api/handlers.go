package main

import (
	"net/http"

	"github.com/evan-buss/watch-together/metadata/scraper/data"
)

func (s *Server) handleTitleSearch(w http.ResponseWriter, r *http.Request) {
	if search, err := getQuery(r, "search"); err == nil {
		var results []data.ImdbData

		err := s.db.Select(&results, `SELECT (title) FROM movies WHERE title LIKE (?)`, "%"+search+"%")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		if len(results) == 0 {
			http.Error(w, "No matching results", 404)
			return
		}
		responder(w, r, results, 200)
	}
}

this is a test

func (s *Server) handleYearSearch(w http.ResponseWriter, r *http.Request) {
	if search, err := getQuery(r, "search"); err == nil {
		var results []data.ImdbData
		err := s.db.Select(&results, `SELECT * FROM movies WHERE year = (?)`, search)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		if len(results) == 0 {
			http.Error(w, "No matching results", 404)
			return
		}
		responder(w, r, results, 200)
	}
}
