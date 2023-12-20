package main

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func main() {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	err := viper.ReadInConfig()
	if err != nil {
		log.Println(err)
	}
	hw := viper.Get("helloWorld")
	fmt.Println(hw)
}
