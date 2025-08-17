package game

import (
	"testing"
)

func TestMakeMoveAndWin(t *testing.T) {
	g := NewGame()
	moves := [][2]int{{0, 0}, {1, 0}, {0, 1}, {1, 1}, {0, 2}} // X wins
	for _, m := range moves {
		ok := g.MakeMove(m[0], m[1])
		if !ok {
			t.Fatalf("Move to %v should be valid", m)
		}
	}
	if g.Status != Win {
		t.Errorf("Expected status Win, got %v", g.Status)
	}
	if g.Winner != "X" {
		t.Errorf("Expected winner X, got %v", g.Winner)
	}
	if len(g.WinningLine) != 3 {
		t.Errorf("Expected winning line of length 3, got %d", len(g.WinningLine))
	}
}

func TestDraw(t *testing.T) {
	g := NewGame()
	moves := [][2]int{{0, 0}, {0, 1}, {0, 2}, {1, 1}, {1, 0}, {1, 2}, {2, 1}, {2, 0}, {2, 2}} // Draw
	for _, m := range moves {
		ok := g.MakeMove(m[0], m[1])
		if !ok {
			t.Fatalf("Move to %v should be valid", m)
		}
	}
	if g.Status != Draw {
		t.Errorf("Expected status Draw, got %v", g.Status)
	}
	if g.Winner != "" {
		t.Errorf("Expected no winner, got %v", g.Winner)
	}
}

func TestInvalidMove(t *testing.T) {
	g := NewGame()
	ok := g.MakeMove(0, 0)
	if !ok {
		t.Fatal("First move should be valid")
	}
	ok = g.MakeMove(0, 0)
	if ok {
		t.Error("Should not allow move to occupied cell")
	}
	ok = g.MakeMove(3, 3)
	if ok {
		t.Error("Should not allow move out of bounds")
	}
}
