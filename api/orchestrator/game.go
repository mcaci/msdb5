package orchestrator

import (
	"errors"
	"strconv"

	"github.com/nikiforosFreespirit/msdb5/api"
	"github.com/nikiforosFreespirit/msdb5/board"
	"github.com/nikiforosFreespirit/msdb5/card"
	"github.com/nikiforosFreespirit/msdb5/companion"
	"github.com/nikiforosFreespirit/msdb5/deck"
	"github.com/nikiforosFreespirit/msdb5/display"
	"github.com/nikiforosFreespirit/msdb5/player"
	"github.com/nikiforosFreespirit/msdb5/playerset"
)

// Game struct
type Game struct {
	playerInTurn uint8
	players      playerset.Players
	companion    companion.Companion
	info         board.Board
	phase        phase
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

// NewAction func
func NewAction() api.Action {
	return NewGame()
}

func (g *Game) setCompanion(c card.ID, p *player.Player) {
	g.companion = *companion.New(c, p)
}

func (g *Game) nextPhase() {
	g.phase++
}

func (g *Game) nextPlayer(generateIndex func() uint8) {
	g.playerInTurn = generateIndex()
}

func (g *Game) phaseCheck(current phase) (err error) {
	if g.phase != current {
		err = errors.New("Phase is not " + strconv.Itoa(int(current)))
	}
	return
}

// Info func
func (g Game) Info() []display.Info {
	gameInfo := g.info.Info()
	gameInfo = append(gameInfo, g.players[g.playerInTurn].Name())
	compCard := display.NewInfo("Companion", ":", g.companion.Card().String(), ";")
	gameInfo = append(gameInfo, compCard)
	return display.Wrap("Game", gameInfo...)
}

func (g Game) String() string {
	gameInfo := g.Info()
	players := display.NewInfo("Players", ":", g.players.String(), ";")
	gameInfo = append(gameInfo, players)
	if g.companion.Ref() != nil {
		gameInfo = append(gameInfo, g.companion.Ref().Name())
	}
	gameInfo = append(gameInfo, g.info.Info()...)
	phase := display.NewInfo("PlayerInTurn", ":", strconv.Itoa(int(g.phase)), ";")
	gameInfo = append(gameInfo, phase)
	return display.All(display.Wrap("Game", gameInfo...)...)
}
