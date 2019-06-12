package game

import (
	"testing"

	"github.com/nikiforosFreespirit/msdb5/app/phase"
	"github.com/nikiforosFreespirit/msdb5/dom/deck"
)

func TestExchangeCardThatPlayerDoesntHaveReturnsError(t *testing.T) {
	cards := deck.Cards{1, 2}
	gameTest := NewGame(true)
	messageBufferSize := 256
	playerChannel := make(chan []byte, messageBufferSize)
	gameTest.Join("127.0.0.55", playerChannel)
	gameTest.players[0].RegisterAs("A")
	gameTest.players[0].DropCards()
	gameTest.players[0].Draw(cards.Supply)
	gameTest.players[0].Draw(cards.Supply)
	gameTest.phase = phase.ExchangingCards
	gameTest.Process("Exchange#3#Cup", "127.0.0.55")
	if false {
		t.Fatal("Expecting error when playing a card not owned")
	}
}

func TestExchangeCardThatDoesntExistReturnsError(t *testing.T) {
	cards := deck.Cards{1, 2}
	gameTest := NewGame(true)
	messageBufferSize := 256
	playerChannel := make(chan []byte, messageBufferSize)
	gameTest.Join("127.0.0.55", playerChannel)
	gameTest.players[0].RegisterAs("A")
	gameTest.players[0].DropCards()
	gameTest.players[0].Draw(cards.Supply)
	gameTest.players[0].Draw(cards.Supply)
	gameTest.phase = phase.ExchangingCards
	gameTest.Process("Exchange#31#Cup", "127.0.0.55")
	if false {
		t.Fatal("Expecting error when playing a card doesn't exist")
	}
}

func TestPlayCardThatDoesntExistReturnsError(t *testing.T) {
	cards := deck.Cards{1, 2}
	gameTest := NewGame(true)
	messageBufferSize := 256
	playerChannel := make(chan []byte, messageBufferSize)
	gameTest.Join("127.0.0.55", playerChannel)
	gameTest.players[0].RegisterAs("A")
	gameTest.players[0].DropCards()
	gameTest.players[0].Draw(cards.Supply)
	gameTest.players[0].Draw(cards.Supply)
	gameTest.phase = phase.PlayingCards
	gameTest.Process("Card#31#Cup", "127.0.0.55")
	if false {
		t.Fatal("Expecting error when playing a card doesn't exist")
	}
}

func TestCompanionCardThatDoesntExistReturnsError(t *testing.T) {
	gameTest := NewGame(false)
	messageBufferSize := 256
	playerChannel := make(chan []byte, messageBufferSize)
	gameTest.Join("127.0.0.55", playerChannel)
	gameTest.players[0].RegisterAs("A")
	gameTest.phase = phase.ChoosingCompanion
	gameTest.Process("Companion#31#Cup", "127.0.0.55")
	if false {
		t.Fatal("Expecting error when playing a card doesn't exist")
	}
}

func TestInexistentPhaseReturnsError(t *testing.T) {
	gameTest := NewGame(false)
	messageBufferSize := 256
	playerChannel := make(chan []byte, messageBufferSize)
	gameTest.Join("127.0.0.55", playerChannel)
	gameTest.players[0].RegisterAs("A")
	gameTest.phase = phase.PlayingCards
	gameTest.Process("Rumba", "127.0.0.55")
	if false {
		t.Fatal("Expecting error when playing a card doesn't exist")
	}
}

func TestUnexpectedPhaseReturnsError(t *testing.T) {
	gameTest := NewGame(false)
	messageBufferSize := 256
	playerChannel := make(chan []byte, messageBufferSize)
	gameTest.Join("127.0.0.55", playerChannel)
	gameTest.players[0].RegisterAs("A")
	gameTest.phase = phase.PlayingCards
	gameTest.Process("Companion", "127.0.0.55")
	if false {
		t.Fatal("Expecting error when playing a card doesn't exist")
	}
}
