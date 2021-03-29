package briscola5

import (
	"errors"
	"log"

	"github.com/mcaci/msdb5/v2/dom/team"
)

// Players is a slice of pointers to players
type Players []*Player

func ToGeneralPlayers(bp5 Players) team.Players {
	pls := make(team.Players, 5)
	for i := range bp5 {
		pls[i] = &bp5[i].Player
	}
	return pls
}

// Add adds a player to the team
func (playerSet *Players) Add(p *Player) {
	*playerSet = append(*playerSet, p)
}

// ErrPlayerNotFound error
var ErrPlayerNotFound = errors.New("Player not found")

// Index returns the index of the player that satisfies the predicate
// or an error if not found
func (playerSet Players) Index(predicate Predicate) (uint8, error) {
	for i, p := range playerSet {
		if predicate(p) {
			return uint8(i), nil
		}
	}
	return 0, ErrPlayerNotFound
}

// MustIndex works as Index but exits fatally in case of error
func (playerSet Players) MustIndex(predicate Predicate) uint8 {
	i, err := playerSet.Index(predicate)
	if err != nil {
		log.Fatalf("error found: %v", err)
	}
	return i
}

// Count counts the number of players satisfying the predicate
func Count(players Players, predicate Predicate) (count uint8) {
	for _, p := range players {
		if predicate(p) {
			count++
		}
	}
	return
}
