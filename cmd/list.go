/*
Package cmd provides a command line application for Pocketsmith.
Copyright © 2022 Todd Turner hi@toddtee.sh

*/
package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/toddtee/pocketsmith-cli/pocketsmith"
	"github.com/toddtee/pocketsmith-cli/wiring"
)

type config struct {
	UserID string
	APIKey string
}

func getConfig() config {
	user, _ := rootCmd.Flags().GetString("user")
	if user == "" {
		user = viper.GetString("user_id")
	}
	apiKey, _ := rootCmd.Flags().GetString("api-key")
	if apiKey == "" {
		apiKey = viper.GetString("api_key")
	}

	c := config{UserID: user, APIKey: apiKey}
	return c
}

// getAccounts lists bank accounts added to the pocketsmith account
func getAccounts(cmd *cobra.Command, args []string) {
	c := getConfig()
	url := fmt.Sprintf("https://api.pocketsmith.com/v2/users/%v/accounts", c.UserID)
	d := wiring.HTTPRequest(url, c.APIKey)
	fmt.Printf("%v", string(d))
}

func getAuthorisedUser(cmd *cobra.Command, args []string) {
	c := getConfig()
	user := pocketsmith.User{}
	url := "https://api.pocketsmith.com/v2/me"
	d := wiring.HTTPRequest(url, c.APIKey)
	err := json.Unmarshal(d, &user)
	if err != nil {
		panic("Was unable to unmarshal user data to user struct.")
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
