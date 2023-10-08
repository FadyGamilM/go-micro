package handlers

import (
	"auth/internal/core"
	"auth/internal/core/dtos"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type signupHandler struct {
	service core.UserService
}

type signupHandlerConfig struct {
	R       *gin.Engine
	Service core.UserService
}

func NewSignupHandler(shc *signupHandlerConfig) *signupHandler {
	h := &signupHandler{
		service: shc.Service,
	}

	hRoutes := shc.R.Group("/auth")

	hRoutes.POST("/signup", h.HandleSignup)

	return h
}

func (sh *signupHandler) HandleSignup(c *gin.Context) {
	var dto dtos.SignupDto
	if err := c.ShouldBind(&dto); err != nil {
		log.Printf("[Signup Handler] | %v \n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}
	user, err := sh.service.Signup(c, &dto)
	if err != nil {
		log.Printf("[Signup Handler] | %v \n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
	}

	c.JSON(http.StatusAccepted, gin.H{
		"data": user,
	})
}
