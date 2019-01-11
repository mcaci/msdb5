package main

import (
	"log"
	"strings"

	"github.com/gorilla/websocket"
	"github.com/nikiforosFreespirit/msdb5/card"
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

func (c *client) read() {
	defer c.socket.Close()
	for {
		_, msg, err := c.socket.ReadMessage()
		if err != nil {
			return
		}
		c.room.forward <- msg
	}
}
func (c *client) write() {
	defer c.socket.Close()
	for msg := range c.send {
		c.actualWrite(msg)
		info := strings.Split(string(msg), " ")
		myCard, _ := card.ByName(info[0], info[1])
		c.actualWrite([]byte(myCard.String()))
		// if err != nil {
		// 	return
		// }
	}
}

func (c *client) actualWrite(msg []byte) {
	err := c.socket.WriteMessage(websocket.TextMessage, msg)
	if err != nil {
		log.Println("Actual Write Error:", err)
	}
}
