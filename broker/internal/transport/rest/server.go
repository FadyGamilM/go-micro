package rest

import (
	"broker/config"
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateServer(r *gin.Engine) *http.Server {
	configs, err := config.LoadServerConfigs("./config")
	if err != nil {
		log.Println("couldn't read the configs")
		os.Exit(1)
	}
	return &http.Server{
		Addr:    configs.Server.Port,
		Handler: r,
	}
}

func InitServer(srv *http.Server) {
	if err := srv.ListenAndServe(); err != nil {
		log.Println("couldn't initate a server")
		os.Exit(1)
	}

	log.Println("[Broker-Microservice] | up and running on port : ", srv.Addr)
}

func ShutdownGracefully(server *http.Server) {
	// Define a context with timeout.
	// This timeout is the time available for the server to finish
	// whatever requests are running currently before being forced to shut down.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	log.Println("The server is shutting down now ...")

	// Start closing all connections.
	// If the context that is passed is expired, we will receive an error.
	if err := server.Shutdown(ctx); err != nil {
		// This will be done only if there is an error while trying to shut down the server.
		log.Fatalf("Server is forced to shutdown: %v\n", err)
	}

	// Log a message indicating successful shutdown.
	log.Println("Server has gracefully shutdown.")
}
