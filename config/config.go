package config

import (
	"fmt"

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
	if err != nil {
		panic(fmt.Sprintf("Couldn't read in config %v", err))
	}
}
