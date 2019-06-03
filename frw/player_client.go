package frw

import (
	"log"

	"github.com/gorilla/websocket"
)

// playerClient represents a playerClient connected to the game.
type playerClient struct {
	// socket is the web socket for this playerClient.
	socket *websocket.Conn
	// infoChannel is a channel on which messages are received from the game.
	infoChannel chan []byte
	// room is the room this playerClient is playing in.
	room *GameRoom
}

// Reads commands from UI
func (p *playerClient) read() {
	defer p.socket.Close()
	for {
		// read playerClient input
		_, msg, err := p.socket.ReadMessage()
		if err != nil {
			log.Println("Error from reading UI input:", err)
			return
		}
		// execute action
		command := playerCommand{string(msg), p.socket.RemoteAddr().String()}
		log.Println(command)
		p.sendToRoom(command)
	}
}

// Writes messages to UI
func (p *playerClient) write() {
	defer p.socket.Close()
	for msg := range p.infoChannel {
		err := p.socket.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			log.Println("Write to UI error:", err)
		}
	}
}

type playerCommand struct {
	request, origin string
}

func (p *playerClient) sendToRoom(command playerCommand) {
	p.room.commandChan <- command
}
