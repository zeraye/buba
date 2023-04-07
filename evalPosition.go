package buba

import (
	"math"

	"github.com/notnil/chess"
)

var (
	pieceToValue = map[chess.Piece]int{
		chess.WhitePawn: 100, chess.WhiteKnight: 320, chess.WhiteBishop: 330, chess.WhiteRook: 500, chess.WhiteQueen: 900, chess.WhiteKing: 20000,
		chess.BlackPawn: -100, chess.BlackKnight: -320, chess.BlackBishop: -330, chess.BlackRook: -500, chess.BlackQueen: -900, chess.BlackKing: -20000,
	}
)

func evalPosition(pos *chess.Position, oldPos *chess.Position) (int, error) {
	status := pos.Status()

	if status >= chess.Stalemate {
		return 0, nil
	}

	if status == chess.Checkmate {
		if pos.Turn() == chess.White {
			return math.MinInt, nil
		}
		return math.MaxInt, nil
	}

	isEndGame, err := checkEndGame(pos)
	if err != nil {
		return 0, err
	}

	eval := 0

	for square, piece := range pos.Board().SquareMap() {
		eval += squareTableEval(piece.Type(), piece.Color(), square, isEndGame)
		eval += pieceToValue[piece]
	}

	eval += 10 * (len(pos.ValidMoves()) - len(oldPos.ValidMoves())) * colorToInt[pos.Turn()]

	/*
		data, err := pos.Board().MarshalBinary()
		if err != nil {
			return 0, err
		}

		for i := 0; i < 12; i++ {
			eval += bits.OnesCount64(binary.BigEndian.Uint64(data[i*8:i*8+8])) * pieceToValue[chess.Piece(i+1)]
		}
	*/

	return eval, nil
}
