package main

import (
	"fmt"

	"github.com/notnil/chess"
	"github.com/zeraye/buba"
)

func main() {
	game := chess.NewGame()

	for game.Outcome() == chess.NoOutcome {
		fmt.Println(game.Position().Board().Draw())
		if game.Position().Turn() == chess.White {
			move := buba.BestMove(game.Position().String())
			game.Move(move)
		} else {
			for {
				var move string
				fmt.Scan(&move)
				err := game.MoveStr(move)

				if err == nil {
					break
				} else {
					fmt.Println("Wrong move, try again!")
				}
			}
		}

	}
}
