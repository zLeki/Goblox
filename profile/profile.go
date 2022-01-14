package profile

import (
	"github.com/zLeki/Goblox/formatter"
	"log"
	"net/http"
	"strconv"
)

var client http.Client
func GetIdFromUsername(username string) UserID {
	req := formatter.FormatRequest(nil, "https://api.roblox.com/users/get-by-username?username="+username, "GET", nil)
	var data UserID
	return formatter.Decode(data, req)
}
func Info(username string) {
	ID := GetIdFromUsername(username)
	req := formatter.FormatRequest(nil, "https://users.roblox.com/v1/users/"+strconv.Itoa(ID.PlayerID), "GET", nil)
	var data UserData
	formatter.Decode(data, req)
}
func CheckUser(name string) int {
	ID := GetIdFromUsername(name)
	resp := formatter.FormatRequest(nil, "https://www.roblox.com/users/"+strconv.Itoa(ID.PlayerID)+"/profile", "GET", nil)
	return resp.StatusCode

}