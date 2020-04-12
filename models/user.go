package models

import (
	"errors"
	"fmt"
)

/*
* el paquete models se utiliza para generar cambios en el estado de la aplicación
 */

type User struct {
	ID        int
	FirstName string
	LastName  string
}

var (
	users []*User // * un slice que tendra como valor un pointer que apunta al struct User
	// * Esto permite utilizar cada item, sin necesidad de hacer una copia o compartir los datos
	nextID = 1 // no requiere sintaxis implicita
)

// función que retorna los usuarios almacenados
func GetUsers() []*User {
	return users
}

// Agrega un usuario a la colección
func AddUser(u User) (User, error) {
	if u.ID != 0 {
		return User{}, errors.New("New user must not include id or it must be set in zero")
	}
	u.ID = nextID // se asigna la variable
	nextID++
	users = append(users, &u)
	return u, nil
}

func GetUserById(id int) (User, error) {
	for _, u := range users {
		if u.ID == id {
			return *u, nil // nil es se devuelve como signo de que no hubo errores
			// referiéndose al pointer retornado
		}
	}
	return User{}, fmt.Errorf("User with ID '%v' not found", id)
	// * Errorf permite retornar strings como objetos de error,
	// * Interpolando el string %v en el valor del id
}

func UpdateUser(u User) (User, error) {
	for i, candidate := range users {
		if candidate.ID == u.ID {
			users[i] = &u
			return u, nil
		}
	}
	return User{}, fmt.Errorf("User with ID '%v' not found", u.ID)
}

func RemoveUserById(id int) error {
	for i, u := range users {
		if u.ID == id {
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("User with ID '%v' not found", id)
}
