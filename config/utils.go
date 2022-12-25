package config

import (
	"fmt"
	"path"
	"runtime"

	"github.com/spf13/viper"
)

type ServeConfig struct {
	Port      int
	Host      string
	StaticDir string
}

type Config struct {
	Serve ServeConfig
}

func GetConfig() Config {
	_, filename, _, _ := runtime.Caller(0)

	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AddConfigPath(path.Dir(filename))
	viper.AutomaticEnv()
	viper.SetConfigType("yml")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	var config Config

	err := viper.Unmarshal(&config)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}

	return config
}
