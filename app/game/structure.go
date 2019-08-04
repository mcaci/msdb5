package game

import (
	"container/list"
	"fmt"

	"github.com/mcaci/msdb5/app/phase"
	"github.com/mcaci/msdb5/app/track"
	"github.com/mcaci/msdb5/dom/auction"
	"github.com/mcaci/msdb5/dom/card"
	"github.com/mcaci/msdb5/dom/deck"
	"github.com/mcaci/msdb5/dom/player"
	"github.com/mcaci/msdb5/dom/team"
	"golang.org/x/text/language"
)

// Game struct
type Game struct {
	lastPlaying  list.List
	players      team.Players
	caller       *player.Player
	companion    *player.Player
	briscolaCard card.ID
	withSide     bool
	side         deck.Cards
	playedCards  deck.Cards
	auctionScore auction.Score
	phase        phase.ID
	lang         language.Tag
}

// NewGame func
func NewGame(withSide bool, lang language.Tag) *Game {
	g := new(Game)
	g.withSide = withSide
	makePlayers(g)
	distributeCards(g)
	track.Player(&g.lastPlaying, g.players[0])
	g.lang = lang
	return g
}

func makePlayers(g *Game) {
	for i := 0; i < 5; i++ {
		g.players.Add(player.New())
	}
}

// Join func
func (g *Game) Join(origin string, channel chan []byte) {
	for _, p := range g.players {
		if p.IsSameHost("") {
			p.Join(origin)
			p.Attach(channel)
			break
		}
	}
}

func distributeCards(g *Game) {
	d := deck.New()
	for i := 0; i < deck.DeckSize; i++ {
		if g.withSide && i >= deck.DeckSize-5 {
			g.side.Add(d.Supply())
		} else {
			track.Player(&g.lastPlaying, g.players[i%5])
			g.CurrentPlayer().Hand().Add(d.Supply())
		}
	}
}

func (g *Game) AuctionScore() *auction.Score  { return &g.auctionScore }
func (g *Game) Briscola() card.ID             { return g.briscolaCard }
func (g *Game) Caller() *player.Player        { return g.caller }
func (g *Game) Companion() *player.Player     { return g.companion }
func (g *Game) CurrentPlayer() *player.Player { return g.lastPlaying.Front().Value.(*player.Player) }
func (g *Game) IsSideUsed() bool              { return g.withSide }
func (g *Game) LastPlayer() *player.Player    { return g.lastPlaying.Back().Value.(*player.Player) }
func (g *Game) LastPlaying() *list.List       { return &g.lastPlaying }
func (g *Game) Phase() phase.ID               { return g.phase }
func (g *Game) Players() team.Players         { return g.players }
func (g *Game) PlayedCards() *deck.Cards      { return &g.playedCards }
func (g *Game) SideDeck() *deck.Cards         { return &g.side }
func (g *Game) Lang() language.Tag            { return g.lang }
func (g *Game) IsRoundOngoing() bool          { return len(g.playedCards) < 5 }

func (g *Game) SetAuction(s auction.Score) { g.auctionScore = s }
func (g *Game) SetBriscola(c card.ID)      { g.briscolaCard = c }
func (g *Game) SetCompanion(idx uint8)     { g.companion = g.players[idx] }

func (g Game) String() string {
	return fmt.Sprintf("(Turn of: %s, Companion is: %s, Played cards: %v, Auction score: %d, Phase: %s, Players: %v, Side Deck: %v, Last Players: %v)",
		g.CurrentPlayer().Name(), g.briscolaCard, g.playedCards, g.auctionScore, g.phase, g.players, g.side, g.lastPlaying)
}
