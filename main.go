package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/TomasCruz/grpc-web-fahrenheit/client"
	"github.com/TomasCruz/grpc-web-fahrenheit/configuration"
	"github.com/TomasCruz/grpc-web-fahrenheit/model"
)

func main() {
	// populate configuration
	config := setupFromEnvVars()
	routes := newRoutes()

	// set gRPC client interface to service
	clientInterface, err := client.InitializeClient(config.GrpcHost, config.GrpcPort)
	if err != nil {
		log.Fatalf("failed to initialize client: %s", err)
	}
	model.SetClient(clientInterface)

	// graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, os.Kill)

	// bindRoutes
	bindRoutes(routes)

	// fire up the server
	var httpServer *http.Server
	log.Printf("starting web server on :%s", config.Port)

	go func() {
		httpServer = &http.Server{Addr: fmt.Sprintf(":%s", config.Port)}
		if err := httpServer.ListenAndServe(); err != nil {
			log.Fatalf("failed to start HTTP server: %s", err)
		}
	}()

	<-stop
	gracefulShutdown(clientInterface, httpServer)
}

func setupFromEnvVars() (config configuration.Config) {
	config.Port = readAndCheckIntEnvVar("GRPC_WEB_PORT")
	config.GrpcPort = readAndCheckIntEnvVar("GRPC_PORT")
	config.GrpcHost = readAndCheckEnvVar("GRPC_HOST")
	return
}

func gracefulShutdown(client model.Client, httpServer *http.Server) {
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdownCancel()

	client.Shutdown()
	httpServer.Shutdown(shutdownCtx)
	fmt.Println("Graceful shutdown")
}
