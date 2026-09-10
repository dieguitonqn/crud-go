package handlers

import (
	"net/http"

	"CRUD_GO/internal/usuarios/service"

	"github.com/gin-gonic/gin"
)

func SignUp(c *gin.Context) {
	// Aquí irá la lógica para registrar un nuevo usuario
	email := c.PostForm("email")
	password := c.PostForm("password")

	if err := service.SignUp(email, password); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Registro de usuario exitoso",
	})
}

func SignIn(c *gin.Context) {
	// Aquí irá la lógica para iniciar sesión de un usuario
	email := c.PostForm("email")
	password := c.PostForm("password")

	if err := service.SignIn(email, password); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Inicio de sesión exitoso",
	})
}

func SignOut(c *gin.Context) {
	// Aquí irá la lógica para cerrar sesión de un usuario
	email := c.PostForm("email")

	if err := service.SignOut(email); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Cierre de sesión exitoso",
	})
}
