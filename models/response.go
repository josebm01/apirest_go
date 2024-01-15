package models

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Estructura para responder al cliente
type Response struct {
	Status      int         `json:"status"`
	Data        interface{} `json:"data"`
	Message     string      `json:"message"`
	contentType string
	respWrite    http.ResponseWriter
}


//! Estructura de respuesta por defecto
func CreateDefaultResponse( rw http.ResponseWriter ) Response {
	return Response {
		Status: http.StatusOK,
		respWrite: rw,
		contentType: "application/json",
	}
}


//! Estructura de respuesta para el cliente
func ( resp *Response ) Send() {
	// Header 
	resp.respWrite.Header().Set("Content-Type", resp.contentType)
	resp.respWrite.WriteHeader(resp.Status)

	// puntero de resp - con toda la información (status, data, mensaje)
	output, _ := json.Marshal(&resp)
	fmt.Fprintln( resp.respWrite, string(output) )
}


//! Devolver respuesta al cliente 
func SendData( rw http.ResponseWriter, data interface{}) {
	response := CreateDefaultResponse(rw)
	response.Message = "data was obtained"
	response.Data = data
	response.Send()
}



//! Método para manejar errores de listar, eliminar o al obtener un dato
func ( resp *Response ) NotFound(){
	resp.Status = http.StatusNotFound
	resp.Message = "Resource Not Found"
}

func SendNoFound( rw http.ResponseWriter ){
	response := CreateDefaultResponse(rw)
	response.NotFound()
	response.Send()
}


//! Manejar error al actualizar o ingresar un registro
func ( resp *Response ) UnprocessableEntity(){
	resp.Status = http.StatusUnprocessableEntity
	resp.Message = "UnprocessableEntity Not Found"
}

func SendUnprocessableEntity( rw http.ResponseWriter ) {
	response := CreateDefaultResponse(rw)
	response.UnprocessableEntity()
	response.Send()
}