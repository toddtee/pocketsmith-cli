/*
Copyright © 2022 Todd Turner hi@toddtee.sh

*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type User struct {
	Id                         int    `json:"id"`
	Login                      string `json:"login"`
	Name                       string `json:"name"`
	Email                      string `json:"email"`
	Avatar_url                 string `json:"avatar_url"`
	Beta_user                  bool   `json:"beta_user"`
	Time_zone                  string `json:"time_zone"`
	Week_start_day             int    `json:"week_start_day"`
	Is_reviewing_transactions  bool   `json:"is_reviewing_transactions"`
	Base_currency_code         string `json:"base_currency_code"`
	Always_show_base_currency  bool   `json:"always_show_base_currency"`
	Using_multiple_currencies  bool   `json:"using_multiple_currencies"`
	Available_accounts         int    `json:"available_accounts"`
	Available_budgets          int    `json:"available_budgets"`
	Forecast_last_updated_at   string `json:"forecast_last_updated_at"`
	Forecast_last_accessed_at  string `json:"forecast_last_accessed_at"`
	Forecast_start_date        string `json:"forecast_start_date"`
	Forecast_end_date          string `json:"forecast_end_date"`
	Forecast_defer_recalculate bool   `json:"forecast_defer_recalculate"`
	Forecast_needs_recalculate bool   `json:"forecast_needs_recalculate"`
	Last_logged_in_at          string `json:"last_logged_in_at"`
	Last_activity_at           string `json:"last_activity_at"`
	Created_at                 string `json:"created_at"`
	Updated_at                 string `json:"updated_at"`
}

type config struct {
	User_id string
	Api_key string
}

func getConfig() config {
	user, _ := rootCmd.Flags().GetString("user")
	if user == "" {
		user = viper.GetString("user_id")
	}
	api_key, _ := rootCmd.Flags().GetString("api_key")
	if api_key == "" {
		api_key = viper.GetString("api_key")
	}

	c := config{User_id: user, Api_key: api_key}
	return c
}

//getAccounts lists bank accounts added to the pocketsmith account
func getAccounts(cmd *cobra.Command, args []string) {
	c := getConfig()
	url := fmt.Sprintf("https://api.pocketsmith.com/v2/users/%v/accounts", c.User_id)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("X-Developer-Key", c.Api_key)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Was unable to list account details.")
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(res)
	fmt.Println(string(body))
}

func getAuthorisedUser(cmd *cobra.Command, args []string) {
	c := getConfig()
	user := User{}
	url := "https://api.pocketsmith.com/v2/me"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("X-Developer-Key", c.Api_key)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Was unable to get user.")
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	b := []byte(body)
	err = json.Unmarshal(b, &user)
	if err != nil {
		fmt.Println("Was unable to unmarshal user data to user struct.")
	}
	fmt.Println(user.Name)
}

// getCmd represents the list command
var getCmd = &cobra.Command{
	Use:   "get", // Need to add a better  [-F file | -D dir]... [-f format] profile
	Short: "retrieves one or many pocketsmith account resources",
	Long:  `prints some import information about the specified pocketsmith resource.`,
}

var getAccountsCmd = &cobra.Command{
	Use:   "accounts", // Need to add a better  [-F file | -D dir]... [-f format] profile
	Short: "retrieves bank account/s associated with the pocketsmith account",
	Long:  `shows where the gold is kept`,
	Run:   getAccounts,
}

var getAuthorisedUserCmd = &cobra.Command{
	Use:   "user", // Need to add a better  [-F file | -D dir]... [-f format] profile
	Short: "retrieves the authorised user of the pocketsmith account",
	Long:  `reveals who is the supreme overlord of the household finances.`,
	Run:   getAuthorisedUser,
}

func init() {
	rootCmd.AddCommand(getCmd)
	getCmd.AddCommand(getAccountsCmd)
	getCmd.AddCommand(getAuthorisedUserCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// accountsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// accountsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
