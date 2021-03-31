package game

import (
	"fmt"

	"github.com/mcaci/msdb5/v2/dom/briscola"
	"github.com/mcaci/msdb5/v2/dom/briscola5"
	"github.com/mcaci/msdb5/v2/dom/briscola5/auction"
)

// Game struct
type Game struct {
	opts         *Options
	players      briscola5.Players
	briscolaCard briscola.Card
	side         briscola5.Side
	auctionScore auction.Score
}

type Options struct {
	WithSide bool
}

func NewGame(gOpts *Options) *Game { return &Game{opts: gOpts} }

// New func
func New() *Game { return &Game{} }

func (g Game) String() string {
	return fmt.Sprintf("(Caller is: %s,\n Companion is: %s,\n Auction score: %d,\n Players: %v,\n Side Deck: %v)",
		g.players.Caller().Name(), g.players.Companion().Name()+" "+g.briscolaCard.String(), g.auctionScore, g.players, g.side)
}
