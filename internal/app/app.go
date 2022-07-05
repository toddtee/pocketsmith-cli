package app

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/spf13/viper"
)

// InitConfig sets up the configuration for the user
func InitConfig() {
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
		fmt.Errorf("couldn't read in config %w", err)
	}
}

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

func GetAuthorisedUser() error {
	u := User{}
	cfg := BuildConfig()
	path := "/me"
	resp, err := SendRequest(cfg, path, http.MethodGet, nil)
	if err != nil {
		return fmt.Errorf("unable to send request: %w", err)
	}
	defer resp.Body.Close()
	d := getBodyData(resp.Body)
	unmarshal(d, &u)

	fmt.Println(&u)
	return nil
}

func getBodyData(r io.Reader) (d []byte) {
	d, _ = ioutil.ReadAll(r)
	return d
}

func unmarshal(d []byte, v interface{}) {
	err := json.Unmarshal(d, &v)
	if err != nil {
		fmt.Errorf("unable to send request: %w", err)
	}
}
