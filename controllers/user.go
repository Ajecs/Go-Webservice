package controllers

/*
	El paquete controllers permite el manejo de peticiones http para recibir y procesar
	- a continuación se crean métodos que permitiran modificar el comportamiento
*/

import (
	"encoding/json"
	"net/http"
	"regexp"
	"strconv"

	"github.com/pluralsight/webservice/models"
)

// * se declara un struct vacío que se enlaza con datos
type userController struct {
	userIDPattern *regexp.Regexp
}

// al enlazar la función a un struct(objeto) la convertimos en un método ->
/*
* Se enlaza el parametro uc al struct userController
* 'uc' es el equivalente al uso de 'this' en otros lenguajes
 */

func (uc userController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// lo anterior pertenece al paquete net/http. el primer parametro es la respuesta web
	if r.URL.Path == "/users" {
		// De ser users se trabaja con la colección entera
		switch r.Method {
		case http.MethodGet:
			uc.getAll(w, r)
		case http.MethodPost:
			uc.post(w, r)
		default:
			w.WriteHeader(http.StatusNotImplemented)
		}
	} else {
		// si no se trabaja con un solo objeto Usuario
		matches := uc.userIDPattern.FindStringSubmatch(r.URL.Path)
		if len(matches) == 0 {
			w.WriteHeader(http.StatusNotFound)
		}
		id, err := strconv.Atoi(matches[1])
		// convierte los valores ingresados a string
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
		}
		switch r.Method {
		case http.MethodGet:
			uc.get(id, w)
		case http.MethodPut:
			uc.put(id, w, r)
		case http.MethodDelete:
			uc.delete(id, w)
		default:
			w.WriteHeader(http.StatusNotImplemented)
		}

	}

	// code w.Write([]byte("Hello from the User Controller!"))
	// Write imprime los datos de la conexión como parte de una respuesta http
	// El método Write recibe un parametro que es un slice de bytes producto de la
	// interacción usualmente con una red externa
	// Con byte puedo convertir los bytes en un string
}

func (uc *userController) getAll(w http.ResponseWriter, r *http.Request) {
	encodeResponseAsJSON(models.GetUsers(), w)
	// * Función que se encuentra en front.go
}

func (uc *userController) get(id int, w http.ResponseWriter) {
	u, err := models.GetUserById(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	encodeResponseAsJSON(u, w)
}

func (uc *userController) post(w http.ResponseWriter, r *http.Request) {
	u, err := uc.parseRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Could not parse User object"))
		return
	}
	u, err = models.AddUser(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	encodeResponseAsJSON(u, w)
}

func (uc *userController) put(id int, w http.ResponseWriter, r *http.Request) {
	// * maneja la actualización de datos ingresados
	u, err := uc.parseRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Could not parse user object"))
		return
	}
	if id != u.ID {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("ID of submitted user must match ID in URL"))
		return
	}
	u, err = models.UpdateUser(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	encodeResponseAsJSON(u, w)
}

func (uc *userController) delete(id int, w http.ResponseWriter) {
	err := models.RemoveUserById(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (uc *userController) parseRequest(r *http.Request) (models.User, error) {
	// * convierte los datos en representación JSON de las funciones en el objeto JSON User
	dec := json.NewDecoder(r.Body)
	var u models.User
	err := dec.Decode(&u)
	if err != nil {
		return models.User{}, err
	}
	return u, nil
}

/*
	* Go no posee una clásico paradigma orientado a objetos.
	* No posee un constructor, sino una función que cumple ese rol
	- Por convención la fución constructora comienza con new seguido del nombre del
	- objeto a configurar
*/
func newUserController() *userController {
	// devolverá un nuevo Controlador de usuario sin necesidad de una copia
	return &userController{
		// Al apuntar a un struct no se necesita de una variable que almacene el pointer
		// Al devolver la variable userController Go reconoce dicha como local y la sube al
		// nivel necesario.
		userIDPattern: regexp.MustCompile(`^/users/(\d+)/?`),
	}
}
