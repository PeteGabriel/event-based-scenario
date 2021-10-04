package main

import (
	"petegabriel/publisher/internal/api"
	"petegabriel/publisher/pkg/config"
)

func main() {
	app := api.New(config.New())
	app.Start()
}
