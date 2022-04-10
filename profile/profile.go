package profile

import (
	"github.com/zLeki/Goblox/account"
	"github.com/zLeki/Goblox/formatter"
	"net/http"
	"strconv"
)

var client http.Client

func GetIdFromUsername(username string) UserID {
	req := formatter.FormatRequest(nil, "https://api.roblox.com/users/get-by-username?username="+username, "GET", nil)
	var data UserID
	return formatter.Decode(data, req)
}
func Info(username string) UserData {
	ID := GetIdFromUsername(username)
	req := formatter.FormatRequest(nil, "https://users.roblox.com/v1/users/"+strconv.Itoa(ID.PlayerID), "GET", nil)
	var data UserData
	formatter.Decode(data, req)
	return data
}
func CheckUser(name string, acc *account.Account) CheckData {
	dataBytes := []byte(`
		{
			"birthday": "1955-06-07T23:00:00.000Z",
			"context": "Signup",
			"username": "` + name + `"
		}`)
	resp := formatter.FormatRequest(acc, "https://auth.roblox.com/v1/usernames/validate", "POST", dataBytes)
	var data CheckData
	return formatter.Decode(data, resp)

}
