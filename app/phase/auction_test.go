package phase

import (
	"testing"

	"github.com/mcaci/msdb5/dom/auction"
)

type auctiontest string

func (a auctiontest) Value() string { return string(a) }

type auctionertest struct {
	folded bool
	score  uint8
}

func (a auctionertest) Folded() bool                 { return a.folded }
func (a auctionertest) AuctionScore() *auction.Score { s := auction.Score(a.score); return &s }

func TestProcessAuctionWithErr(t *testing.T) {
	data := Auction(auctiontest("ciao"), auctionertest{})
	if data.ToFold() != true {
		t.Fatal("Unexpected folded value")
	}
}

func TestProcessAuctionNoFoldFromScore(t *testing.T) {
	data := Auction(auctiontest("80"), auctionertest{score: 65})
	if data.ToFold() != false {
		t.Fatal("Unexpected folded value")
	}
}

func TestProcessAuctionWithFold(t *testing.T) {
	data := Auction(auctiontest("80"), auctionertest{folded: true})
	if data.ToFold() != true {
		t.Fatal("Unexpected folded value")
	}
}
func TestProcessAuctionWithFoldFromScore(t *testing.T) {
	data := Auction(auctiontest("80"), auctionertest{score: 90})
	if data.ToFold() != true {
		t.Fatal("Unexpected folded value")
	}
}

func TestScoreAuctionWithErr(t *testing.T) {
	data := Auction(auctiontest("ciao"), auctionertest{})
	if data.Score() != 61 {
		t.Fatalf("Unexpected value: %d", data.Score())
	}
}

func TestScoreAuctionNoFoldFromScore(t *testing.T) {
	data := Auction(auctiontest("80"), auctionertest{score: 65})
	if data.Score() != 80 {
		t.Fatalf("Unexpected value: %d", data.Score())
	}
}

func TestScoreAuctionWithFold(t *testing.T) {
	data := Auction(auctiontest("70"), auctionertest{folded: true})
	if data.Score() != 70 {
		t.Fatalf("Unexpected value: %d", data.Score())
	}
}
func TestScoreAuctionWithFoldFromScore(t *testing.T) {
	data := Auction(auctiontest("85"), auctionertest{score: 90})
	if data.Score() != 90 {
		t.Fatalf("Unexpected value: %d", data.Score())
	}
}

func TestSideCardsInAuction(t *testing.T) {
	data := Auction(auctiontest("100"), auctionertest{score: 90})
	if data.SideCards() != 2 {
		t.Fatalf("Unexpected value: %d", data.SideCards())
	}
}
