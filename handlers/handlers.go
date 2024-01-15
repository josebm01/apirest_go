package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"apirest/models"

	"github.com/gorilla/mux"
)

func GetUsers(rw http.ResponseWriter, r *http.Request) {
	// fmt.Fprintln(rw, "Lista todos los usuarios")
	// validando si hay error o no 
	if users, err := models.ListUsers(); err != nil {
		models.SendNoFound(rw)
	} else {
		models.SendData(rw, users)
	}

}

func GetUser(rw http.ResponseWriter, r *http.Request) {
	// fmt.Fprintln(rw, "Lista un usuario")
	if user, err := getUserByRequest(r); err != nil {
		models.SendNoFound(rw)
	} else {
		models.SendData(rw, user)
	}
	
}


func CreateUser(rw http.ResponseWriter, r *http.Request) {
	// fmt.Fprintln(rw, "Crea un usuario")
	// mandando JSON, no tipo text
	// rw.Header().Set("Content-Type", "application/json")

	// obteniendo registro - estructura de User
	user := models.User{}
	// decodificar el json a objeto
	decoder := json.NewDecoder(r.Body)

	// devuelve error si no se pudo procesar el json
	if err := decoder.Decode(&user); err != nil {
		models.SendUnprocessableEntity(rw)
	} else {
		user.Save()
		models.SendData(rw, user)
	}

}




func UpdateUser(rw http.ResponseWriter, r *http.Request) {
	// fmt.Fprintln(rw, "Actualiza un usuario")
	// obteniendo registro por id
	var userId int64

	if user, err := getUserByRequest(r); err != nil {
		models.SendNoFound(rw)
	} else {
		userId = user.Id
	}

	user := models.User{}
	// decodificar el json a objeto
	decoder := json.NewDecoder(r.Body)

	// devuelve error si no se pudo procesar el json
	if err := decoder.Decode(&user); err != nil {
		models.SendUnprocessableEntity(rw)
	} else {
		user.Id = userId
		user.Save()
		models.SendData(rw, user)
	}

}

func DeleteUser(rw http.ResponseWriter, r *http.Request) {
	// fmt.Fprintln(rw, "Elimina un usuario")
	// validando si hay error o no 
	if user, err := getUserByRequest(r); err != nil {
		models.SendNoFound(rw)
	} else {
		user.Delete()
		models.SendData(rw, user)
	}
}



//! función para saber si se está editando o eliminando
func getUserByRequest( r *http.Request ) ( models.User, error) {
	// obteniendo id de la url
	vars := mux.Vars(r)
	// convirtiendo a int
	userId, _ := strconv.Atoi(vars["id"])

	if user, err := models.GetUser(userId); err != nil {
		return *user, err
	} else {
		return *user, nil 
	}
}