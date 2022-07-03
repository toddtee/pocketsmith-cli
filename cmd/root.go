/*
Package cmd provides a command line application for Pocketsmith.
Copyright © 2022 Todd Turner hi@toddtee.sh

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/toddtee/pocketsmith-cli/config"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "pocketsmith",
	Short: "pocketsmith-cli interacts with your pocketsmith account",
	Long: `pocketsmith-cli is a command line application for interacting with your pocketsmith account.
	
Built in Go and with much love, the pocketsmith-cli allows programmatic access to your pocketsmith account and all that is contained within it. This application is built in conjunction with the pocketsmith api.
Eat it up you finance nerds.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(config.InitConfig)
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringP("config", "c", "", "config file (default is $HOME/.pocketsmith)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	viper.BindPFlag("config", rootCmd.PersistentFlags().Lookup("config"))
}
