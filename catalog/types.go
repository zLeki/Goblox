package catalog

type ItemData struct {
	Data []struct {
		ID               int           `json:"id"`
		ItemType         string        `json:"itemType"`
		AssetType        int           `json:"assetType"`
		Name             string        `json:"name"`
		Description      string        `json:"description"`
		ProductID        int           `json:"productId"`
		Genres           []string      `json:"genres"`
		ItemStatus       []interface{} `json:"itemStatus"`
		ItemRestrictions []string      `json:"itemRestrictions"`
		CreatorType      string        `json:"creatorType"`
		CreatorTargetID  int           `json:"creatorTargetId"`
		CreatorName      string        `json:"creatorName"`
		LowestPrice      int           `json:"lowestPrice"`
		FavoriteCount    int           `json:"favoriteCount"`
		OffSaleDeadline  interface{}   `json:"offSaleDeadline"`
	} `json:"data"`
}
type ItemsScraped struct {
	Keyword            interface{} `json:"keyword"`
	PreviousPageCursor interface{} `json:"previousPageCursor"`
	NextPageCursor     string      `json:"nextPageCursor"`
	Data               []struct {
		ID       int    `json:"id"`
		ItemType string `json:"itemType"`
	} `json:"data"`
}
