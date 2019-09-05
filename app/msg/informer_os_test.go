package msg

import (
	"testing"

	"github.com/mcaci/ita-cards/card"
	"github.com/mcaci/ita-cards/set"
	"github.com/mcaci/msdb5/app/phase"
	"github.com/mcaci/msdb5/dom/auction"
	"github.com/mcaci/msdb5/dom/player"
	"github.com/mcaci/msdb5/dom/team"
)

type roundInformerFake struct{}

func (roundInformerFake) AuctionScore() *auction.Score  { v := auction.Score(80); return &v }
func (roundInformerFake) Caller() *player.Player        { return player.New() }
func (roundInformerFake) Companion() *player.Player     { return player.New() }
func (roundInformerFake) CurrentPlayer() *player.Player { return player.New() }
func (roundInformerFake) LastPlayer() *player.Player    { return player.New() }
func (roundInformerFake) Players() team.Players {
	p := player.New()
	p.Join("127.0.0.1")
	return team.Players{p}
}
func (roundInformerFake) PlayedCard() *card.Item   { return card.MustID(1) }
func (roundInformerFake) PlayedCards() *set.Cards { return &set.Cards{*card.MustID(2)} }
func (roundInformerFake) Phase() phase.ID         { return phase.PlayingCards }
func (roundInformerFake) Briscola() card.Item     { return *card.MustID(1) }

func (roundInformerFake) RoundError() error      { return nil }
func (roundInformerFake) IsSideUsed() bool       { return true }
func (roundInformerFake) IsSideToShow() bool     { return true }
func (roundInformerFake) SideDeck() *set.Cards   { return &set.Cards{*card.MustID(2)} }
func (roundInformerFake) SideSubset() *set.Cards { return &set.Cards{*card.MustID(2)} }

func TestOutputMessage(t *testing.T) {
	testObj := roundInformerFake{}
	toOS(testObj, "Play#1#Coin", "127.0.0.1")
	t.Log(testObj)
}
