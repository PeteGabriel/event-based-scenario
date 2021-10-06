package provider

import (
	"github.com/pkg/errors"
	"github.com/streadway/amqp"
	"petegabriel/consumer_svc_2/pkg/config"
)

const consumerName = "event_consumer"

type IEventConsumer interface {
	GetEventChannel() (<-chan amqp.Delivery, error)
	StopConsumeEvent() error
}

type EventConsumer struct {
	c *amqp.Channel
	s *config.Settings
}

func New(s *config.Settings) (IEventConsumer, error) {
	ec := &EventConsumer{
		s: s,
	}
	if err := ec.openConnection(); err != nil {
		return nil, err
	}

	return ec, nil
}

func (ec *EventConsumer) GetEventChannel() (<-chan amqp.Delivery, error)  {
	// We consume data in the queue named test using the channel we created in go.
	msgs, err := ec.c.Consume(ec.s.MsgQueueName, consumerName, false, false, false, false, nil)

	if err != nil {
		return nil, errors.Wrap(err, "IEventConsumer:could not consume events from queue")
	}

	return msgs, nil
}

func (ec *EventConsumer) StopConsumeEvent() error {
	//cancel consumer and gracefully end goroutine
	if err := ec.c.Cancel(consumerName, false); err != nil {
		return errors.Wrap(err, "IEventConsumer:could not stop consumer channel")
	}
	return nil
}

func (ec *EventConsumer) openConnection() error {
	conn, err := amqp.Dial(ec.s.MsgQueueConnString)
	if err != nil {
		return errors.Wrap(err, "IEventConsumer:could not dial successfully")
	}

	c, err := conn.Channel()
	if err != nil {
		return errors.Wrap(err, "IEventConsumer:could not open channel")
	}

	if err = c.ExchangeDeclare(ec.s.MsgQueueExchangeName, ec.s.MsgQueueExchangeKind, true, false, false, false, nil); err != nil {
		return errors.Wrap(err, "IEventConsumer:could not declare exchange")
	}

	//create queue
	if  _, err = c.QueueDeclare(ec.s.MsgQueueName, true, false, false, false, nil); err != nil {
		return errors.Wrap(err, "IEventConsumer:could not declare queue.")
	}

	// bind queue to exchange key and name
	if err = c.QueueBind(ec.s.MsgQueueName, ec.s.MsgQueueRoutingKey, ec.s.MsgQueueExchangeName, false, nil); err != nil {
		return errors.Wrap(err, "IEventConsumer:could not bind to queue.")
	}

	ec.c = c

	return nil
}