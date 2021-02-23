package game

import (
	"container/list"
	"fmt"

	"github.com/mcaci/ita-cards/card"
	"github.com/mcaci/ita-cards/set"
	"github.com/mcaci/msdb5/v2/dom/auction"
	"github.com/mcaci/msdb5/v2/dom/phase"
	"github.com/mcaci/msdb5/v2/dom/player"
	"github.com/mcaci/msdb5/v2/dom/team"
)

// Game struct
type Game struct {
	lastPlaying       list.List
	players           team.Players
	caller, companion *player.Player
	briscolaCard      card.Item
	side, playedCards set.Cards
	auctionScore      auction.Score
	phase             phase.ID
	opts              *Options
}

type Options struct {
	WithSide bool
}

func NewGame(gOpts *Options) *Game { return &Game{opts: gOpts} }

// New func
func New() *Game                              { return &Game{} }
func (g *Game) Phase() phase.ID               { return g.phase }
func (g *Game) Caller() *player.Player        { return g.caller }
func (g *Game) Companion() *player.Player     { return g.companion }
func (g *Game) CurrentPlayer() *player.Player { return g.lastPlaying.Front().Value.(*player.Player) }
func (g *Game) LastPlayer() *player.Player    { return g.lastPlaying.Back().Value.(*player.Player) }
func (g *Game) IsRoundOngoing() bool          { return len(g.playedCards) < 5 }
func (g *Game) CurrentPlayerIndex() uint8 {
	for i := range g.players {
		if g.players[i] != g.CurrentPlayer() {
			continue
		}
		return uint8(i)
	}
	return 0
}

func (g Game) String() string {
	return fmt.Sprintf("(Turn of: %s,\n Caller is: %s,\n Companion is: %s,\n Played cards: %v,\n Auction score: %d,\n Phase: %s,\n Players: %v,\n Side Deck: %v,\n Last Players: %v)",
		g.CurrentPlayer().Name(), g.caller.Name(), g.companion.Name()+" "+g.briscolaCard.String(), g.playedCards, g.auctionScore, g.phase, g.players, g.side, g.lastPlaying)
}
