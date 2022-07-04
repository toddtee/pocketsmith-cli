package app

import (
	"fmt"
	"io"
	"net/http"

	"github.com/spf13/viper"
)

// BuildConfig for the Pocketsmith API HTTP request
func BuildConfig() *Config {
	cfg := &Config{
		BaseURL: BaseURL,
		User:    viper.GetString("user_id"),
		APIKey:  viper.GetString("api_key"),
	}

	return cfg
}

// SendRequest sends a HTTP request to the Pocketsmith API
func SendRequest(cfg *Config, path string, method string, body io.Reader) (*http.Response, error) {
	url := fmt.Sprintf("%v%v", cfg.BaseURL, path)
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		fmt.Errorf("unable to create a new request: %w", err)
		return nil, err
	}

	req.Header.Add("X-Developer-Key", cfg.APIKey)
	req.Header.Set("Content-Type", ResourceType)
	req.Header.Set("Accept", ResourceType)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Errorf("unable to do request: %w", err)
		return nil, err
	}

	return resp, nil
}
