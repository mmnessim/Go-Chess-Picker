package randomuser

import (
	"encoding/json"
	"fmt"
	"go-chess/user"
	"io"
	"math/rand"
	"net/http"
)

type UserList struct {
	Players []string `json:"players"`
}

func GetAllUsers() UserList {
	resp, err := http.Get("https://api.chess.com/pub/country/US/players")
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	users := UserList{}
	json.Unmarshal(body, &users)

	return users
}

func (ul *UserList) GetRandomUser() user.ChessUser {
	len := len(ul.Players)
	index := rand.Intn(len)
	username := ul.Players[index]
	user := user.New(username)
	user.GetRandomGame()
	if user.Game.Err {
		ul.Players = append(ul.Players[:index], ul.Players[index+1:]...)
		ul.GetRandomUser()
	}
	return user
}
