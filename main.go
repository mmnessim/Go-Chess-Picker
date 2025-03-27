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
