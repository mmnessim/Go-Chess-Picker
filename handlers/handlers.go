package handlers

import (
	"fmt"
	"html/template"
	"net/http"

	"go-chess/game"
	"go-chess/user"
)

func Index(w http.ResponseWriter, r *http.Request) {
	user := user.New("tenderllama")
	randomGame := game.GetRandomGame(&user)

	templ, err := template.ParseFiles("public/layout.html", "public/index.html")
	if err != nil {
		fmt.Fprintf(w, "Error %s", err)
	}

	templ.Execute(w, randomGame)
}
