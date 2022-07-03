/*
Package cmd provides a command line application for Pocketsmith.
Copyright © 2022 Todd Turner hi@toddtee.sh

*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/spf13/cobra"
	"github.com/toddtee/pocketsmith-cli/pocketsmith"
)

// getAccounts lists bank accounts added to the pocketsmith account
func getAccounts(cmd *cobra.Command, args []string) {
	cl := pocketsmith.New()
	fmt.Printf("base url %v, apikey %v, user %v,", cl.Config.BaseURL, cl.Config.APIKey, cl.Config.User)
	path := fmt.Sprintf("/users/%v/accounts", cl.Config.User)
	reader, err := cl.SendRequest(path, http.MethodGet, nil)
	if err != nil {
		fmt.Println("woops")
	}
	// The below block needs to be extracted to a json decoder and passed to a "print result" function.
	defer reader.Close()
	body, _ := ioutil.ReadAll(reader)
	fmt.Printf("%s", string(body))
}

func getAuthorisedUser(cmd *cobra.Command, args []string) {
	u := pocketsmith.User{}
	cl := pocketsmith.New()
	path := fmt.Sprint("/me")
	reader, err := cl.SendRequest(path, http.MethodGet, nil)
	if err != nil {
		fmt.Println("woops")
	}
	// The below block needs to be extracted to a json decoder and passed to a "print result" function.
	defer reader.Close()
	body, _ := ioutil.ReadAll(reader)
	err = json.Unmarshal(body, &u)
	if err != nil {
		panic("Was unable to unmarshal user data to user struct.")
	}
	fmt.Println(u.Name)
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
