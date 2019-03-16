package orchestrator

// Phase type
type Phase uint8

const (
	Joining Phase = iota
	InsideAuction
	ChosingCompanion
	PlayingCards
	End
)
