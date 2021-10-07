package domain

import (
	"github.com/pkg/errors"
	"log"
)

type IBroadcaster interface {
	RegisterNewClient(client *Client)
	EmitBroadcast(msg []byte)
}

type Broadcaster struct {
	clients map[*Client]bool
}

func New() IBroadcaster {
	return &Broadcaster{clients: make(map[*Client]bool)}
}

//RegisterNewClient registers a new client for later access
//while broadcasting.
func (b *Broadcaster) RegisterNewClient(client *Client) {
	b.clients[client] = true
}

//EmitBroadcast sends the content given by parameter
//to all registered clients.
func (b *Broadcaster) EmitBroadcast(msg []byte) {
	counter := 0
	for client := range b.clients {
		// a more robust solution would take care of different types.
		txtType := 1
		if err := client.Conn.WriteMessage(txtType, msg); err != nil {
			//log error with current connection. allow others to try to write message
			log.Println(errors.Wrap(err, "error broadcasting message to client"))
		}else {
			counter = counter + 1
		}
	}
	log.Printf("message sent to %d client(s)\n", counter )
}