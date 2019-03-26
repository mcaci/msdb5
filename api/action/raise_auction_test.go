package action

import (
	"testing"

	"github.com/nikiforosFreespirit/msdb5/api/game"
	"github.com/nikiforosFreespirit/msdb5/board"

	"github.com/nikiforosFreespirit/msdb5/player"
	"github.com/nikiforosFreespirit/msdb5/playerset"
)

func TestAuctionNextPlayerOf3is4(t *testing.T) {
	testPlayers := playerset.Players{player.New(), player.New(), player.New(), player.New(), player.New()}
	if testObject := NewAuction("", "", testPlayers, nil); testObject.NextPlayer(3) != 4 {
		t.Fatalf("Next player should be 4")
	}
}

func TestAuctionNextPlayerOf1is3WithPlayer2Folded(t *testing.T) {
	testFoldedPlayer := player.New()
	testFoldedPlayer.Fold()
	testPlayers := playerset.Players{player.New(), player.New(), testFoldedPlayer, player.New(), player.New()}
	if testObject := NewAuction("", "", testPlayers, nil); testObject.NextPlayer(1) != 3 {
	}
}

func TestAuctionNextPhaseIsFalse(t *testing.T) {
	testPlayers := playerset.Players{player.New(), player.New(), player.New(), player.New(), player.New()}
	if testObject := NewAuction("", "", testPlayers, board.New()); game.ChosingCompanion == testObject.NextPhase(testPlayers, testObject) {
		t.Fatalf("Should still be in auction phase")
	}
}

func TestAuctionNextPhaseIsTrue(t *testing.T) {
	testFoldedPlayer := player.New()
	testFoldedPlayer.Fold()
	testPlayers := playerset.Players{testFoldedPlayer, testFoldedPlayer, testFoldedPlayer, player.New(), testFoldedPlayer}
	if testObject := NewAuction("", "", testPlayers, board.New()); game.ChosingCompanion != testObject.NextPhase(testPlayers, testObject) {
		t.Fatalf("Should still be in auction phase")
	}
}

func TestAuctionWithSideNextPhaseIsFalse(t *testing.T) {
	testPlayers := playerset.Players{player.New(), player.New(), player.New(), player.New(), player.New()}
	if testObject := NewAuction("", "", testPlayers, board.New()); game.ExchangingCards == testObject.NextPhase(testPlayers, testObject) {
		t.Fatalf("Should still be in auction phase")
	}
}

func TestAuctionWithSideNextPhaseIsTrue(t *testing.T) {
	testFoldedPlayer := player.New()
	testFoldedPlayer.Fold()
	testPlayers := playerset.Players{testFoldedPlayer, testFoldedPlayer, testFoldedPlayer, player.New(), testFoldedPlayer}
	testBoard := board.New()
	testBoard.SideDeck().Add(1)
	if testObject := NewAuction("", "", testPlayers, testBoard); game.ExchangingCards != testObject.NextPhase(testPlayers, testObject) {
		t.Fatalf("Should be in exchanging cards phase")
	}
}

func TestAuctionNextPhaseWithPlayersWithNonFoldedNameIsFalse(t *testing.T) {
	if testObject := NewAuction("", "", nil, nil); testObject.NextPhasePlayerInfo(player.New()) {
		t.Fatalf("Should be false with non folded player")
	}
}

func TestAuctionNextPhaseWithFoldedPlayerIsTrue(t *testing.T) {
	testFoldedPlayer := player.New()
	testFoldedPlayer.Fold()
	if testObject := NewAuction("", "", nil, nil); !testObject.NextPhasePlayerInfo(testFoldedPlayer) {
		t.Fatalf("Should be true with folded player")
	}
}
