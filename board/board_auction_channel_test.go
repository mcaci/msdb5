package board

import (
	"testing"
)

func TestAuctionSentByChannel(t *testing.T) {
	// set up
	b := New()
	b.SetAuctionScore(60)
	cToPlayer := make(chan uint8, 1)
	cFromPlayer := make(chan uint8, 1)

	// send score to player
	cToPlayer <- b.AuctionScore()

	// player sends his score
	go playerTellsScore(cToPlayer, cFromPlayer)

	// board receives player's score
	b.SetAuctionScore(<-cFromPlayer)

	// verify score increase
	if b.AuctionScore() != 61 {
		t.Fatal("score is not the 61")
	}
}

func TestAuctionSentByChannelTwice(t *testing.T) {
	// set up
	b := New()
	b.SetAuctionScore(60)
	cToPlayer := make(chan uint8, 1)
	cFromPlayer := make(chan uint8, 1)

	b.SetAuctionScore(round(b.AuctionScore(), cToPlayer, cFromPlayer))
	b.SetAuctionScore(round(b.AuctionScore(), cToPlayer, cFromPlayer))

	// verify score increase
	if b.AuctionScore() != 62 {
		t.Fatalf("score is not the 62 but %d", b.AuctionScore())
	}
}
func round(previousScore uint8, cToPlayer chan uint8, cFromPlayer chan uint8) uint8 {
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
