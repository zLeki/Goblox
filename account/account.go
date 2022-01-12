package account

import (
	"net/http"
)

type Account struct {
	RobloSecurity *http.Cookie
}

func Validate(RobloSecurity interface{}) *Account {
	if RobloSecurity == nil {
		RobloSecurity = ""
	}
	Cookie := http.Cookie{Name: ".ROBLOSECURITY", Value: RobloSecurity.(string)}
	account := Account{
		RobloSecurity: &Cookie,
	}
	return &account
}
