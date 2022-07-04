package app

import (
	"fmt"
	"net/http"
)

// AuthorisedUser gets the Authorised User of the Pocketsmith Account
func AuthorisedUser(cfg *Config) *http.Response {
	path := "/me"
	resp, err := SendRequest(cfg, path, http.MethodGet, nil)
	if err != nil {
		fmt.Println("woops")
	}
	return resp

	// The below block needs to be extracted to a json decoder and passed to a "print result" function.
	// return reader
}
