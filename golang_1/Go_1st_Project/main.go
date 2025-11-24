package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/users", usersHandler)
	http.HandleFunc("/users/", userHandler)
	http.HandleFunc("/health", healthHandler)


	log.Println("Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
