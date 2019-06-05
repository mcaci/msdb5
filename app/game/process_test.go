package game

// func TestCompletedGameReturningScoreInfo(t *testing.T) {
// 	gameTest := NewGame(false)
// 	gameTest.Join("127.0.0.51", make(chan []byte))
// 	gameTest.Join("127.0.0.52", make(chan []byte))
// 	gameTest.Join("127.0.0.53", make(chan []byte))
// 	gameTest.Join("127.0.0.54", make(chan []byte))
// 	gameTest.Join("127.0.0.55", make(chan []byte))
// 	for i, pl := range gameTest.players {
// 		pl = player.New()
// 		pl.Draw(func() card.ID { return card.ID(2*i + 5) })
// 		if i > 0 {
// 			pl.Fold()
// 		}
// 	}
// 	gameTest.Process("Join#A", "127.0.0.51")
// 	gameTest.Process("Join#B", "127.0.0.52")
// 	gameTest.Process("Join#C", "127.0.0.53")
// 	gameTest.Process("Join#D", "127.0.0.54")
// 	gameTest.Process("Join#E", "127.0.0.55")
// 	gameTest.Process("Auction#80", "127.0.0.51")
// 	gameTest.Process("Companion#7#Coin", "127.0.0.51")
// 	gameTest.Process("Card#5#Coin", "127.0.0.51")
// 	gameTest.Process("Card#7#Coin", "127.0.0.52")
// 	gameTest.Process("Card#9#Coin", "127.0.0.53")
// 	gameTest.Process("Card#1#Cup", "127.0.0.54")
// 	gameTest.Process("Card#3#Cup", "127.0.0.55")
// 	if false {
// 		t.Fatal("Expecting transition to end game and scoring")
// 	}
// }
