package buba

import (
	"errors"
	"math"
	"time"

	"github.com/notnil/chess"
)

var cache *map[[16]byte]int
var timeEnd time.Time
var analysedNodes *int

func miniMax(pos *chess.Position, depth int, mmAnalysedNodes *int, mmTimeEnd time.Time, mmCache *map[[16]byte]int) (*chess.Move, int, error) {
	alpha := math.MinInt
	beta := math.MaxInt
	bestMoves := []*chess.Move{}
	isMax := pos.Turn() == chess.White
	bestScore := math.MaxInt
	if isMax {
		bestScore = math.MinInt
	}
	cache = mmCache
	timeEnd = mmTimeEnd
	analysedNodes = mmAnalysedNodes

	for _, move := range rankMoves(pos.ValidMoves(), pos) {
		newPos := pos.Update(move)

		eval, err := miniMaxRecursive(newPos, pos, depth, !isMax, alpha, beta)
		if err != nil {
			return nil, 0, err
		}

		if (eval > bestScore && isMax) || (eval < bestScore && !isMax) {
			bestMoves = []*chess.Move{move}
			bestScore = eval
		} else if eval == bestScore {
			bestMoves = append(bestMoves, move)
		}
	}

	return bestMoves[0], bestScore, nil
}

func miniMaxRecursive(pos *chess.Position, oldPos *chess.Position, depth int, isMax bool, alpha int, beta int) (int, error) {
	if timeEnd.Compare(time.Now()) == -1 {
		return 0, errors.New("searching time exceeded")
	}

	posEval, evalFound := (*cache)[pos.Hash()]

	if !evalFound {
		posEval, _ = evalPosition(pos, oldPos)
	}

	(*analysedNodes)++

	if pos.Status() != chess.NoMethod || depth == 0 {
		return posEval, nil
	}

	(*cache)[pos.Hash()] = posEval

	bestScore := math.MaxInt
	if isMax {
		bestScore = math.MinInt
	}

	for _, move := range rankMoves(pos.ValidMoves(), pos) {
		newPos := pos.Update(move)

		eval, err := miniMaxRecursive(newPos, pos, depth-1, !isMax, alpha, beta)
		if err != nil {
			return 0, err
		}

		if isMax {
			if eval > bestScore {
				bestScore = eval
			}
			if bestScore > alpha {
				alpha = bestScore
			}
		} else {
			if eval < bestScore {
				bestScore = eval
			}
			if bestScore < beta {
				beta = bestScore
			}
		}

		if beta <= alpha {
			break
		}
	}

	return bestScore, nil
}
