package game

import (
	"fmt"
)

type Game struct {
	ID          int
	Url         string
	Pgn         string // Maybe needs to be map[string]interface{}
	TimeControl string
	Black       Black
	White       White
	Err         bool
}

type Black struct {
	Username string
	Rating   float64
	Result   string
}

type White struct {
	Username string
	Rating   float64
	Result   string
}

func (g *Game) Summary() {
	fmt.Println("Black:", g.Black, "White", g.White)
	fmt.Println("Go Analyze:", g.Url)
}
