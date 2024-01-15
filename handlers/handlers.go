package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"apirest/db"
	"apirest/models"

	"github.com/gorilla/mux"
)

func GetUsers(rw http.ResponseWriter, r *http.Request) {
	// fmt.Fprintln(rw, "Lista todos los usuarios")

	// mandando JSON, no tipo text
	rw.Header().Set("Content-Type", "application/json")

	db.Connect()
	users := models.ListUsers()
	db.Close()

	// respondiendo al cliente con json
	// Marshal transforma el objeto  a json
	output, _ := json.Marshal(users)
	fmt.Fprintln( rw, string(output) )

}

func GetUser(rw http.ResponseWriter, r *http.Request) {
	// fmt.Fprintln(rw, "Lista un usuario")
	// mandando JSON, no tipo text
	rw.Header().Set("Content-Type", "application/json")

	// obteniendo id de la url
	vars := mux.Vars(r)
	// convirtiendo a int
	userId, _ := strconv.Atoi(vars["id"])

	db.Connect()
	users := models.GetUsers(userId)
	db.Close()

	// respondiendo al cliente con json
	// Marshal transforma el objeto  a json
	output, _ := json.Marshal(users)
	fmt.Fprintln( rw, string(output) )
}


func CreateUser(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "Crea un usuario")
}
func UpdateUser(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "Actualiza un usuario")
}

func DeleteUser(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "Elimina un usuario")
}
