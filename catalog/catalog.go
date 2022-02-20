package catalog

import (
	"encoding/json"
	"errors"
	"github.com/zLeki/Goblox/account"
	"github.com/zLeki/Goblox/formatter"
	"log"
	"net/http"
	"strconv"
)
var client http.Client
func GetInfo(id int, acc *account.Account) (ItemData, error) {
	data := []byte(`
{
  "items": [
    {
      "itemType": "Asset",
      "id": ` + strconv.Itoa(id) + `
    }
  ]
}
`)
	resp := formatter.FormatRequest(acc, "https://catalog.roblox.com/v1/catalog/items/details", "POST", data)
	if resp.StatusCode != 200{
		return ItemData{}, errors.New("Error getting item info")
	}
	var itemData ItemData
	log.Println(resp.Body)
	return formatter.Decode(itemData, resp), nil

}
func ScrapeItems() []int {
	resp := formatter.FormatRequest(nil, "https://catalog.roblox.com/v1/search/items?category=Collectibles&limit=60&subcategory=Collectibles", "GET", nil)
	var data ItemsScraped
	err := json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		log.Fatalf("Error saving to struct: %v", err)
	}
	return nil //TODO
}
