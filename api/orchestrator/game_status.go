package orchestrator

type status uint8

const (
	joining status = iota
	scoreAuction
)
