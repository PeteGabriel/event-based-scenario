package domain

import "petegabriel/consumer_svc_2/pkg/domain"

type BroadcasterMock struct {
	RegisterNewClientMock func(*domain.Client)
	EmitBroadcastMock func([]byte)
}

func (mock BroadcasterMock) RegisterNewClient(c *domain.Client) {
	mock.RegisterNewClientMock(c)
}

func (mock BroadcasterMock) EmitBroadcast(m []byte)  {
	mock.EmitBroadcastMock(m)
}


type ConsumerMock struct {
	ConsumeMock func(chan <- domain.IMessage) error
}

func (mock ConsumerMock) Consume(c chan <- domain.IMessage) error {
	return mock.ConsumeMock(c)
}