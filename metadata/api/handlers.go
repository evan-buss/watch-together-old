package main

import (
	"net/http"
)

// handleSearch is responsible for sending search results
func (s *Server) handleSearch(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Query().Get("title")
	year := r.URL.Query().Get("year")
	switch {
	case title != "" && year != "":
		s.SearchTitleYear(w, r, title, year)
	case title != "":
		s.SearchTitle(w, r, title)
	case year != "":
		s.SearchYear(w, r, year)
	default:
		http.Error(w, "You must enter a query.", 404)
		return
	}
}
