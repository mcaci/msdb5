package game

import (
	"testing"

	"github.com/nikiforosFreespirit/msdb5/app/phase"

	"github.com/nikiforosFreespirit/msdb5/dom/deck"
)

// type Game struct {
// 	lastPlaying  list.List
// 	players      team.Players
// 	caller       *player.Player
// 	companion    companion.Companion
// 	side         deck.Cards
// 	playedCards  deck.Cards
// 	auctionScore auction.Score
// 	phase        phase.ID
// }

func TestExchangeCardThatPlayerDoesntHaveReturnsError(t *testing.T) {
	cards := deck.Cards{1, 2}
	gameTest := NewGame(true)
	gameTest.players[0].Join("A", "127.0.0.55")
	gameTest.players[0].Hand().Clear()
	gameTest.players[0].Hand().Add(cards...)
	gameTest.phase = phase.ExchangingCards
	info := gameTest.Process("Exchange#3#Cup", "127.0.0.55")
	if info[0].Err() == nil {
		t.Fatal("Expecting error when playing a card not owned")
	}
}

func TestExchangeCardThatDoesntExistReturnsError(t *testing.T) {
	cards := deck.Cards{1, 2}
	gameTest := NewGame(true)
	gameTest.players[0].Join("A", "127.0.0.55")
	gameTest.players[0].Hand().Clear()
	gameTest.players[0].Hand().Add(cards...)
	gameTest.phase = phase.ExchangingCards
	info := gameTest.Process("Exchange#31#Cup", "127.0.0.55")
	if info[0].Err() == nil {
		t.Fatal("Expecting error when playing a card doesn't exist")
	}
}

func TestPlayCardThatDoesntExistReturnsError(t *testing.T) {
	cards := deck.Cards{1, 2}
	gameTest := NewGame(true)
	gameTest.players[0].Join("A", "127.0.0.55")
	gameTest.players[0].Hand().Clear()
	gameTest.players[0].Hand().Add(cards...)
	gameTest.phase = phase.PlayingCards
	info := gameTest.Process("Card#31#Cup", "127.0.0.55")
	if info[0].Err() == nil {
		t.Fatal("Expecting error when playing a card doesn't exist")
	}
}

func TestCompanionCardThatDoesntExistReturnsError(t *testing.T) {
	gameTest := NewGame(false)
	gameTest.players[0].Join("A", "127.0.0.55")
	gameTest.phase = phase.ChoosingCompanion
	info := gameTest.Process("Companion#31#Cup", "127.0.0.55")
	if info[0].Err() == nil {
		t.Fatal("Expecting error when playing a card doesn't exist")
	}
}

func TestInexistentPhaseReturnsError(t *testing.T) {
	gameTest := NewGame(false)
	gameTest.players[0].Join("A", "127.0.0.55")
	gameTest.phase = phase.PlayingCards
	info := gameTest.Process("Rumba", "127.0.0.55")
	if info[0].Err() == nil {
		t.Fatal("Expecting error when playing a card doesn't exist")
	}
}

func TestUnexpectedPhaseReturnsError(t *testing.T) {
	gameTest := NewGame(false)
	gameTest.players[0].Join("A", "127.0.0.55")
	gameTest.phase = phase.PlayingCards
	info := gameTest.Process("Companion", "127.0.0.55")
	if info[0].Err() == nil {
		t.Fatal("Expecting error when playing a card doesn't exist")
	}
}
