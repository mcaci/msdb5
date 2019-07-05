package game

import (
	"container/list"

	"github.com/nikiforosFreespirit/msdb5/app/notify"
	"github.com/nikiforosFreespirit/msdb5/app/phase"
	"github.com/nikiforosFreespirit/msdb5/app/request"
	"github.com/nikiforosFreespirit/msdb5/app/track"
	"github.com/nikiforosFreespirit/msdb5/dom/auction"
	"github.com/nikiforosFreespirit/msdb5/dom/card"
	"github.com/nikiforosFreespirit/msdb5/dom/deck"
	"github.com/nikiforosFreespirit/msdb5/dom/player"
	"github.com/nikiforosFreespirit/msdb5/dom/team"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
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
	for _, player := range g.players {
		if player.IsSameHost("") {
			player.Join(origin)
			player.Attach(channel)
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
			g.CurrentPlayer().Draw(d.Supply)
		}
	}
}

func (g *Game) AuctionScore() *auction.Score        { return &g.auctionScore }
func (g *Game) Briscola() card.Seed                 { return g.briscolaCard.Seed() }
func (g *Game) Caller() *player.Player              { return g.caller }
func (g *Game) Companion() *player.Player           { return g.companion }
func (g *Game) CurrentPlayer() *player.Player       { return g.lastPlaying.Front().Value.(*player.Player) }
func (g *Game) IsSideUsed() bool                    { return g.withSide }
func (g *Game) LastCardPlayed() card.ID             { return g.playedCards[len(g.playedCards)-1] }
func (g *Game) LastPlayer() *player.Player          { return g.lastPlaying.Back().Value.(*player.Player) }
func (g *Game) LastPlaying() *list.List             { return &g.lastPlaying }
func (g *Game) Phase() phase.ID                     { return g.phase }
func (g *Game) Players() team.Players               { return g.players }
func (g *Game) PlayedCards() *deck.Cards            { return &g.playedCards }
func (g *Game) Sender(origin string) *player.Player { return g.players[g.SenderIndex(origin)] }
func (g *Game) SideDeck() *deck.Cards               { return &g.side }
func (g *Game) Lang() language.Tag                  { return g.lang }
func (g *Game) CardsOnTheBoard() int                { return len(g.playedCards) }
func (g *Game) SenderIndex(origin string) int {
	rq := request.New("Origin", origin)
	criteria := request.FindCriteria(g, rq)
	index, _ := g.players.Find(criteria)
	return index
}

func (g Game) String() (str string) {
	printer := message.NewPrinter(g.lang)

	return printer.Sprintf("(Turn of: %s, Companion is: %s, Played cards: %s, Auction score: %d, Phase: %d)",
		g.CurrentPlayer().Name(), notify.TranslateCard(g.briscolaCard, printer), notify.TranslateCards(g.playedCards, printer), g.auctionScore, g.phase)
}
