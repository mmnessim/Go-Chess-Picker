package game

import (
	"encoding/json"
	"fmt"
	"go-chess/user"
	"io"
	"net/http"

	"math/rand"
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

// Returns Game struct
func GetRandomGame(u *user.ChessUser) Game {
	if u.UsernameNotFound {
		return Game{Err: true}
	}
	randomArchive := u.Archives[rand.Intn(len(u.Archives))]
	//fmt.Println(randomArchive.(string))
	resp, err := http.Get(randomArchive.(string))
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

	// for testing
	//fmt.Println(len(gameArray))
	randomGame := gameArray[rand.Intn(len(gameArray))].(map[string]interface{})
	//fmt.Println(randomGame)

	g := Game{
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
