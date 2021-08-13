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
	"github.com/mcaci/msdb5/v2/app/misc"
	"github.com/mcaci/msdb5/v2/dom/briscola"
	"github.com/mcaci/msdb5/v2/dom/briscola5"
	"github.com/mcaci/msdb5/v2/pb"
)

func Run(g *briscolapp5.Game) []uint32 {
	handSize := 7
	if !g.WithSide() {
		handSize++
	}

	pls := *g.Players()
	// distribute cards to players
	briscolapp.Distribute(&struct {
		Players  misc.Players
		Deck     *briscolapp.Deck
		HandSize int
	}{
		Players:  pls,
		Deck:     g.Deck(),
		HandSize: handSize,
	})
	// set side deck
	for range g.Deck().Cards {
		g.Side().Add(g.Deck().Top())
	}

	// auction phase
	aucInf := auction.Run(struct {
		Players misc.Players
		CmpF    func(briscola5.AuctionScore, briscola5.AuctionScore) int8
	}{
		Players: *g.Players(),
		CmpF:    g.CmpF(),
	})
	briscolapp5.SetAucScore(aucInf.Score, g)
	g.SetCaller((*g.Players())[aucInf.Caller])

	// card exchange phase
	if g.WithSide() {
		exchange.Run(struct {
			Hand, Side *set.Cards
		}{
			Hand: g.Caller().Hand(),
			Side: &g.Side().Cards,
		})
	}

	// companion choice phase
	cmpInf := companion.Run(
		struct {
			Player  misc.Player
			Players misc.Players
		}{
			Player:  (*g.Players())[aucInf.Caller],
			Players: *g.Players(),
		},
	)
	briscolapp5.SetBriscola(cmpInf.Briscola, g)
	g.SetCompanion((*g.Players())[cmpInf.Companion])

	// play phase
	plInfo := play.Run(struct {
		Players      misc.Players
		Caller       misc.Player
		BriscolaCard briscola.Card
		EndRound     func(*struct {
			PlayedCards  briscola.PlayedCards
			BriscolaCard briscola.Card
		}) (*pb.Index, error)
	}{
		Players:      *g.Players(),
		Caller:       g.Caller(),
		BriscolaCard: cmpInf.Briscola,
		EndRound:     g.EndRndF(),
	})

	// end phase
	end.Run(struct {
		PlayedCards  briscola.PlayedCards
		Players      misc.Players
		BriscolaCard briscola.Card
		Side         briscola5.Side
	}{
		PlayedCards:  plInfo.OnBoard,
		Players:      *g.Players(),
		BriscolaCard: cmpInf.Briscola,
		Side:         *g.Side(),
	})

	log.Println("Match over", g)

	t1, t2 := g.Players().Part(misc.IsInCallers(&g.Callers))
	teams := misc.NewPlayers(2)
	(*teams)[0] = misc.New(&misc.Options{For2P: true, Name: "Caller team"})
	(*teams)[0].Pile().Add(misc.CommonPile(t1)...)
	(*teams)[1] = misc.New(&misc.Options{For2P: true, Name: "Non Caller team"})
	(*teams)[1].Pile().Add(misc.CommonPile(t2)...)

	method := g.ScoreF()
	if method == nil {
		log.Println("Using default scoring method")
		method = rem(teams)
	}

	scoreIn := &struct {
		Players *misc.Players
		Method  func(int) (interface{ GetPoints() uint32 }, error)
	}{
		Players: teams,
		Method:  method,
	}
	return briscolapp.Score(scoreIn)
}

func rem(players *misc.Players) func(int) (interface{ GetPoints() uint32 }, error) {
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
		return client.Score(context.Background(), toPBCards(*(*players)[i].Pile()))
	}
}
