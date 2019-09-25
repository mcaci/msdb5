package pl

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mcaci/ita-cards/card"
	"github.com/mcaci/ita-cards/set"
	"github.com/mcaci/msdb5/dom/auction"
	"github.com/mcaci/msdb5/dom/phase"
	"github.com/mcaci/msdb5/dom/player"
	"golang.org/x/text/message"
)

type statusProvider interface {
	AuctionScore() *auction.Score
	Briscola() card.Item
	CurrentPlayer() *player.Player
	Phase() phase.ID
	PlayedCard() card.Item
	PlayedCards() *set.Cards
}

// TranslateGameStatus func
func TranslateGameStatus(g statusProvider, printer *message.Printer) string {
	var c card.Item
	if g.Phase() == phase.PlayingCards {
		c = g.PlayedCard()
	}
	gameElems := []string{
		g.CurrentPlayer().Name(),
		phases(printer, uint8(g.Phase())),
		translateCard(g.Briscola(), printer),
		translateCards(*g.PlayedCards(), printer),
		translateCard(c, printer),
		strconv.Itoa(int(*g.AuctionScore())),
	}
	for i := range gameElems {
		gameElems[i] = gameElemRef(printer, uint8(i))
	}
	return fmt.Sprintf("%s: %s\n", gameRef(printer), strings.Join(gameElems, ", "))
}
