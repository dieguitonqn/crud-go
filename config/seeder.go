package config

import (
	"log"

	"CRUD_GO/internal/usuarios/infrastructure"

	"gorm.io/gorm"
)

func SeedDatabase(db *gorm.DB) {
	// Create sample users
	users := []infrastructure.User{
		{Name: "John Doe", Email: "john.doe@example.com"},
		{Name: "Jane Smith", Email: "jane.smith@example.com"},
	}

	// Insert sample users into the database
	for _, user := range users {
		err := infrastructure.CreateUser(db, &user)
		if err != nil {
			log.Fatalf("Failed to seed user %s: %v", user.Name, err)
		}
	}
}

func seedDevData() error {

	// Crear usuarios de prueba
	users := []infrastructure.User{
		{Name: "Admin", Email: "admin@test.com"},
		{Name: "Test User", Email: "test@test.com"},
	}

	return DB.Create(&users).Error
}
