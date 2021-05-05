package briscola

import (
	"fmt"
	"log"
	"strings"

	"github.com/mcaci/msdb5/v2/dom/briscola"
)

func Score(g *struct {
	Players briscola.Players
	Method  func(int) (interface{ GetPoints() uint32 }, error)
}) string {
	scores := make([]string, len(g.Players.Players))
	for i := range g.Players.Players {
		p, err := g.Method(i)
		if err != nil {
			log.Println(err)
			return ""
		}
		score := fmt.Sprintf("[%s: %d]", g.Players.Players[i].Name(), p.GetPoints())
		log.Println(score)
		scores[i] = score
	}
	return strings.Join(scores, ", ")
}
