package frw

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/nikiforosFreespirit/msdb5/app/game"
)

// Action interface
type Action interface {
	Process(request, origin string) []*game.Info
}

// GameRoom struct
type GameRoom struct {
	// forward is a channel that holds incoming messages
	// that should be forwarded to the other players.
	forward chan []byte
	// join is a channel for players wishing to join the room.
	join chan *player
	// leave is a channel for players wishing to leave the room.
	leave chan *player
	// players holds all current players in this room.
	players map[*player]bool
	// msdb5 game instance
	msdb5game Action
}

// NewGameRoom makes a new room.
func NewGameRoom(side bool) *GameRoom {
	return &GameRoom{
		forward:   make(chan []byte),
		join:      make(chan *player),
		leave:     make(chan *player),
		players:   make(map[*player]bool),
		msdb5game: game.NewGame(side),
	}
}

// Run func
func (r *GameRoom) Run() {
	for {
		select {
		case player := <-r.join:
			// joining
			r.players[player] = true
		case player := <-r.leave:
			// leaving
			delete(r.players, player)
			close(player.send)
		case msg := <-r.forward:
			// forward message to all players
			for player := range r.players {
				player.send <- msg
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

func (r *GameRoom) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	socket, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Fatal("ServeHTTP:", err)
		return
	}
	player := &player{
		socket: socket,
		send:   make(chan []byte, messageBufferSize),
		room:   r,
	}
	r.join <- player
	player.send <- []byte("Enter name and connect")
	defer func() { r.leave <- player }()
	go player.write()
	player.read()
}

func (r *GameRoom) send(destination, message string) {
	for pl := range r.players {
		if pl.socket.RemoteAddr().String() != destination {
			continue
		}
		pl.send <- []byte(message)
		break
	}
	// func send(message string, to chan []byte) {

	// Simpler game info sent to everyone
	// send(info.Msg(), c.room.forward)
	// Player info sent to myself only
	// send(info.ToMe(), c.send)

	// to <- []byte(message)
}
