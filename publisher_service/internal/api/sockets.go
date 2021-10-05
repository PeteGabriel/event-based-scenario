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
		// Upgrade our raw HTTP connection to a websocket based one
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Print("Error during connection upgrade:", err)
			return
		}

		defer conn.Close()

		// The event loop
		for {
			msgType, message, err := conn.ReadMessage()
			if err != nil {
				log.Println("Error during message reading:", err)
				break
			}
			log.Printf("Received: %s", message)

			if err = a.publishService.PublishNewEvent(message); err != nil {
				log.Println("Error publishing new message:", err)
				if err = conn.WriteMessage(msgType, []byte(err.Error())); err != nil {
					log.Println("Error warning client about publishing error:", err)
				}
			}

		}
	}
}
