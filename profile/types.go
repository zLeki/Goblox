package profile

import "time"

type UserID struct {
	PlayerID int `json:"Id"`
}

type UserData struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	Banned      bool      `json:"isBanned"`
	ID          int       `json:"id"`
	DisplayName string    `json:"displayName"`
}
type FriendData struct {
	Data []struct {
		IsOnline               bool        `json:"isOnline"`
		PresenceType           int         `json:"presenceType"`
		IsDeleted              bool        `json:"isDeleted"`
		FriendFrequentRank     int         `json:"friendFrequentRank"`
		Description            interface{} `json:"description"`
		Created                time.Time   `json:"created"`
		IsBanned               bool        `json:"isBanned"`
		ExternalAppDisplayName interface{} `json:"externalAppDisplayName"`
		ID                     int64       `json:"id"`
		Name                   string      `json:"name"`
		DisplayName            string      `json:"displayName"`
	} `json:"data"`
}

type FollowerData struct {
	PreviousPageCursor string `json:"previousPageCursor"`
	NextPageCursor     string `json:"nextPageCursor"`
	Data               []struct {
		IsOnline               bool      `json:"isOnline"`
		PresenceType           string    `json:"presenceType"`
		IsDeleted              bool      `json:"isDeleted"`
		FriendFrequentRank     int       `json:"friendFrequentRank"`
		Description            string    `json:"description"`
		Created                time.Time `json:"created"`
		IsBanned               bool      `json:"isBanned"`
		ExternalAppDisplayName string    `json:"externalAppDisplayName"`
		ID                     int       `json:"id"`
		Name                   string    `json:"name"`
		DisplayName            string    `json:"displayName"`
	} `json:"data"`
}
