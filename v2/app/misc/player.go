package misc

import (
	"fmt"

	"github.com/mcaci/ita-cards/set"
	"github.com/mcaci/msdb5/v2/dom/briscola"
	"github.com/mcaci/msdb5/v2/dom/briscola5"
)

// Options struct
type Options struct {
	Name  string
	For2P bool
	For5P bool
}

// Player interface
type Player interface {
	Name() string
	Hand() *set.Cards
	Pile() *set.Cards
	fmt.Stringer
}

// New func
func New(o *Options) Player {
	var p Player
	switch {
	case o.For2P:
		p = briscola.NewB2Player(o.Name)
	case o.For5P:
		p = briscola5.NewB5Player(o.Name)
	default:
		p = briscola.NewB2Player(o.Name)
	}
	return p
}
