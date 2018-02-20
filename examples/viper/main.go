package main

import (
	"fmt"

	v "github.com/spf13/viper"
)

func main() {

	viper := v.New()
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("properties")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}

	fmt.Println(viper.GetString("hello"))
}
