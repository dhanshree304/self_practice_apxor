package main

import (
	"log"
	"net/http"

	"github.com/rs/cors"

	"backend/db"
	"backend/handlers"
)

func main() {
	db.ConnectMongo()

	mux := http.NewServeMux()

	mux.HandleFunc("/users", handlers.GetUsers)
	mux.HandleFunc("/create", handlers.CreateUser)
	mux.HandleFunc("/update", handlers.UpdateUser)
	mux.HandleFunc("/delete", handlers.DeleteUser)

	//  Fix CORS error
	handler := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
	}).Handler(mux)

	log.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
