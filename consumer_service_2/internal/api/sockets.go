package api

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
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
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Print("Error during connection upgrade:", err)
			return
		}

		//client connected. Register for broadcast later on.
		a.broadcasterSvc.RegisterNewClient(conn)

	}
}

