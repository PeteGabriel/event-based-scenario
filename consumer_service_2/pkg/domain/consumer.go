package domain

import (
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/streadway/amqp"
	"log"
	"petegabriel/consumer_svc_2/pkg/config"
	"petegabriel/consumer_svc_2/pkg/data/provider"
)

type IConsumer interface {
	Consume(chn chan <- IMessage) error
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

//Consume retrieves a channel from queue and starts a new goroutine
//that will traverse the channel eventually consuming messages.
//It returns an error in case it cannot obtain the channel from queue.
//Events consumed are sent to the output channel parameter. If the goroutine is
//terminated it will call Close on the output channel.
func (c *Consumer) Consume(output chan <- IMessage) error {
	events, err := c.eventConsumer.GetEventChannel()
	if err != nil {
		return errors.Wrap(err, "could not get channel of events")
	}

	go handleChannelOfEvents(events, output)

	return nil
}

func handleChannelOfEvents(events <-chan amqp.Delivery, output chan <- IMessage) {
	for e := range events {
		//expect json content type
		if e.ContentType != "application/json" {
			log.Printf("event has invalid content type: %s\n", e.ContentType)
			continue
		}

		//validate this is expected data
		var evt Event
		err := json.Unmarshal(e.Body, &evt)
		if err != nil {
			log.Println("received invalid data:", err)
			continue
		}
		//validate event kind
		if !evt.IsEventTypeOk() {
			log.Printf("consumed event of invalid type: %s\n", evt.EventType)
			continue
		}

		log.Printf("event received with content %s\n", string(e.Body))

		//communicate new event
		output <- &Message{
			ContentType: e.ContentType,
			Timestamp:   e.Timestamp.String(),
			Event: evt.Data,
		}
		//always acknowledge the event received
		e.Ack(false)
	}
	close(output)
}