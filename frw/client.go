package frw

import (
	"log"

	"github.com/gorilla/websocket"
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
		send("----^ NEW MOVE ^----", c.room.forward)
		origin := c.socket.RemoteAddr().String()
		info := c.room.msdb5game.Process(command, origin)
		if info.Err() == nil {
			log.Println(command)
			// Simpler game info sent to everyone
			send(info.ToAll(), c.room.forward)
			// Player info sent to myself only
			send(info.ToMe(), c.send)
		} else {
			log.Println(info.Err())
			send(info.Err().Error(), c.send)
		}
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

func send(message string, to chan []byte) {
	to <- []byte(message)
}
