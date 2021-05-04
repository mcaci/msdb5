package briscola5

import (
	"context"

	"github.com/mcaci/ita-cards/set"
	"github.com/mcaci/msdb5/v2/app/briscola"
	briscolad "github.com/mcaci/msdb5/v2/dom/briscola"
	"github.com/mcaci/msdb5/v2/dom/briscola5"
	"github.com/mcaci/msdb5/v2/dom/team"
	"github.com/mcaci/msdb5/v2/pb"
)

func Score(g *Game) string {
	t1, t2 := briscola5.ToGeneralPlayers(g.players).Part(briscola5.IsInCallers(&g.players))

	pls := briscolad.NewPlayers(2)
	pls.Players[0].RegisterAs("Caller team")
	pls.Players[0].Pile().Add(team.CommonPile(t1)...)
	pls.Players[1].RegisterAs("Non Caller team")
	pls.Players[1].Pile().Add(team.CommonPile(t2)...)

	return briscola.Score(&struct {
		Players briscolad.Players
		Method  func(int) (interface{ GetPoints() uint32 }, error)
	}{
		Players: *pls,
		Method: func(i int) (interface{ GetPoints() uint32 }, error) {
			p := briscolad.Score(*pls.Players[i].Pile())
			return p, nil
		},
	})
}

func ScoreGrpc(g *Game) string {
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

	t1, t2 := briscola5.ToGeneralPlayers(g.players).Part(briscola5.IsInCallers(&g.players))

	pls := briscolad.NewPlayers(2)
	pls.Players[0].RegisterAs("Caller team")
	pls.Players[0].Pile().Add(team.CommonPile(t1)...)
	pls.Players[1].RegisterAs("Non Caller team")
	pls.Players[1].Pile().Add(team.CommonPile(t2)...)

	return briscola.Score(&struct {
		Players briscolad.Players
		Method  func(int) (interface{ GetPoints() uint32 }, error)
	}{
		Players: *pls,
		Method: func(i int) (interface{ GetPoints() uint32 }, error) {
			return client.Score(context.Background(), toPBCards(*pls.Players[i].Pile()))
		},
	})
}
