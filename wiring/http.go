package wiring

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// StatusHelper outputs appropriate error messages to user based on the HTTP response status code.
// Need to offload this to the pocketsmith client
func StatusHelper(resp *http.Response) {
	b := getResponseBody(resp)
	if resp.StatusCode == http.StatusUnauthorized {
		fmt.Printf("ERROR %v: Please check API Key!\n%s\n", resp.StatusCode, b)
	}
	if resp.StatusCode == http.StatusNotFound {
		fmt.Printf("ERROR %v: Please Check Pocketsmith User Id.\n%s\n", resp.StatusCode, b)
	}
}

func getResponseBody(resp *http.Response) []byte {
	body, _ := ioutil.ReadAll(resp.Body)
	return body
}
