package app

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/spf13/viper"
)

// APIError is a custom error to raise when Pocketsmith API returns a non-200 response code.
type APIError struct {
	Message string
}

func (err *APIError) Error() string {
	err.Message = "Bad Request: "
	return err.Message
}

// InitConfig sets up the configuration for the user
func InitConfig() error {
	if viper.GetString("config") != "" {
		viper.SetConfigFile(viper.GetString("config"))
	} else {
		viper.AddConfigPath("$HOME")
		viper.SetConfigName(".pocketsmith")
	}
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	// viper.Unmarshal()
	if err != nil {
		return fmt.Errorf("couldn't read in config %w", err)
	}
	return nil
}

// NewClient for the Pocketsmith API HTTP request
func NewClient() Client {
	var c Client
	c.BaseURL = BaseURL
	c.User = viper.GetString("user_id")
	c.APIKey = viper.GetString("api_key")
	return c
}

// sendRequest to the Pocketsmith API
func (c *Client) sendRequest(path string, method string, body io.Reader) (*http.Response, error) {
	url := fmt.Sprintf("%v%v", c.BaseURL, path)
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, fmt.Errorf("unable to create a new request: %w", err)
	}

	req.Header.Add("X-Developer-Key", c.APIKey)
	req.Header.Set("Content-Type", ResourceType)
	req.Header.Set("Accept", ResourceType)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("unable to do request: %w", err)
	}

	return resp, nil
}

// GetUser gets the authorised user of the Pocketsmith account
func (c *Client) GetUser(auth bool) (*User, error) {
	u := &User{}
	var path string
	// Check if authorised user flag was passed
	if auth == true {
		path = "/me"
	} else {
		path = "/users/" + c.User
	}
	resp, err := c.sendRequest(path, http.MethodGet, nil)
	if err != nil {
		return nil, fmt.Errorf("unable to send request: %w", err)
	}
	err = responseHandler(resp, u)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	return u, nil
}

func responseHandler(resp *http.Response, v interface{}) error {
	defer resp.Body.Close()
	d := getBodyData(resp.Body)
	if resp.StatusCode < http.StatusBadRequest {
		unmarshal(d, v)
	} else {
		err := APIError{}
		return fmt.Errorf("%s%s", err.Error(), string(d))
	}
	return nil
}

func getBodyData(r io.Reader) (d []byte) {
	d, _ = ioutil.ReadAll(r)
	return d
}

func unmarshal(d []byte, v interface{}) error {
	err := json.Unmarshal(d, v)
	if err != nil {
		return fmt.Errorf("unable to send request: %w", err)
	}
	return nil
}
