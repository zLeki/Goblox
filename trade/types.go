package trade

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
