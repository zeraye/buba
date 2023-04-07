package buba

import (
	"fmt"
	"time"

	"github.com/notnil/chess"
	"github.com/notnil/chess/uci"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func CalculateCentipawnLoss(fen_string string) int {
	eng, err := uci.New("bin/stockfish-windows-2022-x86-64-avx2")
	if err != nil {
		panic(err)
	}

	defer eng.Close()

	if err := eng.Run(uci.CmdUCI, uci.CmdIsReady, uci.CmdUCINewGame); err != nil {
		panic(err)
	}

	fen, err := chess.FEN(fen_string)
	if err != nil {
		panic(err)
	}

	moveTime := time.Second * 3
	game := chess.NewGame(fen)

	cmdPos := uci.CmdPosition{Position: game.Position()}
	cmdGo := uci.CmdGo{MoveTime: moveTime}
	if err := eng.Run(cmdPos, cmdGo); err != nil {
		panic(err)
	}

	bestMoveStockfish := eng.SearchResults().BestMove
	bestMoveBuba := BestMove(fen_string, time.Second*60)

	game.Move(bestMoveStockfish)
	endEvalStockfish := stockfishEvaluate(eng, game.Position(), moveTime)

	game = chess.NewGame(fen)
	game.Move(bestMoveBuba.move)
	endEvalBuba := stockfishEvaluate(eng, game.Position(), moveTime)
	cpl := abs(endEvalStockfish - endEvalBuba)
	fmt.Printf("DEBUG INFO\ndepth:\t\t%d\nnodes analysed:\t%d\nmove:\t\t%s\ncpl:\t\t%d\n", bestMoveBuba.depth, bestMoveBuba.nodesAnalysed, bestMoveBuba.move.String(), cpl)

	return cpl
}
