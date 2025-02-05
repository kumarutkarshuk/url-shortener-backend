package controller

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"os"
	"time"
	"urlshortener/model"
	"urlshortener/util"

	"github.com/gorilla/mux"
)

func generateShortURL(url string) string {
	// Create SHA256 hash of the URL
	hasher := sha256.New()
	hasher.Write([]byte(url))
	sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	
	// Return first 8 characters as short URL
	return sha[:8]
}

func CreateShortURL(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	var req model.ShortenRequest

	// check errors
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: "Invalid request"})
		return
	}

	if req.URL == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: "URL is required"})
		return
	}

	// Generate short URL
	shortURL := generateShortURL(req.URL)

	// Store in Redis with 24 hour expiration
	ctx := context.Background()
	err := db.Rdb.Set(ctx, shortURL, req.URL, 24*time.Hour).Err()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: "Failed to store URL"})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(model.ShortenResponse{
		ShortURL: os.Getenv("DOMAIN") + shortURL,
	})
}

func RedirectToOriginal(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	shortURL := vars["shortURL"]
	
	// Get original URL from Redis
	ctx := context.Background()
	originalURL, err := db.Rdb.Get(ctx, shortURL).Result()
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: "URL not found"})
		return
	}

	http.Redirect(w, r, originalURL, http.StatusMovedPermanently)
} 