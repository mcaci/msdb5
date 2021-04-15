package game

import (
	"fmt"

	"github.com/mcaci/msdb5/v2/dom/briscola"
	"github.com/mcaci/msdb5/v2/dom/briscola5"
)

// Game struct
type Game struct {
	opts         *Options
	players      briscola5.Players
	briscolaCard briscola.Card
	side         briscola5.Side
	auctionScore briscola5.AuctionScore
	registration func(string) error
}

type Options struct {
	WithSide bool
	WithName string
}

func NewGame(gOpts *Options) *Game {
	g := &Game{opts: gOpts}
	g.players = *briscola5.NewPlayers()
	g.registration = g.players.Registration()
	return g
}

// New func
func New() *Game { return &Game{} }

func (g *Game) Players() *briscola5.Players { return &g.players }
func (g *Game) Started(name string) bool    { return name == g.opts.WithName }
func Register(name string, g *Game) error   { return g.registration(name) }

func (g Game) String() string {
	return fmt.Sprintf("(Caller is: %s,\n Companion is: %s,\n Auction score: %d,\n Players: %v,\n Side Deck: %v)",
		g.players.Caller().Name(), g.players.Companion().Name()+" "+g.briscolaCard.String(), g.auctionScore, g.players, g.side)
}
