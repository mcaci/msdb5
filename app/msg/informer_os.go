package msg

import (
	"fmt"
	"io"
	"os"

	"github.com/mcaci/msdb5/app/phase"
	"github.com/mcaci/msdb5/app/score"
	"github.com/mcaci/msdb5/dom/briscola"
	"github.com/mcaci/msdb5/dom/player"
)

func toOS(g roundInformer, inputRequest, origin string) {
	rErr := g.RoundError()
	if rErr != nil {
		errMsg := fmtErr(g, inputRequest, rErr)
		io.WriteString(os.Stdout, errMsg)
		return
	}
	s := sender(senderInfo{g.Players(), origin})
	senderInfo := fmt.Sprintf("New Action by %s: %s\nSender info: %+v\nGame info: %+v\n", s.Name(), inputRequest, s, g)
	io.WriteString(os.Stdout, senderInfo)

	// compute score
	scoreTeam1, scoreTeam2 := score.Calc(g, g.Players(), briscola.Points)
	scoreMsg := fmt.Sprintf("Scores -> Callers: %d; Others: %d\n", scoreTeam1, scoreTeam2)
	io.WriteString(os.Stdout, scoreMsg)
}

func fmtErr(g roundInformer, inputRequest string, rErr error) string {
	errMsg := fmt.Sprintf("Error: %+v\n", rErr)
	if rErr == phase.ErrUnexpectedPhase {
		errMsg = fmt.Sprintf("Phase is not %s but %s\n", inputRequest, g.Phase())
	}
	if rErr == player.ErrUnexpectedPlayer {
		errMsg = fmt.Sprintf("Expecting player %s to play\n", g.CurrentPlayer().Name())
	}
	return errMsg
}
