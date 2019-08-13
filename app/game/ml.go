package game

import (
	"fmt"
	"io"
	"os"

	"github.com/mcaci/msdb5/app/phase"
)

func (g *Game) handleMLData() PlMsg {
	// log action to file for ml (TODO: WHEN PUSHED OUTSIDE FUNC -> PROBLEM)
	var dest io.Writer
	var text string
	f, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		dest, text = os.Stdout, err.Error()
	}
	// TODO: put back absolutely
	// defer f.Close()
	// write to file for ml
	switch g.Phase() {
	case phase.ChoosingCompanion:
		dest, text = f, fmt.Sprintf("%s, %s, %d\n", g.CurrentPlayer().Name(), g.Companion().Name(), *(g.AuctionScore()))
	case phase.PlayingCards:
		lastPlayed := g.playedCards[len(g.playedCards)-1]
		dest, text = f, fmt.Sprintf("%s, %d\n", g.CurrentPlayer().Name(), lastPlayed)
	case phase.End:
		// write to file who took all cards at last round
		dest, text = f, fmt.Sprintf("%s\n", g.CurrentPlayer().Name())
	}
	return PlMsg{dest, text}
}
