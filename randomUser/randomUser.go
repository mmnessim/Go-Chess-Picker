package randomuser

import (
	"encoding/json"
	"fmt"
	"go-chess/db"
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

func PopulateAllUsers() {
	allUsers := db.UsersInit()
	defer allUsers.Close()

	usersDataFromChessCom := GetAllUsers()

	for _, u := range usersDataFromChessCom.Players {
		db.AddUser(u, allUsers)
		fmt.Printf("Added user %s\n", u)
	}
}

func GetRandomUser() user.ChessUser {
	allUsers := db.UsersInit()
	defer allUsers.Close()

	type record struct {
		Id       int
		Username string
	}

	query := `SELECT * FROM users ORDER BY RANDOM() LIMIT 1;`
	row, err := allUsers.Query(query)
	if err != nil {
		fmt.Println(err)
	}

	var result record = record{}

	for row.Next() {
		err = row.Scan(&result.Id, &result.Username)
		if err != nil {
			fmt.Println(err)
		}
	}

	user := user.New(result.Username)
	user.GetRandomGame()
	if user.Game.Err {
		fmt.Println("User has no games, removing...")
		db.RemoveUser(result.Username, allUsers)
		GetRandomUser()
	}
	return user
}
