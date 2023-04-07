package buba

import (
	"math/bits"

	"github.com/notnil/chess"
)

func checkEndGame(pos *chess.Position) (bool, error) {
	data, err := pos.Board().MarshalBinary()
	if err != nil {
		return false, err
	}

	piecesCount := 0

	for row := range data {
		piecesCount += bits.OnesCount(uint(row))
	}

	return piecesCount < 12, nil
}
