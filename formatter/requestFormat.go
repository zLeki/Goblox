package formatter
import (
	"bytes"
	"github.com/zLeki/Goblox/account"
	csrf2 "github.com/zLeki/Goblox/csrf"
	"net/http"
)
func FormatRequest(acc *account.Account, url string, method string, json []byte) *http.Request {
	req, _ := http.NewRequest(method, url, bytes.NewBuffer(json))
	if acc != nil {
		csrf, _ := csrf2.GetCSRF(acc)
		req.Header.Set("X-CSRF-TOKEN", csrf)
		req.Header.Set("Content-Type", "application/json")
		return req
	}
	req.Header.Set("Content-Type", "application/json")
	return req
}