package services

import (
	"github.com/pkg/errors"
	"log"
	"petegabriel/publisher/pkg/config"
	"petegabriel/publisher/pkg/data/events"
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
	if err := p.msn.PublishMessage(message); err != nil {
		return errors.Wrap(err, "IMessenger:could not publish event.")
	}
	return nil
}