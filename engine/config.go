package engine

import (
	"fmt"
	"github.com/spf13/viper"
)

type ConfigPlugger struct{}

func (c ConfigPlugger) Plug() (err error) {
	viper.SetConfigName("app.yml")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./var/")
	viper.AddConfigPath("/etc/media-cloud/")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("load config error: %s", err)
	}

	return nil
}

func newConfigPlugger() Plugger {
	return &ConfigPlugger{}
}