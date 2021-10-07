package provider

import "github.com/streadway/amqp"

type EventConsumerMock struct {
	GetEventChannelMock func() (<-chan amqp.Delivery, error)
	StopConsumeEventMock func() error
}

func (mock EventConsumerMock) GetEventChannel () (<-chan amqp.Delivery, error){
	return mock.GetEventChannelMock()
}

func (mock EventConsumerMock) StopConsumeEvent () error{
	return mock.StopConsumeEventMock()
}


