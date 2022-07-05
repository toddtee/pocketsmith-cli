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
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := app.GetAuthorisedUser(); err != nil {
			return fmt.Errorf("unable to get authorised user: %w", err)
		}
		return nil
	},
}

func init() {
	getCmd.AddCommand(getAuthorisedUserCmd)
}
