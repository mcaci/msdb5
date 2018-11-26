package board

import (
	"testing"
)

func TestAuctionSentByChannel(t *testing.T) {
	b, cToPlayer, cFromPlayer := setUp()

	b.SetAuctionScore(playerRound(b.AuctionScore(), cToPlayer, cFromPlayer))

	verification(t, 61, b.AuctionScore())
}

func TestAuctionSentByChannelTwice(t *testing.T) {
	b, cToPlayer, cFromPlayer := setUp()

	b.SetAuctionScore(playerRound(b.AuctionScore(), cToPlayer, cFromPlayer))
	b.SetAuctionScore(playerRound(b.AuctionScore(), cToPlayer, cFromPlayer))

	verification(t, 62, b.AuctionScore())
}

func TestAuctionSentByChannelThrice(t *testing.T) {
	b, cToPlayer, cFromPlayer := setUp()

	b.SetAuctionScore(playerRound(b.AuctionScore(), cToPlayer, cFromPlayer))
	b.SetAuctionScore(playerRound(b.AuctionScore(), cToPlayer, cFromPlayer))
	b.SetAuctionScore(playerRound(b.AuctionScore(), cToPlayer, cFromPlayer))

	verification(t, 63, b.AuctionScore())
}

func setUp() (*Board, chan uint8, chan uint8) {
	b := New()
	b.SetAuctionScore(60)
	cToPlayer := make(chan uint8, 1)
	cFromPlayer := make(chan uint8, 1)
	return b, cToPlayer, cFromPlayer
}

func playerRound(previousScore uint8, cToPlayer chan uint8, cFromPlayer chan uint8) uint8 {
	// send score to player
	cToPlayer <- previousScore
	// player sends his score
	go playerTellsScore(cToPlayer, cFromPlayer)
	// board receives player's score
	return <-cFromPlayer
}

func playerTellsScore(cToPlayer <-chan uint8, cFromPlayer chan<- uint8) {
	actualScore := <-cToPlayer
	cFromPlayer <- actualScore + 1
}

func verification(t *testing.T, expected, actual uint8) {
	// verify score increase
	if actual != expected {
		t.Fatalf("score is not the %d but %d", expected, actual)
	}
}
