/*
Package cmd provides a command line application for Pocketsmith.
Copyright © 2022 Todd Turner hi@toddtee.sh

*/
package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/toddtee/pocketsmith-cli/internal/app"
)

func getAuthorisedUser(cmd *cobra.Command, args []string) {
	u := app.User{}
	cfg := app.BuildConfig()
	resp := app.AuthorisedUser(cfg)
	defer resp.Body.Close()
	d := getBodyData(resp.Body)

	err := json.Unmarshal(d, &u)
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

var getAuthorisedUserCmd = &cobra.Command{
	Use:   "user", // Need to add a better  [-F file | -D dir]... [-f format] profile
	Short: "retrieves the authorised user of the pocketsmith account",
	Long:  `reveals who is the supreme overlord of the household finances.`,
	Run:   getAuthorisedUser,
}

func init() {
	rootCmd.AddCommand(getCmd)
	getCmd.AddCommand(getAuthorisedUserCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// accountsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// accountsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
