package events

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/streadway/amqp"
	"petegabriel/publisher/pkg/config"
	"time"
)

//IMessenger represents the contract this layer provides to
//publish messages to an external interface.
//Replacing this implementation while respecting the contract
//allows replacing how the external communication is done
//without implicating other layers.
type IMessenger interface {
	PublishMessage(msg []byte) error
}

type Messenger struct {
	c *amqp.Channel
	s *config.Settings
}

//NewMessenger creates a new instance of IMessenger
func NewMessenger(s *config.Settings) (IMessenger, error) {
	m := &Messenger{
		s: s,
	}
	if err := m.openConnection(); err != nil {
		return nil, err
	}

	return m, nil
}

//PublishMessage publishes a new event with content given by parameter `msg`.
func (m *Messenger) PublishMessage(msg []byte) error{
	evt := amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		Timestamp:    time.Now(),
		ContentType:  "text/plain",
		Body:         msg,
	}

	if err := m.c.Publish(m.s.MsgQueueExchangeName, m.s.MsgQueueRoutingKey, false, false, evt); err != nil {
		return errors.Wrap(err, "IMessenger:error publishing to queue '" + m.s.MsgQueueName + "'")
	}

	return nil
}


func (m *Messenger) openConnection() error{
	//connection string
	url := fmt.Sprintf("amqp://%s:%s@%s:%s/", m.s.MsgQueueUser, m.s.MsgQueuePassword, m.s.MsgQueueHost, m.s.MsgQueuePort)

	// Connect to the rabbitMQ instance
	conn, err := amqp.Dial(url)
	if err != nil {
		return errors.Wrap(err, "IMessenger:could not dial successfully.")
	}

	c, err := conn.Channel()
	if err != nil {
		return errors.Wrap(err, "IMessenger:could not open channel.")
	}

	//Declare topology
	if err = c.ExchangeDeclare(m.s.MsgQueueExchangeName, m.s.MsgQueueExchangeKind, true, false, false, false, nil); err != nil {
		return errors.Wrap(err, "IMessenger:could not declare exchange.")
	}

	//create queue
	if  _, err = c.QueueDeclare(m.s.MsgQueueName, true, false, false, false, nil); err != nil {
		return errors.Wrap(err, "IMessenger:could not declare queue.")
	}

	// bind queue to exchange key and name
	if err = c.QueueBind(m.s.MsgQueueName, m.s.MsgQueueRoutingKey, m.s.MsgQueueExchangeName, false, nil); err != nil {
		return errors.Wrap(err, "IMessenger:could not bind to queue.")
	}

	m.c = c //assign channel for future communications

	return nil
}