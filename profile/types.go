package profile

import "time"

type UserID struct {
	PlayerID int `json:"Id"`

}
type CheckData struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
type UserData struct {
	Name string `json:"name"`
	Description string `json:"description"`
	CreatedAt time.Time `json:"created_at"`
	Banned bool `json:"isBanned"`
	ID int `json:"id"`
	DisplayName string `json:"displayName"`
}