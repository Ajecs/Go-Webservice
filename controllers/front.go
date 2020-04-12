package controllers

import (
	"encoding/json"
	"io"
	"net/http"
)

/*
este archivo configura todo el enrutamiento de la aplicación.
Permite que cuando la solicitud de red al ser recibida irá al
controlador correcto para ser procesada
*/

func RegisterControllers() {
	uc := newUserController()

	http.Handle("/users", *uc)
	http.Handle("/users/", *uc)
	// manejara la solicitudes que provengan de users y
	// de users + otro subdominio
}

func encodeResponseAsJSON(data interface{}, w io.Writer) {
	// a partir del paquete JSON de Go
	enc := json.NewEncoder(w)
	enc.Encode(data)
}
