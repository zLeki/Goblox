package groups

import (
	"github.com/zLeki/Goblox/account"
)

type Payout struct {
	UserId  int              `json:"Id"`
	Message string           `json:"errors"`
	Acount  *account.Account `json:"-"`
}
type GroupInfo struct {
	Description        string `json:"description"`
	Id                 int    `json:"id"`
	IsBuildersClubOnly bool   `json:"isBuildersClubOnly"`
	MemberCount        int    `json:"memberCount"`
	Name               string `json:"name"`
	PublicEntryAllowed bool   `json:"publicEntryAllowed"`
	Shout              *Shout `json:"shout"`
}
type UserData struct {
	RobloxGroup 	GroupInfo `json:"group"`
	RobloxRole  	RoleData  `json:"role"`
}
type Shout struct {
	Content string `json:"body"`
	Created string `json:"created"`
	Updated string `json:"updated"`
}
type RoleData struct {
	Id          	int    `json:"id"`
	Name        	string `json:"name"`
	Rank        	int    `json:"rank"`
	MemberCount 	int    `json:"memberCount"`
}
