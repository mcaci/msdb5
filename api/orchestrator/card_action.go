package orchestrator

import (
	"strings"

	"github.com/nikiforosFreespirit/msdb5/card"
)

func cardAction(request string) (c card.ID, err error) {
	data := strings.Split(request, "#")
	number := data[1]
	seed := data[2]
	return card.Create(number, seed)
}
