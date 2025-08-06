// Package config centraliza la carga y validación de variables de entorno.
package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

// Config agrupa toda la configuración necesaria para arrancar el servicio.
type Config struct {
	Port string
	Env  string // p.ej. development | staging | production

	API struct {
		BaseURL  string
		User     string
		Password string
	}

	// --- Base de datos ---
	DBHost string
	DBPort string
	DBUser string
	DBPass string
	DBName string
}

// Load lee .env, .env.<ENV> (si existe) y devuelve una Config validada.
// Si falta alguna variable obligatoria, retorna error.
func Load() (Config, error) {
	// Cargar el archivo base .env.
	_ = godotenv.Load()

	// Identificar el entorno.
	env := os.Getenv("ENV")

	// Sobrecarga con .env.<ENV> si corresponde anula/añade variables.
	if env != "" {
		_ = godotenv.Overload(".env." + env) // también ignora si no existe
	}

	var cfg Config
	cfg.Env = env // podría quedar vacío si no se definió

	// Variables obligatorias
	var err error

	if cfg.Port, err = required("PORT"); err != nil {
		return cfg, err
	}
	if cfg.API.BaseURL, err = required("API_BASE_URL"); err != nil {
		return cfg, err
	}
	if cfg.API.User, err = required("API_USER"); err != nil {
		return cfg, err
	}
	if cfg.API.Password, err = required("API_PASSWORD"); err != nil {
		return cfg, err
	}
	if cfg.DBHost, err = required("DB_HOST"); err != nil {
		return cfg, err
	}
	if cfg.DBPort, err = required("DB_PORT"); err != nil {
		return cfg, err
	}
	if cfg.DBUser, err = required("DB_USER"); err != nil {
		return cfg, err
	}
	if cfg.DBPass, err = required("DB_PASS"); err != nil {
		return cfg, err
	}
	if cfg.DBName, err = required("DB_NAME"); err != nil {
		return cfg, err
	}

	return cfg, nil
}

// required devuelve el valor de la variable o un error si está ausente o vacía.
func required(key string) (string, error) {
	if v, ok := os.LookupEnv(key); ok && strings.TrimSpace(v) != "" {
		return v, nil
	}
	return "", fmt.Errorf("config: env var %s is required", key)
}
