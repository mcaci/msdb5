package gamelog

import (
	"testing"

	"github.com/nikiforosFreespirit/msdb5/app/phase"
	"github.com/nikiforosFreespirit/msdb5/dom/auction"
	"github.com/nikiforosFreespirit/msdb5/dom/card"
	"github.com/nikiforosFreespirit/msdb5/dom/deck"
	"github.com/nikiforosFreespirit/msdb5/dom/player"
)

type fakeMiner struct {
	p phase.ID
}

func (m fakeMiner) AuctionScore() *auction.Score  { a := auction.Score(80); return &a }
func (m fakeMiner) Companion() *player.Player     { return player.New() }
func (m fakeMiner) CurrentPlayer() *player.Player { return player.New() }
func (m fakeMiner) LastCardPlayed() card.ID       { return 1 }
func (m fakeMiner) Phase() phase.ID               { return m.p }
func (m fakeMiner) IsSideUsed() bool              { return true }
func (m fakeMiner) SideDeck() *deck.Cards         { a := deck.Cards{1, 2, 3, 4, 5}; return &a }

func TestMinerMsgEmpty(t *testing.T) {
	s := new(fakeWriter)
	ToFile(fakeMiner{1}, s)
	if len(*s) != 0 {
		t.Fatalf("Expecting %s but got %s", "", *s)
	}
}

func TestMinerMsg(t *testing.T) {
	s := new(fakeWriter)
	ToFile(fakeMiner{3}, s)
	if len(*s) == 0 {
		t.Fatalf("Expecting %s but got %s", "", *s)
	}
}
