package main

import (
	"fmt"
	"net/http"

	"go-chess/db"
	"go-chess/handlers"
	"go-chess/middleware"
	"go-chess/user"
)

func main() {
	index := middleware.Logging(handlers.Index)
	showGame := middleware.Logging(handlers.Game)

	database := db.Init()
	defer database.Close()

	u := user.New("tenderllama")
	u.GetRandomGame()

	db.Insert(u.Game, database)
	db.GetById(1, database)

	fmt.Println("Listening on port 8080")

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", index)
	http.HandleFunc("/game", showGame)
	http.ListenAndServe(":8080", nil)
}
