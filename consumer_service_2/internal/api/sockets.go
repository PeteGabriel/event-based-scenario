package api

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

//HandleNewClient returns a function to handle new connections via websocket.
func (a App) HandleNewClient() http.HandlerFunc {

	var upgrader = websocket.Upgrader{
		//allow for localhost testing
		CheckOrigin: func(r *http.Request) bool {
			return a.settings.MsgQueueCheckOrigin
		},
		//limit the maximum length of the message.
		ReadBufferSize: a.settings.ReadBufferSize,
		WriteBufferSize: a.settings.WriteBufferSize,
	}

	return func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Print("Error during connection upgrade:", err)
			return
		}

		//client connected. Register for broadcast later on.
		a.broadcasterSvc.RegisterNewClient(conn)
	}
}

