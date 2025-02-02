package main

import (
	"fmt"
	"net/http"

	"go-chess/game"
	"go-chess/handlers"
	"go-chess/middleware"
	"go-chess/user"
)

func main() {
	index := middleware.Logging(handlers.Index)
	showGame := middleware.Logging(handlers.Game)

	gl := game.GameList{Head: nil, Length: 0}
	u := user.New("tenderllama")
	gl.InsertAtHead(game.GetRandomGame(&u))
	gl.InsertAtHead(game.GetRandomGame(&u))
	gl.InsertAtHead(game.GetRandomGame(&u))

	g, err := gl.GetFromIndex(3)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(g.Pgn)
	}

	fmt.Println("Listening on port 8080")

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", index)
	http.HandleFunc("/game", showGame)
	http.ListenAndServe(":8080", nil)
}
