package publisher

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/streadway/amqp"
	"log"
	"petegabriel/publisher/pkg/config"
	"time"
)

type IPublisher interface {
	Publish(content string) error
}

//Publisher represents the entity capable of
//publishing events to an underlying message queue
type Publisher struct {
	channel *amqp.Channel
}

const (
	exchangeName = "events"
	exchangeKind = "topic"
)

//New instance of type IPublisher.
//This function opens a connection to an underlying message queue
//in order to build the instance. In case of not being able to do this,
//and error is returned instead.
func New(s *config.Settings) (IPublisher, error) {

	c, err := setupConn(s)
	if err != nil {
		return nil, err
	}

	return &Publisher{
		channel: c,
	}, nil
}

func setupConn(s *config.Settings) (*amqp.Channel, error) {
	//connection string
	//TODO replace guest with actual values from .env
	url := fmt.Sprintf("amqp://guest:guest@%s:%s/", s.MsgQueueHost, s.MsgQueuePort)
	log.Println("Url:", url)

	// Connect to the rabbitMQ instance
	conn, err := amqp.Dial(url)


	if err != nil {
		return nil, err
	}

	c, err := conn.Channel()
	if err != nil {
		return nil, errors.Wrap(err, "could not open channel.")
	}

	//Declare topology
	err = c.ExchangeDeclare(exchangeName, exchangeKind, true, false, false, false, nil)
	if err != nil {
		return nil, errors.Wrap(err, "could not declare exchange.")
	}

	// We create a queue named Test
	_, err = c.QueueDeclare("test_queue", true, false, false, false, nil)

	if err != nil {
		return nil, errors.Wrap(err, "could not declare queue.")
	}

	// We bind the queue to the exchange to send and receive data from the queue
	err = c.QueueBind("test_queue", "#", exchangeName, false, nil)

	if err != nil {
		return nil, errors.Wrap(err, "could not bind to queue.")
	}

	return c, nil
}

func (p *Publisher) Publish(content string) error {
	msg := amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		Timestamp:    time.Now(),
		ContentType:  "text/plain",
		Body:         []byte(content),
	}

	err := p.channel.Publish(exchangeName, "info", false, false, msg)
	if err != nil {
		return errors.Wrap(err, "could not publish message.")
	}

	return nil
}