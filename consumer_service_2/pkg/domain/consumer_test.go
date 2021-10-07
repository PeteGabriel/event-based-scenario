package domain

import (
	"bytes"
	"github.com/streadway/amqp"
	"log"
	"petegabriel/consumer_svc_2/pkg/mocks/data/provider"
	"strings"
	"testing"
)

func TestConsumer_Consume(t *testing.T) {
	mock := provider.EventConsumerMock{}
	mock.GetEventChannelMock = func() (<-chan amqp.Delivery, error) {
		ch := make(chan amqp.Delivery, 1)
		del := amqp.Delivery{
			ContentType: "application/json",
			Body: []byte("{\"event_type\": \"TECH_TASK_NAME\",\"data\":\"good data\"}"),
		}
		ch <- del
		close(ch)
		return ch, nil
	}
	consumer := Consumer{eventConsumer: mock}
	eventsConsumed := make(chan IMessage, 1)
	if err := consumer.Consume(eventsConsumed); err != nil {
		t.Fatalf("error found: %s", err)
	}

	for evt := range eventsConsumed {
		if evt.GetContentType() != "application/json" || evt.GetData() !="good data" {
			t.Fatalf("event different than expected: %s", evt)
		}
	}
}

func TestConsumer_Consume_BadContentType(t *testing.T) {
	//redirect log to catch error message
	var buf bytes.Buffer
	log.SetOutput(&buf)
	expected := "event has invalid content type: text/plain\n"

	mock := provider.EventConsumerMock{}
	mock.GetEventChannelMock = func() (<-chan amqp.Delivery, error) {
		ch := make(chan amqp.Delivery, 1)
		del := amqp.Delivery{
			ContentType: "text/plain",
			Body: []byte("{\"event_type\": \"TECH_TASK_NAME\",\"data\":\"good data\"}"),
		}
		ch <- del
		close(ch)
		return ch, nil
	}
	consumer := Consumer{eventConsumer: mock}
	eventsConsumed := make(chan IMessage, 1)
	if err := consumer.Consume(eventsConsumed); err != nil {
		t.Fatalf("error found: %s", err)
	}



	for range eventsConsumed { } //use it as sync mechanism
	if !strings.Contains(buf.String(), expected) {
		t.Fatalf("Unexpected log message: %s", buf.String())
	}
}

func TestConsumer_Consume_BadEventType(t *testing.T) {
	//redirect log to catch error message
	var buf bytes.Buffer
	log.SetOutput(&buf)
	expected := "consumed event of invalid type: BAD_TYPE\n"

	mock := provider.EventConsumerMock{}
	mock.GetEventChannelMock = func() (<-chan amqp.Delivery, error) {
		ch := make(chan amqp.Delivery, 1)
		del := amqp.Delivery{
			ContentType: "application/json",
			Body: []byte("{\"event_type\": \"BAD_TYPE\",\"data\":\"good data\"}"),
		}
		ch <- del
		close(ch)
		return ch, nil
	}
	consumer := Consumer{eventConsumer: mock}
	eventsConsumed := make(chan IMessage, 1)
	if err := consumer.Consume(eventsConsumed); err != nil {
		t.Fatalf("error found: %s", err)
	}



	for range eventsConsumed { } //use it as sync mechanism
	if !strings.Contains(buf.String(), expected) {
		t.Fatalf("Unexpected log message: %s", buf.String())
	}
}