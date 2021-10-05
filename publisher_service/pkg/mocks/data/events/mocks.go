package events

type MessengerMock struct {
	PublishMessageMock  func(msg []byte) error
}

func (mock MessengerMock) PublishMessage(msg []byte) error {
	return mock.PublishMessageMock(msg)
}