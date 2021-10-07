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

//App represents the application state and orchestrate the main flow.
type App struct {
	settings *config.Settings
	consumerSvc services.IConsumerService
	broadcasterSvc services.IBroadcastService
}

//New instance of App
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

//Start initiates the main flow of the application. Starts the consumer routine
//where events will be consumed from the event source.
//It starts an additional goroutine that will allow broadcast
//events received to connected clients via websocket.
func (a *App) Start() {

	//events channel will receive events consumed by service
	events := make(chan domain.IMessage)
	if err := a.consumerSvc.ConsumeEvents(events); err != nil {
		close(events)
		log.Fatal("error initiating consumer", err)
	}

	//start a goroutine to broadcast as events arrive.
	go a.startBroadcast(events)

	a.configureRoutes()

	addr := fmt.Sprintf("%s:%s",a.settings.Host, a.settings.Port)
	log.Printf("starting server at %s\n", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal("error initiating web server:", err)
	}
}

func (a *App) configureRoutes() {
	http.Handle("/", a.HandleNewClient())
}

func (a *App) startBroadcast(events <- chan domain.IMessage){
	for evt := range events {
		a.broadcasterSvc.BroadcastNewMessage([]byte(evt.GetData()))
	}
}

