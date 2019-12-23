package pl

import (
	"fmt"
	"strconv"

	"github.com/mcaci/msdb5/dom/phase"
	"github.com/mcaci/msdb5/dom/team"
	"golang.org/x/text/message"
)

func translateErr(g plInformer, printer *message.Printer, inputRequest string, rErr error) string {
	errMsg := fmt.Sprintf("%s: %+v\n", errRef(printer), rErr)
	if rErr == phase.ErrUnexpectedPhase {
		errMsg = errMsgRef(printer, strconv.Itoa(int(g.Phase())), inputRequest)
	}
	if rErr == team.ErrPlayerNotFound {
		errMsg = errMsgRef(printer, g.CurrentPlayer().Name(), inputRequest)
	}
	return errMsg
}
