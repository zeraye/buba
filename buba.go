package buba

import (
	"fmt"
	"time"

	"github.com/notnil/chess"
)

func BestMove(fen_string string) *chess.Move {
	fen, err := chess.FEN(fen_string)

	if err != nil {
		panic("Invalid FEN string!")
	}

	counter := 0
	pos := chess.NewGame(fen).Position()
	maxDepth := 99
	maxRuntime := 10000.0 // ms

	timerStart := float64(time.Now().UnixMilli())

	var bestMove *chess.Move
	var currMove *chess.Move
	var bestScore float64
	var currScore float64
	var depth int

	for depth = 1; depth < maxDepth; depth++ {
		currMove, currScore, err = miniMax(pos, depth, &counter, timerStart, maxRuntime)

		if err != nil {
			break
		}

		bestMove = currMove
		bestScore = currScore
	}

	fmt.Println("evaluation", bestScore)

	return bestMove
}
