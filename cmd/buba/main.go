package main

import (
	"fmt"
	"math"

	"github.com/zeraye/buba"
)

func main() {

	cplSum := 0
	cplCount := 0

	for _, fen_string := range []string{
		/*
			"R7/B1bb4/3k4/8/1P2p1BK/8/pPp2pp1/Q3R3 w - - 0 1",
			"5rk1/ppp3pp/2n1p2q/3p1r2/6B1/2P1P2n/PPQ3PB/R4RK1 w - - 0 21",
			"3b4/5P2/1n3Bp1/1P5p/1p1r4/1P1kp1K1/1R4pP/8 w - - 0 1",
			"1QNN4/B2k4/1Pb2p2/8/6P1/1r1p1pK1/B6p/R7 w - - 0 1",
			"3R4/4P1K1/2R1Pr2/1p1p2p1/NppN3P/8/1p6/5k2 w - - 0 1",
			"8/1r2nKp1/2P2P1P/2k1qN2/3p4/2p5/1pn4B/Q7 w - - 0 1",
		*/
		"3r1r1k/ppp3pp/1b6/6q1/5p2/PQP5/3B1PBP/3R1R1K b - - 0 1",
	} {

		cpl := buba.CalculateCentipawnLoss(fen_string)
		cplSum += cpl
		cplCount++
	}

	avgCpl := cplSum / cplCount
	fmt.Printf("Average centipawn loss: %d\n", avgCpl)
	elo := 3100 * math.Exp(-0.01*float64(avgCpl))
	fmt.Printf("Elo: %d\n", int(elo))
}
