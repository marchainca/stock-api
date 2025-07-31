package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Modo Release por defecto en prod; Debug en dev
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	// Ruta básica de salud
	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	// Puerto obtenido de var de entorno con fallback
	addr := ":8580"
	log.Printf("🚀 Server listening on %s", addr)
	if err := r.Run(addr); err != nil {
		log.Fatal(err)
	}
}
