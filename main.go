package main

import (
	"log"
	"net/http"
	"urlshortener/router"
)

func main() {
	// router
	r := router.Router()

	// start server
	log.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("Failed to start server:", err)
	}
} 