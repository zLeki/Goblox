package games

import (
	"encoding/json"
	"github.com/zLeki/Goblox/formatter"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

var client http.Client

func GetInfo(placeid int) GameType {
	req := formatter.FormatRequest(nil, "https://games.roblox.com/v1/games?universeIds="+strconv.Itoa(placeid), "GET", nil)
	resp, _ := client.Do(req)
	io, _ := ioutil.ReadAll(resp.Body)
	log.Println(string(io))
	var data GameType
	err := json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		log.Fatalf("Error occured while decoding body: %v", err)
	}
	return data
}
