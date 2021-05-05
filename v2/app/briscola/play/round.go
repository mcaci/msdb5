package play

import (
	"context"
	"log"

	"github.com/mcaci/ita-cards/set"
	"github.com/mcaci/msdb5/v2/dom/briscola"
	"github.com/mcaci/msdb5/v2/pb"
)

type RoundOpts struct {
	PlIdx        uint8
	PlHand       *set.Cards
	CardIdx      uint8
	PlayedCards  *briscola.PlayedCards
	NPlayers     uint8
	BriscolaCard briscola.Card
	EndRound     func(*struct {
		PlayedCards  briscola.PlayedCards
		BriscolaCard briscola.Card
	}) (*pb.Index, error)
}

type RoundInfo struct {
	OnBoard *briscola.PlayedCards
	NextPl  uint8
	NextRnd bool
}

func Round(rOpts *RoundOpts) *RoundInfo {
	isRoundOngoing := func(playedCards set.Cards) bool { return len(playedCards) < int(rOpts.NPlayers) }
	roundRobin := func(idx, off uint8) uint8 { return (idx + off) % rOpts.NPlayers }
	rInfo := &RoundInfo{
		OnBoard: rOpts.PlayedCards,
		NextPl:  roundRobin(rOpts.PlIdx, 1),
	}
	// no cards in hand
	if len(*rOpts.PlHand) <= 0 {
		return rInfo
	}
	// play card
	err := set.MoveOne(&(*rOpts.PlHand)[rOpts.CardIdx], rOpts.PlHand, rOpts.PlayedCards.Cards)
	if err != nil {
		log.Println(err)
		return rInfo
	}
	rInfo.OnBoard = rOpts.PlayedCards
	// round is ongoing
	if isRoundOngoing(*rOpts.PlayedCards.Cards) {
		return rInfo
	}
	// round is finished
	if rOpts.EndRound == nil {
		log.Println("using default end round from remote")
		rOpts.EndRound = endRemote
	}
	win, err := rOpts.EndRound(&struct {
		PlayedCards  briscola.PlayedCards
		BriscolaCard briscola.Card
	}{
		PlayedCards:  *rOpts.PlayedCards,
		BriscolaCard: rOpts.BriscolaCard,
	})
	if err != nil {
		log.Println(err)
		return rInfo
	}
	rInfo.NextPl = roundRobin(rOpts.PlIdx, uint8(win.Id)+1)
	rInfo.NextRnd = true
	return rInfo
}

func endRemote(opts *struct {
	PlayedCards  briscola.PlayedCards
	BriscolaCard briscola.Card
}) (*pb.Index, error) {
	conn := pb.Conn()
	defer conn.Close()
	client := pb.NewBriscolaClient(conn)

	toBoard := func(cards set.Cards) *pb.Board {
		pbcards := make([]*pb.CardID, len(cards))
		for i := range pbcards {
			pbcards[i] = &pb.CardID{Id: uint32(cards[i].ToID())}
		}
		return &pb.Board{Briscola: uint32(opts.BriscolaCard.Seed()), Cards: &pb.Cards{Cards: pbcards}}
	}
	return client.Winner(context.Background(), toBoard(*opts.PlayedCards.Cards))
}
