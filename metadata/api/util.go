package main

import (
	"encoding/json"
	"net/http"
)

func responder(w http.ResponseWriter, r *http.Request, obj interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(obj)
	if err != nil {
		http.Error(w, "Couldn't encode json", http.StatusInternalServerError)
		return
	}
}
