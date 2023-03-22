package main

import (
	"fmt"
	"time"

	"github.com/notnil/chess"
	"github.com/notnil/chess/uci"
	"github.com/zeraye/buba"
)

func main() {
	eng, err := uci.New("stockfish")
	if err != nil {
		panic(err)
	}
	defer eng.Close()
	// initialize uci with new game
	if err := eng.Run(uci.CmdUCI, uci.CmdIsReady, uci.CmdUCINewGame); err != nil {
		panic(err)
	}
	// have stockfish play speed chess against itself (10 msec per move)
	fen_string := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
	//fen_string = "7R/N3P3/8/8/8/8/1k5K/8 b - - 36 84"
	fen_string = "4rr2/2q1p2R/p3k1p1/2ppnb2/1p3Q2/3P1N2/PPP1BPP1/R3K3 b Q - 2 22"
	fen_string = "5rk1/ppp3pp/2n1p2q/3p1r2/6B1/2P1P2n/PPQ3PB/R4RK1 w - - 0 21"
	fen, err := chess.FEN(fen_string)
	if err != nil {
		panic(err)
	}
	game := chess.NewGame(fen)
	cmdPos := uci.CmdPosition{Position: game.Position()}
	cmdGo := uci.CmdGo{MoveTime: time.Second}
	if err := eng.Run(cmdPos, cmdGo); err != nil {
		panic(err)
	}
	var color int
	if game.Position().Turn() == chess.White {
		color = 1
	} else {
		color = -1
	}
	move := eng.SearchResults().Info.Score.CP * color
	fmt.Println(move)
	buba.BestMove(fen_string)
}
