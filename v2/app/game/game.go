package game

import (
	"fmt"

	"github.com/mcaci/ita-cards/card"
	"github.com/mcaci/ita-cards/set"
	"github.com/mcaci/msdb5/v2/dom/auction"
	"github.com/mcaci/msdb5/v2/dom/player"
	"github.com/mcaci/msdb5/v2/dom/team"
)

// Game struct
type Game struct {
	players           team.Players
	c                 callers
	briscolaCard      card.Item
	side, playedCards set.Cards
	auctionScore      auction.Score
	opts              *Options
}

type Options struct {
	WithSide bool
}

func NewGame(gOpts *Options) *Game { return &Game{opts: gOpts} }

// New func
func New() *Game { return &Game{} }

func IsRoundOngoing(playedCards set.Cards) bool { return len(playedCards) < 5 }

func (g Game) String() string {
	return fmt.Sprintf("(Caller is: %s,\n Companion is: %s,\n Played cards: %v,\n Auction score: %d,\n Players: %v,\n Side Deck: %v)",
		g.c.caller.Name(), g.c.companion.Name()+" "+g.briscolaCard.String(), g.playedCards, g.auctionScore, g.players, g.side)
}

type callers struct {
	caller, companion *player.Player
}

func (c callers) Caller() *player.Player    { return c.caller }
func (c callers) Companion() *player.Player { return c.companion }
