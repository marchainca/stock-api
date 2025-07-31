package stock

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct{ svc *Service }

func Register(r *gin.Engine, svc *Service) {
	h := &Handler{svc}
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
	c.JSON(http.StatusOK, res)
}
