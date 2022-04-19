package main

import (
	"log"
	"net/http"
)

const httpPort = ":8080"

func main() {
	println("Server running on", httpPort)

	mux := http.NewServeMux()
	mux.HandleFunc("/health", healthHandler)

	log.Fatal(http.ListenAndServe(httpPort, mux))
}

func healthHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Everything is OK!"))
}
