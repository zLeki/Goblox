package csrf

import (
	"log"
	"net/http"

	"github.com/zLeki/Goblox/account"
)

type token struct {
	XRSF string
}

var client http.Client

func GetCSRF(acc *account.Account) (string, error) {
	resp, err := http.NewRequest("POST", "https://auth.roblox.com/v2/login", nil)

	if err != nil {
		log.Fatalf("Error occured. Error is: %s", err.Error())
	}

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
