package db

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetLeaderboard() {
	url := "https://api.chess.com/pub/leaderboards"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	type Record struct {
		Daily []struct {
			PlayerID   int    `json:"player_id"`
			ID         string `json:"@id"`
			URL        string `json:"url"`
			Username   string `json:"username"`
			Score      int    `json:"score"`
			Rank       int    `json:"rank"`
			Country    string `json:"country"`
			Title      string `json:"title,omitempty"`
			Name       string `json:"name,omitempty"`
			Status     string `json:"status"`
			Avatar     string `json:"avatar"`
			TrendScore struct {
				Direction int `json:"direction"`
				Delta     int `json:"delta"`
			} `json:"trend_score"`
			TrendRank struct {
				Direction int `json:"direction"`
				Delta     int `json:"delta"`
			} `json:"trend_rank"`
			FlairCode string `json:"flair_code"`
			WinCount  int    `json:"win_count"`
			LossCount int    `json:"loss_count"`
			DrawCount int    `json:"draw_count"`
		} `json:"daily"`
	}

	var topPlayers Record = Record{}

	json.Unmarshal(body, &topPlayers)

	fmt.Println(topPlayers)

	database := New()

	for _, p := range topPlayers.Daily {
		database.AddUser(p.Username)
		fmt.Println("Added", p.Username)
	}
}
