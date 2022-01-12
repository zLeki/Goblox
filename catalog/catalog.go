package catalog

import (
	"encoding/json"
	"github.com/zLeki/Goblox/account"
	"github.com/zLeki/Goblox/formatter"
	"log"
	"net/http"
	"strconv"
)

var client http.Client

func GetInfo(id int, acc *account.Account) ItemData {
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
	req := formatter.FormatRequest(acc, "https://catalog.roblox.com/v1/catalog/items/details", "POST", data)
	req.AddCookie(acc.RobloSecurity)
	resp, _ := client.Do(req)
	if resp.StatusCode != 200 {
		log.Println(resp.StatusCode)
		return ItemData{}
	}
	var itemData ItemData
	err := json.NewDecoder(resp.Body).Decode(&itemData)
	if err != nil {
		log.Fatalf("Error saving to struct: %v", err)
	}
	return itemData

}
func ScrapeItems() []int {
	req := formatter.FormatRequest(nil, "https://catalog.roblox.com/v1/search/items?category=Collectibles&limit=60&subcategory=Collectibles", "GET", nil)
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	var data ItemsScraped
	err2 := json.NewDecoder(resp.Body).Decode(&data)
	if err2 != nil {
		log.Fatalf("Error saving to struct: %v", err)
	}
	var items []int
	for _, v := range data.Data {
		items = append(items, v.ID)
	}
	return items
}
