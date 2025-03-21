package db

import (
	"database/sql"
	"fmt"
	"go-chess/game"
	"go-chess/user"

	_ "github.com/glebarez/go-sqlite"
)

// Users Database
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

func AddUser(u user.ChessUser, db *sql.DB) {
	query := `INSERT INTO users (username) values(?);`
	db.Exec(query, u.Username)
}

func RemoveUser(u user.ChessUser, db *sql.DB) {
	query := `DELTE FROM users WHERE username = ` + u.Username + `;`
	db.Exec(query)
}

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
