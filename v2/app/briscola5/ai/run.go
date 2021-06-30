package ai

import (
	"context"
	"log"

	"github.com/mcaci/ita-cards/set"
	briscolapp "github.com/mcaci/msdb5/v2/app/briscola"
	briscolapp5 "github.com/mcaci/msdb5/v2/app/briscola5"
	"github.com/mcaci/msdb5/v2/app/briscola5/auction"
	"github.com/mcaci/msdb5/v2/app/briscola5/companion"
	"github.com/mcaci/msdb5/v2/app/briscola5/end"
	"github.com/mcaci/msdb5/v2/app/briscola5/exchange"
	"github.com/mcaci/msdb5/v2/app/briscola5/play"
	"github.com/mcaci/msdb5/v2/dom/briscola"
	"github.com/mcaci/msdb5/v2/dom/briscola5"
	"github.com/mcaci/msdb5/v2/dom/player"
	"github.com/mcaci/msdb5/v2/dom/team"
	"github.com/mcaci/msdb5/v2/pb"
)

func Run(g *briscolapp5.Game) []uint32 {
	handSize := 7
	if !g.WithSide() {
		handSize++
	}
	// distribute cards to players
	briscola.Distribute(&struct {
		Players  briscola.Players
		Deck     *briscola.Deck
		HandSize int
	}{
		Players:  briscola.Players{Players: briscola5.ToGeneralPlayers(*g.Players())},
		Deck:     g.Deck(),
		HandSize: handSize,
	})
	// set side deck
	for range g.Deck().Cards {
		g.Side().Add(g.Deck().Top())
	}

	// auction phase
	aucInf := auction.Run(struct {
		Players briscola5.Players
		CmpF    func(briscola5.AuctionScore, briscola5.AuctionScore) int8
	}{
		Players: *g.Players(),
		CmpF:    g.CmpF(),
	})
	briscolapp5.SetAucScore(aucInf.Score, g)
	g.Players().SetCaller(aucInf.Caller)

	// card exchange phase
	if g.WithSide() {
		exchange.Run(struct {
			Hand, Side *set.Cards
		}{
			Hand: g.Players().Caller().Hand(),
			Side: &g.Side().Cards,
		})
	}

	// companion choice phase
	cmpInf := companion.Run(
		struct {
			ID      uint8
			Players team.Players
		}{
			ID:      aucInf.Caller,
			Players: briscola5.ToGeneralPlayers(*g.Players()),
		},
	)
	briscolapp5.SetBriscola(cmpInf.Briscola, g)
	g.Players().SetCaller(cmpInf.Companion)

	// play phase
	plInfo := play.Run(struct {
		Players      briscola5.Players
		BriscolaCard briscola.Card
		EndRound     func(*struct {
			PlayedCards  briscola.PlayedCards
			BriscolaCard briscola.Card
		}) (*pb.Index, error)
	}{
		Players:      *g.Players(),
		BriscolaCard: cmpInf.Briscola,
		EndRound:     g.EndRndF(),
	})

	// end phase
	end.Run(struct {
		PlayedCards  briscola.PlayedCards
		Players      team.Players
		BriscolaCard briscola.Card
		Side         briscola5.Side
	}{
		PlayedCards:  plInfo.OnBoard,
		Players:      briscola5.ToGeneralPlayers(*g.Players()),
		BriscolaCard: cmpInf.Briscola,
		Side:         *g.Side(),
	})

	log.Println("Match over", g)

	pls := teams(g.Players())
	method := g.ScoreF()
	if method == nil {
		log.Println("Using default scoring method")
		method = rem(pls)
	}

	scoreIn := &struct {
		Players *briscola.Players
		Method  func(int) (interface{ GetPoints() uint32 }, error)
	}{
		Players: pls,
		Method:  method,
	}
	return briscolapp.Score(scoreIn)
}

func rem(players *briscola.Players) func(int) (interface{ GetPoints() uint32 }, error) {
	return func(i int) (interface{ GetPoints() uint32 }, error) {
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
		return client.Score(context.Background(), toPBCards(*players.At(i).Pile()))
	}
}

func teams(players *briscola5.Players) *briscola.Players {
	t1, t2 := briscola5.ToGeneralPlayers(*players).Part(player.IsInCallers(players))
	pls := briscola.NewPlayers(2)
	pls.Players[0].RegisterAs("Caller team")
	pls.Players[0].Pile().Add(team.CommonPile(t1)...)
	pls.Players[1].RegisterAs("Non Caller team")
	pls.Players[1].Pile().Add(team.CommonPile(t2)...)
	return pls
}
