package game

// func TestCompletedGameReturningScoreInfoWithSide(t *testing.T) {
// 	gameTest := NewGame(true)
// 	gameTest.Join("127.0.0.51", make(chan []byte))
// 	gameTest.Join("127.0.0.52", make(chan []byte))
// 	gameTest.Join("127.0.0.53", make(chan []byte))
// 	gameTest.Join("127.0.0.54", make(chan []byte))
// 	gameTest.Join("127.0.0.55", make(chan []byte))
// 	gameTest.side.Clear()
// 	gameTest.side.Add(card.ID(31))
// 	for i, pl := range gameTest.players {
// 		pl.Hand().Clear()
// 		pl.Hand().Add(card.ID(2*i + 5))
// 		if i > 1 {
// 			pl.Fold()
// 		}
// 	}
// 	gameTest.Process("Join#A", "127.0.0.51")
// 	gameTest.Process("Join#B", "127.0.0.52")
// 	gameTest.Process("Join#C", "127.0.0.53")
// 	gameTest.Process("Join#D", "127.0.0.54")
// 	gameTest.Process("Join#E", "127.0.0.55")
// 	gameTest.Process("Auction#80", "127.0.0.51")
// 	gameTest.Process("Auction#79", "127.0.0.52")
// 	gameTest.Process("Exchange#5#Coin", "127.0.0.51")
// 	gameTest.Process("Exchange#0#Coin", "127.0.0.51")
// 	gameTest.Process("Companion#7#Coin", "127.0.0.51")
// 	gameTest.Process("Card#1#Cudgel", "127.0.0.51")
// 	gameTest.Process("Card#7#Coin", "127.0.0.52")
// 	gameTest.Process("Card#9#Coin", "127.0.0.53")
// 	gameTest.Process("Card#1#Cup", "127.0.0.54")
// 	gameTest.Process("Card#3#Cup", "127.0.0.55")
// 	if false {
// 		t.Fatal("Expecting transition to end game and scoring")
// 	}
// }
