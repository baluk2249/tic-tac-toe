package main

type GameStatus string

const (
	InProgress GameStatus = "in-progress"
	Win        GameStatus = "win"
	Draw       GameStatus = "draw"
)

type Game struct {
	Board         [3][3]string `json:"board"`
	CurrentPlayer string       `json:"current_player"`
	Status        GameStatus   `json:"status"`
}
