package action

import (
	"testing"

	"github.com/nikiforosFreespirit/msdb5/api/game"

	"github.com/nikiforosFreespirit/msdb5/player"
	"github.com/nikiforosFreespirit/msdb5/playerset"
)

func TestAuctionWithSidePhase(t *testing.T) {
	if testObject := NewAuctionWithSide("", "", nil, nil, nil); testObject.Phase() != game.InsideAuction {
		t.Fatalf("Unexpected phase")
	}
}

func TestAuctionWithSideFindsPlayerInTurn(t *testing.T) {
	testPlayer := player.New()
	testPlayer.Join("A", "127.0.0.3")
	if testObject := NewAuctionWithSide("", "127.0.0.3", testPlayer, nil, nil); !testObject.Find(testPlayer) {
		t.Fatalf("Unexpected player")
	}
}

func TestAuctionWithSideDoesNotFindPlayerNotInTurn(t *testing.T) {
	testPlayer := player.New()
	if testObject := NewAuctionWithSide("", "", nil, nil, nil); testObject.Find(testPlayer) {
		t.Fatalf("Unexpected player")
	}
}

func TestAuctionWithSideNextPlayerOf3is4(t *testing.T) {
	testPlayers := playerset.Players{player.New(), player.New(), player.New(), player.New(), player.New()}
	if testObject := NewAuctionWithSide("", "", nil, testPlayers, nil); testObject.NextPlayer(3) != 4 {
		t.Fatalf("Next player should be 4")
	}
}

func TestAuctionWithSideNextPlayerOf1is3WithPlayer2Folded(t *testing.T) {
	testFoldedPlayer := player.New()
	testFoldedPlayer.Fold()
	testPlayers := playerset.Players{player.New(), player.New(), testFoldedPlayer, player.New(), player.New()}
	if testObject := NewAuctionWithSide("", "", nil, testPlayers, nil); testObject.NextPlayer(1) != 3 {
	}
}

func TestAuctionWithSideNextPhaseIsFalse(t *testing.T) {
	testPlayers := playerset.Players{player.New(), player.New(), player.New(), player.New(), player.New()}
	if testObject := NewAuctionWithSide("", "", nil, testPlayers, nil); game.ExchangingCards == testObject.NextPhase(testPlayers, testObject) {
		t.Fatalf("Should still be in auction phase")
	}
}

func TestAuctionWithSideNextPhaseIsTrue(t *testing.T) {
	testFoldedPlayer := player.New()
	testFoldedPlayer.Fold()
	testPlayers := playerset.Players{testFoldedPlayer, testFoldedPlayer, testFoldedPlayer, player.New(), testFoldedPlayer}
	if testObject := NewAuctionWithSide("", "", nil, testPlayers, nil); game.ExchangingCards != testObject.NextPhase(testPlayers, testObject) {
		t.Fatalf("Should still be in auction phase")
	}
}

func TestAuctionWithSideNextPhaseWithPlayersWithNonFoldedNameIsFalse(t *testing.T) {
	if testObject := NewAuctionWithSide("", "", nil, nil, nil); testObject.NextPhasePlayerInfo(player.New()) {
		t.Fatalf("Should be false with non folded player")
	}
}

func TestAuctionWithSideNextPhaseWithFoldedPlayerIsTrue(t *testing.T) {
	testFoldedPlayer := player.New()
	testFoldedPlayer.Fold()
	if testObject := NewAuctionWithSide("", "", nil, nil, nil); !testObject.NextPhasePlayerInfo(testFoldedPlayer) {
		t.Fatalf("Should be true with folded player")
	}
}
