package nextplayer

import (
	"testing"

	"github.com/nikiforosFreespirit/msdb5/app/action"
	"github.com/nikiforosFreespirit/msdb5/app/phase"
	"github.com/nikiforosFreespirit/msdb5/dom/card"
	"github.com/nikiforosFreespirit/msdb5/dom/deck"
	"github.com/nikiforosFreespirit/msdb5/dom/player"
	"github.com/nikiforosFreespirit/msdb5/dom/team"
)

func testObject(phase phase.ID, players team.Players, playedCards *deck.Cards) action.NextPlayerSelector {
	return NewPlayerChanger(phase, players, playedCards, card.Coin)
}

func TestExchangeCardsNextPlayerOf0is1_Joining(t *testing.T) {
	testIndex := testObject(phase.Joining, nil, nil).NextPlayer(0)
	if testIndex != 1 {
		t.Fatalf("Next player should be 1")
	}
}

func TestExchangeCardsNextPlayerOf4is0_JoiningEndEdgeCase(t *testing.T) {
	testIndex := testObject(phase.Joining, nil, nil).NextPlayer(4)
	if testIndex != 0 {
		t.Fatalf("Next player should be 0")
	}
}

func TestExchangeCardsNextPlayerOf1is2_InsideAuction(t *testing.T) {
	testPlayers := team.Players{player.New(), player.New(), player.New(), player.New(), player.New()}
	testIndex := testObject(phase.InsideAuction, testPlayers, nil).NextPlayer(1)
	if testIndex != 2 {
		t.Fatalf("Next player should be 2")
	}
}

func TestExchangeCardsNextPlayerOf1is3_InsideAuctionWithFolded(t *testing.T) {
	testFoldedPlayer := player.New()
	testFoldedPlayer.Fold()
	testPlayers := team.Players{player.New(), player.New(), testFoldedPlayer, player.New(), player.New()}
	testIndex := testObject(phase.InsideAuction, testPlayers, nil).NextPlayer(1)
	if testIndex != 3 {
		t.Fatalf("Next player should be 3")
	}
}

func TestExchangeCardsNextPlayerOf2is2_ExchangingCards(t *testing.T) {
	testIndex := testObject(phase.ExchangingCards, nil, nil).NextPlayer(2)
	if testIndex != 2 {
		t.Fatalf("Next player should be 2")
	}
}

func TestExchangeCardsNextPlayerOf4is4_ChosingCompanion(t *testing.T) {
	testIndex := testObject(phase.ChosingCompanion, nil, nil).NextPlayer(4)
	if testIndex != 4 {
		t.Fatalf("Next player should be 4")
	}
}

func TestExchangeCardsNextPlayerOf3is4_PlayingCards(t *testing.T) {
	testPlayedCards := deck.Cards{1, 2, 3}
	testIndex := testObject(phase.PlayingCards, nil, &testPlayedCards).NextPlayer(3)
	if testIndex != 4 {
		t.Fatalf("Next player should be 4")
	}
}

func TestExchangeCardsNextPlayerOf0is4_PlayingCardsAndRoundEnds(t *testing.T) {
	testPlayedCards := deck.Cards{2, 3, 4, 1, 6}
	testIndex := testObject(phase.PlayingCards, nil, &testPlayedCards).NextPlayer(0)
	if testIndex != 4 {
		t.Fatalf("Next player should be 4")
	}
}

func TestExchangeCardsNextPlayerOf1is0_PlayingCardsAndRoundEnds(t *testing.T) {
	testPlayedCards := deck.Cards{2, 3, 4, 1, 6}
	testIndex := testObject(phase.PlayingCards, nil, &testPlayedCards).NextPlayer(1)
	if testIndex != 0 {
		t.Fatalf("Next player should be 0")
	}
}
