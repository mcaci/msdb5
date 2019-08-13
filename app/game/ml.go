package game

func (g *Game) handleMLData() {
	// log action to file for ml (TODO: WHEN PUSHED OUTSIDE FUNC -> PROBLEM)
	// f, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	// if err != nil {
	// 	cons, consMsg := os.Stdout, err.Error()
	// 	pr.reports = append(pr.reports, PlMsg{cons, consMsg})
	// 	return pr.reports
	// }
	// defer f.Close()
	// // write to file for ml
	// switch g.Phase() {
	// case phase.ChoosingCompanion:
	// 	ml, mlMsg := f, fmt.Sprintf("%s, %s, %d\n", g.CurrentPlayer().Name(), g.Companion().Name(), *(g.AuctionScore()))
	// 	pr.reports = append(pr.reports, PlMsg{ml, mlMsg})
	// case phase.PlayingCards:
	// 	lastPlayed := g.playedCards[len(g.playedCards)-1]
	// 	ml, mlMsg := f, fmt.Sprintf("%s, %d\n", g.CurrentPlayer().Name(), lastPlayed)
	// 	pr.reports = append(pr.reports, PlMsg{ml, mlMsg})
	// }
	// write to file who took all cards at last round
	// ml, mlMsg := f, fmt.Sprintf("%s\n", g.CurrentPlayer().Name())
	// pr.reports = append(pr.reports, PlMsg{ml, mlMsg})
}
