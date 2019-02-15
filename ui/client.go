package main

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
	room *room
}

// Reads commands from UI
func (c *client) read() {
	defer c.socket.Close()
	for {
		_, msg, err := c.socket.ReadMessage()
		if err != nil {
			return
		}
		c.room.msdb5board.Action(string(msg), c.socket.RemoteAddr().String())
		sendMessage(msg, c.room.forward)
		sendMessage([]byte(c.room.msdb5board.String()), c.send)
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

func sendMessage(message []byte, roomChan chan []byte) {
	roomChan <- message
}
