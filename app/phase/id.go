package phase

import "strings"

// ID type
type ID uint8

const (
	Joining ID = iota
	InsideAuction
	ExchangingCards
	ChosingCompanion
	PlayingCards
	End
)

// ToID func
func ToID(request string) ID {
	phase := strings.Split(request, "#")[0]
	switch phase {
	case "Join":
		return Joining
	case "Auction":
		return InsideAuction
	case "Exchange":
		return ExchangingCards
	case "Companion":
		return ChosingCompanion
	case "Card":
		return PlayingCards
	default:
		return End
	}
}
