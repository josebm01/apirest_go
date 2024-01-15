package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// username:password@tcp(localhost:3306)/database
const url = "root:jose@tcp(127.0.0.1:3306)/goweb_db"
// const url = "root:jose@tcp(localhost:3306)/goweb_db"

// Guarda la conexión
var db *sql.DB

// ! Realizando conexión
func Connect() {
	// Recibe nombre del driver y la ruta
	connection, err := sql.Open("mysql", url)
	if err != nil {
		fmt.Println("Connection could not be established")
		panic(err)
	}

	fmt.Println("Successful connection to database")

	// asignando la conexión
	db = connection
}

// ! Cerrar la conexión
func Close() {
	db.Close()
}

// ! Verificar conexión
func Ping() {
	// Si está conectada la bd no manda error pero si o lo está manda error
	if err := db.Ping(); err != nil {
		panic(err)
	}
}

//! Verificando si existe tabla
func ExistsTable( tableName string ) bool {
	sql := fmt.Sprintf("SHOW TABLES LIKE '%s'", tableName)
	rows, err := db.Query(sql)
	if err != nil {
		fmt.Println("Error:", err)
	}

	// valor booleano 
	return rows.Next()
}



//! Crea tabla
func CreateTable( schema string, name string ) {
	
	// comprobando si existe la tabla
	if !ExistsTable(name) {
		fmt.Println("Trying to create table ...")
		if _, err := db.Exec(schema); err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println("Created table ")
	}
	
}


//! Reiniciar datos de una tabla
func TruncateTable( tableName string) {
	sql := fmt.Sprintf("TRUNCATE %s", tableName)
	Exec(sql)
}



//! Polimorfismo de Exec
func Exec( query string, args ...interface{} ) (sql.Result, error) {
	result, err := db.Exec(query, args...)
	if err != nil {
		fmt.Println( err )
	} 

	return result, err
}


//! Polimorfismo de Query
func Query( query string, args ...interface{} ) (*sql.Rows, error) {
	rows, err := db.Query(query, args...)
	if err != nil {
		fmt.Println( err )
	} 

	return rows, err
}


