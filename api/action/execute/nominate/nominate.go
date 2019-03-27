package nominate

import (
	"strings"

	"github.com/nikiforosFreespirit/msdb5/api/action"
	"github.com/nikiforosFreespirit/msdb5/card"
	"github.com/nikiforosFreespirit/msdb5/player"
	"github.com/nikiforosFreespirit/msdb5/playerset"
)

type CompanionStruct struct {
	request, origin string
	players         playerset.Players
	set             func(card.ID, *player.Player)
}

func NewCompanion(request, origin string, players playerset.Players,
	set func(card.ID, *player.Player)) action.Executer {
	return &CompanionStruct{request, origin, players, set}
}

func (cs CompanionStruct) Do(p *player.Player) error {
	data := strings.Split(cs.request, "#")
	number := data[1]
	seed := data[2]
	c, err := card.Create(number, seed)
	if err != nil {
		return err
	}
	pl, err := cs.players.Find(func(p *player.Player) bool { return p.Has(c) })
	if err != nil {
		return err
	}
	cs.set(c, pl)
	return nil
}
