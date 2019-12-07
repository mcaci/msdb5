package pl

import (
	"fmt"
	"strings"

	"github.com/mcaci/ita-cards/card"
	"github.com/mcaci/msdb5/dom/player"
	"github.com/mcaci/msdb5/dom/team"
	"golang.org/x/text/message"
)

func translateTeam(p *player.Player, g team.Callers, printer *message.Printer) string {
	return fmt.Sprintf("%s: %s %s", endRef(printer), translateBool(printer, team.IsInCallers(g)(p), teams), allBriscolaRef(printer))
}

func translatePlayer(pl *player.Player, br card.Item, printer *message.Printer) string {
	var seed *card.Seed
	if br.Number() > 0 {
		s := br.Seed()
		seed = &s
	}
	plElems := []string{
		pl.Name(),
		translateHand(*pl.Hand(), seed, printer),
		translateCards(*pl.Pile(), printer),
		translateBool(printer, player.Folded(pl), yesNoRef),
	}
	for i := range plElems {
		val := plElems[i]
		plElems[i] = plElemRef(printer, uint8(i)) + ": " + val
	}
	return fmt.Sprintf("%s: %s\n", plRef(printer), strings.Join(plElems, ", "))
}

func translateBool(p *message.Printer, info bool, format func(*message.Printer, uint8) string) string {
	if info {
		return format(p, 0)
	}
	return format(p, 1)
}
