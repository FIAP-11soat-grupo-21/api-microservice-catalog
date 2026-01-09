package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type HealthHandler struct{}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) Health(ctx *gin.Context) {
	now := time.Now().UTC().Add(-3 * time.Hour)
	ctx.JSON(http.StatusOK, gin.H{
		"status":    "Serviço estável - OK - " + now.Format("02/01/2006 - 15:04:05"),
		"timestamp": now.Format(time.RFC3339),
	})
}
