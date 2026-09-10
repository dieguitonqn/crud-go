package routes

import (
	"CRUD_GO/internal/usuarios/handlers"

	"github.com/gin-gonic/gin"
)

// SetupRoutes sets up the routes for the application
func SetupRoutes(router *gin.Engine) {
	// Acá se definen las rutas de toda la aplicación

	router.GET("/users", handlers.GetUsers)
	router.POST("/user", handlers.NewUser)

	// Aquí se pueden agregar más rutas según sea necesario

}
