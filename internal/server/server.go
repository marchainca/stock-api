package server

import (
	"github.com/gin-gonic/gin"
	"github.com/marchainca/stock-api/internal/config"
	"github.com/marchainca/stock-api/internal/stock"
)

func New(cfg config.Config) *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	cli := stock.NewClient(cfg.BaseURL, cfg.User, cfg.Password)
	svc := stock.NewService(cli)
	stock.Register(r, svc)

	// health
	r.GET("/healthz", func(c *gin.Context) { c.Status(200) })
	return r
}

func Run(r *gin.Engine, port string) error {
	return r.Run(":" + port)
}
