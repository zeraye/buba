package buba

import (
	"time"

	"github.com/notnil/chess"
)

type result struct {
	move          *chess.Move
	score         int
	nodesAnalysed int
	depth         int
}

func BestMove(fen_string string, maxRuntime time.Duration) result {
	fen, err := chess.FEN(fen_string)
	if err != nil {
		panic("Invalid FEN string!")
	}

	var bestMove *chess.Move
	var currMove *chess.Move
	var bestScore int
	var currScore int
	var depth int

	maxDepth := 99
	nodesAnalysed := 0
	pos := chess.NewGame(fen).Position()
	timeEnd := time.Now().Add(maxRuntime)
	cache := make(map[[16]byte]int)

	for depth = 1; depth < maxDepth; depth++ {
		currMove, currScore, err = miniMax(pos, depth, &nodesAnalysed, timeEnd, &cache)
		if err != nil {
			break
		}

		bestMove = currMove
		bestScore = currScore
	}

	return result{bestMove, bestScore, nodesAnalysed, depth}
}
