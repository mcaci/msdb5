package board

import (
	"log"
	"strconv"
	"strings"

	"github.com/nikiforosFreespirit/msdb5/briscola"
	"github.com/nikiforosFreespirit/msdb5/card"
	"github.com/nikiforosFreespirit/msdb5/player"
)

// Action interface
func (b *Board) Action(request, origin string) {
	data := strings.Split(string(request), "#")
	switch data[0] {
	case "Join":
		b.Join(data[1], origin)
	case "Auction":
		b.RaiseAuction(data[1], origin)
	case "Companion":
		b.Nominate(data[1], data[2], origin)
	case "Card":
		b.Play(data[1], data[2], origin)
	}
}

const minScore = 61
const maxScore = 120

// RaiseAuction func
func (b *Board) RaiseAuction(score, origin string) error {
	prevScore := b.AuctionScore()
	intScore, err := strconv.Atoi(score) // TODO: THIS ERR IS LOST IF ERR BELOW IS NOT
	if err != nil {
		log.Printf("Error was raised during auction: %v\n", err)
	}
	currentScore := uint8(intScore)

	p, err := b.Players().Find(func(p *player.Player) bool { return p.Host() == origin })
	if err == nil {
		updateAuction(0, prevScore, currentScore, p.SetAuctionScore)
	}
	updateAuction(prevScore, prevScore, currentScore, b.SetAuctionScore)
	return err
}

func updateAuction(baseScore, prevScore, currentScore uint8, set func(auctionScore uint8)) {
	if prevScore > 0 && prevScore >= currentScore {
		set(baseScore)
	} else if currentScore < minScore {
		set(minScore)
	} else if currentScore > maxScore {
		set(maxScore)
	} else {
		set(currentScore)
	}
}

// Play func
func (b *Board) Play(number, seed, origin string) error {
	p, err := b.Players().Find(func(p *player.Player) bool { return p.Host() == origin })
	if err == nil {
		c, _ := p.Play(number, seed)
		// c, err := p.Play(number, seed)
		// if err == nil { // TODO: FOR SOME CHECKS IT'S TRUE
		b.PlayedCards().Add(c)
		if len(*b.PlayedCards()) >= 5 {
			playerIndex := briscola.IndexOfWinningCard(*b.PlayedCards(), card.Coin)
			b.PlayedCards().Move(b.Players()[playerIndex].Pile())
		}
		// }
	}
	return err
}

// Nominate func
func (b *Board) Nominate(number, seed, origin string) (card.ID, error) {
	card, err := card.Create(number, seed)
	if err == nil {
		p, err := b.Players().Find(func(p *player.Player) bool { return p.Has(card) })
		if err == nil {
			b.selectedCard = card
			b.selectedPlayer = *p
		}
	}
	return card, err
}

// Join func
func (b *Board) Join(name, origin string) error {
	p, err := b.Players().Find(func(p *player.Player) bool { return p.Name() == "" })
	if err == nil {
		p.SetName(name)
		p.MyHostIs(origin)
	} else {
		log.Println("All players have joined, no further players are expected: " + err.Error())
		log.Println(b.Players())
	}
	return err
}
