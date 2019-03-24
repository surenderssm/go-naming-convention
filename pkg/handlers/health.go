package handlers

import (
	"fmt"
	"log"
	"net/http"
)

// Ping returns 200 if handler is available
func Ping(w http.ResponseWriter, r *http.Request) {
	log.Println("Ping request recieved")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "ok")
}

// Health perfrom exhaustive health check of dependent modules and return the individual status
func Health(w http.ResponseWriter, r *http.Request) {
	// TODO : node signals can be captured here like memory, CPU, I/O
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "healthy")
}
