package domain

import (
	"encoding/json"
	"log"
)

type Event struct {
	EventType string `json:"event_type"`
	Data      string `json:"data"`
}

const eventName = "TECH_TASK_NAME"

func (e *Event) IsEventTypeOk() bool { return e.EventType == eventName }

type IMessage interface {
	GetContentType() string
	GetTimestamp() string
	GetData() string
}

//Message contains some data related to the event that represents
//apart from the actual event data.
type Message struct {
	ContentType string `json:"content_type"`
	Timestamp string `json:"timestamp"`
	Event string `json:"data"`
}

func (m *Message) GetContentType() string { return m.ContentType }
func (m *Message) GetData() string { return m.Event }
func (m *Message) GetTimestamp() string { return m.Timestamp }

func (m Message) String() string{
	j, err := json.Marshal(m)
	if err != nil {
		log.Fatalln(err)
	}
	return string(j)
}

