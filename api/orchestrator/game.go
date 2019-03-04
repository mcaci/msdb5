package orchestrator

import (
	"github.com/nikiforosFreespirit/msdb5/api"
	"github.com/nikiforosFreespirit/msdb5/board"
	"github.com/nikiforosFreespirit/msdb5/companion"
	"github.com/nikiforosFreespirit/msdb5/deck"
	"github.com/nikiforosFreespirit/msdb5/player"
	"github.com/nikiforosFreespirit/msdb5/playerset"
)

// Game struct
type Game struct {
	players    playerset.Players
	companion  companion.Companion
	info       board.Board
	statusInfo status
}

type status uint8

const (
	joining status = iota
	scoreAuction
)

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

// Players func
func (g *Game) Players() playerset.Players {
	return g.players
}
