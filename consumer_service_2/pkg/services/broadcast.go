package services

import (
	"github.com/gorilla/websocket"
	"log"
	"petegabriel/consumer_svc_2/pkg/domain"
)

//IBroadcastService represents the contract this service respects.
//Defines operations related to the broadcast of events to clients.
type IBroadcastService interface {
	RegisterNewClient(conn *websocket.Conn)
	BroadcastNewMessage(content []byte)
}

type BroadcastService struct {
	caster domain.IBroadcaster
}

func NewBroadcastService() IBroadcastService {
	return &BroadcastService{
		caster: domain.New(),
	}
}

//RegisterNewClient for broadcast of events
func (bs *BroadcastService) RegisterNewClient(conn *websocket.Conn) {
	bs.caster.RegisterNewClient(&domain.Client{Conn: conn})
	log.Println("new client registered")
}

//BroadcastNewMessage broadcast content to all registered clients.
func (bs *BroadcastService) BroadcastNewMessage(content []byte){
	log.Println("broadcasting new message")
	bs.caster.EmitBroadcast(content)
}