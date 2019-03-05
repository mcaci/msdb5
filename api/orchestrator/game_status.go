package orchestrator

type phase uint8

const (
	joining phase = iota
	scoreAuction
	companionChoice
)
