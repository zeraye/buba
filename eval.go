package buba

import (
	"math"

	"github.com/notnil/chess"
)

var (
	pieceToValue = map[chess.Piece]float64{
		chess.WhitePawn: 100, chess.WhiteKnight: 320, chess.WhiteBishop: 330, chess.WhiteRook: 500, chess.WhiteQueen: 900, chess.WhiteKing: 20000,
		chess.BlackPawn: -100, chess.BlackKnight: -320, chess.BlackBishop: -330, chess.BlackRook: -500, chess.BlackQueen: -900, chess.BlackKing: -20000,
	}
)

func eval(pos *chess.Position) float64 {
	eval := 0.0

	switch pos.Status() {
	case chess.FivefoldRepetition:
		return 0
	case chess.Stalemate:
		return 0
	case chess.Checkmate:
		if pos.Turn() == chess.White {
			return math.Inf(-1)
		}
		return math.Inf(1)
	}

	board := pos.Board().SquareMap()
	isEndGame := checkEndGame(pos)

	for s, p := range board {
		eval += squareTableEval(p.Type(), p.Color(), s, isEndGame)
		eval += pieceToValue[p]
	}

	return eval
}
