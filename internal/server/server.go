package server

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/marchainca/stock-api/internal/config"
	"github.com/marchainca/stock-api/internal/stock"
)

func New(cfg config.Config) *gin.Engine { // Crea y configura el router de Gin.

	if cfg.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery()) // Loggea peticiones y evita que un panic derrumbe el proceso.

	httpc := &http.Client{Timeout: 10 * time.Second}
	cli := stock.NewClient(cfg.API.BaseURL, cfg.API.User, cfg.API.Password, httpc)
	svc := stock.NewService(cli)
	stock.Register(r, svc)

	// health -> Facilita readiness/liveness probes en Docker/K8s.
	r.GET("/healthz", func(c *gin.Context) { c.Status(200) })
	return r
}

func Run(r *gin.Engine, port string) error { // expone el puerto del servicio
	return r.Run(":" + port)
}
