package services

import (
	"encoding/json"
	"github.com/pkg/errors"
	"log"
	"petegabriel/publisher/pkg/config"
	"petegabriel/publisher/pkg/data/events"
	"petegabriel/publisher/pkg/domain"
)

type IPublishService interface {
	PublishNewEvent(message []byte) error
}

type PublishService struct {
	msn events.IMessenger
}

func New(s *config.Settings)  IPublishService{
	m, err := events.NewMessenger(s)
	if err != nil {
		log.Fatalln("PublishService:", err)
	}
	return &PublishService{
		msn: m,
	}
}

func (p *PublishService) PublishNewEvent(message []byte) error{
	event := domain.New()
	event.Data = string(message)
	//serialize
	evtJSON, err := json.Marshal(event)
	if err != nil {
		return errors.Wrap(err, "IPublisher:could not serialize event")
	}
	if err := p.msn.PublishMessage(evtJSON); err != nil {
		return errors.Wrap(err, "IPublisher:could not publish event")
	}
	log.Printf("event published: %s\n", evtJSON)
	return nil
}