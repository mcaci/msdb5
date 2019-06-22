package team

import (
	"fmt"

	"github.com/nikiforosFreespirit/msdb5/dom/player"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

// Players struct
type Players []*player.Player

// Add func
func (playerSet *Players) Add(p *player.Player) {
	*playerSet = append(*playerSet, p)
}

// Find func
func (playerSet Players) Find(predicate func(p *player.Player) bool) (int, *player.Player, error) {
	for i, p := range playerSet {
		if predicate(p) {
			return i, p, nil
		}
	}
	printer := message.NewPrinter(language.English)
	msg := printer.Sprint("Player not found")
	return -1, nil, fmt.Errorf(msg)
}

func (playerSet Players) String() (str string) {
	printer := message.NewPrinter(language.English)
	for _, p := range playerSet {
		str += printer.Sprintf("- %+v -", *p)
	}
	return

}
