package games

import "time"

type GameType struct {
	Data []struct {
		ID          int64  `json:"id"`
		RootPlaceID int64  `json:"rootPlaceId"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Creator     struct {
			ID           int64  `json:"id"`
			Name         string `json:"name"`
			Type         string `json:"type"`
			IsRNVAccount bool   `json:"isRNVAccount"`
		} `json:"creator"`
		Price                     interface{}   `json:"price"`
		AllowedGearGenres         []string      `json:"allowedGearGenres"`
		AllowedGearCategories     []interface{} `json:"allowedGearCategories"`
		IsGenreEnforced           bool          `json:"isGenreEnforced"`
		CopyingAllowed            bool          `json:"copyingAllowed"`
		Playing                   int           `json:"playing"`
		Visits                    int           `json:"visits"`
		MaxPlayers                int           `json:"maxPlayers"`
		Created                   time.Time     `json:"created"`
		Updated                   time.Time     `json:"updated"`
		StudioAccessToApisAllowed bool          `json:"studioAccessToApisAllowed"`
		CreateVipServersAllowed   bool          `json:"createVipServersAllowed"`
		UniverseAvatarType        string        `json:"universeAvatarType"`
		Genre                     string        `json:"genre"`
		IsAllGenre                bool          `json:"isAllGenre"`
		IsFavoritedByUser         bool          `json:"isFavoritedByUser"`
		FavoritedCount            int           `json:"favoritedCount"`
	} `json:"data"`
}
