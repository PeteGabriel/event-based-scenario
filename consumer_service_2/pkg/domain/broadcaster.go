package domain

import (
	"github.com/pkg/errors"
)

type IBroadcaster interface {
	RegisterNewClient(client *Client)
	EmitBroadcast(msg []byte) error
}

type Broadcaster struct {
	clients map[*Client]bool
}

func New() IBroadcaster {
	return &Broadcaster{clients: make(map[*Client]bool)}
}

func (b *Broadcaster) RegisterNewClient(client *Client) {
	b.clients[client] = true
}

func (b *Broadcaster) EmitBroadcast(msg []byte) error {
	for client := range b.clients {
		//TODO review this msg type
		if err := client.Conn.WriteMessage(1, msg); err != nil {
			return errors.Wrap(err, "error broadcasting message to client")
		}
	}
	return nil
}