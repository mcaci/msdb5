package game

const numberOfPlayers = 5

func roundRobin(idx, off, size uint8) uint8 {
	return (idx + off) % size
}
