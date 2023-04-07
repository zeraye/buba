package buba

import (
	"time"

	"github.com/notnil/chess"
	"github.com/notnil/chess/uci"
)

func stockfishEvaluate(eng *uci.Engine, position *chess.Position, moveTime time.Duration) int {
	cmdPos := uci.CmdPosition{Position: position}
	cmdGo := uci.CmdGo{MoveTime: moveTime}
	if err := eng.Run(cmdPos, cmdGo); err != nil {
		panic(err)
	}
	return eng.SearchResults().Info.Score.CP
}
