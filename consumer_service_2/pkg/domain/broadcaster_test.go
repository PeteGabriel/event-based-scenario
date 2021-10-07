package domain

import (
	"bytes"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func newWSServer(t *testing.T, h http.Handler) *websocket.Conn {
	t.Helper()
	s := httptest.NewServer(h)
	wsURL := strings.ReplaceAll(s.URL, "http://", "ws://")
	ws, _, err := websocket.DefaultDialer.Dial(wsURL, nil)
	if err != nil {
		t.Fatal(err)
	}
	return ws
}

func TestBroadcaster_EmitBroadcast(t *testing.T) {
	clients := make(map[*Client]bool)

	conn := newWSServer(t, func() http.HandlerFunc {
		var upgrader = websocket.Upgrader{}
		return func(w http.ResponseWriter, r *http.Request) {
			upgrader.Upgrade(w, r, nil)
		}
	}())

	clients[&Client{Conn: conn}] = true
	b := &Broadcaster{clients: clients}

	var buf bytes.Buffer
	log.SetOutput(&buf)

	b.EmitBroadcast([]byte("this nice message"))
	expectedOutput := "message sent to 1 client(s)\n"
	if !strings.Contains(buf.String(), expectedOutput) {
		t.Fatalf("Output received: %s", buf.String())
	}
}

func TestBroadcaster_EmitBroadcast_NoClients(t *testing.T) {
	b := &Broadcaster{clients: make(map[*Client]bool)}

	var buf bytes.Buffer
	log.SetOutput(&buf)

	b.EmitBroadcast([]byte("this nice message"))
	expectedOutput := "message sent to 0 client(s)\n"
	if !strings.Contains(buf.String(), expectedOutput) {
		t.Fatalf("Output received: %s", buf.String())
	}
}

func TestBroadcaster_EmitBroadcast_ClientError(t *testing.T) {
	clients := make(map[*Client]bool)

	conn := newWSServer(t, func() http.HandlerFunc {
		var upgrader = websocket.Upgrader{}
		return func(w http.ResponseWriter, r *http.Request) {
			upgrader.Upgrade(w, r, nil)
		}
	}())

	clients[&Client{Conn: conn}] = true
	b := &Broadcaster{clients: clients}


	var buf bytes.Buffer
	log.SetOutput(&buf)

	conn.Close()//fake a situation when the connection closes abruptly

	b.EmitBroadcast([]byte("this nice message"))
	expectedOutput := "error broadcasting message to client"
	if !strings.Contains(buf.String(), expectedOutput) {
		t.Fatalf("Output received: %s", buf.String())
	}
}