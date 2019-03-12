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
	phaseData    map[string]func(string, string) phaseData
}

// NewGame func
func NewGame() *Game {
	g := new(Game)
	makePlayers(g)
	playersDrawAllCards(&g.players)
	initPlayerActions(g)
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

func initPlayerActions(g *Game) {
	g.phaseData = make(map[string]func(string, string) phaseData)
	g.phaseData["Join"] = joinData
	g.phaseData["Auction"] = g.raiseAuctionData
	g.phaseData["Companion"] = g.nominateData
	g.phaseData["Card"] = g.playData
}

// NewAction func
func NewAction() api.Action {
	return NewGame()
}

func (g *Game) setCompanion(c card.ID) (err error) {
	pl, err := g.players.Find(func(p *player.Player) bool { return p.Has(c) })
	if err != nil {
		return
	}
	g.companion = *companion.New(c, pl)
	return
}

func nextPlayer(playerInTurn uint8) uint8 { return (playerInTurn + 1) % 5 }

func isExpectedPlayer(p *player.Player, g *Game, origin string) bool {
	return p.IsSame(g.players[g.playerInTurn]) && p.IsSameHost(origin)
}

func (g *Game) nextPhase(predicate func(playerset.Players, func(*player.Player) bool) bool, playerPredicate func(*player.Player) bool) {
	if predicate(g.players, playerPredicate) {
		g.phase++
	}
}
func (g *Game) phaseCheck(current phase) (err error) {
	if g.phase != current {
		err = errors.New("Phase is not " + strconv.Itoa(int(current)))
	}
	return
}
func (g *Game) nextPlayer(generateIndex func(uint8) uint8) {
	g.playerInTurn = generateIndex(g.playerInTurn)
}

// Info func
func (g Game) Info() []display.Info {
	gameInfo := g.info.Info()
	plInTurn := display.Wrap("Turn of", g.players[g.playerInTurn].Name())
	gameInfo = append(gameInfo, plInTurn...)
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
	phase := display.NewInfo("Phase", ":", strconv.Itoa(int(g.phase)), ";")
	gameInfo = append(gameInfo, phase)
	return display.All(display.Wrap("Game", gameInfo...)...)
}
