package main

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

func main() {
	//home := "viper/config"
	home := "config"
	viper.AddConfigPath(home)
	viper.SetConfigType("toml")
	viper.SetConfigName("conf")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("error:", err)
		panic(err)
	}
	fmt.Println(os.Stderr, "config file: ", viper.ConfigFileUsed())
	fmt.Println(viper.Get("mysql.ip"))
	//var ports interface{}
	ports := viper.GetStringSlice("mysql.port")
	fmt.Println(len(ports))

}
