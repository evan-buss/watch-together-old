package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func responder(w http.ResponseWriter, r *http.Request, obj interface{}, statusCode int) {
	dat, err := json.Marshal(obj)
	if err != nil {
		http.Error(w, "Couldn't encode json", http.StatusInternalServerError)
		return
	}

	fmt.Println(statusCode)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.Write(dat)
}

func getQuery(r *http.Request, param string) (string, error) {
	value := r.URL.Query().Get(param)
	if value != "" {
		return value, nil
	}
	return "", errors.New("query param not found")
}
