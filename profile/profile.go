package profile

import (
	"encoding/json"
	"github.com/zLeki/Goblox/formatter"
	"log"
	"net/http"
	"strconv"
)


var client http.Client
func GetIdFromUsername(username string) UserID {
	req := formatter.FormatRequest(nil, "https://api.roblox.com/users/get-by-username?username="+username, "GET", nil)
	resp, _ := client.Do(req)
	var data UserID
	err := json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		log.Fatalf("Error saving to struct: %v", err)
	}
	return data
}
func Info(username string) UserData {
	ID := GetIdFromUsername(username)
	req := formatter.FormatRequest(nil, "https://users.roblox.com/v1/users/"+strconv.Itoa(ID.PlayerID), "GET", nil)
	resp, _ := client.Do(req)
	var data UserData
	err := json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		log.Fatalf("Error saving to struct: %v", err)
	}
	return data
}
func CheckUser(name string) CheckData {
	dataBytes := []byte(`
	{
	  "username": "`+name+`",
	  "birthday": "2022-01-12T23:36:12.684Z",
	  "context": "Unknown"
	}
	`)
	resp := formatter.FormatRequest(nil, "https://auth.roblox.com/v1/usernames/validate", "POST", dataBytes)
	var data CheckData
	err := json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		log.Fatalf("Error saving to struct: %v", err)
	}
	log.Println(data)
	return data
}