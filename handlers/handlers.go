package handlers

import (
	"fmt"
	"html/template"
	"net/http"

	"go-chess/game"
	"go-chess/user"
)

func Index(w http.ResponseWriter, r *http.Request) {

	templ, err := template.ParseFiles("public/views/layout.html", "public/views/index.html")
	if err != nil {
		fmt.Fprintf(w, "Error %s", err)
	}

	templ.Execute(w, nil)
}

func Game(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	// For testing
	fmt.Println(username)

	person := user.New(username)
	if person.UsernameNotFound {
		// Redirect to index if invalid user
		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else {
		randomGame := game.GetRandomGame(&person)
		// Redirect to index if no random game can be found
		if randomGame.Err {
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}

		templ, err := template.ParseFiles("public/views/layout.html", "public/views/game.html")
		if err != nil {
			fmt.Fprintf(w, "Error %s", err)
		}

		templ.Execute(w, randomGame)
	}

}
