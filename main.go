package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"

	"go-chess/game"
	"go-chess/handlers"
	"go-chess/user"
)

func main() {
	fmt.Println("Hello")

	//test that bad username doesn't crash program
	badUser := user.New("not a real user")
	fmt.Println(badUser.UsernameNotFound)
	badUser.GetArchives()

	http.HandleFunc("/", handlers.Index)
	http.HandleFunc("/game", handlers.Game)
	http.ListenAndServe(":8080", nil)
}

func PickUser() {
	fmt.Println("Enter Chess.com Username:")

	scanner := bufio.NewReader(os.Stdin)

	username, _ := scanner.ReadString('\n')
	u := user.New(strings.TrimSpace(username))
	u.GetArchives()

	for {
		fmt.Println("Get random game?y/n")
		answer, _ := scanner.ReadString('\n')
		if strings.TrimSpace(answer) == "y" {
			randomGame := game.GetRandomGame(&u)
			randomGame.Summary()
		} else if strings.TrimSpace(answer) == "n" {
			break
		} else {
			continue
		}
	}

}
