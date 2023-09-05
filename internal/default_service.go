package internal

import (
	"fmt"
	"github.com/spf13/viper"
)

var initiated = false

func init() {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("toml")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")      // optionally look for config in the working directory
	err := viper.ReadInConfig()   // Find and read the config file
	if err != nil {               // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	initiated = true
}

type DefaultService struct{}

func (s *DefaultService) Init() {}
