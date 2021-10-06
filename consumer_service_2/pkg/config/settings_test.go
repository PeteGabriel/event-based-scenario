package config

import "testing"

func TestSettings(t *testing.T) {

	s := New("../../test.env")

	if s.Host != "localhost" {
		t.Errorf("expected host: localhost. Value received: %s", s.Host)
		t.FailNow()
	}

	if s.Port != "8208" {
		t.Errorf("expected port: 8208. Value received: %s", s.Port)
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

	if s.MsgQueueCheckOrigin != true {
		t.Errorf("expected queue origin check: TRUE. Value received: %t", s.MsgQueueCheckOrigin)
		t.FailNow()
	}

	if s.MsgQueueConnString != "amqp://guest:guest@localhost:5672/" {
		t.Errorf("expected queue conn string: amqp://guest:guest@localhost:5672/. Value received: %s", s.MsgQueueConnString)
		t.FailNow()
	}

	if s.MsgQueueName != "tech_task" {
		t.Errorf("expected queue name: tech_task. Value received: %s", s.MsgQueueName)
		t.FailNow()
	}

	if s.MsgQueueRoutingKey != "info" {
		t.Errorf("expected queue routing key: info. Value received: %s", s.MsgQueueName)
		t.FailNow()
	}

}