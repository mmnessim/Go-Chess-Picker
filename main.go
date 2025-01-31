package main

import (
	"fmt"
	"net/http"

	"go-chess/handlers"
)

func main() {

	fmt.Println("Listening on port 8080")

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", handlers.Index)
	http.HandleFunc("/game", handlers.Game)
	http.ListenAndServe(":8080", nil)
}
