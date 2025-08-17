package game

type Status string

const (
	InProgress Status = "in-progress"
	Win        Status = "win"
	Draw       Status = "draw"
)

type Game struct {
	Board         [3][3]string `json:"board"`
	CurrentPlayer string       `json:"current_player"`
	Status        Status       `json:"status"`
	WinningLine   [][2]int     `json:"winning_line,omitempty"`
	Winner        string       `json:"winner,omitempty"`
}
