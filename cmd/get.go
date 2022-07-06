/*
Package cmd provides a command line application for Pocketsmith.
Copyright © 2022 Todd Turner hi@toddtee.sh

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/toddtee/pocketsmith-cli/internal/app"
)

var authorised bool

// getCmd represents the list command
var getCmd = &cobra.Command{
	Use:   "get", // Need to add a better  [-F file | -D dir]... [-f format] profile
	Short: "retrieves one or many pocketsmith account resources",
	Long:  `prints some import information about the specified pocketsmith resource.`,
}

var getUserCmd = &cobra.Command{
	Use:   "user", // Need to add a better  [-F file | -D dir]... [-f format] profile
	Short: "retrieves the authorised user of the pocketsmith account",
	Long:  `reveals who is the supreme overlord of the household finances.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		c := app.NewClient()
		flag := "authorised"
		b, _ := boolFlagCheck(cmd, flag)
		u, err := c.GetUser(b)
		if err != nil {
			return fmt.Errorf("unable to get user: %w", err)
		}
		printer(u)
		return nil
	},
}

func init() {
	getCmd.AddCommand(getUserCmd)
	getUserCmd.Flags().BoolVarP(&authorised, "authorised", "a", false, "authorised user of account")
}
