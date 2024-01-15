package handlers

import (
	"fmt"
	"net/http"
)

func GetUsers(rw http.ResponseWriter, r *http.Request) {
	fmt.Println(rw, "Lista todos los usuarios")
}
func GetUser(rw http.ResponseWriter, r *http.Request) {
	fmt.Println(rw, "Lista un usuario")
}
func CreateUser(rw http.ResponseWriter, r *http.Request) {
	fmt.Println(rw, "Crea un usuario")
}
func UpdateUser(rw http.ResponseWriter, r *http.Request) {
	fmt.Println(rw, "Actualiza un usuario")
}

func DeleteUser(rw http.ResponseWriter, r *http.Request) {
	fmt.Println(rw, "Elimina un usuario")
}
