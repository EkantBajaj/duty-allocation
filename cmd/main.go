package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/ekantbajaj/duty-allocation/config"
	"github.com/ekantbajaj/duty-allocation/db"
	"github.com/ekantbajaj/duty-allocation/routers"
	"github.com/gin-contrib/cors"
	"github.com/spf13/viper"
)

func main() {
	err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize the database connection
	err = db.Init()
	if err != nil {
		log.Fatalf("Failed to initialize database connection: %v", err)
	}

	// create go-admin engine
	// Create handlers

	// Initialize the HTTP server
	router := routers.SetupRouter()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	router.Use(cors.New(config))

	// Register routes

	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", viper.GetString("server.host"), viper.GetString("server.port")),
		Handler: router,
	}

	// Start the server in a separate goroutine
	go func() {
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	log.Printf("Server started on %s:%s", viper.GetString("server.host"), viper.GetString("server.port"))

	// Wait for a termination signal to gracefully shut down the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	err = server.Shutdown(nil)
	if err != nil {
		log.Printf("Failed to gracefully shut down server: %v", err)
	}

	log.Println("Server shutdown completed.")
}
