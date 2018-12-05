package main

import (
	"fmt"
	"github.com/spf13/viper"
)


func config() {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yml")
	// call multiple times to add many search paths
	viper.AddConfigPath("./conf")               // optionally look for config in the working directory
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}

