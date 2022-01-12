package profile

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/zLeki/Goblox/account"
	"github.com/zLeki/Goblox/formatter"
)

var client http.Client

func GetIdFromUsername(username string) UserID {
	resp := formatter.FormatRequest(nil, "https://api.roblox.com/users/get-by-username?username="+username, "GET", nil)
	var data UserID
	err := json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		log.Fatalf("Error saving to struct: %v", err)
	}
	return data
}
func Info(username string) UserData {
	ID := GetIdFromUsername(username)
	resp := formatter.FormatRequest(nil, "https://users.roblox.com/v1/users/"+strconv.Itoa(ID.PlayerID), "GET", nil)
	var data UserData
	err := json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		log.Fatalf("Error saving to struct: %v", err)
	}
	return data
}
func GetFriends(username string) FriendData {
	ID := GetIdFromUsername(username)
	resp := formatter.FormatRequest(nil, "https://friends.roblox.com/v1/users/"+strconv.Itoa(ID.PlayerID)+"/friends?userSort=Alphabetical&_=1641935441685", "GET", nil)

	var data FriendData
	err := json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		log.Fatalf("Error saving to struct: %v", err)
	}
	return data
}
func GetFollowers(username string, limit int) FollowerData {
	ID := GetIdFromUsername(username)
	resp := formatter.FormatRequest(nil, "https://friends.roblox.com/v1/users/"+strconv.Itoa(ID.PlayerID)+"/followings?sortOrder=Asc&limit="+strconv.Itoa(limit)+"&_=1641940130441", "GET", nil)

	var data FollowerData
	err := json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		log.Fatalf("Error saving to struct: %v", err)
	}
	return data
}
func AcceptFriend(username string, acc *account.Account) (error, string) {
	ID := GetIdFromUsername(username)
	resp := formatter.FormatRequest(nil, "https://friends.roblox.com/v1/users/"+strconv.Itoa(ID.PlayerID)+"/accept-friend-request", "POST", nil)

	switch resp.StatusCode {
	case 200:
		return nil, "Success"
	}
	return errors.New("Error while accepting friend request Status: " + resp.Status), ""
}
func DeclineFriend(username string, acc *account.Account) (error, string) {
	ID := GetIdFromUsername(username)
	resp := formatter.FormatRequest(nil, "https://friends.roblox.com/v1/users/"+strconv.Itoa(ID.PlayerID)+"/decline-friend-request", "POST", nil)

	switch resp.StatusCode {
	case 200:
		return nil, "Success"
	}
	return errors.New("Error while declining friend request Status: " + resp.Status), ""
}
func SendFriendRequest(username string, acc *account.Account) (error, string) {
	ID := GetIdFromUsername(username)
	resp := formatter.FormatRequest(nil, "https://friends.roblox.com/v1/users/"+strconv.Itoa(ID.PlayerID)+"/request-friendship", "POST", nil)

	switch resp.StatusCode {
	case 200:
		return nil, "Success"
	}
	return errors.New("Error while sending friend request Status: " + resp.Status), ""
}
func GetThumbnail(username string) string {
	return "https://www.roblox.com/headshot-thumbnail/image?userId=" + strconv.Itoa(GetIdFromUsername(username).PlayerID) + "&width=420&height=420&format=png"
}
