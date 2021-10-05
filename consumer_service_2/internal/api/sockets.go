package api

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"petegabriel/consumer_svc_2/pkg/domain"
)


//HandleSocketClient returns a function to handle new connections via websocket.
func (a App) HandleSocketClient() http.HandlerFunc {

	var upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { //allow for localhost testing
			return true
		},
		ReadBufferSize: a.settings.ReadBufferSize,
		WriteBufferSize: a.settings.WriteBufferSize,
	}

	return func(w http.ResponseWriter, r *http.Request) {
		// Upgrade our raw HTTP connection to a websocket based one
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Print("Error during connection upgrade:", err)
			return
		}

		//client connected. Register for broadcast later on.
		a.clients[&domain.Client{Conn: conn}] = true

		defer conn.Close()

		// The event loop
		for {

			//TODO consume message from queue whenever ready
			msg := []byte("")

			//TODO validate it matches our internal JSON message structure

			//TODO send a broadcast to all clients
			for client := range a.clients {
				if err = client.Conn.WriteMessage(1, msg); err != nil {
					log.Print("error notifying client about new message:", err)
				}
			}
		}
	}
}

