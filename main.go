package main

import (
	"fmt"

	"go-chess/game"
	"go-chess/user"
)

func main() {
	fmt.Println("Hello")

	me := user.New("tenderllama")
	//fmt.Println(me)
	me.GetArchives()
	//fmt.Println(me.Archives)

	randomGame := game.GetRandomGame(&me)
	randomGame.Summary()
}
