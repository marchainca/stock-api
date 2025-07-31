// cmd/api/main.go
package main

import (
	"log"

	"github.com/marchainca/stock-api/internal/config"
	"github.com/marchainca/stock-api/internal/server"
)

func main() {
	cfg := config.Load()
	r := server.New(cfg)
	log.Printf("🚀 API listening on :%s", cfg.Port)
	if err := server.Run(r, cfg.Port); err != nil {
		log.Fatal(err)
	}
}
