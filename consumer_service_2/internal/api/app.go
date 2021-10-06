package api

import (
	"fmt"
	"github.com/pkg/errors"
	"log"
	"net/http"
	"petegabriel/consumer_svc_2/pkg/config"
	"petegabriel/consumer_svc_2/pkg/domain"
	"petegabriel/consumer_svc_2/pkg/services"
)

//App represents the application state. Contains settings provided via external channel
//as well as a register of all clients willing to receive broadcasts.
type App struct {
	settings *config.Settings
	consumerSvc services.IConsumerService
	broadcasterSvc services.IBroadcastService
}

//New instance of type App
func New(set *config.Settings) (*App, error){
	csvc, err := services.NewConsumerService(set)
	if err != nil {
		return nil, errors.Wrap(err, "could not create Consumer service")
	}
	return &App{
		settings: set,
		consumerSvc: csvc,
		broadcasterSvc: services.NewBroadcastService(),
	}, nil
}

func (a App) configureRoutes() {
	http.Handle("/", a.HandleSocketClient())
}

func (a *App) Start() {

	//events channel will receive events consumed by service
	events := make(chan domain.Message)
	a.consumerSvc.ConsumeEvents(events)

	//start a goroutine to broadcast as events arrive.
	go func() {
		for evt := range events {
			if err := a.broadcasterSvc.BroadcastNewMessage(evt.Content); err != nil {
				log.Print("error broadcasting message", err)
			}
		}
	}()

	a.configureRoutes()

	addr := fmt.Sprintf("%s:%s",a.settings.Host, a.settings.Port)
	log.Printf("starting server at %s\n", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal("error initiating web server:", err)
	}
}

