package app

import (
	"fmt"
	"io"
	"net/http"

	"github.com/spf13/viper"
)

// New pocketsmith client
func New() *Client {
	return &Client{
		HTTPClient: buildHTTPClient(),
		Config:     buildConfig(),
	}
}

func buildConfig() *Config {
	cfg := &Config{
		BaseURL: BaseURL,
		User:    viper.GetString("user_id"),
		APIKey:  viper.GetString("api_key"),
	}

	return cfg
}

func buildHTTPClient() Doer {
	return http.DefaultClient
}

// SendRequest sends a HTTP request to the Pocketsmith API
func (cl *Client) SendRequest(path, method string, body io.Reader) (io.ReadCloser, error) {
	url := fmt.Sprintf("%v%v", cl.Config.BaseURL, path)
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, fmt.Errorf("unable to create a new request: %w", err)
	}

	req.Header.Add("X-Developer-Key", cl.Config.APIKey)
	req.Header.Set("Content-Type", ResourceType)
	req.Header.Set("Accept", ResourceType)

	// if req.Method == "POST" || req.Method == "PUT" || req.Method == "PATCH" {
	// 	req.Header.Set("Idempotency-Key", uuid.NewString())
	// }

	resp, err := cl.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("unable to do request: %w", err)
	}

	// if resp.StatusCode == http.StatusTooManyRequests {
	// 	return nil, &RateLimitError{
	// 		Limit:     resp.Header.Get("X-Ratelimit-Limit"),
	// 		Remaining: resp.Header.Get("X-Ratelimit-Remaining"),
	// 		Reset:     resp.Header.Get("X-Ratelimit-Reset"),
	// 	}
	// }

	// if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusBadRequest {
	// 	apiErr := APIErrors{}
	// 	if err := jsonutil.DecodeJSON(resp.Body, &apiErr); err == nil {
	// 		return nil, apiErr
	// 	}
	// 	return nil, fmt.Errorf("%w: status code: %d", ErrUnknown, resp.StatusCode)
	// }

	return resp.Body, nil
}
