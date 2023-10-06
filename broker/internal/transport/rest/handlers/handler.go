package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type handler struct{}

type HandlerConfig struct {
	R *gin.Engine
}

func New(hc *HandlerConfig) *handler {
	h := &handler{}

	hRoutes := hc.R.Group("/broker")

	hRoutes.POST("/", h.HandleRequests)

	return h
}

func (h *handler) HandleRequests(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": "you hit the broker service",
	})
}
