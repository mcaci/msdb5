package team

import "github.com/mcaci/msdb5/dom/player"

// Part func
func Part(pls Players, pred player.Predicate) (t1, t2 Players) {
	for _, p := range pls {
		if pred(p) {
			t1.Add(p)
			continue
		}
		t2.Add(p)
	}
	return
}
