package provider

import (
	"petegabriel/consumer_svc_2/pkg/config"
	"testing"
)

func TestNew(t *testing.T) {
	s := config.New("../../../test.env")

	if _, err := New(s); err != nil {
		t.Errorf("error creating consumer: %s", err)
		t.FailNow()
	}
}

func TestMessenger_PublishMessage(t *testing.T) {
	s := config.New("../../../test.env")

	ec, err := New(s)
	if err != nil {
		t.FailNow()
	}


	if err = ec.ConsumeEvent(); err != nil {
		t.Errorf("error consuming event from queue: %s", err)
		t.FailNow()
	}
}