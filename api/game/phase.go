package game

// Phase type
type Phase uint8

const (
	Joining Phase = iota
	InsideAuction
	ExchangingCards
	ChosingCompanion
	PlayingCards
	End
)
