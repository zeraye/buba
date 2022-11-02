package buba

import (
	"errors"
	"math"
	"math/rand"
	"time"

	"github.com/notnil/chess"
)

var boolToInt = map[bool]int{
	true: 1, false: -1,
}

func miniMax(pos *chess.Position, depth int, counter *int, timerStart float64, maxRuntime float64) (*chess.Move, float64, error) {
	var bestScore float64

	cache := make(map[[16]byte]float64)
	alpha := math.Inf(-1)
	beta := math.Inf(1)
	bestMoves := []*chess.Move{}
	isMax := pos.Turn() == chess.White

	bestScore = math.Inf(boolToInt[!isMax])

	for _, move := range pos.ValidMoves() {
		new_pos := pos.Update(move)
		eval, err := miniMaxRecursive(new_pos, depth, !isMax, alpha, beta, counter, &cache, timerStart, maxRuntime)

		if err != nil {
			return nil, 0.0, err
		}

		if (eval > bestScore && isMax) || (eval < bestScore && !isMax) {
			bestMoves = []*chess.Move{move}
			bestScore = eval
		} else if eval == bestScore {
			bestMoves = append(bestMoves, move)
		}
	}

	return bestMoves[rand.Intn(len(bestMoves))], bestScore, nil
}

func miniMaxRecursive(pos *chess.Position, depth int, isMax bool, alpha float64, beta float64, counter *int, cache *map[[16]byte]float64, timerStart float64, maxRuntime float64) (float64, error) {
	if float64(time.Now().UnixMilli())-timerStart > maxRuntime {
		return 0.0, errors.New("runtime excedeed")
	}

	(*counter)++

	if pos.Status() != chess.NoMethod || depth == 0 {
		return eval(pos), nil
	}

	cached_eval, ok := (*cache)[pos.Hash()]

	if ok {
		return cached_eval, nil
	}

	(*cache)[pos.Hash()] = eval(pos)

	bestScore := math.Inf(boolToInt[!isMax])

	for _, move := range pos.ValidMoves() {
		new_pos := pos.Update(move)

		eval, err := miniMaxRecursive(new_pos, depth-1, !isMax, alpha, beta, counter, cache, timerStart, maxRuntime)

		if err != nil {
			return 0.0, err
		}

		if isMax {
			bestScore = math.Max(eval, bestScore)
			alpha = math.Max(bestScore, alpha)
		} else {
			bestScore = math.Min(eval, bestScore)
			beta = math.Min(bestScore, beta)
		}

		if beta <= alpha {
			break
		}
	}

	return bestScore, nil
}
