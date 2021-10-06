package services

import (
	"github.com/pkg/errors"
	"petegabriel/consumer_svc_2/pkg/config"
	"petegabriel/consumer_svc_2/pkg/domain"
)

type IConsumerService interface {
	ConsumeEvents(chan domain.Message)
}

type ConsumerService struct {
	consumer domain.IConsumer
	broadcaster IBroadcastService
}

func NewConsumerService(s *config.Settings) (IConsumerService, error){
	c, err := domain.NewConsumer(s)
	if err != nil {
		return nil, errors.Wrap(err, "could not create Consumer")
	}
	return &ConsumerService{
		consumer: c,
	}, nil
}

func (cs *ConsumerService) ConsumeEvents(events chan domain.Message) {
	if err := cs.consumer.Consume(events); err != nil {
		//todo log message
		return
	}
}