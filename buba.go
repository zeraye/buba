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

	timer_end := float64(time.Now().UnixMilli())

	timer_diff := timer_end - timerStart
	timer_diff_sec := timer_diff / float64(time.Second.Milliseconds())

	fmt.Println("summary")
	fmt.Println("runtime[s]", timer_diff_sec)
	fmt.Println("nodes     ", counter)
	fmt.Println("k-nodes/s ", int(float64(counter)/timer_diff))
	fmt.Println("evaluation", bestScore/100) // change evaluation scale from pawn=100 to pawn=1
	fmt.Println("depth     ", depth)
	fmt.Println("best_move ", bestMove.String())

	return bestMove
}
