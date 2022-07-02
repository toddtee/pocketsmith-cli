/*
Copyright © 2022 Todd Turner hi@toddtee.sh

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "pocketsmith",
	Short: "pocketsmith-cli interacts with your pocketsmith account",
	Long: `pocketsmith-cli is a commandline application for interacting with your pocketsmith account.
	
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

//initConfig sets up the configuration for the user
func initConfig() {
	viper.AddConfigPath("$HOME")
	viper.SetConfigType("yaml")
	viper.SetConfigName(".pocketsmith")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Sprintf("Couldn't read in config %v", err))
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.pocketsmith-cli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	viper.BindPFlag("api_key", rootCmd.PersistentFlags().Lookup("api_key"))
	viper.BindPFlag("user", rootCmd.PersistentFlags().Lookup("user"))

	// Config file supported flags with Viper.
	rootCmd.PersistentFlags().StringP("api_key", "a", "", "Developer API Key")
	rootCmd.PersistentFlags().StringP("user", "u", "", "Authorised user ID")
}
