package server

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/marchainca/stock-api/internal/config"
	"github.com/marchainca/stock-api/internal/db"
	"github.com/marchainca/stock-api/internal/stock"
)

func New(cfg config.Config) *gin.Engine { // Crea y configura el router de Gin.

	if cfg.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery()) // Loggea peticiones y evita que un panic derrumbe el proceso.
	d := db.New(cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPass, cfg.DBName)

	httpc := &http.Client{Timeout: 10 * time.Second}
	cli := stock.NewClient(cfg.API.BaseURL, cfg.API.User, cfg.API.Password, httpc)
	svc := stock.NewService(cli)
	repo := stock.NewRepo(d.DB)
	stock.Register(r, svc, repo)

	// health -> Facilita readiness/liveness probes en Docker/K8s.
	r.GET("/healthz", func(c *gin.Context) { c.Status(200) })
	return r
}

func Run(r *gin.Engine, port string) error { // expone el puerto del servicio
	return r.Run(":" + port)
}
