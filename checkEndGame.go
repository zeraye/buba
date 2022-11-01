package buba

import (
	"math/bits"

	"github.com/notnil/chess"
)

func checkEndGame(pos *chess.Position) bool {
	data, _ := pos.Board().MarshalBinary()
	pieces_count := 0

	for entry := range data {
		pieces_count += bits.OnesCount(uint(entry))
	}

	return pieces_count < 16
}
