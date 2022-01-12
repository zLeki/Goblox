package csrf

import (
	"github.com/zLeki/Goblox/account"
	"log"
	"net/http"
)

type token struct {
	XRSF string
}

var client http.Client

func GetCSRF(acc *account.Account) (string, error) {
	req, err := http.NewRequest("POST", "https://auth.roblox.com/v2/login", nil)
	req.AddCookie(acc.RobloSecurity)
	if err != nil {
		log.Fatalf("Error occured. Error is: %s", err.Error())
	}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		log.Fatalf("Error occured. Error is: %s", err.Error())
	}

	csrf := resp.Header.Get("X-CSRF-TOKEN")
	tokenvalue := token{
		csrf,
	}
	returnValue := &tokenvalue
	return (*returnValue).XRSF, nil

}
