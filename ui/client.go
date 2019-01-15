package main

import (
	"log"
	"strconv"
	"strings"

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
		err := c.socket.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			log.Println("Actual Write Error:", err)
		}
		info := strings.Split(string(msg), "#")
		switch info[0] {
		case "Join":
			c.room.msdb5board.Join(info[1], c.socket.RemoteAddr().String())
			c.room.forward <- []byte(c.room.msdb5board.String())
			c.room.forward <- []byte("Wait for other players")
		case "Auction":
			score, _ := strconv.Atoi(info[1])
			c.room.msdb5board.SetAuctionScore(uint8(score))
			c.room.forward <- []byte(c.room.msdb5board.String())
		case "Play":
			c.room.msdb5board.Nominate(info[1], info[2])
			c.room.forward <- []byte(c.room.msdb5board.String())
		}
	}
}
