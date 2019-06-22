package play

import (
	"container/list"
	"testing"

	"golang.org/x/text/language"

	"github.com/nikiforosFreespirit/msdb5/app/phase"
	"github.com/nikiforosFreespirit/msdb5/dom/auction"
	"github.com/nikiforosFreespirit/msdb5/dom/card"
	"github.com/nikiforosFreespirit/msdb5/dom/deck"
	"github.com/nikiforosFreespirit/msdb5/dom/player"
	"github.com/nikiforosFreespirit/msdb5/dom/team"
)

var messageSink = func(p *player.Player, msg string) {}

type fakeGame struct {
	current *player.Player
	phase   phase.ID
	players team.Players
}

func newTestGame(ph phase.ID) fakeGame {
	f := fakeGame{}
	t := team.Players{}
	p := player.New()
	p.Join("127.0.0.51")
	p.RegisterAs("A")
	t.Add(p)
	f.players = t
	f.current = p
	f.phase = ph
	return f
}

func (g fakeGame) CurrentPlayer() *player.Player { return g.current }
func (g fakeGame) Players() team.Players         { return g.players }
func (g fakeGame) LastPlaying() *list.List       { return list.New() }
func (g fakeGame) Lang() language.Tag            { return language.English }
func (g fakeGame) Phase() phase.ID               { return g.phase }
func (g fakeGame) Briscola() card.Seed           { return card.Coin }
func (g fakeGame) IsSideUsed() bool              { return true }
func (g fakeGame) PlayedCards() *deck.Cards      { return new(deck.Cards) }
func (g fakeGame) SideDeck() *deck.Cards         { return new(deck.Cards) }
func (g fakeGame) AuctionScore() *auction.Score  { score := auction.Score(80); return &score }

type fakeRq struct {
	request, value string
	c              card.ID
}

func (r fakeRq) Value() string          { return r.value }
func (r fakeRq) Action() string         { return r.request }
func (r fakeRq) Card() (card.ID, error) { return r.c, nil }
func (r fakeRq) EndExchange() bool      { return false }

func TestProcessRequestWithNoErr(t *testing.T) {
	gameTest := newTestGame(0)
	rq := fakeRq{"Join", "A", 1}
	setCompanion := func(p *player.Player) {}
	setBriscolaCard := func(c card.ID) {}
	err := Request(gameTest, rq, setCompanion, setBriscolaCard, messageSink)
	if err != nil {
		t.Fatal(err)
	}
}

func TestProcessRequestWithErr(t *testing.T) {
	gameTest := newTestGame(4)
	rq := fakeRq{"Card", "A", 50}
	setCompanion := func(p *player.Player) {}
	setBriscolaCard := func(c card.ID) {}
	err := Request(gameTest, rq, setCompanion, setBriscolaCard, messageSink)
	if err == nil {
		t.Fatal("Error was expected")
	}
}

func TestProcessAuctionRequestWithNoErr(t *testing.T) {
	gameTest := newTestGame(3)
	rq := fakeRq{"Auction", "75", 1}
	setCompanion := func(p *player.Player) {}
	setBriscolaCard := func(c card.ID) {}
	err := Request(gameTest, rq, setCompanion, setBriscolaCard, messageSink)
	if err != nil {
		t.Fatal(err)
	}
}
