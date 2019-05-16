package game

import (
	"fmt"
	"log"

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
	board        Board
	phase        Phase
}

// NewGame func
func NewGame(withSide bool) *Game {
	g := new(Game)
	makePlayers(g)
	distributeCards(&g.players, g.board.SideDeck(), withSide)
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

// Scorers func
func (g *Game) Scorers() []player.Scorer {
	scorers := make([]player.Scorer, 0)
	for _, p := range g.players {
		scorers = append(scorers, p)
	}
	return scorers
}

// SetCompanion func
func (g *Game) SetCompanion(c card.ID, pl *player.Player) { g.companion = *companion.New(c, pl) }

// Board func
func (g *Game) Board() *Board { return &g.board }

// BriscolaSeed func
func (g *Game) BriscolaSeed() card.Seed { return g.companion.Card().Seed() }

// CurrentPhase func
func (g *Game) CurrentPhase() Phase { return g.phase }

// NextPhase func
func (g *Game) NextPhase(phase Phase) { g.phase = phase }

// NextPlayer func
func (g *Game) NextPlayer(generateIndex func(uint8) uint8) {
	g.playerInTurn = generateIndex(g.playerInTurn)
}

func (g Game) String() (str string) {
	return fmt.Sprintf("(Turn of: %s, Companion is: %s, Board Info: %+v, Phase: %d)",
		g.PlayerInTurn().Name(), g.companion.Card(), g.board, g.phase)
}

func (g Game) Log(request, origin string, err error) {
	_, playerLogged, err := g.Players().Find(func(p *player.Player) bool { return p.IsSameHost(origin) })
	if err == nil {
		log.Printf("New Action by %s\n", playerLogged.Name())
	}
	log.Printf("Action is %s\n", request)
	log.Printf("Any error raised: %+v\n", err)
	log.Printf("Game info after action: %+v\n", g)
}
