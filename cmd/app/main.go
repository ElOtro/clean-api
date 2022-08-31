package main

import (
	"log"

	"github.com/ElOtro/clean-api/config"
	"github.com/ElOtro/clean-api/internal/app"
)

func main() {
	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Run
	app.Run(cfg)
}
