package buba

import "testing"

type bestMoveTest struct {
	fen      string
	bestMove string
}

var bestMoveTests = []bestMoveTest{
	{"5rk1/ppp3pp/2n1p1q1/3p1r2/6B1/2P1P2P/PPQ4B/R4RK1 w - - 1 22", "c2g2"},
	{"5rk1/ppp3pp/2n1p2q/3p1r2/6B1/2P1P2n/PPQ3PB/R4RK1 w - - 0 21", "g2h3"},
	{"4rr2/2q1p2R/p3k1p1/2ppnb2/1p3Q2/3P1N2/PPP1BPP1/R3K3 b Q - 2 22", "e5d3"},
}

func TestBestMove(t *testing.T) {
	for _, test := range bestMoveTests {
		if bestMove := BestMove(test.fen).String(); bestMove != test.bestMove {
			t.Errorf("Bad move (%s) for position %s", bestMove, test.fen)
		}
	}
}
