package services

import (
	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
	"log"
	"petegabriel/consumer_svc_2/pkg/domain"
)

type IBroadcastService interface {
	RegisterNewClient(conn *websocket.Conn)
	BroadcastNewMessage(content []byte) error
}

type BroadcastService struct {
	caster domain.IBroadcaster
}

func NewBroadcastService() IBroadcastService {
	return &BroadcastService{
		caster: domain.New(),
	}
}

//RegisterNewClient for broadcast later on
func (bs *BroadcastService) RegisterNewClient(conn *websocket.Conn) {
	bs.caster.RegisterNewClient(&domain.Client{Conn: conn})
	log.Println("new client registered")
}

//BroadcastNewMessage broadcast a message to all registered clients.
func (bs *BroadcastService) BroadcastNewMessage(content []byte) error{
	log.Println("broadcasting new message")
	if err := bs.caster.EmitBroadcast(content); err != nil {
		return errors.Wrap(err, "IBroadcastService: error emitting new broadcast")
	}
	return nil
}