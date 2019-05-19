package team

import (
	"github.com/nikiforosFreespirit/msdb5/dom/player"
)

// CountFolded func
func CountFolded(players ...player.FoldedChecker) (count uint8) {
	for _, p := range players {
		testAndIncrement(&count, p.Folded)
	}
	return
}

// CountEmptyNames func
func CountEmptyNames(players ...player.EmptyNameChecker) (count uint8) {
	for _, p := range players {
		testAndIncrement(&count, p.IsNameEmpty)
	}
	return
}

// CountEmptyHands func
func CountEmptyHands(players ...player.EmptyHandChecker) (count uint8) {
	for _, p := range players {
		testAndIncrement(&count, p.IsHandEmpty)
	}
	return
}

func testAndIncrement(count *uint8, predicate func() bool) {
	if predicate() {
		*count++
	}
}
