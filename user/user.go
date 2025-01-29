package user

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ChessUser struct {
	Username         string
	Verified         bool
	Url              string
	ApiUrl           string
	Archives         []interface{}
	Info             map[string]interface{}
	UsernameNotFound bool
}

func New(username string) ChessUser {
	c := ChessUser{Username: username}
	c.init()
	return c
}

func (c *ChessUser) init() {
	apiUrl := "https://api.chess.com/pub/player/" + c.Username
	resp, err := http.Get(apiUrl)
	if err != nil {
		fmt.Println(err)
		c.UsernameNotFound = true
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	bodyMap := make(map[string]interface{})
	json.Unmarshal(body, &bodyMap)

	if len(bodyMap) == 0 || bodyMap["message"] != nil {
		c.UsernameNotFound = true
		return
	}

	c.Info = bodyMap
	if bodyMap["verified"] != nil {
		c.Verified = bodyMap["verified"].(bool)
	}

	if bodyMap["url"] != nil {
		c.Url = bodyMap["url"].(string)
	}

	c.ApiUrl = apiUrl

	c.GetArchives()
}

func (c *ChessUser) GetArchives() {
	if c.UsernameNotFound {
		return
	}
	resp, err := http.Get(c.ApiUrl + "/games/archives")
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	archives := make(map[string]interface{})
	json.Unmarshal(body, &archives)

	c.Archives = archives["archives"].([]interface{})
}
