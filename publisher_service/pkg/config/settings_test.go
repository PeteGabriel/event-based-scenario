package config

import (
	"testing"
)

func TestSettings(t *testing.T) {

	s := New("../../test.env")

	if s.Host != "localhost" {
		t.Errorf("expected host: localhost. Value received: %s", s.Host)
		t.FailNow()
	}

	if s.Port != "8200" {
		t.Errorf("expected port: 8200. Value received: %s", s.Port)
		t.FailNow()
	}

	if s.ReadBufferSize != 1024 {
		t.Errorf("expected read buffer size: 1024. Value received: %d", s.ReadBufferSize)
		t.FailNow()
	}

	if s.WriteBufferSize != 1024 {
		t.Errorf("expected write buffer size: 1024. Value received: %d", s.WriteBufferSize)
		t.FailNow()
	}

	if s.MsgQueueHost != "localhost" {
		t.Errorf("expected queue hostname: localhost. Value received: %s", s.MsgQueueHost)
		t.FailNow()
	}

	if s.MsgQueuePort != "5672" {
		t.Errorf("expected queue port: 5672. Value received: %s", s.MsgQueuePort)
		t.FailNow()
	}

	if s.MsgQueueCheckOrigin != true {
		t.Errorf("expected queue origin check: TRUE. Value received: %t", s.MsgQueueCheckOrigin)
		t.FailNow()
	}

	if s.MsgQueueUser != "guest" {
		t.Errorf("expected queue username: guest. Value received: %s", s.MsgQueueUser)
		t.FailNow()
	}

	if s.MsgQueuePassword != "guest" {
		t.Errorf("expected queue password: guest. Value received: %s", s.MsgQueuePassword)
		t.FailNow()
	}

	if s.MsgQueueName != "tech_task" {
		t.Errorf("expected queue name: tech_task. Value received: %s", s.MsgQueueName)
		t.FailNow()
	}
}
