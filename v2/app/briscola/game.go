package briscola

import (
	"fmt"

	"github.com/mcaci/ita-cards/set"
	"github.com/mcaci/msdb5/v2/app/misc"
	"github.com/mcaci/msdb5/v2/dom/briscola"
)

// Game struct
type Game struct {
	// Game name
	Name         string                `json:"name"`
	PlayerList   misc.Players          `json:"players"`
	briscolaCard briscola.Card         `json:"briscola"`
	BoardSet     *briscola.PlayedCards `json:"board"`
	registration func(string) error
	deck         *Deck
}

type Options struct {
	WithName string
}

var WithDefaultOptions = &Options{}

// NewGame func
func NewGame(gOpts *Options) *Game {
	p, rf := misc.NewWithRegistrator(2)
	g := Game{
		Name:         gOpts.WithName,
		PlayerList:   *p,
		deck:         NewDeck(),
		BoardSet:     briscola.NewPlayedCards(2),
		registration: rf,
	}
	return &g
}

func (g *Game) Players() *misc.Players       { return &g.PlayerList }
func (g *Game) Deck() *Deck                  { return g.deck }
func (g *Game) Board() *briscola.PlayedCards { return g.BoardSet }
func (g *Game) BoardCards() *set.Cards       { return g.BoardSet.Cards }
func (g *Game) Briscola() *briscola.Card     { return &g.briscolaCard }
func (g *Game) Created(name string) bool     { return name == g.Name }
func Register(name string, g *Game) error {
	err := g.registration(name)
	if err != nil {
		return err
	}
	if g.Players().None(func(p misc.Player) bool { return p.Name() == "" }) {
		Start(g)
	}
	return nil
}
func Set(card briscola.Card, g *Game) { g.briscolaCard = card }

func Start(g *Game) {
	Distribute(&struct {
		Players  misc.Players
		Deck     *Deck
		HandSize int
	}{
		Players:  g.PlayerList,
		Deck:     g.deck,
		HandSize: 3,
	})
	Set(briscola.Card{Item: g.deck.Top()}, g)
}

func (g Game) String() string {
	return fmt.Sprintf("(Players: %v, Board: %v, Briscola: %v, Deck: %v)", g.PlayerList, g.BoardSet, g.briscolaCard, g.deck)
}
