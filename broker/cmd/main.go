package main

import (
	"broker/internal/transport/rest"
	"broker/internal/transport/rest/handlers"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	router := rest.CreateRouter()

	handlers.New(&handlers.HandlerConfig{R: router})

	server := rest.CreateServer(router)

	rest.InitServer(server)

	// listen for shutdown or any interrupts
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	// wait for it
	<-quit
	// if we here, thats mean we will shut down the server gracefully
	rest.ShutdownGracefully(server)

}
