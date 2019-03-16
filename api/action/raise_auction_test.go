package action

import (
	"testing"

	"github.com/nikiforosFreespirit/msdb5/player"
	"github.com/nikiforosFreespirit/msdb5/playerset"
)

func TestAuctionPhase(t *testing.T) {
	if testObject := NewAuction("", "", nil, nil, nil); testObject.Phase() != 1 {
		t.Fatalf("Unexpected phase")
	}
}

func TestAuctionFindsPlayerInTurn(t *testing.T) {
	testPlayer := player.New()
	testPlayer.Join("A", "127.0.0.3")
	if testObject := NewAuction("", "127.0.0.3", testPlayer, nil, nil); !testObject.Find(testPlayer) {
		t.Fatalf("Unexpected player")
	}
}

func TestAuctionDoesNotFindPlayerNotInTurn(t *testing.T) {
	testPlayer := player.New()
	if testObject := NewAuction("", "", nil, nil, nil); testObject.Find(testPlayer) {
		t.Fatalf("Unexpected player")
	}
}

func TestAuctionNextPlayerOf3is4(t *testing.T) {
	testPlayers := playerset.Players{player.New(), player.New(), player.New(), player.New(), player.New()}
	if testObject := NewAuction("", "", nil, testPlayers, nil); testObject.NextPlayer(3) != 4 {
		t.Fatalf("Next player should be 4")
	}
}

func TestAuctionNextPlayerOf1is3WithPlayer2Folded(t *testing.T) {
	testFoldedPlayer := player.New()
	testFoldedPlayer.Fold()
	testPlayers := playerset.Players{player.New(), player.New(), testFoldedPlayer, player.New(), player.New()}
	if testObject := NewAuction("", "", nil, testPlayers, nil); testObject.NextPlayer(1) != 3 {
	}
}

func TestAuctionNextPhaseIsFalse(t *testing.T) {
	testPlayers := playerset.Players{player.New(), player.New(), player.New(), player.New(), player.New()}
	if testObject := NewAuction("", "", nil, testPlayers, nil); testObject.NextPhase(testPlayers, testObject) {
		t.Fatalf("Should still be in auction phase")
	}
}

func TestAuctionNextPhaseIsTrue(t *testing.T) {
	testFoldedPlayer := player.New()
	testFoldedPlayer.Fold()
	testPlayers := playerset.Players{testFoldedPlayer, testFoldedPlayer, testFoldedPlayer, player.New(), testFoldedPlayer}
	if testObject := NewAuction("", "", nil, testPlayers, nil); !testObject.NextPhase(testPlayers, testObject) {
		t.Fatalf("Should still be in auction phase")
	}
}

func TestAuctionNextPhaseWithPlayersWithNonFoldedNameIsFalse(t *testing.T) {
	if testObject := NewAuction("", "", nil, nil, nil); testObject.NextPhasePlayerInfo(player.New()) {
		t.Fatalf("Should be false with non folded player")
	}
}

func TestAuctionNextPhaseWithFoldedPlayerIsTrue(t *testing.T) {
	testFoldedPlayer := player.New()
	testFoldedPlayer.Fold()
	if testObject := NewAuction("", "", nil, nil, nil); !testObject.NextPhasePlayerInfo(testFoldedPlayer) {
		t.Fatalf("Should be true with folded player")
	}
}
