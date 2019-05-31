package main

import (
	"log"
	"net/http"

	"github.com/BNPrashanth/auth-server/internal/configs"
	"github.com/BNPrashanth/auth-server/internal/logger"
	"github.com/BNPrashanth/auth-server/internal/services"
)

func main() {
	// Initialize Viper across the application
	configs.InitializeViper()

	// Initialize Logger across the application
	logger.InitializeZapCustomLogger()

	http.HandleFunc("/", services.HandleIndex)
	http.HandleFunc("/signin", services.Signin)
	http.HandleFunc("/welcome", services.Welcome)
	http.HandleFunc("/refresh", services.Refresh)

	log.Fatal(http.ListenAndServe(":9090", nil))
}
