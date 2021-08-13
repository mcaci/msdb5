package briscola

import (
	"fmt"
	"log"
	"strings"

	"github.com/mcaci/msdb5/v2/dom/briscola"
)

func Score(g *struct {
	Players *briscola.Players
	Method  func(int) (interface{ GetPoints() uint32 }, error)
}) []uint32 {
	scores := make([]uint32, len(*g.Players))
	for i := range *g.Players {
		p, err := g.Method(i)
		if err != nil {
			log.Println(err)
			return []uint32{}
		}
		scores[i] = p.GetPoints()
	}
	return scores
}

func PrintScore(g *struct {
	Players *briscola.Players
	Method  func(int) (interface{ GetPoints() uint32 }, error)
}) string {
	scores := make([]string, len(*g.Players))
	scoresN := Score(g)
	if len(scoresN) == 0 {
		return ""
	}
	for i, s := range scoresN {
		score := fmt.Sprintf("[%s: %d]", (*g.Players)[i].Name(), s)
		log.Println(score)
		scores[i] = score
	}
	return strings.Join(scores, ", ")
}
