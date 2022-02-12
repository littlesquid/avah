package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

func Init() {
	viper.SetConfigName("config")
	viper.AddConfigPath("./config")
	viper.AutomaticEnv()

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("exception on loading config file: %v", err))
	}
}
