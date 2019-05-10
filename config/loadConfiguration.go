package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

type Selector struct {
	Type    string
	Pattern string
}

type Environment struct {
	Name                     string
	Namespace                string
	TillerNamespace          string
	EnvironmentVariableFiles []string
	ValueFiles               []string
	Selector                 []Selector
}

type Config struct {
	Chart        string
	ValueFiles   []string
	Environments []Environment
}

func LoadConfiguration() *Config {
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Can't read config:", err)
		os.Exit(1)
	}

	keys := viper.AllKeys()
	fmt.Printf("Keys: %d\n", len(keys))

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		fmt.Printf("Cannot process configuration, %s\n", err)

		os.Exit(1)
	}

	return &config
}
