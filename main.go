package main

import (
	"fmt"
	"net/http"
	"os"

	"go-chess/handlers"
	"go-chess/middleware"
	"go-chess/user"
)

func main() {
	os.Remove("./chess.db")

	randomUser, err := user.GetRandomUser()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(randomUser.Username)

	index := middleware.Logging(handlers.Index)
	showGame := middleware.Logging(handlers.Game)
	showHistory := middleware.Logging(handlers.History)

	fmt.Println("Listening on port 8080")

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", index)
	http.HandleFunc("/game", showGame)
	http.HandleFunc("/history", showHistory)
	http.ListenAndServe(":8080", nil)
}
