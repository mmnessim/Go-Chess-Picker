package game

import (
	"encoding/json"
	"errors"
	"fmt"
	"go-chess/user"
	"io"
	"net/http"

	"math/rand"
)

type Game struct {
	User        user.ChessUser
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

// Returns Game struct
func GetRandomGame(u *user.ChessUser) Game {

	// Handle invalid users or users with no games
	if u.UsernameNotFound || len(u.Archives) == 0 {
		return Game{Err: true}
	}
	randomArchive := u.Archives[rand.Intn(len(u.Archives))]
	resp, err := http.Get(randomArchive)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	games := make(map[string]interface{})
	json.Unmarshal(body, &games)

	gameArray := games["games"].([]interface{})

	randomGame := gameArray[rand.Intn(len(gameArray))].(map[string]interface{})

	g := Game{
		User:        *u,
		Url:         randomGame["url"].(string),
		Pgn:         randomGame["pgn"].(string),
		TimeControl: randomGame["time_control"].(string),
		Black: Black{
			Username: randomGame["black"].(map[string]interface{})["username"].(string),
			Rating:   randomGame["black"].(map[string]interface{})["rating"].(float64),
			Result:   randomGame["black"].(map[string]interface{})["result"].(string),
		},
		White: White{
			Username: randomGame["white"].(map[string]interface{})["username"].(string),
			Rating:   randomGame["white"].(map[string]interface{})["rating"].(float64),
			Result:   randomGame["white"].(map[string]interface{})["result"].(string),
		},
	}
	return g
}
