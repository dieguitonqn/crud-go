package config

import (
	"log"
	"os"

	"CRUD_GO/internal/usuarios/infrastructure"
)

// RunMigrations ejecuta las migraciones según el entorno
func RunMigrations() error {
	env := os.Getenv("ENV")
	if env == "" {
		env = "development"
	}

	log.Printf("🔄 Ejecutando migraciones en entorno: %s", env)

	// Solo en desarrollo: reset y seed
	if env == "development" || env == "test" {
		return runDevMigration()
	}

	// En producción: solo migrar (sin borrar)
	return runProdMigration()
}

func runDevMigration() error {
	log.Println("🧹 Limpiando base de datos de desarrollo...")

	// ⚠️ SOLO en desarrollo
	if err := DB.Exec("DROP SCHEMA public CASCADE; CREATE SCHEMA public;").Error; err != nil {
		return err
	}

	log.Println("📦 Creando tablas...")
	if err := DB.AutoMigrate(
		&infrastructure.User{},
		// Agregar más modelos aquí
	); err != nil {
		return err
	}

	log.Println("🌱 Sembrando datos de prueba...")
	return seedDevData()
}

func runProdMigration() error {
	log.Println("📦 Ejecutando migración en producción...")

	// Solo AutoMigrate (no borra datos)
	if err := DB.AutoMigrate(
		&infrastructure.User{},
	); err != nil {
		return err
	}

	log.Println("✅ Migración completada")
	return nil
}
