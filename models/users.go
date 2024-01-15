package models

import (
	"apirest/db"
)

// campos
type User struct {

	// agregando alias para que el cliente lo reciba con minúscula en el JSON
	Id       int64 `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type Users []User

const UserSchema string = `CREATE TABLE users (
	id INT(6) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
	username VARCHAR(30) NOT NULL,
	password VARCHAR(100) NOT NULL,
	email VARCHAR(50),
	create_data TIMESTAMP DEFAULT CURRENT_TIMESTAMP)`

// ! Construir usuario
func NewUser(username, password, email string) *User {
	// Puntero en base a la estructura
	user := &User{Username: username, Password: password, Email: email}
	return user
}

// ! Crear usuario e insertar
func CreateUser(username, password, email string) *User {
	user := NewUser(username, password, email)
	user.Save()
	return user
}

// ! Insertar registro (método privado)
func (user *User) insert() {
	sql := "INSERT users SET username=?, password=?, email=?"
	result, _ := db.Exec(sql, user.Username, user.Password, user.Email)

	// obteniendo ID para asignarlo y mostrar el id del usuario insertado en la bd
	user.Id, _ = result.LastInsertId()
}

// ! Listar todos los registros
func ListUsers() Users {
	sql := "SELECT id, username, password, email FROM users"
	users := Users{}

	rows, _ := db.Query(sql)

	for rows.Next() {
		user := User{}
		// Obteniendo los valores
		rows.Scan(&user.Id, &user.Username, &user.Password, &user.Email)

		// guardar los registros
		users = append(users, user)
	}

	return users
}

// ! Obtener un registro
func GetUsers(id int) *User {
	user := NewUser("", "", "")

	sql := "SELECT id, username, password, email FROM users WHERE id=?"

	rows, _ := db.Query(sql, id)
	for rows.Next() {
		// Cargando los datos anteriormente vacíos
		rows.Scan(&user.Id, &user.Username, &user.Password, &user.Email)
	}

	return user

}


//! Actualizar Registro
func ( user *User ) update() {
	sql := "UPDATE users SET username=?, password=?, email=? WHERE id=?"
	db.Exec(sql, user.Username, user.Password, user.Email, user.Id)
}

//! Guardar o editar registro
func ( user *User ) Save() {
	// Verificando si id existe en la bd
	if user.Id == 0 {
		user.insert()
	} else {
		user.update()
	}
}


//! Eliminar un registro
func ( user *User ) Delete() {
	sql := "DELETE FROM users WHERE id=?"
	db.Exec( sql, user.Id )
}