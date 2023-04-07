package buba

import (
	"sort"

	"github.com/notnil/chess"
)

type chessMoves []*chess.Move
type moveWithEval struct {
	moves *chess.Move
	eval  int
}
type movesWithEval []moveWithEval

func (moves chessMoves) Len() int { return len(moves) }
func (moves chessMoves) Less(i, j int) bool {
	for _, tag := range []chess.MoveTag{
		chess.Capture,
		chess.EnPassant,
		chess.KingSideCastle,
		chess.QueenSideCastle,
		chess.Check,
	} {
		if !moves[i].HasTag(tag) && moves[j].HasTag((tag)) {
			return true
		}
		if moves[i].HasTag(tag) && !moves[j].HasTag((tag)) {
			return false
		}
	}

	if moves[i].Promo() == chess.NoPieceType && moves[j].Promo() != chess.NoPieceType {
		return true
	}
	if moves[i].Promo() != chess.NoPieceType && moves[j].Promo() == chess.NoPieceType {
		return false
	}
	if moves[i].Promo() != chess.NoPieceType && moves[j].Promo() != chess.NoPieceType {
		return moves[i].Promo() > moves[j].Promo()
	}

	return false
}
func (moves chessMoves) Swap(i, j int) { moves[i], moves[j] = moves[j], moves[i] }

func (moves movesWithEval) Len() int { return len(moves) }
func (moves movesWithEval) Less(i, j int) bool {
	return moves[i].eval < moves[j].eval
}
func (moves movesWithEval) Swap(i, j int) {
	moves[i], moves[j] = moves[j], moves[i]
}

func rankMoves(topMoves []*chess.Move, pos *chess.Position) []*chess.Move {
	sort.Sort(sort.Reverse(chessMoves(topMoves)))
	topMovesCount := 5
	if len(topMoves) < topMovesCount {
		topMovesCount = len(topMoves)
	}
	toptopMoves := topMoves[:topMovesCount]
	toptopMovesWithEval := []moveWithEval{}

	for _, toptopMove := range toptopMoves {
		newPos := pos.Update(toptopMove)
		eval, _ := evalPosition(newPos, pos)
		toptopMovesWithEval = append(toptopMovesWithEval, moveWithEval{toptopMove, eval})
	}

	sort.Sort(sort.Reverse(movesWithEval(toptopMovesWithEval)))
	toptopMovesCount := 3
	if len(toptopMoves) < toptopMovesCount {
		toptopMovesCount = len(toptopMoves)
	}
	return toptopMoves[:toptopMovesCount]
}
