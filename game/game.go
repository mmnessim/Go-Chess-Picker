package game

import (
	"errors"
	"fmt"
)

type Game struct {
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

type GameNode struct {
	Game  Game
	Index int
	Next  *GameNode
}

type GameList struct {
	Head   *GameNode
	Length int
}

func (gl *GameList) InsertAtHead(game Game) {
	temp1 := &GameNode{game, gl.Length, nil}

	if gl.Head == nil {
		gl.Head = temp1
	} else {
		temp2 := gl.Head
		gl.Head = temp1
		temp1.Next = temp2
	}
	gl.Length += 1
}

func (gl *GameList) GetFromIndex(index int) (*Game, error) {
	cur := gl.Head
	for {
		if cur.Index == index {
			return &cur.Game, nil
		} else if cur.Next == nil {
			return nil, errors.New("index outside of bounds")
		}
		cur = cur.Next
	}
}
