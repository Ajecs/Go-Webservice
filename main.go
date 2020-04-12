package main

import (
	"net/http"

	"github.com/pluralsight/webservice/controllers"
)

func main() {
	// registrar el enrutamiento
	controllers.RegisterControllers()
	http.ListenAndServe(":3001", nil)
	// recibe dos parametros.
	// - el 1ro: la dirección ip del servidor
	// - el 2do el servidor multiplexor que cumplira la función de controlador frontal
	// manejando todas las solicitudes, mientras que el controlador trasero(userController)
	// maneja determinada solicitud filtrada por el frontal
}

// * el paquete main puede ejecutarse con 'go run .'
