package game

import (
	"fmt"
	"log"

	"github.com/nikiforosFreespirit/msdb5/board"
	"github.com/nikiforosFreespirit/msdb5/card"
	"github.com/nikiforosFreespirit/msdb5/companion"
	"github.com/nikiforosFreespirit/msdb5/deck"
	"github.com/nikiforosFreespirit/msdb5/player"
	"github.com/nikiforosFreespirit/msdb5/playerset"
)

// Game struct
type Game struct {
	playerInTurn uint8
	players      playerset.Players
	companion    companion.Companion
	board        board.Board
	phase        Phase
}

// NewGame func
func NewGame() *Game {
	g := new(Game)
	makePlayers(g)
	playersDrawAllCards(&g.players)
	return g
}

func makePlayers(g *Game) {
	for i := 0; i < 5; i++ {
		g.players.Add(*player.New())
	}
}

func playersDrawAllCards(players *playerset.Players) {
	d := deck.Deck()
	for i := 0; i < deck.DeckSize; i++ {
		(*players)[i%5].Draw(d)
	}
}

// PlayerInTurn func
func (g *Game) PlayerInTurn() *player.Player { return g.players[g.playerInTurn] }

// Companion func
func (g *Game) Companion() *player.Player { return g.companion.Ref() }

// Players func
func (g *Game) Players() playerset.Players { return g.players }

// SetCompanion func
func (g *Game) SetCompanion(c card.ID, pl *player.Player) { g.companion = *companion.New(c, pl) }

// Board func
func (g *Game) Board() *board.Board { return &g.board }

// BriscolaSeed func
func (g *Game) BriscolaSeed() card.Seed { return g.companion.Card().Seed() }

// CurrentPhase func
func (g *Game) CurrentPhase() Phase { return g.phase }

// NextPhase func
func (g *Game) NextPhase() { g.phase++ }

// NextPlayer func
func (g *Game) NextPlayer(generateIndex func(uint8) uint8) {
	g.playerInTurn = generateIndex(g.playerInTurn)
}

func (g Game) String() (str string) {
	return fmt.Sprintf("(Turn of: %s, Companion is: %s, Board Info: %+v, Phase: %d)",
		g.PlayerInTurn().Name(), g.companion.Card(), g.board, g.phase)
}

func (g Game) Log(request, origin string, err error) {
	playerLogged, err := g.Players().Find(func(p *player.Player) bool { return p.IsSameHost(origin) })
	if err == nil {
		log.Printf("New Action by %s\n", playerLogged.Name())
	}
	log.Printf("Action is %s\n", request)
	log.Printf("Any error raised: %+v\n", err)
	log.Printf("Game info after action: %+v\n", g)
}
