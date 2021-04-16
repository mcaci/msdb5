package play

import (
	"context"
	"log"

	"github.com/mcaci/ita-cards/set"
	"github.com/mcaci/msdb5/v2/dom/briscola"
	"github.com/mcaci/msdb5/v2/dom/briscola5"
	"github.com/mcaci/msdb5/v2/pb"
)

type RoundOpts struct {
	PlIdx        uint8
	PlHand       *set.Cards
	CardIdx      uint8
	PlayedCards  *briscola5.PlayedCards
	NPlayers     uint8
	BriscolaCard briscola.Card
}

type RoundInfo struct {
	OnBoard *briscola5.PlayedCards
	NextPl  uint8
	NextRnd bool
}

func Round(g *RoundOpts) *RoundInfo {
	defaultInfo := &RoundInfo{
		OnBoard: g.PlayedCards,
		NextPl:  roundRobin(g.PlIdx, 1, g.NPlayers),
	}
	if len(*g.PlHand) <= 0 {
		return defaultInfo
	}
	err := set.MoveOne(&(*g.PlHand)[g.CardIdx], g.PlHand, g.PlayedCards.Cards)
	if err != nil {
		return defaultInfo
	}
	if !isRoundOngoing(*g.PlayedCards.Cards) {
		// end current round
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
		toBoard := func(cards set.Cards) *pb.Board {
			pbcards := make([]*pb.CardID, len(cards))
			for i := range pbcards {
				pbcards[i] = &pb.CardID{Id: uint32(cards[i].ToID())}
			}
			return &pb.Board{Briscola: uint32(g.BriscolaCard.Seed()), Cards: toPBCards(cards)}
		}

		win, err := client.Winner(context.Background(), toBoard(*g.PlayedCards.Cards))
		if err != nil {
			log.Println(err)
		}

		return &RoundInfo{
			OnBoard: g.PlayedCards,
			NextPl:  roundRobin(g.PlIdx, uint8(win.Id)+1, g.NPlayers),
			NextRnd: true,
		}
	}
	return &RoundInfo{
		OnBoard: g.PlayedCards,
		NextPl:  roundRobin(g.PlIdx, 1, g.NPlayers),
	}
}

func isRoundOngoing(playedCards set.Cards) bool { return len(playedCards) < 5 }
func roundRobin(idx, off, size uint8) uint8     { return (idx + off) % size }
