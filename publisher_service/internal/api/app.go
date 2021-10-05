package api

import (
	"fmt"
	"log"
	"net/http"
	"petegabriel/publisher/pkg/config"
	"petegabriel/publisher/pkg/services"
)

type App struct {
	settings *config.Settings
	publishService services.IPublishService
}

func New(set *config.Settings) *App{
	return &App{
		settings: set,
		publishService: services.New(set),
	}
}

func (a App) configureRoutes() {
	//upload route
	http.Handle("/", a.HandleSocketClient())
}

func (a *App) Start() {
	a.configureRoutes()

	addr := fmt.Sprintf("%s:%s",a.settings.Host, a.settings.Port)
	log.Printf("starting server at %s\n", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal("error initiating web server:", err)
	}
}