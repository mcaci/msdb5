package ui

import (
	"log"

	"github.com/gorilla/websocket"
	"github.com/nikiforosFreespirit/msdb5/api"
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
			return
		}
		// log action
		log.Println(msg)
		// execute action
		command := string(msg)
		run(c.room.msdb5board, command, c.socket.RemoteAddr().String())
		// TODO: format command with info for others: command with public info for board
		send(command, c.room.forward) // to room
		// TODO: format board with info for myself / my hand and collected cards
		status := c.room.msdb5board.String()
		send(status, c.send) // to myself
	}
}

// Writes messages to UI
func (c *client) write() {
	defer c.socket.Close()
	for msg := range c.send {
		err := c.socket.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			log.Println("Write Error:", err)
		}
	}
}

func run(room api.Action, command, origin string) {
	room.Action(command, origin)
}

func send(message string, to chan []byte) {
	to <- []byte(message)
}
