package main

import (
	"github.com/BNPrashanth/auth-server/internal/configs"
	"github.com/BNPrashanth/auth-server/internal/logger"
)

func main() {
	// Initialize Viper across the application
	configs.InitializeViper()

	// Initialize Logger across the application
	logger.InitializeZapCustomLogger()

}
