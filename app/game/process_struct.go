package game

import (
	"fmt"
	"io"
	"os"

	"github.com/mcaci/ita-cards/card"
	"github.com/mcaci/msdb5/app/phase"
	"github.com/mcaci/msdb5/dom/player"
	"github.com/mcaci/msdb5/dom/team"
)

type report struct {
	reports []PlMsg
	err     error
}

func (r *report) msg(writer io.Writer, msg string) {
	r.reports = append(r.reports, PlMsg{writer, msg})
}

func (r *report) error(s team.SenderInformation, action string, err error) {
	r.msg(os.Stdout, fmt.Sprintf("New Action by %s: %s\n", team.Sender(s).Name(), action))
	if err == nil {
		return
	}
	errMsg := fmt.Sprintf("Error: %+v\n", err)
	r.msg(os.Stdout, errMsg)
	r.msg(team.Sender(s), errMsg)
	r.err = err
}

type phaseInfo struct {
	phase  phase.ID
	action string
}

func (s phaseInfo) Action() string  { return s.action }
func (s phaseInfo) Phase() phase.ID { return s.phase }

type expectedSenderInfo struct {
	senderInfo
	p *player.Player
}

func (s expectedSenderInfo) CurrentPlayer() *player.Player { return s.p }

type senderInfo struct {
	players team.Players
	origin  string
}

func (s senderInfo) From() string          { return s.origin }
func (s senderInfo) Players() team.Players { return s.players }

type gameRound struct {
	*Game
	c    *card.Item
	cErr error
	val  string
}

func (g gameRound) Card() (*card.Item, error) { return g.c, g.cErr }
func (g gameRound) Value() string             { return g.val }
