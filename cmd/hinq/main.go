package main

import (
	"file_management/internal/cmd"
	"file_management/internal/config"
	"file_management/internal/utils/logger"
)

func main() {
	// Initialize logger
	log := logger.New()

	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}

	// Execute root command
	cmd.Execute(cfg, log)
}
