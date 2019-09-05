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
	case phase.PlayingCards:
		if g.PlayedCard() != nil {
			io.WriteString(f, fmt.Sprintf("%s\n", g.CurrentPlayer().Name()))
			io.WriteString(f, fmt.Sprintf("%d\n", g.PlayedCard()))
		}
	case phase.End:
		io.WriteString(f, fmt.Sprintf("%s\n", g.CurrentPlayer().Name()))
		io.WriteString(f, fmt.Sprintf("%d\n", g.PlayedCard()))
		io.WriteString(f, fmt.Sprintf("%s\n", g.Caller().Name()))
		io.WriteString(f, fmt.Sprintf("%s\n", g.Companion().Name()))
		io.WriteString(f, fmt.Sprintf("%d\n", *(g.AuctionScore())))
	}
}
