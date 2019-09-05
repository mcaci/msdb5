package msg

import (
	"fmt"
	"io"
	"os"

	"github.com/mcaci/msdb5/app/phase"
)

func toML(g roundInformer) {
	f, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		io.WriteString(os.Stdout, err.Error())
	}
	defer f.Close()
	switch g.Phase() {
	case phase.ChoosingCompanion:
		io.WriteString(f, fmt.Sprintf("%s, %s, %d\n", g.CurrentPlayer().Name(), g.Companion().Name(), *(g.AuctionScore())))
	case phase.PlayingCards:
		io.WriteString(f, fmt.Sprintf("%s, %d\n", g.CurrentPlayer().Name(), g.PlayedCard()))
	case phase.End:
		io.WriteString(f, fmt.Sprintf("%s\n", g.CurrentPlayer().Name()))
	}
}
