package main

import "net/http"

type Movie struct {
	URL     string `json:"url"`
	Poster  string `json:"poster"`
	Rating  string `json:"rating"`
	Summary string `json:"summary"`
	Title   string `json:"title"`
	Year    string `json:"year"`
}

// TODO: Not sure if I like the package structure I am using. Not really sure the best way to organize the REST code

// SearchYear returns movies made in year
func (s *Server) SearchYear(w http.ResponseWriter, r *http.Request, year string) {
	var movies []Movie
	err := s.db.Select(&movies, `SELECT * FROM movies WHERE year = (?)`, year)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	responder(w, r, movies, 200)
}

// SearchTitle returns movies with matching titles
func (s *Server) SearchTitle(w http.ResponseWriter, r *http.Request, title string) {
	var movies []Movie
	err := s.db.Select(&movies, `SELECT * FROM movies WHERE title LIKE (?)`, title)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	responder(w, r, movies, 200)
}

// SearchTitleYear returns movies with matching titles made in year
func (s *Server) SearchTitleYear(w http.ResponseWriter, r *http.Request, title string, year string) {
	var movies []Movie
	err := s.db.Select(&movies, `SELECT * FROM movies WHERE title LIKE (?) AND year = (?)`, title, year)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	responder(w, r, movies, 200)
}
