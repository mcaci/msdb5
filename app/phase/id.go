package phase

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
func ToID(phase string) ID {
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
