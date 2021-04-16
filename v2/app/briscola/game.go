package briscola

import (
	"fmt"

	"github.com/mcaci/msdb5/v2/dom/briscola"
)

// Game struct
type Game struct {
	opts         *Options
	players      briscola.Players
	briscolaCard briscola.Card
	registration func(string) error
}

type Options struct {
	WithSide bool
	WithName string
}

func NewGame(gOpts *Options) *Game {
	g := &Game{opts: gOpts}
	g.players = *briscola.NewPlayers()
	g.registration = g.players.Registration()
	return g
}

// New func
func New() *Game { return &Game{} }

func (g *Game) Players() *briscola.Players { return &g.players }
func (g *Game) Started(name string) bool   { return name == g.opts.WithName }
func Register(name string, g *Game) error  { return g.registration(name) }

func (g Game) String() string {
	return fmt.Sprintf("(Players: %v,\n)", g.players)
}
