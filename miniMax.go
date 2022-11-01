package buba

import (
	"math"
	"math/rand"

	"github.com/notnil/chess"
)

var boolToInt = map[bool]int{
	true: 1, false: -1,
}

func miniMax(pos *chess.Position, depth int, counter *int) (*chess.Move, float64) {
	var bestScore float64

	cache := make(map[[16]byte]float64)
	alpha := math.Inf(-1)
	beta := math.Inf(1)
	bestMoves := []*chess.Move{}
	isMax := pos.Turn() == chess.White

	bestScore = math.Inf(boolToInt[!isMax])

	for _, move := range pos.ValidMoves() {
		new_pos := pos.Update(move)
		eval := miniMaxRecursive(new_pos, depth, !isMax, alpha, beta, counter, &cache)

		if (eval > bestScore && isMax) || (eval < bestScore && !isMax) {
			bestMoves = []*chess.Move{move}
			bestScore = eval
		} else if eval == bestScore {
			bestMoves = append(bestMoves, move)
		}
	}

	return bestMoves[rand.Intn(len(bestMoves))], bestScore
}

func miniMaxRecursive(pos *chess.Position, depth int, isMax bool, alpha float64, beta float64, counter *int, cache *map[[16]byte]float64) float64 {
	(*counter)++

	if pos.Status() != chess.NoMethod || depth == 0 {
		return eval(pos)
	}

	cached_eval, ok := (*cache)[pos.Hash()]

	if ok {
		return cached_eval
	}

	(*cache)[pos.Hash()] = eval(pos)

	bestScore := math.Inf(boolToInt[!isMax])

	for _, move := range pos.ValidMoves() {
		new_pos := pos.Update(move)

		if isMax {
			bestScore = math.Max(miniMaxRecursive(new_pos, depth-1, !isMax, alpha, beta, counter, cache), bestScore)
			alpha = math.Max(bestScore, alpha)
		} else {
			bestScore = math.Min(miniMaxRecursive(new_pos, depth-1, !isMax, alpha, beta, counter, cache), bestScore)
			beta = math.Min(bestScore, beta)
		}

		if beta <= alpha {
			break
		}
	}

	return bestScore
}
