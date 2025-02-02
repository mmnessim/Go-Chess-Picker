package user

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"

	"go-chess/game"
)

type ChessUser struct {
	Username         string
	Verified         bool
	Url              string
	ApiUrl           string
	Archives         []string
	Info             map[string]interface{}
	Game             game.Game
	UsernameNotFound bool
}

func New(username string) ChessUser {
	c := ChessUser{Username: username}
	c.init()
	return c
}

func (c *ChessUser) init() {
	apiUrl := "https://api.chess.com/pub/player/" + c.Username
	resp, err := http.Get(apiUrl)
	if err != nil {
		fmt.Println(err)
		c.UsernameNotFound = true
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	bodyMap := make(map[string]interface{})
	json.Unmarshal(body, &bodyMap)

	if len(bodyMap) == 0 || bodyMap["message"] != nil {
		c.UsernameNotFound = true
		return
	}

	c.Info = bodyMap
	ver, ok := bodyMap["verified"].(bool)
	if ok {
		c.Verified = ver
	}

	url, ok := bodyMap["url"].(string)
	if ok {
		c.Url = url
	}

	c.ApiUrl = apiUrl

	c.GetArchives()
}

func (c *ChessUser) GetArchives() {
	if c.UsernameNotFound {
		return
	}
	resp, err := http.Get(c.ApiUrl + "/games/archives")
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	archives := make(map[string]interface{})
	json.Unmarshal(body, &archives)

	tempArch := archives["archives"].([]interface{})

	for _, a := range tempArch {
		c.Archives = append(c.Archives, a.(string))
	}
}

func (u *ChessUser) GetRandomGame() {

	// Handle invalid users or users with no games
	if u.UsernameNotFound || len(u.Archives) == 0 {
		u.Game = game.Game{Err: true}
		return
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

	u.Game = game.Game{
		Url:         randomGame["url"].(string),
		Pgn:         randomGame["pgn"].(string),
		TimeControl: randomGame["time_control"].(string),
		Black: game.Black{
			Username: randomGame["black"].(map[string]interface{})["username"].(string),
			Rating:   randomGame["black"].(map[string]interface{})["rating"].(float64),
			Result:   randomGame["black"].(map[string]interface{})["result"].(string),
		},
		White: game.White{
			Username: randomGame["white"].(map[string]interface{})["username"].(string),
			Rating:   randomGame["white"].(map[string]interface{})["rating"].(float64),
			Result:   randomGame["white"].(map[string]interface{})["result"].(string),
		},
	}
}
