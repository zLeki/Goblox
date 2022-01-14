package groups

import (
	"encoding/json"
	"errors"
	"github.com/zLeki/Goblox/account"
	"github.com/zLeki/Goblox/formatter"
	"io"
	"log"
	"net/http"
	"sort"
	"strconv"
)
var client http.Client

func UserPayout(userID string, groupID int, amount int, acc *account.Account) error {
		var payload = []byte(`
		{
		  "PayoutType": "FixedAmount",
		  "Recipients": [
		    {
		      "recipientId": ` + userID + `,
		      "recipientType": "User",
		      "amount": ` + strconv.Itoa(amount) + ` 
		    }
		  ]
		}`)

		resp := formatter.FormatRequest(acc, "https://groups.roblox.com/v1/groups/"+strconv.Itoa(groupID)+"/payouts", "POST", payload)
		if resp.StatusCode == 400 {
			return errors.New("Incorrect userID or the user is new to the group.")
		}else if resp.StatusCode == 401 {
			return errors.New("Incorrect cookie")
		}
	return nil
}
func Groupinfo(ID string) *GroupInfo {
	req := formatter.FormatRequest(nil, "https://groups.roblox.com/v1/groups/"+ID, "GET", nil)
	resp, _ := client.Do(req)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)
	bodyBytes, _ := io.ReadAll(resp.Body)
	var data *GroupInfo
	return formatter.Decode(data, resp)

}
func GetUserRoleInGroup(userid int, groupid int, acc *account.Account) (RoleData, error) {
	resp := formatter.FormatRequest(acc, "https://groups.roblox.com/v1/users/"+strconv.Itoa(userid)+"/groups/roles", "GET", nil)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return RoleData{}, errors.New("Check status code: " + resp.Status)
	}
	type userData struct {
		Data 			[]UserData `json:"data"`
	}

	var jsonResponse userData
	err := json.NewDecoder(resp.Body).Decode(&jsonResponse)
	var role RoleData

	for _, d := range jsonResponse.Data {

		if d.RobloxGroup.Id == groupid {
			role = d.RobloxRole
		}
	}
	if err != nil {
		return RoleData{}, err
	}
	return role, nil
}

func SendShout(acc *account.Account, groupID string, message string) (int, error) {

	jsondata := []byte(`
		{
		  "message": "`+message+`"
		}
	`)
	resp := formatter.FormatRequest(acc, "https://groups.roblox.com/v1/groups/"+groupID+"/status", "PATCH", jsondata)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)
	if resp.StatusCode == http.StatusUnauthorized || resp.StatusCode == http.StatusForbidden {
		return 0, errors.New("You are not authorized to change the shout")
	}
	return resp.StatusCode, nil

}
func GetRoles(groupID int) []RoleData {
	resp := formatter.FormatRequest(nil, "https://groups.roblox.com/v1/groups/"+strconv.Itoa(groupID)+"/roles", "GET", nil)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)
	type roles struct {
		GroupId 		int `json:"groupId"`
		Roles   		[]RoleData `json:"roles"`
	}
	var GroupRole roles
	err := json.NewDecoder(resp.Body).Decode(&GroupRole)
	if err != nil {
		return nil
	}
	sort.Slice(GroupRole.Roles, func(i, j int) bool {
		return GroupRole.Roles[i].Rank < GroupRole.Roles[j].Rank
	})
	return GroupRole.Roles
}
