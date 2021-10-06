package domain

import (
	"github.com/pkg/errors"
	"log"
	"petegabriel/consumer_svc_2/pkg/config"
	"petegabriel/consumer_svc_2/pkg/data/provider"
)

type IConsumer interface {
	Consume(chn chan <- Message) error
}

type Consumer struct {
	eventConsumer provider.IEventConsumer
}

func NewConsumer(s *config.Settings) (IConsumer, error){
	ec, err := provider.New(s)
	if err != nil {
		return nil, errors.Wrap(err, "could not create Consumer")
	}
	return &Consumer{
		eventConsumer: ec,
	}, nil
}

func (c *Consumer) Consume(chn chan <- Message) error {
	events, err := c.eventConsumer.GetEventChannel()
	if err != nil {
		return errors.Wrap(err, "could not get channel of events")
	}


	go func() {
		for evt := range events {
			msg := Message{
				ContentType: evt.ContentType,
				Id:          evt.MessageId,
				ConsumerTag: evt.ConsumerTag,
				Content:     evt.Body,
			}
			log.Println("message received:", msg)
			chn <- msg
			evt.Ack(false)
		}
		close(chn)
	}()

	return nil
}