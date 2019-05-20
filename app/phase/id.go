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
