package main

import "net/http"

// Movie represents a single movie result returned from the database
type Movie struct {
	RowID   int    `json:"id" sql:"rowid"`
	URL     string `json:"url"`
	Poster  string `json:"poster"`
	Rating  string `json:"rating"`
	Summary string `json:"summary"`
	Title   string `json:"title"`
	Year    string `json:"year"`
}

// GetByID returns the movie matching the given movieID
func (s *Server) GetByID(w http.ResponseWriter, r *http.Request, movieID string) {
	var movie Movie
	err := s.db.Get(&movie, `SELECT  rowid, * FROM movies WHERE rowid = (?)`, movieID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	responder(w, r, movie, 200)
}

// SearchYear returns movies made in year
func (s *Server) SearchYear(w http.ResponseWriter, r *http.Request, year string) {
	var movies []Movie
	err := s.db.Select(&movies, `SELECT rowid, * FROM movies WHERE year = (?)`, year)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	responder(w, r, movies, 200)
}

// SearchTitle returns movies with matching titles
func (s *Server) SearchTitle(w http.ResponseWriter, r *http.Request, title string) {
	var movies []Movie
	err := s.db.Select(&movies, `SELECT rowid, * FROM movies WHERE LOWER(title) LIKE (?)`, "%"+title+"%")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	responder(w, r, movies, 200)
}

// SearchTitleYear returns movies with matching titles made in year
func (s *Server) SearchTitleYear(w http.ResponseWriter, r *http.Request, title string, year string) {
	var movies []Movie
	err := s.db.Select(&movies, `SELECT rowid, * FROM movies WHERE LOWER(title) LIKE (?) AND year = (?)`, "%"+title+"%", year)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	responder(w, r, movies, 200)
}
