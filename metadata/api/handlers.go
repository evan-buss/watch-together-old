package main

import (
	"net/http"
	"reflect"

	"github.com/evan-buss/watch-together/metadata/scraper/data"
)

// MovieRequest defines the query parameters for a movie search
type MovieRequest struct {
	Title string
	Year  string
}

func (s *Server) handleSearch(w http.ResponseWriter, r *http.Request) {
	query := getMovieQuery(r)
	var dbStmt string
	switch {
	case query.Title != "" && query.Year != "":
		dbStmt = `SELECT * FROM movies WHERE title LIKE (?) AND year = (?)`
	case query.Title != "":
		dbStmt = `SELECT * FROM movies WHERE title LIKE (?)`
	case query.Year != "":
		dbStmt = `SELECT * FROM movies WHERE year = (?)`
	default:
		http.Error(w, "No matching results", 404)
		return
	}

	var results []data.ImdbData
	err := s.db.Select(&results, dbStmt, structToArray(query)...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	if len(results) == 0 {
		http.Error(w, "No matching results", 404)
		return
	}
	responder(w, r, results, 200)
}

func getMovieQuery(r *http.Request) MovieRequest {
	var req MovieRequest
	req.Title = r.URL.Query().Get("title")
	req.Year = r.URL.Query().Get("year")
	return req
}

// Create an array of interfaces from a struct's fields
func structToArray(obj interface{}) []interface{} {
	values := make([]interface{}, 0)
	st := reflect.TypeOf(obj)
	for i := 0; i < st.NumField(); i++ {
		// field := st.Field(i)
		values = append(values, reflect.ValueOf(obj).Field(i).Interface())
	}
	return values
}
