package main

import (
	"fmt"
	"net/http"

	"go-chess/handlers"
	"go-chess/middleware"
)

func main() {
	index := middleware.Logging(handlers.Index)
	game := middleware.Logging(handlers.Game)

	fmt.Println("Listening on port 8080")

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", index)
	http.HandleFunc("/game", game)
	http.ListenAndServe(":8080", nil)
}
