package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"URL_Shortener/services"

	"github.com/gorilla/mux"
)

// Create an instance of URLShortener
var urlShortener = services.NewURLShortener()

// Handler to shorten a URL
func shortenURLHandler(w http.ResponseWriter, r *http.Request) {
	type RequestBody struct {
		URL string `json:"url"`
	}

	var body RequestBody
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil || body.URL == "" {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
	}

	shortCode := urlShortener.ShortenURL(body.URL)
	response := map[string]string{
		"short_url": fmt.Sprintf("http://localhost:8080/%s", shortCode),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Handler to redirect to the original URL
func redirectURLHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	shortCode := vars["shortcode"]

	longURL, exists := urlShortener.GetLongURL(shortCode)
	if !exists {
		http.NotFound(w, r)
		return
	}

	http.Redirect(w, r, longURL, http.StatusFound)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/shorten", shortenURLHandler).Methods("POST")
	r.HandleFunc("/{shortcode}", redirectURLHandler).Methods("GET")

	fmt.Println("Server started on :8080")
	http.ListenAndServe(":8080", r)
}
