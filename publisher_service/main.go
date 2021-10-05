package main

import (
	"os"
	"petegabriel/publisher/internal/api"
	"petegabriel/publisher/pkg/config"
)

var envPath string
func main() {
	if envPath == "" {
		envPath = os.Getenv("ENV_PATH")
	}
	app := api.New(config.New(envPath))
	app.Start()
}
