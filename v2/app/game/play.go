package game

import (
	"math/rand"
	"time"

	"github.com/mcaci/ita-cards/card"
	"github.com/mcaci/ita-cards/set"
	"github.com/mcaci/msdb5/v2/app/collect"
	"github.com/mcaci/msdb5/v2/app/track"
	"github.com/mcaci/msdb5/v2/dom/briscola"
	"github.com/mcaci/msdb5/v2/dom/phase"
	"github.com/mcaci/msdb5/v2/dom/player"
	"github.com/mcaci/msdb5/v2/dom/team"
)

func runPlay(g *Game) {
	for g.phase == phase.PlayingCards {
		pl := g.CurrentPlayer()
		hnd := pl.Hand()
		if len(*hnd) > 0 {
			rand.Seed(time.Now().Unix())
			idx := rand.Intn(len(*hnd))
			crd := (*hnd)[idx]
			index := hnd.Find(crd)
			g.playedCards.Add((*hnd)[index])
			*hnd = append((*hnd)[:index], (*hnd)[index+1:]...)
		}

		// next phase
		if g.players.All(player.EmptyHanded) || isAnticipatedEnd(g) {
			g.phase++
		}

		// next player
		nextPlayer := roundRobin(g.CurrentPlayerIndex(), 1, numberOfPlayers)
		if !g.IsRoundOngoing() {
			// end current round
			winningCardIndex := indexOfWinningCard(g.playedCards, g.briscolaCard.Seed())
			nextPlayer = roundRobin(nextPlayer, winningCardIndex, numberOfPlayers)

			// collect cards
			cardToCollect := collect.Collector(g.phase, g.players, &g.side, &g.playedCards)
			set.Move(cardToCollect(), g.players[nextPlayer].Pile())
		}
		track.Player(&g.lastPlaying, g.players[nextPlayer])

	}
}

func indexOfWinningCard(cardsOnTheTable set.Cards, b card.Seed) uint8 {
	base := cardsOnTheTable[0]
	max := 0
	for i, other := range cardsOnTheTable {
		if winningCard(base, other, b) == other {
			base = other
			max = i
		}
	}
	return uint8(max)
}

func winningCard(base, other card.Item, b card.Seed) card.Item {
	if &base == nil || doesOtherCardWin(base, other, b) {
		base = other
	}
	return base
}

func doesOtherCardWin(first, other card.Item, briscola card.Seed) bool {
	otherIsBriscola := other.Seed() == briscola
	isSameSeed := first.Seed() == other.Seed()
	return (!isSameSeed && otherIsBriscola) || isOtherHigher(first, other)
}

func isOtherHigher(first, other card.Item) bool {
	isSameSeed := first.Seed() == other.Seed()
	isOtherGreaterOnPoints := briscola.Points(first) < briscola.Points(other)
	isSamePoints := briscola.Points(first) == briscola.Points(other)
	isOtherGreaterOnNumber := first.Number() < other.Number()
	return isSameSeed && ((isSamePoints && isOtherGreaterOnNumber) || isOtherGreaterOnPoints)
}

func isAnticipatedEnd(g *Game) bool {
	var isAnticipatedEnd bool
	const limit = 3
	roundsBefore := uint8(len(*g.players[0].Hand()))
	if roundsBefore <= limit {
		isNewRoundToStart := len(g.playedCards) == 5
		isAnticipatedEnd = isNewRoundToStart && predict(g, roundsBefore)
	}
	return isAnticipatedEnd
}

func predict(g *Game, roundsBefore uint8) bool {
	highbriscolaCard := serie(g.briscolaCard.Seed())
	var teams [2]bool
	var cardsChecked uint8
	for _, card := range highbriscolaCard {
		i, err := g.players.Index(player.IsCardInHand(card))
		if err != nil { // no one has card
			continue
		}
		p := g.players.At(i)
		isPlayerInCallersTeam := team.IsInCallers(g)(p)
		teams[0] = teams[0] || isPlayerInCallersTeam
		teams[1] = teams[1] || !isPlayerInCallersTeam
		if teams[0] == teams[1] {
			return false
		}
		cardsChecked++
		if cardsChecked == roundsBefore {
			return true
		}
	}
	return false
}

func serie(briscola card.Seed) set.Cards {
	serie := []uint8{1, 3, 10, 9, 8, 7, 6, 5, 4, 2}
	cards := make(set.Cards, len(serie))
	for i, id := range serie {
		cards[i] = *card.MustID(id + 10*uint8(briscola))
	}
	return cards
}
