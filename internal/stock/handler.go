package stock

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	svc  *Service
	repo Repository
}

func Register(r *gin.Engine, svc *Service, repo Repository) {
	h := &Handler{svc: svc, repo: repo}
	r.GET("/stocks", h.list)
}

func (h *Handler) list(c *gin.Context) {
	cursor := c.Query("next")
	res, err := h.svc.List(c.Request.Context(), cursor)
	if err != nil {
		// respuesta uniforme del error
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}
	// Persistir los items
	if perr := h.repo.SaveItems(c.Request.Context(), res.Items); perr != nil {
		log.Printf("save items: %v", perr)
	}
	c.JSON(http.StatusOK, res)
}
