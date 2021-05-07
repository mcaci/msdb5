package briscola5

import (
	"fmt"

	"github.com/mcaci/msdb5/v2/dom/briscola"
	"github.com/mcaci/msdb5/v2/dom/briscola5"
	"github.com/mcaci/msdb5/v2/pb"
)

// Game struct
type Game struct {
	opts         *Options
	players      briscola5.Players
	briscolaCard briscola.Card
	side         briscola5.Side
	auctionScore briscola5.AuctionScore
	board        *briscola.PlayedCards
	registration func(string) error
	deck         *briscola.Deck
}

type Options struct {
	WithSide     bool
	WithName     string
	WithCmpF     func(briscola5.AuctionScore, briscola5.AuctionScore) int8
	WithScoreF   func(int) (interface{ GetPoints() uint32 }, error)
	WithEndRound func(*struct {
		PlayedCards  briscola.PlayedCards
		BriscolaCard briscola.Card
	}) (*pb.Index, error)
}

var WithDefaultOptions = &Options{}

func NewGame(gOpts *Options) *Game {
	g := Game{
		opts:    gOpts,
		players: *briscola5.NewPlayers(),
		deck:    briscola.NewDeck(),
		board:   briscola.NewPlayedCards(5),
	}
	g.registration = g.players.Registration()
	return &g
}

func (g *Game) Players() *briscola5.Players                                   { return &g.players }
func (g *Game) Created(name string) bool                                      { return name == g.opts.WithName }
func (g *Game) WithSide() bool                                                { return g.opts.WithSide }
func (g *Game) Deck() *briscola.Deck                                          { return g.deck }
func (g *Game) Side() *briscola5.Side                                         { return &g.side }
func Register(name string, g *Game) error                                     { return g.registration(name) }
func SetAucScore(score briscola5.AuctionScore, g *Game)                       { g.auctionScore = score }
func SetBriscola(card briscola.Card, g *Game)                                 { g.briscolaCard = card }
func SetScoreF(f func(int) (interface{ GetPoints() uint32 }, error), g *Game) { g.opts.WithScoreF = f }

func (g *Game) CmpF() func(briscola5.AuctionScore, briscola5.AuctionScore) int8 {
	return g.opts.WithCmpF
}
func (g *Game) EndRndF() func(*struct {
	PlayedCards  briscola.PlayedCards
	BriscolaCard briscola.Card
}) (*pb.Index, error) {
	return g.opts.WithEndRound
}
func (g *Game) ScoreF() func(int) (interface{ GetPoints() uint32 }, error) {
	return g.opts.WithScoreF
}

func (g Game) String() string {
	return fmt.Sprintf("(Caller is: %s,\n Companion is: %s,\n Auction score: %d,\n Players: %v,\n Side Deck: %v)",
		g.players.Caller().Name(), g.players.Companion().Name()+" "+g.briscolaCard.String(), g.auctionScore, g.players, g.side)
}
