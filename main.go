package main

import (
	"log"
	"net/http"
	"apirest/handlers"

	"github.com/gorilla/mux"
) 

func main() {

	// Rutas
	mux := mux.NewRouter()

	// Endpoint
	mux.HandleFunc("/api/user", handlers.GetUsers).Methods("GET")
	mux.HandleFunc("/api/user/{id:[0-9]+}", handlers.GetUser).Methods("GET")
	mux.HandleFunc("/api/user", handlers.CreateUser).Methods("POST")
	mux.HandleFunc("/api/user/{id:[0-9]+}", handlers.UpdateUser).Methods("PUT")
	mux.HandleFunc("/api/user/{id:[0-9]+}", handlers.DeleteUser).Methods("DELETE")

	// iniciando servidor
	log.Fatal(http.ListenAndServe(":3000", mux))

}
