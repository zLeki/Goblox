package trade

import (
	"encoding/json"
	"log"
	"net/http"

	"errors"
	"strconv"
	"strings"

	"github.com/zLeki/Goblox/account"
	"github.com/zLeki/Goblox/formatter"
	"github.com/zLeki/Goblox/profile"
)

var client http.Client

func UserIDFromCookie(acc *account.Account) int {
	req := formatter.FormatRequest(acc, "http://www.roblox.com/mobileapi/userinfo", "GET", nil)
	req.AddCookie(acc.RobloSecurity)
	resp, _ := client.Do(req)
	var data IDFromCookie
	err := json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		log.Fatalf("Error decoding JSON response: %v", err)
	}
	return data.ID
}
func GetTrades(tradeType string, acc *account.Account) error {
	if tradeType != "outbound" || tradeType != "inbound" {
		return errors.New("Invalid trade type")
	}
	return nil
}
func SendTrade(acc *account.Account, theirUser string, yourItems []string, theirItems []string) (SendTradeInfo, error) {

	data := []byte(`
	{"offers":[{"userId":` + strconv.Itoa(UserIDFromCookie(acc)) + `,"userAssetIds":[` + strings.Join(yourItems, ",") + `],"robux":null},{"userId":` + strconv.Itoa(profile.GetIdFromUsername(theirUser).PlayerID) + `,"userAssetIds":[` + strings.Join(theirItems, ",") + `],"robux":null}]}
	`)
	req := formatter.FormatRequest(acc, "https://trades.roblox.com/v1/trades/send", "POST", data)
	req.AddCookie(acc.RobloSecurity)
	resp, _ := client.Do(req)

	if resp.StatusCode != http.StatusOK {
		var Errdata Error
		err := json.NewDecoder(resp.Body).Decode(&Errdata)
		if err != nil {
			return SendTradeInfo{}, errors.New(err.Error())
		}
		return SendTradeInfo{}, errors.New(Errdata.Errors[0].Message)
	}
	return SendTradeInfo{}, nil
}

// func DeclineTrade(acc *account.Account, tradeID int)
