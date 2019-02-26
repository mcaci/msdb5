package ui

import (
	"log"

	"github.com/gorilla/websocket"
	"github.com/nikiforosFreespirit/msdb5/api"
	"github.com/nikiforosFreespirit/msdb5/board"
	"github.com/nikiforosFreespirit/msdb5/player"
)

// client represents a single chatting user.
type client struct {
	// socket is the web socket for this client.
	socket *websocket.Conn
	// send is a channel on which messages are sent.
	send chan []byte
	// room is the room this client is chatting in.
	room *Room
}

// Reads commands from UI
func (c *client) read() {
	defer c.socket.Close()
	for {
		// read player input
		_, msg, err := c.socket.ReadMessage()
		if err != nil {
			log.Println("Error from reading UI input:", err)
			return
		}
		// execute action
		command := string(msg)
		origin := c.socket.RemoteAddr().String()
		run(c.room.msdb5board, command, origin)
		// Simpler board info sent to everyone
		b, _ := c.room.msdb5board.(*board.Board)
		send(command, c.room.forward)
		send(b.Print(), c.room.forward)
		// Player info sent to myself only
		p, _ := b.Players().Find(func(p *player.Player) bool { return p.Host() == origin })
		send(p.Print(), c.send)
	}
}

// Writes messages to UI
func (c *client) write() {
	defer c.socket.Close()
	for msg := range c.send {
		err := c.socket.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			log.Println("Write to UI error:", err)
		}
	}
}

func run(room api.Action, command, origin string) {
	room.Action(command, origin)
}

func send(message string, to chan []byte) {
	to <- []byte(message)
}
