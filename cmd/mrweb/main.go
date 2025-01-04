package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func main() {
	log.Println("Starting mrweb server on :8080")
	handler := http.NewServeMux()
	handler.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		re := response{
			Success: true,
			Message: "Hello, from Mr. Web!",
		}
		j, err := json.Marshal(re)
		if err != nil {
			log.Printf("Error marshalling response: %s\n", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(j)
	})
	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}
