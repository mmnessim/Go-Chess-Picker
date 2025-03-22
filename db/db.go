package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"go-chess/game"
	"go-chess/user"
	"io"
	"net/http"

	_ "github.com/glebarez/go-sqlite"
)

type AllUsers struct {
	Database *sql.DB
}

func New() *AllUsers {
	db, err := sql.Open("sqlite", "./users.db")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Connected to users database")
	}
	query := `CREATE TABLE IF NOT EXISTS users (
		ID INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT
		);`
	_, err = db.Exec(query)

	if err != nil {
		fmt.Println(err)
	}

	return &AllUsers{
		Database: db,
	}
}

func (al *AllUsers) PopulateAllUsers() {
	userlist := getAllUsers()

	for _, u := range userlist.Players {
		al.AddUser(u)
		fmt.Printf("Added user %s\n", u)
	}
}

// Editing DB functions
func (al *AllUsers) AddUser(username string) {
	query := `INSERT INTO users (username) values(?);`
	al.Database.Exec(query, username)
}

func (al *AllUsers) RemoveUser(username string) {
	query := `DELETE FROM users WHERE username = ` + username + `;`
	_, err := al.Database.Exec(query)
	if err != nil {
		fmt.Println(err)
	}
}

// TODO FIX THIS
func (al *AllUsers) GetRandomUser() user.ChessUser {

	type record struct {
		Id       int
		Username string
	}

	query := `SELECT * FROM users ORDER BY RANDOM() LIMIT 1;`
	row, err := al.Database.Query(query)
	if err != nil {
		fmt.Println(err)
	}

	var result record = record{}

	for row.Next() {
		err = row.Scan(&result.Id, &result.Username)
		if err != nil {
			fmt.Println(err)
		}
	}

	user := user.New(result.Username)
	user.GetRandomGame()
	if user.Game.Err {
		fmt.Println("User has no games, removing...")
		al.RemoveUser(result.Username)
		al.GetRandomUser()
	}
	return user
}

// Helpers to populate the database
type UserList struct {
	Players []string `json:"players"`
}

func getAllUsers() UserList {
	resp, err := http.Get("https://api.chess.com/pub/country/US/players")
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	users := UserList{}
	json.Unmarshal(body, &users)

	return users
}

/* CAN PROBABLY BE DELETED
func UsersInit() *sql.DB {
	db, err := sql.Open("sqlite", "./users.db")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Connected to users database")
	}

	query := `CREATE TABLE IF NOT EXISTS users (
	ID INTEGER PRIMARY KEY AUTOINCREMENT,
	username TEXT
	);`
	_, err = db.Exec(query)

	if err != nil {
		fmt.Println(err)
	}

	return db
}
*/

// History Database functions

func Init() *sql.DB {
	db, err := sql.Open("sqlite", "./chess.db")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Connected to database")
	}

	sql := `CREATE TABLE IF NOT EXISTS history (
	ID INTEGER PRIMARY KEY AUTOINCREMENT,
	pgn TEXT,
	blackUsername TEXT,
	blackRating REAL,
	blackResult TEXT,
	whiteUsername TEXT,
	whiteRating REAL,
	whiteResult TEXT,
	Url TEXT
	);`
	_, err = db.Exec(sql)

	if err != nil {
		fmt.Println(err)
	}

	return db
}

func Insert(g game.Game, db *sql.DB) {
	sql := `INSERT INTO history
	(pgn, blackUsername, blackRating, blackResult,
	whiteUsername, whiteRating, whiteResult, Url)
	values(?, ?, ?, ?, ?, ?, ?, ?);
	`
	db.Exec(sql, g.Pgn, g.Black.Username, g.Black.Rating, g.Black.Result,
		g.White.Username, g.White.Rating, g.White.Result, g.Url)
}

func GetById(id int, db *sql.DB) game.Game {
	sql := fmt.Sprintf("SELECT * FROM history WHERE id = %d;", id)
	rows, err := db.Query(sql)
	if err != nil {
		fmt.Println(err)
	}
	record := game.Game{}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(
			&record.ID, &record.Pgn, &record.Black.Username, &record.Black.Rating,
			&record.Black.Result, &record.White.Username, &record.White.Rating,
			&record.White.Result, &record.Url)
		if err != nil {
			fmt.Println(err)
		}
	}
	return record
}

func GetAll(db *sql.DB) []game.Game {
	sql := "SELECT * FROM history;"
	rows, err := db.Query(sql)
	if err != nil {
		fmt.Println(err)
	}
	var history []game.Game
	defer rows.Close()
	for rows.Next() {
		record := game.Game{}
		err = rows.Scan(
			&record.ID, &record.Pgn, &record.Black.Username, &record.Black.Rating,
			&record.Black.Result, &record.White.Username, &record.White.Rating,
			&record.White.Result, &record.Url)
		if err != nil {
			fmt.Println(err)
		} else {
			history = append(history, record)
		}
	}
	return history
}
