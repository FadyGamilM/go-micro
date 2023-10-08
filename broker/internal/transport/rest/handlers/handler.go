package handlers

import (
	"broker/internal/core"
	"bytes"
	"encoding/json"
	"log"
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

	hRoutes.POST("/handle", h.HandleRequests)

	return h
}

func (h *handler) HandleRequests(c *gin.Context) {
	log.Println("you hit the broker service")

	var reqDto *core.RequestPaylaod
	if err := c.ShouldBind(&reqDto); err != nil {
		log.Printf("[Broker Handler] | %v \n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	switch reqDto.Action {
	case "auth":
		h.Handle_AuthSignup(c, &core.AuthSignupPayload{
			Email:    reqDto.Auth.Email,
			Password: reqDto.Auth.Password,
		})
	default:
		log.Println("[Broker Handler] | Unknown action")
		c.JSON(http.StatusBadGateway, gin.H{
			"error": "Unknown action",
		})
	}
}

func (h *handler) Handle_AuthSignup(c *gin.Context, payload *core.AuthSignupPayload) {
	// call the auth microservice
	// -> convert the payload data type to json for sending it to the request body
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		log.Printf("[Broker Handler - Auth (Signup)] | %v \n", err)
	}

	// -> define a request
	request, err := http.NewRequest("POST", "http://auth-srv/auth/signup", bytes.NewBuffer(jsonPayload))
	if err != nil {
		log.Printf("[Broker Handler - Auth (Signup)] | %v \n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
	}

	// -> call the microservice via a http client
	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		log.Printf("[Broker Handler - Auth (Signup)] | %v \n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
	}
	// we need to close the body and clean resources
	defer response.Body.Close()

	// check response
	if response.StatusCode != http.StatusAccepted {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
	}

	// -> decode the response into appropriate type
	var responseDto *core.AuthResponse

	if err = json.NewDecoder(response.Body).Decode(responseDto); err != nil {
		log.Printf("[Broker Handler - Auth (Signup)] | %v \n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
	}

	var result *core.ResponsePayload
	result.Error = false
	result.Message = "authenticated"
	result.Payload = responseDto

	c.JSON(http.StatusAccepted, gin.H{
		"data": result,
	})

}
