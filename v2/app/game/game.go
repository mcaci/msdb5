package game

import (
	"fmt"
	"log"

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
}

type Options struct {
	WithSide bool
}

func NewGame(gOpts *Options) *Game {
	g := &Game{opts: gOpts}
	g.players = *briscola5.NewPlayers()
	for i, p := range g.players.List() {
		p.RegisterAs(fmt.Sprintf("player%d", i+1))
	}
	return g
}

// New func
func New() *Game { return &Game{} }

func (g Game) String() string {
	return fmt.Sprintf("(Caller is: %s,\n Companion is: %s,\n Auction score: %d,\n Players: %v,\n Side Deck: %v)",
		g.players.Caller().Name(), g.players.Companion().Name()+" "+g.briscolaCard.String(), g.auctionScore, g.players, g.side)
}

func (g *Game) RegisterPlayer() func(string) {
	var index int
	return func(name string) {
		log.Printf("Registering player %d with name %q", index, name)
		g.players.At(index)
		index++
	}
}
