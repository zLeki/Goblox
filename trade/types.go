package trade
import (
	"time"
)
type SendTradeInfo struct {
	ID int `json:"id"`
}
type IDFromCookie struct {
	ID int `json:"UserID"`
}
type Error struct {
	Errors []struct {
		Code              int    `json:"code"`
		Message           string `json:"message"`
		UserFacingMessage string `json:"userFacingMessage"`
		Field             string `json:"field"`
		FieldData         []struct {
			UserAssetID int64  `json:"userAssetId"`
			Reason      string `json:"reason"`
		} `json:"fieldData"`
	} `json:"errors"`
}
type TradeData struct {
	PreviousPageCursor string `json:"previousPageCursor"`
	NextPageCursor     string `json:"nextPageCursor"`
	Data               []struct {
		ID   int `json:"id"`
		User struct {
			ID          int    `json:"id"`
			Name        string `json:"name"`
			DisplayName string `json:"displayName"`
		} `json:"user"`
		Created    time.Time `json:"created"`
		Expiration time.Time `json:"expiration"`
		IsActive   bool      `json:"isActive"`
		Status     string    `json:"status"`
	} `json:"data"`
}