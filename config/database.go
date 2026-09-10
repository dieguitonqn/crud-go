package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// ConnectDatabase establece la conexión con PostgreSQL
func ConnectDatabase() {
	// Cargar variables de entorno
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️  No se encontró archivo .env, usando variables del sistema")
	}

	// Obtener configuración
	cfg := getDBConfig()

	// Construir DSN
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		cfg.host,
		cfg.user,
		cfg.password,
		cfg.dbName,
		cfg.port,
		cfg.sslMode,
		cfg.timeZone,
	)

	// Conectar con GORM
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(getLogLevel()),
		NowFunc: func() time.Time {
			return time.Now().Local()
		},
	})
	if err != nil {
		log.Fatal("❌ Error conectando a PostgreSQL:", err)
	}

	// Configurar pool de conexiones
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatal("❌ Error obteniendo conexión SQL:", err)
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	log.Println("✅ Conectado a PostgreSQL exitosamente")
}

// dbConfig estructura interna para configuración
type dbConfig struct {
	host     string
	user     string
	password string
	dbName   string
	port     string
	sslMode  string
	timeZone string
}

// getDBConfig obtiene configuración de variables de entorno
func getDBConfig() dbConfig {
	return dbConfig{
		host:     getEnv("DB_HOST", "localhost"),
		user:     getEnv("DB_USER", "admin"),
		password: getEnv("DB_PASSWORD", "secret123"),
		dbName:   getEnv("DB_NAME", "mi_db"),
		port:     getEnv("DB_PORT", "5432"),
		sslMode:  getEnv("DB_SSLMODE", "disable"),
		timeZone: getEnv("DB_TIMEZONE", "America/Santiago"),
	}
}

// getLogLevel determina el nivel de logging según entorno
func getLogLevel() logger.LogLevel {
	env := os.Getenv("ENV")
	if env == "production" {
		return logger.Warn
	}
	return logger.Info
}

// getEnv obtiene variable de entorno o valor por defecto
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
