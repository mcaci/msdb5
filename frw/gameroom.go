package frw

import (
	"log"
	"net/http"

	"golang.org/x/text/language"
	"golang.org/x/text/message"

	"github.com/gorilla/websocket"
	"github.com/mcaci/msdb5/app/game"
	"github.com/mcaci/msdb5/app/msg"
)

// GameRoom struct
type GameRoom struct {
	// commandChan is a channel that holds incoming messages
	// that should be forwarded to the other players.
	commandChan chan playerCommand
	// join is a channel for players wishing to join the room.
	join chan *playerClient
	// leave is a channel for players wishing to leave the room.
	leave chan *playerClient
	// players holds all current players in this room.
	players map[*playerClient]bool
	// msdb5 game instance
	msdb5game Action
	// lang language tag
	lang language.Tag
}

// NewGameRoom makes a new room.
func NewGameRoom(side bool, lang language.Tag) *GameRoom {
	return &GameRoom{
		commandChan: make(chan playerCommand),
		join:        make(chan *playerClient),
		leave:       make(chan *playerClient),
		players:     make(map[*playerClient]bool),
		msdb5game:   game.NewGame(side),
		lang:        lang,
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
		case m := <-r.commandChan:
			// commandChan message to all players
			data := r.msdb5game.Process(m.request, m.origin)
			msg.Notify(data, r.lang, m.request, m.origin)
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
	player := r.joinWith(socket)
	defer func() { r.leave <- player }()
	go player.write()
	player.read()
}

func (r *GameRoom) joinWith(socket *websocket.Conn) *playerClient {
	playerChannel := make(chan []byte, messageBufferSize)
	r.msdb5game.Join(socket.RemoteAddr().String(), playerChannel)
	player := &playerClient{
		socket:      socket,
		infoChannel: playerChannel,
		room:        r,
	}
	r.join <- player
	printer := message.NewPrinter(r.lang)
	player.infoChannel <- []byte(printer.Sprintf("Enter name and connect"))
	return player
}
