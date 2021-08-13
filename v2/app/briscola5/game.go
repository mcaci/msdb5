package briscola5

import (
	"fmt"

	briscolapp "github.com/mcaci/msdb5/v2/app/briscola"
	"github.com/mcaci/msdb5/v2/app/register"
	"github.com/mcaci/msdb5/v2/dom/briscola"
	"github.com/mcaci/msdb5/v2/dom/briscola5"
	"github.com/mcaci/msdb5/v2/pb"
)

// Game struct
type Game struct {
	opts         *Options
	players      misc.Players
	briscolaCard briscola.Card
	side         briscola5.Side
	auctionScore briscola5.AuctionScore
	board        *briscola.PlayedCards
	registration func(string) error
	deck         *briscolapp.Deck
	Callers
}

type Callers struct {
	cal, cmp misc.Player
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
	p, rf := register.NewWithRegistrator(5)
	g := Game{
		opts:         gOpts,
		players:      *p,
		deck:         briscolapp.NewDeck(),
		board:        briscola.NewPlayedCards(5),
		registration: rf,
	}
	return &g
}

func (g *Game) Players() *misc.Players                                      { return &g.players }
func (g *Game) Created(name string) bool                                      { return name == g.opts.WithName }
func (g *Game) WithSide() bool                                                { return g.opts.WithSide }
func (g *Game) Deck() *briscolapp.Deck                                        { return g.deck }
func (g *Game) Side() *briscola5.Side                                         { return &g.side }
func Register(name string, g *Game) error                                     { return g.registration(name) }
func SetAucScore(score briscola5.AuctionScore, g *Game)                       { g.auctionScore = score }
func SetBriscola(card briscola.Card, g *Game)                                 { g.briscolaCard = card }
func SetScoreF(f func(int) (interface{ GetPoints() uint32 }, error), g *Game) { g.opts.WithScoreF = f }
func (c *Callers) SetCaller(p misc.Player)                                  { c.cal = p }
func (c *Callers) SetCompanion(p misc.Player)                               { c.cmp = p }
func (c *Callers) Caller() misc.Player                                      { return c.cal }
func (c *Callers) Companion() misc.Player                                   { return c.cmp }

type Callerer interface {
	Caller() misc.Player
	Companion() misc.Player
}

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
		g.Caller().Name(), g.Companion().Name()+" "+g.briscolaCard.String(), g.auctionScore, g.players, g.side)
}
