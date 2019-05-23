package game

import (
	"fmt"

	"github.com/nikiforosFreespirit/msdb5/app/phase"
	"github.com/nikiforosFreespirit/msdb5/dom/auction"
	"github.com/nikiforosFreespirit/msdb5/dom/card"
	"github.com/nikiforosFreespirit/msdb5/dom/companion"
	"github.com/nikiforosFreespirit/msdb5/dom/deck"
	"github.com/nikiforosFreespirit/msdb5/dom/player"
	"github.com/nikiforosFreespirit/msdb5/dom/team"
)

// Game struct
type Game struct {
	playerInTurn uint8
	players      team.Players
	caller       *player.Player
	companion    companion.Companion
	side         deck.Cards
	playedCards  deck.Cards
	auctionScore auction.Score
	phase        phase.ID
}

// NewGame func
func NewGame(withSide bool) *Game {
	g := new(Game)
	makePlayers(g)
	distributeCards(&g.players, &g.side, withSide)
	return g
}

func makePlayers(g *Game) {
	for i := 0; i < 5; i++ {
		g.players.Add(*player.New())
	}
}

func distributeCards(players *team.Players, side *deck.Cards, withSide bool) {
	d := deck.Deck()
	for i := 0; i < deck.DeckSize; i++ {
		if withSide && i >= deck.DeckSize-5 {
			side.Add(d.Supply())
		} else {
			(*players)[i%5].Draw(d)
		}
	}
}

func (g *Game) AuctionScore() auction.Score   { return g.auctionScore }
func (g *Game) Companion() *player.Player     { return g.companion.Ref() }
func (g *Game) CurrentPlayer() *player.Player { return g.players[g.playerInTurn] }
func (g *Game) LastCardPlayed() card.ID       { return g.playedCards[len(g.playedCards)-1] }
func (g *Game) Phase() phase.ID               { return g.phase }
func (g *Game) SideDeck() deck.Cards          { return g.side }
func (g *Game) IsSideUsed() bool              { return len(g.side) > 0 }

func (g *Game) playersRef() team.Players { return g.players }
func (g *Game) briscola() card.Seed      { return g.companion.Card().Seed() }
func (g *Game) cardsOnTheBoard() int     { return len(g.playedCards) }

func (g *Game) setCompanion(c card.ID, pl *player.Player) { g.companion = *companion.New(c, pl) }

func (g Game) String() (str string) {
	return fmt.Sprintf("(Turn of: %s, Companion is: %s, Played cards: %+v, Auction score: %d, ID: %d)",
		g.CurrentPlayer().Name(), g.companion.Card(), g.playedCards, g.auctionScore, g.phase)
}
