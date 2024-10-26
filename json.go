package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// error response
func respondWithERROR(w http.ResponseWriter, code int, msg string) {
	// error code in 400 range is client side error
	if code > 499 {
		log.Println("Responding with 5XX error:", msg)
	}

	type errResponse struct {
		Error string `json:"error"`
	}

	respondWithJSON(w, code, errResponse{Error: msg})
}

// server response
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	// The json.Marshal function is used to convert the payload into a JSON-formatted byte array (data).
	// If this conversion fails (e.g., if payload contains something that can't be serialized to JSON), the error is captured in err.
	data, err := json.Marshal(payload)

	if err != nil {
		log.Printf("Failed to marshal JSON response : %v", payload)
		w.WriteHeader(500)
		return
	}

	// This line sets the response header Content-Type to application/json, indicating to the client that the response body will contain JSON data.
	w.Header().Add("Content-Type", "application/json")
	// sets the HTTP status code.
	w.WriteHeader(code)
	// sends the actual JSON data as the response body.
	w.Write(data)
}
