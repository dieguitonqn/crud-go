package service

import (
	"errors"
)

func SignUp(email, password string) error {
	// Aquí irá la lógica para registrar un nuevo usuario
	// Por ahora, solo devolvemos un error si el email o la contraseña están vacíos
	if email == "" || password == "" {
		return errors.New("email y contraseña son requeridos")
	}
	return nil
}

func SignIn(email, password string) error {
	// Aquí irá la lógica para iniciar sesión de un usuario
	// Por ahora, solo devolvemos un error si el email o la contraseña están vacíos
	if email == "" || password == "" {
		return errors.New("email y contraseña son requeridos")
	}
	return nil
}
func SignOut(email string) error {
	// Aquí irá la lógica para cerrar sesión de un usuario
	// Por ahora, solo devolvemos un error si el email está vacío
	if email == "" {
		return errors.New("email es requerido")
	}
	return nil
}
