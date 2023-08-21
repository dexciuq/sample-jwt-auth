package main

import (
	"github.com/dexciuq/sample-jwt-auth/config"
	"github.com/dexciuq/sample-jwt-auth/internal/app"
	"log"
)

func main() {
	cfg, err := config.LoadConfiguration()
	if err != nil {
		log.Fatalf("Config error: %s", err)
		return
	}
	app.Run(cfg)
}
