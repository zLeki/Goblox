package formatter
import (
	"bytes"
	"encoding/json"
	"github.com/zLeki/Goblox/account"
	csrf2 "github.com/zLeki/Goblox/csrf"
	"log"
	"net/http"
)
var client http.Client
func FormatRequest(acc *account.Account, url string, method string, json []byte) *http.Response {
	req, _ := http.NewRequest(method, url, bytes.NewBuffer(json))
	if acc != nil {
		csrf, _ := csrf2.GetCSRF(acc)
		req.Header.Set("X-CSRF-TOKEN", csrf)
	}
	req.Header.Set("Content-Type", "application/json")
	resp, _ := client.Do(req)
	return resp
}
func Decode[T any](t T, body *http.Response) T {
	err := json.NewDecoder(body.Body).Decode(&t)
	if err != nil {
		log.Fatalf("error decoding response", err)
	}
	return t
}