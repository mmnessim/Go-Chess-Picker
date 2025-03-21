package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"go-chess/handlers"
	"go-chess/middleware"
	randomuser "go-chess/randomUser"
)

func main() {
	os.Remove("./chess.db")
	if _, err := os.Stat("./users.db"); errors.Is(err, os.ErrNotExist) {
		// path/to/whatever does not exist
		randomuser.PopulateAllUsers()
	}
	u := randomuser.GetRandomUser()
	//fmt.Println(u)
	_ = u
	//usersDB := db.UsersInit()
	//usersDB.Close()

	index := middleware.Logging(handlers.Index)
	showGame := middleware.Logging(handlers.Game)
	showHistory := middleware.Logging(handlers.History)
	showGuess := middleware.Logging(handlers.Guess)

	fmt.Println("Listening on port 8080")

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", index)
	http.HandleFunc("/game", showGame)
	http.HandleFunc("/history", showHistory)
	http.HandleFunc("/guess", showGuess)
	http.ListenAndServe(":8080", nil)
}
