package frw

import (
	"log"

	"github.com/gorilla/websocket"
)

// player represents a player connected to the game.
type player struct {
	// socket is the web socket for this player.
	socket *websocket.Conn
	// send is a channel on which messages are sent.
	send chan []byte
	// room is the room this player is playing in.
	room *GameRoom
}

// Reads commands from UI
func (c *player) read() {
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
		log.Println(command)
		origin := c.socket.RemoteAddr().String()
		info := c.room.msdb5game.Process(command, origin)
		// send("----^ NEW MOVE ^----", c.room.forward)
		if info[0].Err() == nil {
			for _, information := range info {
				// Simpler game info sent to everyone
				c.room.send(information.Dest(), information.Msg())
			}
		} else {
			log.Println(info[0].Err())
			c.room.send(info[0].Err().Error(), info[0].Err().Error())
			// send(info[0].Err().Error(), c.send)
		}
	}
}

// Writes messages to UI
func (c *player) write() {
	defer c.socket.Close()
	for msg := range c.send {
		err := c.socket.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			log.Println("Write to UI error:", err)
		}
	}
}
