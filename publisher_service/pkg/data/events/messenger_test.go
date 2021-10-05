package events

import (
	"petegabriel/publisher/pkg/config"
	"testing"
)

func TestNewMessenger(t *testing.T) {
	s := config.New("../../../test.env")

	if _, err := NewMessenger(s); err != nil {
		t.Errorf("error creating messenger: %s", err)
		t.FailNow()
	}
}

func TestNewMessenger_CannotOpenConnection(t *testing.T) {
	s := config.New("../../../empty.env")

	if _, err := NewMessenger(s); err == nil {
		t.Error("error expected to be found")
		t.FailNow()
	}
}

func TestMessenger_PublishMessage(t *testing.T) {
	s := config.New("../../../test.env")

	m, err := NewMessenger(s)
	if err != nil {
		t.FailNow()
	}


	if err = m.PublishMessage([]byte("")); err != nil {
		t.Errorf("error publishing via messenger: %s", err)
		t.FailNow()
	}
}
