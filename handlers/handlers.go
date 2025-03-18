package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"go-chess/db"
	randomuser "go-chess/randomUser"
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
	var username string

	if r.Method == "POST" {
		username = r.FormValue("username")
	} else {
		username = r.URL.Query().Get("username")
	}

	u := user.New(username)

	if u.UsernameNotFound {
		// Redirect to index if invalid user
		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else {
		history := db.Init()
		defer history.Close()
		u.History = db.GetAll(history)

		u.GetRandomGame()
		db.Insert(u.Game, history)

		// Redirect to index if no random game can be found
		if u.Game.Err {
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}

		templ, err := template.ParseFiles("public/views/layout.html", "public/views/game.html")
		if err != nil {
			fmt.Fprintf(w, "Error %s", err)
		}

		templ.Execute(w, u)
	}

}

func Guess(w http.ResponseWriter, r *http.Request) {

	ul := randomuser.GetAllUsers()
	u := ul.GetRandomUser()

	if u.UsernameNotFound {
		// Redirect to index if invalid user
		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else {
		history := db.Init()
		defer history.Close()
		u.History = db.GetAll(history)

		u.GetRandomGame()
		db.Insert(u.Game, history)

		// Redirect to index if no random game can be found
		if u.Game.Err {
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}

		templ, err := template.ParseFiles("public/views/layout.html", "public/views/game.html")
		if err != nil {
			fmt.Fprintf(w, "Error %s", err)
		}

		templ.Execute(w, u)
	}

}

func History(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	idInt, err := strconv.ParseInt(id, 0, 64)
	if err != nil {
		fmt.Println(err)
	}

	history := db.Init()
	defer history.Close()

	g := db.GetById(int(idInt), history)
	_ = g

	templ, err := template.ParseFiles("public/views/layout.html", "public/views/history.html")
	if err != nil {
		fmt.Fprintf(w, "Error %s", err)
	}

	templ.Execute(w, g)

}
