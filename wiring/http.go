package wiring

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// HTTPRequest makes HTTP requests
func HTTPRequest(url string, apiKey string) []byte {
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("X-Developer-Key", apiKey)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic("HTTP Request Failed.")
	}
	if resp.StatusCode != http.StatusOK {
		StatusHelper(resp)
	}
	defer resp.Body.Close()
	d := getResponseBody(resp)
	return d
}

// StatusHelper outputs appropriate error messages to user based on the HTTP response status code.
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
