package services

import (
	"github.com/pkg/errors"
	"petegabriel/publisher/pkg/config"
	"petegabriel/publisher/pkg/mocks/data/events"
	"testing"
)

func TestPublishService_PublishNewEvent(t *testing.T) {
	msnMock := events.MessengerMock{}
	msnMock.PublishMessageMock = func(msg []byte) error {
		return nil
	}

	p := PublishService{msn: msnMock}
	if err := p.PublishNewEvent("new message to be published"); err != nil {
		t.Errorf("error publishing from service: %s", err)
		t.FailNow()
	}
}

func TestPublishService_PublishNewEvent_WithError(t *testing.T) {
	msnMock := events.MessengerMock{}
	msnMock.PublishMessageMock = func(msg []byte) error {
		return errors.New("some error while publishing the message")
	}

	p := PublishService{msn: msnMock}
	if err := p.PublishNewEvent("new message to be published"); err == nil {
		t.Error("error expected to be found")
		t.FailNow()
	}
}

func TestNew(t *testing.T) {
	p := New(config.New("../../test.env"))

	if p == nil {
		t.Error("found nil instead of proper instance")
		t.FailNow()
	}
}