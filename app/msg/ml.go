package msg

import (
	"fmt"
	"io"
	"os"

	"github.com/mcaci/msdb5/app/phase"
)

func HandleMLData(g roundInformer) (io.Writer, string) {
	// log action to file for ml (TODO: WHEN PUSHED OUTSIDE FUNC -> PROBLEM)
	var dest io.Writer
	var text string
	f, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		dest, text = os.Stdout, err.Error()
	}
	defer f.Close()
	// write to file for ml
	switch g.Phase() {
	case phase.ChoosingCompanion:
		dest, text = f, fmt.Sprintf("%s, %s, %d\n", g.CurrentPlayer().Name(), g.Companion().Name(), *(g.AuctionScore()))
	case phase.PlayingCards:
		dest, text = f, fmt.Sprintf("%s, %d\n", g.CurrentPlayer().Name(), g.PlayedCard())
	case phase.End:
		dest, text = f, fmt.Sprintf("%s\n", g.CurrentPlayer().Name())
	}
	return dest, text
}