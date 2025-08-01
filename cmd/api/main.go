package main

import (
	"log"

	"github.com/marchainca/stock-api/internal/config" // Exponer variables de entorno
	"github.com/marchainca/stock-api/internal/server" // Expone New(cfg) para crear e inyectar dependencias
	// en el router y Run(r, port) para iniciar http.ListenAndServe.
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err) // termina con código 1 si la config no es válida
	}

	r := server.New(cfg) // inyecta la config
	if err := server.Run(r, cfg.Port); err != nil {
		log.Fatal(err)
	}
}
