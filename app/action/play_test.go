package action

import (
	"testing"

	"github.com/mcaci/ita-cards/card"
	"github.com/mcaci/ita-cards/set"
	"github.com/mcaci/msdb5/app/phase"
	"github.com/mcaci/msdb5/dom/auction"
	"github.com/mcaci/msdb5/dom/player"
	"github.com/mcaci/msdb5/dom/team"
)

type fakeGameStructure struct {
	auctionScore  auction.Score
	currentPlayer *player.Player
	players       team.Players
	playedCards   *set.Cards
	phase         phase.ID
	sideDeck      *set.Cards
	companion     *player.Player
	briscolaCard  card.Item
}

func (gs fakeGameStructure) AuctionScore() *auction.Score     { return &gs.auctionScore }
func (gs fakeGameStructure) CurrentPlayer() *player.Player    { return gs.currentPlayer }
func (gs fakeGameStructure) Players() team.Players            { return gs.players }
func (gs fakeGameStructure) PlayedCards() *set.Cards          { return gs.playedCards }
func (gs fakeGameStructure) Phase() phase.ID                  { return gs.phase }
func (gs fakeGameStructure) SideDeck() *set.Cards             { return gs.sideDeck }
func (gs fakeGameStructure) SetAuction(score auction.Score)   { gs.auctionScore = score }
func (gs fakeGameStructure) SetBriscola(briscola *card.Item)  { gs.briscolaCard = *briscola }
func (gs fakeGameStructure) SetCompanion(comp *player.Player) { gs.companion = comp }

func TestExecJoin(t *testing.T) {
	gs := fakeGameStructure{
		auctionScore:  auction.Score(80),
		currentPlayer: player.New(),
		players:       team.Players{player.New()},
		playedCards:   &set.Cards{},
		phase:         phase.Joining,
		sideDeck:      &set.Cards{},
	}
	rq := fakeCardValueProv{card.MustID(11), "1"}
	err := Play(gs, rq)
	if err != nil {
		t.Fatal(err)
	}
}

func TestExecAuction(t *testing.T) {
	gs := fakeGameStructure{
		auctionScore:  auction.Score(80),
		currentPlayer: player.New(),
		players:       team.Players{player.New()},
		playedCards:   &set.Cards{},
		phase:         phase.InsideAuction,
		sideDeck:      &set.Cards{},
	}
	rq := fakeCardValueProv{card.MustID(11), "81"}
	err := Play(gs, rq)
	if err != nil {
		t.Fatal(err)
	}
}

func TestExecAuctionFold(t *testing.T) {
	gs := fakeGameStructure{
		auctionScore:  auction.Score(80),
		currentPlayer: player.New(),
		players:       team.Players{player.New()},
		playedCards:   &set.Cards{},
		phase:         phase.InsideAuction,
		sideDeck:      &set.Cards{},
	}
	rq := fakeCardValueProv{card.MustID(11), "79"}
	err := Play(gs, rq)
	if err != nil {
		t.Fatal(err)
	}
}

func TestExecExchange(t *testing.T) {
	p := player.New()
	p.Hand().Add(*card.MustID(11))
	gs := fakeGameStructure{
		auctionScore:  auction.Score(80),
		currentPlayer: p,
		players:       team.Players{p},
		playedCards:   &set.Cards{},
		phase:         phase.ExchangingCards,
		sideDeck:      &set.Cards{*card.MustID(1)},
	}
	rq := fakeCardValueProv{card.MustID(11), "1"}
	err := Play(gs, rq)
	if err != nil {
		t.Fatal(err)
	}
}

func TestExecEndExchange(t *testing.T) {
	p := player.New()
	p.Hand().Add(*card.MustID(11))
	gs := fakeGameStructure{
		auctionScore:  auction.Score(80),
		currentPlayer: p,
		players:       team.Players{p},
		playedCards:   &set.Cards{},
		phase:         phase.ExchangingCards,
		sideDeck:      &set.Cards{*card.MustID(1)},
	}
	rq := fakeCardValueProv{card.MustID(11), "0"}
	err := Play(gs, rq)
	if err != nil {
		t.Fatal(err)
	}
}

func TestExecCompanion(t *testing.T) {
	p := player.New()
	p.Hand().Add(*card.MustID(11))
	gs := fakeGameStructure{
		auctionScore:  auction.Score(80),
		currentPlayer: p,
		players:       team.Players{p},
		playedCards:   &set.Cards{},
		phase:         phase.ChoosingCompanion,
		sideDeck:      &set.Cards{},
	}
	rq := fakeCardValueProv{card.MustID(11), "1"}
	err := Play(gs, rq)
	if err != nil {
		t.Fatal(err)
	}
}

func TestExecPlayCard(t *testing.T) {
	p := player.New()
	p.Hand().Add(*card.MustID(11))
	gs := fakeGameStructure{
		auctionScore:  auction.Score(80),
		currentPlayer: p,
		players:       team.Players{p},
		playedCards:   &set.Cards{},
		phase:         phase.PlayingCards,
		sideDeck:      &set.Cards{},
	}
	rq := fakeCardValueProv{card.MustID(11), "1"}
	err := Play(gs, rq)
	if err != nil {
		t.Fatal(err)
	}
}
