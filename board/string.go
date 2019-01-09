package board

import "strconv"

func (board Board) String() string {
	var str string
	str += "Board("
	str += "Players[" + board.players.String() + "]"
	str += "PlayedCards[" + board.playedCards.String() + "]"
	str += "SelectedCard[" + board.selectedCard.String() + "]"
	str += "AuctionScore[" + strconv.Itoa(int(board.auctionScore)) + "]"
	str += ")"
	return str
}
