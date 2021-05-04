package briscola

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/mcaci/ita-cards/set"
	"github.com/mcaci/msdb5/v2/dom/briscola"
	"github.com/mcaci/msdb5/v2/pb"
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

func ScoreGrpc(g *struct {
	Players briscola.Players
}) string {
	conn := pb.Conn()
	defer conn.Close()
	client := pb.NewBriscolaClient(conn)

	toPBCards := func(cards set.Cards) *pb.Cards {
		pbcards := make([]*pb.CardID, len(cards))
		for i := range pbcards {
			pbcards[i] = &pb.CardID{Id: uint32(cards[i].ToID())}
		}
		return &pb.Cards{Cards: pbcards}
	}

	return Score(&struct {
		Players briscola.Players
		Method  func(int) (interface{ GetPoints() uint32 }, error)
	}{
		Players: g.Players,
		Method: func(i int) (interface{ GetPoints() uint32 }, error) {
			return client.Score(context.Background(), toPBCards(*g.Players.Players[i].Pile()))
		},
	})
}
