package frw

import (
	"testing"

	"golang.org/x/text/language"
)

func TestSendCommand(t *testing.T) {
	testGameRoom := NewGameRoom(true, language.English)
	testPlayer := playerClient{nil, make(chan []byte), testGameRoom}
	testGameRoom.players[&testPlayer] = true
	testPlayerAddress := "localhost"
	testCommand := playerCommand{"Join#A", testPlayerAddress}
	go testPlayer.sendToRoom(testCommand)
	if _, ok := <-testGameRoom.commandChan; !ok {
		t.Fatal("command was not sent successfully")
	}
}
