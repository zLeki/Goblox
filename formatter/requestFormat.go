package formatter

import (
	"bytes"
	"net/http"

	"github.com/zLeki/Goblox/account"
	csrf2 "github.com/zLeki/Goblox/csrf"
)

var client *http.Client

func FormatRequest(acc *account.Account, url string, method string, json []byte) *http.Response {
	req, _ := http.NewRequest(method, url, bytes.NewBuffer(json))
	if acc != nil {
		csrf, _ := csrf2.GetCSRF(acc)
		req.Header.Set("X-CSRF-TOKEN", csrf)

		err, _ := client.Do(req)
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	err, _ := client.Do(req)
	return err
}
