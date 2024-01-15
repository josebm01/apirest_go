package main 

import (
	"github.com/gorilla/mux"
)

func main() {

	// Rutas
	mux := mux.NewRouter()

	// Endpoint 
	mux.HandleFunc("/api/user", nil).Methods("GET")
	mux.HandleFunc("/api/user/{id:[0-9]+}", nil).Methods("GET")
	mux.HandleFunc("/api/user", nil).Methods("POST")
	mux.HandleFunc("/api/user/{id:[0-9]+}", nil).Methods("PUT")
	mux.HandleFunc("/api/user/{id:[0-9]+}", nil).Methods("DELETE")

}