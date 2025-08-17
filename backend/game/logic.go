package game

// Reset sets the game to a new initial state.
func (g *Game) Reset() {
	*g = *NewGame()
	g.Winner = ""
}

// NewGame initializes a new game with an empty board and starting player 'X'.
func NewGame() *Game {
	return &Game{
		Board:         [3][3]string{},
		CurrentPlayer: "X",
		Status:        InProgress,
	}
}

// MakeMove attempts to place the current player's mark at the given row and col.
// Returns true if the move is valid and updates the game state.
func (g *Game) MakeMove(row, col int) bool {
	if row < 0 || row > 2 || col < 0 || col > 2 {
		return false
	}
	if g.Board[row][col] != "" || g.Status != InProgress {
		return false
	}
	g.Board[row][col] = g.CurrentPlayer
	if winLine := g.getWinningLine(g.CurrentPlayer); winLine != nil {
		g.Status = Win
		g.WinningLine = winLine
		g.Winner = g.CurrentPlayer
	} else if g.checkDraw() {
		g.Status = Draw
		g.WinningLine = nil
		g.Winner = ""
	} else {
		g.switchPlayer()
		g.WinningLine = nil
		g.Winner = ""
	}
	return true
}

// getWinningLine returns the winning line as a slice of [row, col] if player has won, else nil
func (g *Game) getWinningLine(player string) [][2]int {
	b := g.Board
	// Rows
	for i := 0; i < 3; i++ {
		if b[i][0] == player && b[i][1] == player && b[i][2] == player {
			return [][2]int{{i, 0}, {i, 1}, {i, 2}}
		}
	}
	// Columns
	for i := 0; i < 3; i++ {
		if b[0][i] == player && b[1][i] == player && b[2][i] == player {
			return [][2]int{{0, i}, {1, i}, {2, i}}
		}
	}
	// Diagonal TL-BR
	if b[0][0] == player && b[1][1] == player && b[2][2] == player {
		return [][2]int{{0, 0}, {1, 1}, {2, 2}}
	}
	// Diagonal TR-BL
	if b[0][2] == player && b[1][1] == player && b[2][0] == player {
		return [][2]int{{0, 2}, {1, 1}, {2, 0}}
	}
	return nil
}

func (g *Game) switchPlayer() {
	if g.CurrentPlayer == "X" {
		g.CurrentPlayer = "O"
	} else {
		g.CurrentPlayer = "X"
	}
}

func (g *Game) checkWin(player string) bool {
	b := g.Board
	for i := 0; i < 3; i++ {
		if b[i][0] == player && b[i][1] == player && b[i][2] == player {
			return true
		}
		if b[0][i] == player && b[1][i] == player && b[2][i] == player {
			return true
		}
	}
	if b[0][0] == player && b[1][1] == player && b[2][2] == player {
		return true
	}
	if b[0][2] == player && b[1][1] == player && b[2][0] == player {
		return true
	}
	return false
}

func (g *Game) checkDraw() bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if g.Board[i][j] == "" {
				return false
			}
		}
	}
	return true
}
