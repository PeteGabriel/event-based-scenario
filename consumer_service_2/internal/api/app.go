package api

import (
	"fmt"
	"log"
	"net/http"
	"petegabriel/consumer_svc_2/pkg/config"
	"petegabriel/consumer_svc_2/pkg/domain"
)

//App represents the application state. Contains settings provided via external channel
//as well as a register of all clients willing to receive broadcasts.
type App struct {
	settings *config.Settings
	clients map[*domain.Client]bool
}

//New instance of type App
func New(set *config.Settings) *App{
	return &App{
		settings: set,
		clients: make(map[*domain.Client]bool),
	}
}

func (a App) configureRoutes() {
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

