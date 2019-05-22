package frw

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/nikiforosFreespirit/msdb5/app"
	"github.com/nikiforosFreespirit/msdb5/app/game"
)

// Room struct
type Room struct {
	// forward is a channel that holds incoming messages
	// that should be forwarded to the other clients.
	forward chan []byte
	// join is a channel for clients wishing to join the room.
	join chan *client
	// leave is a channel for clients wishing to leave the room.
	leave chan *client
	// clients holds all current clients in this room.
	clients map[*client]bool
	// msdb5 game instance
	msdb5game app.Action
}

// NewRoom makes a new room.
func NewRoom(side bool) *Room {
	return &Room{
		forward:   make(chan []byte),
		join:      make(chan *client),
		leave:     make(chan *client),
		clients:   make(map[*client]bool),
		msdb5game: game.NewAction(side),
	}
}

// Run func
func (r *Room) Run() {
	for {
		select {
		case client := <-r.join:
			// joining
			r.clients[client] = true
		case client := <-r.leave:
			// leaving
			delete(r.clients, client)
			close(client.send)
		case msg := <-r.forward:
			// forward message to all clients
			for client := range r.clients {
				client.send <- msg
			}
		}
	}
}

const (
	socketBufferSize  = 1024
	messageBufferSize = 256
)

var upgrader = &websocket.Upgrader{ReadBufferSize: socketBufferSize,
	WriteBufferSize: socketBufferSize}

func (r *Room) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	socket, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Fatal("ServeHTTP:", err)
		return
	}
	client := &client{
		socket: socket,
		send:   make(chan []byte, messageBufferSize),
		room:   r,
	}
	r.join <- client
	client.send <- []byte("Enter name and connect")
	defer func() { r.leave <- client }()
	go client.write()
	client.read()
}
