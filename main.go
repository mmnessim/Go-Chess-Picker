package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"go-chess/game"
	"go-chess/user"
)

func main() {
	fmt.Println("Hello")

	PickUser()

	/*
		me := user.New("tenderllama")
		//fmt.Println(me)
		me.GetArchives()
		//fmt.Println(me.Archives)

		randomGame := game.GetRandomGame(&me)
		randomGame.Summary()
	*/
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
