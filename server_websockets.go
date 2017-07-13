package cistatus

import (
	"time"

	"github.com/gorilla/websocket"
)

const (
	writeWait      = 10 * time.Second
	pongWait       = 60 * time.Second
	pingPeriod     = (pongWait * 2) / 3
	maxMessageSize = 1024 * 1024
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  maxMessageSize,
	WriteBufferSize: maxMessageSize,
}

// wsHub manages web socket connections and communications
type wsHub struct {
	broadcast  chan Summary
	register   chan *wsSubscriber
	unregister chan *wsSubscriber

	subscribers   map[*wsSubscriber]bool
	lastBroadcast Summary
}

// newWSHub creates a new wsHub
func newWSHub() *wsHub {
	return &wsHub{
		subscribers: make(map[*wsSubscriber]bool),
		broadcast:   make(chan Summary),
		register:    make(chan *wsSubscriber),
		unregister:  make(chan *wsSubscriber),
	}
}

// run starts the wsHub recieving on it's channels
func (h *wsHub) run() {
	for {
		select {
		case s := <-h.register:
			h.subscribers[s] = true
			if h.lastBroadcast.Color != "" {
				s.write(h.lastBroadcast)
			}
			break

		case c := <-h.unregister:
			_, ok := h.subscribers[c]
			if ok {
				delete(h.subscribers, c)
				close(c.send)
			}
			break

		case s := <-h.broadcast:
			h.send(s)
			h.lastBroadcast = s
			break
		}
	}
}

// send sends the summary to all subscribers
func (h *wsHub) send(summary Summary) {
	for s := range h.subscribers {
		select {
		case s.send <- summary:
			break

		default:
			go func() {
				h.unregister <- s
			}()
		}
	}
}

// wsSubscriber represents an individual websocket connection to the server
type wsSubscriber struct {
	ws   *websocket.Conn
	send chan Summary
}

// newWSSubscriber creates a populated wsSubscriber
func newWSSubscriber(ws *websocket.Conn) *wsSubscriber {
	return &wsSubscriber{
		ws:   ws,
		send: make(chan Summary),
	}
}

// writePump handles incomming messages from the send channel to and
// deliverers them to clients and sends ping messages.
func (s *wsSubscriber) writePump() {
	ticker := time.NewTicker(pingPeriod)

	defer ticker.Stop()
	defer s.ws.Close()

	for {
		select {
		case message, ok := <-s.send:
			if !ok {
				break
			}
			err := s.write(message)
			if err != nil {
				return
			}
		case <-ticker.C:
			err := s.ping()
			if err != nil {
				return
			}
		}
	}
}

// write sends the summary to the client websocket
func (s *wsSubscriber) write(summary Summary) error {
	s.ws.SetWriteDeadline(time.Now().Add(writeWait))
	return s.ws.WriteJSON(summary)
}

// close sends the websocket close signal to the client
func (s *wsSubscriber) close() error {
	return s.ws.WriteControl(websocket.CloseMessage, []byte{}, time.Now().Add(writeWait))
}

// close sends a ping to the client
func (s *wsSubscriber) ping() error {
	return s.ws.WriteControl(websocket.PingMessage, []byte{}, time.Now().Add(writeWait))
}
