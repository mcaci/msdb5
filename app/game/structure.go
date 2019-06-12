package game

import (
	"container/list"
	"fmt"

	"github.com/nikiforosFreespirit/msdb5/app/phase"
	"github.com/nikiforosFreespirit/msdb5/dom/auction"
	"github.com/nikiforosFreespirit/msdb5/dom/card"
	"github.com/nikiforosFreespirit/msdb5/dom/deck"
	"github.com/nikiforosFreespirit/msdb5/dom/player"
	"github.com/nikiforosFreespirit/msdb5/dom/team"
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
}

// NewGame func
func NewGame(withSide bool) *Game {
	g := new(Game)
	g.withSide = withSide
	makePlayers(g)
	distributeCards(g)
	trackActing(&g.lastPlaying, g.players[0])
	return g
}

func makePlayers(g *Game) {
	for i := 0; i < 5; i++ {
		g.players.Add(*player.New())
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
			trackActing(&g.lastPlaying, g.players[i%5])
			g.CurrentPlayer().Draw(d.Supply)
		}
	}
}

func trackActing(lastPlaying *list.List, actingPlayer *player.Player) {
	lastPlaying.PushFront(actingPlayer)
	if lastPlaying.Len() > 2 {
		lastPlaying.Remove(lastPlaying.Back())
	}
}

func (g *Game) AuctionScore() auction.Score         { return g.auctionScore }
func (g *Game) Companion() *player.Player           { return g.companion }
func (g *Game) LastCardPlayed() card.ID             { return g.playedCards[len(g.playedCards)-1] }
func (g *Game) Phase() phase.ID                     { return g.phase }
func (g *Game) SideDeck() deck.Cards                { return g.side }
func (g *Game) IsSideUsed() bool                    { return g.withSide }
func (g *Game) LastPlayer() *player.Player          { return g.lastPlaying.Back().Value.(*player.Player) }
func (g *Game) CurrentPlayer() *player.Player       { return g.lastPlaying.Front().Value.(*player.Player) }
func (g *Game) Sender(origin string) *player.Player { return g.players[g.senderIndex(origin)] }

func (g *Game) briscola() card.Seed  { return g.briscolaCard.Seed() }
func (g *Game) cardsOnTheBoard() int { return len(g.playedCards) }
func (g *Game) senderIndex(origin string) int {
	rq := newReq("Origin", origin)
	criteria := findCriteria(g, rq)
	index, _, _ := g.players.Find(criteria)
	return index
}

func (g Game) String() (str string) {
	return fmt.Sprintf("(Turn of: %s, Companion is: %s, Played cards: %+v, Auction score: %d, Phase: %d)",
		g.CurrentPlayer().Name(), g.briscolaCard, g.playedCards, g.auctionScore, g.phase)
}
