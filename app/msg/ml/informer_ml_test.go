package ml

import (
	"testing"

	"github.com/mcaci/ita-cards/card"
	"github.com/mcaci/ita-cards/set"
	"github.com/mcaci/msdb5/dom/auction"
	"github.com/mcaci/msdb5/dom/phase"
	"github.com/mcaci/msdb5/dom/player"
	"github.com/mcaci/msdb5/dom/team"
)

type mlInformerFake struct{}

func (mlInformerFake) AuctionScore() *auction.Score  { v := auction.Score(80); return &v }
func (mlInformerFake) Caller() *player.Player        { return player.New() }
func (mlInformerFake) Companion() *player.Player     { return player.New() }
func (mlInformerFake) CurrentPlayer() *player.Player { p := player.New(); p.RegisterAs("Me"); return p }
func (mlInformerFake) LastPlayer() *player.Player    { return player.New() }
func (mlInformerFake) Players() team.Players {
	p := player.New()
	p.Join("127.0.0.1")
	return team.Players{p}
}
func (mlInformerFake) PlayedCard() card.Item   { return *card.MustID(1) }
func (mlInformerFake) PlayedCards() *set.Cards { return &set.Cards{*card.MustID(2)} }
func (mlInformerFake) Phase() phase.ID         { return phase.PlayingCards }
func (mlInformerFake) Briscola() card.Item     { return *card.MustID(1) }

func (mlInformerFake) RoundError() error      { return nil }
func (mlInformerFake) IsSideUsed() bool       { return true }
func (mlInformerFake) IsSideToShow() bool     { return true }
func (mlInformerFake) SideDeck() *set.Cards   { return &set.Cards{*card.MustID(2)} }
func (mlInformerFake) SideSubset() set.Cards { return set.Cards{*card.MustID(2)} }

func TestMLMessage(t *testing.T) {
	testObj := mlInformerFake{}
	Write(testObj)
	t.Log(testObj)
}
