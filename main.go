package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"go-chess/db"
	"go-chess/handlers"
	"go-chess/middleware"
)

// TODO
// Refactor randomUser.go
// UserList struct can probably be removed altogether
// PopulateAllUsers() and GetRandomUser() should be put into db.go
// db.go should maybe be renamed?
//
// PopulateAllUsers() should have the ability to add new users into the db
// maybe with a whole new function
//
// index.html needs a button to get to Guess-The-ELO
// Some kind of result message should appear based on guesses
// Maybe there can be six guesses like wordle?

func main() {
	os.Remove("./chess.db")
	if _, err := os.Stat("./users.db"); errors.Is(err, os.ErrNotExist) {
		list := db.New()
		defer list.Database.Close()
		list.PopulateAllUsers()
		db.GetLeaderboard()
	}

	index := middleware.Logging(handlers.Index)
	showGame := middleware.Logging(handlers.Game)
	showHistory := middleware.Logging(handlers.History)
	showGuess := middleware.Logging(handlers.Guess)
	showAbout := middleware.Logging(handlers.About)

	fmt.Println("Listening on port 8080")

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", index)
	http.HandleFunc("/game", showGame)
	http.HandleFunc("/history", showHistory)
	http.HandleFunc("/guess", showGuess)
	http.HandleFunc("/about", showAbout)
	http.ListenAndServe(":8080", nil)
}
