package briscola

import (
	"errors"
	"fmt"
)

// Game struct
type Game struct {
	// Game name
	Name         string       `json:"name"`
	PlayerList   *Players     `json:"players"`
	Briscola     Card         `json:"briscola"`
	BoardSet     *PlayedCards `json:"board"`
	registration func(string) error
	deck         *Deck
	next         int
}

type Options struct {
	WithName string
}

var WithDefaultOptions = &Options{}

// NewGame func
func NewGame(gOpts *Options) *Game {
	const nPlayers = 2
	g := Game{
		Name:       gOpts.WithName,
		PlayerList: NewPlayers(nPlayers),
		deck:       NewDeck(),
		BoardSet:   NewPlayedCards(nPlayers),
	}
	var i int
	g.registration = func(n string) error {
		if i >= nPlayers {
			return errors.New("noop: max players reached")
		}
		(*g.PlayerList)[i] = NewPlayer(n)
		i++
		return nil
	}
	return &g
}

func (g *Game) Players() *Players   { return g.PlayerList }
func (g *Game) board() *PlayedCards { return g.BoardSet }
func (g *Game) deckCards() *Deck    { return g.deck }
func (g *Game) BriscolaCard() *Card { return &g.Briscola }
func Register(name string, g *Game) error {
	err := g.registration(name)
	if err != nil {
		return err
	}
	if g.Players().None(func(p Player) bool { return p.Name() == "" }) {
		start(g)
	}
	return nil
}
func (g *Game) InTurn() *Player            { return (*g.Players())[g.next] }
func (g *Game) nextPlayer(next func() int) { g.next = next() }
func (g *Game) roundrobin() int            { return (g.next + 1) % 2 }

func start(g *Game) {
	// distribute cards
	const hndSize = 3
	for i := 0; i < hndSize; i++ {
		for _, p := range *g.Players() {
			p.Hand().Add(g.deck.Top())
		}
	}
	// set first card after distribution as briscola
	g.Briscola = Card{Item: g.deck.Top()}
}

func (g Game) String() string {
	return fmt.Sprintf("(Players: %v, board: %v, Briscola: %v, Deck: %v)", g.PlayerList, g.BoardSet, g.Briscola, g.deck)
}
