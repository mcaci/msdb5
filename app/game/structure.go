package game

import (
	"fmt"
	"log"

	"github.com/nikiforosFreespirit/msdb5/app"
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
	companion    companion.Companion
	side         deck.Cards
	playedCards  deck.Cards
	auctionScore auction.Score
	phase        phase.ID
}

// NewAction func
func NewAction(withSide bool) app.Action {
	return NewGame(withSide)
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

// PlayerInTurn func
func (g *Game) PlayerInTurn() *player.Player { return g.players[g.playerInTurn] }

// Companion func
func (g *Game) Companion() *player.Player { return g.companion.Ref() }

// Players func
func (g *Game) Players() team.Players { return g.players }

// SetCompanion func
func (g *Game) SetCompanion(c card.ID, pl *player.Player) { g.companion = *companion.New(c, pl) }

// BriscolaSeed func
func (g *Game) BriscolaSeed() card.Seed { return g.companion.Card().Seed() }

// CurrentPhase func
func (g *Game) CurrentPhase() phase.ID { return g.phase }

// AuctionScore func
func (g *Game) AuctionScore() *auction.Score {
	return &g.auctionScore
}

// PlayedCards func
func (g *Game) PlayedCards() *deck.Cards {
	return &g.playedCards
}

// SideDeck func
func (g *Game) SideDeck() *deck.Cards {
	return &g.side
}

func (g Game) String() (str string) {
	return fmt.Sprintf("(Turn of: %s, Companion is: %s, Played cards: %+v, Auction score: %d, ID: %d)",
		g.PlayerInTurn().Name(), g.companion.Card(), g.PlayedCards(), g.AuctionScore(), g.phase)
}

// Log func
func (g Game) Log(request, origin string, err error) {
	_, playerLogged, err := g.Players().Find(func(p *player.Player) bool { return p.IsSameHost(origin) })
	if err == nil {
		log.Printf("New Action by %s\n", playerLogged.Name())
	}
	log.Printf("Action is %s\n", request)
	log.Printf("Any error raised: %+v\n", err)
	log.Printf("Game info after action: %+v\n", g)
}
