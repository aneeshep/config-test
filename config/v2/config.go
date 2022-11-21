package config

import (
	"fmt"

	"github.com/spf13/viper"
)

var (
	ServiceName string
	config      *viper.Viper
)

func LoadConfig() {
	config = viper.New()
	config.SetConfigName("config")
	config.SetConfigType("yaml")
	config.AddConfigPath(fmt.Sprintf("/etc/services/%s", ServiceName))
	config.AddConfigPath(".")
	err := config.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}
}

func GetString(key string) string { return config.GetString(key) }

func Get(key string) interface{} { return config.Get(key) }

func Set(key string, value interface{}) { config.Set(key, value) }
