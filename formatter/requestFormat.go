package formatter

import (
	"bytes"
	"net/http"

	"io/ioutil"
	"log"

	"github.com/zLeki/Goblox/account"
	csrf2 "github.com/zLeki/Goblox/csrf"
)

var client *http.Client

func FormatRequest(acc *account.Account, url string, method string, json []byte) *http.Response {
	req, _ := http.NewRequest(method, url, bytes.NewBuffer(json))
	if acc != nil {
		csrf, _ := csrf2.GetCSRF(acc)
		req.Header.Set("X-CSRF-TOKEN", csrf)

		resp, _ := client.Do(req)
		return resp
	}
	req.Header.Set("Content-Type", "application/json")
	resp, _ := client.Do(req)
	iou, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("error", err)
	}
	log.Println(string(iou))
	return resp
}
