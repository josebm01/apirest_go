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
	users, _ := models.ListUsers()
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
	users, _ := models.GetUser(userId)
	db.Close()

	// respondiendo al cliente con json
	// Marshal transforma el objeto  a json
	output, _ := json.Marshal(users)
	fmt.Fprintln( rw, string(output) )
}


func CreateUser(rw http.ResponseWriter, r *http.Request) {
	// fmt.Fprintln(rw, "Crea un usuario")
	// mandando JSON, no tipo text
	rw.Header().Set("Content-Type", "application/json")

	// obteniendo registro - estructura de User
	user := models.User{}

	// decodificar el json a objeto
	decoder := json.NewDecoder(r.Body)


	// devuelve error si no se pudo procesar el json
	if err := decoder.Decode(&user); err != nil {
		fmt.Fprintln(rw, http.StatusUnprocessableEntity)
	} else {
		db.Connect()
		user.Save()
		db.Close()
	}

	// respondiendo al cliente con json
	// Marshal transforma el objeto  a json
	output, _ := json.Marshal(user)
	fmt.Fprintln( rw, string(output) )
}





func UpdateUser(rw http.ResponseWriter, r *http.Request) {
	// fmt.Fprintln(rw, "Actualiza un usuario")
	// mandando JSON, no tipo text
	rw.Header().Set("Content-Type", "application/json")

	// obteniendo registro - estructura de User
	user := models.User{}

	// decodificar el json a objeto
	decoder := json.NewDecoder(r.Body)


	// devuelve error si no se pudo procesar el json
	if err := decoder.Decode(&user); err != nil {
		fmt.Fprintln(rw, http.StatusUnprocessableEntity)
	} else {
		db.Connect()
		user.Save()
		db.Close()
	}

	// respondiendo al cliente con json
	// Marshal transforma el objeto  a json
	output, _ := json.Marshal(user)
	fmt.Fprintln( rw, string(output) )	
}

func DeleteUser(rw http.ResponseWriter, r *http.Request) {
	// fmt.Fprintln(rw, "Elimina un usuario")
	// mandando JSON, no tipo text
	rw.Header().Set("Content-Type", "application/json")

	// obteniendo id de la url
	vars := mux.Vars(r)
	// convirtiendo a int
	userId, _ := strconv.Atoi(vars["id"])

	db.Connect()
	user, _ := models.GetUser(userId)
	user.Delete()
	db.Close()

	// respondiendo al cliente con json
	// Marshal transforma el objeto  a json
	output, _ := json.Marshal(user)
	fmt.Fprintln( rw, string(output) )
}
