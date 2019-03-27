package join

import (
	"strings"

	"github.com/nikiforosFreespirit/msdb5/api/action"
	"github.com/nikiforosFreespirit/msdb5/player"
)

type JoinStruct struct {
	request, origin string
}

func NewJoin(request, origin string) action.Executer {
	return &JoinStruct{request, origin}
}

func (js JoinStruct) Do(p *player.Player) error {
	name := strings.Split(js.request, "#")[1]
	p.Join(name, js.origin)
	return nil
}
