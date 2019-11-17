package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

// handleSearch is responsible for sending search results
func (s *Server) handleSearch(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Query().Get("title")
	year := r.URL.Query().Get("year")
	offsetStr := r.URL.Query().Get("offset")

	var offset int
	var err error
	if offsetStr != "" {
		offset, err = strconv.Atoi(offsetStr)
		if err != nil {
			log.Println(err)
			return
		}
	}

	switch {
	case title != "" && year != "":
		s.SearchTitleYear(w, r, title, year, offset)
	case title != "":
		s.SearchTitle(w, r, title, offset)
	case year != "":
		// s.SearchYear(w, r, year)
		// Don't allow year only searches.
		fallthrough
	default:
		http.Error(w, "You must enter a query.", 404)
		return
	}
}

func (s *Server) handleID(w http.ResponseWriter, r *http.Request) {
	movieID := chi.URLParam(r, "movieID")
	s.GetByID(w, r, movieID)
}
