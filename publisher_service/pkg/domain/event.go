package domain


const eventType = "TECH_TASK_NAME"

type Event struct {
	EventType string `json:"event_type"`
	Data      string `json:"data"`
}

func New() *Event {
	return &Event{
		EventType: eventType,
	}
}
