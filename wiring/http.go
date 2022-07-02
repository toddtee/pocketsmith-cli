package wiring

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func HttpRequest(url string, api_key string) []byte {
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("X-Developer-Key", api_key)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic("HTTP Request Failed.")
	}
	if resp.StatusCode != http.StatusOK {
		StatusHelper(resp)
	}
	defer resp.Body.Close()
	d := GetResponseBody(resp)
	return d
}

func StatusHelper(resp *http.Response) {
	b := GetResponseBody(resp)
	if resp.StatusCode == http.StatusUnauthorized {
		fmt.Printf("ERROR %v: Please check API Key!\n%s\n", resp.StatusCode, b)
	}
	if resp.StatusCode == http.StatusNotFound {
		fmt.Printf("ERROR %v: Please Check Pocketsmith User Id.\n%s\n", resp.StatusCode, b)
	}
}

func GetResponseBody(resp *http.Response) []byte {
	body, _ := ioutil.ReadAll(resp.Body)
	return body
}
