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
}

// Load lee .env, .env.<ENV> (si existe) y devuelve una Config validada.
// Si falta alguna variable obligatoria, retorna error.
func Load() (Config, error) {
	// 1) Carga el archivo base .env.
	_ = godotenv.Load()

	// 2) Identifica el entorno (puede venir del .env recién cargado).
	env := os.Getenv("ENV")

	// 3) Sobrecarga con .env.<ENV> si corresponde (anula/añade variables).
	if env != "" {
		_ = godotenv.Overload(".env." + env) // también ignora si no existe
	}

	var cfg Config
	cfg.Env = env // podría quedar vacío si no se definió

	// 4) Variables obligatorias
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

	return cfg, nil
}

// required devuelve el valor de la variable o un error si está ausente o vacía.
func required(key string) (string, error) {
	if v, ok := os.LookupEnv(key); ok && strings.TrimSpace(v) != "" {
		return v, nil
	}
	return "", fmt.Errorf("config: env var %s is required", key)
}
