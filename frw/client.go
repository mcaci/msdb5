package frw

import (
	"log"

	"github.com/gorilla/websocket"
	"github.com/nikiforosFreespirit/msdb5/app"
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
		infoForAll, infoForPlayer, err := run(c.room.msdb5game, command, origin)
		if err == nil {
			log.Println(command)
			// Simpler board info sent to everyone
			send(infoForAll, c.room.forward)
			// Player info sent to myself only
			send(infoForPlayer, c.send)
		} else {
			log.Println(err)
			send(err.Error(), c.send)
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

func run(room app.Action, command, origin string) (string, string, error) {
	return room.Action(command, origin)
}

func send(message string, to chan []byte) {
	to <- []byte(message)
}
