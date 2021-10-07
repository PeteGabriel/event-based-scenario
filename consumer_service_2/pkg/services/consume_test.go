package services

import (
	"github.com/pkg/errors"
	domain2 "petegabriel/consumer_svc_2/pkg/domain"
	"petegabriel/consumer_svc_2/pkg/mocks/domain"
	"testing"
	"time"
)

func TestConsumerService_ConsumeEvents(t *testing.T) {

	tests := []struct {
		ContentType string
		Event string
	}{
		{"text/plain", "event number 0"},
		{"text/plain", "event number 1"},
		{"text/plain", "event number 2"},
	}

	mock := domain.ConsumerMock{}
	mock.ConsumeMock = func(events chan<- domain2.IMessage) error {
		events <- &domain2.Message {
			ContentType: "text/plain",
			Timestamp: time.Now().String(),
			Event:       "event number 0",
		}
		events <- &domain2.Message {
			ContentType: "text/plain",
			Timestamp: time.Now().String(),
			Event:       "event number 1",
		}
		events <- &domain2.Message {
			ContentType: "text/plain",
			Timestamp: time.Now().String(),
			Event:       "event number 2",
		}

		return nil
	}

	cs := ConsumerService{consumer: mock}

	evts := make(chan domain2.IMessage)

	go func(){
		_ = cs.ConsumeEvents(evts)
	}()

	//Closing a channel indicates that no more values will be sent on it.
	//useful here to communicate completion to the channel’s receivers.
	close(evts)
	for evt := range evts {
		found := false
		for _, t := range tests {
			found = t.Event == evt.GetData() && t.ContentType == evt.GetContentType() && evt.GetTimestamp() != ""
			if found {
				break
			}
		}
		if !found {
			t.Fatalf("consumed event expected to be found: %s",evt)
		}
	}
}

func TestConsumerService_ConsumeEvents_ErrorReturned(t *testing.T) {
	mock := domain.ConsumerMock{}
	mock.ConsumeMock = func(events chan<- domain2.IMessage) error {
		return errors.New("error consuming events from queue")
	}

	cs := ConsumerService{consumer: mock}

	evts := make(chan domain2.IMessage)
	defer close(evts)

	if err := cs.ConsumeEvents(evts); err == nil {
		t.Fatalf("error was expected to be found: %s", err)
	}
}