package handlers

import (
	"fmt"
	"html/template"
	"net/http"

	"go-chess/game"
	"go-chess/user"
)

func Index(w http.ResponseWriter, r *http.Request) {
	person := user.New("tenderllama")
	randomGame := game.GetRandomGame(&person)

	templ, err := template.ParseFiles("public/layout.html", "public/index.html")
	if err != nil {
		fmt.Fprintf(w, "Error %s", err)
	}

	templ.Execute(w, randomGame)
}

func Game(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	// For testing
	fmt.Println(username)

	person := user.New(username)
	if person.UsernameNotFound {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

	randomGame := game.GetRandomGame(&person)

	templ, err := template.ParseFiles("public/layout.html", "public/game.html")
	if err != nil {
		fmt.Fprintf(w, "Error %s", err)
	}

	templ.Execute(w, randomGame)
}
