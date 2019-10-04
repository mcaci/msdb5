package cons

import (
	"fmt"
	"io"
	"os"

	"github.com/mcaci/msdb5/app/msg/score"
	"github.com/mcaci/msdb5/dom/phase"
	"github.com/mcaci/msdb5/dom/player"
	"github.com/mcaci/msdb5/dom/team"
)

type osInformer interface {
	CurrentPlayer() *player.Player
	Players() team.Players
	Phase() phase.ID

	team.Callers

	RoundError() error
}

// Write func
func Write(g osInformer, inputRequest, origin string) {
	rErr := g.RoundError()
	if rErr != nil {
		errMsg := fmtErr(g, inputRequest, rErr)
		io.WriteString(os.Stdout, errMsg)
		return
	}
	senderPred := player.MatchingHost(origin)
	s := g.Players().At(g.Players().MustFind(senderPred))
	senderInfo := fmt.Sprintf("New Action by %s: %s\nSender info: %+v\nGame info: %+v\n", s.Name(), inputRequest, s, g)
	io.WriteString(os.Stdout, senderInfo)

	// compute score
	t1, t2 := g.Players().Part(team.IsInCallersPred(g))
	scoreMsg := fmt.Sprintf("Scores -> Callers: %d; Others: %d\n", score.Sum(team.CommonPile(t1)), score.Sum(team.CommonPile(t2)))
	io.WriteString(os.Stdout, scoreMsg)
}

func fmtErr(g osInformer, inputRequest string, rErr error) string {
	errMsg := fmt.Sprintf("Error: %+v\n", rErr)
	if rErr == phase.ErrUnexpectedPhase {
		errMsg = fmt.Sprintf("Phase is not %s but %s\n", inputRequest, g.Phase())
	}
	if rErr == player.ErrUnexpectedPlayer {
		errMsg = fmt.Sprintf("Expecting player %s to play\n", g.CurrentPlayer().Name())
	}
	return errMsg
}
