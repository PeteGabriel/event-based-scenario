package services

import (
	"github.com/pkg/errors"
	"petegabriel/consumer_svc_2/pkg/config"
	"petegabriel/consumer_svc_2/pkg/domain"
)

//IConsumerService represents the contract this service respect.
type IConsumerService interface {
	ConsumeEvents(chan domain.IMessage) error
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

//ConsumeEvents requests events to be consumed and sent to the
//events channel parameter.
//It returns an error if it cannot start the consuming routine.
func (cs *ConsumerService) ConsumeEvents(events chan domain.IMessage) error {
	if err := cs.consumer.Consume(events); err != nil {
		return errors.Wrap(err, "IConsumerService:could not start the consuming routine")
	}
	return nil
}