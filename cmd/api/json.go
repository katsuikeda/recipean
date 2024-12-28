package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithError(w http.ResponseWriter, code int, msg string, err error) {
	if err != nil {
		log.Println(err)
	}
	if code > 499 {
		log.Printf("Responding with 5XX error: %v", err)
	}

	type errorResponse struct {
		Error string `json:"error"`
	}
	respondWithJSON(w, code, errorResponse{
		Error: msg,
	})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")

	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshalling JSON response for payload %v with status code %d: %v", payload, code, err)
		http.Error(w, `{"error": "Internal Server Error"}`, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(code)
	if _, err := w.Write(data); err != nil {
		log.Printf("Error writing response body: %v", err)
		http.Error(w, `{"error": "Internal Server Error"}`, http.StatusInternalServerError)
		return
	}
}
