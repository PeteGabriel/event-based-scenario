package main

import (
	"log"
	"os"
	"petegabriel/consumer_svc_2/internal/api"
	"petegabriel/consumer_svc_2/pkg/config"
)

var envPath string

func main() {
	if envPath == "" {
		envPath = os.Getenv("ENV_PATH")
	}
	app, err := api.New(config.New(envPath))
	if err != nil {
		log.Fatalln(err)
	}
	app.Start()
}