package main

import (
	"fmt"
	"net/http"
	"os"

	"go-chess/db"
	"go-chess/handlers"
	"go-chess/middleware"
)

func main() {
	os.Remove("./chess.db")
	usersDB := db.UsersInit()
	usersDB.Close()

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
